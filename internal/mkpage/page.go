// Package mkpage is the multi-key page (MKSV) of the redesign: one object that
// holds many entries with one value each, in key order, as bytes with no pointer in
// it. Design: docs/redesign/step4-mksv-design.md. The package is the standalone
// stage of step 4.1 and not yet part of the tree.
//
// A page is an object of 32, 64, 128, 256, 384 or 512 bytes. The page of
// values of variable length (strings), Page:
//
//	type | n | cpl | rl 1 ... rl n | common prefix (cpl bytes) | remainder 1, vl 1, value 1 | remainder 2, vl 2, value 2 | ...
//
// and the page of values of one size (a pointer-free T), Fixed:
//
//	type | n | cpl | rl 1 ... rl n | common prefix | remainder 1 ... remainder n | padding to the alignment of T | value 1 ... value n
//
// type is the size class (plus TypeBase, in steps of two), n the number of entries
// (1 to 255), cpl the length of the common prefix: the bytes that all keys of the
// page share after the path to it, stored once. rl i is the length of the
// remainder of entry i, what follows the common prefix. The entries are in key
// order, bytewise, a key that is the prefix of another first: the entry whose
// key ends where the common prefix ends has an empty remainder, and there is at
// most one. The entries are different. The bytes behind the last entry are zero.
//
// The page knows nothing of the tree beyond what its methods take: the tree hands
// in the key from the end of the path on (its "rest"), the page compares the
// common prefix and finds the entry. A key that does not start with the common
// prefix is not in the page, and the tree puts a byte node above the page for it
// (Skip) instead of adding it.
package mkpage

import (
	"unsafe"
)

const (
	// Header is the size of the fixed part: type, n, cpl.
	Header = 3
	// MaxEntries is the largest number of entries of a page: n is one byte. No page
	// of 512 bytes reaches it (an entry takes three bytes at least), so only
	// BuildStrings and BuildFixed look at it.
	MaxEntries = 255
	// MaxRemainder is the longest remainder of an entry, and MaxPrefix the longest
	// common prefix: one byte each. An entry or a set of keys beyond them has no page.
	MaxRemainder = 255
	MaxPrefix    = 255
	// MaxValue is the longest value of a Page: as in the single-key page, a longer
	// one belongs in a value overflow.
	MaxValue = 254
)

// sizes are the object sizes of the classes, those of the single-key page.
var sizes = [...]int{32, 64, 128, 256, 384, 512}

// Classes is the number of size classes.
const Classes = len(sizes)

// TypeBase is added to the type byte of every page, so that a tree whose objects
// tell their type by their first byte gives the pages the types TypeBase,
// TypeBase+2, ... TypeBase+2*(Classes-1). A package that uses pages standalone
// leaves it 0.
var TypeBase uint8

// head is the first three bytes of every page, of both flavors.
type head struct {
	objType uint8 // TypeBase plus twice the size class
	n       uint8 // number of entries
	cpl     uint8 // length of the common prefix
}

// Page is the page of values of variable length (strings), see the package
// comment; Fixed is the page of values of one size.
type Page struct{ head }

// Result says what an Insert did.
type Result int

const (
	// Added: the entry was not there and is now.
	Added Result = iota
	// Present: the entry was there with this value, nothing changed.
	Present
	// Differs: the key is there with another value. The page is unchanged: the
	// key gets a second value, which no multi-key page holds, so the tree builds
	// the subtree again (promote).
	Differs
	// Full: the entry does not go into the page (the content would exceed the
	// largest class, or the remainder or the value are too long). The
	// page is unchanged; the tree bursts it.
	Full
	// Outside: the key does not start with the common prefix of the page. The page
	// is unchanged; the tree puts a byte node above it.
	Outside
)

func (p *head) class() int { return int(p.objType-TypeBase) >> 1 }

// mem returns the whole object.
func (p *head) mem() []byte { return unsafe.Slice((*byte)(unsafe.Pointer(p)), sizes[p.class()]) }

// Size returns the size of the page's object in bytes.
func (p *head) Size() int { return sizes[p.class()] }

// Class returns the index of the page's size class, 0 for 32 bytes.
func (p *head) Class() int { return p.class() }

// Len returns the number of entries.
func (p *head) Len() int { return int(p.n) }

// CP returns the common prefix. The slice aliases the page and is valid until
// the page changes.
func (p *head) CP() []byte { return p.mem()[Header+int(p.n) : Header+int(p.n)+int(p.cpl)] }

// Match returns how many bytes of the common prefix rest starts with. The page
// holds rest if the result is the length of the common prefix.
func (p *head) Match(rest []byte) int { return lcp(p.CP(), rest) }

// PrefixLen returns the length of the common prefix.
func (p *head) PrefixLen() int { return int(p.cpl) }

func lcp(a, b []byte) int {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[i] == b[i] {
		i++
	}
	return i
}

