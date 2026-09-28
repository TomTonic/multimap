package art

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"slices"
	"testing"
)

// TestFlatLeaves makes sure that a key of multimap.Ordered keeps exactly its
// values while their number grows from one to hundreds and shrinks back,
// whatever the key's length. It covers the flat leaves of the ART behind
// Ordered, which hold small pointer-free values right after the key and move
// through size classes as values arrive: after every step the key must hold
// exactly the values added and not removed, a flat leaf must sit in a class
// that fits its values, a key with more values than the largest class holds
// must have spilled into a set leaf, and it must turn flat again once half
// of the largest class would do. The values must also survive a garbage
// collection, since the leaves are memory the collector never scans.
func TestFlatLeaves(t *testing.T) {
	for _, klen := range []int{0, 8, 26, 36, 66, 200, maxFlatKey, maxFlatKey + 1, 300} {
		t.Run(fmt.Sprintf("key of %d bytes", klen), func(t *testing.T) {
			key := bytes.Repeat([]byte{'k'}, klen)
			var m Map[uint64]
			m.Add([]byte("neighbour"), 1) // the key's leaf lives below a node
			maxCap := flatCap[uint64](uint8(len(flatSizes)-1), klen)
			var want []uint64
			check := func() {
				t.Helper()
				runtime.GC()
				got := valuesOf(&m, key)
				slices.Sort(got)
				if !slices.Equal(got, want) {
					t.Fatalf("key holds %v, want %v", got, want)
				}
				if len(want) == 0 {
					return
				}
				l := m.t.find(key)
				switch {
				case klen > maxFlatKey:
					if l.cls() != 0 {
						t.Fatalf("a key of %d bytes got a flat leaf", klen)
					}
				case l.cls() != 0:
					if int(l.n) != len(want) || len(want) > flatCap[uint64](l.cls(), klen) {
						t.Fatalf("flat leaf of class %d holds %d values, want %d", l.cls(), l.n, len(want))
					}
				case len(want) <= maxCap/2:
					t.Fatalf("set leaf with %d values, want a flat leaf again", len(want))
				}
			}
			for v := range uint64(2*maxCap + 10) {
				m.Add(key, v)
				m.Add(key, v) // a duplicate changes nothing
				want = append(want, v)
				check()
				if l := m.t.find(key); klen <= maxFlatKey && len(want) > maxCap && l.cls() != 0 {
					t.Fatalf("%d values in a flat leaf, the largest holds %d", len(want), maxCap)
				}
			}
			for len(want) > 0 {
				v := want[len(want)/2]
				m.Remove(key, v)
				m.Remove(key, v) // an absent value changes nothing
				want = slices.DeleteFunc(want, func(x uint64) bool { return x == v })
				check()
			}
			if m.Has(key) || m.Len() != 1 {
				t.Fatalf("key still present after removing all its values")
			}
			checkInvariants(t, &m.t)
		})
	}
}

// TestFlatLeafHysteresis makes sure that a key whose number of values
// hovers at a size class boundary does not move its leaf on every change.
// In the ART behind multimap.Ordered a flat leaf grows into a class that
// holds twice its values when it is full, and shrinks only once the smaller
// class would be half empty.
func TestFlatLeafHysteresis(t *testing.T) {
	var m Map[uint64]
	key := []byte("hovering")
	full := flatCap[uint64](minGrown, len(key)) // a cache-line leaf holds this many
	for v := range uint64(full + 1) {
		m.Add(key, v)
	}
	l := m.t.find(key)
	if want := flatClass[uint64](len(key), 2*full); l.cls() != want {
		t.Fatalf("leaf in class %d after %d values, want class %d, which holds twice as many", l.cls(), full+1, want)
	}
	for range 10 {
		m.Remove(key, uint64(full))
		m.Add(key, uint64(full))
		if m.t.find(key) != l {
			t.Fatalf("hovering at %d values moved the leaf", full)
		}
	}

	// A key that once grew keeps a cache line while it hovers between one
	// and a few values; its first leaf of 32 bytes holds 3.
	few := []byte("few")
	for v := range uint64(4) { // one more than the smallest leaf holds
		m.Add(few, v)
	}
	l = m.t.find(few)
	for range 10 {
		for v := range uint64(3) {
			m.Remove(few, v+1)
		}
		for v := range uint64(3) {
			m.Add(few, v+1)
		}
		if m.t.find(few) != l || l.cls() != minGrown {
			t.Fatalf("hovering between 1 and 4 values moved the leaf or left the cache line")
		}
	}
}

// TestFlatValueTypes makes sure that multimap.Ordered keeps values of every
// type correctly, whether its ART stores them in flat leaves (small values
// without pointers, of any alignment) or in set leaves (everything else).
func TestFlatValueTypes(t *testing.T) {
	type small struct {
		a uint32
		b uint16
	}
	for _, tc := range []struct {
		name string
		flat bool
		run  func(t *testing.T) int8
	}{
		{"uint64", true, roundTrip(func(i int) uint64 { return uint64(i) * 0x9E3779B97F4A7C15 })},
		{"int8", true, roundTrip(func(i int) int8 { return int8(i) })},
		{"uint32", true, roundTrip(func(i int) uint32 { return uint32(i) })},
		{"float64", true, roundTrip(func(i int) float64 { return float64(i) / 3 })},
		{"complex128", true, roundTrip(func(i int) complex128 { return complex(float64(i), 1) })},
		{"[2]uint64", true, roundTrip(func(i int) [2]uint64 { return [2]uint64{uint64(i), ^uint64(i)} })},
		{"struct of uint32 and uint16", true, roundTrip(func(i int) small { return small{uint32(i), uint16(i)} })},
		{"[3]uint64 is too large", false, roundTrip(func(i int) [3]uint64 { return [3]uint64{uint64(i)} })},
		{"string has a pointer", false, roundTrip(func(i int) string { return fmt.Sprint(i) })},
		{"struct{} is empty", false, roundTrip(func(int) struct{} { return struct{}{} })},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.run(t) > 0; got != tc.flat {
				t.Fatalf("flat leaves = %v, want %v", got, tc.flat)
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

// TestPointerFree checks the type test that decides whether values may live
// in memory the garbage collector does not scan: only types that hold no
// pointer anywhere qualify.
func TestPointerFree(t *testing.T) {
	for _, tc := range []struct {
		typ  reflect.Type
		want bool
	}{
		{reflect.TypeFor[uint64](), true},
		{reflect.TypeFor[[0]*int](), true},
		{reflect.TypeFor[[2]int16](), true},
		{reflect.TypeFor[struct{ a, b uint8 }](), true},
		{reflect.TypeFor[[2]*int](), false},
		{reflect.TypeFor[struct {
			A int
			P *int
		}](), false},
		{reflect.TypeFor[string](), false},
		{reflect.TypeFor[map[int]int](), false},
	} {
		if got := pointerFree(tc.typ); got != tc.want {
			t.Errorf("pointerFree(%v) = %v, want %v", tc.typ, got, tc.want)
		}
	}
}
