package art

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
)

// TestPageLifecycle makes sure that integer keys with one value each stay
// compact however their number changes. It covers the pages of the ART behind
// multimap.Ordered: a page grows through every class as keys arrive, splits
// into two half-full pages below a range node when it overflows, shrinks back
// through the classes as keys leave, moves up when its sibling is gone, and
// every key keeps its value throughout.
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
	m.Add(keys[31], 31) // the 32nd key splits the full page
	checkInvariants(t, &m.t)
	r := m.t.root
	if r.kind != kR || r.plen != 6 || r.count != 2 {
		t.Fatalf("after the split: root kind %d, path %d, %d ranges; want a range node, path 6, 2 ranges", r.kind, r.plen, r.count)
	}
	eachChild(r, func(_ byte, c *header) {
		if c.kind != kPage || asPage(c).count != 16 {
			t.Fatalf("after the split: child kind %d, want a page of 16 keys", c.kind)
		}
	})
	checkValues(0)

	var classes []int
	for i, k := range keys[:16] {
		m.RemoveKey(k)
		checkInvariants(t, &m.t)
		checkValues(i + 1)
		if i < 15 {
			classes = append(classes, int(asPage(asR(m.t.root).children()[0]).class))
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

// TestPagePromotion makes sure that integer keys keep all their values, in
// any mix of one, a few and many values per key, while their pages change
// shape underneath. It covers the ART behind multimap.Ordered where a key of
// a page, which holds exactly one value per key, gets a second value: the key
// gets a leaf with a value set and a range of its own, and the page's other
// keys stay in pages around it. A key with a leaf can become a node's term
// key, and values and keys are removed again. After every step the map must
// match a reference and satisfy the structural invariants.
func TestPagePromotion(t *testing.T) {
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
	isLeaf := func(k []byte) bool {
		n, _ := m.t.find(k)
		return n.kind == kLeaf
	}
	key := func(a, b byte) []byte { return []byte{9, 9, 9, 9, 9, 9, a, b} }

	// 10 keys with one value each share a U8-1 page; a second value gives a
	// key a leaf, with a page on either side.
	for b := range byte(10) {
		add(key(0, b), 1)
	}
	add(key(0, 3), 1) // the same value again changes nothing
	if m.t.root.kind != kPage {
		t.Fatalf("a repeated value changed the page")
	}
	add(key(0, 3), 2)
	if !isLeaf(key(0, 3)) || m.t.root.kind != kR || m.t.root.count != 3 {
		t.Fatalf("a second value did not give the key a leaf between two pages")
	}
	// many values, and more keys than a K page holds
	for v := range uint64(40) {
		add(key(0, 5), 100+v)
	}
	for b := range byte(40) {
		add(key(1, b), 1)
	}
	add(key(1, 0), 2)  // the first key of a page
	add(key(1, 39), 2) // and the last
	// A shorter key makes a key with a leaf the term of a range node.
	add([]byte{9, 9, 9, 9, 9, 9, 0}, 7, 8)
	for b := range byte(40) {
		add([]byte{9, 9, 9, 9, 9, 9, 0, b, 1}, 1)
	}
	// remove values and whole keys
	for _, k := range [][]byte{key(0, 5), key(0, 3), {9, 9, 9, 9, 9, 9, 0}} {
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
}

// TestPageKLifecycle makes sure that string keys of any length keep all their
// values while the pages that hold them change underneath. It covers the K
// pages of the ART behind multimap.Ordered: a page grows through every class
// as keys arrive and learns a shorter shared prefix when a key does not share
// the one it has, splits into pages below a range node when it overflows,
// finds keys that differ only beyond their first 16 suffix bytes or in
// trailing zero bytes, gives a key a leaf of its own when it gets a second
// value, sends a key too long for any page to a leaf, and shrinks back
// through the classes as keys leave. After every step the map must match a
// reference and satisfy the structural invariants.
func TestPageKLifecycle(t *testing.T) {
	r := rand.New(rand.NewPCG(5, 6))
	var m Map[uint64]
	ref := reference{}
	add := func(k string, vs ...uint64) {
		t.Helper()
		for _, v := range vs {
			m.Add([]byte(k), v)
			ref.add([]byte(k), v)
		}
		compare(t, &m, ref, r)
		checkInvariants(t, &m.t)
	}
	page := func(k string) *pageHead {
		t.Helper()
		n, _ := m.t.find([]byte(k))
		if n == nil || n.kind != kPageK {
			t.Fatalf("key %q is not in a K page", k)
		}
		return asPage(n)
	}
	item := func(i int) string { return fmt.Sprintf("tenant/category/item-%02d/%d", i, 1000+i*7) }

	// Keys that share a page grow it through every class; the first key's
	// page knows it alone, so the second teaches it a shorter prefix.
	var classes []int
	for i := range 14 {
		add(item(i), uint64(i))
		classes = append(classes, int(page(item(i)).class))
	}
	if m.t.root.kind != kPageK || !slices.IsSorted(classes) || classes[0] != 0 || classes[len(classes)-1] != 3 {
		t.Fatalf("classes while growing = %v, want one page growing from class 0 to 3", classes)
	}
	if b := int(asPage(m.t.root).base); b != len("tenant/category/item-") {
		t.Fatalf("shared prefix of %d bytes, want %d", b, len("tenant/category/item-"))
	}
	// More keys overflow the largest page: it splits below a range node.
	for i := 14; i < 40; i++ {
		add(item(i), uint64(i))
	}
	if m.t.root.kind != kR {
		t.Fatalf("the full page did not split: root kind %d", m.t.root.kind)
	}

	// Keys that share their first 16 suffix bytes, or differ in trailing
	// zeros, share head words; the full keys and lengths tell them apart.
	for _, k := range []string{"zz/abcdefghijklmnop", "zz/abcdefghijklmnop1", "zz/abcdefghijklmnop2",
		"zz/abcdefghijklmnop12", "zz/ab", "zz/ab\x00", "zz/ab\x00\x00"} {
		add(k, 1)
	}

	// A key too long for any page goes to a leaf among pages.
	long := "zz/abcdefgh1" + string(bytes.Repeat([]byte{'x'}, maxPageKey))
	add(long, 1, 2)
	if n, _ := m.t.find([]byte(long)); n.kind != kLeaf {
		t.Fatalf("a key of %d bytes is not a leaf", len(long))
	}

	// A second value gives a key a leaf of its own, for a short and for a
	// long suffix; more values go to the leaf.
	for _, k := range []string{item(3), "zz/abcdefghijklmnop12"} {
		add(k, 2, 3, 4)
		if n, _ := m.t.find([]byte(k)); n.kind != kLeaf {
			t.Fatalf("a second value did not give %q a leaf", k)
		}
	}
	for v := range uint64(4) {
		m.Remove([]byte(item(3)), 1+v)
		ref.remove([]byte(item(3)), 1+v)
		compare(t, &m, ref, r)
		checkInvariants(t, &m.t)
	}

	// Removing the keys again shrinks the pages back through the classes.
	shrunk := false
	for i := 39; i >= 0; i-- {
		before := page(item(0)).class
		m.RemoveKey([]byte(item(i)))
		delete(ref, item(i))
		if i > 0 && page(item(0)).class < before {
			shrunk = true
		}
	}
	compare(t, &m, ref, r)
	checkInvariants(t, &m.t)
	if !shrunk {
		t.Fatalf("no page shrank while its keys were removed")
	}
	for _, k := range ref.sortedKeys() {
		m.RemoveKey([]byte(k))
		delete(ref, k)
		checkInvariants(t, &m.t)
	}
	if m.Len() != 0 || m.t.root != nil {
		t.Fatalf("%d keys left", m.Len())
	}
}
