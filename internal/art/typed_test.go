package art

import (
	"bytes"
	"fmt"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

// TestTypedLeafLayout makes sure that every typed leaf class holds its values
// and its key where the code reads them, and that the garbage collector keeps
// what the values point to. It covers the allocation of typed leaves, one Go
// type for each value capacity and key area size: for every capacity and every
// key remainder from none to the longest, a leaf is filled to the brim with
// strings that only it refers to, the collector runs, and the strings and the
// key must be intact.
func TestTypedLeafLayout(t *testing.T) {
	for c := 1; c < len(typedCaps); c++ {
		for klen := 0; klen <= maxTypedKey; klen++ {
			key := bytes.Repeat([]byte{byte(klen) + 1}, klen)
			l := typedWithKey[string](uint8(c), key, klen+3)
			if l.cls() != uint8(c) || l.n != 0 || !bytes.Equal(l.stored(), key) || l.keyLen() != klen+3 {
				t.Fatalf("class %d, key of %d bytes: leaf of class %d holds %d values and %q", c, klen, l.cls(), l.n, l.stored())
			}
			for i := range typedCaps[c] {
				appendTyped(l, fmt.Sprintf("value %d of class %d, key %d", i, c, klen))
			}
			runtime.GC()
			for i, v := range typedVals[string](l) {
				if want := fmt.Sprintf("value %d of class %d, key %d", i, c, klen); v != want {
					t.Fatalf("class %d, key of %d bytes: value %d is %q, want %q", c, klen, i, v, want)
				}
			}
			if !bytes.Equal(l.stored(), key) {
				t.Fatalf("class %d: the values overwrote the key of %d bytes: %q", c, klen, l.stored())
			}
		}
	}
}

// TestTypedLeafLife makes sure a key of multimap.Ordered with string values
// keeps all of them while it grows from one value to more than a leaf holds
// and shrinks back, and that its leaf follows: it stays a compact leaf up to
// 16 values, becomes a set leaf beyond, and compact again once the values fill
// half of the largest class. It covers the typed leaves behind such a map:
// growth through the classes, the spill into a set leaf, the return from it,
// and the shrinking of a leaf whose values have gone.
func TestTypedLeafLife(t *testing.T) {
	var m Map[string]
	key := []byte("key")
	val := func(i int) string { return fmt.Sprint("value ", i) }
	check := func(n int) {
		t.Helper()
		got := valuesOf(&m, key)
		slices.Sort(got)
		want := make([]string, n)
		for i := range want {
			want[i] = val(i)
		}
		slices.Sort(want)
		if !slices.Equal(got, want) {
			t.Fatalf("%d values: got %q", n, got)
		}
		if n > 0 {
			largest := typedCaps[len(typedCaps)-1]
			l := findLeaf(&m.t, key)
			switch typed := l.kind != kSet; {
			case n > largest && typed:
				t.Fatalf("%d values in a typed leaf of class %d", n, l.cls())
			case n <= largest/2 && !typed:
				t.Fatalf("%d values in a set leaf, want a typed leaf", n)
			case typed && (int(l.n) != n || typedCaps[l.cls()] < n):
				t.Fatalf("typed leaf of class %d (room for %d) holds %d values, want %d", l.cls(), typedCaps[l.cls()], l.n, n)
			}
		}
		runtime.GC()
	}
	for i := range 20 {
		m.Add(key, val(i))
		m.Add(key, val(i)) // a value that is there already changes nothing
		check(i + 1)
	}
	for i := 19; i >= 0; i-- {
		m.Remove(key, val(i))
		m.Remove(key, val(i)) // and removing it twice neither
		check(i)
	}
	if m.Len() != 0 {
		t.Fatalf("%d keys left", m.Len())
	}
}

// TestTypedLeafForgets makes sure a leaf that loses a value does not keep
// holding on to it, which would stop the garbage collector from freeing what
// the value points to. It covers the removal from a typed leaf: the slot the
// last value moved out of is empty again.
func TestTypedLeafForgets(t *testing.T) {
	var m Map[string]
	key := []byte("key")
	for i := range 5 {
		m.Add(key, fmt.Sprint("value ", i))
	}
	m.Remove(key, "value 1")
	l := findLeaf(&m.t, key)
	slots := unsafe.Slice((*string)(unsafe.Add(unsafe.Pointer(l), typedOff(int(l.klen)))), typedCaps[l.cls()])
	if l.n != 4 || slots[4] != "" {
		t.Fatalf("leaf holds %d values and %q after the last one moved, want 4 and nothing", l.n, slots[4])
	}
}

// TestTypedRekey makes sure that a typed leaf which has to hold more of its
// key, because the node above it went away, keeps all its values and
// stays where it is as long as the longer remainder leaves its values in
// place. It covers rekeyTyped: in place within the same key area, into a new
// typed leaf when the key area grows, into a set leaf when the remainder no
// longer fits a typed leaf, and a set leaf that a typed map holds.
func TestTypedRekey(t *testing.T) {
	key := bytes.Repeat([]byte("abcdefghij"), 12) // 120 bytes
	for _, tc := range []struct {
		name         string
		keyLen, base int
		to           int
		values       int
		want         string // "in place", "typed" or "set"
	}{
		{"within the same key area", 20, 17, 11, 1, "in place"},
		{"with several values", 20, 17, 11, 3, "in place"},
		{"into a longer key area", 40, 37, 24, 3, "typed"},
		{"to the longest remainder", 80, 79, 22, 2, "typed"},
		{"beyond the longest remainder", 100, 99, 20, 2, "set"},
		{"a set leaf", 100, 99, 20, 40, "set leaf"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			k := key[:tc.keyLen]
			var l *leafHead
			if tc.values > typedCaps[len(typedCaps)-1] {
				l = newSetLeaf[string](k, tc.base)
				for i := range tc.values {
					vals[string](l).Add(fmt.Sprint(i))
				}
			} else {
				l = typedWithKey[string](typedClassFor(tc.values), k[tc.base:], tc.keyLen)
				for i := range tc.values {
					appendTyped(l, fmt.Sprint(i))
				}
			}
			nl := rekeyTyped[string](l, k[:tc.base-1], int(k[tc.base-1]), tc.to)
			var got []string
			if nl.kind == kSet {
				vals[string](nl).Each(func(v string) bool { got = append(got, v); return true })
			} else {
				got = typedVals[string](nl)
			}
			slices.Sort(got)
			want := make([]string, tc.values)
			for i := range want {
				want[i] = fmt.Sprint(i)
			}
			slices.Sort(want)
			if !slices.Equal(got, want) {
				t.Errorf("values %q, want %q", got, want)
			}
			if nl.keyLen() != tc.keyLen || !bytes.Equal(nl.from(tc.to), k[tc.to:]) {
				t.Errorf("leaf holds %q from %d of a key of %d bytes, want the key from %d on", nl.stored(), nl.base(), nl.keyLen(), tc.to)
			}
			switch kindOf := map[bool]string{true: "set", false: "typed"}[nl.kind == kSet]; {
			case tc.want == "in place" && nl != l, tc.want == "typed" && (nl == l || kindOf != "typed"), tc.want == "set" && kindOf != "set", tc.want == "set leaf" && kindOf != "set":
				t.Errorf("rekey gave a %s leaf, same leaf: %v, want %s", kindOf, nl == l, tc.want)
			}
		})
	}
}

