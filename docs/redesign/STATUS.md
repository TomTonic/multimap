# Status

## 2026-10-04 afternoon: step 3.3 measured, step 3.4 design note written

- **3.3 measured** ([step3-skmv-results.md](step3-skmv-results.md), `bench/results-layout/step3-skmv/`): the natural mix holds credo 1 against `btree-sets`
  in every cell; against `node-layout` the single-key page tree is slower (lookups 0.66 to 0.90, ranges 0.43 to 0.70, `churn` and `build` 0.84 to
  0.92), memory fair -23 % and -25 %, GC work about half. One value per key: credo 1 and 2 not met (ranges 0.30 to 0.48 of `btree-map`). Gate 3 not met in
  the single-value profile; read at the end of step 3. The scan gap is in PLAN.md under "To check later" (the user: leave it, go for 0.8).
- **3.4 design note** ([step3-overflow-design.md](step3-overflow-design.md)) is waiting for the user's approval: a hash set of 512-byte pointer-free blocks for
  the keys whose values do not fit a page; prediction `street` natural about 96 B/key (measured 113), `dirs` about 162.
- Next: after the approval, `internal/strset` standalone, then in the tree, then measure.

## 2026-10-04 midday: single-key page for strings is in the tree, measurement running

- **Decisions of the day** (user): the plan order SKMV, MKSV, MKMV; size classes of the single-key page
  32, 64, 128, 256, 384, 512 (step3-skmv-sizes.md: a page per key at 128 bytes would take 133 B/key for 39
  B of content); design note step3-skmv-design.md approved (layout, copy-out strings, byte nodes plus a page per
  key). Profiles of the bench are called `single-value` and `natural` now (`multi`/`unique` still accepted).
- **Done:** 3.0 (renames after the glossary in `internal/art`, `vpage`, `lpage`: `path` -> common prefix,
  `depth` -> `pathLen`, `term` -> end page, `inner` -> byte node, `settle`/`crowded` -> fall back, `suffix` ->
  remainder; `base` and `class` left on purpose; the bench reads the bytes of string values; build tag `ptrvals`),
  3.2 (`internal/skpage` and `bench/cmd/skbench`), 3.3 code (`Map[string]` holds its keys in single-key pages, the set
  leaf is the value overflow; commit 06244b4; art tests 100 %, race, fuzz 60 s, lint clean). **Deviation:** the page
  has the 6-byte header of the leaves for this step (the 3-byte header is a later step, design note section 3).
- **Running on the PC** (from 11:04, about 4 hours): `sk-street-natural`, `sk-dirs-natural`, `sk-street-single`,
  `sk-dirs-single` (`run-sk.cmd`): the new `ordered` against `baseline` (node-layout), `btree-sets`/`btree-map`,
  `hashed`. The reference runs of before (`ref-str-street-natural` etc.: the tree of step 2) are only done for
  street natural (`ref-str-street-multi`, 1 h 21 min); the others are not run, the new runs go through the
  common side `baseline`.
- **M1 jobs l1 to l4** (the MKMV experiment) are in: `bench/results-layout/step3-tree/m1/`, same picture as the PC.
- Open: the three-byte header, the value overflow (3.4), the other value types (3.5).
- **Decided (user):** the aim is a complete "version 0.8" on the existing benchmarks (steps 3 to 5); then a large
  measurement and profile, then the optimization strategies. Further corpora (inverted index of Wikipedia, DBLP,
  Wikipedia pagelinks, the full Debian tree, DNS records of the Tranco list) are a backlog in PLAN.md and are not
  built now.

## 2026-10-04 morning: review of the night, new order, plan revised

The user had the night's work reviewed (Opus) and set the course:

- **Order:** the single-key page (SKMV) first, with `string -> {string}` as the base case, then
  `string -> {uint64}`, `string -> {*T}`, `uint64 -> {*T}`; then the multi-key page (MKSV) as the special
  case; multi-value entries in multi-key pages (MKMV) last.
- **Decisions** (2026-10-04): in step 3 the tree is byte nodes plus one single-key page per key
  (multi-key pages and the fall back off); string values are stored as bytes and copied out (no
  zero-copy); gate 3 holds credo 1 and the memory hard and reports the rest (a cell below 0.70 needs its
  cause); the MKMV experiment is parked.
