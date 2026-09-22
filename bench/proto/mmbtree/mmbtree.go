// Package mmbtree implements the multimap API on tidwall/btree.Map in the two
// ways that library supports for mutable values, both with the same value
// container (vset.Set) as the ART prototype:
//
//   - Inline stores vset.Set by value in the B-tree's node arrays. Reads and
//     scans touch no extra memory, but btree.Map has no in-place update, so a
//     mutation is Get, modify, Set (two descents), as the README describes.
//   - Ptr stores *vset.Set. A mutation is one descent plus a pointer write, but
//     every value read follows a pointer.
//
// Lookups use unsafe.String to view the []byte key as a string without
// allocating (btree.Map does not retain it). Inserting a new key clones it,
// because the multimap API takes mutable []byte keys.
package mmbtree

import (
	"iter"
	"unsafe"

	"github.com/tidwall/btree"

	"github.com/TomTonic/multimap/bench/proto/vset"
)

func view(b []byte) string { return unsafe.String(unsafe.SliceData(b), len(b)) }

// Inline is the B-tree multimap with values stored inline.
type Inline[T comparable] struct {
	m btree.Map[string, vset.Set[T]]
}

// Len returns the number of keys.
func (m *Inline[T]) Len() int { return m.m.Len() }

// AddValue adds v to the values of key, creating the key if needed.
func (m *Inline[T]) AddValue(key []byte, v T) {
	s, ok := m.m.Get(view(key))
	if !s.Add(v) && ok {
		return // unchanged, no need to write back
	}
	if ok {
		m.m.Set(view(key), s) // key already stored: the view is not retained
	} else {
		m.m.Set(string(key), s)
	}
}

// RemoveValue removes v from the values of key.
func (m *Inline[T]) RemoveValue(key []byte, v T) {
	s, ok := m.m.Get(view(key))
	if ok && s.Remove(v) {
		m.m.Set(view(key), s)
	}
}

// ValuesFor iterates over the values of key.
func (m *Inline[T]) ValuesFor(key []byte) iter.Seq[T] {
	return func(yield func(T) bool) {
		if s, ok := m.m.Get(view(key)); ok {
			s.Each(yield)
		}
	}
}

// ValuesBetween iterates over the values of all keys in [from, to], in key
// order.
func (m *Inline[T]) ValuesBetween(from, to []byte) iter.Seq[T] {
	return func(yield func(T) bool) {
		hi := view(to)
		m.m.Ascend(view(from), func(k string, s vset.Set[T]) bool {
			if k > hi {
				return false
			}
			return s.Each(yield)
		})
	}
}

// Ptr is the B-tree multimap with values behind a pointer.
type Ptr[T comparable] struct {
	m btree.Map[string, *vset.Set[T]]
}

// Len returns the number of keys.
func (m *Ptr[T]) Len() int { return m.m.Len() }

// AddValue adds v to the values of key, creating the key if needed.
func (m *Ptr[T]) AddValue(key []byte, v T) {
	s, ok := m.m.Get(view(key))
	if !ok {
		s = &vset.Set[T]{}
		m.m.Set(string(key), s)
	}
	s.Add(v)
}

// RemoveValue removes v from the values of key.
func (m *Ptr[T]) RemoveValue(key []byte, v T) {
	if s, ok := m.m.Get(view(key)); ok {
		s.Remove(v)
	}
}

// ValuesFor iterates over the values of key.
func (m *Ptr[T]) ValuesFor(key []byte) iter.Seq[T] {
	return func(yield func(T) bool) {
		if s, ok := m.m.Get(view(key)); ok {
			s.Each(yield)
		}
	}
}

// ValuesBetween iterates over the values of all keys in [from, to], in key
// order.
func (m *Ptr[T]) ValuesBetween(from, to []byte) iter.Seq[T] {
	return func(yield func(T) bool) {
		hi := view(to)
		m.m.Ascend(view(from), func(k string, s *vset.Set[T]) bool {
			if k > hi {
				return false
			}
			return s.Each(yield)
		})
	}
}
