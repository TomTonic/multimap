| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 6 | 72.3 | 66.8 | 0.93× [0.92, 0.94] | -7.6% | [-8.6%, -6.6%] | 1.0 pts | 1.2 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 6 | 72.5 | 94.2 | 1.28× [1.25, 1.31] | +21.8% | [+20.2%, +23.4%] | 1.5 pts | 2.1 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 6 | 2542 | 3687 | 1.46× [1.44, 1.48] | +31.5% | [+30.4%, +32.6%] | 1.0 pts | 0.8 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 6 | 2497 | 1019 | 0.40× [0.39, 0.40] | -152.1% | [-156.4%, -147.8%] | 4.1 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 6 | 366 | 374 | 1.02× [1.00, 1.03] | +1.5% | [+0.3%, +2.7%] | 1.1 pts | 0.8 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 6 | 364 | 177 | 0.48× [0.47, 0.48] | -108.8% | [-110.6%, -107.1%] | 1.7 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 6 | 1971 | 150 | 0.08× [0.08, 0.08] | -1201.3% | [-1229.7%, -1172.9%] | 27.1 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 6 | 1894 | 152 | 0.08× [0.08, 0.08] | -1161.9% | [-1207.9%, -1115.8%] | 43.9 pts | 1.3 | yes | yes |
| single-value-str | street | 4096 | build | ordered | baseline | 6 | 20.51 ms | 1.87 ms | 0.09× [0.09, 0.09] | -980.9% | [-1007.9%, -953.9%] | 25.7 pts | 1.4 | yes | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 6 | 20.46 ms | 2.23 ms | 0.11× [0.11, 0.11] | -813.3% | [-843.5%, -783.0%] | 28.8 pts | 2.6 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 6 | 85.1 | 86.4 | 1.02× [1.01, 1.03] | +1.9% | [+1.0%, +2.8%] | 0.9 pts | 1.0 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 6 | 86.6 | 131 | 1.48× [1.45, 1.52] | +32.7% | [+31.2%, +34.1%] | 1.3 pts | 2.2 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 2890 | 4266 | 1.47× [1.47, 1.48] | +32.2% | [+31.9%, +32.5%] | 0.3 pts | 0.6 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 6 | 2870 | 1290 | 0.45× [0.44, 0.45] | -123.3% | [-125.8%, -120.8%] | 2.4 pts | 1.8 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 6 | 937 | 1351 | 1.43× [1.40, 1.45] | +29.8% | [+28.8%, +30.9%] | 1.0 pts | 1.0 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 6 | 928 | 461 | 0.50× [0.49, 0.51] | -102.0% | [-106.0%, -97.9%] | 3.9 pts | 1.4 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 6 | 2222 | 229 | 0.10× [0.10, 0.11] | -876.5% | [-915.0%, -838.0%] | 36.7 pts | 1.2 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 6 | 2205 | 239 | 0.10× [0.10, 0.11] | -853.5% | [-894.3%, -812.8%] | 38.8 pts | 1.5 | yes | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 6 | 95.04 ms | 9.45 ms | 0.10× [0.10, 0.10] | -895.7% | [-914.8%, -876.6%] | 18.2 pts | 0.9 | yes | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 6 | 95.30 ms | 11.94 ms | 0.12× [0.12, 0.13] | -707.5% | [-719.5%, -695.6%] | 11.4 pts | 0.7 | yes | yes |
| single-value-str | street | 212449 | valuesFor | ordered | baseline | 8 | 142 | 187 | 1.35× [1.30, 1.40] | +25.8% | [+23.2%, +28.4%] | 4.1 pts | 4.3 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | btree-map | 8 | 158 | 252 | 1.62× [1.52, 1.73] | +38.3% | [+34.3%, +42.3%] | 4.6 pts | 5.3 | no | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 3348 | 5669 | 1.71× [1.63, 1.81] | +41.6% | [+38.6%, +44.7%] | 3.6 pts | 6.3 | yes | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | btree-map | 8 | 3256 | 1807 | 0.57× [0.53, 0.61] | -76.3% | [-89.2%, -63.5%] | 14.6 pts | 6.5 | no | yes |
| single-value-str | street | 212449 | prefix | ordered | baseline | 8 | 12.2 µs | 23.1 µs | 1.88× [1.82, 1.95] | +46.8% | [+45.0%, +48.6%] | 2.0 pts | 4.7 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | btree-map | 8 | 12.2 µs | 6936 | 0.58× [0.55, 0.60] | -73.1% | [-80.2%, -66.1%] | 8.4 pts | 3.3 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | baseline | 8 | 2395 | 563 | 0.23× [0.23, 0.24] | -325.7% | [-335.8%, -315.6%] | 9.6 pts | 0.9 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | btree-map | 8 | 2671 | 708 | 0.26× [0.26, 0.27] | -277.8% | [-283.7%, -272.0%] | 6.6 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=4096 build: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.84% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of -2.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 6.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs baseline: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
