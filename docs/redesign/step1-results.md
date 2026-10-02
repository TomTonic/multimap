# Step 1: the page for keys of any length (prototype results)

Prototype in `internal/vpage`, experiments in `bench/cmd/pagefill`, 2026-10-02. Nothing here is in the
tree yet. This is the report for gate 1 of [PLAN.md](PLAN.md); [STATUS.md](STATUS.md) says what
happens next. The first version of the page (sorted arrays with fence search) and its figures are
in the history at the end; the user's idea of a directory, a FAT, in the first 64 bytes made the
second version.

## What was built

A page holds the suffixes of many keys, each with one value, in one object of 128, 256, 512 or 1024
bytes without pointers (`internal/vpage/page.go`, `mutate.go`, `seq.go`):

- **Header** of 8 bytes: kind, class, count, capacity, uniform length, start of the heap.
- **Directory (the FAT)** right after it, one entry per key in key order: in a *uniform* page (all
  suffixes of one length of 1 to 8 bytes: integers) one tag byte, a hash of the key's head word; in a
  *general* page 4 bytes: the tag, the length of the suffix and the offset of its tail. It lies in
  the first line or two of the page.
- **Entries** after it, sorted: the head word (first 8 bytes of the suffix, zero padded) and the
  value; the bytes of a suffix beyond the first 8, its tail, are in a heap that grows down from the
  end of the page. A key costs its suffix plus 12 bytes (general) or 17 bytes in all (uniform).
- **Lookup:** compare the tag of the key with all tags in the directory (8 at a time, 2 entries a
  word in the general flavor, no search of the sorted arrays), then read the head, the value and
  the tail of the entry that matched, which are independent of each other: **two rounds**. A key
  that is not there ends the lookup after the first, with a chance of 1 - n/256.
- **Ordered operations** (insert, delete, range start) still search the sorted heads, with the fence
  scheme of `node-pages`.
- **Capacity chosen by a rebuild** so that slots and heap fill up together; an insert that finds no
  slot or no heap rebuilds the page (also compacting it), grows it into the next class, or says
  `Full` so the tree splits it. Deletes shrink a page that would fit the next smaller class at 70%.
- **Run** (`seq.go`) is a sorted sequence of pages with split and merge, standing in for a range node
  in the tests and the experiments.
- **Tests:** 100% coverage, 9 key shapes against a sorted map (with and without splitting by
  bytes), a fuzz test, race detector and lint clean.

## Memory per key against the leaves of node-pages

The sorted keys of each bench corpus are cut into chunks of 256 keys, standing for the keys below one
range node; what all keys of a chunk share is stripped (the tree holds it in a node); the stripped
suffixes are inserted into a run of pages in the random order of the corpus. Router bytes are added
per chunk in the sizes of `rnode.go`. The model reproduces what is known: for `u64` with the first
version of the page (16 bytes a key, as in `node-pages`) it came out at 24.8 bytes a key, the real tree
(`objstat`) at 24.1. "Today" is the tree of `node-pages` with one value per key
(`objstat-node-pages.md`, same corpora, same seed).

