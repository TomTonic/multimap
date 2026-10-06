| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 6 | 37.0 | 20.1 | 0.55× [0.54, 0.55] | -83.3% | [-84.7%, -82.0%] | 1.3 pts | 3.1 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 36.8 | 87.9 | 2.39× [2.37, 2.41] | +58.1% | [+57.8%, +58.5%] | 0.3 pts | 3.0 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 902 | 1294 | 1.43× [1.43, 1.44] | +30.2% | [+29.9%, +30.6%] | 0.3 pts | 1.2 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 905 | 423 | 0.47× [0.47, 0.47] | -114.1% | [-114.4%, -113.8%] | 0.3 pts | 0.6 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 6 | 74.9 | 57.4 | 0.76× [0.76, 0.77] | -31.0% | [-32.4%, -29.6%] | 1.4 pts | 2.6 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 6 | 74.5 | 121 | 1.62× [1.60, 1.64] | +38.3% | [+37.6%, +39.0%] | 0.7 pts | 1.5 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 6 | 1.10 ms | 987.3 µs | 0.89× [0.88, 0.89] | -12.7% | [-13.1%, -12.3%] | 0.4 pts | 0.4 | yes | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 6 | 1.10 ms | 1.74 ms | 1.57× [1.56, 1.58] | +36.3% | [+35.8%, +36.8%] | 0.5 pts | 1.7 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 8 | 27.8 | 26.0 | 0.93× [0.92, 0.95] | -7.0% | [-9.0%, -5.0%] | 2.2 pts | 3.9 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 27.6 | 110 | 3.98× [3.92, 4.04] | +74.9% | [+74.5%, +75.3%] | 0.4 pts | 4.4 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1630 | 1769 | 1.09× [1.08, 1.10] | +8.4% | [+7.4%, +9.3%] | 1.0 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 1602 | 450 | 0.28× [0.28, 0.28] | -255.4% | [-256.5%, -254.2%] | 2.0 pts | 1.0 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 8 | 82.4 | 64.6 | 0.79× [0.78, 0.79] | -27.1% | [-27.6%, -26.6%] | 1.2 pts | 1.2 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 8 | 83.7 | 167 | 1.99× [1.96, 2.02] | +49.8% | [+49.1%, +50.6%] | 1.3 pts | 3.7 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 8 | 5.80 ms | 4.36 ms | 0.75× [0.75, 0.75] | -32.6% | [-32.8%, -32.5%] | 0.3 pts | 0.4 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 8 | 5.80 ms | 8.92 ms | 1.54× [1.52, 1.55] | +34.9% | [+34.4%, +35.5%] | 0.6 pts | 1.5 | yes | yes |
| single-value | u64 | 262144 | valuesFor | ordered | baseline | 8 | 53.4 | 50.8 | 0.95× [0.92, 0.97] | -5.6% | [-8.6%, -2.6%] | 4.4 pts | 2.6 | no | yes |
| single-value | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 67.9 | 210 | 3.02× [2.67, 3.47] | +66.9% | [+62.5%, +71.2%] | 5.1 pts | 3.1 | yes | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 1246 | 1724 | 1.38× [1.37, 1.40] | +27.8% | [+26.8%, +28.8%] | 1.3 pts | 2.7 | yes | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 1228 | 731 | 0.59× [0.58, 0.61] | -68.2% | [-71.7%, -64.6%] | 3.4 pts | 3.0 | yes | yes |
| single-value | u64 | 262144 | churn | ordered | baseline | 8 | 165 | 193 | 1.13× [1.10, 1.16] | +11.4% | [+9.1%, +13.8%] | 2.3 pts | 0.8 | no | yes |
| single-value | u64 | 262144 | churn | ordered | btree-map | 8 | 241 | 381 | 1.61× [1.58, 1.64] | +37.9% | [+36.8%, +39.1%] | 1.1 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesFor: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesBetween: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
