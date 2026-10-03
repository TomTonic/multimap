| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 75.9 | 108 | 1.43× [1.41, 1.45] | +29.9% | [+29.0%, +30.8%] | 1.0 pts | 1.1 | yes | yes |
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage | 8 | 75.0 | 110 | 1.48× [1.44, 1.52] | +32.3% | [+30.5%, +34.1%] | 1.8 pts | 3.4 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1982 | 1523 | 0.75× [0.74, 0.77] | -32.5% | [-34.7%, -30.3%] | 2.2 pts | 0.8 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage | 8 | 2000 | 1818 | 0.91× [0.90, 0.92] | -9.4% | [-10.5%, -8.2%] | 1.4 pts | 0.8 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 2220 | 1866 | 0.84× [0.83, 0.86] | -18.6% | [-20.4%, -16.7%] | 2.5 pts | 1.1 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage | 8 | 2212 | 2215 | 1.00× [0.98, 1.01] | -0.2% | [-1.5%, +1.1%] | 1.9 pts | 1.1 | yes | no |
| unique-str | dirs | 4096 | churn | ordered | btree-map | 8 | 195 | 182 | 0.93× [0.91, 0.94] | -7.9% | [-9.3%, -6.5%] | 1.8 pts | 1.2 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage | 8 | 198 | 251 | 1.27× [1.26, 1.28] | +21.1% | [+20.7%, +21.6%] | 0.5 pts | 0.6 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | btree-map | 8 | 2.78 ms | 2.67 ms | 0.96× [0.95, 0.97] | -4.0% | [-4.9%, -3.0%] | 0.9 pts | 0.9 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage | 8 | 2.79 ms | 4.76 ms | 1.71× [1.70, 1.72] | +41.4% | [+41.0%, +41.7%] | 0.5 pts | 0.8 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | btree-map | 12 | 105 | 159 | 1.52× [1.49, 1.54] | +34.0% | [+33.0%, +35.0%] | 1.2 pts | 1.4 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage | 12 | 103 | 136 | 1.34× [1.32, 1.36] | +25.3% | [+24.4%, +26.2%] | 1.7 pts | 2.4 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | btree-map | 12 | 2521 | 1975 | 0.78× [0.78, 0.79] | -27.4% | [-28.2%, -26.6%] | 1.1 pts | 1.0 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage | 12 | 2511 | 2206 | 0.88× [0.87, 0.88] | -14.2% | [-15.0%, -13.4%] | 1.5 pts | 1.6 | yes | yes |
| unique-str | dirs | 16384 | prefix | ordered | btree-map | 12 | 10.1 µs | 8899 | 0.86× [0.84, 0.88] | -16.1% | [-18.4%, -13.8%] | 3.7 pts | 0.9 | no | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage | 12 | 10.1 µs | 9039 | 0.88× [0.85, 0.92] | -13.2% | [-17.9%, -8.5%] | 5.2 pts | 1.1 | no | yes |
| unique-str | dirs | 16384 | churn | ordered | btree-map | 12 | 248 | 247 | 0.99× [0.98, 1.00] | -0.8% | [-1.5%, -0.0%] | 1.4 pts | 1.3 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage | 12 | 263 | 320 | 1.23× [1.22, 1.23] | +18.4% | [+17.9%, +19.0%] | 0.8 pts | 0.7 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | btree-map | 12 | 13.67 ms | 13.63 ms | 1.00× [0.99, 1.01] | -0.1% | [-1.0%, +0.7%] | 1.0 pts | 0.8 | yes | no |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage | 12 | 13.51 ms | 21.34 ms | 1.57× [1.56, 1.58] | +36.2% | [+35.7%, +36.7%] | 0.7 pts | 0.9 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | btree-map | 12 | 174 | 233 | 1.35× [1.28, 1.43] | +25.8% | [+21.7%, +29.8%] | 4.6 pts | 3.0 | no | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage | 12 | 170 | 190 | 1.11× [1.07, 1.15] | +10.0% | [+6.7%, +13.2%] | 4.7 pts | 3.3 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | btree-map | 12 | 3011 | 2328 | 0.79× [0.76, 0.82] | -26.8% | [-32.3%, -21.3%] | 7.5 pts | 3.8 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage | 12 | 2899 | 2537 | 0.87× [0.86, 0.88] | -15.2% | [-16.5%, -13.8%] | 1.4 pts | 1.4 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | btree-map | 12 | 76.8 µs | 54.2 µs | 0.69× [0.68, 0.70] | -44.5% | [-46.4%, -42.6%] | 3.3 pts | 1.1 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage | 12 | 74.6 µs | 58.4 µs | 0.77× [0.77, 0.78] | -29.3% | [-30.6%, -28.0%] | 1.5 pts | 0.9 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | btree-map | 12 | 475 | 443 | 0.95× [0.93, 0.96] | -5.5% | [-7.1%, -3.9%] | 1.7 pts | 1.3 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage | 12 | 434 | 473 | 1.10× [1.07, 1.12] | +8.9% | [+6.7%, +11.1%] | 2.7 pts | 1.5 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str dirs n=4096 valuesFor: ordered vs ordered-lpage: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=4096 prefix: ordered vs ordered-lpage: the pooled difference of -0.21% does not clear the 1.33% noise floor, the bound on what the harness reports between identical code in every process
- unique-str dirs n=4096 prefix: ordered vs ordered-lpage: the pooled interval [-1.55%, 1.14%] includes zero
- unique-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 build: ordered vs btree-map: the pooled difference of -0.12% does not clear the 0.46% noise floor, the bound on what the harness reports between identical code in every process
- unique-str dirs n=16384 build: ordered vs btree-map: the pooled interval [-0.97%, 0.73%] includes zero
- unique-str dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the A/A validations found a systematic difference of -0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 prefix: ordered vs ordered-lpage: the A/A validations found a systematic difference of -1.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
