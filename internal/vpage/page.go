// Package vpage is the prototype of the page of the redesign (docs/redesign,
// step 1): a sorted run of keys of any length, each with one value, in one
// object of 128 to 1024 bytes that holds no pointers. It is not part of the
// tree yet; it answers the questions of PLAN step 1 (bytes per key, lines per
// lookup, the cost of inserts and deletes) before the tree is changed.
//
// A page holds the suffixes of its keys below the place the tree routes them
// to; the tree strips the rest. It comes in two flavors, told by Page.ulen:
//
//   - uniform (ulen 1..8): all suffixes have that length. A key is a head
//     word, its suffix big endian and zero padded, and a value word: 16
//     bytes. This is the page of integer keys of branch node-pages.
//   - general (ulen 0): a suffix has any length up to 255. The head word
//     holds its first 8 bytes, zero padded; the bytes beyond them, the tail,
//     are in a heap that grows down from the end of the page. A key costs
//     head, value, and a directory entry of 4 bytes, and its tail.
//
// Both flavors start with a directory in the first bytes of the page, in the
// manner of a file allocation table: the header of 8 bytes, then an entry for
// each key, in the order of the keys. In the uniform flavor it is a tag, one
// byte, a hash of the key's head word. In the general flavor it is 4 bytes: the
// tag, the length of the suffix, and the offset of its tail in the heap. A
// lookup reads that directory, which is in the first line or two of the page,
// compares the tag of its key with the tags of all keys, and learns the position
// of the key, or that it is not there, and for a key with a tail, where the tail
// is; then it reads the head, the value and the tail, which are independent of
// each other, in one round. The arrays are packed by the capacity cap of the
// page, which a rebuild chooses so that the slots and the heap fill up at about
// the same time:
//
//	header 8 B | prefix | directory [cap]uint8 or [cap]uint32 | heads [cap]uint64 | vals [cap]uint64 | free | heap
//
// (the prefix, which all keys of the page share, and the arrays start at a multiple of 8). Heads past count hold pad, so that the
// ordered search, which inserts, deletes and scans use, can look at every slot
// without a branch. The directory is read as words, so the page assumes a
// little endian machine.
package vpage

import (
	"encoding/binary"
	"errors"
	"math/bits"
	"unsafe"
)

const (
	hdr       = 8   // size of the page header
	headLen   = 8   // bytes of a suffix that a head word holds
	maxSuffix = 255 // the longest suffix a page holds
	pad       = ^uint64(0)
)

// sizes are the sizes of the page classes in bytes: Go size classes that are
// multiples of 128 and, up to 512, aligned to their size.
var sizes = [...]int{128, 256, 512, 1024}

// MaxClass is the largest class a page grows into; a page of that class that
// is full has to be split. Experiments may lower it.
var MaxClass = 2

// ErrTooLong is returned for a suffix longer than a page can hold.
var ErrTooLong = errors.New("vpage: suffix longer than 255 bytes")

// Page is the header of a page; the page itself is the object it starts.
type Page struct {
	kind  uint8  // reserved: the tree's kind byte
	class uint8  // index into sizes
	count uint8  // keys
	cap   uint8  // slots in the arrays
	ulen  uint8  // stored length of every suffix in the uniform flavor, else 0
	plen  uint8  // length of the prefix all keys of the page share, stored once at the end
	top   uint16 // start of the heap
}

// Size returns the size of the page's object in bytes.
func (p *Page) Size() int { return sizes[p.class] }

// Len returns the number of keys.
func (p *Page) Len() int { return int(p.count) }

// Class returns the index of the page's class: 0 for 128 bytes, up to 3 for 1024.
func (p *Page) Class() int { return int(p.class) }

// PrefixLen returns the length of the prefix the page stores once.
func (p *Page) PrefixLen() int { return int(p.plen) }

// Uniform reports whether the page is of the uniform flavor.
func (p *Page) Uniform() bool { return p.ulen != 0 }

