package main

import (
	"flag"
	"fmt"
	"maps"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"testing"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/rtcompare"
	"github.com/TomTonic/rtcompare/workload"
)

// TestPairsFor makes sure the benchmark compares exactly what a user choosing
// a multimap wants compared, and skips what cannot be timed sensibly. It
// covers the scenario plan of the benchmark driver: Ordered against every
// other candidate of the value profile in every operation, range queries on
// the scanning candidates and builds only up to their size limits.
func TestPairsFor(t *testing.T) {
	withoutBaseline(t)
	ops := []string{"valuesFor", "valuesBetween", "churn", "build"}
	tests := []struct {
		name    string
		profile string
		n       int
		count   int
		skip    []pair
	}{
		{name: "compares ordered with all three others in every operation", profile: multi, n: 4096, count: 12},
		{
			name: "drops scanning range queries and builds for large scenarios", profile: multi, n: 1 << 20, count: 7,
			skip: []pair{{"valuesBetween", ordered, hashed}, {"valuesBetween", ordered, mapSets}, {"build", ordered, btreeSets}},
		},
		{name: "compares ordered with btree-map for unique values", profile: unique, n: 4096, count: 4},
		{
			name: "keeps range queries on btree-map for large scenarios", profile: unique, n: 1 << 20, count: 3,
			skip: []pair{{"build", ordered, btreeMapC}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pairsFor(keys.U64, tt.n, tt.profile, ops, 1<<16, 1<<16)
			if len(got) != tt.count {
				t.Errorf("%d pairs, want %d: %v", len(got), tt.count, got)
			}
			for _, p := range got {
				if p.a != ordered || !slices.Contains(implsFor(tt.profile)[1:], p.b) {
					t.Errorf("not ordered against another candidate of the profile: %v", p)
				}
			}
			for _, p := range tt.skip {
				if slices.Contains(got, p) {
					t.Errorf("unexpected %v", p)
				}
			}
		})
	}
}

// TestWorkloads makes sure the index workloads behave like a database index
// and never ask a multimap for something impossible. It covers how the
// benchmark driver turns rtcompare's workload streams into insertions and
// deletions of values (see newPairs), for both value profiles: every
// insertion adds a value the key does not hold, every deletion removes one
// it holds, the multimap ends up holding exactly the corpus, and with unique
// values no key ever holds two, up to the largest ratio the driver accepts.
func TestWorkloads(t *testing.T) {
	const n = 3000
	type kv struct {
		key uint32
		val V
	}
	for _, profile := range []string{multi, unique} {
		nums, offs := profileValues(keys.Corpus{}, profile, n)
		vals := toVs(nums)
		corpus := map[kv]bool{}
		for i := range n {
			for _, v := range vals[offs[i]:offs[i+1]] {
				corpus[kv{uint32(i), v}] = true
			}
		}
		for _, r := range []float64{1.5, maxUniqueRatio, 3} {
			if profile == unique && r > maxUniqueRatio {
				continue
			}
			p := newPairs(n, vals, offs, r, profile == unique)
			build, err := workload.Build(len(vals), workloadConfig(r))
			if err != nil {
				t.Fatal(err)
			}
			cycle, err := workload.Cycle(len(vals), workloadConfig(r))
			if err != nil {
				t.Fatal(err)
			}
			for _, tt := range []struct {
				name  string
				ops   []workload.Op
				start map[kv]bool
			}{{"build", build, map[kv]bool{}}, {"churn", cycle, maps.Clone(corpus)}} {
				t.Run(fmt.Sprintf("%s %s with ratio %v", profile, tt.name, r), func(t *testing.T) {
					state, perKey := tt.start, map[uint32]int{}
					for x := range state {
						perKey[x.key]++
					}
					for i, op := range tt.ops {
						x := kv{p.key[op.ID], p.val[op.ID]}
						if del := op.Kind == workload.Delete; del != state[x] {
							t.Fatalf("operation %d (%v): value %v of key %d present=%v", i, op.Kind, x.val, x.key, state[x])
						}
						if op.Kind == workload.Insert {
							state[x] = true
							if perKey[x.key]++; profile == unique && perKey[x.key] > 1 {
								t.Fatalf("operation %d: key %d holds %d values", i, x.key, perKey[x.key])
							}
						} else {
							delete(state, x)
							perKey[x.key]--
						}
					}
					if !maps.Equal(state, corpus) {
						t.Errorf("ends with %d pairs, want the corpus of %d", len(state), len(corpus))
					}
				})
			}
		}
	}
}

