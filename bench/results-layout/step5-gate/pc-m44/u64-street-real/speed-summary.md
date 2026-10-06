| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | street | 4096 | valuesFor | ordered | baseline | 6 | 68.5 | 54.3 | 0.79× [0.79, 0.80] | -26.3% | [-27.2%, -25.5%] | 0.8 pts | 0.7 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 6 | 69.9 | 134 | 1.92× [1.90, 1.94] | +48.0% | [+47.4%, +48.5%] | 0.5 pts | 1.1 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 6 | 1889 | 2570 | 1.39× [1.36, 1.41] | +27.8% | [+26.5%, +29.1%] | 1.2 pts | 0.8 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 1886 | 4566 | 2.42× [2.36, 2.48] | +58.7% | [+57.7%, +59.6%] | 0.9 pts | 1.7 | yes | yes |
| natural | street | 4096 | prefix | ordered | baseline | 6 | 259 | 256 | 0.99× [0.99, 1.00] | -1.0% | [-1.5%, -0.4%] | 0.5 pts | 0.5 | yes | yes |
| natural | street | 4096 | prefix | ordered | btree-sets | 6 | 261 | 525 | 2.02× [2.01, 2.03] | +50.6% | [+50.4%, +50.8%] | 0.2 pts | 0.5 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 6 | 167 | 104 | 0.62× [0.61, 0.63] | -61.0% | [-62.8%, -59.2%] | 1.7 pts | 0.8 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 6 | 168 | 189 | 1.12× [1.10, 1.14] | +10.6% | [+9.3%, +11.9%] | 1.3 pts | 1.1 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 6 | 6.49 ms | 4.10 ms | 0.64× [0.63, 0.65] | -56.8% | [-58.9%, -54.7%] | 2.0 pts | 1.2 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 6 | 6.56 ms | 7.36 ms | 1.13× [1.12, 1.15] | +11.9% | [+10.7%, +13.0%] | 1.1 pts | 1.2 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 6 | 84.2 | 75.0 | 0.88× [0.87, 0.90] | -13.4% | [-15.2%, -11.6%] | 1.7 pts | 1.8 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 6 | 84.8 | 181 | 2.12× [2.09, 2.15] | +52.8% | [+52.2%, +53.4%] | 0.6 pts | 1.6 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 6 | 2249 | 3136 | 1.40× [1.39, 1.40] | +28.4% | [+28.1%, +28.7%] | 0.3 pts | 0.6 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 2269 | 5382 | 2.36× [2.35, 2.38] | +57.7% | [+57.5%, +57.9%] | 0.2 pts | 0.8 | yes | yes |
| natural | street | 16384 | prefix | ordered | baseline | 6 | 714 | 927 | 1.30× [1.28, 1.32] | +23.0% | [+22.0%, +24.0%] | 0.9 pts | 0.9 | yes | yes |
| natural | street | 16384 | prefix | ordered | btree-sets | 6 | 727 | 2043 | 2.82× [2.78, 2.87] | +64.6% | [+64.0%, +65.2%] | 0.6 pts | 1.0 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 6 | 201 | 140 | 0.70× [0.69, 0.71] | -42.9% | [-44.4%, -41.3%] | 1.5 pts | 1.2 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 6 | 204 | 264 | 1.29× [1.26, 1.32] | +22.4% | [+20.4%, +24.5%] | 1.9 pts | 2.2 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 6 | 30.90 ms | 21.46 ms | 0.70× [0.69, 0.70] | -43.6% | [-44.8%, -42.4%] | 1.1 pts | 1.1 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 6 | 31.39 ms | 40.24 ms | 1.29× [1.27, 1.31] | +22.6% | [+21.4%, +23.7%] | 1.1 pts | 1.1 | yes | yes |
| natural | street | 212449 | valuesFor | ordered | baseline | 8 | 150 | 176 | 1.18× [1.14, 1.23] | +15.5% | [+12.5%, +18.5%] | 4.0 pts | 3.4 | no | yes |
| natural | street | 212449 | valuesFor | ordered | btree-sets | 8 | 189 | 477 | 2.54× [2.46, 2.63] | +60.6% | [+59.3%, +62.0%] | 1.4 pts | 2.4 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | baseline | 8 | 3106 | 5074 | 1.64× [1.57, 1.72] | +39.0% | [+36.4%, +41.7%] | 2.6 pts | 3.0 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 3750 | 15.9 µs | 4.33× [4.26, 4.40] | +76.9% | [+76.6%, +77.3%] | 0.5 pts | 1.3 | yes | yes |
| natural | street | 212449 | prefix | ordered | baseline | 8 | 10.1 µs | 17.4 µs | 1.68× [1.65, 1.71] | +40.3% | [+39.2%, +41.4%] | 1.4 pts | 2.5 | yes | yes |
| natural | street | 212449 | prefix | ordered | btree-sets | 8 | 11.8 µs | 50.6 µs | 4.31× [4.17, 4.45] | +76.8% | [+76.0%, +77.5%] | 0.9 pts | 1.6 | yes | yes |
| natural | street | 212449 | churn | ordered | baseline | 8 | 434 | 453 | 1.00× [0.97, 1.03] | +0.2% | [-2.6%, +3.0%] | 3.6 pts | 0.8 | no | no |
| natural | street | 212449 | churn | ordered | btree-sets | 8 | 518 | 725 | 1.40× [1.37, 1.43] | +28.5% | [+27.1%, +29.9%] | 1.8 pts | 1.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=16384 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 churn: ordered vs baseline: the pooled difference of 0.17% does not clear the 2.31% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=212449 churn: ordered vs baseline: the pooled interval [-2.64%, 2.98%] includes zero
