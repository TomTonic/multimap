package art

import (
	"bytes"
	"cmp"
	"encoding/binary"
	"unsafe"

	"github.com/TomTonic/multimap/internal/swar"
)

// S pages hold keys of any length up to maxPageKey bytes, for instance
// strings, with any number of values each. The prefix that all keys of the
// page share is stored once. Of each key's rest, the suffix, the first 8 bytes
// are a word in heads (big endian, zero-padded), as in U8 pages, and the bytes
// beyond them, its tail, lie in a byte area. A search compares words and looks
// at tails only when words are equal. The values are held like in U8-n pages
// (see pagen.go), and the functions there serve S pages as well.
//
// How much room goes to keys, values and bytes is not fixed per class: when a
// page is (re)packed, it gets the smallest class its content fits, and the
// room left over is split in proportion to what its keys need on average (see
// sPack). A page that runs out of one kind of room is repacked. So a page of
// long keys holds fewer of them than a page of short ones, but both fill their
// class.
//
// Layout, after the pageHead:
//
//	ext   [E]unsafe.Pointer  external value sets (E per class, see sExt)
//	heads [kcap]uint64       the first 8 suffix bytes of each key
//	vals  [vcap]uint64       inline values, packed key by key
//	offs  [kcap]uint16       where each key's tail starts in the byte area
//	lens  [kcap]uint8        the length of each key's suffix
//	cnts  [kcap]uint8        the inline values per key, as in U8-n pages
//	bytes                    the shared prefix, then the tails in key order
//
// The external sets come first, so that the garbage collector, which scans an
// object only up to its last pointer, reads a few words of the page, not all
// of it.

// maxPageKey is the longest key a page holds; a longer key gets a leaf.
const maxPageKey = 255

// The S page classes, of 64, 128, 256 and 512 bytes.
type (
	pageS64 struct {
		pageHead
		body [7]uint64
	}
	pageS128 struct {
		pageHead
		ext  [1]unsafe.Pointer
		body [14]uint64
	}
	pageS256 struct {
		pageHead
		ext  [2]unsafe.Pointer
		body [29]uint64
	}
	pageS512 struct {
		pageHead
		ext  [4]unsafe.Pointer
		body [59]uint64
	}
)

var (
	sSizes = [4]int{64, 128, 256, 512}
	sExt   = [4]int{0, 1, 2, 4} // external slots per class
)

// sKeyBytes is what a key costs in an S page besides its values and tail:
// its head word, tail offset, suffix length and value count.
const sKeyBytes = 8 + 2 + 1 + 1

// keyBuf backs a key rebuilt from a page; the extra 8 bytes let a head word
// be stored whole at any position.
type keyBuf [maxPageKey + 8]byte

func newPageS(class int) *pageHead {
	var p *pageHead
	switch class {
	case 0:
		p = &(&pageS64{}).pageHead
	case 1:
		p = &(&pageS128{}).pageHead
	case 2:
		p = &(&pageS256{}).pageHead
	default:
		p = &(&pageS512{}).pageHead
	}
	p.kind, p.class = kPageS, uint8(class)
	return p
}

// sLayout is where an S page keeps its arrays. It follows from the page's
// class and capacities; the part that U8-n pages share is an nLayout.
type sLayout struct {
	nLayout
	offsOff, lensOff, bytesOff, end uintptr
}

func (p *pageHead) sLayout() sLayout {
	e := sExt[p.class]
	k, v := uintptr(p.kcap), uintptr(p.vcap)
	h := headsOff + 8*uintptr(e)
	vo := h + 8*k
	oo := vo + 8*v
	lo := oo + 2*k
	co := lo + k
	return sLayout{nLayout{int(k), int(v), e, h, co, headsOff, vo}, oo, lo, co + k, uintptr(sSizes[p.class])}
}

