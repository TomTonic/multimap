# Pages for keys with one value, on top of node-layout (2026-10-01/02)

`node-pages` (`d606ade`, then the changes listed below) against `node-layout` (`7b8a8d8`,
the baseline) and against `btree-map`. Serial regime, native Windows, Ryzen 9 7900, rtcompare
v0.8.0, 4K to 256K keys, first stage 6 processes, at most 12; memory at 1M keys (3 rounds).
Figures are candidate / baseline, above 1 the candidate is faster; cells give the range over
the key kinds and the median in brackets, `*n` marks n comparisons that are not precise.

## What is in it

Keys with exactly one value, in a map whose values are small and pointer-free (at most 8
bytes), live in **pages**: sorted arrays of up to 31 keys with their values, in one object of
64 to 512 bytes without pointers, below **range nodes** that give every page a range of key
bytes (a B-trie), so the pages stay full however many keys there are. A key that gets a second
value gets a leaf, and where such keys crowd a range node its subtree is rebuilt from inner
nodes and leaves. This is the design of the `leaf-pages` branch, ported onto the pessimistic
paths and the leaves with key remainders of `node-layout`.

Changes against the first port (`d606ade`, run `n1-*`), in the order they were made:

1. **No K pages.** `n1-u64` (K pages for keys of any length up to 255 bytes) lost 10-35% on
   point, `churn` and `build` for every key kind but integers, and needed 15-45% more memory
   than `node-layout` for uuid, email, url and path. Pages now hold integer-like keys only:
   all keys of a page have the length of the first key of the map, at most 8 bytes (`n2-u64`).
2. Keys of other lengths get leaves; a leaf created among pages triggers the fall-back check,
   so a map that does not suit pages becomes a tree of inner nodes at once.
3. A key in a page is removed in one descent (it was two).
4. The page search counts the heads below the key without branching, block by block.
5. A 64-bit bloom filter in the page head (16 B head, the four classes fill 64, 128, 256 and
   512 bytes exactly) answers most lookups of absent keys from the head alone.
6. `upsert` returns the slot only, not a 48-byte struct; a delete looks at the neighbours of a
   page only when it is thin.

## Integer keys (`u64`), one value per key, against `node-layout` (`n3-u64`)

| operation | 4K | 16K | 64K | 256K |
|---|---|---|---|---|
| `valuesFor` | 0.71 | 0.87 | 0.83 | 1.08 * |
| `valuesBetween` | 4.12 | 5.19 | 7.13 | 4.54 |
| `churn` | 1.27 | 0.97 | 1.17 | 1.56 * |
| `build` | 1.33 | 1.06 | 1.07 |  |

Memory at 1M keys: 16 bytes per key against 41 (`node-layout`); `main` needed 73.

## All other key kinds, one value per key (`n3-u64`)

These keys never go into pages (they are longer than 8 bytes, or vary in length), so this is
the cost of the page code on the paths of maps without pages.

| operation | 4K | 16K | 64K | 256K |
|---|---|---|---|---|
| `valuesFor` | 0.97–0.99 (0.97) *1 | 0.98–1.01 (0.99) *1 | 0.97–0.99 (0.98) *1 | 0.96–1.00 (0.99) *1 |
| `churn` | 0.97–1.00 (0.99) | 0.97–0.99 (0.98) | 0.97–0.99 (0.98) | 0.96–0.99 (0.98) *2 |
| `build` | 0.97–1.00 (0.99) | 0.97–0.99 (0.99) | 0.97–0.99 (0.98) |  |

Memory is identical to `node-layout` for every kind (`mem-summary.md`).

## Several values per key (`n3-u64`), all kinds

