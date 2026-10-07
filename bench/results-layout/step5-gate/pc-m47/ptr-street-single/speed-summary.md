| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 59.0 | 49.7 | 0.84× [0.82, 0.86] | -19.2% | [-21.4%, -16.9%] | 2.5 pts | 3.4 | no | yes |
| single-value-ptr | street | 4096 | valuesFor | ordered | btree-map | 8 | 59.3 | 91.5 | 1.54× [1.53, 1.55] | +35.1% | [+34.7%, +35.5%] | 0.4 pts | 0.9 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 1127 | 2009 | 1.80× [1.77, 1.83] | +44.5% | [+43.5%, +45.5%] | 0.9 pts | 1.3 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1111 | 518 | 0.46× [0.46, 0.47] | -116.7% | [-119.4%, -114.0%] | 2.8 pts | 1.4 | yes | yes |
| single-value-ptr | street | 4096 | prefix | ordered | baseline | 8 | 219 | 227 | 1.04× [1.03, 1.05] | +3.9% | [+3.1%, +4.7%] | 0.9 pts | 0.9 | yes | yes |
| single-value-ptr | street | 4096 | prefix | ordered | btree-map | 8 | 218 | 136 | 0.63× [0.62, 0.63] | -59.6% | [-60.7%, -58.4%] | 1.6 pts | 1.2 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | baseline | 8 | 200 | 140 | 0.69× [0.68, 0.70] | -44.6% | [-46.4%, -42.7%] | 2.8 pts | 1.1 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | btree-map | 8 | 200 | 146 | 0.73× [0.71, 0.74] | -37.9% | [-40.2%, -35.7%] | 2.3 pts | 1.6 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | baseline | 8 | 3.02 ms | 2.00 ms | 0.67× [0.65, 0.68] | -49.9% | [-53.0%, -46.7%] | 6.3 pts | 1.8 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | btree-map | 8 | 3.08 ms | 2.25 ms | 0.75× [0.72, 0.78] | -34.1% | [-39.4%, -28.8%] | 5.4 pts | 2.3 | no | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | baseline | 8 | 71.5 | 69.4 | 0.97× [0.95, 0.99] | -2.9% | [-5.2%, -0.6%] | 2.4 pts | 2.6 | no | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | btree-map | 8 | 72.5 | 128 | 1.74× [1.68, 1.81] | +42.7% | [+40.5%, +44.9%] | 2.0 pts | 4.3 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | baseline | 8 | 1307 | 2472 | 1.90× [1.88, 1.92] | +47.3% | [+46.8%, +47.8%] | 0.5 pts | 1.0 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1292 | 650 | 0.50× [0.48, 0.51] | -101.5% | [-107.0%, -95.9%] | 6.1 pts | 5.3 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | baseline | 8 | 460 | 740 | 1.62× [1.61, 1.64] | +38.3% | [+37.8%, +38.8%] | 0.6 pts | 0.8 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | btree-map | 8 | 456 | 302 | 0.65× [0.63, 0.67] | -53.8% | [-59.0%, -48.6%] | 5.1 pts | 3.0 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | baseline | 8 | 216 | 181 | 0.84× [0.83, 0.85] | -19.3% | [-20.9%, -17.7%] | 1.8 pts | 1.5 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | btree-map | 8 | 225 | 206 | 0.92× [0.91, 0.93] | -9.2% | [-10.3%, -8.0%] | 1.2 pts | 0.8 | yes | yes |
| single-value-ptr | street | 16384 | build | ordered | baseline | 8 | 14.25 ms | 10.82 ms | 0.74× [0.72, 0.76] | -34.6% | [-38.4%, -30.8%] | 4.0 pts | 1.8 | no | yes |
| single-value-ptr | street | 16384 | build | ordered | btree-map | 8 | 14.54 ms | 12.20 ms | 0.83× [0.81, 0.85] | -19.9% | [-22.7%, -17.1%] | 3.2 pts | 1.5 | no | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 136 | 193 | 1.35× [1.20, 1.56] | +26.0% | [+16.3%, +35.7%] | 9.1 pts | 8.5 | no | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | btree-map | 8 | 153 | 271 | 1.77× [1.43, 2.33] | +43.6% | [+30.2%, +57.1%] | 12.6 pts | 11.7 | no | yes |
| single-value-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 1804 | 4165 | 2.37× [1.99, 2.95] | +57.9% | [+49.7%, +66.1%] | 7.6 pts | 13.1 | no | yes |
| single-value-ptr | street | 212449 | valuesBetween | ordered | btree-map | 8 | 1778 | 1248 | 0.70× [0.54, 1.03] | -42.0% | [-86.5%, +2.4%] | 41.4 pts | 21.1 | no | no |
| single-value-ptr | street | 212449 | prefix | ordered | baseline | 8 | 5323 | 14.9 µs | 2.86× [2.44, 3.45] | +65.0% | [+59.0%, +71.0%] | 5.7 pts | 6.4 | yes | yes |
| single-value-ptr | street | 212449 | prefix | ordered | btree-map | 8 | 5239 | 3930 | 0.77× [0.62, 1.03] | -29.6% | [-62.6%, +3.4%] | 32.1 pts | 12.7 | no | no |
| single-value-ptr | street | 212449 | churn | ordered | baseline | 8 | 408 | 467 | 1.15× [1.12, 1.18] | +13.2% | [+10.8%, +15.6%] | 2.6 pts | 2.3 | no | yes |
| single-value-ptr | street | 212449 | churn | ordered | btree-map | 8 | 463 | 524 | 1.15× [1.14, 1.17] | +13.3% | [+12.0%, +14.6%] | 2.1 pts | 1.5 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr street n=4096 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr street n=4096 build: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesFor: ordered vs btree-map: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 prefix: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 8.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesFor: ordered vs btree-map: the processes scatter 11.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 13.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs btree-map: the pooled interval [-86.54%, 2.44%] includes zero
- single-value-ptr street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 21.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs btree-map: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr street n=212449 prefix: ordered vs baseline: the processes scatter 6.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 prefix: ordered vs btree-map: the pooled interval [-62.60%, 3.37%] includes zero
- single-value-ptr street n=212449 prefix: ordered vs btree-map: the processes scatter 12.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 prefix: ordered vs btree-map: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr street n=212449 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
