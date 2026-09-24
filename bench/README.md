# Benchmarks for the multimap key index

This module benchmarks the library's ART-backed index against the
alternatives a Go developer would otherwise reach for. It is a separate
module so the library's `go.mod` does not inherit the benchmark
dependencies (rtcompare, tidwall/btree, plar/go-adaptive-radix-tree,
plar/go-hot-trie).

`proto/` holds the ART prototypes that led to the shipped design (pointer
nodes, SWAR/popcount child search, path compression); `cmd/` holds the
benchmark drivers.

This page is intentionally short — it exists to help pick a data structure,
not to document every step of the design process. The full iteration history
(arena allocation, node/leaf sizing, value-set spill strategy) is in git.

## What's compared, and why

Four off-the-shelf alternatives are benchmarked next to this project's own
ART (the `ptr-art` prototype, which ships as `Ordered`). They were picked
because each is pure Go (no cgo), actively maintained, and covers a shape of
the design space this project's index has to justify itself against:

| Candidate | What it is | Why it's here |
|---|---|---|
| Go's built-in `map` | the standard unordered hash map | this library's own `Hashed` variant is built directly on it — the baseline for "what does ordering cost?" |
| [tidwall/btree](https://github.com/tidwall/btree) | a widely used, actively maintained in-memory B-tree with generics | the standard off-the-shelf *ordered* Go container — the natural alternative to hand-rolling an ART when range queries matter |
| [plar/go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) | the most established pure-Go ART implementation | the direct algorithmic peer — answers whether a purpose-built layout is worth it over an existing ART |
| [plar/go-hot-trie](https://github.com/plar/go-hot-trie) | a Go implementation of HOT (Height Optimized Trie), v0.1.2 | HOT keeps the fanout near 32 regardless of key distribution, which is exactly where ART is weak: sparse string keys make our tree twice as deep as random u64 keys |

## Running

```sh
./run.sh    # point lookups, range scan, build: this project's ART vs. the other candidates
./run2.sh   # memory/GC per candidate, in-node child-search strategies
./run3.sh   # realistic multimap: value sets, iterators, vs. tidwall/btree
./run4.sh–run7.sh  # internal design iterations (see git history)
./run8.sh   # HOT vs. this project's ART; memory/GC of all candidates, 3 rounds
./run9.sh–run11.sh # range scan that touches children ahead, prototype and library
```

Each scenario runs in its own process via `rtcompare.Compare` with A/B/B/A
interleaving, which cancels drift from thermal throttling or background
load — never compare absolute numbers from separate `go test -bench` runs.
Memory/GC (`cmd/memgc`) measures one candidate per process instead, since GC
cost depends on the whole live heap. Raw results are in `results/*.jsonl`,
full reports in `results/*.log`. `run.sh` to `run7.sh` together take about
100 minutes.

## Results (Apple M1 Pro, Go 1.27.1, 2026-09-23)

Keys are u64 (random, 8 bytes) or str (path-like, ~26 bytes). Values are
uint32. Figures are median ns/op; percentages are only given where the
difference is resolved (larger than the measured noise floor, 0.1–3%).

### Point lookups (hit)

| keys, n      | multimap (ART) | go-map | tidwall-btree | plar-art |
|--------------|----------------:|-------:|---------------:|---------:|
| u64, 4K      |            16.7 |   11.9 |            116 |     26.8 |
| u64, 1M      |             118 |    111 |            989 |      347 |
| str, 4K      |            56.1 |   14.8 |            134 |     85.8 |
| str, 1M      |             316 |    103 |          1 104 |      940 |

### Range scan, 100 keys from a random existing key

go-map has no order and plar-art wasn't wired into this scenario, so only
the two ordered candidates are compared:

| keys, n      | multimap (ART) | tidwall-btree |
|--------------|----------------:|--------------:|
| u64, 4K      |             581 |           344 |
| u64, 1M      |           3 094 |         1 748 |
| str, 4K      |           1 344 |           365 |
| str, 1M      |           5 497 |         2 264 |

### HOT against this project's ART (`./run8.sh`, 2026-09-24)

Measured as one interleaved rtcompare pair per cell (ART / HOT in ns/op; all
differences resolved). HOT has no seek in v0.1.2, so there is no range-scan
row, and its SIMD search does not run on arm64, so this is its scalar path.

| keys, n | get          | miss         | build              |
|---------|--------------|--------------|--------------------|
| u64, 4K | 16.6 / 57.4  | 8.6 / 61.6   | 0.20 / 0.80 ms     |
| u64, 1M | 107 / 466    | 44.3 / 469   |                    |
| str, 4K | 53.1 / 75.5  | 41.2 / 70.6  | 0.43 / 1.34 ms     |
| str, 1M | 296 / 464    | 180 / 401    |                    |

### Memory & GC, 1M keys, per key beyond the key corpus

GC CPU is per full cycle minus a baseline process that holds only the corpus;
median of three round-robin rounds with 50 cycles each (`./run8.sh`).

| candidate      | u64 heap B | u64 GC CPU/cycle | str heap B | str GC CPU/cycle |
|----------------|-----------:|-----------------:|-----------:|-----------------:|
| multimap (ART) |         71 |          +135 ms |         99 |          +195 ms |
| go-map         |         53 |           +43 ms |         73 |           +54 ms |
| tidwall-btree  |         36 |           +52 ms |         56 |           +68 ms |
| plar-art       |        100 |          +213 ms |        127 |          +300 ms |
| plar-hot       |         83 |           +90 ms |        100 |           +95 ms |

### Realistic multimap (`./run3.sh`): value sets, not bare keys

Each key holds a skewed number of values (50% hold 1, 35% hold 2-4, 12% hold
5-16, 3% hold 17-200), read through `iter.Seq`. This compares the ART
multimap prototype that `Ordered` is built from (`mmart`) against layering the
same value container on
`tidwall/btree.Map[string, vset.Set]`:

| keys, n | op                    | ART / btree-inline (ns) | diff |
|---------|-----------------------|-------------------------:|-----:|
| u64, 4K | ValuesFor             |               86.6 / 194 | +55% |
| u64, 1M | ValuesFor             |               427 / 909  | +53% |
| str, 4K | ValuesFor             |               127 / 204  | +38% |
| str, 1M | ValuesFor             |               648 / 1032 | +37% |
| u64, 4K | ValuesBetween (100 k) |             4 297 / 3 958 |  -9% |
| u64, 1M | ValuesBetween (100 k) |             9 103 / 7 558 | -20% |
| str, 4K | ValuesBetween (100 k) |             5 063 / 3 957 | -28% |
| str, 1M | ValuesBetween (100 k) |            12 650 / 8 222 | -54% |
| u64, 4K | AddValue+RemoveValue  |                53.4 / 480 | +89% |
| u64, 1M | AddValue+RemoveValue  |               324 / 1 287 | +75% |
| str, 4K | AddValue+RemoveValue  |               139 / 503   | +72% |
| str, 1M | AddValue+RemoveValue  |               645 / 1 416 | +54% |
| u64, 4K | build                 |          1.15 ms / 4.52 ms| +75% |
| str, 4K | build                 |          2.03 ms / 4.63 ms| +56% |

Memory converges once real value sets are attached: at 1M keys, the ART
multimap, `btree-inline` and `btree-ptr` all land between 160-206 B/key,
because the value container dominates, not the key index. GC cost does not
converge: +176 ms (u64) and +245 ms (str) per cycle for the ART multimap
against +88 ms for `btree-inline`, because the ART allocates one object per
key where the B-tree packs up to 63 keys into one node array.

### Range scan that touches children ahead (`./run9.sh`–`./run14.sh`, 2026-09-24)

A scan otherwise takes the cache misses for a node's children one after
another. Reading one byte of every child in range before descending lets the
CPU keep those misses in flight at once. The library now always does this,
skipping nodes with fewer than two children in range. The tables above were
measured before, without it. Difference against the same scan without
touching (positive = touching is faster; n.r. = within noise). The library
column compares `Ordered` with the prototype's plain scan and gives the range
over three runs (`./run11.sh`–`./run13.sh`):

| keys | 4K    | 16K   | 64K  | 256K  | 1M    | library `Ordered`, 4K / 1M |
|------|-------|-------|------|-------|-------|----------------------------|
| u64  | n.r.  | n.r.  | n.r. | +14%  | +19%  | -3 to -4% / +14 to +16%    |
| str  | -4%   | n.r.  | +8%  | +16%  | +20%  | -6 to -11% / +14 to +19%   |

Where the gain starts depends on the cache size and on what else the program
keeps in the cache, so the library touches unconditionally: a tree that stays
cached loses at most about 4% to touching, one that does not gains up to 20%.
With touching, the ART ties `btree-inline` on u64 keys at 1M and trails it by
24% on str keys at 1M (previously 54%).

The library first looked 5-6% slower than the prototype's touching scan on
str keys. Two things explain it (`./run14.sh`). Iterating a value set that
had spilled into a hash set copied the whole set on every scan (4 allocations
per 100-key scan, in library and prototype alike); after that fix the gap at
4K is 1-2%. At 1M the sign depends on which structure is built first: -2.2%
with the library built last, +2.2% with it built first.

At 1M keys, the three runs of the same pair disagree by more than each run's
interval allows (str, library vs. touching prototype: -9%, -2%, -6%, each
within about ±2%). Every process builds its trees anew, and where they land
in memory shifts the result. A single-digit difference at 1M therefore needs
several processes before it means anything; the intervals of 4K runs overlap.

## What this means for picking a structure

- **Only need point lookups, no ordering?** Go's `map` wins, especially for
  string keys (3x faster than this project's ART, 8-11x faster than
  tidwall/btree). That's this library's `Hashed` variant.
- **Need ordering but each key has one value?** `tidwall/btree` (or any plain
  ordered map) is simpler than this library and comparable in speed, ahead on
  large range scans over string keys.
- **Need a multimap — ordering *and* multiple values per key?** This
  library's `Ordered`/ART beats layering a value set on `tidwall/btree` by
  37-89% on every op except range scans (up to 54% behind before the scan
  touched children ahead, up to 24% after). The ART compares
  keys only along the two boundary paths; everything in between is emitted
  without a single comparison. What costs time is the traversal itself:
  every key is its own leaf object, reached through a pointer, plus the inner
  nodes on the way. For 100 keys that is 110-135 scattered objects, while
  `tidwall/btree` holds keys and value sets in a few contiguous node arrays.
  String keys make the tree deeper, which makes the gap larger.
- **plar/go-adaptive-radix-tree**, the existing off-the-shelf ART, is
  consistently slower (1.5-3x) and heavier (28-41% more memory) than this
  project's own ART. The specialization — inline value sets, tuned node
  fan-out, no separate leaf allocation for the value container — is what
  buys that gap; a generic ART library doesn't get it for free.
- **plar/go-hot-trie** is built for a low tree height, but in this Go
  implementation that does not turn into speed: this project's ART is 1.4-1.6x
  faster on string lookups, 3.5-4.3x on u64 lookups and 3-4x faster to
  build. Its heap is similar (it boxes each value in an `any`, as plar-art
  does), and its GC cost is one half to two thirds of the ART's. Without a
  seek it cannot serve range queries yet.
- An **arena-allocated** version of the ART was prototyped and dropped: it
  was only 8-22% faster at 1M keys, and the memory safety it gives up
  (dangling indices fail silently on reuse) wasn't worth it for a
  general-purpose library.
