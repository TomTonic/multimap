package art

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// TestPageLifecycle makes sure that integer keys with one value each stay
// compact however their number changes. It covers the pages of the ART behind
// multimap.Ordered: a page grows through every class as keys arrive, bursts
// into a node with smaller pages when it overflows, shrinks back through the
// classes as keys leave, moves up when its sibling is gone, and every key
// keeps its value throughout.
func TestPageLifecycle(t *testing.T) {
	var keys [][]byte // 6 shared bytes, then two groups of 16
	for a := range 2 {
		for b := range 16 {
			keys = append(keys, []byte{1, 2, 3, 4, 5, 6, byte(a), byte(3 * b)})
		}
	}
	var m Map[uint64]
	checkValues := func(from int) {
		t.Helper()
		for i, k := range keys[from:] {
			// Each stops, and reports false, only at the expected value
			if v := m.Values(k); v.Len() != 1 || v.Each(func(x uint64) bool { return x != uint64(from+i) }) {
				t.Fatalf("key %x lost its value %d", k, from+i)
			}
		}
	}
	for i, k := range keys[:31] {
		m.Add(k, uint64(i))
		checkInvariants(t, &m.t)
		if r := m.t.root; r.kind != kPage || int(asPage(r).class) != classFor(i+1) {
			t.Fatalf("after %d keys: root kind %d, want a page of class %d", i+1, r.kind, classFor(i+1))
		}
	}
	m.Add(keys[31], 31) // the 32nd key bursts the full page
	checkInvariants(t, &m.t)
	r := m.t.root
	if r.kind != kN4 || r.plen != 6 || r.count != 2 {
		t.Fatalf("after the burst: root kind %d, path %d, %d children; want a 4-way node, path 6, 2 children", r.kind, r.plen, r.count)
	}
	eachChild(r, func(_ byte, c *header) {
		if c.kind != kPage || asPage(c).count != 16 {
			t.Fatalf("after the burst: child kind %d, want a page of 16 keys", c.kind)
		}
	})
	checkValues(0)

	var classes []int
	for i, k := range keys[:16] {
		m.RemoveKey(k)
		checkInvariants(t, &m.t)
		checkValues(i + 1)
		if i < 15 {
			_, child := sorted(m.t.root)
			classes = append(classes, int(asPage(child[0]).class))
		}
	}
	if want := []int{3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 1, 1, 0, 0}; !slices.Equal(classes, want) {
		t.Fatalf("classes while shrinking = %v, want %v", classes, want)
	}
	if r := m.t.root; r.kind != kPage || asPage(r).count != 16 {
		t.Fatalf("after emptying one page: root kind %d, want the other page", r.kind)
	}
}

// TestSmallPlain makes sure that only values that fit a page's 8-byte word
// and hold no pointers go into pages, where the garbage collector never looks.
// It covers the value-type check of the ART behind multimap.Ordered.
func TestSmallPlain(t *testing.T) {
	type pair struct{ A, B uint32 }
	type withPtr struct{ P *int }
	for _, tc := range []struct {
		name string
		got  bool
		want bool
	}{
		{"uint64 goes into pages", smallPlain[uint64](), true},
		{"int8 goes into pages", smallPlain[int8](), true},
		{"float32 goes into pages", smallPlain[float32](), true},
		{"bool goes into pages", smallPlain[bool](), true},
		{"a struct of two uint32 goes into pages", smallPlain[pair](), true},
		{"an array of four uint16 goes into pages", smallPlain[[4]uint16](), true},
		{"an empty array goes into pages", smallPlain[[0]*int](), true},
		{"an empty struct goes into pages", smallPlain[struct{}](), true},
		{"a string does not", smallPlain[string](), false},
		{"a pointer does not", smallPlain[*int](), false},
		{"a struct with a pointer does not", smallPlain[withPtr](), false},
		{"an array of pointers does not", smallPlain[[1]*int](), false},
		{"a 16-byte value does not", smallPlain[complex128](), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("smallPlain = %v, want %v", tc.got, tc.want)
			}
		})
	}
}

// TestValueTypesInPages makes sure that values of any small plain type come
// back from a page exactly as they went in, including negative numbers and
// types narrower than the page's 8-byte word, and that a key that gets a
// second value leaves its page with both. It covers the raw value storage of
// the pages behind multimap.Ordered.
func TestValueTypesInPages(t *testing.T) {
	type pair struct{ A, B int16 }
	t.Run("int8", func(t *testing.T) { checkValueType(t, []int8{-128, -1, 0, 127}) })
	t.Run("float32", func(t *testing.T) { checkValueType(t, []float32{-1.5, 0, 3.25, 1e30}) })
	t.Run("bool", func(t *testing.T) { checkValueType(t, []bool{true, false}) })
	t.Run("struct", func(t *testing.T) { checkValueType(t, []pair{{-1, 2}, {3, -4}, {0, 0}}) })
	t.Run("string values stay in leaves", func(t *testing.T) { checkValueType(t, []string{"a", "bb", ""}) })
}

