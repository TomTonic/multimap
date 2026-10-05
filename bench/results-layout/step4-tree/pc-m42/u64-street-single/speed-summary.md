| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | street | 4096 | valuesFor | ordered | baseline | 6 | 61.3 | 49.1 | 0.80× [0.80, 0.80] | -24.8% | [-25.1%, -24.4%] | 0.3 pts | 0.4 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 6 | 61.8 | 91.7 | 1.48× [1.47, 1.50] | +32.7% | [+32.0%, +33.3%] | 0.6 pts | 1.2 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 6 | 1152 | 2076 | 1.81× [1.78, 1.84] | +44.7% | [+43.7%, +45.7%] | 1.0 pts | 1.2 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 6 | 1127 | 499 | 0.44× [0.44, 0.45] | -125.7% | [-128.5%, -122.9%] | 2.6 pts | 1.4 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 6 | 244 | 233 | 0.95× [0.95, 0.96] | -4.8% | [-5.7%, -3.8%] | 0.9 pts | 1.0 | yes | yes |
| single-value | street | 4096 | prefix | ordered | btree-map | 6 | 243 | 132 | 0.54× [0.54, 0.55] | -84.4% | [-86.2%, -82.6%] | 1.7 pts | 1.2 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 6 | 1805 | 147 | 0.08× [0.08, 0.08] | -1114.3% | [-1143.1%, -1085.5%] | 27.5 pts | 0.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 6 | 1728 | 148 | 0.08× [0.08, 0.09] | -1090.7% | [-1115.0%, -1066.3%] | 23.2 pts | 1.0 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 6 | 19.09 ms | 1.90 ms | 0.10× [0.10, 0.10] | -901.1% | [-922.2%, -880.0%] | 20.1 pts | 1.8 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 6 | 19.10 ms | 2.09 ms | 0.11× [0.11, 0.12] | -792.1% | [-824.7%, -759.6%] | 31.0 pts | 2.1 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 73.5 | 69.4 | 0.94× [0.92, 0.96] | -6.5% | [-9.3%, -3.7%] | 2.7 pts | 2.7 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 74.7 | 126 | 1.69× [1.65, 1.74] | +40.9% | [+39.3%, +42.5%] | 1.8 pts | 3.8 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1311 | 2510 | 1.90× [1.88, 1.92] | +47.5% | [+46.9%, +48.0%] | 0.5 pts | 1.1 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1301 | 601 | 0.46× [0.45, 0.47] | -117.7% | [-120.7%, -114.8%] | 4.1 pts | 2.8 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 8 | 461 | 762 | 1.65× [1.65, 1.66] | +39.5% | [+39.3%, +39.8%] | 0.3 pts | 0.5 | yes | yes |
| single-value | street | 16384 | prefix | ordered | btree-map | 8 | 455 | 279 | 0.61× [0.59, 0.62] | -64.3% | [-68.1%, -60.5%] | 4.0 pts | 2.2 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 2040 | 211 | 0.10× [0.10, 0.11] | -865.2% | [-888.4%, -842.1%] | 21.9 pts | 1.0 | yes | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 2009 | 220 | 0.11× [0.11, 0.11] | -800.9% | [-819.8%, -782.1%] | 18.6 pts | 0.8 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 88.99 ms | 9.49 ms | 0.11× [0.11, 0.11] | -837.5% | [-851.4%, -823.5%] | 15.1 pts | 0.9 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 89.30 ms | 11.80 ms | 0.13× [0.13, 0.13] | -661.8% | [-674.0%, -649.7%] | 11.9 pts | 0.6 | yes | yes |
| single-value | street | 212449 | valuesFor | ordered | baseline | 8 | 118 | 150 | 1.28× [1.22, 1.36] | +22.2% | [+17.9%, +26.4%] | 4.1 pts | 3.0 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | btree-map | 8 | 132 | 235 | 1.80× [1.76, 1.85] | +44.5% | [+43.0%, +46.0%] | 1.4 pts | 1.4 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | baseline | 8 | 1584 | 3291 | 2.17× [2.01, 2.35] | +53.8% | [+50.1%, +57.5%] | 3.7 pts | 4.9 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | btree-map | 8 | 1512 | 938 | 0.64× [0.59, 0.69] | -56.8% | [-69.1%, -44.5%] | 13.4 pts | 5.6 | no | yes |
| single-value | street | 212449 | prefix | ordered | baseline | 8 | 4459 | 12.7 µs | 2.89× [2.76, 3.03] | +65.4% | [+63.8%, +67.0%] | 1.6 pts | 2.2 | yes | yes |
| single-value | street | 212449 | prefix | ordered | btree-map | 8 | 4405 | 3020 | 0.72× [0.66, 0.78] | -39.8% | [-51.1%, -28.6%] | 12.0 pts | 4.1 | no | yes |
| single-value | street | 212449 | churn | ordered | baseline | 8 | 2264 | 565 | 0.25× [0.24, 0.26] | -297.1% | [-310.0%, -284.1%] | 17.0 pts | 1.9 | yes | yes |
| single-value | street | 212449 | churn | ordered | btree-map | 8 | 2442 | 671 | 0.27× [0.27, 0.28] | -266.1% | [-269.9%, -262.4%] | 4.9 pts | 0.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value street n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -2.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +1.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 prefix: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.90% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -1.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of -2.74% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 prefix: ordered vs btree-map: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
