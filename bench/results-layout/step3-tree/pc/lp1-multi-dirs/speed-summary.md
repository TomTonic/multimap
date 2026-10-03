| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | btree-sets | 12 | 84.8 | 162 | 1.91× [1.88, 1.93] | +47.5% | [+46.9%, +48.1%] | 1.0 pts | 1.6 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage | 12 | 83.3 | 83.9 | 1.01× [1.00, 1.03] | +1.3% | [-0.0%, +2.6%] | 4.0 pts | 3.0 | yes | no |
| multi-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 12 | 2726 | 10.2 µs | 3.78× [3.76, 3.79] | +73.5% | [+73.4%, +73.6%] | 0.3 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage | 12 | 2757 | 2753 | 1.00× [0.98, 1.02] | -0.0% | [-2.5%, +2.4%] | 3.8 pts | 1.4 | no | no |
| multi-str | dirs | 4096 | prefix | ordered | btree-sets | 12 | 2889 | 14.5 µs | 5.00× [4.92, 5.07] | +80.0% | [+79.7%, +80.3%] | 0.3 pts | 0.5 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage | 12 | 2903 | 2867 | 0.98× [0.97, 0.99] | -1.9% | [-3.3%, -0.6%] | 2.4 pts | 1.0 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | btree-sets | 12 | 166 | 236 | 1.44× [1.42, 1.45] | +30.4% | [+29.7%, +31.1%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage | 12 | 166 | 166 | 1.00× [0.99, 1.01] | +0.1% | [-1.2%, +1.4%] | 1.7 pts | 1.3 | yes | no |
| multi-str | dirs | 4096 | build | ordered | btree-sets | 12 | 8.10 ms | 11.87 ms | 1.47× [1.45, 1.48] | +31.8% | [+31.1%, +32.6%] | 0.7 pts | 0.7 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage | 12 | 7.99 ms | 8.37 ms | 1.05× [1.04, 1.05] | +4.4% | [+3.7%, +5.1%] | 0.8 pts | 1.2 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | btree-sets | 6 | 120 | 227 | 1.89× [1.86, 1.92] | +47.1% | [+46.3%, +48.0%] | 0.8 pts | 1.3 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage | 6 | 116 | 117 | 1.01× [0.99, 1.03] | +1.3% | [-0.5%, +3.1%] | 1.7 pts | 1.6 | yes | no |
| multi-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 6 | 3558 | 11.2 µs | 3.14× [3.08, 3.19] | +68.1% | [+67.5%, +68.7%] | 0.5 pts | 1.2 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage | 6 | 3539 | 3582 | 1.01× [1.00, 1.02] | +0.8% | [+0.0%, +1.5%] | 0.7 pts | 0.8 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | btree-sets | 6 | 16.3 µs | 60.4 µs | 3.83× [3.78, 3.87] | +73.9% | [+73.6%, +74.2%] | 0.3 pts | 0.5 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage | 6 | 15.3 µs | 15.4 µs | 1.00× [0.99, 1.02] | +0.4% | [-1.2%, +2.0%] | 1.5 pts | 0.7 | yes | no |
| multi-str | dirs | 16384 | churn | ordered | btree-sets | 6 | 237 | 352 | 1.50× [1.48, 1.51] | +33.2% | [+32.5%, +34.0%] | 0.7 pts | 1.5 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage | 6 | 230 | 230 | 1.00× [0.99, 1.01] | -0.0% | [-0.8%, +0.7%] | 0.7 pts | 1.3 | yes | no |
| multi-str | dirs | 16384 | build | ordered | btree-sets | 6 | 44.21 ms | 66.23 ms | 1.49× [1.47, 1.51] | +32.9% | [+32.2%, +33.7%] | 0.8 pts | 1.4 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage | 6 | 44.89 ms | 49.88 ms | 1.11× [1.10, 1.12] | +10.0% | [+9.4%, +10.6%] | 0.6 pts | 0.6 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | btree-sets | 12 | 258 | 453 | 1.74× [1.69, 1.79] | +42.5% | [+40.7%, +44.2%] | 2.0 pts | 2.5 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage | 12 | 206 | 214 | 0.99× [0.98, 1.00] | -0.6% | [-1.8%, +0.5%] | 2.8 pts | 2.0 | yes | no |
| multi-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 12 | 7964 | 24.8 µs | 3.10× [3.05, 3.14] | +67.7% | [+67.3%, +68.2%] | 1.1 pts | 1.5 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage | 12 | 5455 | 5411 | 1.01× [0.99, 1.03] | +0.7% | [-1.4%, +2.9%] | 2.6 pts | 1.6 | no | no |
| multi-str | dirs | 86215 | prefix | ordered | btree-sets | 12 | 125.3 µs | 515.2 µs | 4.37× [4.25, 4.50] | +77.1% | [+76.5%, +77.8%] | 1.3 pts | 0.8 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage | 12 | 126.0 µs | 121.9 µs | 1.01× [0.99, 1.02] | +0.7% | [-1.1%, +2.4%] | 2.5 pts | 0.4 | yes | no |
| multi-str | dirs | 86215 | churn | ordered | btree-sets | 12 | 567 | 703 | 1.25× [1.23, 1.27] | +20.0% | [+19.0%, +21.0%] | 1.2 pts | 0.7 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage | 12 | 493 | 486 | 0.99× [0.96, 1.03] | -0.6% | [-3.6%, +2.5%] | 3.9 pts | 0.7 | no | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str dirs n=4096 valuesFor: ordered vs ordered-lpage: the pooled interval [-0.04%, 2.62%] includes zero
- multi-str dirs n=4096 valuesFor: ordered vs ordered-lpage: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=4096 valuesBetween: ordered vs ordered-lpage: the pooled difference of -0.04% does not clear the 1.26% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=4096 valuesBetween: ordered vs ordered-lpage: the pooled interval [-2.51%, 2.43%] includes zero
- multi-str dirs n=4096 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -0.71% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=4096 churn: ordered vs ordered-lpage: the pooled difference of 0.07% does not clear the 0.58% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=4096 churn: ordered vs ordered-lpage: the pooled interval [-1.22%, 1.36%] includes zero
- multi-str dirs n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the pooled interval [-0.53%, 3.10%] includes zero
- multi-str dirs n=16384 prefix: ordered vs ordered-lpage: the pooled difference of 0.40% does not clear the 12.97% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=16384 prefix: ordered vs ordered-lpage: the pooled interval [-1.17%, 1.96%] includes zero
- multi-str dirs n=16384 churn: ordered vs ordered-lpage: the pooled difference of -0.02% does not clear the 0.61% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=16384 churn: ordered vs ordered-lpage: the pooled interval [-0.78%, 0.73%] includes zero
- multi-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the pooled difference of -0.65% does not clear the 0.84% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the pooled interval [-1.76%, 0.47%] includes zero
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the pooled interval [-1.37%, 2.86%] includes zero
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage: the pooled difference of 0.68% does not clear the 2.15% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=86215 prefix: ordered vs ordered-lpage: the pooled interval [-1.07%, 2.42%] includes zero
- multi-str dirs n=86215 churn: ordered vs ordered-lpage: the pooled difference of -0.57% does not clear the 1.52% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=86215 churn: ordered vs ordered-lpage: the pooled interval [-3.64%, 2.50%] includes zero
