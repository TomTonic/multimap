| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 6 | 91.1 | 79.1 | 0.87× [0.86, 0.88] | -15.4% | [-16.7%, -14.2%] | 1.2 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 6 | 92.2 | 109 | 1.18× [1.16, 1.20] | +15.5% | [+14.1%, +16.9%] | 1.3 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 1420 | 2329 | 1.62× [1.60, 1.65] | +38.4% | [+37.3%, +39.4%] | 1.0 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 6 | 1392 | 538 | 0.38× [0.38, 0.39] | -161.3% | [-164.9%, -157.8%] | 3.4 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 6 | 1641 | 2669 | 1.61× [1.58, 1.63] | +37.8% | [+36.9%, +38.7%] | 0.8 pts | 0.7 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | btree-map | 6 | 1639 | 642 | 0.39× [0.39, 0.39] | -157.7% | [-158.6%, -156.7%] | 0.9 pts | 0.4 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 6 | 264 | 196 | 0.75× [0.73, 0.77] | -33.4% | [-36.2%, -30.6%] | 2.6 pts | 2.2 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 6 | 263 | 179 | 0.68× [0.67, 0.68] | -47.8% | [-48.7%, -46.8%] | 0.9 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 6 | 3.86 ms | 2.89 ms | 0.75× [0.74, 0.77] | -32.6% | [-34.6%, -30.6%] | 1.9 pts | 2.2 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 6 | 3.91 ms | 2.59 ms | 0.67× [0.66, 0.69] | -48.6% | [-51.8%, -45.3%] | 3.1 pts | 2.2 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 6 | 112 | 107 | 0.96× [0.94, 0.97] | -4.7% | [-6.1%, -3.2%] | 1.4 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 6 | 113 | 156 | 1.39× [1.35, 1.43] | +27.9% | [+26.0%, +29.8%] | 1.8 pts | 2.1 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 1561 | 2847 | 1.81× [1.80, 1.83] | +44.9% | [+44.3%, +45.4%] | 0.5 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 6 | 1560 | 707 | 0.45× [0.44, 0.46] | -122.2% | [-126.6%, -117.8%] | 4.2 pts | 2.3 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 6 | 5794 | 11.7 µs | 1.99× [1.94, 2.05] | +49.7% | [+48.4%, +51.1%] | 1.3 pts | 0.5 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | btree-map | 6 | 5702 | 2340 | 0.41× [0.40, 0.42] | -144.5% | [-149.9%, -139.0%] | 5.2 pts | 0.7 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 6 | 292 | 249 | 0.85× [0.85, 0.86] | -17.1% | [-18.0%, -16.1%] | 0.9 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 6 | 291 | 241 | 0.83× [0.82, 0.84] | -20.7% | [-21.9%, -19.5%] | 1.2 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 6 | 16.92 ms | 13.52 ms | 0.80× [0.79, 0.80] | -25.5% | [-26.5%, -24.5%] | 1.0 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 6 | 16.99 ms | 13.15 ms | 0.77× [0.76, 0.78] | -30.0% | [-31.3%, -28.7%] | 1.2 pts | 1.1 | yes | yes |
| single-value | dirs | 86215 | valuesFor | ordered | baseline | 8 | 148 | 153 | 1.03× [1.02, 1.05] | +3.3% | [+1.9%, +4.6%] | 1.3 pts | 1.5 | yes | yes |
| single-value | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 152 | 229 | 1.50× [1.47, 1.54] | +33.4% | [+31.8%, +35.0%] | 1.5 pts | 1.7 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 1726 | 3164 | 1.83× [1.81, 1.85] | +45.4% | [+44.8%, +45.9%] | 0.5 pts | 1.6 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 1721 | 925 | 0.54× [0.53, 0.54] | -85.9% | [-87.9%, -83.8%] | 2.0 pts | 1.2 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | baseline | 8 | 34.6 µs | 79.4 µs | 2.28× [2.24, 2.33] | +56.1% | [+55.3%, +57.0%] | 0.8 pts | 2.1 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | btree-map | 8 | 34.0 µs | 14.4 µs | 0.42× [0.41, 0.42] | -140.7% | [-144.2%, -137.2%] | 3.9 pts | 1.7 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | baseline | 8 | 395 | 402 | 1.02× [0.99, 1.04] | +1.5% | [-1.2%, +4.2%] | 2.6 pts | 2.3 | no | no |
| single-value | dirs | 86215 | churn | ordered | btree-map | 8 | 430 | 426 | 1.00× [0.98, 1.01] | -0.3% | [-1.5%, +1.0%] | 1.2 pts | 1.4 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +1.05% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 build: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +2.67% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=86215 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 churn: ordered vs baseline: the pooled interval [-1.23%, 4.22%] includes zero
- single-value dirs n=86215 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 churn: ordered vs btree-map: the pooled difference of -0.28% does not clear the 0.48% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=86215 churn: ordered vs btree-map: the pooled interval [-1.54%, 0.99%] includes zero
