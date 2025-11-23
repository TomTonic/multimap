# ART Node Design Decision Matrix (Power-of-Two Objects)

This document summarizes the **decision matrix** for ART nodes optimized for **object size** (powers of two).

Design rules applied:

- **Pack the maximum number of children** for each object size.
- **Add a 256-bit bitmap** (implemented as `[4]uint64`) **only when max children > 32**.
- Lookup strategies:
  - **Node64B / Node128B / Node256B**: *no bitmap*, *UNSORTED*, **linear unrolled**.
  - **Node512B / Node1024B**: *bitmap present*, *SORTED `firstKeyByte[]`*, **binary search**.
  - **FullNode (64 B + 2048 B)**: *bitmap present*, **direct index** (no sorting).

---

## Strategy Matrix — *Bitmap + sorted keys[] + binary search only where it pays*

| **Type**      | **Exact object size** | **Max children packed** | **Bitmap?** | **Sorting?** | **Search Strategy** | **Waste (padding / size, %)** |
|---------------|------------------------|--------------------------|-------------|--------------|------------------------|-------------------------------|
| **Node64**    | 64 bytes               | 4                        | No          | No           | **Linear, unrolled** (4 checks) | **4 / 64 bytes ≈ 6.25%** |
| **Node128**   | 128 bytes              | 11                       | No          | No           | **Linear, unrolled** (~3×4 checks) | **5 / 128 bytes ≈ 3.91%** |
| **Node256**   | 256 bytes              | 25                       | No          | No           | **Linear, unrolled** (3×8 + 1) | **7 / 256 bytes ≈ 2.73%** |
| **Node512**   | 512 bytes              | 50                       | **Yes**     | **Yes**      | **Bitmap → binary search** (≈ log₂50 ≈ 6 steps) | **6 / 512 ≈ 1.17%** |
| **Node1024**  | 1024 bytes             | 107                      | **Yes**     | **Yes**      | **Bitmap → binary search** (≈ log₂107 ≈ 7 steps) | **5 / 1024 ≈ 0.49%** |
| **FullNode**  | **64 bytes** node + **2048 bytes** external array | 256 (via external array) | **Yes** | n/a          | **Bitmap → direct index** (`arr[idx]`) | **Node: 0 / 64 = 0%**; **Array: 0 / 2048 = 0%** |

---

## Why linear (unrolled) for small nodes (64B / 128B / 256B)

For these sizes, the entire `firstKeyByte[]` (and often the first segment of `child[]`) resides within **one or two cache lines** of the node:

- Sequential loads are **prefetch-friendly** and **branch-lean**.
- You can further cut work per lookup using the **branchless, lane-parallel compare** on 64-bit chunks (**`hasAnyMatch64`**):
  - Broadcast the search byte across all lanes.
  - XOR with the chunk to produce `0x00` on matches.
  - Detect any zero byte using a constant-time mask.
- Because the **key bytes already sit in the node’s cache line(s)**, the **expected number of comparisons is lower** than for random access patterns, and **linear + unrolled** tends to outperform binary search on these small arrays.

---

## Why bitmap + binary search for larger nodes (512B / 1024B)

- When **max children > 32**, a **256-bit bitmap** gives an **O(1) negative fast path** (reject absent first key bytes without touching `firstKeyByte[]`).
- With **sorted `firstKeyByte[]`**, **binary search** reduces comparisons (≈ 6–7 steps for 50/107 entries) and **curbs branch mispredictions** compared to broader linear scans.

---

## FullNode (64B + 2048B) — direct index, no sorting

- Node object remains **one cache line** (64 B) and carries the **bitmap** plus a pointer to the external **`[256]*Node` array** (2048 B).
- Lookup: **bitmap check** → **`arr[idx]`** for the hit. The external array line is touched **only** when needed; **no sorting** required.

---

## Allocator, Padding, and GC considerations

- **Cache-line friendliness:** Align node objects so the hot fields used during lookup (`meta`, `numChildren`, `localPrefix`, `firstKeyByte[]`, and nearby `child` pointers) fall into the same cache line (typically 64 bytes). Keeping these fields in the first cache line reduces loads, improves prefetch behavior, and lowers branch misprediction/latency on the critical lookup path.

- **Go allocator size classes:** Go's memory allocator assigns objects to size classes and places same-sized objects in the same spans. Designing node object sizes to match common allocator size classes (the power-of-two / tuned classes used above: 64, 128, 256, 512, 1024) lets the runtime pack nodes densely into spans and reduces internal fragmentation and span churn.

- **Memory fragmentation:** Choosing exact (or intentionally rounded) object sizes reduces both internal and external fragmentation. When nodes fit their target size class, fewer spans are partially occupied and more nodes live closely together, improving locality. Conversely, drifting off class (by a few bytes) can push allocations into a larger class and reduce cache density.

- **Explicit padding rationale:** We add small padding fields intentionally to:
  - Guarantee that `child[]` and any pointer arrays are 8-byte aligned.
  - Keep the layout deterministic so the most frequently accessed bytes remain on the same cache line across platforms.
  - Ensure the total object size lands on the intended size-class boundary (so the allocator places it in the expected span).
  Explicit padding is preferable to relying on implicit compiler-added padding because it documents intent and keeps layout stable across Go versions.

- **GC scan cost and pointer locality:** Grouping pointer fields and keeping pointer-heavy structures near each other reduces GC scan overhead (and the cost of write barriers). If pointer words are scattered, GC scanning touches more cache lines and increases pause/scan cost. Where possible, keep hot, non-pointer lookup data (e.g., `firstKeyByte[]`) in the first cache line and push pointer arrays so the hot path minimizes pointer dereferences.

- **Trade-offs and guidance:**
  - Packing more children into a single node reduces indirections but increases per-node scan work for the GC and can increase allocation size.
  - Splitting very large fanout into an external array (as in `FullNode`) keeps the node header small and GC-friendly while allowing direct indexing when needed.
