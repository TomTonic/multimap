# Step 4: what the benchmark's own streams do to the pages (the probe of 2026-10-05, night)

Words as in [GLOSSARY.md](GLOSSARY.md). Raw: `bench/results-layout/step4-probe/` (`wsl-probe.txt`: the probe, `mem/`: heap per key over the sizes). Machine: WSL2 on the Ryzen 9 7900,
one process, **one run per case: the times scatter by about 15 %, read the shares and the order of magnitude, not the last digit**. The code is commit `1732f5c` plus the counters and the change
of the merge trigger described below. The PC measurement of the fixed code (option A) follows from this (see the end).

The question was: why are `churn` and `build` with one value per entry far below the prediction, and what does that teach for step 5 (multi-value entries in pages)? The answer is not what the
step-4.2 report guessed.

## How the probe works

`TestProbe` in `bench/cmd/bench/probe_test.go` (runs only with `MKPROBE=1`; case from `MKPROBE_KEYS`, `MKPROBE_VALUES`, `MKPROBE_N`, `MKPROBE_PROF` for a CPU profile):

```
cd bench
MKPROBE=1 MKPROBE_KEYS=street,dirs,u64 MKPROBE_VALUES=natural,single-value MKPROBE_N=4096,65536 \
  go test -tags mkstats -run TestProbe -count=1 ./cmd/bench -v
```

