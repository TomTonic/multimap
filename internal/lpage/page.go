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
	"encoding/binary"
	"errors"
	"math/bits"
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

// MinHeader is the smallest header in bytes a page gets, a multiple of 8: a
// larger one costs every page its bytes but leaves room for entries to come
// in place, instead of laying the page out anew each time its header is full.
var MinHeader = 8

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
	// Present is the result of TryInsert for a suffix that is already there:
	// nothing changed.
	Present
)

// KindBase is added to the code byte of every page, so that a tree whose objects
// tell their kind by their first byte can give the pages the kinds
// KindBase to KindBase+Kinds-1. A package that uses pages standalone leaves it 0.
var KindBase uint8

// Kinds is the number of different first bytes of a page.
const Kinds = 128

// Page is the first two bytes of a page; the page is the object they start.
type Page struct {
	code uint8 // KindBase plus: bits 0-1 the class, 2-3 the header's size in 8-byte steps minus one, 4-6 the width code
	cp   uint8 // length of the common prefix
}

func (p *Page) class() int { return int(p.code-KindBase) & 3 }

// hdr returns the size of the header in bytes.
func (p *Page) hdr() int { return 8 * (1 + int((p.code-KindBase)>>2&3)) }

// width returns the length of every value, or 0 if they differ.
func (p *Page) width() int { return widths[(p.code-KindBase)>>4&7] }

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
	for h := max(8, MinHeader); h <= MaxHeader; h += 8 {
		if n <= capacityOf(h, w) {
			return h
		}
	}
	return 0
}

// Len returns the number of entries.
func (p *Page) Len() int {
	r := p.mem()[2 : 2+p.capacity()]
	off := 0
	for ; off+8 <= len(r); off += 8 {
		x := binary.LittleEndian.Uint64(r[off:])
		if z := (x - 0x0101010101010101) &^ x & 0x8080808080808080; z != 0 {
			return off + bits.TrailingZeros64(z)/8
		}
	}
	for ; off < len(r); off++ {
		if r[off] == 0 {
			return off
		}
	}
	return len(r)
}

// sumBytes returns the sum of the bytes of b, a word at a time: the lengths in
// a header are a few bytes and add up to the offsets of the entries.
func sumBytes(b []byte) int {
	s := 0
	for ; len(b) >= 8; b = b[8:] {
		x := binary.LittleEndian.Uint64(b)
		x = x&0x00FF00FF00FF00FF + x>>8&0x00FF00FF00FF00FF
		s += int(x * 0x0001000100010001 >> 48)
	}
	for _, c := range b {
		s += int(c)
	}
	return s
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
	kend = p.hdr() + int(p.cp) + sumBytes(r) // the lengths behind the last entry are 0
	if v == nil {
		return kend, kend + p.Len()*p.width()
	}
	return kend, kend + sumBytes(v)
}

// voffset returns the offset of the value of entry i, given where the remainders
// end.
func (p *Page) voffset(v []byte, kend, i int) int {
	if v == nil {
		return kend + i*p.width()
	}
	return kend + sumBytes(v[:i])
}

