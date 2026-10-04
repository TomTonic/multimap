package art

import (
	"bytes"
	"encoding/binary"
	"math/rand/v2"
	"testing"
)

// typesOf counts the objects below n by type, end pages included.
func typesOf(n *header, count map[objType]int) {
	count[n.objType]++
	if n.objType <= maxMultiKeyByte {
		return
	}
	if endPageOf(n) != nil {
		count[kValueOverflow]++
	}
	if isRange(n.objType) {
		for _, c := range asR(n).children()[:asR(n).n] {
			typesOf(c, count)
		}
		return
	}
	eachChild(n, func(_ byte, c *header) { typesOf(c, count) })
}

// TestRangeNodeClasses makes sure that integer keys whose first bytes spread
// over few or many values are indexed correctly however many ranges that takes.
// It covers the range nodes of the ART behind multimap.Ordered, which hold 8,
// 24, 56 or 256 ranges and move into the next class as ranges arrive and leave:
// each class has to appear. After every phase the map must match a reference
// and hold its invariants.
func TestRangeNodeClasses(t *testing.T) {
	r := rand.New(rand.NewPCG(11, 12))
	seen := map[objType]bool{}
	note := func(m *Map[uint64]) {
		count := map[objType]int{}
		if m.t.root != nil {
			typesOf(m.t.root, count)
		}
		for k := range count {
			seen[k] = true
		}
	}
	for _, fan := range []int{2, 5, 12, 30, 60, 120, 250} {
		var m Map[uint64]
		ref := reference{}
		var keys [][]byte
		for i := 0; i < fan*40; i++ { // more keys per first byte than a page holds
			k := []byte{byte(r.IntN(fan) * (256 / fan))}
			keys = append(keys, binary.BigEndian.AppendUint64(k, r.Uint64())[:8])
		}
		for i, k := range keys {
			m.Add(k, uint64(i))
			ref.add(k, uint64(i))
		}
		note(&m)
		checkInvariants(t, &m.t)
		compare(t, &m, ref, id, r)
		for i, k := range r.Perm(len(keys)) {
			m.RemoveKey(keys[k])
			delete(ref, string(keys[k]))
			if i%97 == 0 {
				checkInvariants(t, &m.t)
				note(&m)
			}
		}
		compare(t, &m, ref, id, r)
	}
	for _, k := range []objType{kR8, kR24, kR56, kR256} {
		if !seen[k] {
			t.Errorf("no range node of type %d appeared", k)
		}
	}
}

// TestPageKeyLength makes sure that keys of any length go into pages, and that
// a key that no page holds, because more than 255 bytes of it lie below the
// page, gets a leaf and stays among them. It covers the pages of the ART behind
// multimap.Ordered, which hold the part of every key below their base: a map
// whose first key is such a long one has pages for the keys that follow, and a
// map that has become empty starts again.
func TestPageKeyLength(t *testing.T) {
	var m Map[uint64]
	ref := reference{}
	r := rand.New(rand.NewPCG(13, 14))
	fill := func(n int) {
		for i := range n {
			ks := [][]byte{binary.BigEndian.AppendUint32(nil, uint32(i)*7919)}
			if i%50 == 0 && m.Len() > 100 { // few keys of other lengths, so that they do not crowd the pages
				ks = append(ks, binary.BigEndian.AppendUint64(nil, uint64(i)*7919), binary.BigEndian.AppendUint64([]byte("long prefix"), uint64(i)))
			}
			for _, k := range ks {
				m.Add(k, uint64(i))
				ref.add(k, uint64(i))
			}
		}
		checkInvariants(t, &m.t)
		compare(t, &m, ref, id, r)
	}
	pages := func() int {
		count := map[objType]int{}
		typesOf(m.t.root, count)
		n := 0
		for k := kMultiKey; k <= kLastMultiKey; k += 2 {
			n += count[k]
		}
		return n
	}
	leaves := func(m *Map[uint64]) int {
		n := 0
		m.Objects(func(o Object) { n += b2i(o.Label == "flat leaf" || o.Label == "value overflow") })
		return n
	}
	fill(2000)
	if pages() == 0 {
		t.Fatal("keys of three lengths did not go into pages")
	}
	if n := leaves(&m); n > 10 { // keys that are the prefix of others end at a node: their leaf is its end page
		t.Fatalf("%d of %d keys of one value have leaves, want pages for nearly all", n, m.Len())
	}
	for k := range ref {
		m.RemoveKey([]byte(k))
		delete(ref, k)
	}
	if m.t.root != nil {
		t.Fatalf("%d keys left", m.Len())
	}
	long := bytes.Repeat([]byte{'x'}, 300) // too long for a page
	m.Add(long, 1)
	ref.add(long, 1)
	fill(2000)
	if pages() == 0 || leaves(&m) < 1 || leaves(&m) > 11 {
		t.Fatalf("after a first key that is too long: %d pages, %d leaves, want pages and a few leaves", pages(), leaves(&m))
	}
}

// TestForkAroundLeaves makes sure that a key goes in next to a leaf that is a
// prefix of it, or that it is a prefix of, in a map whose keys have the length
// of the first one. It covers the range node that a leaf and a key for a page
// get as parents in the ART behind multimap.Ordered: the leaf is the node's
// end page, or the key is.
func TestForkAroundLeaves(t *testing.T) {
	for _, tc := range []struct{ name, first, leaf, then string }{
		{"the leaf is a prefix of the key", "wxyz", "wx", "wxab"},
		{"the key is a prefix of the leaf", "pq", "pqrst", "pq"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var m Map[uint64]
			r := rand.New(rand.NewPCG(17, 18))
			ref := reference{}
			add := func(k string) {
				m.Add([]byte(k), 1)
				ref.add([]byte(k), 1)
			}
			add(tc.first)
			add(tc.leaf)
			m.RemoveKey([]byte(tc.first)) // only the leaf is left, at the root
			delete(ref, tc.first)
			add(tc.then)
			checkInvariants(t, &m.t)
			compare(t, &m, ref, id, r)
		})
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
			m.Remove(k, 99) // a value the key does not hold: the key stays
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
