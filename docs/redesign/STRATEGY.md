# Strategy: an ART that uses whole cache lines

Status: agreed direction (2026-10-02), not yet implemented. The plan that turns it into code is
[PLAN.md](PLAN.md); how to measure is [MEASURING.md](MEASURING.md).

## 1. Goal

`multimap.Ordered` is the general-purpose multimap with point, prefix and range queries: the
structure you take without thinking, between a hash map (fastest points) and a B-tree (fastest
ranges), and never far behind either. The credo, in force for every change:

1. **Band.** In every scenario `Ordered` beats the B-tree (`btree-sets`, `btree-map`) in point
   queries, `churn` and `build`, and beats `hashed` in range queries.
2. **Unique profile.** With one value per key, ranges at least as fast as `btree-map`, memory
   not larger.
3. **Regression.** At most 15% slower than the reference (`main`, or the step before) in any
   cell. *Loosened for step 3 (user, 2026-10-04):* the user had the impression that tuning against
   thresholds had caught the work in a local optimum. In step 3 the regression against the
   reference is measured and reported, not gating; a cell below 0.70 must be explained with its
   cause. Credo 1 and the memory stay hard (PLAN.md, gate 3).
4. **Robustness.** All key kinds (u64, str, uuid, email, url, path, street), all sizes (4K, 16K,
   256K, 1M; 16K and 256K are unlucky fill sizes), both value profiles. Small maps must never
   get noticeably worse: they are far more common than large ones.

## 2. The premise

Modern CPUs read memory in whole cache lines: 64 bytes on x86-64, 128 bytes on Apple's arm64
cores, and x86's spatial prefetcher fetches 64-byte lines in aligned pairs. Once a line is read,
its other bytes are free. The structure must turn that into speed: **every byte of a line a
lookup reads should be useful to that lookup**, and the structure must still be frugal with
memory, which pays off in turn once data leaves the caches.

Rules that follow, binding for every object the tree allocates:

