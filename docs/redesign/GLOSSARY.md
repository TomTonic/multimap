# Glossary

The vocabulary of the redesign, agreed with the user on 2026-10-03, brought up to date on
2026-10-04. Use these words in code, comments and documents. Each entry says what the word names,
and what it was called before. A word marked *planned* names something that the target design needs
and the code does not have yet; the code still has the old object under its old name.

**State of the code (2026-10-04):** the code does not speak this vocabulary yet. `internal/art` still
says `leaf`, `term`, `depth`, `inner`, `settle` and `crowded` (about 900 places), `internal/lpage` and
`internal/vpage` say `suffix`. That two vocabularies were in use is one reason the work stumbled.
Step 3.0 of [PLAN.md](PLAN.md) renames what survives step 3; new or changed code uses only the words
below.

A definition that needs more than two sentences is a warning sign: the concept may be too
complicated. Such places are collected at the end ("Where the vocabulary is still heavy").

## The four cases of the user's sketch

The user's sketch ([whataleafneedstostore.md](whataleafneedstostore.md)) names objects by how many keys
and how many values they hold. The order of the work follows them (PLAN.md, 2026-10-04):

| short | long | object | step |
|---|---|---|---|
| SKSV | single key, single value | a single-key page with one value | 3 (special case of SKMV) |
| SKMV | single key, multiple values | **single-key page** | 3, the base case |
| MKSV | multiple keys, single value each | **multi-key page** | 4 |
| MKMV | multiple keys, multiple values each | a multi-key page with multi-value entries | 5, last |

## Keys, values, entries

- **key**: a byte string. The tree orders its entries by key, bytewise.
- **value**: what a key maps to: a `T` of the user's choice. The values of a key form a set: adding
  a value that is there changes nothing. A page stores it in one of two ways
  *(planned for the single-key page, step 3.1)*:
  - **variable-length value**: its bytes and their length, inline (`string`). A lookup copies the
    bytes out (decided 2026-10-04). Up to 255 bytes while the length is one byte; a longer value
    needs an escape (open).
  - **fixed-size value**: the `T` itself in an array (`uint64`, pointers, any fixed `T`). If `T` holds
    pointers, the page is an object of a Go type the garbage collector scans.
- **entry**: a key with all its values. The unit that objects store.
- **single-value entry**, **multi-value entry**: an entry with exactly one value, or with two or
  more. (Before: single-value or multi-value *key*.) The bench's value profiles are `single-value` (every entry has
  one value) and `natural` (the real mix of both); until 2026-10-04 they were `unique` and `multi`, and results
  from before keep those names. SKSV, SKMV, MKSV and MKMV name page types, not workloads.

## Where a key is cut: path, remainder, common prefix

Every object that stores entries stores only the end of their keys. The tree spells the beginning.

- **path** (of an object): the bytes of the key that the descent has matched when it reaches the
  object: the branch bytes and the common prefixes of the nodes above. The object does not store
  them. (Before: `depth`, `base`, "the path to it".)
- **remainder** (of an entry in an object): the rest of the key, which the object stores. Key =
  path + remainder. (Before: `suffix` in pages, `remainder` in leaves.)
- **path length**: the number of bytes of the path; the remainder starts there. (Before: `base`,
  `depth`.) *Exception, in `internal/art` only:* a page of step 2 may start its remainders a few bytes
  earlier than the path ends, so that every key has a whole head word in the page (`pageBase`);
  those bytes are stored twice. The experiment of layout B in the tree did without it.
- **common prefix**: the bytes that all entries below a node, or in a page, have in common
  after the path to it, stored once there and compared in one step. This is what the ART
  literature calls *path compression*, with the bytes stored as the node's *prefix*. (Before:
  `path`, `compressed path`, the header's `prefix`, the page's `prefix`, `cp`; for the part beyond 12
  bytes in a node: `path tail`.)
- **prefix scan**: the user's operation that visits all entries whose key starts with given bytes.
  Always these two words. "Prefix" alone means the common prefix.

## Objects

