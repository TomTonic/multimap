| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 146 | 126 | 0.86× [0.85, 0.87] | -16.5% | [-17.8%, -15.2%] | 1.4 pts | 1.2 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 149 | 171 | 1.16× [1.14, 1.17] | +13.5% | [+12.2%, +14.8%] | 1.3 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6872 | 6698 | 0.98× [0.97, 0.98] | -2.5% | [-3.3%, -1.8%] | 0.7 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6906 | 10.5 µs | 1.52× [1.52, 1.53] | +34.4% | [+34.1%, +34.7%] | 0.4 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 8070 | 7639 | 0.96× [0.94, 0.99] | -3.7% | [-6.4%, -1.0%] | 4.0 pts | 1.2 | no | yes |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7967 | 15.2 µs | 1.89× [1.84, 1.95] | +47.2% | [+45.5%, +48.8%] | 1.6 pts | 0.6 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 267 | 194 | 0.73× [0.72, 0.74] | -37.1% | [-38.6%, -35.6%] | 1.6 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 267 | 250 | 0.94× [0.93, 0.96] | -6.0% | [-7.6%, -4.4%] | 2.1 pts | 1.3 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 13.44 ms | 8.91 ms | 0.66× [0.65, 0.66] | -52.3% | [-54.1%, -50.6%] | 1.8 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 13.53 ms | 12.57 ms | 0.92× [0.90, 0.94] | -8.9% | [-11.6%, -6.2%] | 2.5 pts | 1.8 | no | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 187 | 169 | 0.89× [0.88, 0.91] | -11.9% | [-13.7%, -10.0%] | 2.8 pts | 3.4 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 196 | 267 | 1.40× [1.27, 1.55] | +28.5% | [+21.5%, +35.5%] | 6.8 pts | 5.5 | no | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 7561 | 7646 | 1.01× [0.98, 1.03] | +0.6% | [-1.8%, +3.1%] | 2.4 pts | 2.2 | no | no |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 7833 | 13.6 µs | 1.76× [1.55, 2.03] | +43.1% | [+35.4%, +50.7%] | 7.7 pts | 8.8 | no | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 8 | 34.6 µs | 38.2 µs | 1.08× [1.07, 1.09] | +7.3% | [+6.6%, +7.9%] | 0.8 pts | 1.2 | yes | no |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 8 | 33.5 µs | 63.5 µs | 1.94× [1.80, 2.09] | +48.4% | [+44.6%, +52.2%] | 3.7 pts | 9.1 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 8 | 333 | 267 | 0.80× [0.79, 0.81] | -24.7% | [-26.3%, -23.1%] | 1.6 pts | 1.6 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 8 | 347 | 384 | 1.11× [1.09, 1.14] | +10.1% | [+8.3%, +11.9%] | 1.9 pts | 2.3 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 8 | 67.08 ms | 50.66 ms | 0.75× [0.73, 0.77] | -33.2% | [-36.3%, -30.1%] | 3.3 pts | 1.8 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 8 | 68.20 ms | 71.37 ms | 1.05× [1.04, 1.06] | +4.9% | [+3.8%, +6.0%] | 1.3 pts | 1.2 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 268 | 273 | 1.00× [0.97, 1.04] | +0.4% | [-2.8%, +3.5%] | 3.6 pts | 3.8 | no | no |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 294 | 507 | 1.68× [1.62, 1.75] | +40.5% | [+38.3%, +42.7%] | 2.4 pts | 3.4 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 10.3 µs | 11.7 µs | 1.13× [1.09, 1.17] | +11.4% | [+8.2%, +14.6%] | 3.2 pts | 2.5 | no | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 10.6 µs | 22.6 µs | 2.14× [2.09, 2.19] | +53.2% | [+52.2%, +54.3%] | 1.1 pts | 1.6 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 239.0 µs | 284.6 µs | 1.17× [1.12, 1.23] | +14.8% | [+10.8%, +18.9%] | 4.7 pts | 1.6 | no | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 257.6 µs | 595.7 µs | 2.34× [2.21, 2.49] | +57.3% | [+54.8%, +59.9%] | 2.7 pts | 1.2 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 580 | 543 | 0.95× [0.92, 0.97] | -5.8% | [-8.6%, -2.9%] | 2.9 pts | 1.4 | no | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 622 | 719 | 1.15× [1.14, 1.16] | +13.3% | [+12.7%, +13.9%] | 1.1 pts | 1.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -0.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.85%, 3.05%] includes zero
- natural-str dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-str dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 8.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 prefix: ordered vs baseline: the pooled difference of 7.26% does not clear the 14.93% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of -9.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 prefix: ordered vs btree-sets: the processes scatter 9.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the pooled difference of 0.37% does not clear the 0.74% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-2.81%, 3.55%] includes zero
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
