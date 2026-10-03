# The page tree for string values (2026-10-03/04)

Report: [docs/redesign/step3-tree-pages.md](../../../docs/redesign/step3-tree-pages.md). Ryzen 9 7900,
Windows, native, `-suite dev -minprocs 6 -maxprocs 12 -memn 262144 -memrounds 3`, three sizes (4,096,
16,384 and the corpus), `-tags strvals`. `pc/<run>/` holds `speed-summary.md`, `mem-summary.md`, the raw
`speed.jsonl`/`mem.jsonl` and `run.json`.

| run | commit | scenario |
|---|---|---|
| `lp1-*` | b8e4111 | `ordered-lpage` against `btree-map` (unique) or `btree-sets` (multi) |
| `lp3-uni-*` | 0a517ae | `ordered-lpage` and `-zc`, unique |
| `lp4-multi-*` | 7717957 | `ordered-lpage-mv` and `-mvzc`, multi |

Every cell: how many times as fast the candidate is as today's `ordered`, from `1 / (A speed vs B)` of the
summary (above 1: faster than today's tree). `*`: the interval is wider than asked; `?`: the difference
does not clear the noise floor. "Against X": the candidate against X through the common side `ordered`
(ratio of the two pairs), only indicative. **The memory tables are not like for like**: only the page
tree holds the string bytes in its heap (10 and 28 bytes a key on `street` for `unique` and `multi`,
15 and 68 on `dirs`); the others hold 16-byte headers into a shared buffer.

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
