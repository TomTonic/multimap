| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 6 | 86.9 | 79.5 | 0.91× [0.90, 0.93] | -9.8% | [-11.7%, -7.9%] | 1.8 pts | 2.2 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 6 | 87.4 | 110 | 1.26× [1.24, 1.29] | +20.9% | [+19.3%, +22.5%] | 1.5 pts | 1.7 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 6 | 250 | 198 | 0.79× [0.78, 0.80] | -26.7% | [-28.3%, -25.1%] | 1.5 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 6 | 250 | 177 | 0.71× [0.69, 0.72] | -41.1% | [-44.1%, -38.1%] | 2.9 pts | 1.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 6 | 3.65 ms | 2.88 ms | 0.79× [0.77, 0.81] | -26.8% | [-29.3%, -24.2%] | 2.4 pts | 2.0 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 6 | 3.68 ms | 2.59 ms | 0.71× [0.70, 0.71] | -41.2% | [-42.6%, -39.9%] | 1.3 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 108 | 108 | 1.00× [0.98, 1.01] | -0.2% | [-1.6%, +1.3%] | 1.5 pts | 1.4 | yes | no |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 109 | 160 | 1.47× [1.43, 1.52] | +32.0% | [+30.0%, +34.1%] | 2.0 pts | 2.3 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 282 | 256 | 0.91× [0.90, 0.92] | -10.1% | [-11.4%, -8.7%] | 1.6 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 279 | 242 | 0.87× [0.86, 0.88] | -15.2% | [-16.1%, -14.3%] | 0.9 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 16.27 ms | 13.66 ms | 0.84× [0.83, 0.85] | -19.2% | [-20.2%, -18.2%] | 1.1 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 16.26 ms | 13.23 ms | 0.81× [0.80, 0.83] | -23.1% | [-25.3%, -20.9%] | 2.1 pts | 1.6 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 57.6 | 49.2 | 0.85× [0.84, 0.86] | -17.5% | [-19.2%, -15.9%] | 1.8 pts | 2.4 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 58.2 | 91.6 | 1.56× [1.55, 1.57] | +36.0% | [+35.7%, +36.4%] | 1.1 pts | 2.7 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 185 | 129 | 0.70× [0.69, 0.70] | -43.6% | [-45.1%, -42.1%] | 2.3 pts | 1.1 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 186 | 139 | 0.75× [0.74, 0.76] | -33.6% | [-34.7%, -32.4%] | 1.4 pts | 1.0 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.57 ms | 1.88 ms | 0.74× [0.73, 0.74] | -35.6% | [-36.8%, -34.4%] | 1.4 pts | 1.3 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.55 ms | 2.13 ms | 0.85× [0.82, 0.87] | -18.1% | [-21.3%, -15.0%] | 4.0 pts | 2.8 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 69.9 | 69.5 | 0.99× [0.98, 1.01] | -0.8% | [-2.3%, +0.7%] | 1.5 pts | 1.6 | yes | no |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 70.0 | 127 | 1.80× [1.77, 1.84] | +44.6% | [+43.6%, +45.6%] | 1.2 pts | 2.5 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 197 | 167 | 0.86× [0.82, 0.90] | -16.9% | [-22.1%, -11.7%] | 5.7 pts | 3.9 | no | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 198 | 195 | 0.98× [0.96, 1.00] | -1.7% | [-3.9%, +0.5%] | 2.2 pts | 2.3 | no | no |
| single-value | street | 16384 | build | ordered | baseline | 8 | 11.26 ms | 9.02 ms | 0.80× [0.80, 0.81] | -24.2% | [-25.2%, -23.2%] | 1.2 pts | 1.2 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 11.56 ms | 11.15 ms | 0.96× [0.95, 0.97] | -4.1% | [-5.2%, -2.9%] | 1.1 pts | 0.4 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 6 | 33.4 | 20.1 | 0.60× [0.60, 0.60] | -65.9% | [-66.2%, -65.6%] | 0.3 pts | 0.8 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 33.2 | 88.8 | 2.68× [2.65, 2.70] | +62.6% | [+62.3%, +63.0%] | 0.4 pts | 3.4 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 6 | 77.9 | 57.4 | 0.74× [0.73, 0.74] | -35.8% | [-37.1%, -34.5%] | 1.2 pts | 1.7 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 6 | 77.9 | 121 | 1.56× [1.55, 1.57] | +35.7% | [+35.4%, +36.1%] | 0.4 pts | 0.9 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 6 | 1.14 ms | 986.8 µs | 0.87× [0.86, 0.87] | -15.5% | [-16.4%, -14.6%] | 0.8 pts | 1.2 | yes | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 6 | 1.14 ms | 1.74 ms | 1.52× [1.52, 1.53] | +34.3% | [+34.0%, +34.6%] | 0.3 pts | 1.1 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 6 | 26.2 | 25.8 | 1.00× [0.98, 1.01] | -0.3% | [-1.9%, +1.3%] | 1.5 pts | 2.9 | yes | no |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 6 | 26.1 | 113 | 4.35× [4.28, 4.42] | +77.0% | [+76.6%, +77.4%] | 0.4 pts | 3.2 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 6 | 73.4 | 66.1 | 0.90× [0.89, 0.90] | -11.7% | [-12.3%, -11.1%] | 0.6 pts | 0.6 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 6 | 74.3 | 167 | 2.25× [2.20, 2.31] | +55.6% | [+54.5%, +56.6%] | 1.0 pts | 2.0 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 6 | 5.36 ms | 4.39 ms | 0.82× [0.81, 0.82] | -22.6% | [-23.4%, -21.8%] | 0.7 pts | 1.1 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 6 | 5.43 ms | 8.88 ms | 1.63× [1.62, 1.65] | +38.8% | [+38.2%, +39.3%] | 0.5 pts | 1.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.18% does not clear the 0.39% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.61%, 1.26%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.26%, 0.74%] includes zero
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs btree-map: the pooled interval [-3.90%, 0.49%] includes zero
- single-value street n=16384 churn: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.86%, 1.28%] includes zero
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.60% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
