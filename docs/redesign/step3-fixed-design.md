# Step 3.5: the single-key page for fixed-size values (`uint64`, `*T`)

Status: **design, decided 2026-10-04 (user)**: (1) the multi-key pages are **off, and cut out of `internal/art`** in 3.5 (they are redesigned later with what has been learned by then; if there is nothing, the old implementation is taken from the tag `before-mk-pages-removal`); (2) pointer pages only for one-word `T`, bigger inline is an optimization for later; (3) values aligned to `T`, compared as `T`. Sections 6 and 8 keep the numbers that led there; section 7 gets the first step below. Words as in [GLOSSARY.md](GLOSSARY.md). Numbers: `bench/cmd/skmodel`
(`-values words|pointers`, `-ovbytes`), `bench/cmd/ovbench` (`-values words|pointers`), the measurements of step 2 (`bench/results-layout/step3-real`).

## 1. What is there today, for the three cases

| case | Go type of the values | object of a key with several values | key with one value | when they do not fit |
|---|---|---|---|---|
| `string -> {string}` | `string` | **single-key page** (`internal/skpage`, bytes with a length byte), done in 3.3/3.4 | same page | value overflow, `Set3[string]` |
| `string -> {uint64}` | `uint64` (8 B, no pointer) | **flat leaf**: head, remainder, `[n]T`, Go size classes 32, 48, 64, 96, 128, 192, 256, 384, 512 | a **multi-key page** (`internal/vpage`, step 2; the tree has it for every `T` of at most 8 bytes) | set leaf, `vset.Set[uint64]` |
| `string -> {*T}`, `uint64 -> {*T}` | `*T` (one word with a pointer) | **typed leaf**: the same layout in an object that is a Go type (value arrays of 1 to 16, key areas of 2 to 58 bytes), so that the garbage collector sees the pointers | no multi-key page: `*T` holds a pointer, a page holds words | set leaf, `vset.Set[*T]` |

So the layout of a single-key page for fixed-size values **exists**: it is the flat and the typed leaf. What 3.5 changes is what the plan of 3.1 wanted for all of them:
one object type of the glossary on **one grid** (32, 64, 128, 256, 384, 512), **one value overflow** (`Set3` for every `T`, `vset` gone from `Ordered`),
the nine-bit remainder (step 3.4b), and **flat.go, typed.go, `leaf[T,K]` and `vset` in the map leave the code** (Gate 3).

## 2. The layout: fixed width, no length byte, values aligned

`singleKeyHead (6 B) | remainder (r bytes) | padding to the alignment of T | values [n]T` (unchanged from the flat leaf; padding to 8 for pointers).

**Why not the layout of the string page** (a length byte before each value, which would let `uint64` reuse `skpage` as it is): the model of 3.1 on the real data
(`skmodel -values words` against `-values words-len`, pages of the grid 32..512, B per key):

| | `street` real | `dirs` real | `street` single-value | `dirs` single-value |
|---|--:|--:|--:|--:|
| 8-byte values with a length byte | 42.5 | 56.1 | 32.2 | 36.3 |
| fixed width, no length byte | 41.1 | 53.0 | 32.1 | 35.7 |

3 to 6 % of the page bytes on the real mix, nothing on one value per key. The pointer page needs the fixed width anyway (a pointer is at a word, the length byte
would break that); with it a value is read as a `T` at a known offset (`unsafe.Slice((*T)(p), n)`: aligned, so `-race` and `checkptr` are content), a scan is a
stride and not a walk, and the same code serves `uint64` and `*T`. Only the allocation differs (section 3).

Values are **unsorted**; a lookup of a value scans, adding checks for the value first (set semantics), removing moves the last value into the gap and clears the last
slot (the typed leaf does that today). Equality is `T`'s `==` on the value read as a `T`, not the bytes (padding).

## 3. Allocation: pointer-free and pointer-holding

