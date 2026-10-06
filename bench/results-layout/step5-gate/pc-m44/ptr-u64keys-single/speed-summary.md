| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | u64 | 4096 | valuesFor | ordered | baseline | 6 | 37.3 | 20.5 | 0.55× [0.54, 0.55] | -82.7% | [-85.1%, -80.3%] | 2.3 pts | 5.3 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 37.1 | 88.4 | 2.38× [2.36, 2.39] | +57.9% | [+57.6%, +58.2%] | 0.3 pts | 2.3 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 923 | 1280 | 1.39× [1.39, 1.39] | +28.0% | [+27.8%, +28.2%] | 0.2 pts | 1.0 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 927 | 433 | 0.47× [0.47, 0.47] | -113.6% | [-114.4%, -112.8%] | 0.8 pts | 1.3 | yes | yes |
| single-value-ptr | u64 | 4096 | churn | ordered | baseline | 6 | 147 | 61.5 | 0.41× [0.41, 0.41] | -144.5% | [-145.7%, -143.2%] | 1.2 pts | 0.3 | yes | yes |
| single-value-ptr | u64 | 4096 | churn | ordered | btree-map | 6 | 145 | 122 | 0.83× [0.83, 0.84] | -19.8% | [-20.8%, -18.8%] | 0.9 pts | 0.9 | yes | yes |
| single-value-ptr | u64 | 4096 | build | ordered | baseline | 6 | 1.86 ms | 1.00 ms | 0.54× [0.54, 0.55] | -84.5% | [-86.8%, -82.3%] | 2.1 pts | 1.1 | yes | yes |
| single-value-ptr | u64 | 4096 | build | ordered | btree-map | 6 | 1.88 ms | 1.75 ms | 0.93× [0.92, 0.95] | -7.4% | [-9.0%, -5.8%] | 1.5 pts | 1.1 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesFor | ordered | baseline | 8 | 29.9 | 27.2 | 0.91× [0.90, 0.92] | -10.3% | [-11.3%, -9.2%] | 1.9 pts | 3.6 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 29.9 | 113 | 3.78× [3.73, 3.83] | +73.5% | [+73.2%, +73.9%] | 0.4 pts | 3.0 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1651 | 1817 | 1.10× [1.08, 1.13] | +9.3% | [+7.3%, +11.3%] | 2.2 pts | 1.6 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 1630 | 494 | 0.30× [0.30, 0.30] | -230.3% | [-232.6%, -228.0%] | 2.4 pts | 1.9 | yes | yes |
| single-value-ptr | u64 | 16384 | churn | ordered | baseline | 8 | 91.9 | 69.5 | 0.76× [0.75, 0.76] | -31.8% | [-32.5%, -31.1%] | 0.8 pts | 0.5 | yes | yes |
| single-value-ptr | u64 | 16384 | churn | ordered | btree-map | 8 | 92.1 | 168 | 1.81× [1.78, 1.85] | +44.7% | [+43.7%, +45.8%] | 1.3 pts | 2.1 | yes | yes |
| single-value-ptr | u64 | 16384 | build | ordered | baseline | 8 | 7.69 ms | 4.63 ms | 0.60× [0.60, 0.61] | -65.4% | [-66.7%, -64.0%] | 1.7 pts | 0.7 | yes | yes |
| single-value-ptr | u64 | 16384 | build | ordered | btree-map | 8 | 7.76 ms | 9.01 ms | 1.17× [1.15, 1.19] | +14.6% | [+13.1%, +16.0%] | 1.4 pts | 0.8 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesFor | ordered | baseline | 8 | 69.0 | 64.7 | 0.96× [0.90, 1.01] | -4.7% | [-10.8%, +1.4%] | 5.7 pts | 3.4 | no | no |
| single-value-ptr | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 82.9 | 214 | 2.60× [2.50, 2.72] | +61.6% | [+59.9%, +63.3%] | 1.8 pts | 1.5 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 1581 | 1989 | 1.26× [1.24, 1.28] | +20.7% | [+19.3%, +22.0%] | 1.6 pts | 1.5 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 1499 | 818 | 0.55× [0.54, 0.56] | -82.3% | [-84.7%, -79.8%] | 2.7 pts | 1.5 | yes | yes |
| single-value-ptr | u64 | 262144 | churn | ordered | baseline | 8 | 264 | 233 | 0.88× [0.86, 0.91] | -13.3% | [-16.1%, -10.4%] | 2.9 pts | 1.0 | no | yes |
| single-value-ptr | u64 | 262144 | churn | ordered | btree-map | 8 | 312 | 395 | 1.27× [1.25, 1.30] | +21.3% | [+19.9%, +22.8%] | 1.4 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of +1.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-10.76%, 1.42%] includes zero
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr u64 n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.82% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