// TestTypedLeafTail makes sure that the scan of a map with typed leaves
// touches a byte that every leaf has. It covers the offset of the last byte of
// the smallest leaf the map creates, which a range query reads ahead.
func TestTypedLeafTail(t *testing.T) {
	m := Map[string]{flat: 2}
	size := unsafe.Sizeof(typedLeaf[string, [1]string, [2]byte]{})
	if tail := m.leafTail(); tail != size-1 {
		t.Fatalf("tail offset %d, want the last byte of the smallest typed leaf of %d bytes", tail, size)
	}
}

// typedOffsets returns where a typed leaf of key area K keeps its key and its
// values.
func typedOffsets[K keyArr]() (key, values, size uintptr) {
	var l typedLeaf[string, [1]string, K]
	return unsafe.Offsetof(l.k), unsafe.Offsetof(l.v), unsafe.Sizeof(l)
}

// TestTypedLeafOffsets makes sure the leaf types put the key right after the
// head and the values where typedOff says, for each key area. It covers the
// layout contract of typed leaves: the code reads the values at typedOff of
// the key remainder's length, whichever key area the leaf was allocated with.
func TestTypedLeafOffsets(t *testing.T) {
	check := func(name string, klen int, key, values, size uintptr) {
		if key != keyOff || values != typedOff(klen) || size != values+unsafe.Sizeof("") {
			t.Errorf("key area of %s: key at %d, values at %d, %d bytes; want %d, %d, %d", name, key, values, size, keyOff, typedOff(klen), typedOff(klen)+unsafe.Sizeof(""))
		}
	}
	k, v, s := typedOffsets[[2]byte]()
	check("2", 2, k, v, s)
	k, v, s = typedOffsets[[10]byte]()
	check("10", 10, k, v, s)
	k, v, s = typedOffsets[[18]byte]()
	check("18", 18, k, v, s)
	k, v, s = typedOffsets[[26]byte]()
	check("26", 26, k, v, s)
	k, v, s = typedOffsets[[34]byte]()
	check("34", 34, k, v, s)
	k, v, s = typedOffsets[[42]byte]()
	check("42", 42, k, v, s)
	k, v, s = typedOffsets[[50]byte]()
	check("50", 50, k, v, s)
	k, v, s = typedOffsets[[58]byte]()
	check("58", 58, k, v, s)
}

// TestTypedLeafHovers makes sure a key that goes back and forth between one
// and a few values does not move its leaf every time. It covers the smallest
// class a typed leaf grows into or shrinks back to: once a key has a second
// value its leaf holds four, and it keeps that leaf while its values come and
// go within them.
func TestTypedLeafHovers(t *testing.T) {
	var m Map[string]
	key := []byte("key")
	m.Add(key, "a")
	if c := findLeaf(&m.t, key).cls(); typedCaps[c] != 1 {
		t.Fatalf("a key with one value has a leaf of room for %d", typedCaps[c])
	}
	m.Add(key, "b")
	l := findLeaf(&m.t, key)
	if typedCaps[l.cls()] != typedCaps[minGrownTyped] {
		t.Fatalf("a key with two values has a leaf of room for %d, want %d", typedCaps[l.cls()], typedCaps[minGrownTyped])
	}
	for range 10 {
		m.Add(key, "c")
		m.Add(key, "d")
		m.Remove(key, "b")
		m.Remove(key, "c")
		m.Remove(key, "d")
		m.Add(key, "b")
		if findLeaf(&m.t, key) != l {
			t.Fatal("a key that hovers within four values moved its leaf")
		}
	}
}
