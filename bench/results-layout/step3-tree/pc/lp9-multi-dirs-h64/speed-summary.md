| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mv | 6 | 85.5 | 150 | 1.75× [1.72, 1.78] | +42.9% | [+41.8%, +43.9%] | 1.0 pts | 1.3 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2733 | 3732 | 1.37× [1.34, 1.40] | +26.9% | [+25.2%, +28.7%] | 1.6 pts | 1.2 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mv | 6 | 2874 | 4243 | 1.46× [1.40, 1.52] | +31.5% | [+28.6%, +34.3%] | 2.7 pts | 1.4 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mv | 6 | 170 | 292 | 1.72× [1.71, 1.74] | +42.0% | [+41.4%, +42.6%] | 0.6 pts | 1.1 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mv | 6 | 8.05 ms | 16.78 ms | 2.08× [2.07, 2.10] | +52.0% | [+51.7%, +52.3%] | 0.3 pts | 0.7 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mv | 6 | 119 | 188 | 1.59× [1.56, 1.63] | +37.2% | [+35.8%, +38.5%] | 1.3 pts | 2.5 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mv | 6 | 3543 | 4116 | 1.16× [1.16, 1.17] | +14.2% | [+13.5%, +14.8%] | 0.6 pts | 0.9 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mv | 6 | 16.5 µs | 20.5 µs | 1.20× [1.19, 1.21] | +16.5% | [+15.9%, +17.1%] | 0.6 pts | 0.3 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mv | 6 | 236 | 357 | 1.49× [1.46, 1.53] | +33.0% | [+31.6%, +34.5%] | 1.4 pts | 1.9 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mv | 6 | 43.64 ms | 80.31 ms | 1.84× [1.83, 1.86] | +45.8% | [+45.2%, +46.3%] | 0.5 pts | 1.3 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mv | 12 | 299 | 331 | 1.11× [1.08, 1.14] | +9.8% | [+7.2%, +12.5%] | 2.8 pts | 2.1 | no | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mv | 12 | 7429 | 6302 | 0.85× [0.81, 0.90] | -17.5% | [-23.8%, -11.1%] | 6.9 pts | 4.2 | no | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mv | 12 | 139.0 µs | 124.0 µs | 0.96× [0.92, 1.00] | -4.3% | [-8.6%, +0.1%] | 5.0 pts | 0.7 | no | no |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mv | 12 | 584 | 680 | 1.17× [1.14, 1.20] | +14.3% | [+11.9%, +16.6%] | 3.0 pts | 0.6 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage-mv: the pooled interval [-8.63%, 0.08%] includes zero
