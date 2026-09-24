package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/layout"
	"github.com/TomTonic/multimap/bench/rtopt"
	"github.com/TomTonic/rtcompare"
)

// sink receives every batch's result, so the compiler cannot drop the work.
var sink uint64

// pair is one comparison a user would ask for.
type pair struct{ op, a, b string }

// pairsFor lists the comparisons of one scenario. Ordered is compared with
// every alternative, Hashed with the hand-written map it replaces. Range
// queries on Hashed scan every key, so they are compared only up to
// hashedRangeMax keys; building is compared only up to buildMax keys, because
// one build of a large multimap takes too long for a batch.
func pairsFor(n int, ops []string, hashedRangeMax, buildMax int) []pair {
	var ps []pair
	for _, op := range ops {
		switch op {
		case "valuesFor", "addRemove":
			ps = append(ps, pair{op, ordered, hashed}, pair{op, ordered, btreeSets},
				pair{op, ordered, mapSets}, pair{op, hashed, mapSets})
		case "valuesBetween":
			ps = append(ps, pair{op, ordered, btreeSets})
			if n <= hashedRangeMax {
				ps = append(ps, pair{op, ordered, hashed})
			}
		case "build":
			if n <= buildMax {
				ps = append(ps, pair{op, ordered, hashed}, pair{op, ordered, btreeSets},
					pair{op, ordered, mapSets}, pair{op, hashed, mapSets})
			}
		}
	}
	return ps
}

// result is one comparison of one process, as written to the JSON lines file.
type result struct {
	Keys       string   `json:"keys"`
	N          int      `json:"n"`
	Op         string   `json:"op"`
	A          string   `json:"a"`
	B          string   `json:"b"`
	NsA        float64  `json:"ns_a"`
	NsB        float64  `json:"ns_b"`
	Delta      float64  `json:"delta"`
	Low        float64  `json:"low"`
	High       float64  `json:"high"`
	Resolved   bool     `json:"resolved"`
	NoiseFloor float64  `json:"noise_floor"`
	InnerLoops uint64   `json:"inner_loops"`
	LayoutSeed uint64   `json:"layout_seed"`
	Warnings   []string `json:"warnings"`
}

// runSpeed is one child process: it builds every candidate, checks that they
// agree, runs the comparisons of the scenario and writes one JSON line per
// comparison to out. Reports go to stderr.
func runSpeed(kind keys.Kind, n int, ps []pair, out io.Writer) error {
	f := newFixture(kind, n, allImpls)
	if err := f.verify(); err != nil {
		return err
	}
	enc := json.NewEncoder(out)
	for _, p := range ps {
		if p.op == "addRemove" {
			f.warmAddRemove()
		}
		a, b := f.candidate(p.op, p.a), f.candidate(p.op, p.b)
		fmt.Fprintf(os.Stderr, "== %s n=%d %s: %s vs %s\n", kind, n, p.op, p.a, p.b)
		rep, err := rtcompare.Compare(a, b, rtopt.Options(a, b, p.op == "build"))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "%s\n\n", rep)
		if err := enc.Encode(result{
			Keys: string(kind), N: n, Op: p.op, A: p.a, B: p.b, NsA: rep.NsPerOpA, NsB: rep.NsPerOpB,
			Delta: rep.Estimate.Delta, Low: rep.Estimate.Low, High: rep.Estimate.High,
			Resolved: rep.Resolved, NoiseFloor: rep.NoiseFloor, InnerLoops: rep.ValidationA.InnerLoops,
			LayoutSeed: layout.Seed(), Warnings: rep.Warnings,
		}); err != nil {
			return err
		}
	}
	fmt.Fprintf(os.Stderr, "checksum %d\n", sink) // the batches' results stay observable
	return nil
}

// verify makes sure all candidates answer identically before anything is
// timed, so that a fast wrong answer cannot win.
func (f *fixture) verify() error {
	n := len(f.c.Hits.B)
	for i := range min(n, 20000) {
		want := sum(f.ord.ValuesForSeq(f.c.Hits.B[i]))
		got := []uint64{sum(f.hsh.ValuesForSeq(f.c.Hits.B[i])), btreeSum(f.bt, f.c.Hits.S[i]), mapSum(f.gm, f.c.Hits.S[i])}
		if slices.ContainsFunc(got, func(g uint64) bool { return g != want }) || want == 0 {
			return fmt.Errorf("candidates disagree on the values of %q", f.c.Hits.S[i])
		}
	}
	for i := range min(len(f.from.B), 2000) {
		want := sum(f.ord.ValuesBetweenInclusiveSeq(f.from.B[i], f.to.B[i]))
		if got := btreeRangeSum(f.bt, f.from.S[i], f.to.S[i]); got != want {
			return fmt.Errorf("btree-sets disagrees on range %q..%q", f.from.S[i], f.to.S[i])
		}
		if i < 20 {
			if got := sum(f.hsh.ValuesBetweenInclusiveSeq(f.from.B[i], f.to.B[i])); got != want {
				return fmt.Errorf("hashed disagrees on range %q..%q", f.from.S[i], f.to.S[i])
			}
		}
	}
	return nil
}

// warmAddRemove adds and removes one absent value at every key once, so that
// every value set has reached the representation the timed operations
// alternate within.
func (f *fixture) warmAddRemove() {
	for i := range f.c.Hits.B {
		v := absent(i)
		f.ord.AddValue(f.c.Hits.B[i], v)
		f.ord.RemoveValue(f.c.Hits.B[i], v)
		f.hsh.AddValue(f.c.Hits.B[i], v)
		f.hsh.RemoveValue(f.c.Hits.B[i], v)
		btreeAdd(f.bt, f.c.Hits.S[i], v)
		btreeRemove(f.bt, f.c.Hits.S[i], v)
		mapAdd(f.gm, f.c.Hits.S[i], v)
		mapRemove(f.gm, f.c.Hits.S[i], v)
	}
}

// absent returns a value no key holds: the corpus values have the top bit clear.
func absent(i int) uint64 { return 1<<63 | uint64(i) }

func sum(s func(func(uint64) bool)) uint64 {
	var a uint64
	for v := range s {
		a += v
	}
	return a
}

func btreeSum(m *btreeMM, k string) uint64 {
	var a uint64
	if s, ok := m.Get(k); ok {
		for v := range s {
			a += v
		}
	}
	return a
}

func mapSum(m mapMM, k string) uint64 {
	var a uint64
	for v := range m[k] {
		a += v
	}
	return a
}

func btreeRangeSum(m *btreeMM, from, to string) uint64 {
	var a uint64
	m.Ascend(from, func(k string, s map[uint64]struct{}) bool {
		if k > to {
			return false
		}
		for v := range s {
			a += v
		}
		return true
	})
	return a
}
