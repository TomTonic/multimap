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

// poolRows pools the rows of one comparison the way multiproc pooled them in
// the run: with Stein's interval (rtcompare.CombineStaged) from the first
// stage the run recorded, which the rows of a run all agree on. Rows without a
// first stage, from runs before rtcompare v0.8.0, are pooled with a plain t
// interval.
func poolRows(g []result) (rtcompare.Pooled, error) {
	reps := make([]rtcompare.Report, len(g))
	first := 0
	for i, r := range g {
		reps[i] = reportOf(r)
		first = max(first, r.FirstStage)
	}
	if first == 0 {
		return rtcompare.Combine(reps, 0)
	}
	return rtcompare.CombineStaged(reps, first, 0)
}

// regime describes how the processes of a run ran, for the summary's notes:
// the serial one, or the parallel one, whose figures are not comparable with
// serial ones.
func regime(c config) string {
	if c.parallel > 1 {
		return fmt.Sprintf("parallel: %d processes at a time, each with GOMAXPROCS %d, sharing the caches and the memory bandwidth", c.parallel, c.childProcs)
	}
	return "serial: one process at a time, each with the machine to itself"
}

// writeSpeed writes speed-summary.md: one row per comparison with the median
// time per operation of both candidates, the pooled difference, its 95%
// interval across processes, and whether that interval is as narrow as the
// run was sized for.
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
	b.WriteString("| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |\n")
	b.WriteString("|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|\n")
	var notes []string
	for _, k := range ks {
		g := groups[k]
		var na, nb []float64
		for _, r := range g {
			na, nb = append(na, r.NsA), append(nb, r.NsB)
		}
		name := fmt.Sprintf("%s %s n=%d %s: %s vs %s", k.values, k.keys, k.n, k.op, k.a, k.b)
		p, err := poolRows(g)
		if err != nil {
			fmt.Fprintf(&b, "| %s | %s | %d | %s | %s | %s | %d | %s | %s | — | — | — | — | — | no | no |\n",
				k.values, k.keys, k.n, k.op, k.a, k.b, len(g), fmtNs(median(na)), fmtNs(median(nb)))
			notes = append(notes, fmt.Sprintf("- %s: %v", name, err))
			continue
		}
		ratio, low, high := rtcompare.Estimate{Delta: p.Delta, Low: p.Low, High: p.High}.Ratio()
		fmt.Fprintf(&b, "| %s | %s | %d | %s | %s | %s | %d | %s | %s | %.2f× [%.2f, %.2f] | %+.1f%% | [%+.1f%%, %+.1f%%] | %.1f pts | %.1f | %s | %s |\n",
			k.values, k.keys, k.n, k.op, k.a, k.b, p.Processes, fmtNs(median(na)), fmtNs(median(nb)),
			ratio, low, high, p.Delta*100, p.Low*100, p.High*100,
			p.SpreadBetween*100, p.Inflation, yesNo(p.Precise(c.abs, c.rel)), yesNo(p.Resolved))
		for _, w := range p.Warnings {
			notes = append(notes, fmt.Sprintf("- %s: %s", name, w))
		}
	}
	b.WriteString("\nRegime: " + regime(c) + ".\n")
	b.WriteString("\nA speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).\n")
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
		label := profile + valueTag
		for _, kind := range c.kinds {
			base := median(field(groups[key{label, kind, "none"}], func(r memResult) float64 { return r.GCCPUMs }))
			for _, impl := range implsFor(profile) {
				g := groups[key{label, kind, impl}]
				if len(g) == 0 {
					continue
				}
				f := func(get func(memResult) float64) float64 { return median(field(g, get)) }
				fmt.Fprintf(&b, "| %s | %s | %d | %s | %d | %.0f | %.0f | %+.0f ms | %.0f |\n", label, kind, g[0].N, impl, len(g),
					f(func(r memResult) float64 { return r.HeapPerKey }),
					f(func(r memResult) float64 { return r.ScanPerKey }),
					f(func(r memResult) float64 { return r.GCCPUMs })-base,
					f(func(r memResult) float64 { return r.HalfHeapPerKey }))
			}
		}
	}
	var notes []string
	for _, profile := range c.profiles {
		for _, kind := range c.kinds {
			if g := groups[key{profile + valueTag, kind, ordered}]; len(g) > 0 && g[0].ValueBytes > 0 {
				notes = append(notes, fmt.Sprintf("- %s %s%s: the values are %.0f bytes of string per key. Only ordered-lpage holds them in its heap; every other candidate holds 16-byte headers that point into one shared buffer (see toVs), so its heap figure does not count them.", profile+valueTag, kind, "", g[0].ValueBytes))
			}
		}
	}
	if len(notes) > 0 {
		b.WriteString("\n" + strings.Join(notes, "\n") + "\n")
	}
	b.WriteString("\nn: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.\n")
	return os.WriteFile(filepath.Join(c.out, "mem-summary.md"), []byte(b.String()), 0o644)
}

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
