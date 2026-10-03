package artstr

import (
	"sort"
	"testing"

	"github.com/TomTonic/multimap/internal/art"
)

// TestShape prints the objects of a tree of 50,000 single-valued keys next to the
// tree of today (go test -run Shape -v).
func TestShape(t *testing.T) {
	keys, vals := benchKeys(50000)
	var a art.Map[string]
	var s Map[string]
	for i, k := range keys {
		a.Add(k, vals[i])
		s.Add(k, vals[i])
	}
	count := func(each func(func(Object))) {
	}
	_ = count
	type agg struct{ n, bytes int }
	got := map[string]*agg{}
	s.Objects(func(o Object) {
		g := got[o.Label]
		if g == nil {
			g = &agg{}
			got[o.Label] = g
		}
		g.n++
		g.bytes += o.Size
	})
	want := map[string]*agg{}
	a.Objects(func(o art.Object) {
		g := want[o.Label]
		if g == nil {
			g = &agg{}
			want[o.Label] = g
		}
		g.n++
		g.bytes += o.Size
	})
	for name, m := range map[string]map[string]*agg{"artstr": got, "art": want} {
		var labels []string
		total := 0
		for l, g := range m {
			labels = append(labels, l)
			total += g.bytes
		}
		sort.Strings(labels)
		for _, l := range labels {
			t.Logf("%-7s %-12s %6d objects %8d bytes", name, l, m[l].n, m[l].bytes)
		}
		t.Logf("%-7s total %d bytes = %.1f per key", name, total, float64(total)/50000)
	}
}
