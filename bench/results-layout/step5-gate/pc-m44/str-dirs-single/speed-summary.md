| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 109 | 97.7 | 0.90× [0.89, 0.92] | -10.6% | [-12.2%, -8.9%] | 1.8 pts | 1.6 | yes | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 114 | 111 | 0.98× [0.95, 1.01] | -2.3% | [-5.2%, +0.6%] | 3.1 pts | 2.1 | no | no |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 3166 | 4149 | 1.30× [1.29, 1.32] | +23.2% | [+22.4%, +24.1%] | 0.8 pts | 0.7 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 3108 | 1414 | 0.44× [0.43, 0.45] | -126.4% | [-131.9%, -120.8%] | 6.1 pts | 1.6 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 3740 | 4833 | 1.27× [1.24, 1.30] | +21.3% | [+19.2%, +23.4%] | 2.3 pts | 1.2 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 3711 | 1772 | 0.48× [0.47, 0.49] | -106.9% | [-111.2%, -102.6%] | 4.7 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 281 | 201 | 0.72× [0.71, 0.73] | -39.0% | [-40.6%, -37.4%] | 1.7 pts | 1.3 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 281 | 182 | 0.65× [0.64, 0.66] | -54.8% | [-57.4%, -52.2%] | 2.6 pts | 1.2 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 4.30 ms | 3.00 ms | 0.70× [0.69, 0.70] | -43.6% | [-45.2%, -42.1%] | 1.6 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 4.48 ms | 2.64 ms | 0.60× [0.59, 0.60] | -68.0% | [-69.9%, -66.1%] | 2.4 pts | 1.2 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 133 | 129 | 0.98× [0.95, 1.00] | -2.5% | [-5.0%, -0.0%] | 2.5 pts | 2.7 | no | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 135 | 164 | 1.21× [1.19, 1.24] | +17.7% | [+15.8%, +19.6%] | 1.8 pts | 1.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3497 | 4761 | 1.36× [1.35, 1.36] | +26.4% | [+26.1%, +26.7%] | 0.3 pts | 0.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 3507 | 1888 | 0.54× [0.53, 0.54] | -86.7% | [-87.5%, -85.9%] | 1.0 pts | 0.8 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 8 | 16.1 µs | 24.5 µs | 1.50× [1.48, 1.52] | +33.4% | [+32.4%, +34.4%] | 1.3 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 8 | 15.1 µs | 8075 | 0.55× [0.52, 0.57] | -83.1% | [-90.8%, -75.3%] | 7.2 pts | 0.5 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 8 | 339 | 265 | 0.79× [0.78, 0.80] | -26.9% | [-28.4%, -25.3%] | 1.9 pts | 1.0 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 8 | 333 | 248 | 0.75× [0.74, 0.76] | -33.9% | [-36.0%, -31.9%] | 1.9 pts | 1.4 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 8 | 19.15 ms | 14.32 ms | 0.75× [0.74, 0.76] | -33.8% | [-35.2%, -32.4%] | 1.6 pts | 1.2 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 8 | 19.68 ms | 13.68 ms | 0.70× [0.69, 0.71] | -43.3% | [-44.9%, -41.8%] | 1.6 pts | 0.7 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 177 | 189 | 1.08× [1.06, 1.10] | +7.4% | [+5.9%, +8.9%] | 1.8 pts | 1.9 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 185 | 249 | 1.34× [1.31, 1.36] | +25.1% | [+23.8%, +26.5%] | 1.3 pts | 1.2 | yes | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3784 | 5285 | 1.40× [1.39, 1.42] | +28.7% | [+27.9%, +29.5%] | 0.9 pts | 1.3 | yes | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 3751 | 2170 | 0.59× [0.57, 0.61] | -69.5% | [-75.7%, -63.3%] | 6.3 pts | 2.2 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | baseline | 8 | 91.9 µs | 149.4 µs | 1.55× [1.54, 1.56] | +35.4% | [+35.0%, +35.8%] | 0.5 pts | 0.7 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | btree-map | 8 | 90.3 µs | 49.2 µs | 0.54× [0.53, 0.55] | -85.2% | [-89.5%, -81.0%] | 4.4 pts | 1.0 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | baseline | 8 | 470 | 458 | 0.96× [0.93, 0.98] | -4.7% | [-7.3%, -2.1%] | 2.6 pts | 1.1 | no | yes |
| single-value-str | dirs | 86215 | churn | ordered | btree-map | 8 | 517 | 459 | 0.90× [0.89, 0.91] | -11.0% | [-12.1%, -10.0%] | 1.2 pts | 0.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the pooled interval [-5.20%, 0.63%] includes zero
- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +7.24% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=16384 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -2.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
