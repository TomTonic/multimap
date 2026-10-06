| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 105 | 87.6 | 0.84× [0.82, 0.85] | -19.6% | [-21.5%, -17.7%] | 3.2 pts | 2.6 | yes | yes |
| natural-ptr | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 107 | 163 | 1.55× [1.53, 1.56] | +35.3% | [+34.6%, +36.1%] | 0.8 pts | 1.0 | yes | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2650 | 2968 | 1.13× [1.12, 1.14] | +11.6% | [+10.7%, +12.6%] | 1.1 pts | 0.6 | yes | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2675 | 5641 | 2.12× [2.10, 2.14] | +52.9% | [+52.5%, +53.3%] | 0.8 pts | 1.4 | yes | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 3023 | 3147 | 1.06× [1.05, 1.06] | +5.2% | [+4.4%, +6.1%] | 0.8 pts | 0.5 | yes | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | btree-sets | 8 | 3026 | 7652 | 2.55× [2.50, 2.60] | +60.8% | [+60.0%, +61.6%] | 0.8 pts | 0.7 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 325 | 157 | 0.48× [0.48, 0.49] | -107.3% | [-109.6%, -104.9%] | 2.3 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | btree-sets | 8 | 326 | 237 | 0.73× [0.72, 0.74] | -37.5% | [-39.9%, -35.2%] | 2.6 pts | 1.2 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | baseline | 8 | 14.79 ms | 7.33 ms | 0.50× [0.49, 0.50] | -101.2% | [-103.5%, -98.9%] | 2.2 pts | 1.2 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | btree-sets | 8 | 14.73 ms | 10.98 ms | 0.74× [0.74, 0.75] | -34.7% | [-35.6%, -33.8%] | 1.4 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 16384 | valuesFor | ordered | baseline | 6 | 136 | 122 | 0.90× [0.89, 0.91] | -11.3% | [-12.1%, -10.4%] | 0.8 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 16384 | valuesFor | ordered | btree-sets | 6 | 138 | 225 | 1.64× [1.63, 1.65] | +39.1% | [+38.7%, +39.5%] | 0.4 pts | 0.6 | yes | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 3152 | 3780 | 1.20× [1.19, 1.20] | +16.5% | [+16.1%, +17.0%] | 0.4 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | btree-sets | 6 | 3154 | 6489 | 2.06× [2.04, 2.08] | +51.4% | [+51.0%, +51.9%] | 0.4 pts | 1.0 | yes | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | baseline | 6 | 12.2 µs | 16.7 µs | 1.31× [1.27, 1.35] | +23.7% | [+21.4%, +26.0%] | 2.2 pts | 0.4 | yes | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | btree-sets | 6 | 11.9 µs | 33.0 µs | 2.78× [2.66, 2.92] | +64.1% | [+62.4%, +65.8%] | 1.6 pts | 0.6 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | baseline | 6 | 389 | 222 | 0.57× [0.56, 0.58] | -75.1% | [-77.8%, -72.4%] | 2.6 pts | 1.6 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | btree-sets | 6 | 388 | 359 | 0.92× [0.91, 0.93] | -8.7% | [-9.9%, -7.5%] | 1.2 pts | 0.9 | yes | yes |
| natural-ptr | dirs | 16384 | build | ordered | baseline | 6 | 74.74 ms | 38.90 ms | 0.53× [0.52, 0.54] | -89.3% | [-92.2%, -86.4%] | 2.8 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 16384 | build | ordered | btree-sets | 6 | 74.66 ms | 61.05 ms | 0.81× [0.80, 0.82] | -23.3% | [-24.7%, -21.9%] | 1.4 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 205 | 201 | 0.98× [0.97, 0.99] | -2.3% | [-3.5%, -1.0%] | 1.7 pts | 1.3 | yes | yes |
| natural-ptr | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 246 | 456 | 1.85× [1.83, 1.88] | +46.0% | [+45.2%, +46.8%] | 1.0 pts | 1.5 | yes | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3963 | 4830 | 1.24× [1.20, 1.27] | +19.1% | [+16.8%, +21.3%] | 2.4 pts | 3.0 | no | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 5083 | 14.1 µs | 2.79× [2.73, 2.85] | +64.1% | [+63.3%, +64.9%] | 1.0 pts | 1.7 | yes | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 89.5 µs | 121.7 µs | 1.33× [1.31, 1.35] | +24.7% | [+23.6%, +25.7%] | 1.2 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | btree-sets | 8 | 96.5 µs | 272.9 µs | 2.96× [2.86, 3.06] | +66.2% | [+65.0%, +67.3%] | 1.1 pts | 0.6 | yes | yes |
| natural-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 618 | 483 | 0.77× [0.75, 0.79] | -29.6% | [-33.1%, -26.1%] | 3.4 pts | 1.0 | no | yes |
| natural-ptr | dirs | 86215 | churn | ordered | btree-sets | 8 | 674 | 679 | 1.03× [0.99, 1.07] | +2.8% | [-1.0%, +6.5%] | 3.8 pts | 4.2 | no | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=16384 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -1.82% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 churn: ordered vs btree-sets: the pooled interval [-0.97%, 6.52%] includes zero
- natural-ptr dirs n=86215 churn: ordered vs btree-sets: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
