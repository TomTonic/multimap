| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 8 | 66.8 | 47.6 | 0.71× [0.69, 0.72] | -41.5% | [-43.9%, -39.1%] | 2.4 pts | 2.5 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 8 | 67.5 | 92.0 | 1.34× [1.31, 1.38] | +25.5% | [+23.6%, +27.5%] | 3.0 pts | 4.6 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 3670 | 1849 | 0.51× [0.50, 0.51] | -97.8% | [-99.6%, -96.0%] | 1.8 pts | 0.7 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 8 | 3639 | 1127 | 0.30× [0.30, 0.31] | -230.5% | [-236.9%, -224.1%] | 6.2 pts | 1.2 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 8 | 365 | 204 | 0.56× [0.55, 0.56] | -79.4% | [-81.5%, -77.4%] | 2.1 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 8 | 365 | 188 | 0.51× [0.50, 0.51] | -97.3% | [-98.7%, -95.9%] | 1.6 pts | 0.6 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 8 | 128 | 118 | 0.92× [0.91, 0.93] | -8.2% | [-9.4%, -7.0%] | 1.2 pts | 1.1 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 8 | 127 | 139 | 1.09× [1.08, 1.10] | +8.3% | [+7.5%, +9.0%] | 1.2 pts | 1.2 | yes | yes |
| single-value-str | street | 4096 | build | ordered | baseline | 8 | 1.99 ms | 1.73 ms | 0.87× [0.85, 0.89] | -15.2% | [-18.2%, -12.3%] | 2.8 pts | 1.7 | no | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 8 | 1.98 ms | 2.17 ms | 1.13× [1.10, 1.15] | +11.2% | [+9.4%, +12.9%] | 1.7 pts | 1.3 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 8 | 87.0 | 66.2 | 0.76× [0.75, 0.78] | -31.2% | [-33.5%, -28.9%] | 2.7 pts | 2.6 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 8 | 89.2 | 127 | 1.42× [1.38, 1.47] | +29.8% | [+27.6%, +32.0%] | 2.3 pts | 3.6 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 4213 | 2241 | 0.53× [0.52, 0.54] | -88.9% | [-91.1%, -86.6%] | 2.2 pts | 1.4 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 8 | 4211 | 1384 | 0.33× [0.33, 0.33] | -202.9% | [-206.6%, -199.3%] | 4.8 pts | 2.5 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 8 | 1330 | 662 | 0.50× [0.49, 0.50] | -100.9% | [-103.6%, -98.3%] | 2.6 pts | 0.7 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 8 | 1323 | 496 | 0.37× [0.37, 0.38] | -166.9% | [-170.8%, -163.0%] | 4.0 pts | 1.0 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 8 | 164 | 156 | 0.95× [0.94, 0.96] | -5.7% | [-6.8%, -4.7%] | 1.1 pts | 1.1 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 8 | 167 | 196 | 1.17× [1.13, 1.21] | +14.5% | [+11.8%, +17.1%] | 2.8 pts | 3.3 | no | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 8 | 9.56 ms | 8.56 ms | 0.90× [0.89, 0.91] | -11.7% | [-12.9%, -10.4%] | 1.2 pts | 1.0 | yes | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 8 | 9.68 ms | 11.57 ms | 1.19× [1.17, 1.21] | +16.1% | [+14.8%, +17.4%] | 1.4 pts | 0.7 | yes | yes |
| single-value-str | street | 212449 | valuesFor | ordered | baseline | 8 | 200 | 169 | 0.88× [0.85, 0.92] | -13.7% | [-18.2%, -9.1%] | 4.4 pts | 2.7 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | btree-map | 8 | 234 | 252 | 1.05× [1.01, 1.09] | +4.7% | [+1.0%, +8.4%] | 3.8 pts | 2.7 | no | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 6232 | 4071 | 0.67× [0.64, 0.69] | -50.3% | [-55.3%, -45.3%] | 5.8 pts | 2.4 | yes | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | btree-map | 8 | 6322 | 2450 | 0.39× [0.38, 0.40] | -156.8% | [-164.4%, -149.2%] | 16.7 pts | 2.3 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | baseline | 8 | 24.6 µs | 14.5 µs | 0.59× [0.57, 0.60] | -69.9% | [-74.3%, -65.6%] | 4.9 pts | 1.7 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | btree-map | 8 | 23.6 µs | 8224 | 0.35× [0.34, 0.36] | -184.5% | [-192.0%, -177.0%] | 11.3 pts | 1.9 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | baseline | 8 | 419 | 395 | 0.96× [0.94, 0.98] | -4.4% | [-6.3%, -2.5%] | 2.0 pts | 1.2 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | btree-map | 8 | 510 | 499 | 0.98× [0.97, 1.00] | -1.6% | [-2.9%, -0.2%] | 1.4 pts | 0.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=16384 churn: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
