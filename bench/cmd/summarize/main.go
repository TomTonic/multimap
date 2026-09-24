// Command summarize pools the results of comparisons that ran in several
// processes (see -layoutseed and PROCS in the run scripts) and prints one row
// per comparison: the median difference, the 95% interval across processes,
// the spread between processes, and how that spread compares with the
// interval a single rtcompare run reports.
//
// Usage:
//
//	go run ./cmd/summarize results/mm.jsonl results/hot.jsonl
//	go run ./cmd/summarize -json results/*.jsonl
//
// Rows without a delta (memgc output) are skipped. Differences are relative:
// positive means candidate A is faster.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
)

type group struct {
	key      string
	deltas   []float64
	halves   []float64
	resolved []bool
}

func main() {
	asJSON := flag.Bool("json", false, "print JSON lines instead of a table")
	flag.Parse()
	var order []string
	groups := map[string]*group{}
	for _, path := range flag.Args() {
		if err := read(path, groups, &order); err != nil {
			fmt.Fprintln(os.Stderr, "summarize:", err)
			os.Exit(1)
		}
	}
	slices.Sort(order)
	enc := json.NewEncoder(os.Stdout)
	if !*asJSON {
		fmt.Println("| comparison | procs | median | 95% across procs | sd between | in-run ± | ratio | resolved |")
		fmt.Println("|---|---:|---:|---|---:|---:|---:|---:|")
	}
	for _, k := range order {
		g := groups[k]
		s := summarize(g.deltas, g.halves, g.resolved)
		if *asJSON {
			if err := enc.Encode(struct {
				Comparison string `json:"comparison"`
				summary
			}{k, s}); err != nil {
				fmt.Fprintln(os.Stderr, "summarize:", err)
				os.Exit(1)
			}
			continue
		}
		interval := "—"
		if !math.IsNaN(s.Low) {
			interval = fmt.Sprintf("[%+.1f, %+.1f]%%", s.Low*100, s.High*100)
		}
		fmt.Printf("| %s | %d | %+.1f%% | %s | %.1f pts | %.1f pts | %.1f | %d/%d |\n",
			k, s.Procs, s.Median*100, interval, s.SD*100, s.InRun*100, s.Ratio, s.Resolved, s.Procs)
	}
}

// read adds every result row of one JSON lines file to its comparison's group.
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
		lo, _ := row["low"].(float64)
		hi, _ := row["high"].(float64)
		res, _ := row["resolved"].(bool)
		k := groupKey(row)
		g := groups[k]
		if g == nil {
			g = &group{key: k}
			groups[k] = g
			*order = append(*order, k)
		}
		g.deltas = append(g.deltas, d)
		g.halves = append(g.halves, (hi-lo)/2)
		g.resolved = append(g.resolved, res)
	}
	return sc.Err()
}
