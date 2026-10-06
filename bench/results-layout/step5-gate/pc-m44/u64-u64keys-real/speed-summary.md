| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | u64 | 4096 | valuesFor | ordered | baseline | 6 | 47.6 | 38.0 | 0.80× [0.79, 0.81] | -25.6% | [-27.4%, -23.8%] | 1.7 pts | 1.5 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 6 | 47.7 | 155 | 3.25× [3.25, 3.26] | +69.3% | [+69.2%, +69.4%] | 0.1 pts | 0.3 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 2898 | 2743 | 0.93× [0.92, 0.94] | -7.4% | [-8.7%, -6.2%] | 1.2 pts | 0.4 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 6 | 2881 | 7023 | 2.44× [2.42, 2.47] | +59.1% | [+58.6%, +59.5%] | 0.4 pts | 0.7 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 6 | 73.3 | 57.6 | 0.79× [0.78, 0.79] | -27.1% | [-28.3%, -26.0%] | 1.1 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 6 | 73.2 | 154 | 2.12× [2.10, 2.14] | +52.8% | [+52.4%, +53.2%] | 0.4 pts | 1.0 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 6 | 9.08 ms | 5.83 ms | 0.64× [0.64, 0.65] | -55.8% | [-57.3%, -54.2%] | 1.5 pts | 2.3 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 6 | 9.10 ms | 14.72 ms | 1.63× [1.61, 1.64] | +38.5% | [+37.9%, +39.0%] | 0.5 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 6 | 50.2 | 47.1 | 0.93× [0.93, 0.94] | -7.0% | [-8.0%, -5.9%] | 1.0 pts | 1.8 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 50.5 | 191 | 3.79× [3.75, 3.84] | +73.6% | [+73.3%, +74.0%] | 0.3 pts | 1.6 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3694 | 3596 | 0.97× [0.97, 0.98] | -2.6% | [-3.4%, -1.8%] | 0.8 pts | 1.1 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 3701 | 7486 | 2.02× [2.00, 2.04] | +50.4% | [+50.0%, +50.9%] | 0.4 pts | 1.9 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 6 | 93.0 | 76.0 | 0.81× [0.81, 0.82] | -22.9% | [-23.6%, -22.2%] | 0.7 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 6 | 99.2 | 229 | 2.33× [2.28, 2.37] | +57.0% | [+56.2%, +57.8%] | 0.7 pts | 1.7 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 6 | 34.95 ms | 28.57 ms | 0.82× [0.81, 0.82] | -22.7% | [-23.5%, -21.8%] | 0.8 pts | 1.0 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 6 | 35.44 ms | 78.38 ms | 2.20× [2.18, 2.21] | +54.5% | [+54.2%, +54.8%] | 0.3 pts | 0.7 | yes | yes |
| natural | u64 | 262144 | valuesFor | ordered | baseline | 8 | 139 | 129 | 0.95× [0.92, 0.97] | -5.7% | [-8.5%, -2.9%] | 2.9 pts | 3.1 | no | yes |
| natural | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 178 | 524 | 3.01× [2.96, 3.05] | +66.7% | [+66.2%, +67.2%] | 0.6 pts | 1.8 | yes | yes |
| natural | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 5647 | 6610 | 1.17× [1.14, 1.20] | +14.6% | [+12.5%, +16.7%] | 2.0 pts | 2.8 | no | yes |
| natural | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 5986 | 21.7 µs | 3.65× [3.60, 3.70] | +72.6% | [+72.2%, +72.9%] | 0.5 pts | 2.0 | yes | yes |
| natural | u64 | 262144 | churn | ordered | baseline | 8 | 338 | 322 | 0.91× [0.89, 0.93] | -9.9% | [-12.8%, -7.1%] | 2.9 pts | 0.5 | no | yes |
| natural | u64 | 262144 | churn | ordered | btree-sets | 8 | 429 | 696 | 1.66× [1.62, 1.71] | +39.9% | [+38.2%, +41.6%] | 1.8 pts | 1.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural u64 n=4096 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
