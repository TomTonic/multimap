package art

import (
	"bytes"
	"fmt"
	"slices"
	"strings"
	"testing"
	"unsafe"

	set3 "github.com/TomTonic/Set3"
	"github.com/TomTonic/multimap/internal/skpage"
)

// inOrder returns the strings of s in order.
func inOrder(s []string) []string { slices.Sort(s); return s }

// isSKPage reports whether leaf l of a string map is a single-key page.
func isSKPage(l *leafHead) bool { return l.kind != kSet }

// TestStringMapUsesPages covers what a user of multimap.Ordered with string
// values gets for each key: one single-key page (SKMV) with the key's end and
// its values as bytes, not a leaf with string headers. The test fills a map and
// checks the kind of every key's object and what the object statistic calls it.
func TestStringMapUsesPages(t *testing.T) {
	var m Map[string]
	for i := range 300 {
		key := []byte(fmt.Sprint("key-", i))
		for j := range i%4 + 1 {
			m.Add(key, fmt.Sprint("value ", j))
		}
	}
	for i := range 300 {
		key := []byte(fmt.Sprint("key-", i))
		l := findLeaf(&m.t, key)
		if l == nil || !isSKPage(l) || l.kind < kSet+1 || l.kind > kSet+kind(skpage.Classes) {
			t.Fatalf("key %q has no single-key page", key)
		}
		if got := inOrder(valuesOf(&m, key)); len(got) != i%4+1 {
			t.Fatalf("key %q has values %q", key, got)
		}
	}
	pages := 0
	m.Objects(func(o Object) {
		if o.Label == "single-key page" {
			pages++
		}
	})
	if pages != 300 {
		t.Errorf("the statistic counts %d single-key pages, want 300", pages)
	}
	if tail := m.leafTail(); tail != 31 {
		t.Errorf("leafTail = %d, want 31 (the smallest page)", tail)
	}
}

// TestStringKeyOverflow covers a key with more values than a page holds. Its
// values go into a set leaf (the value overflow) when the page is full and come
// back into a page once they fill no more than half of the largest class, so
// that a key at the border does not change its object with every added and
// removed value. The test adds values of 30 bytes up to 40, then removes them
// and checks at every step which object the key has and that no value is lost.
func TestStringKeyOverflow(t *testing.T) {
	var m Map[string]
	key := []byte("key")
	val := func(i int) string { return fmt.Sprintf("value-%03d-%s", i, strings.Repeat("x", 20)) } // 30 bytes
	content := func(n int) int { return skpage.Header + 3 + n*(1+len(val(0))) }
	var want []string
	for i := range 40 {
		m.Add(key, val(i))
		want = append(want, val(i))
		isSet := findLeaf(&m.t, key).kind == kSet
		if wantSet := content(i+1) > 512; isSet != wantSet {
			t.Fatalf("after %d values (content %d): set leaf = %v", i+1, content(i+1), isSet)
		}
		if !slices.Equal(inOrder(valuesOf(&m, key)), want) {
			t.Fatalf("after %d values: %q", i+1, valuesOf(&m, key))
		}
	}
	for i := 39; i > 0; i-- {
		m.Remove(key, val(i))
		want = want[:i]
		isSet := findLeaf(&m.t, key).kind == kSet
		// the set leaf stays until the values fill 256 bytes or less; the page, once back, until it is empty
		if isSet && content(i) <= skpage.BackLimit {
			t.Fatalf("with %d values (content %d) the key still has a set leaf", i, content(i))
		}
		if !isSet && content(i) > 512 {
			t.Fatalf("with %d values (content %d) the key has a page", i, content(i))
		}
		if !slices.Equal(inOrder(valuesOf(&m, key)), want) {
			t.Fatalf("after removing down to %d values: %q", i, valuesOf(&m, key))
		}
	}
	// the border: between 256 and 512 bytes the key keeps whichever object it has
	m.Add(key, val(1))
	if findLeaf(&m.t, key).kind == kSet {
		t.Fatal("a key with two values has a set leaf")
	}
	m.Remove(key, val(0))
	m.Remove(key, val(1))
	if m.Has(key) {
		t.Fatal("the key must be gone with its last value")
	}
}

