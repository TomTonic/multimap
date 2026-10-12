# D1: a node step without the switch on the node type (design note, 2026-10-11)

Words as in [GLOSSARY.md](GLOSSARY.md). Candidate D1 of [candidates-2026-10.md](candidates-2026-10.md); written while W1 + W2 are
measured (jobs w12 on the M1 and the same arguments on the PC). Nothing is built before the counts of section 4 and the user's go.

## 1. What the step does today

`Tree.find` (`internal/art/lookup.go`) runs one loop iteration a level: is the object a page; does the node's prefix match
(`swar.Match8`, `longMatch` beyond 12 bytes); does the key end here (end page); then `switch n.objType` with five cases:

| type | size | search | child slot |
|---|--:|---|---|
| N5 | 64 B | `Index8` over the 8 key bytes at 0x10 | `child[i]` at 0x18 + 8i |
| N12 | 128 B | `Index8` over the 16 key bytes at 0x10 (two words) | `child[i]` at 0x20 + 8i |
| N26 | 256 B | `Has` and `Rank` over the bitmap at 0x10..0x2f (a loop over up to three words) | `child[r]` at 0x30 + 8r |
| N58 | 512 B | the same as N26 | the same as N26 |
| N256 | 2,080 B | none | `child[b]` at 0x18 + 8b |

What was measured (descent-analysis.md): on the M1 the lines of the switch, its targets and the key byte take 24 to 26 % of `find`
(by source line, so the wait for a load may sit in it too); on the PC the switch, the loop and the SWAR work are 21 to 36 % of `find`
together (by instruction). The type of the next node is not known before its first line arrives, and along a path of path-like keys
the types change from level to level (url single-value 4K: 2.6 N5, 0.7 N12, 1.1 N26, 1.5 N58 a lookup), so the branch predictor
sees a sequence it can learn only per path.

## 2. Proposal, in three steps that can be measured one by one

**D1a: one case for N26 and N58.** Their layouts are the same up to the length of the child array (bitmap at 0x10, children at
0x30); one case serves both. Nothing else changes; the switch has four targets instead of five.

**Correction (2026-10-12): D1b as written below does not fit.** With the children from 0x18, twelve child slots take 0x18 to 0x77
and the last slot 0x78 to 0x7f: there is no room for key bytes 8 to 11 in 128 bytes. A shared shape of N5 and N12 holds 11 children
(16 + 8 key bytes + 11 slots + 8 key bytes = 128): it costs the middle node a slot. The key arrays are not waste of a choice but the
alignment of the child pointers (N5: 16 + 5 = 21 rounded to 24; N12: 16 + 12 = 28 rounded to 32), byte-exact since the first ART
(`bb4bc0e`: N4/N11 with a 24-byte header; `a47414f`: the 16-byte header, one slot more each). Which prefix length and which node
classes would serve best is analysed in node-classes-analysis.md.

**D1b: N5 and N12 of the same shape.** (First written as "no N5"; withdrawn after the user's objection of 2026-10-11: the N5 is kept
on purpose for the degenerate trees of path-like keys, where nodes are sparse and the routing layer is already too thin; making each
of them 128 bytes would go against that. PLAN.md's point "remove the smallest byte node" is to be read with that exception.) Today
the two differ only in where the children start (N5 0x18, N12 0x20), because the N12 has 16 key bytes. If the N12 keeps its first 8
key bytes at 0x10, its children at 0x18 like the N5, and its key bytes 8 to 11 behind its children (0x78..0x7b, inside its 128
bytes), then both run the same code: `Index8` over the word at 0x10, and only if the byte is not among the first eight and the node
has more than eight children the word at 0x78; the child at 0x18 + 8i in both. No memory changes. The switch then has three
targets (sorted node, bitmap node, N256), and N256 is rare.

**D1c: no branch between the sorted and the bitmap nodes.** Both searches read only the node's first 48 bytes (key word or bitmap),
so both can be computed and one chosen by arithmetic: the index `i` of the sorted search (`i < count`) and the rank `r` of the bitmap
(`Has`, popcount of the words below), the child slot at `0x18 + 8i` or `0x30 + 8r`, picked by a mask from the type byte. The miss is
the same choice between `i >= count` and `!Has`. N256 keeps its own branch.

D1c needs D1b (one layout for the sorted nodes). D1a is part of D1c, but on its own it shows whether the number of targets matters.

**Why the type matters at all (on the user's question whether speculation should not already do the right thing).** The processor
does not wait for the type byte: it predicts the target of the switch and runs it ahead, loading the key word and the child slot of
the guessed type while the node's line is still on its way. When the guess is right, the type costs nothing; when it is wrong (the
types change from level to level, and the guess was the type of another path), the work is thrown away when the line arrives. Two
things speculation cannot do: start the next node before this node's line has arrived (the child's address comes from that line:
this is the wait for the next object, the largest part of `find`), and run one code for all types, because today the types do not
share one: the child slot is at 0x18, 0x20, 0x30 or 0x18 + 8b, and the search is a byte search or a bitmap. D1b and D1c make the
code the same for all but N256, so that there is nothing left to guess.

