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
//	header 8 B | directory [cap]uint8 or [cap]uint32 | heads [cap]uint64 | vals [cap]uint64 | free | heap
//
// (the arrays start at a multiple of 8). Heads past count hold pad, so that the
// ordered search, which inserts, deletes and scans use, can look at every slot
// without a branch. The directory is read as words, so the page assumes a
// little endian machine.
package vpage

import (
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
	kind  uint8 // reserved: the tree's kind byte
	class uint8 // index into sizes
	count uint8 // keys
	cap   uint8 // slots in the arrays
	ulen  uint8 // length of every suffix in the uniform flavor, else 0
	_     uint8
	top   uint16 // start of the heap
}

// Size returns the size of the page's object in bytes.
func (p *Page) Size() int { return sizes[p.class] }

// Len returns the number of keys.
func (p *Page) Len() int { return int(p.count) }

// Class returns the index of the page's class: 0 for 128 bytes, up to 3 for 1024.
func (p *Page) Class() int { return int(p.class) }

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

func (p *Page) at(off int) unsafe.Pointer { return unsafe.Add(unsafe.Pointer(p), off) }

// mem returns the whole object.
func (p *Page) mem() []byte { return unsafe.Slice((*byte)(unsafe.Pointer(p)), sizes[p.class]) }

// tags returns the directory of a uniform page: the tag of each key.
func (p *Page) tags() []uint8 { return unsafe.Slice((*uint8)(p.at(hdr)), p.cap) }

// fat returns the directory of a general page: for each key its tag, in the
// low byte, the length of its suffix, and the offset of its tail.
func (p *Page) fat() []uint32 { return unsafe.Slice((*uint32)(p.at(hdr)), p.cap) }

// base is where the arrays of a page of capacity c start: after the header and
// the directory, at a multiple of 8.
func base(c int, uniform bool) int {
	if uniform {
		return (hdr + c + 7) &^ 7
	}
	return (hdr + 4*c + 7) &^ 7
}

// heads and vals return the arrays, as long as the capacity.
func (p *Page) heads() []uint64 {
	return unsafe.Slice((*uint64)(p.at(base(int(p.cap), p.ulen != 0))), p.cap)
}
func (p *Page) vals() []uint64 {
	return unsafe.Slice((*uint64)(p.at(base(int(p.cap), p.ulen != 0)+8*int(p.cap))), p.cap)
}

// arraysEnd is the end of the arrays of a page of capacity c.
func arraysEnd(c int, uniform bool) int { return base(c, uniform) + 16*c }

// entry makes the directory entry of a general page.
func entry(t uint8, length int, off int) uint32 {
	return uint32(t) | uint32(length)<<8 | uint32(off)<<16
}

// heapFree is the room between the arrays and the heap.
func (p *Page) heapFree() int { return int(p.top) - arraysEnd(int(p.cap), p.ulen != 0) }

// word returns the head word of suffix s: its first 8 bytes, zero padded.
func word(s []byte) uint64 {
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
// position where it would go. In a uniform page a suffix of another length is
// never there, but has a position, between the suffixes it shares a head with.
func (p *Page) find(s []byte, w uint64) (int, bool) {
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

// Get returns the value of suffix s.
func (p *Page) Get(s []byte) (uint64, bool) {
	if len(s) > maxSuffix {
		return 0, false
	}
	if i, ok := p.lookup(s, word(s)); ok {
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
	n, t := int(p.count), tag(w)
	if p.ulen != 0 {
		if len(s) != int(p.ulen) {
			return 0, false
		}
		pat := uint64(t) * ones
		for off := 0; off < n; off += 8 {
			x := *(*uint64)(p.at(hdr + off)) ^ pat
			for m := (x - ones) & ^x & highs; m != 0; m &= m - 1 {
				if i := off + bits.TrailingZeros64(m)>>3; i < n && p.heads()[i] == w {
					return i, true
				}
			}
		}
		return 0, false
	}
	for off := 0; off < 4*n; off += 8 {
		x := *(*uint64)(p.at(hdr + off))
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

// Key returns the suffix at position i; buf backs it.
func (p *Page) Key(i int, buf *[maxSuffix]byte) []byte {
	l := p.length(i)
	w := p.heads()[i]
	for j := range min(l, headLen) {
		buf[j] = byte(w >> (56 - 8*j))
	}
	if l > headLen {
		copy(buf[headLen:], p.tail(i))
	}
	return buf[:l]
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
