package art

import (
	"bytes"
	"fmt"
	"runtime"
	"slices"
	"testing"
)

// TestValueTypes makes sure that multimap.Ordered keeps values of every
// type correctly, whether its ART stores them in single-key pages of fixed-size
// values (small values without pointers, of any alignment, or a pointer), in
// single-key pages of strings, or in value overflows (every other type).
func TestValueTypes(t *testing.T) {
	type small struct {
		a uint32
		b uint16
	}
	type named struct {
		s string
		n int
	}
	ptrs := make([]int, 300)
	for _, tc := range []struct {
		name string
		mode int8 // 1: pages of fixed-size values, 3: pages of strings, -1: value overflows only
		run  func(t *testing.T) int8
	}{
		{"uint64", 1, roundTrip(func(i int) uint64 { return uint64(i) * 0x9E3779B97F4A7C15 })},
		{"int8", 1, roundTrip(func(i int) int8 { return int8(i) })},
		{"uint32", 1, roundTrip(func(i int) uint32 { return uint32(i) })},
		{"float64", 1, roundTrip(func(i int) float64 { return float64(i) / 3 })},
		{"complex128", 1, roundTrip(func(i int) complex128 { return complex(float64(i), 1) })},
		{"[2]uint64", 1, roundTrip(func(i int) [2]uint64 { return [2]uint64{uint64(i), ^uint64(i)} })},
		{"struct of uint32 and uint16", 1, roundTrip(func(i int) small { return small{uint32(i), uint16(i)} })},
		{"[3]uint64 is too large", -1, roundTrip(func(i int) [3]uint64 { return [3]uint64{uint64(i)} })},
		{"string takes single-key pages", 3, roundTrip(func(i int) string { return fmt.Sprint(i) })},
		{"pointer takes pages", 1, roundTrip(func(i int) *int { return &ptrs[i%300] })},
		{"interface is two words", -1, roundTrip(func(i int) any { return i })},
		{"struct of a string and a number", -1, roundTrip(func(i int) named { return named{fmt.Sprint(i), i} })},
		{"[5]string is too large", -1, roundTrip(func(i int) [5]string { return [5]string{fmt.Sprint(i)} })},
		{"struct{} is empty", -1, roundTrip(func(int) struct{} { return struct{}{} })},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.run(t); got != tc.mode {
				t.Fatalf("leaf mode = %d, want %d", got, tc.mode)
			}
		})
	}
}

// roundTrip returns a test that adds 300 values, made by mk, to each of a
// few keys, removes every other one, checks what is left and reports the
// map's leaf mode. Values that mk makes equal count once.
func roundTrip[T comparable](mk func(int) T) func(t *testing.T) int8 {
	return func(t *testing.T) int8 {
		var m Map[T]
		keys := [][]byte{nil, []byte("a"), []byte("ab"), bytes.Repeat([]byte("x"), 70)}
		for _, k := range keys {
			for i := range 300 {
				m.Add(k, mk(i))
			}
			for i := 0; i < 300; i += 2 {
				m.Remove(k, mk(i))
			}
		}
		runtime.GC()
		want := map[T]bool{}
		for i := range 300 {
			want[mk(i)] = true
		}
		for i := 0; i < 300; i += 2 {
			delete(want, mk(i))
		}
		for _, k := range keys {
			n := 0
			m.Each(k, func(v T) bool {
				if !want[v] {
					t.Fatalf("key %q holds %v unexpectedly", k, v)
				}
				n++
				return true
			})
			if n != len(want) {
				t.Fatalf("key %q holds %d values, want %d", k, n, len(want))
			}
		}
		return m.flat
	}
}

// TestRekeyOfOldOverflow makes sure that a value overflow of the maps that still
// hold their values in a vset.Set (values that hold a pointer, and every other
// type) keeps every value when the node above it goes away and it has to hold
// more of its key, and that it stays where it is as long as the longer key
// fits its key area: no new leaf for a delete that merges a node into its only
// leaf.
func TestRekeyOfOldOverflow(t *testing.T) {
	key := bytes.Repeat([]byte("abcdefghij"), 30) // 300 bytes
	for _, tc := range []struct {
		name         string
		keyLen, base int
		to           int
		wantInPlace  bool
	}{
		{"room in its key area", 20, 17, 6, true},
		{"beyond its key area", 20, 17, 2, false},
		{"in the largest key area", 230, 30, 1, true},
		{"up to the longest inline remainder", 300, 100, 46, true},
		{"beyond the longest inline remainder", 300, 100, 45, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			k := key[:tc.keyLen]
			l := newSetLeaf[string](k, tc.base)
			want := []string{"1", "2"}
			for _, v := range want {
				vals[string](l).Add(v)
			}
			nl := rekey[string](l, k[:tc.base-1], int(k[tc.base-1]), tc.to)
			var got []string
			vals[string](nl).Each(func(v string) bool { got = append(got, v); return true })
			slices.Sort(got)
			if !slices.Equal(got, want) {
				t.Errorf("values %v, want %v", got, want)
			}
			if nl.keyLen() != tc.keyLen || !bytes.Equal(nl.from(tc.to), k[tc.to:]) || (nl == l) != tc.wantInPlace {
				t.Errorf("leaf holds %q from %d of a key of %d bytes, in place: %v, want the key from %d on, in place: %v", nl.stored(), nl.base(), nl.keyLen(), nl == l, tc.to, tc.wantInPlace)
			}
		})
	}
}
