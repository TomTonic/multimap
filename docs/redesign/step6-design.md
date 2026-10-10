# Step 6: the routing layer — design note (2026-10-10, before any code)

Words as in [GLOSSARY.md](GLOSSARY.md). Working rules of PLAN.md apply, in particular the user's of 2026-10-10: the structured keys
(str, email, url, path, street, dirs, links) decide, random keys only stay in the band; hot paths in "assembler thinking"; speed only
from rtcompare. This note is to be read and decided by the user before anything is built.

## 1. Why the routing layer

The profiles of every scenario ([profiles-2026-10.md](profiles-2026-10.md)) show the descent through byte nodes as the largest share of
every operation of the structured keys: 39 to 55 % of a lookup, 24 to 40 % of a churn, 16 to 36 % of a build, and, as the walk of the
range scan, 43 to 68 % of a range. A cheaper routing layer helps all four operations; nothing else does.

## 2. What the routing layer looks like today (shape, `MKSHAPE`, as built from the corpus)

| case | nodes a lookup passes | byte nodes | many-key pages | keys a page: mean / median / 10 % / 90 % | single-key pages | block B a key |
|---|--:|--:|--:|--|--:|--:|
| street single-value 4K / 65K | 3.0 / 4.9 | 100 / 1,499 | 536 / 8,894 | 7.1 / 5 / 2 / 17 ; 6.9 / 4 / 2 / 16 | 265 / 4,516 | 30 / 28 |
| dirs single-value 4K / 65K | 7.6 / 11.4 | 221 / 3,291 | 792 / 10,289 | 4.5 / 3 / 2 / 9 ; 5.8 / 4 / 2 / 12 | 514 / 5,701 | 54 / 41 |
| links single-value 4K / 65K | 2.7 / 4.3 | 101 / 1,832 | 670 / 9,807 | 5.4 / 4 / 2 / 12 ; 6.1 / 4 / 2 / 14 | 468 / 5,887 | 36 / 33 |
| url single-value 4K / 65K | 5.9 / 8.2 | 314 / 4,844 | 881 / 13,537 | 3.7 / 3 / 2 / 7 ; 4.0 / 3 / 2 / 7 | 837 / 10,932 | 86 / 77 |
| natural (street, dirs, links, url at 65K) | 5.1, 11.6, 5.7, 8.5 | | | 5.5, 4.8, 2.6, 3.3 (mean) | | 39, 59, 149, 104 |

Two things stand out:
- **The pages are nearly empty.** A full page of 512 bytes holds 24 keys on street single-value (measured: the 443 pages of 512
  bytes at 65K), but the median page holds 3 or 4, and 34 to 45 % of all pages are single-key pages of one key — even in single-value
  maps. The cause is the burst: a page that
  overflows becomes a byte node with one child per distinct next byte, and most of those children get one to four keys. How full
  a page is depends on how the keys spread over the next byte, not on the page size.
- **The trees are deep for path-like keys:** 11.4 nodes for dirs and 8.2 for url at 65K keys, every one a dependent load (a cache
  miss as soon as the tree does not fit the cache), because every divergence point of the keys becomes a node as soon as the keys
  under it do not fit one page, and the small pages make that happen early.

Fewer, fuller pages mean fewer nodes and fewer levels. That is the lever of this step.

## 3. The previous attempt and what it taught (branch `leaf-pages`, `afd1137` and `node-pages`, 2026-09-26 to 10-02; step 3's R8)

Range nodes after the B-trie of Askitis and Zobel: child i of a range node holds the keys whose byte at the node's position lies in
[start_i, start_{i+1}); a full page splits by count at a byte boundary into two pages that share the range; only a page whose keys all
share that byte gets a subtree. The node did not consume the byte (its children held full keys from the same depth); its layout was a
256-bit bitmap of the range starts and the children, found by a floor-rank over the bitmap, in classes of 9, 25, 57 and 256
children (128 to 2,104 bytes).

What it brought and what it cost (rtcompare against `main`, leaf-pages/bench/results-leaf-pages/README.md and the commit message):
1. **Pages full independent of the size**: keys a page 21 for u64 (was 1.1 at 16K), about 10 of 14 for strings; ranges 2 to 6 times
   as fast, memory 20 to 80 % less.
2. **Point queries and writes of small maps lost**: point 0.62 to 0.97 at 4K and 16K, churn 0.53 to 0.96, build 0.37 to 0.63. The
   findings name the cost: a range node cost more per level than `main`'s byte nodes. Likely reasons, read from its design and not
   measured one by one: the floor-rank over a bitmap of 256 bits for every node size, and a level that does not consume its byte, so
   the page below checks it again. And the splits built item lists of the page's entries (the plan of 2026-09-26 named "page splits
   without item lists" as the optimization still to do).
