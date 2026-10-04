# Step 3: the multi-key page in the tree, for string values

Measured on 2026-10-03/04 (Ryzen 9 7900, Windows; the M1 follows). This is the experiment the user asked
for on the evening of 2026-10-03: hang the multi-key page of layout B (the length-header page,
`internal/lpage`) into the tree for string values and measure the tree against today's on the two real
data sets. The single-key page (SKMV) is not built. Raw results:
[bench/results-layout/step3-tree/](../../bench/results-layout/step3-tree/README.md).

## Short version

1. **The decision "a page holds entries with one value" empties the idea on the real data.** In the natural
   mix (`multi`) of `street` and `dirs` the page tree built under that rule *is* today's tree: speed 0.89 to
   1.06 in all 28 cells (the lowest are two `build` cells, pages built early and rebuilt at the fall back),
   memory the same to the byte. One entry in five with several values (street:
   21 %) is enough for the fall back to take every page away; with the fall back switched off the tree is
   larger, not smaller (111 against 84 bytes a key in `objstat`), because the multi-value entries cut the
   pages into pieces of a few keys.
2. **A page that holds a key with several values (the key once, the values behind it; the option `Map.Pairs`,
   candidate `mv`) gets the idea to the real mix.** Memory per key with the string bytes counted on both
   sides: `street` 147 to 88 (-40 %), `dirs` 217 to 139 (-36 %); the bytes the garbage collector scans fall to
   40 % and 43 %, a GC cycle costs 4 and 2.7 times less. Against today's tree: scans 0.7 to 1.6 times as
   fast, point lookups 0.5 to 1.1, `build` 0.4 to 0.6, `churn` 0.5 to 0.9. Against `btree-sets`, the
   competitor of the credo: range scans 2.4 to 6 times as fast, lookups 1.1 to 1.9, `churn` 0.8 to 1.2,
   `build` 0.7 to 0.9. The B-tree needs 3.5 to 4.5 times the memory of `mv`.
3. **Where all keys have one value (`unique`) the pages are clearly good:** memory per key 77 to 30
   (`street`, -61 %) and 93 to 54 (`dirs`, -42 %), scanned bytes 63 to 4 and 72 to 9, range scans 0.9 to 2.6
   times as fast as today's tree and 0.8 to 1.6 times `btree-map` (today's tree is 0.28 to 0.39 of
   `btree-map` on ranges with strings, see step3-real-data.md). Price: `build` 0.5 to 0.7 of today's, small
   lookups 0.5 to 0.7.
4. **Zero-copy strings** (immutable pages, views of the page instead of copies) make scans another 1.4 to 1.6
   times as fast and lookups 20 % faster, and cost writes another 15 to 25 %: `churn` and `build`.
5. **The rest of the memory is the value overflow:** of the 88 bytes a key of `mv` on `street`, about 39 are the
   arrays and hash sets of the 1.5 to 2 % of keys that have more values than a page holds; those keys hold
   about half of all values. No page layout touches that.
6. **A bug in `internal/art` came out of the dirs memory run** and is fixed (see the end).

## What was built

- `internal/artstr`: a copy of `internal/art` for `Map[string]`. Keys with exactly one value live in
  `lpage` pages, which hang below range nodes the way the pages of step 2 do; keys with several values
  have typed or set leaves as before, or stay in the page (`Pairs`). Differences to `internal/art`:
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
  one memmove per area, not a layout from scratch), and the strings: `EachString` copies the values of a
  scan in one allocation per page run. **The page itself changed too, on the way:** a fourth class of 384
  bytes, headers of up to 64 bytes (default 48: up to 23 entries with values of different lengths; it was
  24 bytes and 11 entries, see the sweep below), value widths of 4, 8 and 16 bytes only, and entries that
  repeat the key before them.
- **Immutable pages and zero-copy strings** (`Map.ZeroCopy`): no page is ever changed in place, every
  change builds a new page, and the strings a lookup or a scan hands out are views of the page's bytes
  (`unsafe.String`), with no copy. A string a caller keeps holds its page, up to 512 bytes, alive. The
  option shows what the copy costs; whether a library may do this is another question.
