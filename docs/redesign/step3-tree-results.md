# Step 3.5: the measurement of the tree with single-key pages for every value type

Measured 2026-10-05 on the PC (Ryzen 9 7900, Windows native, rtcompare, `-suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3`), 18 runs, 10:17 to 12:04. No M1 (not available).
Code: commit fe120f5 (step 3.5.5). Words as in [GLOSSARY.md](GLOSSARY.md). Raw: `bench/results-layout/step3-tree/pc-m35/` (summaries, `run.json`, the log of every run, `run-m35.cmd`).

Two references, built with `mkbaseline` as the candidate "baseline":

- **X: `node-layout`** (7b8a8d8): the reference of gate 3 (credo 1, memory not above it).
- **Y: the tree before 3.5** (tag `before-mk-pages-removal`: multi-key pages, flat leaf, typed leaf): what the removal of the multi-key pages and the change to pages cost.

Cases: `string -> {uint64}` and `string -> {*T}` on `street` and `dirs`, real and one value per key; `uint64 -> {*T}` real and one value per key (4,096, 16,384, 262,144 keys).
"A / B" below is the speed of `Ordered` against the reference: below 1.00 is slower.

## Memory (heap bytes per entry, `Ordered` against `node-layout`, then against the tree before 3.5)

| case | `Ordered` | `node-layout` | before 3.5 | prediction |
|---|---:|---:|---:|---|
| `uint64`, street real | 90 | 90 | 90 | 89 |
| `uint64`, dirs real | 108 | 106 | 106 | 107 |
| `uint64`, street, one value per key | 67 | 67 | **40** | 66.6 |
| `uint64`, dirs, one value per key | 77 | 74 | **63** | 73.4 |
| `*T`, street real | 90 | 86 | 86 | 90 |
| `*T`, dirs real | 108 | 106 | 106 | 106 |
| `*T`, street, one value per key | 67 | **58** | 59 | +20 to +41 % against the typed leaf |
| `*T`, dirs, one value per key | 77 | **68** | 68 | same |
| `uint64` keys, `*T`, real | 125 | 124 | | |
| `uint64` keys, `*T`, one value per key | 53 | **45** | | |

- The predictions of the design note hold (`uint64` real 90 and 108, one value 67; `*T` real 90 and 108). The `btree-sets` need 285 to 352 bytes where `Ordered` needs 90 to 125.
- **Gate 3, memory not above `node-layout`: not met.** Two cells are equal (`uint64` street, real and one value), `dirs` and the real `*T` cases are 1 to 4 bytes above (+1 to +5 %), and the
  `*T` maps with one value per key are 8 to 9 bytes above (+13 to +18 %). The 9 bytes are the known price of the smallest class (32 bytes for a page that held a 24-byte typed leaf); a 24-byte class is the candidate
  after 0.8 (decided 2026-10-05: not now). The 1 to 4 bytes on `dirs` and on `*T` real I have not traced to their cause yet.
- Against the tree before 3.5, the single-value maps of `uint64` are **27 and 14 bytes larger**: that tree held them in multi-key pages (40 and 63 bytes). That is the part that comes back in step 4.
- Memory after removing half the entries (not a gate) follows the same order: 35 against 34/35, `*T` 35 against 30.

## Speed against `node-layout` (X), `Ordered` / `node-layout`

| case | `valuesFor` | `valuesBetween` | `prefix` | `churn` | `build` |
|---|---|---|---|---|---|
| `uint64`, street real | 0.94 to 0.97 | 0.88 to 0.96 | 0.88 to 0.93 | 0.91 to 0.95 | 0.89 to 0.91 |
| `uint64`, dirs real | 0.96 to 0.98 | 0.87 to 0.95 | 0.85 to 0.93 | 0.94 to 0.97 | 0.92 to 0.94 |
| `*T`, street real | 0.93 to 0.95 | 0.89 to 0.96 | 0.87 to 0.96 | 0.93 to 0.99 | 0.95 |
| `*T`, dirs real | 0.96 to 0.99 | 0.90 to 0.97 | 0.85 to 0.94 | 0.97 to 0.98 | 0.97 |
| `uint64`, street, one value | 0.95 to 0.97 | 0.86 to 0.93 | 0.86 to 0.90 | 0.90 to 0.96 | 0.85 to 0.89 |
| `*T`, street, one value | 0.90 to 0.94 | 0.84 to 0.88 | 0.83 to 0.87 | 0.90 to 0.95 | 0.87 to 0.91 |
| `*T`, dirs, one value | 0.95 to 0.97 | 0.87 to 0.91 | 0.82 to 0.89 | 0.95 to 0.96 | 0.93 to 0.94 |
| `uint64` keys, `*T`, real | 0.95 to 0.96 | 0.94 to 1.07 | | 0.86 to 0.91 | 0.90 to 0.91 |
| `uint64` keys, `*T`, one value | **0.81 to 0.87** | **0.81 to 0.88** | | **0.83 to 0.91** | **0.80 to 0.83** |

