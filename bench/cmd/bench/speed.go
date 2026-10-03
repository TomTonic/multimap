package main

import (
	"fmt"
	"hash/fnv"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/TomTonic/multimap"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/rtopt"
	"github.com/TomTonic/rtcompare"
	"github.com/TomTonic/rtcompare/multiproc"
	"github.com/TomTonic/rtcompare/workload"
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
	Process    int      `json:"process"`
	Seed       uint64   `json:"seed"`
	NsA        float64  `json:"ns_a"`
	NsB        float64  `json:"ns_b"`
	Delta      float64  `json:"delta"`
	Low        float64  `json:"low"`
	High       float64  `json:"high"`
	Level      float64  `json:"level"`
	Resolved   bool     `json:"resolved"`
	Validated  bool     `json:"validated"`
	NoiseFloor float64  `json:"noise_floor"`
	SuspendedS float64  `json:"suspended_s,omitempty"`
	LiveHeap   uint64   `json:"live_heap"`
	Warnings   []string `json:"warnings"`
	// What rtcompare.CombineStaged needs to pool the rows the way multiproc
	// did: the size of the run's first stage, and the A/A differences of each
	// candidate, from which the pooled noise floor and bias come.
	FirstStage int       `json:"first_stage,omitempty"`
	AAA        []float64 `json:"aa_a,omitempty"`
	AAB        []float64 `json:"aa_b,omitempty"`
	// Parallel is how many processes ran at the same time: 1 is the serial
	// regime, more the parallel one, whose rows are never pooled with serial
	// ones (see multiproc.Options.Parallel).
	Parallel     int     `json:"parallel"`
	Quantization float64 `json:"quantization,omitempty"`
	ReportSeed   uint64  `json:"report_seed,omitempty"`
}

// name is how a speed process records the comparison of p; the scenario is
// implied, since every multiproc run covers one scenario.
func (p pair) name() string { return p.op + " " + p.a + " " + p.b }

// pairOf parses a name made by pair.name.
func pairOf(name string) pair {
	f := strings.Fields(name)
	return pair{f[0], f[1], f[2]}
}

// rowOf turns the report of one process into a row of speed.jsonl. first is
// the size of the run's first stage and parallel the number of processes that
// ran at the same time.
func rowOf(kind, profile string, n int, p pair, process int, seed uint64, first, parallel int, r rtcompare.Report) result {
	return result{
		Keys: kind, Values: profile + valueTag, N: n, Op: p.op, A: p.a, B: p.b, Process: process, Seed: seed,
		NsA: r.NsPerOpA, NsB: r.NsPerOpB, Delta: r.Estimate.Delta, Low: r.Estimate.Low, High: r.Estimate.High,
		Level: r.Estimate.Level, Resolved: r.Resolved, Validated: r.Validated, NoiseFloor: r.NoiseFloor,
		SuspendedS: r.Suspended.Seconds(), LiveHeap: r.LiveHeap, Warnings: slices.DeleteFunc(slices.Clone(r.Warnings), multiprocAdvice),
		FirstStage: first, AAA: r.ValidationA.Deltas, AAB: r.ValidationB.Deltas,
		Parallel: parallel, Quantization: r.Quantization, ReportSeed: r.Seed,
	}
}

// multiprocAdvice reports whether w is rtcompare's advice to run a large
// comparison in several processes, which every speed process already follows.
func multiprocAdvice(w string) bool { return strings.Contains(w, "with the multiproc package") }

// reportOf turns a row back into the parts of a report that rtcompare.Combine
// and CombineStaged pool, so that summaries can be pooled again from
// speed.jsonl.
func reportOf(r result) rtcompare.Report {
	return rtcompare.Report{
		NsPerOpA: r.NsA, NsPerOpB: r.NsB, Resolved: r.Resolved, Validated: r.Validated, NoiseFloor: r.NoiseFloor,
		Estimate:    rtcompare.Estimate{Delta: r.Delta, Low: r.Low, High: r.High, Level: r.Level},
		ValidationA: rtcompare.HarnessValidation{Deltas: r.AAA}, ValidationB: rtcompare.HarnessValidation{Deltas: r.AAB},
		Suspended: time.Duration(r.SuspendedS * float64(time.Second)), LiveHeap: r.LiveHeap,
		Quantization: r.Quantization, Seed: r.ReportSeed,
	}
}

// seedFor derives the seed of one comparison of process p from the process's
// seed and the comparison's name, so that a process's resampling is
// reproducible and no two comparisons draw from the same stream.
func seedFor(p *multiproc.Process, name string) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(name))
	return p.Seed ^ h.Sum64() | 1 // zero would ask rtcompare for a random seed
}

