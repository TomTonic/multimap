# Step 3, first part: two layouts for the multi-key page

2026-10-03. The question: which layout should the multi-key page have, the one of step 2 (`internal/vpage`,
a directory of tags, 16-byte slots, a heap of tails) or the sketch of the user in
[whataleafneedstostore.md](whataleafneedstostore.md) (MKSV: a header of lengths, a common prefix, the
remainders one after another)? Gate 2 was ticked off as "not met, deficits noted" on the same day;
the deficits the layout can still change are small-map lookups, `build` and `u64`.

This part answers with a model of memory per key, a working prototype of the second layout
(`internal/lpage`, 100% coverage, race and lint clean, fuzzed 30 s) and diagnostic
microbenchmarks against `vpage`. It does not wire the prototype into the tree; that is the next
decision. The names are those of [GLOSSARY.md](GLOSSARY.md); in this document **A** is the page of
`internal/vpage` and **B** the length-header page of `internal/lpage`.

## The two layouts

| | A: `vpage` (step 2) | B: `lpage` (the sketch) |
|---|---|---|
| per entry | head word 8 B + value 8 B in a 16-byte slot, directory of 4 B (1 B in a uniform page), the remainder beyond 8 bytes in a heap | length byte 1 B, remainder, value 8 B, remainder and value side by side |
| search | compare the tag with all tags (SWAR), then head, value and tail | read the lengths, compare the remainders in order |
| entries per page | as many as fit (up to about 29 for `uint64`) | the header's length bytes: 6 at 8 bytes, 14 at 16, 22 at 24, 30 at 32 (half as many if each value has a length byte too) |
| common prefix | only if it saves 48 bytes or more | always, the part all remainders share |
| empty remainder | allowed (an entry whose key ends where the page starts) | not allowed (a length of 0 ends the list), so that entry needs its end page |

Two deviations of the prototype from the sketch, both on purpose:
- each value follows its remainder, where the sketch has all values after all remainders. A lookup
  that has read an entry then has its value in the same line, and an insert shifts one block, not two;
