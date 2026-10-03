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

Results will be added here and in `bench/results-layout/step3-real/`.
