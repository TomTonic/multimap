| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 12 | 48.1 | 47.9 | 1.00× [0.98, 1.03] | +0.1% | [-2.4%, +2.6%] | 2.5 pts | 2.1 | no | no |
| multi | email | 4096 | valuesBetween | ordered | baseline | 12 | 2901 | 2885 | 0.99× [0.98, 1.01] | -0.6% | [-1.9%, +0.7%] | 2.8 pts | 1.3 | yes | no |
| multi | email | 4096 | prefix | ordered | baseline | 12 | 77.6 | 77.3 | 1.00× [0.98, 1.01] | -0.2% | [-1.7%, +1.2%] | 1.8 pts | 1.4 | yes | no |
| multi | email | 4096 | churn | ordered | baseline | 12 | 69.0 | 71.0 | 1.03× [1.02, 1.03] | +2.5% | [+2.1%, +2.9%] | 0.8 pts | 1.3 | yes | yes |
| multi | email | 4096 | build | ordered | baseline | 12 | 7.31 ms | 6.96 ms | 0.96× [0.95, 0.96] | -4.7% | [-5.1%, -4.3%] | 0.5 pts | 0.9 | yes | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 59.5 | 59.3 | 1.00× [0.99, 1.01] | +0.2% | [-0.7%, +1.1%] | 0.8 pts | 1.4 | yes | no |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3526 | 3500 | 1.00× [0.99, 1.00] | -0.4% | [-0.8%, +0.1%] | 0.4 pts | 0.7 | yes | no |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 91.1 | 91.7 | 1.01× [1.00, 1.01] | +0.7% | [-0.1%, +1.4%] | 0.7 pts | 1.3 | yes | no |
| multi | email | 16384 | churn | ordered | baseline | 6 | 97.5 | 100 | 1.02× [1.01, 1.03] | +1.9% | [+1.3%, +2.6%] | 0.6 pts | 1.1 | yes | yes |
| multi | email | 16384 | build | ordered | baseline | 6 | 36.77 ms | 35.66 ms | 0.97× [0.95, 0.98] | -3.5% | [-4.8%, -2.2%] | 1.3 pts | 1.5 | yes | yes |
| multi | path | 4096 | valuesFor | ordered | baseline | 8 | 99.4 | 100 | 1.01× [1.00, 1.02] | +1.0% | [+0.4%, +1.7%] | 0.7 pts | 0.7 | yes | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 8 | 3867 | 3910 | 1.02× [1.00, 1.04] | +1.8% | [+0.0%, +3.6%] | 2.2 pts | 0.9 | yes | no |
| multi | path | 4096 | prefix | ordered | baseline | 8 | 312 | 308 | 0.99× [0.98, 1.00] | -0.9% | [-1.6%, -0.3%] | 0.8 pts | 0.6 | yes | no |
| multi | path | 4096 | churn | ordered | baseline | 8 | 157 | 160 | 1.02× [1.01, 1.03] | +1.9% | [+0.6%, +3.3%] | 1.3 pts | 1.2 | yes | yes |
| multi | path | 4096 | build | ordered | baseline | 8 | 14.85 ms | 14.75 ms | 0.99× [0.98, 1.00] | -1.0% | [-1.7%, -0.3%] | 0.7 pts | 1.1 | yes | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 6 | 133 | 133 | 1.00× [0.99, 1.01] | +0.2% | [-0.5%, +1.0%] | 0.7 pts | 0.8 | yes | no |
| multi | path | 16384 | valuesBetween | ordered | baseline | 6 | 4615 | 4630 | 1.00× [1.00, 1.01] | +0.3% | [-0.3%, +0.9%] | 0.6 pts | 1.0 | yes | no |
| multi | path | 16384 | prefix | ordered | baseline | 6 | 744 | 734 | 0.99× [0.97, 1.01] | -0.8% | [-2.6%, +0.9%] | 1.7 pts | 0.6 | yes | no |
| multi | path | 16384 | churn | ordered | baseline | 6 | 216 | 221 | 1.02× [1.02, 1.03] | +2.2% | [+1.7%, +2.7%] | 0.5 pts | 0.9 | yes | yes |
| multi | path | 16384 | build | ordered | baseline | 6 | 74.50 ms | 75.45 ms | 1.01× [1.00, 1.02] | +1.3% | [+0.3%, +2.2%] | 0.9 pts | 1.5 | yes | yes |
| multi | str | 4096 | valuesFor | ordered | baseline | 12 | 60.4 | 60.3 | 1.00× [0.99, 1.01] | +0.0% | [-1.4%, +1.5%] | 2.2 pts | 2.3 | yes | no |
| multi | str | 4096 | valuesBetween | ordered | baseline | 12 | 3337 | 3334 | 1.00× [0.98, 1.02] | +0.0% | [-1.5%, +1.6%] | 3.1 pts | 1.4 | yes | no |
| multi | str | 4096 | prefix | ordered | baseline | 12 | 6995 | 7200 | 1.03× [0.99, 1.07] | +2.8% | [-0.7%, +6.3%] | 4.6 pts | 1.2 | no | no |
| multi | str | 4096 | churn | ordered | baseline | 12 | 85.4 | 89.5 | 1.05× [1.04, 1.06] | +4.7% | [+4.2%, +5.3%] | 0.9 pts | 1.4 | yes | yes |
| multi | str | 4096 | build | ordered | baseline | 12 | 9.00 ms | 8.75 ms | 0.97× [0.97, 0.98] | -2.7% | [-3.4%, -1.9%] | 0.9 pts | 1.5 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 71.0 | 71.0 | 1.00× [0.98, 1.01] | -0.2% | [-1.5%, +1.2%] | 1.3 pts | 2.1 | yes | no |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3640 | 3641 | 1.00× [1.00, 1.01] | +0.3% | [-0.1%, +0.7%] | 0.4 pts | 0.7 | yes | no |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 33.9 µs | 34.3 µs | 1.01× [1.00, 1.02] | +1.1% | [+0.3%, +2.0%] | 0.8 pts | 1.1 | yes | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 116 | 118 | 1.03× [1.02, 1.03] | +2.6% | [+1.9%, +3.2%] | 0.6 pts | 1.0 | yes | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 42.07 ms | 42.00 ms | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.3%] | 0.6 pts | 1.2 | yes | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 8 | 52.2 | 52.0 | 0.99× [0.98, 1.00] | -1.2% | [-2.3%, -0.1%] | 2.0 pts | 2.1 | yes | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 8 | 2268 | 2323 | 1.03× [1.01, 1.05] | +2.6% | [+0.8%, +4.5%] | 1.9 pts | 0.9 | yes | yes |
| multi | street | 4096 | prefix | ordered | baseline | 8 | 223 | 225 | 1.00× [0.99, 1.01] | +0.1% | [-1.3%, +1.4%] | 1.6 pts | 1.2 | yes | no |
| multi | street | 4096 | churn | ordered | baseline | 8 | 94.6 | 97.6 | 1.03× [1.02, 1.05] | +3.2% | [+1.8%, +4.7%] | 1.8 pts | 1.5 | yes | yes |
| multi | street | 4096 | build | ordered | baseline | 8 | 3.69 ms | 3.74 ms | 1.01× [1.01, 1.02] | +1.4% | [+0.5%, +2.3%] | 0.9 pts | 1.5 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 12 | 72.1 | 71.8 | 1.00× [0.99, 1.02] | +0.1% | [-1.4%, +1.6%] | 2.3 pts | 2.6 | yes | no |
| multi | street | 16384 | valuesBetween | ordered | baseline | 12 | 2882 | 2890 | 1.01× [1.00, 1.01] | +0.5% | [-0.1%, +1.1%] | 1.1 pts | 1.6 | yes | no |
| multi | street | 16384 | prefix | ordered | baseline | 12 | 817 | 823 | 1.00× [0.99, 1.02] | +0.4% | [-1.4%, +2.3%] | 1.9 pts | 0.8 | yes | no |
| multi | street | 16384 | churn | ordered | baseline | 12 | 131 | 136 | 1.03× [1.03, 1.04] | +3.1% | [+2.6%, +3.6%] | 0.9 pts | 1.2 | yes | yes |
| multi | street | 16384 | build | ordered | baseline | 12 | 21.50 ms | 20.42 ms | 0.95× [0.94, 0.96] | -5.3% | [-5.9%, -4.7%] | 0.7 pts | 0.9 | yes | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 10 | 36.6 | 36.5 | 1.00× [0.99, 1.02] | +0.4% | [-0.8%, +1.6%] | 1.4 pts | 1.2 | yes | no |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 10 | 2571 | 2545 | 0.99× [0.97, 1.01] | -1.3% | [-3.2%, +0.6%] | 1.9 pts | 0.5 | yes | no |
| multi | u64 | 4096 | churn | ordered | baseline | 10 | 50.2 | 51.5 | 1.03× [1.02, 1.03] | +2.5% | [+1.5%, +3.4%] | 1.0 pts | 1.8 | yes | yes |
| multi | u64 | 4096 | build | ordered | baseline | 10 | 5.63 ms | 5.58 ms | 0.99× [0.99, 1.00] | -0.8% | [-1.5%, -0.0%] | 0.8 pts | 1.4 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 45.4 | 45.5 | 1.00× [1.00, 1.01] | +0.3% | [-0.2%, +0.8%] | 0.5 pts | 0.9 | yes | no |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3487 | 3478 | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.6%] | 0.8 pts | 0.9 | yes | no |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 67.7 | 68.8 | 1.02× [1.02, 1.03] | +2.3% | [+1.8%, +2.9%] | 0.6 pts | 1.1 | yes | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 26.71 ms | 26.97 ms | 1.00× [1.00, 1.01] | +0.2% | [-0.5%, +0.8%] | 0.6 pts | 1.0 | yes | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 10 | 82.5 | 83.0 | 1.01× [0.99, 1.02] | +0.5% | [-0.7%, +1.7%] | 1.4 pts | 1.3 | yes | no |
| multi | url | 4096 | valuesBetween | ordered | baseline | 10 | 3755 | 3777 | 1.00× [0.99, 1.02] | +0.4% | [-1.5%, +2.4%] | 2.6 pts | 1.3 | yes | no |
| multi | url | 4096 | prefix | ordered | baseline | 10 | 181 | 181 | 1.00× [0.99, 1.01] | +0.0% | [-0.6%, +0.6%] | 0.9 pts | 0.8 | yes | no |
| multi | url | 4096 | churn | ordered | baseline | 10 | 126 | 130 | 1.03× [1.02, 1.04] | +2.8% | [+2.1%, +3.4%] | 1.1 pts | 1.1 | yes | yes |
| multi | url | 4096 | build | ordered | baseline | 10 | 12.45 ms | 12.13 ms | 0.97× [0.97, 0.98] | -2.7% | [-3.0%, -2.3%] | 0.4 pts | 0.6 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 107 | 109 | 1.01× [1.00, 1.01] | +0.7% | [+0.0%, +1.4%] | 0.6 pts | 0.8 | yes | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4398 | 4414 | 1.00× [0.99, 1.01] | +0.2% | [-0.5%, +1.0%] | 0.7 pts | 1.3 | yes | no |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 281 | 282 | 1.00× [0.99, 1.01] | +0.0% | [-1.0%, +1.0%] | 0.9 pts | 1.0 | yes | no |
| multi | url | 16384 | churn | ordered | baseline | 6 | 184 | 187 | 1.02× [1.01, 1.03] | +2.2% | [+1.2%, +3.1%] | 0.9 pts | 1.7 | yes | yes |
| multi | url | 16384 | build | ordered | baseline | 6 | 63.53 ms | 63.12 ms | 0.98× [0.97, 0.99] | -1.7% | [-2.7%, -0.7%] | 0.9 pts | 1.0 | yes | yes |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 12 | 53.6 | 53.6 | 1.00× [0.99, 1.01] | -0.3% | [-1.3%, +0.7%] | 1.5 pts | 1.7 | yes | no |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 3159 | 3143 | 0.99× [0.97, 1.02] | -0.7% | [-2.8%, +1.5%] | 2.7 pts | 1.1 | no | no |
| multi | uuid | 4096 | prefix | ordered | baseline | 12 | 89.6 | 89.6 | 1.00× [0.99, 1.00] | -0.4% | [-0.9%, +0.1%] | 0.9 pts | 0.8 | yes | no |
| multi | uuid | 4096 | churn | ordered | baseline | 12 | 79.2 | 81.7 | 1.03× [1.03, 1.03] | +3.0% | [+2.7%, +3.3%] | 0.9 pts | 1.7 | yes | yes |
| multi | uuid | 4096 | build | ordered | baseline | 12 | 8.19 ms | 7.81 ms | 0.96× [0.95, 0.96] | -4.7% | [-5.3%, -4.1%] | 0.8 pts | 1.0 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 63.3 | 62.7 | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.5%] | 1.0 pts | 1.6 | yes | no |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3576 | 3587 | 1.01× [1.00, 1.01] | +0.5% | [-0.1%, +1.1%] | 0.6 pts | 1.0 | yes | no |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 108 | 109 | 1.01× [1.01, 1.02] | +1.1% | [+0.6%, +1.6%] | 0.5 pts | 0.9 | yes | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 108 | 110 | 1.02× [1.00, 1.03] | +1.5% | [+0.5%, +2.6%] | 1.0 pts | 1.8 | yes | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 39.21 ms | 38.64 ms | 0.99× [0.98, 1.00] | -0.9% | [-1.9%, +0.1%] | 0.9 pts | 0.9 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.11% does not clear the 0.38% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.42%, 2.63%] includes zero
- multi email n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=4096 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.59% does not clear the 0.81% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.91%, 0.72%] includes zero
- multi email n=4096 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=4096 prefix: ordered vs baseline: the pooled difference of -0.23% does not clear the 0.97% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi email n=4096 prefix: ordered vs baseline: the pooled interval [-1.65%, 1.20%] includes zero
- multi email n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.20% does not clear the 0.64% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.69%, 1.08%] includes zero
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.38% does not clear the 0.43% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.83%, 0.08%] includes zero
- multi email n=16384 prefix: ordered vs baseline: the pooled interval [-0.07%, 1.44%] includes zero
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.83% does not clear the 1.93% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of -0.94% does not clear the 0.97% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.66% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.21% does not clear the 0.73% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.55%, 0.97%] includes zero
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.28% does not clear the 0.62% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.30%, 0.86%] includes zero
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of -0.85% does not clear the 6.80% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-2.59%, 0.90%] includes zero
- multi path n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi str n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.01% does not clear the 0.39% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.44%, 1.47%] includes zero
- multi str n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=4096 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.05% does not clear the 1.22% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.54%, 1.63%] includes zero
- multi str n=4096 prefix: ordered vs baseline: the pooled interval [-0.68%, 6.34%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.16% does not clear the 0.46% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi str n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.54%, 1.21%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.30% does not clear the 0.47% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.13%, 0.73%] includes zero
- multi str n=16384 build: ordered vs baseline: the pooled interval [-1.05%, 0.28%] includes zero
- multi street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 prefix: ordered vs baseline: the pooled difference of 0.08% does not clear the 1.00% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=4096 prefix: ordered vs baseline: the pooled interval [-1.28%, 1.44%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.08% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.43%, 1.59%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.12%, 1.15%] includes zero
- multi street n=16384 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=16384 prefix: ordered vs baseline: the pooled difference of 0.43% does not clear the 0.96% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=16384 prefix: ordered vs baseline: the pooled interval [-1.42%, 2.27%] includes zero
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.38% does not clear the 0.60% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.79%, 1.55%] includes zero
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.30% does not clear the 1.36% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-3.24%, 0.63%] includes zero
- multi u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.22%, 0.82%] includes zero
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.19% does not clear the 0.86% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.98%, 0.61%] includes zero
- multi u64 n=16384 build: ordered vs baseline: the pooled difference of 0.17% does not clear the 0.45% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 build: ordered vs baseline: the pooled interval [-0.49%, 0.82%] includes zero
- multi url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.68%, 1.68%] includes zero
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.42% does not clear the 1.02% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.51%, 2.36%] includes zero
- multi url n=4096 prefix: ordered vs baseline: the pooled difference of 0.03% does not clear the 0.61% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=4096 prefix: ordered vs baseline: the pooled interval [-0.59%, 0.65%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.23% does not clear the 1.00% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.53%, 0.98%] includes zero
- multi url n=16384 prefix: ordered vs baseline: the pooled difference of 0.04% does not clear the 1.03% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=16384 prefix: ordered vs baseline: the pooled interval [-0.95%, 1.03%] includes zero
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.29% does not clear the 0.71% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.26%, 0.68%] includes zero
- multi uuid n=4096 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.67% does not clear the 1.00% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.85%, 1.51%] includes zero
- multi uuid n=4096 prefix: ordered vs baseline: the pooled difference of -0.39% does not clear the 0.56% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 prefix: ordered vs baseline: the pooled interval [-0.89%, 0.11%] includes zero
- multi uuid n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.57% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.59%, 0.45%] includes zero
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.10%, 1.11%] includes zero
- multi uuid n=16384 build: ordered vs baseline: the pooled interval [-1.87%, 0.10%] includes zero
