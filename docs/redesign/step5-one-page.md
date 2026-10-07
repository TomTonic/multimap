# One page for all: a layout in which the single-key page and the single-value page are the multi-key page with a few bytes left out (proposal, 2026-10-06; nothing built)

Words as in [GLOSSARY.md](GLOSSARY.md). Asked by the user after the head review ([step5-head-layout.md](step5-head-layout.md)): the four page types of today (single-key page for strings and for fixed-size values, multi-key page for strings and for fixed-size values) agree on the head but are four different structures behind it. SKMV and MKSV are special cases of MKMV; what they do not need should be **bytes left out, not a different structure**.

## 1. The one structure

A page is: the head, the key part, **two lists of lengths**, the remainders, free space, and the values, which end with the object.

```
Byte 0      type      size class (32 … 512) and which of the optional parts are there (below); bit 0 = bit 8 of len
Byte 1      len       length of the key part (9 bits)
Byte 2      n         number of slots = number of values (1 … 255); 0 for a value overflow
Byte 3      rawWords  pages whose values hold pointers: size of the byte area in 8-byte words; else 0
Byte 4…     key part  len bytes: the common prefix of all keys of the page
            key lengths    n bytes: the length of the remainder of slot i, or FF for a further value of the key before it
            value lengths  n bytes: the length of value i (0 … 254)
            remainders     the remainders of the keys, one behind the other
            free           zero
            values         ending with the last byte of the object: the values of the slots in order
                           (bytes of variable length one behind the other, or an array of T)
```

**What a page leaves out, and when** (the type byte says which parts are there):

| case | key lengths | value lengths | remainders | the key part is |
|---|---|---|---|---|
| many keys, values of variable length (MKMV, strings) | yes | yes | yes | the common prefix |
| many keys, values of one size (MKMV, `uint64`, `*X`) | yes | **left out** (every value is w bytes) | yes | the common prefix |
| one key, values of variable length (SKMV, strings) | **left out** (every slot is the one key) | yes | **none** | the whole remainder of the key |
| one key, values of one size (SKMV, `uint64`, `*X`) | **left out** | **left out** | **none** | the whole remainder of the key |
| one key, values outside (value overflow) | left out | left out | none | the whole remainder; the "values" are one pointer to the value set |
| one value per key (MKSV) | yes, never FF | as above | yes | the common prefix |

So a single-key page is a multi-key page whose key part is the whole key and whose key-length list is left out; a single-value page is a multi-key page without FF in its list; a value overflow is a single-key page whose values are a pointer. Where the values sit, how a slot is found and how a value is read is the same in all of them.

**Where the values sit:** at the end of the object, for every kind (today: only the multi-key page of fixed-size values). Keys grow from the front, values from the back, the free space is in between: an insert of a key moves no value, an insert of a value moves no key (the slotted page of databases). Fixed-size values are aligned without padding (the object is a multiple of 8 and of the size of T); value i of fixed size w is at `Size − (n − i)·w`. A value of variable length is at `Size − (sum of the value lengths from i on)`. The pointer area of a page of pointers is the end of the object, as the typed objects of Go need it, and `rawWords` is its boundary.

## 2. The examples, byte by byte

**One key, strings** (`Bahnhofstrasse`: `Ost`, `Sued`, `West`; 32 bytes, class 32, as U4; today 34 bytes, class 64):
```
Byte 0       type (one key, values of variable length, class 32)
Byte 1       len = 14          Byte 2  n = 3          Byte 3  rawWords = 0
Byte 4..17   key part "Bahnhofstrasse"
Byte 18..20  value lengths 03 04 04
Byte 21..31  values "Ost" "Sued" "West" (they end with the object; no free byte left)
```
**One key, `uint64`** (values 2, 5, 9; class 64 as today):
```
Byte 0       type (one key, values of one size, class 64)
Byte 1       len = 14          Byte 2  n = 3          Byte 3  rawWords = 0
Byte 4..17   key part "Bahnhofstrasse"
Byte 18..39  free (zero): a further value makes the array one slot longer at its front (it then begins at byte 32, the values move by one slot, as in the multi-key page today); a longer key part would grow into it
Byte 40..63  values 2, 5, 9
```
**Many keys, strings** (the six slots of the design note; 63 bytes, class 64, as today):
```
Byte 0       type (many keys, values of variable length, class 64)
Byte 1       len = cpl = 7     Byte 2  n = 6          Byte 3  rawWords = 0
Byte 4..10   key part "Bahnhof"
Byte 11..16  key lengths   00 06 07 FF FF 03
Byte 17..22  value lengths 05 04 03 04 04 04
Byte 23..38  remainders "sallee" "strasse" "weg"
Byte 39      free
Byte 40..63  values "Mitte" "Nord" "Ost" "Sued" "West" "Ring"
```
**Many keys, `uint64`** (the same slots; class 128, byte for byte the multi-key page of U4):
```
Byte 0       type (many keys, values of one size, class 128)
Byte 1       len = 7           Byte 2  n = 6          Byte 3  rawWords = 0 (with *X: 10)
Byte 4..10   key part "Bahnhof"
Byte 11..16  key lengths 00 06 07 FF FF 03
Byte 17..32  remainders "sallee" "strasse" "weg"
Byte 33..79  free
Byte 80..127 values 7, 1, 2, 5, 9, 4
```
**Value overflow** (`Bahnhofstrasse`, 32 bytes):
```
Byte 0       type (value overflow)    Byte 1  len = 14    Byte 2  n = 0    Byte 3  rawWords = 3
Byte 4..17   key part "Bahnhofstrasse"
Byte 18..23  free
Byte 24..31  the pointer to the value set (the value overflow's "values": one word, the pointer area)
```

