package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/TomTonic/rtcompare"
)

var opOrder = []string{"valuesFor", "valuesBetween", "prefix", "churn", "build"}

// writeSpeed writes speed-summary.md: one row per comparison with the median
// time per operation of both candidates, the median difference, its 95%
// interval across processes, and whether that interval met the stop rule.
func writeSpeed(c config, rows []result) error {
	type key struct {
		keys    string
		values  string
		n       int
		op      string
		a, b    string
		opIndex int
	}
	groups := map[key][]result{}
	for _, r := range rows {
		k := key{r.Keys, r.Values, r.N, r.Op, r.A, r.B, slices.Index(opOrder, r.Op)}
		groups[k] = append(groups[k], r)
	}
	ks := make([]key, 0, len(groups))
	for k := range groups {
		ks = append(ks, k)
	}
	slices.SortFunc(ks, func(x, y key) int {
		return cmpAll(strings.Compare(x.values, y.values), strings.Compare(x.keys, y.keys), x.n-y.n, x.opIndex-y.opIndex,
			strings.Compare(x.a, y.a), strings.Compare(x.b, y.b))
	})
	var b strings.Builder
	b.WriteString("| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |\n")
	b.WriteString("|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|\n")
	var notes []string
	for _, k := range ks {
		g := groups[k]
		reps := make([]rtcompare.Report, len(g))
		var na, nb []float64
		for i, r := range g {
			reps[i] = reportOf(r)
			na, nb = append(na, r.NsA), append(nb, r.NsB)
		}
		name := fmt.Sprintf("%s %s n=%d %s: %s vs %s", k.values, k.keys, k.n, k.op, k.a, k.b)
		p, err := rtcompare.Combine(reps, 0)
		if err != nil {
			fmt.Fprintf(&b, "| %s | %s | %d | %s | %s | %s | %d | %s | %s | — | — | — | — | — | no |\n",
				k.values, k.keys, k.n, k.op, k.a, k.b, len(g), fmtNs(median(na)), fmtNs(median(nb)))
			notes = append(notes, fmt.Sprintf("- %s: %v", name, err))
			continue
		}
		fmt.Fprintf(&b, "| %s | %s | %d | %s | %s | %s | %d | %s | %s | %.2f× [%.2f, %.2f] | %+.1f%% | [%+.1f%%, %+.1f%%] | %.1f pts | %.1f | %s |\n",
			k.values, k.keys, k.n, k.op, k.a, k.b, p.Processes, fmtNs(median(na)), fmtNs(median(nb)),
			speedup(p.Delta), speedup(p.Low), speedup(p.High), p.Delta*100, p.Low*100, p.High*100,
			p.SpreadBetween*100, p.Inflation, yesNo(p.Precise(c.abs, c.rel)))
		for _, w := range p.Warnings {
			notes = append(notes, fmt.Sprintf("- %s: %s", name, w))
		}
	}
	b.WriteString("\nA speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.\n")
	if len(notes) > 0 {
		b.WriteString("\nWarnings from pooling:\n\n" + strings.Join(notes, "\n") + "\n")
	}
	return os.WriteFile(filepath.Join(c.out, "speed-summary.md"), []byte(b.String()), 0o644)
}

// writeMem writes mem-summary.md: medians over the rounds, with the GC cost of
// the baseline process (corpus only) subtracted.
func writeMem(c config, rows []memResult) error {
	type key struct{ values, keys, impl string }
	groups := map[key][]memResult{}
	for _, r := range rows {
		k := key{r.Values, r.Keys, r.Impl}
		groups[k] = append(groups[k], r)
	}
	var b strings.Builder
	b.WriteString("| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |\n")
	b.WriteString("|---|---|---:|---|---:|---:|---:|---:|---:|\n")
	for _, profile := range c.profiles {
		for _, kind := range c.kinds {
			base := median(field(groups[key{profile, kind, "none"}], func(r memResult) float64 { return r.GCCPUMs }))
			for _, impl := range implsFor(profile) {
				g := groups[key{profile, kind, impl}]
				if len(g) == 0 {
					continue
				}
				f := func(get func(memResult) float64) float64 { return median(field(g, get)) }
				fmt.Fprintf(&b, "| %s | %s | %d | %s | %d | %.0f | %.0f | %+.0f ms | %.0f |\n", profile, kind, g[0].N, impl, len(g),
					f(func(r memResult) float64 { return r.HeapPerKey }),
					f(func(r memResult) float64 { return r.ScanPerKey }),
					f(func(r memResult) float64 { return r.GCCPUMs })-base,
					f(func(r memResult) float64 { return r.HalfHeapPerKey }))
			}
		}
	}
	b.WriteString("\nn: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.\n")
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
