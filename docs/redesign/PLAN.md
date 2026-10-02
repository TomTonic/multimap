# Plan: pages for every key

This plan turns [STRATEGY.md](STRATEGY.md) into code, step by step. Each step ends at a gate: a
measurement against fixed criteria, a short report to the user, and the user's decision.
Nothing moves to the next step without that decision.

Branch: `cacheline`, forked from `node-pages` (`node-pages` and `node-layout` are parked and
pushed). Worktree: `/mnt/c/temp/code/multi_map-layout`. `main` stays untouched until the user
decides to merge.

## Working rules

- **Language.** Talk to the user in German. Code, comments, commit messages and documents in the
  repository are English.
- **Commits.** Commit only after the user approved. Push only when the user asks (job queue
  pushes for arm64 included). Never run destructive git commands.
- **Commit messages.** Follow AGENTS.md: imperative mood, subject at most 72 characters, a
  blank line, a body that says why. End with the `Co-Authored-By` trailer of the model that
  did the work.
- **Code checks.** Every commit keeps:
  - 100% coverage per package;
  - `go test ./... -race` passing;
  - `golangci-lint run` clean;
  - the fuzz test running 60 s without a finding (`go test ./internal/art -run '^$' -fuzz FuzzOperations -fuzztime 60s`).
- **Test docs.** Test documentation reads outside-in (AGENTS.md).
- **Measuring.** See [MEASURING.md](MEASURING.md): claims only from interleaved rtcompare runs,
  durations with clock time, a quiet machine during runs.
- **Reports.** Report faithfully. Numbers that speak against a step go into the report and the
  results README just like the good ones. A gate that is missed is reported as missed, with the
  options. It is not explained away.
- **Backups.** Keep backups of uncommitted work outside WSL's `/tmp` (for example
  `C:\temp\bench-win`). Never `git checkout` a file with uncommitted edits.
- **References.** `node-layout` (`7b8a8d8`) is the reference for credo 3 (regression). The
  last step's commit is the reference for diagnosis (`-vs baseline`). The competitors are the
  reference for credo 1 and 2.

## Step 0: groundwork (done 2026-10-02, waiting for the gate)

1. **Object statistic as a tool.** `art.Map.Objects` (`internal/art/objects.go`) walks every
   object of a tree and reports label, size, whether it holds pointers, and the keys it holds;
   `art.Block` turns a size into the block the Go allocator takes. Tests compare both with the
   runtime (`TestBlock`, `TestObjectSizes`) and check that the objects account for every key
   (`TestObjects`). `bench/cmd/objstat` prints the table of
   [objstat-node-pages.md](objstat-node-pages.md) for every bench case, at the corpus maximum for
   `path` and `street`. **Every new object kind of the redesign needs a case in `object()`**;
   `TestObjects` and the tool's check that the objects hold all keys fail without one.
2. **Probe order of `valuesFor`.** `keys.Corpus.Probes` is the hits for corpora of 262 144 keys
   and more, else 262 144 random picks of the keys. `valuesFor` uses it. Documented in
   `bench/README.md`: `valuesFor` numbers at 4K and 16K from before are not comparable.
3. **arm64 job queue.** `bench/remote/queue.txt` and `bench/remote/arm-run.sh`, see
   MEASURING.md. Tested with a dry-run test (`bench/remote`) and once end to end on Linux
   against a local bare repository, with and without a baseline. Not yet tried on macOS.
4. **Test run time.** Under the race detector the key sets of `TestAgainstReference` are a tenth
   of their size and the 64K-path set runs in two leaf modes only (`race_on_test.go`,
   `race_off_test.go`). `go test ./... -race` takes 3.5 minutes instead of more than 20, with
   100% coverage in both modes.

**Gate 0:**
- The user has reviewed the tool, the probe order and the runner.
- The object statistic of `node-pages` is reproduced and committed as the starting point (done:
  `objstat-node-pages.md`, produced by the tool).
- The arm64 A/A run (`node-pages` against itself, u64 and str, 4K and 16K) is between 0.97 and
  1.03, or its deviation is explained. This needs the tools committed and pushed first, and the
  user starting the job.

## Step 1: prototype the page for variable keys (isolated; done 2026-10-02, see step1-results.md)

