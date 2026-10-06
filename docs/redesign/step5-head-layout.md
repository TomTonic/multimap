# Step 5.5: the shared head of the pages, byte by byte (for review; nothing is built)

Words as in [GLOSSARY.md](GLOSSARY.md). The bytes of "today" are dumps of the real objects (`internal/mkpage`, `internal/skpage`, type bytes as in the tree: single-key pages `4 + 2·class`, multi-key pages `16 + 2·class`, value overflow `2`, nodes from `28`); the bytes of the proposals are constructed by hand from the same example and checked for size against the model. The example is the one of the design note: the keys **Bahnhof**, **Bahnhofsallee**, **Bahnhofstrasse** (three values), **Bahnhofweg**, i.e. six slots, common prefix `Bahnhof`, lengths `00 06 07 FF FF 03`.

## 1. Today: three heads

| | offset 0 | 1 | 2 | 3 | 4 | 5 | then |
|---|---|---|---|---|---|---|---|
| **multi-key page** (3 bytes) | type (bit 0 = bit 8 of cpl) | cpl, low 8 bits (length of the common prefix) | n (slots, 1 byte) | common prefix (starts here) | | | length list (n bytes), remainders, values |
| **single-key page** (6 bytes) | type (bit 0 = bit 8 of r) | r, low 8 bits (length of the remainder) | n, low byte | n, high byte (n is 16 bits) | kl, low byte | kl, high byte (length of the **whole key**) | remainder (r bytes), values |
| **value overflow** (6 bytes) | 2 (bit 0 = bit 8 of r) | r (511 = key held as a string) | n (not used) | | kl | | key area (18, 50, 114, 242, 370 or 498 bytes) or a string header, then the pointer to the value set |

What the first three bytes have in common today: type and the length (nine bits) are at the same place in all of them, and the low byte of n at offset 2. They differ in the **width of n** (1 byte in the multi-key page, 2 in the single-key page: a reader of 16 bits at offset 2 gets the first prefix byte as its high byte in a multi-key page) and in **`kl`**, which exists only because the single-key page keeps its key **from its base**: when a node is put above it, it stays as it is and `from(pathLen)` cuts the bytes the path already holds (`base = kl - r`). The multi-key page has no `kl`: it loses the bytes with `Skip` and regains them with `Prepend`, so its keys always begin at its path length.

### The example, today (real bytes)

**Multi-key page, strings** (class 64, 62 bytes used; type `12` = 16 + 2·1):
```
offset: 00 01 02 | 03............09 | 10 11 12 13 14 15 | 16 ............
        12 07 06 | 42 61 68 6E 68 6F 66 | 00 06 07 FF FF 03 | 05 4D 69 74 74 65 ...
        type cpl n | "Bahnhof"           | length list        | slot 0: remainder "" , value (05 "Mitte")
                                                              slot 1: remainder "sallee" (06) , 04 "Nord"
                                                              slot 2: "strasse" (07), 03 "Ost"
                                                              slot 3: (FF) 04 "Sued"   slot 4: (FF) 04 "West"
                                                              slot 5: "weg" (03), 04 "Ring"
   0: 12 07 06 42 61 68 6E 68 6F 66 00 06 07 FF FF 03
  16: 05 4D 69 74 74 65 73 61 6C 6C 65 65 04 4E 6F 72
  32: 64 73 74 72 61 73 73 65 03 4F 73 74 04 53 75 65
  48: 64 04 57 65 73 74 77 65 67 04 52 69 6E 67 00 00
```
**Multi-key page, `uint64` values** (class 128, 33 + 48 = 81 bytes needed; type `14`; the values are the last 48 bytes of the object, the 48 bytes before them are zero):
```
   0: 14 07 06 42 61 68 6E 68 6F 66 00 06 07 FF FF 03     head, "Bahnhof", length list
  16: 73 61 6C 6C 65 65 73 74 72 61 73 73 65 77 65 67     remainders "sallee" "strasse" "weg" (no remainder for the FF slots)
  32: 00 .. 00 (48 zero bytes, to offset 79)
  80: 07 00 00 00 00 00 00 00   01 00 .. 00     values 7, 1
  96: 02 00 .. 00   05 00 .. 00                            2, 5
 112: 09 00 .. 00   04 00 .. 00                            9, 4
```
**Single-key page, strings** for the key `Bahnhofstrasse` at path length 0 with the values Ost, Sued, West (class 64, 34 bytes used; type `06`):
```
   0: 06 0E 03 00 0E 00 42 61 68 6E 68 6F 66 73 74 72     type, r=14, n=3 (2 bytes), kl=14 (2 bytes), "Bahnhofstr"
  16: 61 73 73 65 03 4F 73 74 04 53 75 65 64 04 57 65     "asse", 03 "Ost", 04 "Sued", 04 "We"
  32: 73 74 00 ..                                          "st"
```
**Single-key page, `uint64` values**, same key, values 2, 5, 9 (class 64; the values start at the next multiple of 8 behind the remainder):
```
   0: 06 0E 03 00 0E 00 42 61 68 6E 68 6F 66 73 74 72     head (6 bytes), "Bahnhofstr"
  16: 61 73 73 65 00 00 00 00 02 00 00 00 00 00 00 00     "asse", 4 bytes of padding (20..23), value 2
  32: 05 00 .. 00  09 00 .. 00                              values 5, 9
```

