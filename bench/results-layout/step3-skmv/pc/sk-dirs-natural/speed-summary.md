| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 122 | 87.8 | 0.72× [0.72, 0.73] | -38.1% | [-39.5%, -36.6%] | 1.9 pts | 1.5 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 124 | 168 | 1.35× [1.34, 1.36] | +25.8% | [+25.1%, +26.5%] | 0.8 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | hashed | 8 | 122 | 35.6 | 0.29× [0.29, 0.30] | -240.9% | [-246.5%, -235.2%] | 5.4 pts | 2.9 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6317 | 2967 | 0.47× [0.46, 0.48] | -112.2% | [-116.3%, -108.1%] | 4.1 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6291 | 10.2 µs | 1.61× [1.59, 1.62] | +37.7% | [+37.1%, +38.2%] | 0.7 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | hashed | 8 | 6667 | 57.6 µs | 8.64× [8.56, 8.71] | +88.4% | [+88.3%, +88.5%] | 0.1 pts | 1.3 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 7128 | 3080 | 0.43× [0.42, 0.44] | -131.7% | [-138.2%, -125.3%] | 7.1 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7314 | 14.8 µs | 2.01× [1.98, 2.05] | +50.4% | [+49.5%, +51.3%] | 1.2 pts | 0.5 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | hashed | 8 | 7750 | 57.8 µs | 7.54× [7.34, 7.74] | +86.7% | [+86.4%, +87.1%] | 0.3 pts | 0.7 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 193 | 163 | 0.85× [0.83, 0.86] | -18.2% | [-20.0%, -16.4%] | 1.9 pts | 1.3 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 193 | 239 | 1.23× [1.21, 1.25] | +18.7% | [+17.6%, +19.9%] | 1.1 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | hashed | 8 | 189 | 71.7 | 0.38× [0.38, 0.39] | -162.4% | [-165.4%, -159.3%] | 3.2 pts | 1.3 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 8.99 ms | 7.72 ms | 0.86× [0.86, 0.87] | -16.0% | [-16.8%, -15.2%] | 1.2 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 8.86 ms | 11.57 ms | 1.31× [1.30, 1.32] | +23.7% | [+22.8%, +24.5%] | 1.0 pts | 1.2 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | hashed | 8 | 8.90 ms | 3.65 ms | 0.41× [0.41, 0.42] | -141.1% | [-145.2%, -137.0%] | 4.0 pts | 1.5 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 6 | 157 | 122 | 0.78× [0.77, 0.79] | -28.8% | [-30.3%, -27.3%] | 1.4 pts | 1.3 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 6 | 163 | 237 | 1.45× [1.44, 1.47] | +31.1% | [+30.4%, +31.8%] | 0.7 pts | 0.7 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | hashed | 6 | 157 | 43.1 | 0.27× [0.27, 0.28] | -265.4% | [-268.8%, -262.0%] | 3.2 pts | 0.8 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 7207 | 3719 | 0.52× [0.51, 0.52] | -93.8% | [-96.1%, -91.5%] | 2.2 pts | 1.4 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 6 | 7248 | 11.1 µs | 1.53× [1.52, 1.54] | +34.6% | [+34.2%, +34.9%] | 0.4 pts | 0.6 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | hashed | 6 | 8917 | 249.6 µs | 27.86× [27.03, 28.74] | +96.4% | [+96.3%, +96.5%] | 0.1 pts | 1.0 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 6 | 37.1 µs | 17.8 µs | 0.48× [0.48, 0.49] | -107.6% | [-110.5%, -104.7%] | 2.8 pts | 0.8 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 6 | 33.9 µs | 57.4 µs | 1.71× [1.69, 1.74] | +41.6% | [+40.8%, +42.4%] | 0.8 pts | 1.5 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | hashed | 6 | 39.8 µs | 263.3 µs | 6.63× [6.51, 6.75] | +84.9% | [+84.6%, +85.2%] | 0.3 pts | 0.3 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 6 | 263 | 226 | 0.86× [0.86, 0.87] | -15.9% | [-16.6%, -15.2%] | 0.7 pts | 0.9 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 6 | 270 | 363 | 1.33× [1.32, 1.35] | +25.0% | [+24.1%, +25.9%] | 0.9 pts | 1.2 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | hashed | 6 | 260 | 102 | 0.40× [0.39, 0.41] | -151.8% | [-156.9%, -146.7%] | 4.8 pts | 1.6 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 6 | 49.07 ms | 43.53 ms | 0.88× [0.87, 0.89] | -13.4% | [-14.9%, -11.9%] | 1.5 pts | 1.6 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 6 | 48.67 ms | 66.93 ms | 1.36× [1.34, 1.38] | +26.3% | [+25.1%, +27.5%] | 1.2 pts | 2.7 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | hashed | 6 | 48.31 ms | 20.50 ms | 0.42× [0.41, 0.43] | -138.3% | [-141.9%, -134.7%] | 3.4 pts | 0.8 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 281 | 251 | 0.90× [0.87, 0.93] | -11.2% | [-14.5%, -7.8%] | 3.4 pts | 3.0 | no | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 319 | 509 | 1.58× [1.54, 1.62] | +36.8% | [+35.2%, +38.4%] | 1.7 pts | 2.1 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | hashed | 8 | 284 | 131 | 0.47× [0.45, 0.48] | -114.3% | [-121.6%, -107.1%] | 9.8 pts | 3.2 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 11.0 µs | 7206 | 0.66× [0.65, 0.67] | -51.0% | [-53.0%, -49.1%] | 2.5 pts | 0.9 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 12.9 µs | 24.2 µs | 1.88× [1.84, 1.92] | +46.8% | [+45.6%, +48.0%] | 1.1 pts | 1.4 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 251.7 µs | 130.2 µs | 0.54× [0.50, 0.58] | -86.2% | [-99.2%, -73.3%] | 12.9 pts | 1.0 | no | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 243.2 µs | 500.1 µs | 2.15× [2.07, 2.23] | +53.4% | [+51.7%, +55.1%] | 1.8 pts | 0.6 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 523 | 474 | 0.91× [0.89, 0.92] | -10.2% | [-12.1%, -8.3%] | 2.6 pts | 0.8 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 602 | 688 | 1.16× [1.14, 1.17] | +13.4% | [+12.2%, +14.6%] | 1.6 pts | 1.2 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | hashed | 8 | 525 | 261 | 0.49× [0.47, 0.50] | -105.5% | [-112.8%, -98.2%] | 8.2 pts | 0.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 valuesFor: ordered vs hashed: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of -1.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=4096 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +4.11% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=4096 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.48% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=4096 churn: ordered vs hashed: the A/A validations found a systematic difference of +0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of -4.03% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 build: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs hashed: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of -0.65% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -3.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -3.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=86215 churn: ordered vs hashed: the A/A validations found a systematic difference of +1.96% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
