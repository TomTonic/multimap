package vpage

import "math/bits"

// This file is what the tree (internal/art) needs of a page beyond the
// operations of the prototype: positions of entries, building a page from
// sorted keys, the bytes of keys at an offset, and moving a page to a shallower
// base.

// BuildFill is how full, in percent of its class, a page built from sorted keys
// may be (see Build).
var BuildFill = 100

// Locate returns the position of suffix s, and whether it is there, or the
// position where it would go, for a suffix of any length. It is the ordered search, used where the position
// among other keys matters; Find is the faster test for a key's presence.
func (p *Page) Locate(s []byte) (int, bool) { return p.find(s) }

// LocateIn is Locate for the part of key from the page's base on.
func (p *Page) LocateIn(key []byte) (int, bool) { return p.findAt(key, int(p.base)) }

// Find returns the position of suffix s if it is in the page. It compares the
// tags of the directory instead of searching the heads. A page without a prefix,
// the common case for integers, goes the short way. For a prefix of up to
// maxFast bytes Find takes the head word of the stripped suffix and compares the
// prefix by shifting and masking words, without a branch that depends on the
// length of the page's prefix: such a branch, if the next page makes it go the
// other way, would flush the lookups the CPU has started on other keys while
// this page was loading.
func (p *Page) Find(s []byte) (int, bool) { return p.findTag(s, 0) }

// FindIn is Find for the part of key from the page's base on.
func (p *Page) FindIn(key []byte) (int, bool) { return p.findTag(key, int(p.base)) }

// findTag is Find for the suffix key[off:].
func (p *Page) findTag(key []byte, off int) (int, bool) {
	plen := int(p.plen)
	if len(key)-off > maxSuffix || len(key)-off < plen {
		return 0, false
	}
	var w uint64
	switch {
	case plen == 0:
		if p.ulen != 0 {
			return p.findUniform(key, off)
		}
		w = headWord(key, off)
	case plen <= headLen: // one word covers the prefix
		pw := bswap(*(*uint64)(p.at(hdr)))
		if (headWord(key, off)^pw)&(^uint64(0)<<(64-8*uint(plen))) != 0 {
			return 0, false
		}
		w = headWord(key, off+plen)
	case plen > maxFast:
		if string(key[off:off+plen]) != string(p.prefix()) {
			return 0, false
		}
		w = headWord(key, off+plen)
	default:
		var diff uint64
		for j := range 3 { // the prefix, a word at a time, the bytes past it masked off
			pw := bswap(*(*uint64)(p.at(hdr + headLen*j)))
			valid := uint(min(max(plen-headLen*j, 0), headLen))
			diff |= (headWord(key, off+headLen*j) ^ pw) & (^uint64(0) << (64 - 8*valid))
		}
		if diff != 0 {
			return 0, false
		}
		w = headWord(key, off+plen)
	}
	return p.lookup(key[off+plen:], w)
}

// findUniform is findTag for a page of equal suffixes without a prefix, which is
// the page of integer keys: it compares the key's tag with those of the page 8
// at a time and then its head with that of the entries that have the tag, with
// no slices and no calls in between.
func (p *Page) findUniform(key []byte, off int) (int, bool) {
	if len(key)-off != int(p.ulen) {
		return 0, false
	}
	w := headWord(key, off)
	n, pat := int(p.count), uint64(tag(w))*ones
	heads := (hdr + int(p.cap) + 7) &^ 7
	for o := 0; o < n; o += 8 {
		x := *(*uint64)(p.at(hdr + o)) ^ pat
		for m := (x - ones) & ^x & highs; m != 0; m &= m - 1 {
			if i := o + bits.TrailingZeros64(m)>>3; i < n && *(*uint64)(p.at(heads + 8*i)) == w {
				return i, true
			}
		}
	}
	return 0, false
}

