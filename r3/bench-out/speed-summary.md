| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 77.9 | 184 | 2.37× [2.35, 2.39] | +57.8% | [+57.4%, +58.2%] | 0.5 pts | 1.4 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | hashed | 8 | 75.5 | 34.0 | 0.45× [0.45, 0.45] | -122.0% | [-123.6%, -120.5%] | 1.8 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | map-sets | 8 | 75.6 | 79.3 | 1.05× [1.04, 1.05] | +4.3% | [+3.8%, +4.9%] | 0.6 pts | 1.6 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 3269 | 5398 | 1.65× [1.65, 1.66] | +39.6% | [+39.3%, +39.9%] | 0.4 pts | 0.8 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | hashed | 8 | 3319 | 73.3 µs | 21.90× [21.59, 22.21] | +95.4% | [+95.4%, +95.5%] | 0.1 pts | 2.7 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | map-sets | 8 | 3347 | 71.0 µs | 21.03× [20.73, 21.35] | +95.2% | [+95.2%, +95.3%] | 0.1 pts | 2.3 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | btree-sets | 8 | 321 | 659 | 2.06× [2.04, 2.08] | +51.5% | [+51.1%, +51.9%] | 0.5 pts | 0.8 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | hashed | 8 | 326 | 71.2 µs | 215.18× [207.80, 223.10] | +99.5% | [+99.5%, +99.6%] | 0.0 pts | 2.7 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | map-sets | 8 | 330 | 65.0 µs | 195.77× [189.30, 202.70] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 2.3 | yes | yes |
| multi-str | street | 4096 | churn | ordered | btree-sets | 8 | 137 | 250 | 1.84× [1.81, 1.86] | +45.6% | [+44.8%, +46.3%] | 0.9 pts | 1.8 | yes | yes |
| multi-str | street | 4096 | churn | ordered | hashed | 8 | 136 | 70.8 | 0.52× [0.52, 0.53] | -91.7% | [-93.0%, -90.5%] | 1.5 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | churn | ordered | map-sets | 8 | 136 | 83.7 | 0.62× [0.61, 0.63] | -60.7% | [-63.8%, -57.6%] | 3.7 pts | 3.6 | yes | yes |
| multi-str | street | 4096 | build | ordered | btree-sets | 8 | 5.45 ms | 10.27 ms | 1.88× [1.86, 1.89] | +46.7% | [+46.3%, +47.2%] | 0.5 pts | 2.1 | yes | yes |
| multi-str | street | 4096 | build | ordered | hashed | 8 | 5.46 ms | 3.33 ms | 0.62× [0.61, 0.62] | -62.4% | [-64.8%, -60.1%] | 2.8 pts | 3.3 | yes | yes |
| multi-str | street | 4096 | build | ordered | map-sets | 8 | 5.45 ms | 4.15 ms | 0.76× [0.74, 0.77] | -32.3% | [-34.7%, -29.8%] | 2.9 pts | 3.6 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | btree-sets | 10 | 100 | 235 | 2.37× [2.33, 2.41] | +57.8% | [+57.1%, +58.5%] | 0.9 pts | 1.3 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | hashed | 10 | 96.7 | 36.7 | 0.38× [0.37, 0.40] | -159.8% | [-167.9%, -151.8%] | 10.1 pts | 5.3 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | map-sets | 10 | 99.7 | 90.3 | 0.92× [0.90, 0.94] | -8.6% | [-10.6%, -6.7%] | 2.4 pts | 1.5 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | btree-sets | 10 | 3612 | 5884 | 1.65× [1.62, 1.69] | +39.5% | [+38.2%, +40.7%] | 1.5 pts | 3.7 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | hashed | 10 | 3596 | 285.1 µs | 77.99× [76.85, 79.17] | +98.7% | [+98.7%, +98.7%] | 0.0 pts | 3.5 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | map-sets | 10 | 3688 | 267.5 µs | 72.15× [71.13, 73.19] | +98.6% | [+98.6%, +98.6%] | 0.0 pts | 1.6 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | btree-sets | 10 | 1190 | 2312 | 1.96× [1.94, 1.97] | +48.9% | [+48.5%, +49.3%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | hashed | 10 | 1248 | 290.4 µs | 228.81× [220.89, 237.32] | +99.6% | [+99.5%, +99.6%] | 0.0 pts | 2.1 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | map-sets | 10 | 1264 | 269.0 µs | 211.28× [205.44, 217.46] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 1.7 | yes | yes |
| multi-str | street | 16384 | churn | ordered | btree-sets | 10 | 190 | 323 | 1.71× [1.70, 1.73] | +41.7% | [+41.1%, +42.2%] | 0.8 pts | 0.6 | yes | yes |
| multi-str | street | 16384 | churn | ordered | hashed | 10 | 175 | 83.6 | 0.48× [0.47, 0.48] | -110.5% | [-112.6%, -108.4%] | 2.5 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | churn | ordered | map-sets | 10 | 183 | 111 | 0.62× [0.61, 0.63] | -61.9% | [-64.5%, -59.3%] | 3.1 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | build | ordered | btree-sets | 10 | 27.21 ms | 49.85 ms | 1.83× [1.82, 1.84] | +45.3% | [+45.1%, +45.5%] | 0.4 pts | 3.2 | yes | yes |
| multi-str | street | 16384 | build | ordered | hashed | 10 | 27.25 ms | 14.55 ms | 0.53× [0.53, 0.54] | -87.5% | [-89.2%, -85.7%] | 2.3 pts | 3.9 | yes | yes |
| multi-str | street | 16384 | build | ordered | map-sets | 10 | 27.27 ms | 18.33 ms | 0.67× [0.66, 0.68] | -49.3% | [-51.1%, -47.5%] | 2.4 pts | 4.0 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | btree-sets | 22 | 300 | 654 | 2.15× [2.11, 2.20] | +53.6% | [+52.6%, +54.5%] | 1.9 pts | 1.7 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | hashed | 22 | 258 | 107 | 0.41× [0.41, 0.42] | -142.3% | [-146.3%, -138.4%] | 8.3 pts | 1.9 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | map-sets | 22 | 285 | 247 | 0.87× [0.85, 0.88] | -15.2% | [-17.2%, -13.3%] | 5.0 pts | 2.7 | yes | yes |
| multi-str | street | 212449 | valuesBetween | ordered | btree-sets | 22 | 6702 | 13.5 µs | 2.01× [2.01, 2.01] | +50.3% | [+50.2%, +50.3%] | 1.0 pts | 1.6 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | btree-sets | 22 | 23.6 µs | 51.5 µs | 2.16× [2.13, 2.20] | +53.7% | [+53.0%, +54.5%] | 1.4 pts | 0.6 | yes | yes |
| multi-str | street | 212449 | churn | ordered | btree-sets | 22 | 626 | 875 | 1.38× [1.37, 1.39] | +27.6% | [+26.8%, +28.3%] | 1.7 pts | 1.0 | yes | yes |
| multi-str | street | 212449 | churn | ordered | hashed | 22 | 561 | 256 | 0.44× [0.42, 0.46] | -127.6% | [-137.7%, -117.4%] | 14.1 pts | 0.5 | yes | yes |
| multi-str | street | 212449 | churn | ordered | map-sets | 22 | 580 | 355 | 0.59× [0.58, 0.61] | -68.6% | [-72.4%, -64.9%] | 6.5 pts | 0.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesBetween: ordered vs hashed: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 valuesBetween: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 prefix: ordered vs hashed: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +0.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 prefix: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 churn: ordered vs map-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 build: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 build: ordered vs hashed: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 build: ordered vs map-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +0.60% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesFor: ordered vs hashed: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs hashed: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 prefix: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.67% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 build: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 build: ordered vs hashed: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 build: ordered vs map-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 valuesFor: ordered vs hashed: the A/A validations found a systematic difference of +0.84% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 valuesFor: ordered vs map-sets: the A/A validations found a systematic difference of +0.72% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 valuesFor: ordered vs map-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.97% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 churn: ordered vs hashed: the A/A validations found a systematic difference of +6.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 churn: ordered vs map-sets: the A/A validations found a systematic difference of +2.66% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
