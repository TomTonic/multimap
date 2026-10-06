# Status

## 2026-10-06: step 5.2 built: the tree with several values in a page ([results](step5-tree-results.md)); stopped at one missed prediction

- Promote and `Insert`/`Differs` gone; `reach`, `pair`, `build`, merge, `pageRemove`/`RemoveKey` and the scans take entries with several values; `Remove` of the pages says `Removed`/`Gone`/`Absent`, `KeysUpTo` counts keys cheaply. 100 % (race), lint 0; commit `dbbf333`.
- Probe: 88 % of the keys stay in pages after a cycle (17 % before), fresh memory -31 % (street) / -24 % (dirs) in the census, the six `single-value` cases identical to step 4 in every count, burst 2.3 to 4.8 in 1000 operations. **Missed: after one cycle the tree is 1.18 to 1.24 times a fresh one (predicted 1.05)**: the pages are less full (shrink only at half of the smaller class). Not changed; decision for the user. PC and M1 runs not started.

## 2026-10-06: step 5.1 built: the page with several values a key ([results](step5-mkpage-results.md)); option B built too (values of `Fixed` at the end); waiting for the go for 5.2

- `internal/mkpage` (both flavors): the head `type | cpl (9 bits) | n`, the prefix behind it, `Further` (255) for a further value; `Add`, `Remove`, `EachValue`, `Keys`, `Each` with a flag; 100 % (race), fuzz 40 s, lint 0; the tree compiles on it unchanged in behaviour (`maxPageByte`, 255-byte-prefix pages possible now).
- **Surprise:** Get +3 to 12 %, insert/remove +9 to 17 %, `Each` +36 to 45 % against the page of step 4: the sum of the length list that finds the values now skips the 255s. Option B (the values of `Fixed` at the end of the object, which also matches the pointer pages) was built the same day: Get and changes back at the page of step 4 (1.00 to 1.03; one miss: Get of an absent key in 20 entries +21 %), `Each` +15 %, memory unchanged.

## 2026-10-06: step 5 (MKMV) design note written, approved ([step5-mkmv-design.md](step5-mkmv-design.md))

