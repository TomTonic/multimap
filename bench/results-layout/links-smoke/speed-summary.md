| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | links | 4096 | valuesFor | ordered | btree-sets | 4 | 225 | 435 | 1.93× [1.90, 1.96] | +48.1% | [+47.3%, +48.9%] | 0.5 pts | 0.8 | yes | yes |
| natural | links | 4096 | churn | ordered | btree-sets | 4 | 157 | 199 | 1.24× [1.18, 1.31] | +19.6% | [+15.3%, +23.9%] | 2.7 pts | 2.3 | no | yes |
| single-value | links | 4096 | valuesFor | ordered | btree-map | 4 | 54.2 | 98.5 | 1.82× [1.76, 1.87] | +44.9% | [+43.2%, +46.6%] | 1.1 pts | 2.2 | yes | yes |
| single-value | links | 4096 | churn | ordered | btree-map | 4 | 151 | 157 | 1.04× [1.01, 1.08] | +4.1% | [+0.8%, +7.3%] | 2.0 pts | 1.5 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural links n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
