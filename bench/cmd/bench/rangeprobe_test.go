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

// TestRangeProbe times the ranges of the bench's valuesBetween (rangeKeys consecutive keys) for the ordered map and
// btree-map, and writes a CPU profile of the ordered map's ranges if MKRANGE_PROF names a directory: the diagnosis of
// the single-value ranges against btree-map (docs/redesign/review-2026-10.md). It only runs if MKRANGE is set.
//
// MKRANGE_KEYS, MKRANGE_VALUES and MKRANGE_N select the corpus as MKPROBE_* do for TestProbe.
func TestRangeProbe(t *testing.T) {
	if os.Getenv("MKRANGE") == "" {
		t.Skip("set MKRANGE=1 to run the range probe")
	}
	kind := keys.Kind(envOr("MKRANGE_KEYS", "street"))
	profile := envOr("MKRANGE_VALUES", "single-value")
	n, err := strconv.Atoi(envOr("MKRANGE_N", "4096"))
	must(err)
	f := newFixture(kind, n, profile, []string{ordered, btreeMapC}, stream{ratio: 2, permChurn: 0.25}, func([]string) {})
	from, to := f.from, f.to
	values := 0
	for j := range from.B {
		for range f.ord.ValuesBetweenInclusiveSeq(from.B[j], to.B[j]) {
			values++
		}
	}
	ord := func() {
		var acc uint64
		for j := range from.B {
			for v := range f.ord.ValuesBetweenInclusiveSeq(from.B[j], to.B[j]) {
				acc += weigh(v)
			}
		}
		sink += acc
	}
	total := 0
	for range f.ord.AllValuesSeq() {
		total++
	}
	all := func() {
		var acc uint64
		for v := range f.ord.AllValuesSeq() {
			acc += weigh(v)
		}
		sink += acc
	}
	bm := func() {
		var acc uint64
		for j := range from.B {
			acc += btreeMapRangeSum(f.bm, from.S[j], to.S[j])
		}
		sink += acc
	}
	timed := func(fn func()) float64 {
		var runs []float64
		for range 9 {
			start := time.Now()
			reps := 0
			for time.Since(start) < 100*time.Millisecond {
				fn()
				reps++
			}
			runs = append(runs, float64(time.Since(start))/float64(reps*values))
		}
		slices.Sort(runs)
		return runs[len(runs)/2]
	}
	if dir := os.Getenv("MKRANGE_PROF"); dir != "" {
		must(os.MkdirAll(dir, 0o755))
		file, err := os.Create(fmt.Sprintf("%s/%s-%s-%d-range.pprof", dir, kind, profile, n))
		must(err)
		must(pprof.StartCPUProfile(file))
		for start := time.Now(); time.Since(start) < 3*time.Second; {
			ord()
		}
		pprof.StopCPUProfile()
		must(file.Close())
	}
	fmt.Printf("ranges %s %s n=%d: ordered %.2f ns a value, btree-map %.2f ns a value (%d ranges, %d values)\n",
		kind, profile, n, timed(ord), timed(bm), len(from.B), values)
	values = total
	fmt.Printf("full scan %s %s n=%d: ordered %.2f ns a value (AllValuesSeq, %d values)\n", kind, profile, n, timed(all), total)
}
