package art

import (
	"unsafe"

	"github.com/TomTonic/multimap/internal/vset"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. Every
// leaf holds its key's values inline (vset.Set), so reading them costs no
// pointer chase beyond the leaf. The zero value is an empty map.
type Map[T comparable] struct {
	t Tree
}

// newLeaf allocates a leaf holding a copy of key, in the smallest size class
// that fits it. It captures nothing, so passing it as a newLeafFunc allocates
// no closure.
func newLeaf[T comparable](key []byte) *leafHead {
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
	l.init(len(key), unsafe.Offsetof(l.vals))
	return &l.leafHead
}

// newInline allocates a leaf that holds key inline in an array of type K.
func newInline[T comparable, K [16]byte | [32]byte | [48]byte | [64]byte](key []byte) *leafHead {
	l := &leaf[T, K]{}
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&l.k)), unsafe.Sizeof(l.k)), key)
	l.init(len(key), unsafe.Offsetof(l.vals))
	return &l.leafHead
}

// vals returns the values of a leaf that was created by newLeaf[T].
func vals[T comparable](l *leafHead) *vset.Set[T] {
	return (*vset.Set[T])(unsafe.Add(unsafe.Pointer(l), l.valsOff))
}

// Len returns the number of keys.
func (m *Map[T]) Len() int { return m.t.Len() }

// Clear removes all keys and values.
func (m *Map[T]) Clear() { m.t.Clear() }

// Add adds v to the values of key, creating the key if needed. The map keeps
// its own copy of key.
func (m *Map[T]) Add(key []byte, v T) {
	vals[T](m.t.upsert(key, newLeaf[T])).Add(v)
}

// Remove removes v from the values of key and removes the key once it holds
// no values. Absent keys and values are ignored.
func (m *Map[T]) Remove(key []byte, v T) {
	l := m.t.find(key)
	if l == nil {
		return
	}
	s := vals[T](l)
	if s.Remove(v) && s.Len() == 0 {
		m.t.remove(key)
	}
}

// RemoveKey removes key and all its values. An absent key is ignored.
func (m *Map[T]) RemoveKey(key []byte) { m.t.remove(key) }

// Values returns the values of key, or nil if key is absent. The set belongs
// to the map: it must not be modified and is valid only until the next write.
func (m *Map[T]) Values(key []byte) *vset.Set[T] {
	if l := m.t.find(key); l != nil {
		return vals[T](l)
	}
	return nil
}

// leafTail is the offset of the last byte of the smallest leaf[T], which the
// scan touches ahead (see touchChildren). Every leaf is at least that large,
// and a constant offset keeps the load independent of the leaf's head.
func leafTail[T comparable]() uintptr { return unsafe.Sizeof(leaf[T, [16]byte]{}) - 1 }

// Range calls fn for every key within b, in ascending key order, until fn
// returns false. The key and the set belong to the map: fn must not modify or
// retain them, and must not modify the map.
func (m *Map[T]) Range(b *Bounds, fn func(key []byte, vals *vset.Set[T]) bool) {
	m.t.scan(b, leafTail[T](), func(l *leafHead) bool { return fn(l.key(), vals[T](l)) })
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false. It is Range without the
// per-key callback, for callers that need only the values.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	m.t.scan(b, leafTail[T](), func(l *leafHead) bool { return vals[T](l).Each(yield) })
}
