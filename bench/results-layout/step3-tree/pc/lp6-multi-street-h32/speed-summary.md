| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 8 | 53.2 | 106 | 2.00× [1.96, 2.03] | +49.9% | [+49.1%, +50.7%] | 0.9 pts | 2.0 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 8 | 2371 | 2392 | 1.01× [0.99, 1.03] | +1.0% | [-0.7%, +2.8%] | 2.0 pts | 0.8 | yes | no |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 8 | 234 | 303 | 1.30× [1.28, 1.31] | +23.0% | [+22.0%, +23.9%] | 1.0 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 8 | 112 | 250 | 2.22× [2.20, 2.24] | +55.0% | [+54.6%, +55.3%] | 0.5 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 8 | 4.50 ms | 11.02 ms | 2.46× [2.44, 2.48] | +59.4% | [+59.1%, +59.6%] | 0.5 pts | 1.7 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 6 | 74.4 | 125 | 1.70× [1.67, 1.73] | +41.2% | [+40.3%, +42.1%] | 0.9 pts | 1.6 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2932 | 2901 | 0.99× [0.98, 1.00] | -1.3% | [-2.4%, -0.3%] | 1.0 pts | 1.4 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 6 | 842 | 957 | 1.14× [1.12, 1.16] | +12.2% | [+11.0%, +13.4%] | 1.2 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 6 | 159 | 296 | 1.86× [1.80, 1.91] | +46.1% | [+44.5%, +47.7%] | 1.5 pts | 2.5 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 6 | 23.83 ms | 50.31 ms | 2.11× [2.10, 2.12] | +52.6% | [+52.3%, +52.9%] | 0.3 pts | 0.6 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 235 | 235 | 0.99× [0.95, 1.03] | -1.1% | [-5.3%, +3.2%] | 5.0 pts | 3.7 | no | no |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 6808 | 4903 | 0.74× [0.71, 0.77] | -35.3% | [-40.2%, -30.4%] | 7.1 pts | 3.4 | no | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 19.2 µs | 14.7 µs | 0.78× [0.77, 0.79] | -28.0% | [-29.2%, -26.9%] | 2.1 pts | 1.4 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 522 | 644 | 1.22× [1.20, 1.25] | +18.3% | [+16.9%, +19.8%] | 2.3 pts | 0.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage-mv: the pooled interval [-0.75%, 2.84%] includes zero
- multi-str street n=16384 churn: ordered vs ordered-lpage-mv: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mv: the pooled interval [-5.31%, 3.19%] includes zero
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mv: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
