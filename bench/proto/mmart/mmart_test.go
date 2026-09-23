package mmart

import (
	"testing"
	"unsafe"

	"github.com/TomTonic/multimap/bench/proto/mmtest"
)

// TestAgainstReference checks that the ART multimap returns exactly the values
// a reference multimap holds, for single keys and for key ranges in key order.
func TestAgainstReference(t *testing.T) {
	mmtest.Check(t, func() mmtest.Multimap { return &Map[uint64]{} })
}

// TestLeafLayout guards the leaf size the value layout was chosen for: 80 B is
// a Go size class; one more inline value would move leaves to 96 B.
func TestLeafLayout(t *testing.T) {
	if got := unsafe.Sizeof(leaf[uint64]{}); got != 80 {
		t.Fatalf("leaf[uint64] is %d bytes, want 80", got)
	}
}
