| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 10 | 45.5 | 45.3 | 1.00× [0.98, 1.01] | -0.2% | [-1.6%, +1.2%] | 2.5 pts | 0.8 | yes | no |
| multi | email | 4096 | valuesBetween | ordered | baseline | 10 | 2818 | 2877 | 1.02× [1.00, 1.04] | +1.7% | [-0.3%, +3.6%] | 2.6 pts | 1.0 | yes | no |
| multi | email | 4096 | prefix | ordered | baseline | 10 | 76.2 | 83.9 | 1.10× [1.08, 1.12] | +9.0% | [+7.5%, +10.5%] | 1.6 pts | 1.5 | yes | yes |
| multi | email | 4096 | churn | ordered | baseline | 10 | 69.2 | 67.2 | 0.97× [0.96, 0.98] | -3.0% | [-4.0%, -2.0%] | 1.3 pts | 2.2 | yes | yes |
| multi | email | 4096 | build | ordered | baseline | 10 | 7.07 ms | 6.70 ms | 0.95× [0.95, 0.96] | -5.0% | [-5.5%, -4.4%] | 0.8 pts | 1.3 | yes | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 59.4 | 59.1 | 1.00× [0.99, 1.00] | -0.5% | [-1.1%, +0.1%] | 0.6 pts | 0.9 | yes | no |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3451 | 3466 | 1.01× [1.00, 1.03] | +1.4% | [-0.1%, +3.0%] | 1.5 pts | 2.1 | yes | no |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 90.3 | 95.5 | 1.06× [1.05, 1.07] | +5.8% | [+5.0%, +6.6%] | 0.8 pts | 1.6 | yes | yes |
| multi | email | 16384 | churn | ordered | baseline | 6 | 96.0 | 94.3 | 0.97× [0.97, 0.98] | -2.8% | [-3.4%, -2.3%] | 0.5 pts | 1.1 | yes | yes |
| multi | email | 16384 | build | ordered | baseline | 6 | 34.58 ms | 33.11 ms | 0.96× [0.95, 0.97] | -4.3% | [-5.0%, -3.6%] | 0.7 pts | 1.4 | yes | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 12 | 197 | 197 | 0.99× [0.98, 1.01] | -0.9% | [-2.4%, +0.7%] | 2.5 pts | 4.2 | yes | no |
| multi | email | 262144 | valuesBetween | ordered | baseline | 12 | 7315 | 7349 | 1.00× [1.00, 1.01] | +0.4% | [-0.5%, +1.3%] | 1.3 pts | 1.8 | yes | no |
| multi | email | 262144 | prefix | ordered | baseline | 12 | 266 | 272 | 1.02× [1.01, 1.04] | +2.4% | [+0.9%, +4.0%] | 2.1 pts | 4.0 | yes | yes |
| multi | email | 262144 | churn | ordered | baseline | 12 | 411 | 395 | 0.93× [0.89, 0.98] | -7.4% | [-12.6%, -2.3%] | 5.7 pts | 0.8 | no | yes |
| multi | path | 4096 | valuesFor | ordered | baseline | 12 | 102 | 101 | 1.00× [0.99, 1.01] | +0.0% | [-1.3%, +1.3%] | 1.6 pts | 1.0 | yes | no |
| multi | path | 4096 | valuesBetween | ordered | baseline | 12 | 3821 | 3890 | 1.02× [1.00, 1.05] | +2.2% | [-0.2%, +4.5%] | 2.9 pts | 1.2 | no | no |
| multi | path | 4096 | prefix | ordered | baseline | 12 | 301 | 333 | 1.11× [1.09, 1.13] | +9.7% | [+8.2%, +11.2%] | 1.7 pts | 1.4 | yes | yes |
| multi | path | 4096 | churn | ordered | baseline | 12 | 155 | 149 | 0.96× [0.96, 0.97] | -3.7% | [-4.4%, -3.0%] | 1.2 pts | 1.2 | yes | yes |
| multi | path | 4096 | build | ordered | baseline | 12 | 14.52 ms | 13.73 ms | 0.95× [0.94, 0.96] | -5.4% | [-6.1%, -4.7%] | 0.7 pts | 1.0 | yes | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 10 | 132 | 133 | 1.00× [0.99, 1.01] | +0.0% | [-1.2%, +1.3%] | 1.6 pts | 2.1 | yes | no |
| multi | path | 16384 | valuesBetween | ordered | baseline | 10 | 4545 | 4610 | 1.01× [1.01, 1.02] | +1.1% | [+0.7%, +1.6%] | 0.5 pts | 0.9 | yes | yes |
| multi | path | 16384 | prefix | ordered | baseline | 10 | 699 | 734 | 1.04× [1.02, 1.07] | +4.2% | [+2.2%, +6.1%] | 2.0 pts | 1.3 | yes | yes |
| multi | path | 16384 | churn | ordered | baseline | 10 | 218 | 210 | 0.96× [0.96, 0.97] | -3.7% | [-4.0%, -3.4%] | 1.0 pts | 1.6 | yes | yes |
| multi | path | 16384 | build | ordered | baseline | 10 | 73.19 ms | 70.04 ms | 0.96× [0.95, 0.96] | -4.3% | [-4.9%, -3.7%] | 0.7 pts | 1.1 | yes | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 12 | 399 | 396 | 1.00× [1.00, 1.01] | +0.3% | [-0.4%, +1.1%] | 0.9 pts | 1.1 | yes | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 12 | 10.4 µs | 10.4 µs | 1.01× [1.00, 1.01] | +0.9% | [+0.4%, +1.5%] | 0.8 pts | 1.1 | yes | yes |
| multi | path | 262144 | prefix | ordered | baseline | 12 | 9284 | 9576 | 1.00× [0.99, 1.01] | +0.0% | [-1.2%, +1.3%] | 2.0 pts | 1.0 | yes | no |
| multi | path | 262144 | churn | ordered | baseline | 12 | 698 | 681 | 0.93× [0.89, 0.97] | -7.7% | [-12.0%, -3.3%] | 4.9 pts | 0.8 | no | yes |
| multi | str | 4096 | valuesFor | ordered | baseline | 10 | 59.4 | 59.2 | 1.00× [0.99, 1.02] | +0.4% | [-0.8%, +1.6%] | 2.1 pts | 0.9 | yes | no |
| multi | str | 4096 | valuesBetween | ordered | baseline | 10 | 3294 | 3319 | 1.00× [0.98, 1.02] | +0.3% | [-1.6%, +2.3%] | 3.0 pts | 1.2 | yes | no |
| multi | str | 4096 | prefix | ordered | baseline | 10 | 6820 | 7019 | 1.02× [1.00, 1.05] | +2.4% | [+0.4%, +4.4%] | 3.1 pts | 0.7 | yes | no |
| multi | str | 4096 | churn | ordered | baseline | 10 | 86.7 | 83.2 | 0.96× [0.95, 0.97] | -4.3% | [-5.0%, -3.6%] | 1.4 pts | 1.4 | yes | yes |
| multi | str | 4096 | build | ordered | baseline | 10 | 8.73 ms | 8.26 ms | 0.95× [0.94, 0.95] | -5.5% | [-6.0%, -5.0%] | 0.6 pts | 0.8 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 71.8 | 71.8 | 1.00× [0.99, 1.00] | -0.1% | [-0.7%, +0.4%] | 0.5 pts | 1.0 | yes | no |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3625 | 3624 | 1.00× [1.00, 1.00] | +0.2% | [+0.0%, +0.4%] | 0.2 pts | 0.2 | yes | no |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 33.4 µs | 33.6 µs | 1.00× [0.99, 1.02] | +0.3% | [-1.2%, +1.8%] | 1.4 pts | 1.8 | yes | no |
| multi | str | 16384 | churn | ordered | baseline | 6 | 112 | 108 | 0.96× [0.96, 0.97] | -3.7% | [-4.2%, -3.1%] | 0.6 pts | 1.0 | yes | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 40.73 ms | 39.02 ms | 0.96× [0.95, 0.97] | -4.3% | [-5.1%, -3.5%] | 0.8 pts | 1.7 | yes | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 12 | 221 | 221 | 1.00× [0.98, 1.01] | -0.3% | [-1.9%, +1.3%] | 2.4 pts | 2.8 | yes | no |
| multi | str | 262144 | valuesBetween | ordered | baseline | 12 | 7805 | 7827 | 1.00× [0.99, 1.02] | +0.5% | [-0.6%, +1.6%] | 1.8 pts | 2.5 | yes | no |
| multi | str | 262144 | prefix | ordered | baseline | 12 | 1.25 ms | 1.25 ms | 1.00× [0.99, 1.01] | -0.1% | [-1.4%, +1.1%] | 1.7 pts | 3.2 | yes | no |
| multi | str | 262144 | churn | ordered | baseline | 12 | 429 | 408 | 0.94× [0.92, 0.97] | -6.3% | [-9.0%, -3.6%] | 3.6 pts | 0.5 | no | yes |
| multi | street | 4096 | valuesFor | ordered | baseline | 12 | 48.8 | 47.7 | 0.97× [0.95, 1.00] | -2.7% | [-5.0%, -0.3%] | 2.2 pts | 0.8 | no | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 12 | 2247 | 2282 | 1.02× [0.99, 1.05] | +1.8% | [-1.0%, +4.5%] | 2.7 pts | 1.1 | no | no |
| multi | street | 4096 | prefix | ordered | baseline | 12 | 218 | 237 | 1.09× [1.07, 1.10] | +7.8% | [+6.2%, +9.5%] | 1.8 pts | 1.7 | yes | yes |
| multi | street | 4096 | churn | ordered | baseline | 12 | 96.9 | 94.7 | 0.98× [0.97, 0.99] | -2.1% | [-2.9%, -1.3%] | 0.9 pts | 0.8 | yes | yes |
| multi | street | 4096 | build | ordered | baseline | 12 | 3.77 ms | 3.65 ms | 0.97× [0.96, 0.97] | -3.3% | [-4.0%, -2.6%] | 0.9 pts | 1.0 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 73.2 | 72.5 | 0.99× [0.98, 1.01] | -0.8% | [-2.1%, +0.6%] | 1.3 pts | 1.4 | yes | no |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2835 | 2865 | 1.01× [1.00, 1.02] | +1.0% | [-0.3%, +2.2%] | 1.2 pts | 1.2 | yes | no |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 799 | 828 | 1.04× [1.02, 1.05] | +3.6% | [+2.2%, +5.0%] | 1.4 pts | 0.6 | yes | yes |
| multi | street | 16384 | churn | ordered | baseline | 6 | 132 | 130 | 0.98× [0.97, 0.99] | -1.6% | [-2.7%, -0.6%] | 1.0 pts | 1.5 | yes | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 20.14 ms | 19.54 ms | 0.97× [0.97, 0.98] | -2.8% | [-3.2%, -2.4%] | 0.4 pts | 0.7 | yes | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 12 | 32.0 | 31.6 | 1.01× [0.99, 1.03] | +0.9% | [-0.9%, +2.6%] | 3.4 pts | 0.6 | yes | no |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 12 | 2507 | 2582 | 1.04× [1.02, 1.06] | +3.7% | [+1.9%, +5.6%] | 2.2 pts | 0.6 | yes | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 12 | 51.7 | 49.9 | 0.97× [0.96, 0.97] | -3.6% | [-4.0%, -3.1%] | 0.8 pts | 1.0 | yes | yes |
| multi | u64 | 4096 | build | ordered | baseline | 12 | 5.67 ms | 5.14 ms | 0.90× [0.89, 0.91] | -10.8% | [-11.8%, -9.7%] | 1.7 pts | 1.8 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 46.1 | 45.6 | 0.98× [0.98, 0.99] | -1.6% | [-2.5%, -0.7%] | 0.9 pts | 1.9 | yes | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3445 | 3458 | 1.00× [1.00, 1.01] | +0.5% | [-0.3%, +1.2%] | 0.7 pts | 0.9 | yes | no |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 65.5 | 62.1 | 0.95× [0.94, 0.95] | -5.6% | [-6.1%, -5.2%] | 0.4 pts | 0.8 | yes | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 26.61 ms | 24.73 ms | 0.93× [0.92, 0.94] | -7.3% | [-8.1%, -6.5%] | 0.8 pts | 1.2 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 12 | 121 | 119 | 0.98× [0.95, 1.00] | -2.3% | [-5.0%, +0.4%] | 4.4 pts | 5.8 | no | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 6367 | 6404 | 1.00× [0.98, 1.03] | +0.4% | [-1.7%, +2.5%] | 3.2 pts | 4.3 | no | no |
| multi | u64 | 262144 | churn | ordered | baseline | 12 | 290 | 282 | 0.93× [0.92, 0.95] | -7.2% | [-8.7%, -5.6%] | 2.4 pts | 0.4 | yes | yes |
| multi | url | 4096 | valuesFor | ordered | baseline | 10 | 83.2 | 83.0 | 1.00× [0.99, 1.00] | -0.4% | [-1.0%, +0.3%] | 0.6 pts | 0.4 | yes | no |
| multi | url | 4096 | valuesBetween | ordered | baseline | 10 | 3673 | 3731 | 1.01× [0.99, 1.03] | +1.0% | [-0.9%, +2.9%] | 2.6 pts | 1.2 | yes | no |
| multi | url | 4096 | prefix | ordered | baseline | 10 | 177 | 200 | 1.13× [1.11, 1.14] | +11.2% | [+10.0%, +12.5%] | 1.5 pts | 1.7 | yes | yes |
| multi | url | 4096 | churn | ordered | baseline | 10 | 129 | 124 | 0.96× [0.95, 0.98] | -3.9% | [-5.2%, -2.5%] | 1.4 pts | 1.3 | yes | yes |
| multi | url | 4096 | build | ordered | baseline | 10 | 11.96 ms | 11.34 ms | 0.95× [0.95, 0.95] | -5.2% | [-5.5%, -5.0%] | 0.5 pts | 0.7 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 108 | 107 | 1.00× [0.99, 1.01] | +0.4% | [-0.5%, +1.3%] | 0.9 pts | 1.4 | yes | no |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4342 | 4404 | 1.01× [1.00, 1.02] | +1.2% | [-0.1%, +2.4%] | 1.2 pts | 1.6 | yes | no |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 277 | 301 | 1.08× [1.07, 1.09] | +7.6% | [+6.9%, +8.3%] | 0.7 pts | 0.8 | yes | yes |
| multi | url | 16384 | churn | ordered | baseline | 6 | 177 | 172 | 0.97× [0.96, 0.98] | -3.3% | [-4.2%, -2.4%] | 0.8 pts | 1.5 | yes | yes |
| multi | url | 16384 | build | ordered | baseline | 6 | 60.67 ms | 58.18 ms | 0.96× [0.95, 0.97] | -4.2% | [-5.0%, -3.4%] | 0.8 pts | 1.9 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 12 | 365 | 354 | 1.00× [0.97, 1.02] | -0.3% | [-2.8%, +2.2%] | 2.9 pts | 3.6 | no | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 12 | 10.5 µs | 10.7 µs | 1.01× [1.00, 1.01] | +0.9% | [+0.5%, +1.4%] | 0.8 pts | 1.5 | yes | yes |
| multi | url | 262144 | prefix | ordered | baseline | 12 | 2312 | 2334 | 1.01× [1.00, 1.02] | +1.0% | [+0.1%, +2.0%] | 1.6 pts | 1.3 | yes | no |
| multi | url | 262144 | churn | ordered | baseline | 12 | 639 | 637 | 0.97× [0.94, 1.01] | -2.6% | [-6.4%, +1.1%] | 4.3 pts | 0.7 | no | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 12 | 52.1 | 52.2 | 1.00× [0.98, 1.02] | +0.3% | [-1.6%, +2.3%] | 2.3 pts | 0.8 | yes | no |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 3110 | 3134 | 1.00× [0.98, 1.03] | +0.5% | [-2.3%, +3.2%] | 3.1 pts | 1.2 | no | no |
| multi | uuid | 4096 | prefix | ordered | baseline | 12 | 88.4 | 96.4 | 1.09× [1.07, 1.10] | +8.0% | [+6.7%, +9.3%] | 1.3 pts | 1.6 | yes | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 12 | 80.1 | 77.9 | 0.97× [0.97, 0.98] | -2.9% | [-3.4%, -2.3%] | 0.8 pts | 1.0 | yes | yes |
| multi | uuid | 4096 | build | ordered | baseline | 12 | 7.91 ms | 7.54 ms | 0.96× [0.95, 0.96] | -4.4% | [-4.8%, -4.0%] | 0.8 pts | 1.1 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 63.2 | 63.3 | 1.00× [1.00, 1.01] | +0.4% | [-0.5%, +1.3%] | 0.8 pts | 1.6 | yes | no |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3527 | 3557 | 1.01× [1.01, 1.01] | +1.0% | [+0.6%, +1.5%] | 0.4 pts | 0.7 | yes | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 107 | 113 | 1.06× [1.06, 1.07] | +5.9% | [+5.4%, +6.4%] | 0.5 pts | 1.0 | yes | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 104 | 102 | 0.98× [0.97, 0.99] | -2.0% | [-2.7%, -1.2%] | 0.7 pts | 1.3 | yes | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 36.91 ms | 35.74 ms | 0.97× [0.96, 0.97] | -3.3% | [-4.0%, -2.6%] | 0.7 pts | 1.5 | yes | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 12 | 215 | 214 | 0.99× [0.97, 1.01] | -1.2% | [-3.1%, +0.7%] | 2.3 pts | 3.5 | yes | no |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 8037 | 8076 | 1.01× [1.00, 1.02] | +0.7% | [-0.3%, +1.7%] | 1.1 pts | 1.8 | yes | no |
| multi | uuid | 262144 | prefix | ordered | baseline | 12 | 531 | 538 | 1.01× [1.00, 1.02] | +1.0% | [-0.2%, +2.2%] | 1.2 pts | 2.7 | yes | no |
| multi | uuid | 262144 | churn | ordered | baseline | 12 | 449 | 428 | 0.94× [0.89, 0.99] | -6.8% | [-12.4%, -1.1%] | 6.6 pts | 0.9 | no | yes |
| unique | email | 4096 | valuesFor | ordered | baseline | 12 | 22.3 | 21.9 | 0.97× [0.94, 1.00] | -2.9% | [-5.8%, -0.0%] | 3.1 pts | 1.7 | no | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 12 | 1243 | 1303 | 1.05× [1.03, 1.06] | +4.3% | [+3.2%, +5.5%] | 1.4 pts | 1.4 | yes | yes |
| unique | email | 4096 | prefix | ordered | baseline | 12 | 57.7 | 65.2 | 1.13× [1.12, 1.13] | +11.2% | [+10.6%, +11.9%] | 0.8 pts | 1.5 | yes | yes |
| unique | email | 4096 | churn | ordered | baseline | 12 | 72.8 | 70.3 | 0.97× [0.97, 0.98] | -2.9% | [-3.5%, -2.3%] | 0.9 pts | 1.1 | yes | yes |
| unique | email | 4096 | build | ordered | baseline | 12 | 1.07 ms | 1.02 ms | 0.96× [0.96, 0.97] | -4.0% | [-4.5%, -3.4%] | 0.9 pts | 1.1 | yes | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 37.6 | 37.2 | 0.99× [0.99, 1.00] | -0.7% | [-1.4%, -0.0%] | 0.7 pts | 0.9 | yes | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 1588 | 1635 | 1.03× [1.01, 1.05] | +3.1% | [+1.4%, +4.8%] | 1.6 pts | 1.6 | yes | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 71.8 | 76.6 | 1.07× [1.06, 1.08] | +6.6% | [+5.6%, +7.7%] | 1.0 pts | 2.3 | yes | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 94.1 | 91.9 | 0.98× [0.96, 0.99] | -2.4% | [-3.9%, -1.0%] | 1.4 pts | 1.1 | yes | yes |
| unique | email | 16384 | build | ordered | baseline | 6 | 5.22 ms | 5.03 ms | 0.96× [0.95, 0.97] | -3.9% | [-4.8%, -3.0%] | 0.9 pts | 1.2 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 12 | 134 | 133 | 0.99× [0.97, 1.01] | -1.3% | [-3.3%, +0.7%] | 2.1 pts | 1.9 | no | no |
| unique | email | 262144 | valuesBetween | ordered | baseline | 12 | 3216 | 3246 | 1.01× [1.00, 1.03] | +1.2% | [-0.3%, +2.7%] | 2.3 pts | 1.6 | yes | no |
| unique | email | 262144 | prefix | ordered | baseline | 12 | 179 | 185 | 1.03× [1.02, 1.04] | +2.8% | [+1.8%, +3.7%] | 1.7 pts | 1.9 | yes | yes |
| unique | email | 262144 | churn | ordered | baseline | 12 | 265 | 260 | 0.99× [0.98, 1.01] | -0.8% | [-2.5%, +0.9%] | 2.1 pts | 0.9 | yes | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 8 | 77.0 | 76.5 | 0.98× [0.97, 0.99] | -2.2% | [-3.5%, -0.9%] | 1.4 pts | 0.9 | yes | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 8 | 2024 | 2066 | 1.02× [1.00, 1.04] | +1.6% | [-0.3%, +3.5%] | 2.2 pts | 1.1 | yes | no |
| unique | path | 4096 | prefix | ordered | baseline | 8 | 228 | 258 | 1.14× [1.13, 1.15] | +12.2% | [+11.1%, +13.3%] | 1.2 pts | 1.1 | yes | yes |
| unique | path | 4096 | churn | ordered | baseline | 8 | 204 | 196 | 0.96× [0.95, 0.98] | -3.8% | [-5.1%, -2.5%] | 1.7 pts | 1.4 | yes | yes |
| unique | path | 4096 | build | ordered | baseline | 8 | 2.72 ms | 2.62 ms | 0.97× [0.95, 0.98] | -3.6% | [-5.3%, -1.9%] | 1.8 pts | 2.1 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 6 | 111 | 110 | 1.00× [0.99, 1.01] | -0.2% | [-1.3%, +0.9%] | 1.1 pts | 1.2 | yes | no |
| unique | path | 16384 | valuesBetween | ordered | baseline | 6 | 2532 | 2585 | 1.02× [1.01, 1.04] | +2.3% | [+0.9%, +3.7%] | 1.3 pts | 1.4 | yes | yes |
| unique | path | 16384 | prefix | ordered | baseline | 6 | 469 | 506 | 1.09× [1.08, 1.09] | +8.0% | [+7.3%, +8.6%] | 0.7 pts | 0.4 | yes | yes |
| unique | path | 16384 | churn | ordered | baseline | 6 | 259 | 250 | 0.98× [0.97, 0.98] | -2.5% | [-3.2%, -1.8%] | 0.7 pts | 0.8 | yes | yes |
| unique | path | 16384 | build | ordered | baseline | 6 | 13.39 ms | 12.92 ms | 0.96× [0.95, 0.97] | -4.1% | [-5.2%, -3.1%] | 1.0 pts | 1.3 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 12 | 323 | 322 | 1.00× [0.99, 1.01] | +0.1% | [-0.8%, +0.9%] | 1.7 pts | 1.4 | yes | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 12 | 5885 | 6063 | 1.01× [0.99, 1.04] | +1.3% | [-1.3%, +4.0%] | 3.2 pts | 2.9 | no | no |
| unique | path | 262144 | prefix | ordered | baseline | 12 | 4922 | 4815 | 1.00× [0.98, 1.03] | +0.2% | [-2.6%, +2.9%] | 3.7 pts | 1.3 | no | no |
| unique | path | 262144 | churn | ordered | baseline | 12 | 630 | 617 | 0.97× [0.96, 0.99] | -2.6% | [-4.2%, -0.9%] | 2.2 pts | 0.7 | yes | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 8 | 32.4 | 31.5 | 0.99× [0.97, 1.01] | -0.7% | [-2.6%, +1.2%] | 2.2 pts | 0.5 | yes | no |
| unique | str | 4096 | valuesBetween | ordered | baseline | 8 | 1597 | 1645 | 1.03× [1.02, 1.05] | +3.3% | [+2.1%, +4.5%] | 1.5 pts | 0.9 | yes | yes |
| unique | str | 4096 | prefix | ordered | baseline | 8 | 3060 | 3212 | 1.05× [1.04, 1.06] | +4.7% | [+3.5%, +5.8%] | 1.1 pts | 1.1 | yes | yes |
| unique | str | 4096 | churn | ordered | baseline | 8 | 98.1 | 95.0 | 0.97× [0.95, 0.98] | -3.6% | [-4.7%, -2.4%] | 1.4 pts | 1.6 | yes | yes |
| unique | str | 4096 | build | ordered | baseline | 8 | 1.41 ms | 1.35 ms | 0.96× [0.95, 0.97] | -4.4% | [-5.3%, -3.5%] | 1.1 pts | 0.7 | yes | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 6 | 48.5 | 48.2 | 0.99× [0.98, 1.00] | -0.9% | [-2.1%, +0.4%] | 1.2 pts | 1.3 | yes | no |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 1621 | 1681 | 1.03× [1.02, 1.05] | +3.2% | [+1.8%, +4.6%] | 1.3 pts | 1.4 | yes | yes |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 13.3 µs | 13.5 µs | 1.01× [0.99, 1.03] | +0.8% | [-0.7%, +2.4%] | 1.5 pts | 1.3 | yes | no |
| unique | str | 16384 | churn | ordered | baseline | 6 | 116 | 112 | 0.97× [0.97, 0.98] | -2.6% | [-3.3%, -1.9%] | 0.7 pts | 0.6 | yes | yes |
| unique | str | 16384 | build | ordered | baseline | 6 | 6.56 ms | 6.32 ms | 0.96× [0.95, 0.97] | -4.1% | [-4.8%, -3.3%] | 0.7 pts | 1.0 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 12 | 140 | 136 | 0.98× [0.96, 1.01] | -1.7% | [-4.4%, +1.1%] | 3.5 pts | 2.1 | no | no |
| unique | str | 262144 | valuesBetween | ordered | baseline | 12 | 2645 | 2711 | 1.02× [1.00, 1.03] | +1.7% | [+0.1%, +3.2%] | 2.2 pts | 1.0 | yes | yes |
| unique | str | 262144 | prefix | ordered | baseline | 12 | 380.6 µs | 385.9 µs | 1.01× [1.00, 1.02] | +1.3% | [+0.2%, +2.3%] | 2.5 pts | 1.1 | yes | no |
| unique | str | 262144 | churn | ordered | baseline | 12 | 312 | 306 | 0.97× [0.96, 0.98] | -3.0% | [-4.3%, -1.8%] | 1.8 pts | 0.7 | yes | yes |
| unique | street | 4096 | valuesFor | ordered | baseline | 12 | 40.7 | 39.2 | 0.96× [0.94, 0.99] | -3.7% | [-5.9%, -1.5%] | 2.6 pts | 0.7 | no | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 12 | 1764 | 1801 | 1.01× [0.99, 1.03] | +0.9% | [-0.9%, +2.6%] | 2.3 pts | 1.2 | yes | no |
| unique | street | 4096 | prefix | ordered | baseline | 12 | 193 | 212 | 1.09× [1.08, 1.10] | +8.2% | [+7.5%, +9.0%] | 1.1 pts | 1.1 | yes | yes |
| unique | street | 4096 | churn | ordered | baseline | 12 | 116 | 114 | 0.98× [0.97, 0.99] | -2.1% | [-3.0%, -1.2%] | 1.3 pts | 1.0 | yes | yes |
| unique | street | 4096 | build | ordered | baseline | 12 | 1.65 ms | 1.61 ms | 0.97× [0.97, 0.98] | -2.6% | [-3.0%, -2.2%] | 0.5 pts | 0.8 | yes | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 8 | 66.1 | 65.7 | 1.00× [0.98, 1.01] | -0.4% | [-1.9%, +1.1%] | 1.6 pts | 1.6 | yes | no |
| unique | street | 16384 | valuesBetween | ordered | baseline | 8 | 2155 | 2204 | 1.02× [1.00, 1.04] | +2.0% | [+0.2%, +3.8%] | 1.8 pts | 1.3 | yes | yes |
| unique | street | 16384 | prefix | ordered | baseline | 8 | 620 | 649 | 1.05× [1.04, 1.07] | +5.2% | [+4.0%, +6.4%] | 1.4 pts | 0.7 | yes | yes |
| unique | street | 16384 | churn | ordered | baseline | 8 | 150 | 149 | 0.99× [0.98, 1.00] | -1.1% | [-2.2%, +0.1%] | 1.1 pts | 1.0 | yes | no |
| unique | street | 16384 | build | ordered | baseline | 8 | 8.20 ms | 8.02 ms | 0.98× [0.97, 0.99] | -2.2% | [-2.8%, -1.5%] | 0.7 pts | 1.7 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 11.7 | 13.0 | 1.09× [1.08, 1.11] | +8.7% | [+7.7%, +9.6%] | 0.9 pts | 1.1 | yes | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 236 | 1060 | 4.49× [4.47, 4.52] | +77.7% | [+77.6%, +77.9%] | 0.1 pts | 1.0 | yes | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 45.2 | 46.0 | 1.02× [1.01, 1.03] | +1.9% | [+0.8%, +2.9%] | 1.0 pts | 1.9 | yes | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 704.4 µs | 764.0 µs | 1.09× [1.08, 1.09] | +7.8% | [+7.0%, +8.6%] | 0.8 pts | 0.9 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 12 | 29.6 | 22.8 | 0.77× [0.76, 0.78] | -29.4% | [-31.3%, -27.6%] | 1.9 pts | 2.8 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 12 | 280 | 1548 | 5.53× [5.51, 5.56] | +81.9% | [+81.8%, +82.0%] | 0.1 pts | 1.6 | yes | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 12 | 67.5 | 53.9 | 0.79× [0.78, 0.81] | -26.3% | [-29.0%, -23.5%] | 3.4 pts | 3.1 | no | yes |
| unique | u64 | 16384 | build | ordered | baseline | 12 | 4.00 ms | 3.55 ms | 0.89× [0.88, 0.90] | -12.8% | [-14.2%, -11.3%] | 1.5 pts | 1.0 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 12 | 41.2 | 44.0 | 1.05× [1.01, 1.09] | +4.7% | [+1.0%, +8.3%] | 5.0 pts | 3.6 | no | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 300 | 1428 | 4.74× [4.61, 4.87] | +78.9% | [+78.3%, +79.5%] | 0.7 pts | 1.4 | yes | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 12 | 120 | 177 | 1.39× [1.35, 1.43] | +28.1% | [+26.2%, +30.1%] | 3.0 pts | 0.8 | yes | yes |
| unique | url | 4096 | valuesFor | ordered | baseline | 6 | 60.7 | 59.0 | 0.99× [0.97, 1.00] | -1.4% | [-3.0%, +0.2%] | 1.5 pts | 0.9 | yes | no |
| unique | url | 4096 | valuesBetween | ordered | baseline | 6 | 1894 | 1958 | 1.04× [1.02, 1.05] | +3.5% | [+1.9%, +5.1%] | 1.5 pts | 0.8 | yes | yes |
| unique | url | 4096 | prefix | ordered | baseline | 6 | 144 | 166 | 1.16× [1.15, 1.18] | +14.0% | [+13.0%, +15.1%] | 1.0 pts | 1.1 | yes | yes |
| unique | url | 4096 | churn | ordered | baseline | 6 | 162 | 156 | 0.96× [0.94, 0.97] | -4.3% | [-5.9%, -2.7%] | 1.5 pts | 1.2 | yes | yes |
| unique | url | 4096 | build | ordered | baseline | 6 | 2.26 ms | 2.17 ms | 0.96× [0.94, 0.97] | -4.7% | [-5.9%, -3.4%] | 1.2 pts | 0.7 | yes | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 8 | 87.2 | 86.9 | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.4%] | 1.0 pts | 1.2 | yes | no |
| unique | url | 16384 | valuesBetween | ordered | baseline | 8 | 2366 | 2425 | 1.02× [1.02, 1.03] | +2.0% | [+1.5%, +2.5%] | 0.9 pts | 1.1 | yes | yes |
| unique | url | 16384 | prefix | ordered | baseline | 8 | 216 | 241 | 1.11× [1.10, 1.12] | +10.2% | [+9.5%, +10.9%] | 0.7 pts | 0.8 | yes | yes |
| unique | url | 16384 | churn | ordered | baseline | 8 | 208 | 202 | 0.97× [0.95, 0.98] | -3.5% | [-5.4%, -1.7%] | 1.8 pts | 1.4 | yes | yes |
| unique | url | 16384 | build | ordered | baseline | 8 | 10.66 ms | 10.31 ms | 0.96× [0.96, 0.97] | -3.8% | [-4.2%, -3.4%] | 0.6 pts | 0.9 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 12 | 289 | 289 | 0.99× [0.96, 1.03] | -0.6% | [-4.0%, +2.7%] | 4.0 pts | 3.4 | no | no |
| unique | url | 262144 | valuesBetween | ordered | baseline | 12 | 6196 | 6348 | 1.01× [1.00, 1.02] | +1.0% | [-0.2%, +2.2%] | 1.8 pts | 1.6 | yes | no |
| unique | url | 262144 | prefix | ordered | baseline | 12 | 1297 | 1301 | 1.02× [1.00, 1.03] | +1.5% | [-0.2%, +3.3%] | 2.8 pts | 2.2 | yes | no |
| unique | url | 262144 | churn | ordered | baseline | 12 | 566 | 555 | 0.98× [0.97, 1.00] | -1.9% | [-3.6%, -0.2%] | 1.9 pts | 0.8 | yes | yes |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 12 | 26.8 | 26.4 | 0.99× [0.97, 1.02] | -0.9% | [-3.5%, +1.6%] | 2.4 pts | 1.0 | no | no |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 1426 | 1485 | 1.04× [1.03, 1.05] | +3.7% | [+2.9%, +4.5%] | 1.3 pts | 0.8 | yes | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 12 | 68.7 | 76.2 | 1.11× [1.10, 1.11] | +9.6% | [+9.1%, +10.1%] | 0.6 pts | 1.1 | yes | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 12 | 85.2 | 82.0 | 0.96× [0.95, 0.98] | -3.7% | [-5.0%, -2.4%] | 1.5 pts | 1.3 | yes | yes |
| unique | uuid | 4096 | build | ordered | baseline | 12 | 1.19 ms | 1.13 ms | 0.96× [0.95, 0.97] | -4.5% | [-5.5%, -3.5%] | 1.1 pts | 1.0 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 6 | 41.6 | 41.8 | 1.01× [0.99, 1.03] | +0.7% | [-1.2%, +2.5%] | 1.8 pts | 2.5 | yes | no |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 1573 | 1629 | 1.04× [1.03, 1.05] | +3.8% | [+2.7%, +4.8%] | 1.0 pts | 1.1 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 83.7 | 90.0 | 1.08× [1.07, 1.09] | +7.3% | [+6.4%, +8.2%] | 0.9 pts | 2.4 | yes | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 6 | 101 | 97.5 | 0.98× [0.96, 1.00] | -2.3% | [-4.3%, -0.4%] | 1.9 pts | 1.6 | yes | yes |
| unique | uuid | 16384 | build | ordered | baseline | 6 | 5.62 ms | 5.42 ms | 0.96× [0.96, 0.97] | -3.8% | [-4.4%, -3.2%] | 0.6 pts | 0.6 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 12 | 157 | 154 | 0.97× [0.95, 0.98] | -3.3% | [-4.8%, -1.8%] | 1.6 pts | 1.3 | yes | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 3959 | 3981 | 1.01× [1.00, 1.03] | +1.3% | [-0.1%, +2.7%] | 1.4 pts | 1.0 | yes | no |
| unique | uuid | 262144 | prefix | ordered | baseline | 12 | 318 | 329 | 1.03× [1.01, 1.05] | +2.9% | [+1.1%, +4.7%] | 2.1 pts | 2.1 | yes | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 12 | 318 | 313 | 1.00× [0.98, 1.02] | -0.3% | [-2.4%, +1.8%] | 3.1 pts | 1.1 | no | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.20% does not clear the 1.34% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.59%, 1.20%] includes zero
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.65% does not clear the 2.02% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.25%, 3.56%] includes zero
- multi email n=4096 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.12%, 0.14%] includes zero
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.13%, 2.99%] includes zero
- multi email n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.43%, 0.65%] includes zero
- multi email n=262144 valuesFor: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.40% does not clear the 0.76% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.46%, 1.25%] includes zero
- multi email n=262144 prefix: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.05% does not clear the 0.97% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.26%, 1.35%] includes zero
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.20%, 4.51%] includes zero
- multi path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.02% does not clear the 0.36% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.21%, 1.26%] includes zero
- multi path n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.33% does not clear the 0.55% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.42%, 1.08%] includes zero
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of 0.04% does not clear the 7.12% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of -4.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi path n=262144 prefix: ordered vs baseline: the pooled interval [-1.18%, 1.25%] includes zero
- multi str n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.38% does not clear the 0.85% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.85%, 1.61%] includes zero
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.32% does not clear the 1.20% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.64%, 2.27%] includes zero
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of 2.36% does not clear the 3.40% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.14% does not clear the 0.30% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.67%, 0.39%] includes zero
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.24% does not clear the 0.51% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 prefix: ordered vs baseline: the pooled difference of 0.26% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 prefix: ordered vs baseline: the pooled interval [-1.25%, 1.76%] includes zero
- multi str n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.94%, 1.32%] includes zero
- multi str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.64%, 1.63%] includes zero
- multi str n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesBetween: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=262144 prefix: ordered vs baseline: the pooled difference of -0.14% does not clear the 0.56% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=262144 prefix: ordered vs baseline: the pooled interval [-1.37%, 1.10%] includes zero
- multi str n=262144 prefix: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 prefix: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.97%, 4.51%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.77% does not clear the 0.96% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.10%, 0.56%] includes zero
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.34%, 2.24%] includes zero
- multi street n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.60% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.86% does not clear the 2.49% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.90%, 2.62%] includes zero
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.45% does not clear the 0.63% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.27%, 1.18%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-5.04%, 0.40%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 5.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.66%, 2.55%] includes zero
- multi u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesBetween: ordered vs baseline: 6 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.39% does not clear the 1.20% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.03%, 0.26%] includes zero
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.00% does not clear the 1.59% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.89%, 2.89%] includes zero
- multi url n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.41% does not clear the 0.51% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.51%, 1.34%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.09%, 2.40%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.27% does not clear the 0.51% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.76%, 2.23%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of 1.04% does not clear the 1.58% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=262144 churn: ordered vs baseline: the pooled interval [-6.36%, 1.09%] includes zero
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.34% does not clear the 1.54% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.65%, 2.33%] includes zero
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.48% does not clear the 1.30% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.28%, 3.24%] includes zero
- multi uuid n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.47%, 1.27%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.05%, 0.71%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.31%, 1.72%] includes zero
- multi uuid n=262144 valuesBetween: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 prefix: ordered vs baseline: the pooled interval [-0.17%, 2.25%] includes zero
- multi uuid n=262144 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 prefix: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.34%, 0.69%] includes zero
- unique email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.30%, 2.72%] includes zero
- unique email n=262144 churn: ordered vs baseline: the pooled difference of -0.78% does not clear the 1.69% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=262144 churn: ordered vs baseline: the pooled interval [-2.48%, 0.91%] includes zero
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.35%, 3.54%] includes zero
- unique path n=4096 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.19% does not clear the 0.53% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.32%, 0.94%] includes zero
- unique path n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.06% does not clear the 0.49% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.81%, 0.93%] includes zero
- unique path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.29%, 3.96%] includes zero
- unique path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesBetween: ordered vs baseline: 7 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=262144 prefix: ordered vs baseline: the pooled difference of 0.18% does not clear the 8.40% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=262144 prefix: ordered vs baseline: the pooled interval [-2.55%, 2.91%] includes zero
- unique str n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.70% does not clear the 2.88% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.61%, 1.21%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.12%, 0.36%] includes zero
- unique str n=16384 prefix: ordered vs baseline: the pooled interval [-0.75%, 2.45%] includes zero
- unique str n=262144 valuesFor: ordered vs baseline: the pooled interval [-4.43%, 1.12%] includes zero
- unique str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique str n=262144 prefix: ordered vs baseline: the pooled difference of 1.25% does not clear the 1.88% noise floor, the bound on what the harness reports between identical code in every process
- unique street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.87%, 2.58%] includes zero
- unique street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.41% does not clear the 0.61% noise floor, the bound on what the harness reports between identical code in every process
- unique street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.93%, 1.10%] includes zero
- unique street n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.61% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=16384 churn: ordered vs baseline: the pooled interval [-2.21%, 0.10%] includes zero
- unique u64 n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 churn: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=4096 valuesFor: ordered vs baseline: the pooled difference of -1.42% does not clear the 1.53% noise floor, the bound on what the harness reports between identical code in every process
- unique url n=4096 valuesFor: ordered vs baseline: the pooled interval [-3.01%, 0.18%] includes zero
- unique url n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.69% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.59%, 0.35%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.99%, 2.70%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.17%, 2.22%] includes zero
- unique url n=262144 prefix: ordered vs baseline: the pooled interval [-0.23%, 3.31%] includes zero
- unique url n=262144 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 prefix: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.94% does not clear the 1.36% noise floor, the bound on what the harness reports between identical code in every process
- unique uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-3.46%, 1.58%] includes zero
- unique uuid n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.19%, 2.54%] includes zero
- unique uuid n=16384 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 prefix: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.07%, 2.70%] includes zero
- unique uuid n=262144 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 churn: ordered vs baseline: the pooled difference of -0.28% does not clear the 1.76% noise floor, the bound on what the harness reports between identical code in every process
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-2.39%, 1.83%] includes zero
