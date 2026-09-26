package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/layout"
	"github.com/TomTonic/multimap/bench/rtopt"
	"github.com/TomTonic/rtcompare"
)

// sink receives every batch's result, so the compiler cannot drop the work.
var sink uint64

// pair is one comparison a user would ask for.
type pair struct{ op, a, b string }

// pairsFor lists the comparisons of one scenario: Ordered against every
// other candidate of the value profile. Range and prefix queries on hashed
// and map-sets scan every key, so they are compared only up to scanMax keys;
// building is compared only up to buildMax keys. Both take too long per
// operation beyond that for the batch size the fast Ordered side needs.
// Prefix queries need text keys (see keys.Text).
func pairsFor(kind keys.Kind, n int, profile string, ops []string, scanMax, buildMax int) []pair {
	var ps []pair
	for _, op := range ops {
		for _, b := range implsFor(profile)[1:] {
			ranged := op == "valuesBetween" || op == "prefix"
			switch {
			case op == "prefix" && !keys.Text(kind):
			case ranged && (b == hashed || b == mapSets) && n > scanMax:
			case op == "build" && n > buildMax:
			default:
				ps = append(ps, pair{op, ordered, b})
			}
		}
	}
	return ps
}

// result is one comparison of one process, as written to the JSON lines file.
type result struct {
	Keys       string   `json:"keys"`
	Values     string   `json:"values"`
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
func runSpeed(kind keys.Kind, profile string, n int, ratio float64, ps []pair, out io.Writer) error {
	f := newFixture(kind, n, profile, implsFor(profile))
	f.ratio = ratio
	if err := f.verify(); err != nil {
		return err
	}
	enc := json.NewEncoder(out)
	for i, p := range ps {
		if p.op == "build" {
			if err := f.verifyBuild(p.a, p.b); err != nil {
				return err
			}
		}
		a, b := f.candidate(p.op, p.a), f.candidate(p.op, p.b)
		fmt.Fprintf(os.Stderr, "== %s %s n=%d %s: %s vs %s\n", kind, profile, n, p.op, p.a, p.b)
		rep, err := rtcompare.Compare(a, b, rtopt.Options(a, b, p.op == "build"))
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stderr, "%s\n\n", rep)
		if err := enc.Encode(result{
			Keys: string(kind), Values: profile, N: n, Op: p.op, A: p.a, B: p.b, NsA: rep.NsPerOpA, NsB: rep.NsPerOpB,
			Delta: rep.Estimate.Delta, Low: rep.Estimate.Low, High: rep.Estimate.High,
			Resolved: rep.Resolved, NoiseFloor: rep.NoiseFloor, InnerLoops: rep.ValidationA.InnerLoops,
			LayoutSeed: layout.Seed(), Warnings: rep.Warnings,
		}); err != nil {
			return err
		}
		if p.op == "churn" && (i+1 == len(ps) || ps[i+1].op != "churn") {
			f.settle()
			if err := f.verify(); err != nil {
				return fmt.Errorf("after churn: %w", err)
			}
		}
	}
	fmt.Fprintf(os.Stderr, "checksum %d\n", sink) // the batches' results stay observable
	return nil
}

// verify makes sure all candidates answer identically before anything is
// timed, so that a fast wrong answer cannot win: point queries, ranges and
// prefix ranges.
func (f *fixture) verify() error {
	for _, r := range [][2]keys.Set{{f.from, f.to}, {f.pfrom, f.pto}} {
		if err := f.verifyRanges(r[0], r[1]); err != nil {
			return err
		}
	}
	for i := range min(len(f.c.Hits.B), 20000) {
		want := sum(f.ord.ValuesForSeq(f.c.Hits.B[i]))
		for _, impl := range f.others {
			if got := f.pointSum(impl, i); got != want || want == 0 {
				return fmt.Errorf("%s disagrees on the values of %q", impl, f.c.Hits.S[i])
			}
		}
	}
	return nil
}

// verifyRanges makes sure all candidates return the same values for the
// ranges [from[i], to[i]]; the scanning candidates are checked on a few
// ranges only, as each check scans every key.
func (f *fixture) verifyRanges(from, to keys.Set) error {
	for i := range min(len(from.B), 2000) {
		want := sum(f.ord.ValuesBetweenInclusiveSeq(from.B[i], to.B[i]))
		for _, impl := range f.others {
			if (impl == hashed || impl == mapSets) && i >= 20 {
				continue
			}
			if f.rangeSum(impl, from, to, i) != want {
				return fmt.Errorf("%s disagrees on range %q..%q", impl, from.S[i], to.S[i])
			}
		}
	}
	return nil
}

