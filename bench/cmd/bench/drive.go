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
	"github.com/TomTonic/rtcompare/multiproc"
)

// drive runs the whole suite: every speed scenario with as many processes as
// its comparisons need (see driveScenario), then the memory measurements, and
// writes the raw results and two summary tables to c.out.
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
	if release, err := awake.Hold("benchmark run"); err != nil {
		logf("warning: %v; the run goes on, and processes the machine slept through are flagged", err)
	} else {
		defer release()
	}
	if !c.skipSpeed {
		path := filepath.Join(c.out, "speed.jsonl")
		if err := removeIfExists(path); err != nil {
			return err
		}
		for _, profile := range c.profiles {
			for _, kind := range c.kinds {
				for _, n := range c.sizes {
					if err := driveScenario(c, kind, profile, n); err != nil {
						return err
					}
				}
			}
		}
		all, err := readLines[result](path)
		if err != nil {
			return err
		}
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
	return writeRunInfo(c, start)
}

// driveScenario runs the speed processes of one key kind, value profile and
// size through multiproc: each child builds the scenario's fixture after its
// own heap perturbation and in its own build order, and processes are added
// until every comparison's pooled interval is precise (after at least
// c.minProcs) or c.maxProcs have run. Every process's reports go to
// speed.jsonl, the children's rtcompare reports to one log per scenario.
func driveScenario(c config, kind, profile string, n int) error {
	ps := pairsFor(keys.Kind(kind), n, profile, c.ops, c.scanMax, c.buildMax)
	if len(ps) == 0 {
		return nil
	}
	if limit := keys.Capacity(keys.Kind(kind)); n > limit {
		logf("%s %s n=%d: skipped, the corpus holds keys for n <= %d", kind, profile, n, limit)
		return nil
	}
	name := fmt.Sprintf("%s-%s-%d", kind, profile, n)
	log, err := os.Create(filepath.Join(c.out, "logs", "speed-"+name+".log"))
	if err != nil {
		return err
	}
	defer func() { _ = log.Close() }()
	res, err := multiproc.Run(c.procOptions(childArgs(c, kind, profile, n), log, func(r multiproc.Results) {
		open, widest := imprecise(r, c.abs, c.rel)
		logf("%s %s n=%d process %d: %d of %d comparisons not yet precise%s", kind, profile, n, r.Processes, open, len(ps), widest)
	}), speedSuite(keys.Kind(kind), profile, n, c.ratio, ps))
	if err != nil {
		return fmt.Errorf("%s: %w (see logs/speed-%s.log)", name, err, name)
	}
	var rows []result
	for _, cmp := range res.Comparisons {
		for i, r := range cmp.Reports {
			rows = append(rows, rowOf(kind, profile, n, pairOf(cmp.Name), i+1, res.Seeds[i], r))
			if r.Suspended > 0 {
				logf("warning: the machine slept %s during process %d of %s", r.Suspended.Round(time.Second), i+1, name)
			}
		}
	}
	b, err := encodeLines(rows)
	if err != nil {
		return err
	}
	return appendFile(filepath.Join(c.out, "speed.jsonl"), b)
}

// procOptions are the multiproc options of a speed scenario. The seed is
// fixed per scenario, so that a run can be repeated process for process.
func (c config) procOptions(args []string, log io.Writer, progress func(multiproc.Results)) multiproc.Options {
	return multiproc.Options{
		MinProcesses: c.minProcs, MaxProcesses: c.maxProcs, AbsPrecision: c.abs, RelPrecision: c.rel,
		Seed: 0x5EED, Args: args, Stderr: log, Progress: progress,
	}
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
		"minprocs": c.minProcs, "maxprocs": c.maxProcs, "ratio": c.ratio, "abs": c.abs, "rel": c.rel,
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
