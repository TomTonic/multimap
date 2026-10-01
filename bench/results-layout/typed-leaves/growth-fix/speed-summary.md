| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | path | 4096 | churn | ordered | baseline | 6 | 173 | 167 | 0.97× [0.96, 0.98] | -3.2% | [-3.9%, -2.5%] | 0.7 pts | 0.7 | yes | yes |
| multi-str | path | 4096 | build | ordered | baseline | 6 | 16.96 ms | 16.10 ms | 0.95× [0.94, 0.96] | -5.1% | [-6.0%, -4.3%] | 0.8 pts | 1.5 | yes | yes |
| multi-str | path | 16384 | churn | ordered | baseline | 6 | 255 | 261 | 1.03× [1.02, 1.03] | +2.5% | [+1.7%, +3.3%] | 0.8 pts | 1.2 | yes | yes |
| multi-str | path | 16384 | build | ordered | baseline | 6 | 85.70 ms | 83.62 ms | 0.98× [0.98, 0.98] | -2.3% | [-2.5%, -2.0%] | 0.2 pts | 0.4 | yes | yes |
| multi-str | street | 4096 | churn | ordered | baseline | 6 | 114 | 108 | 0.94× [0.94, 0.95] | -6.0% | [-6.6%, -5.4%] | 0.6 pts | 0.6 | yes | yes |
| multi-str | street | 4096 | build | ordered | baseline | 6 | 4.51 ms | 4.06 ms | 0.90× [0.89, 0.91] | -11.5% | [-12.5%, -10.5%] | 1.0 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | churn | ordered | baseline | 6 | 150 | 143 | 0.96× [0.94, 0.97] | -4.6% | [-5.9%, -3.4%] | 1.2 pts | 1.2 | yes | yes |
| multi-str | street | 16384 | build | ordered | baseline | 6 | 23.91 ms | 21.63 ms | 0.90× [0.90, 0.91] | -10.6% | [-11.1%, -10.2%] | 0.4 pts | 0.5 | yes | yes |
| multi-str | u64 | 4096 | churn | ordered | baseline | 6 | 67.9 | 63.7 | 0.93× [0.93, 0.94] | -7.1% | [-7.8%, -6.3%] | 0.7 pts | 0.9 | yes | yes |
| multi-str | u64 | 4096 | build | ordered | baseline | 6 | 7.43 ms | 6.84 ms | 0.92× [0.92, 0.92] | -8.7% | [-9.2%, -8.3%] | 0.5 pts | 0.9 | yes | yes |
| multi-str | u64 | 16384 | churn | ordered | baseline | 10 | 91.5 | 84.6 | 0.92× [0.91, 0.94] | -8.6% | [-10.5%, -6.8%] | 2.1 pts | 1.7 | yes | yes |
| multi-str | u64 | 16384 | build | ordered | baseline | 10 | 36.84 ms | 32.33 ms | 0.88× [0.88, 0.89] | -13.3% | [-14.2%, -12.5%] | 1.0 pts | 1.4 | yes | yes |
| multi-str | url | 4096 | churn | ordered | baseline | 6 | 147 | 140 | 0.96× [0.94, 0.97] | -4.3% | [-5.9%, -2.8%] | 1.5 pts | 1.6 | yes | yes |
| multi-str | url | 4096 | build | ordered | baseline | 6 | 14.36 ms | 13.46 ms | 0.94× [0.93, 0.94] | -6.5% | [-7.1%, -5.9%] | 0.5 pts | 0.7 | yes | yes |
| multi-str | url | 16384 | churn | ordered | baseline | 6 | 221 | 222 | 1.01× [1.00, 1.02] | +0.8% | [-0.3%, +1.8%] | 1.0 pts | 1.1 | yes | no |
| multi-str | url | 16384 | build | ordered | baseline | 6 | 73.53 ms | 70.79 ms | 0.96× [0.96, 0.97] | -3.7% | [-4.5%, -2.9%] | 0.8 pts | 1.5 | yes | yes |
| multi-str | uuid | 4096 | churn | ordered | baseline | 6 | 95.9 | 90.3 | 0.94× [0.93, 0.95] | -6.1% | [-7.3%, -4.9%] | 1.2 pts | 1.0 | yes | yes |
| multi-str | uuid | 4096 | build | ordered | baseline | 6 | 9.83 ms | 8.96 ms | 0.91× [0.91, 0.92] | -9.6% | [-10.4%, -8.8%] | 0.8 pts | 1.6 | yes | yes |
| multi-str | uuid | 16384 | churn | ordered | baseline | 6 | 138 | 133 | 0.97× [0.96, 0.97] | -3.4% | [-4.2%, -2.7%] | 0.8 pts | 0.6 | yes | yes |
| multi-str | uuid | 16384 | build | ordered | baseline | 6 | 47.64 ms | 44.37 ms | 0.93× [0.92, 0.94] | -7.0% | [-8.2%, -5.8%] | 1.1 pts | 1.8 | yes | yes |
| unique-str | path | 4096 | churn | ordered | baseline | 6 | 203 | 203 | 1.00× [0.98, 1.01] | -0.4% | [-2.0%, +1.1%] | 1.5 pts | 1.3 | yes | no |
| unique-str | path | 4096 | build | ordered | baseline | 6 | 2.78 ms | 2.62 ms | 0.94× [0.93, 0.95] | -6.3% | [-7.6%, -5.0%] | 1.2 pts | 1.3 | yes | yes |
| unique-str | path | 16384 | churn | ordered | baseline | 12 | 261 | 270 | 1.04× [1.02, 1.06] | +3.9% | [+2.1%, +5.8%] | 2.2 pts | 2.2 | yes | yes |
| unique-str | path | 16384 | build | ordered | baseline | 12 | 13.65 ms | 13.63 ms | 1.00× [0.98, 1.01] | -0.4% | [-2.0%, +1.2%] | 1.5 pts | 1.3 | yes | no |
| unique-str | street | 4096 | churn | ordered | baseline | 6 | 122 | 119 | 0.98× [0.97, 0.99] | -2.4% | [-3.3%, -1.4%] | 0.9 pts | 0.6 | yes | yes |
| unique-str | street | 4096 | build | ordered | baseline | 6 | 1.72 ms | 1.61 ms | 0.93× [0.92, 0.94] | -7.6% | [-8.5%, -6.6%] | 0.9 pts | 1.2 | yes | yes |
| unique-str | street | 16384 | churn | ordered | baseline | 10 | 158 | 155 | 1.00× [0.98, 1.02] | +0.0% | [-1.8%, +1.9%] | 1.9 pts | 1.5 | yes | no |
| unique-str | street | 16384 | build | ordered | baseline | 10 | 8.49 ms | 8.14 ms | 0.96× [0.95, 0.96] | -4.4% | [-5.1%, -3.7%] | 0.7 pts | 1.0 | yes | yes |
| unique-str | u64 | 4096 | churn | ordered | baseline | 8 | 50.7 | 50.3 | 1.00× [0.98, 1.02] | +0.0% | [-1.9%, +1.9%] | 2.0 pts | 2.9 | yes | no |
| unique-str | u64 | 4096 | build | ordered | baseline | 8 | 827.4 µs | 813.8 µs | 0.99× [0.98, 0.99] | -1.5% | [-2.2%, -0.8%] | 0.7 pts | 1.2 | yes | yes |
| unique-str | u64 | 16384 | churn | ordered | baseline | 8 | 59.8 | 61.7 | 1.04× [1.02, 1.06] | +3.9% | [+2.1%, +5.8%] | 3.5 pts | 2.4 | yes | yes |
| unique-str | u64 | 16384 | build | ordered | baseline | 8 | 3.82 ms | 3.79 ms | 0.99× [0.97, 1.00] | -1.3% | [-2.7%, +0.0%] | 1.3 pts | 1.5 | yes | no |
| unique-str | url | 4096 | churn | ordered | baseline | 6 | 167 | 166 | 1.00× [0.99, 1.01] | -0.4% | [-1.4%, +0.6%] | 1.0 pts | 1.1 | yes | no |
| unique-str | url | 4096 | build | ordered | baseline | 6 | 2.31 ms | 2.20 ms | 0.94× [0.93, 0.95] | -6.0% | [-7.1%, -4.9%] | 1.0 pts | 1.1 | yes | yes |
| unique-str | url | 16384 | churn | ordered | baseline | 6 | 213 | 216 | 1.02× [1.00, 1.03] | +1.6% | [+0.3%, +3.0%] | 1.3 pts | 1.2 | yes | yes |
| unique-str | url | 16384 | build | ordered | baseline | 6 | 11.22 ms | 10.87 ms | 0.98× [0.97, 0.99] | -2.4% | [-3.6%, -1.1%] | 1.2 pts | 1.8 | yes | yes |
| unique-str | uuid | 4096 | churn | ordered | baseline | 6 | 87.8 | 87.4 | 0.99× [0.98, 1.00] | -1.2% | [-2.1%, -0.4%] | 0.8 pts | 0.7 | yes | no |
| unique-str | uuid | 4096 | build | ordered | baseline | 6 | 1.24 ms | 1.20 ms | 0.97× [0.96, 0.97] | -3.6% | [-4.0%, -3.2%] | 0.4 pts | 0.5 | yes | yes |
| unique-str | uuid | 16384 | churn | ordered | baseline | 6 | 103 | 106 | 1.03× [1.02, 1.04] | +3.0% | [+1.7%, +4.3%] | 1.2 pts | 1.0 | yes | yes |
| unique-str | uuid | 16384 | build | ordered | baseline | 6 | 5.82 ms | 5.71 ms | 0.98× [0.97, 0.99] | -1.9% | [-2.9%, -0.9%] | 0.9 pts | 1.5 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str url n=16384 churn: ordered vs baseline: the pooled interval [-0.32%, 1.83%] includes zero
- multi-str uuid n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of -0.15% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str path n=4096 churn: ordered vs baseline: the pooled difference of -0.44% does not clear the 0.64% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=4096 churn: ordered vs baseline: the pooled interval [-1.99%, 1.11%] includes zero
- unique-str path n=16384 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=16384 build: ordered vs baseline: the pooled difference of -0.41% does not clear the 0.51% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=16384 build: ordered vs baseline: the pooled interval [-2.00%, 1.19%] includes zero
- unique-str path n=16384 build: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str street n=16384 churn: ordered vs baseline: the pooled difference of 0.05% does not clear the 0.30% noise floor, the bound on what the harness reports between identical code in every process
- unique-str street n=16384 churn: ordered vs baseline: the pooled interval [-1.77%, 1.86%] includes zero
- unique-str u64 n=4096 churn: ordered vs baseline: the pooled difference of 0.00% does not clear the 0.65% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=4096 churn: ordered vs baseline: the pooled interval [-1.86%, 1.86%] includes zero
- unique-str u64 n=4096 churn: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=16384 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=16384 build: ordered vs baseline: the pooled interval [-2.71%, 0.02%] includes zero
- unique-str url n=4096 churn: ordered vs baseline: the pooled difference of -0.38% does not clear the 1.03% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=4096 churn: ordered vs baseline: the pooled interval [-1.38%, 0.62%] includes zero
- unique-str uuid n=4096 churn: ordered vs baseline: the pooled difference of -1.23% does not clear the 1.81% noise floor, the bound on what the harness reports between identical code in every process
