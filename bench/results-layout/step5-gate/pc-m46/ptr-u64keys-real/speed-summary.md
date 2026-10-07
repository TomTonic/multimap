| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | u64 | 4096 | valuesFor | ordered | baseline | 6 | 56.2 | 40.4 | 0.71× [0.71, 0.72] | -40.1% | [-41.6%, -38.7%] | 1.4 pts | 1.2 | yes | yes |
| natural-ptr | u64 | 4096 | valuesFor | ordered | btree-sets | 6 | 56.1 | 161 | 2.89× [2.87, 2.90] | +65.3% | [+65.2%, +65.5%] | 0.1 pts | 0.5 | yes | yes |
| natural-ptr | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 3577 | 2873 | 0.81× [0.80, 0.81] | -24.0% | [-24.9%, -23.1%] | 0.9 pts | 0.3 | yes | yes |
| natural-ptr | u64 | 4096 | valuesBetween | ordered | btree-sets | 6 | 3570 | 7634 | 2.13× [2.12, 2.15] | +53.2% | [+52.7%, +53.6%] | 0.4 pts | 1.1 | yes | yes |
| natural-ptr | u64 | 4096 | churn | ordered | baseline | 6 | 79.8 | 61.3 | 0.77× [0.76, 0.77] | -30.5% | [-31.4%, -29.6%] | 0.9 pts | 0.7 | yes | yes |
| natural-ptr | u64 | 4096 | churn | ordered | btree-sets | 6 | 80.4 | 165 | 2.03× [2.00, 2.06] | +50.8% | [+50.1%, +51.5%] | 0.7 pts | 1.8 | yes | yes |
| natural-ptr | u64 | 4096 | build | ordered | baseline | 6 | 10.29 ms | 6.32 ms | 0.61× [0.61, 0.62] | -62.8% | [-63.9%, -61.8%] | 1.0 pts | 1.4 | yes | yes |
| natural-ptr | u64 | 4096 | build | ordered | btree-sets | 6 | 10.32 ms | 15.48 ms | 1.50× [1.48, 1.51] | +33.2% | [+32.5%, +33.9%] | 0.7 pts | 1.6 | yes | yes |
| natural-ptr | u64 | 16384 | valuesFor | ordered | baseline | 6 | 59.9 | 50.1 | 0.84× [0.83, 0.84] | -19.2% | [-20.0%, -18.4%] | 0.8 pts | 1.5 | yes | yes |
| natural-ptr | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 60.7 | 199 | 3.26× [3.20, 3.33] | +69.4% | [+68.7%, +70.0%] | 0.6 pts | 2.1 | yes | yes |
| natural-ptr | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 4477 | 3824 | 0.86× [0.85, 0.86] | -16.9% | [-18.0%, -15.7%] | 1.1 pts | 1.6 | yes | yes |
| natural-ptr | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 4474 | 8204 | 1.82× [1.81, 1.84] | +45.2% | [+44.7%, +45.7%] | 0.5 pts | 1.6 | yes | yes |
| natural-ptr | u64 | 16384 | churn | ordered | baseline | 6 | 99.5 | 83.2 | 0.84× [0.83, 0.84] | -19.6% | [-20.7%, -18.5%] | 1.1 pts | 1.5 | yes | yes |
| natural-ptr | u64 | 16384 | churn | ordered | btree-sets | 6 | 107 | 247 | 2.29× [2.23, 2.34] | +56.3% | [+55.2%, +57.3%] | 1.0 pts | 1.9 | yes | yes |
| natural-ptr | u64 | 16384 | build | ordered | baseline | 6 | 40.10 ms | 32.33 ms | 0.80× [0.80, 0.81] | -24.3% | [-25.6%, -22.9%] | 1.3 pts | 1.1 | yes | yes |
| natural-ptr | u64 | 16384 | build | ordered | btree-sets | 6 | 40.50 ms | 84.76 ms | 2.10× [2.06, 2.14] | +52.4% | [+51.6%, +53.3%] | 0.8 pts | 1.7 | yes | yes |
| natural-ptr | u64 | 262144 | valuesFor | ordered | baseline | 8 | 205 | 189 | 0.93× [0.91, 0.95] | -7.2% | [-9.5%, -5.0%] | 2.1 pts | 2.8 | no | yes |
| natural-ptr | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 247 | 634 | 2.54× [2.45, 2.64] | +60.7% | [+59.3%, +62.1%] | 1.3 pts | 3.7 | yes | yes |
| natural-ptr | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 9816 | 10.1 µs | 1.02× [0.96, 1.07] | +1.5% | [-3.7%, +6.8%] | 5.4 pts | 2.3 | no | no |
| natural-ptr | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 10.9 µs | 28.9 µs | 2.60× [2.54, 2.67] | +61.6% | [+60.6%, +62.5%] | 1.2 pts | 5.2 | yes | yes |
| natural-ptr | u64 | 262144 | churn | ordered | baseline | 8 | 367 | 339 | 0.90× [0.86, 0.94] | -11.2% | [-16.0%, -6.5%] | 5.3 pts | 1.1 | no | yes |
| natural-ptr | u64 | 262144 | churn | ordered | btree-sets | 8 | 456 | 714 | 1.59× [1.55, 1.63] | +37.0% | [+35.3%, +38.7%] | 1.7 pts | 3.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr u64 n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of -0.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesFor: ordered vs btree-sets: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-3.72%, 6.76%] includes zero
- natural-ptr u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 churn: ordered vs btree-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
