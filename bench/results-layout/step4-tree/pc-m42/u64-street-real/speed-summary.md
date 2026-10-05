| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 60.7 | 55.3 | 0.90× [0.89, 0.91] | -10.8% | [-12.1%, -9.4%] | 1.4 pts | 1.3 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 61.3 | 136 | 2.23× [2.20, 2.26] | +55.1% | [+54.5%, +55.7%] | 0.6 pts | 1.6 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 2190 | 2619 | 1.19× [1.18, 1.20] | +16.0% | [+15.0%, +17.0%] | 1.5 pts | 0.9 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 2198 | 4563 | 2.09× [2.05, 2.12] | +52.1% | [+51.3%, +52.9%] | 0.8 pts | 1.2 | yes | yes |
| natural | street | 4096 | prefix | ordered | baseline | 8 | 233 | 259 | 1.11× [1.10, 1.12] | +9.9% | [+9.3%, +10.6%] | 1.4 pts | 1.2 | yes | yes |
| natural | street | 4096 | prefix | ordered | btree-sets | 8 | 238 | 528 | 2.22× [2.19, 2.26] | +55.0% | [+54.3%, +55.8%] | 0.8 pts | 1.6 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 179 | 107 | 0.60× [0.59, 0.60] | -67.4% | [-69.0%, -65.8%] | 1.6 pts | 0.6 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 180 | 193 | 1.07× [1.05, 1.08] | +6.3% | [+5.0%, +7.7%] | 1.3 pts | 1.1 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 7.93 ms | 4.29 ms | 0.54× [0.53, 0.55] | -85.9% | [-88.8%, -82.9%] | 3.0 pts | 2.2 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 8.09 ms | 7.56 ms | 0.95× [0.93, 0.98] | -4.8% | [-7.3%, -2.3%] | 2.6 pts | 3.7 | no | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 6 | 78.7 | 76.5 | 0.98× [0.97, 0.98] | -2.3% | [-3.0%, -1.6%] | 0.7 pts | 0.7 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 6 | 79.8 | 184 | 2.31× [2.27, 2.36] | +56.7% | [+55.9%, +57.5%] | 0.8 pts | 2.2 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 6 | 2735 | 3174 | 1.16× [1.15, 1.16] | +13.4% | [+12.9%, +14.0%] | 0.6 pts | 0.8 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 2759 | 5403 | 1.95× [1.93, 1.96] | +48.7% | [+48.3%, +49.1%] | 0.4 pts | 1.1 | yes | yes |
| natural | street | 16384 | prefix | ordered | baseline | 6 | 790 | 947 | 1.19× [1.18, 1.21] | +16.2% | [+14.9%, +17.5%] | 1.2 pts | 0.8 | yes | yes |
| natural | street | 16384 | prefix | ordered | btree-sets | 6 | 803 | 2045 | 2.52× [2.47, 2.58] | +60.4% | [+59.5%, +61.2%] | 0.8 pts | 1.2 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 6 | 216 | 145 | 0.68× [0.67, 0.69] | -47.2% | [-49.5%, -44.9%] | 2.2 pts | 1.6 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 6 | 223 | 275 | 1.23× [1.22, 1.25] | +19.0% | [+17.9%, +20.2%] | 1.1 pts | 1.0 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 6 | 36.41 ms | 22.46 ms | 0.61× [0.61, 0.62] | -62.6% | [-64.7%, -60.6%] | 1.9 pts | 1.9 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 6 | 36.86 ms | 41.69 ms | 1.14× [1.12, 1.16] | +12.2% | [+10.8%, +13.6%] | 1.3 pts | 1.7 | yes | yes |
| natural | street | 212449 | valuesFor | ordered | baseline | 8 | 179 | 197 | 1.10× [1.05, 1.15] | +8.9% | [+4.4%, +13.4%] | 4.3 pts | 3.2 | no | yes |
| natural | street | 212449 | valuesFor | ordered | btree-sets | 8 | 228 | 481 | 2.08× [2.02, 2.13] | +51.8% | [+50.6%, +53.1%] | 1.3 pts | 1.5 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | baseline | 8 | 4623 | 5485 | 1.19× [1.15, 1.23] | +15.8% | [+13.1%, +18.5%] | 2.7 pts | 2.8 | no | yes |
| natural | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 5621 | 16.5 µs | 2.93× [2.89, 2.98] | +65.9% | [+65.4%, +66.4%] | 0.7 pts | 1.5 | yes | yes |
| natural | street | 212449 | prefix | ordered | baseline | 8 | 13.8 µs | 17.5 µs | 1.27× [1.23, 1.31] | +21.2% | [+18.4%, +23.9%] | 2.9 pts | 3.5 | no | yes |
| natural | street | 212449 | prefix | ordered | btree-sets | 8 | 17.8 µs | 52.7 µs | 3.02× [2.94, 3.10] | +66.9% | [+66.0%, +67.7%] | 0.8 pts | 1.0 | yes | yes |
| natural | street | 212449 | churn | ordered | baseline | 8 | 541 | 459 | 0.84× [0.81, 0.87] | -19.1% | [-23.4%, -14.7%] | 4.2 pts | 0.8 | no | yes |
| natural | street | 212449 | churn | ordered | btree-sets | 8 | 636 | 737 | 1.16× [1.14, 1.19] | +14.0% | [+12.3%, +15.8%] | 1.8 pts | 1.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 build: ordered vs btree-sets: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs baseline: the A/A validations found a systematic difference of -3.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=212449 prefix: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
