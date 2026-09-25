package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/TomTonic/multimap/bench/stats"
)

var opOrder = []string{"valuesFor", "valuesBetween", "churn", "build"}

// writeSpeed writes speed-summary.md: one row per comparison with the median
// time per operation of both candidates, the median difference, its 95%
// interval across processes, and whether that interval met the stop rule.
func writeSpeed(c config, rows []result) error {
	type key struct {
		keys    string
		n       int
		op      string
		a, b    string
		opIndex int
	}
	groups := map[key][]result{}
	for _, r := range rows {
		k := key{r.Keys, r.N, r.Op, r.A, r.B, slices.Index(opOrder, r.Op)}
		groups[k] = append(groups[k], r)
	}
	ks := make([]key, 0, len(groups))
	for k := range groups {
		ks = append(ks, k)
	}
	slices.SortFunc(ks, func(x, y key) int {
		return cmpAll(strings.Compare(x.keys, y.keys), x.n-y.n, x.opIndex-y.opIndex,
			strings.Compare(x.a, y.a), strings.Compare(x.b, y.b))
	})
	var b strings.Builder
	b.WriteString("| keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |\n")
	b.WriteString("|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|\n")
	for _, k := range ks {
		g := groups[k]
		var d, h, na, nb []float64
		var res []bool
		for _, r := range g {
			d, h, res = append(d, r.Delta), append(h, (r.High-r.Low)/2), append(res, r.Resolved)
			na, nb = append(na, r.NsA), append(nb, r.NsB)
		}
		s := stats.Summarize(d, h, res)
		iv, fiv := "—", ""
		if !math.IsNaN(s.Low) {
			iv = fmt.Sprintf("[%+.1f%%, %+.1f%%]", s.Low*100, s.High*100)
			fiv = fmt.Sprintf(" [%.2f, %.2f]", speedup(s.Low), speedup(s.High))
		}
		fmt.Fprintf(&b, "| %s | %d | %s | %s | %s | %d | %s | %s | %.2f×%s | %+.1f%% | %s | %.1f pts | %.1f | %s |\n",
			k.keys, k.n, k.op, k.a, k.b, s.Procs, fmtNs(median(na)), fmtNs(median(nb)), speedup(s.Median), fiv,
			s.Median*100, iv, s.SD*100, s.Ratio, yesNo(s.Precise(c.abs, c.rel)))
	}
	b.WriteString("\nA speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.\n")
	return os.WriteFile(filepath.Join(c.out, "speed-summary.md"), []byte(b.String()), 0o644)
}

// writeMem writes mem-summary.md: medians over the rounds, with the GC cost of
// the baseline process (corpus only) subtracted.
func writeMem(c config, rows []memResult) error {
	type key struct{ keys, impl string }
	groups := map[key][]memResult{}
	for _, r := range rows {
		groups[key{r.Keys, r.Impl}] = append(groups[key{r.Keys, r.Impl}], r)
	}
	var b strings.Builder
	b.WriteString("| keys | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|\n")
	for _, kind := range c.kinds {
		base := median(field(groups[key{kind, "none"}], func(r memResult) float64 { return r.GCCPUMs }))
		for _, impl := range c.impls {
			g := groups[key{kind, impl}]
			if len(g) == 0 {
				continue
			}
			f := func(get func(memResult) float64) float64 { return median(field(g, get)) }
			fmt.Fprintf(&b, "| %s | %s | %d | %.0f | %.0f | %+.0f ms | %.0f |\n", kind, impl, len(g),
				f(func(r memResult) float64 { return r.HeapPerKey }),
				f(func(r memResult) float64 { return r.ScanPerKey }),
				f(func(r memResult) float64 { return r.GCCPUMs })-base,
				f(func(r memResult) float64 { return r.HalfHeapPerKey }))
		}
	}
	fmt.Fprintf(&b, "\n%d keys. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.\n", c.memN)
	return os.WriteFile(filepath.Join(c.out, "mem-summary.md"), []byte(b.String()), 0o644)
}

// speedup turns rtcompare's relative difference (1 - timeA/timeB) into how
// many times as fast A is as B.
func speedup(delta float64) float64 { return 1 / (1 - delta) }

func field[T any](rows []T, get func(T) float64) []float64 {
	out := make([]float64, len(rows))
	for i, r := range rows {
		out[i] = get(r)
	}
	return out
}

func median(xs []float64) float64 {
	if len(xs) == 0 {
		return math.NaN()
	}
	s := slices.Sorted(slices.Values(xs))
	if len(s)%2 == 1 {
		return s[len(s)/2]
	}
	return (s[len(s)/2-1] + s[len(s)/2]) / 2
}

func fmtNs(ns float64) string {
	switch {
	case ns >= 1e6:
		return fmt.Sprintf("%.2f ms", ns/1e6)
	case ns >= 1e4:
		return fmt.Sprintf("%.1f µs", ns/1e3)
	case ns >= 100:
		return fmt.Sprintf("%.0f", ns)
	default:
		return fmt.Sprintf("%.1f", ns)
	}
}

func yesNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

func cmpAll(cs ...int) int {
	for _, c := range cs {
		if c != 0 {
			return c
		}
	}
	return 0
}
