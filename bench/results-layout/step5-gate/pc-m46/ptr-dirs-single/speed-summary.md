| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 99.5 | 78.3 | 0.78× [0.77, 0.80] | -27.5% | [-29.8%, -25.2%] | 2.1 pts | 2.4 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 100 | 109 | 1.08× [1.07, 1.09] | +7.3% | [+6.5%, +8.1%] | 1.8 pts | 1.4 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1636 | 2258 | 1.37× [1.35, 1.39] | +27.2% | [+26.1%, +28.3%] | 1.0 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1605 | 570 | 0.35× [0.35, 0.35] | -185.5% | [-188.9%, -182.0%] | 4.7 pts | 1.5 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 1996 | 2552 | 1.28× [1.26, 1.31] | +22.0% | [+20.5%, +23.5%] | 1.5 pts | 1.2 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | btree-map | 8 | 1977 | 667 | 0.34× [0.33, 0.34] | -198.0% | [-202.0%, -194.0%] | 4.8 pts | 1.4 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 274 | 207 | 0.77× [0.76, 0.77] | -30.6% | [-31.8%, -29.4%] | 2.0 pts | 1.4 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | btree-map | 8 | 274 | 187 | 0.68× [0.67, 0.68] | -47.7% | [-49.0%, -46.4%] | 2.3 pts | 1.2 | yes | yes |
| single-value-ptr | dirs | 4096 | build | ordered | baseline | 8 | 4.25 ms | 3.02 ms | 0.71× [0.69, 0.73] | -40.8% | [-44.4%, -37.3%] | 3.4 pts | 1.8 | yes | yes |
| single-value-ptr | dirs | 4096 | build | ordered | btree-map | 8 | 4.22 ms | 2.77 ms | 0.65× [0.64, 0.67] | -52.9% | [-56.6%, -49.3%] | 3.4 pts | 1.9 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | baseline | 8 | 124 | 109 | 0.89× [0.86, 0.92] | -12.6% | [-16.9%, -8.4%] | 4.3 pts | 4.8 | no | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 125 | 159 | 1.27× [1.25, 1.29] | +21.3% | [+20.3%, +22.2%] | 1.1 pts | 1.2 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1774 | 2839 | 1.60× [1.58, 1.62] | +37.6% | [+36.8%, +38.3%] | 0.7 pts | 1.5 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 1760 | 769 | 0.44× [0.43, 0.44] | -127.8% | [-129.9%, -125.7%] | 2.2 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | baseline | 8 | 6690 | 11.2 µs | 1.64× [1.56, 1.72] | +38.9% | [+35.7%, +42.0%] | 2.9 pts | 0.7 | yes | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | btree-map | 8 | 6925 | 2683 | 0.40× [0.39, 0.41] | -151.0% | [-157.6%, -144.4%] | 6.2 pts | 0.8 | yes | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | baseline | 8 | 310 | 266 | 0.86× [0.85, 0.88] | -15.9% | [-17.8%, -14.1%] | 2.2 pts | 1.9 | yes | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | btree-map | 8 | 319 | 265 | 0.85× [0.82, 0.88] | -17.6% | [-21.5%, -13.8%] | 4.1 pts | 4.0 | no | yes |
| single-value-ptr | dirs | 16384 | build | ordered | baseline | 8 | 20.59 ms | 16.05 ms | 0.78× [0.76, 0.80] | -28.3% | [-31.1%, -25.5%] | 3.3 pts | 1.9 | no | yes |
| single-value-ptr | dirs | 16384 | build | ordered | btree-map | 8 | 20.99 ms | 15.89 ms | 0.75× [0.73, 0.77] | -33.1% | [-36.4%, -29.9%] | 3.7 pts | 2.5 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 169 | 164 | 1.01× [0.99, 1.04] | +1.4% | [-0.8%, +3.7%] | 11.2 pts | 11.2 | no | no |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 185 | 269 | 1.53× [1.27, 1.93] | +34.7% | [+21.3%, +48.1%] | 13.8 pts | 11.4 | no | yes |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 2017 | 3415 | 1.75× [1.71, 1.78] | +42.7% | [+41.6%, +43.8%] | 8.6 pts | 14.4 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 2075 | 1706 | 0.76× [0.58, 1.09] | -32.3% | [-73.1%, +8.5%] | 43.6 pts | 23.4 | no | no |
| single-value-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 41.0 µs | 82.0 µs | 2.09× [2.07, 2.11] | +52.2% | [+51.7%, +52.6%] | 8.5 pts | 21.1 | yes | yes |
| single-value-ptr | dirs | 86215 | prefix | ordered | btree-map | 8 | 41.8 µs | 25.5 µs | 0.60× [0.42, 1.01] | -67.5% | [-135.9%, +0.9%] | 75.6 pts | 26.9 | no | no |
| single-value-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 431 | 436 | 1.01× [0.98, 1.04] | +0.9% | [-2.4%, +4.1%] | 3.2 pts | 2.4 | no | no |
| single-value-ptr | dirs | 86215 | churn | ordered | btree-map | 8 | 479 | 469 | 0.99× [0.97, 1.00] | -1.5% | [-3.4%, +0.5%] | 2.3 pts | 1.8 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr dirs n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 churn: ordered vs btree-map: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 build: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-0.80%, 3.67%] includes zero
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 11.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 11.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 14.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesBetween: ordered vs btree-map: the pooled interval [-73.10%, 8.46%] includes zero
- single-value-ptr dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 23.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesBetween: ordered vs btree-map: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr dirs n=86215 prefix: ordered vs baseline: the processes scatter 21.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 prefix: ordered vs btree-map: the pooled interval [-135.94%, 0.88%] includes zero
- single-value-ptr dirs n=86215 prefix: ordered vs btree-map: the processes scatter 26.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 prefix: ordered vs btree-map: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr dirs n=86215 churn: ordered vs baseline: the pooled difference of 0.85% does not clear the 1.35% noise floor, the bound on what the harness reports between identical code in every process
- single-value-ptr dirs n=86215 churn: ordered vs baseline: the pooled interval [-2.44%, 4.14%] includes zero
- single-value-ptr dirs n=86215 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 churn: ordered vs btree-map: the pooled interval [-3.45%, 0.48%] includes zero
