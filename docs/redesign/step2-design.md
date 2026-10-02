# Step 2: pages for every key in the tree (design notes)

How the page of step 1 (`internal/vpage`) sits in the tree (`internal/art`), what changed against
`node-pages`, and the decisions behind it. Results are in `step2-results.md` once measured.

## What a page is in the tree

- **Every key with exactly one value lives in a page**, whatever its length, if the map's values are
  small and pointer-free (`Tree.small`, unchanged). The rule "the first key fixes the page key length"
  (`Tree.pk`, `chooseKeyLen`) is gone. Keys with several values keep their flat, typed or set leaves.
- **A page holds its keys from its base on.** The base is the depth the page was created at, stored in
  the page header (one byte); the bytes before it are the path to the page, checked by the nodes
  above. A lookup passes `key[base:]` to `Page.Find`. This is what a leaf does with its `kl`.
- **Page kinds.** The page of class c (128, 256, 512, 1024 bytes) has the kind `kPage+c`; `isPage` and
  `kLastPage` cover the four. This freed the class byte of the header for the base. `vpage.KindBase` is
  set by `art` at start-up. The tree uses classes 0 to 2; class 3 is for rebasing (below).
- **Limits.** A page is made only where its base is at most 255 (`maxPageDepth`) and what lies below
  the base is at most 255 bytes (`maxPageSuffix`); other keys get a leaf, as before. Keys with 300 bytes
  and paths of 64 KB work as before.
- **Fit depends on the base.** A page's capacity is bytes, not a count, and a page with a prefix holds
  the shared bytes once, so a set of keys may fit a page at base d and not at base d-1 (the prefix is
  chosen by a gain rule, see step1-results.md). `pageFor(items, depth)` therefore takes the depth.

- **Base extension for short keys.** A page's head word costs 8 bytes whatever it holds, so a page
  whose shortest key has less than 8 bytes below the depth starts earlier (`pageBase`): integer keys
  are held whole (base 0), and a lookup takes the head word from the key in one load. Without this,
  `u64` lookups lost 17-21% against `node-pages` to a byte loop and a shift per lookup.

## Operations

| operation | how |
|---|---|
| lookup | `findInPage`: `p.Find(key[p.Base():])`, two rounds of loads |
| insert | `p.LocateIn` (ordered search; also tells whether the key is there), then `p.InsertIn`; a full page splits at a byte boundary of its range node (`splitFull`, `SplitOff`: the first half stays in place), or, if its keys share the byte, `burst` |
| second value | the key leaves its page for a leaf (`place`, `leafWith`: a flat leaf, a set leaf if the key remainder is longer than 254 bytes) |
| delete | `p.Find`, then `p.DeleteAt`; thin pages merge with a neighbour (`rMerge`, `vpage.Merge`), only if both have the same base |
| range scan | `seek` with `Locate` for the bounds; keys are assembled from the path bytes up to the base plus `AppendKey` |
| rebuild | `pageItems(p, pre)` gives items with whole keys (`pre` are the key bytes before the base), `pageFor` turns items back into a page |

## Decisions

1. **Range nodes carry long paths now.** Pages below a range node used to hold whole keys of at most 8
   bytes, so a range node's path was never longer than that and had no tail. Now a range node can have a
   path of any length (200 bytes in the `deep` test set), with a tail like inner nodes. `copyFixed` had to
   learn the four range node classes: it copied them as 256-way nodes, and the first test run on path keys
   crashed in `fork` on garbage in `starts`.
2. **A page moves up with a new base.** When a range node with one child page goes away
   (`collapse`), the page's keys get the node's path bytes in front (`Page.Rebase`, `pageUp`). It may
   not fit (a full page, a long path); then the node stays with its one child. This is the "pinned"
   node: a range node of one page whose base lies below the node's path. It also arises in `build`
   when a set of keys fits a page at the base below the shared bytes and not above. `checkNode` allows
   it; every later delete through the node retries the rebase.
3. **`build` of a single key.** A single item that no page holds (long suffix, deep base) becomes a leaf;
   before, a single item always fit a page by count.
4. **Merging pages** only for equal bases, so that the merged page's suffixes line up.
5. **`crowded`/`settle`** count a page's keys with `Len()`; leaves that stand where no page can (long
   suffixes, deep keys) count as keys with several values, so a range node over many such keys falls
   back to inner nodes. That is the intended safety net and not a bug.

6. **`burst` instead of a rebuild from items.** When a full page cannot be split by the byte at its
   depth, the first version rebuilt the subtree from the keys (`pageItems`, `build`, `ranges`): 40% of
   all allocation in a build of string keys. Now `burst` asks the page for the bytes all its keys share
   (`Shared`): a key that leaves them forks the page off with a range node and a new page or leaf (`fork`
   learned pages); a key that shares them gets the page wrapped in a range node with those bytes as its
   path, in which `splitFull` finds the byte where the keys differ. A key of the page that equals the
   shared bytes exactly becomes the node's term leaf. The items path remains for keys that no page holds.
7. **Churn without allocation.** A page whose keys come and go ran out of heap, not of slots, and
   rebuilt itself in a new object every few operations: now `compactFor` moves the tails together in
   place. Growth into a larger class copies the heads, values and heap in blocks (`copyGeneral`,
   `copyUniform`), and a page of equal short suffixes needs no plan.
8. **`MinGain` 48.** A page prefix of fewer than 48 saved bytes is not worth the shifts and masks it
   costs every lookup of that page: memory +1% at most (`pagefill`), `str` lookups faster.

## Tests

- `keySets` gained `deep`: two keys that fork at byte 200, 500 keys that share 270 more bytes and force
  rebuilds at a depth beyond 255, and one key of 300 bytes.
- `TestPageLongSuffix`, `TestPageMovesUp`, `TestRangeNodeLongPath` cover the limits, the rebase and the
  copy of long-path range nodes of every class; `TestPageLifecycle`, `TestPageKeyLength` and
  `checkPage` were rewritten for pages of any key length.
- `internal/vpage/api_test.go` covers `Build`, `Locate`, `Find`, `Rebase`, `ByteAt`, `DeleteAt`,
  `SplitAt` and the accessors.
