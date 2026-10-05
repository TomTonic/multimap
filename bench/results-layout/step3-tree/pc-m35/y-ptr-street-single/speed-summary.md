| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 49.1 | 47.8 | 0.97× [0.96, 0.98] | -3.0% | [-4.1%, -2.0%] | 1.4 pts | 1.7 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 2020 | 1844 | 0.92× [0.90, 0.94] | -8.7% | [-11.3%, -6.1%] | 2.7 pts | 1.7 | no | yes |
| single-value-ptr | street | 4096 | prefix | ordered | baseline | 8 | 225 | 202 | 0.90× [0.88, 0.91] | -11.7% | [-13.1%, -10.2%] | 1.9 pts | 2.1 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | baseline | 8 | 129 | 121 | 0.94× [0.93, 0.95] | -6.6% | [-7.9%, -5.3%] | 1.6 pts | 1.2 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | baseline | 8 | 1.92 ms | 1.74 ms | 0.91× [0.90, 0.91] | -10.4% | [-11.5%, -9.3%] | 1.0 pts | 2.1 | yes | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | baseline | 8 | 68.5 | 66.7 | 0.98× [0.96, 0.99] | -2.4% | [-4.2%, -0.7%] | 1.8 pts | 1.9 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | baseline | 8 | 2423 | 2228 | 0.92× [0.91, 0.93] | -8.5% | [-9.3%, -7.7%] | 1.4 pts | 1.5 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | baseline | 8 | 738 | 646 | 0.89× [0.88, 0.91] | -12.1% | [-13.7%, -10.5%] | 1.5 pts | 0.8 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | baseline | 8 | 164 | 158 | 0.97× [0.97, 0.97] | -3.1% | [-3.5%, -2.7%] | 0.4 pts | 0.4 | yes | yes |
| single-value-ptr | street | 16384 | build | ordered | baseline | 8 | 9.19 ms | 8.56 ms | 0.93× [0.93, 0.93] | -7.5% | [-7.9%, -7.0%] | 0.6 pts | 1.2 | yes | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | baseline | 6 | 163 | 151 | 0.93× [0.91, 0.94] | -8.0% | [-9.6%, -6.3%] | 1.6 pts | 1.4 | yes | yes |
| single-value-ptr | street | 212449 | valuesBetween | ordered | baseline | 6 | 3602 | 3273 | 0.92× [0.91, 0.93] | -8.7% | [-10.2%, -7.2%] | 1.4 pts | 1.0 | yes | yes |
| single-value-ptr | street | 212449 | prefix | ordered | baseline | 6 | 13.3 µs | 12.1 µs | 0.92× [0.90, 0.93] | -8.9% | [-10.8%, -7.0%] | 1.8 pts | 1.8 | yes | yes |
| single-value-ptr | street | 212449 | churn | ordered | baseline | 6 | 382 | 375 | 0.96× [0.95, 0.97] | -4.1% | [-5.2%, -3.0%] | 1.1 pts | 0.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr street n=4096 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.80% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
