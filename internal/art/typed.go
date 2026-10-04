package art

import (
	"reflect"
	"slices"
	"unsafe"
)

// A typed leaf holds a key and its values like a flat leaf, in one object, for
// values that hold a pointer (strings, pointers, interfaces, structs of them):
//
//	leafHead (6 B) | key remainder (klen bytes) | padding to 8 | values [n]T
//
// Such an object must be allocated with its real type, or the garbage
// collector would not see the pointers in the values. The type is a struct of
// the head, a key array and a value array, so a leaf has one of eight value
// capacities (typedCaps, the leaf's kind) and one of eight key area sizes
// (2 to 58 bytes, so that the head and the key fill whole words and the values
// start at typedOff(klen) in every one of them). A longer remainder makes the
// leaf a set leaf, as does a key with more values than the largest capacity
// holds; the set leaf turns into a typed one again once its values fill half
// of it.
//
// Values are unsorted as in a flat leaf: a lookup scans them, an insertion
// appends after the scan found no duplicate, and a removal moves the last value
// into the gap, and clears the last slot so that the leaf does not hold on to
// what it no longer contains.

// typedCaps are the value capacities of the typed leaf classes, indexed by
// leafHead.cls(); class 0 marks a set leaf.
var typedCaps = [...]int{0, 1, 2, 3, 4, 6, 8, 12, 16}

// minGrownTyped is the smallest class a typed leaf grows into or shrinks back
// to, four values: a key with a second value is likely to get a few more, and
// each move to a bigger class allocates an object the garbage collector must
// scan, which costs more than the room that is left over.
const minGrownTyped = 4

// maxTypedKey is the longest key remainder a typed leaf holds, and
// maxTypedValue the largest T that takes typed leaves.
const (
	maxTypedKey   = 58
	maxTypedValue = 64
)

// typedType reports whether T takes typed leaves: pointer-holding, so that it
// cannot take flat ones, small enough to leave room for several per leaf, and
// aligned to a word, which the leaf's layout assumes.
func typedType[T comparable]() bool {
	var z T
	return unsafe.Sizeof(z) > 0 && unsafe.Sizeof(z) <= maxTypedValue && unsafe.Alignof(z) == 8 && !pointerFree(reflect.TypeFor[T]())
}

// typedOff returns the offset of the values in a typed leaf with a key
// remainder of klen bytes: the next multiple of 8 after the key.
func typedOff(klen int) uintptr { return (keyOff + uintptr(klen) + 7) &^ 7 }

// typedClassFor returns the smallest class that holds n values, or 0 if none
// does.
func typedClassFor(n int) uint8 {
	for c := 1; c < len(typedCaps); c++ {
		if typedCaps[c] >= n {
			return uint8(c)
		}
	}
	return 0
}

// valArr and keyArr are the value arrays and key areas of the typed leaf
// types.
type valArr[T any] interface {
	[1]T | [2]T | [3]T | [4]T | [6]T | [8]T | [12]T | [16]T
}

type keyArr interface {
	[2]byte | [10]byte | [18]byte | [26]byte | [34]byte | [42]byte | [50]byte | [58]byte
}

// typedLeaf is the type of a typed leaf object.
type typedLeaf[T any, V valArr[T], K keyArr] struct {
	leafHead
	k K
	v V
}

// allocTyped allocates an empty typed leaf of class c for a key remainder of
// klen bytes, at most maxTypedKey; the caller stores the key.
func allocTyped[T comparable](c uint8, klen int) *leafHead {
	j := int(typedOff(klen)-8) / 8
	var l *leafHead
	switch c {
	case 1:
		l = allocTypedK[T, [1]T](j)
	case 2:
		l = allocTypedK[T, [2]T](j)
	case 3:
		l = allocTypedK[T, [3]T](j)
	case 4:
		l = allocTypedK[T, [4]T](j)
	case 5:
		l = allocTypedK[T, [6]T](j)
	case 6:
		l = allocTypedK[T, [8]T](j)
	case 7:
		l = allocTypedK[T, [12]T](j)
	default:
		l = allocTypedK[T, [16]T](j)
	}
	l.kind = kSet + kind(c)
	return l
}

// allocTypedK allocates a typed leaf with value array V and the key area of
// size class j.
func allocTypedK[T comparable, V valArr[T]](j int) *leafHead {
	switch j {
	case 0:
		return &new(typedLeaf[T, V, [2]byte]).leafHead
	case 1:
		return &new(typedLeaf[T, V, [10]byte]).leafHead
	case 2:
		return &new(typedLeaf[T, V, [18]byte]).leafHead
	case 3:
		return &new(typedLeaf[T, V, [26]byte]).leafHead
	case 4:
		return &new(typedLeaf[T, V, [34]byte]).leafHead
	case 5:
		return &new(typedLeaf[T, V, [42]byte]).leafHead
	case 6:
		return &new(typedLeaf[T, V, [50]byte]).leafHead
	default:
		return &new(typedLeaf[T, V, [58]byte]).leafHead
	}
}

// newTypedLeaf allocates a leaf for key that holds it from base on, a typed
// leaf of the smallest class or, if the remainder is longer than a typed leaf
// holds, a set leaf. It captures nothing, so passing it as a newLeafFunc
// allocates no closure.
func newTypedLeaf[T comparable](key []byte, base int) *leafHead {
	if len(key) > maxKeyLen || len(key)-base > maxTypedKey {
		return newSetLeaf[T](key, base)
	}
	return typedWithKey[T](1, key[base:], len(key))
}

