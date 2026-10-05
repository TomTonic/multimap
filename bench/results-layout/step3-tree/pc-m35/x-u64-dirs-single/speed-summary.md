| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 77.9 | 75.0 | 0.96× [0.94, 0.98] | -4.4% | [-6.3%, -2.5%] | 1.9 pts | 1.9 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 79.2 | 108 | 1.38× [1.36, 1.40] | +27.3% | [+26.3%, +28.3%] | 1.2 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2332 | 1994 | 0.85× [0.84, 0.87] | -17.2% | [-19.7%, -14.6%] | 2.6 pts | 1.1 | no | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 2293 | 554 | 0.24× [0.24, 0.24] | -314.6% | [-317.5%, -311.6%] | 3.8 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 8 | 2705 | 2239 | 0.82× [0.82, 0.83] | -21.3% | [-22.3%, -20.2%] | 1.5 pts | 0.7 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | btree-map | 8 | 2699 | 673 | 0.25× [0.24, 0.25] | -305.5% | [-310.2%, -300.7%] | 4.6 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 196 | 184 | 0.94× [0.92, 0.96] | -6.6% | [-8.6%, -4.7%] | 2.2 pts | 2.1 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 197 | 176 | 0.89× [0.88, 0.90] | -12.5% | [-13.7%, -11.2%] | 1.6 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 2.92 ms | 2.58 ms | 0.89× [0.88, 0.90] | -12.6% | [-14.0%, -11.3%] | 1.5 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 2.98 ms | 2.57 ms | 0.87× [0.86, 0.88] | -14.8% | [-15.7%, -14.0%] | 1.8 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 106 | 103 | 0.98× [0.95, 1.00] | -2.5% | [-4.8%, -0.2%] | 2.3 pts | 2.2 | no | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 108 | 158 | 1.47× [1.44, 1.51] | +32.2% | [+30.7%, +33.6%] | 1.4 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 2823 | 2532 | 0.89× [0.88, 0.90] | -12.0% | [-13.3%, -10.8%] | 1.8 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 2822 | 722 | 0.26× [0.25, 0.26] | -291.1% | [-296.3%, -286.0%] | 5.7 pts | 1.6 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 8 | 11.3 µs | 9571 | 0.87× [0.86, 0.88] | -15.1% | [-16.8%, -13.5%] | 3.0 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | btree-map | 8 | 11.4 µs | 2444 | 0.21× [0.21, 0.21] | -377.8% | [-385.6%, -370.1%] | 8.0 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 248 | 236 | 0.94× [0.93, 0.95] | -6.4% | [-7.0%, -5.8%] | 0.9 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 249 | 240 | 0.97× [0.96, 0.98] | -3.3% | [-4.2%, -2.5%] | 1.0 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 13.54 ms | 12.43 ms | 0.92× [0.91, 0.93] | -8.9% | [-9.8%, -8.1%] | 0.8 pts | 1.6 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 13.58 ms | 12.97 ms | 0.96× [0.95, 0.96] | -4.5% | [-5.2%, -3.7%] | 0.8 pts | 1.0 | yes | yes |
| single-value | dirs | 86215 | valuesFor | ordered | baseline | 8 | 155 | 150 | 0.98× [0.96, 1.00] | -2.3% | [-4.5%, -0.1%] | 2.1 pts | 2.0 | no | yes |
| single-value | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 161 | 230 | 1.42× [1.38, 1.46] | +29.6% | [+27.7%, +31.6%] | 2.3 pts | 2.7 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3146 | 2867 | 0.91× [0.91, 0.92] | -9.6% | [-10.4%, -8.8%] | 0.8 pts | 0.9 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 3146 | 980 | 0.31× [0.30, 0.32] | -219.4% | [-227.9%, -210.9%] | 14.5 pts | 3.1 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | baseline | 8 | 78.2 µs | 71.8 µs | 0.90× [0.89, 0.91] | -11.5% | [-12.5%, -10.5%] | 1.0 pts | 1.0 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | btree-map | 8 | 78.2 µs | 15.5 µs | 0.19× [0.18, 0.20] | -432.7% | [-453.1%, -412.2%] | 19.9 pts | 2.3 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | baseline | 8 | 409 | 383 | 0.95× [0.94, 0.97] | -4.8% | [-6.8%, -2.9%] | 1.8 pts | 1.3 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | btree-map | 8 | 472 | 435 | 0.93× [0.92, 0.94] | -7.3% | [-8.5%, -6.1%] | 1.6 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of -6.99% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -2.21% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=86215 prefix: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
