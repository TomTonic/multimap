package art

import (
	"math/bits"
	"slices"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
	"github.com/TomTonic/multimap/internal/vpage"
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
// Pages hold full keys, a leaf holds its key from a base that is at most that
// depth, and a node below a range node has that byte as the first byte of its
// path. So a child's range may be wider than its keys, and no range needs an
// empty child: a key whose byte falls into a range belongs to that range's
// child, and a lookup that finds nothing there is a miss.
//
// Only trees with pages have range nodes. Such a tree creates range nodes
// wherever keys hold one value, and inner nodes below where keys with several
// values crowd it (see settle.go); an inner node below a range node starts its
// path with its byte, like every child of a range node.
//
// Layout: the header, a 256-bit bitmap of the range starts (bit 0 is always
// set), the number of starts in the words before each word, and the children
// in byte order; the child of byte b is child[r.index(b)]. After the last
// child comes the slot of the term leaf, as in every node (see termOf). The
// head fills 56 bytes, so the four classes of 8, 24, 56 and 256 ranges fill
// 128, 256, 512 and 2112 bytes, three of them Go size classes. A range node
// keeps header.count at 255, as the 256-way node does, and its number of
// ranges in n.
type rhead struct {
	header
	starts [4]uint64
	before [4]uint8 // before[w]: the starts in words 0 to w-1
	n      uint16   // ranges
	_      [2]byte
}

type (
	rnode8 struct {
		rhead
		child [9]*header
	}
	rnode24 struct {
		rhead
		child [25]*header
	}
	rnode56 struct {
		rhead
		child [57]*header
	}
	rnode256 struct {
		rhead
		child [257]*header
	}
)

// rCaps are the numbers of ranges the classes of range nodes hold, indexed by
// kind - kR8; a node shrinks into the next smaller class once it holds
// rShrink ranges or fewer.
var (
	rCaps   = [4]int{8, 24, 56, 256}
	rShrink = [4]int{0, 5, 17, 44}
)

// rChildOff is the offset of the children in every range node class.
const rChildOff = unsafe.Sizeof(rhead{})

func asR(h *header) *rhead { return (*rhead)(unsafe.Pointer(h)) }

// class returns the index of r's class in rCaps.
func (r *rhead) class() int { return int(r.kind-kR8) & 3 }

// children returns the child slots of r, as many as its class holds; the
// first n are in use.
func (r *rhead) children() []*header {
	return unsafe.Slice((**header)(unsafe.Add(unsafe.Pointer(r), rChildOff)), rCaps[r.class()])
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

// rng is one range of a range node while it is built: its first byte and its
// child.
type rng struct {
	b byte
	c *header
}

// ranges returns the ranges of r in order.
func (r *rhead) ranges() []rng {
	out := make([]rng, 0, r.n)
	ch := r.children()
	for w, set := range r.starts {
		for ; set != 0; set &= set - 1 {
			out = append(out, rng{byte(w<<6 + bits.TrailingZeros64(set)), ch[len(out)]})
		}
	}
	return out
}

// rClass returns the smallest class that holds n ranges.
func rClass(n int) int {
	c := 0
	for rCaps[c] < n {
		c++
	}
	return c
}

// makeR returns a new range node with the path p and the term leaf term, which
// may be nil, and the ranges rs, of which the first must start at 0.
func makeR(p []byte, term *leafHead, rs []rng) *header {
	n := newNode(kR8+kind(rClass(len(rs))), len(p))
	storePath(n, p)
	fillR(n, rs)
	setTermSlot(n, term)
	return n
}

// fillR stores the ranges rs in the new range node n, which has room for them.
func fillR(n *header, rs []rng) {
	r := asR(n)
	r.n = uint16(len(rs))
	ch := r.children()
	for i, x := range rs {
		swar.Set(&r.starts, x.b)
		ch[i] = x.c
	}
	r.recount()
}

// remakeR returns a range node with n's path and term and the ranges rs, in the
// smallest class that holds them.
func remakeR(n *header, rs []rng) *header {
	y := newLike(n, kR8+kind(rClass(len(rs))))
	fillR(y, rs)
	setTermSlot(y, termOf(n))
	return y
}

// rSplice replaces range i of n by rs, which cover the same bytes: rs[0]
// takes over the range's start, the others start within the range. It
// returns n or its larger replacement.
func rSplice(n *header, i int, rs []rng) *header {
	r := asR(n)
	cnt := int(r.n) + len(rs) - 1
	if cnt > rCaps[r.class()] {
		all := r.ranges()
		rs[0].b = all[i].b
		return remakeR(n, slices.Replace(all, i, i+1, rs...))
	}
	ch := r.children()
	copy(ch[i+len(rs):cnt], ch[i+1:r.n])
	for k, x := range rs {
		ch[i+k] = x.c
		if k > 0 {
			swar.Set(&r.starts, x.b)
		}
	}
	r.n = uint16(cnt)
	r.recount()
	return n
}

// rRemove removes range i of n, whose child is gone. Its bytes go to the
// range before it, or, for the first range, to the one after it. It returns
// n or its smaller replacement.
func rRemove(n *header, i int) *header {
	r := asR(n)
	cnt := int(r.n) - 1
	if cnt > 0 {
		s := r.start(max(i, 1))
		r.starts[s>>6] &^= uint64(1) << (s & 63)
		r.recount()
	}
	ch := r.children()
	copy(ch[i:cnt], ch[i+1:cnt+1])
	ch[cnt] = nil
	r.n = uint16(cnt)
	if c := r.class(); c > 0 && cnt <= rShrink[c] {
		return remakeR(n, r.ranges())
	}
	return n
}

// rMerge merges the page of range i of n into a neighbouring page when the
// two fit one page with room to spare (see vpage.Merge), so that pages thinned
// out by deletes do not stay behind half empty. A page that has just split holds
// about half of that already, so the two halves do not merge back after a few
// deletes; and the check costs nothing on most deletes, which leave their page
// well above it. It returns n or its smaller replacement.
func rMerge(n *header, i int) *header {
	r := asR(n)
	ch := r.children()
	for _, j := range [2]int{i - 1, i + 1} {
		if j < 0 || j >= int(r.n) {
			continue
		}
		a, b := ch[min(i, j)], ch[max(i, j)]
		if !isPage(a.kind) || !isPage(b.kind) || asPage(a).Base() != asPage(b).Base() {
			continue
		}
		if m := vpage.Merge(asPage(a), asPage(b)); m != nil {
			ch[min(i, j)] = pageHdr(m)
			return rRemove(n, max(i, j))
		}
	}
	return n
}
