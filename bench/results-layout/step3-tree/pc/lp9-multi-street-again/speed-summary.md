| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 10 | 54.6 | 111 | 2.05× [2.03, 2.08] | +51.3% | [+50.7%, +51.9%] | 0.6 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 10 | 2327 | 2209 | 0.95× [0.94, 0.97] | -4.9% | [-6.7%, -3.1%] | 2.0 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 10 | 226 | 299 | 1.32× [1.31, 1.34] | +24.3% | [+23.5%, +25.2%] | 1.3 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 10 | 118 | 235 | 2.02× [2.00, 2.04] | +50.4% | [+50.0%, +50.9%] | 0.6 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 10 | 4.55 ms | 10.52 ms | 2.31× [2.30, 2.33] | +56.8% | [+56.5%, +57.1%] | 0.4 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 10 | 74.3 | 131 | 1.76× [1.74, 1.79] | +43.3% | [+42.6%, +44.0%] | 0.7 pts | 1.4 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 10 | 2897 | 2627 | 0.90× [0.89, 0.92] | -11.0% | [-12.9%, -9.1%] | 1.9 pts | 2.0 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 10 | 833 | 903 | 1.08× [1.06, 1.10] | +7.4% | [+5.8%, +9.0%] | 1.9 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 10 | 155 | 271 | 1.74× [1.70, 1.77] | +42.4% | [+41.2%, +43.6%] | 1.3 pts | 2.0 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 10 | 24.09 ms | 47.60 ms | 1.98× [1.97, 2.00] | +49.6% | [+49.3%, +50.0%] | 0.4 pts | 0.9 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 238 | 222 | 0.95× [0.92, 0.98] | -5.1% | [-8.3%, -1.9%] | 3.2 pts | 2.4 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 6453 | 4147 | 0.65× [0.63, 0.66] | -54.7% | [-58.0%, -51.3%] | 5.7 pts | 2.7 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 18.2 µs | 12.3 µs | 0.68× [0.66, 0.70] | -47.2% | [-52.3%, -42.0%] | 5.5 pts | 2.5 | no | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 525 | 603 | 1.15× [1.14, 1.17] | +13.3% | [+12.1%, +14.5%] | 3.0 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of -0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 churn: ordered vs ordered-lpage-mv: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs ordered-lpage-mv: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
