| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | street | 4096 | valuesFor | ordered | btree-sets | 12 | 52.1 | 135 | 2.60× [2.57, 2.63] | +61.5% | [+61.1%, +62.0%] | 0.5 pts | 1.7 | yes | yes |
| multi | street | 4096 | valuesFor | ordered | hashed | 12 | 52.0 | 22.1 | 0.43× [0.42, 0.43] | -134.8% | [-136.8%, -132.8%] | 2.6 pts | 1.3 | yes | yes |
| multi | street | 4096 | valuesFor | ordered | map-sets | 12 | 51.8 | 64.6 | 1.24× [1.23, 1.25] | +19.5% | [+18.9%, +20.2%] | 1.3 pts | 2.2 | yes | yes |
| multi | street | 4096 | valuesBetween | ordered | btree-sets | 12 | 2264 | 4581 | 2.02× [1.99, 2.05] | +50.5% | [+49.8%, +51.2%] | 1.1 pts | 2.0 | yes | yes |
| multi | street | 4096 | valuesBetween | ordered | hashed | 12 | 2339 | 53.8 µs | 23.01× [22.71, 23.32] | +95.7% | [+95.6%, +95.7%] | 0.1 pts | 1.0 | yes | yes |
| multi | street | 4096 | valuesBetween | ordered | map-sets | 12 | 2329 | 54.2 µs | 23.36× [22.92, 23.81] | +95.7% | [+95.6%, +95.8%] | 0.1 pts | 1.2 | yes | yes |
| multi | street | 4096 | prefix | ordered | btree-sets | 12 | 223 | 527 | 2.35× [2.28, 2.42] | +57.4% | [+56.2%, +58.6%] | 1.4 pts | 2.7 | yes | yes |
| multi | street | 4096 | prefix | ordered | hashed | 12 | 254 | 49.5 µs | 197.10× [193.91, 200.40] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 1.0 | yes | yes |
| multi | street | 4096 | prefix | ordered | map-sets | 12 | 248 | 46.3 µs | 186.12× [181.93, 190.51] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 0.8 | yes | yes |
| multi | street | 4096 | churn | ordered | btree-sets | 12 | 98.1 | 190 | 1.95× [1.93, 1.98] | +48.8% | [+48.1%, +49.6%] | 0.9 pts | 1.3 | yes | yes |
| multi | street | 4096 | churn | ordered | hashed | 12 | 94.0 | 50.2 | 0.53× [0.52, 0.54] | -88.4% | [-90.5%, -86.3%] | 3.8 pts | 2.3 | yes | yes |
| multi | street | 4096 | churn | ordered | map-sets | 12 | 94.7 | 63.1 | 0.67× [0.66, 0.68] | -49.4% | [-52.4%, -46.4%] | 4.0 pts | 2.6 | yes | yes |
| multi | street | 4096 | build | ordered | btree-sets | 12 | 3.84 ms | 7.45 ms | 1.95× [1.91, 1.98] | +48.6% | [+47.8%, +49.5%] | 0.9 pts | 1.3 | yes | yes |
| multi | street | 4096 | build | ordered | hashed | 12 | 3.85 ms | 2.43 ms | 0.63× [0.61, 0.65] | -58.3% | [-63.0%, -53.7%] | 4.6 pts | 1.8 | yes | yes |
| multi | street | 4096 | build | ordered | map-sets | 12 | 3.85 ms | 3.01 ms | 0.79× [0.76, 0.81] | -27.3% | [-31.1%, -23.5%] | 5.0 pts | 1.5 | no | yes |
| multi | street | 16384 | valuesFor | ordered | btree-sets | 8 | 73.8 | 184 | 2.49× [2.44, 2.55] | +59.9% | [+59.0%, +60.7%] | 1.0 pts | 4.2 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | hashed | 8 | 71.5 | 27.7 | 0.39× [0.38, 0.39] | -159.7% | [-163.4%, -156.0%] | 3.6 pts | 2.0 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | map-sets | 8 | 71.7 | 76.8 | 1.08× [1.06, 1.09] | +7.1% | [+5.7%, +8.6%] | 1.4 pts | 2.2 | yes | yes |
| multi | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 2860 | 5314 | 1.85× [1.84, 1.87] | +46.1% | [+45.7%, +46.5%] | 0.6 pts | 1.7 | yes | yes |
| multi | street | 16384 | valuesBetween | ordered | hashed | 8 | 3155 | 228.7 µs | 68.50× [62.86, 75.25] | +98.5% | [+98.4%, +98.7%] | 0.1 pts | 1.7 | yes | yes |
| multi | street | 16384 | valuesBetween | ordered | map-sets | 8 | 3232 | 222.8 µs | 67.74× [66.09, 69.49] | +98.5% | [+98.5%, +98.6%] | 0.1 pts | 0.7 | yes | yes |
| multi | street | 16384 | prefix | ordered | btree-sets | 8 | 810 | 2021 | 2.50× [2.47, 2.53] | +59.9% | [+59.4%, +60.4%] | 0.5 pts | 0.8 | yes | yes |
| multi | street | 16384 | prefix | ordered | hashed | 8 | 1138 | 227.1 µs | 199.98× [196.22, 203.89] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 0.7 | yes | yes |
| multi | street | 16384 | prefix | ordered | map-sets | 8 | 1147 | 216.3 µs | 188.87× [185.63, 192.22] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 0.6 | yes | yes |
| multi | street | 16384 | churn | ordered | btree-sets | 8 | 136 | 263 | 1.91× [1.88, 1.95] | +47.7% | [+46.7%, +48.7%] | 1.0 pts | 1.4 | yes | yes |
| multi | street | 16384 | churn | ordered | hashed | 8 | 129 | 63.8 | 0.49× [0.48, 0.50] | -103.0% | [-106.3%, -99.8%] | 3.5 pts | 1.7 | yes | yes |
| multi | street | 16384 | churn | ordered | map-sets | 8 | 130 | 82.0 | 0.63× [0.62, 0.65] | -58.1% | [-62.4%, -53.9%] | 5.4 pts | 2.5 | yes | yes |
| multi | street | 16384 | build | ordered | btree-sets | 8 | 22.20 ms | 40.65 ms | 1.83× [1.81, 1.85] | +45.4% | [+44.7%, +46.0%] | 1.2 pts | 1.5 | yes | yes |
| multi | street | 16384 | build | ordered | hashed | 8 | 21.73 ms | 11.81 ms | 0.54× [0.53, 0.55] | -84.2% | [-87.5%, -80.8%] | 5.1 pts | 1.8 | yes | yes |
| multi | street | 16384 | build | ordered | map-sets | 8 | 22.16 ms | 15.03 ms | 0.68× [0.66, 0.70] | -47.7% | [-52.1%, -43.4%] | 4.3 pts | 1.1 | yes | yes |
| multi | street | 212449 | valuesFor | ordered | btree-sets | 12 | 249 | 480 | 1.90× [1.87, 1.92] | +47.2% | [+46.6%, +47.9%] | 2.1 pts | 2.6 | yes | yes |
| multi | street | 212449 | valuesFor | ordered | hashed | 12 | 194 | 94.4 | 0.49× [0.48, 0.50] | -102.8% | [-107.3%, -98.2%] | 6.2 pts | 2.3 | yes | yes |
| multi | street | 212449 | valuesFor | ordered | map-sets | 12 | 212 | 269 | 1.23× [1.18, 1.29] | +18.8% | [+15.2%, +22.4%] | 4.2 pts | 4.8 | no | yes |
| multi | street | 212449 | valuesBetween | ordered | btree-sets | 12 | 6934 | 16.7 µs | 2.38× [2.33, 2.44] | +58.0% | [+57.0%, +58.9%] | 1.3 pts | 3.0 | yes | yes |
| multi | street | 212449 | prefix | ordered | btree-sets | 12 | 22.3 µs | 53.7 µs | 2.44× [2.37, 2.51] | +59.0% | [+57.8%, +60.2%] | 1.3 pts | 1.4 | yes | yes |
| multi | street | 212449 | churn | ordered | btree-sets | 12 | 518 | 705 | 1.36× [1.35, 1.38] | +26.6% | [+25.9%, +27.4%] | 1.2 pts | 1.7 | yes | yes |
| multi | street | 212449 | churn | ordered | hashed | 12 | 408 | 242 | 0.57× [0.56, 0.58] | -76.0% | [-78.7%, -73.4%] | 6.6 pts | 1.2 | yes | yes |
| multi | street | 212449 | churn | ordered | map-sets | 12 | 429 | 297 | 0.69× [0.69, 0.70] | -44.7% | [-46.0%, -43.5%] | 1.9 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi street n=4096 valuesFor: ordered vs map-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 valuesBetween: ordered vs hashed: the A/A validations found a systematic difference of +0.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=4096 prefix: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 prefix: ordered vs hashed: the A/A validations found a systematic difference of +1.89% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=4096 churn: ordered vs hashed: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 churn: ordered vs map-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=4096 build: ordered vs map-sets: the A/A validations found a systematic difference of -1.08% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 valuesFor: ordered vs map-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +3.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +4.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=16384 churn: ordered vs map-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=212449 valuesFor: ordered vs hashed: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=212449 valuesFor: ordered vs map-sets: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=212449 churn: ordered vs map-sets: the A/A validations found a systematic difference of +0.37% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
