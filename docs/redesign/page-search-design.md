# The search in a page: word compare or fingerprint (design note, 2026-10-07)

Words as in [GLOSSARY.md](GLOSSARY.md). Origin: analysis (b) of the review ([review-2026-10.md](review-2026-10.md)): a lookup of a
small `u64` map takes 33 ns against 17 in `main`, and 64 % of it is the search in the page (`locate` and `compare`); street single-value
58.6 against 47.1. The measurement of the first bytes on the real corpora (same file, last section) rules out a list of first bytes: of the
5 to 7 keys of a page, 3 to 5 share the first byte.

## Where the time goes today

`locate(m, kl, n, rem, r)` walks the key lengths; for every first slot it compares the remainder with the rest of the key byte by byte
(`compare`: a loop up to the first difference) and adds its length to the offset of the next remainder; it stops at the first remainder
that is not smaller. A hit compares about half the keys of the page (3 to 8), each up to its first differing byte; a miss as many. The
compare loop and the sum are serial: each step waits for the one before.

`main` reaches the key with a few word-wide matches in its byte nodes and one compare of the leaf's key: no linear search.

## Option 1: compare a word at a time (no change of the layout)

`compare` of two remainders loads 8 bytes of each at once (big-endian, so that the order of the words is the order of the bytes), compares
them as integers and goes to the next word only if they are equal; the last, shorter word is read as a whole word from the page and masked
to its length. A page is at least 32 bytes and the values follow the remainders, so an 8-byte load from the start of a remainder never
leaves the object; only at the very end of the object, where the values of a `Fixed` page are, would it need a guard, and remainders never
lie there. The rest of the key, the caller's slice, is loaded the same way with a guard at its end. The search stays linear and in order,
so it serves lookups, the place of a new key and the bounds of a range alike.

**Prediction:** a compare of remainders of 7 to 25 bytes from 3 to 10 ns to 1 to 3 ns; lookups `u64` single-value 4,096 33 → **26 to 29 ns**,
street single-value 58.6 → **52 to 56**, dirs (remainders of 15 to 25 bytes) a little more; writes 3 to 8 % faster (the add and remove of a
key call `locate`); memory unchanged.

## Option 2: a fingerprint byte a key (one more list in the page)

A list of n bytes behind the key lengths (many-key form only), one byte a slot: for a first slot a hash of its remainder (the length and the
first and last up to 8 bytes, multiplied by a constant, the top byte), for a further value a byte that is never compared (the key length
says `Further`). A lookup of a key (`Get`, `EachValue`, the add of a value to a key that is there, the remove of a value or a key) hashes
the rest of the key once, matches the byte against 8 fingerprints at a time with SWAR (`swar.Index8` / a mask of equal bytes, as the byte
nodes do), and compares only the remainders of the matches, on average one; the offset of a candidate's remainder is the sum of the key
lengths before it (a SWAR sum of the bytes, the `Further` slots masked to zero). A miss is decided without a compare in most cases. The
place of a **new** key and the bounds of a range still need the ordered search (`locate`), since a hash says nothing about order: the add of a
new key gets slower by the hash and by one more list to shift.

**Prediction:** lookups `u64` single-value 4,096 33 → **20 to 24 ns** (`main` 17), street single-value 58.6 → **45 to 50** (`main` 47); a miss
of a key 30 to 60 % faster; the add of a value to a key and every remove 5 to 15 % faster; the add of a new key 2 to 5 % slower; ranges
unchanged; memory **+1 byte a key** in multi-key pages (street single-value 27 → 28 block bytes a key, +3 to 4 %; a page holds that much
less, so a few more pages burst).

## Both

Option 1 makes every compare cheaper, also the ones option 2 leaves (the candidate, the place of a new key, the bounds of a range); they do
not exclude each other. In the order of cost: option 1 first (no layout change, every path profits, no memory), measured; then option 2 if
the lookups are still well behind `main`.

## How it is measured

The lookup probe against `main` (`lookprobe_test.go`: `u64`, street, dirs; single-value and natural; 4,096 and 65,536 keys), the write probe
(build and replay), the memory of the probe, the range probe (it must not get slower); then the cells of E5 on the PC (`main`, `node-layout`,
`btree`). Tests: the model tests of the page (`TestPageAgainstModel` and the fuzz) cover the search through every operation; a test of the
word compare at every length 0 to 24 and every position of the first difference, and at the end of the object.

