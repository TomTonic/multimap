| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 102 | 96.8 | 0.95× [0.95, 0.95] | -5.3% | [-5.7%, -4.9%] | 0.4 pts | 0.3 | yes | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 105 | 111 | 1.05× [1.01, 1.09] | +4.9% | [+1.2%, +8.6%] | 4.6 pts | 3.3 | no | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 3108 | 4072 | 1.31× [1.31, 1.32] | +23.9% | [+23.4%, +24.3%] | 0.6 pts | 0.5 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 3055 | 1531 | 0.49× [0.48, 0.49] | -104.5% | [-106.5%, -102.5%] | 3.0 pts | 0.8 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 3650 | 4775 | 1.31× [1.30, 1.32] | +23.8% | [+23.2%, +24.3%] | 0.6 pts | 0.3 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 3711 | 1886 | 0.51× [0.50, 0.53] | -94.2% | [-98.6%, -89.9%] | 4.1 pts | 0.7 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 266 | 200 | 0.75× [0.74, 0.76] | -33.2% | [-35.5%, -30.9%] | 2.4 pts | 1.3 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 261 | 180 | 0.69× [0.68, 0.69] | -45.5% | [-46.6%, -44.4%] | 1.1 pts | 0.7 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 4.05 ms | 2.94 ms | 0.73× [0.72, 0.74] | -37.4% | [-39.1%, -35.6%] | 1.8 pts | 1.4 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 4.17 ms | 2.63 ms | 0.63× [0.62, 0.64] | -58.5% | [-60.3%, -56.7%] | 1.9 pts | 1.0 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 127 | 127 | 1.00× [0.99, 1.02] | +0.1% | [-1.4%, +1.6%] | 1.7 pts | 1.9 | yes | no |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 129 | 160 | 1.24× [1.21, 1.28] | +19.3% | [+17.0%, +21.6%] | 2.1 pts | 2.0 | no | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3401 | 4719 | 1.39× [1.38, 1.40] | +28.0% | [+27.4%, +28.5%] | 0.5 pts | 1.3 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 3398 | 1966 | 0.58× [0.57, 0.58] | -73.4% | [-75.2%, -71.6%] | 2.4 pts | 1.9 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 8 | 15.9 µs | 25.1 µs | 1.53× [1.50, 1.55] | +34.4% | [+33.4%, +35.5%] | 1.0 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 8 | 15.0 µs | 8347 | 0.56× [0.53, 0.58] | -79.5% | [-87.1%, -71.8%] | 7.2 pts | 0.7 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 8 | 318 | 255 | 0.81× [0.80, 0.82] | -23.4% | [-25.3%, -21.6%] | 2.1 pts | 1.6 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 8 | 321 | 248 | 0.77× [0.76, 0.79] | -29.1% | [-31.4%, -26.8%] | 2.2 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 8 | 18.20 ms | 14.34 ms | 0.79× [0.78, 0.79] | -27.3% | [-28.5%, -26.1%] | 1.5 pts | 0.9 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 8 | 18.64 ms | 13.60 ms | 0.73× [0.72, 0.74] | -36.9% | [-38.8%, -34.9%] | 1.9 pts | 0.7 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 173 | 193 | 1.10× [1.07, 1.14] | +9.2% | [+6.1%, +12.3%] | 2.9 pts | 2.6 | no | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 180 | 245 | 1.37× [1.36, 1.39] | +27.2% | [+26.2%, +28.1%] | 1.0 pts | 1.0 | yes | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3701 | 5364 | 1.45× [1.43, 1.46] | +30.9% | [+30.0%, +31.7%] | 0.8 pts | 1.4 | yes | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 3666 | 2329 | 0.64× [0.61, 0.67] | -56.1% | [-63.9%, -48.3%] | 7.3 pts | 3.0 | no | yes |
| single-value-str | dirs | 86215 | prefix | ordered | baseline | 8 | 91.4 µs | 154.6 µs | 1.60× [1.59, 1.62] | +37.6% | [+37.0%, +38.1%] | 0.5 pts | 0.7 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | btree-map | 8 | 86.6 µs | 53.1 µs | 0.58× [0.57, 0.60] | -71.1% | [-74.6%, -67.6%] | 3.3 pts | 0.8 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | baseline | 8 | 444 | 432 | 0.98× [0.97, 1.00] | -1.7% | [-3.6%, +0.1%] | 2.9 pts | 1.4 | yes | no |
| single-value-str | dirs | 86215 | churn | ordered | btree-map | 8 | 516 | 471 | 0.93× [0.92, 0.93] | -8.1% | [-9.1%, -7.1%] | 1.1 pts | 0.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.11% does not clear the 0.37% noise floor, the bound on what the harness reports between identical code in every process
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.36%, 1.57%] includes zero
- single-value-str dirs n=16384 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -2.04% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 churn: ordered vs baseline: the pooled interval [-3.57%, 0.11%] includes zero
