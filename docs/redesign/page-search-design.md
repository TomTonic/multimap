# The search in a page: word compare or fingerprint (design note, 2026-10-07)

Words as in [GLOSSARY.md](GLOSSARY.md). Origin: analysis (b) of the review ([review-2026-10.md](review-2026-10.md)): a lookup of a
small `u64` map takes 33 ns against 17 in `main`, and 64 % of it is the search in the page (`locate` and `compare`); street single-value
58.6 against 47.1. The measurement of the first bytes on the real corpora (same file, last section) rules out a list of first bytes: of the
5 to 7 keys of a page, 3 to 5 share the first byte.

## Where the time goes today

`locate(m, kl, n, rem, r)` walks the key lengths; for every first slot it compares the remainder with the rest of the key byte by byte
(`compare`: a loop up to the first difference) and adds its length to the offset of the next remainder; it stops at the first remainder
that is not smaller. A hit compares about half the keys of the page (3 to 8), each up to its first differing byte; a miss as many. The
compare loop and the sum are serial: each step waits for the one before.

`main` reaches the key with a few word-wide matches in its byte nodes and one compare of the leaf's key: no linear search.

## Option 1: compare a word at a time (no change of the layout)

`compare` of two remainders loads 8 bytes of each at once (big-endian, so that the order of the words is the order of the bytes), compares
them as integers and goes to the next word only if they are equal; the last, shorter word is read as a whole word from the page and masked
to its length. A page is at least 32 bytes and the values follow the remainders, so an 8-byte load from the start of a remainder never
leaves the object; only at the very end of the object, where the values of a `Fixed` page are, would it need a guard, and remainders never
lie there. The rest of the key, the caller's slice, is loaded the same way with a guard at its end. The search stays linear and in order,
so it serves lookups, the place of a new key and the bounds of a range alike.

**Prediction:** a compare of remainders of 7 to 25 bytes from 3 to 10 ns to 1 to 3 ns; lookups `u64` single-value 4,096 33 → **26 to 29 ns**,
street single-value 58.6 → **52 to 56**, dirs (remainders of 15 to 25 bytes) a little more; writes 3 to 8 % faster (the add and remove of a
key call `locate`); memory unchanged.

## Option 2: a fingerprint byte a key (one more list in the page)

A list of n bytes behind the key lengths (many-key form only), one byte a slot: for a first slot a hash of its remainder (the length and the
first and last up to 8 bytes, multiplied by a constant, the top byte), for a further value a byte that is never compared (the key length
says `Further`). A lookup of a key (`Get`, `EachValue`, the add of a value to a key that is there, the remove of a value or a key) hashes
the rest of the key once, matches the byte against 8 fingerprints at a time with SWAR (`swar.Index8` / a mask of equal bytes, as the byte
nodes do), and compares only the remainders of the matches, on average one; the offset of a candidate's remainder is the sum of the key
lengths before it (a SWAR sum of the bytes, the `Further` slots masked to zero). A miss is decided without a compare in most cases. The
place of a **new** key and the bounds of a range still need the ordered search (`locate`), since a hash says nothing about order: the add of a
new key gets slower by the hash and by one more list to shift.

**Prediction:** lookups `u64` single-value 4,096 33 → **20 to 24 ns** (`main` 17), street single-value 58.6 → **45 to 50** (`main` 47); a miss
of a key 30 to 60 % faster; the add of a value to a key and every remove 5 to 15 % faster; the add of a new key 2 to 5 % slower; ranges
unchanged; memory **+1 byte a key** in multi-key pages (street single-value 27 → 28 block bytes a key, +3 to 4 %; a page holds that much
less, so a few more pages burst).

## Both

Option 1 makes every compare cheaper, also the ones option 2 leaves (the candidate, the place of a new key, the bounds of a range); they do
not exclude each other. In the order of cost: option 1 first (no layout change, every path profits, no memory), measured; then option 2 if
the lookups are still well behind `main`.

## How it is measured

The lookup probe against `main` (`lookprobe_test.go`: `u64`, street, dirs; single-value and natural; 4,096 and 65,536 keys), the write probe
(build and replay), the memory of the probe, the range probe (it must not get slower); then the cells of E5 on the PC (`main`, `node-layout`,
`btree`). Tests: the model tests of the page (`TestPageAgainstModel` and the fuzz) cover the search through every operation; a test of the
word compare at every length 0 to 24 and every position of the first difference, and at the end of the object.

