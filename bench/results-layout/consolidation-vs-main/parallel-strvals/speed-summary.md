| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | path | 262144 | valuesFor | ordered | baseline | 8 | 654 | 710 | 1.07× [1.05, 1.09] | +6.8% | [+4.9%, +8.6%] | 2.2 pts | 3.3 | yes | yes |
| multi-str | path | 262144 | valuesBetween | ordered | baseline | 8 | 16.2 µs | 16.2 µs | 0.99× [0.98, 1.01] | -0.6% | [-2.0%, +0.8%] | 1.7 pts | 4.6 | yes | no |
| multi-str | path | 262144 | prefix | ordered | baseline | 8 | 18.0 µs | 17.4 µs | 0.98× [0.97, 1.00] | -1.9% | [-3.5%, -0.3%] | 1.9 pts | 1.8 | yes | yes |
| multi-str | path | 262144 | churn | ordered | baseline | 8 | 1029 | 1086 | 1.05× [1.04, 1.07] | +5.2% | [+4.1%, +6.3%] | 1.3 pts | 2.1 | yes | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | baseline | 16 | 202 | 203 | 1.02× [1.01, 1.03] | +2.1% | [+1.2%, +3.0%] | 1.5 pts | 3.3 | yes | yes |
| multi-str | u64 | 262144 | valuesBetween | ordered | baseline | 16 | 10.7 µs | 10.6 µs | 0.99× [0.98, 1.01] | -0.6% | [-1.6%, +0.5%] | 1.8 pts | 5.8 | yes | no |
| multi-str | u64 | 262144 | churn | ordered | baseline | 16 | 525 | 507 | 0.98× [0.96, 1.00] | -2.3% | [-4.3%, -0.4%] | 4.5 pts | 1.4 | yes | yes |
| multi-str | u64 | 1048576 | valuesFor | ordered | baseline | 8 | 249 | 252 | 1.01× [1.00, 1.02] | +1.3% | [+0.1%, +2.4%] | 1.4 pts | 2.6 | yes | yes |
| multi-str | u64 | 1048576 | valuesBetween | ordered | baseline | 8 | 10.1 µs | 9905 | 0.98× [0.97, 1.00] | -2.0% | [-3.6%, -0.5%] | 1.9 pts | 6.0 | yes | yes |
| multi-str | u64 | 1048576 | churn | ordered | baseline | 8 | 638 | 612 | 0.96× [0.94, 0.97] | -4.6% | [-6.1%, -3.1%] | 1.8 pts | 0.9 | yes | yes |
| multi-str | url | 262144 | valuesFor | ordered | baseline | 8 | 610 | 638 | 1.07× [1.05, 1.08] | +6.2% | [+4.8%, +7.7%] | 1.7 pts | 2.6 | yes | yes |
| multi-str | url | 262144 | valuesBetween | ordered | baseline | 8 | 16.3 µs | 15.5 µs | 0.95× [0.95, 0.95] | -5.3% | [-5.7%, -4.8%] | 0.5 pts | 1.7 | yes | yes |
| multi-str | url | 262144 | prefix | ordered | baseline | 8 | 4668 | 4585 | 0.98× [0.97, 0.99] | -2.2% | [-3.1%, -1.4%] | 1.0 pts | 0.5 | yes | yes |
| multi-str | url | 262144 | churn | ordered | baseline | 8 | 911 | 935 | 1.02× [1.01, 1.04] | +2.4% | [+1.2%, +3.7%] | 1.5 pts | 1.4 | yes | yes |
| multi-str | url | 1048576 | valuesFor | ordered | baseline | 8 | 861 | 933 | 1.06× [1.06, 1.07] | +6.1% | [+5.3%, +6.9%] | 1.0 pts | 1.2 | yes | yes |
| multi-str | url | 1048576 | valuesBetween | ordered | baseline | 8 | 18.3 µs | 17.5 µs | 0.96× [0.95, 0.97] | -4.1% | [-5.0%, -3.3%] | 1.0 pts | 2.2 | yes | yes |
| multi-str | url | 1048576 | prefix | ordered | baseline | 8 | 21.6 µs | 21.5 µs | 1.00× [0.98, 1.02] | +0.1% | [-1.6%, +1.8%] | 2.1 pts | 2.0 | yes | no |
| multi-str | url | 1048576 | churn | ordered | baseline | 8 | 1319 | 1345 | 1.02× [1.00, 1.04] | +2.0% | [+0.3%, +3.7%] | 2.0 pts | 2.1 | yes | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | baseline | 16 | 375 | 358 | 0.96× [0.95, 0.97] | -4.1% | [-4.9%, -3.2%] | 1.7 pts | 3.1 | yes | yes |
| multi-str | uuid | 262144 | valuesBetween | ordered | baseline | 16 | 13.1 µs | 12.3 µs | 0.94× [0.93, 0.94] | -6.9% | [-7.7%, -6.2%] | 1.5 pts | 3.3 | yes | yes |
| multi-str | uuid | 262144 | prefix | ordered | baseline | 16 | 903 | 868 | 0.97× [0.96, 0.97] | -3.6% | [-4.5%, -2.6%] | 1.7 pts | 4.3 | yes | yes |
| multi-str | uuid | 262144 | churn | ordered | baseline | 16 | 690 | 672 | 0.97× [0.96, 0.99] | -2.9% | [-4.4%, -1.5%] | 2.7 pts | 1.3 | yes | yes |
| multi-str | uuid | 1048576 | valuesFor | ordered | baseline | 16 | 526 | 503 | 0.97× [0.95, 0.98] | -3.4% | [-5.0%, -1.7%] | 2.3 pts | 4.0 | yes | yes |
| multi-str | uuid | 1048576 | valuesBetween | ordered | baseline | 16 | 14.4 µs | 13.5 µs | 0.92× [0.91, 0.93] | -8.4% | [-9.6%, -7.1%] | 2.2 pts | 4.0 | yes | yes |
| multi-str | uuid | 1048576 | prefix | ordered | baseline | 16 | 2866 | 2638 | 0.93× [0.91, 0.94] | -7.9% | [-9.5%, -6.3%] | 2.5 pts | 4.7 | yes | yes |
| multi-str | uuid | 1048576 | churn | ordered | baseline | 16 | 930 | 918 | 0.98× [0.97, 0.99] | -2.3% | [-3.4%, -1.3%] | 2.0 pts | 1.1 | yes | yes |
| unique-str | path | 262144 | valuesFor | ordered | baseline | 24 | 563 | 621 | 1.10× [1.09, 1.10] | +8.9% | [+8.5%, +9.3%] | 1.1 pts | 1.2 | yes | yes |
| unique-str | path | 262144 | valuesBetween | ordered | baseline | 24 | 10.6 µs | 10.5 µs | 0.99× [0.99, 1.00] | -0.9% | [-1.4%, -0.3%] | 1.1 pts | 1.9 | yes | yes |
| unique-str | path | 262144 | prefix | ordered | baseline | 24 | 10.4 µs | 10.1 µs | 0.96× [0.94, 0.98] | -4.2% | [-5.9%, -2.5%] | 2.9 pts | 0.9 | yes | yes |
| unique-str | path | 262144 | churn | ordered | baseline | 24 | 894 | 958 | 1.08× [1.07, 1.08] | +7.0% | [+6.5%, +7.5%] | 1.3 pts | 1.6 | yes | yes |
| unique-str | u64 | 262144 | valuesFor | ordered | baseline | 8 | 131 | 133 | 1.03× [1.01, 1.04] | +2.5% | [+0.9%, +4.0%] | 1.9 pts | 3.6 | yes | yes |
| unique-str | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 5252 | 5348 | 1.02× [1.01, 1.03] | +2.1% | [+0.8%, +3.4%] | 1.5 pts | 2.7 | yes | yes |
| unique-str | u64 | 262144 | churn | ordered | baseline | 8 | 359 | 362 | 1.00× [0.99, 1.01] | -0.1% | [-1.0%, +0.8%] | 1.1 pts | 0.7 | yes | no |
| unique-str | u64 | 1048576 | valuesFor | ordered | baseline | 8 | 167 | 169 | 1.01× [1.00, 1.01] | +0.8% | [+0.2%, +1.4%] | 0.7 pts | 2.0 | yes | yes |
| unique-str | u64 | 1048576 | valuesBetween | ordered | baseline | 8 | 4258 | 4111 | 0.96× [0.96, 0.97] | -3.8% | [-4.3%, -3.3%] | 0.6 pts | 1.7 | yes | yes |
| unique-str | u64 | 1048576 | churn | ordered | baseline | 8 | 449 | 443 | 0.99× [0.97, 1.00] | -1.3% | [-2.8%, +0.1%] | 1.7 pts | 1.9 | yes | no |
| unique-str | url | 262144 | valuesFor | ordered | baseline | 8 | 522 | 571 | 1.10× [1.09, 1.10] | +9.1% | [+8.6%, +9.5%] | 0.5 pts | 0.8 | yes | yes |
| unique-str | url | 262144 | valuesBetween | ordered | baseline | 8 | 10.7 µs | 10.0 µs | 0.94× [0.93, 0.95] | -6.4% | [-7.1%, -5.7%] | 0.9 pts | 2.5 | yes | yes |
| unique-str | url | 262144 | prefix | ordered | baseline | 8 | 2481 | 2396 | 0.97× [0.96, 0.98] | -3.5% | [-4.4%, -2.5%] | 1.1 pts | 0.7 | yes | yes |
| unique-str | url | 262144 | churn | ordered | baseline | 8 | 789 | 828 | 1.05× [1.04, 1.06] | +4.9% | [+3.9%, +5.9%] | 1.2 pts | 1.9 | yes | yes |
| unique-str | url | 1048576 | valuesFor | ordered | baseline | 16 | 720 | 778 | 1.08× [1.07, 1.09] | +7.3% | [+6.5%, +8.1%] | 1.2 pts | 1.6 | yes | yes |
| unique-str | url | 1048576 | valuesBetween | ordered | baseline | 16 | 11.4 µs | 10.9 µs | 0.95× [0.95, 0.96] | -4.9% | [-5.4%, -4.3%] | 1.0 pts | 2.2 | yes | yes |
| unique-str | url | 1048576 | prefix | ordered | baseline | 16 | 11.3 µs | 10.8 µs | 0.97× [0.95, 0.98] | -3.2% | [-4.7%, -1.6%] | 4.0 pts | 1.1 | yes | yes |
| unique-str | url | 1048576 | churn | ordered | baseline | 16 | 1086 | 1138 | 1.05× [1.04, 1.06] | +4.7% | [+3.4%, +6.0%] | 1.9 pts | 2.2 | yes | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | baseline | 8 | 297 | 283 | 0.96× [0.95, 0.97] | -4.2% | [-4.8%, -3.6%] | 0.7 pts | 1.1 | yes | yes |
| unique-str | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 7440 | 6782 | 0.91× [0.90, 0.91] | -10.4% | [-11.4%, -9.5%] | 1.1 pts | 1.9 | yes | yes |
| unique-str | uuid | 262144 | prefix | ordered | baseline | 8 | 566 | 538 | 0.95× [0.94, 0.96] | -5.2% | [-6.1%, -4.3%] | 1.1 pts | 2.0 | yes | yes |
| unique-str | uuid | 262144 | churn | ordered | baseline | 8 | 471 | 474 | 1.00× [0.99, 1.01] | -0.1% | [-0.8%, +0.5%] | 0.8 pts | 0.7 | yes | no |
| unique-str | uuid | 1048576 | valuesFor | ordered | baseline | 8 | 398 | 380 | 0.96× [0.95, 0.97] | -3.9% | [-4.8%, -3.0%] | 1.1 pts | 1.7 | yes | yes |
| unique-str | uuid | 1048576 | valuesBetween | ordered | baseline | 8 | 8379 | 7371 | 0.88× [0.87, 0.88] | -13.7% | [-14.4%, -13.1%] | 0.8 pts | 2.8 | yes | yes |
| unique-str | uuid | 1048576 | prefix | ordered | baseline | 8 | 1602 | 1447 | 0.90× [0.90, 0.91] | -10.8% | [-11.5%, -10.2%] | 0.8 pts | 1.4 | yes | yes |
| unique-str | uuid | 1048576 | churn | ordered | baseline | 8 | 691 | 700 | 1.00× [0.98, 1.01] | -0.3% | [-1.9%, +1.2%] | 1.8 pts | 2.1 | yes | no |

