# Step 3, first part: two layouts for the multi-key page

*(Note of 2026-10-04: `internal/lpage` has changed since the measurements below, which are from commit
`b680abf`: a 384-byte class, headers up to 64 bytes with 48 as the default (the tables use 24), value
widths of 4, 8 and 16 bytes only, and entries with several values. The numbers below need that commit.)*

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

> **Correction, 2026-10-03 (user).** This comparison is not a fair test of B, and the recommendation
> to drop B that an earlier version of this document gave is withdrawn.
> - B was designed for values of variable length (strings and the like). Everything below measures
>   8-byte scalar values, the case A is built for (its slot has one word for the value), and the
>   "fixed values" variant of B is the scalar specialization the user wants later, not B itself.
> - The keys of the microbenchmarks (`u64`, `uuid`, `path`) are synthetic generators, and the
>   value counts of the model are the synthetic skew of the bench. The only valid data we have is
>   `street`: real names and real locality counts.
> - A has no counterpart for values of variable length, so there is no comparison for B's own case yet.
> A special case for `u64` and other scalars is fine; rejecting B on the strength of `u64` is not.
> Valid so far, `street` only, memory model (bytes per key, pages plus routers): A 35.7; B with a
> length byte per value 30.8 (header 16, 7 entries) and 29.9 (header 24, 11 entries); B with fixed
> values 29.0 and 27.9. The microbenchmarks have no `street` yet.

## On the real data (street and dirs), 2026-10-03 evening

`internal/lpage` now holds values of any length as in the sketch (a length byte per value, all remainders
then all values, 3 to 15 entries per page; pages of values of one width, the scalar specialization, 6 to
30). `bench/cmd/pagebench` compares it with the directory page of `vpage` on the real keys, the first
value of each key (the unique profile) as a number and as the real name, interleaved with rtcompare. Raw
output: [bench/results-layout/step3-pages/](../../bench/results-layout/step3-pages/) (`pc-native.md`: Windows
native; `pc-wsl.md`: WSL). A single process each, so a diagnosis of the layouts in isolation; the page
is looked up directly, the choice of the page is not timed.

**Point lookups, numbers as values** (A = directory of tags, B = length header; Windows native):

| | present keys | absent keys | page bytes per key |
|---|--:|--:|--:|
| `street` (212,449 keys) | A 69.0 ns, B 85.6 ns: **B 1.19 x as long** | A 33.5, B 51.3: **B 1.53 x** | A 35.2, **B 27.3 (-22%)** |
| `dirs` (86,215 keys) | A 62.5, B 73.1: **B 1.18 x** | A 48.3, B 48.2: **level** | A 55.9, **B 46.6 (-17%)** |

WSL gives the same picture (street 1.38 x and 1.59 x, dirs 1.16 x and level).

**What the variable length costs B** (names instead of numbers, both in length-header pages):
- looking a name up instead of a number costs nothing: street 0.95 x as long (faster), dirs 1.01 x;
- giving the name out as a string, which has to be copied out of the page (it must not alias a page
  that changes), costs 7% (street) and 9% (dirs) on top;
- the pages with the names inside take 31.0 (`street`) and 59.9 (`dirs`) bytes per key, string bytes included
  (the pages with numbers hold nothing else; today's tree holds 16 bytes of header per string and the
  bytes elsewhere).

**Reading.** On the real keys A is faster than B for numbers by 16-19% for present keys, and for absent keys
of `street`; B is 17-22% smaller. Both agree with the synthetic keys, so the earlier verdict stands as a
finding about **scalar values**, and only about those. For values of variable length A has no counterpart,
B costs no more per lookup than with numbers, and what is missing is the comparison with what the tree does
for strings today (a leaf per key with a pointer to the string). The unique profile of `street` and `dirs`
with string values is being measured on the PC for that (job script `run-uni.cmd`, results to follow), and
`p1` on the M1 repeats the table above.

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

## The same on the M1 Pro (2026-10-03, run by the user by hand)

Raw files: `m1-vpage.txt`, `m1-lpage.txt` in the results directory (the header size of B is
presumably the 16 bytes the instruction named). ns per lookup; the cache line there is 128 bytes.

| | A | B | B against A |
|---|--:|--:|--:|
| cold, present: u64 / uuid / path | 108 / 199 / 235 | 276 / 235 / 200 | **2.55** / 1.18 / 0.85 |
| hot, present: u64 / uuid / path | 9.5 / 14.0 / 22.3 | 10.7 / 17.0 / 11.1 | 1.13 / 1.21 / 0.50 |
| cold, absent: u64 / uuid / path | 50 / 83 / 165 | 187 / 199 / 153 | **3.7** / 2.4 / 0.93 |
| insert and delete in a page: u64 / uuid / path | 52 / 119 / 164 | 88 / 91 / 109 | 1.7 / 0.76 / 0.67 |
| fill until split: u64 / uuid / path | 1512 / 1081 / 1487 | 2715 / 2946 / 3736 | 1.8 / 2.7 / 2.5 |

The picture is the one of the Ryzen, and for `u64` it is worse: cold lookups of B take 2.5 times
as long, a cold lookup of A about one memory latency (108 ns) and B's more than two. **So the slow
cold lookup is a property of the layout (or of my loop), not of the Ryzen.**

## What did not help B

- Touching every cache line of the page before the comparison, so that the loads overlap, changed
  `u64` cold lookups from 240 to 229 ns and made `path` worse (166 to 197 ns).
- A loop of fixed length that compares the first 8 bytes of every remainder without branches
  was slower hot (9 to 26-38 ns) and cold (230-250 to 270-320 ns).
- The microbenchmark `first line only` (read byte 0 of a random page) takes 23 ns and `all lines`
  71 ns, so the pages themselves load fast; the `Get` of B takes 250 ns. Where the other 180 ns
  are spent is not understood. Candidates: the keys of the benchmark are fetched from memory in
  both, so they cancel; a branch that depends on a missing line may stop the CPU from running ahead
  to the next lookup. The M1 shows the same gap (above), so it is not this CPU.

## Conclusion so far (corrected)

What the numbers support, and nothing more:

- **For scalar values** (8 bytes) A is faster than B for short keys and for absent keys (Ryzen and M1), B is
  smaller everywhere in the model and faster for the long synthetic keys. That is a case for a
  scalar specialization, not a verdict on B.
- **For `street`**, the one real corpus, B is 16-22% smaller in the model (above). Its lookups and changes
  have not been measured on `street`.
- **For values of variable length** nothing has been measured. The prototype stores 8-byte values only.
  A comparison needs a real benchmark: `street` offers one, since its corpus has the names of
  the localities (10,199, variable length) next to the indexes, so a street can map to the real names
  of its localities.

What to do next (the user decides): add values of variable length to the prototype (a length byte per
value, as in the sketch); build the `street` benchmark with the locality names as string values and
run the microbenchmarks and the model on it; compare B with what the tree does today for string values
(typed and set leaves, `strvals`); only then choose.
