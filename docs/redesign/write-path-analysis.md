# What a write does with objects, and where its time goes (2026-10-10)

Words as in [GLOSSARY.md](GLOSSARY.md). Lever 3 of the order of 2026-10-10 (PLAN.md): the writes of small maps, where the ordered map
is 0.65 to 0.98 times `btree-map` (single value a key) and 0.89 to 0.96 times `btree-sets` (strings, 4,096 keys). The user's
direction: count, measure and think before building anything; the ideas stay open side by side. The user's first idea: a free list
of objects a size class, LIFO, so that an object taken again (above all in churn) is likely still in the cache.

## Questions

a) **Free list, LIFO.** How many objects does a write make and drop, by type and size? How long does a dropped object lie before an
   object of the same type and size is wanted again, in operations and in bytes made since? That gives the hit rate of a LIFO stack
   a type and size, how deep the stack gets (memory held), and how likely the object is still in a cache.
b) **Growing without copying.** How many writes change their page in place, how many replace it by a new object (another class),
   how many change the tree (burst, merge, nodes)?
c) **The head helpers** (`mem`, `lay`, `cpl`: 8 % of a churn at 4,096 keys): how often a write computes the same again.
d) **Changing the tree:** events a write (bursts, merges) and their cost.
e) **Allocation and garbage collection:** allocations and bytes a write in all (with temporaries), against the objects of the tree.
f) **Where a page change spends its time:** waiting for a fresh, cold object (stores that miss), copying, or the work; at the
   level of source lines, as descent-analysis.md did for the descent. This also prices what LIFO's warmth could give.

## Method

No change of the library but one diagnostic field: `art.Object` gets the address of the object, so that a probe can tell objects
apart. A new probe (`MKLIFE`, bench/cmd/bench) builds the tree as the bench does, replays one churn cycle to the steady state, then
replays one more cycle operation by operation and takes the set of the tree's objects (`Map.Objects`) before and after every
operation: what is new was made by it, what is gone was dropped. The same for the bench's build stream from empty. A LIFO stack a
key (label and size of the object) is simulated over that stream: an object made can be taken from the stack only if an earlier
operation dropped one of its key. Not seen: the `Set3` of a value overflow (Objects does not report what it allocates) and
temporaries (they are in e, from `runtime.MemStats`). Questions c, d and f from the probe's event counts (build tag mkstats) and
from CPU profiles of 10 seconds of churn and of the build stream.

Cases: the seven structured kinds (str, email, url, path, street, dirs, links), single-value and natural, 4,096 and 16,384 keys (the
gap is in small maps).

## Prediction (written before any of these counts)

1. In a steady-state churn of single-value maps at 4,096 keys, 60 to 75 % of the writes change their page in place (no object
   made or dropped); most of the rest replace one object by one of another class; bursts and merges are under 2 % of the writes.
2. Objects of the tree made a write: 0.25 to 0.5 in churn, 0.4 to 0.8 in the build stream; allocations of the runtime a write
   (with temporaries) at most twice that.
3. LIFO hit rate (object made from an earlier drop of its key): at least 80 % in churn, 40 to 70 % in the build stream (it grows,
   so it drops fewer than it makes). Median distance from drop to reuse at most 30 writes in churn, so a reused object is likely
   still in L1 or L2.
4. Allocation and garbage collection are at most 8 % of a churn and at most 20 % of a build of 4,096 keys (the profiles of
   2026-10-10 said 5 and 17 % in the mean); what a free list saves directly is bounded by that.
5. In the page change, waiting for a fresh object is visible but a minority: at most 30 % of the page change's time.

## Runs

- Life probe (`MKLIFE`, `lifeprobe_test.go`, build tag mkstats for the events), PC/WSL, 2026-10-10 13:23 to 16:22 (counts, no times; the
  last cases in parallel processes). All 28 cases but the build stream of links natural 16,384 (about a million writes, each with a
  census of the tree: stopped after two hours). `bench/results-layout/write-2026-10/life.txt`, summary `life-summary.md`.
- CPU profiles of 10 seconds of churn and of the build stream (`MKPROBE`, the knob `MKPROBE_SECS` added), the 28 cases, PC/WSL alone,
  16:23 to 16:32; shares by category as in profiles-2026-10.md (`profile-shares.md`), functions by share (`top-*.txt`).

## Result

**What a write does with objects** (single-value, 4,096 keys; natural in the text below the table):

| kind | churn: in place | churn: objects made a write | LIFO hit | writes from drop to reuse, median | build: objects made a write | build: LIFO hit |
|---|--:|--:|--:|--:|--:|--:|
| str | 93 % | 0.06 | 69 % | 74 | 0.21 | 66 % |
| email | 94 % | 0.06 | 87 % | 77 | 0.29 | 68 % |
| street | 83 % | 0.13 | 81 % | 26 | 0.23 | 64 % |
| links | 75 % | 0.19 | 80 % | 21 | 0.31 | 66 % |
| dirs | 70 % | 0.23 | 80 % | 23 | 0.42 | 68 % |
| path | 66 % | 0.27 | 80 % | 21 | 0.49 | 68 % |
| url | 56 % | 0.34 | 83 % | 16 | 0.53 | 69 % |

