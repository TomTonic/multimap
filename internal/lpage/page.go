// Package lpage is the second candidate for the multi-key page of the redesign
// (docs/redesign, PLAN step 3): a page whose header lists the lengths of its
// entries instead of holding a directory of tags. It follows the sketch of the
// user in docs/redesign/whataleafneedstostore.md (MKSV) and is built to be
// compared with the page of internal/vpage, in memory per key, in lookup time
// and in the cost of changes. It is a prototype and not yet part of the tree.
//
// A page holds the stripped suffixes of keys, each with one value, in one object
// of 128, 256 or 512 bytes without pointers:
//
//	header | common prefix | remainder 1 .. remainder n | value 1 .. value n
//
// The header is a code byte (the class of the page, the size of its header and
// the width of its values), the length of the common prefix, then one length byte
// per remainder, in key order, 0 for "no more entries", and, if the values have
// different lengths, one length byte per value. Its size is the smallest multiple
// of 8 of 8 to MaxHeader that has room for the entries: with values of different
// lengths 3, 7, 11 or 15 entries, with values of one width (a scalar
// specialization) 6, 14, 22 or 30. The common prefix is the part all suffixes of
// the page share, stored once; a remainder is what follows it and is at least one
// byte long. A value is a string of up to 255 bytes (a longer one needs another
// form, which the prototype does not have).
//
// A lookup reads the header, compares the common prefix, then the remainders in
// order; the offset of each follows from the lengths in the header, so the lines
// of the remainders are known after the first round. The value of the entry found
// follows from the lengths of the values before it.
//
// The version of this package with values of 8 bytes, whose value follows its
// remainder, is in the history (commit 4576cf6).
package lpage

import (
	"bytes"
	"errors"
	"unsafe"
)

const maxField = 255 // the longest suffix and the longest value

// sizes are the object sizes of the classes.
var sizes = [...]int{128, 256, 512}

// widths are the value widths of the pages with values of one length, by
// their code; code 0 is a page with values of different lengths.
var widths = [...]int{0, 1, 2, 4, 8, 16, 32, 64}

// MaxHeader is the largest header in bytes (a multiple of 8, at most 32); it
// limits the entries of a page. Experiments may change it.
var MaxHeader = 24

// ShrinkFill is how full a page may be in the next smaller class for it to move
// there after a removal, in percent.
var ShrinkFill = 70

// ErrTooLong is returned for suffixes or values that do not fit a page.
var ErrTooLong = errors.New("lpage: suffix or value longer than 255 bytes, or too many entries")

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
	code uint8 // bits 0-1 the class, 2-3 the header's size in 8-byte steps minus one, 4-6 the width code
	cp   uint8 // length of the common prefix
}

func (p *Page) class() int { return int(p.code & 3) }

// hdr returns the size of the header in bytes.
func (p *Page) hdr() int { return 8 * (1 + int(p.code>>2&3)) }

// width returns the length of every value, or 0 if they differ.
func (p *Page) width() int { return widths[p.code>>4&7] }

// capacity returns how many entries the header has room for.
func (p *Page) capacity() int { return capacityOf(p.hdr(), p.width()) }

// capacityOf returns how many entries a header of h bytes has room for when the
// values have the width w (0: they differ).
func capacityOf(h, w int) int {
	if w != 0 {
		return h - 2
	}
	return (h - 2) / 2
}

// Size returns the size of the page's object in bytes.
func (p *Page) Size() int { return sizes[p.class()] }

// Class returns the index of the page's size class.
func (p *Page) Class() int { return p.class() }

// PrefixLen returns the length of the common prefix.
func (p *Page) PrefixLen() int { return int(p.cp) }

// HeaderLen returns the size of the header in bytes.
func (p *Page) HeaderLen() int { return p.hdr() }

// Width returns the length of every value of the page, or 0 if they differ.
func (p *Page) Width() int { return p.width() }

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

// headerFor returns the smallest header in bytes that has room for n entries
// with values of width w, or 0 if there is none.
func headerFor(n, w int) int {
	for h := 8; h <= MaxHeader; h += 8 {
		if n <= capacityOf(h, w) {
			return h
		}
	}
	return 0
}

// Len returns the number of entries.
func (p *Page) Len() int {
	m := p.mem()
	c := p.capacity()
	if i := bytes.IndexByte(m[2:2+c], 0); i >= 0 {
		return i
	}
	return c
}

// lens returns the length arrays of the page: the remainders and, for values of
// different lengths, the values (otherwise nil).
func (p *Page) lens() (r, v []byte) {
	m := p.mem()
	c := p.capacity()
	r = m[2 : 2+c]
	if p.width() == 0 {
		v = m[2+c : 2+2*c]
	}
	return r, v
}