// TestStringValueTooLong covers a string of 255 bytes or more as a value:
// such a value cannot go into a page, so the key becomes a set leaf with all its
// values, and a page again once the long value is gone and the rest is small.
func TestStringValueTooLong(t *testing.T) {
	var m Map[string]
	key := []byte("key")
	long := strings.Repeat("L", 300)
	m.Add(key, "a")
	m.Add(key, "b")
	m.Add(key, long)
	if findLeaf(&m.t, key).kind != kSet {
		t.Fatal("a value of 300 bytes must make a set leaf")
	}
	if got := inOrder(valuesOf(&m, key)); !slices.Equal(got, inOrder([]string{"a", "b", long})) {
		t.Fatalf("got %d values", len(got))
	}
	m.Remove(key, "a") // the set still holds the long value: it stays a set leaf
	if findLeaf(&m.t, key).kind != kSet {
		t.Fatal("with the long value in it the key must stay a set leaf")
	}
	m.Remove(key, long)
	if findLeaf(&m.t, key).kind == kSet || !slices.Equal(valuesOf(&m, key), []string{"b"}) {
		t.Fatalf("the key must be a page with the value b, has %q", valuesOf(&m, key))
	}
}

// TestStringKeyTooLong covers keys whose end does not fit a page: a lone key of
// 300 bytes has a remainder of 300, a key of more than 64 KiB cannot even
// have its length in the page's header. Both get a set leaf that holds the key
// as a string, and keep their values, also when values go.
func TestStringKeyTooLong(t *testing.T) {
	for _, n := range []int{255, 300, maxKeyLen + 10} {
		t.Run(fmt.Sprint(n, " bytes"), func(t *testing.T) {
			var m Map[string]
			key := bytes.Repeat([]byte("k"), n)
			for _, v := range []string{"a", "b", "c"} {
				m.Add(key, v)
			}
			if findLeaf(&m.t, key).kind != kSet {
				t.Fatal("expected a set leaf")
			}
			m.Remove(key, "b")
			if got := inOrder(valuesOf(&m, key)); !slices.Equal(got, []string{"a", "c"}) {
				t.Fatalf("got %q", got)
			}
		})
	}
	t.Run("the longest remainder a page holds", func(t *testing.T) {
		var m Map[string]
		key := bytes.Repeat([]byte("k"), skpage.MaxRemainder)
		m.Add(key, "a")
		if findLeaf(&m.t, key).kind == kSet {
			t.Fatal("a remainder of 254 bytes fits a page")
		}
	})
}

// TestStringSetStaysSet covers a set leaf that cannot become a page: one with
// too many values for a page (every value takes two bytes at least; 255 values
// of one byte are the most there can be), however few of them go.
func TestStringSetStaysSet(t *testing.T) {
	var m Map[string]
	key := []byte("key")
	one := func(i int) string { return string([]byte{byte(i)}) }
	for i := range 255 {
		m.Add(key, one(i))
	}
	m.Remove(key, one(0))
	if findLeaf(&m.t, key).kind != kSet || len(valuesOf(&m, key)) != 254 {
		t.Fatal("254 values of one byte are a set leaf")
	}
}

