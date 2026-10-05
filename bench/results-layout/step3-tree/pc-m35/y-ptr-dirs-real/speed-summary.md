| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 86.8 | 86.4 | 0.99× [0.98, 1.01] | -0.8% | [-2.1%, +0.5%] | 1.6 pts | 1.4 | yes | no |
| natural-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2974 | 2838 | 0.95× [0.93, 0.97] | -5.2% | [-7.6%, -2.8%] | 2.4 pts | 1.0 | no | yes |
| natural-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 3225 | 2941 | 0.91× [0.90, 0.92] | -10.2% | [-11.7%, -8.7%] | 1.7 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 153 | 158 | 1.03× [1.01, 1.05] | +3.0% | [+1.3%, +4.7%] | 1.7 pts | 1.4 | yes | yes |
| natural-ptr | dirs | 4096 | build | ordered | baseline | 8 | 7.37 ms | 7.49 ms | 1.02× [1.01, 1.03] | +2.1% | [+1.3%, +2.8%] | 0.9 pts | 1.6 | yes | yes |
| natural-ptr | dirs | 16384 | valuesFor | ordered | baseline | 6 | 122 | 121 | 1.00× [0.99, 1.01] | -0.2% | [-1.5%, +1.2%] | 1.3 pts | 1.6 | yes | no |
| natural-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 3772 | 3671 | 0.97× [0.96, 0.98] | -3.2% | [-4.0%, -2.3%] | 0.8 pts | 1.1 | yes | yes |
| natural-ptr | dirs | 16384 | prefix | ordered | baseline | 6 | 18.3 µs | 16.9 µs | 0.96× [0.94, 0.97] | -4.4% | [-6.1%, -2.7%] | 1.6 pts | 0.8 | yes | no |
| natural-ptr | dirs | 16384 | churn | ordered | baseline | 6 | 208 | 215 | 1.03× [1.02, 1.05] | +3.2% | [+1.8%, +4.6%] | 1.3 pts | 1.5 | yes | yes |
| natural-ptr | dirs | 16384 | build | ordered | baseline | 6 | 38.93 ms | 40.36 ms | 1.04× [1.03, 1.05] | +3.7% | [+3.0%, +4.4%] | 0.7 pts | 1.2 | yes | yes |
| natural-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 227 | 228 | 1.01× [0.99, 1.02] | +0.8% | [-0.6%, +2.3%] | 1.4 pts | 1.9 | yes | no |
| natural-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 5197 | 5270 | 1.00× [0.98, 1.03] | +0.0% | [-2.5%, +2.6%] | 2.5 pts | 2.6 | no | no |
| natural-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 120.4 µs | 122.8 µs | 0.98× [0.96, 1.01] | -1.7% | [-4.2%, +0.7%] | 2.8 pts | 0.8 | no | no |
| natural-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 430 | 429 | 1.00× [0.99, 1.02] | +0.4% | [-1.5%, +2.2%] | 3.0 pts | 1.1 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr dirs n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.78% does not clear the 0.90% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.09%, 0.53%] includes zero
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.17% does not clear the 0.36% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.51%, 1.17%] includes zero
- natural-ptr dirs n=16384 prefix: ordered vs baseline: the pooled difference of -4.39% does not clear the 14.26% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-0.64%, 2.30%] includes zero
- natural-ptr dirs n=86215 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: the pooled difference of 0.04% does not clear the 0.42% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: the pooled interval [-2.49%, 2.57%] includes zero
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr dirs n=86215 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr dirs n=86215 prefix: ordered vs baseline: the pooled difference of -1.75% does not clear the 1.96% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=86215 prefix: ordered vs baseline: the pooled interval [-4.18%, 0.69%] includes zero
- natural-ptr dirs n=86215 churn: ordered vs baseline: the pooled difference of 0.37% does not clear the 1.85% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr dirs n=86215 churn: ordered vs baseline: the pooled interval [-1.49%, 2.24%] includes zero
