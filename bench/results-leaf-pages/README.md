# Leaf pages against main, head to head (2026-09-27)

Interim result of the `leaf-pages` branch, kept as the basis for deciding
how pages go on. Ordered at `6fe5dc9` (range nodes, pages only for keys with
one value) was compared with Ordered at `main 7c7e6a2`, built into the bench
as the candidate `baseline` (see `cmd/mkbaseline`), in the same process and
interleaved by rtcompare.

## Settings

- Ryzen 9 7900, WSL2, go1.27.1, one process at a time.
- Dev precision: 41 samples and 2 A/A runs per comparison.
- 4K, 16K, 64K, 256K and 1M keys for all kinds (path and street up to their
  corpus sizes); 5 to 10 processes per scenario up to uuid multi 256K, then
  5 to 6 (the run was continued with `-continue -maxprocs 6`).
- Build up to 256K for u64, str and uuid multi, otherwise up to 64K: building
  256K keys made a process take 15 minutes.
- 4M keys for u64, str and uuid with 3 processes (`4m/`); these intervals are
  wide.
- Memory at 1M keys (path 300K, street 212K), 5 rounds.

Many scenarios from 64K up did not reach the stop rule: between processes,
that is between memory layouts, results scatter by 10 to 40 points, while
each process is precise in itself. The summaries give the 95% interval
across processes for every figure.

## Speed, leaf-pages relative to main

Range of the kinds, median in brackets; above 1 is faster.

**unique** (one value per key)

| n | point | churn | build | range | prefix |
|---|---|---|---|---|---|
| 4K | 0.65–0.91 (0.85) | 0.74–0.94 (0.89) | 0.66–0.99 (0.71) | 3.0–4.2× | 0.88–5.2× |
| 16K | 0.67–1.00 (0.85) | 0.70–0.95 (0.87) | 0.60–0.79 (0.73) | 3.0–5.5× | 0.82–5.4× |
| 64K | 0.68–0.98 (0.80) | 0.93–1.13 (0.95) | 0.69–0.92 (0.82) | 2.4–6.8× | 0.98–4.5× |
| 256K | 0.75–1.56 (0.87) | 0.71–1.19 (1.02) | – | 2.4–6.6× | 0.93–5.4× |
| 1M | 0.70–0.97 (0.78) | 0.69–1.05 (0.80) | – | 2.8–4.0× | 0.91–4.7× |
| 4M | 0.59–1.03 (0.77) | 0.58–0.95 (0.86) | – | 3.1–3.8× | 4.1–5.7× |

**multi** (50% of keys with one value, the others with 2 to 200)

| n | point | churn | build | range | prefix |
|---|---|---|---|---|---|
| 4K | 0.75–0.87 (0.82) | 0.66–0.80 (0.72) | 0.50–0.69 (0.60) | 0.93–1.26 (0.97) | 0.84–1.22 (0.90) |
| 16K | 0.73–0.85 (0.82) | 0.61–0.70 (0.66) | 0.51–0.66 (0.57) | 0.89–1.21 (0.90) | 0.82–1.32 (0.86) |
| 64K | 0.52–0.75 (0.66) | 0.54–0.64 (0.59) | 0.56–0.69 (0.60) | 0.70–0.96 (0.72) | 0.60–1.16 (0.78) |
| 256K | 0.59–0.78 (0.68) | 0.56–0.68 (0.58) | 0.68–0.72 | 0.81–0.86 (0.85) | 0.71–0.89 (0.84) |
| 1M | 0.66–0.83 (0.71) | 0.55–0.64 (0.57) | – | 0.80–0.93 (0.87) | 0.84–1.13 (0.90) |
| 4M | 0.60–0.78 (0.63) | 0.59–0.66 (0.60) | – | 0.69–0.95 (0.81) | 0.71–0.90 (0.80) |

## Memory at 1M keys, heap bytes per key

| | u64 | str | uuid | email | url | path | street |
|---|---|---|---|---|---|---|---|
| unique, leaf-pages | 16 | 44 | 95 | 72 | 117 | 105 | 48 |
| unique, main | 73 | 101 | 117 | 100 | 154 | 157 | 96 |
| multi, leaf-pages | 140 | 194 | 223 | 201 | 251 | 259 | 116 |
| multi, main | 149 | 177 | 192 | 175 | 229 | 232 | 120 |

GC time per cycle drops by 50 to 99% for unique values and rises by 15 to 45%
for multi values with text keys.

## Findings

