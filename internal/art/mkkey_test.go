package art

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
	"testing"
)

// pageCount returns how many multi-key pages the map holds and how many keys are in them.
func pageCount[T comparable](m *Map[T]) (pages, keys int) {
	m.Objects(func(o Object) {
		if o.Label == "multi-key page" {
			pages++
			keys += o.Keys
		}
	})
	return pages, keys
}

// mkKeys returns n keys with a common start, in no order that the tree could rely on.
func mkKeys(n int) [][]byte {
	keys := make([][]byte, n)
	for i := range keys {
		keys[i] = fmt.Appendf(nil, "street-%03d", (i*37)%n)
	}
	return keys
}

// TestMultiKeyPages makes sure that the keys of a map with one value each share pages,
// and that adding a second value to a key, and removing keys and values, keep every
// answer right.
//
// A user of Ordered who maps names to one value each (ids, locations) gets the memory
// of the multi-key pages (docs/redesign/step4-mksv-design.md): many keys in one object
// instead of one object per key. A key that gets a second value must not disturb its
// neighbours, and a page that loses keys must give way to what is left.
//
// Expected, for strings and for uint64 values: after adding 300 keys with one value
// each most of them are in multi-key pages (and every one is found); a second value
// stays in the key's page; removing the keys one by one leaves the others
// intact down to the last key, which is a single-key page.
func TestMultiKeyPages(t *testing.T) {
	t.Run("strings", func(t *testing.T) { runMultiKeyPages(t, func(i int) string { return fmt.Sprint("v", i) }) })
	t.Run("uint64", func(t *testing.T) { runMultiKeyPages(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runMultiKeyPages(t, recPool()) })
}

func runMultiKeyPages[T comparable](t *testing.T, val func(i int) T) {
	var m Map[T]
	keys := mkKeys(300)
	for i, k := range keys {
		m.Add(k, val(i))
		m.Add(k, val(i)) // the same value again changes nothing
	}
	checkInvariants(t, &m.t)
	if pages, inPages := pageCount(&m); pages == 0 || inPages < 250 {
		t.Fatalf("%d multi-key pages hold %d of 300 keys", pages, inPages)
	}
	for i, k := range keys {
		if got := valuesOf(&m, k); !slices.Equal(got, []T{val(i)}) {
			t.Fatalf("key %q holds %v, want %v", k, got, val(i))
		}
	}
	if m.Has([]byte("street-")) || m.Has([]byte("street-0000")) || m.Has([]byte("x")) {
		t.Fatal("found a key that was never added")
	}
	_, before := pageCount(&m)
	m.Add(keys[10], val(1000+10)) // a second value: the key stays in its page
	if _, after := pageCount(&m); after != before {
		t.Fatalf("a second value moved a key out of its page: %d keys in pages before, %d after", before, after)
	}
	if got := valuesOf(&m, keys[10]); len(got) != 2 {
		t.Fatalf("key with two values holds %v", got)
	}
	checkInvariants(t, &m.t)
	for i, k := range keys {
		m.Remove(k, val(i+1000)) // a value the key does not have, or not one at all
		m.Remove(k, val(i))
		if m.Has(k) && i != 10 {
			t.Fatalf("key %q still there after removing its value", k)
		}
		checkInvariants(t, &m.t)
		for _, other := range keys[i+1:] {
			if !m.Has(other) {
				t.Fatalf("removing %q lost %q", k, other)
			}
		}
	}
	m.RemoveKey(keys[10])
	if m.Len() != 0 || m.t.root != nil {
		t.Fatalf("map not empty: %d keys", m.Len())
	}
}

// TestMultiKeyPageAboveAndUp puts a byte node above a full multi-key page with a key that
// leaves its common prefix, in every way a key can, and takes the node away again.
//
// The tree adds keys all the time that share some of the bytes of a page and not all: a page
// with room takes them in and shortens its common prefix; a full page gives the front of its
// common prefix to a node above it, and takes it back when the node goes away.
//
// Expected: after each added key, all earlier keys and the new one are found, the new key
// being a key that differs inside the common prefix, one that ends inside it, the empty
// key; and after the other keys are removed, the page is back where it started and holds
// its keys.
func TestMultiKeyPageAboveAndUp(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		runAboveAndUp(t, 3, func(i int) string { return fmt.Sprint(i, strings.Repeat("v", 150)) })
	})
	t.Run("uint64", func(t *testing.T) { runAboveAndUp(t, 49, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runAboveAndUp(t, 49, recPool()) })
}