| rule | detail |
|---|---|
| R1 sizes | 128, 256 or 512 bytes; pages also 384 (the user's sketch, whataleafneedstostore.md; a Go size class, aligned to 128). Pointer-free objects may also be 1024 or 2048. Exception (2026-10-03): an *oversized object*, a single-key page whose remainder does not fit 512 bytes, keeps the remainder inline and takes the size Go gives it; it saves the cache miss of a pointer to the key. |
| R2 64 bytes | Only for anomalous inner nodes, such as chains of path bytes in file paths. Never for pages, leaves or the objects of small maps. |
| R3 alignment | Use Go size classes that are multiples of 128. Up to 512 bytes they are aligned to their size (384 to 128). Objects with pointers above 512 bytes get an 8-byte malloc header and lose alignment, so they are allowed only for the 256-way nodes (N256, R256): rare, at the top of the tree, and always hot. |
| R4 (struck 2026-10-03) | It meant "no separately allocated string object on the Go heap for a key". That follows from keeping long remainders inline (the oversized object), so it needs no rule of its own. The number is kept so that the other rules keep theirs. |
| R5 rounds per lookup | A lookup inside an object takes at most two rounds of cache-line loads (defined below the table). Round 1 reads only the first 128 bytes of the object (if possible its first 64 bytes), which hold everything that says where the key's data lies. Round 2 reads all of that data at once. |
| R6 values | A key's few values sit next to it in the page. Many values go to a value object built from 128- or 256-byte blocks. |

**What a round is.** A round is a set of cache-line loads whose addresses are all known before
the first of them completes, so the CPU issues them together and they cost one memory latency, not
one each. The next round starts when its address depends on data of the round before (a pointer, an
offset, a position found by a search). When the object is not in the cache a round costs one miss
(about 50-100 ns from RAM); when it is, a few cycles. So R5 does not limit how many lines a lookup
touches, only how many times it must wait for memory: a page lookup reads the directory in the
first lines (round 1), finds the position, and then loads head word, value and tail in round 2, in
whichever lines they lie. Step 1 measured that rounds, not lines, set the cost of a cold lookup
(loading all lines at once changed nothing; cutting rounds from 3-4 to 2 cut the cold lookup by 35-45%).

The object statistic (`objstat`, see PLAN step 0) is the measure of these rules. Target: 100% of
objects at multiples of 64 bytes, at least 95% at multiples of 128, no object crossing more lines
than its size needs.

## 3. Where we stand

Branch `node-pages` (on `node-layout`) holds the measured state:

- **Inner nodes** N5, N12, N26, N58: 64, 128, 256, 512 bytes, pessimistic paths with up to 12
  bytes in the 16-byte header.
- **Longer paths** go into a tail in the same object. This gives 80-, 144-, 176- and 624-byte
  nodes.
- **Leaves, one per key:**
  - flat leaves for small pointer-free values (32 to 512 bytes);
  - typed leaves for small values with pointers (sizes depend on key and value capacity);
  - set leaves otherwise.
- **Pages** (64 to 512 bytes, branch-free block search, bloom filter) below **range nodes**
  (B-trie, 128 to 512 bytes). They hold only keys of one length of at most 8 bytes with exactly
  one small pointer-free value.

The object statistic of that state ([objstat-node-pages.md](objstat-node-pages.md)):

- Leaves are 65-92% of all objects.
- 30-92% of the objects miss "multiple of 64" and 84-98% miss "multiple of 128". The exceptions
  are u64 with one value per key (pages): 0%.
- Up to 37% of the objects cross more lines than needed (uuid and email with one value per key).

The leaf per key is the cause. A lookup ends with a miss on a small object, and the rest of that
line belongs to unrelated keys. Resizing leaves only moves the waste around.

The pages prove the alternative. For u64 keys with one value each:

| measure | pages | comparison |
|---|---|---|
| object-statistic violations | 0% | |
| memory per key | 16 bytes | `btree-map` 37 bytes |
| ranges | 1.6-2.4× | `btree-map` (leaves: 0.3-0.5×) |
| point queries | 4-5× | `btree-map` |
| `churn` | 2.2-3.3× | `btree-map` |
| `build` | 2.6-2.9× | `btree-map` |

The strategy is to make pages the rule for every key, not the exception for integer keys.

## 4. Target architecture

### 4.1 Pages for every key (replace all leaf kinds)

A page holds the keys of one subtree in key order:

- each key's suffix below the page's position;
- the suffix bytes all its keys share, stored once;
- the values.

It never holds a pointer per key and never a separate copy of a key. Starting point for the
layout of a pointer-free page (the prototype in PLAN step 1 settles it):

```
header 8 B      kind, class, count, capacity, uniform length, prefix length, heap top
prefix          the bytes all keys of the page share, padded to 8 (only if it saves 16 B or more)
directory       one entry per key, in key order: a tag byte (uniform page) or tag, suffix
                length and tail offset, 4 B (general page); in the first line or two
heads           the first 8 suffix bytes of each key, big endian, zero padded; sorted
values          one word per key
heap            the suffix bytes beyond the first 8, growing down from the end
```

(The layout of the prototype, `internal/vpage`; see docs/redesign/step1-results.md.)

What follows from this layout:

- **A key costs its suffix bytes plus about 12 bytes** (17 in all for an integer key), less what the
  prefix saves.
- **Lookup:** compare the key's tag with all tags of the directory (first line or two), then read
  head, value and tail, which are independent, in one round: two rounds in all (R5).
- **A key that ends where the page starts** (today a node's term leaf) is an entry with an empty
  suffix.
- **The rule "the first key fixes the page key length" disappears.**
- **Sizes and splits:** classes of 128, 256 and 512 bytes, pointer-free up to 1024 if long
  suffixes need it. A page splits by count or by bytes, whichever runs out first.

### 4.2 Several values per key

**Decision of 2026-10-03 (see PLAN step 3 and GLOSSARY.md).** First step: a *single-key page* holds one
entry with all its values inline, and replaces the leaf kinds. No page mixes entries with different
numbers of values yet. The plan below, with values inline next to the key in a multi-key page, is the
step after: it is what removes the per-key object and the fall back, and is decided by what the
multi profile shows.

**Order of 2026-10-04 (user):** the single-key page first, with `string -> {string}` as the base case
and `string -> {uint64}`, `string -> {*T}` and `uint64 -> {*T}` after it; then the multi-key page with
single-value entries (MKSV) as the special case; multi-value entries in multi-key pages (MKMV, the
plan below) last, because it makes the code much more complex. An experiment of MKMV exists (branch
`mkmv-experiment`, [step3-tree-pages.md](step3-tree-pages.md)).

- **Up to about 4 values stay inline**, next to the key's entry. That covers 85% of the keys of
  the multi profile.
- **More values go to a value object** of 128- or 256-byte blocks. It is a B-tree of value blocks
  or the existing `vset`, whichever measures better.
- **The pointer to such an object** lives in a page kind of its own, with a small pointer array.
  Pages without such keys stay pointer-free, and the garbage collector never scans them.

### 4.3 Values with pointers (e.g. strings)

- The page is allocated with its real generic type, as typed leaves are today, so the garbage
  collector sees the pointers.
- Its capacity is chosen so that the size lands on 256 or 512 bytes. Never above 512 (R3).

### 4.4 The routing layer

- **Range nodes** (128, 256, 512; R256 as the exception) directly above the pages, as now.
- **Inner nodes** where keys share bytes: N12 (128), N26 (256), N58 (512), N256.
- **N5 (64 bytes) goes.** With 10-30 keys per page there are 10-30 times fewer nodes, so a
  128-byte minimum costs little. It is allowed only under R2.
- **No path tails.**
  - Up to 12 path bytes stay in the node header.
  - Longer shared prefixes go into the page's shared-prefix field.
  - Long chains above inner nodes get a path node of 128 bytes: up to about 100 path bytes and
    one child. 64 bytes is acceptable here under R2.

## 5. Lessons that bind the new design

| lesson | source |
|---|---|
| K pages (keys up to 255 bytes, 16-byte heads, a pointer to a separate full-key copy per long key, capacities 1/3/7/14) lost 10-35% on point, `churn` and `build` and needed 15-45% more memory than leaves. Never a pointer or object per key; capacity by bytes, not by tiny fixed slot counts. | `node-pages` `n1-u64` |
| Pages for keys with several values (S, U8-n) won at 256K (point 1.04-1.25, ranges 1.2-1.7, memory -20..-51%) but lost at 4K (`churn` 0.57-0.81, `build` 0.38-0.71). Small maps are the risk; measure 4K and 16K first, every time. | `leaf-pages`, night run 2026-09-28 |
| Point lookups in pages at 16K-64K: 0.84-0.88 of leaves. The cause is serial misses (head, fences, block, values) and range nodes costing one load more per level than 256-way nodes. R5 is the answer: decide in the first line(s), then one round for the payload. A blocked layout (4 keys and 4 values per row) gained only 1.5 ns at 64K in a prototype. | `node-pages` `n3-lp` |
| 63-key pages were worse at small sizes, 15-key pages too; 31 stayed. Binary search, full-count search and an inlined search lost against the branch-free block search. | `node-pages` diagnosis |
| 256-byte pages measured worse than 512-byte ones (an extra tree level). | `leaf-pages` |
| A tree that turns out unsuitable for pages must fall back at once (settle, crowded ratio 4). Keep that robustness, even if the page becomes the rule. | `node-pages` |
| rtcompare's `valuesFor` walks the keys as one cycle of length n. At 4K the branch predictor learns it, which favours branchy code (pages 0.71 with the cycle, 0.99 with a long random order). | `node-pages` `n3-lp` |
| rtcompare before v0.7.0 favoured the last-validated candidate by 5-22% from 64K up. Compare only with v0.8.0 or later, interleaved, never pooled across regimes. | memory `project_bench_bias` |
| Multi `churn` at 1M u64 is 0.82 against `main` since `node-layout`, cause not found (spread, no hotspot). Step 3 must look at it again. | `node-layout` |

## 6. What to reuse

From `node-pages`:

- page search (fences, `bits.Sub64` block count, pad words);
- bloom filter in the head;
- range nodes (`rnode.go`);
- `settle`/`crowded`/`inner` fallback;
- `removeRaw` in one descent;
- slot-only `upsert` with `t.at`;
- the test harness (`checkNode`, `checkPage`, reference model, fuzzing);
- the bench with `-vs baseline` and `mkbaseline`.

The flat, typed and set leaves remain the fallback until pages cover their cases, and are then
removed.
