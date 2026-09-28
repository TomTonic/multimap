package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/vset"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. Every
// leaf holds its key's values: a flat leaf right after the key (see flat.go),
// a set leaf in a vset.Set. Reading them costs no pointer chase beyond the
// leaf until a key holds more values than a flat leaf of 512 bytes. The zero
// value is an empty map.
type Map[T comparable] struct {
	t    Tree
	flat int8 // 1: T takes flat leaves, -1: set leaves only, 0: not decided yet
}

// newSetLeaf allocates a set leaf holding a copy of key, in the smallest size
// class that fits it. It captures nothing, so passing it as a newLeafFunc
// allocates no closure.
func newSetLeaf[T comparable](key []byte) *leafHead {
	switch n := len(key); {
	case n <= 16:
		return newInline[T, [16]byte](key)
	case n <= 32:
		return newInline[T, [32]byte](key)
	case n <= 48:
		return newInline[T, [48]byte](key)
	case n <= maxInline:
		return newInline[T, [64]byte](key)
	}
	l := &leaf[T, string]{k: string(key)}
	l.kind, l.klen = kLeaf, longKey
	return &l.leafHead
}

// newInline allocates a set leaf that holds key inline in an array of type K.
func newInline[T comparable, K [16]byte | [32]byte | [48]byte | [64]byte](key []byte) *leafHead {
	l := &leaf[T, K]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), key)
	l.kind, l.klen = kLeaf, uint8(len(key))
	return &l.leafHead
}

// Offsets of the value set in the set leaf of each key size class. They do
// not depend on T: a vset.Set holds a pointer, so it is always 8-aligned.
var (
	setOff16  = unsafe.Offsetof(leaf[struct{}, [16]byte]{}.vals)
	setOff32  = unsafe.Offsetof(leaf[struct{}, [32]byte]{}.vals)
	setOff48  = unsafe.Offsetof(leaf[struct{}, [48]byte]{}.vals)
	setOff64  = unsafe.Offsetof(leaf[struct{}, [64]byte]{}.vals)
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
	case k <= maxInline:
		off = setOff64
	}
	return (*vset.Set[T])(unsafe.Add(unsafe.Pointer(l), off))
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
	if l.cls == 0 {
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
	if l.cls != 0 {
		switch c, empty := flatRemove(l, v); {
		case empty:
			m.t.remove(key)
		case c != 0:
			*m.t.findSlot(key) = leafHdr(resize[T](l, c))
		}
		return
	}
	s := vals[T](l)
	switch {
	case !s.Remove(v):
	case s.Len() == 0:
		m.t.remove(key)
	case m.flat > 0:
		if c := unspillClass[T](l); c != 0 {
			*m.t.findSlot(key) = leafHdr(unspill[T](l, c))
		}
	}
}

// RemoveKey removes key and all its values. An absent key is ignored.
func (m *Map[T]) RemoveKey(key []byte) { m.t.remove(key) }

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
	if l.cls == 0 {
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
// returns false. The key belongs to the map: fn must not modify or retain
// it, and must not modify the map.
func (m *Map[T]) Range(b *Bounds, fn func(key []byte) bool) {
	m.t.scan(b, m.leafTail(), func(l *leafHead) bool { return fn(l.key()) })
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	m.t.scan(b, m.leafTail(), func(l *leafHead) bool { return eachValue(l, yield) })
}
