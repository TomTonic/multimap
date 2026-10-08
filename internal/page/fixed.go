package page

import (
	"bytes"
	"unsafe"
)

// Fixed is the page of values of one size, see the package comment: for the maps whose values are small
// and hold no pointer (a uint64), or are a word with a pointer (a *X). The values are an array of T at the
// end of the object, one for each slot in the order of the slots, and compared as T (a == b), not as
// bytes. Between the key area and the values the bytes are zero.
//
// A page of a T with a pointer is a typed object (allocPtr): the first rawWords words are
// the byte area (head, key part, key lengths, remainders: no pointer in them), the rest are slots that hold
// a value or nil. The type of an object cannot change, so the byte area of such a page is fixed for
// its life: a key that does not fit it makes a new object, a value that has a free slot does not.
// Values are moved as T (never as bytes: a byte move of pointers hides them from the write barrier);
// for a T without a pointer the same code is a plain memmove.
//
// The methods that may make a new object (BuildFixedOf, Add, Remove, Widen, Prepend) take ptr, whether
// T holds a pointer (HoldsPointers[T]()): it is a reflection of the type, about 3 ns, too slow to ask
// on every change, so the map that owns the pages asks once and passes it.
//
// The methods are generic in T and do not carry it: the tree holds untyped *Fixed pointers, and the map that
// owns the page names T in every call. T must be the type the page was made with.
type Fixed struct{ head }

func size[T comparable]() int {
	var z T
	return int(unsafe.Sizeof(z))
}

// valuesIn returns the n values at the end of the object m as a slice, to be moved with copy (a move as
// T keeps the write barrier for a T with a pointer).
func valuesIn[T comparable](m []byte, n int) []T {
	return unsafe.Slice((*T)(unsafe.Pointer(&m[len(m)-n*size[T]()])), n)
}

// newFixed returns a zeroed object of class c, with its head set for n slots and a key part of l
// bytes; for a T with a pointer (ptr) the typed object whose byte area is the first words that hold e bytes.
func newFixed[T comparable](many bool, c, n, l, e int, ptr bool) *Fixed {
	var p *Fixed
	raw := 0
	if ptr {
		raw = (e + 7) / 8
		p = (*Fixed)(allocPtr[T](c, raw))
	} else {
		p = (*Fixed)(allocRaw(c))
	}
	p.setHead(many, c, n, l, raw)
	return p
}

// BuildFixed returns a page for the values vals of the keys rests (see BuildStrings), or nil if T has no
// page (see Supported) or they do not fit: no value, more than MaxEntries, a remainder beyond MaxRemainder
// (many keys) or content beyond the largest class.
func BuildFixed[T comparable](rests [][]byte, vals []T) *Fixed {
	if !Supported[T]() {
		return nil
	}
	return BuildFixedOf(rests, vals, HoldsPointers[T]())
}

// BuildFixedOf is BuildFixed for a T that is known to be Supported, and ptr says whether it holds a pointer:
// the checks are reflection of the type, too slow for a tree that builds a page for every burst.
func BuildFixedOf[T comparable](rests [][]byte, vals []T, ptr bool) *Fixed {
	n := len(rests)
	if n == 0 || n > MaxEntries || len(vals) != n {
		return nil
	}
	one := string(rests[0]) == string(rests[n-1])
	l := min(lcp(rests[0], rests[n-1]), MaxKeyPart)
	e := Header
	if one {
		l = len(rests[0])
		e += l
	} else {
		e += l + 2*n // key lengths and fingerprints
		for i, r := range rests {
			if !further(rests, i) {
				if len(r)-l > MaxRemainder {
					return nil
				}
				e += len(r) - l
			}
		}
	}
	c := classFor(e + n*size[T]())
	if c < 0 {
		return nil
	}
	p := newFixed[T](!one, c, n, l, e, ptr)
	m := p.mem()
	copy(m[Header:], rests[0][:l])
	la := p.lay(false)
	off := la.rem
	nv := valuesIn[T](m, n)
	for i, r := range rests {
		if !one {
			if further(rests, i) {
				m[la.kl+i] = Further
			} else {
				m[la.kl+i] = uint8(len(r) - l)
				m[la.fp+i] = Fingerprint(r[l:])
				off += copy(m[off:], r[l:])
			}
		}
		nv[i] = vals[i]
	}
	return p
}

