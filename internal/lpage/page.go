// Package lpage is the second candidate for the multi-key page of the redesign
// (docs/redesign, PLAN step 3): a page whose header lists the lengths of its
// entries instead of holding a directory of tags. It follows the sketch of the
// user in docs/redesign/whataleafneedstostore.md (MKSV) and is built to be
// compared with the page of internal/vpage, in memory per key, in lookup time
// and in the cost of changes. It is a prototype and not yet part of the tree.
//
// A page holds the stripped suffixes of keys, each with one value of 8 bytes, in
// one object of 128, 256 or 512 bytes without pointers:
//
//	header | common prefix | remainder 1, value 1 | remainder 2, value 2 | ...
//
// The header is a code byte (the class of the page and the size of its header), the
// length of the common prefix, then one length byte per remainder, in key
// order, 0 for "no more entries"; its size is the smallest multiple of 8 of 8 to
// MaxHeader that has room for the entries (6, 14, 22 or 30). The common
// prefix is the part all suffixes of the page share, stored once; a remainder is
// what follows it and is at least one byte long. Each remainder is followed by its
// value, so that the line a lookup reads to compare a key holds its value as well
// (the sketch has all values after all remainders).
//
// A lookup reads the header, compares the common prefix, then the remainders in
// order; the offset of each follows from the lengths in the header, so the lines
// of the remainders are known after the first round.
package lpage

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unsafe"
)

const (
	valBytes  = 8
	maxSuffix = 255 // the longest suffix: a remainder takes at most 255 bytes after the prefix, which takes one byte less
)

// sizes are the object sizes of the classes.
var sizes = [...]int{128, 256, 512}

// MaxHeader is the largest header in bytes (a multiple of 8, at most 32); it
// limits the entries of a page to MaxHeader-2. Experiments may change it.
var MaxHeader = 24

// ShrinkFill is how full a page may be in the next smaller class for it to move
// there after a removal, in percent.
var ShrinkFill = 70

// ErrTooLong is returned for suffixes that do not fit a page.
var ErrTooLong = errors.New("lpage: suffix longer than 255 bytes, or too many")

// ErrEmpty is returned for an empty suffix: a length of 0 in the header ends the
// entries, so every remainder has at least one byte. A key that ends where the
// page starts is not an entry of a page.
var ErrEmpty = errors.New("lpage: empty suffix")

// Result says what Insert did.
type Result int

const (
	Inserted Result = iota
	Updated
	Full // the page is of the largest class or has the most entries: split it
)

// Page is the first two bytes of a page; the page is the object they start.
type Page struct {
	code uint8 // class in the low two bits, the header's size in 8-byte steps minus one above
	cp   uint8 // length of the common prefix
}

func (p *Page) class() int { return int(p.code & 3) }

// hdr returns the size of the header in bytes.
func (p *Page) hdr() int { return 8 * (1 + int(p.code>>2)) }

// Size returns the size of the page's object in bytes.
func (p *Page) Size() int { return sizes[p.class()] }

// Class returns the index of the page's size class.
func (p *Page) Class() int { return p.class() }

// PrefixLen returns the length of the common prefix.
func (p *Page) PrefixLen() int { return int(p.cp) }

// HeaderLen returns the size of the header in bytes.
func (p *Page) HeaderLen() int { return p.hdr() }

func (p *Page) mem() []byte { return unsafe.Slice((*byte)(unsafe.Pointer(p)), sizes[p.class()]) }

// alloc returns a zeroed object of class c.
func alloc(c int) *Page {
	switch c {
	case 0:
		return (*Page)(unsafe.Pointer(new([16]uint64)))
	case 1:
		return (*Page)(unsafe.Pointer(new([32]uint64)))
	}
	return (*Page)(unsafe.Pointer(new([64]uint64)))
}

// headerFor returns the smallest header in bytes that has room for n entries,
// or 0 if there is none.
func headerFor(n int) int {
	for h := 8; h <= MaxHeader; h += 8 {
		if n <= h-2 {
			return h
		}
	}
	return 0
}

