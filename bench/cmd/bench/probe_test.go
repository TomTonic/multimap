package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/art"
	"github.com/TomTonic/rtcompare/workload"
)

// TestProbe shows what the benchmark's own churn and build streams do to the tree with
// multi-key pages, which the speed figures alone do not say: how the number of values of a
// key changes along the stream, what each kind of change costs, and which events (promote,
// burst, merge ...) happen how often. It only runs if MKPROBE is set, because it is no test
// but a measurement; see docs/redesign/step4-probe.md for how to run and read it.
//
// For every case (key kind, value profile, size) it builds the tree as the benchmark does,
// replays the steady-state cycle of workload.Cycle twice (the first time to reach the steady
// state), and then, if MKPROBE_PROF names a directory, under the CPU profiler: 40 cycles and at least MKPROBE_SECS seconds (default
// 3), and the build stream from empty for MKPROBE_SECS seconds (a second file, -build.pprof). The events are counted only by a build
// with the tag mkstats (art.EventsEnabled).
//
// The expectation: it prints, per case, the transitions of the value counts with their
// share of the operations and their time per operation, the events per operation of one
// cycle, and the objects of the tree before and after the cycle.
func TestProbe(t *testing.T) {
	if os.Getenv("MKPROBE") == "" {
		t.Skip("set MKPROBE=1 to run the probe")
	}
	kindsF := envOr("MKPROBE_KEYS", "street")
	valuesF := envOr("MKPROBE_VALUES", "natural,single-value")
	sizesF := envOr("MKPROBE_N", "4096")
	for _, ks := range strings.Split(kindsF, ",") {
		kind := keys.Kind(ks)
		if !slices.Contains(keys.Kinds, kind) {
			t.Fatalf("unknown key kind %q", ks)
		}
		for _, profile := range strings.Split(valuesF, ",") {
			for _, ns := range strings.Split(sizesF, ",") {
				n, err := strconv.Atoi(ns)
				if err != nil {
					t.Fatal(err)
				}
				probeCase(t, kind, profile, n)
			}
		}
	}
}

func envOr(name, def string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return def
}

// transition classes of one operation: its kind and the number of values its key had before
// (0, 1, 2 or 3 and more).
const nClass = 8

func classOf(del bool, before int) int {
	c := min(before, 3)
	if del {
		c += 4
	}
	return c
}

func className(c int) string {
	k := [...]string{"insert into", "delete from"}[c/4]
	b := [...]string{"0 values (new key)", "1 value", "2 values", "3+ values"}[c%4]
	return k + " key with " + b
}

type probeRun struct {
	ops      int
	elapsed  time.Duration
	classN   [nClass]int
	classNs  [nClass]time.Duration
	hist     [2][8]int // the number of values of the keys at the start and in the middle of the cycle
	histKeys [2]int
}

