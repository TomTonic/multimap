# Smoke run of the corpus `links`

A check that the new key kind `links` (see `bench/keys/testdata/README.md`) runs through the bench in all its
parts. It answers no question about the library and draws no conclusion; the tables are as rtcompare wrote them.

- When and where: 2026-10-10, 09:28 to 09:29, the PC (Ryzen 9 7900) through WSL, nothing else running; 26 s.
- Commit: `6e3e9c9` (branch `corpus-links`), `uint64` values (no `strvals` build).
- Command, from `bench`:

  ```
  go run ./cmd/bench -suite dev -keys links -values natural,single-value -sizes 4096 -ops valuesFor,churn \
      -vs btree-sets,btree-map -minprocs 4 -maxprocs 4 -memn 4096 -memrounds 1 -out <dir>
  ```

- Files: [speed-summary.md](speed-summary.md), [mem-summary.md](mem-summary.md).
- Limits: 4 processes a comparison, which is below the precision the bench asks for (`churn` intervals are wider
  than ±2 points, see the warnings in the speed summary); one memory round at 4,096 keys.
