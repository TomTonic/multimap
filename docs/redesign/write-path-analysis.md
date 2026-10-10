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
