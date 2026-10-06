| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | u64 | 4096 | valuesFor | ordered | baseline | 8 | 50.5 | 40.8 | 0.80× [0.79, 0.81] | -25.2% | [-26.6%, -23.9%] | 1.5 pts | 1.3 | yes | yes |
| natural-ptr | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 50.6 | 162 | 3.20× [3.17, 3.24] | +68.8% | [+68.4%, +69.1%] | 0.3 pts | 1.8 | yes | yes |
| natural-ptr | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2964 | 2894 | 0.98× [0.96, 1.00] | -2.1% | [-4.3%, -0.0%] | 2.4 pts | 0.8 | no | yes |
| natural-ptr | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 2950 | 7509 | 2.55× [2.52, 2.58] | +60.8% | [+60.4%, +61.3%] | 0.5 pts | 1.0 | yes | yes |
| natural-ptr | u64 | 4096 | churn | ordered | baseline | 8 | 86.8 | 59.8 | 0.69× [0.68, 0.70] | -44.1% | [-46.0%, -42.2%] | 2.0 pts | 1.4 | yes | yes |
| natural-ptr | u64 | 4096 | churn | ordered | btree-sets | 8 | 87.5 | 160 | 1.83× [1.81, 1.86] | +45.4% | [+44.6%, +46.2%] | 0.8 pts | 1.3 | yes | yes |
| natural-ptr | u64 | 4096 | build | ordered | baseline | 8 | 12.72 ms | 6.26 ms | 0.49× [0.49, 0.50] | -102.6% | [-104.4%, -100.8%] | 2.4 pts | 1.0 | yes | yes |
| natural-ptr | u64 | 4096 | build | ordered | btree-sets | 8 | 12.81 ms | 15.11 ms | 1.19× [1.17, 1.21] | +16.1% | [+14.5%, +17.6%] | 1.7 pts | 1.0 | yes | yes |
| natural-ptr | u64 | 16384 | valuesFor | ordered | baseline | 6 | 53.9 | 50.9 | 0.94× [0.94, 0.95] | -5.9% | [-6.5%, -5.2%] | 0.6 pts | 1.1 | yes | yes |
| natural-ptr | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 54.1 | 199 | 3.67× [3.64, 3.71] | +72.8% | [+72.5%, +73.0%] | 0.2 pts | 1.3 | yes | yes |
| natural-ptr | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3826 | 3805 | 1.00× [0.99, 1.01] | -0.4% | [-1.3%, +0.5%] | 0.9 pts | 1.5 | yes | no |
| natural-ptr | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 3816 | 8054 | 2.11× [2.10, 2.13] | +52.7% | [+52.3%, +53.1%] | 0.4 pts | 1.6 | yes | yes |
| natural-ptr | u64 | 16384 | churn | ordered | baseline | 6 | 117 | 82.8 | 0.71× [0.70, 0.72] | -40.7% | [-43.3%, -38.0%] | 2.5 pts | 2.0 | yes | yes |
| natural-ptr | u64 | 16384 | churn | ordered | btree-sets | 6 | 119 | 237 | 1.99× [1.94, 2.04] | +49.8% | [+48.5%, +51.1%] | 1.2 pts | 1.9 | yes | yes |
| natural-ptr | u64 | 16384 | build | ordered | baseline | 6 | 44.14 ms | 30.63 ms | 0.70× [0.68, 0.71] | -43.8% | [-46.1%, -41.4%] | 2.3 pts | 0.9 | yes | yes |
| natural-ptr | u64 | 16384 | build | ordered | btree-sets | 6 | 43.02 ms | 80.07 ms | 1.86× [1.84, 1.88] | +46.2% | [+45.6%, +46.8%] | 0.6 pts | 1.7 | yes | yes |
| natural-ptr | u64 | 262144 | valuesFor | ordered | baseline | 8 | 204 | 199 | 0.96× [0.94, 0.98] | -4.3% | [-6.1%, -2.6%] | 2.4 pts | 2.3 | yes | yes |
| natural-ptr | u64 | 262144 | valuesFor | ordered | btree-sets | 8 | 236 | 617 | 2.62× [2.59, 2.65] | +61.8% | [+61.4%, +62.3%] | 0.6 pts | 1.7 | yes | yes |
| natural-ptr | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 8695 | 9502 | 1.07× [1.03, 1.12] | +6.8% | [+2.9%, +10.7%] | 4.3 pts | 2.3 | no | yes |
| natural-ptr | u64 | 262144 | valuesBetween | ordered | btree-sets | 8 | 10.5 µs | 28.6 µs | 2.74× [2.71, 2.78] | +63.5% | [+63.1%, +64.0%] | 0.4 pts | 1.6 | yes | yes |
| natural-ptr | u64 | 262144 | churn | ordered | baseline | 8 | 452 | 359 | 0.78× [0.76, 0.81] | -28.0% | [-31.8%, -24.1%] | 4.0 pts | 0.8 | no | yes |
| natural-ptr | u64 | 262144 | churn | ordered | btree-sets | 8 | 517 | 721 | 1.42× [1.35, 1.49] | +29.5% | [+26.1%, +32.9%] | 3.3 pts | 3.2 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr u64 n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr u64 n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of +1.93% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.28%, 0.54%] includes zero
- natural-ptr u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr u64 n=262144 churn: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