// alloc returns a zeroed object of class c.
func alloc(c int) *Page {
	switch c {
	case 0:
		return (*Page)(unsafe.Pointer(new([16]uint64)))
	case 1:
		return (*Page)(unsafe.Pointer(new([32]uint64)))
	case 2:
		return (*Page)(unsafe.Pointer(new([64]uint64)))
	}
	return (*Page)(unsafe.Pointer(new([128]uint64)))
}

// prefix returns the bytes all keys of the page start with; they follow the
// header, in the first line, so that a lookup that has read the header has read
// the prefix as well.
func (p *Page) prefix() []byte { return p.mem()[hdr : hdr+int(p.plen)] }

// dirAt is where the directory of a page with a prefix of plen bytes starts:
// after the header and the prefix, at a multiple of 8.
func dirAt(plen int) int { return hdr + (plen+7)&^7 }

// strip returns suffix s without the page's prefix and 0. A suffix that does
// not start with the prefix is not in the page: strip returns nil and -1 if it
// sorts below all keys of the page, +1 if above them.
func (p *Page) strip(s []byte) ([]byte, int) {
	n := int(p.plen)
	if n == 0 {
		return s, 0
	}
	pre := p.prefix()
	if len(s) >= n && string(s[:n]) == string(pre) {
		return s[n:], 0
	}
	if compare(s[:min(len(s), n)], pre) < 0 {
		return nil, -1
	}
	return nil, 1
}

func (p *Page) at(off int) unsafe.Pointer { return unsafe.Add(unsafe.Pointer(p), off) }

// mem returns the whole object.
func (p *Page) mem() []byte { return unsafe.Slice((*byte)(unsafe.Pointer(p)), sizes[p.class]) }

// tags returns the directory of a uniform page: the tag of each key.
func (p *Page) tags() []uint8 { return unsafe.Slice((*uint8)(p.at(dirAt(int(p.plen)))), p.cap) }

// fat returns the directory of a general page: for each key its tag, in the
// low byte, the length of its suffix, and the offset of its tail.
func (p *Page) fat() []uint32 { return unsafe.Slice((*uint32)(p.at(dirAt(int(p.plen)))), p.cap) }

// base is where the arrays of a page of capacity c and a prefix of plen bytes
// start: after the header, the prefix and the directory, at a multiple of 8.
func base(c int, uniform bool, plen int) int {
	if uniform {
		return (dirAt(plen) + c + 7) &^ 7
	}
	return (dirAt(plen) + 4*c + 7) &^ 7
}

// heads and vals return the arrays, as long as the capacity.
func (p *Page) heads() []uint64 {
	return unsafe.Slice((*uint64)(p.at(base(int(p.cap), p.ulen != 0, int(p.plen)))), p.cap)
}
func (p *Page) vals() []uint64 {
	return unsafe.Slice((*uint64)(p.at(base(int(p.cap), p.ulen != 0, int(p.plen))+8*int(p.cap))), p.cap)
}

// arraysEnd is the end of the arrays of a page of capacity c.
func arraysEnd(c int, uniform bool, plen int) int { return base(c, uniform, plen) + 16*c }

// entry makes the directory entry of a general page.
func entry(t uint8, length int, off int) uint32 {
	return uint32(t) | uint32(length)<<8 | uint32(off)<<16
}

// heapFree is the room between the arrays and the heap.
func (p *Page) heapFree() int { return int(p.top) - arraysEnd(int(p.cap), p.ulen != 0, int(p.plen)) }

// word returns the head word of suffix s: its first 8 bytes, zero padded.
func word(s []byte) uint64 {
	if len(s) >= headLen {
		return binary.BigEndian.Uint64(s)
	}
	var w uint64
	n := min(len(s), headLen)
	for i := range n {
		w |= uint64(s[i]) << (56 - 8*i)
	}
	return w
}

// tag returns the tag of head word w: its hash, a byte. A key that is not in
// a page has the tag of one of the page's n keys with a chance of n in 256, and
// is then told apart by its head word.
func tag(w uint64) uint8 { return uint8(w * 0x9E3779B97F4A7C15 >> 56) }

