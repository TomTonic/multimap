package vset

import (
	"math/rand/v2"
	"testing"
	"unsafe"
)

type container interface {
	Add(uint64) bool
	Remove(uint64) bool
	Contains(uint64) bool
	Len() int
	Each(func(uint64) bool) bool
}

// TestAgainstMap checks that the value set of one multimap key behaves like a
// plain Go set through random adds and removes, in phases that make it cross
// every representation boundary (inline, array, hash) in both directions.
func TestAgainstMap(t *testing.T) {
	for name, s := range map[string]container{"Set": &Set[uint64]{}} {
		t.Run(name, func(t *testing.T) {
			r := rand.New(rand.NewPCG(5, 6))
			ref := map[uint64]bool{}
			for i := range 400000 {
				domain := []int{5, 40, 300, 80}[i/10000%4] // small, array, hash, back down
				v := uint64(r.IntN(domain))
				grow := r.IntN(100) < []int{50, 50, 60, 40}[i/10000%4]
				if grow {
					if s.Add(v) == ref[v] {
						t.Fatalf("step %d: Add(%d) result wrong", i, v)
					}
					ref[v] = true
				} else {
					if s.Remove(v) != ref[v] {
						t.Fatalf("step %d: Remove(%d) result wrong", i, v)
					}
					delete(ref, v)
				}
				if s.Len() != len(ref) {
					t.Fatalf("step %d: Len = %d, want %d", i, s.Len(), len(ref))
				}
				if i%997 == 0 {
					check(t, s, ref)
				}
			}
			check(t, s, ref)
		})
	}
}

func check(t *testing.T, s container, ref map[uint64]bool) {
	t.Helper()
	seen := map[uint64]bool{}
	s.Each(func(v uint64) bool {
		if !ref[v] || seen[v] {
			t.Fatalf("Each yielded %d unexpectedly", v)
		}
		seen[v] = true
		return true
	})
	if len(seen) != len(ref) {
		t.Fatalf("Each yielded %d values, want %d", len(seen), len(ref))
	}
	for v := range ref {
		if !s.Contains(v) {
			t.Fatalf("Contains(%d) = false", v)
		}
	}
}

// TestRepresentations checks that sizes map to the documented representation.
func TestRepresentations(t *testing.T) {
	var s Set[uint64]
	mode := func() string {
		switch {
		case s.ext == nil:
			return "inline"
		case s.cap > 0:
			return "array"
		}
		return "hash"
	}
	for i := uint64(1); i <= ArrayMax+1; i++ {
		s.Add(i)
		want := map[bool]string{true: "inline", false: "array"}[i <= InlineCap]
		if i > ArrayMax {
			want = "hash"
		}
		if mode() != want {
			t.Fatalf("with %d values: %s, want %s", i, mode(), want)
		}
	}
	for i := uint64(ArrayMax + 1); i > 0; i-- {
		s.Remove(i)
		n := i - 1
		want := "hash"
		switch {
		case n < InlineCap:
			want = "inline"
		case n <= ArrayBack:
			want = "array"
		}
		if mode() != want {
			t.Fatalf("down to %d values: %s, want %s", n, mode(), want)
		}
	}
}

// TestLayout guards the size the ART leaf layout depends on.
func TestLayout(t *testing.T) {
	if got := unsafe.Sizeof(Set[uint64]{}); got != 40 {
		t.Fatalf("Set[uint64] is %d bytes, want 40", got)
	}
}

// TestIteration makes sure that reading the values of a multimap key yields
// every value exactly once and stops as soon as the caller stops, however many
// values the key holds. It covers All and Each of the value set in each of its
// representations (empty, inline, array, hash) and checks a full pass and a
// pass that the caller ends after the first value.
func TestIteration(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    int
	}{
		{"an empty set yields nothing", 0},
		{"a full inline set yields every value", InlineCap},
		{"a full array set yields every value", ArrayMax},
		{"a hash set yields every value", ArrayMax + 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var s Set[uint64]
			for i := range tc.n {
				s.Add(uint64(i))
			}
			seen := map[uint64]bool{}
			for v := range s.All() {
				if seen[v] || v >= uint64(tc.n) {
					t.Fatalf("All yielded %d twice or unexpectedly", v)
				}
				seen[v] = true
			}
			if len(seen) != tc.n {
				t.Fatalf("All yielded %d values, want %d", len(seen), tc.n)
			}
			calls := 0
			done := s.Each(func(uint64) bool { calls++; return false })
			if calls != min(1, tc.n) || done != (tc.n == 0) {
				t.Fatalf("Each after stopping at once: %d calls, completed %v", calls, done)
			}
			calls = 0
			for range s.All() {
				calls++
				break
			}
			if calls != min(1, tc.n) {
				t.Fatalf("All after a break: %d calls", calls)
			}
		})
	}
}
