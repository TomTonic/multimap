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