// vlen returns the length of the value of entry i; v is the array of the lengths
// of the values, nil for a page of values of one width.
func (p *Page) vlen(v []byte, i int) int {
	if v == nil {
		return p.width()
	}
	return int(v[i])
}

// span returns the offsets where the remainders end and the used bytes end.
func (p *Page) span() (kend, used int) {
	r, v := p.lens()
	n := p.Len()
	kend = p.hdr() + int(p.cp)
	for _, x := range r[:n] {
		kend += int(x)
	}
	used = kend
	if v == nil {
		return kend, used + n*p.width()
	}
	for _, x := range v[:n] {
		used += int(x)
	}
	return kend, used
}

// voffset returns the offset of the value of entry i, given where the remainders
// end.
func (p *Page) voffset(v []byte, kend, i int) int {
	if v == nil {
		return kend + i*p.width()
	}
	for _, x := range v[:i] {
		kend += int(x)
	}
	return kend
}

// Used returns the bytes of the page that hold something.
func (p *Page) Used() int {
	_, used := p.span()
	return used
}

// plan chooses the prefix and the sizes for sorted suffixes keys with values of
// the lengths vlen: the code byte of the page, the prefix length and the bytes
// needed. ok is false if there are too many entries, a remainder or a value is
// too long, or the largest class is too small. Values of one length that is a
// width of a page (1, 2, 4, 8, 16, 32 or 64) go into a page of that width.
func plan(keys [][]byte, vlen []int) (code uint8, cp, need int, ok bool) {
	n := len(keys)
	wcode := 0
	for i := 1; i < len(widths); i++ {
		if widths[i] == vlen[0] {
			wcode = i
		}
	}
	for _, x := range vlen {
		if x != vlen[0] {
			wcode = 0
		}
	}
	h := headerFor(n, widths[wcode])
	if h == 0 {
		return 0, 0, 0, false
	}
	if n > 1 {
		shortest := len(keys[0])
		for _, k := range keys {
			shortest = min(shortest, len(k))
		}
		cp = min(lcp(keys[0], keys[n-1]), shortest-1, maxField)
	}
	need = h + cp
	for i, k := range keys {
		if r := len(k) - cp; r < 1 || r > maxField || vlen[i] > maxField {
			return 0, 0, 0, false
		}
		need += len(k) - cp + vlen[i]
	}
	for c, s := range sizes {
		if need <= s {
			return uint8(c) | uint8(h/8-1)<<2 | uint8(wcode)<<4, cp, need, true
		}
	}
	return 0, 0, 0, false
}

// tooBig reports whether suffix s and value val do not fit a page of the
// largest class together, with the smallest header.
func tooBig(s, val []byte) bool {
	return len(s) > maxField || len(val) > maxField || len(s)+len(val) > sizes[len(sizes)-1]-8
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

func lengths(vals [][]byte) []int {
	out := make([]int, len(vals))
	for i, v := range vals {
		out[i] = len(v)
	}
	return out
}

// Build returns a page that holds the sorted, distinct suffixes keys with
// their values, or an error if they do not fit one. keys must not be empty.
func Build(keys, vals [][]byte) (*Page, error) {
	for _, k := range keys {
		if len(k) == 0 {
			return nil, ErrEmpty
		}
	}
	code, cp, _, ok := plan(keys, lengths(vals))
	if !ok {
		return nil, ErrTooLong
	}
	return write(code, cp, keys, vals), nil
}

// write lays out the page chosen by plan.
func write(code uint8, cp int, keys, vals [][]byte) *Page {
	p := alloc(int(code & 3))
	p.code, p.cp = code, uint8(cp)
	m := p.mem()
	r, v := p.lens()
	off := p.hdr()
	off += copy(m[off:], keys[0][:cp])
	for i, k := range keys {
		r[i] = byte(len(k) - cp)
		off += copy(m[off:], k[cp:])
	}
	for i, val := range vals {
		if v != nil {
			v[i] = byte(len(val))
		}
		off += copy(m[off:], val)
	}
	return p
}

// Get returns the value of suffix s. The slice aliases the page and is valid
// until the page changes.
func (p *Page) Get(s []byte) ([]byte, bool) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	if len(s) <= cp || !bytes.Equal(m[h:h+cp], s[:cp]) {
		return nil, false
	}
	rem := s[cp:]
	r, v := p.lens()
	off := h + cp
	first := rem[0]
	for i, x := range r {
		if x == 0 {
			break
		}
		if int(x) == len(rem) && m[off] == first && bytes.Equal(m[off:off+int(x)], rem) {
			kend := h + cp
			for _, y := range r[:p.Len()] {
				kend += int(y)
			}
			voff := p.voffset(v, kend, i)
			return m[voff : voff+p.vlen(v, i)], true
		}
		off += int(x)
	}
	return nil, false
}

