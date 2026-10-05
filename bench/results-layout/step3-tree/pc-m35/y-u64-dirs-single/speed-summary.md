| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 77.6 | 88.9 | 1.14× [1.11, 1.17] | +12.3% | [+10.2%, +14.5%] | 2.0 pts | 3.3 | no | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2149 | 660 | 0.30× [0.30, 0.31] | -228.3% | [-237.4%, -219.2%] | 8.7 pts | 1.6 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 8 | 2512 | 678 | 0.27× [0.27, 0.28] | -267.3% | [-271.9%, -262.7%] | 4.3 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 196 | 198 | 1.00× [0.99, 1.01] | +0.3% | [-0.6%, +1.3%] | 1.0 pts | 1.0 | yes | no |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 2.94 ms | 3.53 ms | 1.21× [1.19, 1.23] | +17.3% | [+15.7%, +18.9%] | 1.5 pts | 1.7 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 106 | 113 | 1.07× [1.06, 1.07] | +6.3% | [+5.7%, +6.9%] | 1.2 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 2709 | 821 | 0.30× [0.30, 0.31] | -230.2% | [-233.4%, -227.0%] | 3.6 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 8 | 10.8 µs | 2189 | 0.21× [0.20, 0.21] | -387.4% | [-398.2%, -376.7%] | 10.2 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 246 | 240 | 0.97× [0.95, 0.99] | -3.4% | [-5.7%, -1.0%] | 2.3 pts | 3.2 | no | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 13.52 ms | 15.94 ms | 1.19× [1.18, 1.19] | +15.7% | [+15.2%, +16.2%] | 0.7 pts | 1.0 | yes | yes |
| single-value | dirs | 86215 | valuesFor | ordered | baseline | 8 | 151 | 152 | 1.00× [0.97, 1.03] | -0.4% | [-3.5%, +2.7%] | 3.3 pts | 3.4 | no | no |
| single-value | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3028 | 985 | 0.32× [0.32, 0.33] | -209.8% | [-212.7%, -206.9%] | 3.1 pts | 1.1 | yes | yes |
| single-value | dirs | 86215 | prefix | ordered | baseline | 8 | 75.6 µs | 14.5 µs | 0.19× [0.19, 0.20] | -418.0% | [-429.2%, -406.8%] | 11.4 pts | 1.2 | yes | yes |
| single-value | dirs | 86215 | churn | ordered | baseline | 8 | 398 | 324 | 0.83× [0.81, 0.85] | -20.4% | [-23.4%, -17.4%] | 3.5 pts | 1.3 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 churn: ordered vs baseline: the pooled difference of 0.31% does not clear the 0.78% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 churn: ordered vs baseline: the pooled interval [-0.63%, 1.25%] includes zero
- single-value dirs n=16384 churn: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesFor: ordered vs baseline: the pooled difference of -0.37% does not clear the 1.27% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-3.47%, 2.73%] includes zero
- single-value dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=86215 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
