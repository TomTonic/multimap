// Package page is the one page of the redesign (docs/redesign/step5-one-page.md): one object of
// 32 to 512 bytes that holds the entries of a subtree, each a key with one value or more, in two
// forms of one structure and in two flavors, for values of variable length (Str, strings) and of
// one size (Fixed, a pointer-free T or one word with a pointer). Before it there were four
// structures (the single-key and the multi-key page, each for both flavors).
//
// Every page begins with the same head of four bytes:
//
//	byte 0  type     the size class (plus TypeOneKey or TypeManyKeys, in steps of two); bit 0 is bit 8 of len
//	byte 1  len      the length of the key part, nine bits
//	byte 2  n        the number of slots, which is the number of values (1 to 255)
//	byte 3  rawWords the size of the byte area in 8-byte words, for a page whose values are pointers; else 0
//
// and the key part from byte 4: in the many-key form the common prefix of all keys (the bytes
// that all keys of the page share after the path to it, stored once), in the one-key form the whole
// remainder of the one key. A page begins exactly at its path length: when the path changes, Skip
// and Prepend change the key part in place. Behind the key part, as far as the page has them:
//
//	key lengths    n bytes, many-key form only: the length of the remainder of slot i, or Further (255)
//	               for a slot that holds a further value of the key before it (no remainder)
//	value lengths  n bytes, Str only: the length of value i
//	remainders     many-key form only: the remainders of the keys, one behind the other
//	free           zero
//	values         they end with the object: Str: the bytes of the values of all slots, in order; Fixed:
//	               an array of T, value i at Size - (n - i) * size of T
//
// So a one-key page is a many-key page with the key part of the whole key and without key
// lengths and remainders; a page of fixed-size values has no value lengths. The type byte says
// the form (the one-key types come before the many-key types, as the tree needs them: a
// descent ends at a type byte up to the last page type, and isSingleKey is a comparison). The
// keys are in order, bytewise, a key that is the prefix of another first; the values of a key sit
// next to each other in the order they came in and are a set (no two equal).
//
// The page knows nothing of the tree beyond what its methods take: the tree hands in the key
// from the end of the path on (its "rest").
package page

import (
	"unsafe"
)

const (
	// Header is the size of the head: type, len, n, rawWords.
	Header = 4
	// MaxEntries is the largest number of slots (values) of a page: n is one byte.
	MaxEntries = 255
	// Further is the key length of a slot that holds a further value of the key before it.
	// A remainder of that length would be taken for it, so MaxRemainder is one less.
	Further = 255
	// MaxRemainder is the longest remainder of a key of the many-key form; MaxKeyPart is the
	// longest key part (nine bits). An entry or a set of keys beyond them has no page.
	MaxRemainder = 254
	MaxKeyPart   = 511
	// MaxValue is the longest value of a Str page: a longer one belongs in a value overflow.
	MaxValue = 254
)

// TypeOneKey and TypeManyKeys are added to the type byte of every page: a one-key page has TypeOneKey+2c,
// a many-key page TypeManyKeys+2c for the size class c, and the one-key types are the smaller. They are the
// two places where the tree's type bytes of the pages begin (art: the single-key page after the value
// overflow, the multi-key page after the single-key pages; a test of the tree checks it); constants, so
// that the form of a page is a comparison with a constant.
const (
	TypeOneKey   uint8 = 4
	TypeManyKeys uint8 = 16
)

// sizes are the object sizes of the classes.
var sizes = [...]int{32, 64, 128, 256, 384, 512}

// Classes is the number of size classes.
const Classes = len(sizes)

// Result says what an Add did.
type Result int

const (
	// Added: the key was not there and is now.
	Added Result = iota
	// AddedValue: the key was there with other values and has one more now.
	AddedValue
	// Present: the key was there with this value, nothing changed.
	Present
	// Full: the entry does not go into the page (the content would exceed the largest class, or
	// a remainder or the value is too long). The page is unchanged; the tree bursts it.
	Full
	// Outside: the key does not start with the common prefix of the page, and Add of the
	// many-key form leaves it as it is. Widen takes such a key in.
	Outside
)

// Removal says what a Remove did.
type Removal int

const (
	// Absent: the key was not there with that value (or is outside the common prefix); the
	// page is unchanged.
	Absent Removal = iota
	// Removed: the value is gone, the key has other values.
	Removed
	// Gone: the value was the last of its key, and the key is gone.
	Gone
)

