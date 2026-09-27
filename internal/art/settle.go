package art

import "github.com/TomTonic/multimap/internal/swar"

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
// node below a range node starts its path with the byte it sits under, like
// every child of a range node, so the two kinds of subtree mix freely. There
// is no way back to pages, except that a subtree that falls to a single leaf
// takes new keys in pages again.

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
	for n := *loc; n.kind == kR; n = *loc {
		path = append(path, rpos{loc, depth})
		depth += int(n.plen)
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
		*p.loc = t.inner(t.items(*p.loc, nil), p.depth)
	}
}

// crowded reports whether the leaves and inner-node subtrees below r
// outweigh its pages and range-node subtrees, with gained of the range
// nodes about to fall back (see crowdedRatio).
func crowded(r *rhead, gained int) bool {
	multi, keys, ranges := gained, 0, -gained
	for _, c := range r.children()[:r.count] {
		switch {
		case isPage(c):
			keys += int(asPage(c).count)
		case c.kind == kR:
			ranges++
		default:
			multi++
		}
	}
	return crowdedRatio*multi > keys+crowdedRatio*ranges
}

// items appends the keys below n to out in order, with their leaves or, for
// keys in pages, their raw values, and returns out.
func (t *Tree) items(n *header, out []item) []item {
	switch {
	case n.kind == kLeaf:
		return append(out, item{key: asLeaf(n).key(), leaf: asLeaf(n)})
	case isPage(n):
		return append(out, pageItems(asPage(n))...)
	}
	if n.term != nil {
		out = append(out, item{key: n.term.key(), leaf: n.term})
	}
	if n.kind == kR {
		for _, c := range asR(n).children()[:n.count] {
			out = t.items(c, out)
		}
		return out
	}
	eachInner(n, func(c *header) { out = t.items(c, out) })
	return out
}

// eachInner calls fn for every child of the inner node n in byte order.
func eachInner(n *header, fn func(*header)) {
	switch n.kind {
	case kN4, kN11:
		_, child := sorted(n)
		for _, c := range child {
			fn(c)
		}
	case kN25, kN57:
		_, child := bitmapOf(n)
		for _, c := range child[:n.count] {
			fn(c)
		}
	default:
		for _, c := range asN256(n).child {
			if c != nil {
				fn(c)
			}
		}
	}
}

// inner returns a subtree of inner nodes and leaves that holds items, which
// are sorted, distinct and share their first depth bytes, as a tree without
// pages holds them. Keys from pages get a leaf with their value; keys with a
// leaf keep it.
func (t *Tree) inner(items []item, depth int) *header {
	if len(items) == 1 {
		return leafHdr(t.leafOf(items[0]))
	}
	first, last := items[0].key, items[len(items)-1].key
	plen := swar.Lcp(first[depth:], last[depth:])
	d := depth + plen
	var h header
	h.setPrefix(first[depth:d], plen)
	if len(first) == d {
		// A key that ends here sorts first.
		h.term = t.leafOf(items[0])
		items = items[1:]
	}
	groups := 1
	for i := 1; i < len(items); i++ {
		groups += b2i(items[i].key[d] != items[i-1].key[d])
	}
	n := newInner(h, groups)
	for i := 0; i < len(items); {
		j, b := i+1, items[i].key[d]
		for j < len(items) && items[j].key[d] == b {
			j++
		}
		n = addChild(n, b, t.inner(items[i:j], d+1))
		i = j
	}
	return n
}

// leafOf returns the leaf of it, making one for a key from a page.
func (t *Tree) leafOf(it item) *leafHead {
	if it.leaf != nil {
		return it.leaf
	}
	return t.mk(it)
}

// newInner returns an empty inner node of the smallest kind that holds count
// children, with the path and term of h.
func newInner(h header, count int) *header {
	var n *header
	switch {
	case count <= 4:
		n, h.kind = &(&node4{}).header, kN4
	case count <= 11:
		n, h.kind = &(&node11{}).header, kN11
	case count <= 25:
		n, h.kind = &(&node25{}).header, kN25
	case count <= 57:
		n, h.kind = &(&node57{}).header, kN57
	default:
		n, h.kind = &(&node256{}).header, kN256
	}
	*n = h
	return n
}
