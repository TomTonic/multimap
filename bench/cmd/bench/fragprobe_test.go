package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/metrics"
	"strconv"
	"testing"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/rtcompare/workload"
)

// TestFragProbe runs one candidate through a long life of churn, as a database index that runs for weeks,
// and prints after the build and after every MKFRAG_EVERY cycles how much memory the heap holds for the live
// objects and how much of it is lost between them (docs/redesign/review-2026-10.md, E3). Go never moves an
// object: a span of one size class is given back only when every object in it is dead, so a structure whose
// objects change their size class often can leave the heap full of half-used spans. It only runs if MKFRAG
// is set; one candidate per process (MKFRAG_IMPL), so that the heap holds nothing else.
//
// MKFRAG_KEYS, MKFRAG_VALUES, MKFRAG_N, MKFRAG_CYCLES (default 200) and MKFRAG_EVERY (default 20) select the
// corpus and the length of the run.
func TestFragProbe(t *testing.T) {
	if os.Getenv("MKFRAG") == "" {
		t.Skip("set MKFRAG=1 to run the fragmentation probe")
	}
	impl := envOr("MKFRAG_IMPL", ordered)
	kind := keys.Kind(envOr("MKFRAG_KEYS", "street"))
	profile := envOr("MKFRAG_VALUES", "natural")
	n, err := strconv.Atoi(envOr("MKFRAG_N", "65536"))
	must(err)
	cycles, err := strconv.Atoi(envOr("MKFRAG_CYCLES", "200"))
	must(err)
	every, err := strconv.Atoi(envOr("MKFRAG_EVERY", "20"))
	must(err)
	st := stream{ratio: 2, permChurn: 0.25}
	f := newFixture(kind, n, profile, nil, st, func([]string) {})
	cfg := workloadConfig(st)
	build, err := workload.Build(len(f.vals), cfg)
	must(err)
	cycle, err := workload.Cycle(len(f.vals), cfg)
	must(err)
	s := f.structure(impl)
	m := s.New()
	s.Apply(m, build)
	fmt.Printf("frag %s %s %s n=%d: %d values, cycle of %d operations\n", impl, kind, profile, n, len(f.vals), len(cycle))
	fmt.Println("cycle | live objects MB | unused in spans MB | free pages MB | released MB | heap total MB | objects | unused/live")
	report := func(c int) {
		runtime.GC()
		v := heapMetrics()
		fmt.Printf("%d | %.1f | %.1f | %.1f | %.1f | %.1f | %d | %.3f\n", c, v[0]/1e6, v[1]/1e6, v[2]/1e6, v[3]/1e6,
			(v[0]+v[1]+v[2]+v[3])/1e6, int(v[4]), v[1]/v[0])
	}
	report(0)
	for c := 1; c <= cycles; c++ {
		s.Apply(m, cycle)
		if c%every == 0 {
			report(c)
		}
	}
	runtime.KeepAlive(m)
}

// heapMetrics reads the bytes of live heap objects, the bytes unused in the spans that hold them, the free
// pages the heap keeps, the pages it gave back to the system, and the number of live objects.
func heapMetrics() [5]float64 {
	names := []string{
		"/memory/classes/heap/objects:bytes",
		"/memory/classes/heap/unused:bytes",
		"/memory/classes/heap/free:bytes",
		"/memory/classes/heap/released:bytes",
		"/gc/heap/objects:objects",
	}
	samples := make([]metrics.Sample, len(names))
	for i, name := range names {
		samples[i].Name = name
	}
	metrics.Read(samples)
	var v [5]float64
	for i, s := range samples {
		v[i] = float64(s.Value.Uint64())
	}
	return v
}
