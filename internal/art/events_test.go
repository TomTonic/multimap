package art

import "testing"

// TestEvents checks the counters of the multi-key pages, which a build with the tag mkstats
// keeps for the probe of the benchmark (docs/redesign/step4-probe.md): a user of the
// library never sees them, and without the tag they do not exist.
//
// The counters belong to the multi-key pages of Map. The test makes two keys share a page
// and expects that build to count it, or to count nothing at all in a build without the tag.
func TestEvents(t *testing.T) {
	ResetEvents()
	var m Map[uint64]
	m.Add([]byte("alpha"), 1)
	m.Add([]byte("alphb"), 2)
	got := Events()
	if !EventsEnabled {
		if got != nil {
			t.Fatalf("Events() = %v without the tag mkstats, want nil", got)
		}
		return
	}
	count := func() (sum uint64) {
		for _, c := range Events() {
			sum += c.Total
		}
		return sum
	}
	if count() == 0 {
		t.Fatal("no event counted for a pair of keys that share a page")
	}
	ResetEvents()
	if n := count(); n != 0 {
		t.Fatalf("events after reset: %d, want 0", n)
	}
}
