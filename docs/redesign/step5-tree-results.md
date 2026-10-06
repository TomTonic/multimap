# Step 5.2: the tree with several values in a page, results (probe)

Design: [step5-mkmv-design.md](step5-mkmv-design.md) (section "5.2"). Words as in [GLOSSARY.md](GLOSSARY.md). Commit `dbbf333`. Page format: [step5-mkpage-results.md](step5-mkpage-results.md).

## What is built

- **`reach`** uses `Add`: `Added` (key +1), `AddedValue` (the key was there), `Present`, `Full` (burst; the key is counted only if it was new). **The promote and `Insert`/`Differs` are gone** (event `promote` is `value added to a key in a page in place`).
- **`pair`**: a single-key page with any number of values (not a value overflow) and a new key make a page if they fit; **`build`/`pageOf`/`pageItems`**: an item is a key with its values (`multi`), the page is sized by slots; **merge** (`mergeFits`/`tryMerge`/`leafItem`) counts slots and keys.
- **Removal**: `Remove` of both pages says `Removed`/`Gone`/`Absent`, so the tree counts the keys right; `RemoveKey` takes the values of the key one by one; `KeysUpTo(mergeBelow+1)` counts the keys left (stops after the third key, so a page of one value each costs no more than before); one key left is a single-key page.
- **Scans**: `Range` reports a key once, `RangeValues` every value; the bounds are compared once for the run.
- Tests: `mkmulti_test.go` (several values a key, burst on a further value, pairs, removals, merge; both flavors), `TestMultiKeyPages` (a second value stays in the page), the random reference tests (up to five values a key, 6 phases, scans) now run on pages with runs; `internal/art` 100 %, race, lint 0, `go vet` of `bench`.

## Probe (WSL, `MKPROBE=1 ... -tags mkstats`, 5 s for twelve cases), step 4 (`3e1e952`) → now

Raw: `bench/results-layout/step5-probe/wsl-probe.txt` (and step 4: `step4-probe/wsl-probe.txt`). B/key is "block bytes per key" of the census.

| case | fresh B/key | after one cycle B/key | after / fresh | keys in pages after the cycle, % | burst per 1000 ops (cycle) |
|---|--:|--:|--:|--:|--:|
| street natural 4,096 | 59.1 → **40.5** | 73.8 → **49.4** | 1.25 → **1.22** | 17 → **88** | 0.1 → 2.3 |
| street natural 65,536 | 58.3 → **39.2** | 73.2 → **48.5** | 1.26 → 1.24 | 16 → 88 | 0.1 → 2.5 |
| dirs natural 4,096 | 94.9 → **72.3** | 108.9 → **85.6** | 1.15 → 1.18 | 11 → 82 | 0.8 → 4.8 |
| dirs natural 65,536 | 81.2 → **58.7** | 95.3 → **71.2** | 1.17 → 1.21 | 10 → 85 | 0.6 → 3.4 |
| u64 natural 4,096 | 72.4 → 64.4 | 79.6 → 75.4 | 1.10 → 1.17 | 0 → 14 | 0 → 0.4 |
| u64 natural 65,536 | 77.2 → 62.6 | 85.5 → 73.6 | 1.11 → 1.18 | 3 → 59 | 0 → 0.1 |
| all six `single-value` cases | unchanged to the byte | unchanged to the byte | 1.02 to 1.22, as before | as before | as before |

**The six `single-value` cases are identical to step 4 in every count** (objects, bytes, events): the format and the paths for one value do not move. Their times in the probe are the noise of ten-millisecond runs (four runs each, new and old code: `dirs single-value` 4,096 replay 261 to 277 against 263 to 273 ns, 65,536 426 to 529 against 416 to 467); the figures that count are the PC and M1 runs.

## Against the prediction

| prediction | result |
|---|---|
| promote gone | **met** (0) |
| keys in pages after a cycle at least 90 % of the fresh tree's | **met**: 88.0 % of keys against 88.8 % fresh (street 4,096); 82 to 88 % for street/dirs |
| burst at most 5 in 1000 operations (real mix) | **met** for street (2.3, 2.5) and `u64`; dirs 4.8 and 3.4; **build stream** 6.0 to 11.2 (the build adds all keys of the corpus: pages fill) |
| memory real mix street 52.9, dirs 72.6 B a key (model, fresh) | the probe's census is not the heap of the benchmark (59.1 in the census against 73 measured in m43 for street): **fresh -31 % (street) and -24 % (dirs)** against -27 % and -24 % predicted from the model; the heap is measured in the gate |
| **after a cycle at most 1.05 times the fresh tree** | **missed: 1.18 to 1.24** (street 1.22, dirs 1.18) |

## The one surprise: the decay that is left is the fill of the pages

Keys are no longer thrown out of their pages (88 % stay in pages, as fresh, against 17 %), yet the tree is 22 % larger after a cycle than a fresh one with the same keys. The census by object size shows where (street 4,096, values a page of 128 / 256 / 512 bytes):

| | fresh | after a cycle | with every second key removed |
|---|--:|--:|--:|
| 128 bytes | 5.5 | 3.8 | 3.6 |
| 256 bytes | 11.6 | 8.5 | 6.7 |
| 512 bytes | 30.5 | 24.9 | 22.1 |

The pages have the same classes but hold fewer values: a page that loses entries moves to a smaller class only when its content fits **half of the smaller class** (`classFor(2*content)`: the hysteresis of steps 3 and 4, so that a page at a class border does not change its object with every insert and remove), so a page in class `c` is between 25 % and 100 % full; a page built fresh gets the smallest class that holds the content, 50 % to 100 %. The bytes a value: 23.2 → 33.8 (128), 22.2 → 29.9 (256), 16.8 → 20.6 (512). The decay of step 4 (1.25) was mostly the pages that dissolved; what is left of it is this hysteresis, and the prediction of 1.05 forgot it.

This is a knob with a reason, not a defect: the choice is between memory after use (shrink earlier) and moves of objects at class borders (time and garbage). It is **not changed in this step**.
