package mmbtree

import (
	"testing"

	"github.com/TomTonic/multimap/bench/proto/mmtest"
)

// TestAgainstReference checks both B-tree multimaps (values inline and behind
// a pointer) against a reference multimap, for single keys and key ranges.
func TestAgainstReference(t *testing.T) {
	t.Run("inline", func(t *testing.T) {
		mmtest.Check(t, func() mmtest.Multimap { return &Inline[uint64]{} })
	})
	t.Run("ptr", func(t *testing.T) {
		mmtest.Check(t, func() mmtest.Multimap { return &Ptr[uint64]{} })
	})
}
