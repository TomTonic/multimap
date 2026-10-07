| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 45.6 | 38.2 | 0.83× [0.82, 0.85] | -20.1% | [-22.2%, -18.0%] | 2.1 pts | 1.8 | no | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 45.7 | 160 | 3.52× [3.48, 3.56] | +71.6% | [+71.3%, +71.9%] | 0.3 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2900 | 2850 | 0.99× [0.97, 1.01] | -1.2% | [-3.0%, +0.6%] | 2.3 pts | 0.8 | yes | no |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 2868 | 7213 | 2.54× [2.51, 2.56] | +60.6% | [+60.2%, +61.0%] | 0.5 pts | 0.8 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 76.5 | 60.2 | 0.79× [0.78, 0.80] | -26.8% | [-28.1%, -25.4%] | 1.4 pts | 1.2 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 8 | 77.2 | 163 | 2.10× [2.07, 2.13] | +52.4% | [+51.7%, +53.0%] | 0.6 pts | 1.4 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 9.73 ms | 6.25 ms | 0.65× [0.63, 0.68] | -53.9% | [-59.9%, -47.9%] | 5.6 pts | 2.7 | no | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 8 | 9.65 ms | 16.07 ms | 1.64× [1.62, 1.66] | +38.9% | [+38.2%, +39.7%] | 1.5 pts | 2.0 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 8 | 48.7 | 47.8 | 0.98× [0.97, 0.98] | -2.4% | [-2.9%, -2.0%] | 0.4 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 8 | 48.3 | 196 | 4.04× [4.01, 4.08] | +75.3% | [+75.0%, +75.5%] | 0.2 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 3712 | 3676 | 1.00× [0.99, 1.00] | -0.3% | [-0.9%, +0.3%] | 0.5 pts | 0.8 | yes | no |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 8 | 3695 | 7701 | 2.10× [2.07, 2.13] | +52.4% | [+51.6%, +53.1%] | 0.8 pts | 3.6 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 8 | 101 | 87.4 | 0.87× [0.85, 0.89] | -15.0% | [-17.0%, -13.0%] | 2.0 pts | 1.7 | no | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 8 | 106 | 249 | 2.28× [2.14, 2.44] | +56.1% | [+53.2%, +59.0%] | 2.8 pts | 5.3 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 8 | 38.83 ms | 33.72 ms | 0.86× [0.84, 0.88] | -16.6% | [-19.1%, -14.1%] | 2.5 pts | 1.1 | no | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 8 | 40.25 ms | 87.84 ms | 2.13× [1.97, 2.32] | +53.0% | [+49.1%, +57.0%] | 3.7 pts | 5.4 | yes | yes |
| natural | u64 | 262144 | valuesFor | ordered | baseline | 8 | 120 | 121 | 1.02× [0.98, 1.06] | +1.7% | [-1.9%, +5.2%] | 3.6 pts | 4.4 | no | no |
| natural | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 169 | 552 | 3.22× [3.16, 3.28] | +68.9% | [+68.3%, +69.5%] | 1.1 pts | 2.8 | yes | yes |
| natural | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 5555 | 6454 | 1.16× [1.15, 1.18] | +14.1% | [+12.7%, +15.6%] | 1.5 pts | 2.8 | yes | yes |
| natural | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 6035 | 21.9 µs | 3.65× [3.60, 3.70] | +72.6% | [+72.2%, +73.0%] | 0.4 pts | 2.1 | yes | yes |
| natural | u64 | 262144 | churn | ordered | baseline | 8 | 361 | 345 | 0.93× [0.90, 0.95] | -8.1% | [-11.0%, -5.1%] | 7.5 pts | 1.3 | no | yes |
| natural | u64 | 262144 | churn | ordered | btree-sets | 8 | 448 | 698 | 1.58× [1.52, 1.66] | +36.9% | [+34.1%, +39.6%] | 2.6 pts | 5.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural u64 n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.99%, 0.56%] includes zero
- natural u64 n=4096 build: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 build: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.90%, 0.25%] includes zero
- natural u64 n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs btree-sets: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 build: ordered vs btree-sets: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.94%, 5.25%] includes zero
- natural u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 churn: ordered vs btree-sets: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
