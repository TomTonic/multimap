# Step 5: entries with several values in the multi-key page (MKMV), design note

Written 2026-10-06, for the user's approval before code (PLAN.md working rules). Words as in [GLOSSARY.md](GLOSSARY.md); the glossary entries that this note changes are listed at the end.
The prediction comes from the model of step 4 extended for this step (`bench/cmd/skmodel -multi -mkmv`), which reproduces the measured memory of the trees of steps 3.5 and 4.2 within 6 %
(73.1 and 56.4 modelled against 73 measured for the real mix of `street`, `uint64` values; the value overflows included at 16.3 bytes a value).

## What we know (the input, not decisions)

1. **The page that holds single-value entries only decays on the real mix** ([step4-probe.md](step4-probe.md)). `street` `natural`, 65,536 keys: 58.3 bytes a key fresh, 73.2 after one cycle of the benchmark's own stream with the
   same content (+25 %), 15.9 % of the keys in pages instead of 41.6 %. A key that gets a second value leaves its page (promote: 32 per 1000 operations, the page is built again); when the value goes the key stays a single-key page.
   Putting the key back (demote) restores 63 bytes and costs 37 % of the time. **What does not dissolve does not need to be put back.**
2. **Most keys have few values.** Of the keys that have values, 79 % (`street`), 62 % (`dirs`), 50 % (`uint64` keys of the bench) have one, and 94 %, 89 %, 85 % have at most four. The keys with more than 512 bytes of content
   (0.36 % of `street`, 0.5 % of `dirs`) hold half of the values and, as value overflows with their sets, 15 bytes of every key's 53 or 73 in the prediction below. That part is not in this step.
3. **A page that holds a key with several values has been built once** (the parked branch `mkmv-experiment`, `Map.Pairs`, [step3-tree-pages.md](step3-tree-pages.md)): a further value is an entry whose remainder length is 255 and which stores no key bytes.
   Measured there: -36 to -40 % memory on the real mix, the scanned bytes to 40 %; mutation was untuned (0.4 to 0.9 of the single-key tree).
4. **The scan is not limited by the page** (`BenchmarkFixedScan`, hot cache, `uint64` values, pages of 7 and 20 entries): `Each` with a function per entry 2.3 and 2.0 ns a value, a loop over the same arrays 1.0 and 0.84. The range scan of the
   tree costs 8.2 ns a value: the rest is the glue (`scanPage`: two closures, the bound comparisons, the key buffer, the walk over nodes). So **the format of the page does not bound option B** of the step-4.2 report; B can be done
   after this step on whichever format stands.
5. **State of gate 4 on m43** (PC and M1, [step4-probe.md](step4-probe.md)): memory as predicted; real mix credo 1 and 2 met; `single-value`: `churn`/`build` 0.63 to 0.98 of `btree-map` at 4,096 and 16,384 keys with `string` keys,
   ranges 0.28 to 0.71. Step 5 does not change what happens with one value per entry (D1 below), so these numbers are the reference for "no regression".

## Decision proposals

**D1. A further value is an entry of its own in the length list: the byte 255, no remainder.** The page of step 4 is the special case without any such entry: **a page of single-value entries is the same bytes as today** (no cost for the
case step 4 made fast); a key with k values costs k-1 length bytes more and the k-1 values. The model, bytes a key (value overflows included, fresh tree, `uint64` values / `string` values):

| layout of the entries | `street` | `dirs` | one value only, `street` / `dirs` |
|---|--:|--:|--:|
| single-key pages only (step 3.5) | 88.0 / 115.5 | 101.6 / 152.9 | 66.5 / 72.2 |
| one value per entry (step 4) | 71.4 / 100.3 | 92.0 / 144.1 | 26.9 / 39.5 |
| **length byte 255 for a further value (proposed)** | **52.9 / 82.8** | **72.6 / 125.6** | **26.9 / 39.5** |
| a count byte per entry (the values after the key) | 53.2 / 83.2 | 72.4 / 125.6 | 28.4 / 40.7 |
| the key repeated for a further value | 57.3 / 86.8 | 83.3 / 133.5 | 26.9 / 39.5 |

