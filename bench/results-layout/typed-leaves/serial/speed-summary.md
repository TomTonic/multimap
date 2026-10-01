| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | email | 4096 | valuesFor | ordered | baseline | 12 | 49.6 | 49.6 | 0.99× [0.97, 1.01] | -1.0% | [-3.0%, +1.0%] | 5.3 pts | 2.0 | no | no |
| multi-str | email | 4096 | valuesBetween | ordered | baseline | 12 | 2978 | 2903 | 0.98× [0.96, 0.99] | -2.3% | [-4.0%, -0.7%] | 2.3 pts | 0.8 | yes | yes |
| multi-str | email | 4096 | prefix | ordered | baseline | 12 | 79.0 | 75.8 | 0.96× [0.96, 0.97] | -3.9% | [-4.5%, -3.2%] | 0.8 pts | 0.7 | yes | yes |
| multi-str | email | 4096 | churn | ordered | baseline | 12 | 89.3 | 82.8 | 0.93× [0.92, 0.94] | -7.5% | [-8.5%, -6.5%] | 1.4 pts | 1.8 | yes | yes |
| multi-str | email | 4096 | build | ordered | baseline | 12 | 9.26 ms | 8.43 ms | 0.91× [0.90, 0.91] | -10.2% | [-10.6%, -9.7%] | 0.7 pts | 1.0 | yes | yes |
| multi-str | email | 16384 | valuesFor | ordered | baseline | 6 | 61.9 | 62.6 | 1.01× [1.01, 1.02] | +1.4% | [+0.8%, +2.0%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | email | 16384 | valuesBetween | ordered | baseline | 6 | 3595 | 3564 | 0.99× [0.98, 1.00] | -0.7% | [-1.7%, +0.3%] | 0.9 pts | 1.3 | yes | no |
| multi-str | email | 16384 | prefix | ordered | baseline | 6 | 93.1 | 92.4 | 0.99× [0.98, 1.00] | -1.0% | [-1.6%, -0.5%] | 0.5 pts | 0.9 | yes | yes |
| multi-str | email | 16384 | churn | ordered | baseline | 6 | 132 | 123 | 0.94× [0.93, 0.95] | -6.4% | [-7.3%, -5.6%] | 0.8 pts | 0.9 | yes | yes |
| multi-str | email | 16384 | build | ordered | baseline | 6 | 47.27 ms | 43.12 ms | 0.91× [0.90, 0.92] | -9.9% | [-10.9%, -8.9%] | 0.9 pts | 1.4 | yes | yes |
| multi-str | email | 262144 | valuesFor | ordered | baseline | 12 | 230 | 250 | 1.10× [1.08, 1.11] | +8.8% | [+7.6%, +10.1%] | 2.6 pts | 3.8 | yes | yes |
| multi-str | email | 262144 | valuesBetween | ordered | baseline | 12 | 8821 | 9916 | 1.13× [1.10, 1.16] | +11.5% | [+9.5%, +13.6%] | 3.0 pts | 4.2 | no | yes |
| multi-str | email | 262144 | prefix | ordered | baseline | 12 | 317 | 349 | 1.11× [1.10, 1.13] | +10.1% | [+9.0%, +11.3%] | 1.7 pts | 3.8 | yes | yes |
| multi-str | email | 262144 | churn | ordered | baseline | 12 | 507 | 474 | 0.91× [0.88, 0.95] | -9.4% | [-13.1%, -5.8%] | 3.8 pts | 0.5 | no | yes |
| multi-str | path | 4096 | valuesFor | ordered | baseline | 6 | 108 | 105 | 0.97× [0.96, 0.99] | -2.6% | [-4.0%, -1.3%] | 1.3 pts | 1.1 | yes | yes |
| multi-str | path | 4096 | valuesBetween | ordered | baseline | 6 | 4106 | 3972 | 0.97× [0.95, 0.99] | -3.3% | [-5.2%, -1.4%] | 1.8 pts | 0.7 | yes | yes |
| multi-str | path | 4096 | prefix | ordered | baseline | 6 | 317 | 304 | 0.96× [0.94, 0.97] | -4.6% | [-6.2%, -3.0%] | 1.5 pts | 1.2 | yes | yes |
| multi-str | path | 4096 | churn | ordered | baseline | 6 | 180 | 165 | 0.93× [0.91, 0.94] | -8.1% | [-9.7%, -6.5%] | 1.5 pts | 1.9 | yes | yes |
| multi-str | path | 4096 | build | ordered | baseline | 6 | 17.23 ms | 15.86 ms | 0.92× [0.92, 0.92] | -8.5% | [-8.9%, -8.2%] | 0.3 pts | 0.5 | yes | yes |
| multi-str | path | 16384 | valuesFor | ordered | baseline | 12 | 138 | 137 | 0.99× [0.99, 0.99] | -1.2% | [-1.5%, -0.8%] | 1.6 pts | 2.0 | yes | yes |
| multi-str | path | 16384 | valuesBetween | ordered | baseline | 12 | 4764 | 4709 | 0.99× [0.98, 0.99] | -1.2% | [-1.7%, -0.8%] | 0.8 pts | 1.2 | yes | yes |
| multi-str | path | 16384 | prefix | ordered | baseline | 12 | 739 | 723 | 0.98× [0.95, 1.02] | -1.6% | [-5.2%, +2.0%] | 3.5 pts | 1.1 | no | no |
| multi-str | path | 16384 | churn | ordered | baseline | 12 | 262 | 260 | 0.99× [0.99, 1.00] | -0.8% | [-1.4%, -0.2%] | 0.8 pts | 0.9 | yes | yes |
| multi-str | path | 16384 | build | ordered | baseline | 12 | 87.91 ms | 83.73 ms | 0.95× [0.94, 0.96] | -4.9% | [-6.0%, -3.8%] | 1.2 pts | 1.8 | yes | yes |
| multi-str | path | 262144 | valuesFor | ordered | baseline | 12 | 427 | 480 | 1.13× [1.11, 1.14] | +11.4% | [+10.2%, +12.5%] | 2.4 pts | 3.7 | yes | yes |
| multi-str | path | 262144 | valuesBetween | ordered | baseline | 12 | 12.0 µs | 13.2 µs | 1.11× [1.09, 1.12] | +9.5% | [+8.1%, +10.9%] | 1.5 pts | 2.7 | yes | yes |
| multi-str | path | 262144 | prefix | ordered | baseline | 12 | 10.6 µs | 11.3 µs | 1.06× [1.00, 1.13] | +6.0% | [+0.3%, +11.6%] | 6.6 pts | 1.0 | no | yes |
| multi-str | path | 262144 | churn | ordered | baseline | 12 | 815 | 824 | 1.01× [0.98, 1.04] | +0.6% | [-2.5%, +3.7%] | 3.4 pts | 0.8 | no | no |
| multi-str | str | 4096 | valuesFor | ordered | baseline | 12 | 63.2 | 62.6 | 1.00× [0.98, 1.02] | +0.0% | [-1.7%, +1.7%] | 1.9 pts | 1.0 | yes | no |
| multi-str | str | 4096 | valuesBetween | ordered | baseline | 12 | 3434 | 3362 | 0.98× [0.96, 1.00] | -2.0% | [-4.3%, +0.2%] | 2.6 pts | 1.1 | no | no |
| multi-str | str | 4096 | prefix | ordered | baseline | 12 | 7225 | 7168 | 0.99× [0.95, 1.04] | -0.6% | [-4.7%, +3.5%] | 4.4 pts | 1.1 | no | no |
| multi-str | str | 4096 | churn | ordered | baseline | 12 | 107 | 98.8 | 0.92× [0.92, 0.92] | -8.7% | [-9.1%, -8.4%] | 1.0 pts | 1.4 | yes | yes |
| multi-str | str | 4096 | build | ordered | baseline | 12 | 11.09 ms | 9.93 ms | 0.90× [0.89, 0.90] | -11.3% | [-11.8%, -10.9%] | 0.7 pts | 1.1 | yes | yes |
| multi-str | str | 16384 | valuesFor | ordered | baseline | 6 | 74.7 | 76.4 | 1.03× [1.02, 1.03] | +2.6% | [+1.9%, +3.2%] | 0.6 pts | 1.3 | yes | yes |
| multi-str | str | 16384 | valuesBetween | ordered | baseline | 6 | 3744 | 3727 | 1.00× [0.99, 1.00] | -0.2% | [-0.9%, +0.5%] | 0.7 pts | 1.0 | yes | no |
| multi-str | str | 16384 | prefix | ordered | baseline | 6 | 35.0 µs | 34.7 µs | 1.00× [0.99, 1.01] | -0.5% | [-1.5%, +0.5%] | 0.9 pts | 1.0 | yes | no |
| multi-str | str | 16384 | churn | ordered | baseline | 6 | 148 | 140 | 0.95× [0.94, 0.96] | -5.5% | [-6.3%, -4.7%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | str | 16384 | build | ordered | baseline | 6 | 52.98 ms | 49.25 ms | 0.93× [0.92, 0.93] | -7.8% | [-8.6%, -7.0%] | 0.8 pts | 1.3 | yes | yes |
| multi-str | str | 262144 | valuesFor | ordered | baseline | 12 | 255 | 289 | 1.13× [1.11, 1.16] | +11.7% | [+9.5%, +13.9%] | 2.6 pts | 4.7 | no | yes |
| multi-str | str | 262144 | valuesBetween | ordered | baseline | 12 | 9402 | 10.7 µs | 1.14× [1.13, 1.15] | +12.4% | [+11.5%, +13.2%] | 1.5 pts | 2.3 | yes | yes |
| multi-str | str | 262144 | prefix | ordered | baseline | 12 | 1.50 ms | 1.72 ms | 1.14× [1.13, 1.16] | +12.5% | [+11.3%, +13.7%] | 1.6 pts | 2.7 | yes | yes |
| multi-str | str | 262144 | churn | ordered | baseline | 12 | 537 | 513 | 0.95× [0.91, 0.99] | -5.5% | [-9.8%, -1.2%] | 5.0 pts | 0.9 | no | yes |
| multi-str | street | 4096 | valuesFor | ordered | baseline | 10 | 49.3 | 51.0 | 1.04× [1.03, 1.06] | +4.2% | [+2.6%, +5.8%] | 2.7 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | baseline | 10 | 2303 | 2183 | 0.95× [0.93, 0.96] | -5.6% | [-7.0%, -4.3%] | 1.7 pts | 0.6 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | baseline | 10 | 224 | 210 | 0.93× [0.91, 0.95] | -7.5% | [-9.4%, -5.6%] | 2.0 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | churn | ordered | baseline | 10 | 114 | 106 | 0.93× [0.92, 0.93] | -7.8% | [-8.5%, -7.2%] | 1.0 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | build | ordered | baseline | 10 | 4.50 ms | 3.99 ms | 0.88× [0.88, 0.89] | -13.0% | [-13.8%, -12.3%] | 0.9 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | baseline | 6 | 75.1 | 75.8 | 1.01× [0.99, 1.02] | +0.6% | [-0.8%, +2.0%] | 1.4 pts | 1.6 | yes | no |
| multi-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 2892 | 2791 | 0.97× [0.96, 0.98] | -3.2% | [-4.2%, -2.3%] | 0.9 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | baseline | 6 | 827 | 797 | 0.95× [0.94, 0.97] | -4.8% | [-6.2%, -3.3%] | 1.4 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | churn | ordered | baseline | 6 | 155 | 147 | 0.95× [0.94, 0.96] | -5.5% | [-6.7%, -4.3%] | 1.2 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | build | ordered | baseline | 6 | 23.56 ms | 21.53 ms | 0.91× [0.91, 0.92] | -9.5% | [-10.1%, -8.8%] | 0.6 pts | 0.9 | yes | yes |
| multi-str | u64 | 4096 | valuesFor | ordered | baseline | 12 | 35.5 | 35.5 | 1.02× [1.00, 1.04] | +1.7% | [-0.3%, +3.7%] | 3.7 pts | 0.7 | yes | no |
| multi-str | u64 | 4096 | valuesBetween | ordered | baseline | 12 | 2675 | 2594 | 0.98× [0.96, 1.00] | -2.1% | [-4.4%, +0.1%] | 2.6 pts | 0.7 | no | no |
| multi-str | u64 | 4096 | churn | ordered | baseline | 12 | 70.8 | 62.8 | 0.89× [0.88, 0.89] | -12.4% | [-13.0%, -11.7%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | u64 | 4096 | build | ordered | baseline | 12 | 7.61 ms | 6.78 ms | 0.89× [0.89, 0.90] | -12.0% | [-12.6%, -11.3%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | u64 | 16384 | valuesFor | ordered | baseline | 6 | 47.9 | 46.8 | 0.98× [0.98, 0.98] | -2.2% | [-2.5%, -1.9%] | 0.3 pts | 0.7 | yes | yes |
| multi-str | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3586 | 3539 | 0.99× [0.98, 1.00] | -1.2% | [-2.2%, -0.1%] | 1.0 pts | 1.4 | yes | yes |
| multi-str | u64 | 16384 | churn | ordered | baseline | 6 | 97.5 | 90.0 | 0.92× [0.91, 0.93] | -8.9% | [-10.4%, -7.3%] | 1.5 pts | 1.4 | yes | yes |
| multi-str | u64 | 16384 | build | ordered | baseline | 6 | 37.40 ms | 33.26 ms | 0.89× [0.88, 0.90] | -12.4% | [-13.2%, -11.6%] | 0.7 pts | 0.8 | yes | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | baseline | 12 | 149 | 154 | 1.06× [1.04, 1.08] | +5.5% | [+3.5%, +7.6%] | 2.6 pts | 4.0 | no | yes |
| multi-str | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 7776 | 8642 | 1.11× [1.08, 1.14] | +9.9% | [+7.7%, +12.0%] | 2.5 pts | 3.8 | no | yes |
| multi-str | u64 | 262144 | churn | ordered | baseline | 12 | 398 | 372 | 0.86× [0.81, 0.91] | -16.2% | [-22.8%, -9.5%] | 8.0 pts | 0.6 | no | yes |
| multi-str | url | 4096 | valuesFor | ordered | baseline | 12 | 88.8 | 89.1 | 1.00× [0.99, 1.01] | +0.0% | [-0.9%, +0.9%] | 1.9 pts | 1.7 | yes | no |
| multi-str | url | 4096 | valuesBetween | ordered | baseline | 12 | 4099 | 3779 | 0.92× [0.90, 0.94] | -8.5% | [-10.8%, -6.2%] | 3.1 pts | 1.3 | no | yes |
| multi-str | url | 4096 | prefix | ordered | baseline | 12 | 184 | 174 | 0.94× [0.93, 0.95] | -6.2% | [-7.4%, -5.0%] | 1.5 pts | 1.2 | yes | yes |
| multi-str | url | 4096 | churn | ordered | baseline | 12 | 149 | 138 | 0.92× [0.91, 0.93] | -8.4% | [-9.6%, -7.1%] | 1.4 pts | 1.5 | yes | yes |
| multi-str | url | 4096 | build | ordered | baseline | 12 | 14.56 ms | 13.33 ms | 0.92× [0.91, 0.92] | -9.2% | [-9.6%, -8.8%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | url | 16384 | valuesFor | ordered | baseline | 6 | 113 | 115 | 1.02× [1.00, 1.03] | +1.6% | [+0.3%, +2.9%] | 1.2 pts | 1.7 | yes | yes |
| multi-str | url | 16384 | valuesBetween | ordered | baseline | 6 | 4769 | 4555 | 0.95× [0.95, 0.96] | -4.8% | [-5.6%, -4.1%] | 0.7 pts | 0.9 | yes | yes |
| multi-str | url | 16384 | prefix | ordered | baseline | 6 | 289 | 279 | 0.97× [0.95, 0.98] | -3.6% | [-5.2%, -2.0%] | 1.5 pts | 1.1 | yes | yes |
| multi-str | url | 16384 | churn | ordered | baseline | 6 | 229 | 225 | 0.98× [0.98, 0.99] | -1.6% | [-2.6%, -0.6%] | 1.0 pts | 1.2 | yes | yes |
| multi-str | url | 16384 | build | ordered | baseline | 6 | 75.65 ms | 71.96 ms | 0.95× [0.94, 0.96] | -5.5% | [-6.4%, -4.5%] | 0.9 pts | 1.5 | yes | yes |
| multi-str | url | 262144 | valuesFor | ordered | baseline | 12 | 394 | 444 | 1.13× [1.12, 1.14] | +11.6% | [+10.6%, +12.6%] | 1.6 pts | 2.3 | yes | yes |
| multi-str | url | 262144 | valuesBetween | ordered | baseline | 12 | 12.3 µs | 12.8 µs | 1.04× [1.03, 1.04] | +3.4% | [+3.0%, +3.8%] | 1.0 pts | 1.6 | yes | yes |
| multi-str | url | 262144 | prefix | ordered | baseline | 12 | 2621 | 2765 | 1.05× [1.03, 1.08] | +5.2% | [+2.8%, +7.6%] | 2.3 pts | 1.3 | no | yes |
| multi-str | url | 262144 | churn | ordered | baseline | 12 | 753 | 750 | 1.00× [0.99, 1.01] | +0.1% | [-1.3%, +1.5%] | 2.6 pts | 0.6 | yes | no |
| multi-str | uuid | 4096 | valuesFor | ordered | baseline | 12 | 56.1 | 56.2 | 1.00× [0.99, 1.01] | -0.3% | [-1.4%, +0.8%] | 1.3 pts | 0.6 | yes | no |
| multi-str | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 3247 | 3184 | 0.97× [0.95, 0.99] | -2.8% | [-5.1%, -0.5%] | 2.4 pts | 0.8 | no | yes |
| multi-str | uuid | 4096 | prefix | ordered | baseline | 12 | 90.7 | 88.5 | 0.97× [0.97, 0.98] | -2.7% | [-3.4%, -2.1%] | 1.4 pts | 1.4 | yes | yes |
| multi-str | uuid | 4096 | churn | ordered | baseline | 12 | 99.1 | 86.7 | 0.87× [0.86, 0.87] | -15.4% | [-16.3%, -14.5%] | 0.9 pts | 1.1 | yes | yes |
| multi-str | uuid | 4096 | build | ordered | baseline | 12 | 10.04 ms | 8.92 ms | 0.88× [0.88, 0.89] | -13.1% | [-13.4%, -12.7%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | uuid | 16384 | valuesFor | ordered | baseline | 6 | 66.5 | 65.5 | 0.99× [0.98, 1.00] | -1.1% | [-1.7%, -0.4%] | 0.6 pts | 1.2 | yes | yes |
| multi-str | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3657 | 3607 | 0.99× [0.98, 0.99] | -1.2% | [-1.5%, -0.8%] | 0.3 pts | 0.6 | yes | yes |
| multi-str | uuid | 16384 | prefix | ordered | baseline | 6 | 110 | 108 | 0.98× [0.97, 0.99] | -1.9% | [-2.6%, -1.2%] | 0.7 pts | 1.4 | yes | yes |
| multi-str | uuid | 16384 | churn | ordered | baseline | 6 | 145 | 132 | 0.91× [0.90, 0.91] | -10.1% | [-10.8%, -9.3%] | 0.7 pts | 0.8 | yes | yes |
| multi-str | uuid | 16384 | build | ordered | baseline | 6 | 49.59 ms | 44.32 ms | 0.89× [0.89, 0.90] | -12.0% | [-12.7%, -11.2%] | 0.7 pts | 1.0 | yes | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | baseline | 12 | 238 | 260 | 1.07× [1.04, 1.11] | +6.5% | [+3.5%, +9.6%] | 3.8 pts | 7.8 | no | yes |
| multi-str | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 9156 | 10.2 µs | 1.11× [1.10, 1.12] | +9.8% | [+8.7%, +10.8%] | 1.6 pts | 2.7 | yes | yes |
| multi-str | uuid | 262144 | prefix | ordered | baseline | 12 | 590 | 655 | 1.11× [1.10, 1.13] | +10.1% | [+8.7%, +11.5%] | 1.9 pts | 4.4 | yes | yes |
| multi-str | uuid | 262144 | churn | ordered | baseline | 12 | 537 | 498 | 0.92× [0.90, 0.94] | -8.5% | [-10.6%, -6.4%] | 4.4 pts | 0.6 | no | yes |
| unique-str | email | 4096 | valuesFor | ordered | baseline | 8 | 22.5 | 22.6 | 1.00× [0.98, 1.01] | -0.3% | [-2.0%, +1.4%] | 1.8 pts | 0.9 | yes | no |
| unique-str | email | 4096 | valuesBetween | ordered | baseline | 8 | 1290 | 1176 | 0.92× [0.91, 0.93] | -8.7% | [-9.3%, -8.1%] | 1.0 pts | 0.8 | yes | yes |
| unique-str | email | 4096 | prefix | ordered | baseline | 8 | 57.6 | 54.2 | 0.94× [0.93, 0.96] | -5.9% | [-7.7%, -4.0%] | 2.0 pts | 1.9 | yes | yes |
| unique-str | email | 4096 | churn | ordered | baseline | 8 | 75.8 | 76.9 | 1.01× [1.00, 1.02] | +1.2% | [+0.3%, +2.1%] | 1.0 pts | 0.9 | yes | yes |
| unique-str | email | 4096 | build | ordered | baseline | 8 | 1.11 ms | 1.11 ms | 1.00× [0.98, 1.01] | -0.2% | [-1.6%, +1.1%] | 1.5 pts | 1.3 | yes | no |
| unique-str | email | 16384 | valuesFor | ordered | baseline | 6 | 37.7 | 39.0 | 1.03× [1.02, 1.05] | +3.4% | [+1.8%, +4.9%] | 1.5 pts | 2.2 | yes | yes |
| unique-str | email | 16384 | valuesBetween | ordered | baseline | 6 | 1614 | 1541 | 0.96× [0.95, 0.97] | -4.3% | [-5.4%, -3.1%] | 1.1 pts | 1.2 | yes | yes |
| unique-str | email | 16384 | prefix | ordered | baseline | 6 | 72.0 | 69.5 | 0.96× [0.95, 0.98] | -3.7% | [-5.1%, -2.3%] | 1.3 pts | 2.7 | yes | yes |
| unique-str | email | 16384 | churn | ordered | baseline | 6 | 94.0 | 96.0 | 1.02× [1.01, 1.03] | +1.9% | [+0.8%, +3.0%] | 1.0 pts | 0.9 | yes | yes |
| unique-str | email | 16384 | build | ordered | baseline | 6 | 5.37 ms | 5.28 ms | 0.98× [0.97, 0.98] | -2.4% | [-3.2%, -1.6%] | 0.8 pts | 1.2 | yes | yes |
| unique-str | email | 262144 | valuesFor | ordered | baseline | 12 | 154 | 172 | 1.13× [1.11, 1.15] | +11.5% | [+9.7%, +13.2%] | 2.4 pts | 2.8 | yes | yes |
| unique-str | email | 262144 | valuesBetween | ordered | baseline | 12 | 4109 | 4724 | 1.16× [1.14, 1.17] | +13.5% | [+12.3%, +14.8%] | 1.7 pts | 1.2 | yes | yes |
| unique-str | email | 262144 | prefix | ordered | baseline | 12 | 201 | 224 | 1.12× [1.09, 1.14] | +10.4% | [+8.6%, +12.3%] | 1.8 pts | 2.6 | yes | yes |
| unique-str | email | 262144 | churn | ordered | baseline | 12 | 297 | 306 | 1.03× [1.01, 1.06] | +3.2% | [+1.0%, +5.3%] | 3.2 pts | 1.2 | no | yes |
| unique-str | path | 4096 | valuesFor | ordered | baseline | 6 | 77.7 | 78.6 | 1.01× [0.99, 1.02] | +0.5% | [-0.9%, +2.0%] | 1.4 pts | 1.0 | yes | no |
| unique-str | path | 4096 | valuesBetween | ordered | baseline | 6 | 2097 | 1942 | 0.92× [0.90, 0.93] | -8.9% | [-10.8%, -7.0%] | 1.8 pts | 0.9 | yes | yes |
| unique-str | path | 4096 | prefix | ordered | baseline | 6 | 236 | 220 | 0.94× [0.93, 0.95] | -6.8% | [-8.0%, -5.6%] | 1.1 pts | 0.8 | yes | yes |
| unique-str | path | 4096 | churn | ordered | baseline | 6 | 201 | 201 | 1.00× [0.99, 1.00] | -0.4% | [-1.2%, +0.3%] | 0.7 pts | 0.8 | yes | no |
| unique-str | path | 4096 | build | ordered | baseline | 6 | 2.75 ms | 2.59 ms | 0.94× [0.93, 0.95] | -6.4% | [-7.2%, -5.5%] | 0.8 pts | 0.8 | yes | yes |
| unique-str | path | 16384 | valuesFor | ordered | baseline | 12 | 112 | 111 | 1.00× [0.99, 1.00] | -0.2% | [-0.8%, +0.4%] | 0.9 pts | 1.1 | yes | no |
| unique-str | path | 16384 | valuesBetween | ordered | baseline | 12 | 2630 | 2541 | 0.97× [0.96, 0.97] | -3.4% | [-3.8%, -2.9%] | 0.7 pts | 0.8 | yes | yes |
| unique-str | path | 16384 | prefix | ordered | baseline | 12 | 487 | 469 | 0.97× [0.95, 0.99] | -2.9% | [-4.7%, -1.0%] | 2.2 pts | 1.1 | yes | yes |
| unique-str | path | 16384 | churn | ordered | baseline | 12 | 257 | 267 | 1.05× [1.04, 1.05] | +4.5% | [+3.9%, +5.0%] | 1.2 pts | 1.2 | yes | yes |
| unique-str | path | 16384 | build | ordered | baseline | 12 | 13.36 ms | 13.48 ms | 1.01× [1.00, 1.02] | +0.9% | [+0.4%, +1.5%] | 0.9 pts | 1.1 | yes | yes |
| unique-str | path | 262144 | valuesFor | ordered | baseline | 12 | 346 | 411 | 1.19× [1.18, 1.20] | +16.1% | [+15.5%, +16.7%] | 1.2 pts | 1.3 | yes | yes |
| unique-str | path | 262144 | valuesBetween | ordered | baseline | 12 | 6773 | 7831 | 1.14× [1.13, 1.16] | +12.6% | [+11.6%, +13.7%] | 1.3 pts | 1.5 | yes | yes |
| unique-str | path | 262144 | prefix | ordered | baseline | 12 | 5100 | 5421 | 1.09× [1.07, 1.12] | +8.5% | [+6.4%, +10.7%] | 2.4 pts | 0.7 | no | yes |
| unique-str | path | 262144 | churn | ordered | baseline | 12 | 681 | 752 | 1.10× [1.07, 1.12] | +8.7% | [+6.8%, +10.7%] | 2.0 pts | 0.9 | yes | yes |
| unique-str | str | 4096 | valuesFor | ordered | baseline | 12 | 33.4 | 33.4 | 1.00× [0.98, 1.02] | -0.1% | [-2.5%, +2.3%] | 3.3 pts | 0.9 | no | no |
| unique-str | str | 4096 | valuesBetween | ordered | baseline | 12 | 1639 | 1543 | 0.94× [0.94, 0.95] | -5.9% | [-6.8%, -4.9%] | 1.7 pts | 1.1 | yes | yes |
| unique-str | str | 4096 | prefix | ordered | baseline | 12 | 3193 | 2971 | 0.93× [0.92, 0.94] | -7.2% | [-8.2%, -6.2%] | 1.4 pts | 1.4 | yes | yes |
| unique-str | str | 4096 | churn | ordered | baseline | 12 | 103 | 98.4 | 0.95× [0.92, 0.98] | -5.2% | [-8.2%, -2.3%] | 3.1 pts | 2.3 | no | yes |
| unique-str | str | 4096 | build | ordered | baseline | 12 | 1.47 ms | 1.36 ms | 0.92× [0.91, 0.93] | -8.8% | [-10.1%, -7.4%] | 2.0 pts | 1.6 | yes | yes |
| unique-str | str | 16384 | valuesFor | ordered | baseline | 6 | 49.0 | 51.1 | 1.05× [1.04, 1.07] | +5.0% | [+3.5%, +6.6%] | 1.5 pts | 2.3 | yes | yes |
| unique-str | str | 16384 | valuesBetween | ordered | baseline | 6 | 1675 | 1636 | 0.98× [0.97, 0.99] | -1.9% | [-3.1%, -0.8%] | 1.1 pts | 1.4 | yes | yes |
| unique-str | str | 16384 | prefix | ordered | baseline | 6 | 13.6 µs | 13.5 µs | 0.99× [0.98, 1.01] | -0.5% | [-2.3%, +1.2%] | 1.7 pts | 1.3 | yes | no |
| unique-str | str | 16384 | churn | ordered | baseline | 6 | 121 | 121 | 1.00× [0.99, 1.01] | +0.1% | [-0.6%, +0.8%] | 0.7 pts | 0.6 | yes | no |
| unique-str | str | 16384 | build | ordered | baseline | 6 | 6.87 ms | 6.59 ms | 0.95× [0.95, 0.96] | -4.7% | [-5.6%, -3.8%] | 0.9 pts | 1.1 | yes | yes |
| unique-str | str | 262144 | valuesFor | ordered | baseline | 12 | 156 | 207 | 1.30× [1.26, 1.35] | +23.2% | [+20.7%, +25.7%] | 3.3 pts | 3.1 | no | yes |
| unique-str | str | 262144 | valuesBetween | ordered | baseline | 12 | 4259 | 5628 | 1.31× [1.29, 1.33] | +23.7% | [+22.7%, +24.7%] | 2.2 pts | 2.1 | yes | yes |
| unique-str | str | 262144 | prefix | ordered | baseline | 12 | 634.8 µs | 862.2 µs | 1.35× [1.33, 1.37] | +25.9% | [+24.6%, +27.1%] | 2.2 pts | 1.4 | yes | yes |
| unique-str | str | 262144 | churn | ordered | baseline | 12 | 324 | 351 | 1.07× [1.05, 1.09] | +6.6% | [+5.1%, +8.1%] | 2.9 pts | 1.4 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | baseline | 6 | 39.4 | 42.5 | 1.09× [1.07, 1.11] | +8.1% | [+6.3%, +9.8%] | 1.7 pts | 0.6 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | baseline | 6 | 1777 | 1674 | 0.94× [0.93, 0.96] | -6.1% | [-8.0%, -4.1%] | 1.9 pts | 0.8 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | baseline | 6 | 198 | 185 | 0.94× [0.92, 0.95] | -6.8% | [-8.4%, -5.3%] | 1.5 pts | 1.4 | yes | yes |
| unique-str | street | 4096 | churn | ordered | baseline | 6 | 121 | 117 | 0.97× [0.96, 0.99] | -2.6% | [-4.1%, -1.0%] | 1.5 pts | 1.1 | yes | yes |
| unique-str | street | 4096 | build | ordered | baseline | 6 | 1.70 ms | 1.60 ms | 0.94× [0.94, 0.94] | -6.5% | [-6.9%, -6.0%] | 0.4 pts | 0.7 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | baseline | 6 | 66.0 | 68.5 | 1.03× [1.02, 1.05] | +3.4% | [+2.0%, +4.7%] | 1.3 pts | 1.5 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 2166 | 2122 | 0.98× [0.97, 0.99] | -2.2% | [-3.5%, -0.8%] | 1.3 pts | 1.1 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | baseline | 6 | 637 | 606 | 0.96× [0.95, 0.98] | -4.0% | [-5.5%, -2.5%] | 1.4 pts | 0.7 | yes | yes |
| unique-str | street | 16384 | churn | ordered | baseline | 6 | 159 | 159 | 1.00× [0.99, 1.01] | +0.1% | [-1.0%, +1.1%] | 1.0 pts | 0.8 | yes | no |
| unique-str | street | 16384 | build | ordered | baseline | 6 | 8.42 ms | 8.15 ms | 0.97× [0.96, 0.97] | -3.5% | [-3.9%, -3.0%] | 0.4 pts | 0.6 | yes | yes |
| unique-str | u64 | 4096 | valuesFor | ordered | baseline | 6 | 13.4 | 13.4 | 1.01× [1.00, 1.02] | +1.1% | [+0.0%, +2.3%] | 1.1 pts | 1.3 | yes | yes |
| unique-str | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1063 | 935 | 0.89× [0.87, 0.90] | -12.9% | [-14.7%, -11.1%] | 1.7 pts | 1.6 | yes | yes |
| unique-str | u64 | 4096 | churn | ordered | baseline | 6 | 50.7 | 50.1 | 0.99× [0.98, 1.01] | -0.7% | [-2.1%, +0.6%] | 1.3 pts | 1.4 | yes | no |
| unique-str | u64 | 4096 | build | ordered | baseline | 6 | 820.5 µs | 810.8 µs | 0.98× [0.97, 0.99] | -1.6% | [-2.6%, -0.6%] | 1.0 pts | 1.1 | yes | yes |
| unique-str | u64 | 16384 | valuesFor | ordered | baseline | 12 | 23.2 | 23.0 | 1.00× [0.99, 1.01] | -0.1% | [-0.9%, +0.8%] | 1.4 pts | 2.4 | yes | no |
| unique-str | u64 | 16384 | valuesBetween | ordered | baseline | 12 | 1571 | 1511 | 0.97× [0.95, 0.98] | -3.4% | [-4.7%, -2.1%] | 1.5 pts | 0.9 | yes | yes |
| unique-str | u64 | 16384 | churn | ordered | baseline | 12 | 58.7 | 60.0 | 1.03× [1.00, 1.06] | +3.3% | [+0.5%, +6.0%] | 4.3 pts | 3.9 | no | yes |
| unique-str | u64 | 16384 | build | ordered | baseline | 12 | 3.82 ms | 3.77 ms | 0.99× [0.98, 1.00] | -0.9% | [-1.9%, +0.2%] | 1.5 pts | 1.6 | yes | no |
| unique-str | u64 | 262144 | valuesFor | ordered | baseline | 12 | 54.1 | 87.2 | 1.62× [1.55, 1.68] | +38.1% | [+35.6%, +40.6%] | 2.6 pts | 2.6 | yes | yes |
| unique-str | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 2251 | 3710 | 1.66× [1.61, 1.71] | +39.7% | [+37.9%, +41.5%] | 3.5 pts | 2.6 | yes | yes |
| unique-str | u64 | 262144 | churn | ordered | baseline | 12 | 214 | 233 | 1.09× [1.06, 1.13] | +8.5% | [+5.8%, +11.2%] | 3.5 pts | 1.1 | no | yes |
| unique-str | url | 4096 | valuesFor | ordered | baseline | 12 | 61.6 | 63.8 | 1.05× [1.02, 1.08] | +4.5% | [+1.7%, +7.3%] | 2.7 pts | 1.7 | no | yes |
| unique-str | url | 4096 | valuesBetween | ordered | baseline | 12 | 2078 | 1802 | 0.86× [0.85, 0.88] | -15.8% | [-17.8%, -13.8%] | 2.3 pts | 1.1 | no | yes |
| unique-str | url | 4096 | prefix | ordered | baseline | 12 | 147 | 137 | 0.93× [0.93, 0.94] | -7.0% | [-7.8%, -6.1%] | 1.1 pts | 1.0 | yes | yes |
| unique-str | url | 4096 | churn | ordered | baseline | 12 | 165 | 163 | 0.99× [0.98, 1.00] | -0.9% | [-1.8%, +0.0%] | 1.4 pts | 1.3 | yes | no |
| unique-str | url | 4096 | build | ordered | baseline | 12 | 2.30 ms | 2.17 ms | 0.94× [0.94, 0.95] | -6.1% | [-6.7%, -5.5%] | 0.7 pts | 0.5 | yes | yes |
| unique-str | url | 16384 | valuesFor | ordered | baseline | 10 | 88.9 | 91.6 | 1.03× [1.02, 1.04] | +2.9% | [+2.3%, +3.5%] | 1.8 pts | 2.7 | yes | yes |
| unique-str | url | 16384 | valuesBetween | ordered | baseline | 10 | 2636 | 2324 | 0.88× [0.87, 0.90] | -13.5% | [-15.4%, -11.6%] | 1.8 pts | 1.7 | yes | yes |
| unique-str | url | 16384 | prefix | ordered | baseline | 10 | 218 | 208 | 0.95× [0.94, 0.96] | -5.0% | [-5.9%, -4.1%] | 0.9 pts | 0.9 | yes | yes |
| unique-str | url | 16384 | churn | ordered | baseline | 10 | 209 | 215 | 1.03× [1.02, 1.04] | +3.1% | [+2.2%, +4.0%] | 1.1 pts | 1.0 | yes | yes |
| unique-str | url | 16384 | build | ordered | baseline | 10 | 11.05 ms | 10.89 ms | 0.98× [0.97, 1.00] | -1.6% | [-2.9%, -0.3%] | 1.3 pts | 1.6 | yes | yes |
| unique-str | url | 262144 | valuesFor | ordered | baseline | 12 | 326 | 375 | 1.15× [1.13, 1.18] | +13.4% | [+11.3%, +15.5%] | 3.5 pts | 5.0 | no | yes |
| unique-str | url | 262144 | valuesBetween | ordered | baseline | 12 | 7390 | 7673 | 1.03× [1.02, 1.04] | +3.0% | [+1.8%, +4.2%] | 2.2 pts | 2.6 | yes | yes |
| unique-str | url | 262144 | prefix | ordered | baseline | 12 | 1399 | 1477 | 1.04× [1.03, 1.06] | +3.9% | [+2.6%, +5.3%] | 2.7 pts | 2.1 | yes | yes |
| unique-str | url | 262144 | churn | ordered | baseline | 12 | 594 | 640 | 1.07× [1.05, 1.10] | +6.9% | [+4.8%, +9.0%] | 2.4 pts | 0.9 | no | yes |
| unique-str | uuid | 4096 | valuesFor | ordered | baseline | 6 | 27.2 | 26.9 | 0.99× [0.98, 1.00] | -1.2% | [-2.1%, -0.3%] | 0.8 pts | 0.3 | yes | no |
| unique-str | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1456 | 1359 | 0.94× [0.93, 0.95] | -6.5% | [-7.9%, -5.1%] | 1.3 pts | 0.6 | yes | yes |
| unique-str | uuid | 4096 | prefix | ordered | baseline | 6 | 69.2 | 65.6 | 0.95× [0.93, 0.97] | -5.2% | [-7.0%, -3.4%] | 1.7 pts | 2.5 | yes | yes |
| unique-str | uuid | 4096 | churn | ordered | baseline | 6 | 86.0 | 84.7 | 0.98× [0.97, 0.99] | -2.0% | [-2.9%, -1.2%] | 0.8 pts | 0.9 | yes | no |
| unique-str | uuid | 4096 | build | ordered | baseline | 6 | 1.22 ms | 1.18 ms | 0.97× [0.95, 0.98] | -3.6% | [-5.1%, -2.1%] | 1.4 pts | 1.3 | yes | yes |
| unique-str | uuid | 16384 | valuesFor | ordered | baseline | 6 | 42.2 | 41.1 | 0.97× [0.96, 0.98] | -2.9% | [-4.0%, -1.8%] | 1.1 pts | 1.7 | yes | yes |
| unique-str | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 1592 | 1529 | 0.96× [0.95, 0.97] | -4.1% | [-4.9%, -3.2%] | 0.8 pts | 1.3 | yes | yes |
| unique-str | uuid | 16384 | prefix | ordered | baseline | 6 | 83.5 | 80.3 | 0.96× [0.95, 0.97] | -4.0% | [-5.1%, -2.8%] | 1.1 pts | 2.5 | yes | yes |
| unique-str | uuid | 16384 | churn | ordered | baseline | 6 | 103 | 105 | 1.03× [1.01, 1.04] | +2.6% | [+1.3%, +4.0%] | 1.3 pts | 0.8 | yes | yes |
| unique-str | uuid | 16384 | build | ordered | baseline | 6 | 5.75 ms | 5.66 ms | 0.99× [0.98, 0.99] | -1.4% | [-2.2%, -0.6%] | 0.8 pts | 0.9 | yes | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | baseline | 12 | 163 | 177 | 1.10× [1.08, 1.13] | +9.3% | [+7.3%, +11.2%] | 2.4 pts | 3.6 | yes | yes |
| unique-str | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 4267 | 5250 | 1.23× [1.21, 1.24] | +18.4% | [+17.3%, +19.5%] | 1.7 pts | 2.0 | yes | yes |
| unique-str | uuid | 262144 | prefix | ordered | baseline | 12 | 313 | 362 | 1.16× [1.13, 1.18] | +13.4% | [+11.3%, +15.6%] | 2.1 pts | 3.7 | no | yes |
| unique-str | uuid | 262144 | churn | ordered | baseline | 12 | 292 | 317 | 1.07× [1.06, 1.09] | +6.7% | [+5.3%, +8.0%] | 1.6 pts | 0.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str email n=4096 valuesFor: ordered vs baseline: the pooled interval [-3.04%, 0.99%] includes zero
- multi-str email n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.72%, 0.25%] includes zero
- multi-str email n=262144 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str email n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 prefix: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=16384 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str path n=16384 prefix: ordered vs baseline: the pooled difference of -1.62% does not clear the 2.21% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=16384 prefix: ordered vs baseline: the pooled interval [-5.20%, 1.96%] includes zero
- multi-str path n=262144 valuesFor: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 churn: ordered vs baseline: the pooled difference of 0.61% does not clear the 1.77% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.94% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str path n=262144 churn: ordered vs baseline: the pooled interval [-2.51%, 3.73%] includes zero
- multi-str str n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.01% does not clear the 0.50% noise floor, the bound on what the harness reports between identical code in every process
- multi-str str n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.71%, 1.73%] includes zero
- multi-str str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-4.32%, 0.23%] includes zero
- multi-str str n=4096 prefix: ordered vs baseline: the pooled difference of -0.59% does not clear the 1.87% noise floor, the bound on what the harness reports between identical code in every process
- multi-str str n=4096 prefix: ordered vs baseline: the pooled interval [-4.73%, 3.54%] includes zero
- multi-str str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.20% does not clear the 1.02% noise floor, the bound on what the harness reports between identical code in every process
- multi-str str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.89%, 0.49%] includes zero
- multi-str str n=16384 prefix: ordered vs baseline: the pooled difference of -0.47% does not clear the 0.82% noise floor, the bound on what the harness reports between identical code in every process
- multi-str str n=16384 prefix: ordered vs baseline: the pooled interval [-1.46%, 0.52%] includes zero
- multi-str str n=262144 valuesFor: ordered vs baseline: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.81%, 2.05%] includes zero
- multi-str street n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.28% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.27%, 3.70%] includes zero
- multi-str u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-4.36%, 0.11%] includes zero
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.03% does not clear the 1.07% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.87%, 0.93%] includes zero
- multi-str url n=4096 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str url n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 churn: ordered vs baseline: the pooled difference of 0.07% does not clear the 1.16% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=262144 churn: ordered vs baseline: the pooled interval [-1.34%, 1.47%] includes zero
- multi-str uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.31% does not clear the 0.84% noise floor, the bound on what the harness reports between identical code in every process
- multi-str uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.44%, 0.81%] includes zero
- multi-str uuid n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.24% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 7.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.28% does not clear the 2.25% noise floor, the bound on what the harness reports between identical code in every process
- unique-str email n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.98%, 1.42%] includes zero
- unique-str email n=4096 build: ordered vs baseline: the pooled difference of -0.22% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- unique-str email n=4096 build: ordered vs baseline: the pooled interval [-1.57%, 1.13%] includes zero
- unique-str email n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=16384 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 prefix: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.53% does not clear the 1.44% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str path n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.94%, 2.00%] includes zero
- unique-str path n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str path n=4096 churn: ordered vs baseline: the pooled difference of -0.44% does not clear the 1.22% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=4096 churn: ordered vs baseline: the pooled interval [-1.21%, 0.34%] includes zero
- unique-str path n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.19% does not clear the 0.20% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.81%, 0.43%] includes zero
- unique-str path n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str str n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.09% does not clear the 2.43% noise floor, the bound on what the harness reports between identical code in every process
- unique-str str n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.45%, 2.28%] includes zero
- unique-str str n=4096 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=16384 prefix: ordered vs baseline: the pooled difference of -0.53% does not clear the 1.14% noise floor, the bound on what the harness reports between identical code in every process
- unique-str str n=16384 prefix: ordered vs baseline: the pooled interval [-2.31%, 1.25%] includes zero
- unique-str str n=16384 churn: ordered vs baseline: the pooled difference of 0.08% does not clear the 1.10% noise floor, the bound on what the harness reports between identical code in every process
- unique-str str n=16384 churn: ordered vs baseline: the pooled interval [-0.64%, 0.80%] includes zero
- unique-str str n=262144 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 churn: ordered vs baseline: the pooled difference of 0.07% does not clear the 1.02% noise floor, the bound on what the harness reports between identical code in every process
- unique-str street n=16384 churn: ordered vs baseline: the pooled interval [-1.00%, 1.13%] includes zero
- unique-str u64 n=4096 churn: ordered vs baseline: the pooled difference of -0.71% does not clear the 0.87% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=4096 churn: ordered vs baseline: the pooled interval [-2.06%, 0.65%] includes zero
- unique-str u64 n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.08% does not clear the 0.32% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.93%, 0.77%] includes zero
- unique-str u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=16384 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str u64 n=16384 churn: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=16384 build: ordered vs baseline: the pooled interval [-1.87%, 0.17%] includes zero
- unique-str u64 n=16384 build: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.82% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str url n=4096 churn: ordered vs baseline: the pooled interval [-1.83%, 0.03%] includes zero
- unique-str url n=16384 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -1.20% does not clear the 2.94% noise floor, the bound on what the harness reports between identical code in every process
- unique-str uuid n=4096 prefix: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=4096 churn: ordered vs baseline: the pooled difference of -2.03% does not clear the 2.04% noise floor, the bound on what the harness reports between identical code in every process
- unique-str uuid n=16384 prefix: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
