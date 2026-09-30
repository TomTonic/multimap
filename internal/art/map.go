package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/vset"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. Every
// leaf holds its key's values: a flat leaf right after the key remainder
// (see flat.go), a set leaf in a vset.Set. Reading them costs no pointer chase beyond the
// leaf until a key holds more values than a flat leaf of 512 bytes. The zero
// value is an empty map.
type Map[T comparable] struct {
	t    Tree
	flat int8 // 1: T takes flat leaves, -1: set leaves only, 0: not decided yet
}

// newSetLeaf allocates a set leaf that holds key from base on, in the
// smallest size class that fits it, or the whole key as a string when the
// rest is too long to hold inline. It captures nothing, so passing it as a
// newLeafFunc allocates no closure.
func newSetLeaf[T comparable](key []byte, base int) *leafHead {
	if len(key) > maxKeyLen || len(key)-base > maxInline {
		l := &leaf[T, string]{k: string(key)}
		l.kind, l.klen = kSet, longKey
		return &l.leafHead
	}
	return newSetLeafOf[T](key[base:], len(key))
}

// newSetLeafOf allocates a set leaf holding the key remainder s, at most
// maxInline bytes, of a key of kl bytes.
func newSetLeafOf[T comparable](s []byte, kl int) *leafHead {
	switch n := len(s); {
	case n <= 16:
		return newInline[T, [16]byte](s, kl)
	case n <= 32:
		return newInline[T, [32]byte](s, kl)
	case n <= 48:
		return newInline[T, [48]byte](s, kl)
	case n <= 64:
		return newInline[T, [64]byte](s, kl)
	case n <= 96:
		return newInline[T, [96]byte](s, kl)
	case n <= 128:
		return newInline[T, [128]byte](s, kl)
	case n <= 192:
		return newInline[T, [192]byte](s, kl)
	}
	return newInline[T, [256]byte](s, kl)
}

// newInline allocates a set leaf that holds the key remainder s inline in an
// array of type K.
func newInline[T comparable, K [16]byte | [32]byte | [48]byte | [64]byte | [96]byte | [128]byte | [192]byte | [256]byte](s []byte, kl int) *leafHead {
	l := &leaf[T, K]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), s)
	l.kind, l.klen, l.kl = kSet, uint8(len(s)), uint16(kl)
	return &l.leafHead
}

// Offsets of the value set in the set leaf of each key size class. They do
// not depend on T: a vset.Set holds a pointer, so it is always 8-aligned.
var (
	setOff16  = unsafe.Offsetof(leaf[struct{}, [16]byte]{}.vals)
	setOff32  = unsafe.Offsetof(leaf[struct{}, [32]byte]{}.vals)
	setOff48  = unsafe.Offsetof(leaf[struct{}, [48]byte]{}.vals)
	setOff64  = unsafe.Offsetof(leaf[struct{}, [64]byte]{}.vals)
	setOff96  = unsafe.Offsetof(leaf[struct{}, [96]byte]{}.vals)
	setOff128 = unsafe.Offsetof(leaf[struct{}, [128]byte]{}.vals)
	setOff192 = unsafe.Offsetof(leaf[struct{}, [192]byte]{}.vals)
	setOff256 = unsafe.Offsetof(leaf[struct{}, [256]byte]{}.vals)
	setOffStr = unsafe.Offsetof(leaf[struct{}, string]{}.vals)
)

// vals returns the value set of a set leaf, created by newSetLeaf[T].
func vals[T comparable](l *leafHead) *vset.Set[T] {
	off := setOffStr
	switch k := l.klen; {
	case k <= 16:
		off = setOff16
	case k <= 32:
		off = setOff32
	case k <= 48:
		off = setOff48
	case k <= 64:
		off = setOff64
	case k <= 96:
		off = setOff96
	case k <= 128:
		off = setOff128
	case k <= 192:
		off = setOff192
	case k <= maxInline:
		off = setOff256
	}
	return (*vset.Set[T])(unsafe.Add(unsafe.Pointer(l), off))
}

// rekey returns a leaf with l's values that holds its key from depth on (see
// rekeyFunc). It captures nothing, so passing it allocates no closure. It
// keeps l where it is when the longer remainder fits its size class, and
// builds the whole key only when it has to copy l.
func rekey[T comparable](l *leafHead, pre []byte, b, depth int) *leafHead {
	if l.kind != kSet {
		if flatPrepend[T](l, pre, b, depth) {
			return l
		}
		return reflat[T](l, wholeKey(l, pre, b), depth)
	}
	if setPrepend(l, pre, b, depth) {
		return l
	}
	nl := newSetLeaf[T](wholeKey(l, pre, b), depth)
	*vals[T](nl) = *vals[T](l)
	return nl
}

