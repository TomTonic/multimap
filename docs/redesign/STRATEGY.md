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
   cell.
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
| R1 sizes | 128, 256 or 512 bytes. Pointer-free objects may also be 1024 or 2048. |
| R2 64 bytes | Only for anomalous inner nodes, such as chains of path bytes in file paths. Never for pages, leaves or the objects of small maps. |
| R3 alignment | Use Go size classes that are multiples of 128. Up to 512 bytes they are aligned to their size. Objects with pointers above 512 bytes get an 8-byte malloc header and lose alignment, so they are allowed only for the 256-way nodes (N256, R256): rare, at the top of the tree, and always hot. |
| R4 no object per key | A key never gets an object of its own. Keys live, many to an object, in pages. The only per-key objects are the value sets of keys with many values (R6). |
| R5 lines per lookup | Inside an object, a lookup decides in the first 128 bytes and reads its payload in at most one more line of the same object. |
| R6 values | A key's few values sit next to it in the page. Many values go to a value object built from 128- or 256-byte blocks. |

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
line 0 (128 B): head 16 B  kind, class, count, shared-prefix length, heap top, stale, bloom (8 B)
                heads      first 8 suffix bytes of each key, big endian (fences + blocks as now)
lines 1..:      more heads, then per key: suffix length, heap offset, value count
                values     one word per value
                heap       suffix bytes beyond the first 8, growing down from the end
```

What follows from this layout:

- **A key costs its suffix bytes plus about 11 bytes.**
- **Lookup:** search the heads (line 0, maybe line 1), then compare the tail and read the value.
  Both are usually in one more line of the same object.
- **A key that ends where the page starts** (today a node's term leaf) is an entry with an empty
  suffix.
- **The rule "the first key fixes the page key length" disappears.**
- **Sizes and splits:** classes of 128, 256 and 512 bytes, pointer-free up to 1024 if long
  suffixes need it. A page splits by count or by bytes, whichever runs out first.

### 4.2 Several values per key

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
| Point lookups in pages at 16K-64K: 0.84-0.88 of leaves. The cause is serial misses (head, fences, block, values) and range nodes costing one load more per level than 256-way nodes. R5 is the answer: decide in line 0, payload in one more line. A blocked layout (4 keys and 4 values per row) gained only 1.5 ns at 64K in a prototype. | `node-pages` `n3-lp` |
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
