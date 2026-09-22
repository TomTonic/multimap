# Benchmarks for the multimap key index

This module compares ART prototypes against existing ordered and unordered Go
structures. It is a separate module so that the library's `go.mod` does not
inherit the benchmark dependencies (rtcompare, tidwall/btree, plar ART).

- `proto/ptrart`: ART with ordinary Go pointers
- `proto/arenaart`: ART in chunked slabs with 32-bit refs (noscan)
- `proto/arenaflat`: like arenaart, one contiguous slice per node kind
- `proto/swar`: the node-search primitives shared by all three

All three prototypes use the same algorithm: sorted keys, SWAR search for
small nodes, bitmap + popcount rank for the 52/57-way node, direct index for
the 256-way node, lazy expansion, and 8-byte pessimistic plus optimistic path
compression with the full key in the leaf. `go test ./proto/...` checks each
one against a sorted reference.

## Running

```sh
./run.sh    # rtcompare: A = arena-art against every other structure  (~40 min)
./run2.sh   # memory/GC per structure, node search, arena-flat         (~15 min)
./run3.sh   # realistic multimap: value sets, iterators, vs B-tree     (~25 min)
```

The comparisons follow the rtcompare HOWTO. Each uses `rtcompare.Compare` with
A/A validation. `MaxQuantizationError` is set to 1e-4, because a smoke test at
the default showed an 18% tie rate. Build comparisons use `GCBetween`. Each
scenario runs in its own process. Raw results are in `results/*.jsonl`, and the
full reports with warnings are in `results/*.log`.

Memory and GC figures (`cmd/memgc`) are not rtcompare comparisons. GC cost
depends on the whole live heap, so two candidates in one process would each pay
for the other. Each structure is therefore measured alone in its own process.

## Results (Apple M1 Pro, Go 1.27.1, 2026-09-22)

Keys are u64 (random, 8 bytes, big endian) or str (path-like strings,
"tenant/category/word/NNNNN", about 26 bytes). Values are uint32. The tables
give the median ns/op of each candidate. All differences cited are resolved,
meaning they are larger than the measured noise floor, which was 0.2-1.7%.

### Point lookups (hit)

| keys, n    | ptr-art | arena-art | arena-flat | go-map | tidwall-btree | plar-art |
|------------|--------:|----------:|-----------:|-------:|--------------:|---------:|
| u64, 4K    |    16.8 |      18.6 |       17.2 |   12.0 |           116 |     27.5 |
| u64, 1M    |     102 |        83 |         79 |     97 |           843 |      344 |
| str, 4K    |    53.0 |      58.0 |       52.3 |   14.9 |           130 |     81.6 |
| str, 1M    |     303 |       278 |        274 |    115 |           978 |      931 |

### Range scan, 100 keys from a random existing key

| keys, n    | ptr-art | arena-art | arena-flat | tidwall-btree |
|------------|--------:|----------:|-----------:|--------------:|
| u64, 4K    |     580 |       621 |        569 |           345 |
| u64, 1M    |   2 521 |     3 543 |      3 312 |         1 962 |
| str, 4K    |   1 285 |     1 314 |      1 263 |           361 |
| str, 1M    |   5 575 |     7 009 |      6 070 |         1 855 |

### Memory and GC with 1M keys (per key, beyond the key corpus)

| structure     | u64 heap B | u64 GC CPU/cycle | str heap B | str GC CPU/cycle |
|---------------|-----------:|-----------------:|-----------:|-----------------:|
| arena-art     |         46 |          +0 ms   |         81 |           +0 ms  |
| arena-flat    |         78 |          +0 ms   |        132 |           +0 ms  |
| ptr-art       |         79 |        +134 ms   |        107 |         +196 ms  |
| tidwall-btree |         45 |         +50 ms   |         65 |          +57 ms  |
| go-map        |         61 |         +47 ms   |         81 |          +47 ms  |
| plar-art      |        108 |        +215 ms   |        135 |         +284 ms  |

In this table, arena-flat includes up to 2x slack capacity from slice doubling.
The GC column gives the CPU time of a full cycle minus a baseline process that
holds only the corpus.

### In-node child search (all nodes in cache, random hits)

| capacity | winner             | ns   | against            | ns   |
|---------:|--------------------|-----:|--------------------|-----:|
|  4 / 8   | SWAR (1 word)      |  1.3 | linear loop        | 8-10 |
|       22 | SWAR (3 words)     |  7.8 | bitmap + binary    | 22.6 |
|       22 | bitmap + popcount  |  8.0 | SWAR (3 words)     |  7.9 |
|       52 | bitmap + popcount  |  7.8 | bitmap + binary    | 27.3 |

## What the results say

1. ART is the right family. For point lookups it is 3-10x faster than
   tidwall/btree and plar-art. For u64 keys at 1M it is faster than Go's map.
2. Go's map is 2.5-4x faster than any tree for string point lookups. It
   provides no order, however, and a map plus an ordered index costs roughly
   double the memory.
