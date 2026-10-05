| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | dirs | 4096 | valuesFor | ordered | baseline | 6 | 78.2 | 76.5 | 0.99× [0.97, 1.00] | -1.4% | [-3.0%, +0.2%] | 1.5 pts | 1.6 | yes | no |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 2242 | 2102 | 0.94× [0.93, 0.95] | -6.6% | [-7.8%, -5.4%] | 1.2 pts | 0.5 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | baseline | 6 | 2555 | 2311 | 0.90× [0.89, 0.91] | -11.0% | [-11.9%, -10.0%] | 0.9 pts | 0.4 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | baseline | 6 | 198 | 199 | 1.00× [0.99, 1.01] | +0.3% | [-0.7%, +1.4%] | 1.0 pts | 0.9 | yes | no |
| single-value-ptr | dirs | 4096 | build | ordered | baseline | 6 | 2.86 ms | 2.79 ms | 0.98× [0.97, 0.99] | -2.3% | [-3.5%, -1.1%] | 1.2 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | baseline | 8 | 108 | 107 | 0.99× [0.98, 1.00] | -1.1% | [-2.0%, -0.2%] | 0.9 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 2790 | 2654 | 0.95× [0.95, 0.95] | -5.2% | [-5.6%, -4.8%] | 0.8 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | baseline | 8 | 11.2 µs | 10.3 µs | 0.92× [0.91, 0.94] | -8.1% | [-10.0%, -6.2%] | 2.4 pts | 0.6 | yes | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | baseline | 8 | 247 | 253 | 1.02× [1.01, 1.03] | +1.9% | [+0.7%, +3.1%] | 1.2 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 16384 | build | ordered | baseline | 8 | 13.70 ms | 13.63 ms | 0.99× [0.99, 1.00] | -0.7% | [-1.3%, -0.0%] | 0.9 pts | 1.4 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 159 | 157 | 0.99× [0.97, 1.01] | -1.5% | [-3.4%, +0.5%] | 2.0 pts | 1.9 | yes | no |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3179 | 3038 | 0.95× [0.94, 0.97] | -4.9% | [-6.2%, -3.5%] | 1.4 pts | 1.8 | yes | yes |
| single-value-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 78.4 µs | 72.8 µs | 0.94× [0.93, 0.95] | -6.4% | [-7.6%, -5.3%] | 1.1 pts | 1.2 | yes | yes |
| single-value-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 406 | 402 | 1.00× [0.98, 1.01] | -0.4% | [-1.8%, +0.9%] | 1.5 pts | 1.2 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr dirs n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.98%, 0.18%] includes zero
- single-value-ptr dirs n=4096 churn: ordered vs baseline: the pooled difference of 0.33% does not clear the 0.74% noise floor, the bound on what the harness reports between identical code in every process
- single-value-ptr dirs n=4096 churn: ordered vs baseline: the pooled interval [-0.74%, 1.39%] includes zero
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-3.40%, 0.50%] includes zero
- single-value-ptr dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -3.19% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=86215 churn: ordered vs baseline: the pooled difference of -0.45% does not clear the 0.96% noise floor, the bound on what the harness reports between identical code in every process
- single-value-ptr dirs n=86215 churn: ordered vs baseline: the pooled interval [-1.75%, 0.86%] includes zero
