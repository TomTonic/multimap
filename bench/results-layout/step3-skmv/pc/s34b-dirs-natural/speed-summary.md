| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 124 | 88.6 | 0.71× [0.70, 0.72] | -40.1% | [-41.9%, -38.3%] | 1.7 pts | 1.4 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 127 | 168 | 1.33× [1.27, 1.39] | +24.6% | [+21.3%, +27.9%] | 3.1 pts | 3.1 | no | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | hashed | 8 | 124 | 34.5 | 0.28× [0.28, 0.28] | -259.6% | [-262.8%, -256.4%] | 3.2 pts | 1.7 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6383 | 2999 | 0.47× [0.46, 0.47] | -113.5% | [-116.0%, -111.0%] | 2.5 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6393 | 10.3 µs | 1.62× [1.60, 1.63] | +38.1% | [+37.7%, +38.5%] | 0.4 pts | 0.7 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | hashed | 8 | 6683 | 57.3 µs | 8.74× [8.42, 9.09] | +88.6% | [+88.1%, +89.0%] | 0.4 pts | 4.0 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 7092 | 3133 | 0.44× [0.43, 0.45] | -128.1% | [-135.3%, -120.9%] | 8.2 pts | 1.2 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7306 | 14.6 µs | 1.98× [1.95, 2.02] | +49.6% | [+48.6%, +50.6%] | 1.7 pts | 0.6 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | hashed | 8 | 7684 | 57.4 µs | 7.57× [7.36, 7.79] | +86.8% | [+86.4%, +87.2%] | 0.4 pts | 1.2 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 192 | 162 | 0.83× [0.82, 0.84] | -20.8% | [-22.3%, -19.3%] | 1.5 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 195 | 239 | 1.24× [1.22, 1.26] | +19.3% | [+17.9%, +20.6%] | 1.4 pts | 1.2 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | hashed | 8 | 188 | 71.6 | 0.38× [0.38, 0.38] | -164.1% | [-165.6%, -162.6%] | 2.1 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 9.13 ms | 7.70 ms | 0.85× [0.84, 0.86] | -18.3% | [-19.7%, -16.9%] | 1.4 pts | 1.2 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 9.08 ms | 11.81 ms | 1.31× [1.30, 1.33] | +23.7% | [+22.9%, +24.6%] | 0.8 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | hashed | 8 | 9.05 ms | 3.68 ms | 0.41× [0.40, 0.42] | -143.8% | [-148.1%, -139.6%] | 4.0 pts | 1.5 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 161 | 123 | 0.77× [0.76, 0.78] | -30.3% | [-32.0%, -28.5%] | 2.3 pts | 1.7 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 166 | 237 | 1.41× [1.39, 1.43] | +29.1% | [+27.9%, +30.2%] | 1.7 pts | 1.8 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | hashed | 8 | 160 | 43.8 | 0.28× [0.26, 0.30] | -258.0% | [-284.0%, -232.0%] | 24.8 pts | 3.2 | no | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 7306 | 3771 | 0.52× [0.51, 0.53] | -91.4% | [-95.1%, -87.6%] | 3.9 pts | 2.0 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 7390 | 11.5 µs | 1.56× [1.52, 1.61] | +36.0% | [+34.1%, +37.9%] | 1.8 pts | 2.2 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | hashed | 8 | 9153 | 248.9 µs | 26.58× [24.47, 29.09] | +96.2% | [+95.9%, +96.6%] | 0.3 pts | 3.1 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 8 | 37.4 µs | 17.2 µs | 0.48× [0.47, 0.48] | -110.5% | [-114.7%, -106.2%] | 4.1 pts | 1.3 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 8 | 33.9 µs | 58.0 µs | 1.72× [1.70, 1.73] | +41.8% | [+41.2%, +42.3%] | 1.0 pts | 1.6 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | hashed | 8 | 39.8 µs | 262.8 µs | 6.56× [6.36, 6.78] | +84.8% | [+84.3%, +85.2%] | 0.6 pts | 0.8 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 8 | 273 | 227 | 0.85× [0.84, 0.86] | -17.8% | [-19.5%, -16.2%] | 1.6 pts | 1.9 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 8 | 276 | 372 | 1.33× [1.31, 1.34] | +24.6% | [+23.8%, +25.3%] | 0.9 pts | 1.3 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | hashed | 8 | 264 | 105 | 0.41× [0.39, 0.42] | -145.9% | [-154.2%, -137.7%] | 10.4 pts | 3.2 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 8 | 49.43 ms | 43.46 ms | 0.87× [0.86, 0.88] | -14.5% | [-15.9%, -13.2%] | 1.4 pts | 1.4 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 8 | 49.01 ms | 67.17 ms | 1.38× [1.37, 1.40] | +27.6% | [+26.8%, +28.4%] | 0.9 pts | 1.1 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | hashed | 8 | 49.30 ms | 21.12 ms | 0.43× [0.41, 0.46] | -131.1% | [-143.1%, -119.1%] | 11.4 pts | 3.2 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 281 | 256 | 0.91× [0.89, 0.92] | -10.4% | [-12.5%, -8.4%] | 2.2 pts | 1.8 | no | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 313 | 507 | 1.56× [1.51, 1.61] | +35.9% | [+33.9%, +37.8%] | 2.3 pts | 3.1 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | hashed | 8 | 289 | 136 | 0.47× [0.46, 0.49] | -111.4% | [-117.9%, -104.9%] | 7.3 pts | 2.2 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 11.3 µs | 7670 | 0.67× [0.65, 0.69] | -49.6% | [-53.9%, -45.3%] | 4.2 pts | 1.8 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 12.8 µs | 24.5 µs | 1.90× [1.84, 1.95] | +47.2% | [+45.8%, +48.7%] | 1.4 pts | 2.3 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 254.7 µs | 133.9 µs | 0.58× [0.54, 0.61] | -73.8% | [-84.8%, -62.7%] | 11.0 pts | 0.9 | no | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 253.6 µs | 534.3 µs | 2.13× [1.99, 2.29] | +53.1% | [+49.9%, +56.3%] | 3.7 pts | 1.2 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 530 | 472 | 0.88× [0.86, 0.91] | -13.0% | [-16.3%, -9.8%] | 3.1 pts | 0.8 | no | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 617 | 710 | 1.15× [1.14, 1.16] | +12.9% | [+12.2%, +13.6%] | 1.0 pts | 1.2 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | hashed | 8 | 531 | 260 | 0.48× [0.46, 0.50] | -109.3% | [-117.2%, -101.3%] | 8.9 pts | 1.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 valuesFor: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 valuesBetween: ordered vs hashed: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesFor: ordered vs hashed: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs hashed: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of -3.48% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 churn: ordered vs hashed: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 build: ordered vs hashed: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs hashed: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -3.70% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -2.08% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
