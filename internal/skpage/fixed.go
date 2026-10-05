package skpage

//go:generate go run ./gen

import (
	"reflect"
	"unsafe"
)

// Fixed is the single-key page of values of one size (docs/redesign/step3-fixed-design.md):
// one object that holds the remainder of one key and all its values as an array
// of T, for the maps whose values are a word with a pointer (a *X) or small and
// without any pointer (a uint64).
//
//	type | r | n (2 bytes) | kl (2 bytes) | remainder (r bytes) | padding to the alignment of T | values [n]T
//
// The head is that of Page (and of the leaves of internal/art). The values are
// unsorted: Has scans them, Add appends after a scan that found no duplicate,
// Remove moves the last value into the gap and clears the last slot, so that
// the page holds on to nothing it no longer contains. The values of a key are a
// set, compared as T (a == b), not as bytes.
//
// A page whose T holds no pointer is an array of words that the garbage
// collector never scans (Supported tells which T take it). A page whose T is a
// word with a pointer must be allocated as a Go type with its pointers marked:
// the words that hold head and remainder are not pointers, the words behind
// them are (fixed_ptr_gen.go has one type for every class and every number of
// words in front, 166). The methods are generic in T and do not carry it: the tree
// holds untyped *Fixed pointers, as it holds *Page, and the map that owns the
// page names T in every call. T must be the type the page was made with.
type Fixed struct{ head }

// fixedKind says how the page of a T is allocated.
type fixedKind int

const (
	fixedNone fixedKind = iota // not a page
	fixedRaw                   // no pointer: an array of words
	fixedPtr                   // one word that is a pointer: a typed object
)

// maxFixedValue is the largest T without a pointer that takes pages; larger
// values would leave too few per class.
const maxFixedValue = 16

// kindOf returns how the page of T is allocated, or fixedNone.
func kindOf[T comparable]() fixedKind {
	var z T
	switch w := unsafe.Sizeof(z); {
	case w == 0:
		return fixedNone
	case pointerFree(reflect.TypeFor[T]()):
		if w <= maxFixedValue {
			return fixedRaw
		}
	case w == 8 && unsafe.Alignof(z) == 8:
		return fixedPtr // 8 bytes with a pointer in them: the pointer is all of it
	}
	return fixedNone
}

// Supported reports whether values of type T go into a Fixed page: T of 1 to 16
// bytes without a pointer, or one word that is a pointer (a pointer, a map, a
// channel, a function, an unsafe.Pointer, a struct or array of one of them).
// Any other T, an interface or a struct with two words, has no page.
func Supported[T comparable]() bool { return kindOf[T]() != fixedNone }

// HoldsPointers reports whether the page of T is an object with pointers: T is
// a word that is a pointer. Such an object is scanned by the garbage
// collector and, in the tree's statistic, an object with pointers.
func HoldsPointers[T comparable]() bool { return kindOf[T]() == fixedPtr }

func pointerFree(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return true
	case reflect.Array:
		return t.Len() == 0 || pointerFree(t.Elem())
	case reflect.Struct:
		for i := range t.NumField() {
			if !pointerFree(t.Field(i).Type) {
				return false
			}
		}
		return true
	}
	return false
}

// valuesAt returns the offset of the values of a page of T with a remainder of
// rem bytes: the next multiple of the alignment of T after the remainder.
func valuesAt[T comparable](rem int) int {
	var z T
	a := int(unsafe.Alignof(z))
	return (Header + rem + a - 1) &^ (a - 1)
}

// capacityOf returns how many values of T a page of class c holds with a
// remainder of rem bytes, which may be none.
func capacityOf[T comparable](c, rem int) int {
	var z T
	return max(sizes[c]-valuesAt[T](rem), 0) / int(unsafe.Sizeof(z))
}

// classHolding returns the smallest class that holds n values of T behind a
// remainder of rem bytes, or -1.
func classHolding[T comparable](n, rem int) int {
	for c := range sizes {
		if capacityOf[T](c, rem) >= n {
			return c
		}
	}
	return -1
}

