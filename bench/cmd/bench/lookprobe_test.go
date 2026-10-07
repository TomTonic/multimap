package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/TomTonic/multimap/bench/keys"
)

// TestLookProbe times the bench's valuesFor (all values of random existing keys) of the ordered map and of the
// baseline candidate (build tag baseline, made with cmd/mkbaseline from the version to compare), and writes a CPU
// profile of each if MKLOOK_PROF names a directory: the diagnosis of the lookups of small maps against main
// (docs/redesign/review-2026-10.md). It only runs if MKLOOK is set.
//
// MKLOOK_KEYS, MKLOOK_VALUES and MKLOOK_N select the corpus as MKPROBE_* do for TestProbe.
func TestLookProbe(t *testing.T) {
	if os.Getenv("MKLOOK") == "" {
		t.Skip("set MKLOOK=1 to run the lookup probe")
	}
	if baseKit == nil {
		t.Skip("the lookup probe needs the build tag baseline")
	}
	kind := keys.Kind(envOr("MKLOOK_KEYS", "u64"))
	profile := envOr("MKLOOK_VALUES", "single-value")
	n, err := strconv.Atoi(envOr("MKLOOK_N", "4096"))
	must(err)
	f := newFixture(kind, n, profile, []string{ordered, baseline}, stream{ratio: 2, permChurn: 0.25}, func([]string) {})
	const batch = 1 << 14
	timed := func(fn func(uint64)) float64 {
		var runs []float64
		for range 9 {
			start := time.Now()
			reps := 0
			for time.Since(start) < 100*time.Millisecond {
				fn(batch)
				reps++
			}
			runs = append(runs, float64(time.Since(start))/float64(reps*batch))
		}
		slices.Sort(runs)
		return runs[len(runs)/2]
	}
	ord, base := f.valuesFor(ordered), f.valuesFor(baseline)
	if dir := os.Getenv("MKLOOK_PROF"); dir != "" {
		must(os.MkdirAll(dir, 0o755))
		for _, w := range []struct {
			name string
			fn   func(uint64)
		}{{"ordered", ord}, {"baseline", base}} {
			file, err := os.Create(fmt.Sprintf("%s/%s-%s-%d-%s.pprof", dir, kind, profile, n, w.name))
			must(err)
			must(pprof.StartCPUProfile(file))
			for start := time.Now(); time.Since(start) < 3*time.Second; {
				w.fn(batch)
			}
			pprof.StopCPUProfile()
			must(file.Close())
		}
	}
	fmt.Printf("lookups %s %s n=%d: ordered %.1f ns, baseline (%s) %.1f ns\n", kind, profile, n, timed(ord), baseKit.ref, timed(base))
}
