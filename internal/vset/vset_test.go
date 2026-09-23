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
