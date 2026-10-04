| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | street | 4096 | valuesFor | ordered | baseline | 6 | 81.1 | 54.3 | 0.66× [0.65, 0.68] | -50.7% | [-54.1%, -47.3%] | 3.2 pts | 2.4 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 6 | 82.2 | 138 | 1.70× [1.66, 1.73] | +41.0% | [+39.8%, +42.3%] | 1.2 pts | 1.9 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | hashed | 6 | 81.2 | 24.6 | 0.30× [0.29, 0.31] | -231.7% | [-241.1%, -222.3%] | 8.9 pts | 3.6 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 6 | 5050 | 2449 | 0.48× [0.47, 0.50] | -106.5% | [-111.5%, -101.4%] | 4.8 pts | 1.5 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 5068 | 6977 | 1.38× [1.35, 1.40] | +27.3% | [+26.1%, +28.5%] | 1.1 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | hashed | 6 | 5248 | 54.5 µs | 10.46× [10.32, 10.59] | +90.4% | [+90.3%, +90.6%] | 0.1 pts | 1.2 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 6 | 460 | 235 | 0.51× [0.51, 0.52] | -95.7% | [-97.7%, -93.7%] | 1.9 pts | 0.7 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 6 | 466 | 702 | 1.52× [1.51, 1.53] | +34.2% | [+33.7%, +34.7%] | 0.5 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | hashed | 6 | 545 | 50.2 µs | 92.46× [91.15, 93.82] | +98.9% | [+98.9%, +98.9%] | 0.0 pts | 0.4 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 6 | 129 | 109 | 0.85× [0.84, 0.85] | -18.3% | [-19.5%, -17.2%] | 1.1 pts | 0.7 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 6 | 127 | 199 | 1.56× [1.55, 1.58] | +36.0% | [+35.3%, +36.7%] | 0.6 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | churn | ordered | hashed | 6 | 126 | 56.4 | 0.45× [0.44, 0.46] | -122.1% | [-126.6%, -117.6%] | 4.3 pts | 2.1 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 6 | 4.80 ms | 4.35 ms | 0.91× [0.90, 0.92] | -10.3% | [-11.5%, -9.1%] | 1.1 pts | 1.0 | yes | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 6 | 4.87 ms | 8.05 ms | 1.66× [1.62, 1.70] | +39.6% | [+38.2%, +41.0%] | 1.3 pts | 1.2 | yes | yes |
| natural-str | street | 4096 | build | ordered | hashed | 6 | 4.77 ms | 2.70 ms | 0.57× [0.56, 0.58] | -75.7% | [-77.8%, -73.6%] | 2.0 pts | 0.7 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 8 | 104 | 75.2 | 0.72× [0.70, 0.74] | -39.2% | [-42.9%, -35.5%] | 3.5 pts | 3.1 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 8 | 105 | 188 | 1.78× [1.75, 1.81] | +43.8% | [+42.8%, +44.7%] | 1.1 pts | 2.3 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | hashed | 8 | 104 | 29.9 | 0.29× [0.29, 0.29] | -247.8% | [-250.3%, -245.2%] | 2.4 pts | 1.0 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 5738 | 3010 | 0.53× [0.52, 0.53] | -90.2% | [-92.6%, -87.7%] | 2.4 pts | 1.5 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 5717 | 8062 | 1.41× [1.40, 1.42] | +29.3% | [+28.8%, +29.8%] | 0.7 pts | 1.4 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | hashed | 8 | 6577 | 233.2 µs | 35.15× [34.19, 36.17] | +97.2% | [+97.1%, +97.2%] | 0.1 pts | 1.8 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 8 | 1868 | 880 | 0.48× [0.47, 0.48] | -110.2% | [-112.7%, -107.7%] | 2.5 pts | 0.7 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 8 | 1902 | 3006 | 1.55× [1.51, 1.59] | +35.4% | [+33.9%, +36.9%] | 1.5 pts | 1.3 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | hashed | 8 | 2427 | 230.3 µs | 93.82× [90.95, 96.88] | +98.9% | [+98.9%, +99.0%] | 0.0 pts | 0.8 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 8 | 173 | 149 | 0.86× [0.85, 0.88] | -15.7% | [-17.7%, -13.7%] | 1.9 pts | 1.6 | no | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 8 | 175 | 280 | 1.58× [1.55, 1.61] | +36.6% | [+35.4%, +37.9%] | 1.2 pts | 2.1 | yes | yes |
| natural-str | street | 16384 | churn | ordered | hashed | 8 | 169 | 76.1 | 0.44× [0.44, 0.45] | -125.3% | [-126.7%, -123.9%] | 3.2 pts | 1.0 | yes | yes |
| natural-str | street | 16384 | build | ordered | baseline | 8 | 25.57 ms | 23.27 ms | 0.91× [0.90, 0.92] | -10.0% | [-11.1%, -9.0%] | 1.1 pts | 1.1 | yes | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 8 | 25.63 ms | 44.03 ms | 1.72× [1.70, 1.73] | +41.8% | [+41.2%, +42.3%] | 0.7 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | build | ordered | hashed | 8 | 25.48 ms | 13.79 ms | 0.54× [0.53, 0.55] | -84.4% | [-87.2%, -81.7%] | 2.8 pts | 1.2 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | baseline | 8 | 264 | 225 | 0.85× [0.83, 0.87] | -17.6% | [-20.0%, -15.2%] | 2.3 pts | 1.7 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | btree-sets | 8 | 316 | 498 | 1.58× [1.52, 1.63] | +36.5% | [+34.3%, +38.7%] | 2.4 pts | 2.5 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | hashed | 8 | 254 | 110 | 0.44× [0.43, 0.44] | -129.1% | [-131.5%, -126.7%] | 4.7 pts | 2.0 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 9961 | 7081 | 0.70× [0.69, 0.72] | -42.6% | [-45.8%, -39.5%] | 5.7 pts | 3.7 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 11.4 µs | 23.9 µs | 2.09× [2.05, 2.14] | +52.2% | [+51.2%, +53.3%] | 1.5 pts | 3.0 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | baseline | 8 | 37.1 µs | 21.8 µs | 0.60× [0.58, 0.62] | -66.9% | [-73.2%, -60.5%] | 7.7 pts | 2.5 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | btree-sets | 8 | 39.1 µs | 79.4 µs | 2.02× [2.00, 2.03] | +50.4% | [+50.0%, +50.8%] | 0.4 pts | 0.3 | yes | yes |
| natural-str | street | 212449 | churn | ordered | baseline | 8 | 522 | 483 | 0.91× [0.88, 0.94] | -9.9% | [-13.9%, -6.0%] | 4.1 pts | 0.8 | no | yes |
| natural-str | street | 212449 | churn | ordered | btree-sets | 8 | 626 | 761 | 1.23× [1.21, 1.25] | +18.8% | [+17.4%, +20.2%] | 1.6 pts | 1.4 | yes | yes |
| natural-str | street | 212449 | churn | ordered | hashed | 8 | 530 | 283 | 0.52× [0.50, 0.54] | -92.7% | [-98.9%, -86.5%] | 6.7 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 valuesFor: ordered vs hashed: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of -0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=4096 churn: ordered vs hashed: the A/A validations found a systematic difference of -0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=4096 churn: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +2.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=16384 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs hashed: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
