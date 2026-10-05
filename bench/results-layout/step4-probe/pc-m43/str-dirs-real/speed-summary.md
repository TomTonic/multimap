| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 126 | 123 | 0.98× [0.97, 1.00] | -1.7% | [-3.1%, -0.2%] | 1.4 pts | 1.3 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 127 | 167 | 1.32× [1.30, 1.33] | +24.0% | [+23.3%, +24.6%] | 1.0 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6205 | 6442 | 1.04× [1.03, 1.05] | +4.0% | [+3.3%, +4.7%] | 0.9 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6219 | 10.1 µs | 1.63× [1.62, 1.64] | +38.7% | [+38.3%, +39.0%] | 0.3 pts | 0.6 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 7052 | 7289 | 1.03× [1.01, 1.05] | +3.2% | [+1.4%, +5.0%] | 1.7 pts | 0.7 | yes | no |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7077 | 14.5 µs | 2.04× [1.99, 2.11] | +51.1% | [+49.6%, +52.6%] | 1.5 pts | 0.5 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 241 | 185 | 0.77× [0.77, 0.78] | -29.6% | [-30.4%, -28.8%] | 0.9 pts | 0.7 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 242 | 243 | 1.00× [0.98, 1.01] | -0.3% | [-1.8%, +1.2%] | 1.4 pts | 0.8 | yes | no |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 11.35 ms | 8.22 ms | 0.72× [0.72, 0.72] | -38.8% | [-39.3%, -38.2%] | 0.6 pts | 0.6 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 11.53 ms | 11.35 ms | 0.99× [0.99, 1.00] | -0.9% | [-1.5%, -0.3%] | 0.7 pts | 0.9 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 160 | 158 | 0.99× [0.98, 1.00] | -1.2% | [-1.9%, -0.5%] | 0.8 pts | 0.9 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 170 | 229 | 1.37× [1.29, 1.45] | +26.7% | [+22.6%, +30.9%] | 4.1 pts | 4.7 | no | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 6959 | 7266 | 1.05× [1.04, 1.06] | +4.7% | [+4.1%, +5.2%] | 0.7 pts | 0.8 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 7000 | 11.0 µs | 1.58× [1.57, 1.59] | +36.6% | [+36.3%, +37.0%] | 0.4 pts | 0.6 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 8 | 35.8 µs | 37.6 µs | 1.06× [1.05, 1.06] | +5.5% | [+5.2%, +5.9%] | 0.4 pts | 0.5 | yes | no |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 8 | 32.5 µs | 57.2 µs | 1.80× [1.78, 1.83] | +44.5% | [+43.7%, +45.4%] | 0.8 pts | 1.7 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 8 | 323 | 262 | 0.82× [0.81, 0.82] | -22.7% | [-23.7%, -21.6%] | 2.1 pts | 1.8 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 8 | 328 | 372 | 1.14× [1.13, 1.15] | +12.3% | [+11.4%, +13.2%] | 0.9 pts | 1.0 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 8 | 60.73 ms | 45.79 ms | 0.75× [0.75, 0.76] | -32.8% | [-33.8%, -31.7%] | 1.0 pts | 0.8 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 8 | 60.79 ms | 65.44 ms | 1.07× [1.06, 1.08] | +6.7% | [+5.9%, +7.6%] | 0.8 pts | 1.1 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 280 | 290 | 1.02× [0.99, 1.05] | +2.1% | [-0.7%, +4.9%] | 2.6 pts | 3.0 | no | no |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 306 | 504 | 1.64× [1.59, 1.69] | +38.9% | [+37.0%, +40.7%] | 2.0 pts | 2.2 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 10.4 µs | 11.0 µs | 1.07× [1.06, 1.08] | +6.8% | [+5.7%, +7.8%] | 1.5 pts | 2.0 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 11.8 µs | 23.7 µs | 2.02× [1.95, 2.08] | +50.4% | [+48.8%, +52.0%] | 1.5 pts | 2.3 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 226.1 µs | 255.4 µs | 1.07× [1.02, 1.12] | +6.5% | [+2.2%, +10.7%] | 5.5 pts | 1.0 | no | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 229.3 µs | 505.8 µs | 2.15× [2.05, 2.26] | +53.5% | [+51.2%, +55.8%] | 2.5 pts | 0.8 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 605 | 534 | 0.89× [0.87, 0.91] | -12.5% | [-15.5%, -9.5%] | 3.0 pts | 0.9 | no | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 668 | 700 | 1.06× [1.04, 1.09] | +6.1% | [+4.0%, +8.1%] | 2.0 pts | 2.0 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 prefix: ordered vs baseline: the pooled difference of 3.20% does not clear the 4.48% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=4096 churn: ordered vs btree-sets: the pooled difference of -0.33% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=4096 churn: ordered vs btree-sets: the pooled interval [-1.82%, 1.15%] includes zero
- natural-str dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 prefix: ordered vs baseline: the pooled difference of 5.53% does not clear the 13.39% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of -11.88% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -5.67% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-0.73%, 4.89%] includes zero
- natural-str dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -3.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