| operation | 4K | 16K | 64K | 256K |
|---|---|---|---|---|
| `valuesFor` | 0.97–1.01 (1.00) *1 | 0.99–1.01 (0.99) | 0.98–1.00 (0.99) *2 | 0.98–1.00 (0.98) *3 |
| `valuesBetween` | 0.99–1.03 (1.01) *3 | 0.99–1.01 (1.00) | 0.99–1.01 (1.00) | 0.99–1.00 (1.00) |
| `prefix` | 0.99–1.04 (1.02) *1 | 1.00–1.03 (1.00) | 0.99–1.03 (1.00) *2 | 0.99–1.01 (1.00) |
| `churn` | 0.97–0.99 (0.97) | 0.96–0.98 (0.97) | 0.96–0.98 (0.97) *2 | 0.92–0.96 (0.94) *6 |
| `build` | 0.93–1.00 (0.98) | 0.94–0.99 (0.97) | 0.95–0.99 (0.98) |  |

Memory is identical to `node-layout`.

## The probe order of `valuesFor` (`n3-lp`)

rtcompare's `valuesFor` walks the keys in one random order, over and over: a cycle as long as
the map. At 4K keys a branch predictor learns most of such a cycle, which favours the code with
fewer branches, here the leaves; the page search, which has none, does the same work every time.
`n3-lp` uses a probe order of 262144 random picks, too long to learn (a local change to the
bench's corpus, not in the repository), for `u64`:

| operation | 4K | 16K | 64K | 256K |
|---|---|---|---|---|
| `valuesFor` | 0.99 | 0.88 * | 0.84 | 1.15 * |

At 4K keys the 0.71 of `n3-u64` is 0.99 here. 16K and 64K are real: they come from the pages
costing more cache misses in a row (head, fence heads, block, values) than a leaf, and from
range nodes needing one more load per level than the 256-way nodes of the leaves tree.

## Against `btree-map` (`c7-unique-u64`), one value per key

`u64`:

| operation | 4K | 16K | 256K |
|---|---|---|---|
| `valuesFor` | 4.74 | 4.24 | 5.44 |
| `valuesBetween` | 1.80 | 1.63 | 2.44 |
| `churn` | 3.29 | 3.07 | 2.23 |
| `build` | 2.90 | 2.63 |  |

All other key kinds, which have leaves as before:

| operation | 4K | 16K | 256K |
|---|---|---|---|
| `valuesFor` | 1.44–4.20 (2.45) | 1.44–3.47 (2.24) | 1.08–1.45 (1.38) *2 |
| `valuesBetween` | 0.28–0.36 (0.29) | 0.27–0.34 (0.31) | 0.46–0.53 (0.50) *1 |
| `prefix` | 0.27–1.67 (0.79) | 0.27–1.88 (0.74) | 0.44–1.20 (0.60) |
| `churn` | 0.87–1.96 (1.32) | 0.95–2.15 (1.48) | 1.02–1.56 (1.42) |
| `build` | 0.93–1.86 (1.34) | 0.98–2.00 (1.44) |  |

Heap bytes per key at 1M keys, candidate / `btree-map`:
u64 16 / 37, str 52 / 56, uuid 68 / 77, email 61 / 57, url 86 / 101, path 72 / 101, street 59 / 47.

## Findings

1. **For integer keys the credo's unique profile holds.** `valuesBetween` is 1.6-2.4 times as
   fast as `btree-map` (it was 0.3-0.5 with leaves), points are 4-5 times, `churn` 2.2-3.3,
   `build` 2.6-2.9, memory is 16 against 37 bytes per key.
2. **Against `node-layout`, integer keys win** `valuesBetween` (4.1-7.1 times), `churn`
   (0.97-1.56) and `build` (1.06-1.33), and, from 256K keys on, lookups (1.08-1.15).
3. **Point lookups of integer keys are 12-17% slower than with leaves at 16K to 64K keys**
   (0.88 and 0.84, `n3-lp`), which breaks the credo's 15% budget for one cell. At 4K the rtcompare
   figure is a predictor effect (0.71 with its probe cycle, 0.99 without).
4. **Maps without pages pay 1-4%** for the page code on `churn` and `build` with several values
   per key and with other key kinds (0.92-1.00), and nothing on memory.
5. **Keys that are not integers are as before**: ranges 0.3-0.5 of `btree-map` and prefix
   searches 0.3-1.9, memory ahead of `btree-map` except for email (61 against 57) and street
   (59 against 47). This is the credo's remaining gap; K pages did not close it (`n1-u64`).
