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
// Only trees with pages have range nodes. Such a tree creates range nodes
// wherever keys hold one value, and inner nodes below where keys with
// several values crowd it (see settle.go); an inner node below a range node
// starts its path with its byte, like every child of a range node.
//
// Layout: the header, a 256-bit bitmap of the range starts (bit 0 is always
// set), the number of starts in the words before each word, and the
// children in byte order; the child of byte b is child[r.index(b)]. The
// head fills the first 64 bytes, so the children start on a cache line of
// their own. The four classes of 8, 24, 56 and 256 children fill 128, 256,
// 512 and 2112 bytes.
type rhead struct {
	header
	starts [4]uint64
	before [4]uint8 // before[w]: the starts in words 0 to w-1
	_      [4]byte
}

type (
	rnode8 struct {
		rhead
		child [8]*header
	}
	rnode24 struct {
		rhead
		child [24]*header
	}
	rnode56 struct {
		rhead
		child [56]*header
	}
	rnode256 struct {
		rhead
		child [256]*header
	}
)

// rCaps are the capacities of the range node classes; a node shrinks into
// the next smaller class once it holds rShrink ranges or fewer.
var (
	rCaps   = [4]int{8, 24, 56, 256}
	rShrink = [4]int{0, 5, 17, 44}
)

// rChildOff is the offset of the children in every range node class.
const rChildOff = unsafe.Sizeof(rhead{})

func asR(h *header) *rhead { return (*rhead)(unsafe.Pointer(h)) }

func newR(class int) *rhead {
	var r *rhead
	switch class {
	case 0:
		r = &(&rnode8{}).rhead
	case 1:
		r = &(&rnode24{}).rhead
	case 2:
		r = &(&rnode56{}).rhead
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

// index returns the index of the range that holds byte b: the range starts
// up to b, less one. It counts one word, with the count of the words before
// it at hand, instead of looping over them: how many words come before b's
// is as random as b, and that loop's branch would mispredict on most
// lookups.
func (r *rhead) index(b byte) int {
	w := b >> 6
	return int(r.before[w&3]) + bits.OnesCount64(r.starts[w&3]&(uint64(2)<<(b&63)-1)) - 1
}

// recount updates before after starts changed.
func (r *rhead) recount() {
	c := 0
	for w, set := range r.starts {
		r.before[w] = uint8(c)
		c += bits.OnesCount64(set)
	}
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
	r.recount()
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
	r.recount()
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
		r.recount()
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
// maxKeys). Such a merged page always fits, and pages thinned out by deletes
// do not stay behind half empty. A page that has just split holds about half
// of that already, so the two halves do not merge back after a few deletes;
// and the check costs nothing on most deletes, which leave their page well
// above it. It returns n or its smaller replacement.
func rMerge(n *header, i int) *header {
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
		ch[min(i, j)] = pageHdr(pageFor(append(pageItems(asPage(a)), pageItems(asPage(b))...)))
		return rRemove(n, max(i, j))
	}
	return n
}

// isPage reports whether n is a page of any type.
func isPage(n *header) bool { return n.kind != kLeaf && n.kind <= kLastPage }

// maxKeys returns how many keys the largest page of p's type holds.
func (p *pageHead) maxKeys() int {
	if p.kind == kPage {
		return pageCaps[len(pageCaps)-1]
	}
	return kCaps[len(kCaps)-1]
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
