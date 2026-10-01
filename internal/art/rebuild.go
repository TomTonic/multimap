package art

import (
	"math/bits"
	"slices"

	"github.com/TomTonic/multimap/internal/swar"
)

// item is one key during a rebuild of a subtree (see build): a key with its
// one raw value, which may go into a page, or a key that has a leaf.
type item struct {
	key  []byte    // the whole key
	val  uint64    // the key's value, if it has no leaf
	leaf *leafHead // holds key from a base of some depth on, or nil
	full []byte    // an immutable copy of key a K page may keep, or nil
}

// leafOf returns the leaf of it, ready to stand at depth: the leaf it has, or a
// new one for a key from a page. A leaf never lands above its base: the
// subtrees that are rebuilt around leaves hold them in inner nodes only (see
// settle), and in a tree of inner nodes a leaf stands where its key first
// differs from every other key. Every node above a leaf branches on a byte its
// key shares with another one, so the leaf stood no deeper before.
func (t *Tree) leafOf(it item, depth int) *leafHead {
	if it.leaf == nil {
		return t.mk(it.key, depth, it.val)
	}
	return it.leaf
}

// build returns a subtree holding exactly items, which are sorted by key,
// distinct, and all start with the same depth bytes. It is how pages burst
// and how pages change type when nothing simpler fits: the subtree is rebuilt
// from its keys. Items that fit a page become one; otherwise a range node
// takes their common path and splits them into ranges (see ranges).
func (t *Tree) build(items []item, depth int) *header {
	if len(items) == 1 && items[0].leaf != nil {
		return leafHdr(t.leafOf(items[0], depth))
	}
	if p := pageFor(items); p != nil {
		return pageHdr(p)
	}
	first, last := items[0].key, items[len(items)-1].key
	plen := swar.Lcp(first[depth:], last[depth:])
	d := depth + plen
	var term *leafHead
	if len(first) == d {
		// A key that ends here sorts first.
		term = t.leafOf(items[0], d)
		items = items[1:]
	}
	rs := t.ranges(items, d, nil)
	rs[0].b = 0
	return makeR(first[depth:d], term, rs)
}

// ranges appends to out the ranges that hold items, which are sorted,
// distinct, share their first d bytes and are all longer than that, and
// returns out. Items that fit a page become one. Otherwise the byte group of
// the first key with a leaf gets a range of its own, so that the keys around
// it fill pages as large as they can; without leaves, the items split in two
// at the byte boundary nearest their middle, and each half is split the same
// way. Items that all share byte d and do not fit a page become a subtree of
// their own. The first range starts at the byte of the first item.
func (t *Tree) ranges(items []item, d int, out []rng) []rng {
	n := len(items)
	first, last := items[0].key[d], items[n-1].key[d]
	if first == last {
		return append(out, rng{first, t.build(items, d)})
	}
	if p := pageFor(items); p != nil {
		return append(out, rng{first, pageHdr(p)})
	}
	k := slices.IndexFunc(items, func(it item) bool { return it.leaf != nil })
	lo := n / 2
	if k >= 0 {
		lo = k
	}
	b, hi := items[lo].key[d], lo
	for lo > 0 && items[lo-1].key[d] == b {
		lo--
	}
	for hi < n && items[hi].key[d] == b {
		hi++
	}
	if k >= 0 {
		// the leaf's group, between the keys before and after it
		if lo > 0 {
			out = t.ranges(items[:lo], d, out)
		}
		out = append(out, rng{b, t.build(items[lo:hi], d)})
		if hi < n {
			out = t.ranges(items[hi:], d, out)
		}
		return out
	}
	m, s := n/2, lo
	if lo == 0 || (hi < n && hi-m < m-lo) {
		s = hi
	}
	return t.ranges(items[s:], d, t.ranges(items[:s], d, out))
}

// Pages pay off only where keys hold one value. Where many keys hold several,
// their leaves cut the pages between them into pieces of a key or two, and
// range nodes cost more per level than inner nodes: such a subtree is
// faster, and no larger, as a tree without pages.
//
// So a tree with pages starts optimistic and falls back where its keys turn
// out to hold several values. Every key that gets a second value leaves its
// page for a leaf (see promote). Then the range nodes on its path are
// checked from the bottom up: once the leaves and inner-node subtrees below
// a range node outweigh its pages (see crowded), the range node and its
// subtree are rebuilt from inner nodes and leaves, as a tree without pages
// would hold them. A parent that gains such a subtree is checked the same
// way, so the fallback spreads upwards as far as the keys call for it.
//
// Below an inner node, a tree never holds pages or range nodes: a new key
// there gets a leaf, and splits make inner nodes (see upsert). An inner
// node below a range node starts its path with its byte, like every child of a
// range node, so the two kinds of subtree mix freely. There is no way back to
// pages, except that a subtree that falls to a single leaf takes new keys in
// pages again.

// crowdedRatio is how many keys in pages one leaf or inner-node subtree below
// a range node outweighs: a range node falls back once more than about one
// in crowdedRatio+1 of its keys holds several values.
const crowdedRatio = 4

// rpos is a range node on the path of a key: its slot and the key depth its
// path starts at.
type rpos struct {
	loc   **header
	depth int
}

