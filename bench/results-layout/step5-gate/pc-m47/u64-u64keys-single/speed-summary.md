| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 8 | 33.4 | 20.1 | 0.60× [0.60, 0.60] | -66.4% | [-67.5%, -65.3%] | 1.0 pts | 1.9 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 8 | 33.2 | 88.5 | 2.67× [2.66, 2.69] | +62.6% | [+62.4%, +62.8%] | 0.2 pts | 1.7 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 883 | 1307 | 1.48× [1.47, 1.49] | +32.5% | [+32.1%, +32.9%] | 0.4 pts | 1.7 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 8 | 884 | 443 | 0.50× [0.49, 0.50] | -100.5% | [-102.2%, -98.7%] | 2.7 pts | 2.7 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 8 | 78.8 | 58.8 | 0.75× [0.74, 0.76] | -34.0% | [-36.0%, -32.1%] | 1.9 pts | 2.4 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 8 | 79.3 | 123 | 1.55× [1.54, 1.55] | +35.3% | [+35.2%, +35.4%] | 0.2 pts | 0.5 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 8 | 1.21 ms | 1.15 ms | 0.93× [0.87, 0.99] | -7.9% | [-14.5%, -1.4%] | 6.5 pts | 1.8 | no | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 8 | 1.20 ms | 1.82 ms | 1.53× [1.49, 1.58] | +34.7% | [+32.7%, +36.7%] | 1.9 pts | 2.0 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 8 | 25.9 | 26.3 | 1.01× [1.01, 1.01] | +1.2% | [+0.8%, +1.5%] | 0.7 pts | 1.6 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 25.6 | 111 | 4.26× [4.19, 4.34] | +76.6% | [+76.1%, +77.0%] | 1.5 pts | 13.9 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1607 | 1817 | 1.13× [1.12, 1.14] | +11.2% | [+10.4%, +12.1%] | 1.1 pts | 0.8 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 1579 | 474 | 0.30× [0.30, 0.30] | -233.0% | [-235.4%, -230.6%] | 2.5 pts | 1.3 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 8 | 78.5 | 71.9 | 0.91× [0.90, 0.91] | -10.1% | [-10.7%, -9.4%] | 0.7 pts | 0.6 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 8 | 79.2 | 173 | 2.15× [2.06, 2.25] | +53.5% | [+51.4%, +55.6%] | 2.2 pts | 3.1 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 8 | 6.01 ms | 5.03 ms | 0.83× [0.80, 0.87] | -19.9% | [-24.9%, -14.8%] | 4.8 pts | 1.9 | no | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 8 | 6.21 ms | 9.53 ms | 1.54× [1.49, 1.59] | +35.1% | [+33.1%, +37.0%] | 1.9 pts | 1.4 | yes | yes |
| single-value | u64 | 262144 | valuesFor | ordered | baseline | 8 | 53.2 | 62.0 | 1.09× [0.92, 1.33] | +8.1% | [-8.7%, +24.9%] | 18.1 pts | 19.9 | no | no |
| single-value | u64 | 262144 | valuesFor | ordered | btree-map | 8 | 75.1 | 256 | 3.31× [2.89, 3.86] | +69.8% | [+65.4%, +74.1%] | 4.4 pts | 6.4 | yes | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 1281 | 2175 | 1.68× [1.44, 2.02] | +40.5% | [+30.7%, +50.4%] | 11.9 pts | 29.8 | no | yes |
| single-value | u64 | 262144 | valuesBetween | ordered | btree-map | 8 | 1297 | 957 | 0.72× [0.63, 0.85] | -38.1% | [-59.1%, -17.1%] | 20.3 pts | 14.7 | no | yes |
| single-value | u64 | 262144 | churn | ordered | baseline | 8 | 212 | 237 | 1.11× [1.07, 1.15] | +9.7% | [+6.8%, +12.7%] | 2.9 pts | 1.2 | no | yes |
| single-value | u64 | 262144 | churn | ordered | btree-map | 8 | 278 | 478 | 1.80× [1.72, 1.89] | +44.6% | [+41.9%, +47.2%] | 2.9 pts | 2.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value u64 n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 13.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-8.69%, 24.92%] includes zero
- single-value u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 19.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=262144 valuesFor: ordered vs btree-map: the processes scatter 6.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 29.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 valuesBetween: ordered vs btree-map: the processes scatter 14.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=262144 churn: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