// NewFixed returns the one-key page of the key part rest (the key from the end of the path on) with the one value v,
// or nil if they do not fit the largest class (T must be Supported); ptr is HoldsPointers[T](). The tree calls it
// for a key that no page holds yet. It is BuildFixedOf of one entry without the slices.
func NewFixed[T comparable](rest []byte, v T, ptr bool) *Fixed {
	e := Header + len(rest)
	c := classFor(e + size[T]())
	if c < 0 {
		return nil
	}
	p := newFixed[T](false, c, 1, len(rest), e, ptr)
	m := p.mem()
	copy(m[Header:], rest)
	valuesIn[T](m, 1)[0] = v
	return p
}

// Used returns the bytes of the page that hold keys: where the remainders end. The values are at the end
// of the object.
func (p *Fixed) Used() int {
	la := p.lay(false)
	return la.keyEnd(p.mem())
}

// Get returns the first value of the key rest (the key from the end of the path on), or false if the page
// does not hold it.
func (p *Fixed) Get[T comparable](rest []byte) (T, bool) {
	var zero T
	m := p.mem()
	n, l := int(p.n), p.cpl()
	pos := 0
	if p.one() {
		if string(m[Header:Header+l]) != string(rest) {
			return zero, false
		}
	} else {
		if lcp(m[Header:Header+l], rest) < l {
			return zero, false
		}
		var found bool
		if pos, _, found = locate(m, Header+l, n, Header+l+2*n, rest[l:]); !found {
			return zero, false
		}
	}
	return valuesIn[T](m, n)[pos], true
}

// EachValue calls fn with every value of the key rest, in the order they came in, until fn returns
// false, and reports whether the page holds the key.
func (p *Fixed) EachValue[T comparable](rest []byte, fn func(v T) bool) bool {
	m := p.mem()
	n, l := int(p.n), p.cpl()
	pos, end := 0, n
	if p.one() {
		if string(m[Header:Header+l]) != string(rest) {
			return false
		}
	} else {
		if lcp(m[Header:Header+l], rest) < l {
			return false
		}
		var found bool
		if pos, _, found = locate(m, Header+l, n, Header+l+2*n, rest[l:]); !found {
			return false
		}
		for end = pos + 1; end < n && m[Header+l+end] == Further; end++ {
		}
	}
	vs := valuesIn[T](m, n)
	for i := pos; i < end; i++ {
		if !fn(vs[i]) {
			return true
		}
	}
	return true
}

func indexOf[T comparable](vs []T, pos, end int, v T) int {
	for i := pos; i < end; i++ {
		if vs[i] == v {
			return i
		}
	}
	return -1
}

// fits reports whether a page of p's object can take a key area that ends at ne and n+d values in
// place: for a page of pointers the byte area (rawWords) and the slots must hold them, for the others the
// object.
func (p *Fixed) fits(ne, n, w int) bool {
	if p.raw != 0 {
		return ne <= int(p.raw)*8 && int(p.raw)*8+n*8 <= p.Size()
	}
	return ne+n*w <= p.Size()
}

// regrow returns a new page of class c with the key area (the first e bytes, the new end ne) and the values
// of p, for a key area that ends at ne and n slots.
func regrow[T comparable](p *Fixed, c, e, ne int, ptr bool) *Fixed {
	q := newFixed[T](!p.one(), c, int(p.n), p.cpl(), ne, ptr)
	pm, qm := p.mem(), q.mem()
	copy(qm[Header:e], pm[Header:e])
	copy(valuesIn[T](qm, int(p.n)), valuesIn[T](pm, int(p.n)))
	return q
}