- **Multi-value entries in the page** (`Map.Pairs`): a page entry is a (key, value) pair; a further
  value of the same key is an entry whose remainder length is 255, which stores no key bytes. The key
  of a multi-value entry is stored once, the values follow in the value area. A full page splits at a
  byte boundary or is laid out anew; a key whose values fit no page gets its leaf (a typed or set leaf).
  The fall back does nothing in this mode (the tree is the same for any ratio from 0 to 4).
- Candidates in the bench (`-tags strvals`, `-vs`): `ordered-lpage` (one value per entry, copy-out strings),
  `ordered-lpage-zc` (zero-copy), `ordered-lpage-mv` (several values per entry), `ordered-lpage-mvzc` (both).
  The environment variables `LPAGE_MAXHEADER` and `LPAGE_MINHEADER` (header sizes, 8 to 64) and
  `ARTSTR_CROWDED` (the ratio at which a subtree falls back; 0: never) change the knobs of a whole run.
- `bench/cmd/objstat -pages [-detail]` prints the object statistic of the page tree (`ARTSTR_PAIRS=1` for
  the several-values mode), with the real names as values.
- Checks: a differential test of the tree against a model (18 workloads in four modes: copy-out and
  zero-copy, one value and several per entry, with headers of 8 to 64 bytes), a page-level model test of
  multi-value pages, strings that stay unchanged while the tree changes (zero-copy), a fuzz test of the tree
  (60 s without a finding), the bench's own verification of the new candidates against `ordered` (points,
  ranges, prefixes, the contents of a build) on four key kinds and both profiles, `-race`, and lint.

`internal/artstr` is an experiment and not held to the project's gates (100 % coverage); it carries the
dead code of the copy (the flat leaves, for example). If a page tree is adopted, the code goes back into
`internal/art`.

## How the numbers are to be read

- Speed: how many times as fast the candidate is as today's `ordered` (above 1: faster), from rtcompare,
  6 to 12 processes a scenario, one process at a time, three sizes (4,096 and 16,384 keys, and the corpus).
  `*`: the interval is wider than asked; `?`: the difference does not clear the noise floor. Tables
  "against `btree-…`" go through the common side `ordered` (the ratio of the two comparisons) and are
  indicative.
- **Memory is not like for like.** The page tree stores the string bytes in its pages; every other
  candidate stores 16-byte string headers that point into one buffer the bench owns (`toVs`). "Fair"
  adds the bytes to the others: `street` values are 10 bytes a key (`unique`) and 28 (`multi`), `dirs` 15
  and 68.
- `build` and `churn` are the ns per insertion of rtcompare's workload streams. The mutation code of the
  page tree is not tuned.

## Results (commit `c4dc4b6`, header 48, four classes)

### One value per key (`unique`)

Speed against today's `ordered`, strings; `lp`: copy-out, `lp-zc`: zero-copy.

`street` (4,096, 16,384 and 212,449 keys):

| operation | lp 4096 | lp 16384 | lp 212449 | lp-zc 4096 | lp-zc 16384 | lp-zc 212449 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.52 | 0.63 | 1.30* | 0.63 | 0.75 | 1.52 |
| valuesBetween | 1.61 | 1.59 | 2.50 | 2.27 | 2.38 | 3.57 |
| prefix | 0.88 | 1.25 | 2.56 | 1.09 | 1.82 | 4.00 |
| churn | 0.66 | 0.81 | 1.00? | 0.49 | 0.62 | 0.87 |
| build | 0.53 | 0.60 | - | 0.40 | 0.45 | - |

`dirs` (4,096, 16,384 and 86,215 keys):

