# The scan without calls on the hot path (design note, 2026-10-07)

Words as in [GLOSSARY.md](GLOSSARY.md). Origin: analysis (a) of the review ([review-2026-10.md](review-2026-10.md)): the ranges of
single-value maps take 9 to 14 ns a value against 4.4 to 7.5 for `btree-map`, and 75 % of the time is `scanPage`, with three calls through
function values a value (the page's `Each` → a closure → `visit` → the iterator's `yield`). The user's direction: no needless calls or
stack frames on the hot path; the code is written out where we know the path is hot, not left to the compiler.

## Step 1: the page's values handed out directly

`scanPage`, for a scan of values (`RangeValues`), no longer goes through `Each` and `visit`: the page has a method that walks its slots
itself and calls `yield` once a value. A page that lies wholly inside the bounds (`lo` and `hi` false) is a plain loop over its values
(`Fixed`: the array of `T`; `Str`: the value bytes, walked with the value lengths); a page on a bound compares the key of each first slot
with the bound, as `visit` does today, inside the same loop. The scan of keys (`Range`) stays as it is in this step.

**Prediction** (`rangeprobe_test.go`, WSL, medians of 9): single-value ranges street 4,096 11.2 → **6 to 7.5 ns a value**, `u64` 8.8 →
5 to 6, dirs 14.2 → 8 to 10, street 65,536 13.5 → 8.5 to 10; against `btree-map` 0.7 to 0.9. The full scan of the read probe 10 to 30 %
faster; lookups and writes unchanged.

## Step 2: the scan as a cursor, without recursion and without the call of `yield`

The recursive `scanRange` / `scanChildren` becomes a cursor with an explicit stack (one frame per node on the way: the node, the next
child, the bound flags, the path length; the depth is bounded by the key length, a fixed array of 32 frames with a slice behind it for
deeper trees). The cursor's `next` has a fast path (the next value of the current page: an index and a bound check) small enough for the
compiler to inline, and a slow path (the next page: up the stack, down the next child) that is a function of its own. The iterator that
`ValuesBetween…Seq` returns is then a loop of a few lines, `for c.next() { if !yield(c.value()) { return } }`, which the compiler inlines
into the caller's `for range`; the loop body of the caller (the `yield`) becomes plain code in it, so a value costs no call at all. This
is checked with `go build -gcflags=-m` before it is measured; if the compiler does not inline it, the note says why and the user decides.
`touchChildren` (the prefetch of the children of a node) is kept, at the descent into a node. The scan of keys (`Range`) uses the same
cursor, which assembles the key only when asked.

**Prediction:** single-value ranges 3 to 5 ns a value at 4,096 keys (**1.0 to 1.4 times `btree-map`**), natural ranges a further 10 to
30 % faster than after step 1; the full scan of a page's values close to the speed of a loop over an array (1 to 2 ns a value for `uint64`
values); lookups and writes unchanged.

Both steps keep the semantics (bounds inclusive and exclusive, prefix scans, early stop) and are tested by the existing tests of the scans
and the reference tests; 100 %, race, lint 0. Each step is measured on its own (range probe, read probe, and the bench on the PC for the
ranges against `btree-map` and `btree-sets`).

## Step 1 result (2026-10-07)

`page.Fixed.ScanValues` and `page.Str.ScanValues` (new, `internal/page/scan.go`): the page walks its slots itself, a plain loop over the
values when it lies wholly inside the bounds, the bound compare of each first slot inside the same loop otherwise; `scanPage` hands a scan
of values to them and keeps `Each` with `visit` only for the scan of keys. New test `TestPageScanValues`; 100 %, race, lint 0.

Range probe (WSL, medians of 9), ns a value before → after, `btree-map` alongside:

| case | before | after | `btree-map` |
|---|--:|--:|--:|
| street single-value 4,096 | 10.95 | **6.70** | 5.25 |
| `u64` single-value 4,096 | 8.80 | **4.39** | 4.43 |
| dirs single-value 4,096 | 13.94 | **9.46** | 5.87 |
| street single-value 65,536 | 13.44 | **9.01** | 7.53 |
| street natural 4,096 | 7.19 | **4.32** | (not the competitor of the natural mix) |