// sKeys returns the head words of the page's keys; sLens, sOffs and sBytes
// the other key arrays, over their full capacity.
func (p *pageHead) sKeys(l *sLayout) []uint64 {
	return unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), l.headsOff)), p.count)
}

func (p *pageHead) sLens(l *sLayout) []uint8 {
	return unsafe.Slice((*uint8)(unsafe.Add(unsafe.Pointer(p), l.lensOff)), l.keys)
}

func (p *pageHead) sOffs(l *sLayout) []uint16 {
	return unsafe.Slice((*uint16)(unsafe.Add(unsafe.Pointer(p), l.offsOff)), l.keys)
}

// The byte area may be empty and end with the page, so it is cut from the
// whole page: a pointer to its start would point past the page.
func (p *pageHead) sBytes(l *sLayout) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(p)), l.end)[l.bytesOff:]
}

// tailLen is the length of the tail of a suffix of n bytes.
func tailLen[N uint8 | int](n N) int { return max(int(n)-8, 0) }

// headWord returns the word a suffix is sorted and found by: its first 8
// bytes, zero-padded. Words order suffixes like bytes.Compare, except that
// suffixes that differ only beyond them, or in trailing zeros, share a word.
func headWord(s []byte) uint64 {
	if len(s) >= 8 {
		return binary.BigEndian.Uint64(s)
	}
	return keyWord(s)
}

// sUsed returns how many bytes of the byte area are in use: the tails end
// with the last key's. A page always holds a key.
func (p *pageHead) sUsed(l *sLayout) int {
	n := int(p.count)
	return int(p.sOffs(l)[n-1]) + tailLen(p.sLens(l)[n-1])
}

// sMatch reports whether key may be in the page: short enough, and starting
// with the page's shared prefix. Only then can sSearch look for its suffix.
func (p *pageHead) sMatch(key []byte) bool {
	b := int(p.base)
	if len(key) < b || len(key) > maxPageKey {
		return false
	}
	l := p.sLayout()
	return bytes.Equal(key[:b], p.sBytes(&l)[:b])
}

// sSearch returns the position of the key with suffix s in the page, or where
// it would be inserted, and whether it is there.
func (p *pageHead) sSearch(s []byte) (int, bool) {
	l := p.sLayout()
	h := p.sKeys(&l)
	w := headWord(s)
	i, ok := search(h, w)
	if !ok {
		return i, false
	}
	lens, offs, by := p.sLens(&l), p.sOffs(&l), p.sBytes(&l)
	st := s[min(len(s), 8):]
	for ; i < len(h) && h[i] == w; i++ {
		o := int(offs[i])
		c := bytes.Compare(by[o:o+tailLen(lens[i])], st)
		if c == 0 {
			c = cmp.Compare(int(lens[i]), len(s))
		}
		if c >= 0 {
			return i, c == 0
		}
	}
	return i, false
}

// sKey returns key i of the page, rebuilt in buf.
func (p *pageHead) sKey(i int, buf *keyBuf) []byte {
	l := p.sLayout()
	by := p.sBytes(&l)
	b := copy(buf[:], by[:p.base])
	binary.BigEndian.PutUint64(buf[b:], p.sKeys(&l)[i])
	n := int(p.sLens(&l)[i])
	if n > 8 {
		o := int(p.sOffs(&l)[i])
		copy(buf[b+8:], by[o:o+n-8])
	}
	return buf[:b+n]
}

// sRoom reports whether a key with a suffix of n bytes and one value fits the
// page without repacking.
func (p *pageHead) sRoom(n int) bool {
	l := p.sLayout()
	return p.count < p.kcap && p.nv < p.vcap && p.sUsed(&l)+tailLen(n) <= len(p.sBytes(&l))
}