Build the page of STRATEGY 4.1 as a self-contained unit, not yet wired into the tree (built as its
own package, `internal/vpage`, so that it stays out of the tree's coverage until step 2).
- Operations: build from sorted items, search, insert, delete, split, merge, iterate from a
  bound.
- One value per key only (step 3 adds more).

**Inputs.** Real suffixes: take the keys of each bench corpus, cut off what a tree of
`node-pages` routes in nodes (the depth of each leaf), and fill pages of 128, 256 and 512 bytes
in random insertion order.

**Questions this step must answer** (diagnostic microbenchmarks are fine here, this is a
prototype):

| question | measure |
|---|---|
| bytes per key by key kind, at the fill random insertion leaves | must not exceed today's leaf plus its child pointer (`node-pages` objstat and memory figures) |
| head format: first 8 suffix bytes vs 7 bytes + length | compares per lookup, misses per lookup |
| tail compare: rounds per lookup | R5: at most two rounds (directory, then everything else together) |
| page classes and split rule (by count or by bytes) | fill after random inserts, bytes per key |
| insert and delete cost (memmove, heap compaction) | ns per operation against appending to a leaf |
| bloom filter: keep or drop | misses answered from the head |

**Gate 1:** report to the user with numbers per key kind. Decide on a layout. Continue only if
memory per key does not exceed today's leaves for any key kind and a lookup inside a page takes at
most two rounds of cache-line loads (R5; a round is defined in STRATEGY.md section 2). Otherwise
stop and discuss. (Decided 2026-10-02: the criterion was "at most two lines after the head"; step 1
showed that rounds, not lines, are what a cold lookup pays for, and the user restated it as rounds.)

## Step 2: pages for keys with one value, every key kind

Wire the page of step 1 into the tree. It replaces the U8 pages.
- Every key with exactly one value lives in a page below range nodes, whatever its length.
- Keys with several values keep their leaves for now.
- Keep the fallback (`settle`/`crowded`): a subtree whose keys mostly hold several values
  becomes inner nodes and leaves.

**Measure** (Windows, dev suite):
- unique against `node-pages`, all key kinds, with memory;
- unique against `btree-map`;
- multi against `node-pages` (must stay neutral);
- the object statistic of every case.

**Gate 2:**
- **Lookup gate:** point lookups (`valuesFor`) of string keys at 16K-64K keys at least 0.85 of
  `node-layout`, else pages only for short suffixes. (Step 1 measured pages alone; this is the first
  time they sit below a tree.)
- Credo 1 and 2 for every key kind with one value per key. Ranges at least `btree-map`, memory
  at most `btree-map`. That closes the str, uuid, email, url, path and street gap.
- No cell below 0.85 against `node-layout` (credo 3), 4K and 16K included.
- u64 not worse than `node-pages` beyond noise.
- Objects of unique cases: at least 95% at multiples of 128 bytes.

## Step 3: several values per key

- Values of a key sit inline in its page entry, up to a threshold found by measuring (start at
  4).
- Beyond the threshold: a value object of 128- or 256-byte blocks. Measure a B-tree of value
  blocks against `vset`.
- A page kind with a pointer array for those objects; pages without them stay pointer-free.
- Remove flat and set leaves for maps whose values are small and pointer-free. Remove the
  fallback if it is no longer needed. Decide that by measuring, not by taste.
- Look at multi `churn` at 1M u64 (0.82 since `node-layout`) with the new structure.

**Gate 3:**
- Multi against `node-layout` (all key kinds, 4K-256K, 1M spot check): no cell below 0.85.
- Against `hashed`, `btree-sets` and `map-sets`: credo 1.
- Memory not above `node-layout`.
- Objects: 100% at multiples of 64, at least 95% at multiples of 128 (value objects included).

## Step 4: the routing layer

- **Remove N5.** The smallest inner node is N12 with 128 bytes. N5 stays only where R2 allows it
  (anomalous chains), and only if measuring shows it is needed there.
- **Remove path tails.**
  - Up to 12 bytes stay in the header.
  - Longer shared suffix bytes go into the page's shared-prefix field.
  - Long chains above inner nodes get a path node (128 bytes, R2 allows 64).
- Check whether range nodes or inner nodes should route where, now that pages hold every key.

**Gate 4:**
- Object statistic: no object outside R1-R3.
- No regression beyond noise against step 3, url and path in particular.

## Step 5: values with pointers

- Typed pages for values with pointers (string values, `strvals`): generic, sized to 256 or 512
  bytes.
- Remove typed leaves.

**Gate 5:**
- The `strvals` suite against `node-layout` with string values: no cell below 0.85.
- Memory not above `node-layout`.
- GC CPU per cycle not above `node-layout`.

## Step 6: clean-up and release measurement

- Remove dead code, so that one leaf-free structure remains.
- Update the package documentation in `node.go`, `page.go` and `rnode.go`.
- Release suite on Windows: all key kinds, 4K-1M, serial and parallel, against `main`,
  `node-layout` and all competitors.
- Selected arm64 jobs.
- Update `bench/README.md` and the published numbers.
- The user decides about merging into `main`.

## Status

The state of work, the next action and open questions live in [STATUS.md](STATUS.md). Update it
at every gate and whenever work stops.
