# What an object stores: the system as it is today

(Written before the vocabulary of [GLOSSARY.md](GLOSSARY.md); its words are the old ones.)

This is the case analysis of [whataleafneedstostore.md](whataleafneedstostore.md), filled in with what
the code on `cacheline` (commit `89708db`, end of step 2) really does. Nothing here is a proposal.
Where the code differs from the outline's wording, the text says so.

## Notation

As in the outline: B is the capacity of an object, K the length of a key, W the length of a value,
all in bytes. Two refinements are needed because the code works that way:

- **S, not K, is what an object stores.** A page or leaf holds a key only from its *base* on; the
  bytes before are the path to it, kept once in the nodes above. So S = K - base, with base at most
  the depth at which the object hangs. A page with a base smaller than that depth (`pageBase`)
  holds up to 8 bytes more of each key than the path, which makes the lookup cheaper; S counts those.
- **The kind of T decides which objects exist at all.** The map looks at its value type T once
  (`Map.decide`):

| T | what the map uses | pages? |
|---|---|---|
| pointer-free, at most 8 bytes (`uint64`, `int32`, ...) | pages, flat leaves, set leaves | yes |
| pointer-free, 9 to 16 bytes | flat leaves, set leaves | no |
| holds a pointer, at most 64 bytes, 8-aligned (`string`, `*X`, ...) | typed leaves, set leaves | no |
| anything else | set leaves | no |

Only the first row has the "many keys in one object" case (Case 3). In the other rows every key has
a leaf of its own, and the cases below that talk about pages do not apply.

The three kinds of object that hold keys:

- **Page** (`internal/vpage`): many keys, one value each, no pointer inside. 128, 256 or 512 bytes
  (a class of 1024 exists in the code but `MaxClass` stops at 512).
- **Leaf**: one key with its values. A *flat leaf* (pointer-free T, at most 16 bytes) and a *typed
  leaf* (T with pointers) hold the values inline, unsorted. A *set leaf* holds them in a
  `vset.Set`.
- **Node** (N5 to N256, R8 to R256): holds no keys of its own except the *term* leaf of the key that
  ends at the node. It routes: it holds 8-byte pointers to pages, leaves and nodes.

## What a page costs

```
header 8 B | prefix | directory | slots, 16-aligned | free | heap of tails (grows down)
```

- A key's **slot** is 16 bytes: its first 8 suffix bytes as a head word (zero padded) and its value
  as one 8-byte word. The value is always inline and always 8 bytes, whatever T's size.
- A **uniform** page (all suffixes of equal length 1 to 8): directory of 1 byte per key. A key
  costs 17 bytes. Example: u64 keys fill at most 29 per 512-byte page, by the formula; the real fill
  after random inserts is lower (split at 85 %, shrink below 70 %).
- A **general** page (any suffix up to 255 bytes): directory of 4 bytes per key (tag, suffix length,
  tail offset). A key costs 20 + max(0, S - 8) bytes: the part of the suffix beyond the head word is
  its *tail* in the heap.
- A **prefix** that all keys of the page share is stored once if that saves at least 48 bytes
  (`MinGain`).
- Overhead: the header, the prefix rounded up to 8, padding of the slots to a multiple of 16, and
  the free space that a growth step leaves.

Growth, shrink, split:

- A full page **grows** to the next class by building a new object (128, 256, 512). A page that
  empties to at most 70 % of the next smaller class moves back down.
- A full page of 512 bytes **splits** at a byte boundary near its middle into two pages below the
  same range node. If all its keys share the byte (or it has no range node above), it first gets a
  range node of its own (`burst`).
- Two neighbouring pages **merge** if together they fill at most 75 % of 512 bytes.

## Case 1: one key, one value

The case occurs for every key that does not share its object with other keys, and as the first key
of a map. What happens depends on S and T.

### Case 1.a: pointer-free T of at most 8 bytes, S <= 255

The key and its value are in a page (`newPageFor`): a new page of 128 bytes if the key has no page to
join. Key suffix and value are stored whole and inline. There is no reference to either.

### Case 1.b: the entry does not fit

For pages this cannot happen through size alone: the longest entry is 20 + 247 bytes, and a page of
512 bytes holds it. What can happen:

- **S > 255** (or the page would sit deeper than depth 255): the key does not go into a page. It
  gets a **leaf**, as for the other kinds of T:
  - S <= 254 and the value fits: a flat leaf with the suffix inline and the value inline;
  - S > 254, or a whole key of more than 65535 bytes: a set leaf that holds **a reference to the
    key** (a separate string with the *whole* key, not only the suffix), and the value inline in
    the `vset.Set` (inline up to 3 values).
- **The page is full** and holds other keys: grow, else split (see above). Neither makes a
  reference.

### Case 1 for the other kinds of T

Every key has a leaf, so one key with one value is always a leaf:

- **flat T**: a flat leaf of 6-byte head, key suffix, padding to T's alignment, the value. The
  smallest class that holds them: 32, 48, 64, 96, 128, 192, 256, 384 or 512 bytes. If key suffix and
  one value do not fit 512 bytes (S > 254): a set leaf with the key as a string.
- **typed T**: a typed leaf with room for 1 value (classes with 1, 2, 3, 4, 6, 8, 12 or 16 values;
  key area 2 to 58 bytes). If S > 58: a set leaf, and its key is inline (up to 254 bytes) or a
  string.
- **other T**: a set leaf. Its key is inline in an array of 16 to 256 bytes or, beyond 254 bytes, a
  string; its value is inline in the `vset.Set`.

In none of these is the value a reference. The key is a reference only in a set leaf, only for long
keys.

## Case 2: one key, several values