## 2. What `kl` stores, and why one kind of page has it and the other does not

**The problem is the same for both kinds of page:** a page hangs below a byte node, and the **path** (the bytes the descent has matched) changes while the page does not: a byte node is put above it when a new key shares a part of its path, and a byte node above it goes away when the keys that made it are removed. When that happens, bytes that the page stores as **remainder** (or, in a multi-key page, as **common prefix**) are now also in the path, or bytes that were in the path now have to be stored by the page. The page must always know **where in the key its stored bytes begin**.

**Example.** The single-key page P holds the entry `Bahnhofstrasse`, made when the path length was 0. It stores the remainder `Bahnhofstrasse` (r = 14) and `kl` = 14, the length of the whole key. A new key `Bahnhofweg` arrives: a byte node with the common prefix `Bahnhof` and the branch bytes `s` / `w` is put above P. P's path is now `Bahnhofs`, path length 8. **P itself does not change**: it still stores all 14 bytes, the first 8 of which are now stored twice (in the node's common prefix and branch byte, and in P). On the next descent P must skip them. It knows how many: the path length (8, counted by the descent) minus the place in the key where its stored bytes begin, `kl − r` = 0 (this is what the code calls `base`; the glossary's word is *path length at the time the page was made*). `kl` is therefore **the position of the first stored byte, written as "length of the whole key minus length of the remainder"**.

**The two kinds, side by side:**

| | single-key page (today) | multi-key page (today) |
|---|---|---|
| where its stored bytes begin | **may begin before the path length**; the page says where (`kl − r`) | **exactly at the path length**, always |
| a byte node is put above it | nothing happens to the page | the page drops the first k bytes of its common prefix in place (`Skip`: a `memmove` of its content, up to 512 bytes) |
| a byte node above it goes away | page keeps its bytes if they reach back far enough, else a new page with the missing bytes (`rekey`, `Prepend`) | `Prepend`: the same, in place when the size class holds it |
| costs | 2 bytes a page (`kl`); bytes stored twice; every use of the key goes through `from(pathLen)`, `keyLen`, `base` (76 places) | the `memmove` at the event; no field, no stored twice |
| why it is so | historical: the leaf of step 1 held its key from where it was made, and the single-key page copied the head of the leaf so that the old code worked on it (design note step 3, "the header of 3 bytes is a later step") | new code of step 4, which started without the old |

So **it is not a property of one kind of page that needs `kl` and the other not**; two ways of one solution exist, and the tree uses both. The user's impression is right: it is needed everywhere (a field that says where the stored bytes begin; for the multi-key page that would be "how many bytes of the common prefix the path already holds", 9 bits) **or nowhere** (the page always begins at the path length and pays a `memmove` when the path changes).

**What the price of "nowhere" is: counted.** The events that change the path of a single-key page, in the benchmark's own streams (the probe, events per 1,000 operations; a cell that is not listed is 0; 4,096 and 65,536 keys):

| | node put above a single-key page (`splitLeaf`) | node above a single-key page goes away (`rekey`) | node put above a multi-key page |
|---|--:|--:|--:|
| street natural | 0.1, 0.2 | 0.0, 0.0 | 0.0, 0.0 |
| street single-value | 0 | 0.0 | 0.0 |
| dirs natural | 0.3, 0.4 | 0.5, 0.3 | 0.2, 0.2 |
| dirs single-value | 0 | 0.7, 0.9 | 0.5, 0.6 |
| `u64` keys natural | 0.1, 0.7 | 0.0 | 0 |
| `u64` keys single-value | 0 | 0 | 0 |

**At most 1 in 1,000 operations changes the path of a page.** The reason is that a new key that meets a single-key page almost never makes a node above it: it makes a **multi-key page** with it (`pair`: 6 to 184 in 1,000 operations). A `memmove` of 10 to 30 ns at those events costs **0.01 to 0.03 ns an operation on average**: nothing. (My first draft of this note said that a split of a single-key page is one of the common events of a new key; that was the situation before the multi-key pages and is wrong now.) What "everywhere" would cost instead: 2 bytes a page for the single-key pages **and** 9 bits for the multi-key page, the bytes stored twice, and the code of `base`/`from(pathLen)` in all pages.

**Result for the head:** the single-key page does not need `kl`; "nowhere" is the cheaper and simpler choice, and U3/U4 below have no `kl`.

## 3a. The two other questions of the head

1. **How wide is n?** One byte for everyone is enough: a multi-key page has at most 255 slots (a slot takes two bytes at least), a single-key page of strings at most 254 values, and a set of `uint8` values has at most 256 distinct values (the 256th would go to the value overflow). One byte lets every reader take `type | length | n` as the first three bytes of any page.
2. **Is there a spare byte?** The in-place pointer page (5.6) needs the number of untyped words in front, J (6 bits), stored; the multi-key page has no spare bit (nine bits of length, eight of n). A fourth byte `aux` is that place: J for the pages of pointers, 0 for all others, and room for what we do not know yet.

## 3. The two proposals side by side (same example)

| field | **today** multi-key | **today** single-key | **U3** (3 bytes, all kinds) | **U4** (4 bytes, all kinds) |
|---|---|---|---|---|
| 0 | type | type | type | type |
| 1 | cpl | r | len (cpl or r) | len (cpl or r) |
| 2 | n (1 byte) | n (low) | n (1 byte) | n (1 byte) |
| 3 | prefix starts | n (high) | prefix / remainder starts | **aux** (J for pointer pages, else 0) |
| 4 | | kl (low) | | prefix / remainder starts |
| 5 | | kl (high) | | |
| head size | 3 | 6 | 3 | 4 |
| single-key page keeps its key from | (n/a: Skip) | its base (`kl`) | the path length (`Skip`) | the path length (`Skip`) |

**Single-key page, strings, U4** (the same key and values; **32 bytes: it fits the class 32 where today's page takes the class 64**):
```
   0: 06 0E 03 00 42 61 68 6E 68 6F 66 73 74 72 61 73     type, r=14, n=3, aux=0, "Bahnhofstras"
  16: 73 65 03 4F 73 74 04 53 75 65 64 04 57 65 73 74     "se", 03 "Ost", 04 "Sued", 04 "West"
```
(U3: 31 bytes, the same class 32.)

**Single-key page, `uint64`, U4** (class 64 as today: 4 + 14 = 18, padding to 24, three values):
```
   0: 06 0E 03 00 42 61 68 6E 68 6F 66 73 74 72 61 73     head (4 bytes), "Bahnhofstras"
  16: 73 65 00 00 00 00 00 00 02 00 .. 00                  "se", padding 18..23, value 2
  32: 05 00 .. 00  09 00 .. 00                              values 5, 9
```
**Multi-key page, strings, U4** (63 bytes used, class 64 as today's 62; one byte more):
```
   0: 12 07 06 00 42 61 68 6E 68 6F 66 00 06 07 FF FF     type, cpl=7, n=6, aux=0, "Bahnhof", length list ...
  16: 03 05 4D 69 74 74 65 73 61 6C 6C 65 65 04 4E 6F     (... 03) 05 "Mitte", "sallee", 04 "No"
  32: 72 64 73 74 72 61 73 73 65 03 4F 73 74 04 53 75     "rd", "strasse", 03 "Ost", 04 "Su"
  48: 65 64 04 57 65 73 74 77 65 67 04 52 69 6E 67 00     "ed", 04 "West", "weg", 04 "Ring"
```
**Multi-key page, `*X` values, U4** (class 128; the typed object has J = 10 untyped words in front, the last 6 words are the values; `aux = 0A`; today J is derived from n, here it is stored and may stay when n changes):
```
   0: 14 07 06 0A 42 61 68 6E 68 6F 66 00 06 07 FF FF     type, cpl, n, aux = J = 10, ...
  16: 03 73 61 6C 6C 65 65 73 74 72 61 73 73 65 77 65     ...
  32: 67 00 ..                                              "g", zeros up to offset 79 (the nil gap inside the 10 untyped words: bytes only up to 79)
  80: 07 .. 01 .. 02 .. 05 .. 09 .. 04 ..                   the six typed words (values)
```
The multi-key page of strings and of `uint64` gets one zero byte (`aux`) at offset 3 and everything behind it moves by one; for `uint64` the object is still the class 128 (4 + 7 + 6 + 16 = 33, 33 + 48 = 81).

## 4. What each proposal costs and brings

**Memory** (`skmodel -multi -mkmv marker -header N`, bytes a key of the whole tree, model; the model gives the same header to both kinds, so it overstates the effect of U4 on the multi-key pages): the head size moves the real mix by 0.1 to 0.8 B a key from 3 to 6 bytes (street natural 56.4 / 56.5 / 56.7 / 57.2 B for 3 / 4 / 5 / 6; dirs natural 76.7 / 77.3 / 77.9 / 78.5; with several values in pages 37.9 / 37.9 / 37.9 / 38.0 and 57.3 / 57.4 / 57.6 / 57.6) and **nothing** for single-value entries (street 26.9, dirs 24.0 and 39.5 for every header size): the grid of 32 to 512 bytes is coarse, and 93 % of the single-value keys sit in multi-key pages. The size of the head is **not** the deciding question.

| | U3 | U4 |
|---|---|---|
| memory (model) | best by 0.1 to 0.2 B a key | +0.1 to 0.2 B a key for the multi-key pages; single-key pages 2 bytes smaller than today (the example's strings page falls from class 64 to class 32) |
| pointer pages (5.6) | J needs a place: variable head (3 or 4 bytes by kind) or a bit stolen from n | `aux` = J, one head for all |
| generic code: read `type | len | n` of any page | yes | yes (and `aux`) |
| single-key page without `kl` | yes | yes |
| risk | the 5.6 work needs a variable head | one byte more in the multi-key page |

**Time**: the multi-key page does not change (`Header` 3 → 4: every offset moves by one). The single-key page without `kl` gains the simpler key code (no `base`, no `from(pathLen)`: `rest` against `remainder`) and pays `Skip` in place at a `splitLeaf` and `Prepend` at a collapse (today `rekey`): at most 1 event in 1,000 operations (section 2). Prediction for the single-value cells: **±0 to +1 % in churn and build** (the gain: no `kl`, a few checks fewer on every use of a single-key page); a measurement on the PC and the M1 decides, after the model (5.5a: multi-key page to 4 bytes alone; 5.5b: single-key page to the shared head with `Skip`; each its own commit and measurement).

**Recommendation: U4**, in two steps (5.5a the multi-key page gets `aux`, 5.5b the single-key page and the value overflow take the head and `Skip`), because (1) the pointer page of 5.6 needs the byte and the user's priority (maps of pointers churn) makes that page important, (2) the head is then one head for every object that ends a descent, (3) the memory costs 0.1 to 0.2 B a key. What U3 gains over it (0.1 to 0.2 B a key) does not pay for a variable head.

Open for the review: n as one byte, the cap of 255 slots and values, `kl` gone (section 2), `aux`, and the value overflow, which would get the same head (`n` unused, `aux` 0) and a key area 2 bytes larger in the grid (20, 52, 116, 244, 372, 500).
