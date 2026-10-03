| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique | dirs | 4096 | valuesFor | ordered | btree-map | 12 | 91.0 | 110 | 1.21× [1.18, 1.24] | +17.1% | [+15.0%, +19.3%] | 2.1 pts | 2.8 | no | yes |
| unique | dirs | 4096 | valuesBetween | ordered | btree-map | 12 | 667 | 542 | 0.81× [0.79, 0.83] | -23.8% | [-26.8%, -20.8%] | 3.8 pts | 2.9 | no | yes |
| unique | dirs | 4096 | prefix | ordered | btree-map | 12 | 700 | 655 | 0.93× [0.91, 0.96] | -7.2% | [-10.5%, -3.9%] | 3.7 pts | 2.6 | no | yes |
| unique | dirs | 4096 | churn | ordered | btree-map | 12 | 201 | 181 | 0.90× [0.90, 0.91] | -10.8% | [-11.5%, -10.1%] | 1.2 pts | 0.8 | yes | yes |
| unique | dirs | 4096 | build | ordered | btree-map | 12 | 3.55 ms | 2.64 ms | 0.74× [0.74, 0.75] | -34.3% | [-36.0%, -32.6%] | 1.6 pts | 1.6 | yes | yes |
| unique | dirs | 16384 | valuesFor | ordered | btree-map | 12 | 116 | 160 | 1.40× [1.39, 1.41] | +28.7% | [+28.2%, +29.2%] | 0.7 pts | 1.1 | yes | yes |
| unique | dirs | 16384 | valuesBetween | ordered | btree-map | 12 | 829 | 716 | 0.86× [0.85, 0.88] | -15.7% | [-17.1%, -14.3%] | 1.6 pts | 1.5 | yes | yes |
| unique | dirs | 16384 | prefix | ordered | btree-map | 12 | 2162 | 2341 | 1.08× [1.06, 1.11] | +7.7% | [+5.5%, +9.9%] | 2.1 pts | 1.5 | no | yes |
| unique | dirs | 16384 | churn | ordered | btree-map | 12 | 246 | 246 | 1.01× [1.00, 1.02] | +0.7% | [-0.2%, +1.6%] | 1.1 pts | 1.3 | yes | no |
| unique | dirs | 16384 | build | ordered | btree-map | 12 | 16.25 ms | 13.45 ms | 0.82× [0.82, 0.83] | -21.6% | [-22.6%, -20.6%] | 1.5 pts | 0.9 | yes | yes |
| unique | dirs | 86215 | valuesFor | ordered | btree-map | 12 | 165 | 236 | 1.44× [1.42, 1.47] | +30.8% | [+29.8%, +31.7%] | 1.5 pts | 1.7 | yes | yes |
| unique | dirs | 86215 | valuesBetween | ordered | btree-map | 12 | 1013 | 998 | 0.99× [0.96, 1.02] | -1.0% | [-3.9%, +1.9%] | 4.1 pts | 3.4 | no | no |
| unique | dirs | 86215 | prefix | ordered | btree-map | 12 | 15.0 µs | 16.3 µs | 1.04× [1.03, 1.05] | +4.1% | [+3.2%, +5.1%] | 1.4 pts | 0.7 | yes | yes |
| unique | dirs | 86215 | churn | ordered | btree-map | 12 | 432 | 465 | 1.08× [1.08, 1.09] | +7.5% | [+7.1%, +7.9%] | 1.5 pts | 1.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique dirs n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique dirs n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique dirs n=4096 prefix: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique dirs n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique dirs n=16384 churn: ordered vs btree-map: the pooled interval [-0.16%, 1.65%] includes zero
- unique dirs n=86215 valuesBetween: ordered vs btree-map: the pooled interval [-3.88%, 1.91%] includes zero
- unique dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique dirs n=86215 valuesBetween: ordered vs btree-map: 2 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
