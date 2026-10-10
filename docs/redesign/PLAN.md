# Plan: pages for every key

This plan turns [STRATEGY.md](STRATEGY.md) into code, step by step. Each step ends at a gate: a
measurement against fixed criteria, a short report to the user, and the user's decision.
Nothing moves to the next step without that decision.

Branch: `cacheline`, forked from `node-pages` (`node-pages` and `node-layout` are parked and
pushed). Worktree: `/mnt/c/temp/code/multi_map-layout`. `main` stays untouched until the user
decides to merge.

**Revised 2026-10-04** (review of the night of 2026-10-03/04): the order of the work is now
single-key page (SKMV) first, then the multi-key page (MKSV) as a special case, then multi-value
entries in multi-key pages (MKMV) last. The base case is `string -> {string}`: a key with a set of
strings. The rules below that are marked *new* come from that review.

## Working rules

- **Language.** Talk to the user in German. Code, comments, commit messages and documents in the
  repository are English.
- **Explainable before fast** *(new)*. The user must be able to follow every step: what is built,
  why, and what the measurement says about it. Optimizing without a model of the cause is not
  allowed. Concretely:
  - **Design before code.** Every sub-step that builds an object starts with a design note (at
    most about two pages, in the words of [GLOSSARY.md](GLOSSARY.md)): the layout byte by byte, the
    operations, the rounds of cache-line loads of each (R5), and a **prediction** in numbers
    (memory per key from a model of the real data, rounds per lookup). The user approves the note
    before code is written.
  - **Predict, then measure.** The report compares the measurement with the prediction. A deviation
    of more than about 10 % is explained (cause found and shown, for example with a profile or an
    object statistic) before anything else changes.
  - **No knob without a reason.** Every constant (size classes, thresholds, fill limits, header
    sizes) comes from the design with its reason: a cache line, a size class, an R rule, a share of
    the real data. Sweeping a constant to pick the best value is not a reason; a sweep may only
    confirm that the chosen value is not next to a cliff. New tuning variables (`var X` for
    experiments, environment variables) need the user's approval.
  - **One question per run.** Write down the question a benchmark run answers before starting it.
- **Stop at surprises** *(new)*. If a measurement says that a decision of the user or the design
  does not work, stop, write it up with the numbers and the options, and ask. Do not build the
  alternative. (The night of 2026-10-03/04 built multi-value entries in multi-key pages where the
  multi-key page with single-value entries had been asked for.)
- **Glossary** *(new)*. Code, comments and documents use the words of [GLOSSARY.md](GLOSSARY.md).
  A new concept gets a glossary entry before it gets a name in code. A retired word in new or
  changed code is a review finding. Up to 2026-10-04 the code still spoke the old words (`leaf`,
  `term`, `depth`, `inner`, `settle`, `crowded`, `suffix`); step 3.0 fixes that.
- **Commits.** Commit and push at every valuable point without asking (user decision 2026-10-02);
  push from WSL with `git -c credential.helper= -c credential.helper='!gh auth git-credential' push origin cacheline`.
  **Never rewrite history** *(new)*: no `filter-branch`,
  `rebase`, `reset --hard`, `commit --amend` of pushed commits, or force-push. A file committed by
  mistake is removed in a new commit. Before every commit run `git status` and check that no binary
  or scratch file is staged (`*.test`, `bench/cmd/bench/bench`, `*.exe` are ignored).
- **Commit messages.** Follow AGENTS.md: imperative mood, subject at most 72 characters, a
  blank line, a body that says why. End with the `Co-Authored-By` trailer of the model that
  did the work.
- **Code checks.** Every commit keeps:
  - 100% coverage per package (prototypes in their own package included);
  - `go test ./... -race` passing;
  - `golangci-lint run` clean (the generated copy in `bench/baseline` is not ours);
  - the fuzz test running 60 s without a finding (`go test ./internal/art -run '^$' -fuzz FuzzOperations -fuzztime 60s`).
