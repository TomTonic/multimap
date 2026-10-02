package art

import (
	"math/rand/v2"
	"slices"
	"testing"
	"unsafe"
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
			if v := valuesOf(&m, k); !slices.Equal(v, []uint64{uint64(from + i)}) {
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
	if !isRange(r.kind) || r.plen != 6 || asR(r).n != 2 {
		t.Fatalf("after the split: root kind %d, path %d, %d ranges; want a range node, path 6, 2 ranges", r.kind, r.plen, asR(r).n)
	}
	for _, c := range asR(r).children()[:2] {
		if c.kind != kPage || asPage(c).count != 16 {
			t.Fatalf("after the split: child kind %d, want a page of 16 keys", c.kind)
		}
	}
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

// TestPageType makes sure that only values that fit a page's 8-byte word
// and hold no pointers go into pages, where the garbage collector never looks.
// It covers the value-type check of the ART behind multimap.Ordered.
func TestPageType(t *testing.T) {
	type pair struct{ A, B uint32 }
	type withPtr struct{ P *int }
	for _, tc := range []struct {
		name string
		got  bool
		want bool
	}{
		{"uint64 goes into pages", takesPages[uint64](), true},
		{"int8 goes into pages", takesPages[int8](), true},
		{"float32 goes into pages", takesPages[float32](), true},
		{"bool goes into pages", takesPages[bool](), true},
		{"a struct of two uint32 goes into pages", takesPages[pair](), true},
		{"an array of four uint16 goes into pages", takesPages[[4]uint16](), true},
		{"a string does not", takesPages[string](), false},
		{"a pointer does not", takesPages[*int](), false},
		{"a struct with a pointer does not", takesPages[withPtr](), false},
		{"an array of pointers does not", takesPages[[1]*int](), false},
		{"a 16-byte value does not", takesPages[complex128](), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.got != tc.want {
				t.Fatalf("takesPages = %v, want %v", tc.got, tc.want)
			}
		})
	}
}

// takesPages reports whether a map of T holds keys with one value in pages.
func takesPages[T comparable]() bool {
	var m Map[T]
	m.decide()
	return m.t.small
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
	if got, want := m.t.root.kind == kPage, takesPages[T](); got != want {
		t.Fatalf("root is a page: %v, want %v", got, want)
	}
	var got []T
	m.RangeValues(&Bounds{}, func(v T) bool { got = append(got, v); return true })
	if !slices.Equal(got, vs) {
		t.Fatalf("RangeValues = %v, want %v", got, vs)
	}
	for i, v := range vs {
		if s := valuesOf(&m, key(i)); !slices.Equal(s, []T{v}) {
			t.Fatalf("Values(key %d) = %v, want exactly %v", i, s, v)
		}
	}
	m.Add(key(0), vs[1]) // a second value
	checkInvariants(t, &m.t)
	if s := valuesOf(&m, key(0)); len(s) != 2 {
		t.Fatalf("key 0 holds %d values after a second one, want 2", len(s))
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
		compare(t, &m, ref, id, r)
		checkInvariants(t, &m.t)
	}
	leafAt := func(k []byte) bool { return findLeaf(&m.t, k) != nil }
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
	if !leafAt(key(0, 3)) || !isRange(m.t.root.kind) || asR(m.t.root).n != 3 {
		t.Fatalf("a second value did not give the key a leaf between two pages")
	}
	// many values, and more keys than a leaf's flat values hold
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
			compare(t, &m, ref, id, r)
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

// TestPageLayout makes sure that every page class has the size and the layout
// the code that reaches into it assumes. It covers the pages of the ART behind
// multimap.Ordered, which are read and written through offsets, not fields: a
// page fills one Go size class, starts with its kind, and keeps its arrays one
// after another, each key's data at the same index in every array.
func TestPageLayout(t *testing.T) {
	for _, tc := range []struct {
		name      string
		got, want uintptr
	}{
		{"page head", unsafe.Sizeof(pageHead{}), 16},
		{"page of 3 keys", unsafe.Sizeof(page3{}), 64},
		{"page of 7 keys", unsafe.Sizeof(page7{}), 128},
		{"page of 15 keys", unsafe.Sizeof(page15{}), 256},
		{"page of 31 keys", unsafe.Sizeof(page31{}), 512},
		{"heads of a page of 7 keys", unsafe.Offsetof(page7{}.heads), headsOff},
		{"values of a page of 7 keys", unsafe.Offsetof(page7{}.vals), headsOff + 7*8},
		{"heads of a page of 3 keys", unsafe.Offsetof(page3{}.heads), headsOff},
		{"heads of a page of 15 keys", unsafe.Offsetof(page15{}.heads), headsOff},
		{"heads of a page of 31 keys", unsafe.Offsetof(page31{}.heads), headsOff},
		{"values of a page of 3 keys", unsafe.Offsetof(page3{}.vals), headsOff + 3*8},
		{"values of a page of 15 keys", unsafe.Offsetof(page15{}.vals), headsOff + 15*8},
		{"values of a page of 31 keys", unsafe.Offsetof(page31{}.vals), headsOff + 31*8},
		{"kind at the start of a page", unsafe.Offsetof(pageHead{}.kind), 0},
		{"head of a range node of 8 ranges", unsafe.Offsetof(rnode8{}.rhead), 0},
		{"children of a range node of 8 ranges", unsafe.Offsetof(rnode8{}.child), rChildOff},
		{"head of a range node of 24 ranges", unsafe.Offsetof(rnode24{}.rhead), 0},
		{"children of a range node of 24 ranges", unsafe.Offsetof(rnode24{}.child), rChildOff},
		{"head of a range node of 56 ranges", unsafe.Offsetof(rnode56{}.rhead), 0},
		{"children of a range node of 56 ranges", unsafe.Offsetof(rnode56{}.child), rChildOff},
		{"head of a range node of 256 ranges", unsafe.Offsetof(rnode256{}.rhead), 0},
		{"children of a range node of 256 ranges", unsafe.Offsetof(rnode256{}.child), rChildOff},
		{"range node head", unsafe.Sizeof(rhead{}), 56},
		{"range node of 8 ranges", unsafe.Sizeof(rnode8{}), 128},
		{"range node of 24 ranges", unsafe.Sizeof(rnode24{}), 256},
		{"range node of 56 ranges", unsafe.Sizeof(rnode56{}), 512},
		{"range node of 256 ranges", unsafe.Sizeof(rnode256{}), 2112},
	} {
		if tc.got != tc.want {
			t.Errorf("%s: %d bytes, want %d", tc.name, tc.got, tc.want)
		}
	}
}