// Add adds the value v to the key rest and returns the page that holds the result, which is p itself
// unless the content no longer fits p's object or the page of one key meets another key, and says what
// happened (see Result). A key that is there gets the value behind its others.
//
// ptr is HoldsPointers[T]() (see Fixed).
func (p *Fixed) Add[T comparable](rest []byte, v T, ptr bool) (*Fixed, Result) {
	m := p.mem()
	la := p.lay(false)
	w := size[T]()
	var slot, ro int
	var kl, fp byte
	var r []byte
	res := AddedValue
	e := la.rem
	vs := valuesIn[T](m, la.n)
	switch {
	case !la.many:
		if !bytes.Equal(m[Header:Header+la.l], rest) {
			return pairWithFixed(p, rest, v, ptr)
		}
		if indexOf(vs, 0, la.n, v) >= 0 {
			return p, Present
		}
		slot = la.n
	default:
		if p.Match(rest) < la.l {
			return p, Outside
		}
		r = rest[la.l:]
		pos, off, found := locate(m, la.kl, la.n, la.rem, r)
		ro, slot = off-la.rem, pos
		if found {
			end := la.runEnd(m, pos)
			if indexOf(vs, pos, end, v) >= 0 {
				return p, Present
			}
			slot, kl, r = end, Further, nil
		} else {
			if len(r) > MaxRemainder {
				return p, Full
			}
			kl, fp, res = uint8(len(r)), Fingerprint(r), Added
		}
		e = la.keyEndFrom(m, pos, off)
	}
	if la.n >= MaxEntries {
		return p, Full
	}
	ne := e + 2*b2i(la.many) + len(r)
	q := p
	if !p.fits(ne, la.n+1, w) {
		c := classFor(ne + (la.n+1)*w)
		if c < 0 {
			return p, Full
		}
		q = regrow[T](p, c, e, ne, ptr)
		m = q.mem()
	}
	// the array grows at its front: the values before the new one move down by one slot
	nv := valuesIn[T](m, la.n+1)
	copy(nv[:slot], nv[1:slot+1])
	nv[slot] = v
	if la.many { // insertSlot of the page without value lengths, written out: the key lengths and the head of the remainders move together
		rlen := len(r)
		copy(m[la.rem+ro+2+rlen:e+2+rlen], m[la.rem+ro:e])
		copy(m[la.fp+slot+2:la.rem+ro+2], m[la.fp+slot:la.rem+ro])
		copy(m[la.kl+slot+1:la.fp+slot+1], m[la.kl+slot:la.fp+slot])
		m[la.kl+slot] = kl
		m[la.fp+1+slot] = fp
		copy(m[la.rem+2+ro:], r)
	}
	q.n++
	return q, res
}

// pairWithFixed makes the page of the two keys that a one-key page and a new key rest make, in the many-key
// form. It returns the page, or p and Full if they do not fit.
func pairWithFixed[T comparable](p *Fixed, rest []byte, v T, ptr bool) (*Fixed, Result) {
	m := p.mem()
	la := p.lay(false)
	key := m[Header : Header+la.l]
	var rb [9][]byte
	var vb [9]T
	rests, vals := rb[:0], vb[:0]
	first := bytes.Compare(rest, key) < 0
	if first {
		rests, vals = append(rests, rest), append(vals, v)
	}
	for _, x := range valuesIn[T](m, la.n) {
		rests, vals = append(rests, key), append(vals, x)
	}
	if !first {
		rests, vals = append(rests, rest), append(vals, v)
	}
	if q := BuildFixedOf(rests, vals, ptr); q != nil {
		return q, Added
	}
	return p, Full
}

