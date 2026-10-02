| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique | u64 | 4096 | valuesFor | ordered | baseline | 8 | 17.9 | 17.7 | 0.99× [0.98, 0.99] | -1.3% | [-1.8%, -0.9%] | 0.5 pts | 2.1 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 12 | 25.7 | 22.7 | 0.88× [0.85, 0.90] | -14.3% | [-17.1%, -11.4%] | 6.2 pts | 18.3 | no | yes |
| unique | u64 | 65536 | valuesFor | ordered | baseline | 8 | 31.3 | 26.3 | 0.84× [0.83, 0.84] | -19.2% | [-19.9%, -18.6%] | 0.8 pts | 2.0 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 12 | 37.1 | 43.4 | 1.15× [1.10, 1.20] | +12.9% | [+9.0%, +16.9%] | 4.8 pts | 2.1 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.15% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 18.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
