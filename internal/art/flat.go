package art

import (
	"reflect"
	"slices"
	"unsafe"
)

// A flat leaf holds a key and its values in one object without pointers:
//
//	leafHead (6 B) | key remainder (klen bytes) | padding to T's alignment | values [n]T
//
// The object is one of the Go size classes in flatSizes, allocated as an
// array of uint64 so that the garbage collector never scans it. A new key
// gets the smallest class that holds its first value, so that a key with one
// value, the most common case, takes as little memory as possible. When the
// class is full, the leaf moves to the smallest class that holds twice its
// values, and to at least a cache line (minGrown); it moves back to a
// smaller class only once that class would still be half empty, and never
// below a cache line. A key whose count hovers at a class boundary, or
// between one and a few values, therefore does not move on every change. A key with more values than the largest class holds becomes
// a set leaf (spill), and a set leaf becomes flat again once its values fill
// half of the largest class (unspill).
//
// Values are unsorted: a lookup scans them, an insertion appends after the
// scan found no duplicate, and a removal moves the last value into the gap.
// Up to 512 bytes the scan reads a few adjacent cache lines, which costs less
// than the extra cache miss of a separate value array.

// flatSizes are the size classes of flat leaves, indexed by leafHead.cls();
// class 0 marks a set leaf.
var flatSizes = [...]uintptr{0, 32, 48, 64, 96, 128, 192, 256, 384, 512}

// minGrown is the smallest class a flat leaf grows into or shrinks back to:
// 64 bytes, one cache line.
const minGrown = 3

// maxFlatValue is the largest T that takes flat leaves; larger values would
// leave too few per class.
const maxFlatValue = 16

// flatType reports whether T takes flat leaves: small, not empty and free of
// pointers, so that its values may live in memory the garbage collector does
// not scan.
func flatType[T comparable]() bool {
	var z T
	return unsafe.Sizeof(z) > 0 && unsafe.Sizeof(z) <= maxFlatValue && pointerFree(reflect.TypeFor[T]())
}

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

// flatOff returns the offset of the values in a flat leaf with a key of klen
// bytes.
func flatOff[T comparable](klen int) uintptr {
	var z T
	a := unsafe.Alignof(z)
	return (keyOff + uintptr(klen) + a - 1) &^ (a - 1)
}

// flatCap returns how many values a flat leaf of class cls holds after a key
// of klen bytes.
func flatCap[T comparable](cls uint8, klen int) int {
	var z T
	return int((flatSizes[cls] - flatOff[T](klen)) / unsafe.Sizeof(z))
}

// flatClass returns the smallest class that holds n values after a key of
// klen bytes, or 0 if none does.
func flatClass[T comparable](klen, n int) uint8 {
	var z T
	need := flatOff[T](klen) + uintptr(n)*unsafe.Sizeof(z)
	for c := 1; c < len(flatSizes); c++ {
		if flatSizes[c] >= need {
			return uint8(c)
		}
	}
	return 0
}

// allocFlat allocates an empty flat leaf of class cls.
func allocFlat(cls uint8) *leafHead {
	var p unsafe.Pointer
	switch flatSizes[cls] {
	case 32:
		p = unsafe.Pointer(new([4]uint64))
	case 48:
		p = unsafe.Pointer(new([6]uint64))
	case 64:
		p = unsafe.Pointer(new([8]uint64))
	case 96:
		p = unsafe.Pointer(new([12]uint64))
	case 128:
		p = unsafe.Pointer(new([16]uint64))
	case 192:
		p = unsafe.Pointer(new([24]uint64))
	case 256:
		p = unsafe.Pointer(new([32]uint64))
	case 384:
		p = unsafe.Pointer(new([48]uint64))
	default:
		p = unsafe.Pointer(new([64]uint64))
	}
	l := (*leafHead)(p)
	l.kind = kSet + kind(cls)
	return l
}

// newFlatLeaf allocates a flat leaf that holds key from base on, in the
// smallest class that holds one value; a key that no flat leaf holds gets a
// set leaf. It captures nothing, so passing it as a newLeafFunc allocates no
// closure.
func newFlatLeaf[T comparable](key []byte, base int) *leafHead {
	c := flatClassFor[T](key, base, 1)
	if c == 0 {
		return newSetLeaf[T](key, base)
	}
	return flatWithKey(c, key[base:], len(key))
}

// flatClassFor returns the smallest class of a flat leaf that holds key from
// base on and n values, or 0 if none does.
func flatClassFor[T comparable](key []byte, base, n int) uint8 {
	if len(key) > maxKeyLen || len(key)-base > maxInline {
		return 0
	}
	return flatClass[T](len(key)-base, n)
}

// flatWithKey allocates a flat leaf of class c holding a copy of the key
// remainder s of a key of kl bytes, and no values.
func flatWithKey(c uint8, s []byte, kl int) *leafHead {
	l := allocFlat(c)
	l.klen, l.kl = uint8(len(s)), uint16(kl)
	copy(unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), len(s)), s)
	return l
}

