package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

// measured are the fields that differ between processes of the same
// comparison; every other field names the comparison.
var measured = map[string]bool{
	"ns_a": true, "ns_b": true, "delta": true, "low": true, "high": true,
	"resolved": true, "noise_floor": true, "autocorr": true, "inner_loops": true,
	"layout_seed": true, "warnings": true,
}

// groupKey names the comparison a result row belongs to, so that the rows of
// all its processes end up together; fields are sorted for a stable key.
func groupKey(row map[string]any) string {
	var parts []string
	for k, v := range row {
		if measured[k] || v == nil {
			continue
		}
		if f, ok := v.(float64); ok { // JSON numbers: print 1048576, not 1.048576e+06
			v = strconv.FormatFloat(f, 'f', -1, 64)
		}
		parts = append(parts, fmt.Sprintf("%s=%v", k, v))
	}
	slices.Sort(parts)
	return strings.Join(parts, " ")
}

// summary describes one comparison across the processes that measured it.
type summary struct {
	Procs    int     `json:"procs"`
	Median   float64 `json:"median"`
	Mean     float64 `json:"mean"`
	SD       float64 `json:"sd"`   // between processes
	Low      float64 `json:"low"`  // 95% t-interval of the mean; NaN below 2 processes
	High     float64 `json:"high"` //
	InRun    float64 `json:"in_run_half_width"`
	Ratio    float64 `json:"ratio"` // SD over the standard error one run reports
	Resolved int     `json:"resolved"`
}

// summarize pools the relative differences (deltas) that separate processes
// measured for one comparison, with the half-widths of their in-run 95%
// intervals and whether each run resolved the difference. The interval it
// returns treats each process as one observation, so it includes the
// variation between heap layouts that no single run can see.
func summarize(deltas, halfWidths []float64, resolved []bool) summary {
	k := len(deltas)
	s := summary{Procs: k, Low: math.NaN(), High: math.NaN()}
	if k == 0 {
		return s
	}
	sorted := slices.Sorted(slices.Values(deltas))
	s.Median = sorted[k/2]
	if k%2 == 0 {
		s.Median = (sorted[k/2-1] + sorted[k/2]) / 2
	}
	for _, d := range deltas {
		s.Mean += d
	}
	s.Mean /= float64(k)
	hw := slices.Sorted(slices.Values(halfWidths))
	s.InRun = hw[len(hw)/2]
	for _, r := range resolved {
		if r {
			s.Resolved++
		}
	}
	if k < 2 {
		return s
	}
	var ss float64
	for _, d := range deltas {
		ss += (d - s.Mean) * (d - s.Mean)
	}
	s.SD = math.Sqrt(ss / float64(k-1))
	half := tQuantile975(k-1) * s.SD / math.Sqrt(float64(k))
	s.Low, s.High = s.Mean-half, s.Mean+half
	if s.InRun > 0 {
		s.Ratio = s.SD / (s.InRun / 1.96)
	}
	return s
}

// tTable holds the 97.5% quantiles of Student's t distribution by degrees of
// freedom; the benchmark uses a handful of processes, where 1.96 would be far
// too optimistic.
var tTable = []float64{0, 12.706, 4.303, 3.182, 2.776, 2.571, 2.447, 2.365, 2.306, 2.262, 2.228,
	2.201, 2.179, 2.160, 2.145, 2.131, 2.120, 2.110, 2.101, 2.093, 2.086,
	2.080, 2.074, 2.069, 2.064, 2.060, 2.056, 2.052, 2.048, 2.045, 2.042}

// tQuantile975 returns the two-sided 95% critical value for df degrees of
// freedom; beyond the table it uses the value for 30, which errs wide.
func tQuantile975(df int) float64 {
	if df < 1 {
		return math.NaN()
	}
	return tTable[min(df, len(tTable)-1)]
}
