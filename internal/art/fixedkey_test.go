package art

import (
	"bytes"
	"fmt"
	"runtime"
	"slices"
	"testing"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/skpage"
)

// TestFixedKeys makes sure that a key of multimap.Ordered with small values
// without pointers (uint64) keeps exactly its values while their number grows
// from one to hundreds and shrinks back, whatever the key's length. It covers
// the single-key page of fixed-size values (skpage.Fixed) behind Ordered, which
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
			pageable := rem <= skpage.MaxRemainderFixed[uint64]()
			capLargest := 0
			if pageable {
				capLargest = (512 - (skpage.Header+rem+7)&^7) / 8
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
				if len(want) == 0 {
					return
				}
				switch l := m.t.find(key); {
				case l.isValueOverflow() && pageable && !removing && len(want) <= capLargest:
					t.Fatalf("a value overflow with %d values, which a page holds", len(want))
				case l.isValueOverflow() && pageable && removing && skpage.BackFits(rem, 8*len(want)):
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
			m.Add(k2, 7) // k2 and k3 make the node holding the common part
			m.Add(k3, 7)
			var want []uint64
			for v := range uint64(n) {
				m.Add(k1, v)
				want = append(want, v)
			}
			if l := m.t.find(k1); len(l.stored()) != 0 {
				t.Fatalf("leaf below the common part holds %q, want nothing", l.stored())
			}
			m.RemoveKey(k2)
			m.RemoveKey(k3)
			checkInvariants(t, &m.t)
			l := m.t.find(k1)
			got := valuesOf(&m, k1)
			slices.Sort(got)
			if !slices.Equal(got, want) || l.base() != 0 || !bytes.Equal(l.stored(), k1) {
				t.Fatalf("after moving up the leaf holds %d bytes from %d and values %v, want the whole key and %v", len(l.stored()), l.base(), got, want)
			}
			// a remainder of 201 bytes puts the values at 208: 38 of 8 bytes fit
			if overflow := n > 38; overflow != l.isValueOverflow() {
				t.Fatalf("%d values: value overflow = %v", n, l.isValueOverflow())
			}
		})
	}
}

// TestFixedRekey makes sure a leaf that has to hold more of its key, because the
// node above it went away, keeps every value, and that it stays where it is as long
// as the longer key still fits its size class (and, for a value overflow, its key
// area): a delete that merges a node into its only leaf then costs no new leaf,
// which with one value per key is what churn and build pay for most; a longer key
// than the class holds moves the leaf, and one that no page holds makes it a value
// overflow.
func TestFixedRekey(t *testing.T) {
	key := bytes.Repeat([]byte("abcdefghij"), 70) // 700 bytes
	for _, tc := range []struct {
		name         string
		overflow     bool // the key starts as a value overflow
		keyLen, base int
		to           int
		values       int
		wantInPlace  bool
		wantOverflow bool
	}{
		{"page with room", false, 20, 17, 11, 1, true, false},
		{"page with several values and room", false, 20, 17, 11, 3, true, false},
		{"page whose values move up a word", false, 30, 27, 13, 1, true, false},
		{"page without room in its class", false, 40, 37, 0, 1, false, false},
		{"page whose key no page holds", false, 600, 300, 0, 2, false, true},
		{"value overflow with room in its key area", true, 20, 17, 6, 1, true, true},
		{"value overflow beyond its key area", true, 20, 17, 0, 1, false, true},
		{"value overflow up to the longest inline remainder", true, 600, 120, 102, 1, true, true},
		{"value overflow beyond it: the key as a string", true, 600, 120, 101, 1, false, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			k := key[:tc.keyLen]
			l := newFixedLeaf[uint64](k, tc.base)
			slot := singleKeyHdr(l)
			if tc.overflow {
				l = newValueOverflow(k, tc.base, set3.EmptyWithCapacity[uint64](4))
				slot = singleKeyHdr(l)
			}
			var want []uint64
			for v := range uint64(tc.values) {
				if l.isValueOverflow() {
					overflowAdd(l, v)
				} else {
					addFixed(&slot, l, k, v)
					l = asSingleKey(slot)
				}
				want = append(want, v)
			}
			nl := rekeyFixed[uint64](l, k[:tc.base-1], int(k[tc.base-1]), tc.to)
			var got []uint64
			eachValue(nl, 1, func(v uint64) bool { got = append(got, v); return true })
			slices.Sort(got)
			if !slices.Equal(got, want) {
				t.Errorf("values %v, want %v", got, want)
			}
			if nl.keyLen() != tc.keyLen || nl.base() > tc.to || !bytes.Equal(nl.from(tc.to), k[tc.to:]) {
				t.Errorf("leaf holds %q from %d of a key of %d bytes, want the key from %d on", nl.stored(), nl.base(), nl.keyLen(), tc.to)
			}
			if (nl == l) != tc.wantInPlace || nl.isValueOverflow() != tc.wantOverflow {
				t.Errorf("in place: %v (want %v), value overflow: %v (want %v)", nl == l, tc.wantInPlace, nl.isValueOverflow(), tc.wantOverflow)
			}
		})
	}
}
