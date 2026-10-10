package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime/pprof"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/TomTonic/multimap/bench/keys"
)

// TestLookProbe times the bench's valuesFor (all values of random existing keys) of the ordered map and of the
// baseline candidate (build tag baseline, made with cmd/mkbaseline from the version to compare), and writes a CPU
// profile of each if MKLOOK_PROF names a directory: the diagnosis of the lookups of small maps against main
// (docs/redesign/review-2026-10.md) and of the descent (descent-analysis.md). It only runs if MKLOOK is set.
//
// MKLOOK_KEYS, MKLOOK_VALUES and MKLOOK_N select the cases as MKPROBE_* do for TestProbe (lists separated by commas or
// by +, since the M1 queue separates its environment variables by commas); MKLOOK_SECS sets how long each profile runs
// (default 3 seconds; longer for a view of single instructions). Without the build tag baseline only the ordered map is
// timed and profiled. MKLOOK_LINES prints, for each profile of the ordered map, the 80 source lines with the most own
// time (go tool pprof -top -lines), so that a job whose files are not kept (the M1 queue) reports them in its log.
func TestLookProbe(t *testing.T) {
	if os.Getenv("MKLOOK") == "" {
		t.Skip("set MKLOOK=1 to run the lookup probe")
	}
	list := func(s string) []string {
		return strings.FieldsFunc(s, func(r rune) bool { return r == ',' || r == '+' })
	}
	for _, k := range list(envOr("MKLOOK_KEYS", "u64")) {
		for _, v := range list(envOr("MKLOOK_VALUES", "single-value")) {
			for _, ns := range list(envOr("MKLOOK_N", "4096")) {
				n, err := strconv.Atoi(ns)
				must(err)
				lookCase(keys.Kind(k), v, n)
			}
		}
	}
}

func lookCase(kind keys.Kind, profile string, n int) {
	cands := []string{ordered}
	if baseKit != nil {
		cands = append(cands, baseline)
	}
	f := newFixture(kind, n, profile, cands, stream{ratio: 2, permChurn: 0.25}, func([]string) {})
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
	type who struct {
		name string
		fn   func(uint64)
	}
	runs := []who{{"ordered", f.valuesFor(ordered)}}
	if baseKit != nil {
		runs = append(runs, who{"baseline", f.valuesFor(baseline)})
	}
	if dir := os.Getenv("MKLOOK_PROF"); dir != "" {
		secs, err := strconv.Atoi(envOr("MKLOOK_SECS", "3"))
		must(err)
		must(os.MkdirAll(dir, 0o755))
		for _, w := range runs {
			name := fmt.Sprintf("%s/%s-%s-%d-%s.pprof", dir, kind, profile, n, w.name)
			file, err := os.Create(name)
			must(err)
			must(pprof.StartCPUProfile(file))
			for start := time.Now(); time.Since(start) < time.Duration(secs)*time.Second; {
				w.fn(batch)
			}
			pprof.StopCPUProfile()
			must(file.Close())
			if os.Getenv("MKLOOK_LINES") != "" && w.name == "ordered" {
				out, err := exec.Command("go", "tool", "pprof", "-top", "-lines", "-nodecount=80", os.Args[0], name).CombinedOutput()
				must(err)
				fmt.Printf("lines of %s %s n=%d (ordered):\n%s\n", kind, profile, n, out)
			}
		}
	}
	if baseKit == nil {
		fmt.Printf("lookups %s %s n=%d: ordered %.1f ns\n", kind, profile, n, timed(runs[0].fn))
		return
	}
	fmt.Printf("lookups %s %s n=%d: ordered %.1f ns, baseline (%s) %.1f ns\n", kind, profile, n, timed(runs[0].fn), baseKit.ref, timed(runs[1].fn))
}
