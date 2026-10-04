# Step 3.4b: a remainder of up to 510 bytes, nine bits of length in two adjacent bytes

Status: **design, before code** (2026-10-04; the idea and the encoding are the user's). Words as in [GLOSSARY.md](GLOSSARY.md).

## 1. What and why

Today `leafHead.klen` is one byte, so a leaf or page holds a remainder of 254 bytes at most (255 = `longKey`: the whole key is held as a string, one
more object and one more pointer chase per comparison). With one more bit the remainder is 9 bits long: **inline up to 510 bytes**, 511 = `longKey`.
What it buys: the set leaf of a string map gets the classes **384 and 512** (key area 370 and 498 bytes), and a page can hold a remainder up to the
largest page. It costs nothing in memory on `street` and `dirs` (their remainders are short: the byte nodes hold the common prefixes); it is for the
data of the backlog (long paths, long keys such as URLs) and for the grid, which should have no hole between 242 and 510.

## 2. The encoding

The two bytes after each other are `kind` (byte 0) and `klen` (byte 1), as in the user's design: the **lowest bit of the kind byte is bit 8 of the
remainder length**. Read as a 16-bit number with the kind byte on top, `(kind<<8 | klen) & 0x1FF` is the length and `kind &^ 1` the type.

- **Kinds step by two** (`kSet = 2`, flat/typed leaf of class c = `kSet+2c`, pages likewise, nodes 2 apart); a leaf or page of a long remainder has the
  kind byte `+1`. Byte nodes never set the bit.
- **Type test:** `kind&^1 == kSet` (the user's `&0xFE`); **leaf or page**: still *one range comparison*, against the largest byte value of the class,
  `maxLeafByte = kLastLeaf|1`, `maxPageByte = kLastPage|1`, so the descent in `find`, `scan`, `insert` and `delete` costs nothing more.
- **Class** of a leaf: `(kind&^1 - kSet) >> 1`. **Tables indexed by kind** (`endPageOff`, `fixedSize`, `slotCap`, `kindMask`) are indexed by `kind>>1`.
- **Length** (`leafHead.rem()`, `setRem`): `int(klen) | int(kind&1)<<8`; one load of the two bytes, a rotate and a mask on a little-endian machine. The
  compiler is left to it; the accessors stay small enough to inline (checked with `-gcflags=-m`, `matches` must still inline into `find`).
- **`maxInline`** of the generic leaves (`Map[T]` other than strings: a key area of 16 to 256 bytes) stays 254 and is no longer tied to `longKey`; `longKey`
  becomes 511.

- **Why for every class and not only for 384 and 512** (user's question): the bit can only be set in an object of 384 or 512 bytes (a remainder of
  256 bytes needs 262), but `rem()` is the same branch-free expression for all of them, and the kind byte is already loaded; a rule by class would be
  a compare and a second path in `matches`. In the small classes the bit is just 0.

## 3. What changes (all of it)

| where | change |
|---|---|
| `internal/art/node.go` | kinds step by 2, `maxLeafByte`, `maxPageByte`, `isLeaf`, `isPage`, `cls`, `rem`/`setRem`/`isSet`, `longKey` 511, table indexes |
| `insert.go`, `lookup.go`, `scan.go`, `delete.go` | `<= kLastPage` and `> kLastLeaf` use the byte maxima (8 places); nothing else: node kinds are even and unchanged |
| `flat.go`, `typed.go`, `map.go`, `skleaf.go`, `objects.go` | `l.klen` becomes `l.rem()`, assignments `setRem`; `kSet + kind(c)` becomes `kSet + 2*kind(c)`; `l.kind == kSet` becomes `l.isSet()` (11 places) |
| `rnode.go` | range node class from `(kind-kR8)>>1` |
| `internal/vpage`, `internal/skpage`, `internal/lpage` | `KindBase + 2c` and `class()` from `>>1`; the single-key page reads and writes its remainder length in nine bits; `MaxRemainder` 505 (a page of 512 bytes holds 6 bytes of header and one length byte) |
| set leaf of a string map | key areas 18, 50, 114, 242, **370, 498** (32, 64, 128, 256, **384, 512** bytes); a remainder above 498 as a string |
| tests | the kind arithmetic of the tests, a remainder of 255 to 510 bytes in every leaf and page kind, in the fuzz test and the layout test; coverage stays 100 % |

## 4. Prediction

- **Memory:** unchanged on `street` and `dirs` (no remainder above 254 there; the objects keep their size). Longer remainders cost what they hold and do not
  fall back to a string any more: a key of 300 bytes held by a set leaf of 384 bytes instead of 32 + 320.
- **Speed:** `matches` and `stored` read one more bit (two instructions) and the descent is unchanged: below 1 ns on a lookup of 80 to 260 ns, so not
  measurable in the benchmarks; `go test -bench` of `Find` on the PC (the kernel of the lookup) checks that it does not move by more than the noise.
- **Check:** the whole test suite on the new encoding, `-race`, fuzz, coverage 100 %, then the measurement of 3.4 again (the numbers of the memory runs
  with the grid), as the next step.

## 5. Done (same day)

Built as described; `matches` and `rem` still inline (15 call sites of `matches` inlined, 11 before). The page's remainder is at most 505 bytes (`skpage.MaxRemainder`);
while doing it a latent fault of `skpage.New` showed: a remainder and a value that together need more than 512 bytes (254 + 254) gave class -1 and
a crash instead of nil; it returns nil now. Tests: remainders from 254 to 600 bytes in the page, the set leaf of every class, and the rekey of a set leaf in
each key area.
