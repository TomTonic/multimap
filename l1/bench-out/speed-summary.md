| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | street | 4096 | valuesFor | ordered | btree-map | 8 | 67.4 | 127 | 1.87× [1.86, 1.89] | +46.6% | [+46.1%, +47.1%] | 0.6 pts | 1.4 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | ordered-lpage | 8 | 67.4 | 127 | 1.88× [1.87, 1.90] | +46.9% | [+46.6%, +47.3%] | 0.4 pts | 1.3 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | ordered-lpage-zc | 8 | 67.1 | 107 | 1.60× [1.59, 1.61] | +37.5% | [+37.2%, +37.7%] | 0.3 pts | 0.8 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | btree-map | 8 | 2508 | 1601 | 0.64× [0.64, 0.65] | -56.0% | [-57.4%, -54.7%] | 1.6 pts | 0.8 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | ordered-lpage | 8 | 2503 | 1318 | 0.53× [0.52, 0.54] | -87.6% | [-91.0%, -84.1%] | 4.1 pts | 2.0 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | ordered-lpage-zc | 8 | 2489 | 1060 | 0.43× [0.43, 0.43] | -133.9% | [-134.9%, -133.0%] | 1.1 pts | 0.3 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | btree-map | 8 | 285 | 248 | 0.86× [0.85, 0.87] | -16.6% | [-17.8%, -15.4%] | 1.5 pts | 1.2 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | ordered-lpage | 8 | 285 | 273 | 0.96× [0.95, 0.97] | -4.1% | [-5.5%, -2.7%] | 1.7 pts | 1.5 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | ordered-lpage-zc | 8 | 284 | 235 | 0.83× [0.82, 0.83] | -20.6% | [-21.4%, -19.8%] | 1.0 pts | 0.9 | yes | yes |
| unique-str | street | 4096 | churn | ordered | btree-map | 8 | 159 | 188 | 1.17× [1.16, 1.19] | +14.9% | [+14.1%, +15.6%] | 0.9 pts | 1.1 | yes | yes |
| unique-str | street | 4096 | churn | ordered | ordered-lpage | 8 | 160 | 243 | 1.52× [1.51, 1.53] | +34.2% | [+34.0%, +34.5%] | 0.3 pts | 0.5 | yes | yes |
| unique-str | street | 4096 | churn | ordered | ordered-lpage-zc | 8 | 163 | 312 | 1.92× [1.90, 1.93] | +47.8% | [+47.4%, +48.2%] | 0.5 pts | 1.0 | yes | yes |
| unique-str | street | 4096 | build | ordered | btree-map | 8 | 2.18 ms | 2.75 ms | 1.26× [1.25, 1.27] | +20.8% | [+20.1%, +21.5%] | 0.9 pts | 4.0 | yes | yes |
| unique-str | street | 4096 | build | ordered | ordered-lpage | 8 | 2.19 ms | 3.97 ms | 1.81× [1.81, 1.82] | +44.8% | [+44.6%, +45.1%] | 0.3 pts | 1.1 | yes | yes |
| unique-str | street | 4096 | build | ordered | ordered-lpage-zc | 8 | 2.19 ms | 4.95 ms | 2.25× [2.24, 2.27] | +55.6% | [+55.3%, +55.9%] | 0.4 pts | 1.2 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | btree-map | 10 | 88.5 | 175 | 1.98× [1.96, 2.00] | +49.4% | [+48.9%, +50.0%] | 0.7 pts | 2.3 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | ordered-lpage | 10 | 87.9 | 149 | 1.68× [1.67, 1.70] | +40.7% | [+40.2%, +41.1%] | 0.5 pts | 1.5 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | ordered-lpage-zc | 10 | 87.6 | 129 | 1.47× [1.46, 1.48] | +32.0% | [+31.7%, +32.3%] | 0.4 pts | 1.1 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | btree-map | 10 | 2705 | 1737 | 0.64× [0.64, 0.65] | -55.8% | [-56.8%, -54.8%] | 1.2 pts | 1.9 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | ordered-lpage | 10 | 2694 | 1438 | 0.54× [0.53, 0.54] | -86.7% | [-88.9%, -84.5%] | 2.8 pts | 2.6 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | ordered-lpage-zc | 10 | 2690 | 1140 | 0.42× [0.42, 0.42] | -136.5% | [-137.4%, -135.6%] | 1.2 pts | 0.6 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | btree-map | 10 | 930 | 720 | 0.77× [0.76, 0.78] | -29.7% | [-30.8%, -28.7%] | 1.6 pts | 0.9 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | ordered-lpage | 10 | 919 | 549 | 0.59× [0.59, 0.60] | -68.1% | [-70.0%, -66.2%] | 2.4 pts | 1.2 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | ordered-lpage-zc | 10 | 907 | 438 | 0.48× [0.48, 0.48] | -107.8% | [-108.5%, -107.0%] | 1.0 pts | 0.4 | yes | yes |
| unique-str | street | 16384 | churn | ordered | btree-map | 10 | 204 | 244 | 1.17× [1.14, 1.20] | +14.6% | [+12.6%, +16.5%] | 2.4 pts | 3.4 | yes | yes |
| unique-str | street | 16384 | churn | ordered | ordered-lpage | 10 | 204 | 275 | 1.33× [1.29, 1.37] | +24.7% | [+22.6%, +26.8%] | 2.6 pts | 4.7 | yes | yes |
| unique-str | street | 16384 | churn | ordered | ordered-lpage-zc | 10 | 231 | 356 | 1.53× [1.49, 1.57] | +34.6% | [+33.1%, +36.2%] | 2.1 pts | 2.5 | yes | yes |
| unique-str | street | 16384 | build | ordered | btree-map | 10 | 10.86 ms | 14.23 ms | 1.31× [1.30, 1.32] | +23.5% | [+22.8%, +24.2%] | 1.0 pts | 3.5 | yes | yes |
| unique-str | street | 16384 | build | ordered | ordered-lpage | 10 | 10.86 ms | 17.92 ms | 1.65× [1.65, 1.66] | +39.5% | [+39.2%, +39.7%] | 0.4 pts | 1.5 | yes | yes |
| unique-str | street | 16384 | build | ordered | ordered-lpage-zc | 10 | 10.96 ms | 22.87 ms | 2.09× [2.07, 2.12] | +52.2% | [+51.7%, +52.8%] | 0.8 pts | 1.6 | yes | yes |
| unique-str | street | 212449 | valuesFor | ordered | btree-map | 24 | 270 | 389 | 1.50× [1.48, 1.51] | +33.1% | [+32.4%, +33.9%] | 2.0 pts | 1.1 | yes | yes |
| unique-str | street | 212449 | valuesFor | ordered | ordered-lpage | 24 | 236 | 229 | 0.98× [0.97, 0.99] | -2.4% | [-3.4%, -1.4%] | 5.1 pts | 2.6 | yes | yes |
| unique-str | street | 212449 | valuesFor | ordered | ordered-lpage-zc | 24 | 225 | 208 | 0.91× [0.90, 0.91] | -10.2% | [-10.9%, -9.6%] | 5.0 pts | 2.7 | yes | yes |
| unique-str | street | 212449 | valuesBetween | ordered | btree-map | 24 | 4169 | 2560 | 0.62× [0.60, 0.63] | -62.3% | [-65.4%, -59.2%] | 4.9 pts | 1.4 | yes | yes |
| unique-str | street | 212449 | valuesBetween | ordered | ordered-lpage | 24 | 3991 | 1808 | 0.43× [0.43, 0.44] | -130.1% | [-133.2%, -127.1%] | 8.9 pts | 2.5 | yes | yes |
| unique-str | street | 212449 | valuesBetween | ordered | ordered-lpage-zc | 24 | 3818 | 1380 | 0.35× [0.35, 0.35] | -184.7% | [-187.4%, -182.1%] | 10.9 pts | 2.1 | yes | yes |
| unique-str | street | 212449 | prefix | ordered | btree-map | 24 | 15.8 µs | 9585 | 0.60× [0.59, 0.61] | -65.8% | [-68.2%, -63.4%] | 3.4 pts | 1.0 | yes | yes |
| unique-str | street | 212449 | prefix | ordered | ordered-lpage | 24 | 16.4 µs | 5918 | 0.36× [0.35, 0.36] | -180.0% | [-182.5%, -177.5%] | 7.2 pts | 1.6 | yes | yes |
| unique-str | street | 212449 | prefix | ordered | ordered-lpage-zc | 24 | 14.5 µs | 4053 | 0.27× [0.27, 0.27] | -267.5% | [-271.4%, -263.7%] | 11.4 pts | 0.7 | yes | yes |
| unique-str | street | 212449 | churn | ordered | btree-map | 24 | 589 | 646 | 1.08× [1.08, 1.09] | +7.8% | [+7.0%, +8.6%] | 2.7 pts | 1.6 | yes | yes |
| unique-str | street | 212449 | churn | ordered | ordered-lpage | 24 | 524 | 461 | 0.85× [0.83, 0.88] | -17.0% | [-20.5%, -13.6%] | 4.5 pts | 0.6 | no | yes |
| unique-str | street | 212449 | churn | ordered | ordered-lpage-zc | 24 | 567 | 570 | 0.96× [0.94, 0.98] | -4.2% | [-6.5%, -2.0%] | 3.6 pts | 0.6 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str street n=4096 valuesBetween: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=4096 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.60% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=4096 prefix: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=4096 build: ordered vs btree-map: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of +0.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=16384 valuesBetween: ordered vs ordered-lpage: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 churn: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 churn: ordered vs ordered-lpage: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 churn: ordered vs ordered-lpage-zc: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=16384 build: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesFor: ordered vs ordered-lpage: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesBetween: ordered vs ordered-lpage: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesBetween: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=212449 valuesBetween: ordered vs ordered-lpage-zc: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -2.93% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=212449 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=212449 churn: ordered vs ordered-lpage: the A/A validations found a systematic difference of -0.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
