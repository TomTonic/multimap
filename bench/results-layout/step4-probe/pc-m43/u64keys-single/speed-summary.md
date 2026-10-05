| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 6 | 38.2 | 20.1 | 0.53× [0.52, 0.53] | -90.4% | [-91.9%, -88.8%] | 1.5 pts | 3.0 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 38.1 | 88.0 | 2.31× [2.28, 2.34] | +56.7% | [+56.2%, +57.3%] | 0.5 pts | 4.4 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 795 | 1326 | 1.66× [1.65, 1.68] | +39.9% | [+39.5%, +40.3%] | 0.4 pts | 1.6 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 798 | 430 | 0.54× [0.53, 0.54] | -85.9% | [-88.0%, -83.7%] | 2.0 pts | 1.9 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 6 | 64.5 | 56.6 | 0.88× [0.87, 0.89] | -13.6% | [-14.7%, -12.6%] | 1.0 pts | 1.8 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 6 | 64.5 | 120 | 1.85× [1.84, 1.86] | +46.1% | [+45.8%, +46.3%] | 0.3 pts | 0.7 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 6 | 958.6 µs | 985.4 µs | 1.03× [1.02, 1.04] | +2.7% | [+1.8%, +3.5%] | 0.8 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 6 | 955.5 µs | 1.72 ms | 1.80× [1.79, 1.81] | +44.4% | [+44.2%, +44.7%] | 0.2 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 8 | 27.6 | 25.9 | 0.95× [0.93, 0.96] | -5.7% | [-7.5%, -3.9%] | 2.2 pts | 3.9 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 27.5 | 111 | 4.05× [4.00, 4.11] | +75.3% | [+75.0%, +75.6%] | 0.4 pts | 3.1 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1532 | 1794 | 1.17× [1.15, 1.19] | +14.3% | [+12.7%, +16.0%] | 2.0 pts | 1.8 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 1504 | 457 | 0.30× [0.30, 0.31] | -228.7% | [-231.1%, -226.4%] | 2.7 pts | 1.3 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 8 | 79.7 | 64.7 | 0.82× [0.81, 0.83] | -22.0% | [-23.6%, -20.5%] | 1.5 pts | 1.4 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 8 | 80.3 | 166 | 2.07× [2.03, 2.12] | +51.8% | [+50.8%, +52.7%] | 0.9 pts | 1.5 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 8 | 5.40 ms | 4.36 ms | 0.81× [0.80, 0.81] | -24.1% | [-24.7%, -23.5%] | 0.6 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 8 | 5.45 ms | 8.80 ms | 1.62× [1.60, 1.63] | +38.1% | [+37.5%, +38.7%] | 0.6 pts | 1.2 | yes | yes |
| single-value | u64 | 262144 | valuesFor | ordered | baseline | 8 | 54.8 | 52.8 | 0.98× [0.94, 1.02] | -2.3% | [-6.5%, +1.8%] | 5.3 pts | 2.9 | no | no |
| single-value | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 65.8 | 209 | 3.21× [2.93, 3.55] | +68.8% | [+65.8%, +71.8%] | 3.2 pts | 2.0 | yes | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 1094 | 1736 | 1.60× [1.58, 1.61] | +37.4% | [+36.8%, +38.1%] | 0.9 pts | 1.5 | yes | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 1054 | 719 | 0.68× [0.68, 0.69] | -46.2% | [-47.7%, -44.7%] | 1.5 pts | 1.7 | yes | yes |
| single-value | u64 | 262144 | churn | ordered | baseline | 8 | 150 | 188 | 1.24× [1.21, 1.27] | +19.4% | [+17.4%, +21.3%] | 3.0 pts | 1.0 | yes | yes |
| single-value | u64 | 262144 | churn | ordered | btree-map | 8 | 233 | 390 | 1.68× [1.65, 1.72] | +40.5% | [+39.2%, +41.8%] | 1.2 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-6.49%, 1.84%] includes zero
- single-value u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
