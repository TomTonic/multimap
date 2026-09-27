| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | str | 262144 | valuesFor | ordered | baseline | 6 | 375 | 321 | 0.85× [0.84, 0.90] | -17.6% | [-19.7%, -11.1%] | 4.1 pts | 1.8 | no |
| multi | str | 262144 | churn | ordered | baseline | 6 | 487 | 422 | 0.94× [0.80, 1.07] | -6.1% | [-25.6%, +6.2%] | 15.1 pts | 5.9 | no |
| multi | str | 1048576 | valuesFor | ordered | baseline | 6 | 486 | 442 | 0.91× [0.87, 0.96] | -9.4% | [-15.4%, -4.5%] | 5.2 pts | 2.6 | no |
| multi | str | 1048576 | churn | ordered | baseline | 6 | 619 | 565 | 0.88× [0.79, 1.05] | -14.2% | [-27.2%, +4.5%] | 15.1 pts | 6.0 | no |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 6 | 171 | 149 | 0.88× [0.80, 1.00] | -13.5% | [-24.7%, -0.1%] | 11.8 pts | 6.0 | no |
| multi | u64 | 262144 | churn | ordered | baseline | 6 | 304 | 264 | 0.98× [0.76, 1.21] | -1.6% | [-32.2%, +17.4%] | 23.6 pts | 7.7 | no |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 6 | 215 | 200 | 0.92× [0.90, 0.95] | -8.9% | [-11.1%, -5.3%] | 2.7 pts | 1.3 | no |
| multi | u64 | 1048576 | churn | ordered | baseline | 6 | 436 | 388 | 0.85× [0.74, 1.08] | -17.2% | [-34.7%, +7.5%] | 20.1 pts | 5.7 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
