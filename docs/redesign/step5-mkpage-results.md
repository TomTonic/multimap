# Step 5.1: the page with several values a key (`internal/mkpage`), results

Design: [step5-mkmv-design.md](step5-mkmv-design.md). Words as in [GLOSSARY.md](GLOSSARY.md). The code is `internal/mkpage` (both flavors) and the small changes in `internal/art` that keep the tree running on it; **the tree
does not yet put a second value into a page** (that is 5.2): `Insert` still answers `Differs`, as the tree expects.

## What is built

- **Format** as decided: head of three bytes `type | cpl | n` (the lowest bit of `type` is bit 8 of `cpl`, so a common prefix of up to 511 bytes: `MaxPrefix` 255 → 511), the common prefix right behind the head,
  then the length list, `Further` (255) for a further value of the key before it (no remainder), `MaxRemainder` 254. The bytes of the design note's example are pinned in `TestMultiValuePageLayout` (class 128 for `uint64` values,
  class 64 for strings: used 62 bytes, not the 63 of the note's hand count).
- **API** (both flavors): `Add` (new key `Added`, further value `AddedValue`, `Present`, `Full`, `Outside`), `Insert` (the old answers, `Differs` for another value; goes when the tree uses `Add`), `Remove(key, value)` (any value of the
  key: the middle or last value takes out its length byte and slot; the first value of a key with others lets the next one take the key's place; the last value takes the remainder too), `EachValue` (the values of a key in
  the order they came in), `Each` (with a `first` flag; a further value reports its key's remainder), `Keys()`, `BuildStrings/BuildFixed` (a key with several values appears once for each), `Widen` (heads grow, `Further` stays; refuses a
  remainder beyond 254), `Skip`/`Prepend` (the prefix sits at `Header` now).
- **Tests:** `TestMultiValuePageAgainstModel` (60 seeds × 800 random adds, removes and widenings against a model, both flavors), the layout, removals in every order, `EachValue`, build limits, `Skip`/`Prepend`, the nine-bit prefix
  (with `TypeBase` 0 and 20: odd type byte, growth into the next class, `Skip` below 256 and `Prepend` back), `Widen` into a 300-byte remainder, `FuzzMultiValuePage` (40 s, 2.4 million runs, nothing found). `internal/mkpage` **100 %**, race,
  lint 0. In the tree: `maxPageByte = kLastMultiKey|1` for `isPage` (the odd type byte of a page with a long prefix), and `TestMultiKeyPageLongPrefix` (three keys with 300 bytes in common are one page; a node comes and goes above it).

## Page level, against the page of step 4 (hot cache, WSL, median of 4 runs; below 1.00 is faster)

| | n=3 | n=7 | n=20 |
|---|--:|--:|--:|
| `Page` Get hit | 1.03 | 1.04 | 1.03 |
| `Page` Get miss | 1.09 | 1.12 | 1.05 |
| `Page` insert and remove | 1.16 | 1.17 | 1.12 |
| `Fixed` Get hit | 1.08 | 1.06 | 1.06 |
| `Fixed` Get miss | 1.11 | 1.08 | 1.12 |
| `Fixed` insert and remove | 1.11 | 1.11 | 1.09 |
| `Fixed` scan with `Each` (ns a value) | | 2.3 → 3.3 | 2.0 → 2.7 |
| `Fixed` scan, loop without a function per entry | | 1.0 → 1.5 | 0.84 → 1.1 |

**This misses the prediction** (Get within 5 %, `Each` at most 2.5 ns a value): Get is 3 to 12 % slower, insert and remove 9 to 17 %, the scan 30 to 45 %. One run each on WSL (the scatter is a few percent), but the order is clear.

## Cause, as far as understood

The layout has no cost in bytes, but **the position of the values**. In `Fixed` the values array starts behind the last remainder, so every Get and every `Each` first adds up the length list to find where it starts (`keyEnd`; the page of step 4 did
`sum += rl`). With `Further` the sum has to leave the 255s out: two adds (`sum`, and a count of the 255s) where there was one. The tuning I tried in the same step (a sub-slice for the length list so that the bounds checks go, the
count instead of a mask) took Get miss from 1.24 to 1.12 and no more. The scan's `loop` benchmark (which also calls `Used`) moved 7.3 → 10.6 ns for 7 values: the sum is about three nanoseconds of the seven that `Each` of a page of 7 lost (16.0 → 22.8). The string page has
no such sum (its values are behind their keys), and its Get is 3 to 5 % slower only.

## Options (for the user; none is built)

- **A. Accept and measure in the tree** (5.2). A page is one object in a descent of three or four; 3 to 12 % at the page may be 1 to 3 % of an operation. The risk is the scans, where the page is most of the work.
- **B. Values at the end of the object (built, see below)** (`Fixed` only). The array of values sits at `Size - n*w`, not behind the remainders: **no sum is needed to read** (Get and `Each` read `m[Size-(n-pos)*w]`); only a change of the page adds up the lengths. It is also what the typed
  pointer pages of 5.3 are anyway (`ptrObject[T, [J]uint64, [N]T]` has its `N` slots at the end), so both flavors would have one layout. It costs nothing in bytes (the padding before the values becomes a gap of zeros, and a page needs `keyEnd + n*w`,
  not `align(keyEnd) + n*w`: on average a little less, some pages one class smaller). Prediction: Get and `Each` as before the change or better (`Each` 2.0 ns a value or less, the loop 0.9), insert and remove as now (the values array moves
  by one slot at an insert at the front, as it moved before); memory unchanged to a few tenths of a byte an entry.
- **C. Cache the end of the keys in the head**: a fourth byte (nine bits do not fit the one) for both flavors; it belongs to the question of the shared head (step 5.5) and costs a byte a page (+0.15 B an entry).
- **D. The 8-byte sum with SWAR** (eight length bytes at a time): helps `n` above 8, costs a page-edge case at the end of the object; not before B or C show they are not enough.

My proposal was **B**; the user agreed (2026-10-06). Built the same day, results below.

## Option B built: the values of `Fixed` at the end of the object

`Fixed.vs() = Size - n*w`; reads (`Get`, `EachValue`, `Each`) never add up the length list, only changes do (`keyEndFrom`). The class of a page needs `keyEnd + n*w` (no padding to the alignment of `T`: the object's size and the array's size are multiples of
what `T` needs). All 100 % tests pass, race, fuzz; the one bug of the change was found by the model test at once (the end of the keys must be measured before a slot is turned into the key's). Page level against the page of step 4 (WSL, median of 4):

| `Fixed` | n=3 | n=7 | n=20 |
|---|--:|--:|--:|
| Get hit | 1.01 | 0.96 | 1.00 |
| Get miss | 1.00 | 1.09 | 1.21 |
| insert and remove | 1.00 | 1.00 | 1.03 |
| scan with `Each`, ns a value | | 2.3 → 2.6 | 2.0 → 2.3 |
| scan, loop without a function per entry, ns a value | | 1.04 → 0.97 | 0.85 → 0.65 |

**Gets and changes are back at the page of step 4 (the prediction); the scan with `Each` is 15 % slower (2.3 ns a value for 20 entries, within the 2.5 of the prediction; for 7 entries 2.6); the plain loop is faster than before. The one miss is `Get` of an absent key in a page of 20 entries
(+21 %: one more compare of the length byte against `Further` for every slot).** Memory is unchanged to the byte (model 26.9 / 39.5 / 24.0; `objstat` 28.1, 41.0, 24.0 at the sizes that match). The pages of strings (no sum to find their values) are as in the first table: Get +3 to 12 %,
insert and remove +9 to 17 % (not changed in this step).
