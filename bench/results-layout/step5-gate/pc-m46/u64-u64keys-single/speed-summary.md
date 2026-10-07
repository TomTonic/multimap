| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 6 | 45.7 | 20.0 | 0.43× [0.42, 0.45] | -130.1% | [-137.8%, -122.3%] | 7.4 pts | 10.4 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 45.4 | 88.2 | 1.94× [1.94, 1.95] | +48.6% | [+48.4%, +48.8%] | 0.2 pts | 1.6 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 910 | 1279 | 1.40× [1.40, 1.41] | +28.6% | [+28.3%, +28.9%] | 0.3 pts | 1.0 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 914 | 446 | 0.49× [0.48, 0.49] | -105.3% | [-107.2%, -103.3%] | 1.8 pts | 2.6 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 6 | 78.4 | 58.8 | 0.75× [0.74, 0.76] | -33.8% | [-35.3%, -32.3%] | 1.4 pts | 2.1 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 6 | 78.6 | 123 | 1.56× [1.55, 1.56] | +35.8% | [+35.5%, +36.1%] | 0.3 pts | 0.7 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 6 | 1.17 ms | 1.04 ms | 0.88× [0.86, 0.89] | -14.0% | [-15.8%, -12.1%] | 1.8 pts | 0.9 | yes | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 6 | 1.16 ms | 1.78 ms | 1.52× [1.50, 1.53] | +34.1% | [+33.5%, +34.8%] | 0.6 pts | 0.6 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 8 | 33.7 | 25.7 | 0.76× [0.74, 0.77] | -32.1% | [-35.2%, -29.1%] | 3.1 pts | 4.0 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 33.5 | 112 | 3.31× [3.22, 3.39] | +69.8% | [+69.0%, +70.5%] | 0.8 pts | 5.6 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 2141 | 1800 | 0.84× [0.83, 0.86] | -18.8% | [-20.6%, -16.9%] | 1.8 pts | 1.3 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 2098 | 476 | 0.23× [0.22, 0.23] | -343.5% | [-346.1%, -341.0%] | 2.8 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 8 | 74.5 | 68.9 | 0.92× [0.91, 0.94] | -8.4% | [-9.9%, -6.9%] | 1.4 pts | 2.0 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 8 | 77.5 | 170 | 2.21× [2.13, 2.29] | +54.7% | [+53.0%, +56.4%] | 1.9 pts | 3.7 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 8 | 5.53 ms | 4.53 ms | 0.83× [0.82, 0.83] | -21.2% | [-21.8%, -20.5%] | 0.6 pts | 0.7 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 8 | 5.60 ms | 9.07 ms | 1.60× [1.55, 1.66] | +37.6% | [+35.5%, +39.8%] | 2.1 pts | 1.7 | yes | yes |
| single-value | u64 | 262144 | valuesFor | ordered | baseline | 8 | 74.4 | 60.4 | 0.79× [0.70, 0.90] | -26.9% | [-42.3%, -11.5%] | 18.9 pts | 16.0 | no | yes |
| single-value | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 88.3 | 234 | 2.72× [2.38, 3.18] | +63.3% | [+58.0%, +68.6%] | 6.0 pts | 8.2 | yes | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 1347 | 2334 | 1.58× [1.40, 1.81] | +36.6% | [+28.6%, +44.6%] | 10.3 pts | 21.3 | no | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 1336 | 852 | 0.66× [0.60, 0.73] | -52.2% | [-67.0%, -37.3%] | 18.4 pts | 18.4 | no | yes |
| single-value | u64 | 262144 | churn | ordered | baseline | 8 | 187 | 212 | 1.13× [1.10, 1.17] | +11.7% | [+8.9%, +14.5%] | 2.9 pts | 1.6 | no | yes |
| single-value | u64 | 262144 | churn | ordered | btree-map | 8 | 240 | 438 | 1.78× [1.68, 1.89] | +43.8% | [+40.5%, +47.0%] | 4.0 pts | 3.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 10.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=16384 churn: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 16.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesFor: ordered vs btree-map: the processes scatter 8.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 21.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesBetween: ordered vs btree-map: the processes scatter 18.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.21% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=262144 churn: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