// locate returns the position of s among the entries, whether it is there, and
// the offset of the remainder at that position (the end of the remainders if s
// is greater than all). s must start with the prefix.
func (p *Page) locate(s []byte) (i, off int, found bool) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	rem := s[cp:]
	r, _ := p.lens()
	off = h + cp
	for i = 0; i < len(r); i++ {
		x := int(r[i])
		if x == 0 {
			break
		}
		switch c := bytes.Compare(m[off:off+x], rem); {
		case c == 0:
			return i, off, true
		case c > 0:
			return i, off, false
		}
		off += x
	}
	return i, off, false
}

// Entries returns the suffixes and values of the page in order, as copies.
func (p *Page) Entries() (keys, vals [][]byte) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	r, v := p.lens()
	n := p.Len()
	kend, _ := p.span()
	off, voff := h+cp, kend
	for i := range n {
		x := int(r[i])
		k := make([]byte, 0, cp+x)
		k = append(append(k, m[h:h+cp]...), m[off:off+x]...)
		keys = append(keys, k)
		vl := p.vlen(v, i)
		vals = append(vals, bytes.Clone(m[voff:voff+vl]))
		off, voff = off+x, voff+vl
	}
	return keys, vals
}

// Insert sets the value of suffix s and returns the page that holds the result,
// which is p itself unless the page had to move to another class or lay out its
// entries anew. Full means the page is as big as it gets and has no room: split
// it (the result is p). A suffix or value longer than 255 bytes is an error. A
// page of values of one width takes a value of another length by laying its
// entries out anew.
func (p *Page) Insert(s, val []byte) (*Page, Result, error) {
	switch {
	case len(s) == 0:
		return p, Inserted, ErrEmpty
	case tooBig(s, val):
		return p, Inserted, ErrTooLong
	}
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	if w := p.width(); (w == 0 || len(val) == w) && len(s) > cp && bytes.Equal(m[h:h+cp], s[:cp]) {
		i, koff, found := p.locate(s)
		n := p.Len()
		r, v := p.lens()
		kend, used := p.span()
		voff := p.voffset(v, kend, i)
		if found {
			if p.vlen(v, i) == len(val) {
				copy(m[voff:], val)
				return p, Updated, nil
			}
			return p.insertSlow(s, val) // another length moves the values behind it
		}
		x := len(s) - cp
		if n < len(r) && used+x+len(val) <= len(m) {
			copy(m[voff+len(val):], m[voff:used])
			copy(m[voff:], val)
			copy(m[koff+x:], m[koff:used+len(val)])
			copy(m[koff:], s[cp:])
			copy(r[i+1:], r[i:n])
			r[i] = byte(x)
			if v != nil {
				copy(v[i+1:], v[i:n])
				v[i] = byte(len(val))
			}
			return p, Inserted, nil
		}
	}
	return p.insertSlow(s, val)
}

// insertSlow lays the entries out anew with s among them: the prefix may get
// shorter, the header longer, the page bigger, or a value change its length.
func (p *Page) insertSlow(s, val []byte) (*Page, Result, error) {
	keys, vals := p.Entries()
	i := 0
	for i < len(keys) && bytes.Compare(keys[i], s) < 0 {
		i++
	}
	res := Inserted
	if i < len(keys) && bytes.Equal(keys[i], s) {
		vals[i], res = val, Updated
	} else {
		keys = append(keys[:i], append([][]byte{s}, keys[i:]...)...)
		vals = append(vals[:i], append([][]byte{val}, vals[i:]...)...)
	}
	code, cp, _, ok := plan(keys, lengths(vals))
	if !ok {
		return p, Full, nil
	}
	return write(code, cp, keys, vals), res, nil
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
	i, koff, found := p.locate(s)
	if !found {
		return p, false
	}
	n := p.Len()
	if n == 1 {
		return nil, true
	}
	r, v := p.lens()
	kend, used := p.span()
	voff := p.voffset(v, kend, i)
	x, vl := int(r[i]), p.vlen(v, i)
	copy(m[voff:], m[voff+vl:used])
	copy(m[koff:], m[koff+x:used-vl])
	clear(m[used-x-vl : used])
	copy(r[i:], r[i+1:n])
	r[n-1] = 0
	if v != nil {
		copy(v[i:], v[i+1:n])
		v[n-1] = 0
	}
	if c := p.class(); c > 0 && p.Used()*100 <= ShrinkFill*sizes[c-1] {
		keys, vals := p.Entries()
		if code, cp, need, ok := plan(keys, lengths(vals)); ok && need*100 <= ShrinkFill*sizes[c-1] {
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