3. **A weak band at 16K to 64K keys**: one level more (R256 → R8 → page), about 2.5 cache lines more a lookup; the point lookups
   turned from losing to winning between 32K and 64K keys.
4. Pages of 256 bytes were worse than pages of 512 (one level more).
5. A range node of 128 bytes held only 2.6 to 3.1 ranges in step 3 (artstr): 3 to 14 bytes a key, 11 to 16 % of the objects.
6. Keys with several values cut the pages between them into pieces (one leaf a key with several values). **That cause is gone
   since step 5**: a page holds several values of a key.

## 4. The proposal: byte nodes whose children may be ranges ("split instead of explode")

Keep today's byte nodes and their descent, and change one thing: **a child of a byte node may stand for a range of bytes**, not
only for one. Concretely:
- **Node layout unchanged.** N5 and N12 keep their sorted byte arrays, N26 and N58 their bitmaps; a byte in the array (a bit in the
  bitmap) is the **start** of the child's range, which ends before the next start. A child for a single byte is a range of one.
- **The search is a floor instead of an exact match**: N5/N12: the number of starts at most b, by SWAR on the 8 or 16 bytes (a
  subtraction and a mask a word, branch-free, the same instructions as `swar.Index8`); N26/N58: the rank of the highest set bit at
  most b (as today's `Has` plus `Rank`, one word more masked). No extra level and no extra cache line: this answers point 2 of
  section 3, where the bitmap of 256 bits and the separate node kind were the cost.
- **Who consumes the byte**: a child whose range is one byte and which is a byte node consumes the byte as today. A **page** under a
  range child holds its keys **from the node's byte on** (the byte is the first byte of the page's key part or of each remainder), so
  the page can hold keys of several bytes; it checks the byte itself, inside the search it does anyway. So the descent passes no
  level that does not consume a byte (point 2 of section 3).
- **A burst splits instead of exploding**: a page that overflows is replaced by a byte node with **two** range children (the page cut
  in two at the byte boundary nearest the middle of its entries), or, when its keys all share their next byte, by a node on the first
  byte where they differ, as today. Later a full page under a range splits again into two ranges of the same node, until the node is
  full and grows as today (N5 → N12 → N26 → N58 → N256). The split moves the second half of the page's bytes into a new page in one
  copy of the key area and one of the values (no item lists: point 2 of section 3).
- **A key with its own object** (a value overflow, or a key too long for a page) gets a range of its own byte, a single slot, so it
  does not cut its neighbours' range in two (the user's idea of 2026-10-03, "single slots in the range node").
- **Merges**: two neighbouring range children whose pages together need at most half a page merge into one range (the hysteresis of
  E6, `mergeFill`, unchanged).
- **The scans**: the cursor walks the children of a node as today; a range child's page gives its values in order as today.

What it does not change: the pages (internal/page), the value overflow, the API, the byte nodes' sizes (so the objects stay within
R1 to R3).

## 5. Prediction (to be checked by a model first, then by rtcompare)

- **Pages fuller**: keys a page from a median of 3 to 4 to 8 to 15 (a split page is half full and fills; the B-tree expectation of
  69 %); single-key pages from a third to a half of all pages to a few percent (only keys with their own object).
- **Fewer nodes and levels**: nodes a lookup from 4.9 to 3 to 4 (street 65K), 11.4 to 6 to 8 (dirs 65K), 8.2 to 5 to 6 (url 65K);
  byte nodes 40 to 60 % fewer.
- **Memory**: block bytes a key 10 to 25 % less (fewer pages, less slack, fewer nodes).
- **Speed** (structured keys, against the state before the step, M1 job `c1` and the same on the PC): lookups and ranges 10 to 30 %
  faster from about 64K keys on; churn about equal (fewer bursts, but splits and merges of ranges); build ±10 % (a split costs one
  copy instead of a burst into many pages). **For point queries of small maps no improvement is predicted, and a loss is possible**
  (section 5a).

### 5a. The risk for point queries (added on the user's question, 2026-10-10)

