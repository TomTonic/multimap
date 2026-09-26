package art

import (
	"bytes"
	"cmp"
	"encoding/binary"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// K pages hold keys of any length up to maxPageKey,
// each with exactly one value, in a fixed layout. The prefix all keys share is
// held once; of each key's rest, the suffix, the first 16 bytes are two words
// in heads (big endian, zero-padded), which order and find the keys. A key
// whose suffix is longer than 16 bytes keeps a copy of its full key, which
// only a hit on it, a burst and the bounds of a range read.
//
// Unlike S pages, nothing is packed: every array has a fixed place per class,
// so no offsets or counts are summed up, and inserting or removing a key only
// shifts the keys after it.
//
// Layout, after the pageKHead:
//
//	tails [k]unsafe.Pointer  full key of each key with a suffix over 16 bytes, else nil
//	heads [k][2]uint64       the first 16 suffix bytes of each key
//	vals  [k]uint64          the value of each key
//	lens  [k]uint8           the length of each key's suffix
//
// The pointers come first, so that the garbage collector reads only them.

// pageKHead starts every K page (24 B). klen is the length of the key at pre,
// kcap the capacity of the class, base the length of the shared prefix.
type pageKHead struct {
	pageHead
	pre unsafe.Pointer // a key of at least base bytes: the shared prefix
	win [8]byte        // the last up to 8 bytes of the shared prefix
}

type (
	pageK1 struct {
		pageKHead
		tails [1]unsafe.Pointer
		heads [2]uint64
		vals  [1]uint64
		lens  [1]uint8
	}
	pageK3 struct {
		pageKHead
		tails [3]unsafe.Pointer
		heads [6]uint64
		vals  [3]uint64
		lens  [3]uint8
	}
	pageK7 struct {
		pageKHead
		tails [7]unsafe.Pointer
		heads [14]uint64
		vals  [7]uint64
		lens  [7]uint8
	}
	pageK14 struct {
		pageKHead
		tails [14]unsafe.Pointer
		heads [28]uint64
		vals  [14]uint64
		lens  [14]uint8
	}
)

// kCaps are the capacities of the four K classes of 64, 128, 256 and 512
// bytes; a page shrinks into the next smaller class at kShrink keys.
var (
	kCaps   = [4]int{1, 3, 7, 14}
	kShrink = [4]int{0, 0, 2, 5}
)

const kHeadOff = unsafe.Sizeof(pageKHead{})

func newPageK(class int) *pageHead {
	var p *pageHead
	switch class {
	case 0:
		p = &(&pageK1{}).pageHead
	case 1:
		p = &(&pageK3{}).pageHead
	case 2:
		p = &(&pageK7{}).pageHead
	default:
		p = &(&pageK14{}).pageHead
	}
	p.kind, p.class, p.kcap = kPageK, uint8(class), uint8(kCaps[class])
	return p
}

func (p *pageHead) kh() *pageKHead { return (*pageKHead)(unsafe.Pointer(p)) }

func (p *pageHead) kTails() []unsafe.Pointer {
	return unsafe.Slice((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(p), kHeadOff)), p.kcap)
}

func (p *pageHead) kHeads() []uint64 {
	k := uintptr(p.kcap)
	return unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), kHeadOff+8*k)), 2*k)
}

func (p *pageHead) kVals() []uint64 {
	k := uintptr(p.kcap)
	return unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), kHeadOff+24*k)), k)
}

func (p *pageHead) kLens() []uint8 {
	k := uintptr(p.kcap)
	return unsafe.Slice((*uint8)(unsafe.Add(unsafe.Pointer(p), kHeadOff+32*k)), k)
}

// kWords returns the two head words of suffix s.
func kWords(s []byte) (uint64, uint64) {
	if len(s) >= 16 {
		return binary.BigEndian.Uint64(s), binary.BigEndian.Uint64(s[8:])
	}
	var b [16]byte
	copy(b[:], s)
	return binary.BigEndian.Uint64(b[:8]), binary.BigEndian.Uint64(b[8:])
}

// kPrefix returns the shared prefix of the page.
func (p *pageHead) kPrefix() []byte {
	h := p.kh()
	return unsafe.Slice((*byte)(h.pre), h.klen)[:p.base]
}

// kMatch reports whether key, whose first from bytes are known to match,
// starts with the page's shared prefix. The last 8 bytes of the prefix are in
// the page; only a check that reaches further back reads the prefix key.
func (p *pageHead) kMatch(key []byte, from int) bool {
	b := int(p.base)
	if from >= b {
		return true
	}
	if ws := max(b-8, 0); from >= ws {
		return bytes.Equal(key[from:b], p.kh().win[from-ws:b-ws])
	}
	return bytes.Equal(key[from:b], p.kPrefix()[from:])
}

// kFind looks for key in the page like sFind.
func (p *pageHead) kFind(key []byte, from int) (i int, found, match bool) {
	b := int(p.base)
	if len(key) < b || len(key) > maxPageKey || !p.kMatch(key, from) {
		return 0, false, false
	}
	s := key[b:]
	w0, w1 := kWords(s)
	n, h := int(p.count), p.kHeads()
	if n > 7 && (h[12] < w0 || h[12] == w0 && h[13] < w1) {
		i = 7
	}
	for i < n && (h[2*i] < w0 || h[2*i] == w0 && h[2*i+1] < w1) {
		i++
	}
	for ; i < n && h[2*i] == w0 && h[2*i+1] == w1; i++ {
		if c := p.kCmp(i, s); c >= 0 {
			return i, c == 0, true
		}
	}
	return i, false, true
}

