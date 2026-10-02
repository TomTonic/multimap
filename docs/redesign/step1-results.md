# Step 1: the page for keys of any length (prototype results)

Prototype in `internal/vpage`, experiments in `bench/cmd/pagefill`, 2026-10-02. Nothing here is in the
tree yet. This is the report for gate 1 of [PLAN.md](PLAN.md); [STATUS.md](STATUS.md) says what
happens next. The first version of the page (sorted arrays with fence search) and its figures are
in the history at the end; the user's idea of a directory, a FAT, in the first 64 bytes made the
second version.

## What was built

A page holds the suffixes of many keys, each with one value, in one object of 128, 256, 512 or 1024
bytes without pointers (`internal/vpage/page.go`, `mutate.go`, `seq.go`):

- **Header** of 8 bytes: kind, class, count, capacity, uniform length, prefix length, start of the heap.
- **Prefix** (added in the evening, see "The page prefix"): the bytes all keys of a page share, stored once
  directly behind the header, so in the first line; the stripped keys are what the rest holds.
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

| kind | keys | leaves today B/key | pages without prefix (with router) | **pages with prefix** (with router) | fill (no prefix) | keys per page (no prefix) |
|---|---|--:|--:|--:|--:|--:|
| u64 | 4K | 24.2 | 26.9 (+11 %) | 26.9 (+11 %) | 68 % | 19.8 |
| u64 | 16K | 26.1 | 26.5 (+2 %) | 26.5 (+2 %) | 69 % | 20.1 |
| u64 | 256K | 24.1 | 26.5 (+10 %) | 26.5 (+10 %) | 69 % | 20.1 |
| str | 4K | 66.9 | 57.1 (-15 %) | 46.4 (-31 %) | 67 % | 9.3 |
| str | 16K | 60.0 | 49.0 (-18 %) | 46.3 (-23 %) | 68 % | 10.7 |
| str | 256K | 59.3 | 45.3 (-24 %) | 38.0 (-36 %) | 69 % | 11.6 |
| uuid | 4K | 75.4 | 74.1 (-2 %) | 74.1 (-2 %) | 68 % | 7.1 |
| uuid | 16K | 75.6 | 74.3 (-2 %) | 74.3 (-2 %) | 67 % | 7.1 |
| uuid | 256K | 75.5 | 74.1 (-2 %) | 74.1 (-2 %) | 66 % | 7.1 |
| email | 4K | 68.6 | 57.2 (-17 %) | 57.2 (-17 %) | 69 % | 9.2 |
| email | 16K | 72.6 | 56.8 (-22 %) | 56.2 (-23 %) | 69 % | 9.3 |
| email | 256K | 68.8 | 55.5 (-19 %) | 55.1 (-20 %) | 69 % | 9.6 |
| url | 4K | 107.5 | 102.8 (-4 %) | 99.0 (-8 %) | 70 % | 4.8 |
| url | 16K | 104.2 | 99.9 (-4 %) | 94.1 (-10 %) | 70 % | 5.0 |
| url | 256K | 97.8 | 95.1 (-3 %) | 85.6 (-12 %) | 70 % | 5.2 |
| path | 4K | 99.2 | 101.2 (+2 %) | 88.0 (-11 %) | 70 % | 4.9 |
| path | 16K | 92.2 | 95.3 (+3 %) | 79.9 (-13 %) | 70 % | 5.2 |
| path | 256K | 80.0 | 84.8 (+6 %) | 64.9 (-19 %) | 69 % | 6.0 |
| street | 4K | 65.3 | 39.9 (-39 %) | 37.9 (-42 %) | 69 % | 13.1 |
| street | 16K | 65.6 | 39.0 (-41 %) | 37.0 (-44 %) | 68 % | 13.5 |
| street | 212K | 66.5 | 36.2 (-46 %) | 34.9 (-48 %) | 69 % | 14.5 |

- **Without the prefix** (middle column) `street` -39..-46%, `str` -15..-24% and `email` -17..-22% win,
  `uuid` and `url` tie (-2..-4%), and `path` (+2..+6%) and `u64` (+2..+11%) lose.
