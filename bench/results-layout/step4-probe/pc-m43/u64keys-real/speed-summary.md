| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 38.7 | 38.3 | 0.99× [0.98, 1.00] | -1.0% | [-2.4%, +0.4%] | 1.3 pts | 1.3 | yes | no |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 38.6 | 157 | 4.14× [3.94, 4.35] | +75.8% | [+74.6%, +77.0%] | 1.2 pts | 7.1 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2731 | 2838 | 1.05× [1.02, 1.08] | +4.9% | [+2.3%, +7.6%] | 2.6 pts | 0.9 | no | yes |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 2734 | 7049 | 2.60× [2.56, 2.65] | +61.6% | [+60.9%, +62.2%] | 0.6 pts | 1.1 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 63.5 | 57.5 | 0.91× [0.90, 0.92] | -9.7% | [-10.7%, -8.6%] | 1.5 pts | 1.9 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 8 | 63.7 | 155 | 2.44× [2.43, 2.45] | +59.0% | [+58.8%, +59.2%] | 0.2 pts | 0.4 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 6.42 ms | 5.85 ms | 0.91× [0.91, 0.92] | -9.3% | [-9.9%, -8.7%] | 0.6 pts | 0.8 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 8 | 6.49 ms | 14.77 ms | 2.28× [2.26, 2.30] | +56.1% | [+55.7%, +56.5%] | 0.4 pts | 1.0 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 6 | 47.3 | 46.8 | 0.99× [0.98, 0.99] | -1.4% | [-2.4%, -0.5%] | 0.9 pts | 1.9 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 48.0 | 192 | 4.01× [3.99, 4.04] | +75.1% | [+74.9%, +75.2%] | 0.1 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3514 | 3663 | 1.04× [1.03, 1.04] | +3.8% | [+3.4%, +4.2%] | 0.4 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 3528 | 7504 | 2.13× [2.12, 2.15] | +53.2% | [+52.8%, +53.5%] | 0.3 pts | 1.1 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 6 | 85.8 | 75.9 | 0.89× [0.89, 0.89] | -12.4% | [-13.0%, -11.9%] | 0.5 pts | 0.6 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 6 | 89.5 | 228 | 2.54× [2.52, 2.56] | +60.7% | [+60.3%, +61.0%] | 0.3 pts | 0.9 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 6 | 31.00 ms | 28.54 ms | 0.92× [0.91, 0.93] | -9.1% | [-10.2%, -8.0%] | 1.1 pts | 1.3 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 6 | 32.03 ms | 77.10 ms | 2.43× [2.39, 2.48] | +58.9% | [+58.1%, +59.7%] | 0.7 pts | 1.2 | yes | yes |
| natural | u64 | 262144 | valuesFor | ordered | baseline | 8 | 156 | 139 | 0.90× [0.86, 0.95] | -11.0% | [-16.5%, -5.6%] | 5.5 pts | 7.4 | no | yes |
| natural | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 183 | 526 | 2.87× [2.82, 2.92] | +65.1% | [+64.6%, +65.7%] | 0.9 pts | 2.4 | yes | yes |
| natural | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 6532 | 6652 | 1.01× [0.99, 1.04] | +1.4% | [-1.0%, +3.9%] | 2.3 pts | 3.4 | no | no |
| natural | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 6987 | 22.1 µs | 3.18× [3.12, 3.24] | +68.5% | [+67.9%, +69.1%] | 0.6 pts | 2.5 | yes | yes |
| natural | u64 | 262144 | churn | ordered | baseline | 8 | 337 | 316 | 0.91× [0.89, 0.92] | -10.0% | [-11.8%, -8.2%] | 3.0 pts | 0.5 | yes | yes |
| natural | u64 | 262144 | churn | ordered | btree-sets | 8 | 429 | 709 | 1.67× [1.61, 1.74] | +40.2% | [+37.9%, +42.5%] | 2.5 pts | 2.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.42%, 0.42%] includes zero
- natural u64 n=4096 valuesFor: ordered vs btree-sets: the processes scatter 7.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 7.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.01%, 3.86%] includes zero
- natural u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of +2.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=262144 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