| kind | keys | leaves today B/key | pages B/key (with router) | change | fill | keys per page | shared prefix, B/key a page could save | with that prefix |
|---|---|--:|--:|--:|--:|--:|--:|--:|
| u64 | 4K | 24.2 | 26.9 | +11 % | 68 % | 19.8 | 0.1 | 26.8 (+11 %) |
| u64 | 16K | 26.1 | 26.5 | +2 % | 69 % | 20.1 | 0.7 | 25.8 (-1 %) |
| u64 | 256K | 24.1 | 26.5 | +10 % | 69 % | 20.1 | 0.2 | 26.3 (+9 %) |
| str | 4K | 66.9 | 57.1 | -15 % | 67 % | 9.3 | 5.1 | 52.0 (-22 %) |
| str | 16K | 60.0 | 49.0 | -18 % | 68 % | 10.7 | 2.0 | 47.0 (-22 %) |
| str | 256K | 59.3 | 45.3 | -24 % | 69 % | 11.6 | 4.7 | 40.6 (-32 %) |
| uuid | 4K | 75.4 | 74.1 | -2 % | 68 % | 7.1 | 1.2 | 72.9 (-3 %) |
| uuid | 16K | 75.6 | 74.3 | -2 % | 67 % | 7.1 | 1.0 | 73.3 (-3 %) |
| uuid | 256K | 75.5 | 74.1 | -2 % | 66 % | 7.1 | 1.0 | 73.1 (-3 %) |
| email | 4K | 68.6 | 57.2 | -17 % | 69 % | 9.2 | 0.9 | 56.3 (-18 %) |
| email | 16K | 72.6 | 56.8 | -22 % | 69 % | 9.3 | 0.9 | 55.9 (-23 %) |
| email | 256K | 68.8 | 55.5 | -19 % | 69 % | 9.6 | 0.9 | 54.6 (-21 %) |
| url | 4K | 107.5 | 102.8 | -4 % | 70 % | 4.8 | 3.3 | 99.5 (-7 %) |
| url | 16K | 104.2 | 99.9 | -4 % | 70 % | 5.0 | 4.8 | 95.1 (-9 %) |
| url | 256K | 97.8 | 95.1 | -3 % | 70 % | 5.2 | 6.9 | 88.2 (-10 %) |
| path | 4K | 99.2 | 101.2 | +2 % | 70 % | 4.9 | 9.0 | 92.2 (-7 %) |
| path | 16K | 92.2 | 95.3 | +3 % | 70 % | 5.2 | 10.4 | 84.9 (-8 %) |
| path | 300000 | 80.0 | 84.8 | +6 % | 69 % | 6.0 | 12.7 | 72.1 (-10 %) |
| street | 4K | 65.3 | 39.9 | -39 % | 69 % | 13.1 | 1.7 | 38.2 (-42 %) |
| street | 16K | 65.6 | 39.0 | -41 % | 68 % | 13.5 | 1.9 | 37.1 (-43 %) |
| street | 212449 | 66.5 | 36.2 | -46 % | 69 % | 14.5 | 1.8 | 34.4 (-48 %) |

- **Shorter keys win clearly**: `street` -39..-46%, `str` -15..-24%, `email` -17..-22%.
- **Long keys tie**: `uuid` -2%, `url` -3..-4%.
- **`path` loses 2-6%**, and **`u64` 2-11%**: the directory costs a byte a key (17 instead of 16
  bytes) and a page of 512 bytes holds 29 integers instead of 31. The `u64` figure at 16K is the
  model's noise; the real tree will tell in step 2.
- **A page could hold the prefix its keys share once.** The column "shared prefix" is what that would
  save, measured on the finished pages (`(n-1) x` the common prefix of first and last key; not
  implemented). It would make `path` -7..-10%, `url` -7..-10% and `str` -22..-32%.
- **Fill is 66-70% everywhere**, whatever the key or the page size: the fill of random inserts into
  sorted pages (ln 2). Splitting by bytes changes nothing (±1%). A split that lets the halves fill
  their class completely (`SplitFill` 100 instead of 85) saves 2-8% (`uuid` 74.1 to 68.1) at the cost of a
  rebuild when the next key arrives; pages of 1024 bytes save about as much (`uuid` 68.3). The chunk
  size, which stands for how much the tree strips above the pages, matters for keys with long common
  prefixes: 64 keys per chunk gives `path` 77.5 bytes a key, 1024 gives 90.4; the real tree decides that.
- Two keys of the `path` corpus have a suffix over 255 bytes, which a page does not hold (the length
  byte); the tree needs a fallback for them (a bigger class with a 16-bit length, or a leaf).

## Cache lines and rounds of a lookup

A trace of `Get` (a mirror in `bench_test.go`, checked against `Get`) for pages of 512 bytes with random
keys. A round is a set of loads that can all be issued at once; each round costs about one memory latency
when the page is not in a cache.

| keys | lines read, key present | rounds | before the directory (version 1) |
|---|--:|--:|--:|
| `u64` (uniform) | 2.9 | 2 | 5.0 lines, 3 rounds |
| `uuid` | 4.1 | 2 | 4.9 lines, 4 rounds |
| `path` | 5.1 | 2 | 4.8 lines, 3.6 rounds |

**Every lookup now takes two rounds**: the directory, then the head, value and tail together.
Lines within a round are free in this measurement (the cold timings below do not follow the line
count). The first wording of gate 1 (at most two lines after the head) is not met by general pages
(3 to 4 lines after the first), but its purpose, few dependent waits, is; the user restated the rule as
"at most two rounds" (R5, STRATEGY.md), where a round is a set of loads that can be issued together. A key that is not there takes one round when its tag matches none of the page's (a chance of
1 - n/256), and the timing of those lookups is in the table; `path` keys share their first 8 bytes
often, so the tags rule out fewer of them (187 ns).

