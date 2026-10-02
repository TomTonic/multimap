package art

import (
	"encoding/binary"
	"math/bits"
	"unsafe"
)

// A page holds all keys of one subtree in a single object of 64, 128, 256 or
// 512 bytes: Go size classes that are aligned to their size and carry no
// malloc header. It replaces a leaf per key, and the pointer to each leaf, by
// sorted arrays, so that a lookup ends in one object and a range scan walks
// contiguous memory.
//
// Pages hold values of a small pointer-free type only (see Tree.small), and
// only keys with exactly one value keep it in the page. All keys of a page have
// one length of at most 8 bytes (integer keys are 8), the length of the first
// key of the tree (see Tree.pk). The keys are stored as big-endian words; keys
// and values are plain words, so the page contains no pointers and the garbage
// collector never scans it. A key of another length, or with more than one
// value, has a leaf, which a range node holds like a page (see rnode.go).
//
// Every page can rebuild its full keys, so it never depends on where in the
// tree it sits: when a node above it collapses, the page just moves up.

// pageCaps are the capacities of the four page classes.
var pageCaps = [4]int{3, 7, 15, 31}

// maxPageKey is the longest key a page holds; a longer key gets a leaf.
const maxPageKey = 8

// pageHead is the start of every page (16 B).
type pageHead struct {
	kind  kind
	class uint8 // index into pageCaps
	count uint8 // keys
	klen  uint8 // length of every key in the page, 0..8
	stale uint8 // keys removed since bloom was computed
	_     [3]byte
	bloom uint64 // one bit per key (see bloomBit); keys removed since leave theirs set
}

// bloomBit returns the bit of key word w in a page's bloom filter. A lookup that
// finds its bit unset knows that the key is not in the page, without reading
// anything but the head: most lookups of absent keys end there, where a
// lookup in a tree of leaves ends one node earlier than its keys' depth. The 64
// bits are about half set in a page of 31 keys.
func bloomBit(w uint64) uint64 { return 1 << ((w * 0x9E3779B97F4A7C15) >> 58) }

// rebloom recomputes the page's bloom filter from its keys.
func (p *pageHead) rebloom() {
	p.bloom, p.stale = 0, 0
	for _, w := range p.keys() {
		p.bloom |= bloomBit(w)
	}
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
	h := p.heads()
	for i := range h {
		h[i] = pad
	}
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

// pad fills the head slots of a page that hold no key, so that comparisons
// against them are never "below" and search can look at whole blocks.
const pad = ^uint64(0)

// search returns the position of key w in a page, or where it would be
// inserted, and whether it is there.
//
// It counts the keys below w without branching: the heads of a page are
// compared in blocks of 8, one fence head per block (the last of the block)
// picks the block, and the other 7 heads of that block are counted. All these
// loads are independent of each other, so the cache misses of a page that is
// not in the L1 cache overlap, and no branch depends on where w lies, which
// mispredicted about once per lookup in a linear search (measured against it
// and against binary searches: 15-25% faster on 16K to 64K integer keys).
func (p *pageHead) search(w uint64) (int, bool) {
	h := p.heads()
	var i int
	switch p.class {
	case 3:
		blk := below(h[7], w) + below(h[15], w) + below(h[23], w)
		i = 8*blk + count7((*[7]uint64)(h[8*blk:]), w)
	case 2:
		blk := below(h[7], w)
		i = 8*blk + count7((*[7]uint64)(h[8*blk:]), w)
	case 1:
		i = count7((*[7]uint64)(h), w)
	default:
		i = below(h[0], w) + below(h[1], w) + below(h[2], w)
	}
	return i, i < int(p.count) && h[i] == w
}

// below returns 1 if x < w, else 0, without a branch.
func below(x, w uint64) int {
	_, b := bits.Sub64(x, w, 0)
	return int(b)
}

// count7 returns how many of the 7 heads in h are below w.
func count7(h *[7]uint64, w uint64) int {
	return below(h[0], w) + below(h[1], w) + below(h[2], w) + below(h[3], w) + below(h[4], w) + below(h[5], w) + below(h[6], w)
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
	p.bloom |= bloomBit(w)
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
	h[n-1] = pad
	p.count--
	if p.stale++; p.stale >= 8 {
		p.rebloom()
	}
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
	for i := s; i < n; i++ {
		p.heads()[i] = pad
	}
	p.count = uint8(s)
	q.rebloom()
	if s <= pageShrink[p.class] {
		p = p.resize(classFor(s))
	}
	p.rebloom()
	return p, q
}

// resize copies the page into a new page of the given class.
func (p *pageHead) resize(class int) *pageHead {
	q := newPage(class)
	q.count, q.klen, q.bloom, q.stale = p.count, p.klen, p.bloom, p.stale
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

// key returns key i of a page, rebuilt in buf.
func (p *pageHead) key(i int, buf *[8]byte) []byte {
	return wordKey(p.keys()[i], int(p.klen), buf)
}

// pageItems returns the keys of page p as items, in order.
func pageItems(p *pageHead) []item {
	n := int(p.count)
	out := make([]item, n)
	buf := make([]byte, 0, n*int(p.klen))
	for i, w := range p.keys() {
		var kb [8]byte
		start := len(buf)
		buf = append(buf, wordKey(w, int(p.klen), &kb)...)
		out[i].key = buf[start:len(buf):len(buf)]
		out[i].val = p.vals()[i]
	}
	return out
}

// pageFor returns the page that holds items, or nil if they do not fit one:
// there are too many, they differ in length, or one of them has a leaf.
func pageFor(items []item) *pageHead {
	if len(items) > pageCaps[len(pageCaps)-1] {
		return nil
	}
	l := len(items[0].key)
	for _, it := range items {
		if it.leaf != nil || len(it.key) != l {
			return nil
		}
	}
	p := newPage(classFor(len(items)))
	p.count, p.klen = uint8(len(items)), uint8(l)
	h, vs := p.heads(), p.vals()
	for i, it := range items {
		h[i], vs[i] = keyWord(it.key), it.val
	}
	p.rebloom()
	return p
}

// b2i returns 1 for true and 0 for false.
func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
