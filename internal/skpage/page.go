// Package skpage is the single-key page (SKMV) of the redesign for values of
// variable length, in particular strings: one object that holds the remainder
// of one key and all its values as bytes, with no pointer in it. Design:
// docs/redesign/step3-skmv-design.md. The package is a prototype for the
// measurements of step 3.2 and not yet part of the tree.
//
// A page is an object of 32, 64, 128, 256, 384 or 512 bytes:
//
//	kind | r | n (2 bytes) | kl (2 bytes) | remainder (r bytes) | length 1, value 1 | length 2, value 2 | ...
//
// kind is the size class (plus KindBase), r the length of the remainder (0 to
// 254), n the number of values (1 to 254: every value takes a byte at least, and
// the largest class has 512), kl the length of the whole key. The first six
// bytes are those of the leaves of internal/art (kind, key remainder length,
// number of values, whole key length), so that the tree's code for keys, which
// holds the key from the base it was made at and compares the end of a key with
// the remainder, works on a page as on a leaf (design note, section 5: the
// header of 3 bytes is a later step). Each value is a length byte (0 to 254)
// followed by that many bytes. The bytes behind the last value are zero. The
// values of a key are a set: no two are equal, and their order is the order of
// their arrival.
//
// The page knows nothing of the tree beyond what the leaves of the tree know:
// the tree hands in the key, of which the page keeps the end.
package skpage

import (
	"bytes"
	"unsafe"
)

const (
	// MaxValue is the longest value a page holds; a longer one belongs in a
	// value set. 255 is not a length, so that a length is one byte.
	MaxValue = 254
	// MaxRemainder is the longest remainder a page holds.
	MaxRemainder = 254
	// Header is the size of the fixed part: kind, r, n, kl.
	Header = 6
	// BackLimit is the content, in bytes, up to which the value set of a key
	// goes back into a page (half of the largest class). A key moves into a
	// set when its content no longer fits the largest class; moving back only
	// at half of that keeps a key at the border from changing its object with
	// every added and removed value.
	BackLimit = 256
)

// sizes are the object sizes of the classes.
var sizes = [...]int{32, 64, 128, 256, 384, 512}

// Classes is the number of size classes.
const Classes = len(sizes)

// KindBase is added to the kind byte of every page, so that a tree whose
// objects tell their kind by their first byte can give the pages the kinds
// KindBase to KindBase+Classes-1. A package that uses pages standalone
// leaves it 0.
var KindBase uint8

// Page is the first three bytes of a page; the page is the object they start.
type Page struct {
	kind uint8  // KindBase plus the size class
	r    uint8  // length of the remainder
	n    uint16 // number of values
	kl   uint16 // length of the whole key
}

// Result says what Add did.
type Result int

const (
	// Added: the value was not there and is now.
	Added Result = iota
	// Present: the value was there, nothing changed.
	Present
	// Full: the value cannot go into a page (it is longer than MaxValue or the
	// content would exceed the largest class). The page is unchanged; the key
	// belongs in a value set.
	Full
)

func (p *Page) class() int { return int(p.kind - KindBase) }

// mem returns the whole object.
func (p *Page) mem() []byte { return unsafe.Slice((*byte)(unsafe.Pointer(p)), sizes[p.class()]) }

// alloc returns a zeroed object of class c.
func alloc(c int) *Page {
	switch c {
	case 0:
		return (*Page)(unsafe.Pointer(new([4]uint64)))
	case 1:
		return (*Page)(unsafe.Pointer(new([8]uint64)))
	case 2:
		return (*Page)(unsafe.Pointer(new([16]uint64)))
	case 3:
		return (*Page)(unsafe.Pointer(new([32]uint64)))
	case 4:
		return (*Page)(unsafe.Pointer(new([48]uint64)))
	}
	return (*Page)(unsafe.Pointer(new([64]uint64)))
}

// classFor returns the smallest class that holds need bytes, or -1.
func classFor(need int) int {
	for c, s := range sizes {
		if need <= s {
			return c
		}
	}
	return -1
}

// newPage returns an empty page of class c for a remainder of r bytes of a key
// of kl bytes.
func newPage(c, r, kl int) *Page {
	p := alloc(c)
	p.kind, p.r, p.kl = KindBase+uint8(c), uint8(r), uint16(kl)
	return p
}

