| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 6 | 87.8 | 86.3 | 0.99× [0.97, 1.00] | -1.5% | [-2.6%, -0.4%] | 1.0 pts | 1.3 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 1394 | 2004 | 1.44× [1.41, 1.47] | +30.6% | [+29.2%, +31.9%] | 1.3 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 6 | 206 | 193 | 0.94× [0.92, 0.95] | -6.8% | [-8.7%, -4.9%] | 1.8 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 6 | 3.42 ms | 2.73 ms | 0.79× [0.78, 0.80] | -26.7% | [-28.8%, -24.7%] | 1.9 pts | 1.2 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 137 | 152 | 1.13× [1.07, 1.19] | +11.2% | [+6.8%, +15.6%] | 4.8 pts | 4.5 | no | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1691 | 2815 | 1.68× [1.60, 1.76] | +40.4% | [+37.6%, +43.2%] | 2.7 pts | 6.7 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 330 | 370 | 1.12× [1.11, 1.14] | +11.0% | [+9.6%, +12.5%] | 1.8 pts | 2.6 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 82.07 ms | 78.74 ms | 0.95× [0.94, 0.96] | -5.3% | [-6.5%, -4.1%] | 1.4 pts | 0.9 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 6 | 57.4 | 57.2 | 0.99× [0.97, 1.01] | -1.1% | [-2.7%, +0.6%] | 1.6 pts | 2.7 | yes | no |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 6 | 1099 | 1788 | 1.64× [1.61, 1.66] | +38.9% | [+37.9%, +39.9%] | 1.0 pts | 1.1 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 6 | 139 | 119 | 0.86× [0.85, 0.86] | -16.7% | [-17.8%, -15.6%] | 1.0 pts | 0.6 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 6 | 2.30 ms | 1.70 ms | 0.74× [0.73, 0.75] | -35.1% | [-37.4%, -32.8%] | 2.2 pts | 0.9 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 84.4 | 99.2 | 1.17× [1.16, 1.18] | +14.4% | [+13.5%, +15.3%] | 1.0 pts | 1.6 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 1333 | 2464 | 1.85× [1.84, 1.86] | +45.9% | [+45.6%, +46.3%] | 0.4 pts | 1.1 | yes | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 198 | 214 | 1.08× [1.07, 1.09] | +7.4% | [+6.7%, +8.1%] | 1.4 pts | 2.0 | yes | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 50.97 ms | 46.56 ms | 0.91× [0.89, 0.94] | -9.7% | [-12.7%, -6.7%] | 3.8 pts | 1.7 | no | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 6 | 32.4 | 26.5 | 0.82× [0.82, 0.82] | -22.0% | [-22.6%, -21.4%] | 0.6 pts | 1.1 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 870 | 974 | 1.12× [1.11, 1.12] | +10.6% | [+10.3%, +10.9%] | 0.3 pts | 0.7 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 6 | 78.3 | 47.2 | 0.60× [0.60, 0.60] | -66.5% | [-67.3%, -65.7%] | 0.8 pts | 0.9 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 6 | 1.17 ms | 816.6 µs | 0.69× [0.69, 0.70] | -44.2% | [-45.0%, -43.5%] | 0.7 pts | 0.4 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | baseline | 8 | 34.8 | 33.3 | 0.95× [0.94, 0.96] | -5.0% | [-6.3%, -3.8%] | 1.4 pts | 2.7 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 1723 | 1998 | 1.16× [1.15, 1.17] | +13.8% | [+13.1%, +14.5%] | 0.7 pts | 1.9 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | baseline | 8 | 110 | 86.9 | 0.80× [0.78, 0.82] | -24.7% | [-27.4%, -22.0%] | 3.0 pts | 3.1 | no | yes |
| single-value | u64 | 65536 | build | ordered | baseline | 8 | 25.68 ms | 19.78 ms | 0.78× [0.78, 0.79] | -27.9% | [-28.6%, -27.2%] | 6.3 pts | 2.0 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | baseline | 6 | 79.8 | 73.5 | 0.91× [0.90, 0.92] | -9.5% | [-10.5%, -8.5%] | 1.0 pts | 1.3 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | baseline | 6 | 1441 | 1926 | 1.33× [1.32, 1.34] | +24.7% | [+24.0%, +25.4%] | 0.7 pts | 0.4 | yes | yes |
| single-value | url | 4096 | churn | ordered | baseline | 6 | 210 | 167 | 0.79× [0.78, 0.80] | -26.8% | [-28.2%, -25.3%] | 1.4 pts | 0.7 | yes | yes |
| single-value | url | 4096 | build | ordered | baseline | 6 | 3.24 ms | 2.27 ms | 0.71× [0.70, 0.72] | -40.6% | [-42.9%, -38.3%] | 2.2 pts | 0.9 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | baseline | 8 | 132 | 137 | 1.05× [1.02, 1.08] | +4.5% | [+1.8%, +7.1%] | 2.7 pts | 2.0 | no | yes |
| single-value | url | 65536 | valuesBetween | ordered | baseline | 8 | 1867 | 2667 | 1.43× [1.41, 1.45] | +30.0% | [+28.9%, +31.2%] | 1.1 pts | 1.9 | yes | yes |
| single-value | url | 65536 | churn | ordered | baseline | 8 | 375 | 356 | 0.94× [0.93, 0.95] | -6.2% | [-7.7%, -4.8%] | 1.4 pts | 1.8 | yes | yes |
| single-value | url | 65536 | build | ordered | baseline | 8 | 90.81 ms | 74.03 ms | 0.82× [0.80, 0.83] | -22.6% | [-24.4%, -20.8%] | 2.0 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesBetween: ordered vs baseline: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.73%, 0.60%] includes zero
- single-value street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.69% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.07% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