- Reported, not gating: **no cell is below 0.70** (the lowest is 0.80, `uint64` keys with one value per key, `build` at 4,096). `Ordered` is 3 to 17 % slower than `node-layout` on almost every cell of
  `street` and `dirs`; the gap is largest for the range scans at 4,096 and 16,384 entries and closes towards the largest size (0.90 to 1.00, partly within the noise of the 8-process runs).
- The 3 to 17 % come from the page and not from the tree: against the tree before 3.5 (Y) the real cases are 0.87 to 1.06, and the `*T` real cases of `dirs` are 0.97 to 1.04 (churn, build).
  The step from `node-layout` to the tree before 3.5 had already cost the rest (the 9-bit remainder, step 3.4b; the build/churn 2 to 3 % has not been traced).
- The `uint64` keys with one value per key are the weakest cell: 53 bytes against 45 and 15 to 20 % slower. These are `uint64` keys with a remainder of at most 7 bytes in a 32-byte page, the
  worst relation of head to payload; the multi-key page (step 4) is meant for exactly this case.

## Credo 1 (gate 3, hard)

- **Several values per key (real), against `btree-sets`: met in every cell.** Lookups 1.8 to 2.5 times, ranges 1.7 to 2.4 times, `churn` 1.26 at the lowest (the largest size), `build` 1.8 times.
- **One value per key, against `btree-map`: not met** (not new: gate 3 was already marked not met for this profile in step3-skmv-results.md). `valuesFor` is 1.1 to 1.2 times at the largest size (within the noise at some),
  `churn` and `build` 0.87 to 1.20 (street above 1, `dirs` below 1), **range scans 0.19 to 0.57 times**. These are the cells of step 4 (credo 2).
- `hashed` was not in these runs (not needed for the question); the range-scan comparison with `hashed` of gate 3 is still to do before the gate is closed.

## The tree before 3.5 (Y)

- Real maps: the removal of the multi-key pages and of the flat and typed leaves cost nothing worth the name: **0.87 to 1.06** over all cells (the low ones are `prefix` at 4,096 entries), memory equal for `uint64` and 2 to 4 bytes up for `*T`.
  This was the expectation: on real data the multi-key pages had been given up by the fall back in every subtree (step3-tree-pages.md).
- One value per key, `uint64`: **the multi-key pages are missing**: `valuesBetween` 0.20 to 0.32 and `prefix` 0.15 to 0.48 of the old tree (cells below 0.70, cause: the page holds one entry where the multi-key
  page held many), `churn` 0.71 to 0.82 on `street` (0.83 to 1.00 on `dirs`), memory 67 against 40. `valuesFor` stays at 0.90 to 1.14 up to 16,384 entries and falls to 0.73 on `street` at 212,449 (a cell below 0.70; probably the cache misses of one page per entry against one per many, not examined); `build` is 0.98 on `street` and 1.19 to 1.21 on `dirs`.
  This is the price of the decision "multi-key pages out first, redesign in step 4"; the task of step 4 is to win these cells back (gate 4: credo 2 for one value per key) without touching the real ones.
- `*T`, one value per key: `Ordered` against the tree before 3.5: 0.91 to 1.02 for lookup, `churn` and `build`, 0.89 to 0.95 for the scans: the 32-byte class costs 9 bytes, the speed little.

## What this says

1. **The page for every type of value works and costs nothing on real data** against the old tree; the predictions of memory hold.
2. **Gate 3 is not met in two places**, both with a known cause: memory of the `*T` maps with one value per key (+13 to +18 %, 32-byte class), and credo 1 for one value per key (needs the multi-key page).
   The 1 to 4 bytes more on `dirs` and on the real `*T` cases and the 3 to 17 % against `node-layout` are open and are to be traced (the page head against the flat leaf, the 9-bit remainder; M1 and
   a profile of `valuesBetween` at 4,096 entries would tell).
3. Step 4 (MKSV) is the natural next step: it is where the single-value cells (credo 2, memory 67 against 40, the weak `uint64` keys) can be repaired.