func runAboveAndUp[T comparable](t *testing.T, n int, val func(i int) T) {
	var m Map[T]
	var page [][]byte
	for i := range n {
		page = append(page, fmt.Appendf(nil, "abcdef%c", 'A'+i))
		m.Add(page[i], val(i))
	}
	if pages, keys := pageCount(&m); pages != 1 || keys != n {
		t.Fatalf("%d pages with %d keys, want one with %d", pages, keys, n)
	}
	if m.t.findLeaf(page[0]) != nil || !m.Has(page[0]) {
		t.Fatal("the key of a multi-key page has a single-key page")
	}
	others := [][]byte{[]byte("abX"), []byte("ab"), []byte("a"), []byte("z"), []byte("")}
	for i, k := range others {
		m.Add(k, val(100+i))
		m.Add(k, val(200+i)) // two values: a single-key page, which stays when the others go
		checkInvariants(t, &m.t)
		for j, p := range page {
			if got := valuesOf(&m, p); !slices.Equal(got, []T{val(j)}) {
				t.Fatalf("after adding %q the page key %q holds %v", k, p, got)
			}
		}
	}
	for _, k := range others {
		m.RemoveKey(k)
		checkInvariants(t, &m.t)
		for j, p := range page {
			if got := valuesOf(&m, p); !slices.Equal(got, []T{val(j)}) {
				t.Fatalf("after removing %q the page key %q holds %v", k, p, got)
			}
		}
	}
	if pages, keys := pageCount(&m); pages != 1 || keys != n {
		t.Fatalf("%d pages with %d keys after removing the others, want one with %d", pages, keys, n)
	}
}

// TestMultiKeyPageAbsorbs shows that a key which leaves the common prefix of a page joins
// the page when they fit together, and does not make a node of its own.
//
// Pages are the memory saving; a key that differs from its neighbours in an early byte must
// not cost a node and a single-key page when a page has room for it.
//
// Expected: three keys and a fourth that differs in the first byte are one page of four keys.
func TestMultiKeyPageAbsorbs(t *testing.T) {
	var m Map[uint64]
	for i, k := range []string{"abc1", "abc2", "abc3", "xyz"} {
		m.Add([]byte(k), uint64(i))
	}
	if pages, keys := pageCount(&m); pages != 1 || keys != 4 {
		t.Fatalf("%d pages with %d keys, want one with 4", pages, keys)
	}
	checkInvariants(t, &m.t)
}

