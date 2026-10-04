# The page tree for string values (2026-10-03/04)

Report: [docs/redesign/step3-tree-pages.md](../../../docs/redesign/step3-tree-pages.md). Ryzen 9 7900,
Windows, native, `-suite dev -minprocs 6 -maxprocs 12 -memn 262144 -memrounds 3`, three sizes (4,096,
16,384 and the corpus), `-tags strvals`. `pc/<run>/` holds `speed-summary.md`, `mem-summary.md`, the raw
`speed.jsonl`/`mem.jsonl` and `run.json`. The M1 jobs `l1` to `l4` of `bench/remote/queue.txt` repeat the
`lp8` runs; their results will be on the branch `arm-results`. The code that produced these runs (the
candidates `ordered-lpage*`, `internal/artstr`) is on the branch `mkmv-experiment`, not on `cacheline`.

| run | code | scenario |
|---|---|---|
| `lp8-*` | `c4dc4b6` (final) | `unique` and `multi` of `street` and `dirs`; `ordered-lpage`, `-zc`, `-mv`, `-mvzc` against `btree-map` or `btree-sets` |
| `lp9-*` | `c4dc4b6` | `street` repeated (`mv`; the same numbers as `lp8` to within 0.03) and `dirs` with a header of 64 |
| `lp6-multi-*` | `ef06198` | header size 32, 48 and 64 (`LPAGE_MAXHEADER`), larger minimum header (`LPAGE_MINHEADER=24`: no gain); the `dirs` run with 48 crashed in its memory phase (the bug of the report) and has no memory table |
| `lp4-*` | `7717957` | first several-values runs (3 classes, header 24; `-h16`, `-h32`: 16 and 32) |
| `lp3-uni-*` | `0a517ae` | first zero-copy runs |
| `lp1-*` | `b8e4111` | first runs (one value per entry, header 24, 3 classes) |