// speedSuite is what one speed process of a scenario measures: it builds
// every candidate of the value profile, in the order process p calls for,
// checks that they agree, and compares the pairs ps, recording each report
// under the pair's name. Churn and build come from one workload.Compare per
// pair of candidates (see workload.go). Reports go to stderr.
func speedSuite(kind keys.Kind, profile string, n int, st stream, ps []pair) func(*multiproc.Process) error {
	return func(p *multiproc.Process) error {
		f := newFixture(kind, n, profile, implsFor(profile), st, buildOrder(p))
		if err := f.verify(); err != nil {
			return err
		}
		mutated := map[string]bool{}
		for _, pr := range ps {
			if pr.op != "churn" && pr.op != "build" {
				a, b := f.candidate(pr.op, pr.a), f.candidate(pr.op, pr.b)
				rep, err := rtcompare.Compare(a, b, rtopt.Options(a, b, false, seedFor(p, pr.name())))
				if err != nil {
					return err
				}
				fmt.Fprintf(os.Stderr, "== %s %s n=%d %s: %s vs %s\n%s\n\n", kind, profile, n, pr.op, pr.a, pr.b, rep)
				p.Record(pr.name(), rep)
				continue
			}
			if mutated[pr.b] {
				continue
			}
			mutated[pr.b] = true
			churn, build := pair{"churn", pr.a, pr.b}, pair{"build", pr.a, pr.b}
			wantChurn, wantBuild := slices.Contains(ps, churn), slices.Contains(ps, build)
			if wantBuild {
				if err := f.verifyBuild(pr.a, pr.b); err != nil {
					return err
				}
			}
			seed := seedFor(p, "churn and build "+pr.b)
			res, err := workload.Compare(len(f.vals), f.structure(pr.a), f.structure(pr.b), workload.Options{
				Config: workloadConfig(st), SteadyState: rtopt.Plain(seed), Build: rtopt.Build(seed + 2), SkipBuild: !wantBuild,
			})
			if err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "== %s %s n=%d churn and build: %s vs %s\n%s\n\n", kind, profile, n, pr.a, pr.b, res)
			if wantChurn {
				p.Record(churn.name(), res.SteadyState)
			}
			if wantBuild {
				p.Record(build.name(), res.Build)
			}
		}
		fmt.Fprintf(os.Stderr, "checksum %d\n", sink) // the batches' results stay observable
		return nil
	}
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
		return checkWeigh(v)
	case orderedLP:
		return sum(f.lp.ValuesForSeq(k))
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
	case orderedLP:
		return sum(f.lp.ValuesBetweenInclusiveSeq(from.B[i], to.B[i]))
	}
	return baseKit.rangeSum(f.base, from.B[i], to.B[i])
}

// verifyBuild makes sure the build stream leaves both candidates with
// exactly the corpus: as many keys, and the same values for every key.
func (f *fixture) verifyBuild(impls ...string) error {
	ops, err := workload.Build(len(f.vals), workloadConfig(f.stream))
	if err != nil {
		return err
	}
	n := len(f.c.Keys.B)
	for _, impl := range impls {
		s := f.structure(impl)
		m := s.New()
		s.Apply(m, ops)
		keys, got := f.inspect(impl, m)
		if keys != n {
			return fmt.Errorf("build stream leaves %s with %d keys, want %d", impl, keys, n)
		}
		for i := range n {
			var want uint64
			for _, v := range f.vals[f.offs[i]:f.offs[i+1]] {
				want += checkWeigh(v)
			}
			if got(i) != want {
				return fmt.Errorf("build stream leaves %s with wrong values for %q", impl, f.c.Keys.S[i])
			}
		}
	}
	return nil
}

// inspect returns the number of keys of m, a multimap of candidate impl, and
// a function that sums the values of corpus key i in it.
func (f *fixture) inspect(impl string, m any) (int, func(i int) uint64) {
	kb, ks := f.c.Keys.B, f.c.Keys.S
	switch impl {
	case ordered:
		o := m.(*multimap.Ordered[V])
		return int(o.NumberOfKeys()), func(i int) uint64 { return sum(o.ValuesForSeq(kb[i])) }
	case hashed:
		h := m.(*multimap.Hashed[V])
		return int(h.NumberOfKeys()), func(i int) uint64 { return sum(h.ValuesForSeq(kb[i])) }
	case btreeSets:
		b := m.(*btreeMM)
		return b.Len(), func(i int) uint64 { return btreeSum(b, ks[i]) }
	case mapSets:
		g := m.(mapMM)
		return len(g), func(i int) uint64 { return mapSum(g, ks[i]) }
	case btreeMapC:
		b := m.(*btreeMap)
		return b.Len(), func(i int) uint64 { v, _ := b.Get(ks[i]); return checkWeigh(v) }
	case orderedLP:
		l := m.(*lpMap)
		return l.NumberOfKeys(), func(i int) uint64 { return sum(l.ValuesForSeq(kb[i])) }
	}
	return baseKit.keys(m), func(i int) uint64 { return baseKit.sum(m, kb[i]) }
}

func sum(s func(func(V) bool)) uint64 {
	var a uint64
	for v := range s {
		a += checkWeigh(v)
	}
	return a
}

func btreeSum(m *btreeMM, k string) uint64 {
	var a uint64
	if s, ok := m.Get(k); ok {
		for v := range s {
			a += checkWeigh(v)
		}
	}
	return a
}

func mapSum(m mapMM, k string) uint64 {
	var a uint64
	for v := range m[k] {
		a += checkWeigh(v)
	}
	return a
}

func btreeRangeSum(m *btreeMM, from, to string) uint64 {
	var a uint64
	m.Ascend(from, func(k string, s map[V]struct{}) bool {
		if k > to {
			return false
		}
		for v := range s {
			a += checkWeigh(v)
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
				a += checkWeigh(v)
			}
		}
	}
	return a
}

func btreeMapRangeSum(m *btreeMap, from, to string) uint64 {
	var a uint64
	m.Ascend(from, func(k string, v V) bool {
		if k > to {
			return false
		}
		a += checkWeigh(v)
		return true
	})
	return a
}
