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
