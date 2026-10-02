# Status

## 2026-10-02 afternoon: A/A job a0 on the M1 Pro

`a0` (same code against itself, 6b06dd2, u64 and str, 4K and 16K, multi and unique, four
operations) ran in 4.5 minutes, from 12:46 on 2026-10-02, results on `origin/arm-results`
(`a0/`). Of 32 comparisons:

- all 32 point estimates are within 0.97-1.01 (differences -2.8% to +0.9%), 30 within ±1.1%;
- 30 are precise by the harness's criterion (95% interval within ±2 points or 10% of the
  difference); the other two have wide intervals (±3.7 and ±4.0 points): `u64 unique 16K
  valuesFor` has one outlier process (+7.1% against -0.4% to +0.6% for the other seven),
  `u64 multi 16K churn` scatters between all eight (-4.7% to +3.0%). The scenarios were capped
  at 8 processes, the harness wanted 29 and 32;
- **systematic bias:** `valuesBetween` with str keys and one value per key reads 0.97 (4K) and
  0.98 (16K) in every one of its four processes (-2.0% to -3.2%) for identical code. The
  harness marks these "resolved"; they are not: its A/A validation runs inside one binary and
  does not see the layout difference between the library and the baseline copy;
- the runner's "load at start" in `env.txt` was my own compile (29): fixed, the runner now notes the load
  before the build, waits 30 s after it, and warns at a load of 2;
- the machine was not at rest (load 3.8, 15-minute average 6.1 at the start).

Gate 0 reading: numerically met (0.97-1.01), with the two caveats above. Job `a1` (queued, not
pushed yet) repeats a0 with 8 to 24 processes per scenario on a machine at rest.

## 2026-10-02: step 0 done, waiting for gate 0

**Parked and pushed:**

| branch | commit | content |
|---|---|---|
| `node-layout` | `7b8a8d8` | pessimistic paths, leaves with key remainders, flat and typed leaves, results against `main` and the competitors |
| `node-pages` | `cf6944b` | pages for integer-like keys with one value, range nodes, results in `bench/results-layout/node-pages/` |
| `leaf-pages` | `b0aeed1` | the earlier page experiments (K, S, U8-n pages) and their measurements |

**Working branch:** `cacheline`, forked from `node-pages`. It holds these documents and step 0.

**Step 0 (not yet committed, see PLAN.md):**

- 0.1 `art.Map.Objects`, `art.Block`, `bench/cmd/objstat`, tests, the statistic of `node-pages`.
- 0.2 `keys.Corpus.Probes`, used by `valuesFor`; `bench/README.md` says numbers are not comparable.
- 0.3 `bench/remote/` (queue, runner, test); run once end to end on Linux.
- 0.4 `go test ./... -race` 3.5 minutes instead of more than 20.

**Next action:** the user reviews step 0 and decides about the commit. After that the tools are
committed and pushed, the A/A job goes into the queue, and the user starts it on the arm64
machine. Then step 1 (PLAN.md).

**Open points:**

- **arm64 runner on macOS** is untested (see MEASURING.md).
- **Probe order**: the 64K-and-above numbers are expected to be comparable with the old order but
  that was not measured. A short A/A of old and new probes at 64K would settle it.
- **Race test time**: still 3.5 minutes. If CI needs less, shrink the key sets further.

## Test run time (2026-10-02)

`node-pages`, WSL, machine otherwise idle:

| run | `internal/art` | `TestAgainstReference` |
|---|---|---|
| `go test -race -cover ./...` (default 10-minute timeout) | failed: timeout | still running at 600 s |
| `go test -race -cover -timeout 60m -v ./internal/art` | passed, 100% coverage, 1175 s | 1121 s |
| `go test -run TestAgainstReference ./internal/art` (no race detector) | | 43.5 s |

All other tests together take about 55 s with the race detector; the slowest are
`TestRemoveAbsentNearMisses` (25 s) and `TestLongPaths` (7 s).

**Fixed in step 0.4:** under `-race` the big key sets of `TestAgainstReference` are a tenth of
their size and the 64K-path set runs in two of the five leaf modes. `go test -race -cover
./internal/art` then takes 205 s with 100% coverage (the whole `./...` run 208 s). One
deterministic set was added (a term arriving at a full node) because the smaller random sets no
longer reached `setTerm`'s growth.

`TestAgainstReference` runs five leaf modes over every key set, against the reference model.
Each run has six phases, and after each phase it compares the whole map and checks the
invariants of the whole tree. The race detector slows it down about 26 times.
