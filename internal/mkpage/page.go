// Package mkpage is the multi-key page of the redesign: one object that holds many
// entries, each a key with one or more values, in key order, as bytes with no pointer
// in it. Design: docs/redesign/step4-mksv-design.md (the page for entries with one
// value) and docs/redesign/step5-mkmv-design.md (the values of a key).
//
// A page is an object of 32, 64, 128, 256, 384 or 512 bytes. The page of
// values of variable length (strings), Page:
//
//	type | cpl | n | common prefix (cpl bytes) | rl 1 ... rl n | remainder 1, vl 1, value 1 | vl 2, value 2 | remainder 3, ...
//
// and the page of values of one size (a pointer-free T), Fixed:
//
//	type | cpl | n | common prefix | rl 1 ... rl n | remainders | padding to the alignment of T | value 1 ... value n
//
// type is the size class (plus TypeBase, in steps of two); its lowest bit is bit 8 of
// the next byte, so that cpl, the length of the common prefix (the bytes that all keys of
// the page share after the path to it, stored once), has nine bits: the head starts like
// the head of a single-key page (type, length of the key part stored). n is the number
// of values (1 to 255, a "slot" each), rl i the length of the remainder of slot i, what
// follows the common prefix, or Further (255) for a slot that holds a further value of the
// key before it: it has no remainder. The keys are in order, bytewise, a key that is the
// prefix of another first: the key that ends where the common prefix ends has an empty
// remainder, and there is at most one. The keys are different; the values of a key sit
// next to each other in the order they came in. A page whose keys have one value each has
// no Further. The bytes behind the last entry are zero.
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
	// Header is the size of the fixed part: type, cpl, n.
	Header = 3
	// MaxEntries is the largest number of slots (values) of a page: n is one byte. No page
	// of 512 bytes reaches it (a slot takes two bytes at least, and the head and the common
	// prefix come on top), so only BuildStrings and BuildFixed look at it.
	MaxEntries = 255
	// Further is the length byte of a slot that holds a further value of the key before it.
	// A remainder of that length would be taken for it, so MaxRemainder is one less.
	Further = 255
	// MaxRemainder is the longest remainder of an entry, and MaxPrefix the longest
	// common prefix (nine bits, as the remainder of a single-key page). An entry or a set of
	// keys beyond them has no page.
	MaxRemainder = 254
	MaxPrefix    = 511
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
// TypeBase+2, ... TypeBase+2*(Classes-1), and one more for a common prefix of 256 bytes or
// more. TypeBase must be even. A package that uses pages standalone leaves it 0.
var TypeBase uint8

// head is the first three bytes of every page, of both flavors.
type head struct {
	objType uint8 // TypeBase plus twice the size class; the lowest bit is bit 8 of klen
	klen    uint8 // bits 0 to 7 of the length of the common prefix
	n       uint8 // number of slots
}

// Page is the page of values of variable length (strings), see the package
// comment; Fixed is the page of values of one size.
type Page struct{ head }

// Result says what an Insert or Add did.
type Result int

