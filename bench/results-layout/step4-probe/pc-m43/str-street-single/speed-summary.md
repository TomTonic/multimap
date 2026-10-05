| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 8 | 71.7 | 66.4 | 0.92× [0.91, 0.93] | -8.6% | [-10.0%, -7.2%] | 1.5 pts | 2.1 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 8 | 73.4 | 92.1 | 1.25× [1.22, 1.28] | +19.7% | [+17.7%, +21.7%] | 2.2 pts | 2.7 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 2551 | 3702 | 1.44× [1.43, 1.46] | +30.7% | [+30.1%, +31.3%] | 0.6 pts | 0.5 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 8 | 2515 | 1129 | 0.44× [0.43, 0.44] | -128.9% | [-130.9%, -126.8%] | 2.5 pts | 0.7 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 8 | 363 | 372 | 1.03× [1.02, 1.04] | +3.0% | [+1.8%, +4.2%] | 1.2 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 8 | 359 | 188 | 0.52× [0.51, 0.52] | -92.4% | [-94.2%, -90.6%] | 2.0 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 8 | 182 | 126 | 0.69× [0.69, 0.69] | -45.4% | [-45.9%, -44.9%] | 0.6 pts | 0.3 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 8 | 183 | 140 | 0.76× [0.76, 0.77] | -31.1% | [-31.7%, -30.4%] | 0.8 pts | 0.6 | yes | yes |
| single-value-str | street | 4096 | build | ordered | baseline | 8 | 2.78 ms | 1.82 ms | 0.66× [0.65, 0.67] | -52.2% | [-54.3%, -50.2%] | 2.0 pts | 1.2 | yes | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 8 | 2.80 ms | 2.13 ms | 0.76× [0.76, 0.77] | -30.8% | [-32.3%, -29.2%] | 2.5 pts | 1.6 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 6 | 84.4 | 86.5 | 1.02× [1.01, 1.04] | +2.2% | [+0.7%, +3.6%] | 1.4 pts | 2.0 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 6 | 85.6 | 126 | 1.48× [1.43, 1.53] | +32.3% | [+30.0%, +34.5%] | 2.1 pts | 5.0 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 2853 | 4193 | 1.47× [1.45, 1.49] | +32.0% | [+31.2%, +32.8%] | 0.8 pts | 1.8 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 6 | 2856 | 1372 | 0.48× [0.48, 0.48] | -108.7% | [-110.1%, -107.4%] | 1.3 pts | 1.0 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 6 | 940 | 1338 | 1.43× [1.42, 1.45] | +30.2% | [+29.4%, +30.9%] | 0.7 pts | 0.7 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 6 | 926 | 491 | 0.53× [0.52, 0.54] | -90.1% | [-93.5%, -86.8%] | 3.2 pts | 1.2 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 6 | 207 | 163 | 0.80× [0.79, 0.81] | -25.4% | [-27.0%, -23.9%] | 1.5 pts | 0.7 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 6 | 208 | 194 | 0.94× [0.93, 0.95] | -6.4% | [-7.8%, -5.1%] | 1.3 pts | 0.8 | yes | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 6 | 11.79 ms | 8.84 ms | 0.75× [0.74, 0.76] | -33.8% | [-35.4%, -32.3%] | 1.5 pts | 1.5 | yes | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 6 | 12.18 ms | 11.19 ms | 0.93× [0.92, 0.95] | -7.2% | [-8.7%, -5.6%] | 1.5 pts | 0.6 | yes | yes |
| single-value-str | street | 212449 | valuesFor | ordered | baseline | 8 | 140 | 188 | 1.34× [1.30, 1.40] | +25.6% | [+22.8%, +28.3%] | 3.1 pts | 3.5 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | btree-map | 8 | 160 | 244 | 1.50× [1.43, 1.57] | +33.3% | [+30.3%, +36.4%] | 3.6 pts | 3.4 | yes | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 3346 | 5777 | 1.73× [1.71, 1.76] | +42.2% | [+41.4%, +43.0%] | 1.3 pts | 1.6 | yes | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | btree-map | 8 | 3205 | 1878 | 0.58× [0.55, 0.61] | -72.3% | [-80.9%, -63.6%] | 8.4 pts | 4.2 | no | yes |
| single-value-str | street | 212449 | prefix | ordered | baseline | 8 | 12.4 µs | 23.1 µs | 1.87× [1.83, 1.91] | +46.4% | [+45.2%, +47.6%] | 1.4 pts | 2.9 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | btree-map | 8 | 12.7 µs | 7538 | 0.60× [0.56, 0.65] | -66.5% | [-78.6%, -54.3%] | 11.3 pts | 6.6 | no | yes |
| single-value-str | street | 212449 | churn | ordered | baseline | 8 | 382 | 420 | 1.11× [1.09, 1.13] | +9.7% | [+8.3%, +11.2%] | 1.7 pts | 1.1 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | btree-map | 8 | 466 | 493 | 1.09× [1.06, 1.12] | +8.1% | [+5.8%, +10.4%] | 2.4 pts | 1.7 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.49% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs btree-map: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