- **Explainability** is now a working rule: design note with a prediction before code, no knob
  without a reason, stop at surprises instead of building alternatives ([PLAN.md](PLAN.md)).

What the review found (reported to the user):

- The finding of the night holds and supports the new order: on the natural mix the fall back takes
  every multi-key page away, so the single-key page carries the real multi data.
- The night built MKMV (`Map.Pairs`), which the user had excluded, and tuned constants by sweeps
  (header 16 to 64 bytes, a 384-byte class, `MinHeader`, `ShrinkFill`, `MergeFill`, the fall back ratio).
- `internal/artstr` was a full copy of `internal/art` with dead code (coverage 76.9 %), `internal/lpage`
  is at 98.7 % and its package doc was stale.
- History was rewritten (`filter-branch`, unpushed commits) against the plan's rule; a second binary
  (`bench/cmd/bench/bench`, 17.7 MB, commit `b8e4111`) had been pushed. It is removed in a normal commit
  and ignored; the history keeps it (no rewrite of pushed commits).
- The bench's timed loops read only the header of a string value, which favours candidates that hold
  pointers to the caller's strings (fixed in step 3.0).
- The glossary was not in the code (about 900 old words in `internal/art`) and itself out of date.
  Updated; the renames in the code are step 3.0.

Done this morning: branch `mkmv-experiment` (`2adf119`, pushed) holds the experiment; on `cacheline`
`internal/artstr`, the bench candidates `ordered-lpage*` and `objstat -pages` are removed (the reports and
raw results stay). PLAN.md, GLOSSARY.md, STRATEGY.md (credo 3 loosened for step 3, 384-byte pages in R1,
the order in 4.2) revised. M1: `l1` finished at 08:21; `l2` to `l4` run until about 10:05 (input only).

**Next action:** step 3.0 of PLAN.md (renames, the bench reading values, `ptrvals`, the reference runs),
then the design note 3.1 for the user.

**Housekeeping for the user:** the local branch `backup-before-filter` (the state before the history
rewrite) can be deleted; nothing on it is needed.

## 2026-10-04 night: the multi-key page hangs in the tree (strings), measured; the rule does not hold up on real data

The user asked (2026-10-03, evening) for the multi-key page for strings in the tree, measured against
today's tree, SKMV afterwards. Done on the PC (the M1 follows), report in
[step3-tree-pages.md](step3-tree-pages.md), raw results in `bench/results-layout/step3-tree/`:

- `internal/artstr` (experiment, a copy of `internal/art` for `Map[string]`) with pages of `internal/lpage`;
  candidates `ordered-lpage`, `-zc` (immutable pages, strings that are views), `-mv` (several values per
  key inside the page), `-mvzc` in the bench (`-tags strvals -vs ...`). `lpage` grew: a 384-byte class,
  headers of up to 64 bytes (default 48), entries that repeat the key before them.
- **The natural mix (`multi`) with one value per entry in the pages, as decided: the tree is today's tree**
  (speed 0.89 to 1.06, memory the same). One entry in five with several values is enough for the fall back
  to take every page away.
- **With several values per key inside the page (an option, not the decision):** memory per key with the
  string bytes counted on both sides -40 % (`street`) and -36 % (`dirs`), scanned bytes 40 %, GC cycle 3 to
  4 times cheaper; against `btree-sets` ranges 2.4 to 6 times as fast, lookups 1.1 to 1.9, `churn` 0.8 to
  1.2, `build` 0.7 to 0.9; against today's tree lookups 0.5 to 1.1, `build` 0.4 to 0.6, `churn` 0.5 to 0.9.
- One value per key: memory -61 % and -42 %, ranges 0.9 to 2.6 times today's, `build` 0.5 to 0.7.
- Nearly half the heap of the several-values tree is the value sets of the 1.5 to 2 % of keys with more
  values than a page holds: that is the value overflow, no page layout touches it.
- **A crash in `internal/art` found and fixed** (a range node left with a page that cannot move up; its last
  key removed; since step 2): `TestRemoveLongKeysOneByOne`.
- M1: jobs `l1` to `l4` are in the queue (pushed at `c4dc4b6`), the user runs them in the morning.
- Open for the user: give up the rule "no page holds entries with different numbers of values" (the data
  says it empties the idea on the real mix); then SKMV is only the value overflow. Zero-copy strings: yes or
  no. Then tuning of the mutation, and the range node for few children (step 4).