// New returns a page for the key of keyLen bytes whose end is rest, with the one
// value val, or nil if they do not fit a page (see MaxRemainder and MaxValue; the largest class always
// holds the rest).
// The page copies both. The tree calls New when a key arrives that no page holds
// yet, with the remainder it has cut from the key.
func New(rest []byte, keyLen int, val []byte) *Page {
	if len(rest) > MaxRemainder || len(val) > MaxValue {
		return nil
	}
	p := newPage(classFor(Header+len(rest)+1+len(val)), len(rest), keyLen) // at most 512
	p.n = 1
	m := p.mem()
	copy(m[Header:], rest)
	m[Header+len(rest)] = uint8(len(val))
	copy(m[Header+len(rest)+1:], val)
	return p
}

// Empty returns a page without values for the key of keyLen bytes whose end is
// rest, or nil if rest is longer than MaxRemainder. The tree makes it when a key
// arrives and adds the first value at once with Add; a page without values is
// not a state a key stays in. Do not use any other operation on it.
func Empty(rest []byte, keyLen int) *Page {
	if len(rest) > MaxRemainder {
		return nil
	}
	p := newPage(classFor(Header+len(rest)+1), len(rest), keyLen)
	copy(p.mem()[Header:], rest)
	return p
}

// Build returns a page for the key of keyLen bytes whose end is rest, with the
// values vals (all different),
// or nil if they do not fit. The tree calls Build when the value set of a key
// has shrunk to BackLimit bytes of content or less.
func Build(rest []byte, keyLen int, vals [][]byte) *Page {
	if len(rest) > MaxRemainder || len(vals) == 0 {
		return nil
	}
	need := Header + len(rest)
	for _, v := range vals {
		if len(v) > MaxValue {
			return nil
		}
		need += 1 + len(v)
		if need > sizes[Classes-1] {
			return nil // also keeps n below 256
		}
	}
	p := newPage(classFor(need), len(rest), keyLen)
	p.n = uint16(len(vals))
	m := p.mem()
	off := Header + copy(m[Header:], rest)
	for _, v := range vals {
		m[off] = uint8(len(v))
		off += 1 + copy(m[off+1:], v)
	}
	return p
}

// Size returns the size of the page's object in bytes.
func (p *Page) Size() int { return sizes[p.class()] }

// Class returns the index of the page's size class, 0 for 32 bytes.
func (p *Page) Class() int { return p.class() }

// Len returns the number of values.
func (p *Page) Len() int { return int(p.n) }

// Rest returns the remainder. The slice aliases the page and is valid until the
// page changes.
func (p *Page) Rest() []byte { return p.mem()[Header : Header+int(p.r)] }

// KeyLen returns the length of the whole key.
func (p *Page) KeyLen() int { return int(p.kl) }

// Match reports whether key is the key of the page: it has the length of the
// whole key, and its last bytes are the remainder. The tree has matched the
// bytes before them on its way down.
func (p *Page) Match(key []byte) bool {
	r := int(p.r)
	return len(key) == int(p.kl) && string(key[len(key)-r:]) == unsafe.String((*byte)(unsafe.Add(unsafe.Pointer(p), Header)), r)
}

// Used returns the bytes of the page that hold something: the header, the
// remainder and the values with their lengths.
func (p *Page) Used() int {
	m := p.mem()
	off := Header + int(p.r)
	for range p.n {
		off += 1 + int(m[off])
	}
	return off
}

// find returns the offset of the length byte of the value val, or -1, and the
// end of the used part.
func (p *Page) find(val []byte) (at, used int) {
	m := p.mem()
	off, at := Header+int(p.r), -1
	for range p.n {
		l := int(m[off])
		if at < 0 && l == len(val) && string(m[off+1:off+1+l]) == string(val) {
			at = off
		}
		off += 1 + l
	}
	return at, off
}

// Has reports whether val is one of the values.
func (p *Page) Has(val []byte) bool {
	if len(val) > MaxValue {
		return false
	}
	m := p.mem()
	off := Header + int(p.r)
	for range p.n {
		l := int(m[off])
		if l == len(val) && string(m[off+1:off+1+l]) == string(val) {
			return true
		}
		off += 1 + l
	}
	return false
}

