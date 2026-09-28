| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 105 | 121 | 1.15× [1.13, 1.18] | +13.4% | [+11.7%, +15.1%] | 1.6 pts | 2.2 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4113 | 4370 | 1.06× [1.05, 1.07] | +5.6% | [+4.4%, +6.8%] | 1.2 pts | 0.7 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 328 | 366 | 1.12× [1.11, 1.13] | +10.7% | [+9.9%, +11.4%] | 0.7 pts | 0.3 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 163 | 200 | 1.23× [1.22, 1.24] | +18.6% | [+17.9%, +19.3%] | 0.7 pts | 0.5 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 11.63 ms | 16.63 ms | 1.43× [1.42, 1.44] | +30.1% | [+29.6%, +30.7%] | 0.6 pts | 0.9 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 596 | 696 | 1.17× [1.16, 1.19] | +14.7% | [+13.7%, +15.8%] | 1.5 pts | 1.4 | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 13.6 µs | 14.8 µs | 1.08× [1.07, 1.09] | +7.4% | [+6.3%, +8.5%] | 1.6 pts | 1.6 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 12.1 µs | 13.2 µs | 1.07× [1.04, 1.09] | +6.1% | [+4.0%, +8.3%] | 3.0 pts | 0.2 | no |
| multi | path | 262144 | churn | ordered | baseline | 10 | 992 | 1116 | 1.10× [1.08, 1.12] | +9.1% | [+7.4%, +10.8%] | 2.4 pts | 2.0 | yes |
| multi | str | 4096 | valuesFor | ordered | baseline | 10 | 62.3 | 77.4 | 1.24× [1.24, 1.25] | +19.4% | [+19.0%, +19.7%] | 0.5 pts | 0.6 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 10 | 3505 | 3660 | 1.04× [1.03, 1.06] | +4.2% | [+2.9%, +5.5%] | 1.8 pts | 0.8 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 10 | 7863 | 7977 | 1.01× [0.99, 1.03] | +0.8% | [-1.2%, +2.8%] | 2.8 pts | 0.8 | yes |
| multi | str | 4096 | churn | ordered | baseline | 10 | 85.4 | 118 | 1.38× [1.35, 1.42] | +27.7% | [+26.0%, +29.4%] | 2.4 pts | 1.3 | yes |
| multi | str | 4096 | build | ordered | baseline | 10 | 6.60 ms | 10.65 ms | 1.63× [1.61, 1.66] | +38.8% | [+37.8%, +39.7%] | 1.3 pts | 1.9 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 352 | 432 | 1.22× [1.20, 1.24] | +18.0% | [+16.4%, +19.7%] | 2.3 pts | 2.6 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 11.0 µs | 12.3 µs | 1.12× [1.11, 1.13] | +11.0% | [+10.1%, +11.8%] | 1.2 pts | 1.0 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.71 ms | 1.99 ms | 1.16× [1.15, 1.17] | +13.6% | [+12.9%, +14.4%] | 1.0 pts | 0.8 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 475 | 589 | 1.21× [1.18, 1.24] | +17.4% | [+15.4%, +19.4%] | 2.8 pts | 2.3 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 8 | 55.6 | 70.5 | 1.27× [1.26, 1.28] | +21.1% | [+20.5%, +21.7%] | 0.7 pts | 0.7 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 8 | 2308 | 1865 | 0.80× [0.78, 0.81] | -25.2% | [-27.6%, -22.8%] | 2.9 pts | 1.1 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 8 | 249 | 208 | 0.84× [0.82, 0.85] | -19.7% | [-21.4%, -17.9%] | 2.1 pts | 0.9 | yes |
| multi | street | 4096 | churn | ordered | baseline | 8 | 94.8 | 131 | 1.36× [1.34, 1.39] | +26.7% | [+25.5%, +27.9%] | 1.5 pts | 1.3 | yes |
| multi | street | 4096 | build | ordered | baseline | 8 | 2.84 ms | 5.60 ms | 1.98× [1.95, 2.00] | +49.4% | [+48.8%, +50.0%] | 0.7 pts | 1.1 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 8 | 37.1 | 44.4 | 1.20× [1.19, 1.22] | +16.9% | [+15.9%, +17.9%] | 1.2 pts | 1.1 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2651 | 2683 | 1.01× [1.00, 1.03] | +1.5% | [-0.2%, +3.1%] | 2.0 pts | 0.8 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 8 | 46.1 | 58.3 | 1.27× [1.26, 1.28] | +21.1% | [+20.3%, +21.8%] | 0.9 pts | 0.6 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 8 | 4.06 ms | 5.76 ms | 1.42× [1.41, 1.43] | +29.6% | [+29.1%, +30.0%] | 0.5 pts | 0.8 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 8 | 173 | 246 | 1.43× [1.41, 1.44] | +29.9% | [+29.2%, +30.5%] | 0.8 pts | 1.2 | yes |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 9007 | 10.5 µs | 1.17× [1.16, 1.18] | +14.3% | [+13.6%, +15.0%] | 0.8 pts | 0.8 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 8 | 303 | 344 | 1.13× [1.11, 1.16] | +11.9% | [+10.1%, +13.7%] | 2.2 pts | 1.6 | yes |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 10 | 55.5 | 71.9 | 1.29× [1.28, 1.30] | +22.6% | [+21.8%, +23.3%] | 1.0 pts | 1.6 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 10 | 3312 | 3362 | 1.01× [1.01, 1.02] | +1.3% | [+0.6%, +2.0%] | 1.0 pts | 0.4 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 10 | 89.1 | 98.5 | 1.11× [1.08, 1.14] | +9.8% | [+7.6%, +12.0%] | 3.1 pts | 4.3 | no |
| multi | uuid | 4096 | churn | ordered | baseline | 10 | 78.4 | 111 | 1.43× [1.42, 1.45] | +30.3% | [+29.7%, +30.8%] | 0.8 pts | 0.4 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 10 | 6.10 ms | 10.08 ms | 1.63× [1.61, 1.66] | +38.8% | [+37.8%, +39.8%] | 1.4 pts | 1.5 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 326 | 469 | 1.44× [1.42, 1.47] | +30.7% | [+29.4%, +32.0%] | 1.8 pts | 3.0 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 11.2 µs | 12.5 µs | 1.12× [1.11, 1.13] | +10.9% | [+10.1%, +11.6%] | 1.0 pts | 1.0 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 746 | 813 | 1.09× [1.08, 1.10] | +8.3% | [+7.6%, +8.9%] | 0.9 pts | 0.9 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 487 | 585 | 1.19× [1.16, 1.22] | +16.1% | [+14.1%, +18.0%] | 2.7 pts | 2.4 | yes |
| unique | path | 4096 | valuesFor | ordered | baseline | 10 | 86.4 | 102 | 1.19× [1.17, 1.20] | +15.7% | [+14.9%, +16.6%] | 1.2 pts | 1.3 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 10 | 2111 | 692 | 0.33× [0.32, 0.35] | -199.4% | [-211.5%, -187.3%] | 16.9 pts | 3.6 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 10 | 252 | 235 | 0.97× [0.92, 1.04] | -2.7% | [-8.9%, +3.4%] | 8.6 pts | 3.9 | no |
| unique | path | 4096 | churn | ordered | baseline | 10 | 204 | 222 | 1.10× [1.09, 1.11] | +8.8% | [+8.0%, +9.6%] | 1.1 pts | 0.5 | yes |
| unique | path | 4096 | build | ordered | baseline | 10 | 2.16 ms | 3.05 ms | 1.40× [1.40, 1.41] | +28.8% | [+28.3%, +29.2%] | 0.6 pts | 0.7 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 519 | 448 | 0.86× [0.84, 0.88] | -16.2% | [-19.1%, -13.3%] | 4.1 pts | 3.3 | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 7110 | 2745 | 0.38× [0.37, 0.38] | -163.5% | [-167.0%, -160.1%] | 4.8 pts | 1.0 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 6303 | 1951 | 0.33× [0.31, 0.35] | -203.3% | [-217.9%, -188.7%] | 20.4 pts | 0.6 | yes |
| unique | path | 262144 | churn | ordered | baseline | 10 | 928 | 903 | 0.99× [0.96, 1.01] | -1.5% | [-3.8%, +0.8%] | 3.2 pts | 0.8 | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 6 | 37.7 | 45.8 | 1.22× [1.21, 1.22] | +17.7% | [+17.1%, +18.3%] | 0.6 pts | 0.5 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 6 | 1582 | 391 | 0.25× [0.25, 0.25] | -306.0% | [-308.0%, -304.0%] | 1.9 pts | 0.4 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 6 | 3103 | 619 | 0.20× [0.20, 0.20] | -401.3% | [-405.4%, -397.2%] | 3.9 pts | 0.9 | yes |
| unique | str | 4096 | churn | ordered | baseline | 6 | 87.1 | 112 | 1.27× [1.26, 1.29] | +21.6% | [+20.4%, +22.7%] | 1.1 pts | 1.0 | yes |
| unique | str | 4096 | build | ordered | baseline | 6 | 1.05 ms | 1.49 ms | 1.42× [1.39, 1.45] | +29.6% | [+28.0%, +31.1%] | 1.5 pts | 1.1 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 283 | 170 | 0.62× [0.58, 0.66] | -61.9% | [-71.8%, -52.0%] | 13.8 pts | 4.9 | no |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 5347 | 1268 | 0.24× [0.24, 0.24] | -317.5% | [-321.7%, -313.3%] | 5.9 pts | 0.5 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 781.6 µs | 134.1 µs | 0.18× [0.17, 0.18] | -469.2% | [-480.6%, -457.9%] | 15.9 pts | 0.5 | yes |
| unique | str | 262144 | churn | ordered | baseline | 10 | 355 | 333 | 0.93× [0.90, 0.96] | -7.2% | [-10.5%, -3.8%] | 4.7 pts | 1.9 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 48.3 | 54.3 | 1.12× [1.11, 1.13] | +10.7% | [+10.2%, +11.3%] | 0.5 pts | 0.5 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 1797 | 481 | 0.27× [0.26, 0.27] | -273.6% | [-277.4%, -269.8%] | 3.7 pts | 0.8 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 212 | 110 | 0.52× [0.51, 0.52] | -93.0% | [-94.7%, -91.3%] | 1.6 pts | 0.6 | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 111 | 121 | 1.09× [1.08, 1.10] | +8.3% | [+7.3%, +9.3%] | 1.0 pts | 0.9 | yes |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.28 ms | 1.91 ms | 1.49× [1.48, 1.51] | +33.0% | [+32.4%, +33.7%] | 0.6 pts | 0.7 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 15.4 | 17.1 | 1.11× [1.09, 1.13] | +9.7% | [+7.9%, +11.5%] | 1.7 pts | 0.6 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1012 | 243 | 0.24× [0.24, 0.24] | -316.9% | [-320.1%, -313.7%] | 3.1 pts | 1.4 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 44.1 | 46.4 | 1.05× [1.04, 1.07] | +5.1% | [+3.8%, +6.3%] | 1.2 pts | 1.3 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 609.1 µs | 629.6 µs | 1.03× [1.02, 1.04] | +3.2% | [+2.3%, +4.1%] | 0.9 pts | 0.7 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 10 | 97.9 | 51.1 | 0.52× [0.50, 0.54] | -92.6% | [-98.4%, -86.7%] | 8.2 pts | 1.9 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 3424 | 462 | 0.13× [0.13, 0.14] | -640.8% | [-653.5%, -628.1%] | 17.7 pts | 0.7 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 10 | 185 | 125 | 0.67× [0.65, 0.69] | -50.2% | [-54.5%, -45.8%] | 6.1 pts | 1.1 | yes |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 6 | 30.7 | 45.4 | 1.47× [1.42, 1.52] | +31.9% | [+29.5%, +34.3%] | 2.3 pts | 2.9 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1460 | 413 | 0.28× [0.28, 0.29] | -254.0% | [-257.2%, -250.8%] | 3.1 pts | 0.9 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 67.4 | 75.0 | 1.12× [1.11, 1.13] | +10.5% | [+9.9%, +11.2%] | 0.6 pts | 1.1 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 6 | 82.9 | 99.6 | 1.20× [1.19, 1.21] | +16.7% | [+15.8%, +17.5%] | 0.8 pts | 0.6 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 6 | 970.1 µs | 1.45 ms | 1.49× [1.47, 1.50] | +32.8% | [+32.1%, +33.5%] | 0.7 pts | 0.9 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 259 | 279 | 1.04× [1.00, 1.09] | +3.8% | [-0.5%, +8.1%] | 6.0 pts | 5.5 | no |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 5459 | 1471 | 0.27× [0.27, 0.27] | -270.0% | [-275.6%, -264.5%] | 7.7 pts | 1.6 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 460 | 253 | 0.57× [0.54, 0.59] | -76.8% | [-83.5%, -70.0%] | 9.5 pts | 3.9 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 345 | 368 | 1.04× [0.99, 1.09] | +3.5% | [-1.1%, +8.0%] | 6.3 pts | 0.5 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi path n=4096 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of 6.12% does not clear the 11.06% median noise floor of the processes
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of 0.81% does not clear the 3.89% median noise floor of the processes
- multi str n=4096 prefix: ordered vs baseline: the pooled interval [-1.17%, 2.78%] includes zero
- multi str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.19%, 3.14%] includes zero
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.29% does not clear the 2.26% median noise floor of the processes
- multi uuid n=4096 prefix: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=4096 valuesBetween: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=4096 prefix: ordered vs baseline: the pooled interval [-8.87%, 3.43%] includes zero
- unique path n=4096 prefix: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 churn: ordered vs baseline: the pooled interval [-3.80%, 0.84%] includes zero
- unique str n=262144 valuesFor: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.48%, 8.11%] includes zero
- unique uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesFor: ordered vs baseline: 7 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=262144 prefix: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-1.06%, 8.00%] includes zero
