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
