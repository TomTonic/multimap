package art

import (
	"math/bits"
	"slices"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// Range nodes keep pages full.
//
// Below a node that branches on single bytes, a page that overflows bursts
// into one page per byte of its keys. For keys whose next bytes are spread
// out, most of these pages hold a key or two, and they fill up only as slowly
// as keys arrive for their byte: how full the pages are depends on the number
// of keys, not on the page size.
//
// A range node gives each child a range of bytes instead: child i holds the
// keys whose byte at the node's branch position lies in [start_i,
// start_{i+1}). A page that overflows splits by count, at a byte boundary,
// into pages that share its range, like a B-tree leaf; only a page whose keys
// all share that byte gets a subtree of its own. This follows the B-trie of
// Askitis and Zobel (VLDB Journal 18, 2009), whose buckets are referenced by
// ranges of trie pointers.
//
// Unlike an inner node, a range node does not consume the byte it branches
// on: its children start at the same depth and check that byte themselves.
// Pages and leaves hold full keys, and a node below a range node has that
// byte as the first byte of its path. So a child's range may be wider than
// its keys, and no range needs an empty child: a key whose byte falls into a
// range belongs to that range's child, and a lookup that finds nothing there
// is a miss.
//
// Only trees with pages have range nodes, and they have no other inner
// nodes: every node such a tree creates is a range node.
//
// Layout: the header, a 256-bit bitmap of the range starts (bit 0 is always
// set), and the children in byte order; the child of byte b is
// child[swar.Floor(starts, b)]. The four classes of 9, 25, 57 and 256
// children fill 128, 256, 512 and 2104 bytes.
type rhead struct {
	header
	starts [4]uint64
}

type (
	rnode9 struct {
		rhead
		child [9]*header
	}
	rnode25 struct {
		rhead
		child [25]*header
	}
	rnode57 struct {
		rhead
		child [57]*header
	}
	rnode256 struct {
		rhead
		child [256]*header
	}
)

// rCaps are the capacities of the range node classes; a node shrinks into
// the next smaller class once it holds rShrink ranges or fewer.
var (
	rCaps   = [4]int{9, 25, 57, 256}
	rShrink = [4]int{0, 6, 18, 44}
)

// rChildOff is the offset of the children in every range node class.
const rChildOff = unsafe.Sizeof(rhead{})

func asR(h *header) *rhead { return (*rhead)(unsafe.Pointer(h)) }

func newR(class int) *rhead {
	var r *rhead
	switch class {
	case 0:
		r = &(&rnode9{}).rhead
	case 1:
		r = &(&rnode25{}).rhead
	case 2:
		r = &(&rnode57{}).rhead
	default:
		r = &(&rnode256{}).rhead
	}
	r.kind, r.class = kR, uint8(class)
	return r
}

// children returns the child slots of r, as many as its class holds; the
// first count are in use.
func (r *rhead) children() []*header {
	return unsafe.Slice((**header)(unsafe.Add(unsafe.Pointer(r), rChildOff)), rCaps[r.class])
}

// start returns the first byte of range k.
func (r *rhead) start(k int) byte {
	w := 0
	for k >= bits.OnesCount64(r.starts[w]) {
		k -= bits.OnesCount64(r.starts[w])
		w++
	}
	set := r.starts[w]
	for ; k > 0; k-- {
		set &= set - 1
	}
	return byte(w<<6 + bits.TrailingZeros64(set))
}

// rng is one range of a range node while it is built: its first byte and
// its child.
type rng struct {
	b byte
	c *header
}

// ranges returns the ranges of r in order.
func (r *rhead) ranges() []rng {
	out := make([]rng, 0, r.count)
	ch := r.children()
	for w, set := range r.starts {
		for ; set != 0; set &= set - 1 {
			out = append(out, rng{byte(w<<6 + bits.TrailingZeros64(set)), ch[len(out)]})
		}
	}
	return out
}

// makeR returns a new range node with the path and term of h and the ranges
// rs, of which the first must start at 0.
func makeR(h header, rs []rng) *header {
	c := 0
	for rCaps[c] < len(rs) {
		c++
	}
	r := newR(c)
	h.kind, h.class, h.count = kR, r.class, uint16(len(rs))
	r.header = h
	ch := r.children()
	for i, x := range rs {
		swar.Set(&r.starts, x.b)
		ch[i] = x.c
	}
	return &r.header
}

// rSplice replaces range i of n by rs, which cover the same bytes: rs[0]
// takes over the range's start, the others start within the range. It
// returns n or its larger replacement.
func rSplice(n *header, i int, rs []rng) *header {
	r := asR(n)
	cnt := int(r.count) + len(rs) - 1
	if cnt > rCaps[r.class] {
		all := r.ranges()
		rs[0].b = all[i].b
		return makeR(r.header, slices.Replace(all, i, i+1, rs...))
	}
	ch := r.children()
	copy(ch[i+len(rs):cnt], ch[i+1:r.count])
	for k, x := range rs {
		ch[i+k] = x.c
		if k > 0 {
			swar.Set(&r.starts, x.b)
		}
	}
	r.count = uint16(cnt)
	return n
}

// rRemove removes range i of n, whose child is gone. Its bytes go to the
// range before it, or, for the first range, to the one after it. It returns
// n or its smaller replacement.
func rRemove(n *header, i int) *header {
	r := asR(n)
	cnt := int(r.count) - 1
	if cnt > 0 {
		s := r.start(max(i, 1))
		r.starts[s>>6] &^= uint64(1) << (s & 63)
	}
	ch := r.children()
	copy(ch[i:cnt], ch[i+1:cnt+1])
	ch[cnt] = nil
	r.count = uint16(cnt)
	if c := int(r.class); c > 0 && cnt <= rShrink[c] {
		return makeR(r.header, r.ranges())
	}
	return n
}

// rMerge merges the page of range i of n into a neighbouring page when the
// two hold at most half the keys the largest page of either type holds (see
// maxKeys) and fit a page below the largest class, so that pages thinned out
// by deletes do not stay behind half empty. A page that has just split holds
// about half of that already, so the two halves do not merge back after a
// few deletes; and the check costs nothing on most deletes, which leave
// their page well above it.
func (t *Tree) rMerge(n *header, i int) *header {
	r := asR(n)
	ch := r.children()
	for _, j := range [2]int{i - 1, i + 1} {
		if j < 0 || j >= int(r.count) {
			continue
		}
		a, b := ch[min(i, j)], ch[max(i, j)]
		if !isPage(a) || !isPage(b) ||
			2*(int(asPage(a).count)+int(asPage(b).count)) > min(asPage(a).maxKeys(), asPage(b).maxKeys()) {
			continue
		}
		q := t.pageFor(append(pageItems(asPage(a)), pageItems(asPage(b))...))
		if q == nil || q.largest() {
			continue
		}
		ch[min(i, j)] = pageHdr(q)
		return rRemove(n, max(i, j))
	}
	return n
}

// isPage reports whether n is a page of any type.
func isPage(n *header) bool { return n.kind != kLeaf && n.kind <= kLastPage }

// maxKeys returns about how many keys the largest page of p's type holds:
// for an S page, whose capacity depends on its keys, as many as its own
// capacity scaled to the largest class.
func (p *pageHead) maxKeys() int {
	switch p.kind {
	case kPage:
		return pageCaps[len(pageCaps)-1]
	case kPageK:
		return kCaps[len(kCaps)-1]
	case kPageN:
		return nLayouts[len(nLayouts)-1].keys()
	}
	return int(p.kcap) * sSizes[len(sSizes)-1] / sSizes[p.class]
}

// largest reports whether p is of the largest class of its type.
func (p *pageHead) largest() bool {
	if p.kind == kPageN {
		return int(p.class) == len(nLayouts)-1
	}
	return p.class == 3
}

// ranges appends to out the ranges that hold items, which are sorted,
// distinct, share their first d bytes and are all longer than that, and
// returns out. Items that fit a page become one; otherwise they split in two
// at the byte boundary nearest their middle, and each half is split the same
// way. Items that all share byte d and do not fit a page become a subtree of
// their own. The first range starts at the byte of the first item.
func (t *Tree) ranges(items []item, d int, out []rng) []rng {
	n := len(items)
	first, last := items[0].key[d], items[n-1].key[d]
	if first == last || n == 1 {
		return append(out, rng{first, t.build(items, d)})
	}
	if p := t.pageFor(items); p != nil {
		return append(out, rng{first, pageHdr(p)})
	}
	m := n / 2
	b := items[m].key[d]
	lo, hi := m, m
	for lo > 0 && items[lo-1].key[d] == b {
		lo--
	}
	for hi < n && items[hi].key[d] == b {
		hi++
	}
	s := lo
	if lo == 0 || (hi < n && hi-m < m-lo) {
		s = hi
	}
	return t.ranges(items[s:], d, t.ranges(items[:s], d, out))
}
