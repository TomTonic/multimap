package art

import (
	"bytes"
	"testing"
)

// TestRemoveLongKeysOneByOne checks that removing two keys of more than 255
// bytes that share a long prefix, one after the other, leaves an empty map.
//
// A user who stores long keys such as URLs or paths of 300 bytes with a common
// beginning removes them again. In the tree the two keys sit below a range node
// with the common bytes as its path; after the first removal the node keeps its
// last page, because the page could not hold its keys with the path in front of
// them (a page holds at most 255 bytes of a key). Removing the last key then
// empties the node, which must disappear: it once crashed the map.
func TestRemoveLongKeysOneByOne(t *testing.T) {
	var m Map[uint64]
	prefix := bytes.Repeat([]byte{'p'}, 100)
	k1 := append(append(bytes.Clone(prefix), '1'), bytes.Repeat([]byte{'x'}, 199)...)
	k2 := append(append(bytes.Clone(prefix), '2'), bytes.Repeat([]byte{'x'}, 199)...)
	m.Add(k1, 1)
	m.Add(k2, 2)
	m.RemoveKey(k1)
	if m.Len() != 1 || !m.Has(k2) {
		t.Fatalf("after the first removal: Len %d, Has(k2) %v", m.Len(), m.Has(k2))
	}
	m.RemoveKey(k2)
	if m.Len() != 0 || m.Has(k1) || m.Has(k2) {
		t.Fatalf("after the second removal: Len %d", m.Len())
	}
	m.Add(k1, 3) // and the map still works
	if !m.Has(k1) {
		t.Fatal("the key is gone after adding it again")
	}
}

// TestTreeRemoveAbsentKeys shows that the tree's removal of a key that is not in it, however
// near the tree it gets, changes nothing and says so.
//
// The map looks up a key before it removes it, but the tree does not rely on that: a key that
// is absent, that ends inside a node's common prefix, that leaves it, that has no child under its
// byte, that is the end page of a node that has none, or that falls into a multi-key page, is
// not removed.
//
// Expected: false for each, the size unchanged, and every key still found.
func TestTreeRemoveAbsentKeys(t *testing.T) {
	var empty Tree
	if empty.remove([]byte("a"), nil) {
		t.Fatal("removed a key from an empty tree")
	}
	m := Map[uint64]{flat: 1} // single-key pages, so that the keys make nodes
	for _, k := range []string{"common-1", "common-2", "common-3", "other", "x"} {
		m.Add([]byte(k), 1)
		m.Add([]byte(k), 2)
	}
	rk := m.rekey
	for _, k := range []string{"", "comm", "common", "common-4", "commonX", "common-1-", "zzz", "otherwise", "w"} {
		if m.t.remove([]byte(k), rk) {
			t.Fatalf("removed %q, which is not in the tree", k)
		}
	}
	var p Map[uint64]
	for _, k := range []string{"page-1", "page-2", "page-3"} {
		p.Add([]byte(k), 1) // one value each: a multi-key page
	}
	if p.t.remove([]byte("page-1"), p.rekey) || p.t.remove([]byte("page-9"), p.rekey) {
		t.Fatal("the tree removed an entry of a multi-key page")
	}
	if m.Len() != 5 || p.Len() != 3 {
		t.Fatalf("sizes %d and %d", m.Len(), p.Len())
	}
	for _, k := range []string{"common-1", "common-2", "common-3", "other", "x"} {
		if !m.Has([]byte(k)) {
			t.Fatalf("lost %q", k)
		}
	}
}