// settle rebuilds the subtree around key, which has just got a leaf, from
// inner nodes when keys with several values crowd it: the topmost range
// node on the key's path that is crowded, counting the subtrees below it
// that fall back too.
func (t *Tree) settle(key []byte) {
	var buf [16]rpos
	path := buf[:0]
	loc, depth := &t.root, 0
	for n := *loc; isRange(n.kind); n = *loc {
		path = append(path, rpos{loc, depth})
		depth += n.pathLen()
		if depth == len(key) {
			break // the key's leaf is n's term
		}
		x := asR(n)
		loc = &x.children()[x.index(key[depth])]
	}
	top, gained := -1, 0
	for i := len(path) - 1; i >= 0; i-- {
		if !crowded(asR(*path[i].loc), gained) {
			break
		}
		top, gained = i, 1
	}
	if top >= 0 {
		p := path[top]
		*p.loc = t.inner(t.items(*p.loc, p.depth, key[:p.depth]), p.depth)
	}
}

// crowded reports whether the leaves and inner-node subtrees below r
// outweigh its pages and range-node subtrees, with gained of the range
// nodes about to fall back (see crowdedRatio).
func crowded(r *rhead, gained int) bool {
	multi, keys, ranges := gained, 0, -gained
	for _, c := range r.children()[:r.n] {
		switch {
		case isPage(c.kind):
			keys += int(asPage(c).count)
		case isRange(c.kind):
			ranges++
		default:
			multi++
		}
	}
	return crowdedRatio*multi > keys+crowdedRatio*ranges
}

// walker collects the keys below a node as items.
type walker struct {
	path  []byte // the key bytes above the node being visited
	arena []byte // the keys of the leaves, one after another
	out   []item
}

// items returns the keys below n, whose path starts at depth, in order, with
// their leaves or, for keys in pages, their raw values; pre is the key bytes
// before depth.
func (t *Tree) items(n *header, depth int, pre []byte) []item {
	w := walker{path: slices.Clone(pre)}
	w.walk(n, depth)
	return w.out
}

// leaf appends the item of leaf l, which holds its key from a base within the
// current path.
func (w *walker) leaf(l *leafHead) {
	start := len(w.arena)
	w.arena = append(append(w.arena, w.path[:l.base()]...), l.stored()...)
	k := w.arena[start:len(w.arena):len(w.arena)]
	w.out = append(w.out, item{key: k, leaf: l})
}

// walk appends the items of the subtree n, whose path starts at depth, to
// w.out.
func (w *walker) walk(n *header, depth int) {
	switch {
	case isLeaf(n.kind):
		w.leaf(asLeaf(n))
		return
	case isPage(n.kind):
		w.out = append(w.out, pageItems(asPage(n))...)
		return
	}
	w.path = appendPath(w.path[:depth], n)
	depth += n.pathLen()
	if t := termOf(n); t != nil {
		w.leaf(t)
	}
	if isRange(n.kind) {
		for _, c := range asR(n).children()[:asR(n).n] {
			w.walk(c, depth)
		}
		return
	}
	eachInner(n, func(b byte, c *header) {
		w.path = append(w.path[:depth], b)
		w.walk(c, depth+1)
	})
}

// eachInner calls fn for every byte child of the inner node n in byte order.
func eachInner(n *header, fn func(b byte, c *header)) {
	switch n.kind {
	case kN5, kN12:
		keys, child := sorted(n)
		for i, c := range child {
			fn(keys[i], c)
		}
	case kN26, kN58:
		bm, child := bitmapOf(n)
		i := 0
		for w, set := range bm {
			for ; set != 0; set &= set - 1 {
				fn(byte(w<<6+bits.TrailingZeros64(set)), child[i])
				i++
			}
		}
	default:
		for b, c := range asN256(n).child[:256] {
			if c != nil {
				fn(byte(b), c)
			}
		}
	}
}

// innerKind returns the smallest kind of inner node with room for groups byte
// children and, if term, a term leaf.
func innerKind(groups int, term bool) kind {
	switch need := groups + b2i(term); {
	case need <= 5:
		return kN5
	case need <= 12:
		return kN12
	case need <= 26:
		return kN26
	case need <= 58:
		return kN58
	}
	return kN256
}

// inner returns a subtree of inner nodes and leaves that holds items, which
// are sorted, distinct and share their first depth bytes, as a tree without
// pages holds them. Keys from pages get a leaf with their value; keys with a
// leaf keep it.
func (t *Tree) inner(items []item, depth int) *header {
	if len(items) == 1 {
		return leafHdr(t.leafOf(items[0], depth))
	}
	first, last := items[0].key, items[len(items)-1].key
	plen := swar.Lcp(first[depth:], last[depth:])
	d := depth + plen
	var term *leafHead
	if len(first) == d {
		// A key that ends here sorts first.
		term = t.leafOf(items[0], d)
		items = items[1:]
	}
	groups := 1
	for i := 1; i < len(items); i++ {
		groups += b2i(items[i].key[d] != items[i-1].key[d])
	}
	n := newNode(innerKind(groups, term != nil), plen)
	storePath(n, first[depth:d])
	setTermSlot(n, term)
	for i := 0; i < len(items); {
		j, b := i+1, items[i].key[d]
		for j < len(items) && items[j].key[d] == b {
			j++
		}
		n, _ = addChild(n, b, t.inner(items[i:j], d+1))
		i = j
	}
	return n
}