- values have a fixed length of 8 bytes (the map's `uint64`), so the header has no length byte per
  value. The model also shows the sketch's variable-length values, which cost one more header
  byte per entry.

## Memory per key (model)

`bench/cmd/pagefill -layout lens` models B as `-layout vpage` models A: the sorted keys of each corpus
are cut into chunks of 256 keys below a range node, their shared bytes are stripped, the suffixes are
inserted at random into a run of pages, a page that is full is split in the middle, and the router
bytes are added. For B only sizes are modelled, not the bytes; A is the real page. 262,144 keys
(`street`: 212,449). Total bytes per key, pages plus routers; raw tables in
[bench/results-layout/step3-layout/](../../bench/results-layout/step3-layout/pagefill-raw.md).

| kind | A (step 2) | B, values with a length byte | | B, fixed values | | |
|---|--:|--:|--:|--:|--:|--:|
| | | 7 entries (16 B header) | 11 (24 B) | 14 (16 B) | 22 (24 B) | 30 (32 B, class 1024) |
| u64 | 26.5 | 27.2 | 26.5 | 26.2 | 24.9 | 24.8 |
| str | 38.9 | 36.1 | 35.8 | 34.0 | 32.6 | 33.8 |
| uuid | 74.1 | 69.1 | 67.2 | 58.7 | 58.7 | 60.6 |
| email | 55.5 | 51.7 | 50.4 | 48.9 | 47.9 | 48.0 |
| url | 88.4 | 80.6 | 78.4 | 76.5 | 76.3 | 77.3 |
| path | 66.8 | 61.4 | 59.7 | 57.8 | 57.3 | 59.2 |
| street | 35.7 | 30.8 | 29.9 | 29.0 | 27.9 | 28.7 |

With a header of 8 bytes only (3 or 6 entries) B costs 54-90 bytes per key and is out: the
pages are too small, and the router grows with their number (8 bytes per key at 6 entries).

**Reading.** With room for 14 or more fixed-length entries B takes 1% (`u64`) to 21% (`uuid`)
less than A, for every kind of key. With a length byte per value and 7 to 11 entries it is at
most 4% worse for `u64` and 4-10% better for the others. The advantage comes from three places:
no directory (3 to 4 bytes per key), the common prefix always on, and a fuller page (76% instead of
67% for `uuid`, where the class steps fit better). A could take the first two over:
its prefix threshold of 48 bytes is a knob, and the directory is its design.

Caveats: the model does not charge B for rebuilding when the prefix gets shorter (an insert that
differs in the prefix cuts all remainders anew), and A's fill comes from its split rule
(85%) while B splits at the median of a full page.

## Lookups and changes (diagnosis, not a claim)

`internal/lpage/bench_test.go` repeats the benchmarks of `internal/vpage` with the same keys
(`u64` 3M, `uuid` 1M, `path` 1M, in random order, far more than the caches hold).
Windows, native, Ryzen 9 7900, Go 1.27.1, 2026-10-03 15:35-15:39.

Raw logs: [bench/results-layout/step3-layout/](../../bench/results-layout/step3-layout/). ns per
lookup; "B 14" is B with a header of 16 bytes (14 entries), "B 22" with 24 bytes (22 entries).

| lookup | kind | A | B 14 | B 22 | B 14 against A |
|---|---|--:|--:|--:|--:|
| cold, present | u64 | 135 | 228 | 242 | **1.69 x** |
| | uuid | 186 | 194 | 190 | 1.04 x |
| | path | 213 | 165 | 168 | 0.77 x |
| hot, present | u64 | 7.6 | 9.7 | 29.9 | 1.28 x (B 22: 3.9 x) |
| | uuid | 12.5 | 13.9 | 13.9 | 1.11 x |
| | path | 19.2 | 10.4 | 13.7 | 0.54 x |
| cold, absent | u64 | 59 | 142 | 174 | **2.4 x** |
| | uuid | 70 | 163 | 170 | 2.3 x |
| | path | 122 | 113 | 127 | 0.92 x |
| cold, same head, other tail | u64 | 167 | 274 | 275 | 1.64 x |
| | uuid | 189 | 248 | 251 | 1.31 x |
| | path | 295 | 275 | 286 | 0.93 x |

| change | kind | A | B 14 | B 22 |
|---|---|--:|--:|--:|
| insert and delete the same key (page with room) | u64 | 40 | 73 | 73 |
| | uuid | 100 | 76 | 76 |
| | path | 133 | 86 | 86 |
| delete and insert back | u64 | 44 | 83 | 110 |
| | uuid | 123 | 77 | 77 |
| | path | 133 | 72 | 72 |
| fill a page until it splits | u64 | 2134 | 6840 | 13398 |
| | uuid | 1209 | 10157 | 10453 |
| | path | 1845 | 10047 | 10038 |

The same run on WSL (Linux) gave the same picture within 10%. These are single runs of microbenchmarks
on one machine: they say where to look, not how fast the tree is.

**Reading.**
- **Present keys.** B is slower for short keys (`u64`: +69% cold, +28% hot), equal for `uuid`, and faster
  for long keys (`path`: -23% cold, -46% hot). A's hot lookup is a compare of tags that does not grow with the
  key; B compares the remainders in order, and a header with room for 22 entries makes the loop long (30 ns
  for `u64`).
- **Absent keys** take B 2.3-2.4 times as long for short keys: A answers from its directory with a chance
  of 1 - n/256 and reads nothing else; B walks all remainders. The tree asks for absent keys with every
  insert of a new key.
- **Changes in a page with room:** B is faster for long keys (one block shifts, not slots, heap and
  directory), 1.8 times slower for `u64`. **Splitting a full page** is 3-9 times slower in B, but the
  prototype builds the halves from scratch with allocations; A's split is tuned. That says little about
  what B could do.

## What did not help B

- Touching every cache line of the page before the comparison, so that the loads overlap, changed
  `u64` cold lookups from 240 to 229 ns and made `path` worse (166 to 197 ns).
- A loop of fixed length that compares the first 8 bytes of every remainder without branches
  was slower hot (9 to 26-38 ns) and cold (230-250 to 270-320 ns).
- The microbenchmark `first line only` (read byte 0 of a random page) takes 23 ns and `all lines`
  71 ns, so the pages themselves load fast; the `Get` of B takes 250 ns. Where the other 180 ns
  are spent is not understood. Candidates: the keys of the benchmark are fetched from memory in
  both, so they cancel; a branch that depends on a missing line may stop the CPU from running ahead
  to the next lookup. A native measurement on arm64 (the user's M1 offer) can say whether it is
  a property of the layout or of this CPU.

## Conclusion so far

Memory per key and the speed of lookups point in opposite directions, and only the first is
solid:

- **B is smaller than A by 1-21% in the model, for every kind of key.** It is not the prefix policy
  (A with a prefix threshold of 4 bytes instead of 48 gains 0.1-2.4%); it is that B has no directory and a
  fuller page.
- **B is not faster for the short keys that gate 2 missed.** Gate 2's deficits were small maps and `u64`: B's
  lookups for `u64` are 28-69% slower and absent keys 2.4 times, with the prototype's loop. B pays back
  for long keys (`path`, `url` presumably), where A's gap to `btree-map` was.
- **Nothing is known about the tree.** The tree adds routing, scans and the end entries; a layout
  that wins on a microbenchmark has lost to one that loses it before (the first page of step 1).

What I would do, in this order (the user decides):

1. **Do not wire B into the tree yet.** It lacks the operations the tree needs (range scan from a
   bound, merge, a cut at a byte boundary, rebase), and the evidence so far is a model and
   microbenchmarks.
2. **Find out why the cold lookups of B are slow for `u64`**, on the M1 the user offered and on
   this machine with `perf` or the like: 23 ns for the first line, 71 for all lines, 250 for `Get` is
   unexplained. If it is a property of the loop (a branch on missing data), a different loop or a
   tag byte per entry (which makes B the A with a better memory layout) might fix it; if it is a property of
   the layout, B stays for long keys.
3. **Consider one layout per kind of key, as the user suggested for fixed-size values**: A's uniform page
   for keys of at most 8 bytes (`u64`, where A is at 26.5 and B at 24.9-26.2, and the lookup is
   faster), a length-header page for longer keys. That is two page layouts to maintain; the glossary
   already names the second as a candidate.
4. **Do not expect to get B's memory into A by policy.** A lower prefix threshold gains 0.1-2.4%; the rest
   of B's advantage is the missing directory and the fuller pages, that is, the layout.
