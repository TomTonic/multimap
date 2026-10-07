| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 97.5 | 78.1 | 0.80× [0.79, 0.81] | -25.2% | [-27.0%, -23.5%] | 1.7 pts | 1.5 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 98.8 | 109 | 1.11× [1.10, 1.13] | +10.3% | [+9.4%, +11.2%] | 1.0 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1615 | 2334 | 1.45× [1.43, 1.48] | +31.3% | [+30.3%, +32.2%] | 1.0 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1581 | 558 | 0.35× [0.35, 0.36] | -184.6% | [-188.8%, -180.3%] | 4.8 pts | 1.7 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 8 | 1947 | 2672 | 1.35× [1.33, 1.38] | +25.9% | [+24.5%, +27.3%] | 1.3 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | btree-map | 8 | 1961 | 667 | 0.34× [0.33, 0.34] | -196.4% | [-200.7%, -192.1%] | 4.4 pts | 1.3 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 257 | 204 | 0.80× [0.79, 0.80] | -25.3% | [-26.3%, -24.3%] | 1.1 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 260 | 184 | 0.70× [0.70, 0.71] | -42.1% | [-43.3%, -40.8%] | 1.8 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 3.83 ms | 3.02 ms | 0.79× [0.78, 0.81] | -26.4% | [-28.9%, -24.0%] | 2.3 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 3.82 ms | 2.65 ms | 0.70× [0.69, 0.71] | -42.6% | [-44.8%, -40.4%] | 3.1 pts | 2.1 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 119 | 108 | 0.91× [0.88, 0.93] | -10.5% | [-14.0%, -7.0%] | 3.4 pts | 2.8 | no | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 120 | 159 | 1.32× [1.26, 1.39] | +24.3% | [+20.8%, +27.8%] | 3.3 pts | 3.0 | no | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1730 | 2910 | 1.68× [1.66, 1.70] | +40.4% | [+39.7%, +41.1%] | 0.6 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 1706 | 734 | 0.44× [0.42, 0.47] | -125.8% | [-140.7%, -110.9%] | 14.8 pts | 6.7 | no | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 8 | 6502 | 11.5 µs | 1.73× [1.69, 1.79] | +42.4% | [+40.7%, +44.0%] | 1.8 pts | 0.5 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | btree-map | 8 | 6715 | 2498 | 0.37× [0.36, 0.39] | -167.8% | [-176.1%, -159.5%] | 9.1 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 292 | 270 | 0.92× [0.91, 0.93] | -8.9% | [-9.8%, -8.0%] | 1.0 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 305 | 260 | 0.86× [0.85, 0.87] | -15.9% | [-17.4%, -14.4%] | 1.5 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 18.31 ms | 15.66 ms | 0.87× [0.85, 0.88] | -15.4% | [-17.3%, -13.5%] | 3.0 pts | 1.7 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 18.42 ms | 14.80 ms | 0.80× [0.79, 0.81] | -25.2% | [-27.4%, -23.0%] | 3.2 pts | 1.6 | yes | yes |
| single-value | dirs | 86215 | valuesFor | ordered | baseline | 8 | 159 | 166 | 1.04× [1.03, 1.06] | +4.2% | [+2.8%, +5.5%] | 1.6 pts | 1.5 | yes | yes |
| single-value | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 163 | 236 | 1.45× [1.40, 1.51] | +31.2% | [+28.5%, +33.9%] | 2.8 pts | 2.4 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 1888 | 3287 | 1.76× [1.67, 1.85] | +43.1% | [+40.1%, +46.0%] | 2.9 pts | 5.4 | yes | yes |
| single-value | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 1880 | 1030 | 0.56× [0.52, 0.59] | -79.9% | [-90.5%, -69.3%] | 13.4 pts | 4.3 | no | yes |
| single-value | dirs | 86215 | prefix | ordered | baseline | 8 | 38.8 µs | 81.2 µs | 2.11× [2.06, 2.16] | +52.5% | [+51.4%, +53.6%] | 1.0 pts | 3.5 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | btree-map | 8 | 36.6 µs | 14.8 µs | 0.39× [0.38, 0.40] | -155.9% | [-162.9%, -148.9%] | 8.6 pts | 3.1 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | baseline | 8 | 413 | 442 | 1.06× [1.05, 1.07] | +5.8% | [+4.9%, +6.7%] | 0.9 pts | 1.2 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | btree-map | 8 | 453 | 470 | 1.06× [1.02, 1.09] | +5.3% | [+2.3%, +8.3%] | 3.0 pts | 3.0 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 build: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.59% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of +1.92% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=86215 prefix: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 prefix: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=86215 churn: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
