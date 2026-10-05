# Step 4.2: multi-key pages in the tree: the measurement of 2026-10-05

Design: [step4-mksv-design.md](step4-mksv-design.md). Words as in [GLOSSARY.md](GLOSSARY.md). Raw: `bench/results-layout/step4-tree/pc-m42/` (PC, Ryzen 9 7900, Windows native, rtcompare,
`-suite dev -minprocs 6 -maxprocs 8 -memn 262144 -memrounds 3`, 10 runs, 15:50 to 21:20). **The code measured is commit `a166cad`**; the reference ("baseline") is step 3.5, commit `fe120f5`
(`mkbaseline`). Cases: `string -> {uint64}` and `string -> {string}` on `street` and `dirs`, real and one value per entry (`single-value`), and `uint64 -> {uint64}` keys. `*T` is step 4.3.

## Result in one paragraph

**Memory is as predicted and much better than before; the speed of changes (`churn`, `build`) is far below the prediction, and the range scans are 1.1 to 2.9 times as fast as step 3.5
but still 0.3 to 0.7 of `btree-map`.** Gate 4 is not met: credo 2 (ranges at least `btree-map`) and credo 1 for `churn` and `build` with one value per entry. The cause of the first part
was found in a profile and is fixed after the measurement (below); the second part is a finding that needs the user's decision.

## Memory (heap bytes per entry, `Ordered` / step 3.5 / `btree-map` or `btree-sets`)

| case | `Ordered` | step 3.5 | competitor | prediction |
|---|--:|--:|--:|--:|
| `street` `uint64`, one value | **27** | 67 | 55 | 26.9 |
| `dirs` `uint64`, one value | **40** | 77 | 95 | 39.5 |
| `street` strings, one value | **31** | 69 | 66 | 30.9 (model) |
| `dirs` strings, one value | **51** | 87 | 107 | 50.3 |
| `uint64` keys, one value | **24** | 53 | 45 | 24.0 |
| `street` `uint64`, real | **73** | 90 | 285 | at most 58 + 15 (value overflow) |
| `dirs` `uint64`, real | **96** | 108 | 336 | at most 79 + 15 |
| `street` strings, real | 98 | 114 | 367 | |
| `dirs` strings, real | 153 | 165 | 422 | |
| `uint64` keys, real | 123 | 125 | 352 | |

- The model reproduces the build to the byte (`objstat`: 26.9, 39.7, 24.0). **Credo 2 for memory is met** (27 against 55, 40 against 95, 24 against 45) and credo 1's memory is not above `node-layout` by far.
- After removing every second key (heap per key of the corpus): 19, 29, 22 (strings), 20 against 28, 48, 34, 27 of `btree-map`; the pages shrink and merge (`objstat -removehalf`: 38 bytes per remaining key
  against 27 fresh).
- The real mix has pages too (21 % of the entries of `street` have several values): -19 % (`street`) and -11 % (`dirs`), as the model's upper bound said.
- GC: the scannable bytes of the single-value maps fall to 3 to 6 per entry (34 to 37 before), and a GC cycle takes 2 to 5 ms against 11 to 27.

## Speed against step 3.5 (`Ordered` / baseline; below 1 is slower)

| case | `valuesFor` | `valuesBetween` | `prefix` | `churn` | `build` |
|---|---|---|---|---|---|
| `uint64`, `street` real | 0.90 to 1.10 | 1.16 to 1.19 | 1.11 to 1.27 | 0.60 to 0.84 | 0.54 to 0.61 |
| `uint64`, `dirs` real | 0.96 to 1.01 | 1.09 to 1.11 | 1.14 to 1.24 | 0.69 to 0.84 | 0.62 to 0.68 |
| strings, `street` real | 0.96 to 1.07 | 1.07 to 1.15 | 1.07 to 1.16 | 0.61 to 0.87 | 0.54 to 0.61 |
| strings, `dirs` real | 0.97 to 1.03 | 1.04 to 1.08 | 1.05 to 1.10 | 0.71 to 0.86 | 0.63 to 0.69 |
| `uint64`, `street`, one value | 0.80 to 1.28 | 1.81 to 2.17 | 0.95 to 2.89 | **0.08 to 0.25** | **0.10 to 0.11** |
| `uint64`, `dirs`, one value | 0.87 to 1.05 | 1.67 to 1.80 | 1.74 to 2.44 | **0.14 to 0.28** | **0.17 to 0.19** |
| strings, `street`, one value | 0.93 to 1.35 | 1.46 to 1.71 | 1.02 to 1.88 | **0.08 to 0.23** | **0.09 to 0.10** |
| strings, `dirs`, one value | 0.95 to 1.08 | 1.34 to 1.44 | 1.35 to 1.60 | **0.14 to 0.28** | **0.17 to 0.19** |
| `uint64` keys, real | 0.88 to 1.01 | 1.01 to 1.02 | | 0.57 to 0.85 | 0.60 to 0.74 |
| `uint64` keys, one value | 0.53 to 0.97 | 1.13 to 1.52 | | **0.03 to 0.14** | **0.04 to 0.06** |