The count byte costs a byte for every entry (+5.6 % with one value per entry, +17 % for the `uint64` keys' 24.0 bytes), the repeated key costs the remainders of the further values (+8 to +15 %). The proposal has neither.
`MaxRemainder` goes from 255 to 254 (a remainder of 255 bytes stays a single-key page); `n`, the number of length bytes, is at most 255 as before, so a page has at most 255 values.

```
head (3 bytes, as the single-key page's first two, user 2026-10-06): type (bit 0 = bit 8 of the next byte) | cpl (low 8 bits: the length of the common prefix, 9 bits in all) | n (number of length bytes)
strings:  head | cp ... cp | rl1 rl2 ... rln | r1 vl1 v1 | vl v (a further value of entry 1: rl = 255) | r2 vl2 v2 | ...
fixed T:  head | cp ... cp | rl1 rl2 ... rln | the remainders of the entries with rl != 255 | padding to a word | v1 v2 ... vn  (one slot a length byte, the entry's values side by side)
```

**The head is the single-key page's first two bytes in meaning and place:** `type` with its lowest bit as bit 8, then the nine-bit length of the stored key part (`klen`) - for the single-key page its remainder, for this page the common prefix. Code that
compares the stored key part with the search key reads them the same way for both, while it still looks at the type. `n` moves from byte 1 to byte 2 (it is the second byte of the single-key page's count). The common prefix now sits directly behind the head,
the length list behind it (at `3 + cpl`); nothing else changes in size. Whether the stored part also sits at the same offset (3 in this page, 6 in the single-key page) is the question of step 5.5.

The values of a key sit next to each other, in the order they came in; inserting a value appends it to the key's run. Lookup of a key finds its first entry as today (the walk skips the 255 entries) and its values are the entry and the 255
entries that follow. The type bytes, the size classes and the pointer pages (`ptrObject[T, [J]uint64, [N]T]` with N = n) stay: **no new object type, no new node**. Removing a value takes out its length byte and slot; the first value of a key with
more values lets the next one take its place.

**D2. No limit on the values of an entry in a page: it lives in the page if it fits with the others.** The model with a limit (street, `uint64` values; `string` values in brackets): 2 values 64.7 (93.9), 4 values 59.4 (88.9), 8 values 56.1 (85.6),
16 values 54.2 (83.6), no limit 52.9 (82.8). A limit saves nothing in code (the rule "fits 512 bytes" is there anyway) and costs up to 12 bytes a key; if the measurement shows that a key with 30 values makes the page too heavy to change, the limit is a knob
to add then, with the model's number next to it.

**D3. The promote goes; the burst, `pair` and the merge learn the new entry.** `reach` has no `Differs` case: a second value is an `Added` of one length byte and one slot (or the `Full` that bursts the page). `pair` makes a page of a single-key page with
any number of values and a new key when both fit; `build` and `pageOf` take entries with several values (`item.multi`); `mergeFits` and `tryMerge` take children that are pages with several values. `mergeBelow` counts keys, not slots. What stays as
today: the merge trigger of step 4 (a removal that leaves a page nearly empty), the burst, `Widen`, `Skip`, `Prepend`, `abovePage`. A key whose values do not fit a page stays a single-key page or value overflow, and does not come back into a page when
its values go (a residual drift, only for the heavy 0.4 to 0.5 % of the keys; the probe's census after a cycle will show it).

**D4. The pointer pages (4.3 of step 4) are built with this format, once.** 4.3 is still open from step 4; building the pointer page for single-value entries first and changing it here would build it twice. Proposal: 4.3 becomes the last part
of this step (5.3), for `*T` and `uint64 -> *T` in the format of D1 (the existing `ptrObject` family, N = n; every added value is a new shape: a copy of at most 512 bytes, as for a new key today).

## Operations and rounds (R5)

| operation | how | rounds in the page |
|---|---|---|
| lookup (point) | as today to the page; compare the prefix; walk the length list (the 255s are skipped), compare the remainder of an entry whose length matches; the values are the entry's slot and the 255 slots behind it | 2: head and lengths, then the entries |
| add a value to a key in a page | the key's run is found as in a lookup; the value is looked for in the run (Present); else one length byte and one slot are inserted behind the run, in place if the class holds it, else the next class, if 512 do not hold it: burst | 2 |
| insert a new key | as today | 2 |
| remove a value | the run is found; the last value of the key: the entry goes (as today); else its length byte and slot go; the first value of the entry: the next value takes its place | 2 |
| scan | node by node; the page's entries in order, each value of a run on its own: the key is built once for the run | |
| merge, burst, promote | burst and merge as in step 4; **promote goes** | |

## Prediction

**Memory**, bytes a key, fresh tree, model, tolerance 10 % (heap of the measurement: nodes, pages, value overflows; the model's value overflows at 16.3 B a value for `uint64`, the estimate of `skmodel` for strings):

| case | measured now (m43) | step 3.5 | **step 5 (model)** | `btree-sets` |
|---|--:|--:|--:|--:|
| `street` `uint64`, real | 73 | 90 | **52.9** (-27 %) | 285 |
| `dirs` `uint64`, real | 96 | 108 | **72.6** (-24 %) | 336 |
| `street` strings, real | 98 | 114 | **82.8** (-15 %) | 367 |
| `dirs` strings, real | 153 | 165 | **125.6** (-18 %) | 422 |
| one value per entry (five cases) | 27, 41, 31, 51, 24 | | **the same** | |

(The model's value overflow bytes are the weakest part of the real-mix rows; 15 of the 53 bytes of `street` are value overflows.) **After the benchmark's stream (a cycle of `churn`): at most 1.05 times the fresh tree**, today 1.25 times; the keys in pages after a cycle: at least 90 % of those in the fresh
tree (today 38 %). GC: the scannable bytes of the real mix fall by the pages' share (today 21 of 73 for `street`).

**Speed.** The format is the same for one value per entry; so **every `single-value` cell of m43 stays where it is, within the noise of the run (ratio to m43 0.95 to 1.05)**, and a point lookup in a page that has no further values costs the same
(`BenchmarkFixed`/`BenchmarkPage`: Get within 5 %). The real mix: `churn` and `build` 0.95 to 1.0 of the time of m43 (the 32 rebuilds in 1000 operations go; an insert into a key with one value was 322 ns against 255 for a new key, 18 % of the operations: -4 %),
ranges and `prefix` 1.0 to 1.3 times m43 (the tree is a quarter smaller and has fewer objects to scan), point lookups 0.95 to 1.1. Page level: **`Each` of a page with further values at most 2.5 ns a value (hot, as measured now: 2.3), the loop without a function per entry 1.0**. The burst:
at most 5 in 1000 operations on the real mix (0.1 now; pages fill with values), counted by the probe.

## Steps

- **5.0** Glossary (below), the length byte 255 and `MaxRemainder` 254 in `internal/mkpage`.
- **5.1** `internal/mkpage` for strings and `uint64`: further values (insert into the run, remove, Get of the run, Each with the run), tests with the model of pairs, fuzz, microbenchmark (`BenchmarkFixedScan` and the two Get benchmarks against the figures above). **Stop for the user.**
- **5.2** In the tree for strings and `uint64`: promote gone, `pair`/`build`/`pageOf`/merge with several values, scans; probe (census after a cycle, events) and measurement against m43 (single-value must not move; the real mix as predicted). **Stop.**
- **5.3** Pointer pages (`*T`, `uint64 -> *T`), the old 4.3, in this format. **Stop.**
- **5.4** Gate 5 on the PC and the M1 (job in the queue), the report against this prediction.
- **5.5** The shared head (user, 2026-10-06): the single-key page was meant to have the 3-byte head too; it has the leaf's 6 bytes (`type | klen | n (2) | kl (2)`) since step 3.3 only for the code that holds a key from its `base` and compares its *end* (no `Skip` when a node comes in above). That code is the single-key page itself (76 places in `internal/art`, and the value overflow shares the head), not old code that has gone. Once MKMV stands, the single-key page and the value overflow move to `Skip` in place, as the multi-key page does, and the head size for all of them (3, 4, 5 or 6 bytes: the fields are type, nine-bit key length, count; `kl` only if it pays) is found by measurement with a prediction, not decided now.
- Then the decision about option B (scan glue and allocation-free changes), which this step does not touch.

## Glossary changes this note proposes (made after approval)

`entry` (a key with its values) stays; `multi-key page` becomes "a page of entries with one or more values each"; `single-key page` stays "a page of one entry". New: **continuation** (the length byte 255: a further value of the entry before it). `promote` goes with the code;
`fall back` is already gone. `page` grows nothing.

## For the user to decide

1. D1: the length byte 255 for a further value (the alternatives in the table: count byte, repeated key).
2. D2: no limit on the values of an entry in a page.
3. D4: the pointer pages are built in this step (5.3), not before it.
4. Option B (scan glue, allocation-free change) after gate 5, not inside it.

Decided with the user on 2026-10-06: the head of this page has `cpl` as its nine-bit byte 1 and `n` at byte 2 (as above); the shared head of all page kinds is step 5.5.
