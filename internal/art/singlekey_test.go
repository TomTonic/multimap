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
func isSKPage(l *singleKeyHead) bool { return !l.isValueOverflow() }

// TestStringMapUsesPages covers what a user of multimap.Ordered with string
// values gets for each key: one single-key page (SKMV) with the key's end and
// its values as bytes, not a leaf with string headers. The test fills a map and
// checks the type of every key's object and what the object statistic calls it.
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
		l := m.t.find(key)
		if l == nil || !isSKPage(l) || l.cls() < 1 || int(l.cls()) > skpage.Classes {
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
	if leafTail != 31 {
		t.Errorf("leafTail = %d, want 31 (the smallest page)", leafTail)
	}
}

// TestStringKeyOverflow covers a key with more values than a page holds. Its
// values go into a value overflow (the value overflow) when the page is full and come
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
		isValueOverflow := m.t.find(key).isValueOverflow()
		if wantSet := content(i+1) > 512; isValueOverflow != wantSet {
			t.Fatalf("after %d values (content %d): value overflow = %v", i+1, content(i+1), isValueOverflow)
		}
		if !slices.Equal(inOrder(valuesOf(&m, key)), want) {
			t.Fatalf("after %d values: %q", i+1, valuesOf(&m, key))
		}
	}
	for i := 39; i > 0; i-- {
		m.Remove(key, val(i))
		want = want[:i]
		isValueOverflow := m.t.find(key).isValueOverflow()
		// the value overflow stays until the values take half the room of a page or less; the page, once back, until it is empty
		if isValueOverflow && skpage.BackFits(3, i*(1+len(val(0)))) {
			t.Fatalf("with %d values (content %d) the key still has a value overflow", i, content(i))
		}
		if !isValueOverflow && content(i) > 512 {
			t.Fatalf("with %d values (content %d) the key has a page", i, content(i))
		}
		if !slices.Equal(inOrder(valuesOf(&m, key)), want) {
			t.Fatalf("after removing down to %d values: %q", i, valuesOf(&m, key))
		}
	}
	// the border: between 256 and 512 bytes the key keeps whichever object it has
	m.Add(key, val(1))
	if m.t.find(key).isValueOverflow() {
		t.Fatal("a key with two values has a value overflow")
	}
	m.Remove(key, val(0))
	m.Remove(key, val(1))
	if m.Has(key) {
		t.Fatal("the key must be gone with its last value")
	}
}

// TestStringValueTooLong covers a string of 255 bytes or more as a value:
// such a value cannot go into a page, so the key becomes a value overflow with all its
// values, and a page again once the long value is gone and the rest is small.
func TestStringValueTooLong(t *testing.T) {
	var m Map[string]
	key := []byte("key")
	long := strings.Repeat("L", 300)
	m.Add(key, "a")
	m.Add(key, "b")
	m.Add(key, long)
	if !m.t.find(key).isValueOverflow() {
		t.Fatal("a value of 300 bytes must make a value overflow")
	}
	if got := inOrder(valuesOf(&m, key)); !slices.Equal(got, inOrder([]string{"a", "b", long})) {
		t.Fatalf("got %d values", len(got))
	}
	m.Remove(key, "a") // the set still holds the long value: it stays a value overflow
	if !m.t.find(key).isValueOverflow() {
		t.Fatal("with the long value in it the key must stay a value overflow")
	}
	m.Remove(key, long)
	if m.t.find(key).isValueOverflow() || !slices.Equal(valuesOf(&m, key), []string{"b"}) {
		t.Fatalf("the key must be a page with the value b, has %q", valuesOf(&m, key))
	}
}

// TestStringKeyTooLong covers keys whose end does not fit a page: a lone key of
// 505 bytes has a remainder of 505, a key of more than 64 KiB cannot even
// have its length in the page's header. Both get a value overflow that holds the key
// as a string, and keep their values, also when values go.
func TestStringKeyTooLong(t *testing.T) {
	for _, n := range []int{505, 600, maxKeyLen + 10} {
		t.Run(fmt.Sprint(n, " bytes"), func(t *testing.T) {
			var m Map[string]
			key := bytes.Repeat([]byte("k"), n)
			for _, v := range []string{"a", "b", "c"} {
				m.Add(key, v)
			}
			if !m.t.find(key).isValueOverflow() {
				t.Fatal("expected a value overflow")
			}
			m.Remove(key, "b")
			if got := inOrder(valuesOf(&m, key)); !slices.Equal(got, []string{"a", "c"}) {
				t.Fatalf("got %q", got)
			}
		})
	}
	t.Run("the longest remainder a page holds", func(t *testing.T) {
		var m Map[string]
		key := bytes.Repeat([]byte("k"), skpage.MaxRemainder-1) // and a value of one byte
		m.Add(key, "a")
		if m.t.find(key).isValueOverflow() {
			t.Fatal("a remainder of 504 bytes fits a page")
		}
	})
}

