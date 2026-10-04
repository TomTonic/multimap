# Step 3.4: measurement of the value overflow with `Set3`

Written 2026-10-04. What was built: [step3-overflow-design.md](step3-overflow-design.md) section 0 (option B). Words as in
[GLOSSARY.md](GLOSSARY.md). Raw results: `bench/results-layout/step3-skmv/pc/s34-*` (PC, Ryzen 9 7900, `rtcompare`, 6 to 8 processes,
sizes 4,096, 16,384 and the corpus), compared with the runs of step 3.3 (`sk-*`, same machine, same flags; commit 1cf35ec against 06244b4).

## Prediction and result

| Prediction (section 0 of the note) | Measured |
|---|---|
| Memory stays: `street` natural about 113 B/key, `dirs` about 163 | **114** and **165** (+1 and +2 B/key: +2 B a value in the overflow) |
| The containers are 15 to 40 % faster than `vset` | **not visible** in any cell of the four profiles: every ratio against `node-layout` is within 0.01 to 0.03 of step 3.3, the ns/op within about 3 % (noise) |
| Scannable bytes and GC work stay | 65 / 79 scannable B/key (3.3: 65 / 76), GC +34 / +14 ms (3.3: +33 / +14) |

**Reading.** The overflow keys are 0.5 % (`street`) and 1.6 % (`dirs`) of the keys. The operations of the benchmark (`valuesFor` of a random key,
a range over many keys, `churn`) touch them in about one case in 200 to 60, so a container that is 30 % faster moves a mean by well under
1 %: it cannot show in these cells. The 15 to 40 % of `ovbench` are real, but only an operation on the overflow keys themselves would show them;
none of the profiles has one. **The step is neutral on the existing benchmarks, and that is what it was for:** one overflow form instead of
two, any length of value, no `vset` in the string map. The gain in time belongs to a workload with many overflow keys (the inverted index of
the backlog).

## What stays open

- The set leaf of this measurement had the key areas of the old set leaf (16 to 256 bytes, odd sizes). After the measurement it was moved onto the grid of the
  pages (32, 64, 128 and 256 bytes; key area 18, 50, 114 and 242; a longer remainder as a string): the bytes that the 64-byte `vset.Set` took are
  the key area and not padding. The effect on memory is below 1 B/key (1,090 and 1,380 leaves); the final numbers come with the next measurement.
- The value of 255 bytes or more and the key of more than 242 bytes are in the tests, not in the benchmark.

## Measurement of 3.4b: the set leaf on the grid and the nine-bit remainder length (`s34b-*`, same runs as above)

Commit 71e386d against 1cf35ec (`s34-*`), PC, same flags. Raw: `bench/results-layout/step3-skmv/pc/s34b-*`.

- **Memory:** unchanged: `street` real 114 B/key, `dirs` real 165 (prediction: the same, at most 1 B).
- **Lookups and ranges (`valuesFor`, `valuesBetween`, `prefix`):** unchanged, the ratio to `node-layout` within 0.01 to 0.03 (one cell of `prefix` 86,215 on `dirs`: 0.54 to 0.58, one of
  `valuesFor` 212,449 `street` single-value: 0.85 to 0.88; the ns/op of `ordered` the same). Prediction met.
- **`build` and `churn`: 2 to 3 % slower, in all four profiles** (`build` at 4,096 keys 0.91 to 0.88 on `street` real, 0.90 to 0.87 single-value; `churn` at 4,096 and 16,384 keys 0.85 to 0.82 and
  0.86 to 0.83; at the full size 0.91 to 0.90; ns/op of `ordered` +2 to +3 % on `build`). Not predicted ("not measurable"). **Cause not investigated**; the
  hypothesis is the write paths (`cls()` and `setRem` do more operations than a byte read, and the writes call them per added value). It is the price of the nine bits and of the types in steps of two.
