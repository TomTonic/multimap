# Step 3: the multi-key page in the tree, for string values

Measured on 2026-10-03/04. This is the experiment the user asked for on the evening of 2026-10-03: hang
the multi-key page of layout B (the length-header page, `internal/lpage`) into the tree for string
values, and measure the tree against today's on the two real data sets. The single-key page (SKMV) is
not built. All raw results: [bench/results-layout/step3-tree/](../../bench/results-layout/step3-tree/README.md).

## Short version

1. **One value per entry in the page, as decided: it brings nothing on the real data.** In the natural
   `multi` profile of `street` and `dirs` the page tree is today's tree (speed 0.99 to 1.04 in all
   but the `build` of 16K keys, which is 0.88 to 0.90 because pages built early are rebuilt at the fall back;
   memory identical to the byte): the fall back (GLOSSARY) rebuilds every subtree from inner nodes and
   leaves as soon as one entry in five has several values, and `street` has 21 %. No page is left.
   With the fall back switched off the tree is larger (111 instead of 84 bytes a key in `objstat`, 16K
   keys), because the multi-value entries cut the pages into pieces of a few keys.
2. **Where all entries have one value (the unique profile), the page tree is a different animal:**
   the memory per key falls from 67 to 40 bytes on `street` (string bytes inside the page included,
   which the others do not count), the bytes the garbage collector scans from 63 to 8, and range scans
   are 1.1 to 1.7 times as fast as today's tree (1.2 to 2.6 times with zero-copy strings). Price:
   `build` is 1.6 to 2.9 and `churn` 1.1 to 2.5 times slower, point lookups on small trees 1.4 to 1.7
   times slower (and a little faster on the large ones).
3. **A page that holds multi-value entries inside (key once, values behind it) gets part of that to the
   real mix.** Memory per key, string bytes included on both sides: `street` 147 to 105 bytes (-29 %),
   `dirs` 217 to 158 (-27 %); scanned bytes halve, the GC cycle costs 2 to 3 times less. Speed is a
   mixed picture: scans 0.65 to 1.3 times today's (1.0 to 1.7 with zero-copy), lookups 0.5 to 1.0,
   `build` 0.4 to 0.5, `churn` 0.4 to 0.8. This is a deviation from the decision "no page holds entries
   with different numbers of values"; it is built as an option (`Map.Pairs`) and measured, not adopted.
4. **What this does to SKMV:** in this mode only the entries that do not fit a page keep a leaf (about
   4 % of the objects on `street` at 16K keys: more than 8 to 10 values, or too large). That is the part left for a
   single-key page.

## What was built

- `internal/artstr`: a copy of `internal/art` for `Map[string]`. Keys with exactly one value live in
  `lpage` pages, which hang below range nodes the way the pages of step 2 do; keys with several values
  have typed or set leaves as before (or stay in the page, see below). Differences to `internal/art`:
  - A page has no base of its own: it holds its keys from the depth it hangs at. The `pageBase` trick
    of step 2 (a page starting a few bytes early so that every key has a whole head word) is gone, and
    a page that changes its depth is rebuilt (`Skip` when a node is split in above it, `Prepend` when
    a node above it collapses).
  - The first byte of an object is its kind. A page's first byte is its code (class, header size,
    value width), so the pages take 128 kinds and the nodes come after them (`kindMask` 255).
  - `TryInsert` finds and inserts in one pass; a full page splits at a byte boundary exactly as before.
- `internal/lpage` gets what a tree needs: positions (`Find`, `Seek`), `AddValueAt`, `DeleteAt`,
  `SplitAt`, `Merge` (with a quick refusal, since a tree tries it after almost every delete), `Skip`,
  `Prepend`, a growth in place (`insertCopy`: a new object with a longer header or a bigger class and
  one memmove per area, not a layout from scratch), and the strings: `EachString` copies the values of
  a scan in one allocation per page run.
- **Immutable pages and zero-copy strings** (`Map.ZeroCopy`): no page is ever changed in place, every
  change builds a new page, and the strings a lookup or a scan hands out are views of the page's bytes
  (`unsafe.String`), with no copy. A string a caller keeps holds its page, up to 512 bytes, alive. The
  option shows what the copy costs; whether a library may do this is another question.