// head is the first four bytes of every page of either flavor; the page is the object they start.
type head struct {
	objType uint8 // TypeOneKey or TypeManyKeys plus twice the size class, plus bit 8 of the length of the key part
	klen    uint8 // bits 0 to 7 of the length of the key part
	n       uint8 // the number of slots
	raw     uint8 // rawWords
}

// Str is the page of values of variable length (strings); Fixed is the page of values of one size.
type Str struct{ head }

// one reports whether the page has the one-key form.
func (p *head) one() bool { return p.objType < TypeManyKeys }

func (p *head) class() int {
	if p.one() {
		return int(p.objType-TypeOneKey) >> 1
	}
	return int(p.objType-TypeManyKeys) >> 1
}

// cpl returns the length of the key part, nine bits.
func (p *head) cpl() int { return int(p.klen) | int(p.objType&1)<<8 }

// setCpl sets the length of the key part and keeps the type.
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

// Len returns the number of values (slots) of the page.
func (p *head) Len() int { return int(p.n) }

// OneKey reports whether the page has the one-key form: its key part is the whole key and it
// has no key lengths and remainders.
func (p *head) OneKey() bool { return p.one() }

// RawWords returns the size of the byte area in 8-byte words for a page whose values are
// pointers, else 0.
func (p *head) RawWords() int { return int(p.raw) }

// CP returns the key part: the common prefix of the many-key form, the whole remainder of the
// one-key form. The slice aliases the page and is valid until the page changes.
func (p *head) CP() []byte { return p.mem()[Header : Header+p.cpl()] }

// PrefixLen returns the length of the key part.
func (p *head) PrefixLen() int { return p.cpl() }

// Match returns how many bytes of rest the key part matches.
func (p *head) Match(rest []byte) int { return lcp(p.CP(), rest) }

func lcp(a, b []byte) int {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[i] == b[i] {
		i++
	}
	return i
}

// lay holds the offsets of the parts behind the key part of a page with n slots.
type lay struct {
	n, l        int  // slots, length of the key part
	kl, vl, rem int  // offsets: key lengths, value lengths, remainders
	many, str   bool // the many-key form has key lengths and remainders; Str has value lengths
}

func (p *head) lay(str bool) lay {
	la := lay{n: int(p.n), l: p.cpl(), many: !p.one(), str: str}
	la.kl = Header + la.l
	la.vl = la.kl
	if la.many {
		la.vl += la.n
	}
	la.rem = la.vl
	if str {
		la.rem += la.n
	}
	return la
}

// keyEnd returns where the remainders end: the end of the key area.
func (la *lay) keyEnd(m []byte) int {
	if !la.many {
		return la.rem
	}
	return la.keyEndFrom(m, 0, la.rem)
}

// keyEndFrom returns where the remainders end, given the offset off of the remainder of the
// key at slot pos (the first slot of a key, or n).
func (la *lay) keyEndFrom(m []byte, pos, off int) int {
	sum, further := 0, 0
	for _, rl := range m[la.kl+pos : la.kl+la.n] {
		sum += int(rl)
		further += (int(rl) + 1) >> 8 // 1 for Further
	}
	return off + sum - Further*further
}

// locate returns the position (slot) of the first value of the key whose remainder is r (after the key
// part), or of the key that would follow it, with the offset of its remainder in the object (the end of the
// key area, if it goes at the end). Many-key form only; kl is the offset of the key lengths, n the number
// of slots, rem the offset of the remainders.
func locate(m []byte, kl, n, rem int, r []byte) (pos, off int, found bool) {
	off = rem
	pos = n
	for i, rl := range m[kl : kl+n] {
		if rl == Further {
			continue
		}
		c := compare(m[off:off+int(rl)], r)
		if c >= 0 {
			pos, found = i, c == 0
			break
		}
		off += int(rl)
	}
	return pos, off, found
}

func (la *lay) locate(m []byte, r []byte) (pos, off int, found bool) {
	return locate(m, la.kl, la.n, la.rem, r)
}

// runEnd returns the slot after the last value of the key whose first slot is pos.
func (la *lay) runEnd(m []byte, pos int) int {
	i := pos + 1
	for i < la.n && m[la.kl+i] == Further {
		i++
	}
	return i
}

// compare is bytes.Compare for the short remainders of a page: most differ in the first bytes, and the
// call of the library's routine costs more than the compare.
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

