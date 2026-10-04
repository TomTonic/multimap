# Step 3.4: design note of the value overflow for string values

Status: **decided 2026-10-04 (user): no new hash set.** Step 3.4 is option **B** of section 3a: the set leaf of a string map
holds a `Set3` on its own, `vset` is gone from it. The block set (C) and the arena (D) of sections 2 to 6 are **not built in this
step**; they stay in this note as a candidate for the profile after version 0.8, next to the dictionary of values (H). Section 0
says what B is; sections 1 and 3a still hold. Written 2026-10-04. Words as in
[GLOSSARY.md](GLOSSARY.md). Numbers from `bench/cmd/skmodel -overflow` (a model on the real keys and values), the
measurement of step 3.3 and `internal/vset`.

## 0. What is built: the set leaf of a string map holds a `Set3`

- **One overflow form.** The set leaf of a map of strings (`Map.flat == 3`) holds, behind the leaf head and the key area, **a pointer to a
  `Set3[string]`** (github.com/TomTonic/Set3, a swiss-table hash set that exists, is tested and is the user's) instead of the 64-byte `vset.Set[string]`
  with its inline stage and array stage. It takes any string, so the second form of the block design (values of 255 bytes or more)
  is not needed: **one** form for every key that does not fit a page, and for keys of more than 254 bytes.
- **The leaf is on the grid** (user, 2026-10-04): 32, 64, 128 or 256 bytes, the head 6 bytes, the pointer the last 8, the key area what is
  left (18, 50, 114, 242 bytes; a longer remainder: the whole key as a string, 32 bytes + the string). The bytes the 64-byte `vset.Set` took are
  not saved but go to the key. 384 and 512 are of no use as long as `klen` is one byte (the remainder is at most 254); see 3.4b in PLAN.md.
  The pointer stays last: moving it first would make the key offset of this leaf differ from every other leaf (`matches` is the hot path) for about
  1 ns a leaf and GC cycle (measured: 1 M objects of 256 bytes, pointer first 7.1 ms, last 7.9 ms).
- **Created** with room for what the page held (`EmptyWithCapacity` of 1.5 times its values, so the spill does not rehash).
- **Transitions** as in the table of section 2, with the block-set leaf left out: page -> set leaf when the content no longer fits
  512 bytes or a value of 255 bytes or more arrives; set leaf -> page when the content is 256 bytes or less and every value is shorter than 255
  bytes (the check of step 3.3, `unspillSK`).
- **`vset` stays** for `Hashed` and for the maps of other value types; nothing in it changes (option F).
- **Prediction** (from the measurement of section 3a): the containers are 15 to 40 % faster than `vset` (a hit 55 against 65 ns, a miss 17
  against 31, add + remove 38 against 63) and a leaf is smaller; the heap stays where it is, **+2 B a value** (50.4 against 48.6 on `street`), so the
  memory of the natural mix does not fall: `street` about 113 B/key, `dirs` about 163, as measured in 3.3. What this step does *not* bring: the third of the
  heap that is value sets, and the pointer to every string.
- **How it is checked:** the tests of 3.3 for the transitions run on the new leaf (the reference test with values that overflow, the fuzz test,
  the rekey test), the object statistic knows the new leaf, `go test -race`, the coverage of `internal/art` stays 100 %; measurement as
  in 3.3 against this prediction.

## 1. The problem, with the data

A key whose content does not fit the largest page (512 bytes: header, remainder, the values with their length
bytes) keeps its values in a value set (today the set leaf with a `vset.Set[string]`: 3 string headers inline, an array of
headers up to 64, then a hash set). On the real data these are few keys with most of the values:

| data | keys with overflow | values they hold | value bytes | most values of a key |
|---|--:|--:|--:|--:|
| `street` natural | about 1,090 (0.51 % of the keys) | about 212,000 (37 % of all) | 2.2 MB (10.3 B a value) | 4,177 |
| `dirs` natural | about 1,380 (1.6 %) | about 112,000 (38 %) | 2.9 MB (25.9 B a value) | 3,064 |

Their values by key: `street` 337 keys have up to 64 values, 573 up to 256, 142 up to 1,024, 34 more; `dirs` 954
keys up to 64, 370 up to 256, 52 more. In the model of step 3.1 these value sets are about **a third of the
heap**: 38.6 of 117 B/key on `street`, 46 of 157 on `dirs`, counted as 24 to 40 bytes of headers and hash
a value, *without* the bytes of the strings, which a set of string headers leaves with the caller.

Two things are wrong with that for the aim of the redesign: it is the part of the structure that is not
cache-line shaped (a pointer to every string, a scattered hash set), and the garbage collector has to scan every
one of the headers.

## 2. The design: a hash set of 512-byte blocks

The values of an overflow key go into **pointer-free blocks of 512 bytes**, which the set finds by the hash of the
value (extendible hashing). Nothing in a block points anywhere.

```
set        : the number of values, the bytes of all values, the global depth d, and
             a directory of 2^d pointers to blocks
block (512): local depth (1 B) | count (2 B) | used (2 B) | values: (length, bytes) ... | free
```

- **Where a value lives:** in the block `dir[hash(v) >> (64-d)]` (the top d bits of a 64-bit hash). Several
  directory slots point to one block when its local depth is below d.
- **Contains / Add:** hash the value, take the block, walk its `(length, bytes)` records comparing length first
  (a block holds about 22 values of 10 bytes, 9 of 25), done. Add appends at `used` if the value fits.
- **A full block splits:** a new block, the records of the old one are divided by the next hash bit, the
  directory entries are set (the directory doubles if the local depth equals d). One split costs about as much as
  45 hashes and copies, and happens once per about 20 added values.
- **Remove:** find the record, move the rest of the block down (at most 512 bytes), count minus one. No index to
  keep right.
- **Each:** block by block through the directory (stepping over repeated entries); the values of a block are
  copied out together, **one allocation for a block** and not for every value (so an overflow key of 4,000
  values costs about 180 allocations in a scan).
- **A value of 255 bytes or more** does not fit the one-byte length. A key that has such a value keeps its
  values in a `vset.Set[string]` as today (the heavy set leaf, which stays: also for keys of more than 254 bytes).
  Both forms leave and enter the page as follows.

**Objects and kinds.** In a string map a key has one of three objects: a *single-key page* (kinds `kSet+1` to
`+6`, step 3.3), a *block-set leaf* (new kind `kSet+7`: leaf head, the key remainder as in the set leaf, and one
pointer to the set), or the heavy *set leaf* (`kSet`). The directory is a Go slice of `*block`, the blocks are
`[512]byte`: a block is a 512-byte object in the Go heap and aligned to its size.

**Transitions** (hysteresis as in the design note of the page):

| from | to | when |
|---|---|---|
| page | block-set leaf | the content no longer fits 512 bytes, all values below 255 bytes |
| page | set leaf | a value of 255 bytes or more arrives, or the remainder is too long for a page |
| block-set leaf | page | the content is 256 bytes or less |
| block-set leaf | set leaf | a value of 255 bytes or more arrives |
| set leaf | block-set leaf | the last value of 255 bytes or more is removed (checked when such a value is removed) and the remainder is short enough |
| set leaf | page | content 256 bytes or less and no long value (as in step 3.3) |

**Shrinking:** blocks that have become sparse after many removals are not merged one by one: the set is rebuilt
(all values added again to a fresh set) when its blocks are less than a quarter full on average. The set goes back to a
page long before that for most keys (at 256 bytes of content).

## 3. Why this and not the alternatives

- **One arena of values and an open-addressing table of 32-bit positions** (about 11 bytes of index a value, one
  random byte comparison for each probe): every value needs an index entry, and a removal leaves a hole to
  fill or to mark. The blocks keep the index where the data is (the hash decides the block, the scan of one block
  is the whole search) and a removal is a `memmove` in one block.
- **Chunks of values without a hash** (an array of 512-byte blocks in arrival order): a membership test would walk
  all of them, up to 180 blocks for the largest key of `street`. Rejected: Add must be cheap too.
- **Sorted values:** scans would come out sorted, but every insertion moves bytes in a long structure. Not needed
  (the order of the values of a key is unspecified).

## 3a. All the options, looked at together (added after the user's question, 2026-10-04)

The first version of this note compared the block set only with an arena and a table, with chunks and with
sorted values. It did not look at keeping `vset`, tuning it, or using `Set3` on its own. Here is everything, with the
numbers that exist. **What changed with the single-key page:** a key with up to about 40 values of 10 bytes (`street`)
or 8 values of 55 bytes (`dirs`) now lives in a page, so the inline stage (3 values) and the array stage (up to
64 values) of `vset` are mostly unused: a key reaches the value set with a few dozen values or more. Of the keys that overflow, 31 % of
`street` and 69 % of `dirs` still have up to 64 values, but they hold only 3 % and 11 % of the values; the weight is in the
hash stage.

**Measured** with `bench/cmd/ovbench` (the real overflow keys, about 1,090 and 1,380 keys; the strings allocated one by one, as an
application would; the working set is 3 to 8 MB, so a hit is a hop in L2 or L3, not a trip to the memory; a diagnosis of the containers, not of
the tree), median of 5 rounds:

| data | container | heap B/value, string bytes included | contains, hit | contains, miss | add + remove | read a value |
|---|---|--:|--:|--:|--:|--:|
| `street` | `vset` as it is | 48.6 | 65 ns | 31 ns | 63 ns | 3.9 ns |
| `street` | `Set3` on its own | 50.4 | 55 ns | 17 ns | 38 ns | 3.9 ns |
| `dirs` | `vset` as it is | 60.9 | 58 ns | 35 ns | 55 ns | 3.6 ns |
| `dirs` | `Set3` on its own | 64.1 | 46 ns | 17 ns | 37 ns | 4.0 ns |

| | option | misses after the leaf for a hit | bytes a value (`street` / `dirs`) | pointer-free | effort | verdict |
|---|---|---|---|---|---|---|
| A | keep `vset` and tune it (drop the inline stage for strings, lower the array stage) | leaf, `Set3` header, control bytes, slot group, string: 4 | about 48 / 60, a little less | no | days | gains a few percent; the structure stays what it is |
| B | `Set3` on its own, `vset` gone (the user's idea), embedded in the leaf or behind one pointer | control bytes, slot group, string: 3 (one less than A: the hop to the header is gone) | **50 / 64** (measured) | no | one or two days | **15 to 30 % faster, no smaller**: a hit 55 against 65 ns, a miss 17 against 31, add + remove 38 against 63; the memory a little larger; the pointers stay |
| C | blocks of 512 bytes with a hash by extendible hashing (this note) | directory (can sit in the leaf for a few blocks), block: 2, in one block, no pointer to a string | **17 / 40** (prediction) | **yes** | three or four days | memory -65 % / -35 % against A and B; speed to be measured |
| D | one arena of values and an open-addressing table of 32-bit positions | table slot, record: 2 | about 22 / 38 (11 bytes of index a value) | **yes** | about three days | as C in memory; a removal leaves holes, a rebuild fills them; no `memmove` in a block |
| E | sorted values in blocks | like C | like C | yes | more | scans would be sorted; nothing asks for it; an insertion moves bytes in a long structure |
| F | a different set for `Hashed` and for `Ordered` | | | | | **yes, as a consequence**: `Hashed` stays on `vset` (`map[string]*vset.Set`, there is no cache-line premise for it, only the credo as the competitor); `Ordered` gets its own for each value type |
| H | a **dictionary of values**: each different value is stored once and the keys hold numbers (4 bytes) | key: number array in a page or blocks; a dictionary lookup on add | **about 4 / 4 plus the dictionary** if the values repeat a lot (`street`: a few thousand different localities among 579,000 values), **worse** if they do not (`dirs` file names) | yes for the keys | weeks (a reference count to free values, a dictionary that is itself a map) | the biggest possible saving for repeating values, a different data structure; for the profile after version 0.8, with the corpus of the inverted index in front of us |

What decides between B and C: **B is a replacement of a container, C is a new data structure.** B is cheap and fast to try and
makes the structure a little faster, but does nothing for what the redesign is about (a pointer to every string; the garbage collector scans 16
bytes a value; the strings stay where the caller allocated them, so a string that is a piece of a big buffer keeps the whole buffer alive, which a copy
into a block does not). C is the cache-line shaped answer and the only one that removes a third of the heap of the natural mix
(`street` 113 to about 96 B/key). A and B cannot get there.

A way to decide with data and not with opinions, before the tree is touched: build the standalone `internal/strset`
(C) first, put it into `ovbench` next to A and B, and take C if it holds **at most 60 % of the bytes of B** and its hit and its add + remove cost
**at most 20 % more than B's** (55 and 38 ns on `street`). If it does not, B is the fall back and C goes to the list of what comes after
version 0.8. D can go through the same bench cheaply and is the second candidate if C's blocks are too slow on a hit.

## 4. What it costs and what it saves (prediction)

Memory a value: the bytes and one length byte, divided by the fill of the blocks, which is about 69 % for hashed
splits (ln 2) plus the directory (8 bytes a block, with about 1.5 slots a block, so 12 bytes a block, 0.5 byte a
value at 22 values).

| data | bytes a value now (headers, hash set, string bytes counted) | block set | overflow of the tree, B/key now | block set, B/key |
|---|--:|--:|--:|--:|
| `street` | 24 to 40 + 10.3 | (10.3 + 1) / 0.69 + 0.5 = 16.9 | 38.6 (+ 10.3 bytes of strings) = 48.9 | 17.0 |
| `dirs` | 24 to 40 + 25.9 | (25.9 + 1) / 0.69 + 0.5 = 39.5 | 46.0 (+ 33.7) = 79.7 | 51.4 |

So the whole tree of step 3.3 (measured in the bench without the bytes of the strings that it does not copy), expected
with the block sets and the bytes of the strings **inside**: `street` natural about **96 B/key** (measured 113
without the bytes in the overflow, 123 with them: -22 %), `dirs` natural about **162 B/key** (model 157
without the bytes of the overflow, 191 with them: -15 %). The bytes the garbage collector scans for the overflow fall
to the directories (about 0.5 B a value, now 16 B a value); in the memory tables `scannable B/key` falls by the
same: on `street` from 65 to about 20.

Speed, in the order of the weight of the case:

- **Scans:** a key with many values costs one allocation for 22 or 9 values and a walk through contiguous
  bytes: better than today's walk through headers that point to scattered strings. Expected: the 37 % of the values that sit
  in overflow keys are read faster; the scan gap to `node-layout` (see "To check later" in PLAN.md) should narrow where these keys
  are hit, which `valuesBetween` does for the big ones.
- **Add of a value to an overflow key:** hash (about 10 ns for 10 bytes), a block scan of about 10 records (about
  15 ns), an append; every 20th add a split (about 1 µs). Expected about 50 to 70 ns, today's array and hash set
  take about the same.
- **Remove:** hash, scan, `memmove` of about 250 bytes (about 10 ns). About the same as today.
- **Contains / Has:** the same as Add without the append.
- **A split's pause** is about 1 µs and happens in the set that grows; no rebuild of the whole set.

Risk: the block scan reads up to 8 cache lines (512 bytes) for a miss; the arrays of 64 headers of today read
16 lines. If `churn` on overflow keys shows it, blocks of 256 bytes (about 10 values of 10 bytes) are the first
thing to try; not before a profile.

## 5. How it is built and checked

1. **`internal/strset`**, standalone: `Set` with `Add`, `Remove`, `Contains`, `Len`, `Bytes`, `Each` (blocks
   copied out), `Strings`, a rebuild; 100 % coverage, model test against a Go map, a fuzz test, a test of
   the split at every depth up to 12, a test with values that all collide in the first hash bits (the split must stop).
2. **Microbenchmark** on the real overflow keys (about 1,090 and 1,380 keys) against `vset.Set[string]`: Add, Remove,
   Contains, Each, bytes a value.
3. **In the tree:** the block-set leaf as a leaf kind (`Objects` label "block-set leaf"), the transitions of
   section 2, tests of every transition, the reference and fuzz tests with values of 5 to 300 bytes.
4. **Measurement** as in 3.3 (the four runs) and the report against this prediction.

## 6. For the user to decide

1. Blocks of 512 bytes (with 256 as the first thing to try if `churn` asks for it).
2. The heavy set leaf stays for values of 255 bytes or more and keys of more than 254 bytes (a second form of
   overflow, kept because the one-byte length does not reach them).
3. A hash of 64 bits from `hash/maphash` with a seed per process (no stable order of values between processes; the order
   of the values of a key is unspecified anyway).
4. Order of work: this (3.4) before the other value types (3.5), as in the plan.
