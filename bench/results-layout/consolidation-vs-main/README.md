# node-layout against main (2026-09-30)

The first head-to-head of the whole `node-layout` branch against `main`: the
16-byte node header, flat leaves with small values, the leaf class in the
kind, and pessimistic path compression with leaves that hold only their key
rest, including the two fixes of the delete path (`fe6151c`, `970c08c`).

- **Candidate:** `node-layout` at `970c08c`.
- **Baseline:** `main` at `818e861`, copied into the bench with
  `cmd/mkbaseline -ref main`.
- **Driver:** `5e7f7a5` and rtcompare v0.8.0, the dev suite, serial,
  4K, 16K and 256K keys, a first stage of 6 processes and at most 12, memory at
  1M keys (3 rounds). `uint64` values and string values (`-tags strvals`).
- **Machine:** Ryzen 9 7900, **native Windows** (the drivers ran as Windows
  programs, not in WSL2): 1 h 13 min and 1 h 20 min. Numbers from WSL2 runs are
  not comparable in absolute terms.

```sh
bench -suite dev -vs baseline -sizes 4096,16384,262144 -minprocs 6 -maxprocs 12 \
  -memn 1048576 -memrounds 3
```

Speed is relative to the baseline (above 1: the candidate is faster). Each
cell gives the range over the key kinds and the median in brackets; `*n` marks
n comparisons whose interval is wider than ±2 points and ±10% of the
difference. 36 of 182 comparisons with `uint64` values and 23 of 182 with
strings are not precise, mostly churn at 256K keys.

## Results with `uint64` values

| values | operation | 4K | 16K | 256K |
|---|---|---|---|---|
| multi | `valuesFor` | 0.92–1.00 (0.98) *2 | 0.96–1.01 (0.98) | 1.07–1.17 (1.12) *2 |
| multi | `valuesBetween` | 0.94–0.99 (0.98) *5 | 0.96–1.00 (1.00) | 1.08–1.18 (1.15) |
| multi | `prefix` | 0.94–0.98 (0.95) *2 | 0.94–0.99 (0.98) *1 | 1.05–1.19 (1.13) *1 |
| multi | `churn` | 0.93–1.00 (0.94) | 0.90–1.02 (0.96) | 0.90–1.04 (0.95) *3 |
| multi | `build` | 0.91–0.98 (0.94) | 0.91–1.01 (0.95) | |
| unique | `valuesFor` | 0.98–1.03 (0.98) *4 | 0.95–1.03 (0.99) *1 | 1.10–1.40 (1.17) *3 |
| unique | `valuesBetween` | 0.94–0.98 (0.96) *1 | 0.96–1.01 (0.97) *1 | 1.04–1.44 (1.19) *1 |
| unique | `prefix` | 0.92–0.97 (0.94) | 0.94–0.99 (0.96) *2 | 1.01–1.51 (1.11) *2 |
| unique | `churn` | 0.95–1.03 (0.98) | 0.94–1.06 (1.01) | 1.02–1.08 (1.06) *4 |
| unique | `build` | 0.91–1.03 (0.96) *1 | 0.93–1.03 (0.97) | |

## Results with string values

| values | operation | 4K | 16K | 256K |
|---|---|---|---|---|
| multi-str | `valuesFor` | 0.95–1.01 (0.98) *1 | 0.97–1.03 (0.99) | 1.02–1.09 (1.06) |
| multi-str | `valuesBetween` | 0.87–0.94 (0.92) *2 | 0.88–0.96 (0.93) | 0.95–1.03 (0.98) |
| multi-str | `prefix` | 0.89–0.95 (0.93) *1 | 0.91–0.97 (0.94) *1 | 0.96–1.04 (0.98) *1 |
| multi-str | `churn` | 0.91–1.00 (0.99) | 0.96–1.03 (0.98) | 0.90–1.02 (0.95) *6 |
| multi-str | `build` | 0.94–0.99 (0.97) | 0.96–1.02 (0.98) | |
| unique-str | `valuesFor` | 0.90–1.00 (0.96) *2 | 0.92–1.02 (0.96) *1 | 1.01–1.14 (1.04) |
| unique-str | `valuesBetween` | 0.78–0.86 (0.82) | 0.77–0.90 (0.85) *2 | 0.96–1.00 (0.97) *1 |
| unique-str | `prefix` | 0.83–0.91 (0.90) | 0.85–0.94 (0.90) *1 | 0.94–1.04 (0.99) *1 |
| unique-str | `churn` | 0.91–0.97 (0.96) | 0.95–1.01 (0.96) | 0.97–1.07 (1.01) *3 |
| unique-str | `build` | 0.89–0.96 (0.92) | 0.91–0.97 (0.94) | |

## Memory at 1M keys

Heap bytes per key, candidate against `main` (path 300K, street 212K keys).

