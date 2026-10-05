| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | street | 4096 | valuesFor | ordered | baseline | 6 | 54.4 | 51.3 | 0.94× [0.93, 0.96] | -6.0% | [-7.4%, -4.5%] | 1.4 pts | 1.2 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 6 | 55.0 | 135 | 2.46× [2.43, 2.49] | +59.3% | [+58.8%, +59.9%] | 0.5 pts | 1.3 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 6 | 2622 | 2286 | 0.88× [0.86, 0.89] | -14.2% | [-16.0%, -12.4%] | 1.7 pts | 0.8 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 2606 | 4563 | 1.76× [1.72, 1.80] | +43.1% | [+41.9%, +44.3%] | 1.2 pts | 1.6 | yes | yes |
| natural | street | 4096 | prefix | ordered | baseline | 6 | 257 | 224 | 0.88× [0.87, 0.89] | -13.7% | [-15.0%, -12.4%] | 1.3 pts | 0.8 | yes | yes |
| natural | street | 4096 | prefix | ordered | btree-sets | 6 | 257 | 529 | 2.06× [2.03, 2.10] | +51.5% | [+50.7%, +52.3%] | 0.7 pts | 1.4 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 6 | 105 | 94.8 | 0.91× [0.89, 0.92] | -10.4% | [-12.3%, -8.5%] | 1.8 pts | 1.9 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 6 | 105 | 188 | 1.78× [1.76, 1.81] | +44.0% | [+43.2%, +44.7%] | 0.7 pts | 1.0 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 6 | 4.17 ms | 3.71 ms | 0.89× [0.88, 0.90] | -11.9% | [-13.0%, -10.8%] | 1.0 pts | 1.3 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 6 | 4.37 ms | 7.57 ms | 1.76× [1.72, 1.80] | +43.1% | [+41.7%, +44.4%] | 1.3 pts | 1.7 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 75.3 | 71.9 | 0.95× [0.93, 0.97] | -4.9% | [-7.2%, -2.6%] | 2.2 pts | 2.3 | no | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 76.6 | 181 | 2.38× [2.34, 2.43] | +58.1% | [+57.3%, +58.8%] | 0.8 pts | 2.4 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 3150 | 2868 | 0.91× [0.90, 0.93] | -9.5% | [-10.9%, -8.1%] | 1.5 pts | 1.6 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 3159 | 5343 | 1.70× [1.68, 1.72] | +41.1% | [+40.4%, +41.8%] | 0.7 pts | 1.7 | yes | yes |
| natural | street | 16384 | prefix | ordered | baseline | 8 | 934 | 816 | 0.88× [0.86, 0.89] | -14.2% | [-15.9%, -12.5%] | 1.8 pts | 0.7 | yes | yes |
| natural | street | 16384 | prefix | ordered | btree-sets | 8 | 943 | 2037 | 2.16× [2.14, 2.17] | +53.6% | [+53.4%, +53.9%] | 0.4 pts | 0.5 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 141 | 130 | 0.92× [0.91, 0.93] | -9.1% | [-10.2%, -8.0%] | 1.1 pts | 1.1 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 146 | 264 | 1.80× [1.73, 1.87] | +44.4% | [+42.2%, +46.5%] | 2.0 pts | 3.1 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 22.14 ms | 20.08 ms | 0.91× [0.91, 0.91] | -10.0% | [-10.4%, -9.7%] | 0.8 pts | 1.1 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 22.69 ms | 41.21 ms | 1.83× [1.79, 1.87] | +45.3% | [+44.2%, +46.4%] | 1.1 pts | 1.4 | yes | yes |
| natural | street | 212449 | valuesFor | ordered | baseline | 8 | 198 | 194 | 0.97× [0.95, 0.98] | -3.5% | [-5.3%, -1.7%] | 1.7 pts | 1.1 | yes | yes |
| natural | street | 212449 | valuesFor | ordered | btree-sets | 8 | 274 | 499 | 1.82× [1.78, 1.86] | +45.0% | [+43.7%, +46.3%] | 1.5 pts | 1.9 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | baseline | 8 | 5846 | 5563 | 0.96× [0.94, 0.97] | -4.6% | [-6.2%, -3.1%] | 1.5 pts | 1.5 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 7297 | 16.6 µs | 2.28× [2.23, 2.33] | +56.1% | [+55.2%, +57.0%] | 0.9 pts | 1.6 | yes | yes |
| natural | street | 212449 | prefix | ordered | baseline | 8 | 19.3 µs | 17.9 µs | 0.93× [0.90, 0.95] | -7.7% | [-10.7%, -4.7%] | 2.9 pts | 2.2 | no | yes |
| natural | street | 212449 | prefix | ordered | btree-sets | 8 | 23.0 µs | 55.0 µs | 2.35× [2.31, 2.40] | +57.5% | [+56.7%, +58.3%] | 0.8 pts | 0.8 | yes | yes |
| natural | street | 212449 | churn | ordered | baseline | 8 | 458 | 436 | 0.95× [0.90, 1.00] | -5.6% | [-11.1%, -0.1%] | 5.6 pts | 1.2 | no | yes |
| natural | street | 212449 | churn | ordered | btree-sets | 8 | 540 | 697 | 1.29× [1.25, 1.32] | +22.2% | [+20.2%, +24.2%] | 2.0 pts | 1.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=4096 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +1.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=212449 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
