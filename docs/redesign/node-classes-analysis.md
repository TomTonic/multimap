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