- **Test docs.** Test documentation reads outside-in (AGENTS.md).
- **Weight of random keys** *(user, 2026-10-10)*. A range index is used in practice on keys that are not scattered at random
  (text, paths, names, addresses); for random keys (u64, uuid) a user wants point queries and takes the hashed map, which is
  unbeatable there. Ordered stays a general-purpose structure, but in every trade-off the structured keys (str, email, url, path,
  street, dirs) decide; the random ones must not fall out of the band, they do not set the course.
- **Hot path in "assembler thinking"** *(user, 2026-10-10)*. On the hot paths (lookup, add, remove, the range walk): no
  reflection, no interfaces or other dynamic dispatch, no yield chains or iterators, no closures as callbacks, no channels, few
  or no method calls; plain loops over bytes and words, the code written out where Go's inliner would not inline it (a Go stack
  frame rarely pays). No real assembler. The public API may return `iter.Seq`; behind it, one plain loop per page that the
  compiler inlines into the caller's range loop (checked with `-gcflags=-m`), as the range walk does since scan-design.md.
- **Measuring.** See [MEASURING.md](MEASURING.md): claims only from interleaved rtcompare runs,
  durations with clock time, a quiet machine during runs, one run at a time on the PC. Tell the
  user the duration and the expected end of every run longer than a few minutes. Jobs for the M1
  only through `bench/remote/queue.txt`.
- **Data.** The real data sets `street` and `dirs` (real keys, real value counts, real names as
  string values) decide. A design is never discarded because of a synthetic corpus; synthetic keys
  (`u64`, `str`, ...) may show a problem, not settle a decision.
- **Reports.** Report faithfully. Numbers that speak against a step go into the report and the
  results README just like the good ones. A gate that is missed is reported as missed, with the
  options. It is not explained away.
- **Backups.** Keep backups of uncommitted work outside WSL's `/tmp` (for example
  `C:\temp\bench-win`). Never `git checkout` a file with uncommitted edits.
- **References.** `node-layout` (`7b8a8d8`: byte nodes and a leaf per key, no pages) is the
  reference of step 3: the single-key page replaces its leaves one to one. The competitors
  (`btree-sets`, `map-sets`, `hashed`; `btree-map` for one value per key) are the reference for
  credo 1 and 2.

## Steps 0 to 2 (done)

- **Step 0, groundwork** (2026-10-02, gate 0 met): the object statistic (`art.Map.Objects`,
  `bench/cmd/objstat`), the probe order of `valuesFor`, the arm64 job queue (`bench/remote/`), a
  shorter race test. A/A runs on the M1: noise about ±1 % for point operations, ±3 % for ranges.
- **Step 1, the page prototype** (2026-10-02, [step1-results.md](step1-results.md)): the multi-key
  page of layout A (`internal/vpage`, a directory of tags), two rounds per lookup (R5 restated as
  rounds), the common prefix of a page.
- **Step 2, pages in the tree for every key with one value** (2026-10-02,
  [step2-design.md](step2-design.md), [step2-results.md](step2-results.md)): gate 2 not met; the user
  ticked it off on 2026-10-03 as "not met, deficits noted" (small maps and `u64` slower; ranges,
  memory and GC much better).

## Step 3 so far (2026-10-03/04): what was learned before the order changed

Three pieces of work are done and stay valid as input; none of them is a decision.

- [step3-layout.md](step3-layout.md): the two candidate layouts of the multi-key page, A
  (`internal/vpage`, directory of tags) and B (`internal/lpage`, a header of lengths, the user's
  sketch), compared standalone. B is smaller, A faster for short keys.
- [step3-real-data.md](step3-real-data.md): today's tree on the real data. With string values its
  range scans are 2.6 to 4.2 times slower than `btree-map`'s, because of the leaf per key with a
  string header.