| operation | lp 4096 | lp 16384 | lp 86215 | lp-zc 4096 | lp-zc 16384 | lp-zc 86215 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.63 | 0.73 | 1.08 | 0.75 | 0.85 | 1.23 |
| valuesBetween | 1.10 | 1.15 | 1.37* | 1.69 | 1.85 | 2.00 |
| prefix | 1.03 | 1.18* | 1.43 | 1.59 | 2.04 | 2.56 |
| churn | 0.79 | 0.89 | 0.99*? | 0.63 | 0.74 | 0.87 |
| build | 0.58 | 0.66 | - | 0.50 | 0.54 | - |

Against `btree-map` (the plain competitor for one value per key). `street`:

| operation | lp 4096 | lp 16384 | lp 212449 | lp-zc 4096 | lp-zc 16384 | lp-zc 212449 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.01 | 1.18 | 1.39* | 1.23 | 1.40 | 1.62 |
| valuesBetween | 1.03 | 1.03 | 1.55 | 1.45 | 1.55 | 2.21 |
| prefix | 0.84 | 0.99 | 1.67 | 1.04 | 1.44 | 2.60 |
| churn | 0.77 | 0.90* | 1.20? | 0.57 | 0.69* | 1.04 |
| build | 0.68* | 0.77 | - | 0.52* | 0.58 | - |

`dirs`:

| operation | lp 4096 | lp 16384 | lp 86215 | lp-zc 4096 | lp-zc 16384 | lp-zc 86215 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 0.91 | 1.08 | 1.12 | 1.08 | 1.25 | 1.28 |
| valuesBetween | 0.82 | 0.91 | 1.27** | 1.27 | 1.46 | 1.86* |
| prefix | 0.88 | 1.02** | 1.13* | 1.35 | 1.78* | 2.03* |
| churn | 0.72 | 0.80* | 0.99*?? | 0.58 | 0.66* | 0.87? |
| build | 0.57 | 0.66? | - | 0.49 | 0.53? | - |

Today's tree was 0.28 to 0.39 of `btree-map` on ranges here ([step3-real-data.md](step3-real-data.md)): the
collapse of the scans is gone.

### The natural mix (`multi`)

`lp` is the page tree under the decision (one value per entry), `mv` with several values per entry.
`street`:

| operation | lp 4096 | lp 16384 | lp 212449 | mv 4096 | mv 16384 | mv 212449 | mv-zc 4096 | mv-zc 16384 | mv-zc 212449 |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.01? | 1.00? | 1.00? | 0.49 | 0.57 | 1.08 | 0.59 | 0.68 | 1.18* |
| valuesBetween | 0.98 | 0.99 | 1.00? | 1.05 | 1.12 | 1.56* | 1.49 | 1.67 | 2.17 |
| prefix | 0.99? | 0.98? | 1.00? | 0.76 | 0.93 | 1.54 | 1.01 | 1.39 | 2.22 |
| churn | 1.00? | 1.00? | 1.02*? | 0.50 | 0.58 | 0.88 | 0.37 | 0.47 | 0.79* |
| build | 0.99 | 0.89 | - | 0.44 | 0.52 | - | 0.35 | 0.40 | - |

`dirs`:

| operation | lp 4096 | lp 16384 | lp 86215 | mv 4096 | mv 16384 | mv 86215 | mv-zc 4096 | mv-zc 16384 | mv-zc 86215 |
|---|--:|--:|--:|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.00? | 1.00? | 1.00? | 0.56 | 0.65 | 0.92 | 0.70 | 0.78 | 1.10* |
| valuesBetween | 0.98*? | 0.99? | 1.00*? | 0.73 | 0.85 | 1.25 | 1.18* | 1.37 | 1.79 |
| prefix | 0.97 | 0.99? | 0.94* | 0.67 | 0.83 | 1.06* | 1.11 | 1.49* | 1.79 |
| churn | 0.99? | 1.00 | 1.06* | 0.58 | 0.71 | 0.87* | 0.46 | 0.60 | 0.78 |
| build | 0.96 | 0.93 | - | 0.49 | 0.57 | - | 0.41 | 0.45 | - |

Against `btree-sets`, `street`:

