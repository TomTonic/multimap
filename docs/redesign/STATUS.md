# Status

## 2026-10-02 evening: step 1 done, waiting for gate 1

The page prototype (`internal/vpage`, 100% coverage, fuzzed, race and lint clean) and its
experiments are built and described in [step1-results.md](step1-results.md). It has a directory
(the user's idea of a FAT) in the first line or two of the page: a tag byte per key, in general
pages also the suffix length and the tail's offset. Every lookup takes **two dependent rounds**
(directory; then head, value and tail together), against 3 to 4 before, and cold lookups in the
microbenchmark got 23-44% faster (`u64` 157 ns, `uuid` 246, `path` 259; hot 9-14 ns).

Memory against the leaves of `node-pages`: `street` -39..-46%, `str` -15..-24%, `email` -17..-22%,
`uuid` -2%, `url` -3..-4%, but `path` +2..+6% and `u64` +2..+11% (the directory costs a byte a
key). A page prefix (not built) would make `path` and `url` -7..-10%.

Gate 1 as first written (memory not above leaves for any kind; at most two lines after the head)
was not met for `path` and `u64`, and literally not for the lines of general pages (3 to 4 after
the first), though they are read in one round.

**User's decision (2026-10-02):** the rule is restated as "at most two rounds of cache-line loads"
(R5 in STRATEGY.md, with a precise definition of a round). Continue with the page prefix, but its
cost in `churn` is the user's worry and decides whether it stays (see below). `u64`'s 6-10% more
memory is the price of 44% faster cold lookups. Commit and push at every valuable point: allowed
without asking.

**Next:** (1) page prefix in the prototype, measured for memory and for mutate cost (insert/delete
and split/merge), kept only if `churn` does not pay for it; (2) step 2 gets a lookup gate, in
plain words: after wiring the pages into the tree, point lookups of string keys at 16K-64K keys
must reach at least 0.85 of `node-layout` (the microbenchmark has no tree above the pages, so it
can promise nothing about that); if they do not, pages are used only for short suffixes.

## 2026-10-02 13:06: A/A job a1, gate 0 met

`a1` (the same as a0 with `-minprocs 8 -maxprocs 24`) ran 12:59-13:06, 7 minutes (estimated 20),
results in `origin/arm-results` (`a1/`). Every scenario was precise with its first 8 processes,
none needed more.

- **All 32 comparisons are precise** by the harness's criterion (a0: 30). The bounds of their
  intervals stay within ±1.9 points, except for `valuesBetween` (up to -2.8).
- **All 32 point estimates lie between 0.98 and 1.01** for identical code (differences -2.5% to
  +1.5%); 29 of them within ±1%.
- **`valuesBetween` is the noisy operation** on this machine: `str unique` 0.98 at 4K (-1.7%) and
  16K (-2.5%) as in a0, `u64 multi 4K` +1.5%, and the other range cells ±0.7%. The harness marks
  nine cells "resolved" although the code is identical (all of them range scans or within 1%);
  read an arm64 `valuesBetween` difference below 3%, and any other below 1.5%, as noise.
- The machine was not at rest again: the runner warned (load 6.6 in the last minute before the
  build, 7.5 at the end), and I did not ask what else was running. Precision was good anyway; a
  quieter machine may be better still, and nothing here says it would not.

Gate 0 reading: met. The M1 Pro can measure; its noise floor is about ±1% for point operations,
`churn` and `build`, and about ±3% for ranges.

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

Gate 0 reading: numerically met (0.97-1.01), with the two caveats above; `a1` above settled them.

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

**Next action:** the user decides whether gate 0 is passed. Then step 1 (PLAN.md).

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