func sum(b []byte) int {
	s := 0
	for _, x := range b {
		s += int(x)
	}
	return s
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
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

// ShrinkLimit returns how much of the room of a smaller class the content of a page may fill for
// the page to move into that class when it loses content: 171/256, two thirds. (The one threshold of
// all pages of the redesign.)
func ShrinkLimit(avail int) int { return (avail*171 + 128) >> 8 }

// shrinkClass returns the smallest class whose size holds used bytes within ShrinkLimit, or -1.
func shrinkClass(used int) int {
	for c, size := range sizes {
		if used <= ShrinkLimit(size) {
			return c
		}
	}
	return -1
}

// allocRaw returns a zeroed object of class c that holds no pointer: an array of words, which the
// garbage collector never scans.
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

// setHead sets the head of a new page of class c: many-key form or one-key form, n slots, a key part
// of l bytes, and raw rawWords.
func (p *head) setHead(many bool, c, n, l, raw int) {
	base := TypeOneKey
	if many {
		base = TypeManyKeys
	}
	p.objType, p.n, p.raw = base+uint8(c)<<1, uint8(n), uint8(raw)
	p.setCpl(l)
}

// insertSlot makes room for one more slot at position pos in the key area of the object m, whose key area ends
// at e, and writes its entries: kl in the key lengths (many-key form), vlb in the value lengths (Str), and
// rest, the remainder of a new key, at offset ro of the remainders (nothing for a further value). The values
// area is the caller's, and so is the count n. The lists and the remainders are one block, so the move is three
// copies at most: the tail of the remainders, the tail of the value lengths with the head of the remainders,
// and the tail of the key lengths with the head of the value lengths.
func insertSlot(m []byte, la *lay, e, pos, ro int, kl, vlb byte, rest []byte) {
	dk, dv := b2i(la.many), b2i(la.str)
	d := dk + dv
	rlen := len(rest)
	if la.many {
		copy(m[la.rem+ro+d+rlen:e+d+rlen], m[la.rem+ro:e])
	}
	switch {
	case la.str:
		copy(m[la.vl+pos+d:la.rem+ro+d], m[la.vl+pos:la.rem+ro])
		if la.many {
			copy(m[la.kl+pos+dk:la.vl+pos+dk], m[la.kl+pos:la.vl+pos])
		}
	case la.many:
		copy(m[la.kl+pos+dk:la.rem+ro+dk], m[la.kl+pos:la.rem+ro])
	}
	if la.many {
		m[la.kl+pos] = kl
	}
	if la.str {
		m[la.vl+dk+pos] = vlb
	}
	copy(m[la.rem+d+ro:], rest)
}

// removeSlot takes slot `slot` out of the key area of the object m, whose key area ends at e: its key length and
// value length, and remLen bytes of remainder at offset ro of the remainders (none if the key keeps other
// values). The values area and the count are the caller's. Three copies at most, as insertSlot.
func removeSlot(m []byte, la *lay, e, slot, ro, remLen int) {
	dk, dv := b2i(la.many), b2i(la.str)
	d := dk + dv
	switch {
	case la.str:
		if la.many {
			copy(m[la.kl+slot:la.vl+slot-dk], m[la.kl+slot+dk:la.vl+slot])
		}
		copy(m[la.vl-dk+slot:la.rem+ro-d], m[la.vl+slot+1:la.rem+ro])
	case la.many:
		copy(m[la.kl+slot:la.rem+ro-dk], m[la.kl+slot+1:la.rem+ro])
	}
	if la.many {
		copy(m[la.rem-d+ro:e-d-remLen], m[la.rem+ro+remLen:e])
	}
	clear(m[e-d-remLen : e])
}

// toOneKey turns the many-key form of the object m with one key into the one-key form in place: the
// remainder joins the key part, the key lengths go. The page has la.n slots, the key part la.l bytes and the
// key area ends at e. It returns the new end of the key area. (The new key part fits nine bits: the page is
// at most 512 bytes.)
func toOneKey(p *head, m []byte, la lay, e int) int {
	rl := e - la.rem
	var tmp [MaxRemainder]byte
	copy(tmp[:], m[la.rem:e])
	nl := la.l + rl
	nvl := Header + nl
	if la.str {
		copy(m[nvl:nvl+la.n], m[la.vl:la.vl+la.n])
	}
	copy(m[Header+la.l:], tmp[:rl])
	ne := nvl + b2i(la.str)*la.n
	if ne < e {
		clear(m[ne:e])
	}
	p.setHead(false, p.class(), la.n, nl, int(p.raw))
	return ne
}
