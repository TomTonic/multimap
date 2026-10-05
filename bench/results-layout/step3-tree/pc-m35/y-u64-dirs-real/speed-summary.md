| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 84.9 | 84.0 | 0.99× [0.98, 1.00] | -1.0% | [-2.4%, +0.5%] | 1.4 pts | 1.1 | yes | no |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2879 | 2697 | 0.93× [0.91, 0.95] | -7.6% | [-9.9%, -5.2%] | 3.4 pts | 1.1 | no | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 8 | 3139 | 2747 | 0.87× [0.86, 0.89] | -14.4% | [-15.8%, -12.9%] | 1.4 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 153 | 150 | 0.98× [0.96, 0.99] | -2.4% | [-4.0%, -0.8%] | 1.9 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 7.22 ms | 7.15 ms | 0.99× [0.99, 1.00] | -0.9% | [-1.4%, -0.4%] | 0.5 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 116 | 115 | 0.99× [0.98, 1.00] | -1.2% | [-2.3%, -0.1%] | 1.2 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3630 | 3481 | 0.96× [0.95, 0.97] | -4.0% | [-5.5%, -2.6%] | 1.7 pts | 2.8 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 8 | 14.8 µs | 14.3 µs | 0.95× [0.93, 0.97] | -5.3% | [-7.7%, -2.8%] | 2.4 pts | 0.8 | no | no |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 203 | 198 | 0.97× [0.97, 0.98] | -2.6% | [-3.3%, -1.9%] | 0.7 pts | 0.9 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 38.37 ms | 40.43 ms | 1.06× [1.05, 1.06] | +5.4% | [+4.9%, +5.8%] | 0.8 pts | 1.3 | yes | yes |
| natural | dirs | 86215 | valuesFor | ordered | baseline | 8 | 184 | 187 | 1.01× [1.01, 1.02] | +1.5% | [+0.5%, +2.4%] | 1.4 pts | 0.9 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 4506 | 4333 | 0.97× [0.96, 0.98] | -3.0% | [-3.9%, -2.1%] | 1.2 pts | 1.3 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | baseline | 8 | 117.4 µs | 113.1 µs | 0.96× [0.96, 0.97] | -4.0% | [-4.5%, -3.5%] | 1.5 pts | 0.9 | yes | yes |
| natural | dirs | 86215 | churn | ordered | baseline | 8 | 414 | 404 | 0.96× [0.93, 1.00] | -3.9% | [-7.5%, -0.3%] | 4.3 pts | 1.4 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.98% does not clear the 1.01% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.43%, 0.47%] includes zero
- natural dirs n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.99% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 prefix: ordered vs baseline: the pooled difference of -5.28% does not clear the 10.43% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +8.72% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
