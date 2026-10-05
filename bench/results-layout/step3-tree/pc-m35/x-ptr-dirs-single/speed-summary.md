| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 78.0 | 75.1 | 0.95× [0.93, 0.97] | -5.0% | [-7.2%, -2.8%] | 2.1 pts | 2.1 | no | yes |
| single-value-ptr | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 79.5 | 108 | 1.37× [1.35, 1.39] | +27.0% | [+25.8%, +28.2%] | 1.2 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2361 | 2039 | 0.87× [0.85, 0.89] | -15.0% | [-17.1%, -12.8%] | 2.0 pts | 0.9 | no | yes |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 2308 | 552 | 0.24× [0.23, 0.24] | -320.9% | [-325.8%, -315.9%] | 5.1 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 2716 | 2236 | 0.82× [0.81, 0.83] | -21.7% | [-23.5%, -19.8%] | 1.8 pts | 0.8 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | btree-map | 8 | 2730 | 656 | 0.24× [0.24, 0.24] | -317.3% | [-319.1%, -315.5%] | 2.8 pts | 0.7 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 199 | 190 | 0.95× [0.95, 0.96] | -4.8% | [-5.7%, -3.8%] | 1.0 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | btree-map | 8 | 200 | 176 | 0.88× [0.87, 0.89] | -13.6% | [-15.4%, -11.9%] | 1.9 pts | 1.2 | yes | yes |
| single-value-ptr | dirs | 4096 | build | ordered | baseline | 8 | 2.87 ms | 2.66 ms | 0.93× [0.92, 0.94] | -7.8% | [-8.7%, -6.9%] | 0.9 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 4096 | build | ordered | btree-map | 8 | 2.90 ms | 2.63 ms | 0.90× [0.89, 0.91] | -11.0% | [-11.9%, -10.1%] | 1.3 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | baseline | 8 | 109 | 106 | 0.97× [0.95, 0.99] | -3.5% | [-5.6%, -1.5%] | 2.0 pts | 2.2 | no | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 110 | 158 | 1.44× [1.42, 1.46] | +30.5% | [+29.5%, +31.5%] | 1.0 pts | 1.2 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 2879 | 2586 | 0.90× [0.89, 0.91] | -11.2% | [-12.6%, -9.8%] | 1.3 pts | 1.5 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 2889 | 745 | 0.26× [0.26, 0.26] | -286.7% | [-291.6%, -281.8%] | 4.9 pts | 1.4 | yes | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | baseline | 8 | 11.4 µs | 9879 | 0.83× [0.81, 0.86] | -20.2% | [-24.2%, -16.2%] | 4.2 pts | 0.8 | no | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | btree-map | 8 | 12.0 µs | 2671 | 0.23× [0.22, 0.23] | -343.0% | [-358.6%, -327.3%] | 15.2 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | baseline | 8 | 252 | 243 | 0.96× [0.95, 0.98] | -3.6% | [-4.8%, -2.4%] | 1.2 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | btree-map | 8 | 250 | 239 | 0.95× [0.93, 0.97] | -5.3% | [-7.3%, -3.3%] | 1.9 pts | 1.4 | yes | yes |
| single-value-ptr | dirs | 16384 | build | ordered | baseline | 8 | 13.73 ms | 12.92 ms | 0.94× [0.93, 0.94] | -6.5% | [-7.1%, -5.9%] | 0.6 pts | 0.8 | yes | yes |
| single-value-ptr | dirs | 16384 | build | ordered | btree-map | 8 | 13.91 ms | 13.22 ms | 0.95× [0.95, 0.96] | -4.9% | [-5.8%, -4.1%] | 1.2 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | baseline | 6 | 159 | 152 | 0.96× [0.95, 0.97] | -3.9% | [-4.9%, -2.9%] | 1.0 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | btree-map | 6 | 170 | 231 | 1.37× [1.33, 1.41] | +27.1% | [+24.9%, +29.3%] | 2.1 pts | 2.2 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 6 | 3198 | 2903 | 0.91× [0.91, 0.92] | -9.7% | [-10.2%, -9.2%] | 0.5 pts | 0.6 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | btree-map | 6 | 3233 | 1009 | 0.32× [0.30, 0.34] | -213.0% | [-230.4%, -195.6%] | 16.6 pts | 2.5 | yes | yes |
| single-value-ptr | dirs | 86215 | prefix | ordered | baseline | 6 | 81.7 µs | 71.1 µs | 0.89× [0.88, 0.89] | -12.4% | [-13.1%, -11.8%] | 0.6 pts | 0.7 | yes | yes |
| single-value-ptr | dirs | 86215 | prefix | ordered | btree-map | 6 | 80.4 µs | 16.8 µs | 0.20× [0.19, 0.21] | -397.9% | [-414.7%, -381.1%] | 16.0 pts | 1.8 | yes | yes |
| single-value-ptr | dirs | 86215 | churn | ordered | baseline | 6 | 394 | 379 | 0.96× [0.95, 0.98] | -4.0% | [-5.8%, -2.2%] | 1.7 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 86215 | churn | ordered | btree-map | 6 | 465 | 414 | 0.91× [0.90, 0.93] | -9.7% | [-11.7%, -7.7%] | 1.9 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -1.60% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of -4.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=86215 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.24% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 prefix: ordered vs btree-map: the A/A validations found a systematic difference of +2.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
