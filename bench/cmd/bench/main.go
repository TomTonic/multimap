// Command bench measures what a user choosing a Go multimap wants to know:
// multimap.Ordered and multimap.Hashed against the hand-written alternatives
// (a tidwall/btree.Map or a Go map of Go-map sets), for reading a key's
// values, range queries, the insertions and deletions of a database index
// (churn and build), memory, GC cost, and memory after removing keys.
//
// Every speed comparison runs in separate processes, each with its own heap
// layout (-layoutseed), because rtcompare's interval covers only the noise
// within one process. The driver starts at least -minprocs processes per
// scenario and adds more until every comparison's 95% interval across
// processes is within -abs percentage points or -rel of the difference, or
// -maxprocs is reached. See README.md.
//
// Usage:
//
//	go run ./cmd/bench                          # everything; takes hours
//	go run ./cmd/bench -sizes 4096 -skipmem     # a quicker subset
//
// The driver re-executes its own binary for every process (-child, -memchild).
package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/TomTonic/multimap/bench/keys"
)

type config struct {
	kinds, ops, impls  []string
	sizes              []int
	minProcs, maxProcs int
	abs, rel           float64
	scanMax            int
	ratio              float64
	buildMax           int
	memN, memRounds    int
	cycles             int
	out                string
	skipSpeed, skipMem bool
}

func main() {
	child := flag.Bool("child", false, "internal: run one speed process for -keys and -n")
	memChild := flag.String("memchild", "", "internal: run one memory process for this candidate")
	kindsF := flag.String("keys", "u64,str", "key kinds (u64, str); a single kind for -child")
	sizesF := flag.String("sizes", "4096,1048576", "numbers of keys for the speed comparisons")
	n := flag.Int("n", 4096, "internal: number of keys of a -child or -memchild process")
	opsF := flag.String("ops", "valuesFor,valuesBetween,churn,build", "operations to compare")
	var c config
	flag.IntVar(&c.minProcs, "minprocs", 5, "processes per scenario before the stop rule applies")
	flag.IntVar(&c.maxProcs, "maxprocs", 20, "processes per scenario at most")
	flag.Float64Var(&c.abs, "abs", 0.02, "stop once every 95% interval is within this many (fractional) points ...")
	flag.Float64Var(&c.rel, "rel", 0.10, "... or within this fraction of its own difference")
	flag.IntVar(&c.scanMax, "scanmax", 1<<16, "compare range queries on hashed and map-sets (a scan of all keys) only up to this many keys")
	flag.Float64Var(&c.ratio, "ratio", 2, "churn and build: insertions per value the multimap holds in the end (at least 1)")
	flag.IntVar(&c.buildMax, "buildmax", 1<<16, "compare building only up to this many keys")
	flag.IntVar(&c.memN, "memn", 1<<20, "number of keys for the memory measurements")
	flag.IntVar(&c.memRounds, "memrounds", 5, "processes per candidate and key kind for the memory measurements")
	flag.IntVar(&c.cycles, "cycles", 50, "forced GC cycles to time per memory process")
	flag.StringVar(&c.out, "out", "results", "directory for results and logs")
	flag.BoolVar(&c.skipSpeed, "skipspeed", false, "skip the speed comparisons")
	flag.BoolVar(&c.skipMem, "skipmem", false, "skip the memory measurements")
	flag.Parse()

	c.kinds, c.ops, c.impls = split(*kindsF), split(*opsF), allImpls
	var err error
	if c.ratio < 1 || c.ratio == 1 && slices.Contains(c.ops, "churn") {
		err = fmt.Errorf("-ratio %v: must be at least 1, and more than 1 for churn", c.ratio)
	}
	switch {
	case err != nil:
	case *child:
		err = runSpeed(keys.Kind(*kindsF), *n, c.ratio, pairsFor(*n, c.ops, c.scanMax, c.buildMax), os.Stdout)
	case *memChild != "":
		err = runMem(keys.Kind(*kindsF), *n, *memChild, c.cycles, os.Stdout)
	default:
		if c.sizes, err = atoiAll(split(*sizesF)); err == nil {
			err = drive(c)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "bench:", err)
		os.Exit(1)
	}
}

func split(s string) []string {
	var out []string
	for f := range strings.SplitSeq(s, ",") {
		if f = strings.TrimSpace(f); f != "" {
			out = append(out, f)
		}
	}
	return out
}

func atoiAll(ss []string) ([]int, error) {
	out := make([]int, len(ss))
	for i, s := range ss {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("bad size %q: %w", s, err)
		}
		out[i] = v
	}
	return out, nil
}
