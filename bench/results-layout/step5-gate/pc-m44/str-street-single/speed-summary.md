| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 8 | 76.2 | 67.2 | 0.88× [0.86, 0.89] | -14.3% | [-16.0%, -12.6%] | 1.6 pts | 2.3 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 8 | 77.0 | 92.9 | 1.18× [1.15, 1.21] | +15.2% | [+13.3%, +17.1%] | 2.4 pts | 3.7 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 2647 | 3722 | 1.40× [1.39, 1.41] | +28.6% | [+28.1%, +29.2%] | 0.6 pts | 0.6 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 8 | 2581 | 1007 | 0.38× [0.37, 0.38] | -164.7% | [-167.0%, -162.4%] | 2.3 pts | 0.5 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 8 | 379 | 375 | 0.99× [0.98, 1.01] | -0.7% | [-2.1%, +0.7%] | 1.3 pts | 1.1 | yes | no |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 8 | 375 | 178 | 0.47× [0.46, 0.47] | -114.9% | [-117.0%, -112.8%] | 2.1 pts | 1.0 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 8 | 197 | 130 | 0.65× [0.64, 0.66] | -54.2% | [-56.1%, -52.4%] | 2.2 pts | 1.1 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 8 | 196 | 141 | 0.71× [0.71, 0.72] | -40.0% | [-40.9%, -39.2%] | 0.9 pts | 0.7 | yes | yes |
| single-value-str | street | 4096 | build | ordered | baseline | 8 | 3.03 ms | 1.89 ms | 0.61× [0.61, 0.62] | -62.8% | [-64.8%, -60.8%] | 2.7 pts | 2.6 | yes | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 8 | 3.05 ms | 2.17 ms | 0.72× [0.71, 0.73] | -39.3% | [-41.2%, -37.3%] | 1.9 pts | 1.3 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 8 | 89.2 | 87.7 | 0.98× [0.96, 0.99] | -2.4% | [-3.6%, -1.2%] | 1.4 pts | 1.7 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 8 | 91.1 | 131 | 1.42× [1.39, 1.44] | +29.4% | [+28.1%, +30.6%] | 1.2 pts | 1.8 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 2967 | 4245 | 1.43× [1.42, 1.43] | +29.8% | [+29.4%, +30.2%] | 0.5 pts | 1.0 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 8 | 2963 | 1289 | 0.43× [0.43, 0.44] | -130.7% | [-132.9%, -128.5%] | 3.2 pts | 1.8 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 8 | 975 | 1346 | 1.38× [1.36, 1.40] | +27.4% | [+26.3%, +28.4%] | 1.0 pts | 0.8 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 8 | 960 | 463 | 0.48× [0.47, 0.49] | -109.4% | [-113.0%, -105.8%] | 4.0 pts | 1.2 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 8 | 222 | 167 | 0.76× [0.74, 0.77] | -32.4% | [-34.7%, -30.2%] | 2.1 pts | 0.9 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 8 | 221 | 197 | 0.89× [0.88, 0.90] | -12.4% | [-13.4%, -11.4%] | 1.1 pts | 0.8 | yes | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 8 | 12.58 ms | 8.95 ms | 0.71× [0.70, 0.72] | -40.7% | [-43.3%, -38.1%] | 2.5 pts | 2.7 | yes | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 8 | 12.97 ms | 11.32 ms | 0.88× [0.86, 0.90] | -13.3% | [-15.6%, -10.9%] | 2.3 pts | 1.0 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | baseline | 8 | 144 | 183 | 1.32× [1.22, 1.43] | +24.0% | [+17.8%, +30.1%] | 6.1 pts | 6.3 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | btree-map | 8 | 163 | 236 | 1.48× [1.45, 1.51] | +32.4% | [+31.2%, +33.6%] | 2.2 pts | 2.3 | yes | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 3402 | 5541 | 1.63× [1.60, 1.66] | +38.6% | [+37.4%, +39.9%] | 1.6 pts | 2.3 | yes | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | btree-map | 8 | 3447 | 1921 | 0.56× [0.52, 0.60] | -79.4% | [-93.1%, -65.7%] | 15.3 pts | 4.6 | no | yes |
| single-value-str | street | 212449 | prefix | ordered | baseline | 8 | 12.6 µs | 22.9 µs | 1.80× [1.76, 1.83] | +44.3% | [+43.1%, +45.5%] | 1.2 pts | 3.1 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | btree-map | 8 | 12.7 µs | 6921 | 0.53× [0.52, 0.54] | -89.4% | [-93.1%, -85.7%] | 9.1 pts | 2.8 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | baseline | 8 | 414 | 445 | 1.06× [1.04, 1.08] | +5.9% | [+4.3%, +7.5%] | 1.7 pts | 1.0 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | btree-map | 8 | 500 | 503 | 1.05× [1.03, 1.06] | +4.6% | [+3.2%, +5.9%] | 1.8 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 prefix: ordered vs baseline: the pooled difference of -0.71% does not clear the 1.22% noise floor, the bound on what the harness reports between identical code in every process
- single-value-str street n=4096 prefix: ordered vs baseline: the pooled interval [-2.11%, 0.68%] includes zero
- single-value-str street n=4096 build: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 build: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.66% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 6.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
