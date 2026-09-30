package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/TomTonic/multimap/bench/awake"
	"github.com/TomTonic/multimap/bench/keys"
	"github.com/TomTonic/multimap/bench/rtopt"
	"github.com/TomTonic/rtcompare"
	"github.com/TomTonic/rtcompare/multiproc"
)

// drive runs the whole suite: every speed scenario with as many processes as
// its comparisons need (see driveScenario), then the memory measurements, and
// writes the raw results and two summary tables to c.out. All speed scenarios
// of one run share a regime, serial or parallel (see config.parallel), which
// run.json records.
func drive(c config) error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(c.out, "logs"), 0o755); err != nil {
		return err
	}
	start := time.Now()
	logf("start: %s", strings.Join(os.Args[1:], " "))
	if c.parallel > 1 {
		logf("parallel regime: %d processes at a time; the results are not comparable with serial runs%s", c.parallel, memoryPerProcess(c.parallel))
	}
	if release, err := awake.Hold("benchmark run"); err != nil {
		logf("warning: %v; the run goes on, and processes the machine slept through are flagged", err)
	} else {
		defer release()
	}
	childProcs := 0 // GOMAXPROCS of the speed processes of a parallel run, see multiproc.Results
	if !c.skipSpeed {
		path := filepath.Join(c.out, "speed.jsonl")
		if err := removeIfExists(path); err != nil {
			return err
		}
		for _, profile := range c.profiles {
			for _, kind := range c.kinds {
				for _, n := range c.sizes {
					g, err := driveScenario(c, kind, profile, n)
					if err != nil {
						return err
					}
					childProcs = max(childProcs, g)
				}
			}
		}
		all, err := readLines[result](path)
		if err != nil {
			return err
		}
		c.childProcs = childProcs
		if err := writeSpeed(c, all); err != nil {
			return err
		}
	}
	if !c.skipMem {
		if err := removeIfExists(filepath.Join(c.out, "mem.jsonl")); err != nil {
			return err
		}
		if err := driveMem(c, self); err != nil {
			return err
		}
	}
	logf("done after %s", time.Since(start).Round(time.Second))
	c.childProcs = childProcs
	return writeRunInfo(c, start)
}

// driveScenario runs the speed processes of one key kind, value profile and
// size through multiproc: each child builds the scenario's fixture after its
// own heap perturbation and in its own build order. The first c.minProcs
// processes decide once how many the scenario needs for every comparison's
// pooled interval to be precise (within c.abs or c.rel), and that many run,
// at most c.maxProcs, one after another or c.parallel at a time. Every
// process's reports go to speed.jsonl, the children's rtcompare reports to one
// log per scenario. It returns the GOMAXPROCS the children ran with, zero if
// they kept the runtime's default.
func driveScenario(c config, kind, profile string, n int) (int, error) {
	ps := pairsFor(keys.Kind(kind), n, profile, c.ops, c.scanMax, c.buildMax)
	if len(ps) == 0 {
		return 0, nil
	}
	if limit := keys.Capacity(keys.Kind(kind)); n > limit {
		logf("%s %s n=%d: skipped, the corpus holds keys for n <= %d", kind, profile, n, limit)
		return 0, nil
	}
	name := fmt.Sprintf("%s-%s-%d", kind, profile, n)
	log, err := os.Create(filepath.Join(c.out, "logs", "speed-"+name+".log"))
	if err != nil {
		return 0, err
	}
	defer func() { _ = log.Close() }()
	res, err := multiproc.Run(c.procOptions(childArgs(c, kind, profile, n), log, c.progress(fmt.Sprintf("%s %s n=%d", kind, profile, n), len(ps))),
		speedSuite(keys.Kind(kind), profile, n, c.stream(), ps))
	if err != nil {
		return 0, fmt.Errorf("%s: %w (see logs/speed-%s.log)", name, err, name)
	}
	if !res.Precise {
		logf("warning: %s %s n=%d stopped after %d processes, before every interval was as precise as asked", kind, profile, n, res.Processes)
	}
	var rows []result
	for _, cmp := range res.Comparisons {
		for i, r := range cmp.Reports {
			rows = append(rows, rowOf(kind, profile, n, pairOf(cmp.Name), i+1, res.Seeds[i], cmp.Pooled.FirstStage, res.Parallel, r))
			if r.Suspended > 0 {
				logf("warning: the machine slept %s during process %d of %s", r.Suspended.Round(time.Second), i+1, name)
			}
		}
	}
	b, err := encodeLines(rows)
	if err != nil {
		return 0, err
	}
	if err := appendFile(filepath.Join(c.out, "speed.jsonl"), b); err != nil {
		return 0, err
	}
	checkPooled(name, res, rows)
	return res.ChildGOMAXPROCS, nil
}