| operation | mv 4096 | mv 16384 | mv 212449 | mv-zc 4096 | mv-zc 16384 | mv-zc 212449 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.27 | 1.43 | 1.86 | 1.52 | 1.72 | 2.04* |
| valuesBetween | 3.19 | 3.13 | 4.69* | 4.52 | 4.65 | 6.52 |
| prefix | 2.37 | 3.18 | 4.86 | 3.16 | 4.76 | 7.02 |
| churn | 0.88 | 1.05 | 1.19 | 0.66 | 0.85 | 1.07* |
| build | 0.78 | 0.93 | - | 0.62 | 0.71 | - |

`dirs`:

| operation | mv 4096 | mv 16384 | mv 86215 | mv-zc 4096 | mv-zc 16384 | mv-zc 86215 |
|---|--:|--:|--:|--:|--:|--:|
| valuesFor | 1.07 | 1.21 | 1.45 | 1.35 | 1.46 | 1.74* |
| valuesBetween | 2.79 | 2.70 | 4.10 | 4.49* | 4.33 | 5.86 |
| prefix | 3.37 | 3.20 | 6.19* | 5.62 | 5.73* | 10.39 |
| churn | 0.83 | 1.02 | 1.12* | 0.66 | 0.87 | 1.01 |
| build | 0.71 | 0.85 | - | 0.60 | 0.67 | - |

### Memory

| data | candidate | heap B/key as measured | heap B/key, fair | scannable B/key | GC CPU per cycle | heap after removing half |
|---|---|--:|--:|--:|--:|--:|
| street, one value | today | 67 | 77 | 63 | +39 ms | 35 |
| street, one value | btree-map | 66 | 76 | 49 | +31 ms | 34 |
| street, one value | page tree | 30 | 30 | 4 | +1 ms | 20 |
| dirs, one value | today | 78 | 93 | 72 | +19 ms | 41 |
| dirs, one value | btree-map | 108 | 123 | 49 | +14 ms | 55 |
| dirs, one value | page tree | 54 | 54 | 9 | +3 ms | 36 |
| street, natural mix | today | 119 | 147 | 111 | +56 ms | 61 |
| street, natural mix | btree-sets | 367 | 395 | 336 | +92 ms | 185 |
| street, natural mix | page tree, one value per entry (no page is left) | 119 | 147 | 111 | +55 ms | 61 |
| street, natural mix | page tree, several values per entry | 88 | 88 | 44 | +14 ms | 50 |
| dirs, natural mix | today | 149 | 217 | 138 | +27 ms | 75 |
| dirs, natural mix | btree-sets | 421 | 489 | 352 | +36 ms | 210 |
| dirs, natural mix | page tree, one value per entry (no page is left) | 150 | 218 | 138 | +27 ms | 75 |
| dirs, natural mix | page tree, several values per entry | 139 | 139 | 60 | +10 ms | 80 |

### The header size (`mv`, `street`, natural mix)

The largest header in bytes limits the entries of a page (3, 7, 11, 15, 19, 23, 27, 31 with values of
different lengths). More entries a page: fewer pages and range nodes, less memory, faster scans, and a
little slower small lookups. The first two rows are from before the 384-byte class.

| largest header | heap B/key | scannable B/key | valuesFor 16K / 212K | valuesBetween 16K / 212K | prefix 16K / 212K | churn 16K / 212K | build 16K |
|---|--:|--:|--:|--:|--:|--:|--:|
| 16 (3 classes) | 118 | 64 | 0.62 / 0.98 | 0.87 / 1.23 | 0.78 / 1.20* | 0.58 / 0.78 | 0.44 |
| 24 (3 classes) | 105 | 55 | 0.61 / 0.99 | 0.97* / 1.27* | 0.85 / 1.23* | 0.52 / 0.78 | 0.45 |
| 32 | 97 | 50 | 0.59 / 1.01*? | 1.01 / 1.35* | 0.88 / 1.28 | 0.54 / 0.82 | 0.47 |
| 48 (default) | 88 | 44 | 0.56 / 1.05* | 1.09 / 1.54* | 0.94 / 1.47 | 0.56 / 0.85 | 0.50 |
| 64 | 83 | 41 | 0.54 / 1.03* | 1.12 / 1.67 | 0.95 / 1.56 | 0.56 / 0.86 | 0.49 |

