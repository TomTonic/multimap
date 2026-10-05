# Step 4: the multi-key page for entries with one value (MKSV), design note

Written 2026-10-05, for the user's approval before code (PLAN.md working rules). Words as in [GLOSSARY.md](GLOSSARY.md); the glossary entries that this note changes are listed at the end.
The prediction comes from a model of the real data (`bench/cmd/skmodel -multi`, new; it reproduces the measured memory of the single-key tree to 1 byte: 66.5 against 67 for `street` with
`uint64`, 52.6 against 53 for `uint64` keys with `*T`).

## What we know (the input, not decisions)

1. **The memory of one object per entry is the problem.** The tree of step 3.5 takes 67 bytes an entry for `street` with one value (`btree-map`: 55), 53 for `uint64` keys with `*T` (45). The tree before 3.5 with
   multi-key pages of layout A took 40 and 63 ([step3-tree-results.md](step3-tree-results.md)); its range scans were 3 to 5 times the single-key tree's.
2. **On the real mix the multi-key page needs help.** With one entry in five holding several values, a multi-key page "that holds single-value entries only" was taken away by the fall back everywhere
   ([step3-tree-pages.md](step3-tree-pages.md): 0.89 to 1.06 of the single-key tree, same memory). So step 4 is about one value per entry; the real mix must stay where step 3.5 left it (gate 4), and
   what multi-value entries in pages can bring (-36 to -40 %) is step 5.
3. **Layout B (the user's sketch, `lpage`) is 16 to 22 % smaller than layout A, and slower for short keys** (cold lookups of `u64` 1.7 times A in the microbenchmark; absent keys 2.3 times: it walks all lengths).
   In the tree with strings it was 0.5 to 0.6 of the old tree for small maps (4,096 entries) and 1.3 to 2.5 times for ranges at 212,449; against `btree-map` ranges 1.03 to 1.55 (`street`), 0.82 to 1.27 (`dirs`), churn 0.77 to 1.2,
   `build` 0.57 to 0.77 (untuned code).
4. **A range node costs 11 to 16 % of the objects** (R8: 128 bytes for 2.6 to 3.1 ranges, 3 to 14 bytes an entry, step3-tree-pages.md).
5. **The page of the single-key tree already has what a multi-key page needs:** the grid 32 to 512, the head, the hysteresis, `BackFits`, and for pointers the 166 typed objects (`ptrObject[T, [J]uint64, [N]T]`:
   class and number of slots; 3 + 7 + 15 + 31 + 47 + 63 = 166 is every class with 1 to class/8 - 1 pointer slots at the end).

## Decision proposals

**D1. No range nodes: pages hang below byte nodes and burst.** A subtree whose entries all have one value and whose content fits 512 bytes is one multi-key page, hung where the single-key page of its
first entry would hang (at the byte that tells its keys from the keys beside it); the bytes its keys share beyond that are the page's common prefix, stored once. Insert into a full page *bursts* it:
the byte node on the next byte in which its keys differ, one child per byte value, each child a page again (one entry: a single-key page). That is the burst trie; the tree has no new node type and the model below has no free
parameter. *Why not range nodes:* the model puts the byte tree at 30.9 bytes an entry for `street` with strings (the measured page tree with range nodes: 30.4) and 50.3 for `dirs` (54.3): nothing to gain.
*What it costs:* keys that spread evenly over the bytes. Random `uint64` keys give pages of 4.3 entries (24.0 bytes an entry; layout A with range nodes took 16): still half of `btree-map`'s 45, but this is where
range nodes would pay, and the question belongs to step 6 (routing), which the user decides after step 5.

**D2. Layout B, one grid.** The page is the sketch's: head, one length byte per entry, the common prefix once, the remainders. One grid for every page, 32 to 512 bytes (a multi-key page of two entries
takes 64: the model gives 31.1 against 34.3 bytes an entry on `street` with strings and 50.4 against 52.4 on `dirs` without the 64-byte class; 32 adds nothing). Types: three ranges, single-key pages, multi-key pages, nodes, so that `isPage` is still one
compare (`objType <= kLastMultiKey`) and `isSingleKey` the one of today.

```
strings:  type | n | cpl | rl1 rl2 ... rln | cp ... cp | r1 ... r1 vl1 v1 ... v1 | r2 ... r2 vl2 v2 ... | ...
fixed T:  type | n | cpl | rl1 rl2 ... rln | cp ... cp | all remainders | padding to a word | v1 v2 ... vn   (the values behind the padding, after the last remainder)
```

