package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"runtime/pprof"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/internal/art"
)

// TestReadProbe times the reads of a tree built from the corpus (what the valuesFor and the scan
// benchmarks of the gate do): Each of every key in random order, and a scan of all values, and
// writes a CPU profile of both if MKREAD_PROF names a directory. It only runs if MKREAD is set.
//
// MKREAD_KEYS, MKREAD_VALUES and MKREAD_N select the corpus as MKPROBE_* do for TestProbe.
func TestReadProbe(t *testing.T) {
	if os.Getenv("MKREAD") == "" {
		t.Skip("set MKREAD=1 to run the read probe")
	}
	kind := keys.Kind(envOr("MKREAD_KEYS", "street"))
	profile := envOr("MKREAD_VALUES", "single-value")
	n, err := strconv.Atoi(envOr("MKREAD_N", "4096"))
	if err != nil {
		t.Fatal(err)
	}
	f := newFixture(kind, n, profile, []string{ordered}, stream{ratio: 2, permChurn: 0.25}, func([]string) {})
	var m art.Map[V]
	nvals := 0
	for i := range f.c.Keys.B {
		for _, v := range f.vals[f.offs[i]:f.offs[i+1]] {
			m.Add(f.c.Keys.B[i], v)
			nvals++
		}
	}
	order := rand.New(rand.NewPCG(1, 2)).Perm(len(f.c.Keys.B))
	sink := 0
	each := func() {
		for _, i := range order {
			m.Each(f.c.Keys.B[i], func(V) bool { sink++; return true })
		}
	}
	scan := func() {
		m.RangeValues(&art.Bounds{}, func(V) bool { sink++; return true })
	}
	timed := func(fn func(), per int) float64 {
		var best []float64
		for range 9 {
			start := time.Now()
			reps := 0
			for time.Since(start) < 100*time.Millisecond {
				fn()
				reps++
			}
			best = append(best, float64(time.Since(start))/float64(reps*per))
		}
		slices.Sort(best)
		return best[len(best)/2]
	}
	each()
	scan()
	if dir := os.Getenv("MKREAD_PROF"); dir != "" {
		must(os.MkdirAll(dir, 0o755))
		for _, w := range []struct {
			name string
			fn   func()
		}{{"each", each}, {"scan", scan}} {
			file, err := os.Create(fmt.Sprintf("%s/%s-%s-%d-%s.pprof", dir, kind, profile, n, w.name))
			must(err)
			must(pprof.StartCPUProfile(file))
			start := time.Now()
			for time.Since(start) < 3*time.Second {
				w.fn()
			}
			pprof.StopCPUProfile()
			must(file.Close())
		}
	}
	fmt.Printf("reads %s %s n=%d: each %.1f ns a key, scan %.2f ns a value (%d keys, %d values, sink %d)\n",
		kind, profile, n, timed(each, len(order)), timed(scan, nvals), len(order), nvals, sink%2)
}
