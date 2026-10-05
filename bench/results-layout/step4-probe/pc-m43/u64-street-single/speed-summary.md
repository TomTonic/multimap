| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | street | 4096 | valuesFor | ordered | baseline | 6 | 61.7 | 48.8 | 0.80× [0.78, 0.81] | -25.8% | [-27.6%, -23.9%] | 1.8 pts | 2.0 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 6 | 62.4 | 92.0 | 1.47× [1.45, 1.49] | +31.9% | [+30.9%, +32.8%] | 0.9 pts | 1.9 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 6 | 1052 | 2111 | 2.01× [1.98, 2.04] | +50.3% | [+49.5%, +51.0%] | 0.7 pts | 1.0 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 6 | 1032 | 497 | 0.48× [0.48, 0.49] | -108.1% | [-110.5%, -105.6%] | 2.3 pts | 1.1 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 6 | 204 | 233 | 1.14× [1.14, 1.15] | +12.6% | [+12.1%, +13.2%] | 0.5 pts | 0.8 | yes | yes |
| single-value | street | 4096 | prefix | ordered | btree-map | 6 | 203 | 131 | 0.64× [0.64, 0.65] | -55.1% | [-56.1%, -54.0%] | 1.0 pts | 0.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 6 | 178 | 125 | 0.70× [0.70, 0.71] | -42.2% | [-43.1%, -41.3%] | 0.8 pts | 0.6 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 6 | 178 | 137 | 0.77× [0.76, 0.79] | -29.5% | [-31.7%, -27.3%] | 2.1 pts | 1.4 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 6 | 2.43 ms | 1.87 ms | 0.77× [0.76, 0.78] | -30.0% | [-31.9%, -28.2%] | 1.8 pts | 1.4 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 6 | 2.46 ms | 2.08 ms | 0.85× [0.84, 0.86] | -17.3% | [-19.0%, -15.7%] | 1.6 pts | 1.3 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 73.3 | 68.5 | 0.93× [0.92, 0.94] | -7.2% | [-8.2%, -6.1%] | 1.0 pts | 1.0 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 73.8 | 124 | 1.66× [1.63, 1.70] | +39.9% | [+38.7%, +41.1%] | 1.3 pts | 2.5 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1215 | 2520 | 2.08× [2.07, 2.09] | +51.9% | [+51.7%, +52.1%] | 0.4 pts | 1.0 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1213 | 591 | 0.49× [0.48, 0.50] | -104.9% | [-109.1%, -100.8%] | 5.2 pts | 3.7 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 8 | 422 | 762 | 1.81× [1.78, 1.84] | +44.7% | [+43.9%, +45.5%] | 0.8 pts | 1.4 | yes | yes |
| single-value | street | 16384 | prefix | ordered | btree-map | 8 | 418 | 274 | 0.65× [0.64, 0.67] | -53.3% | [-56.6%, -50.1%] | 3.5 pts | 2.1 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 190 | 160 | 0.85× [0.84, 0.85] | -18.1% | [-18.8%, -17.4%] | 2.0 pts | 1.5 | yes | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 191 | 186 | 0.98× [0.97, 0.99] | -1.9% | [-3.3%, -0.6%] | 1.6 pts | 1.4 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 10.70 ms | 8.96 ms | 0.83× [0.82, 0.84] | -20.5% | [-21.8%, -19.2%] | 1.5 pts | 1.7 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 10.96 ms | 10.90 ms | 0.99× [0.97, 1.01] | -0.8% | [-2.8%, +1.3%] | 1.9 pts | 1.0 | no | no |
| single-value | street | 212449 | valuesFor | ordered | baseline | 8 | 116 | 153 | 1.33× [1.19, 1.52] | +25.0% | [+15.8%, +34.3%] | 8.8 pts | 5.6 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | btree-map | 8 | 124 | 224 | 1.79× [1.70, 1.89] | +44.1% | [+41.1%, +47.1%] | 3.3 pts | 5.3 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | baseline | 8 | 1425 | 3159 | 2.24× [2.14, 2.35] | +55.3% | [+53.2%, +57.4%] | 2.1 pts | 3.7 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | btree-map | 8 | 1367 | 875 | 0.65× [0.60, 0.71] | -53.8% | [-66.4%, -41.3%] | 11.8 pts | 5.5 | no | yes |
| single-value | street | 212449 | prefix | ordered | baseline | 8 | 4547 | 13.0 µs | 2.86× [2.75, 2.97] | +65.0% | [+63.7%, +66.3%] | 1.5 pts | 1.5 | yes | yes |
| single-value | street | 212449 | prefix | ordered | btree-map | 8 | 4385 | 2940 | 0.68× [0.67, 0.69] | -46.8% | [-49.1%, -44.5%] | 2.5 pts | 0.8 | yes | yes |
| single-value | street | 212449 | churn | ordered | baseline | 8 | 343 | 414 | 1.21× [1.19, 1.24] | +17.6% | [+16.1%, +19.0%] | 1.7 pts | 0.9 | yes | yes |
| single-value | street | 212449 | churn | ordered | btree-map | 8 | 402 | 470 | 1.17× [1.12, 1.23] | +14.8% | [+10.7%, +18.8%] | 3.8 pts | 3.7 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 prefix: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs btree-map: the pooled difference of -0.77% does not clear the 0.78% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=16384 build: ordered vs btree-map: the pooled interval [-2.81%, 1.27%] includes zero
- single-value street n=212449 valuesFor: ordered vs baseline: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs btree-map: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 churn: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
