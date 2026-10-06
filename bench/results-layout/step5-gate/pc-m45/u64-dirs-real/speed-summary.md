| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 102 | 87.5 | 0.87× [0.83, 0.91] | -15.4% | [-20.4%, -10.3%] | 4.8 pts | 3.6 | no | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 104 | 162 | 1.56× [1.54, 1.59] | +36.1% | [+35.1%, +37.0%] | 0.9 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2601 | 3045 | 1.17× [1.14, 1.20] | +14.4% | [+12.4%, +16.4%] | 2.6 pts | 1.3 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2602 | 5458 | 2.12× [2.09, 2.14] | +52.7% | [+52.2%, +53.3%] | 0.7 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 8 | 2903 | 3339 | 1.13× [1.11, 1.15] | +11.5% | [+9.6%, +13.3%] | 1.8 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 8 | 2919 | 7532 | 2.58× [2.53, 2.62] | +61.2% | [+60.5%, +61.9%] | 0.7 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 232 | 156 | 0.67× [0.66, 0.68] | -48.9% | [-50.7%, -47.1%] | 1.8 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 232 | 233 | 1.00× [0.98, 1.02] | +0.4% | [-1.6%, +2.3%] | 2.1 pts | 1.6 | yes | no |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 10.98 ms | 7.34 ms | 0.67× [0.66, 0.67] | -49.5% | [-50.7%, -48.4%] | 1.4 pts | 1.6 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 11.04 ms | 10.78 ms | 0.98× [0.97, 0.99] | -1.8% | [-2.7%, -1.0%] | 1.0 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 128 | 119 | 0.92× [0.91, 0.94] | -8.2% | [-10.1%, -6.4%] | 2.0 pts | 1.7 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 130 | 222 | 1.72× [1.69, 1.75] | +41.8% | [+40.7%, +42.9%] | 1.1 pts | 1.9 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3003 | 3771 | 1.26× [1.24, 1.27] | +20.4% | [+19.7%, +21.2%] | 0.7 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3006 | 6201 | 2.06× [2.06, 2.07] | +51.6% | [+51.4%, +51.7%] | 0.2 pts | 0.8 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 8 | 11.8 µs | 16.8 µs | 1.44× [1.36, 1.53] | +30.4% | [+26.3%, +34.6%] | 4.0 pts | 0.7 | no | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 8 | 10.8 µs | 31.0 µs | 2.92× [2.79, 3.06] | +65.7% | [+64.1%, +67.3%] | 1.7 pts | 0.8 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 274 | 208 | 0.77× [0.74, 0.79] | -30.6% | [-35.0%, -26.3%] | 4.2 pts | 3.8 | no | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 274 | 335 | 1.22× [1.20, 1.24] | +18.0% | [+16.7%, +19.2%] | 1.2 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 52.92 ms | 38.89 ms | 0.73× [0.73, 0.74] | -36.2% | [-36.9%, -35.6%] | 0.8 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 53.48 ms | 60.08 ms | 1.12× [1.12, 1.13] | +10.9% | [+10.3%, +11.5%] | 0.6 pts | 1.0 | yes | yes |
| natural | dirs | 86215 | valuesFor | ordered | baseline | 8 | 179 | 183 | 1.03× [1.00, 1.07] | +3.2% | [-0.3%, +6.7%] | 3.4 pts | 2.4 | no | no |
| natural | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 201 | 417 | 2.10× [2.05, 2.15] | +52.3% | [+51.2%, +53.4%] | 1.2 pts | 1.8 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3447 | 4588 | 1.34× [1.28, 1.41] | +25.6% | [+22.1%, +29.1%] | 3.4 pts | 3.6 | no | yes |
| natural | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 4232 | 12.6 µs | 3.00× [2.87, 3.15] | +66.7% | [+65.2%, +68.2%] | 1.5 pts | 2.0 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | baseline | 8 | 83.6 µs | 121.1 µs | 1.43× [1.41, 1.44] | +30.0% | [+29.2%, +30.8%] | 0.8 pts | 1.0 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | btree-sets | 8 | 88.1 µs | 249.2 µs | 2.82× [2.74, 2.91] | +64.6% | [+63.5%, +65.6%] | 1.5 pts | 1.0 | yes | yes |
| natural | dirs | 86215 | churn | ordered | baseline | 8 | 422 | 413 | 0.98× [0.95, 1.01] | -2.3% | [-5.1%, +0.5%] | 2.9 pts | 0.8 | no | no |
| natural | dirs | 86215 | churn | ordered | btree-sets | 8 | 514 | 668 | 1.30× [1.28, 1.33] | +23.3% | [+22.0%, +24.7%] | 1.3 pts | 1.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=4096 churn: ordered vs btree-sets: the pooled difference of 0.38% does not clear the 1.06% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=4096 churn: ordered vs btree-sets: the pooled interval [-1.56%, 2.31%] includes zero
- natural dirs n=16384 churn: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-0.28%, 6.72%] includes zero
- natural dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 churn: ordered vs baseline: the pooled interval [-5.11%, 0.55%] includes zero
