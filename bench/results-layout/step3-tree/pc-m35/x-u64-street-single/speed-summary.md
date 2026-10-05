| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 48.6 | 46.0 | 0.95× [0.95, 0.95] | -5.3% | [-5.7%, -4.9%] | 1.1 pts | 1.3 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 48.9 | 90.7 | 1.86× [1.83, 1.88] | +46.1% | [+45.5%, +46.8%] | 0.7 pts | 2.2 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 2102 | 1805 | 0.86× [0.85, 0.87] | -16.3% | [-18.2%, -14.4%] | 1.8 pts | 1.0 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 2058 | 515 | 0.25× [0.25, 0.25] | -300.6% | [-306.0%, -295.2%] | 6.0 pts | 1.8 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 8 | 231 | 198 | 0.86× [0.85, 0.87] | -16.7% | [-18.2%, -15.3%] | 1.7 pts | 1.2 | yes | yes |
| single-value | street | 4096 | prefix | ordered | btree-map | 8 | 229 | 132 | 0.57× [0.57, 0.58] | -74.2% | [-75.7%, -72.8%] | 1.6 pts | 0.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 126 | 113 | 0.90× [0.89, 0.91] | -11.0% | [-11.8%, -10.2%] | 0.9 pts | 0.7 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 127 | 135 | 1.07× [1.05, 1.08] | +6.1% | [+4.9%, +7.4%] | 1.6 pts | 1.5 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 1.88 ms | 1.62 ms | 0.85× [0.85, 0.86] | -17.0% | [-17.8%, -16.2%] | 0.9 pts | 1.2 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 1.91 ms | 2.10 ms | 1.12× [1.09, 1.15] | +10.9% | [+8.6%, +13.1%] | 2.2 pts | 2.7 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 67.5 | 64.8 | 0.96× [0.92, 0.99] | -4.5% | [-8.2%, -0.8%] | 3.5 pts | 3.6 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 68.1 | 124 | 1.82× [1.77, 1.87] | +45.1% | [+43.6%, +46.6%] | 1.6 pts | 3.3 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 2479 | 2210 | 0.89× [0.88, 0.91] | -11.9% | [-13.3%, -10.5%] | 1.7 pts | 1.8 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 2475 | 610 | 0.24× [0.24, 0.25] | -310.2% | [-319.1%, -301.2%] | 9.1 pts | 3.5 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 8 | 754 | 647 | 0.86× [0.85, 0.86] | -16.9% | [-18.1%, -15.8%] | 1.1 pts | 0.5 | yes | yes |
| single-value | street | 16384 | prefix | ordered | btree-map | 8 | 743 | 285 | 0.38× [0.37, 0.39] | -164.5% | [-169.3%, -159.7%] | 6.0 pts | 1.8 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 166 | 152 | 0.93× [0.92, 0.93] | -8.0% | [-8.7%, -7.3%] | 0.7 pts | 0.5 | yes | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 161 | 191 | 1.18× [1.17, 1.20] | +15.4% | [+14.4%, +16.4%] | 1.3 pts | 1.5 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 9.00 ms | 8.00 ms | 0.89× [0.89, 0.90] | -12.2% | [-12.7%, -11.7%] | 0.5 pts | 0.8 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 9.13 ms | 11.00 ms | 1.20× [1.19, 1.21] | +16.8% | [+16.2%, +17.4%] | 0.7 pts | 0.6 | yes | yes |
| single-value | street | 212449 | valuesFor | ordered | baseline | 8 | 172 | 166 | 0.97× [0.93, 1.01] | -3.1% | [-7.2%, +1.1%] | 4.2 pts | 3.0 | no | no |
| single-value | street | 212449 | valuesFor | ordered | btree-map | 8 | 193 | 234 | 1.23× [1.18, 1.28] | +18.5% | [+14.9%, +22.0%] | 4.7 pts | 5.0 | no | yes |
| single-value | street | 212449 | valuesBetween | ordered | baseline | 8 | 4006 | 3637 | 0.93× [0.90, 0.97] | -7.4% | [-11.1%, -3.6%] | 4.1 pts | 2.8 | no | yes |
| single-value | street | 212449 | valuesBetween | ordered | btree-map | 8 | 3691 | 1261 | 0.34× [0.32, 0.36] | -193.8% | [-213.7%, -174.0%] | 18.5 pts | 1.3 | no | yes |
| single-value | street | 212449 | prefix | ordered | baseline | 8 | 13.7 µs | 12.5 µs | 0.90× [0.87, 0.93] | -10.6% | [-14.3%, -7.0%] | 4.2 pts | 3.6 | no | yes |
| single-value | street | 212449 | prefix | ordered | btree-map | 8 | 13.3 µs | 3435 | 0.27× [0.25, 0.28] | -276.9% | [-293.1%, -260.8%] | 16.2 pts | 1.3 | yes | yes |
| single-value | street | 212449 | churn | ordered | baseline | 8 | 383 | 371 | 0.96× [0.95, 0.97] | -4.2% | [-5.7%, -2.7%] | 1.6 pts | 1.2 | yes | yes |
| single-value | street | 212449 | churn | ordered | btree-map | 8 | 483 | 466 | 0.98× [0.96, 1.00] | -2.0% | [-4.4%, +0.3%] | 2.5 pts | 1.7 | no | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=212449 valuesFor: ordered vs baseline: the pooled interval [-7.18%, 1.08%] includes zero
- single-value street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs btree-map: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.96% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=212449 prefix: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 churn: ordered vs btree-map: the pooled interval [-4.38%, 0.32%] includes zero