- **Multi-value entries in the page** (`Map.Pairs`): a page entry is a (key, value) pair; a further
  value of the same key is an entry whose remainder length is 255, which stores no key bytes. The key
  of a multi-value entry is stored once, the values follow in the value area. A full page splits at a
  byte boundary; a key whose values fit no page gets its leaf.
- Candidates in the bench (`-tags strvals`, `-vs`): `ordered-lpage` (copy-out strings, one value per
  entry), `ordered-lpage-zc` (zero-copy), `ordered-lpage-mv` (multi-value entries in pages),
  `ordered-lpage-mvzc` (both). The environment variables `LPAGE_MAXHEADER` (largest header, 8 to 32
  bytes, default 24) and `ARTSTR_CROWDED` (the ratio at which a subtree falls back; 0: never) change the
  knobs of a whole run.
- `bench/cmd/objstat -pages` prints the object statistic of the page tree (`ARTSTR_PAIRS=1` for the
  multi-value mode).
- Checks: a differential test of the tree against a model for 18 workloads in four modes (copy-out and
  zero-copy, one value and several per entry), a page-level model test of multi-value pages, strings
  handed out stay unchanged while the tree changes (zero-copy), and the bench's own verification of
  the new candidates against `ordered` (points, ranges, prefixes, the contents of a build) on four key
  kinds and both profiles.

`internal/artstr` is an experiment and not held to the project's gates (100 % coverage, race, fuzz); it
carries the dead code of the copy (the flat leaves, for example). If a page tree is adopted, the code
goes back into `internal/art`.

## How the numbers are to be read

- Speed: how many times as fast the candidate is as today's `ordered` (above 1: faster), from
  rtcompare, Ryzen 9 7900 on Windows, 6 to 12 processes a scenario, one process at a time. `*`: the
  interval is wider than asked; `?`: the difference does not clear the noise floor.
- **Memory is not like for like.** The page tree stores the string bytes in its pages; every other
  candidate stores 16-byte string headers that point into one buffer the bench owns (`toVs`). The mem
  tables add the bytes to the others when the text says "fair". `street` values are 10 bytes a key
  (`unique`) and 28 (`multi`); `dirs` 15 and 68.
- `build` and `churn` are the ns per insertion of rtcompare's workload streams; the code of the
  page tree is not tuned (see below).

## Results

All tables with every size, and the raw summaries: the README of the results directory. The
essentials, Ryzen, three sizes (small, medium, the corpus):

### One value per key (`unique`), `street` and `dirs`

Speed against today's `ordered`, string values; copy-out strings (`lp`) and zero-copy (`lp-zc`):

| operation | street 4K | 16K | 212K | `lp-zc` 4K | 16K | 212K | dirs 4K | 16K | 86K | `lp-zc` 4K | 16K | 86K |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.58 | 0.68 | 1.18 | 0.72 | 0.83 | 1.32 | 0.66 | 0.75 | 1.01 | 0.79 | 0.88 | 1.10 |
| valuesBetween | 1.35 | 1.39 | 1.59 | 1.92 | 2.08 | 2.22 | 1.12 | 1.15 | 1.20 | 1.72 | 1.92 | 1.85 |
| prefix | 0.92 | 1.15 | 1.69 | 1.18 | 1.72 | 2.56 | 1.05 | 1.16 | 1.30 | 1.64 | 2.08 | 2.33 |
| churn | 0.48 | 0.54 | 0.81 | 0.40 | 0.48 | 0.75 | 0.80 | 0.85 | 0.91 | 0.63 | 0.66 | 0.79 |
| build | 0.40 | 0.47 | - | 0.34 | 0.39 | - | 0.59 | 0.63 | - | 0.50 | 0.51 | - |

Memory per key at the corpus size (heap / scannable / GC CPU per cycle):

| | street today | street page tree | dirs today | dirs page tree |
|---|--:|--:|--:|--:|
| heap B/key (page tree incl. 10 or 15 bytes of string) | 67 | 40 | 78 | 63 |
| scannable B/key | 63 | 8 | 72 | 11 |
| GC CPU per cycle | +41 ms | +4 to +6 ms | +18 ms | +3 ms |
| `btree-map` heap / scannable | 66 / 49 | | 107 / 49 | |