// memoryPerProcess says how much memory each of parallel processes may use,
// where the system reports what is free (Linux), for the log: a process of 1M
// keys holds every candidate of its scenario and the streams' structures, and
// the kernel ends a run that outgrows the machine's memory without a word.
func memoryPerProcess(parallel int) string {
	b, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return ""
	}
	for line := range strings.Lines(string(b)) {
		if rest, ok := strings.CutPrefix(line, "MemAvailable:"); ok {
			if kb, err := strconv.ParseFloat(strings.Fields(rest)[0], 64); err == nil {
				return fmt.Sprintf("; %.1f GB of memory are available, %.1f GB for each process (a process at 1M keys needs 4-5 GB against the baseline, up to 8 GB with all candidates)", kb/1e6, kb/1e6/float64(parallel))
			}
		}
	}
	return ""
}

// checkPooled warns if pooling the rows of a comparison, as the summary will,
// gives another result than multiproc's own pooling in the run: the summary
// would then state figures the run never saw.
func checkPooled(scenario string, res multiproc.Results, rows []result) {
	byName := map[string][]result{}
	for _, r := range rows {
		k := pair{r.Op, r.A, r.B}.name()
		byName[k] = append(byName[k], r)
	}
	for _, cmp := range res.Comparisons {
		if cmp.Pooled.Processes == 0 {
			continue
		}
		got, err := poolRows(byName[cmp.Name])
		if err != nil || got.Delta != cmp.Pooled.Delta || got.Low != cmp.Pooled.Low || got.High != cmp.Pooled.High || got.NoiseFloor != cmp.Pooled.NoiseFloor {
			logf("warning: %s: %s pools to %+v (%v) from its rows, but multiproc pooled %+v", scenario, cmp.Name, got, err, cmp.Pooled)
		}
	}
}

// procOptions are the multiproc options of a speed scenario. The seed is
// fixed per scenario, so that a run can be repeated process for process.
func (c config) procOptions(args []string, log io.Writer, progress func(multiproc.Results)) multiproc.Options {
	return multiproc.Options{
		MinProcesses: c.minProcs, MaxProcesses: c.maxProcs, Parallel: c.parallel, AbsPrecision: c.abs, RelPrecision: c.rel,
		Seed: 0x5EED, Args: args, Stderr: log, Progress: progress,
	}
}

// progress returns the callback that logs a scenario's progress after each
// process (after each wave in a parallel run), and, once the first stage has
// run, how many processes the scenario was sized for.
func (c config) progress(scenario string, comparisons int) func(multiproc.Results) {
	sized := false
	return func(r multiproc.Results) {
		first := 0
		for _, cmp := range r.Comparisons {
			first = max(first, cmp.Pooled.FirstStage)
		}
		if first == 0 {
			logf("%s: %d processes done; the first stage is at least %d", scenario, r.Processes, c.minProcs)
			return
		}
		if !sized {
			sized = true
			need, name := needed(r, first, c.abs, c.rel)
			logf("%s: first stage of %d processes done; %s needs %d processes for its interval, run in whole %s (at most %s)",
				scenario, first, name, need, c.unit(), c.limit())
		}
		open, widest := imprecise(r, c.abs, c.rel)
		logf("%s: %d processes done; %d of %d comparisons not yet within the interval asked for%s", scenario, r.Processes, open, comparisons, widest)
	}
}

// unit and limit describe how a run's size is rounded and capped, for the log.
func (c config) unit() string {
	if c.parallel > 1 {
		return fmt.Sprintf("waves of %d", c.parallel)
	}
	return "pairs of processes"
}