### The shape of the page tree (`objstat -pages -detail`, corpus size, header 48)

Bytes a key by kind of object, strings included for the pages; "used" is the share of a page's bytes that
hold something. What it does **not** count is the arrays and hash sets of the set leaves.

| | street unique | street multi (mv) | dirs unique | dirs multi (mv) |
|---|--:|--:|--:|--:|
| pages 128 / 256 / 384 / 512 (bytes a key) | 2.4 / 4.8 / 10.5 / 8.7 | 8.9 / 11.0 / 14.2 / 4.7 | 5.4 / 8.0 / 14.3 / 17.3 | 9.4 / 14.5 / 21.5 / 22.4 |
| keys a page | 2.8 / 9.9 / 15.2 / 20.2 | 2.2 / 6.8 / 10.6 / 16.5 | 1.8 / 5.8 / 9.0 / 11.9 | 1.5 / 3.7 / 5.7 / 7.5 |
| used | 46 / 79 / 83 / 85 % | 43 / 76 / 81 / 83 % | 47 / 76 / 84 / 88 % | 50 / 75 / 83 / 87 % |
| range nodes R8 (bytes a key; ranges each) | 3.3; 3.1 | 8.1; 2.9 | 7.7; 2.7 | 14.2; 2.6 |
| leaves (set + typed, bytes a key) | 0.0 | 1.2 | 0.1 | 3.1 |
| total bytes a key | 30.4 | 49.2 | 54.3 | 87.9 |

The measured heap of `mv` on `street` is 88 bytes a key against 49 here: the other **39 bytes, nearly half, are the
value sets** of the set leaves, about 1.5 to 2 % of the keys that have more values than a page holds, each value 16
bytes in an array or a hash set (today's tree has them too).

| values of a key | street keys | street values | dirs keys | dirs values |
|---|--:|--:|--:|--:|
| 1 | 79.3 % | 29.1 % | 62.1 % | 18.0 % |
| 2 to 4 | 15.3 % | 13.8 % | 27.2 % | 19.9 % |
| 5 to 15 | 4.0 % | 11.3 % | 8.1 % | 18.0 % |
| 16 to 64 | 1.1 % | 12.3 % | 2.0 % | 17.0 % |
| 65 and more | 0.35 % | 33.5 % | 0.5 % | 27.0 % |

(212,449 street names with 579,000 values, 86,215 directories with 296,800.) A page of up to 23 entries
holds the keys with up to a handful of values, and with them half of the values (54 % on `street`, 56 % on
`dirs`); the other half belongs to 1.5 and 2.5 % of the keys, which need the value overflow.

Two more things stand out. The pages are 43 to 88 % full (a B-tree's 69 % is the expectation for random
inserts), and **an R8 node, 128 bytes for 2.6 to 3.1 ranges, costs 3 to 14 bytes a key, 11 to 16 % of the
objects**: the routing layer is as heavy as the pages' slack. A range node for few children (64 bytes, the start
bytes in an array instead of a bitmap) is what step 4 of the plan is about; it would save about 5 to 8 %.

## What follows from it

1. **The rule "no page holds entries with different numbers of values" cannot be kept for the real
   data.** The unique profile shows what pages can do; the natural profile shows that one entry in five
   with several values is enough to take the pages away, with or without the fall back. Either an entry with
   several values lives inside a page (built here as an option), or the single-key page and the multi-key
   page must be able to share a range without cutting each other to pieces, for which nobody has a layout.
2. **SKMV as sketched would hold only what does not fit a page.** In the `mv` mode a key with as many values
   as fit a page (up to 22 at a header of 48 bytes, fewer if the values are long) stays in the page; the
   1.5 to 2 % of the keys with more are the part left for a single-key page, which then is the **value
   overflow** of the plan, not the bulk of the multi-value entries. It holds nearly half of the memory of
   the tree.