3. B-trees win range scans by 2-3.6x. Part of that comes from this benchmark,
   where the B-tree stores uint32 values inline. The rest is structural: the
   ART keeps one leaf object per key (see the scan analysis below).
4. Arena against pointers: the arena is 8-22% faster only at 1M keys, and
   within ±5% (sometimes slower) in cache. The arena removes the tree's GC cost
   and uses about 40% less memory. It has no memory reclamation of its own, and
   dangling indices would fail silently.
5. Node search: SWAR and popcount beat binary search by 3-3.5x, and linear
   search by 1.5-7x.

### Why the ART scan is slower (profile + tree shape, ptr-art, 1M keys)

The CPU profile of `BenchmarkScan100` (proto/ptrart) shows the scan stalled on
memory. `minLeaf` (reading long prefixes from a leaf) does not appear at all.
Tree shape: u64 keys give 0.09 inner nodes per key at an average leaf depth of
3.1. Str keys give 0.33 inner nodes per key at an average leaf depth of 6.1,
with long (> 8 byte) prefixes on only 173 nodes. A scan over 100 keys therefore
touches about 110-135 separately allocated objects, one of which is a leaf per
key. A B-tree touches a few contiguous node arrays.

## Realistic multimap benchmark (`./run3.sh`)

This benchmark gives each key a set of uint64 values. Set sizes are skewed:
50% of keys hold 1 value, 35% hold 2-4, 12% hold 5-16 and 3% hold 17-200. All
results are consumed through `iter.Seq`. Three multimaps share the same value
container, `vset.Set`, which stores 3 values inline and spills to set3 above
that:

- `mmart`: the pointer ART with the value set inline in the 80-byte leaf.
- `btree-inline`: `tidwall/btree.Map[string, vset.Set]`. An update is Get,
  modify, Set, as the btree README describes.
- `btree-ptr`: `tidwall/btree.Map[string, *vset.Set]`.

`GCBetween` is used for build only. The first run of this suite also used it
for addRemove, which forced a full GC over a 1 GB heap before every batch.
That run was stopped and addRemove was re-measured without it.

Each cell gives ART / B-tree in ns/op and the resolved difference. A positive
difference means the ART is faster.

| keys, n  | op                    | vs btree-inline             | vs btree-ptr               |
|----------|-----------------------|-----------------------------|----------------------------|
| u64, 4K  | ValuesFor             | 115 / 229 (+50%)            | 114 / 225 (+49%)           |
| u64, 1M  | ValuesFor             | 476 / 1 006 (+53%)          | 482 / 1 082 (+56%)         |
| str, 4K  | ValuesFor             | 165 / 251 (+34%)            | 153 / 230 (+34%)           |
| str, 1M  | ValuesFor             | 771 / 1 268 (+39%)          | 728 / 1 217 (+40%)         |
| u64, 4K  | ValuesBetween (100 k) | 7 490 / 6 958 (-8%)         | 6 813 / 6 405 (-6%)        |
| u64, 1M  | ValuesBetween (100 k) | 12 248 / 12 638 (+3%)       | 12 797 / 15 213 (+16%)     |
| str, 4K  | ValuesBetween (100 k) | 7 614 / 6 462 (-18%)        | 8 361 / 7 265 (-15%)       |
| str, 1M  | ValuesBetween (100 k) | 18 537 / 13 118 (-41%)      | 18 793 / 15 539 (-21%)     |
| u64, 4K  | AddValue+RemoveValue  | 56 / 487 (+89%)             | 55 / 245 (+78%)            |
| u64, 1M  | AddValue+RemoveValue  | 372 / 1 483 (+75%)          | 371 / 1 159 (+68%)         |
| str, 4K  | AddValue+RemoveValue  | 140 / 509 (+73%)            | 139 / 266 (+48%)           |
| str, 1M  | AddValue+RemoveValue  | 676 / 1 496 (+55%)          | 678 / 1 208 (+44%)         |
| u64, 4K  | build                 | 1.42 ms / 4.79 ms (+70%)    | 1.43 ms / 3.23 ms (+56%)   |
| str, 4K  | build                 | 2.28 ms / 4.94 ms (+54%)    | 2.29 ms / 3.37 ms (+32%)   |

Memory and GC with 1M keys, including the values. Heap and scannable bytes
are per key; GC CPU is per cycle, minus baseline:

| multimap      | u64 heap | u64 scannable | u64 GC CPU | str heap | str GC CPU |
|---------------|---------:|--------------:|-----------:|---------:|-----------:|
| mmart         |    202 B |         128 B |    +224 ms |    230 B |    +288 ms |
| btree-inline  |    201 B |         121 B |    +130 ms |    209 B |    +129 ms |
| btree-ptr     |    184 B |          94 B |    +218 ms |    204 B |    +223 ms |

What these results say:

- **Point operations and build.** The ART is 1.5-9x faster than both B-tree
  variants.
