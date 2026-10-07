| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 66.6 | 56.7 | 0.85× [0.84, 0.85] | -18.3% | [-19.1%, -17.5%] | 1.8 pts | 1.7 | yes | yes |
| natural-ptr | street | 4096 | valuesFor | ordered | btree-sets | 8 | 67.7 | 138 | 2.06× [2.01, 2.10] | +51.4% | [+50.3%, +52.5%] | 1.1 pts | 2.6 | yes | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 1938 | 2567 | 1.34× [1.32, 1.37] | +25.5% | [+24.1%, +26.8%] | 1.7 pts | 1.3 | yes | yes |
| natural-ptr | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 1918 | 4832 | 2.56× [2.49, 2.63] | +60.9% | [+59.9%, +62.0%] | 1.1 pts | 2.5 | yes | yes |
| natural-ptr | street | 4096 | prefix | ordered | baseline | 8 | 263 | 258 | 0.98× [0.97, 0.99] | -2.0% | [-2.8%, -1.2%] | 0.8 pts | 0.7 | yes | yes |
| natural-ptr | street | 4096 | prefix | ordered | btree-sets | 8 | 264 | 548 | 2.10× [2.07, 2.14] | +52.4% | [+51.6%, +53.2%] | 0.9 pts | 2.5 | yes | yes |
| natural-ptr | street | 4096 | churn | ordered | baseline | 8 | 175 | 115 | 0.66× [0.65, 0.66] | -52.6% | [-53.8%, -51.4%] | 1.5 pts | 0.8 | yes | yes |
| natural-ptr | street | 4096 | churn | ordered | btree-sets | 8 | 176 | 206 | 1.17× [1.14, 1.20] | +14.6% | [+12.6%, +16.7%] | 2.0 pts | 2.0 | no | yes |
| natural-ptr | street | 4096 | build | ordered | baseline | 8 | 7.32 ms | 4.48 ms | 0.61× [0.60, 0.62] | -64.1% | [-66.9%, -61.4%] | 2.6 pts | 1.0 | yes | yes |
| natural-ptr | street | 4096 | build | ordered | btree-sets | 8 | 7.35 ms | 8.16 ms | 1.11× [1.08, 1.14] | +9.7% | [+7.3%, +12.1%] | 3.2 pts | 2.3 | no | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | baseline | 8 | 85.7 | 78.9 | 0.91× [0.90, 0.93] | -9.5% | [-11.7%, -7.2%] | 2.2 pts | 2.2 | no | yes |
| natural-ptr | street | 16384 | valuesFor | ordered | btree-sets | 8 | 86.8 | 191 | 2.30× [2.01, 2.70] | +56.6% | [+50.2%, +62.9%] | 6.1 pts | 15.5 | no | yes |
| natural-ptr | street | 16384 | valuesBetween | ordered | baseline | 8 | 2350 | 3204 | 1.37× [1.36, 1.37] | +26.8% | [+26.5%, +27.1%] | 0.3 pts | 0.5 | yes | yes |
| natural-ptr | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 2354 | 5816 | 2.68× [2.20, 3.43] | +62.7% | [+54.6%, +70.9%] | 8.0 pts | 31.6 | no | yes |
| natural-ptr | street | 16384 | prefix | ordered | baseline | 8 | 742 | 936 | 1.27× [1.25, 1.28] | +21.1% | [+20.1%, +22.1%] | 1.0 pts | 0.7 | yes | yes |
| natural-ptr | street | 16384 | prefix | ordered | btree-sets | 8 | 747 | 2188 | 3.05× [2.69, 3.52] | +67.2% | [+62.9%, +71.6%] | 4.2 pts | 8.7 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | baseline | 8 | 215 | 162 | 0.75× [0.75, 0.76] | -32.6% | [-32.8%, -32.4%] | 1.0 pts | 0.7 | yes | yes |
| natural-ptr | street | 16384 | churn | ordered | btree-sets | 8 | 219 | 290 | 1.30× [1.27, 1.34] | +23.2% | [+21.2%, +25.3%] | 2.2 pts | 2.6 | yes | yes |
| natural-ptr | street | 16384 | build | ordered | baseline | 8 | 35.58 ms | 24.65 ms | 0.71× [0.68, 0.74] | -41.1% | [-46.5%, -35.8%] | 5.2 pts | 2.2 | no | yes |
| natural-ptr | street | 16384 | build | ordered | btree-sets | 8 | 35.79 ms | 45.97 ms | 1.29× [1.25, 1.32] | +22.2% | [+20.1%, +24.3%] | 2.0 pts | 2.6 | yes | yes |
| natural-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 156 | 194 | 1.24× [1.15, 1.34] | +19.3% | [+13.3%, +25.3%] | 9.2 pts | 10.6 | no | yes |
| natural-ptr | street | 212449 | valuesFor | ordered | btree-sets | 8 | 210 | 531 | 2.59× [2.44, 2.77] | +61.4% | [+59.0%, +63.8%] | 3.0 pts | 4.3 | yes | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 3330 | 5877 | 1.74× [1.60, 1.90] | +42.6% | [+37.7%, +47.5%] | 7.2 pts | 13.2 | no | yes |
| natural-ptr | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 3835 | 17.1 µs | 4.47× [4.33, 4.62] | +77.6% | [+76.9%, +78.4%] | 0.7 pts | 3.4 | yes | yes |
| natural-ptr | street | 212449 | prefix | ordered | baseline | 8 | 10.4 µs | 19.6 µs | 1.87× [1.73, 2.03] | +46.5% | [+42.3%, +50.7%] | 6.7 pts | 13.9 | yes | yes |
| natural-ptr | street | 212449 | prefix | ordered | btree-sets | 8 | 12.2 µs | 59.3 µs | 4.87× [4.52, 5.28] | +79.5% | [+77.9%, +81.1%] | 1.8 pts | 7.3 | yes | yes |
| natural-ptr | street | 212449 | churn | ordered | baseline | 8 | 480 | 501 | 1.05× [1.00, 1.10] | +4.3% | [-0.1%, +8.7%] | 5.1 pts | 1.5 | no | no |
| natural-ptr | street | 212449 | churn | ordered | btree-sets | 8 | 564 | 782 | 1.35× [1.31, 1.38] | +25.9% | [+23.9%, +27.8%] | 1.9 pts | 3.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-ptr street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=4096 prefix: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=4096 churn: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=4096 build: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 15.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 31.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 prefix: ordered vs btree-sets: the processes scatter 8.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 churn: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=16384 build: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 10.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 13.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 prefix: ordered vs baseline: the A/A validations found a systematic difference of +2.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-ptr street n=212449 prefix: ordered vs baseline: the processes scatter 13.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 prefix: ordered vs btree-sets: the processes scatter 7.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-ptr street n=212449 churn: ordered vs baseline: the pooled interval [-0.07%, 8.69%] includes zero
- natural-ptr street n=212449 churn: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
