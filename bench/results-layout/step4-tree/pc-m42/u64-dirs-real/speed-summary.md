| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 90.4 | 86.8 | 0.96× [0.94, 0.98] | -4.0% | [-5.9%, -2.0%] | 2.1 pts | 1.5 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 92.2 | 161 | 1.78× [1.75, 1.81] | +43.7% | [+42.7%, +44.8%] | 1.1 pts | 1.7 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2728 | 3026 | 1.11× [1.08, 1.14] | +9.9% | [+7.5%, +12.2%] | 2.7 pts | 1.3 | no | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2742 | 5427 | 1.99× [1.96, 2.01] | +49.6% | [+49.1%, +50.2%] | 0.7 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 8 | 2888 | 3301 | 1.16× [1.14, 1.18] | +13.7% | [+12.4%, +15.1%] | 1.3 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 8 | 2872 | 7464 | 2.60× [2.55, 2.66] | +61.6% | [+60.9%, +62.4%] | 0.7 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 222 | 153 | 0.69× [0.68, 0.70] | -45.2% | [-46.8%, -43.6%] | 1.5 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 225 | 230 | 1.02× [1.00, 1.04] | +2.0% | [+0.4%, +3.6%] | 1.5 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 11.69 ms | 7.25 ms | 0.62× [0.62, 0.62] | -60.8% | [-61.5%, -60.0%] | 0.8 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 11.55 ms | 10.67 ms | 0.93× [0.91, 0.94] | -7.9% | [-9.3%, -6.6%] | 1.4 pts | 1.7 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 119 | 118 | 1.00× [0.98, 1.01] | -0.3% | [-1.7%, +1.0%] | 1.4 pts | 1.5 | yes | no |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 120 | 223 | 1.85× [1.83, 1.87] | +46.0% | [+45.5%, +46.6%] | 0.5 pts | 1.0 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3403 | 3752 | 1.10× [1.09, 1.11] | +9.3% | [+8.5%, +10.1%] | 1.0 pts | 1.2 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3448 | 6229 | 1.80× [1.79, 1.82] | +44.5% | [+44.1%, +45.0%] | 0.4 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 8 | 13.3 µs | 16.3 µs | 1.24× [1.20, 1.29] | +19.6% | [+16.9%, +22.3%] | 3.1 pts | 0.6 | no | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 8 | 14.2 µs | 35.4 µs | 2.45× [2.31, 2.62] | +59.3% | [+56.7%, +61.8%] | 2.8 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 293 | 217 | 0.74× [0.73, 0.75] | -34.6% | [-36.3%, -33.0%] | 1.5 pts | 1.0 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 284 | 338 | 1.19× [1.18, 1.21] | +16.1% | [+15.1%, +17.1%] | 1.9 pts | 1.8 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 57.46 ms | 39.20 ms | 0.68× [0.67, 0.69] | -46.7% | [-48.3%, -45.1%] | 1.7 pts | 1.6 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 57.81 ms | 60.34 ms | 1.05× [1.04, 1.06] | +4.5% | [+3.5%, +5.5%] | 1.0 pts | 1.9 | yes | yes |
| natural | dirs | 86215 | valuesFor | ordered | baseline | 8 | 183 | 183 | 1.01× [0.98, 1.03] | +0.7% | [-1.9%, +3.2%] | 2.6 pts | 1.7 | no | no |
| natural | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 217 | 418 | 1.93× [1.90, 1.97] | +48.3% | [+47.4%, +49.1%] | 0.8 pts | 1.0 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 4264 | 4665 | 1.09× [1.04, 1.14] | +8.4% | [+4.2%, +12.5%] | 4.0 pts | 3.7 | no | yes |
| natural | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 5689 | 12.4 µs | 2.19× [2.14, 2.25] | +54.4% | [+53.3%, +55.6%] | 1.2 pts | 1.5 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | baseline | 8 | 104.1 µs | 119.9 µs | 1.14× [1.12, 1.16] | +12.0% | [+10.5%, +13.4%] | 1.5 pts | 1.2 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | btree-sets | 8 | 112.6 µs | 245.6 µs | 2.28× [2.23, 2.34] | +56.1% | [+55.1%, +57.2%] | 2.1 pts | 0.8 | yes | yes |
| natural | dirs | 86215 | churn | ordered | baseline | 8 | 504 | 436 | 0.84× [0.81, 0.88] | -18.5% | [-23.4%, -13.7%] | 4.6 pts | 1.2 | no | yes |
| natural | dirs | 86215 | churn | ordered | btree-sets | 8 | 601 | 663 | 1.11× [1.09, 1.12] | +9.6% | [+8.5%, +10.8%] | 1.5 pts | 1.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.34% does not clear the 0.35% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.73%, 1.04%] includes zero
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled difference of 0.68% does not clear the 1.09% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-1.88%, 3.24%] includes zero
- natural dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.90% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -2.43% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
