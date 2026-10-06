| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 139 | 124 | 0.89× [0.88, 0.90] | -12.0% | [-13.3%, -10.7%] | 1.3 pts | 1.2 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 141 | 169 | 1.20× [1.18, 1.22] | +16.7% | [+15.6%, +17.8%] | 1.2 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6451 | 6571 | 1.01× [1.01, 1.02] | +1.4% | [+0.6%, +2.3%] | 1.2 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6484 | 10.5 µs | 1.62× [1.61, 1.63] | +38.4% | [+37.9%, +38.8%] | 0.5 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 7289 | 7597 | 1.01× [0.99, 1.04] | +1.4% | [-0.7%, +3.5%] | 2.2 pts | 0.7 | no | no |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7527 | 14.8 µs | 2.03× [1.98, 2.09] | +50.8% | [+49.4%, +52.2%] | 1.4 pts | 0.6 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 268 | 185 | 0.69× [0.68, 0.71] | -44.3% | [-46.8%, -41.8%] | 2.5 pts | 1.5 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 268 | 241 | 0.91× [0.90, 0.91] | -10.5% | [-11.6%, -9.3%] | 1.2 pts | 0.7 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 12.77 ms | 8.26 ms | 0.65× [0.64, 0.65] | -54.8% | [-56.4%, -53.1%] | 1.6 pts | 2.0 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 12.77 ms | 11.33 ms | 0.89× [0.88, 0.90] | -12.4% | [-13.1%, -11.6%] | 0.8 pts | 1.1 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 172 | 158 | 0.92× [0.91, 0.94] | -8.2% | [-9.6%, -6.8%] | 1.4 pts | 1.5 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 182 | 236 | 1.32× [1.28, 1.37] | +24.3% | [+21.6%, +27.0%] | 2.5 pts | 2.8 | no | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 7015 | 7477 | 1.07× [1.05, 1.08] | +6.2% | [+5.1%, +7.4%] | 1.1 pts | 1.4 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 7013 | 11.4 µs | 1.63× [1.62, 1.65] | +38.8% | [+38.3%, +39.3%] | 0.7 pts | 1.1 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 8 | 34.4 µs | 38.6 µs | 1.14× [1.12, 1.15] | +12.0% | [+11.1%, +13.0%] | 0.9 pts | 1.2 | yes | no |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 8 | 30.7 µs | 59.5 µs | 1.95× [1.94, 1.97] | +48.8% | [+48.4%, +49.1%] | 0.4 pts | 0.6 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 8 | 333 | 264 | 0.79× [0.78, 0.81] | -26.3% | [-28.8%, -23.8%] | 2.9 pts | 2.4 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 8 | 329 | 371 | 1.12× [1.12, 1.13] | +11.0% | [+10.4%, +11.7%] | 0.9 pts | 1.3 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 8 | 63.14 ms | 45.29 ms | 0.72× [0.71, 0.73] | -38.8% | [-40.6%, -37.0%] | 2.0 pts | 1.8 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 8 | 63.14 ms | 65.90 ms | 1.04× [1.03, 1.05] | +3.6% | [+2.5%, +4.8%] | 1.2 pts | 1.6 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 264 | 278 | 1.07× [1.04, 1.09] | +6.3% | [+4.3%, +8.2%] | 2.0 pts | 2.1 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 290 | 509 | 1.73× [1.70, 1.76] | +42.2% | [+41.3%, +43.2%] | 1.3 pts | 1.3 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 9326 | 11.5 µs | 1.22× [1.19, 1.26] | +18.3% | [+16.1%, +20.4%] | 2.1 pts | 1.5 | no | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 10.1 µs | 23.9 µs | 2.40× [2.33, 2.47] | +58.3% | [+57.1%, +59.5%] | 1.2 pts | 1.5 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 207.1 µs | 261.3 µs | 1.25× [1.23, 1.26] | +19.7% | [+18.5%, +20.9%] | 1.4 pts | 0.4 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 209.8 µs | 517.2 µs | 2.50× [2.39, 2.63] | +60.1% | [+58.2%, +61.9%] | 1.9 pts | 0.9 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 581 | 540 | 0.93× [0.91, 0.95] | -7.4% | [-9.8%, -5.0%] | 2.9 pts | 0.8 | no | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 635 | 709 | 1.11× [1.10, 1.13] | +10.1% | [+9.1%, +11.2%] | 1.5 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 prefix: ordered vs baseline: the pooled difference of 1.40% does not clear the 2.23% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=4096 prefix: ordered vs baseline: the pooled interval [-0.75%, 3.54%] includes zero
- natural-str dirs n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=4096 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=4096 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 prefix: ordered vs baseline: the pooled difference of 12.05% does not clear the 12.80% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=16384 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -4.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -2.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
