# Flat leaf class in the leaf kind (2026-09-28)

A flat leaf's size class moves from its own byte in `leafHead` into the
leaf's kind: `kSet` for a set leaf, `kSet+c` for a flat leaf of class `c`,
and the inner nodes after them. A leaf is now any kind up to `kLastLeaf`,
tested with one comparison as before. The freed byte widens the value count
to 16 bits.

This is Ordered with the folded class against flat leaves
([flat-leaves](../flat-leaves/README.md), `7062206`), compared head to head
as `-vs baseline`. Both sides use the new `url` corpus of real Tranco hosts
(`57cc38d`), so its figures are not comparable with earlier runs:

```sh
bench -suite dev -buildmax 65536 -vs baseline -sizes 4096,16384,262144,1048576 \
  -minprocs 5 -maxprocs 10 -memn 1048576 -memrounds 3
```

The run took 1 h 42 min and did not sleep. 194 of 220 comparisons are precise;
the other 26 are mostly churn and lookups from 256K keys up.

## Results

Speed relative to the baseline (above 1 means faster). Each cell gives the
range over the key kinds and the median in brackets:

| values | operation | 4K | 16K | 256K | 1M |
|---|---|---|---|---|---|
| multi | `valuesFor` | 0.99–1.00 (1.00) | 1.00–1.00 (1.00) | 0.98–1.02 (1.00)* | 1.00–1.04 (1.00) |
| multi | `valuesBetween` | 0.99–1.02 (1.00) | 0.99–1.00 (1.00) | 0.99–1.00 (1.00) | 0.99–1.00 (1.00) |
| multi | `prefix` | 0.98–1.02 (1.00) | 0.99–1.01 (1.00)* | 0.99–1.01 (1.00)* | 1.00–1.02 (1.00) |
| multi | `churn` | 0.99–1.01 (1.00) | 1.00–1.01 (1.00) | 0.94–1.03 (1.00)* | 0.99–1.03 (1.00)* |
| multi | `build` | 0.99–1.00 (1.00) | 0.99–1.00 (1.00) | | |
| unique | `valuesFor` | 0.99–1.00 (1.00) | 0.99–1.01 (1.00) | 0.98–1.01 (1.00)* | 0.99–1.01 (1.00)* |
| unique | `valuesBetween` | 0.96–1.01 (0.99) | 0.99–1.00 (0.99) | 0.99–1.01 (1.00) | 0.99–1.01 (1.00) |
| unique | `prefix` | 0.99–1.01 (1.00)* | 0.99–1.00 (1.00) | 0.98–1.01 (0.99)* | 1.00–1.02 (1.00)* |
| unique | `churn` | 1.00–1.01 (1.00) | 0.99–1.01 (1.01) | 0.99–1.01 (0.99)* | 0.98–1.01 (0.99)* |
| unique | `build` | 0.99–1.00 (1.00) | 0.99–1.00 (1.00) | | |

`*` marks cells with imprecise comparisons.

**Memory:** identical for every kind; the GC time per cycle stays within the
spread of the rounds.

## Findings

1. **The change is neutral.** Every median lies at 0.99–1.01. Leaves keep
   their size, so memory cannot change.
2. **Two outliers:** unique u64 `valuesBetween` at 4K keys is 4% slower and
   precise (0.96 [0.95, 0.97]); multi email churn at 256K keys is 6% slower,
   but imprecise (0.94 [0.91, 0.97]). Neither shows at the other sizes or
   kinds.
3. The change is kept for what it simplifies: one byte less in the leaf
   head, and a value count that no longer overflows with one-byte values.
