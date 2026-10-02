| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | str | 4096 | valuesFor | ordered | baseline | 8 | 86.7 | 87.1 | 1.00× [1.00, 1.01] | +0.4% | [+0.1%, +0.7%] | 0.4 pts | 1.0 | yes | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 8 | 4960 | 4946 | 1.00× [0.99, 1.00] | -0.3% | [-0.6%, +0.1%] | 0.4 pts | 0.4 | yes | no |
| multi | str | 4096 | churn | ordered | baseline | 8 | 105 | 105 | 1.00× [0.99, 1.00] | -0.1% | [-0.6%, +0.3%] | 0.5 pts | 1.8 | yes | no |
| multi | str | 4096 | build | ordered | baseline | 8 | 10.57 ms | 10.55 ms | 1.00× [1.00, 1.00] | -0.1% | [-0.2%, +0.1%] | 0.2 pts | 0.8 | yes | no |
| multi | str | 16384 | valuesFor | ordered | baseline | 8 | 91.6 | 91.8 | 1.00× [0.99, 1.01] | -0.2% | [-1.3%, +0.8%] | 1.2 pts | 2.4 | yes | no |
| multi | str | 16384 | valuesBetween | ordered | baseline | 8 | 4716 | 4678 | 0.99× [0.99, 1.00] | -0.7% | [-1.1%, -0.3%] | 0.4 pts | 0.8 | yes | yes |
| multi | str | 16384 | churn | ordered | baseline | 8 | 126 | 127 | 1.00× [0.99, 1.01] | +0.0% | [-0.6%, +0.6%] | 0.7 pts | 2.6 | yes | no |
| multi | str | 16384 | build | ordered | baseline | 8 | 44.82 ms | 44.86 ms | 1.00× [1.00, 1.00] | +0.2% | [-0.1%, +0.4%] | 0.3 pts | 1.2 | yes | no |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 8 | 51.1 | 51.4 | 1.01× [1.00, 1.01] | +0.6% | [-0.2%, +1.3%] | 0.9 pts | 1.8 | yes | no |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 3876 | 3917 | 1.01× [1.01, 1.02] | +1.5% | [+0.8%, +2.1%] | 0.7 pts | 0.6 | yes | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 8 | 62.6 | 62.5 | 1.00× [0.99, 1.00] | -0.3% | [-1.0%, +0.3%] | 0.8 pts | 3.0 | yes | no |
| multi | u64 | 4096 | build | ordered | baseline | 8 | 6.76 ms | 6.77 ms | 1.00× [1.00, 1.00] | +0.0% | [-0.3%, +0.3%] | 0.3 pts | 0.9 | yes | no |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 8 | 56.0 | 56.1 | 1.01× [1.00, 1.01] | +0.6% | [-0.1%, +1.4%] | 0.9 pts | 2.3 | yes | no |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 4550 | 4514 | 0.99× [0.99, 1.00] | -0.7% | [-1.1%, -0.4%] | 0.4 pts | 1.0 | yes | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 8 | 67.3 | 67.7 | 1.01× [1.00, 1.02] | +0.7% | [-0.5%, +1.9%] | 1.4 pts | 1.3 | yes | no |
| multi | u64 | 16384 | build | ordered | baseline | 8 | 27.46 ms | 27.47 ms | 1.00× [1.00, 1.00] | -0.0% | [-0.3%, +0.2%] | 0.3 pts | 0.9 | yes | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 8 | 58.6 | 58.5 | 1.00× [0.99, 1.00] | -0.2% | [-0.6%, +0.2%] | 0.5 pts | 1.9 | yes | no |
| unique | str | 4096 | valuesBetween | ordered | baseline | 8 | 2466 | 2425 | 0.98× [0.98, 0.99] | -1.7% | [-2.4%, -1.1%] | 0.8 pts | 1.0 | yes | yes |
| unique | str | 4096 | churn | ordered | baseline | 8 | 126 | 126 | 1.00× [1.00, 1.01] | +0.4% | [+0.0%, +0.7%] | 0.4 pts | 0.6 | yes | yes |
| unique | str | 4096 | build | ordered | baseline | 8 | 1.90 ms | 1.89 ms | 1.00× [0.99, 1.01] | -0.1% | [-0.8%, +0.6%] | 0.8 pts | 1.0 | yes | no |
| unique | str | 16384 | valuesFor | ordered | baseline | 8 | 60.8 | 60.7 | 1.00× [0.99, 1.00] | -0.3% | [-0.9%, +0.3%] | 0.7 pts | 2.3 | yes | no |
| unique | str | 16384 | valuesBetween | ordered | baseline | 8 | 2214 | 2166 | 0.98× [0.97, 0.98] | -2.5% | [-2.8%, -2.1%] | 0.4 pts | 1.5 | yes | yes |
| unique | str | 16384 | churn | ordered | baseline | 8 | 148 | 147 | 0.99× [0.99, 1.00] | -0.6% | [-1.4%, +0.2%] | 1.0 pts | 2.3 | yes | no |
| unique | str | 16384 | build | ordered | baseline | 8 | 8.39 ms | 8.36 ms | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.3%] | 0.8 pts | 0.8 | yes | no |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 8 | 17.7 | 17.7 | 0.99× [0.99, 1.00] | -0.6% | [-0.9%, -0.3%] | 0.4 pts | 2.8 | yes | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 330 | 331 | 1.00× [1.00, 1.00] | +0.1% | [-0.2%, +0.4%] | 0.3 pts | 1.6 | yes | no |
| unique | u64 | 4096 | churn | ordered | baseline | 8 | 51.6 | 51.4 | 1.00× [0.99, 1.00] | -0.1% | [-0.7%, +0.4%] | 0.6 pts | 3.3 | yes | no |
| unique | u64 | 4096 | build | ordered | baseline | 8 | 858.3 µs | 864.5 µs | 1.01× [1.00, 1.02] | +0.6% | [-0.2%, +1.5%] | 1.0 pts | 1.3 | yes | no |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 8 | 27.3 | 27.4 | 1.00× [1.00, 1.01] | +0.4% | [-0.1%, +0.9%] | 0.6 pts | 7.3 | yes | no |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 412 | 414 | 1.00× [1.00, 1.01] | +0.5% | [+0.3%, +0.7%] | 0.2 pts | 0.9 | yes | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 8 | 77.7 | 77.9 | 1.00× [1.00, 1.01] | +0.2% | [-0.1%, +0.5%] | 0.4 pts | 1.3 | yes | no |
| unique | u64 | 16384 | build | ordered | baseline | 8 | 4.92 ms | 4.99 ms | 1.01× [1.00, 1.02] | +0.9% | [+0.0%, +1.8%] | 1.1 pts | 0.8 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.27% does not clear the 0.44% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.63%, 0.10%] includes zero
- multi str n=4096 churn: ordered vs baseline: the pooled difference of -0.15% does not clear the 0.25% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 churn: ordered vs baseline: the pooled interval [-0.58%, 0.28%] includes zero
- multi str n=4096 build: ordered vs baseline: the pooled difference of -0.08% does not clear the 0.19% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 build: ordered vs baseline: the pooled interval [-0.24%, 0.08%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.22% does not clear the 0.29% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.26%, 0.82%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=16384 churn: ordered vs baseline: the pooled difference of 0.03% does not clear the 0.19% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 churn: ordered vs baseline: the pooled interval [-0.58%, 0.63%] includes zero
- multi str n=16384 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=16384 build: ordered vs baseline: the pooled interval [-0.07%, 0.38%] includes zero
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.24%, 1.35%] includes zero
- multi u64 n=4096 churn: ordered vs baseline: the pooled difference of -0.33% does not clear the 0.45% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 churn: ordered vs baseline: the pooled interval [-0.97%, 0.31%] includes zero
- multi u64 n=4096 churn: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=4096 build: ordered vs baseline: the pooled difference of 0.01% does not clear the 0.17% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 build: ordered vs baseline: the pooled interval [-0.28%, 0.30%] includes zero
- multi u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.10%, 1.35%] includes zero
- multi u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=16384 churn: ordered vs baseline: the pooled difference of 0.70% does not clear the 0.87% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 churn: ordered vs baseline: the pooled interval [-0.50%, 1.91%] includes zero
- multi u64 n=16384 build: ordered vs baseline: the pooled difference of -0.03% does not clear the 0.12% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 build: ordered vs baseline: the pooled interval [-0.30%, 0.24%] includes zero
- unique str n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.20% does not clear the 0.22% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.62%, 0.21%] includes zero
- unique str n=4096 build: ordered vs baseline: the pooled difference of -0.13% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 build: ordered vs baseline: the pooled interval [-0.84%, 0.58%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.30% does not clear the 0.38% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.91%, 0.30%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 churn: ordered vs baseline: the pooled interval [-1.39%, 0.22%] includes zero
- unique str n=16384 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 build: ordered vs baseline: the pooled difference of -0.43% does not clear the 0.64% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=16384 build: ordered vs baseline: the pooled interval [-1.14%, 0.27%] includes zero
- unique u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.08% does not clear the 0.12% noise floor, the bound on what the harness reports between identical code in every process
- unique u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.19%, 0.35%] includes zero
- unique u64 n=4096 churn: ordered vs baseline: the pooled difference of -0.13% does not clear the 0.26% noise floor, the bound on what the harness reports between identical code in every process
- unique u64 n=4096 churn: ordered vs baseline: the pooled interval [-0.65%, 0.39%] includes zero
- unique u64 n=4096 churn: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 churn: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=4096 build: ordered vs baseline: the pooled interval [-0.23%, 1.51%] includes zero
- unique u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.05%, 0.94%] includes zero
- unique u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 7.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.06% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=16384 churn: ordered vs baseline: the pooled interval [-0.11%, 0.52%] includes zero
- unique u64 n=16384 build: ordered vs baseline: the pooled difference of 0.91% does not clear the 1.02% noise floor, the bound on what the harness reports between identical code in every process
