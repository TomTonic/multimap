| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | u64 | 4096 | valuesFor | ordered | baseline | 8 | 47.8 | 40.4 | 0.83× [0.82, 0.84] | -20.1% | [-21.7%, -18.5%] | 1.9 pts | 2.4 | yes | yes |
| natural-ptr | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 47.7 | 162 | 3.41× [3.38, 3.44] | +70.6% | [+70.4%, +70.9%] | 0.3 pts | 1.9 | yes | yes |
| natural-ptr | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2987 | 2904 | 0.98× [0.96, 0.99] | -2.5% | [-4.5%, -0.5%] | 3.0 pts | 1.0 | yes | yes |
| natural-ptr | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 2979 | 7602 | 2.56× [2.54, 2.58] | +60.9% | [+60.7%, +61.2%] | 0.3 pts | 0.8 | yes | yes |
| natural-ptr | u64 | 4096 | churn | ordered | baseline | 8 | 82.3 | 64.4 | 0.78× [0.77, 0.78] | -28.8% | [-30.1%, -27.6%] | 1.2 pts | 1.1 | yes | yes |
| natural-ptr | u64 | 4096 | churn | ordered | btree-sets | 8 | 83.3 | 166 | 2.00× [1.98, 2.02] | +50.1% | [+49.6%, +50.6%] | 0.9 pts | 2.2 | yes | yes |
| natural-ptr | u64 | 4096 | build | ordered | baseline | 8 | 10.72 ms | 6.55 ms | 0.62× [0.61, 0.63] | -61.6% | [-63.9%, -59.2%] | 2.3 pts | 1.0 | yes | yes |
| natural-ptr | u64 | 4096 | build | ordered | btree-sets | 8 | 10.65 ms | 16.18 ms | 1.50× [1.47, 1.53] | +33.3% | [+31.9%, +34.7%] | 1.5 pts | 1.9 | yes | yes |
| natural-ptr | u64 | 16384 | valuesFor | ordered | baseline | 8 | 52.2 | 51.2 | 0.99× [0.96, 1.01] | -1.5% | [-4.0%, +1.0%] | 2.4 pts | 4.4 | no | no |
| natural-ptr | u64 | 16384 | valuesFor | ordered | btree-sets | 8 | 52.3 | 202 | 3.86× [3.79, 3.94] | +74.1% | [+73.6%, +74.6%] | 0.5 pts | 2.0 | yes | yes |
| natural-ptr | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 3855 | 3825 | 1.00× [0.98, 1.02] | -0.2% | [-2.4%, +2.0%] | 2.1 pts | 3.4 | no | no |
| natural-ptr | u64 | 16384 | valuesBetween | ordered | btree-sets | 8 | 3846 | 8172 | 2.14× [2.10, 2.17] | +53.2% | [+52.5%, +54.0%] | 0.8 pts | 2.8 | yes | yes |
| natural-ptr | u64 | 16384 | churn | ordered | baseline | 8 | 105 | 90.6 | 0.86× [0.84, 0.87] | -17.0% | [-19.0%, -15.0%] | 2.0 pts | 2.2 | yes | yes |
| natural-ptr | u64 | 16384 | churn | ordered | btree-sets | 8 | 110 | 254 | 2.29× [2.26, 2.33] | +56.3% | [+55.7%, +57.0%] | 0.8 pts | 1.2 | yes | yes |
| natural-ptr | u64 | 16384 | build | ordered | baseline | 8 | 41.98 ms | 34.82 ms | 0.82× [0.81, 0.83] | -21.4% | [-22.7%, -20.1%] | 2.3 pts | 1.4 | yes | yes |
| natural-ptr | u64 | 16384 | build | ordered | btree-sets | 8 | 42.50 ms | 88.46 ms | 2.09× [2.03, 2.15] | +52.2% | [+50.8%, +53.6%] | 1.4 pts | 2.2 | yes | yes |
| natural-ptr | u64 | 262144 | valuesFor | ordered | baseline | 8 | 187 | 187 | 1.00× [0.97, 1.03] | -0.4% | [-3.4%, +2.6%] | 3.0 pts | 4.3 | no | no |
| natural-ptr | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 230 | 638 | 2.77× [2.73, 2.82] | +63.9% | [+63.4%, +64.5%] | 0.6 pts | 1.5 | yes | yes |
| natural-ptr | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 9800 | 10.2 µs | 1.02× [1.00, 1.05] | +2.3% | [+0.3%, +4.3%] | 2.5 pts | 1.0 | no | yes |
| natural-ptr | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 10.9 µs | 28.8 µs | 2.66× [2.63, 2.70] | +62.5% | [+62.0%, +63.0%] | 0.5 pts | 1.5 | yes | yes |
| natural-ptr | u64 | 262144 | churn | ordered | baseline | 8 | 386 | 364 | 0.94× [0.90, 0.98] | -6.7% | [-11.2%, -2.2%] | 6.4 pts | 1.4 | no | yes |
| natural-ptr | u64 | 262144 | churn | ordered | btree-sets | 8 | 458 | 712 | 1.55× [1.50, 1.61] | +35.6% | [+33.2%, +38.1%] | 2.4 pts | 4.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=4096 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.97%, 0.98%] includes zero
- natural-ptr u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 7 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr u64 n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.19% does not clear the 0.33% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-2.40%, 2.02%] includes zero
- natural-ptr u64 n=16384 valuesBetween: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=16384 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr u64 n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=16384 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=16384 build: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.39% does not clear the 0.72% noise floor, the bound on what the harness reports between identical code in every process
- natural-ptr u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.40%, 2.62%] includes zero
- natural-ptr u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-ptr u64 n=262144 churn: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
