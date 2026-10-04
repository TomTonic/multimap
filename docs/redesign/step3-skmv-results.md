# Step 3.3: the single-key page in the tree, for string values (measured 2026-10-04)

Commit `06244b4` (a `Map[string]` holds its keys in single-key pages, the set leaf is the value overflow; the
6-byte header of the leaves; copy-out strings). Ryzen 9 7900, Windows native, `-suite dev -minprocs 6 -maxprocs 8
-memn 262144 -memrounds 3`, `-tags baseline,strvals`, sizes 4,096, 16,384 and the corpus (212,449 `street` keys, 86,215
`dirs` keys), the machine otherwise idle, the four runs took 51 minutes. `baseline` is `node-layout` (`7b8a8d8`: byte
nodes and a leaf per key, no pages). Design: [step3-skmv-design.md](step3-skmv-design.md), prediction there in section 6.
Raw results: `bench/results-layout/step3-skmv/pc/`. The speed is the new `ordered` as many times as fast as the
competitor (above 1: `ordered` faster); `*`: the interval is wider than asked, `?`: the difference does not clear the
noise. The step-2 tree is the same as `node-layout` for strings (it had no pages for them; measured in
`ref-str-street-multi`: 0.97 to 1.02 in every cell), so the comparison with `baseline` is also the comparison with before.

## Short version

1. **The natural mix: credo 1 holds.** Against `btree-sets`, `ordered` is faster in every cell: point lookups 1.35 to
   1.75, ranges 1.34 to 2.04, prefix 1.51 to 2.15, `churn` 1.16 to 1.58, `build` 1.31 to 1.69. Against `hashed` ranges are 7 to 93 times
   as fast (and lookups 0.27 to 0.47, as they must be).
2. **Against `node-layout` it is slower, in every operation:** point lookups 0.66 to 0.90, ranges 0.46 to 0.70, prefix 0.43 to 0.60, `churn`
   0.84 to 0.91, `build` 0.86 to 0.90. The prediction (design note, section 6) was "even for lookups, 1.1 to 1.5 for scans, `build`
   and `churn` within 15 %": lookups and mutation are inside it (lookups at 4K to 16K are not: 0.66 to 0.78 against
   0.9 to 1.1), **the scans are not**. The cause, concluded and not measured, is the copy of the values (one allocation for every
   key a scan visits); the user decided on 2026-10-04 to leave it and go on (PLAN.md, "To check later").
3. **One value per key (`single-value`): credo 1 and 2 are not met.** Against `btree-map`: lookups hold (1.06 to 1.43, 0.97 at 212K `street`),
   `churn` and `build` on `street` hold at 4K and 16K (1.11 to 1.24), on `dirs` they do not (churn 0.89 to 0.94, build
   0.82 to 0.92); ranges are 0.30 to 0.48 and prefix 0.34 to 0.50 of `btree-map`'s (credo 2 asks for at least 1). This is what a page
   per key can be: nothing in it is contiguous. The pages that make ranges fast, step 4, are not in yet.
4. **Memory, counted fairly** (the string bytes included for every candidate): `street` natural 113 B/key against `node-layout` 147
   (-23 %), `dirs` natural 163 against 218 (-25 %), `street` single-value 69 against 77 (-10 %), `dirs` single-value 87 against 93 (-6 %).
   As the bench counts it (the bytes of the strings a candidate does not copy are left out): 113 against 119 (-5 %), 163
   against 150 (+9 %), 69 against 67, 87 against 78.
5. **The garbage collector's work falls a lot:** scannable bytes `street` natural 65 against 111 B/key, `dirs` 76 against 139,
   `street` single-value 34 against 62, `dirs` 37 against 72; one cycle costs +33 instead of +54 ms, +14 instead of +26, +26 instead of +36,
   +11 instead of +17.
6. **Gate 3 (as written 2026-10-04: credo 1 hard, memory hard, the rest reported): not met** in the single-value profile (item 3);
   met in the natural mix. Step 3 is not done: the value overflow (3.4) and the other value types (3.5) follow, and
   credo 2 belongs to step 4. The gate is read at the end of step 3.

## Prediction against measurement (design note, section 6; the deviation is explained or marked open)

| | predicted | measured | |
|---|--:|--:|---|
| `street` natural, B/key (model: nodes 34.4, pages 44.5, overflow 38.6) | 117.5 | 113 | within 4 % |
| `dirs` natural | 157.3 | 163 | within 4 % |
| `street` single-value | 68.1 | 69 | within 2 % |
| `dirs` single-value | 83.0 | 87 | within 5 % |
| lookups against `node-layout` | even, within 10 % | 0.66 to 0.90 | **deviates** at 4K to 16K keys (-28 to -34 %); the allocation of the copy (about 20 to 25 ns, in the page microbenchmark one fifth of a lookup) fits, not measured in the tree |
| scans against `node-layout` | 1.1 to 1.5 | 0.43 to 0.70 | **deviates**: the copy per key; open, noted for later |
| `churn`, `build` | within 15 % | 0.84 to 0.91, 0.85 to 0.92 | inside |

## Results


### street natural: ordered (new) as many times as fast as ...

