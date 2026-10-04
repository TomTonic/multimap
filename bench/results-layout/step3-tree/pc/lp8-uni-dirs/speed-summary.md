| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | dirs | 4096 | valuesFor | ordered | btree-map | 12 | 78.4 | 112 | 1.44× [1.43, 1.45] | +30.6% | [+30.1%, +31.1%] | 0.9 pts | 0.9 | yes | yes |
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage | 12 | 78.5 | 123 | 1.58× [1.56, 1.61] | +36.9% | [+35.7%, +38.0%] | 1.4 pts | 2.5 | yes | yes |
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-zc | 12 | 78.1 | 104 | 1.33× [1.32, 1.34] | +24.8% | [+24.3%, +25.3%] | 0.6 pts | 1.1 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | btree-map | 12 | 2047 | 1571 | 0.75× [0.75, 0.76] | -32.5% | [-33.1%, -32.0%] | 1.5 pts | 0.6 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage | 12 | 2048 | 1863 | 0.91× [0.89, 0.93] | -9.9% | [-11.8%, -7.9%] | 2.3 pts | 1.2 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-zc | 12 | 2076 | 1226 | 0.59× [0.58, 0.60] | -69.1% | [-71.0%, -67.2%] | 2.3 pts | 1.2 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | btree-map | 12 | 2240 | 1966 | 0.85× [0.84, 0.87] | -17.2% | [-18.8%, -15.5%] | 3.2 pts | 1.5 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage | 12 | 2310 | 2260 | 0.97× [0.97, 0.98] | -2.6% | [-3.1%, -2.1%] | 1.8 pts | 1.1 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage-zc | 12 | 2241 | 1392 | 0.63× [0.61, 0.64] | -59.9% | [-63.8%, -56.0%] | 4.2 pts | 1.2 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | btree-map | 12 | 212 | 195 | 0.92× [0.90, 0.94] | -8.7% | [-10.6%, -6.8%] | 2.4 pts | 1.0 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage | 12 | 210 | 264 | 1.27× [1.25, 1.28] | +21.0% | [+20.1%, +21.9%] | 1.4 pts | 1.0 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage-zc | 12 | 217 | 340 | 1.58× [1.57, 1.59] | +36.8% | [+36.4%, +37.2%] | 1.0 pts | 0.7 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | btree-map | 12 | 2.93 ms | 2.87 ms | 0.98× [0.97, 0.99] | -2.0% | [-3.0%, -0.9%] | 1.2 pts | 1.1 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage | 12 | 2.89 ms | 4.97 ms | 1.71× [1.71, 1.72] | +41.7% | [+41.5%, +41.9%] | 0.5 pts | 1.0 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage-zc | 12 | 2.93 ms | 5.93 ms | 2.01× [1.99, 2.02] | +50.2% | [+49.9%, +50.5%] | 0.4 pts | 0.7 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | btree-map | 12 | 112 | 165 | 1.48× [1.45, 1.50] | +32.2% | [+31.1%, +33.4%] | 1.3 pts | 1.1 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage | 12 | 112 | 153 | 1.37× [1.33, 1.41] | +26.9% | [+25.0%, +28.9%] | 2.3 pts | 2.7 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-zc | 12 | 109 | 130 | 1.18× [1.16, 1.19] | +14.9% | [+13.9%, +15.9%] | 1.2 pts | 1.3 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | btree-map | 12 | 2600 | 2066 | 0.79× [0.78, 0.80] | -26.0% | [-27.6%, -24.5%] | 1.8 pts | 1.1 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage | 12 | 2588 | 2234 | 0.87× [0.86, 0.88] | -15.3% | [-16.4%, -14.2%] | 1.7 pts | 1.4 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-zc | 12 | 2579 | 1394 | 0.54× [0.54, 0.55] | -84.8% | [-86.1%, -83.4%] | 2.5 pts | 1.3 | yes | yes |
| unique-str | dirs | 16384 | prefix | ordered | btree-map | 12 | 10.6 µs | 9232 | 0.87× [0.84, 0.89] | -15.3% | [-18.5%, -12.1%] | 4.2 pts | 0.9 | no | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage | 12 | 10.9 µs | 9245 | 0.85× [0.82, 0.88] | -17.7% | [-21.4%, -14.0%] | 4.6 pts | 0.9 | no | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage-zc | 12 | 10.2 µs | 5022 | 0.49× [0.48, 0.50] | -105.5% | [-109.9%, -101.2%] | 6.3 pts | 0.7 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | btree-map | 12 | 329 | 291 | 0.90× [0.89, 0.93] | -10.5% | [-13.0%, -8.0%] | 2.6 pts | 1.8 | no | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage | 12 | 316 | 350 | 1.12× [1.10, 1.13] | +10.4% | [+9.4%, +11.5%] | 1.3 pts | 1.5 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage-zc | 12 | 329 | 443 | 1.36× [1.34, 1.37] | +26.2% | [+25.6%, +26.8%] | 1.0 pts | 1.0 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | btree-map | 12 | 15.77 ms | 15.78 ms | 0.99× [0.98, 1.01] | -0.6% | [-2.4%, +1.2%] | 2.3 pts | 1.5 | yes | no |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage | 12 | 15.60 ms | 23.64 ms | 1.51× [1.50, 1.53] | +33.9% | [+33.2%, +34.5%] | 0.7 pts | 0.9 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage-zc | 12 | 15.51 ms | 28.88 ms | 1.86× [1.84, 1.87] | +46.1% | [+45.7%, +46.6%] | 0.7 pts | 1.1 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | btree-map | 12 | 304 | 309 | 1.04× [1.02, 1.05] | +3.6% | [+2.2%, +4.9%] | 2.6 pts | 1.1 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage | 12 | 276 | 249 | 0.93× [0.92, 0.94] | -7.5% | [-8.7%, -6.4%] | 3.6 pts | 1.6 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-zc | 12 | 252 | 204 | 0.81× [0.80, 0.83] | -22.8% | [-25.0%, -20.7%] | 3.7 pts | 1.4 | yes | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | btree-map | 12 | 5061 | 4694 | 0.93× [0.91, 0.96] | -7.5% | [-10.5%, -4.5%] | 2.8 pts | 1.1 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage | 12 | 4084 | 3031 | 0.73× [0.70, 0.75] | -37.7% | [-42.8%, -32.6%] | 5.6 pts | 2.3 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-zc | 12 | 3712 | 1882 | 0.50× [0.49, 0.51] | -101.4% | [-105.6%, -97.2%] | 9.4 pts | 2.0 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | btree-map | 12 | 90.1 µs | 68.6 µs | 0.79× [0.73, 0.87] | -25.9% | [-37.1%, -14.8%] | 14.0 pts | 0.8 | no | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage | 12 | 92.7 µs | 60.3 µs | 0.70× [0.68, 0.72] | -42.7% | [-46.7%, -38.6%] | 7.0 pts | 0.5 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage-zc | 12 | 84.3 µs | 32.3 µs | 0.39× [0.38, 0.40] | -156.3% | [-159.8%, -152.7%] | 8.5 pts | 0.7 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | btree-map | 12 | 627 | 625 | 1.00× [0.99, 1.01] | -0.2% | [-1.1%, +0.7%] | 1.6 pts | 1.1 | yes | no |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage | 12 | 579 | 583 | 1.01× [0.99, 1.03] | +1.2% | [-1.0%, +3.4%] | 2.2 pts | 1.6 | no | no |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage-zc | 12 | 596 | 679 | 1.15× [1.13, 1.16] | +12.8% | [+11.9%, +13.7%] | 1.0 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str dirs n=4096 valuesFor: ordered vs ordered-lpage: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=4096 build: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.70% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=16384 build: ordered vs btree-map: the pooled difference of -0.60% does not clear the 0.70% noise floor, the bound on what the harness reports between identical code in every process
- unique-str dirs n=16384 build: ordered vs btree-map: the pooled interval [-2.43%, 1.24%] includes zero
- unique-str dirs n=16384 build: ordered vs btree-map: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the A/A validations found a systematic difference of -0.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesBetween: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.87% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 prefix: ordered vs ordered-lpage: the A/A validations found a systematic difference of -0.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 churn: ordered vs btree-map: the pooled difference of -0.21% does not clear the 0.46% noise floor, the bound on what the harness reports between identical code in every process
- unique-str dirs n=86215 churn: ordered vs btree-map: the pooled interval [-1.13%, 0.71%] includes zero
- unique-str dirs n=86215 churn: ordered vs btree-map: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str dirs n=86215 churn: ordered vs ordered-lpage: the pooled interval [-1.03%, 3.37%] includes zero
- unique-str dirs n=86215 churn: ordered vs ordered-lpage: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
