| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 108 | 87.5 | 0.82× [0.80, 0.83] | -22.6% | [-25.2%, -20.1%] | 2.6 pts | 2.1 | no | yes |
| natural-ptr | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 110 | 164 | 1.51× [1.48, 1.53] | +33.6% | [+32.6%, +34.6%] | 1.2 pts | 1.4 | yes | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2688 | 3012 | 1.12× [1.09, 1.16] | +11.0% | [+8.3%, +13.6%] | 2.5 pts | 1.3 | no | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2697 | 5688 | 2.11× [2.09, 2.14] | +52.7% | [+52.2%, +53.2%] | 0.5 pts | 0.9 | yes | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 3019 | 3251 | 1.07× [1.06, 1.08] | +6.5% | [+5.5%, +7.4%] | 1.0 pts | 0.7 | yes | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | btree-sets | 8 | 3032 | 7785 | 2.56× [2.47, 2.66] | +61.0% | [+59.5%, +62.4%] | 1.5 pts | 1.5 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 327 | 160 | 0.49× [0.48, 0.49] | -104.9% | [-107.6%, -102.1%] | 2.7 pts | 0.9 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | btree-sets | 8 | 327 | 238 | 0.73× [0.70, 0.76] | -37.2% | [-42.1%, -32.3%] | 4.8 pts | 2.1 | no | yes |
| natural-ptr | dirs | 4096 | build | ordered | baseline | 8 | 14.75 ms | 7.34 ms | 0.50× [0.50, 0.50] | -100.1% | [-101.1%, -99.1%] | 1.2 pts | 0.7 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | btree-sets | 8 | 14.68 ms | 11.01 ms | 0.75× [0.75, 0.75] | -33.4% | [-34.2%, -32.6%] | 1.0 pts | 0.7 | yes | yes |
| natural-ptr | dirs | 16384 | valuesFor | ordered | baseline | 8 | 138 | 122 | 0.88× [0.87, 0.89] | -13.6% | [-14.4%, -12.8%] | 0.8 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 141 | 228 | 1.63× [1.61, 1.65] | +38.8% | [+38.1%, +39.5%] | 0.7 pts | 1.0 | yes | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3171 | 3805 | 1.20× [1.19, 1.21] | +16.9% | [+16.3%, +17.5%] | 0.6 pts | 1.0 | yes | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3185 | 6516 | 2.05× [2.04, 2.07] | +51.2% | [+50.9%, +51.6%] | 0.4 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | baseline | 8 | 12.4 µs | 16.8 µs | 1.37× [1.32, 1.43] | +27.2% | [+24.4%, +29.9%] | 2.7 pts | 0.5 | no | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | btree-sets | 8 | 12.0 µs | 33.5 µs | 2.88× [2.76, 3.00] | +65.3% | [+63.8%, +66.7%] | 1.5 pts | 0.6 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | baseline | 8 | 392 | 232 | 0.58× [0.57, 0.58] | -72.7% | [-74.1%, -71.2%] | 2.3 pts | 1.2 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | btree-sets | 8 | 411 | 378 | 0.91× [0.90, 0.93] | -9.5% | [-11.4%, -7.6%] | 2.0 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 16384 | build | ordered | baseline | 8 | 75.15 ms | 39.84 ms | 0.53× [0.53, 0.54] | -87.8% | [-89.6%, -86.0%] | 1.7 pts | 0.6 | yes | yes |
| natural-ptr | dirs | 16384 | build | ordered | btree-sets | 8 | 75.41 ms | 60.86 ms | 0.81× [0.80, 0.82] | -23.4% | [-25.1%, -21.8%] | 1.6 pts | 0.9 | yes | yes |
| natural-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 218 | 229 | 1.03× [1.00, 1.06] | +3.0% | [+0.2%, +5.9%] | 3.7 pts | 3.2 | no | yes |
| natural-ptr | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 248 | 463 | 1.86× [1.83, 1.89] | +46.2% | [+45.2%, +47.2%] | 0.9 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3926 | 4829 | 1.25× [1.20, 1.31] | +20.1% | [+16.6%, +23.6%] | 3.3 pts | 3.2 | no | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 5291 | 14.7 µs | 2.80× [2.70, 2.90] | +64.3% | [+63.0%, +65.5%] | 1.2 pts | 2.0 | yes | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 89.1 µs | 124.5 µs | 1.36× [1.33, 1.38] | +26.2% | [+24.9%, +27.6%] | 1.5 pts | 1.3 | yes | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | btree-sets | 8 | 100.8 µs | 280.4 µs | 3.03× [2.94, 3.12] | +67.0% | [+66.0%, +68.0%] | 1.1 pts | 0.6 | yes | yes |
| natural-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 627 | 493 | 0.78× [0.75, 0.82] | -27.6% | [-32.9%, -22.3%] | 5.6 pts | 1.5 | no | yes |
| natural-ptr | dirs | 86215 | churn | ordered | btree-sets | 8 | 660 | 680 | 1.03× [1.01, 1.05] | +2.8% | [+0.7%, +5.0%] | 2.8 pts | 3.1 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -1.60% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=86215 churn: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
