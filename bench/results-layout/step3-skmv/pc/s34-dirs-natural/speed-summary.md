| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 122 | 88.7 | 0.73× [0.72, 0.74] | -37.4% | [-39.7%, -35.2%] | 2.2 pts | 1.9 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 124 | 168 | 1.35× [1.34, 1.37] | +26.2% | [+25.2%, +27.2%] | 1.1 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | hashed | 8 | 122 | 34.7 | 0.28× [0.28, 0.29] | -251.3% | [-256.3%, -246.3%] | 4.9 pts | 2.1 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6354 | 3025 | 0.47× [0.47, 0.48] | -110.6% | [-114.4%, -106.8%] | 3.6 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6410 | 10.3 µs | 1.60× [1.59, 1.61] | +37.6% | [+37.2%, +38.1%] | 0.5 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | hashed | 8 | 6681 | 57.2 µs | 8.58× [8.46, 8.70] | +88.3% | [+88.2%, +88.5%] | 0.2 pts | 1.4 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 7231 | 3107 | 0.44× [0.43, 0.44] | -129.2% | [-131.7%, -126.6%] | 3.9 pts | 0.6 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7180 | 14.6 µs | 1.98× [1.94, 2.02] | +49.5% | [+48.4%, +50.6%] | 1.2 pts | 0.5 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | hashed | 8 | 7709 | 57.3 µs | 7.48× [7.29, 7.68] | +86.6% | [+86.3%, +87.0%] | 0.4 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 188 | 162 | 0.86× [0.85, 0.87] | -16.4% | [-18.2%, -14.7%] | 1.8 pts | 1.3 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 190 | 239 | 1.25× [1.24, 1.27] | +20.3% | [+19.1%, +21.4%] | 1.1 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | hashed | 8 | 187 | 72.0 | 0.38× [0.38, 0.39] | -160.3% | [-163.6%, -157.0%] | 3.1 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 9.00 ms | 7.68 ms | 0.86× [0.85, 0.88] | -15.7% | [-18.0%, -13.4%] | 2.5 pts | 2.1 | no | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 8.88 ms | 11.72 ms | 1.33× [1.32, 1.34] | +24.7% | [+24.1%, +25.4%] | 0.7 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | hashed | 8 | 8.91 ms | 3.66 ms | 0.42× [0.41, 0.42] | -139.4% | [-141.2%, -137.7%] | 1.8 pts | 0.5 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 6 | 157 | 121 | 0.77× [0.76, 0.79] | -29.4% | [-31.6%, -27.2%] | 2.1 pts | 1.6 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 6 | 163 | 234 | 1.45× [1.43, 1.47] | +31.0% | [+30.0%, +32.1%] | 1.0 pts | 1.3 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | hashed | 6 | 159 | 42.4 | 0.27× [0.26, 0.29] | -264.7% | [-282.3%, -247.1%] | 16.7 pts | 2.8 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 7293 | 3766 | 0.51× [0.51, 0.52] | -94.3% | [-95.0%, -93.6%] | 0.6 pts | 0.4 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 6 | 7314 | 11.2 µs | 1.54× [1.51, 1.56] | +34.9% | [+34.0%, +35.9%] | 0.9 pts | 1.1 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | hashed | 6 | 9060 | 249.5 µs | 26.62× [24.50, 29.13] | +96.2% | [+95.9%, +96.6%] | 0.3 pts | 2.8 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 6 | 36.2 µs | 17.6 µs | 0.47× [0.47, 0.48] | -110.9% | [-114.2%, -107.6%] | 3.1 pts | 0.9 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 6 | 34.3 µs | 57.9 µs | 1.73× [1.71, 1.75] | +42.2% | [+41.4%, +43.0%] | 0.8 pts | 1.3 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | hashed | 6 | 40.1 µs | 263.2 µs | 6.52× [6.23, 6.83] | +84.7% | [+84.0%, +85.4%] | 0.7 pts | 0.7 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 6 | 271 | 233 | 0.87× [0.86, 0.87] | -15.4% | [-16.3%, -14.4%] | 0.9 pts | 0.9 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 6 | 268 | 365 | 1.35× [1.34, 1.36] | +26.0% | [+25.6%, +26.4%] | 0.4 pts | 0.8 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | hashed | 6 | 271 | 107 | 0.40× [0.39, 0.42] | -149.1% | [-159.4%, -138.8%] | 9.8 pts | 2.5 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 6 | 48.95 ms | 43.33 ms | 0.89× [0.88, 0.90] | -12.4% | [-13.9%, -10.9%] | 1.4 pts | 1.2 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 6 | 48.48 ms | 66.93 ms | 1.39× [1.37, 1.40] | +27.9% | [+26.9%, +28.8%] | 0.9 pts | 1.0 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | hashed | 6 | 48.50 ms | 20.67 ms | 0.42× [0.41, 0.44] | -136.3% | [-143.6%, -128.9%] | 7.0 pts | 1.9 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 276 | 247 | 0.90× [0.88, 0.91] | -11.3% | [-13.1%, -9.4%] | 3.2 pts | 2.5 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 319 | 495 | 1.59× [1.55, 1.62] | +37.0% | [+35.7%, +38.2%] | 1.8 pts | 2.2 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | hashed | 8 | 274 | 128 | 0.46× [0.45, 0.48] | -116.2% | [-124.0%, -108.4%] | 8.0 pts | 3.3 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 11.6 µs | 7502 | 0.66× [0.65, 0.68] | -51.2% | [-54.3%, -48.0%] | 4.6 pts | 1.9 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 12.9 µs | 24.3 µs | 1.91× [1.87, 1.95] | +47.7% | [+46.6%, +48.8%] | 1.5 pts | 2.2 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 255.6 µs | 133.7 µs | 0.54× [0.51, 0.57] | -85.0% | [-95.7%, -74.2%] | 11.5 pts | 0.9 | no | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 256.7 µs | 520.5 µs | 2.11× [1.98, 2.27] | +52.7% | [+49.5%, +55.9%] | 3.7 pts | 1.2 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 520 | 472 | 0.90× [0.87, 0.94] | -10.6% | [-15.1%, -6.1%] | 5.2 pts | 0.9 | no | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 609 | 703 | 1.16× [1.14, 1.19] | +14.0% | [+12.3%, +15.7%] | 1.6 pts | 1.5 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | hashed | 8 | 511 | 257 | 0.50× [0.48, 0.52] | -99.9% | [-107.9%, -91.8%] | 8.5 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 valuesFor: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.48% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=4096 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesFor: ordered vs hashed: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs hashed: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +7.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +3.99% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of -3.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 churn: ordered vs hashed: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs hashed: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -3.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -3.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
