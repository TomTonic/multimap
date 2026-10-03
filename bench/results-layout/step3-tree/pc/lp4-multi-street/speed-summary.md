| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 6 | 53.5 | 104 | 1.94× [1.92, 1.97] | +48.6% | [+47.9%, +49.3%] | 0.7 pts | 1.7 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mvzc | 6 | 53.7 | 84.8 | 1.58× [1.56, 1.61] | +36.7% | [+35.7%, +37.7%] | 0.9 pts | 1.9 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2347 | 2456 | 1.05× [1.03, 1.07] | +4.7% | [+3.3%, +6.2%] | 1.4 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mvzc | 6 | 2322 | 1767 | 0.76× [0.75, 0.78] | -31.3% | [-33.8%, -28.8%] | 2.4 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 6 | 227 | 296 | 1.31× [1.28, 1.33] | +23.5% | [+22.0%, +24.9%] | 1.4 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mvzc | 6 | 228 | 227 | 0.99× [0.98, 1.01] | -0.8% | [-2.2%, +0.7%] | 1.4 pts | 1.0 | yes | no |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 6 | 116 | 271 | 2.35× [2.31, 2.40] | +57.5% | [+56.7%, +58.3%] | 0.7 pts | 1.6 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mvzc | 6 | 116 | 324 | 2.81× [2.78, 2.85] | +64.5% | [+64.1%, +64.9%] | 0.4 pts | 0.7 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 6 | 4.50 ms | 11.58 ms | 2.58× [2.56, 2.59] | +61.2% | [+60.9%, +61.4%] | 0.3 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mvzc | 6 | 4.49 ms | 13.39 ms | 2.97× [2.93, 3.02] | +66.4% | [+65.8%, +66.9%] | 0.5 pts | 1.5 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 12 | 74.2 | 122 | 1.65× [1.63, 1.67] | +39.4% | [+38.6%, +40.3%] | 1.1 pts | 2.2 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mvzc | 12 | 74.3 | 102 | 1.38× [1.38, 1.39] | +27.7% | [+27.3%, +28.0%] | 0.7 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 12 | 2930 | 3024 | 1.03× [1.01, 1.06] | +3.4% | [+1.3%, +5.4%] | 2.3 pts | 3.1 | no | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mvzc | 12 | 2943 | 2085 | 0.71× [0.70, 0.71] | -41.4% | [-42.1%, -40.7%] | 1.0 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 12 | 842 | 979 | 1.17× [1.16, 1.19] | +14.8% | [+13.9%, +15.7%] | 1.5 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mvzc | 12 | 834 | 660 | 0.79× [0.79, 0.80] | -26.2% | [-27.4%, -25.0%] | 1.6 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 12 | 162 | 318 | 1.94× [1.89, 1.98] | +48.4% | [+47.2%, +49.6%] | 1.6 pts | 2.7 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mvzc | 12 | 166 | 375 | 2.25× [2.21, 2.29] | +55.6% | [+54.8%, +56.3%] | 0.9 pts | 1.3 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 12 | 24.20 ms | 54.04 ms | 2.24× [2.23, 2.25] | +55.3% | [+55.2%, +55.5%] | 0.4 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mvzc | 12 | 24.10 ms | 63.11 ms | 2.60× [2.57, 2.63] | +61.6% | [+61.1%, +62.0%] | 0.5 pts | 1.0 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 249 | 246 | 1.01× [1.00, 1.02] | +1.0% | [+0.0%, +1.9%] | 1.6 pts | 1.2 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mvzc | 12 | 226 | 217 | 0.93× [0.91, 0.95] | -7.5% | [-10.1%, -4.9%] | 3.4 pts | 2.2 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 6754 | 5381 | 0.79× [0.76, 0.83] | -26.3% | [-32.1%, -20.4%] | 6.8 pts | 4.5 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mvzc | 12 | 6308 | 3875 | 0.61× [0.58, 0.64] | -64.4% | [-71.8%, -57.0%] | 7.8 pts | 3.3 | no | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 19.9 µs | 16.2 µs | 0.81× [0.79, 0.84] | -23.0% | [-26.6%, -19.4%] | 4.0 pts | 2.4 | no | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mvzc | 12 | 18.5 µs | 10.9 µs | 0.60× [0.59, 0.62] | -65.5% | [-69.0%, -61.9%] | 5.2 pts | 2.6 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 540 | 690 | 1.28× [1.24, 1.31] | +21.6% | [+19.5%, +23.7%] | 2.9 pts | 1.1 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mvzc | 12 | 551 | 746 | 1.36× [1.32, 1.41] | +26.6% | [+24.1%, +29.0%] | 2.6 pts | 1.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 prefix: ordered vs ordered-lpage-mvzc: the pooled difference of -0.75% does not clear the 1.90% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 prefix: ordered vs ordered-lpage-mvzc: the pooled interval [-2.25%, 0.75%] includes zero
- multi-str street n=4096 churn: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of +0.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 churn: ordered vs ordered-lpage-mv: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mv: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mvzc: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mvzc: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs ordered-lpage-mv: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs ordered-lpage-mvzc: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
