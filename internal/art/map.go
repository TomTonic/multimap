package art

import (
	"reflect"
	"unsafe"

	"github.com/TomTonic/multimap/internal/vset"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. A key
// with several values, or one that does not fit a page, has a leaf that holds
// its values inline (vset.Set); a short key with a single small value sits in
// a page (see page.go). The zero value is an empty map.
type Map[T comparable] struct {
	t Tree
}

// View is the values of one key: the value set of its leaf, or its single
// value in a page. The zero View stands for an absent key. It belongs to the
// map: it must not be modified and is valid only until the next write.
type View[T comparable] struct {
	set *vset.Set[T]
	one *T
}

// Found reports whether the key is present.
func (v View[T]) Found() bool { return v.set != nil || v.one != nil }

// Len returns the number of values.
func (v View[T]) Len() int {
	switch {
	case v.set != nil:
		return v.set.Len()
	case v.one != nil:
		return 1
	}
	return 0
}

// Each calls yield for every value until it returns false, and reports
// whether it ran to completion.
func (v View[T]) Each(yield func(T) bool) bool {
	switch {
	case v.set != nil:
		return v.set.Each(yield)
	case v.one != nil:
		return yield(*v.one)
	}
	return true
}

// newLeaf allocates a leaf holding a copy of key, inline or as a string. It
// captures nothing, so passing it as a newLeafFunc allocates no closure.
func newLeaf[T comparable](key []byte) *leafHead {
	if len(key) <= maxInline {
		l := &leaf[T, [16]byte]{}
		copy(l.k[:], key)
		l.init(len(key), unsafe.Offsetof(l.vals))
		return &l.leafHead
	}
	l := &leaf[T, string]{k: string(key)}
	l.init(len(key), unsafe.Offsetof(l.vals))
	return &l.leafHead
}

// vals returns the values of a leaf that was created by newLeaf[T].
func vals[T comparable](l *leafHead) *vset.Set[T] {
	return (*vset.Set[T])(unsafe.Add(unsafe.Pointer(l), l.valsOff))
}

// pageVal returns the value at position i of a page of a Map[T].
func pageVal[T comparable](p *pageHead, i int) *T {
	return (*T)(unsafe.Pointer(&p.vals()[i]))
}

// view returns the values found at n, a leaf or a page with position i.
func view[T comparable](n *header, i int) View[T] {
	if n.kind == kLeaf {
		return View[T]{set: vals[T](asLeaf(n))}
	}
	return View[T]{one: pageVal[T](asPage(n), i)}
}

// smallPlain reports whether T may be stored in pages: at most 8 bytes and
// free of pointers, so that a value fits a page's word and the page needs no
// scanning by the garbage collector.
func smallPlain[T comparable]() bool {
	var z T
	return unsafe.Sizeof(z) <= 8 && pointerFree(reflect.TypeFor[T]())
}

func pointerFree(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64:
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

// prep tells the tree about T before the first write.
func (m *Map[T]) prep() {
	if m.t.mk != nil {
		return
	}
	m.t.small = smallPlain[T]()
	m.t.mk = func(key []byte, raw uint64) *leafHead {
		l := newLeaf[T](key)
		vals[T](l).Add(*(*T)(unsafe.Pointer(&raw)))
		return l
	}
}

// Len returns the number of keys.
func (m *Map[T]) Len() int { return m.t.Len() }

// Clear removes all keys and values.
func (m *Map[T]) Clear() { m.t.Clear() }

// Add adds v to the values of key, creating the key if needed. The map keeps
// its own copy of key.
func (m *Map[T]) Add(key []byte, v T) {
	m.prep()
	var raw uint64
	if m.t.small {
		*(*T)(unsafe.Pointer(&raw)) = v
	}
	sp := m.t.upsert(key, raw, newLeaf[T])
	switch {
	case sp.leaf != nil:
		vals[T](sp.leaf).Add(v)
	case sp.created:
	case *pageVal[T](asPage(*sp.at), sp.i) != v:
		vals[T](m.t.promote(sp)).Add(v) // a second value: the key needs a leaf
	}
}

// Remove removes v from the values of key and removes the key once it holds
// no values. Absent keys and values are ignored.
func (m *Map[T]) Remove(key []byte, v T) {
	n, i := m.t.find(key)
	switch {
	case n == nil:
	case n.kind == kLeaf:
		if s := vals[T](asLeaf(n)); s.Remove(v) && s.Len() == 0 {
			m.t.remove(key)
		}
	case *pageVal[T](asPage(n), i) == v:
		m.t.remove(key)
	}
}

// RemoveKey removes key and all its values. An absent key is ignored.
func (m *Map[T]) RemoveKey(key []byte) { m.t.remove(key) }

// Values returns the values of key; the View is empty if key is absent.
func (m *Map[T]) Values(key []byte) View[T] {
	if n, i := m.t.find(key); n != nil {
		return view[T](n, i)
	}
	return View[T]{}
}

// leafTail is the offset of the last byte of the smallest leaf[T], which the
// scan touches ahead (see touchChildren). Every leaf and page is at least
// that large, and a constant offset keeps the load independent of the head.
func leafTail[T comparable]() uintptr { return unsafe.Sizeof(leaf[T, [16]byte]{}) - 1 }

// Range calls fn for every key within b, in ascending key order, until fn
// returns false. The key and the values belong to the map: fn must not modify
// or retain them, and must not modify the map.
func (m *Map[T]) Range(b *Bounds, fn func(key []byte, vals View[T]) bool) {
	var buf [8]byte
	m.t.scan(b, leafTail[T](), func(n *header, i, j int) bool {
		if n.kind == kLeaf {
			return fn(asLeaf(n).key(), view[T](n, 0))
		}
		p := asPage(n)
		h := p.heads()
		for ; i < j; i++ {
			if !fn(wordKey(h[i], int(p.klen), &buf), view[T](n, i)) {
				return false
			}
		}
		return true
	})
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false. It is Range without the
// per-key callback, for callers that need only the values.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	m.t.scan(b, leafTail[T](), func(n *header, i, j int) bool {
		if n.kind == kLeaf {
			return vals[T](asLeaf(n)).Each(yield)
		}
		vs := asPage(n).vals()
		for ; i < j; i++ {
			if !yield(*(*T)(unsafe.Pointer(&vs[i]))) {
				return false
			}
		}
		return true
	})
}