- **Pointer-free `T`** (of 1 to 16 bytes, as `flatType` today): the object is `[N]uint64`, never scanned, as the string page is. One code path for the classes.
- **`T` that is one word with a pointer** (`*X`, `unsafe.Pointer`, a struct or array of one pointer): the object must be a Go type with its pointers marked. Words 0 to J-1 hold
  head and remainder (no pointer), words J on hold the values: `struct{ h [J]uint64; v [W-J]T }`, `W` = size class / 8 (4, 8, 16, 32, 48, 64) and `J` = 1 to `W-1` (at least one
  value slot). **166 types** (3 + 7 + 15 + 31 + 47 + 63), a switch like `allocTyped` has today, generated. (First version of this note: 42 types, from a remainder of at most 58 bytes
  taken over from the typed leaf. **Wrong, user's correction 2026-10-04:** the remainder is limited only by the page, see below.) The count does not depend on the size of the head
  (6 bytes today for historical reasons, 3 later): the head and the remainder fill `J` words whatever the head is.
- **The limit of the remainder** is what the page leaves: 512 - head - 8 (one value slot), i.e. 498 bytes with today's head. If the remainder is long, few values fit (a remainder
  of 480 bytes: 3 pointers), and when more come, the key becomes a value overflow with `Set3[T]`; that is the one rule for every length. A remainder above 498 bytes is held as a string by the
  value overflow, as for strings. There is no extra case for "long remainder".
- **Return from the value overflow** (a point to settle in the code, not new): the string page returns at 256 bytes of *content* including the remainder, so a key with a remainder of
  300 bytes could never return. The rule to build for both flavors: return when the **values** take at most half of the room the page has for them (512 - head - remainder).
- **Every other `T`** (an interface, a string-holding struct, more than 16 bytes): *not a page*; the key is a value overflow from its first value. Slow and large per key, but
  exact; **this narrows what `Ordered` is good for** (the typed leaf took pointer-holding `T` up to 64 bytes). Decision 2.

## 4. The value overflow: `Set3[T]` for every `T`

As in 3.4. `ovbench -values words|pointers` on the keys that overflow (more than (512 - 6 - r) / 8 values: `street` 767 keys with 195,348 values, `dirs` 446 with 81,670):

| | heap B/value | hit ns | miss ns | add + remove ns |
|---|--:|--:|--:|--:|
| `vset.Set[uint64]` today, `street` / `dirs` | 17.0 / 17.2 | 31 / 27 | 17 / 17 | 40 / 38 |
| `Set3[uint64]`, `street` / `dirs` | 16.3 / 16.1 | 29 / 25 | 15 / 15 | 34 / 32 |
| `vset.Set[*rec]` | 17.6 / 17.8 | 30 / 26 | 18 / 17 | 40 / 39 |
| `Set3[*rec]` | 18.0 / 17.8 | 29 / 25 | 16 / 14 | 35 / 33 |

The same as for strings: memory equal, 5 to 15 % faster. `Set3` is generic in `comparable`, so one form serves all `T`. `vset` stays for `Hashed`.