// TestResultRows makes sure speed.jsonl keeps what a later summary needs to
// pool the processes again. It covers how the benchmark driver writes a
// process's report as a row and rebuilds it for rtcompare.Combine: the
// comparison's name survives, and the pooled result of the rebuilt reports is
// the one of the originals.
func TestResultRows(t *testing.T) {
	p := pair{"valuesBetween", ordered, btreeMapC}
	if got := pairOf(p.name()); got != p {
		t.Fatalf("pairOf(%q) = %v", p.name(), got)
	}
	var orig, back []rtcompare.Report
	for i, d := range []float64{0.10, 0.12, 0.08, 0.11} {
		r := rtcompare.Report{NsPerOpA: 90, NsPerOpB: 100, Resolved: true, Validated: true, NoiseFloor: 0.01,
			Estimate: rtcompare.Estimate{Delta: d, Low: d - 0.01, High: d + 0.01, Level: 0.95}}
		orig = append(orig, r)
		back = append(back, reportOf(rowOf("u64", multi, 4096, p, i+1, uint64(i), r)))
	}
	want, err := rtcompare.Combine(orig, 0)
	if err != nil {
		t.Fatal(err)
	}
	got, err := rtcompare.Combine(back, 0)
	if err != nil || got.Delta != want.Delta || got.Low != want.Low || got.High != want.High || got.Resolved != want.Resolved {
		t.Errorf("rebuilt reports pool to %+v, %v; want %+v", got, err, want)
	}
}

// TestProfileValues makes sure the unique profile really gives every key
// exactly one value, taken from the same values as the multi profile. It
// covers the value profiles of the benchmark driver.
func TestProfileValues(t *testing.T) {
	mv, mo := profileValues(keys.Corpus{}, multi, 1000)
	uv, uo := profileValues(keys.Corpus{}, unique, 1000)
	if len(uv) != 1000 || len(uo) != 1001 {
		t.Fatalf("unique: %d values, %d offsets; want 1000 and 1001", len(uv), len(uo))
	}
	for i := range 1000 {
		if uo[i+1]-uo[i] != 1 || uv[uo[i]] != mv[mo[i]] {
			t.Fatalf("key %d: %d values, first %d; want 1 value, %d", i, uo[i+1]-uo[i], uv[uo[i]], mv[mo[i]])
		}
	}
}

// TestValidate makes sure the driver refuses settings under which a workload
// could not be generated, before any process starts. It covers the flag
// checks of the benchmark driver.
func TestValidate(t *testing.T) {
	for _, tt := range []struct {
		name    string
		c       config
		wantErr bool
	}{
		{"accepts the defaults", config{profiles: []string{multi, unique}, ops: []string{"churn"}, ratio: 2, minProcs: 5, maxProcs: 5}, false},
		{"rejects fewer than three processes", config{profiles: []string{multi}, ratio: 2, minProcs: 2, maxProcs: 5}, true},
		{"rejects fewer processes at most than at least", config{profiles: []string{multi}, ratio: 2, minProcs: 6, maxProcs: 5}, true},
		{"rejects an unknown profile", config{profiles: []string{"few"}, ratio: 2, minProcs: 5, maxProcs: 5}, true},
		{"rejects a ratio below 1", config{profiles: []string{multi}, ratio: 0.5, minProcs: 5, maxProcs: 5}, true},
		{"rejects churn with ratio 1", config{profiles: []string{multi}, ops: []string{"churn"}, ratio: 1, minProcs: 5, maxProcs: 5}, true},
		{"accepts a large ratio for multi", config{profiles: []string{multi}, ratio: 10, minProcs: 5, maxProcs: 5}, false},
		{"rejects a ratio beyond the extra keys for unique", config{profiles: []string{unique}, ratio: maxUniqueRatio + 1, minProcs: 5, maxProcs: 5}, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.validate(); (err != nil) != tt.wantErr {
				t.Fatalf("validate() = %v, want error %v", err, tt.wantErr)
			}
		})
	}
}

// TestSpeedup makes sure the summary states speed the way a reader expects:
// "2.00×" when A is twice as fast, "0.50×" when it is half as fast. It covers
// the conversion of rtcompare's relative difference (1 - timeA/timeB) in the
// benchmark driver's summary table.
func TestSpeedup(t *testing.T) {
	for _, tt := range []struct{ delta, want float64 }{{0.5, 2}, {0, 1}, {-1, 0.5}} {
		if got := speedup(tt.delta); math.Abs(got-tt.want) > 1e-12 {
			t.Errorf("speedup(%v) = %v, want %v", tt.delta, got, tt.want)
		}
	}
}

// TestMedian makes sure the summary tables report the typical process, not
// an outlier. It covers the median helper of the benchmark driver for odd
// and even counts and for no values at all.
func TestMedian(t *testing.T) {
	if got := median([]float64{3, 1, 2}); got != 2 {
		t.Errorf("odd: %v", got)
	}
	if got := median([]float64{4, 1, 3, 2}); got != 2.5 {
		t.Errorf("even: %v", got)
	}
	if got := median(nil); !math.IsNaN(got) {
		t.Errorf("empty: %v", got)
	}
}