## 3. Costs and conflicts

- **Memory (D1b):** none; the N12 only moves its last four key bytes behind its children.
- **Lines (D1b):** an N12 whose searched byte is not among its first eight key bytes reads the word at 0x78, in its second line on
  the PC (where its children from the sixth on are already); on the M1 (128-byte lines) it is the same line. How often a lookup
  needs the second key word is counted first (section 4).
- **Work (D1c):** both searches every level, about ten more simple instructions, against one hard-to-predict branch less. On the M1
  a mispredicted branch costs about 13 cycles, on the PC about 15 to 20.
- **Writes:** inserts and removals descend the same way (`findLoc`, `findSlot`): the same change applies there. Growing and shrinking
  nodes lose the N5 cases (simpler code).
- **Ranges:** the walk (`nextChild`, `enter`) knows the node types too; D1b removes a case there, D1c does not touch it.
- **Hot path style:** D1c is the kind of code the user asked for (arithmetic instead of dispatch); it must stay inlined, which
  `-gcflags=-m` checks.

## 4. Counts before any code (to run when the PC is free again)

1. In the model (`rangemodel`, today's tree): the lookups' node types level by level, and how often the class changes from one level
   to the next (five types today, three after D1b); and how often a lookup through an N12 finds its byte only among key bytes 8 to
   11 (the second key word of D1b).
2. With `-gcflags=-m`: that `find` stays as it is inlined today with D1c's arithmetic (no call left in the loop).

## 5. Prediction (before the counts; the counts may change it, and then this section says so)

- D1a alone: within noise (±2 %) everywhere: the switch is compiled to compares and jumps either way.
- D1b + D1c against today: lookups of single-value maps at 4,096 and 16,384 keys +4 to +10 % on the M1, +2 to +6 % on the PC, most
  for url, path, dirs (most levels); writes +2 to +5 % (their descent); ranges within noise; memory unchanged.
- A result below this on the M1 means the M1's switch lines were mostly the wait for the next object after all (section 1), not
  the branch: then D1 stops, and the descent is a question of structure only (D3).

## 4a. The counts (2026-10-11, 17:55; `bench/results-layout/descent-2026-10/d1-counts.md`)

The model, today's tree, single-value, the seven structured kinds, 4K/16K/64K:

| | type changes a lookup, 5 types → 3 (after D1b) | share of level pairs that change, 5 → 3 | N12 visits that need the second key word, a lookup (share of N12 visits) |
|---|---|---|---|
| str | 1.4 to 2.7 → 1.1 to 2.0 | 99 to 100 → 73 to 84 % | 0.26 to 0.35 (18 to 25 %) |
| email | 0 to 0.06 (all N26) | 0 to 3 % | 0 to 0.02 |
| url | 4.0 to 4.9 → 2.6 to 3.0 | 61 to 83 → 41 to 54 % | 0.05 to 0.12 (5 to 8 %) |
| path | 6.1 to 7.1 → 3.1 to 3.8 | 65 to 85 → 35 to 42 % | 0.84 to 0.87 (24 to 39 %) |
| street | 1.7 to 2.8 → 1.2 to 1.7 | 73 to 84 → 44 to 61 % | 0.10 to 0.16 (9 to 11 %) |
| dirs | 6.0 to 7.9 → 3.1 to 3.8 | 76 to 91 → 37 to 46 % | 0.05 to 0.81 (2 to 39 %) |
| links | 1.3 to 2.6 → 0.6 to 1.9 | 56 to 79 → 23 to 56 % | 0.05 to 0.08 (9 to 14 %) |

The node type changes on most levels of a lookup (but for email, whose trees are all N26): a predictor that guesses "the same type
as last time at this place" is wrong on 56 to 100 % of the level pairs today and still on 23 to 84 % with three types (D1b). Only
D1c (no branch between the sorted and the bitmap nodes) takes the guess away. The second key word of D1b is needed on 0 to 0.9 N12
visits a lookup (path and dirs at 64K most), an extra line on the PC for those.