// kCmp compares the suffix of key i with s, whose head words are equal.
func (p *pageHead) kCmp(i int, s []byte) int {
	var kt, st []byte
	l := int(p.kLens()[i])
	if l > 16 {
		b := int(p.base)
		kt = unsafe.Slice((*byte)(p.kTails()[i]), b+l)[b+16:]
	}
	if len(s) > 16 {
		st = s[16:]
	}
	if c := bytes.Compare(kt, st); c != 0 {
		return c
	}
	return cmp.Compare(l, len(s))
}

// kKey returns key i, from its full key or rebuilt in buf.
func (p *pageHead) kKey(i int, buf *keyBuf) []byte {
	l, b := int(p.kLens()[i]), int(p.base)
	if l > 16 {
		return unsafe.Slice((*byte)(p.kTails()[i]), b+l)
	}
	n := copy(buf[:], p.kPrefix())
	h := p.kHeads()
	binary.BigEndian.PutUint64(buf[n:], h[2*i])
	binary.BigEndian.PutUint64(buf[n+8:], h[2*i+1])
	return buf[:n+l]
}

// kInsertAt inserts key with value v at position i. The page must have room
// and key must start with its shared prefix.
func (p *pageHead) kInsertAt(i int, key []byte, v uint64) {
	n, s := int(p.count), key[p.base:]
	t, h, vs, ls := p.kTails(), p.kHeads(), p.kVals(), p.kLens()
	copy(t[i+1:n+1], t[i:n])
	copy(h[2*i+2:2*n+2], h[2*i:2*n])
	copy(vs[i+1:n+1], vs[i:n])
	copy(ls[i+1:n+1], ls[i:n])
	t[i] = nil
	if len(s) > 16 {
		t[i] = unsafe.Pointer(unsafe.SliceData(bytes.Clone(key)))
	}
	h[2*i], h[2*i+1] = kWords(s)
	vs[i], ls[i] = v, uint8(len(s))
	p.count++
}

// kRemoveAt removes key i and returns the page, its smaller replacement, or
// nil once it is empty.
func (p *pageHead) kRemoveAt(i int) *pageHead {
	n := int(p.count)
	if n == 1 {
		return nil
	}
	t, h, vs, ls := p.kTails(), p.kHeads(), p.kVals(), p.kLens()
	copy(t[i:n-1], t[i+1:n])
	copy(h[2*i:2*n-2], h[2*i+2:2*n])
	copy(vs[i:n-1], vs[i+1:n])
	copy(ls[i:n-1], ls[i+1:n])
	t[n-1] = nil
	p.count--
	if c := int(p.class); c > 0 && n-1 <= kShrink[c] {
		return p.kResize(c - 1)
	}
	return p
}

// kResize copies the page into a new page of the given class.
func (p *pageHead) kResize(class int) *pageHead {
	q := newPageK(class)
	n := int(p.count)
	q.count, q.base, q.klen = p.count, p.base, p.klen
	q.kh().pre, q.kh().win = p.kh().pre, p.kh().win
	copy(q.kTails(), p.kTails()[:n])
	copy(q.kHeads(), p.kHeads()[:2*n])
	copy(q.kVals(), p.kVals()[:n])
	copy(q.kLens(), p.kLens()[:n])
	return q
}

// kSingle reports whether items can go into a K page as far as their values
// go: every key with exactly one value, none a leaf.
func kSingle(items []item) bool {
	for _, it := range items {
		if it.leaf != nil || it.set != nil || len(it.vals) != 1 {
			return false
		}
	}
	return true
}

// kPack returns the K page that holds items, which are sorted, distinct and
// single-valued (see kSingle), or nil if there are too many.
func kPack(items []item) *pageHead {
	k := len(items)
	if k > kCaps[len(kCaps)-1] {
		return nil
	}
	b := swar.Lcp(items[0].key, items[k-1].key)
	c := 0
	for kCaps[c] < k {
		c++
	}
	p := newPageK(c)
	p.count, p.base = uint8(k), uint8(b)
	t, h, vs, ls := p.kTails(), p.kHeads(), p.kVals(), p.kLens()
	var pre []byte
	for i, it := range items {
		s := it.key[b:]
		if len(s) > 16 {
			full := it.full
			if full == nil {
				full = bytes.Clone(it.key)
			}
			t[i] = unsafe.Pointer(unsafe.SliceData(full))
			if pre == nil {
				pre = full
			}
		}
		h[2*i], h[2*i+1] = kWords(s)
		vs[i], ls[i] = it.vals[0], uint8(len(s))
	}
	if b > 0 {
		if pre == nil {
			pre = bytes.Clone(items[0].key[:b])
		}
		p.kh().pre, p.klen = unsafe.Pointer(unsafe.SliceData(pre)), uint8(len(pre))
	}
	copy(p.kh().win[:], items[0].key[max(b-8, 0):b])
	return p
}
