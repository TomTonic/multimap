| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 8 | 52.5 | 104 | 1.99× [1.96, 2.01] | +49.7% | [+49.1%, +50.3%] | 0.6 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 8 | 2335 | 2364 | 1.03× [1.01, 1.05] | +2.6% | [+0.7%, +4.5%] | 2.4 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 8 | 232 | 302 | 1.30× [1.28, 1.31] | +22.9% | [+21.8%, +23.9%] | 1.4 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 8 | 113 | 246 | 2.20× [2.17, 2.22] | +54.5% | [+54.0%, +55.0%] | 0.5 pts | 0.7 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 8 | 4.42 ms | 10.80 ms | 2.42× [2.41, 2.43] | +58.7% | [+58.5%, +58.9%] | 0.3 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 6 | 79.8 | 132 | 1.63× [1.58, 1.69] | +38.8% | [+36.7%, +40.9%] | 2.0 pts | 2.6 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 6 | 3055 | 3037 | 0.99× [0.98, 1.00] | -0.8% | [-2.0%, +0.4%] | 1.1 pts | 1.1 | yes | no |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 6 | 926 | 1027 | 1.13× [1.11, 1.15] | +11.3% | [+9.8%, +12.8%] | 1.5 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 6 | 205 | 329 | 1.58× [1.53, 1.64] | +36.8% | [+34.6%, +39.0%] | 2.1 pts | 1.9 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 6 | 26.18 ms | 53.59 ms | 2.05× [2.03, 2.07] | +51.2% | [+50.7%, +51.7%] | 0.5 pts | 1.1 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 331 | 315 | 0.94× [0.92, 0.96] | -6.0% | [-8.2%, -3.8%] | 2.5 pts | 1.3 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 8286 | 5524 | 0.67× [0.64, 0.69] | -50.2% | [-55.5%, -44.9%] | 6.3 pts | 4.4 | no | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 27.0 µs | 18.0 µs | 0.66× [0.64, 0.68] | -50.9% | [-55.1%, -46.7%] | 4.7 pts | 1.8 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 661 | 797 | 1.19× [1.17, 1.22] | +16.1% | [+14.2%, +17.9%] | 1.9 pts | 0.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=16384 valuesFor: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of +0.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs ordered-lpage-mv: the pooled interval [-2.02%, 0.37%] includes zero
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
