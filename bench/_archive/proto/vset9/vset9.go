// Package vset9 is package vset with 9 inline values instead of 3, generated
// for the leaf-size comparison (a Set[uint64] is 88 bytes, so an ART leaf is
// 128). Go generics cannot take the capacity as a parameter.
//
// The original package documentation follows.
//
// Package vset is the value container of one multimap key.
//
// Most keys of a multimap hold one or a few values, a minority hold dozens and
// a few hold hundreds. The container has one representation for each:
//
//   - inline: up to InlineCap values inside the Set itself, so inside the
//     owning ART leaf or B-tree item. No allocation, no pointer chase.
//   - array: up to ArrayMax values in one plain, unsorted array, searched
//     linearly. For pointer-free T it is a noscan object the GC never scans.
//     It is 2-4x smaller than a hash set of the same size.
//   - hash: a set3 hash set beyond that, where linear search would cost too
//     much.
//
// Shrinking uses hysteresis (hash back to array at ArrayBack values, array
// back to inline below InlineCap), so a key sitting at a boundary does not
// reallocate on every add/remove pair.
//
// Every index structure in the benchmark uses this container, so they differ
// only in the index.
package vset9

import (
	"iter"
	"slices"
	"unsafe"

	set3 "github.com/TomTonic/Set3"
)

const (
	// InlineCap 9 makes a Set[uint64] 88 bytes and the ART leaf 128.
	InlineCap = 9
	// ArrayMax is the largest set kept as an array. A linear scan over 64
	// uint64 values reads 512 contiguous bytes, which is about what a cold hash
	// lookup costs in cache misses.
	ArrayMax = 64
	// ArrayBack is where a hash set turns back into an array.
	ArrayBack = ArrayMax / 2
)

// Set holds the values of one key. The zero value is an empty set.
//
// The representation is encoded in ext and cap: ext == nil means inline,
// cap > 0 means ext points to the first element of an array with that
// capacity, and cap == 0 with ext != nil means ext is a *set3.Set3[T].
type Set[T comparable] struct {
	small [InlineCap]T
	n     uint32 // number of values, in every representation
	cap   uint32 // capacity of the array; 0 when inline or hashed
	ext   unsafe.Pointer
}

func (s *Set[T]) array() []T {
	return unsafe.Slice((*T)(s.ext), s.cap)[:s.n]
}

func (s *Set[T]) hash() *set3.Set3[T] { return (*set3.Set3[T])(s.ext) }

func (s *Set[T]) setArray(a []T) {
	a = a[:cap(a)]
	s.ext, s.cap = unsafe.Pointer(unsafe.SliceData(a)), uint32(len(a))
}

// Len returns the number of values.
func (s *Set[T]) Len() int { return int(s.n) }

// Contains reports whether v is in the set.
func (s *Set[T]) Contains(v T) bool {
	switch {
	case s.ext == nil:
		for i := 0; i < int(s.n); i++ {
			if s.small[i] == v {
				return true
			}
		}
		return false
	case s.cap > 0:
		return slices.Contains(s.array(), v)
	default:
		return s.hash().Contains(v)
	}
}

// Add inserts v and reports whether it was not present before.
func (s *Set[T]) Add(v T) bool {
	switch {
	case s.ext == nil:
		if s.Contains(v) {
			return false
		}
		if s.n < InlineCap {
			s.small[s.n] = v
		} else {
			a := make([]T, 0, 2*InlineCap)
			s.setArray(append(append(a, s.small[:]...), v))
			s.small = [InlineCap]T{}
		}
	case s.cap > 0:
		if s.Contains(v) {
			return false
		}
		if s.n == ArrayMax {
			h := set3.EmptyWithCapacity[T](2 * ArrayMax)
			h.AddAllFromArray(s.array())
			h.Add(v)
			s.ext, s.cap = unsafe.Pointer(h), 0
		} else {
			// append grows into Go's size classes; it only reallocates when
			// the array is full.
			s.setArray(append(s.array(), v))
		}
	default:
		h := s.hash()
		if h.Contains(v) {
			return false
		}
		h.Add(v)
	}
	s.n++
	return true
}

// Remove deletes v and reports whether it was present.
func (s *Set[T]) Remove(v T) bool {
	switch {
	case s.ext == nil:
		for i := 0; i < int(s.n); i++ {
			if s.small[i] == v {
				s.n--
				s.small[i] = s.small[s.n]
				var zero T
				s.small[s.n] = zero
				return true
			}
		}
		return false
	case s.cap > 0:
		a := s.array()
		for i, x := range a {
			if x == v {
				last := len(a) - 1
				a[i] = a[last]
				var zero T
				a[last] = zero
				s.n--
				if s.n < InlineCap {
					copy(s.small[:], a[:s.n])
					s.ext, s.cap = nil, 0
				}
				return true
			}
		}
		return false
	default:
		h := s.hash()
		if !h.Remove(v) {
			return false
		}
		s.n--
		if s.n <= ArrayBack {
			a := make([]T, 0, ArrayBack+ArrayBack/2)
			for x := range h.ImmutableRange() {
				a = append(a, x)
			}
			s.setArray(a) // setArray keeps the spare capacity
		}
		return true
	}
}

// All returns an iterator over the values in unspecified order.
func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) { s.Each(yield) }
}

// Each calls yield for every value until it returns false, and reports
// whether it ran to completion. Index structures iterating over many sets call
// this directly instead of building an iterator per set.
func (s *Set[T]) Each(yield func(T) bool) bool {
	switch {
	case s.ext == nil:
		for i := 0; i < int(s.n); i++ {
			if !yield(s.small[i]) {
				return false
			}
		}
	case s.cap > 0:
		for _, x := range s.array() {
			if !yield(x) {
				return false
			}
		}
	default:
		for x := range s.hash().ImmutableRange() {
			if !yield(x) {
				return false
			}
		}
	}
	return true
}
