| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 90.9 | 79.1 | 0.87× [0.86, 0.89] | -14.7% | [-16.8%, -12.7%] | 2.2 pts | 2.3 | no | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 92.0 | 110 | 1.20× [1.18, 1.22] | +16.7% | [+15.2%, +18.2%] | 1.7 pts | 1.5 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1386 | 2328 | 1.67× [1.65, 1.69] | +40.2% | [+39.5%, +41.0%] | 0.8 pts | 0.7 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1357 | 544 | 0.40× [0.40, 0.40] | -149.6% | [-152.1%, -147.1%] | 2.5 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 8 | 1527 | 2698 | 1.74× [1.73, 1.75] | +42.5% | [+42.2%, +42.8%] | 0.7 pts | 0.6 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | btree-map | 8 | 1537 | 657 | 0.42× [0.42, 0.43] | -135.3% | [-137.6%, -133.1%] | 2.2 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 1533 | 219 | 0.14× [0.14, 0.14] | -628.0% | [-636.0%, -620.0%] | 8.9 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 1520 | 187 | 0.12× [0.12, 0.12] | -738.7% | [-747.6%, -729.9%] | 8.6 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 16.90 ms | 2.93 ms | 0.18× [0.17, 0.18] | -469.4% | [-474.8%, -464.0%] | 9.4 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 17.03 ms | 2.65 ms | 0.16× [0.16, 0.16] | -531.2% | [-536.4%, -525.9%] | 7.8 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 6 | 113 | 108 | 0.97× [0.95, 0.98] | -3.4% | [-5.1%, -1.7%] | 1.6 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 6 | 114 | 159 | 1.39× [1.37, 1.41] | +27.9% | [+27.0%, +28.8%] | 0.9 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 1583 | 2849 | 1.80× [1.77, 1.83] | +44.5% | [+43.6%, +45.3%] | 0.8 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 6 | 1564 | 713 | 0.45× [0.44, 0.46] | -121.5% | [-125.6%, -117.4%] | 3.9 pts | 1.7 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 6 | 5267 | 11.6 µs | 2.14× [2.07, 2.22] | +53.3% | [+51.6%, +55.0%] | 1.6 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | btree-map | 6 | 5222 | 2420 | 0.46× [0.45, 0.48] | -116.5% | [-122.7%, -110.2%] | 6.0 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 6 | 1644 | 305 | 0.19× [0.18, 0.19] | -434.9% | [-448.7%, -421.2%] | 13.1 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 6 | 1663 | 279 | 0.17× [0.17, 0.18] | -484.1% | [-498.5%, -469.6%] | 13.8 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 6 | 79.18 ms | 14.57 ms | 0.19× [0.18, 0.19] | -433.4% | [-449.3%, -417.4%] | 15.2 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 6 | 78.94 ms | 14.34 ms | 0.18× [0.18, 0.19] | -445.0% | [-456.6%, -433.5%] | 11.0 pts | 0.6 | yes | yes |
| single-value | dirs | 86215 | valuesFor | ordered | baseline | 8 | 150 | 155 | 1.05× [1.01, 1.09] | +4.9% | [+1.1%, +8.7%] | 4.2 pts | 3.3 | no | yes |
| single-value | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 155 | 236 | 1.49× [1.45, 1.54] | +33.1% | [+30.9%, +35.2%] | 2.6 pts | 2.5 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 1754 | 3133 | 1.78× [1.77, 1.80] | +44.0% | [+43.7%, +44.3%] | 0.4 pts | 1.0 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 1740 | 938 | 0.54× [0.53, 0.54] | -86.6% | [-89.1%, -84.2%] | 2.5 pts | 1.3 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | baseline | 8 | 32.6 µs | 78.4 µs | 2.44× [2.40, 2.48] | +59.0% | [+58.3%, +59.7%] | 0.7 pts | 1.8 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | btree-map | 8 | 31.6 µs | 14.1 µs | 0.45× [0.45, 0.45] | -122.6% | [-123.2%, -122.0%] | 0.7 pts | 0.3 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | baseline | 8 | 1939 | 545 | 0.28× [0.28, 0.29] | -255.1% | [-259.7%, -250.5%] | 4.6 pts | 0.7 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | btree-map | 8 | 1942 | 564 | 0.30× [0.29, 0.30] | -237.8% | [-246.3%, -229.3%] | 9.1 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -1.11% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -1.93% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -1.80% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of -3.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -4.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
