package art

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"slices"
	"testing"
	"unsafe"

	"github.com/TomTonic/multimap/internal/vpage"
)

// TestPageLifecycle makes sure that integer keys with one value each stay
// compact however their number changes. It covers the pages of the ART behind
// multimap.Ordered: a page grows through its classes as keys arrive, splits
// into two pages below a range node when it overflows, shrinks back through the
// classes and merges with its neighbour as keys leave, moves up when its
// sibling is gone, and every key keeps its value throughout.
func TestPageLifecycle(t *testing.T) {
	var keys [][]byte // 6 shared bytes, then two groups of 20
	for a := range 2 {
		for b := range 20 {
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
	class, split := -1, 0
	for i, k := range keys {
		m.Add(k, uint64(i))
		checkInvariants(t, &m.t)
		if r := m.t.root; !isMultiKey(r.objType) {
			if split == 0 {
				split = i + 1 // the number of keys the page held when it overflowed, plus one
			}
			continue
		} else if c := asMultiKey(r).Class(); c < class {
			t.Fatalf("after %d keys the root page shrank from class %d to %d", i+1, class, c)
		} else {
			class = c
		}
	}
	if class != 2 || split < 20 || split > 40 {
		t.Fatalf("the page grew to class %d and overflowed at key %d, want class 2 and a split between the 20th and 40th key", class, split)
	}
	r := m.t.root
	if !isRange(r.objType) || r.plen != 6 {
		t.Fatalf("after the split: root type %d, path %d; want a range node, path 6", r.objType, r.plen)
	}
	pages := 0
	for _, c := range asR(r).children()[:asR(r).n] {
		if !isMultiKey(c.objType) {
			t.Fatalf("after the split: child type %d, want a page", c.objType)
		}
		pages++
	}
	if pages < 2 {
		t.Fatalf("after the split: %d pages, want at least 2", pages)
	}
	checkValues(0)

	for i, k := range keys[:20] {
		m.RemoveKey(k)
		checkInvariants(t, &m.t)
		checkValues(i + 1)
	}
	if m.Len() != 20 {
		t.Fatalf("after emptying one group: %d keys, want the 20 of the other", m.Len())
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
	if got, want := isMultiKey(m.t.root.objType), takesPages[T](); got != want {
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
// keys stay in pages around it. A key with a leaf can become a node's end page
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
	leafAt := func(k []byte) bool { return findSingleKey(&m.t, k) != nil }
	key := func(a, b byte) []byte { return []byte{9, 9, 9, 9, 9, 9, a, b} }

	// 10 keys with one value each share a U8-1 page; a second value gives a
	// key a leaf, with a page on either side.
	for b := range byte(10) {
		add(key(0, b), 1)
	}
	add(key(0, 3), 1) // the same value again changes nothing
	if !isMultiKey(m.t.root.objType) {
		t.Fatalf("a repeated value changed the page")
	}
	add(key(0, 3), 2)
	if !leafAt(key(0, 3)) || !isRange(m.t.root.objType) || asR(m.t.root).n != 3 {
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
	// A shorter key makes a key with a leaf the end page of a range node.
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
// multimap.Ordered, which the tree reads through the type byte at their start:
// a page of class c has the type kMultiKey+c, the header is 8 bytes, and the range
// nodes, which hold pages, have the layout their offsets say.
func TestPageLayout(t *testing.T) {
	for class := range 4 {
		h := (*header)(unsafe.Pointer(vpage.New(class, 0)))
		if k := asMultiKey(h).Class(); k != class || h.objType != kMultiKey+objType(class)<<1 || !isMultiKey(h.objType) || isSingleKey(h.objType) {
			t.Errorf("page of class %d has class %d, type %d", class, k, h.objType)
		}
	}
	for _, tc := range []struct {
		name      string
		got, want uintptr
	}{
		{"page header", unsafe.Sizeof(vpage.Page{}), 8},
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

// TestPageLongRemainder makes sure that a key with as much below its page as a page
// holds, and one with more, keep their values, also when they get a second
// one. It covers the limits of the pages of the ART behind multimap.Ordered
// (255 bytes below the base) and the leaf a key gets that is too long for a page,
// a flat leaf up to 254 bytes of its key and a value overflow beyond.
func TestPageLongRemainder(t *testing.T) {
	r := rand.New(rand.NewPCG(5, 6))
	var m Map[uint64]
	ref := reference{}
	// a key with more than 255 bytes, added to a root page of short keys
	for _, k := range []string{"abc1", "abc2", "abd", "ab"} {
		m.Add([]byte(k), 1)
		ref.add([]byte(k), 1)
	}
	m.Add(bytes.Repeat([]byte("abc"), 100), 1)
	ref.add(bytes.Repeat([]byte("abc"), 100), 1)
	compare(t, &m, ref, id, r)
	checkInvariants(t, &m.t)
	m.Clear()
	ref = reference{}
	key := func(n int, last byte) []byte { return append(bytes.Repeat([]byte("k"), n-1), last) }
	m.Add(key(255, 'a'), 1) // a key with 255 bytes below the root page, the most it holds,
	m.Add(key(255, 'b'), 1)
	m.Add(key(255, 'a'), 2) // gets a second value: a value overflow, as the key is too long for a flat one
	ref.add(key(255, 'a'), 1)
	ref.add(key(255, 'b'), 1)
	ref.add(key(255, 'a'), 2)
	compare(t, &m, ref, id, r)
	checkInvariants(t, &m.t)
	for _, k := range [][]byte{key(255, 'a'), key(255, 'b'), key(300, 'a'), key(254, 'c'), key(256, 'd')} {
		m.Add(k, 1)
		ref.add(k, 1)
		checkInvariants(t, &m.t)
	}
	for _, k := range [][]byte{key(255, 'a'), key(300, 'a'), key(254, 'c'), key(256, 'd')} {
		m.Add(k, 2) // a second value: the key leaves its page, or already has a leaf
		ref.add(k, 2)
		compare(t, &m, ref, id, r)
		checkInvariants(t, &m.t)
	}
	for _, k := range [][]byte{key(255, 'a'), key(255, 'b'), key(300, 'a'), key(254, 'c'), key(256, 'd')} {
		m.RemoveKey(k)
		delete(ref, string(k))
		compare(t, &m, ref, id, r)
		checkInvariants(t, &m.t)
	}
}

// TestPageMovesUp makes sure that a page whose range node goes away takes the
// bytes of the node's path into its keys, as far as they fit it, and stays
// below the node when they do not. It covers the collapse of range nodes in the
// ART behind multimap.Ordered: the page of the one child left holds its keys
// from a base below the node's path, so it must start higher up (see pageUp).
func TestPageMovesUp(t *testing.T) {
	path := bytes.Repeat([]byte("p"), 40)
	for _, tc := range []struct {
		name   string
		keys   int // keys of 8 bytes in each of two families below the path
		pinned bool
	}{
		{"a page with room takes the path", 3, false},
		{"a full page cannot, and stays below the node", 29, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var m Map[uint64]
			ref := reference{}
			key := func(fam byte, i int) []byte {
				return slices.Concat(path, []byte{fam}, binary.BigEndian.AppendUint64(nil, uint64(i))[1:])
			}
			for fam := byte(0); fam < 2; fam++ {
				for i := range tc.keys {
					m.Add(key(fam, i), uint64(i))
					ref.add(key(fam, i), uint64(i))
				}
			}
			checkInvariants(t, &m.t)
			for i := range tc.keys { // the first family goes: the node has one child
				m.RemoveKey(key(0, i))
				delete(ref, string(key(0, i)))
			}
			checkInvariants(t, &m.t)
			compare(t, &m, ref, id, rand.New(rand.NewPCG(7, 8)))
			if got := isRange(m.t.root.objType); got != tc.pinned {
				t.Fatalf("root is a range node: %v, want %v", got, tc.pinned)
			}
		})
	}
}

// TestRangeNodeLongPath makes sure that a range node with a long path keeps its
// ranges and children when another key splits its path. It covers the ART
// behind multimap.Ordered, whose range nodes carry the shared bytes of keys of
// any length: the node is copied into a new object when its path changes the
// tail class, and every class of range nodes (8, 24, 56 and 256 ranges) must
// survive the copy.
func TestRangeNodeLongPath(t *testing.T) {
	for _, fan := range []int{4, 20, 50, 250} {
		t.Run(fmt.Sprintf("%d groups", fan), func(t *testing.T) {
			r := rand.New(rand.NewPCG(uint64(fan), 9))
			var m Map[uint64]
			ref := reference{}
			prefix := []byte("a-path-of-twenty-bytes/")
			add := func(k []byte) {
				m.Add(k, 1)
				ref.add(k, 1)
			}
			for g := range fan {
				for range 40 {
					add(slices.Concat(prefix, []byte{byte(g)}, binary.BigEndian.AppendUint32(nil, r.Uint32())[1:]))
				}
			}
			checkInvariants(t, &m.t)
			add(slices.Concat(prefix[:11], []byte("-splits-the-path"))) // leaves the node a path of 12 bytes: no tail
			add(slices.Concat(prefix[:3], []byte("-and-again")))        // and a path of 8 above it
			compare(t, &m, ref, id, r)
			checkInvariants(t, &m.t)
		})
	}
}
