| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 96.7 | 86.3 | 0.89× [0.88, 0.91] | -11.9% | [-13.8%, -10.1%] | 2.0 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 98.2 | 167 | 1.72× [1.70, 1.74] | +41.8% | [+41.1%, +42.5%] | 0.9 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 225 | 161 | 0.71× [0.71, 0.72] | -40.7% | [-41.6%, -39.7%] | 1.3 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 228 | 243 | 1.07× [1.06, 1.09] | +6.6% | [+5.2%, +7.9%] | 1.7 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 10.50 ms | 7.49 ms | 0.71× [0.70, 0.72] | -41.0% | [-42.8%, -39.3%] | 1.6 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 10.51 ms | 11.04 ms | 1.05× [1.04, 1.06] | +4.9% | [+3.7%, +6.0%] | 1.2 pts | 1.4 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 6 | 123 | 117 | 0.95× [0.94, 0.97] | -5.1% | [-6.8%, -3.4%] | 1.6 pts | 1.6 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 6 | 125 | 226 | 1.81× [1.80, 1.82] | +44.7% | [+44.4%, +45.1%] | 0.3 pts | 0.4 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 6 | 267 | 212 | 0.79× [0.78, 0.80] | -26.8% | [-28.0%, -25.7%] | 1.1 pts | 0.9 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 6 | 270 | 339 | 1.25× [1.23, 1.28] | +20.3% | [+18.5%, +22.1%] | 1.7 pts | 1.4 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 6 | 50.95 ms | 39.29 ms | 0.77× [0.76, 0.78] | -29.4% | [-30.9%, -27.8%] | 1.5 pts | 1.2 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 6 | 51.67 ms | 61.59 ms | 1.19× [1.18, 1.20] | +15.9% | [+15.1%, +16.8%] | 0.8 pts | 1.1 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 65.3 | 54.5 | 0.83× [0.82, 0.85] | -19.8% | [-22.3%, -17.3%] | 2.3 pts | 2.0 | no | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 66.0 | 140 | 2.13× [2.11, 2.14] | +53.0% | [+52.7%, +53.3%] | 0.5 pts | 0.9 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 167 | 109 | 0.65× [0.64, 0.65] | -54.2% | [-55.6%, -52.8%] | 1.3 pts | 0.7 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 168 | 191 | 1.15× [1.13, 1.17] | +13.2% | [+11.6%, +14.9%] | 1.7 pts | 2.5 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 6.45 ms | 4.25 ms | 0.66× [0.65, 0.66] | -52.6% | [-54.2%, -51.0%] | 1.6 pts | 1.6 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 6.52 ms | 7.77 ms | 1.19× [1.17, 1.21] | +16.1% | [+14.6%, +17.6%] | 1.6 pts | 1.2 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 6 | 81.9 | 75.7 | 0.92× [0.91, 0.94] | -8.5% | [-10.1%, -6.9%] | 1.5 pts | 1.5 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 6 | 82.9 | 190 | 2.28× [2.25, 2.32] | +56.2% | [+55.6%, +56.8%] | 0.6 pts | 1.6 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 6 | 205 | 149 | 0.73× [0.72, 0.73] | -37.7% | [-38.8%, -36.6%] | 1.0 pts | 0.7 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 6 | 204 | 271 | 1.33× [1.29, 1.36] | +24.6% | [+22.6%, +26.5%] | 1.9 pts | 2.7 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 6 | 30.55 ms | 22.17 ms | 0.73× [0.72, 0.74] | -37.4% | [-38.9%, -35.8%] | 1.5 pts | 1.3 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 6 | 31.53 ms | 42.21 ms | 1.35× [1.32, 1.37] | +25.8% | [+24.3%, +27.2%] | 1.4 pts | 1.2 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 6 | 46.4 | 38.1 | 0.82× [0.81, 0.83] | -21.9% | [-22.8%, -20.9%] | 0.9 pts | 0.7 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 6 | 46.5 | 161 | 3.46× [3.43, 3.49] | +71.1% | [+70.8%, +71.4%] | 0.3 pts | 1.1 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 6 | 77.3 | 59.7 | 0.77× [0.76, 0.78] | -29.7% | [-31.3%, -28.2%] | 1.5 pts | 1.4 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 6 | 77.7 | 160 | 2.07× [2.05, 2.10] | +51.8% | [+51.2%, +52.4%] | 0.6 pts | 1.5 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 6 | 9.41 ms | 5.95 ms | 0.63× [0.63, 0.64] | -58.2% | [-59.0%, -57.3%] | 0.8 pts | 2.3 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 6 | 9.41 ms | 15.13 ms | 1.61× [1.60, 1.62] | +37.7% | [+37.3%, +38.1%] | 0.4 pts | 0.9 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 8 | 49.6 | 47.7 | 0.95× [0.94, 0.97] | -4.8% | [-6.8%, -2.9%] | 2.1 pts | 3.9 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 8 | 49.2 | 195 | 3.97× [3.96, 3.99] | +74.8% | [+74.7%, +74.9%] | 0.2 pts | 1.3 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 8 | 95.0 | 79.4 | 0.84× [0.83, 0.84] | -19.6% | [-20.0%, -19.2%] | 0.6 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 8 | 97.1 | 232 | 2.38× [2.36, 2.40] | +58.0% | [+57.7%, +58.4%] | 0.4 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 8 | 35.94 ms | 29.51 ms | 0.82× [0.82, 0.83] | -21.5% | [-22.6%, -20.4%] | 1.3 pts | 1.1 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 8 | 37.20 ms | 81.82 ms | 2.19× [2.16, 2.22] | +54.4% | [+53.7%, +55.1%] | 0.7 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 churn: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
