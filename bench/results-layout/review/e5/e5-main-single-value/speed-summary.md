| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 86.6 | 73.1 | 0.84× [0.83, 0.86] | -18.9% | [-20.9%, -16.9%] | 1.9 pts | 2.4 | no | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 87.5 | 108 | 1.24× [1.23, 1.26] | +19.6% | [+18.9%, +20.4%] | 0.9 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1388 | 1892 | 1.38× [1.36, 1.41] | +27.8% | [+26.6%, +28.9%] | 1.1 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1352 | 538 | 0.40× [0.39, 0.40] | -152.4% | [-155.2%, -149.7%] | 2.6 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 198 | 182 | 0.92× [0.91, 0.93] | -8.8% | [-9.9%, -7.7%] | 1.0 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 198 | 177 | 0.89× [0.88, 0.91] | -11.8% | [-13.4%, -10.1%] | 1.6 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 3.37 ms | 2.47 ms | 0.73× [0.71, 0.75] | -36.7% | [-40.3%, -33.1%] | 3.4 pts | 1.7 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 3.35 ms | 2.60 ms | 0.78× [0.77, 0.79] | -27.8% | [-29.6%, -26.1%] | 1.8 pts | 1.5 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 137 | 158 | 1.15× [1.10, 1.20] | +13.1% | [+9.4%, +16.8%] | 3.5 pts | 3.5 | no | yes |
| single-value | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 141 | 219 | 1.56× [1.54, 1.58] | +35.7% | [+35.0%, +36.5%] | 1.0 pts | 1.4 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1695 | 2763 | 1.63× [1.62, 1.64] | +38.6% | [+38.3%, +38.8%] | 0.7 pts | 1.2 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 1681 | 897 | 0.53× [0.53, 0.54] | -87.6% | [-90.2%, -85.1%] | 2.6 pts | 1.7 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 308 | 382 | 1.25× [1.24, 1.27] | +20.3% | [+19.3%, +21.3%] | 1.1 pts | 1.1 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | btree-map | 8 | 325 | 368 | 1.12× [1.11, 1.13] | +10.9% | [+10.3%, +11.5%] | 0.6 pts | 0.8 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 75.78 ms | 74.44 ms | 0.99× [0.97, 1.00] | -1.3% | [-2.6%, -0.0%] | 1.4 pts | 1.0 | yes | yes |
| single-value | dirs | 65536 | build | ordered | btree-map | 8 | 76.85 ms | 76.73 ms | 1.00× [0.99, 1.01] | +0.0% | [-1.1%, +1.1%] | 1.3 pts | 1.4 | yes | no |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 57.8 | 46.9 | 0.81× [0.80, 0.82] | -23.9% | [-25.6%, -22.3%] | 1.6 pts | 2.3 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 58.2 | 91.7 | 1.56× [1.55, 1.57] | +35.9% | [+35.3%, +36.4%] | 0.6 pts | 1.5 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 1091 | 1671 | 1.55× [1.53, 1.56] | +35.3% | [+34.7%, +36.0%] | 0.6 pts | 0.6 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1073 | 498 | 0.46× [0.45, 0.47] | -117.6% | [-121.5%, -113.7%] | 3.8 pts | 1.6 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 136 | 110 | 0.81× [0.79, 0.82] | -23.8% | [-26.0%, -21.7%] | 2.2 pts | 1.3 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 136 | 137 | 1.01× [0.99, 1.02] | +0.5% | [-0.6%, +1.6%] | 1.7 pts | 1.6 | yes | no |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.27 ms | 1.53 ms | 0.67× [0.67, 0.67] | -49.0% | [-49.8%, -48.2%] | 2.2 pts | 1.2 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.28 ms | 2.13 ms | 0.96× [0.93, 0.99] | -4.4% | [-7.8%, -1.1%] | 5.3 pts | 4.8 | no | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 85.7 | 87.8 | 1.04× [1.01, 1.07] | +3.9% | [+1.2%, +6.5%] | 2.8 pts | 2.9 | no | yes |
| single-value | street | 65536 | valuesFor | ordered | btree-map | 8 | 87.1 | 181 | 2.06× [2.03, 2.10] | +51.6% | [+50.9%, +52.3%] | 0.7 pts | 1.9 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 1329 | 2344 | 1.76× [1.75, 1.78] | +43.3% | [+42.8%, +43.8%] | 0.5 pts | 1.2 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | btree-map | 8 | 1319 | 736 | 0.56× [0.55, 0.56] | -79.5% | [-80.6%, -78.4%] | 2.0 pts | 2.0 | yes | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 201 | 240 | 1.17× [1.14, 1.20] | +14.2% | [+12.0%, +16.4%] | 3.9 pts | 2.9 | no | yes |
| single-value | street | 65536 | churn | ordered | btree-map | 8 | 203 | 254 | 1.25× [1.22, 1.27] | +19.7% | [+18.3%, +21.1%] | 1.7 pts | 2.6 | yes | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 50.12 ms | 44.72 ms | 0.88× [0.86, 0.90] | -13.6% | [-15.6%, -11.6%] | 3.2 pts | 2.4 | no | yes |
| single-value | street | 65536 | build | ordered | btree-map | 8 | 50.46 ms | 58.03 ms | 1.15× [1.14, 1.16] | +12.9% | [+12.0%, +13.7%] | 1.0 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 6 | 32.9 | 17.0 | 0.52× [0.51, 0.52] | -93.0% | [-94.9%, -91.1%] | 1.8 pts | 3.6 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 32.8 | 88.0 | 2.70× [2.67, 2.73] | +63.0% | [+62.6%, +63.3%] | 0.4 pts | 3.5 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 867 | 1000 | 1.15× [1.14, 1.16] | +13.0% | [+12.5%, +13.6%] | 0.5 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 870 | 424 | 0.49× [0.48, 0.49] | -105.4% | [-106.6%, -104.2%] | 1.2 pts | 2.0 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 6 | 76.8 | 43.9 | 0.57× [0.56, 0.58] | -74.7% | [-78.4%, -71.1%] | 3.5 pts | 3.3 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 6 | 76.6 | 121 | 1.58× [1.57, 1.60] | +36.9% | [+36.2%, +37.5%] | 0.6 pts | 1.8 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 6 | 1.13 ms | 737.9 µs | 0.66× [0.64, 0.67] | -52.5% | [-55.2%, -49.8%] | 2.6 pts | 1.1 | yes | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 6 | 1.13 ms | 1.74 ms | 1.54× [1.53, 1.56] | +35.2% | [+34.6%, +35.7%] | 0.5 pts | 1.9 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | baseline | 8 | 34.5 | 24.8 | 0.72× [0.71, 0.72] | -39.6% | [-41.1%, -38.1%] | 1.7 pts | 2.4 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | btree-map | 8 | 35.3 | 156 | 4.45× [4.37, 4.53] | +77.5% | [+77.1%, +77.9%] | 0.4 pts | 3.1 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 1687 | 1921 | 1.13× [1.13, 1.14] | +11.9% | [+11.6%, +12.2%] | 0.6 pts | 1.5 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | btree-map | 8 | 1692 | 577 | 0.34× [0.34, 0.34] | -193.4% | [-195.9%, -190.9%] | 2.4 pts | 2.0 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | baseline | 8 | 111 | 93.0 | 0.87× [0.82, 0.92] | -15.4% | [-22.2%, -8.7%] | 6.6 pts | 4.2 | no | yes |
| single-value | u64 | 65536 | churn | ordered | btree-map | 8 | 124 | 229 | 1.89× [1.82, 1.96] | +47.1% | [+45.2%, +49.0%] | 3.3 pts | 3.0 | yes | yes |
| single-value | u64 | 65536 | build | ordered | baseline | 8 | 24.09 ms | 18.42 ms | 0.75× [0.74, 0.76] | -33.0% | [-34.3%, -31.7%] | 2.9 pts | 1.7 | yes | yes |
| single-value | u64 | 65536 | build | ordered | btree-map | 8 | 24.66 ms | 50.19 ms | 2.04× [2.02, 2.05] | +50.9% | [+50.6%, +51.3%] | 0.4 pts | 0.5 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | baseline | 8 | 76.6 | 65.0 | 0.85× [0.84, 0.85] | -18.3% | [-19.4%, -17.2%] | 1.1 pts | 1.4 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | btree-map | 8 | 77.2 | 112 | 1.45× [1.44, 1.47] | +31.2% | [+30.6%, +31.7%] | 0.8 pts | 1.5 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | baseline | 8 | 1429 | 1821 | 1.29× [1.27, 1.32] | +22.7% | [+21.2%, +24.1%] | 1.5 pts | 0.8 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | btree-map | 8 | 1386 | 554 | 0.39× [0.39, 0.40] | -154.4% | [-159.6%, -149.1%] | 5.1 pts | 1.2 | yes | yes |
| single-value | url | 4096 | churn | ordered | baseline | 8 | 201 | 157 | 0.79× [0.76, 0.81] | -27.0% | [-31.1%, -23.0%] | 3.8 pts | 2.6 | no | yes |
| single-value | url | 4096 | churn | ordered | btree-map | 8 | 202 | 179 | 0.89× [0.88, 0.90] | -12.5% | [-14.0%, -10.9%] | 1.8 pts | 1.4 | yes | yes |
| single-value | url | 4096 | build | ordered | baseline | 8 | 3.12 ms | 2.10 ms | 0.67× [0.66, 0.68] | -49.6% | [-51.4%, -47.7%] | 2.1 pts | 1.3 | yes | yes |
| single-value | url | 4096 | build | ordered | btree-map | 8 | 3.13 ms | 2.44 ms | 0.78× [0.77, 0.79] | -28.1% | [-29.3%, -26.9%] | 1.2 pts | 1.4 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | baseline | 8 | 132 | 151 | 1.15× [1.09, 1.22] | +13.0% | [+8.1%, +18.0%] | 4.9 pts | 3.9 | no | yes |
| single-value | url | 65536 | valuesFor | ordered | btree-map | 8 | 136 | 220 | 1.61× [1.59, 1.64] | +38.1% | [+37.0%, +39.1%] | 1.1 pts | 1.3 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | baseline | 8 | 1875 | 2648 | 1.41× [1.40, 1.43] | +29.3% | [+28.5%, +30.0%] | 0.8 pts | 1.5 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | btree-map | 8 | 1842 | 900 | 0.49× [0.48, 0.49] | -105.2% | [-107.2%, -103.3%] | 1.9 pts | 1.0 | yes | yes |
| single-value | url | 65536 | churn | ordered | baseline | 8 | 354 | 374 | 1.05× [1.04, 1.07] | +5.1% | [+3.5%, +6.8%] | 1.7 pts | 1.6 | yes | yes |
| single-value | url | 65536 | churn | ordered | btree-map | 8 | 377 | 381 | 1.01× [1.00, 1.02] | +0.8% | [-0.4%, +2.0%] | 1.2 pts | 1.1 | yes | no |
| single-value | url | 65536 | build | ordered | baseline | 8 | 77.50 ms | 64.39 ms | 0.83× [0.81, 0.85] | -20.3% | [-22.9%, -17.6%] | 2.5 pts | 2.3 | no | yes |
| single-value | url | 65536 | build | ordered | btree-map | 8 | 78.19 ms | 75.72 ms | 0.97× [0.96, 0.98] | -3.2% | [-4.3%, -2.2%] | 1.1 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: the pooled difference of 0.02% does not clear the 0.63% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=65536 build: ordered vs btree-map: the pooled interval [-1.08%, 1.13%] includes zero
- single-value street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 churn: ordered vs btree-map: the pooled difference of 0.52% does not clear the 0.61% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=4096 churn: ordered vs btree-map: the pooled interval [-0.56%, 1.61%] includes zero
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs btree-map: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=65536 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=65536 churn: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 churn: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesBetween: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.27% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=4096 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.67% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=65536 valuesFor: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=65536 churn: ordered vs btree-map: the pooled interval [-0.39%, 2.04%] includes zero
- single-value url n=65536 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