func probeCase(t *testing.T, kind keys.Kind, profile string, n int) {
	st := stream{ratio: 2, permChurn: 0.25}
	f := newFixture(kind, n, profile, []string{ordered}, st, func([]string) {})
	cfg := workloadConfig(st)
	cycle, err := workload.Cycle(len(f.vals), cfg)
	if err != nil {
		t.Fatal(err)
	}
	build, err := workload.Build(len(f.vals), cfg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\n## %s %s n=%d (%d corpus values, %d operations in the cycle, %d in the build)\n\n",
		kind, profile, n, len(f.vals), len(cycle), len(build))

	// --- the build stream, from empty
	var m art.Map[V]
	art.ResetEvents()
	start := time.Now()
	apply(&m, f, build)
	fmt.Printf("build stream: %.0f ns per operation\n\n", float64(time.Since(start))/float64(len(build)))
	printEvents("build stream", len(build))
	census("tree after the build stream", &m)

	// --- the tree as the fixture builds it, then two cycles
	var tm art.Map[V]
	counts := make([]int, len(f.ck.B))
	for i := range f.c.Keys.B {
		for _, v := range f.vals[f.offs[i]:f.offs[i+1]] {
			tm.Add(f.c.Keys.B[i], v)
			counts[i]++
		}
	}
	census("tree as built from the corpus", &tm)
	apply(&tm, f, cycle) // reach the steady state
	census("tree after one cycle", &tm)

	art.ResetEvents()
	run := replayTimed(&tm, f, cycle, counts)
	printRun(run)
	printEvents("one steady-state cycle", len(cycle))

	// --- the corpus tree with every second key's values removed value by value
	var hm art.Map[V]
	for i := range f.c.Keys.B {
		for _, v := range f.vals[f.offs[i]:f.offs[i+1]] {
			hm.Add(f.c.Keys.B[i], v)
		}
	}
	start = time.Now()
	removed := 0
	for i := 1; i < len(f.c.Keys.B); i += 2 {
		for _, v := range f.vals[f.offs[i]:f.offs[i+1]] {
			hm.Remove(f.c.Keys.B[i], v)
			removed++
		}
	}
	fmt.Printf("removing every second key value by value: %.0f ns per removal\n\n", float64(time.Since(start))/float64(max(1, removed)))
	census("tree with every second key removed (Remove, value by value)", &hm)

	// --- a plain replay (no per-operation timing) under the profiler, and the build stream from empty
	dir := os.Getenv("MKPROBE_PROF")
	if dir != "" {
		secs, err := strconv.Atoi(envOr("MKPROBE_SECS", "3"))
		must(err)
		dur := time.Duration(secs) * time.Second
		must(os.MkdirAll(dir, 0o755))
		file, err := os.Create(fmt.Sprintf("%s/%s-%s-%d.pprof", dir, kind, profile, n))
		must(err)
		must(pprof.StartCPUProfile(file))
		for start, i := time.Now(), 0; i < cyclesEnv() || time.Since(start) < dur; i++ {
			apply(&tm, f, cycle)
		}
		pprof.StopCPUProfile()
		must(file.Close())
		file, err = os.Create(fmt.Sprintf("%s/%s-%s-%d-build.pprof", dir, kind, profile, n))
		must(err)
		must(pprof.StartCPUProfile(file))
		for start := time.Now(); time.Since(start) < dur; {
			var bm art.Map[V]
			apply(&bm, f, build)
		}
		pprof.StopCPUProfile()
		must(file.Close())
	}
	start = time.Now()
	apply(&tm, f, cycle)
	fmt.Printf("plain replay of a cycle: %.0f ns per operation\n", float64(time.Since(start))/float64(len(cycle)))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// apply replays ops on m through the fixture's pairs.
func apply(m *art.Map[V], f *fixture, ops []workload.Op) {
	p := f.pairs
	kb := f.ck.B
	for _, op := range ops {
		if op.Kind == workload.Insert {
			m.Add(kb[p.key[op.ID]], p.val[op.ID])
		} else {
			m.Remove(kb[p.key[op.ID]], p.val[op.ID])
		}
	}
}

// replayTimed replays ops on m, timing every operation and filing it under the number of
// values its key had before, which counts tracks. The cost of the clock reads is measured
// and taken off.
func replayTimed(m *art.Map[V], f *fixture, ops []workload.Op, counts []int) probeRun {
	p := f.pairs
	kb := f.ck.B
	r := probeRun{ops: len(ops)}
	var overhead time.Duration
	{
		const k = 200000
		for range k {
			a := time.Now()
			overhead += time.Since(a)
		}
		overhead = time.Duration(float64(overhead)/k + 0.5)
	}
	for i, op := range ops {
		if i == 0 || i == len(ops)/2 {
			r.sample(counts, i != 0)
		}
		k := p.key[op.ID]
		before := counts[k]
		c := classOf(op.Kind != workload.Insert, before)
		t0 := time.Now()
		if op.Kind == workload.Insert {
			m.Add(kb[k], p.val[op.ID])
			counts[k]++
		} else {
			m.Remove(kb[k], p.val[op.ID])
			counts[k]--
		}
		d := time.Since(t0) - overhead
		r.classN[c]++
		r.classNs[c] += d
		r.elapsed += d
	}
	return r
}

func (r *probeRun) sample(counts []int, mid bool) {
	i := 0
	if mid {
		i = 1
	}
	for _, c := range counts {
		r.hist[i][min(c, 7)]++
	}
	r.histKeys[i] = len(counts)
}

func printRun(r probeRun) {
	fmt.Printf("per operation, by what the operation does to its key (clock overhead taken off):\n\n")
	fmt.Println("| operation | share of operations | ns per operation |")
	fmt.Println("|---|--:|--:|")
	for c := range nClass {
		if r.classN[c] == 0 {
			continue
		}
		fmt.Printf("| %s | %.1f %% | %.0f |\n", className(c), 100*float64(r.classN[c])/float64(r.ops), float64(r.classNs[c])/float64(r.classN[c]))
	}
	fmt.Printf("| all | 100 %% | %.0f |\n\n", float64(r.elapsed)/float64(r.ops))
	fmt.Println("keys by number of values (of all keys the stream knows, including those without a value):")
	fmt.Println()
	fmt.Println("| when | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7+ |")
	fmt.Println("|---|--:|--:|--:|--:|--:|--:|--:|--:|")
	for i, when := range []string{"start of the cycle", "middle of the cycle"} {
		fmt.Printf("| %s |", when)
		for _, c := range r.hist[i] {
			fmt.Printf(" %.1f %% |", 100*float64(c)/float64(max(1, r.histKeys[i])))
		}
		fmt.Println()
	}
	fmt.Println()
}

func printEvents(what string, ops int) {
	if !art.EventsEnabled {
		return
	}
	ev := art.Events()
	names := make([]string, 0, len(ev))
	for k, c := range ev {
		if c.Total > 0 {
			names = append(names, k)
		}
	}
	sort.Strings(names)
	fmt.Printf("events of %s (%d operations):\n\n| event | count | per 1000 operations | entries of the page: median / 90th percentile / largest |\n|---|--:|--:|---|\n", what, ops)
	for _, k := range names {
		c := ev[k]
		fmt.Printf("| %s | %d | %.1f | %s |\n", k, c.Total, 1000*float64(c.Total)/float64(ops), quantiles(c))
	}
	fmt.Println()
	if os.Getenv("MKPROBE_HIST") != "" { // the whole histograms over the entries, for a closer look
		for _, k := range names {
			fmt.Printf("histogram | %s |", k)
			for n, x := range ev[k].Entries {
				if x > 0 {
					fmt.Printf(" %d:%d", n, x)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func quantiles(c *art.EventCount) string {
	var seen uint64
	var med, p90, largest int
	for n, k := range c.Entries {
		if k == 0 {
			continue
		}
		largest = n
		seen += k
		if med == 0 && seen*2 >= c.Total {
			med = n
		}
		if p90 == 0 && seen*10 >= c.Total*9 {
			p90 = n
		}
	}
	return fmt.Sprintf("%d / %d / %d", med, p90, largest)
}

// census prints how many objects of which kind the tree has, how its keys spread over them,
// and how the multi-key pages fill.
func census(what string, m *art.Map[V]) {
	type row struct{ n, bytes, keys, values int }
	rows := map[string]*row{}
	var pageEntries []int
	var skValues [4]int // single-key pages by number of values: 1, 2, 3-4, 5+
	var bySize [513]row // multi-key pages by the size of the object
	skBySize := map[int]int{} // single-key pages by the size of the object
	total, blocks := 0, 0
	m.Objects(func(o art.Object) {
		r := rows[o.Label]
		if r == nil {
			r = &row{}
			rows[o.Label] = r
		}
		r.n++
		blk, _ := art.Block(o.Size, o.Pointers)
		blocks += blk
		r.bytes += o.Size
		r.keys += o.Keys
		r.values += o.Values
		total += o.Keys
		switch o.Label {
		case "multi-key page":
			pageEntries = append(pageEntries, o.Keys)
			bySize[o.Size].n++
			bySize[o.Size].keys += o.Keys
			bySize[o.Size].values += o.Values
		case "single-key page":
			skValues[valueBucket(o.Values)]++
			skBySize[o.Size]++
		}
	})
	fmt.Printf("objects of the %s:\n\n| object | count | bytes | keys | share of keys |\n|---|--:|--:|--:|--:|\n", what)
	labels := make([]string, 0, len(rows))
	for l := range rows {
		labels = append(labels, l)
	}
	sort.Strings(labels)
	for _, l := range labels {
		r := rows[l]
		fmt.Printf("| %s | %d | %d | %d | %.1f %% |\n", l, r.n, r.bytes, r.keys, 100*float64(r.keys)/float64(max(1, total)))
	}
	fmt.Printf("\nblock bytes per key: %.1f\n\n", float64(blocks)/float64(max(1, total)))
	if len(pageEntries) > 0 {
		slices.Sort(pageEntries)
		sum := 0
		for _, e := range pageEntries {
			sum += e
		}
		fmt.Printf("multi-key pages: %d, entries per page: mean %.1f, median %d, 10th percentile %d, 90th percentile %d, largest %d\n\n",
			len(pageEntries), float64(sum)/float64(len(pageEntries)), pageEntries[len(pageEntries)/2],
			pageEntries[len(pageEntries)/10], pageEntries[len(pageEntries)*9/10], pageEntries[len(pageEntries)-1])
	}
	if len(pageEntries) > 0 {
		fmt.Printf("multi-key pages by size (count, keys and values a page, bytes a value):")
		for size := range bySize {
			if r := bySize[size]; r.n > 0 {
				fmt.Printf(" %d: %d, %.1f, %.1f, %.1f;", size, r.n, float64(r.keys)/float64(r.n), float64(r.values)/float64(r.n), float64(size*r.n)/float64(max(1, r.values)))
			}
		}
		fmt.Print("\n\n")
	}
	fmt.Printf("single-key pages by number of values: 1: %d, 2: %d, 3-4: %d, 5+: %d\n\n", skValues[0], skValues[1], skValues[2], skValues[3])
	sizes := make([]int, 0, len(skBySize))
	for size := range skBySize {
		sizes = append(sizes, size)
	}
	slices.Sort(sizes)
	fmt.Printf("single-key pages by size of the object:")
	for _, size := range sizes {
		fmt.Printf(" %d: %d;", size, skBySize[size])
	}
	fmt.Print("\n\n")
}

func valueBucket(n int) int {
	switch {
	case n <= 1:
		return 0
	case n == 2:
		return 1
	case n <= 4:
		return 2
	}
	return 3
}

// cyclesEnv is how many cycles the profiled replay runs (MKPROBE_CYCLES, 40 by default), so that
// the profile has enough samples.
func cyclesEnv() int {
	n, err := strconv.Atoi(envOr("MKPROBE_CYCLES", "40"))
	if err != nil || n < 1 {
		return 40
	}
	return n
}
