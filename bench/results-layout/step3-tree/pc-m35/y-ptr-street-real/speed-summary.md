| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 56.4 | 54.5 | 0.96× [0.96, 0.97] | -3.6% | [-4.4%, -2.9%] | 1.5 pts | 1.5 | yes | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 2552 | 2388 | 0.94× [0.91, 0.98] | -6.1% | [-9.8%, -2.5%] | 3.6 pts | 1.4 | no | yes |
| natural-ptr | street | 4096 | prefix | ordered | baseline | 8 | 251 | 230 | 0.92× [0.90, 0.94] | -8.2% | [-10.5%, -5.9%] | 2.5 pts | 1.9 | no | yes |
| natural-ptr | street | 4096 | churn | ordered | baseline | 8 | 105 | 104 | 0.99× [0.98, 1.00] | -1.1% | [-2.4%, +0.2%] | 1.3 pts | 1.1 | yes | no |
| natural-ptr | street | 4096 | build | ordered | baseline | 8 | 4.24 ms | 4.15 ms | 0.98× [0.97, 0.99] | -2.0% | [-2.7%, -1.3%] | 0.7 pts | 0.9 | yes | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | baseline | 6 | 76.3 | 75.7 | 0.99× [0.97, 1.01] | -1.5% | [-3.4%, +0.5%] | 1.9 pts | 2.0 | yes | no |
| natural-ptr | street | 16384 | valuesBetween | ordered | baseline | 6 | 3113 | 2989 | 0.96× [0.95, 0.97] | -4.3% | [-5.2%, -3.4%] | 0.9 pts | 1.1 | yes | yes |
| natural-ptr | street | 16384 | prefix | ordered | baseline | 6 | 915 | 852 | 0.93× [0.92, 0.95] | -7.5% | [-9.2%, -5.8%] | 1.6 pts | 0.8 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | baseline | 6 | 143 | 143 | 0.99× [0.98, 1.00] | -0.8% | [-1.7%, +0.1%] | 0.8 pts | 0.8 | yes | no |
| natural-ptr | street | 16384 | build | ordered | baseline | 6 | 21.87 ms | 21.98 ms | 1.00× [1.00, 1.01] | +0.3% | [-0.3%, +1.0%] | 0.6 pts | 0.9 | yes | no |
| natural-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 213 | 199 | 0.95× [0.91, 0.99] | -5.6% | [-9.6%, -1.5%] | 3.9 pts | 2.9 | no | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 5961 | 5872 | 0.98× [0.95, 1.01] | -2.2% | [-5.6%, +1.2%] | 3.3 pts | 2.3 | no | no |
| natural-ptr | street | 212449 | prefix | ordered | baseline | 8 | 18.6 µs | 18.1 µs | 0.97× [0.94, 1.00] | -3.4% | [-6.7%, -0.0%] | 3.2 pts | 2.0 | no | yes |
| natural-ptr | street | 212449 | churn | ordered | baseline | 8 | 451 | 445 | 0.99× [0.95, 1.02] | -1.5% | [-5.1%, +2.1%] | 3.3 pts | 0.8 | no | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr street n=4096 churn: ordered vs baseline: the pooled interval [-2.36%, 0.19%] includes zero
- natural-ptr street n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.45%, 0.53%] includes zero
- natural-ptr street n=16384 churn: ordered vs baseline: the pooled difference of -0.80% does not clear the 0.88% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr street n=16384 churn: ordered vs baseline: the pooled interval [-1.69%, 0.09%] includes zero
- natural-ptr street n=16384 build: ordered vs baseline: the pooled difference of 0.34% does not clear the 0.49% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr street n=16384 build: ordered vs baseline: the pooled interval [-0.30%, 0.99%] includes zero
- natural-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesBetween: ordered vs baseline: the pooled interval [-5.64%, 1.19%] includes zero
- natural-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 prefix: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 churn: ordered vs baseline: the pooled difference of -1.48% does not clear the 1.59% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr street n=212449 churn: ordered vs baseline: the pooled interval [-5.07%, 2.10%] includes zero