Natural: 85 to 97 % of the churn writes in place, 0.03 to 0.09 objects made a write. The runtime allocates about what the tree makes
(churn: nothing besides; build: 10 to 20 % temporaries). Between a drop and its reuse only a few hundred bytes of objects are made
(median 160 to 1,300 bytes); a reused object is likely still in L1 or L2. A free list would hold memory: up to 0.8 MB for natural
16,384, and links natural keeps some objects for hundreds of thousands of writes (p90), held at most 1.7 MB: it needs a bound.

**Events a 1,000 writes of a churn at 4,096 keys** (mkstats): bursts and merges done are at most 3; but

| kind, single-value | merges tried | merges done | page with one key left becomes a single-key page | single-key page and a new key make a page | page shrinks |
|---|--:|--:|--:|--:|--:|
| str | 49 | 0 | 20 | 20 | 15 |
| email (16,384 keys) | 7 (29) | 0 (0) | 7 (182) | 7 (182) | 26 (135) |
| street | 82 | 0 | 39 | 39 | 34 |
| links | 82 | 0 | 57 | 57 | 43 |
| dirs | 131 | 0.1 | 66 | 66 | 49 |
| path | 168 | 0.5 | 71 | 71 | 57 |
| url | 216 | 1.1 | 90 | 89 | 73 |

Two things nobody looked for:
- **Merges are tried in 5 to 22 % of the writes and almost never done.** The try (`mergeUp`, inclusive) costs 11 to 16 % of a churn
  of single-value url, path, street, dirs, links at 4,096 keys (str 6 %, email 1 %; natural 1 to 9 %).
- **A pendulum between the two forms of a page:** a page that keeps one key becomes a single-key page (a new object), and the next
  key beside it makes a page of two again (a new object): 20 to 90 times each way a 1,000 writes, 182 for email at 16,384 keys. Most
  of the "one object for one" writes are this. Its own code (`toOneKey`, `pairWithFixed`, inclusive) costs 1 to 6.5 % of a churn,
  besides the allocations.

**Where a churn's time goes** (single-value, 4,096 keys, mean over the seven kinds): descent 22 %, search in the page 19 %, the page
head's layout (`lay`, `mem`) 7 %, changing the page 16 % (`Add`, `Remove` own code 12 %), memmove 9 % (85 to 90 % of it the shifts
inside a page in `Add` and `Remove`, the rest copies into new objects), changing the tree 15 % (the merge tries the largest part),
allocation and garbage collection about 5 %. Build: allocation and garbage collection about 14 %, the rest alike.

**Against the prediction:**
1. In place 60 to 75 %: missed in its range (56 to 94 %: the path-like keys replace their page far more often than str and email).
   Bursts and merges done under 2 %: hit; but the merge tries (not predicted) are 5 to 22 %.
2. Objects made 0.25 to 0.5 a churn write, 0.4 to 0.8 a build write: missed, fewer (0.06 to 0.34, build 0.21 to 0.53; email at 16,384
   keys 0.53). The runtime's allocations at most twice that: hit (1.0 to 1.2 times).
3. LIFO hit at least 80 % in churn: mostly (69 to 87 %); 40 to 70 % in build: hit (64 to 69 %). Median distance at most 30 writes: hit
   for the path-like keys (16 to 31), missed for str, email and natural (50 to 150); in bytes made in between it is small everywhere.
4. Allocation and garbage collection at most 8 % of a churn and 20 % of a build: hit (about 5 and 14 %).
5. Waiting for a fresh object at most 30 % of the page change: hit, far below (copies into new objects and the clearing of new
   objects together about 2 to 3 % of a churn).

**What it says for the ideas of lever 3:**
- the free list (LIFO) can save at most about 5 % of a churn and 14 % of a build directly, and the warmth of a reused object concerns
  only the 2 to 3 % that touch fresh objects; it would need a bound on the memory it holds;
- the futile merge tries (11 to 16 % of a churn of five of the seven kinds) and the pendulum between the page forms (1 to 6.5 % plus
  the objects; most of the object traffic of churn) are larger and were not on the list; both are questions of when the tree
  changes its shape (hysteresis), not of allocation;
- the shifts inside a page (memmove 9 % with `Add`/`Remove` 12 %) and the recomputed layout (`lay` 5 %) are the page's own costs.

## Follow-up (user, 2026-10-10 evening): why the merge tries fail, and what a hysteresis of the page forms would cost

**What the code does.** After a removal that leaves a page at most `mergeBelow` (2) entries, `Remove` and `RemoveKey` call `mergeUp`,
which descends from the root again along the key and tries, from the lowest node up, to merge a node whose children are all pages
(`tryMerge`, `mergeFits`: it walks every child page with `Each` and a closure and adds up the bytes the merged page would need,
against `mergeFill`, half of the largest page). The pendulum: a page left with one key turns into the one-key form in place (no new
object), and then often shrinks to a smaller class (a new object); a new key beside a single-key page always makes a new page of two
keys (`pairWithFixed` builds one), even if the two keys would fit the object it has.