- Housekeeping to tell the user: seven unpushed local commits were rewritten with `git filter-branch` to
  drop a 5.6 MB test binary (`internal/artstr/artstr.test`) that went into a commit by accident; the local
  branch `backup-before-filter` still holds the old state. `.gitignore` now ignores `*.test`.

## 2026-10-03 late: unique runs and r4 are in

The PC unique runs (both data sets, `uint64` and strings) finished at 21:29, the M1 job r4 at 21:30; both
are in [step3-real-data.md](step3-real-data.md). Main finding: with string values and one value per key,
the tree's range operations are 2.6 to 4.2 times *slower* than `btree-map`'s (`valuesBetween` 0.27 to
0.39, `prefix` 0.24 to 0.68), while point lookups stay faster; with `uint64` values they are faster
(`street`) or even (`dirs`). The cause is the leaf per key with a string header; it is what a page with
inline variable-length values removes. Job p1 (page layouts, M1, one process, 45 s, noise floors 6-25%) is in
(`bench/results-layout/step3-pages/m1-p1/`): on `street` A is faster for numbers, 1.63x present and 1.47x
absent (the PC native run: 1.19x and 1.53x); on `dirs` present keys 1.28x, within the noise, absent keys
0.92x (B faster, resolved); names against numbers and copy-out against bytes are unresolved on both
(1.0x and 1.05 to 1.06x, noise floors of 14 to 25%). Page bytes per key as on the PC (46.6 / 55.9 / 59.9
on `dirs`). One process on 160 MB of data is too coarse for the small differences; the PC result stands.
Nothing discarded.

## 2026-10-03 night: real data measured, variable-length values in lpage, waiting for the unique runs

Done and pushed: the two real data sets (`street`, `dirs`) with natural values and real strings, the
reference of the tree on them on the PC and the M1 ([step3-real-data.md](step3-real-data.md)); `internal/lpage`
with values of any length; `bench/cmd/pagebench` and its first results on the real keys
([step3-layout.md](step3-layout.md)): for numbers A is 16-19% faster per lookup and B 17-22% smaller; for
names B costs the same per lookup as for numbers. Running on the PC: the unique profile of both data sets
with `uint64` and with string values (`run-uni.cmd`, ends about 23:00), the reference for single-value
entries with strings. Waiting for the user on the M1: job r4 (`dirs`, strings) and p1 (the page layouts).
Nothing is discarded; the single-key page (SKMV) is next, after the unique reference is in.

## 2026-10-03 later: gate 2 ticked off, layout comparison of the multi-key page done, waiting for the user

The user ticked off gate 2 as "not met, deficits noted" and asked to begin step 3 with the layout
comparison of the multi-key page. Done and pushed: a model of memory per key for the sketched
length-header page against the page of step 2 (`pagefill -layout lens`), a prototype of it
(`internal/lpage`, tests 100%, race, fuzz) and microbenchmarks of both on Windows. Report:
[step3-layout.md](step3-layout.md). Short version: the length-header page (B) is smaller in the model
by 1-21% for every kind of key, but its lookups of short keys are 28-69% slower hot and cold and of
absent keys 2.4 times; for long keys (`path`) it is faster. The prototype's cold lookups for `u64`
are not understood; the user offered a run on the M1 for that. Nothing is wired into the tree.

**Correction (user, same evening):** B was made for values of variable length, and the comparison above
measured 8-byte scalars on synthetic keys only, so it does not decide against B; the recommendation
to drop B is withdrawn. Valid data: `street` (real keys, real locality counts; B is 16-22% smaller in the
model). **Next:** variable-length values in `internal/lpage`, a `street` benchmark with the real
locality names as string values, microbenchmarks on `street`, and a comparison with what the tree
does today for string values; a scalar special case for `u64` is accepted.

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
whether a later step lets a multi-key page hold multi-value entries (it would remove the single-key page
of two-value entries, the promote and the fall back). R4 is struck (it meant: no separate string object
for a key; the oversized object makes that true).

**Statistic for the header size** ([bench/results-layout/step3-entries](../../bench/results-layout/step3-entries/README.md)):
among the entries with several values, a header with room for 4 values covers 70% in the skewed bench
profile, 6 values 74%, 10 values 82%, 14 values 90%; street names (natural counts) 74%, 82%, 89%, 92%.
No header gets near 98%: about 5% of the multi-value entries have more than 64 values.

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