// pointSum returns the sum of the values of hit i in candidate impl, which
// is not ordered.
func (f *fixture) pointSum(impl string, i int) uint64 {
	k, s := f.c.Hits.B[i], f.c.Hits.S[i]
	switch impl {
	case hashed:
		return sum(f.hsh.ValuesForSeq(k))
	case btreeSets:
		return btreeSum(f.bt, s)
	case mapSets:
		return mapSum(f.gm, s)
	case btreeMapC:
		v, _ := f.bm.Get(s)
		return v
	}
	return baseKit.sum(f.base, k)
}

// rangeSum returns the sum of the values of the keys in [from[i], to[i]] in
// candidate impl, which is not ordered.
func (f *fixture) rangeSum(impl string, from, to keys.Set, i int) uint64 {
	switch impl {
	case hashed:
		return sum(f.hsh.ValuesBetweenInclusiveSeq(from.B[i], to.B[i]))
	case btreeSets:
		return btreeRangeSum(f.bt, from.S[i], to.S[i])
	case mapSets:
		return mapRangeSum(f.gm, from.S[i], to.S[i])
	case btreeMapC:
		return btreeMapRangeSum(f.bm, from.S[i], to.S[i])
	}
	return baseKit.rangeSum(f.base, from.B[i], to.B[i])
}

// verifyBuild makes sure the build workload leaves both candidates with
// exactly the corpus: as many keys, and the same values for every key.
func (f *fixture) verifyBuild(impls ...string) error {
	ms, n := f.buildWorkload(), len(f.c.Keys.B)
	for _, impl := range impls {
		var keys int
		var got func(i int) uint64
		switch impl {
		case ordered:
			m := multimap.NewOrdered[uint64]()
			applyOrdered(m, f.ck.B, ms)
			keys, got = int(m.NumberOfKeys()), func(i int) uint64 { return sum(m.ValuesForSeq(f.c.Keys.B[i])) }
		case hashed:
			m := multimap.NewHashed[uint64]()
			applyHashed(m, f.ck.B, ms)
			keys, got = int(m.NumberOfKeys()), func(i int) uint64 { return sum(m.ValuesForSeq(f.c.Keys.B[i])) }
		case btreeSets:
			m := &btreeMM{}
			applyBtree(m, f.ck.S, ms)
			keys, got = m.Len(), func(i int) uint64 { return btreeSum(m, f.c.Keys.S[i]) }
		case mapSets:
			m := mapMM{}
			applyMap(m, f.ck.S, ms)
			keys, got = len(m), func(i int) uint64 { return mapSum(m, f.c.Keys.S[i]) }
		case btreeMapC:
			m := &btreeMap{}
			applyBtreeMap(m, f.ck.S, ms)
			keys, got = m.Len(), func(i int) uint64 { v, _ := m.Get(f.c.Keys.S[i]); return v }
		case baseline:
			m := baseKit.empty()
			baseKit.apply(m)(f.ck.B, ms)
			keys, got = baseKit.keys(m), func(i int) uint64 { return baseKit.sum(m, f.c.Keys.B[i]) }
		}
		if keys != n {
			return fmt.Errorf("build workload leaves %s with %d keys, want %d", impl, keys, n)
		}
		for i := range n {
			var want uint64
			for _, v := range f.vals[f.offs[i]:f.offs[i+1]] {
				want += v
			}
			if got(i) != want {
				return fmt.Errorf("build workload leaves %s with wrong values for %q", impl, f.c.Keys.S[i])
			}
		}
	}
	return nil
}

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

func mapRangeSum(m mapMM, from, to string) uint64 {
	var a uint64
	for k, s := range m {
		if k >= from && k <= to {
			for v := range s {
				a += v
			}
		}
	}
	return a
}

func btreeMapRangeSum(m *btreeMap, from, to string) uint64 {
	var a uint64
	m.Ascend(from, func(k string, v uint64) bool {
		if k > to {
			return false
		}
		a += v
		return true
	})
	return a
}
