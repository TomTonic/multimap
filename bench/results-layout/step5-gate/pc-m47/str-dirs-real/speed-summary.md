| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 140 | 126 | 0.90× [0.89, 0.91] | -10.9% | [-12.1%, -9.6%] | 1.6 pts | 1.3 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 142 | 177 | 1.25× [1.19, 1.32] | +20.2% | [+16.2%, +24.3%] | 4.2 pts | 2.7 | no | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6686 | 6652 | 0.99× [0.98, 1.00] | -0.9% | [-2.1%, +0.2%] | 1.3 pts | 1.2 | yes | no |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6727 | 11.1 µs | 1.66× [1.61, 1.71] | +39.7% | [+38.1%, +41.4%] | 1.7 pts | 2.5 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 7524 | 7587 | 1.01× [0.99, 1.03] | +1.0% | [-1.0%, +2.9%] | 2.0 pts | 0.7 | yes | no |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7550 | 16.0 µs | 2.06× [1.96, 2.18] | +51.5% | [+48.9%, +54.2%] | 2.6 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 271 | 196 | 0.72× [0.71, 0.73] | -38.5% | [-40.1%, -36.9%] | 1.7 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 273 | 261 | 0.94× [0.93, 0.96] | -5.9% | [-7.5%, -4.3%] | 1.6 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 14.08 ms | 9.43 ms | 0.67× [0.66, 0.68] | -49.4% | [-52.6%, -46.1%] | 3.5 pts | 1.5 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 14.15 ms | 13.23 ms | 0.94× [0.92, 0.96] | -6.2% | [-8.3%, -4.1%] | 2.3 pts | 1.8 | no | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 178 | 170 | 0.96× [0.88, 1.06] | -4.0% | [-13.8%, +5.9%] | 10.0 pts | 11.1 | no | no |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 189 | 301 | 1.61× [1.47, 1.78] | +37.8% | [+31.8%, +43.8%] | 8.4 pts | 8.2 | no | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 7356 | 7785 | 1.07× [0.97, 1.19] | +6.3% | [-3.3%, +15.8%] | 9.6 pts | 10.8 | no | no |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 7603 | 15.8 µs | 2.06× [1.83, 2.35] | +51.4% | [+45.4%, +57.4%] | 7.9 pts | 11.5 | no | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 8 | 33.7 µs | 38.1 µs | 1.11× [1.07, 1.15] | +10.0% | [+6.9%, +13.2%] | 3.0 pts | 4.0 | no | no |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 8 | 32.1 µs | 68.4 µs | 2.15× [2.05, 2.25] | +53.4% | [+51.2%, +55.6%] | 5.6 pts | 11.0 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 8 | 358 | 287 | 0.81× [0.79, 0.83] | -23.7% | [-27.1%, -20.3%] | 3.4 pts | 3.4 | no | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 8 | 387 | 422 | 1.08× [1.04, 1.12] | +7.2% | [+4.0%, +10.3%] | 3.1 pts | 3.7 | no | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 8 | 72.77 ms | 56.74 ms | 0.77× [0.76, 0.78] | -30.1% | [-32.4%, -27.8%] | 2.9 pts | 1.9 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 8 | 73.65 ms | 77.27 ms | 1.03× [1.00, 1.06] | +2.9% | [+0.4%, +5.4%] | 2.5 pts | 2.8 | no | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 264 | 282 | 1.05× [1.01, 1.09] | +4.9% | [+1.2%, +8.6%] | 4.0 pts | 4.1 | no | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 297 | 506 | 1.75× [1.67, 1.84] | +42.9% | [+40.2%, +45.6%] | 2.7 pts | 4.1 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 10.1 µs | 11.4 µs | 1.16× [1.12, 1.20] | +13.6% | [+10.8%, +16.4%] | 3.3 pts | 2.7 | no | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 10.5 µs | 23.5 µs | 2.23× [2.13, 2.34] | +55.2% | [+53.1%, +57.2%] | 2.0 pts | 3.2 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 228.9 µs | 289.0 µs | 1.22× [1.18, 1.26] | +17.8% | [+15.0%, +20.5%] | 4.4 pts | 1.5 | no | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 233.6 µs | 598.2 µs | 2.47× [2.31, 2.66] | +59.5% | [+56.7%, +62.4%] | 2.8 pts | 1.3 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 592 | 552 | 0.90× [0.85, 0.95] | -11.2% | [-17.4%, -4.9%] | 8.0 pts | 2.9 | no | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 617 | 704 | 1.15× [1.14, 1.17] | +13.2% | [+12.1%, +14.3%] | 1.5 pts | 1.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.08%, 0.19%] includes zero
- natural-str dirs n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 prefix: ordered vs baseline: the pooled difference of 0.96% does not clear the 4.82% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=4096 prefix: ordered vs baseline: the pooled interval [-1.01%, 2.93%] includes zero
- natural-str dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-13.84%, 5.87%] includes zero
- natural-str dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 11.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-str dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 8.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs baseline: the pooled interval [-3.28%, 15.83%] includes zero
- natural-str dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 10.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs baseline: 6 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-str dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 11.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 prefix: ordered vs baseline: the pooled difference of 10.01% does not clear the 12.99% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of -6.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 prefix: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 prefix: ordered vs btree-sets: the processes scatter 11.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 churn: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 churn: ordered vs btree-sets: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 build: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.37% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -2.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=86215 churn: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
