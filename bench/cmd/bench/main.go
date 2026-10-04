// Command bench measures what a user choosing a Go multimap wants to know:
// multimap.Ordered and multimap.Hashed against the hand-written alternatives
// (a tidwall/btree.Map or a Go map of Go-map sets), for reading a key's
// values, range queries, the insertions and deletions of a database index
// (churn and build), memory, GC cost, and memory after removing keys. For
// keys that hold exactly one value, it also compares multimap.Ordered with a
// plain tidwall/btree.Map (-values single-value).
//
// Every speed comparison runs in separate processes, through rtcompare's
// multiproc package, because rtcompare's interval covers only the noise
// within one process: each process perturbs its heap from its own seed and
// builds the candidates in its own order, and the driver pools the processes
// with rtcompare.CombineStaged. The first -minprocs processes of a scenario
// show how much the processes scatter, from that the run works out once how
// many processes every comparison's pooled 95% interval needs to be within
// -abs percentage points or -rel of the difference, and runs that many, at
// most -maxprocs. They run one after another (the serial regime, each with
// the machine to itself) or, with -parallel, several at a time (the parallel
// regime, sharing the caches and the memory bandwidth). The two regimes
// measure different things and are never pooled. See README.md.
//
// Usage:
//
//	go run ./cmd/bench                          # the dev suite (-suite dev)
//	go run ./cmd/bench -suite release           # everything serial; takes a day
//	go run ./cmd/bench -suite parallel          # the out-of-cache sizes, 8 processes at a time
//	go run ./cmd/bench -sizes 4096 -skipmem     # a quicker subset
//	go run -tags baseline ./cmd/bench -vs baseline      # head to head with an earlier Ordered (see cmd/mkbaseline)
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
	"github.com/TomTonic/rtcompare/multiproc"
)

type config struct {
	kinds, ops         []string
	profiles           []string
	sizes              []int
	minProcs, maxProcs int
	parallel           int
	childProcs         int // GOMAXPROCS the children of a parallel run had, set by drive
	abs, rel           float64
	scanMax            int
	ratio, permChurn   float64
	buildMax           int
	memN, memRounds    int
	cycles             int
	seed               uint64
	out                string
	skipSpeed, skipMem bool
}

