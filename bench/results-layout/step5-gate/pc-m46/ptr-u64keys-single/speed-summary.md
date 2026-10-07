| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | u64 | 4096 | valuesFor | ordered | baseline | 6 | 46.3 | 20.1 | 0.44× [0.43, 0.44] | -129.8% | [-132.8%, -126.8%] | 2.9 pts | 4.5 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 46.0 | 87.9 | 1.91× [1.89, 1.92] | +47.6% | [+47.1%, +48.0%] | 0.4 pts | 2.9 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 924 | 1251 | 1.35× [1.34, 1.36] | +26.0% | [+25.6%, +26.5%] | 0.4 pts | 1.2 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 928 | 440 | 0.47× [0.47, 0.48] | -111.2% | [-112.0%, -110.5%] | 0.7 pts | 1.0 | yes | yes |
| single-value-ptr | u64 | 4096 | churn | ordered | baseline | 6 | 81.5 | 60.3 | 0.74× [0.74, 0.74] | -35.3% | [-36.0%, -34.5%] | 0.7 pts | 1.1 | yes | yes |
| single-value-ptr | u64 | 4096 | churn | ordered | btree-map | 6 | 82.0 | 124 | 1.51× [1.51, 1.52] | +34.0% | [+33.7%, +34.2%] | 0.2 pts | 0.6 | yes | yes |
| single-value-ptr | u64 | 4096 | build | ordered | baseline | 6 | 1.41 ms | 1.02 ms | 0.72× [0.71, 0.72] | -39.5% | [-40.7%, -38.3%] | 1.1 pts | 1.9 | yes | yes |
| single-value-ptr | u64 | 4096 | build | ordered | btree-map | 6 | 1.41 ms | 1.78 ms | 1.27× [1.26, 1.27] | +21.0% | [+20.4%, +21.5%] | 0.5 pts | 1.3 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesFor | ordered | baseline | 8 | 36.3 | 27.2 | 0.74× [0.73, 0.76] | -34.4% | [-36.7%, -32.1%] | 2.8 pts | 4.2 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 36.1 | 113 | 3.13× [3.11, 3.14] | +68.0% | [+67.9%, +68.2%] | 0.3 pts | 1.3 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 2127 | 1822 | 0.86× [0.85, 0.87] | -16.7% | [-18.3%, -15.1%] | 1.6 pts | 1.0 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 2104 | 499 | 0.24× [0.24, 0.24] | -321.3% | [-325.2%, -317.3%] | 5.1 pts | 3.4 | yes | yes |
| single-value-ptr | u64 | 16384 | churn | ordered | baseline | 8 | 78.4 | 69.7 | 0.89× [0.87, 0.90] | -12.9% | [-14.7%, -11.0%] | 1.7 pts | 1.9 | yes | yes |
| single-value-ptr | u64 | 16384 | churn | ordered | btree-map | 8 | 79.5 | 173 | 2.17× [2.09, 2.26] | +53.9% | [+52.1%, +55.8%] | 2.0 pts | 4.5 | yes | yes |
| single-value-ptr | u64 | 16384 | build | ordered | baseline | 8 | 6.38 ms | 4.66 ms | 0.74× [0.72, 0.76] | -35.0% | [-38.3%, -31.6%] | 3.2 pts | 4.5 | yes | yes |
| single-value-ptr | u64 | 16384 | build | ordered | btree-map | 8 | 6.44 ms | 9.29 ms | 1.46× [1.44, 1.49] | +31.6% | [+30.4%, +32.8%] | 1.3 pts | 1.9 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesFor | ordered | baseline | 8 | 81.4 | 78.6 | 0.90× [0.78, 1.06] | -11.3% | [-28.7%, +6.0%] | 16.3 pts | 13.4 | no | no |
| single-value-ptr | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 115 | 233 | 2.11× [1.91, 2.35] | +52.6% | [+47.7%, +57.4%] | 4.5 pts | 9.3 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 1628 | 2678 | 1.51× [1.26, 1.87] | +33.6% | [+20.7%, +46.4%] | 13.9 pts | 28.3 | no | yes |
| single-value-ptr | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 1605 | 992 | 0.59× [0.53, 0.68] | -68.3% | [-89.5%, -47.2%] | 23.3 pts | 9.8 | no | yes |
| single-value-ptr | u64 | 262144 | churn | ordered | baseline | 8 | 200 | 221 | 1.11× [1.04, 1.18] | +9.7% | [+4.2%, +15.2%] | 5.3 pts | 1.8 | no | yes |
| single-value-ptr | u64 | 262144 | churn | ordered | btree-map | 8 | 251 | 460 | 1.82× [1.73, 1.92] | +45.1% | [+42.3%, +47.9%] | 3.4 pts | 3.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 churn: ordered vs btree-map: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 build: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-28.66%, 6.02%] includes zero
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 13.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr u64 n=262144 valuesFor: ordered vs btree-map: the processes scatter 9.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 28.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesBetween: ordered vs btree-map: the processes scatter 9.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 churn: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
