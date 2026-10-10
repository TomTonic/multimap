# Task: the benchmark corpus `links` (Wikipedia page links), for a fresh session

Written 2026-10-10 by the session that works on the redesign (branch `cacheline`), for a new session that builds one benchmark
corpus and nothing else. Read it completely before you start.

## Who you are working with, and the rules

- Talk to the user in **German**. Code, comments, commit messages and documents in the repository are **English**.
- The repository is a Go multimap (`multimap.Ordered`, an adaptive radix tree with pages). Read first: `AGENTS.md` (code style,
  test documentation outside-in, 100 % coverage of the library), `docs/redesign/PLAN.md` (section "Working rules"; the backlog
  entry "Wikipedia pagelinks" under "Backlog: further benchmark corpora"), `docs/redesign/MEASURING.md`, and
  `bench/keys/testdata/README.md` (how the existing real corpora `street` and `dirs` are documented).
- **Stop and ask the user** at every decision that this task does not settle, at every surprise (a number that does not match
  the expectation below, a license doubt, a file larger than planned), and before every benchmark run. Do not build an
  alternative on your own.
- Speed claims come only from rtcompare runs (`cmd/bench`), never from your own timing loops or `go test -bench`.
- One measurement at a time on the machine; WSL and the Windows PC are the same machine. Before any run, ask the user whether
  the machine is free, and tell the duration and the expected end time.
- Never kill processes by pattern (`pkill -f` kills your own shell); only by exact name or PID, and only your own.

## Where you work

- **Your own worktree and branch**: `git -C /mnt/c/temp/code/multi_map-layout worktree add -b corpus-links
  /mnt/c/temp/code/multi_map-links origin/cacheline` (after `git fetch`). Work only there.
