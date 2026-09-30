# Pessimistic paths, leaves hold their rest (2026-09-29/30)

Every byte of a compressed path is now stored in its node: the first 12
bytes in the header as before, the rest in a tail after the node's fixed
part, in the same object (16, 48 or 112 bytes inline, a string beyond). A
lookup checks every key byte on its way down, and a leaf holds only its key
from the depth it was created at, plus the whole key's length. A leaf whose
node above it disappears on a delete is copied into one that holds its
longer rest.

The change is measured against `27d7e93` (kind class, stabilized string
values), head to head as `-vs baseline`, once with `uint64` values and once
with string values (`-tags strvals`), in two runs:

1. **`run1-v0.7/`**, the driver of `d715b81`-`27d7e93` with rtcompare
   v0.7.0, sizes up to 1M keys, with memory (1 h 50 min and 2 h 8 min,
   2026-09-29). Its intervals are pooled here with the sample standard
   deviation; the summaries in the directory show the narrower ones of the
   run (rtcompare#118).
2. **`run2-v0.8/`**, after two changes to the lookup path (below), with the
   driver of `5e7f7a5` and rtcompare v0.8.0: the dev suite, 4K, 16K and 256K
   keys, a first stage of 4 processes and at most 8, without memory
   (46 min and 55 min, 2026-09-30, machine otherwise idle):

```sh
bench -suite dev -vs baseline -sizes 4096,16384,262144 -skipmem
```

Run 2's churn and build also take out and put back long-lived values
(`-permchurn 0.25`), which reaches the rekey of a leaf whose node above it
disappears; run 1's did not. Its figures are therefore not comparable with
run 1's for churn and build. With at most 8 processes, 57 of 182 comparisons
with `uint64` values and 49 of 182 with strings are not precise (interval
wider than ±2 points and ±10%), mostly churn and lookups from 256K keys up;
they are counted with `*n` below.

## Results of run 2

Speed relative to the baseline (above 1 means faster). Each cell gives the
range over the key kinds and the median in brackets; `*n` marks n imprecise
comparisons.

| values | operation | 4K | 16K | 256K |
|---|---|---|---|---|
| multi | valuesFor | 0.96–1.00 (0.99) *2 | 0.96–1.00 (0.98) *3 | 1.00–1.04 (1.02) *4 |
| multi | valuesBetween | 0.96–1.01 (0.99) *3 | 0.97–1.00 (0.98) | 0.99–1.05 (1.02) *1 |
| multi | prefix | 0.97–1.02 (0.97) *2 | 0.98–1.02 (0.98) *2 | 1.00–1.05 (1.03) |
| multi | churn | 0.95–0.98 (0.96) | 0.95–1.00 (0.97) *1 | 0.91–1.00 (0.95) *5 |
| multi | build | 0.94–0.97 (0.95) | 0.96–0.98 (0.97) |  |
| unique | valuesFor | 0.92–1.01 (0.98) *3 | 0.93–0.99 (0.97) *1 | 0.96–1.12 (1.04) *5 |
| unique | valuesBetween | 0.93–0.99 (0.99) | 0.95–0.99 (0.98) *1 | 0.96–1.14 (1.04) *4 |
| unique | prefix | 0.95–1.01 (0.96) *1 | 0.96–0.99 (0.97) | 0.98–1.14 (1.05) *5 |
| unique | churn | 0.89–0.97 (0.93) *3 | 0.93–0.98 (0.95) *3 | 0.95–1.01 (0.98) *5 |
| unique | build | 0.87–0.94 (0.90) *3 | 0.90–0.93 (0.92) |  |
| multi-str | valuesFor | 0.95–1.02 (0.99) *2 | 0.94–1.01 (0.99) *2 | 1.00–1.07 (1.02) *2 |
| multi-str | valuesBetween | 0.95–1.00 (0.96) | 0.94–1.00 (0.98) | 0.97–1.05 (0.99) *1 |
| multi-str | prefix | 0.95–0.99 (0.96) *3 | 0.96–1.02 (0.97) *4 | 0.98–1.06 (1.00) |
| multi-str | churn | 0.92–1.01 (0.98) *1 | 0.97–1.02 (0.98) *1 | 0.93–0.99 (0.96) *6 |
| multi-str | build | 0.94–0.99 (0.97) | 0.96–1.02 (0.97) |  |
| unique-str | valuesFor | 0.92–1.02 (0.97) *1 | 0.90–1.01 (0.99) *1 | 0.99–1.11 (1.00) *4 |
| unique-str | valuesBetween | 0.92–1.00 (0.94) *4 | 0.87–1.01 (0.95) *1 | 0.98–1.02 (0.98) *1 |
| unique-str | prefix | 0.94–0.99 (0.95) *2 | 0.92–1.01 (0.96) | 0.97–1.02 (1.00) *2 |
| unique-str | churn | 0.92–0.98 (0.95) *1 | 0.93–1.01 (0.97) | 0.96–1.09 (0.97) *5 |
| unique-str | build | 0.90–0.94 (0.93) *5 | 0.92–0.97 (0.94) |  |

## Results of run 1

**`uint64` values:**

| values | operation | 4K | 16K | 256K | 1M |
|---|---|---|---|---|---|
| multi | `valuesFor` | 0.96–0.98 (0.97) | 0.96–0.99 (0.97) | 0.96–1.04 (1.01)*1 | 0.94–1.02 (1.00)*1 |
| multi | `valuesBetween` | 0.97–1.00 (0.98) | 0.98–1.00 (0.99) | 0.99–1.06 (1.01) | 0.99–1.04 (1.01) |
| multi | `prefix` | 0.95–0.99 (0.97) | 0.96–1.00 (0.98) | 1.00–1.04 (1.03)*1 | 1.00–1.05 (1.02) |
| multi | `churn` | 0.93–0.97 (0.95) | 0.92–1.01 (0.96) | 0.96–1.01 (0.98)*6 | 0.93–1.02 (0.97)*5 |
| multi | `build` | 0.94–0.98 (0.96)*1 | 0.94–0.98 (0.96) | | |
| unique | `valuesFor` | 0.92–0.98 (0.95)*1 | 0.91–0.98 (0.96) | 0.97–1.08 (1.00)*1 | 0.99–1.03 (1.00)*1 |
| unique | `valuesBetween` | 0.92–0.99 (0.98) | 0.95–1.00 (0.97) | 0.98–1.14 (1.02) | 0.98–1.05 (1.00) |
| unique | `prefix` | 0.95–0.98 (0.96) | 0.95–1.01 (0.98)*1 | 1.01–1.15 (1.03)*3 | 1.00–1.07 (1.03)*1 |
| unique | `churn` | 0.92–0.96 (0.93) | 0.91–0.97 (0.95)*3 | 0.98–1.02 (0.99)*3 | 0.99–1.00 (1.00)*5 |
| unique | `build` | 0.90–0.94 (0.92) | 0.91–0.94 (0.92)*1 | | |

**String values:**

| values | operation | 4K | 16K | 256K | 1M |
|---|---|---|---|---|---|
| multi-str | `valuesFor` | 0.95–0.99 (0.97) | 0.96–0.99 (0.99)*1 | 0.98–1.07 (1.02) | 0.98–1.05 (1.01)*1 |
| multi-str | `valuesBetween` | 0.97–1.01 (0.98) | 0.97–1.02 (0.99) | 0.97–1.06 (0.99) | 0.97–1.08 (1.00) |
| multi-str | `prefix` | 0.94–0.99 (0.96) | 0.94–1.02 (0.98)*2 | 0.96–1.05 (1.01)*2 | 0.96–1.08 (1.01)*1 |
| multi-str | `churn` | 0.92–0.99 (0.98) | 0.96–1.01 (0.99)*2 | 0.97–0.99 (0.98)*6 | 0.95–1.03 (0.97)*5 |
| multi-str | `build` | 0.94–0.98 (0.97) | 0.95–1.00 (0.97) | | |
| unique-str | `valuesFor` | 0.91–0.98 (0.96)*2 | 0.89–0.99 (0.95) | 0.98–1.10 (1.00) | 0.98–1.08 (1.01) |
| unique-str | `valuesBetween` | 0.94–1.01 (0.96) | 0.89–1.03 (0.96) | 0.97–1.02 (0.99)*1 | 0.93–1.08 (0.99) |
| unique-str | `prefix` | 0.91–0.99 (0.96)*1 | 0.91–1.01 (0.95) | 0.96–1.03 (0.99)*2 | 0.93–1.06 (0.99)*1 |
| unique-str | `churn` | 0.93–0.97 (0.94) | 0.94–1.00 (0.96)*2 | 0.98–1.05 (1.00)*2 | 0.98–1.02 (0.99)*3 |
| unique-str | `build` | 0.90–0.95 (0.93) | 0.92–0.98 (0.94) | | |

**Memory** at 1M keys (path 300K, street 212K), heap bytes per key, and the
change against the baseline:

| values | u64 | str | uuid | email | url | path | street |
|---|---:|---:|---:|---:|---:|---:|---:|
| multi | 116 (±0%) | 129 (−8%) | 147 (±0%) | 137 (−1%) | 161 (−16%) | 149 (−24%) | 82 (−4%) |
| unique | 41 (±0%) | 51 (−23%) | 68 (±0%) | 62 (−5%) | 85 (−27%) | 72 (−40%) | 59 (−2%) |
| multi-str | 250 (±0%) | 260 (−6%) | 277 (−5%) | 273 (−1%) | 294 (−11%) | 282 (−15%) | 169 (−2%) |
| unique-str | 105 (±0%) | 115 (−12%) | 132 (−11%) | 128 (−2%) | 149 (−19%) | 137 (−27%) | 123 (−3%) |

The GC time per cycle stays within the spread of the rounds.

## Findings

1. **Memory falls wherever keys are long** (run 1). Real URLs and file paths
   need 16-40% less with `uint64` values and 11-27% less with strings; the
   leaves come within 2-3 bytes per key of the smallest possible rest. uuid
   with `uint64` values keeps its leaf size class, so its memory does not
   move.
2. **Large sizes stay neutral or gain** (both runs). From 256K keys up the
   medians of lookups and ranges lie at 0.98-1.05; ranges and prefix
   searches over real paths, URLs and synthetic strings gain up to 14%.
3. **The lookup path is repaired; insertion and deletion are not.** Run 1
   lost 3-5% in lookups at 4K and 16K keys. Two changes brought it to
   1-3% (medians 0.97-0.99): the leaf compares its rest against the end of
   the key, which needs no depth and lets `matches` inline into the lookup
   again, and the check of a path beyond eight bytes moved out of the
   lookup loop. In a microbenchmark with 16K keys, lookups with `u64` keys
   went from 8% to 1% slower, with uuid keys they stay 4-8% slower, without
   a hot spot; the uuid rows are the weakest in run 2 as well (0.94-0.95 in
   several operations).
4. **`churn` and `build` remain 2-10% slower at small sizes.** With `uint64`
   values `churn` has medians 0.93-0.97 at 4K and 16K keys and `build`
   0.90-0.97; with strings 0.95-0.99 and 0.93-0.97, worst with one value per
   key (0.87-0.90 in single cells). The cause is not found. The costs sit
   where leaves and nodes are created and taken apart: a leaf is
   created with a remainder, a split needs the common prefix and a tail
   class, a delete may copy a leaf into one with a longer remainder. This
   breaks the credo's fourth point for small sizes, though no cell is worse
   than 0.87, inside the regression budget of 15%.
