| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | btree-sets | 6 | 85.1 | 160 | 1.89× [1.85, 1.93] | +47.0% | [+46.0%, +48.1%] | 1.0 pts | 1.7 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | hashed | 6 | 83.5 | 31.6 | 0.38× [0.37, 0.38] | -165.4% | [-168.9%, -162.0%] | 3.3 pts | 1.8 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | map-sets | 6 | 83.0 | 72.1 | 0.87× [0.85, 0.88] | -15.2% | [-17.0%, -13.4%] | 1.7 pts | 2.2 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 6 | 2728 | 5485 | 2.02× [2.01, 2.03] | +50.4% | [+50.2%, +50.7%] | 0.2 pts | 0.4 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | hashed | 6 | 2836 | 56.2 µs | 19.84× [19.57, 20.13] | +95.0% | [+94.9%, +95.0%] | 0.1 pts | 1.2 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | map-sets | 6 | 2818 | 56.5 µs | 20.05× [19.58, 20.54] | +95.0% | [+94.9%, +95.1%] | 0.1 pts | 1.5 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | btree-sets | 6 | 2835 | 7527 | 2.65× [2.62, 2.68] | +62.3% | [+61.9%, +62.7%] | 0.4 pts | 0.3 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | hashed | 6 | 3015 | 55.9 µs | 18.71× [18.13, 19.31] | +94.7% | [+94.5%, +94.8%] | 0.2 pts | 1.3 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | map-sets | 6 | 2957 | 58.7 µs | 19.86× [19.59, 20.14] | +95.0% | [+94.9%, +95.0%] | 0.1 pts | 0.5 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | btree-sets | 6 | 166 | 238 | 1.42× [1.41, 1.44] | +29.8% | [+29.2%, +30.3%] | 0.5 pts | 0.6 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | hashed | 6 | 164 | 72.4 | 0.44× [0.43, 0.45] | -127.1% | [-130.2%, -123.9%] | 3.0 pts | 1.6 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | map-sets | 6 | 165 | 81.4 | 0.50× [0.48, 0.51] | -101.2% | [-106.9%, -95.5%] | 5.4 pts | 3.0 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | btree-sets | 6 | 7.87 ms | 11.76 ms | 1.48× [1.47, 1.49] | +32.4% | [+31.8%, +33.1%] | 0.6 pts | 0.7 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | hashed | 6 | 7.92 ms | 3.76 ms | 0.47× [0.46, 0.48] | -112.4% | [-116.5%, -108.3%] | 3.9 pts | 1.4 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | map-sets | 6 | 7.91 ms | 4.59 ms | 0.58× [0.56, 0.59] | -73.9% | [-78.1%, -69.6%] | 4.0 pts | 0.9 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | btree-sets | 12 | 117 | 221 | 1.87× [1.85, 1.90] | +46.6% | [+45.8%, +47.4%] | 1.0 pts | 1.7 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | hashed | 12 | 114 | 37.6 | 0.33× [0.32, 0.33] | -207.0% | [-210.8%, -203.2%] | 6.8 pts | 3.4 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | map-sets | 12 | 114 | 85.0 | 0.75× [0.73, 0.77] | -33.5% | [-36.8%, -30.2%] | 4.5 pts | 6.5 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 12 | 3542 | 6215 | 1.75× [1.75, 1.76] | +43.0% | [+42.8%, +43.1%] | 0.5 pts | 1.3 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | hashed | 12 | 4249 | 243.2 µs | 57.81× [57.28, 58.36] | +98.3% | [+98.3%, +98.3%] | 0.1 pts | 0.6 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | map-sets | 12 | 4181 | 233.9 µs | 55.41× [54.88, 55.95] | +98.2% | [+98.2%, +98.2%] | 0.0 pts | 0.3 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | btree-sets | 12 | 16.1 µs | 35.3 µs | 2.10× [2.07, 2.12] | +52.3% | [+51.7%, +52.9%] | 0.8 pts | 1.1 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | hashed | 12 | 17.5 µs | 255.6 µs | 14.88× [14.20, 15.63] | +93.3% | [+93.0%, +93.6%] | 0.4 pts | 0.9 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | map-sets | 12 | 17.8 µs | 265.0 µs | 14.77× [14.44, 15.11] | +93.2% | [+93.1%, +93.4%] | 0.2 pts | 0.4 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | btree-sets | 12 | 233 | 351 | 1.49× [1.48, 1.50] | +32.9% | [+32.6%, +33.2%] | 0.3 pts | 0.5 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | hashed | 12 | 222 | 100.0 | 0.45× [0.44, 0.45] | -123.4% | [-127.0%, -119.8%] | 4.6 pts | 2.2 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | map-sets | 12 | 226 | 127 | 0.56× [0.55, 0.57] | -78.0% | [-81.2%, -74.8%] | 8.4 pts | 2.5 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | btree-sets | 12 | 44.89 ms | 65.95 ms | 1.48× [1.47, 1.48] | +32.2% | [+31.8%, +32.6%] | 0.8 pts | 1.4 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | hashed | 12 | 43.87 ms | 20.70 ms | 0.47× [0.46, 0.48] | -112.1% | [-115.1%, -109.2%] | 3.8 pts | 1.6 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | map-sets | 12 | 44.24 ms | 26.52 ms | 0.59× [0.58, 0.61] | -69.1% | [-73.0%, -65.2%] | 8.2 pts | 2.7 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | btree-sets | 12 | 238 | 429 | 1.83× [1.79, 1.87] | +45.3% | [+44.0%, +46.5%] | 1.5 pts | 1.8 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | hashed | 12 | 198 | 90.0 | 0.46× [0.45, 0.47] | -118.8% | [-123.1%, -114.4%] | 6.4 pts | 2.5 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | map-sets | 12 | 220 | 221 | 1.00× [0.98, 1.03] | +0.4% | [-1.9%, +2.8%] | 2.8 pts | 2.3 | no | no |
| multi-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 12 | 7081 | 14.5 µs | 2.06× [2.00, 2.11] | +51.4% | [+50.1%, +52.7%] | 1.8 pts | 1.8 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | btree-sets | 12 | 122.6 µs | 265.4 µs | 2.26× [2.16, 2.36] | +55.7% | [+53.7%, +57.7%] | 2.5 pts | 0.6 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | btree-sets | 12 | 541 | 666 | 1.24× [1.23, 1.25] | +19.4% | [+18.7%, +20.1%] | 1.0 pts | 1.3 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | hashed | 12 | 454 | 250 | 0.53× [0.52, 0.55] | -87.3% | [-93.8%, -80.7%] | 7.9 pts | 0.8 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | map-sets | 12 | 455 | 300 | 0.65× [0.64, 0.67] | -52.8% | [-56.2%, -49.5%] | 3.8 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str dirs n=4096 valuesFor: ordered vs map-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=4096 churn: ordered vs map-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=4096 build: ordered vs map-sets: the A/A validations found a systematic difference of -1.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesFor: ordered vs hashed: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesFor: ordered vs map-sets: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesBetween: ordered vs hashed: the A/A validations found a systematic difference of +1.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesBetween: ordered vs map-sets: the A/A validations found a systematic difference of +2.66% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +3.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +3.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 churn: ordered vs hashed: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 churn: ordered vs map-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 build: ordered vs hashed: the A/A validations found a systematic difference of +0.37% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 build: ordered vs map-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs hashed: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs map-sets: the pooled difference of 0.42% does not clear the 0.50% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=86215 valuesFor: ordered vs map-sets: the pooled interval [-1.93%, 2.77%] includes zero
- multi-str dirs n=86215 valuesFor: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs map-sets: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str dirs n=86215 churn: ordered vs hashed: the A/A validations found a systematic difference of +1.82% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
