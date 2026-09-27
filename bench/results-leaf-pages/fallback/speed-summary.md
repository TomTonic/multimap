| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 6 | 51.7 | 49.3 | 0.95× [0.95, 0.96] | -5.1% | [-5.6%, -4.6%] | 0.5 pts | 0.5 | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 6 | 2916 | 2939 | 1.00× [0.98, 1.02] | +0.0% | [-1.5%, +2.1%] | 1.8 pts | 0.8 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 6 | 76.5 | 75.0 | 0.98× [0.98, 0.98] | -2.0% | [-2.5%, -1.8%] | 0.3 pts | 0.4 | yes |
| multi | email | 4096 | churn | ordered | baseline | 6 | 65.2 | 57.9 | 0.88× [0.88, 0.90] | -13.2% | [-14.2%, -11.6%] | 1.2 pts | 0.6 | yes |
| multi | email | 4096 | build | ordered | baseline | 6 | 5.68 ms | 4.82 ms | 0.86× [0.83, 0.90] | -16.8% | [-20.4%, -11.0%] | 4.5 pts | 3.3 | no |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 64.8 | 60.4 | 0.92× [0.87, 1.00] | -8.7% | [-14.5%, +0.2%] | 7.0 pts | 3.8 | no |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3495 | 3545 | 1.01× [0.98, 1.02] | +0.8% | [-1.6%, +2.1%] | 1.8 pts | 1.4 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 95.9 | 92.8 | 0.96× [0.94, 0.98] | -4.1% | [-6.9%, -1.7%] | 2.5 pts | 1.7 | no |
| multi | email | 16384 | churn | ordered | baseline | 6 | 95.6 | 81.8 | 0.86× [0.84, 0.89] | -16.0% | [-19.1%, -12.5%] | 3.2 pts | 1.0 | no |
| multi | email | 16384 | build | ordered | baseline | 6 | 28.88 ms | 24.53 ms | 0.85× [0.85, 0.85] | -17.6% | [-18.1%, -17.2%] | 0.5 pts | 0.5 | yes |
| multi | email | 65536 | valuesFor | ordered | baseline | 6 | 192 | 102 | 0.53× [0.49, 0.59] | -89.6% | [-104.0%, -69.8%] | 16.3 pts | 0.9 | no |
| multi | email | 65536 | valuesBetween | ordered | baseline | 6 | 6198 | 4694 | 0.79× [0.72, 0.86] | -26.8% | [-39.5%, -16.5%] | 10.9 pts | 1.5 | no |
| multi | email | 65536 | prefix | ordered | baseline | 6 | 236 | 172 | 0.78× [0.70, 0.84] | -28.4% | [-43.2%, -18.7%] | 11.6 pts | 1.4 | no |
| multi | email | 65536 | churn | ordered | baseline | 6 | 270 | 174 | 0.64× [0.59, 0.72] | -55.7% | [-69.2%, -39.6%] | 14.1 pts | 2.7 | no |
| multi | email | 65536 | build | ordered | baseline | 6 | 216.71 ms | 190.25 ms | 0.87× [0.85, 0.91] | -15.3% | [-17.2%, -10.0%] | 3.4 pts | 2.6 | no |
| multi | email | 262144 | valuesFor | ordered | baseline | 6 | 395 | 255 | 0.65× [0.59, 0.71] | -54.9% | [-69.9%, -40.5%] | 14.0 pts | 4.1 | no |
| multi | email | 262144 | valuesBetween | ordered | baseline | 6 | 9063 | 9057 | 1.00× [0.99, 1.01] | +0.0% | [-1.5%, +1.0%] | 1.2 pts | 0.7 | yes |
| multi | email | 262144 | prefix | ordered | baseline | 6 | 432 | 319 | 0.75× [0.67, 0.84] | -33.4% | [-49.4%, -18.9%] | 14.5 pts | 3.5 | no |
| multi | email | 262144 | churn | ordered | baseline | 6 | 410 | 320 | 0.78× [0.72, 0.92] | -28.9% | [-39.7%, -9.2%] | 14.5 pts | 4.9 | no |
| multi | email | 1048576 | valuesFor | ordered | baseline | 6 | 494 | 377 | 0.76× [0.70, 0.82] | -32.0% | [-42.3%, -22.3%] | 9.5 pts | 4.8 | no |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 6 | 11.1 µs | 11.3 µs | 1.02× [1.01, 1.03] | +1.6% | [+0.7%, +2.9%] | 1.1 pts | 0.8 | yes |
| multi | email | 1048576 | prefix | ordered | baseline | 6 | 653 | 641 | 0.98× [0.95, 1.00] | -2.3% | [-5.3%, -0.4%] | 2.3 pts | 1.0 | no |
| multi | email | 1048576 | churn | ordered | baseline | 6 | 607 | 468 | 0.78× [0.70, 0.92] | -28.1% | [-43.1%, -8.8%] | 16.4 pts | 4.7 | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 5 | 108 | 105 | 0.97× [0.97, 0.97] | -3.3% | [-3.6%, -2.6%] | 0.4 pts | 0.4 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 5 | 4121 | 4074 | 0.99× [0.98, 1.01] | -1.1% | [-2.1%, +0.8%] | 1.2 pts | 0.8 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 5 | 332 | 322 | 0.98× [0.96, 0.99] | -2.5% | [-4.4%, -1.3%] | 1.2 pts | 0.6 | yes |
| multi | path | 4096 | churn | ordered | baseline | 5 | 153 | 147 | 0.96× [0.95, 0.98] | -4.4% | [-5.1%, -2.5%] | 1.1 pts | 0.7 | yes |
| multi | path | 4096 | build | ordered | baseline | 5 | 12.12 ms | 11.05 ms | 0.91× [0.91, 0.92] | -9.5% | [-10.4%, -9.1%] | 0.5 pts | 0.9 | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 6 | 152 | 136 | 0.89× [0.83, 0.95] | -12.5% | [-20.3%, -5.5%] | 7.0 pts | 1.3 | no |
| multi | path | 16384 | valuesBetween | ordered | baseline | 6 | 4791 | 4718 | 0.99× [0.98, 1.00] | -1.0% | [-2.4%, -0.0%] | 1.1 pts | 0.6 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 6 | 660 | 639 | 0.95× [0.91, 1.00] | -4.7% | [-9.8%, -0.2%] | 4.5 pts | 0.8 | no |
| multi | path | 16384 | churn | ordered | baseline | 6 | 243 | 222 | 0.92× [0.86, 0.99] | -8.8% | [-15.8%, -0.9%] | 7.1 pts | 1.7 | no |
| multi | path | 16384 | build | ordered | baseline | 6 | 66.51 ms | 63.42 ms | 0.95× [0.95, 0.96] | -4.7% | [-5.6%, -4.0%] | 0.8 pts | 0.8 | yes |
| multi | path | 65536 | valuesFor | ordered | baseline | 6 | 376 | 291 | 0.77× [0.71, 0.84] | -30.6% | [-40.6%, -19.7%] | 9.9 pts | 2.5 | no |
| multi | path | 65536 | valuesBetween | ordered | baseline | 6 | 7919 | 6443 | 0.82× [0.74, 0.86] | -21.7% | [-34.4%, -16.5%] | 8.5 pts | 1.4 | no |
| multi | path | 65536 | prefix | ordered | baseline | 6 | 2326 | 2132 | 0.93× [0.88, 0.98] | -8.0% | [-14.1%, -2.0%] | 5.8 pts | 0.7 | no |
| multi | path | 65536 | churn | ordered | baseline | 6 | 511 | 431 | 0.83× [0.74, 0.94] | -20.8% | [-34.7%, -6.1%] | 13.6 pts | 3.3 | no |
| multi | path | 65536 | build | ordered | baseline | 6 | 459.76 ms | 450.44 ms | 0.99× [0.95, 1.03] | -1.4% | [-5.3%, +2.7%] | 3.8 pts | 4.4 | no |
| multi | path | 262144 | valuesFor | ordered | baseline | 6 | 602 | 502 | 0.84× [0.81, 0.87] | -19.3% | [-23.5%, -15.3%] | 3.9 pts | 1.3 | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 6 | 12.0 µs | 12.1 µs | 1.01× [1.00, 1.04] | +1.5% | [-0.2%, +3.8%] | 1.9 pts | 1.1 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 6 | 10.6 µs | 11.0 µs | 1.04× [1.00, 1.07] | +3.7% | [-0.2%, +6.5%] | 3.2 pts | 0.2 | no |
| multi | path | 262144 | churn | ordered | baseline | 6 | 840 | 798 | 0.96× [0.87, 1.02] | -4.2% | [-14.8%, +2.4%] | 8.2 pts | 3.0 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 5 | 62.5 | 60.9 | 0.97× [0.96, 0.98] | -2.7% | [-3.6%, -1.8%] | 0.7 pts | 0.9 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 5 | 3408 | 3447 | 1.02× [1.00, 1.03] | +1.7% | [+0.1%, +3.2%] | 1.2 pts | 0.6 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 5 | 7546 | 7754 | 1.03× [1.00, 1.04] | +2.5% | [+0.2%, +4.2%] | 1.6 pts | 0.4 | yes |
| multi | str | 4096 | churn | ordered | baseline | 5 | 79.2 | 72.4 | 0.91× [0.90, 0.93] | -10.0% | [-11.1%, -8.1%] | 1.2 pts | 0.9 | yes |
| multi | str | 4096 | build | ordered | baseline | 5 | 6.63 ms | 5.79 ms | 0.87× [0.87, 0.88] | -14.7% | [-15.3%, -13.3%] | 0.8 pts | 1.1 | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 75.4 | 72.8 | 0.97× [0.94, 0.98] | -2.8% | [-5.9%, -1.5%] | 2.1 pts | 2.2 | no |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3701 | 3641 | 0.99× [0.96, 1.00] | -1.5% | [-3.8%, -0.3%] | 1.7 pts | 1.2 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 34.4 µs | 34.2 µs | 1.00× [0.97, 1.01] | -0.3% | [-3.0%, +0.9%] | 1.9 pts | 1.3 | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 114 | 101 | 0.89× [0.81, 0.94] | -12.6% | [-22.9%, -6.8%] | 7.7 pts | 1.8 | no |
| multi | str | 16384 | build | ordered | baseline | 6 | 33.71 ms | 29.48 ms | 0.88× [0.87, 0.88] | -14.2% | [-14.7%, -13.9%] | 0.4 pts | 0.6 | yes |
| multi | str | 65536 | valuesFor | ordered | baseline | 6 | 204 | 133 | 0.65× [0.53, 0.83] | -53.8% | [-89.4%, -19.9%] | 33.1 pts | 2.5 | no |
| multi | str | 65536 | valuesBetween | ordered | baseline | 6 | 6285 | 5346 | 0.89× [0.76, 1.03] | -12.9% | [-31.7%, +3.3%] | 16.7 pts | 2.8 | no |
| multi | str | 65536 | prefix | ordered | baseline | 6 | 231.7 µs | 192.8 µs | 0.84× [0.78, 0.96] | -19.2% | [-27.6%, -4.6%] | 10.9 pts | 1.5 | no |
| multi | str | 65536 | churn | ordered | baseline | 6 | 268 | 197 | 0.74× [0.66, 0.87] | -34.3% | [-51.0%, -15.5%] | 16.9 pts | 3.2 | no |
| multi | str | 65536 | build | ordered | baseline | 6 | 226.00 ms | 204.20 ms | 0.91× [0.90, 0.92] | -10.1% | [-11.2%, -9.2%] | 1.0 pts | 0.9 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 6 | 402 | 293 | 0.74× [0.70, 0.83] | -34.8% | [-43.4%, -20.6%] | 10.9 pts | 2.9 | no |
| multi | str | 262144 | valuesBetween | ordered | baseline | 6 | 9729 | 9902 | 1.02× [1.00, 1.03] | +1.6% | [-0.2%, +3.1%] | 1.6 pts | 1.1 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 6 | 1.56 ms | 1.60 ms | 1.02× [0.99, 1.06] | +2.0% | [-1.2%, +5.3%] | 3.1 pts | 1.4 | no |
| multi | str | 262144 | churn | ordered | baseline | 6 | 456 | 357 | 0.79× [0.74, 0.85] | -27.0% | [-35.5%, -17.7%] | 8.5 pts | 2.4 | no |
| multi | str | 1048576 | valuesFor | ordered | baseline | 6 | 554 | 435 | 0.82× [0.75, 0.85] | -21.7% | [-32.6%, -18.0%] | 7.0 pts | 2.2 | no |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 6 | 11.3 µs | 11.5 µs | 1.02× [1.01, 1.04] | +2.3% | [+0.8%, +3.4%] | 1.2 pts | 0.8 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 6 | 6.65 ms | 7.13 ms | 1.07× [1.06, 1.08] | +6.7% | [+5.9%, +7.2%] | 0.6 pts | 1.6 | yes |
| multi | str | 1048576 | churn | ordered | baseline | 6 | 639 | 513 | 0.81× [0.72, 0.92] | -23.2% | [-39.0%, -8.3%] | 14.6 pts | 5.6 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 6 | 57.7 | 54.6 | 0.94× [0.93, 0.96] | -6.2% | [-7.0%, -4.7%] | 1.1 pts | 0.9 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 6 | 2367 | 2364 | 0.99× [0.98, 1.02] | -0.5% | [-2.0%, +2.3%] | 2.1 pts | 1.2 | no |
| multi | street | 4096 | prefix | ordered | baseline | 6 | 251 | 257 | 1.02× [1.00, 1.04] | +2.2% | [+0.1%, +3.7%] | 1.7 pts | 0.8 | yes |
| multi | street | 4096 | churn | ordered | baseline | 6 | 92.8 | 85.8 | 0.93× [0.91, 0.94] | -7.1% | [-9.5%, -6.1%] | 1.6 pts | 0.8 | yes |
| multi | street | 4096 | build | ordered | baseline | 6 | 3.12 ms | 2.62 ms | 0.84× [0.83, 0.84] | -19.2% | [-20.4%, -18.4%] | 0.9 pts | 1.0 | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 79.4 | 75.4 | 0.95× [0.93, 0.96] | -5.0% | [-7.0%, -3.7%] | 1.6 pts | 1.0 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2873 | 2857 | 0.99× [0.98, 1.01] | -0.7% | [-2.3%, +0.6%] | 1.4 pts | 1.0 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 909 | 919 | 1.00× [0.99, 1.02] | +0.4% | [-0.7%, +2.1%] | 1.3 pts | 0.4 | yes |
| multi | street | 16384 | churn | ordered | baseline | 6 | 131 | 120 | 0.92× [0.91, 0.93] | -8.9% | [-10.1%, -7.2%] | 1.4 pts | 0.6 | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 16.74 ms | 14.26 ms | 0.86× [0.85, 0.86] | -16.7% | [-17.5%, -16.2%] | 0.6 pts | 1.2 | yes |
| multi | street | 65536 | valuesFor | ordered | baseline | 6 | 160 | 109 | 0.71× [0.65, 0.76] | -40.1% | [-53.8%, -31.2%] | 10.8 pts | 1.2 | no |
| multi | street | 65536 | valuesBetween | ordered | baseline | 6 | 4332 | 3625 | 0.84× [0.80, 0.96] | -19.0% | [-25.2%, -4.0%] | 10.1 pts | 0.9 | no |
| multi | street | 65536 | prefix | ordered | baseline | 6 | 4966 | 4670 | 0.93× [0.91, 0.96] | -7.3% | [-9.9%, -4.1%] | 2.8 pts | 0.5 | no |
| multi | street | 65536 | churn | ordered | baseline | 6 | 270 | 210 | 0.79× [0.72, 0.92] | -25.9% | [-38.7%, -8.5%] | 14.4 pts | 2.7 | no |
| multi | street | 65536 | build | ordered | baseline | 6 | 105.34 ms | 94.31 ms | 0.92× [0.89, 0.93] | -9.0% | [-12.4%, -7.1%] | 2.5 pts | 2.9 | no |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 6 | 38.2 | 36.9 | 0.96× [0.95, 0.97] | -3.7% | [-5.3%, -2.8%] | 1.2 pts | 1.0 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 2654 | 2687 | 1.01× [0.99, 1.02] | +0.6% | [-0.8%, +2.4%] | 1.5 pts | 0.6 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 6 | 43.7 | 39.6 | 0.91× [0.89, 0.92] | -10.4% | [-12.6%, -8.8%] | 1.8 pts | 0.9 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 6 | 4.13 ms | 3.55 ms | 0.86× [0.85, 0.87] | -16.9% | [-18.0%, -14.9%] | 1.5 pts | 0.8 | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 46.8 | 44.0 | 0.94× [0.91, 0.96] | -6.1% | [-9.3%, -4.4%] | 2.4 pts | 3.0 | no |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3531 | 3549 | 1.00× [0.99, 1.01] | +0.2% | [-0.7%, +1.1%] | 0.9 pts | 0.8 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 53.6 | 46.4 | 0.87× [0.80, 0.90] | -15.3% | [-25.2%, -11.1%] | 6.8 pts | 1.7 | no |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 18.68 ms | 16.77 ms | 0.90× [0.89, 0.90] | -11.4% | [-11.7%, -11.0%] | 0.3 pts | 0.5 | yes |
| multi | u64 | 65536 | valuesFor | ordered | baseline | 6 | 75.3 | 67.2 | 0.85× [0.81, 0.93] | -17.3% | [-24.0%, -7.0%] | 8.1 pts | 1.2 | no |
| multi | u64 | 65536 | valuesBetween | ordered | baseline | 6 | 5329 | 4724 | 0.90× [0.84, 0.94] | -11.3% | [-18.4%, -6.2%] | 5.8 pts | 1.0 | no |
| multi | u64 | 65536 | churn | ordered | baseline | 6 | 128 | 91.1 | 0.70× [0.64, 0.78] | -42.5% | [-55.5%, -28.8%] | 12.7 pts | 1.8 | no |
| multi | u64 | 65536 | build | ordered | baseline | 6 | 120.71 ms | 106.73 ms | 0.89× [0.87, 0.90] | -12.2% | [-14.5%, -10.6%] | 1.9 pts | 0.8 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 6 | 168 | 137 | 0.81× [0.77, 0.84] | -23.9% | [-30.1%, -19.4%] | 5.1 pts | 2.2 | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 6 | 8636 | 7854 | 0.91× [0.87, 0.97] | -10.0% | [-15.0%, -2.8%] | 5.8 pts | 3.8 | no |
| multi | u64 | 262144 | churn | ordered | baseline | 6 | 260 | 195 | 0.76× [0.73, 0.78] | -31.9% | [-37.5%, -27.7%] | 4.7 pts | 1.2 | no |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 6 | 221 | 194 | 0.87× [0.86, 0.88] | -15.5% | [-16.8%, -13.7%] | 1.4 pts | 0.8 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 6 | 8103 | 7891 | 0.97× [0.95, 0.99] | -2.8% | [-5.0%, -1.4%] | 1.7 pts | 1.4 | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 6 | 397 | 316 | 0.80× [0.77, 0.90] | -24.9% | [-30.3%, -11.2%] | 9.1 pts | 2.7 | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 5 | 84.6 | 81.8 | 0.97× [0.95, 0.98] | -3.4% | [-5.0%, -2.2%] | 1.1 pts | 2.1 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 5 | 3670 | 3607 | 0.97× [0.96, 0.99] | -3.0% | [-4.5%, -0.6%] | 1.6 pts | 0.9 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 5 | 412 | 402 | 0.98× [0.96, 0.99] | -2.4% | [-4.5%, -0.5%] | 1.6 pts | 0.9 | yes |
| multi | url | 4096 | churn | ordered | baseline | 5 | 115 | 111 | 0.96× [0.95, 0.97] | -4.6% | [-5.5%, -3.4%] | 0.8 pts | 0.5 | yes |
| multi | url | 4096 | build | ordered | baseline | 5 | 9.30 ms | 8.64 ms | 0.93× [0.92, 0.94] | -7.9% | [-8.9%, -6.9%] | 0.8 pts | 1.4 | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 101 | 95.6 | 0.95× [0.92, 0.96] | -5.7% | [-8.2%, -3.8%] | 2.1 pts | 0.8 | no |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4038 | 3999 | 0.98× [0.96, 1.00] | -2.3% | [-3.8%, +0.3%] | 1.9 pts | 1.5 | no |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 1361 | 1354 | 1.00× [0.98, 1.01] | +0.2% | [-2.5%, +1.3%] | 1.8 pts | 0.7 | yes |
| multi | url | 16384 | churn | ordered | baseline | 6 | 171 | 144 | 0.84× [0.80, 0.87] | -18.7% | [-25.7%, -14.7%] | 5.2 pts | 1.1 | no |
| multi | url | 16384 | build | ordered | baseline | 6 | 47.85 ms | 44.05 ms | 0.92× [0.91, 0.93] | -8.1% | [-9.3%, -7.1%] | 1.0 pts | 1.1 | yes |
| multi | url | 65536 | valuesFor | ordered | baseline | 6 | 323 | 226 | 0.68× [0.62, 0.73] | -46.4% | [-60.1%, -36.2%] | 11.4 pts | 2.5 | no |
| multi | url | 65536 | valuesBetween | ordered | baseline | 6 | 7311 | 5853 | 0.81× [0.73, 0.92] | -23.2% | [-36.7%, -8.3%] | 13.5 pts | 2.0 | no |
| multi | url | 65536 | prefix | ordered | baseline | 6 | 7057 | 6707 | 0.97× [0.89, 1.01] | -3.3% | [-11.8%, +0.9%] | 6.0 pts | 0.8 | no |
| multi | url | 65536 | churn | ordered | baseline | 6 | 380 | 299 | 0.78× [0.75, 0.82] | -27.7% | [-32.6%, -22.0%] | 5.0 pts | 1.6 | no |
| multi | url | 65536 | build | ordered | baseline | 6 | 351.97 ms | 322.42 ms | 0.91× [0.87, 0.97] | -10.0% | [-14.7%, -3.5%] | 5.3 pts | 3.9 | no |
| multi | url | 262144 | valuesFor | ordered | baseline | 6 | 474 | 419 | 0.89× [0.85, 0.92] | -12.3% | [-18.3%, -8.6%] | 4.6 pts | 3.1 | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 6 | 10.3 µs | 10.3 µs | 1.00× [0.93, 1.03] | +0.0% | [-7.3%, +3.2%] | 5.0 pts | 3.6 | no |
| multi | url | 262144 | prefix | ordered | baseline | 6 | 49.7 µs | 49.4 µs | 1.00× [0.96, 1.04] | +0.2% | [-4.1%, +4.0%] | 3.9 pts | 0.7 | no |
| multi | url | 262144 | churn | ordered | baseline | 6 | 522 | 488 | 0.92× [0.88, 0.96] | -8.7% | [-13.2%, -3.8%] | 4.5 pts | 1.3 | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 6 | 606 | 538 | 0.89× [0.83, 0.93] | -11.8% | [-20.0%, -7.8%] | 5.8 pts | 2.9 | no |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 6 | 12.0 µs | 12.3 µs | 1.03× [0.99, 1.06] | +2.7% | [-1.3%, +5.3%] | 3.1 pts | 3.1 | no |
| multi | url | 1048576 | prefix | ordered | baseline | 6 | 192.8 µs | 203.2 µs | 1.06× [1.04, 1.07] | +5.4% | [+4.0%, +6.7%] | 1.3 pts | 0.1 | yes |
| multi | url | 1048576 | churn | ordered | baseline | 6 | 750 | 628 | 0.84× [0.78, 0.92] | -18.8% | [-28.6%, -9.3%] | 9.2 pts | 4.8 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 5 | 56.7 | 54.8 | 0.97× [0.96, 0.98] | -3.5% | [-4.4%, -2.0%] | 1.0 pts | 1.1 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 5 | 3227 | 3291 | 1.02× [1.00, 1.04] | +2.0% | [-0.2%, +3.5%] | 1.5 pts | 0.8 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 5 | 88.8 | 87.7 | 0.99× [0.98, 1.00] | -1.3% | [-1.9%, -0.5%] | 0.6 pts | 0.6 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 5 | 70.8 | 64.9 | 0.90× [0.90, 0.92] | -10.6% | [-11.4%, -9.1%] | 0.9 pts | 0.7 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 5 | 6.12 ms | 5.39 ms | 0.88× [0.87, 0.89] | -13.7% | [-14.5%, -12.5%] | 0.8 pts | 0.9 | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 67.6 | 64.4 | 0.95× [0.91, 0.97] | -5.2% | [-9.6%, -2.8%] | 3.2 pts | 1.8 | no |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3640 | 3616 | 1.00× [0.98, 1.01] | -0.2% | [-1.8%, +1.2%] | 1.4 pts | 1.3 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 111 | 110 | 0.99× [0.98, 1.00] | -1.3% | [-2.1%, -0.1%] | 0.9 pts | 0.7 | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 109 | 90.8 | 0.87× [0.83, 0.91] | -15.5% | [-20.7%, -9.7%] | 5.2 pts | 1.1 | no |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 31.17 ms | 27.16 ms | 0.87× [0.86, 0.89] | -15.4% | [-16.7%, -13.0%] | 1.8 pts | 1.8 | yes |
| multi | uuid | 65536 | valuesFor | ordered | baseline | 6 | 220 | 125 | 0.57× [0.51, 0.68] | -76.6% | [-95.5%, -46.6%] | 23.3 pts | 2.1 | no |
| multi | uuid | 65536 | valuesBetween | ordered | baseline | 6 | 6841 | 5432 | 0.86× [0.74, 1.00] | -16.6% | [-35.0%, -0.0%] | 16.7 pts | 2.8 | no |
| multi | uuid | 65536 | prefix | ordered | baseline | 6 | 280 | 249 | 0.84× [0.80, 0.93] | -18.7% | [-25.7%, -7.8%] | 8.5 pts | 1.8 | no |
| multi | uuid | 65536 | churn | ordered | baseline | 6 | 253 | 189 | 0.73× [0.64, 0.77] | -36.4% | [-55.2%, -29.6%] | 12.2 pts | 2.1 | no |
| multi | uuid | 65536 | build | ordered | baseline | 6 | 230.56 ms | 194.80 ms | 0.85× [0.83, 0.86] | -18.1% | [-19.8%, -15.9%] | 1.9 pts | 1.3 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 6 | 390 | 266 | 0.69× [0.65, 0.72] | -45.4% | [-52.8%, -39.2%] | 6.5 pts | 3.0 | no |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 6 | 9761 | 9894 | 1.02× [1.00, 1.03] | +2.0% | [-0.0%, +3.1%] | 1.5 pts | 0.9 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 6 | 653 | 641 | 0.97× [0.93, 1.02] | -3.0% | [-7.7%, +2.4%] | 4.8 pts | 2.2 | no |
| multi | uuid | 262144 | churn | ordered | baseline | 6 | 437 | 363 | 0.86× [0.81, 0.92] | -16.8% | [-23.0%, -9.1%] | 6.6 pts | 2.8 | no |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 6 | 549 | 413 | 0.77× [0.75, 0.79] | -29.7% | [-33.0%, -26.1%] | 3.3 pts | 1.2 | no |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 6 | 11.3 µs | 11.9 µs | 1.06× [1.04, 1.07] | +5.8% | [+3.7%, +6.7%] | 1.4 pts | 1.1 | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 6 | 2034 | 2134 | 1.05× [1.02, 1.08] | +5.0% | [+2.2%, +7.3%] | 2.4 pts | 1.8 | no |
| multi | uuid | 1048576 | churn | ordered | baseline | 6 | 614 | 549 | 0.90× [0.80, 0.97] | -11.4% | [-25.1%, -3.0%] | 10.5 pts | 3.1 | no |
| unique | email | 4096 | valuesFor | ordered | baseline | 5 | 40.1 | 26.1 | 0.64× [0.63, 0.66] | -55.1% | [-59.4%, -50.5%] | 3.6 pts | 1.7 | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 5 | 336 | 1276 | 3.79× [3.70, 3.95] | +73.6% | [+73.0%, +74.7%] | 0.7 pts | 3.3 | yes |
| unique | email | 4096 | prefix | ordered | baseline | 5 | 59.4 | 55.8 | 0.94× [0.92, 0.96] | -6.8% | [-8.2%, -4.4%] | 1.5 pts | 1.9 | yes |
| unique | email | 4096 | churn | ordered | baseline | 5 | 92.1 | 68.2 | 0.74× [0.73, 0.75] | -35.0% | [-36.5%, -33.7%] | 1.1 pts | 0.7 | yes |
| unique | email | 4096 | build | ordered | baseline | 5 | 1.25 ms | 850.1 µs | 0.68× [0.66, 0.69] | -47.4% | [-50.6%, -44.9%] | 2.3 pts | 1.0 | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 57.5 | 37.9 | 0.66× [0.65, 0.68] | -51.3% | [-54.2%, -47.1%] | 3.4 pts | 2.9 | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 448 | 1600 | 3.58× [3.51, 3.62] | +72.1% | [+71.5%, +72.4%] | 0.4 pts | 1.3 | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 84.7 | 70.4 | 0.83× [0.82, 0.84] | -20.0% | [-21.6%, -19.2%] | 1.2 pts | 1.0 | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 128 | 92.6 | 0.72× [0.69, 0.75] | -38.8% | [-44.1%, -32.6%] | 5.5 pts | 2.7 | no |
| unique | email | 16384 | build | ordered | baseline | 6 | 6.74 ms | 4.06 ms | 0.60× [0.60, 0.60] | -66.6% | [-67.7%, -65.3%] | 1.1 pts | 0.8 | yes |
| unique | email | 65536 | valuesFor | ordered | baseline | 6 | 71.0 | 58.1 | 0.82× [0.78, 0.88] | -21.7% | [-28.1%, -13.1%] | 7.1 pts | 1.7 | no |
| unique | email | 65536 | valuesBetween | ordered | baseline | 6 | 518 | 1584 | 3.06× [2.97, 3.17] | +67.3% | [+66.3%, +68.5%] | 1.0 pts | 1.7 | yes |
| unique | email | 65536 | prefix | ordered | baseline | 6 | 102 | 104 | 1.00× [0.96, 1.05] | -0.3% | [-4.2%, +5.1%] | 4.4 pts | 1.2 | no |
| unique | email | 65536 | churn | ordered | baseline | 6 | 165 | 153 | 0.95× [0.90, 0.99] | -5.1% | [-10.6%, -1.1%] | 4.5 pts | 1.4 | no |
| unique | email | 65536 | build | ordered | baseline | 6 | 34.75 ms | 28.30 ms | 0.82× [0.81, 0.83] | -22.0% | [-24.1%, -20.2%] | 1.8 pts | 1.1 | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 6 | 281 | 187 | 0.72× [0.62, 0.84] | -38.8% | [-61.7%, -18.7%] | 20.5 pts | 1.9 | no |
| unique | email | 262144 | valuesBetween | ordered | baseline | 6 | 1527 | 3962 | 2.57× [2.50, 2.65] | +61.0% | [+60.0%, +62.3%] | 1.1 pts | 1.1 | yes |
| unique | email | 262144 | prefix | ordered | baseline | 6 | 276 | 242 | 0.88× [0.69, 1.03] | -14.2% | [-45.8%, +2.8%] | 23.1 pts | 3.8 | no |
| unique | email | 262144 | churn | ordered | baseline | 6 | 518 | 348 | 0.71× [0.65, 0.79] | -40.5% | [-53.9%, -26.9%] | 12.9 pts | 3.1 | no |
| unique | email | 1048576 | valuesFor | ordered | baseline | 6 | 431 | 311 | 0.73× [0.67, 0.76] | -37.0% | [-48.3%, -30.9%] | 8.3 pts | 3.1 | no |
| unique | email | 1048576 | valuesBetween | ordered | baseline | 6 | 1880 | 6490 | 3.44× [3.39, 3.49] | +71.0% | [+70.5%, +71.4%] | 0.4 pts | 1.1 | yes |
| unique | email | 1048576 | prefix | ordered | baseline | 6 | 456 | 431 | 0.94× [0.86, 1.02] | -6.4% | [-16.5%, +2.4%] | 9.0 pts | 2.7 | no |
| unique | email | 1048576 | churn | ordered | baseline | 6 | 671 | 590 | 0.92× [0.80, 1.08] | -9.0% | [-25.5%, +7.7%] | 15.8 pts | 5.7 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 5 | 97.8 | 84.4 | 0.86× [0.85, 0.87] | -16.8% | [-18.1%, -15.1%] | 1.2 pts | 1.2 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 5 | 668 | 2095 | 3.13× [3.07, 3.30] | +68.0% | [+67.4%, +69.7%] | 0.9 pts | 1.7 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 5 | 227 | 247 | 1.09× [1.07, 1.11] | +8.1% | [+6.4%, +9.8%] | 1.4 pts | 0.8 | yes |
| unique | path | 4096 | churn | ordered | baseline | 5 | 218 | 199 | 0.91× [0.90, 0.92] | -10.3% | [-11.5%, -8.5%] | 1.2 pts | 0.8 | yes |
| unique | path | 4096 | build | ordered | baseline | 5 | 2.89 ms | 2.05 ms | 0.71× [0.71, 0.72] | -40.3% | [-41.8%, -39.1%] | 1.1 pts | 0.8 | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 6 | 130 | 114 | 0.88× [0.87, 0.89] | -13.6% | [-15.5%, -12.5%] | 1.4 pts | 1.3 | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 6 | 849 | 2642 | 3.11× [3.08, 3.14] | +67.8% | [+67.6%, +68.2%] | 0.3 pts | 0.7 | yes |
| unique | path | 16384 | prefix | ordered | baseline | 6 | 327 | 433 | 1.35× [1.27, 1.41] | +26.0% | [+21.2%, +29.2%] | 3.8 pts | 1.4 | no |
| unique | path | 16384 | churn | ordered | baseline | 6 | 297 | 277 | 0.94× [0.91, 0.97] | -6.1% | [-10.5%, -3.6%] | 3.3 pts | 1.2 | no |
| unique | path | 16384 | build | ordered | baseline | 6 | 15.12 ms | 11.10 ms | 0.73× [0.73, 0.74] | -36.7% | [-37.1%, -35.7%] | 0.7 pts | 0.5 | yes |
| unique | path | 65536 | valuesFor | ordered | baseline | 6 | 236 | 207 | 0.88× [0.85, 0.96] | -14.1% | [-17.4%, -3.7%] | 6.5 pts | 1.5 | no |
| unique | path | 65536 | valuesBetween | ordered | baseline | 6 | 1197 | 3026 | 2.54× [2.51, 2.59] | +60.6% | [+60.1%, +61.4%] | 0.6 pts | 0.6 | yes |
| unique | path | 65536 | prefix | ordered | baseline | 6 | 677 | 1264 | 1.86× [1.74, 2.00] | +46.2% | [+42.4%, +50.0%] | 3.6 pts | 1.1 | yes |
| unique | path | 65536 | churn | ordered | baseline | 6 | 442 | 432 | 1.00× [0.90, 1.05] | -0.2% | [-10.5%, +4.6%] | 7.2 pts | 1.4 | no |
| unique | path | 65536 | build | ordered | baseline | 6 | 91.23 ms | 71.02 ms | 0.78× [0.77, 0.79] | -28.7% | [-30.2%, -26.2%] | 1.9 pts | 1.6 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 6 | 527 | 435 | 0.84× [0.77, 0.93] | -19.6% | [-29.9%, -7.0%] | 10.9 pts | 2.5 | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 6 | 2684 | 6708 | 2.49× [2.34, 2.56] | +59.8% | [+57.2%, +60.9%] | 1.8 pts | 1.6 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 6 | 1975 | 5707 | 2.89× [2.71, 3.23] | +65.4% | [+63.1%, +69.1%] | 2.8 pts | 0.7 | yes |
| unique | path | 262144 | churn | ordered | baseline | 6 | 968 | 810 | 0.87× [0.76, 0.97] | -15.3% | [-31.3%, -2.7%] | 13.6 pts | 3.2 | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 6 | 45.8 | 38.3 | 0.83× [0.82, 0.84] | -19.8% | [-22.1%, -18.4%] | 1.8 pts | 1.8 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 6 | 373 | 1639 | 4.42× [4.37, 4.46] | +77.4% | [+77.1%, +77.6%] | 0.2 pts | 0.9 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 6 | 561 | 3220 | 5.75× [5.73, 5.77] | +82.6% | [+82.6%, +82.7%] | 0.1 pts | 0.4 | yes |
| unique | str | 4096 | churn | ordered | baseline | 6 | 113 | 87.4 | 0.77× [0.76, 0.78] | -29.9% | [-31.1%, -28.2%] | 1.3 pts | 0.9 | yes |
| unique | str | 4096 | build | ordered | baseline | 6 | 1.47 ms | 1.04 ms | 0.71× [0.70, 0.71] | -41.6% | [-43.0%, -40.4%] | 1.2 pts | 0.5 | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 6 | 59.1 | 50.2 | 0.85× [0.84, 0.85] | -18.3% | [-18.9%, -17.2%] | 0.8 pts | 0.8 | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 470 | 1707 | 3.64× [3.61, 3.66] | +72.5% | [+72.3%, +72.7%] | 0.2 pts | 0.7 | yes |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 2574 | 14.2 µs | 5.49× [5.27, 5.65] | +81.8% | [+81.0%, +82.3%] | 0.6 pts | 2.3 | yes |
| unique | str | 16384 | churn | ordered | baseline | 6 | 139 | 113 | 0.82× [0.80, 0.84] | -22.4% | [-25.6%, -18.9%] | 3.2 pts | 1.4 | no |
| unique | str | 16384 | build | ordered | baseline | 6 | 7.14 ms | 4.95 ms | 0.69× [0.68, 0.70] | -44.8% | [-46.1%, -42.4%] | 1.8 pts | 1.8 | yes |
| unique | str | 65536 | valuesFor | ordered | baseline | 6 | 89.4 | 70.1 | 0.78× [0.78, 0.79] | -27.4% | [-28.4%, -26.1%] | 1.1 pts | 0.5 | yes |
| unique | str | 65536 | valuesBetween | ordered | baseline | 6 | 669 | 1824 | 2.73× [2.69, 2.75] | +63.4% | [+62.9%, +63.7%] | 0.4 pts | 0.6 | yes |
| unique | str | 65536 | prefix | ordered | baseline | 6 | 15.7 µs | 62.4 µs | 3.97× [3.89, 4.06] | +74.8% | [+74.3%, +75.4%] | 0.5 pts | 1.4 | yes |
| unique | str | 65536 | churn | ordered | baseline | 6 | 195 | 179 | 0.93× [0.90, 0.99] | -8.0% | [-11.2%, -0.6%] | 5.1 pts | 0.9 | no |
| unique | str | 65536 | build | ordered | baseline | 6 | 45.63 ms | 31.35 ms | 0.69× [0.67, 0.70] | -45.5% | [-49.4%, -43.7%] | 2.8 pts | 1.6 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 6 | 170 | 219 | 1.39× [1.15, 1.51] | +28.1% | [+12.7%, +33.7%] | 10.0 pts | 1.8 | no |
| unique | str | 262144 | valuesBetween | ordered | baseline | 6 | 1083 | 4415 | 4.04× [3.92, 4.13] | +75.3% | [+74.5%, +75.8%] | 0.6 pts | 1.3 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 6 | 118.0 µs | 639.0 µs | 5.43× [5.23, 5.74] | +81.6% | [+80.9%, +82.6%] | 0.8 pts | 1.1 | yes |
| unique | str | 262144 | churn | ordered | baseline | 6 | 383 | 412 | 1.19× [1.04, 1.36] | +16.3% | [+4.0%, +26.5%] | 10.7 pts | 2.4 | no |
| unique | str | 1048576 | valuesFor | ordered | baseline | 6 | 347 | 348 | 0.99× [0.94, 1.08] | -1.2% | [-6.7%, +7.7%] | 6.8 pts | 2.5 | no |
| unique | str | 1048576 | valuesBetween | ordered | baseline | 6 | 1773 | 6673 | 3.76× [3.66, 3.86] | +73.4% | [+72.7%, +74.1%] | 0.7 pts | 2.1 | yes |
| unique | str | 1048576 | prefix | ordered | baseline | 6 | 896.3 µs | 4.21 ms | 4.70× [4.60, 4.80] | +78.7% | [+78.2%, +79.2%] | 0.4 pts | 2.6 | yes |
| unique | str | 1048576 | churn | ordered | baseline | 6 | 673 | 548 | 0.83× [0.80, 0.88] | -19.9% | [-25.7%, -13.6%] | 5.8 pts | 2.2 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 5 | 54.3 | 48.3 | 0.89× [0.88, 0.90] | -12.4% | [-13.6%, -11.4%] | 0.9 pts | 0.8 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 5 | 469 | 1828 | 3.90× [3.83, 3.97] | +74.4% | [+73.9%, +74.8%] | 0.4 pts | 1.4 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 5 | 107 | 220 | 2.05× [2.03, 2.06] | +51.2% | [+50.7%, +51.4%] | 0.3 pts | 0.6 | yes |
| unique | street | 4096 | churn | ordered | baseline | 5 | 123 | 111 | 0.89× [0.88, 0.91] | -12.5% | [-13.1%, -10.4%] | 1.1 pts | 0.8 | yes |
| unique | street | 4096 | build | ordered | baseline | 5 | 1.87 ms | 1.23 ms | 0.66× [0.66, 0.66] | -51.2% | [-52.2%, -50.7%] | 0.6 pts | 0.5 | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 6 | 70.9 | 70.2 | 0.99× [0.98, 1.00] | -1.3% | [-2.5%, -0.5%] | 1.0 pts | 1.0 | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 6 | 567 | 2214 | 3.90× [3.87, 3.92] | +74.4% | [+74.2%, +74.5%] | 0.2 pts | 0.7 | yes |
| unique | street | 16384 | prefix | ordered | baseline | 6 | 204 | 723 | 3.57× [3.47, 3.63] | +72.0% | [+71.2%, +72.4%] | 0.6 pts | 1.5 | yes |
| unique | street | 16384 | churn | ordered | baseline | 6 | 156 | 151 | 0.97× [0.95, 0.98] | -3.3% | [-5.2%, -1.5%] | 1.7 pts | 1.0 | yes |
| unique | street | 16384 | build | ordered | baseline | 6 | 8.79 ms | 6.27 ms | 0.71× [0.71, 0.72] | -40.4% | [-41.5%, -39.1%] | 1.2 pts | 1.5 | yes |
| unique | street | 65536 | valuesFor | ordered | baseline | 6 | 101 | 98.5 | 0.97× [0.90, 1.13] | -2.7% | [-11.7%, +11.8%] | 11.2 pts | 3.8 | no |
| unique | street | 65536 | valuesBetween | ordered | baseline | 6 | 692 | 2401 | 3.46× [3.32, 3.56] | +71.1% | [+69.9%, +71.9%] | 0.9 pts | 1.9 | yes |
| unique | street | 65536 | prefix | ordered | baseline | 6 | 658 | 3153 | 4.75× [4.59, 4.93] | +79.0% | [+78.2%, +79.7%] | 0.7 pts | 1.4 | yes |
| unique | street | 65536 | churn | ordered | baseline | 6 | 243 | 252 | 1.03× [0.99, 1.07] | +2.7% | [-1.0%, +6.7%] | 3.7 pts | 0.6 | no |
| unique | street | 65536 | build | ordered | baseline | 6 | 48.04 ms | 39.60 ms | 0.83× [0.81, 0.84] | -20.8% | [-22.7%, -19.3%] | 1.7 pts | 1.7 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 18.0 | 15.7 | 0.87× [0.86, 0.89] | -14.5% | [-16.4%, -12.9%] | 1.7 pts | 0.4 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 240 | 1044 | 4.36× [4.31, 4.40] | +77.1% | [+76.8%, +77.3%] | 0.2 pts | 2.0 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 45.1 | 42.2 | 0.94× [0.93, 0.95] | -6.3% | [-7.6%, -5.2%] | 1.2 pts | 1.3 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 598.0 µs | 593.0 µs | 0.99× [0.98, 1.00] | -1.1% | [-1.9%, -0.2%] | 0.8 pts | 0.7 | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 5 | 31.5 | 22.9 | 0.72× [0.71, 0.74] | -38.1% | [-41.3%, -35.2%] | 2.5 pts | 2.6 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 5 | 288 | 1613 | 5.59× [5.53, 5.65] | +82.1% | [+81.9%, +82.3%] | 0.2 pts | 1.4 | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 5 | 65.7 | 45.5 | 0.69× [0.68, 0.70] | -45.9% | [-48.0%, -42.5%] | 2.2 pts | 1.1 | yes |
| unique | u64 | 16384 | build | ordered | baseline | 5 | 3.41 ms | 2.68 ms | 0.78× [0.78, 0.79] | -27.5% | [-28.3%, -26.3%] | 0.8 pts | 0.7 | yes |
| unique | u64 | 65536 | valuesFor | ordered | baseline | 6 | 36.7 | 25.3 | 0.69× [0.68, 0.71] | -44.6% | [-47.9%, -40.3%] | 3.6 pts | 2.5 | yes |
| unique | u64 | 65536 | valuesBetween | ordered | baseline | 6 | 285 | 1969 | 6.89× [6.86, 6.95] | +85.5% | [+85.4%, +85.6%] | 0.1 pts | 1.0 | yes |
| unique | u64 | 65536 | churn | ordered | baseline | 6 | 80.1 | 79.0 | 0.98× [0.96, 1.00] | -2.1% | [-4.2%, +0.3%] | 2.1 pts | 1.1 | no |
| unique | u64 | 65536 | build | ordered | baseline | 6 | 15.65 ms | 14.39 ms | 0.92× [0.88, 0.97] | -8.9% | [-13.5%, -3.1%] | 5.0 pts | 4.5 | no |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 6 | 47.0 | 79.6 | 1.65× [1.58, 1.77] | +39.6% | [+36.7%, +43.5%] | 3.2 pts | 1.3 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 6 | 381 | 2495 | 6.47× [6.17, 6.65] | +84.5% | [+83.8%, +85.0%] | 0.6 pts | 1.1 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 6 | 149 | 216 | 1.51× [1.23, 1.81] | +33.9% | [+18.7%, +44.8%] | 12.4 pts | 2.7 | no |
| unique | u64 | 1048576 | valuesFor | ordered | baseline | 6 | 134 | 127 | 0.95× [0.94, 0.98] | -4.7% | [-6.8%, -2.0%] | 2.3 pts | 0.6 | no |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 6 | 763 | 3172 | 4.14× [4.01, 4.26] | +75.9% | [+75.0%, +76.5%] | 0.7 pts | 2.0 | yes |
| unique | u64 | 1048576 | churn | ordered | baseline | 6 | 311 | 370 | 1.22× [0.89, 1.59] | +17.7% | [-12.8%, +37.1%] | 23.8 pts | 5.4 | no |
| unique | url | 4096 | valuesFor | ordered | baseline | 5 | 69.1 | 60.9 | 0.88× [0.87, 0.90] | -13.9% | [-15.2%, -11.6%] | 1.5 pts | 1.8 | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 5 | 536 | 1722 | 3.21× [3.19, 3.25] | +68.8% | [+68.6%, +69.2%] | 0.2 pts | 0.5 | yes |
| unique | url | 4096 | prefix | ordered | baseline | 5 | 170 | 244 | 1.44× [1.42, 1.46] | +30.3% | [+29.4%, +31.4%] | 0.8 pts | 1.0 | yes |
| unique | url | 4096 | churn | ordered | baseline | 5 | 152 | 136 | 0.89× [0.89, 0.90] | -11.7% | [-12.3%, -10.8%] | 0.6 pts | 0.4 | yes |
| unique | url | 4096 | build | ordered | baseline | 5 | 2.26 ms | 1.57 ms | 0.70× [0.67, 0.72] | -43.8% | [-48.5%, -39.8%] | 3.5 pts | 1.2 | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 81.9 | 75.3 | 0.92× [0.88, 0.95] | -8.8% | [-13.4%, -5.6%] | 3.7 pts | 3.6 | no |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 575 | 1936 | 3.36× [3.27, 3.42] | +70.2% | [+69.5%, +70.7%] | 0.6 pts | 1.2 | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 255 | 654 | 2.57× [2.53, 2.58] | +61.1% | [+60.5%, +61.3%] | 0.4 pts | 0.6 | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 187 | 167 | 0.88× [0.86, 0.92] | -13.1% | [-15.7%, -8.6%] | 3.4 pts | 1.7 | no |
| unique | url | 16384 | build | ordered | baseline | 6 | 9.52 ms | 7.46 ms | 0.78× [0.77, 0.79] | -28.3% | [-29.2%, -26.8%] | 1.2 pts | 1.3 | yes |
| unique | url | 65536 | valuesFor | ordered | baseline | 6 | 178 | 145 | 0.85× [0.73, 1.08] | -17.7% | [-36.7%, +7.4%] | 21.0 pts | 3.3 | no |
| unique | url | 65536 | valuesBetween | ordered | baseline | 6 | 864 | 2364 | 2.74× [2.69, 2.78] | +63.5% | [+62.8%, +64.0%] | 0.6 pts | 0.7 | yes |
| unique | url | 65536 | prefix | ordered | baseline | 6 | 776 | 2571 | 3.30× [3.23, 3.37] | +69.7% | [+69.1%, +70.3%] | 0.6 pts | 1.1 | yes |
| unique | url | 65536 | churn | ordered | baseline | 6 | 295 | 304 | 1.00× [0.94, 1.12] | +0.3% | [-6.1%, +10.6%] | 8.0 pts | 1.8 | no |
| unique | url | 65536 | build | ordered | baseline | 6 | 67.52 ms | 56.40 ms | 0.82× [0.78, 0.86] | -22.3% | [-28.2%, -15.7%] | 6.0 pts | 3.1 | no |
| unique | url | 262144 | valuesFor | ordered | baseline | 6 | 333 | 355 | 1.06× [0.98, 1.14] | +5.6% | [-2.5%, +12.7%] | 7.2 pts | 2.3 | no |
| unique | url | 262144 | valuesBetween | ordered | baseline | 6 | 1601 | 4978 | 3.16× [3.08, 3.23] | +68.3% | [+67.5%, +69.0%] | 0.7 pts | 1.0 | yes |
| unique | url | 262144 | prefix | ordered | baseline | 6 | 4733 | 20.8 µs | 4.37× [4.30, 4.48] | +77.1% | [+76.7%, +77.7%] | 0.4 pts | 0.4 | yes |
| unique | url | 262144 | churn | ordered | baseline | 6 | 612 | 613 | 1.02× [0.89, 1.13] | +2.4% | [-12.8%, +11.8%] | 11.7 pts | 3.6 | no |
| unique | url | 1048576 | valuesFor | ordered | baseline | 6 | 606 | 477 | 0.78× [0.74, 0.87] | -28.9% | [-35.7%, -14.3%] | 10.2 pts | 3.1 | no |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 6 | 2597 | 7345 | 2.84× [2.79, 2.87] | +64.7% | [+64.1%, +65.1%] | 0.5 pts | 0.8 | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 6 | 32.6 µs | 130.4 µs | 4.02× [3.90, 4.12] | +75.1% | [+74.3%, +75.7%] | 0.7 pts | 0.6 | yes |
| unique | url | 1048576 | churn | ordered | baseline | 6 | 920 | 797 | 0.86× [0.76, 1.00] | -16.4% | [-31.3%, +0.3%] | 15.1 pts | 4.9 | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 5 | 44.0 | 31.2 | 0.71× [0.69, 0.73] | -40.6% | [-44.9%, -37.6%] | 2.9 pts | 1.9 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 5 | 396 | 1495 | 3.76× [3.73, 3.79] | +73.4% | [+73.2%, +73.6%] | 0.2 pts | 0.5 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 5 | 74.2 | 66.3 | 0.90× [0.89, 0.90] | -11.7% | [-12.3%, -10.5%] | 0.7 pts | 0.9 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 5 | 96.6 | 80.6 | 0.83× [0.83, 0.84] | -19.9% | [-20.7%, -19.4%] | 0.5 pts | 0.5 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 5 | 1.41 ms | 953.6 µs | 0.68× [0.68, 0.69] | -46.3% | [-48.0%, -45.3%] | 1.1 pts | 0.4 | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 6 | 50.7 | 40.6 | 0.80× [0.80, 0.81] | -24.4% | [-25.1%, -23.9%] | 0.6 pts | 0.6 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 410 | 1630 | 3.99× [3.84, 4.08] | +75.0% | [+74.0%, +75.5%] | 0.7 pts | 1.9 | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 82.8 | 81.2 | 0.98× [0.97, 0.99] | -1.9% | [-3.4%, -1.0%] | 1.2 pts | 1.5 | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 6 | 118 | 103 | 0.88× [0.86, 0.92] | -14.0% | [-16.1%, -8.7%] | 3.5 pts | 1.6 | no |
| unique | uuid | 16384 | build | ordered | baseline | 6 | 5.97 ms | 4.63 ms | 0.78× [0.77, 0.78] | -29.0% | [-30.0%, -27.7%] | 1.1 pts | 0.9 | yes |
| unique | uuid | 65536 | valuesFor | ordered | baseline | 6 | 89.4 | 67.9 | 0.76× [0.70, 0.81] | -31.9% | [-43.4%, -23.9%] | 9.3 pts | 1.3 | no |
| unique | uuid | 65536 | valuesBetween | ordered | baseline | 6 | 703 | 1928 | 2.73× [2.68, 2.81] | +63.4% | [+62.7%, +64.4%] | 0.8 pts | 1.0 | yes |
| unique | uuid | 65536 | prefix | ordered | baseline | 6 | 132 | 142 | 1.08× [0.97, 1.17] | +7.4% | [-3.4%, +14.6%] | 8.6 pts | 2.7 | no |
| unique | uuid | 65536 | churn | ordered | baseline | 6 | 186 | 170 | 0.90× [0.88, 0.94] | -11.2% | [-14.0%, -6.8%] | 3.4 pts | 0.7 | no |
| unique | uuid | 65536 | build | ordered | baseline | 6 | 44.96 ms | 32.44 ms | 0.72× [0.71, 0.73] | -39.7% | [-41.7%, -37.1%] | 2.2 pts | 1.0 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 6 | 255 | 198 | 0.78× [0.69, 0.85] | -28.9% | [-44.0%, -17.6%] | 12.6 pts | 1.4 | no |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 6 | 1314 | 4647 | 3.55× [3.47, 3.62] | +71.9% | [+71.1%, +72.4%] | 0.6 pts | 1.0 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 6 | 274 | 359 | 1.31× [1.19, 1.51] | +24.0% | [+15.7%, +33.6%] | 8.5 pts | 2.5 | no |
| unique | uuid | 262144 | churn | ordered | baseline | 6 | 420 | 408 | 1.04× [0.84, 1.21] | +3.8% | [-19.1%, +17.4%] | 17.4 pts | 4.5 | no |
| unique | uuid | 1048576 | valuesFor | ordered | baseline | 6 | 523 | 329 | 0.66× [0.61, 0.70] | -51.5% | [-64.4%, -42.1%] | 10.6 pts | 2.8 | no |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 6 | 2466 | 6899 | 2.81× [2.75, 2.85] | +64.4% | [+63.6%, +64.9%] | 0.6 pts | 1.1 | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 6 | 560 | 1294 | 2.36× [2.07, 2.53] | +57.7% | [+51.8%, +60.4%] | 4.1 pts | 2.4 | yes |
| unique | uuid | 1048576 | churn | ordered | baseline | 6 | 910 | 591 | 0.66× [0.63, 0.71] | -51.1% | [-59.9%, -41.0%] | 9.0 pts | 2.0 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
