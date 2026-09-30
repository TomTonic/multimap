| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 8 | 51.3 | 50.9 | 1.00× [0.96, 1.04] | -0.2% | [-4.1%, +3.7%] | 2.9 pts | 1.6 | no | no |
| multi | email | 4096 | valuesBetween | ordered | baseline | 8 | 3112 | 3070 | 0.99× [0.96, 1.01] | -1.1% | [-3.6%, +1.4%] | 1.9 pts | 0.6 | no | no |
| multi | email | 4096 | prefix | ordered | baseline | 8 | 80.8 | 78.3 | 0.97× [0.96, 0.97] | -3.4% | [-4.1%, -2.6%] | 0.7 pts | 0.5 | yes | yes |
| multi | email | 4096 | churn | ordered | baseline | 8 | 71.5 | 68.7 | 0.96× [0.96, 0.97] | -4.1% | [-4.7%, -3.5%] | 0.5 pts | 0.4 | yes | yes |
| multi | email | 4096 | build | ordered | baseline | 8 | 7.08 ms | 6.86 ms | 0.97× [0.96, 0.98] | -3.2% | [-4.2%, -2.2%] | 0.9 pts | 0.8 | yes | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 8 | 61.5 | 60.6 | 0.98× [0.96, 1.01] | -1.6% | [-3.9%, +0.7%] | 1.5 pts | 1.3 | no | no |
| multi | email | 16384 | valuesBetween | ordered | baseline | 8 | 3604 | 3523 | 0.98× [0.97, 0.98] | -2.4% | [-3.3%, -1.6%] | 0.7 pts | 0.6 | yes | yes |
| multi | email | 16384 | prefix | ordered | baseline | 8 | 93.4 | 91.4 | 0.98× [0.97, 0.99] | -1.9% | [-3.1%, -0.7%] | 1.3 pts | 1.3 | yes | yes |
| multi | email | 16384 | churn | ordered | baseline | 8 | 109 | 105 | 0.98× [0.97, 0.99] | -2.3% | [-3.4%, -1.3%] | 0.9 pts | 1.0 | yes | yes |
| multi | email | 16384 | build | ordered | baseline | 8 | 34.88 ms | 33.74 ms | 0.97× [0.96, 0.97] | -3.5% | [-4.1%, -2.9%] | 0.7 pts | 1.1 | yes | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 8 | 224 | 229 | 1.02× [1.00, 1.04] | +1.9% | [-0.4%, +4.2%] | 1.9 pts | 2.0 | no | no |
| multi | email | 262144 | valuesBetween | ordered | baseline | 8 | 7949 | 7898 | 0.99× [0.98, 1.00] | -0.8% | [-2.1%, +0.5%] | 1.1 pts | 1.3 | yes | no |
| multi | email | 262144 | prefix | ordered | baseline | 8 | 296 | 296 | 1.00× [0.99, 1.01] | -0.4% | [-1.4%, +0.7%] | 1.0 pts | 1.3 | yes | no |
| multi | email | 262144 | churn | ordered | baseline | 8 | 467 | 446 | 0.94× [0.87, 1.02] | -6.2% | [-14.4%, +2.0%] | 5.7 pts | 0.7 | no | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 8 | 110 | 107 | 0.96× [0.96, 0.97] | -3.7% | [-4.6%, -2.7%] | 1.3 pts | 0.9 | yes | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 8 | 4212 | 4219 | 1.01× [0.99, 1.02] | +0.6% | [-0.9%, +2.2%] | 1.5 pts | 0.6 | yes | no |
| multi | path | 4096 | prefix | ordered | baseline | 8 | 330 | 322 | 1.02× [0.96, 1.07] | +1.5% | [-3.6%, +6.6%] | 3.2 pts | 1.1 | no | no |
| multi | path | 4096 | churn | ordered | baseline | 8 | 161 | 156 | 0.98× [0.96, 0.99] | -2.5% | [-3.7%, -1.3%] | 1.1 pts | 0.7 | yes | yes |
| multi | path | 4096 | build | ordered | baseline | 8 | 14.78 ms | 14.10 ms | 0.95× [0.95, 0.96] | -5.0% | [-5.3%, -4.7%] | 0.3 pts | 0.7 | yes | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 8 | 139 | 137 | 0.98× [0.98, 0.99] | -1.8% | [-2.5%, -1.1%] | 0.7 pts | 0.4 | yes | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 8 | 4736 | 4728 | 1.00× [0.99, 1.00] | -0.4% | [-1.2%, +0.4%] | 0.7 pts | 0.7 | yes | no |
| multi | path | 16384 | prefix | ordered | baseline | 8 | 716 | 758 | 1.02× [0.97, 1.07] | +1.7% | [-2.8%, +6.1%] | 3.0 pts | 0.8 | no | no |
| multi | path | 16384 | churn | ordered | baseline | 8 | 251 | 248 | 1.00× [0.99, 1.02] | +0.3% | [-1.4%, +2.0%] | 1.2 pts | 0.7 | yes | no |
| multi | path | 16384 | build | ordered | baseline | 8 | 76.89 ms | 75.93 ms | 0.98× [0.98, 0.99] | -1.5% | [-2.2%, -0.9%] | 0.7 pts | 1.1 | yes | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 8 | 445 | 448 | 1.02× [1.01, 1.03] | +1.8% | [+0.9%, +2.7%] | 2.6 pts | 1.3 | yes | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 8 | 11.6 µs | 12.0 µs | 1.05× [1.04, 1.06] | +4.7% | [+3.5%, +5.9%] | 2.0 pts | 2.3 | yes | yes |
| multi | path | 262144 | prefix | ordered | baseline | 8 | 9677 | 10.2 µs | 1.05× [1.04, 1.06] | +4.5% | [+3.7%, +5.3%] | 1.9 pts | 2.5 | yes | yes |
| multi | path | 262144 | churn | ordered | baseline | 8 | 804 | 815 | 0.99× [0.96, 1.03] | -0.6% | [-4.2%, +3.0%] | 4.2 pts | 0.8 | no | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 8 | 64.2 | 63.6 | 0.99× [0.97, 1.01] | -0.8% | [-2.6%, +1.1%] | 1.4 pts | 0.7 | yes | no |
| multi | str | 4096 | valuesBetween | ordered | baseline | 8 | 3591 | 3519 | 0.99× [0.98, 1.00] | -1.2% | [-2.4%, -0.1%] | 2.3 pts | 0.7 | yes | no |
| multi | str | 4096 | prefix | ordered | baseline | 8 | 8059 | 8052 | 1.02× [1.00, 1.05] | +2.4% | [+0.0%, +4.8%] | 1.8 pts | 0.4 | no | yes |
| multi | str | 4096 | churn | ordered | baseline | 8 | 88.1 | 85.8 | 0.97× [0.96, 0.98] | -3.1% | [-3.8%, -2.4%] | 1.1 pts | 0.8 | yes | yes |
| multi | str | 4096 | build | ordered | baseline | 8 | 8.67 ms | 8.35 ms | 0.96× [0.95, 0.97] | -3.8% | [-5.0%, -2.6%] | 0.8 pts | 1.5 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 74.3 | 73.0 | 0.99× [0.98, 0.99] | -1.3% | [-2.1%, -0.6%] | 0.7 pts | 0.6 | yes | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3737 | 3651 | 0.98× [0.98, 0.99] | -1.9% | [-2.4%, -1.3%] | 0.4 pts | 0.4 | yes | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 35.2 µs | 34.5 µs | 0.99× [0.97, 1.01] | -1.1% | [-2.9%, +0.6%] | 1.1 pts | 1.0 | yes | no |
| multi | str | 16384 | churn | ordered | baseline | 6 | 123 | 120 | 0.98× [0.96, 0.99] | -2.3% | [-3.7%, -0.9%] | 1.0 pts | 1.0 | yes | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 40.68 ms | 39.97 ms | 0.98× [0.97, 0.99] | -1.8% | [-3.0%, -0.6%] | 0.8 pts | 1.4 | yes | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 6 | 254 | 266 | 1.04× [1.02, 1.05] | +3.5% | [+1.9%, +5.0%] | 1.2 pts | 0.8 | yes | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 6 | 8391 | 8540 | 1.02× [1.01, 1.03] | +2.2% | [+1.5%, +2.9%] | 0.7 pts | 0.9 | yes | yes |
| multi | str | 262144 | prefix | ordered | baseline | 6 | 1.28 ms | 1.31 ms | 1.03× [1.01, 1.05] | +3.1% | [+1.2%, +5.0%] | 1.2 pts | 1.9 | yes | yes |
| multi | str | 262144 | churn | ordered | baseline | 6 | 462 | 452 | 0.96× [0.94, 0.97] | -4.5% | [-5.9%, -3.1%] | 3.8 pts | 0.5 | yes | yes |
| multi | street | 4096 | valuesFor | ordered | baseline | 8 | 54.2 | 54.1 | 1.00× [0.97, 1.03] | +0.0% | [-3.0%, +3.0%] | 2.0 pts | 1.2 | no | no |
| multi | street | 4096 | valuesBetween | ordered | baseline | 8 | 2458 | 2443 | 0.99× [0.96, 1.02] | -1.3% | [-4.4%, +1.8%] | 2.6 pts | 0.9 | no | no |
| multi | street | 4096 | prefix | ordered | baseline | 8 | 247 | 237 | 0.97× [0.95, 0.98] | -3.5% | [-5.0%, -1.9%] | 1.2 pts | 0.5 | yes | yes |
| multi | street | 4096 | churn | ordered | baseline | 8 | 101 | 96.6 | 0.95× [0.95, 0.95] | -5.4% | [-5.6%, -5.2%] | 0.8 pts | 0.4 | yes | yes |
| multi | street | 4096 | build | ordered | baseline | 8 | 3.89 ms | 3.64 ms | 0.94× [0.92, 0.95] | -6.6% | [-8.1%, -5.1%] | 1.2 pts | 1.3 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 8 | 76.6 | 76.5 | 1.00× [0.98, 1.03] | +0.3% | [-2.0%, +2.5%] | 1.4 pts | 0.9 | no | no |
| multi | street | 16384 | valuesBetween | ordered | baseline | 8 | 2927 | 2841 | 0.97× [0.97, 0.98] | -2.6% | [-3.3%, -1.9%] | 0.6 pts | 0.4 | yes | yes |
| multi | street | 16384 | prefix | ordered | baseline | 8 | 903 | 885 | 0.98× [0.97, 1.00] | -1.7% | [-3.2%, -0.2%] | 2.0 pts | 0.8 | yes | no |
| multi | street | 16384 | churn | ordered | baseline | 8 | 143 | 138 | 0.96× [0.95, 0.98] | -3.8% | [-5.8%, -1.9%] | 1.3 pts | 0.8 | yes | yes |
| multi | street | 16384 | build | ordered | baseline | 8 | 20.50 ms | 19.54 ms | 0.96× [0.95, 0.96] | -4.6% | [-5.4%, -3.8%] | 0.6 pts | 1.5 | yes | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 8 | 38.8 | 38.7 | 1.00× [1.00, 1.01] | +0.4% | [-0.5%, +1.2%] | 0.6 pts | 0.3 | yes | no |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2854 | 2752 | 0.96× [0.95, 0.98] | -3.9% | [-5.7%, -2.0%] | 1.9 pts | 0.6 | yes | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 8 | 51.8 | 49.9 | 0.96× [0.95, 0.97] | -4.4% | [-5.3%, -3.4%] | 1.0 pts | 0.8 | yes | yes |
| multi | u64 | 4096 | build | ordered | baseline | 8 | 5.39 ms | 5.14 ms | 0.96× [0.94, 0.98] | -4.1% | [-6.0%, -2.2%] | 1.2 pts | 0.9 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 8 | 47.1 | 46.2 | 0.98× [0.98, 0.99] | -1.6% | [-2.2%, -1.1%] | 0.7 pts | 0.8 | yes | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 3607 | 3480 | 0.97× [0.96, 0.97] | -3.5% | [-4.1%, -3.0%] | 0.8 pts | 1.0 | yes | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 8 | 73.6 | 70.3 | 0.95× [0.93, 0.97] | -5.3% | [-7.4%, -3.2%] | 1.5 pts | 1.3 | no | yes |
| multi | u64 | 16384 | build | ordered | baseline | 8 | 25.75 ms | 24.74 ms | 0.96× [0.95, 0.97] | -4.2% | [-5.0%, -3.4%] | 0.7 pts | 1.1 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 8 | 138 | 139 | 1.00× [0.96, 1.04] | +0.0% | [-3.6%, +3.7%] | 3.1 pts | 3.5 | no | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 6954 | 6806 | 0.99× [0.95, 1.02] | -1.5% | [-4.8%, +1.8%] | 2.3 pts | 3.3 | no | no |
| multi | u64 | 262144 | churn | ordered | baseline | 8 | 312 | 292 | 0.92× [0.88, 0.98] | -8.3% | [-14.1%, -2.6%] | 5.5 pts | 0.8 | no | yes |
| multi | url | 4096 | valuesFor | ordered | baseline | 4 | 91.1 | 89.5 | 0.99× [0.97, 1.01] | -1.4% | [-3.2%, +0.5%] | 1.2 pts | 0.9 | yes | no |
| multi | url | 4096 | valuesBetween | ordered | baseline | 4 | 4061 | 4075 | 1.00× [0.99, 1.02] | +0.5% | [-0.9%, +1.9%] | 0.9 pts | 0.4 | yes | no |
| multi | url | 4096 | prefix | ordered | baseline | 4 | 194 | 188 | 0.98× [0.96, 1.00] | -1.8% | [-3.8%, +0.1%] | 1.2 pts | 0.7 | yes | no |
| multi | url | 4096 | churn | ordered | baseline | 4 | 135 | 131 | 0.96× [0.94, 0.98] | -4.5% | [-6.4%, -2.6%] | 1.2 pts | 0.8 | yes | yes |
| multi | url | 4096 | build | ordered | baseline | 4 | 12.23 ms | 11.59 ms | 0.95× [0.94, 0.96] | -5.3% | [-6.5%, -4.1%] | 0.8 pts | 1.8 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 8 | 114 | 113 | 0.99× [0.97, 1.01] | -1.3% | [-3.5%, +1.0%] | 1.9 pts | 1.1 | no | no |
| multi | url | 16384 | valuesBetween | ordered | baseline | 8 | 4555 | 4508 | 0.98× [0.97, 1.00] | -1.5% | [-2.9%, -0.2%] | 1.5 pts | 1.3 | yes | yes |
| multi | url | 16384 | prefix | ordered | baseline | 8 | 306 | 295 | 0.98× [0.95, 1.00] | -2.6% | [-4.9%, -0.3%] | 2.7 pts | 1.4 | no | yes |
| multi | url | 16384 | churn | ordered | baseline | 8 | 209 | 205 | 0.97× [0.97, 0.98] | -2.6% | [-3.6%, -1.7%] | 1.0 pts | 0.8 | yes | yes |
| multi | url | 16384 | build | ordered | baseline | 8 | 65.76 ms | 63.50 ms | 0.97× [0.96, 0.98] | -3.4% | [-4.5%, -2.3%] | 1.0 pts | 1.1 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 8 | 420 | 427 | 1.02× [0.97, 1.09] | +2.3% | [-3.5%, +8.2%] | 4.8 pts | 3.6 | no | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 8 | 11.4 µs | 11.8 µs | 1.03× [1.02, 1.04] | +3.3% | [+2.4%, +4.1%] | 0.9 pts | 1.1 | yes | yes |
| multi | url | 262144 | prefix | ordered | baseline | 8 | 2517 | 2648 | 1.04× [1.02, 1.05] | +3.6% | [+2.3%, +4.9%] | 2.4 pts | 1.0 | yes | yes |
| multi | url | 262144 | churn | ordered | baseline | 8 | 722 | 723 | 1.00× [0.92, 1.09] | +0.1% | [-8.2%, +8.5%] | 5.2 pts | 0.8 | no | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 8 | 56.9 | 54.4 | 0.96× [0.94, 0.97] | -4.6% | [-6.5%, -2.8%] | 1.3 pts | 0.8 | yes | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 8 | 3388 | 3351 | 0.99× [0.96, 1.01] | -1.4% | [-4.1%, +1.3%] | 2.3 pts | 0.7 | no | no |
| multi | uuid | 4096 | prefix | ordered | baseline | 8 | 93.2 | 91.0 | 0.97× [0.96, 0.98] | -2.7% | [-3.7%, -1.7%] | 0.9 pts | 0.8 | yes | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 8 | 82.3 | 79.1 | 0.96× [0.94, 0.97] | -4.7% | [-6.3%, -3.1%] | 1.0 pts | 0.8 | yes | yes |
| multi | uuid | 4096 | build | ordered | baseline | 8 | 7.94 ms | 7.57 ms | 0.95× [0.95, 0.96] | -4.8% | [-5.5%, -4.1%] | 1.3 pts | 1.2 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 65.0 | 63.2 | 0.96× [0.95, 0.98] | -3.6% | [-5.4%, -1.8%] | 1.2 pts | 1.1 | yes | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3669 | 3606 | 0.98× [0.98, 0.99] | -1.6% | [-2.0%, -1.2%] | 0.3 pts | 0.3 | yes | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 111 | 109 | 0.98× [0.96, 1.00] | -2.0% | [-4.0%, -0.1%] | 1.4 pts | 0.9 | yes | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 118 | 113 | 0.95× [0.94, 0.97] | -4.8% | [-6.7%, -2.8%] | 1.3 pts | 1.2 | yes | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 37.97 ms | 36.28 ms | 0.96× [0.95, 0.97] | -4.2% | [-5.2%, -3.2%] | 0.6 pts | 0.9 | yes | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 8 | 242 | 248 | 1.02× [1.00, 1.05] | +2.4% | [-0.0%, +4.9%] | 1.9 pts | 2.5 | no | no |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 8810 | 8865 | 1.01× [0.99, 1.03] | +0.6% | [-1.3%, +2.5%] | 1.4 pts | 2.0 | yes | no |
| multi | uuid | 262144 | prefix | ordered | baseline | 8 | 575 | 579 | 1.00× [0.98, 1.02] | +0.1% | [-1.6%, +1.8%] | 1.1 pts | 1.4 | yes | no |
| multi | uuid | 262144 | churn | ordered | baseline | 8 | 521 | 491 | 0.91× [0.88, 0.94] | -9.9% | [-13.7%, -6.2%] | 5.6 pts | 0.7 | no | yes |
| unique | email | 4096 | valuesFor | ordered | baseline | 8 | 26.6 | 26.4 | 0.99× [0.95, 1.04] | -0.8% | [-5.3%, +3.8%] | 3.3 pts | 1.9 | no | no |
| unique | email | 4096 | valuesBetween | ordered | baseline | 8 | 1320 | 1292 | 0.99× [0.97, 1.00] | -1.3% | [-2.7%, +0.1%] | 2.8 pts | 2.0 | yes | no |
| unique | email | 4096 | prefix | ordered | baseline | 8 | 59.3 | 56.7 | 0.95× [0.94, 0.96] | -4.9% | [-6.0%, -3.7%] | 1.3 pts | 1.6 | yes | yes |
| unique | email | 4096 | churn | ordered | baseline | 8 | 74.3 | 71.8 | 0.97× [0.95, 0.98] | -3.4% | [-5.1%, -1.6%] | 1.1 pts | 0.9 | yes | yes |
| unique | email | 4096 | build | ordered | baseline | 8 | 1.07 ms | 1.00 ms | 0.94× [0.92, 0.96] | -6.5% | [-8.7%, -4.2%] | 1.6 pts | 0.9 | no | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 8 | 38.8 | 37.9 | 0.97× [0.96, 0.99] | -2.6% | [-4.0%, -1.2%] | 1.0 pts | 0.7 | yes | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 8 | 1654 | 1607 | 0.98× [0.97, 0.99] | -2.2% | [-3.6%, -0.9%] | 1.7 pts | 1.4 | yes | yes |
| unique | email | 16384 | prefix | ordered | baseline | 8 | 72.8 | 69.9 | 0.96× [0.95, 0.97] | -4.4% | [-5.6%, -3.2%] | 1.4 pts | 1.1 | yes | yes |
| unique | email | 16384 | churn | ordered | baseline | 8 | 94.0 | 89.3 | 0.95× [0.93, 0.97] | -5.4% | [-7.6%, -3.2%] | 1.6 pts | 0.9 | no | yes |
| unique | email | 16384 | build | ordered | baseline | 8 | 5.27 ms | 4.94 ms | 0.93× [0.93, 0.94] | -7.1% | [-7.9%, -6.3%] | 0.6 pts | 0.7 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 8 | 163 | 168 | 1.04× [1.02, 1.07] | +4.1% | [+1.8%, +6.4%] | 2.0 pts | 2.2 | no | yes |
| unique | email | 262144 | valuesBetween | ordered | baseline | 8 | 3740 | 3844 | 1.01× [1.00, 1.01] | +0.7% | [+0.4%, +1.0%] | 2.1 pts | 2.0 | yes | no |
| unique | email | 262144 | prefix | ordered | baseline | 8 | 208 | 210 | 1.01× [0.99, 1.04] | +1.2% | [-1.3%, +3.6%] | 2.3 pts | 1.9 | no | no |
| unique | email | 262144 | churn | ordered | baseline | 8 | 269 | 266 | 0.98× [0.97, 0.99] | -2.1% | [-2.8%, -1.4%] | 0.8 pts | 0.5 | yes | yes |
| unique | path | 4096 | valuesFor | ordered | baseline | 8 | 86.1 | 82.1 | 0.96× [0.94, 0.98] | -4.2% | [-6.9%, -1.6%] | 2.0 pts | 0.9 | no | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 8 | 2216 | 2181 | 0.99× [0.98, 1.00] | -1.0% | [-2.4%, +0.3%] | 2.2 pts | 0.9 | yes | no |
| unique | path | 4096 | prefix | ordered | baseline | 8 | 256 | 246 | 0.96× [0.94, 0.99] | -3.7% | [-5.9%, -1.5%] | 1.5 pts | 0.4 | no | yes |
| unique | path | 4096 | churn | ordered | baseline | 8 | 217 | 193 | 0.89× [0.87, 0.91] | -12.7% | [-15.2%, -10.2%] | 1.7 pts | 1.0 | no | yes |
| unique | path | 4096 | build | ordered | baseline | 8 | 2.87 ms | 2.52 ms | 0.87× [0.86, 0.88] | -14.8% | [-16.1%, -13.5%] | 0.9 pts | 1.0 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 8 | 118 | 114 | 0.96× [0.94, 0.98] | -4.1% | [-6.0%, -2.1%] | 1.5 pts | 0.8 | yes | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 8 | 2666 | 2653 | 0.99× [0.98, 1.00] | -0.7% | [-1.8%, +0.5%] | 0.9 pts | 0.6 | yes | no |
| unique | path | 16384 | prefix | ordered | baseline | 8 | 491 | 494 | 0.99× [0.98, 1.01] | -0.5% | [-2.3%, +1.2%] | 2.7 pts | 1.1 | yes | no |
| unique | path | 16384 | churn | ordered | baseline | 8 | 291 | 273 | 0.94× [0.92, 0.96] | -6.3% | [-8.3%, -4.4%] | 1.9 pts | 1.1 | yes | yes |
| unique | path | 16384 | build | ordered | baseline | 8 | 13.96 ms | 12.65 ms | 0.91× [0.90, 0.92] | -10.4% | [-11.5%, -9.3%] | 1.0 pts | 2.1 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 8 | 371 | 381 | 1.03× [1.00, 1.06] | +2.8% | [-0.2%, +5.7%] | 2.1 pts | 1.2 | no | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 8 | 6926 | 7865 | 1.14× [1.12, 1.16] | +12.2% | [+10.7%, +13.7%] | 2.1 pts | 2.2 | yes | yes |
| unique | path | 262144 | prefix | ordered | baseline | 8 | 5070 | 5468 | 1.08× [1.05, 1.11] | +7.3% | [+4.8%, +9.8%] | 2.7 pts | 2.4 | no | yes |
| unique | path | 262144 | churn | ordered | baseline | 8 | 751 | 742 | 1.01× [0.98, 1.04] | +1.1% | [-2.1%, +4.2%] | 2.7 pts | 0.9 | no | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 8 | 39.1 | 38.4 | 0.98× [0.95, 1.01] | -2.2% | [-5.0%, +0.7%] | 1.7 pts | 1.2 | no | no |
| unique | str | 4096 | valuesBetween | ordered | baseline | 8 | 1710 | 1692 | 0.98× [0.97, 1.00] | -1.9% | [-3.2%, -0.5%] | 1.1 pts | 0.5 | yes | yes |
| unique | str | 4096 | prefix | ordered | baseline | 8 | 3318 | 3291 | 1.01× [0.99, 1.03] | +0.7% | [-1.0%, +2.5%] | 1.5 pts | 0.7 | yes | no |
| unique | str | 4096 | churn | ordered | baseline | 8 | 103 | 94.8 | 0.93× [0.92, 0.95] | -7.0% | [-9.0%, -4.9%] | 1.5 pts | 1.1 | no | yes |
| unique | str | 4096 | build | ordered | baseline | 8 | 1.46 ms | 1.31 ms | 0.90× [0.89, 0.91] | -11.7% | [-13.0%, -10.3%] | 2.3 pts | 0.9 | yes | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 8 | 50.6 | 50.1 | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.4%] | 1.1 pts | 0.7 | yes | no |
| unique | str | 16384 | valuesBetween | ordered | baseline | 8 | 1713 | 1679 | 0.98× [0.98, 0.99] | -1.7% | [-2.4%, -1.0%] | 0.6 pts | 0.5 | yes | yes |
| unique | str | 16384 | prefix | ordered | baseline | 8 | 14.3 µs | 14.1 µs | 0.99× [0.99, 1.00] | -0.7% | [-1.4%, -0.0%] | 1.2 pts | 0.7 | yes | no |
| unique | str | 16384 | churn | ordered | baseline | 8 | 124 | 121 | 0.98× [0.96, 1.00] | -2.3% | [-4.5%, -0.0%] | 2.0 pts | 1.0 | no | yes |
| unique | str | 16384 | build | ordered | baseline | 8 | 6.69 ms | 6.21 ms | 0.93× [0.93, 0.94] | -7.4% | [-7.9%, -6.8%] | 0.5 pts | 0.5 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 8 | 164 | 191 | 1.12× [1.11, 1.14] | +11.1% | [+9.6%, +12.5%] | 3.4 pts | 2.6 | yes | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 8 | 3381 | 3777 | 1.12× [1.08, 1.16] | +10.5% | [+7.5%, +13.4%] | 2.3 pts | 1.8 | no | yes |
| unique | str | 262144 | prefix | ordered | baseline | 8 | 501.6 µs | 585.2 µs | 1.14× [1.11, 1.17] | +12.2% | [+10.0%, +14.4%] | 1.9 pts | 2.5 | no | yes |
| unique | str | 262144 | churn | ordered | baseline | 8 | 384 | 381 | 1.00× [0.95, 1.04] | -0.4% | [-4.9%, +4.0%] | 3.8 pts | 0.8 | no | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 8 | 47.0 | 47.6 | 1.01× [0.99, 1.02] | +0.6% | [-1.0%, +2.2%] | 1.4 pts | 1.0 | yes | no |
| unique | street | 4096 | valuesBetween | ordered | baseline | 8 | 1929 | 1898 | 0.99× [0.97, 1.00] | -1.5% | [-2.7%, -0.3%] | 0.8 pts | 0.5 | yes | yes |
| unique | street | 4096 | prefix | ordered | baseline | 8 | 216 | 211 | 0.98× [0.97, 1.00] | -1.7% | [-3.4%, +0.1%] | 1.4 pts | 0.7 | yes | no |
| unique | street | 4096 | churn | ordered | baseline | 8 | 124 | 114 | 0.93× [0.92, 0.94] | -7.3% | [-8.6%, -6.0%] | 1.0 pts | 0.7 | yes | yes |
| unique | street | 4096 | build | ordered | baseline | 8 | 1.73 ms | 1.56 ms | 0.90× [0.89, 0.91] | -10.9% | [-11.8%, -10.1%] | 0.7 pts | 1.0 | yes | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 4 | 69.9 | 69.1 | 0.99× [0.97, 1.01] | -1.4% | [-3.3%, +0.6%] | 1.2 pts | 0.7 | yes | no |
| unique | street | 16384 | valuesBetween | ordered | baseline | 4 | 2282 | 2208 | 0.97× [0.95, 0.98] | -3.4% | [-5.2%, -1.7%] | 1.1 pts | 0.6 | yes | yes |
| unique | street | 16384 | prefix | ordered | baseline | 4 | 724 | 703 | 0.97× [0.96, 0.98] | -3.0% | [-4.0%, -1.9%] | 0.6 pts | 0.3 | yes | yes |
| unique | street | 16384 | churn | ordered | baseline | 4 | 164 | 153 | 0.95× [0.93, 0.96] | -5.6% | [-7.2%, -3.9%] | 1.0 pts | 0.6 | yes | yes |
| unique | street | 16384 | build | ordered | baseline | 4 | 8.58 ms | 7.85 ms | 0.92× [0.91, 0.93] | -8.9% | [-9.7%, -8.0%] | 0.5 pts | 0.7 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 8 | 15.1 | 15.1 | 1.01× [1.00, 1.01] | +0.8% | [+0.5%, +1.2%] | 0.7 pts | 0.4 | yes | no |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 1076 | 1006 | 0.93× [0.93, 0.94] | -7.1% | [-7.7%, -6.5%] | 0.6 pts | 0.5 | yes | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 8 | 47.8 | 45.5 | 0.95× [0.93, 0.96] | -5.6% | [-7.0%, -4.3%] | 0.9 pts | 1.0 | yes | yes |
| unique | u64 | 4096 | build | ordered | baseline | 8 | 767.0 µs | 721.1 µs | 0.94× [0.93, 0.96] | -5.9% | [-8.0%, -3.7%] | 1.6 pts | 1.1 | no | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 8 | 24.2 | 23.4 | 0.97× [0.96, 0.97] | -3.4% | [-4.2%, -2.6%] | 0.9 pts | 1.0 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1634 | 1558 | 0.95× [0.93, 0.96] | -5.6% | [-7.0%, -4.2%] | 1.7 pts | 0.9 | yes | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 8 | 57.4 | 52.8 | 0.93× [0.91, 0.95] | -7.3% | [-9.8%, -4.8%] | 1.6 pts | 0.9 | no | yes |
| unique | u64 | 16384 | build | ordered | baseline | 8 | 3.70 ms | 3.44 ms | 0.93× [0.91, 0.94] | -8.1% | [-9.3%, -6.8%] | 1.6 pts | 1.2 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 8 | 77.5 | 75.4 | 0.96× [0.93, 1.00] | -4.0% | [-7.6%, -0.4%] | 3.2 pts | 2.4 | no | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 2046 | 1965 | 0.96× [0.93, 0.99] | -4.3% | [-7.9%, -0.7%] | 2.5 pts | 1.8 | no | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 8 | 226 | 221 | 0.98× [0.96, 1.00] | -2.1% | [-4.7%, +0.4%] | 1.8 pts | 0.6 | no | no |
| unique | url | 4096 | valuesFor | ordered | baseline | 8 | 66.6 | 65.3 | 0.97× [0.97, 0.98] | -2.7% | [-3.3%, -2.2%] | 1.2 pts | 0.7 | yes | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 8 | 2059 | 2016 | 0.98× [0.97, 0.99] | -1.6% | [-2.6%, -0.6%] | 1.0 pts | 0.4 | yes | yes |
| unique | url | 4096 | prefix | ordered | baseline | 8 | 156 | 152 | 0.97× [0.95, 0.98] | -3.5% | [-4.7%, -2.3%] | 1.2 pts | 0.7 | yes | yes |
| unique | url | 4096 | churn | ordered | baseline | 8 | 170 | 154 | 0.91× [0.90, 0.92] | -10.5% | [-11.7%, -9.2%] | 1.4 pts | 0.7 | yes | yes |
| unique | url | 4096 | build | ordered | baseline | 8 | 2.30 ms | 2.04 ms | 0.88× [0.87, 0.90] | -13.3% | [-15.5%, -11.2%] | 1.6 pts | 1.1 | no | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 8 | 91.2 | 90.1 | 0.99× [0.96, 1.01] | -1.4% | [-3.6%, +0.9%] | 1.4 pts | 0.9 | no | no |
| unique | url | 16384 | valuesBetween | ordered | baseline | 8 | 2502 | 2440 | 0.98× [0.96, 1.01] | -1.8% | [-4.7%, +1.1%] | 1.8 pts | 1.8 | no | no |
| unique | url | 16384 | prefix | ordered | baseline | 8 | 233 | 229 | 0.98× [0.97, 0.98] | -2.6% | [-3.2%, -1.9%] | 0.7 pts | 0.4 | yes | yes |
| unique | url | 16384 | churn | ordered | baseline | 8 | 227 | 211 | 0.94× [0.92, 0.96] | -6.4% | [-8.4%, -4.4%] | 1.5 pts | 0.7 | yes | yes |
| unique | url | 16384 | build | ordered | baseline | 8 | 11.18 ms | 10.12 ms | 0.90× [0.89, 0.91] | -11.4% | [-12.5%, -10.2%] | 0.9 pts | 1.3 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 8 | 366 | 382 | 1.03× [0.96, 1.11] | +3.0% | [-4.2%, +10.2%] | 6.0 pts | 4.1 | no | no |
| unique | url | 262144 | valuesBetween | ordered | baseline | 8 | 7252 | 7939 | 1.07× [1.05, 1.10] | +6.9% | [+4.9%, +8.9%] | 2.0 pts | 1.9 | no | yes |
| unique | url | 262144 | prefix | ordered | baseline | 8 | 1564 | 1638 | 1.05× [0.98, 1.12] | +4.4% | [-2.0%, +10.9%] | 4.5 pts | 2.2 | no | no |
| unique | url | 262144 | churn | ordered | baseline | 8 | 698 | 687 | 0.98× [0.95, 1.02] | -1.7% | [-5.2%, +1.7%] | 2.6 pts | 0.8 | no | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 8 | 32.3 | 29.8 | 0.92× [0.91, 0.93] | -8.6% | [-9.4%, -7.7%] | 0.9 pts | 0.6 | yes | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 8 | 1525 | 1504 | 0.99× [0.97, 1.00] | -1.2% | [-2.7%, +0.3%] | 1.0 pts | 0.6 | yes | no |
| unique | uuid | 4096 | prefix | ordered | baseline | 8 | 70.8 | 68.1 | 0.96× [0.96, 0.97] | -3.9% | [-4.4%, -3.3%] | 0.9 pts | 0.9 | yes | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 8 | 86.4 | 81.4 | 0.94× [0.92, 0.96] | -6.4% | [-8.4%, -4.3%] | 1.3 pts | 0.8 | no | yes |
| unique | uuid | 4096 | build | ordered | baseline | 8 | 1.19 ms | 1.10 ms | 0.92× [0.91, 0.94] | -8.4% | [-10.1%, -6.7%] | 1.5 pts | 0.8 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 6 | 42.6 | 39.4 | 0.93× [0.92, 0.93] | -7.9% | [-8.7%, -7.2%] | 0.5 pts | 0.4 | yes | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 1669 | 1613 | 0.97× [0.95, 0.98] | -3.5% | [-5.4%, -1.6%] | 1.2 pts | 1.1 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 85.4 | 82.3 | 0.96× [0.95, 0.97] | -4.2% | [-5.0%, -3.5%] | 0.7 pts | 0.6 | yes | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 6 | 106 | 98.6 | 0.95× [0.94, 0.96] | -5.0% | [-6.3%, -3.7%] | 1.3 pts | 0.8 | yes | yes |
| unique | uuid | 16384 | build | ordered | baseline | 6 | 5.73 ms | 5.27 ms | 0.92× [0.91, 0.93] | -8.6% | [-10.3%, -7.0%] | 1.1 pts | 1.1 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 8 | 182 | 184 | 1.04× [1.01, 1.08] | +4.3% | [+1.2%, +7.3%] | 2.4 pts | 3.0 | no | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 4469 | 4313 | 0.97× [0.93, 1.02] | -3.0% | [-7.6%, +1.6%] | 3.0 pts | 2.2 | no | no |
| unique | uuid | 262144 | prefix | ordered | baseline | 8 | 347 | 342 | 0.98× [0.95, 1.02] | -1.6% | [-5.2%, +1.9%] | 2.7 pts | 4.2 | no | no |
| unique | uuid | 262144 | churn | ordered | baseline | 8 | 404 | 397 | 0.95× [0.92, 0.99] | -5.0% | [-8.9%, -1.2%] | 4.7 pts | 0.9 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.20% does not clear the 0.77% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=4096 valuesFor: ordered vs baseline: the pooled interval [-4.14%, 3.73%] includes zero
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-3.63%, 1.44%] includes zero
- multi email n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of -0.37% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi email n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.91%, 0.69%] includes zero
- multi email n=262144 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi email n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.41%, 4.20%] includes zero
- multi email n=262144 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.12%, 0.45%] includes zero
- multi email n=262144 prefix: ordered vs baseline: the pooled difference of -0.35% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=262144 prefix: ordered vs baseline: the pooled interval [-1.39%, 0.69%] includes zero
- multi email n=262144 churn: ordered vs baseline: the pooled interval [-14.42%, 1.99%] includes zero
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.65% does not clear the 1.00% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.90%, 2.19%] includes zero
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of 1.48% does not clear the 2.75% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 prefix: ordered vs baseline: the pooled interval [-3.63%, 6.60%] includes zero
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.39% does not clear the 0.47% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.22%, 0.44%] includes zero
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of 1.67% does not clear the 5.19% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-2.81%, 6.14%] includes zero
- multi path n=16384 churn: ordered vs baseline: the pooled difference of 0.32% does not clear the 0.51% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 churn: ordered vs baseline: the pooled interval [-1.40%, 2.03%] includes zero
- multi path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=262144 prefix: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=262144 churn: ordered vs baseline: the pooled difference of -0.59% does not clear the 1.91% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 churn: ordered vs baseline: the pooled interval [-4.17%, 2.99%] includes zero
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.62%, 1.06%] includes zero
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.22% does not clear the 1.74% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 prefix: ordered vs baseline: the pooled interval [-2.87%, 0.57%] includes zero
- multi street n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.00% does not clear the 0.63% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=4096 valuesFor: ordered vs baseline: the pooled interval [-3.00%, 3.00%] includes zero
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-4.43%, 1.82%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.26% does not clear the 0.67% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.97%, 2.50%] includes zero
- multi street n=16384 prefix: ordered vs baseline: the pooled difference of -1.69% does not clear the 1.83% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.38% does not clear the 0.77% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.47%, 1.22%] includes zero
- multi u64 n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi u64 n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.05% does not clear the 0.29% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.63%, 3.72%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-4.81%, 1.77%] includes zero
- multi u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of +1.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi url n=4096 valuesFor: ordered vs baseline: the pooled interval [-3.22%, 0.52%] includes zero
- multi url n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.49% does not clear the 1.17% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.92%, 1.89%] includes zero
- multi url n=4096 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi url n=4096 prefix: ordered vs baseline: the pooled interval [-3.77%, 0.15%] includes zero
- multi url n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi url n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi url n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.53%, 1.00%] includes zero
- multi url n=16384 prefix: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.50%, 8.19%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=262144 churn: ordered vs baseline: the pooled difference of 0.14% does not clear the 0.85% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=262144 churn: ordered vs baseline: the pooled interval [-8.22%, 8.50%] includes zero
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-4.08%, 1.31%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.04%, 4.89%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.33%, 2.50%] includes zero
- multi uuid n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 prefix: ordered vs baseline: the pooled difference of 0.14% does not clear the 0.71% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=262144 prefix: ordered vs baseline: the pooled interval [-1.56%, 1.84%] includes zero
- unique email n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.77% does not clear the 1.07% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=4096 valuesFor: ordered vs baseline: the pooled interval [-5.32%, 3.78%] includes zero
- unique email n=4096 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.69%, 0.11%] includes zero
- unique email n=4096 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=4096 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=262144 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=262144 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.70% does not clear the 0.76% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=262144 prefix: ordered vs baseline: the pooled interval [-1.31%, 3.61%] includes zero
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.04% does not clear the 1.27% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.37%, 0.29%] includes zero
- unique path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.82%, 0.50%] includes zero
- unique path n=16384 prefix: ordered vs baseline: the pooled difference of -0.53% does not clear the 2.47% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=16384 prefix: ordered vs baseline: the pooled interval [-2.26%, 1.21%] includes zero
- unique path n=16384 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.16%, 5.73%] includes zero
- unique path n=262144 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 prefix: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 churn: ordered vs baseline: the pooled interval [-2.09%, 4.20%] includes zero
- unique str n=4096 valuesFor: ordered vs baseline: the pooled interval [-5.00%, 0.65%] includes zero
- unique str n=4096 prefix: ordered vs baseline: the pooled difference of 0.75% does not clear the 1.16% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 prefix: ordered vs baseline: the pooled interval [-1.05%, 2.54%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.58% does not clear the 0.94% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.60%, 0.45%] includes zero
- unique str n=16384 prefix: ordered vs baseline: the pooled difference of -0.70% does not clear the 0.94% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 prefix: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 churn: ordered vs baseline: the pooled difference of -0.45% does not clear the 1.66% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=262144 churn: ordered vs baseline: the pooled interval [-4.88%, 3.98%] includes zero
- unique street n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.62% does not clear the 0.71% noise floor, the bound on what the harness reports between identical code in every process
- unique street n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.97%, 2.21%] includes zero
- unique street n=4096 prefix: ordered vs baseline: the pooled interval [-3.44%, 0.09%] includes zero
- unique street n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique street n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.33%, 0.59%] includes zero
- unique street n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique street n=16384 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique street n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique street n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.84% does not clear the 1.23% noise floor, the bound on what the harness reports between identical code in every process
- unique u64 n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 churn: ordered vs baseline: the pooled interval [-4.67%, 0.41%] includes zero
- unique url n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.64%, 0.92%] includes zero
- unique url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-4.69%, 1.15%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the pooled interval [-4.23%, 10.16%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 prefix: ordered vs baseline: the pooled interval [-2.05%, 10.94%] includes zero
- unique url n=262144 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 churn: ordered vs baseline: the pooled interval [-5.16%, 1.74%] includes zero
- unique uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.68%, 0.29%] includes zero
- unique uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-7.59%, 1.61%] includes zero
- unique uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 prefix: ordered vs baseline: the pooled interval [-5.18%, 1.93%] includes zero
- unique uuid n=262144 prefix: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
