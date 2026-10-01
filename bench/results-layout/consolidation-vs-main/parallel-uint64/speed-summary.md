| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | path | 262144 | valuesFor | ordered | baseline | 8 | 577 | 634 | 1.10× [1.09, 1.11] | +8.9% | [+8.1%, +9.6%] | 0.9 pts | 1.3 | yes | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 8 | 12.7 µs | 14.1 µs | 1.12× [1.11, 1.13] | +10.5% | [+9.8%, +11.2%] | 0.8 pts | 1.9 | yes | yes |
| multi | path | 262144 | prefix | ordered | baseline | 8 | 13.6 µs | 14.8 µs | 1.12× [1.10, 1.13] | +10.4% | [+9.0%, +11.7%] | 1.6 pts | 0.2 | yes | yes |
| multi | path | 262144 | churn | ordered | baseline | 8 | 917 | 945 | 1.04× [1.02, 1.05] | +3.5% | [+2.4%, +4.5%] | 1.2 pts | 1.4 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 32 | 179 | 177 | 1.00× [1.00, 1.00] | -0.0% | [-0.5%, +0.4%] | 1.5 pts | 3.0 | yes | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 32 | 8087 | 9385 | 1.16× [1.15, 1.17] | +13.8% | [+13.3%, +14.2%] | 1.0 pts | 2.9 | yes | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 32 | 413 | 366 | 0.88× [0.87, 0.90] | -13.2% | [-15.1%, -11.4%] | 4.1 pts | 1.3 | yes | yes |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 32 | 222 | 220 | 0.99× [0.99, 0.99] | -1.1% | [-1.4%, -0.7%] | 0.8 pts | 1.6 | yes | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 32 | 7312 | 8590 | 1.18× [1.18, 1.18] | +15.1% | [+14.9%, +15.3%] | 0.6 pts | 2.2 | yes | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 32 | 556 | 499 | 0.89× [0.88, 0.91] | -12.0% | [-13.8%, -10.2%] | 3.7 pts | 1.7 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 8 | 524 | 584 | 1.10× [1.08, 1.12] | +9.1% | [+7.5%, +10.8%] | 2.0 pts | 3.0 | yes | yes |
| multi | url | 262144 | valuesBetween | ordered | baseline | 8 | 12.6 µs | 13.8 µs | 1.09× [1.09, 1.10] | +8.6% | [+8.0%, +9.1%] | 0.7 pts | 1.9 | yes | yes |
| multi | url | 262144 | prefix | ordered | baseline | 8 | 3452 | 3834 | 1.10× [1.08, 1.12] | +9.1% | [+7.6%, +10.5%] | 1.8 pts | 1.3 | yes | yes |
| multi | url | 262144 | churn | ordered | baseline | 8 | 819 | 821 | 1.00× [0.99, 1.01] | -0.3% | [-1.4%, +0.8%] | 1.3 pts | 1.0 | yes | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 8 | 742 | 805 | 1.08× [1.07, 1.09] | +7.3% | [+6.1%, +8.5%] | 1.4 pts | 2.2 | yes | yes |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 8 | 13.5 µs | 14.8 µs | 1.10× [1.10, 1.11] | +9.5% | [+9.1%, +9.9%] | 0.5 pts | 1.7 | yes | yes |
| multi | url | 1048576 | prefix | ordered | baseline | 8 | 14.8 µs | 16.7 µs | 1.15× [1.14, 1.16] | +12.7% | [+11.9%, +13.5%] | 0.9 pts | 1.1 | yes | yes |
| multi | url | 1048576 | churn | ordered | baseline | 8 | 1265 | 1285 | 1.02× [1.00, 1.03] | +1.7% | [+0.5%, +2.9%] | 1.5 pts | 1.4 | yes | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 24 | 326 | 333 | 1.02× [1.02, 1.03] | +2.4% | [+1.7%, +3.0%] | 1.1 pts | 2.1 | yes | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 24 | 9854 | 11.4 µs | 1.15× [1.15, 1.16] | +13.2% | [+12.9%, +13.5%] | 0.7 pts | 2.4 | yes | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 24 | 707 | 799 | 1.13× [1.13, 1.13] | +11.5% | [+11.3%, +11.7%] | 0.8 pts | 2.0 | yes | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 24 | 582 | 550 | 0.92× [0.91, 0.94] | -8.3% | [-10.2%, -6.5%] | 3.8 pts | 1.5 | yes | yes |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 40 | 432 | 437 | 1.02× [1.02, 1.03] | +2.1% | [+1.8%, +2.4%] | 1.1 pts | 2.2 | yes | yes |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 40 | 10.7 µs | 12.5 µs | 1.17× [1.16, 1.17] | +14.3% | [+14.1%, +14.5%] | 0.6 pts | 2.3 | yes | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 40 | 2029 | 2356 | 1.17× [1.16, 1.17] | +14.3% | [+13.9%, +14.7%] | 0.8 pts | 2.1 | yes | yes |
| multi | uuid | 1048576 | churn | ordered | baseline | 40 | 835 | 821 | 0.96× [0.94, 0.98] | -4.2% | [-6.0%, -2.3%] | 3.7 pts | 2.1 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 24 | 526 | 581 | 1.11× [1.10, 1.12] | +9.6% | [+8.8%, +10.5%] | 1.3 pts | 1.4 | yes | yes |
| unique | path | 262144 | valuesBetween | ordered | baseline | 24 | 8948 | 9274 | 1.04× [1.03, 1.05] | +3.6% | [+2.8%, +4.5%] | 1.6 pts | 3.0 | yes | yes |
| unique | path | 262144 | prefix | ordered | baseline | 24 | 8132 | 8510 | 1.05× [1.03, 1.06] | +4.4% | [+2.7%, +6.1%] | 3.9 pts | 1.0 | yes | yes |
| unique | path | 262144 | churn | ordered | baseline | 24 | 841 | 904 | 1.08× [1.07, 1.08] | +7.0% | [+6.3%, +7.8%] | 1.3 pts | 1.8 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 16 | 120 | 121 | 1.01× [1.01, 1.02] | +1.4% | [+1.0%, +1.7%] | 1.0 pts | 1.6 | yes | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 16 | 4341 | 4720 | 1.09× [1.07, 1.11] | +8.4% | [+6.9%, +9.9%] | 2.2 pts | 3.0 | yes | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 16 | 324 | 316 | 1.00× [0.99, 1.01] | -0.3% | [-1.4%, +0.7%] | 2.0 pts | 1.1 | yes | no |
| unique | u64 | 1048576 | valuesFor | ordered | baseline | 8 | 156 | 158 | 1.02× [1.01, 1.02] | +1.6% | [+0.7%, +2.4%] | 1.0 pts | 2.2 | yes | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 8 | 3517 | 3544 | 1.01× [1.00, 1.01] | +0.6% | [-0.1%, +1.3%] | 0.8 pts | 3.3 | yes | no |
| unique | u64 | 1048576 | churn | ordered | baseline | 8 | 409 | 412 | 0.99× [0.98, 1.00] | -1.0% | [-1.9%, -0.1%] | 1.1 pts | 1.0 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 8 | 475 | 519 | 1.10× [1.08, 1.11] | +8.7% | [+7.3%, +10.1%] | 1.6 pts | 2.1 | yes | yes |
| unique | url | 262144 | valuesBetween | ordered | baseline | 8 | 8764 | 8938 | 1.02× [1.01, 1.03] | +2.1% | [+1.4%, +2.8%] | 0.8 pts | 1.6 | yes | yes |
| unique | url | 262144 | prefix | ordered | baseline | 8 | 2015 | 2154 | 1.06× [1.05, 1.07] | +5.6% | [+4.6%, +6.7%] | 1.2 pts | 1.0 | yes | yes |
| unique | url | 262144 | churn | ordered | baseline | 8 | 736 | 776 | 1.06× [1.05, 1.08] | +5.9% | [+4.6%, +7.2%] | 1.6 pts | 2.0 | yes | yes |
| unique | url | 1048576 | valuesFor | ordered | baseline | 8 | 660 | 720 | 1.08× [1.06, 1.10] | +7.5% | [+5.6%, +9.5%] | 2.3 pts | 2.4 | yes | yes |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 8 | 9495 | 9703 | 1.02× [1.01, 1.03] | +2.1% | [+1.3%, +2.9%] | 1.0 pts | 2.6 | yes | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 8 | 8588 | 9226 | 1.09× [1.07, 1.12] | +8.5% | [+6.5%, +10.4%] | 2.3 pts | 0.6 | yes | yes |
| unique | url | 1048576 | churn | ordered | baseline | 8 | 1019 | 1064 | 1.07× [1.05, 1.08] | +6.4% | [+5.0%, +7.7%] | 1.6 pts | 1.7 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 8 | 265 | 262 | 1.00× [0.99, 1.01] | -0.1% | [-1.1%, +0.9%] | 1.2 pts | 2.0 | yes | no |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 6221 | 6596 | 1.06× [1.05, 1.07] | +6.0% | [+5.2%, +6.9%] | 1.0 pts | 2.4 | yes | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 8 | 503 | 529 | 1.05× [1.04, 1.06] | +4.8% | [+3.8%, +5.7%] | 1.1 pts | 1.7 | yes | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 8 | 444 | 444 | 1.02× [1.01, 1.03] | +2.1% | [+1.1%, +3.2%] | 1.2 pts | 1.0 | yes | yes |
| unique | uuid | 1048576 | valuesFor | ordered | baseline | 8 | 357 | 352 | 0.99× [0.99, 1.00] | -0.8% | [-1.2%, -0.3%] | 0.6 pts | 1.4 | yes | yes |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 8 | 7117 | 7437 | 1.05× [1.04, 1.05] | +4.6% | [+4.0%, +5.2%] | 0.7 pts | 2.4 | yes | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 8 | 1363 | 1445 | 1.06× [1.05, 1.07] | +5.7% | [+5.1%, +6.4%] | 0.7 pts | 2.7 | yes | yes |
| unique | uuid | 1048576 | churn | ordered | baseline | 8 | 610 | 622 | 1.02× [1.00, 1.04] | +2.3% | [+0.3%, +4.2%] | 2.3 pts | 2.6 | yes | yes |

Regime: parallel: 8 processes at a time, each with GOMAXPROCS 3, sharing the caches and the memory bandwidth.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi u64 n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.01% does not clear the 0.04% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.48%, 0.45%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs baseline: 7 processes resolved A as faster and 8 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 churn: ordered vs baseline: the pooled interval [-1.42%, 0.78%] includes zero
- multi url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 prefix: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.02% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 churn: ordered vs baseline: the pooled interval [-1.39%, 0.74%] includes zero
- unique u64 n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.10%, 1.31%] includes zero
- unique u64 n=1048576 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 valuesBetween: ordered vs baseline: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.15%, 0.86%] includes zero
- unique uuid n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
