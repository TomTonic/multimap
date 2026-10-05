| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 49.6 | 46.1 | 0.93× [0.92, 0.95] | -7.2% | [-8.7%, -5.8%] | 1.5 pts | 2.5 | yes | yes |
| single-value-ptr | street | 4096 | valuesFor | ordered | btree-map | 8 | 50.2 | 91.6 | 1.83× [1.80, 1.85] | +45.3% | [+44.5%, +46.0%] | 0.9 pts | 2.5 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 2139 | 1800 | 0.84× [0.83, 0.86] | -18.6% | [-20.5%, -16.6%] | 2.1 pts | 1.1 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | btree-map | 8 | 2112 | 506 | 0.24× [0.24, 0.24] | -319.0% | [-323.2%, -314.8%] | 5.6 pts | 1.3 | yes | yes |
| single-value-ptr | street | 4096 | prefix | ordered | baseline | 8 | 234 | 198 | 0.85× [0.84, 0.86] | -17.2% | [-18.6%, -15.8%] | 1.5 pts | 1.2 | yes | yes |
| single-value-ptr | street | 4096 | prefix | ordered | btree-map | 8 | 233 | 134 | 0.57× [0.56, 0.58] | -74.2% | [-77.0%, -71.3%] | 2.7 pts | 1.6 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | baseline | 8 | 129 | 117 | 0.90× [0.89, 0.91] | -10.6% | [-12.0%, -9.3%] | 1.3 pts | 0.9 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | btree-map | 8 | 129 | 136 | 1.06× [1.03, 1.08] | +5.2% | [+3.3%, +7.1%] | 1.8 pts | 2.1 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | baseline | 8 | 1.92 ms | 1.67 ms | 0.87× [0.87, 0.88] | -14.4% | [-15.1%, -13.7%] | 0.7 pts | 0.9 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | btree-map | 8 | 1.92 ms | 2.08 ms | 1.10× [1.07, 1.13] | +9.1% | [+6.5%, +11.7%] | 2.8 pts | 2.5 | no | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | baseline | 8 | 69.1 | 64.8 | 0.94× [0.92, 0.95] | -6.8% | [-8.7%, -4.9%] | 2.5 pts | 2.9 | yes | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | btree-map | 8 | 70.2 | 126 | 1.81× [1.77, 1.84] | +44.6% | [+43.5%, +45.7%] | 1.4 pts | 3.5 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | baseline | 8 | 2523 | 2199 | 0.87× [0.86, 0.88] | -14.4% | [-15.6%, -13.2%] | 1.3 pts | 1.2 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | btree-map | 8 | 2522 | 628 | 0.25× [0.24, 0.25] | -306.7% | [-319.1%, -294.3%] | 14.6 pts | 5.0 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | baseline | 8 | 774 | 647 | 0.83× [0.82, 0.84] | -20.3% | [-21.7%, -18.8%] | 1.6 pts | 0.9 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | btree-map | 8 | 768 | 289 | 0.37× [0.37, 0.38] | -167.2% | [-171.9%, -162.5%] | 5.5 pts | 2.4 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | baseline | 8 | 163 | 152 | 0.92× [0.92, 0.93] | -8.3% | [-9.2%, -7.3%] | 1.0 pts | 1.0 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | btree-map | 8 | 162 | 188 | 1.17× [1.15, 1.18] | +14.3% | [+13.2%, +15.4%] | 1.0 pts | 1.4 | yes | yes |
| single-value-ptr | street | 16384 | build | ordered | baseline | 8 | 9.15 ms | 8.33 ms | 0.91× [0.90, 0.92] | -9.9% | [-10.5%, -9.2%] | 0.7 pts | 1.1 | yes | yes |
| single-value-ptr | street | 16384 | build | ordered | btree-map | 8 | 9.22 ms | 11.03 ms | 1.17× [1.14, 1.20] | +14.5% | [+12.5%, +16.6%] | 2.5 pts | 2.3 | no | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 180 | 158 | 0.90× [0.86, 0.94] | -11.5% | [-16.6%, -6.4%] | 4.7 pts | 3.4 | no | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | btree-map | 8 | 227 | 240 | 1.09× [0.99, 1.20] | +8.0% | [-0.5%, +16.5%] | 8.7 pts | 5.7 | no | no |
| single-value-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 3843 | 3435 | 0.88× [0.85, 0.91] | -13.7% | [-17.7%, -9.8%] | 3.7 pts | 2.4 | no | yes |
| single-value-ptr | street | 212449 | valuesBetween | ordered | btree-map | 8 | 3702 | 1239 | 0.33× [0.32, 0.35] | -202.4% | [-216.1%, -188.7%] | 16.6 pts | 1.1 | yes | yes |
| single-value-ptr | street | 212449 | prefix | ordered | baseline | 8 | 14.4 µs | 12.1 µs | 0.87× [0.84, 0.90] | -15.2% | [-19.4%, -11.1%] | 3.9 pts | 3.0 | no | yes |
| single-value-ptr | street | 212449 | prefix | ordered | btree-map | 8 | 13.6 µs | 3565 | 0.26× [0.26, 0.27] | -281.1% | [-286.0%, -276.2%] | 7.9 pts | 0.6 | yes | yes |
| single-value-ptr | street | 212449 | churn | ordered | baseline | 8 | 381 | 362 | 0.95× [0.94, 0.96] | -5.1% | [-5.8%, -4.4%] | 0.7 pts | 0.4 | yes | yes |
| single-value-ptr | street | 212449 | churn | ordered | btree-map | 8 | 490 | 464 | 0.96× [0.93, 1.00] | -4.1% | [-8.0%, -0.3%] | 3.7 pts | 3.0 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.21% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr street n=4096 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 build: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 prefix: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of -0.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr street n=16384 build: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.70% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr street n=212449 valuesFor: ordered vs btree-map: the pooled interval [-0.53%, 16.48%] includes zero
- single-value-ptr street n=212449 valuesFor: ordered vs btree-map: the processes scatter 5.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 prefix: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 churn: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
