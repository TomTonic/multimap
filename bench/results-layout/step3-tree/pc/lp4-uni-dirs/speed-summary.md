| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage | 12 | 75.2 | 118 | 1.59× [1.56, 1.62] | +37.1% | [+36.1%, +38.1%] | 1.4 pts | 2.6 | yes | yes |
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-zc | 12 | 74.7 | 101 | 1.35× [1.34, 1.36] | +26.0% | [+25.6%, +26.5%] | 0.9 pts | 1.4 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage | 12 | 2021 | 1815 | 0.89× [0.88, 0.91] | -12.0% | [-14.0%, -10.0%] | 2.6 pts | 1.4 | no | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-zc | 12 | 2020 | 1178 | 0.58× [0.58, 0.59] | -71.0% | [-72.9%, -69.0%] | 2.3 pts | 1.0 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage | 12 | 2227 | 2184 | 0.98× [0.97, 0.98] | -2.5% | [-3.3%, -1.7%] | 1.3 pts | 1.2 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage-zc | 12 | 2233 | 1360 | 0.61× [0.61, 0.62] | -62.6% | [-64.4%, -60.8%] | 4.9 pts | 1.6 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage | 12 | 198 | 254 | 1.28× [1.27, 1.30] | +22.1% | [+21.2%, +23.0%] | 1.0 pts | 1.2 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage-zc | 12 | 201 | 325 | 1.61× [1.58, 1.63] | +37.7% | [+36.6%, +38.8%] | 1.4 pts | 1.4 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage | 12 | 2.80 ms | 4.73 ms | 1.68× [1.67, 1.70] | +40.6% | [+40.0%, +41.3%] | 0.9 pts | 1.3 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage-zc | 12 | 2.79 ms | 5.56 ms | 1.99× [1.96, 2.01] | +49.7% | [+49.0%, +50.3%] | 0.7 pts | 1.2 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage | 12 | 109 | 150 | 1.38× [1.35, 1.41] | +27.5% | [+25.7%, +29.3%] | 2.0 pts | 2.7 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-zc | 12 | 107 | 131 | 1.20× [1.18, 1.22] | +16.8% | [+15.3%, +18.3%] | 1.9 pts | 2.1 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage | 12 | 2607 | 2292 | 0.88× [0.87, 0.88] | -14.1% | [-15.3%, -13.0%] | 1.4 pts | 1.3 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-zc | 12 | 2598 | 1416 | 0.54× [0.54, 0.55] | -83.9% | [-85.6%, -82.2%] | 2.1 pts | 1.0 | yes | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage | 12 | 10.7 µs | 9191 | 0.87× [0.84, 0.90] | -15.1% | [-19.2%, -11.0%] | 4.4 pts | 1.2 | no | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage-zc | 12 | 10.2 µs | 4974 | 0.50× [0.48, 0.52] | -100.9% | [-108.2%, -93.6%] | 7.5 pts | 0.7 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage | 12 | 312 | 364 | 1.17× [1.15, 1.19] | +14.7% | [+13.3%, +16.1%] | 1.6 pts | 1.6 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage-zc | 12 | 331 | 458 | 1.41× [1.37, 1.45] | +28.9% | [+26.8%, +30.9%] | 2.3 pts | 1.9 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage | 12 | 15.51 ms | 23.86 ms | 1.53× [1.52, 1.56] | +34.8% | [+34.0%, +35.7%] | 1.2 pts | 1.3 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage-zc | 12 | 15.77 ms | 30.08 ms | 1.89× [1.87, 1.90] | +47.0% | [+46.5%, +47.5%] | 0.7 pts | 0.8 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage | 12 | 270 | 255 | 0.96× [0.94, 0.98] | -4.2% | [-5.8%, -2.5%] | 2.7 pts | 1.4 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-zc | 12 | 255 | 210 | 0.84× [0.82, 0.85] | -19.2% | [-21.2%, -17.1%] | 3.3 pts | 1.4 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage | 12 | 4315 | 3274 | 0.77× [0.74, 0.80] | -30.5% | [-35.3%, -25.6%] | 6.0 pts | 2.5 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-zc | 12 | 3794 | 2009 | 0.53× [0.51, 0.54] | -90.2% | [-95.4%, -85.0%] | 7.6 pts | 1.7 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage | 12 | 97.0 µs | 65.5 µs | 0.74× [0.69, 0.82] | -34.3% | [-45.9%, -22.7%] | 12.0 pts | 0.9 | no | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage-zc | 12 | 85.1 µs | 35.0 µs | 0.43× [0.42, 0.44] | -132.4% | [-135.8%, -128.9%] | 9.4 pts | 0.8 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage | 12 | 599 | 659 | 1.10× [1.09, 1.12] | +9.4% | [+8.5%, +10.4%] | 0.9 pts | 0.8 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage-zc | 12 | 607 | 750 | 1.24× [1.22, 1.26] | +19.2% | [+17.9%, +20.6%] | 1.7 pts | 1.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str dirs n=4096 valuesFor: ordered vs ordered-lpage: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.89% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 churn: ordered vs ordered-lpage: the A/A validations found a systematic difference of +0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