1. **No size threshold.** Neither value profile has a size from which pages
   pay off; the 1.37–1.56 for point lookups at 256K u64 and str is a cache
   effect of that one size and gone again at 1M.
2. **The value profile decides.** With unique values, pages make range
   queries 2.4 to 6.8 times as fast and need 20 to 78% less memory, at 10 to
   35% slower point lookups (uuid and email at every size) and slower builds.
   With multi values, pages lose everywhere: the leaves of keys with several
   values cut the pages between them into pieces of a key or two, and range
   nodes cost more per level than main's nodes.
3. **Criterion 3** (at most 15% slower than main in point and churn) fails for
   multi values everywhere, and for unique values with uuid, email and u64 at
   several sizes.
4. **Criterion 2** (range ≥ btree-map, memory ≤ btree-map) is not covered by
   this run: it compared with main only.
5. **Scatter between processes** grows from 64K up and is largest where the
   index is about the size of the L3 cache; see the intervals in the
   summaries.

Decision (Tom, 2026-09-27): keep main's structure for keys with several
values and use pages only where keys hold one value, switching automatically:
optimistic pages, and a subtree falls back to main's structure once keys with
several values show up in it.

## After the fallback (`0b7e400`, same day)

Measured again with the fallback to inner nodes (see `internal/art/settle.go`),
with the same settings but build only up to 64K and 5 to 6 processes
throughout; files in `fallback/`. Unique values are unchanged (same code).
Multi values, median of the kinds, before → after:

| n | point | churn | build | range | prefix |
|---|---|---|---|---|---|
| 4K | 0.82 → 0.97 | 0.72 → 0.91 | 0.60 → 0.87 | 0.97 → 1.00 | 0.90 → 0.98 |
| 16K | 0.82 → 0.95 | 0.66 → 0.87 | 0.57 → 0.88 | 0.90 → 0.99 | 0.86 → 0.99 |
| 64K | 0.66 → 0.68 | 0.59 → 0.74 | 0.60 → 0.91 | 0.72 → 0.84 | 0.78 → 0.89 |
| 256K | 0.68 → 0.78 | 0.58 → 0.82 | – | 0.85 → 1.00 | 0.84 → 1.00 |
| 1M | 0.71 → 0.82 | 0.57 → 0.81 | – | 0.87 → 1.02 | 0.90 → 1.06 |
| 4M | 0.63 → 0.78 | 0.60 → 0.80 | – | 0.81 → 0.88 | 0.80 → 1.01 |

The trees of multi maps end up exactly like main's (no pages, no range
nodes left), and memory for u64 keys is identical to main's. What remains
from 64K up comes from the leaves: main holds keys of up to 64 bytes inline
in leaves of 64 to 112 bytes, this branch only up to 16 bytes in leaves of
64 bytes, and nearly all text keys of the bench are longer (str 17-32, uuid
36, email 15-34, url 51-82, path 10-300 bytes; street 25%). Each such key
costs a separate string and a cache miss per comparison at its leaf, and
2 to 8% more memory. K pages hold key tails beyond 16 bytes out of line in
the same way, which fits the slower point lookups of uuid, email and str with
unique values.

## Leaf size classes, and a bias of the bench

Which leaves should the multi-value keys get? Two gradings were compared on
top of `0b7e400`, multi values only (files in `leaves/`): main's leaves of
64, 80, 96 and 112 bytes with keys of up to 64 bytes inline (M, baseline)
against leaves of 64 and 128 bytes with keys of up to 80 bytes inline (P).
P needs 18 to 27% more memory for text keys (str 225 instead of 177 bytes
per key) and is no faster: median 0.94-0.98 up to 16K, and from 256K no
lower than the A/A test below. M was taken.

The A/A test (files in `aa/`) compared P with an exact copy of itself,
multi values, 6 processes. Point lookups measured 0.88 (u64 256K), 0.92 (u64
1M), 0.85 (str 256K) and 0.91 (str 1M) instead of 1.0; up to 16K, identical
code measured 1.00-1.04. The build order does not explain it (with the
baseline built first, still 0.87-0.96). **Every head-to-head ratio from 64K up
in this directory is biased against Ordered by roughly 5 to 15%** until the
cause is found.

## Files

- `speed-summary.md`, `mem-summary.md`: pooled tables; `speed.jsonl`,
  `mem.jsonl`: one line per comparison and process; `run.json`: settings.
- `4m/`: the same for 4M keys.
- `fallback/`: the run after the fallback, laid out the same way.
- `leaves/`: P against M; `aa/`: the A/A test.
