# Step 3.1: design note of the single-key page (SKMV) for string values

Status: **for the user's approval, no code yet** (PLAN.md, step 3.1). Written 2026-10-04. Words as in
[GLOSSARY.md](GLOSSARY.md). Numbers from `bench/cmd/skmodel` (a model on the real keys and values, no
tree is built) and from [step3-skmv-sizes.md](step3-skmv-sizes.md).

## 1. What this note decides, and what it leaves

Decides: the layout of the single-key page for **variable-length values** (`Map[string]`, the base case
`string -> {string}`), its size classes, its operations, how it enters the tree, and what the
measurement of 3.3 has to show.

Leaves for later notes: fixed-size values (`uint64`) and values with pointers (`*T`) are step 3.5; the
value overflow (a key with more values than a page holds) is step 3.4 and stays, for step 3.3, what it is
today: the set leaf with a value set. The multi-key page (MKSV) is step 4, multi-value entries in multi-key
pages (MKMV) step 5.

## 2. The page

One object per key, **no pointer in it**. It holds the key's remainder and all its values as bytes.

```
offset 0   kind        1 B   the size class: kSK32, kSK64, kSK128, kSK256, kSK384, kSK512
offset 1   n           1 B   number of values, 1 to 255
offset 2   r           1 B   length of the remainder, 0 to 254
offset 3   remainder   r B   the key from the page's path length on
offset 3+r value 1     1 B length l1 (0 to 254), then l1 bytes
           value 2     1 B length l2, then l2 bytes
           ...         up to value n; the rest of the object is unused
```

Example, `street` p50, from the model: key "Schlosspark Wendorf", the tree consumes the path up to and
including the branch byte, the remainder is 4 bytes ("dorf"), one value "Möllenhagen" (12 bytes in UTF-8):
`kSK32 | 1 | 4 | d o r f | 12 | M ö l l e n h a g e n` = 3 + 4 + 1 + 12 = 20 bytes in a 32-byte object.

Why this and not the sketch of [whataleafneedstostore.md](whataleafneedstostore.md) (all lengths in the
header, values behind the remainder):

- **The lengths sit in front of their values.** Adding a value is an append (no shifting of the remainder
  and the values), lookup of a value walks the lengths and the bytes in one pass over contiguous memory, and
  the number of values is not limited by a header with room for six or seven lengths (the user's decision:
  "what fits"). The price is that the i-th value cannot be found without walking; no operation needs
  that: the values of a key are a set, read as a whole.
- **The header is 3 bytes**, the sketch's 8. Every key costs 5 bytes more with 8: in the model the share of
  `street` keys (natural mix) in the 32-byte class falls from 84.8 % to 76.7 % and the pages from 44.8 to 47.9
  B/key; for `dirs` from 49 % to 37 % and from 76.3 to 82.2 B/key.
- **The remainder comes first**, because a lookup first decides whether this is the key.

Limits that make every length one byte, and what happens beyond them (all of them go to the value
overflow of today, the set leaf; see 5): a value of 255 bytes or more, more than 255 values, a content of
more than 512 bytes. A remainder of 255 bytes or more also goes to the set leaf, which holds a key inline
up to 254 bytes and longer ones as a string (this is today's behaviour; the *oversized object* of the glossary
is not built in this step).

**Size classes** (decided 2026-10-04): 32, 64, 128, 256, 384 and 512 bytes. 32 and 64 are Go size classes
that are aligned to their size, so an object of 64 bytes or less lies in one 64-byte cache line (two 32-byte
pages share one); 128, 256, 384 and 512 are multiples of 128 as before. R1 gets an exception for the
single-key page: 32 and 64.

## 3. Operations

All are on one page; the tree finds the page and tells it the remainder of the key and the path length.

| operation | what it does | cost |
|---|---|---|
| `Has(rest)` / `Values(rest)` | compare `r` with len(rest) and the remainder bytes with `rest`; then walk the values | one pass over at most 512 bytes |
| `Add(rest, v)` | `Values` found the key: walk the values, return if `v` is there (set semantics); else append `1+len(v)` bytes at the end, `n++`; if it does not fit the class, copy to the next class that holds it; if none (512) or `n==255` or `len(v)>=255`, convert to a set leaf | append in place; a copy when the class changes |
| `Remove(rest, v)` | walk to the value, shift the values behind it down, `n--`; if `n==0` the key is gone; shrink (below) | one memmove of at most 512 bytes |
| `New(rest, v)` | the smallest class that holds 3 + r + 1 + len(v) | one allocation |
| `Skip(k)` | the page moves k bytes deeper (a new node was split in above it): the remainder loses its first k bytes, everything behind it shifts down; in place | one memmove |
| `Prepend(pre)` | the page moves up (the node above it went away): the remainder gets `pre` in front; copy to a larger class if it no longer fits | a copy |
| `Each` / key assembly for scans | the remainder is appended to the path; the values are handed out as below | |

