| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | street | 4096 | valuesFor | ordered | baseline | 8 | 84.9 | 81.9 | 0.97× [0.94, 0.99] | -3.4% | [-5.9%, -0.8%] | 2.5 pts | 2.7 | no | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 85.5 | 138 | 1.62× [1.61, 1.64] | +38.5% | [+38.0%, +38.9%] | 0.8 pts | 1.5 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 4732 | 5087 | 1.06× [1.05, 1.08] | +6.0% | [+5.0%, +7.1%] | 1.0 pts | 0.7 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 4744 | 6955 | 1.47× [1.46, 1.48] | +31.9% | [+31.4%, +32.4%] | 0.9 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 8 | 442 | 469 | 1.05× [1.04, 1.07] | +5.2% | [+4.1%, +6.3%] | 1.1 pts | 0.8 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 8 | 441 | 704 | 1.58× [1.55, 1.60] | +36.6% | [+35.6%, +37.6%] | 1.0 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 8 | 183 | 122 | 0.67× [0.66, 0.67] | -50.1% | [-50.9%, -49.2%] | 0.9 pts | 0.5 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 8 | 184 | 198 | 1.07× [1.07, 1.08] | +6.7% | [+6.2%, +7.3%] | 0.5 pts | 0.5 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 8 | 7.15 ms | 4.57 ms | 0.64× [0.64, 0.65] | -55.1% | [-56.9%, -53.3%] | 2.0 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 8 | 7.22 ms | 7.74 ms | 1.07× [1.06, 1.09] | +6.8% | [+5.6%, +8.1%] | 1.3 pts | 1.1 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 6 | 105 | 104 | 0.99× [0.98, 1.00] | -1.4% | [-2.5%, -0.4%] | 1.0 pts | 1.0 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 6 | 113 | 193 | 1.71× [1.66, 1.76] | +41.6% | [+39.8%, +43.3%] | 1.7 pts | 3.6 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 5392 | 5784 | 1.08× [1.07, 1.08] | +7.0% | [+6.2%, +7.8%] | 0.7 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 5397 | 8160 | 1.50× [1.49, 1.52] | +33.5% | [+32.7%, +34.4%] | 0.8 pts | 1.7 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 6 | 1763 | 1939 | 1.09× [1.07, 1.12] | +8.5% | [+6.6%, +10.5%] | 1.9 pts | 0.9 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 6 | 1801 | 3057 | 1.67× [1.60, 1.73] | +40.0% | [+37.7%, +42.3%] | 2.2 pts | 3.2 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 6 | 242 | 181 | 0.74× [0.73, 0.75] | -34.6% | [-36.2%, -33.1%] | 1.5 pts | 0.8 | yes | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 6 | 248 | 298 | 1.21× [1.18, 1.23] | +17.1% | [+15.2%, +18.9%] | 1.8 pts | 1.6 | yes | yes |
| natural-str | street | 16384 | build | ordered | baseline | 6 | 36.40 ms | 25.37 ms | 0.69× [0.69, 0.70] | -44.5% | [-45.2%, -43.8%] | 0.7 pts | 0.7 | yes | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 6 | 36.58 ms | 44.26 ms | 1.21× [1.19, 1.22] | +17.0% | [+15.8%, +18.2%] | 1.1 pts | 1.9 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | baseline | 8 | 235 | 248 | 1.08× [1.04, 1.13] | +7.6% | [+3.6%, +11.6%] | 3.7 pts | 3.4 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | btree-sets | 8 | 275 | 498 | 1.81× [1.78, 1.85] | +44.9% | [+43.7%, +46.0%] | 1.3 pts | 2.4 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 8462 | 9690 | 1.14× [1.12, 1.17] | +12.5% | [+10.6%, +14.3%] | 1.8 pts | 2.0 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 9172 | 23.5 µs | 2.50× [2.42, 2.60] | +60.1% | [+58.7%, +61.5%] | 1.9 pts | 4.1 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | baseline | 8 | 30.9 µs | 35.0 µs | 1.15× [1.12, 1.17] | +12.9% | [+11.1%, +14.7%] | 1.8 pts | 2.1 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | btree-sets | 8 | 31.3 µs | 76.6 µs | 2.41× [2.36, 2.45] | +58.4% | [+57.7%, +59.2%] | 0.8 pts | 0.9 | yes | yes |
| natural-str | street | 212449 | churn | ordered | baseline | 8 | 606 | 538 | 0.87× [0.85, 0.89] | -15.1% | [-17.6%, -12.6%] | 2.7 pts | 0.5 | no | yes |
| natural-str | street | 212449 | churn | ordered | btree-sets | 8 | 713 | 776 | 1.12× [1.09, 1.14] | +10.4% | [+8.5%, +12.3%] | 2.0 pts | 1.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 prefix: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.72% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
