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