// typedWithKey allocates a typed leaf of class c holding a copy of the key
// remainder s of a key of kl bytes, and no values.
func typedWithKey[T comparable](c uint8, s []byte, kl int) *leafHead {
	l := allocTyped[T](c, len(s))
	l.klen, l.kl = uint8(len(s)), uint16(kl)
	copy(unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), len(s)), s)
	return l
}

// typedVals returns the values of typed leaf l. The slice aliases the leaf.
func typedVals[T comparable](l *leafHead) []T {
	return unsafe.Slice((*T)(unsafe.Add(unsafe.Pointer(l), typedOff(int(l.klen)))), l.n)
}

// typedAdd adds v to typed leaf l. It returns the leaf that replaces l when l
// had no room, or nil when l still holds the key's values.
func typedAdd[T comparable](l *leafHead, v T) *leafHead {
	vs := typedVals[T](l)
	if slices.Contains(vs, v) {
		return nil
	}
	n := int(l.n)
	if n < typedCaps[l.cls()] {
		appendTyped(l, v)
		return nil
	}
	c := typedClassFor(2 * n)
	if c == 0 {
		c = uint8(len(typedCaps) - 1) // the largest class
	}
	c = max(c, minGrownTyped)
	if typedCaps[c] <= n {
		return typedSpill(l, v)
	}
	nl := resizeTyped[T](l, c)
	appendTyped(nl, v)
	return nl
}

// appendTyped stores v after the values of typed leaf l, which has room.
func appendTyped[T comparable](l *leafHead, v T) {
	n := int(l.n)
	unsafe.Slice((*T)(unsafe.Add(unsafe.Pointer(l), typedOff(int(l.klen)))), n+1)[n] = v
	l.n++
}

// typedRemove removes v from typed leaf l. empty reports that l held v as its
// last value; shrink is the smaller class l should move to (see resizeTyped),
// or 0. The caller moves it, since only it can reach the leaf's slot.
func typedRemove[T comparable](l *leafHead, v T) (shrink uint8, empty bool) {
	vs := typedVals[T](l)
	i := slices.Index(vs, v)
	if i < 0 {
		return 0, false
	}
	last := len(vs) - 1
	vs[i] = vs[last]
	var zero T
	vs[last] = zero
	l.n--
	if l.n == 0 {
		return 0, true
	}
	if c := typedClassFor(2 * int(l.n)); c != 0 && max(c, minGrownTyped) < l.cls() {
		return max(c, minGrownTyped), false
	}
	return 0, false
}

// resizeTyped copies typed leaf l into a new leaf of class c, which holds its
// values.
func resizeTyped[T comparable](l *leafHead, c uint8) *leafHead {
	nl := typedWithKey[T](c, l.stored(), l.keyLen())
	copy(unsafe.Slice((*T)(unsafe.Add(unsafe.Pointer(nl), typedOff(int(nl.klen)))), l.n), typedVals[T](l))
	nl.n = l.n
	return nl
}

// typedSpill turns typed leaf l, which is full in the largest class, into a
// set leaf holding its values and v.
func typedSpill[T comparable](l *leafHead, v T) *leafHead {
	sl := newSetLeafOf[T](l.stored(), l.keyLen())
	s := vals[T](sl)
	for _, x := range typedVals[T](l) {
		s.Add(x)
	}
	s.Add(v)
	return sl
}

// unspillTypedClass returns the class that set leaf l should turn into once
// its values fill at most half of the largest class, or 0 if l stays a set
// leaf.
func unspillTypedClass[T comparable](l *leafHead) uint8 {
	if len(l.stored()) > maxTypedKey || l.keyLen() > maxKeyLen {
		return 0
	}
	return typedClassFor(2 * vals[T](l).Len())
}

// unspillTyped copies set leaf l into a typed leaf of class c.
func unspillTyped[T comparable](l *leafHead, c uint8) *leafHead {
	nl := typedWithKey[T](c, l.stored(), l.keyLen())
	vals[T](l).Each(func(v T) bool { appendTyped(nl, v); return true })
	return nl
}

// rekeyTyped is rekey for a map of typed leaves: a typed leaf grows to hold
// its longer key in place when the values stay where they are, and moves into
// a typed leaf of the longer key, or a set leaf, otherwise.
func rekeyTyped[T comparable](l *leafHead, pre []byte, b, pathLen int) *leafHead {
	if l.kind == kSet {
		return rekey[T](l, pre, b, pathLen)
	}
	old, klen := int(l.klen), l.keyLen()-pathLen
	if klen <= maxTypedKey && typedOff(klen) == typedOff(old) {
		area := unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), klen)
		copy(area[klen-old:], area[:old])
		fillHead(area[:klen-old], pre, b, pathLen)
		l.klen = uint8(klen)
		return l
	}
	key := wholeKey(l, pre, b)
	if klen > maxTypedKey {
		nl := newSetLeaf[T](key, pathLen)
		s := vals[T](nl)
		for _, v := range typedVals[T](l) {
			s.Add(v)
		}
		return nl
	}
	nl := typedWithKey[T](l.cls(), key[pathLen:], len(key))
	copy(unsafe.Slice((*T)(unsafe.Add(unsafe.Pointer(nl), typedOff(klen))), l.n), typedVals[T](l))
	nl.n = l.n
	return nl
}
