# Where pages for keys with several values start to pay off (2026-09-28)

`afd1137` (range nodes, K and S pages, keys with several values in pages)
against main's Ordered (`7c7e6a2`) as A, several values per key only, between
the sizes of `../night-v0.7` (4K and 256K). Dev suite with rtcompare v0.7.0,
5-10 processes per scenario, builds up to 64K keys, memory at 64K keys:

```sh
bench-vs-afd1137 -suite dev -buildmax 65536 -vs baseline -values multi \
  -keys u64,str,uuid,path,street -sizes 16384,32768,65536,131072 \
  -minprocs 5 -maxprocs 10 -memn 65536 -memrounds 1
```

It took 1 h 50 min and did not sleep. Figures are `afd1137`'s speed relative
to main (above 1 = faster); `*` marks comparisons whose interval across
processes stayed wider than ±2 points or ±10% after 10 processes.

| operation | keys | 16K | 32K | 64K | 128K |
|---|---|---:|---:|---:|---:|
| `valuesFor` | u64 | 0.78 | 0.71 | 0.87* | 1.20* |
| | str | 0.85 | 0.85 | 1.18* | 1.52* |
| | uuid | 0.76 | 0.81* | 1.08* | 1.06* |
| | path | 0.90 | 0.96* | 1.23* | 1.23* |
| | street | 0.93 | 0.93* | 1.16* | 1.52* |
| `valuesBetween` | u64 | 1.54 | 1.69 | 1.61 | 1.67 |
| | str | 1.35 | 1.27 | 1.35 | 1.64 |
| | uuid | 1.30 | 1.19 | 1.37 | 1.67 |
| | path | 1.32 | 1.28 | 1.37 | 1.45 |
| | street | 1.79 | 1.69 | 1.75 | 2.00 |
| `prefix` | str | 1.49 | 1.37 | 1.45 | 1.79 |
| | uuid | 0.53 | 0.58 | 0.88 | 0.95 |
| | path | 0.97 | 1.12* | 1.20* | 1.41 |
| | street | 1.67 | 1.92 | 2.13 | 2.38 |
| `churn` | u64 | 0.56 | 0.68 | 0.94* | 1.06* |
| | str | 0.71 | 0.83* | 1.02* | 0.99* |
| | uuid | 0.67 | 0.83* | 0.96* | 0.90* |
| | path | 0.88 | 1.04 | 1.03* | 0.98* |
| | street | 0.73 | 0.87 | 1.03* | 1.05* |
| `build` | u64 | 0.55 | 0.60 | 0.71 | |
| | str | 0.48 | 0.57 | 0.68 | |
| | uuid | 0.48 | 0.52 | 0.61 | |
| | path | 0.60 | 0.68 | 0.76 | |
| | street | 0.44 | 0.51 | 0.58 | |

Memory at 64K keys, bytes per key, main → `afd1137`: u64 159 → 110, str
177 → 121, uuid 194 → 154, path 235 → 163, street 119 → 60 (−18 to −50%);
GC CPU per cycle falls by 40-75%.

## Findings

1. **Point lookups turn between 32K and 64K keys** for every kind: 0.71-0.96
   at 32K, 0.87-1.23 at 64K, 1.06-1.52 at 128K. The jump is abrupt. At 64K,
   main's tree holds about 10-15 MB, close to the reach of Zen 4's L2 TLB
   (3,072 pages of 4 KB = 12 MB); pages then touch fewer pages per lookup.
   That is a guess, not a measurement.
2. **Churn turns at about 64K** (0.94-1.03) and is level at 128K
   (0.90-1.06): it does not gain the way point lookups do.
3. **Range queries win at every size** (1.19-2.00), as they did at 4K;
   prefix searches too, except uuid (random keys, 0.53-0.95).
4. **Builds from empty lose at every size measured** (0.44-0.76): a build
   passes through all the small sizes on its way. A size switch would build
   with leaves up to the threshold and would pay once for the conversion.
5. **A size switch at about 64K keys** would therefore keep main's speed below
   and give pages' gains above: point 1.1-1.5, range 1.4-2.0, memory −20 to
   −50% from 128K keys up, churn level.
