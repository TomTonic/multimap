| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 6 | 90.6 | 78.3 | 0.87× [0.86, 0.88] | -15.3% | [-16.4%, -14.2%] | 1.0 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 6 | 92.1 | 108 | 1.16× [1.14, 1.19] | +14.0% | [+12.3%, +15.7%] | 1.6 pts | 1.7 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 1364 | 2335 | 1.71× [1.66, 1.76] | +41.5% | [+39.8%, +43.3%] | 1.7 pts | 1.6 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 6 | 1341 | 538 | 0.40× [0.39, 0.41] | -149.3% | [-154.2%, -144.3%] | 4.7 pts | 1.6 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 6 | 1603 | 2709 | 1.68× [1.65, 1.71] | +40.5% | [+39.5%, +41.4%] | 0.9 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | btree-map | 6 | 1607 | 644 | 0.40× [0.40, 0.40] | -150.1% | [-152.6%, -147.7%] | 2.4 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 6 | 252 | 198 | 0.79× [0.78, 0.80] | -27.3% | [-28.8%, -25.7%] | 1.5 pts | 1.2 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 6 | 252 | 176 | 0.69× [0.68, 0.71] | -44.9% | [-48.0%, -41.7%] | 3.0 pts | 1.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 6 | 3.63 ms | 2.90 ms | 0.80× [0.79, 0.81] | -24.7% | [-25.9%, -23.6%] | 1.1 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 6 | 3.67 ms | 2.55 ms | 0.70× [0.69, 0.71] | -43.0% | [-44.2%, -41.8%] | 1.1 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 112 | 106 | 0.95× [0.94, 0.96] | -5.1% | [-5.9%, -4.2%] | 0.9 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 114 | 157 | 1.38× [1.37, 1.40] | +27.6% | [+26.8%, +28.5%] | 0.9 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1517 | 2845 | 1.87× [1.87, 1.88] | +46.6% | [+46.4%, +46.8%] | 0.4 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 1508 | 701 | 0.47× [0.46, 0.48] | -113.8% | [-117.7%, -109.9%] | 4.1 pts | 2.1 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 8 | 5419 | 11.5 µs | 2.11× [2.03, 2.20] | +52.6% | [+50.6%, +54.5%] | 2.5 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | btree-map | 8 | 5400 | 2359 | 0.43× [0.43, 0.43] | -132.1% | [-133.9%, -130.3%] | 3.7 pts | 0.5 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 280 | 247 | 0.89× [0.87, 0.92] | -12.1% | [-14.9%, -9.2%] | 2.7 pts | 2.6 | no | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 281 | 237 | 0.85× [0.84, 0.86] | -17.3% | [-18.7%, -15.9%] | 1.6 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 16.12 ms | 13.56 ms | 0.84× [0.84, 0.85] | -18.4% | [-19.6%, -17.2%] | 1.3 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 16.08 ms | 13.06 ms | 0.81× [0.80, 0.82] | -23.0% | [-24.3%, -21.7%] | 1.3 pts | 0.9 | yes | yes |
| single-value | dirs | 86215 | valuesFor | ordered | baseline | 8 | 150 | 155 | 1.03× [1.00, 1.07] | +3.2% | [+0.2%, +6.3%] | 3.0 pts | 3.0 | no | yes |
| single-value | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 152 | 228 | 1.49× [1.47, 1.50] | +32.7% | [+31.9%, +33.5%] | 1.1 pts | 1.4 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 1666 | 3138 | 1.88× [1.87, 1.89] | +46.8% | [+46.5%, +47.1%] | 0.4 pts | 0.9 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 1655 | 915 | 0.56× [0.54, 0.57] | -79.9% | [-83.5%, -76.2%] | 3.4 pts | 1.9 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | baseline | 8 | 33.3 µs | 79.3 µs | 2.39× [2.36, 2.42] | +58.1% | [+57.6%, +58.7%] | 0.6 pts | 1.4 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | btree-map | 8 | 32.5 µs | 13.9 µs | 0.43× [0.42, 0.44] | -131.4% | [-136.7%, -126.2%] | 5.1 pts | 2.1 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | baseline | 8 | 373 | 386 | 1.04× [1.01, 1.07] | +3.8% | [+1.4%, +6.1%] | 2.4 pts | 1.6 | no | yes |
| single-value | dirs | 86215 | churn | ordered | btree-map | 8 | 420 | 428 | 1.02× [1.00, 1.04] | +1.9% | [+0.2%, +3.5%] | 1.8 pts | 1.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 prefix: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
