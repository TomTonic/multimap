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

The user has an arm64 machine: a MacBook with an M1 Pro (8 performance and 2 efficiency cores,
16 GB, macOS, 128-byte cache lines). They use it for other work during the day and start single
jobs by hand. Nothing runs without them. The sync goes through git.

- **Queue:** `bench/remote/queue.txt` on the working branch (`cacheline`). One job per line:
  `<id> <ref> <baseline-ref or -> <duration-minutes> [tags=a,b] <arguments of cmd/bench>`, for
  example `a1 <commit> <baseline-commit> 90 -suite dev -keys u64,str -sizes 4096,16384 -vs baseline
  -minprocs 4 -maxprocs 8 -skipmem`. Refs are commits, so that a job means the same thing when it
  is run later. With a baseline ref the runner calls `mkbaseline -ref <baseline-ref>` and builds
  with the tag `baseline`. The agent appends jobs, commits and pushes (with the user's consent).
- **Runner:** `bench/remote/arm-run.sh [id]`, run by the user in their clone, on macOS or Linux
  arm64. Its header is the user's how-to. With no id it picks the first job with no result.
  `--list` shows the queue, `--dry-run` shows what it would do. Steps:
  1. fetch the remote, read the queue from `origin/cacheline`;
  2. print the machine, the duration and the expected end, and warn about battery power and a
     high load (and ask, unless `--yes`);
  3. build the job's commit in a temporary worktree (the user's working tree is not touched),
     with the baseline if the job names one;
  4. run the bench under `caffeinate -i` on macOS, with the output also on the terminal;
  5. write the results, `env.txt` (system, CPU, performance cores, memory, cache line size from
     `sysctl hw.cachelinesize` or `getconf LEVEL1_DCACHE_LINESIZE`, Go version, power source, load
     before and after), `args.txt`, `run.log` and a `DONE` marker into `<id>/` on the branch
     `arm-results`, commit it and push that branch.
- **Evaluation:** the agent runs `git fetch origin arm-results` and reads the results like the
  Windows ones: `git show origin/arm-results:<id>/bench-out/speed-summary.md`. Every arm64 claim
  names the machine (`env.txt`).
- **Limits of this machine:**
  - 16 GB: a process at 1M keys holds 4-5 GB, so no 1M keys on this machine and no memory phase
    beyond 256K (`-memn 262144`).
  - Eight performance cores. In the serial regime the processes of a scenario run one after
    another, so `-maxprocs` is only how many a scenario may use at most: use `-minprocs 8
    -maxprocs 24`. macOS cannot pin processes, so a process may land on one of the two
    efficiency cores, which shows as one outlier among the processes of a scenario (see the A/A
    job a0 in STATUS.md); more processes dilute it.
  - A/A figures of this machine (job a0, identical code as library and as baseline copy): 30 of
    32 comparisons within ±1.1% (all 32 between 0.97 and 1.01), but `valuesBetween` with str keys and one value per key reads
    0.97-0.98 in every process, so treat differences below 3% there as noise. The harness's own
    "resolved" mark does not know this bias (its A/A validations run inside one binary).
  - It runs on power only and idle: the runner warns about battery and load.
- **Durations:** job a0 (u64 and str, 4K and 16K, both profiles, four operations, 4-8 processes)
  took 4.5 minutes on this machine, one tenth of what I guessed from the Windows runs. Measure
  before you estimate.
- **When to run:** jobs of up to 2-3 h fit a lunch break or an evening. Results from a machine in
  use are invalid: the runner warns when the load is 2 or more (1 minute) or 3 or more (5 minutes),
  and waits 30 seconds after the build so that the compile does not count.
- **Prerequisites** on the arm64 machine: Go at the version in `go.mod`, a clone of the repository
  with push rights, and nothing else running. The runner is tested on Linux only; the macOS
  branches (`sysctl`, `pmset`, `caffeinate`, `date -v`) are untested until the first job.
