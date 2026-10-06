package art

import (
	"bytes"
	"fmt"
	"runtime"
	"slices"
	"testing"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/page"
)

// TestFixedKeys makes sure that a key of multimap.Ordered with small values
// without pointers (uint64) keeps exactly its values while their number grows
// from one to hundreds and shrinks back, whatever the key's length. It covers
// the single-key page of fixed-size values (the one-key form of page.Fixed) behind Ordered, which
// holds the values right after the key's remainder and moves through its size
// classes: after every step the key must hold exactly the values added and not
// removed, a page must hold no more than the largest class takes, a key with more
// values than that (or a remainder too long for a page) must be a value overflow,
// and it must go back into a page once its values take half of the room that is
// left. The values must also survive a garbage collection.
func TestFixedKeys(t *testing.T) {
	for _, klen := range []int{0, 8, 26, 36, 66, 200, 499, 500, 600} {
		t.Run(fmt.Sprintf("key of %d bytes", klen), func(t *testing.T) {
			key := bytes.Repeat([]byte{'k'}, klen)
			var m Map[uint64]
			m.Add([]byte("neighbour"), 1) // the key's leaf lives below a node, which holds its first byte
			rem := max(klen-1, 0)         // what the leaf holds
			pageable := page.Header+rem+8 <= 512
			capLargest := 0
			if pageable {
				capLargest = (512 - page.Header - rem) / 8
			}
			var want []uint64
			check := func(removing bool) {
				t.Helper()
				runtime.GC()
				got := valuesOf(&m, key)
				slices.Sort(got)
				if !slices.Equal(got, want) {
					t.Fatalf("key holds %v, want %v", got, want)
				}
				if len(want) < 2 { // an entry with one value may be in a multi-key page
					return
				}
				l := m.t.findLeaf(key)
				if l == nil { // the key shares a multi-key page with its neighbour
					return
				}
				switch {
				case l.isValueOverflow() && pageable && !removing && len(want) <= capLargest:
					t.Fatalf("a value overflow with %d values, which a page holds", len(want))
				case l.isValueOverflow() && pageable && removing && page.BackFits(rem, 8*len(want)):
					t.Fatalf("a value overflow with %d values, which fit back into a page", len(want))
				case !l.isValueOverflow() && (!pageable || asFixed(l).Len() != len(want) || len(want) > capLargest):
					t.Fatalf("page holding %d values of at most %d", asFixed(l).Len(), capLargest)
				}
			}
			for v := range uint64(2*capLargest + 10) {
				m.Add(key, v)
				m.Add(key, v) // a duplicate changes nothing
				want = append(want, v)
				check(false)
			}
			for len(want) > 0 {
				v := want[len(want)/2]
				m.Remove(key, v)
				m.Remove(key, v) // an absent value changes nothing
				want = slices.DeleteFunc(want, func(x uint64) bool { return x == v })
				check(true)
			}
			if m.Has(key) || m.Len() != 1 {
				t.Fatalf("key still present after removing all its values")
			}
			checkInvariants(t, &m.t)
		})
	}
}

