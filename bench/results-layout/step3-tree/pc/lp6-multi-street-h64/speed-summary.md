| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 6 | 52.7 | 116 | 2.21× [2.19, 2.23] | +54.7% | [+54.2%, +55.1%] | 0.4 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2324 | 2193 | 0.95× [0.94, 0.96] | -5.5% | [-6.6%, -4.4%] | 1.0 pts | 0.5 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 6 | 232 | 311 | 1.35× [1.32, 1.37] | +25.7% | [+24.3%, +27.1%] | 1.3 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 6 | 111 | 238 | 2.15× [2.11, 2.19] | +53.4% | [+52.6%, +54.2%] | 0.8 pts | 1.5 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 6 | 4.41 ms | 10.63 ms | 2.40× [2.39, 2.42] | +58.4% | [+58.1%, +58.6%] | 0.2 pts | 0.6 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 8 | 74.1 | 137 | 1.86× [1.84, 1.87] | +46.1% | [+45.6%, +46.6%] | 0.6 pts | 1.4 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 8 | 2920 | 2575 | 0.89× [0.87, 0.90] | -12.6% | [-14.5%, -10.7%] | 2.2 pts | 3.4 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 8 | 837 | 875 | 1.05× [1.04, 1.07] | +5.0% | [+3.6%, +6.3%] | 1.4 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 8 | 154 | 276 | 1.80× [1.77, 1.83] | +44.4% | [+43.5%, +45.3%] | 0.9 pts | 1.3 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 8 | 24.00 ms | 49.22 ms | 2.06× [2.03, 2.09] | +51.5% | [+50.8%, +52.2%] | 0.7 pts | 1.7 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 244 | 231 | 0.97× [0.94, 0.99] | -3.5% | [-6.2%, -0.9%] | 4.2 pts | 3.1 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 6471 | 3889 | 0.60× [0.58, 0.62] | -67.3% | [-72.3%, -62.3%] | 4.9 pts | 1.9 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 18.6 µs | 12.0 µs | 0.64× [0.62, 0.67] | -55.4% | [-60.8%, -49.9%] | 6.4 pts | 2.9 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 513 | 596 | 1.16× [1.14, 1.18] | +14.0% | [+12.7%, +15.4%] | 3.4 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=16384 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of -0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs ordered-lpage-mv: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