## Option 1 result (2026-10-08): missed, dropped

Built as described (the key's first word once, one 8-byte load from the page a slot, masked to the shorter length; the rest word by word),
all tests and a new test of the compare at every length and difference (`TestPageCompareOrdersAsBytes`, kept) passing. Lookup probe against
`main` (WSL, medians of 9), ns a lookup, before → option 1: `u64` single-value 4,096 32.8 → **40.5** (+23 %), street single-value 57.9 →
59.6, dirs 86.4 → 86.5, `u64` natural 45.2 → 46.4, street natural 65.9 → 67.8, `u64` 65,536 35.4 → 36.5, street 65,536 85.6 → 87.1.
**Prediction missed in the sign** (predicted -10 to -20 %).

**Cause** (profile, `u64` single-value): the byte-wise compare was already cheap where it runs: random keys differ in the first byte, and
the remainders of a real page differ within the first bytes after the page's common prefix, so the loop ends after one or two bytes. The
word path adds work to every slot (a variable shift for the mask, `min`, a bounds-checked load with a byte swap, two more branches: the
mask and the compare lines alone took 19 % of the time). **The cost of the search is the number of slots it walks, not the width of a
compare**: each slot is a serial chain (its key length, the offset of its remainder, a load from there, a compare). The code is back as it
was; the lever left is to walk fewer slots: option 2 (a fingerprint, on average one candidate) or a search that does not walk (offsets of
the remainders, a binary search).

## The fingerprint: collisions within the pages of the real corpora (2026-10-08)

**Hash** (the user's proposal): wyhash's mixing as optimized in the Set3 project (`hashing.WH64Det`, 2.72 ns a call on the M1), on a block
of 16 bytes: the length of the input (2 bytes) and its last up to 14 bytes, right-aligned with zeros in front (the length tells `"a"` from
`"\x00a"`), as two words; the hash of the first word is the seed of the second. Measured in `Map.Shape` (tag mkstats, `shape_on.go`) on the
trees as built from the corpora: for every key of a multi-key page, how many **other keys of the same page** have the same fingerprint (the
compares a lookup would make in vain), against chance ((keys of the page − 1) / 128 or / 256). Three inputs: (A) the remainder of the key in
its page; (B) the whole key; (A') the remainder, but its first 6 and last 8 bytes when it is longer than 14.

| corpus | keys a page | (A) 7 bits | (A) 8 bits | (B) 8 bits | (A') 8 bits | chance 8 bits | keys with a twin (A, 8 bits) | most twins in a page |
|---|--:|--:|--:|--:|--:|--:|--:|--:|
| street natural 212,449 | 5.6 | 0.066 | 0.033 | 0.033 | 0.033 | 0.033 | 3.2 % | 3 |
| street single-value 212,449 | 7.1 | 0.092 | 0.045 | 0.046 | 0.046 | 0.046 | 4.4 % | 3 |
| dirs natural 86,215 | 4.9 | 0.061 | 0.038 | 0.038 | 0.044 | 0.025 | 3.4 % | 8 |
| dirs single-value 86,215 | 5.9 | 0.080 | 0.049 | 0.048 | 0.055 | 0.032 | 4.4 % | 8 |
| street single-value 4,096 | 7.1 | 0.092 | 0.041 | 0.038 | 0.042 | 0.043 | 3.9 % | 2 |
| dirs natural 4,096 | 4.0 | 0.044 | 0.024 | 0.031 | 0.022 | 0.017 | 2.3 % | 2 |

**What it says:** for street the fingerprint collides exactly as often as chance; for dirs about 1.5 times as often, because path keys of
the same length and the same last 14 bytes differ only in the middle — the worst page holds nine keys like
`ervice/mgmt/2019-08-01/containerservice/`, `…/2020-03-01/containerservice/`, …, which share length and ending (16 bytes of
`/containerservice/`). Taking the first 6 bytes instead of 6 of the last (A') does not help there (the dates differ at the seventh byte)
and is a little worse elsewhere. The whole key (B) is no better than the remainder (A), as expected (the path and the key part are the same
for all keys of a page). **In numbers:** with 8 bits a lookup compares on average 0.02 to 0.05 keys in vain (today: half the page, 3 to 4);
3 to 4 % of the keys have a twin in their page; the worst page (nine twins) is searched as today. 8 bits halve the collisions of 7.

---

## Option 2: implementation spec (decided 2026-10-08; written to be built as it stands)

The user decided to build option 2 after the collision statistics above. This section is the whole task; it is written so that it can be
carried out without the conversation that led to it. Working rules that apply (PLAN.md, "Working rules"): the conversation with the user is
German, documents and code comments English; words as in GLOSSARY.md; every exported function and every test has the doc comment
AGENTS.md asks for (tests outside-in: user, context, expectation); 100 % coverage, `go test ./... -race`, `golangci-lint run` 0 issues before
every commit; commit messages imperative, at most 72 characters in the subject; one measurement at a time on the machine (WSL and the PC
are the same machine); **stop and report to the user at any surprise** (a prediction missed in its sign, a failing test of the tree that is
not a page-size number, a result that is not understood) instead of building an alternative.

### What a fingerprint is

`fp(r)` for a remainder `r` (the bytes of a key after the page's key part): the lowest 8 bits of a wyhash-style mix of a 16-byte block. The
block is the length of `r` as 2 bytes big-endian, then the last `min(len(r), 14)` bytes of `r`, right-aligned in the remaining 14 bytes
with zeros in front. It is hashed as two little-endian words `w1 = block[0:8]`, `w2 = block[8:16]`: `h = wh64(w2, wh64(w1, 0))`, where

```go
// wh64 is hashing.WH64Det of github.com/TomTonic/Set3 (branch new_hasher): wyhash's mixing for one 64-bit word.
func wh64(val, seed uint64) uint64 {
	const m5, p1 = 0x1d8e4e27c47d124f, 0xf20a3e5e0b7b9731
	hi, lo := bits.Mul64(val^p1, bits.RotateLeft64(val, 32)^seed)
	hi, lo = bits.Mul64(m5^8, hi^lo)
	return hi ^ lo
}
```

and `fp(r) = byte(h)`. This is exactly the function measured above (`fingerprint` in `internal/art/shape_on.go`, input (A)); when it moves
into `internal/page` (file `internal/page/fingerprint.go`, unexported `fingerprint(r []byte) byte`), the diagnostic in `shape_on.go` must call
the page's function (export it as `page.Fingerprint` for that, documented) and its own copy goes, so that the statistics measure what the
page stores. Build the block with a `[16]byte` on the stack and `copy`; do not optimize the loads in this step (an over-reading variant is a
later step, only if the profile shows the hash above 15 % of a lookup). The fingerprint is not stable across versions and never leaves
memory.

### Where it is stored

Many-key form only (a one-key page has one key and no lists), a list of n bytes **directly behind the key lengths**:

```
head (4) | key part (l) | key lengths (n) | fingerprints (n) | value lengths (n, Str only) | remainders | free | values
```

Slot i's fingerprint is `fp(remainder of slot i)` for a first slot and **0** for a further value (`Further` in its key length); a search
skips candidates whose key length is `Further`. New offsets (`lay` in page.go): `kl = Header + l`, `fp = kl + n` (new field), `vl = fp + n`,
`rem = vl + (n if str)`; in the one-key form `fp` and `kl` are not used and `vl = Header + l` as today. `NeedStrings` becomes
`Header + cpl + 3*n + remBytes + valBytes`, `NeedFixed` `Header + cpl + 2*n + remBytes + n*size`. The package comment of page.go (the layout
drawing) and GLOSSARY.md ("The one page": a line for "fingerprints") are updated.

### Every place that computes offsets or moves the lists (all must take the new list into account)

Search `internal/page` for `Header+l`, `Header + l`, `kl+n`, `kl + n`, `vl`, `la.rem`, `la.vl`, `la.kl`, `+ n` and `2*n`; as of commit
`30bdc5d` these are:

- page.go: `lay` and `head.lay` (add `fp`), `keyEnd`/`keyEndFrom` (unchanged, they read the key lengths and take `rem` from `lay`),
  `locate` (signature unchanged; callers pass the new `rem`), `NeedStrings`, `NeedFixed`, `toOneKey` (the one-key form drops the key lengths
  and the fingerprints: the value lengths move to `Header + newl`, and everything between must be cleared), `oneKeyLeft` (reads key lengths
  only: unchanged).
- str.go: `BuildStrings` (`need`, the lists it fills: also the fingerprints), `Get` and `EachValue` (they compute `Header+l+n+n` for the
  remainders and `vl += n`: both shift by n; and they search with the fingerprint, below), `Add` (the written-out `insertSlot`: one more list
  to shift by one byte, the new slot's fingerprint written; the value lengths and remainders now sit one list further), `regrow`, `pairWith`
  (builds through `BuildStrings`: fine), `Remove` (the written-out removal of a slot: one more list to shift back), `Each`, `Skip`/`Prepend`
  (they move `Header..e` as a block: the lists come along; the fingerprints stay valid, since the remainders do not change), `Widen` (it
  rebuilds the lists into `out`: build the fingerprint list there too, and **recompute every fingerprint**, since every remainder grows by
  the bytes the key part loses), `Used`.
- fixed.go: the same functions for `Fixed` (`BuildFixedOf`, `Get`, `EachValue`, `Add`, `pairWithFixed`, `Remove`, `Each`, `Skip`, `Prepend`,
  `Widen`, `regrow`, `Used`, `fits`).
- scan.go: `slotRange` callers (`ValuesIn`: `Header+l+n` → plus n; `AppendStrings`: `vl += n` → `vl = kl + 2n` and its remainders `vl + n`),
  `AppendKeys` (`rem := kl + n` → `kl + 2n`, plus n for strings).
- Outside the package: nothing computes page offsets (the tree uses `NeedStrings`/`NeedFixed` in `mergeFits`, which follow). The model in
  `bench/cmd/skmodel` is not updated (it models the layouts of step 5 and is not used any more).

### The search with the fingerprint

A function in page.go, used by `Get`, `EachValue` (both flavors), and by `Add` and `Remove` to find a key that is there:

```go
// find returns the first slot of the key whose remainder is r in a many-key page (kl: offset of the key lengths, n: slots,
// rem: offset of the remainders), or false. It compares only the remainders whose fingerprint equals r's.
func find(m []byte, kl, n, rem int, r []byte) (pos int, found bool)
```

Compute `f := fingerprint(r)`. Walk the fingerprint list `m[kl+n : kl+2n]` eight bytes at a time (`swar.Index8` on a little-endian word; for
the last, shorter group build the word from the bytes left, padding with a byte that is not `f`, e.g. `^f`). For each match at slot i: if
`m[kl+i] == Further` or `int(m[kl+i]) != len(r)`, it is not the key (the key length is part of the check, cheap); else compute the offset of
slot i's remainder as `rem + sum of the key lengths of the first slots before i` (the `Further` ones count 0) and compare the bytes (`string(a)
== string(b)`). To find the next match in the same word after a candidate that was not the key, overwrite that byte of the word with `^f`
and call `Index8` again (its false positives lie only above a true match, so the lowest flagged byte is exact). Return the first slot found,
or false after the last group. `locate` stays as it is and is used only where the place of a key that is not there is needed: `Add` of a
new key (after `find` said false), `Widen`, and the scans (`slotRange` compares bounds, not keys: unchanged).

`Add`: `find` first; found → the further-value path as today (the new slot gets fingerprint 0); not found → `locate` for the place, then
the new-key path (the new slot gets `fingerprint(r)`). `Remove`: `find` instead of `locate`; the rest as today. `Get`/`EachValue`: `find`
instead of `locate`, then the run of the key (`runEnd`) as today.

### Tests (all in `internal/page`, plus what the tree's tests need)

- `TestPageLayout` pins bytes: add the fingerprint list to the expected bytes, computed with `fingerprint` (not as literal numbers), and
  check `Used()`; the classes of the example pages may change (a page needs n more bytes): adjust the expected class with a comment.
- A test of `fingerprint` against a reference that builds the block byte by byte as described (lengths 0 to 40, random bytes, 10,000
  inputs), and that it depends on the length (`"a"` and `"\x00a"` differ).
- A test of `find` with a crafted page in which two keys of the same length share a fingerprint (find two such remainders by brute force in
  the test, e.g. 3-byte strings): both are found, a third key with the same fingerprint and length but not in the page is not; a key whose
  fingerprint matches a `Further` slot's 0 is not confused with it; groups of more than 8 slots (a page of 20 keys) with matches in the second
  and third word.
- `Widen`: after a widen every key of the page is still found (its fingerprint was recomputed).
- `TestPageNeed` (layout_test.go) checks `NeedStrings`/`NeedFixed`: update its expected numbers by n.
- The model test (`TestPageAgainstModel` and its fuzz test `FuzzPage` in model_test.go, run with `go test -run ^$ -fuzz FuzzPage -fuzztime 60s ./internal/page`) must pass unchanged: they drive every
  operation through the new search.
- The tree (`internal/art`): tests that build pages of an exact size (e.g. values of 237 bytes so that a page is full, the tests of
  `mergeFill`) may need other numbers, since a page now needs n more bytes; change the numbers, not the logic, and list each such change in
  the results section. Any other failing test of the tree is a stop.

### Prediction (write the results against it)

Lookup probe against `main` (`MKLOOK=1`, ns a lookup, WSL, medians of 9), today → with fingerprints: `u64` single-value 4,096 32.8 → 20 to 26
(`main` 17); street single-value 4,096 57.9 → 45 to 52 (`main` 45 to 47); dirs single-value 4,096 86.4 → 70 to 80 (`main` 75); `u64` natural
4,096 45.2 → 35 to 41; street natural 4,096 65.9 → 52 to 60. Write probe (`MKPROBE`, build / replay ns an operation): single-value cells
±5 % (the add of a new key costs the hash and a list more to shift, the add of a value and every remove are faster); natural cells -3 to
-10 %. Range probe (`MKRANGE`): unchanged ±3 % (scans do not use the fingerprint; the lists are one longer). Memory (`MKSHAPE`, block bytes a
key as built): +1 byte a key in multi-key pages (street single-value 27 → 28, street natural 38 → 39, dirs natural 58 → 59); slightly more
bursts in the probe's events (a page holds n bytes less). If `u64` and street lookups are not at least 10 % faster: **stop** and report with a
profile; do not tune the hash or the loads before the user has seen it.

### How to measure (commands; run one at a time)

The lookup probe needs the bench built with `main` as the baseline (the bench's `bench/baseline` holds step 3.5 by default and is not in
git; save and restore it):

```sh
cd bench
cp -r baseline /tmp/baseline-keep && go run ./cmd/mkbaseline -ref main
go test -c -tags baseline -o /tmp/look-new.test ./cmd/bench
rm -rf baseline && cp -r /tmp/baseline-keep baseline      # check: tail -1 baseline/ref.go says fe120f5
# the same at the commit before the change, into /tmp/look-old.test (git worktree add /tmp/wt-old <commit>; copy baseline in; build there)
for c in "u64 single-value 4096" "street single-value 4096" "dirs single-value 4096" "u64 natural 4096" "street natural 4096"; do
  set -- $c; for v in old new; do MKLOOK=1 MKLOOK_KEYS=$1 MKLOOK_VALUES=$2 MKLOOK_N=$3 /tmp/look-$v.test -test.run TestLookProbe -test.count=1 | grep lookups; done
done
```

The write probe: binaries built with `-tags mkstats`, run interleaved with a small script (see `/tmp/claude-1000/bis3.sh` on this machine,
or write one: for each of 7 rounds run old and new with `MKPROBE=1 MKPROBE_KEYS=… MKPROBE_VALUES=… MKPROBE_N=…`, take the medians of the
lines `build stream:` and `plain replay`). The range probe: `MKRANGE=1 MKRANGE_KEYS=… MKRANGE_VALUES=… MKRANGE_N=…`. Memory and collisions:
`MKREAD=1 MKSHAPE=1 MKREAD_KEYS=… MKREAD_VALUES=… MKREAD_N=…` with `-tags mkstats` (prints the census with block bytes a key, the shape and
the fingerprint statistics). Then, if the probes meet the prediction, the cells of E5 on the PC against `main` and `node-layout`
(review-2026-10.md, E5: `bench -suite dev -skipmem -keys street,dirs,u64,url -values single-value|natural -sizes 4096,65536 -ops
valuesFor,valuesBetween,churn,build`, executables built with `GOOS=windows GOARCH=amd64 go build -tags baseline` after `mkbaseline -ref main`
resp. `-ref 7b8a8d8`, run through a `.cmd` file started with `powershell.exe Start-Process`; about 2.5 hours; say the end time to the user).

### Order of work and commits

1. `internal/page/fingerprint.go` with `fingerprint` and its test; the diagnostic in `shape_on.go` switched to it (commit).
2. The layout: the list in every place listed above, built and kept on every operation, `find` not yet used; all tests green, `TestPageLayout`
   updated (commit). Measure memory (`MKSHAPE`) here: the +1 byte a key must show.
3. `find`, used by `Get`, `EachValue`, `Add`, `Remove`; its tests (commit).
4. The probes against the prediction, results written into this file under "Option 2 result" (commit), push, and report to the user before
   the PC run.

---

## Coding of further values and fingerprints per key (2026-10-08, on the user's question)

Question: could the fingerprints start right behind the key part with one entry per key (not per slot), and is the `Further`
(FF) coding of the key lengths worth it against "every key gets a byte with its number of values"? Measured with `Shape()`
(`valueCoding`, tag `mkstats`) over the multi-key pages of 65,536 keys as built at 52b785c, `MKREAD=1 MKSHAPE=1`. n slots (values), K
keys of a page; bytes a key = bytes of the lists over all keys of all multi-key pages.

| keys | values | values a key | FF: key lengths n | counts: lengths + counts 2K | today kl+fp 2n | fp per key: n+K | all per key 3K |
|---|---|--:|--:|--:|--:|--:|--:|
| street (real keys and values) | single-value | 1.00 | 1.00 | 2.00 | 2.00 | 2.00 | 3.00 |
| street | natural (the localities) | 1.72 | 1.72 | 2.00 | 3.45 | 2.72 | 3.00 |
| dirs (real keys and values) | single-value | 1.00 | 1.00 | 2.00 | 2.00 | 2.00 | 3.00 |
| dirs | natural (the file names) | 2.18 | 2.18 | 2.00 | 4.36 | 3.18 | 3.00 |
| path (real keys, generated values) | single-value / natural | 1.00 / 2.95 | 1.00 / 2.95 | 2.00 | 2.00 / 5.90 | 2.00 / 3.95 | 3.00 |
| url (real hosts, shaped paths; generated values) | single-value / natural | 1.00 / 2.93 | 1.00 / 2.93 | 2.00 | 2.00 / 5.86 | 2.00 / 3.93 | 3.00 |
| u64, str, uuid, email (generated) | single-value / natural | 1.00 / 2.91 to 3.03 | 1.00 / 2.91 to 3.03 | 2.00 | 2.00 / 5.8 to 6.1 | 2.00 / 3.9 to 4.0 | 3.00 |

Pages where FF is smaller than the counts (n < 2K): single-value 100 %; street natural 77 %; dirs natural 61 %; the generated
natural profiles 33 to 37 % (equal in 11 to 13 %). The generated natural profile has 2.9 to 3.0 values a key; the two real
profiles have 1.7 and 2.2, so the counts look better on generated keys than they are on real ones.

Reading: the break-even of FF and counts is 2 values a key. With one value a key (the profile of the typical index) FF costs half;
on street natural it is smaller too; on dirs natural the counts save 0.18 B a key. FF also keeps the slot of a match direct (the value
array is indexed by slot), where counts need a prefix sum for it. A fingerprint per key (n+K instead of 2n) saves n-K bytes a page:
0 for single-value, 0.72 B a key on street natural, 1.18 on dirs natural. Cost of it: K is not in the head (the lists behind it
start at 11+K), so one byte of the page must hold K, and a match at key k must be mapped to a slot (one pass over the key lengths
counting keys, the pass that sums the lengths for the offset anyway). Not built: the fingerprint has not yet paid in speed (see
"Option 2 result").

---

## Search without a hash: (length, first byte) with SWAR (2026-10-08, thought through on the user's question; not built)

**Why a hash is not needed.** Collisions within a page of another key's signature, single-value profile, 65,536 keys (`cheapKeys` in
`shape_on.go`, other keys of the page with the same signature a key; "≥1" is the share of keys that have one or more):

| keys | fingerprint, 8 bits | length + first byte | length + first and last byte | length + first two bytes |
|---|--:|--:|--:|--:|
| street (real) | 0.041 (4.0 % ≥1) | 0.360 (23.4 %) | 0.093 (7.6 %) | 0.178 (12.3 %) |
| dirs (real) | 0.044 (4.0 %) | 0.332 (17.2 %) | 0.332 (all end in `/`) | 0.286 (14.2 %) |
| path (real keys) | 0.034 | 0.330 (15.9 %) | 0.269 | 0.269 |
| url | 0.016 | 0.071 | 0.012 | 0.050 |
| u64 | 0.006 | 0.004 | 0.000 | 0.000 |
| str / uuid / email (generated) | 0.058 / 0.009 / 0.015 | 0.112 / 0.120 / 0.011 | 0.010 / 0.006 / 0.005 | 0.022 / 0.007 / 0.000 |

A false candidate costs one failed compare (about 2 ns); 0.36 of them (street) cost less than one hash (4 ns at best, 8.6 ns as built:
`Fingerprint` of 7 bytes measured alone, 6.0 ns with the block built in registers, 4.1 ns the two `wh64` alone).

**Kernel measured alone** (scratch benchmark, not in the repository; `internal/page`, a page of K street-like remainders of 8 to
11 bytes with 6 first letters, 65,536 random lookups of present keys, ns a lookup, medians of 3, WSL):

| K | `locate` (today's order search) | hash + `Index8` (as committed) | SWAR on first byte and length, both lists | SWAR on first byte, length per candidate |
|--:|--:|--:|--:|--:|
| 4 | 12.0 | 19.9 | **9.9** | 10.3 |
| 7 | 15.3 | 22.4 | **11.2** | 13.1 |
| 12 | 19.4 | 25.4 | **13.8** | 15.7 |
| 16 | 21.6 | 26.2 | **17.7** | 19.9 |

So the hash search is slower than `locate` in the kernel by 7 ns, as in the lookup probe, and the hash-free SWAR search on both
lists is 2 to 6 ns faster than `locate`. Of a lookup of 57 ns on street (4.9 nodes on the way) that is at most 4 to 10 %: the page
search is a part of the lookup, the descent is the larger one.

**Layout that follows (smallest change from the committed one).** The list of fingerprints keeps its place and its length (one byte for
every value, 0 for a continuation); the byte is the first byte of the remainder (0 for the empty remainder) instead of the hash. The
key lengths are the second list, parallel to it, and a continuation's `Further` (255) is never a length searched for. The search is
`zero bytes of (list ^ first) & zero bytes of (key lengths ^ length)` per eight values, exact (no borrow false positives), the lowest set
byte is the candidate; the offset of its remainder is the byte sum of the key lengths before it, less 255 for each `Further`
(`sum8(w) - 255 * count of FF bytes`); then one compare of the remainder. Nothing else changes: no byte for the number of keys, no
change of the head, `Fingerprint` and `wh64` go.

**The "simple way" (a byte for the number of keys K, lists per key).** It saves n-K bytes a page (0 for single-value, 0.7 B a key
on street natural, 1.2 on dirs natural) and costs a byte a page and a mapping from the key's index to its value (a pass over the key
lengths counting keys, unless the page has no continuation, V = K, where they are the same). The lists are then no longer parallel
for pages with continuations. Not now; it is a memory optimization of 1 to 2 %.

**Prediction for the hash-free search** (lookup probe, ns, today → predicted, `main`): u64 single-value 4,096 33.5 → 28 to 31 (17);
street 57.8 → 52 to 55 (47); dirs 87.4 → 82 to 86 (75). The ceiling is the share of the page search in a lookup, about a third on
street: even a free search would not reach `main`. Writes: the add of a new key shifts the same lists as the committed layout, no hash:
better than the committed layout, about as `locate`. Memory: as the committed layout (+1.5 to +2.9 B a key against 5a63bd0).

---

## Option 2 result (2026-10-08, written 2026-10-10) and what the numbers of this note are worth

Option 2 as specified was built (`5a63bd0`, `828bf1f`, `52b785c`) and missed: lookup probe, ns a lookup at 4,096 keys, before → after
(`main`): u64 single-value 33.5 → 42.2 (17.3), street single-value 57.8 → 57.1 (46.5), dirs single-value 87.4 → 84.8 (75.6), u64
natural 46.1 → 44.8 (36.6), street natural 65.9 → 67.2 (53.0); predicted u64 20 to 26, street 45 to 52. The profile: `Fingerprint`
24 % of a u64 lookup (building the 16-byte block in memory, two wyhash mixes in a chain); memory +1.5 to +2.9 block bytes a key. Taken
out again on 2026-10-10 (`6e76275`); the renaming to `currentValues` and the diagnostics stay.

**Not rtcompare.** The speed figures of this note (options 1 and 2, the cost of `Fingerprint`, the search kernel of the last section)
come from the lookup probe (its own loop, medians of 9 runs of 100 ms, the versions run one after the other) and from `go test -bench`
scratch benchmarks, not from interleaved rtcompare comparisons as MEASURING.md requires for speed claims. They agree with the profiles,
but they are not verified. If the page search is taken up again, its variants are compared as rtcompare candidates (in `cmd/bench`, as
the `ordered-mkN` candidates of exp-maxkeys) before any decision.
