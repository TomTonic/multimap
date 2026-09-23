package mmart2

import (
	"testing"
	"unsafe"

	"github.com/TomTonic/multimap/bench/proto/mmtest"
)

// TestAgainstReference checks that the ART multimap with node25 and 9 inline values per leaf returns exactly the values
// a reference multimap holds, for single keys and for key ranges in key order.
func TestAgainstReference(t *testing.T) {
	mmtest.Check(t, func() mmtest.Multimap { return &Map[uint64]{} })
}

// TestLeafLayout guards the sizes the layout was chosen for: 128 and 256
// bytes are Go size classes.
func TestLeafLayout(t *testing.T) {
	if got := unsafe.Sizeof(leaf[uint64]{}); got != 128 {
		t.Fatalf("leaf[uint64] is %d bytes, want 128", got)
	}
	if got := unsafe.Sizeof(node25{}); got != 256 {
		t.Fatalf("node25 is %d bytes, want 256", got)
	}
}
