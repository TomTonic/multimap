| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 6 | 55.1 | 102 | 1.86× [1.83, 1.89] | +46.2% | [+45.2%, +47.1%] | 0.9 pts | 1.8 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2377 | 2904 | 1.22× [1.20, 1.24] | +18.0% | [+16.9%, +19.2%] | 1.1 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 6 | 232 | 324 | 1.40× [1.39, 1.41] | +28.5% | [+28.0%, +29.0%] | 0.5 pts | 0.6 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 6 | 126 | 290 | 2.32× [2.27, 2.37] | +56.9% | [+56.0%, +57.8%] | 0.9 pts | 0.8 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 6 | 4.70 ms | 12.59 ms | 2.68× [2.66, 2.70] | +62.7% | [+62.4%, +63.0%] | 0.3 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 6 | 79.4 | 128 | 1.60× [1.54, 1.66] | +37.5% | [+35.1%, +39.9%] | 2.3 pts | 3.1 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 6 | 3055 | 3506 | 1.15× [1.13, 1.16] | +12.8% | [+11.8%, +13.7%] | 0.9 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 6 | 947 | 1208 | 1.28× [1.25, 1.30] | +21.8% | [+20.3%, +23.2%] | 1.4 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 6 | 220 | 381 | 1.72× [1.66, 1.80] | +42.0% | [+39.6%, +44.4%] | 2.3 pts | 1.8 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 6 | 27.30 ms | 61.90 ms | 2.26× [2.23, 2.29] | +55.7% | [+55.1%, +56.3%] | 0.6 pts | 0.8 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 338 | 337 | 1.02× [1.00, 1.03] | +1.6% | [+0.1%, +3.1%] | 2.3 pts | 1.3 | yes | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 8299 | 6849 | 0.81× [0.79, 0.82] | -24.1% | [-25.9%, -22.3%] | 4.4 pts | 4.7 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 29.2 µs | 24.5 µs | 0.83× [0.80, 0.85] | -21.0% | [-24.8%, -17.2%] | 5.0 pts | 1.8 | no | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 674 | 876 | 1.29× [1.27, 1.30] | +22.2% | [+21.3%, +23.1%] | 1.5 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 prefix: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of -0.57% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
