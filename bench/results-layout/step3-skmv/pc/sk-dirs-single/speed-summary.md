| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 97.1 | 76.9 | 0.79× [0.77, 0.81] | -26.4% | [-29.4%, -23.4%] | 3.3 pts | 2.7 | no | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 104 | 110 | 1.06× [1.03, 1.10] | +5.9% | [+2.5%, +9.2%] | 3.7 pts | 2.3 | no | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 4064 | 2078 | 0.51× [0.50, 0.51] | -96.4% | [-98.2%, -94.6%] | 1.7 pts | 0.5 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 4052 | 1509 | 0.37× [0.36, 0.37] | -173.8% | [-175.5%, -172.0%] | 2.4 pts | 0.5 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 4746 | 2305 | 0.48× [0.47, 0.49] | -108.2% | [-112.4%, -104.0%] | 4.1 pts | 0.7 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 4817 | 1845 | 0.39× [0.38, 0.41] | -153.3% | [-161.4%, -145.1%] | 8.0 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 202 | 190 | 0.95× [0.94, 0.96] | -5.4% | [-6.2%, -4.7%] | 1.2 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 201 | 180 | 0.89× [0.89, 0.90] | -11.8% | [-12.6%, -10.9%] | 1.4 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 3.20 ms | 2.69 ms | 0.85× [0.84, 0.86] | -18.1% | [-19.7%, -16.5%] | 2.2 pts | 1.5 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 3.22 ms | 2.65 ms | 0.82× [0.81, 0.83] | -21.9% | [-23.1%, -20.7%] | 1.3 pts | 0.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 128 | 107 | 0.82× [0.80, 0.85] | -21.2% | [-24.7%, -17.8%] | 3.8 pts | 3.4 | no | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 131 | 162 | 1.23× [1.20, 1.27] | +18.8% | [+16.4%, +21.1%] | 2.9 pts | 2.9 | no | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 4712 | 2677 | 0.57× [0.56, 0.57] | -76.9% | [-78.2%, -75.6%] | 2.1 pts | 1.6 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 4720 | 1976 | 0.42× [0.42, 0.42] | -139.1% | [-140.7%, -137.5%] | 1.6 pts | 0.9 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 8 | 22.7 µs | 11.3 µs | 0.48× [0.47, 0.50] | -107.4% | [-114.5%, -100.3%] | 7.5 pts | 0.5 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 8 | 22.5 µs | 8388 | 0.37× [0.35, 0.38] | -173.0% | [-182.4%, -163.7%] | 15.6 pts | 0.9 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 8 | 262 | 248 | 0.94× [0.93, 0.95] | -6.7% | [-7.8%, -5.5%] | 1.1 pts | 1.0 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 8 | 261 | 244 | 0.94× [0.91, 0.97] | -6.3% | [-9.9%, -2.6%] | 3.7 pts | 2.1 | no | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 8 | 14.86 ms | 13.04 ms | 0.88× [0.87, 0.89] | -14.0% | [-15.0%, -13.0%] | 1.0 pts | 0.9 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 8 | 14.84 ms | 13.54 ms | 0.92× [0.90, 0.93] | -9.3% | [-10.8%, -7.7%] | 1.7 pts | 0.9 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 202 | 175 | 0.88× [0.86, 0.91] | -13.1% | [-16.6%, -9.5%] | 3.5 pts | 2.4 | no | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 221 | 249 | 1.15× [1.09, 1.21] | +13.0% | [+8.7%, +17.4%] | 5.2 pts | 3.8 | no | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 5424 | 3230 | 0.61× [0.59, 0.63] | -65.1% | [-70.3%, -60.0%] | 6.1 pts | 3.3 | yes | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 5623 | 2646 | 0.48× [0.45, 0.51] | -109.0% | [-122.8%, -95.1%] | 14.7 pts | 2.7 | no | yes |
| single-value-str | dirs | 86215 | prefix | ordered | baseline | 8 | 151.7 µs | 86.8 µs | 0.56× [0.55, 0.58] | -77.8% | [-82.1%, -73.6%] | 4.7 pts | 0.9 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | btree-map | 8 | 151.3 µs | 59.1 µs | 0.39× [0.38, 0.41] | -155.7% | [-164.5%, -146.8%] | 9.5 pts | 0.9 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | baseline | 8 | 417 | 387 | 0.94× [0.93, 0.95] | -6.8% | [-7.8%, -5.7%] | 1.0 pts | 0.6 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | btree-map | 8 | 486 | 432 | 0.89× [0.88, 0.90] | -12.6% | [-13.7%, -11.5%] | 1.1 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of -1.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -3.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