Against `btree-map` (the plain competitor for one value per key, computed through the common `ordered`
side): lookups 1.15 to 1.35 times as fast, ranges 0.82 to 0.90 (`btree-map` is faster), `churn` 0.54 to
0.88, `build` 0.50 to 0.60 on `street`. Today's tree was 0.28 to 0.39 on ranges against `btree-map`
([step3-real-data.md](step3-real-data.md)): the scan collapse is gone, the build is now the weak spot.

### The natural mix (`multi`), one value per entry in the pages

Speed against `ordered` is 0.99 to 1.04 in 26 of the 28 cells (`street` and `dirs`, all operations, three
sizes; nearly all "?", inside the noise); the two others are the `build` of 16K keys, 0.88 and 0.90; the bytes per key are 119 against 119 and 149 against 150, scannable
111 against 111 and 139 against 138. The tree has no page. With the fall back off (`ARTSTR_CROWDED=0`,
shape only, `objstat -pages`): `street` 111 bytes a key against 84.

### The natural mix, multi-value entries in the pages (`Map.Pairs`)

| operation | street 4K | 16K | 212K | `mvzc` 4K | 16K | 212K | dirs 4K | 16K | 86K | `mvzc` 4K | 16K | 86K |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.52 | 0.61 | 0.99 | 0.63 | 0.72 | 1.08 | 0.56 | 0.63 | 0.84 | 0.70 | 0.77 | 0.94 |
| valuesBetween | 0.95 | 0.97 | 1.27 | 1.32 | 1.41 | 1.64 | 0.70 | 0.81 | 0.96 | 1.15 | 1.28 | 1.30 |
| prefix | 0.76 | 0.85 | 1.23 | 1.01 | 1.27 | 1.67 | 0.65 | 0.77 | 0.93 | 1.09 | 1.37 | 1.37 |
| churn | 0.43 | 0.52 | 0.78 | 0.36 | 0.44 | 0.74 | 0.56 | 0.64 | 0.78 | 0.46 | 0.56 | 0.70 |
| build | 0.39 | 0.45 | - | 0.34 | 0.38 | - | 0.47 | 0.53 | - | 0.41 | 0.43 | - |

