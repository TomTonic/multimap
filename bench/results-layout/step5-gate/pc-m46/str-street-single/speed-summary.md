| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 6 | 82.7 | 66.9 | 0.81× [0.80, 0.81] | -24.2% | [-25.0%, -23.4%] | 0.8 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 6 | 83.2 | 93.4 | 1.10× [1.08, 1.13] | +9.5% | [+7.5%, +11.4%] | 1.8 pts | 2.6 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 6 | 2855 | 3774 | 1.32× [1.31, 1.33] | +24.4% | [+23.7%, +25.1%] | 0.6 pts | 0.6 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 6 | 2789 | 1162 | 0.41× [0.40, 0.42] | -146.0% | [-151.1%, -140.9%] | 4.9 pts | 1.2 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 6 | 399 | 381 | 0.95× [0.94, 0.97] | -4.8% | [-6.2%, -3.5%] | 1.3 pts | 1.1 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 6 | 396 | 193 | 0.48× [0.48, 0.49] | -107.8% | [-109.4%, -106.1%] | 1.5 pts | 0.8 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 6 | 219 | 134 | 0.61× [0.59, 0.62] | -64.4% | [-68.1%, -60.6%] | 3.6 pts | 1.2 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 6 | 218 | 146 | 0.67× [0.66, 0.68] | -50.0% | [-51.9%, -48.1%] | 1.8 pts | 1.3 | yes | yes |
| single-value-str | street | 4096 | build | ordered | baseline | 6 | 3.34 ms | 1.96 ms | 0.58× [0.57, 0.59] | -72.0% | [-75.2%, -68.9%] | 3.0 pts | 1.1 | yes | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 6 | 3.30 ms | 2.21 ms | 0.67× [0.65, 0.68] | -49.6% | [-53.1%, -46.0%] | 3.4 pts | 1.7 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 8 | 95.8 | 87.7 | 0.91× [0.90, 0.92] | -9.9% | [-11.3%, -8.5%] | 1.9 pts | 2.3 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 8 | 96.3 | 130 | 1.33× [1.30, 1.36] | +24.8% | [+23.1%, +26.5%] | 2.0 pts | 2.7 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 3211 | 4327 | 1.35× [1.33, 1.36] | +25.8% | [+25.0%, +26.7%] | 0.8 pts | 1.8 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 8 | 3195 | 1413 | 0.44× [0.43, 0.45] | -127.0% | [-130.9%, -123.0%] | 3.8 pts | 3.1 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 8 | 1062 | 1364 | 1.28× [1.26, 1.30] | +21.7% | [+20.6%, +22.9%] | 1.4 pts | 1.1 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 8 | 1053 | 508 | 0.48× [0.47, 0.49] | -109.8% | [-114.0%, -105.5%] | 4.1 pts | 1.6 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 8 | 246 | 173 | 0.71× [0.69, 0.73] | -40.2% | [-44.1%, -36.4%] | 4.5 pts | 1.9 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 8 | 255 | 212 | 0.83× [0.82, 0.84] | -20.7% | [-22.1%, -19.3%] | 1.6 pts | 1.1 | yes | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 8 | 14.86 ms | 10.46 ms | 0.70× [0.68, 0.71] | -43.2% | [-46.4%, -39.9%] | 3.7 pts | 1.4 | yes | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 8 | 15.07 ms | 12.33 ms | 0.81× [0.79, 0.83] | -23.9% | [-26.9%, -20.9%] | 3.2 pts | 1.6 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | baseline | 8 | 164 | 200 | 1.22× [1.11, 1.34] | +17.8% | [+10.3%, +25.3%] | 8.0 pts | 6.8 | no | yes |
| single-value-str | street | 212449 | valuesFor | ordered | btree-map | 8 | 173 | 273 | 1.61× [1.34, 2.04] | +38.1% | [+25.1%, +51.0%] | 12.5 pts | 11.1 | no | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 3775 | 5884 | 1.50× [1.42, 1.59] | +33.5% | [+29.7%, +37.3%] | 5.3 pts | 13.3 | no | yes |
| single-value-str | street | 212449 | valuesBetween | ordered | btree-map | 8 | 3684 | 2034 | 0.58× [0.49, 0.74] | -71.0% | [-106.1%, -35.9%] | 36.4 pts | 21.4 | no | yes |
| single-value-str | street | 212449 | prefix | ordered | baseline | 8 | 13.7 µs | 23.4 µs | 1.69× [1.63, 1.74] | +40.7% | [+38.8%, +42.5%] | 2.6 pts | 6.8 | yes | yes |
| single-value-str | street | 212449 | prefix | ordered | btree-map | 8 | 13.9 µs | 8278 | 0.63× [0.52, 0.79] | -58.9% | [-90.5%, -27.4%] | 32.2 pts | 25.4 | no | yes |
| single-value-str | street | 212449 | churn | ordered | baseline | 8 | 422 | 422 | 0.99× [0.97, 1.01] | -1.0% | [-3.0%, +1.1%] | 2.6 pts | 2.2 | no | no |
| single-value-str | street | 212449 | churn | ordered | btree-map | 8 | 507 | 543 | 1.06× [1.03, 1.10] | +5.8% | [+2.7%, +8.9%] | 3.1 pts | 2.5 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-str street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.85% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.68% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 6.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesFor: ordered vs btree-map: the processes scatter 11.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 13.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 21.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs baseline: the processes scatter 6.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 prefix: ordered vs btree-map: the processes scatter 25.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 churn: ordered vs baseline: the pooled interval [-3.00%, 1.05%] includes zero
- single-value-str street n=212449 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=212449 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str street n=212449 churn: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
