package art

import (
	"reflect"
	"unsafe"

	"github.com/TomTonic/multimap/internal/vset"
)

// Map is a multimap from byte-string keys to sets of T on top of Tree. When T
// is small and pointer-free, keys of up to maxPageKey bytes sit in pages with
// their values (see page.go); every other key has a leaf that holds its
// values in a vset.Set. The zero value is an empty map.
type Map[T comparable] struct {
	t Tree
}

// View is the values of one key: the value set of its leaf or its page, or
// its values held inline in a page, as raw words. The zero View stands for an
// absent key. It belongs to the map: it must not be modified and is valid
// only until the next write.
type View[T comparable] struct {
	set *vset.Set[T]
	raw []uint64
}

// Found reports whether the key is present.
func (v View[T]) Found() bool { return v.set != nil || v.raw != nil }

// Len returns the number of values.
func (v View[T]) Len() int {
	if v.set != nil {
		return v.set.Len()
	}
	return len(v.raw)
}

// Each calls yield for every value until it returns false, and reports
// whether it ran to completion.
func (v View[T]) Each(yield func(T) bool) bool {
	if v.set != nil {
		return v.set.Each(yield)
	}
	return eachRaw(v.raw, yield)
}

// eachRaw calls yield for every raw value until it returns false.
func eachRaw[T comparable](raw []uint64, yield func(T) bool) bool {
	for i := range raw {
		if !yield(*(*T)(unsafe.Pointer(&raw[i]))) {
			return false
		}
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

// pageVal returns the value at position i of a U8-1 page of a Map[T].
func pageVal[T comparable](p *pageHead, i int) *T {
	return (*T)(unsafe.Pointer(&p.vals()[i]))
}

// view returns the values found at n, a leaf or a page with position i.
func view[T comparable](n *header, i int) View[T] {
	switch n.kind {
	case kLeaf:
		return View[T]{set: vals[T](asLeaf(n))}
	case kPage:
		return View[T]{raw: asPage(n).vals()[i : i+1]}
	}
	p := asPage(n)
	l := p.layout()
	off, cnt, e := p.runIn(&l, i)
	if e >= 0 {
		return View[T]{set: (*vset.Set[T])(p.extsIn(&l)[e])}
	}
	return View[T]{raw: p.nvalsIn(&l)[off : off+cnt]}
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
	m.t.mk = func(it item) *leafHead {
		l := newLeaf[T](it.key)
		if it.set != nil {
			*vals[T](l) = *(*vset.Set[T])(it.set)
			return l
		}
		eachRaw(it.vals, func(v T) bool { vals[T](l).Add(v); return true })
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
	case (*sp.at).kind == kPage:
		if *pageVal[T](asPage(*sp.at), sp.i) != v {
			m.t.addToSingle(sp, raw)
		}
	default:
		m.addToPageN(sp, v, raw)
	}
}

// addToPageN adds v to the values of the key at sp, in a U8-n page: to its
// external set, inline, or, as its (inlineMax+1)-th value, to a new external
// set that takes over its inline values.
func (m *Map[T]) addToPageN(sp spot, v T, raw uint64) {
	p := asPage(*sp.at)
	off, n, e := p.run(sp.i)
	if e >= 0 {
		(*vset.Set[T])(p.exts()[e]).Add(v)
		return
	}
	run := p.nvals()[off : off+n]
	if !eachRaw(run, func(x T) bool { return x != v }) {
		return // already there
	}
	if n < inlineMax {
		m.t.addInline(sp, raw)
		return
	}
	s := &vset.Set[T]{}
	eachRaw(run, func(x T) bool { s.Add(x); return true })
	s.Add(v)
	m.t.externalize(sp, unsafe.Pointer(s))
}

// Remove removes v from the values of key and removes the key once it holds
// no values. Absent keys and values are ignored.
func (m *Map[T]) Remove(key []byte, v T) {
	n, i := m.t.find(key)
	if n == nil {
		return
	}
	var s *vset.Set[T]
	switch n.kind {
	case kLeaf:
		s = vals[T](asLeaf(n))
	case kPage:
		if *pageVal[T](asPage(n), i) == v {
			m.t.remove(key)
		}
		return
	default:
		p := asPage(n)
		off, cnt, e := p.run(i)
		if e >= 0 {
			s = (*vset.Set[T])(p.exts()[e])
			break
		}
		for k := range cnt {
			if *(*T)(unsafe.Pointer(&p.nvals()[off+k])) == v {
				if cnt == 1 {
					m.t.remove(key)
				} else {
					p.nRemoveVal(i, k)
				}
				return
			}
		}
		return
	}
	if s.Remove(v) && s.Len() == 0 {
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
	var buf keyBuf
	m.t.scan(b, leafTail[T](), func(n *header, i, j int) bool {
		if n.kind == kLeaf {
			return fn(asLeaf(n).key(), view[T](n, 0))
		}
		p := asPage(n)
		for k := i; k < j; k++ {
			if !fn(p.key(k, &buf), view[T](n, k)) {
				return false
			}
		}
		return true
	})
}

// RangeValues calls yield for every value of every key within b, key by key
// in ascending key order, until yield returns false. It is Range without the
// per-key callback, for callers that need only the values; it walks a page's
// values directly.
func (m *Map[T]) RangeValues(b *Bounds, yield func(T) bool) {
	m.t.scan(b, leafTail[T](), func(n *header, i, j int) bool {
		switch n.kind {
		case kLeaf:
			return vals[T](asLeaf(n)).Each(yield)
		case kPage:
			return eachRaw(asPage(n).vals()[i:j], yield)
		}
		p := asPage(n)
		off, _, _ := p.run(i)
		vs, ex := p.nvals(), p.exts()
		for _, c := range p.cnts()[i:j] {
			if c >= extBit {
				if !(*vset.Set[T])(ex[c&^extBit]).Each(yield) {
					return false
				}
				continue
			}
			if !eachRaw(vs[off:off+int(c)], yield) {
				return false
			}
			off += int(c)
		}
		return true
	})
}