// ByteAt returns byte off of the key at position i, counted from the page's
// base, or 0 if the key is shorter.
func (p *Page) ByteAt(i, off int) byte {
	if off < int(p.plen) {
		return p.prefix()[off]
	}
	off -= int(p.plen)
	switch l := p.length(i); {
	case off >= l:
		return 0
	case off < headLen:
		return byte(p.heads()[i] >> (56 - 8*off))
	}
	return p.tail(i)[off-headLen]
}

// Build returns a page of base base that holds the n keys that key returns, in
// sorted order, with the values val returns, or nil if they do not fit a page of
// the largest class with BuildFill percent of it. key must return slices that
// stay valid until Build returns; a key is the suffix from the base.
func Build(base, n int, key func(i int) []byte, val func(i int) uint64) *Page {
	if n == 0 || n > 255 {
		return nil
	}
	shape := func(q int) plan {
		pl := plan{n: n, plen: q}
		tails, uni := 0, -1
		for i := range n {
			if len(key(i)) > maxSuffix {
				return plan{n: -1}
			}
			tails, uni = see(tails, uni, len(key(i))-q)
		}
		pl.tails = tails
		if uni >= 1 && uni <= headLen {
			pl.ulen = uni
		}
		return pl
	}
	pl := shape(0)
	if pl.n < 0 {
		return nil
	}
	first := key(0)
	if q := lcp(first, key(n-1)) - PrefixSlack; n >= 2 && q >= max(MinPrefix, 1) && pl.tails >= MinGain {
		if with := shape(q); need(with)+MinGain <= need(pl) {
			pl = with
		}
	}
	c := fitClass(pl, 0, BuildFill)
	if c < 0 {
		return nil
	}
	q := newPage(c, pl, first[:pl.plen])
	q.base = uint8(base)
	for i := range n {
		q.appendFull(key(i), val(i))
	}
	return q
}

// Rebase returns a page with the same entries whose suffixes start at the
// shallower depth base: pre are the bytes of the path between base and the
// page's base, which each suffix gets in front. It returns nil if they do not fit
// a page.
func (p *Page) Rebase(base int, pre []byte) *Page {
	var buf [maxSuffix]byte
	keys := make([][]byte, p.Len())
	for i := range keys {
		keys[i] = append(append(make([]byte, 0, len(pre)+maxSuffix), pre...), p.Key(i, &buf)...)
	}
	return Build(base, len(keys), func(i int) []byte { return keys[i] }, p.Val)
}

// ValPtr returns the address of the value at position i, which the caller may
// read or set.
func (p *Page) ValPtr(i int) *uint64 { return &p.vals()[i] }

// Vals returns the values of the page's entries, in key order.
func (p *Page) Vals() []uint64 { return p.vals()[:p.count] }

// AppendKey appends the suffix at position i, with the page's prefix, to dst
// and returns the result.
func (p *Page) AppendKey(dst []byte, i int) []byte {
	dst = append(dst, p.prefix()...)
	l := p.length(i)
	w := p.heads()[i]
	for j := range min(l, headLen) {
		dst = append(dst, byte(w>>(56-8*j)))
	}
	if l > headLen {
		dst = append(dst, p.tail(i)...)
	}
	return dst
}

// Thin reports whether the page holds so little that it may fit a page with
// a neighbour (see Merge): its entries and heap take at most half of what a merge
// allows. It is a cheap test that most deletes fail, which spares looking at the
// neighbours.
func (p *Page) Thin() bool {
	used := arraysEnd(int(p.count), p.ulen != 0, int(p.plen)) + sizes[p.class()] - int(p.top)
	return 2*used*100 <= MergeFill*sizes[MaxClass]
}

// Shared returns how many bytes, counted from the page's base, all keys of the
// page start with: for a page of one key, its whole length.
func (p *Page) Shared() int {
	n := int(p.count)
	if n == 1 {
		return int(p.plen) + p.length(0)
	}
	return int(p.plen) + p.lcpEntries(0, n-1)
}
