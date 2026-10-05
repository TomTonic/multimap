| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 127 | 124 | 0.97× [0.96, 0.99] | -2.6% | [-3.7%, -1.5%] | 1.1 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 129 | 170 | 1.31× [1.29, 1.32] | +23.6% | [+22.7%, +24.5%] | 1.6 pts | 1.5 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6238 | 6477 | 1.04× [1.03, 1.05] | +3.7% | [+2.8%, +4.7%] | 1.3 pts | 1.4 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6283 | 10.5 µs | 1.67× [1.66, 1.68] | +40.2% | [+39.8%, +40.5%] | 0.4 pts | 0.7 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 6982 | 7209 | 1.05× [1.03, 1.07] | +4.7% | [+2.9%, +6.5%] | 2.2 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7250 | 14.9 µs | 2.10× [2.03, 2.19] | +52.5% | [+50.8%, +54.2%] | 2.2 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 268 | 189 | 0.71× [0.70, 0.72] | -41.4% | [-43.5%, -39.2%] | 2.1 pts | 1.2 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 265 | 243 | 0.91× [0.90, 0.92] | -9.9% | [-11.1%, -8.6%] | 1.5 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 13.28 ms | 8.35 ms | 0.63× [0.63, 0.63] | -58.8% | [-59.7%, -57.9%] | 1.0 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 13.25 ms | 11.66 ms | 0.88× [0.88, 0.89] | -13.2% | [-13.9%, -12.5%] | 0.7 pts | 0.8 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 160 | 159 | 0.98× [0.97, 0.99] | -1.6% | [-2.8%, -0.5%] | 1.2 pts | 1.4 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 167 | 237 | 1.37× [1.30, 1.45] | +26.9% | [+22.9%, +30.9%] | 4.4 pts | 4.9 | no | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 7052 | 7312 | 1.04× [1.03, 1.05] | +4.0% | [+3.0%, +5.0%] | 1.0 pts | 1.2 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 7112 | 11.7 µs | 1.64× [1.60, 1.67] | +38.9% | [+37.7%, +40.1%] | 1.4 pts | 2.2 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 8 | 34.8 µs | 36.6 µs | 1.05× [1.05, 1.06] | +5.2% | [+4.8%, +5.6%] | 0.7 pts | 1.0 | yes | no |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 8 | 32.9 µs | 60.4 µs | 1.85× [1.82, 1.87] | +45.8% | [+45.1%, +46.6%] | 0.8 pts | 1.7 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 8 | 341 | 261 | 0.76× [0.76, 0.77] | -30.8% | [-31.4%, -30.2%] | 0.6 pts | 0.6 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 8 | 338 | 369 | 1.09× [1.08, 1.10] | +8.1% | [+7.2%, +8.9%] | 0.9 pts | 1.2 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 8 | 68.22 ms | 46.60 ms | 0.69× [0.68, 0.69] | -45.8% | [-47.0%, -44.5%] | 1.7 pts | 1.3 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 8 | 66.82 ms | 66.31 ms | 1.00× [0.99, 1.01] | -0.5% | [-1.5%, +0.5%] | 1.0 pts | 1.2 | yes | no |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 269 | 278 | 1.03× [1.02, 1.04] | +2.6% | [+1.7%, +3.5%] | 1.0 pts | 1.1 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 305 | 510 | 1.62× [1.57, 1.67] | +38.1% | [+36.3%, +40.0%] | 2.1 pts | 2.0 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 10.3 µs | 10.9 µs | 1.08× [1.05, 1.10] | +7.2% | [+4.9%, +9.5%] | 2.1 pts | 1.8 | no | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 11.8 µs | 23.4 µs | 2.01× [1.96, 2.07] | +50.4% | [+48.9%, +51.8%] | 1.4 pts | 1.5 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 229.2 µs | 259.9 µs | 1.10× [1.03, 1.18] | +9.0% | [+2.5%, +15.4%] | 7.1 pts | 1.3 | no | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 234.3 µs | 527.6 µs | 2.21× [2.10, 2.33] | +54.7% | [+52.4%, +57.0%] | 3.0 pts | 1.0 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 631 | 547 | 0.86× [0.83, 0.88] | -16.9% | [-20.4%, -13.4%] | 3.8 pts | 0.9 | no | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 684 | 706 | 1.04× [1.02, 1.07] | +4.2% | [+1.8%, +6.6%] | 2.4 pts | 2.0 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 prefix: ordered vs baseline: the pooled difference of 5.21% does not clear the 15.31% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 build: ordered vs btree-sets: the pooled difference of -0.46% does not clear the 0.50% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=16384 build: ordered vs btree-sets: the pooled interval [-1.48%, 0.55%] includes zero
- natural-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -1.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
