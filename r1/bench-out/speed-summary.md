| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | street | 4096 | valuesFor | ordered | btree-sets | 8 | 76.3 | 180 | 2.35× [2.33, 2.38] | +57.5% | [+57.1%, +57.9%] | 0.5 pts | 1.9 | yes | yes |
| multi | street | 4096 | valuesFor | ordered | hashed | 8 | 75.5 | 34.1 | 0.45× [0.45, 0.45] | -121.0% | [-122.1%, -120.0%] | 1.3 pts | 1.0 | yes | yes |
| multi | street | 4096 | valuesFor | ordered | map-sets | 8 | 75.8 | 79.4 | 1.04× [1.04, 1.05] | +4.2% | [+3.8%, +4.6%] | 0.5 pts | 1.0 | yes | yes |
| multi | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 3393 | 5362 | 1.59× [1.57, 1.60] | +36.9% | [+36.5%, +37.4%] | 0.5 pts | 0.9 | yes | yes |
| multi | street | 4096 | valuesBetween | ordered | hashed | 8 | 3414 | 72.5 µs | 21.00× [20.62, 21.39] | +95.2% | [+95.2%, +95.3%] | 0.1 pts | 2.9 | yes | yes |
| multi | street | 4096 | valuesBetween | ordered | map-sets | 8 | 3437 | 70.9 µs | 20.47× [20.11, 20.84] | +95.1% | [+95.0%, +95.2%] | 0.1 pts | 3.4 | yes | yes |
| multi | street | 4096 | prefix | ordered | btree-sets | 8 | 329 | 653 | 2.00× [1.99, 2.01] | +50.0% | [+49.8%, +50.3%] | 0.3 pts | 0.5 | yes | yes |
| multi | street | 4096 | prefix | ordered | hashed | 8 | 343 | 71.1 µs | 206.47× [202.05, 211.09] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 1.3 | yes | yes |
| multi | street | 4096 | prefix | ordered | map-sets | 8 | 342 | 65.8 µs | 191.15× [185.28, 197.41] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 2.4 | yes | yes |
| multi | street | 4096 | churn | ordered | btree-sets | 8 | 126 | 237 | 1.89× [1.88, 1.91] | +47.2% | [+46.8%, +47.6%] | 0.5 pts | 1.0 | yes | yes |
| multi | street | 4096 | churn | ordered | hashed | 8 | 123 | 64.4 | 0.52× [0.52, 0.53] | -90.9% | [-91.9%, -89.9%] | 1.2 pts | 1.2 | yes | yes |
| multi | street | 4096 | churn | ordered | map-sets | 8 | 123 | 70.9 | 0.58× [0.57, 0.60] | -71.3% | [-74.9%, -67.8%] | 4.3 pts | 5.7 | yes | yes |
| multi | street | 4096 | build | ordered | btree-sets | 8 | 4.80 ms | 9.48 ms | 1.97× [1.96, 1.98] | +49.3% | [+49.0%, +49.5%] | 0.3 pts | 1.6 | yes | yes |
| multi | street | 4096 | build | ordered | hashed | 8 | 4.84 ms | 3.00 ms | 0.63× [0.62, 0.63] | -59.9% | [-62.1%, -57.7%] | 2.6 pts | 3.5 | yes | yes |
| multi | street | 4096 | build | ordered | map-sets | 8 | 4.80 ms | 3.48 ms | 0.73× [0.72, 0.74] | -37.7% | [-39.5%, -35.9%] | 2.1 pts | 2.8 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | btree-sets | 24 | 100 | 231 | 2.30× [2.29, 2.31] | +56.5% | [+56.3%, +56.8%] | 0.5 pts | 1.4 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | hashed | 24 | 97.4 | 36.1 | 0.37× [0.37, 0.37] | -169.8% | [-171.1%, -168.6%] | 1.8 pts | 1.2 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | map-sets | 24 | 99.3 | 87.6 | 0.89× [0.86, 0.91] | -12.7% | [-15.9%, -9.4%] | 4.1 pts | 4.4 | no | yes |
| multi | street | 16384 | valuesBetween | ordered | btree-sets | 24 | 3707 | 5734 | 1.54× [1.53, 1.55] | +35.2% | [+34.8%, +35.5%] | 0.6 pts | 1.3 | yes | yes |
| multi | street | 16384 | valuesBetween | ordered | hashed | 24 | 3726 | 284.3 µs | 75.30× [71.93, 79.01] | +98.7% | [+98.6%, +98.7%] | 0.1 pts | 7.4 | yes | yes |
| multi | street | 16384 | valuesBetween | ordered | map-sets | 24 | 3773 | 265.0 µs | 69.74× [68.92, 70.57] | +98.6% | [+98.5%, +98.6%] | 0.0 pts | 3.0 | yes | yes |
| multi | street | 16384 | prefix | ordered | btree-sets | 24 | 1220 | 2260 | 1.86× [1.85, 1.86] | +46.1% | [+45.9%, +46.3%] | 0.4 pts | 0.6 | yes | yes |
| multi | street | 16384 | prefix | ordered | hashed | 24 | 1295 | 289.3 µs | 221.07× [216.33, 226.02] | +99.5% | [+99.5%, +99.6%] | 0.0 pts | 1.8 | yes | yes |
| multi | street | 16384 | prefix | ordered | map-sets | 24 | 1300 | 268.3 µs | 205.86× [204.14, 207.62] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 0.9 | yes | yes |
| multi | street | 16384 | churn | ordered | btree-sets | 24 | 168 | 301 | 1.81× [1.80, 1.82] | +44.8% | [+44.4%, +45.1%] | 0.7 pts | 0.6 | yes | yes |
| multi | street | 16384 | churn | ordered | hashed | 24 | 158 | 71.8 | 0.46× [0.45, 0.46] | -118.5% | [-121.6%, -115.4%] | 5.0 pts | 3.0 | yes | yes |
| multi | street | 16384 | churn | ordered | map-sets | 24 | 164 | 89.7 | 0.55× [0.55, 0.56] | -80.6% | [-83.5%, -77.7%] | 5.2 pts | 1.5 | yes | yes |
| multi | street | 16384 | build | ordered | btree-sets | 24 | 25.45 ms | 46.72 ms | 1.83× [1.82, 1.85] | +45.5% | [+45.1%, +45.9%] | 0.6 pts | 7.0 | yes | yes |
| multi | street | 16384 | build | ordered | hashed | 24 | 25.47 ms | 12.68 ms | 0.50× [0.50, 0.50] | -100.8% | [-101.8%, -99.9%] | 2.3 pts | 5.9 | yes | yes |
| multi | street | 16384 | build | ordered | map-sets | 24 | 25.44 ms | 14.90 ms | 0.58× [0.58, 0.59] | -71.2% | [-72.3%, -70.2%] | 2.3 pts | 5.8 | yes | yes |
| multi | street | 212449 | valuesFor | ordered | btree-sets | 8 | 295 | 649 | 2.22× [2.18, 2.26] | +54.9% | [+54.1%, +55.7%] | 1.0 pts | 0.9 | yes | yes |
| multi | street | 212449 | valuesFor | ordered | hashed | 8 | 251 | 100 | 0.41× [0.40, 0.41] | -146.9% | [-148.4%, -145.5%] | 1.7 pts | 0.5 | yes | yes |
| multi | street | 212449 | valuesFor | ordered | map-sets | 8 | 274 | 241 | 0.89× [0.87, 0.90] | -12.8% | [-14.4%, -11.1%] | 2.0 pts | 1.2 | yes | yes |
| multi | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 6386 | 12.7 µs | 2.00× [1.98, 2.02] | +50.0% | [+49.6%, +50.5%] | 0.5 pts | 0.8 | yes | yes |
| multi | street | 212449 | prefix | ordered | btree-sets | 8 | 24.1 µs | 48.0 µs | 2.13× [2.07, 2.19] | +53.0% | [+51.8%, +54.2%] | 1.5 pts | 0.8 | yes | yes |
| multi | street | 212449 | churn | ordered | btree-sets | 8 | 569 | 852 | 1.51× [1.48, 1.54] | +33.7% | [+32.3%, +35.1%] | 1.6 pts | 1.1 | yes | yes |
| multi | street | 212449 | churn | ordered | hashed | 8 | 497 | 217 | 0.43× [0.41, 0.44] | -133.5% | [-141.2%, -125.9%] | 9.2 pts | 0.3 | yes | yes |
| multi | street | 212449 | churn | ordered | map-sets | 8 | 513 | 285 | 0.53× [0.52, 0.54] | -88.7% | [-91.1%, -86.3%] | 2.9 pts | 0.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi street n=4096 valuesBetween: ordered vs hashed: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 valuesBetween: ordered vs map-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 prefix: ordered vs hashed: the A/A validations found a systematic difference of +0.69% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=4096 prefix: ordered vs map-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 churn: ordered vs map-sets: the processes scatter 5.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 build: ordered vs hashed: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 build: ordered vs map-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 valuesFor: ordered vs map-sets: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 valuesBetween: ordered vs hashed: the processes scatter 7.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 valuesBetween: ordered vs map-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +0.77% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=16384 churn: ordered vs hashed: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 build: ordered vs btree-sets: the processes scatter 7.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 build: ordered vs hashed: the processes scatter 5.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 build: ordered vs map-sets: the processes scatter 5.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=212449 valuesFor: ordered vs hashed: the A/A validations found a systematic difference of +0.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=212449 churn: ordered vs hashed: the A/A validations found a systematic difference of +5.59% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=212449 churn: ordered vs map-sets: the A/A validations found a systematic difference of +3.91% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
