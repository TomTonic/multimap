| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | street | 4096 | valuesFor | ordered | baseline | 8 | 106 | 82.2 | 0.77× [0.77, 0.78] | -29.7% | [-30.4%, -28.9%] | 0.9 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 107 | 140 | 1.30× [1.28, 1.31] | +23.0% | [+22.1%, +24.0%] | 1.2 pts | 1.6 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 5074 | 5285 | 1.05× [1.04, 1.06] | +4.6% | [+3.5%, +5.7%] | 1.3 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 5058 | 7127 | 1.41× [1.38, 1.44] | +29.2% | [+27.7%, +30.7%] | 1.8 pts | 2.6 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 8 | 552 | 484 | 0.88× [0.87, 0.89] | -13.8% | [-15.6%, -12.1%] | 1.9 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 8 | 555 | 728 | 1.31× [1.29, 1.33] | +23.6% | [+22.6%, +24.6%] | 2.2 pts | 2.6 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 8 | 230 | 127 | 0.55× [0.54, 0.56] | -81.9% | [-84.6%, -79.3%] | 2.5 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 8 | 234 | 209 | 0.89× [0.89, 0.90] | -12.1% | [-12.9%, -11.3%] | 1.0 pts | 0.7 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 8 | 8.81 ms | 5.12 ms | 0.58× [0.58, 0.59] | -72.1% | [-73.3%, -70.9%] | 2.1 pts | 0.9 | yes | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 8 | 8.84 ms | 8.37 ms | 0.96× [0.93, 0.99] | -4.5% | [-7.5%, -1.4%] | 3.2 pts | 2.6 | no | yes |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 8 | 124 | 107 | 0.86× [0.85, 0.88] | -16.0% | [-18.3%, -13.6%] | 2.5 pts | 2.5 | no | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 8 | 130 | 194 | 1.50× [1.47, 1.54] | +33.5% | [+32.1%, +34.9%] | 1.9 pts | 2.3 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 5628 | 5993 | 1.07× [1.05, 1.10] | +6.6% | [+4.4%, +8.9%] | 2.2 pts | 3.4 | no | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 5748 | 8454 | 1.51× [1.47, 1.54] | +33.7% | [+32.1%, +35.2%] | 3.8 pts | 7.0 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 8 | 1901 | 1985 | 1.05× [1.03, 1.06] | +4.3% | [+3.3%, +5.4%] | 1.6 pts | 0.9 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 8 | 1884 | 3061 | 1.63× [1.61, 1.64] | +38.5% | [+38.1%, +39.0%] | 1.0 pts | 1.4 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 8 | 251 | 182 | 0.72× [0.70, 0.75] | -38.4% | [-42.7%, -34.1%] | 4.3 pts | 2.5 | no | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 8 | 261 | 309 | 1.16× [1.14, 1.18] | +14.0% | [+12.4%, +15.5%] | 3.5 pts | 3.4 | yes | yes |
| natural-str | street | 16384 | build | ordered | baseline | 8 | 40.23 ms | 28.48 ms | 0.70× [0.68, 0.72] | -43.6% | [-47.6%, -39.5%] | 4.3 pts | 2.6 | yes | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 8 | 41.35 ms | 48.91 ms | 1.16× [1.13, 1.20] | +13.9% | [+11.2%, +16.6%] | 4.0 pts | 3.4 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | baseline | 8 | 207 | 234 | 1.13× [1.08, 1.17] | +11.2% | [+7.6%, +14.9%] | 4.1 pts | 4.4 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | btree-sets | 8 | 259 | 526 | 2.03× [1.94, 2.13] | +50.7% | [+48.5%, +53.0%] | 2.5 pts | 3.5 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 7394 | 9430 | 1.28× [1.24, 1.32] | +22.1% | [+19.6%, +24.5%] | 2.6 pts | 3.1 | no | yes |
| natural-str | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 7788 | 22.2 µs | 2.85× [2.73, 2.97] | +64.9% | [+63.4%, +66.3%] | 1.5 pts | 4.7 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | baseline | 8 | 28.0 µs | 36.2 µs | 1.30× [1.26, 1.35] | +23.3% | [+20.6%, +25.9%] | 2.7 pts | 4.4 | no | yes |
| natural-str | street | 212449 | prefix | ordered | btree-sets | 8 | 27.3 µs | 75.0 µs | 2.75× [2.64, 2.87] | +63.7% | [+62.1%, +65.2%] | 1.8 pts | 5.0 | yes | yes |
| natural-str | street | 212449 | churn | ordered | baseline | 8 | 535 | 529 | 0.99× [0.96, 1.01] | -1.5% | [-4.5%, +1.5%] | 3.3 pts | 1.2 | no | no |
| natural-str | street | 212449 | churn | ordered | btree-sets | 8 | 619 | 770 | 1.26× [1.24, 1.28] | +20.7% | [+19.2%, +22.1%] | 1.8 pts | 2.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str street n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 prefix: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 build: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 7.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 churn: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 build: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 build: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs btree-sets: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 churn: ordered vs baseline: the pooled difference of -1.51% does not clear the 1.79% noise floor, the bound on what the harness reports between identical code in every process
- natural-str street n=212449 churn: ordered vs baseline: the pooled interval [-4.49%, 1.48%] includes zero
- natural-str street n=212449 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
