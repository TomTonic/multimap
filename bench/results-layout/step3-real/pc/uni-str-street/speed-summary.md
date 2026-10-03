| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | street | 4096 | valuesFor | ordered | btree-map | 10 | 47.3 | 92.3 | 1.96× [1.94, 1.97] | +48.9% | [+48.6%, +49.2%] | 0.7 pts | 2.1 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | btree-map | 10 | 1757 | 501 | 0.28× [0.28, 0.29] | -251.8% | [-256.1%, -247.6%] | 5.6 pts | 1.8 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | btree-map | 10 | 198 | 134 | 0.68× [0.67, 0.68] | -48.0% | [-49.6%, -46.5%] | 1.9 pts | 1.2 | yes | yes |
| unique-str | street | 4096 | churn | ordered | btree-map | 10 | 127 | 144 | 1.14× [1.13, 1.14] | +12.2% | [+11.8%, +12.7%] | 1.0 pts | 0.8 | yes | yes |
| unique-str | street | 4096 | build | ordered | btree-map | 10 | 1.74 ms | 2.17 ms | 1.27× [1.24, 1.31] | +21.3% | [+19.3%, +23.4%] | 2.6 pts | 3.7 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | btree-map | 10 | 66.9 | 129 | 1.95× [1.90, 1.99] | +48.6% | [+47.4%, +49.8%] | 1.5 pts | 3.5 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | btree-map | 10 | 2169 | 612 | 0.28× [0.27, 0.29] | -257.9% | [-265.4%, -250.5%] | 8.9 pts | 3.7 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | btree-map | 10 | 627 | 286 | 0.45× [0.44, 0.46] | -121.6% | [-126.3%, -116.9%] | 5.3 pts | 2.1 | yes | yes |
| unique-str | street | 16384 | churn | ordered | btree-map | 10 | 165 | 202 | 1.22× [1.19, 1.25] | +18.1% | [+16.3%, +19.9%] | 1.7 pts | 1.8 | yes | yes |
| unique-str | street | 16384 | build | ordered | btree-map | 10 | 8.71 ms | 11.49 ms | 1.32× [1.30, 1.33] | +24.1% | [+23.2%, +25.0%] | 0.9 pts | 0.9 | yes | yes |
| unique-str | street | 212449 | valuesFor | ordered | btree-map | 12 | 235 | 248 | 1.05× [1.01, 1.10] | +5.1% | [+0.8%, +9.5%] | 5.1 pts | 4.2 | no | yes |
| unique-str | street | 212449 | valuesBetween | ordered | btree-map | 12 | 3944 | 1495 | 0.39× [0.37, 0.42] | -155.3% | [-170.5%, -140.0%] | 18.9 pts | 1.5 | yes | yes |
| unique-str | street | 212449 | prefix | ordered | btree-map | 12 | 12.4 µs | 3956 | 0.31× [0.30, 0.32] | -223.7% | [-234.5%, -212.9%] | 13.2 pts | 1.2 | yes | yes |
| unique-str | street | 212449 | churn | ordered | btree-map | 12 | 531 | 559 | 1.05× [1.04, 1.07] | +5.1% | [+3.7%, +6.5%] | 2.9 pts | 2.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=4096 build: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 prefix: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesFor: ordered vs btree-map: 9 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str street n=212449 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -3.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=212449 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -1.61% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=212449 churn: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 churn: ordered vs btree-map: 10 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
