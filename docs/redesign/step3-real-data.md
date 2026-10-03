# Step 3: the two real data sets, measured once through

2026-10-03. Before any layout is chosen or tuned, the tree as it stands (step 2: pages for keys with
one value, leaves for keys with several) is measured on the only data that is real from end to end,
with values of variable length where the data has them. Nothing is discarded or optimized on the
strength of it; it is the reference the candidates of step 3 are compared with.

## The data

| | `street` | `dirs` |
|---|---|---|
| keys | 425,000 German street names (OpenStreetMap), about 14 bytes | 172,431 directories of 600,000 Debian file paths, with the closing slash, about 25 bytes |
| values of a key | the localities that have a street of that name: 79% of the names one, "Hauptstr." 5,913 | the names of the files in the directory: 62% of the directories one, 89% at most four, the biggest 6,372 |
| value as a string | the locality's name, 2 to 33 bytes (10,199 localities) | the file's name, 1 to 136 bytes, median 16 (418k distinct) |
| keys for 2n of the corpus | up to 212,449 | up to 86,215 |

`dirs` is new (commit `9847a43`): the directories of the existing `path` sample, which is a random 8% of
the files of Debian 12, so its directories hold fewer files than real ones. The value counts of both are
natural; the skew of the other corpora is synthetic. Value counts of both match the synthetic skew
closely (2-4 values: 72-74% of the multi-value keys, 5-16: 20-22%, 17 and more: 6%), see
`bench/results-layout/step3-entries/`.

## What is measured

Both data sets, multi profile only, with the dev suite (4 to 8 processes at least 6, memory at
262,144 keys or the corpus maximum), against all competitors of the profile (`hashed`, `btree-sets`,
`map-sets`), once with `uint64` values (the value numbers) and once with strings (`-tags strvals`, the real
names): sizes 4,096, 16,384 and the corpus maximum.

- **PC** (Ryzen 9 7900, Windows, native): started 2026-10-03 17:11, four invocations in sequence
  (`C:\temp\bench-win\run-real.cmd`), logs `real-{u64,str}-{street,dirs}.log`.
- **M1 Pro**: jobs r1 to r4 of `bench/remote/queue.txt`, started by the user with
  `bench/remote/arm-run.sh r1` and so on.

## Results

All tables (every size, both machines): [bench/results-layout/step3-real/README.md](../../bench/results-layout/step3-real/README.md);
raw files beside it. Factors are how many times as fast `ordered` is (above 1: faster). The PC and the M1
Pro agree closely, cell for cell (within about 25%, the M1 mostly a little lower); where one says "faster",
so does the other. The PC took 1 h 02 min (`street`, `uint64`), 22 min (`dirs`, `uint64`), 1 h 16 min
(`street`, strings) and 29 min (`dirs`, strings); the M1 50, 33, 34 and 37 minutes for r1 to r4.

### Against `btree-sets`, the ordered competitor (credo 1)

Range over all sizes and both machines, `uint64` and string values alike:

| operation | `street` | `dirs` |
|---|--:|--:|
| point lookup | 1.82 - 2.60 | 1.61 - 1.92 |
| range (100 keys) | 1.54 - 2.38 | 1.58 - 2.13 |
| prefix | 1.86 - 2.50 | 1.63 - 2.68 |
| churn | 1.33 - 1.95 | 1.23 - 1.62 |
| build | 1.83 - 1.97 | 1.42 - 1.51 |

`ordered` is faster in every cell, on both machines, with both value types; the smallest advantage is
churn of `dirs` at 86,215 keys (1.23 - 1.33). `dirs` is the harder data set: more keys have several values
(38% against 21%), its keys are longer, and each value is a string of its own length.

### Against `hashed` and `map-sets` (credo 2 and what the hash maps win)

- **Ranges and prefix searches:** `ordered` is 14 - 230 times as fast as `hashed` and `map-sets`
  (`street` 20 - 230, `dirs` 14 - 68). That is the reason to have an ordered multimap.
- **Point lookups:** `hashed` is 2 - 3.5 times as fast (`ordered` 0.29 - 0.50). Against `map-sets`
  `ordered` is level on `street` (0.87 - 1.27) and mostly 13 - 37% slower on `dirs` (0.63 - 0.87; level at the largest size on the PC).
- **Churn and build:** the hash maps are 1.4 - 2.8 times as fast (`ordered` 0.36 - 0.73 for churn, 0.36 -
  0.79 for build, `dirs` the lower ones).

