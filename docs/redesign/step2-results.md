# Step 2: pages for every key in the tree (results)

Code state: branch `cacheline`, the commit that carries this file. The design is in
[step2-design.md](step2-design.md). Raw files, with every cell's interval, are in
`bench/results-layout/step2-pages/`. Serial regime, native Windows, Ryzen 9 7900, rtcompare v0.8.0, dev
suite sizes 4K, 16K and 256K (`build` up to 64K), first stage 6 processes, at most 12; memory at 256K keys
(3 rounds). Figures are new / reference, above 1 the new tree is faster, `*` marks a cell whose
interval is not as narrow as asked. The reference of the step is `node-pages` (`cf6944b`): it was the
last state, and `node-layout` (`7b8a8d8`, plan reference for credo 3) has the same leaves, so a cell
below 0.85 here is one below 0.85 there for the string kinds.

## Result in one paragraph

Every key with one value now lives in a page, whatever its length. Against `node-pages` ranges and prefix
queries of string keys are 2.5-5.6 times faster, memory per key is 4-46% lower for five of seven kinds, the
garbage collector has 2-9 times less to scan, the multi-value profile is neutral (0.95-1.05), and at 256K
keys the point lookups of string keys are 8-29% faster. **But at 4K and 16K keys point lookups are 5-19%
slower, `churn` of email and uuid 10-20% slower, `build` 3-30% slower, and `u64` is 5-14% slower in
every operation.** Gate 2 is not met; see the reading below.

## Speed: unique profile (one value per key)

### unique, new / node-pages (>1: new faster)

| kind | n | valuesFor | valuesBetween | prefix | churn | build |
|---|---:|--:|--:|--:|--:|--:|
| u64 | 4096 | 0.87 | 0.95 |  | 0.89 | 0.86 |
| u64 | 16384 | 0.92 | 0.95 |  | 0.91 | 0.93 |
| u64 | 262144 | 0.91 * | 0.94 |  | 0.92 * |  |
| str | 4096 | 0.91 | 4.37 | 5.60 | 1.06 | 0.89 |
| str | 16384 | 0.86 | 3.66 | 5.60 | 0.96 | 0.85 |
| str | 262144 | 1.29 | 3.83 | 5.47 | 1.15 |  |
| uuid | 4096 | 0.88 * | 3.28 | 0.94 | 0.90 | 0.76 |
| uuid | 16384 | 0.95 | 3.65 | 1.06 | 0.89 | 0.78 |
| uuid | 262144 | 1.16 * | 2.99 | 1.59 | 0.99 |  |
| email | 4096 | 0.83 | 3.68 | 0.96 | 0.81 | 0.72 |
| email | 16384 | 0.81 | 3.63 | 0.91 | 0.80 | 0.72 |
| email | 262144 | 1.13 | 2.79 | 1.24 | 0.87 * |  |
| url | 4096 | 0.83 | 2.59 | 0.79 | 0.90 | 0.70 |
| url | 16384 | 0.90 | 2.54 | 0.94 | 0.98 | 0.72 |
| url | 262144 | 1.08 | 2.43 | 1.70 | 1.02 * |  |
| path | 4096 | 0.83 | 2.57 | 0.91 | 1.03 | 0.77 |
| path | 16384 | 0.90 | 2.68 | 1.31 | 1.02 | 0.79 |
| path | 262144 | 1.10 | 2.57 | 2.90 | 1.08 |  |
| street | 4096 | 0.97 | 4.40 | 1.93 | 1.21 * | 0.94 |
| street | 16384 | 1.08 | 4.64 | 3.52 | 1.25 | 0.97 |

| operation | min | median | max | cells below 0.85 |
|---|--:|--:|--:|--:|
| valuesFor | 0.81 | 0.91 | 1.29 | 4 of 20 |
| valuesBetween | 0.94 | 2.89 | 4.64 | 0 of 20 |
| prefix | 0.79 | 1.31 | 5.60 | 1 of 17 |
| churn | 0.80 | 0.97 | 1.25 | 2 of 20 |
| build | 0.70 | 0.79 | 0.97 | 8 of 14 |

### unique, new / btree-map

