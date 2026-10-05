| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 90.2 | 85.6 | 0.95× [0.94, 0.96] | -5.3% | [-6.2%, -4.5%] | 0.9 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 91.5 | 160 | 1.75× [1.74, 1.77] | +43.0% | [+42.4%, +43.6%] | 0.8 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2761 | 3045 | 1.12× [1.09, 1.15] | +10.9% | [+8.3%, +13.4%] | 2.4 pts | 1.1 | no | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2776 | 5403 | 1.95× [1.92, 1.98] | +48.7% | [+47.9%, +49.6%] | 0.9 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 8 | 2976 | 3351 | 1.14× [1.12, 1.15] | +12.0% | [+10.5%, +13.4%] | 1.9 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 8 | 3028 | 7429 | 2.48× [2.46, 2.51] | +59.7% | [+59.3%, +60.1%] | 0.7 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 207 | 153 | 0.74× [0.73, 0.75] | -35.4% | [-37.4%, -33.3%] | 1.9 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 206 | 225 | 1.08× [1.07, 1.10] | +7.7% | [+6.4%, +9.0%] | 1.3 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 9.98 ms | 7.15 ms | 0.71× [0.71, 0.72] | -39.9% | [-40.8%, -39.1%] | 1.0 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 10.08 ms | 10.70 ms | 1.06× [1.05, 1.08] | +6.1% | [+5.1%, +7.0%] | 1.0 pts | 0.9 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 118 | 116 | 0.98× [0.96, 1.00] | -2.1% | [-3.7%, -0.5%] | 1.7 pts | 2.2 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 120 | 219 | 1.84× [1.82, 1.86] | +45.6% | [+45.0%, +46.3%] | 0.6 pts | 1.0 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3404 | 3759 | 1.11× [1.10, 1.11] | +9.6% | [+9.0%, +10.2%] | 0.9 pts | 1.4 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3417 | 6185 | 1.81× [1.80, 1.82] | +44.7% | [+44.4%, +45.0%] | 0.3 pts | 0.8 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 8 | 13.7 µs | 16.2 µs | 1.18× [1.14, 1.23] | +15.5% | [+12.4%, +18.6%] | 3.5 pts | 1.0 | no | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 8 | 14.6 µs | 35.3 µs | 2.31× [2.24, 2.38] | +56.7% | [+55.4%, +58.0%] | 1.3 pts | 1.2 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 265 | 205 | 0.78× [0.78, 0.78] | -28.4% | [-28.9%, -27.9%] | 0.9 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 267 | 330 | 1.22× [1.20, 1.25] | +18.3% | [+17.0%, +19.7%] | 1.4 pts | 1.5 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 52.12 ms | 38.62 ms | 0.74× [0.73, 0.75] | -34.8% | [-36.1%, -33.5%] | 1.6 pts | 1.8 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 52.67 ms | 59.52 ms | 1.13× [1.13, 1.14] | +11.8% | [+11.2%, +12.4%] | 0.6 pts | 0.9 | yes | yes |
| natural | dirs | 86215 | valuesFor | ordered | baseline | 8 | 178 | 177 | 1.00× [0.98, 1.02] | -0.2% | [-2.4%, +2.0%] | 2.3 pts | 1.9 | no | no |
| natural | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 214 | 417 | 1.94× [1.91, 1.97] | +48.4% | [+47.5%, +49.2%] | 0.9 pts | 1.2 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 4154 | 4447 | 1.09× [1.08, 1.11] | +8.6% | [+7.2%, +10.1%] | 1.5 pts | 1.9 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 5686 | 12.4 µs | 2.22× [2.16, 2.27] | +54.9% | [+53.7%, +56.0%] | 1.3 pts | 1.9 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | baseline | 8 | 106.7 µs | 121.5 µs | 1.13× [1.12, 1.13] | +11.3% | [+10.8%, +11.8%] | 0.5 pts | 0.5 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | btree-sets | 8 | 113.9 µs | 236.1 µs | 2.16× [2.10, 2.23] | +53.8% | [+52.4%, +55.2%] | 1.4 pts | 0.5 | yes | yes |
| natural | dirs | 86215 | churn | ordered | baseline | 8 | 472 | 414 | 0.86× [0.84, 0.89] | -15.9% | [-19.0%, -12.8%] | 3.1 pts | 1.1 | no | yes |
| natural | dirs | 86215 | churn | ordered | btree-sets | 8 | 575 | 663 | 1.15× [1.13, 1.17] | +13.0% | [+11.4%, +14.6%] | 1.5 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.43% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +5.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled difference of -0.20% does not clear the 1.00% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-2.39%, 2.00%] includes zero