// MaxRemainderFixed returns the longest remainder a page of T holds: the one that
// leaves room for a single value in the largest class.
func MaxRemainderFixed[T comparable]() int {
	var z T
	a := int(unsafe.Alignof(z))
	// the values start at the first multiple of a that leaves room for one of them
	off := (sizes[Classes-1] - int(unsafe.Sizeof(z))) &^ (a - 1)
	return off - Header
}

// valuesOf returns the first n slots of the values area of p, which is a page of
// T; n may be up to the capacity of the page. The slice aliases the page.
func valuesOf[T comparable](p *Fixed, n int) []T {
	return unsafe.Slice((*T)(unsafe.Add(unsafe.Pointer(p), valuesAt[T](p.rem()))), n)
}

// allocFixed returns a zeroed page of class c with the head set, for a
// remainder of r bytes of a key of kl bytes. It does not check that a value
// fits.
func allocFixed[T comparable](c, r, kl int) *Fixed {
	var p *Fixed
	if kindOf[T]() == fixedPtr {
		p = (*Fixed)(allocPtr[T](c, valuesAt[T](r)/8))
	} else {
		p = (*Fixed)(allocRaw(c))
	}
	p.objType, p.kl = TypeBase+uint8(c)<<1, uint16(kl)
	p.setRem(r)
	return p
}

// NewFixed returns a page for the key of keyLen bytes whose end is rest, with
// the one value v, or nil if T has no page (see Supported) or they do not fit
// (see MaxRemainderFixed). The page copies the remainder. The tree calls it when
// a key arrives that no page holds yet, with the remainder it has cut from the
// key.
func NewFixed[T comparable](rest []byte, keyLen int, v T) *Fixed {
	if kindOf[T]() == fixedNone || len(rest) > MaxRemainderFixed[T]() {
		return nil
	}
	p := allocFixed[T](classHolding[T](1, len(rest)), len(rest), keyLen)
	copy(p.mem()[Header:], rest)
	valuesOf[T](p, 1)[0] = v
	p.n = 1
	return p
}

// EmptyFixed returns a page without values for the key of keyLen bytes whose end is
// rest, or nil if T has no page or rest is longer than MaxRemainderFixed. The tree
// makes it when a key arrives and adds the first value at once with Add; a page
// without values is not a state a key stays in.
func EmptyFixed[T comparable](rest []byte, keyLen int) *Fixed {
	if kindOf[T]() == fixedNone || len(rest) > MaxRemainderFixed[T]() {
		return nil
	}
	p := allocFixed[T](classHolding[T](1, len(rest)), len(rest), keyLen)
	copy(p.mem()[Header:], rest)
	return p
}

// BuildFixed returns a page for the key of keyLen bytes whose end is rest, with
// the values vals (all different, at least one), or nil if T has no page or
// they do not fit the largest class. The tree calls it when the value set of a
// key has shrunk to what BackFits allows.
func BuildFixed[T comparable](rest []byte, keyLen int, vals []T) *Fixed {
	if kindOf[T]() == fixedNone || len(vals) == 0 || len(rest) > MaxRemainderFixed[T]() {
		return nil
	}
	c := classHolding[T](len(vals), len(rest))
	if c < 0 {
		return nil
	}
	p := allocFixed[T](c, len(rest), keyLen)
	copy(p.mem()[Header:], rest)
	copy(valuesOf[T](p, len(vals)), vals)
	p.n = uint16(len(vals))
	return p
}

// Values returns the values of the page, in no particular order. The slice
// aliases the page and is valid until the page changes.
func (p *Fixed) Values[T comparable]() []T { return valuesOf[T](p, int(p.n)) }

// Has reports whether v is one of the values.
func (p *Fixed) Has[T comparable](v T) bool {
	for _, x := range valuesOf[T](p, int(p.n)) {
		if x == v {
			return true
		}
	}
	return false
}

