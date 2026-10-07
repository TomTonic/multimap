| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 112 | 97.7 | 0.87× [0.86, 0.88] | -15.2% | [-16.4%, -14.1%] | 1.2 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 116 | 112 | 0.96× [0.94, 0.99] | -3.9% | [-6.9%, -1.0%] | 3.2 pts | 1.9 | no | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 3485 | 4211 | 1.21× [1.20, 1.22] | +17.4% | [+16.8%, +17.9%] | 0.6 pts | 0.5 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 3418 | 1585 | 0.46× [0.45, 0.47] | -117.0% | [-120.1%, -113.9%] | 4.0 pts | 1.5 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 4116 | 4900 | 1.20× [1.19, 1.22] | +16.8% | [+15.9%, +17.7%] | 0.9 pts | 0.4 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 4038 | 1960 | 0.48× [0.47, 0.50] | -106.4% | [-111.3%, -101.5%] | 6.3 pts | 1.3 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 295 | 213 | 0.72× [0.71, 0.73] | -39.1% | [-41.5%, -36.7%] | 2.6 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 293 | 189 | 0.65× [0.64, 0.65] | -54.7% | [-56.4%, -52.9%] | 1.8 pts | 0.8 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 4.58 ms | 3.15 ms | 0.70× [0.68, 0.71] | -43.8% | [-46.1%, -41.5%] | 2.4 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 4.73 ms | 2.77 ms | 0.59× [0.57, 0.61] | -69.1% | [-74.7%, -63.5%] | 5.4 pts | 1.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 137 | 128 | 0.93× [0.92, 0.94] | -7.5% | [-8.6%, -6.5%] | 1.2 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 139 | 165 | 1.19× [1.16, 1.23] | +16.3% | [+13.9%, +18.6%] | 3.7 pts | 3.5 | no | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3791 | 4869 | 1.29× [1.28, 1.30] | +22.2% | [+21.6%, +22.8%] | 0.6 pts | 1.3 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 3771 | 2070 | 0.56× [0.55, 0.57] | -79.5% | [-83.1%, -75.9%] | 10.7 pts | 6.9 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 8 | 17.9 µs | 25.4 µs | 1.40× [1.39, 1.40] | +28.4% | [+27.9%, +28.8%] | 0.7 pts | 0.7 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 8 | 17.2 µs | 9178 | 0.53× [0.51, 0.55] | -88.3% | [-95.0%, -81.5%] | 10.4 pts | 0.7 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 8 | 358 | 280 | 0.78× [0.78, 0.79] | -27.6% | [-28.8%, -26.5%] | 1.2 pts | 0.8 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 8 | 360 | 263 | 0.75× [0.73, 0.78] | -33.0% | [-37.0%, -29.0%] | 3.9 pts | 2.6 | no | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 8 | 21.48 ms | 15.76 ms | 0.75× [0.73, 0.77] | -33.6% | [-37.7%, -29.5%] | 4.0 pts | 2.5 | no | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 8 | 21.75 ms | 14.96 ms | 0.69× [0.67, 0.71] | -45.4% | [-49.6%, -41.2%] | 3.9 pts | 2.0 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 184 | 193 | 1.06× [1.02, 1.11] | +5.6% | [+1.5%, +9.6%] | 5.9 pts | 4.8 | no | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 203 | 274 | 1.44× [1.32, 1.60] | +30.8% | [+24.1%, +37.5%] | 8.7 pts | 6.9 | no | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 4333 | 6176 | 1.42× [1.34, 1.51] | +29.5% | [+25.3%, +33.8%] | 5.4 pts | 9.3 | no | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 4174 | 3015 | 0.74× [0.66, 0.86] | -34.3% | [-52.0%, -16.5%] | 22.6 pts | 13.1 | no | yes |
| single-value-str | dirs | 86215 | prefix | ordered | baseline | 8 | 103.1 µs | 159.0 µs | 1.54× [1.49, 1.60] | +35.0% | [+32.7%, +37.4%] | 3.5 pts | 6.1 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | btree-map | 8 | 103.2 µs | 63.3 µs | 0.66× [0.57, 0.78] | -52.6% | [-76.1%, -29.0%] | 29.5 pts | 9.0 | no | yes |
| single-value-str | dirs | 86215 | churn | ordered | baseline | 8 | 508 | 474 | 0.92× [0.91, 0.94] | -8.2% | [-10.4%, -6.1%] | 3.4 pts | 2.6 | no | yes |
| single-value-str | dirs | 86215 | churn | ordered | btree-map | 8 | 547 | 482 | 0.89× [0.86, 0.91] | -12.7% | [-15.7%, -9.7%] | 2.9 pts | 2.6 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str dirs n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 6.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 churn: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 build: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 build: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: 6 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 6.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 9.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 13.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 prefix: ordered vs baseline: the processes scatter 6.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 prefix: ordered vs btree-map: the processes scatter 9.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 churn: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
