| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 8 | 66.0 | 47.4 | 0.71× [0.70, 0.73] | -40.1% | [-42.9%, -37.3%] | 2.7 pts | 2.7 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 8 | 66.9 | 93.4 | 1.37× [1.34, 1.40] | +26.8% | [+25.1%, +28.5%] | 2.2 pts | 3.3 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 3635 | 1860 | 0.51× [0.50, 0.52] | -95.5% | [-99.1%, -91.8%] | 3.5 pts | 1.3 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 8 | 3611 | 1137 | 0.31× [0.30, 0.31] | -224.8% | [-228.8%, -220.7%] | 3.9 pts | 0.7 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 8 | 366 | 207 | 0.56× [0.56, 0.57] | -77.3% | [-79.2%, -75.5%] | 1.9 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 8 | 366 | 189 | 0.51× [0.50, 0.51] | -96.5% | [-98.5%, -94.6%] | 2.5 pts | 1.0 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 8 | 125 | 120 | 0.96× [0.95, 0.97] | -4.5% | [-5.6%, -3.5%] | 1.2 pts | 1.1 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 8 | 125 | 139 | 1.10× [1.08, 1.12] | +9.1% | [+7.8%, +10.5%] | 1.3 pts | 1.4 | yes | yes |
| single-value-str | street | 4096 | build | ordered | baseline | 8 | 1.93 ms | 1.74 ms | 0.90× [0.88, 0.92] | -11.2% | [-13.7%, -8.7%] | 2.4 pts | 1.5 | no | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 8 | 1.93 ms | 2.20 ms | 1.15× [1.12, 1.18] | +13.2% | [+10.9%, +15.6%] | 2.2 pts | 1.7 | no | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 6 | 86.2 | 65.9 | 0.76× [0.75, 0.78] | -30.8% | [-33.2%, -28.5%] | 2.2 pts | 2.2 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 6 | 87.3 | 129 | 1.46× [1.42, 1.51] | +31.5% | [+29.4%, +33.6%] | 2.0 pts | 3.8 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 4210 | 2246 | 0.53× [0.53, 0.54] | -88.3% | [-89.9%, -86.6%] | 1.6 pts | 0.9 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 6 | 4203 | 1393 | 0.33× [0.33, 0.33] | -202.1% | [-205.5%, -198.6%] | 3.3 pts | 1.7 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 6 | 1327 | 666 | 0.50× [0.49, 0.51] | -100.0% | [-102.1%, -97.9%] | 2.0 pts | 0.6 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 6 | 1320 | 499 | 0.37× [0.37, 0.38] | -167.8% | [-172.5%, -163.1%] | 4.5 pts | 1.2 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 6 | 164 | 158 | 0.97× [0.95, 0.98] | -3.6% | [-4.9%, -2.3%] | 1.3 pts | 1.6 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 6 | 161 | 195 | 1.21× [1.18, 1.23] | +17.2% | [+15.4%, +18.9%] | 1.7 pts | 2.3 | yes | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 6 | 9.31 ms | 8.45 ms | 0.91× [0.91, 0.92] | -9.5% | [-9.8%, -9.2%] | 0.2 pts | 0.2 | yes | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 6 | 9.39 ms | 11.53 ms | 1.23× [1.21, 1.25] | +18.9% | [+17.6%, +20.1%] | 1.2 pts | 0.6 | yes | yes |
| single-value-str | street | 212449 | valuesFor | ordered | baseline | 8 | 203 | 172 | 0.85× [0.82, 0.89] | -17.6% | [-22.6%, -12.6%] | 4.9 pts | 3.3 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | btree-map | 8 | 246 | 248 | 1.02× [1.00, 1.05] | +2.2% | [-0.3%, +4.7%] | 2.5 pts | 2.1 | no | no |
| single-value-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 6404 | 4111 | 0.65× [0.63, 0.66] | -54.7% | [-58.7%, -50.7%] | 4.1 pts | 2.1 | yes | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | btree-map | 8 | 6356 | 2348 | 0.38× [0.37, 0.39] | -164.8% | [-170.1%, -159.5%] | 7.7 pts | 2.2 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | baseline | 8 | 24.0 µs | 13.8 µs | 0.57× [0.55, 0.59] | -75.9% | [-81.6%, -70.1%] | 5.7 pts | 2.0 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | btree-map | 8 | 22.9 µs | 7900 | 0.34× [0.33, 0.36] | -190.9% | [-201.5%, -180.3%] | 10.9 pts | 1.8 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | baseline | 8 | 411 | 396 | 0.96× [0.95, 0.98] | -4.1% | [-5.7%, -2.4%] | 1.6 pts | 1.2 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | btree-map | 8 | 524 | 509 | 1.00× [0.98, 1.02] | -0.3% | [-2.5%, +1.8%] | 2.1 pts | 1.6 | no | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 churn: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the pooled interval [-0.33%, 4.67%] includes zero
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs btree-map: the A/A validations found a systematic difference of +2.84% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 churn: ordered vs btree-map: the pooled difference of -0.31% does not clear the 0.39% noise floor, the bound on what the harness reports between identical code in every process
- single-value-str street n=212449 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 churn: ordered vs btree-map: the pooled interval [-2.45%, 1.82%] includes zero
- single-value-str street n=212449 churn: ordered vs btree-map: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
