| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 87.9 | 79.2 | 0.90× [0.89, 0.91] | -11.4% | [-13.0%, -9.8%] | 1.6 pts | 2.0 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 88.5 | 109 | 1.24× [1.22, 1.25] | +19.0% | [+18.0%, +20.1%] | 1.0 pts | 1.2 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1436 | 2269 | 1.58× [1.56, 1.60] | +36.7% | [+36.0%, +37.5%] | 1.0 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1400 | 563 | 0.40× [0.40, 0.41] | -147.4% | [-152.4%, -142.4%] | 4.9 pts | 1.6 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 1627 | 2582 | 1.56× [1.54, 1.59] | +36.1% | [+35.2%, +36.9%] | 1.2 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | btree-map | 8 | 1642 | 668 | 0.41× [0.40, 0.41] | -144.7% | [-148.0%, -141.4%] | 4.0 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 272 | 210 | 0.77× [0.76, 0.78] | -29.8% | [-32.0%, -27.6%] | 2.9 pts | 2.4 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | btree-map | 8 | 275 | 187 | 0.68× [0.67, 0.69] | -47.6% | [-49.4%, -45.7%] | 2.2 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 4096 | build | ordered | baseline | 8 | 4.32 ms | 3.06 ms | 0.72× [0.70, 0.73] | -39.5% | [-42.6%, -36.4%] | 3.0 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 4096 | build | ordered | btree-map | 8 | 4.32 ms | 2.79 ms | 0.64× [0.62, 0.66] | -55.8% | [-60.9%, -50.7%] | 4.9 pts | 2.0 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | baseline | 8 | 115 | 109 | 0.97× [0.94, 0.99] | -3.4% | [-6.3%, -0.5%] | 2.7 pts | 2.5 | no | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 115 | 164 | 1.43× [1.36, 1.51] | +30.1% | [+26.5%, +33.6%] | 3.3 pts | 2.9 | no | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1650 | 2849 | 1.76× [1.70, 1.81] | +43.0% | [+41.2%, +44.8%] | 1.7 pts | 3.4 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 1645 | 801 | 0.50× [0.46, 0.55] | -100.9% | [-119.3%, -82.5%] | 17.6 pts | 9.5 | no | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | baseline | 8 | 5855 | 11.6 µs | 1.98× [1.90, 2.07] | +49.5% | [+47.3%, +51.8%] | 2.1 pts | 0.7 | yes | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | btree-map | 8 | 5909 | 2783 | 0.50× [0.46, 0.55] | -100.3% | [-117.5%, -83.2%] | 17.3 pts | 1.7 | no | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | baseline | 8 | 317 | 276 | 0.88× [0.87, 0.89] | -14.0% | [-15.6%, -12.5%] | 1.7 pts | 1.9 | yes | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | btree-map | 8 | 333 | 270 | 0.82× [0.81, 0.84] | -21.7% | [-23.9%, -19.5%] | 3.2 pts | 2.0 | no | yes |
| single-value-ptr | dirs | 16384 | build | ordered | baseline | 8 | 21.09 ms | 16.97 ms | 0.79× [0.77, 0.81] | -26.1% | [-29.1%, -23.1%] | 3.1 pts | 1.3 | no | yes |
| single-value-ptr | dirs | 16384 | build | ordered | btree-map | 8 | 21.36 ms | 15.43 ms | 0.72× [0.71, 0.74] | -38.5% | [-41.8%, -35.3%] | 3.4 pts | 1.8 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 162 | 193 | 1.22× [1.02, 1.50] | +17.9% | [+2.3%, +33.5%] | 15.7 pts | 12.9 | no | yes |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 175 | 266 | 1.52× [1.40, 1.65] | +34.0% | [+28.6%, +39.5%] | 5.7 pts | 4.0 | no | yes |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 1952 | 4216 | 2.14× [1.75, 2.77] | +53.3% | [+42.8%, +63.9%] | 10.8 pts | 18.5 | no | yes |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 2000 | 1463 | 0.73× [0.65, 0.83] | -37.1% | [-53.2%, -20.9%] | 27.2 pts | 10.8 | no | yes |
| single-value-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 39.6 µs | 98.1 µs | 2.64× [2.09, 3.59] | +62.2% | [+52.2%, +72.1%] | 9.8 pts | 30.1 | no | yes |
| single-value-ptr | dirs | 86215 | prefix | ordered | btree-map | 8 | 37.7 µs | 19.1 µs | 0.51× [0.44, 0.60] | -97.1% | [-128.0%, -66.2%] | 33.5 pts | 8.7 | no | yes |
| single-value-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 462 | 477 | 1.01× [0.98, 1.05] | +1.4% | [-2.3%, +5.0%] | 3.8 pts | 4.0 | no | no |
| single-value-ptr | dirs | 86215 | churn | ordered | btree-map | 8 | 518 | 508 | 0.99× [0.98, 1.01] | -0.9% | [-2.4%, +0.6%] | 1.5 pts | 1.3 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr dirs n=4096 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=4096 build: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 9.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 12.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 18.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 10.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 prefix: ordered vs baseline: the processes scatter 30.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 prefix: ordered vs btree-map: the processes scatter 8.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 churn: ordered vs baseline: the pooled interval [-2.31%, 5.04%] includes zero
- single-value-ptr dirs n=86215 churn: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 churn: ordered vs btree-map: the pooled interval [-2.41%, 0.58%] includes zero
