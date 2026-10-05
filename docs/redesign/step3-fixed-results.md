# Step 3.5.1: the single-key page for fixed-size values, standalone

Written 2026-10-04. Design: [step3-fixed-design.md](step3-fixed-design.md). Words as in [GLOSSARY.md](GLOSSARY.md). Code: `internal/skpage/fixed.go`
(`Fixed`, generic methods since Go 1.27), `fixed_ptr_gen.go` (166 types, from `internal/skpage/gen`), tests `fixed_test.go`; the benchmark `internal/art/fixed_bench_test.go`
with the entries of `street` and `dirs` (`internal/art/testdata`, from `bench/cmd/skmodel -histogram`). Raw: `bench/results-layout/step3-fixed/pc/` (PC, Ryzen 9 7900, Windows native,
`-test.benchtime 500ms -test.count 5`, median).

## What was built

- `Fixed`: the head of `Page`, then remainder, padding to the alignment of `T`, `[n]T`. `NewFixed`, `BuildFixed`, `Has`, `Add`, `Remove`, `Each`, `Values`, `Prepend`, `Supported`, `MaxRemainderFixed`.
  Pointer-free `T` of 1 to 16 bytes: a `[N]uint64`; one word with a pointer: one of the 166 typed objects (class × words in front), generated. `Page` and `Fixed` share one `head`.
- **Gates:** `internal/skpage` 100 %, `-race`, `FuzzFixed` 60 s (1.1 million runs), lint 0. Tests among others: a model test for 9 types of value and 15 remainder lengths (alignment of every kind), and **a GC test
  for all 166 layouts**: every record the page holds survives a collection, a removed one does not, and a record that only the *bytes of the remainder* point to is collected
  (the remainder is not scanned).
- **The rule for the way back** (user's correction): `skpage.BackFits(rem, valueBytes)`: the values of a key go back from the value overflow into a page when they take at most half the room the page has for
  them (512 less head and remainder). `internal/art` uses it for strings now (it was 256 bytes of content including the remainder, so a key with a remainder of 300 bytes never came back); test
  `TestStringBackWithLongRemainder`.
- **Shrinking with hysteresis** (a decision made on a measurement, below): a page shrinks only when its values fill at most half of the smaller class (a class that holds twice as many as are left); it grows to
  the smallest class that holds one more.

## Measured against the prediction

| | prediction | measured |
|---|---|---|
| memory, `uint64`, per key held by both (street real / dirs real) | page = flat leaf (model 41.1 / 53.0) | **39.4 against 39.6; 50.6 against 49.7 B** |
| memory, `uint64`, one value per key | equal | **32.1 = 32.1; 35.7 against 34.1 B** |
| memory, `*rec` real | (not predicted against the typed leaf) | page **36.2 against 30.3 B (+20 %)**; dirs 45.2 against 41.0 (+10 %) |
| memory, `*rec`, one value per key | (the model said 32.1) | page **32.1 against 22.8 B (+41 %)**; dirs 35.6 against 26.8 (+33 %) |
| hit and read (adding a value that is there, summing the values) | within the noise | **0.95 to 1.02** (pointers, all keys, read: 1.06 to 1.11) |
| churn (add and remove a new value) | within the noise | `uint64`: **1.02 to 1.06**; `*rec`: **1.15** (hot) and **1.31 to 1.42** (all keys) |
| build (all values of every key, one by one) | within the noise | `uint64`: **1.17 to 1.22**; `*rec`: 0.95 to 1.00 |

Page / old leaf, time (hot: 4,096 keys in the cache; all: every key). Full table in `pc/bench.txt`.

## Findings

1. **A first version of `Remove` shrank at once** (`values fit a class of at most half the size`, the rule of the string page): a key at a class border was copied on every added and removed value. Measured churn:
   `uint64` 1.6 to 1.9 times the flat leaf, `*rec` 2 times. With the hysteresis above it is 1.0 to 1.06 for `uint64`. **The string page has the same rule and the same weakness** (3.3 called it "the worst case of the page");
   not changed yet, see the question below.
2. **The page is not smaller than the typed leaf for pointers, it is bigger.** The typed leaf is an object of 16, 24 or 32 bytes for a key with one value and a remainder of up to 10 bytes (head, two words of key
   area, one pointer: 24 bytes, a Go size class); the page starts at 32. That is the grid, which was chosen for the strings. 72 % of the `street` keys and 40 % of the `dirs` keys have one value and a remainder of up to 10
   bytes, so they would take 24 bytes in an object of that class and take 32 now: **+5.8 B/key on `street` (the typed leaf's 22.8 against 32.1 on one value per key is the same effect)**; in the whole tree about +5 % on `street` real
   and +14 % on one value per key.
3. **Churn of pointers** is 3 ns slower hot and 10 ns slower on all keys, build and read are equal. A tree operation costs about 500 ns at the full size (`churn`, step 3.4b), so this is about 2 % there; not investigated.
   The profile puts it on the first touch of the object (`Has`: a cache miss, the same for both) and the arithmetic of class and capacity (`capacityOf`, `classHolding`: a loop per call, which a table per class would take away).
4. **Build** for `uint64` is slower because the flat leaf grows by doubling to at least 64 bytes and the page by one class: more small objects when keys get their values one by one. In a map built from a bulk the page is
   cheaper (`BuildFixed` makes the page of the right class at once; the tree does not use it for that yet).

## Decisions of the user (2026-10-05)

- **No 24-byte class for now** ("erstmal keine"): the pages stay on the grid 32/64/128/256/384/512; the pointer maps are 5 to 14 % bigger than the typed leaf for the time being, a candidate for the profile after 0.8.
- **The string page gets the same hysteresis** (done, commit below; tests, race, both fuzz tests, lint green). One rule for both flavors.

The questions as asked:

- **A 24-byte class for the pages?** For one word of value and a remainder of up to 10 bytes a page of 24 bytes (3 words: head and remainder in two, the value in one) is what the typed leaf has now, and what the
  `uint64` page would take too (`flat leaf` starts at 32 for `uint64`, so no gain there). It brings the pointer page to the typed leaf's memory (-5.8 B/key on `street`) and costs: one class that is not on the grid 32/64/128/...,
  an object that straddles a cache line in one of four positions (24 B objects in a span), two more of the 166 types. Recommended only if memory of `*T` maps is a goal of 0.8; without it the pointer maps are about
  5 % (real) to 14 % (one value per key) bigger than today. The glossary and the string page stay on the grid either way.
- **The shrink rule of the string page** (`Page.Remove`): give it the same hysteresis? It changes a measured behavior of 3.3 and 3.4 (churn of strings gets better, memory of a map after many removals
  gets a little worse). Recommended, as the one rule for both flavors.