3. **Zero-copy strings are what makes scans fast and costs writes.** The copy of the values is a quarter of a
   scan of 100 keys. Without it the page tree is 0.9 to 2.6 times as fast as today's on scans of single
   values; with it 1.1 to 4.0. But it needs pages that are never changed in place (every insertion and deletion
   builds a page: `churn` and `build` fall another 15 to 25 %) and lets a string pin up to 512 bytes. A middle
   way would be a scan that hands out views and a lookup that copies; the API of a library that promises
   strings that stay valid does not have it.
4. **Mutation is the open wound.** `build` is 0.4 to 0.6 and `churn` 0.5 to 0.9 of today's tree. Against the
   B-tree competitors the page tree is level (`build` 0.7 to 0.9 of `btree-sets` and 0.6 to 0.8 of
   `btree-map`; `churn` 0.8 to 1.2 and 0.7 to 1.2), where today's tree was ahead of them. The page tree
   is not tuned: a page splits every few insertions, a split allocates two pages and splices the range
   node, and `Insert` makes several passes over the header that one pass could make. The profile of the
   several-values mode (50,000 street-like keys): a third in `TryInsert`, a fifth in splits, a tenth reading
   the first byte of every child (also in today's tree).
5. **Small trees are where the pages lose.** Lookups at 4,096 keys are 0.5 to 0.6 of today's, because the
   search in a page (a linear scan of up to 23 lengths and the first bytes) costs more than a leaf per key
   that sits in the cache; from 200,000 keys on the cache misses decide and the pages win. The tag directory
   of the page of step 2 was faster here (16 to 19 % in the page-level comparison), and larger than this one.
6. **The vocabulary:** if multi-value entries live in the pages, "multi-key page" and "single-key page" blur:
   a page holds entries, each with one or more values. GLOSSARY.md listed that as the first heavy place; the
   data says the vocabulary may lose its two layouts.

## A bug found on the way

The memory phase of the benchmark (remove every second key) crashed on `dirs` with a nil pointer in `collapse`.
Cause, in `internal/art` since step 2 and copied into `artstr`: when a range node is left with one page below
it and no term, the tree moves the page up if its keys still fit a page with the node's path in front of them,
and leaves the node standing if they do not (keys of more than 255 bytes). Removing the last key of that page
left a node with no child, which `collapse` treated as a node with a term. Reproduced with two keys of 300 bytes
that share 100 bytes, added and removed one after the other (`TestRemoveLongKeysOneByOne` in `internal/art`);
fixed in both packages (the empty node now disappears). The main branch, which has no range nodes, is
not affected.

## Not done / open

- Everything on the M1: the jobs `l1` to `l4` in the queue (commit `c4dc4b6`), which the user starts in the
  morning.
- Tuning of the mutation, of the search in a page, of the range node for few children.
- Keys larger than a page, values larger than 255 bytes, and the oversized object stay as they were (a
  leaf). The value overflow is not touched.
- 1M keys, the other key kinds and `uint64` values are not measured; the pages are for strings only.
- Memory of the page tree is measured with the string bytes inside; a user who keeps the strings anyway pays
  them twice, and the "fair" column says only that today's tree does not copy them.

## Reproduce

```
cd bench
go build -tags strvals -o lp.exe ./cmd/bench
lp.exe -suite dev -keys street -values multi -sizes 4096,16384,212449 -minprocs 6 -maxprocs 12 \
       -memn 262144 -memrounds 3 -vs ordered-lpage,ordered-lpage-mv,ordered-lpage-mvzc,btree-sets -out out
lp.exe -suite dev -keys street -values unique ...  -vs ordered-lpage,ordered-lpage-zc,btree-map
go run ./cmd/objstat -keys street -values multi -sizes 262144 -pages -detail   # ARTSTR_PAIRS=1 for mv
```
