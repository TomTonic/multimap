| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 76.2 | 78.8 | 1.03× [1.01, 1.05] | +2.8% | [+1.4%, +4.3%] | 1.5 pts | 1.6 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 77.2 | 109 | 1.43× [1.42, 1.44] | +30.1% | [+29.6%, +30.6%] | 0.8 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 207 | 205 | 0.99× [0.98, 1.00] | -1.2% | [-1.9%, -0.5%] | 1.0 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 204 | 181 | 0.88× [0.86, 0.89] | -13.9% | [-15.6%, -12.2%] | 2.4 pts | 1.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 2.82 ms | 2.90 ms | 1.02× [1.02, 1.03] | +2.3% | [+1.7%, +2.8%] | 0.6 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 2.92 ms | 2.64 ms | 0.91× [0.89, 0.92] | -10.4% | [-12.4%, -8.3%] | 2.0 pts | 2.0 | no | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 103 | 108 | 1.04× [1.03, 1.06] | +4.2% | [+2.9%, +5.4%] | 1.6 pts | 1.6 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 105 | 158 | 1.52× [1.50, 1.53] | +34.0% | [+33.3%, +34.8%] | 1.5 pts | 1.7 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 264 | 260 | 0.99× [0.98, 1.00] | -0.8% | [-2.0%, +0.3%] | 1.1 pts | 1.0 | yes | no |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 264 | 246 | 0.93× [0.91, 0.95] | -7.1% | [-9.4%, -4.7%] | 2.8 pts | 2.5 | no | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 13.45 ms | 13.67 ms | 1.02× [1.01, 1.02] | +1.7% | [+1.4%, +2.0%] | 0.4 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 13.68 ms | 13.39 ms | 0.97× [0.95, 0.99] | -3.0% | [-4.8%, -1.2%] | 1.7 pts | 1.6 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 47.5 | 49.4 | 1.04× [1.03, 1.05] | +4.0% | [+3.3%, +4.8%] | 1.1 pts | 1.5 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 47.8 | 90.8 | 1.91× [1.89, 1.92] | +47.5% | [+47.1%, +47.9%] | 0.5 pts | 1.5 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 130 | 132 | 1.00× [0.98, 1.03] | +0.5% | [-2.0%, +3.0%] | 2.7 pts | 1.9 | no | no |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 131 | 141 | 1.07× [1.06, 1.09] | +6.8% | [+5.3%, +8.3%] | 1.8 pts | 1.5 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 1.83 ms | 1.90 ms | 1.04× [1.04, 1.05] | +4.3% | [+3.8%, +4.8%] | 0.5 pts | 0.8 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 1.84 ms | 2.15 ms | 1.20× [1.16, 1.23] | +16.3% | [+14.1%, +18.6%] | 2.7 pts | 3.4 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 65.3 | 68.7 | 1.05× [1.02, 1.08] | +4.8% | [+2.2%, +7.4%] | 2.5 pts | 3.3 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 66.3 | 125 | 1.89× [1.83, 1.95] | +47.1% | [+45.3%, +48.8%] | 1.8 pts | 3.9 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 174 | 172 | 1.00× [0.99, 1.00] | -0.2% | [-0.8%, +0.4%] | 0.7 pts | 0.6 | yes | no |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 171 | 197 | 1.16× [1.14, 1.17] | +13.5% | [+12.2%, +14.8%] | 1.5 pts | 1.9 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 8.86 ms | 9.14 ms | 1.03× [1.03, 1.04] | +3.2% | [+2.9%, +3.4%] | 0.2 pts | 0.7 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 8.99 ms | 11.34 ms | 1.26× [1.24, 1.27] | +20.5% | [+19.6%, +21.4%] | 0.9 pts | 0.7 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 6 | 19.8 | 20.6 | 1.03× [1.02, 1.05] | +3.2% | [+2.1%, +4.4%] | 1.1 pts | 4.0 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 19.5 | 88.2 | 4.53× [4.50, 4.55] | +77.9% | [+77.8%, +78.0%] | 0.1 pts | 1.8 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 6 | 53.9 | 58.3 | 1.08× [1.08, 1.09] | +7.7% | [+7.3%, +8.1%] | 0.4 pts | 0.6 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 6 | 53.6 | 123 | 2.30× [2.28, 2.32] | +56.5% | [+56.1%, +56.9%] | 0.4 pts | 1.2 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 6 | 856.9 µs | 981.1 µs | 1.15× [1.13, 1.16] | +12.7% | [+11.7%, +13.6%] | 0.9 pts | 1.6 | yes | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 6 | 865.5 µs | 1.77 ms | 2.06× [2.03, 2.08] | +51.4% | [+50.7%, +52.0%] | 0.6 pts | 2.6 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 8 | 24.6 | 26.8 | 1.09× [1.06, 1.11] | +7.9% | [+5.5%, +10.2%] | 2.2 pts | 5.7 | no | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 24.2 | 112 | 4.61× [4.53, 4.70] | +78.3% | [+77.9%, +78.7%] | 0.4 pts | 4.4 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 8 | 67.8 | 68.1 | 1.01× [1.00, 1.02] | +1.2% | [+0.3%, +2.2%] | 2.1 pts | 2.5 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 8 | 68.1 | 172 | 2.56× [2.50, 2.63] | +60.9% | [+60.0%, +61.9%] | 1.0 pts | 3.0 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 8 | 4.05 ms | 4.45 ms | 1.09× [1.09, 1.10] | +8.6% | [+8.1%, +9.0%] | 0.4 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 8 | 4.07 ms | 9.11 ms | 2.24× [2.23, 2.25] | +55.4% | [+55.2%, +55.6%] | 0.2 pts | 1.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 churn: ordered vs baseline: the pooled interval [-1.96%, 0.29%] includes zero
- single-value dirs n=16384 churn: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 churn: ordered vs baseline: the pooled difference of 0.49% does not clear the 1.29% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=4096 churn: ordered vs baseline: the pooled interval [-1.99%, 2.97%] includes zero
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs baseline: the pooled difference of -0.16% does not clear the 0.75% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=16384 churn: ordered vs baseline: the pooled interval [-0.76%, 0.44%] includes zero
- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 build: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 5.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
