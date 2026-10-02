# Step 2: pages for every key (2026-10-02)

The tree with the pages of `internal/vpage` for every key with one value (`cacheline`, see
`docs/redesign/step2-design.md`) against `node-pages` (`cf6944b`) and `btree-map`. Serial regime, native
Windows, Ryzen 9 7900, rtcompare v0.8.0, 4K, 16K and 256K keys. The report with tables and the reading
against gate 2 is `docs/redesign/step2-results.md`.

| folder | what |
|---|---|
| `unique-vs-node-pages` | one value per key, `-vs baseline` (the copy of `node-pages`), speed and memory phase of the run |
| `unique-vs-btree-map` | one value per key, `-vs btree-map` |
| `multi-vs-node-pages` | several values per key (must stay neutral), 4K and 16K |
| `memory-vs-node-pages`, `memory-vs-btree-map` | memory at 256K keys, 3 rounds, after the last code change |

Every folder holds `speed-summary.md` or `mem-summary.md` (the cells with their intervals), the raw
`*.jsonl` and `run.json` (the flags).
