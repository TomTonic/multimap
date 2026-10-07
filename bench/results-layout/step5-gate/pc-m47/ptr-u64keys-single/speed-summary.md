| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | u64 | 4096 | valuesFor | ordered | baseline | 8 | 33.9 | 20.3 | 0.60× [0.60, 0.61] | -66.2% | [-67.2%, -65.2%] | 1.3 pts | 3.5 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesFor | ordered | btree-map | 8 | 33.7 | 88.3 | 2.63× [2.62, 2.64] | +62.0% | [+61.8%, +62.1%] | 0.2 pts | 1.2 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 900 | 1274 | 1.41× [1.40, 1.42] | +29.0% | [+28.6%, +29.4%] | 0.4 pts | 1.5 | yes | yes |
| single-value-ptr | u64 | 4096 | valuesBetween | ordered | btree-map | 8 | 902 | 444 | 0.49× [0.49, 0.50] | -103.4% | [-105.5%, -101.2%] | 2.0 pts | 2.9 | yes | yes |
| single-value-ptr | u64 | 4096 | churn | ordered | baseline | 8 | 82.2 | 60.7 | 0.74× [0.74, 0.75] | -34.3% | [-35.9%, -32.7%] | 1.8 pts | 2.2 | yes | yes |
| single-value-ptr | u64 | 4096 | churn | ordered | btree-map | 8 | 81.6 | 125 | 1.54× [1.50, 1.59] | +35.1% | [+33.3%, +37.0%] | 1.8 pts | 4.5 | yes | yes |
| single-value-ptr | u64 | 4096 | build | ordered | baseline | 8 | 1.50 ms | 1.07 ms | 0.72× [0.70, 0.73] | -39.4% | [-42.4%, -36.4%] | 2.9 pts | 0.9 | yes | yes |
| single-value-ptr | u64 | 4096 | build | ordered | btree-map | 8 | 1.57 ms | 1.83 ms | 1.21× [1.18, 1.24] | +17.1% | [+15.2%, +19.1%] | 2.0 pts | 0.9 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesFor | ordered | baseline | 8 | 27.7 | 27.3 | 0.98× [0.97, 0.99] | -2.0% | [-3.0%, -0.9%] | 1.5 pts | 3.6 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 27.4 | 113 | 4.11× [4.09, 4.14] | +75.7% | [+75.5%, +75.8%] | 0.3 pts | 2.0 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1644 | 1823 | 1.11× [1.09, 1.13] | +9.7% | [+8.2%, +11.2%] | 1.6 pts | 1.4 | yes | yes |
| single-value-ptr | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 1607 | 503 | 0.31× [0.31, 0.32] | -219.9% | [-224.1%, -215.7%] | 5.0 pts | 3.8 | yes | yes |
| single-value-ptr | u64 | 16384 | churn | ordered | baseline | 8 | 82.3 | 74.0 | 0.89× [0.88, 0.91] | -11.8% | [-13.2%, -10.4%] | 1.3 pts | 1.1 | yes | yes |
| single-value-ptr | u64 | 16384 | churn | ordered | btree-map | 8 | 83.1 | 175 | 2.00× [1.89, 2.12] | +50.0% | [+47.1%, +52.9%] | 3.6 pts | 6.5 | yes | yes |
| single-value-ptr | u64 | 16384 | build | ordered | baseline | 8 | 6.82 ms | 5.24 ms | 0.77× [0.75, 0.78] | -30.5% | [-33.4%, -27.5%] | 2.9 pts | 1.4 | yes | yes |
| single-value-ptr | u64 | 16384 | build | ordered | btree-map | 8 | 7.01 ms | 9.52 ms | 1.37× [1.32, 1.42] | +26.7% | [+24.1%, +29.3%] | 2.7 pts | 1.6 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesFor | ordered | baseline | 8 | 103 | 89.8 | 0.98× [0.85, 1.14] | -2.4% | [-17.2%, +12.3%] | 14.4 pts | 15.5 | no | no |
| single-value-ptr | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 106 | 250 | 2.50× [2.31, 2.72] | +59.9% | [+56.7%, +63.2%] | 5.0 pts | 8.4 | yes | yes |
| single-value-ptr | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 1976 | 2856 | 1.48× [1.34, 1.66] | +32.5% | [+25.2%, +39.8%] | 9.2 pts | 17.7 | no | yes |
| single-value-ptr | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 1820 | 1158 | 0.64× [0.59, 0.69] | -56.5% | [-68.9%, -44.1%] | 15.4 pts | 10.5 | no | yes |
| single-value-ptr | u64 | 262144 | churn | ordered | baseline | 8 | 211 | 241 | 1.14× [1.10, 1.19] | +12.3% | [+8.9%, +15.7%] | 3.2 pts | 1.5 | no | yes |
| single-value-ptr | u64 | 262144 | churn | ordered | btree-map | 8 | 269 | 463 | 1.74× [1.66, 1.82] | +42.5% | [+39.8%, +45.2%] | 2.8 pts | 2.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=4096 churn: ordered vs btree-map: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=16384 churn: ordered vs btree-map: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-17.18%, 12.32%] includes zero
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 15.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr u64 n=262144 valuesFor: ordered vs btree-map: the processes scatter 8.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 17.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 valuesBetween: ordered vs btree-map: the processes scatter 10.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr u64 n=262144 churn: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
