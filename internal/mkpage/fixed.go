package mkpage

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/skpage"
)

// Fixed is the multi-key page of values of one size, see the package comment: for
// the maps whose values are small and hold no pointer (a uint64). The values are
// an array of T behind the remainders, in the order of the entries, and compared as
// T (a == b), not as bytes. The pointer page (a T that is a word with a pointer,
// the typed objects of skpage) is step 4.3.
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

// valuesAt returns the offset of the values of a page whose keys end at keyEnd:
// the next multiple of the alignment of T.
func valuesAt[T comparable](keyEnd int) int {
	_, a := sizeAlign[T]()
	return (keyEnd + a - 1) &^ (a - 1)
}

// NeedFixed returns the bytes a Fixed page of T takes for n entries whose
// remainders (after a common prefix of cpl bytes) are remBytes in all. The tree
// calls it to decide, before it builds anything, whether the entries of a subtree
// fit a page (at most 512).
func NeedFixed[T comparable](n, cpl, remBytes int) int {
	w, _ := sizeAlign[T]()
	return valuesAt[T](Header+n+cpl+remBytes) + n*w
}

// BuildFixed returns a page for the entries with the keys rests (the keys from the
// end of the path on, in key order, all different) and the values vals, or nil if
// T has no page (see Supported) or they do not fit: no entry, more than MaxEntries, a
// remainder beyond MaxRemainder or content beyond the largest class. The common
// prefix of the page is the longest one all keys share (up to MaxPrefix).
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
	for _, r := range rests {
		if len(r)-cp > MaxRemainder {
			return nil
		}
		keyEnd += len(r) - cp
	}
	w, _ := sizeAlign[T]()
	vs := valuesAt[T](keyEnd)
	c := classFor(vs + n*w)
	if c < 0 {
		return nil
	}
	p := (*Fixed)(allocRaw(c))
	p.objType, p.n, p.cpl = TypeBase+uint8(c)<<1, uint8(n), uint8(cp)
	m := p.mem()
	off := Header + n
	off += copy(m[off:], rests[0][:cp])
	for i, r := range rests {
		m[Header+i] = uint8(len(r) - cp)
		off += copy(m[off:], r[cp:])
		*(*T)(unsafe.Pointer(&m[vs+i*w])) = vals[i]
	}
	return p
}

// locate returns the position of the entry with the remainder r, or where it would
// go, with the offset of its remainder in the object, and the offset where the
// remainders end.
func (p *Fixed) locate(r []byte) (pos, off, keyEnd int, found bool) {
	m := p.mem()
	n := int(p.n)
	off = Header + n + int(p.cpl)
	pos = n
	for i := range n {
		rl := int(m[Header+i])
		c := compare(m[off:off+rl], r)
		if c >= 0 {
			pos, found = i, c == 0
			break
		}
		off += rl
	}
	keyEnd = off
	for _, rl := range m[Header+pos : Header+n] {
		keyEnd += int(rl)
	}
	return pos, off, keyEnd, found
}

func valueAt[T comparable](m []byte, vs, i int) *T {
	w, _ := sizeAlign[T]()
	return (*T)(unsafe.Pointer(&m[vs+i*w]))
}

// Used returns the bytes of the page that hold something.
func (p *Fixed) Used() int {
	m := p.mem()
	keyEnd := Header + int(p.n) + int(p.cpl)
	for _, rl := range m[Header : Header+int(p.n)] {
		keyEnd += int(rl)
	}
	return keyEnd
}

// Get returns the value of the key rest (the key from the end of the path on), or
// false if the page does not hold it.
func (p *Fixed) Get[T comparable](rest []byte) (T, bool) {
	var zero T
	if p.Match(rest) < int(p.cpl) {
		return zero, false
	}
	pos, _, keyEnd, found := p.locate(rest[p.cpl:])
	if !found {
		return zero, false
	}
	return *valueAt[T](p.mem(), valuesAt[T](keyEnd), pos), true
}

// Insert adds the entry rest -> v and returns the page that holds the result, which is
// p itself unless the content no longer fits p's class, and says what happened (see
// Result).
func (p *Fixed) Insert[T comparable](rest []byte, v T) (*Fixed, Result) {
	if p.Match(rest) < int(p.cpl) {
		return p, Outside
	}
	r := rest[p.cpl:]
	pos, off, keyEnd, found := p.locate(r)
	n := int(p.n)
	w, _ := sizeAlign[T]()
	vs := valuesAt[T](keyEnd)
	m := p.mem()
	if found {
		if *valueAt[T](m, vs, pos) == v {
			return p, Present
		}
		return p, Differs
	}
	if len(r) > MaxRemainder {
		return p, Full
	}
	nk := keyEnd + 1 + len(r)
	nvs := valuesAt[T](nk)
	c := classFor(nvs + (n+1)*w)
	if c < 0 {
		return p, Full
	}
	q := p
	if c > p.class() {
		q = (*Fixed)(grow(&p.head, c, vs+n*w))
		m = q.mem()
	}
	copy(m[nvs+(pos+1)*w:nvs+(n+1)*w], m[vs+pos*w:vs+n*w])
	copy(m[nvs:nvs+pos*w], m[vs:vs+pos*w])
	a := Header + pos
	copy(m[off+1+len(r):nk], m[off:keyEnd])
	copy(m[a+1:off+1], m[a:off])
	m[a] = uint8(len(r))
	copy(m[off+1:], r)
	clear(m[nk:nvs])
	*valueAt[T](m, nvs, pos) = v
	q.n++
	return q, Added
}

