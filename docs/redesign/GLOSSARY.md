# Glossary

The vocabulary of the redesign, agreed with the user on 2026-10-03. Use these words in code,
comments and documents from now on. Each entry says what the word names, and what it was called
before. A word marked *planned* names something that the target design needs and the code does not
have yet; the code still has the old object under its old name.

A definition that needs more than two sentences is a warning sign: the concept may be too
complicated. Such places are collected at the end ("Where the vocabulary is still heavy").

## Keys, values, entries

- **key**: a byte string. The tree orders its entries by key, bytewise.
- **value**: what a key maps to. *Planned:* a byte string of any length (the layouts under design,
  [whataleafneedstostore.md](whataleafneedstostore.md), store lengths); for fixed-size values a
  specialized variant follows later. *Today:* a fixed-size `T`, held as a word of 8 bytes in a page.
- **entry**: a key with all its values. The unit that objects store.
- **single-value entry**, **multi-value entry**: an entry with exactly one value, or with two or
  more. (Before: single-value or multi-value *key*; in the bench: `unique` and `multi`.)

## Where a key is cut: path, remainder, common prefix

Every object that stores entries stores only the end of their keys. The tree spells the beginning.

- **path** (of an object): the bytes of the key that the descent has matched when it reaches the
  object: the branch bytes and the common prefixes of the nodes above. The object does not store
  them. (Before: `depth`, `base`, "the path to it".)
- **remainder** (of an entry in an object): the rest of the key, which the object stores. Key =
  path + remainder. (Before: `suffix` in pages, `remainder` in leaves.)
- **path length**: the number of bytes of the path; the remainder starts there. (Before: `base`,
  `depth`.) *Exception:* a page may start its remainder a few bytes earlier than the path ends,
  so that every key has a whole head word in the page (`pageBase`); those bytes are stored twice.