## Against the competitors (credo 1 and 2)

- **Real mix against `btree-sets`:** lookups 1.3 to 4.1, ranges 1.5 to 3.2, `churn` 0.91 to 2.0, `build` 0.88 to 1.9: credo 1 is **not met in four cells at 4,096 entries** (`build` 0.88 to 0.95 on `street` and `dirs`
  with strings and `uint64`; `churn` 0.91 and 0.97 with strings).
- **One value per entry against `btree-map`:** lookups 1.06 to 4.1 (met), **ranges 0.30 to 0.72 (credo 2 not met)**, **`churn` 0.05 to 0.30, `build` 0.07 to 0.19 (credo 1 not met)**.

## What was found, and what was changed after the measurement

1. **`churn` and `build` with one value per entry were a factor of 6 to 30 slow because of the merge.** Every removal called `mergeUp`, which collected the entries of the node's children into new objects before it knew whether they fit one page. Profile
   (`go test -cpuprofile`, 100,000 random keys, remove and add): 58 % in `tryMerge`, 25 % allocation. **Fixed:** the sizes are added up first without building anything (`mergeFits`), and only nodes with up to 12 children are looked at
   (`mergeChildren`, the capacity of a 12-way node: a node with more holds more entries than one page takes). The same loop: 6.2 µs per remove and add before, 0.21 µs after (a tree of single-key pages only: 0.18 µs). Memory after removing half the keys did not change (38.6 against 38.3).
2. **A key that leaves the common prefix of a page** was added by building the page again from its entries (47 % of all allocations in a build). **Fixed:** `Widen` builds the new page in one pass; the pages of a burst reuse their scratch slices
   (1,050,000 to 530,000 allocations for 400 builds of 4,096 keys; 171 to 150 ns a key).
3. **After these changes** the quick comparison on WSL (not the PC, 4,096 `street` entries, 4 processes, indicative): one value per entry `churn` 0.08 to **0.50**, `build` 0.09 to **0.55** of step 3.5; real mix `churn` 0.61 to 0.67, `build` 0.54 to 0.64. **That is still not the
   prediction (0.8 to 1.2 of `btree-map`, which means about 1.0 of step 3.5), and the PC has not measured it.**
4. **Why it is still slow: corrected on 2026-10-05 night by the probe ([step4-probe.md](step4-probe.md)).** The first guess here (transient values promote the pages) holds for `natural` only: it makes the pages decay (a tree with a history has 73 bytes per key, a fresh one 58), but under `single-value` the streams never give a key a second value, and the probe counts no promote. What costs there is the merge on every removal of a key (28 % of the time, 4 % of the tries succeed); trying it only after a removal that leaves a page nearly empty takes a third off the time of `churn` for 3.5 % more memory after a mass removal.
5. **The range scan** of a page costs 8.2 ns a value in a hot loop (single-key pages: 16.9), about half; `btree-map` is faster still on cached data. Not tuned: `scanPage` calls two closures per entry.

## Gate 4 and the options

Not met: credo 2 for ranges (0.3 to 0.7 of `btree-map`) and credo 1 for `churn` and `build` with one value per entry. The memory part is met and exceeds the prediction on the real mix.

Options (the user decides; none has been built):

- **A.** Measure the fixed code on the PC (one run), with a profile of the benchmark's own stream, before deciding anything.
- **B.** Make the promote and the burst cheaper (no allocations: build from the page in place); tune the scan of a page (one loop over the entries, no closures). Both are explainable and may bring `churn`/`build` towards 1 and the ranges towards `btree-map`.
- **C.** Go to step 5 (multi-value entries in pages), which removes the promote altogether, and keep step 4 as it is (memory first).
- **D.** Let an entry that is down to one value go back into a page (merge on `Remove` of a value, not only of a key): fewer single-key pages after churn, but more work per removal.
