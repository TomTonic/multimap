package main

import (
	"fmt"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/art"
	"github.com/TomTonic/rtcompare/workload"
)

// TestLifeProbe shows what the writes of the benchmark do with the objects of the tree, which neither the speed figures nor
// the CPU profiles say: how many objects a write makes and drops, of which type and size, and how a free list of dropped
// objects (LIFO, one stack for every type and size) would serve the objects made, and how soon (docs/redesign/
// write-path-analysis.md). It only runs if MKLIFE is set, because it is no test but a measurement.
//
// MKLIFE_KEYS, MKLIFE_VALUES and MKLIFE_N select the cases as MKPROBE_* do for TestProbe. For every case it builds the tree
// as the benchmark does, replays the steady-state cycle of workload.Cycle once to reach the steady state, then once more
// write by write, taking the set of the tree's objects before and after every write; the same for the build stream from
// empty. Objects that the census does not report (the Set3 of a value overflow) and temporaries are not seen there; the
// allocations of the runtime a write, with them, come from runtime.MemStats over a plain replay.
//
// The expectation: it prints, per case and stream, the writes by what they did to objects, the objects made and dropped
// a write, the hit rate of the free list and the distances from drop to reuse, and the runtime's allocations a write.
func TestLifeProbe(t *testing.T) {
	if os.Getenv("MKLIFE") == "" {
		t.Skip("set MKLIFE=1 to run the life probe")
	}
	for _, k := range strings.Split(envOr("MKLIFE_KEYS", "street"), ",") {
		for _, v := range strings.Split(envOr("MKLIFE_VALUES", "single-value,natural"), ",") {
			for _, ns := range strings.Split(envOr("MKLIFE_N", "4096"), ",") {
				n, err := strconv.Atoi(ns)
				must(err)
				lifeCase(t, keys.Kind(k), v, n)
			}
		}
	}
}

func lifeCase(t *testing.T, kind keys.Kind, profile string, n int) {
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
	fmt.Printf("\n## life %s %s n=%d (%d writes in the cycle, %d in the build)\n\n", kind, profile, n, len(cycle), len(build))

	var tm art.Map[V]
	for i := range f.c.Keys.B {
		for _, v := range f.vals[f.offs[i]:f.offs[i+1]] {
			tm.Add(f.c.Keys.B[i], v)
		}
	}
	apply(&tm, f, cycle) // reach the steady state
	art.ResetEvents()
	trace(&tm, f, cycle).print("churn (one steady-state cycle)")
	if art.EventsEnabled {
		printEvents("one steady-state cycle", len(cycle))
	}
	mallocs("churn", len(cycle), func() { apply(&tm, f, cycle) })

	var bm art.Map[V]
	art.ResetEvents()
	trace(&bm, f, build).print("build stream from empty")
	if art.EventsEnabled {
		printEvents("build stream", len(build))
	}
	mallocs("build", len(build), func() {
		var m art.Map[V]
		apply(&m, f, build)
	})
}

// mallocs prints the allocations of the runtime a write over one run of fn of the given number of writes.
func mallocs(what string, writes int, fn func()) {
	var a, b runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&a)
	fn()
	runtime.ReadMemStats(&b)
	fmt.Printf("runtime, %s: %.3f allocations and %.1f bytes a write\n\n", what,
		float64(b.Mallocs-a.Mallocs)/float64(writes), float64(b.TotalAlloc-a.TotalAlloc)/float64(writes))
}

type objKey struct {
	label string
	size  int
}

type objInfo struct {
	key   objKey
	block int
}

// freed is an object on a stack of the simulated free list: when it was dropped, and how many bytes had been made by then.
type freed struct {
	write int
	made  int
}

type lifeStats struct {
	writes                   int
	inPlace, one2one, others int // writes that made and dropped nothing, one object each, anything else
	made, dropped            int
	madeBytes, droppedBytes  int
	byKey                    map[objKey]*[3]int // made, dropped, taken from the free list
	hits                     int
	distW                    []int // writes from drop to reuse
	distB                    []int // bytes made from drop to reuse
	held, maxHeld            int   // bytes on the stacks, now and at most
}

// snapshot puts the objects of m into dst.
func snapshot(m *art.Map[V], dst map[uintptr]objInfo) {
	clear(dst)
	m.Objects(func(o art.Object) {
		blk, _ := art.Block(o.Size, o.Pointers)
		dst[o.Addr] = objInfo{objKey{o.Label, o.Size}, blk}
	})
}

