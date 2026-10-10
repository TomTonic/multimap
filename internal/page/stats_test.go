package page

import "testing"

// TestLayCallsCount: a probe that asks how often a write computes the layout of a page reads LayCalls. In the page
// package's statistics (counted only in a build with the tag mkstats), the count never goes down while pages are made
// and read, and without the tag it stays 0.
func TestLayCallsCount(t *testing.T) {
	before := LayCalls()
	p := BuildFixedOf([][]byte{[]byte("ab"), []byte("ac")}, []uint64{1, 2}, false)
	if p == nil || p.Len() != 2 {
		t.Fatalf("page of two keys: %v", p)
	}
	if after := LayCalls(); after < before {
		t.Fatalf("LayCalls went from %d down to %d", before, after)
	}
}
