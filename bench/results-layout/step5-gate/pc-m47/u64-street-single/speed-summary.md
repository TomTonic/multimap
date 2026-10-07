| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 57.5 | 48.8 | 0.85× [0.84, 0.85] | -18.3% | [-19.2%, -17.5%] | 1.0 pts | 1.5 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 58.1 | 91.3 | 1.57× [1.56, 1.58] | +36.4% | [+36.0%, +36.7%] | 0.4 pts | 1.3 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 1106 | 2129 | 1.93× [1.92, 1.94] | +48.2% | [+47.9%, +48.4%] | 0.3 pts | 0.5 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1093 | 515 | 0.47× [0.46, 0.48] | -113.6% | [-116.8%, -110.3%] | 3.1 pts | 1.7 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 8 | 216 | 236 | 1.09× [1.08, 1.10] | +8.2% | [+7.3%, +9.0%] | 0.9 pts | 1.3 | yes | yes |
| single-value | street | 4096 | prefix | ordered | btree-map | 8 | 216 | 134 | 0.62× [0.62, 0.62] | -61.1% | [-61.8%, -60.3%] | 1.3 pts | 1.0 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 188 | 133 | 0.70× [0.69, 0.72] | -41.9% | [-45.0%, -38.8%] | 3.1 pts | 1.5 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 188 | 142 | 0.75× [0.74, 0.76] | -33.0% | [-34.5%, -31.4%] | 1.6 pts | 1.0 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.68 ms | 1.99 ms | 0.76× [0.72, 0.82] | -30.8% | [-39.1%, -22.5%] | 7.9 pts | 2.4 | no | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.72 ms | 2.27 ms | 0.83× [0.81, 0.85] | -20.6% | [-23.4%, -17.8%] | 3.9 pts | 2.1 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 69.1 | 68.8 | 1.00× [0.98, 1.01] | -0.3% | [-2.1%, +1.4%] | 1.7 pts | 1.7 | yes | no |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 69.6 | 125 | 1.78× [1.75, 1.82] | +43.9% | [+42.8%, +45.1%] | 1.3 pts | 2.9 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1279 | 2575 | 2.01× [2.00, 2.02] | +50.3% | [+49.9%, +50.6%] | 0.4 pts | 1.1 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1265 | 615 | 0.48× [0.47, 0.50] | -107.1% | [-112.3%, -101.9%] | 5.5 pts | 3.8 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 8 | 447 | 774 | 1.74× [1.72, 1.75] | +42.4% | [+41.8%, +42.9%] | 0.5 pts | 0.8 | yes | yes |
| single-value | street | 16384 | prefix | ordered | btree-map | 8 | 442 | 282 | 0.63× [0.61, 0.65] | -58.9% | [-63.3%, -54.6%] | 4.2 pts | 2.7 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 202 | 176 | 0.87× [0.86, 0.89] | -14.6% | [-16.9%, -12.4%] | 2.1 pts | 2.2 | no | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 212 | 204 | 0.98× [0.97, 1.00] | -1.6% | [-3.2%, -0.0%] | 1.6 pts | 1.2 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 12.33 ms | 10.44 ms | 0.84× [0.83, 0.85] | -18.8% | [-20.2%, -17.3%] | 3.4 pts | 2.2 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 12.52 ms | 12.30 ms | 0.97× [0.95, 0.99] | -3.6% | [-5.7%, -1.5%] | 2.1 pts | 1.4 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | baseline | 8 | 119 | 156 | 1.24× [1.10, 1.42] | +19.2% | [+8.8%, +29.7%] | 12.4 pts | 13.8 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | btree-map | 8 | 134 | 321 | 2.13× [1.66, 2.94] | +52.9% | [+39.9%, +66.0%] | 12.4 pts | 12.9 | no | yes |
| single-value | street | 212449 | valuesBetween | ordered | baseline | 8 | 1688 | 3745 | 2.18× [1.91, 2.56] | +54.2% | [+47.5%, +60.9%] | 7.7 pts | 16.0 | no | yes |
| single-value | street | 212449 | valuesBetween | ordered | btree-map | 8 | 1568 | 1490 | 0.82× [0.63, 1.18] | -21.3% | [-58.0%, +15.4%] | 35.1 pts | 26.9 | no | no |
| single-value | street | 212449 | prefix | ordered | baseline | 8 | 4902 | 13.4 µs | 2.72× [2.48, 3.01] | +63.3% | [+59.7%, +66.8%] | 4.0 pts | 3.7 | yes | yes |
| single-value | street | 212449 | prefix | ordered | btree-map | 8 | 4899 | 4454 | 0.84× [0.67, 1.13] | -18.7% | [-48.9%, +11.5%] | 30.1 pts | 9.9 | no | no |
| single-value | street | 212449 | churn | ordered | baseline | 8 | 357 | 426 | 1.19× [1.17, 1.21] | +15.9% | [+14.7%, +17.1%] | 1.6 pts | 1.5 | yes | yes |
| single-value | street | 212449 | churn | ordered | btree-map | 8 | 421 | 474 | 1.17× [1.14, 1.21] | +14.8% | [+12.2%, +17.3%] | 2.8 pts | 2.7 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.35% does not clear the 0.90% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.09%, 1.40%] includes zero
- single-value street n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 prefix: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs baseline: the processes scatter 13.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs baseline: 7 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=212449 valuesFor: ordered vs btree-map: the processes scatter 12.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs baseline: the processes scatter 16.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs btree-map: the pooled interval [-57.99%, 15.42%] includes zero
- single-value street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 26.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs btree-map: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=212449 prefix: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 prefix: ordered vs btree-map: the pooled interval [-48.94%, 11.45%] includes zero
- single-value street n=212449 prefix: ordered vs btree-map: the processes scatter 9.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 prefix: ordered vs btree-map: 2 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=212449 churn: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