// Remove removes the value v of the key rest and says whether it was there (Removed) and whether the key
// went with it (Gone, its last value). It returns the page that holds the rest: p itself, a page of a
// smaller class once the content fills at most ShrinkLimit of it, or nil if the value was the only one
// (the page is gone). A many-key page that is left with one key becomes a one-key page in place, if the
// key fits the key part.
//
// ptr is HoldsPointers[T]() (see Fixed).
func (p *Fixed) Remove[T comparable](rest []byte, v T, ptr bool) (*Fixed, Removal) {
	m := p.mem()
	la := p.lay(false)
	w := size[T]()
	var slot, ro, remLen int
	res := Removed
	e := la.rem
	vs := valuesIn[T](m, la.n)
	if !la.many {
		if !bytes.Equal(m[Header:Header+la.l], rest) {
			return p, Absent
		}
		if slot = indexOf(vs, 0, la.n, v); slot < 0 {
			return p, Absent
		}
		if la.n == 1 {
			return nil, Gone
		}
	} else {
		if p.Match(rest) < la.l {
			return p, Absent
		}
		r := rest[la.l:]
		pos, off, found := locate(m, la.kl, la.n, la.rem, r)
		if !found {
			return p, Absent
		}
		end := la.runEnd(m, pos)
		if slot = indexOf(vs, pos, end, v); slot < 0 {
			return p, Absent
		}
		e = la.keyEndFrom(m, pos, off)
		if slot == pos {
			if end > pos+1 { // the next value takes the key's place: its slot becomes the key's
				m[la.kl+slot+1], m[la.fp+slot+1] = uint8(len(r)), m[la.fp+slot]
			} else { // (a many-key page has two keys at least)
				remLen, res = len(r), Gone
				ro = off - la.rem
			}
		}
	}
	// the array shrinks at its front: the values before the removed one move up by one slot
	var zero T
	copy(vs[1:slot+1], vs[:slot])
	vs[0] = zero
	if la.many { // removeSlot of the page without value lengths, written out
		copy(m[la.kl+slot:la.fp+slot-1], m[la.kl+slot+1:la.fp+slot])
		copy(m[la.fp-1+slot:la.rem+ro-2], m[la.fp+slot+1:la.rem+ro])
		copy(m[la.rem-2+ro:e-2-remLen], m[la.rem+ro+remLen:e])
		clear(m[e-2-remLen : e])
	}
	p.n--
	e -= 2*b2i(la.many) + remLen
	if la.many && res == Gone && oneKeyLeft(m, la.kl, int(p.n)) {
		e = toOneKey(&p.head, m, p.lay(false), e)
	}
	if c := shrinkClass(p.class(), e+int(p.n)*w); c < p.class() {
		return regrow[T](p, c, e, e, ptr), res
	}
	return p, res
}

// Each calls fn with the remainder (after the key part), the value and whether it is the first value of its
// key, for every value in key order until fn returns false, and reports whether it ran to completion.
// The remainders alias the page; that of a further value is its key's, and that of the one-key form is
// empty.
func (p *Fixed) Each[T comparable](fn func(rem []byte, v T, first bool) bool) bool {
	m := p.mem()
	n, l := int(p.n), p.cpl()
	vs := valuesIn[T](m, n)
	if p.one() {
		for i, v := range vs {
			if !fn(nil, v, i == 0) {
				return false
			}
		}
		return true
	}
	kl := Header + l
	off := kl + 2*n
	var rem []byte
	for i, rl := range m[kl : kl+n] {
		first := rl != Further
		if first {
			rem = m[off : off+int(rl)]
			off += int(rl)
		}
		if !fn(rem, vs[i], first) {
			return false
		}
	}
	return true
}

