# Autotune of the keys per page: concept (2026-10-07), measured and dropped (2026-10-10)

Words as in [GLOSSARY.md](GLOSSARY.md). Origin: the review of the approach ([review-2026-10.md](review-2026-10.md), E1 to E6) and the
user's decision of 2026-10-07. **This note fixes the concept; it is implemented later**, after the hysteresis (E6) is adopted.

## The problem it solves

Two forms of the same structure, measured on today's code:

| | one key per page (E1) | multi-key pages (today) |
|---|---|---|
| small maps (4,096 keys), writes | as fast as step 3.5, 25 to 45 % faster than today | slower |
| memory a key | about twice | 35 to 70 % less than `main`, `node-layout`, step 3.5 |
| reserved memory over a long churn (E3) | 1.3 to 3 times | the least of all candidates |
| scan of single-value maps (E4) | 1.3 to 2.1 times the time | faster |

A small map fits the caches, where the instructions per operation decide and memory is cheap; a large map lives in memory, where
bytes and cache misses decide. One fixed form is wrong for one of the two.

## The concept

1. **One number for the whole map:** the largest number of keys a page may hold, `maxKeys`. `maxKeys = 1` is the form of E1 (every key its
   own page; a different key that meets a page puts a node above it), the largest value is today's form (the page's own limits: 512
   bytes, 255 slots). Both ends work alone (E1 showed it); the border moves between them by this one number.
