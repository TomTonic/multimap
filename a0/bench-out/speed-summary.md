| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | str | 4096 | valuesFor | ordered | baseline | 4 | 87.2 | 87.2 | 1.00× [0.99, 1.01] | +0.4% | [-0.7%, +1.4%] | 0.7 pts | 2.3 | yes | no |
| multi | str | 4096 | valuesBetween | ordered | baseline | 4 | 4946 | 4953 | 1.00× [0.99, 1.01] | +0.3% | [-0.8%, +1.5%] | 0.7 pts | 0.7 | yes | no |
| multi | str | 4096 | churn | ordered | baseline | 4 | 105 | 104 | 1.00× [0.99, 1.00] | -0.3% | [-0.8%, +0.2%] | 0.3 pts | 1.0 | yes | no |
| multi | str | 4096 | build | ordered | baseline | 4 | 10.56 ms | 10.54 ms | 1.00× [1.00, 1.00] | -0.2% | [-0.4%, +0.1%] | 0.1 pts | 0.7 | yes | no |
| multi | str | 16384 | valuesFor | ordered | baseline | 4 | 91.2 | 90.8 | 1.00× [0.98, 1.02] | +0.1% | [-1.6%, +1.8%] | 1.0 pts | 3.3 | yes | no |
| multi | str | 16384 | valuesBetween | ordered | baseline | 4 | 4716 | 4673 | 0.99× [0.99, 0.99] | -1.1% | [-1.5%, -0.7%] | 0.2 pts | 0.5 | yes | no |
| multi | str | 16384 | churn | ordered | baseline | 4 | 127 | 127 | 1.00× [0.99, 1.00] | -0.1% | [-0.6%, +0.4%] | 0.3 pts | 1.3 | yes | no |
| multi | str | 16384 | build | ordered | baseline | 4 | 44.60 ms | 44.55 ms | 1.00× [1.00, 1.00] | -0.0% | [-0.5%, +0.5%] | 0.3 pts | 1.5 | yes | no |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 4 | 51.0 | 51.1 | 1.00× [1.00, 1.01] | +0.2% | [-0.2%, +0.7%] | 0.3 pts | 0.7 | yes | no |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 4 | 3869 | 3914 | 1.01× [0.99, 1.03] | +0.9% | [-0.7%, +2.6%] | 1.0 pts | 0.9 | yes | no |
| multi | u64 | 4096 | churn | ordered | baseline | 4 | 62.0 | 61.9 | 1.00× [0.99, 1.01] | +0.0% | [-0.8%, +0.9%] | 0.5 pts | 2.5 | yes | no |
| multi | u64 | 4096 | build | ordered | baseline | 4 | 6.78 ms | 6.78 ms | 1.00× [1.00, 1.00] | +0.1% | [-0.3%, +0.5%] | 0.2 pts | 0.7 | yes | no |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 8 | 55.8 | 56.3 | 1.00× [1.00, 1.01] | +0.4% | [+0.1%, +0.7%] | 0.8 pts | 2.1 | yes | no |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 4522 | 4508 | 1.00× [0.99, 1.00] | -0.4% | [-0.8%, -0.1%] | 0.3 pts | 0.9 | yes | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 8 | 69.0 | 68.9 | 1.00× [0.97, 1.04] | +0.2% | [-3.5%, +4.0%] | 2.5 pts | 2.0 | no | no |
| multi | u64 | 16384 | build | ordered | baseline | 8 | 27.48 ms | 27.51 ms | 1.00× [1.00, 1.00] | +0.1% | [-0.0%, +0.2%] | 0.2 pts | 0.7 | yes | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 4 | 57.4 | 57.2 | 1.00× [0.99, 1.00] | -0.4% | [-0.9%, +0.1%] | 0.3 pts | 1.5 | yes | no |
| unique | str | 4096 | valuesBetween | ordered | baseline | 4 | 2416 | 2373 | 0.97× [0.97, 0.98] | -2.8% | [-3.4%, -2.2%] | 0.4 pts | 0.4 | yes | yes |
| unique | str | 4096 | churn | ordered | baseline | 4 | 123 | 123 | 1.00× [1.00, 1.01] | +0.3% | [-0.3%, +0.9%] | 0.4 pts | 0.6 | yes | no |
| unique | str | 4096 | build | ordered | baseline | 4 | 1.85 ms | 1.86 ms | 1.00× [1.00, 1.01] | +0.5% | [-0.1%, +1.0%] | 0.4 pts | 0.6 | yes | no |
| unique | str | 16384 | valuesFor | ordered | baseline | 4 | 59.6 | 59.3 | 1.00× [0.99, 1.01] | -0.2% | [-1.1%, +0.6%] | 0.5 pts | 2.5 | yes | no |
| unique | str | 16384 | valuesBetween | ordered | baseline | 4 | 2161 | 2106 | 0.98× [0.97, 0.98] | -2.5% | [-3.1%, -1.9%] | 0.4 pts | 0.9 | yes | yes |
| unique | str | 16384 | churn | ordered | baseline | 4 | 139 | 138 | 1.00× [0.99, 1.00] | -0.4% | [-1.0%, +0.2%] | 0.4 pts | 1.2 | yes | no |
| unique | str | 16384 | build | ordered | baseline | 4 | 7.81 ms | 7.81 ms | 1.00× [0.99, 1.00] | -0.2% | [-0.8%, +0.5%] | 0.4 pts | 0.7 | yes | no |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 4 | 17.5 | 17.4 | 1.00× [0.98, 1.01] | -0.4% | [-1.7%, +0.9%] | 0.8 pts | 6.7 | yes | no |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 4 | 326 | 326 | 1.00× [1.00, 1.01] | +0.1% | [-0.5%, +0.7%] | 0.4 pts | 2.6 | yes | no |
| unique | u64 | 4096 | churn | ordered | baseline | 4 | 50.9 | 51.2 | 1.00× [1.00, 1.01] | +0.1% | [-0.5%, +0.6%] | 0.3 pts | 2.3 | yes | no |
| unique | u64 | 4096 | build | ordered | baseline | 4 | 844.7 µs | 842.2 µs | 1.00× [0.99, 1.02] | +0.4% | [-0.8%, +1.7%] | 0.8 pts | 1.0 | yes | no |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 8 | 26.8 | 26.9 | 1.01× [0.97, 1.05] | +0.9% | [-3.0%, +4.9%] | 2.5 pts | 25.9 | no | no |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 404 | 405 | 1.00× [1.00, 1.01] | +0.3% | [-0.2%, +0.8%] | 0.3 pts | 1.5 | yes | no |
| unique | u64 | 16384 | churn | ordered | baseline | 8 | 76.1 | 76.4 | 1.00× [1.00, 1.01] | +0.3% | [-0.2%, +0.8%] | 0.3 pts | 2.0 | yes | no |
| unique | u64 | 16384 | build | ordered | baseline | 8 | 4.64 ms | 4.66 ms | 1.00× [0.98, 1.02] | -0.1% | [-1.9%, +1.7%] | 1.1 pts | 0.7 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi str n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi str n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.36% does not clear the 0.62% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.69%, 1.41%] includes zero
- multi str n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.33% does not clear the 0.48% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.82%, 1.47%] includes zero
- multi str n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi str n=4096 churn: ordered vs baseline: the pooled interval [-0.75%, 0.18%] includes zero
- multi str n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi str n=4096 build: ordered vs baseline: the pooled difference of -0.17% does not clear the 0.55% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 build: ordered vs baseline: the pooled interval [-0.40%, 0.06%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi str n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.10% does not clear the 0.56% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.56%, 1.77%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.12% does not clear the 1.27% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi str n=16384 churn: ordered vs baseline: the pooled difference of -0.10% does not clear the 0.59% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 churn: ordered vs baseline: the pooled interval [-0.61%, 0.40%] includes zero
- multi str n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi str n=16384 build: ordered vs baseline: the pooled difference of -0.01% does not clear the 0.11% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 build: ordered vs baseline: the pooled interval [-0.48%, 0.47%] includes zero
- multi u64 n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.22% does not clear the 0.50% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.25%, 0.69%] includes zero
- multi u64 n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.67%, 2.55%] includes zero
- multi u64 n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi u64 n=4096 churn: ordered vs baseline: the pooled difference of 0.02% does not clear the 0.98% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 churn: ordered vs baseline: the pooled interval [-0.85%, 0.89%] includes zero
- multi u64 n=4096 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi u64 n=4096 build: ordered vs baseline: the pooled difference of 0.10% does not clear the 0.29% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 build: ordered vs baseline: the pooled interval [-0.26%, 0.46%] includes zero
- multi u64 n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.41% does not clear the 0.51% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=16384 churn: ordered vs baseline: the pooled difference of 0.25% does not clear the 0.46% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 churn: ordered vs baseline: the pooled interval [-3.50%, 3.99%] includes zero
- multi u64 n=16384 build: ordered vs baseline: the pooled difference of 0.11% does not clear the 0.22% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 build: ordered vs baseline: the pooled interval [-0.01%, 0.22%] includes zero
- unique str n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique str n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.86%, 0.13%] includes zero
- unique str n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique str n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique str n=4096 churn: ordered vs baseline: the pooled difference of 0.27% does not clear the 0.56% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 churn: ordered vs baseline: the pooled interval [-0.33%, 0.88%] includes zero
- unique str n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique str n=4096 build: ordered vs baseline: the pooled difference of 0.47% does not clear the 0.77% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 build: ordered vs baseline: the pooled interval [-0.09%, 1.04%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique str n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.12%, 0.62%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique str n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.13% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique str n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique str n=16384 churn: ordered vs baseline: the pooled difference of -0.40% does not clear the 0.41% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=16384 churn: ordered vs baseline: the pooled interval [-1.00%, 0.20%] includes zero
- unique str n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique str n=16384 build: ordered vs baseline: the pooled difference of -0.15% does not clear the 0.44% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=16384 build: ordered vs baseline: the pooled interval [-0.79%, 0.48%] includes zero
- unique u64 n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.71%, 0.92%] includes zero
- unique u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.13% does not clear the 0.23% noise floor, the bound on what the harness reports between identical code in every process
- unique u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.45%, 0.72%] includes zero
- unique u64 n=4096 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique u64 n=4096 churn: ordered vs baseline: the pooled difference of 0.06% does not clear the 0.15% noise floor, the bound on what the harness reports between identical code in every process
- unique u64 n=4096 churn: ordered vs baseline: the pooled interval [-0.45%, 0.58%] includes zero
- unique u64 n=4096 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique u64 n=4096 build: ordered vs baseline: the pooled difference of 0.43% does not clear the 1.52% noise floor, the bound on what the harness reports between identical code in every process
- unique u64 n=4096 build: ordered vs baseline: the pooled interval [-0.83%, 1.70%] includes zero
- unique u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.03%, 4.91%] includes zero
- unique u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 25.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.20%, 0.78%] includes zero
- unique u64 n=16384 churn: ordered vs baseline: the pooled interval [-0.19%, 0.80%] includes zero
- unique u64 n=16384 churn: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 build: ordered vs baseline: the pooled difference of -0.06% does not clear the 0.71% noise floor, the bound on what the harness reports between identical code in every process
- unique u64 n=16384 build: ordered vs baseline: the pooled interval [-1.86%, 1.73%] includes zero
