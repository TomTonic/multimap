| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 64.4 | 54.7 | 0.85× [0.83, 0.86] | -18.1% | [-20.1%, -16.1%] | 2.5 pts | 2.2 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 65.3 | 140 | 2.12× [2.08, 2.16] | +52.9% | [+52.0%, +53.8%] | 0.9 pts | 2.1 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 1868 | 2597 | 1.39× [1.37, 1.41] | +28.2% | [+27.1%, +29.3%] | 1.0 pts | 0.7 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 1870 | 4617 | 2.48× [2.45, 2.51] | +59.7% | [+59.2%, +60.2%] | 1.1 pts | 2.1 | yes | yes |
| natural | street | 4096 | prefix | ordered | baseline | 8 | 254 | 258 | 1.01× [1.00, 1.03] | +1.3% | [+0.1%, +2.5%] | 1.4 pts | 1.3 | yes | yes |
| natural | street | 4096 | prefix | ordered | btree-sets | 8 | 255 | 528 | 2.07× [2.04, 2.11] | +51.8% | [+51.0%, +52.5%] | 0.8 pts | 1.7 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 160 | 105 | 0.66× [0.65, 0.66] | -52.6% | [-54.2%, -51.0%] | 1.6 pts | 0.9 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 161 | 189 | 1.18× [1.16, 1.19] | +14.9% | [+13.7%, +16.2%] | 1.4 pts | 1.6 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 6.26 ms | 4.18 ms | 0.67× [0.66, 0.68] | -50.2% | [-52.4%, -48.0%] | 2.3 pts | 2.0 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 6.32 ms | 7.79 ms | 1.22× [1.20, 1.23] | +17.7% | [+16.9%, +18.5%] | 1.1 pts | 1.3 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 81.2 | 77.8 | 0.95× [0.93, 0.98] | -4.9% | [-8.1%, -1.7%] | 3.5 pts | 2.5 | no | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 82.3 | 188 | 2.44× [2.25, 2.68] | +59.1% | [+55.5%, +62.7%] | 6.2 pts | 14.6 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 2250 | 3247 | 1.46× [1.42, 1.50] | +31.4% | [+29.5%, +33.3%] | 2.0 pts | 1.9 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 2236 | 5486 | 2.80× [2.39, 3.37] | +64.2% | [+58.1%, +70.4%] | 9.0 pts | 26.8 | yes | yes |
| natural | street | 16384 | prefix | ordered | baseline | 8 | 714 | 955 | 1.35× [1.33, 1.37] | +26.0% | [+25.0%, +26.9%] | 1.0 pts | 0.7 | yes | yes |
| natural | street | 16384 | prefix | ordered | btree-sets | 8 | 716 | 2080 | 3.09× [2.88, 3.32] | +67.6% | [+65.3%, +69.9%] | 4.1 pts | 8.8 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 199 | 148 | 0.74× [0.73, 0.75] | -34.9% | [-36.2%, -33.6%] | 1.7 pts | 1.0 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 209 | 277 | 1.31× [1.28, 1.35] | +23.8% | [+21.8%, +25.8%] | 2.3 pts | 2.2 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 30.85 ms | 22.95 ms | 0.75× [0.72, 0.77] | -33.8% | [-38.3%, -29.2%] | 4.7 pts | 5.3 | no | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 31.18 ms | 43.19 ms | 1.39× [1.35, 1.43] | +27.9% | [+25.7%, +30.1%] | 2.0 pts | 4.6 | yes | yes |
| natural | street | 212449 | valuesFor | ordered | baseline | 8 | 162 | 202 | 1.29× [1.22, 1.36] | +22.2% | [+18.0%, +26.5%] | 6.3 pts | 6.3 | no | yes |
| natural | street | 212449 | valuesFor | ordered | btree-sets | 8 | 198 | 516 | 2.64× [2.47, 2.83] | +62.1% | [+59.5%, +64.7%] | 3.4 pts | 4.7 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | baseline | 8 | 3337 | 5828 | 1.75× [1.65, 1.86] | +42.9% | [+39.5%, +46.2%] | 4.9 pts | 7.2 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 3843 | 17.1 µs | 4.46× [4.37, 4.56] | +77.6% | [+77.1%, +78.1%] | 0.7 pts | 2.7 | yes | yes |
| natural | street | 212449 | prefix | ordered | baseline | 8 | 10.2 µs | 19.3 µs | 1.88× [1.80, 1.97] | +46.8% | [+44.4%, +49.2%] | 5.0 pts | 8.4 | yes | yes |
| natural | street | 212449 | prefix | ordered | btree-sets | 8 | 12.2 µs | 58.9 µs | 4.89× [4.69, 5.10] | +79.5% | [+78.7%, +80.4%] | 1.1 pts | 2.4 | yes | yes |
| natural | street | 212449 | churn | ordered | baseline | 8 | 463 | 491 | 1.06× [1.01, 1.10] | +5.4% | [+1.4%, +9.3%] | 3.7 pts | 0.9 | no | yes |
| natural | street | 212449 | churn | ordered | btree-sets | 8 | 545 | 753 | 1.41× [1.37, 1.44] | +28.9% | [+27.0%, +30.7%] | 2.0 pts | 3.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 14.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 26.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 prefix: ordered vs btree-sets: the processes scatter 8.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs baseline: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs btree-sets: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs baseline: the processes scatter 6.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs baseline: the processes scatter 7.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs baseline: the processes scatter 8.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 churn: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