Timing, `go test -bench` on the Ryzen 9 7900, for diagnosis only (not a claim of speed): one lookup
in a random page of 40-70 MB of pages (cold) or in one page (hot), for a key that is there and for
one that is not (a random key of the same kind), and a leaf as a 64-byte object of an array as large as
the keys need, one cache miss:

| keys | page, cold | page, hot | page, cold, key absent | leaf in an array, cold | version 1, cold / hot |
|---|--:|--:|--:|--:|--:|
| `u64` | 157 ns | 8.8 ns | 66 ns | 46 ns | 283 / 13.7 ns |
| `uuid` | 246 ns | 13 ns | 90 ns | 65 ns | 317 / 23 ns |
| `path` | 259 ns | 14 ns | 187 ns | 54 ns | 409 / 51 ns |

Cold lookups are 23-44% faster than with version 1, hot ones 36-72%. **The leaf figure is a lower
bound** (no routing above it, a perfect array), and the page figure has no routing either; whether a
tree of general pages keeps up with a tree of leaves is a question for step 2. The uniform page has now
fewer rounds than the integer page of `node-pages`, which measured 0.84-0.88 of the leaves at 16K-64K
keys and 1.08-1.15 at 256K (`bench/results-layout/node-pages`, `n3-lp` and `n3-u64`).

What an experiment of version 1 showed and still holds: reading every line of the page at once after
the header changed nothing (291 against 283 ns). The cost is dependent loads and the instructions the
search runs between them, not missing lines.

## Cost of changes

`go test -bench Mutate`, one core, nothing else running; version 2 (in brackets version 1):

| keys | insert and delete one key in a page with room | delete and insert back | fill a page from empty through its classes, then split it |
|---|--:|--:|--:|
| `u64` | 47 ns (43) | 68 ns (63) | 1.65 us (1.6) |
| `uuid` | 170 ns (170) | 276 ns (283) | 0.96 us (1.0) |
| `path` | 199 ns (209) | 243 ns (227) | 1.43 us (1.5) |

The directory adds a byte (or four) to move on each insert and delete, and the numbers show it
only for `u64`. For comparison, a leaf is an allocation per key, which was not measured here. The
deletes pay for the check whether the page fits a smaller class (a scan of the length bytes).

## Questions of the plan, answered

| question | answer |
|---|---|
| bytes per key by key kind | table above: better for `street`, `str`, `email`, a tie for `uuid`, `url`, 2-6% worse for `path` and `u64` (a page prefix would turn `path` to -10%) |
| head format: 8 bytes against 7 bytes and a length | 7 bytes and a length cannot order keys longer than 7 bytes; 8-byte heads with tags and a tie-break on length and tail are used |
| tail compare: lines | the tail is read in the same round as the head and value, since the directory holds its offset |
| classes and split rule | 128/256/512; splitting by bytes or count makes no difference (±1%); filling halves to 100% or 1024-byte pages gain 2-8% |
| insert and delete cost | see above: 47-199 ns for an insert and delete in a page with room |
| bloom filter | replaced by the tags: 1 byte a key instead of 8 bytes a page, and they find the key, too |

## Reading against gate 1

Gate 1 says: continue only if memory per key does not exceed today's leaves for any key kind and a
lookup touches at most two lines after the head (restated by the user as at most two rounds, see
STRATEGY.md R5).

- **Memory:** met for `street`, `str`, `email`, `uuid`, `url`; **not met for `path` (+2..+6%) and
  `u64` (+2..+11%)**. A page prefix turns `path` into -7..-10%. `u64` pays a byte a key for
  tags that cut its lookup from 3 rounds to 2; its memory stays at 26.5 bytes a key against 37 for
  `btree-map`.
- **Lookups:** every lookup takes two rounds; the literal count of lines (3 to 4 after the first for
  general pages) is not met, the substance is.

## History: version 1 (no directory)

Sorted arrays of heads and values searched with fences, a 64-bit bloom filter and a 16-byte
header (16 bytes a key uniform, 19 general). Memory per key was within ±1% of the version above
(`u64` 24.8, `str` 44.3, `uuid` 69.4, `email` 54.8, `url` 95.4, `path` 84.7, `street` 35.3 at
256K). A lookup took 3 to 4 rounds and 283-409 ns cold. The code of that version is in
`C:\temp\bench-win\vpage-v1` (not in git).