- **Do not touch** `/mnt/c/temp/code/multi_map-layout` (the other session's worktree, branch `cacheline`) or any other branch.
- Push only your branch: `git -c credential.helper= -c "credential.helper=!gh auth git-credential" push -u origin corpus-links`.
  The other session merges it into `cacheline` later.
- Commit messages: imperative mood, subject at most 72 characters, blank line, body; end with the co-author line your
  environment prescribes. Before every commit: `go test ./... -race` (library) and `cd bench && go test ./...`, `golangci-lint
  run` 0 issues in both modules (use golangci-lint v2.14.0 built with Go 1.27.2: `GOTOOLCHAIN=go1.27.2 go install
  github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.14.0` if `golangci-lint version` shows an older Go).

## Why this corpus

The benchmark has two real corpora with natural values: `street` (street name -> localities) and `dirs` (directory -> file
names). Both have the same shape: most keys hold one value. The page-link graph of Wikipedia has the opposite shape: the typical
key holds many values. Measured on the dump below (Simple English Wikipedia, articles only, namespace 0), against the existing
corpora:

| corpus | keys with 1 value | 2-8 | 9-64 | 65+ | values in sets of 65+ | values a key (mean / median) | key length |
|---|--:|--:|--:|--:|--:|--:|--:|
| street natural | 79 % | 18 % | 2.5 % | 0.4 % | 34 % | 2.7 / 1 | 14 B |
| dirs natural | 62 % | 33 % | 4.5 % | 0.5 % | 28 % | 3.5 / 1 | 51 B |
| **page links** (page -> linked pages) | **28 %** | **13 %** | **43 %** | **17 %** | **76 %** | **47.7 / 13** | titles: to be measured |

Full numbers of the page links (all 400,659 linking pages of namespace 0, 19,110,260 links to any namespace; the 28 % with one link
are probably redirects, not checked): keys by values a key 1: 27.7 %, 2: 0.6 %, 3-4: 2.1 %, 5-8: 10.0 %, 9-16: 16.6 %, 17-64:
25.9 %, 65-256: 12.9 %, 257+: 4.2 %; values by values a key 1: 0.6 %, 2: 0.0 %, 3-4: 0.2 %, 5-8: 1.4 %, 9-16: 4.2 %, 17-64:
17.4 %, 65-256: 34.0 %, 257+: 42.3 %; max 5,932. Your corpus will differ a little (targets restricted to namespace 0, see below);
say by how much.

## The corpus

- **Source** (fixed, reproducible; do not use `latest`): `https://dumps.wikimedia.org/simplewiki/20261001/` with the files
  `simplewiki-20261001-page.sql.gz` (33 MB), `simplewiki-20261001-pagelinks.sql.gz` (84 MB),
  `simplewiki-20261001-linktarget.sql.gz` (38 MB). Read each table's `CREATE TABLE` for the columns; as of this dump
  `pagelinks` is `(pl_from, pl_from_namespace, pl_target_id)` and `pl_target_id` refers to `linktarget(lt_id, lt_namespace,
  lt_title)`. The value tuples of the `INSERT` statements stand on lines of their own; titles are quoted SQL strings with
  escapes (`\'`, `\\`, ...): parse them correctly, do not split on commas naively.
- **Keys**: the titles (as stored, with underscores, UTF-8) of the namespace-0 pages that link to at least one namespace-0 page.
  Redirect pages stay in (they are real keys with one link); count them (`page.page_is_redirect`) and report their share.
- **Natural values**: for each key, the namespace-0 titles it links to, as numbers: the index plus one of the target title in
  the sorted list of all target titles of the corpus (the same scheme as `street`'s localities), and the titles themselves
  as the names for the string-value profile (`Corpus.Names`, as `street` and `dirs` do).
- **Size: the decision is the user's.** All keys with all values are far too large to embed (19 million values). The bench needs
  at least twice its largest size in keys (`fromList` panics below that: 2 x 65,536 = 131,072 keys for the dev suite's sizes),
  and the existing corpus files are about 5 MB each. Sample **keys**, never values (a sampled key keeps all its links, so the
  distribution of values a key stays), deterministically with a fixed seed. Compute two or three options (number of keys, number
  of values, size of the gzipped file, the distribution table above for each), show them to the user and **stop until the user
  chooses**.
- **File**: `bench/keys/testdata/links.tsv.gz` in a format like `streets.tsv.gz` (document it in the README: the target titles
  first, then one line per key: title, tab, comma-separated target indexes; or what fits better, with the reason).

## License (check before the first commit of the data)

The dumps are under CC BY-SA 4.0 (and GFDL for the text); titles and links are data of the dump. Check
<https://dumps.wikimedia.org/legal.html> and the Wikimedia terms, write a section in `bench/keys/testdata/README.md` like the
others (source with date, what the file holds, the license that applies to this file only, attribution "Wikipedia contributors,
Simple English Wikipedia"). If anything is unclear, ask the user before committing the file.

## The code

Follow how `street` and `dirs` are done (search for `Street`, `streetCorpus`, `Dirs` in `bench/`):
1. `bench/cmd/mkcorpora/main.go`: a corpus `links` that downloads the three files from the fixed addresses, builds the file
   deterministically, and prints the statistics of the table above (keys, values, buckets, key length, redirect share).
2. `bench/keys`: the kind `Links Kind = "links"` with its doc comment (source, shape, sizes), in `Kinds` after `dirs`, `Text`
   true; the loader in `corpora.go` (`go:embed`, `sync.OnceValue`), `Generate` through `fromList` with the natural values and
   names; tests as the existing kinds have (`keys_test.go`), with outside-in doc comments.
3. `bench/cmd/bench`: whatever the other real kinds need there (natural values, string values, prefixes of text keys); the
   help texts that list the kinds.
4. Nothing in the library (`internal/`, the root package) changes.

## Check and report

1. The statistics of the chosen corpus against the table above (mkcorpora's output), in the README and in your report.
2. With the user's go, a short rtcompare smoke run on the PC side through WSL, e.g. `cd bench && go run ./cmd/bench -suite dev
   -keys links -values natural,single-value -sizes 4096 -ops valuesFor,churn -vs btree-sets,btree-map -minprocs 4 -maxprocs 4
   -memn 4096 -memrounds 1 -out <scratch dir>`; report the summary tables as they are, without conclusions about the library.
3. Update `docs/redesign/PLAN.md`: the backlog entry "Wikipedia pagelinks" says what was built (corpus `links`, branch, commit)
   and its numbers; add a line to `docs/redesign/STATUS.md`. Push the branch and report to the user in German: what was built,
   the corpus's numbers against the expectation, open points.
