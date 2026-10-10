# What the descent's time is made of (2026-10-10)

Words as in [GLOSSARY.md](GLOSSARY.md). Lever 1 of the order after the profiles (profiles-2026-10.md): the descent through the byte
nodes is 41 to 55 % of a lookup of the structured keys and 24 to 40 % of a churn. Before choosing a structure or a code change, find
out what that time is: waiting for dependent loads (only fewer levels or fewer objects help: structure), mispredicted branches
(the switch on the node type, the loop exits: a branch-poor step helps: code), or the work of the steps themselves.

## Method

No new measurement: the CPU profiles of profiles-2026-10.md (code at `6e76275`, unchanged in the library since), resolved to single
instructions with `pprof -disasm` for `Tree.find` and what is inlined into it. A timer sample lands on the instruction that waits to
retire: after a load that misses, on the first instruction that uses its result (or the load itself); after a mispredicted branch,
on the first instructions of the right path. So samples are put into:
- **load wait**: the instruction that loads, or first uses, the node's type byte, the key word of a node search, the prefix word, the
  child pointer;
- **branch**: the first instructions at the targets of the type switch and of the loop's exits;
- **work**: the rest (SWAR arithmetic, key byte, bounds).

There are no hardware counters here (no `perf` in WSL); the split is an attribution by position, rough by a few instructions.

## Prediction (written before the instruction view; the function top list of dirs single-value 4,096 was seen: `isPage`, the
## load of the type byte, has 12 % of the whole lookup)

- load wait is the larger part of `find`'s own time: at least 50 % at 4,096 keys and at least 65 % at 65,536;
- branch at most 25 % at every size;
- consequence if it holds: the descent is a structure question (fewer dependent objects a lookup), a branch-poor node step gains at
  most a tenth of the descent.

## Run

The probe got a knob for the length of its profiles (`MKLOOK_SECS`). Lookup profiles of 10 seconds each (`MKLOOK`, as the bench's
valuesFor), the seven structured kinds (str, email, url, path, street, dirs, links), single-value, 4,096, 16,384 and 65,536 keys;
WSL on the PC, nothing else running, 2026-10-10 12:43 to 12:51. Code `6e76275`'s library (unchanged). The samples of each
instruction of `Tree.find` (with what is inlined into it), summed over the seven kinds, by `annot.py`; the categories by address
(`cat.py`); both in `bench/results-layout/descent-2026-10/`.

## Result

`find`'s own time is 40 % of a lookup at 4,096 keys, 53 % at 16,384, 57 % at 65,536 (the search in the page and the reading of the
values are the rest). Inside it:

| % of `find` | 4,096 | 16,384 | 65,536 |
|---|--:|--:|--:|
| waiting for the **type byte of the next object** (the first line of every node and of the page, `isPage`) | 34 | 46 | 57 |
| waiting for the **child pointer** in the node (its line differs from the line of the keys or the bitmap) | 19 | 18 | 16 |
| waiting for other loads in the node (prefix length, prefix word, key words, bitmap) | 11 | 8 | 6 |
| **load wait in all** | **64** | **72** | **79** |
| rest: switch, loop, SWAR work, the check of the page type | 36 | 28 | 21 |

By kind the load wait is 56 to 82 % at 4,096 keys and 67 to 93 % at 65,536; the rest is largest for the path-like keys at 4,096 keys
(dirs 45 %, path 41 %, url 30 %: many levels whose lines are in the cache, so the work of each level shows), smallest for email (7 %
at 65,536). The single largest wait inside a node is the child pointer of an N26 (7 % of `find` at every size): its bitmap is in the
node's first line, its pointers in the three lines behind it, so an N26 costs two dependent lines.

**Against the prediction:** load wait at least 50 % at 4,096 and 65 % at 65,536: hit (64 and 79 %). Branch at most 25 %: not
separable from the work by position; the rest as a whole, an upper bound for both, is 21 to 36 %. The consequence holds in its first
half: the descent is mainly a question of how many lines a lookup waits for one after the other, not of the code of a step. Its
second half (a branch-poor step gains at most a tenth) is not shown: such a step could take a part of the 21 to 36 %, of up to 45 %
for small maps of path-like keys, and how large a part is unknown.