**Growing and shrinking.** A page grows to the smallest class that holds the new content. It shrinks, on a
removal, to the smallest class that holds it if that saves at least half of its size (the sketch's rule): a
128-byte page whose content falls to 64 or less becomes 64; a 384-byte page becomes smaller only if the content
fits 128 or less. The page the garbage collector frees is not recycled (the sketch's LIFO page cache is not built;
Go's allocator is the cache).

**From the page to the set leaf and back.** The conversion to a set leaf happens when the content no longer
fits 512 bytes, `n` would be 256, or a value is 255 bytes or longer. The set leaf goes back to a page when all
its values fit a page of 256 bytes (half of the largest class, so a key at the boundary does not flap between
the two on every add and remove) and every value is shorter than 255 bytes.

**Values given out.** The page holds bytes; the caller gets `string`s. A lookup or a scan allocates **one
buffer per key** (the value bytes of the key together, without the length bytes) and hands out the values as
substrings of it. A string the caller keeps holds the buffer of its key alive. No zero-copy (decided
2026-10-04). Adding a value copies the caller's bytes into the page; the page keeps no reference to the
caller's string.

## 4. Rounds of cache-line loads (R5)

The descent to the page costs what it costs today (nodes). In the page: round 1 reads the first line, which
holds the header, the remainder (the median remainder is 4 to 10 bytes) and, for a content up to 64 bytes
(95 % of the `street` keys, 81 % of the `dirs` keys, 100 % and 98 % with one value) everything. Where the
content is longer, the addresses of the further lines are known after round 1 (the position is the sum of
the lengths walked), and they are read together: round 2. R5 holds.

## 5. How it enters the tree

- **Only for `T = string`** in this step. `Map[T]` decides at its first use, as it decides between flat and
  typed leaves today (`decide`), and a string map gets the kinds `kSK32` to `kSK512` instead of typed leaves.
  Every other `T` keeps its leaves until 3.5.
- **The set leaf stays** as the value overflow and for keys with a remainder of 255 bytes or more. It is
  redesigned in 3.4, not here.
- **No multi-key pages and no fall back in this step** for a string map (decided 2026-10-04): the tree is byte
  nodes and one single-key page per key, plus range nodes nowhere. This needs a switch in `Tree` (pages and
  range nodes off); today the tree creates pages whenever a key with one value arrives.
- **Where the code changes**, by the functions that touch a leaf: key match in `find`/`upsert`/`del`,
  `splitLeaf` and `rekey` (Skip/Prepend), the key assembly in `scan.go` and `rebuild.go`, `Each`, `Objects`
  (a new label per class). They dispatch on the kind; the new page does not share `leafHead` (its header is 3
  bytes, the leaf's is 6).
- **Rounds of work:** 3.2 builds the page alone (`internal/skpage`: the table above as functions on a byte
  slice, 100 % coverage, a fuzz test against a model, microbenchmarks on the real entries), 3.3 wires it in.

## 6. Prediction (to be compared with the measurement of 3.3)

Model of the tree of byte nodes (the nodes of `node-layout`) with a page per key, on the whole corpora:

| data | nodes (a key) | pages | value overflow (est.) | total B/key |
|---|--:|--:|--:|--:|
| street, natural mix | 34.4 | 42.8 | 38.6 | **115.8** |
| street, one value | 34.4 | 32.8 | 0 | **67.2** |
| dirs, natural mix | 37.7 | 70.2 | 46.0 | **153.9** |
| dirs, one value | 37.7 | 42.2 | 0 | **80.0** |

The string bytes of the values are in the pages. The value overflow row is a rough estimate (the value sets
of 0.5 % of the `street` keys and 1.6 % of the `dirs` keys, which hold 37 % of the values; 24 or 40 bytes a
value). Compared with what is measured today (the tree of step 2: 119 and 149 B/key in the natural mix
without the string bytes, 147 and 217 with them; 67 and 78 with one value): **expect between -20 % and -30 %
in the natural mix, little in the one-value profile** (the pages replace a leaf of about the same size; the
nodes, a third of the total, do not change). The number to compare against is the memory of `node-layout`
from the reference runs (to be filled in when they are in).

Expected speed against `node-layout` (the leaf per key it replaces):

- **Point lookups:** the same descent; the leaf is one line either way; the string bytes are in that line
  instead of behind a pointer (a second miss when the caller reads them); against that, one allocation of
  about 28 bytes per `ValuesFor` for the copy (a leaf hands out the stored header). Expected: even, within
  10 % either way. **This is the risk for credo 1** (lookups against `btree-sets`).
- **Range scans:** one miss less per key where the caller reads the bytes; the pages are not contiguous
  (one object per key, in allocation order), so no gain from layout. Expected 1.1 to 1.5 times.
- **Mutation:** a page is a 32 to 512-byte object that is copied when its class changes; adding a value
  is an append. Expected: even for `churn`, `build` within ±15 %.

If a measurement deviates by more than about 10 % from this, the cause is found and written down before
anything else changes (PLAN.md).

## 7. For the user to decide

1. The layout of section 2: remainder first, then length-prefixed values, a 3-byte header, one-byte lengths
   (so: value < 255 bytes, remainder < 255 bytes, n < 256 per page), the rest to the set leaf. Alternative with
   all lengths in the header as in the sketch: +5 bytes a key, a fixed number of values.
2. The classes 32, 64, 128, 256, 384, 512 (decided), shrink only if that saves half, no page cache.
3. Convert back from a set leaf at a content of 256 bytes or less.
4. One allocation per lookup or scan of a key for the strings (copy-out), no zero-copy.
5. In 3.3 the tree is byte nodes plus one page per key; no multi-key pages, no fall back for string maps.