const (
	// Added: the key was not there and is now.
	Added Result = iota
	// AddedValue: the key was there with other values and has one more now (Add only).
	AddedValue
	// Present: the key was there with this value, nothing changed.
	Present
	// Differs: the key is there with another value, and Insert, which knows one value a key, left
	// the page unchanged. The tree builds the subtree again (promote). Add never answers it.
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

// cpl returns the length of the common prefix, nine bits.
func (p *head) cpl() int { return int(p.klen) | int(p.objType&1)<<8 }

// setCpl sets the length of the common prefix and keeps the type.
func (p *head) setCpl(n int) {
	p.objType = p.objType&^1 | uint8(n>>8)
	p.klen = uint8(n)
}

// mem returns the whole object.
func (p *head) mem() []byte { return unsafe.Slice((*byte)(unsafe.Pointer(p)), sizes[p.class()]) }

// Size returns the size of the page's object in bytes.
func (p *head) Size() int { return sizes[p.class()] }

// Class returns the index of the page's size class, 0 for 32 bytes.
func (p *head) Class() int { return p.class() }

// Len returns the number of values (slots) of the page; it is the number of keys if every key
// has one value.
func (p *head) Len() int { return int(p.n) }

// Keys returns the number of keys of the page.
func (p *head) Keys() int {
	m := p.mem()
	lo := Header + p.cpl()
	k := 0
	for _, rl := range m[lo : lo+int(p.n)] {
		k += b2i(rl != Further)
	}
	return k
}

// CP returns the common prefix. The slice aliases the page and is valid until
// the page changes.
func (p *head) CP() []byte { return p.mem()[Header : Header+p.cpl()] }

// Match returns how many bytes of the common prefix rest starts with. The page
// holds rest if the result is the length of the common prefix.
func (p *head) Match(rest []byte) int { return lcp(p.CP(), rest) }

// PrefixLen returns the length of the common prefix.
func (p *head) PrefixLen() int { return p.cpl() }

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

// NeedStrings returns the bytes a Page takes for n slots whose remainders (after
// a common prefix of cpl bytes) are remBytes in all (the slots with Further have none) and
// whose values are valBytes in all. The tree calls it to decide, before it builds
// anything, whether the entries of a subtree fit a page (at most 512).
func NeedStrings(n, cpl, remBytes, valBytes int) int {
	return Header + n + cpl + remBytes + n + valBytes
}

// grow returns a new page of class c with the first used bytes of p: the layout
// does not depend on the class, so the content is where it was.
func grow(p *head, c, used int) unsafe.Pointer {
	q := allocRaw(c)
	copy(unsafe.Slice((*byte)(q), used), p.mem()[:used])
	(*head)(q).objType = TypeBase + uint8(c)<<1 | p.objType&1
	return q
}

// newHead sets the head of a new page of class c with n slots and a common prefix of cpl bytes.
func newHead(p *head, c, n, cpl int) {
	p.objType, p.n = TypeBase+uint8(c)<<1, uint8(n)
	p.setCpl(cpl)
}

// prefixOf returns the common prefix of the first and the last of rests, the
// longest one a page takes.
func prefixOf(rests [][]byte) int {
	return min(lcp(rests[0], rests[len(rests)-1]), MaxPrefix)
}

// further reports whether rests[i] is a further value of the key of rests[i-1].
func further(rests [][]byte, i int) bool {
	return i > 0 && string(rests[i]) == string(rests[i-1])
}

// BuildStrings returns a page for the values vals of the keys rests (the keys from the
// end of the path on, in key order: a key with several values appears once for each, one
// after the other), or nil if they do not fit: no value, more than MaxEntries, a remainder
// beyond MaxRemainder, a value beyond MaxValue or content beyond the largest class. The
// common prefix of the page is the longest one all keys share (up to MaxPrefix). The tree
// calls it when single-key pages meet and when it builds the subtree of a burst.
func BuildStrings(rests, vals [][]byte) *Page {
	n := len(rests)
	if n == 0 || n > MaxEntries || len(vals) != n {
		return nil
	}
	cp := prefixOf(rests)
	need := Header + n + cp
	for i, r := range rests {
		if !further(rests, i) {
			if len(r)-cp > MaxRemainder {
				return nil
			}
			need += len(r) - cp
		}
		if len(vals[i]) > MaxValue {
			return nil
		}
		need += 1 + len(vals[i])
	}
	c := classFor(need)
	if c < 0 {
		return nil
	}
	p := (*Page)(allocRaw(c))
	newHead(&p.head, c, n, cp)
	m := p.mem()
	lo := Header + cp
	off := lo + n
	copy(m[Header:], rests[0][:cp])
	for i, r := range rests {
		if further(rests, i) {
			m[lo+i] = Further
		} else {
			m[lo+i] = uint8(len(r) - cp)
			off += copy(m[off:], r[cp:])
		}
		m[off] = uint8(len(vals[i]))
		off += 1 + copy(m[off+1:], vals[i])
	}
	return p
}

// locate returns the position (slot) of the first value of the key with the remainder r, or
// of the key that would follow it, with the offset of that slot's entry in the object (the used
// size, if it goes at the end).
func (p *Page) locate(r []byte) (pos, off int, found bool) {
	m := p.mem()
	n := int(p.n)
	lo := Header + p.cpl()
	off = lo + n
	for i, rl := range m[lo : lo+n] {
		if rl == Further {
			off += 1 + int(m[off])
			continue
		}
		switch c := compare(m[off:off+int(rl)], r); {
		case c == 0:
			return i, off, true
		case c > 0:
			return i, off, false
		}
		off += int(rl) + 1 + int(m[off+int(rl)])
	}
	return n, off, false
}

// usedFrom returns the used size of the page, given the offset of its slot pos.
func (p *Page) usedFrom(pos, off int) int {
	m := p.mem()
	lo := Header + p.cpl()
	for i := pos; i < int(p.n); i++ {
		if rl := int(m[lo+i]); rl == Further {
			off += 1 + int(m[off])
		} else {
			off += rl + 1 + int(m[off+rl])
		}
	}
	return off
}

// Used returns the bytes of the page that hold something.
func (p *Page) Used() int {
	return p.usedFrom(0, Header+p.cpl()+int(p.n))
}

// Get returns the first value of the key rest (the key from the end of the path on), or
// false if the page does not hold it. The slice aliases the page.
func (p *Page) Get(rest []byte) ([]byte, bool) {
	if p.Match(rest) < p.cpl() {
		return nil, false
	}
	r := rest[p.cpl():]
	_, off, found := p.locate(r)
	if !found {
		return nil, false
	}
	m := p.mem()
	vl := int(m[off+len(r)])
	return m[off+len(r)+1 : off+len(r)+1+vl], true
}

// EachValue calls fn with every value of the key rest, in the order they came in, until fn
// returns false, and reports whether the page holds the key. The slices alias the page.
func (p *Page) EachValue(rest []byte, fn func(val []byte) bool) bool {
	if p.Match(rest) < p.cpl() {
		return false
	}
	r := rest[p.cpl():]
	pos, off, found := p.locate(r)
	if !found {
		return false
	}
	m := p.mem()
	n := int(p.n)
	lo := Header + p.cpl()
	o := off + len(r)
	for i := pos; ; {
		vl := int(m[o])
		if !fn(m[o+1 : o+1+vl]) {
			return true
		}
		o += 1 + vl
		if i++; i >= n || m[lo+i] != Further {
			return true
		}
	}
}

// Insert is Add for a page whose keys have one value each: a key that is there with another
// value is left alone, and the answer is Differs (the tree builds it again). Add replaces it
// once the tree holds several values in a page.
func (p *Page) Insert(rest, val []byte) (*Page, Result) { return p.add(rest, val, false) }

// Add adds the value val to the key rest and returns the page that holds the result, which
// is p itself unless the content no longer fits p's class, and says what happened (see
// Result). val is copied. A key that is there gets the value behind its others.
func (p *Page) Add(rest, val []byte) (*Page, Result) { return p.add(rest, val, true) }

func (p *Page) add(rest, val []byte, multi bool) (*Page, Result) {
	cpl := p.cpl()
	if p.Match(rest) < cpl {
		return p, Outside
	}
	r := rest[cpl:]
	pos, off, found := p.locate(r)
	m := p.mem()
	rl := uint8(len(r))
	slot, at := pos, off
	if found {
		n := int(p.n)
		lo := Header + cpl
		o := off + len(r)
		for {
			vl := int(m[o])
			if string(m[o+1:o+1+vl]) == string(val) {
				return p, Present
			}
			o += 1 + vl
			if slot++; slot >= n || m[lo+slot] != Further {
				break
			}
		}
		if !multi {
			return p, Differs
		}
		rl, r, at = Further, nil, o
	}
	if len(r) > MaxRemainder || len(val) > MaxValue {
		return p, Full
	}
	used := p.usedFrom(slot, at)
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
	a := Header + cpl + slot
	copy(m[at+1+elen:], m[at:used])
	copy(m[a+1:], m[a:at])
	m[a] = rl
	o := at + 1
	o += copy(m[o:], r)
	m[o] = uint8(len(val))
	copy(m[o+1:], val)
	q.n++
	return q, Added + Result(b2i(found))
}

// Widen adds the entry rest -> val, a key that leaves the common prefix of the page after
// mis = Match(rest) < PrefixLen() bytes, and shortens the common prefix to those mis bytes: the
// page is built again, with the rest of the old prefix in front of every remainder. It returns
// the new page, or nil if the entry does not go in (the content would exceed the largest class, or
// a remainder or the value is too long); p is then unchanged. The tree calls it where it would
// otherwise put a byte node above the page, so that a key that differs early joins its
// neighbours when they fit one page.
func (p *Page) Widen(rest, val []byte) *Page {
	cpl := p.cpl()
	mis := p.Match(rest)
	d := cpl - mis
	n := int(p.n)
	newRem := rest[mis:]
	if len(val) > MaxValue || len(newRem) > MaxRemainder {
		return nil
	}
	m := p.mem()
	lo := Header + cpl
	used := p.Used()
	heads, longest := 0, 0
	for _, rl := range m[lo : lo+n] {
		if rl != Further {
			heads++
			longest = max(longest, int(rl))
		}
	}
	need := used + 1 - d + heads*d + len(newRem) + 1 + len(val)
	c := classFor(need)
	if c < 0 || longest+d > MaxRemainder {
		return nil
	}
	// the key differs from the prefix at mis: it is below every entry or above every one
	first := len(newRem) == 0 || newRem[0] < p.CP()[mis]
	q := (*Page)(allocRaw(c))
	newHead(&q.head, c, n+1, mis)
	w := q.mem()
	nlo := Header + mis
	at := 0 // where the new entry goes
	if !first {
		at = n
	}
	for i := range n {
		rl := m[lo+i]
		if rl != Further {
			rl += uint8(d)
		}
		w[nlo+i+b2i(i >= at)] = rl
	}
	w[nlo+at] = uint8(len(newRem))
	off := nlo + n + 1
	copy(w[Header:], p.CP()[:mis])
	src := lo + n
	extra := p.CP()[mis:]
	put := func() {
		off += copy(w[off:], newRem)
		w[off] = uint8(len(val))
		off += 1 + copy(w[off+1:], val)
	}
	if first {
		put()
	}
	for i := range n {
		if rl := int(m[lo+i]); rl != Further {
			off += copy(w[off:], extra)
			copy(w[off:], m[src:src+rl])
			off += rl
			src += rl
		}
		vl := int(m[src])
		copy(w[off:], m[src:src+1+vl])
		off += 1 + vl
		src += 1 + vl
	}
	if !first {
		put()
	}
	return q
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Remove removes the value val of the key rest and reports whether it was there; a key
// that has no value left is gone. It returns the page that holds the rest: p itself, a page
// of a smaller class once the content fills at most half of it, or nil if the value was the
// only one (the page is gone). Like Fixed, a page grows by the smallest class that holds one
// more value, so a page at a class border does not change its object with every insert and
// remove. The first value of a key with others lets the next one take its place.
func (p *Page) Remove(rest, val []byte) (*Page, bool) {
	cpl := p.cpl()
	if p.Match(rest) < cpl {
		return p, false
	}
	r := rest[cpl:]
	pos, off, found := p.locate(r)
	if !found {
		return p, false
	}
	m := p.mem()
	n := int(p.n)
	lo := Header + cpl
	slot, vo := pos, off+len(r) // the slot and the offset of the length of the value looked at
	for {
		vl := int(m[vo])
		if string(m[vo+1:vo+1+vl]) == string(val) {
			break
		}
		vo += 1 + vl
		if slot++; slot >= n || m[lo+slot] != Further {
			return p, false
		}
	}
	vl := int(m[vo])
	more := slot+1 < n && m[lo+slot+1] == Further
	from, to := vo, vo+1+vl   // what goes: the value, with its length, ...
	if slot == pos && !more { // ... and the key's remainder, if the value is the only one
		if n == 1 {
			return nil, true
		}
		from = off
	}
	used := p.usedFrom(slot, vo-b2i(slot == pos)*len(r))
	if slot == pos && more { // the next value takes the key's place: its slot becomes the key's
		m[lo+slot+1] = uint8(len(r))
	}
	a := lo + slot
	copy(m[a:], m[a+1:from])
	copy(m[from-1:], m[to:used])
	used -= 1 + to - from
	clear(m[used : used+1+to-from])
	p.n--
	if c := classFor(2 * used); c >= 0 && c < p.class() {
		return (*Page)(grow(&p.head, c, used)), true
	}
	return p, true
}

// Each calls fn with the remainder (after the common prefix), the value and whether it is the
// first value of its key, for every value in key order until fn returns false, and reports
// whether it ran to completion. The slices alias the page; the remainder of a further value
// is that of its key. The key of an entry is the path, then CP, then the remainder.
func (p *Page) Each(fn func(rem, val []byte, first bool) bool) bool {
	m := p.mem()
	n := int(p.n)
	lo := Header + p.cpl()
	off := lo + n
	var rem []byte
	for i := range n {
		rl := int(m[lo+i])
		first := rl != Further
		if first {
			rem = m[off : off+rl]
			off += rl
		}
		vl := int(m[off])
		if !fn(rem, m[off+1:off+1+vl], first) {
			return false
		}
		off += 1 + vl
	}
	return true
}

// Skip drops the first k bytes of the common prefix (k at most its length): the
// tree has put a byte node above the page that consumes them. It happens in place.
func (p *Page) Skip(k int) {
	m := p.mem()
	used := p.Used()
	copy(m[Header:], m[Header+k:used])
	clear(m[used-k : used])
	p.setCpl(p.cpl() - k)
}

// Prepend returns the page with pre in front of the common prefix: the page moves
// up in the tree, when the node above it goes away. It is p itself if the content
// still fits, else a page of a larger class; nil if the common prefix would exceed
// MaxPrefix or the content the largest class.
func (p *Page) Prepend(pre []byte) *Page {
	used := p.Used()
	if p.cpl()+len(pre) > MaxPrefix || used+len(pre) > sizes[Classes-1] {
		return nil
	}
	q := p
	if used+len(pre) > p.Size() {
		q = (*Page)(grow(&p.head, classFor(used+len(pre)), used))
	}
	m := q.mem()
	copy(m[Header+len(pre):], m[Header:used])
	copy(m[Header:], pre)
	q.setCpl(q.cpl() + len(pre))
	return q
}

// Equal reports whether two pages hold the same values of the same keys (key from the end of the
// path on), whatever their common prefixes and classes. Tests use it.
func Equal(a, b *Page) bool {
	if a.n != b.n {
		return false
	}
	collect := func(p *Page) (ks, vs []string) {
		p.Each(func(rem, val []byte, _ bool) bool {
			ks = append(ks, string(p.CP())+string(rem))
			vs = append(vs, string(val))
			return true
		})
		return
	}
	ka, va := collect(a)
	kb, vb := collect(b)
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
