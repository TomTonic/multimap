| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 107 | 97.0 | 0.90× [0.88, 0.92] | -11.1% | [-14.0%, -8.2%] | 2.8 pts | 2.5 | no | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 110 | 111 | 1.00× [0.97, 1.03] | +0.1% | [-3.0%, +3.2%] | 3.9 pts | 2.7 | no | no |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 3271 | 4208 | 1.28× [1.27, 1.29] | +21.9% | [+21.2%, +22.6%] | 0.7 pts | 0.8 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 3211 | 1469 | 0.45× [0.44, 0.46] | -122.0% | [-127.9%, -116.0%] | 5.5 pts | 1.8 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 3794 | 4859 | 1.27× [1.23, 1.31] | +21.3% | [+19.0%, +23.5%] | 2.1 pts | 0.9 | no | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 3749 | 1819 | 0.50× [0.49, 0.50] | -101.6% | [-104.9%, -98.2%] | 3.7 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 296 | 215 | 0.73× [0.71, 0.74] | -37.9% | [-40.3%, -35.5%] | 2.5 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 298 | 191 | 0.64× [0.63, 0.64] | -57.3% | [-58.7%, -55.8%] | 1.9 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 4.64 ms | 3.31 ms | 0.70× [0.69, 0.71] | -42.4% | [-44.9%, -39.9%] | 3.3 pts | 1.6 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 4.88 ms | 2.88 ms | 0.59× [0.57, 0.61] | -69.2% | [-74.6%, -63.7%] | 5.4 pts | 1.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 132 | 135 | 1.07× [0.97, 1.20] | +6.8% | [-3.0%, +16.6%] | 11.1 pts | 8.6 | no | no |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 134 | 169 | 1.28× [1.23, 1.33] | +21.9% | [+18.9%, +24.9%] | 5.4 pts | 4.9 | no | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3644 | 4905 | 1.44× [1.30, 1.62] | +30.7% | [+23.1%, +38.2%] | 7.7 pts | 15.7 | no | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 3618 | 2040 | 0.60× [0.55, 0.66] | -67.6% | [-83.4%, -51.9%] | 26.4 pts | 11.6 | no | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 8 | 16.7 µs | 26.3 µs | 1.54× [1.46, 1.63] | +35.0% | [+31.4%, +38.7%] | 3.8 pts | 4.5 | no | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 8 | 15.5 µs | 8583 | 0.57× [0.55, 0.60] | -74.4% | [-81.6%, -67.2%] | 14.2 pts | 1.1 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 8 | 367 | 286 | 0.78× [0.77, 0.79] | -28.5% | [-30.3%, -26.7%] | 2.5 pts | 2.2 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 8 | 385 | 291 | 0.75× [0.71, 0.78] | -34.1% | [-40.3%, -27.8%] | 6.0 pts | 4.5 | no | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 8 | 23.26 ms | 17.97 ms | 0.77× [0.76, 0.79] | -29.3% | [-32.3%, -26.3%] | 3.2 pts | 1.9 | no | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 8 | 23.59 ms | 16.32 ms | 0.69× [0.67, 0.72] | -43.9% | [-49.2%, -38.7%] | 5.3 pts | 3.1 | no | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 178 | 202 | 1.13× [1.02, 1.26] | +11.3% | [+2.2%, +20.4%] | 9.2 pts | 8.8 | no | yes |
| single-value-str | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 195 | 303 | 1.57× [1.39, 1.82] | +36.5% | [+28.1%, +44.9%] | 8.4 pts | 6.5 | no | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3991 | 6263 | 1.52× [1.37, 1.71] | +34.3% | [+27.0%, +41.6%] | 7.8 pts | 14.3 | no | yes |
| single-value-str | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 4039 | 3370 | 0.81× [0.69, 0.96] | -24.2% | [-44.5%, -3.9%] | 21.6 pts | 14.2 | no | yes |
| single-value-str | dirs | 86215 | prefix | ordered | baseline | 8 | 101.9 µs | 156.0 µs | 1.63× [1.51, 1.77] | +38.7% | [+33.8%, +43.5%] | 5.9 pts | 9.9 | no | yes |
| single-value-str | dirs | 86215 | prefix | ordered | btree-map | 8 | 95.4 µs | 71.0 µs | 0.71× [0.60, 0.88] | -39.9% | [-65.6%, -14.2%] | 31.2 pts | 11.9 | no | yes |
| single-value-str | dirs | 86215 | churn | ordered | baseline | 8 | 508 | 477 | 0.92× [0.91, 0.94] | -8.3% | [-10.4%, -6.2%] | 2.0 pts | 0.9 | no | yes |
| single-value-str | dirs | 86215 | churn | ordered | btree-map | 8 | 591 | 522 | 0.90× [0.88, 0.92] | -11.2% | [-13.5%, -9.0%] | 3.0 pts | 2.2 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the pooled difference of 0.10% does not clear the 1.79% noise floor, the bound on what the harness reports between identical code in every process
- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the pooled interval [-3.01%, 3.22%] includes zero
- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.66% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.97%, 16.55%] includes zero
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 8.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 15.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 11.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 prefix: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 churn: ordered vs btree-map: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=16384 build: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 8.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesFor: ordered vs baseline: 7 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 14.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: the processes scatter 14.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 valuesBetween: ordered vs btree-map: 1 processes resolved A as faster and 7 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str dirs n=86215 prefix: ordered vs baseline: the processes scatter 9.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 prefix: ordered vs btree-map: the processes scatter 11.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=86215 prefix: ordered vs btree-map: 1 processes resolved A as faster and 7 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str dirs n=86215 churn: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
