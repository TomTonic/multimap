| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | btree-sets | 12 | 52.9 | 135 | 2.58× [2.55, 2.60] | +61.2% | [+60.8%, +61.6%] | 0.5 pts | 1.6 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | hashed | 12 | 52.7 | 22.9 | 0.44× [0.43, 0.44] | -128.3% | [-129.9%, -126.8%] | 3.1 pts | 1.9 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | map-sets | 12 | 52.2 | 64.6 | 1.27× [1.21, 1.32] | +21.0% | [+17.6%, +24.3%] | 3.4 pts | 6.6 | no | yes |
| multi-str | street | 4096 | valuesBetween | ordered | btree-sets | 12 | 2289 | 4636 | 2.03× [2.00, 2.07] | +50.8% | [+50.1%, +51.6%] | 0.9 pts | 1.9 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | hashed | 12 | 2372 | 53.9 µs | 22.78× [22.44, 23.12] | +95.6% | [+95.5%, +95.7%] | 0.1 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | map-sets | 12 | 2373 | 54.0 µs | 23.00× [22.52, 23.50] | +95.7% | [+95.6%, +95.7%] | 0.1 pts | 1.4 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | btree-sets | 12 | 225 | 529 | 2.35× [2.32, 2.38] | +57.4% | [+56.9%, +57.9%] | 0.7 pts | 1.4 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | hashed | 12 | 257 | 49.4 µs | 195.89× [190.07, 202.07] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | map-sets | 12 | 253 | 46.3 µs | 185.05× [179.58, 190.86] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | churn | ordered | btree-sets | 12 | 111 | 200 | 1.81× [1.79, 1.84] | +44.9% | [+44.1%, +45.6%] | 1.1 pts | 1.4 | yes | yes |
| multi-str | street | 4096 | churn | ordered | hashed | 12 | 108 | 56.6 | 0.52× [0.51, 0.52] | -93.0% | [-94.9%, -91.0%] | 2.6 pts | 1.4 | yes | yes |
| multi-str | street | 4096 | churn | ordered | map-sets | 12 | 109 | 69.4 | 0.65× [0.64, 0.66] | -53.7% | [-55.2%, -52.1%] | 4.6 pts | 3.1 | yes | yes |
| multi-str | street | 4096 | build | ordered | btree-sets | 12 | 4.38 ms | 8.00 ms | 1.83× [1.80, 1.86] | +45.4% | [+44.6%, +46.3%] | 1.1 pts | 1.5 | yes | yes |
| multi-str | street | 4096 | build | ordered | hashed | 12 | 4.38 ms | 2.74 ms | 0.63× [0.61, 0.65] | -58.9% | [-62.8%, -55.0%] | 4.8 pts | 1.9 | yes | yes |
| multi-str | street | 4096 | build | ordered | map-sets | 12 | 4.39 ms | 3.40 ms | 0.77× [0.74, 0.80] | -29.8% | [-34.3%, -25.3%] | 6.1 pts | 2.2 | no | yes |
| multi-str | street | 16384 | valuesFor | ordered | btree-sets | 12 | 74.7 | 183 | 2.46× [2.42, 2.50] | +59.3% | [+58.7%, +60.0%] | 0.9 pts | 2.4 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | hashed | 12 | 73.1 | 28.2 | 0.39× [0.38, 0.40] | -156.8% | [-160.7%, -152.9%] | 4.6 pts | 2.4 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | map-sets | 12 | 72.7 | 76.7 | 1.07× [1.04, 1.10] | +6.4% | [+3.5%, +9.3%] | 2.9 pts | 5.2 | no | yes |
| multi-str | street | 16384 | valuesBetween | ordered | btree-sets | 12 | 2889 | 5384 | 1.85× [1.84, 1.86] | +46.0% | [+45.8%, +46.3%] | 0.6 pts | 1.9 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | hashed | 12 | 3227 | 229.9 µs | 69.97× [68.56, 71.43] | +98.6% | [+98.5%, +98.6%] | 0.0 pts | 0.5 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | map-sets | 12 | 3379 | 223.3 µs | 64.47× [62.89, 66.14] | +98.4% | [+98.4%, +98.5%] | 0.1 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | btree-sets | 12 | 826 | 2042 | 2.47× [2.43, 2.50] | +59.4% | [+58.9%, +60.0%] | 0.8 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | hashed | 12 | 1157 | 227.3 µs | 195.38× [191.75, 199.14] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 0.4 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | map-sets | 12 | 1153 | 216.9 µs | 184.14× [179.97, 188.50] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 0.6 | yes | yes |
| multi-str | street | 16384 | churn | ordered | btree-sets | 12 | 151 | 274 | 1.82× [1.80, 1.84] | +45.1% | [+44.5%, +45.6%] | 0.8 pts | 1.2 | yes | yes |
| multi-str | street | 16384 | churn | ordered | hashed | 12 | 147 | 73.6 | 0.50× [0.50, 0.51] | -98.7% | [-101.0%, -96.4%] | 3.8 pts | 2.2 | yes | yes |
| multi-str | street | 16384 | churn | ordered | map-sets | 12 | 148 | 99.4 | 0.66× [0.65, 0.68] | -51.3% | [-55.0%, -47.6%] | 7.0 pts | 3.4 | yes | yes |
| multi-str | street | 16384 | build | ordered | btree-sets | 12 | 24.04 ms | 44.01 ms | 1.83× [1.81, 1.84] | +45.2% | [+44.8%, +45.6%] | 0.6 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | build | ordered | hashed | 12 | 23.48 ms | 13.59 ms | 0.58× [0.57, 0.59] | -73.1% | [-76.0%, -70.2%] | 3.2 pts | 1.9 | yes | yes |
| multi-str | street | 16384 | build | ordered | map-sets | 12 | 24.02 ms | 17.90 ms | 0.73× [0.70, 0.77] | -36.3% | [-42.7%, -29.9%] | 6.3 pts | 2.3 | no | yes |
| multi-str | street | 212449 | valuesFor | ordered | btree-sets | 12 | 264 | 480 | 1.82× [1.79, 1.86] | +45.2% | [+44.2%, +46.2%] | 1.5 pts | 2.0 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | hashed | 12 | 205 | 103 | 0.50× [0.49, 0.51] | -100.7% | [-104.0%, -97.4%] | 5.8 pts | 3.0 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | map-sets | 12 | 230 | 273 | 1.19× [1.16, 1.22] | +15.9% | [+13.7%, +18.2%] | 2.4 pts | 3.2 | no | yes |
| multi-str | street | 212449 | valuesBetween | ordered | btree-sets | 12 | 7679 | 17.2 µs | 2.25× [2.22, 2.27] | +55.5% | [+54.9%, +56.0%] | 1.1 pts | 2.8 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | btree-sets | 12 | 24.2 µs | 55.5 µs | 2.32× [2.25, 2.39] | +56.9% | [+55.5%, +58.2%] | 1.9 pts | 1.2 | yes | yes |
| multi-str | street | 212449 | churn | ordered | btree-sets | 12 | 560 | 729 | 1.33× [1.29, 1.36] | +24.6% | [+22.7%, +26.4%] | 1.8 pts | 1.9 | yes | yes |
| multi-str | street | 212449 | churn | ordered | hashed | 12 | 456 | 275 | 0.58× [0.57, 0.59] | -72.1% | [-74.3%, -69.9%] | 6.5 pts | 1.0 | yes | yes |
| multi-str | street | 212449 | churn | ordered | map-sets | 12 | 443 | 313 | 0.73× [0.72, 0.74] | -37.4% | [-39.0%, -35.9%] | 1.8 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesFor: ordered vs hashed: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 valuesFor: ordered vs map-sets: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 prefix: ordered vs hashed: process 2 was suspended for 1s during its comparison
- multi-str street n=4096 churn: ordered vs hashed: the A/A validations found a systematic difference of +0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 churn: ordered vs map-sets: the A/A validations found a systematic difference of +0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 churn: ordered vs map-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 build: ordered vs hashed: the A/A validations found a systematic difference of -0.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 build: ordered vs map-sets: the A/A validations found a systematic difference of -1.15% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 build: ordered vs map-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesFor: ordered vs hashed: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesFor: ordered vs map-sets: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs hashed: the A/A validations found a systematic difference of +1.68% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 valuesBetween: ordered vs map-sets: the A/A validations found a systematic difference of +2.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +3.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +3.97% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 churn: ordered vs hashed: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 churn: ordered vs map-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 build: ordered vs map-sets: the A/A validations found a systematic difference of +0.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 build: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs hashed: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs map-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
