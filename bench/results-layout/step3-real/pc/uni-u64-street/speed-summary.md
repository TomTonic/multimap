| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique | street | 4096 | valuesFor | ordered | btree-map | 12 | 48.8 | 92.0 | 1.89× [1.88, 1.89] | +47.0% | [+46.8%, +47.2%] | 0.4 pts | 1.8 | yes | yes |
| unique | street | 4096 | valuesBetween | ordered | btree-map | 12 | 405 | 498 | 1.23× [1.22, 1.24] | +18.8% | [+17.9%, +19.7%] | 0.9 pts | 1.3 | yes | yes |
| unique | street | 4096 | prefix | ordered | btree-map | 12 | 103 | 131 | 1.27× [1.26, 1.28] | +21.2% | [+20.5%, +21.8%] | 0.7 pts | 1.6 | yes | yes |
| unique | street | 4096 | churn | ordered | btree-map | 12 | 104 | 138 | 1.33× [1.32, 1.34] | +24.6% | [+24.1%, +25.1%] | 0.9 pts | 1.1 | yes | yes |
| unique | street | 4096 | build | ordered | btree-map | 12 | 1.90 ms | 2.12 ms | 1.14× [1.11, 1.18] | +12.4% | [+9.9%, +15.0%] | 3.3 pts | 4.4 | no | yes |
| unique | street | 16384 | valuesFor | ordered | btree-map | 6 | 61.9 | 128 | 2.06× [2.00, 2.12] | +51.4% | [+49.9%, +52.8%] | 1.4 pts | 5.2 | yes | yes |
| unique | street | 16384 | valuesBetween | ordered | btree-map | 6 | 482 | 607 | 1.26× [1.24, 1.27] | +20.3% | [+19.3%, +21.3%] | 1.0 pts | 1.6 | yes | yes |
| unique | street | 16384 | prefix | ordered | btree-map | 6 | 179 | 277 | 1.54× [1.50, 1.59] | +35.1% | [+33.4%, +36.9%] | 1.7 pts | 4.1 | yes | yes |
| unique | street | 16384 | churn | ordered | btree-map | 6 | 134 | 197 | 1.48× [1.46, 1.51] | +32.6% | [+31.5%, +33.6%] | 1.0 pts | 2.0 | yes | yes |
| unique | street | 16384 | build | ordered | btree-map | 6 | 8.93 ms | 11.10 ms | 1.24× [1.22, 1.26] | +19.3% | [+17.8%, +20.8%] | 1.4 pts | 1.2 | yes | yes |
| unique | street | 212449 | valuesFor | ordered | btree-map | 12 | 140 | 238 | 1.70× [1.63, 1.78] | +41.2% | [+38.8%, +43.7%] | 3.4 pts | 3.5 | yes | yes |
| unique | street | 212449 | valuesBetween | ordered | btree-map | 12 | 758 | 1044 | 1.37× [1.32, 1.43] | +27.1% | [+24.3%, +29.9%] | 3.3 pts | 1.5 | no | yes |
| unique | street | 212449 | prefix | ordered | btree-map | 12 | 1740 | 3042 | 1.75× [1.72, 1.78] | +43.0% | [+42.0%, +44.0%] | 1.1 pts | 1.0 | yes | yes |
| unique | street | 212449 | churn | ordered | btree-map | 12 | 398 | 488 | 1.24× [1.22, 1.26] | +19.3% | [+18.2%, +20.4%] | 2.5 pts | 2.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique street n=4096 build: ordered vs btree-map: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 valuesFor: ordered vs btree-map: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=16384 prefix: ordered vs btree-map: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=212449 valuesFor: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=212449 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.03% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=212449 churn: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
