# Where leaf pages stand, measured without bias (2026-09-28)

The first measurement with rtcompare v0.7.0, which fixed the bias against
candidate A ([rtcompare#111](https://github.com/TomTonic/rtcompare/issues/111)).
It replaces every head-to-head figure from 64K keys up in the directories
next to this one.

## Setup

- Ryzen 9 7900, WSL2, go1.27.1, rtcompare v0.7.0, bench refactored onto
  `multiproc`, `Combine` and `workload` (see `../../README.md`).
- Dev precision (41 samples, 2 A/A runs; builds 21 samples, 2 A/A runs). 5-10
  processes per scenario (`main/`: 5-6), each with its own heap perturbation
  and alternating build order.
- **A is always main's Ordered (`7c7e6a2`).** B is a leaf-pages commit copied
  in as `baseline`, or one of the other candidates in `main/`.
- The script is `run.sh`; it took 22:14 to 06:18.
- A/A check beforehand with identical code, u64 and str with several values
  per key at 256K keys: point lookup, range query and str churn pooled to
  0.99-1.01. u64 churn pooled to 1.00, but its processes scattered by about
  ±8 points.

Figures below are **B's speed relative to main** (1 − Δ; above 1 = B faster):
the range over the key kinds and the median in brackets. For `main/`, it is
main's speed relative to the other candidate.

## 1. Current leaf-pages (`e1f6ad3`: pages for single-value keys, fallback, main's leaves) vs main

All 7 key kinds.

| values | n | point | range | prefix | churn | build |
|---|---|---|---|---|---|---|
| unique | 4K | 0.66–0.89 (0.85) | 3.14–4.35 (3.79) | 0.90–5.62 (1.26) | 0.75–0.94 (0.91) | 0.67–0.97 (0.71) |
| unique | 16K | 0.67–0.99 (0.85) | 3.16–5.61 (3.64) | 0.83–5.64 (1.96) | 0.70–1.01 (0.92) | 0.60–0.79 (0.73) |
| unique | 256K | 0.98–1.77 (1.16) | 2.62–6.88 (3.49) | 1.20–5.39 (3.13) | 0.87–1.43 (1.03) | – |
| unique | 1M | 0.75–1.21 (0.94) | 2.83–4.38 (3.48) | 1.18–4.72 (3.33) | 0.88–1.21 (0.97) | – |
| multi | 4K | 0.94–0.97 (0.96) | 0.97–1.02 (0.99) | 0.96–1.00 (0.99) | 0.93–0.96 (0.94) | 0.83–0.92 (0.90) |
| multi | 16K | 0.94–0.98 (0.96) | 0.99–1.00 (0.99) | 0.97–1.00 (0.99) | 0.90–0.99 (0.95) | 0.88–0.95 (0.91) |
| multi | 256K | 0.93–0.99 (0.97) | 0.99–1.00 (1.00) | 0.99–1.01 (1.00) | 0.93–1.02 (0.98) | – |
| multi | 1M | 0.95–0.99 (0.97) | 0.99–1.00 (1.00) | 0.99–1.00 (1.00) | 0.95–1.00 (0.99) | – |

**Memory at 1M:**
- **multi:** identical to main for every kind; the fallback leaves main's tree.
- **unique:** per key 16 instead of 73 B (u64), 44/101 (str), 95/117 (uuid), 72/100 (email), 117/154 (url), 105/158 (path), 48/96 (street); GC time per cycle falls by 54-98%.

Where unique still loses:
- **Point:** uuid and email at 4K/16K (0.66-0.82), uuid at 1M (0.75), u64 at 16K (0.71).
- **Churn:** u64 at 16K (0.70) and email at 4K/16K (0.75).
- **Build:** 0.60-0.97 up to 16K.

## 2. The page types dropped earlier vs main

5 kinds (u64, str, uuid, path, street); 4K and 256K keys.

| variant | values | n | point | range | prefix | churn | build |
|---|---|---|---|---|---|---|---|
| `8052a18` single-value pages, range nodes, no fallback | unique | 4K | 0.68–0.90 (0.84) | 2.99–4.17 (3.74) | 0.90–5.01 (1.48) | 0.78–0.95 (0.91) | 0.67–0.97 (0.70) |
| | unique | 256K | 0.96–1.93 (1.39) | 2.63–7.41 (3.94) | 1.77–5.69 (3.03) | 0.96–1.50 (1.04) | – |
| | multi | 4K | 0.77–0.87 (0.81) | 0.94–1.25 (0.98) | 0.89–1.20 (0.95) | 0.70–0.81 (0.73) | 0.51–0.70 (0.61) |
| | multi | 256K | 0.69–0.85 (0.76) | 0.86–0.93 (0.89) | 0.86–0.94 (0.92) | 0.83–0.91 (0.86) | – |
| `afd1137` range nodes, K and S pages, keys with several values in pages | unique | 4K | 0.65–0.89 (0.78) | 2.68–3.22 (2.92) | 0.57–4.83 (1.12) | 0.76–0.93 (0.87) | 0.51–0.64 (0.51) |
| | unique | 256K | 0.98–1.60 (1.39) | 2.39–4.87 (3.26) | 1.03–5.34 (2.72) | 0.97–1.35 (1.03) | – |
| | multi | 4K | 0.73–0.88 (0.84) | 1.19–1.63 (1.22) | 0.45–1.45 (0.81) | 0.57–0.77 (0.61) | 0.38–0.55 (0.46) |
| | multi | 256K | **0.95–1.25 (1.18)** | **1.43–1.66 (1.48)** | 1.05–1.77 (1.48) | 0.89–0.98 (0.94) | – |
| `7157f4b` K pages next to U8-1, U8-n and S, no range nodes | unique | 4K | 0.66–0.91 (0.86) | 0.81–3.06 (1.46) | 0.60–1.95 (0.87) | 0.65–1.07 (0.70) | 0.41–1.16 (0.48) |
| | unique | 256K | 1.15–1.87 (1.31) | 1.79–3.16 (1.87) | 1.37–1.86 (1.80) | 0.98–1.35 (1.07) | – |
| | multi | 4K | 0.78–0.92 (0.87) | 0.88–1.31 (1.06) | 0.68–1.27 (0.88) | 0.62–0.81 (0.66) | 0.44–0.69 (0.53) |
| | multi | 256K | 1.04–1.16 (1.12) | 1.20–1.34 (1.28) | 1.14–1.31 (1.18) | 0.92–1.00 (0.96) | – |
| `527d9b7` U8-1, U8-n and S pages | unique | 4K | 0.67–0.94 (0.85) | 0.83–3.27 (1.48) | 0.61–2.03 (0.88) | 0.66–1.11 (0.69) | 0.41–1.19 (0.47) |
| | unique | 256K | 1.14–1.74 (1.31) | 1.78–3.46 (1.88) | 1.37–1.88 (1.76) | 0.96–1.28 (1.07) | – |
| | multi | 4K | 0.79–0.92 (0.87) | 0.90–1.34 (1.08) | 0.68–1.31 (0.89) | 0.62–0.81 (0.67) | 0.44–0.71 (0.53) |
| | multi | 256K | 1.04–1.17 (1.12) | 1.22–1.35 (1.29) | 1.14–1.32 (1.21) | 0.95–1.02 (0.97) | – |

Memory at 256K keys, bytes per key, main → variant (1 round):

| variant | values | u64 | str | uuid | path | street |
|---|---|---|---|---|---|---|
| `8052a18` | multi | 154 → 152 | 176 → 196 | 193 → 221 | 233 → 259 | 119 → 116 |
| `8052a18` | unique | 79 → 16 | 101 → 45 | 118 → 92 | 158 → 106 | 96 → 48 |
| `afd1137` | multi | 155 → **109** | 176 → **115** | 193 → **147** | 233 → **153** | 120 → **59** |
| `afd1137` | unique | 79 → 15 | 101 → 44 | 118 → 90 | 158 → 104 | 96 → 48 |
| `7157f4b` | multi | 154 → 115 | 176 → 131 | 193 → 153 | 233 → 165 | 120 → 62 |
| `7157f4b` | unique | 79 → 21 | 101 → 43 | 118 → **65** | 158 → **76** | 96 → **34** |
| `527d9b7` | multi | 154 → 115 | 176 → 131 | 193 → 153 | 233 → 165 | 119 → 62 |
| `527d9b7` | unique | 79 → 21 | 101 → 43 | 118 → 65 | 157 → 76 | 96 → 34 |

## 3. The band: main vs the other candidates

All 7 kinds; main's speed relative to the other candidate.

| B | values | n | point | range | prefix | churn | build |
|---|---|---|---|---|---|---|---|
| btree-sets | multi | 4K | 1.88–4.48 (2.87) | 1.89–2.76 (2.21) | 1.46–2.40 (2.20) | 1.49–3.62 (2.31) | 1.54–3.10 (2.27) |
| | multi | 1M | 2.21–4.59 (2.60) | 2.46–3.55 (2.62) | 2.56–2.74 (2.68) | 2.10–3.13 (2.26) | – |
| hashed | multi | 4K | 0.39–1.04 (0.67) | 15–23× | 7–500× | 0.43–1.07 (0.69) | 0.47–1.09 (0.72) |
| | multi | 1M | 0.46–1.03 (0.53) | – | – | 0.67–0.96 (0.73) | – |
| map-sets | multi | 4K | 0.93–2.54 (1.53) | 16–24× | 8–500× | 0.51–1.24 (0.78) | 0.54–1.19 (0.88) |
| | multi | 1M | 0.92–2.29 (1.19) | – | – | 0.73–1.08 (0.82) | – |
| btree-map | unique | 4K | 1.44–5.75 (2.67) | **0.31–0.47 (0.34)** | 0.32–1.82 (0.72) | 0.96–2.83 (1.62) | 1.04–2.29 (1.56) |
| | unique | 256K | 0.94–2.30 (1.19) | **0.32–0.58 (0.44)** | 0.29–1.10 (0.53) | 0.91–1.67 (1.49) | – |
| | unique | 1M | 1.64–2.99 (1.91) | **0.54–0.69 (0.55)** | 0.42–1.64 (0.76) | 1.61–2.23 (1.75) | – |

At 16K and 256K the figures lie between those of 4K and 1M; all rows are in
`main/speed-summary.md`.

Memory at 1M keys:
- **multi:** main needs 119-232 B per key, hashed 114-242, the hand-written sets 273-425.
- **unique:** main needs 73-157 B per key, btree-map 37-102, so main needs 1.5 to 2 times as much.

**Current leaf-pages vs btree-map with unique values** is derived: main vs btree-map times leaf-pages vs main, both from this night.
- **Point:** 1.06–4.86.
- **Range:** 1.13–3.04, except path at 4K/16K (0.95–0.99).
- **Churn:** 1.10–2.70, except path (0.88–0.92).
- **Build:** 0.74–2.22.
- **Memory:** below btree-map for u64 and str, 4% above for path and street, 15–26% above for uuid, email and url.

## Findings

1. **The bias hid most of what pages do for large maps.** Every page type is faster than main from 256K keys up in point lookups, range queries and prefix searches, except the variant without fallback (`8052a18`) with several values per key.
2. **Unique values:** the current leaf-pages is the best variant. Range queries are 2.6–6.9× as fast as main, memory is 19–78% lower, and GC 54–98% lower. Against btree-map (derived) it now wins almost everywhere, including range queries. What is left are small indexes: point lookups at 4K/16K for keys longer than 16 bytes (uuid, email: K-page tails outside the page), u64 at 16K, and builds.
3. **Several values per key:** the fallback makes the current leaf-pages as fast as main (0.93–1.02 in point, range and churn) with the same memory. The old types that keep such keys in pages (S/U8-n, with or without range nodes) are clearly better at 256K: points 1.04–1.25, range queries 1.2–1.7, 20–51% less memory and 50–77% less GC. At 4K they lose clearly: churn 0.57–0.81, build 0.38–0.71, point lookups 0.73–0.92. **For several values per key, a size switch pays off after all:** leaves while a map or subtree is small, pages above a threshold between 4K and 256K that remains to be measured.
4. **S pages are the most compact** for unique uuid and path (65/76 against 90/104 B per key for K pages), but slower in range queries (1.8–1.9× instead of 3.3–3.9× main's).
