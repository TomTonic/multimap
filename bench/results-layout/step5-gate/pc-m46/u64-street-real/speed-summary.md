| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | street | 4096 | valuesFor | ordered | baseline | 6 | 75.6 | 53.9 | 0.71× [0.70, 0.72] | -41.1% | [-43.2%, -39.1%] | 2.0 pts | 1.7 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 6 | 76.7 | 134 | 1.77× [1.72, 1.82] | +43.5% | [+41.9%, +45.1%] | 1.5 pts | 3.3 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 6 | 1999 | 2535 | 1.27× [1.24, 1.30] | +21.1% | [+19.2%, +23.0%] | 1.8 pts | 1.4 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 2018 | 4544 | 2.26× [2.21, 2.31] | +55.8% | [+54.8%, +56.7%] | 0.9 pts | 1.9 | yes | yes |
| natural | street | 4096 | prefix | ordered | baseline | 6 | 266 | 254 | 0.95× [0.93, 0.96] | -5.8% | [-7.5%, -4.0%] | 1.6 pts | 1.5 | yes | yes |
| natural | street | 4096 | prefix | ordered | btree-sets | 6 | 268 | 526 | 1.97× [1.92, 2.02] | +49.2% | [+48.0%, +50.4%] | 1.1 pts | 2.4 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 6 | 160 | 103 | 0.65× [0.64, 0.67] | -53.3% | [-56.9%, -49.8%] | 3.4 pts | 2.3 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 6 | 160 | 188 | 1.18× [1.16, 1.20] | +15.2% | [+13.7%, +16.7%] | 1.4 pts | 1.4 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 6 | 6.29 ms | 4.15 ms | 0.66× [0.65, 0.67] | -51.5% | [-53.2%, -49.9%] | 1.6 pts | 1.6 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 6 | 6.34 ms | 7.58 ms | 1.19× [1.17, 1.20] | +15.7% | [+14.4%, +16.9%] | 1.2 pts | 1.4 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 91.7 | 75.3 | 0.82× [0.81, 0.84] | -21.3% | [-23.8%, -18.7%] | 2.8 pts | 2.6 | no | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 93.1 | 188 | 2.12× [1.87, 2.45] | +52.9% | [+46.6%, +59.2%] | 7.4 pts | 16.0 | no | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 2385 | 3140 | 1.33× [1.31, 1.34] | +24.6% | [+23.6%, +25.6%] | 0.9 pts | 1.7 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 2406 | 5679 | 2.57× [2.11, 3.29] | +61.1% | [+52.6%, +69.6%] | 10.1 pts | 34.3 | no | yes |
| natural | street | 16384 | prefix | ordered | baseline | 8 | 763 | 931 | 1.22× [1.20, 1.24] | +18.0% | [+16.3%, +19.6%] | 1.6 pts | 1.2 | yes | yes |
| natural | street | 16384 | prefix | ordered | btree-sets | 8 | 767 | 2080 | 2.85× [2.53, 3.26] | +64.9% | [+60.5%, +69.3%] | 5.1 pts | 10.5 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 195 | 142 | 0.73× [0.72, 0.75] | -36.3% | [-38.8%, -33.9%] | 2.4 pts | 1.9 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 201 | 277 | 1.37× [1.35, 1.40] | +27.2% | [+25.8%, +28.5%] | 1.5 pts | 1.7 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 30.24 ms | 22.17 ms | 0.73× [0.72, 0.74] | -36.8% | [-38.6%, -35.0%] | 1.8 pts | 2.0 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 30.57 ms | 41.72 ms | 1.37× [1.32, 1.43] | +27.2% | [+24.4%, +30.1%] | 2.9 pts | 3.5 | no | yes |
| natural | street | 212449 | valuesFor | ordered | baseline | 8 | 162 | 202 | 1.25× [1.15, 1.37] | +19.8% | [+12.8%, +26.8%] | 7.7 pts | 6.4 | no | yes |
| natural | street | 212449 | valuesFor | ordered | btree-sets | 8 | 201 | 499 | 2.50× [2.42, 2.59] | +60.0% | [+58.7%, +61.4%] | 1.4 pts | 2.2 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | baseline | 8 | 3278 | 5946 | 1.82× [1.75, 1.90] | +45.1% | [+42.7%, +47.4%] | 2.9 pts | 4.6 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 3724 | 16.5 µs | 4.47× [4.40, 4.54] | +77.6% | [+77.3%, +78.0%] | 0.4 pts | 1.1 | yes | yes |
| natural | street | 212449 | prefix | ordered | baseline | 8 | 10.2 µs | 19.1 µs | 1.87× [1.80, 1.95] | +46.5% | [+44.3%, +48.8%] | 2.7 pts | 4.3 | yes | yes |
| natural | street | 212449 | prefix | ordered | btree-sets | 8 | 12.4 µs | 56.1 µs | 4.64× [4.34, 4.97] | +78.4% | [+77.0%, +79.9%] | 1.4 pts | 3.6 | yes | yes |
| natural | street | 212449 | churn | ordered | baseline | 8 | 449 | 472 | 1.05× [1.00, 1.11] | +4.8% | [-0.2%, +9.9%] | 4.7 pts | 1.0 | no | no |
| natural | street | 212449 | churn | ordered | btree-sets | 8 | 528 | 749 | 1.40× [1.38, 1.42] | +28.7% | [+27.8%, +29.6%] | 1.1 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 prefix: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 16.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 34.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 prefix: ordered vs btree-sets: the processes scatter 10.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs baseline: the processes scatter 6.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs baseline: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 prefix: ordered vs btree-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 churn: ordered vs baseline: the pooled interval [-0.24%, 9.91%] includes zero