// allocRaw returns a zeroed object of class c that holds no pointer: an array of
// words, which the garbage collector never scans.
func allocRaw(c int) unsafe.Pointer {
	switch c {
	case 0:
		return unsafe.Pointer(new([4]uint64))
	case 1:
		return unsafe.Pointer(new([8]uint64))
	case 2:
		return unsafe.Pointer(new([16]uint64))
	case 3:
		return unsafe.Pointer(new([32]uint64))
	case 4:
		return unsafe.Pointer(new([48]uint64))
	}
	return unsafe.Pointer(new([64]uint64))
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

// NeedStrings returns the bytes a Page takes for n entries whose remainders (after
// a common prefix of cpl bytes) are remBytes in all and whose values are valBytes
// in all. The tree calls it to decide, before it builds anything, whether the
// entries of a subtree fit a page (at most 512).
func NeedStrings(n, cpl, remBytes, valBytes int) int {
	return Header + n + cpl + remBytes + n + valBytes
}

// grow returns a new page of class c with the first used bytes of p: the layout
// does not depend on the class, so the content is where it was.
func grow(p *head, c, used int) unsafe.Pointer {
	q := allocRaw(c)
	copy(unsafe.Slice((*byte)(q), used), p.mem()[:used])
	(*head)(q).objType = TypeBase + uint8(c)<<1
	return q
}

// prefixOf returns the common prefix of the first and the last of rests, the
// longest one a page takes.
func prefixOf(rests [][]byte) int {
	return min(lcp(rests[0], rests[len(rests)-1]), MaxPrefix)
}

// BuildStrings returns a page for the entries with the keys rests (the keys from the
// end of the path on, in key order, all different) and the values vals, or nil if
// they do not fit: no entry, more than MaxEntries, a remainder beyond MaxRemainder, a
// value beyond MaxValue or content beyond the largest class. The common prefix of
// the page is the longest one all keys share (up to MaxPrefix). The tree calls it
// when two single-value entries meet and when it builds the subtree of a burst.
func BuildStrings(rests, vals [][]byte) *Page {
	n := len(rests)
	if n == 0 || n > MaxEntries || len(vals) != n {
		return nil
	}
	cp := prefixOf(rests)
	need := Header + n + cp
	for i, r := range rests {
		if len(r)-cp > MaxRemainder || len(vals[i]) > MaxValue {
			return nil
		}
		need += len(r) - cp + 1 + len(vals[i])
	}
	c := classFor(need)
	if c < 0 {
		return nil
	}
	p := (*Page)(allocRaw(c))
	p.objType, p.n, p.cpl = TypeBase+uint8(c)<<1, uint8(n), uint8(cp)
	m := p.mem()
	off := Header + n
	off += copy(m[off:], rests[0][:cp])
	for i, r := range rests {
		m[Header+i] = uint8(len(r) - cp)
		off += copy(m[off:], r[cp:])
		m[off] = uint8(len(vals[i]))
		off += 1 + copy(m[off+1:], vals[i])
	}
	return p
}

// locate returns the position of the entry with the remainder r, or where it would
// go, with the offset of that entry in the object (the used size, if it goes at
// the end).
func (p *Page) locate(r []byte) (pos, off int, found bool) {
	m := p.mem()
	off = Header + int(p.n) + int(p.cpl)
	for i := range int(p.n) {
		rl := int(m[Header+i])
		switch c := compare(m[off:off+rl], r); {
		case c == 0:
			return i, off, true
		case c > 0:
			return i, off, false
		}
		off += rl + 1 + int(m[off+rl])
	}
	return int(p.n), off, false
}

// usedFrom returns the used size of the page, given the offset of its entry pos.
func (p *Page) usedFrom(pos, off int) int {
	m := p.mem()
	for i := pos; i < int(p.n); i++ {
		off += int(m[Header+i]) + 1 + int(m[off+int(m[Header+i])])
	}
	return off
}

// Used returns the bytes of the page that hold something.
func (p *Page) Used() int {
	return p.usedFrom(0, Header+int(p.n)+int(p.cpl))
}

// Get returns the value of the key rest (the key from the end of the path on), or
// false if the page does not hold it. The slice aliases the page.
func (p *Page) Get(rest []byte) ([]byte, bool) {
	if p.Match(rest) < int(p.cpl) {
		return nil, false
	}
	r := rest[p.cpl:]
	_, off, found := p.locate(r)
	if !found {
		return nil, false
	}
	m := p.mem()
	vl := int(m[off+len(r)])
	return m[off+len(r)+1 : off+len(r)+1+vl], true
}

// Insert adds the entry rest -> val and returns the page that holds the result,
// which is p itself unless the content no longer fits p's class, and says what
// happened (see Result). val is copied.
func (p *Page) Insert(rest, val []byte) (*Page, Result) {
	if p.Match(rest) < int(p.cpl) {
		return p, Outside
	}
	if len(val) > MaxValue {
		return p, Full
	}
	r := rest[p.cpl:]
	pos, off, found := p.locate(r)
	m := p.mem()
	if found {
		if vl := int(m[off+len(r)]); string(m[off+len(r)+1:off+len(r)+1+vl]) == string(val) {
			return p, Present
		}
		return p, Differs
	}
	if len(r) > MaxRemainder {
		return p, Full
	}
	used := p.usedFrom(pos, off)
	elen := len(r) + 1 + len(val)
	need := used + 1 + elen
	q := p
	if need > p.Size() {
		c := classFor(need)
		if c < 0 {
			return p, Full
		}
		q = (*Page)(grow(&p.head, c, used))
		m = q.mem()
	}
	a := Header + pos
	copy(m[off+1+elen:], m[off:used])
	copy(m[a+1:], m[a:off])
	m[a] = uint8(len(r))
	o := off + 1
	o += copy(m[o:], r)
	m[o] = uint8(len(val))
	copy(m[o+1:], val)
	q.n++
	return q, Added
}

// Remove removes the entry rest -> val and reports whether it was there. It returns
// the page that holds the rest: p itself, a page of a smaller class once the
// content fills at most half of it, or nil if the entry was the only one (the
// page is gone). Like Fixed, a page grows by the smallest class that holds one more
// entry, so a page at a class border does not change its object with every
// insert and remove.
func (p *Page) Remove(rest, val []byte) (*Page, bool) {
	if p.Match(rest) < int(p.cpl) {
		return p, false
	}
	r := rest[p.cpl:]
	pos, off, found := p.locate(r)
	if !found {
		return p, false
	}
	m := p.mem()
	vl := int(m[off+len(r)])
	if string(m[off+len(r)+1:off+len(r)+1+vl]) != string(val) {
		return p, false
	}
	if p.n == 1 {
		return nil, true
	}
	elen := len(r) + 1 + vl
	used := p.usedFrom(pos, off)
	a := Header + pos
	copy(m[a:], m[a+1:off])
	copy(m[off-1:], m[off+elen:used])
	used -= 1 + elen
	clear(m[used : used+1+elen])
	p.n--
	if c := classFor(2 * used); c >= 0 && c < p.class() {
		return (*Page)(grow(&p.head, c, used)), true
	}
	return p, true
}

// Each calls fn with the remainder (after the common prefix) and the value of every
// entry in key order until fn returns false, and reports whether it ran to
// completion. The slices alias the page. The key of an entry is the path, then
// CP, then the remainder.
func (p *Page) Each(fn func(rem, val []byte) bool) bool {
	m := p.mem()
	off := Header + int(p.n) + int(p.cpl)
	for i := range int(p.n) {
		rl := int(m[Header+i])
		vl := int(m[off+rl])
		if !fn(m[off:off+rl], m[off+rl+1:off+rl+1+vl]) {
			return false
		}
		off += rl + 1 + vl
	}
	return true
}

// Skip drops the first k bytes of the common prefix (k at most its length): the
// tree has put a byte node above the page that consumes them. It happens in place.
func (p *Page) Skip(k int) {
	m := p.mem()
	used := p.Used()
	at := Header + int(p.n)
	copy(m[at:], m[at+k:used])
	clear(m[used-k : used])
	p.cpl -= uint8(k)
}

// Prepend returns the page with pre in front of the common prefix: the page moves
// up in the tree, when the node above it goes away. It is p itself if the content
// still fits, else a page of a larger class; nil if the common prefix would exceed
// MaxPrefix or the content the largest class.
func (p *Page) Prepend(pre []byte) *Page {
	used := p.Used()
	if int(p.cpl)+len(pre) > MaxPrefix || used+len(pre) > sizes[Classes-1] {
		return nil
	}
	q := p
	if used+len(pre) > p.Size() {
		q = (*Page)(grow(&p.head, classFor(used+len(pre)), used))
	}
	m := q.mem()
	at := Header + int(q.n)
	copy(m[at+len(pre):], m[at:used])
	copy(m[at:], pre)
	q.cpl += uint8(len(pre))
	return q
}

// Equal reports whether two pages hold the same entries (key from the end of the
// path on, and value), whatever their common prefixes and classes. Tests use it.
func Equal(a, b *Page) bool {
	if a.n != b.n {
		return false
	}
	var ka, va, kb, vb []string
	collect := func(p *Page) (ks, vs []string) {
		p.Each(func(rem, val []byte) bool {
			ks = append(ks, string(p.CP())+string(rem))
			vs = append(vs, string(val))
			return true
		})
		return
	}
	ka, va = collect(a)
	kb, vb = collect(b)
	return equalStrings(ka, kb) && equalStrings(va, vb)
}

func equalStrings(a, b []string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// compare is bytes.Compare for the short remainders of a page: most differ in the
// first bytes, and the call of the library's routine costs more than the compare.
func compare(a, b []byte) int {
	n := min(len(a), len(b))
	for i := range n {
		if a[i] != b[i] {
			if a[i] < b[i] {
				return -1
			}
			return 1
		}
	}
	return len(a) - len(b)
}
