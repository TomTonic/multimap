| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | u64 | 4096 | valuesFor | ordered | baseline | 6 | 20.7 | 18.1 | 0.87× [0.87, 0.88] | -14.5% | [-15.3%, -13.8%] | 0.7 pts | 2.2 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 20.7 | 88.4 | 4.28× [4.24, 4.32] | +76.6% | [+76.4%, +76.9%] | 0.2 pts | 3.0 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1321 | 1070 | 0.81× [0.80, 0.82] | -23.3% | [-24.4%, -22.1%] | 1.1 pts | 1.1 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 1313 | 443 | 0.34× [0.33, 0.34] | -197.6% | [-200.5%, -194.6%] | 2.8 pts | 1.8 | yes | yes |
| single-value-ptr | u64 | 4096 | churn | ordered | baseline | 6 | 59.4 | 49.1 | 0.83× [0.81, 0.84] | -21.0% | [-22.8%, -19.2%] | 1.7 pts | 2.8 | yes | yes |
| single-value-ptr | u64 | 4096 | churn | ordered | btree-map | 6 | 58.6 | 121 | 2.06× [2.04, 2.08] | +51.4% | [+50.9%, +51.9%] | 0.5 pts | 1.4 | yes | yes |
| single-value-ptr | u64 | 4096 | build | ordered | baseline | 6 | 1.01 ms | 809.9 µs | 0.80× [0.80, 0.81] | -24.9% | [-25.7%, -24.1%] | 0.8 pts | 0.9 | yes | yes |
| single-value-ptr | u64 | 4096 | build | ordered | btree-map | 6 | 1.01 ms | 1.72 ms | 1.72× [1.70, 1.75] | +41.9% | [+41.1%, +42.8%] | 0.8 pts | 2.0 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesFor | ordered | baseline | 8 | 26.6 | 23.5 | 0.87× [0.86, 0.89] | -14.4% | [-16.5%, -12.3%] | 2.1 pts | 4.6 | no | yes |
| single-value-ptr | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 26.9 | 112 | 4.12× [3.93, 4.33] | +75.7% | [+74.5%, +76.9%] | 1.2 pts | 11.1 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1819 | 1612 | 0.88× [0.87, 0.90] | -13.0% | [-14.4%, -11.6%] | 1.4 pts | 0.8 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 1790 | 488 | 0.27× [0.27, 0.28] | -266.3% | [-271.0%, -261.5%] | 4.6 pts | 2.0 | yes | yes |
| single-value-ptr | u64 | 16384 | churn | ordered | baseline | 8 | 66.3 | 56.6 | 0.85× [0.84, 0.86] | -17.3% | [-18.4%, -16.2%] | 1.2 pts | 1.0 | yes | yes |
| single-value-ptr | u64 | 16384 | churn | ordered | btree-map | 8 | 67.4 | 166 | 2.46× [2.39, 2.52] | +59.3% | [+58.2%, +60.4%] | 1.0 pts | 2.5 | yes | yes |
| single-value-ptr | u64 | 16384 | build | ordered | baseline | 8 | 4.51 ms | 3.71 ms | 0.83× [0.82, 0.83] | -21.2% | [-21.4%, -21.0%] | 0.8 pts | 1.0 | yes | yes |
| single-value-ptr | u64 | 16384 | build | ordered | btree-map | 8 | 4.49 ms | 8.91 ms | 1.98× [1.96, 2.00] | +49.5% | [+49.0%, +49.9%] | 0.5 pts | 1.6 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesFor | ordered | baseline | 8 | 77.8 | 62.7 | 0.81× [0.79, 0.84] | -22.8% | [-26.4%, -19.2%] | 4.9 pts | 3.1 | no | yes |
| single-value-ptr | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 95.0 | 218 | 2.29× [2.26, 2.32] | +56.4% | [+55.8%, +56.9%] | 0.9 pts | 0.8 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 2580 | 2063 | 0.83× [0.80, 0.86] | -20.9% | [-24.9%, -16.8%] | 4.2 pts | 2.3 | no | yes |
| single-value-ptr | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 2282 | 962 | 0.42× [0.41, 0.44] | -137.0% | [-146.5%, -127.5%] | 10.4 pts | 1.1 | yes | yes |
| single-value-ptr | u64 | 262144 | churn | ordered | baseline | 8 | 194 | 181 | 0.91× [0.90, 0.93] | -9.6% | [-11.6%, -7.7%] | 1.9 pts | 0.6 | yes | yes |
| single-value-ptr | u64 | 262144 | churn | ordered | btree-map | 8 | 273 | 379 | 1.41× [1.37, 1.45] | +29.0% | [+27.2%, +30.8%] | 1.7 pts | 1.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr u64 n=4096 churn: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.21% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr u64 n=4096 build: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 11.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 churn: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of -0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.48% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
