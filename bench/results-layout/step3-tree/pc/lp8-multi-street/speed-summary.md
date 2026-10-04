| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 53.8 | 138 | 2.57× [2.52, 2.63] | +61.1% | [+60.3%, +62.0%] | 0.8 pts | 2.7 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage | 8 | 53.5 | 53.1 | 0.99× [0.99, 1.00] | -0.5% | [-1.2%, +0.1%] | 0.6 pts | 0.6 | yes | no |
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 8 | 54.0 | 109 | 2.03× [1.98, 2.08] | +50.7% | [+49.4%, +51.9%] | 1.2 pts | 3.0 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mvzc | 8 | 53.5 | 89.9 | 1.69× [1.66, 1.72] | +40.8% | [+39.9%, +41.7%] | 0.9 pts | 1.8 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 2324 | 6996 | 3.03× [2.97, 3.10] | +67.0% | [+66.4%, +67.7%] | 0.7 pts | 1.6 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage | 8 | 2331 | 2361 | 1.02× [1.00, 1.04] | +2.1% | [+0.3%, +4.0%] | 1.9 pts | 0.8 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 8 | 2307 | 2171 | 0.95× [0.94, 0.96] | -5.7% | [-6.8%, -4.5%] | 1.7 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mvzc | 8 | 2293 | 1540 | 0.67× [0.66, 0.68] | -49.7% | [-52.4%, -47.0%] | 3.0 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | btree-sets | 8 | 224 | 708 | 3.13× [3.06, 3.21] | +68.1% | [+67.3%, +68.8%] | 0.8 pts | 2.0 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage | 8 | 225 | 226 | 1.01× [0.99, 1.03] | +0.6% | [-1.3%, +2.6%] | 1.9 pts | 1.2 | yes | no |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 8 | 225 | 296 | 1.32× [1.30, 1.34] | +24.1% | [+23.0%, +25.2%] | 1.1 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mvzc | 8 | 225 | 224 | 0.99× [0.98, 1.00] | -1.2% | [-2.1%, -0.3%] | 1.4 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | churn | ordered | btree-sets | 8 | 113 | 203 | 1.78× [1.75, 1.81] | +43.9% | [+42.9%, +44.8%] | 1.1 pts | 1.5 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage | 8 | 112 | 112 | 1.00× [1.00, 1.01] | +0.5% | [-0.5%, +1.4%] | 1.2 pts | 1.2 | yes | no |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 8 | 117 | 233 | 2.02× [2.00, 2.05] | +50.6% | [+49.9%, +51.3%] | 0.7 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mvzc | 8 | 116 | 308 | 2.68× [2.66, 2.71] | +62.7% | [+62.4%, +63.0%] | 0.4 pts | 0.7 | yes | yes |
| multi-str | street | 4096 | build | ordered | btree-sets | 8 | 4.61 ms | 8.16 ms | 1.76× [1.74, 1.79] | +43.3% | [+42.5%, +44.1%] | 0.9 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage | 8 | 4.50 ms | 4.54 ms | 1.01× [1.00, 1.02] | +1.2% | [+0.4%, +2.0%] | 0.8 pts | 0.8 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 8 | 4.61 ms | 10.38 ms | 2.27× [2.25, 2.29] | +55.9% | [+55.5%, +56.3%] | 0.4 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mvzc | 8 | 4.63 ms | 13.17 ms | 2.86× [2.82, 2.90] | +65.0% | [+64.6%, +65.5%] | 0.5 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | btree-sets | 6 | 76.2 | 189 | 2.51× [2.50, 2.52] | +60.2% | [+60.0%, +60.4%] | 0.2 pts | 0.5 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage | 6 | 74.1 | 74.3 | 1.00× [0.99, 1.02] | +0.3% | [-1.0%, +1.5%] | 1.2 pts | 1.2 | yes | no |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 6 | 73.7 | 129 | 1.75× [1.73, 1.77] | +42.9% | [+42.2%, +43.6%] | 0.7 pts | 1.5 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mvzc | 6 | 74.3 | 108 | 1.46× [1.43, 1.48] | +31.3% | [+30.2%, +32.5%] | 1.1 pts | 1.8 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 2894 | 8122 | 2.79× [2.76, 2.83] | +64.2% | [+63.7%, +64.7%] | 0.5 pts | 1.4 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage | 6 | 2908 | 2931 | 1.01× [1.00, 1.01] | +0.8% | [+0.3%, +1.2%] | 0.4 pts | 0.5 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 6 | 2903 | 2614 | 0.89× [0.88, 0.90] | -11.8% | [-13.0%, -10.5%] | 1.2 pts | 1.7 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mvzc | 6 | 2929 | 1761 | 0.60× [0.60, 0.60] | -66.8% | [-67.3%, -66.3%] | 0.4 pts | 0.4 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | btree-sets | 6 | 862 | 2965 | 3.43× [3.35, 3.50] | +70.8% | [+70.2%, +71.5%] | 0.6 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage | 6 | 829 | 848 | 1.02× [1.00, 1.04] | +1.6% | [-0.1%, +3.4%] | 1.7 pts | 0.9 | yes | no |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 6 | 837 | 901 | 1.08× [1.07, 1.09] | +7.4% | [+6.2%, +8.6%] | 1.1 pts | 0.6 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mvzc | 6 | 816 | 586 | 0.72× [0.71, 0.73] | -39.8% | [-41.8%, -37.8%] | 1.9 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | churn | ordered | btree-sets | 6 | 156 | 281 | 1.79× [1.77, 1.82] | +44.2% | [+43.4%, +45.0%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage | 6 | 155 | 154 | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +0.9%] | 0.8 pts | 0.9 | yes | no |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 6 | 161 | 278 | 1.71× [1.65, 1.77] | +41.4% | [+39.4%, +43.5%] | 1.9 pts | 2.3 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mvzc | 6 | 173 | 361 | 2.11× [2.09, 2.14] | +52.7% | [+52.1%, +53.3%] | 0.6 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | build | ordered | btree-sets | 6 | 24.41 ms | 43.99 ms | 1.80× [1.78, 1.82] | +44.5% | [+43.9%, +45.0%] | 0.5 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage | 6 | 24.16 ms | 27.06 ms | 1.12× [1.11, 1.14] | +11.0% | [+9.9%, +12.2%] | 1.1 pts | 1.4 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 6 | 24.75 ms | 47.86 ms | 1.93× [1.92, 1.95] | +48.3% | [+47.9%, +48.6%] | 0.3 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mvzc | 6 | 24.11 ms | 60.48 ms | 2.52× [2.51, 2.54] | +60.4% | [+60.1%, +60.7%] | 0.3 pts | 0.4 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | btree-sets | 12 | 311 | 537 | 1.73× [1.70, 1.77] | +42.3% | [+41.3%, +43.4%] | 1.7 pts | 2.2 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage | 12 | 241 | 241 | 1.00× [0.99, 1.01] | +0.1% | [-1.0%, +1.3%] | 4.2 pts | 3.3 | yes | no |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 12 | 248 | 228 | 0.93× [0.92, 0.94] | -7.4% | [-8.9%, -5.9%] | 2.5 pts | 1.6 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mvzc | 12 | 219 | 186 | 0.85× [0.83, 0.86] | -18.3% | [-20.7%, -15.9%] | 4.0 pts | 2.7 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | btree-sets | 12 | 8197 | 24.2 µs | 3.00× [2.96, 3.04] | +66.7% | [+66.3%, +67.1%] | 0.5 pts | 1.1 | yes | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage | 12 | 7024 | 7058 | 1.00× [0.98, 1.01] | -0.4% | [-1.6%, +0.9%] | 3.6 pts | 3.2 | yes | no |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 12 | 6762 | 4267 | 0.64× [0.61, 0.67] | -57.3% | [-64.2%, -50.3%] | 9.4 pts | 4.9 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mvzc | 12 | 6386 | 2943 | 0.46× [0.45, 0.47] | -115.6% | [-120.4%, -110.8%] | 9.9 pts | 3.3 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | btree-sets | 12 | 26.3 µs | 81.1 µs | 3.16× [3.04, 3.29] | +68.3% | [+67.1%, +69.6%] | 1.5 pts | 1.0 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage | 12 | 22.0 µs | 22.3 µs | 1.00× [0.99, 1.01] | -0.0% | [-1.1%, +1.0%] | 4.3 pts | 2.3 | yes | no |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 12 | 20.2 µs | 12.8 µs | 0.65× [0.64, 0.66] | -54.0% | [-57.3%, -50.6%] | 7.1 pts | 3.1 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mvzc | 12 | 18.9 µs | 8651 | 0.45× [0.44, 0.46] | -121.2% | [-125.4%, -116.9%] | 8.3 pts | 1.0 | yes | yes |
| multi-str | street | 212449 | churn | ordered | btree-sets | 12 | 587 | 781 | 1.35× [1.32, 1.38] | +25.9% | [+24.0%, +27.8%] | 2.1 pts | 1.5 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage | 12 | 510 | 505 | 0.98× [0.96, 1.01] | -1.8% | [-4.5%, +0.9%] | 4.2 pts | 0.9 | no | no |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 12 | 517 | 592 | 1.13× [1.12, 1.15] | +11.9% | [+10.6%, +13.1%] | 3.3 pts | 0.9 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mvzc | 12 | 541 | 691 | 1.26× [1.22, 1.30] | +20.5% | [+18.1%, +22.9%] | 4.1 pts | 0.8 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 valuesFor: ordered vs ordered-lpage: the pooled interval [-1.19%, 0.12%] includes zero
- multi-str street n=4096 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage: the A/A validations found a systematic difference of -0.90% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of +0.66% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 prefix: ordered vs ordered-lpage: the pooled difference of 0.62% does not clear the 0.83% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 prefix: ordered vs ordered-lpage: the pooled interval [-1.34%, 2.58%] includes zero
- multi-str street n=4096 churn: ordered vs ordered-lpage: the pooled interval [-0.49%, 1.42%] includes zero
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage: the pooled difference of 0.25% does not clear the 0.40% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage: the pooled interval [-0.98%, 1.49%] includes zero
- multi-str street n=16384 prefix: ordered vs ordered-lpage: the pooled interval [-0.10%, 3.39%] includes zero
- multi-str street n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.94% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 churn: ordered vs ordered-lpage: the pooled difference of 0.07% does not clear the 0.73% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 churn: ordered vs ordered-lpage: the pooled interval [-0.72%, 0.87%] includes zero
- multi-str street n=16384 churn: ordered vs ordered-lpage-mv: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage: the pooled difference of 0.13% does not clear the 0.89% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage: the pooled interval [-1.04%, 1.31%] includes zero
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage-mvzc: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: the pooled difference of -0.39% does not clear the 0.47% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: the pooled interval [-1.65%, 0.88%] includes zero
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mvzc: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs ordered-lpage: the pooled difference of -0.01% does not clear the 2.00% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 prefix: ordered vs ordered-lpage: the pooled interval [-1.06%, 1.03%] includes zero
- multi-str street n=212449 prefix: ordered vs ordered-lpage: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs ordered-lpage: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 prefix: ordered vs ordered-lpage-mv: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 churn: ordered vs ordered-lpage: the pooled difference of -1.80% does not clear the 1.97% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 churn: ordered vs ordered-lpage: the pooled interval [-4.45%, 0.85%] includes zero
