| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 8 | 66.2 | 47.1 | 0.71× [0.68, 0.73] | -41.6% | [-46.5%, -36.7%] | 4.8 pts | 4.1 | no | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 8 | 66.9 | 91.9 | 1.36× [1.34, 1.38] | +26.5% | [+25.3%, +27.6%] | 2.3 pts | 3.5 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 3661 | 1835 | 0.50× [0.49, 0.51] | -99.0% | [-102.3%, -95.6%] | 3.2 pts | 1.2 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 8 | 3649 | 1130 | 0.30× [0.30, 0.31] | -232.6% | [-238.0%, -227.1%] | 5.3 pts | 1.0 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 8 | 369 | 202 | 0.55× [0.54, 0.56] | -82.6% | [-84.9%, -80.2%] | 2.8 pts | 1.2 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 8 | 367 | 188 | 0.50× [0.50, 0.51] | -98.9% | [-101.0%, -96.8%] | 2.3 pts | 0.8 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 8 | 127 | 120 | 0.95× [0.95, 0.96] | -4.7% | [-5.8%, -3.7%] | 1.4 pts | 1.0 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 8 | 125 | 139 | 1.11× [1.10, 1.13] | +9.9% | [+8.7%, +11.2%] | 1.3 pts | 1.5 | yes | yes |
| single-value-str | street | 4096 | build | ordered | baseline | 8 | 1.91 ms | 1.76 ms | 0.91× [0.89, 0.94] | -9.3% | [-11.8%, -6.8%] | 2.5 pts | 1.5 | no | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 8 | 1.90 ms | 2.17 ms | 1.16× [1.14, 1.18] | +13.7% | [+11.9%, +15.5%] | 1.8 pts | 1.5 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 8 | 86.0 | 67.0 | 0.77× [0.75, 0.78] | -30.1% | [-32.7%, -27.5%] | 2.8 pts | 2.5 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 8 | 87.5 | 126 | 1.43× [1.39, 1.47] | +30.0% | [+28.2%, +31.8%] | 2.5 pts | 3.9 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 4203 | 2219 | 0.53× [0.53, 0.53] | -89.0% | [-90.2%, -87.8%] | 1.8 pts | 1.4 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 8 | 4217 | 1371 | 0.32× [0.32, 0.33] | -208.4% | [-211.6%, -205.3%] | 3.1 pts | 1.6 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 8 | 1335 | 666 | 0.50× [0.50, 0.50] | -100.1% | [-102.0%, -98.2%] | 2.1 pts | 0.6 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 8 | 1319 | 486 | 0.37× [0.36, 0.38] | -169.7% | [-174.2%, -165.3%] | 4.4 pts | 1.1 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 8 | 161 | 155 | 0.97× [0.96, 0.98] | -3.4% | [-4.5%, -2.2%] | 1.4 pts | 1.3 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 8 | 163 | 194 | 1.19× [1.17, 1.21] | +16.2% | [+14.8%, +17.6%] | 1.4 pts | 1.5 | yes | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 8 | 9.25 ms | 8.49 ms | 0.92× [0.91, 0.93] | -8.8% | [-10.0%, -7.5%] | 1.5 pts | 1.4 | yes | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 8 | 9.39 ms | 11.68 ms | 1.24× [1.22, 1.26] | +19.6% | [+18.3%, +20.9%] | 1.6 pts | 0.8 | yes | yes |
| single-value-str | street | 212449 | valuesFor | ordered | baseline | 8 | 214 | 189 | 0.86× [0.83, 0.89] | -16.7% | [-20.8%, -12.5%] | 4.6 pts | 2.4 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | btree-map | 8 | 275 | 258 | 0.97× [0.91, 1.02] | -3.5% | [-9.5%, +2.4%] | 7.9 pts | 5.1 | no | no |
| single-value-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 6192 | 3911 | 0.65× [0.62, 0.68] | -54.5% | [-62.2%, -46.9%] | 7.5 pts | 2.8 | no | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | btree-map | 8 | 6300 | 2322 | 0.38× [0.37, 0.39] | -162.5% | [-170.2%, -154.7%] | 11.5 pts | 2.5 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | baseline | 8 | 23.5 µs | 13.2 µs | 0.57× [0.56, 0.58] | -75.9% | [-79.5%, -72.3%] | 4.3 pts | 2.0 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | btree-map | 8 | 22.6 µs | 7866 | 0.34× [0.34, 0.35] | -190.1% | [-196.2%, -184.1%] | 6.2 pts | 1.5 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | baseline | 8 | 405 | 385 | 0.96× [0.95, 0.97] | -4.1% | [-5.6%, -2.6%] | 2.0 pts | 1.3 | yes | yes |
| single-value-str | street | 212449 | churn | ordered | btree-map | 8 | 515 | 504 | 1.01× [1.00, 1.02] | +0.6% | [-0.4%, +1.5%] | 1.6 pts | 1.1 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.12% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.59% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of +1.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the pooled interval [-9.52%, 2.43%] includes zero
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: 2 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs baseline: the A/A validations found a systematic difference of -2.48% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 prefix: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 churn: ordered vs btree-map: the pooled interval [-0.40%, 1.54%] includes zero
