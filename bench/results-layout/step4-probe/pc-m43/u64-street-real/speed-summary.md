| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 60.6 | 54.3 | 0.89× [0.88, 0.90] | -12.2% | [-14.0%, -10.5%] | 1.9 pts | 1.6 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 61.4 | 135 | 2.21× [2.19, 2.23] | +54.8% | [+54.4%, +55.2%] | 0.4 pts | 1.0 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 2214 | 2620 | 1.17× [1.14, 1.20] | +14.2% | [+12.1%, +16.3%] | 2.0 pts | 1.1 | no | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 2213 | 4582 | 2.09× [2.06, 2.12] | +52.2% | [+51.5%, +52.8%] | 0.8 pts | 1.1 | yes | yes |
| natural | street | 4096 | prefix | ordered | baseline | 8 | 227 | 259 | 1.14× [1.13, 1.16] | +12.6% | [+11.6%, +13.6%] | 1.0 pts | 0.8 | yes | yes |
| natural | street | 4096 | prefix | ordered | btree-sets | 8 | 228 | 526 | 2.30× [2.28, 2.32] | +56.6% | [+56.2%, +57.0%] | 0.4 pts | 0.7 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 163 | 105 | 0.64× [0.63, 0.65] | -55.6% | [-58.1%, -53.1%] | 2.9 pts | 1.3 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 161 | 186 | 1.16× [1.15, 1.17] | +13.8% | [+13.1%, +14.4%] | 0.7 pts | 0.7 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 6.32 ms | 4.12 ms | 0.65× [0.64, 0.66] | -53.8% | [-55.4%, -52.2%] | 1.6 pts | 1.1 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 6.31 ms | 7.35 ms | 1.16× [1.13, 1.18] | +13.5% | [+11.6%, +15.5%] | 2.2 pts | 2.3 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 6 | 78.7 | 74.0 | 0.93× [0.92, 0.95] | -7.0% | [-8.7%, -5.2%] | 1.6 pts | 1.8 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 6 | 80.0 | 181 | 2.27× [2.24, 2.31] | +56.0% | [+55.3%, +56.7%] | 0.7 pts | 1.8 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 6 | 2746 | 3179 | 1.15× [1.13, 1.17] | +13.2% | [+11.7%, +14.6%] | 1.4 pts | 1.7 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 2746 | 5361 | 1.95× [1.94, 1.96] | +48.7% | [+48.4%, +49.1%] | 0.3 pts | 1.1 | yes | yes |
| natural | street | 16384 | prefix | ordered | baseline | 6 | 802 | 940 | 1.16× [1.15, 1.18] | +14.1% | [+12.7%, +15.5%] | 1.3 pts | 0.8 | yes | yes |
| natural | street | 16384 | prefix | ordered | btree-sets | 6 | 811 | 2047 | 2.53× [2.48, 2.59] | +60.5% | [+59.7%, +61.3%] | 0.8 pts | 1.0 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 6 | 201 | 143 | 0.72× [0.71, 0.72] | -39.7% | [-40.8%, -38.6%] | 1.0 pts | 0.7 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 6 | 203 | 261 | 1.29× [1.27, 1.30] | +22.4% | [+21.5%, +23.3%] | 0.8 pts | 1.0 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 6 | 31.08 ms | 21.81 ms | 0.70× [0.69, 0.71] | -42.8% | [-44.0%, -41.5%] | 1.2 pts | 1.3 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 6 | 31.73 ms | 40.59 ms | 1.28× [1.27, 1.30] | +22.0% | [+21.1%, +22.9%] | 0.8 pts | 0.9 | yes | yes |
| natural | street | 212449 | valuesFor | ordered | baseline | 8 | 186 | 204 | 1.08× [1.04, 1.12] | +7.4% | [+4.0%, +10.9%] | 3.2 pts | 2.4 | no | yes |
| natural | street | 212449 | valuesFor | ordered | btree-sets | 8 | 240 | 484 | 2.03× [1.96, 2.10] | +50.7% | [+49.0%, +52.4%] | 1.8 pts | 2.5 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | baseline | 8 | 4818 | 5620 | 1.18× [1.15, 1.21] | +15.2% | [+12.9%, +17.6%] | 2.8 pts | 2.5 | no | yes |
| natural | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 5538 | 16.1 µs | 2.89× [2.82, 2.96] | +65.4% | [+64.5%, +66.3%] | 0.9 pts | 2.0 | yes | yes |
| natural | street | 212449 | prefix | ordered | baseline | 8 | 13.8 µs | 17.1 µs | 1.24× [1.21, 1.26] | +19.1% | [+17.6%, +20.6%] | 1.6 pts | 2.9 | yes | yes |
| natural | street | 212449 | prefix | ordered | btree-sets | 8 | 17.8 µs | 53.0 µs | 2.95× [2.86, 3.03] | +66.1% | [+65.1%, +67.0%] | 1.0 pts | 1.4 | yes | yes |
| natural | street | 212449 | churn | ordered | baseline | 8 | 512 | 459 | 0.88× [0.86, 0.89] | -13.9% | [-15.8%, -11.9%] | 3.6 pts | 0.6 | yes | yes |
| natural | street | 212449 | churn | ordered | btree-sets | 8 | 607 | 717 | 1.19× [1.17, 1.21] | +16.2% | [+14.8%, +17.5%] | 1.8 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.59% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 build: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +0.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -1.12% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
