| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 86.2 | 78.6 | 0.91× [0.90, 0.93] | -9.6% | [-11.4%, -7.9%] | 1.9 pts | 2.2 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 86.9 | 109 | 1.26× [1.24, 1.28] | +20.5% | [+19.4%, +21.7%] | 1.3 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 199 | 198 | 1.01× [1.00, 1.01] | +0.8% | [+0.5%, +1.2%] | 0.5 pts | 0.4 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 198 | 178 | 0.90× [0.89, 0.91] | -11.3% | [-12.6%, -9.9%] | 1.4 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 3.32 ms | 2.95 ms | 0.89× [0.88, 0.90] | -12.7% | [-14.2%, -11.1%] | 1.4 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 3.36 ms | 2.59 ms | 0.78× [0.77, 0.78] | -29.0% | [-29.8%, -28.2%] | 1.1 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 108 | 108 | 1.00× [0.97, 1.04] | +0.2% | [-3.2%, +3.6%] | 3.2 pts | 3.1 | no | no |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 109 | 158 | 1.48× [1.46, 1.50] | +32.6% | [+31.7%, +33.5%] | 2.5 pts | 3.4 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 236 | 249 | 1.06× [1.05, 1.08] | +6.0% | [+4.4%, +7.6%] | 1.5 pts | 1.7 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 238 | 243 | 1.03× [1.02, 1.04] | +2.6% | [+1.6%, +3.6%] | 1.0 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 14.52 ms | 13.50 ms | 0.93× [0.92, 0.93] | -7.8% | [-8.2%, -7.4%] | 0.5 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 14.70 ms | 13.13 ms | 0.89× [0.88, 0.90] | -12.0% | [-13.3%, -10.6%] | 1.3 pts | 1.3 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 57.2 | 49.0 | 0.85× [0.85, 0.86] | -17.0% | [-17.9%, -16.1%] | 0.9 pts | 1.1 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 58.0 | 91.9 | 1.60× [1.54, 1.67] | +37.5% | [+35.0%, +40.1%] | 2.5 pts | 7.3 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 136 | 126 | 0.93× [0.92, 0.95] | -7.1% | [-8.4%, -5.7%] | 1.5 pts | 1.2 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 136 | 138 | 1.02× [1.00, 1.03] | +1.7% | [+0.4%, +3.1%] | 1.3 pts | 1.2 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.24 ms | 1.88 ms | 0.84× [0.83, 0.85] | -18.8% | [-20.2%, -17.4%] | 1.3 pts | 1.5 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.26 ms | 2.11 ms | 0.95× [0.91, 0.98] | -5.5% | [-9.3%, -1.7%] | 4.4 pts | 3.4 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 68.9 | 68.5 | 0.99× [0.97, 1.02] | -0.6% | [-2.7%, +1.6%] | 2.1 pts | 2.3 | no | no |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 69.7 | 127 | 1.81× [1.78, 1.85] | +44.8% | [+43.7%, +45.9%] | 1.3 pts | 2.9 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 162 | 164 | 1.01× [1.00, 1.03] | +1.3% | [-0.3%, +2.9%] | 1.6 pts | 1.5 | yes | no |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 162 | 194 | 1.20× [1.19, 1.21] | +16.4% | [+15.7%, +17.1%] | 0.7 pts | 0.9 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 10.20 ms | 9.12 ms | 0.90× [0.88, 0.91] | -11.5% | [-13.2%, -9.9%] | 1.7 pts | 1.2 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 10.44 ms | 11.17 ms | 1.06× [1.04, 1.08] | +5.4% | [+3.8%, +7.1%] | 1.6 pts | 0.7 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 8 | 33.4 | 20.1 | 0.60× [0.60, 0.60] | -66.3% | [-66.7%, -65.8%] | 0.5 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 8 | 33.2 | 88.6 | 2.67× [2.66, 2.69] | +62.6% | [+62.4%, +62.8%] | 0.2 pts | 2.3 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 8 | 77.9 | 57.4 | 0.74× [0.73, 0.74] | -35.5% | [-36.6%, -34.5%] | 1.0 pts | 1.5 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 8 | 77.9 | 121 | 1.57× [1.51, 1.64] | +36.4% | [+33.9%, +38.9%] | 2.4 pts | 5.9 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 8 | 1.14 ms | 995.9 µs | 0.87× [0.85, 0.89] | -14.7% | [-17.5%, -11.9%] | 2.7 pts | 2.9 | no | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 8 | 1.14 ms | 1.73 ms | 1.52× [1.51, 1.52] | +34.1% | [+33.8%, +34.4%] | 0.5 pts | 2.2 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 6 | 26.2 | 26.3 | 1.00× [0.98, 1.02] | -0.2% | [-1.9%, +1.5%] | 1.6 pts | 3.4 | yes | no |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 6 | 26.1 | 112 | 4.31× [4.27, 4.35] | +76.8% | [+76.6%, +77.0%] | 0.2 pts | 2.3 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 6 | 73.6 | 66.3 | 0.91× [0.90, 0.92] | -10.0% | [-11.4%, -8.5%] | 1.4 pts | 1.5 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 6 | 74.6 | 167 | 2.26× [2.24, 2.28] | +55.7% | [+55.3%, +56.1%] | 0.4 pts | 0.7 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 6 | 5.35 ms | 4.38 ms | 0.82× [0.81, 0.82] | -22.5% | [-22.7%, -22.3%] | 0.2 pts | 0.3 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 6 | 5.44 ms | 8.90 ms | 1.63× [1.62, 1.64] | +38.7% | [+38.3%, +39.2%] | 0.4 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.20% does not clear the 0.59% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.21%, 3.61%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs btree-map: the processes scatter 7.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.71%, 1.59%] includes zero
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs baseline: the pooled interval [-0.33%, 2.91%] includes zero
- single-value u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs btree-map: the processes scatter 5.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 build: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 build: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.16% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.86%, 1.54%] includes zero
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