- **With the prefix** (next section) `path` becomes -11..-19%, `url` -8..-12%, `str` -23..-36%; only
  `u64` stays above the leaves (+2..+11%): the directory costs a byte a key (17 bytes instead of 16)
  and a page of 512 bytes holds 29 integers instead of 31. The `u64` figure at 16K is the model's
  noise; the real tree will tell in step 2.
- **Fill is 66-70% everywhere**, whatever the key or the page size: the fill of random inserts into
  sorted pages (ln 2). Splitting by bytes changes nothing (±1%). A split that lets the halves fill
  their class completely (`SplitFill` 100 instead of 85) saves 2-8% (`uuid` 74.1 to 68.1) at the cost of a
  rebuild when the next key arrives; pages of 1024 bytes save about as much (`uuid` 68.3). The chunk
  size, which stands for how much the tree strips above the pages, matters for keys with long common
  prefixes: 64 keys per chunk gives `path` 77.5 bytes a key, 1024 gives 90.4; the real tree decides that.
- Two keys of the `path` corpus have a suffix over 255 bytes, which a page does not hold (the length
  byte); the tree needs a fallback for them (a bigger class with a 16-bit length, or a leaf).

## The page prefix

Added on the evening of 2026-10-02 at the user's request to try it, with the worry that it costs too
much in `churn`. A page may hold the bytes all its keys share once, right behind the header, and the keys
without them.

**Design.**
- The prefix lies in the first line (after the 8-byte header, padded to 8), so a lookup that has read the
  header has read the prefix too. The first design put it at the end of the page; the cold lookup
  then waited for a second line (`u64` 267 against 157 ns) and it was moved.
- A rebuild (growth, shrink, split, merge) takes the prefix the keys share, **if it saves at least
  `MinGain` = 16 bytes**: the saving is the difference of `need()` with and without it. The heads are
  8 bytes wide whatever the key, so a prefix shortens only the tails (and may make a page uniform);
  pages of short keys never get one, and `u64` never does. An insert in place does not change the
  prefix. An insert that does not start with it makes a rebuild with a shorter or no prefix.
- Lookups (`Get`) of a page with a prefix of up to 23 bytes compare it and cut the key's head word out
  of three words by shifting and masking, **without a branch on the prefix length** (a branch that the
  next page makes go the other way flushes the lookups the CPU has begun on other keys while this page
  loaded: with such a branch `u64` was 38% slower cold). A page with no prefix goes the old, short way.
- 100% coverage, a brute-force test of the shared bytes (`TestSharedBy`), the reference test with the
  shapes of keys (including a deep common prefix), fuzz and race detector clean.

**Memory** is in the table above (`path` -11..-19%, `url` -8..-12%, `str` -23..-36% against the
leaves; the prefix saves `path` 20 bytes a key at 256K).

**Cost of changes** (`pagefill -timing`: build and then 400,000 times delete a random key and put it
back, 256K keys in the chunk model, prefix off and on alternating, three rounds, fastest counts; a
microbenchmark for diagnosis, noise about 5%). "Rebuilds" are pages built anew, "with a new prefix" those
that had to cut every key anew because the prefix changed:

| kind | build ns/key off / on | churn ns/op off / on | rebuilds per 1000 ops off / on | with a new prefix per 1000 ops | pages off / on |
|---|--:|--:|--:|--:|--:|
| `u64` | 258 / 272 | 474 / 460 | 0 / 0 | 0 | 13055 / 13055 |
| `str` | 351 / 357 | 800 / 752 | 179 / 126 | 0.3 | 22660 / 19522 |
| `uuid` | 401 / 411 | 912 / 984 | 224 / 224 | 0.6 | 36906 / 36904 |
| `email` | 350 / 361 | 815 / 856 | 235 / 233 | 2.7 | 27447 / 27288 |
| `url` | 583 / 578 | 1306 / 1365 | 441 / 383 | 12.5 | 49828 / 45233 |
| `path` | 562 / 546 | 1306 / 1224 | 380 / 276 | 7.3 | 43529 / 34221 |
| `street` | 307 / 327 | 651 / 681 | 106 / 85 | 3.5 | 14619 / 14089 |

