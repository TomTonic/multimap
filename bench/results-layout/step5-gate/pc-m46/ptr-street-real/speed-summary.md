| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 78.6 | 56.2 | 0.72× [0.71, 0.72] | -39.7% | [-40.2%, -39.3%] | 1.9 pts | 1.7 | yes | yes |
| natural-ptr | street | 4096 | valuesFor | ordered | btree-sets | 8 | 79.7 | 138 | 1.74× [1.71, 1.77] | +42.5% | [+41.6%, +43.4%] | 1.1 pts | 2.3 | yes | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 2121 | 2569 | 1.23× [1.21, 1.25] | +18.5% | [+17.1%, +19.9%] | 1.7 pts | 1.1 | yes | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 2110 | 4904 | 2.33× [2.27, 2.39] | +57.0% | [+55.9%, +58.2%] | 1.2 pts | 2.8 | yes | yes |
| natural-ptr | street | 4096 | prefix | ordered | baseline | 8 | 275 | 257 | 0.93× [0.92, 0.94] | -7.3% | [-8.2%, -6.5%] | 1.0 pts | 1.0 | yes | yes |
| natural-ptr | street | 4096 | prefix | ordered | btree-sets | 8 | 277 | 555 | 2.00× [1.97, 2.03] | +50.0% | [+49.2%, +50.7%] | 0.9 pts | 2.5 | yes | yes |
| natural-ptr | street | 4096 | churn | ordered | baseline | 8 | 173 | 110 | 0.64× [0.63, 0.65] | -56.6% | [-58.2%, -55.0%] | 1.6 pts | 1.0 | yes | yes |
| natural-ptr | street | 4096 | churn | ordered | btree-sets | 8 | 174 | 202 | 1.16× [1.15, 1.17] | +13.9% | [+13.0%, +14.8%] | 1.1 pts | 1.6 | yes | yes |
| natural-ptr | street | 4096 | build | ordered | baseline | 8 | 7.27 ms | 4.34 ms | 0.60× [0.59, 0.61] | -66.8% | [-69.0%, -64.5%] | 2.1 pts | 1.6 | yes | yes |
| natural-ptr | street | 4096 | build | ordered | btree-sets | 8 | 7.21 ms | 8.00 ms | 1.10× [1.08, 1.13] | +9.3% | [+7.4%, +11.3%] | 2.1 pts | 2.1 | yes | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | baseline | 6 | 96.8 | 78.1 | 0.80× [0.79, 0.82] | -24.3% | [-26.2%, -22.3%] | 1.8 pts | 2.4 | yes | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | btree-sets | 6 | 98.5 | 192 | 1.95× [1.85, 2.05] | +48.6% | [+45.9%, +51.3%] | 2.5 pts | 5.6 | yes | yes |
| natural-ptr | street | 16384 | valuesBetween | ordered | baseline | 6 | 2519 | 3193 | 1.27× [1.26, 1.27] | +21.0% | [+20.5%, +21.4%] | 0.4 pts | 1.0 | yes | yes |
| natural-ptr | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 2519 | 5782 | 2.40× [2.12, 2.76] | +58.3% | [+52.8%, +63.7%] | 5.2 pts | 22.4 | yes | yes |
| natural-ptr | street | 16384 | prefix | ordered | baseline | 6 | 806 | 935 | 1.17× [1.16, 1.18] | +14.6% | [+13.8%, +15.5%] | 0.8 pts | 0.7 | yes | yes |
| natural-ptr | street | 16384 | prefix | ordered | btree-sets | 6 | 811 | 2177 | 2.72× [2.61, 2.84] | +63.3% | [+61.8%, +64.8%] | 1.4 pts | 2.6 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | baseline | 6 | 211 | 150 | 0.71× [0.70, 0.72] | -40.4% | [-42.5%, -38.2%] | 2.1 pts | 2.3 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | btree-sets | 6 | 214 | 285 | 1.32× [1.29, 1.37] | +24.5% | [+22.2%, +26.8%] | 2.2 pts | 2.6 | yes | yes |
| natural-ptr | street | 16384 | build | ordered | baseline | 6 | 34.86 ms | 23.99 ms | 0.68× [0.67, 0.70] | -46.5% | [-49.6%, -43.4%] | 2.9 pts | 1.4 | yes | yes |
| natural-ptr | street | 16384 | build | ordered | btree-sets | 6 | 35.08 ms | 43.57 ms | 1.25× [1.22, 1.28] | +19.9% | [+18.2%, +21.6%] | 1.6 pts | 1.7 | yes | yes |
| natural-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 171 | 190 | 1.12× [1.08, 1.16] | +10.5% | [+7.2%, +13.9%] | 7.4 pts | 8.9 | no | yes |
| natural-ptr | street | 212449 | valuesFor | ordered | btree-sets | 8 | 229 | 540 | 2.36× [2.29, 2.43] | +57.6% | [+56.3%, +58.9%] | 1.3 pts | 2.0 | yes | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 3496 | 5657 | 1.62× [1.55, 1.69] | +38.2% | [+35.5%, +40.9%] | 5.3 pts | 10.9 | yes | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 3949 | 17.1 µs | 4.30× [4.25, 4.34] | +76.7% | [+76.5%, +77.0%] | 0.4 pts | 1.7 | yes | yes |
| natural-ptr | street | 212449 | prefix | ordered | baseline | 8 | 10.9 µs | 18.3 µs | 1.71× [1.63, 1.81] | +41.6% | [+38.5%, +44.6%] | 5.8 pts | 11.9 | yes | yes |
| natural-ptr | street | 212449 | prefix | ordered | btree-sets | 8 | 13.0 µs | 58.7 µs | 4.56× [4.43, 4.71] | +78.1% | [+77.4%, +78.8%] | 0.8 pts | 2.7 | yes | yes |
| natural-ptr | street | 212449 | churn | ordered | baseline | 8 | 459 | 471 | 0.99× [0.97, 1.02] | -0.7% | [-3.0%, +1.7%] | 3.3 pts | 1.4 | no | no |
| natural-ptr | street | 212449 | churn | ordered | btree-sets | 8 | 547 | 752 | 1.39× [1.37, 1.41] | +28.1% | [+27.1%, +29.0%] | 1.0 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=4096 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=4096 prefix: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=4096 build: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 22.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 prefix: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 churn: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 8.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 10.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 prefix: ordered vs baseline: the processes scatter 11.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +1.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=212449 prefix: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 churn: ordered vs baseline: the pooled difference of -0.68% does not clear the 2.75% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr street n=212449 churn: ordered vs baseline: the pooled interval [-3.04%, 1.68%] includes zero