// Len returns the number of entries.
func (p *Page) Len() int {
	m := p.mem()
	if i := bytes.IndexByte(m[2:p.hdr()], 0); i >= 0 {
		return i
	}
	return p.hdr() - 2
}

// used returns the bytes of the page that hold something.
func (p *Page) used() int {
	m := p.mem()
	n := p.Len()
	total := p.hdr() + int(p.cp) + n*valBytes
	for _, r := range m[2 : 2+n] {
		total += int(r)
	}
	return total
}

// Used returns the bytes of the page that hold something.
func (p *Page) Used() int { return p.used() }

// plan chooses the prefix and the sizes for sorted suffixes keys: the code byte
// of the page, the prefix length and the bytes needed. ok is false if there are
// too many entries, a remainder is too long, or the largest class is too small.
func plan(keys [][]byte) (code uint8, cp, need int, ok bool) {
	n := len(keys)
	h := headerFor(n)
	if h == 0 {
		return 0, 0, 0, false
	}
	if n > 1 {
		shortest := len(keys[0])
		for _, k := range keys {
			shortest = min(shortest, len(k))
		}
		cp = min(lcp(keys[0], keys[n-1]), shortest-1, 255)
	}
	rem := 0
	for _, k := range keys {
		if r := len(k) - cp; r < 1 || r > 255 {
			return 0, 0, 0, false
		}
		rem += len(k) - cp
	}
	need = h + cp + rem + n*valBytes
	for c, s := range sizes {
		if need <= s {
			return uint8(c) | uint8(h/8-1)<<2, cp, need, true
		}
	}
	return 0, 0, 0, false
}

func lcp(a, b []byte) int {
	n := min(len(a), len(b))
	for i := range n {
		if a[i] != b[i] {
			return i
		}
	}
	return n
}

// Build returns a page that holds the sorted, distinct suffixes keys with
// their values, or an error if they do not fit one. keys must not be empty.
func Build(keys [][]byte, vals []uint64) (*Page, error) {
	for _, k := range keys {
		if len(k) == 0 {
			return nil, ErrEmpty
		}
	}
	code, cp, _, ok := plan(keys)
	if !ok {
		return nil, ErrTooLong
	}
	return write(code, cp, keys, vals), nil
}

// write lays out the page chosen by plan.
func write(code uint8, cp int, keys [][]byte, vals []uint64) *Page {
	p := alloc(int(code & 3))
	p.code, p.cp = code, uint8(cp)
	m := p.mem()
	h := p.hdr()
	off := h
	off += copy(m[off:], keys[0][:cp])
	for i, k := range keys {
		m[2+i] = byte(len(k) - cp)
		off += copy(m[off:], k[cp:])
		binary.LittleEndian.PutUint64(m[off:], vals[i])
		off += valBytes
	}
	return p
}

// Get returns the value of suffix s.
func (p *Page) Get(s []byte) (uint64, bool) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	if len(s) <= cp || !bytes.Equal(m[h:h+cp], s[:cp]) {
		return 0, false
	}
	rem := s[cp:]
	off := h + cp
	lens := m[2:h]
	first := rem[0]
	for _, r := range lens {
		if r == 0 {
			break
		}
		if int(r) == len(rem) && m[off] == first && bytes.Equal(m[off:off+int(r)], rem) {
			return binary.LittleEndian.Uint64(m[off+int(r):]), true
		}
		off += int(r) + valBytes
	}
	return 0, false
}

// locate returns the position of s among the entries, whether it is there, and
// the offset of the entry at that position (the end of the entries if s is
// greater than all). s must start with the prefix.
func (p *Page) locate(s []byte) (i, off int, found bool) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	rem := s[cp:]
	off = h + cp
	for i = 0; i < h-2; i++ {
		r := int(m[2+i])
		if r == 0 {
			break
		}
		switch c := bytes.Compare(m[off:off+r], rem); {
		case c == 0:
			return i, off, true
		case c > 0:
			return i, off, false
		}
		off += r + valBytes
	}
	return i, off, false
}