- [step3-tree-pages.md](step3-tree-pages.md): layout B in a copy of the tree, for string values.
  **On the natural mix of `street` and `dirs`, the fall back takes away every multi-key page**, so the
  tree is nodes plus one object per key: for real multi data, the single-key page is what decides.
  The same night built multi-value entries in multi-key pages (`Map.Pairs`, MKMV) as an option; that
  code is parked on the branch `mkmv-experiment` (`2adf119`, pushed) and comes back in step 5. Its
  numbers (memory -36 to -40 %, mutation 0.4 to 0.9 of today's) are the starting point there.

## Step 3: the single-key page (SKMV) (decided 2026-10-04)

**Goal.** One object per key that holds the key's remainder and all its values, replacing flat,
typed and set leaves, and efficient for the four cases the user named:

| case | key | values | why |
|---|---|---|---|
| base case, first | `string` | set of `string` | the user's main use; values stored as bytes |
| | `string` | set of `uint64` | fixed-size values without pointers |
| | `string` | set of `*T` | values with pointers: the page is typed for the GC |
| | `uint64` | set of `*T` | integer keys with pointers as values |

**The tree in this step** (decided 2026-10-04): byte nodes and one single-key page per key. A key
with one value has a single-key page with one value (SKSV is the special case of SKMV). The
multi-key pages of step 2 and the fall back are **switched off** in this step, so that the
measurement shows the single-key page against the leaf it replaces and nothing else. They come back
in step 4. (A switch in `internal/art` is needed; today `crowdedRatio` is a constant and pages are
always on.)

**String values** (decided 2026-10-04): stored as bytes inside the page, pointer-free; a lookup
copies them out (one allocation for the values of a key, as `EachString` in `internal/lpage` does).
Zero-copy (immutable pages, strings that are views) is not pursued. For comparison, the layout for
fixed-size values with `T = string` (16-byte headers, the page typed for the GC, no copy) is
measured once in 3.3, since that layout exists anyway for pointers.

### 3.0 Groundwork (no new structure)

1. **Glossary in the code.** Pure renames in `internal/art`, `internal/vpage` and `internal/lpage`,
   one commit per package, no change of behaviour (tests unchanged). Rename what survives step 3:
   `term` to end page, `depth`/`base` to path length, `inner` to byte node, `settle`/`crowded` to fall
   back, `suffix` to remainder, `class` to size class where it means one. `leaf` is not renamed: the
   leaves disappear in 3.3.
2. **The bench reads the values it times.** Today the timed loops read only the header of a string
   value (`weigh` in `bench/cmd/bench/value_str.go`), which favours every candidate that holds
   pointers to the caller's strings: a real caller reads the bytes and pays the cache miss there.
   Make `weigh` read the length and the first byte. Note in `bench/README.md` that string-value numbers
   before and after are not comparable.
3. **Pointer values in the bench.** A build tag `ptrvals` with `V = *T` (`T` a small record of 16
   bytes, one object per distinct value, as an application would have), next to `strvals`.
4. **The reference runs** (PC; then the same as M1 jobs): `node-layout` (`mkbaseline -ref 7b8a8d8`,
   check that it builds with today's bench) and the competitors, for the four cases on `street`
   and `dirs` (multi and unique; sizes 4,096, 16,384 and the corpus) and `u64` keys with `*T`
   values (4K, 16K, 256K). These are the numbers step 3 is measured against.

### 3.1 Design note (`step3-skmv-design.md`, approved by the user before code)

Start from the user's sketch ([whataleafneedstostore.md](whataleafneedstostore.md), SKMV cases 2a
to 2d). The note must answer:

- **Layout for variable-length values** (strings): kind byte, remainder length, number of values,
  value lengths, remainder bytes, value bytes. Where the remainder lies relative to the value
  lengths, so that the key check of a lookup needs one round and `ValuesFor` two at most (R5).
  How many values fit: what fits, not a number fixed by the header. How a value of more than 255
  bytes is stored.
- **Layout for fixed-size values** (`uint64`, `*T`, any fixed `T`): the remainder and an array of `T`;
  with pointers, an object typed for the GC (one Go type per size class).
- **Set semantics.** Adding a value checks that it is not there: the cost by number of values,
  and from which number a value overflow is cheaper than a linear check.
- **Size classes and growth.** 32, 64, 128, 256, 384 and 512 bytes (decided 2026-10-04 from the
  model of step3-skmv-sizes.md: 83.6 % of the `street` keys fit 32 bytes, memory per key 45 instead of
  133); when a page grows, shrinks (the sketch: only if that saves 50 %), and what happens beyond 512.
- **Value overflow** (sketch 2b): a pointer to a value set when the values do not fit. First with
  the existing value set (`internal/vset`: array, then hash set); whether its inline stage is still
  needed.
- **Oversized object** (sketch 2c/2d, decided 2026-10-03): a remainder that does not fit 512 bytes
  stays inline in a larger object, exempt from R1; no pointer to a key.
- **The end page** of a byte node is a single-key page.
- **How the generic tree picks the layout for `T`** (variable-length for `string`, fixed-size
  otherwise) without a second copy of the tree (the experiment `internal/artstr` was a full copy;
  that is not to be repeated).
- **The prediction:** memory per key of the four cases on `street` and `dirs` from a model of the
  real entries (key remainder lengths, value counts and lengths), against the leaves of
  `node-layout` and the competitors; rounds per operation.

### 3.2 The single-key page standalone

Its own package (as `vpage` and `lpage` were), 100 % coverage, fuzzed, with microbenchmarks on the
real entries of `street` and `dirs` against the leaf kinds of today (flat, typed, set leaf with its
value set): lookup of the key, `ValuesFor`, add and remove of a value. Short report: memory and
cost per operation against the prediction. **Stop for the user.**

### 3.3 In the tree: `string -> {string}`

The single-key page replaces the leaves for string values; multi-key pages and the fall back off.
Measure against `node-layout` and the competitors on `street` and `dirs`, multi and unique, three
sizes, with memory and the object statistic; plus one run of the fixed-size layout with
`T = string` (see above). Report against the prediction. **Stop for the user.**

### 3.4 Value overflow (decided 2026-10-04: the set leaf holds a `Set3` on its own; no new hash set, see step3-overflow-design.md section 0)

On the real data, 1.5 to 2.5 % of the keys hold about half of the values (step3-tree-pages.md), and
their value sets were nearly half of the heap in the experiment. Design note first: the existing
value set against value blocks of 128-byte multiples (STRATEGY R6), with the prediction. Then build,
measure as in 3.3. **Stop for the user.**

**Done 2026-10-04**, measured: [step3-overflow-results.md](step3-overflow-results.md) (memory as predicted, speed neutral on the
existing benchmarks: the overflow keys are 0.5 to 1.6 % of the keys). The set leaf of a string map is on the grid 32/64/128/256.

### 3.4b Remainder of up to 505 bytes (user, 2026-10-04; [step3-klen9-design.md](step3-klen9-design.md); **done 2026-10-04**, not yet measured)

One more bit of the kind byte extends `klen` to nine bits, so a leaf and a page hold a remainder of up to 511 bytes inline (now 254; `longKey` = 255 marks
a key as a string). Then the set leaf gets the classes 384 and 512 (key area 370 and 498) and a page a longer remainder. It touches every `klen`
and every comparison of `kind` (`isLeaf`, `isPage`, the descent), so: note with the list of places, the cost in `matches` and `find`, a prediction
(memory: none on the existing corpora, their remainders are short; speed: not measurable) and then the change, before the next measurement.

### 3.5 The other cases

**3.5.1 done 2026-10-05** ([step3-fixed-results.md](step3-fixed-results.md)); decided: no 24-byte class for now, the string page shrinks with the same hysteresis as `Fixed`. A 24-byte class (the typed leaf's size for one pointer) is a candidate for the profile after 0.8.

**Decided 2026-10-04 (user), see [step3-fixed-design.md](step3-fixed-design.md):** the multi-key pages are off and cut out of `internal/art` first (tag `before-mk-pages-removal`; they are redesigned in step 4 with what is learned by then, or the old code comes back from the tag); pointer pages for one-word `T` only; values aligned to `T` and compared as `T`.

**First, the vocabulary of what survives** (user's remark of 2026-10-04: new code still said "leaf", and the code's `isPage` means the multi-key page while the single-key page
sits under `leaf` names). When the flat and the typed leaf are replaced by single-key pages (below), their code goes, so rename only what stays: `leafHead`
(the head of a single-key page), `isLeaf`/`isPage` (to `isSingleKey`/`isMultiKey`), `kSet` and the "set leaf" (the single-key page in value overflow: `kValueOverflow`, `isValueOverflow()`, in prose "value overflow", never "overflow page" or a bare "set"; user, 2026-10-04: unambiguous, no guessing; `isSingleKey`/`isMultiKey` agreed). Likewise the word `kind` (the type byte of an object) is to be called *type* or *class*, as *child* is the parent-child relation only: the field and the Go type `kind` become `objType` (agreed 2026-10-04: byte 0 of every object, nodes and pages alike; `pageType` and `nodeType` would be wrong for half of them; only the prefix of the constants `kN5`, `kR8`, `kValueOverflow` is still open); the bench profile `natural` is called "real" in prose (`street real`) and may get that name in code too, if the user says so, `skleaf.go`,
the comments and test names of the string map. Names to the user first, then one commit with nothing else in it.

`string -> {uint64}`, `string -> {*T}`, `uint64 -> {*T}`, each measured as in 3.3 (`u64` keys at 4K,
16K, 256K and a 1M spot check on the PC).

**3.5 measured 2026-10-05 on the PC** ([step3-tree-results.md](step3-tree-results.md)): predictions hold; gate 3 not met for memory of `*T` with one value per key and for credo 1 with one value per key (both known causes, step 4); M1 not available.

### Gate 3 (decided 2026-10-04: credo 1 hard, the rest reported)

- **Hard:** credo 1 for the four cases, multi and unique, every size measured: `Ordered` beats
  `btree-sets` (and `btree-map` for one value per key) in point lookups, `churn` and `build`, and
  `hashed` in range scans.
- **Hard:** memory per key not above `node-layout`, with the bytes of string values counted for
  every candidate (the "fair" column of step3-tree-pages.md).
- **Reported, not gating:** every cell against `node-layout`. A cell below 0.70 is explained in
  the report with its cause.
- Objects: 100 % at multiples of 128, except oversized objects and the value sets of the value
  overflow, which are listed on their own.
- Flat, typed and set leaves are gone from the code, or the report says why one stays.

## Step 4: the multi-key page (MKSV) as the special case

**Design approved 2026-10-05** ([step4-mksv-design.md](step4-mksv-design.md): burst under byte nodes, no range nodes now; layout B with a 3-byte head and one grid; one `build` function for burst and promote, no fall back). **4.1 done 2026-10-05** ([step4-mkpage-results.md](step4-mkpage-results.md)): `internal/mkpage`, memory equal to the model to the byte. **4.2 built and measured 2026-10-05** ([step4-tree-results.md](step4-tree-results.md)): memory as predicted, speed of changes and ranges below credo 1 and 2; options for the user in the report.

Many keys with one value each in one page, the object of steps 1 and 2. It comes back on top of
the single-key page of step 3.

- Choose the layout (A, `vpage`, or B, `lpage`) by a measurement on `street` and `dirs`, with what
  step3-layout.md and step3-tree-pages.md found as input (the header and class runs of
  `lpage` there are input, not decisions).
- The interplay with the single-key page: an entry that gets a second value leaves its multi-key
  page (promote); a subtree crowded with multi-value entries falls back. On the natural mix the fall
  back took every multi-key page away (step3-tree-pages.md); this step shows where multi-key pages
  pay (one value per key, `u64` keys) and that the natural mix keeps at least the speed of step 3.
- Design note with prediction first, as in step 3.

**Gate 4:** gate 3 again, plus credo 2 for one value per key (ranges at least `btree-map`, memory
at most `btree-map`), and no cell of the natural mix below step 3 beyond noise.

## Step 5: multi-value entries in multi-key pages (MKMV), last

**Design note approved 2026-10-06** (cpl at byte 1 as in the single-key head; shared head: 5.5). **5.1 built 2026-10-06** ([step5-mkpage-results.md](step5-mkpage-results.md)): `internal/mkpage` with continuations, 100 %, race, fuzz; page-level Get +3 to 12 %, `Each` +36 %, then option B (values of `Fixed` at the end of the object): Get and changes as in step 4, `Each` +15 %. The design: a further value is an entry of the length list with the byte 255 (no new object type, single-value pages unchanged), no limit on values per entry, the pointer pages (4.3) built in this step. Prediction: memory of the real mix -15 to -27 % and no decay by use. **5.2 built 2026-10-06** ([step5-tree-results.md](step5-tree-results.md)): the tree puts values into pages; keys stay in pages (88 %), fresh memory as predicted, `single-value` identical; the decay after use is 1.18 to 1.24 (predicted 1.05) because of the page fill (shrink hysteresis): open question for the user.

**2026-10-06:** 5.3 built (pointer pages, the stage where every change of n is a new object), gate 5 on the PC run ([step5-gate-results.md](step5-gate-results.md): memory as predicted; maps of pointers below `btree-sets` up to 16,384 keys; a crash found and fixed; the single-value slowdown analysed: the reflection check, the `pager` interface and the conversion fixed), the M1 deferred to the end of step 5. **Then the user asked for one page layout instead of four: [step5-one-page.md](step5-one-page.md); its section 5 is the plan from now on (5.5a model, 5.5b package `internal/page` and the placement benchmark, 5.5c the tree in four steps, 5.5d gate).**

Decided by the user on 2026-10-04 to be the last of the three: it makes the code much more complex.
Starting point: the parked branch `mkmv-experiment` (`internal/artstr` with `Map.Pairs`, layout B with
entries that repeat the key before them) and its measurements in step3-tree-pages.md. If it works,
the promote and the fall back go. Design note with prediction first.

## Open after step 5 (credo gaps and tuning, entered 2026-10-07 from gate m46)

Not part of a step yet; each gets a design note with a prediction before any code, and is decided after the report of gate 5 (5.5d).

- **Strings against `btree-sets` at 4,096 keys.** Churn and build of the real mix (street, dirs) are 0.89 to 0.96 times `btree-sets` in m46 (credo 1 asks for at least 1.0; m44: 0.89 to 0.97). The string pages gained nothing in churn since m44 and lost 6 to 10 % in build; above 16,384 keys the credo holds (1.04 to 1.26). To find out: where the string page's `Add` and `Remove` lose against the `uint64` page (value lengths as a list of their own, the values moved at the end of the object), with the probe's profile on `street natural 4096` strings.
- **Single value a key against `btree-map` at 4,096 and 16,384 keys.** In m46 the ordered map is 0.65 to 0.98 times `btree-map` for `uint64` street and dirs and for strings (as in m43; m46: `uint64` street 0.75 / 0.83 at 4,096, 0.98 / 0.96 at 16,384; dirs 0.70 / 0.70 and 0.86 / 0.80; strings 0.65 to 0.83), and above it only for the full corpus (1.06 to 1.18, strings dirs 0.89). A small map of one value a key is the case in which a B-tree is a single array of cache lines; to find out: how much of the gap is the descent through byte nodes (three or four objects) and how much the page, with the `Each` and the `Add` profile of a single-value tree of 4,096 keys.
- **Hysteresis between burst and merge: done 2026-10-07** (review E6, adopted): merges only into a page of at most 256 bytes.
- **The fast path of the one-key page** (see below): +13 % build and +11 % replay for `uint64` single-value in the probe (5.5c), +1 to +3 % build in the gate; belongs to the two points above.
- **Reads** were measured late (5.5e); the probes of later steps include `readprobe_test.go`.

## Feature: autotune of the keys per page (user, 2026-10-07; from the review, review-2026-10.md) — measured and dropped 2026-10-10

**Decision (user, 2026-10-10): B, always the full page; the autotune is not built.** Experiment exp-maxkeys (autotune-design.md,
result and decision): limits of 2 and 4 keys a page never win; one key a page wins point queries and writes up to 16K keys by 5 to 45 %
but loses ranges three to four times (below `main` and far below `btree-map`) and doubles the memory. The gap of small maps to `main`
goes to the search in the page and the write path of the tree ("Open after step 5").

**Idea (the user's):** the hybrid of E1 as an automatic feature. The map keeps count of its pages (and keys) and infers its memory from
them; with that count it raises, step by step and by a suitable heuristic, the largest number of keys a page may hold. A small map,
which fits the caches, keeps one key (or few keys) per page and runs the fast form (E1: as fast as step 3.5, churn up to 45 % faster
than today at 4,096 keys); as it grows, pages may take more keys, so that pages filled from then on are fuller and the memory stays low
(E1: one key per page costs about twice the memory). The two ends work alone (E1 showed it for one key per page; today's code is the
other end), and the border moves between them by one number.

**What the experiments already say about it:**
- the limit 1 costs, besides memory, 1.3 to 2.1 times the time of a scan of single-value maps (E4) and 1.3 to 3 times the reserved
  memory over a long churn (E3); in a small map that is a small absolute amount, which is the point of the idea;
- the hysteresis of E6 belongs to it either way (fewer bursts and merges for every limit above 1);
- the step at which the limit rises should follow memory, not keys alone: the bytes a key take differ by kind by a factor of three
  (E4: 27 to 114 block bytes a key).

**Concept fixed by the user (2026-10-07), implemented later:** [autotune-design.md](autotune-design.md). One number for the whole map
(`maxKeys`), growing linearly with the number of pages; nothing merged or split artificially (existing pages fill as keys come); the
feature can be switched off and `maxKeys` fixed, to find the curve in a concrete context; the curve itself comes from experiments.

## Step 6: the routing layer

- **Remove the smallest byte node** (N5, 64 bytes). The smallest is then N12 with 128 bytes. N5
  stays only where R2 allows it (anomalous chains), and only if measuring shows it is needed there.
- **Remove the tails of common prefixes in nodes.**
  - Up to 12 bytes stay in the header.
  - Longer common prefixes go into the page's common prefix.
  - Long chains above byte nodes get a path node (128 bytes, R2 allows 64).
- **A range node for few children.** An R8 node (128 bytes) held 2.6 to 3.1 ranges and cost 11 to
  16 % of the objects in the experiment (step3-tree-pages.md).
- **Single slots in the range node** (user's idea, 2026-10-03): check exact byte values first, then
  the ranges, so that an entry with a page of its own does not cut its neighbours' page in two.
  Measure how often that case occurs before building it.
- Whether byte nodes are still needed once pages hold every key.

**Gate 6:** object statistic: no object outside R1-R3; no regression beyond noise against step 5,
`url` and `path` in particular.

## Step 7: clean-up and release measurement

- Remove dead code, so that one leaf-free structure remains.
- Update the package documentation.
- Release suite on Windows: all key kinds, 4K-1M, serial and parallel, against `main`,
  `node-layout` and all competitors.
- Selected arm64 jobs.
- Update `bench/README.md` and the published numbers.
- The user decides about merging into `main`.

## To check later (notes to pick up when the time comes)

- **Range and prefix scans of string maps are slower than `node-layout`** (measured 2026-10-04, commit `06244b4`,
  `street`, natural mix, PC, `ordered` as many times as fast as `baseline`; the full table is in the report of 3.3):
  `valuesBetween` 0.46, 0.51, 0.70 and `prefix` 0.49, 0.47, 0.60 at 4,096, 16,384 and 212,449 keys, point lookups
  0.66, 0.72, 0.85, `churn` 0.84 to 0.89, `build` 0.89 to 0.90. Against `btree-sets` the credo holds (1.2 to 2.0).
  *Cause (concluded, not yet measured):* the copy-out of the values, one allocation of about 28 bytes for every key a scan
  visits (`skpage.Page.Strings`); `node-layout` hands out the string headers it holds, and in the bench the bytes
  behind them lie in a small buffer (a few thousand different locality names) that stays in the cache. About 20 to 25 ns
  an allocation against about 8 ns for a leaf in the cache fits the shape: worst at 4K keys, better as the misses
  take over. *Decision of the user, 2026-10-04:* leave it as it is and go for version 0.8. *Options for later:* (1) take
  the buffers of a scan, and perhaps of a lookup, from a chunk of 512 bytes to 4 KB shared by several keys (fewer
  allocations; a string that is kept holds its chunk); (2) a scan that gives out views of the page, with a
  string that is kept holding its page (the user excluded zero-copy as the default); (3) the bench with the strings
  scattered over the heap, as an application would have them, to see how much of the gap is the bench. *How to check:* a
  profile of `valuesBetween` (allocation and GC share), the same run with chunked buffers as an experiment, and the
  corpus of the inverted index (many values per key, so that the copy is amortized or not).

## Version 0.8 and what follows (decided 2026-10-04)

The goal of the current work is a **complete implementation, called version 0.8**, on the benchmarks that exist
(`street`, `dirs`, and the synthetic key kinds for the checks): the single-key page (step 3), the multi-key page
(step 4) and multi-value entries in multi-key pages (step 5), each with its gate. Only then:

1. **Measure and profile at large**, with the corpora of the backlog below, on the PC and the M1, against all
   competitors.
2. **Fix the optimization strategies** from what that shows (mutation, the three-byte header, the routing layer
   of step 6, the value overflow, zero-copy as an option), each with its prediction, as the working rules say.

No new corpus is built before that point; the backlog says what to build then. (Open: whether step 6, the
routing layer, belongs to 0.8 or to the optimization that follows it.)

## Backlog: further benchmark corpora (todo for after 0.8, not to be started now)

Each answers a question the current corpora cannot. Raw data does not go into the repository: a download script
with the version of the source, a loader like `keys/corpora.go`, a statistic of the value distribution, and the
license checked before the first commit.

- **Inverted index (term -> documents) from a Wikipedia dump. A must.** The classic multimap use, a Zipf
  distribution: few terms with millions of values, most with one. It tests the value overflow (step 3.4) and
  integer values (documents as numbers), and gives the largest value sets. Terms are strings, values `uint64`
  (document numbers) and, as a second profile, strings (titles).
- **DBLP (author -> publications), CC0.** Strings as keys and values, one publication for most authors, hundreds
  for a few: `street` with a realistic tail. The smallest of the heavy-tailed corpora.
- **Wikipedia pagelinks (page number -> linked page numbers).** `uint64` -> `uint64` with a power law: the
  integer case with a real distribution.
- **The complete Debian file tree** (`Contents-*` of several releases; path -> packages): `dirs` without
  the reduction, past one million keys, for the 1M spot check. Possibly with a skewed choice of the keys.
- **DNS records of the Tranco top list** (idea of the user). Keys: the domain names of one *numbered* Tranco list
  (the list id makes it reproducible). Values: the records that are queried (A, AAAA, HTTPS, CNAME, MX, NS, TXT, ...),
  with their types. Only the **distribution** matters: the number of records of each type per domain and their
  lengths are measured once and stored as a histogram in the repository; the benchmark draws the values as
  random bytes of those lengths from a seeded generator. That is reproducible without storing the records and
  raises no license question (the data is meant to be cached). A real DNS cache would be hash-based, since DNS has no range
  queries; here the corpus stands for variable-length values with a realistic spread, and for the case that a hash map
  no longer fits the memory.
- **After those, for robustness and not for decisions:** the French address base (BAN, Licence Ouverte;
  street -> municipality, ten times `street`), GeoNames (alternative name -> places, CC BY, Unicode, "Springfield"),
  US addresses (TIGER, public domain; OpenAddresses by source), a package registry (package -> dependents,
  Debian `Packages`, crates.io).
- **Not now, expensive:** the Common Crawl URL index (URL -> captures) for real URLs instead of the synthetic
  `url` kind, as a sample; OpenStreetMap cells -> places for range queries.

The first two (inverted index, DBLP) are due before the value overflow of step 3.4 is designed *if* the
real data of 0.8 is not enough for it: decide at 3.4 whether to pull them forward.

## Status

The state of work, the next action and open questions live in [STATUS.md](STATUS.md). Update it
at every gate and whenever work stops.
