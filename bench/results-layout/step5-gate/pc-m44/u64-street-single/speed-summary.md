| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 61.1 | 48.7 | 0.79× [0.79, 0.80] | -25.9% | [-26.9%, -25.0%] | 1.3 pts | 1.6 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 61.6 | 90.0 | 1.46× [1.45, 1.47] | +31.5% | [+30.9%, +32.1%] | 0.5 pts | 1.1 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 1109 | 2078 | 1.86× [1.84, 1.89] | +46.3% | [+45.5%, +47.0%] | 0.7 pts | 1.1 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1096 | 492 | 0.44× [0.44, 0.45] | -125.4% | [-127.3%, -123.4%] | 2.1 pts | 0.9 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 8 | 217 | 229 | 1.06× [1.05, 1.06] | +5.4% | [+4.8%, +6.0%] | 0.6 pts | 0.8 | yes | yes |
| single-value | street | 4096 | prefix | ordered | btree-map | 8 | 216 | 131 | 0.60× [0.60, 0.61] | -65.6% | [-67.0%, -64.2%] | 1.5 pts | 1.3 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 187 | 125 | 0.67× [0.67, 0.67] | -49.6% | [-50.4%, -48.8%] | 1.1 pts | 0.6 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 186 | 136 | 0.73× [0.72, 0.73] | -37.2% | [-38.0%, -36.4%] | 1.0 pts | 0.8 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.60 ms | 1.89 ms | 0.73× [0.72, 0.73] | -37.7% | [-39.1%, -36.3%] | 1.6 pts | 1.2 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.59 ms | 2.17 ms | 0.84× [0.81, 0.86] | -19.6% | [-22.9%, -16.4%] | 4.1 pts | 3.4 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 71.9 | 68.5 | 0.95× [0.94, 0.95] | -5.8% | [-6.5%, -5.0%] | 1.0 pts | 1.3 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 72.8 | 124 | 1.70× [1.66, 1.73] | +41.1% | [+39.9%, +42.3%] | 1.5 pts | 2.9 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1264 | 2505 | 1.97× [1.96, 1.98] | +49.2% | [+48.9%, +49.6%] | 0.4 pts | 0.8 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1265 | 594 | 0.47× [0.46, 0.48] | -113.8% | [-117.0%, -110.5%] | 4.6 pts | 3.0 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 8 | 453 | 755 | 1.67× [1.65, 1.68] | +40.0% | [+39.5%, +40.5%] | 0.5 pts | 0.8 | yes | yes |
| single-value | street | 16384 | prefix | ordered | btree-map | 8 | 449 | 274 | 0.61× [0.60, 0.61] | -65.1% | [-67.3%, -62.9%] | 3.4 pts | 2.1 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 198 | 163 | 0.82× [0.80, 0.83] | -22.7% | [-24.7%, -20.6%] | 1.9 pts | 1.3 | yes | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 199 | 191 | 0.96× [0.95, 0.97] | -4.2% | [-5.0%, -3.3%] | 1.4 pts | 1.4 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 11.42 ms | 8.92 ms | 0.78× [0.78, 0.79] | -27.9% | [-28.9%, -26.9%] | 1.0 pts | 1.2 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 11.64 ms | 10.97 ms | 0.94× [0.93, 0.95] | -6.4% | [-7.3%, -5.5%] | 1.3 pts | 0.7 | yes | yes |
| single-value | street | 212449 | valuesFor | ordered | baseline | 8 | 118 | 150 | 1.26× [1.22, 1.31] | +20.7% | [+17.9%, +23.4%] | 2.6 pts | 2.4 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | btree-map | 8 | 126 | 226 | 1.81× [1.76, 1.87] | +44.9% | [+43.3%, +46.5%] | 1.7 pts | 2.4 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | baseline | 8 | 1476 | 3142 | 2.12× [2.07, 2.18] | +52.9% | [+51.7%, +54.0%] | 1.2 pts | 2.3 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | btree-map | 8 | 1453 | 923 | 0.63× [0.59, 0.67] | -58.8% | [-68.8%, -48.9%] | 9.5 pts | 4.9 | no | yes |
| single-value | street | 212449 | prefix | ordered | baseline | 8 | 4883 | 12.8 µs | 2.61× [2.56, 2.66] | +61.7% | [+60.9%, +62.4%] | 0.7 pts | 0.8 | yes | yes |
| single-value | street | 212449 | prefix | ordered | btree-map | 8 | 4741 | 3062 | 0.64× [0.62, 0.66] | -56.5% | [-61.1%, -52.0%] | 4.9 pts | 1.8 | yes | yes |
| single-value | street | 212449 | churn | ordered | baseline | 8 | 339 | 405 | 1.20× [1.17, 1.23] | +16.6% | [+14.6%, +18.6%] | 2.4 pts | 1.3 | yes | yes |
| single-value | street | 212449 | churn | ordered | btree-map | 8 | 410 | 459 | 1.13× [1.11, 1.14] | +11.2% | [+10.0%, +12.5%] | 1.4 pts | 1.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value street n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.81% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 prefix: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
