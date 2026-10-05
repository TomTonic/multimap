| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 57.1 | 52.9 | 0.93× [0.92, 0.95] | -7.1% | [-9.1%, -5.0%] | 2.1 pts | 2.4 | no | yes |
| natural-ptr | street | 4096 | valuesFor | ordered | btree-sets | 8 | 57.5 | 136 | 2.36× [2.33, 2.39] | +57.6% | [+57.1%, +58.2%] | 0.5 pts | 1.3 | yes | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 2664 | 2343 | 0.89× [0.87, 0.91] | -12.4% | [-14.7%, -10.1%] | 2.2 pts | 1.1 | no | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 2679 | 4705 | 1.76× [1.74, 1.79] | +43.3% | [+42.5%, +44.1%] | 0.9 pts | 1.5 | yes | yes |
| natural-ptr | street | 4096 | prefix | ordered | baseline | 8 | 261 | 226 | 0.87× [0.86, 0.88] | -14.4% | [-15.6%, -13.1%] | 1.2 pts | 0.8 | yes | yes |
| natural-ptr | street | 4096 | prefix | ordered | btree-sets | 8 | 263 | 537 | 2.05× [2.02, 2.07] | +51.1% | [+50.6%, +51.7%] | 0.5 pts | 1.1 | yes | yes |
| natural-ptr | street | 4096 | churn | ordered | baseline | 8 | 106 | 98.4 | 0.93× [0.92, 0.94] | -7.7% | [-8.6%, -6.7%] | 1.2 pts | 1.0 | yes | yes |
| natural-ptr | street | 4096 | churn | ordered | btree-sets | 8 | 106 | 190 | 1.80× [1.77, 1.82] | +44.3% | [+43.5%, +45.1%] | 0.9 pts | 1.6 | yes | yes |
| natural-ptr | street | 4096 | build | ordered | baseline | 8 | 4.21 ms | 4.01 ms | 0.95× [0.94, 0.95] | -5.7% | [-6.6%, -4.8%] | 1.1 pts | 1.4 | yes | yes |
| natural-ptr | street | 4096 | build | ordered | btree-sets | 8 | 4.23 ms | 7.58 ms | 1.80× [1.78, 1.82] | +44.5% | [+43.8%, +45.1%] | 0.7 pts | 1.1 | yes | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | baseline | 6 | 78.3 | 73.8 | 0.95× [0.93, 0.96] | -5.6% | [-7.5%, -3.7%] | 1.8 pts | 2.8 | yes | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | btree-sets | 6 | 79.4 | 185 | 2.35× [2.28, 2.44] | +57.5% | [+56.1%, +59.0%] | 1.4 pts | 4.0 | yes | yes |
| natural-ptr | street | 16384 | valuesBetween | ordered | baseline | 6 | 3242 | 2942 | 0.91× [0.90, 0.91] | -10.4% | [-11.2%, -9.6%] | 0.8 pts | 1.2 | yes | yes |
| natural-ptr | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 3246 | 5619 | 1.72× [1.70, 1.74] | +41.9% | [+41.2%, +42.6%] | 0.7 pts | 2.1 | yes | yes |
| natural-ptr | street | 16384 | prefix | ordered | baseline | 6 | 959 | 841 | 0.87× [0.86, 0.88] | -14.3% | [-15.6%, -13.1%] | 1.2 pts | 0.6 | yes | yes |
| natural-ptr | street | 16384 | prefix | ordered | btree-sets | 6 | 982 | 2128 | 2.16× [2.14, 2.18] | +53.7% | [+53.4%, +54.0%] | 0.3 pts | 0.4 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | baseline | 6 | 145 | 135 | 0.93× [0.92, 0.95] | -7.1% | [-8.8%, -5.3%] | 1.7 pts | 1.4 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | btree-sets | 6 | 144 | 259 | 1.80× [1.77, 1.82] | +44.3% | [+43.6%, +45.1%] | 0.7 pts | 1.0 | yes | yes |
| natural-ptr | street | 16384 | build | ordered | baseline | 6 | 22.14 ms | 21.10 ms | 0.95× [0.95, 0.96] | -5.0% | [-5.8%, -4.2%] | 0.8 pts | 1.3 | yes | yes |
| natural-ptr | street | 16384 | build | ordered | btree-sets | 6 | 22.41 ms | 40.32 ms | 1.80× [1.77, 1.84] | +44.5% | [+43.4%, +45.6%] | 1.0 pts | 2.4 | yes | yes |
| natural-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 204 | 193 | 0.94× [0.92, 0.97] | -6.0% | [-9.2%, -2.9%] | 5.2 pts | 3.8 | no | yes |
| natural-ptr | street | 212449 | valuesFor | ordered | btree-sets | 8 | 266 | 482 | 1.80× [1.74, 1.86] | +44.3% | [+42.5%, +46.1%] | 1.8 pts | 3.0 | yes | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 6169 | 5774 | 0.96× [0.94, 0.99] | -3.7% | [-6.4%, -1.0%] | 4.3 pts | 3.9 | no | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 7369 | 17.8 µs | 2.42× [2.37, 2.47] | +58.7% | [+57.8%, +59.6%] | 1.2 pts | 2.3 | yes | yes |
| natural-ptr | street | 212449 | prefix | ordered | baseline | 8 | 18.7 µs | 17.8 µs | 0.96× [0.92, 0.99] | -4.7% | [-8.2%, -1.1%] | 4.9 pts | 3.7 | no | yes |
| natural-ptr | street | 212449 | prefix | ordered | btree-sets | 8 | 23.2 µs | 54.9 µs | 2.38× [2.29, 2.48] | +58.0% | [+56.4%, +59.6%] | 1.8 pts | 1.8 | yes | yes |
| natural-ptr | street | 212449 | churn | ordered | baseline | 8 | 442 | 428 | 0.99× [0.95, 1.02] | -1.4% | [-5.0%, +2.2%] | 4.3 pts | 1.0 | no | no |
| natural-ptr | street | 212449 | churn | ordered | btree-sets | 8 | 547 | 703 | 1.29× [1.27, 1.31] | +22.6% | [+21.3%, +23.8%] | 1.2 pts | 0.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.83% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.49% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=16384 build: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.67% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 prefix: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 churn: ordered vs baseline: the pooled interval [-5.02%, 2.16%] includes zero
