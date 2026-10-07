# Gate 5 on the PC: the tree with several values in the pages (5.2) and the pointer pages (5.3)

Design and prediction: [step5-mkmv-design.md](step5-mkmv-design.md). Words as in [GLOSSARY.md](GLOSSARY.md). The M1 Pro is **not** measured yet (user, 2026-10-06: after 5.5 and 5.6, so that it measures the code that ships; `bench/remote/queue.txt`).

## What was measured

The Ryzen 9 7900, `-suite dev`, 6 to 8 processes, the cases of m43 (street, dirs, `u64` keys; natural and single-value; sizes 4,096 and 16,384 and the full corpus; memory at 262,144 / 212,449 / 86,215 keys), against step 3.5 (`fe120f5`, "baseline") and `btree-sets` / `btree-map`, in three builds: values `uint64`, strings (`strvals`), pointers (`ptrvals`, new: no m43 exists). Run script `step5-gate/pc-m44/../run-m44.cmd`, 14:38 to 16:47, commit `7929687` (code of `e2c912d`). **The memory phase of `dirs natural` crashed in all three builds**: a node was left with one single-key page and no end page when a page that could not move up shrank to one key (the bug is in step 4's code too, the new structure met it). Fixed in `54e9297` with two tests that crash without the fix and a random test with the invariants checked after every operation; the three `dirs natural` cases were run again on `54e9297` (`pc-m45`, 17:00 to 17:44); the other 13 cases are of `7929687` (the fix does not touch what they measure; all 18 memory children also run clean in WSL on `54e9297`). Raw: `bench/results-layout/step5-gate/`.

## Memory (heap bytes a key, m43 → now)

| case | step 3.5 | m43 | **now** | predicted (model) |
|---|--:|--:|--:|--:|
| `uint64`, street real | 90 | 73 | **54** | 52.9 |
| `uint64`, dirs real | 108 | 96 | **74** | 72.6 |
| strings, street real | 114 | 98 | **79** | 82.8 |
| strings, dirs real | 165 | 153 | **131** | 125.6 |
| `uint64` keys real (262,144) | 125 | 123 | **102** | (random keys: not in the model) |
| pointers (`*rec`), street / dirs / `u64` keys real | 90 / 108 / 125 | | **53 / 73 / 101** | as `uint64` |
| one value a key (street, dirs, `u64`; `uint64`) | 67, 78, 53 | 27, 41, 24 | **27, 40, 24** | the same |

All within the tolerance of 10 % (the largest: strings dirs +4 %). The scannable bytes (what the collector reads) of the real mix fall with the pages' share: `uint64` street 21 → **5**, dirs 28 → **8**, `u64` keys 22 → 7, strings street 52 → 36. After removing every second key value by value: street 40 → 34, dirs 51 → 47. Maps of pointers have the memory of the `uint64` map, as predicted (their pages are scanned: 52 of 53).

## Speed: time of `Ordered` now over m43 (below 1.00 is faster; ranges over the sizes)

| | valuesFor | valuesBetween | prefix | churn | build |
|---|--|--|--|--|--|
| `uint64` real mix (street, dirs) | 0.79–1.14 | **0.68–0.94** | **0.66–0.96** | 0.85–1.13 | 0.99–1.10 |
| `uint64` keys real | 0.97–1.24 | 0.86–1.05 | | 1.00–1.15 | 1.11–1.40 |
| strings real mix | 0.87–1.14 | **0.77–1.04** | 0.82–1.18 | 0.89–1.20 | 1.04–1.11 |
| `uint64` single-value | 0.99–1.03 | 1.03–1.17 | 1.02–1.08 | 1.02–1.16 | 1.05–1.15 |
| strings single-value | 1.02–1.09 | 1.02–1.08 | 1.00–1.04 | 1.00–1.08 | 1.06–1.09 |

## Against the prediction

| prediction | result |
|---|---|
| memory of the real mix as the model | **met** (table above) |
| one value a key: memory the same | **met** |
| single-value cells within 0.95 to 1.05 of m43 | **missed slightly**: churn and build 1.00 to 1.09 (street, dirs `uint64`/strings 1.02–1.08), the scans 1.00 to 1.08, `u64` keys up to 1.16: the same code on the same pages costs 2 to 8 % more (not understood yet; candidates: the removal now asks what went and counts keys, `pageOf` loops over slots and keys, the scratch for `pair`) |
| real mix churn and build 0.95 to 1.0 of m43 | **missed**: churn 0.85 to 1.20, build 0.99 to 1.11 (not faster; the gain from no rebuild is eaten by the same 2 to 8 %) |
| real-mix ranges 1.0 to 1.3 times (faster) | **met**: `valuesBetween` and `prefix` 0.66 to 0.98 of the time (`uint64`), 0.77 to 1.04 (strings) |
| after a cycle at most 1.05 times fresh | missed in 5.2 (1.2; the fill of a randomly changed structure, no lever found) |
| burst at most 5 in 1,000 operations | met in the probe (5.2) |

## Credo, gate 5 (ordered over the competitor, churn / build; the full corpus has no `build` here)

Real mix against `btree-sets` (credo 1: at least 1.0):

| | 4,096 | 16,384 | large (churn) |
|---|--|--|--|
| `uint64` street / dirs | 1.12 / 1.13, 1.00 / 0.98 | 1.29 / 1.29, 1.22 / 1.12 | 1.40, 1.30 |
| `uint64` keys | 2.12 / 1.63 | 2.33 / 2.20 | 1.66 |
| **strings** street / dirs | **0.93 / 0.97, 0.91 / 0.89** | 1.22 / 1.18, 1.12 / 1.04 | 1.23, 1.11 |
| **pointers** street / dirs | **0.79 / 0.82, 0.73 / 0.74** | **0.95 / 0.89, 0.92 / 0.81** | 1.10, 1.03 |
| pointers `u64` keys | 1.83 / 1.19 | 1.99 / 1.86 | 1.42 |

Single value against `btree-map` (credo 1; as in m43, below 1.0 at 4,096 and 16,384): `uint64` street 0.73 / 0.84 and dirs 0.68 / 0.67 at 4,096, 1.13 and 1.00 at the large size; strings 0.65 to 0.72; **pointers 0.51 to 0.61 at 4,096 and 0.76 to 0.82 at the large size**. Credo 2 (ranges): real mix 2.0 to 4.3 times `btree-sets` (met), single value 0.28 to 0.63 of `btree-map` (as before: option B of the step-4.2 report, the scan glue).

## Findings

1. **The real mix is met for `uint64`, nearly for strings, and not for pointers.** Strings are 0.89 to 0.97 at 4,096 keys (m43: 0.99 to 1.07); pointers 0.73 to 0.89 up to 16,384 keys. The pointer page is the stage of 5.3 (every change of the number of values is a new object; the user's order: it is optimised at the end of step 5, 5.6); against step 3.5 it is 0.43 to 0.83 in churn and 0.46 to 0.62 in build, which the typed leaf did not lose.
2. **The 2 to 8 % that single-value cells lost against m43 and the missing gain for build** are the open item before 5.5: they sit in code the step touched for every page, and the baseline of all later work. A profile of `churn` and `build` of `uint64` single-value (street, 4,096) against `3e1e952` is the next measurement.
3. **The crash** (above) would have shipped: the random tests of the tree never made a page shrink to one key under a node it could not join. The new test (`TestMultiKeyPagesChurnKeepsTheStructure`) does not reproduce it either (it passes without the fix); the two tests of `TestMultiKeyPageShrunkBelowItsNode` do. Open: a random test that finds it, to be written when a way to build such a state at random is found.

## The 2 to 8 % of the single-value cells: where it comes from (analysis, 2026-10-06 evening, nothing changed in the code)

Method: the probe (`TestProbe`, the benchmark's own `churn` and `build` streams, single-value, 4,096 keys, WSL, medians of 11 interleaved runs, test binaries built from the commits) at five commits, plus the page-level benchmark `BenchmarkFixed` (hot, best of 6) and a CPU profile of 4,000 cycles at m43 and now.

| commit (what it is) | `uint64` build / replay ns an op | street | dirs |
|---|--|--|--|
| `3e1e952` (m43) | 75 / 71 | 197 / 216 | 269 / 261 |
| `70407c6` (5.1: the page format, tree unchanged) | 77 / 70 | 198 / 221 | 271 / 261 |
| `dbbf333` (5.2: the tree with several values) | 77 / 70 | 194 / 227 | 280 / 271 |
| `e2c912d` (5.3: pages for pointers) | **88 / 75** | 207 / 228 | 273 / 274 |
| `54e9297` (now) | 85 / 76 | 214 / 226 | 282 / 273 |
| now, without the `HoldsPointers` calls in `Add`/`Remove` (experiment) | 80 / 72 | 208 / 226 | 277 / 267 |

Page level, one insert and one remove (n = 3 / 7 / 20, ns): `3e1e952` 51.1 / 62.8 / 83.1; 5.1 51.7 / 63.1 / 85.6 (+1 to 3 %); **5.3 61.6 / 73.5 / 94.8 (+11 to 19 %)**; 5.3 without the `HoldsPointers` calls 53.1 / 62.7 / 86.1 (back at 5.1). Get is the same at every commit (17.2 / 23.1 / 43 ns).

**Findings.**
1. The page format of 5.1 costs nothing measurable in the tree (0 to 3 %). The tree of 5.2 costs 2 to 4 % in the replay of street and dirs (221 → 227, 261 → 271), nothing for `uint64` keys. **5.3 costs the most: +10 to 13 % in the build of `uint64` keys (77 → 88) and +3 to 7 % elsewhere.**
2. About half of that is one cause, found by removing it: **`HoldsPointers[T]()` is called in every `Add` and `Remove`** (to decide whether the page is a new object) and is reflection (`reflect.TypeFor`, `pointerFree`): about 5 ns a call, 10 ns for an insert and a remove (`internal/abi.TypeFor` and `newFixed` in the profile). Without the calls the page level is at 5.1, and the tree at 80 / 72 instead of 85 / 76 ns.
3. The rest of the 5.3 step (typed copies of the values with `copy` on `[]T`, `newFixed` in place of `grow`) is 2 to 5 ns an operation; the 2 to 4 % of 5.2 sits in the removal (`Removal`, `removeFrom`, `KeysUpTo`) and in the merge and burst paths that walk pages with the `first` flag (`mergeFits.func1` 470 → 600 ms, `pageItems` closures 200 → 320 ms of 9.5 s in the profile). Not separated further.
4. So of the 4 to 6 % (street, dirs) and 8 to 13 % (`uint64` keys, build) that single-value churn and build lost, **about half is the reflection check, which a flag decided once per map (`decide`, passed to the page functions) removes**; the other half is spread over 5.2 and the typed copies and is the price of the new code until a profile of those paths says more. The PC figures (churn and build 1.02 to 1.09, `u64` keys up to 1.16) agree.
5. The real-mix build that was predicted 0.95 to 1.0 of m43 (the 32 rebuilds in 1,000 operations are gone) is 0.99 to 1.11: its gain is smaller than these costs, the cost per operation of the same code rises by the same 4 to 6 %.

Options (not decided): (A) leave it, (B) pass `ptr bool` (decided once in `decide`) to `Add`, `Remove`, `Widen`, `Prepend`, `BuildFixedOf`; expected back to about 2 to 3 % over m43 for `uint64`; (C) B and then a look at the removal and the merge walk with the `first` flag.

### Option B done (user, 2026-10-06): the pointer flag is decided once per map

`Map.decide` sets `m.ptr` (`mkpage.HoldsPointers[T]()`, asked once); `mkpage.Fixed.Add/Remove/Widen/Prepend` and `BuildFixedOf` take it as `ptr bool` and no longer ask the type; `skpage.NewFixed/EmptyFixed/BuildFixed` ask `kindOf` once instead of twice. `kindOf` is the only use of `reflect` in the library (2.9 ns a call for `uint64`, 8 ns for a two-word array). Tests, race, lint and 100 % unchanged.

Page level (insert and remove, n = 3 / 7 / 20, ns): m43 51.1 / 62.8 / 83.1; 5.3 61.6 / 73.5 / 94.8; **now 52.4 / 63.2 / 86.9** (+1 to 4 % over m43). Tree level (probe, medians of 11, build / replay ns an operation, m43 → 5.3 → now):

| case | m43 | 5.3 | now |
|---|--|--|--|
| `uint64` single-value 4,096 | 75 / 72 | 87 / 77 | **81 / 74** |
| `uint64` single-value 65,536 | 118 / 124 | 130 / 135 | 122 / 134 |
| street single-value 4,096 | 197 / 219 | 214 / 227 | **197 / 219** |
| dirs single-value 4,096 | 264 / 262 | 279 / 272 | 274 / 267 |
| street natural 4,096 | 192 / 174 | 206 / 181 | 201 / 174 |
| `uint64` natural 4,096 | 72 / 67 | 103 / 89 | 99 / 81 |

What is left of the single-value gap is 0 to 8 %, mostly in `uint64` keys. **A second finding that this run shows and the PC run confirms** (`u64` keys real: build 1.11 to 1.40 of m43's time): the real mix on random `uint64` keys (pages of 4 entries, many bursts and merges) is 20 to 40 % slower in build and 20 % in replay than m43; it is not the reflection (it stays at 99 / 81). Cause not found yet: candidates are the items with `multi` (a slice for every key with several values in `pageItems`, `leafItem`, `appendValue`) and the larger `pageOf`.

### The `pager` interface is gone (user, 2026-10-06)

`Tree.upsert` and `Tree.splitLeaf` are methods of `Map[T]` now (`insert.go`); `reach` and `pair` are called directly, and the interface `pager` (an itab call for every `Add` that reached a page or met a single-key page) is deleted. Prediction: 2 to 3 ns an `Add` that reaches a page. Measured with a new benchmark, `BenchmarkMapAddPresent` (`Add` of a value that is there already: the descent, `upsert`, `reach`, `Page.Add` that answers `Present`; best of 8): **pages 32.1 → 30.0 ns (-2.0 ns, -6 %)**, single-key pages 24.4 → 24.9 (no change: they never called `reach`; `pair` is called only for a new key). In the probe at tree level (15 rounds, build / replay ns an operation, before → after) the change is below the noise: `uint64` single-value 77 / 70 → 77 / 72, street 211 / 221 → 201 / 219, dirs 271 / 267 → 265 / 268, `uint64` natural 100 / 84 → 99 / 77: two nanoseconds are 1 % of 200 and 3 % of 70.

## Random `uint64` keys, real mix: the build that is 20 to 40 % slower than m43 (analysis, 2026-10-06 evening, nothing changed in the code)

Cases: `u64` keys, natural values, 4,096 keys (the probe at `3e1e952`, `70407c6` = 5.1, `dbbf333` = 5.2, `e2c912d` = 5.3, `3b94663` = now; medians of 11 interleaved runs).

**Where it enters:** build / replay ns an operation `3e1e952` 73 / 67 → 5.1 72 / 68 → **5.2 99 / 87** → 5.3 102 / 86 → now 99 / 83. It is the step of 5.2 (the tree puts several values into a page), and only for few keys: at 65,536 keys the build is unchanged (184 → 180) and the replay +5 %. The PC says the same (`u64` keys real: build 1.11 to 1.40, churn 1.00 to 1.15 of m43's time; the memory is 12 to 20 % smaller: 72.4 → 64.4 B a key fresh at 4,096, 77.2 → 62.6 at 65,536).

**What changes in the tree** (events a cycle of 69,746 operations; 5.1 → 5.2): keys in pages after the cycle 0.2 % → 14 %; page changes (entries removed from a page) 118 → **4,196**, values added to a key in a page in place 0 → 3,113, keys added to a page 0 → 719; pairs of a single-key page and a new key 224 → 438 (a single-key page with several values may pair now), page with one entry left becomes a single-key page 118 → 439; bursts 0 → 31; promotes 106 → 0. So 11.5 % of the operations are now operations in a page (before: nearly none: with random keys a key with several values was a single-key page of its own).

**Cost by kind of operation** (the probe's table, ns, 3 runs, 5.1 → 5.2): insert of a new key 101 → **131**; insert into a key with 1 / 2 / 3+ values 62 / 91 / 63 → 66 / 101 / 67; delete of the last value 96 → **142**; delete from a key with 2 / 3+ values 68 / 60 → 72 / 66; all 72 → 84. The new key and the last value are 20 % of the operations and carry about 60 % of the loss (+30 and +46 ns); the others lose 5 to 11 %.

**Where the time goes** (a temporary timing of the functions, inclusive, with its own cost of about 25 ns a call; per operation of the cycle, mkstats build, nothing committed):

| function | calls a cycle | ns a call | ns an operation |
|---|--:|--:|--:|
| `reach` (an `Add` that reaches a page) | 3,863 | 185 to 204 | 10.3 to 11.3 |
| `pageRemove` | 4,196 | 212 to 247 | 12.7 to 14.9 |
| of it `removeFrom` (the page's `Remove`) | 4,196 | 102 to 108 | 6.1 to 6.5 |
| of it the conversion to a single-key page | 439 | 379 to 762 | 2.4 to 4.8 |
| `pageItems` (burst and conversion) | 499 | **608 to 1,090** | **4.3 to 7.8** |
| `pair` | 446 | 179 to 241 | 1.1 to 1.5 |
| `build` | 459 | 239 to 323 | 1.6 to 2.1 |

An allocation profile (`-test.memprofile`) agrees: 5.2 allocates 8 % more objects, 21 % of them in `pageItems` (the `items` slice of 56 bytes an entry, the key buffer, and `appendValue`'s slice for every key with several values), against 20 % fewer single-key pages.

**Findings.**
1. Most of the loss is **page operations that replace cheap single-key page operations** for keys that do not share bytes (random keys: a page holds 2 to 4 entries, 14 % of the keys): about 11.5 % of the operations now cost 100 to 140 ns more (a page `Add` or `Remove` that walks the length list and compares the remainder, cold in the cache like the single-key page it replaces) where the single-key page was one load.
2. **The conversion of a page with one entry left to a single-key page is wasteful:** it calls `pageItems`, which builds the slice of all entries, a buffer and a slice for every key with several values, to read one entry (500 to 1,100 ns a call, 439 times in 69,746 operations: 4 to 8 ns an operation, 30 to 50 % of the 12 to 17 ns that the real mix of random keys lost), and then walks the tree a second time (`findSlot`).
3. Not found (to be settled with counters per event if the user wants it): how the rest splits between `reach` (page `Add`, 10 ns an operation) and the removal in the page (`removeFrom`, 6 ns).

Options (not decided): (A) leave it (the memory of these keys is 12 to 20 % better); (B) convert without `pageItems` (read the one entry from the page, build the single-key page from it, and the slot from the descent that `Remove` has made already) and let `pageItems` make `multi` slices from one backing array: expected -5 to -8 ns an operation, i.e. back to +5 to +9 ns over m43 for this case; (C) B and a cheaper path for the page of two entries (a sorted pair compared without the length-list walk).

### Option B done: the conversion to a single-key page without `pageItems` (user, 2026-10-06)

`onlyItem` reads the one entry of a page that is about to become a single-key page (one slice for its key, one for its values if it has several), and `pageItems` makes the `multi` slices of its items from one array (`append` into a slice with room for the whole page) instead of one slice for every key with several values. Tests, race, lint and the 100 % unchanged. Prediction: -5 to -8 ns an operation for `u64` keys with the real mix at 4,096 keys (build 99 → about 92, replay 83 → about 77).

**Prediction missed.** Objects allocated in a cycle fall by 17 % (allocation profile: 114,991 → 95,737; `pageItems` and `appendValue` nearly gone), but the time falls only a little (probe, medians of 15 interleaved runs, build / replay ns an operation, before → after):

| case | before B | with B |
|---|--|--|
| `u64` natural 4,096 | 99 / 84 | 98 / 83 |
| `u64` natural 65,536 | 178 / 190 | 170 / 180 |
| street natural 4,096 | 191 / 176 | 188 / 172 |
| dirs natural 4,096 | 244 / 229 | 237 / 222 |
| `u64` single-value 4,096 | 81 / 73 | 79 / 71 |
| street single-value 4,096 | 201 / 223 | 196 / 221 |

So -1 to -4 % (-1 ns for `u64` natural at 4,096, where -5 to -8 was predicted). **The explanation of the analysis was too strong:** the timing of the functions that put 30 to 50 % of the loss into `pageItems` was inclusive and carried its own cost of about 25 ns a call, and allocations in this tree are cheap (a few ns of the cycle); what is left of the 12 to 17 ns an operation is in the page operations themselves (`reach`, `removeFrom`), which no refactoring of the conversion touches. The kept change is still right (17 % fewer allocations, 1 to 5 % less time in five of six cells, simpler conversion).


## Gate 5, second measurement `m46` (2026-10-07): the tree on the one page (5.5c), commit `69c65ea`

Same cases, same machine and script as `m44` (`run-m46.cmd`, 06:53 to 09:12, 2 h 19 min; the `dirs natural` cases run through this time, no crash; no warning other than the usual "stopped after 8 processes"), three builds (`uint64`, strings, pointers) against `baseline` (`fe120f5`) and `btree-sets` / `btree-map`. Results: `bench/results-layout/step5-gate/pc-m46/` (with `tables.txt` and the script `gate46.py` that made the tables). The column "m44" is `pc-m44`, `dirs real` from `pc-m45` (after the crash fix).

### Memory (heap bytes a key; baseline → m43 → m44 → m46)

| build | case | n | baseline | m43 | m44 | m46 | scannable m44 -> m46 | after half m44 -> m46 |
|---|---|--:|--|--|--|--|--|--|
| u64 | street-real | 212449 | 90 | 73 | 54 | **54** | 5 → 6 | 34 → 33 |
| u64 | dirs-real | 86215 | 108 | 96 | 74 | **73** | 8 → 8 | 47 → 45 |
| u64 | street-single | 212449 | 67 | 27 | 27 | **27** | 3 → 3 | 20 → 18 |
| u64 | dirs-single | 86215 | 77 | 41 | 40 | **40** | 5 → 5 | 30 → 28 |
| u64 | u64keys-real | 262144 | 125 | 123 | 102 | **102** | 7 → 7 | 61 → 59 |
| u64 | u64keys-single | 262144 | 53 | 24 | 24 | **24** | 2 → 2 | 20 → 18 |
| str | street-real | 212449 | 114 | 98 | 79 | **80** | 36 → 36 | 47 → 46 |
| str | dirs-real | 86215 | 165 | 153 | 131 | **132** | 52 → 52 | 75 → 74 |
| str | street-single | 212449 | 69 | 31 | 31 | **31** | 4 → 4 | 23 → 21 |
| str | dirs-single | 86215 | 87 | 51 | 51 | **50** | 6 → 6 | 37 → 34 |
| ptr | street-real | 212449 | 89 | – | 53 | **54** | 52 → 52 | 34 → 33 |
| ptr | dirs-real | 86215 | 108 | – | 73 | **73** | 72 → 71 | 46 → 45 |
| ptr | street-single | 212449 | 67 | – | 27 | **28** | 27 → 27 | 20 → 19 |
| ptr | dirs-single | 86215 | 77 | – | 40 | **40** | 40 → 40 | 30 → 28 |
| ptr | u64keys-real | 262144 | 125 | – | 101 | **101** | 96 → 96 | 61 → 58 |
| ptr | u64keys-single | 262144 | 53 | – | 24 | **24** | 24 → 24 | 20 → 18 |

**Memory is as in m44** (-1 to +1 B a key in every case; strings +1 on street and dirs real, `uint64` dirs real -1), as predicted (±0.3 B in the probe, a few percent more or less at 86,000 to 262,000 keys). After removing half the keys: -1 to -3 B (the pages shrink and merge a little better).

### Speed: time m46 over m44 (below 1.00 is faster; ranges over the sizes)

| build | case | valuesFor | valuesBetween | prefix | churn | build |
|---|---|--|--|--|--|--|
| u64 | street-real | 1.06–1.10 | 0.99–1.07 | 1.03–1.06 | 0.95–1.02 | 0.97–0.97 |
| u64 | dirs-real | 1.03–1.05 | 1.00–1.09 | 1.09–1.15 | 0.97–1.00 | 0.98–0.99 |
| u64 | street-single | 1.12–1.13 | 1.09–1.11 | 1.05–1.10 | 1.01–1.02 | 1.02–1.03 |
| u64 | dirs-single | 1.06–1.07 | 1.09–1.14 | 1.08–1.20 | 0.99–1.05 | 0.98–1.08 |
| u64 | u64keys-real | 1.04–1.12 | 1.00–1.19 | – | 1.03–1.06 | 1.04–1.09 |
| u64 | u64keys-single | 1.21–1.30 | 1.01–1.31 | – | 0.93–1.06 | 0.97–1.05 |
| str | street-real | 1.07–1.10 | 1.09–1.10 | 1.06–1.08 | 0.98–1.09 | 1.08–1.10 |
| str | dirs-real | 1.01–1.08 | 1.05–1.12 | 1.06–1.23 | 0.98–1.05 | 1.06–1.08 |
| str | street-single | 1.06–1.08 | 1.07–1.08 | 1.06–1.10 | 1.01–1.15 | 1.08–1.16 |
| str | dirs-single | 1.02–1.10 | 1.08–1.11 | 1.09–1.14 | 1.04–1.08 | 1.06–1.11 |
| ptr | street-real | 1.08–1.10 | 0.97–1.09 | 1.02–1.08 | 0.69–0.81 | 0.77–0.79 |
| ptr | dirs-real | 0.96–1.06 | 1.00–1.12 | 1.06–1.20 | 0.73–0.80 | 0.80–0.81 |
| ptr | street-single | 1.09–1.15 | 1.08–1.12 | 1.05–1.11 | 0.73–0.75 | 0.82–0.86 |
| ptr | dirs-single | 1.03–1.10 | 1.09–1.14 | 1.15–1.18 | 0.78–0.81 | 0.87–0.91 |
| ptr | u64keys-real | 1.05–1.12 | 1.04–1.21 | – | 0.88–0.92 | 0.81–0.94 |
| ptr | u64keys-single | 1.21–1.39 | 1.00–1.29 | – | 0.57–0.86 | 0.75–0.83 |

### Speed: time m46 over m43 (the last measurement before step 5)

| build | case | valuesFor | valuesBetween | prefix | churn | build |
|---|---|--|--|--|--|--|
| u64 | street-real | 0.84–1.25 | 0.67–0.91 | 0.70–1.18 | 0.87–0.99 | 0.96–1.00 |
| u64 | dirs-real | 0.97–1.19 | 0.75–1.02 | 0.84–1.07 | 0.88–1.10 | 1.01–1.07 |
| u64 | street-single | 1.10–1.15 | 1.13–1.17 | 1.12–1.19 | 1.03–1.06 | 1.08–1.08 |
| u64 | dirs-single | 1.05–1.07 | 1.13–1.18 | 1.13–1.24 | 1.03–1.09 | 1.04–1.15 |
| u64 | u64keys-real | 1.02–1.39 | 0.86–1.26 | – | 1.03–1.21 | 1.20–1.46 |
| u64 | u64keys-single | 1.19–1.34 | 1.15–1.39 | – | 0.97–1.22 | 1.03–1.21 |
| str | street-real | 0.94–1.25 | 0.85–1.07 | 0.87–1.26 | 0.87–1.27 | 1.13–1.22 |
| str | dirs-real | 0.96–1.17 | 0.90–1.12 | 1.03–1.13 | 0.93–1.10 | 1.12–1.17 |
| str | street-single | 1.08–1.13 | 1.11–1.15 | 1.09–1.14 | 1.09–1.23 | 1.18–1.24 |
| str | dirs-single | 1.08–1.13 | 1.11–1.14 | 1.09–1.19 | 1.06–1.12 | 1.13–1.17 |

### Credo (ordered speed over the competitor, churn / build, m44 → m46)

Real mix against `btree-sets` (credo 1: at least 1.0):

| build | keys | 4,096 | 16,384 | large |
|---|---|--|--|--|
| u64 | street | 1.12 / 1.13 → **1.18 / 1.19** | 1.29 / 1.29 → **1.37 / 1.37** | 1.40 / – → **1.40 / –** |
| u64 | dirs | 1.00 / 0.98 → **1.07 / 1.06** | 1.22 / 1.12 → **1.29 / 1.22** | 1.30 / – → **1.33 / –** |
| u64 | u64keys | 2.12 / 1.63 → **2.11 / 1.61** | 2.33 / 2.20 → **2.32 / 2.20** | 1.66 / – → **1.60 / –** |
| str | street | 0.93 / 0.97 → **0.89 / 0.96** | 1.22 / 1.18 → **1.16 / 1.16** | 1.23 / – → **1.26 / –** |
| str | dirs | 0.91 / 0.89 → **0.94 / 0.92** | 1.12 / 1.04 → **1.11 / 1.05** | 1.11 / – → **1.15 / –** |
| ptr | street | 0.79 / 0.82 → **1.16 / 1.10** | 0.95 / 0.89 → **1.32 / 1.25** | 1.10 / – → **1.39 / –** |
| ptr | dirs | 0.73 / 0.74 → **1.03 / 1.02** | 0.92 / 0.81 → **1.20 / 1.15** | 1.03 / – → **1.28 / –** |
| ptr | u64keys | 1.83 / 1.19 → **2.03 / 1.50** | 1.99 / 1.86 → **2.29 / 2.10** | 1.42 / – → **1.59 / –** |

Single value against `btree-map` (as in m43, below 1.0 at 4,096 and 16,384 for `uint64` street, dirs and strings):

| build | keys | 4,096 | 16,384 | large |
|---|---|--|--|--|
| u64 | street | 0.73 / 0.84 → **0.75 / 0.83** | 0.96 / 0.94 → **0.98 / 0.96** | 1.13 / – → **1.18 / –** |
| u64 | dirs | 0.68 / 0.67 → **0.70 / 0.70** | 0.83 / 0.77 → **0.86 / 0.80** | 1.00 / – → **1.06 / –** |
| u64 | u64keys | 1.62 / 1.57 → **1.56 / 1.52** | 1.99 / 1.54 → **2.21 / 1.60** | 1.61 / – → **1.78 / –** |
| str | street | 0.71 / 0.72 → **0.67 / 0.67** | 0.89 / 0.88 → **0.83 / 0.81** | 1.05 / – → **1.06 / –** |
| str | dirs | 0.65 / 0.60 → **0.65 / 0.59** | 0.75 / 0.70 → **0.75 / 0.69** | 0.90 / – → **0.89 / –** |
| ptr | street | 0.52 / 0.61 → **0.72 / 0.73** | 0.66 / 0.67 → **0.94 / 0.84** | 0.82 / – → **1.18 / –** |
| ptr | dirs | 0.51 / 0.54 → **0.68 / 0.65** | 0.61 / 0.60 → **0.85 / 0.75** | 0.76 / – → **0.99 / –** |
| ptr | u64keys | 0.83 / 0.93 → **1.51 / 1.27** | 1.81 / 1.17 → **2.17 / 1.46** | 1.27 / – → **1.82 / –** |

Credo 2 (ranges, ordered over `btree-sets`, ranges over the sizes): `uint64` street `valuesBetween` 2.3 to 4.5, prefix 2.0 to 4.6; dirs 2.0 to 2.9 and 2.4 to 3.1; `u64` keys 1.9 to 3.6; strings street 1.4 to 2.9 and 1.3 to 2.8, dirs 1.5 to 2.1 and 1.9 to 2.3; pointers as `uint64`.

### Against the predictions

| prediction | result |
|---|---|
| memory as the model (real mix -0.1 to -0.5 B a key against m44, single-value ±0.1) | **met** to ±1 B a key; the sign is not clean (see above) |
| **maps of pointers: churn and build -30 to -50 % against the page of 5.3** | **met for churn and build of the real mix and single-value** (0.69 to 0.91 of m44; `u64` keys 0.57 to 0.94); the credo cells of pointers moved from 0.73 to 0.89 to **1.03 to 1.39 at 4,096 and 16,384 over `btree-sets`**: credo 1 is now met for pointers in the real mix |
| single-value cells ±3 % of m44 (`uint64`: +0 to +10 % in the probe) | churn and build **met** (0.93 to 1.08); **reads missed: `valuesFor` +2 to +30 %, `valuesBetween` +0 to +31 %, `prefix` +2 to +23 %** (`u64` keys single the worst) |
| real mix ±3 % | churn and build met (0.95 to 1.09, strings build +6 to +10 %); **reads -3 to +23 %** (mostly +3 to +12 %) |
| ranges still 2 to 4 times `btree-sets` | met (`uint64`, pointers 2.0 to 4.6; strings 1.3 to 2.9) |

### Findings

1. **What 5.5 gives:** pointer maps are 10 to 40 % faster in churn and build and pass credo 1 for the real mix (1.03 to 1.39 over `btree-sets` where it was 0.73 to 0.89); memory unchanged; the structure is one page, the code `skpage` and `mkpage` is gone.
2. **What it costs: reads got slower, typically by 5 to 12 %,** in every build and case (single-value `uint64` `valuesFor` +12 %, strings +6 to +8 %, `valuesBetween` and `prefix` +5 to +20 %), not by a few cells: it is the same direction everywhere. The page-level benchmarks (hot) had Get at -1 to -11 % against `mkpage` and `skpage`, and the probe measured writes only: **reads in the tree were not measured between m44 and now**, so the prediction ("Get ±3 %, strings 0 to +5 %") is missed by the tree, not by the page. Open question (not answered here, nothing was changed): where the 5 to 12 % of a lookup in the tree go. Candidates, to be checked with a profile of `valuesFor` and of a scan: the `lay` set-up in `Get` and `Each` of the many-key form (the page level measured it hot, the tree reads cold pages), the compare of the key part in the one-key form (`matches` against the head's length), the form check, and the 4-byte head against the 3-byte one of m44 (one more byte before the key part in every many-key page, a different alignment of the lists).
3. **The credo of strings** at 4,096 keys is still below 1.0 for churn and build over `btree-sets` (0.89 to 0.96; m44 0.89 to 0.97): the string pages gained nothing in churn and lost 6 to 10 % in build.
4. **The tuning candidate "fast path for the one-key add and remove of a value"** (+13 % `uint64` single-value in the probe) shows in the gate as +1 to +3 % build, +0 to +5 % churn: smaller than the reads.


## Gate 5, third measurement `m47` (2026-10-07): the tree on the one page with the reads fixed (5.5e), commit `36f4765`

Same cases, machine and script as `m44` and `m46` (`run-m47.cmd`, 10:25 to 12:50, 2 h 25 min; no crash; no warning but the usual "stopped after 8 processes"). Results in `bench/results-layout/step5-gate/pc-m47/` with `tables.txt` and `gate47.py`. Tables as in the section of m46; "m44" is `pc-m44` (`dirs real` from `pc-m45`).

### Speed: time m47 over m46 (the effect of the fix; below 1.00 is faster)

| build | case | valuesFor | valuesBetween | prefix | churn | build |
|---|---|--|--|--|--|--|
| u64 | street-real | 0.85–0.99 | 0.93–1.03 | 0.93–0.98 | 1.01–1.04 | 1.00–1.02 |
| u64 | dirs-real | 0.89–0.95 | 0.93–0.97 | 0.93–0.94 | 1.00–1.12 | 1.01–1.08 |
| u64 | street-single | 0.84–0.94 | 0.90–0.98 | 0.89–0.95 | 1.01–1.05 | 1.02–1.05 |
| u64 | dirs-single | 0.88–0.98 | 0.87–0.95 | 0.82–0.95 | 0.99–1.03 | 1.01–1.06 |
| u64 | u64keys-real | 0.85–0.91 | 0.83–1.01 | – | 1.00–1.02 | 1.02–1.04 |
| u64 | u64keys-single | 0.73–0.85 | 0.75–0.97 | – | 1.01–1.16 | 1.03–1.11 |
| str | street-real | 0.93–0.99 | 0.97–1.00 | 0.97–0.98 | 1.01–1.14 | 1.02–1.08 |
| str | dirs-real | 0.95–1.01 | 0.97–0.99 | 0.91–0.96 | 0.99–1.12 | 1.05–1.08 |
| str | street-single | 0.92–0.98 | 0.96–0.96 | 0.95–0.98 | 1.00–1.05 | 1.03–1.05 |
| str | dirs-single | 0.95–0.96 | 0.94–0.97 | 0.90–0.93 | 1.02–1.08 | 1.03–1.08 |
| ptr | street-real | 0.85–0.92 | 0.91–0.97 | 0.92–0.95 | 1.01–1.03 | 1.02–1.02 |
| ptr | dirs-real | 0.89–0.96 | 0.91–0.97 | 0.81–0.98 | 0.98–1.10 | 0.99–1.00 |
| ptr | street-single | 0.84–0.97 | 0.90–1.01 | 0.91–0.95 | 1.01–1.06 | 1.00–1.06 |
| ptr | dirs-single | 0.89–0.95 | 0.87–0.96 | 0.83–0.90 | 1.00–1.08 | 1.02–1.02 |
| ptr | u64keys-real | 0.85–0.93 | 0.83–1.00 | – | 1.00–1.04 | 1.03–1.05 |
| ptr | u64keys-single | 0.73–0.92 | 0.76–1.13 | – | 1.00–1.07 | 1.09–1.11 |

### Speed: time m47 over m44/m45

| build | case | valuesFor | valuesBetween | prefix | churn | build |
|---|---|--|--|--|--|--|
| u64 | street-real | 0.93–1.05 | 0.99–1.02 | 0.98–1.03 | 0.96–1.05 | 0.96–0.99 |
| u64 | dirs-real | 0.93–0.99 | 0.95–1.05 | 1.01–1.07 | 0.97–1.12 | 0.99–1.07 |
| u64 | street-single | 0.94–1.06 | 1.00–1.08 | 0.98–1.03 | 1.01–1.07 | 1.05–1.08 |
| u64 | dirs-single | 0.94–1.05 | 0.99–1.04 | 0.98–1.02 | 0.98–1.08 | 0.99–1.15 |
| u64 | u64keys-real | 0.95–0.96 | 1.00–1.01 | – | 1.04–1.07 | 1.06–1.14 |
| u64 | u64keys-single | 0.90–1.11 | 0.98–1.06 | – | 0.95–1.15 | 1.07–1.09 |
| str | street-real | 1.02–1.07 | 1.06–1.10 | 1.03–1.06 | 1.01–1.24 | 1.12–1.17 |
| str | dirs-real | 1.01–1.04 | 1.04–1.08 | 1.00–1.11 | 0.97–1.18 | 1.11–1.17 |
| str | street-single | 0.99–1.04 | 1.02–1.04 | 1.01–1.07 | 1.06–1.16 | 1.14–1.19 |
| str | dirs-single | 0.96–1.05 | 1.03–1.08 | 1.01–1.06 | 1.06–1.16 | 1.09–1.20 |
| ptr | street-real | 0.92–1.00 | 0.94–1.00 | 0.95–0.99 | 0.70–0.84 | 0.79–0.80 |
| ptr | dirs-real | 0.92–0.97 | 0.98–1.01 | 0.97–1.04 | 0.74–0.85 | 0.79–0.81 |
| ptr | street-single | 0.92–1.12 | 0.99–1.14 | 0.99–1.04 | 0.74–0.80 | 0.86–0.87 |
| ptr | dirs-single | 0.92–1.04 | 0.99–1.09 | 0.98–1.03 | 0.78–0.88 | 0.89–0.93 |
| ptr | u64keys-real | 0.94–0.97 | 1.01–1.04 | – | 0.89–0.95 | 0.83–0.99 |
| ptr | u64keys-single | 0.91–1.28 | 0.97–1.21 | – | 0.56–0.90 | 0.84–0.90 |

### Speed: time m47 over m43

| build | case | valuesFor | valuesBetween | prefix | churn | build |
|---|---|--|--|--|--|--|
| u64 | street-real | 0.82–1.06 | 0.69–0.85 | 0.69–1.12 | 0.90–1.03 | 0.98–1.00 |
| u64 | dirs-real | 0.92–1.07 | 0.71–0.95 | 0.78–0.99 | 0.91–1.15 | 1.08–1.09 |
| u64 | street-single | 0.93–1.08 | 1.04–1.15 | 1.06–1.12 | 1.05–1.11 | 1.11–1.14 |
| u64 | dirs-single | 0.94–1.05 | 1.02–1.08 | 1.00–1.07 | 1.02–1.11 | 1.05–1.21 |
| u64 | u64keys-real | 0.92–1.18 | 0.86–1.05 | – | 1.04–1.21 | 1.26–1.49 |
| u64 | u64keys-single | 0.87–1.14 | 1.05–1.23 | – | 0.99–1.23 | 1.14–1.26 |
| str | street-real | 0.93–1.16 | 0.85–1.04 | 0.85–1.22 | 0.90–1.28 | 1.22–1.25 |
| str | dirs-real | 0.97–1.12 | 0.89–1.09 | 0.99–1.07 | 0.92–1.18 | 1.21–1.23 |
| str | street-single | 1.04–1.06 | 1.06–1.10 | 1.06–1.08 | 1.14–1.24 | 1.24–1.27 |
| str | dirs-single | 1.04–1.08 | 1.05–1.10 | 1.01–1.10 | 1.14–1.20 | 1.17–1.27 |

### Memory (heap bytes a key; step 3.5 baseline, m43, m44, m47)

| build | case | n | baseline | m43 | m44 | m47 | scannable m44 -> m47 | after half m44 -> m47 |
|---|---|--:|--|--|--|--|--|--|
| u64 | street-real | 212449 | 90 | 73 | 54 | **54** | 5 → 6 | 34 → 33 |
| u64 | dirs-real | 86215 | 108 | 96 | 74 | **73** | 8 → 8 | 47 → 46 |
| u64 | street-single | 212449 | 67 | 27 | 27 | **27** | 3 → 3 | 20 → 18 |
| u64 | dirs-single | 86215 | 77 | 41 | 40 | **40** | 5 → 5 | 30 → 28 |
| u64 | u64keys-real | 262144 | 125 | 123 | 102 | **102** | 7 → 7 | 61 → 59 |
| u64 | u64keys-single | 262144 | 53 | 24 | 24 | **24** | 2 → 2 | 20 → 18 |
| str | street-real | 212449 | 114 | 98 | 79 | **80** | 36 → 36 | 47 → 46 |
| str | dirs-real | 86215 | 165 | 153 | 131 | **133** | 52 → 52 | 75 → 75 |
| str | street-single | 212449 | 69 | 31 | 31 | **31** | 4 → 4 | 23 → 21 |
| str | dirs-single | 86215 | 87 | 51 | 51 | **52** | 6 → 6 | 37 → 35 |
| ptr | street-real | 212449 | 90 | – | 53 | **54** | 52 → 52 | 34 → 33 |
| ptr | dirs-real | 86215 | 108 | – | 73 | **73** | 72 → 71 | 46 → 45 |
| ptr | street-single | 212449 | 66 | – | 27 | **28** | 27 → 27 | 20 → 18 |
| ptr | dirs-single | 86215 | 78 | – | 40 | **40** | 40 → 40 | 30 → 28 |
| ptr | u64keys-real | 262144 | 125 | – | 101 | **101** | 96 → 96 | 61 → 58 |
| ptr | u64keys-single | 262144 | 53 | – | 24 | **24** | 24 → 24 | 20 → 18 |

### Credo (ordered speed over the competitor, churn / build, m44 → m47)

Real mix against `btree-sets`:

| build | keys | 4,096 | 16,384 | large |
|---|---|--|--|--|
| u64 | street | 1.12 / 1.13 → **1.18 / 1.22** | 1.29 / 1.29 → **1.31 / 1.39** | 1.40 / – → **1.41 / –** |
| u64 | dirs | 1.00 / 0.98 → **1.09 / 1.09** | 1.22 / 1.12 → **1.24 / 1.24** | 1.30 / – → **1.33 / –** |
| u64 | u64keys | 2.12 / 1.63 → **2.10 / 1.64** | 2.33 / 2.20 → **2.28 / 2.13** | 1.66 / – → **1.58 / –** |
| str | street | 0.93 / 0.97 → **0.89 / 0.97** | 1.22 / 1.18 → **1.11 / 1.11** | 1.23 / – → **1.22 / –** |
| str | dirs | 0.91 / 0.89 → **0.94 / 0.94** | 1.12 / 1.04 → **1.08 / 1.03** | 1.11 / – → **1.15 / –** |
| ptr | street | 0.79 / 0.82 → **1.17 / 1.11** | 0.95 / 0.89 → **1.30 / 1.29** | 1.10 / – → **1.35 / –** |
| ptr | dirs | 0.73 / 0.74 → **1.07 / 1.03** | 0.92 / 0.81 → **1.21 / 1.14** | 1.03 / – → **1.24 / –** |
| ptr | u64keys | 1.83 / 1.19 → **2.00 / 1.50** | 1.99 / 1.86 → **2.29 / 2.09** | 1.42 / – → **1.55 / –** |

Single value against `btree-map`:

| build | keys | 4,096 | 16,384 | large |
|---|---|--|--|--|
| u64 | street | 0.73 / 0.84 → **0.75 / 0.83** | 0.96 / 0.94 → **0.98 / 0.97** | 1.13 / – → **1.17 / –** |
| u64 | dirs | 0.68 / 0.67 → **0.71 / 0.71** | 0.83 / 0.77 → **0.88 / 0.82** | 1.00 / – → **1.04 / –** |
| u64 | u64keys | 1.62 / 1.57 → **1.55 / 1.53** | 1.99 / 1.54 → **2.15 / 1.54** | 1.61 / – → **1.80 / –** |
| str | street | 0.71 / 0.72 → **0.67 / 0.67** | 0.89 / 0.88 → **0.82 / 0.81** | 1.05 / – → **1.03 / –** |
| str | dirs | 0.65 / 0.60 → **0.64 / 0.59** | 0.75 / 0.70 → **0.75 / 0.69** | 0.90 / – → **0.90 / –** |
| ptr | street | 0.52 / 0.61 → **0.73 / 0.75** | 0.66 / 0.67 → **0.92 / 0.83** | 0.82 / – → **1.15 / –** |
| ptr | dirs | 0.51 / 0.54 → **0.68 / 0.64** | 0.61 / 0.60 → **0.82 / 0.72** | 0.76 / – → **0.99 / –** |
| ptr | u64keys | 0.83 / 0.93 → **1.54 / 1.21** | 1.81 / 1.17 → **2.00 / 1.37** | 1.27 / – → **1.74 / –** |

Credo 2 (ranges, ordered over `btree-sets`, over the sizes): `uint64` street `valuesBetween` 2.5 to 4.5, prefix 2.1 to 4.9; dirs 2.2 to 3.2 and 2.6 to 3.5; `u64` keys 2.1 to 3.6; strings street 1.5 to 3.0 and 1.4 to 3.0, dirs 1.7 to 2.2 and 2.1 to 2.5; pointers as `uint64`.

### Findings

1. **The fix of the reads holds on the PC.** Against m46: `valuesFor` 0.73 to 1.01 (typically 0.85 to 0.95), `valuesBetween` 0.75 to 1.13 (typically 0.90 to 0.97), `prefix` 0.81 to 0.98 in every build; the single-value `u64` keys cells (the worst of m46) 0.73 to 0.85. Against m44 the reads are back to where they were: `uint64` and pointers 0.91 to 1.08 (outliers: `ptr u64keys single` `valuesFor` up to 1.28 and `valuesBetween` 1.21 at one size), **strings still +1 to +11 %** (`valuesBetween` 1.02 to 1.10, `prefix` 1.00 to 1.11).
2. **The writes did not change with the fix, as predicted** (m47 over m46: churn 0.98 to 1.16, build 1.00 to 1.11; the whole range is above 1.0 for build, i.e. +2 to +5 % in the typical cell: either the difference between two runs of the PC or the larger code; there is no repetition of one commit to tell which). **Strings build is +9 to +20 % against m44** (m46: +6 to +11 %) and is the largest remaining loss of speed.
3. **Memory** as m44 and m46 (±1 to 2 B a key; strings dirs real 131 → 133, single-value strings dirs 51 → 52).
4. **Credo 1 (real mix over `btree-sets`):** `uint64` 1.09 to 1.41 at the sizes (dirs 4,096: 1.09 / 1.09, where m44 had 1.00 / 0.98), pointers 1.07 to 2.29 (m44: 0.73 to 0.89 at 4,096 and 16,384), **strings street 0.89 / 0.97 at 4,096 still below 1.0**, dirs 0.94 / 0.94; above 16,384 keys 1.03 to 1.22. Credo 2 holds (1.4 to 4.9).
5. **Single value against `btree-map`:** unchanged from m46 (`uint64` street 0.75 / 0.83 and dirs 0.71 / 0.71 at 4,096, 0.98 / 0.97 and 0.88 / 0.82 at 16,384; strings 0.59 to 0.82; pointers 0.64 to 0.92 up to 16,384); above it only for the full corpus (1.03 to 1.80, strings dirs 0.90, pointers dirs 0.99). These are the two open points of the plan ("Open after step 5").
