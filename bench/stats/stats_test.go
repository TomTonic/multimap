package stats

import (
	"math"
	"testing"
)

// TestSummarize makes sure a benchmark comparison that ran in several
// processes gets an honest pooled answer: the typical difference, and an
// interval wide enough for the processes' disagreement. It covers the pooling
// used by the benchmark driver and the summarize command, which treat each
// process (one heap layout) as one observation. The cases check the median, the sample
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
			s := Summarize(tt.deltas, tt.halves, tt.resolved)
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
		s := Summarize([]float64{0.05}, []float64{0.01}, []bool{true})
		if !math.IsNaN(s.Low) || !math.IsNaN(s.High) || s.SD != 0 || s.Median != 0.05 {
			t.Errorf("single process: %+v", s)
		}
	})
}

// TestGroupKey makes sure the results of one comparison from different
// processes are pooled together and never with another comparison. It covers
// the grouping in the benchmark's stats package: fields that describe the comparison
// (keys, size, operation, candidates) form the key, fields that vary per
// process (timings, interval, layout seed) do not, and the key does not depend
// on field order.
func TestGroupKey(t *testing.T) {
	a := map[string]any{"keys": "str", "n": 4096.0, "op": "get", "a": "x", "b": "y", "delta": 0.1, "layout_seed": 1.0}
	b := map[string]any{"b": "y", "a": "x", "op": "get", "n": 4096.0, "keys": "str", "delta": -0.3, "layout_seed": 2.0}
	c := map[string]any{"keys": "str", "n": 4096.0, "op": "miss", "a": "x", "b": "y", "delta": 0.1}
	if GroupKey(a) != GroupKey(b) {
		t.Errorf("same comparison, different keys: %q vs %q", GroupKey(a), GroupKey(b))
	}
	if GroupKey(a) == GroupKey(c) {
		t.Errorf("different comparisons share key %q", GroupKey(a))
	}
}

// TestPrecise makes sure the benchmark driver stops adding processes exactly
// when a comparison is known well enough, and not before. It covers the stop
// rule of the benchmark's stats package: an interval across processes is
// precise when its half-width is within an absolute bound or within a
// fraction of the difference itself; with fewer than two processes there is
// no interval and it never is.
func TestPrecise(t *testing.T) {
	tests := []struct {
		name string
		s    Summary
		want bool
	}{
		{"accepts a narrow interval around a small difference", Summary{Mean: 0.01, Low: 0.0, High: 0.02}, true},
		{"rejects a wide interval around a small difference", Summary{Mean: 0.01, Low: -0.04, High: 0.06}, false},
		{"accepts a wide interval that is small relative to a large difference", Summary{Mean: 0.60, Low: 0.55, High: 0.65}, true},
		{"rejects a single process", Summary{Mean: 0.5, Low: math.NaN(), High: math.NaN()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Precise(0.02, 0.1); got != tt.want {
				t.Errorf("Precise = %v, want %v (half-width %v)", got, tt.want, tt.s.HalfWidth())
			}
		})
	}
}