// setKeyCap returns the size of the key area of a set leaf whose key
// remainder is klen bytes long, at most maxInline: the class that klen falls
// in (see vals), whose largest holds 256 bytes.
func setKeyCap(klen int) int {
	for _, c := range [...]int{16, 32, 48, 64, 96, 128, 192} {
		if klen <= c {
			return c
		}
	}
	return 256
}

// setPrepend makes set leaf l hold its key from depth on in place when the
// longer remainder still fits the key area of l's class, which keeps its value
// set where it is. It saves rekey the allocation of a new leaf, see
// flatPrepend, and reports whether it did.
func setPrepend(l *leafHead, pre []byte, b, depth int) bool {
	old, klen := int(l.klen), l.keyLen()-depth // a string leaf holds its whole key, and never gets here
	if klen > maxInline || klen > setKeyCap(old) {
		return false
	}
	area := unsafe.Slice((*byte)(unsafe.Add(unsafe.Pointer(l), keyOff)), klen)
	copy(area[klen-old:], area[:old])
	fillHead(area[:klen-old], pre, b, depth)
	l.klen = uint8(klen)
	return true
}

// decide settles once per map whether T takes flat leaves (see flatType).
func (m *Map[T]) decide() {
	m.flat = -1
	if flatType[T]() {
		m.flat = 1
	}
}

// Len returns the number of keys.
func (m *Map[T]) Len() int { return m.t.Len() }

// Clear removes all keys and values.
func (m *Map[T]) Clear() { m.t.Clear() }

// Add adds v to the values of key, creating the key if needed. The map keeps
// its own copy of key.
func (m *Map[T]) Add(key []byte, v T) {
	if m.flat == 0 {
		m.decide()
	}
	// Chosen here rather than returned from a helper: a function value that
	// does not escape stays on the stack, one that is returned is allocated.
	var nl newLeafFunc = newSetLeaf[T]
	if m.flat > 0 {
		nl = newFlatLeaf[T]
	}
	loc := m.t.upsert(key, nl)
	l := asLeaf(*loc)
	if l.kind == kSet {
		vals[T](l).Add(v)
		return
	}
	if nl := flatAdd(l, v); nl != nil {
		*loc = leafHdr(nl)
	}
}

// Remove removes v from the values of key and removes the key once it holds
// no values. Absent keys and values are ignored.
//
// It finds the leaf as a lookup does and looks for the leaf's slot only when
// the leaf has to move into a smaller one.
func (m *Map[T]) Remove(key []byte, v T) {
	l := m.t.find(key)
	if l == nil {
		return
	}
	if l.kind != kSet {
		switch c, empty := flatRemove(l, v); {
		case empty:
			m.t.remove(key, rekey[T])
		case c != 0:
			*m.t.findSlot(key) = leafHdr(resize[T](l, c))
		}
		return
	}
	s := vals[T](l)
	switch {
	case !s.Remove(v):
	case s.Len() == 0:
		m.t.remove(key, rekey[T])
	case m.flat > 0:
		if c := unspillClass[T](l); c != 0 {
			*m.t.findSlot(key) = leafHdr(unspill[T](l, c))
		}
	}
}

// RemoveKey removes key and all its values. An absent key is ignored.
func (m *Map[T]) RemoveKey(key []byte) { m.t.remove(key, rekey[T]) }

// Has reports whether key holds any values.
func (m *Map[T]) Has(key []byte) bool { return m.t.find(key) != nil }

// Each calls yield for every value of key, in unspecified order, until it
// returns false. yield must not modify the map.
func (m *Map[T]) Each(key []byte, yield func(T) bool) {
	if l := m.t.find(key); l != nil {
		eachValue(l, yield)
	}
}

// eachValue calls yield for every value of leaf l and reports whether it ran
// to completion.
func eachValue[T comparable](l *leafHead, yield func(T) bool) bool {
	if l.kind == kSet {
		return vals[T](l).Each(yield)
	}
	for _, v := range flatVals[T](l) {
		if !yield(v) {
			return false
		}
	}
	return true
}

// leafTail is the offset of the last byte of the smallest leaf this map
// creates, which the scan touches ahead (see touchChildren). Every leaf is
// at least that large, and a constant offset keeps the load independent of
// the leaf's head.
func (m *Map[T]) leafTail() uintptr {
	if m.flat > 0 {
		return flatSizes[1] - 1
	}
	return unsafe.Sizeof(leaf[T, [16]byte]{}) - 1
}

// Range calls fn for every key within b, in ascending key order, until fn
// returns false. The key is assembled for fn, from the path to its leaf and
// the rest the leaf holds: fn must not modify or retain it, and must not
// modify the map.
func (m *Map[T]) Range(b *Bounds, fn func(key []byte) bool) {
	var kb keyBuf
	m.t.scan(b, m.leafTail(), &kb, func(*leafHead) bool { return fn(kb.key) })
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	m.t.scan(b, m.leafTail(), nil, func(l *leafHead) bool { return eachValue(l, yield) })
}
