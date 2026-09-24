package arenaart

import (
	"testing"

	"github.com/TomTonic/multimap/bench/proto/prototest"
)

// TestAgainstReference checks that the arena-based ART prototype stores,
// finds and orders keys exactly like a sorted reference map, across corpora
// that exercise prefix keys, long keys and every node kind.
func TestAgainstReference(t *testing.T) {
	prototest.Check(t, func() prototest.Map { return &Tree{} })
}
