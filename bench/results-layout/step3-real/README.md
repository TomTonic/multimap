# The two real data sets, measured through (2026-10-03)

The tree as of commit `9847a43` (step 2: pages for keys with one value, leaves for keys with several)
against the competitors of the multi profile, on `street` (425,000 street names with the localities
that have such a street) and `dirs` (172,431 directories with the names of the files in them), both with
their natural values, once with `uint64` values (the value numbers, `u64`) and once with strings
(`-tags strvals`, the real names, `str`). Report: [docs/redesign/step3-real-data.md](../../../docs/redesign/step3-real-data.md).

- `pc/<values>-<kind>/`: Ryzen 9 7900, Windows, native; `-suite dev -values multi -minprocs 6 -maxprocs 12
  -memn 262144 -memrounds 3`, sizes 4,096, 16,384 and the corpus maximum.
- `m1/r1..r3/`: Apple M1 Pro, 16 GB, jobs of `bench/remote/queue.txt` (`-minprocs 8 -maxprocs 24`); r4
  (`dirs`, strings) had not run when this was written. The load average before r1-r3 was 3.8 to 7.9,
  so the machine was not at rest (see each `env.txt`).

Every cell: how many times as fast `ordered` is as the competitor (above 1 it is faster), from the pooled
rtcompare comparison of the processes. `*` marks an interval wider than asked for, `?` a difference
that does not clear the noise floor; a `-` is a comparison the suite leaves out (ranges and prefixes against
the hash maps, and `build`, at the largest size).


## street, uint64 values

| operation | against | PC 4096 | PC 16384 | PC 212449 | M1 4096 | M1 16384 | M1 212449 |
|---|---|--:|--:|--:|--:|--:|--:|
| valuesFor | btree-sets | 2.60 | 2.49 | 1.90 | 2.35 | 2.30 | 2.22 |
| valuesFor | hashed | 0.43 | 0.39 | 0.49 | 0.45 | 0.37 | 0.41 |
| valuesFor | map-sets | 1.24 | 1.08 | 1.23* | 1.04 | 0.89* | 0.89 |
| valuesBetween | btree-sets | 2.02 | 1.85 | 2.38 | 1.59 | 1.54 | 2.00 |
| valuesBetween | hashed | 23.01 | 68.50 | - | 21.00 | 75.30 | - |
| valuesBetween | map-sets | 23.36 | 67.74 | - | 20.47 | 69.74 | - |
| prefix | btree-sets | 2.35 | 2.50 | 2.44 | 2.00 | 1.86 | 2.13 |
| prefix | hashed | 197.10 | 199.98 | - | 206.47 | 221.07 | - |
| prefix | map-sets | 186.12 | 188.87 | - | 191.15 | 205.86 | - |
| churn | btree-sets | 1.95 | 1.91 | 1.36 | 1.89 | 1.81 | 1.51 |
| churn | hashed | 0.53 | 0.49 | 0.57 | 0.52 | 0.46 | 0.43 |
| churn | map-sets | 0.67 | 0.63 | 0.69 | 0.58 | 0.55 | 0.53 |
| build | btree-sets | 1.95 | 1.83 | - | 1.97 | 1.83 | - |
| build | hashed | 0.63 | 0.54 | - | 0.63 | 0.50 | - |
| build | map-sets | 0.79* | 0.68 | - | 0.73 | 0.58 | - |

## dirs, uint64 values

| operation | against | PC 4096 | PC 16384 | PC 86215 | M1 4096 | M1 16384 | M1 86215 |
|---|---|--:|--:|--:|--:|--:|--:|
| valuesFor | btree-sets | 1.91 | 1.91 | 1.92 | 1.61 | 1.63 | 1.70 |
| valuesFor | hashed | 0.36 | 0.32 | 0.42 | 0.33 | 0.29 | 0.36 |
| valuesFor | map-sets | 0.87 | 0.75 | 1.04 | 0.68 | 0.63 | 0.69 |
| valuesBetween | btree-sets | 2.04 | 1.79 | 2.13 | 1.62 | 1.58 | 1.82 |
| valuesBetween | hashed | 20.63 | 64.79 | - | 19.99 | 68.32 | - |
| valuesBetween | map-sets | 20.85 | 60.46 | - | 19.61 | 61.44 | - |
| prefix | btree-sets | 2.68 | 2.16 | 2.10 | 1.89 | 1.63 | 2.17 |
| prefix | hashed | 19.97 | 16.14 | - | 17.07 | 14.34 | - |
| prefix | map-sets | 20.72 | 15.99 | - | 17.18 | 14.44 | - |
| churn | btree-sets | 1.52 | 1.62 | 1.33 | 1.40 | 1.37 | 1.23 |
| churn | hashed | 0.43 | 0.42 | 0.52 | 0.36 | 0.37 | 0.36 |
| churn | map-sets | 0.49 | 0.51 | 0.65 | 0.39 | 0.44 | 0.45 |
| build | btree-sets | 1.51 | 1.49 | - | 1.44 | 1.42 | - |
| build | hashed | 0.44 | 0.42 | - | 0.40 | 0.36 | - |
| build | map-sets | 0.54 | 0.55 | - | 0.45 | 0.43 | - |