The key has a **leaf**. A page never holds a key with several values (its slot has room for one).

How it gets there: a key that has one value in a page and receives a second, different value
(`Map.Add`) leaves the page. A leaf with the key and both values is built; the page is cut around
the key, or the leaf replaces the page if the page held only this key (`place`). A key with a
leaf stays with a leaf: removing values down to one does not give the key a slot in a page again.
Removing the last value removes the leaf.

### Case 2.a: the values fit the leaf's largest class

- **flat T**: a flat leaf holds the suffix S and all n values inline: 6 + S + padding + n * sizeof(T)
  <= 512. Values are unsorted; a lookup scans them. When full the leaf moves to the smallest class
  that holds twice its values (at least 64 bytes); it shrinks back when the smaller class would still
  be half empty, never below 64 bytes.
- **typed T**: a typed leaf holds up to 16 values inline, in classes of 1, 2, 3, 4, 6, 8, 12 and 16
  values; the smallest it grows into or shrinks back to is 4 values.

### Case 2.b: they do not fit

The leaf becomes a **set leaf** (`spill`): the leaf holds the key (inline up to 254 bytes, else a
string) and a `vset.Set` for the values:

- up to 3 values inline in the leaf (the `Set` is in the leaf object);
- up to 64 values: a **reference** to one unsorted array, a separate object (no pointers for
  pointer-free T);
- more than 64: a **reference** to a hash set (`Set3`).

A set leaf becomes a flat or typed leaf again once its values fill half of the largest class
(`unspill`); a hash set becomes an array again at 32 values.

So here the values are a reference, and only here (and the key is a reference only if it is long).
The wording of the outline's Case 2.b.i, "a reference to a Set3, array or vset", is exactly this:
`vset` is the container, and Set3 and the array are its two out-of-line forms.

## Case 3: several keys, one value each

This is the case pages are for, and only for pointer-free T of at most 8 bytes. All keys sit in
**pages**: sorted by suffix, inline, with their values inline, in a page of 128, 256 or 512 bytes.

- Capacity: the keys fit if header + prefix + the directory bytes + 16 per key + the tails <= B,
  with B = 128, 256 or 512; the page takes the smallest class that holds them with room (see "What
  a page costs").
- Keys whose suffix exceeds 255 bytes are the exception: those keys have leaves, next to the pages
  (Case 4 below).
- Pages hang below **range nodes**. A range node gives each child a range of key bytes: the page
  of the range holds all keys whose byte falls in it. A lookup goes node by node to the page, takes
  the directory and the slot there, and compares the tail.
- If T is not a page type: no pages; this case is a tree of leaves, one per key (Case 1, other
  kinds of T).

## Case 4: several keys, at least one with several values

Pages and leaves live side by side below the same range nodes. Concretely:

- The keys with one value are in pages. Each key with several values (or a key too long for a page)
  is a leaf, a child of the range node next to the pages, or the term leaf of a node. The page
  that held the key is cut around it (left page, leaf, right page).
- A key that gets a leaf can crowd the subtree: **fallback** (`settle`, `crowded`). If in the
  subtree of a range node the leaves and inner-node subtrees outweigh its pages by a ratio of 4
  (more than about 1 key in 5 holds several values), the range node and everything below it are
  rebuilt as inner nodes and leaves, as a tree without pages. Below an inner node, a tree never has
  pages again; new keys there get leaves. Only a subtree that falls back to a single leaf takes new
  keys in pages again.
- So a map whose keys mostly have several values ends up as a tree of leaves, and a map whose keys
  mostly have one value ends up as a tree of pages; maps in between are mixed, and what is mixed
  is decided per subtree.
- A leaf with a long key keeps a reference to the key (a string); a leaf with more than a few
  values keeps the values inline or by reference as in Case 2.

## What each object holds, in one table

| object | keys | value(s) | references from the object |
|---|---|---|---|
| page | many; suffix inline, up to 255 bytes, tails in the heap | one per key, inline, 8-byte word | none (pointer-free) |
| flat leaf | one; suffix inline up to 254 bytes | n inline, unsorted, up to 512 bytes | none (pointer-free) |
| typed leaf | one; suffix inline up to 58 bytes | up to 16 inline | the pointers inside the values |
| set leaf | one; suffix inline up to 254 bytes, else the whole key as a string | `vset.Set`: 3 inline, then an array, then a hash set | the key string if long; the array or hash set |
| node | none (the term leaf is a child) | none | the children |

## Differences from the outline

- **No reference to a value from a page.** A page's slot is 16 bytes; a value that is not a
  pointer-free word of at most 8 bytes makes the whole map leaf-based, and a second value makes the
  key a leaf. The outline's "the page stores a reference to the value" (Cases 1.b.i, 1.b.iii) does
  not exist for pages; it exists only in a set leaf (Case 2.b).
- **No reference to a key from a page.** A suffix above 255 bytes is not stored in a page; the key
  gets a leaf. Only a set leaf holds a reference to a key.
- **Growing is by class, then splitting.** The outline's "if possible and if it suffices: enlarge
  the page" is the growth by classes 128, 256, 512. Beyond 512 there is no growth: the page splits
  into two, which the outline has no step for.
- **Leaves are not capped at multiples of 128 bytes.** Flat leaves have 9 size classes from 32 to
  512 bytes, set leaves and typed leaves have their own sizes. The cache-line premise of
  [STRATEGY.md](STRATEGY.md) is met only by pages and by the nodes of 128 bytes and more. Step 3
  of the [plan](PLAN.md) is about this.
- **Decided per map, by T; per key, by the number of values; per subtree, by the fallback.** The
  outline's cases are decided per key and value size only; the code also decides per value type.