| operation | competitor | 4096 | 16384 | 212449 |
|---|---|--:|--:|--:|
| valuesFor | baseline | 0.66 | 0.72 | 0.85* |
| valuesFor | btree-sets | 1.68 | 1.75 | 1.57 |
| valuesFor | hashed | 0.32 | 0.30 | 0.45 |
| valuesBetween | baseline | 0.46 | 0.51 | 0.70 |
| valuesBetween | btree-sets | 1.34 | 1.38 | 2.04 |
| valuesBetween | hashed | 10.09 | 34.53 | - |
| prefix | baseline | 0.49 | 0.47 | 0.60 |
| prefix | btree-sets | 1.51 | 1.53 | 2.01 |
| prefix | hashed | 93.09 | 92.46 | - |
| churn | baseline | 0.84 | 0.85 | 0.89* |
| churn | btree-sets | 1.55 | 1.58 | 1.22 |
| churn | hashed | 0.44 | 0.44 | 0.50 |
| build | baseline | 0.89 | 0.90 | - |
| build | btree-sets | 1.63 | 1.69 | - |
| build | hashed | 0.57 | 0.54 | - |

| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-str | street | 212449 | ordered | 3 | 113 | 65 | +33 ms | 58 |
| natural-str | street | 212449 | hashed | 3 | 161 | 141 | +48 ms | 98 |
| natural-str | street | 212449 | btree-sets | 3 | 367 | 335 | +93 ms | 185 |
| natural-str | street | 212449 | baseline | 3 | 119 | 111 | +54 ms | 62 |

### dirs natural: ordered (new) as many times as fast as ...

| operation | competitor | 4096 | 16384 | 86215 |
|---|---|--:|--:|--:|
| valuesFor | baseline | 0.72 | 0.78 | 0.90* |
| valuesFor | btree-sets | 1.35 | 1.45 | 1.58 |
| valuesFor | hashed | 0.29 | 0.27 | 0.47 |
| valuesBetween | baseline | 0.47 | 0.52 | 0.66 |
| valuesBetween | btree-sets | 1.61 | 1.53 | 1.88 |
| valuesBetween | hashed | 8.64 | 27.86 | - |
| prefix | baseline | 0.43 | 0.48 | 0.54* |
| prefix | btree-sets | 2.01 | 1.71 | 2.15 |
| prefix | hashed | 7.54 | 6.63 | - |
| churn | baseline | 0.85 | 0.86 | 0.91 |
| churn | btree-sets | 1.23 | 1.33 | 1.16 |
| churn | hashed | 0.38 | 0.40 | 0.49 |
| build | baseline | 0.86 | 0.88 | - |
| build | btree-sets | 1.31 | 1.36 | - |
| build | hashed | 0.41 | 0.42 | - |

| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-str | dirs | 86215 | ordered | 3 | 163 | 76 | +14 ms | 81 |
| natural-str | dirs | 86215 | hashed | 3 | 224 | 164 | +21 ms | 130 |
| natural-str | dirs | 86215 | btree-sets | 3 | 421 | 355 | +35 ms | 209 |
| natural-str | dirs | 86215 | baseline | 3 | 150 | 139 | +26 ms | 75 |

### street single-value: ordered (new) as many times as fast as ...

| operation | competitor | 4096 | 16384 | 212449 |
|---|---|--:|--:|--:|
| valuesFor | baseline | 0.71* | 0.77 | 0.86* |
| valuesFor | btree-map | 1.36 | 1.43 | 0.97*? |
| valuesBetween | baseline | 0.50 | 0.53 | 0.65* |
| valuesBetween | btree-map | 0.30 | 0.32 | 0.38 |
| prefix | baseline | 0.55 | 0.50 | 0.57 |
| prefix | btree-map | 0.50 | 0.37 | 0.34 |
| churn | baseline | 0.95 | 0.97 | 0.96 |
| churn | btree-map | 1.11 | 1.19 | 1.01? |
| build | baseline | 0.91* | 0.92 | - |
| build | btree-map | 1.16 | 1.24 | - |

| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value-str | street | 212449 | ordered | 3 | 69 | 34 | +26 ms | 36 |
| single-value-str | street | 212449 | btree-map | 3 | 66 | 49 | +28 ms | 34 |
| single-value-str | street | 212449 | baseline | 3 | 67 | 62 | +36 ms | 35 |

### dirs single-value: ordered (new) as many times as fast as ...

| operation | competitor | 4096 | 16384 | 86215 |
|---|---|--:|--:|--:|
| valuesFor | baseline | 0.79* | 0.82* | 0.88* |
| valuesFor | btree-map | 1.06* | 1.23* | 1.15* |
| valuesBetween | baseline | 0.51 | 0.57 | 0.61 |
| valuesBetween | btree-map | 0.37 | 0.42 | 0.48* |
| prefix | baseline | 0.48 | 0.48 | 0.56 |
| prefix | btree-map | 0.39 | 0.37 | 0.39 |
| churn | baseline | 0.95 | 0.94 | 0.94 |
| churn | btree-map | 0.89 | 0.94* | 0.89 |
| build | baseline | 0.85 | 0.88 | - |
| build | btree-map | 0.82 | 0.92 | - |

| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value-str | dirs | 86215 | ordered | 3 | 87 | 37 | +11 ms | 46 |
| single-value-str | dirs | 86215 | btree-map | 3 | 107 | 49 | +12 ms | 55 |
| single-value-str | dirs | 86215 | baseline | 3 | 78 | 72 | +17 ms | 41 |

## What follows

Nothing is changed because of these numbers; the order of the plan holds (user, 2026-10-04): 3.4, the value overflow
([step3-overflow-design.md](step3-overflow-design.md), for the user's approval), then 3.5, then step 4. What the measurement puts on the list
for the profile after version 0.8: the copy-out of strings in scans and lookups (the options are in PLAN.md), the three-byte header, the
bench with scattered strings.
