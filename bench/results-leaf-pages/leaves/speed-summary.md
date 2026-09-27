| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 6 | 53.7 | 52.4 | 0.97× [0.97, 0.98] | -2.7% | [-3.0%, -2.2%] | 0.4 pts | 0.4 | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 6 | 2978 | 2970 | 1.00× [0.98, 1.01] | -0.3% | [-2.3%, +1.3%] | 1.7 pts | 1.0 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 6 | 77.0 | 75.9 | 0.99× [0.97, 1.00] | -1.4% | [-3.0%, -0.5%] | 1.2 pts | 1.3 | yes |
| multi | email | 4096 | churn | ordered | baseline | 6 | 66.6 | 63.4 | 0.95× [0.94, 0.96] | -5.6% | [-6.8%, -4.4%] | 1.2 pts | 0.5 | yes |
| multi | email | 4096 | build | ordered | baseline | 6 | 5.68 ms | 5.48 ms | 0.97× [0.96, 0.97] | -3.4% | [-4.6%, -2.6%] | 1.0 pts | 1.1 | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 69.5 | 64.7 | 0.94× [0.88, 0.97] | -6.7% | [-13.2%, -2.7%] | 5.0 pts | 1.5 | no |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3590 | 3548 | 0.99× [0.98, 0.99] | -1.4% | [-2.0%, -0.8%] | 0.6 pts | 0.4 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 95.1 | 93.5 | 0.97× [0.96, 0.99] | -2.9% | [-4.5%, -0.7%] | 1.8 pts | 1.7 | yes |
| multi | email | 16384 | churn | ordered | baseline | 6 | 102 | 87.6 | 0.85× [0.82, 0.90] | -18.2% | [-22.2%, -11.4%] | 5.2 pts | 1.3 | no |
| multi | email | 16384 | build | ordered | baseline | 6 | 30.09 ms | 27.78 ms | 0.92× [0.92, 0.93] | -8.5% | [-9.2%, -7.1%] | 1.0 pts | 1.2 | yes |
| multi | email | 65536 | valuesFor | ordered | baseline | 6 | 184 | 124 | 0.63× [0.58, 0.79] | -58.3% | [-71.2%, -26.5%] | 21.3 pts | 2.6 | no |
| multi | email | 65536 | valuesBetween | ordered | baseline | 6 | 6677 | 5213 | 0.82× [0.67, 0.89] | -21.5% | [-48.6%, -11.7%] | 17.5 pts | 3.1 | no |
| multi | email | 65536 | prefix | ordered | baseline | 6 | 231 | 161 | 0.70× [0.62, 0.80] | -42.4% | [-62.3%, -24.6%] | 18.0 pts | 2.7 | no |
| multi | email | 65536 | churn | ordered | baseline | 6 | 242 | 182 | 0.74× [0.71, 0.80] | -34.3% | [-40.3%, -25.4%] | 7.1 pts | 1.5 | no |
| multi | email | 65536 | build | ordered | baseline | 6 | 213.59 ms | 193.12 ms | 0.91× [0.89, 0.93] | -10.4% | [-11.9%, -7.8%] | 1.9 pts | 1.4 | no |
| multi | email | 262144 | valuesFor | ordered | baseline | 6 | 329 | 263 | 0.79× [0.75, 0.86] | -27.1% | [-34.2%, -16.3%] | 8.5 pts | 3.3 | no |
| multi | email | 262144 | valuesBetween | ordered | baseline | 6 | 10.5 µs | 9266 | 0.88× [0.86, 0.90] | -14.0% | [-15.7%, -11.4%] | 2.1 pts | 1.4 | no |
| multi | email | 262144 | prefix | ordered | baseline | 6 | 433 | 323 | 0.76× [0.73, 0.79] | -32.4% | [-37.8%, -26.0%] | 5.6 pts | 1.6 | no |
| multi | email | 262144 | churn | ordered | baseline | 6 | 418 | 377 | 0.88× [0.79, 1.01] | -13.1% | [-27.1%, +1.1%] | 13.5 pts | 5.0 | no |
| multi | email | 1048576 | valuesFor | ordered | baseline | 6 | 511 | 436 | 0.89× [0.84, 0.94] | -12.3% | [-18.8%, -6.9%] | 5.6 pts | 2.2 | no |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 6 | 12.5 µs | 11.4 µs | 0.91× [0.90, 0.93] | -9.9% | [-11.6%, -7.7%] | 1.9 pts | 0.9 | yes |
| multi | email | 1048576 | prefix | ordered | baseline | 6 | 738 | 649 | 0.88× [0.85, 0.90] | -13.8% | [-17.1%, -11.0%] | 2.9 pts | 1.3 | no |
| multi | email | 1048576 | churn | ordered | baseline | 6 | 587 | 557 | 0.96× [0.79, 1.06] | -4.6% | [-27.0%, +5.8%] | 15.6 pts | 7.2 | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 112 | 109 | 0.98× [0.97, 0.98] | -2.2% | [-3.0%, -1.8%] | 0.6 pts | 0.6 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4267 | 4193 | 0.98× [0.96, 1.01] | -2.0% | [-3.7%, +0.6%] | 2.0 pts | 1.1 | no |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 344 | 340 | 0.99× [0.97, 1.02] | -1.1% | [-2.8%, +1.5%] | 2.1 pts | 0.6 | no |
| multi | path | 4096 | churn | ordered | baseline | 6 | 162 | 163 | 0.98× [0.97, 1.01] | -2.0% | [-3.0%, +0.5%] | 1.7 pts | 0.8 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 12.11 ms | 12.16 ms | 1.01× [1.00, 1.01] | +0.6% | [+0.3%, +1.0%] | 0.3 pts | 0.8 | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 6 | 160 | 139 | 0.87× [0.81, 0.97] | -15.5% | [-22.9%, -2.6%] | 9.7 pts | 2.2 | no |
| multi | path | 16384 | valuesBetween | ordered | baseline | 6 | 4842 | 4765 | 0.98× [0.98, 0.99] | -1.7% | [-2.5%, -0.8%] | 0.8 pts | 0.4 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 6 | 672 | 646 | 0.97× [0.92, 1.01] | -3.6% | [-9.0%, +0.8%] | 4.7 pts | 0.8 | no |
| multi | path | 16384 | churn | ordered | baseline | 6 | 275 | 248 | 0.90× [0.85, 0.95] | -11.2% | [-17.0%, -5.5%] | 5.5 pts | 1.4 | no |
| multi | path | 16384 | build | ordered | baseline | 6 | 66.83 ms | 67.07 ms | 1.00× [1.00, 1.01] | -0.1% | [-0.4%, +0.6%] | 0.5 pts | 0.8 | yes |
| multi | path | 65536 | valuesFor | ordered | baseline | 6 | 368 | 282 | 0.79× [0.72, 0.87] | -26.4% | [-39.5%, -14.5%] | 11.9 pts | 3.9 | no |
| multi | path | 65536 | valuesBetween | ordered | baseline | 6 | 9569 | 7510 | 0.85× [0.78, 0.93] | -17.1% | [-28.0%, -7.9%] | 9.6 pts | 1.8 | no |
| multi | path | 65536 | prefix | ordered | baseline | 6 | 2422 | 2272 | 0.92× [0.87, 1.01] | -8.5% | [-14.8%, +0.5%] | 7.3 pts | 0.7 | no |
| multi | path | 65536 | churn | ordered | baseline | 6 | 539 | 448 | 0.87× [0.77, 1.01] | -15.5% | [-30.5%, +0.7%] | 14.9 pts | 4.1 | no |
| multi | path | 65536 | build | ordered | baseline | 6 | 477.33 ms | 460.47 ms | 0.98× [0.96, 0.99] | -2.5% | [-4.6%, -0.6%] | 1.9 pts | 1.8 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 6 | 637 | 532 | 0.85× [0.81, 0.93] | -18.1% | [-23.5%, -7.5%] | 7.6 pts | 5.0 | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 6 | 13.7 µs | 12.3 µs | 0.90× [0.89, 0.91] | -10.8% | [-11.8%, -9.8%] | 1.0 pts | 0.6 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 6 | 12.0 µs | 11.0 µs | 0.93× [0.90, 0.97] | -7.5% | [-11.6%, -3.5%] | 3.9 pts | 0.2 | no |
| multi | path | 262144 | churn | ordered | baseline | 6 | 867 | 800 | 0.98× [0.85, 1.05] | -1.9% | [-17.5%, +4.9%] | 10.7 pts | 5.0 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 6 | 66.8 | 64.4 | 0.96× [0.96, 0.97] | -3.9% | [-4.5%, -3.2%] | 0.6 pts | 0.6 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 6 | 3552 | 3571 | 1.00× [0.99, 1.02] | -0.1% | [-1.4%, +1.8%] | 1.5 pts | 0.7 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 6 | 7976 | 7975 | 0.99× [0.97, 1.02] | -0.9% | [-3.2%, +2.0%] | 2.5 pts | 0.8 | no |
| multi | str | 4096 | churn | ordered | baseline | 6 | 81.6 | 78.4 | 0.96× [0.94, 0.97] | -4.5% | [-6.0%, -2.9%] | 1.5 pts | 1.1 | yes |
| multi | str | 4096 | build | ordered | baseline | 6 | 6.80 ms | 6.58 ms | 0.97× [0.96, 0.97] | -3.4% | [-3.7%, -2.9%] | 0.4 pts | 0.6 | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 86.4 | 77.6 | 0.90× [0.84, 0.98] | -11.0% | [-19.0%, -2.3%] | 8.0 pts | 3.2 | no |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3772 | 3728 | 0.99× [0.97, 1.00] | -1.4% | [-2.7%, -0.2%] | 1.2 pts | 1.2 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 35.7 µs | 35.1 µs | 0.98× [0.97, 0.99] | -1.8% | [-3.1%, -0.6%] | 1.2 pts | 0.8 | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 123 | 107 | 0.88× [0.83, 0.91] | -14.0% | [-20.3%, -9.9%] | 4.9 pts | 1.4 | no |
| multi | str | 16384 | build | ordered | baseline | 6 | 35.50 ms | 32.83 ms | 0.93× [0.92, 0.93] | -7.5% | [-8.6%, -7.1%] | 0.7 pts | 1.3 | yes |
| multi | str | 65536 | valuesFor | ordered | baseline | 6 | 223 | 134 | 0.61× [0.56, 0.67] | -64.3% | [-79.1%, -50.0%] | 13.8 pts | 1.7 | no |
| multi | str | 65536 | valuesBetween | ordered | baseline | 6 | 6313 | 5425 | 0.84× [0.78, 0.87] | -19.6% | [-28.6%, -14.4%] | 6.8 pts | 1.1 | no |
| multi | str | 65536 | prefix | ordered | baseline | 6 | 281.7 µs | 210.2 µs | 0.77× [0.68, 0.83] | -29.9% | [-47.3%, -21.1%] | 12.5 pts | 1.5 | no |
| multi | str | 65536 | churn | ordered | baseline | 6 | 288 | 214 | 0.80× [0.73, 0.91] | -25.2% | [-37.5%, -10.3%] | 13.0 pts | 2.9 | no |
| multi | str | 65536 | build | ordered | baseline | 6 | 236.69 ms | 217.43 ms | 0.92× [0.89, 0.96] | -8.9% | [-12.1%, -4.6%] | 3.6 pts | 2.4 | no |
| multi | str | 262144 | valuesFor | ordered | baseline | 6 | 360 | 299 | 0.85× [0.80, 0.87] | -18.2% | [-24.9%, -14.5%] | 4.9 pts | 1.5 | no |
| multi | str | 262144 | valuesBetween | ordered | baseline | 6 | 10.9 µs | 9923 | 0.91× [0.90, 0.93] | -9.9% | [-10.8%, -8.1%] | 1.3 pts | 1.1 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 6 | 1.75 ms | 1.61 ms | 0.92× [0.88, 0.94] | -8.6% | [-13.2%, -5.8%] | 3.5 pts | 1.6 | no |
| multi | str | 262144 | churn | ordered | baseline | 6 | 436 | 428 | 0.92× [0.82, 1.13] | -8.6% | [-21.4%, +11.3%] | 15.6 pts | 4.6 | no |
| multi | str | 1048576 | valuesFor | ordered | baseline | 6 | 500 | 435 | 0.87× [0.80, 0.92] | -14.6% | [-24.4%, -8.4%] | 7.6 pts | 3.1 | no |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 6 | 12.7 µs | 11.6 µs | 0.92× [0.90, 0.93] | -8.2% | [-11.0%, -7.2%] | 1.8 pts | 1.1 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 6 | 7.44 ms | 7.22 ms | 0.97× [0.97, 0.97] | -3.2% | [-3.5%, -2.9%] | 0.3 pts | 0.6 | yes |
| multi | str | 1048576 | churn | ordered | baseline | 6 | 613 | 586 | 0.97× [0.84, 1.11] | -3.4% | [-19.4%, +10.3%] | 14.2 pts | 7.1 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 6 | 58.2 | 57.8 | 0.99× [0.98, 1.01] | -1.0% | [-2.3%, +1.0%] | 1.6 pts | 1.3 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 6 | 2373 | 2411 | 1.01× [1.00, 1.02] | +1.4% | [+0.4%, +2.3%] | 0.9 pts | 0.5 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 6 | 258 | 259 | 1.00× [0.99, 1.02] | +0.2% | [-1.5%, +2.2%] | 1.7 pts | 0.7 | yes |
| multi | street | 4096 | churn | ordered | baseline | 6 | 94.2 | 91.5 | 0.97× [0.96, 0.98] | -2.6% | [-3.8%, -1.7%] | 1.0 pts | 0.8 | yes |
| multi | street | 4096 | build | ordered | baseline | 6 | 3.13 ms | 3.08 ms | 0.98× [0.97, 0.99] | -2.2% | [-2.6%, -1.3%] | 0.6 pts | 0.7 | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 81.9 | 80.3 | 0.98× [0.95, 1.00] | -1.9% | [-5.1%, +0.2%] | 2.5 pts | 1.6 | no |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2862 | 2888 | 1.00× [0.99, 1.02] | +0.1% | [-1.3%, +1.9%] | 1.5 pts | 1.2 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 947 | 947 | 1.00× [0.97, 1.03] | -0.2% | [-3.6%, +2.7%] | 3.0 pts | 1.0 | no |
| multi | street | 16384 | churn | ordered | baseline | 6 | 136 | 129 | 0.96× [0.93, 0.98] | -4.5% | [-7.1%, -2.0%] | 2.5 pts | 0.7 | no |
| multi | street | 16384 | build | ordered | baseline | 6 | 16.81 ms | 16.48 ms | 0.98× [0.98, 0.98] | -2.1% | [-2.5%, -1.6%] | 0.4 pts | 1.0 | yes |
| multi | street | 65536 | valuesFor | ordered | baseline | 6 | 162 | 115 | 0.73× [0.67, 0.80] | -37.5% | [-49.1%, -25.0%] | 11.5 pts | 1.4 | no |
| multi | street | 65536 | valuesBetween | ordered | baseline | 6 | 4105 | 3464 | 0.85× [0.78, 0.88] | -18.2% | [-28.7%, -13.6%] | 7.2 pts | 1.3 | no |
| multi | street | 65536 | prefix | ordered | baseline | 6 | 5000 | 4586 | 0.91× [0.86, 0.95] | -10.2% | [-16.2%, -4.9%] | 5.4 pts | 1.0 | no |
| multi | street | 65536 | churn | ordered | baseline | 6 | 289 | 245 | 0.82× [0.76, 0.92] | -22.6% | [-31.4%, -8.2%] | 11.0 pts | 2.1 | no |
| multi | street | 65536 | build | ordered | baseline | 6 | 115.75 ms | 110.63 ms | 0.95× [0.95, 0.96] | -4.8% | [-5.4%, -4.3%] | 0.5 pts | 0.6 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 5 | 38.5 | 39.7 | 1.03× [1.02, 1.04] | +2.8% | [+2.0%, +3.5%] | 0.6 pts | 0.6 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 5 | 2664 | 2706 | 1.02× [1.00, 1.03] | +1.6% | [+0.3%, +2.9%] | 1.0 pts | 0.6 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 5 | 43.9 | 43.3 | 0.99× [0.97, 1.01] | -0.7% | [-2.9%, +0.6%] | 1.4 pts | 0.8 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 5 | 4.14 ms | 4.09 ms | 1.00× [0.99, 1.00] | -0.5% | [-1.5%, +0.5%] | 0.8 pts | 1.0 | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 47.6 | 47.7 | 1.00× [0.99, 1.02] | +0.3% | [-0.7%, +1.5%] | 1.0 pts | 1.0 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3596 | 3595 | 1.00× [0.99, 1.01] | -0.3% | [-0.7%, +0.6%] | 0.6 pts | 0.4 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 55.6 | 53.1 | 0.93× [0.90, 0.98] | -7.5% | [-11.5%, -2.0%] | 4.5 pts | 0.9 | no |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 18.98 ms | 18.92 ms | 1.00× [0.99, 1.00] | -0.3% | [-0.7%, -0.1%] | 0.3 pts | 0.6 | yes |
| multi | u64 | 65536 | valuesFor | ordered | baseline | 6 | 79.5 | 69.7 | 0.96× [0.79, 1.07] | -3.8% | [-26.4%, +6.2%] | 15.6 pts | 2.3 | no |
| multi | u64 | 65536 | valuesBetween | ordered | baseline | 6 | 5441 | 4679 | 0.87× [0.81, 1.00] | -14.4% | [-23.7%, -0.3%] | 11.2 pts | 1.7 | no |
| multi | u64 | 65536 | churn | ordered | baseline | 6 | 132 | 116 | 0.86× [0.78, 0.92] | -16.4% | [-27.5%, -8.5%] | 9.0 pts | 1.3 | no |
| multi | u64 | 65536 | build | ordered | baseline | 6 | 118.86 ms | 117.36 ms | 0.99× [0.97, 1.02] | -0.9% | [-3.1%, +2.0%] | 2.5 pts | 3.1 | no |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 6 | 171 | 140 | 0.81× [0.76, 0.88] | -22.8% | [-32.2%, -14.0%] | 8.7 pts | 2.8 | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 6 | 8208 | 7747 | 0.95× [0.90, 0.98] | -5.8% | [-10.7%, -2.3%] | 4.0 pts | 2.3 | no |
| multi | u64 | 262144 | churn | ordered | baseline | 6 | 271 | 247 | 0.96× [0.87, 1.05] | -4.5% | [-15.0%, +4.9%] | 9.5 pts | 2.8 | no |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 6 | 212 | 197 | 0.93× [0.92, 0.94] | -7.4% | [-9.1%, -5.8%] | 1.6 pts | 1.2 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 6 | 8097 | 7816 | 0.97× [0.95, 0.98] | -3.5% | [-5.6%, -2.0%] | 1.7 pts | 1.4 | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 6 | 427 | 387 | 0.94× [0.77, 1.14] | -6.0% | [-29.9%, +12.0%] | 20.0 pts | 7.2 | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 6 | 87.0 | 86.1 | 0.99× [0.97, 1.00] | -0.8% | [-2.7%, +0.2%] | 1.3 pts | 1.6 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 6 | 3770 | 3758 | 0.99× [0.98, 1.01] | -0.6% | [-2.4%, +1.1%] | 1.6 pts | 1.2 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 6 | 422 | 419 | 0.99× [0.98, 1.00] | -1.5% | [-2.3%, +0.5%] | 1.3 pts | 0.8 | yes |
| multi | url | 4096 | churn | ordered | baseline | 6 | 117 | 117 | 1.00× [1.00, 1.01] | +0.2% | [-0.2%, +1.3%] | 0.7 pts | 0.4 | yes |
| multi | url | 4096 | build | ordered | baseline | 6 | 9.32 ms | 9.47 ms | 1.01× [1.01, 1.02] | +1.2% | [+0.7%, +2.1%] | 0.7 pts | 1.5 | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 119 | 99.3 | 0.83× [0.78, 0.91] | -19.9% | [-28.9%, -9.7%] | 9.1 pts | 1.5 | no |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4036 | 4001 | 1.00× [0.98, 1.01] | -0.2% | [-2.2%, +0.9%] | 1.5 pts | 0.9 | yes |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 1395 | 1348 | 0.98× [0.96, 1.00] | -1.6% | [-3.8%, +0.5%] | 2.1 pts | 0.8 | no |
| multi | url | 16384 | churn | ordered | baseline | 6 | 187 | 157 | 0.85× [0.82, 0.86] | -18.2% | [-21.5%, -16.6%] | 2.3 pts | 0.6 | no |
| multi | url | 16384 | build | ordered | baseline | 6 | 47.31 ms | 48.20 ms | 1.02× [1.01, 1.02] | +1.6% | [+0.9%, +2.4%] | 0.8 pts | 1.1 | yes |
| multi | url | 65536 | valuesFor | ordered | baseline | 6 | 286 | 198 | 0.71× [0.65, 0.77] | -39.9% | [-53.6%, -30.2%] | 11.2 pts | 2.0 | no |
| multi | url | 65536 | valuesBetween | ordered | baseline | 6 | 8425 | 6011 | 0.73× [0.64, 0.83] | -36.9% | [-55.5%, -20.5%] | 16.7 pts | 2.9 | no |
| multi | url | 65536 | prefix | ordered | baseline | 6 | 9965 | 7174 | 0.78× [0.69, 0.84] | -28.5% | [-44.6%, -18.7%] | 12.3 pts | 1.4 | no |
| multi | url | 65536 | churn | ordered | baseline | 6 | 378 | 326 | 0.85× [0.78, 0.95] | -18.3% | [-27.7%, -4.8%] | 10.9 pts | 3.1 | no |
| multi | url | 65536 | build | ordered | baseline | 6 | 347.19 ms | 354.10 ms | 1.03× [1.00, 1.04] | +2.6% | [+0.4%, +3.8%] | 1.6 pts | 1.6 | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 6 | 461 | 425 | 0.93× [0.87, 0.96] | -8.0% | [-14.4%, -4.1%] | 4.9 pts | 2.5 | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 6 | 11.7 µs | 10.7 µs | 0.92× [0.91, 0.94] | -8.9% | [-9.8%, -6.4%] | 1.7 pts | 1.2 | yes |
| multi | url | 262144 | prefix | ordered | baseline | 6 | 55.7 µs | 51.2 µs | 0.92× [0.89, 0.94] | -8.3% | [-11.8%, -5.9%] | 2.8 pts | 0.5 | no |
| multi | url | 262144 | churn | ordered | baseline | 6 | 546 | 551 | 1.01× [0.80, 1.16] | +1.1% | [-24.4%, +13.8%] | 18.2 pts | 5.1 | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 6 | 578 | 547 | 0.95× [0.88, 1.00] | -5.7% | [-13.3%, +0.3%] | 6.5 pts | 3.2 | no |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 6 | 13.3 µs | 12.4 µs | 0.93× [0.93, 0.94] | -7.2% | [-8.0%, -6.0%] | 0.9 pts | 0.6 | yes |
| multi | url | 1048576 | prefix | ordered | baseline | 6 | 210.5 µs | 201.0 µs | 0.96× [0.89, 1.19] | -4.3% | [-12.9%, +15.7%] | 13.6 pts | 0.8 | no |
| multi | url | 1048576 | churn | ordered | baseline | 6 | 793 | 696 | 0.86× [0.73, 1.02] | -16.8% | [-37.0%, +1.8%] | 18.5 pts | 8.3 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 58.9 | 57.8 | 0.98× [0.97, 0.99] | -1.8% | [-2.8%, -1.1%] | 0.8 pts | 0.9 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3372 | 3392 | 1.00× [0.99, 1.01] | +0.2% | [-0.9%, +1.3%] | 1.0 pts | 0.5 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 90.7 | 89.5 | 0.99× [0.98, 0.99] | -1.2% | [-2.3%, -0.5%] | 0.8 pts | 0.9 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 73.7 | 70.9 | 0.96× [0.95, 0.98] | -3.9% | [-5.4%, -1.5%] | 1.8 pts | 0.9 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 6.18 ms | 6.04 ms | 0.97× [0.97, 0.98] | -2.6% | [-3.1%, -1.5%] | 0.7 pts | 1.0 | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 72.0 | 67.4 | 0.94× [0.83, 1.01] | -6.7% | [-21.1%, +1.3%] | 10.7 pts | 1.1 | no |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3687 | 3655 | 0.98× [0.95, 1.00] | -1.7% | [-5.2%, -0.1%] | 2.4 pts | 1.6 | no |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 115 | 110 | 0.96× [0.94, 0.98] | -3.9% | [-5.9%, -2.5%] | 1.6 pts | 1.0 | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 109 | 99.7 | 0.90× [0.85, 0.93] | -11.3% | [-17.9%, -7.8%] | 4.8 pts | 1.1 | no |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 31.57 ms | 30.02 ms | 0.95× [0.94, 0.96] | -5.5% | [-6.3%, -4.6%] | 0.8 pts | 1.3 | yes |
| multi | uuid | 65536 | valuesFor | ordered | baseline | 6 | 189 | 153 | 0.81× [0.70, 0.92] | -22.7% | [-43.0%, -8.2%] | 16.6 pts | 2.5 | no |
| multi | uuid | 65536 | valuesBetween | ordered | baseline | 6 | 8056 | 5666 | 0.82× [0.71, 0.91] | -22.2% | [-41.7%, -10.3%] | 15.0 pts | 2.7 | no |
| multi | uuid | 65536 | prefix | ordered | baseline | 6 | 309 | 292 | 1.01× [0.84, 1.11] | +1.5% | [-18.6%, +9.6%] | 13.4 pts | 3.2 | no |
| multi | uuid | 65536 | churn | ordered | baseline | 6 | 275 | 200 | 0.74× [0.67, 0.91] | -34.4% | [-50.1%, -9.7%] | 19.3 pts | 4.2 | no |
| multi | uuid | 65536 | build | ordered | baseline | 6 | 240.77 ms | 219.31 ms | 0.92× [0.90, 0.93] | -8.8% | [-11.4%, -7.6%] | 1.8 pts | 1.3 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 6 | 351 | 296 | 0.84× [0.82, 0.86] | -19.4% | [-22.3%, -16.6%] | 2.7 pts | 1.0 | no |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 6 | 10.8 µs | 10.3 µs | 0.95× [0.93, 0.96] | -5.5% | [-7.4%, -3.9%] | 1.7 pts | 0.9 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 6 | 746 | 656 | 0.90× [0.85, 0.92] | -11.7% | [-17.4%, -8.5%] | 4.2 pts | 2.1 | no |
| multi | uuid | 262144 | churn | ordered | baseline | 6 | 410 | 373 | 0.92× [0.79, 0.98] | -9.0% | [-26.6%, -1.6%] | 11.9 pts | 5.3 | no |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 6 | 443 | 415 | 0.94× [0.87, 0.97] | -6.9% | [-14.7%, -3.5%] | 5.3 pts | 2.7 | no |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 6 | 12.5 µs | 11.9 µs | 0.95× [0.94, 0.96] | -5.0% | [-6.6%, -3.6%] | 1.4 pts | 1.2 | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 6 | 2325 | 2165 | 0.93× [0.91, 0.94] | -8.0% | [-9.6%, -5.9%] | 1.7 pts | 1.1 | yes |
| multi | uuid | 1048576 | churn | ordered | baseline | 6 | 611 | 561 | 0.95× [0.77, 1.15] | -5.5% | [-30.4%, +12.9%] | 20.6 pts | 9.2 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
