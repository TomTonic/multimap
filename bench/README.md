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

- **`valuesFor`:** iterate over all values of a random existing key.
- **`valuesBetween`:** iterate over all values of 100 consecutive keys. `hashed` and `map-sets` have no order and scan every key, so they are compared only up to 64K keys: their cost grows linearly with the number of keys.
- **`churn`:** add and remove values like a database index in use. On the filled multimap, bursts of 1-16 insertions alternate with bursts of deletions of values inserted earlier. Half of the new values go to existing keys, whose value sets grow and shrink; half go to keys that appear and disappear. The cycle inserts as many values as the multimap holds (`-ratio 2`) and ends in its start state, so it repeats endlessly. Timed per insertion or deletion.
- **`build`:** build the whole multimap from empty with the same bursts: twice as many insertions as values in the end, until exactly the corpus is left. Only up to 64K keys, because at 1M one build takes too long for a timing sample.
- **Memory**, at 1M keys: retained heap per key, including the keys' copies and values but excluding the input corpus. Also the CPU time of a full GC cycle while the multimap is alive, and the heap still retained after removing every second key. Go maps do not shrink.

Each comparison runs with 4,096 keys, which fit in the CPU caches, and with
1,048,576 keys, which do not. Keys are either `u64` (random 64-bit integers,
8 bytes big endian) or `str` (path-like strings such as `tenant/category/word/12345`,
about 26 bytes). Values are `uint64`, and their number per key is skewed like
a real index (`-values multi`): 50% of keys hold 1 value, 35% hold 2-4, 12%
hold 5-16 and 3% hold 17-200.

With `-values unique`, every key holds exactly one value, like an index on a
unique column, and `ordered` is compared with `btree-map`. In `churn` and
`build`, every new value then goes to a key of its own, which appears with
the value and disappears with it: the simulation hands each new value a free
key from a pool and takes the key back when the value is deleted.

## Results

AMD Ryzen 9 7900 (12 cores, 24 threads), Linux 6.18 under WSL2 on Windows,
Go 1.27.1, 2026-09-25. Every factor is how many operations the first
candidate completes in the time the second needs for one: above 1 it is
faster. All intervals across processes are within ±2 percentage points or
±10% of the difference, except for the two marked `*`: `u64` `churn` at 1M
keys spreads 9-12 points between processes, so ±2 would take 70-130
processes, and the driver stopped at 40 (`-continue`, see below). Their 95%
intervals are given below the tables. The full tables
are [`results/speed-summary.md`](results/speed-summary.md) and
[`results/mem-summary.md`](results/mem-summary.md).

**`ordered` against `hashed`:** about equal for point operations with integer
keys, slower with string keys, far faster for range queries.

| operation | u64 4K | u64 1M | str 4K | str 1M |
|---|---:|---:|---:|---:|
| `valuesFor` | 0.98× | 0.97× | 0.65× | 0.47× |
| `churn` | 1.08× | 1.06×* | 0.67× | 0.64× |
| `build` | 1.11× | | 0.72× | |
| `valuesBetween` | 18.4× | | 15.0× | |

\* `churn` u64 1M: [1.00×, 1.08×].

**`ordered` against the hand-written candidates:** always faster than
`btree-sets`, faster than `map-sets` except for changes with string keys.

| operation | vs | u64 4K | u64 1M | str 4K | str 1M |
|---|---|---:|---:|---:|---:|
| `valuesFor` | `btree-sets` | 4.41× | 4.26× | 2.85× | 2.25× |
| `valuesBetween` | `btree-sets` | 2.61× | 3.40× | 2.13× | 2.52× |
| `churn` | `btree-sets` | 3.73× | 2.86× | 2.32× | 1.98× |
| `build` | `btree-sets` | 3.55× | | 2.40× | |
| `valuesFor` | `map-sets` | 2.50× | 2.15× | 1.54× | 0.97× |
| `valuesBetween` | `map-sets` | 19.7× | | 15.6× | |
| `churn` | `map-sets` | 1.20× | 1.12×* | 0.75× | 0.73× |
| `build` | `map-sets` | 1.28× | | 0.83× | |

\* `churn` u64 1M: [1.05×, 1.12×].

**`ordered` against `btree-map`, one value per key:** faster for point
operations and changes, slower for range queries.

| operation | u64 4K | u64 1M | str 4K | str 1M |
|---|---:|---:|---:|---:|
| `valuesFor` | 5.63× | 2.94× | 2.65× | 1.71× |
| `valuesBetween` | 0.42× | 0.64× | 0.28× | 0.52× |
| `churn` | 2.82× | 1.76× | 1.55× | 1.43× |
| `build` | 2.36× | | 1.54× | |

**Memory at 1M keys** (bytes per key; scannable: the part of the heap the GC
must scan for pointers, as `runtime/metrics` counts it, which can exceed the
live heap; GC CPU time per full cycle, minus a process that holds only the
input):