// sInsertKey inserts the key with suffix s and the one value v at position i.
// The page must have room (see sRoom).
func (p *pageHead) sInsertKey(i int, s []byte, v uint64) {
	l := p.sLayout()
	n := int(p.count)
	lens, offs, by := p.sLens(&l), p.sOffs(&l), p.sBytes(&l)
	used, t := p.sUsed(&l), tailLen(len(s))
	start := used
	if i < n {
		start = int(offs[i])
	}
	copy(by[start+t:used+t], by[start:used])
	copy(by[start:start+t], s[min(len(s), 8):])
	copy(lens[i+1:n+1], lens[i:n])
	copy(offs[i+1:n+1], offs[i:n])
	lens[i], offs[i] = uint8(len(s)), uint16(start)
	for j := i + 1; j <= n; j++ {
		offs[j] += uint16(t)
	}
	p.nInsertKey(i, headWord(s), v)
}

// sRemoveKey removes key i with all its values and returns the page, its
// repacked replacement once it fits well into a smaller class, or nil once it
// is empty.
func (p *pageHead) sRemoveKey(i int) *pageHead {
	n := int(p.count)
	if n == 1 {
		return nil
	}
	l := p.sLayout()
	lens, offs, by := p.sLens(&l), p.sOffs(&l), p.sBytes(&l)
	used, t, start := p.sUsed(&l), tailLen(lens[i]), int(offs[i])
	copy(by[start:used-t], by[start+t:used])
	copy(lens[i:n-1], lens[i+1:n])
	copy(offs[i:n-1], offs[i+1:n])
	for j := i; j < n-1; j++ {
		offs[j] -= uint16(t)
	}
	p.nDropKey(i)
	// Shrink with a margin, so that a page that has just grown does not
	// shrink back after one removal.
	if c := int(p.class) - 1; c >= 0 && p.extUsed() <= sExt[c] &&
		sNeed(c, int(p.count), int(p.nv), used-t) <= 2*sSizes[c]/3 {
		return sPack(pageItems(p))
	}
	return p
}

// sNeed returns the bytes an S page of class c needs for k keys, v inline
// values and a byte area of b bytes in use.
func sNeed(c, k, v, b int) int {
	return int(headsOff) + 8*sExt[c] + sKeyBytes*k + 8*v + b
}

// sPack returns an S page holding items, which are sorted by key and
// distinct, in the smallest class they fit, or nil if they fit none or one of
// them is a leaf (its key is too long for a page).
func sPack(items []item) *pageHead {
	k := len(items)
	base := swar.Lcp(items[0].key, items[k-1].key)
	tails, inline, sets := 0, 0, 0
	for _, it := range items {
		if it.leaf != nil {
			return nil
		}
		tails += tailLen(len(it.key) - base)
		if it.set != nil {
			sets++
		} else {
			inline += len(it.vals)
		}
	}
	for c, size := range sSizes {
		need := sNeed(c, k, inline, base+tails)
		if sets > sExt[c] || need > size {
			continue
		}
		// The room left over goes to extra keys, each with the average
		// values and tail of the keys so far; what the values leave over
		// stays in the byte area.
		extra := (size - need) * k / (sKeyBytes*k + 8*inline + tails)
		p := newPageS(c)
		p.count, p.nv, p.base = uint8(k), uint8(inline), uint8(base)
		p.kcap, p.vcap = uint8(min(k+extra, 255)), uint8(inline+extra*inline/k)
		p.sFill(items)
		return p
	}
	return nil
}

// sFill stores items in the new page p, whose counts and capacities are set.
func (p *pageHead) sFill(items []item) {
	l := p.sLayout()
	h := unsafe.Slice((*uint64)(unsafe.Add(unsafe.Pointer(p), l.headsOff)), l.keys)
	lens, offs, by := p.sLens(&l), p.sOffs(&l), p.sBytes(&l)
	b := int(p.base)
	o := copy(by, items[0].key[:b])
	for i, it := range items {
		s := it.key[b:]
		h[i], lens[i], offs[i] = headWord(s), uint8(len(s)), uint16(o)
		o += copy(by[o:], s[min(len(s), 8):])
	}
	p.fillVals(items)
}