func (c config) limit() string {
	switch {
	case c.maxProcs != 0:
		return strconv.Itoa(c.maxProcs)
	case c.parallel > 1:
		return fmt.Sprintf("%d waves", multiproc.DefaultMaxWaves)
	}
	return strconv.Itoa(multiproc.DefaultMaxProcesses)
}

// needed returns how many processes the most demanding comparison of r asks
// for, judged from the first stage's processes, and its name; multiproc sizes
// the run from the same figures, then rounds up and caps it.
func needed(r multiproc.Results, first int, abs, rel float64) (int, string) {
	most, name := first, ""
	for _, cmp := range r.Comparisons {
		p, err := rtcompare.Combine(cmp.Reports[:first], cmp.Pooled.Level)
		if err != nil {
			continue
		}
		if k := p.ProcessesFor(abs, rel); k > most {
			most, name = k, cmp.Name
		}
	}
	if name == "" {
		return most, "no comparison"
	}
	return most, name
}

// childArgs is the command line of a scenario's speed processes.
func childArgs(c config, kind, profile string, n int) []string {
	args := []string{"-child", "-keys", kind, "-values", profile, "-n", strconv.Itoa(n), "-ops", strings.Join(c.ops, ","),
		"-scanmax", strconv.Itoa(c.scanMax), "-buildmax", strconv.Itoa(c.buildMax),
		"-ratio", strconv.FormatFloat(c.ratio, 'g', -1, 64),
		"-minprocs", strconv.Itoa(c.minProcs), "-maxprocs", strconv.Itoa(c.maxProcs), "-vs", strings.Join(vsOnly, ",")}
	return append(args, rtopt.Forward()...)
}

// imprecise counts the comparisons whose pooled interval is not yet tight
// enough, or not yet pooled, and describes the widest one.
func imprecise(r multiproc.Results, abs, rel float64) (int, string) {
	open, worst, name := 0, 0.0, ""
	for _, cmp := range r.Comparisons {
		p := cmp.Pooled
		if p.Processes >= 3 && p.Precise(abs, rel) {
			continue
		}
		open++
		if h := (p.High - p.Low) / 2; p.Processes >= 3 && h > worst {
			worst, name = h, cmp.Name
		}
	}
	if name == "" {
		return open, ""
	}
	return open, fmt.Sprintf("; widest ±%.1f pts: %s", worst*100, name)
}

// driveMem measures every candidate of every value profile, and a baseline
// per profile, in c.memRounds rounds of separate processes, in a new random
// order each round.
func driveMem(c config, self string) error {
	type job struct{ profile, impl string }
	var jobs []job
	for _, p := range c.profiles {
		for _, impl := range append([]string{"none"}, implsFor(p)...) {
			jobs = append(jobs, job{p, impl})
		}
	}
	path := filepath.Join(c.out, "mem.jsonl")
	var rows []memResult
	for round := 1; round <= c.memRounds; round++ {
		order := slices.Clone(jobs)
		rand.New(rand.NewPCG(uint64(round), 7)).Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })
		for _, kind := range c.kinds {
			for _, jb := range order {
				args := []string{"-memchild", jb.impl, "-keys", kind, "-values", jb.profile, "-n", strconv.Itoa(memN(c, kind)),
					"-cycles", strconv.Itoa(c.cycles), "-seed", strconv.Itoa(round)}
				out, err := runChild(self, args, filepath.Join(c.out, "logs", fmt.Sprintf("mem-%s-%s-%s-r%d.log", kind, jb.profile, jb.impl, round)))
				if err != nil {
					return err
				}
				if err := appendFile(path, out); err != nil {
					return err
				}
				got, err := decodeLines[memResult](out)
				if err != nil {
					return err
				}
				rows = append(rows, got...)
			}
		}
		logf("memory round %d of %d done", round, c.memRounds)
	}
	return writeMem(c, rows)
}

// memN is the number of keys the memory measurements use for kind: c.memN,
// or less where the kind's corpus is smaller (see keys.Capacity).
func memN(c config, kind string) int { return min(c.memN, keys.Capacity(keys.Kind(kind))) }