It builds the fixture as the benchmark does, replays the very streams of `workload.Cycle` and `workload.Build` (the benchmark's `churn` and `build`: `-ratio 2 -permchurn 0.25`) on an `art.Map`, and prints:
the objects of the tree (before, after a cycle, after the build stream, after removing half the entries value by value), the **number of values a key has before each operation** with the time of the operation by
that class (the clock's own cost taken off), how the number of values of the keys is spread at the start and in the middle of a cycle, and (with the build tag `mkstats`) the events of the pages
(`internal/art/events_on.go`: pair, promote, burst, widen, merge tried and done, ...), each with the number of entries of the page it concerned. Without the tag the counters do not exist (`events_off.go`: empty
functions); `go test -run TestEvents` checks both builds.

## Finding 1: the hypothesis in the step-4.2 report was wrong for `single-value`

The report said: the streams add transient values to existing keys, they promote the pages, the entries never come back. **That holds for `natural`, not for `single-value`**: there every transient value gets a key of its own
(`newPairs`), no key ever has two values, and the probe counts **no promote at all**. What costs time there is something else.

`street` `single-value`, 65,536 keys (171,410 operations in a cycle), before the change below:

| what | share of operations | time |
|---|--:|--:|
| insert a new key | 50 % | 225 ns |
| remove a key | 50 % | **451 ns** |

The profile of the replay (40 cycles): **`mergeUp` and `tryMerge` take 28 % of all the time, `isPage` (the first touch of a child's header: a cache miss) another 14 %.** Of 100 merges tried
(after every removal of a key), **4 succeed** (368 tries and 15 successes per 1000 operations). A page in use holds some ten entries (median 10 when an entry is removed), its siblings as many, a node has up to 12 children: the
sum never fits one page of 512 bytes. Every try touches the headers of up to 12 children, which are not in the cache.

Without any merge (`mergeChildren = 0`, an experiment) the same cycle takes **196 ns per operation instead of 349** (the replay 190 against 331), and the tree has 91.7 % of its keys in pages instead of 92.9 %: the merge
earns almost nothing in a steady state. It earns its keep after mass removals (the step-4.2 measurement of the memory after removing half the keys).

**Change (this commit): try a merge only when the page the removal came from is nearly empty (`mergeBelow = 2` entries left), or when the removed entry was not in a page.** The sweep (`street` `single-value`, 65,536 keys;
churn in ns per operation, and bytes per key of a tree that had every second key removed value by value; one run each):

| entries left in the page at most | 0 | 1 | 2 (chosen) | 3 | 6 | 12 | any (before) |
|---|--:|--:|--:|--:|--:|--:|--:|
| churn, ns per operation | 205 | 224 | 225 | 239 to 279 | 247 | 297 | 349 |
| bytes per key after removing half | 43.1 | 42.2 | 41.8 | 41.6 | 41.0 | 40.6 | 40.4 |

Two entries left: **a third off the time of `churn`, 3.5 % more memory after a mass removal.** The choice is the knee of the table, not a derivation; the number is a named constant (`mergeBelow`) with a test that
shows the behaviour (`TestMultiKeyPageMergeWaitsForNearlyEmptyPage`). `RemoveKey` follows the same rule. `natural` does not change (its time is in the other findings).

## Finding 2: with several values per key the pages decay, and that is the real price of "pages hold one value per entry"

`street` `natural`, 65,536 keys, the same tree at three moments (heap blocks per key, `art.Block`):

| tree | bytes per key | keys in multi-key pages |
|---|--:|--:|
| built from the corpus (what the benchmark's memory phase measures) | **58.3** | 41.6 % |
| after the build stream (`workload.Build`: a history with transient values) | **70.7** | |
| after one steady-state cycle (the same keys and values as at the start) | **73.2** | **15.9 %** |

**The tree gets 25 % bigger by use, with the same content.** A key that gets a second value leaves its page (promote: the page is built again around it, as a byte node with single-key pages and smaller pages),
and when the second value goes away the key stays a single-key page: nothing puts it back. Per 1000 operations there are 32 promotes (median page of 2 entries, 90th percentile 4, largest 15), and 27 pages with one entry
left that become a single-key page; 41,557 single-key pages with one value after a cycle against 24,685 in the fresh tree. The step-3.5 pages (one entry each) did not have this drift. **The memory advantage the benchmark's
memory phase reports for `natural` (a fresh build) is therefore 58 per key; a tree with a history has 73 (still 285 for `btree-sets`).**

The number of values of the keys the streams touch (`street` `natural`; the first column is keys that only the streams know, absent at the start):

| when | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7+ |
|---|--:|--:|--:|--:|--:|--:|--:|--:|
| start of the cycle | 33.3 % | 52.8 % | 6.5 % | 2.4 % | 1.3 % | 0.7 % | 0.5 % | 2.4 % |
| middle of the cycle | 27.1 % | 49.7 % | 14.5 % | 3.6 % | 1.4 % | 0.8 % | 0.5 % | 2.3 % |

Of the keys that have values: 79 % have one, 10 % two, 4 % three, 2 % four, 3.6 % seven and more (`dirs`: 62 / 17 / 7 / 4 / 7; `uint64` keys: 50 / 11 / 12 / 12 / 13: values 1 to 4 are 94 %, 89 %, 85 %).
The time of an operation by what it does (`street` `natural`, 65,536 keys, ns): insert a new key 255; insert into a key with 1 value 322 (the promote); with 2 values 265; with 3 or more 240; **remove a key 555**;
remove a value from a key with 2 values 234; with 3 or more 253. Removing a key costs twice as much as the rest in all the cases of the table (`dirs`: 835 against 330; `uint64`: 259 against 187).

**Experiment: demote.** A removal that leaves a key with one value in a single-key page also tries the merge (`MKDEMOTE`, not kept): the pages come back (41.5 % of the keys) and the tree after a cycle is 63.1 bytes per key
instead of 73.2 (the fresh tree has 58.3), but the time of a cycle grows by 37 % (337 against 463 ns), because the merge tries fail again and again (the siblings are not all single-value pages). That is option D of the step-4.2 report:
**it works, and it is too dear.** What would not have a price is a page that does not dissolve.

## Finding 3: the memory depends on the number of keys where keys are random (not in the 4.2 report, which measured 262,144)

Heap bytes per key, `Ordered` / `btree-map` (`single-value`, `-memrounds 2`, WSL; `mem/`; the sizes of the dev suite for speed are 4,096 and 16,384, for memory only 262,144):

| keys | 1,024 | 4,096 | 16,384 | 65,536 | 262,144 |
|---|--:|--:|--:|--:|--:|
| `uint64` (random) | 40 / 57 | 23 / 48 | **51 / 45** | 38 / 45 | 24 / 45 |
| `uuid` | 76 / 96 | 77 / 88 | 63 / 85 | 66 / 85 | 57 / 85 |
| `email` | 68 / 77 | 47 / 69 | 63 / 67 | 48 / 66 | 48 / 66 |
| `str`, `url`, `path`, `street`, `dirs` | all below `btree-map`, 20 to 60 % | | | | |

(`objstat`: `uint64` single-value 16K = 50.6, 32K = 49.1, 64K = 37.9, 128K = 29.5, 256K = 24.0, 512K = 22.9.) With random 64-bit keys, 16,384 keys spread over 2 levels of 256 branches: most
branches below the second level hold one key, a page needs two keys under one byte node, so (reading of the table, not a measurement of the cause) **86 % of the keys sit in single-key pages and the byte nodes above them are the cost**. **At 16,384 `uint64` keys `Ordered`
is above `btree-map` in memory (51 against 45): the one case where credo 1 fails for memory.** At `natural` the numbers are always far below `btree-sets` (123 to 222 against 280 to 450). This is
exactly what range nodes (step 6: a page above a byte node that holds the keys of several branches) would change; it is an argument for step 6 with a number, not a reason to change step 4 (D1).

## What this means for step 5 (multi-value entries in pages)

1. **Entries with several values must stay in their page.** The promote is the cause of the 25 % growth and of the 32 rebuilds per 1000 operations; the demote experiment shows that bringing the keys back costs 37 % of the time.
   In a page with multi-value entries, a second value is a change of the page in place (like `Added`), and a removed value does not change the page's place in the tree.
2. **How many values per entry?** In the real data 85 to 94 % of the keys that have values have at most 4 (`street` 79 % have exactly one). An entry in a page with up to 4 to 8 values
   covers nearly all of the keys; the long tail (7 and more: 4 to 13 % of the keys) stays in single-key pages and value overflows, as now. The design note of step 5 has to pick the count from these tables and the page model, not from the
   probe alone (the model `skmodel -multi` takes the corpora's value counts).
3. **The merge is for removals of keys only**, and the trigger of Finding 1 is the right shape for it; with step 5 a removed value never needs a merge.
4. **Removing a key costs twice an insertion** (555 against 255 ns): a descent to the page, the shrink of the page (a copy of up to 512 bytes), the merge tries. The profile of `natural` shows `isPage` at 30 % of the time
   (the first touch of a header on the way down is a cache miss at 65,536 keys and more): the number of levels, not the pages, is the cost there. Step 6 (range nodes) takes levels away.
5. **The benchmark's memory phase measures a fresh tree.** For a multimap in use the figure that counts is after a history (58 against 73 for `street` `natural`); the large measurement of 0.8 should report both
   (the probe's `census` of the tree after `workload.Build` and after a cycle does it for a case).

## The PC measurement of the fixed code (option A), 2026-10-06 night

Commit `3e1e952` (the merge trigger of Finding 1), the same ten cases as step 4.2 (`run-m43.cmd` in `bench/results-layout/step4-probe/pc-m43/`, 23:07 to 00:29, 6 to 8 processes), against step 3.5 (`fe120f5`) and
`btree-map` / `btree-sets`. Ratio is `Ordered` speed over the other (below 1 is slower), step 4.2 (`a166cad`) → now.

**`churn` and `build` with one value per entry** (`single-value`), against step 3.5: `churn` 0.03 to 0.28 → **0.69 to 1.21**, `build` 0.04 to 0.19 → **0.66 to 1.03**; against `btree-map` `churn` 0.05 to 0.30 → **0.69 to 2.07**
(`uint64` keys 1.68 to 2.07), `build` 0.07 to 0.19 → **0.63 to 1.80**. The 4,096-key cases are the lowest (0.63 to 0.88 against `btree-map` for `street`, `dirs`), the large sizes (86,215, 212,449, 262,144) are at 0.93 to 1.24.
**The real mix** (`natural`) against `btree-sets`: `churn` 0.91 to 1.64 → **1.00 to 2.54**, `build` 0.88 to 1.62 → **0.99 to 2.43**: credo 1 is met in all cells (the lowest are 4,096 `dirs` `build` 0.99 and `churn` 1.00), against step 3.5 0.64 to 0.92 (was 0.54 to 0.86).
`valuesFor`, `valuesBetween`, `prefix` do not move (0.80 to 1.33, 1.01 to 2.24, 0.95 to 2.86 against step 3.5), as the change does not touch them.

**Memory** is unchanged to the byte per entry (27, 41, 31, 51, 24 with one value; 73, 96, 98, 153, 123 real); after removing every second key +0 to +2 bytes (the 3.5 % of the sweep) and for two cases -2.
GC per cycle +2 to +6 ms (`single-value`), +8 to +22 ms (real).

**Gate 4 now:** credo 1 for `churn` and `build` is met for the real mix and for the large sizes of `single-value`, **not** for `single-value` at 4,096 and 16,384 keys against `btree-map` (0.63 to 0.89); credo 2 for ranges is met for the real mix
(1.5 to 3.0 against `btree-sets`) and **not for `single-value` against `btree-map`** (`valuesBetween` 0.30 to 0.68, `prefix` 0.40 to 0.65; `valuesFor` 1.05 to 4.1 is met). The M1 job `m43` (same commit) is queued.

What is left of the gap is, from the probe: the scan of a page (two closures per entry, 8.2 ns a value against `btree-map`'s fewer ns; `valuesBetween` in a page-heavy tree), and at 4,096 and 16,384 keys the remaining cost of
the removal of a key (a page copy, the merge tries at pages with one or two entries). Neither is a question of the design; they are the options B of the step-4.2 report (scan, allocation-free page change).

