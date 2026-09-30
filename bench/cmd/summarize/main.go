// Command summarize pools the results of comparisons that ran in several
// processes (see cmd/bench) with rtcompare.CombineStaged, or rtcompare.Combine
// for rows of a run before rtcompare v0.8.0, and prints one row per
// comparison: the mean difference, its 95% interval across processes, the
// spread between processes, and how far that spread exceeds the interval a
// single process reports. Rows of the serial and the parallel regime (the
// "parallel" field) are pooled separately, never together.
//
// Usage:
//
//	go run ./cmd/summarize results/speed.jsonl
//	go run ./cmd/summarize -json results/*.jsonl
//
// Rows without a delta (memory results) are skipped. Differences are
// relative: positive means candidate A is faster.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/TomTonic/rtcompare"
)

// measured are the fields of a result row that vary between the processes of
// one comparison; all others name the comparison.
var measured = map[string]bool{
	"ns_a": true, "ns_b": true, "delta": true, "low": true, "high": true, "level": true,
	"resolved": true, "validated": true, "noise_floor": true, "autocorr": true, "inner_loops": true,
	"layout_seed": true, "seed": true, "process": true, "suspended_s": true, "live_heap": true, "warnings": true,
	"first_stage": true, "aa_a": true, "aa_b": true, "quantization": true, "report_seed": true,
}

func main() {
	asJSON := flag.Bool("json", false, "print JSON lines instead of a table")
	flag.Parse()
	var order []string
	groups := map[string]*group{}
	for _, path := range flag.Args() {
		if err := read(path, groups, &order); err != nil {
			fail(err)
		}
	}
	slices.Sort(order)
	enc := json.NewEncoder(os.Stdout)
	if !*asJSON {
		fmt.Println("| comparison | procs | delta | 95% across procs | sd between | inflation | I² | resolved |")
		fmt.Println("|---|---:|---:|---|---:|---:|---:|---|")
	}
	for _, k := range order {
		p, err := groups[k].pool()
		if err != nil {
			fmt.Fprintf(os.Stderr, "summarize: %s: %v\n", k, err)
			continue
		}
		if *asJSON {
			if err := enc.Encode(struct {
				Comparison string `json:"comparison"`
				rtcompare.Pooled
			}{k, p}); err != nil {
				fail(err)
			}
			continue
		}
		fmt.Printf("| %s | %d | %+.1f%% | [%+.1f, %+.1f]%% | %.1f pts | %.1f | %.2f | %v |\n",
			k, p.Processes, p.Delta*100, p.Low*100, p.High*100, p.SpreadBetween*100, p.Inflation, p.I2, p.Resolved)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "summarize:", err)
	os.Exit(1)
}

// group is the reports of one comparison, one per process, and the size of the
// first stage of the run that made them, zero if the rows do not say.
type group struct {
	reports []rtcompare.Report
	first   int
}

// pool pools the reports the way multiproc did in the run: with Stein's
// interval from the recorded first stage, or with a plain t interval when the
// rows lack one.
func (g *group) pool() (rtcompare.Pooled, error) {
	if g.first == 0 {
		return rtcompare.Combine(g.reports, 0)
	}
	return rtcompare.CombineStaged(g.reports, g.first, 0)
}

// read adds every result row of one JSON lines file to its comparison's
// reports, rebuilt from the fields that rtcompare.Combine and CombineStaged
// pool.
func read(path string, groups map[string]*group, order *[]string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	for line := 1; sc.Scan(); line++ {
		var row map[string]any
		if err := json.Unmarshal(sc.Bytes(), &row); err != nil {
			return fmt.Errorf("%s:%d: %w", path, line, err)
		}
		d, ok := row["delta"].(float64)
		if !ok {
			continue
		}
		num := func(k string) float64 { v, _ := row[k].(float64); return v }
		res, _ := row["resolved"].(bool)
		val, _ := row["validated"].(bool)
		level := num("level")
		if level == 0 {
			level = 0.95 // rows from before the level was recorded
		}
		r := rtcompare.Report{
			NsPerOpA: num("ns_a"), NsPerOpB: num("ns_b"), Resolved: res, NoiseFloor: num("noise_floor"),
			Validated:    val || num("noise_floor") > 0,
			Estimate:     rtcompare.Estimate{Delta: d, Low: num("low"), High: num("high"), Level: level},
			ValidationA:  rtcompare.HarnessValidation{Deltas: floats(row["aa_a"])},
			ValidationB:  rtcompare.HarnessValidation{Deltas: floats(row["aa_b"])},
			Suspended:    time.Duration(num("suspended_s") * float64(time.Second)),
			Quantization: num("quantization"),
		}
		k := groupKey(row)
		g, ok := groups[k]
		if !ok {
			g = &group{}
			groups[k] = g
			*order = append(*order, k)
		}
		g.reports = append(g.reports, r)
		g.first = max(g.first, int(num("first_stage")))
	}
	return sc.Err()
}

// floats decodes a JSON array of numbers, nil for anything else.
func floats(v any) []float64 {
	a, _ := v.([]any)
	var out []float64
	for _, x := range a {
		if f, ok := x.(float64); ok {
			out = append(out, f)
		}
	}
	return out
}

// groupKey names the comparison a decoded result row belongs to, so that the
// rows of all its processes can be pooled. Fields are sorted, so the key does
// not depend on field order.
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