| kind | n | valuesFor | valuesBetween | prefix | churn | build |
|---|---:|--:|--:|--:|--:|--:|
| u64 | 4096 | 4.27 | 1.67 |  | 2.96 | 2.52 |
| u64 | 16384 | 3.95 | 1.51 |  | 2.80 | 2.43 |
| u64 | 262144 | 3.72 | 2.25 |  | 1.96 |  |
| str | 4096 | 2.26 | 1.16 | 1.45 | 1.46 | 1.19 |
| str | 16384 | 2.28 | 1.10 | 1.44 | 1.61 | 1.29 |
| str | 262144 | 1.73 | 1.83 | 2.16 | 1.49 |  |
| uuid | 4096 | 2.51 | 0.96 | 1.36 | 1.47 | 1.26 |
| uuid | 16384 | 3.04 | 1.23 | 1.79 | 1.73 | 1.45 |
| uuid | 262144 | 1.66 | 1.63 * | 1.92 | 1.50 |  |
| email | 4096 | 2.93 | 1.25 | 1.66 | 1.64 | 1.36 |
| email | 16384 | 2.85 | 1.22 | 1.72 | 1.79 | 1.50 |
| email | 262144 | 1.46 | 1.27 * | 1.44 | 1.37 |  |
| url | 4096 | 1.45 | 0.69 | 0.68 | 0.98 | 0.76 |
| url | 16384 | 1.58 | 0.76 | 0.84 | 1.09 | 0.85 |
| url | 262144 | 1.23 | 1.10 * | 1.07 * | 1.14 |  |
| path | 4096 | 1.15 | 0.70 | 0.56 | 0.90 | 0.70 |
| path | 16384 | 1.28 | 0.74 | 0.72 | 0.96 | 0.78 |
| path | 262144 | 1.18 * | 1.25 | 1.26 * | 1.12 |  |
| street | 4096 | 1.88 | 1.22 | 1.27 | 1.42 | 1.20 |
| street | 16384 | 2.01 | 1.22 | 1.50 | 1.55 | 1.29 |

| operation | min | median | max | cells below 0.85 |
|---|--:|--:|--:|--:|
| valuesFor | 1.15 | 1.94 | 4.27 | 0 of 20 |
| valuesBetween | 0.69 | 1.22 | 2.25 | 4 of 20 |
| prefix | 0.56 | 1.44 | 2.16 | 4 of 17 |
| churn | 0.90 | 1.48 | 2.96 | 0 of 20 |
| build | 0.70 | 1.27 | 2.52 | 3 of 14 |

### multi, new / node-pages

| kind | n | valuesFor | valuesBetween | prefix | churn | build |
|---|---:|--:|--:|--:|--:|--:|
| u64 | 4096 | 1.00 | 0.99 |  | 1.03 | 0.99 |
| u64 | 16384 | 1.00 | 1.00 |  | 1.02 | 1.00 |
| str | 4096 | 1.00 | 1.00 | 1.03 * | 1.05 | 0.97 |
| str | 16384 | 1.00 | 1.00 | 1.01 | 1.03 | 1.00 |
| uuid | 4096 | 1.00 | 0.99 * | 1.00 | 1.03 | 0.96 |
| uuid | 16384 | 0.99 | 1.01 | 1.01 | 1.02 | 0.99 |
| email | 4096 | 1.00 * | 0.99 | 1.00 | 1.03 | 0.96 |
| email | 16384 | 1.00 | 1.00 | 1.01 | 1.02 | 0.97 |
| url | 4096 | 1.01 | 1.00 | 1.00 | 1.03 | 0.97 |
| url | 16384 | 1.01 | 1.00 | 1.00 | 1.02 | 0.98 |
| path | 4096 | 1.01 | 1.02 | 0.99 | 1.02 | 0.99 |
| path | 16384 | 1.00 | 1.00 | 0.99 | 1.02 | 1.01 |
| street | 4096 | 0.99 | 1.03 | 1.00 | 1.03 | 1.01 |
| street | 16384 | 1.00 | 1.01 | 1.00 | 1.03 | 0.95 |

| operation | min | median | max | cells below 0.85 |
|---|--:|--:|--:|--:|
| valuesFor | 0.99 | 1.00 | 1.01 | 0 of 14 |
| valuesBetween | 0.99 | 1.00 | 1.03 | 0 of 14 |
| prefix | 0.99 | 1.00 | 1.03 | 0 of 12 |
| churn | 1.02 | 1.03 | 1.05 | 0 of 14 |
| build | 0.95 | 0.98 | 1.01 | 0 of 14 |

## Memory at 256K keys, unique profile (heap bytes per key, 3 rounds)

| kind | new | node-pages | btree-map | scannable by the GC: new / node-pages / btree-map | after removing half the keys: new / node-pages / btree-map |
|---|--:|--:|--:|---|---|
| u64 | 18 | 16 | 37 | 1 / 1 / 37 | 10 / 10 / 19 |
| str | 31 | 51 | 56 | 3 / 27 / 36 | 21 / 23 / 25 |
| uuid | 65 | 68 | 77 | 4 / 27 / 37 | 44 / 32 / 35 |
| email | 56 | 61 | 58 | 7 / 25 / 37 | 32 / 27 / 26 |
| url | 91 | 90 | 101 | 15 / 36 / 43 | 60 / 42 / 47 |
| path | 67 | 73 | 101 | 11 / 38 / 37 | 47 / 34 / 47 |
| street | 32 | 59 | 47 | 4 / 35 / 37 | 20 / 26 / 20 |

