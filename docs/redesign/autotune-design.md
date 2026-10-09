# Autotune of the keys per page: concept (2026-10-07, not implemented)

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
