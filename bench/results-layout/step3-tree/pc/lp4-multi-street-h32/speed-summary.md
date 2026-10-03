| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 12 | 54.7 | 109 | 2.00× [1.99, 2.01] | +50.1% | [+49.8%, +50.4%] | 0.6 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 12 | 2361 | 2359 | 1.01× [0.99, 1.03] | +1.1% | [-0.9%, +3.1%] | 2.8 pts | 1.4 | no | no |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 12 | 231 | 303 | 1.32× [1.30, 1.33] | +24.1% | [+23.2%, +25.1%] | 1.2 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 12 | 123 | 269 | 2.18× [2.16, 2.20] | +54.1% | [+53.7%, +54.6%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 12 | 4.71 ms | 11.63 ms | 2.50× [2.49, 2.50] | +59.9% | [+59.8%, +60.0%] | 0.2 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 10 | 80.1 | 132 | 1.65× [1.61, 1.71] | +39.6% | [+37.7%, +41.4%] | 1.7 pts | 2.2 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 10 | 3024 | 2928 | 0.97× [0.95, 0.99] | -3.3% | [-5.3%, -1.3%] | 1.9 pts | 2.1 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 10 | 890 | 990 | 1.11× [1.10, 1.12] | +9.9% | [+8.9%, +11.0%] | 1.1 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 10 | 209 | 342 | 1.66× [1.65, 1.68] | +39.8% | [+39.3%, +40.4%] | 0.7 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 10 | 26.94 ms | 56.46 ms | 2.10× [2.05, 2.14] | +52.3% | [+51.3%, +53.3%] | 1.2 pts | 2.4 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 334 | 318 | 0.95× [0.93, 0.97] | -5.1% | [-7.0%, -3.3%] | 2.1 pts | 1.2 | yes | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 8316 | 5230 | 0.64× [0.63, 0.66] | -55.4% | [-58.6%, -52.1%] | 4.8 pts | 3.9 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 27.5 µs | 18.0 µs | 0.65× [0.64, 0.65] | -54.9% | [-56.1%, -53.8%] | 3.6 pts | 1.3 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 665 | 809 | 1.21× [1.19, 1.23] | +17.3% | [+15.8%, +18.7%] | 1.7 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage-mv: the pooled interval [-0.87%, 3.14%] includes zero
- multi-str street n=4096 build: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of +0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 prefix: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of -0.68% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 build: ordered vs ordered-lpage-mv: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