(`mem-summary.md` in `memory-vs-*`; the GC CPU per cycle falls from 17-37 ms to 1-9 ms for the string kinds.)
The pages are below `btree-map` for every kind, below `node-pages` for five, and 2 bytes (`u64`, the
directory) and 1 byte (`url`) above it. After removing every second key they keep 20-60% more per key than
the leaves do (uuid, email, url, path): a page that loses keys at random keeps a third of its room until its
neighbour is thin too (`MergeFill` 75 after this run: 60 gave 46, 37, 65 and 51).

## Objects (`objstat`, unique, 4K to 256K)

Objects at multiples of 128 bytes: 100% for `u64`, `str`, `uuid` and `email`, 97-99% for `url`, `path` and
`street`; the rest are range nodes with a path tail (144 or 176 bytes). Line overflow at most 0.5%. The
pages are 128, 256 and 512 bytes, without pointers. Gate 2 asks for 95%: met.

## Reading against gate 2

| criterion | result |
|---|---|
| credo 1 and 2, every kind with one value: ranges at least `btree-map`, memory at most `btree-map` | memory met for all seven kinds. Ranges: met at 256K for every kind (1.10-2.25) and at 4K-16K for `u64`, `str`, `email`, `street`, `uuid`; **not met for `url` and `path` at 4K-16K (0.69-0.76)** and `uuid` 4K (0.96). `churn`, `build`, `prefix` of `url`/`path` at 4K-16K are below `btree-map` too (0.56-0.98). Points are 1.15-4.3 times `btree-map` everywhere |
| no cell below 0.85 against the reference (credo 3) | **not met**: 4 of 20 `valuesFor`, 2 of 20 `churn`, 8 of 14 `build` cells, 1 `prefix` cell |
| `u64` not worse than `node-pages` beyond noise | **not met**: 0.86-0.95 in `valuesFor`, `valuesBetween`, `churn`, `build` at 4K-256K |
| objects at multiples of 128: 95% | met |
| multi profile neutral | met (0.95-1.05) |

## Why: what the pages cost, and what was done about it

- **A lookup touches the directory line, the slot line and the tail line of a 512-byte object**; a
  leaf is one line. Two rounds against one. Resident in L2 (4K and 16K keys: 100 KB to 1 MB) that costs
  a few ns per lookup; at 256K, where a leaf is a cache miss, the page's shorter routing wins.
- **A page insert is a memmove of the slots, the directory and the tail heap, and a page grows through
  three classes and splits;** a leaf insert allocates 32-64 bytes. The first measurement of `build`
  against `node-pages` was 0.57-0.70, `churn` 0.65 for `u64`; it is 0.70-0.97 and 0.80-1.25 now.
  What got it there (all in step2-design.md): the head word taken from the end of the key for short
  suffixes (`u64` lookups 0.83 to 0.91), `burst` instead of rebuilding a subtree from its keys (40%
  of the allocation of a string build), splitting in place, compacting the heap in place, block copies
  on growth, one slot (head and value) instead of two arrays, a branch-free tag scan for integers.
- **The tree has one range node per two or three pages** (`R8` is 21-29% of the objects of the string
  kinds): a router of 128 bytes for two or three children, nested wherever keys share bytes. For `url`
  and `path`, whose pages hold 5-7 keys, this costs 9-10 bytes a key, depth, and the ranges and `build`
  of those two kinds against `btree-map` at 4K-16K. A node of 64 bytes for up to four ranges (R2 allows
  it for the anomalous chains of path keys) is the step 4 remedy.
- **`u64`: the directory does not pay.** One byte a key of memory (18 against 16 bytes), a second array to
  maintain on every insert and delete, and the tag scan on every lookup, for 2 rounds instead of 3 that the
  old page's bloom filter and fence search needed only when the page is out of cache; at 256K it is
  still 0.91. A uniform page that keeps the bloom filter in place of the directory would be the old page.

## What speaks against the change

The credo's third rule fails at the sizes people use most (4K-16K), for `build` and `u64` in particular,
and the page that was designed for strings costs the integer keys that `node-pages` served well. The gains
(ranges, memory, GC, 256K) are real and large. Whether they are worth the loss is the user's call.

## Options

1. **Keep the pages for every kind, accept the small-map loss as the price,** and run steps 3 to 6 (several
   values per key, the routing layer with 64-byte routers, values with pointers). Restate credo 3 for
   this redesign: at most 15% against the reference at 256K, at most 25% at 4K-16K.
2. **Keep the page of `node-pages` for integer keys** (uniform page with bloom filter and fences, no
   directory) next to the new general page: `u64` back to 1.0 and 16 bytes a key, one more page kind
   in the code.
3. **Size switch:** leaves below some number of keys per map, pages above it. Leaves are best where
   everything is in cache; the user parked this idea earlier. It doubles the code for strings.
4. **Stop** and keep `node-pages`.

My recommendation is 2 and 1 together: the integer page gets its old speed back at little cost, the
rest of the loss (5-19% for point lookups at 4K-16K, 3-30% for `build`) is what the design costs, and
step 4 (64-byte routers) should take another 5-10% off lookups, `build` and `url`/`path` ranges.
