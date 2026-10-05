| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 54.3 | 53.2 | 0.97× [0.96, 0.99] | -2.7% | [-4.3%, -1.0%] | 1.7 pts | 1.7 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 2468 | 2318 | 0.94× [0.91, 0.96] | -6.7% | [-9.6%, -3.7%] | 2.7 pts | 1.0 | no | yes |
| natural | street | 4096 | prefix | ordered | baseline | 8 | 248 | 224 | 0.90× [0.89, 0.92] | -10.7% | [-12.7%, -8.6%] | 2.6 pts | 2.0 | no | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 103 | 96.1 | 0.93× [0.92, 0.93] | -7.8% | [-8.6%, -7.1%] | 0.9 pts | 0.8 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 4.13 ms | 3.79 ms | 0.92× [0.92, 0.93] | -8.5% | [-9.2%, -7.8%] | 0.7 pts | 1.0 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 75.0 | 73.5 | 0.99× [0.97, 1.01] | -0.9% | [-2.8%, +0.9%] | 1.7 pts | 2.0 | yes | no |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 3046 | 2914 | 0.95× [0.94, 0.96] | -5.1% | [-5.9%, -4.3%] | 0.8 pts | 0.8 | yes | yes |
| natural | street | 16384 | prefix | ordered | baseline | 8 | 896 | 820 | 0.92× [0.91, 0.92] | -9.0% | [-9.8%, -8.2%] | 0.9 pts | 0.5 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 142 | 133 | 0.93× [0.92, 0.94] | -7.1% | [-8.1%, -6.2%] | 0.9 pts | 0.9 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 21.71 ms | 21.95 ms | 1.01× [1.00, 1.02] | +1.0% | [+0.4%, +1.6%] | 0.8 pts | 1.2 | yes | yes |
| natural | street | 212449 | valuesFor | ordered | baseline | 8 | 194 | 192 | 1.00× [0.97, 1.03] | -0.2% | [-3.4%, +3.1%] | 3.1 pts | 2.2 | no | no |
| natural | street | 212449 | valuesBetween | ordered | baseline | 8 | 5765 | 5700 | 0.99× [0.95, 1.04] | -1.1% | [-5.7%, +3.5%] | 4.5 pts | 4.5 | no | no |
| natural | street | 212449 | prefix | ordered | baseline | 8 | 17.7 µs | 17.5 µs | 0.98× [0.92, 1.05] | -2.0% | [-8.4%, +4.4%] | 6.0 pts | 4.0 | no | no |
| natural | street | 212449 | churn | ordered | baseline | 8 | 447 | 428 | 0.95× [0.94, 0.97] | -4.7% | [-6.6%, -2.9%] | 4.1 pts | 0.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.76%, 0.93%] includes zero
- natural street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs baseline: the pooled difference of -0.17% does not clear the 1.02% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=212449 valuesFor: ordered vs baseline: the pooled interval [-3.41%, 3.07%] includes zero
- natural street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs baseline: the pooled interval [-5.74%, 3.55%] includes zero
- natural street n=212449 valuesBetween: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural street n=212449 prefix: ordered vs baseline: the pooled difference of -2.03% does not clear the 2.63% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=212449 prefix: ordered vs baseline: the pooled interval [-8.44%, 4.38%] includes zero
- natural street n=212449 prefix: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
