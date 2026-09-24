package main

import (
	"math"
	"slices"
	"testing"
)

// TestPairsFor makes sure the benchmark compares exactly what a user choosing
// a multimap wants compared, and skips what cannot be timed sensibly. It
// covers the scenario plan of the benchmark driver: Ordered against every
// alternative and Hashed against the hand-written map for point operations,
// no range queries for the unordered map-sets, Hashed range queries (a scan of
// all keys) and builds only up to their size limits.
func TestPairsFor(t *testing.T) {
	ops := []string{"valuesFor", "valuesBetween", "addRemove", "build"}
	tests := []struct {
		name string
		n    int
		want []pair
		skip []pair
	}{
		{
			name: "compares everything for small scenarios",
			n:    4096,
			want: []pair{{"valuesFor", ordered, mapSets}, {"valuesFor", hashed, mapSets},
				{"valuesBetween", ordered, hashed}, {"valuesBetween", ordered, btreeSets}, {"build", ordered, btreeSets}},
		},
		{
			name: "drops hashed range queries and builds for large scenarios",
			n:    1 << 20,
			want: []pair{{"valuesBetween", ordered, btreeSets}, {"addRemove", hashed, mapSets}},
			skip: []pair{{"valuesBetween", ordered, hashed}, {"build", ordered, hashed}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pairsFor(tt.n, ops, 1<<16, 1<<16)
			for _, p := range tt.want {
				if !slices.Contains(got, p) {
					t.Errorf("missing %v", p)
				}
			}
			for _, p := range tt.skip {
				if slices.Contains(got, p) {
					t.Errorf("unexpected %v", p)
				}
			}
			for _, p := range got {
				if p.op == "valuesBetween" && (p.a == mapSets || p.b == mapSets) {
					t.Errorf("range query on the unordered map-sets: %v", p)
				}
			}
		})
	}
}

// TestSpeedup makes sure the summary states speed the way a reader expects:
// "2.00×" when A is twice as fast, "0.50×" when it is half as fast. It covers
// the conversion of rtcompare's relative difference (1 - timeA/timeB) in the
// benchmark driver's summary table.
func TestSpeedup(t *testing.T) {
	for _, tt := range []struct{ delta, want float64 }{{0.5, 2}, {0, 1}, {-1, 0.5}} {
		if got := speedup(tt.delta); math.Abs(got-tt.want) > 1e-12 {
			t.Errorf("speedup(%v) = %v, want %v", tt.delta, got, tt.want)
		}
	}
}

// TestMedian makes sure the summary tables report the typical process, not
// an outlier. It covers the median helper of the benchmark driver for odd
// and even counts and for no values at all.
func TestMedian(t *testing.T) {
	if got := median([]float64{3, 1, 2}); got != 2 {
		t.Errorf("odd: %v", got)
	}
	if got := median([]float64{4, 1, 3, 2}); got != 2.5 {
		t.Errorf("even: %v", got)
	}
	if got := median(nil); !math.IsNaN(got) {
		t.Errorf("empty: %v", got)
	}
}