| per key at the corpus size | street today | street `mv` | dirs today | dirs `mv` |
|---|--:|--:|--:|--:|
| heap B/key as measured | 119 | 105 | 149 | 158 |
| heap B/key, fair (today's plus the string bytes: 28 and 68) | 147 | 105 | 217 | 158 |
| scannable B/key | 111 | 55 | 139 | 72 |
| GC CPU per cycle | +56 ms | +20 ms | +27 ms | +12 ms |

### The shape of the page tree (`objstat -pages -detail`, corpus size, header 32, 384-byte class)

Bytes a key by kind of object, strings included for the pages; "used" is the share of a page's bytes
that hold something.

| | street unique | street multi (pairs) | dirs unique | dirs multi (pairs) |
|---|--:|--:|--:|--:|
| pages 128 / 256 / 384 / 512 (bytes a key) | 5.0 / 14.3 / 9.6 / 0.2 | 15.0 / 21.2 / 5.3 / 0.1 | 5.6 / 9.4 / 14.9 / 15.5 | 12.6 / 21.3 / 20.9 / 12.6 |
| keys a page (128 / 256 / 384 / 512) | 3.2 / 9.5 / 13.5 / 14.6 | 2.3 / 6.6 / 11.4 / 13.8 | 1.9 / 6.1 / 9.0 / 11.2 | 1.6 / 3.7 / 5.7 / 7.3 |
| used | 51 / 76 / 77 / 80 % | 46 / 73 / 75 / 79 % | 48 / 76 / 83 / 87 % | 52 / 74 / 82 / 86 % |
| range nodes R8 (bytes a key; ranges each) | 4.9; 3.1 | 11.4; 2.9 | 8.0; 2.7 | 16.3; 2.6 |
| leaves (set + typed, bytes a key) | 0.1 | 2.1 | 0.1 | 4.4 |
| total bytes a key (not counting the value sets of set leaves) | 34.9 | 56.6 | 55.0 | 91.5 |

The total is what `objstat` sees: it does not count the arrays and hash sets that set leaves allocate for
keys with many values, which the bench's heap figure does (street multi: 97 bytes a key measured, header 32,
against 56.6 here; the difference is those value sets, which today's tree has as well).

Two things stand out. The pages are 46 to 87 % full (a B-tree's 69 % is the expectation for random
inserts), and **an R8 node, which is 128 bytes and holds 2.6 to 3.1 ranges, costs 5 to 16 bytes a key, 14 to
20 % of the tree**: the routing layer is as heavy as the pages' slack. A range node for few children
(a 64-byte class with the start bytes in an array instead of a bitmap) is what step 4 of the plan is about;
it would save about 7 to 10 % of the total.

## What follows from it

1. **The decision "no page holds entries with different numbers of values" cannot be kept for the
   real data.** The unique profile shows what pages can do; the natural profile shows that one
   multi-value entry in five is enough to take the pages away, with or without the fall back. The
   fall back is not the culprit: with it off the tree is larger. Either an entry with several values
   lives inside a page (the key once, the values behind it: built here as an option), or the
   single-key page and the multi-key page must be able to share a range without cutting each other
   to pieces, which nobody has found a layout for.
2. **SKMV as sketched would hold only what does not fit a page.** In the pairs mode about 4 % of the
   objects on `street` are leaves (more than 8 to 10 values, with the header of 24 bytes: 11 entries),
   and they keep the old leaf formats. What SKMV adds in that world is a pointer-free format for the
   oversized entries and the value overflow, not the bulk of the multi-value entries.
3. **Zero-copy strings are what makes scans fast and cost writes.** The copy of the values is 25 to 35 %
   of a scan of 100 keys; without it the page tree is 1.2 to 2.6 times as fast as today's on scans of
   single values, with the copy 0.9 to 1.7, and on the natural mix 1.0 to 1.7 against 0.65 to 1.3. But it needs pages that are never changed in
   place (every insertion and deletion builds a page: `churn` and `build` are 15 to 25 % slower again)
   and lets a string pin up to 512 bytes. A middle way would be a scan that hands out views and a lookup
   that copies; the API of a library that promises strings that stay valid does not have it.
4. **Mutation is the open wound.** `build` is 0.4 to 0.6 and `churn` 0.4 to 0.9 of today's tree, and 0.5
   to 0.6 of `btree-map` on `build`; the credo (robust between the hash map and the B-tree) is in
   danger there, not at the reads. The page tree is not tuned: a page of 11 entries splits every
   fifth insertion (the vpage held 29), a split allocates two pages and splices the range node, and
   `Insert` runs several passes over the header that one pass could do. The profile of the pairs mode
   (50,000 street-like keys): 33 % in `TryInsert`, 19 % in splits, 10 % reading the children's
   first bytes (also in the tree of today). Headers of 32 bytes (15 entries) and of 16 (7) are
   measured in the PC run `lp4-multi-street-h32` and `-h16`.
5. **The vocabulary:** if multi-value entries live in the pages, "multi-key page" and "single-key
   page" blur: a page holds entries, each with one or more values. This is the case GLOSSARY.md
   listed as the first heavy place; the data says the vocabulary may lose its two layouts.

## Not done / open

- Everything on the M1: the jobs `l1` to `l4` in the queue (see below). The M1 was not available at
  night.
- Tuning of mutation (single pass over the header in `Insert`, fewer allocations on split, a larger
  page for the many small keys).
- Keys larger than the page, values larger than 255 bytes, and the oversized object stay as they were
  (a leaf), so does everything below the ranges that the fall back (`settle`) leaves alone.
- The fuzz test of `internal/art` does not run on `artstr`; the differential test is the check.
- 1M keys, the other key kinds, `u64` values: not measured. The pages are for strings only.

## Reproduce

```
cd bench
go build -tags strvals -o lp.exe ./cmd/bench
lp.exe -suite dev -keys street -values unique -sizes 4096,16384,212449 -minprocs 6 -maxprocs 12 \
       -memn 262144 -memrounds 3 -vs ordered-lpage,ordered-lpage-zc -out out
lp.exe -suite dev -keys street -values multi  -sizes 4096,16384,212449 ... -vs ordered-lpage-mv,ordered-lpage-mvzc
```