// TestMultiKeyPageCannotMoveUp shows a page that stays below a node with no other child
// when the bytes of the node no longer fit in front of its common prefix.
//
// A node that is left with one child goes away and the child takes its bytes; a multi-key
// page that has filled up since the node was made cannot, and the tree then keeps the node.
//
// Expected: after the other child is removed the page's keys are all found, the tree
// passes its invariants, and the node is still there.
func TestMultiKeyPageCannotMoveUp(t *testing.T) {
	var m Map[string]
	p := "0123456789"
	val := func(n int) string { return strings.Repeat("v", n) }
	page := func() *header { h, _ := m.t.find([]byte(p + "q1")); return h }
	for _, k := range []string{"1", "2"} {
		m.Add([]byte(p+"q"+k), val(240))
	}
	m.Add([]byte(p+"q3"), val(9)) // a page of 512 bytes: the next key does not fit it, and a node goes above
	m.Add([]byte(p+"r"), "x")
	m.Add([]byte(p+"r"), "y") // two values: a single-key page, the node's other child
	if isPage(m.t.root.objType) {
		t.Fatal("no node above the page")
	}
	used := asMKStr(page()).Used()
	m.Add([]byte(p+"q5"), val(505-used-3)) // the page is now at 505 bytes: eleven more do not fit
	m.RemoveKey([]byte(p + "r"))
	checkInvariants(t, &m.t)
	if isPage(m.t.root.objType) {
		t.Fatalf("the page moved up although it is full")
	}
	for _, k := range []string{"1", "2", "3", "5"} {
		if !m.Has([]byte(p + "q" + k)) {
			t.Fatalf("lost key %q", k)
		}
	}
}

// TestMultiKeyPageShrunkBelowItsNode makes sure that a page that could not move up and shrinks to
// one key leaves no node with a single leaf behind, and that the last key can be removed.
//
// A user who deletes most of an index must not crash on the last keys of a subtree (it once did,
// at the end of removing half the keys of a directory listing): the page below a node that it could
// not join becomes a single-key page when one key is left, and the node, which then has one
// child that can move up, goes away.
//
// Expected: the page is brought down to one key without the merge that a removal would try (so
// the node is left with one single-key page); the merge of the removal then takes the node away and
// the tree passes its invariants; the same state, left as it is, still lets the last key be removed.
func TestMultiKeyPageShrunkBelowItsNode(t *testing.T) {
	p := "0123456789"
	val := func(n int) string { return strings.Repeat("v", n) }
	stuck := func() *Map[string] {
		var m Map[string]
		for _, k := range []string{"1", "2"} {
			m.Add([]byte(p+"q"+k), val(240))
		}
		m.Add([]byte(p+"q3"), val(9))
		m.Add([]byte(p+"r"), "x")
		m.Add([]byte(p+"r"), "y")
		h, _ := m.t.find([]byte(p + "q1"))
		m.Add([]byte(p+"q5"), val(505-asMKStr(h).Used()-3))
		m.RemoveKey([]byte(p + "r")) // the page stays below the node: it is full
		checkInvariants(t, &m.t)
		if page, _ := m.t.find([]byte(p + "q1")); !isMultiKey(page.objType) || isPage(m.t.root.objType) {
			t.Fatal("the page is not below a node")
		}
		for _, k := range []string{"3", "2", "1"} { // the removals of Remove without its merge
			key := []byte(p + "q" + k)
			page, pathLen := m.t.find(key)
			m.pageRemove(page, pathLen, key, "", true)
		}
		if h, _ := m.t.find([]byte(p + "q5")); !isSingleKey(h.objType) || m.t.root.count != 1 {
			t.Fatal("the page did not become a single-key page below the node")
		}
		return &m
	}
	t.Run("the merge of a removal takes the node away", func(t *testing.T) {
		m := stuck()
		m.mergeUp(&m.t.root, []byte(p+"q1"), 0)
		if !isPage(m.t.root.objType) {
			t.Fatal("the node with one single-key page is still there")
		}
		checkInvariants(t, &m.t)
		if got := valuesOf(m, []byte(p+"q5")); len(got) != 1 || m.Len() != 1 {
			t.Fatalf("the key is lost: %v, Len %d", got, m.Len())
		}
	})
	t.Run("the last key can be removed from the state before the merge", func(t *testing.T) {
		m := stuck()
		m.t.remove([]byte(p+"q5"), m.rekey)
		if m.t.root != nil || m.t.size != 0 {
			t.Fatalf("tree not empty: %d keys", m.t.size)
		}
	})
}