// trace replays ops on m write by write and records what each write made and dropped, and a LIFO free list over it: an
// object made is taken from the stack of its key if an earlier write dropped one.
func trace(m *art.Map[V], f *fixture, ops []workload.Op) *lifeStats {
	s := &lifeStats{writes: len(ops), byKey: map[objKey]*[3]int{}}
	stacks := map[objKey][]freed{}
	prev, cur := map[uintptr]objInfo{}, map[uintptr]objInfo{}
	snapshot(m, prev)
	p, kb := f.pairs, f.ck.B
	row := func(k objKey) *[3]int {
		r := s.byKey[k]
		if r == nil {
			r = &[3]int{}
			s.byKey[k] = r
		}
		return r
	}
	for i, op := range ops {
		if op.Kind == workload.Insert {
			m.Add(kb[p.key[op.ID]], p.val[op.ID])
		} else {
			m.Remove(kb[p.key[op.ID]], p.val[op.ID])
		}
		snapshot(m, cur)
		made, dropped := 0, 0
		for a, o := range cur {
			if _, ok := prev[a]; ok {
				continue
			}
			made++
			r := row(o.key)
			r[0]++
			if st := stacks[o.key]; len(st) > 0 {
				top := st[len(st)-1]
				stacks[o.key] = st[:len(st)-1]
				r[2]++
				s.hits++
				s.held -= o.block
				s.distW = append(s.distW, i-top.write)
				s.distB = append(s.distB, s.madeBytes-top.made)
			}
			s.madeBytes += o.block
		}
		for a, o := range prev {
			if _, ok := cur[a]; ok {
				continue
			}
			dropped++
			row(o.key)[1]++
			s.droppedBytes += o.block
			stacks[o.key] = append(stacks[o.key], freed{i, s.madeBytes})
			s.held += o.block
		}
		s.maxHeld = max(s.maxHeld, s.held)
		s.made += made
		s.dropped += dropped
		switch {
		case made == 0 && dropped == 0:
			s.inPlace++
		case made == 1 && dropped == 1:
			s.one2one++
		default:
			s.others++
		}
		prev, cur = cur, prev
	}
	return s
}

func pct(a, b int) float64 { return 100 * float64(a) / float64(max(b, 1)) }

// quantile returns the q-quantile of v (sorted in place), or -1 for an empty v.
func quantile(v []int, q float64) int {
	if len(v) == 0 {
		return -1
	}
	slices.Sort(v)
	return v[min(len(v)-1, int(q*float64(len(v))))]
}

func (s *lifeStats) print(what string) {
	w := float64(s.writes)
	fmt.Printf("%s: %d writes; in place %.1f %%, one object for one %.1f %%, other %.1f %%\n", what, s.writes,
		pct(s.inPlace, s.writes), pct(s.one2one, s.writes), pct(s.others, s.writes))
	fmt.Printf("  objects made %.3f a write (%.1f bytes), dropped %.3f (%.1f bytes)\n",
		float64(s.made)/w, float64(s.madeBytes)/w, float64(s.dropped)/w, float64(s.droppedBytes)/w)
	fmt.Printf("  LIFO free list: %.1f %% of the objects made taken from it; writes from drop to reuse median %d, p90 %d; "+
		"bytes made in between median %d, p90 %d; held at most %d bytes, at the end %d\n",
		pct(s.hits, s.made), quantile(s.distW, 0.5), quantile(s.distW, 0.9), quantile(s.distB, 0.5), quantile(s.distB, 0.9),
		s.maxHeld, s.held)
	type kv struct {
		k objKey
		r [3]int
	}
	var rows []kv
	for k, r := range s.byKey {
		rows = append(rows, kv{k, *r})
	}
	slices.SortFunc(rows, func(a, b kv) int { return b.r[0] + b.r[1] - a.r[0] - a.r[1] })
	fmt.Printf("  %-24s %6s %10s %10s %8s\n", "object", "size", "made/w", "dropped/w", "reused")
	for _, r := range rows[:min(len(rows), 10)] {
		fmt.Printf("  %-24s %6d %10.4f %10.4f %7.1f%%\n", r.k.label, r.k.size, float64(r.r[0])/w, float64(r.r[1])/w, pct(r.r[2], r.r[0]))
	}
	fmt.Println()
}