// TestReadLines makes sure the summary finds every result of a run, and that
// a run without results reads as none. It covers how the benchmark driver
// reads speed.jsonl: every line is one result, and a missing file holds none.
func TestReadLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "speed.jsonl")
	if rows, err := readLines[result](path); err != nil || rows != nil {
		t.Fatalf("missing file: %v, %v; want no rows and no error", rows, err)
	}
	if err := os.WriteFile(path, []byte(`{"keys":"u64","seed":3}`+"\n"+`{"keys":"str","seed":4}`+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	rows, err := readLines[result](path)
	if err != nil || len(rows) != 2 || rows[1].Keys != "str" || rows[1].Seed != 4 {
		t.Fatalf("got %+v, %v; want the two rows", rows, err)
	}
	if _, err := readLines[result](dir); err == nil {
		t.Error("a directory must not read as results")
	}
}

// TestCPUInfoModel makes sure the results name the processor they were
// measured on under Linux too. It covers how the benchmark driver reads the
// model name from /proc/cpuinfo, and that it gives up quietly without one.
func TestCPUInfoModel(t *testing.T) {
	for _, tt := range []struct{ name, in, want string }{
		{"reads the first model name", "processor\t: 0\nvendor_id\t: AuthenticAMD\nmodel name\t: AMD Ryzen 9 7900 12-Core Processor\nmodel name\t: other\n", "AMD Ryzen 9 7900 12-Core Processor"},
		{"returns nothing without a model name", "processor\t: 0\nCPU part\t: 0xd0c\n", ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := cpuInfoModel(tt.in); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
	if runtime.GOOS == "linux" && cpuName() == "" {
		t.Log("no model name in /proc/cpuinfo on this machine")
	}
}

// withoutBaseline removes the baseline candidate for the rest of the test,
// so that a test of the scenario plan counts the same pairs in a bench built
// with the baseline tag.
func withoutBaseline(t *testing.T) {
	t.Helper()
	kit := baseKit
	baseKit = nil
	t.Cleanup(func() { baseKit = kit })
}

// TestPairsForPrefix makes sure prefix searches are compared only where they
// mean something: on text keys, where a user types the first characters, and
// on the scanning candidates only while a scan of all keys stays affordable.
// It covers the scenario plan of the benchmark driver for the prefix
// operation.
func TestPairsForPrefix(t *testing.T) {
	withoutBaseline(t)
	ops := []string{"prefix"}
	for _, tt := range []struct {
		name  string
		kind  keys.Kind
		n     int
		count int
	}{
		{"skips integer keys", keys.U64, 4096, 0},
		{"compares text keys with all three others", keys.Street, 4096, 3},
		{"drops the scanning candidates for large scenarios", keys.Path, 1 << 20, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairsFor(tt.kind, tt.n, multi, ops, 1<<16, 1<<16); len(got) != tt.count {
				t.Errorf("%d pairs, want %d: %v", len(got), tt.count, got)
			}
		})
	}
}

// TestApplySuite makes sure a suite preset fills in what the user left open
// and never overrides what the user asked for. It covers the -suite presets
// of the benchmark driver: every preset names only existing flags, and a
// flag given on the command line keeps its value.
func TestApplySuite(t *testing.T) {
	for name := range suites {
		t.Run(name, func(t *testing.T) {
			fs := flag.NewFlagSet("bench", flag.ContinueOnError)
			vals := map[string]*string{}
			for k := range suites[name] {
				vals[k] = fs.String(k, "unset", "")
			}
			if err := fs.Parse([]string{"-maxprocs", "7"}); err != nil {
				t.Fatal(err)
			}
			if err := applySuite(fs, name); err != nil {
				t.Fatal(err)
			}
			for k, v := range vals {
				want := suites[name][k]
				if k == "maxprocs" {
					want = "7"
				}
				if *v != want {
					t.Errorf("-%s = %q, want %q", k, *v, want)
				}
			}
		})
	}
	if err := applySuite(flag.NewFlagSet("bench", flag.ContinueOnError), "nightly"); err == nil {
		t.Error("an unknown suite must be an error")
	}
	if err := applySuite(flag.NewFlagSet("bench", flag.ContinueOnError), "dev"); err == nil {
		t.Error("a preset for flags that do not exist must be an error")
	}
}

// TestProfileValuesNatural makes sure street names keep their real
// localities as values under the multi profile, and one value each under
// unique. It covers how the benchmark driver chooses values for a corpus
// with natural values.
func TestProfileValuesNatural(t *testing.T) {
	c := keys.Generate(keys.Street, 1000, 1)
	vals, offs := profileValues(c, multi, 1000)
	for i := range 1000 {
		if !slices.Equal(vals[offs[i]:offs[i+1]], c.Natural[i]) {
			t.Fatalf("key %q: values %v, want its localities %v", c.Keys.S[i], vals[offs[i]:offs[i+1]], c.Natural[i])
		}
	}
	if uv, uo := profileValues(c, unique, 1000); len(uv) != 1000 || uo[1000] != 1000 {
		t.Errorf("unique: %d values for 1000 keys", len(uv))
	}
}
