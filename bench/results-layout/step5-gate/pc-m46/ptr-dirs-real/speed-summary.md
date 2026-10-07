| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 111 | 89.2 | 0.80× [0.79, 0.81] | -25.6% | [-27.2%, -23.9%] | 1.6 pts | 1.2 | yes | yes |
| natural-ptr | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 113 | 168 | 1.51× [1.49, 1.53] | +33.6% | [+32.8%, +34.4%] | 3.3 pts | 3.1 | yes | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2948 | 3075 | 1.04× [1.02, 1.07] | +4.1% | [+2.0%, +6.2%] | 2.8 pts | 1.4 | no | yes |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2986 | 5904 | 2.02× [1.99, 2.05] | +50.5% | [+49.7%, +51.3%] | 1.9 pts | 3.4 | yes | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 3359 | 3305 | 0.98× [0.96, 1.01] | -1.6% | [-4.1%, +1.0%] | 2.5 pts | 1.4 | no | no |
| natural-ptr | dirs | 4096 | prefix | ordered | btree-sets | 8 | 3330 | 8122 | 2.46× [2.38, 2.56] | +59.4% | [+57.9%, +60.9%] | 2.0 pts | 2.3 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 237 | 163 | 0.68× [0.67, 0.69] | -46.8% | [-48.3%, -45.3%] | 1.8 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | btree-sets | 8 | 238 | 247 | 1.03× [1.02, 1.04] | +3.0% | [+2.0%, +3.9%] | 1.1 pts | 0.8 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | baseline | 8 | 11.66 ms | 7.63 ms | 0.65× [0.65, 0.66] | -53.1% | [-54.4%, -51.8%] | 1.9 pts | 1.6 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | btree-sets | 8 | 11.88 ms | 12.04 ms | 1.02× [1.00, 1.05] | +2.1% | [-0.3%, +4.6%] | 2.8 pts | 2.8 | no | no |
| natural-ptr | dirs | 16384 | valuesFor | ordered | baseline | 8 | 143 | 124 | 0.95× [0.85, 1.08] | -5.1% | [-17.7%, +7.6%] | 14.6 pts | 17.2 | no | no |
| natural-ptr | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 144 | 239 | 1.72× [1.64, 1.82] | +42.0% | [+38.9%, +45.0%] | 5.6 pts | 6.6 | yes | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3353 | 3854 | 1.29× [1.12, 1.51] | +22.4% | [+11.1%, +33.8%] | 13.7 pts | 20.2 | no | yes |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3401 | 7235 | 2.23× [2.06, 2.44] | +55.2% | [+51.4%, +59.1%] | 8.0 pts | 21.1 | yes | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | baseline | 8 | 13.0 µs | 16.7 µs | 1.31× [1.22, 1.41] | +23.7% | [+18.2%, +29.3%] | 6.0 pts | 1.0 | no | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | btree-sets | 8 | 14.3 µs | 38.2 µs | 2.73× [2.66, 2.82] | +63.4% | [+62.4%, +64.5%] | 4.2 pts | 3.2 | yes | yes |
| natural-ptr | dirs | 16384 | churn | ordered | baseline | 8 | 284 | 221 | 0.78× [0.76, 0.81] | -27.9% | [-31.7%, -24.1%] | 3.8 pts | 2.6 | no | yes |
| natural-ptr | dirs | 16384 | churn | ordered | btree-sets | 8 | 312 | 383 | 1.20× [1.17, 1.25] | +17.0% | [+14.3%, +19.7%] | 3.6 pts | 3.5 | no | yes |
| natural-ptr | dirs | 16384 | build | ordered | baseline | 8 | 60.52 ms | 43.29 ms | 0.73× [0.70, 0.76] | -37.0% | [-42.3%, -31.7%] | 5.6 pts | 3.0 | no | yes |
| natural-ptr | dirs | 16384 | build | ordered | btree-sets | 8 | 59.70 ms | 68.06 ms | 1.15× [1.11, 1.20] | +13.1% | [+9.5%, +16.6%] | 4.5 pts | 5.5 | no | yes |
| natural-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 208 | 214 | 1.03× [0.96, 1.12] | +3.3% | [-4.1%, +10.8%] | 8.2 pts | 9.7 | no | no |
| natural-ptr | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 235 | 454 | 1.95× [1.85, 2.06] | +48.6% | [+45.9%, +51.3%] | 2.6 pts | 3.0 | yes | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 4309 | 6172 | 1.39× [1.26, 1.56] | +28.3% | [+20.8%, +35.7%] | 8.8 pts | 11.7 | no | yes |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 5106 | 15.2 µs | 2.95× [2.74, 3.20] | +66.1% | [+63.5%, +68.7%] | 3.2 pts | 6.7 | yes | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 96.6 µs | 146.5 µs | 1.52× [1.35, 1.75] | +34.2% | [+25.7%, +42.8%] | 9.7 pts | 10.8 | no | yes |
| natural-ptr | dirs | 86215 | prefix | ordered | btree-sets | 8 | 102.6 µs | 332.0 µs | 3.31× [2.88, 3.88] | +69.8% | [+65.3%, +74.2%] | 4.8 pts | 4.1 | yes | yes |
| natural-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 476 | 458 | 0.95× [0.92, 0.99] | -5.0% | [-8.9%, -1.2%] | 6.7 pts | 2.9 | no | yes |
| natural-ptr | dirs | 86215 | churn | ordered | btree-sets | 8 | 523 | 672 | 1.28× [1.25, 1.32] | +22.2% | [+20.1%, +24.2%] | 2.1 pts | 2.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr dirs n=4096 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +0.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr dirs n=4096 valuesFor: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 prefix: ordered vs baseline: the pooled interval [-4.10%, 0.97%] includes zero
- natural-ptr dirs n=4096 prefix: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 build: ordered vs btree-sets: the pooled interval [-0.32%, 4.57%] includes zero
- natural-ptr dirs n=4096 build: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=4096 build: ordered vs btree-sets: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-17.69%, 7.59%] includes zero
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 17.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 20.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 21.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 prefix: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 churn: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 build: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=16384 build: ordered vs btree-sets: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-4.13%, 10.83%] includes zero
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 9.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 11.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 prefix: ordered vs baseline: the processes scatter 10.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 prefix: ordered vs btree-sets: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 churn: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 churn: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