// runChild runs this binary with args, returns its stdout and writes its
// stderr (the rtcompare reports) to logPath.
func runChild(self string, args []string, logPath string) ([]byte, error) {
	log, err := os.Create(logPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = log.Close() }()
	var out bytes.Buffer
	cmd := exec.Command(self, args...)
	cmd.Stdout, cmd.Stderr = &out, log
	start := time.Now()
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s %s: %w (see %s)", filepath.Base(self), strings.Join(args, " "), err, logPath)
	}
	if s := awake.Slept(start); s > 0 {
		logf("warning: the machine slept %s during %s", s.Round(time.Second), filepath.Base(logPath))
	}
	return out.Bytes(), nil
}

// readLines decodes a JSON-lines file; a missing file holds no lines.
func readLines[T any](path string) ([]T, error) {
	b, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return decodeLines[T](b)
}

func decodeLines[T any](b []byte) ([]T, error) {
	var out []T
	sc := bufio.NewScanner(bytes.NewReader(b))
	for sc.Scan() {
		var v T
		if err := json.Unmarshal(sc.Bytes(), &v); err != nil {
			return nil, err
		}
		out = append(out, v)
	}
	return out, sc.Err()
}

func encodeLines[T any](rows []T) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	for _, r := range rows {
		if err := enc.Encode(r); err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil
}

// appendFile appends b to path; each process's lines are kept as they came.
func appendFile(path string, b []byte) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	if _, err := f.Write(b); err != nil {
		_ = f.Close()
		return err
	}
	return f.Close()
}

func removeIfExists(path string) error {
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func logf(format string, args ...any) {
	fmt.Printf("%s  %s\n", time.Now().Format("15:04:05"), fmt.Sprintf(format, args...))
}

// writeRunInfo records when and on what the suite ran.
func writeRunInfo(c config, start time.Time) error {
	path := filepath.Join(c.out, "run.json")
	info := map[string]any{
		"start": start.Format(time.RFC3339), "end": time.Now().Format(time.RFC3339),
		"go": runtime.Version(), "os": runtime.GOOS, "arch": runtime.GOARCH, "cpus": runtime.NumCPU(),
		"cpu": cpuName(), "args": os.Args[1:], "rtcompare": rtcompareVersion(),
		"minprocs": c.minProcs, "maxprocs": c.maxProcs, "ratio": c.ratio, "permchurn": c.permChurn, "abs": c.abs, "rel": c.rel,
		"parallel": c.parallel, "child_gomaxprocs": c.childProcs,
		"values": c.profiles, "keys": c.kinds, "sizes": c.sizes, "ops": c.ops, "memn": c.memN,
		"suite": flag.Lookup("suite").Value.String(), "vs": vsOnly, "value_type": fmt.Sprintf("%T", *new(V)),
	}
	if baseKit != nil {
		info["baseline"] = baseKit.ref
	}
	b, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

// rtcompareVersion is the version of rtcompare this binary was built with.
func rtcompareVersion() string {
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, d := range bi.Deps {
			if d.Path == "github.com/TomTonic/rtcompare" {
				return d.Version
			}
		}
	}
	return ""
}

// cpuName names the processor, so that results from different machines are
// never mistaken for each other. It returns "" where it cannot tell.
func cpuName() string {
	switch runtime.GOOS {
	case "darwin":
		out, _ := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
		return strings.TrimSpace(string(out))
	case "windows":
		out, _ := exec.Command("reg", "query", `HKLM\HARDWARE\DESCRIPTION\System\CentralProcessor\0`, "/v", "ProcessorNameString").Output()
		_, name, _ := strings.Cut(string(out), "REG_SZ")
		return strings.TrimSpace(name)
	}
	b, _ := os.ReadFile("/proc/cpuinfo")
	return cpuInfoModel(string(b))
}

// cpuInfoModel extracts the model name from the text of /proc/cpuinfo.
func cpuInfoModel(s string) string {
	for line := range strings.Lines(s) {
		if k, v, ok := strings.Cut(line, ":"); ok && strings.TrimSpace(k) == "model name" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
