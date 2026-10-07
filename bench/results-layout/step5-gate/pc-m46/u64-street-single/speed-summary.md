| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 68.9 | 48.8 | 0.71× [0.70, 0.71] | -41.6% | [-42.6%, -40.7%] | 1.1 pts | 1.4 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 69.4 | 91.4 | 1.31× [1.30, 1.32] | +23.6% | [+22.9%, +24.4%] | 0.8 pts | 1.5 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 1230 | 2101 | 1.71× [1.68, 1.73] | +41.4% | [+40.6%, +42.2%] | 0.8 pts | 1.3 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1212 | 513 | 0.42× [0.42, 0.42] | -137.2% | [-138.3%, -136.1%] | 1.1 pts | 0.6 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 8 | 228 | 232 | 1.01× [1.01, 1.02] | +1.4% | [+0.8%, +2.1%] | 0.6 pts | 0.8 | yes | yes |
| single-value | street | 4096 | prefix | ordered | btree-map | 8 | 227 | 134 | 0.59× [0.58, 0.59] | -70.1% | [-71.9%, -68.3%] | 1.8 pts | 1.3 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 189 | 134 | 0.70× [0.70, 0.71] | -42.0% | [-43.7%, -40.2%] | 2.3 pts | 1.0 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 187 | 142 | 0.75× [0.74, 0.76] | -33.1% | [-34.4%, -31.8%] | 2.5 pts | 1.9 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.58 ms | 1.96 ms | 0.76× [0.74, 0.78] | -31.9% | [-35.5%, -28.4%] | 3.5 pts | 2.0 | no | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.66 ms | 2.18 ms | 0.83× [0.80, 0.86] | -20.3% | [-24.7%, -15.9%] | 4.4 pts | 2.4 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 79.9 | 67.6 | 0.85× [0.84, 0.86] | -17.8% | [-19.7%, -15.9%] | 2.0 pts | 2.1 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 81.2 | 125 | 1.53× [1.49, 1.57] | +34.5% | [+32.7%, +36.2%] | 1.8 pts | 3.1 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1386 | 2545 | 1.83× [1.82, 1.84] | +45.3% | [+45.0%, +45.7%] | 0.4 pts | 0.9 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1376 | 609 | 0.44× [0.43, 0.45] | -127.9% | [-132.7%, -123.2%] | 5.3 pts | 3.6 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 8 | 495 | 765 | 1.54× [1.52, 1.55] | +35.0% | [+34.4%, +35.7%] | 0.6 pts | 0.9 | yes | yes |
| single-value | street | 16384 | prefix | ordered | btree-map | 8 | 494 | 282 | 0.57× [0.56, 0.58] | -75.8% | [-78.8%, -72.7%] | 3.0 pts | 2.0 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 200 | 174 | 0.86× [0.85, 0.88] | -15.8% | [-17.6%, -14.0%] | 2.4 pts | 1.9 | yes | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 202 | 198 | 0.98× [0.97, 1.00] | -1.7% | [-3.1%, -0.2%] | 1.5 pts | 1.4 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 11.59 ms | 9.50 ms | 0.82× [0.81, 0.83] | -22.0% | [-23.7%, -20.3%] | 1.9 pts | 1.4 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 11.89 ms | 11.51 ms | 0.96× [0.93, 1.00] | -3.7% | [-7.1%, -0.2%] | 3.5 pts | 1.4 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | baseline | 8 | 129 | 160 | 1.30× [1.12, 1.54] | +23.0% | [+11.0%, +35.0%] | 12.5 pts | 13.7 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | btree-map | 8 | 143 | 255 | 1.78× [1.64, 1.94] | +43.7% | [+39.1%, +48.4%] | 6.0 pts | 5.3 | no | yes |
| single-value | street | 212449 | valuesBetween | ordered | baseline | 8 | 1670 | 4021 | 2.39× [2.01, 2.96] | +58.2% | [+50.2%, +66.2%] | 8.1 pts | 17.8 | no | yes |
| single-value | street | 212449 | valuesBetween | ordered | btree-map | 8 | 1600 | 1132 | 0.70× [0.61, 0.83] | -42.8% | [-65.1%, -20.5%] | 23.2 pts | 10.4 | no | yes |
| single-value | street | 212449 | prefix | ordered | baseline | 8 | 5322 | 14.1 µs | 2.79× [2.40, 3.31] | +64.1% | [+58.4%, +69.8%] | 5.9 pts | 7.4 | yes | yes |
| single-value | street | 212449 | prefix | ordered | btree-map | 8 | 5207 | 3546 | 0.69× [0.59, 0.82] | -45.1% | [-68.9%, -21.3%] | 23.3 pts | 5.5 | no | yes |
| single-value | street | 212449 | churn | ordered | baseline | 8 | 356 | 427 | 1.19× [1.17, 1.22] | +16.2% | [+14.6%, +17.8%] | 1.7 pts | 1.3 | yes | yes |
| single-value | street | 212449 | churn | ordered | btree-map | 8 | 415 | 485 | 1.18× [1.14, 1.22] | +15.3% | [+12.3%, +18.2%] | 3.0 pts | 3.4 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.71% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs baseline: the processes scatter 13.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs btree-map: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs baseline: the processes scatter 17.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 10.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs btree-map: 1 processes resolved A as faster and 7 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=212449 prefix: ordered vs baseline: the processes scatter 7.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 prefix: ordered vs btree-map: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 prefix: ordered vs btree-map: 1 processes resolved A as faster and 7 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=212449 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=212449 churn: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
