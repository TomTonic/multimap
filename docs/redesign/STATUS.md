# Status

## 2026-10-03: vocabulary agreed, single-key page decided, step 2 gate still open

Worked out with the user: [GLOSSARY.md](GLOSSARY.md) (new words and the old ones they replace) and a
redesigned step 3 in [PLAN.md](PLAN.md). The user's decisions:

- A page may hold **one multi-value entry** (single-key page, "SKMV"); it replaces the flat, typed and
  set leaves. No page holds entries with different numbers of values for now.
- A remainder too long for the page stays inline in an **oversized object**, exempt from R1: no pointer
  to a key any more, one random cache miss less.
- Values are byte strings in the layouts under design, at most 255 bytes each while one length byte is
  used (a longer value needs an escape, open); fixed-size values get a specialized variant later. The
  number of values per entry is "what fits", not capped by the header.
- The common prefix of a page should, if possible, decide a mismatch within its first 64 or 128 bytes;
  the threshold is found by measuring. The entry count per multi-key page is not fixed: 3, 7, 11, 15 in
  the first sketch, 2, 6, 10 if another header byte is needed, or other with four size classes.
- Words: *path*, *remainder*, *common prefix*, *byte node*, *range node*, *end page*; `leaf` no longer
  names an object.

**Still open:** the decision on gate 2 (the four options of step2-results.md); the layout of the
multi-key page (sketch against `internal/vpage`); whether the byte node is needed once step 3 is done;
whether a later step lets a multi-key page hold multi-value entries (it would remove the per-key object,
the promote and the fall back). R4 is violated knowingly for multi-value entries until then.

## 2026-10-02 night: step 2 done, gate 2 not met, waiting for the user's decision

The pages of `internal/vpage` hold every key with one value in the tree (`internal/art`), whatever its
length; design in [step2-design.md](step2-design.md), results and the reading against gate 2 in
[step2-results.md](step2-results.md), raw files in `bench/results-layout/step2-pages/`. Tests: 100%
coverage in `art` and `vpage`, `go test -race ./internal/...` and the fuzz test (`FuzzOperations`, 60 s)
pass, `golangci-lint` is clean.

**What the pages bring** against `node-pages` (and `btree-map`): string ranges and prefix queries 2.5-5.6
times faster, memory per key -4..-46% for five of seven kinds and below `btree-map` for all seven, 2-9
times less for the garbage collector to scan, point lookups of string keys at 256K 8-29% faster, multi
profile neutral.

**What they cost:** at 4K-16K keys point lookups 5-19% slower, `build` 3-30% slower, `churn` of email
and uuid 10-20% slower, `u64` 5-14% slower in every operation; `url` and `path` are below `btree-map`
at 4K-16K in ranges, `churn` and `build`. Gate 2 as written (no cell below 0.85, `u64` unchanged,
credo 1 and 2 for every kind) is not met.

**Decision for the user:** the four options and my recommendation (keep the old integer page next to the
new general page, accept the rest as the price of the design, and take the routing layer of step 4 on
next) are at the end of step2-results.md. Nothing in steps 3-6 starts before the decision.

## 2026-10-02 evening: step 1 done, page prefix built, waiting for the go for step 2

The page prototype (`internal/vpage`, 100% coverage, fuzzed, race and lint clean) and its
experiments are built and described in [step1-results.md](step1-results.md). It has a directory
(the user's idea of a FAT) in the first line or two of the page: a tag byte per key, in general
pages also the suffix length and the tail's offset. Every lookup takes **two dependent rounds**
(directory; then head, value and tail together), against 3 to 4 before. In the microbenchmark
(Ryzen, diagnosis only) cold lookups take `u64` 148 ns, `uuid` 212, `path` 247 (hot 6-12 ns).

**User's decisions (2026-10-02):** the rule R5 is restated as "at most two rounds of cache-line loads"
(STRATEGY.md, with a precise definition of a round); the page prefix may be tried, with the worry that
it costs too much in `churn`; commit and push at every valuable point without asking.

**The page prefix** (the bytes all keys of a page share, once, in the first line; chosen by a
rebuild if it saves 16 bytes or more) is built. Against the leaves of `node-pages` the pages now
need: `street` -42..-48%, `str` -23..-36%, `email` -17..-23%, `path` -11..-19%, `url` -8..-12%,
`uuid` -2%, but `u64` +2..+11% (a byte a key for the directory). **Churn** in the page model with and
without the prefix: -6% to +8%, build -3% to +7%, noise about 5%; the prefix changes in at most 12 of
1000 operations. **Lookup:** a page without prefix is as fast as before (faster: a one-load head
word), a page with one costs +2.7 ns hot (`path` 11.9 to 14.6), nothing cold, and absent keys get
16% faster. The first versions of the prefix cost more (30-40% cold for `u64`, 5-20% in churn); the
reasons and the repairs are in step1-results.md ("The page prefix").

Gate 1 as first written was not met for `u64` (memory) and, literally, for the lines of general
pages; as restated (rounds) and with the prefix it is met for every kind except `u64`'s memory,
which is the price of 40% faster cold lookups. **My recommendation:** go to step 2.

**Step 2 gets a lookup gate** (in plain words): the microbenchmark has no tree above the pages
and loops over pages that are all in order, so it can promise nothing about a lookup in the tree.
After the pages are wired in, point lookups of string keys at 16K-64K keys, measured against
`node-layout` interleaved as always, must reach at least 0.85, else pages are used only for short
suffixes. (0.85 is the credo's limit for any cell.)

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
