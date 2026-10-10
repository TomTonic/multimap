# Where the time goes: profiles of every scenario (2026-10-10)

Words as in [GLOSSARY.md](GLOSSARY.md). Asked by the user after the decision for the full page (autotune-design.md): before working on
any of the open gaps, profile all scenarios, so that a change for one gap is not one that is bad for another.

## Method

Code at `6e76275` (`cacheline`: the full page, `locate`, the hysteresis). WSL on the PC, nothing else running, 2026-10-10 07:46 to 08:04.
For each of the 8 key kinds (u64, str, uuid, email, url, path, street, dirs), both value profiles and 4,096, 16,384 and 65,536 keys,
four CPU profiles of about 3 seconds each (the probes of `bench/cmd/bench`, `prof-all.sh`):
- **lookup**: `MKLOOK` (all values of random existing keys, as the bench's valuesFor);
- **churn**: `MKPROBE`, the steady-state cycle of the bench's churn stream, replayed for at least 3 seconds;
- **build**: `MKPROBE`, the bench's build stream from an empty map, repeated for 3 seconds (added to the probe for this);
- **range**: `MKRANGE`, ranges of 100 consecutive keys, as the bench's valuesBetween.

Every function's own (flat) time is put into one category by its name (`classify.py`):
- **desc**: the descent through byte nodes (`Tree.find`, node searches, prefix checks);
- **search**: the search in a page (`locate`, `compare`, the key checks of a one-key page);
- **head**: the page's head helpers (`mem`, `lay`, `cpl`), used by every page operation;
- **read**: reading the values of a page;
- **pchg**: changing a page (`Add`, `Remove`, `Widen`, `Build`, `regrow`);
- **tree**: changing the tree (bursts, merges, building subtrees, new nodes, splits);
- **scan**: the range walk (`Cursor`, bounds, values of a page);
- **sets**: the value sets of keys with many values (value overflows);
- **GC**: allocation and garbage collection; **mem**: `memmove`, `memequal`, compares of the runtime; **rt**: the rest of the runtime;
- **harn**: the probe's own loop.

**These are shares, not speeds.** The probes' times are not comparisons (MEASURING.md: speed claims come only from rtcompare); the
tables per scenario are in `bench/results-layout/profiles-2026-10/shares.md`, the split of the range walk in `range-split.txt`.

## Shares, mean over the 8 key kinds (%)

| operation | values | n | desc | search | head | read | pchg | tree | scan | sets | GC | mem | harn |
|---|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| lookup | single-value | 4,096 | 37 | **38** | 2 | 15 | | | | | | 1 | 6 |
| | | 65,536 | **56** | 25 | 1 | 12 | | | | | | 1 | 4 |
| | natural | 4,096 | **36** | 26 | 2 | 16 | | | | 7 | | 1 | 12 |
| | | 65,536 | **52** | 19 | 1 | 14 | | | | 7 | | 1 | 7 |
| churn | single-value | 4,096 | **21** | 19 | 8 | 3 | 16 | 14 | | | 3 | 7 | 6 |
| | | 65,536 | **39** | 11 | 5 | 1 | 11 | 12 | | | 2 | 6 | 12 |
| | natural | 4,096 | **25** | 18 | 8 | 2 | 15 | 12 | | 2 | 1 | 7 | 8 |
| | | 65,536 | **38** | 12 | 3 | 1 | 8 | 9 | | 6 | | 4 | 19 |
| build | single-value | 4,096 | 14 | 15 | 8 | 2 | **18** | 13 | | | 8 | 6 | 6 |
| | | 65,536 | **28** | 9 | 7 | 1 | 12 | 13 | | | 5 | 7 | 12 |
| | natural | 4,096 | **19** | **19** | 7 | 2 | 14 | 12 | | 2 | 5 | 7 | 8 |
| | | 65,536 | **34** | 11 | 5 | 1 | 8 | 10 | | 3 | 2 | 4 | 19 |
| range | single-value | 4,096 | 5 | | 6 | 3 | | | **77** | | | 7 | |
| | | 65,536 | 6 | | 6 | 2 | | | **79** | | | 5 | 1 |
| | natural | 4,096 | 5 | | 5 | 2 | | | **70** | 15 | | 2 | |
| | | 65,536 | 5 | | 4 | 2 | | | **69** | 17 | | 2 | |

16,384 lies between the two sizes everywhere. Build has 2 to 8 % more in "rt" (the runtime's page allocation and clearing) besides GC.

**The range walk split** (`range-split.txt`, single-value): the **walk through the byte nodes** (`NextPage`, `enter`, `nextChild`) is 43
to 70 % of a range (u64 at 4,096 keys: 26 %, where the bounds take 40 %), the bounds 2 to 23 %, the values of the pages 3 to 8 %, the loop 8 to 23 %. The walk is largest where pages hold few
keys: u64 at 65,536 keys 70 % (2.4 keys a page, see page-search-design.md "Coding of further values"), uuid 65 to 70 %, url 59 to 68 %.

**By kind** (`shares.md`): the descent is larger for path-like keys (url, path, dirs: 29 to 33 % of a single-value churn at 4,096
keys, 42 to 52 % at 65,536) and for long keys; the page search is largest for u64 (38 % of a churn, 58 % of a lookup at 4,096) and for
keys with long common prefixes (str, street, email: 23 to 24 % of a churn at 4,096). Changing the tree (bursts, merges) is largest for
uuid (29 % of a single-value churn at 4,096).

## What it says for the open gaps (PLAN.md, "Open after step 5")

1. **The routing part is the largest share in every operation** from 16,384 keys on, and at 4,096 keys for path-like keys: 36 to 56 %
   of a lookup, 21 to 39 % of a churn, 14 to 34 % of a build, and as the walk 43 to 70 % of a range. A change that makes the routing
   part cheaper (fewer or denser byte nodes, pages that hold more keys where random keys leave them small) helps all four operations
   at once; it conflicts with none. That is step 6 of the plan (the routing layer), which waited behind step 5.
2. **The search in the page** is 19 to 38 % of a point operation at 4,096 keys (u64 most) and 9 to 25 % at 65,536, and **nothing** of a
   range (the walk compares bounds, not keys). A faster search helps lookups and writes and is neutral for ranges **as long as the page
   does not grow**: a list in the page costs bytes, so pages hold fewer keys, more pages and nodes are walked, and the walk is the
   largest part of a range (the fingerprint list cost +1.5 to +2.9 B a key and more bursts). A search that adds no bytes to the page
   has no such conflict.
3. **The writes of single-value maps at 4,096 keys** (the band against btree-map, dirs and url): the descent 29 to 33 %, changing the
   page 12 to 15 %, changing the tree 15 to 16 %, the page search 10 to 15 %; in build, GC and the runtime's allocation 19 to 24 %. There is no single
   part to blame; the descent is the largest for these keys, so step 6 is the first lever here too.
4. **The single-value ranges of u64** (0.28 to 0.40 of btree-map at 16,384 and 65,536 keys): the walk is 70 % there, because random
   keys leave pages of 2 to 3 keys below the byte nodes. This is the routing part again, not the scan code.

## Weighted as the user asks (2026-10-10): structured keys decide, random keys only stay in the band

The user's two rules of 2026-10-10 (PLAN.md, working rules): random keys (u64, uuid) do not set the course, a user takes the hashed
map for them; and the hot paths are written in "assembler thinking" (no reflection, interfaces, yield chains, closures, channels,
few methods). The shares over the structured keys only (str, email, url, path, street, dirs), single-value / natural:

| operation | n | desc | search | head | read | pchg | tree | GC + rt |
|---|--:|--:|--:|--:|--:|--:|--:|--:|
| lookup | 4,096 | 41 / 39 | 36 / 28 | 2 / 2 | 13 / 15 | | | |
| | 65,536 | 55 / 55 | 28 / 21 | 1 / 1 | 10 / 12 | | | |
| churn | 4,096 | 24 / 26 | 18 / 21 | 8 / 7 | 2 / 2 | 16 / 15 | 13 / 12 | 5 / 1 |
| | 65,536 | 40 / 40 | 14 / 13 | 4 / 3 | 1 / 1 | 11 / 8 | 11 / 8 | 2 / 0 |
| build | 4,096 | 16 / 20 | 16 / 20 | 7 / 6 | 2 / 2 | 16 / 14 | 13 / 12 | 17 / 9 |
| | 65,536 | 30 / 36 | 11 / 12 | 7 / 4 | 1 / 1 | 12 / 8 | 12 / 10 | 10 / 4 |
| range (walk of the scan 43 to 64 %) | 4,096 to 65,536 | 6 to 8 | 0 | 4 to 6 | 1 to 3 | | | |

The order of the levers stays: the routing part first (it is the largest share of every operation and helps all four), the search
in the page second (21 to 36 % of a lookup at every size, without growing the page), the changes of page and tree third. Point 4
above (the u64 ranges) loses its weight under the first rule: u64 has to stay in the band, not to win.