Regime: parallel: 8 processes at a time, each with GOMAXPROCS 3, sharing the caches and the memory bandwidth.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str path n=262144 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.03%, 0.80%] includes zero
- multi-str path n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str path n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.89% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str path n=262144 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.64%, 0.50%] includes zero
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 5.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: 4 processes resolved A as faster and 8 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str u64 n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=1048576 valuesBetween: ordered vs baseline: the processes scatter 6.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 prefix: ordered vs baseline: the pooled difference of 0.07% does not clear the 0.72% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=1048576 prefix: ordered vs baseline: the pooled interval [-1.65%, 1.80%] includes zero
- multi-str url n=1048576 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str uuid n=1048576 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 valuesBetween: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 prefix: ordered vs baseline: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=262144 churn: ordered vs baseline: the pooled difference of -0.10% does not clear the 0.20% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=262144 churn: ordered vs baseline: the pooled interval [-1.01%, 0.81%] includes zero
- unique-str u64 n=1048576 churn: ordered vs baseline: the pooled interval [-2.80%, 0.11%] includes zero
- unique-str url n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=1048576 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.04% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str url n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=1048576 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 churn: ordered vs baseline: the pooled difference of -0.12% does not clear the 0.19% noise floor, the bound on what the harness reports between identical code in every process
- unique-str uuid n=262144 churn: ordered vs baseline: the pooled interval [-0.76%, 0.52%] includes zero
- unique-str uuid n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=1048576 churn: ordered vs baseline: the pooled interval [-1.86%, 1.19%] includes zero
- unique-str uuid n=1048576 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