// TestMultiKeyPageRefusals shows what keeps an entry out of a multi-key page: a value or a
// remainder that is too long, or entries that together do not fit.
//
// The tree must never lose an entry because a page refuses it; it makes the entry a single-key
// page or a value overflow, and keeps the others in pages.
//
// Expected: all keys are found with their values in each case below.
func TestMultiKeyPageRefusals(t *testing.T) {
	long := strings.Repeat("v", 300)
	t.Run("a value of 300 bytes joins a page of short ones", func(t *testing.T) {
		var m Map[string]
		for i := range 10 {
			m.Add(fmt.Appendf(nil, "k%d", i), "short")
		}
		m.Add([]byte("k5x"), long) // refused by the page: it bursts
		m.Add([]byte("q1"), "x")
		m.Add([]byte("q2"), long) // meets the single-key page of q1, and the two do not make a page
		if got := valuesOf(&m, []byte("q1")); !slices.Equal(got, []string{"x"}) || m.t.findLeaf([]byte("zzz")) != nil {
			t.Fatalf("q1 holds %v", got)
		}
		checkInvariants(t, &m.t)
		for _, k := range []string{"k0", "k9"} {
			if got := valuesOf(&m, []byte(k)); !slices.Equal(got, []string{"short"}) {
				t.Fatalf("%q holds %v", k, got)
			}
		}
		if got := valuesOf(&m, []byte("k5x")); !slices.Equal(got, []string{long}) {
			t.Fatal("the long value is lost")
		}
	})
	t.Run("a remainder of 300 bytes joins a page of short ones", func(t *testing.T) {
		var m Map[uint64]
		for i := range 10 {
			m.Add(fmt.Appendf(nil, "k%d", i), uint64(i))
		}
		longKey := append([]byte("k5"), bytes.Repeat([]byte("y"), 298)...)
		m.Add(longKey, 99)
		m.Add([]byte("a"), 1)
		other := append([]byte("b"), bytes.Repeat([]byte("y"), 299)...) // two long keys do not make a page either
		m.Add(other, 2)
		m.Add(append(slices.Clone(other), 'z'), 3)
		checkInvariants(t, &m.t)
		if got := valuesOf(&m, longKey); !slices.Equal(got, []uint64{99}) {
			t.Fatalf("long key holds %v", got)
		}
		for i := range 10 {
			if got := valuesOf(&m, fmt.Appendf(nil, "k%d", i)); !slices.Equal(got, []uint64{uint64(i)}) {
				t.Fatalf("k%d holds %v", i, got)
			}
		}
	})
	t.Run("a merge whose remainder is too long for a page", func(t *testing.T) {
		var m Map[string]
		longKey := "a" + strings.Repeat("y", 300)
		m.Add([]byte(longKey), "1") // a single-key page: no page holds a remainder of 300 bytes
		m.Add([]byte("b"), "2")
		m.Add([]byte("c"), "3")
		m.RemoveKey([]byte("c")) // the node's children are single-value pages that fit one page, but not this key
		checkInvariants(t, &m.t)
		if got := valuesOf(&m, []byte(longKey)); !slices.Equal(got, []string{"1"}) || m.Len() != 2 {
			t.Fatalf("long key holds %v, Len %d", got, m.Len())
		}
	})
	t.Run("two values of 300 bytes make a value overflow", func(t *testing.T) {
		var m Map[string]
		m.Add([]byte("a"), "1")
		m.Add([]byte("b"), "2") // a page with "a"
		m.Add([]byte("b"), long)
		m.Add([]byte("b"), long+"!") // a second value: the page's entries are built again
		checkInvariants(t, &m.t)
		if got := valuesOf(&m, []byte("b")); len(got) != 3 {
			t.Fatalf("b holds %d values", len(got))
		}
		if !m.t.findLeaf([]byte("b")).isValueOverflow() {
			t.Fatal("three long values are in a page")
		}
	})
	t.Run("a page of ten keys bursts into pages", func(t *testing.T) {
		var m Map[uint64]
		for i := range 400 {
			m.Add(fmt.Appendf(nil, "k%04d", i), uint64(i))
		}
		checkInvariants(t, &m.t)
		for i := range 400 {
			if got := valuesOf(&m, fmt.Appendf(nil, "k%04d", i)); !slices.Equal(got, []uint64{uint64(i)}) {
				t.Fatalf("k%04d holds %v", i, got)
			}
		}
	})
}

