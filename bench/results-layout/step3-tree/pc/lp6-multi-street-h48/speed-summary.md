| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 6 | 52.9 | 110 | 2.09× [2.07, 2.11] | +52.2% | [+51.7%, +52.7%] | 0.5 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2328 | 2203 | 0.96× [0.95, 0.97] | -4.0% | [-5.3%, -2.7%] | 1.2 pts | 0.6 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 6 | 232 | 302 | 1.31× [1.28, 1.33] | +23.5% | [+21.9%, +25.0%] | 1.5 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 6 | 111 | 232 | 2.10× [2.07, 2.13] | +52.4% | [+51.7%, +53.1%] | 0.7 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 6 | 4.45 ms | 10.47 ms | 2.36× [2.34, 2.38] | +57.6% | [+57.2%, +58.0%] | 0.4 pts | 1.6 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 6 | 76.6 | 135 | 1.77× [1.74, 1.79] | +43.4% | [+42.6%, +44.2%] | 0.8 pts | 1.8 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2978 | 2721 | 0.92× [0.91, 0.92] | -9.3% | [-10.2%, -8.3%] | 0.9 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 6 | 862 | 921 | 1.06× [1.05, 1.07] | +5.8% | [+5.0%, +6.5%] | 0.7 pts | 0.5 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 6 | 159 | 281 | 1.77× [1.76, 1.79] | +43.7% | [+43.2%, +44.1%] | 0.4 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 6 | 23.87 ms | 48.13 ms | 2.01× [1.99, 2.02] | +50.1% | [+49.7%, +50.6%] | 0.4 pts | 1.0 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 256 | 238 | 0.95× [0.93, 0.97] | -5.2% | [-7.7%, -2.7%] | 3.2 pts | 2.6 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 6505 | 4259 | 0.65× [0.61, 0.69] | -54.0% | [-63.6%, -44.3%] | 9.1 pts | 5.6 | no | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 18.9 µs | 12.7 µs | 0.68× [0.66, 0.70] | -48.1% | [-52.5%, -43.7%] | 5.7 pts | 2.5 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 523 | 604 | 1.18× [1.16, 1.19] | +15.0% | [+13.9%, +16.1%] | 2.6 pts | 0.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs ordered-lpage-mv: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
