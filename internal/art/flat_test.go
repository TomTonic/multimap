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
// Ordered, which hold small pointer-free values right after the key's
// remainder and move through size classes as values arrive: after every step
// the key must hold exactly the values added and not removed, a flat leaf
// must sit in a class that fits its values, a key with more values than the
// largest class holds must have spilled into a set leaf, and it must turn
// flat again once half of the largest class would do. The values must also
// survive a garbage collection, since the leaves are memory the collector
// never scans.
func TestFlatLeaves(t *testing.T) {
	for _, klen := range []int{0, 8, 26, 36, 66, 200, maxInline + 1, maxInline + 2, 300} {
		t.Run(fmt.Sprintf("key of %d bytes", klen), func(t *testing.T) {
			key := bytes.Repeat([]byte{'k'}, klen)
			var m Map[uint64]
			leavesOnly(&m)
			m.Add([]byte("neighbour"), 1) // the key's leaf lives below a node, which holds its first byte
			rem := max(klen-1, 0)         // what the leaf holds
			maxCap := flatCap[uint64](uint8(len(flatSizes)-1), rem)
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
				l := findLeaf(&m.t, key)
				switch {
				case rem > maxInline:
					if l.cls() != 0 {
						t.Fatalf("a key of %d bytes got a flat leaf", klen)
					}
				case l.cls() != 0:
					if int(l.n) != len(want) || len(want) > flatCap[uint64](l.cls(), rem) {
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
				if l := findLeaf(&m.t, key); rem <= maxInline && len(want) > maxCap && l.cls() != 0 {
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
	leavesOnly(&m)
	key := []byte("hovering")
	full := flatCap[uint64](minGrown, len(key)) // a cache-line leaf holds this many
	for v := range uint64(full + 1) {
		m.Add(key, v)
	}
	l := findLeaf(&m.t, key)
	if want := flatClass[uint64](len(key), 2*full); l.cls() != want {
		t.Fatalf("leaf in class %d after %d values, want class %d, which holds twice as many", l.cls(), full+1, want)
	}
	for range 10 {
		m.Remove(key, uint64(full))
		m.Add(key, uint64(full))
		if findLeaf(&m.t, key) != l {
			t.Fatalf("hovering at %d values moved the leaf", full)
		}
	}

	// A key that once grew keeps a cache line while it hovers between one
	// and a few values; its first leaf of 32 bytes holds 3.
	few := []byte("few")
	for v := range uint64(4) { // one more than the smallest leaf holds
		m.Add(few, v)
	}
	l = findLeaf(&m.t, few)
	for range 10 {
		for v := range uint64(3) {
			m.Remove(few, v+1)
		}
		for v := range uint64(3) {
			m.Add(few, v+1)
		}
		if findLeaf(&m.t, few) != l || l.cls() != minGrown {
			t.Fatalf("hovering between 1 and 4 values moved the leaf or left the cache line")
		}
	}
}

// TestFlatValueTypes makes sure that multimap.Ordered keeps values of every
// type correctly, whether its ART stores them in flat leaves (small values
// without pointers, of any alignment), in typed leaves (small values with a
// pointer) or in set leaves (everything else).
func TestFlatValueTypes(t *testing.T) {
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
		mode int8 // 1: flat leaves, 2: typed leaves, -1: set leaves only
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
		{"pointer", 2, roundTrip(func(i int) *int { return &ptrs[i%300] })},
		{"interface", 2, roundTrip(func(i int) any { return i })},
		{"struct of a string and a number", 2, roundTrip(func(i int) named { return named{fmt.Sprint(i), i} })},
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
		leavesOnly(&m)
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

// TestLeafMovesUp makes sure that a key of multimap.Ordered keeps all its
// values when the keys it shared a long common part with go away. In the ART
// behind Ordered a leaf holds only the part of its key below its node; when
// that node goes, the leaf takes its place and must hold more of its key
// (rekey). A flat leaf then moves into a class that holds the longer key and
// its values, or, if no flat leaf does, into a set leaf.
func TestLeafMovesUp(t *testing.T) {
	common := bytes.Repeat([]byte("c"), 200)
	for _, n := range []int{1, 5, 62} {
		t.Run(fmt.Sprintf("%d values", n), func(t *testing.T) {
			var m Map[uint64]
			leavesOnly(&m)
			k1, k2, k3 := append(slices.Clip(common), '1'), append(slices.Clip(common), '2'), append(slices.Clip(common), '3')
			m.Add(k2, 7) // k2 and k3 make the node holding the common part
			m.Add(k3, 7)
			var want []uint64
			for v := range uint64(n) {
				m.Add(k1, v)
				want = append(want, v)
			}
			if l := findLeaf(&m.t, k1); len(l.stored()) != 0 {
				t.Fatalf("leaf below the common part holds %q, want nothing", l.stored())
			}
			m.RemoveKey(k2)
			m.RemoveKey(k3)
			checkInvariants(t, &m.t)
			l := findLeaf(&m.t, k1)
			got := valuesOf(&m, k1)
			slices.Sort(got)
			if !slices.Equal(got, want) || l.base() != 0 || !bytes.Equal(l.stored(), k1) {
				t.Fatalf("after moving up the leaf holds %d bytes from %d and values %v, want the whole key and %v", len(l.stored()), l.base(), got, want)
			}
			if flat := flatClassFor[uint64](k1, 0, n) != 0; flat != (!l.isSet()) {
				t.Fatalf("leaf of kind %d, want a flat leaf: %v", l.kind, flat)
			}
		})
	}
}

// TestRekeyInPlace makes sure a leaf that has to hold more of its key, because
// the node above it went away, keeps every value, and that it stays where it is
// as long as the longer key still fits its size class. It covers rekey for
// flat leaves (integer values) and set leaves (values that hold a pointer): a
// delete that merges a node into its only leaf then costs no new leaf, which
// with one value per key is what churn and build pay for most; a longer key
// than the class holds moves the leaf.
func TestRekeyInPlace(t *testing.T) {
	key := bytes.Repeat([]byte("abcdefghij"), 30) // 300 bytes
	for _, tc := range []struct {
		name         string
		flat         bool
		keyLen, base int
		to           int
		values       int
		wantInPlace  bool
	}{
		{"flat leaf with room", true, 20, 17, 11, 1, true},
		{"flat leaf with several values and room", true, 20, 17, 11, 3, true},
		{"flat leaf without room", true, 40, 37, 0, 1, false},
		{"set leaf with room in its class", false, 20, 17, 6, 1, true},
		{"set leaf beyond its class", false, 20, 17, 2, 1, false},
		{"set leaf in the largest class", false, 230, 30, 1, 1, true},
		{"set leaf up to the longest inline remainder", false, 300, 100, 46, 1, true},
		{"set leaf beyond the longest inline remainder", false, 300, 100, 45, 1, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			k := key[:tc.keyLen]
			var l, nl *leafHead
			var got []string
			want := make([]string, tc.values)
			for i := range want {
				want[i] = fmt.Sprint(i + 1)
			}
			if tc.flat {
				l = newFlatLeaf[uint64](k, tc.base)
				for i := range tc.values {
					if g := flatAdd(l, uint64(i+1)); g != nil {
						l = g
					}
				}
				nl = rekey[uint64](l, k[:tc.base-1], int(k[tc.base-1]), tc.to)
				for _, v := range flatVals[uint64](nl) {
					got = append(got, fmt.Sprint(v))
				}
			} else {
				l = newSetLeaf[string](k, tc.base)
				for _, v := range want {
					vals[string](l).Add(v)
				}
				nl = rekey[string](l, k[:tc.base-1], int(k[tc.base-1]), tc.to)
				vals[string](nl).Each(func(v string) bool { got = append(got, v); return true })
			}
			slices.Sort(got)
			if !slices.Equal(got, want) {
				t.Errorf("values %v, want %v", got, want)
			}
			if nl.keyLen() != tc.keyLen || nl.base() > tc.to || !bytes.Equal(nl.from(tc.to), k[tc.to:]) {
				t.Errorf("leaf holds %q from %d of a key of %d bytes, want the key from %d on", nl.stored(), nl.base(), nl.keyLen(), tc.to)
			}
			if inPlace := nl == l; inPlace != tc.wantInPlace {
				t.Errorf("leaf stayed in place: %v, want %v", inPlace, tc.wantInPlace)
			}
		})
	}
}
