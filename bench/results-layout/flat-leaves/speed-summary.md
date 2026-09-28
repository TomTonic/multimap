| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 6 | 49.4 | 49.4 | 1.00× [0.98, 1.01] | -0.4% | [-1.9%, +1.2%] | 1.5 pts | 1.1 | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 6 | 2987 | 2953 | 1.00× [0.98, 1.01] | -0.4% | [-1.8%, +0.9%] | 1.3 pts | 0.7 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 6 | 76.2 | 76.9 | 1.01× [1.00, 1.01] | +0.9% | [+0.5%, +1.4%] | 0.4 pts | 0.4 | yes |
| multi | email | 4096 | churn | ordered | baseline | 6 | 64.6 | 62.3 | 0.96× [0.95, 0.97] | -3.9% | [-4.7%, -3.1%] | 0.8 pts | 0.6 | yes |
| multi | email | 4096 | build | ordered | baseline | 6 | 5.45 ms | 5.25 ms | 0.97× [0.96, 0.98] | -3.3% | [-4.5%, -2.1%] | 1.2 pts | 0.6 | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 10 | 61.1 | 60.1 | 0.99× [0.98, 0.99] | -1.5% | [-2.0%, -0.9%] | 0.8 pts | 0.7 | yes |
| multi | email | 16384 | valuesBetween | ordered | baseline | 10 | 3528 | 3600 | 1.02× [1.01, 1.02] | +1.8% | [+1.3%, +2.3%] | 0.7 pts | 0.8 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 10 | 92.4 | 94.4 | 1.02× [1.02, 1.03] | +2.3% | [+1.8%, +2.8%] | 0.7 pts | 0.7 | yes |
| multi | email | 16384 | churn | ordered | baseline | 10 | 94.8 | 93.8 | 0.99× [0.97, 1.01] | -0.6% | [-2.7%, +1.4%] | 2.9 pts | 1.9 | no |
| multi | email | 16384 | build | ordered | baseline | 10 | 27.80 ms | 26.96 ms | 0.97× [0.97, 0.97] | -2.8% | [-3.1%, -2.6%] | 0.4 pts | 0.6 | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 10 | 254 | 263 | 1.02× [1.00, 1.04] | +1.8% | [-0.1%, +3.8%] | 2.7 pts | 1.8 | yes |
| multi | email | 262144 | valuesBetween | ordered | baseline | 10 | 8302 | 9756 | 1.18× [1.17, 1.19] | +15.2% | [+14.2%, +16.3%] | 1.5 pts | 1.6 | yes |
| multi | email | 262144 | prefix | ordered | baseline | 10 | 299 | 338 | 1.13× [1.12, 1.15] | +11.7% | [+10.4%, +12.9%] | 1.8 pts | 1.7 | yes |
| multi | email | 262144 | churn | ordered | baseline | 10 | 384 | 366 | 0.95× [0.92, 0.98] | -5.3% | [-8.4%, -2.2%] | 4.3 pts | 1.9 | no |
| multi | email | 1048576 | valuesFor | ordered | baseline | 10 | 375 | 377 | 1.00× [0.99, 1.01] | -0.1% | [-1.3%, +1.0%] | 1.6 pts | 1.8 | yes |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 10 | 9775 | 11.6 µs | 1.19× [1.18, 1.19] | +15.8% | [+15.5%, +16.2%] | 0.5 pts | 0.6 | yes |
| multi | email | 1048576 | prefix | ordered | baseline | 10 | 570 | 640 | 1.12× [1.11, 1.13] | +10.6% | [+9.9%, +11.3%] | 1.0 pts | 1.0 | yes |
| multi | email | 1048576 | churn | ordered | baseline | 10 | 579 | 546 | 0.96× [0.93, 1.00] | -3.8% | [-8.0%, +0.3%] | 5.8 pts | 1.3 | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 107 | 108 | 1.01× [1.00, 1.01] | +0.6% | [-0.0%, +1.3%] | 0.6 pts | 0.6 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4181 | 4263 | 1.02× [1.00, 1.03] | +1.5% | [+0.3%, +2.8%] | 1.2 pts | 0.7 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 331 | 337 | 1.02× [1.00, 1.03] | +1.5% | [-0.0%, +3.1%] | 1.5 pts | 0.5 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 156 | 155 | 1.00× [0.98, 1.01] | -0.4% | [-1.5%, +0.8%] | 1.1 pts | 0.7 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 11.64 ms | 11.54 ms | 0.99× [0.99, 0.99] | -1.1% | [-1.5%, -0.6%] | 0.4 pts | 0.4 | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 10 | 137 | 137 | 1.00× [1.00, 1.01] | +0.2% | [-0.2%, +0.6%] | 0.6 pts | 0.7 | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 10 | 4780 | 4838 | 1.01× [1.01, 1.02] | +1.3% | [+0.7%, +1.8%] | 0.8 pts | 0.6 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 10 | 618 | 630 | 1.02× [0.99, 1.04] | +1.9% | [-0.5%, +4.3%] | 3.4 pts | 0.6 | no |
| multi | path | 16384 | churn | ordered | baseline | 10 | 241 | 239 | 0.99× [0.98, 1.00] | -0.7% | [-1.6%, +0.3%] | 1.3 pts | 0.9 | yes |
| multi | path | 16384 | build | ordered | baseline | 10 | 64.59 ms | 64.30 ms | 1.00× [0.99, 1.01] | -0.1% | [-0.8%, +0.6%] | 1.0 pts | 1.1 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 462 | 511 | 1.11× [1.09, 1.12] | +9.5% | [+8.4%, +10.7%] | 1.6 pts | 1.2 | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 12.3 µs | 12.9 µs | 1.04× [1.03, 1.05] | +3.9% | [+2.6%, +5.2%] | 1.8 pts | 1.4 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 12.0 µs | 12.1 µs | 1.01× [1.00, 1.03] | +1.4% | [-0.3%, +3.2%] | 2.4 pts | 0.2 | yes |
| multi | path | 262144 | churn | ordered | baseline | 10 | 814 | 798 | 0.99× [0.97, 1.01] | -1.0% | [-3.2%, +1.1%] | 3.0 pts | 1.1 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 10 | 61.1 | 61.7 | 1.01× [1.00, 1.02] | +1.0% | [+0.2%, +1.9%] | 1.2 pts | 1.3 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 10 | 3551 | 3570 | 1.01× [1.00, 1.02] | +1.0% | [-0.4%, +2.3%] | 1.9 pts | 0.9 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 10 | 8040 | 8106 | 1.01× [0.99, 1.03] | +0.6% | [-1.5%, +2.7%] | 2.9 pts | 1.0 | no |
| multi | str | 4096 | churn | ordered | baseline | 10 | 80.9 | 79.2 | 0.98× [0.97, 0.99] | -1.9% | [-2.7%, -1.2%] | 1.1 pts | 0.8 | yes |
| multi | str | 4096 | build | ordered | baseline | 10 | 6.58 ms | 6.39 ms | 0.97× [0.97, 0.97] | -3.1% | [-3.5%, -2.7%] | 0.6 pts | 0.6 | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 8 | 72.0 | 72.0 | 1.00× [0.98, 1.01] | -0.5% | [-2.4%, +1.4%] | 2.2 pts | 2.4 | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 8 | 3687 | 3763 | 1.02× [1.01, 1.03] | +2.3% | [+1.5%, +3.1%] | 1.0 pts | 1.2 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 8 | 34.8 µs | 35.3 µs | 1.02× [1.01, 1.02] | +1.6% | [+1.3%, +1.9%] | 0.3 pts | 0.2 | yes |
| multi | str | 16384 | churn | ordered | baseline | 8 | 109 | 111 | 1.02× [1.01, 1.04] | +2.1% | [+0.5%, +3.7%] | 1.9 pts | 1.2 | yes |
| multi | str | 16384 | build | ordered | baseline | 8 | 32.73 ms | 32.12 ms | 0.98× [0.98, 0.98] | -1.8% | [-2.1%, -1.6%] | 0.3 pts | 0.4 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 278 | 304 | 1.10× [1.09, 1.11] | +9.1% | [+8.2%, +10.0%] | 1.3 pts | 1.3 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 8781 | 10.4 µs | 1.18× [1.17, 1.18] | +15.1% | [+14.7%, +15.5%] | 0.6 pts | 0.6 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.42 ms | 1.69 ms | 1.19× [1.18, 1.19] | +15.7% | [+15.2%, +16.3%] | 0.7 pts | 0.8 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 419 | 409 | 0.97× [0.95, 1.00] | -2.7% | [-5.0%, -0.4%] | 3.2 pts | 1.6 | no |
| multi | str | 1048576 | valuesFor | ordered | baseline | 10 | 441 | 464 | 1.06× [1.04, 1.07] | +5.3% | [+4.1%, +6.6%] | 1.8 pts | 0.9 | yes |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 10 | 10.3 µs | 12.2 µs | 1.18× [1.17, 1.20] | +15.4% | [+14.4%, +16.5%] | 1.5 pts | 1.3 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 10 | 6.48 ms | 7.60 ms | 1.19× [1.18, 1.19] | +15.6% | [+15.1%, +16.2%] | 0.8 pts | 0.5 | yes |
| multi | str | 1048576 | churn | ordered | baseline | 10 | 703 | 666 | 0.97× [0.93, 1.02] | -3.1% | [-7.7%, +1.6%] | 6.6 pts | 1.9 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 6 | 54.1 | 54.6 | 1.01× [0.99, 1.02] | +0.7% | [-0.5%, +2.0%] | 1.2 pts | 1.0 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 6 | 2421 | 2438 | 1.01× [1.00, 1.02] | +0.9% | [+0.1%, +1.7%] | 0.7 pts | 0.4 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 6 | 256 | 258 | 1.00× [0.99, 1.01] | +0.4% | [-0.6%, +1.4%] | 0.9 pts | 0.3 | yes |
| multi | street | 4096 | churn | ordered | baseline | 6 | 94.2 | 90.9 | 0.96× [0.96, 0.97] | -3.7% | [-4.1%, -3.3%] | 0.4 pts | 0.3 | yes |
| multi | street | 4096 | build | ordered | baseline | 6 | 2.94 ms | 2.80 ms | 0.96× [0.94, 0.97] | -4.6% | [-6.5%, -2.8%] | 1.7 pts | 1.4 | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 76.5 | 75.6 | 0.99× [0.99, 1.00] | -0.7% | [-1.5%, +0.0%] | 0.7 pts | 0.5 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2909 | 2874 | 0.99× [0.98, 1.00] | -1.0% | [-2.1%, +0.0%] | 1.0 pts | 0.8 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 942 | 940 | 1.00× [0.99, 1.01] | -0.0% | [-1.4%, +1.4%] | 1.3 pts | 0.4 | yes |
| multi | street | 16384 | churn | ordered | baseline | 6 | 133 | 130 | 0.98× [0.97, 0.99] | -2.4% | [-3.5%, -1.2%] | 1.1 pts | 0.5 | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 16.50 ms | 15.58 ms | 0.94× [0.94, 0.95] | -6.0% | [-6.5%, -5.5%] | 0.5 pts | 0.3 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 8 | 37.6 | 37.5 | 0.99× [0.98, 1.01] | -0.5% | [-1.5%, +0.5%] | 1.2 pts | 0.8 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2706 | 2689 | 0.99× [0.99, 0.99] | -1.0% | [-1.4%, -0.5%] | 0.5 pts | 0.2 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 8 | 45.7 | 44.1 | 0.97× [0.95, 0.98] | -3.1% | [-4.7%, -1.5%] | 1.9 pts | 1.5 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 8 | 4.17 ms | 4.00 ms | 0.96× [0.96, 0.97] | -4.0% | [-4.4%, -3.5%] | 0.5 pts | 1.0 | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 10 | 45.8 | 45.2 | 0.98× [0.98, 0.99] | -1.6% | [-2.1%, -1.2%] | 0.6 pts | 1.0 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 10 | 3531 | 3560 | 1.01× [1.00, 1.02] | +0.9% | [+0.3%, +1.5%] | 0.9 pts | 0.8 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 10 | 61.4 | 57.8 | 0.95× [0.93, 0.97] | -5.3% | [-7.5%, -3.2%] | 3.0 pts | 1.5 | no |
| multi | u64 | 16384 | build | ordered | baseline | 10 | 20.63 ms | 19.62 ms | 0.95× [0.95, 0.95] | -5.2% | [-5.5%, -4.8%] | 0.5 pts | 0.7 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 140 | 145 | 1.03× [1.02, 1.05] | +3.1% | [+1.7%, +4.5%] | 1.9 pts | 1.5 | yes |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 7093 | 8254 | 1.16× [1.15, 1.17] | +13.6% | [+12.8%, +14.4%] | 1.1 pts | 1.2 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 255 | 241 | 0.93× [0.90, 0.96] | -7.5% | [-11.0%, -4.0%] | 4.9 pts | 2.0 | no |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 207 | 197 | 0.95× [0.95, 0.96] | -5.1% | [-5.7%, -4.4%] | 0.9 pts | 0.7 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 6572 | 7926 | 1.20× [1.19, 1.21] | +17.0% | [+16.3%, +17.7%] | 1.0 pts | 1.3 | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 10 | 379 | 361 | 0.96× [0.92, 1.00] | -4.6% | [-9.0%, -0.1%] | 6.2 pts | 1.1 | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 6 | 81.8 | 82.2 | 1.01× [0.99, 1.02] | +0.7% | [-0.7%, +2.0%] | 1.3 pts | 1.5 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 6 | 3717 | 3764 | 1.01× [1.00, 1.02] | +1.3% | [+0.3%, +2.2%] | 0.9 pts | 0.6 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 6 | 416 | 412 | 1.00× [0.98, 1.01] | -0.4% | [-2.1%, +1.3%] | 1.6 pts | 1.0 | yes |
| multi | url | 4096 | churn | ordered | baseline | 6 | 117 | 118 | 1.00× [0.99, 1.01] | +0.3% | [-0.6%, +1.2%] | 0.9 pts | 0.5 | yes |
| multi | url | 4096 | build | ordered | baseline | 6 | 9.21 ms | 9.17 ms | 0.99× [0.98, 1.01] | -0.5% | [-1.6%, +0.6%] | 1.1 pts | 1.3 | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 8 | 96.5 | 97.4 | 1.01× [1.00, 1.01] | +0.7% | [+0.1%, +1.4%] | 0.8 pts | 0.8 | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 8 | 3970 | 4059 | 1.02× [1.01, 1.03] | +1.8% | [+1.0%, +2.6%] | 0.9 pts | 0.7 | yes |
| multi | url | 16384 | prefix | ordered | baseline | 8 | 1323 | 1343 | 1.02× [1.01, 1.03] | +1.9% | [+1.1%, +2.7%] | 1.0 pts | 0.4 | yes |
| multi | url | 16384 | churn | ordered | baseline | 8 | 170 | 173 | 1.02× [1.00, 1.04] | +1.9% | [-0.0%, +3.8%] | 2.3 pts | 1.6 | yes |
| multi | url | 16384 | build | ordered | baseline | 8 | 45.95 ms | 46.23 ms | 1.01× [1.01, 1.02] | +1.0% | [+0.5%, +1.5%] | 0.6 pts | 0.6 | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 10 | 358 | 418 | 1.17× [1.16, 1.18] | +14.3% | [+13.5%, +15.1%] | 1.2 pts | 1.3 | yes |
| multi | url | 262144 | valuesBetween | ordered | baseline | 10 | 10.4 µs | 11.1 µs | 1.07× [1.05, 1.08] | +6.1% | [+5.0%, +7.3%] | 1.6 pts | 1.3 | yes |
| multi | url | 262144 | prefix | ordered | baseline | 10 | 48.3 µs | 50.6 µs | 1.04× [1.02, 1.06] | +4.1% | [+2.1%, +6.1%] | 2.8 pts | 0.4 | yes |
| multi | url | 262144 | churn | ordered | baseline | 10 | 569 | 554 | 0.96× [0.93, 0.99] | -3.8% | [-7.0%, -0.5%] | 4.5 pts | 2.2 | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 10 | 491 | 537 | 1.10× [1.09, 1.11] | +9.2% | [+8.0%, +10.3%] | 1.6 pts | 2.1 | yes |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 10 | 11.5 µs | 12.6 µs | 1.08× [1.08, 1.09] | +7.8% | [+7.1%, +8.5%] | 1.0 pts | 0.9 | yes |
| multi | url | 1048576 | prefix | ordered | baseline | 10 | 208.0 µs | 221.1 µs | 1.07× [1.06, 1.08] | +6.3% | [+5.4%, +7.3%] | 1.3 pts | 0.1 | yes |
| multi | url | 1048576 | churn | ordered | baseline | 10 | 741 | 731 | 1.00× [0.97, 1.03] | -0.1% | [-3.0%, +2.9%] | 4.1 pts | 1.3 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 54.2 | 54.0 | 1.00× [0.99, 1.00] | -0.2% | [-0.7%, +0.3%] | 0.5 pts | 0.4 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3322 | 3344 | 1.01× [1.00, 1.01] | +0.6% | [+0.3%, +0.9%] | 0.3 pts | 0.1 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 89.8 | 90.8 | 1.01× [1.01, 1.02] | +1.1% | [+0.5%, +1.6%] | 0.5 pts | 0.8 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 76.2 | 74.1 | 0.97× [0.96, 0.98] | -3.2% | [-4.5%, -2.0%] | 1.2 pts | 0.9 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 6.15 ms | 5.96 ms | 0.97× [0.96, 0.97] | -3.2% | [-3.6%, -2.8%] | 0.4 pts | 0.5 | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 10 | 65.4 | 65.3 | 1.00× [0.99, 1.01] | -0.2% | [-1.3%, +1.0%] | 1.6 pts | 1.4 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 10 | 3673 | 3749 | 1.01× [1.00, 1.03] | +1.3% | [-0.4%, +3.0%] | 2.4 pts | 2.1 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 10 | 113 | 115 | 1.02× [1.01, 1.03] | +2.0% | [+1.2%, +2.9%] | 1.2 pts | 0.9 | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 10 | 118 | 125 | 1.04× [1.01, 1.07] | +3.9% | [+1.4%, +6.3%] | 3.4 pts | 1.1 | no |
| multi | uuid | 16384 | build | ordered | baseline | 10 | 36.58 ms | 37.11 ms | 1.01× [0.99, 1.04] | +1.2% | [-1.0%, +3.4%] | 3.1 pts | 1.0 | no |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 333 | 338 | 1.03× [1.01, 1.05] | +2.7% | [+0.7%, +4.7%] | 2.9 pts | 1.4 | no |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 9785 | 11.3 µs | 1.16× [1.16, 1.17] | +14.1% | [+13.6%, +14.6%] | 0.7 pts | 0.5 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 664 | 753 | 1.12× [1.11, 1.14] | +11.0% | [+9.9%, +12.0%] | 1.4 pts | 0.8 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 453 | 463 | 0.99× [0.95, 1.04] | -0.6% | [-4.9%, +3.7%] | 6.0 pts | 1.5 | no |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 10 | 410 | 419 | 1.02× [1.01, 1.03] | +1.9% | [+0.5%, +3.2%] | 1.9 pts | 1.9 | yes |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 10 | 10.4 µs | 12.2 µs | 1.16× [1.14, 1.18] | +13.9% | [+12.6%, +15.3%] | 1.8 pts | 1.8 | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 10 | 1900 | 2234 | 1.17× [1.15, 1.18] | +14.2% | [+13.4%, +15.0%] | 1.2 pts | 1.2 | yes |
| multi | uuid | 1048576 | churn | ordered | baseline | 10 | 642 | 610 | 0.96× [0.93, 0.99] | -4.5% | [-8.0%, -0.9%] | 5.0 pts | 1.0 | no |
| unique | email | 4096 | valuesFor | ordered | baseline | 8 | 26.0 | 25.8 | 0.99× [0.98, 1.01] | -0.9% | [-2.5%, +0.8%] | 2.0 pts | 1.0 | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 8 | 1290 | 1272 | 0.98× [0.97, 1.00] | -1.5% | [-2.6%, -0.5%] | 1.3 pts | 1.0 | yes |
| unique | email | 4096 | prefix | ordered | baseline | 8 | 56.7 | 57.1 | 1.00× [1.00, 1.01] | +0.5% | [-0.2%, +1.1%] | 0.8 pts | 1.1 | yes |
| unique | email | 4096 | churn | ordered | baseline | 8 | 67.3 | 68.3 | 1.02× [1.01, 1.03] | +1.8% | [+1.0%, +2.6%] | 0.9 pts | 1.0 | yes |
| unique | email | 4096 | build | ordered | baseline | 8 | 819.8 µs | 839.3 µs | 1.02× [1.01, 1.03] | +2.1% | [+0.8%, +3.3%] | 1.5 pts | 0.7 | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 8 | 38.1 | 37.7 | 0.99× [0.99, 0.99] | -1.1% | [-1.3%, -1.0%] | 0.2 pts | 0.3 | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 8 | 1640 | 1618 | 0.99× [0.98, 0.99] | -1.5% | [-2.1%, -0.9%] | 0.7 pts | 0.7 | yes |
| unique | email | 16384 | prefix | ordered | baseline | 8 | 70.4 | 70.9 | 1.01× [1.00, 1.03] | +1.3% | [-0.3%, +2.8%] | 1.8 pts | 2.2 | yes |
| unique | email | 16384 | churn | ordered | baseline | 8 | 83.8 | 86.7 | 1.04× [1.03, 1.06] | +3.9% | [+2.5%, +5.4%] | 1.7 pts | 0.7 | yes |
| unique | email | 16384 | build | ordered | baseline | 8 | 4.03 ms | 4.16 ms | 1.03× [1.02, 1.04] | +2.9% | [+2.2%, +3.6%] | 0.8 pts | 0.6 | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 10 | 174 | 178 | 1.03× [1.01, 1.04] | +2.6% | [+1.3%, +3.8%] | 1.8 pts | 1.5 | yes |
| unique | email | 262144 | valuesBetween | ordered | baseline | 10 | 4069 | 4761 | 1.15× [1.13, 1.17] | +13.0% | [+11.9%, +14.2%] | 1.6 pts | 1.5 | yes |
| unique | email | 262144 | prefix | ordered | baseline | 10 | 230 | 254 | 1.11× [1.10, 1.12] | +9.9% | [+8.8%, +10.9%] | 1.4 pts | 1.2 | yes |
| unique | email | 262144 | churn | ordered | baseline | 10 | 242 | 258 | 1.05× [1.04, 1.07] | +5.2% | [+3.5%, +6.9%] | 2.3 pts | 0.2 | yes |
| unique | email | 1048576 | valuesFor | ordered | baseline | 10 | 313 | 300 | 0.96× [0.94, 0.97] | -4.4% | [-6.0%, -2.7%] | 2.3 pts | 1.8 | yes |
| unique | email | 1048576 | valuesBetween | ordered | baseline | 10 | 6392 | 6733 | 1.05× [1.04, 1.06] | +4.7% | [+4.0%, +5.4%] | 1.0 pts | 1.1 | yes |
| unique | email | 1048576 | prefix | ordered | baseline | 10 | 419 | 438 | 1.05× [1.04, 1.05] | +4.4% | [+3.9%, +5.0%] | 0.8 pts | 0.7 | yes |
| unique | email | 1048576 | churn | ordered | baseline | 10 | 428 | 433 | 1.01× [0.99, 1.04] | +1.4% | [-1.3%, +4.1%] | 3.7 pts | 2.5 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 6 | 86.7 | 88.6 | 1.02× [1.01, 1.03] | +1.7% | [+0.6%, +2.9%] | 1.1 pts | 0.8 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 6 | 2226 | 2235 | 1.01× [1.00, 1.02] | +0.6% | [-0.3%, +1.5%] | 0.9 pts | 0.5 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 6 | 254 | 257 | 1.01× [1.00, 1.03] | +1.3% | [+0.0%, +2.7%] | 1.2 pts | 0.7 | yes |
| unique | path | 4096 | churn | ordered | baseline | 6 | 193 | 202 | 1.05× [1.04, 1.06] | +4.6% | [+3.7%, +5.5%] | 0.8 pts | 0.6 | yes |
| unique | path | 4096 | build | ordered | baseline | 6 | 2.10 ms | 2.19 ms | 1.04× [1.03, 1.05] | +3.9% | [+3.3%, +4.6%] | 0.6 pts | 0.5 | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 8 | 114 | 116 | 1.01× [1.00, 1.02] | +1.2% | [+0.4%, +1.9%] | 0.9 pts | 0.9 | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 8 | 2712 | 2664 | 0.98× [0.97, 0.99] | -1.7% | [-2.7%, -0.8%] | 1.1 pts | 1.1 | yes |
| unique | path | 16384 | prefix | ordered | baseline | 8 | 427 | 432 | 1.01× [0.99, 1.02] | +0.7% | [-1.1%, +2.4%] | 2.1 pts | 0.5 | yes |
| unique | path | 16384 | churn | ordered | baseline | 8 | 253 | 268 | 1.05× [1.04, 1.07] | +5.1% | [+3.8%, +6.5%] | 1.6 pts | 0.8 | yes |
| unique | path | 16384 | build | ordered | baseline | 8 | 10.53 ms | 11.02 ms | 1.05× [1.04, 1.06] | +4.7% | [+4.1%, +5.3%] | 0.7 pts | 0.7 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 404 | 450 | 1.13× [1.11, 1.14] | +11.3% | [+10.0%, +12.6%] | 1.8 pts | 1.6 | yes |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 8756 | 8421 | 0.96× [0.94, 0.97] | -4.6% | [-6.1%, -3.1%] | 2.2 pts | 1.6 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 5645 | 5496 | 0.97× [0.94, 1.01] | -2.8% | [-6.6%, +1.0%] | 5.3 pts | 0.4 | no |
| unique | path | 262144 | churn | ordered | baseline | 10 | 587 | 632 | 1.07× [1.05, 1.08] | +6.2% | [+5.1%, +7.3%] | 1.5 pts | 1.1 | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 10 | 37.0 | 37.3 | 1.01× [1.00, 1.02] | +1.1% | [+0.0%, +2.1%] | 1.5 pts | 0.9 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 10 | 1717 | 1691 | 0.98× [0.97, 0.99] | -1.9% | [-3.2%, -0.6%] | 1.8 pts | 1.2 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 10 | 3377 | 3296 | 0.98× [0.96, 1.00] | -2.0% | [-4.1%, -0.0%] | 2.8 pts | 1.0 | no |
| unique | str | 4096 | churn | ordered | baseline | 10 | 86.7 | 87.2 | 1.00× [1.00, 1.01] | +0.5% | [+0.0%, +0.9%] | 0.6 pts | 0.5 | yes |
| unique | str | 4096 | build | ordered | baseline | 10 | 1.04 ms | 1.05 ms | 1.01× [0.99, 1.02] | +0.8% | [-0.7%, +2.3%] | 2.0 pts | 0.9 | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 6 | 47.9 | 47.5 | 0.99× [0.98, 1.00] | -0.9% | [-1.6%, -0.1%] | 0.7 pts | 0.7 | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 1697 | 1688 | 1.00× [0.99, 1.00] | -0.5% | [-1.4%, +0.4%] | 0.8 pts | 0.8 | yes |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 14.3 µs | 14.3 µs | 1.00× [0.99, 1.02] | +0.4% | [-1.3%, +2.0%] | 1.6 pts | 1.0 | yes |
| unique | str | 16384 | churn | ordered | baseline | 6 | 104 | 105 | 1.02× [1.01, 1.04] | +2.2% | [+1.0%, +3.5%] | 1.2 pts | 0.5 | yes |
| unique | str | 16384 | build | ordered | baseline | 6 | 4.97 ms | 4.99 ms | 1.01× [0.99, 1.02] | +0.5% | [-0.7%, +1.7%] | 1.1 pts | 1.1 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 206 | 223 | 1.07× [1.04, 1.09] | +6.2% | [+4.2%, +8.3%] | 2.8 pts | 2.4 | no |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 4696 | 5606 | 1.17× [1.14, 1.20] | +14.6% | [+12.4%, +16.8%] | 3.1 pts | 3.0 | no |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 779.5 µs | 936.5 µs | 1.19× [1.15, 1.24] | +16.2% | [+13.1%, +19.4%] | 4.4 pts | 1.2 | no |
| unique | str | 262144 | churn | ordered | baseline | 10 | 319 | 337 | 1.04× [1.02, 1.05] | +3.7% | [+2.3%, +5.2%] | 2.1 pts | 1.0 | yes |
| unique | str | 1048576 | valuesFor | ordered | baseline | 10 | 346 | 349 | 1.01× [1.00, 1.02] | +0.6% | [-0.4%, +1.5%] | 1.3 pts | 1.5 | yes |
| unique | str | 1048576 | valuesBetween | ordered | baseline | 10 | 6746 | 7006 | 1.04× [1.04, 1.05] | +4.2% | [+3.4%, +5.0%] | 1.1 pts | 1.2 | yes |
| unique | str | 1048576 | prefix | ordered | baseline | 10 | 4.09 ms | 4.24 ms | 1.04× [1.03, 1.05] | +3.8% | [+3.0%, +4.5%] | 1.1 pts | 1.7 | yes |
| unique | str | 1048576 | churn | ordered | baseline | 10 | 503 | 517 | 1.03× [1.01, 1.04] | +2.5% | [+0.7%, +4.2%] | 2.4 pts | 1.7 | yes |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 47.4 | 48.3 | 1.02× [1.01, 1.03] | +1.7% | [+0.6%, +2.8%] | 1.0 pts | 0.9 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 1923 | 1926 | 1.00× [1.00, 1.01] | +0.2% | [-0.5%, +0.9%] | 0.7 pts | 0.5 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 218 | 218 | 1.00× [0.98, 1.02] | +0.1% | [-1.8%, +1.9%] | 1.7 pts | 0.7 | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 111 | 111 | 1.00× [0.99, 1.01] | +0.2% | [-0.7%, +1.0%] | 0.8 pts | 0.6 | yes |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.30 ms | 1.31 ms | 1.00× [0.99, 1.02] | +0.3% | [-1.1%, +1.6%] | 1.3 pts | 1.0 | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 6 | 70.1 | 69.5 | 0.99× [0.98, 1.00] | -0.9% | [-1.9%, +0.1%] | 1.0 pts | 0.9 | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 6 | 2263 | 2229 | 0.98× [0.97, 0.99] | -2.0% | [-3.2%, -0.7%] | 1.2 pts | 1.2 | yes |
| unique | street | 16384 | prefix | ordered | baseline | 6 | 769 | 760 | 0.98× [0.97, 1.00] | -1.6% | [-3.1%, -0.1%] | 1.5 pts | 0.5 | yes |
| unique | street | 16384 | churn | ordered | baseline | 6 | 141 | 142 | 1.01× [1.00, 1.02] | +0.6% | [-0.4%, +1.7%] | 1.0 pts | 0.5 | yes |
| unique | street | 16384 | build | ordered | baseline | 6 | 6.51 ms | 6.58 ms | 1.01× [1.00, 1.02] | +0.8% | [-0.0%, +1.7%] | 0.8 pts | 0.7 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 14.9 | 15.3 | 1.03× [1.01, 1.05] | +2.7% | [+0.7%, +4.7%] | 1.9 pts | 0.8 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1030 | 1024 | 0.99× [0.98, 1.00] | -0.8% | [-1.5%, -0.1%] | 0.7 pts | 0.5 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 44.0 | 44.3 | 1.00× [1.00, 1.01] | +0.2% | [-0.2%, +0.7%] | 0.5 pts | 0.4 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 615.2 µs | 609.8 µs | 1.00× [0.98, 1.01] | -0.2% | [-1.6%, +1.3%] | 1.4 pts | 0.8 | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 6 | 23.2 | 23.2 | 1.00× [0.98, 1.01] | -0.4% | [-1.7%, +0.9%] | 1.2 pts | 1.5 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 1612 | 1597 | 0.99× [0.98, 1.00] | -0.9% | [-1.6%, -0.3%] | 0.6 pts | 0.5 | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 6 | 46.7 | 48.0 | 1.03× [1.01, 1.05] | +2.8% | [+0.9%, +4.7%] | 1.8 pts | 0.7 | yes |
| unique | u64 | 16384 | build | ordered | baseline | 6 | 2.86 ms | 2.89 ms | 1.01× [1.00, 1.02] | +1.1% | [+0.5%, +1.7%] | 0.6 pts | 0.4 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 10 | 63.8 | 82.8 | 1.28× [1.22, 1.33] | +21.7% | [+18.4%, +25.0%] | 4.7 pts | 2.5 | no |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 2295 | 3210 | 1.39× [1.36, 1.43] | +28.3% | [+26.4%, +30.1%] | 2.6 pts | 1.8 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 10 | 166 | 179 | 1.08× [1.06, 1.09] | +7.3% | [+6.0%, +8.5%] | 1.7 pts | 1.0 | yes |
| unique | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 136 | 137 | 1.01× [1.00, 1.01] | +0.5% | [-0.4%, +1.5%] | 1.3 pts | 1.6 | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 3141 | 3186 | 1.02× [1.01, 1.03] | +2.2% | [+1.1%, +3.2%] | 1.5 pts | 2.0 | yes |
| unique | u64 | 1048576 | churn | ordered | baseline | 10 | 300 | 295 | 0.98× [0.95, 1.01] | -2.5% | [-5.5%, +0.5%] | 4.3 pts | 3.7 | no |
| unique | url | 4096 | valuesFor | ordered | baseline | 8 | 58.8 | 61.0 | 1.04× [1.03, 1.04] | +3.6% | [+3.1%, +4.1%] | 0.6 pts | 0.6 | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 8 | 1856 | 1843 | 1.00× [0.99, 1.01] | -0.2% | [-1.2%, +0.9%] | 1.3 pts | 0.7 | yes |
| unique | url | 4096 | prefix | ordered | baseline | 8 | 249 | 250 | 1.00× [1.00, 1.01] | +0.1% | [-0.3%, +0.6%] | 0.5 pts | 0.5 | yes |
| unique | url | 4096 | churn | ordered | baseline | 8 | 132 | 137 | 1.04× [1.03, 1.04] | +3.6% | [+3.1%, +4.0%] | 0.6 pts | 0.4 | yes |
| unique | url | 4096 | build | ordered | baseline | 8 | 1.57 ms | 1.63 ms | 1.03× [1.01, 1.04] | +2.8% | [+1.4%, +4.2%] | 1.7 pts | 0.6 | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 73.9 | 76.1 | 1.03× [1.02, 1.03] | +2.5% | [+2.1%, +2.9%] | 0.4 pts | 0.5 | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 1982 | 1959 | 0.99× [0.97, 1.00] | -1.5% | [-2.8%, -0.2%] | 1.2 pts | 1.0 | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 674 | 667 | 0.99× [0.98, 1.00] | -0.9% | [-2.3%, +0.5%] | 1.3 pts | 0.8 | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 160 | 174 | 1.09× [1.06, 1.11] | +7.9% | [+6.0%, +9.8%] | 1.8 pts | 0.7 | yes |
| unique | url | 16384 | build | ordered | baseline | 6 | 7.32 ms | 7.59 ms | 1.04× [1.03, 1.06] | +4.1% | [+2.9%, +5.3%] | 1.2 pts | 0.8 | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 10 | 303 | 359 | 1.18× [1.16, 1.19] | +14.9% | [+14.0%, +15.9%] | 1.3 pts | 1.4 | yes |
| unique | url | 262144 | valuesBetween | ordered | baseline | 10 | 6532 | 6198 | 0.95× [0.94, 0.96] | -5.3% | [-5.9%, -4.6%] | 0.9 pts | 0.7 | yes |
| unique | url | 262144 | prefix | ordered | baseline | 10 | 30.1 µs | 27.1 µs | 0.91× [0.89, 0.93] | -10.1% | [-12.5%, -7.8%] | 3.3 pts | 0.5 | no |
| unique | url | 262144 | churn | ordered | baseline | 10 | 481 | 522 | 1.08× [1.07, 1.10] | +7.7% | [+6.5%, +8.8%] | 1.6 pts | 1.3 | yes |
| unique | url | 1048576 | valuesFor | ordered | baseline | 10 | 461 | 498 | 1.10× [1.09, 1.11] | +8.9% | [+7.9%, +9.8%] | 1.3 pts | 0.8 | yes |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 10 | 8141 | 7801 | 0.96× [0.95, 0.97] | -3.8% | [-4.8%, -2.8%] | 1.4 pts | 1.2 | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 10 | 139.1 µs | 130.8 µs | 0.95× [0.93, 0.97] | -5.5% | [-8.0%, -3.0%] | 3.5 pts | 0.2 | no |
| unique | url | 1048576 | churn | ordered | baseline | 10 | 684 | 702 | 1.03× [1.02, 1.04] | +2.9% | [+1.5%, +4.2%] | 1.8 pts | 1.2 | yes |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 6 | 29.1 | 29.7 | 1.02× [1.01, 1.03] | +1.9% | [+1.0%, +2.9%] | 0.9 pts | 0.7 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1531 | 1507 | 0.98× [0.98, 0.99] | -1.7% | [-2.3%, -1.1%] | 0.6 pts | 0.4 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 68.4 | 68.5 | 1.00× [1.00, 1.01] | +0.4% | [+0.0%, +0.8%] | 0.4 pts | 0.4 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 6 | 76.9 | 81.1 | 1.06× [1.05, 1.06] | +5.3% | [+4.6%, +6.0%] | 0.7 pts | 0.7 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 6 | 913.8 µs | 975.5 µs | 1.06× [1.05, 1.08] | +6.0% | [+4.5%, +7.6%] | 1.4 pts | 0.8 | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 8 | 39.6 | 38.5 | 0.97× [0.96, 0.98] | -3.0% | [-3.7%, -2.3%] | 0.8 pts | 0.9 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 8 | 1631 | 1625 | 0.99× [0.98, 1.01] | -0.6% | [-2.0%, +0.9%] | 1.7 pts | 1.5 | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 8 | 83.3 | 83.7 | 1.00× [1.00, 1.01] | +0.4% | [-0.1%, +0.8%] | 0.5 pts | 0.6 | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 8 | 90.4 | 97.3 | 1.08× [1.06, 1.10] | +7.6% | [+5.7%, +9.4%] | 2.2 pts | 1.2 | yes |
| unique | uuid | 16384 | build | ordered | baseline | 8 | 4.39 ms | 4.68 ms | 1.07× [1.06, 1.07] | +6.3% | [+5.7%, +6.9%] | 0.7 pts | 0.6 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 221 | 218 | 0.99× [0.96, 1.01] | -1.4% | [-3.7%, +1.0%] | 3.2 pts | 2.3 | no |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 4553 | 5378 | 1.18× [1.16, 1.21] | +15.5% | [+13.7%, +17.4%] | 2.6 pts | 2.4 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 363 | 396 | 1.09× [1.08, 1.10] | +8.2% | [+7.3%, +9.1%] | 1.3 pts | 1.5 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 300 | 316 | 1.05× [1.03, 1.07] | +4.5% | [+2.6%, +6.4%] | 2.6 pts | 1.6 | yes |
| unique | uuid | 1048576 | valuesFor | ordered | baseline | 10 | 342 | 333 | 0.98× [0.97, 1.00] | -1.8% | [-3.2%, -0.4%] | 2.0 pts | 1.4 | yes |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 10 | 6701 | 7197 | 1.07× [1.06, 1.08] | +6.3% | [+5.4%, +7.2%] | 1.2 pts | 1.3 | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 10 | 1262 | 1352 | 1.08× [1.06, 1.09] | +7.0% | [+6.0%, +8.0%] | 1.4 pts | 1.3 | yes |
| unique | uuid | 1048576 | churn | ordered | baseline | 10 | 495 | 508 | 1.02× [0.99, 1.06] | +2.4% | [-0.9%, +5.7%] | 4.6 pts | 3.0 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.38% does not clear the 0.70% median noise floor of the processes
- multi email n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.91%, 1.15%] includes zero
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.43% does not clear the 2.29% median noise floor of the processes
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.81%, 0.94%] includes zero
- multi email n=4096 prefix: ordered vs baseline: the pooled difference of 0.94% does not clear the 1.25% median noise floor of the processes
- multi email n=16384 churn: ordered vs baseline: the pooled difference of -0.61% does not clear the 1.93% median noise floor of the processes
- multi email n=16384 churn: ordered vs baseline: the pooled interval [-2.66%, 1.43%] includes zero
- multi email n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.08%, 3.78%] includes zero
- multi email n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.14% does not clear the 1.03% median noise floor of the processes
- multi email n=1048576 valuesFor: ordered vs baseline: the pooled interval [-1.31%, 1.02%] includes zero
- multi email n=1048576 churn: ordered vs baseline: the pooled interval [-7.99%, 0.32%] includes zero
- multi email n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.63% does not clear the 0.94% median noise floor of the processes
- multi path n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.01%, 1.26%] includes zero
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.52% does not clear the 1.63% median noise floor of the processes
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of 1.55% does not clear the 6.50% median noise floor of the processes
- multi path n=4096 prefix: ordered vs baseline: the pooled interval [-0.02%, 3.11%] includes zero
- multi path n=4096 churn: ordered vs baseline: the pooled difference of -0.38% does not clear the 1.64% median noise floor of the processes
- multi path n=4096 churn: ordered vs baseline: the pooled interval [-1.54%, 0.79%] includes zero
- multi path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.21% does not clear the 1.32% median noise floor of the processes
- multi path n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.21%, 0.62%] includes zero
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 1.27% does not clear the 1.65% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of 1.88% does not clear the 5.45% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-0.53%, 4.29%] includes zero
- multi path n=16384 churn: ordered vs baseline: the pooled difference of -0.66% does not clear the 1.57% median noise floor of the processes
- multi path n=16384 churn: ordered vs baseline: the pooled interval [-1.57%, 0.25%] includes zero
- multi path n=16384 build: ordered vs baseline: the pooled difference of -0.11% does not clear the 1.09% median noise floor of the processes
- multi path n=16384 build: ordered vs baseline: the pooled interval [-0.82%, 0.60%] includes zero
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of 1.42% does not clear the 10.95% median noise floor of the processes
- multi path n=262144 prefix: ordered vs baseline: the pooled interval [-0.31%, 3.15%] includes zero
- multi path n=262144 churn: ordered vs baseline: the pooled difference of -1.05% does not clear the 2.56% median noise floor of the processes
- multi path n=262144 churn: ordered vs baseline: the pooled interval [-3.18%, 1.08%] includes zero
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.97% does not clear the 1.90% median noise floor of the processes
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.37%, 2.31%] includes zero
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of 0.62% does not clear the 2.70% median noise floor of the processes
- multi str n=4096 prefix: ordered vs baseline: the pooled interval [-1.48%, 2.72%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.48% does not clear the 1.58% median noise floor of the processes
- multi str n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.36%, 1.40%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 prefix: ordered vs baseline: the pooled difference of 1.60% does not clear the 1.86% median noise floor of the processes
- multi str n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=1048576 churn: ordered vs baseline: the pooled difference of -3.06% does not clear the 3.82% median noise floor of the processes
- multi str n=1048576 churn: ordered vs baseline: the pooled interval [-7.75%, 1.63%] includes zero
- multi str n=1048576 churn: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.74% does not clear the 1.72% median noise floor of the processes
- multi street n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.55%, 2.03%] includes zero
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.91% does not clear the 1.91% median noise floor of the processes
- multi street n=4096 prefix: ordered vs baseline: the pooled difference of 0.43% does not clear the 2.63% median noise floor of the processes
- multi street n=4096 prefix: ordered vs baseline: the pooled interval [-0.56%, 1.42%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.72% does not clear the 1.84% median noise floor of the processes
- multi street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.49%, 0.05%] includes zero
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.01% does not clear the 1.28% median noise floor of the processes
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-2.07%, 0.04%] includes zero
- multi street n=16384 prefix: ordered vs baseline: the pooled difference of -0.00% does not clear the 5.53% median noise floor of the processes
- multi street n=16384 prefix: ordered vs baseline: the pooled interval [-1.37%, 1.37%] includes zero
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.51% does not clear the 1.74% median noise floor of the processes
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.54%, 0.52%] includes zero
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.96% does not clear the 2.21% median noise floor of the processes
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.90% does not clear the 1.41% median noise floor of the processes
- multi u64 n=262144 churn: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.69%, 2.04%] includes zero
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.25% does not clear the 1.72% median noise floor of the processes
- multi url n=4096 prefix: ordered vs baseline: the pooled difference of -0.38% does not clear the 1.77% median noise floor of the processes
- multi url n=4096 prefix: ordered vs baseline: the pooled interval [-2.09%, 1.33%] includes zero
- multi url n=4096 churn: ordered vs baseline: the pooled difference of 0.31% does not clear the 2.42% median noise floor of the processes
- multi url n=4096 churn: ordered vs baseline: the pooled interval [-0.59%, 1.21%] includes zero
- multi url n=4096 build: ordered vs baseline: the pooled difference of -0.52% does not clear the 1.11% median noise floor of the processes
- multi url n=4096 build: ordered vs baseline: the pooled interval [-1.63%, 0.59%] includes zero
- multi url n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.72% does not clear the 1.04% median noise floor of the processes
- multi url n=16384 prefix: ordered vs baseline: the pooled difference of 1.91% does not clear the 2.23% median noise floor of the processes
- multi url n=16384 churn: ordered vs baseline: the pooled difference of 1.90% does not clear the 2.21% median noise floor of the processes
- multi url n=16384 churn: ordered vs baseline: the pooled interval [-0.03%, 3.83%] includes zero
- multi url n=16384 build: ordered vs baseline: the pooled difference of 0.99% does not clear the 1.00% median noise floor of the processes
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of 4.11% does not clear the 8.41% median noise floor of the processes
- multi url n=262144 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 prefix: ordered vs baseline: the pooled difference of 6.32% does not clear the 34.28% median noise floor of the processes
- multi url n=1048576 churn: ordered vs baseline: the pooled difference of -0.08% does not clear the 2.71% median noise floor of the processes
- multi url n=1048576 churn: ordered vs baseline: the pooled interval [-3.01%, 2.85%] includes zero
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.18% does not clear the 0.56% median noise floor of the processes
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.69%, 0.34%] includes zero
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.59% does not clear the 2.34% median noise floor of the processes
- multi uuid n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.15% does not clear the 1.91% median noise floor of the processes
- multi uuid n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.30%, 0.99%] includes zero
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of 1.29% does not clear the 2.85% median noise floor of the processes
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.44%, 3.03%] includes zero
- multi uuid n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=16384 prefix: ordered vs baseline: the pooled difference of 2.03% does not clear the 2.28% median noise floor of the processes
- multi uuid n=16384 churn: ordered vs baseline: the pooled difference of 3.86% does not clear the 3.95% median noise floor of the processes
- multi uuid n=16384 build: ordered vs baseline: the pooled difference of 1.21% does not clear the 3.75% median noise floor of the processes
- multi uuid n=16384 build: ordered vs baseline: the pooled interval [-0.98%, 3.41%] includes zero
- multi uuid n=262144 churn: ordered vs baseline: the pooled difference of -0.58% does not clear the 3.77% median noise floor of the processes
- multi uuid n=262144 churn: ordered vs baseline: the pooled interval [-4.89%, 3.73%] includes zero
- multi uuid n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=1048576 valuesFor: ordered vs baseline: the pooled difference of 1.86% does not clear the 2.17% median noise floor of the processes
- multi uuid n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.89% does not clear the 2.74% median noise floor of the processes
- unique email n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.54%, 0.77%] includes zero
- unique email n=4096 prefix: ordered vs baseline: the pooled difference of 0.45% does not clear the 1.10% median noise floor of the processes
- unique email n=4096 prefix: ordered vs baseline: the pooled interval [-0.25%, 1.15%] includes zero
- unique email n=4096 build: ordered vs baseline: the pooled difference of 2.06% does not clear the 2.14% median noise floor of the processes
- unique email n=16384 valuesFor: ordered vs baseline: the pooled difference of -1.12% does not clear the 1.15% median noise floor of the processes
- unique email n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.49% does not clear the 1.55% median noise floor of the processes
- unique email n=16384 prefix: ordered vs baseline: the pooled interval [-0.25%, 2.80%] includes zero
- unique email n=16384 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=1048576 churn: ordered vs baseline: the pooled difference of 1.42% does not clear the 1.65% median noise floor of the processes
- unique email n=1048576 churn: ordered vs baseline: the pooled interval [-1.25%, 4.09%] includes zero
- unique email n=1048576 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=1048576 churn: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=4096 valuesFor: ordered vs baseline: the pooled difference of 1.72% does not clear the 1.84% median noise floor of the processes
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.61% does not clear the 2.41% median noise floor of the processes
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.29%, 1.51%] includes zero
- unique path n=4096 prefix: ordered vs baseline: the pooled difference of 1.35% does not clear the 2.44% median noise floor of the processes
- unique path n=16384 prefix: ordered vs baseline: the pooled difference of 0.67% does not clear the 3.96% median noise floor of the processes
- unique path n=16384 prefix: ordered vs baseline: the pooled interval [-1.10%, 2.43%] includes zero
- unique path n=262144 prefix: ordered vs baseline: the pooled difference of -2.78% does not clear the 12.03% median noise floor of the processes
- unique path n=262144 prefix: ordered vs baseline: the pooled interval [-6.59%, 1.04%] includes zero
- unique str n=4096 valuesFor: ordered vs baseline: the pooled difference of 1.06% does not clear the 2.16% median noise floor of the processes
- unique str n=4096 prefix: ordered vs baseline: the pooled difference of -2.03% does not clear the 3.65% median noise floor of the processes
- unique str n=4096 churn: ordered vs baseline: the pooled difference of 0.46% does not clear the 0.72% median noise floor of the processes
- unique str n=4096 build: ordered vs baseline: the pooled difference of 0.81% does not clear the 2.54% median noise floor of the processes
- unique str n=4096 build: ordered vs baseline: the pooled interval [-0.66%, 2.28%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.85% does not clear the 0.96% median noise floor of the processes
- unique str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.49% does not clear the 1.22% median noise floor of the processes
- unique str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.38%, 0.40%] includes zero
- unique str n=16384 prefix: ordered vs baseline: the pooled difference of 0.38% does not clear the 1.89% median noise floor of the processes
- unique str n=16384 prefix: ordered vs baseline: the pooled interval [-1.27%, 2.02%] includes zero
- unique str n=16384 build: ordered vs baseline: the pooled difference of 0.50% does not clear the 1.79% median noise floor of the processes
- unique str n=16384 build: ordered vs baseline: the pooled interval [-0.68%, 1.68%] includes zero
- unique str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.56% does not clear the 1.40% median noise floor of the processes
- unique str n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.38%, 1.50%] includes zero
- unique street n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.23% does not clear the 1.96% median noise floor of the processes
- unique street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.48%, 0.94%] includes zero
- unique street n=4096 prefix: ordered vs baseline: the pooled difference of 0.05% does not clear the 1.77% median noise floor of the processes
- unique street n=4096 prefix: ordered vs baseline: the pooled interval [-1.77%, 1.87%] includes zero
- unique street n=4096 churn: ordered vs baseline: the pooled difference of 0.16% does not clear the 1.51% median noise floor of the processes
- unique street n=4096 churn: ordered vs baseline: the pooled interval [-0.73%, 1.04%] includes zero
- unique street n=4096 build: ordered vs baseline: the pooled difference of 0.26% does not clear the 2.11% median noise floor of the processes
- unique street n=4096 build: ordered vs baseline: the pooled interval [-1.06%, 1.59%] includes zero
- unique street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.90% does not clear the 1.48% median noise floor of the processes
- unique street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.93%, 0.14%] includes zero
- unique street n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.97% does not clear the 2.91% median noise floor of the processes
- unique street n=16384 prefix: ordered vs baseline: the pooled difference of -1.60% does not clear the 4.58% median noise floor of the processes
- unique street n=16384 churn: ordered vs baseline: the pooled difference of 0.62% does not clear the 2.06% median noise floor of the processes
- unique street n=16384 churn: ordered vs baseline: the pooled interval [-0.41%, 1.66%] includes zero
- unique street n=16384 build: ordered vs baseline: the pooled difference of 0.82% does not clear the 1.36% median noise floor of the processes
- unique street n=16384 build: ordered vs baseline: the pooled interval [-0.03%, 1.66%] includes zero
- unique u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 2.71% does not clear the 4.79% median noise floor of the processes
- unique u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.82% does not clear the 1.80% median noise floor of the processes
- unique u64 n=4096 churn: ordered vs baseline: the pooled difference of 0.24% does not clear the 1.28% median noise floor of the processes
- unique u64 n=4096 churn: ordered vs baseline: the pooled interval [-0.25%, 0.74%] includes zero
- unique u64 n=4096 build: ordered vs baseline: the pooled difference of -0.20% does not clear the 3.68% median noise floor of the processes
- unique u64 n=4096 build: ordered vs baseline: the pooled interval [-1.64%, 1.25%] includes zero
- unique u64 n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.40% does not clear the 1.10% median noise floor of the processes
- unique u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.69%, 0.90%] includes zero
- unique u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.93% does not clear the 1.15% median noise floor of the processes
- unique u64 n=16384 build: ordered vs baseline: the pooled difference of 1.07% does not clear the 1.15% median noise floor of the processes
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.53% does not clear the 1.08% median noise floor of the processes
- unique u64 n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.40%, 1.46%] includes zero
- unique u64 n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 churn: ordered vs baseline: the pooled interval [-5.55%, 0.54%] includes zero
- unique u64 n=1048576 churn: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.16% does not clear the 1.72% median noise floor of the processes
- unique url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.21%, 0.89%] includes zero
- unique url n=4096 prefix: ordered vs baseline: the pooled difference of 0.11% does not clear the 1.46% median noise floor of the processes
- unique url n=4096 prefix: ordered vs baseline: the pooled interval [-0.33%, 0.56%] includes zero
- unique url n=4096 build: ordered vs baseline: the pooled difference of 2.84% does not clear the 3.56% median noise floor of the processes
- unique url n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.49% does not clear the 1.55% median noise floor of the processes
- unique url n=16384 prefix: ordered vs baseline: the pooled difference of -0.89% does not clear the 2.47% median noise floor of the processes
- unique url n=16384 prefix: ordered vs baseline: the pooled interval [-2.28%, 0.50%] includes zero
- unique url n=1048576 prefix: ordered vs baseline: the pooled difference of -5.52% does not clear the 17.10% median noise floor of the processes
- unique uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of 1.94% does not clear the 2.06% median noise floor of the processes
- unique uuid n=4096 prefix: ordered vs baseline: the pooled difference of 0.43% does not clear the 1.11% median noise floor of the processes
- unique uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.56% does not clear the 1.30% median noise floor of the processes
- unique uuid n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.99%, 0.88%] includes zero
- unique uuid n=16384 prefix: ordered vs baseline: the pooled difference of 0.38% does not clear the 1.48% median noise floor of the processes
- unique uuid n=16384 prefix: ordered vs baseline: the pooled interval [-0.05%, 0.80%] includes zero
- unique uuid n=262144 valuesFor: ordered vs baseline: the pooled difference of -1.36% does not clear the 1.65% median noise floor of the processes
- unique uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.66%, 0.95%] includes zero
- unique uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=1048576 churn: ordered vs baseline: the pooled interval [-0.91%, 5.71%] includes zero
- unique uuid n=1048576 churn: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 churn: ordered vs baseline: 6 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