Every cell: how many times as fast the candidate is as today's `ordered`, from `1 / (A speed vs B)` of the
summary (above 1: faster than today's tree). `*`: the interval is wider than asked; `?`: the difference
does not clear the noise floor. "Against X": the candidate against X through the common side `ordered`
(ratio of the two pairs), only indicative. **The memory tables are not like for like**: only the page
tree holds the string bytes in its heap (10 and 28 bytes a key on `street` for `unique` and `multi`,
15 and 68 on `dirs`); the others hold 16-byte headers into a shared buffer.

# Final runs (`lp8`)

### street, one value per key (final)

| operation | lp 4096 | lp 16384 | lp 212449 | lp-zc 4096 | lp-zc 16384 | lp-zc 212449 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.52 | 0.63 | 1.30* | 0.63 | 0.75 | 1.52 |
| valuesBetween | 1.61 | 1.59 | 2.50 | 2.27 | 2.38 | 3.57 |
| prefix | 0.88 | 1.25 | 2.56 | 1.09 | 1.82 | 4.00 |
| churn | 0.66 | 0.81 | 1.00? | 0.49 | 0.62 | 0.87 |
| build | 0.53 | 0.60 | - | 0.40 | 0.45 | - |

Against `btree-map` (above 1: faster than it):

| operation | lp 4096 | lp 16384 | lp 212449 | lp-zc 4096 | lp-zc 16384 | lp-zc 212449 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.01 | 1.18 | 1.39* | 1.23 | 1.40 | 1.62 |
| valuesBetween | 1.03 | 1.03 | 1.55 | 1.45 | 1.55 | 2.21 |
| prefix | 0.84 | 0.99 | 1.67 | 1.04 | 1.44 | 2.60 |
| churn | 0.77 | 0.90* | 1.20? | 0.57 | 0.69* | 1.04 |
| build | 0.68* | 0.77 | - | 0.52* | 0.58 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 67 | 63 | +39 ms | 35 |
| btree-map | 66 | 49 | +31 ms | 34 |
| ordered-lpage | 30 | 4 | +1 ms | 20 |
| ordered-lpage-zc | 30 | 4 | +4 ms | 20 |

### dirs, one value per key (final)

| operation | lp 4096 | lp 16384 | lp 86215 | lp-zc 4096 | lp-zc 16384 | lp-zc 86215 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.63 | 0.73 | 1.08 | 0.75 | 0.85 | 1.23 |
| valuesBetween | 1.10 | 1.15 | 1.37* | 1.69 | 1.85 | 2.00 |
| prefix | 1.03 | 1.18* | 1.43 | 1.59 | 2.04 | 2.56 |
| churn | 0.79 | 0.89 | 0.99*? | 0.63 | 0.74 | 0.87 |
| build | 0.58 | 0.66 | - | 0.50 | 0.54 | - |

Against `btree-map` (above 1: faster than it):

| operation | lp 4096 | lp 16384 | lp 86215 | lp-zc 4096 | lp-zc 16384 | lp-zc 86215 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.91 | 1.08 | 1.12 | 1.08 | 1.25 | 1.28 |
| valuesBetween | 0.82 | 0.91 | 1.27** | 1.27 | 1.46 | 1.86* |
| prefix | 0.88 | 1.02** | 1.13* | 1.35 | 1.78* | 2.03* |
| churn | 0.72 | 0.80* | 0.99*?? | 0.58 | 0.66* | 0.87? |
| build | 0.57 | 0.66? | - | 0.49 | 0.53? | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 78 | 72 | +19 ms | 41 |
| btree-map | 108 | 49 | +14 ms | 55 |
| ordered-lpage | 54 | 9 | +3 ms | 36 |
| ordered-lpage-zc | 54 | 9 | +2 ms | 36 |

### street, the natural mix (final)

| operation | lp 4096 | lp 16384 | lp 212449 | lp-mv 4096 | lp-mv 16384 | lp-mv 212449 | lp-mvzc 4096 | lp-mvzc 16384 | lp-mvzc 212449 |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.01? | 1.00? | 1.00? | 0.49 | 0.57 | 1.08 | 0.59 | 0.68 | 1.18* |
| valuesBetween | 0.98 | 0.99 | 1.00? | 1.05 | 1.12 | 1.56* | 1.49 | 1.67 | 2.17 |
| prefix | 0.99? | 0.98? | 1.00? | 0.76 | 0.93 | 1.54 | 1.01 | 1.39 | 2.22 |
| churn | 1.00? | 1.00? | 1.02*? | 0.50 | 0.58 | 0.88 | 0.37 | 0.47 | 0.79* |
| build | 0.99 | 0.89 | - | 0.44 | 0.52 | - | 0.35 | 0.40 | - |

Against `btree-sets` (above 1: faster than it):

| operation | lp 4096 | lp 16384 | lp 212449 | lp-mv 4096 | lp-mv 16384 | lp-mv 212449 | lp-mvzc 4096 | lp-mvzc 16384 | lp-mvzc 212449 |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| valuesFor | 2.60? | 2.51? | 1.73? | 1.27 | 1.43 | 1.86 | 1.52 | 1.72 | 2.04* |
| valuesBetween | 2.97 | 2.76 | 3.00? | 3.19 | 3.13 | 4.69* | 4.52 | 4.65 | 6.52 |
| prefix | 3.10? | 3.36? | 3.16? | 2.37 | 3.18 | 4.86 | 3.16 | 4.76 | 7.02 |
| churn | 1.78? | 1.79? | 1.38*? | 0.88 | 1.05 | 1.19 | 0.66 | 0.85 | 1.07* |
| build | 1.74 | 1.61 | - | 0.78 | 0.93 | - | 0.62 | 0.71 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 119 | 111 | +56 ms | 61 |
| btree-sets | 367 | 336 | +92 ms | 185 |
| ordered-lpage | 119 | 111 | +55 ms | 61 |
| ordered-lpage-mv | 88 | 44 | +14 ms | 50 |
| ordered-lpage-mvzc | 88 | 44 | +11 ms | 50 |

### dirs, the natural mix (final)

| operation | lp 4096 | lp 16384 | lp 86215 | lp-mv 4096 | lp-mv 16384 | lp-mv 86215 | lp-mvzc 4096 | lp-mvzc 16384 | lp-mvzc 86215 |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.00? | 1.00? | 1.00? | 0.56 | 0.65 | 0.92 | 0.70 | 0.78 | 1.10* |
| valuesBetween | 0.98*? | 0.99? | 1.00*? | 0.73 | 0.85 | 1.25 | 1.18* | 1.37 | 1.79 |
| prefix | 0.97 | 0.99? | 0.94* | 0.67 | 0.83 | 1.06* | 1.11 | 1.49* | 1.79 |
| churn | 0.99? | 1.00 | 1.06* | 0.58 | 0.71 | 0.87* | 0.46 | 0.60 | 0.78 |
| build | 0.96 | 0.93 | - | 0.49 | 0.57 | - | 0.41 | 0.45 | - |

Against `btree-sets` (above 1: faster than it):

| operation | lp 4096 | lp 16384 | lp 86215 | lp-mv 4096 | lp-mv 16384 | lp-mv 86215 | lp-mvzc 4096 | lp-mvzc 16384 | lp-mvzc 86215 |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.91? | 1.87? | 1.58? | 1.07 | 1.21 | 1.45 | 1.35 | 1.46 | 1.74* |
| valuesBetween | 3.75*? | 3.13? | 3.28*? | 2.79 | 2.70 | 4.10 | 4.49* | 4.33 | 5.86 |
| prefix | 4.91 | 3.80? | 5.49* | 3.37 | 3.20 | 6.19* | 5.62 | 5.73* | 10.39 |
| churn | 1.41? | 1.44 | 1.37* | 0.83 | 1.02 | 1.12* | 0.66 | 0.87 | 1.01 |
| build | 1.39 | 1.38 | - | 0.71 | 0.85 | - | 0.60 | 0.67 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 149 | 138 | +27 ms | 75 |
| btree-sets | 421 | 352 | +36 ms | 210 |
| ordered-lpage | 150 | 138 | +27 ms | 75 |
| ordered-lpage-mv | 139 | 60 | +10 ms | 80 |
| ordered-lpage-mvzc | 139 | 60 | +10 ms | 79 |

# History (older code)

### street, one value per key (unique), copy-out strings (commit b8e4111)

| operation | lp 4096 | lp 16384 | lp 212449 |
|---|--:|--:|--:|
| valuesFor | 0.59 | 0.70 | 1.12* |
| valuesBetween | 1.30 | 1.35 | 1.49* |
| prefix | 0.86 | 1.11 | 1.59 |
| churn | 0.47 | 0.55 | 0.83 |
| build | 0.40 | 0.46 | - |

Against `btree-map` (above 1: faster than it):

| operation | lp 4096 | lp 16384 | lp 212449 |
|---|--:|--:|--:|
| valuesFor | 1.15 | 1.35 | 1.25** |
| valuesBetween | 0.82 | 0.86 | 0.90** |
| prefix | 0.82 | 0.88 | 1.00 |
| churn | 0.54 | 0.67* | 0.88 |
| build | 0.50 | 0.60 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 67 | 63 | +37 ms | 35 |
| btree-map | 66 | 49 | +29 ms | 34 |
| ordered-lpage | 40 | 8 | +4 ms | 23 |

### street, one value per key, copy-out and zero-copy strings

| operation | lp 4096 | lp 16384 | lp 212449 | lp-zc 4096 | lp-zc 16384 | lp-zc 212449 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.58 | 0.68 | 1.18* | 0.72 | 0.83 | 1.32* |
| valuesBetween | 1.35 | 1.39 | 1.59* | 1.92 | 2.08 | 2.22* |
| prefix | 0.92 | 1.15* | 1.69* | 1.18 | 1.72 | 2.56 |
| churn | 0.48 | 0.54 | 0.81 | 0.40 | 0.48 | 0.75 |
| build | 0.40 | 0.47 | - | 0.34 | 0.39 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 67 | 63 | +41 ms | 35 |
| ordered-lpage | 40 | 8 | +6 ms | 23 |
| ordered-lpage-zc | 40 | 8 | +4 ms | 23 |

### dirs, one value per key (unique), copy-out strings

| operation | lp 4096 | lp 16384 | lp 86215 |
|---|--:|--:|--:|
| valuesFor | 0.68 | 0.75 | 0.90* |
| valuesBetween | 1.10 | 1.14 | 1.15 |
| prefix | 1.00? | 1.14* | 1.30 |
| churn | 0.79 | 0.81 | 0.91* |
| build | 0.58 | 0.64 | - |

Against `btree-map` (above 1: faster than it):

| operation | lp 4096 | lp 16384 | lp 86215 |
|---|--:|--:|--:|
| valuesFor | 0.97 | 1.13 | 1.22** |
| valuesBetween | 0.82 | 0.89 | 0.91* |
| prefix | 0.84? | 0.98** | 0.90 |
| churn | 0.73 | 0.80 | 0.86* |
| build | 0.56 | 0.64? | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 78 | 72 | +17 ms | 41 |
| btree-map | 107 | 49 | +11 ms | 54 |
| ordered-lpage | 64 | 11 | +4 ms | 44 |

### dirs, one value per key, copy-out and zero-copy strings

| operation | lp 4096 | lp 16384 | lp 86215 | lp-zc 4096 | lp-zc 16384 | lp-zc 86215 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.66 | 0.75 | 1.01*? | 0.79 | 0.88 | 1.10* |
| valuesBetween | 1.12 | 1.15 | 1.20* | 1.72 | 1.92 | 1.85 |
| prefix | 1.05 | 1.16* | 1.30 | 1.64 | 2.08 | 2.33 |
| churn | 0.80 | 0.85 | 0.91 | 0.63 | 0.66 | 0.79 |
| build | 0.59 | 0.63 | - | 0.50 | 0.51 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 78 | 72 | +18 ms | 41 |
| ordered-lpage | 63 | 11 | +3 ms | 43 |
| ordered-lpage-zc | 64 | 11 | +3 ms | 44 |

### street, the natural mix (multi), one value per entry in the pages (as decided)

| operation | lp 4096 | lp 16384 | lp 212449 |
|---|--:|--:|--:|
| valuesFor | 0.99? | 0.99? | 0.99*? |
| valuesBetween | 1.01? | 1.00? | 0.99*? |
| prefix | 1.02? | 1.00? | 0.99*? |
| churn | 1.00? | 1.00? | 1.04* |
| build | 0.99? | 0.88 | - |

Against `btree-sets` (above 1: faster than it):

| operation | lp 4096 | lp 16384 | lp 212449 |
|---|--:|--:|--:|
| valuesFor | 2.53? | 2.46? | 1.73*? |
| valuesBetween | 3.05? | 2.77? | 2.98*? |
| prefix | 3.18? | 3.46? | 3.08*? |
| churn | 1.82? | 1.81? | 1.38* |
| build | 1.76? | 1.59 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 119 | 111 | +54 ms | 61 |
| btree-sets | 367 | 337 | +91 ms | 185 |
| ordered-lpage | 119 | 111 | +53 ms | 61 |

### dirs, the natural mix (multi), one value per entry in the pages (as decided)

| operation | lp 4096 | lp 16384 | lp 86215 |
|---|--:|--:|--:|
| valuesFor | 0.99? | 0.99? | 1.01? |
| valuesBetween | 1.00*? | 0.99 | 0.99*? |
| prefix | 1.02 | 1.00? | 0.99? |
| churn | 1.00? | 1.00? | 1.01*? |
| build | 0.95 | 0.90 | - |

Against `btree-sets` (above 1: faster than it):

| operation | lp 4096 | lp 16384 | lp 86215 |
|---|--:|--:|--:|
| valuesFor | 1.89? | 1.87? | 1.76? |
| valuesBetween | 3.78*? | 3.11 | 3.07*? |
| prefix | 5.10 | 3.83? | 4.33? |
| churn | 1.44? | 1.50? | 1.26*? |
| build | 1.40 | 1.34 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 149 | 139 | +26 ms | 75 |
| btree-sets | 421 | 353 | +33 ms | 210 |
| ordered-lpage | 150 | 138 | +27 ms | 76 |

### street, the natural mix, multi-value entries inside the pages (experiment beyond the decision)

| operation | lp-mv 4096 | lp-mv 16384 | lp-mv 212449 | lp-mvzc 4096 | lp-mvzc 16384 | lp-mvzc 212449 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.52 | 0.61 | 0.99 | 0.63 | 0.72 | 1.08* |
| valuesBetween | 0.95 | 0.97* | 1.27* | 1.32 | 1.41 | 1.64* |
| prefix | 0.76 | 0.85 | 1.23* | 1.01? | 1.27 | 1.67 |
| churn | 0.43 | 0.52 | 0.78 | 0.36 | 0.44 | 0.74 |
| build | 0.39 | 0.45 | - | 0.34 | 0.38 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 119 | 111 | +56 ms | 61 |
| ordered-lpage-mv | 105 | 55 | +20 ms | 56 |
| ordered-lpage-mvzc | 105 | 55 | +22 ms | 56 |

### dirs, the natural mix, multi-value entries inside the pages

| operation | lp-mv 4096 | lp-mv 16384 | lp-mv 86215 | lp-mvzc 4096 | lp-mvzc 16384 | lp-mvzc 86215 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.56 | 0.63 | 0.84 | 0.70 | 0.77 | 0.94 |
| valuesBetween | 0.70 | 0.81 | 0.96* | 1.15 | 1.28 | 1.30* |
| prefix | 0.65 | 0.77 | 0.93* | 1.09 | 1.37* | 1.37 |
| churn | 0.56 | 0.64 | 0.78 | 0.46 | 0.56 | 0.70 |
| build | 0.47 | 0.53 | - | 0.41 | 0.43 | - |

Memory per key (heap / scannable / GC CPU per cycle / heap after removing half):

| candidate | heap B/key | scannable B/key | GC CPU | half |
|---|--:|--:|--:|--:|
| ordered | 149 | 139 | +27 ms | 76 |
| ordered-lpage-mv | 158 | 72 | +12 ms | 87 |
| ordered-lpage-mvzc | 156 | 72 | +13 ms | 87 |
