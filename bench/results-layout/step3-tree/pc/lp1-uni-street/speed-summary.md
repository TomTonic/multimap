| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | street | 4096 | valuesFor | ordered | btree-map | 10 | 46.5 | 91.0 | 1.96× [1.95, 1.98] | +49.1% | [+48.7%, +49.5%] | 0.7 pts | 2.1 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | ordered-lpage | 10 | 46.5 | 78.2 | 1.70× [1.68, 1.72] | +41.3% | [+40.6%, +41.9%] | 0.7 pts | 2.0 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | btree-map | 10 | 1732 | 1127 | 0.63× [0.62, 0.65] | -57.8% | [-60.8%, -54.7%] | 2.9 pts | 1.1 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | ordered-lpage | 10 | 1769 | 1351 | 0.77× [0.76, 0.77] | -30.5% | [-31.6%, -29.4%] | 1.4 pts | 0.8 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | btree-map | 10 | 197 | 187 | 0.95× [0.93, 0.96] | -5.7% | [-7.6%, -3.8%] | 1.8 pts | 1.1 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | ordered-lpage | 10 | 197 | 230 | 1.16× [1.15, 1.18] | +14.0% | [+13.0%, +15.0%] | 1.3 pts | 1.1 | yes | yes |
| unique-str | street | 4096 | churn | ordered | btree-map | 10 | 122 | 141 | 1.15× [1.13, 1.18] | +13.3% | [+11.5%, +15.1%] | 1.8 pts | 1.9 | yes | yes |
| unique-str | street | 4096 | churn | ordered | ordered-lpage | 10 | 123 | 261 | 2.13× [2.11, 2.15] | +53.0% | [+52.5%, +53.5%] | 0.6 pts | 0.9 | yes | yes |
| unique-str | street | 4096 | build | ordered | btree-map | 10 | 1.76 ms | 2.17 ms | 1.26× [1.24, 1.29] | +20.9% | [+19.1%, +22.6%] | 3.5 pts | 5.1 | yes | yes |
| unique-str | street | 4096 | build | ordered | ordered-lpage | 10 | 1.76 ms | 4.44 ms | 2.52× [2.50, 2.54] | +60.3% | [+60.1%, +60.6%] | 0.3 pts | 0.7 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | btree-map | 12 | 66.3 | 127 | 1.93× [1.86, 1.99] | +48.1% | [+46.3%, +49.8%] | 2.0 pts | 3.8 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | ordered-lpage | 12 | 65.1 | 91.8 | 1.43× [1.39, 1.47] | +29.8% | [+27.8%, +31.8%] | 2.5 pts | 3.9 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | btree-map | 12 | 2142 | 1390 | 0.64× [0.64, 0.65] | -55.2% | [-57.5%, -53.0%] | 2.6 pts | 2.1 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | ordered-lpage | 12 | 2148 | 1598 | 0.74× [0.74, 0.75] | -34.7% | [-35.8%, -33.7%] | 2.3 pts | 2.0 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | btree-map | 12 | 630 | 499 | 0.79× [0.78, 0.79] | -27.0% | [-28.1%, -25.8%] | 1.8 pts | 0.8 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | ordered-lpage | 12 | 627 | 563 | 0.90× [0.89, 0.91] | -11.3% | [-12.8%, -9.8%] | 1.9 pts | 1.2 | yes | yes |
| unique-str | street | 16384 | churn | ordered | btree-map | 12 | 161 | 199 | 1.23× [1.19, 1.27] | +18.8% | [+16.1%, +21.5%] | 2.6 pts | 3.1 | no | yes |
| unique-str | street | 16384 | churn | ordered | ordered-lpage | 12 | 163 | 300 | 1.83× [1.80, 1.87] | +45.5% | [+44.4%, +46.5%] | 1.1 pts | 1.6 | yes | yes |
| unique-str | street | 16384 | build | ordered | btree-map | 12 | 8.69 ms | 11.48 ms | 1.31× [1.30, 1.33] | +23.9% | [+23.1%, +24.6%] | 0.9 pts | 0.8 | yes | yes |
| unique-str | street | 16384 | build | ordered | ordered-lpage | 12 | 8.62 ms | 18.75 ms | 2.18× [2.16, 2.19] | +54.1% | [+53.8%, +54.3%] | 0.4 pts | 0.9 | yes | yes |
| unique-str | street | 212449 | valuesFor | ordered | btree-map | 12 | 218 | 241 | 1.11× [1.07, 1.15] | +9.6% | [+6.4%, +12.7%] | 3.6 pts | 2.3 | no | yes |
| unique-str | street | 212449 | valuesFor | ordered | ordered-lpage | 12 | 176 | 154 | 0.89× [0.86, 0.93] | -11.9% | [-16.3%, -7.5%] | 6.7 pts | 3.7 | no | yes |
| unique-str | street | 212449 | valuesBetween | ordered | btree-map | 12 | 3635 | 2163 | 0.60× [0.58, 0.63] | -66.5% | [-73.7%, -59.3%] | 8.0 pts | 2.7 | no | yes |
| unique-str | street | 212449 | valuesBetween | ordered | ordered-lpage | 12 | 3538 | 2344 | 0.67× [0.64, 0.70] | -49.3% | [-56.1%, -42.5%] | 10.5 pts | 3.4 | no | yes |
| unique-str | street | 212449 | prefix | ordered | btree-map | 12 | 12.2 µs | 7769 | 0.63× [0.61, 0.65] | -59.3% | [-65.1%, -53.5%] | 5.7 pts | 3.1 | yes | yes |
| unique-str | street | 212449 | prefix | ordered | ordered-lpage | 12 | 11.7 µs | 7375 | 0.63× [0.61, 0.65] | -59.6% | [-64.5%, -54.7%] | 4.9 pts | 1.2 | yes | yes |
| unique-str | street | 212449 | churn | ordered | btree-map | 12 | 500 | 518 | 1.06× [1.05, 1.07] | +5.6% | [+4.6%, +6.6%] | 2.1 pts | 1.1 | yes | yes |
| unique-str | street | 212449 | churn | ordered | ordered-lpage | 12 | 447 | 545 | 1.21× [1.21, 1.22] | +17.5% | [+17.2%, +17.8%] | 1.9 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 build: ordered vs btree-map: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 churn: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesFor: ordered vs ordered-lpage: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesBetween: ordered vs ordered-lpage: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 prefix: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
