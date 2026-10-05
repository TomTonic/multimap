| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | street | 4096 | valuesFor | ordered | baseline | 6 | 85.4 | 82.4 | 0.96× [0.94, 0.97] | -4.7% | [-6.5%, -2.9%] | 1.7 pts | 1.6 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 6 | 86.8 | 139 | 1.61× [1.59, 1.64] | +37.9% | [+37.0%, +38.8%] | 0.9 pts | 1.8 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 6 | 4757 | 5124 | 1.08× [1.07, 1.09] | +7.4% | [+6.3%, +8.4%] | 1.0 pts | 0.9 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 4744 | 7181 | 1.51× [1.50, 1.53] | +33.9% | [+33.3%, +34.4%] | 0.5 pts | 0.7 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 6 | 443 | 472 | 1.07× [1.05, 1.08] | +6.2% | [+5.0%, +7.4%] | 1.1 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 6 | 443 | 727 | 1.62× [1.61, 1.64] | +38.4% | [+37.7%, +39.1%] | 0.7 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 6 | 206 | 125 | 0.61× [0.60, 0.61] | -65.2% | [-67.8%, -62.6%] | 2.5 pts | 1.2 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 6 | 207 | 201 | 0.97× [0.96, 0.97] | -3.6% | [-4.1%, -3.0%] | 0.6 pts | 0.5 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 6 | 8.80 ms | 4.69 ms | 0.54× [0.53, 0.54] | -86.9% | [-88.1%, -85.7%] | 1.1 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 6 | 8.84 ms | 7.84 ms | 0.89× [0.88, 0.91] | -11.9% | [-13.7%, -10.0%] | 1.8 pts | 1.3 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 8 | 106 | 105 | 0.99× [0.98, 1.00] | -1.2% | [-2.1%, -0.3%] | 1.0 pts | 1.1 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 8 | 109 | 188 | 1.69× [1.66, 1.73] | +41.0% | [+39.9%, +42.1%] | 1.5 pts | 3.3 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 5416 | 5794 | 1.07× [1.06, 1.08] | +6.6% | [+5.9%, +7.3%] | 0.7 pts | 1.4 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 5427 | 8341 | 1.51× [1.48, 1.55] | +34.0% | [+32.5%, +35.4%] | 1.4 pts | 3.2 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 8 | 1743 | 1931 | 1.11× [1.07, 1.14] | +9.5% | [+7.0%, +12.1%] | 2.4 pts | 1.3 | no | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 8 | 1781 | 3088 | 1.72× [1.68, 1.75] | +41.8% | [+40.6%, +42.9%] | 1.4 pts | 1.4 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 8 | 257 | 177 | 0.69× [0.68, 0.70] | -45.1% | [-47.1%, -43.0%] | 1.9 pts | 1.1 | yes | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 8 | 256 | 292 | 1.13× [1.12, 1.14] | +11.7% | [+11.1%, +12.2%] | 0.6 pts | 0.6 | yes | yes |
| natural-str | street | 16384 | build | ordered | baseline | 8 | 42.36 ms | 25.66 ms | 0.61× [0.60, 0.61] | -65.2% | [-66.9%, -63.4%] | 1.8 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 8 | 42.03 ms | 43.94 ms | 1.05× [1.04, 1.06] | +5.0% | [+4.2%, +5.9%] | 0.8 pts | 1.3 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | baseline | 8 | 231 | 250 | 1.07× [1.04, 1.11] | +6.9% | [+3.9%, +9.9%] | 2.9 pts | 2.5 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | btree-sets | 8 | 277 | 495 | 1.82× [1.77, 1.86] | +44.9% | [+43.6%, +46.2%] | 1.6 pts | 2.5 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 8417 | 9539 | 1.15× [1.12, 1.18] | +12.9% | [+10.5%, +15.4%] | 2.3 pts | 3.2 | no | yes |
| natural-str | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 9242 | 23.3 µs | 2.52× [2.45, 2.59] | +60.3% | [+59.2%, +61.4%] | 1.7 pts | 3.9 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | baseline | 8 | 31.1 µs | 35.5 µs | 1.16× [1.14, 1.19] | +13.8% | [+12.0%, +15.7%] | 1.9 pts | 2.0 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | btree-sets | 8 | 32.1 µs | 77.2 µs | 2.45× [2.41, 2.50] | +59.3% | [+58.5%, +60.0%] | 0.7 pts | 0.8 | yes | yes |
| natural-str | street | 212449 | churn | ordered | baseline | 8 | 636 | 555 | 0.87× [0.83, 0.90] | -15.4% | [-20.3%, -10.6%] | 4.9 pts | 0.9 | no | yes |
| natural-str | street | 212449 | churn | ordered | btree-sets | 8 | 734 | 786 | 1.08× [1.07, 1.09] | +7.6% | [+6.7%, +8.6%] | 0.9 pts | 0.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +1.48% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
