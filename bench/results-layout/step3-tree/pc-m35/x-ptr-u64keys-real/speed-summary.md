| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | u64 | 4096 | valuesFor | ordered | baseline | 8 | 41.1 | 38.7 | 0.95× [0.94, 0.97] | -5.2% | [-6.9%, -3.6%] | 1.9 pts | 1.7 | yes | yes |
| natural-ptr | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 40.9 | 160 | 3.94× [3.90, 3.99] | +74.6% | [+74.4%, +74.9%] | 0.3 pts | 2.0 | yes | yes |
| natural-ptr | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2932 | 2711 | 0.94× [0.92, 0.96] | -6.2% | [-8.7%, -3.7%] | 2.6 pts | 0.8 | no | yes |
| natural-ptr | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 2933 | 7490 | 2.59× [2.55, 2.64] | +61.5% | [+60.8%, +62.1%] | 0.7 pts | 1.1 | yes | yes |
| natural-ptr | u64 | 4096 | churn | ordered | baseline | 8 | 59.0 | 53.4 | 0.91× [0.90, 0.92] | -10.1% | [-11.2%, -9.1%] | 1.4 pts | 1.6 | yes | yes |
| natural-ptr | u64 | 4096 | churn | ordered | btree-sets | 8 | 58.8 | 158 | 2.70× [2.67, 2.73] | +62.9% | [+62.6%, +63.3%] | 0.4 pts | 1.5 | yes | yes |
| natural-ptr | u64 | 4096 | build | ordered | baseline | 8 | 6.17 ms | 5.64 ms | 0.91× [0.90, 0.91] | -10.3% | [-11.1%, -9.4%] | 1.2 pts | 1.6 | yes | yes |
| natural-ptr | u64 | 4096 | build | ordered | btree-sets | 8 | 6.16 ms | 15.16 ms | 2.47× [2.45, 2.49] | +59.5% | [+59.1%, +59.8%] | 0.3 pts | 1.1 | yes | yes |
| natural-ptr | u64 | 16384 | valuesFor | ordered | baseline | 6 | 50.6 | 48.3 | 0.96× [0.95, 0.97] | -4.2% | [-5.2%, -3.3%] | 0.9 pts | 1.4 | yes | yes |
| natural-ptr | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 50.6 | 197 | 3.89× [3.86, 3.92] | +74.3% | [+74.1%, +74.5%] | 0.2 pts | 1.0 | yes | yes |
| natural-ptr | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3789 | 3657 | 0.97× [0.96, 0.97] | -3.3% | [-3.7%, -2.8%] | 0.4 pts | 0.7 | yes | yes |
| natural-ptr | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 3818 | 8031 | 2.10× [2.08, 2.12] | +52.4% | [+51.9%, +52.9%] | 0.5 pts | 1.9 | yes | yes |
| natural-ptr | u64 | 16384 | churn | ordered | baseline | 6 | 76.6 | 66.7 | 0.87× [0.86, 0.88] | -15.1% | [-16.3%, -13.8%] | 1.2 pts | 1.2 | yes | yes |
| natural-ptr | u64 | 16384 | churn | ordered | btree-sets | 6 | 81.5 | 233 | 2.84× [2.78, 2.91] | +64.8% | [+64.1%, +65.6%] | 0.7 pts | 1.5 | yes | yes |
| natural-ptr | u64 | 16384 | build | ordered | baseline | 6 | 29.64 ms | 26.78 ms | 0.90× [0.90, 0.91] | -10.7% | [-11.7%, -9.7%] | 0.9 pts | 1.2 | yes | yes |
| natural-ptr | u64 | 16384 | build | ordered | btree-sets | 6 | 29.95 ms | 80.22 ms | 2.67× [2.65, 2.69] | +62.6% | [+62.3%, +62.9%] | 0.3 pts | 0.7 | yes | yes |
| natural-ptr | u64 | 262144 | valuesFor | ordered | baseline | 8 | 197 | 183 | 0.95× [0.91, 0.99] | -5.5% | [-10.4%, -0.6%] | 5.5 pts | 4.7 | no | yes |
| natural-ptr | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 224 | 615 | 2.76× [2.66, 2.87] | +63.8% | [+62.4%, +65.2%] | 1.5 pts | 3.5 | yes | yes |
| natural-ptr | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 9253 | 9778 | 1.07× [1.01, 1.14] | +6.8% | [+1.0%, +12.6%] | 5.9 pts | 4.2 | no | yes |
| natural-ptr | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 10.9 µs | 28.7 µs | 2.63× [2.53, 2.73] | +62.0% | [+60.5%, +63.4%] | 1.4 pts | 4.6 | yes | yes |
| natural-ptr | u64 | 262144 | churn | ordered | baseline | 8 | 325 | 288 | 0.86× [0.83, 0.90] | -16.2% | [-21.1%, -11.2%] | 5.3 pts | 0.7 | no | yes |
| natural-ptr | u64 | 262144 | churn | ordered | btree-sets | 8 | 414 | 702 | 1.73× [1.71, 1.75] | +42.2% | [+41.4%, +42.9%] | 1.9 pts | 2.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr u64 n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr u64 n=262144 valuesFor: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of -0.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr u64 n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
