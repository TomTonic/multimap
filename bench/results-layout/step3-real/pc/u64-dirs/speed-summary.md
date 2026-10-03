| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | dirs | 4096 | valuesFor | ordered | btree-sets | 6 | 83.4 | 160 | 1.91× [1.86, 1.95] | +47.5% | [+46.3%, +48.8%] | 1.2 pts | 2.1 | yes | yes |
| multi | dirs | 4096 | valuesFor | ordered | hashed | 6 | 81.6 | 29.4 | 0.36× [0.36, 0.36] | -177.4% | [-179.7%, -175.2%] | 2.1 pts | 1.2 | yes | yes |
| multi | dirs | 4096 | valuesFor | ordered | map-sets | 6 | 81.8 | 71.3 | 0.87× [0.85, 0.88] | -15.4% | [-17.2%, -13.6%] | 1.7 pts | 2.0 | yes | yes |
| multi | dirs | 4096 | valuesBetween | ordered | btree-sets | 6 | 2646 | 5383 | 2.04× [2.02, 2.06] | +51.0% | [+50.5%, +51.5%] | 0.5 pts | 0.6 | yes | yes |
| multi | dirs | 4096 | valuesBetween | ordered | hashed | 6 | 2748 | 56.2 µs | 20.63× [20.42, 20.84] | +95.2% | [+95.1%, +95.2%] | 0.0 pts | 0.6 | yes | yes |
| multi | dirs | 4096 | valuesBetween | ordered | map-sets | 6 | 2716 | 56.6 µs | 20.85× [20.54, 21.18] | +95.2% | [+95.1%, +95.3%] | 0.1 pts | 1.1 | yes | yes |
| multi | dirs | 4096 | prefix | ordered | btree-sets | 6 | 2764 | 7400 | 2.68× [2.61, 2.75] | +62.7% | [+61.7%, +63.6%] | 0.9 pts | 0.9 | yes | yes |
| multi | dirs | 4096 | prefix | ordered | hashed | 6 | 2817 | 55.8 µs | 19.97× [19.56, 20.41] | +95.0% | [+94.9%, +95.1%] | 0.1 pts | 1.0 | yes | yes |
| multi | dirs | 4096 | prefix | ordered | map-sets | 6 | 2810 | 58.6 µs | 20.72× [20.35, 21.09] | +95.2% | [+95.1%, +95.3%] | 0.1 pts | 0.8 | yes | yes |
| multi | dirs | 4096 | churn | ordered | btree-sets | 6 | 150 | 226 | 1.52× [1.51, 1.54] | +34.3% | [+33.6%, +35.0%] | 0.6 pts | 0.7 | yes | yes |
| multi | dirs | 4096 | churn | ordered | hashed | 6 | 147 | 62.8 | 0.43× [0.43, 0.43] | -132.8% | [-134.5%, -131.1%] | 1.6 pts | 0.8 | yes | yes |
| multi | dirs | 4096 | churn | ordered | map-sets | 6 | 147 | 72.6 | 0.49× [0.49, 0.50] | -102.2% | [-104.1%, -100.3%] | 1.8 pts | 1.2 | yes | yes |
| multi | dirs | 4096 | build | ordered | btree-sets | 6 | 7.20 ms | 10.89 ms | 1.51× [1.50, 1.52] | +33.7% | [+33.2%, +34.2%] | 0.5 pts | 0.8 | yes | yes |
| multi | dirs | 4096 | build | ordered | hashed | 6 | 7.16 ms | 3.12 ms | 0.44× [0.43, 0.45] | -128.0% | [-131.4%, -124.5%] | 3.3 pts | 2.1 | yes | yes |
| multi | dirs | 4096 | build | ordered | map-sets | 6 | 7.18 ms | 3.83 ms | 0.54× [0.52, 0.56] | -85.4% | [-90.7%, -80.2%] | 5.0 pts | 1.3 | yes | yes |
| multi | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 115 | 219 | 1.91× [1.89, 1.93] | +47.8% | [+47.2%, +48.3%] | 0.6 pts | 1.3 | yes | yes |
| multi | dirs | 16384 | valuesFor | ordered | hashed | 8 | 113 | 36.0 | 0.32× [0.31, 0.33] | -212.4% | [-218.7%, -206.0%] | 6.0 pts | 2.3 | yes | yes |
| multi | dirs | 16384 | valuesFor | ordered | map-sets | 8 | 112 | 83.8 | 0.75× [0.75, 0.76] | -32.9% | [-33.8%, -31.9%] | 1.0 pts | 1.3 | yes | yes |
| multi | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3421 | 6127 | 1.79× [1.77, 1.81] | +44.2% | [+43.6%, +44.8%] | 0.6 pts | 1.5 | yes | yes |
| multi | dirs | 16384 | valuesBetween | ordered | hashed | 8 | 3695 | 243.8 µs | 64.79× [62.15, 67.67] | +98.5% | [+98.4%, +98.5%] | 0.1 pts | 1.2 | yes | yes |
| multi | dirs | 16384 | valuesBetween | ordered | map-sets | 8 | 3922 | 233.2 µs | 60.46× [58.37, 62.72] | +98.3% | [+98.3%, +98.4%] | 0.1 pts | 0.8 | yes | yes |
| multi | dirs | 16384 | prefix | ordered | btree-sets | 8 | 15.6 µs | 34.8 µs | 2.16× [2.14, 2.18] | +53.7% | [+53.3%, +54.1%] | 1.0 pts | 1.1 | yes | yes |
| multi | dirs | 16384 | prefix | ordered | hashed | 8 | 15.6 µs | 255.0 µs | 16.14× [15.43, 16.92] | +93.8% | [+93.5%, +94.1%] | 0.3 pts | 0.7 | yes | yes |
| multi | dirs | 16384 | prefix | ordered | map-sets | 8 | 16.4 µs | 263.5 µs | 15.99× [15.18, 16.90] | +93.7% | [+93.4%, +94.1%] | 0.3 pts | 0.8 | yes | yes |
| multi | dirs | 16384 | churn | ordered | btree-sets | 8 | 206 | 333 | 1.62× [1.61, 1.63] | +38.2% | [+37.7%, +38.6%] | 0.7 pts | 0.9 | yes | yes |
| multi | dirs | 16384 | churn | ordered | hashed | 8 | 194 | 80.9 | 0.42× [0.41, 0.43] | -137.5% | [-144.5%, -130.5%] | 7.0 pts | 2.1 | yes | yes |
| multi | dirs | 16384 | churn | ordered | map-sets | 8 | 196 | 99.2 | 0.51× [0.49, 0.54] | -94.3% | [-103.2%, -85.3%] | 9.3 pts | 3.1 | yes | yes |
| multi | dirs | 16384 | build | ordered | btree-sets | 8 | 40.73 ms | 60.58 ms | 1.49× [1.48, 1.51] | +33.1% | [+32.5%, +33.6%] | 0.5 pts | 0.9 | yes | yes |
| multi | dirs | 16384 | build | ordered | hashed | 8 | 39.51 ms | 16.63 ms | 0.42× [0.42, 0.42] | -137.9% | [-140.1%, -135.8%] | 2.7 pts | 0.8 | yes | yes |
| multi | dirs | 16384 | build | ordered | map-sets | 8 | 40.33 ms | 21.89 ms | 0.55× [0.53, 0.57] | -83.2% | [-89.6%, -76.8%] | 6.0 pts | 1.9 | yes | yes |
| multi | dirs | 86215 | valuesFor | ordered | btree-sets | 12 | 216 | 414 | 1.92× [1.89, 1.95] | +47.9% | [+47.1%, +48.7%] | 1.0 pts | 1.2 | yes | yes |
| multi | dirs | 86215 | valuesFor | ordered | hashed | 12 | 177 | 73.5 | 0.42× [0.40, 0.44] | -139.3% | [-152.2%, -126.4%] | 12.6 pts | 4.0 | yes | yes |
| multi | dirs | 86215 | valuesFor | ordered | map-sets | 12 | 191 | 199 | 1.04× [1.02, 1.05] | +3.5% | [+2.2%, +4.7%] | 2.2 pts | 2.1 | yes | yes |
| multi | dirs | 86215 | valuesBetween | ordered | btree-sets | 12 | 5909 | 12.5 µs | 2.13× [2.10, 2.16] | +53.1% | [+52.4%, +53.8%] | 1.4 pts | 1.6 | yes | yes |
| multi | dirs | 86215 | prefix | ordered | btree-sets | 12 | 121.4 µs | 242.9 µs | 2.10× [2.06, 2.15] | +52.5% | [+51.6%, +53.4%] | 1.2 pts | 0.4 | yes | yes |
| multi | dirs | 86215 | churn | ordered | btree-sets | 12 | 478 | 633 | 1.33× [1.31, 1.35] | +25.0% | [+23.9%, +26.2%] | 1.5 pts | 2.6 | yes | yes |
| multi | dirs | 86215 | churn | ordered | hashed | 12 | 365 | 192 | 0.52× [0.51, 0.54] | -90.6% | [-95.6%, -85.6%] | 7.2 pts | 1.1 | yes | yes |
| multi | dirs | 86215 | churn | ordered | map-sets | 12 | 409 | 267 | 0.65× [0.64, 0.65] | -54.4% | [-55.2%, -53.6%] | 1.4 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi dirs n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.76% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=4096 build: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=4096 build: ordered vs map-sets: the A/A validations found a systematic difference of -0.80% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 valuesFor: ordered vs hashed: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 valuesBetween: ordered vs map-sets: the A/A validations found a systematic difference of +2.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +10.77% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +3.04% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +2.97% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=16384 churn: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=16384 churn: ordered vs map-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=86215 valuesFor: ordered vs hashed: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=86215 valuesFor: ordered vs map-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -2.57% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi dirs n=86215 churn: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
