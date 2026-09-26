package art

import (
	"encoding/binary"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// A page holds all keys of one subtree in a single object of 64, 128, 256 or
// 512 bytes: Go size classes that are aligned to their size and carry no
// malloc header. It replaces a leaf per key, and the pointer to each leaf, by
// sorted arrays, so that a lookup ends in one object and a range scan walks
// contiguous memory.
//
// Pages hold values of a small pointer-free type only (see Tree.small), and
// only keys with exactly one value keep it in the page. There are two page
// types:
//
//   - U8-1 (this file): keys of one common length of at most 8 bytes (integer
//     keys are 8), each with exactly one value. The keys are stored as
//     big-endian words; keys and values are plain words, so the page contains
//     no pointers and the garbage collector never scans it.
//   - K (pagek.go): keys of any lengths up to maxPageKey, each with exactly
//     one value, in a fixed layout.
//
// A key with more than one value has a generic leaf with a value set, as in a
// tree without pages, which a range node holds like a page (see rnode.go).
//
// Every page can rebuild its full keys, so it never depends on where in the
// tree it sits: when a node above it collapses, the page just moves up.

// pageCaps are the capacities of the four page classes.
var pageCaps = [4]int{3, 7, 15, 31}

// maxPageKey is the longest key a page holds; a longer key gets a leaf.
const maxPageKey = 255

// keyBuf holds a key that a page rebuilds from its parts.
type keyBuf [maxPageKey + 16]byte

// pageHead is the start of every page (8 B).
type pageHead struct {
	kind  kind
	class uint8  // index into pageCaps (U8-1) or kCaps (K)
	count uint8  // keys
	klen  uint8 // U8-1: length of every key in the page, 0..8; K: see pageKHead
	_     uint16
	kcap  uint8 // K: key slots
	base  uint8 // K: length of the prefix all keys share
}

// The page classes: 56, 120, 248 and 504 bytes, allocated as 64, 128, 256
// and 512. heads holds the keys in ascending order, vals the value of each
// key as raw bits.
type (
	page3 struct {
		pageHead
		heads, vals [3]uint64
	}
	page7 struct {
		pageHead
		heads, vals [7]uint64
	}
	page15 struct {
		pageHead
		heads, vals [15]uint64
	}
	page31 struct {
		pageHead
		heads, vals [31]uint64
	}
)

// headsOff is the offset of heads in every page class; vals follows heads.
const headsOff = unsafe.Sizeof(pageHead{})

// Pages shrink into the next smaller class once they hold this many keys or
// fewer, well below that class's capacity, so that a page that has just
// grown does not shrink back after one removal.
var pageShrink = [4]int{0, 2, 4, 10}

func asPage(h *header) *pageHead  { return (*pageHead)(unsafe.Pointer(h)) }
func pageHdr(p *pageHead) *header { return (*header)(unsafe.Pointer(p)) }

func newPage(class int) *pageHead {
	var p *pageHead
	switch class {
	case 0:
		p = &(&page3{}).pageHead
	case 1:
		p = &(&page7{}).pageHead
	case 2:
		p = &(&page15{}).pageHead
	default:
		p = &(&page31{}).pageHead
	}
	p.kind, p.class = kPage, uint8(class)
	return p
}

// keys returns the keys of a U8-1 page in use, in ascending order.
func (p *pageHead) keys() []uint64 {
	return unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), headsOff)), p.count)
}

// heads returns the keys of a U8-1 page; vals their values. Both slices span
// the page's capacity, of which the first count entries are in use.
func (p *pageHead) heads() []uint64 {
	return unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), headsOff)), pageCaps[p.class])
}

func (p *pageHead) vals() []uint64 {
	c := pageCaps[p.class]
	return unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), headsOff+8*uintptr(c))), c)
}

// keyWord turns a key of at most 8 bytes into the word a page stores: big
// endian and zero-padded, so that among keys of one length the words compare
// like the keys.
func keyWord(k []byte) uint64 {
	if len(k) == 8 {
		return binary.BigEndian.Uint64(k)
	}
	var b [8]byte
	copy(b[:], k)
	return binary.BigEndian.Uint64(b[:])
}

// wordKey is the inverse of keyWord for keys of length l; buf backs the result.
func wordKey(w uint64, l int, buf *[8]byte) []byte {
	binary.BigEndian.PutUint64(buf[:], w)
	return buf[:l]
}

// search returns the position of key w in a U8 page, or where it would be
// inserted, and whether it is there.
func (p *pageHead) search(w uint64) (int, bool) { return search(p.keys(), w) }

// search returns the position of the first word in h, which is sorted, that
// is not below w, and whether it is w.
//
// Heads 0-14 fill the page's first 128 bytes, one cache line on Apple
// silicon. One comparison with head 14 picks the line that holds w, and a
// linear search within that line follows. Measured against a branch-free
// binary search, this is 1.4x as fast while the tree is cached and as fast
// once it is not; a linear search over all heads is as fast in the cache but
// loses 25-35% without it, and counting all heads without branching is
// slower in both.
func search(h []uint64, w uint64) (int, bool) {
	i := 0
	if len(h) > 15 && h[14] < w {
		i = 15
	}
	for i < len(h) && h[i] < w {
		i++
	}
	return i, i < len(h) && h[i] == w
}

