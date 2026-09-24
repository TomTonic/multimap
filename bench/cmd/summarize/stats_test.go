package main

import (
	"math"
	"testing"
)

// TestSummarize makes sure a benchmark comparison that ran in several
// processes gets an honest pooled answer: the typical difference, and an
// interval wide enough for the processes' disagreement. It covers the
// summarize command of the benchmark suite, which treats each process (one
// heap layout) as one observation. The cases check the median, the sample
// standard deviation, the t-based interval, the ratio to the in-run standard
// error, and the degenerate single-process case.
func TestSummarize(t *testing.T) {
	tests := []struct {
		name                 string
		deltas, halves       []float64
		resolved             []bool
		median, sd, low, hig float64
		ratio                float64
		nResolved            int
	}{
		{
			name:   "pools three processes with a t interval",
			deltas: []float64{0.10, 0.20, 0.30}, halves: []float64{0.0196, 0.0196, 0.0196},
			resolved: []bool{true, true, false},
			median:   0.20, sd: 0.10, low: 0.20 - 4.303*0.10/math.Sqrt(3), hig: 0.20 + 4.303*0.10/math.Sqrt(3),
			ratio: 10, nResolved: 2,
		},
		{
			name:   "takes the mean of the middle pair as median for even counts",
			deltas: []float64{-0.02, 0.01, 0.03, 0.02}, halves: []float64{0.01, 0.02, 0.01, 0.02},
			resolved: []bool{false, false, false, false},
			median:   0.015, sd: math.Sqrt((0.03*0.03 + 0 + 0.02*0.02 + 0.01*0.01) / 3), // mean 0.01
			ratio: -1, nResolved: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := summarize(tt.deltas, tt.halves, tt.resolved)
			near := func(what string, got, want float64) {
				if math.Abs(got-want) > 1e-9 {
					t.Errorf("%s = %v, want %v", what, got, want)
				}
			}
			near("median", s.Median, tt.median)
			near("sd", s.SD, tt.sd)
			if tt.low != 0 || tt.hig != 0 {
				near("low", s.Low, tt.low)
				near("high", s.High, tt.hig)
			}
			if tt.ratio >= 0 {
				near("ratio", s.Ratio, tt.ratio)
			}
			if s.Resolved != tt.nResolved {
				t.Errorf("resolved = %d, want %d", s.Resolved, tt.nResolved)
			}
		})
	}
	t.Run("gives no interval for a single process", func(t *testing.T) {
		s := summarize([]float64{0.05}, []float64{0.01}, []bool{true})
		if !math.IsNaN(s.Low) || !math.IsNaN(s.High) || s.SD != 0 || s.Median != 0.05 {
			t.Errorf("single process: %+v", s)
		}
	})
}

// TestGroupKey makes sure the results of one comparison from different
// processes are pooled together and never with another comparison. It covers
// the grouping in the summarize command: fields that describe the comparison
// (keys, size, operation, candidates) form the key, fields that vary per
// process (timings, interval, layout seed) do not, and the key does not depend
// on field order.
func TestGroupKey(t *testing.T) {
	a := map[string]any{"keys": "str", "n": 4096.0, "op": "get", "a": "x", "b": "y", "delta": 0.1, "layout_seed": 1.0}
	b := map[string]any{"b": "y", "a": "x", "op": "get", "n": 4096.0, "keys": "str", "delta": -0.3, "layout_seed": 2.0}
	c := map[string]any{"keys": "str", "n": 4096.0, "op": "miss", "a": "x", "b": "y", "delta": 0.1}
	if groupKey(a) != groupKey(b) {
		t.Errorf("same comparison, different keys: %q vs %q", groupKey(a), groupKey(b))
	}
	if groupKey(a) == groupKey(c) {
		t.Errorf("different comparisons share key %q", groupKey(a))
	}
}