// below returns 1 if x < w, else 0, without a branch.
func below(x, w uint64) int {
	_, b := bits.Sub64(x, w, 0)
	return int(b)
}

// lower returns how many of the heads h are below w. The heads are sorted and
// padded with the largest word, which is below nothing. It compares one fence
// head per block of 8 first, which picks the block, then the other heads of
// that block; all these loads are independent, so the cache misses of a page
// that is not in the L1 cache overlap, and no branch depends on where w lies.
func lower(h []uint64, w uint64) int {
	n := len(h)
	if n <= 8 {
		c := 0
		for _, x := range h {
			c += below(x, w)
		}
		return c
	}
	blk := 0
	for f := 7; f < n; f += 8 {
		blk += below(h[f], w)
	}
	start := 8 * blk
	c := 0
	for _, x := range h[start:min(start+7, n)] { // a block's last head is its fence: known not to be below
		c += below(x, w)
	}
	return start + c
}

// length returns the length of the suffix at position i.
func (p *Page) length(i int) int {
	if p.ulen != 0 {
		return int(p.ulen)
	}
	return int(p.fat()[i] >> 8 & 0xff)
}

// tail returns the bytes of the suffix at position i beyond its head word.
func (p *Page) tail(i int) []byte {
	if p.ulen != 0 {
		return nil
	}
	e := p.fat()[i]
	l := int(e>>8&0xff) - headLen
	if l <= 0 {
		return nil
	}
	off := int(e >> 16)
	return p.mem()[off : off+l]
}

// cmpEntry compares the suffix at position i, whose head word equals the one
// of s, with s: negative if the entry is below s. Entries of equal head words
// are ordered by their suffixes: a short one, which the zero padding makes
// equal to a prefix, before longer ones.
func (p *Page) cmpEntry(i int, s []byte) int {
	la, lb := p.length(i), len(s)
	switch {
	case la <= headLen && lb <= headLen:
		return la - lb
	case la <= headLen:
		return -1
	case lb <= headLen:
		return 1
	}
	return compare(p.tail(i), s[headLen:])
}

func compare(a, b []byte) int {
	n := min(len(a), len(b))
	for i := range n {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
	}
	return len(a) - len(b)
}

// find returns the position of suffix s and whether it is there, or the
// position where it would go, also for a suffix that does not start with the
// page's prefix.
func (p *Page) find(s []byte) (int, bool) {
	rest, rel := p.strip(s)
	switch rel {
	case -1:
		return 0, false
	case 1:
		return int(p.count), false
	}
	return p.findRest(rest, word(rest))
}

// findRest is find for a suffix without the prefix, whose head word is w. In a
// uniform page a suffix of another length is never there, but has a position,
// between the suffixes it shares a head with.
func (p *Page) findRest(s []byte, w uint64) (int, bool) {
	h := p.heads()
	i := lower(h, w)
	n := int(p.count)
	if p.ulen != 0 && len(s) == int(p.ulen) {
		return i, i < n && h[i] == w
	}
	for ; i < n && h[i] == w; i++ {
		switch c := p.cmpEntry(i, s); {
		case c == 0:
			return i, true
		case c > 0:
			return i, false
		}
	}
	return i, false
}

// maxFast is the longest prefix Get handles without a branch on its length.
const maxFast = 3*headLen - 1