Fuller pages work against a point query in two ways: `locate` walks the keys of a page one after the other, so 8 to 15 keys a page
instead of 3 to 4 mean more compares (the page search is already 21 to 43 % of a lookup); and in a page of 256 to 512 bytes the
value lies at the end of the object, a cache line away from the key, where a page of 64 bytes has both in one line (the "serial cache
hops" of node-pages). Against that stands only the saving of levels, and a level costs a few nanoseconds while the tree fits L1/L2
(4K keys) and an L3 or memory access from about 64K keys on. That is the pattern of `afd1137` (losing at 4K and 16K, winning from
32K to 64K), and cheaper levels do not change it. So:
- the model counts, besides the nodes a lookup, **the remainders compared in the page a lookup** and **the cache lines touched a
  lookup** (node lines, page lines up to the key, the value's line);
- the gate includes the small maps: at 4K keys no more work a lookup than today (compares and lines);
- **when a range child's page splits** (by bytes or by keys) is chosen from the model, as one fixed, explained rule, to balance the
  levels against the page search; the lesson "pages of 256 bytes were worse" is the check of that rule;
- the page search may have to come into this step: the offsets of the remainders follow from the key-length list by a SWAR prefix
  sum, which allows a binary search in the page (3 to 4 compares instead of a walk) without a byte more in the page (an idea, not
  measured).
- Random keys (u64, uuid): pages fuller as well; only checked to stay in the band.

## 6. How it is checked, in this order

1. **A model first** (no change of the tree): a program that builds the tree of section 4 from the sorted keys of street, dirs, links
   and url at 4K, 16K and 65K (both value profiles, the page sizes of `internal/page` with `NeedStrings`/`NeedFixed`), inserting in the
   corpus order of the bench's build, and counts nodes a lookup, remainders compared and cache lines touched a lookup, pages, keys a page, single-key
   pages and block bytes a key, next to the same counts of today's tree. **Gate**: dirs and url at 65K with at least 30 % fewer nodes a lookup and fewer bytes a key, street and links
   not worse, **and at 4K for every corpus no more compares and cache lines a lookup than today** (section 5a); else stop and report.
2. Then the code, in steps with their own tests: the floor search and pages under range children (lookup, scan); the split; the merge;
   single slots. Every step keeps the tests, 100 % coverage, `-race`, lint, and the model's counts as a check of the real tree.
3. rtcompare on the PC and the M1 against the state before the step (`c1`'s commit), on street, dirs, links, url (both profiles; 4K,
   16K, 65K), and the band against the B-trees; memory with three rounds.

## 7. Not in this step

The other points of PLAN.md's step 6: removing N5 (memory only: 2 to 4 bytes a key in dirs and url, no speed) and the tails of long
common prefixes (90 to 307 nodes with 13 bytes or more at 65K in dirs and url, few) wait until the model shows what is left of them.
The search in the page and the value sets (profiles-2026-10.md) are separate levers after this step.

## 8. Result of the model (2026-10-10): the gate is missed — fuller pages, but not one level less

`bench/cmd/rangemodel` (counts, no times; uint64 values; tables in `bench/results-layout/step6-model/model-uint64.md`). Its model of
today's tree reproduces the measured shape exactly (street single-value 65K: 4.87 nodes a lookup, 13,410 pages, 28.3 block bytes a key;
dirs 11.37 nodes), so its counts of the proposal can be trusted as counts.

| at 65K keys, single-value | nodes a lookup | compares walk / binary | cache lines walk | keys a page (mean) | one-key pages | objects a key | bytes a key |
|---|--|--|--|--|--|--|--|
| street: today → ranges 512 | 4.87 → **4.87** | 6.2 → 9.9 / 2.7 → 3.5 | 10.8 → 11.0 | 4.9 → 14.4 | 4,516 → 315 | 0.227 → **0.092** | 28.3 → 24.6 |
| dirs | 11.37 → **11.37** | 4.7 → 6.7 / 2.5 → 3.0 | 18.9 → 20.6 | 4.1 → 8.4 | 5,701 → 1,168 | 0.297 → **0.173** | 40.9 → 38.7 |
| links | 4.32 → **4.32** | 5.3 → 8.6 / 2.6 → 3.3 | 10.2 → 10.6 | 4.2 → 11.9 | 5,887 → 577 | 0.267 → **0.112** | 33.1 → 28.8 |
| url | 8.19 → **8.19** | 2.9 → 4.0 / 2.0 → 2.4 | 15.9 → 16.0 | 2.7 → 5.2 | 10,932 → 1,633 | 0.449 → **0.268** | 76.7 → 70.5 |

Splitting above 384 or 256 bytes adds levels (dirs 11.58 and 11.87) and gives back the memory. Natural and 4K alike.

**Why the levels stay:** a node on the path to a key exists because the keys below its prefix do not fit one page; how the
neighbouring subtrees are packed into pages does not change that. Ranges pack the siblings (objects a key 40 to 65 % fewer: the walk
of a range scan, 43 to 68 % of its time), but the depth comes from the subtree sizes along the key paths, and for path-like keys
(dirs 11.4, url 8.2) from chains of small nodes (dirs: 2,304 of its 3,291 nodes are N5 with 2.8 children on average). **The prediction
of section 5 ("dirs 11.4 to 6 to 8 nodes") was wrong**; section 5a's risk is real: a lookup compares more and touches 1 to 9 % more
cache lines, for the same number of nodes. The binary search in the page takes the compares back to about today's (2.4 to 3.5), not
the lines.

So the proposal is a **range-scan and memory** change (fewer objects to walk, 5 to 15 % less memory), with slightly worse point
queries; it is not the routing change this step was meant to be. Depth for path-like keys needs nodes that discriminate on more than
one byte position each. Stopped here for the user's decision.
