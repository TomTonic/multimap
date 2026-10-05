# Step 4.1: the multi-key page standalone

Written 2026-10-05. Design: [step4-mksv-design.md](step4-mksv-design.md) (approved by the user on 2026-10-05: D1 burst under byte nodes, D2 layout B with one grid, D3 one `build` function, order strings, `uint64`, pointers).
Words as in [GLOSSARY.md](GLOSSARY.md). Code: `internal/mkpage` (`page.go`: `Page` for strings; `fixed.go`: `Fixed` for pointer-free `T`, generic methods as `skpage.Fixed`), the model `bench/cmd/skmodel -multi`.

## What was built

- `mkpage.Page` and `mkpage.Fixed`, the layout of the note byte for byte: head of 3 bytes (type, n, cpl), one remainder length per entry, the common prefix once, then (strings) remainder, value length and value of each entry,
  or (fixed) all remainders, padding to the alignment of `T`, the values. Operations: `BuildStrings`/`BuildFixed` (from sorted entries, the common prefix is the longest all keys share), `Get`, `Insert` (answers `Added`, `Present`, `Differs`,
  `Full`, `Outside`), `Remove` (shrinks with the hysteresis of step 3.5: the content fills at most half of the smaller class), `Each`, `Skip` (a byte node arrives above: the prefix loses its first bytes in place), `Prepend` (the node above
  goes away), `NeedStrings`/`NeedFixed` (what a set of entries takes, so that the tree decides page or byte node before it builds anything). Growth copies the used bytes into the next class: the layout does not depend on the class.
- **Gates:** `internal/mkpage` 100 %, `-race`, `FuzzPage` and `FuzzFixed` 20 s each (3.6 and 2.7 million runs, no finding), lint 0. Tests: both pages against a sorted-list model for 40 seeds of 600 random inserts and removes
  (strings; `uint8`, `uint16`, `uint64`, a pair with alignment 4, two words), limits (refused and accepted at the borders), every `Result`, growth through all classes, hysteresis, `Skip`/`Prepend`, early stop of scans.
- The 255-entry limit of the head byte cannot be reached by `Insert`: an entry takes three bytes at least, so 512 bytes hold at most 170. Only the builders look at it.

## Against the prediction

1. **Memory: the model and the package agree to the byte.** `skmodel -multi` builds every page of the model with `mkpage` (strings and `uint64`; `street`, `dirs`, random `uint64` keys; real mix and one value per entry), checks every entry
   with `Get` and adds up the sizes: **10 of 10 data sets give exactly the bytes of the model, no page refused, none of another class** (for example `street`, one value, `uint64`: 4,593,888 B in 27,887 pages; model 4,593,888 B). So the prediction of the
   note (26.9 B an entry for `street`, 39.5 `dirs`, strings 30.9 and 50.3) is a statement about this layout and holds as far as the tree builds the same subtrees.
2. **Cost per operation** (`go test -bench`, WSL on the Ryzen 9 7900, 4,096 pages per set; indicative, the PC native and the M1 follow with the tree). A page of 7 entries is the average of the model (6.5), 20 is the upper part (up to 29).

| ns | n=3 | n=7 | n=20 |
|---|--:|--:|--:|
| `Get`, key there (strings / `uint64`) | 16.4 / 16.8 | 23.8 / 23.4 | 45.7 / 42.7 |
| `Get`, key not there | 9.3 / 9.1 | 15.6 / 14.3 | 40.2 / 31.6 |
| `Insert` and `Remove` of one entry, room in the page | 43 / 52 | 55 / 64 | 102 / 84 |
| `BuildStrings` | 71 | 123 | 256 |
| the single-key page: `Match` of its one key | 3.1 | | |

   About 2 ns an entry walked: the lengths are read one after the other. **A first version used `bytes.Compare` per entry: 30 / 67 ns for the key there at 7 / 20 entries, 29 / 96 ns for the key not there.** Replacing it by a loop
   that stops at the first differing byte (the remainders differ in the first bytes) took 20 to 60 % off (`compare` in `page.go`); that is the only tuning done, and it has a cause: a call of a library routine against a few compares.
3. What this says for the tree (a conclusion, not a measurement): a lookup in a page of 7 entries costs about 20 ns more than the one compare of a single-key page; the descent saves the object per entry: one cache miss (50 to 100 ns when
   the page is not in the cache) per lookup that now finds its page together with its neighbours. The weak spot stays small maps and absent keys in full pages, as in step3-layout.md; the pages of 20 entries and up are where the linear walk
   shows (45 ns). Whether a first-byte filter (a byte per entry in the head) pays is a question for the tree measurement, not for this step.

## The likely numbers of entries (the user's question after 4.1)

The 255 (really 170) entries are a bound of the head byte that no data comes near: in the model of `street` with one value 90 % of the pages hold 16 entries or fewer, the average is 6.5 (7.1 with `uint64`), the largest 29.
What would a limit that fits the likely numbers cost? `skmodel -multi -mkmaxn N` (a page with more entries bursts) and `-mkgrid` (the largest class), bytes an entry, one value per entry, model:

| limit | `street`, strings | `street`, `uint64` | `dirs`, strings | `dirs`, `uint64` |
|---|--:|--:|--:|--:|
| none (512 bytes) | 30.9 | 26.9 | 50.3 | 39.5 |
| at most 24 entries | 31.2 (+1 %) | 27.8 (+3 %) | 50.3 | 39.6 |
| at most 16 entries | 33.4 (+8 %) | 29.9 (+11 %) | 50.7 | 40.5 |
| at most 12 entries | 35.1 (+14 %) | 31.7 (+18 %) | 51.9 | 42.0 |
| largest class 384 | 32.5 (+5 %) | 28.3 (+5 %) | 52.0 | 40.7 |
| largest class 256 | 34.8 (+13 %) | 30.4 (+13 %) | 55.1 | 42.6 |

So the cost of the walk and of the shift (about 2 ns an entry for a lookup, a memmove of at most 512 bytes for an insert) can be bounded at 24 entries for 1 to 3 % of memory on `street` and nothing on `dirs`, but not much lower: at 16 the pages burst into more, smaller ones
and the nodes above them grow (`street`: 3.7 to 5.0 bytes an entry). Nothing in the layout depends on 255, and a head of two bytes (n and cpl in one) would save one byte a page, 0.15 bytes an entry. **No change now**; the cap of 24 is a free
option for 4.2 if the tree measurement shows the walk in the profile.

## Not done / next

- The pointer page (`*T`) is step 4.3. `Skip` and `Prepend` have no allocation-free path for a class change; `Prepend` copies once.
- 4.2: the tree. `build` (page or byte node, from sorted entries), burst, promote, removal and merge, scans; the types of the pages in `internal/art/node.go` (three ranges: single-key pages, multi-key pages, nodes). **Stop for the user** before it.