// Get returns the value of suffix s. A page without a prefix, the common case
// for integers, goes the short way. For a prefix of up to maxFast bytes Get
// takes the head word of the stripped suffix and compares the prefix by
// shifting and masking words, without a branch that depends on the length of the
// page's prefix: such a branch, if the next page makes it go the other way,
// would flush the lookups the CPU has started on other keys while this page was
// loading.
func (p *Page) Get(s []byte) (uint64, bool) {
	plen := int(p.plen)
	if len(s) > maxSuffix || len(s) < plen {
		return 0, false
	}
	rest, w := s, uint64(0)
	switch {
	case plen == 0:
		w = word(s)
	case plen > maxFast:
		var rel int
		if rest, rel = p.strip(s); rel != 0 {
			return 0, false
		}
		w = word(rest)
	default:
		var ws [4]uint64 // the first 32 bytes of s as head words
		for j := 0; j < 4 && headLen*j < len(s); j++ {
			ws[j] = word(s[headLen*j:])
		}
		var diff uint64
		for j := range 3 { // the prefix, a word at a time, the bytes past it masked off
			pw := bits.ReverseBytes64(*(*uint64)(p.at(hdr + headLen*j)))
			valid := uint(min(max(plen-headLen*j, 0), headLen))
			diff |= (ws[j] ^ pw) & (^uint64(0) << (64 - 8*valid))
		}
		if diff != 0 {
			return 0, false
		}
		k, r := plen/headLen, uint(plen%headLen)*8
		w = ws[k]<<r | ws[k+1]>>(64-r)
		rest = s[plen:]
	}
	if i, ok := p.lookup(rest, w); ok {
		return p.vals()[i], true
	}
	return 0, false
}

const (
	ones  = 0x0101010101010101
	highs = 0x8080808080808080
)

// lookup finds the position of suffix s, whose head word is w, by the
// directory. In a uniform page it compares the tag of s with 8 tags at a time (a
// byte is zero in x ^ pattern where the tags are equal, which the subtraction
// marks in the high bit, with false marks possible above a true one); in a
// general page it compares the tags of two entries per word. It then checks the
// head word, and the length and the tail, of the keys with that tag.
func (p *Page) lookup(s []byte, w uint64) (int, bool) {
	n, t, d := int(p.count), tag(w), dirAt(int(p.plen))
	if p.ulen != 0 {
		if len(s) != int(p.ulen) {
			return 0, false
		}
		pat := uint64(t) * ones
		for off := 0; off < n; off += 8 {
			x := *(*uint64)(p.at(d + off)) ^ pat
			for m := (x - ones) & ^x & highs; m != 0; m &= m - 1 {
				if i := off + bits.TrailingZeros64(m)>>3; i < n && p.heads()[i] == w {
					return i, true
				}
			}
		}
		return 0, false
	}
	for off := 0; off < 4*n; off += 8 {
		x := *(*uint64)(p.at(d + off))
		if i := off / 4; uint8(x) == t && i < n && p.sameGeneral(i, uint32(x), s, w) {
			return i, true
		} else if uint8(x>>32) == t && i+1 < n && p.sameGeneral(i+1, uint32(x>>32), s, w) {
			return i + 1, true
		}
	}
	return 0, false
}

// sameGeneral reports whether the suffix at position i of a general page, whose
// directory entry is e and whose tag is the one of s, is s.
func (p *Page) sameGeneral(i int, e uint32, s []byte, w uint64) bool {
	l := int(e >> 8 & 0xff)
	if l != len(s) || p.heads()[i] != w {
		return false
	}
	if l <= headLen {
		return true
	}
	off := int(e >> 16)
	return string(p.mem()[off:off+l-headLen]) == string(s[headLen:])
}

// Key returns the suffix at position i, with the page's prefix; buf backs it.
func (p *Page) Key(i int, buf *[maxSuffix]byte) []byte {
	pl := int(p.plen)
	copy(buf[:], p.prefix())
	l := p.length(i)
	w := p.heads()[i]
	for j := range min(l, headLen) {
		buf[pl+j] = byte(w >> (56 - 8*j))
	}
	if l > headLen {
		copy(buf[pl+headLen:], p.tail(i))
	}
	return buf[:pl+l]
}

// Val returns the value at position i.
func (p *Page) Val(i int) uint64 { return p.vals()[i] }

// tag returns the tag of the key at position i, from the directory.
func (p *Page) tag(i int) uint8 {
	if p.ulen != 0 {
		return p.tags()[i]
	}
	return uint8(p.fat()[i])
}