**Prediction met or beaten:** street 6 to 7.5 (6.70), `u64` 5 to 6 (4.39: as fast as `btree-map`), dirs 8 to 10 (9.46), street 65,536
8.5 to 10 (9.01); natural -10 to -30 % (-40 %). Against `btree-map` now 0.62 to 1.01.

## Step 2 result (2026-10-07)

`art.Cursor` (new, `internal/art/cursor.go`) replaces the recursive scan (`scanRange`, `scanChildren`, `scanLeaf`, `scanPage`, `cmpParts`
are gone): a stack of compact frames (24 bytes, 32 in the cursor, a slice behind them for deeper trees), `NextPage` once a page, and the
page's methods `ValuesIn` (a slice of the page's own array for `uint64` and pointers), `AppendStrings` and `AppendKeys` (`internal/page/scan.go`,
which also replace step 1's `ScanValues`); a value overflow hands its set to the loop instead of being copied. The iterators of `Ordered`
(`rangeSeq`, `AllKeysSeq`) are the loop over a page's values: the compiler inlines them, `Cursor.Init` and the caller's loop body into the
caller (checked with `-gcflags=-m` at the bench's call site), and the bounds no longer escape to the heap. New tests `TestScanOfDeepTrees`
(two chains of 40 byte nodes, the frames beyond 32 used twice), `TestRangeStopsInsideAKeyWithManyValues`, `TestPageValuesIn`; 100 %, race,
lint 0.

**A surprise on the way, understood and fixed:** the first version was 10 to 30 % slower than step 1. The profile showed three faults of the
implementation, not of the design: the next child of a node was searched from the start every time (`childFrom`, 12 to 16 %), a page at a
bound compared every slot with both bounds (20 %), and a value overflow's set was copied into the cursor's scratch (with an allocation a
range and 20 % garbage collection). Fixed by an incremental walk over the children, compares only as long as they can decide, and the set
handed to the loop; the frames were made compact (56 → 24 bytes) and filled in place.

Range probe, ns a value (the bench's ranges of 100 keys), before step 1 → step 1 → step 2, `btree-map` alongside:

| case | before | step 1 | step 2 | `btree-map` |
|---|--:|--:|--:|--:|
| street single-value 4,096 | 10.97 | 6.66 | **6.03** | 4.96 |
| `u64` single-value 4,096 | 8.82 | 4.37 | **2.74** | 4.27 |
| dirs single-value 4,096 | 14.24 | 9.55 | **9.89** | 5.72 |
| street single-value 65,536 | 13.44 | 9.03 | **8.48** | 7.38 |
| street natural 4,096 | 7.24 | 4.37 | **3.42** | |
| `u64` natural 4,096 | 4.17 | 3.74 | **2.68** | |
| dirs natural 4,096 | | 5.02 | **4.42** | |

The full scan through the public iterator (`AllValuesSeq`), step 1 → step 2: street single-value 3.29 → 3.15, `u64` single-value 1.93 →
**0.99**, dirs single-value 4.90 → 5.29, `u64` natural 3.12 → 2.20, street natural 65,536 4.41 → 3.55. (The read probe's scan goes through
`art.Map.RangeValues` with a closure, a path that the public API no longer takes; it is 12 to 29 % slower than in step 1, since the per-page
work of the cursor is not paid back by an inlined loop there.)

**Against the prediction** (single-value 3 to 5 ns a value at 4,096 keys, 1.0 to 1.4 times `btree-map`; natural a further 10 to 30 %; the
full scan 1 to 2 ns a value for `uint64`): `u64` met (2.74, **1.56 times `btree-map`**; full scan 0.99); street (6.03, 0.82 of `btree-map`)
and dirs (9.89, 0.58) **missed**; natural met (-12 to -28 % against step 1). **dirs is the one cell slower than step 1** (+3.5 % ranges, +8 %
full scan): its pages are small (4.5 keys a multi-key page, 13 % single-key pages, 7.6 nodes on the way), so the walk's cost a page (the call
of `NextPage`, the frame, the dispatch) is shared by few values. What remains for street and dirs against `btree-map` is that walk and the
bound compares of the pages at the ends of a range; the lever there is fuller pages (the autotune) or a cheaper step from page to page.
