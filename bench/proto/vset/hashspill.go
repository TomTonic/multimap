package vset

import (
	"iter"

	set3 "github.com/TomTonic/Set3"
)

// HashSpill is the first value container: InlineCap values inline, spilling
// straight into a set3 hash set. It is kept only as the baseline that Set is
// measured against (cmd/vsetcompare, cmd/memgc).
type HashSpill[T comparable] struct {
	small [InlineCap]T
	n     uint8         // number of values in small; 0 when spilled
	big   *set3.Set3[T] // non-nil once more than InlineCap values were added
}

// Len returns the number of values.
func (s *HashSpill[T]) Len() int {
	if s.big != nil {
		return int(s.big.Size())
	}
	return int(s.n)
}

// Contains reports whether v is in the set.
func (s *HashSpill[T]) Contains(v T) bool {
	if s.big != nil {
		return s.big.Contains(v)
	}
	for i := 0; i < int(s.n); i++ {
		if s.small[i] == v {
			return true
		}
	}
	return false
}

// Add inserts v and reports whether it was not present before.
func (s *HashSpill[T]) Add(v T) bool {
	if s.big != nil {
		n := s.big.Size()
		s.big.Add(v)
		return s.big.Size() != n
	}
	if s.Contains(v) {
		return false
	}
	if int(s.n) < InlineCap {
		s.small[s.n] = v
		s.n++
		return true
	}
	s.big = set3.EmptyWithCapacity[T](2 * InlineCap)
	for _, x := range s.small {
		s.big.Add(x)
	}
	s.big.Add(v)
	s.small, s.n = [InlineCap]T{}, 0
	return true
}

// Remove deletes v and reports whether it was present. A spilled set moves
// back inline only when one value is left, so that a key sitting at the
// inline boundary does not allocate on every add/remove pair.
func (s *HashSpill[T]) Remove(v T) bool {
	if s.big != nil {
		if !s.big.Remove(v) {
			return false
		}
		if s.big.Size() <= 1 {
			var small [InlineCap]T
			n := uint8(0)
			for x := range s.big.ImmutableRange() {
				small[n] = x
				n++
			}
			s.small, s.n, s.big = small, n, nil
		}
		return true
	}
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
}

// All returns an iterator over the values in unspecified order.
func (s *HashSpill[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) { s.Each(yield) }
}

// Each calls yield for every value until it returns false, and reports
// whether it ran to completion. Index structures iterating over many sets call
// this directly instead of building an iterator per set.
func (s *HashSpill[T]) Each(yield func(T) bool) bool {
	if s.big != nil {
		for v := range s.big.ImmutableRange() {
			if !yield(v) {
				return false
			}
		}
		return true
	}
	for i := 0; i < int(s.n); i++ {
		if !yield(s.small[i]) {
			return false
		}
	}
	return true
}