- **leaf** (the abstract word): what a descent ends in, the key with its values. It is a concept of the model, **not an object**: whatever it is in
  memory is a **page**. So "leaf" does not name a type in new code, and "leaf or page" is not a distinction (the user, 2026-10-04: "leaf is the
  computation-independent model, everything is mapped to pages"). The code of `internal/art` still has the old object under that word: the flat leaf and the typed leaf
  (`leaf[T,K]`, `findLeaf` ... ) are single-key pages in older forms and keep those names until step 3.5 replaces them. **Renamed 2026-10-04** (commit 622d634):
  `leafHead` -> `singleKeyHead`, `isLeaf`/`isPage` -> `isSingleKey`/`isMultiKey`, `kSet` -> `kValueOverflow`, `kind` -> `objType` (the type byte, byte 0 of
  every object), the set leaf -> value overflow.
- **page**: an object of 32, 64, 128, 256, 384 or 512 bytes that stores entries (the grid of
  [step3-skmv-sizes.md](step3-skmv-sizes.md), decided 2026-10-04 for the single-key page: most keys hold 20 to 60 bytes; decided 2026-10-05 for the multi-key page too, step 4:
  [step4-mksv-design.md](step4-mksv-design.md), a page of two entries takes 64). It has no pointer in it unless its values need
  one. There are two kinds, and until step 5 no page mixes them: no page holds entries with
  different numbers of values.
  - **single-key page** *(planned, step 3)*: exactly one entry with all its values inline, as many as
    fit (not a number fixed by the header). (Before: *flat leaf*, *typed leaf*, and the inline part
    of the *set leaf*. SKMV; SKSV is the case of one value.)
  - **multi-key page** (changed 2026-10-06, step 5): many entries, each with one or more values, in key order; a page whose entries have one value each is the case of step 4. Two layouts existed as
    prototypes: **layout A** (`internal/vpage`, a directory of tags; in the tree since step 2) and
    **layout B** (`internal/lpage`, a header of lengths, the user's sketch). (Before: `vpage`, `page`,
    leaf-page, U8 page. MKSV.)
  - **end page**: the page that a node holds for the entry whose key ends exactly at the node,
    which happens when that key is a prefix of the other keys below (`ab` beside `abc`). It is a
    single-key page. (Before: `term`, `term leaf`.)
- **oversized object** *(planned)*: a single-key page whose remainder does not fit the largest
  size class. It stores the remainder inline anyway, in an object of the size Go gives it (576 bytes
  and up). It is exempt from R1 and counted separately in the object statistic. It saves the
  random cache miss that a pointer to the key would cost. (Before: *key overflow*, a set leaf with the
  key as a string.)
- **value overflow** *(planned)*: when the values of a single-key page do not fit it, the page
  holds a pointer to a **value set** instead of the values. (Before: the *set leaf*.) On the
  real data 1.5 to 2.5 % of the keys need it, and they hold about half of all values.
- **value set**: the object of a value overflow. Today the value set of `internal/vset`: an
  array (up to 64 values), then a hash set. (Before: `vset.Set`, with its inline stage of 3 values,
  and `Set3` for the hash stage.) Whether value blocks of 128-byte multiples do better is step 3.4.
- **node**: an object that routes: it picks the child for one byte of the key. It stores no
  entries; the entry whose key ends at the node is in its end page. There are two kinds, which
  differ only in how fine they cut the byte:
  - **byte node**: one child per byte value. The child never sees that byte again; the node
    consumes it. (Before: `inner node`, N5 to N256.)
  - **range node**: one child per range of byte values. The byte stays in the keys below, because a
    child holds several values of it. (R8 to R256.) *Not part of the tree since 2026-10-05* (step 4 hangs the multi-key pages below byte nodes); whether it comes back is the question of step 6 (routing).

  A byte node is the special case of a range node with one value per range, plus the consumed byte.

## Memory

- **size class**: one of the sizes our objects have: pages 128 to 512 bytes, nodes 128 to
  512 (R256: 2112), and so on. (Before: `class`, used in six tables.)
- **block**: the size the Go allocator really gives an object (`art.Block`), which is the next
  Go size class.
- **round**: a set of cache-line loads whose addresses are all known before the first completes
  (R5, [STRATEGY.md](STRATEGY.md)).
- **R1 to R6**: the binding rules of [STRATEGY.md](STRATEGY.md), section 2.

## Operations

- **grow**, **shrink**: a page moves to the next larger or smaller size class (a new object).
- **burst** (changed 2026-10-05, step 4): a full multi-key page is replaced by a byte node on the next byte in which its keys differ, with a page (or, for one entry, a single-key page)
  for each byte value. The tree has no range nodes in step 4; *before:* a full page was split at a byte boundary, or got a range node of its own.
- **promote** (changed 2026-10-05; **goes with step 5.2**): an entry that gets a second value leaves its multi-key page: the page's entries are built again with that entry as a single-key page
  (`build`, the function that also does the burst). In step 5 a second value is a **continuation** in the page.
- **slot**, **continuation** (2026-10-06, `internal/page`): a slot is one value of a multi-key page with its length byte; a continuation is the slot of a further value of the key before it: its length byte is `Further` (255) and it has no remainder.
- **merge**: after a removal, a byte node whose children are single-value pages that fit one page together becomes one page.
- **fall back**, **split**: retired (2026-10-05). A multi-key page never holds a multi-value entry, so no subtree needs a fall back; there is no range node to split at. (Before: `settle`, `crowded`.)

## The one page (2026-10-07, `internal/page`, step 5.5)

The single-key page and the multi-key page are **one structure** in two forms; the value overflow is its one-key form with a pointer for the values. Every leaf of the tree has the same head. Offsets are bytes from the start of the object.

- **one page**: the object that holds entries inline: head, key area, key lengths, value lengths, free zeros, values. The package is `internal/page`, with `Str` (values of variable length: strings) and `Fixed` (values of one size: words, pointers) as its two flavors.
- **one-key form** (the *single-key page*): the page of one key. Its key part is the whole remainder of the key; it has **no key-length list**. A page with many keys left with one key turns into it in place.
- **many-key form** (the *multi-key page*): the page of several keys. Its key part is the common prefix of the keys; behind the lists come the remainders.
- **head** (4 bytes): `type | len | n | rawWords`. `len` has nine bits (bit 8 is the lowest bit of the type byte); `n` is the number of slots (1 to 255; 0 in a value overflow); `rawWords` is the size of the byte area in words for a page of pointers, and in a value overflow the offset of its value set in words.
- **key part**: the `len` bytes at byte 4: the remainder in the one-key form, the common prefix in the many-key form. A page begins exactly at its path length: when the path changes, `Skip` and `Prepend` change the key part in place.
- **key lengths**: many-key form only, one byte per slot: the length of the slot's remainder, or `Further` (255) for a further value of the key before it (the *continuation*).
- **value lengths**: `Str` only, one byte per slot, a list of its own in front of the remainders.
- **byte area**: the bytes of the object in front of the values (head, key part, lists, remainders, free zeros); for a page of pointers the first `rawWords` words, whose size is fixed for the life of the object.
- **pointer area**: for a page of pointers, the words behind the byte area: the values, moved as `T` so that the garbage collector keeps its barrier. The values of every flavor end the object (`Fixed`: the array of `T` at `Size-(n-i)*w`).
- **value overflow**: the one-key form of a key whose values do not fit one page (or whose remainder is longer than a page holds): head, key area of 20, 52, 116, 244, 372 or 500 bytes (a longer remainder is held as a string), and the pointer to the value set in the last word. Back into a page when the values fill at most half the room (`BackFits`).

## Parts of the multi-key page

Layout A (`internal/vpage`):

- **head word**: the first 8 bytes of a remainder, big endian, zero padded.
- **slot**: head word and value, side by side, 16 bytes.
- **directory**: one byte per entry (uniform page) or four (general page), at the start.
- **tail**: the remainder bytes beyond the head word, in the **heap** that grows down from the end.
- **uniform page**, **general page**: all remainders of equal length up to 8, or any length.

Layout B (`internal/lpage`):

- **code byte**: the first byte; size class, header size and value width.
- **length header**: the length of the common prefix, one length byte per remainder (0 ends them),
  and one per value if the values differ in length. The offsets of the data follow from it.

## Names that are retired

| old | new |
|---|---|
| vpage, U8 page, leaf-page | page, multi-key page (layout A) |
| flat leaf, typed leaf | single-key page *(planned)* |
| set leaf | single-key page with value overflow, or an oversized object *(planned)* |
| vset, Set3 | value set |
| inner node, N5 to N256 | byte node |
| R8 to R256 | range node |
| term, term leaf | end page |
| path, compressed path, path tail, shared prefix, cp | common prefix |
| base, depth | path length |
| suffix, remainder (both) | remainder |
| unique / multi (as entry kinds) | single-value / multi-value entry |
| settle, crowded | fall back |
| class | size class |

`leaf` no longer names an object. In a drawing it still says "an object without children".

## Where the vocabulary is still heavy

1. **Two page kinds.** A single-key page and a multi-key page are the same idea at the two ends
   (one entry with many values; many entries with one value). The user's decision is to keep them
   apart until step 5. The price is that a multi-value entry cuts its multi-key page in two, and
   everything that follows from that stays: the promote, the fall back, and the byte node as the
   routing the fall back needs. **Measured 2026-10-04** ([step3-tree-pages.md](step3-tree-pages.md)):
   on the natural mix of `street` and `dirs` the fall back takes every multi-key page away, so
   the single-key page carries the real multi data. A multi-key page that holds multi-value entries
   (the experiment `Map.Pairs`, MKMV, parked) saved 36 to 40 % of the memory there, at the price of
   slower mutation; it is step 5.
2. **The common prefix of a page and of a node** are one idea with two places to live. They are
   one word on purpose. The exception to the path length (`pageBase`) is a trick that has made it
   into the definition; it should be measured and probably dropped.
3. **Byte node and range node** are one idea in two granularities (see above).
4. **Variable-length and fixed-size values** are two layouts of one page kind. Whether they share
   a header is for the design note of step 3.1 to show.
