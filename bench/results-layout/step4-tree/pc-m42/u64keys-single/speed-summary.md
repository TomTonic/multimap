| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 6 | 38.2 | 20.3 | 0.53× [0.53, 0.54] | -87.6% | [-89.0%, -86.2%] | 1.3 pts | 2.5 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 38.0 | 88.2 | 2.32× [2.31, 2.34] | +57.0% | [+56.7%, +57.2%] | 0.2 pts | 1.9 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 954 | 1287 | 1.35× [1.35, 1.36] | +26.1% | [+25.7%, +26.6%] | 0.4 pts | 1.4 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 952 | 438 | 0.46× [0.46, 0.46] | -117.9% | [-119.1%, -116.8%] | 1.1 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 6 | 2471 | 63.3 | 0.03× [0.02, 0.03] | -3885.2% | [-4001.6%, -3768.9%] | 110.9 pts | 1.0 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 6 | 2472 | 130 | 0.05× [0.05, 0.05] | -1816.7% | [-1846.7%, -1786.7%] | 28.6 pts | 1.1 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 6 | 25.32 ms | 998.0 µs | 0.04× [0.04, 0.04] | -2431.8% | [-2482.8%, -2380.8%] | 48.6 pts | 2.1 | yes | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 6 | 25.08 ms | 1.75 ms | 0.07× [0.07, 0.07] | -1331.9% | [-1353.1%, -1310.6%] | 20.3 pts | 2.1 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 8 | 27.8 | 26.5 | 0.95× [0.93, 0.98] | -5.0% | [-7.6%, -2.4%] | 2.5 pts | 4.8 | no | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 28.0 | 113 | 4.09× [4.01, 4.17] | +75.5% | [+75.1%, +76.0%] | 0.5 pts | 5.6 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1592 | 1794 | 1.13× [1.11, 1.14] | +11.2% | [+10.2%, +12.1%] | 1.1 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 1553 | 471 | 0.30× [0.30, 0.31] | -230.2% | [-234.4%, -226.0%] | 4.0 pts | 1.9 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 8 | 2080 | 104 | 0.05× [0.05, 0.05] | -1904.3% | [-1950.2%, -1858.4%] | 52.0 pts | 0.5 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 8 | 2090 | 188 | 0.09× [0.09, 0.09] | -987.3% | [-1002.2%, -972.4%] | 20.8 pts | 1.2 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 8 | 71.68 ms | 4.66 ms | 0.06× [0.06, 0.07] | -1444.5% | [-1477.3%, -1411.7%] | 31.5 pts | 1.5 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 8 | 71.59 ms | 9.15 ms | 0.13× [0.13, 0.13] | -683.0% | [-688.1%, -677.8%] | 7.0 pts | 1.1 | yes | yes |
| single-value | u64 | 262144 | valuesFor | ordered | baseline | 8 | 53.6 | 53.3 | 0.97× [0.92, 1.03] | -2.8% | [-8.9%, +3.3%] | 5.8 pts | 3.8 | no | no |
| single-value | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 64.2 | 212 | 3.22× [2.90, 3.62] | +69.0% | [+65.6%, +72.4%] | 3.3 pts | 1.8 | yes | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 1159 | 1763 | 1.52× [1.44, 1.60] | +34.0% | [+30.7%, +37.4%] | 3.4 pts | 5.7 | yes | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 1119 | 722 | 0.65× [0.64, 0.65] | -54.2% | [-55.7%, -52.8%] | 1.7 pts | 1.7 | yes | yes |
| single-value | u64 | 262144 | churn | ordered | baseline | 8 | 2716 | 360 | 0.14× [0.13, 0.14] | -638.5% | [-653.6%, -623.3%] | 14.5 pts | 0.6 | yes | yes |
| single-value | u64 | 262144 | churn | ordered | btree-map | 8 | 2911 | 624 | 0.21× [0.19, 0.22] | -386.5% | [-415.4%, -357.6%] | 31.3 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -2.88% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -1.11% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=4096 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 build: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -3.98% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of -1.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-8.86%, 3.29%] includes zero
- single-value u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 5.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of -2.08% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=262144 churn: ordered vs btree-map: the A/A validations found a systematic difference of -5.24% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