**Counted** (build tag mkstats, counters only, the code does not change its behaviour): what started a `mergeUp` (1 or 2 entries
left); every refusal by its reason (the node has more than 12 children; a child is a node or a value overflow; the children hold more
than 64 values; the merged page would need more than `mergeFill`, and by how much); for every pair, the bytes of the page of two keys
against the object of the single-key page (whether it could have been made in place); in the census, the single-key pages by the size
of their object (what keeping them one class larger would cost). Cases: the seven structured kinds, single-value and natural, 4,096
and 16,384 keys, one steady-state churn cycle (TestProbe).

**Prediction:**
1. At least 70 % of the `mergeUp` calls start from a page left with one key (the pairs losing a key).
2. At least 80 % of the refusals are by size, and the merged page would need at least twice `mergeFill` in the median: the
   siblings hold some ten entries each, so the tries are not near misses.
3. In 30 to 60 % of the pairs the page of two keys would fit the object of the single-key page.
4. Keeping every single-key page one class larger (the most a hysteresis against shrinking could cost) costs at most 10 bytes a key.

**Run:** PC/WSL, 2026-10-10 17:17 to 17:20, one process a kind (counts only); `bench/results-layout/write-2026-10/events-merge-pair.txt`,
summary `merge-pair-summary.md` (`evsum.py`).

**Result** (one steady-state churn cycle; single-value / natural at 4,096 keys unless said):

| kind | merge ups a 1,000 writes | from 1 entry left | refused: >12 children / a child no page / too big | too big: median need, share under 1.5x | pairs a 1,000 writes | pair fits the old object | single-key pages one class larger, B a key |
|---|--:|--:|---|---|--:|--:|--:|
| str | 54 / 58 | 37 / 39 % | 8/0/92 / 0/32/58 % | 1.8x, 8 % | 20 / 23 | 88 / 95 % | 1.5 / 17.9 |
| email (16,384) | 28 (462) / 47 | 23 (39) % | 73/0/27 (94/0/6) % | 1.6x, 33 % | 6 (182) / 11 | 93 (93) / 97 % | 0.9 (23.1) / 31.1 |
| url | 280 / 70 | 32 / 30 % | 23/28/49 / 16/40/41 % | 2.2x, 12 % | 89 / 21 | 70 / 83 % | 25.0 / 40.0 |
| path | 232 / 60 | 31 / 32 % | 28/35/37 / 20/44/35 % | 2.2x, 10 % | 70 / 19 | 79 / 88 % | 15.2 / 28.8 |
| street | 113 / 57 | 34 / 35 % | 28/22/50 / 23/33/35 % | 2.1x, 3 % | 39 / 20 | 97 / 98 % | 3.4 / 5.9 |
| dirs | 201 / 76 | 33 / 32 % | 35/28/37 / 29/36/34 % | 2.1x, 9 % | 66 / 24 | 84 / 92 % | 9.0 / 15.1 |
| links | 168 / 6 | 34 / 27 % | 51/11/38 / 18/60/15 % | 2.3x, 4 % | 57 / 2 | 91 / 97 % | 5.6 / 49.7 |

(The rest of the refusals, more than 64 values, is 0 to 13 %.) Every `mergeUp` tries exactly one node (1.00 tries a merge up): the
lowest one refuses and the ones above are not tried. Where the time of a try goes (CPU profiles above, single-value 4,096): the
size check `mergeFits` 6 to 10 % of a churn, the second descent from the root and the rest of `mergeUp` 2 to 8 %.

**Against the prediction:**
1. At least 70 % of the merge ups from a page left with one key: missed. Only 22 to 39 % are; 61 to 78 % start from a page left
   with two entries.
2. At least 80 % of the refusals by size: missed (6 to 92 %; the node with more than 12 children and the child that is a node are
   as frequent). Not near misses, at least twice `mergeFill` in the median: about hit (1.6 to 4.0 times, 2.1 to 2.3 for most;
   under 1.5 times only 1 to 12 %, links natural and email 4K about 30 %).
3. The page of two keys fits the single-key page's object in 30 to 60 % of the pairs: missed, far more: 70 to 99 %.
4. Keeping every single-key page one class larger costs at most 10 bytes a key: missed for url, path, email at 16,384 and all
   natural (12 to 50 bytes a key); hit for str, street, dirs, links single-value (0.7 to 9).

**What it says (no decision, for the user):**
- the merge tries fail for reasons that are known before any page is read: a node of more than 12 children (refused at once, but
  only after the second descent from the root), a child that is a node (found while walking the children); the walk over all
  sibling pages (`mergeFits`, 6 to 10 % of a churn) is spent on a size that is twice the limit in the median. A trigger of one
  entry left instead of two would take away 61 to 78 % of the tries (whether merges that matter would be lost then is not counted);
  remembering the path of the removal would take away the second descent;
- the pair could be made in place, in the object the single-key page already has, in 70 to 99 % of the cases: that takes away most
  of the pendulum's objects without any memory, unlike keeping single-key pages one class larger (up to 50 bytes a key).