`n` entries (at most 255), `cpl` the length of the common prefix, `rl_i` the length of the remainder after it (an entry with the empty remainder, the key that ends at the prefix, is allowed: `n` ends the lists, not a zero).
Entries are in key order. For strings each value follows its remainder (`vl` 1 byte, value up to 255): an insert shifts one block, a lookup that has found the key has its value in the same lines
(the deviation of the prototype from the sketch, step3-layout.md). For fixed-size `T` the values are one array at the end, which is also what the typed objects of `*T` need: **the pointer page is the existing `ptrObject[T, [J]uint64, [N]T]`
with N = n** (head, lengths, prefix and remainders in the J words in front); no new Go type. The head is 3 bytes in this page (the single-key page's 6 are history, step 3.5 note).
*Admissible entry:* one value, remainder after the prefix at most 255 bytes, value at most 254 bytes (as in the single-key page), common prefix at most 255 bytes. Anything else stays a single-key page (value overflow, long remainder, as today).

**D3. One function builds subtrees: `build(entries, path length)`.** It is the model's recursion: all entries single-value and fitting 512 bytes: a page; one entry: a single-key page; else a byte node on the next differing byte with a
`build` per child. It is used by the burst of a full page and by the **promote** (an entry gets a second value: the page's entries are built again with that entry as a single-key page; with the sorted entries at hand
this is one pass). The **fall back** and the crowded ratio are not needed, because a page can never hold a multi-value entry. Pages form where two single-value entries meet: the insert that reaches a single-key page
of a single-value entry with a key that differs makes a page of two (instead of a byte node with two children), when both are admissible.

## Operations and rounds (R5)

| operation | how | rounds in the page |
|---|---|---|
| lookup | descent through byte nodes as today to the page; compare the common prefix; walk the lengths (head and lengths sit in the first 64 bytes up to 30 entries), compare the remainder of an entry whose length matches | 2: head and lengths, then the entries |
| insert a new key | same walk finds the position; shift the tail by one entry in place if the class holds it, else grow into the smallest class that holds one more (a new object), if 512 do not hold it: burst | 2 |
| add a value to an entry | promote (D3) | |
| key outside the common prefix | the byte node on the first differing byte goes above the page, the page's prefix loses its first bytes in place (the sketch of `Skip`); the new key is a single-key page | |
| remove | shift out; shrink with the hysteresis of step 3.5 (values fill half of the smaller class); one entry left: a single-key page; none: the child goes | |
| merge | after a removal, a byte node whose children are single-value pages that fit together in 512 bytes becomes one page: the rule is made in 4.2 from the memory after removing half the keys (the memory run does that) | |
| scan | node by node, the page's entries in order; strings copied out as for the single-key page | |

For pointers every added entry changes the shape (N + 1 slots) and so the object: a copy of at most 512 bytes for every insert of a new key. The words page (`uint64`) and the string page grow in place within a class.
Whether the pointer page keeps spare slots is for the measurement of 4.3, not for this note.

## Prediction

Memory, bytes an entry (nodes and pages; the model, one value per entry; no value overflow):

| case | single-key tree (model, measured) | multi-key pages (model) | `btree-map` (measured) |
|---|---|---|---|
| `street` -> `uint64`, one value | 66.5, 67 | **26.9** | 55 |
| `dirs` -> `uint64`, one value | 72.2, 77 | **39.5** | 96 |
| `street` -> `string`, one value (bytes included) | 67.2 | **30.9** | (76 fair) |
| `dirs` -> `string` | 80.0 | **50.3** | (123 fair) |
| `uint64` -> `*T` | 52.6, 53 | **24.0** | 45 |
| `street` -> `*T`, one value | 66.5, 67 | 26.9 | 55 |

Tolerance: 10 % (the model's SK numbers are within 1 byte of the measurements; the pages' slack after real inserts and deletes is the unknown). The natural mix: **at most** 58.0 against 74.7 (`street`, `uint64`) and 78.9
against 88.5 (`dirs`) in the model where pages re-form among single-value neighbours; that is an upper bound (-22 %, -11 %), the build order decides, and step 3.5 measured the real mix's tree = the single-key tree. **Gate 4 asks
only for no cell of the real mix below step 3.5 beyond noise.**

Speed (against `btree-map`, one value per entry; from the page tree with range nodes of step3-tree-pages.md, so a hope and not a model): point lookups 1.0 to 1.4 times, ranges 1.0 to 1.5 (`street`) and 0.8 to 1.3 (`dirs`, small sizes the weakest), `churn` and
`build` 0.8 to 1.2: **credo 1 for `churn` and `build` and credo 2 for the ranges of `dirs` at 4,096 entries are at risk** and will be reported cell by cell. The real mix: `build` 0.85 to 1.0 of step 3.5 (a page is made and, at the entry's second
value, burst again: one entry in five), everything else 0.95 to 1.05. GC: pages of `uint64` and strings hold no pointer; the scannable bytes of the single-value map fall from 34 to about 4 an entry.

## Steps

- **4.0** Glossary and vocabulary (below), the types of pages in `internal/art/node.go`.
- **4.1** `internal/mkpage` standalone for strings and for `uint64` (Go 1.27 generic methods as `skpage.Fixed`): 100 %, fuzz, microbenchmark on the entries of `street` and `dirs` against the model. **Stop for the user.**
- **4.2** In the tree for `string -> {string}` and `string -> {uint64}`: `build`, burst, promote, removal and merge, scans; measure single-value and real mix against step 3.5 and `btree-map`. **Stop.**
- **4.3** Pointer pages (`*T`, `uint64` -> `*T`). **Stop.**
- **4.4** Gate 4 on the PC (M1 when available), the report against this prediction.

## Glossary changes this note proposes (made after approval)

`range node`: kept as the name of an object the tree may get in step 6, no longer part of the tree. `burst` becomes "a full multi-key page is replaced by a byte node with a page or single-key page for each byte value". `split`, `merge`:
`merge` stays (two thin pages), `split` goes. `promote`: "the entry that gets a second value leaves its multi-key page: the page is built again". `fall back` goes. `page` grows a size class 32 and 64 for the multi-key page.

## For the user to decide

1. D1: bursting under byte nodes, no range nodes now (the alternative: bring back `rnode.go` from the tag `before-mk-pages-removal`).
2. D2: layout B with the head of 3 bytes for this page, one grid 32 to 512.
3. D3: the one `build` function for burst and promote; the fall back goes.
4. The order: strings first (4.2), then `uint64`, pointers last (4.3).