// TestMultiKeyPageScans makes sure that range scans and prefix scans visit the keys and
// values of multi-key pages in key order, with the bounds a caller sets, and stop when the caller
// says so.
//
// Expected: for several bounds, including ones that fall inside a page, the keys and values
// come in order, a callback that returns false ends the scan, and every key lies within.
func TestMultiKeyPageScans(t *testing.T) {
	t.Run("strings", func(t *testing.T) { runScans(t, func(i int) string { return fmt.Sprint("v", i) }) })
	t.Run("uint64", func(t *testing.T) { runScans(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runScans(t, recPool()) })
}

func runScans[T comparable](t *testing.T, val func(i int) T) {
	var m Map[T]
	keys := mkKeys(200)
	for i, k := range keys {
		m.Add(k, val(i))
	}
	sorted := slices.Clone(keys)
	slices.SortFunc(sorted, bytes.Compare)
	for _, b := range []*Bounds{
		{},
		{From: []byte("street-050"), HasFrom: true, FromIncl: true},
		{From: []byte("street-050"), HasFrom: true},
		{To: []byte("street-150"), HasTo: true, ToIncl: true},
		{To: []byte("street-150"), HasTo: true},
		{From: []byte("street-0505"), To: []byte("street-1005"), HasFrom: true, HasTo: true},
		{From: []byte("street-100"), To: []byte("street-100"), HasFrom: true, HasTo: true, FromIncl: true, ToIncl: true},
		{From: []byte("a"), To: []byte("b"), HasFrom: true, HasTo: true},
	} {
		var want [][]byte
		for _, k := range sorted {
			if b.Contains(k) {
				want = append(want, k)
			}
		}
		var got [][]byte
		m.Range(b, func(k []byte) bool { got = append(got, bytes.Clone(k)); return true })
		if !slices.EqualFunc(got, want, bytes.Equal) {
			t.Fatalf("bounds %+v: %d keys, want %d", b, len(got), len(want))
		}
		n := 0
		m.RangeValues(b, func(T) bool { n++; return true })
		if n != len(want) {
			t.Fatalf("bounds %+v: %d values, want %d", b, n, len(want))
		}
		if len(want) > 3 {
			stop := 0
			m.Range(b, func([]byte) bool { stop++; return stop < 3 })
			m.RangeValues(b, func(T) bool { stop++; return stop < 6 })
			if stop != 6 {
				t.Fatalf("the scan did not stop when told to: %d", stop)
			}
		}
	}
}

// TestMultiKeyPageOfWordsCannotMoveUp is TestMultiKeyPageCannotMoveUp for a page of
// uint64 values, which holds 49 keys behind a common prefix of 11 bytes and 50 without it:
// the bytes of the node above it would take its content beyond 512.
//
// Expected: after the node's other child is removed, every key of the page is
// found, the invariants hold and the node is still there.
func TestMultiKeyPageOfWordsCannotMoveUp(t *testing.T) {
	var m Map[uint64]
	p := "0123456789"
	key := func(i int) []byte { return []byte(fmt.Sprint(p, "q", string(rune('A'+i)))) }
	for i := range 49 {
		m.Add(key(i), uint64(i))
	}
	m.Add([]byte(p+"r"), 1) // does not fit the page: a node goes above it
	m.Add([]byte(p+"r"), 2) // two values: a single-key page
	m.Add(key(49), 49)      // 50 keys fit a page without the 11 bytes
	m.RemoveKey([]byte(p + "r"))
	checkInvariants(t, &m.t)
	if isPage(m.t.root.objType) {
		t.Fatalf("the page moved up although it is full")
	}
	for i := range 50 {
		if got := valuesOf(&m, key(i)); !slices.Equal(got, []uint64{uint64(i)}) {
			t.Fatalf("key %d holds %v", i, got)
		}
	}
}

// TestMultiKeyPageMergeWaitsForNearlyEmptyPage makes sure that removing values from a
// map with multi-key pages does not try to merge pages while they are still well filled.
//
// A user who deletes entries from a large index of single-value entries should not pay for
// a merge attempt at every deletion: with pages that hold some ten entries, nearly every
// attempt fails (docs/redesign/step4-probe.md). The tree comes together again once a page
// is nearly empty.
//
// Expected: two sibling pages of six entries each, whose entries together would fit one
// page after four removals, stay two pages while the page the removals come from has
// more than mergeBelow entries left, and become one page with the removal that leaves it
// mergeBelow entries.
func TestMultiKeyPageMergeWaitsForNearlyEmptyPage(t *testing.T) {
	var m Map[uint64]
	key := func(side byte, i int) []byte { return fmt.Appendf(nil, "%c%02d%s", side, i, strings.Repeat("x", 40)) }
	for i := range 6 {
		m.Add(key('a', i), uint64(i))
		m.Add(key('b', i), uint64(100+i))
	}
	if pages, keys := pageCount(&m); pages != 2 || keys != 12 {
		t.Fatalf("setup: %d pages hold %d keys, want 2 pages with 12", pages, keys)
	}
	for i := range 6 - mergeBelow - 1 { // the page keeps mergeBelow+1 entries
		m.Remove(key('a', i), uint64(i))
	}
	if pages, _ := pageCount(&m); pages != 2 {
		t.Fatalf("%d pages with %d entries left in one of them, want 2 (no merge yet)", pages, mergeBelow+1)
	}
	m.Remove(key('a', 6-mergeBelow-1), uint64(6-mergeBelow-1))
	if pages, keys := pageCount(&m); pages != 1 || keys != 6+mergeBelow {
		t.Fatalf("after the removal that leaves %d entries: %d pages with %d keys, want 1 page with %d", mergeBelow, pages, keys, 6+mergeBelow)
	}
	checkInvariants(t, &m.t)
}

// TestMultiKeyPageLongPrefix makes sure that keys which share 300 bytes below the last byte
// node still share one page: the common prefix of a page has nine bits (the lowest bit of the
// type byte is the ninth), and the page keeps it through the byte nodes that come and go above it.
//
// A user whose keys are long paths or URLs with a long common start, a few of them below one
// branch, gets one object for them and not a node and a single-key page each.
//
// Expected: three keys with a 300-byte common start are one page with that prefix; a key that
// leaves the prefix after 100 bytes puts a node above the page, which then holds 199 bytes of
// prefix; when that key goes the page has 300 again; every key is found all the time.
func TestMultiKeyPageLongPrefix(t *testing.T) {
	t.Run("strings", func(t *testing.T) { runLongPrefix(t, func(i int) string { return fmt.Sprint("v", i) }) })
	t.Run("uint64", func(t *testing.T) { runLongPrefix(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runLongPrefix(t, recPool()) })
}

func runLongPrefix[T comparable](t *testing.T, val func(i int) T) {
	var m Map[T]
	long := strings.Repeat("p", 300)
	var keys [][]byte
	for i := range 3 {
		keys = append(keys, fmt.Appendf(nil, "%s%c", long, 'A'+i))
		m.Add(keys[i], val(i))
	}
	prefix := func() (n int) {
		m.Objects(func(o Object) {
			if o.Label == "multi-key page" {
				n = o.Remainder
			}
		})
		return n
	}
	if pages, inPages := pageCount(&m); pages != 1 || inPages != 3 || prefix() < 255 {
		t.Fatalf("%d pages with %d keys and a prefix of %d bytes, want one with 3 and 300 or so", pages, inPages, prefix())
	}
	check := func() {
		t.Helper()
		checkInvariants(t, &m.t)
		for i, k := range keys {
			if got := valuesOf(&m, k); !slices.Equal(got, []T{val(i)}) {
				t.Fatalf("key %d holds %v", i, got)
			}
		}
	}
	check()
	other := []byte(long[:100] + "X")
	m.Add(other, val(9))
	m.Add(other, val(10)) // two values: a single-key page, which stays when the others go
	check()
	if p := prefix(); p >= 255 {
		t.Fatalf("with a node above, the page's prefix is %d bytes", p)
	}
	m.RemoveKey(other)
	check()
	if pages, inPages := pageCount(&m); pages != 1 || inPages != 3 || prefix() < 255 {
		t.Fatalf("after the other key is gone: %d pages with %d keys and a prefix of %d bytes", pages, inPages, prefix())
	}
}

// TestMultiKeyPagesChurnKeepsTheStructure makes sure that adding and removing keys and values
// of a map with long shared paths, in any order, never leaves the tree in a state the next
// removal cannot handle.
//
// A user whose keys are file paths (long common starts, a few keys below each directory) adds and
// removes keys all the time; a page that shrinks to one key below a node that it could not join must
// not leave a node with a single leaf behind (the removal of that leaf once crashed).
//
// Expected: after every operation of a long random sequence, for strings and for uint64 values, the
// tree passes its invariants and the map agrees with a model of the keys and values it should hold.
func TestMultiKeyPagesChurnKeepsTheStructure(t *testing.T) {
	t.Run("strings", func(t *testing.T) { runStructureChurn(t, func(i int) string { return fmt.Sprint("v", i) }) })
	t.Run("uint64", func(t *testing.T) { runStructureChurn(t, func(i int) uint64 { return uint64(i) }) })
	t.Run("pointers", func(t *testing.T) { runStructureChurn(t, recPool()) })
}

func runStructureChurn[T comparable](t *testing.T, val func(i int) T) {
	for seed := range uint64(12) {
		r := rand.New(rand.NewPCG(seed, 9))
		var m Map[T]
		model := map[string][]int{}
		dirs := []string{"/usr/share/doc/libtss2-dev/html/", "/usr/share/doc/libtss2-mu0/", "/usr/share/doc/libtss2-esys0/", "/usr/lib/x86_64-linux-gnu/", "/etc/"}
		key := func() []byte {
			d := dirs[r.IntN(len(dirs))]
			return fmt.Appendf(nil, "%s%s%d", d, strings.Repeat("n", r.IntN(30)), r.IntN(40))
		}
		for step := range 1500 {
			k := key()
			switch op := r.IntN(10); {
			case op < 5:
				v := r.IntN(6)
				m.Add(k, val(v))
				if !slices.Contains(model[string(k)], v) {
					model[string(k)] = append(model[string(k)], v)
				}
			case op < 7:
				m.RemoveKey(k)
				delete(model, string(k))
			default:
				v := r.IntN(6)
				m.Remove(k, val(v))
				if vs := model[string(k)]; slices.Contains(vs, v) {
					if vs = slices.DeleteFunc(vs, func(x int) bool { return x == v }); len(vs) == 0 {
						delete(model, string(k))
					} else {
						model[string(k)] = vs
					}
				}
			}
			checkInvariants(t, &m.t)
			if m.Len() != len(model) {
				t.Fatalf("seed %d step %d: Len %d, model %d", seed, step, m.Len(), len(model))
			}
		}
		for k, vs := range model {
			if got := valuesOf(&m, []byte(k)); len(got) != len(vs) {
				t.Fatalf("seed %d: key %q holds %d values, want %d", seed, k, len(got), len(vs))
			}
		}
	}
}
