| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 130 | 211 | 1.60× [1.58, 1.63] | +37.6% | [+36.7%, +38.5%] | 1.1 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | hashed | 8 | 126 | 41.5 | 0.33× [0.33, 0.33] | -202.0% | [-204.6%, -199.5%] | 3.1 pts | 1.9 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | map-sets | 8 | 127 | 86.2 | 0.68× [0.68, 0.68] | -47.4% | [-47.7%, -47.2%] | 0.3 pts | 0.4 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 3901 | 6352 | 1.63× [1.62, 1.64] | +38.8% | [+38.4%, +39.1%] | 0.4 pts | 1.1 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | hashed | 8 | 3905 | 79.3 µs | 20.37× [20.28, 20.45] | +95.1% | [+95.1%, +95.1%] | 0.0 pts | 0.7 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | map-sets | 8 | 3949 | 77.4 µs | 19.70× [19.59, 19.81] | +94.9% | [+94.9%, +95.0%] | 0.0 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 4471 | 8672 | 1.89× [1.86, 1.92] | +47.1% | [+46.3%, +47.9%] | 1.0 pts | 0.4 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | hashed | 8 | 4490 | 78.6 µs | 17.48× [17.21, 17.75] | +94.3% | [+94.2%, +94.4%] | 0.1 pts | 0.7 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | map-sets | 8 | 4526 | 78.7 µs | 17.29× [17.07, 17.52] | +94.2% | [+94.1%, +94.3%] | 0.1 pts | 0.5 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 214 | 284 | 1.33× [1.32, 1.34] | +25.0% | [+24.5%, +25.6%] | 0.7 pts | 0.7 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | hashed | 8 | 207 | 78.8 | 0.38× [0.38, 0.38] | -162.5% | [-164.4%, -160.5%] | 2.3 pts | 1.7 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | map-sets | 8 | 209 | 85.0 | 0.41× [0.40, 0.41] | -145.1% | [-147.2%, -142.9%] | 2.5 pts | 1.9 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | btree-sets | 8 | 10.26 ms | 14.13 ms | 1.37× [1.37, 1.38] | +27.2% | [+26.8%, +27.5%] | 0.4 pts | 1.9 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | hashed | 8 | 10.31 ms | 4.40 ms | 0.43× [0.43, 0.43] | -132.6% | [-134.0%, -131.2%] | 1.6 pts | 1.2 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | map-sets | 8 | 10.29 ms | 5.09 ms | 0.49× [0.49, 0.50] | -102.2% | [-105.1%, -99.3%] | 3.5 pts | 2.3 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | btree-sets | 18 | 177 | 289 | 1.66× [1.64, 1.68] | +39.8% | [+39.1%, +40.5%] | 1.4 pts | 0.8 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | hashed | 18 | 167 | 50.6 | 0.31× [0.30, 0.32] | -221.4% | [-229.0%, -213.7%] | 12.1 pts | 1.2 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | map-sets | 18 | 175 | 118 | 0.70× [0.68, 0.72] | -43.4% | [-47.1%, -39.7%] | 5.3 pts | 1.3 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 18 | 4623 | 7508 | 1.65× [1.63, 1.67] | +39.4% | [+38.7%, +40.0%] | 2.1 pts | 1.9 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | hashed | 18 | 4474 | 309.2 µs | 67.03× [63.51, 70.96] | +98.5% | [+98.4%, +98.6%] | 0.1 pts | 6.9 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | map-sets | 18 | 4835 | 292.6 µs | 59.30× [57.16, 61.61] | +98.3% | [+98.3%, +98.4%] | 0.1 pts | 3.6 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | btree-sets | 18 | 20.8 µs | 33.8 µs | 1.67× [1.65, 1.69] | +40.3% | [+39.6%, +41.0%] | 1.3 pts | 2.5 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | hashed | 18 | 21.6 µs | 318.2 µs | 14.54× [14.22, 14.86] | +93.1% | [+93.0%, +93.3%] | 0.2 pts | 0.5 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | map-sets | 18 | 22.1 µs | 323.8 µs | 14.61× [14.39, 14.83] | +93.2% | [+93.1%, +93.3%] | 0.2 pts | 0.4 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | btree-sets | 18 | 316 | 410 | 1.29× [1.28, 1.30] | +22.5% | [+21.9%, +23.2%] | 1.2 pts | 1.2 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | hashed | 18 | 284 | 115 | 0.42× [0.40, 0.43] | -140.7% | [-151.3%, -130.1%] | 15.9 pts | 5.7 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | map-sets | 18 | 301 | 144 | 0.49× [0.47, 0.51] | -103.9% | [-113.2%, -94.7%] | 13.4 pts | 9.1 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | btree-sets | 18 | 51.52 ms | 73.10 ms | 1.41× [1.41, 1.42] | +29.3% | [+29.0%, +29.6%] | 0.4 pts | 2.7 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | hashed | 18 | 51.44 ms | 21.62 ms | 0.42× [0.42, 0.43] | -136.7% | [-140.2%, -133.2%] | 4.9 pts | 7.9 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | map-sets | 18 | 51.48 ms | 27.09 ms | 0.53× [0.52, 0.54] | -88.7% | [-91.3%, -86.0%] | 3.6 pts | 7.5 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | btree-sets | 20 | 366 | 583 | 1.62× [1.60, 1.64] | +38.4% | [+37.7%, +39.1%] | 2.0 pts | 0.9 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | hashed | 20 | 305 | 106 | 0.36× [0.35, 0.36] | -181.1% | [-184.1%, -178.2%] | 6.2 pts | 0.8 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | map-sets | 20 | 340 | 229 | 0.69× [0.68, 0.70] | -45.3% | [-47.0%, -43.6%] | 2.9 pts | 0.8 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 20 | 7099 | 13.1 µs | 1.85× [1.84, 1.87] | +46.0% | [+45.6%, +46.4%] | 0.7 pts | 0.9 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | btree-sets | 20 | 150.7 µs | 313.8 µs | 2.10× [2.09, 2.12] | +52.4% | [+52.1%, +52.8%] | 1.0 pts | 1.1 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | btree-sets | 20 | 686 | 805 | 1.18× [1.16, 1.19] | +15.0% | [+14.1%, +15.9%] | 1.8 pts | 0.8 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | hashed | 20 | 592 | 218 | 0.37× [0.35, 0.39] | -170.2% | [-186.5%, -154.0%] | 25.3 pts | 0.8 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | map-sets | 20 | 615 | 292 | 0.46× [0.45, 0.47] | -116.8% | [-120.8%, -112.8%] | 8.7 pts | 0.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str dirs n=4096 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.76% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=4096 valuesBetween: ordered vs map-sets: the A/A validations found a systematic difference of +0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=4096 build: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesFor: ordered vs hashed: the A/A validations found a systematic difference of -1.89% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesFor: ordered vs map-sets: the A/A validations found a systematic difference of +1.01% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesBetween: ordered vs hashed: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesBetween: ordered vs hashed: the processes scatter 6.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesBetween: ordered vs map-sets: the A/A validations found a systematic difference of +1.81% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesBetween: ordered vs map-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 prefix: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of -3.05% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of -3.28% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 churn: ordered vs hashed: the processes scatter 5.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 churn: ordered vs map-sets: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 churn: ordered vs map-sets: the processes scatter 9.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs hashed: the processes scatter 7.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs map-sets: the processes scatter 7.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +2.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 valuesFor: ordered vs hashed: the A/A validations found a systematic difference of +1.71% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 valuesFor: ordered vs map-sets: the A/A validations found a systematic difference of +1.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +0.59% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -1.03% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 churn: ordered vs hashed: the A/A validations found a systematic difference of +6.97% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 churn: ordered vs map-sets: the A/A validations found a systematic difference of +4.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
