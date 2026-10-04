| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 95.9 | 76.4 | 0.80× [0.79, 0.81] | -25.1% | [-27.2%, -23.1%] | 2.2 pts | 1.8 | yes | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 101 | 112 | 1.12× [1.08, 1.16] | +10.7% | [+7.5%, +14.0%] | 4.0 pts | 2.7 | no | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 4039 | 2111 | 0.52× [0.51, 0.53] | -92.2% | [-94.7%, -89.7%] | 2.8 pts | 0.9 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 4028 | 1544 | 0.37× [0.37, 0.38] | -167.1% | [-172.0%, -162.1%] | 4.9 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 4715 | 2302 | 0.50× [0.48, 0.51] | -102.0% | [-107.2%, -96.8%] | 5.6 pts | 0.9 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 4725 | 1914 | 0.41× [0.40, 0.42] | -145.4% | [-150.7%, -140.0%] | 5.4 pts | 0.8 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 204 | 192 | 0.93× [0.92, 0.93] | -7.9% | [-8.7%, -7.1%] | 5.4 pts | 4.9 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 202 | 182 | 0.90× [0.88, 0.91] | -11.7% | [-13.8%, -9.6%] | 2.2 pts | 1.5 | no | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 3.18 ms | 2.70 ms | 0.85× [0.83, 0.87] | -18.0% | [-21.2%, -14.8%] | 3.2 pts | 2.5 | no | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 3.21 ms | 2.65 ms | 0.84× [0.82, 0.86] | -19.1% | [-21.6%, -16.6%] | 3.0 pts | 1.7 | no | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 128 | 107 | 0.83× [0.80, 0.85] | -21.1% | [-25.0%, -17.3%] | 4.4 pts | 3.7 | no | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 129 | 164 | 1.26× [1.21, 1.31] | +20.5% | [+17.5%, +23.4%] | 3.4 pts | 2.9 | no | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 4720 | 2684 | 0.57× [0.56, 0.57] | -75.8% | [-77.4%, -74.2%] | 1.9 pts | 1.3 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 4701 | 2020 | 0.43× [0.42, 0.43] | -133.9% | [-136.1%, -131.7%] | 2.4 pts | 1.4 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 8 | 23.2 µs | 10.8 µs | 0.46× [0.46, 0.47] | -115.3% | [-119.0%, -111.6%] | 3.5 pts | 0.2 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 8 | 24.1 µs | 8779 | 0.36× [0.35, 0.38] | -176.2% | [-187.0%, -165.3%] | 11.2 pts | 0.7 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 8 | 266 | 249 | 0.94× [0.93, 0.96] | -6.1% | [-7.6%, -4.6%] | 1.5 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 8 | 259 | 246 | 0.94× [0.91, 0.97] | -6.4% | [-9.5%, -3.2%] | 3.1 pts | 2.4 | no | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 8 | 15.01 ms | 13.13 ms | 0.88× [0.87, 0.88] | -14.1% | [-15.2%, -13.1%] | 1.2 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 8 | 15.04 ms | 13.70 ms | 0.92× [0.91, 0.93] | -8.7% | [-10.2%, -7.2%] | 1.5 pts | 1.0 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 201 | 176 | 0.88× [0.85, 0.92] | -13.6% | [-17.9%, -9.3%] | 5.0 pts | 3.6 | no | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 223 | 259 | 1.15× [1.12, 1.20] | +13.4% | [+10.5%, +16.4%] | 4.2 pts | 2.9 | no | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 5540 | 3329 | 0.61× [0.58, 0.63] | -65.0% | [-71.2%, -58.7%] | 6.2 pts | 2.6 | yes | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 5694 | 2770 | 0.49× [0.46, 0.53] | -102.5% | [-116.7%, -88.3%] | 13.2 pts | 3.4 | no | yes |
| single-value-str | dirs | 86215 | prefix | ordered | baseline | 8 | 150.5 µs | 87.1 µs | 0.56× [0.56, 0.57] | -77.5% | [-79.4%, -75.6%] | 3.4 pts | 0.5 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | btree-map | 8 | 147.7 µs | 63.5 µs | 0.40× [0.39, 0.41] | -152.0% | [-157.5%, -146.5%] | 5.6 pts | 0.6 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | baseline | 8 | 432 | 401 | 0.93× [0.92, 0.95] | -7.3% | [-8.9%, -5.8%] | 1.5 pts | 0.9 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | btree-map | 8 | 485 | 434 | 0.90× [0.89, 0.92] | -10.9% | [-12.7%, -9.1%] | 2.0 pts | 1.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 churn: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 build: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 churn: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -3.02% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