### Memory (same on both machines; 212,449 keys of `street`, 86,215 of `dirs`)

| bytes per key | `ordered` | `hashed` | `btree-sets` | `map-sets` |
|---|--:|--:|--:|--:|
| `street`, `uint64` | **82** | 114 | 277 | 274 |
| `dirs`, `uint64` | **98** | 170 | 327 | 331 |
| `street`, strings | **111** | 153 | 359 | 356 |
| `dirs`, strings | **141** | 216 | 413 | 417 |
| `street`, `uint64`, scannable by the GC | **35** | 73 - 81 | 62 - 80 | 58 - 79 |
| `street`, strings, scannable | 111 | 141 | 301 - 336 | 298 - 343 |
| after removing half the keys, `street` `uint64` | **38** | 70 | 136 | 150 |

`ordered` is the smallest in every row. With string values everything it holds is scannable (111 and 139
bytes a key, against 35 and 38 with `uint64`): the typed and set leaves hold pointers, and the GC cycle
takes about twice as long (`street` +55 ms against +27 ms on the PC).

### Strings against `uint64`

The speeds are the same within about 0.1 in the factors (the string headers are 16 bytes against 8 and
the comparisons go through the same leaves), but the memory is 35% (`street`) and 44% (`dirs`) higher and
three times as much of it is scannable. The bench counts the 16-byte header of each string value but not
the bytes behind it (they are views into one buffer that no candidate owns), so **a layout that stores the
bytes of the values inside its pages would pay for them in this table and the others would not**; a
comparison with such a layout has to add the bytes to every candidate.

### The unique profile: one value per key (PC)

Against `btree-map`, the reference for single-value entries, which the page layouts of step 3 replace or
keep. Full table: the README of the results directory (`pc/uni-*`). Factors as above.

- **`uint64` values:** `valuesFor` is 1.2 to 2.1 times as fast, ranges and prefixes 1.2 to 1.8 times on
  `street`. On `dirs` the ranges and `build` are the weak spot: `valuesBetween` 0.81 to 0.99, `build`
  0.74 to 0.82, `churn` 0.9 to 1.08. Memory: `street` 32 against 47 bytes a key (4 against 37 scannable),
  `dirs` 54 against 87.
- **String values: the range operations collapse.** `valuesFor` stays good (1.0 to 2.0; 1.05 at the
  largest `street`, within the noise), but `valuesBetween` is **0.27 to 0.39** and `prefix` **0.24 to 0.68**
  of `btree-map`'s speed, on both data sets and at every size, with intervals that clear the noise by far.
  With several values per key (multi profile) the same operations were 1.6 to 2.1 times *faster* than
  `btree-sets`. Memory: 59 against 58 bytes a key (`street`), 69 against 99 (`dirs`), and all of it
  scannable (62 and 72), as `btree-map` has 49.
- **Why, in short (not yet traced by a profile):** in a map with string values every key is a typed leaf
  of its own; a scan of single-value entries goes leaf by leaf, each a separate object with a string
  header that points to the bytes, which is a random cache miss per entry. `btree-map` keeps header and
  key in one node array. This is exactly the cost that a page with inline values removes, so it is the
  strongest argument so far that a page for string values (layout B with variable-length values) belongs
  into the tree; the scalar pages already show it for `uint64` on `street` (ranges 1.2 to 1.8).

## Caveats

- **The M1 was not at rest.** The load average before the jobs r1 to r4 was 7.9, 3.8, 6.2 and 5.9 (see
  `env.txt`). The intervals are narrow and agree with the PC, but the M1 figures are less trustworthy than
  the PC's. r4 (`dirs`, strings) agrees with the PC cell by cell, a little lower, as the others do.
- **Precision.** At 4,096 keys, `build` against `map-sets` was not as precise as asked on the PC (the
  suite says so; the cells are marked `*`); a few other cells of `valuesFor` against `map-sets` are
  marked. None of them changes a conclusion above.
- **`dirs` comes from a sample.** The 600,000 paths are a random 8% of the files of Debian 12, so its
  directories hold fewer files than real ones; the counts are natural for the sample, not for Debian.
- **Not measured:** the unique profile on the M1, 1M keys (the corpora are smaller), and the
  effect of the layouts of step 3. This is the reference for them: commit `9847a43`.
