| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 97.7 | 76.9 | 0.78× [0.77, 0.79] | -27.7% | [-29.1%, -26.3%] | 1.9 pts | 1.7 | yes | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 103 | 110 | 1.08× [1.04, 1.12] | +7.0% | [+3.6%, +10.4%] | 4.2 pts | 2.6 | no | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 4066 | 2124 | 0.52× [0.51, 0.53] | -92.9% | [-95.5%, -90.4%] | 3.2 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 4090 | 1521 | 0.37× [0.36, 0.37] | -173.7% | [-178.7%, -168.6%] | 5.4 pts | 1.2 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 4713 | 2332 | 0.49× [0.48, 0.50] | -103.7% | [-106.3%, -101.2%] | 2.7 pts | 0.6 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 4661 | 1881 | 0.40× [0.39, 0.41] | -149.0% | [-154.0%, -144.1%] | 5.9 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 204 | 189 | 0.93× [0.92, 0.94] | -7.7% | [-9.1%, -6.4%] | 1.3 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 203 | 180 | 0.88× [0.87, 0.89] | -14.0% | [-15.5%, -12.6%] | 2.0 pts | 1.3 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 3.23 ms | 2.71 ms | 0.85× [0.84, 0.86] | -18.2% | [-19.6%, -16.8%] | 1.4 pts | 0.9 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 3.22 ms | 2.65 ms | 0.83× [0.82, 0.84] | -20.3% | [-21.7%, -18.9%] | 1.7 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 127 | 107 | 0.84× [0.81, 0.87] | -19.4% | [-23.3%, -15.6%] | 3.6 pts | 3.0 | no | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 134 | 161 | 1.21× [1.17, 1.27] | +17.7% | [+14.2%, +21.2%] | 3.6 pts | 3.8 | no | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 4690 | 2681 | 0.57× [0.57, 0.57] | -75.2% | [-76.1%, -74.3%] | 0.9 pts | 0.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 4694 | 1998 | 0.42× [0.42, 0.43] | -135.7% | [-137.5%, -133.9%] | 1.8 pts | 1.0 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 8 | 23.6 µs | 10.7 µs | 0.45× [0.44, 0.47] | -120.6% | [-126.7%, -114.4%] | 10.5 pts | 0.8 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 8 | 23.5 µs | 8299 | 0.37× [0.36, 0.39] | -167.2% | [-177.6%, -156.8%] | 10.3 pts | 0.5 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 8 | 264 | 247 | 0.93× [0.93, 0.94] | -7.2% | [-7.9%, -6.5%] | 0.9 pts | 0.8 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 8 | 262 | 245 | 0.93× [0.93, 0.94] | -7.2% | [-8.0%, -6.4%] | 0.8 pts | 0.7 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 8 | 15.40 ms | 13.25 ms | 0.87× [0.85, 0.88] | -15.3% | [-17.2%, -13.4%] | 2.0 pts | 2.0 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 8 | 15.32 ms | 14.03 ms | 0.92× [0.91, 0.94] | -8.4% | [-10.0%, -6.9%] | 1.6 pts | 1.1 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 204 | 178 | 0.89× [0.87, 0.92] | -11.9% | [-15.4%, -8.4%] | 3.3 pts | 2.1 | no | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 231 | 262 | 1.14× [1.07, 1.22] | +12.3% | [+6.6%, +17.9%] | 5.3 pts | 2.9 | no | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 5618 | 3449 | 0.61× [0.58, 0.65] | -63.3% | [-72.6%, -53.9%] | 11.0 pts | 4.2 | no | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 5881 | 3156 | 0.53× [0.50, 0.56] | -89.2% | [-99.7%, -78.6%] | 12.0 pts | 2.8 | no | yes |
| single-value-str | dirs | 86215 | prefix | ordered | baseline | 8 | 153.7 µs | 86.9 µs | 0.56× [0.55, 0.58] | -77.8% | [-83.2%, -72.5%] | 5.3 pts | 0.8 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | btree-map | 8 | 155.2 µs | 66.5 µs | 0.40× [0.38, 0.42] | -149.2% | [-160.2%, -138.2%] | 10.3 pts | 0.8 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | baseline | 8 | 441 | 416 | 0.93× [0.91, 0.95] | -7.7% | [-9.8%, -5.6%] | 2.2 pts | 1.3 | no | yes |
| single-value-str | dirs | 86215 | churn | ordered | btree-map | 8 | 493 | 447 | 0.92× [0.90, 0.94] | -8.3% | [-10.5%, -6.0%] | 3.0 pts | 2.1 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -1.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=86215 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
