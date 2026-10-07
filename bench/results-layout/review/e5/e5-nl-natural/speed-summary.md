| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 6 | 97.6 | 105 | 1.07× [1.06, 1.09] | +6.9% | [+5.6%, +8.1%] | 1.2 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 2580 | 2686 | 1.04× [1.02, 1.05] | +3.8% | [+2.4%, +5.2%] | 1.3 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 6 | 201 | 151 | 0.76× [0.74, 0.77] | -31.7% | [-34.3%, -29.1%] | 2.5 pts | 1.7 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 6 | 9.89 ms | 7.05 ms | 0.71× [0.70, 0.72] | -41.1% | [-43.1%, -39.0%] | 2.0 pts | 1.4 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 165 | 193 | 1.20× [1.10, 1.31] | +16.5% | [+9.2%, +23.8%] | 7.1 pts | 6.0 | no | yes |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 3410 | 4422 | 1.37× [1.20, 1.58] | +26.8% | [+16.7%, +36.9%] | 10.1 pts | 13.8 | no | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 359 | 342 | 0.97× [0.96, 0.98] | -3.1% | [-4.3%, -1.9%] | 1.9 pts | 1.5 | yes | yes |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 292.19 ms | 257.12 ms | 0.87× [0.86, 0.89] | -14.3% | [-16.8%, -11.9%] | 2.4 pts | 2.7 | no | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 65.2 | 67.9 | 1.03× [1.03, 1.04] | +3.4% | [+2.8%, +3.9%] | 0.5 pts | 0.7 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 1873 | 2281 | 1.24× [1.21, 1.28] | +19.3% | [+17.1%, +21.6%] | 2.5 pts | 1.5 | no | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 150 | 99.3 | 0.66× [0.66, 0.67] | -50.4% | [-51.4%, -49.5%] | 2.0 pts | 1.3 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 6.02 ms | 3.81 ms | 0.64× [0.63, 0.65] | -57.0% | [-60.0%, -53.9%] | 3.4 pts | 2.3 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 103 | 115 | 1.11× [1.11, 1.12] | +10.2% | [+9.6%, +10.7%] | 1.2 pts | 1.1 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 2483 | 3271 | 1.33× [1.32, 1.33] | +24.6% | [+24.2%, +25.0%] | 0.6 pts | 1.0 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 244 | 223 | 0.90× [0.87, 0.92] | -11.4% | [-14.3%, -8.5%] | 3.0 pts | 2.6 | no | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 166.19 ms | 137.17 ms | 0.83× [0.81, 0.84] | -20.6% | [-22.9%, -18.4%] | 2.2 pts | 1.8 | no | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 45.0 | 59.5 | 1.32× [1.30, 1.34] | +24.1% | [+23.0%, +25.2%] | 1.0 pts | 1.4 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2825 | 2495 | 0.88× [0.86, 0.91] | -13.4% | [-16.5%, -10.3%] | 2.9 pts | 0.8 | no | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 73.7 | 51.0 | 0.69× [0.69, 0.70] | -44.3% | [-45.5%, -43.1%] | 1.1 pts | 1.1 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 9.39 ms | 5.28 ms | 0.56× [0.56, 0.57] | -77.6% | [-78.5%, -76.8%] | 1.0 pts | 0.9 | yes | yes |
| natural | u64 | 65536 | valuesFor | ordered | baseline | 8 | 65.0 | 72.9 | 1.12× [1.09, 1.14] | +10.4% | [+8.2%, +12.5%] | 6.1 pts | 8.2 | no | yes |
| natural | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 4269 | 4156 | 0.97× [0.96, 0.99] | -2.9% | [-4.5%, -1.4%] | 3.3 pts | 5.1 | yes | yes |
| natural | u64 | 65536 | churn | ordered | baseline | 8 | 195 | 157 | 0.79× [0.77, 0.81] | -26.7% | [-29.4%, -24.0%] | 2.7 pts | 1.0 | no | yes |
| natural | u64 | 65536 | build | ordered | baseline | 8 | 269.24 ms | 216.22 ms | 0.80× [0.78, 0.83] | -24.9% | [-28.9%, -20.9%] | 3.8 pts | 3.1 | no | yes |
| natural | url | 4096 | valuesFor | ordered | baseline | 6 | 100 | 109 | 1.08× [1.07, 1.09] | +7.3% | [+6.1%, +8.4%] | 1.1 pts | 1.7 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | baseline | 6 | 3804 | 3796 | 1.00× [0.98, 1.02] | -0.3% | [-2.1%, +1.5%] | 1.7 pts | 0.9 | yes | no |
| natural | url | 4096 | churn | ordered | baseline | 6 | 182 | 133 | 0.73× [0.72, 0.73] | -37.4% | [-38.4%, -36.4%] | 1.0 pts | 0.8 | yes | yes |
| natural | url | 4096 | build | ordered | baseline | 6 | 18.14 ms | 12.24 ms | 0.68× [0.67, 0.69] | -46.8% | [-48.6%, -45.0%] | 1.7 pts | 1.1 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | baseline | 8 | 180 | 207 | 1.15× [1.13, 1.18] | +13.4% | [+11.8%, +14.9%] | 1.6 pts | 1.6 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | baseline | 8 | 5555 | 6239 | 1.13× [1.07, 1.19] | +11.2% | [+6.8%, +15.7%] | 4.9 pts | 5.1 | no | yes |
| natural | url | 65536 | churn | ordered | baseline | 8 | 414 | 391 | 0.94× [0.93, 0.95] | -6.1% | [-7.5%, -4.8%] | 1.4 pts | 1.0 | yes | yes |
| natural | url | 65536 | build | ordered | baseline | 8 | 546.80 ms | 458.59 ms | 0.84× [0.83, 0.85] | -19.1% | [-20.1%, -18.0%] | 1.1 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=65536 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 6.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesBetween: ordered vs baseline: the processes scatter 13.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 build: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.99% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 8.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs baseline: 7 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=65536 valuesBetween: ordered vs baseline: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 build: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.37% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.25% does not clear the 2.33% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.06%, 1.55%] includes zero
- natural url n=65536 valuesBetween: ordered vs baseline: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
