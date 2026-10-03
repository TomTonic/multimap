| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mv | 6 | 83.9 | 146 | 1.77× [1.70, 1.84] | +43.4% | [+41.2%, +45.6%] | 2.1 pts | 3.1 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mvzc | 6 | 83.3 | 119 | 1.43× [1.41, 1.45] | +29.9% | [+28.9%, +31.0%] | 1.0 pts | 1.2 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2732 | 3857 | 1.43× [1.40, 1.46] | +30.1% | [+28.7%, +31.5%] | 1.3 pts | 0.9 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mvzc | 6 | 2743 | 2413 | 0.87× [0.86, 0.88] | -14.8% | [-16.1%, -13.6%] | 1.1 pts | 0.5 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mv | 6 | 2882 | 4441 | 1.53× [1.50, 1.57] | +34.7% | [+33.3%, +36.2%] | 1.4 pts | 0.6 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mvzc | 6 | 2876 | 2672 | 0.92× [0.91, 0.93] | -8.7% | [-9.3%, -8.0%] | 0.6 pts | 0.4 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mv | 6 | 170 | 303 | 1.79× [1.77, 1.81] | +44.2% | [+43.6%, +44.8%] | 0.6 pts | 1.1 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mvzc | 6 | 173 | 376 | 2.18× [2.14, 2.22] | +54.1% | [+53.2%, +54.9%] | 0.8 pts | 1.4 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mv | 6 | 8.06 ms | 17.14 ms | 2.11× [2.08, 2.14] | +52.7% | [+52.0%, +53.4%] | 0.7 pts | 1.4 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mvzc | 6 | 8.17 ms | 20.14 ms | 2.45× [2.42, 2.48] | +59.2% | [+58.7%, +59.7%] | 0.5 pts | 0.9 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mv | 12 | 116 | 186 | 1.58× [1.56, 1.61] | +36.7% | [+35.8%, +37.7%] | 2.0 pts | 3.6 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mvzc | 12 | 116 | 151 | 1.30× [1.29, 1.31] | +23.3% | [+22.7%, +23.9%] | 0.7 pts | 0.9 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mv | 12 | 3538 | 4384 | 1.24× [1.23, 1.24] | +19.1% | [+18.6%, +19.6%] | 0.7 pts | 0.8 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mvzc | 12 | 3554 | 2795 | 0.78× [0.78, 0.79] | -27.5% | [-28.5%, -26.5%] | 1.1 pts | 0.9 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mv | 12 | 16.9 µs | 22.3 µs | 1.30× [1.29, 1.31] | +22.9% | [+22.2%, +23.5%] | 1.1 pts | 0.7 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mvzc | 12 | 15.5 µs | 11.3 µs | 0.73× [0.71, 0.76] | -36.3% | [-41.4%, -31.2%] | 6.3 pts | 0.7 | no | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mv | 12 | 243 | 382 | 1.57× [1.56, 1.59] | +36.5% | [+35.8%, +37.2%] | 0.8 pts | 1.3 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mvzc | 12 | 260 | 471 | 1.80× [1.77, 1.83] | +44.6% | [+43.7%, +45.5%] | 1.1 pts | 1.7 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mv | 12 | 44.41 ms | 83.54 ms | 1.87× [1.86, 1.89] | +46.6% | [+46.3%, +47.0%] | 0.5 pts | 1.3 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mvzc | 12 | 43.96 ms | 101.22 ms | 2.30× [2.28, 2.32] | +56.5% | [+56.1%, +56.8%] | 0.4 pts | 0.7 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mv | 12 | 232 | 280 | 1.19× [1.17, 1.22] | +16.2% | [+14.4%, +18.0%] | 3.3 pts | 2.9 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mvzc | 12 | 227 | 234 | 1.06× [1.04, 1.09] | +6.0% | [+4.0%, +7.9%] | 3.7 pts | 2.8 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mv | 12 | 6165 | 6460 | 1.04× [1.01, 1.07] | +3.7% | [+0.5%, +6.9%] | 4.1 pts | 2.8 | no | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mvzc | 12 | 5558 | 4282 | 0.77× [0.74, 0.79] | -30.3% | [-34.5%, -26.1%] | 5.3 pts | 3.0 | no | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mv | 12 | 133.7 µs | 135.1 µs | 1.07× [1.03, 1.10] | +6.2% | [+3.1%, +9.4%] | 3.4 pts | 0.6 | no | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mvzc | 12 | 123.0 µs | 91.2 µs | 0.73× [0.72, 0.74] | -37.4% | [-39.5%, -35.3%] | 3.9 pts | 0.8 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mv | 12 | 543 | 700 | 1.28× [1.27, 1.29] | +22.1% | [+21.4%, +22.7%] | 2.0 pts | 1.0 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mvzc | 12 | 528 | 755 | 1.42× [1.41, 1.44] | +29.8% | [+28.8%, +30.7%] | 1.9 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str dirs n=4096 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of -0.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage-mvzc: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of -0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage-mvzc: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of -1.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