2. **The map counts its pages** (single-key and multi-key pages; the count changes where pages are made and freed) and derives `maxKeys`
   from it: **`maxKeys` grows linearly with the number of pages** (the factor is what the experiments of point 5 find; a floor of 1 and
   the page's own limits as the ceiling).
3. **Nothing is merged or split artificially.** When `maxKeys` rises, the existing pages stay as they are; a page takes another key when
   one arrives and the limit allows it, so the pages filled from then on are fuller. When the map shrinks, `maxKeys` follows the count
   down; pages that hold more keys than the new limit stay until their keys go. The hysteresis between burst and merge (E6) applies at
   every limit.
4. **It can be switched off**, and `maxKeys` can be fixed: to find the right curve for a concrete context (the cache sizes of a
   production machine are not known in advance, several maps share them), and to measure both ends.
5. **The curve is found by experiments** (to be planned when it is built): the gate cells at 4K, 16K, 65K and the full corpus, the key
   kinds of E4, against today, E1 and `main`; speed, memory and reserved memory (E3's probe), on the PC and the M1.

## What is not decided

- the factor of the linear rule, and whether it counts pages or the bytes the pages take (a key takes 27 to 114 block bytes by kind, E4);
- how `maxKeys` is set from outside (an option of `NewOrdered`, an environment variable for the experiments, or both);
- whether the limit, once the map is large, should ever stop growing before the page's own limits.

---

## Experiment exp-maxkeys: the curve, first measurement (2026-10-09, before the run)

Why now: two days went into the search inside multi-key pages at 4,096 keys (page-search-design.md, options 1 and 2, then a
hash-free SWAR search): the size where the autotune would hold one key a page anyway. Whether the page search matters at all
depends on where multi-key pages live, which is what the curve decides; the page search is frozen until then.

**Code** (branch `exp-maxkeys`, c69349e, from c3367b0: the search is today's `locate`, no fingerprints): `Ordered.SetMaxKeys(k)`
(experiment only). A new key bursts a page that holds k keys; a key outside the common prefix of such a page gets a node above it
instead of `Widen`; a burst builds no page of more than k keys; a merge needs at most k/2 keys (the hysteresis of `mergeFill`, in
keys; for k = 2 there are no merges); k = 1 is E1 (no multi-key pages, no merge attempts). The bench has the candidates
`ordered-mk1`, `ordered-mk2`, `ordered-mk4`; `ordered` (A) has no limit but the page's own (5 to 16 keys a page by kind).

**Runs** (one command each machine; the same arguments on both):
- M1 (job `mx1`) and PC run 1: `-suite dev -keys u64,street,dirs,url -values natural,single-value -sizes 4096,16384,65536 -ops
  valuesFor,valuesBetween,churn,build -memn 262144 -memrounds 3 -vs baseline,btree-sets,btree-map,ordered-mk1,ordered-mk2,ordered-mk4
  -minprocs 4 -maxprocs 8`, baseline = `main` (cc47992). About 4.5 to 5 hours each.
- PC run 2, the large end: `-keys u64,street,url -sizes 212449 -ops valuesFor,valuesBetween,churn -vs
  baseline,ordered-mk1,ordered-mk2,ordered-mk4 -skipmem`. About 1.5 to 2 hours.

**Prediction** (the smoke run of 22:13, street single-value 4,096, low precision, gave: valuesFor ordered 0.83 of mk1, 0.90 of mk2,
0.94 of mk4; churn ordered 0.93 of mk1, 1.38 of mk2, 1.36 of mk4; heap B a key ordered 33, mk4 46, mk2 54, mk1 70, `main` 105):
- valuesFor, 4K and 16K: mk1 the fastest ordered variant, 1.1 to 1.25 times ordered on street, dirs, url, 1.4 to 1.6 on u64 (the
  lookup probe: one key a page is 0.96 to 0.97 of `main` on street and dirs, 0.80 to 0.85 on u64); mk2 and mk4 in between, in order.
  65K: the variants within ±10 % of each other.
- churn and build, 4K and 16K: mk1 1.0 to 1.15 times ordered (E1: as fast as step 3.5); **mk2 and mk4 slower than ordered by 20 to
  40 %** (the smoke run; my explanation, not yet checked: every new key at a full page bursts it and rebuilds the subtree, and a
  page of 2 never merges). 65K: within ±10 %.
- valuesBetween (100 keys): single-value mk1 0.5 to 0.75 of ordered (E4: the scan of single-value maps takes 1.3 to 2.1 times with
  one key a page), mk2 and mk4 in between; natural closer (the values of a key are read together either way).
- memory, heap B a key at 262,144 (or the whole corpus): ordered as today (u64 natural 101, street natural 54), mk4 +30 to +50 %,
  mk2 +50 to +80 %, mk1 about twice; all below `main` except mk1 on u64 natural, where it comes near.
- the band of the credo: every variant beats btree-sets / btree-map in point queries and churn at every size.

**Decision it prepares** (with the user, 2026-10-10): the rule that gives `maxKeys` from the number of pages, or a different
answer if the curve says so (for example: mk2 and mk4 never pay, then the rule is a switch between 1 and the page's own limit).

## Experiment exp-maxkeys: result (2026-10-10; PC 22:37 to 05:14, M1 job mx1 the same night)

Code at 99b00af (c69349e with Go 1.27.2 and rtcompare v0.8.1; rtcompare's own code unchanged). Results: `bench/results-layout/maxkeys/`
(PC: `pc-mx1`, `pc-mx2`; M1: branch `arm-results`, `mx1`); the tables "speed vs main" of every cell were made with `vs-main.py`
(`pc-mx1/vs-main.txt`, `pc-mx2/vs-main.txt`, `m1-mx1-vs-main.txt`): each variant's speed relative to `main` (above 1 faster),
derived from the pairs ordered/main and ordered/variant. Many cells stopped at 8 processes before the precision asked for (the
`*` in those tables); the differences below are larger than that.

**PC and M1 agree** in every finding (the M1 ratios differ by a few points, never in direction).

**Speed against `main`, PC, single-value (natural is alike; u64 below):**

| | n | ordered | mk1 | mk2 | mk4 |
|---|--:|--:|--:|--:|--:|
| valuesFor street / dirs / url | 4,096 | 0.81 / 0.85 / 0.84 | **0.99 / 0.96 / 1.00** | 0.89 / 0.89 / 0.88 | 0.87 / 0.88 / 0.86 |
| | 65,536 | **1.05 / 1.07 / 1.15** | 1.03 / 1.04 / 1.14 | 0.99 / 1.04 / 1.10 | 0.99 / 1.05 / 1.10 |
| | 212,449 | **1.48** / – / **1.40** | 1.06 / – / 1.21 | 1.17 / – / 1.25 | 1.19 / – / 1.30 |
| churn street / dirs / url | 4,096 | 0.81 / 0.92 / 0.81 | **0.86 / 0.91 / 0.93** | 0.58 / 0.65 / 0.65 | 0.59 / 0.71 / 0.70 |
| | 65,536 | **1.14 / 1.26 / 1.05** | 0.96 / 1.05 / 1.04 | 0.77 / 0.87 / 0.84 | 0.81 / 0.91 / 0.89 |
| | 212,449 | **1.28** / – / **1.08** | 0.96 / – / 1.06 | 0.87 / – / 0.91 | 0.92 / – / 0.92 |
| valuesBetween street / dirs / url | 4,096 | **2.83 / 1.92 / 1.60** | 0.66 / 0.67 / 0.67 | 0.84 / 0.86 / 0.84 | 1.12 / 1.17 / 1.13 |
| | 212,449 | **3.70** / – / **2.01** | 0.92 / – / 0.91 | 1.15 / – / 1.09 | 1.51 / – / 1.44 |
| valuesFor u64 | 4,096 / 65,536 / 212,449 | 0.51 / 0.71 / 1.02 | **0.86 / 0.85 / 1.20** | 0.85 / 0.75 / 1.07 | 0.85 / 0.71 / 0.97 |
| churn u64 | 4,096 / 65,536 / 212,449 | 0.58 / 0.83 / **1.10** | **0.84 / 0.86** / 1.03 | 0.67 / 0.71 / 0.82 | 0.68 / 0.81 / 0.83 |

**Memory** (heap B a key at 262,144 keys or the whole corpus; the same on both machines):

| | ordered | mk4 | mk2 | mk1 | `main` | btree-map / btree-sets |
|---|--:|--:|--:|--:|--:|--:|
| single-value street / dirs / u64 / url | **27 / 41 / 24 / 72** | 40 / 50 / 39 / 83 | 51 / 58 / 50 / 92 | 67 / 76 / 53 / 108 | 104 / 149 / 87 / 163 | 55 / 96 / 45 / 109 |
| natural street / dirs / u64 / url | **54 / 74 / 102 / 150** | 65 / 83 / 114 / 157 | 75 / 90 / 122 / 167 | 89 / 106 / 125 / 182 | 127 / 179 / 163 / 239 | 285 / 335 / 352 / 416 |

**Against the prediction:**
- valuesFor: mk1 the fastest ordered variant at 4K and 16K: **met** (1.13 to 1.22 times ordered on text keys at 4K, 1.07 to 1.12 at
  16K; u64 1.69 at 4K, but only 1.06 at 16K, predicted 1.4 to 1.6). 65K within ±10 %: met for text keys; u64 1.20 (missed).
- churn and build, mk1 against ordered at 4K and 16K: predicted 1.0 to 1.15, **more**: 1.06 to 1.16 single-value text, 1.3 to 1.45 natural
  street and u64.
- **mk2 and mk4 slower than ordered in writes: met for single-value text keys** (28 to 30 % at 4K), not for natural and u64 (equal).
  They are slower than mk1 everywhere, and slower than or equal to ordered everywhere except valuesFor at 4K and 16K. Explanation (not
  checked): a page at its limit bursts with every new key, so a small limit means a burst for nearly every insert, and a limit of 2 never
  merges.
- valuesBetween, mk1 against ordered: predicted 0.5 to 0.75, **missed, much worse**: 0.23 to 0.42 (single-value; mk1 is slower than `main`
  in ranges, 0.62 to 0.78). E4 measured the full scan, not 100-key ranges, which carry the descent and the bounds for every one-key page.
- memory: mk1 about twice ordered (single-value 2.2 to 2.5 times, natural 1.2 to 1.65 times), every variant below `main` (met); mk4 and
  mk2 as predicted on street, more on u64.
- the band: every variant beats the B-trees in point queries and in the writes of the natural profile (met); **single-value writes are
  not in the band at small sizes, with any variant**: against btree-map (PC), churn at 4K ordered 0.88 to 1.01 and mk1 0.88 to 1.07, build at
  4K ordered 0.77 to 0.94 and mk1 0.88 to 1.09 (dirs loses with both, url with ordered); at 65K ordered 1.00 to 1.27, mk1 loses on dirs (0.91).

**What follows** (for the decision with the user):
1. Limits between 1 and the page's own are never the best choice for any operation at any size: the autotune is a **switch** between
   one key a page and the full page, not a linear curve.
2. The switch point for point queries and writes is between 16K and 65K keys for text keys (65K: the full page is as fast as one key a
   page or faster, and from there it pulls ahead: 212K street valuesFor 1.48 against 1.06); for u64 lookups beyond 212K.
3. Its price below the switch: ranges 3 to 4 times slower (one key a page is below `main` in ranges) and twice the memory, for 5 to 45 %
   faster point queries and writes. Whether that pays is a question of the workload (the credo asks for both).
4. The writes of single-value maps at small sizes are below btree-map with either form: that gap is not the page form's, it is the
   write path of the tree (a separate question).
5. One anomaly not understood: valuesBetween u64 single-value at 4,096 keys, ordered 3.69 and 3.54 times `main` on both machines but 0.94
   and 0.86 at 16,384 (`main` is slow at 4K there, not ordered fast).

## Decision (user, 2026-10-10): B, always the full page; the autotune is not built

Asked where A (the switch: one key a page up to some 32K keys, the full page above) and B (always the full page) stand against the
B-trees, the user chose **B**. Against the B-trees (PC, see the result above): both beat them in every point query and in every
write of the natural profile; in the single-value writes of small maps both lose on dirs and url (A 0.88 to 1.0, B 0.77 to 0.93); in
the single-value ranges `btree-map` wins against both, but A falls to 0.19 to 0.29 of it at 4K and 16K keys, B stays at 0.5 to 0.9
(u64 0.28 to 0.40 at 16K and 65K). A would trade a large loss in ranges and twice the memory for 5 to 45 % in point queries and writes
of small maps; B keeps ranges and memory and leaves the gap to `main` in point queries and writes of small maps (credo 3) to the two
paths that cost there: the search in the page and the write path of the tree. `hashed` was not in the run (the bench compares it in
the natural profile only); A and B both beat it in ranges, where it scans every key.

The branch `exp-maxkeys` stays as the record of the experiment; `SetMaxKeys` does not go into the library.
