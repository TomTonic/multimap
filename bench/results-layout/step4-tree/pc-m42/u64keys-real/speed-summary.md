| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 38.4 | 38.6 | 1.01× [0.99, 1.03] | +0.6% | [-1.5%, +2.7%] | 2.1 pts | 2.0 | no | no |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 38.5 | 157 | 4.10× [4.05, 4.15] | +75.6% | [+75.3%, +75.9%] | 0.3 pts | 1.9 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2729 | 2783 | 1.02× [0.99, 1.05] | +1.7% | [-0.9%, +4.4%] | 2.5 pts | 0.8 | no | no |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 2719 | 7036 | 2.61× [2.58, 2.63] | +61.6% | [+61.3%, +62.0%] | 0.4 pts | 0.6 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 77.0 | 58.0 | 0.75× [0.74, 0.77] | -32.5% | [-34.4%, -30.7%] | 1.9 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 8 | 76.4 | 157 | 2.04× [2.02, 2.06] | +51.0% | [+50.5%, +51.5%] | 0.5 pts | 1.0 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 7.90 ms | 5.88 ms | 0.74× [0.74, 0.75] | -34.2% | [-34.9%, -33.6%] | 0.7 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 8 | 7.89 ms | 14.77 ms | 1.88× [1.86, 1.89] | +46.7% | [+46.3%, +47.2%] | 0.5 pts | 1.4 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 6 | 47.6 | 47.5 | 1.00× [0.99, 1.01] | -0.4% | [-1.5%, +0.7%] | 1.1 pts | 2.0 | yes | no |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 47.8 | 192 | 4.03× [4.00, 4.06] | +75.2% | [+75.0%, +75.4%] | 0.2 pts | 0.9 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3568 | 3639 | 1.02× [1.01, 1.03] | +1.8% | [+1.0%, +2.6%] | 0.7 pts | 1.1 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 3574 | 7482 | 2.09× [2.07, 2.11] | +52.2% | [+51.7%, +52.6%] | 0.5 pts | 1.7 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 6 | 135 | 76.8 | 0.57× [0.57, 0.57] | -75.5% | [-76.6%, -74.5%] | 1.0 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 6 | 143 | 232 | 1.64× [1.62, 1.66] | +39.0% | [+38.4%, +39.7%] | 0.6 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 6 | 47.73 ms | 28.90 ms | 0.60× [0.60, 0.61] | -65.6% | [-66.1%, -65.0%] | 0.5 pts | 0.6 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 6 | 48.05 ms | 77.64 ms | 1.62× [1.60, 1.63] | +38.2% | [+37.6%, +38.8%] | 0.5 pts | 1.3 | yes | yes |
| natural | u64 | 262144 | valuesFor | ordered | baseline | 8 | 146 | 130 | 0.88× [0.84, 0.92] | -13.8% | [-18.7%, -8.8%] | 4.8 pts | 5.8 | no | yes |
| natural | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 186 | 530 | 2.86× [2.82, 2.90] | +65.0% | [+64.6%, +65.5%] | 0.7 pts | 2.2 | yes | yes |
| natural | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 6485 | 6557 | 1.01× [0.98, 1.04] | +0.9% | [-2.1%, +3.9%] | 2.8 pts | 4.8 | no | no |
| natural | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 7016 | 21.9 µs | 3.15× [3.13, 3.17] | +68.2% | [+68.0%, +68.5%] | 0.4 pts | 1.6 | yes | yes |
| natural | u64 | 262144 | churn | ordered | baseline | 8 | 367 | 330 | 0.85× [0.84, 0.87] | -17.0% | [-19.4%, -14.7%] | 4.0 pts | 0.6 | no | yes |
| natural | u64 | 262144 | churn | ordered | btree-sets | 8 | 450 | 716 | 1.63× [1.57, 1.69] | +38.7% | [+36.5%, +40.9%] | 2.1 pts | 2.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.48%, 2.74%] includes zero
- natural u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.74% does not clear the 2.24% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.91%, 4.40%] includes zero
- natural u64 n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.27% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.49%, 0.72%] includes zero
- natural u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 5.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.15%, 3.90%] includes zero
- natural u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=262144 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=262144 churn: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