- **Range scans with integer keys.** The earlier 2-3.6x gap has disappeared
  once real value sets are read. At 1M keys the ART is at least as fast as the
  B-tree.
- **Range scans with string keys.** The ART remains 15-41% behind. Keys longer
  than the 16 inline bytes live in a separate string, and the scan compares
  every key against the upper bound, so each key costs an extra cache miss.
  Checking the upper bound structurally (only along the boundary path) would
  remove that miss.
- **Memory is dominated by the value container, not the index.** A key costs
  about 200 bytes in every variant, while the ART leaf is 80 bytes. Most of the
  rest is set3 for the roughly 15% of keys with more than 3 values. The biggest
  lever for memory and GC is therefore a more compact spill representation, for
  example a sorted slice up to 16-32 values before switching to a hash set. The
  index layout matters less.

## Range scan with structural upper bound (`./run4.sh`)

`mmart.ValuesBetween` now checks both bounds structurally. Prefix and child
bytes are compared with `from` and `to` only along their two paths, and
subtrees strictly inside the range are visited without reading any key. The
former scan compared every leaf key with `to` and remains available as
`ValuesBetweenLinear` (`art-linear`) for this comparison. Because `art-linear`
runs on the same tree, its probe cursor starts half-way through the probes.
Without that offset, ABBA ordering lets each candidate re-scan ranges the
other just cached. A first run without the offset was discarded for that
reason.

Each cell gives ART / other in ns/op and the difference; a positive difference
means the ART is faster. "n.r." means not resolved.

| keys, n | vs art-linear            | vs btree-inline         | vs btree-ptr            |
|---------|--------------------------|-------------------------|-------------------------|
| u64, 4K | 6 417 / 6 652 (n.r.)     | 6 720 / 6 330 (-6%)     | 6 634 / 6 439 (n.r.)    |
| u64, 1M | 13 495 / 13 342 (n.r.)   | 12 622 / 13 063 (+3%)   | 12 312 / 15 057 (+18%)  |
| str, 4K | 7 160 / 7 398 (+3%)      | 7 597 / 6 386 (-19%)    | 7 268 / 6 610 (n.r.)    |
| str, 1M | 15 637 / 18 331 (+15%)   | 15 701 / 12 977 (-21%)  | 15 852 / 15 613 (-1.5%) |

The structural bound saves 15% for long string keys at 1M, because it avoids
one key read per leaf. It makes no difference for u64 keys, which are inline
in the leaf anyway. The remaining gap to btree-inline on long string keys is
the one leaf access per key, against the B-tree's contiguous item arrays.

## Value container: array spill instead of hash spill (`./run5.sh`)

`vset.Set` now has three representations: up to 3 values inline, then up to
64 values in a plain unsorted array (noscan for pointer-free T, grown by
`append` along Go's size classes), then a set3 hash set. It shrinks with
hysteresis: from hash to array at 32 values, from array to inline below 3. The
struct stays at 40 bytes, so the ART leaf stays at 80 bytes. The former design
(3 inline, then straight to set3) remains as `vset.HashSpill` for comparison.

Standalone, a set3 costs 2-4x an array for 4-32 values (4 values: 153 B vs
32 B; 32 values: 656 B vs 256 B). The two converge at about 100 values.

Memory and GC with 1M keys and the skewed value distribution:

| what                    | hash spill | array spill | change |
|-------------------------|-----------:|------------:|-------:|
| containers alone, heap  |  131 B/key |   107 B/key |   -18% |
| containers alone, GC    |    +26 ms  |     +14 ms  |   -49% |
| mmart u64, heap         |  202 B/key |   178 B/key |   -12% |
| mmart str, heap         |  230 B/key |   206 B/key |   -10% |
| btree-inline u64, heap  |  201 B/key |   177 B/key |   -12% |
| btree-ptr u64, heap     |  184 B/key |   160 B/key |   -13% |

Speed, measured with rtcompare on n contiguous containers with random probes.
Each cell gives array spill / hash spill in ns/op:

| n  | Contains              | iterate all values     | Add+Remove (absent)    |
|----|-----------------------|------------------------|------------------------|
| 4K | 12.6 / 11.7 (-8%)     | 40.3 / 59.4 (+32%)     | 19.3 / 22.0 (+12%)     |
| 1M | 50.8 / 76.9 (+34%)    | 82.1 / 143.3 (+43%)    | 76.5 / 114.1 (+33%)    |

The array spill wins everywhere except `Contains` on a warm cache. There, a
linear scan over up to 64 values loses 8% to a hashed probe. Once the data is
not cached, one contiguous array beats the hash set's pointer chases by a
third or more.

With the new container, the remaining GC difference is in the index itself.
mmart costs about 2x btree-inline in GC CPU per cycle (+178 ms against
+91 ms for u64) at a similar number of scannable bytes. The ART allocates one
object per key (the leaf) plus inner nodes. btree-inline packs many keys into
each node array, so the GC visits far fewer objects.
