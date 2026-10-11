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

**D1b: no N5** (already a point of step 6 in PLAN.md, from the cache-line rules: the smallest node becomes N12 with 128 bytes). A
node of up to 12 children is then always an N12; the switch has three targets (N12, bitmap node, N256), and N256 is rare (near the
root of large maps), so the branch that matters is one: "up to 12 children or more".

**D1c: no branch between N12 and the bitmap nodes.** Both searches read only the node's first 48 bytes (keys or bitmap), so both can
be computed and one chosen by arithmetic: the index `i` of the sorted search (`Index8`, `i < count`) and the rank `r` of the bitmap
(`Has`, popcount of the words below), the child slot at `0x20 + 8i` or `0x30 + 8r`, picked by a mask from the type byte. The miss
(no child) is the same choice between `i >= count` and `!Has`. N256 keeps its own branch. The loop body then has the branches of the
page check, the prefix and the end of the key, all of which go the same way on most levels.

D1c only makes sense after D1b (with N5 the sorted search has two layouts). D1a is part of D1c, but on its own it shows whether the
number of targets matters at all.

## 3. Costs and conflicts

- **Memory (D1b):** every N5 becomes an N12, 64 bytes more. In the census of url single-value 4K (wprof run of 2026-10-10) there
  are 223 N5 for 4,096 keys: +3.5 bytes a key; to be counted for all cases (section 4).
- **Lines (D1b):** an N12 that holds five children has its fifth child in its second line on the PC (lever D2); a lookup through a
  former N5 may wait for one line more on the PC. On the M1 (128-byte lines) an N12 is one line.
- **Work (D1c):** both searches every level, about ten more simple instructions, against one hard-to-predict branch less. On the M1
  a mispredicted branch costs about 13 cycles, on the PC about 15 to 20.
- **Writes:** inserts and removals descend the same way (`findLoc`, `findSlot`): the same change applies there. Growing and shrinking
  nodes lose the N5 cases (simpler code).
- **Ranges:** the walk (`nextChild`, `enter`) knows the node types too; D1b removes a case there, D1c does not touch it.
- **Hot path style:** D1c is the kind of code the user asked for (arithmetic instead of dispatch); it must stay inlined, which
  `-gcflags=-m` checks.

## 4. Counts before any code (to run when the PC is free again)

1. In the model (`rangemodel`, today's tree): the lookups' node types level by level, and how often the class changes from one level
   to the next in the two-way view of D1b ("up to 12 children" or "more"); and the N5 a key over the seven structured kinds and three
   sizes (the memory of D1b).
2. With `-gcflags=-m`: that `find` stays as it is inlined today with D1c's arithmetic (no call left in the loop).

## 5. Prediction (before the counts; the counts may change it, and then this section says so)

- D1a alone: within noise (±2 %) everywhere: the switch is compiled to compares and jumps either way.
- D1b + D1c against today: lookups of single-value maps at 4,096 and 16,384 keys +4 to +10 % on the M1, +2 to +6 % on the PC, most
  for url, path, dirs (most levels); writes +2 to +5 % (their descent); ranges within noise; memory +1 to +5 bytes a key.
- A result below this on the M1 means the M1's switch lines were mostly the wait for the next object after all (section 1), not
  the branch: then D1 stops, and the descent is a question of structure only (D3).
