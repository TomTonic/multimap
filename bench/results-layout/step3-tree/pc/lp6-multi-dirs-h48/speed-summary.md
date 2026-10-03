| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mv | 6 | 86.8 | 155 | 1.79× [1.76, 1.82] | +44.3% | [+43.3%, +45.2%] | 0.9 pts | 1.5 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2832 | 3858 | 1.35× [1.33, 1.37] | +26.0% | [+25.0%, +27.0%] | 0.9 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mv | 6 | 3026 | 4554 | 1.49× [1.46, 1.53] | +33.1% | [+31.4%, +34.8%] | 1.6 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mv | 6 | 183 | 305 | 1.66× [1.64, 1.68] | +39.7% | [+39.0%, +40.5%] | 0.7 pts | 0.9 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mv | 6 | 8.72 ms | 17.60 ms | 2.04× [2.02, 2.06] | +50.9% | [+50.4%, +51.4%] | 0.5 pts | 1.2 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mv | 6 | 138 | 202 | 1.48× [1.44, 1.52] | +32.4% | [+30.5%, +34.2%] | 1.8 pts | 1.9 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mv | 6 | 3747 | 4367 | 1.17× [1.15, 1.19] | +14.5% | [+13.3%, +15.8%] | 1.2 pts | 1.3 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mv | 6 | 17.8 µs | 22.4 µs | 1.18× [1.16, 1.20] | +15.4% | [+14.0%, +16.8%] | 1.3 pts | 0.5 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mv | 6 | 312 | 414 | 1.33× [1.31, 1.35] | +24.8% | [+23.6%, +26.1%] | 1.2 pts | 1.7 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mv | 6 | 50.41 ms | 89.60 ms | 1.78× [1.74, 1.82] | +43.9% | [+42.6%, +45.1%] | 1.2 pts | 2.0 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mv | 12 | 329 | 351 | 1.09× [1.07, 1.12] | +8.4% | [+6.2%, +10.6%] | 2.4 pts | 1.3 | no | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mv | 12 | 8386 | 6740 | 0.81× [0.79, 0.83] | -23.6% | [-26.6%, -20.5%] | 3.8 pts | 1.8 | no | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mv | 12 | 146.3 µs | 128.7 µs | 0.95× [0.91, 1.00] | -4.9% | [-9.3%, -0.5%] | 4.9 pts | 0.6 | no | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mv | 12 | 623 | 733 | 1.18× [1.16, 1.21] | +15.6% | [+13.8%, +17.4%] | 2.2 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).