// TestStringRekey covers what happens to the object of a key when the node above
// it goes away and its path gets shorter (see rekeyFunc): a page takes the bytes
// in front of its remainder, in place if they fit its class, in a page of a larger
// class, or, when the remainder would be longer than a page holds, as a
// set leaf; a set leaf is left to the set leaf code. The values and the key
// stay.
func TestStringRekey(t *testing.T) {
	key := bytes.Repeat([]byte("abcdefghij"), 30) // 300 bytes
	for _, tc := range []struct {
		name         string
		keyLen, base int
		to           int
		values       []string
		want         string // "in place", "page" or "set"
	}{
		{"fits the class", 40, 37, 30, []string{"v"}, "in place"},
		{"grows into a larger class", 40, 37, 10, []string{strings.Repeat("v", 17)}, "page"},
		{"a remainder of 254 bytes", 280, 270, 26, []string{"v", "w"}, "page"},
		{"a remainder beyond 254 bytes", 290, 280, 20, []string{"v", "w"}, "set"},
		{"a set leaf", 40, 37, 30, []string{strings.Repeat("L", 300)}, "set"},
		{"a set leaf that needs a larger key area", 100, 99, 20, []string{strings.Repeat("L", 300)}, "set"},
		{"a set leaf in the largest key area", 200, 80, 70, []string{strings.Repeat("L", 300)}, "set"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			k := key[:tc.keyLen]
			l := newSK(k, tc.base)
			slot := leafHdr(l)
			for _, v := range tc.values {
				if l.kind == kSet {
					strSetAdd(l, v)
				} else {
					addSK(&slot, l, k, v)
					l = asLeaf(slot)
				}
			}
			nl := rekeySK(l, k[:tc.base-1], int(k[tc.base-1]), tc.to)
			var got []string
			if nl.kind == kSet {
				strSetEach(nl, func(v string) bool { got = append(got, v); return true })
			} else {
				asSK(nl).Strings(func(v string) bool { got = append(got, v); return true })
			}
			if !slices.Equal(inOrder(got), inOrder(slices.Clone(tc.values))) {
				t.Errorf("values %d, want %d", len(got), len(tc.values))
			}
			if nl.keyLen() != tc.keyLen || !bytes.Equal(nl.from(tc.to), k[tc.to:]) {
				t.Errorf("leaf holds %q from %d of a key of %d bytes, want the key from %d on", nl.stored(), nl.base(), nl.keyLen(), tc.to)
			}
			switch {
			case tc.want == "in place" && nl != l, tc.want == "page" && (nl == l || nl.kind == kSet), tc.want == "set" && nl.kind != kSet:
				t.Errorf("rekey gave kind %d, same leaf: %v, want %s", nl.kind, nl == l, tc.want)
			}
		})
	}
}

// TestStringSetLeafLayout covers the set leaf of a string map as an object: it
// lies on the grid of the pages (32, 64, 128 or 256 bytes, 32 for a key held as
// a string), the pointer to its value set is the last word of the object whatever
// the key area, the key area is the one that fills the object, and the object
// statistic knows its size, which is the size the runtime allocates.
func TestStringSetLeafLayout(t *testing.T) {
	var strs Map[string]
	strs.flat = 3
	for _, tc := range []struct {
		n, size int
	}{{5, 32}, {18, 32}, {19, 64}, {50, 64}, {51, 128}, {114, 128}, {115, 256}, {242, 256}, {243, 32}, {300, 32}} {
		key := bytes.Repeat([]byte("k"), tc.n)
		l := newSetLeaf3(key, 0, set3.EmptyWithCapacity[string](4))
		if off := int(uintptr(unsafe.Pointer(strSetOf(l))) - uintptr(unsafe.Pointer(l))); tc.n <= maxInline3 && off != tc.size-8 {
			t.Errorf("key of %d bytes: set at %d, want the last word of %d", tc.n, off, tc.size)
		}
		strSetAdd(l, "v")
		if got := strs.leafValues(l); got != 1 {
			t.Errorf("key of %d bytes: %d values", tc.n, got)
		}
		o := strs.leafKind(l)
		if o.Size != tc.size || int(setLeaf3Size(l.klen)) != tc.size || o.Label != "set leaf" || !o.Pointers {
			t.Errorf("key of %d bytes: object %+v, want a set leaf of %d bytes", tc.n, o, tc.size)
		}
		if string(l.stored()) != string(key) && tc.n <= maxInline3 || l.keyLen() != tc.n {
			t.Errorf("key of %d bytes: stored %q, length %d", tc.n, l.stored(), l.keyLen())
		}
	}
}
