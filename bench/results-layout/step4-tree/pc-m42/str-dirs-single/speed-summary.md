| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 103 | 97.1 | 0.95× [0.94, 0.96] | -5.4% | [-6.6%, -4.2%] | 1.1 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 106 | 112 | 1.06× [1.03, 1.09] | +5.2% | [+2.5%, +8.0%] | 2.7 pts | 2.2 | no | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 3065 | 4111 | 1.34× [1.33, 1.35] | +25.5% | [+24.9%, +26.1%] | 0.8 pts | 0.8 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 3020 | 1429 | 0.46× [0.46, 0.47] | -115.7% | [-118.9%, -112.4%] | 3.1 pts | 0.9 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 3545 | 4759 | 1.35× [1.33, 1.38] | +26.2% | [+25.0%, +27.4%] | 1.2 pts | 0.7 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 3570 | 1775 | 0.50× [0.49, 0.50] | -100.9% | [-103.0%, -98.7%] | 2.6 pts | 0.5 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 1584 | 222 | 0.14× [0.13, 0.14] | -638.6% | [-656.7%, -620.5%] | 18.4 pts | 0.8 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 1589 | 193 | 0.12× [0.12, 0.12] | -739.1% | [-749.1%, -729.1%] | 10.5 pts | 0.7 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 18.03 ms | 3.05 ms | 0.17× [0.17, 0.17] | -487.4% | [-497.3%, -477.4%] | 20.6 pts | 2.5 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 18.12 ms | 2.73 ms | 0.15× [0.15, 0.16] | -552.5% | [-569.5%, -535.4%] | 19.5 pts | 1.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 127 | 126 | 1.00× [0.99, 1.01] | -0.3% | [-1.5%, +0.8%] | 1.2 pts | 1.2 | yes | no |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 129 | 162 | 1.26× [1.23, 1.30] | +20.9% | [+18.7%, +23.2%] | 2.2 pts | 2.1 | no | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3414 | 4747 | 1.39× [1.38, 1.40] | +28.0% | [+27.4%, +28.6%] | 0.5 pts | 1.4 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 3405 | 1884 | 0.55× [0.55, 0.56] | -81.3% | [-82.6%, -80.0%] | 1.5 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 8 | 16.2 µs | 24.6 µs | 1.54× [1.52, 1.56] | +35.1% | [+34.0%, +36.1%] | 1.0 pts | 0.9 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 8 | 14.4 µs | 8034 | 0.55× [0.54, 0.57] | -80.3% | [-84.0%, -76.6%] | 5.6 pts | 0.6 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 8 | 1664 | 328 | 0.20× [0.19, 0.20] | -408.0% | [-418.0%, -398.0%] | 10.1 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 8 | 1677 | 289 | 0.17× [0.17, 0.18] | -473.5% | [-492.6%, -454.4%] | 18.8 pts | 1.3 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 8 | 77.00 ms | 14.83 ms | 0.19× [0.19, 0.20] | -420.5% | [-428.7%, -412.3%] | 10.7 pts | 0.9 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 8 | 77.34 ms | 14.50 ms | 0.19× [0.19, 0.19] | -430.9% | [-440.4%, -421.5%] | 11.2 pts | 0.8 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 172 | 185 | 1.08× [1.06, 1.10] | +7.6% | [+5.9%, +9.3%] | 2.4 pts | 2.4 | yes | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 183 | 260 | 1.41× [1.38, 1.43] | +28.9% | [+27.5%, +30.2%] | 1.4 pts | 1.3 | yes | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3713 | 5295 | 1.44× [1.42, 1.46] | +30.5% | [+29.4%, +31.6%] | 1.1 pts | 1.7 | yes | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 3712 | 2413 | 0.65× [0.63, 0.67] | -53.4% | [-58.3%, -48.6%] | 8.2 pts | 3.3 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | baseline | 8 | 91.1 µs | 145.8 µs | 1.60× [1.59, 1.62] | +37.6% | [+37.1%, +38.2%] | 0.6 pts | 1.1 | yes | yes |
| single-value-str | dirs | 86215 | prefix | ordered | btree-map | 8 | 86.6 µs | 52.0 µs | 0.59× [0.57, 0.60] | -70.3% | [-75.2%, -65.4%] | 5.3 pts | 1.4 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | baseline | 8 | 2067 | 585 | 0.28× [0.28, 0.29] | -251.7% | [-260.7%, -242.6%] | 9.9 pts | 1.7 | yes | yes |
| single-value-str | dirs | 86215 | churn | ordered | btree-map | 8 | 2118 | 594 | 0.28× [0.28, 0.29] | -252.5% | [-257.3%, -247.8%] | 6.3 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.27% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 build: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.34% does not clear the 0.41% noise floor, the bound on what the harness reports between identical code in every process
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.45%, 0.77%] includes zero
- single-value-str dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
