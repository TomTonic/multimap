| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 6 | 83.9 | 87.0 | 1.04× [1.03, 1.05] | +3.9% | [+3.2%, +4.5%] | 0.6 pts | 0.5 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 6 | 84.7 | 160 | 1.92× [1.89, 1.95] | +47.8% | [+47.0%, +48.6%] | 0.7 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 6 | 165 | 158 | 0.95× [0.94, 0.96] | -4.8% | [-5.8%, -3.8%] | 0.9 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 6 | 165 | 235 | 1.43× [1.41, 1.46] | +30.3% | [+29.0%, +31.6%] | 1.3 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 6 | 7.57 ms | 7.32 ms | 0.97× [0.96, 0.98] | -3.1% | [-3.8%, -2.4%] | 0.7 pts | 1.9 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 6 | 7.85 ms | 11.25 ms | 1.45× [1.41, 1.49] | +31.1% | [+29.2%, +32.9%] | 1.8 pts | 1.6 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 6 | 115 | 119 | 1.04× [1.02, 1.06] | +3.4% | [+1.6%, +5.3%] | 1.8 pts | 1.7 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 6 | 117 | 223 | 1.92× [1.89, 1.95] | +47.9% | [+47.0%, +48.8%] | 0.8 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 6 | 218 | 211 | 0.97× [0.96, 0.97] | -3.5% | [-4.1%, -2.9%] | 0.6 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 6 | 222 | 335 | 1.50× [1.48, 1.52] | +33.3% | [+32.4%, +34.2%] | 0.8 pts | 0.9 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 6 | 39.70 ms | 38.66 ms | 0.97× [0.96, 0.98] | -2.8% | [-3.8%, -1.8%] | 1.0 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 6 | 40.79 ms | 61.91 ms | 1.52× [1.50, 1.54] | +34.3% | [+33.4%, +35.1%] | 0.8 pts | 1.0 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 6 | 53.0 | 54.9 | 1.04× [1.02, 1.06] | +3.7% | [+2.0%, +5.3%] | 1.6 pts | 2.0 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 6 | 53.5 | 135 | 2.52× [2.50, 2.55] | +60.4% | [+60.0%, +60.8%] | 0.4 pts | 1.3 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 6 | 112 | 109 | 0.97× [0.96, 0.98] | -3.0% | [-4.4%, -1.6%] | 1.4 pts | 0.8 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 6 | 114 | 196 | 1.73× [1.71, 1.75] | +42.2% | [+41.5%, +42.9%] | 0.7 pts | 0.8 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 6 | 4.21 ms | 4.14 ms | 0.99× [0.98, 1.00] | -1.2% | [-1.9%, -0.5%] | 0.7 pts | 1.4 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 6 | 4.35 ms | 7.77 ms | 1.78× [1.72, 1.84] | +43.8% | [+41.7%, +45.8%] | 1.9 pts | 1.7 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 72.3 | 75.9 | 1.04× [1.01, 1.07] | +4.0% | [+1.3%, +6.7%] | 2.5 pts | 2.7 | no | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 73.1 | 181 | 2.48× [2.42, 2.53] | +59.6% | [+58.7%, +60.5%] | 1.0 pts | 3.1 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 151 | 147 | 0.98× [0.97, 0.99] | -2.5% | [-3.5%, -1.5%] | 0.9 pts | 1.0 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 152 | 268 | 1.77× [1.75, 1.79] | +43.6% | [+42.9%, +44.2%] | 0.7 pts | 1.5 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 22.45 ms | 22.32 ms | 0.99× [0.98, 1.00] | -0.9% | [-2.0%, +0.2%] | 1.2 pts | 1.7 | yes | no |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 23.56 ms | 42.72 ms | 1.81× [1.79, 1.83] | +44.7% | [+44.1%, +45.2%] | 1.8 pts | 1.8 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 6 | 37.9 | 39.0 | 1.02× [1.01, 1.04] | +2.3% | [+0.9%, +3.8%] | 1.4 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 6 | 37.5 | 156 | 4.18× [4.14, 4.21] | +76.1% | [+75.9%, +76.3%] | 0.2 pts | 1.0 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 6 | 65.3 | 59.1 | 0.91× [0.90, 0.92] | -10.0% | [-10.7%, -9.2%] | 0.7 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 6 | 65.1 | 160 | 2.46× [2.41, 2.51] | +59.3% | [+58.4%, +60.2%] | 0.8 pts | 2.2 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 6 | 6.46 ms | 5.90 ms | 0.92× [0.91, 0.92] | -9.2% | [-9.7%, -8.7%] | 0.5 pts | 1.5 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 6 | 6.49 ms | 15.11 ms | 2.34× [2.32, 2.37] | +57.3% | [+56.8%, +57.8%] | 0.5 pts | 1.3 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 6 | 46.7 | 48.0 | 1.03× [1.02, 1.04] | +2.7% | [+1.7%, +3.7%] | 1.0 pts | 1.9 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 46.5 | 192 | 4.16× [4.13, 4.19] | +75.9% | [+75.8%, +76.1%] | 0.2 pts | 0.9 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 6 | 85.2 | 77.6 | 0.92× [0.91, 0.92] | -9.0% | [-9.9%, -8.2%] | 0.8 pts | 1.1 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 6 | 87.5 | 230 | 2.59× [2.56, 2.63] | +61.5% | [+60.9%, +62.0%] | 0.6 pts | 1.3 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 6 | 31.62 ms | 29.12 ms | 0.93× [0.92, 0.93] | -7.9% | [-8.9%, -7.0%] | 0.9 pts | 0.8 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 6 | 32.69 ms | 82.43 ms | 2.50× [2.45, 2.57] | +60.1% | [+59.1%, +61.0%] | 0.9 pts | 1.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of -0.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs baseline: the pooled interval [-1.98%, 0.22%] includes zero
- natural u64 n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=4096 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