// checkValueType stores one value per key, reads the values back through
// Values and RangeValues, gives one key a second value, and removes it all.
func checkValueType[T comparable](t *testing.T, vs []T) {
	t.Helper()
	var m Map[T]
	key := func(i int) []byte { return []byte{0, 0, 0, 0, 0, 0, 0, byte(i)} }
	for i, v := range vs {
		m.Add(key(i), v)
	}
	if got, want := m.t.root.kind == kPage, smallPlain[T](); got != want {
		t.Fatalf("root is a page: %v, want %v", got, want)
	}
	var got []T
	m.RangeValues(&Bounds{}, func(v T) bool { got = append(got, v); return true })
	if !slices.Equal(got, vs) {
		t.Fatalf("RangeValues = %v, want %v", got, vs)
	}
	for i, v := range vs {
		if s := m.Values(key(i)); s.Len() != 1 || s.Each(func(x T) bool { return x != v }) {
			t.Fatalf("Values(key %d) does not hold exactly %v", i, v)
		}
	}
	m.Add(key(0), vs[1]) // a second value
	checkInvariants(t, &m.t)
	if s := m.Values(key(0)); s.Len() != 2 {
		t.Fatalf("key 0 holds %d values after a second one, want 2", s.Len())
	}
	for i, v := range vs {
		m.Remove(key(i), v)
	}
	m.Remove(key(0), vs[1])
	if m.Len() != 0 {
		t.Fatalf("%d keys left after removing every value", m.Len())
	}
}

// TestPageNLifecycle makes sure that integer keys keep all their values, in
// any mix of one, a few and many values per key, while their pages change
// shape underneath. It covers the U8-n pages of the ART behind
// multimap.Ordered: a U8-1 page turning into a U8-n page in place and by
// rebuilding, values moving out to external sets as a key passes inlineMax,
// pages growing for lack of external slots and bursting beyond the largest
// class, an externally stored key becoming a node's term key, and the removal
// of values and keys. After every step the map must match a reference and
// satisfy the structural invariants.
func TestPageNLifecycle(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 4))
	var m Map[uint64]
	ref := reference{}
	add := func(k []byte, vs ...uint64) {
		t.Helper()
		for _, v := range vs {
			m.Add(k, v)
			ref.add(k, v)
		}
		compare(t, &m, ref, r)
		checkInvariants(t, &m.t)
	}
	kind := func(k []byte) kind {
		n, _ := m.t.find(k)
		return n.kind
	}
	many := func(from, n int) []uint64 {
		out := make([]uint64, n)
		for i := range out {
			out[i] = uint64(from + i)
		}
		return out
	}
	key := func(a, b byte) []byte { return []byte{9, 9, 9, 9, 9, 9, a, b} }

	// 16 keys with one value each, then a second value: U8-1 turns into U8-n in place
	for b := range byte(16) {
		add(key(0, b), 1)
	}
	add(key(0, 3), 2)
	if kind(key(0, 3)) != kPageN {
		t.Fatalf("a second value did not turn the page into a U8-n page")
	}
	// two keys with more than inlineMax values: the second needs a larger class for its slot
	add(key(0, 5), many(100, inlineMax+1)...)
	add(key(0, 6), many(200, inlineMax+1)...)
	// three more: the largest class has 4 external slots, so the fifth rebuilds the subtree
	for _, b := range []byte{7, 8, 9} {
		add(key(0, b), many(300, inlineMax+2)...)
	}
	// a shorter key makes an externally stored key a term
	add([]byte{9, 9, 9, 9, 9, 9, 0}, 7)
	// 20 keys with one value each in one page, then a second value: too many keys for U8-n
	for b := range byte(20) {
		add(key(1, b), 1)
	}
	add(key(1, 4), 2)
	// remove values and whole keys, external ones included
	for _, k := range [][]byte{key(0, 5), key(0, 7), key(0, 3), key(1, 4)} {
		for v := range ref[string(k)] {
			m.Remove(k, v)
			ref.remove(k, v)
			compare(t, &m, ref, r)
			checkInvariants(t, &m.t)
		}
	}
	for _, k := range ref.sortedKeys() {
		m.RemoveKey([]byte(k))
		delete(ref, k)
		checkInvariants(t, &m.t)
	}
	if m.Len() != 0 {
		t.Fatalf("%d keys left", m.Len())
	}

	// In the smallest class (4 keys, 1 external slot), a second key with many
	// values needs a larger class for its slot.
	add([]byte("ab"), many(0, inlineMax+1)...)
	add([]byte("ac"), many(50, inlineMax+1)...)
	if n, _ := m.t.find([]byte("ab")); n.kind != kPageN || asPage(n).extUsed() != 2 {
		t.Fatalf("two externally stored keys do not share one U8-n page")
	}
	// Removing an externally stored key from a page that keeps other keys.
	add([]byte("ad"), 1, 2)
	m.RemoveKey([]byte("ac"))
	delete(ref, "ac")
	compare(t, &m, ref, r)
	checkInvariants(t, &m.t)
	// A longer key makes the externally stored "ab" the term of a new node.
	add([]byte("abx"), 1)
	if n, _ := m.t.find([]byte("ab")); n.kind != kLeaf {
		t.Fatalf("the term key \"ab\" is not a leaf")
	}
}
