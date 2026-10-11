# Candidates after the counts of 2026-10-10: a decision paper

Words as in [GLOSSARY.md](GLOSSARY.md). Written 2026-10-11 for the user's decision: which changes are built, in which order, and how
each is judged. It only gathers what the notes of 2026-10-10 found; the evidence is there:
[profiles-2026-10.md](profiles-2026-10.md) (shares of all scenarios), [descent-analysis.md](descent-analysis.md) (the descent, PC by
instruction, M1 by line, the model's second lines), [write-path-analysis.md](write-path-analysis.md) (objects a write, merge tries,
the pendulum of the page forms, free list), `bench/results-layout/c1-m1/matrix.md` (the state before any of this on the M1).

The working rules that shape the list (PLAN.md): structured keys set the course, random keys only stay in the band; no candidate may
buy one scenario with another; hot paths in "assembler thinking"; speed claims only from rtcompare.

## Where the ordered map stands (c1, M1, today's code)

Against the B-trees (the band of the credo): natural is faster everywhere (1.07 to 4.85 times `btree-sets`). Single-value: lookups
faster everywhere (1.17 to 2.10 times `btree-map`); **writes of dirs and url up to 16,384 keys slower (churn 0.79 to 0.96, build
0.71 to 0.87)**, street and links about even or faster; **ranges 0.31 to 0.73 times `btree-map`** (the known weakness of single-value
ranges). Against main: small maps slower (0.63 to 0.99 at 4,096 keys), about even at 65,536 (0.85 to 1.22), ranges 1.4 to 2.5 times faster. Memory: the
least of all candidates everywhere.

So the gaps are: **writes of small single-value maps of path-like keys** (outside the band), **single-value ranges** (outside the
band), and the distance to main in small maps (inside the band, but the largest everyday cost).

## The candidates

Shares are of the CPU profiles (PC, 4,096 keys, single-value unless said), not speeds. "Gain at most" is the share the candidate can
remove at best.

| | candidate | evidence | gain at most | touches | risk |
|---|---|---|---|---|---|
| **W1** | **Cheaper merge tries.** Keep the merge decisions exactly as they are, but (a) decide the clear refusals from the page heads (`Used`, O(1) a child) before walking every entry with `Each` and a closure; (b) remember the path of the removal instead of a second descent from the root; (c) no closures in this path. | tries in 5 to 22 % of the writes, 0 to 1 a 1,000 done; `mergeUp` 11 to 16 % of a churn for url, path, street, dirs, links (str 6, email 1, natural 1 to 9); of it `mergeFits` 6 to 10 %, the rest 2 to 8 %; refusals by size at 2.1 to 2.3 times the limit in the median | 11 to 16 % of a churn of path-like keys, a few % of a build | writes only | low: same decisions, so the same tree; a test can compare the shapes |
| **W2** | **Pair in place.** When a new key meets a single-key page and the two fit its object, turn the object into the many-key form in place instead of building a new page. | pairs 20 to 90 a 1,000 writes (182 for email at 16K); they fit the old object in 43 to 63 % for most single-value cases (corrected 2026-10-11; first counted as 70 to 99 %); `toOneKey` + `pairWithFixed` 1 to 6.5 % of a churn plus the allocations | 1 to 4 % of a churn | writes only; memory: the pair keeps its object's class (may be one class larger than today's new page; to be counted) | low |
| **W3** | **Free list, LIFO, bounded.** | 0.06 to 0.34 objects made a churn write; 69 to 87 % could come from the list, a few hundred bytes after the drop; allocation and GC about 5 % of a churn, 14 % of a build; fresh-object warmth only 2 to 3 % | about 5 % of a churn, 14 % of a build (less after W2) | writes; memory held by the list (needs a bound: up to 1.7 MB seen) | medium: per map state, the bound |
| **D1** | **A node step without the switch on the node type** (one search for the small nodes, the type deciding by arithmetic or a table rather than by branches). | M1: the switch and the key byte 24 to 26 % of `find` (by line); PC: within the rest of 21 to 36 % (by instruction); `find` is 40 to 57 % of a lookup and 22 to 40 % of a write | up to a quarter of the descent on the M1, less on the PC | every operation | medium: needs a design note first; the gain on the PC is uncertain |
| **D2** | **The child pointer in the node's first line** (N12 from the fifth child, N26, N58). | PC: 16 to 19 % of `find`, the model 1.5 to 5.5 second lines a lookup; **M1: 0.4 to 1 %** | PC only | every operation | medium: a new node layout; helps one machine |
| **S1** | **The search in the page** without a larger page (the SWAR on length and first byte of page-search-design.md). | 19 to 43 % of a lookup, 10 to 21 % of a write, nothing of a range | large, but not counted yet | point operations | medium: not yet counted what it saves |
| **D3** | **Fewer objects a lookup** (step 6: the routing layer). | the wait for the next object is the largest single part of the descent on both machines (PC 34 to 57 %, M1 35 to 39 % of `find`); the packing model cuts objects a lookup by 11 to 52 % | the largest | every operation | high: no design yet that keeps writes cheap (the bundle did not convince) |
| — | not a lever: the value sets (`Set3`) | the order of the values does not matter; their share is reading many values | | | |
| — | not a lever: the page layout recomputed (`lay`) | computed about once a write (0.75 to 1.26); its 5 % is the one computation (a function not inlined) | at most a part of 5 % | | |

