| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 96.2 | 86.4 | 0.90× [0.89, 0.91] | -10.8% | [-11.7%, -9.9%] | 0.9 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 98.2 | 166 | 1.71× [1.69, 1.73] | +41.5% | [+40.8%, +42.2%] | 0.8 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 196 | 158 | 0.80× [0.78, 0.81] | -25.2% | [-27.6%, -22.9%] | 2.5 pts | 1.7 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 193 | 231 | 1.20× [1.19, 1.21] | +16.6% | [+15.7%, +17.6%] | 0.9 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 9.51 ms | 7.36 ms | 0.77× [0.76, 0.78] | -29.9% | [-31.2%, -28.6%] | 1.5 pts | 1.6 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 9.61 ms | 11.02 ms | 1.14× [1.13, 1.15] | +12.2% | [+11.6%, +12.8%] | 0.8 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 123 | 117 | 0.95× [0.95, 0.95] | -5.2% | [-5.7%, -4.7%] | 0.6 pts | 0.6 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 125 | 225 | 1.84× [1.79, 1.88] | +45.5% | [+44.2%, +46.8%] | 1.3 pts | 2.0 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 243 | 214 | 0.89× [0.87, 0.91] | -12.4% | [-15.2%, -9.6%] | 3.5 pts | 2.4 | no | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 240 | 338 | 1.39× [1.36, 1.43] | +28.3% | [+26.7%, +30.0%] | 2.1 pts | 2.2 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 47.36 ms | 39.21 ms | 0.83× [0.82, 0.83] | -21.0% | [-22.0%, -20.0%] | 1.1 pts | 1.9 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 47.84 ms | 61.28 ms | 1.28× [1.27, 1.29] | +21.6% | [+21.0%, +22.2%] | 0.6 pts | 1.3 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 6 | 65.1 | 54.5 | 0.84× [0.83, 0.85] | -18.8% | [-20.6%, -17.0%] | 1.7 pts | 2.0 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 6 | 66.0 | 140 | 2.12× [2.08, 2.16] | +52.9% | [+52.0%, +53.8%] | 0.9 pts | 2.0 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 6 | 146 | 106 | 0.72× [0.71, 0.73] | -39.3% | [-41.2%, -37.5%] | 1.8 pts | 1.3 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 6 | 146 | 188 | 1.30× [1.28, 1.32] | +23.0% | [+22.0%, +24.0%] | 0.9 pts | 1.1 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 6 | 5.90 ms | 4.17 ms | 0.71× [0.70, 0.71] | -41.1% | [-42.4%, -39.9%] | 1.2 pts | 1.5 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 6 | 5.98 ms | 7.58 ms | 1.28× [1.25, 1.31] | +21.9% | [+20.2%, +23.6%] | 1.6 pts | 1.4 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 81.5 | 75.5 | 0.92× [0.91, 0.94] | -8.3% | [-10.1%, -6.4%] | 1.7 pts | 1.5 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 82.4 | 186 | 2.25× [2.21, 2.30] | +55.6% | [+54.7%, +56.5%] | 0.9 pts | 2.5 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 176 | 143 | 0.81× [0.81, 0.82] | -23.3% | [-24.2%, -22.3%] | 1.1 pts | 0.9 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 177 | 267 | 1.51× [1.48, 1.54] | +33.7% | [+32.4%, +35.1%] | 1.5 pts | 2.1 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 28.41 ms | 21.98 ms | 0.77× [0.76, 0.78] | -29.6% | [-30.8%, -28.4%] | 1.2 pts | 1.7 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 29.00 ms | 42.08 ms | 1.45× [1.44, 1.47] | +31.1% | [+30.4%, +31.9%] | 0.8 pts | 0.8 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 46.7 | 37.9 | 0.81× [0.80, 0.83] | -23.3% | [-25.5%, -21.1%] | 2.1 pts | 2.1 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 46.8 | 161 | 3.45× [3.42, 3.48] | +71.0% | [+70.8%, +71.2%] | 0.3 pts | 2.0 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 70.4 | 57.4 | 0.81× [0.80, 0.81] | -23.7% | [-24.5%, -23.0%] | 1.2 pts | 1.6 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 8 | 70.6 | 156 | 2.21× [2.18, 2.24] | +54.7% | [+54.1%, +55.3%] | 0.6 pts | 1.4 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 9.09 ms | 5.80 ms | 0.64× [0.64, 0.65] | -55.9% | [-57.3%, -54.4%] | 1.7 pts | 2.3 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 8 | 9.13 ms | 14.80 ms | 1.63× [1.62, 1.64] | +38.6% | [+38.3%, +39.0%] | 0.6 pts | 1.5 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 8 | 49.0 | 47.1 | 0.96× [0.95, 0.97] | -4.3% | [-5.3%, -3.2%] | 1.0 pts | 1.8 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 8 | 49.1 | 195 | 3.95× [3.92, 3.99] | +74.7% | [+74.5%, +74.9%] | 0.2 pts | 1.4 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 8 | 90.9 | 77.4 | 0.85× [0.84, 0.86] | -17.7% | [-18.4%, -16.9%] | 0.7 pts | 1.0 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 8 | 96.4 | 234 | 2.41× [2.38, 2.44] | +58.4% | [+57.9%, +59.0%] | 0.6 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 8 | 34.76 ms | 28.92 ms | 0.83× [0.82, 0.85] | -19.8% | [-21.6%, -17.9%] | 2.0 pts | 3.0 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 8 | 35.63 ms | 79.84 ms | 2.24× [2.22, 2.26] | +55.3% | [+54.9%, +55.7%] | 0.5 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.49% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=16384 build: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