| values | u64 | str | uuid | email | url | path | street |
|---|---|---|---|---|---|---|---|
| multi, `uint64` | 117 / 149 | 129 / 177 | 147 / 193 | 137 / 176 | 162 / 230 | 149 / 233 | 82 / 119 |
| | −21% | −27% | −24% | −22% | −30% | −36% | −31% |
| unique, `uint64` | 41 / 73 | 52 / 102 | 68 / 117 | 61 / 100 | 86 / 155 | 72 / 157 | 59 / 96 |
| | −44% | −49% | −42% | −39% | −45% | −54% | −39% |
| multi, strings | 250 / 251 | 261 / 279 | 277 / 294 | 273 / 277 | 296 / 332 | 282 / 334 | 168 / 174 |
| | ±0% | −6% | −6% | −1% | −11% | −16% | −3% |
| unique, strings | 105 / 105 | 116 / 134 | 132 / 149 | 128 / 132 | 150 / 187 | 137 / 189 | 123 / 128 |
| | ±0% | −13% | −11% | −3% | −20% | −28% | −4% |

## Findings

1. **Memory is the branch's win.** With `uint64` values every kind needs
   21-54% less than `main`. With string values, which keep their values in set
   leaves, the saving shrinks to 0-28%, and only the long keys (url, path) save
   more than 10%.
2. **Out of the cache, it is faster.** From 256K keys up, lookups and ranges
   with `uint64` values gain 10-50% (medians 1.11-1.19). With string values only
   lookups gain (1.04-1.06); ranges and prefix searches are at 0.97-0.99.
3. **Small sizes cost little with `uint64` values.** At 4K and 16K keys the
   medians lie at 0.94-1.01; the worst cells (churn and build, 0.90-0.93) are
   within the regression budget of 15%, but the credo's "small sizes never get
   worse" still does not hold for churn and build with several values per key
   (medians 0.94-0.96).
4. **String values lose in range queries at small sizes**, the worst result of
   the branch: `valuesBetween` and `prefix` at 4K and 16K keys have medians
   0.82-0.94, single cells down to 0.77 (unique-str, 16K). Range queries with
   `uint64` values do not have this (0.94-1.00), so the cost sits in how a scan
   walks set leaves, not in the paths. String values are what typed leaves
   (next step) are for.
5. **Speed at 256K and 1M keys is below, in the parallel regime.**

## Out of the cache, in the parallel regime

The same candidate and baseline, 256K and 1M keys, the kinds `u64`, `uuid`,
`url` and `path` (`path` has no 1M keys), no `build` (it stops at 64K keys),
in the **parallel regime**: 8 processes at a time, each with GOMAXPROCS 3,
sharing the caches and the memory bandwidth (`parallel-uint64/`,
`parallel-strvals/`, 66 and 46 minutes, 2026-09-30 evening). The first stage of
8 processes sufficed for 31 of the 52 comparisons with `uint64` values and 33
with strings; the others ran 16 to 40 processes. All 104 are precise. These figures are **not comparable
with the serial ones above** and must not be put in one table with them.

```sh
bench -suite parallel -vs baseline -keys u64,uuid,url,path -minprocs 8 -maxprocs 160
```

| values | operation | 256K | 1M |
|---|---|---|---|
| multi | `valuesFor` | 1.00–1.10 (1.06) | 0.99–1.08 (1.02) |
| multi | `valuesBetween` | 1.09–1.16 (1.14) | 1.10–1.18 (1.17) |
| multi | `prefix` | 1.10–1.13 (1.12) | 1.15–1.17 (1.16) |
| multi | `churn` | 0.88–1.04 (0.96) | 0.89–1.02 (0.96) |
| unique | `valuesFor` | 1.00–1.11 (1.06) | 0.99–1.08 (1.02) |
| unique | `valuesBetween` | 1.02–1.09 (1.05) | 1.01–1.05 (1.02) |
| unique | `prefix` | 1.05–1.06 (1.05) | 1.06–1.09 (1.08) |
| unique | `churn` | 1.00–1.08 (1.04) | 0.99–1.07 (1.02) |
| multi-str | `valuesFor` | 0.96–1.07 (1.04) | 0.97–1.06 (1.01) |
| multi-str | `valuesBetween` | 0.94–0.99 (0.97) | 0.92–0.98 (0.96) |
| multi-str | `prefix` | 0.97–0.98 (0.98) | 0.93–1.00 (0.97) |
| multi-str | `churn` | 0.97–1.05 (1.00) | 0.96–1.02 (0.98) |
| unique-str | `valuesFor` | 0.96–1.10 (1.06) | 0.96–1.08 (1.01) |
| unique-str | `valuesBetween` | 0.91–1.02 (0.96) | 0.88–0.96 (0.95) |
| unique-str | `prefix` | 0.95–0.97 (0.96) | 0.90–0.97 (0.94) |
| unique-str | `churn` | 1.00–1.08 (1.02) | 0.99–1.05 (1.00) |

6. **`uint64` values keep their win at 1M keys**: ranges and prefix searches
   are 1.02-1.17 times as fast, lookups 1.02, and churn is level (0.96-1.02),
   with `u64` multi churn as the weak spot (0.88-0.89).
7. **String values lose in ranges at every size**, now also at 256K and 1M
   keys: `valuesBetween` and `prefix` have medians 0.94-0.98 and are worst with
   uuid keys (0.88-0.94). It is the same effect as at 4K and 16K, so it does
   not go away with the size.
8. **`competitors-partial/`** holds the unfinished run against all
   competitors (`u64` and `str`, several values per key, 4K and 16K keys); it
   stopped because it would have needed seven hours, and should be repeated
   on the final code.