// TestStringValueOverflowStays covers a value overflow that cannot become a page: one with
// too many values for a page (every value takes two bytes at least; 255 values
// of one byte are the most there can be), however few of them go.
func TestStringValueOverflowStays(t *testing.T) {
	var m Map[string]
	key := []byte("key")
	one := func(i int) string { return string([]byte{byte(i)}) }
	for i := range 255 {
		m.Add(key, one(i))
	}
	m.Remove(key, one(0))
	if !m.t.find(key).isValueOverflow() || len(valuesOf(&m, key)) != 254 {
		t.Fatal("254 values of one byte are a value overflow")
	}
}

// TestStringRekey covers what happens to the object of a key when the node above
// it goes away and its path gets shorter (see rekeyFunc): a page takes the bytes
// in front of its remainder, in place if they fit its class, in a page of a larger
// class, or, when the remainder would be longer than a page holds, as a
// value overflow; a value overflow is left to the value overflow code. The values and the key
// stay.
func TestStringRekey(t *testing.T) {
	key := bytes.Repeat([]byte("abcdefghij"), 70) // 700 bytes
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
		{"a remainder of 270 bytes (nine bits of length)", 290, 280, 20, []string{"v", "w"}, "page"},
		{"a remainder beyond what a page holds", 600, 590, 20, []string{"v", "w"}, "set"},
		{"a value overflow", 40, 37, 30, []string{strings.Repeat("L", 300)}, "set"},
		{"a value overflow that needs a larger key area", 100, 99, 20, []string{strings.Repeat("L", 300)}, "set"},
		{"a value overflow in the largest key area", 200, 80, 70, []string{strings.Repeat("L", 300)}, "set"},
		{"a value overflow in the key area of 384 bytes", 400, 100, 90, []string{strings.Repeat("L", 300)}, "set"},
		{"a value overflow in the key area of 512 bytes", 600, 200, 190, []string{strings.Repeat("L", 300)}, "set"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			k := key[:tc.keyLen]
			l := newSK(k, tc.base)
			slot := singleKeyHdr(l)
			for _, v := range tc.values {
				if l.isValueOverflow() {
					overflowAdd(l, v)
				} else {
					addSK(&slot, l, k, v)
					l = asSingleKey(slot)
				}
			}
			nl := rekeySK(l, k[:tc.base-1], int(k[tc.base-1]), tc.to)
			var got []string
			if nl.isValueOverflow() {
				overflowEach(nl, func(v string) bool { got = append(got, v); return true })
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
			case tc.want == "in place" && nl != l, tc.want == "page" && (nl == l || nl.isValueOverflow()), tc.want == "set" && !nl.isValueOverflow():
				t.Errorf("rekey gave type %d, same leaf: %v, want %s", nl.objType, nl == l, tc.want)
			}
		})
	}
}

// TestStringValueOverflowLayout covers the value overflow of a string map as an object: it
// lies on the grid of the pages (32 to 512 bytes, 32 for a key held as a
// string), the pointer to its value set is the last word of the object whatever
// the key area, the key area is the one that fills the object, and the object
// statistic knows its size, which is the size the runtime allocates.
func TestStringValueOverflowLayout(t *testing.T) {
	var strs Map[string]
	strs.flat = 3
	for _, tc := range []struct {
		n, size int
	}{{5, 32}, {18, 32}, {19, 64}, {50, 64}, {51, 128}, {114, 128}, {115, 256}, {242, 256}, {243, 384}, {370, 384}, {371, 512}, {498, 512}, {499, 32}, {600, 32}} {
		key := bytes.Repeat([]byte("k"), tc.n)
		l := newValueOverflow(key, 0, set3.EmptyWithCapacity[string](4))
		if off := int(uintptr(unsafe.Pointer(overflowSetOf[string](l))) - uintptr(unsafe.Pointer(l))); tc.n <= maxInlineOverflow && off != tc.size-8 {
			t.Errorf("key of %d bytes: set at %d, want the last word of %d", tc.n, off, tc.size)
		}
		overflowAdd(l, "v")
		if got := strs.leafValues(l); got != 1 {
			t.Errorf("key of %d bytes: %d values", tc.n, got)
		}
		o := strs.leafKind(l)
		if o.Size != tc.size || int(valueOverflowSize(l.rem())) != tc.size || o.Label != "value overflow" || !o.Pointers {
			t.Errorf("key of %d bytes: object %+v, want a value overflow of %d bytes", tc.n, o, tc.size)
		}
		if string(l.stored()) != string(key) && tc.n <= maxInlineOverflow || l.keyLen() != tc.n {
			t.Errorf("key of %d bytes: stored %q, length %d", tc.n, l.stored(), l.keyLen())
		}
	}
}