- Proposal: a further value of a key is an entry of the length list with the byte 255 and no remainder (a page of single-value entries stays byte-identical, so `single-value` must not move); no limit on values per entry; the old 4.3 (pointer pages) built in this format as 5.3. Model: `street`/`dirs` real mix 52.9 / 72.6 bytes a key (measured now 73 / 96), no decay after use, promote gone.
- `skmodel -multi -mkmv marker|count|repeat -mkmvcap N` and `BenchmarkFixedScan` (page-level scan 2.3 ns a value, a plain loop 1.0: the scan's 8.2 ns is mostly glue, so option B does not depend on the format).

## 2026-10-05 night: probe of the benchmark's streams ([step4-probe.md](step4-probe.md)); option A chosen by the user, with "profiling or statistics to learn for step 5"

- The step-4.2 hypothesis was wrong for `single-value` (no promote there): the cost was the merge tried after every removal of a key (28 % of the time, 4 % success). Now tried only when the page is nearly empty (`mergeBelow = 2`): a third off the time of `churn`, 3.5 % more memory after a mass removal.
- With several values per key the pages decay: 58 to 73 bytes per key (+25 %) after one cycle, 16 % of the keys in pages instead of 42 %; demoting them again works (63) but costs 37 % of the time. For step 5: entries with several values stay in their page.
- Random `uint64` keys: memory depends on the size, 51 bytes per key at 16,384 against 45 of `btree-map` (24 at 262,144): the one case below credo 1 in memory, and an argument for the question of step 6.
- **PC measurement of the fixed code (`3e1e952`, `m43`, finished 00:29):** `single-value` `churn` 0.03-0.28 → 0.69-1.21 of step 3.5, `build` 0.04-0.19 → 0.66-1.03; real mix against `btree-sets` `churn` 1.00-2.54, `build` 0.99-2.43 (credo 1 met); memory unchanged. Not met: `single-value` at 4,096/16,384 keys against `btree-map` (0.63-0.98), and ranges of `single-value` against `btree-map` (0.30-0.68). M1 job `m43` done: real mix credo 1 and 2 met, `single-value` ranges 0.28-0.71 of `btree-map` (see step4-probe.md).
- Counters (build tag `mkstats`) and the probe (`TestProbe`, `MKPROBE=1`) are in the tree; no cost without the tag.

## 2026-10-05 evening: step 4.2 built and measured on the PC ([results](step4-tree-results.md)); gate 4 not met

- Multi-key pages in the tree for strings and `uint64` (`internal/art/mkkey.go`, `internal/mkpage`): 100 %, race, fuzz, lint. **Memory as predicted** (27, 40, 31, 51, 24 B an entry with one value against 55, 95, 66, 107, 45 of `btree-map`; the real mix -19 % and -11 %).
- **Speed not as predicted:** range scans 1.1 to 2.9 times step 3.5 but 0.3 to 0.7 of `btree-map`; `churn` and `build` with one value per entry 0.03 to 0.28 of step 3.5 (measured on `a166cad`).
  Cause 1 found and fixed after the measurement (merge on every removal, rebuild for a key outside a page's prefix): churn 0.08 to 0.50, build 0.09 to 0.55 on a quick WSL check; the PC has not measured it. Cause 2 (transient values promote pages, which never come back) is the design rule meeting the benchmark's streams: options A to D in the report, for the user.
- Gates after the changes: green (`internal/art` 100 % also under `-race`, fuzz 60 s, lint 0).

## 2026-10-05: step 4 approved (D1 burst, no range nodes; D3 one `build`), 4.1 done: `internal/mkpage` ([results](step4-mkpage-results.md))

- The user approved the design note: bursting under byte nodes (range nodes: question of step 6), `build` for burst and promote, no fall back, order strings, `uint64`, pointers. Glossary updated (burst, promote, merge changed; fall back, split retired; range node out of the tree).
- `internal/mkpage` (`Page` strings, `Fixed` pointer-free `T`): 100 %, race, fuzz, lint. Built into the model, **the pages have exactly the bytes of the model on all 10 data sets**. `Get` 24 ns at 7 entries, 46 ns at 20 (one `compare` loop instead of `bytes.Compare` took 20 to 60 % off).
- Next: 4.2 (the tree for strings and `uint64`); stop for the user before it.

## 2026-10-05: step 4 (MKSV) design note written, waiting for the user's approval ([step4-mksv-design.md](step4-mksv-design.md))

- Proposal: no range nodes, multi-key pages below byte nodes that burst when full (model: `street` strings 30.9 B an entry against 30.4 measured with range nodes, `dirs` 50.3 against 54.3; random `uint64` keys 24.0 against 16 with range nodes: the case for step 6); layout B with a 3-byte head, one grid 32 to 512, pointer pages as the existing `ptrObject` with N = n;
  one `build` function for burst and promote, the fall back goes. Prediction: single-value `street`/`uint64` 66.5 to 26.9 B an entry (`btree-map` 55), `dirs` 72.2 to 39.5, `uint64` -> `*T` 52.6 to 24.0 (45); the real mix at most -22 % / -11 %, gate 4 asks only for no loss.
- New: `bench/cmd/skmodel -multi [-values string|words|pointers] [-mkgrid ...]` (the model, reproduces the single-key tree's measured memory to 1 B).

## 2026-10-05: step 3.5 measured on the PC ([results](step3-tree-results.md)); gate 3 not met in two places

- 18 runs (10:17 to 12:04): `node-layout` and the tree before 3.5 as references. Memory as predicted (`uint64` real 90 and 108 B per entry, one value 67; `*T` real 90 and 108). Real maps: credo 1 met in every cell against `btree-sets` (churn 1.26 at the lowest), speed 0.85 to 0.99 of `node-layout`, no cell below 0.70, 0.93 to 1.06 of the tree before 3.5.
- **Not met:** memory above `node-layout` (equal for `uint64` street; +1 to +5 % for `dirs` and real `*T`, **+13 to +18 % for `*T` with one value per key**, the 32-byte class) and credo 1 for one value per key (ranges 0.19 to 0.57 of `btree-map`). Against the tree before 3.5 the `uint64` single-value maps lose their multi-key pages (memory 67 against 40, scans 0.15 to 0.48): that is step 4.
- Open: the 1 to 4 bytes and 3 to 17 % against `node-layout` are not traced; `hashed` was not in the runs; no M1 data.

## 2026-10-05: step 3.5.5 done, step 3.5 complete: no flat leaf, typed leaf or `vset` left in `internal/art`

- Every map holds an entry in a single-key page (`skpage.Page` for strings, `skpage.Fixed` for `uint64`-like and one-pointer values) or in a value overflow with a `Set3[T]`; the types without a page (an interface, a struct with a string, more than 16 bytes)
  have a value overflow for every entry (`newOverflowLeaf`, `rekeyOverflow`). Gone: `leaf[T, K]`, `newSetLeaf*`, `vals`, `setKeyCap`, the keyArea types, `rekey[T]`. `vset` is used by `Hashed` only. `leafTail` is a constant (31).
  Gates: race, fuzz, lint, **100 % of `internal/art` also under `-race`**.
- Gate 3 of the plan ("flat, typed and set leaves are gone from the code") is met. Next (PLAN 3.5 last step): measure the three cases (`street`, `dirs`, `u64` keys with `*T`) on the PC and as M1 jobs, report against the predictions.

## 2026-10-05: step 3.5.4 done, `*T` maps hold their entries in pointer pages; the typed leaf is gone

- `Map.flat == 1` covers every `T` that `skpage.Supported` takes: small and pointer-free, or one word that is a pointer (the 166 typed objects). Every other `T` (an interface, a struct with a string, more than 16 bytes) is mode -1: each entry a value overflow.
  `typed.go` and its tests are deleted; the reference test, the fuzz test and the object tests run pointer maps. Gates green (race, fuzz, lint, 100 % without `-race`).
- Left of the old code: `leaf[T, K]`, `newSetLeaf*`, `vals`, `vset` inside `internal/art` for mode -1. Next: 3.5.5, mode -1 on the generic value overflow with `Set3[T]`, and those go.

## 2026-10-05: step 3.5.3 done, the tree for pointer-free values (`uint64`) holds `skpage.Fixed` pages

- `Map.flat == 1` is now `Fixed` pages with `Set3[T]` as the value overflow (`internal/art/fixedkey.go`: `newFixedLeaf`, `addFixed`, `removeFixed`, `removeFromFixedOverflow`, `rekeyFixed`, `fixedFromOverflow`); **`flat.go` and the flat leaf are gone**.
  The tests of the flat leaf were ported to the page (`fixedkey_test.go`: `TestFixedKeys`, `TestFixedKeyMovesUp`, `TestFixedRekey`; the value types test and the old overflow's rekey test are in `values_test.go`). Gates: race, 100 % without `-race`, fuzz, lint.
- **Memory as predicted** (`objstat`, object bytes per key, value sets of the value overflow not counted): `street` one value per key **66.6** (model: 66.6), `street` real **74.1** + about 15 for the `Set3` of the value overflow = 89 (model 88.9),
  `dirs` real 92.1 (+ about 15 = 107; model 104). Next: 3.5.4 the pointer pages (`*T`: 166 typed objects) replace the typed leaf.

## 2026-10-05: step 3.5.2 done, the value overflow is generic in `T`

- `valueOverflow[K, T]`, `newValueOverflow[T]`, `overflowSetOf[T]`, `overflowAdd[T]`, `overflowEach[T]` (`internal/art/singlekey.go`): the object, its size and the offset of the pointer do not depend on `T`; strings call it with `T = string`
  and nothing changed for them (all string tests unchanged, race, fuzz, lint green). New test `TestValueOverflowOfOtherValues`: every key area of the grid and a key held as a string, for `uint64` and `*rec`.
- Next: 3.5.3 the tree for `uint64` with `Fixed` pages (the flat leaf and the old set leaf with `vset` go for these maps).

## 2026-10-04 late night: step 3.5.1 done, `skpage.Fixed` standalone ([results](step3-fixed-results.md))

- `Fixed` with 166 generated pointer types, GC test for all of them, fuzz, 100 %; `BackFits` (the way back from the value overflow, also for strings). Against flat and typed leaf on the real entries: hit and read equal, churn +2 to 6 % (`uint64`) and
  +15 to 42 % (`*rec`), build +17 to 22 % (`uint64`); memory equal for `uint64`, **+20 % / +41 % for `*rec`** (32-byte minimum class against the typed leaf's 24). User's answers (2026-10-05): no 24-byte class for now (backlog for the profile after 0.8); the string page got the same shrink hysteresis. Next: step 3.5.2, the value overflow generic in `T`.

## 2026-10-04 night: step 3.5.0 done, the multi-key pages are out of the code

- **Cut out** (user's decision: no code or mental dependency now): the page paths of insert, lookup, scan, delete and rebuild, `page.go`, `rnode.go` (range nodes), `rebuild.go` (build, ranges, fall back),
  `internal/vpage`, `internal/lpage` and the benches `pagebench`, `pagefill`. Every map holds its keys in single-key pages (and the flat/typed leaves that 3.5 replaces) below byte nodes; `find` returns
  the leaf again. **Way back:** branch `mk-pages` and tag `before-mk-pages-removal` (both pushed, the tree of commit 1984f4c with the experiment `mkmv-experiment` beside it). Gates green, `internal/art` 100 %.

## 2026-10-04 night: rename done, step 3.5 design note waiting for three decisions

- **Rename** (622d634, names only): `singleKeyHead`, `isSingleKey`/`isMultiKey`, `kValueOverflow`, `objType`; `skleaf.go` is `singlekey.go`. Flat and typed leaves keep their names until 3.5 deletes them.
- **`s34b` measured** ([step3-overflow-results.md](step3-overflow-results.md)): memory and lookups unchanged, `build`/`churn` 2 to 3 % slower (not investigated).
- **3.5 design note** ([step3-fixed-design.md](step3-fixed-design.md)): one fixed-width layout for `uint64` and `*T`, `Set3[T]` for every `T`, flat/typed/`vset` leave `Ordered`. Waits for the user: (1) multi-key pages
  on or off in 3.5 (recommended: on; off costs 8 % on the real mix and doubles one-value-per-key memory), (2) pointer pages only for one-word `T`, (3) alignment and equality of values.
- Routing layer (step 6): whether it is part of 0.8 is decided at the end of step 5 (user).

## 2026-10-04 evening: step 3.4 measured, set leaf on the grid

- **3.4 measured** ([step3-overflow-results.md](step3-overflow-results.md)): `Set3` in the set leaf. Memory as predicted (`street` natural 114 B/key, `dirs` 165),
  speed neutral on all four profiles (the overflow keys are 0.5 to 1.6 % of the keys), GC the same.
- **User:** the leaf stays on the grid 32/64/128/256 (done: key areas 18, 50, 114, 242); pointer position: left last (a measured 1 ns a leaf and cycle
  does not pay for a second key offset); **new:** nine-bit remainder length through the lowest bit of the kind byte (user's encoding; [step3-klen9-design.md](step3-klen9-design.md)). Done: kinds step by two, set leaf of strings on the grid 32 to 512, page remainder up to 505; all gates green; fixed a crash of `skpage.New` for 512+ bytes of content. Not measured yet.

## 2026-10-04 afternoon: step 3.3 measured, step 3.4 design note written

- **3.3 measured** ([step3-skmv-results.md](step3-skmv-results.md), `bench/results-layout/step3-skmv/`): the natural mix holds credo 1 against `btree-sets`
  in every cell; against `node-layout` the single-key page tree is slower (lookups 0.66 to 0.90, ranges 0.43 to 0.70, `churn` and `build` 0.84 to
  0.92), memory fair -23 % and -25 %, GC work about half. One value per key: credo 1 and 2 not met (ranges 0.30 to 0.48 of `btree-map`). Gate 3 not met in
  the single-value profile; read at the end of step 3. The scan gap is in PLAN.md under "To check later" (the user: leave it, go for 0.8).
- **3.4 design note** ([step3-overflow-design.md](step3-overflow-design.md)) is waiting for the user's approval: a hash set of 512-byte pointer-free blocks for
  the keys whose values do not fit a page; prediction `street` natural about 96 B/key (measured 113), `dirs` about 162.
- Next: after the approval, `internal/strset` standalone, then in the tree, then measure.

## 2026-10-04 midday: single-key page for strings is in the tree, measurement running

- **Decisions of the day** (user): the plan order SKMV, MKSV, MKMV; size classes of the single-key page
  32, 64, 128, 256, 384, 512 (step3-skmv-sizes.md: a page per key at 128 bytes would take 133 B/key for 39
  B of content); design note step3-skmv-design.md approved (layout, copy-out strings, byte nodes plus a page per
  key). Profiles of the bench are called `single-value` and `natural` now (`multi`/`unique` still accepted).
- **Done:** 3.0 (renames after the glossary in `internal/art`, `vpage`, `lpage`: `path` -> common prefix,
  `depth` -> `pathLen`, `term` -> end page, `inner` -> byte node, `settle`/`crowded` -> fall back, `suffix` ->
  remainder; `base` and `class` left on purpose; the bench reads the bytes of string values; build tag `ptrvals`),
  3.2 (`internal/skpage` and `bench/cmd/skbench`), 3.3 code (`Map[string]` holds its keys in single-key pages, the set
  leaf is the value overflow; commit 06244b4; art tests 100 %, race, fuzz 60 s, lint clean). **Deviation:** the page
  has the 6-byte header of the leaves for this step (the 3-byte header is a later step, design note section 3).
- **Running on the PC** (from 11:04, about 4 hours): `sk-street-natural`, `sk-dirs-natural`, `sk-street-single`,
  `sk-dirs-single` (`run-sk.cmd`): the new `ordered` against `baseline` (node-layout), `btree-sets`/`btree-map`,
  `hashed`. The reference runs of before (`ref-str-street-natural` etc.: the tree of step 2) are only done for
  street natural (`ref-str-street-multi`, 1 h 21 min); the others are not run, the new runs go through the
  common side `baseline`.
- **M1 jobs l1 to l4** (the MKMV experiment) are in: `bench/results-layout/step3-tree/m1/`, same picture as the PC.
- Open: the three-byte header, the value overflow (3.4), the other value types (3.5).
- **Decided (user):** the aim is a complete "version 0.8" on the existing benchmarks (steps 3 to 5); then a large
  measurement and profile, then the optimization strategies. Further corpora (inverted index of Wikipedia, DBLP,
  Wikipedia pagelinks, the full Debian tree, DNS records of the Tranco list) are a backlog in PLAN.md and are not
  built now.

## 2026-10-04 morning: review of the night, new order, plan revised

The user had the night's work reviewed (Opus) and set the course:

- **Order:** the single-key page (SKMV) first, with `string -> {string}` as the base case, then
  `string -> {uint64}`, `string -> {*T}`, `uint64 -> {*T}`; then the multi-key page (MKSV) as the special
  case; multi-value entries in multi-key pages (MKMV) last.
- **Decisions** (2026-10-04): in step 3 the tree is byte nodes plus one single-key page per key
  (multi-key pages and the fall back off); string values are stored as bytes and copied out (no
  zero-copy); gate 3 holds credo 1 and the memory hard and reports the rest (a cell below 0.70 needs its
  cause); the MKMV experiment is parked.
- **Explainability** is now a working rule: design note with a prediction before code, no knob
  without a reason, stop at surprises instead of building alternatives ([PLAN.md](PLAN.md)).

What the review found (reported to the user):

- The finding of the night holds and supports the new order: on the natural mix the fall back takes
  every multi-key page away, so the single-key page carries the real multi data.
- The night built MKMV (`Map.Pairs`), which the user had excluded, and tuned constants by sweeps
  (header 16 to 64 bytes, a 384-byte class, `MinHeader`, `ShrinkFill`, `MergeFill`, the fall back ratio).
- `internal/artstr` was a full copy of `internal/art` with dead code (coverage 76.9 %), `internal/lpage`
  is at 98.7 % and its package doc was stale.
- History was rewritten (`filter-branch`, unpushed commits) against the plan's rule; a second binary
  (`bench/cmd/bench/bench`, 17.7 MB, commit `b8e4111`) had been pushed. It is removed in a normal commit
  and ignored; the history keeps it (no rewrite of pushed commits).
- The bench's timed loops read only the header of a string value, which favours candidates that hold
  pointers to the caller's strings (fixed in step 3.0).
- The glossary was not in the code (about 900 old words in `internal/art`) and itself out of date.
  Updated; the renames in the code are step 3.0.

Done this morning: branch `mkmv-experiment` (`2adf119`, pushed) holds the experiment; on `cacheline`
`internal/artstr`, the bench candidates `ordered-lpage*` and `objstat -pages` are removed (the reports and
raw results stay). PLAN.md, GLOSSARY.md, STRATEGY.md (credo 3 loosened for step 3, 384-byte pages in R1,
the order in 4.2) revised. M1: `l1` finished at 08:21; `l2` to `l4` run until about 10:05 (input only).

**Next action:** step 3.0 of PLAN.md (renames, the bench reading values, `ptrvals`, the reference runs),
then the design note 3.1 for the user.

**Housekeeping for the user:** the local branch `backup-before-filter` (the state before the history
rewrite) can be deleted; nothing on it is needed.

## 2026-10-04 night: the multi-key page hangs in the tree (strings), measured; the rule does not hold up on real data

The user asked (2026-10-03, evening) for the multi-key page for strings in the tree, measured against
today's tree, SKMV afterwards. Done on the PC (the M1 follows), report in
[step3-tree-pages.md](step3-tree-pages.md), raw results in `bench/results-layout/step3-tree/`:

- `internal/artstr` (experiment, a copy of `internal/art` for `Map[string]`) with pages of `internal/lpage`;
  candidates `ordered-lpage`, `-zc` (immutable pages, strings that are views), `-mv` (several values per
  key inside the page), `-mvzc` in the bench (`-tags strvals -vs ...`). `lpage` grew: a 384-byte class,
  headers of up to 64 bytes (default 48), entries that repeat the key before them.
- **The natural mix (`multi`) with one value per entry in the pages, as decided: the tree is today's tree**
  (speed 0.89 to 1.06, memory the same). One entry in five with several values is enough for the fall back
  to take every page away.
- **With several values per key inside the page (an option, not the decision):** memory per key with the
  string bytes counted on both sides -40 % (`street`) and -36 % (`dirs`), scanned bytes 40 %, GC cycle 3 to
  4 times cheaper; against `btree-sets` ranges 2.4 to 6 times as fast, lookups 1.1 to 1.9, `churn` 0.8 to
  1.2, `build` 0.7 to 0.9; against today's tree lookups 0.5 to 1.1, `build` 0.4 to 0.6, `churn` 0.5 to 0.9.
- One value per key: memory -61 % and -42 %, ranges 0.9 to 2.6 times today's, `build` 0.5 to 0.7.
- Nearly half the heap of the several-values tree is the value sets of the 1.5 to 2 % of keys with more
  values than a page holds: that is the value overflow, no page layout touches it.
- **A crash in `internal/art` found and fixed** (a range node left with a page that cannot move up; its last
  key removed; since step 2): `TestRemoveLongKeysOneByOne`.
- M1: jobs `l1` to `l4` are in the queue (pushed at `c4dc4b6`), the user runs them in the morning.
- Open for the user: give up the rule "no page holds entries with different numbers of values" (the data
  says it empties the idea on the real mix); then SKMV is only the value overflow. Zero-copy strings: yes or
  no. Then tuning of the mutation, and the range node for few children (step 4).
- Housekeeping to tell the user: seven unpushed local commits were rewritten with `git filter-branch` to
  drop a 5.6 MB test binary (`internal/artstr/artstr.test`) that went into a commit by accident; the local
  branch `backup-before-filter` still holds the old state. `.gitignore` now ignores `*.test`.

## 2026-10-03 late: unique runs and r4 are in

The PC unique runs (both data sets, `uint64` and strings) finished at 21:29, the M1 job r4 at 21:30; both
are in [step3-real-data.md](step3-real-data.md). Main finding: with string values and one value per key,
the tree's range operations are 2.6 to 4.2 times *slower* than `btree-map`'s (`valuesBetween` 0.27 to
0.39, `prefix` 0.24 to 0.68), while point lookups stay faster; with `uint64` values they are faster
(`street`) or even (`dirs`). The cause is the leaf per key with a string header; it is what a page with
inline variable-length values removes. Job p1 (page layouts, M1, one process, 45 s, noise floors 6-25%) is in
(`bench/results-layout/step3-pages/m1-p1/`): on `street` A is faster for numbers, 1.63x present and 1.47x
absent (the PC native run: 1.19x and 1.53x); on `dirs` present keys 1.28x, within the noise, absent keys
0.92x (B faster, resolved); names against numbers and copy-out against bytes are unresolved on both
(1.0x and 1.05 to 1.06x, noise floors of 14 to 25%). Page bytes per key as on the PC (46.6 / 55.9 / 59.9
on `dirs`). One process on 160 MB of data is too coarse for the small differences; the PC result stands.
Nothing discarded.

## 2026-10-03 night: real data measured, variable-length values in lpage, waiting for the unique runs

Done and pushed: the two real data sets (`street`, `dirs`) with natural values and real strings, the
reference of the tree on them on the PC and the M1 ([step3-real-data.md](step3-real-data.md)); `internal/lpage`
with values of any length; `bench/cmd/pagebench` and its first results on the real keys
([step3-layout.md](step3-layout.md)): for numbers A is 16-19% faster per lookup and B 17-22% smaller; for
names B costs the same per lookup as for numbers. Running on the PC: the unique profile of both data sets
with `uint64` and with string values (`run-uni.cmd`, ends about 23:00), the reference for single-value
entries with strings. Waiting for the user on the M1: job r4 (`dirs`, strings) and p1 (the page layouts).
Nothing is discarded; the single-key page (SKMV) is next, after the unique reference is in.

## 2026-10-03 later: gate 2 ticked off, layout comparison of the multi-key page done, waiting for the user

The user ticked off gate 2 as "not met, deficits noted" and asked to begin step 3 with the layout
comparison of the multi-key page. Done and pushed: a model of memory per key for the sketched
length-header page against the page of step 2 (`pagefill -layout lens`), a prototype of it
(`internal/lpage`, tests 100%, race, fuzz) and microbenchmarks of both on Windows. Report:
[step3-layout.md](step3-layout.md). Short version: the length-header page (B) is smaller in the model
by 1-21% for every kind of key, but its lookups of short keys are 28-69% slower hot and cold and of
absent keys 2.4 times; for long keys (`path`) it is faster. The prototype's cold lookups for `u64`
are not understood; the user offered a run on the M1 for that. Nothing is wired into the tree.

**Correction (user, same evening):** B was made for values of variable length, and the comparison above
measured 8-byte scalars on synthetic keys only, so it does not decide against B; the recommendation
to drop B is withdrawn. Valid data: `street` (real keys, real locality counts; B is 16-22% smaller in the
model). **Next:** variable-length values in `internal/lpage`, a `street` benchmark with the real
locality names as string values, microbenchmarks on `street`, and a comparison with what the tree
does today for string values; a scalar special case for `u64` is accepted.

## 2026-10-03: vocabulary agreed, single-key page decided, step 2 gate still open

Worked out with the user: [GLOSSARY.md](GLOSSARY.md) (new words and the old ones they replace) and a
redesigned step 3 in [PLAN.md](PLAN.md). The user's decisions:

- A page may hold **one multi-value entry** (single-key page, "SKMV"); it replaces the flat, typed and
  set leaves. No page holds entries with different numbers of values for now.
- A remainder too long for the page stays inline in an **oversized object**, exempt from R1: no pointer
  to a key any more, one random cache miss less.
- Values are byte strings in the layouts under design, at most 255 bytes each while one length byte is
  used (a longer value needs an escape, open); fixed-size values get a specialized variant later. The
  number of values per entry is "what fits", not capped by the header.
- The common prefix of a page should, if possible, decide a mismatch within its first 64 or 128 bytes;
  the threshold is found by measuring. The entry count per multi-key page is not fixed: 3, 7, 11, 15 in
  the first sketch, 2, 6, 10 if another header byte is needed, or other with four size classes.
- Words: *path*, *remainder*, *common prefix*, *byte node*, *range node*, *end page*; `leaf` no longer
  names an object.

**Still open:** the decision on gate 2 (the four options of step2-results.md); the layout of the
multi-key page (sketch against `internal/vpage`); whether the byte node is needed once step 3 is done;
whether a later step lets a multi-key page hold multi-value entries (it would remove the single-key page
of two-value entries, the promote and the fall back). R4 is struck (it meant: no separate string object
for a key; the oversized object makes that true).

**Statistic for the header size** ([bench/results-layout/step3-entries](../../bench/results-layout/step3-entries/README.md)):
among the entries with several values, a header with room for 4 values covers 70% in the skewed bench
profile, 6 values 74%, 10 values 82%, 14 values 90%; street names (natural counts) 74%, 82%, 89%, 92%.
No header gets near 98%: about 5% of the multi-value entries have more than 64 values.

## 2026-10-02 night: step 2 done, gate 2 not met, waiting for the user's decision

The pages of `internal/vpage` hold every key with one value in the tree (`internal/art`), whatever its
length; design in [step2-design.md](step2-design.md), results and the reading against gate 2 in
[step2-results.md](step2-results.md), raw files in `bench/results-layout/step2-pages/`. Tests: 100%
coverage in `art` and `vpage`, `go test -race ./internal/...` and the fuzz test (`FuzzOperations`, 60 s)
pass, `golangci-lint` is clean.

**What the pages bring** against `node-pages` (and `btree-map`): string ranges and prefix queries 2.5-5.6
times faster, memory per key -4..-46% for five of seven kinds and below `btree-map` for all seven, 2-9
times less for the garbage collector to scan, point lookups of string keys at 256K 8-29% faster, multi
profile neutral.

**What they cost:** at 4K-16K keys point lookups 5-19% slower, `build` 3-30% slower, `churn` of email
and uuid 10-20% slower, `u64` 5-14% slower in every operation; `url` and `path` are below `btree-map`
at 4K-16K in ranges, `churn` and `build`. Gate 2 as written (no cell below 0.85, `u64` unchanged,
credo 1 and 2 for every kind) is not met.

**Decision for the user:** the four options and my recommendation (keep the old integer page next to the
new general page, accept the rest as the price of the design, and take the routing layer of step 4 on
next) are at the end of step2-results.md. Nothing in steps 3-6 starts before the decision.

## 2026-10-02 evening: step 1 done, page prefix built, waiting for the go for step 2

The page prototype (`internal/vpage`, 100% coverage, fuzzed, race and lint clean) and its
experiments are built and described in [step1-results.md](step1-results.md). It has a directory
(the user's idea of a FAT) in the first line or two of the page: a tag byte per key, in general
pages also the suffix length and the tail's offset. Every lookup takes **two dependent rounds**
(directory; then head, value and tail together), against 3 to 4 before. In the microbenchmark
(Ryzen, diagnosis only) cold lookups take `u64` 148 ns, `uuid` 212, `path` 247 (hot 6-12 ns).

**User's decisions (2026-10-02):** the rule R5 is restated as "at most two rounds of cache-line loads"
(STRATEGY.md, with a precise definition of a round); the page prefix may be tried, with the worry that
it costs too much in `churn`; commit and push at every valuable point without asking.

**The page prefix** (the bytes all keys of a page share, once, in the first line; chosen by a
rebuild if it saves 16 bytes or more) is built. Against the leaves of `node-pages` the pages now
need: `street` -42..-48%, `str` -23..-36%, `email` -17..-23%, `path` -11..-19%, `url` -8..-12%,
`uuid` -2%, but `u64` +2..+11% (a byte a key for the directory). **Churn** in the page model with and
without the prefix: -6% to +8%, build -3% to +7%, noise about 5%; the prefix changes in at most 12 of
1000 operations. **Lookup:** a page without prefix is as fast as before (faster: a one-load head
word), a page with one costs +2.7 ns hot (`path` 11.9 to 14.6), nothing cold, and absent keys get
16% faster. The first versions of the prefix cost more (30-40% cold for `u64`, 5-20% in churn); the
reasons and the repairs are in step1-results.md ("The page prefix").

Gate 1 as first written was not met for `u64` (memory) and, literally, for the lines of general
pages; as restated (rounds) and with the prefix it is met for every kind except `u64`'s memory,
which is the price of 40% faster cold lookups. **My recommendation:** go to step 2.

**Step 2 gets a lookup gate** (in plain words): the microbenchmark has no tree above the pages
and loops over pages that are all in order, so it can promise nothing about a lookup in the tree.
After the pages are wired in, point lookups of string keys at 16K-64K keys, measured against
`node-layout` interleaved as always, must reach at least 0.85, else pages are used only for short
suffixes. (0.85 is the credo's limit for any cell.)

## 2026-10-02 13:06: A/A job a1, gate 0 met

`a1` (the same as a0 with `-minprocs 8 -maxprocs 24`) ran 12:59-13:06, 7 minutes (estimated 20),
results in `origin/arm-results` (`a1/`). Every scenario was precise with its first 8 processes,
none needed more.

- **All 32 comparisons are precise** by the harness's criterion (a0: 30). The bounds of their
  intervals stay within ±1.9 points, except for `valuesBetween` (up to -2.8).
- **All 32 point estimates lie between 0.98 and 1.01** for identical code (differences -2.5% to
  +1.5%); 29 of them within ±1%.
- **`valuesBetween` is the noisy operation** on this machine: `str unique` 0.98 at 4K (-1.7%) and
  16K (-2.5%) as in a0, `u64 multi 4K` +1.5%, and the other range cells ±0.7%. The harness marks
  nine cells "resolved" although the code is identical (all of them range scans or within 1%);
  read an arm64 `valuesBetween` difference below 3%, and any other below 1.5%, as noise.
- The machine was not at rest again: the runner warned (load 6.6 in the last minute before the
  build, 7.5 at the end), and I did not ask what else was running. Precision was good anyway; a
  quieter machine may be better still, and nothing here says it would not.

Gate 0 reading: met. The M1 Pro can measure; its noise floor is about ±1% for point operations,
`churn` and `build`, and about ±3% for ranges.

## 2026-10-02 afternoon: A/A job a0 on the M1 Pro

`a0` (same code against itself, 6b06dd2, u64 and str, 4K and 16K, multi and unique, four
operations) ran in 4.5 minutes, from 12:46 on 2026-10-02, results on `origin/arm-results`
(`a0/`). Of 32 comparisons:

- all 32 point estimates are within 0.97-1.01 (differences -2.8% to +0.9%), 30 within ±1.1%;
- 30 are precise by the harness's criterion (95% interval within ±2 points or 10% of the
  difference); the other two have wide intervals (±3.7 and ±4.0 points): `u64 unique 16K
  valuesFor` has one outlier process (+7.1% against -0.4% to +0.6% for the other seven),
  `u64 multi 16K churn` scatters between all eight (-4.7% to +3.0%). The scenarios were capped
  at 8 processes, the harness wanted 29 and 32;
- **systematic bias:** `valuesBetween` with str keys and one value per key reads 0.97 (4K) and
  0.98 (16K) in every one of its four processes (-2.0% to -3.2%) for identical code. The
  harness marks these "resolved"; they are not: its A/A validation runs inside one binary and
  does not see the layout difference between the library and the baseline copy;
- the runner's "load at start" in `env.txt` was my own compile (29): fixed, the runner now notes the load
  before the build, waits 30 s after it, and warns at a load of 2;
- the machine was not at rest (load 3.8, 15-minute average 6.1 at the start).

Gate 0 reading: numerically met (0.97-1.01), with the two caveats above; `a1` above settled them.

## 2026-10-02: step 0 done, waiting for gate 0

**Parked and pushed:**

| branch | commit | content |
|---|---|---|
| `node-layout` | `7b8a8d8` | pessimistic paths, leaves with key remainders, flat and typed leaves, results against `main` and the competitors |
| `node-pages` | `cf6944b` | pages for integer-like keys with one value, range nodes, results in `bench/results-layout/node-pages/` |
| `leaf-pages` | `b0aeed1` | the earlier page experiments (K, S, U8-n pages) and their measurements |

**Working branch:** `cacheline`, forked from `node-pages`. It holds these documents and step 0.

**Step 0 (not yet committed, see PLAN.md):**

- 0.1 `art.Map.Objects`, `art.Block`, `bench/cmd/objstat`, tests, the statistic of `node-pages`.
- 0.2 `keys.Corpus.Probes`, used by `valuesFor`; `bench/README.md` says numbers are not comparable.
- 0.3 `bench/remote/` (queue, runner, test); run once end to end on Linux.
- 0.4 `go test ./... -race` 3.5 minutes instead of more than 20.

**Next action:** the user decides whether gate 0 is passed. Then step 1 (PLAN.md).

**Open points:**

- **arm64 runner on macOS** is untested (see MEASURING.md).
- **Probe order**: the 64K-and-above numbers are expected to be comparable with the old order but
  that was not measured. A short A/A of old and new probes at 64K would settle it.
- **Race test time**: still 3.5 minutes. If CI needs less, shrink the key sets further.

## Test run time (2026-10-02)

`node-pages`, WSL, machine otherwise idle:

| run | `internal/art` | `TestAgainstReference` |
|---|---|---|
| `go test -race -cover ./...` (default 10-minute timeout) | failed: timeout | still running at 600 s |
| `go test -race -cover -timeout 60m -v ./internal/art` | passed, 100% coverage, 1175 s | 1121 s |
| `go test -run TestAgainstReference ./internal/art` (no race detector) | | 43.5 s |

All other tests together take about 55 s with the race detector; the slowest are
`TestRemoveAbsentNearMisses` (25 s) and `TestLongPaths` (7 s).

**Fixed in step 0.4:** under `-race` the big key sets of `TestAgainstReference` are a tenth of
their size and the 64K-path set runs in two of the five leaf modes. `go test -race -cover
./internal/art` then takes 205 s with 100% coverage (the whole `./...` run 208 s). One
deterministic set was added (a term arriving at a full node) because the smaller random sets no
longer reached `setTerm`'s growth.

`TestAgainstReference` runs five leaf modes over every key set, against the reference model.
Each run has six phases, and after each phase it compares the whole map and checks the
invariants of the whole tree. The race detector slows it down about 26 times.
