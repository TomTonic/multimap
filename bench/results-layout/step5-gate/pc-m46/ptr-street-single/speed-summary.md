| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| single-value-ptr | street | 4096 | valuesFor | ordered | baseline | 8 | 70.6 | 49.6 | 0.70× [0.69, 0.72] | -42.2% | [-44.6%, -39.7%] | 2.4 pts | 3.1 | yes | yes |
| single-value-ptr | street | 4096 | valuesFor | ordered | btree-map | 8 | 71.0 | 90.5 | 1.27× [1.26, 1.28] | +21.3% | [+20.7%, +21.8%] | 0.8 pts | 1.8 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | baseline | 8 | 1251 | 2016 | 1.61× [1.60, 1.62] | +37.9% | [+37.4%, +38.4%] | 1.0 pts | 1.3 | yes | yes |
| single-value-ptr | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1235 | 509 | 0.41× [0.41, 0.42] | -143.4% | [-146.7%, -140.1%] | 5.8 pts | 2.6 | yes | yes |
| single-value-ptr | street | 4096 | prefix | ordered | baseline | 8 | 231 | 229 | 0.99× [0.98, 0.99] | -1.5% | [-2.3%, -0.6%] | 1.5 pts | 2.2 | yes | yes |
| single-value-ptr | street | 4096 | prefix | ordered | btree-map | 8 | 230 | 135 | 0.58× [0.58, 0.59] | -71.3% | [-72.8%, -69.8%] | 4.1 pts | 3.1 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | baseline | 8 | 198 | 133 | 0.67× [0.66, 0.68] | -49.1% | [-51.2%, -47.0%] | 2.0 pts | 1.4 | yes | yes |
| single-value-ptr | street | 4096 | churn | ordered | btree-map | 8 | 199 | 143 | 0.72× [0.71, 0.73] | -39.1% | [-41.3%, -36.9%] | 2.1 pts | 1.6 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | baseline | 8 | 3.03 ms | 1.96 ms | 0.66× [0.64, 0.67] | -52.3% | [-55.6%, -49.0%] | 3.6 pts | 1.3 | yes | yes |
| single-value-ptr | street | 4096 | build | ordered | btree-map | 8 | 3.07 ms | 2.22 ms | 0.73× [0.70, 0.76] | -37.3% | [-42.3%, -32.4%] | 5.6 pts | 2.2 | no | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | baseline | 8 | 83.4 | 69.0 | 0.83× [0.81, 0.84] | -20.9% | [-23.1%, -18.7%] | 2.2 pts | 2.5 | no | yes |
| single-value-ptr | street | 16384 | valuesFor | ordered | btree-map | 8 | 84.3 | 127 | 1.49× [1.45, 1.53] | +32.9% | [+31.2%, +34.5%] | 2.0 pts | 3.4 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | baseline | 8 | 1423 | 2461 | 1.73× [1.71, 1.75] | +42.1% | [+41.5%, +42.7%] | 0.6 pts | 1.5 | yes | yes |
| single-value-ptr | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1412 | 642 | 0.45× [0.44, 0.46] | -123.8% | [-129.0%, -118.5%] | 7.5 pts | 5.4 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | baseline | 8 | 507 | 741 | 1.46× [1.44, 1.48] | +31.5% | [+30.5%, +32.5%] | 1.0 pts | 1.4 | yes | yes |
| single-value-ptr | street | 16384 | prefix | ordered | btree-map | 8 | 502 | 296 | 0.58× [0.56, 0.59] | -73.7% | [-78.1%, -69.2%] | 4.9 pts | 2.5 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | baseline | 8 | 211 | 170 | 0.81× [0.80, 0.82] | -23.6% | [-25.0%, -22.1%] | 1.4 pts | 1.1 | yes | yes |
| single-value-ptr | street | 16384 | churn | ordered | btree-map | 8 | 212 | 200 | 0.94× [0.93, 0.95] | -6.2% | [-7.0%, -5.3%] | 1.0 pts | 1.2 | yes | yes |
| single-value-ptr | street | 16384 | build | ordered | baseline | 8 | 13.55 ms | 9.79 ms | 0.72× [0.70, 0.75] | -38.0% | [-42.0%, -34.0%] | 4.0 pts | 2.0 | no | yes |
| single-value-ptr | street | 16384 | build | ordered | btree-map | 8 | 13.75 ms | 11.61 ms | 0.84× [0.83, 0.86] | -18.8% | [-21.1%, -16.5%] | 3.9 pts | 2.0 | no | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | baseline | 8 | 149 | 183 | 1.20× [1.01, 1.48] | +16.7% | [+1.1%, +32.2%] | 15.1 pts | 14.6 | no | yes |
| single-value-ptr | street | 212449 | valuesFor | ordered | btree-map | 8 | 157 | 306 | 1.80× [1.53, 2.17] | +44.3% | [+34.7%, +53.9%] | 11.1 pts | 12.0 | no | yes |
| single-value-ptr | street | 212449 | valuesBetween | ordered | baseline | 8 | 1814 | 3816 | 2.05× [1.66, 2.66] | +51.1% | [+39.9%, +62.4%] | 11.0 pts | 25.7 | no | yes |
| single-value-ptr | street | 212449 | valuesBetween | ordered | btree-map | 8 | 1755 | 1642 | 0.76× [0.59, 1.10] | -31.0% | [-70.8%, +8.7%] | 46.5 pts | 37.3 | no | no |
| single-value-ptr | street | 212449 | prefix | ordered | baseline | 8 | 5642 | 14.0 µs | 2.45× [2.14, 2.87] | +59.2% | [+53.3%, +65.2%] | 6.2 pts | 6.2 | yes | yes |
| single-value-ptr | street | 212449 | prefix | ordered | btree-map | 8 | 5605 | 5027 | 0.80× [0.63, 1.09] | -25.7% | [-59.7%, +8.3%] | 39.4 pts | 16.2 | no | no |
| single-value-ptr | street | 212449 | churn | ordered | baseline | 8 | 359 | 413 | 1.13× [1.10, 1.16] | +11.5% | [+9.1%, +13.9%] | 2.3 pts | 1.7 | no | yes |
| single-value-ptr | street | 212449 | churn | ordered | btree-map | 8 | 437 | 510 | 1.18× [1.15, 1.21] | +15.2% | [+12.8%, +17.6%] | 2.2 pts | 1.7 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- single-value-ptr street n=4096 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 prefix: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=4096 build: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 prefix: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=16384 build: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesFor: ordered vs baseline: the processes scatter 14.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesFor: ordered vs baseline: 7 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr street n=212449 valuesFor: ordered vs btree-map: the processes scatter 12.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs baseline: the processes scatter 25.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs btree-map: the pooled interval [-70.78%, 8.73%] includes zero
- single-value-ptr street n=212449 valuesBetween: ordered vs btree-map: the processes scatter 37.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 valuesBetween: ordered vs btree-map: 2 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-ptr street n=212449 prefix: ordered vs baseline: the processes scatter 6.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 prefix: ordered vs btree-map: the pooled interval [-59.71%, 8.28%] includes zero
- single-value-ptr street n=212449 prefix: ordered vs btree-map: the processes scatter 16.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-ptr street n=212449 prefix: ordered vs btree-map: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
