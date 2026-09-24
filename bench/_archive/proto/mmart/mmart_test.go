package mmart

import (
	"iter"
	"testing"
	"unsafe"

	"github.com/TomTonic/multimap/bench/proto/mmtest"
)

// TestAgainstReference checks that the ART multimap returns exactly the values
// a reference multimap holds, for single keys and for key ranges in key order.
func TestAgainstReference(t *testing.T) {
	mmtest.Check(t, func() mmtest.Multimap { return &Map[uint64]{} })
}

// touchMap routes ValuesBetween to ValuesBetweenTouch for the reference test.
type touchMap struct{ *Map[uint64] }

func (m touchMap) ValuesBetween(from, to []byte) iter.Seq[uint64] {
	return m.ValuesBetweenTouch(from, to)
}

// TestTouchAgainstReference makes sure that range reads return the same values
// in the same order whether or not the scan touches children ahead. It covers
// the benchmark's ART multimap prototype, whose touch-ahead scan variant only
// changes memory access timing, and checks it against the reference multimap
// for the same key sets and ranges as TestAgainstReference.
func TestTouchAgainstReference(t *testing.T) {
	mmtest.Check(t, func() mmtest.Multimap { return touchMap{&Map[uint64]{}} })
}

// TestLeafLayout guards the leaf size the value layout was chosen for: 80 B is
// a Go size class; one more inline value would move leaves to 96 B.
func TestLeafLayout(t *testing.T) {
	if got := unsafe.Sizeof(leaf[uint64]{}); got != 80 {
		t.Fatalf("leaf[uint64] is %d bytes, want 80", got)
	}
	var l leaf[uint64]
	if got := unsafe.Offsetof(l.vals) + unsafe.Sizeof(l.vals) - 8; got != leafTailOff {
		t.Fatalf("last word of leaf[uint64] is at %d, leafTailOff is %d", got, leafTailOff)
	}
}