- **common prefix**: the bytes that all entries below a node, or in a page, have in common
  after the path to it, stored once there and compared in one step. This is what the ART
  literature calls *path compression*, with the bytes stored as the node's *prefix* (my reading of
  the paper, not checked word by word). (Before: `path`, `compressed path`, the header's `prefix`,
  the page's `prefix`; for the part beyond 12 bytes in a node: `path tail`.)
- **prefix scan**: the user's operation that visits all entries whose key starts with given bytes.
  Always these two words. "Prefix" alone means the common prefix.

## Objects

- **page**: an object of 128, 256 or 512 bytes that stores entries, with no pointer in it unless
  the values need one (Go's garbage collector forces a separate allocation type for those).
  There are two layouts, and no page mixes them: no page holds entries with different numbers of
  values.
  - **multi-key page**: many entries with one value each, in key order. This is the page of
    `internal/vpage`, and the object steps 1 and 2 built. (Before: `vpage`, `page`, leaf-page,
    U8 page. In the sketches: MKSV.)
  - **single-key page** *(planned)*: exactly one entry with all its values inline, whatever fits:
    any number of values of any length. (Before: *flat leaf*, *typed leaf*, and the inline part of
    the *set leaf*. In the sketches: SKMV; SKSV is the case of one value.) An entry with a second
    value leaves its multi-key page for a single-key page (*promote*).
  - **end page**: the page that a node holds for the entry whose key ends exactly at the node,
    which happens when that key is a prefix of the other keys below (`ab` beside `abc`). It is a
    single-key page. (Before: `term`, `term leaf`.)
- **oversized object** *(planned)*: a single-key page whose remainder does not fit the largest
  class. It stores the remainder inline anyway, in an object of the size Go gives it (576 bytes
  and up). It is exempt from R1 and counted separately in the object statistic. It saves the
  random cache miss that a pointer to the key would cost. (Before: *key overflow*, a set leaf with the
  key as a string.)
- **value overflow** *(planned)*: when the values of a single-key page do not fit it, the page
  holds one pointer, at offset 8, to a **value set** and no values. (Before: the *set leaf*.)
- **value set**: a container for many values of one key: an array (up to 64 values), then a
  hash set. (Before: `vset.Set`, with its inline stage of 3 values, and `Set3` for the hash
  stage.) The inline stage becomes redundant once values sit in the single-key page.
- **node**: an object that routes: it picks the child for one byte of the key. It stores no
  entries; the entry whose key ends at the node is in its end page. There are two kinds, which
  differ only in how fine they cut the byte:
  - **byte node**: one child per byte value. The child never sees that byte again; the node
    consumes it. (Before: `inner node`, N5 to N256.)
  - **range node**: one child per range of byte values. The byte stays in the keys below, because a
    child holds several values of it. (R8 to R256.)

  A byte node is the special case of a range node with one value per range, plus the consumed byte.
  After step 3 the byte node may not be needed at all (see the plan, step 4).

## Memory

- **size class**: one of the sizes our objects have: pages 128, 256 or 512 bytes, nodes 128 to
  512 (R256: 2112), and so on. (Before: `class`, used in six tables.)
- **block**: the size the Go allocator really gives an object (`art.Block`), which is the next
  Go size class.
- **round**: a set of cache-line loads whose addresses are all known before the first completes
  (R5, [STRATEGY.md](STRATEGY.md)).
- **R1 to R6**: the binding rules of [STRATEGY.md](STRATEGY.md), section 2.

## Operations

- **grow**, **shrink**: a page moves to the next larger or smaller size class (a new object).
- **split**, **merge**: a full multi-key page becomes two, at a byte boundary; two thin
  neighbours become one.
- **burst**: a full page that cannot be split at a byte boundary gets a range node of its own.
- **promote**: an entry that gets a second value leaves its multi-key page for a single-key page.
- **fall back**: a subtree is rebuilt from byte nodes and single-key pages, because entries with
  several values crowd it. (Before: `settle`, `crowded`. Only exists as long as a single-key page
  cuts a multi-key page in two, see below.)

## The implemented multi-key page

These words name parts of the layout of `internal/vpage`; the sketches in
[whataleafneedstostore.md](whataleafneedstostore.md) are a second candidate with a header of
lengths and no directory, which a measurement has to compare.

- **head word**: the first 8 bytes of a remainder, big endian, zero padded.
- **slot**: head word and value, side by side, 16 bytes.
- **directory**: one byte per entry (uniform page) or four (general page), at the start.
- **tail**: the remainder bytes beyond the head word, in the **heap** that grows down from the end.
- **uniform page**, **general page**: all remainders of equal length up to 8, or any length.

## Names that are retired

| old | new |
|---|---|
| vpage, U8 page, leaf-page | page, multi-key page |
| flat leaf, typed leaf | single-key page *(planned)* |
| set leaf | single-key page with value overflow, or an oversized object *(planned)* |
| vset, Set3 | value set |
| inner node, N5 to N256 | byte node |
| R8 to R256 | range node |
| term, term leaf | end page |
| path, compressed path, path tail, shared prefix | common prefix |
| base, depth | path length |
| suffix, remainder (both) | remainder |
| unique / multi (as entry kinds) | single-value / multi-value entry |
| settle, crowded | fall back |
| class | size class |

`leaf` no longer names an object. In a drawing it still says "an object without children".

## Where the vocabulary is still heavy

1. **Two page layouts.** A single-key page and a multi-key page are the same idea at the two ends
   (one entry with many values; many entries with one value). The user's decision is to keep them
   apart for now. The price is that a multi-value entry cuts its multi-key page in two, and
   everything that follows from that stays: the promote, the fall back, and the byte node as the
   routing the fall back needs. If a multi-key page could hold entries with different numbers of
   values, all of these words would go.
2. **The common prefix of a page and of a node** are one idea with two places to live. They are
   one word on purpose. The exception to the path length (`pageBase`) is a trick that has made it
   into the definition; it should be measured and probably dropped.
3. **Byte node and range node** are one idea in two granularities (see above).
4. **Fixed-size and variable-length values.** The layouts under design take values of any length;
   the code has fixed-size ones. The vocabulary says "value" for both. A specialized layout for
   fixed-size values, which can swap instead of shifting when it compacts, will need its own
   name when it exists.