## street, string values

| operation | against | PC 4096 | PC 16384 | PC 212449 | M1 4096 | M1 16384 | M1 212449 |
|---|---|--:|--:|--:|--:|--:|--:|
| valuesFor | btree-sets | 2.58 | 2.46 | 1.82 | 2.37 | 2.37 | 2.15 |
| valuesFor | hashed | 0.44 | 0.39 | 0.50 | 0.45 | 0.38 | 0.41 |
| valuesFor | map-sets | 1.27* | 1.07* | 1.19* | 1.05 | 0.92 | 0.87 |
| valuesBetween | btree-sets | 2.03 | 1.85 | 2.25 | 1.65 | 1.65 | 2.01 |
| valuesBetween | hashed | 22.78 | 69.97 | - | 21.90 | 77.99 | - |
| valuesBetween | map-sets | 23.00 | 64.47 | - | 21.03 | 72.15 | - |
| prefix | btree-sets | 2.35 | 2.47 | 2.32 | 2.06 | 1.96 | 2.16 |
| prefix | hashed | 195.89 | 195.38 | - | 215.18 | 228.81 | - |
| prefix | map-sets | 185.05 | 184.14 | - | 195.77 | 211.28 | - |
| churn | btree-sets | 1.81 | 1.82 | 1.33 | 1.84 | 1.71 | 1.38 |
| churn | hashed | 0.52 | 0.50 | 0.58 | 0.52 | 0.48 | 0.44 |
| churn | map-sets | 0.65 | 0.66 | 0.73 | 0.62 | 0.62 | 0.59 |
| build | btree-sets | 1.83 | 1.83 | - | 1.88 | 1.83 | - |
| build | hashed | 0.63 | 0.58 | - | 0.62 | 0.53 | - |
| build | map-sets | 0.77* | 0.73* | - | 0.76 | 0.67 | - |

## dirs, string values

| operation | against | PC 4096 | PC 16384 | PC 86215 |
|---|---|--:|--:|--:|
| valuesFor | btree-sets | 1.89 | 1.87 | 1.83 |
| valuesFor | hashed | 0.38 | 0.33 | 0.46 |
| valuesFor | map-sets | 0.87 | 0.75 | 1.00*? |
| valuesBetween | btree-sets | 2.02 | 1.75 | 2.06 |
| valuesBetween | hashed | 19.84 | 57.81 | - |
| valuesBetween | map-sets | 20.05 | 55.41 | - |
| prefix | btree-sets | 2.65 | 2.10 | 2.26 |
| prefix | hashed | 18.71 | 14.88 | - |
| prefix | map-sets | 19.86 | 14.77 | - |
| churn | btree-sets | 1.42 | 1.49 | 1.24 |
| churn | hashed | 0.44 | 0.45 | 0.53 |
| churn | map-sets | 0.50 | 0.56 | 0.65 |
| build | btree-sets | 1.48 | 1.48 | - |
| build | hashed | 0.47 | 0.47 | - |
| build | map-sets | 0.58 | 0.59 | - |

## Memory (identical on both machines, except the GC figures)

| data | candidate | heap B/key | scannable B/key (PC / M1) | GC CPU per cycle (PC / M1) | heap B/key after removing half |
|---|---|--:|--:|--:|--:|
| street uint64 | ordered | 82 | 35 / 35 | +27 ms / +21 ms | 38 |
| street uint64 | hashed | 114 | 81 / 73 | +38 ms / +29 ms | 70 |
| street uint64 | btree-sets | 277 | 80 / 62 | +50 ms / +39 ms | 136 |
| street uint64 | map-sets | 274 | 79 / 58 | +43 ms / +29 ms | 150 |
| dirs uint64 | ordered | 98 | 38 / 37 | +12 ms / +8 ms | 46 |
| dirs uint64 | hashed | 170 | 88 / 81 | +11 ms / +12 ms | 100 |
| dirs uint64 | btree-sets | 327 | 82 / 62 | +20 ms / +15 ms | 159 |
| dirs uint64 | map-sets | 331 | 88 / 66 | +16 ms / +12 ms | 180 |
| street string | ordered | 111 | 111 / 108 | +55 ms / +42 ms | 53 |
| street string | hashed | 153 | 141 / 141 | +49 ms / +36 ms | 90 |
| street string | btree-sets | 359 | 336 / 301 | +90 ms / +67 ms | 177 |
| street string | map-sets | 356 | 343 / 298 | +65 ms / +42 ms | 192 |
| dirs string | ordered | 141 | 139 / - | +27 ms / - | 67 |
| dirs string | hashed | 216 | 164 / - | +21 ms / - | 123 |
| dirs string | btree-sets | 413 | 353 / - | +31 ms / - | 202 |
| dirs string | map-sets | 417 | 362 / - | +28 ms / - | 223 |
