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

The hand-written candidates own their keys, as the library does, and remove
a key once its last value is gone. Every candidate is called through its
concrete type, as you would call it.

## What is measured

- **`valuesFor`:** iterate over all values of a random existing key.
- **`valuesBetween`:** iterate over all values of 100 consecutive keys. `map-sets` has no order. `hashed` scans every key, so it is compared only up to 64K keys: its cost grows linearly with the number of keys.
- **`addRemove`:** add a value that a random existing key does not hold, then remove it again.
- **`build`:** build the whole multimap from scratch. Only up to 64K keys, because at 1M one build takes too long for a timing sample.
- **Memory**, at 1M keys: retained heap per key, including the keys' copies and values but excluding the input corpus. Also the CPU time of a full GC cycle while the multimap is alive, and the heap still retained after removing every second key. Go maps do not shrink.

Each comparison runs with 4,096 keys, which fit in the CPU caches, and with
1,048,576 keys, which do not. Keys are either `u64` (random 64-bit integers,
8 bytes big endian) or `str` (path-like strings such as `tenant/category/word/12345`,
about 26 bytes). Values are `uint64`, and their number per key is skewed like
a real index: 50% of keys hold 1 value, 35% hold 2-4, 12% hold 5-16 and 3%
hold 17-200.

## Results

Pending: the first run of this suite is planned for the night of 2026-09-24.
`cmd/bench` writes `results/speed-summary.md` and `results/mem-summary.md`;
the tables will be summarized here.

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

**Memory** is not an rtcompare comparison. GC cost depends on the whole live
heap, so each candidate is measured alone in its own process, together with a
baseline process that holds only the input. Five rounds run in a new random
order each, and the tables show medians.

## Running

```sh
go run ./cmd/bench                                  # the full suite: about 2-3.5 hours on an M1 Pro
go run ./cmd/bench -sizes 4096 -skipmem             # a quicker subset
go run ./cmd/bench -out /tmp/smoke -repeats 21 -validation 2 -minprocs 2 -maxprocs 2 -memn 16384 -memrounds 1
go run ./cmd/summarize results/speed.jsonl          # pool the raw results again
```

The driver starts its own binary for every process and writes to `results/`:
- `speed.jsonl` and `mem.jsonl`: one line per comparison per process;
- `speed-summary.md` and `mem-summary.md`: the pooled tables;
- `run.json`: settings, machine and time;
- `logs/`: rtcompare's full report for every process.

`-repeats`, `-loopscale` and `-validation` pass through to rtcompare. The
defaults are rtcompare's. Lower values are for smoke tests only.

## Earlier benchmarks

The benchmarks from the design phase are in [`_archive/`](_archive/README.md):
prototypes of the tree, comparisons with
[plar/go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree)
and [plar/go-hot-trie](https://github.com/plar/go-hot-trie), node sizes, the
value container, and the range scan that touches children ahead. They no
longer build from there. The last commit in which they ran is `57a4fa9`.
