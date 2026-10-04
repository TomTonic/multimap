| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | street | 4096 | valuesFor | ordered | btree-map | 12 | 48.1 | 93.5 | 1.95× [1.92, 1.98] | +48.7% | [+48.0%, +49.4%] | 0.8 pts | 2.1 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | ordered-lpage | 12 | 48.1 | 92.4 | 1.93× [1.90, 1.96] | +48.2% | [+47.4%, +49.1%] | 1.0 pts | 3.2 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | ordered-lpage-zc | 12 | 48.0 | 76.0 | 1.59× [1.57, 1.61] | +37.0% | [+36.2%, +37.8%] | 1.0 pts | 2.5 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | btree-map | 12 | 1763 | 1154 | 0.64× [0.63, 0.64] | -57.5% | [-58.9%, -56.0%] | 2.4 pts | 0.8 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | ordered-lpage | 12 | 1765 | 1095 | 0.62× [0.62, 0.62] | -60.9% | [-61.7%, -60.0%] | 1.6 pts | 0.8 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | ordered-lpage-zc | 12 | 1763 | 783 | 0.44× [0.44, 0.45] | -125.9% | [-127.9%, -123.8%] | 2.9 pts | 1.2 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | btree-map | 12 | 197 | 191 | 0.96× [0.94, 0.97] | -4.7% | [-6.1%, -3.3%] | 1.4 pts | 1.0 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | ordered-lpage | 12 | 197 | 224 | 1.14× [1.13, 1.15] | +12.1% | [+11.5%, +12.8%] | 1.1 pts | 0.9 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | ordered-lpage-zc | 12 | 196 | 181 | 0.92× [0.91, 0.92] | -9.2% | [-10.1%, -8.4%] | 1.3 pts | 1.2 | yes | yes |
| unique-str | street | 4096 | churn | ordered | btree-map | 12 | 128 | 150 | 1.17× [1.15, 1.19] | +14.5% | [+13.4%, +15.7%] | 1.3 pts | 1.0 | yes | yes |
| unique-str | street | 4096 | churn | ordered | ordered-lpage | 12 | 128 | 194 | 1.52× [1.50, 1.53] | +34.0% | [+33.3%, +34.8%] | 1.2 pts | 1.2 | yes | yes |
| unique-str | street | 4096 | churn | ordered | ordered-lpage-zc | 12 | 134 | 273 | 2.06× [2.03, 2.10] | +51.6% | [+50.8%, +52.3%] | 1.2 pts | 0.9 | yes | yes |
| unique-str | street | 4096 | build | ordered | btree-map | 12 | 1.81 ms | 2.30 ms | 1.29× [1.24, 1.34] | +22.4% | [+19.5%, +25.3%] | 3.6 pts | 3.6 | no | yes |
| unique-str | street | 4096 | build | ordered | ordered-lpage | 12 | 1.83 ms | 3.47 ms | 1.90× [1.87, 1.93] | +47.4% | [+46.7%, +48.2%] | 0.9 pts | 1.8 | yes | yes |
| unique-str | street | 4096 | build | ordered | ordered-lpage-zc | 12 | 1.80 ms | 4.51 ms | 2.48× [2.46, 2.51] | +59.7% | [+59.3%, +60.1%] | 0.4 pts | 1.0 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | btree-map | 12 | 69.9 | 130 | 1.88× [1.82, 1.93] | +46.7% | [+45.1%, +48.3%] | 1.9 pts | 2.4 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | ordered-lpage | 12 | 68.7 | 108 | 1.59× [1.54, 1.64] | +37.1% | [+35.2%, +39.0%] | 2.0 pts | 3.1 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | ordered-lpage-zc | 12 | 68.1 | 91.0 | 1.34× [1.31, 1.38] | +25.5% | [+23.7%, +27.4%] | 2.1 pts | 2.7 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | btree-map | 12 | 2186 | 1425 | 0.65× [0.64, 0.66] | -53.6% | [-55.9%, -51.4%] | 2.6 pts | 2.0 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | ordered-lpage | 12 | 2180 | 1381 | 0.63× [0.62, 0.64] | -58.2% | [-60.0%, -56.4%] | 2.1 pts | 1.7 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | ordered-lpage-zc | 12 | 2175 | 919 | 0.42× [0.42, 0.43] | -137.1% | [-139.3%, -134.9%] | 2.4 pts | 1.0 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | btree-map | 12 | 645 | 519 | 0.79× [0.78, 0.80] | -26.5% | [-28.1%, -25.0%] | 2.9 pts | 1.2 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | ordered-lpage | 12 | 646 | 518 | 0.80× [0.78, 0.81] | -25.3% | [-27.7%, -22.9%] | 2.5 pts | 1.3 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | ordered-lpage-zc | 12 | 630 | 348 | 0.55× [0.54, 0.55] | -82.4% | [-84.5%, -80.2%] | 2.6 pts | 0.8 | yes | yes |
| unique-str | street | 16384 | churn | ordered | btree-map | 12 | 197 | 218 | 1.11× [1.07, 1.14] | +9.5% | [+6.9%, +12.2%] | 4.0 pts | 2.2 | no | yes |
| unique-str | street | 16384 | churn | ordered | ordered-lpage | 12 | 192 | 235 | 1.23× [1.21, 1.26] | +18.9% | [+17.0%, +20.8%] | 2.6 pts | 2.4 | yes | yes |
| unique-str | street | 16384 | churn | ordered | ordered-lpage-zc | 12 | 206 | 328 | 1.62× [1.60, 1.65] | +38.4% | [+37.4%, +39.4%] | 1.4 pts | 0.9 | yes | yes |
| unique-str | street | 16384 | build | ordered | btree-map | 12 | 9.82 ms | 12.68 ms | 1.29× [1.28, 1.30] | +22.3% | [+21.6%, +23.0%] | 1.6 pts | 1.0 | yes | yes |
| unique-str | street | 16384 | build | ordered | ordered-lpage | 12 | 9.38 ms | 15.74 ms | 1.67× [1.64, 1.69] | +40.0% | [+39.2%, +40.8%] | 0.8 pts | 1.3 | yes | yes |
| unique-str | street | 16384 | build | ordered | ordered-lpage-zc | 12 | 9.52 ms | 21.53 ms | 2.23× [2.20, 2.26] | +55.1% | [+54.5%, +55.7%] | 0.9 pts | 1.0 | yes | yes |
| unique-str | street | 212449 | valuesFor | ordered | btree-map | 10 | 341 | 359 | 1.07× [1.05, 1.09] | +6.5% | [+4.7%, +8.3%] | 2.7 pts | 1.5 | yes | yes |
| unique-str | street | 212449 | valuesFor | ordered | ordered-lpage | 10 | 288 | 222 | 0.77× [0.76, 0.79] | -29.4% | [-32.4%, -26.4%] | 3.0 pts | 1.2 | no | yes |
| unique-str | street | 212449 | valuesFor | ordered | ordered-lpage-zc | 10 | 279 | 178 | 0.66× [0.64, 0.67] | -52.3% | [-55.3%, -49.3%] | 3.4 pts | 1.0 | yes | yes |
| unique-str | street | 212449 | valuesBetween | ordered | btree-map | 10 | 6333 | 3868 | 0.62× [0.61, 0.64] | -60.2% | [-64.7%, -55.8%] | 5.8 pts | 2.0 | yes | yes |
| unique-str | street | 212449 | valuesBetween | ordered | ordered-lpage | 10 | 5564 | 2248 | 0.40× [0.39, 0.42] | -148.8% | [-156.6%, -140.9%] | 14.8 pts | 4.1 | yes | yes |
| unique-str | street | 212449 | valuesBetween | ordered | ordered-lpage-zc | 10 | 5009 | 1463 | 0.28× [0.27, 0.29] | -259.4% | [-273.2%, -245.7%] | 13.2 pts | 1.8 | yes | yes |
| unique-str | street | 212449 | prefix | ordered | btree-map | 10 | 18.7 µs | 12.1 µs | 0.65× [0.64, 0.65] | -54.5% | [-56.2%, -52.8%] | 4.5 pts | 0.8 | yes | yes |
| unique-str | street | 212449 | prefix | ordered | ordered-lpage | 10 | 17.7 µs | 6665 | 0.39× [0.38, 0.40] | -155.4% | [-163.8%, -146.9%] | 11.2 pts | 1.4 | yes | yes |
| unique-str | street | 212449 | prefix | ordered | ordered-lpage-zc | 10 | 15.0 µs | 3858 | 0.25× [0.25, 0.26] | -293.0% | [-306.1%, -280.0%] | 17.4 pts | 1.4 | yes | yes |
| unique-str | street | 212449 | churn | ordered | btree-map | 10 | 569 | 689 | 1.20× [1.18, 1.23] | +16.7% | [+15.1%, +18.4%] | 2.0 pts | 1.3 | yes | yes |
| unique-str | street | 212449 | churn | ordered | ordered-lpage | 10 | 567 | 563 | 1.00× [0.99, 1.02] | +0.4% | [-0.9%, +1.6%] | 1.5 pts | 1.4 | yes | no |
| unique-str | street | 212449 | churn | ordered | ordered-lpage-zc | 10 | 548 | 622 | 1.15× [1.13, 1.17] | +12.9% | [+11.6%, +14.3%] | 1.2 pts | 1.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 valuesFor: ordered vs ordered-lpage: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 build: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesBetween: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.37% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=16384 prefix: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.21% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=16384 churn: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 churn: ordered vs ordered-lpage: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesFor: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesBetween: ordered vs ordered-lpage: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesBetween: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -1.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=212449 churn: ordered vs ordered-lpage: the pooled difference of 0.36% does not clear the 0.99% noise floor, the bound on what the harness reports between identical code in every process
- unique-str street n=212449 churn: ordered vs ordered-lpage: the pooled interval [-0.92%, 1.63%] includes zero
- unique-str street n=212449 churn: ordered vs ordered-lpage: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
