| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | path | 262144 | valuesFor | ordered | baseline | 16 | 627 | 699 | 1.11× [1.09, 1.13] | +10.0% | [+8.6%, +11.4%] | 2.2 pts | 3.5 | yes | yes |
| multi-str | path | 262144 | valuesBetween | ordered | baseline | 16 | 14.4 µs | 16.0 µs | 1.12× [1.10, 1.13] | +10.4% | [+9.4%, +11.3%] | 1.6 pts | 3.5 | yes | yes |
| multi-str | path | 262144 | prefix | ordered | baseline | 16 | 15.1 µs | 17.3 µs | 1.12× [1.10, 1.14] | +10.8% | [+9.1%, +12.5%] | 3.2 pts | 2.1 | yes | yes |
| multi-str | path | 262144 | churn | ordered | baseline | 16 | 1071 | 1105 | 1.04× [1.03, 1.04] | +3.4% | [+2.7%, +4.1%] | 1.2 pts | 1.7 | yes | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | baseline | 32 | 205 | 203 | 0.99× [0.98, 1.00] | -1.0% | [-1.9%, -0.1%] | 2.0 pts | 3.4 | yes | yes |
| multi-str | u64 | 262144 | valuesBetween | ordered | baseline | 32 | 9631 | 10.6 µs | 1.10× [1.09, 1.12] | +9.5% | [+8.6%, +10.3%] | 1.7 pts | 5.3 | yes | yes |
| multi-str | u64 | 262144 | churn | ordered | baseline | 32 | 536 | 496 | 0.92× [0.90, 0.93] | -8.9% | [-10.7%, -7.0%] | 4.0 pts | 1.0 | yes | yes |
| multi-str | u64 | 1048576 | valuesFor | ordered | baseline | 16 | 264 | 252 | 0.95× [0.94, 0.96] | -5.3% | [-6.7%, -3.8%] | 2.2 pts | 4.4 | yes | yes |
| multi-str | u64 | 1048576 | valuesBetween | ordered | baseline | 16 | 9539 | 9922 | 1.04× [1.03, 1.05] | +3.9% | [+2.9%, +4.9%] | 1.7 pts | 5.2 | yes | yes |
| multi-str | u64 | 1048576 | churn | ordered | baseline | 16 | 717 | 636 | 0.88× [0.87, 0.89] | -13.2% | [-14.7%, -11.8%] | 3.2 pts | 1.7 | yes | yes |
| multi-str | url | 262144 | valuesFor | ordered | baseline | 8 | 584 | 653 | 1.11× [1.10, 1.12] | +10.1% | [+9.1%, +11.1%] | 1.2 pts | 2.1 | yes | yes |
| multi-str | url | 262144 | valuesBetween | ordered | baseline | 8 | 14.9 µs | 15.4 µs | 1.04× [1.04, 1.04] | +4.1% | [+3.9%, +4.2%] | 0.2 pts | 0.5 | yes | yes |
| multi-str | url | 262144 | prefix | ordered | baseline | 8 | 4009 | 4348 | 1.09× [1.08, 1.11] | +8.7% | [+7.4%, +9.9%] | 1.5 pts | 0.9 | yes | yes |
| multi-str | url | 262144 | churn | ordered | baseline | 8 | 1002 | 1013 | 1.00× [0.98, 1.01] | -0.3% | [-1.8%, +1.2%] | 1.8 pts | 2.1 | yes | no |
| multi-str | url | 1048576 | valuesFor | ordered | baseline | 8 | 815 | 907 | 1.10× [1.08, 1.12] | +8.8% | [+7.2%, +10.4%] | 1.9 pts | 3.1 | yes | yes |
| multi-str | url | 1048576 | valuesBetween | ordered | baseline | 8 | 16.5 µs | 17.5 µs | 1.06× [1.05, 1.07] | +5.9% | [+5.0%, +6.9%] | 1.1 pts | 2.7 | yes | yes |
| multi-str | url | 1048576 | prefix | ordered | baseline | 8 | 17.5 µs | 20.3 µs | 1.11× [1.10, 1.13] | +10.3% | [+9.3%, +11.3%] | 1.2 pts | 1.3 | yes | yes |
| multi-str | url | 1048576 | churn | ordered | baseline | 8 | 1310 | 1315 | 1.02× [1.01, 1.03] | +1.7% | [+0.7%, +2.7%] | 1.2 pts | 1.3 | yes | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | baseline | 8 | 332 | 352 | 1.04× [1.03, 1.06] | +4.2% | [+2.5%, +5.9%] | 2.1 pts | 5.1 | yes | yes |
| multi-str | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 11.2 µs | 12.2 µs | 1.10× [1.09, 1.11] | +9.1% | [+8.1%, +10.2%] | 1.3 pts | 3.6 | yes | yes |
| multi-str | uuid | 262144 | prefix | ordered | baseline | 8 | 786 | 863 | 1.10× [1.09, 1.11] | +9.2% | [+8.2%, +10.2%] | 1.2 pts | 4.1 | yes | yes |
| multi-str | uuid | 262144 | churn | ordered | baseline | 8 | 725 | 672 | 0.92× [0.91, 0.94] | -8.3% | [-10.1%, -6.6%] | 2.1 pts | 1.4 | yes | yes |
| multi-str | uuid | 1048576 | valuesFor | ordered | baseline | 16 | 477 | 492 | 1.05× [1.03, 1.07] | +4.7% | [+2.9%, +6.5%] | 3.2 pts | 6.1 | yes | yes |
| multi-str | uuid | 1048576 | valuesBetween | ordered | baseline | 16 | 12.3 µs | 13.1 µs | 1.08× [1.07, 1.09] | +7.7% | [+6.7%, +8.6%] | 1.7 pts | 4.5 | yes | yes |
| multi-str | uuid | 1048576 | prefix | ordered | baseline | 16 | 2354 | 2542 | 1.09× [1.07, 1.10] | +8.1% | [+7.0%, +9.2%] | 1.9 pts | 4.9 | yes | yes |
| multi-str | uuid | 1048576 | churn | ordered | baseline | 16 | 951 | 905 | 0.96× [0.95, 0.97] | -4.3% | [-5.6%, -3.0%] | 1.7 pts | 0.9 | yes | yes |
| unique-str | path | 262144 | valuesFor | ordered | baseline | 8 | 536 | 612 | 1.14× [1.13, 1.16] | +12.7% | [+11.8%, +13.5%] | 1.0 pts | 1.1 | yes | yes |
| unique-str | path | 262144 | valuesBetween | ordered | baseline | 8 | 9085 | 10.3 µs | 1.13× [1.12, 1.14] | +11.5% | [+10.5%, +12.5%] | 1.2 pts | 2.0 | yes | yes |
| unique-str | path | 262144 | prefix | ordered | baseline | 8 | 8486 | 9781 | 1.12× [1.11, 1.14] | +11.1% | [+10.2%, +11.9%] | 1.0 pts | 0.4 | yes | yes |
| unique-str | path | 262144 | churn | ordered | baseline | 8 | 858 | 926 | 1.10× [1.08, 1.11] | +8.9% | [+7.6%, +10.3%] | 1.6 pts | 1.8 | yes | yes |
| unique-str | u64 | 262144 | valuesFor | ordered | baseline | 16 | 121 | 133 | 1.10× [1.09, 1.12] | +9.4% | [+8.4%, +10.4%] | 1.5 pts | 2.5 | yes | yes |
| unique-str | u64 | 262144 | valuesBetween | ordered | baseline | 16 | 4369 | 5220 | 1.20× [1.18, 1.21] | +16.4% | [+15.6%, +17.2%] | 1.9 pts | 3.3 | yes | yes |
| unique-str | u64 | 262144 | churn | ordered | baseline | 16 | 332 | 344 | 1.07× [1.05, 1.09] | +6.4% | [+4.9%, +7.9%] | 2.4 pts | 1.7 | yes | yes |
| unique-str | u64 | 1048576 | valuesFor | ordered | baseline | 8 | 159 | 170 | 1.06× [1.05, 1.07] | +5.9% | [+5.1%, +6.6%] | 0.9 pts | 2.3 | yes | yes |
| unique-str | u64 | 1048576 | valuesBetween | ordered | baseline | 8 | 3578 | 4068 | 1.14× [1.13, 1.14] | +12.0% | [+11.3%, +12.6%] | 0.8 pts | 3.3 | yes | yes |
| unique-str | u64 | 1048576 | churn | ordered | baseline | 8 | 441 | 455 | 1.01× [0.99, 1.03] | +1.3% | [-0.6%, +3.3%] | 2.3 pts | 2.3 | yes | no |
| unique-str | url | 262144 | valuesFor | ordered | baseline | 8 | 477 | 542 | 1.13× [1.11, 1.15] | +11.3% | [+9.7%, +12.8%] | 1.8 pts | 2.7 | yes | yes |
| unique-str | url | 262144 | valuesBetween | ordered | baseline | 8 | 9599 | 9826 | 1.02× [1.01, 1.03] | +2.0% | [+1.3%, +2.7%] | 0.9 pts | 1.7 | yes | yes |
| unique-str | url | 262144 | prefix | ordered | baseline | 8 | 2108 | 2330 | 1.10× [1.08, 1.11] | +8.7% | [+7.5%, +10.0%] | 1.5 pts | 1.3 | yes | yes |
| unique-str | url | 262144 | churn | ordered | baseline | 8 | 746 | 814 | 1.08× [1.06, 1.09] | +7.0% | [+5.8%, +8.2%] | 1.4 pts | 1.6 | yes | yes |
| unique-str | url | 1048576 | valuesFor | ordered | baseline | 16 | 670 | 747 | 1.12× [1.11, 1.13] | +10.8% | [+10.2%, +11.5%] | 1.4 pts | 2.1 | yes | yes |
| unique-str | url | 1048576 | valuesBetween | ordered | baseline | 16 | 10.2 µs | 10.6 µs | 1.05× [1.04, 1.05] | +4.7% | [+4.1%, +5.2%] | 0.7 pts | 1.8 | yes | yes |
| unique-str | url | 1048576 | prefix | ordered | baseline | 16 | 9333 | 10.4 µs | 1.12× [1.09, 1.14] | +10.5% | [+8.5%, +12.5%] | 2.8 pts | 0.8 | yes | yes |
| unique-str | url | 1048576 | churn | ordered | baseline | 16 | 994 | 1088 | 1.07× [1.06, 1.08] | +6.8% | [+5.9%, +7.6%] | 1.1 pts | 1.7 | yes | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | baseline | 8 | 247 | 271 | 1.09× [1.08, 1.11] | +8.6% | [+7.7%, +9.5%] | 1.1 pts | 1.9 | yes | yes |
| unique-str | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 5630 | 6595 | 1.17× [1.16, 1.19] | +14.7% | [+13.7%, +15.6%] | 1.1 pts | 2.6 | yes | yes |
| unique-str | uuid | 262144 | prefix | ordered | baseline | 8 | 440 | 510 | 1.14× [1.13, 1.16] | +12.5% | [+11.5%, +13.5%] | 1.2 pts | 2.7 | yes | yes |
| unique-str | uuid | 262144 | churn | ordered | baseline | 8 | 434 | 454 | 1.05× [1.03, 1.07] | +4.7% | [+3.3%, +6.1%] | 1.7 pts | 1.6 | yes | yes |
| unique-str | uuid | 1048576 | valuesFor | ordered | baseline | 8 | 340 | 364 | 1.08× [1.07, 1.09] | +7.4% | [+6.7%, +8.1%] | 0.8 pts | 1.8 | yes | yes |
| unique-str | uuid | 1048576 | valuesBetween | ordered | baseline | 8 | 6407 | 7148 | 1.12× [1.11, 1.12] | +10.6% | [+10.2%, +10.9%] | 0.4 pts | 1.9 | yes | yes |
| unique-str | uuid | 1048576 | prefix | ordered | baseline | 8 | 1247 | 1397 | 1.12× [1.12, 1.12] | +10.7% | [+10.4%, +11.0%] | 0.4 pts | 1.6 | yes | yes |
| unique-str | uuid | 1048576 | churn | ordered | baseline | 8 | 600 | 633 | 1.05× [1.03, 1.06] | +4.7% | [+3.3%, +6.1%] | 1.6 pts | 2.0 | yes | yes |

Regime: parallel: 8 processes at a time, each with GOMAXPROCS 3, sharing the caches and the memory bandwidth.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str path n=262144 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesFor: ordered vs baseline: 6 processes resolved A as faster and 16 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=1048576 valuesFor: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=1048576 valuesBetween: ordered vs baseline: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str url n=262144 churn: ordered vs baseline: the pooled interval [-1.81%, 1.17%] includes zero
- multi-str url n=262144 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.43% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 valuesFor: ordered vs baseline: the processes scatter 6.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 valuesBetween: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 prefix: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.06% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=1048576 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=1048576 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.10% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str u64 n=1048576 churn: ordered vs baseline: the pooled interval [-0.63%, 3.28%] includes zero
- unique-str u64 n=1048576 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=1048576 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=1048576 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.08% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str uuid n=1048576 churn: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
