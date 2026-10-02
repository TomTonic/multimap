# Measuring

How speed and memory are measured for the redesign, on the Windows desktop (x86-64) and on an
arm64 machine. The bench itself is documented in [bench/README.md](../../bench/README.md).

## Rules

- **Speed claims** come only from rtcompare comparisons that run both candidates interleaved
  (the bench's `-vs` modes). Microbenchmarks and pprof are for diagnosis only, never for a claim.
- **Never pool** the serial regime with the parallel one (`-suite parallel`).
- **Announce every long run** with its duration and expected end time (clock time), and say so
  when the estimate changes.
- **Keep the machine quiet** while a run is going: no builds, no tests, no other measurements.
- **The reference** is the step before (`-vs baseline`, see below). For the credo also compare
  with the competitors: `btree-map` for unique values, `hashed`, `btree-sets` and `map-sets` for
  several.
- **Sizes:** the dev suite measures 4K, 16K and 256K, and `build` up to 64K. 1M only as a spot
  check or in the release suite. Always look at 4K and 16K first (credo 4).
- **Windows processes:** stop only your own (by PID, found via `Get-CimInstance Win32_Process`
  where the command line holds the script name). Never `taskkill /IM cmd.exe`: that kills the
  user's windows.

## Baseline

`bench/cmd/mkbaseline` copies the library as of a git ref into `bench/baseline` (generated,
ignored by git). Build tag `baseline` compares against that copy:

```
cd bench
go run ./cmd/mkbaseline -ref <commit>      # e.g. the last commit of the previous step
go run -tags baseline ./cmd/bench -vs baseline ...
```

## Windows (the main measuring machine, Ryzen 9 7900)

WSL measurements are noisy. Measure natively on Windows: cross-compile, copy into
`C:\temp\bench-win`, start detached.

```
cd bench
GOOS=windows GOARCH=amd64 go build -tags baseline         -o /mnt/c/temp/bench-win/<name>.exe     ./cmd/bench
GOOS=windows GOARCH=amd64 go build -tags baseline,strvals -o /mnt/c/temp/bench-win/<name>-str.exe ./cmd/bench
```

A run script `C:\temp\bench-win\runN.cmd` chains the measurements, for example:

```
@echo off
cd /d C:\temp\bench-win
<name>.exe -suite dev -vs baseline -sizes 4096,16384,262144 -minprocs 6 -maxprocs 12 -memn 1048576 -memrounds 3 -out C:\temp\bench-win\<result> > <result>.log 2>&1
<name>.exe -suite dev -values unique -vs btree-map -sizes 4096,16384,262144 -minprocs 6 -maxprocs 12 -skipmem -out C:\temp\bench-win\<result2> > <result2>.log 2>&1
```

Start it from WSL so that it survives the session:

```
powershell.exe -NoProfile -Command "Start-Process -FilePath cmd.exe -ArgumentList '/c','C:\temp\bench-win\runN.cmd' -WindowStyle Hidden"
```

**Durations** on this machine:
- the whole dev suite: about 2.5 h;
- u64 only, 4 sizes, with memory: about 1 h;
- `n3-u64` + `c7-unique-u64` + `n3-lp` together: 3 h 10 min.

Each process writes its progress into the `.log`. Estimate the end from the first finished
processes.

**Results** go into the repository under `bench/results-layout/<branch-or-step>/` with a
README in the style of `bench/results-layout/node-pages/README.md`:
- what was measured, and against what;
- tables with range and median over key kinds;
- findings, including the ones that speak against the change.

`C:\temp\bench-win` keeps the raw files. WSL's `/tmp` and the session scratchpad are wiped
when the VM restarts: never keep results only there.

Useful flags:
- `-suite dev|release|parallel`
- `-vs baseline|hashed|btree-sets|map-sets|btree-map`
- `-values multi|unique`
- `-keys u64,str,...`
- `-ops valuesFor,valuesBetween,prefix,churn,build`
- `-sizes`
- `-minprocs`, `-maxprocs`
- `-memn`, `-memrounds`
- `-skipmem`

## arm64 via a git job queue

The user has an arm64 machine (128-byte cache lines) that they use for other work during the
day. They start single jobs there by hand. Nothing runs without them. The sync goes through
git; PLAN step 0 builds it.

- **Queue:** `bench/remote/queue.txt` on the working branch. One job per line:
  `<id> <ref> <baseline-ref or -> <duration-minutes> <bench arguments>`, for example
  `a1 cacheline-s1 node-pages 90 -suite dev -keys u64,str,uuid -vs baseline`. With a baseline
  ref the runner calls `mkbaseline -ref <baseline-ref>` and builds with tag `baseline`.
  The agent appends jobs, commits and pushes (with the user's consent).
- **Runner:** `bench/remote/arm-run.sh [id]`, for macOS (arm64) and Linux arm64, started by the
  user. With no id it picks the first job that has no result yet. Steps:
  1. `git fetch`; check out the job's ref in a separate worktree;
  2. run `mkbaseline` if the job names a baseline ref;
  3. print the duration and expected end;
  4. run the bench under `caffeinate -i` (macOS) and `nice`;
  5. write the results and an `env.txt` (`uname -a`, CPU model, cache line size via
     `sysctl hw.cachelinesize` or `getconf LEVEL1_DCACHE_LINESIZE`, Go version, power source)
     into `<id>/` on branch `arm-results`;
  6. commit and push that branch.
- **Evaluation:** the agent runs `git fetch origin arm-results` and reads the results like the
  Windows ones. Every arm64 claim names the machine.
- **When to run:** jobs should take at most 2-3 h, so that a lunch break or an evening fits one.
  Results from a machine in use are invalid. The runner warns if the user is active (load
  average) and says so in `env.txt`.
- **Prerequisites** on the arm64 machine:
  - Go at the version in `go.mod`;
  - a clone of the repository with push rights;
  - for Linux, `nice` and `getconf`.