Two parts of the wait are separate levers:
- the **objects on the way** (57 % at 65,536): only fewer levels or fewer objects a lookup reduce it (the structure of step 6);
- the **second line inside a node** (16 to 19 %): N12 (child 4 and up), N26 and N58 keep the searched bytes and the pointer in
  different lines. That is a layout question of the node alone, without the maintenance problem of packing. On the M1 (128-byte
  lines) an N12 is one line, so this part may be smaller there; these profiles are the PC's only.

## Lever 2 counted in the model (2026-10-10 afternoon)

`rangemodel -nodelines` (today's tree, exact against the measured shape; `bench/results-layout/descent-2026-10/nodelines-model.md`):
the byte nodes a lookup passes by class, and how many of them read the child slot from a second line of the node.

| single-value | objects a lookup | second lines a lookup, 64-byte lines (PC) | 128-byte lines (M1) |
|---|--:|--:|--:|
| str 4K / 64K | 3.4 / 4.8 | 1.5 / 2.0 | 0.7 / 0.7 |
| email 4K / 64K | 3.0 / 4.0 | 1.9 / 2.8 (all N26) | 1.2 / 1.8 |
| url 4K / 64K | 6.9 / 9.2 | 2.8 / 4.5 | 1.9 / 2.9 |
| path 4K / 64K | 9.4 / 12.8 | 3.1 / 5.5 | 1.0 / 1.9 |
| street 4K / 64K | 4.0 / 5.9 | 2.1 / 3.5 | 0.7 / 1.8 |
| dirs 4K / 64K | 8.6 / 12.4 | 2.9 / 4.5 | 1.1 / 2.4 |
| links 4K / 64K | 3.7 / 5.3 | 2.4 / 3.4 | 1.4 / 2.6 |

On 64-byte lines 40 to 90 % of the nodes passed read a second line (all N26 and N58 visits but a few, N12 from the fifth child);
on 128-byte lines a third to two thirds as many remain. The lever exists on both machines, smaller on the M1; the M1's profile
(by source line, so that it reads on arm64) is still to be taken.

## The M1 (job lp1, 2026-10-10 21:16 to 21:27)

Lookup profiles of 10 seconds on the M1 (128-byte lines), the six structured kinds but links (its cache was not found by the job),
single-value, 4K/16K/64K, the ordered map alone, code `a0e75c5` (the library of `6e76275` with the counters of the evening, which
cost nothing without mkstats). The probe printed the source lines with the most own time (`m1-lp1.log`); `m1lines.py` sums the lines
of `Tree.find` and what is inlined into it over the kinds (`m1-lines.txt`). Lines, not instructions: on arm64 the wait for a load
shows on the line of the first instruction that uses it, which may be the next line, so the split is rougher than on the PC.

`find` is 41, 49 and 50 % of a lookup (PC: 40, 53, 57). Inside it (% of `find`, 4K / 16K / 64K):

| line | what | M1 | PC (same place) |
|---|---|--:|--:|
| node.go:131 (`isSingleKey`), lookup.go:34/35/38/39/43 | the type byte and header of the next object, the loop head that uses the child pointer | 38 / 39 / 35 | 34 / 46 / 57 (type byte) + 2 to 4 (loop head) |
| lookup.go:85 | the child pointer of an N26 (second line on the PC) | 0.4 / 0.4 / 1.0 | 7.3 / 7.3 / 7.3 |
| lookup.go:60 to 63, 80, 86 | the key byte and the switch on the node type (its targets) | 24 / 26 / 24 | in the rest of 36 / 28 / 21 |
| swar.go, lookup.go:66/73/76 | the SWAR searches and prefix compares | 27 / 27 / 26 | in the rest |

So on the M1 the second line inside an N26 costs next to nothing (as the model says: on 128-byte lines most child slots of an N26
are in the node's first two lines, and the M1 seems to have them both), and the switch on the node type is a quarter of `find`.
**Lever 2 (the second line inside a node) is a lever of the PC; on the M1 the dispatch of the node step is the larger part.** The
wait for the next object is the largest part on both machines.
