# Status

## 2026-10-02: plan written, nothing implemented

**Parked and pushed:**

| branch | commit | content |
|---|---|---|
| `node-layout` | `7b8a8d8` | pessimistic paths, leaves with key remainders, flat and typed leaves, results against `main` and the competitors |
| `node-pages` | last commit | pages for integer-like keys with one value, range nodes, results in `bench/results-layout/node-pages/` |
| `leaf-pages` | `b0aeed1` | the earlier page experiments (K, S, U8-n pages) and their measurements |

**Working branch:** `cacheline`, forked from `node-pages`. It holds these documents.

**Next action:** step 0 of [PLAN.md](PLAN.md), after the user's go-ahead.

**Open points:**

- **Test run time.** `go test ./... -race` fails on the default timeout (see below). PLAN step
  0.4.
- **arm64 machine.** The user owns one and starts jobs by hand. Model and OS are not yet known:
  ask when building the runner.
## Test run time (2026-10-02)

`node-pages`, WSL, machine otherwise idle:

| run | `internal/art` | `TestAgainstReference` |
|---|---|---|
| `go test -race -cover ./...` (default 10-minute timeout) | failed: timeout | still running at 600 s |
| `go test -race -cover -timeout 60m -v ./internal/art` | passed, 100% coverage, 1175 s | 1121 s |
| `go test -run TestAgainstReference ./internal/art` (no race detector) | | 43.5 s |

All other tests together take about 55 s with the race detector; the slowest are
`TestRemoveAbsentNearMisses` (25 s) and `TestLongPaths` (7 s).

`TestAgainstReference` runs five leaf modes over every key set, against the reference model.
Each run has six phases, and after each phase it compares the whole map and checks the
invariants of the whole tree. The race detector slows it down about 26 times. Step 0.4:
cut its work under `-race`, for example fewer rounds when `testing.Short()` holds or behind a
`race` build tag, or check the invariants less often, without losing coverage. Then measure
the run time again.