> **Corrections made while building (2026-10-04):** Go 1.27 has generic methods (the user's remark; I had said it had not), so `Fixed` is an untyped head with methods `p.Has[T](v)`, not a
> generic type. Growing takes the smallest class that holds one more, **shrinking has hysteresis** (only when the values fill at most half of the smaller class), because the first version, which shrank at once like the
> string page, copied a key at a class border on every added and removed value (churn 1.6 to 2 times the flat leaf). Results: [step3-fixed-results.md](step3-fixed-results.md).

## 5. Size classes, growth, return (as 3.3)

Smallest class that holds the first value; when full, the smallest class that holds one more (the classes double, so one step); shrink on removal when a class of at most half
the size holds the values; value overflow when the values do not fit 512 bytes; back into a page at **256 bytes of content or less** (hysteresis, `skpage.BackLimit`). Capacity of a
page of 8-byte values with a remainder of 5 bytes: 3, 7, 15, 31, 47, 63 values.

## 6. The multi-key pages (decision 1: off, and cut out)

The plan of 3.1 says: switch the multi-key pages off while the single-key page is measured, bring them back in step 4. For strings they never existed, so nothing changed.
**For `T` of at most 8 bytes the tree has them today** (step 2), and they carry the keys with one value. Switching them off in 3.5 costs (model, B per key whole tree:
byte nodes + pages + value sets):

| | tree of today (measured, step 2) | 3.5 with the multi-key pages **on** (predicted) | 3.5 with them **off** (predicted) |
|---|--:|--:|--:|
| `street` real, `uint64` | 82 | **83** | 89 |
| `dirs` real, `uint64` | 98 | **101** | 104 |
| `street` single-value, `uint64` | 32 | **32** | 67 |
| `dirs` single-value, `uint64` | 54 | **54** | 73 |
| `string -> {*T}` (no multi-key page for pointers either way) `street` / `dirs` real | - | 90 / 106 | 90 / 106 |
| same, single-value | - | 67 / 73 | 67 / 73 |

(*Off*: `skmodel -nodes -values words -ovbytes 16.3`; *on*: today's measurement plus the change of the objects of the keys with several values, which are 20.7 % (`street`) and 37.9 % (`dirs`) of the keys:
a page of the grid takes 75.3 B against 67.8 B for the flat leaf's classes on `street`, 80.5 against 70.8 on `dirs`: +1.6 and +3.7 B/key, and `Set3` takes 0.6 to 0.9 B/key less than `vset`.)
Memory of `node-layout` (the baseline, no pages at all) for `uint64` values is not measured yet; the reference run for it is part of the work.

**Recommendation: keep the multi-key pages on for `T` of at most 8 bytes in 3.5** (the tree stays as it is; the single-key page replaces the flat leaf, which is its role under the multi-key
page too: an entry that gets a second value is promoted into it). The reason to switch them off was a clean comparison "page against leaf and nothing else"; the clean comparison is also
possible against the commit before 3.5 (`mkbaseline -ref`), and it does not give up a half of the memory of the case that matters most for `uint64 -> {*T}`-like data (one value per key).
`*T` has no multi-key page either way, and that case (one pointer per key: a page of 32 bytes per key, 67 B/key whole tree against 32 for a `uint64`) is where step 4 has most to gain: **the gap
between the two lines is the case for MKSV.**

## 7. What changes in the code

| where | change |
|---|---|
| `internal/skpage` | a second page flavor `Fixed` (same head, same classes, `BackLimit`, shrink rule): `New/Add/Remove/Has/Each/Len` generic in `T`; the allocation of the two kinds (section 3); 100 % coverage, a model test against `map[string]map[T]`, a fuzz test; remainder in nine bits as the string page |
| `internal/art` | `Map.flat`: 1 = fixed-size pointer-free, 2 = one word with a pointer, 3 = string, -1 = every key a value overflow; `leafWith` (the multi-key page's `mk`) makes a `Fixed` page; `fall back` and the promote path use the same functions; `flat.go`, `typed.go`, `leaf[T,K]`, `newSetLeaf`, `vals[T]` deleted; the value overflow generic in `T` (`valueOverflow[K,T]`) |
| `bench` | `objstat` knows the new objects; reference runs of the commit before 3.5 and of `node-layout` for `uint64` |

Order of work, each with its gate (race, 100 %, fuzz 60 s, lint): (0) **DONE 2026-10-04**: **cut the multi-key pages out of `internal/art`** (one commit of its own, tag `before-mk-pages-removal` is the way back): the page paths of insert, lookup, scan, delete and rebuild, `page.go`, the range nodes if only pages need them, `internal/vpage` and `internal/lpage` stay in the repository as standalone packages but nothing in the tree uses them; (1) `skpage.Fixed` standalone with microbenchmark against the flat and the typed leaf (`skbench -values words|pointers`);
(2) the value overflow generic in `T` for strings first (no change of behavior, tests as today); (3) the tree for `uint64`; (4) the tree for `*T`; (5) delete the old code; (6) measure
(`street`, `dirs`, `u64` keys with `*T` at 4K, 16K, 256K and a 1M spot check, PC; then M1).

## 8. Prediction and decisions

- **Memory** per key as in section 6 (pages on): `street` real, `uint64` 83 (+1 over today), `dirs` 101 (+3); one value per key unchanged; `*T` as `uint64`'s single-value lines plus the real mix, 90 / 106.
  A miss of more than 5 % stops the work.
- **Speed:** lookups, ranges, `churn`, `build` within the noise of the tree before 3.5 (the objects change their class grid and nothing about the access); the value overflow 5 to 15 % faster, not
  visible in these benchmarks (0.4 to 0.5 % of the keys).
- **Decision 1** (section 6): multi-key pages **on** (recommended) or off as the plan of 3.1 says.
- **Decision 2** (section 3): pointer pages only for one-word `T`; every other pointer-holding `T` becomes a value overflow per key (recommended: the three cases the user named are covered; the
  typed leaf's 64-byte `T` goes), or keep a typed page for `T` up to 16 bytes (more types).
- **Decision 3** (section 2): values aligned to the alignment of `T` and compared as `T` (recommended), or all values copied as bytes.
