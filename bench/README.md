# Benchmarks

These benchmarks answer one question: which multimap should you use? They
compare this library's two implementations with what you would otherwise
write by hand, for the operations a multimap exists for.

The module is separate from the library, so the library's `go.mod` does not
inherit the benchmark's dependencies.

## Candidates

| Name | What it is | Why it is here |
|---|---|---|
| `ordered` | [`multimap.Ordered`](../ordered.go), an adaptive radix tree with value sets in its leaves | this library's default (`New` wraps it) |
| `hashed` | [`multimap.Hashed`](../hashed.go), a Go map from keys to the same value sets | this library's choice when range queries are rare |
| `btree-sets` | [`tidwall/btree`](https://github.com/tidwall/btree)`.Map[string, map[T]struct{}]` | the ordered multimap you would build by hand, on the most widely used Go B-tree |
| `map-sets` | `map[string]map[T]struct{}` | the unordered multimap you would build by hand |
| `btree-map` | [`tidwall/btree`](https://github.com/tidwall/btree)`.Map[string, T]`, one value per key | what you would use if your keys (almost) never hold more than one value: is `ordered` still a good choice then? |

The hand-written candidates own their keys, as the library does, and remove
a key once its last value is gone. Every candidate is called through its
concrete type, as you would call it. Every comparison is `ordered` against one
of the others: `hashed`, `btree-sets` and `map-sets` with the usual number of
values per key, `btree-map` with exactly one value per key (see below).

The insertions and deletions of `churn` and `build` are computed before
timing starts, by simulating the workload: a deletion is drawn only from
values that are already in. The timed loop reads the next mutation from an
array and applies it. New values are running row IDs, as a database index
stores them, so every insertion adds a value its key does not hold yet.

## What is measured

- **`valuesFor`:** iterate over all values of a random existing key. The keys come in a random order that repeats only after 262,144 lookups (`keys.ProbeLen`): the hits of a corpus of that many keys or more, else that many random picks from its keys. A sequence of just n lookups, as it was before, repeats every n lookups, and at a few thousand keys a branch predictor learns that cycle and favours code with few branches over code with many: a page search, which has no branches, measured 0.71 against leaves at 4K keys with the cycle and 0.99 without it. **Numbers of `valuesFor` from before this change are not comparable at 4K and 16K keys**; the other operations are unaffected. From 64K keys on the old cycle was too long for a predictor to learn, so those numbers should be comparable, but that has not been measured.
- **`valuesBetween`:** iterate over all values of 100 consecutive keys. `hashed` and `map-sets` have no order and scan every key, so they are compared only up to 64K keys: their cost grows linearly with the number of keys.
- **`prefix`:** iterate over all values of the keys that start with what a user searches by: the first four characters of a random key, as typed into a search field, or, for paths and URLs, the directory the key lies in. Frequent prefixes come up as often as users would search them. Text keys only; `hashed` and `map-sets` again only up to 64K keys.
- **`churn`:** add and remove values like a database index in use, as the steady-state comparison of rtcompare's [`workload`](https://pkg.go.dev/github.com/TomTonic/rtcompare/workload) package: on the filled multimap, bursts of 1-16 insertions alternate with bursts of deletions of values inserted earlier. Half of the new values go to existing keys, whose value sets grow and shrink; half go to keys that appear and disappear. A quarter of the deletions (`-permchurn 0.25`) take out one of the multimap's own, long-lived values instead and put it back later in the cycle, so that the merges and shrinks of real deletions reach the whole multimap and not only its newest part. The cycle inserts as many values as the multimap holds (`-ratio 2`) and ends in its start state, so it repeats endlessly; its first pass, where the multimap grows to its peak size, is played untimed. Timed per insertion or deletion.
- **`build`:** build the whole multimap from empty with the same bursts, as the build comparison of the same `workload.Compare`: twice as many insertions as values in the end, with the same share of long-lived values taken out and put back, until exactly the corpus is left. Every sample is a whole build, so it uses its own repeats (`-buildrepeats`, `-buildvalidation`). Only up to 64K keys, because at 1M one build takes too long for a timing sample.
- **Memory**, at 1M keys (`-suite release`) or 128K keys (`-suite dev`), or fewer where a real-world corpus holds fewer keys: retained heap per key, including the keys' copies and values but excluding the input corpus. Also the CPU time of a full GC cycle while the multimap is alive, and the heap still retained after removing every second key. Go maps do not shrink.

Each comparison runs with 4,096 and 16,384 keys, which fit in the CPU caches,
with 262,144 keys, which fit only partly, and in the release suite with
1,048,576 keys, which do not. The key kinds:

| kind | what | length |
|---|---|---|
| `u64` | random 64-bit integers, 8 bytes big endian | 8 B |
| `str` | synthetic paths such as `tenant/category/word/12345` | about 26 B |
| `uuid` | random UUIDs in lowercase hex: no shared prefixes | 36 B |
| `email` | `first.last@domain` over five domains: the shared part comes last | about 25 B |
| `url` | real host names from the Tranco list with synthetic paths shaped like real sites | about 63 B, up to 190 |
| `path` | real file paths from the packages of Debian 12 | about 65 B, up to 300 |
| `street` | real German street names from OpenStreetMap | about 14 B |
| `dirs` | the directories of the same Debian file paths, with the closing slash | about 25 B |
| `links` | the titles of Simple English Wikipedia pages that link to other articles | about 17 B, up to 255 |

`dirs` is derived from `path`'s sample: its keys are the 172,431 directories of the 600,000 paths and
its natural values the file names in them (see below). `path`, `street` and the hosts of `url` come from
[`keys/testdata`](keys/testdata/README.md), where their sources and licenses
are documented. `path` and `street` hold enough keys for 262,144 and 212,000
keys respectively, and larger scenarios are skipped; `dirs` holds enough for 86,215 and `links` for 70,000; `url` has no limit.

Values are `uint64`, and their number per key is skewed like a real index
(`-values natural`, called `multi` until 2026-10-04): 50% of keys hold 1 value, 35% hold 2-4, 12% hold 5-16 and
3% hold 17-200. Street names hold their real localities instead: 79% of the
names have one, "Hauptstr." has 5,913. Directories hold the names of the files in them: 62% of the
directories hold one file, 89% at most four, and the biggest holds 6,372 (the sample thins directories
out, so real directories hold more). Pages hold the pages they link to, the opposite shape: 29% of the
pages hold one link (the redirects), 16% hold 2-8, 38% hold 9-64 and 17% hold 65 or more, 46 links on average and
5,693 for the biggest, and the pages with 65 or more hold 77% of all links.

With `-values single-value` (before 2026-10-04: `unique`), every key holds exactly one value, like an index on a
unique column, and `ordered` is compared with `btree-map`. In `churn` and
`build`, every new value then goes to a key of its own, which appears with
the value and disappears with it; there are as many such keys as corpus
keys, so `-ratio` is at most 2 with single-value entries. The old names `multi` and `unique` are still accepted
by `-values`, and result files of before keep them.

Built with the tag `strvals`, the bench uses `string` values instead: each
value number as 16 hex digits, like a record ID (for `street`, `dirs` and `links` the real names instead: the
locality, 2 to 33 bytes, the file name, 1 to 136 bytes, and the title of the linked page, 1 to 216 bytes, whose lengths vary; the transient values of
churn and build stay hex)
(`go run -tags strvals ./cmd/bench`). Values that hold a pointer take other
paths than integers in some candidates, `ordered` among them, and the
garbage collector has to scan them. The profiles are then reported as
`natural-str` and `single-value-str`. All value strings are views into one buffer
that no candidate owns, as the key corpus is, so the memory figures count
each value's 16-byte string header but not its bytes. The timed loops add
up each value's length and first and last byte, which reads the bytes
as a caller that uses a value does (since 2026-10-04; before, they added up the
address of the bytes, which read only the string header the candidate holds,
so string-value speed figures from before are not comparable: they favoured a
candidate that holds pointers to the caller's strings over one that holds the
bytes itself). The checks compare an FNV hash of all bytes.

Built with the tag `ptrvals`, the values are pointers to records of 16 bytes
(`*rec`, one object per different value, allocated one by one), as an
application indexes its objects by key. The profiles are reported as
`natural-ptr` and `single-value-ptr`. The records are the caller's and cost every
candidate the same, so the memory figures do not count them; the timed loops
read the record's id. `strvals` and `ptrvals` are exclusive. All three builds
share one source; the `uint64` build compiles its timed loops exactly as before.

## Results

AMD Ryzen 9 7900 (12 cores, 24 threads), Linux 6.18 under WSL2 on Windows,
Go 1.27.1, rtcompare v0.7.0, 2026-09-28: the dev suite with 5-6 processes per
scenario, all four sizes and memory at 1M keys (`-suite dev -sizes
4096,16384,262144,1048576 -maxprocs 6 -memn 1048576 -memrounds 3`). The
`url` rows come from a run with the same flags on 2026-09-29
(`-keys url`, [`results/run-url.log`](results/run-url.log)), after `url`
changed to real Tranco hosts. Every
factor is how many operations the first candidate completes in the time the
second needs for one: above 1 it is faster.

352 of the 400 comparisons reached an interval across processes within ±2
percentage points or ±10% of the difference within 6 processes. The other
48, mostly `churn` from 256K keys up, were measured again on 2026-09-29,
each scenario with only those comparisons and up to 30 processes, and the 6
still open with up to 100 more (`results/logs/stab-*`). Their rows pool all
processes of these runs and replace the first ones. 398 comparisons are now
precise. The other 2 spread 14 points between processes: `btree-map` range
queries with u64 keys at 1M (38 processes, see below) and `btree-map` churn
with url keys at 256K (130 processes, 1.07× [1.05×, 1.10×]). They are marked
`*` in the tables below, with their 95% intervals below the table. The full tables are
[`results/speed-summary.md`](results/speed-summary.md) and
[`results/mem-summary.md`](results/mem-summary.md).

**`ordered` against `hashed`:** about equal for point operations with integer
keys, slower with string keys, far faster for range queries.

| operation | u64 4K | u64 1M | str 4K | str 1M |
|---|---:|---:|---:|---:|
| `valuesFor` | 1.04× | 1.01× | 0.67× | 0.50× |
| `churn` | 1.07× | 0.96× | 0.69× | 0.73× |
| `build` | 1.09× | | 0.72× | |
| `valuesBetween` | 18.9× | | 15.3× | |

**`ordered` against the hand-written candidates:** always faster than
`btree-sets`, faster than `map-sets` except for changes with string keys.

| operation | vs | u64 4K | u64 1M | str 4K | str 1M |
|---|---|---:|---:|---:|---:|
| `valuesFor` | `btree-sets` | 4.48× | 4.59× | 2.87× | 2.39× |
| `valuesBetween` | `btree-sets` | 2.76× | 3.55× | 2.21× | 2.62× |
| `churn` | `btree-sets` | 3.63× | 3.13× | 2.31× | 2.12× |
| `build` | `btree-sets` | 3.10× | | 2.23× | |
| `valuesFor` | `map-sets` | 2.54× | 2.29× | 1.53× | 1.09× |
| `valuesBetween` | `map-sets` | 20.4× | | 16.0× | |
| `churn` | `map-sets` | 1.24× | 1.09× | 0.76× | 0.81× |
| `build` | `map-sets` | 1.19× | | 0.81× | |

**`ordered` against `btree-map`, one value per key:** faster for point
operations and changes, slower for range queries.

| operation | u64 4K | u64 1M | str 4K | str 1M |
|---|---:|---:|---:|---:|
| `valuesFor` | 5.76× | 2.99× | 2.67× | 1.80× |
| `valuesBetween` | 0.47× | 0.68×* | 0.32× | 0.54× |
| `churn` | 2.83× | 2.23× | 1.62× | 1.75× |
| `build` | 2.29× | | 1.56× | |

\* `valuesBetween` u64 1M: [0.66×, 0.70×] over 38 processes of two runs, which
gave 0.69× and 0.65× on their own; its processes spread 14 points.

**Memory at 1M keys** (bytes per key; scannable: the part of the heap the GC
must scan for pointers, as `runtime/metrics` counts it, which can exceed the
live heap; GC CPU time per full cycle, minus a process that holds only the
input; medians of 3 rounds):

| values | candidate | heap u64 | heap str | scannable u64 | scannable str | GC u64 | GC str | heap after removing half, u64 | str |
|---|---|---:|---:|---:|---:|---:|---:|---:|---:|
| natural | `ordered` | 149 | 177 | 83 | 111 | 215 ms | 305 ms | 72 | 86 |
| natural | `hashed` | 177 | 196 | 101 | 101 | 218 ms | 236 ms | 115 | 121 |
| natural | `btree-sets` | 343 | 363 | 74 | 72 | 351 ms | 364 ms | 172 | 178 |
| natural | `map-sets` | 360 | 379 | 85 | 87 | 316 ms | 340 ms | 206 | 212 |
| single-value | `ordered` | 73 | 101 | 81 | 109 | 179 ms | 264 ms | 35 | 48 |
| single-value | `btree-map` | 37 | 56 | 37 | 37 | 59 ms | 76 ms | 19 | 25 |

What follows for a choice:
- **Integer or other short fixed-size keys:** `ordered` is as fast as `hashed` (0.96-1.09×) and adds range queries.
- **String keys:** `hashed` is 1.4-2.0× as fast for point operations. Take `ordered` when you need range queries or ordered iteration: a range query on `hashed` scans every key and costs 15× as much already at 4K keys.
- **Against writing it yourself:** both use about half the memory of a map or B-tree of Go sets, and `ordered` is 2.1-4.6× faster than the B-tree.
- **One value per key:** `ordered` is 1.6-5.8× faster than `btree-map` for lookups and changes, but `btree-map` answers range queries 1.5-3.1× faster, needs half the memory and a third of the GC time. The B-tree keeps each key and value in one flat array per node; `ordered` has a leaf object per key with a pointer to its key.
- **Memory after deletions:** `ordered` shrinks with its keys; the maps keep their tables.
- **GC:** with string keys `ordered` costs about 1.3× the GC time per cycle of `hashed`, with integer keys about the same; both cost less than the hand-written candidates. Its leaves and nodes are many small objects with pointers.

## How the numbers are made

**Speed:** every comparison of two candidates is one
[rtcompare](https://github.com/TomTonic/rtcompare) run. It interleaves the
candidates in A/B/B/A order, which cancels drift from throttling or background
load, and validates itself with A/A runs. Never compare absolute numbers from
separate `go test -bench` runs.

**One process is one sample.** rtcompare's interval covers only the noise
within one process. Each process also lays out its data in memory
differently, and with large pointer-heavy data that layout shifts a
comparison by several points. On 2026-09-24, repeated processes with
identical code scattered 2-5 points apart at 1M keys, while each process
reported an interval of about ±0.7 points. At 4K keys they agreed within
their intervals. Longer runs did not help: 1.5x samples and batches narrowed
each interval but not the scatter
([rtcompare#109](https://github.com/TomTonic/rtcompare/issues/109)).

The driver `cmd/bench` therefore runs every scenario (key kind × number of
keys) in separate processes, through rtcompare's
[`multiproc`](https://pkg.go.dev/github.com/TomTonic/rtcompare/multiproc)
package (rtcompare v0.8.0 or later):
- Each process perturbs its heap from its own seed before it builds anything (`rtcompare.PerturbHeap`) and builds the candidates in its own order: two candidates alternate from process to process, more are shuffled. The processes thus sample different layouts instead of repeating one biased layout, and whichever candidate is built last has no lasting advantage. The seed of a process also seeds each comparison's resampling, so a run can be repeated (`Options.Seed` of the driver is fixed).
- Each process counts as one observation. `rtcompare.CombineStaged` pools them: the mean difference, a 95% t-interval across processes, the spread between processes and how far it exceeds a single process's interval.
- The run is sized once, from its first stage (Stein's two-stage procedure): the first `-minprocs` processes (6 in the release suite; the dev suite takes 4 and accepts wider intervals) show how much the processes scatter, from that the run works out how many processes every comparison needs for its interval to be within ±2 percentage points or within ±10% of the difference itself, and runs exactly that many, up to `-maxprocs` (rtcompare's default: 40 serial, 10 waves parallel). It does not look at the intervals again and stop as soon as they are narrow enough: that stops preferably where the processes happened to agree, and such intervals covered the truth only 92-94% of the time instead of 95%. Five or six processes are plenty in the cache: a spread of 0.1-0.4 points gives about ±0.5. At 1M keys a spread of 2-5 points needs 10-40 for ±2. A fixed number would either waste the night on 4K or stop too early at 1M. A run that would need more than `-maxprocs` says so in the log and in `speed-summary.md` (precise: no).
- `speed.jsonl` keeps what pooling needs: the size of the first stage and the A/A differences of every process, so the summary pools the rows exactly as the run did, and the driver warns if it does not.

**Two regimes.** By default the processes run one after another, each with the
machine to itself: the serial regime says how the candidates compare on an idle
machine. With `-parallel N` they run N at a time, in waves: the parallel
regime. The processes then share the last-level cache, the memory bandwidth and
the clock headroom, much as a program shares them with its neighbours in
production, and the data falls out of the cache at smaller sizes. That is a
different question, not a faster answer to the same one. Out of the cache it
is also the way to precision in reasonable time: the scatter between processes
dominates, and only more processes narrow it, so N at a time buy N times as
many per hour. A and B still run interleaved within each process, so the
neighbours widen the noise but favour neither. The two regimes' tables are
never pooled or compared with each other: the rows carry `parallel`,
`run.json` and the summary say which regime ran. Keep `-parallel` at or below
the physical cores (12 on the Ryzen 9 7900) and what fits in memory (see below): two processes on the two threads
of one core disturb each other far more than neighbours on other cores. Each
child then runs with `GOMAXPROCS` = CPUs / N, at least 2, so that one child's
garbage collector cannot take cores from its neighbours' measurements.

The measurement plan follows from that: the **serial** `dev` and `release`
suites answer "does a small index get worse, does the change pay on an idle
machine", and the **parallel** suite answers the same for the out-of-cache
sizes under load, where the serial suite would need days for a precise answer.
Both are reported side by side, never merged.

Before rtcompare v0.7.0, every comparison from about 64K keys up favoured
candidate B, the other candidate, by 5-20%: rtcompare validated B last, and B
started the measurement with the cache full of its own data
([rtcompare#111](https://github.com/TomTonic/rtcompare/issues/111)). Results
measured before 2026-09-28 carry that bias against `ordered` at large sizes.

Since rtcompare v0.8.0 (2026-09-30) the numbers are not comparable with
earlier ones: `NewDPRNG` draws other sequences, so every synthetic key set and
every sample of the real corpora differs; `Compare` takes its difference from
the ratios of pairs of neighbouring batches, which cancels disturbance that hits
both candidates; the pooled interval is Stein's, and its noise floor is the
systematic bias of the A/A runs; and `churn` and `build` delete and put back
long-lived values too. Only comparisons within one run, and runs made with the
same version of the driver, can be read against each other.

**Memory** is not an rtcompare comparison. GC cost depends on the whole live
heap, so each candidate is measured alone in its own process, together with a
baseline process that holds only the input. Five rounds run in a new random
order each, and the tables show medians.

## Running

```sh
go run ./cmd/bench                                  # the dev suite, for frequent runs while trying a change
go run ./cmd/bench -suite release                   # the complete suite, before a release
go run ./cmd/bench -keys str,street -sizes 16384    # a quicker subset
go run ./cmd/bench -out /tmp/smoke -repeats 21 -validation 2 -minprocs 4 -maxprocs 4 -memn 16384 -memrounds 1
go run ./cmd/summarize results/speed.jsonl          # pool the raw results again
```

The three suites:
- **`dev`** (the default) leaves out 1M keys and measures 4K, 16K and 256K instead: small indexes are the common case and must never get worse, and the trend up to 256K shows where larger ones are headed. It trades precision for time: a first stage of 4 processes, at most 8, 2 A/A runs and 41 samples per rtcompare comparison, about 1.5 hours on a Ryzen 9 7900.
- **`release`** adds 1M keys and measures at full precision, one process after another: a first stage of 6 processes, at most 40, rtcompare's defaults.
- **`parallel`** measures the out-of-cache sizes, 256K and 1M keys, in the parallel regime: 8 processes at a time, a first stage of one wave, at most 10 waves (80 processes), rtcompare's defaults, no memory measurements. See "Memory and machines" below for what 8 processes need.

Flags given on the command line override the suite's presets. `-maxprocs 0`
takes rtcompare's default.

### Memory and machines

Every process builds all candidates of its scenario and, for `churn`, the
workload's structures next to them, so a process at 1M keys is large:
a process of `url` keys with several values each, 1M keys, needs 4.4 GB
with `uint64` values and 5.1 GB with string values against a baseline, and up
to 8 GB with all four candidates of the natural profile (measured as the
largest resident set of a wave of four processes, `-ops valuesFor,churn`).
Most of it is the garbage collector's headroom over two live structures and
the streams. Eight processes at 1M keys therefore need 35-41 GB, twelve 53-61
GB, which leaves too little of 64 GB. WSL2 offers only half of the machine's
memory by default (30 GB here, see `memory=` in `.wslconfig`), enough for four
processes at 1M keys; `-parallel 8` at 256K keys needs 11 GB. The driver logs
the free memory and the share per process when a parallel run starts.

Windows can also run the driver natively, which takes the WSL2 virtual
machine out of memory-bound comparisons and gives all of the machine's memory
to the run: build it with `GOOS=windows go build -o bench.exe
./cmd/bench` (add `-tags baseline,strvals` as needed) and start it from
PowerShell; the driver, its `awake` package and rtcompare's clock all support
Windows. Results from another operating system or machine are recorded in
`run.json` (`os`, `cpu`) and are not comparable with these either.

The driver starts its own binary for every process and writes to `results/`:
- `speed.jsonl` and `mem.jsonl`: one line per comparison per process;
- `speed-summary.md` and `mem-summary.md`: the pooled tables;
- `run.json`: settings, machine and time;
- `logs/`: rtcompare's full report for every process.

While it runs, the driver keeps the machine from sleeping (package
[`awake`](awake/awake.go): `caffeinate` on macOS, `systemd-inhibit` on Linux,
`SetThreadExecutionState` on Windows) and warns in the log about every process
that the machine slept through anyway (rtcompare reports it per
comparison). Under WSL2, `systemd-inhibit` holds only the Linux VM awake, not
Windows: switch off sleep in the Windows power plan for the run. The run of
2026-09-28 above took 4 h 49 min on the Ryzen and did not sleep.

To measure a change against an earlier commit, copy that commit's library
into the bench as the candidate `baseline` and compare Ordered with it alone,
in the same process and interleaved like every other pair:

```sh
go run ./cmd/mkbaseline -ref main                   # writes ./baseline (ignored by git)
go run -tags baseline ./cmd/bench -vs baseline      # Ordered against main's Ordered
```

`-vs` limits the candidates Ordered is compared with; `run.json` records the
commit of the baseline.

`-repeats`, `-loopscale` and `-validation` pass through to rtcompare. The
defaults are rtcompare's. Lower values are for smoke tests only.

## Earlier benchmarks

The benchmarks from the design phase are in [`_archive/`](_archive/README.md):
prototypes of the tree, comparisons with
[plar/go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree)
and [plar/go-hot-trie](https://github.com/plar/go-hot-trie), node sizes, the
value container, and the range scan that touches children ahead. They no
longer build from there. The last commit in which they ran is `57a4fa9`.