## 3. What changes against today, and what it costs or brings

| | today | one page |
|---|---|---|
| structures behind the head | four (+ the value overflow) | one, with two optional lists |
| single-key page, strings | `vl v vl v …` behind the remainder | value lengths behind the key part, values at the end (the same bytes, in another order) |
| single-key page, fixed size | values right behind the remainder (padding), free slots at the end | values at the end, free space between |
| multi-key page, strings | value behind its remainder (`r vl v | vl v | r …`) | value lengths as a second list, remainders together, values at the end |
| multi-key page, fixed size | unchanged | unchanged |
| a single-key page meets a new key (`pair`) | a new multi-key page (both entries rebuilt) | the same as a multi-key page that gets a key outside its common prefix (`Widen`): one new object, as today |
| a multi-key page left with one key | a new single-key page (`onlyItem`, `leafFor`) | **in place**: the key-length list goes, the remainder joins the key part (a `memmove` of the page's bytes) |
| `kl`, `base`, `from(pathLen)` | in the single-key page and the value overflow | gone (section 2 of the head review) |
| code | `skpage` (2 flavors), `mkpage` (2 flavors), the tree's single-key and multi-key paths | one package with 2 flavors (variable and fixed size), one path in the tree |

**Memory** (estimate, to be confirmed with the model before anything is built): the multi-key pages are the same number of bytes (strings: the same bytes reordered; fixed size: unchanged). The single-key pages lose `kl` and their padding (−2 to −9 bytes a page) and gain nothing (their key-length list is left out): **−0.1 to −0.5 B a key** for the real mix (single-key pages hold 7 to 12 % of the keys), **±0** for one value a key. A few pages fall a class (the strings example: 64 → 32).

**Speed** (prediction, to be measured): the cases of one value per key are the multi-key pages of today, **±0** except **the multi-key page of strings**, whose value is no longer next to its remainder: a hit reads the value lengths (in the first line, next to the key lengths) and the value at the end of the object, in a page of 128 bytes or more possibly one cache line more than today: Get hit **0 to +5 %** on the PC (64-byte lines), ±0 on the M1 (128-byte lines). The real mix of random `uint64` keys (the case of 5.2 that lost 20 to 40 % in build) **gains** the conversions in place (439 a cycle of 69,746 operations) and one code path: **−3 to −10 %** there. The scan of a page (`Each`) reads the remainders and the values as two runs: ±0.

**What it does not change:** the head (U4), the grid of size classes, the byte nodes, the merge, the burst, the value set.

## 4. Questions for the user

1. **The value lengths as a list of their own** for strings (and the value away from its remainder) instead of `vl v` behind each remainder: one structure for every flavor, at the possible price of a cache line in a hit of a large page of strings.
2. **The key-length list left out for one key** (type byte: one key or many) or always there (a page with one key would have `00 FF FF …`, n bytes more, and no page type of its own: then even the type of the page does not change when the second key leaves).
3. **The values at the end** for every page (the two growing ends), against "behind the byte area" (your sketch of the pointer page).
4. **The order of the work:** this replaces 5.5 (head) and the core of 5.6 (the page of pointers with `rawWords` and free slots is this page with T = `*X`). Proposed: (a) the model (`skmodel`: one page with the parts left out) for the memory, with this prediction; (b) a new package with both flavors, its tests against a model as in 5.1, page-level benchmarks against `mkpage` and `skpage`; stop; (c) the tree on it, the probe; stop; (d) gate on the PC, then the M1 as decided.

## 5. Decided (user, 2026-10-06) and the plan

1. **The value lengths are a list of their own** for strings (values away from their remainders, at the end of the object). Yes.
2. **The key-length list is left out for one key**; the type byte says "one key" or "many keys" (two ranges as today: one-key pages `4 + 2·c`, many-key pages `16 + 2·c`, value overflow `2`, nodes from 28; `isPage`, `isSingleKey` keep their meaning). Yes.
3. **Where the values sit (end of the object or right behind the byte area) is found by a benchmark** (step 5.5b), not decided now.
4. **The order** below is accepted; it replaces the old 5.5 (head only) and the core of 5.6 (the page of pointers is this page with T = `*X`, `rawWords` stored, free slots).

### Fixed points of the layout for every step

- Head (4 bytes): `type | len | n | rawWords`; `len` nine bits (bit 0 of the type byte is bit 8); `n` slots 1 to 255 (0 in a value overflow); `rawWords` = size of the byte area in 8-byte words for pages whose values hold pointers (the boundary of the typed object, fixed for the life of the object), else 0.
- Byte 4: the key part (`len` bytes): the common prefix (many keys) or the whole remainder (one key, value overflow). Every page begins exactly at its path length (no `kl`, no `base`): when the path changes, `Skip` and `Prepend` in place (at most 1 in 1,000 operations, see step5-head-layout.md section 2).
- Then, as far as present: key lengths (n bytes, `FF` = further value of the key before; many keys only), value lengths (n bytes, 0 to 254; strings only), remainders (many keys only), free bytes (zero), values (strings: bytes in slot order; fixed size: an array of T, aligned without padding).
- A page with one key left is the one-key form (in place: the key-length list goes, the remainder joins the key part); a one-key page that gets a second key becomes the many-key form (in place when the key shares the whole key part and the page has room; else `Widen`, a new object).
- Limits as today: values up to 254 bytes, remainders up to 254 (many keys) and the key part up to 511, n up to 255; beyond: burst (many keys) or value overflow (one key; back into a page when the values fill at most half: `BackFits`).
- The value overflow: head, key part (key area of the grid 20, 52, 116, 244, 372, 500 bytes), the pointer to the value set in the last word (4 + area + 8 = 32 … 512); a remainder beyond 500 bytes: the key as a string at byte 8, `len` = 511.
- Values with a pointer: the object is a typed object (`ptrObject[T, [rawWords]uint64, [N]T]`, the generator moves to the new package); the map decides `ptr` once (`decide`) and passes it; values move as T (never as bytes).

### 5.5a The model (memory), no code of the library

Extend `bench/cmd/skmodel` with `-onepage`: the pages as in section 1 (one-key pages: 4 + r + value lengths (strings) + values; many-key pages: 4 + cpl + n + value lengths (strings) + remainders + values; value overflow: 4 + area + 8), against the model of today (`-multi -mkmv marker`, many-key head 3, single-key head 6 with `kl`). Cases: street, dirs, `u64` keys; natural and single-value; strings, words, pointers.
**Prediction:** real mix −0.1 to −0.5 B a key, single-value ±0.1 B a key, nothing worse than +0.3 B a key in any case. **Stop** with the table if a case is worse than +0.5 B a key or better than −1.0 (then the model or the layout is not understood).

### 5.5b The package `internal/page` with both flavors, alone (no tree)

- `page.Str` (values of variable length) and `page.Fixed` (values of one size T, pointer-free or one word with a pointer); one head type; `TypeOneKey`, `TypeManyKeys` set by the tree (as `TypeBase` today).
- API (both flavors, the names of `mkpage`): `BuildStrings` / `BuildFixedOf(…, ptr)` (one key if all rests are equal), `Match`, `CP`, `Get`, `EachValue`, `Add` (Added, AddedValue, Present, Full, Outside), `Widen`, `Remove` (Removed, Gone, Absent; converts to the one-key form in place when one key is left), `Each(fn(rem, val, first))`, `Keys`, `KeysUpTo`, `Len`, `Skip`, `Prepend`, `OneKey()`, `Size`, `Used`. Values of a key as a set (`Present` for a value that is there), in the order of arrival.
- Tests: the model test of 5.1 (`mvModel`, random adds, removes, widenings, skips, prepends; flavors strings, `uint64`, `*uint64`), the transitions one key ↔ many keys in place, the layout pinned byte by byte for the examples of section 2, limits, the collector test for pointers (finalizers, a collection every few steps), `FuzzOnePage`; 100 %, race, lint 0.
- **The placement benchmark (decision 3):** the value offset is one function (`valueAt`, `valuesStart`) with two implementations behind a build tag (`valfront`: values right behind the byte area, free space behind them; default: values at the end, free space in between). `BenchmarkOnePage`: Get hit and miss, add a key, add a value to a key, remove, `Each`, for n = 3, 7, 20, strings, `uint64`, `*X`; **hot** (4,096 pages) and **cold** (pages spread over 64 MB, random order). Run on the PC (WSL, medians of 6) and, as a short queue job, on the M1. **Rule:** values at the front only if the cold Get hit and `Each` are at least 3 % faster and no change is more than 5 % slower; else at the end. The losing variant and the tag are deleted. **Prediction:** at the end is no slower in the changes (keys and values never move each other) and equal in the hot Get; at the front is 0 to 5 % faster in the cold Get of fixed values when keys and values of a page share a cache line (64-byte lines), ±0 on the M1.
- Page-level comparison against `mkpage` and `skpage` of today (`BenchmarkFixed`, `BenchmarkPage`, `BenchmarkFixedPointers`, `BenchmarkFixedScan`, the single-key page benchmarks of `skpage`): **prediction** many keys `uint64` ±3 %; many keys strings Get 0 to +5 % (the value is away from its remainder), `Each` ±5 %; one key: Get ±3 %, add a value: +0 to +10 % for strings (a length byte inserted into the list), ±3 % for fixed size; pointers: add and remove **in place** (no new object while the byte area and the free slots hold): −50 to −70 % against the page of 5.3.
**Stop** with the results.

### 5.5c The tree on the package (`internal/art`), in steps that each pass the gates

1. The multi-key pages from `mkpage` to `page` (many-key form only; the single-key pages still `skpage`); `pair`/`onlyItem` unchanged in function. Probe and `BenchmarkMapAddPresent`: ±3 % against `2d5d08b`.
2. The single-key pages to the one-key form of `page`: `kl`, `base`, `from(pathLen)` go (the 76 places); `splitLeaf` and `collapse` call `Skip`/`Prepend` in place; `pair` becomes the many-key form of the same page (in place or `Widen`); the conversion of a page with one key left disappears from the tree (the page does it in `Remove`). New events: one-key ↔ many-key in place, `Skip`/`Prepend` of a one-key page.
3. The value overflow to the new head (key area 20 … 500, the pointer in the last word).
4. `skpage` and `mkpage` deleted; the glossary: page (one structure, one-key and many-key form), key lengths, value lengths, byte area, pointer area, `rawWords`.
**Prediction** (probe, against `2d5d08b`): single-value cells ±3 %; real mix of random `uint64` keys −3 to −10 % (conversions in place); real mix street/dirs ±3 %; maps of pointers: churn and build −30 to −50 % against the page of 5.3 (changes in place), memory as `uint64`. **Stop** after 2 with the probe, and after 4.

### 5.5d Gate on the PC (run script as `run-m44.cmd`, three builds: `uint64`, strings, pointers), then the M1 (the jobs deferred in `bench/remote/queue.txt`, written again with that commit); the report against all predictions of step 5; then the decision about step 6 (routing).

### Working rules for this plan (unchanged)

A design note with a prediction before code; stop at a surprise and ask; the vocabulary of the glossary; one measurement at a time (duration and end time said before); the M1 only through the queue; every step its own commit with tests (100 %, race, lint 0) and its own measurement; push with `git -c credential.helper= -c "credential.helper=!gh auth git-credential" push`.

## 6. Results

### 5.5a The model (memory), 2026-10-06

`skmodel -multi -onepage -mkmv marker -mkgrid 32,64,128,256,384,512 -values string|words|pointers -ovbytes 16.3`: the tree of byte nodes with the pages of today (single-key head 6 with the padding of the remainder to a word before fixed-size values, multi-key head 3, several values in the many-key pages: marker) against the one page (head 4 for both, one-key pages without key lengths and without padding). Bytes a key, whole tree, with the value overflows' value sets:

| values | data | today | one page | difference |
|---|---|--:|--:|--:|
| strings | street natural | 44.5 | 44.6 | +0.1 |
| strings | street single-value | 31.0 | 31.1 | +0.1 |
| strings | dirs natural | 80.3 | 80.0 | **-0.3** |
| strings | dirs single-value | 50.7 | 50.7 | 0.0 |
| `uint64` and `*X` | street natural | 38.0 | 38.1 | +0.1 |
| `uint64` and `*X` | street single-value | 26.9 | 27.1 | +0.2 |
| `uint64` and `*X` | dirs natural | 57.6 | 57.7 | +0.1 |
| `uint64` and `*X` | dirs single-value | 39.7 | 39.7 | 0.0 |
| all | random `u64` keys, one value | 36.5 (strings), 24.0 (words, pointers) | the same | 0.0 |

**Prediction** (real mix -0.1 to -0.5 B a key, single-value ±0.1, nothing worse than +0.3): **partly missed in the sign**: the real mix is -0.3 to +0.1 and street single-value `uint64` +0.2. The single-key pages lose `kl` (two bytes, 7 to 14 % of the keys) but the multi-key pages (86 to 93 % of the keys in the single-value cases) gain a head byte (`aux`/`rawWords`), a page more or less crosses a class border (the single-key pages of the grid are coarse), and a many-key page one byte fuller bursts a little earlier (pages +59 in street single-value). **Memory is neutral: -0.3 to +0.2 B a key**; no stop condition of the plan (worse than +0.5 or better than -1.0) is met. The memory is therefore no argument for or against the one page; the arguments are the code (one structure, no `kl`/`base`, conversions in place) and the speed of the pages with a pointer (5.6), which the next steps measure.

### 5.5b The package `internal/page` alone (2026-10-06, commit `0c4eb26`; benchmarks WSL, hot, best of 4)

**Built:** `internal/page` with `Str` (values of variable length) and `Fixed[T]` (one size, pointer-free or one word with a pointer), one head of four bytes, the one-key and the many-key form in one structure (the layout of section 1 and 2, byte for byte), the API of `mkpage` (`BuildStrings`, `BuildFixedOf`, `Match`, `CP`, `Get`, `EachValue`, `Add`, `Widen`, `Remove`, `Each`, `Keys`, `KeysUpTo`, `Len`, `Skip`, `Prepend`, `Used`) plus `OneKey()` and `RawWords()`. Behaviour as designed: a many-key page left with one key becomes the one-key form **in place** (the key-length list goes, the remainder joins the key part; this always fits: a page is at most 512 bytes); a one-key page that meets another key becomes a new many-key page (`Add` makes it, as `pair` does); a page of pointers keeps its byte area (`rawWords`) for its life and changes in place while the byte area and the free slots hold. Tests: the model test (60 seeds × 800 random adds, removes, widenings, skips and prepends, flavors strings, `uint64`, `*uint64`, the form checked after every step), the layout pinned byte for byte for the examples of section 2, limits, the transitions between the forms, Skip/Prepend, `Widen` refusals, lookups and removals of what is not there, the collector test for pointers (finalizers), `TestPointerPageChangesInPlace`, `FuzzPage` (30 s, 2·10^5 runs, nothing found). **100 %, race, lint 0.** Three simplifications found on the way: `MaxKeyPart` can never be reached (a page is at most 512 bytes), so neither the conversion to the one-key form nor `Prepend` can fail for the length of the key part; a many-key page has two keys at least; the type bytes of the two forms are constants (`TypeOneKey` 4, `TypeManyKeys` 16: the tree's ranges, a test of the tree will pin them).

**Placement of the values (decision 3): at the end of the object stays.** `BenchmarkPlacement` (a Get hit in a page with the values at the end against a copy with the values right behind the byte area; pages hot (4,096) and cold (512 Ki pages, 64 MiB and more, random order); n = 3 / 7 / 20, `uint64`): hot end 30.8 / 37.8 / 78.5 ns against front 31.0 / 36.7 / 77.2; **cold end 155 / 170 / 232 against front 160 / 179 / 235 ns** (the front layout is 2 to 5 % **slower** when cold, ±2 % hot). The rule of the plan (front only if the cold Get hit is at least 3 % faster) is not met, so the values stay at the end, the two-ended page where an insert of a key moves no value and an insert of a value moves no key. (The reason is in the grid: a page of 128 bytes or more holds more than 64 bytes of content, so its values are in the second 64-byte line behind the keys in either layout, except the first one or two of a small key area.) The M1 runs the same benchmark as the queue job `pg1` (128-byte lines); nothing is built for the front layout.

**Page level against `mkpage` and `skpage` of today** (ns, hot; below 0 % is faster than today):

| | n = 3 | n = 7 | n = 20 |
|---|--:|--:|--:|
| many keys, `uint64`: Get hit | 17.1 → 16.3 (-4 %) | 23.1 → 22.7 (-1 %) | 42.8 → 41.3 (-3 %) |
| many keys, `uint64`: Get miss | 9.7 → 9.2 (-5 %) | 15.8 → 14.0 (-11 %) | 35.1 → 29.5 (-16 %) |
| many keys, `uint64`: add and remove a key | 52.7 → 61.1 (**+16 %**) | 63.1 → 70.0 (**+11 %**) | 86.9 → 92.9 (+7 %) |
| many keys, `uint64`: `Each` (a value) | | 18.3 → 19.2 (+5 %) | 46.6 → 44.9 (-4 %) |
| many keys, strings: Get hit | 17.5 → 17.0 (-3 %) | 24.6 → 22.6 (-8 %) | 45.9 → 41.0 (-11 %) |
| many keys, strings: Get miss | 10.4 → 9.2 (-11 %) | 17.6 → 14.2 (-19 %) | 41.5 → 30.1 (-27 %) |
| many keys, strings: add and remove a key | 51.4 → 76.3 (**+48 %**) | 66.3 → 87.4 (**+32 %**) | 116.5 → 123.4 (+6 %) |
| many keys, strings: build | 75.8 → 83.2 (+10 %) | 140 → 154 (+10 %) | 303 → 326 (+8 %) |
| pointers: Get hit | 17.1 → 16.4 (-4 %) | 23.0 → 22.5 (-2 %) | 42.4 → 41.5 (-2 %) |
| **pointers: add and remove a key** | 158.6 → 64.1 (**-60 %**) | 266 → 72.7 (**-73 %**) | 417 → 96.2 (**-77 %**) |
| one key, strings (3 values): Get hit | 7.0 (match and has) → 5.6 | | |
| one key, strings: add and remove a value | 28.5 → **68.1** | | |
| one key, `uint64`: Get hit | 3.7 (match and has) → 4.8 | | |
| one key, `uint64`: add and remove a value | 8.6 → **30.9** | | |

**Against the prediction** (many keys `uint64` ±3 %; many keys strings Get 0 to +5 %, `Each` ±5 %; one key Get ±3 %, add a value strings +0 to +10 %, fixed ±3 %; pointers -50 to -70 %):
- **Met:** Get of both flavors and of pointers (-1 to -7 %, a miss -5 to -27 %), `Each` (-4 to +5 %), pointers' changes (**-57 to -76 %: the page of pointers changes in place, which is what 5.6 wanted**).
- **Missed:** the changes of the many-key page: `uint64` +13 to +25 %, strings +11 to +55 % (small pages the most: the fixed cost of the general code is a larger part of a short operation: about +7 ns for `uint64`, +14 to +28 ns for strings, which write three lists and the values where the page of today wrote one entry), and the one-key page's add and remove of a value: +12 ns (`uint64`) and +40 ns (strings) over the single-key page, which appended to an unsorted array in place (the one-key form keeps the order of arrival, like the many-key form).
- Cost, as far as measured: `p.lay` and the `lay` struct 13 % of an `Add`, three memmoves for the key area (one more than the entry insert of `mkpage` for strings), `KeysUpTo(2)` in `Remove` of a key; after the changes already made (the constants for the type bytes, the scalar `locate`, one block move of the lists) the many-key changes cost 7 ns more than `mkpage` in `uint64`. **At the level of the tree the changes are about 10 to 20 % of an operation and the page change in it one half: expected +1 to +4 % of an operation for the many-key page of strings, ±1 % for `uint64`, which 5.5c measures.** Tuning candidates, not done: scalar locals in `Add`/`Remove` instead of the `lay` struct, a fast path for the in-place `Add` and `Remove` of a value of a one-key page, the value lengths of strings next to the values.

**Stop** for the user: the page is built, measured and tested alone; the tree is not touched.

### 5.5b tuned before the tree (user: "first tune, so that the comparison in the tree is cleaner")

What was done (commit after `0c4eb26`): the type bytes as constants, `locate` with scalars and one `compare`, the lists and remainders moved as one block (three copies), `oneKeyLeft` (a scan that stops at the second key) instead of counting keys on a removal, `shrinkClass` with a table (one comparison when the page stays), the value comparison of strings after the length. The profile of what is left (`Add`/`Remove` of a many-key page of strings, n = 3): 37 % of `Add` and 31 % of `Remove` are the three `memmove` calls of the key area (the lists and the remainders are four regions where `mkpage` had one entry to move), 8 % each the first touch of the page (a cache miss, as in `mkpage`), `locate` 20 %, `hasValue` 20 % of the one-key `Add` (the compare of the values, as the single-key page did). **Tuned page against today's, hot:** Get -1 to -11 % (a miss -11 to -27 %), `Each` -4 to +5 %, pointers' changes **-60 to -77 %**, many keys `uint64` changes **+7 to +16 %** (were +13 to +25), strings **+6 to +48 %** (were +11 to +55: n = 3 +48 %, n = 7 +32 %, n = 20 +6 %), the one-key add and remove of a value +40 ns (strings) and +22 ns (`uint64`). What is left is structural, the price of the layout chosen with the user (value lengths as a list of their own, values at the end): more regions to shift on a change of a short page; the tree-level cost is measured in 5.5c (at most a few percent of an operation). 100 %, race, lint 0.

### 5.5c part 1: the multi-key pages of the tree on `internal/page` (2026-10-07, probe, WSL, medians of 11, against `2d5d08b`)

The tree holds many-key pages of the one page; single-key pages are still `skpage`; a many-key page left with one key is converted at once by the tree (`onlyItem`, `leafFor`: transitional work that step 2 removes, since the page does it in place). Tests adapted to the new sizes (two tests of full pages: 237-byte values), a test pins the type bytes (`TestPageTypesAreThoseOfTheTree`), 100 % everywhere, race, lint 0.

**Prediction** (single-value cells ±3 %, real mix ±3 %, strings at most +4 %): **missed.** Build / replay ns an operation, before → after the tuning of the page (inlined key-area moves, no `lay` wrapper for `locate`):

| case | before | now |
|---|--|--|
| `uint64` single-value 4,096 | 79 / 70 | 87 / 79 (+10 % / +13 %) |
| street single-value 4,096 | 205 / 219 | 217 / 229 (+6 % / +5 %) |
| dirs single-value 4,096 | 266 / 264 | 280 / 275 (+5 % / +4 %) |
| `uint64` natural 4,096 | 100 / 82 | 103 / 88 (+3 % / +7 %) |
| street natural 4,096 / 65,536 | 191 / 175, 261 / 298 | 200 / 183, 275 / 290 (+5 % / +5 %, +5 % / -3 %) |
| dirs natural 4,096 | 238 / 221 | 245 / 228 (+3 % / +3 %) |
| `uint64` natural 65,536 | 176 / 189 | 194 / 196 (+10 % / +4 %) |

**Cause** (profile of `uint64` single-value, 4,096 keys, where nearly every operation is an operation in a page): the page's `Add` and `Remove` cost 1.21 and 1.26 s where `mkpage` took 1.08 and 1.06 s of 3.1 s: **+10 to +20 % in the tree**, the page level's +0 to +6 % (after inlining the key-area moves of the page of fixed-size values, which took the changes from +7 to +16 % to +0 to +6 %, and of strings from +12 to +48 % to +1 to +39 %). In the tree the pages are colder than in the benchmark and every instruction counts as much as a cache miss; what is left is the structure's cost (the `lay` set-up with the first touch of the page, the form check, one more branch in every step). Part of the loss is **transitional** (the conversion of a page with one key left, done twice: in the page and in the tree) and not measurable here (6 conversions in 1,000 operations).

### 5.5c parts 2 and 3: the single-key pages and the value overflow on the one page (plan, 2026-10-07; user: "A", go on)

**Parts 2 and 3 are done together** (a deviation from the plan, for a reason): the tree treats a single-key page and a value overflow as one thing (`singleKeyHead`: `base`, `keyLen`, `from`, `matches`, `rekey`: 76 places); a state in which the pages have no `kl` and the overflow still has it would need both in every one of them.

**What changes.** (1) The head of every leaf is the head of the page (`type | len | n | rawWords`, 4 bytes); the key part (the remainder) begins at byte 4 and **begins exactly at the path length**: `kl`, `base`, `keyLen`, `from(pathLen)` and `maxKeyLen` are gone; `matches(rest)` compares `key[pathLen:]` with the key part. (2) A single-key page is the one-key form of `page.Str`/`page.Fixed`; the value overflow keeps its type (2), its key area is 20, 52, 116, 244, 372 or 500 bytes (4 + area + 8 = 32 … 512), the key as a string above 500. (3) A new leaf is made **with its first value** (`page.NewStr`, `page.NewFixed`; no page without a value). (4) `upsert` is the add: it reaches a leaf and calls `Add` of the page (a different key: the page makes the many-key page itself, `pair` is gone; `Full` for the same key: value overflow; `Full` for another key: a node above, `splitLeaf`) or `overflowAdd`; the leaf that a node is put above loses the bytes of the path with `Skip` (page and overflow, in place), and the leaf that a node above it leaves regains them with `Prepend` (a page that no longer fits becomes a value overflow, an overflow that has no room gets a larger one). (5) A many-key page left with one key is converted by the page itself: the tree's conversion (`onlyItem`, `leafFor`) is gone. (6) The back-conversion of an overflow into a page: `page.BackFits`. (7) `skpage` is not used by the tree any more (deleted in part 4).

**Prediction** (probe, against the state after part 1, `1ea4ddd`, and against `2d5d08b` = the state before 5.5c): single-value cells -3 to +1 % against part 1 (no double conversion, no `kl` compare in a lookup, `from(pathLen)` gone; the one-key page's add and remove of a value is slower, 10 % of the operations), i.e. **+0 to +10 % against `2d5d08b`**; real mix of random `uint64` keys -3 to -10 % against part 1 (no conversions: pair and the page left with one key); maps of pointers: churn and build **-30 to -50 %** against `2d5d08b` (pages change in place; against m43's typed leaf the credo question is open); memory **±0.3 B a key** against part 1 (the one-key pages lose 2 bytes, the overflow grid moves by 2). Stop after the part with the probe.

### 5.5c parts 2 and 3: results (2026-10-06, probe, WSL, medians of 11, three builds one after the other; commit `bbbabd7` and the fix `Test the size before allocating in backToPage`)

Done as planned (the leaves are the one-key form of the page, the value overflow has the head of the page). Two things the plan did not say: (1) **the value overflow keeps the offset of its value set in `rawWords`** (the head's last byte, in words): its key area is chosen when the object is made and the key part shrinks with `Skip`, so the offset cannot follow the key part's length (the first run crashed on this: the set was read at the offset of the shorter key); the size of the object, for the statistic, is the same number (`valueOverflowSize`). (2) `Fixed.Add` refused no 256th slot (`n` is one byte): a key of 256 values of one byte overflowed the head; found by `TestValueTypes`, now `TestPageSlotLimit`. 100 % in `art`, `page`, root; race, lint 0, vet with `mkstats`, `ptrvals`, `baseline`.

**First probe run: a surprise, understood and fixed.** The cells with several values a key were +100 % (`uint64` natural 4,096 replay 78 → 168 ns, street 185 → 333). The events were the same as in part 1, so the cause was not the structure: the profile showed `backToPage` (14 %, with `makeslice` 7 %): a value overflow tries to become a page **at every removal**, and the slice of keys was allocated before the size test. A bug of the implementation (the size test now comes first and allocates nothing), not a question of the design.

**Probe, build / replay in ns an operation: before 5.5c (`2d5d08b`) → part 1 (`1ea4ddd`) → now:**

| case | before | part 1 | now | now against part 1 | now against before |
|---|--|--|--|--|--|
| `uint64` single-value 4,096 | 77 / 70 | 89 / 83 | 87 / 78 | -2 % / -6 % | +13 % / +11 % |
| street single-value 4,096 | 201 / 222 | 222 / 231 | 209 / 223 | -6 % / -3 % | +4 % / +0 % |
| dirs single-value 4,096 | 266 / 263 | 283 / 273 | 264 / 262 | -7 % / -4 % | -1 % / -0 % |
| `uint64` natural 4,096 | 99 / 84 | 102 / 77 | 105 / 86 | +3 % / +12 % | +6 % / +2 % |
| street natural 4,096 | 190 / 173 | 222 / 181 | 191 / 177 | -14 % / -2 % | +1 % / +2 % |
| dirs natural 4,096 | 237 / 223 | 247 / 234 | 238 / 227 | -4 % / -3 % | +0 % / +2 % |
| `uint64` natural 65,536 | 171 / 180 | 184 / 191 | 181 / 184 | -2 % / -4 % | +6 % / +2 % |
| street natural 65,536 | 259 / 276 | 271 / 290 | 262 / 282 | -3 % / -3 % | +1 % / +2 % |

**Against the prediction:** single-value against part 1 -3 to +1 %: **better** (-2 to -7 %); against `2d5d08b` +0 to +10 %: met except `uint64` single-value (+13 / +11 %: the one-key page's add and remove of a value, 10 % of its operations, is the tuning candidate "fast path for the in-place `Add` and `Remove` of a value of a one-key page"). Real mix of random `uint64` keys -3 to -10 % against part 1: **missed** for `uint64` (+3 / +12 % at 4,096, -2 / -4 % at 65,536; part 1's replay of 77 ns is lower than `2d5d08b`'s 84, i.e. the cell varies by a few percent between builds); street and dirs -2 to -14 %, as predicted. Memory (block bytes a key, as built): **±0.3 B against part 1 and `2d5d08b`** (`uint64` single-value 21.2 / 21.2 / 21.2, street natural 40.5 / 40.8 / 40.8, dirs natural 72.3 / 72.6 / 72.1): as predicted.

**Maps of pointers** (probe with `ptrvals`, medians of 7), `2d5d08b` → part 1 → now: `uint64` single-value 4,096 build 321 → 236 → 222 (-31 %), replay 374 → 84 → 85 (**-77 %**); natural 4,096 169 → 115 → 122 / 95 → 80 → 84 (-28 % / -12 %); street natural 408 → 237 → 236 / 457 → 194 → 193 (-42 % / -58 %); `uint64` natural 65,536 276 → 211 → 206 / 284 → 213 → 210 (-25 % / -26 %). Prediction -30 to -50 % for churn and build: met for the single-value and street cells, **below it** (-12 to -28 %) for the natural cells of `uint64`, where most keys stay in pages with several values and the old typed leaf churned less. The gain came with part 1 (the many-key pages in place); parts 2 and 3 add ±3 %.

**Next (the plan says: stop here).** Part 4: delete `skpage` and `mkpage` (the `AllocPtr` generator moves into `page`), glossary; then 5.5d, the gate on the PC and the jobs for the M1. Open for the user: the one tuning candidate that is left, the fast path of the one-key page (+13 % in `uint64` single-value against before 5.5c), before or after the gate.

### 5.5c part 4: `skpage` and `mkpage` are gone (2026-10-07)

Deleted: `internal/skpage`, `internal/mkpage`, `bench/cmd/skbench` (the diagnosis of the old single-key page) and `internal/art/fixed_bench_test.go` (a standalone benchmark of `skpage.Fixed`; the page-level benchmarks of `internal/page` replace it). `internal/page` now holds what the tree still took from them: `kind.go` (`Supported`, `HoldsPointers`, which `T` takes which page: reflection once per map) and the generator of the 166 typed objects of the pages of pointers (`gen/`, `fixed_ptr_gen.go`; `go generate ./internal/page` gives the same file). `skmodel` lost its `realPage` check against `mkpage` (the model of the one page is pinned by `TestPageLayout` and the memory numbers of the probe). New tests `TestSupportedTypes` and `TestAllocPtrEveryShape` (the old ones lived in `skpage`). Glossary: the section "The one page" (one-key form, many-key form, head, key part, key lengths, value lengths, byte area, pointer area, value overflow). 100 % in the root, `art` and `page`, race, lint 0, vet with `mkstats`, `ptrvals`, `baseline`.

**Next: 5.5d**, the gate on the PC (three builds as `run-m44.cmd`), the M1 jobs written with that commit (`m44`, `m44s`, `m44p`; `pg1` is queued), the report against the predictions of step 5, then the decision on step 6. The one-key fast path (+13 % `uint64` single-value against before 5.5c) comes after the gate, as decided.

### 5.5e the reads in the tree (2026-10-07; after the gate m46: reads 5 to 12 % slower)

**Measurement** (new: `bench/cmd/bench/readprobe_test.go`, `MKREAD=1`; a tree built from the corpus, `Each` of every key in random order and a scan of all values, WSL, medians of 5 interleaved runs; builds at `7929687` = m44, `2d5d08b` = before 5.5c, `1ea4ddd` = part 1, now). `Each` ns a key / scan ns a value:

| case | m44 | before 5.5c | part 1 | now |
|---|--|--|--|--|
| `uint64` single-value 4,096 | 36.0 / 6.61 | 35.9 / 6.60 | **47.4** / 6.06 | 47.7 / 6.03 |
| `uint64` natural 4,096 | 45.8 / 3.38 | 44.6 / 3.25 | 47.4 / 3.23 | 52.2 / **4.04** |
| street natural 4,096 | 72.8 / 4.98 | 73.2 / 4.95 | 81.9 / 5.18 | 82.8 / 5.56 |
| street single-value 4,096 | 63.4 / 7.32 | 63.6 / 7.28 | 73.4 / 7.88 | 73.2 / 8.22 |
| `uint64` natural 65,536 | 86.0 / 6.29 | 85.7 / 6.28 | 93.1 / 6.31 | 98.1 / 6.88 |

**The regression came with part 1** (the many-key pages on the one page): a lookup of a key costs +10 ns (+30 % for `uint64` single-value, where the tree's descent is short) and part 2+3 adds to the scan of the single-key pages in the natural mix (+10 to +25 %).

**Cause** (CPU profile of `Each`, `uint64` single-value 4,096, 3 s each): `Fixed.EachValue` of the one page builds the `lay` struct (`head.lay`, 13.8 % of the time; it did not exist in `mkpage`), passes it by value to `slotOf` (6.4 % of its own, plus the `Match` call that walks the common prefix again) and `EachValue`'s own time is 15 % where `mkpage`'s was 11 %; `locate` and `compare` cost the same as before. The page's `Get` was already written with scalars in 5.5b (the tuning of 5.5b skipped `EachValue` and `Each`, which the tree's `Each` and scans use).

**Prediction** (the same fix as `Get`: `EachValue` and `Each` of both flavors with scalars, no `lay`, no `slotOf`): `Each` of a key `uint64` single-value 47.7 → **37 to 40 ns** (m44: 36.0), street 73 → 64 to 67 (m44: 63); the scan of the natural mix back to m44 ±3 % (3.4 / 5.0 ns a value). Writes unchanged (±2 %).

**Result** (commit after `69c65ea`: `EachValue` and `Each` of `Fixed` and `Str` without `lay`/`slotOf`, `slotOf` deleted; new `EachSingle` for the scan of the single-key pages, which saves the closure around the closure; 100 %, race, lint 0). `Each` ns a key / scan ns a value, medians of 5 interleaved runs, m44 → before the fix → now:

| case | m44 | before the fix | now |
|---|--|--|--|
| `uint64` single-value 4,096 | 36.3 / 6.64 | 47.6 / 6.03 | **33.2** / 5.84 |
| `uint64` natural 4,096 | 45.6 / 3.37 | 52.2 / 4.04 | **42.3** / 3.04 |
| street natural 4,096 | 73.1 / 4.96 | 81.6 / 5.51 | **68.4** / 4.91 |
| street single-value 4,096 | 63.5 / 7.35 | 73.4 / 8.20 | **60.6** / 7.40 |
| `uint64` natural 65,536 | 85.0 / 6.26 | 95.5 / 6.79 | **81.2** / 6.13 |

**Prediction met and beaten** (`Each` of a key `uint64` single-value 37 to 40 ns: 33.2; street 64 to 67: 60.6 / 68.4; scan of the natural mix within ±3 % of m44: -10 % to +0 %): the reads are **-4 to -9 % against m44** for lookups (single-value `uint64` -9 %) and -9 % to +1 % for scans, i.e. the layout of the one page reads faster than `mkpage`/`skpage` once the page's own code has no per-call set-up. Writes unchanged (probe, medians of 7: `uint64` single-value 4,096 build 86 → 87 / replay 81 → 78, street natural 192 → 193 / 177 → 175, `uint64` natural 65,536 182 → 182 / 194 → 184). The lesson for the method: **a probe that measures writes only does not pin the reads**; `readprobe_test.go` (`MKREAD=1`) is now part of the probes, and the gate has to be run again for the numbers (the PC: `m47`).
