| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 88.3 | 85.2 | 0.97× [0.95, 0.98] | -3.3% | [-4.7%, -1.9%] | 1.5 pts | 1.5 | yes | yes |
| natural-ptr | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 89.8 | 164 | 1.85× [1.82, 1.87] | +45.8% | [+45.0%, +46.6%] | 0.8 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 3146 | 2817 | 0.90× [0.89, 0.92] | -10.7% | [-12.5%, -8.9%] | 2.3 pts | 0.9 | yes | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 3150 | 5667 | 1.81× [1.79, 1.83] | +44.8% | [+44.2%, +45.5%] | 0.6 pts | 0.9 | yes | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 3453 | 2917 | 0.85× [0.84, 0.87] | -17.5% | [-19.8%, -15.2%] | 2.5 pts | 1.4 | no | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | btree-sets | 8 | 3451 | 7675 | 2.24× [2.17, 2.31] | +55.3% | [+53.9%, +56.7%] | 1.3 pts | 1.3 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 153 | 148 | 0.98× [0.96, 0.99] | -2.3% | [-3.9%, -0.6%] | 1.7 pts | 1.4 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | btree-sets | 8 | 154 | 229 | 1.48× [1.45, 1.52] | +32.5% | [+30.9%, +34.0%] | 1.5 pts | 1.5 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | baseline | 8 | 7.35 ms | 7.12 ms | 0.97× [0.96, 0.98] | -3.1% | [-4.1%, -2.2%] | 1.0 pts | 1.4 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | btree-sets | 8 | 7.41 ms | 11.23 ms | 1.51× [1.50, 1.52] | +33.8% | [+33.2%, +34.4%] | 0.5 pts | 0.7 | yes | yes |
| natural-ptr | dirs | 16384 | valuesFor | ordered | baseline | 8 | 121 | 117 | 0.96× [0.95, 0.97] | -3.9% | [-5.1%, -2.7%] | 1.2 pts | 1.5 | yes | yes |
| natural-ptr | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 124 | 227 | 1.83× [1.81, 1.85] | +45.4% | [+44.8%, +46.0%] | 0.6 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3867 | 3602 | 0.93× [0.92, 0.94] | -7.5% | [-8.2%, -6.7%] | 0.7 pts | 0.9 | yes | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3904 | 6497 | 1.66× [1.65, 1.67] | +39.8% | [+39.4%, +40.3%] | 0.6 pts | 1.8 | yes | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | baseline | 8 | 18.2 µs | 15.9 µs | 0.91× [0.89, 0.93] | -10.5% | [-13.0%, -7.9%] | 2.6 pts | 1.0 | no | no |
| natural-ptr | dirs | 16384 | prefix | ordered | btree-sets | 8 | 18.2 µs | 35.7 µs | 1.91× [1.88, 1.94] | +47.7% | [+46.9%, +48.4%] | 0.8 pts | 1.3 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | baseline | 8 | 205 | 200 | 0.97× [0.97, 0.98] | -2.7% | [-3.4%, -1.9%] | 0.8 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | btree-sets | 8 | 210 | 330 | 1.56× [1.54, 1.58] | +36.0% | [+35.2%, +36.8%] | 0.9 pts | 1.3 | yes | yes |
| natural-ptr | dirs | 16384 | build | ordered | baseline | 8 | 39.23 ms | 38.27 ms | 0.97× [0.97, 0.98] | -2.6% | [-3.1%, -2.2%] | 0.5 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 16384 | build | ordered | btree-sets | 8 | 39.51 ms | 60.84 ms | 1.54× [1.53, 1.55] | +35.0% | [+34.6%, +35.4%] | 0.4 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 195 | 197 | 0.99× [0.96, 1.02] | -1.2% | [-4.7%, +2.2%] | 3.2 pts | 3.1 | no | no |
| natural-ptr | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 263 | 455 | 1.76× [1.73, 1.80] | +43.3% | [+42.2%, +44.4%] | 1.8 pts | 2.4 | yes | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 5298 | 5047 | 0.97× [0.94, 1.00] | -3.2% | [-6.3%, -0.1%] | 3.0 pts | 2.3 | no | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 7387 | 15.2 µs | 2.08× [2.03, 2.13] | +51.9% | [+50.7%, +53.0%] | 1.2 pts | 1.2 | yes | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 125.0 µs | 120.2 µs | 0.94× [0.91, 0.97] | -6.3% | [-9.7%, -2.9%] | 3.2 pts | 1.1 | no | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | btree-sets | 8 | 127.0 µs | 282.5 µs | 2.16× [2.09, 2.22] | +53.6% | [+52.2%, +55.0%] | 1.9 pts | 0.5 | yes | yes |
| natural-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 402 | 395 | 0.97× [0.94, 1.01] | -3.1% | [-6.9%, +0.8%] | 3.6 pts | 1.3 | no | no |
| natural-ptr | dirs | 86215 | churn | ordered | btree-sets | 8 | 523 | 657 | 1.26× [1.24, 1.29] | +20.7% | [+19.1%, +22.3%] | 1.5 pts | 1.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr dirs n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=16384 prefix: ordered vs baseline: the pooled difference of -10.45% does not clear the 12.21% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of -7.01% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-4.68%, 2.21%] includes zero
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 churn: ordered vs baseline: the pooled interval [-6.90%, 0.79%] includes zero
