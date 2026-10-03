| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mv | 6 | 86.9 | 152 | 1.77× [1.71, 1.83] | +43.4% | [+41.5%, +45.3%] | 1.8 pts | 3.0 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2850 | 3934 | 1.39× [1.37, 1.42] | +28.2% | [+27.0%, +29.3%] | 1.1 pts | 0.9 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mv | 6 | 3079 | 4627 | 1.51× [1.48, 1.54] | +33.6% | [+32.3%, +34.9%] | 1.3 pts | 0.6 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mv | 6 | 185 | 304 | 1.64× [1.62, 1.66] | +39.0% | [+38.2%, +39.7%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mv | 6 | 8.70 ms | 17.37 ms | 2.00× [1.98, 2.01] | +49.9% | [+49.6%, +50.2%] | 0.3 pts | 0.6 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mv | 8 | 137 | 204 | 1.46× [1.40, 1.53] | +31.6% | [+28.8%, +34.5%] | 2.9 pts | 3.1 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mv | 8 | 3748 | 4515 | 1.20× [1.19, 1.21] | +16.6% | [+15.8%, +17.4%] | 1.3 pts | 1.6 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mv | 8 | 17.5 µs | 22.7 µs | 1.23× [1.21, 1.25] | +18.8% | [+17.3%, +20.3%] | 1.4 pts | 0.7 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mv | 8 | 315 | 427 | 1.34× [1.32, 1.36] | +25.4% | [+24.3%, +26.5%] | 1.1 pts | 1.5 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mv | 8 | 51.21 ms | 89.69 ms | 1.75× [1.72, 1.78] | +42.8% | [+41.7%, +43.8%] | 1.1 pts | 1.3 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mv | 10 | 320 | 344 | 1.08× [1.06, 1.11] | +7.7% | [+5.9%, +9.5%] | 2.4 pts | 1.2 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mv | 10 | 8418 | 7194 | 0.84× [0.84, 0.85] | -18.4% | [-19.1%, -17.7%] | 2.2 pts | 1.6 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mv | 10 | 138.7 µs | 138.1 µs | 0.99× [0.97, 1.01] | -1.1% | [-3.0%, +0.9%] | 3.9 pts | 0.5 | yes | no |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mv | 10 | 633 | 769 | 1.21× [1.18, 1.24] | +17.2% | [+15.4%, +19.1%] | 1.9 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str dirs n=4096 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage-mv: the pooled difference of -1.09% does not clear the 2.58% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage-mv: the pooled interval [-3.04%, 0.86%] includes zero
