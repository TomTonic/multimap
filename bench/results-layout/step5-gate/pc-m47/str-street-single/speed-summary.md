| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 6 | 76.0 | 67.1 | 0.88× [0.86, 0.89] | -13.8% | [-15.6%, -12.0%] | 1.7 pts | 2.1 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 6 | 76.4 | 92.8 | 1.20× [1.18, 1.22] | +16.4% | [+15.0%, +17.8%] | 1.4 pts | 1.9 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 6 | 2741 | 3779 | 1.38× [1.36, 1.40] | +27.5% | [+26.6%, +28.4%] | 0.9 pts | 0.8 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 6 | 2666 | 1026 | 0.38× [0.37, 0.38] | -166.0% | [-170.4%, -161.5%] | 4.2 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 6 | 386 | 380 | 0.98× [0.97, 1.00] | -1.6% | [-2.9%, -0.3%] | 1.2 pts | 0.7 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 6 | 380 | 181 | 0.47× [0.46, 0.48] | -112.2% | [-116.0%, -108.4%] | 3.6 pts | 1.3 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 6 | 221 | 133 | 0.60× [0.59, 0.62] | -66.5% | [-70.6%, -62.4%] | 3.9 pts | 1.8 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 6 | 219 | 147 | 0.67× [0.66, 0.67] | -50.4% | [-51.7%, -49.0%] | 1.3 pts | 0.8 | yes | yes |
| single-value-str | street | 4096 | build | ordered | baseline | 6 | 3.42 ms | 2.01 ms | 0.59× [0.58, 0.61] | -69.0% | [-72.9%, -65.1%] | 3.7 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 6 | 3.48 ms | 2.36 ms | 0.67× [0.65, 0.69] | -49.7% | [-53.8%, -45.7%] | 3.9 pts | 1.3 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 8 | 89.1 | 87.9 | 0.99× [0.98, 1.00] | -0.7% | [-1.8%, +0.4%] | 1.8 pts | 2.1 | yes | no |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 8 | 90.1 | 130 | 1.43× [1.40, 1.46] | +29.9% | [+28.5%, +31.4%] | 2.2 pts | 3.4 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 3098 | 4372 | 1.41× [1.40, 1.42] | +29.2% | [+28.8%, +29.6%] | 0.4 pts | 0.8 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 8 | 3083 | 1327 | 0.43× [0.42, 0.44] | -133.6% | [-138.2%, -129.0%] | 5.8 pts | 4.0 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 8 | 1012 | 1386 | 1.36× [1.34, 1.39] | +26.7% | [+25.4%, +28.0%] | 1.2 pts | 1.2 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 8 | 1003 | 484 | 0.47× [0.46, 0.49] | -110.5% | [-116.4%, -104.6%] | 6.8 pts | 2.0 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 8 | 247 | 179 | 0.72× [0.71, 0.73] | -38.2% | [-40.0%, -36.4%] | 2.5 pts | 1.1 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 8 | 257 | 211 | 0.82× [0.81, 0.83] | -21.4% | [-22.9%, -19.9%] | 1.8 pts | 1.0 | yes | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 8 | 15.15 ms | 10.93 ms | 0.71× [0.68, 0.73] | -41.8% | [-46.2%, -37.3%] | 4.9 pts | 2.0 | no | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 8 | 15.45 ms | 12.59 ms | 0.81× [0.80, 0.82] | -23.2% | [-24.7%, -21.7%] | 1.6 pts | 0.9 | yes | yes |
| single-value-str | street | 212449 | valuesFor | ordered | baseline | 8 | 148 | 206 | 1.34× [1.16, 1.59] | +25.5% | [+13.8%, +37.1%] | 12.1 pts | 14.9 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | btree-map | 8 | 169 | 307 | 1.78× [1.45, 2.30] | +43.8% | [+31.2%, +56.4%] | 11.7 pts | 11.7 | no | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 3624 | 6153 | 1.68× [1.47, 1.96] | +40.5% | [+32.1%, +48.9%] | 8.0 pts | 17.3 | no | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | btree-map | 8 | 3519 | 2282 | 0.62× [0.51, 0.79] | -61.8% | [-97.0%, -26.5%] | 33.3 pts | 25.3 | no | yes |
| single-value-str | street | 212449 | prefix | ordered | baseline | 8 | 13.6 µs | 25.0 µs | 1.85× [1.68, 2.07] | +46.0% | [+40.4%, +51.7%] | 5.4 pts | 15.3 | no | yes |
| single-value-str | street | 212449 | prefix | ordered | btree-map | 8 | 13.6 µs | 8293 | 0.63× [0.52, 0.79] | -59.5% | [-91.7%, -27.3%] | 30.6 pts | 18.3 | no | yes |
| single-value-str | street | 212449 | churn | ordered | baseline | 8 | 414 | 416 | 0.99× [0.96, 1.02] | -1.3% | [-4.7%, +2.1%] | 3.3 pts | 2.1 | no | no |
| single-value-str | street | 212449 | churn | ordered | btree-map | 8 | 531 | 537 | 1.03× [1.02, 1.05] | +3.3% | [+1.7%, +4.9%] | 1.7 pts | 1.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.67% does not clear the 0.78% noise floor, the bound on what the harness reports between identical code in every process
- single-value-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.79%, 0.44%] includes zero
- single-value-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.59% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=16384 prefix: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.89% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 14.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 11.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 17.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 25.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs baseline: the processes scatter 15.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs btree-map: the processes scatter 18.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.61% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 churn: ordered vs baseline: the pooled interval [-4.67%, 2.09%] includes zero
- single-value-str street n=212449 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
