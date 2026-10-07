//go:build !mkstats

package art

import "testing"

// TestShapeNeedsStats shows that the description of the routing part is a diagnosis of the stats build.
//
// A user of the map never sees it; the probes of step 6 print it from a build with the tag mkstats, and the
// ordinary build carries no code for it.
//
// Expected: Shape answers an empty string without the tag, for an empty and a filled map.
func TestShapeNeedsStats(t *testing.T) {
	var m Map[uint64]
	m.Add([]byte("a"), 1)
	if (&Map[uint64]{}).Shape() != "" || m.Shape() != "" {
		t.Error("Shape says something without the tag mkstats")
	}
}
