| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 63.8 | 50.2 | 0.78× [0.77, 0.79] | -27.9% | [-29.8%, -26.1%] | 1.9 pts | 3.0 | yes | yes |
| single-value-ptr | street | 4096 | valuesFor | ordered | btree-map | 8 | 64.4 | 91.3 | 1.42× [1.41, 1.43] | +29.4% | [+29.0%, +29.9%] | 0.7 pts | 1.4 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 1137 | 2002 | 1.77× [1.76, 1.78] | +43.5% | [+43.2%, +43.7%] | 0.2 pts | 0.4 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1120 | 501 | 0.44× [0.44, 0.45] | -126.3% | [-129.6%, -122.9%] | 3.3 pts | 1.4 | yes | yes |
| single-value-ptr | street | 4096 | prefix | ordered | baseline | 8 | 221 | 226 | 1.03× [1.02, 1.04] | +2.7% | [+1.9%, +3.4%] | 0.7 pts | 1.0 | yes | yes |
| single-value-ptr | street | 4096 | prefix | ordered | btree-map | 8 | 220 | 133 | 0.60× [0.60, 0.61] | -65.5% | [-67.0%, -64.0%] | 1.5 pts | 1.2 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | baseline | 8 | 271 | 135 | 0.49× [0.48, 0.50] | -103.7% | [-107.5%, -99.9%] | 4.3 pts | 1.1 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | btree-map | 8 | 269 | 140 | 0.52× [0.52, 0.53] | -91.2% | [-92.9%, -89.5%] | 1.9 pts | 0.9 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | baseline | 8 | 3.59 ms | 1.90 ms | 0.53× [0.53, 0.54] | -88.1% | [-90.0%, -86.2%] | 2.0 pts | 1.3 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | btree-map | 8 | 3.58 ms | 2.17 ms | 0.61× [0.58, 0.63] | -64.6% | [-71.3%, -57.8%] | 7.9 pts | 5.4 | no | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | baseline | 8 | 76.4 | 69.5 | 0.91× [0.89, 0.92] | -10.4% | [-12.3%, -8.5%] | 1.9 pts | 2.2 | yes | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | btree-map | 8 | 77.1 | 127 | 1.64× [1.60, 1.68] | +39.1% | [+37.7%, +40.5%] | 1.9 pts | 3.8 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | baseline | 8 | 1309 | 2443 | 1.86× [1.85, 1.87] | +46.2% | [+45.8%, +46.6%] | 0.4 pts | 1.1 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1302 | 632 | 0.48× [0.47, 0.49] | -106.9% | [-111.5%, -102.3%] | 5.4 pts | 3.8 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | baseline | 8 | 464 | 733 | 1.58× [1.57, 1.60] | +36.8% | [+36.4%, +37.3%] | 0.6 pts | 1.0 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | btree-map | 8 | 462 | 290 | 0.62× [0.61, 0.63] | -60.8% | [-64.0%, -57.7%] | 3.9 pts | 2.6 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | baseline | 8 | 290 | 174 | 0.60× [0.59, 0.61] | -67.6% | [-70.4%, -64.7%] | 2.8 pts | 1.1 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | btree-map | 8 | 291 | 193 | 0.66× [0.65, 0.68] | -50.5% | [-53.9%, -47.2%] | 3.5 pts | 2.1 | yes | yes |
| single-value-ptr | street | 16384 | build | ordered | baseline | 8 | 16.49 ms | 9.37 ms | 0.57× [0.57, 0.58] | -74.0% | [-75.5%, -72.6%] | 1.5 pts | 0.4 | yes | yes |
| single-value-ptr | street | 16384 | build | ordered | btree-map | 8 | 16.79 ms | 11.14 ms | 0.67× [0.65, 0.69] | -49.3% | [-54.5%, -44.0%] | 5.0 pts | 1.1 | no | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 127 | 147 | 1.19× [1.09, 1.32] | +16.0% | [+7.9%, +24.1%] | 7.6 pts | 7.4 | no | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | btree-map | 8 | 137 | 230 | 1.66× [1.59, 1.75] | +39.9% | [+37.0%, +42.8%] | 3.0 pts | 3.6 | yes | yes |
| single-value-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 1609 | 3124 | 1.94× [1.86, 2.02] | +48.4% | [+46.3%, +50.4%] | 2.1 pts | 3.1 | yes | yes |
| single-value-ptr | street | 212449 | valuesBetween | ordered | btree-map | 8 | 1563 | 970 | 0.62× [0.60, 0.65] | -61.0% | [-67.7%, -54.4%] | 8.0 pts | 3.1 | no | yes |
| single-value-ptr | street | 212449 | prefix | ordered | baseline | 8 | 5200 | 13.1 µs | 2.45× [2.39, 2.51] | +59.2% | [+58.2%, +60.1%] | 0.9 pts | 1.2 | yes | yes |
| single-value-ptr | street | 212449 | prefix | ordered | btree-map | 8 | 5047 | 3166 | 0.63× [0.61, 0.65] | -59.9% | [-64.8%, -55.0%] | 4.8 pts | 1.3 | yes | yes |
| single-value-ptr | street | 212449 | churn | ordered | baseline | 8 | 526 | 474 | 0.91× [0.88, 0.93] | -10.3% | [-13.1%, -7.6%] | 2.6 pts | 1.3 | no | yes |
| single-value-ptr | street | 212449 | churn | ordered | btree-map | 8 | 582 | 470 | 0.82× [0.82, 0.82] | -22.1% | [-22.7%, -21.4%] | 1.3 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr street n=4096 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 build: ordered vs btree-map: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 prefix: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 7.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesFor: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
