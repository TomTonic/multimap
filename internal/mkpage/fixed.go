package mkpage

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/skpage"
)

// Fixed is the multi-key page of values of one size, see the package comment: for
// the maps whose values are small and hold no pointer (a uint64). The values are
// an array of T at the end of the object, one for each slot in the order of the slots,
// and compared as T (a == b), not as bytes: the array does not move when the keys grow or
// shrink in front of it, and a reader finds it without adding up the length list (it ends
// with the object, as the slots of the typed objects of a pointer page do). Between the
// remainders and the values the bytes are zero. The pointer page (a T that is a word with
// a pointer, the typed objects of skpage) is step 5.3.
//
// The methods are generic in T and do not carry it: the tree holds untyped *Fixed
// pointers, as it holds *Page, and the map that owns the page names T in every
// call. T must be the type the page was made with.
type Fixed struct{ head }

// Supported reports whether values of type T go into a Fixed page: T of 1 to 16
// bytes without a pointer.
func Supported[T comparable]() bool {
	return skpage.Supported[T]() && !skpage.HoldsPointers[T]()
}

func sizeAlign[T comparable]() (w, a int) {
	var z T
	return int(unsafe.Sizeof(z)), int(unsafe.Alignof(z))
}

// NeedFixed returns the bytes a Fixed page of T takes for n slots whose remainders
// (after a common prefix of cpl bytes) are remBytes in all (the slots with Further have none).
// The tree calls it to decide, before it builds anything, whether the entries of a subtree
// fit a page (at most 512).
func NeedFixed[T comparable](n, cpl, remBytes int) int {
	w, _ := sizeAlign[T]()
	return Header + n + cpl + remBytes + n*w
}

// vs returns where the values start: the array ends with the object. The object is a
// multiple of 8 bytes and the array a multiple of the size of T, so T is aligned.
func (p *Fixed) vs(w int) int { return p.Size() - int(p.n)*w }

// BuildFixed returns a page for the values vals of the keys rests (the keys from the
// end of the path on, in key order: a key with several values appears once for each, one
// after the other), or nil if T has no page (see Supported) or they do not fit: no value, more
// than MaxEntries, a remainder beyond MaxRemainder or content beyond the largest class. The
// common prefix of the page is the longest one all keys share (up to MaxPrefix).
func BuildFixed[T comparable](rests [][]byte, vals []T) *Fixed {
	if !Supported[T]() {
		return nil
	}
	return BuildFixedOf(rests, vals)
}

// BuildFixedOf is BuildFixed for a T that is known to be Supported: the check is a reflection
// of the type, too slow for a tree that builds a page for every burst.
func BuildFixedOf[T comparable](rests [][]byte, vals []T) *Fixed {
	n := len(rests)
	if n == 0 || n > MaxEntries || len(vals) != n {
		return nil
	}
	cp := prefixOf(rests)
	keyEnd := Header + n + cp
	for i, r := range rests {
		if further(rests, i) {
			continue
		}
		if len(r)-cp > MaxRemainder {
			return nil
		}
		keyEnd += len(r) - cp
	}
	w, _ := sizeAlign[T]()
	c := classFor(keyEnd + n*w)
	if c < 0 {
		return nil
	}
	p := (*Fixed)(allocRaw(c))
	newHead(&p.head, c, n, cp)
	m := p.mem()
	vs := p.vs(w)
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
		*(*T)(unsafe.Pointer(&m[vs+i*w])) = vals[i]
	}
	return p
}