// Add adds val to the values of the key. It returns the page that holds the
// result, which is p itself unless the content no longer fits p's class, and
// says what happened. With Full the page is unchanged.
func (p *Page) Add(val []byte) (*Page, Result) {
	if len(val) > MaxValue {
		return p, Full
	}
	at, used := p.find(val)
	switch {
	case at >= 0:
		return p, Present
	}
	need := used + 1 + len(val)
	q := p
	if need > p.Size() {
		c := classFor(need)
		if c < 0 {
			return p, Full
		}
		q = newPage(c, int(p.r), int(p.kl))
		q.n = p.n
		copy(q.mem()[Header:], p.mem()[Header:used])
	}
	m := q.mem()
	m[used] = uint8(len(val))
	copy(m[used+1:], val)
	q.n++
	return q, Added
}

// Remove removes val from the values and reports whether it was there. It
// returns the page that holds the rest: p itself, a page of a smaller class if
// the content now fits one that saves at least half of p's size, or nil if val
// was the only value, when the key is gone.
func (p *Page) Remove(val []byte) (*Page, bool) {
	if len(val) > MaxValue {
		return p, false
	}
	at, used := p.find(val)
	if at < 0 {
		return p, false
	}
	if p.n == 1 {
		return nil, true
	}
	m := p.mem()
	end := at + 1 + len(val)
	copy(m[at:], m[end:used])
	clear(m[used-(end-at) : used])
	p.n--
	used -= end - at
	if c := classFor(used); c >= 0 && 2*sizes[c] <= p.Size() {
		q := newPage(c, int(p.r), int(p.kl))
		q.n = p.n
		copy(q.mem()[Header:], m[Header:used])
		return q, true
	}
	return p, true
}

// Each calls fn with every value in the order of their arrival until fn returns
// false, and reports whether it ran to completion. The slices alias the page.
func (p *Page) Each(fn func(val []byte) bool) bool {
	m := p.mem()
	off := Header + int(p.r)
	for range p.n {
		l := int(m[off])
		if !fn(m[off+1 : off+1+l]) {
			return false
		}
		off += 1 + l
	}
	return true
}

// Strings calls fn with every value as a string, in the order of their arrival,
// until fn returns false, and reports whether it ran to completion. The strings
// are copies, made together: one allocation for all the values of the key. A
// string that fn keeps holds the bytes of its neighbours alive.
func (p *Page) Strings(fn func(val string) bool) bool {
	m := p.mem()
	start := Header + int(p.r)
	total, off := 0, start
	for range p.n {
		l := int(m[off])
		total += l
		off += 1 + l
	}
	buf := make([]byte, total)
	at := 0
	off = start
	for range p.n {
		l := int(m[off])
		str := ""
		if l > 0 {
			copy(buf[at:], m[off+1:off+1+l])
			str = unsafe.String(&buf[at], l)
		}
		at += l
		off += 1 + l
		if !fn(str) {
			return false
		}
	}
	return true
}

// AppendKey appends the remainder of the page to dst. The tree assembles the key
// of a scan from the path and the remainder.
func (p *Page) AppendKey(dst []byte) []byte { return append(dst, p.Rest()...) }

// Prepend returns the page with pre in front of the remainder: the page moves
// up in the tree, when the node above it goes away. It is p itself if the
// content still fits, else a page of a larger class; nil if the remainder
// would exceed MaxRemainder or the content the largest class.
func (p *Page) Prepend(pre []byte) *Page {
	r := int(p.r) + len(pre)
	used := p.Used()
	if r > MaxRemainder || used+len(pre) > sizes[Classes-1] {
		return nil
	}
	q := p
	if used+len(pre) > p.Size() {
		q = newPage(classFor(used+len(pre)), r, int(p.kl))
		q.n = p.n
		copy(q.mem()[Header+len(pre):], p.mem()[Header:used])
	} else {
		m := p.mem()
		copy(m[Header+len(pre):], m[Header:used])
	}
	q.r = uint8(r)
	copy(q.mem()[Header:], pre)
	return q
}

// Equal reports whether two pages hold the same remainder and the same values
// in the same order, whatever their classes. Tests use it.
func Equal(a, b *Page) bool {
	return a.n == b.n && a.r == b.r && a.kl == b.kl && bytes.Equal(a.mem()[Header:a.Used()], b.mem()[Header:b.Used()])
}
