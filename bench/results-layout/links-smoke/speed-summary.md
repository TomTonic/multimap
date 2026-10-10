| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | links | 4096 | valuesFor | ordered | btree-sets | 4 | 201 | 410 | 2.02× [1.97, 2.08] | +50.6% | [+49.2%, +52.0%] | 0.9 pts | 1.1 | yes | yes |
| natural | links | 4096 | churn | ordered | btree-sets | 4 | 156 | 194 | 1.23× [1.18, 1.29] | +19.0% | [+15.4%, +22.6%] | 2.2 pts | 1.5 | no | yes |
| single-value | links | 4096 | valuesFor | ordered | btree-map | 4 | 54.4 | 99.1 | 1.82× [1.81, 1.84] | +45.2% | [+44.8%, +45.6%] | 0.3 pts | 0.4 | yes | yes |
| single-value | links | 4096 | churn | ordered | btree-map | 4 | 141 | 154 | 1.09× [1.03, 1.16] | +8.4% | [+3.2%, +13.5%] | 3.2 pts | 1.9 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural links n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.61% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