// Widen is Page.Widen for values of one size: it adds the entry rest -> v, a key that leaves
// the common prefix after mis = Match(rest) < PrefixLen() bytes, and shortens the prefix to those
// bytes. It returns the new page, or nil if the entry does not go in.
func (p *Fixed) Widen[T comparable](rest []byte, v T) *Fixed {
	mis := p.Match(rest)
	d := int(p.cpl) - mis
	n := int(p.n)
	newRem := rest[mis:]
	if len(newRem) > MaxRemainder {
		return nil
	}
	w, _ := sizeAlign[T]()
	m := p.mem()
	keyEnd := p.Used()
	vs := valuesAt[T](keyEnd)
	nk := keyEnd + 1 - d + n*d + len(newRem)
	nvs := valuesAt[T](nk)
	c := classFor(nvs + (n+1)*w)
	if c < 0 {
		return nil
	}
	first := len(newRem) == 0 || newRem[0] < p.CP()[mis]
	at := 0
	if !first {
		at = n
	}
	q := (*Fixed)(allocRaw(c))
	q.objType, q.n, q.cpl = TypeBase+uint8(c)<<1, p.n+1, uint8(mis)
	out := q.mem()
	for i := range n { // an old remainder cannot exceed 255 bytes: the content would exceed 512
		out[Header+i+b2i(i >= at)] = m[Header+i] + uint8(d)
	}
	out[Header+at] = uint8(len(newRem))
	off := Header + n + 1
	off += copy(out[off:], p.CP()[:mis])
	src := Header + n + int(p.cpl)
	extra := p.CP()[mis:]
	if first {
		off += copy(out[off:], newRem)
	}
	for i := range n {
		rl := int(m[Header+i])
		off += copy(out[off:], extra)
		off += copy(out[off:], m[src:src+rl])
		src += rl
	}
	if !first {
		copy(out[off:], newRem)
	}
	copy(out[nvs+b2i(first)*w:nvs+b2i(first)*w+n*w], m[vs:vs+n*w])
	*valueAt[T](out, nvs, at) = v
	return q
}

// Remove removes the entry rest -> v and reports whether it was there. It returns the
// page that holds the rest: p itself, a page of a smaller class once the content
// fills at most half of it, or nil if the entry was the only one (the page is
// gone); see Page.Remove.
func (p *Fixed) Remove[T comparable](rest []byte, v T) (*Fixed, bool) {
	if p.Match(rest) < int(p.cpl) {
		return p, false
	}
	r := rest[p.cpl:]
	pos, off, keyEnd, found := p.locate(r)
	n := int(p.n)
	w, _ := sizeAlign[T]()
	vs := valuesAt[T](keyEnd)
	m := p.mem()
	if !found || *valueAt[T](m, vs, pos) != v {
		return p, false
	}
	if n == 1 {
		return nil, true
	}
	nk := keyEnd - 1 - len(r)
	nvs := valuesAt[T](nk)
	a := Header + pos
	copy(m[a:], m[a+1:off])
	copy(m[off-1:], m[off+len(r):keyEnd])
	copy(m[nvs:nvs+pos*w], m[vs:vs+pos*w])
	copy(m[nvs+pos*w:], m[vs+(pos+1)*w:vs+n*w])
	end := nvs + (n-1)*w
	clear(m[end : vs+n*w])
	clear(m[nk:nvs])
	p.n--
	if c := classFor(2 * end); c >= 0 && c < p.class() {
		return (*Fixed)(grow(&p.head, c, end)), true
	}
	return p, true
}

// Each calls fn with the remainder (after the common prefix) and the value of every
// entry in key order until fn returns false, and reports whether it ran to
// completion. The remainders alias the page.
func (p *Fixed) Each[T comparable](fn func(rem []byte, v T) bool) bool {
	m := p.mem()
	n := int(p.n)
	off := Header + n + int(p.cpl)
	vs := valuesAt[T](p.Used())
	for i := range n {
		rl := int(m[Header+i])
		if !fn(m[off:off+rl], *valueAt[T](m, vs, i)) {
			return false
		}
		off += rl
	}
	return true
}

// Skip drops the first k bytes of the common prefix (k at most its length): the
// tree has put a byte node above the page that consumes them. It happens in place.
func (p *Fixed) Skip[T comparable](k int) {
	m := p.mem()
	n := int(p.n)
	w, _ := sizeAlign[T]()
	keyEnd := p.Used()
	vs := valuesAt[T](keyEnd)
	nk := keyEnd - k
	nvs := valuesAt[T](nk)
	at := Header + n
	copy(m[at:], m[at+k:keyEnd])
	copy(m[nvs:nvs+n*w], m[vs:vs+n*w])
	clear(m[nk:nvs])
	clear(m[nvs+n*w : vs+n*w])
	p.cpl -= uint8(k)
}

// Prepend returns the page with pre in front of the common prefix: the page moves up
// in the tree, when the node above it goes away. It is p itself if the content still
// fits, else a page of a larger class; nil if the common prefix would exceed MaxPrefix
// or the content the largest class.
func (p *Fixed) Prepend[T comparable](pre []byte) *Fixed {
	n := int(p.n)
	w, _ := sizeAlign[T]()
	keyEnd := p.Used()
	vs := valuesAt[T](keyEnd)
	nk := keyEnd + len(pre)
	nvs := valuesAt[T](nk)
	c := classFor(nvs + n*w)
	if int(p.cpl)+len(pre) > MaxPrefix || c < 0 {
		return nil
	}
	q := p
	if c > p.class() {
		q = (*Fixed)(grow(&p.head, c, vs+n*w))
	}
	m := q.mem()
	copy(m[nvs:nvs+n*w], m[vs:vs+n*w])
	at := Header + n
	copy(m[at+len(pre):nk], m[at:keyEnd])
	copy(m[at:], pre)
	clear(m[nk:nvs])
	q.cpl += uint8(len(pre))
	return q
}
