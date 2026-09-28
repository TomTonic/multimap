# Flat leaves (2026-09-28)

A key whose values are small and free of pointers (up to 16 bytes, such as
integers or small structs of them) keeps its values in its own leaf instead
of a `vset.Set` behind it:

```text
leafHead (4 B) | key (klen bytes) | padding to T's alignment | values [n]T
```

The leaf is one Go size class of 32 to 512 bytes, allocated as an array of
`uint64`, so the garbage collector never scans it. A new key gets the
smallest class that holds one value. A full leaf moves to the smallest class
that holds twice its values, and to at least 64 bytes. It moves back only
once a smaller class would still be half empty. A key with more values than
512 bytes hold becomes a set leaf as before, and becomes flat again at half
of that. Other value types and keys over 254 bytes keep set leaves.

This is Ordered with flat leaves against the 16-byte header of
[header16](../header16/README.md) (`cff024d`), compared head to head as
`-vs baseline`, with the same command line:

```sh
bench -suite dev -buildmax 65536 -vs baseline -sizes 4096,16384,262144,1048576 \
  -minprocs 5 -maxprocs 10 -memn 1048576 -memrounds 3
```

The run took 1 h 46 min and did not sleep. 190 of 220 comparisons are precise;
the other 30 are mostly churn and 256K lookups.

## Results

**Memory** at 1M keys (path 300K, street 212K), flat leaves against the
baseline:

| values | keys | heap B/key | scannable B/key | GC CPU per cycle |
|---|---|---:|---:|---:|
| multi | u64 | 148 → 116 | 82 → 20 | 209 → 83 ms |
| multi | str | 175 → 140 | 109 → 31 | 317 → 144 ms |
| multi | uuid | 191 → 147 | 125 → 32 | 315 → 139 ms |
| multi | email | 174 → 139 | 108 → 30 | 281 → 130 ms |
| multi | url | 228 → 190 | 117 → 34 | 365 → 162 ms |
| multi | path | 231 → 196 | 126 → 41 | 89 → 43 ms |
| multi | street | 118 → 85 | 103 → 35 | 41 → 24 ms |
| unique | u64 | 73 → 41 | 81 → 17 | 180 → 73 ms |
| unique | str | 99 → 66 | 107 → 27 | 270 → 131 ms |
| unique | uuid | 116 → 68 | 124 → 28 | 268 → 124 ms |
| unique | email | 98 → 65 | 106 → 26 | 239 → 117 ms |
| unique | url | 153 → 119 | 115 → 31 | 322 → 143 ms |
| unique | path | 156 → 121 | 124 → 37 | 80 → 39 ms |
| unique | street | 95 → 60 | 103 → 34 | 42 → 23 ms |

**Speed** of flat leaves relative to the baseline (above 1 means faster).
Each cell gives the range over the key kinds and the median in brackets:

| values | operation | 4K | 16K | 256K | 1M |
|---|---|---|---|---|---|
| multi | `valuesFor` | 0.99–1.01 (1.01) | 0.98–1.01 (1.00) | 1.02–1.17 (1.06)* | 0.95–1.10 (1.02) |
| multi | `valuesBetween` | 0.99–1.02 (1.01) | 0.99–1.02 (1.01) | 1.04–1.18 (1.16) | 1.08–1.20 (1.18) |
| multi | `prefix` | 1.00–1.02 (1.01)* | 1.00–1.02 (1.02)* | 1.01–1.19 (1.12) | 1.07–1.19 (1.15) |
| multi | `churn` | 0.96–1.00 (0.97) | 0.95–1.04 (0.99)* | 0.93–0.99 (0.96)* | 0.96–1.00 (0.96)* |
| multi | `build` | 0.96–0.99 (0.97) | 0.94–1.01 (0.98)* | | |
| unique | `valuesFor` | 0.99–1.04 (1.02) | 0.97–1.03 (0.99) | 0.99–1.28 (1.10)* | 0.96–1.10 (1.01) |
| unique | `valuesBetween` | 0.98–1.01 (0.99) | 0.98–1.00 (0.99) | 0.95–1.39 (1.16)* | 0.96–1.07 (1.04) |
| unique | `prefix` | 0.98–1.01 (1.00)* | 0.98–1.01 (1.00) | 0.91–1.19 (1.09)* | 0.95–1.08 (1.04)* |
| unique | `churn` | 1.00–1.06 (1.02) | 1.01–1.09 (1.04) | 1.04–1.08 (1.06) | 0.98–1.03 (1.02)* |
| unique | `build` | 1.00–1.06 (1.02) | 1.01–1.07 (1.03) | | |

`*` marks cells with imprecise comparisons.

## Findings

1. **Memory and GC are the large win.** The heap shrinks by 15–44% per
   key and the GC time per cycle by 41–60%. Scannable memory falls to a
   fifth, since only inner nodes and set leaves hold pointers. Unique u64
   keys now take 41 bytes per key (btree-map on main: 37).
2. **Large maps read faster.** Multi range and prefix queries gain 4–20%
   from 256K keys up, since a key's values no longer sit behind a second
   pointer. Point lookups gain up to 28% at 256K keys.
3. **Unique writes are faster** (up to 9%): a new key takes one pointer-free
   allocation of 32–48 bytes instead of 64–96 bytes with pointers.
4. **Multi writes lose 0–7%**, most with u64 and street keys. A leaf that
   outgrows its class moves into a new object, where the set only appended to
   its array.
5. **Unique url and path range queries lose 4–9%** from 256K keys up. Their
   long key pushes the value out of the leaf's first 32 bytes, which the scan
   touches ahead (`touchChildren`), so the value costs a second miss.
6. **Regression criterion:** multiplying both runs per scenario estimates
   the 16-byte header and flat leaves against main. The lowest products are
   unique url prefix at 256K (0.92) and multi u64 churn at 1M (0.92).
   Criterion 3 (at most 15% regression against main) holds.