// locate returns the position (slot) of the first value of the key with the remainder r, or
// of the key that would follow it, with the offset of its remainder in the object.
func (p *Fixed) locate(r []byte) (pos, off int, found bool) {
	m := p.mem()
	n := int(p.n)
	lo := Header + p.cpl()
	off = lo + n
	pos = n
	for i, rl := range m[lo : lo+n] {
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

// keyEndFrom returns the offset where the remainders end, given the offset off of the
// remainder of the key at slot pos (the first slot of a key, or n).
func (p *Fixed) keyEndFrom(pos, off int) int {
	m := p.mem()
	lo := Header + p.cpl()
	sum, further := 0, 0
	for _, rl := range m[lo+pos : lo+int(p.n)] {
		sum += int(rl)
		further += (int(rl) + 1) >> 8 // 1 for Further
	}
	return off + sum - Further*further
}

func valueAt[T comparable](m []byte, vs, i int) *T {
	w, _ := sizeAlign[T]()
	return (*T)(unsafe.Pointer(&m[vs+i*w]))
}

// Used returns the bytes of the page that hold keys: where the remainders end. The values
// are at the end of the object.
func (p *Fixed) Used() int { return p.keyEndFrom(0, Header+p.cpl()+int(p.n)) }

// Get returns the first value of the key rest (the key from the end of the path on), or
// false if the page does not hold it.
func (p *Fixed) Get[T comparable](rest []byte) (T, bool) {
	var zero T
	if p.Match(rest) < p.cpl() {
		return zero, false
	}
	pos, _, found := p.locate(rest[p.cpl():])
	if !found {
		return zero, false
	}
	w, _ := sizeAlign[T]()
	return *valueAt[T](p.mem(), p.vs(w), pos), true
}

// EachValue calls fn with every value of the key rest, in the order they came in, until fn
// returns false, and reports whether the page holds the key.
func (p *Fixed) EachValue[T comparable](rest []byte, fn func(v T) bool) bool {
	if p.Match(rest) < p.cpl() {
		return false
	}
	pos, _, found := p.locate(rest[p.cpl():])
	if !found {
		return false
	}
	m := p.mem()
	n := int(p.n)
	lo := Header + p.cpl()
	w, _ := sizeAlign[T]()
	vs := p.vs(w)
	for i := pos; ; {
		if !fn(*valueAt[T](m, vs, i)) {
			return true
		}
		if i++; i >= n || m[lo+i] != Further {
			return true
		}
	}
}

// Insert is Add for a page whose keys have one value each: a key that is there with another
// value is left alone, and the answer is Differs (the tree builds it again). Add replaces it
// once the tree holds several values in a page.
func (p *Fixed) Insert[T comparable](rest []byte, v T) (*Fixed, Result) { return p.add(rest, v, false) }

// Add adds the value v to the key rest and returns the page that holds the result, which is
// p itself unless the content no longer fits p's class, and says what happened (see Result).
// A key that is there gets the value behind its others.
func (p *Fixed) Add[T comparable](rest []byte, v T) (*Fixed, Result) { return p.add(rest, v, true) }

func (p *Fixed) add[T comparable](rest []byte, v T, multi bool) (*Fixed, Result) {
	cpl := p.cpl()
	if p.Match(rest) < cpl {
		return p, Outside
	}
	r := rest[cpl:]
	pos, off, found := p.locate(r)
	n := int(p.n)
	w, _ := sizeAlign[T]()
	vs := p.vs(w)
	m := p.mem()
	lo := Header + cpl
	rl, slot, at := uint8(len(r)), pos, off
	if found {
		for {
			if *valueAt[T](m, vs, slot) == v {
				return p, Present
			}
			if slot++; slot >= n || m[lo+slot] != Further {
				break
			}
		}
		if !multi {
			return p, Differs
		}
		rl, r, at = Further, nil, off+len(r)
	}
	if len(r) > MaxRemainder {
		return p, Full
	}
	keyEnd := p.keyEndFrom(pos, off)
	nk := keyEnd + 1 + len(r)
	c := classFor(nk + (n+1)*w)
	if c < 0 {
		return p, Full
	}
	q, old := p, m
	if c > p.class() {
		q = (*Fixed)(grow(&p.head, c, keyEnd))
		m = q.mem()
	}
	// the array grows at its front: the values before the new one move down by one slot,
	// those behind it stay where they are (from the old object, if the page grew)
	nvs := len(m) - (n+1)*w
	copy(m[nvs:nvs+slot*w], old[vs:vs+slot*w])
	if q != p {
		copy(m[nvs+(slot+1)*w:], old[vs+slot*w:vs+n*w])
	}
	a := lo + slot
	copy(m[at+1+len(r):nk], m[at:keyEnd])
	copy(m[a+1:at+1], m[a:at])
	m[a] = rl
	copy(m[at+1:], r)
	*valueAt[T](m, nvs, slot) = v
	q.n++
	return q, Added + Result(b2i(found))
}

// Widen is Page.Widen for values of one size: it adds the entry rest -> v, a key that leaves
// the common prefix after mis = Match(rest) < PrefixLen() bytes, and shortens the prefix to those
// bytes. It returns the new page, or nil if the entry does not go in.
func (p *Fixed) Widen[T comparable](rest []byte, v T) *Fixed {
	cpl := p.cpl()
	mis := p.Match(rest)
	d := cpl - mis
	n := int(p.n)
	newRem := rest[mis:]
	if len(newRem) > MaxRemainder {
		return nil
	}
	w, _ := sizeAlign[T]()
	m := p.mem()
	lo := Header + cpl
	keyEnd := p.Used()
	vs := p.vs(w)
	heads, longest := 0, 0
	for _, rl := range m[lo : lo+n] {
		if rl != Further {
			heads++
			longest = max(longest, int(rl))
		}
	}
	nk := keyEnd + 1 - d + heads*d + len(newRem)
	c := classFor(nk + (n+1)*w)
	if c < 0 || longest+d > MaxRemainder {
		return nil
	}
	first := len(newRem) == 0 || newRem[0] < p.CP()[mis]
	at := 0
	if !first {
		at = n
	}
	q := (*Fixed)(allocRaw(c))
	newHead(&q.head, c, n+1, mis)
	out := q.mem()
	nlo := Header + mis
	for i := range n {
		rl := m[lo+i]
		if rl != Further {
			rl += uint8(d)
		}
		out[nlo+i+b2i(i >= at)] = rl
	}
	out[nlo+at] = uint8(len(newRem))
	off := nlo + n + 1
	copy(out[Header:], p.CP()[:mis])
	src := lo + n
	extra := p.CP()[mis:]
	if first {
		off += copy(out[off:], newRem)
	}
	for i := range n {
		if rl := int(m[lo+i]); rl != Further {
			off += copy(out[off:], extra)
			off += copy(out[off:], m[src:src+rl])
			src += rl
		}
	}
	if !first {
		copy(out[off:], newRem)
	}
	nvs := len(out) - (n+1)*w
	copy(out[nvs+b2i(first)*w:nvs+b2i(first)*w+n*w], m[vs:vs+n*w])
	*valueAt[T](out, nvs, at) = v
	return q
}

// Remove removes the value v of the key rest and reports whether it was there; a key that has
// no value left is gone. It returns the page that holds the rest: p itself, a page of a smaller
// class once the content fills at most half of it, or nil if the value was the only one (the
// page is gone); see Page.Remove.
func (p *Fixed) Remove[T comparable](rest []byte, v T) (*Fixed, bool) {
	cpl := p.cpl()
	if p.Match(rest) < cpl {
		return p, false
	}
	r := rest[cpl:]
	pos, off, found := p.locate(r)
	if !found {
		return p, false
	}
	n := int(p.n)
	w, _ := sizeAlign[T]()
	vs := p.vs(w)
	m := p.mem()
	lo := Header + cpl
	slot := pos
	for *valueAt[T](m, vs, slot) != v {
		if slot++; slot >= n || m[lo+slot] != Further {
			return p, false
		}
	}
	keyEnd := p.keyEndFrom(pos, off) // before the length list changes
	more := slot+1 < n && m[lo+slot+1] == Further
	remOff, remLen := off+len(r), 0 // the remainder goes only with the last value of its key
	if slot == pos {
		if more { // the next value takes the key's place: its slot becomes the key's
			m[lo+slot+1] = uint8(len(r))
		} else {
			if n == 1 {
				return nil, true
			}
			remOff, remLen = off, len(r)
		}
	}
	nk := keyEnd - 1 - remLen
	a := lo + slot
	copy(m[a:], m[a+1:remOff])
	copy(m[remOff-1:], m[remOff+remLen:keyEnd])
	clear(m[nk:keyEnd])
	// the array shrinks at its front: the values before the removed one move up by one slot
	copy(m[vs+w:vs+(slot+1)*w], m[vs:vs+slot*w])
	clear(m[vs : vs+w])
	p.n--
	if c := classFor(2 * (nk + (n-1)*w)); c >= 0 && c < p.class() {
		q := (*Fixed)(grow(&p.head, c, nk))
		out := q.mem()
		copy(out[len(out)-(n-1)*w:], m[vs+w:vs+n*w])
		return q, true
	}
	return p, true
}

// Each calls fn with the remainder (after the common prefix), the value and whether it is the
// first value of its key, for every value in key order until fn returns false, and reports
// whether it ran to completion. The remainders alias the page; that of a further value is its
// key's.
func (p *Fixed) Each[T comparable](fn func(rem []byte, v T, first bool) bool) bool {
	m := p.mem()
	n := int(p.n)
	lo := Header + p.cpl()
	off := lo + n
	w, _ := sizeAlign[T]()
	vs := p.vs(w)
	var rem []byte
	for i, rl := range m[lo : lo+n] {
		first := rl != Further
		if first {
			rem = m[off : off+int(rl)]
			off += int(rl)
		}
		if !fn(rem, *valueAt[T](m, vs, i), first) {
			return false
		}
	}
	return true
}

// Skip drops the first k bytes of the common prefix (k at most its length): the
// tree has put a byte node above the page that consumes them. It happens in place.
func (p *Fixed) Skip[T comparable](k int) {
	m := p.mem()
	keyEnd := p.Used()
	copy(m[Header:], m[Header+k:keyEnd])
	clear(m[keyEnd-k : keyEnd])
	p.setCpl(p.cpl() - k)
}

// Prepend returns the page with pre in front of the common prefix: the page moves up
// in the tree, when the node above it goes away. It is p itself if the content still
// fits, else a page of a larger class; nil if the common prefix would exceed MaxPrefix
// or the content the largest class.
func (p *Fixed) Prepend[T comparable](pre []byte) *Fixed {
	n := int(p.n)
	w, _ := sizeAlign[T]()
	keyEnd := p.Used()
	nk := keyEnd + len(pre)
	c := classFor(nk + n*w)
	if p.cpl()+len(pre) > MaxPrefix || c < 0 {
		return nil
	}
	q := p
	if c > p.class() {
		q = (*Fixed)(grow(&p.head, c, keyEnd))
		copy(q.mem()[q.Size()-n*w:], p.mem()[p.Size()-n*w:])
	}
	m := q.mem()
	copy(m[Header+len(pre):nk], m[Header:keyEnd])
	copy(m[Header:], pre)
	q.setCpl(q.cpl() + len(pre))
	return q
}
