# String values (2026-09-28/29)

The first measurement of values that hold a pointer. Ordered keeps them in
set leaves: the key's leaf holds a `vset.Set` with up to three values
inline, then a separate array, then a hash set. Flat leaves apply only to
small values without pointers, and every earlier run used `uint64` values.

This is `node-layout` at `d715b81` (16-byte node header, flat leaves, leaf
class in the kind) in a bench built with `-tags strvals`, against all
candidates, with the flags of main's night run:

```sh
go build -tags strvals -o bench-str ./cmd/bench
bench-str -suite dev -sizes 4096,16384,262144,1048576 -maxprocs 6 \
  -memn 1048576 -memrounds 3
```

Values are 16 hex digits in one shared buffer (see the bench README). The
run took 5 h 7 min and did not sleep. 346 of 400 comparisons were precise
within 6 processes. The other 54, mostly churn and point lookups from 256K
keys up, were measured again on 2026-09-29, each scenario with only those
comparisons and up to 30 processes, and the 7 still open with up to 80 more
(`logs/stab-*`). Their rows pool all processes of these runs and replace the
first ones. 399 comparisons are now precise; the last, churn against
`hashed` with str keys at 1M, spreads 14 points between its 30 processes.

## Results

Speed of `ordered` against each candidate (above 1 means `ordered` is
faster). Each cell gives the range over the key kinds and the median in
brackets; `*` marks cells with imprecise comparisons.

**Several values per key (`multi-str`):**

| vs | operation | 4K | 16K | 256K | 1M |
|---|---|---|---|---|---|
| `hashed` | `valuesFor` | 0.39–0.99 (0.67) | 0.37–0.94 (0.63) | 0.42–1.05 (0.66) | 0.35–1.00 (0.54) |
| `hashed` | `valuesBetween` | 13.0–19.3 (14.9) | 38.3–59.4 (46.9) | | |
| `hashed` | `churn` | 0.46–1.01 (0.69) | 0.48–1.00 (0.69) | 0.58–0.96 (0.84) | 0.51–0.94 (0.78)* |
| `hashed` | `build` | 0.50–1.01 (0.74) | 0.49–0.95 (0.73) | | |
| `btree-sets` | `valuesFor` | 1.78–4.03 (2.75) | 1.76–4.03 (2.93) | 1.51–3.30 (2.10) | 1.60–4.13 (2.44) |
| `btree-sets` | `valuesBetween` | 1.65–2.52 (1.95) | 1.58–2.14 (2.03) | 1.89–2.73 (2.37) | 1.99–3.33 (2.45) |
| `btree-sets` | `prefix` | 1.33–2.25 (2.02) | 1.44–2.38 (2.10) | 1.84–2.59 (2.13) | 2.05–2.65 (2.50) |
| `btree-sets` | `churn` | 1.35–2.82 (1.99) | 1.37–2.80 (2.03) | 1.22–1.84 (1.69) | 1.39–2.59 (2.04) |
| `btree-sets` | `build` | 1.36–2.41 (1.95) | 1.34–2.56 (1.98) | | |
| `map-sets` | `valuesFor` | 0.89–2.35 (1.49) | 0.78–2.17 (1.43) | 0.81–2.34 (1.36) | 0.70–2.22 (1.13) |
| `map-sets` | `churn` | 0.52–1.06 (0.74) | 0.55–1.13 (0.82) | 0.60–0.99 (0.89) | 0.54–0.97 (0.81) |
| `map-sets` | `build` | 0.52–0.99 (0.81) | 0.55–1.12 (0.86) | | |

**One value per key (`unique-str`), against `btree-map`:**

| operation | 4K | 16K | 256K | 1M |
|---|---|---|---|---|
| `valuesFor` | 1.38–5.35 (2.60) | 1.44–4.81 (2.75) | 0.95–2.03 (1.22) | 1.24–3.18 (1.90) |
| `valuesBetween` | 0.24–0.37 (0.25) | 0.24–0.32 (0.28) | 0.28–0.46 (0.41) | 0.44–0.68 (0.53) |
| `prefix` | 0.24–1.78 (0.76) | 0.24–1.98 (0.74) | 0.29–1.16 (0.60) | 0.39–1.67 (0.77) |
| `churn` | 0.93–2.52 (1.53) | 0.93–3.03 (1.76) | 0.93–1.64 (1.46) | 1.14–2.26 (1.81) |
| `build` | 1.00–2.11 (1.49) | 1.01–2.38 (1.69) | | |

Range queries on `hashed` and `map-sets` scan every key; like `prefix`, they
far favour `ordered` (6× to 700×) and are left out above.

**Memory** at 1M keys (path 300K, street 212K), heap bytes per key:

| values | u64 | str | uuid | email | url | path | street |
|---|---:|---:|---:|---:|---:|---:|---:|
| multi-str `ordered` | 250 | 276 | 293 | 275 | 329 | 332 | 173 |
| multi-str `hashed` | 262 | 282 | 302 | 283 | 325 | 319 | 153 |
| multi-str `btree-sets` | 452 | 472 | 492 | 473 | 514 | 515 | 359 |
| unique-str `ordered` | 105 | 131 | 148 | 130 | 184 | 188 | 127 |
| unique-str `btree-map` | 48 | 68 | 88 | 69 | 111 | 112 | 58 |

The GC time per cycle of `ordered` is 1.3–1.7 times that of `hashed`
(multi) and 2.2–3.6 times that of `btree-map` (unique): almost all of its
heap is scannable.

## Findings

1. **The band holds for several values per key.** `ordered` is faster than
   `btree-sets` in every operation (1.2× to 4.1×), and range queries are far
   faster than on `hashed`. It needs about the memory of `hashed` and half
   that of the B-tree of sets.
2. **One value per key misses criterion 2 clearly.** `ordered` needs about
   twice the memory of `btree-map` (105–188 against 48–112 bytes per key),
   and range queries reach only 0.24–0.68 of its speed. A set leaf with one
   string takes 96 bytes or more: 4 bytes head, the key, and a 64-byte
   `vset.Set` with three inline slots, a count, a capacity and a pointer, of
   which one slot is used.
3. **Against `hashed`, point operations lose more than with integers**
   (median 0.54–0.84). With integer values, a key's values are in its flat
   leaf; with strings, every key with more than three values reaches them
   through a second object.
4. This is the case typed leaves are for: a leaf that holds its key and a
   small array of values of their own type, which the garbage collector
   scans, without the set's overhead and its second object for four to
   sixteen values.