| values | candidate | heap u64 | heap str | scannable u64 | scannable str | GC u64 | GC str | heap after removing half, u64 | str |
|---|---|---:|---:|---:|---:|---:|---:|---:|---:|
| multi | `ordered` | 149 | 177 | 83 | 111 | 238 ms | 336 ms | 72 | 86 |
| multi | `hashed` | 177 | 196 | 101 | 101 | 235 ms | 255 ms | 115 | 121 |
| multi | `btree-sets` | 343 | 363 | 73 | 73 | 384 ms | 400 ms | 172 | 178 |
| multi | `map-sets` | 360 | 379 | 85 | 86 | 347 ms | 367 ms | 206 | 212 |
| unique | `ordered` | 73 | 101 | 81 | 109 | 202 ms | 288 ms | 35 | 48 |
| unique | `btree-map` | 37 | 56 | 37 | 37 | 64 ms | 84 ms | 19 | 25 |

What follows for a choice:
- **Integer or other short fixed-size keys:** `ordered` is as fast as `hashed` (0.97-1.11×) and adds range queries.
- **String keys:** `hashed` is 1.4-2.1× as fast for point operations. Take `ordered` when you need range queries or ordered iteration: a range query on `hashed` scans every key and costs 15× as much already at 4K keys.
- **Against writing it yourself:** both use about half the memory of a map or B-tree of Go sets, and `ordered` is 2.0-4.4× faster than the B-tree.
- **One value per key:** `ordered` is 1.4-5.6× faster than `btree-map` for lookups and changes, but `btree-map` answers range queries 1.6-3.6× faster, needs half the memory and a third of the GC time. The B-tree keeps each key and value in one flat array per node; `ordered` has a leaf object per key with a pointer to its key.
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
keys) in separate processes:
- Each process gets its own `-layoutseed`. Package [`layout`](layout/layout.go) then puts random spacers between the candidates and builds them in a random order, so the processes sample different layouts instead of repeating one biased layout.
- Each process counts as one observation. Package [`stats`](stats/stats.go) reports the median difference and a 95% t-interval across processes.
- Five processes are the minimum. After that, the driver adds processes until every comparison's interval is within ±2 percentage points or within ±10% of the difference itself, up to 20. Five are plenty in the cache: a spread of 0.1-0.4 points gives about ±0.5. At 1M keys a spread of 2-5 points gives ±2.5-6 with five processes and needs 10-20 for ±2. A fixed number would either waste the night on 4K or stop too early at 1M.
- Fewer than five would estimate the spread from too few processes, and a run that happened to scatter little would stop early by luck.
- Where 20 were not enough, `-continue` adds processes to an existing `results/` under a higher `-maxprocs`: every scenario resumes after its last process, and scenarios that are already precise are skipped. The run of 2026-09-25 went on this way at 1M keys with several values: `str` became precise after 27 processes, `u64` stopped at 40.

**Memory** is not an rtcompare comparison. GC cost depends on the whole live
heap, so each candidate is measured alone in its own process, together with a
baseline process that holds only the input. Five rounds run in a new random
order each, and the tables show medians.

## Running

```sh
go run ./cmd/bench                                  # the full suite: about 3 hours on an M1 Pro
go run ./cmd/bench -sizes 4096 -skipmem             # a quicker subset
go run ./cmd/bench -continue -skipmem -maxprocs 40  # more processes where 20 were not enough
go run ./cmd/bench -out /tmp/smoke -repeats 21 -validation 2 -minprocs 2 -maxprocs 2 -memn 16384 -memrounds 1
go run ./cmd/summarize results/speed.jsonl          # pool the raw results again
```

The driver starts its own binary for every process and writes to `results/`:
- `speed.jsonl` and `mem.jsonl`: one line per comparison per process;
- `speed-summary.md` and `mem-summary.md`: the pooled tables;
- `run.json`: settings, machine and time;
- `logs/`: rtcompare's full report for every process.

While it runs, the driver keeps the machine from sleeping (package
[`awake`](awake/awake.go): `caffeinate` on macOS, `systemd-inhibit` on Linux,
`SetThreadExecutionState` on Windows) and warns in the log about every process
that the machine slept through anyway. Under WSL2, `systemd-inhibit` holds
only the Linux VM awake, not Windows: switch off sleep in the Windows power
plan for the run. The run of 2026-09-25 took 3 h 38 min on the Ryzen above,
plus 49 min to continue it, and did not sleep.

`-repeats`, `-loopscale` and `-validation` pass through to rtcompare. The
defaults are rtcompare's. Lower values are for smoke tests only.

## Earlier benchmarks

The benchmarks from the design phase are in [`_archive/`](_archive/README.md):
prototypes of the tree, comparisons with
[plar/go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree)
and [plar/go-hot-trie](https://github.com/plar/go-hot-trie), node sizes, the
value container, and the range scan that touches children ahead. They no
longer build from there. The last commit in which they ran is `57a4fa9`.
