| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value | street | 4096 | valuesFor | ordered | baseline | 6 | 48.6 | 48.2 | 0.99× [0.99, 1.00] | -0.8% | [-1.4%, -0.2%] | 0.6 pts | 1.1 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 6 | 1927 | 405 | 0.21× [0.21, 0.21] | -377.8% | [-382.9%, -372.7%] | 4.9 pts | 0.9 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 6 | 218 | 104 | 0.48× [0.48, 0.48] | -109.6% | [-110.3%, -108.9%] | 0.7 pts | 0.6 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 6 | 126 | 104 | 0.82× [0.81, 0.82] | -22.0% | [-22.7%, -21.3%] | 0.7 pts | 0.5 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 6 | 1.92 ms | 1.88 ms | 0.98× [0.97, 0.99] | -1.9% | [-3.1%, -0.7%] | 1.1 pts | 1.2 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 6 | 67.8 | 60.8 | 0.90× [0.89, 0.91] | -11.1% | [-11.9%, -10.4%] | 0.7 pts | 0.9 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 6 | 2375 | 481 | 0.20× [0.20, 0.20] | -396.7% | [-400.2%, -393.2%] | 3.3 pts | 0.7 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 6 | 705 | 180 | 0.26× [0.25, 0.26] | -289.3% | [-293.9%, -284.7%] | 4.4 pts | 1.4 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 6 | 161 | 130 | 0.81× [0.80, 0.82] | -23.6% | [-25.7%, -21.4%] | 2.0 pts | 2.0 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 6 | 9.08 ms | 8.85 ms | 0.98× [0.97, 0.98] | -2.6% | [-3.1%, -2.0%] | 0.5 pts | 0.6 | yes | yes |
| single-value | street | 212449 | valuesFor | ordered | baseline | 8 | 145 | 106 | 0.73× [0.70, 0.75] | -37.3% | [-42.1%, -32.6%] | 6.0 pts | 3.3 | no | yes |
| single-value | street | 212449 | valuesBetween | ordered | baseline | 8 | 3246 | 768 | 0.24× [0.24, 0.25] | -313.7% | [-323.3%, -304.0%] | 12.9 pts | 0.7 | yes | yes |
| single-value | street | 212449 | prefix | ordered | baseline | 8 | 12.1 µs | 1778 | 0.15× [0.14, 0.15] | -579.0% | [-603.2%, -554.9%] | 24.0 pts | 1.5 | yes | yes |
| single-value | street | 212449 | churn | ordered | baseline | 8 | 404 | 275 | 0.71× [0.69, 0.73] | -41.8% | [-45.9%, -37.7%] | 5.5 pts | 1.5 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value street n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