// Each calls fn with every value until fn returns false, and reports whether
// it ran to completion.
func (p *Fixed) Each[T comparable](fn func(v T) bool) bool {
	for _, x := range valuesOf[T](p, int(p.n)) {
		if !fn(x) {
			return false
		}
	}
	return true
}

// Add adds v to the values of the key. It returns the page that holds the
// result, which is p itself unless it was full, and says what happened. With
// Full (the largest class does not hold one more value) the page is unchanged
// and the key belongs in a value overflow.
func (p *Fixed) Add[T comparable](v T) (*Fixed, Result) {
	n := int(p.n)
	if p.Has(v) {
		return p, Present
	}
	if n < capacityOf[T](p.class(), p.rem()) {
		valuesOf[T](p, n+1)[n] = v
		p.n++
		return p, Added
	}
	c := classHolding[T](n+1, p.rem())
	if c < 0 {
		return p, Full
	}
	q := allocFixed[T](c, p.rem(), int(p.kl))
	copy(q.mem()[Header:], p.Rest())
	copy(valuesOf[T](q, n+1), valuesOf[T](p, n))
	valuesOf[T](q, n+1)[n] = v
	q.n = p.n + 1
	return q, Added
}

// Remove removes v from the values and reports whether it was there. It returns
// the page that holds the rest: p itself, a page of a smaller class once the
// values fill at most half of it (a class that holds twice as many as are left),
// or nil if v was the only value, when the key is gone. Growing takes the
// smallest class that holds one more, so a key whose values hover at a class
// border does not change its object with every added and removed value: after
// shrinking, the page is half empty and takes as many additions to grow again.
func (p *Fixed) Remove[T comparable](v T) (*Fixed, bool) {
	vs := valuesOf[T](p, int(p.n))
	i := 0
	for i < len(vs) && vs[i] != v {
		i++
	}
	switch {
	case i == len(vs):
		return p, false
	case len(vs) == 1:
		return nil, true
	}
	var zero T
	vs[i] = vs[len(vs)-1]
	vs[len(vs)-1] = zero
	p.n--
	if c := classHolding[T](2*int(p.n), p.rem()); c >= 0 && c < p.class() {
		q := allocFixed[T](c, p.rem(), int(p.kl))
		copy(q.mem()[Header:], p.Rest())
		copy(valuesOf[T](q, int(p.n)), valuesOf[T](p, int(p.n)))
		q.n = p.n
		return q, true
	}
	return p, true
}

// Prepend returns the page with pre in front of the remainder: the page moves
// up in the tree, when the node above it goes away. It is p itself if the
// values still fit and, for a page with pointers, the words in front of them
// are as many as before; else a page of the class and the layout that fit. It
// is nil, and p unchanged, if the remainder would exceed MaxRemainderFixed or
// the values no longer fit the largest class behind it: the key belongs in a
// value overflow.
func (p *Fixed) Prepend[T comparable](pre []byte) *Fixed {
	r := p.rem() + len(pre)
	n := int(p.n)
	if r > MaxRemainderFixed[T]() {
		return nil
	}
	c := classHolding[T](n, r)
	if c < 0 {
		return nil
	}
	from, to := valuesAt[T](p.rem()), valuesAt[T](r)
	if c == p.class() && (from == to || kindOf[T]() == fixedRaw) {
		if from != to { // no pointer: the values move as bytes
			var z T
			m := p.mem()
			w := n * int(unsafe.Sizeof(z))
			copy(m[to:to+w], m[from:from+w])
			clear(m[from:min(to, from+w)]) // what the move left behind
		}
		m := p.mem()
		copy(m[Header+len(pre):], m[Header:Header+p.rem()])
		copy(m[Header:], pre)
		p.setRem(r)
		return p
	}
	q := allocFixed[T](c, r, int(p.kl))
	m := q.mem()
	copy(m[Header:], pre)
	copy(m[Header+len(pre):], p.Rest())
	copy(valuesOf[T](q, n), valuesOf[T](p, n))
	q.n = p.n
	return q
}
