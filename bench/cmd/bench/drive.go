package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/TomTonic/multimap/bench/awake"
	"github.com/TomTonic/multimap/bench/rtopt"
	"github.com/TomTonic/multimap/bench/stats"
)

// drive runs the whole suite: every speed scenario with as many processes as
// its comparisons need, then the memory measurements, and writes the raw
// results and two summary tables to c.out.
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
		if !c.cont {
			if err := removeIfExists(path); err != nil {
				return err
			}
		}
		prior, err := readLines[result](path)
		if err != nil {
			return err
		}
		for _, profile := range c.profiles {
			for _, kind := range c.kinds {
				for _, n := range c.sizes {
					if err := driveScenario(c, self, kind, profile, n, prior); err != nil {
						return err
					}
				}
			}
		}
		// The summary pools the whole file: with -continue, it also holds
		// the scenarios this run did not touch.
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

// driveScenario runs speed processes for one key kind, value profile and size
// until every comparison is precise (after at least c.minProcs) or c.maxProcs
// is reached. The processes of prior that belong to the scenario count as
// already run, so -continue resumes after the last of them.
func driveScenario(c config, self, kind, profile string, n int, prior []result) error {
	ps := pairsFor(n, profile, c.ops, c.scanMax, c.buildMax)
	if len(ps) == 0 {
		return nil
	}
	rows, done := scenarioRows(prior, kind, profile, n, ps)
	if open, _ := imprecise(rows, c.abs, c.rel); done > 0 && (done >= c.maxProcs || done >= c.minProcs && open == 0) {
		logf("%s %s n=%d: %d processes already run, %d of %d comparisons not yet precise", kind, profile, n, done, open, len(ps))
		return nil
	}
	for i := done + 1; i <= c.maxProcs; i++ {
		args := []string{"-child", "-keys", kind, "-values", profile, "-n", strconv.Itoa(n), "-ops", strings.Join(c.ops, ","),
			"-scanmax", strconv.Itoa(c.scanMax), "-buildmax", strconv.Itoa(c.buildMax),
			"-ratio", strconv.FormatFloat(c.ratio, 'g', -1, 64),
			"-layoutseed", strconv.Itoa(i)}
		out, err := runChild(self, append(args, rtopt.Forward()...), filepath.Join(c.out, "logs", fmt.Sprintf("speed-%s-%s-%d-p%02d.log", kind, profile, n, i)))
		if err != nil {
			return err
		}
		if err := appendFile(filepath.Join(c.out, "speed.jsonl"), out); err != nil {
			return err
		}
		got, err := decodeLines[result](out)
		if err != nil {
			return err
		}
		rows = append(rows, got...)
		open, widest := imprecise(rows, c.abs, c.rel)
		logf("%s %s n=%d process %d: %d of %d comparisons not yet precise%s", kind, profile, n, i, open, len(ps), widest)
		if i >= c.minProcs && open == 0 {
			break
		}
	}
	return nil
}

// scenarioRows picks the rows of prior that compare one of ps in the scenario
// and returns them with the number of processes behind them, which is the
// highest layout seed: the driver numbers a scenario's processes 1, 2, ...
// and passes that number as the seed. -continue needs both to resume.
func scenarioRows(prior []result, kind, profile string, n int, ps []pair) (rows []result, done int) {
	for _, r := range prior {
		if r.Keys == kind && r.Values == profile && r.N == n && slices.Contains(ps, pair{r.Op, r.A, r.B}) {
			rows = append(rows, r)
			done = max(done, int(r.LayoutSeed))
		}
	}
	return rows, done
}

// imprecise counts the comparisons whose interval across processes is not yet
// tight enough and describes the widest one.
func imprecise(rows []result, abs, rel float64) (int, string) {
	open, worst, name := 0, 0.0, ""
	for k, s := range summaries(rows) {
		if s.Precise(abs, rel) {
			continue
		}
		open++
		if h := s.HalfWidth(); h > worst {
			worst, name = h, k
		}
	}
	if open == 0 {
		return 0, ""
	}
	if math.IsInf(worst, 1) {
		return open, ""
	}
	return open, fmt.Sprintf("; widest ±%.1f pts: %s", worst*100, name)
}

// summaries pools the rows of each comparison across processes.
func summaries(rows []result) map[string]stats.Summary {
	type acc struct {
		d, h []float64
		r    []bool
	}
	g := map[string]*acc{}
	for _, r := range rows {
		k := fmt.Sprintf("%s %s n=%d %s %s vs %s", r.Keys, r.Values, r.N, r.Op, r.A, r.B)
		if g[k] == nil {
			g[k] = &acc{}
		}
		g[k].d = append(g[k].d, r.Delta)
		g[k].h = append(g[k].h, (r.High-r.Low)/2)
		g[k].r = append(g[k].r, r.Resolved)
	}
	out := map[string]stats.Summary{}
	for k, a := range g {
		out[k] = stats.Summarize(a.d, a.h, a.r)
	}
	return out
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
				args := []string{"-memchild", jb.impl, "-keys", kind, "-values", jb.profile, "-n", strconv.Itoa(c.memN),
					"-cycles", strconv.Itoa(c.cycles), "-layoutseed", strconv.Itoa(round)}
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

// writeRunInfo records when and on what the suite ran. With -continue, it
// keeps the record of the run it continues and adds this one under
// "continued".
func writeRunInfo(c config, start time.Time) error {
	path := filepath.Join(c.out, "run.json")
	info := map[string]any{
		"start": start.Format(time.RFC3339), "end": time.Now().Format(time.RFC3339),
		"go": runtime.Version(), "os": runtime.GOOS, "arch": runtime.GOARCH, "cpus": runtime.NumCPU(),
		"cpu": cpuName(), "args": os.Args[1:],
		"minprocs": c.minProcs, "maxprocs": c.maxProcs, "ratio": c.ratio, "abs": c.abs, "rel": c.rel,
		"values": c.profiles,
	}
	if c.cont {
		info = continued(path, info)
	}
	b, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

// continued returns the record in path with info appended to its
// "continued" list, or info itself if path holds no record. A record from
// before the CPU was known on Linux gets it now.
func continued(path string, info map[string]any) map[string]any {
	var old map[string]any
	if b, err := os.ReadFile(path); err != nil || json.Unmarshal(b, &old) != nil || old == nil {
		return info
	}
	if cpu, _ := old["cpu"].(string); cpu == "" {
		old["cpu"] = info["cpu"]
	}
	list, _ := old["continued"].([]any)
	old["continued"] = append(list, info)
	return old
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
