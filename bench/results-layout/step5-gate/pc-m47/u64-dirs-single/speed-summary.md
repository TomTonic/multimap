| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 85.4 | 78.4 | 0.92× [0.91, 0.93] | -9.0% | [-10.2%, -7.7%] | 1.2 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 86.7 | 108 | 1.25× [1.23, 1.27] | +19.9% | [+18.8%, +21.1%] | 1.1 pts | 1.3 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1406 | 2374 | 1.70× [1.68, 1.73] | +41.3% | [+40.4%, +42.2%] | 0.9 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1372 | 549 | 0.40× [0.40, 0.41] | -149.1% | [-151.7%, -146.6%] | 2.6 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 8 | 1616 | 2723 | 1.68× [1.64, 1.71] | +40.3% | [+39.1%, +41.6%] | 1.2 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | btree-map | 8 | 1610 | 663 | 0.41× [0.41, 0.41] | -143.3% | [-144.5%, -142.0%] | 2.2 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 255 | 205 | 0.81× [0.79, 0.82] | -24.0% | [-25.8%, -22.2%] | 1.8 pts | 1.5 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 257 | 184 | 0.71× [0.70, 0.72] | -41.5% | [-43.6%, -39.3%] | 2.3 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 3.86 ms | 3.12 ms | 0.80× [0.79, 0.82] | -24.2% | [-26.6%, -21.9%] | 2.2 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 3.87 ms | 2.74 ms | 0.71× [0.70, 0.72] | -41.6% | [-43.8%, -39.4%] | 2.1 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 108 | 108 | 1.00× [0.99, 1.02] | +0.4% | [-1.0%, +1.8%] | 1.5 pts | 1.5 | yes | no |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 109 | 157 | 1.45× [1.43, 1.48] | +31.2% | [+30.0%, +32.3%] | 1.4 pts | 1.8 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1594 | 2913 | 1.83× [1.82, 1.85] | +45.4% | [+45.0%, +45.8%] | 0.4 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 1572 | 722 | 0.46× [0.46, 0.46] | -117.6% | [-119.3%, -115.9%] | 3.4 pts | 2.0 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 8 | 5724 | 12.2 µs | 2.09× [2.02, 2.16] | +52.1% | [+50.5%, +53.8%] | 1.5 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | btree-map | 8 | 5604 | 2467 | 0.44× [0.42, 0.46] | -126.6% | [-136.0%, -117.2%] | 11.6 pts | 1.8 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 300 | 280 | 0.93× [0.91, 0.95] | -7.8% | [-10.1%, -5.4%] | 2.3 pts | 2.4 | no | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 304 | 270 | 0.88× [0.85, 0.91] | -13.4% | [-17.2%, -9.5%] | 4.4 pts | 2.8 | no | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 19.04 ms | 16.67 ms | 0.89× [0.86, 0.92] | -12.7% | [-16.5%, -8.9%] | 3.9 pts | 1.9 | no | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 19.47 ms | 15.58 ms | 0.82× [0.79, 0.86] | -21.3% | [-26.0%, -16.6%] | 4.6 pts | 3.1 | no | yes |
| single-value | dirs | 86215 | valuesFor | ordered | baseline | 8 | 148 | 171 | 1.15× [1.02, 1.32] | +12.9% | [+1.5%, +24.2%] | 10.7 pts | 13.2 | no | yes |
| single-value | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 159 | 258 | 1.59× [1.37, 1.91] | +37.2% | [+26.8%, +47.6%] | 9.7 pts | 8.8 | no | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 1823 | 3701 | 2.03× [1.70, 2.52] | +50.7% | [+41.1%, +60.4%] | 9.1 pts | 20.0 | no | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 1793 | 1217 | 0.69× [0.54, 0.96] | -45.5% | [-86.6%, -4.5%] | 38.4 pts | 13.0 | no | yes |
| single-value | dirs | 86215 | prefix | ordered | baseline | 8 | 34.5 µs | 84.6 µs | 2.51× [2.20, 2.91] | +60.1% | [+54.6%, +65.7%] | 5.3 pts | 13.2 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | btree-map | 8 | 34.8 µs | 16.2 µs | 0.51× [0.40, 0.71] | -96.2% | [-152.1%, -40.4%] | 55.6 pts | 17.7 | no | yes |
| single-value | dirs | 86215 | churn | ordered | baseline | 8 | 424 | 453 | 1.04× [1.00, 1.08] | +4.1% | [+0.4%, +7.8%] | 3.9 pts | 3.1 | no | yes |
| single-value | dirs | 86215 | churn | ordered | btree-map | 8 | 465 | 487 | 1.04× [1.03, 1.06] | +4.3% | [+2.6%, +6.0%] | 2.9 pts | 2.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.42% does not clear the 0.63% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.96%, 1.79%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 churn: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 build: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 13.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesFor: ordered vs baseline: 7 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 8.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 20.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 13.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesBetween: ordered vs btree-map: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value dirs n=86215 prefix: ordered vs baseline: the processes scatter 13.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 prefix: ordered vs btree-map: the processes scatter 17.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 churn: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 churn: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
