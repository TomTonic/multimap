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
		if err := removeIfExists(filepath.Join(c.out, "speed.jsonl")); err != nil {
			return err
		}
		var all []result
		for _, kind := range c.kinds {
			for _, n := range c.sizes {
				rows, err := driveScenario(c, self, kind, n)
				if err != nil {
					return err
				}
				all = append(all, rows...)
			}
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

// driveScenario runs speed processes for one key kind and size until every
// comparison is precise (after at least c.minProcs) or c.maxProcs is reached.
func driveScenario(c config, self, kind string, n int) ([]result, error) {
	ps := pairsFor(n, c.ops, c.scanMax, c.buildMax)
	if len(ps) == 0 {
		return nil, nil
	}
	var rows []result
	for i := 1; i <= c.maxProcs; i++ {
		args := []string{"-child", "-keys", kind, "-n", strconv.Itoa(n), "-ops", strings.Join(c.ops, ","),
			"-scanmax", strconv.Itoa(c.scanMax), "-buildmax", strconv.Itoa(c.buildMax),
			"-ratio", strconv.FormatFloat(c.ratio, 'g', -1, 64),
			"-layoutseed", strconv.Itoa(i)}
		out, err := runChild(self, append(args, rtopt.Forward()...), filepath.Join(c.out, "logs", fmt.Sprintf("speed-%s-%d-p%02d.log", kind, n, i)))
		if err != nil {
			return nil, err
		}
		if err := appendFile(filepath.Join(c.out, "speed.jsonl"), out); err != nil {
			return nil, err
		}
		got, err := decodeLines[result](out)
		if err != nil {
			return nil, err
		}
		rows = append(rows, got...)
		open, widest := imprecise(rows, c.abs, c.rel)
		logf("%s n=%d process %d: %d of %d comparisons not yet precise%s", kind, n, i, open, len(ps), widest)
		if i >= c.minProcs && open == 0 {
			break
		}
	}
	return rows, nil
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
		k := fmt.Sprintf("%s n=%d %s %s vs %s", r.Keys, r.N, r.Op, r.A, r.B)
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

// driveMem measures every candidate and a baseline in c.memRounds rounds of
// separate processes, in a new random order each round.
func driveMem(c config, self string) error {
	impls := append([]string{"none"}, c.impls...)
	path := filepath.Join(c.out, "mem.jsonl")
	var rows []memResult
	for round := 1; round <= c.memRounds; round++ {
		order := slices.Clone(impls)
		rand.New(rand.NewPCG(uint64(round), 7)).Shuffle(len(order), func(i, j int) { order[i], order[j] = order[j], order[i] })
		for _, kind := range c.kinds {
			for _, impl := range order {
				args := []string{"-memchild", impl, "-keys", kind, "-n", strconv.Itoa(c.memN),
					"-cycles", strconv.Itoa(c.cycles), "-layoutseed", strconv.Itoa(round)}
				out, err := runChild(self, args, filepath.Join(c.out, "logs", fmt.Sprintf("mem-%s-%s-r%d.log", kind, impl, round)))
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

// writeRunInfo records when and on what the suite ran.
func writeRunInfo(c config, start time.Time) error {
	cpu, _ := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
	info := map[string]any{
		"start": start.Format(time.RFC3339), "end": time.Now().Format(time.RFC3339),
		"go": runtime.Version(), "os": runtime.GOOS, "arch": runtime.GOARCH, "cpus": runtime.NumCPU(),
		"cpu": strings.TrimSpace(string(cpu)), "args": os.Args[1:],
		"minprocs": c.minProcs, "maxprocs": c.maxProcs, "ratio": c.ratio, "abs": c.abs, "rel": c.rel,
	}
	b, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(c.out, "run.json"), append(b, '\n'), 0o644)
}
