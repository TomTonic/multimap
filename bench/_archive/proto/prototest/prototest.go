// Package prototest checks an ordered byte-key map implementation against a
// sorted reference, so both ART prototypes are verified by the same test.
package prototest

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"sort"
	"testing"
)

// Map is the surface both prototypes implement.
type Map interface {
	Put(key []byte, val uint32)
	Get(key []byte) (uint32, bool)
	Scan(from []byte, fn func(key []byte, val uint32) bool)
	Len() int
}

// KeySets returns named key corpora that exercise the structural edge cases:
// keys that are prefixes of other keys, the empty key, keys longer than both
// the inline prefix (8) and the inline leaf key (16), zero bytes, dense and
// sparse integers, and shared string prefixes.
func KeySets() map[string][][]byte {
	r := rand.New(rand.NewPCG(1, 2))
	sets := map[string][][]byte{}

	var u64 [][]byte
	for range 20000 {
		u64 = append(u64, binary.BigEndian.AppendUint64(nil, r.Uint64()))
	}
	sets["u64-random"] = u64

	var dense [][]byte
	for i := range uint64(70000) {
		dense = append(dense, binary.BigEndian.AppendUint64(nil, i))
	}
	sets["u64-dense"] = dense

	var prefixes [][]byte
	base := []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	for i := 0; i <= len(base); i++ {
		prefixes = append(prefixes, append([]byte(nil), base[:i]...))
		prefixes = append(prefixes, append(append([]byte(nil), base[:i]...), 0))
		prefixes = append(prefixes, append(append([]byte(nil), base[:i]...), 0xff, 'x'))
	}
	sets["prefix-chains"] = prefixes

	var strs [][]byte
	words := []string{"user", "users", "item", "items", "order", "a", "", "product", "productcatalogue", "x"}
	for range 30000 {
		k := fmt.Sprintf("%s/%s/%s/%d", words[r.IntN(len(words))], words[r.IntN(len(words))],
			words[r.IntN(len(words))], r.IntN(300))
		strs = append(strs, []byte(k))
	}
	sets["strings"] = strs

	// wide fan-out below a long compressed path, then keys that split that path:
	// exercises optimistic prefixes above the 57/52- and 256-way nodes.
	var wide [][]byte
	long := []byte("a-compressed-path-longer-than-sixteen-bytes/")
	for _, fan := range []int{40, 256} {
		p := append(append([]byte(nil), long...), byte(fan))
		for b := range fan {
			wide = append(wide, append(append(append([]byte(nil), p...), byte(b)), "tail"...))
		}
	}
	wide = append(wide, long[:20], append(append([]byte(nil), long[:30]...), 'Z'))
	sets["long-prefix-wide"] = wide

	var random [][]byte
	for range 20000 {
		k := make([]byte, r.IntN(40))
		for i := range k {
			k[i] = byte(r.IntN(4)) // tiny alphabet -> deep, long shared paths
		}
		random = append(random, k)
	}
	sets["random-small-alphabet"] = random
	return sets
}

// Check inserts keys into m and verifies Get, Len and Scan against a reference.
func Check(t *testing.T, newMap func() Map) {
	for name, keys := range KeySets() {
		t.Run(name, func(t *testing.T) {
			m := newMap()
			ref := map[string]uint32{}
			for i, k := range keys {
				m.Put(k, uint32(i))
				ref[string(k)] = uint32(i)
			}
			if m.Len() != len(ref) {
				t.Fatalf("Len = %d, want %d", m.Len(), len(ref))
			}
			for k, v := range ref {
				got, ok := m.Get([]byte(k))
				if !ok || got != v {
					t.Fatalf("Get(%q) = %d,%v want %d", k, got, ok, v)
				}
			}
			sorted := make([]string, 0, len(ref))
			for k := range ref {
				sorted = append(sorted, k)
			}
			sort.Strings(sorted)

			r := rand.New(rand.NewPCG(3, 4))
			for i := range 300 {
				// misses: mutate a present key
				k := []byte(sorted[r.IntN(len(sorted))])
				miss := append(append([]byte(nil), k...), byte(r.IntN(256)))
				if len(k) > 0 && i%2 == 0 {
					miss = append([]byte(nil), k[:len(k)-1]...)
				}
				if _, present := ref[string(miss)]; !present {
					if _, ok := m.Get(miss); ok {
						t.Fatalf("Get(%q) found a key that was never inserted", miss)
					}
				}
			}

			froms := [][]byte{nil, {}, {0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}}
			for range 300 {
				k := []byte(sorted[r.IntN(len(sorted))])
				froms = append(froms, k)
				if len(k) > 0 {
					j := r.IntN(len(k))
					m2 := append([]byte(nil), k...)
					m2[j] += byte(r.IntN(3)) - 1
					froms = append(froms, m2, k[:j])
				}
			}
			for _, from := range froms {
				start := sort.SearchStrings(sorted, string(from))
				want := sorted[start:min(start+50, len(sorted))]
				var got []string
				m.Scan(from, func(k []byte, v uint32) bool {
					if ref[string(k)] != v {
						t.Fatalf("Scan value mismatch for %q", k)
					}
					got = append(got, string(k))
					return len(got) < 50
				})
				if len(got) != len(want) {
					t.Fatalf("Scan(%q): got %d keys, want %d\n got=%q\nwant=%q", from, len(got), len(want), got, want)
				}
				for i := range got {
					if got[i] != want[i] {
						t.Fatalf("Scan(%q)[%d] = %q, want %q", from, i, got[i], want[i])
					}
				}
			}
			// full ordered iteration
			n := 0
			var prev []byte
			m.Scan(nil, func(k []byte, _ uint32) bool {
				if n > 0 && bytes.Compare(prev, k) >= 0 {
					t.Fatalf("Scan not ascending: %q then %q", prev, k)
				}
				prev = append(prev[:0], k...)
				n++
				return true
			})
			if n != len(ref) {
				t.Fatalf("full Scan visited %d keys, want %d", n, len(ref))
			}
		})
	}
}