// Entries returns the suffixes and values of the page in order.
func (p *Page) Entries() (keys [][]byte, vals []uint64) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	off := h + cp
	for i := 0; i < p.Len(); i++ {
		r := int(m[2+i])
		k := make([]byte, 0, cp+r)
		k = append(append(k, m[h:h+cp]...), m[off:off+r]...)
		keys = append(keys, k)
		vals = append(vals, binary.LittleEndian.Uint64(m[off+r:]))
		off += r + valBytes
	}
	return keys, vals
}

// Insert sets the value of suffix s and returns the page that holds the result,
// which is p itself unless the page had to move to another class or lay out its
// entries anew. Full means the page is as big as it gets and has no room: split
// it (the result is p). A suffix longer than 255 bytes is an error.
func (p *Page) Insert(s []byte, v uint64) (*Page, Result, error) {
	switch {
	case len(s) == 0:
		return p, Inserted, ErrEmpty
	case len(s) > maxSuffix:
		return p, Inserted, ErrTooLong
	}
	n := p.Len()
	h, cp := p.hdr(), int(p.cp)
	m := p.mem()
	if len(s) > cp && bytes.Equal(m[h:h+cp], s[:cp]) {
		i, off, found := p.locate(s)
		if found {
			binary.LittleEndian.PutUint64(m[off+int(m[2+i]):], v)
			return p, Updated, nil
		}
		r := len(s) - cp
		if used := p.used(); n < h-2 && used+r+valBytes <= len(m) {
			copy(m[off+r+valBytes:], m[off:used])
			copy(m[off:], s[cp:])
			binary.LittleEndian.PutUint64(m[off+r:], v)
			copy(m[2+i+1:], m[2+i:2+n])
			m[2+i] = byte(r)
			return p, Inserted, nil
		}
	}
	return p.insertSlow(s, v)
}

// insertSlow lays the entries out anew with s among them: the prefix may get
// shorter, the header longer or the page bigger.
func (p *Page) insertSlow(s []byte, v uint64) (*Page, Result, error) {
	keys, vals := p.Entries()
	i := 0
	for i < len(keys) && bytes.Compare(keys[i], s) < 0 {
		i++
	}
	keys = append(keys[:i], append([][]byte{s}, keys[i:]...)...)
	vals = append(vals[:i], append([]uint64{v}, vals[i:]...)...)
	code, cp, _, ok := plan(keys)
	if !ok {
		return p, Full, nil
	}
	return write(code, cp, keys, vals), Inserted, nil
}

// Delete removes suffix s and returns the page that holds the rest, nil if it
// is empty, and whether s was there. A page that has become thin moves to a
// smaller class.
func (p *Page) Delete(s []byte) (*Page, bool) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	if len(s) <= cp || !bytes.Equal(m[h:h+cp], s[:cp]) {
		return p, false
	}
	i, off, found := p.locate(s)
	if !found {
		return p, false
	}
	n := p.Len()
	if n == 1 {
		return nil, true
	}
	r := int(m[2+i])
	used := p.used()
	copy(m[off:], m[off+r+valBytes:used])
	clear(m[used-r-valBytes : used])
	copy(m[2+i:], m[2+i+1:2+n])
	m[2+n-1] = 0
	if c := p.class(); c > 0 && p.used()*100 <= ShrinkFill*sizes[c-1] {
		keys, vals := p.Entries()
		if code, cp, need, ok := plan(keys); ok && need*100 <= ShrinkFill*sizes[c-1] {
			return write(code, cp, keys, vals), true
		}
	}
	return p, true
}

// Split divides a page of at least two entries in the middle by count and
// returns the two halves, each in the smallest class that holds it.
func (p *Page) Split() (left, right *Page) {
	keys, vals := p.Entries()
	m := len(keys) / 2
	left, _ = Build(keys[:m], vals[:m])
	right, _ = Build(keys[m:], vals[m:])
	return left, right
}
