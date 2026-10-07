package art

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

// TestScanOfDeepTrees shows that a range scan finds every key of a tree deeper than the cursor's fixed stack.
//
// A user whose keys share long paths that fork again and again (deep directory trees) gets every key and value
// of a scan in order: the walk keeps a frame for every byte node on the way, 32 of them in the cursor itself and
// the rest in a slice that it reuses when the walk goes deep a second time.
//
// Expected: two chains of 40 forking nodes each (keys that differ at every level and are too long to share a
// page) are scanned completely and in order, by keys and by values, and the bounded scan of the second chain
// finds exactly its keys.
func TestScanOfDeepTrees(t *testing.T) {
	var m Map[uint64]
	tail := strings.Repeat("t", 300) // two such keys do not fit one page: every fork is a byte node
	var keys [][]byte
	for _, side := range []string{"x", "y"} {
		for i := range 40 {
			keys = append(keys, []byte(side+strings.Repeat("a", i)+"b"+tail))
		}
	}
	slices.SortFunc(keys, bytes.Compare)
	for i, k := range keys {
		m.Add(k, uint64(i))
	}
	var got [][]byte
	m.Range(&Bounds{}, func(k []byte) bool { got = append(got, slices.Clone(k)); return true })
	if !slices.EqualFunc(got, keys, bytes.Equal) {
		t.Fatalf("the scan of keys found %d of %d keys or not in order", len(got), len(keys))
	}
	var vals []uint64
	m.RangeValues(&Bounds{}, func(v uint64) bool { vals = append(vals, v); return true })
	for i, v := range vals {
		if v != uint64(i) {
			t.Fatalf("value %d is %d", i, v)
		}
	}
	if len(vals) != len(keys) {
		t.Fatalf("%d values, want %d", len(vals), len(keys))
	}
	n := 0
	m.RangeValues(&Bounds{From: []byte("y"), HasFrom: true, FromIncl: true}, func(uint64) bool { n++; return true })
	if n != 40 {
		t.Fatalf("the scan from y found %d values, want 40", n)
	}
}
