| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 53.7 | 37.8 | 0.70× [0.69, 0.71] | -42.3% | [-44.1%, -40.5%] | 2.1 pts | 1.7 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 53.6 | 157 | 2.92× [2.89, 2.94] | +65.7% | [+65.4%, +66.0%] | 0.3 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 3413 | 2787 | 0.82× [0.80, 0.84] | -21.6% | [-24.4%, -18.8%] | 2.7 pts | 1.0 | no | yes |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 3437 | 7158 | 2.10× [2.08, 2.12] | +52.3% | [+52.0%, +52.7%] | 0.4 pts | 0.8 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 76.4 | 59.6 | 0.78× [0.77, 0.78] | -28.3% | [-29.0%, -27.5%] | 1.2 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 8 | 76.9 | 162 | 2.11× [2.09, 2.13] | +52.7% | [+52.2%, +53.1%] | 0.7 pts | 1.6 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 9.49 ms | 6.05 ms | 0.64× [0.63, 0.64] | -56.9% | [-58.2%, -55.6%] | 1.7 pts | 1.6 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 8 | 9.50 ms | 15.45 ms | 1.61× [1.59, 1.62] | +37.7% | [+37.1%, +38.3%] | 0.6 pts | 0.6 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 8 | 56.4 | 47.1 | 0.84× [0.83, 0.84] | -19.7% | [-20.5%, -19.0%] | 0.7 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 8 | 56.7 | 195 | 3.50× [3.26, 3.79] | +71.4% | [+69.3%, +73.6%] | 2.0 pts | 9.4 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 4292 | 3664 | 0.85× [0.85, 0.86] | -17.0% | [-18.1%, -16.0%] | 1.4 pts | 1.7 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 8 | 4307 | 7667 | 1.85× [1.66, 2.09] | +46.0% | [+39.9%, +52.1%] | 5.9 pts | 20.3 | no | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 8 | 97.8 | 82.7 | 0.86× [0.85, 0.86] | -16.9% | [-17.9%, -15.9%] | 1.8 pts | 2.4 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 8 | 105 | 247 | 2.32× [2.24, 2.41] | +56.9% | [+55.4%, +58.5%] | 1.9 pts | 2.5 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 8 | 37.92 ms | 31.57 ms | 0.84× [0.82, 0.85] | -19.7% | [-22.2%, -17.1%] | 2.4 pts | 1.1 | no | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 8 | 38.53 ms | 85.67 ms | 2.20× [2.18, 2.22] | +54.5% | [+54.0%, +55.0%] | 1.2 pts | 1.7 | yes | yes |
| natural | u64 | 262144 | valuesFor | ordered | baseline | 8 | 135 | 120 | 0.89× [0.87, 0.91] | -12.3% | [-14.6%, -10.0%] | 2.4 pts | 3.2 | no | yes |
| natural | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 186 | 529 | 2.87× [2.80, 2.93] | +65.1% | [+64.3%, +65.9%] | 0.7 pts | 2.5 | yes | yes |
| natural | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 5543 | 6457 | 1.16× [1.14, 1.19] | +14.0% | [+12.1%, +15.8%] | 1.8 pts | 2.6 | yes | yes |
| natural | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 6004 | 21.7 µs | 3.64× [3.58, 3.70] | +72.5% | [+72.1%, +73.0%] | 0.5 pts | 2.3 | yes | yes |
| natural | u64 | 262144 | churn | ordered | baseline | 8 | 346 | 325 | 0.92× [0.89, 0.95] | -8.5% | [-11.9%, -5.1%] | 4.5 pts | 1.1 | no | yes |
| natural | u64 | 262144 | churn | ordered | btree-sets | 8 | 441 | 704 | 1.60× [1.54, 1.66] | +37.5% | [+35.0%, +39.9%] | 2.5 pts | 6.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural u64 n=16384 valuesFor: ordered vs btree-sets: the processes scatter 9.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 20.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 churn: ordered vs btree-sets: the processes scatter 6.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
