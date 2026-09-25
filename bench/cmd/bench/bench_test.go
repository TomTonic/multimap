package main

import (
	"fmt"
	"maps"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"testing"
)

// TestPairsFor makes sure the benchmark compares exactly what a user choosing
// a multimap wants compared, and skips what cannot be timed sensibly. It
// covers the scenario plan of the benchmark driver: Ordered against every
// other candidate of the value profile in every operation, range queries on
// the scanning candidates and builds only up to their size limits.
func TestPairsFor(t *testing.T) {
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
			got := pairsFor(tt.n, tt.profile, ops, 1<<16, 1<<16)
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
// and never ask a multimap for something impossible. It covers the build and
// churn workloads of the benchmark driver for both value profiles: every
// insertion adds a value the key does not hold, every deletion removes a
// value inserted earlier, insertions and deletions interleave, the ratio of
// insertions to final values is as requested, and the multimap ends up
// holding exactly the corpus. With unique values, no key ever holds two
// values, up to the largest ratio the driver accepts.
func TestWorkloads(t *testing.T) {
	const n = 3000
	for _, profile := range []string{multi, unique} {
		vals, offs := profileValues(profile, n)
		corpus := map[kv]bool{}
		for i := range n {
			for _, v := range vals[offs[i]:offs[i+1]] {
				corpus[kv{uint32(i), v}] = true
			}
		}
		u := profile == unique
		for _, r := range []float64{1, 1.5, 2, 3, maxUniqueRatio} {
			t.Run(fmt.Sprintf("%s build with ratio %v", profile, r), func(t *testing.T) {
				checkWorkload(t, buildWorkload(n, vals, offs, r, 1, u), map[kv]bool{}, corpus, r*float64(len(vals)), u)
			})
			t.Run(fmt.Sprintf("%s churn with ratio %v", profile, r), func(t *testing.T) {
				checkWorkload(t, churnWorkload(n, vals, r, 1, u), maps.Clone(corpus), corpus, (r-1)*float64(len(vals)), u)
			})
		}
	}
}

// checkWorkload replays ms on start and checks it against the expectations
// of TestWorkloads; with unique, also that no key ever holds two values.
func checkWorkload(t *testing.T, ms []mutation, start, end map[kv]bool, inserts float64, unique bool) {
	t.Helper()
	state, adds, switches := start, 0, 0
	perKey := map[uint32]int{}
	for p := range start {
		perKey[p.key]++
	}
	for i, m := range ms {
		p := kv{m.key, m.val}
		if m.del != state[p] {
			t.Fatalf("mutation %d: del=%v but present=%v", i, m.del, state[p])
		}
		if m.del {
			delete(state, p)
			perKey[m.key]--
		} else {
			state[p] = true
			perKey[m.key]++
			adds++
			if unique && perKey[m.key] > 1 {
				t.Fatalf("mutation %d: key %d holds %d values", i, m.key, perKey[m.key])
			}
		}
		if i > 0 && m.del != ms[i-1].del {
			switches++
		}
	}
	if !maps.Equal(state, end) {
		t.Errorf("ends with %d pairs, want the corpus of %d", len(state), len(end))
	}
	if math.Abs(float64(adds)-inserts) > 1 {
		t.Errorf("%d insertions, want %.0f", adds, inserts)
	}
	if inserts > float64(len(end)) && switches < adds/100 {
		t.Errorf("only %d switches between inserting and deleting in %d mutations", switches, len(ms))
	}
}

// TestProfileValues makes sure the unique profile really gives every key
// exactly one value, taken from the same values as the multi profile. It
// covers the value profiles of the benchmark driver.
func TestProfileValues(t *testing.T) {
	mv, mo := profileValues(multi, 1000)
	uv, uo := profileValues(unique, 1000)
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
		{"accepts the defaults", config{profiles: []string{multi, unique}, ops: []string{"churn"}, ratio: 2}, false},
		{"rejects an unknown profile", config{profiles: []string{"few"}, ratio: 2}, true},
		{"rejects a ratio below 1", config{profiles: []string{multi}, ratio: 0.5}, true},
		{"rejects churn with ratio 1", config{profiles: []string{multi}, ops: []string{"churn"}, ratio: 1}, true},
		{"accepts a large ratio for multi", config{profiles: []string{multi}, ratio: 10}, false},
		{"rejects a ratio beyond the key pool for unique", config{profiles: []string{unique}, ratio: maxUniqueRatio + 1}, true},
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

// TestScenarioRows makes sure a run continued with -continue picks up
// exactly where the earlier run of a scenario stopped. It covers how the
// benchmark driver finds a scenario's earlier processes in speed.jsonl: only
// rows of the same key kind, value profile, size and planned comparisons
// count, and the number of processes run is the highest layout seed.
func TestScenarioRows(t *testing.T) {
	ps := pairsFor(1<<20, multi, []string{"valuesFor"}, 1<<16, 1<<16)
	row := func(keys, values string, n int, op, b string, seed uint64) result {
		return result{Keys: keys, Values: values, N: n, Op: op, A: ordered, B: b, LayoutSeed: seed}
	}
	prior := []result{
		row("u64", multi, 1<<20, "valuesFor", hashed, 1),
		row("u64", multi, 1<<20, "valuesFor", mapSets, 20),
		row("u64", multi, 1<<20, "churn", hashed, 25),     // not planned
		row("str", multi, 1<<20, "valuesFor", hashed, 30), // other key kind
		row("u64", unique, 1<<20, "valuesFor", btreeMapC, 30),
		row("u64", multi, 4096, "valuesFor", hashed, 30),
	}
	for _, tt := range []struct {
		name     string
		kind     string
		wantRows int
		wantDone int
	}{
		{"counts only the scenario's planned comparisons", "u64", 2, 20},
		{"starts from scratch when the scenario never ran", "none", 0, 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			rows, done := scenarioRows(prior, tt.kind, multi, 1<<20, ps)
			if len(rows) != tt.wantRows || done != tt.wantDone {
				t.Errorf("%d rows, %d processes; want %d and %d", len(rows), done, tt.wantRows, tt.wantDone)
			}
		})
	}
}

// TestReadLines makes sure a continued run finds the results of the run it
// continues, and that a fresh run starts from nothing. It covers how the
// benchmark driver reads speed.jsonl: every line is one result, and a
// missing file holds none.
func TestReadLines(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "speed.jsonl")
	if rows, err := readLines[result](path); err != nil || rows != nil {
		t.Fatalf("missing file: %v, %v; want no rows and no error", rows, err)
	}
	if err := os.WriteFile(path, []byte(`{"keys":"u64","layout_seed":3}`+"\n"+`{"keys":"str","layout_seed":4}`+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	rows, err := readLines[result](path)
	if err != nil || len(rows) != 2 || rows[1].Keys != "str" || rows[1].LayoutSeed != 4 {
		t.Fatalf("got %+v, %v; want the two rows", rows, err)
	}
	if _, err := readLines[result](dir); err == nil {
		t.Error("a directory must not read as results")
	}
}

// TestContinued makes sure run.json still tells when and on what the first
// run happened after a run continued it, and when the continuation ran. It
// covers the run record of the benchmark driver: the continuation is
// appended under "continued", a missing CPU name is filled in, and without
// an earlier record the continuation's own record is written.
func TestContinued(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "run.json")
	info := map[string]any{"start": "later", "cpu": "Some CPU"}
	if got := continued(path, info); got["start"] != "later" {
		t.Errorf("without an earlier record: %v", got)
	}
	if err := os.WriteFile(path, []byte(`{"start":"first","cpu":"","continued":[{"start":"second"}]}`), 0o644); err != nil {
		t.Fatal(err)
	}
	got := continued(path, info)
	list, _ := got["continued"].([]any)
	if got["start"] != "first" || got["cpu"] != "Some CPU" || len(list) != 2 {
		t.Errorf("got %v; want the first run with the CPU filled in and two continuations", got)
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
