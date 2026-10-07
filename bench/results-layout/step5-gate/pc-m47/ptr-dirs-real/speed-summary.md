| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 98.8 | 88.4 | 0.88× [0.88, 0.89] | -13.1% | [-14.3%, -12.0%] | 1.4 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 101 | 165 | 1.69× [1.62, 1.77] | +40.8% | [+38.3%, +43.4%] | 3.8 pts | 4.5 | yes | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2683 | 3043 | 1.13× [1.10, 1.16] | +11.2% | [+9.0%, +13.5%] | 2.6 pts | 1.3 | no | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2705 | 5775 | 2.24× [2.10, 2.41] | +55.4% | [+52.4%, +58.5%] | 3.7 pts | 5.3 | yes | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 2968 | 3291 | 1.11× [1.09, 1.12] | +9.6% | [+8.2%, +11.0%] | 1.4 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | btree-sets | 8 | 3013 | 7888 | 2.70× [2.62, 2.79] | +63.0% | [+61.8%, +64.2%] | 1.9 pts | 1.8 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 237 | 167 | 0.70× [0.70, 0.71] | -42.2% | [-43.7%, -40.7%] | 1.8 pts | 1.2 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | btree-sets | 8 | 242 | 258 | 1.07× [1.05, 1.10] | +6.9% | [+4.9%, +9.0%] | 3.4 pts | 2.2 | no | yes |
| natural-ptr | dirs | 4096 | build | ordered | baseline | 8 | 11.71 ms | 8.03 ms | 0.68× [0.67, 0.70] | -46.5% | [-50.3%, -42.7%] | 4.6 pts | 2.8 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | btree-sets | 8 | 11.90 ms | 12.13 ms | 1.03× [1.01, 1.05] | +2.9% | [+0.8%, +5.1%] | 2.5 pts | 2.3 | no | yes |
| natural-ptr | dirs | 16384 | valuesFor | ordered | baseline | 8 | 131 | 124 | 1.02× [0.91, 1.15] | +1.6% | [-10.0%, +13.2%] | 12.8 pts | 12.1 | no | no |
| natural-ptr | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 134 | 234 | 1.78× [1.71, 1.84] | +43.7% | [+41.6%, +45.8%] | 2.2 pts | 3.1 | yes | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3157 | 3847 | 1.36× [1.20, 1.58] | +26.7% | [+16.6%, +36.8%] | 11.6 pts | 17.6 | no | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3163 | 6643 | 2.18× [1.98, 2.42] | +54.1% | [+49.4%, +58.7%] | 4.4 pts | 15.5 | yes | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | baseline | 8 | 12.3 µs | 17.7 µs | 1.47× [1.33, 1.64] | +31.9% | [+24.7%, +39.2%] | 7.2 pts | 1.5 | no | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | btree-sets | 8 | 11.6 µs | 34.8 µs | 2.90× [2.81, 2.98] | +65.5% | [+64.5%, +66.5%] | 1.7 pts | 0.7 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | baseline | 8 | 290 | 231 | 0.80× [0.80, 0.81] | -24.3% | [-25.1%, -23.6%] | 2.1 pts | 1.6 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | btree-sets | 8 | 305 | 373 | 1.21× [1.17, 1.26] | +17.3% | [+14.3%, +20.4%] | 3.6 pts | 3.7 | no | yes |
| natural-ptr | dirs | 16384 | build | ordered | baseline | 8 | 59.04 ms | 44.98 ms | 0.76× [0.75, 0.77] | -31.2% | [-32.9%, -29.6%] | 2.3 pts | 1.3 | yes | yes |
| natural-ptr | dirs | 16384 | build | ordered | btree-sets | 8 | 59.26 ms | 67.97 ms | 1.14× [1.11, 1.17] | +12.5% | [+10.1%, +14.9%] | 2.3 pts | 3.3 | no | yes |
| natural-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 198 | 200 | 1.03× [0.96, 1.11] | +3.2% | [-3.7%, +10.1%] | 8.3 pts | 10.6 | no | no |
| natural-ptr | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 226 | 452 | 2.03× [1.96, 2.11] | +50.7% | [+48.8%, +52.6%] | 2.1 pts | 3.4 | yes | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 4268 | 5498 | 1.33× [1.19, 1.52] | +25.0% | [+15.7%, +34.2%] | 10.9 pts | 11.4 | no | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 4978 | 14.7 µs | 2.93× [2.74, 3.16] | +65.9% | [+63.5%, +68.4%] | 2.7 pts | 7.9 | yes | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 92.5 µs | 135.0 µs | 1.47× [1.35, 1.61] | +31.9% | [+25.8%, +37.9%] | 7.9 pts | 8.3 | no | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | btree-sets | 8 | 100.8 µs | 334.5 µs | 3.34× [3.04, 3.71] | +70.1% | [+67.1%, +73.0%] | 2.9 pts | 2.6 | yes | yes |
| natural-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 506 | 494 | 0.97× [0.89, 1.06] | -3.0% | [-11.8%, +5.8%] | 10.4 pts | 4.5 | no | no |
| natural-ptr | dirs | 86215 | churn | ordered | btree-sets | 8 | 576 | 701 | 1.24× [1.22, 1.27] | +19.5% | [+17.7%, +21.3%] | 1.9 pts | 3.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr dirs n=4096 valuesFor: ordered vs btree-sets: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.43% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=4096 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 build: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=4096 build: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of 1.58% does not clear the 1.64% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-10.05%, 13.21%] includes zero
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 12.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 17.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 15.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 churn: ordered vs btree-sets: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=16384 build: ordered vs btree-sets: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-3.74%, 10.07%] includes zero
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 10.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 11.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 7.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 prefix: ordered vs baseline: the processes scatter 8.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 prefix: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 churn: ordered vs baseline: the pooled difference of -3.02% does not clear the 4.60% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=86215 churn: ordered vs baseline: the pooled interval [-11.84%, 5.80%] includes zero
- natural-ptr dirs n=86215 churn: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 churn: ordered vs btree-sets: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
