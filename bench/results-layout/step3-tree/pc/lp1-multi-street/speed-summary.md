| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 53.7 | 136 | 2.56× [2.54, 2.58] | +60.9% | [+60.6%, +61.2%] | 0.3 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage | 8 | 52.8 | 53.3 | 1.01× [0.99, 1.02] | +0.5% | [-1.2%, +2.2%] | 1.7 pts | 1.8 | yes | no |
| multi-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 2311 | 6965 | 3.02× [2.99, 3.04] | +66.8% | [+66.6%, +67.1%] | 0.3 pts | 0.6 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage | 8 | 2342 | 2322 | 0.99× [0.98, 1.00] | -0.7% | [-1.8%, +0.4%] | 1.9 pts | 0.9 | yes | no |
| multi-str | street | 4096 | prefix | ordered | btree-sets | 8 | 226 | 709 | 3.12× [3.09, 3.16] | +68.0% | [+67.6%, +68.3%] | 0.5 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage | 8 | 227 | 223 | 0.98× [0.98, 0.99] | -1.5% | [-2.4%, -0.6%] | 1.0 pts | 0.7 | yes | no |
| multi-str | street | 4096 | churn | ordered | btree-sets | 8 | 108 | 197 | 1.82× [1.79, 1.84] | +45.0% | [+44.3%, +45.8%] | 0.8 pts | 1.5 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage | 8 | 110 | 111 | 1.00× [0.99, 1.02] | +0.1% | [-1.3%, +1.6%] | 1.4 pts | 0.9 | yes | no |
| multi-str | street | 4096 | build | ordered | btree-sets | 8 | 4.48 ms | 8.00 ms | 1.78× [1.76, 1.80] | +43.7% | [+43.1%, +44.4%] | 0.9 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage | 8 | 4.46 ms | 4.55 ms | 1.01× [1.00, 1.02] | +1.1% | [-0.3%, +2.4%] | 1.3 pts | 1.3 | yes | no |
| multi-str | street | 16384 | valuesFor | ordered | btree-sets | 6 | 75.1 | 187 | 2.48× [2.39, 2.58] | +59.7% | [+58.2%, +61.2%] | 1.4 pts | 3.6 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage | 6 | 73.9 | 74.6 | 1.01× [1.00, 1.02] | +1.1% | [-0.2%, +2.4%] | 1.2 pts | 1.6 | yes | no |
| multi-str | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 2899 | 8042 | 2.77× [2.74, 2.79] | +63.8% | [+63.5%, +64.2%] | 0.3 pts | 1.2 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage | 6 | 2922 | 2932 | 1.00× [1.00, 1.01] | +0.5% | [-0.3%, +1.3%] | 0.8 pts | 0.9 | yes | no |
| multi-str | street | 16384 | prefix | ordered | btree-sets | 6 | 867 | 2991 | 3.46× [3.39, 3.52] | +71.1% | [+70.5%, +71.6%] | 0.5 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage | 6 | 836 | 833 | 1.00× [0.99, 1.02] | +0.3% | [-1.2%, +1.8%] | 1.4 pts | 0.7 | yes | no |
| multi-str | street | 16384 | churn | ordered | btree-sets | 6 | 154 | 279 | 1.81× [1.79, 1.84] | +44.9% | [+44.2%, +45.6%] | 0.7 pts | 1.8 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage | 6 | 157 | 158 | 1.00× [1.00, 1.00] | +0.0% | [-0.1%, +0.2%] | 0.2 pts | 0.2 | yes | no |
| multi-str | street | 16384 | build | ordered | btree-sets | 6 | 24.08 ms | 43.21 ms | 1.81× [1.79, 1.83] | +44.8% | [+44.1%, +45.4%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage | 6 | 24.04 ms | 27.29 ms | 1.14× [1.13, 1.14] | +12.2% | [+11.9%, +12.6%] | 0.4 pts | 0.5 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | btree-sets | 12 | 299 | 513 | 1.75× [1.72, 1.78] | +42.8% | [+41.9%, +43.7%] | 2.2 pts | 2.5 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage | 12 | 232 | 239 | 1.01× [0.98, 1.04] | +1.1% | [-1.8%, +4.0%] | 3.6 pts | 3.0 | no | no |
| multi-str | street | 212449 | valuesBetween | ordered | btree-sets | 12 | 8062 | 23.9 µs | 3.01× [2.96, 3.05] | +66.7% | [+66.3%, +67.2%] | 0.8 pts | 2.2 | yes | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage | 12 | 6766 | 6862 | 1.01× [0.98, 1.03] | +0.9% | [-1.6%, +3.3%] | 3.4 pts | 3.5 | no | no |
| multi-str | street | 212449 | prefix | ordered | btree-sets | 12 | 25.3 µs | 77.4 µs | 3.11× [3.00, 3.22] | +67.8% | [+66.7%, +68.9%] | 1.5 pts | 0.9 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage | 12 | 21.8 µs | 21.6 µs | 1.01× [0.99, 1.04] | +1.3% | [-1.3%, +4.0%] | 3.7 pts | 1.9 | no | no |
| multi-str | street | 212449 | churn | ordered | btree-sets | 12 | 571 | 743 | 1.32× [1.30, 1.34] | +24.4% | [+23.2%, +25.6%] | 2.2 pts | 1.4 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage | 12 | 497 | 494 | 0.96× [0.94, 0.99] | -3.8% | [-6.4%, -1.2%] | 5.0 pts | 0.9 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesFor: ordered vs ordered-lpage: the pooled difference of 0.50% does not clear the 0.80% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 valuesFor: ordered vs ordered-lpage: the pooled interval [-1.24%, 2.25%] includes zero
- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage: the pooled difference of -0.68% does not clear the 1.55% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage: the A/A validations found a systematic difference of +0.85% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage: the pooled interval [-1.80%, 0.44%] includes zero
- multi-str street n=4096 prefix: ordered vs ordered-lpage: the pooled difference of -1.54% does not clear the 1.62% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 churn: ordered vs ordered-lpage: the pooled difference of 0.15% does not clear the 0.53% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 churn: ordered vs ordered-lpage: the pooled interval [-1.28%, 1.57%] includes zero
- multi-str street n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.74% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 build: ordered vs ordered-lpage: the A/A validations found a systematic difference of +0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 build: ordered vs ordered-lpage: the pooled interval [-0.26%, 2.40%] includes zero
- multi-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage: the pooled difference of 1.10% does not clear the 1.12% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage: the pooled interval [-0.18%, 2.37%] includes zero
- multi-str street n=16384 valuesBetween: ordered vs ordered-lpage: the pooled difference of 0.48% does not clear the 0.94% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 valuesBetween: ordered vs ordered-lpage: the pooled interval [-0.34%, 1.30%] includes zero
- multi-str street n=16384 prefix: ordered vs ordered-lpage: the pooled difference of 0.29% does not clear the 1.06% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 prefix: ordered vs ordered-lpage: the pooled interval [-1.22%, 1.80%] includes zero
- multi-str street n=16384 churn: ordered vs ordered-lpage: the pooled difference of 0.05% does not clear the 0.89% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 churn: ordered vs ordered-lpage: the pooled interval [-0.15%, 0.24%] includes zero
- multi-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage: the pooled interval [-1.82%, 3.99%] includes zero
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: the pooled interval [-1.56%, 3.33%] includes zero
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: 5 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 prefix: ordered vs ordered-lpage: the pooled difference of 1.33% does not clear the 1.89% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 prefix: ordered vs ordered-lpage: the pooled interval [-1.33%, 4.00%] includes zero
- multi-str street n=212449 prefix: ordered vs ordered-lpage: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=212449 churn: ordered vs ordered-lpage: the A/A validations found a systematic difference of +1.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