// reflat copies flat leaf l into a leaf that holds key from base on (see
// rekeyFunc): a flat leaf at least l's class if one holds the values, a set
// leaf otherwise.
func reflat[T comparable](l *leafHead, key []byte, base int) *leafHead {
	vs := flatVals[T](l)
	c := flatClassFor[T](key, base, len(vs))
	if c == 0 {
		nl := newSetLeaf[T](key, base)
		s := vals[T](nl)
		for _, v := range vs {
			s.Add(v)
		}
		return nl
	}
	nl := flatWithKey(max(c, l.cls()), key[base:], len(key))
	for _, v := range vs {
		appendFlat(nl, v)
	}
	return nl
}

// flatPrepend makes flat leaf l hold its key from depth on in place, by moving
// its key remainder and values up, when they still fit its size class: the copy
// that reflat would make costs an allocation, which a delete that merges a node
// into its only leaf would otherwise pay every time. The bytes it adds come
// from a rekeyFunc's pre and b. It reports whether it did.
func flatPrepend[T comparable](l *leafHead, pre []byte, b, depth int) bool {
	var z T
	old, klen := int(l.klen), l.keyLen()-depth
	if klen > maxInline || flatOff[T](klen)+uintptr(l.n)*unsafe.Sizeof(z) > flatSizes[l.cls()] {
		return false
	}
	nb := uintptr(l.n) * unsafe.Sizeof(z)
	from, to := flatOff[T](old), flatOff[T](klen)
	copy(unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), to)), nb), unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), from)), nb))
	area := unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), klen)
	copy(area[klen-old:], area[:old])
	fillHead(area[:klen-old], pre, b, depth)
	l.klen = uint8(klen)
	return true
}

// flatVals returns the values of flat leaf l. The slice aliases the leaf.
func flatVals[T comparable](l *leafHead) []T {
	return unsafe.Slice((*T)(unsafe.Add(unsafe.Pointer(l), flatOff[T](int(l.klen)))), l.n)
}

// flatAdd adds v to flat leaf l. It returns the leaf that replaces l when l
// had no room, or nil when l still holds the key's values.
func flatAdd[T comparable](l *leafHead, v T) *leafHead {
	if slices.Contains(flatVals[T](l), v) {
		return nil
	}
	n := int(l.n)
	if n == flatCap[T](l.cls(), int(l.klen)) {
		if flatClass[T](int(l.klen), n+1) == 0 {
			return spill(l, v)
		}
		c := flatClass[T](int(l.klen), 2*n)
		if c == 0 {
			c = uint8(len(flatSizes) - 1) // the largest class, which holds n+1
		}
		l = resize[T](l, max(c, minGrown))
		appendFlat(l, v)
		return l
	}
	appendFlat(l, v)
	return nil
}

// appendFlat stores v after the values of flat leaf l, which has room.
func appendFlat[T comparable](l *leafHead, v T) {
	n := int(l.n)
	unsafe.Slice((*T)(unsafe.Add(unsafe.Pointer(l), flatOff[T](int(l.klen)))), n+1)[n] = v
	l.n++
}

// flatRemove removes v from flat leaf l. empty reports that l held v as its
// last value; shrink is the smaller class l should move to (see resize), or
// 0. The caller moves it, since only it can reach the leaf's slot.
func flatRemove[T comparable](l *leafHead, v T) (shrink uint8, empty bool) {
	vs := flatVals[T](l)
	i := slices.Index(vs, v)
	if i < 0 {
		return 0, false
	}
	last := len(vs) - 1
	vs[i] = vs[last]
	l.n--
	if l.n == 0 {
		return 0, true
	}
	if c := flatClass[T](int(l.klen), 2*int(l.n)); c != 0 && max(c, minGrown) < l.cls() {
		return max(c, minGrown), false
	}
	return 0, false
}

// resize copies flat leaf l into a new leaf of class c, which holds its values.
func resize[T comparable](l *leafHead, c uint8) *leafHead {
	var z T
	used := flatOff[T](int(l.klen)) + uintptr(l.n)*unsafe.Sizeof(z)
	nl := allocFlat(c)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(nl)), used), unsafe.Slice((*byte)(unsafe.Pointer(l)), used))
	nl.kind = kSet + kind(c)
	return nl
}

// spill turns flat leaf l, which is full in the largest class, into a set
// leaf holding its values and v.
func spill[T comparable](l *leafHead, v T) *leafHead {
	sl := newSetLeafOf[T](l.stored(), l.keyLen())
	s := vals[T](sl)
	for _, x := range flatVals[T](l) {
		s.Add(x)
	}
	s.Add(v)
	return sl
}

// unspillClass returns the class that set leaf l should turn into once its
// values fill at most half of the largest class, or 0 if l stays a set leaf.
func unspillClass[T comparable](l *leafHead) uint8 {
	s := l.stored()
	if len(s) > maxInline || l.keyLen() > maxKeyLen {
		return 0
	}
	return flatClass[T](len(s), 2*vals[T](l).Len())
}

// unspill copies set leaf l into a flat leaf of class c.
func unspill[T comparable](l *leafHead, c uint8) *leafHead {
	nl := flatWithKey(c, l.stored(), l.keyLen())
	vals[T](l).Each(func(v T) bool { appendFlat(nl, v); return true })
	return nl
}