// TestFixedKeyMovesUp makes sure that a key of multimap.Ordered keeps all its
// values when the keys it shared a long common part with go away. A leaf holds
// only the part of its key below its node; when that node goes, the leaf takes its
// place and must hold more of its key (rekey). A page then moves into a class
// that holds the longer key and its values, or, if no page does, into a value
// overflow.
func TestFixedKeyMovesUp(t *testing.T) {
	common := bytes.Repeat([]byte("c"), 200)
	for _, n := range []int{1, 5, 62} {
		t.Run(fmt.Sprintf("%d values", n), func(t *testing.T) {
			var m Map[uint64]
			k1, k2, k3 := append(slices.Clip(common), '1'), append(slices.Clip(common), '2'), append(slices.Clip(common), '3')
			for _, k := range [][]byte{k2, k3} { // k2 and k3 make the node holding the common part; with two values they are single-key pages
				m.Add(k, 7)
				m.Add(k, 8)
			}
			var want []uint64
			for v := range uint64(n) {
				m.Add(k1, v)
				want = append(want, v)
			}
			if l := m.t.findLeaf(k1); l != nil && len(l.stored()) != 0 { // nil: the keys share a multi-key page
				t.Fatalf("leaf below the common part holds %q, want nothing", l.stored())
			}
			m.RemoveKey(k2)
			m.RemoveKey(k3)
			checkInvariants(t, &m.t)
			l := m.t.findLeaf(k1)
			got := valuesOf(&m, k1)
			slices.Sort(got)
			if !slices.Equal(got, want) || !bytes.Equal(l.stored(), k1) {
				t.Fatalf("after moving up the leaf holds %d bytes and values %v, want the whole key and %v", len(l.stored()), got, want)
			}
			// a key part of 201 bytes leaves 307 bytes behind the head: 38 values of 8 bytes fit
			if overflow := n > 38; overflow != l.isValueOverflow() {
				t.Fatalf("%d values: value overflow = %v", n, l.isValueOverflow())
			}
		})
	}
}

// TestFixedRekey makes sure a leaf that has to hold more of its key, because the
// node above it went away, keeps every value, and that it stays where it is as long
// as the longer key part still fits its size class (and, for a value overflow, its key
// area): a delete that merges a node into its only leaf then costs no new leaf,
// which with one value per key is what churn and build pay for most; a longer key
// part than the class holds moves the leaf, and one that no page holds makes it a value
// overflow.
func TestFixedRekey(t *testing.T) {
	rest := bytes.Repeat([]byte("abcdefghij"), 70) // 700 bytes
	for _, tc := range []struct {
		name         string
		overflow     bool // the key starts as a value overflow
		restLen      int  // the key part of the leaf
		front        int  // the bytes that come in front of it
		values       int
		wantInPlace  bool
		wantOverflow bool
	}{
		{"page with room", false, 3, 6, 1, true, false},
		{"page with several values and room", false, 3, 6, 4, true, false},
		{"page without room in its class", false, 20, 3, 1, false, false},
		{"page whose key no page holds", false, 300, 250, 2, false, true},
		{"value overflow with room in its key area", true, 3, 6, 1, true, true},
		{"value overflow beyond its key area", true, 20, 6, 1, false, true},
		{"value overflow up to the longest inline remainder", true, 400, 100, 1, true, true},
		{"value overflow beyond it: the key as a string", true, 400, 101, 1, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var m Map[uint64]
			m.decide()
			k := rest[:tc.restLen]
			var l *singleKeyHead
			var want []uint64
			if tc.overflow {
				l = newValueOverflow(k, set3.EmptyWithCapacity[uint64](4))
			}
			var rests [][]byte
			for v := range uint64(tc.values) {
				if tc.overflow {
					overflowAdd(l, v)
				}
				rests = append(rests, k)
				want = append(want, v)
			}
			if !tc.overflow {
				l = fixedHead(page.BuildFixed(rests, want))
			}
			front := rest[tc.restLen : tc.restLen+tc.front]
			nl := m.rekey(l, front, -1, 0)
			var got []uint64
			eachValue(nl, 1, func(v uint64) bool { got = append(got, v); return true })
			slices.Sort(got)
			if !slices.Equal(got, want) {
				t.Errorf("values %v, want %v", got, want)
			}
			if !bytes.Equal(nl.stored(), append(slices.Clone(front), k...)) {
				t.Errorf("leaf holds %d bytes, want the %d bytes of front and key part", len(nl.stored()), tc.front+tc.restLen)
			}
			if (nl == l) != tc.wantInPlace || nl.isValueOverflow() != tc.wantOverflow {
				t.Errorf("in place: %v (want %v), value overflow: %v (want %v)", nl == l, tc.wantInPlace, nl.isValueOverflow(), tc.wantOverflow)
			}
		})
	}
}