// TestStringBackWithLongRemainder covers a key whose remainder is long, which
// leaves little room for values: the key moves into a value overflow when
// its few values no longer fit, and comes back into a page when they take half
// of the room that is left, which a rule of fixed content, counting the
// remainder, never allowed (a remainder of 300 bytes took all of the 256).
func TestStringBackWithLongRemainder(t *testing.T) {
	var m Map[string]
	key := bytes.Repeat([]byte("k"), 400)
	val := func(i int) string { return fmt.Sprintf("value-%02d-%s", i, strings.Repeat("x", 20)) } // 30 bytes, 31 with its length
	for i := range 4 {
		m.Add(key, val(i))
	}
	if !m.t.find(key).isValueOverflow() {
		t.Fatal("four values of 31 bytes behind a remainder of 400 bytes (106 left) fit a page")
	}
	m.Remove(key, val(3))
	m.Remove(key, val(2))
	if !m.t.find(key).isValueOverflow() {
		t.Fatal("two values (62 bytes) take more than half of the 106 bytes: the key goes back too early")
	}
	m.Remove(key, val(1))
	if m.t.find(key).isValueOverflow() {
		t.Fatal("one value (31 bytes) takes less than half of the room: the key stays in the value overflow")
	}
	if got := valuesOf(&m, key); !slices.Equal(got, []string{val(0)}) {
		t.Fatalf("got %q", got)
	}
}

// TestValueOverflowOfOtherValues covers that the value overflow, the object
// that holds the values of a key outside its page, is the same object for every type of value: its size
// and the place of the pointer to the set depend on the key only, and the set
// holds what it is given, whether the values are strings, words or pointers
// (step 3.5: the maps of uint64 and *T use it). Every key area of the grid and a
// key held as a string are made for each type, filled and read back.
func TestValueOverflowOfOtherValues(t *testing.T) {
	type rec struct{ id int }
	recs := make([]*rec, 100)
	for i := range recs {
		recs[i] = &rec{i}
	}
	check := func(t *testing.T, name string, n int, make func(key []byte) (l *singleKeyHead, add func(i int), read func() int)) {
		for _, klen := range []int{0, 18, 19, 50, 114, 242, 370, 498, 499} {
			key := bytes.Repeat([]byte("k"), klen)
			l, add, read := make(key)
			for i := range n {
				add(i)
				add(i) // a set: the second time changes nothing
			}
			if got := read(); got != n {
				t.Errorf("%s, key of %d bytes: %d values, want %d", name, klen, got, n)
			}
			if !l.isValueOverflow() || len(l.stored()) != klen || l.keyLen() != klen {
				t.Errorf("%s, key of %d bytes: the key is not what was given", name, klen)
			}
			if want := int(valueOverflowSize(l.rem())); want != 32 && want != 64 && want != 128 && want != 256 && want != 384 && want != 512 {
				t.Errorf("%s, key of %d bytes: an object of %d bytes is off the grid", name, klen, want)
			}
		}
	}
	check(t, "uint64", 100, func(key []byte) (*singleKeyHead, func(int), func() int) {
		l := newValueOverflow(key, 0, set3.EmptyWithCapacity[uint64](4))
		return l, func(i int) { overflowAdd(l, uint64(i)) }, func() int {
			n := 0
			overflowEach(l, func(uint64) bool { n++; return true })
			return n
		}
	})
	check(t, "*rec", 100, func(key []byte) (*singleKeyHead, func(int), func() int) {
		l := newValueOverflow(key, 0, set3.EmptyWithCapacity[*rec](4))
		return l, func(i int) { overflowAdd(l, recs[i]) }, func() int {
			n := 0
			overflowEach(l, func(r *rec) bool { n += min(r.id, 0) + 1; return true })
			return n
		}
	})
}
