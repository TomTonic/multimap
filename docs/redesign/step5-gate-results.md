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