## Option 1 result (2026-10-08): missed, dropped

Built as described (the key's first word once, one 8-byte load from the page a slot, masked to the shorter length; the rest word by word),
all tests and a new test of the compare at every length and difference (`TestPageCompareOrdersAsBytes`, kept) passing. Lookup probe against
`main` (WSL, medians of 9), ns a lookup, before → option 1: `u64` single-value 4,096 32.8 → **40.5** (+23 %), street single-value 57.9 →
59.6, dirs 86.4 → 86.5, `u64` natural 45.2 → 46.4, street natural 65.9 → 67.8, `u64` 65,536 35.4 → 36.5, street 65,536 85.6 → 87.1.
**Prediction missed in the sign** (predicted -10 to -20 %).

**Cause** (profile, `u64` single-value): the byte-wise compare was already cheap where it runs: random keys differ in the first byte, and
the remainders of a real page differ within the first bytes after the page's common prefix, so the loop ends after one or two bytes. The
word path adds work to every slot (a variable shift for the mask, `min`, a bounds-checked load with a byte swap, two more branches: the
mask and the compare lines alone took 19 % of the time). **The cost of the search is the number of slots it walks, not the width of a
compare**: each slot is a serial chain (its key length, the offset of its remainder, a load from there, a compare). The code is back as it
was; the lever left is to walk fewer slots: option 2 (a fingerprint, on average one candidate) or a search that does not walk (offsets of
the remainders, a binary search).

## The fingerprint: collisions within the pages of the real corpora (2026-10-08)

**Hash** (the user's proposal): wyhash's mixing as optimized in the Set3 project (`hashing.WH64Det`, 2.72 ns a call on the M1), on a block
of 16 bytes: the length of the input (2 bytes) and its last up to 14 bytes, right-aligned with zeros in front (the length tells `"a"` from
`"\x00a"`), as two words; the hash of the first word is the seed of the second. Measured in `Map.Shape` (tag mkstats, `shape_on.go`) on the
trees as built from the corpora: for every key of a multi-key page, how many **other keys of the same page** have the same fingerprint (the
compares a lookup would make in vain), against chance ((keys of the page − 1) / 128 or / 256). Three inputs: (A) the remainder of the key in
its page; (B) the whole key; (A') the remainder, but its first 6 and last 8 bytes when it is longer than 14.

| corpus | keys a page | (A) 7 bits | (A) 8 bits | (B) 8 bits | (A') 8 bits | chance 8 bits | keys with a twin (A, 8 bits) | most twins in a page |
|---|--:|--:|--:|--:|--:|--:|--:|--:|
| street natural 212,449 | 5.6 | 0.066 | 0.033 | 0.033 | 0.033 | 0.033 | 3.2 % | 3 |
| street single-value 212,449 | 7.1 | 0.092 | 0.045 | 0.046 | 0.046 | 0.046 | 4.4 % | 3 |
| dirs natural 86,215 | 4.9 | 0.061 | 0.038 | 0.038 | 0.044 | 0.025 | 3.4 % | 8 |
| dirs single-value 86,215 | 5.9 | 0.080 | 0.049 | 0.048 | 0.055 | 0.032 | 4.4 % | 8 |
| street single-value 4,096 | 7.1 | 0.092 | 0.041 | 0.038 | 0.042 | 0.043 | 3.9 % | 2 |
| dirs natural 4,096 | 4.0 | 0.044 | 0.024 | 0.031 | 0.022 | 0.017 | 2.3 % | 2 |

**What it says:** for street the fingerprint collides exactly as often as chance; for dirs about 1.5 times as often, because path keys of
the same length and the same last 14 bytes differ only in the middle — the worst page holds nine keys like
`ervice/mgmt/2019-08-01/containerservice/`, `…/2020-03-01/containerservice/`, …, which share length and ending (16 bytes of
`/containerservice/`). Taking the first 6 bytes instead of 6 of the last (A') does not help there (the dates differ at the seventh byte)
and is a little worse elsewhere. The whole key (B) is no better than the remainder (A), as expected (the path and the key part are the same
for all keys of a page). **In numbers:** with 8 bits a lookup compares on average 0.02 to 0.05 keys in vain (today: half the page, 3 to 4);
3 to 4 % of the keys have a twin in their page; the worst page (nine twins) is searched as today. 8 bits halve the collisions of 7.
