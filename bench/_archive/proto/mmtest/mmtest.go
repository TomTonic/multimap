// Package mmtest checks a multimap implementation against a reference built
// from Go maps, so every multimap prototype is verified by the same test.
package mmtest

import (
	"iter"
	"math/rand/v2"
	"slices"
	"sort"
	"testing"

	"github.com/TomTonic/multimap/bench/proto/prototest"
)

// Multimap is the API all multimap prototypes implement.
type Multimap interface {
	AddValue(key []byte, v uint64)
	RemoveValue(key []byte, v uint64)
	ValuesFor(key []byte) iter.Seq[uint64]
	ValuesBetween(from, to []byte) iter.Seq[uint64]
	Len() int
}

// Check fills m with random values per key, removes some, and compares point
// and range reads against the reference. Range results must come in key order
// (the values of one key in any order).
func Check(t *testing.T, newMM func() Multimap) {
	for name, keys := range prototest.KeySets() {
		t.Run(name, func(t *testing.T) {
			m := newMM()
			ref := map[string]map[uint64]bool{}
			r := rand.New(rand.NewPCG(7, 8))
			for _, k := range keys {
				n := 1 + r.IntN(6)
				if r.IntN(20) == 0 {
					n = 40 // spilled sets
				}
				for range n {
					v := uint64(r.IntN(50))
					m.AddValue(k, v)
					if ref[string(k)] == nil {
						ref[string(k)] = map[uint64]bool{}
					}
					ref[string(k)][v] = true
				}
				if r.IntN(3) == 0 { // remove present and absent values
					v := uint64(r.IntN(50))
					m.RemoveValue(k, v)
					delete(ref[string(k)], v)
				}
			}
			if m.Len() != len(ref) {
				t.Fatalf("Len = %d, want %d", m.Len(), len(ref))
			}
			for k, want := range ref {
				got := slices.Sorted(m.ValuesFor([]byte(k)))
				if !slices.Equal(got, sortedSet(want)) {
					t.Fatalf("ValuesFor(%q) = %v, want %v", k, got, sortedSet(want))
				}
			}
			sorted := make([]string, 0, len(ref))
			for k := range ref {
				sorted = append(sorted, k)
			}
			sort.Strings(sorted)
			for i := range 600 {
				a, b := sorted[r.IntN(len(sorted))], sorted[r.IntN(len(sorted))]
				switch i % 6 {
				case 1:
					a += "\x00" // bounds that are not keys
				case 2:
					b = b[:r.IntN(len(b)+1)] // bounds that are prefixes of keys
				case 3:
					a = a[:r.IntN(len(a)+1)]
				case 4:
					b = a // single key
				case 5:
					b += "\xff\xff"
				}
				checkRange(t, m, ref, sorted, a, b)
			}
			checkRange(t, m, ref, sorted, "", "\xff\xff\xff\xff\xff\xff\xff\xff\xff") // everything
		})
	}
}

func checkRange(t *testing.T, m Multimap, ref map[string]map[uint64]bool, sorted []string, from, to string) {
	var want []uint64
	var perKey [][]uint64
	for _, k := range sorted {
		if k >= from && k <= to {
			s := sortedSet(ref[k])
			want = append(want, s...)
			perKey = append(perKey, s)
		}
	}
	got := slices.Collect(m.ValuesBetween([]byte(from), []byte(to)))
	if len(got) != len(want) {
		t.Fatalf("ValuesBetween(%q, %q) yielded %d values, want %d", from, to, len(got), len(want))
	}
	// values of each key form a contiguous group in key order
	i := 0
	for _, s := range perKey {
		g := slices.Sorted(slices.Values(got[i : i+len(s)]))
		if !slices.Equal(g, s) {
			t.Fatalf("ValuesBetween(%q, %q): group %v, want %v", from, to, g, s)
		}
		i += len(s)
	}
}

func sortedSet(s map[uint64]bool) []uint64 {
	out := make([]uint64, 0, len(s))
	for v := range s {
		out = append(out, v)
	}
	slices.Sort(out)
	return out
}
