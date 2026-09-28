| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesBetween | ordered | baseline | 6 | 3079 | 3052 | 0.99× [0.98, 1.00] | -1.1% | [-2.6%, +0.3%] | 1.4 pts | 0.6 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 6 | 76.6 | 76.0 | 0.99× [0.98, 1.01] | -0.6% | [-1.8%, +0.5%] | 1.1 pts | 1.5 | yes |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3830 | 3627 | 0.94× [0.93, 0.95] | -6.2% | [-7.0%, -5.3%] | 0.8 pts | 0.7 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 92.1 | 91.6 | 0.99× [0.98, 1.00] | -0.9% | [-1.6%, -0.1%] | 0.7 pts | 0.7 | yes |
| multi | email | 262144 | valuesBetween | ordered | baseline | 6 | 8290 | 8034 | 0.97× [0.95, 0.99] | -3.0% | [-4.9%, -1.0%] | 1.9 pts | 1.7 | yes |
| multi | email | 262144 | prefix | ordered | baseline | 6 | 300 | 302 | 1.00× [0.98, 1.02] | +0.1% | [-1.8%, +2.0%] | 1.8 pts | 1.3 | yes |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 6 | 9880 | 9652 | 0.97× [0.97, 0.98] | -2.7% | [-3.5%, -1.9%] | 0.8 pts | 0.8 | yes |
| multi | email | 1048576 | prefix | ordered | baseline | 6 | 563 | 559 | 1.00× [0.99, 1.01] | -0.3% | [-1.3%, +0.6%] | 0.9 pts | 0.9 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4377 | 4272 | 0.97× [0.96, 0.99] | -3.0% | [-4.7%, -1.3%] | 1.6 pts | 0.7 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 334 | 335 | 1.01× [1.00, 1.01] | +0.7% | [+0.0%, +1.4%] | 0.6 pts | 0.2 | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 10 | 4960 | 4803 | 0.97× [0.97, 0.98] | -2.8% | [-3.4%, -2.2%] | 0.9 pts | 0.7 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 10 | 646 | 632 | 0.99× [0.96, 1.01] | -1.4% | [-3.9%, +1.2%] | 3.6 pts | 0.6 | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 13.0 µs | 12.4 µs | 0.96× [0.95, 0.97] | -3.7% | [-4.8%, -2.7%] | 1.5 pts | 1.2 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 12.6 µs | 12.0 µs | 0.95× [0.92, 0.98] | -5.4% | [-8.4%, -2.4%] | 4.2 pts | 0.3 | no |
| multi | str | 4096 | valuesBetween | ordered | baseline | 6 | 3702 | 3589 | 0.97× [0.96, 0.98] | -2.7% | [-3.9%, -1.6%] | 1.1 pts | 0.5 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 6 | 8370 | 8106 | 0.97× [0.95, 0.98] | -3.4% | [-5.0%, -1.9%] | 1.5 pts | 0.5 | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3849 | 3725 | 0.97× [0.96, 0.99] | -2.7% | [-4.0%, -1.5%] | 1.2 pts | 0.9 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 36.2 µs | 34.9 µs | 0.97× [0.96, 0.98] | -3.1% | [-4.1%, -2.0%] | 1.0 pts | 0.7 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 6 | 9033 | 8642 | 0.97× [0.95, 0.99] | -2.9% | [-4.8%, -1.0%] | 1.8 pts | 1.4 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 6 | 1.44 ms | 1.40 ms | 0.97× [0.96, 0.99] | -2.6% | [-4.4%, -0.9%] | 1.7 pts | 0.9 | yes |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 6 | 10.1 µs | 9773 | 0.96× [0.96, 0.97] | -3.9% | [-4.2%, -3.6%] | 0.3 pts | 0.3 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 6 | 6.33 ms | 6.11 ms | 0.96× [0.96, 0.97] | -3.6% | [-4.0%, -3.3%] | 0.3 pts | 0.6 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 8 | 2719 | 2456 | 0.90× [0.89, 0.91] | -11.4% | [-12.9%, -9.9%] | 1.8 pts | 0.8 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 8 | 279 | 267 | 0.95× [0.94, 0.97] | -4.8% | [-6.4%, -3.2%] | 1.9 pts | 0.5 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 8 | 3235 | 2957 | 0.91× [0.91, 0.92] | -9.7% | [-10.2%, -9.3%] | 0.5 pts | 0.4 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 8 | 1070 | 965 | 0.90× [0.89, 0.91] | -11.2% | [-13.0%, -9.4%] | 2.2 pts | 0.7 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 2812 | 2793 | 1.00× [0.98, 1.01] | -0.2% | [-1.6%, +1.2%] | 1.3 pts | 0.7 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3977 | 3627 | 0.91× [0.90, 0.92] | -9.8% | [-10.6%, -9.0%] | 0.8 pts | 0.7 | yes |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 6 | 7245 | 7211 | 1.00× [0.98, 1.02] | -0.2% | [-2.0%, +1.6%] | 1.7 pts | 1.5 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 6 | 6726 | 6701 | 0.99× [0.98, 1.00] | -0.7% | [-1.9%, +0.5%] | 1.1 pts | 1.0 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 6 | 4143 | 3992 | 0.97× [0.96, 0.98] | -3.4% | [-4.6%, -2.1%] | 1.2 pts | 0.6 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 6 | 186 | 188 | 1.02× [1.00, 1.03] | +1.5% | [+0.1%, +2.9%] | 1.3 pts | 0.9 | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4708 | 4529 | 0.96× [0.95, 0.96] | -4.5% | [-5.3%, -3.7%] | 0.8 pts | 0.6 | yes |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 275 | 274 | 1.00× [0.99, 1.01] | -0.2% | [-1.2%, +0.8%] | 0.9 pts | 0.5 | yes |
| multi | url | 262144 | valuesBetween | ordered | baseline | 6 | 12.2 µs | 11.8 µs | 0.97× [0.96, 0.98] | -3.2% | [-4.3%, -2.1%] | 1.1 pts | 0.8 | yes |
| multi | url | 262144 | prefix | ordered | baseline | 6 | 2265 | 2275 | 0.99× [0.97, 1.01] | -1.1% | [-3.1%, +0.9%] | 1.9 pts | 0.3 | yes |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 10 | 13.5 µs | 13.1 µs | 0.97× [0.96, 0.98] | -3.1% | [-3.7%, -2.5%] | 0.8 pts | 0.8 | yes |
| multi | url | 1048576 | prefix | ordered | baseline | 10 | 9641 | 9399 | 0.97× [0.95, 1.00] | -2.9% | [-5.5%, -0.3%] | 3.7 pts | 0.3 | no |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3485 | 3395 | 0.97× [0.96, 0.99] | -2.6% | [-4.0%, -1.2%] | 1.3 pts | 0.6 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 88.8 | 89.8 | 1.01× [1.01, 1.01] | +1.0% | [+0.7%, +1.3%] | 0.3 pts | 0.4 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3735 | 3686 | 0.99× [0.98, 0.99] | -1.5% | [-1.9%, -1.1%] | 0.4 pts | 0.4 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 110 | 111 | 1.01× [0.99, 1.02] | +0.8% | [-0.7%, +2.3%] | 1.4 pts | 1.3 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 6 | 8991 | 8789 | 0.98× [0.96, 0.99] | -2.3% | [-3.6%, -1.0%] | 1.2 pts | 1.0 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 6 | 593 | 592 | 1.00× [0.99, 1.01] | -0.1% | [-1.0%, +0.7%] | 0.8 pts | 0.5 | yes |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 6 | 10.6 µs | 10.2 µs | 0.97× [0.96, 0.98] | -3.1% | [-4.6%, -1.7%] | 1.4 pts | 1.2 | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 6 | 1925 | 1886 | 0.98× [0.96, 1.00] | -1.8% | [-3.7%, +0.0%] | 1.8 pts | 1.6 | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 6 | 1422 | 1321 | 0.93× [0.92, 0.94] | -7.6% | [-8.9%, -6.2%] | 1.3 pts | 0.8 | yes |
| unique | email | 4096 | prefix | ordered | baseline | 6 | 56.4 | 55.9 | 0.99× [0.99, 0.99] | -0.9% | [-1.2%, -0.6%] | 0.3 pts | 0.5 | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 1943 | 1674 | 0.86× [0.85, 0.86] | -16.5% | [-17.4%, -15.7%] | 0.8 pts | 0.6 | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 69.8 | 69.1 | 0.99× [0.98, 1.00] | -1.1% | [-1.9%, -0.4%] | 0.7 pts | 0.7 | yes |
| unique | email | 262144 | valuesBetween | ordered | baseline | 6 | 4136 | 3834 | 0.92× [0.91, 0.94] | -8.4% | [-9.9%, -6.8%] | 1.5 pts | 0.9 | yes |
| unique | email | 262144 | prefix | ordered | baseline | 6 | 225 | 230 | 1.01× [0.99, 1.02] | +0.5% | [-1.0%, +2.0%] | 1.5 pts | 0.9 | yes |
| unique | email | 1048576 | valuesBetween | ordered | baseline | 6 | 6751 | 6331 | 0.94× [0.93, 0.95] | -6.4% | [-7.8%, -5.0%] | 1.4 pts | 1.3 | yes |
| unique | email | 1048576 | prefix | ordered | baseline | 6 | 420 | 416 | 0.98× [0.96, 1.00] | -2.0% | [-4.0%, -0.1%] | 1.9 pts | 1.5 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 6 | 2424 | 2252 | 0.93× [0.92, 0.94] | -7.5% | [-9.0%, -5.9%] | 1.5 pts | 0.8 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 6 | 248 | 254 | 1.02× [1.00, 1.03] | +1.8% | [+0.4%, +3.2%] | 1.3 pts | 0.8 | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 10 | 2907 | 2712 | 0.94× [0.93, 0.94] | -6.9% | [-7.6%, -6.3%] | 0.9 pts | 0.8 | yes |
| unique | path | 16384 | prefix | ordered | baseline | 10 | 446 | 435 | 0.99× [0.97, 1.01] | -1.2% | [-3.6%, +1.1%] | 3.3 pts | 0.7 | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 8630 | 8247 | 0.95× [0.93, 0.97] | -5.0% | [-7.0%, -3.0%] | 2.7 pts | 2.5 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 6448 | 6048 | 0.92× [0.90, 0.95] | -8.2% | [-11.5%, -4.9%] | 4.7 pts | 0.4 | no |
| unique | str | 4096 | valuesBetween | ordered | baseline | 10 | 1864 | 1729 | 0.92× [0.92, 0.93] | -8.2% | [-9.2%, -7.3%] | 1.3 pts | 0.7 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 10 | 3819 | 3598 | 0.93× [0.91, 0.95] | -7.2% | [-9.6%, -4.8%] | 3.4 pts | 1.0 | no |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 1894 | 1737 | 0.91× [0.91, 0.92] | -9.3% | [-9.9%, -8.7%] | 0.6 pts | 0.5 | yes |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 16.4 µs | 14.6 µs | 0.89× [0.88, 0.91] | -12.1% | [-13.8%, -10.3%] | 1.7 pts | 0.9 | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 4650 | 4305 | 0.93× [0.92, 0.93] | -8.0% | [-8.9%, -7.1%] | 1.2 pts | 0.6 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 756.5 µs | 696.1 µs | 0.93× [0.91, 0.96] | -7.1% | [-9.6%, -4.6%] | 3.5 pts | 0.4 | no |
| unique | str | 1048576 | valuesBetween | ordered | baseline | 6 | 7097 | 6728 | 0.95× [0.94, 0.97] | -4.7% | [-6.2%, -3.2%] | 1.4 pts | 1.1 | yes |
| unique | str | 1048576 | prefix | ordered | baseline | 6 | 4.25 ms | 4.09 ms | 0.95× [0.94, 0.96] | -5.1% | [-6.3%, -3.9%] | 1.1 pts | 1.6 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 2194 | 1919 | 0.87× [0.87, 0.88] | -14.6% | [-15.4%, -13.8%] | 0.8 pts | 0.5 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 241 | 224 | 0.93× [0.92, 0.95] | -7.3% | [-8.9%, -5.7%] | 1.5 pts | 0.6 | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 10 | 2583 | 2285 | 0.89× [0.88, 0.90] | -12.5% | [-13.6%, -11.4%] | 1.5 pts | 1.1 | yes |
| unique | street | 16384 | prefix | ordered | baseline | 10 | 876 | 778 | 0.90× [0.88, 0.92] | -11.6% | [-14.2%, -9.0%] | 3.7 pts | 1.0 | no |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1111 | 1059 | 0.95× [0.95, 0.96] | -4.9% | [-5.8%, -4.0%] | 0.9 pts | 0.7 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 2081 | 1638 | 0.79× [0.79, 0.80] | -26.1% | [-27.3%, -24.9%] | 1.2 pts | 0.8 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 2189 | 2102 | 0.97× [0.95, 0.98] | -3.5% | [-5.5%, -1.6%] | 2.7 pts | 1.5 | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 6 | 3304 | 3163 | 0.95× [0.94, 0.96] | -5.0% | [-6.0%, -4.0%] | 0.9 pts | 1.0 | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 6 | 2259 | 2076 | 0.92× [0.91, 0.93] | -8.8% | [-10.1%, -7.5%] | 1.2 pts | 0.6 | yes |
| unique | url | 4096 | prefix | ordered | baseline | 6 | 153 | 156 | 1.01× [1.01, 1.02] | +1.4% | [+0.6%, +2.3%] | 0.8 pts | 0.6 | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 2706 | 2495 | 0.92× [0.92, 0.92] | -8.4% | [-8.7%, -8.1%] | 0.3 pts | 0.2 | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 219 | 220 | 1.00× [0.99, 1.01] | +0.0% | [-1.0%, +1.0%] | 1.0 pts | 0.6 | yes |
| unique | url | 262144 | valuesBetween | ordered | baseline | 10 | 8338 | 7902 | 0.95× [0.94, 0.96] | -5.1% | [-6.4%, -3.8%] | 1.8 pts | 1.4 | yes |
| unique | url | 262144 | prefix | ordered | baseline | 10 | 1441 | 1413 | 0.97× [0.93, 1.00] | -3.4% | [-7.2%, +0.3%] | 5.3 pts | 1.2 | no |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 10 | 9899 | 9633 | 0.97× [0.96, 0.98] | -3.6% | [-4.7%, -2.4%] | 1.6 pts | 1.7 | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 10 | 5478 | 5186 | 0.95× [0.93, 0.98] | -4.9% | [-8.1%, -1.8%] | 4.4 pts | 0.4 | no |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1692 | 1577 | 0.93× [0.91, 0.94] | -8.0% | [-9.8%, -6.2%] | 1.7 pts | 1.1 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 66.9 | 67.3 | 1.00× [0.99, 1.01] | +0.2% | [-0.6%, +1.0%] | 0.8 pts | 0.9 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 1747 | 1646 | 0.94× [0.93, 0.95] | -6.2% | [-7.2%, -5.3%] | 0.9 pts | 0.9 | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 82.7 | 82.8 | 1.00× [0.99, 1.00] | -0.0% | [-0.5%, +0.4%] | 0.5 pts | 0.6 | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 4763 | 4486 | 0.94× [0.92, 0.96] | -6.3% | [-8.2%, -4.4%] | 2.3 pts | 1.0 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 8 | 382 | 376 | 0.99× [0.97, 1.00] | -1.2% | [-2.9%, +0.5%] | 2.0 pts | 1.4 | yes |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 6 | 7127 | 6762 | 0.94× [0.93, 0.96] | -5.9% | [-7.2%, -4.7%] | 1.2 pts | 1.1 | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 6 | 1288 | 1261 | 0.97× [0.96, 0.98] | -3.3% | [-4.2%, -2.3%] | 0.9 pts | 0.9 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.11% does not clear the 2.08% median noise floor of the processes
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.55%, 0.34%] includes zero
- multi email n=4096 prefix: ordered vs baseline: the pooled interval [-1.78%, 0.50%] includes zero
- multi email n=16384 prefix: ordered vs baseline: the pooled difference of -0.87% does not clear the 1.02% median noise floor of the processes
- multi email n=262144 prefix: ordered vs baseline: the pooled difference of 0.13% does not clear the 2.20% median noise floor of the processes
- multi email n=262144 prefix: ordered vs baseline: the pooled interval [-1.75%, 2.01%] includes zero
- multi email n=1048576 prefix: ordered vs baseline: the pooled difference of -0.35% does not clear the 1.47% median noise floor of the processes
- multi email n=1048576 prefix: ordered vs baseline: the pooled interval [-1.30%, 0.61%] includes zero
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of 0.71% does not clear the 2.40% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of -1.38% does not clear the 6.06% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-3.92%, 1.17%] includes zero
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of -5.41% does not clear the 12.78% median noise floor of the processes
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.23% does not clear the 2.21% median noise floor of the processes
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.61%, 1.16%] includes zero
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.19% does not clear the 1.17% median noise floor of the processes
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.98%, 1.60%] includes zero
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.71% does not clear the 1.28% median noise floor of the processes
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-1.89%, 0.48%] includes zero
- multi url n=16384 prefix: ordered vs baseline: the pooled difference of -0.18% does not clear the 1.43% median noise floor of the processes
- multi url n=16384 prefix: ordered vs baseline: the pooled interval [-1.15%, 0.79%] includes zero
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of -1.09% does not clear the 6.87% median noise floor of the processes
- multi url n=262144 prefix: ordered vs baseline: the pooled interval [-3.07%, 0.88%] includes zero
- multi url n=1048576 prefix: ordered vs baseline: the pooled difference of -2.92% does not clear the 9.78% median noise floor of the processes
- multi uuid n=4096 prefix: ordered vs baseline: the pooled difference of 1.00% does not clear the 1.12% median noise floor of the processes
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.51% does not clear the 1.86% median noise floor of the processes
- multi uuid n=16384 prefix: ordered vs baseline: the pooled difference of 0.80% does not clear the 1.13% median noise floor of the processes
- multi uuid n=16384 prefix: ordered vs baseline: the pooled interval [-0.68%, 2.28%] includes zero
- multi uuid n=262144 prefix: ordered vs baseline: the pooled difference of -0.15% does not clear the 1.50% median noise floor of the processes
- multi uuid n=262144 prefix: ordered vs baseline: the pooled interval [-0.99%, 0.69%] includes zero
- multi uuid n=1048576 prefix: ordered vs baseline: the pooled interval [-3.73%, 0.05%] includes zero
- unique email n=262144 prefix: ordered vs baseline: the pooled difference of 0.50% does not clear the 2.31% median noise floor of the processes
- unique email n=262144 prefix: ordered vs baseline: the pooled interval [-1.04%, 2.04%] includes zero
- unique path n=16384 prefix: ordered vs baseline: the pooled difference of -1.24% does not clear the 5.99% median noise floor of the processes
- unique path n=16384 prefix: ordered vs baseline: the pooled interval [-3.57%, 1.09%] includes zero
- unique path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 prefix: ordered vs baseline: the pooled difference of -8.20% does not clear the 14.14% median noise floor of the processes
- unique str n=262144 prefix: ordered vs baseline: the pooled difference of -7.10% does not clear the 13.53% median noise floor of the processes
- unique url n=16384 prefix: ordered vs baseline: the pooled difference of 0.01% does not clear the 1.88% median noise floor of the processes
- unique url n=16384 prefix: ordered vs baseline: the pooled interval [-1.01%, 1.02%] includes zero
- unique url n=262144 prefix: ordered vs baseline: the pooled difference of -3.43% does not clear the 5.28% median noise floor of the processes
- unique url n=262144 prefix: ordered vs baseline: the pooled interval [-7.18%, 0.33%] includes zero
- unique url n=1048576 prefix: ordered vs baseline: the pooled difference of -4.93% does not clear the 7.73% median noise floor of the processes
- unique uuid n=4096 prefix: ordered vs baseline: the pooled difference of 0.22% does not clear the 0.90% median noise floor of the processes
- unique uuid n=4096 prefix: ordered vs baseline: the pooled interval [-0.59%, 1.04%] includes zero
- unique uuid n=16384 prefix: ordered vs baseline: the pooled difference of -0.05% does not clear the 0.87% median noise floor of the processes
- unique uuid n=16384 prefix: ordered vs baseline: the pooled interval [-0.53%, 0.44%] includes zero
- unique uuid n=262144 prefix: ordered vs baseline: the pooled difference of -1.22% does not clear the 1.46% median noise floor of the processes
- unique uuid n=262144 prefix: ordered vs baseline: the pooled interval [-2.92%, 0.47%] includes zero