// insertAt inserts key w with value v at position i, growing the page into
// the next class when it is full. It returns the page or its replacement.
// The page must not hold 31 keys already.
func (p *pageHead) insertAt(i int, w, v uint64) *pageHead {
	if int(p.count) == pageCaps[p.class] {
		p = p.resize(int(p.class) + 1)
	}
	n := int(p.count)
	h, vs := p.heads(), p.vals()
	copy(h[i+1:n+1], h[i:n])
	copy(vs[i+1:n+1], vs[i:n])
	h[i], vs[i] = w, v
	p.count++
	return p
}

// removeAt removes the key at position i and returns the page, its smaller
// replacement once it holds few enough keys, or nil once it is empty.
func (p *pageHead) removeAt(i int) *pageHead {
	n := int(p.count)
	if n == 1 {
		return nil
	}
	h, vs := p.heads(), p.vals()
	copy(h[i:n-1], h[i+1:n])
	copy(vs[i:n-1], vs[i+1:n])
	p.count--
	if n-1 <= pageShrink[p.class] {
		return p.resize(int(p.class) - 1)
	}
	return p
}

// split moves the keys from position s on into a new page and returns the
// page with the keys before s, or its smaller replacement, and the new page.
func (p *pageHead) split(s int) (*pageHead, *pageHead) {
	n := int(p.count)
	q := newPage(classFor(n - s))
	q.count, q.klen = uint8(n-s), p.klen
	copy(q.heads(), p.heads()[s:n])
	copy(q.vals(), p.vals()[s:n])
	p.count = uint8(s)
	if s <= pageShrink[p.class] {
		p = p.resize(classFor(s))
	}
	return p, q
}

// resize copies the page into a new page of the given class.
func (p *pageHead) resize(class int) *pageHead {
	q := newPage(class)
	q.count, q.klen = p.count, p.klen
	n := int(p.count)
	copy(q.heads()[:n], p.heads()[:n])
	copy(q.vals()[:n], p.vals()[:n])
	return q
}

// classFor returns the smallest page class that holds n keys.
func classFor(n int) int {
	c := 0
	for pageCaps[c] < n {
		c++
	}
	return c
}

// item is one key during a rebuild of a subtree (see build): a key with its
// one raw value, which may go into a page, or a key that has a leaf.
type item struct {
	key  []byte
	val  uint64
	leaf *leafHead
	full []byte // an immutable copy of key a K page may keep, or nil
}

// key returns key i of a page of either type, rebuilt in buf.
func (p *pageHead) key(i int, buf *keyBuf) []byte {
	if p.kind == kPageK {
		return p.kKey(i, buf)
	}
	return wordKey(p.keys()[i], int(p.klen), (*[8]byte)(buf[:8]))
}

// pageItems returns the keys of page p, of either type, as items, in order.
func pageItems(p *pageHead) []item {
	n := int(p.count)
	var buf []byte
	out := make([]item, n)
	var kb keyBuf
	ends := make([]int, n)
	for i := range n {
		buf = append(buf, p.key(i, &kb)...)
		ends[i] = len(buf)
	}
	start := 0
	for i, e := range ends {
		out[i].key = buf[start:e:e]
		start = e
	}
	if p.kind == kPage {
		for i, v := range p.vals()[:n] {
			out[i].val = v
		}
		return out
	}
	vs, ts, ls := p.kVals(), p.kTails(), p.kLens()
	for i := range out {
		if ls[i] > 16 {
			out[i].full = unsafe.Slice((*byte)(ts[i]), len(out[i].key))
		}
		out[i].val = vs[i]
	}
	return out
}

// pageFor returns the page that holds items, or nil if they do not fit one
// or one of them has a leaf. Keys of one length of at most 8 bytes go into a
// U8-1 page, all others into a K page (see kPack).
func pageFor(items []item) *pageHead {
	l := len(items[0].key)
	u8 := true
	for _, it := range items {
		if it.leaf != nil {
			return nil
		}
		u8 = u8 && len(it.key) == l && l <= 8
	}
	if !u8 {
		return kPack(items)
	}
	if len(items) > pageCaps[len(pageCaps)-1] {
		return nil
	}
	p := newPage(classFor(len(items)))
	p.count, p.klen = uint8(len(items)), uint8(l)
	h, vs := p.heads(), p.vals()
	for i, it := range items {
		h[i], vs[i] = keyWord(it.key), it.val
	}
	return p
}

// b2i returns 1 for true and 0 for false.
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// build returns a subtree holding exactly items, which are sorted by key,
// distinct, and all start with the same depth bytes. It is how pages burst
// and how pages change type when nothing simpler fits: the subtree is rebuilt
// from its keys. Items that fit a page become one; otherwise a range node
// takes their common path and splits them into ranges (see ranges).
func (t *Tree) build(items []item, depth int) *header {
	if len(items) == 1 && items[0].leaf != nil {
		return leafHdr(items[0].leaf)
	}
	if p := pageFor(items); p != nil {
		return pageHdr(p)
	}
	first, last := items[0].key, items[len(items)-1].key
	plen := swar.Lcp(first[depth:], last[depth:])
	d := depth + plen
	var h header
	h.setPrefix(first[depth:d], plen)
	if len(first) == d {
		// A key that ends here sorts first.
		h.term = items[0].leaf
		if h.term == nil {
			h.term = t.mk(items[0])
		}
		items = items[1:]
	}
	rs := t.ranges(items, d, nil)
	rs[0].b = 0
	return makeR(h, rs)
}