func main() {
	child := flag.Bool("child", false, "internal: run one speed process for -keys and -n")
	memChild := flag.String("memchild", "", "internal: run one memory process for this candidate")
	suite := flag.String("suite", "dev", "presets for the flags not given: dev (small and medium sizes, fewer processes and A/A runs, for frequent runs) or release (all sizes up to 1M at full precision)")
	kindsF := flag.String("keys", strings.Join(kindNames(), ","), "key kinds ("+strings.Join(kindNames(), ", ")+"); a single kind for -child")
	profilesF := flag.String("values", "natural,single-value", "value profiles: natural (a skewed number of values per key), single-value (one value per key); a single profile for -child")
	sizesF := flag.String("sizes", "4096,1048576", "numbers of keys for the speed comparisons")
	n := flag.Int("n", 4096, "internal: number of keys of a -child or -memchild process")
	opsF := flag.String("ops", "valuesFor,valuesBetween,prefix,churn,build", "operations to compare")
	vsF := flag.String("vs", "", "compare ordered only with these candidates (default: all of each value profile; baseline needs -tags baseline, see cmd/mkbaseline)")
	var c config
	flag.IntVar(&c.minProcs, "minprocs", 6, "processes in a scenario's first stage, whose scatter decides how many more it needs (in a parallel run: rounded up to whole waves)")
	flag.IntVar(&c.maxProcs, "maxprocs", 0, "processes per scenario at most (0: rtcompare's default, "+strconv.Itoa(multiproc.DefaultMaxProcesses)+" serial or "+strconv.Itoa(multiproc.DefaultMaxWaves)+" waves parallel)")
	flag.IntVar(&c.parallel, "parallel", 1, "processes to run at the same time: 1 runs them one after another (the serial regime); more runs them in waves (the parallel regime: the processes share caches and memory bandwidth, so results are never comparable with serial ones); keep it at or below the physical cores, and mind the memory, see README.md")
	flag.Float64Var(&c.abs, "abs", 0.02, "size the run so that every 95% interval is within this many (fractional) points ...")
	flag.Float64Var(&c.rel, "rel", 0.10, "... or within this fraction of its own difference")
	flag.IntVar(&c.scanMax, "scanmax", 1<<16, "compare range queries on hashed and map-sets (a scan of all keys) only up to this many keys")
	flag.Float64Var(&c.ratio, "ratio", 2, "churn and build: insertions per value the multimap holds in the end (at least 1)")
	flag.Float64Var(&c.permChurn, "permchurn", 0.25, "churn and build: share of the deletions that take out a long-lived value, which is put back later (0 to below 1; 0 leaves the corpus untouched, as rtcompare's default does)")
	flag.IntVar(&c.buildMax, "buildmax", 1<<16, "compare building only up to this many keys")
	flag.IntVar(&c.memN, "memn", 1<<20, "number of keys for the memory measurements")
	flag.IntVar(&c.memRounds, "memrounds", 5, "processes per candidate and key kind for the memory measurements")
	flag.IntVar(&c.cycles, "cycles", 50, "forced GC cycles to time per memory process")
	flag.Uint64Var(&c.seed, "seed", 0, "internal: heap perturbation seed of a -memchild process (0: none)")
	flag.StringVar(&c.out, "out", "results", "directory for results and logs")
	flag.BoolVar(&c.skipSpeed, "skipspeed", false, "skip the speed comparisons")
	flag.BoolVar(&c.skipMem, "skipmem", false, "skip the memory measurements")
	flag.Parse()
	if err := applySuite(flag.CommandLine, *suite); err != nil {
		fmt.Fprintln(os.Stderr, "bench:", err)
		os.Exit(2)
	}

	*profilesF = canonicalProfiles(*profilesF)
	c.kinds, c.ops, c.profiles = split(*kindsF), split(*opsF), split(*profilesF)
	vsOnly = split(*vsF)
	err := c.validate()
	switch {
	case err != nil:
	case *child:
		ps := pairsFor(keys.Kind(*kindsF), *n, *profilesF, c.ops, c.scanMax, c.buildMax)
		_, err = multiproc.Run(c.procOptions(nil, nil, nil), speedSuite(keys.Kind(*kindsF), *profilesF, *n, c.stream(), ps))
	case *memChild != "":
		err = runMem(keys.Kind(*kindsF), *profilesF, *n, *memChild, c.cycles, c.seed, os.Stdout)
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

// suites are the presets of -suite. dev is for the frequent runs while
// trying a change: it leaves out 1M keys, where one process tells least
// (see README.md), and trades precision for time with fewer processes and
// A/A runs. release is the complete suite at full precision, one process
// after another. parallel is its out-of-cache sizes, where the scatter between
// processes dominates and only more processes bring precision in reasonable
// time, run in the parallel regime (8 at a time: a process at 1M keys needs
// up to 5 GB, and 8 of them leave the 12 physical cores of the reference
// machine room for the garbage collectors; one wave is the first stage) and
// without the memory measurements, which the release suite already takes.
var suites = map[string]map[string]string{
	"dev": {"sizes": "4096,16384,262144", "minprocs": "4", "maxprocs": "8",
		"validation": "2", "repeats": "41", "buildrepeats": "21", "buildvalidation": "2", "memn": "131072", "memrounds": "3"},
	"release": {"sizes": "4096,16384,262144,1048576", "minprocs": "6", "maxprocs": "0",
		"validation": "0", "repeats": "0", "memn": "1048576", "memrounds": "5"},
	"parallel": {"sizes": "262144,1048576", "parallel": "8", "minprocs": "8", "maxprocs": "0",
		"validation": "0", "repeats": "0", "skipmem": "true"},
}

// applySuite sets the flags of fs that the named preset lists and the
// command line did not set itself.
func applySuite(fs *flag.FlagSet, name string) error {
	preset, ok := suites[name]
	if !ok {
		return fmt.Errorf("-suite: unknown suite %q", name)
	}
	given := map[string]bool{}
	fs.Visit(func(f *flag.Flag) { given[f.Name] = true })
	for k, v := range preset {
		if !given[k] {
			if err := fs.Set(k, v); err != nil {
				return fmt.Errorf("-suite %s: %w", name, err)
			}
		}
	}
	return nil
}

func kindNames() []string {
	out := make([]string, len(keys.Kinds))
	for i, k := range keys.Kinds {
		out[i] = string(k)
	}
	return out
}

// maxSingleValueRatio bounds -ratio for the single-value profile: every transient value
// takes an extra key of its own, and there are as many extra keys as corpus
// keys (see newPairs and extraKeys).
const maxSingleValueRatio = 2

// stream is the shape of the churn and build streams the flags ask for.
func (c config) stream() stream { return stream{c.ratio, c.permChurn} }

func (c *config) validate() error {
	for _, v := range vsOnly {
		switch {
		case !slices.Contains([]string{hashed, btreeSets, mapSets, btreeMapC, baseline}, v):
			return fmt.Errorf("-vs: unknown candidate %q", v)
		case v == baseline && baseKit == nil:
			return fmt.Errorf("-vs baseline: build the bench with -tags baseline after go run ./cmd/mkbaseline")
		}
	}
	for _, p := range c.profiles {
		if p != natural && p != singleValue {
			return fmt.Errorf("-values: unknown profile %q", p)
		}
	}
	switch {
	case c.ratio < 1 || c.ratio == 1 && slices.Contains(c.ops, "churn"):
		return fmt.Errorf("-ratio %v: must be at least 1, and more than 1 for churn", c.ratio)
	case c.minProcs < 3 || c.maxProcs != 0 && c.maxProcs < c.minProcs:
		return fmt.Errorf("-minprocs %d, -maxprocs %d: need 3 <= minprocs <= maxprocs, or maxprocs 0", c.minProcs, c.maxProcs)
	case c.parallel < 1:
		return fmt.Errorf("-parallel %d: must be at least 1", c.parallel)
	case c.permChurn < 0 || c.permChurn >= 1:
		return fmt.Errorf("-permchurn %v: must be at least 0 and below 1", c.permChurn)
	case c.ratio > maxSingleValueRatio && slices.Contains(c.profiles, singleValue):
		return fmt.Errorf("-ratio %v: at most %d with -values single-value", c.ratio, maxSingleValueRatio)
	}
	return nil
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

// canonicalProfiles returns the -values list with the names the profiles had
// until 2026-10-04 (multi, unique) replaced by their names now (natural,
// single-value), so that queue jobs and scripts from before still run.
func canonicalProfiles(list string) string {
	return strings.NewReplacer("multi", natural, "unique", singleValue).Replace(list)
}
