| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 72.6 | 57.0 | 0.78× [0.76, 0.80] | -28.2% | [-31.0%, -25.4%] | 3.1 pts | 3.0 | yes | yes |
| natural-ptr | street | 4096 | valuesFor | ordered | btree-sets | 8 | 73.5 | 138 | 1.88× [1.85, 1.90] | +46.7% | [+45.9%, +47.5%] | 0.8 pts | 1.9 | yes | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 1965 | 2534 | 1.30× [1.27, 1.33] | +23.0% | [+21.2%, +24.7%] | 1.7 pts | 1.0 | yes | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 1943 | 4730 | 2.44× [2.40, 2.48] | +59.0% | [+58.3%, +59.7%] | 0.8 pts | 1.7 | yes | yes |
| natural-ptr | street | 4096 | prefix | ordered | baseline | 8 | 265 | 255 | 0.96× [0.95, 0.97] | -4.1% | [-5.2%, -3.0%] | 1.3 pts | 1.1 | yes | yes |
| natural-ptr | street | 4096 | prefix | ordered | btree-sets | 8 | 268 | 540 | 2.02× [2.00, 2.03] | +50.4% | [+50.0%, +50.9%] | 0.6 pts | 1.3 | yes | yes |
| natural-ptr | street | 4096 | churn | ordered | baseline | 8 | 251 | 110 | 0.43× [0.42, 0.44] | -132.3% | [-137.0%, -127.6%] | 4.8 pts | 1.4 | yes | yes |
| natural-ptr | street | 4096 | churn | ordered | btree-sets | 8 | 252 | 199 | 0.79× [0.78, 0.80] | -26.4% | [-28.4%, -24.3%] | 1.9 pts | 0.9 | yes | yes |
| natural-ptr | street | 4096 | build | ordered | baseline | 8 | 9.24 ms | 4.24 ms | 0.46× [0.45, 0.46] | -119.1% | [-120.8%, -117.3%] | 1.9 pts | 1.0 | yes | yes |
| natural-ptr | street | 4096 | build | ordered | btree-sets | 8 | 9.17 ms | 7.48 ms | 0.82× [0.81, 0.84] | -21.7% | [-23.7%, -19.7%] | 2.2 pts | 1.7 | yes | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | baseline | 6 | 90.3 | 78.2 | 0.86× [0.85, 0.88] | -15.8% | [-17.8%, -13.8%] | 1.9 pts | 2.1 | yes | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | btree-sets | 6 | 91.0 | 188 | 2.03× [1.96, 2.10] | +50.6% | [+49.0%, +52.3%] | 1.6 pts | 3.7 | yes | yes |
| natural-ptr | street | 16384 | valuesBetween | ordered | baseline | 6 | 2346 | 3145 | 1.34× [1.33, 1.35] | +25.4% | [+25.0%, +25.9%] | 0.4 pts | 0.6 | yes | yes |
| natural-ptr | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 2353 | 5574 | 2.36× [2.34, 2.39] | +57.7% | [+57.2%, +58.1%] | 0.5 pts | 2.0 | yes | yes |
| natural-ptr | street | 16384 | prefix | ordered | baseline | 6 | 748 | 940 | 1.25× [1.24, 1.27] | +20.1% | [+19.2%, +21.1%] | 0.9 pts | 0.7 | yes | yes |
| natural-ptr | street | 16384 | prefix | ordered | btree-sets | 6 | 752 | 2122 | 2.80× [2.76, 2.85] | +64.3% | [+63.7%, +64.9%] | 0.6 pts | 1.1 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | baseline | 6 | 295 | 157 | 0.52× [0.52, 0.53] | -90.8% | [-93.8%, -87.9%] | 2.8 pts | 1.2 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | btree-sets | 6 | 302 | 286 | 0.95× [0.94, 0.96] | -5.2% | [-6.0%, -4.4%] | 0.8 pts | 0.6 | yes | yes |
| natural-ptr | street | 16384 | build | ordered | baseline | 6 | 45.17 ms | 22.60 ms | 0.50× [0.50, 0.51] | -98.7% | [-101.0%, -96.4%] | 2.2 pts | 0.7 | yes | yes |
| natural-ptr | street | 16384 | build | ordered | btree-sets | 6 | 45.36 ms | 40.81 ms | 0.89× [0.88, 0.91] | -11.8% | [-13.8%, -9.8%] | 1.9 pts | 0.8 | yes | yes |
| natural-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 174 | 195 | 1.15× [1.09, 1.21] | +13.0% | [+8.3%, +17.6%] | 5.7 pts | 5.8 | no | yes |
| natural-ptr | street | 212449 | valuesFor | ordered | btree-sets | 8 | 209 | 484 | 2.32× [2.26, 2.38] | +56.9% | [+55.7%, +58.0%] | 1.6 pts | 2.6 | yes | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 3517 | 5366 | 1.55× [1.49, 1.61] | +35.4% | [+32.7%, +38.0%] | 2.7 pts | 2.7 | yes | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 4071 | 17.0 µs | 4.23× [4.17, 4.29] | +76.4% | [+76.0%, +76.7%] | 0.4 pts | 1.0 | yes | yes |
| natural-ptr | street | 212449 | prefix | ordered | baseline | 8 | 10.7 µs | 17.4 µs | 1.62× [1.56, 1.68] | +38.3% | [+36.0%, +40.6%] | 2.4 pts | 3.9 | yes | yes |
| natural-ptr | street | 212449 | prefix | ordered | btree-sets | 8 | 12.8 µs | 53.1 µs | 4.16× [4.06, 4.26] | +75.9% | [+75.4%, +76.5%] | 0.6 pts | 1.3 | yes | yes |
| natural-ptr | street | 212449 | churn | ordered | baseline | 8 | 604 | 515 | 0.83× [0.81, 0.84] | -20.8% | [-23.3%, -18.4%] | 2.7 pts | 0.7 | no | yes |
| natural-ptr | street | 212449 | churn | ordered | btree-sets | 8 | 673 | 732 | 1.10× [1.08, 1.11] | +8.8% | [+7.6%, +9.9%] | 1.1 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr street n=4096 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of +1.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 5.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 prefix: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
