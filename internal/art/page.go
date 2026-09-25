package art

import (
	"encoding/binary"
	"slices"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// A page holds all keys of one subtree in a single object of 64, 128, 256 or
// 512 bytes: Go size classes that are aligned to their size and carry no
// malloc header. It replaces a leaf per key, and the pointer to each leaf, by
// sorted arrays, so that a lookup ends in one object and a range scan walks
// contiguous memory.
//
// Pages hold values of a small pointer-free type only (see Tree.small). There
// are three page types:
//
//   - U8-1 (this file): keys of one common length of at most 8 bytes (integer
//     keys are 8), each with exactly one value. The keys are stored as
//     big-endian words; keys and values are plain words, so the page contains
//     no pointers and the garbage collector never scans it.
//   - U8-n (pagen.go): keys like U8-1, with any number of values each.
//   - S (pages.go): keys of any lengths up to maxPageKey, with any number of
//     values each.
//
// Every page can rebuild its full keys, so it never depends on where in the
// tree it sits: when a node above it collapses, the page just moves up.

// pageCaps are the capacities of the four page classes.
var pageCaps = [4]int{3, 7, 15, 31}

// pageHead is the start of every page (8 B), U8-1 and U8-n alike.
type pageHead struct {
	kind  kind
	class uint8 // index into pageCaps (U8-1), nLayouts (U8-n) or sSizes (S)
	count uint8 // keys
	klen  uint8 // U8: length of every key in the page, 0..8
	nv    uint8 // U8-n and S: values held inline
	kcap  uint8 // S: key slots
	vcap  uint8 // S: inline value slots
	base  uint8 // S: length of the prefix all keys share
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

// keys returns the keys of a U8 page, in ascending order. Pages of both U8
// types keep them at the same place.
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
// raw values (at most inlineMax) or with an external value set, either of
// which may go into a page, or a key that must stay a leaf.
type item struct {
	key  []byte
	vals []uint64
	set  unsafe.Pointer
	leaf *leafHead
}

// key returns key i of a page of any type, rebuilt in buf.
func (p *pageHead) key(i int, buf *keyBuf) []byte {
	if p.kind == kPageS {
		return p.sKey(i, buf)
	}
	return wordKey(p.keys()[i], int(p.klen), (*[8]byte)(buf[:8]))
}

// pageItems returns the keys of page p, of any type, as items, in order.
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
		vs := p.vals()
		for i := range out {
			out[i].vals = vs[i : i+1 : i+1]
		}
		return out
	}
	vs, ex := slices.Clone(p.nvals()[:p.nv]), p.exts()
	for i := range out {
		off, cnt, e := p.run(i)
		if e >= 0 {
			out[i].set = ex[e]
		} else {
			out[i].vals = vs[off : off+cnt : off+cnt]
		}
	}
	return out
}

// pageFor returns the page that holds items, or nil if they do not fit one.
// Keys of one length of at most 8 bytes go into a U8 page, all others into an
// S page (see sPack). Only trees with pages rebuild subtrees, so the values
// may go into pages.
func pageFor(items []item) *pageHead {
	l := len(items[0].key)
	for _, it := range items {
		if len(it.key) != l || l > 8 {
			return sPack(items)
		}
	}
	return u8Pack(items)
}

// u8Pack returns the U8 page that holds items, keys of one length of at most
// 8 bytes, or nil if they do not fit one. Keys with one value each go into a
// U8-1 page if there are at most 31; otherwise a U8-n page takes them if its
// largest class holds them.
func u8Pack(items []item) *pageHead {
	l := len(items[0].key)
	single, inline, sets := true, 0, 0
	for _, it := range items {
		// No item is a leaf: a key that must be one is longer than 8 bytes.
		switch {
		case it.set != nil:
			sets++
			single = false
		default:
			inline += len(it.vals)
			single = single && len(it.vals) == 1
		}
	}
	if single && len(items) <= pageCaps[len(pageCaps)-1] {
		p := newPage(classFor(len(items)))
		p.count, p.klen = uint8(len(items)), uint8(l)
		h, vs := p.heads(), p.vals()
		for i, it := range items {
			h[i], vs[i] = keyWord(it.key), it.vals[0]
		}
		return p
	}
	c := nClassFor(len(items), inline, sets)
	if c < 0 {
		return nil
	}
	p := newPageN(c)
	p.count, p.klen, p.nv = uint8(len(items)), uint8(l), uint8(inline)
	h := p.nHeads()
	for i, it := range items {
		h[i] = keyWord(it.key)
	}
	p.fillVals(items)
	return p
}

// build returns a subtree holding exactly items, which are sorted by key,
// distinct, and all start with the same depth bytes. It is how pages burst
// and how pages change type when nothing simpler fits: the subtree is rebuilt
// from its keys. Groups of keys that fit a page become pages; the rest becomes
// nodes and leaves.
func (t *Tree) build(items []item, depth int) *header {
	if len(items) == 1 && items[0].leaf != nil {
		return leafHdr(items[0].leaf)
	}
	if p := pageFor(items); p != nil {
		return pageHdr(p)
	}
	first, last := items[0].key, items[len(items)-1].key
	plen := swar.Lcp(first[depth:], last[depth:])
	n := &node4{}
	n.kind = kN4
	n.setPrefix(first[depth:depth+plen], plen)
	d := depth + plen
	h := &n.header
	if len(items[0].key) == d {
		// A key that ends here sorts first. It is never a leaf already: a key
		// that must be a leaf is longer than any key a page holds, and all
		// other items come from one page.
		h.term = t.mk(items[0])
		items = items[1:]
	}
	for len(items) > 0 {
		b, j := items[0].key[d], 1
		for j < len(items) && items[j].key[d] == b {
			j++
		}
		h = addChild(h, b, t.build(items[:j], d+1))
		items = items[j:]
	}
	return h
}
