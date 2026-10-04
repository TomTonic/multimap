| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | baseline | 12 | 54.0 | 53.7 | 0.99× [0.97, 1.00] | -1.2% | [-2.6%, +0.2%] | 1.7 pts | 1.9 | yes | no |
| multi-str | street | 4096 | valuesFor | ordered | btree-sets | 12 | 54.9 | 138 | 2.52× [2.48, 2.56] | +60.3% | [+59.7%, +61.0%] | 1.0 pts | 2.8 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | hashed | 12 | 54.5 | 24.3 | 0.45× [0.45, 0.46] | -121.7% | [-123.8%, -119.5%] | 5.7 pts | 3.9 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | map-sets | 12 | 54.4 | 67.5 | 1.24× [1.23, 1.26] | +19.6% | [+18.5%, +20.7%] | 1.6 pts | 2.7 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | baseline | 12 | 2391 | 2434 | 1.02× [1.00, 1.03] | +1.5% | [+0.3%, +2.7%] | 1.9 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | btree-sets | 12 | 2378 | 7042 | 2.99× [2.96, 3.02] | +66.5% | [+66.2%, +66.9%] | 0.7 pts | 1.7 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | hashed | 12 | 2497 | 54.1 µs | 21.69× [21.36, 22.03] | +95.4% | [+95.3%, +95.5%] | 0.1 pts | 1.5 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | map-sets | 12 | 2462 | 57.4 µs | 23.38× [23.00, 23.76] | +95.7% | [+95.7%, +95.8%] | 0.1 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | baseline | 12 | 228 | 231 | 1.01× [0.99, 1.03] | +1.0% | [-0.9%, +3.0%] | 2.2 pts | 1.9 | yes | no |
| multi-str | street | 4096 | prefix | ordered | btree-sets | 12 | 229 | 713 | 3.11× [3.02, 3.20] | +67.8% | [+66.8%, +68.8%] | 1.1 pts | 2.3 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | hashed | 12 | 264 | 49.7 µs | 190.04× [187.40, 192.75] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | map-sets | 12 | 261 | 47.1 µs | 183.26× [174.50, 192.95] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | churn | ordered | baseline | 12 | 109 | 106 | 0.98× [0.97, 0.99] | -2.2% | [-3.2%, -1.2%] | 1.1 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | churn | ordered | btree-sets | 12 | 109 | 200 | 1.84× [1.82, 1.86] | +45.7% | [+45.2%, +46.3%] | 0.7 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | churn | ordered | hashed | 12 | 107 | 56.8 | 0.53× [0.53, 0.54] | -88.0% | [-89.7%, -86.4%] | 3.0 pts | 2.1 | yes | yes |
| multi-str | street | 4096 | churn | ordered | map-sets | 12 | 108 | 72.9 | 0.67× [0.66, 0.69] | -48.5% | [-51.7%, -45.3%] | 3.8 pts | 2.3 | yes | yes |
| multi-str | street | 4096 | build | ordered | baseline | 12 | 4.43 ms | 4.34 ms | 0.97× [0.97, 0.98] | -2.7% | [-3.3%, -2.2%] | 1.0 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | build | ordered | btree-sets | 12 | 4.49 ms | 8.08 ms | 1.79× [1.77, 1.81] | +44.1% | [+43.4%, +44.7%] | 0.9 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | build | ordered | hashed | 12 | 4.45 ms | 2.75 ms | 0.62× [0.61, 0.64] | -60.0% | [-62.8%, -57.2%] | 4.6 pts | 1.8 | yes | yes |
| multi-str | street | 4096 | build | ordered | map-sets | 12 | 4.48 ms | 3.33 ms | 0.75× [0.73, 0.77] | -32.8% | [-36.3%, -29.3%] | 5.8 pts | 1.7 | no | yes |
| multi-str | street | 16384 | valuesFor | ordered | baseline | 12 | 75.7 | 75.4 | 0.99× [0.98, 1.01] | -0.7% | [-2.0%, +0.6%] | 1.3 pts | 1.8 | yes | no |
| multi-str | street | 16384 | valuesFor | ordered | btree-sets | 12 | 77.5 | 187 | 2.42× [2.39, 2.45] | +58.7% | [+58.1%, +59.2%] | 0.6 pts | 2.0 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | hashed | 12 | 76.2 | 30.1 | 0.39× [0.39, 0.40] | -153.2% | [-155.9%, -150.6%] | 3.4 pts | 2.1 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | map-sets | 12 | 75.7 | 79.5 | 1.06× [1.05, 1.07] | +5.4% | [+4.5%, +6.2%] | 1.1 pts | 2.0 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | baseline | 12 | 2963 | 3018 | 1.02× [1.01, 1.02] | +1.6% | [+0.7%, +2.4%] | 0.9 pts | 1.3 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | btree-sets | 12 | 2974 | 8241 | 2.76× [2.75, 2.78] | +63.8% | [+63.6%, +64.0%] | 0.4 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | hashed | 12 | 3663 | 233.4 µs | 62.02× [59.15, 65.19] | +98.4% | [+98.3%, +98.5%] | 0.1 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | map-sets | 12 | 3802 | 233.2 µs | 59.67× [58.79, 60.57] | +98.3% | [+98.3%, +98.3%] | 0.1 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | baseline | 12 | 853 | 874 | 1.02× [1.01, 1.03] | +2.3% | [+1.4%, +3.2%] | 1.6 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | btree-sets | 12 | 884 | 2994 | 3.40× [3.34, 3.47] | +70.6% | [+70.0%, +71.2%] | 0.7 pts | 1.4 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | hashed | 12 | 1214 | 231.0 µs | 187.95× [180.84, 195.65] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | map-sets | 12 | 1217 | 223.0 µs | 183.84× [181.64, 186.08] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | churn | ordered | baseline | 12 | 151 | 147 | 0.98× [0.97, 0.98] | -2.4% | [-2.8%, -1.9%] | 0.6 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | churn | ordered | btree-sets | 12 | 159 | 296 | 1.84× [1.82, 1.85] | +45.5% | [+45.1%, +46.0%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | churn | ordered | hashed | 12 | 145 | 74.1 | 0.51× [0.51, 0.52] | -95.0% | [-96.6%, -93.5%] | 2.9 pts | 1.5 | yes | yes |
| multi-str | street | 16384 | churn | ordered | map-sets | 12 | 151 | 102 | 0.68× [0.67, 0.69] | -46.9% | [-48.6%, -45.2%] | 6.2 pts | 2.2 | yes | yes |
| multi-str | street | 16384 | build | ordered | baseline | 12 | 24.15 ms | 23.62 ms | 0.97× [0.96, 0.98] | -2.7% | [-3.7%, -1.8%] | 1.2 pts | 1.4 | yes | yes |
| multi-str | street | 16384 | build | ordered | btree-sets | 12 | 24.38 ms | 45.19 ms | 1.84× [1.82, 1.86] | +45.6% | [+44.9%, +46.2%] | 0.7 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | build | ordered | hashed | 12 | 23.96 ms | 13.98 ms | 0.58× [0.57, 0.59] | -72.2% | [-74.6%, -69.8%] | 4.1 pts | 1.7 | yes | yes |
| multi-str | street | 16384 | build | ordered | map-sets | 12 | 24.36 ms | 18.31 ms | 0.74× [0.72, 0.77] | -34.7% | [-39.3%, -30.0%] | 6.6 pts | 2.1 | no | yes |
| multi-str | street | 212449 | valuesFor | ordered | baseline | 12 | 263 | 256 | 0.99× [0.97, 1.01] | -0.8% | [-3.0%, +1.3%] | 2.7 pts | 2.2 | no | no |
| multi-str | street | 212449 | valuesFor | ordered | btree-sets | 12 | 311 | 529 | 1.73× [1.70, 1.76] | +42.1% | [+41.1%, +43.1%] | 1.9 pts | 2.1 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | hashed | 12 | 233 | 117 | 0.50× [0.49, 0.51] | -101.8% | [-106.1%, -97.4%] | 6.9 pts | 2.7 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | map-sets | 12 | 266 | 315 | 1.17× [1.13, 1.20] | +14.3% | [+11.7%, +16.9%] | 3.6 pts | 3.8 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | baseline | 12 | 7365 | 7209 | 1.00× [0.97, 1.02] | -0.2% | [-2.8%, +2.4%] | 3.6 pts | 4.2 | no | no |
| multi-str | street | 212449 | valuesBetween | ordered | btree-sets | 12 | 8250 | 24.8 µs | 2.97× [2.94, 2.99] | +66.3% | [+66.0%, +66.5%] | 0.6 pts | 1.2 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | baseline | 12 | 22.1 µs | 22.8 µs | 0.99× [0.95, 1.02] | -1.2% | [-4.8%, +2.3%] | 4.9 pts | 2.3 | no | no |
| multi-str | street | 212449 | prefix | ordered | btree-sets | 12 | 26.9 µs | 81.4 µs | 3.07× [2.98, 3.16] | +67.4% | [+66.5%, +68.4%] | 1.1 pts | 0.8 | yes | yes |
| multi-str | street | 212449 | churn | ordered | baseline | 12 | 535 | 525 | 0.97× [0.95, 0.99] | -3.2% | [-5.8%, -0.6%] | 2.9 pts | 0.7 | no | yes |
| multi-str | street | 212449 | churn | ordered | btree-sets | 12 | 594 | 789 | 1.34× [1.32, 1.36] | +25.6% | [+24.5%, +26.6%] | 1.6 pts | 1.8 | yes | yes |
| multi-str | street | 212449 | churn | ordered | hashed | 12 | 488 | 288 | 0.57× [0.56, 0.57] | -76.3% | [-77.8%, -74.8%] | 10.0 pts | 1.4 | yes | yes |
| multi-str | street | 212449 | churn | ordered | map-sets | 12 | 505 | 363 | 0.72× [0.71, 0.74] | -38.6% | [-41.2%, -35.9%] | 4.1 pts | 1.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.58%, 0.18%] includes zero
- multi-str street n=4096 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 valuesFor: ordered vs hashed: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 valuesFor: ordered vs map-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 prefix: ordered vs baseline: the pooled interval [-0.92%, 3.01%] includes zero
- multi-str street n=4096 prefix: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 prefix: ordered vs hashed: the A/A validations found a systematic difference of +1.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +2.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 churn: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 churn: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of -0.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.95%, 0.64%] includes zero
- multi-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesFor: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesFor: ordered vs map-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs hashed: the A/A validations found a systematic difference of +3.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesBetween: ordered vs map-sets: the A/A validations found a systematic difference of +4.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +4.05% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +3.00% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 churn: ordered vs map-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 build: ordered vs map-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs baseline: the pooled interval [-2.97%, 1.33%] includes zero
- multi-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs hashed: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs map-sets: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs baseline: the pooled difference of -0.18% does not clear the 0.50% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 valuesBetween: ordered vs baseline: the pooled interval [-2.78%, 2.43%] includes zero
- multi-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs baseline: 4 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 prefix: ordered vs baseline: the pooled difference of -1.24% does not clear the 1.63% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 prefix: ordered vs baseline: the pooled interval [-4.77%, 2.29%] includes zero
- multi-str street n=212449 prefix: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