Churn changes by -6% (`path`, `str`) to +8% (`uuid`, whose pages are identical with and without the
prefix: noise), build by -3% to +7%; the first version, which planned with a 255-byte array and
copied keys to find the prefix, cost 5-20% more, and the plan was slimmed until a prefix that is never
chosen costs 0-2%. **The prefix changes in 0.3-12 of 1000 operations**, and each such rebuild costs
about a microsecond, about as much as the operation itself: a share of 1% at most. The pages with a
prefix are fewer and fuller, so fewer inserts hit a full page (`path` 276 rebuilds against 380).
Churn in the tree is the real test and belongs to step 2.

**Lookup** (`go test -bench Get`, Ryzen 9 7900, off and on alternating, diagnosis only):

| | cold ns off / on | hot ns off / on | cold, key absent, ns off / on |
|---|--:|--:|--:|
| `u64` | 148 / 147 | 6.2 / 6.2 | 58 / 59 |
| `uuid` | 212 / 217 | 10.5 / 9.9 | 72 / 76 |
| `path` | 247 / 249 | 11.9 / 14.6 | 180 / 151 |

Only the pages with a prefix, `path` here, cost something: +2.7 ns hot (+23%), nothing cold, and
absent keys are 16% faster (they end at the prefix). The word that reads the head of a key in one load
(instead of a loop over 8 bytes) made all lookups faster than in the morning (`u64` hot 8.8 to 6.2 ns,
`uuid` 13.0 to 10.5, `path` 14.3 to 11.9 without prefix). Still two rounds for every lookup (the
trace test mirrors it).

**What the prefix does not settle:** the cost of the prefix on lookups in the tree, where a tree level
in front of the page already strips bytes (the page prefix is what the keys of one page share beyond
that); the microbenchmark builds one run of 3 million keys, whose pages have longer prefixes than the
tree's would. Step 2 measures it.

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
| bytes per key by key kind | table above: better for every kind except `u64` (+2..+11%), with the page prefix: `street` -42..-48%, `str` -23..-36%, `email` -17..-23%, `path` -11..-19%, `url` -8..-12%, `uuid` -2% |
| head format: 8 bytes against 7 bytes and a length | 7 bytes and a length cannot order keys longer than 7 bytes; 8-byte heads with tags and a tie-break on length and tail are used |
| tail compare: lines | the tail is read in the same round as the head and value, since the directory holds its offset |
| classes and split rule | 128/256/512; splitting by bytes or count makes no difference (±1%); filling halves to 100% or 1024-byte pages gain 2-8% |
| insert and delete cost | see above: 47-199 ns for an insert and delete in a page with room; churn and build with the prefix within -6..+8% of without |
| bloom filter | replaced by the tags: 1 byte a key instead of 8 bytes a page, and they find the key, too |

## Reading against gate 1

Gate 1 says: continue only if memory per key does not exceed today's leaves for any key kind and a
lookup touches at most two lines after the head (restated by the user as at most two rounds, see
STRATEGY.md R5).

- **Memory:** met for every kind except `u64` (+2..+11%): with the page prefix `path` is -11..-19%
  and `url` -8..-12% below the leaves. `u64` pays a byte a key for tags that cut its lookup from 3
  rounds to 2 (and, on the Ryzen, its cold lookup from 283 to 148 ns); its memory stays at 26.5 bytes a
  key against 37 for `btree-map`.
- **Lookups:** every lookup takes two rounds; the literal count of lines (3 to 4 after the first for
  general pages) is not met, the substance is.

## History: version 1 (no directory)

Sorted arrays of heads and values searched with fences, a 64-bit bloom filter and a 16-byte
header (16 bytes a key uniform, 19 general). Memory per key was within ±1% of the version above
(`u64` 24.8, `str` 44.3, `uuid` 69.4, `email` 54.8, `url` 95.4, `path` 84.7, `street` 35.3 at
256K). A lookup took 3 to 4 rounds and 283-409 ns cold. The code of that version is in
`C:\temp\bench-win\vpage-v1` (not in git).
