# Redesign: an ART that uses whole cache lines

Start here when you resume the work.

1. [STRATEGY.md](STRATEGY.md): the goal, the cache-line premise with its binding rules (R1-R6),
   the diagnosis of the current state, the target architecture and the lessons that bind it.
1a. [GLOSSARY.md](GLOSSARY.md): the vocabulary (page, entry, path, remainder, common prefix, ...) and the names it retires.
2. [PLAN.md](PLAN.md): the steps, each with its gate, and the working rules.
3. [STATUS.md](STATUS.md): where the work stands, the next action and the open questions.
4. [MEASURING.md](MEASURING.md): how to measure on Windows and arm64.
5. [objstat-node-pages.md](objstat-node-pages.md): the object statistic of the starting point.
   [step1-results.md](step1-results.md): the results of the page prototype (step 1);
   [step2-design.md](step2-design.md) and [step2-results.md](step2-results.md): the pages in the tree
   (step 2) and their measurements.
6. The tools: `bench/cmd/objstat` (the object statistic), `bench/cmd/pagefill` (bytes per key of
   the page prototype `internal/vpage`), `bench/remote/` (the arm64 job queue and its runner, see
   MEASURING.md).

Measured history:
- `bench/results-layout/` (`node-layout`, `node-pages`);
- `bench/results-leaf-pages/` on branch `leaf-pages`.
