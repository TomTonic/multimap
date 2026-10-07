| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 96.1 | 86.4 | 0.90× [0.88, 0.91] | -11.7% | [-14.0%, -9.5%] | 2.4 pts | 2.1 | no | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 97.0 | 165 | 1.73× [1.70, 1.77] | +42.3% | [+41.2%, +43.4%] | 1.8 pts | 2.1 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2631 | 3125 | 1.20× [1.17, 1.23] | +16.9% | [+14.7%, +19.0%] | 2.1 pts | 1.2 | no | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2640 | 5610 | 2.16× [2.13, 2.19] | +53.6% | [+53.0%, +54.3%] | 2.0 pts | 4.0 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 8 | 2987 | 3447 | 1.15× [1.11, 1.18] | +12.9% | [+10.2%, +15.5%] | 2.5 pts | 1.4 | no | yes |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 8 | 2999 | 7759 | 2.59× [2.52, 2.68] | +61.4% | [+60.3%, +62.6%] | 1.7 pts | 2.3 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 223 | 162 | 0.72× [0.72, 0.73] | -38.5% | [-39.7%, -37.2%] | 2.1 pts | 1.3 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 225 | 244 | 1.09× [1.07, 1.10] | +7.9% | [+6.5%, +9.3%] | 2.1 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 10.69 ms | 7.80 ms | 0.73× [0.72, 0.74] | -36.3% | [-38.3%, -34.3%] | 3.1 pts | 1.8 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 10.91 ms | 11.93 ms | 1.09× [1.06, 1.11] | +8.1% | [+6.0%, +10.2%] | 2.8 pts | 2.1 | no | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 125 | 120 | 1.02× [0.91, 1.16] | +1.7% | [-10.4%, +13.7%] | 11.9 pts | 10.4 | no | no |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 129 | 271 | 2.15× [1.88, 2.49] | +53.4% | [+46.9%, +59.9%] | 6.9 pts | 6.1 | no | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3037 | 3869 | 1.38× [1.18, 1.64] | +27.3% | [+15.5%, +39.1%] | 11.5 pts | 20.0 | no | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3146 | 8966 | 2.81× [2.35, 3.51] | +64.5% | [+57.4%, +71.5%] | 8.0 pts | 13.5 | no | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 8 | 12.0 µs | 17.6 µs | 1.54× [1.40, 1.72] | +35.2% | [+28.6%, +41.8%] | 6.2 pts | 1.2 | no | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 8 | 11.6 µs | 36.6 µs | 3.30× [2.77, 4.10] | +69.7% | [+63.9%, +75.6%] | 6.7 pts | 3.1 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 283 | 232 | 0.82× [0.80, 0.84] | -21.6% | [-24.8%, -18.4%] | 3.1 pts | 2.5 | no | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 308 | 384 | 1.24× [1.19, 1.29] | +19.2% | [+16.0%, +22.4%] | 3.2 pts | 4.3 | no | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 57.16 ms | 47.57 ms | 0.81× [0.79, 0.84] | -23.0% | [-26.9%, -19.1%] | 3.7 pts | 2.6 | no | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 57.35 ms | 73.73 ms | 1.24× [1.20, 1.28] | +19.4% | [+16.8%, +22.1%] | 3.4 pts | 4.7 | no | yes |
| natural | dirs | 86215 | valuesFor | ordered | baseline | 8 | 174 | 181 | 1.03× [0.99, 1.08] | +3.4% | [-0.5%, +7.3%] | 4.2 pts | 3.5 | no | no |
| natural | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 196 | 430 | 2.20× [2.07, 2.34] | +54.5% | [+51.8%, +57.3%] | 2.7 pts | 4.2 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3502 | 4815 | 1.38× [1.27, 1.51] | +27.6% | [+21.5%, +33.7%] | 6.2 pts | 8.7 | no | yes |
| natural | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 4031 | 13.3 µs | 3.23× [2.98, 3.52] | +69.0% | [+66.5%, +71.6%] | 2.7 pts | 5.9 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | baseline | 8 | 84.4 µs | 120.3 µs | 1.49× [1.41, 1.58] | +33.0% | [+29.0%, +36.9%] | 3.8 pts | 6.5 | no | yes |
| natural | dirs | 86215 | prefix | ordered | btree-sets | 8 | 89.3 µs | 311.9 µs | 3.46× [3.08, 3.94] | +71.1% | [+67.5%, +74.6%] | 4.3 pts | 6.5 | yes | yes |
| natural | dirs | 86215 | churn | ordered | baseline | 8 | 472 | 459 | 0.98× [0.95, 1.01] | -2.5% | [-5.5%, +0.6%] | 2.8 pts | 1.4 | no | no |
| natural | dirs | 86215 | churn | ordered | btree-sets | 8 | 525 | 701 | 1.33× [1.29, 1.36] | +24.6% | [+22.6%, +26.5%] | 2.7 pts | 4.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=4096 prefix: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=4096 build: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-10.37%, 13.71%] includes zero
- natural dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 10.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural dirs n=16384 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.13% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 6.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 20.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.83% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 13.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 prefix: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs btree-sets: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-0.55%, 7.26%] includes zero
- natural dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 8.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 5.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -1.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=86215 prefix: ordered vs baseline: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 prefix: ordered vs btree-sets: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 churn: ordered vs baseline: the pooled interval [-5.52%, 0.56%] includes zero
- natural dirs n=86215 churn: ordered vs btree-sets: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
