# node-layout against the competitors (2026-10-01)

`node-layout` at `092952a` (typed leaves) against the structures it is meant to sit between: a B-tree
with set values and a hash map. Serial regime, native Windows, Ryzen 9 7900, rtcompare v0.8.0, 4K, 16K and 256K keys.

- `c5-unique-u64/`, `c5-unique-str/`: one value per key, against `btree-map` (the B-tree with one value per
  key, the credo's reference for the unique profile), values `uint64` and string (`-tags strvals`),
  first stage of 6 processes, at most 12, memory at 1M keys (3 rounds). About 1 h each.
- `c6-multi-u64/`: several values per key, `uint64`, against `hashed`, `btree-sets` and `map-sets`, key kinds
  `u64`, `str`, `uuid`, `url`, no memory. 6 h 13 min.

```sh
bench -suite dev -values unique -vs btree-map -sizes 4096,16384,262144 -minprocs 6 -maxprocs 12 -memn 1048576 -memrounds 3
bench -suite dev -values multi -vs hashed,btree-sets,map-sets -keys u64,str,uuid,url -sizes 4096,16384,262144 -minprocs 6 -maxprocs 12 -skipmem
```

Figures are **ordered / competitor**: above 1 the candidate is faster. Each cell gives the range over the
key kinds and the median in brackets; `*n` marks n comparisons that are not precise. `valuesBetween` and
`prefix` against `hashed` and `map-sets` are not a fair race (they must sort first) and show where the
ordered index earns its keep. `build` stops at 64K keys, so it has no 256K column.

## One value per key, against `btree-map`

### `uint64` values
| operation | 4K | 16K | 256K |
|---|---|---|---|
| `valuesFor` | 1.45–6.23 (2.85) | 1.45–4.79 (2.70) | 1.06–3.06 (1.40) |
| `valuesBetween` | 0.27–0.41 (0.29) | 0.26–0.34 (0.30) | 0.47–0.55 (0.51) *1 |
| `prefix` | 0.27–1.71 (0.77) *1 | 0.27–1.86 (0.73) | 0.44–1.19 (0.62) *2 |
| `churn` | 0.92–2.61 (1.48) | 0.98–3.00 (1.74) | 1.03–1.58 (1.46) |
| `build` | 0.97–2.25 (1.44) | 1.02–2.51 (1.64) |  |

### String values
| operation | 4K | 16K | 256K |
|---|---|---|---|
| `valuesFor` | 1.44–6.32 (2.79) | 1.45–4.98 (2.72) | 1.08–2.88 (1.38) *1 |
| `valuesBetween` | 0.26–0.39 (0.28) | 0.26–0.34 (0.29) | 0.43–0.53 (0.48) |
| `prefix` | 0.26–1.72 (0.76) | 0.27–1.89 (0.71) | 0.42–1.20 (0.62) |
| `churn` | 0.90–2.48 (1.40) | 0.95–2.87 (1.65) | 1.07–1.61 (1.46) |
| `build` | 0.94–2.15 (1.33) | 1.00–2.41 (1.54) |  |

### Memory at 1M keys, heap bytes per key (path 300K, street 212K keys)

`uint64` values:

| keys | u64 | str | uuid | email | url | path | street |
|---|---|---|---|---|---|---|---|
| ordered / btree-map | 41 / 37 | 52 / 56 | 68 / 77 | 61 / 57 | 86 / 101 | 72 / 101 | 59 / 47 |
| | +11% | -7% | -12% | +7% | -15% | -29% | +26% |

String values:

| keys | u64 | str | uuid | email | url | path | street |
|---|---|---|---|---|---|---|---|
| ordered / btree-map | 41 / 48 | 56 / 68 | 84 / 89 | 67 / 69 | 100 / 112 | 80 / 112 | 59 / 58 |
| | -15% | -18% | -6% | -3% | -11% | -29% | +2% |

## Several values per key, `uint64`

### Against `btree-sets`
| operation | 4K | 16K | 256K |
|---|---|---|---|
| `valuesFor` | 2.24–4.95 (3.13) | 2.22–4.20 (3.16) | 1.77–3.56 (2.27) |
| `valuesBetween` | 2.01–2.78 (2.32) | 1.81–2.16 (2.13) | 2.25–3.18 (2.80) |
| `prefix` | 1.38–2.76 (2.01) | 1.60–2.29 (2.27) | 2.08–3.01 (2.88) |
| `churn` | 1.68–3.18 (2.21) | 1.75–3.45 (2.35) | 1.37–1.88 (1.70) |
| `build` | 1.72–2.87 (2.09) | 1.80–3.10 (2.33) |  |

### Against `map-sets`
| operation | 4K | 16K | 256K |
|---|---|---|---|
| `valuesFor` | 1.15–2.95 (1.71) | 0.96–2.20 (1.50) | 1.08–2.56 (1.56) *1 |
| `valuesBetween` | 16.34–20.51 (16.94) | 47.81–53.68 (52.64) |  |
| `prefix` | 9.40–401 (223) | 8.23–1030 (438) |  |
| `churn` | 0.59–1.13 (0.79) | 0.64–1.18 (0.86) *2 | 0.63–0.98 (0.82) |
| `build` | 0.61–1.08 (0.77) | 0.62–1.11 (0.83) *1 |  |

### Against `hashed`
| operation | 4K | 16K | 256K |
|---|---|---|---|
| `valuesFor` | 0.50–0.96 (0.66) *1 | 0.47–0.94 (0.67) | 0.54–1.13 (0.73) |
| `valuesBetween` | 15.38–19.07 (16.05) | 50.83–58.16 (55.11) |  |
| `prefix` | 7.62–438 (239) | 6.93–1099 (439) |  |
| `churn` | 0.55–1.06 (0.72) | 0.52–1.06 (0.72) | 0.53–0.89 (0.71) *3 |
| `build` | 0.56–1.02 (0.72) | 0.54–0.99 (0.72) |  |

## Findings

1. **The credo's unique-profile range claim does not hold.** `valuesBetween` against `btree-map` is at
   0.26-0.41 at 4K and 16K and 0.43-0.55 at 256K; `prefix` has medians of 0.62-0.77 (worst with `str` keys,
   0.27). The B-tree keeps a run of keys in one node, the ART walks one leaf per key.
2. **Memory against `btree-map` is mixed**: ahead for uuid, url and path, behind for `u64` (41 / 37),
   email and street.
3. **Point, churn and build in the unique profile win everywhere that matters**: lookups 1.1-6.2, churn
   0.9-3.0, build 0.94-2.5; only path and url at 4K are close to level.
4. **Against `btree-sets` (several values) it wins every operation**, medians 1.7-3.2.
5. **Against `hashed` it sits where the credo wants it**: 0.5-1.1 on point operations, 15-55 times on range queries.
6. **Against `map-sets` it wins lookups and ranges but loses churn and build** (medians 0.77-0.86).
7. Not measured: 1M keys and the parallel regime against the competitors.
