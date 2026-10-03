| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 129 | 208 | 1.61× [1.60, 1.63] | +38.0% | [+37.4%, +38.7%] | 0.8 pts | 1.0 | yes | yes |
| multi | dirs | 4096 | valuesFor | ordered | hashed | 8 | 126 | 41.4 | 0.33× [0.33, 0.34] | -201.2% | [-206.8%, -195.5%] | 6.7 pts | 4.7 | yes | yes |
| multi | dirs | 4096 | valuesFor | ordered | map-sets | 8 | 126 | 85.5 | 0.68× [0.68, 0.68] | -47.1% | [-47.7%, -46.5%] | 0.7 pts | 1.2 | yes | yes |
| multi | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 3918 | 6319 | 1.62× [1.61, 1.62] | +38.1% | [+37.7%, +38.4%] | 0.4 pts | 1.0 | yes | yes |
| multi | dirs | 4096 | valuesBetween | ordered | hashed | 8 | 3958 | 78.9 µs | 19.99× [19.91, 20.07] | +95.0% | [+95.0%, +95.0%] | 0.0 pts | 0.7 | yes | yes |
| multi | dirs | 4096 | valuesBetween | ordered | map-sets | 8 | 3931 | 77.1 µs | 19.61× [19.54, 19.68] | +94.9% | [+94.9%, +94.9%] | 0.0 pts | 0.9 | yes | yes |
| multi | dirs | 4096 | prefix | ordered | btree-sets | 8 | 4443 | 8367 | 1.89× [1.84, 1.94] | +47.1% | [+45.7%, +48.6%] | 1.7 pts | 0.9 | yes | yes |
| multi | dirs | 4096 | prefix | ordered | hashed | 8 | 4555 | 77.9 µs | 17.07× [16.78, 17.36] | +94.1% | [+94.0%, +94.2%] | 0.1 pts | 0.8 | yes | yes |
| multi | dirs | 4096 | prefix | ordered | map-sets | 8 | 4586 | 78.7 µs | 17.18× [16.93, 17.43] | +94.2% | [+94.1%, +94.3%] | 0.1 pts | 0.7 | yes | yes |
| multi | dirs | 4096 | churn | ordered | btree-sets | 8 | 196 | 275 | 1.40× [1.39, 1.42] | +28.7% | [+28.0%, +29.4%] | 0.9 pts | 0.9 | yes | yes |
| multi | dirs | 4096 | churn | ordered | hashed | 8 | 191 | 69.8 | 0.36× [0.36, 0.37] | -174.2% | [-175.6%, -172.8%] | 1.7 pts | 1.2 | yes | yes |
| multi | dirs | 4096 | churn | ordered | map-sets | 8 | 191 | 73.1 | 0.39× [0.38, 0.39] | -159.7% | [-161.2%, -158.2%] | 1.8 pts | 1.4 | yes | yes |
| multi | dirs | 4096 | build | ordered | btree-sets | 8 | 9.14 ms | 13.10 ms | 1.44× [1.43, 1.44] | +30.4% | [+30.1%, +30.7%] | 0.4 pts | 2.2 | yes | yes |
| multi | dirs | 4096 | build | ordered | hashed | 8 | 9.17 ms | 3.69 ms | 0.40× [0.40, 0.40] | -148.5% | [-149.5%, -147.5%] | 1.2 pts | 1.5 | yes | yes |
| multi | dirs | 4096 | build | ordered | map-sets | 8 | 9.14 ms | 4.16 ms | 0.45× [0.45, 0.46] | -120.3% | [-121.4%, -119.2%] | 1.3 pts | 1.2 | yes | yes |
| multi | dirs | 16384 | valuesFor | ordered | btree-sets | 20 | 168 | 272 | 1.63× [1.60, 1.65] | +38.5% | [+37.6%, +39.5%] | 1.3 pts | 1.5 | yes | yes |
| multi | dirs | 16384 | valuesFor | ordered | hashed | 20 | 163 | 46.4 | 0.29× [0.28, 0.30] | -244.4% | [-254.0%, -234.8%] | 12.1 pts | 2.1 | yes | yes |
| multi | dirs | 16384 | valuesFor | ordered | map-sets | 20 | 166 | 102 | 0.63× [0.62, 0.64] | -59.3% | [-62.2%, -56.4%] | 4.5 pts | 1.7 | yes | yes |
| multi | dirs | 16384 | valuesBetween | ordered | btree-sets | 20 | 4490 | 6995 | 1.58× [1.56, 1.61] | +36.8% | [+35.9%, +37.8%] | 1.5 pts | 2.9 | yes | yes |
| multi | dirs | 16384 | valuesBetween | ordered | hashed | 20 | 4450 | 306.9 µs | 68.32× [66.98, 69.71] | +98.5% | [+98.5%, +98.6%] | 0.0 pts | 4.4 | yes | yes |
| multi | dirs | 16384 | valuesBetween | ordered | map-sets | 20 | 4562 | 289.1 µs | 61.44× [58.94, 64.16] | +98.4% | [+98.3%, +98.4%] | 0.1 pts | 7.9 | yes | yes |
| multi | dirs | 16384 | prefix | ordered | btree-sets | 20 | 20.7 µs | 33.0 µs | 1.63× [1.59, 1.67] | +38.6% | [+37.1%, +40.1%] | 1.9 pts | 6.5 | yes | yes |
| multi | dirs | 16384 | prefix | ordered | hashed | 20 | 21.9 µs | 316.5 µs | 14.34× [14.12, 14.57] | +93.0% | [+92.9%, +93.1%] | 0.1 pts | 0.5 | yes | yes |
| multi | dirs | 16384 | prefix | ordered | map-sets | 20 | 22.1 µs | 320.7 µs | 14.44× [14.19, 14.69] | +93.1% | [+93.0%, +93.2%] | 0.2 pts | 0.4 | yes | yes |
| multi | dirs | 16384 | churn | ordered | btree-sets | 20 | 277 | 378 | 1.37× [1.36, 1.38] | +27.0% | [+26.2%, +27.8%] | 1.1 pts | 1.1 | yes | yes |
| multi | dirs | 16384 | churn | ordered | hashed | 20 | 248 | 88.7 | 0.37× [0.35, 0.38] | -173.2% | [-184.0%, -162.5%] | 16.8 pts | 5.9 | yes | yes |
| multi | dirs | 16384 | churn | ordered | map-sets | 20 | 258 | 110 | 0.44× [0.42, 0.46] | -128.7% | [-139.9%, -117.5%] | 19.1 pts | 8.3 | yes | yes |
| multi | dirs | 16384 | build | ordered | btree-sets | 20 | 47.44 ms | 66.83 ms | 1.42× [1.40, 1.44] | +29.7% | [+28.7%, +30.6%] | 1.4 pts | 12.9 | yes | yes |
| multi | dirs | 16384 | build | ordered | hashed | 20 | 47.43 ms | 17.07 ms | 0.36× [0.36, 0.37] | -177.1% | [-181.1%, -173.1%] | 5.0 pts | 7.9 | yes | yes |
| multi | dirs | 16384 | build | ordered | map-sets | 20 | 47.44 ms | 20.27 ms | 0.43× [0.42, 0.44] | -130.8% | [-136.3%, -125.2%] | 7.8 pts | 12.1 | yes | yes |
| multi | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 344 | 573 | 1.70× [1.65, 1.74] | +41.1% | [+39.5%, +42.7%] | 1.9 pts | 0.8 | yes | yes |
| multi | dirs | 86215 | valuesFor | ordered | hashed | 8 | 281 | 100 | 0.36× [0.35, 0.37] | -176.4% | [-183.9%, -168.9%] | 8.9 pts | 1.1 | yes | yes |
| multi | dirs | 86215 | valuesFor | ordered | map-sets | 8 | 325 | 218 | 0.69× [0.67, 0.70] | -45.8% | [-49.4%, -42.2%] | 4.3 pts | 1.1 | yes | yes |
| multi | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 6633 | 12.1 µs | 1.82× [1.80, 1.83] | +45.0% | [+44.6%, +45.5%] | 0.5 pts | 0.9 | yes | yes |
| multi | dirs | 86215 | prefix | ordered | btree-sets | 8 | 139.7 µs | 310.2 µs | 2.17× [2.13, 2.22] | +53.9% | [+53.0%, +54.9%] | 1.1 pts | 1.2 | yes | yes |
| multi | dirs | 86215 | churn | ordered | btree-sets | 8 | 628 | 779 | 1.23× [1.21, 1.25] | +18.8% | [+17.6%, +20.1%] | 1.5 pts | 0.7 | yes | yes |
| multi | dirs | 86215 | churn | ordered | hashed | 8 | 519 | 196 | 0.36× [0.35, 0.37] | -180.2% | [-189.3%, -171.1%] | 10.9 pts | 0.4 | yes | yes |
| multi | dirs | 86215 | churn | ordered | map-sets | 8 | 578 | 269 | 0.45× [0.43, 0.46] | -124.3% | [-131.1%, -117.4%] | 8.2 pts | 0.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi dirs n=4096 valuesFor: ordered vs hashed: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=4096 churn: ordered vs hashed: the A/A validations found a systematic difference of -0.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=4096 build: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 valuesFor: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 valuesFor: ordered vs map-sets: the A/A validations found a systematic difference of +0.71% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 valuesBetween: ordered vs hashed: the A/A validations found a systematic difference of +0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 valuesBetween: ordered vs hashed: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 valuesBetween: ordered vs map-sets: the processes scatter 7.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -4.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 prefix: ordered vs btree-sets: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of -3.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of -3.57% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 churn: ordered vs hashed: the processes scatter 5.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 churn: ordered vs map-sets: the A/A validations found a systematic difference of +0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 churn: ordered vs map-sets: the processes scatter 8.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 build: ordered vs btree-sets: the processes scatter 12.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 build: ordered vs hashed: the processes scatter 7.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 build: ordered vs map-sets: the processes scatter 12.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=86215 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +3.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=86215 valuesFor: ordered vs hashed: the A/A validations found a systematic difference of +2.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=86215 valuesFor: ordered vs map-sets: the A/A validations found a systematic difference of +3.03% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=86215 churn: ordered vs hashed: the A/A validations found a systematic difference of +5.82% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=86215 churn: ordered vs map-sets: the A/A validations found a systematic difference of +3.77% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