// EachSingle calls fn with every value of a page of the one-key form, in slot order, until fn returns false, and
// reports whether it ran to completion. It is Each without the remainders for the tree's scans of the single-key
// pages, which would otherwise call a closure around a closure for every value.
func (p *Fixed) EachSingle[T comparable](fn func(v T) bool) bool {
	for _, v := range valuesIn[T](p.mem(), int(p.n)) {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Skip drops the first k bytes of the key part (k at most its length): the tree has put a byte node above
// the page that consumes them. It happens in place.
func (p *Fixed) Skip(k int) {
	m := p.mem()
	e := p.Used()
	copy(m[Header:], m[Header+k:e])
	clear(m[e-k : e])
	p.setCpl(p.cpl() - k)
}

// Prepend returns the page with pre in front of the key part: the page moves up in the tree, when the node
// above it goes away. It is p itself if the content still fits, else a new page; nil if the key part would
// exceed the largest class.
//
// ptr is HoldsPointers[T]() (see Fixed).
func (p *Fixed) Prepend[T comparable](pre []byte, ptr bool) *Fixed {
	m := p.mem()
	la := p.lay(false)
	w := size[T]()
	e := la.keyEnd(m)
	ne := e + len(pre)
	if classFor(ne+la.n*w) < 0 {
		return nil
	}
	q := p
	if !p.fits(ne, la.n, w) {
		q = regrow[T](p, classFor(ne+la.n*w), e, ne, ptr)
		m = q.mem()
	}
	copy(m[Header+len(pre):ne], m[Header:e])
	copy(m[Header:], pre)
	q.setCpl(la.l + len(pre))
	return q
}

// Widen adds the entry rest -> v, a key that leaves the common prefix of the many-key page after mis =
// Match(rest) < PrefixLen() bytes, and shortens the common prefix to those mis bytes: the page is built
// again, with the rest of the old prefix in front of every remainder. It returns the new page, or nil if the
// entry does not go in (the content would exceed the largest class, or a remainder is too long); p is then
// unchanged.
//
// ptr is HoldsPointers[T]() (see Fixed).
func (p *Fixed) Widen[T comparable](rest []byte, v T, ptr bool) *Fixed {
	m := p.mem()
	la := p.lay(false)
	mis := p.Match(rest)
	d := la.l - mis
	newRem := rest[mis:]
	if len(newRem) > MaxRemainder {
		return nil
	}
	w := size[T]()
	e := la.keyEnd(m)
	heads, longest := 0, 0
	for _, rl := range m[la.kl : la.kl+la.n] {
		if rl != Further {
			heads++
			longest = max(longest, int(rl))
		}
	}
	if longest+d > MaxRemainder {
		return nil
	}
	nk := Header + mis + 2*(la.n+1) + (e - la.rem) + heads*d + len(newRem)
	c := classFor(nk + (la.n+1)*w)
	if c < 0 {
		return nil
	}
	first := len(newRem) == 0 || newRem[0] < p.CP()[mis]
	at := 0
	if !first {
		at = la.n
	}
	q := newFixed[T](true, c, la.n+1, mis, nk, ptr)
	out := q.mem()
	copy(out[Header:], m[Header:Header+mis])
	nl := q.lay(false)
	for i := range la.n {
		rl := m[la.kl+i]
		if rl != Further {
			rl += uint8(d)
		}
		out[nl.kl+i+b2i(i >= at)] = rl
	}
	out[nl.kl+at] = uint8(len(newRem))
	out[nl.fp+at] = Fingerprint(newRem)
	off := nl.rem
	if first {
		off += copy(out[off:], newRem)
	}
	extra := p.CP()[mis:]
	src := la.rem
	for i := range la.n {
		if rl := int(m[la.kl+i]); rl != Further {
			start := off
			off += copy(out[off:], extra)
			off += copy(out[off:], m[src:src+rl])
			src += rl
			out[nl.fp+i+b2i(i >= at)] = Fingerprint(out[start:off]) // the remainder grew: its fingerprint is new
		}
	}
	if !first {
		copy(out[off:], newRem)
	}
	vals := valuesIn[T](out, la.n+1)
	copy(vals[b2i(first):], valuesIn[T](m, la.n))
	vals[at] = v
	return q
}
