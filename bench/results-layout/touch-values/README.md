# Touching the first value behind long keys (2026-09-28, rejected)

Flat leaves hold their values after the key. With a key of more than about
20 bytes, the first value lies beyond the 32 bytes that the range scan
touches ahead in every leaf (`touchChildren`), in a cache line that is
fetched only when the scan reaches the leaf. The idea was a second wave in
`touchChildren`: after the first wave has touched every child, read each
flat leaf's first value, at an offset its key length gives
([`touch-values.patch`](touch-values.patch), on top of `1d8082a`).

This is Ordered with the second wave against `63130b7`, compared head to
head as `-vs baseline`, range queries only:

```sh
bench -suite dev -vs baseline -ops valuesBetween,prefix \
  -sizes 4096,16384,262144,1048576 -minprocs 5 -maxprocs 10 -skipmem
```

The run took 19 min and did not sleep. 82 of 92 comparisons are precise.

## Results

Speed relative to the baseline (above 1 means faster). Each cell gives the
range over the key kinds and the median in brackets:

| values | operation | 4K | 16K | 256K | 1M |
|---|---|---|---|---|---|
| multi | `valuesBetween` | 0.90–1.00 (0.97) | 0.91–0.99 (0.96) | 0.96–1.00 (0.97) | 0.96–0.99 (0.97) |
| multi | `prefix` | 0.95–1.02 (1.00) | 0.90–1.01 (0.99)* | 0.95–1.00 (0.99)* | 0.96–1.00 (0.97)* |
| unique | `valuesBetween` | 0.87–0.95 (0.93) | 0.79–0.94 (0.91) | 0.92–0.97 (0.94) | 0.94–0.97 (0.95) |
| unique | `prefix` | 0.93–1.02 (0.99)* | 0.89–1.00 (0.99)* | 0.92–1.01 (0.97)* | 0.95–0.98 (0.96)* |

`*` marks cells with imprecise comparisons.

## Findings

1. **The second wave makes range queries slower everywhere**, by 3–9% in
   the median and up to 21% (unique u64 at 16K keys).
2. **Its cost is the second pass, not its loads.** u64 and street keys
   lose too, although their first value lies within the touched 32 bytes,
   so that the second wave adds no cache miss for them. Walking every child
   in range again, with a dependent branch per child, costs more than the
   scan saves; a 256-way node walks all 256 slots twice.
3. **It does not help where it should.** url and path, the long keys it was
   meant for, lose as well, also at 1M keys where the tree does not fit in
   the cache. Presumably the CPU's adjacent-line prefetcher already fetches
   the value's line together with the leaf's first line, so that the loss of
   unique url and path range queries in [flat-leaves](../flat-leaves/README.md)
   has another cause. That is a guess; it was not measured.
4. The change is rejected. The patch stays here for reference.
