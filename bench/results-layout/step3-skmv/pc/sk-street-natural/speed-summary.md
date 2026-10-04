| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | street | 4096 | valuesFor | ordered | baseline | 8 | 82.8 | 54.4 | 0.66× [0.65, 0.67] | -51.7% | [-54.2%, -49.3%] | 2.5 pts | 1.6 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 83.6 | 142 | 1.68× [1.66, 1.70] | +40.5% | [+39.8%, +41.3%] | 1.0 pts | 2.0 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | hashed | 8 | 83.3 | 26.4 | 0.32× [0.31, 0.32] | -216.8% | [-221.3%, -212.3%] | 5.2 pts | 2.2 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 5217 | 2441 | 0.46× [0.46, 0.47] | -115.9% | [-117.8%, -114.0%] | 3.2 pts | 1.0 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 5272 | 7062 | 1.34× [1.32, 1.36] | +25.2% | [+24.1%, +26.3%] | 1.4 pts | 1.8 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | hashed | 8 | 5448 | 55.0 µs | 10.09× [9.98, 10.21] | +90.1% | [+90.0%, +90.2%] | 0.1 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 8 | 468 | 231 | 0.49× [0.49, 0.50] | -103.6% | [-105.9%, -101.2%] | 2.9 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 8 | 469 | 711 | 1.51× [1.50, 1.52] | +33.6% | [+33.1%, +34.1%] | 0.9 pts | 1.0 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | hashed | 8 | 540 | 50.5 µs | 93.09× [89.62, 96.84] | +98.9% | [+98.9%, +99.0%] | 0.0 pts | 1.0 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 8 | 129 | 108 | 0.84× [0.83, 0.84] | -19.2% | [-20.0%, -18.4%] | 1.3 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 8 | 129 | 200 | 1.55× [1.52, 1.58] | +35.4% | [+34.1%, +36.6%] | 1.2 pts | 1.7 | yes | yes |
| natural-str | street | 4096 | churn | ordered | hashed | 8 | 128 | 56.4 | 0.44× [0.43, 0.45] | -125.5% | [-130.5%, -120.4%] | 5.0 pts | 2.2 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 8 | 4.91 ms | 4.38 ms | 0.89× [0.88, 0.91] | -12.0% | [-13.8%, -10.1%] | 2.0 pts | 1.8 | yes | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 8 | 4.95 ms | 8.02 ms | 1.63× [1.60, 1.65] | +38.6% | [+37.6%, +39.6%] | 0.9 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | build | ordered | hashed | 8 | 4.94 ms | 2.76 ms | 0.57× [0.55, 0.59] | -76.4% | [-82.1%, -70.7%] | 5.6 pts | 1.6 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 6 | 105 | 75.6 | 0.72× [0.71, 0.73] | -39.6% | [-41.3%, -37.9%] | 1.6 pts | 1.4 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 6 | 107 | 190 | 1.75× [1.72, 1.78] | +42.8% | [+41.8%, +43.9%] | 1.0 pts | 2.1 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | hashed | 6 | 105 | 31.7 | 0.30× [0.30, 0.30] | -233.5% | [-236.5%, -230.6%] | 2.8 pts | 1.1 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 5855 | 2984 | 0.51× [0.49, 0.52] | -97.8% | [-102.2%, -93.4%] | 4.2 pts | 2.8 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 5859 | 8117 | 1.38× [1.37, 1.39] | +27.4% | [+27.0%, +27.9%] | 0.4 pts | 0.9 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | hashed | 6 | 6656 | 235.1 µs | 34.53× [33.48, 35.64] | +97.1% | [+97.0%, +97.2%] | 0.1 pts | 1.1 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 6 | 1882 | 876 | 0.47× [0.46, 0.48] | -114.9% | [-119.4%, -110.4%] | 4.2 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 6 | 1975 | 3005 | 1.53× [1.50, 1.55] | +34.4% | [+33.5%, +35.4%] | 0.9 pts | 0.7 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | hashed | 6 | 2479 | 232.2 µs | 92.46× [90.96, 94.02] | +98.9% | [+98.9%, +98.9%] | 0.0 pts | 0.4 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 6 | 176 | 149 | 0.85× [0.85, 0.86] | -17.0% | [-18.3%, -15.8%] | 1.2 pts | 1.3 | yes | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 6 | 181 | 284 | 1.58× [1.56, 1.59] | +36.6% | [+35.9%, +37.3%] | 0.6 pts | 0.9 | yes | yes |
| natural-str | street | 16384 | churn | ordered | hashed | 6 | 171 | 75.4 | 0.44× [0.43, 0.45] | -127.3% | [-131.6%, -122.9%] | 4.2 pts | 1.5 | yes | yes |
| natural-str | street | 16384 | build | ordered | baseline | 6 | 25.88 ms | 23.12 ms | 0.90× [0.89, 0.90] | -11.6% | [-12.3%, -11.0%] | 0.6 pts | 0.7 | yes | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 6 | 26.09 ms | 43.83 ms | 1.69× [1.67, 1.72] | +40.9% | [+40.0%, +41.8%] | 0.8 pts | 1.4 | yes | yes |
| natural-str | street | 16384 | build | ordered | hashed | 6 | 25.86 ms | 13.93 ms | 0.54× [0.53, 0.55] | -85.9% | [-89.2%, -82.7%] | 3.1 pts | 1.3 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | baseline | 8 | 266 | 224 | 0.85× [0.82, 0.89] | -17.4% | [-21.9%, -12.8%] | 4.7 pts | 3.6 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | btree-sets | 8 | 328 | 505 | 1.57× [1.51, 1.64] | +36.3% | [+33.6%, +39.1%] | 2.9 pts | 4.1 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | hashed | 8 | 260 | 114 | 0.45× [0.43, 0.46] | -124.7% | [-133.3%, -116.1%] | 9.7 pts | 3.7 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 10.2 µs | 7138 | 0.70× [0.69, 0.73] | -41.9% | [-45.9%, -37.9%] | 4.6 pts | 2.9 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 11.6 µs | 23.5 µs | 2.04× [1.98, 2.10] | +50.9% | [+49.4%, +52.5%] | 1.7 pts | 3.5 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | baseline | 8 | 37.7 µs | 21.9 µs | 0.60× [0.58, 0.61] | -67.6% | [-72.2%, -62.9%] | 4.6 pts | 1.6 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | btree-sets | 8 | 39.8 µs | 80.3 µs | 2.01× [1.93, 2.09] | +50.2% | [+48.2%, +52.2%] | 2.3 pts | 1.3 | yes | yes |
| natural-str | street | 212449 | churn | ordered | baseline | 8 | 536 | 493 | 0.89× [0.87, 0.91] | -12.8% | [-15.2%, -10.4%] | 4.2 pts | 0.8 | no | yes |
| natural-str | street | 212449 | churn | ordered | btree-sets | 8 | 613 | 744 | 1.22× [1.19, 1.25] | +18.1% | [+16.3%, +19.9%] | 2.2 pts | 1.1 | yes | yes |
| natural-str | street | 212449 | churn | ordered | hashed | 8 | 509 | 276 | 0.50× [0.48, 0.53] | -98.5% | [-107.3%, -89.8%] | 11.7 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 valuesFor: ordered vs hashed: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=4096 churn: ordered vs hashed: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of -0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs hashed: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
