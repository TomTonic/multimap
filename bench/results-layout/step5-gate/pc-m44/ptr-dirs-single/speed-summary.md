| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | dirs | 4096 | valuesFor | ordered | baseline | 8 | 95.3 | 78.2 | 0.82× [0.81, 0.84] | -21.3% | [-23.7%, -18.9%] | 2.4 pts | 2.3 | no | yes |
| single-value-ptr | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 96.2 | 108 | 1.13× [1.11, 1.14] | +11.2% | [+9.9%, +12.6%] | 1.5 pts | 1.5 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1436 | 2246 | 1.57× [1.55, 1.59] | +36.3% | [+35.5%, +37.0%] | 0.8 pts | 0.8 | yes | yes |
| single-value-ptr | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1411 | 553 | 0.39× [0.38, 0.39] | -157.7% | [-160.8%, -154.6%] | 3.2 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | baseline | 8 | 1660 | 2605 | 1.53× [1.52, 1.55] | +34.8% | [+34.1%, +35.6%] | 0.8 pts | 0.6 | yes | yes |
| single-value-ptr | dirs | 4096 | prefix | ordered | btree-map | 8 | 1679 | 651 | 0.39× [0.38, 0.39] | -159.6% | [-162.2%, -157.0%] | 2.5 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | baseline | 8 | 354 | 205 | 0.58× [0.57, 0.59] | -72.7% | [-75.0%, -70.4%] | 2.8 pts | 1.4 | yes | yes |
| single-value-ptr | dirs | 4096 | churn | ordered | btree-map | 8 | 351 | 179 | 0.51× [0.50, 0.51] | -97.5% | [-99.6%, -95.5%] | 2.4 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 4096 | build | ordered | baseline | 8 | 4.78 ms | 2.87 ms | 0.60× [0.59, 0.60] | -66.8% | [-68.3%, -65.3%] | 2.3 pts | 1.5 | yes | yes |
| single-value-ptr | dirs | 4096 | build | ordered | btree-map | 8 | 4.83 ms | 2.63 ms | 0.54× [0.53, 0.55] | -84.0% | [-87.2%, -80.9%] | 3.1 pts | 2.2 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | baseline | 6 | 119 | 108 | 0.91× [0.90, 0.92] | -9.7% | [-10.5%, -8.9%] | 0.8 pts | 0.9 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesFor | ordered | btree-map | 6 | 121 | 159 | 1.33× [1.32, 1.34] | +24.6% | [+24.0%, +25.2%] | 0.5 pts | 0.6 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 1631 | 2797 | 1.72× [1.71, 1.73] | +41.8% | [+41.6%, +42.1%] | 0.3 pts | 0.4 | yes | yes |
| single-value-ptr | dirs | 16384 | valuesBetween | ordered | btree-map | 6 | 1612 | 757 | 0.47× [0.46, 0.47] | -114.9% | [-117.2%, -112.7%] | 2.2 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | baseline | 6 | 5999 | 11.3 µs | 1.86× [1.77, 1.95] | +46.2% | [+43.5%, +48.8%] | 2.5 pts | 0.8 | yes | yes |
| single-value-ptr | dirs | 16384 | prefix | ordered | btree-map | 6 | 5963 | 2625 | 0.45× [0.43, 0.46] | -123.6% | [-130.3%, -116.9%] | 6.4 pts | 0.8 | yes | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | baseline | 6 | 398 | 262 | 0.66× [0.65, 0.68] | -50.8% | [-53.5%, -48.1%] | 2.6 pts | 1.7 | yes | yes |
| single-value-ptr | dirs | 16384 | churn | ordered | btree-map | 6 | 406 | 247 | 0.61× [0.59, 0.62] | -64.4% | [-68.2%, -60.7%] | 3.6 pts | 1.5 | yes | yes |
| single-value-ptr | dirs | 16384 | build | ordered | baseline | 6 | 22.96 ms | 14.32 ms | 0.62× [0.61, 0.63] | -60.4% | [-62.6%, -58.2%] | 2.1 pts | 0.8 | yes | yes |
| single-value-ptr | dirs | 16384 | build | ordered | btree-map | 6 | 22.97 ms | 13.82 ms | 0.60× [0.58, 0.61] | -67.5% | [-71.0%, -64.0%] | 3.3 pts | 1.0 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | baseline | 8 | 161 | 156 | 1.00× [0.94, 1.06] | -0.1% | [-5.9%, +5.6%] | 5.6 pts | 5.9 | no | no |
| single-value-ptr | dirs | 86215 | valuesFor | ordered | btree-map | 8 | 168 | 233 | 1.40× [1.38, 1.41] | +28.5% | [+27.7%, +29.3%] | 1.3 pts | 1.5 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 1826 | 3137 | 1.71× [1.70, 1.72] | +41.4% | [+41.1%, +41.8%] | 0.4 pts | 1.1 | yes | yes |
| single-value-ptr | dirs | 86215 | valuesBetween | ordered | btree-map | 8 | 1829 | 996 | 0.54× [0.54, 0.54] | -84.0% | [-84.6%, -83.5%] | 1.5 pts | 0.7 | yes | yes |
| single-value-ptr | dirs | 86215 | prefix | ordered | baseline | 8 | 37.5 µs | 78.6 µs | 2.12× [2.10, 2.13] | +52.8% | [+52.4%, +53.2%] | 0.4 pts | 1.2 | yes | yes |
| single-value-ptr | dirs | 86215 | prefix | ordered | btree-map | 8 | 36.5 µs | 15.5 µs | 0.42× [0.42, 0.43] | -137.0% | [-140.4%, -133.6%] | 3.3 pts | 1.4 | yes | yes |
| single-value-ptr | dirs | 86215 | churn | ordered | baseline | 8 | 542 | 443 | 0.81× [0.80, 0.83] | -23.3% | [-25.7%, -20.9%] | 2.5 pts | 1.3 | no | yes |
| single-value-ptr | dirs | 86215 | churn | ordered | btree-map | 8 | 590 | 450 | 0.76× [0.75, 0.78] | -31.5% | [-34.0%, -29.0%] | 3.4 pts | 2.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of +0.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=4096 build: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of -1.28% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: the pooled difference of -0.14% does not clear the 1.17% noise floor, the bound on what the harness reports between identical code in every process
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-5.86%, 5.58%] includes zero
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 5.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr dirs n=86215 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of +2.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-ptr dirs n=86215 churn: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
