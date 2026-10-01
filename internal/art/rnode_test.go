package art

import (
	"bytes"
	"encoding/binary"
	"math/rand/v2"
	"slices"
	"testing"
)

// kindsOf counts the objects below n by kind, terms included.
func kindsOf(n *header, count map[kind]int) {
	count[n.kind]++
	if n.kind <= kLastPage {
		return
	}
	if termOf(n) != nil {
		count[kSet]++
	}
	if isRange(n.kind) {
		for _, c := range asR(n).children()[:asR(n).n] {
			kindsOf(c, count)
		}
		return
	}
	eachChild(n, func(_ byte, c *header) { kindsOf(c, count) })
}

// TestRangeNodeClasses makes sure that integer keys whose first bytes spread
// over few or many values are indexed correctly however many ranges that takes.
// It covers the range nodes of the ART behind multimap.Ordered, which hold 8,
// 24, 56 or 256 ranges and move into the next class as ranges arrive and leave:
// each class has to appear, with and without a path long enough to need a tail
// that the node carries along when a key leaves its path and the path shrinks.
// After every phase the map must match a reference and hold its invariants.
func TestRangeNodeClasses(t *testing.T) {
	r := rand.New(rand.NewPCG(11, 12))
	seen := map[kind]bool{}
	for _, fan := range []int{2, 5, 12, 30, 60, 120, 250} {
		for _, prefix := range []int{0, 30} {
			var m Map[uint64]
			ref := reference{}
			pre := bytes.Repeat([]byte{'p'}, prefix)
			var keys [][]byte
			for i := 0; i < fan*40; i++ { // more keys per first byte than a page holds
				k := append(bytes.Clone(pre), byte(r.IntN(fan)*(256/fan)))
				keys = append(keys, binary.BigEndian.AppendUint64(k, r.Uint64()))
			}
			for i, k := range keys {
				m.Add(k, uint64(i))
				ref.add(k, uint64(i))
			}
			count := map[kind]int{}
			kindsOf(m.t.root, count)
			for k := range count {
				seen[k] = true
			}
			checkInvariants(t, &m.t)
			compare(t, &m, ref, id, r)
			if prefix > 0 {
				k := bytes.Clone(keys[0]) // leaves the root's path after 10 bytes, which shrinks it
				k[10] = 'q'
				m.Add(k, 1)
				ref.add(k, 1)
				checkInvariants(t, &m.t)
				compare(t, &m, ref, id, r)
			}
			for i, k := range r.Perm(len(keys)) {
				m.RemoveKey(keys[k])
				delete(ref, string(keys[k]))
				if i%97 == 0 {
					checkInvariants(t, &m.t)
					count := map[kind]int{}
					if m.t.root != nil {
						kindsOf(m.t.root, count)
					}
					for k := range count {
						seen[k] = true
					}
				}
			}
			compare(t, &m, ref, id, r)
		}
	}
	for _, k := range []kind{kR8, kR24, kR56, kR256} {
		if !seen[k] {
			t.Errorf("no range node of kind %d appeared", k)
		}
	}
}

// TestLongestPageKey makes sure that the longest key a page holds and keys
// around that length work like every other key. It covers the limit of the
// pages of the ART behind multimap.Ordered: a key of 255 bytes still goes into a
// page, whose key length byte is full, and keeps all its values when it gets
// a leaf, which has to hold the whole key as a string since no inline
// remainder is that long; a key of 256 bytes goes to a leaf at once.
func TestLongestPageKey(t *testing.T) {
	for _, n := range []int{maxPageKey - 1, maxPageKey, maxPageKey + 1} {
		var m Map[uint64]
		k := bytes.Repeat([]byte{'k'}, n)
		m.Add(k, 1)
		m.Add(k, 2)
		m.Add(k, 3)
		if got := valuesOf(&m, k); len(got) != 3 {
			t.Fatalf("key of %d bytes holds %v, want 3 values", n, got)
		}
		checkInvariants(t, &m.t)
		m.RemoveKey(k)
		if m.Len() != 0 {
			t.Fatalf("%d keys left after removing the key of %d bytes", m.Len(), n)
		}
	}
}

// TestRemoveAbsentNearMisses makes sure that removing keys that are not there
// changes nothing, however much they look like keys that are. It covers the
// delete path of the ART behind multimap.Ordered where a key is routed into a
// page, or past a range, but is not in it: it is shorter, longer, or differs in
// one byte from keys the page holds, and the page and the tree stay as they
// were.
func TestRemoveAbsentNearMisses(t *testing.T) {
	r := rand.New(rand.NewPCG(13, 14))
	for _, keys := range keySets() {
		var m Map[uint64]
		ref := reference{}
		for i, k := range keys {
			m.Add(k, uint64(i%5))
			ref.add(k, uint64(i%5))
		}
		for _, k := range keys {
			for _, probe := range nearMisses(k) {
				if _, there := ref[string(probe)]; there {
					continue
				}
				m.RemoveKey(probe)
				m.Remove(probe, 1)
			}
		}
		checkInvariants(t, &m.t)
		compare(t, &m, ref, id, r)
	}
}

// TestThinPagesOfOtherLengthsStayApart makes sure that deleting keys never
// loses or mixes up keys when two neighbouring pages thin out but cannot become
// one page. It covers the page merge of the ART behind multimap.Ordered: pages
// of integer keys of three bytes and of four bytes would fit one page by their
// number of keys, but a page of one length holds no other, and the keys of
// both lengths in one page would be too many for the page type that holds
// them; the pages stay as they are and every key keeps its value.
func TestThinPagesOfOtherLengthsStayApart(t *testing.T) {
	var m Map[uint64]
	ref := reference{}
	r := rand.New(rand.NewPCG(15, 16))
	var short, long [][]byte
	for i := range 40 {
		short = append(short, []byte{byte(i), 1, 1})
		long = append(long, []byte{byte(100 + i), 1, 1, 1})
	}
	for i, k := range slices.Concat(short, long) {
		m.Add(k, uint64(i))
		ref.add(k, uint64(i))
	}
	checkInvariants(t, &m.t)
	for i := range 40 {
		if i%8 != 0 {
			for _, k := range [][]byte{short[i], long[i]} {
				m.RemoveKey(k)
				delete(ref, string(k))
				checkInvariants(t, &m.t)
			}
		}
	}
	compare(t, &m, ref, id, r)
}
