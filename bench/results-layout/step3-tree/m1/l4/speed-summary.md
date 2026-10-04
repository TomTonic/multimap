| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | btree-sets | 10 | 130 | 212 | 1.62× [1.61, 1.64] | +38.4% | [+37.7%, +39.1%] | 0.9 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage | 10 | 129 | 128 | 0.99× [0.99, 1.00] | -0.6% | [-0.9%, -0.3%] | 0.4 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mv | 10 | 131 | 227 | 1.73× [1.71, 1.75] | +42.2% | [+41.7%, +42.7%] | 0.7 pts | 1.4 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mvzc | 10 | 130 | 203 | 1.56× [1.55, 1.57] | +35.9% | [+35.4%, +36.4%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 10 | 3941 | 12.3 µs | 3.11× [3.08, 3.13] | +67.8% | [+67.5%, +68.1%] | 0.4 pts | 1.5 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage | 10 | 3919 | 3920 | 1.00× [1.00, 1.01] | +0.4% | [-0.3%, +1.2%] | 0.9 pts | 0.5 | yes | no |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mv | 10 | 3958 | 4042 | 1.02× [1.01, 1.03] | +2.1% | [+1.3%, +2.9%] | 1.0 pts | 0.9 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mvzc | 10 | 3887 | 3012 | 0.78× [0.77, 0.78] | -28.3% | [-29.1%, -27.5%] | 1.0 pts | 0.5 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | btree-sets | 10 | 4412 | 17.1 µs | 3.73× [3.68, 3.78] | +73.2% | [+72.9%, +73.6%] | 0.7 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage | 10 | 4394 | 4491 | 1.01× [1.00, 1.02] | +1.2% | [+0.3%, +2.0%] | 1.0 pts | 0.5 | yes | no |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mv | 10 | 4494 | 5061 | 1.11× [1.09, 1.13] | +9.9% | [+8.4%, +11.4%] | 1.9 pts | 0.6 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mvzc | 10 | 4440 | 3557 | 0.80× [0.79, 0.82] | -24.5% | [-26.7%, -22.4%] | 3.1 pts | 0.5 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | btree-sets | 10 | 215 | 287 | 1.33× [1.32, 1.34] | +25.0% | [+24.3%, +25.6%] | 1.2 pts | 1.2 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage | 10 | 213 | 214 | 1.00× [1.00, 1.01] | +0.3% | [+0.0%, +0.5%] | 0.3 pts | 0.8 | yes | no |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mv | 10 | 217 | 395 | 1.81× [1.81, 1.82] | +44.9% | [+44.7%, +45.1%] | 0.2 pts | 0.5 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mvzc | 10 | 220 | 473 | 2.12× [2.10, 2.15] | +52.9% | [+52.3%, +53.5%] | 1.0 pts | 2.2 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | btree-sets | 10 | 10.50 ms | 14.21 ms | 1.36× [1.35, 1.36] | +26.4% | [+26.1%, +26.7%] | 0.4 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage | 10 | 10.54 ms | 10.79 ms | 1.02× [1.01, 1.03] | +2.3% | [+1.3%, +3.4%] | 1.3 pts | 1.4 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mv | 10 | 10.49 ms | 21.32 ms | 2.03× [2.02, 2.04] | +50.8% | [+50.5%, +51.1%] | 0.4 pts | 1.4 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mvzc | 10 | 10.49 ms | 24.58 ms | 2.36× [2.35, 2.37] | +57.6% | [+57.4%, +57.7%] | 0.6 pts | 1.8 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | btree-sets | 10 | 183 | 295 | 1.65× [1.60, 1.70] | +39.3% | [+37.5%, +41.1%] | 2.3 pts | 1.2 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage | 10 | 166 | 165 | 1.00× [0.99, 1.02] | +0.3% | [-1.1%, +1.8%] | 1.8 pts | 3.0 | yes | no |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mv | 10 | 170 | 276 | 1.62× [1.58, 1.65] | +38.1% | [+36.8%, +39.4%] | 1.7 pts | 2.7 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mvzc | 10 | 167 | 251 | 1.48× [1.43, 1.52] | +32.3% | [+30.3%, +34.2%] | 2.5 pts | 6.1 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 10 | 4799 | 15.3 µs | 3.16× [3.10, 3.23] | +68.4% | [+67.8%, +69.0%] | 0.7 pts | 1.1 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage | 10 | 4388 | 4366 | 0.99× [0.99, 1.00] | -0.7% | [-1.3%, -0.1%] | 0.7 pts | 0.7 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mv | 10 | 4545 | 4353 | 0.95× [0.93, 0.97] | -5.1% | [-7.0%, -3.2%] | 2.5 pts | 2.2 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mvzc | 10 | 4329 | 3236 | 0.74× [0.74, 0.75] | -34.3% | [-34.9%, -33.7%] | 0.9 pts | 0.8 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | btree-sets | 10 | 19.9 µs | 68.7 µs | 3.49× [3.43, 3.55] | +71.3% | [+70.8%, +71.8%] | 0.6 pts | 1.3 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage | 10 | 21.9 µs | 21.5 µs | 1.00× [0.99, 1.00] | -0.3% | [-1.0%, +0.4%] | 0.9 pts | 0.9 | yes | no |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mv | 10 | 22.1 µs | 19.6 µs | 0.90× [0.89, 0.91] | -10.9% | [-12.0%, -9.8%] | 1.3 pts | 1.0 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mvzc | 10 | 22.2 µs | 14.6 µs | 0.67× [0.66, 0.67] | -49.9% | [-50.7%, -49.1%] | 1.1 pts | 0.8 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | btree-sets | 10 | 364 | 462 | 1.27× [1.24, 1.29] | +21.0% | [+19.2%, +22.7%] | 2.4 pts | 2.1 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage | 10 | 298 | 297 | 0.99× [0.99, 1.00] | -0.5% | [-1.5%, +0.5%] | 1.2 pts | 1.0 | yes | no |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mv | 10 | 304 | 476 | 1.50× [1.43, 1.57] | +33.2% | [+30.2%, +36.3%] | 4.9 pts | 11.3 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mvzc | 10 | 347 | 586 | 1.66× [1.58, 1.75] | +39.7% | [+36.5%, +42.8%] | 3.8 pts | 7.6 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | btree-sets | 10 | 53.46 ms | 74.26 ms | 1.39× [1.38, 1.40] | +28.2% | [+27.7%, +28.8%] | 0.7 pts | 2.9 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage | 10 | 53.49 ms | 57.13 ms | 1.07× [1.06, 1.07] | +6.2% | [+5.6%, +6.9%] | 1.1 pts | 3.8 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mv | 10 | 53.33 ms | 99.47 ms | 1.84× [1.79, 1.89] | +45.6% | [+44.2%, +47.0%] | 1.7 pts | 6.9 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mvzc | 10 | 53.80 ms | 118.70 ms | 2.18× [2.15, 2.21] | +54.2% | [+53.6%, +54.8%] | 0.8 pts | 2.4 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | btree-sets | 22 | 393 | 626 | 1.61× [1.56, 1.65] | +37.7% | [+36.0%, +39.5%] | 3.1 pts | 1.5 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage | 22 | 329 | 334 | 1.02× [1.01, 1.02] | +1.6% | [+0.8%, +2.4%] | 2.0 pts | 0.7 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mv | 22 | 338 | 398 | 1.15× [1.12, 1.17] | +12.9% | [+10.9%, +14.9%] | 4.7 pts | 2.1 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mvzc | 22 | 369 | 388 | 1.09× [1.07, 1.11] | +7.9% | [+6.1%, +9.7%] | 3.4 pts | 1.7 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 22 | 7477 | 23.6 µs | 3.13× [3.09, 3.16] | +68.0% | [+67.7%, +68.4%] | 0.6 pts | 1.1 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage | 22 | 6706 | 6715 | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.2%] | 1.3 pts | 1.0 | yes | no |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mv | 22 | 6752 | 5609 | 0.82× [0.81, 0.83] | -22.3% | [-23.8%, -20.9%] | 3.2 pts | 1.8 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mvzc | 22 | 6716 | 4479 | 0.66× [0.65, 0.67] | -51.7% | [-53.7%, -49.6%] | 3.3 pts | 1.3 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | btree-sets | 22 | 161.2 µs | 619.4 µs | 3.92× [3.83, 4.02] | +74.5% | [+73.9%, +75.1%] | 1.5 pts | 0.6 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage | 22 | 150.9 µs | 153.8 µs | 1.01× [1.00, 1.02] | +1.3% | [+0.3%, +2.3%] | 2.0 pts | 0.9 | yes | no |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mv | 22 | 155.9 µs | 119.7 µs | 0.76× [0.75, 0.77] | -31.6% | [-33.5%, -29.6%] | 4.1 pts | 1.4 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mvzc | 22 | 155.2 µs | 86.5 µs | 0.55× [0.54, 0.57] | -81.0% | [-85.2%, -76.8%] | 6.9 pts | 2.1 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | btree-sets | 22 | 753 | 875 | 1.17× [1.16, 1.18] | +14.5% | [+13.6%, +15.5%] | 2.0 pts | 1.0 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage | 22 | 590 | 572 | 0.85× [0.84, 0.86] | -17.3% | [-18.9%, -15.7%] | 3.6 pts | 0.4 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mv | 22 | 631 | 699 | 1.06× [1.05, 1.07] | +5.4% | [+4.3%, +6.6%] | 3.5 pts | 0.5 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mvzc | 22 | 701 | 886 | 1.16× [1.14, 1.17] | +13.4% | [+12.1%, +14.8%] | 4.7 pts | 0.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str dirs n=4096 valuesBetween: ordered vs ordered-lpage: the pooled difference of 0.42% does not clear the 0.73% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=4096 valuesBetween: ordered vs ordered-lpage: the pooled interval [-0.31%, 1.15%] includes zero
- multi-str dirs n=4096 prefix: ordered vs ordered-lpage: the pooled difference of 1.17% does not clear the 5.46% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=4096 churn: ordered vs ordered-lpage: the pooled difference of 0.27% does not clear the 0.36% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=4096 churn: ordered vs ordered-lpage-mvzc: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the pooled difference of 0.35% does not clear the 0.41% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the pooled interval [-1.14%, 1.83%] includes zero
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage-mvzc: the processes scatter 6.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesBetween: ordered vs ordered-lpage: the A/A validations found a systematic difference of +0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesBetween: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of +0.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 prefix: ordered vs ordered-lpage: the pooled difference of -0.27% does not clear the 11.96% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=16384 prefix: ordered vs ordered-lpage: the A/A validations found a systematic difference of -7.76% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 prefix: ordered vs ordered-lpage: the pooled interval [-0.96%, 0.42%] includes zero
- multi-str dirs n=16384 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 churn: ordered vs ordered-lpage: the pooled difference of -0.52% does not clear the 0.80% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=16384 churn: ordered vs ordered-lpage: the pooled interval [-1.52%, 0.48%] includes zero
- multi-str dirs n=16384 churn: ordered vs ordered-lpage-mv: the processes scatter 11.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 churn: ordered vs ordered-lpage-mvzc: the processes scatter 7.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs ordered-lpage: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs ordered-lpage-mv: the processes scatter 6.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of -0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 build: ordered vs ordered-lpage-mvzc: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of +0.85% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of +0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the pooled interval [-1.14%, 0.25%] includes zero
- multi-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +1.19% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage: the pooled difference of 1.30% does not clear the 1.66% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage: the A/A validations found a systematic difference of -1.01% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of -1.13% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage-mvzc: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.77% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 churn: ordered vs ordered-lpage: the A/A validations found a systematic difference of +0.97% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