The single-value ranges (0.31 to 0.73 of `btree-map`) are in none of the cheap candidates: their walk is the routing layer (D3),
profiles-2026-10.md.

## Recommended order

1. **W1 and W2 together, as one candidate** (`ordered` on a branch against `baseline` = `cacheline` before them). Both are cheap,
   change no decision of the tree (W1) or only where an object is made (W2), touch the writes only and aim straight at the gap that is
   outside the band (writes of small maps of dirs and url). Measuring them together saves a night; if the result is a surprise they
   are split.
2. **D1**, after a design note with a prediction: the one descent lever that the M1 shows clearly and that touches every operation.
3. **S1**, after a count of what it saves (comparisons and lines a lookup, as the model counts).
4. **W3**, after W2 (it will have less to do then).
5. **D2** only if the PC's result matters on its own; **D3** stays the large open question of step 6.

## How a candidate is judged (rtcompare, PC and M1)

The candidate on a branch is `ordered`, the code before it is `baseline` (cmd/mkbaseline from the commit before), and the B-trees run
alongside: `-keys street,dirs,links,url -values natural,single-value -sizes 4096,16384,65536 -ops valuesFor,valuesBetween,churn,build
-vs baseline,btree-sets,btree-map`, memory included. All four operations, so that a gain in writes cannot hide a loss in lookups or
ranges. Then u64 and uuid once, only to see that they stay in the band.

**Prediction for W1 + W2** (ordered against baseline; written before W2's count was corrected, kept as written): churn single-value at 4,096 keys +8 to +15 % for url, path and dirs, +5 to
+12 % for street and links; churn natural +2 to +8 %; build +2 to +6 %; lookups and ranges within noise (±2 %); memory within ±2 %. On
the M1 the same direction; dirs and url churn single-value at 4,096 keys from 0.86 and 0.79 of `btree-map` to about 0.95 and 0.88:
better, still below 1. A result outside these ranges stops the work for a look before the next step.

## Decisions asked

1. The order above, or another.
2. W1 and W2 measured together (recommended) or apart.
3. D1: a design note next, while W1 + W2 are measured.

## Decision (user, 2026-10-11 morning) and what was built

The order as recommended; W1 and W2 measured together; D1's note depending on the time.

Built on `cacheline` in `e8d370e` (its parent `df110e6` is the baseline of the measurement):
- **W1** (`internal/art/mkkey.go`): `mergeFits` first takes `mergeBound`, a lower bound of the merged page's need from the heads of
  the child pages (`Used`, `Len`, `PrefixLen`: every value's bytes and every key's own bytes); only a bound within `mergeFill` is
  checked exactly (`mergeNeed`, the old walk). The bound is never above the need (`TestMergeBoundNeverAboveNeed`, every node of maps
  of fixed, string and pointer values), so the decisions are the same: the probe counts the same bursts (27,542), merges (429) and
  pairs (159,361) over the 28 cases as before. The second descent of `mergeUp` (the path of the removal) is not changed yet.
- **W2** (`internal/page`): a one-key page that meets a new key makes the page of the two in its own object when they fit its
  class and the page holds no pointer (`pairWithFixed`, `Str.pairWith`): the old content goes to a buffer on the stack (at most 512
  bytes), the object is cleared and written in the many-key form. `BuildFixedOf` and `BuildStrings` are split into the layout
  (`fixedLayout`, `strLayout`) and the writing (`fillFixed`, `Str.fill`) for it. Pages of pointers keep the old path (their object's
  type fixes its split of words). A pair in place keeps the class of the old object, which may be one class above what a new page
  would take: the memory check of the run tells whether it matters.
- Tests: 100 % of both packages, `-race`, fuzz 60 s each of `FuzzPage` and `FuzzOperations`, lint 0.
