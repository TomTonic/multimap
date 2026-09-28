# Inner nodes with a 16-byte header (2026-09-28)

The node header shrinks from 24 to 16 bytes: `kind` 1, `count` 1, `plen` 2
and 12 path bytes instead of 8. The term leaf moves from a header field into
the node's last child slot. The same size classes then hold 5, 12, 26 and 58
children instead of 4, 11, 25 and 57.

This is Ordered with the new header against main's Ordered (`37ca7c8`),
compared head to head as `-vs baseline`. It used the dev suite with rtcompare
v0.7.0, 5 to 10 processes per scenario, builds up to 16K keys and memory at
1M keys:

```sh
bench -suite dev -buildmax 65536 -vs baseline -sizes 4096,16384,262144,1048576 \
  -minprocs 5 -maxprocs 10 -memn 1048576 -memrounds 3
```

The run took 1 h 38 min and did not sleep. 201 of 220 comparisons are precise;
the other 19 are mostly churn from 256K keys up.

## Results

The figures are the new header's speed relative to main (above 1 means
faster). Each cell gives the range over the key kinds and the median in
brackets.

| values | operation | 4K | 16K | 256K | 1M |
|---|---|---|---|---|---|
| multi | `valuesFor` | 0.98–1.02 (1.00) | 0.98–1.04 (1.01) | 1.00–1.04 (1.03) | 1.01–1.03 (1.02) |
| multi | `valuesBetween` | 0.99–1.05 (1.01) | 1.00–1.05 (1.01) | 1.00–1.04 (1.02) | 1.00–1.02 (1.01) |
| multi | `prefix` | 0.98–1.00 (0.99) | 0.98–1.01 (1.00) | 0.99–1.03 (1.01) | 0.99–1.03 (1.01) |
| multi | `churn` | 0.99–1.01 (0.99) | 0.99–1.01 (1.00) | 0.94–1.02 (1.01)* | 0.96–1.02 (0.98)* |
| multi | `build` | 0.99–1.01 (1.00) | 0.98–1.01 (1.00) | | |
| unique | `valuesFor` | 0.97–1.01 (1.01) | 0.99–1.06 (1.01) | 1.01–1.06 (1.02) | 1.00–1.04 (1.02) |
| unique | `valuesBetween` | 0.99–1.09 (1.02) | 1.00–1.09 (1.02) | 1.00–1.05 (1.02) | 1.00–1.06 (1.01) |
| unique | `prefix` | 0.96–1.03 (1.00) | 0.98–1.03 (1.00) | 1.01–1.05 (1.02) | 0.99–1.05 (1.02) |
| unique | `churn` | 0.98–0.99 (0.98) | 0.98–1.00 (0.99) | 0.98–1.03 (1.01) | 1.00–1.02 (1.00) |
| unique | `build` | 0.97–1.00 (0.99) | 0.96–1.00 (0.99) | | |

`*` marks cells with imprecise comparisons. The processes of multi churn from
256K keys up spread several points.

**Memory:** the new header uses 0–2 bytes less per key for every kind.
Scannable memory is 1–2 bytes lower. The GC time per cycle stays within the
spread of the rounds.

## Findings

1. **The new header is neutral to slightly faster.** Point lookups from 16K
   keys up are 1–6% faster. Range queries gain most with u64 keys (4–9%),
   presumably because each node kind holds one more child and fewer nodes
   grow into the next kind. That is a guess; the node mix was not counted.
2. **Small losses remain within 2–4%:**
   - path point lookups at 4K keys (0.97–0.98);
   - unique churn at 4K keys (0.98);
   - u64 builds (0.96–0.97);
   - multi churn at 1M keys (median 0.98, imprecise).
3. **Memory gains little**, because inner nodes are only a small part of a
   key's bytes: there are 0.1–0.5 inner nodes per key. The large
   items are the leaves with their value sets. The flat leaf targets those.
4. **Regression criterion:** no operation loses more than 6%, and the median
   loses at most 2%. Criterion 3 (at most 15% regression against main) holds.