// Used returns the bytes of the page that hold something.
func (p *Page) Used() int {
	_, used := p.span()
	return used
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

// Build returns a page that holds the sorted, distinct suffixes keys with
// their values, or an error if they do not fit one. keys must not be empty.
func Build(keys, vals [][]byte) (*Page, error) {
	if len(keys) > maxEnts {
		return nil, ErrTooLong
	}
	var es [maxEnts]ent
	for i, k := range keys {
		if len(k) == 0 {
			return nil, ErrEmpty
		}
		es[i] = ent{k1: k, v: vals[i]}
	}
	if p := buildEnts(es[:len(keys)]); p != nil {
		return p, nil
	}
	return nil, ErrTooLong
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
		if b := m[off]; b != rem[0] { // most entries differ at the first byte
			if b > rem[0] {
				return i, off, false
			}
			off += x
			continue
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
	q, res, _, err := p.insert(s, val, true, false)
	return q, res, err
}

// TryInsert adds suffix s with value val if it is not there, and returns the
// page that holds the result. If s is there it changes nothing and returns
// Present and the position of s; the position of a new entry is not reported. It
// serves a tree that keeps one value per key in its pages and moves a key with a
// second value elsewhere. Full and errors are as for Insert. With cow the page
// is never changed in place: the result is always another page, as in a tree
// whose pages are immutable (see StringAt).
func (p *Page) TryInsert(s, val []byte, cow bool) (*Page, Result, int, error) {
	return p.insert(s, val, false, cow)
}

func (p *Page) insert(s, val []byte, replace, cow bool) (*Page, Result, int, error) {
	switch {
	case len(s) == 0:
		return p, Inserted, 0, ErrEmpty
	case tooBig(s, val):
		return p, Inserted, 0, ErrTooLong
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
			if !replace {
				return p, Present, i, nil
			}
			if p.vlen(v, i) == len(val) {
				copy(m[voff:], val)
				return p, Updated, i, nil
			}
			q, res, err := p.insertSlow(s, val) // another length moves the values behind it
			return q, res, i, err
		}
		x := len(s) - cp
		if !cow && n < len(r) && used+x+len(val) <= len(m) {
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
			return p, Inserted, i, nil
		}
		q, res, err := p.insertCopy(i, koff, voff, used, s, val)
		return q, res, i, err
	}
	if !replace {
		if i, ok := p.Find(s); ok {
			return p, Present, i, nil
		}
	}
	q, res, err := p.insertSlow(s, val)
	return q, res, 0, err
}

// insertCopy lays out the page that holds the entries of p and the new entry
// (s, val) at position i, in the prefix p has: the entry's remainder goes to koff
// and its value to voff of p's bytes, which end at used. The header may get
// longer and the page bigger; it is Full if the entry fits no page. It is Insert's
// way to grow a page without laying all its entries out anew.
func (p *Page) insertCopy(i, koff, voff, used int, s, val []byte) (*Page, Result, error) {
	m := p.mem()
	h, cp := p.hdr(), int(p.cp)
	r, v := p.lens()
	n := p.Len()
	h2 := headerFor(n+1, p.width())
	if h2 == 0 {
		return p, Full, nil
	}
	x := len(s) - cp
	need := used - h + h2 + x + len(val)
	c := 0
	for c < len(sizes) && need > sizes[c] {
		c++
	}
	if c == len(sizes) {
		return p, Full, nil
	}
	q := alloc(c)
	q.code = uint8(c) | uint8(h2/8-1)<<2 | (p.code-KindBase)&0x70 + KindBase
	q.cp = p.cp
	m2 := q.mem()
	r2, v2 := q.lens()
	copy(r2, r[:i])
	r2[i] = byte(x)
	copy(r2[i+1:], r[i:n])
	if v != nil {
		copy(v2, v[:i])
		v2[i] = byte(len(val))
		copy(v2[i+1:], v[i:n])
	}
	off := h2
	off += copy(m2[off:], m[h:koff])
	off += copy(m2[off:], s[cp:])
	off += copy(m2[off:], m[koff:voff])
	off += copy(m2[off:], val)
	copy(m2[off:], m[voff:used])
	return q, Inserted, nil
}

// insertSlow lays the entries out anew with s among them: the prefix may get
// shorter, the header longer, the page bigger, or a value change its length.
func (p *Page) insertSlow(s, val []byte) (*Page, Result, error) {
	var es [maxEnts]ent
	n := p.load(es[:])
	i, found := position(es[:n], s)
	res := Inserted
	if found {
		es[i].v, res = val, Updated
	} else {
		if n == maxEnts-1 {
			return p, Full, nil
		}
		copy(es[i+1:n+1], es[i:n])
		es[i] = ent{k1: s, v: val}
		n++
	}
	if q := buildEnts(es[:n]); q != nil {
		return q, res, nil
	}
	return p, Full, nil
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
	q, _ := p.deleteEntry(i, koff, false)
	return q, true
}

// deleteAt removes entry i; see Delete. With cow the page is not changed in
// place.
func (p *Page) deleteAt(i int, cow bool) (*Page, bool) {
	koff, _ := p.offsets(i)
	return p.deleteEntry(i, koff, cow)
}

// deleteEntry removes entry i, whose remainder starts at offset koff.
func (p *Page) deleteEntry(i, koff int, cow bool) (*Page, bool) {
	n := p.Len()
	if n == 1 {
		return nil, true
	}
	r, v := p.lens()
	kend, used := p.span()
	voff := p.voffset(v, kend, i)
	x, vl := int(r[i]), p.vlen(v, i)
	if cow {
		p = p.without(i, koff, voff, used)
	} else {
		m := p.mem()
		copy(m[voff:], m[voff+vl:used])
		copy(m[koff:], m[koff+x:used-vl])
		clear(m[used-x-vl : used])
		copy(r[i:], r[i+1:n])
		r[n-1] = 0
		if v != nil {
			copy(v[i:], v[i+1:n])
			v[n-1] = 0
		}
	}
	if c := p.class(); c > 0 && p.Used()*100 <= ShrinkFill*sizes[c-1] {
		var es [maxEnts]ent
		if q := fitEnts(es[:p.load(es[:])], c, ShrinkFill); q != nil {
			return q, true
		}
	}
	return p, true
}

// without returns a page of the class and header of p that holds its entries but
// entry i, whose remainder starts at koff and whose value at voff of the used
// bytes of p.
func (p *Page) without(i, koff, voff, used int) *Page {
	m := p.mem()
	h := p.hdr()
	r, v := p.lens()
	n := p.Len()
	x, vl := int(r[i]), p.vlen(v, i)
	q := alloc(p.class())
	q.code, q.cp = p.code, p.cp
	m2 := q.mem()
	r2, v2 := q.lens()
	copy(r2, r[:i])
	copy(r2[i:], r[i+1:n])
	if v != nil {
		copy(v2, v[:i])
		copy(v2[i:], v[i+1:n])
	}
	off := h
	off += copy(m2[off:], m[h:koff])
	off += copy(m2[off:], m[koff+x:voff])
	copy(m2[off:], m[voff+vl:used])
	return q
}

// Split divides a page of at least two entries in the middle by count and
// returns the two halves, each in the smallest class that holds it.
func (p *Page) Split() (left, right *Page) {
	return p.SplitAt(p.Len() / 2)
}
