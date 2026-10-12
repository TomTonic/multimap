# Which inline prefix and which node classes? (statistical analysis, 2026-10-12)

Words as in [GLOSSARY.md](GLOSSARY.md). Asked by the user after the layouts of N5 and N12 (d1-node-step-design.md): which length
of the prefix in the node header (8, 12, 16 ... bytes) and which branching steps (today 5, 12, 26, 58, 256) would be best, and where
the sweet spot lies.

## Method

The node classes do not change the shape of the tree (which nodes exist, their children, their prefixes): that is set by the pages
and the bursts. They change the bytes of every node, whether its prefix needs a tail, and which line of the node a lookup reads its
child from. So the analysis takes today's tree once (the model `rangemodel`, exact against the measured shape), records for every
byte node its slots (children and end page), its prefix length and the keys below each slot (each key looked up once), and then
evaluates any design without rebuilding a tree:

- **a design** is a prefix length P inline in the header (header = 4 + P bytes: type, count, prefix length, prefix) and a set of node
  sizes; each size holds as many slots as fit: a sorted node (key bytes, then 8-byte aligned child pointers) up to 24 slots, a bitmap
  node (32-byte bitmap) above; a 256-way node (direct) for more than the largest class;
- **memory**: every node in the smallest class that holds its slots, plus a tail of the prefix beyond P in the same object (16-byte
  steps), rounded to Go's allocation classes; bytes a key;
- **lines a lookup** (64- and 128-byte lines): every node passed costs its first line, the line of the child slot taken if it is
  another one, and the line of the tail if its prefix is longer than P;
- **node types on the way**: how many classes a design has (each is a case of the descent, D1).

Cases: the seven structured kinds, single-value, 4,096, 16,384 and 65,536 keys (natural gives pages with fewer keys and more nodes:
to be added if the result differs by kind). Designs: P from 4 to 28 in steps of 4, node sizes chosen from 64, 96, 128, 192, 256, 384,
512, 768, 1,024 bytes (two to five classes) plus the 256-way node.

## Prediction (before any number)

1. Today's P = 12 is within 4 bytes of the best: prefixes beyond 12 bytes are rare (0 to 11 % of the N5 and N12, nodefill.md), and
   every byte of P costs every node.
2. The class sizes matter more than P: a class between 128 and 256 bytes (about 20 slots) and a small node of 6 to 8 slots save 5
   to 15 % of the routing bytes, with no more than 0.2 lines a lookup more.
3. No design is better than today in both memory and lines on the PC by more than 10 % in either.

## Run

`rangemodel -nodedump` (today's tree, 92,617 byte nodes in 42 cases: seven kinds, single-value and natural, 4K/16K/64K;
`bench/results-layout/node-classes/nodes.txt.gz`), evaluated by `classes.py` (every P from 4 to 28 and every set of two to five
node sizes of the menu, 2,667 designs) and `few.py` (chosen designs). 2026-10-12. Of the nodes, 62 % have no prefix at all, 88 %
one of at most 8 bytes.

## Result

**How much the routing nodes weigh at all** (today, single-value): 1.4 to 10.3 bytes a key (str, email), 3.4 to 4.6 (street,
links), 5.4 to 8.4 (dirs, path, url): against 27 to 72 bytes a key in all (c1, M1 memory of the ordered map, street to url), the
nodes are 11 to 16 % of the memory. Lines a lookup: 3.9 to 17.6 on 64-byte lines, 3.1 to 13.9 on 128-byte lines; every node passed
costs its first line, and that part no node class changes.

**Chosen designs against today** (single-value, mean of the ratios over the 21 cases; `few.md`):

| P | sizes (and the 256-way node) | slots | routing B a key | memory | lines 64 B | lines 128 B |
|--:|---|---|--:|--:|--:|--:|
| 12 | **64/128/256/512 (today)** | 5/12/26/58/256 | 5.2 | 1.00 | 1.000 | 1.000 |
| 12 | 64/128/192/256/512 | 5/12/19/26/58/256 | 4.8 | 0.94 | 0.997 | 0.995 |
| 8 | 64/128/192/256/512 | 5/12/20/26/58/256 | 4.8 | 0.94 | 0.980 | 0.994 |
| 4 | 64/128/192/256/512 | 6/13/20/27/59/256 | 4.6 | **0.91** | 0.990 | 1.003 |
| 8 | 64/128/256/512 | 5/12/26/58/256 | 5.2 | 1.00 | 0.987 | 1.004 |
| 20 | 64/128/256/512 | 4/11/25/57/256 | 5.8 | 1.18 | 1.051 | 1.021 |
| 28 | 64/128/256/512 | 3/10/24/56/256 | 6.2 | 1.26 | 1.088 | 1.042 |
| 12 | 64/128/256 | 5/12/26/256 | 5.9 | 1.13 | 0.997 | 1.016 |
| 12 | 64/128/512 | 5/12/58/256 | 7.1 | 1.41 | 1.000 | 1.000 |
| 12 | 128/512 | 12/58/256 | 8.5 | 1.64 | 1.002 | 1.002 |
| 12 | 64/448 ("5, 50, 256") | 5/50/256 | 10.7 | 2.12 | 1.037 | 1.007 |
| 12 | 128 | 12/256 | 22.7 | 4.66 | 1.009 | 1.111 |

Natural alike (best 0.92 to 0.94 of today's routing memory, lines 0.98 to 1.00; `result.txt`). The search over all 2,667 designs
finds nothing better than the best rows above (Pareto front in `result.txt`).

**The sweet spot is flat, and today's design is in it.**
- **Prefix length:** 4 to 12 bytes are equal within 4 % of the routing memory and 2 % of the lines; beyond 12 every class loses
  slots and both get worse (16: the 58-slot class no longer fits 512 bytes; 20: +18 % memory, +5 % lines; 28: +26 %, +9 %). Short
  prefixes cost few tails, because long prefixes are rare; today's 12 is at the upper edge of the flat part, 8 is as good.
- **Branching steps:** today's 5/12/26/58/256 is within 6 to 9 % of the best routing memory (a class of 19 to 20 slots at 192 bytes
  added; 6/13/20/27/59 with a 4-byte prefix). Fewer classes cost a lot: three classes (5/50/256 or 12/58/256) take 1.6 to 2.1 times
  the routing bytes, a single class 4.7 to 9.9 times.
- **In the whole map:** the best design saves about 0.5 bytes a key (about 1 % of the memory) and at most 2 % of the lines on the
  PC, nothing on the M1. The node classes are not a lever; the descent's cost is the number of nodes passed (each a dependent line),
  which only the structure (pages, bursts, D3) changes.

**Against the prediction:** 1. P = 12 within 4 bytes of the best: hit (4 and 8 are a little better, by 0 to 4 %). 2. Class sizes
matter more than P and save 5 to 15 % of the routing bytes with at most 0.2 lines more: about hit (6 to 9 %, lines even 0 to 2 %
fewer). 3. No design better than today by more than 10 % in memory or lines: hit (9 % and 2 % at most).

**For D1:** fewer node types for the sake of the dispatch cost routing memory (three types: 1.6 to 2.1 times, about +3 to +6 bytes a
key in all, 5 to 10 % of the memory); keeping five sizes and making the code the same for all (D1c) is the cheaper way, if the
dispatch matters at all.
