| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 12 | 47.5 | 46.6 | 0.98× [0.97, 0.99] | -2.3% | [-3.6%, -1.0%] | 1.6 pts | 0.6 | yes | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 12 | 2889 | 2832 | 0.99× [0.98, 0.99] | -1.5% | [-2.3%, -0.7%] | 1.8 pts | 0.6 | yes | yes |
| multi | email | 4096 | prefix | ordered | baseline | 12 | 77.1 | 73.4 | 0.95× [0.93, 0.97] | -5.4% | [-7.5%, -3.3%] | 2.2 pts | 1.8 | no | yes |
| multi | email | 4096 | churn | ordered | baseline | 12 | 67.4 | 63.7 | 0.94× [0.94, 0.95] | -5.9% | [-6.6%, -5.3%] | 0.8 pts | 1.1 | yes | yes |
| multi | email | 4096 | build | ordered | baseline | 12 | 6.82 ms | 6.41 ms | 0.94× [0.93, 0.95] | -6.3% | [-7.3%, -5.3%] | 1.0 pts | 1.4 | yes | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 60.6 | 59.1 | 0.98× [0.97, 0.98] | -2.3% | [-2.9%, -1.8%] | 0.5 pts | 1.0 | yes | yes |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3519 | 3511 | 1.00× [0.99, 1.01] | -0.1% | [-1.2%, +1.0%] | 1.1 pts | 1.8 | yes | no |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 91.0 | 89.2 | 0.98× [0.97, 0.99] | -2.1% | [-2.9%, -1.4%] | 0.7 pts | 1.1 | yes | yes |
| multi | email | 16384 | churn | ordered | baseline | 6 | 94.6 | 90.8 | 0.96× [0.95, 0.98] | -3.9% | [-5.6%, -2.3%] | 1.6 pts | 2.2 | yes | yes |
| multi | email | 16384 | build | ordered | baseline | 6 | 33.25 ms | 31.25 ms | 0.94× [0.94, 0.95] | -6.2% | [-6.7%, -5.6%] | 0.5 pts | 0.9 | yes | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 12 | 208 | 224 | 1.07× [1.03, 1.12] | +6.8% | [+3.2%, +10.4%] | 3.8 pts | 5.5 | no | yes |
| multi | email | 262144 | valuesBetween | ordered | baseline | 12 | 7446 | 8745 | 1.17× [1.16, 1.17] | +14.4% | [+14.0%, +14.7%] | 0.9 pts | 1.4 | yes | yes |
| multi | email | 262144 | prefix | ordered | baseline | 12 | 271 | 305 | 1.14× [1.13, 1.14] | +12.0% | [+11.4%, +12.6%] | 1.1 pts | 2.0 | yes | yes |
| multi | email | 262144 | churn | ordered | baseline | 12 | 398 | 374 | 0.92× [0.90, 0.93] | -9.0% | [-11.0%, -7.1%] | 2.5 pts | 0.5 | yes | yes |
| multi | path | 4096 | valuesFor | ordered | baseline | 12 | 104 | 99.9 | 0.97× [0.96, 0.98] | -3.0% | [-4.5%, -1.5%] | 1.5 pts | 1.1 | yes | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 12 | 3897 | 3813 | 0.98× [0.95, 1.01] | -2.0% | [-4.7%, +0.7%] | 3.6 pts | 1.5 | no | no |
| multi | path | 4096 | prefix | ordered | baseline | 12 | 309 | 292 | 0.95× [0.94, 0.95] | -5.7% | [-6.6%, -4.8%] | 1.2 pts | 0.8 | yes | yes |
| multi | path | 4096 | churn | ordered | baseline | 12 | 148 | 147 | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.6%] | 1.0 pts | 1.2 | yes | no |
| multi | path | 4096 | build | ordered | baseline | 12 | 13.85 ms | 13.51 ms | 0.98× [0.97, 0.98] | -2.6% | [-3.3%, -1.8%] | 0.7 pts | 1.3 | yes | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 12 | 135 | 133 | 0.98× [0.97, 0.98] | -2.4% | [-3.3%, -1.6%] | 1.2 pts | 1.5 | yes | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 12 | 4599 | 4615 | 1.00× [1.00, 1.01] | +0.2% | [-0.4%, +0.8%] | 0.8 pts | 1.3 | yes | no |
| multi | path | 16384 | prefix | ordered | baseline | 12 | 713 | 715 | 0.98× [0.96, 1.00] | -2.5% | [-4.6%, -0.3%] | 2.6 pts | 1.4 | no | yes |
| multi | path | 16384 | churn | ordered | baseline | 12 | 205 | 210 | 1.02× [1.01, 1.03] | +1.8% | [+0.7%, +3.0%] | 1.5 pts | 2.1 | yes | yes |
| multi | path | 16384 | build | ordered | baseline | 12 | 70.29 ms | 70.29 ms | 1.01× [1.00, 1.01] | +0.6% | [+0.0%, +1.2%] | 0.7 pts | 1.7 | yes | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 12 | 392 | 454 | 1.14× [1.13, 1.15] | +12.5% | [+11.7%, +13.2%] | 0.9 pts | 1.2 | yes | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 12 | 10.5 µs | 11.6 µs | 1.10× [1.09, 1.11] | +8.8% | [+7.8%, +9.7%] | 1.1 pts | 1.6 | yes | yes |
| multi | path | 262144 | prefix | ordered | baseline | 12 | 9958 | 10.5 µs | 1.06× [1.03, 1.10] | +6.1% | [+2.9%, +9.2%] | 4.2 pts | 0.9 | no | yes |
| multi | path | 262144 | churn | ordered | baseline | 12 | 705 | 726 | 1.04× [1.01, 1.08] | +4.3% | [+1.3%, +7.2%] | 6.0 pts | 1.2 | no | yes |
| multi | str | 4096 | valuesFor | ordered | baseline | 12 | 61.3 | 60.2 | 0.96× [0.93, 1.00] | -3.7% | [-7.6%, +0.2%] | 4.1 pts | 2.3 | no | no |
| multi | str | 4096 | valuesBetween | ordered | baseline | 12 | 3352 | 3275 | 0.98× [0.96, 1.00] | -1.9% | [-4.3%, +0.5%] | 2.5 pts | 0.8 | no | no |
| multi | str | 4096 | prefix | ordered | baseline | 12 | 6999 | 6842 | 0.98× [0.95, 1.02] | -1.7% | [-4.9%, +1.6%] | 4.8 pts | 1.2 | no | no |
| multi | str | 4096 | churn | ordered | baseline | 12 | 84.8 | 81.9 | 0.96× [0.95, 0.97] | -4.1% | [-4.9%, -3.2%] | 1.1 pts | 1.5 | yes | yes |
| multi | str | 4096 | build | ordered | baseline | 12 | 8.47 ms | 7.92 ms | 0.94× [0.93, 0.94] | -6.6% | [-7.0%, -6.2%] | 0.7 pts | 1.2 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 73.1 | 74.3 | 1.01× [1.01, 1.02] | +1.2% | [+0.7%, +1.7%] | 0.5 pts | 0.8 | yes | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3652 | 3627 | 1.00× [0.99, 1.00] | -0.2% | [-1.0%, +0.5%] | 0.7 pts | 1.0 | yes | no |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 34.1 µs | 33.8 µs | 0.99× [0.98, 0.99] | -1.5% | [-2.0%, -0.9%] | 0.6 pts | 0.7 | yes | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 109 | 108 | 0.99× [0.98, 1.00] | -1.2% | [-2.1%, -0.3%] | 0.9 pts | 1.2 | yes | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 39.34 ms | 38.10 ms | 0.97× [0.96, 0.98] | -3.3% | [-4.3%, -2.2%] | 1.0 pts | 1.9 | yes | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 12 | 227 | 270 | 1.17× [1.16, 1.19] | +14.8% | [+13.9%, +15.7%] | 1.4 pts | 2.1 | yes | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 12 | 8012 | 9462 | 1.18× [1.17, 1.19] | +15.3% | [+14.5%, +16.1%] | 1.1 pts | 1.5 | yes | yes |
| multi | str | 262144 | prefix | ordered | baseline | 12 | 1.26 ms | 1.51 ms | 1.19× [1.18, 1.20] | +15.9% | [+15.1%, +16.6%] | 0.8 pts | 1.0 | yes | yes |
| multi | str | 262144 | churn | ordered | baseline | 12 | 423 | 413 | 0.97× [0.95, 0.98] | -3.5% | [-5.3%, -1.6%] | 2.8 pts | 0.6 | yes | yes |
| multi | street | 4096 | valuesFor | ordered | baseline | 12 | 48.8 | 48.7 | 1.00× [0.99, 1.01] | -0.1% | [-1.3%, +1.2%] | 2.4 pts | 0.9 | yes | no |
| multi | street | 4096 | valuesBetween | ordered | baseline | 12 | 2286 | 2166 | 0.94× [0.92, 0.96] | -6.1% | [-8.1%, -4.0%] | 2.8 pts | 1.0 | no | yes |
| multi | street | 4096 | prefix | ordered | baseline | 12 | 224 | 212 | 0.95× [0.93, 0.96] | -5.8% | [-7.5%, -4.1%] | 2.1 pts | 1.5 | yes | yes |
| multi | street | 4096 | churn | ordered | baseline | 12 | 94.1 | 88.3 | 0.93× [0.92, 0.94] | -7.1% | [-8.2%, -6.0%] | 1.3 pts | 1.2 | yes | yes |
| multi | street | 4096 | build | ordered | baseline | 12 | 3.68 ms | 3.36 ms | 0.91× [0.90, 0.93] | -9.6% | [-11.3%, -7.9%] | 1.8 pts | 1.7 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 74.6 | 73.4 | 0.98× [0.97, 0.99] | -2.1% | [-3.0%, -1.1%] | 0.9 pts | 1.2 | yes | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2870 | 2755 | 0.96× [0.95, 0.98] | -3.9% | [-5.7%, -2.1%] | 1.7 pts | 1.9 | yes | yes |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 817 | 772 | 0.94× [0.94, 0.95] | -5.9% | [-6.9%, -4.8%] | 1.0 pts | 0.5 | yes | yes |
| multi | street | 16384 | churn | ordered | baseline | 6 | 128 | 119 | 0.94× [0.92, 0.95] | -6.8% | [-8.4%, -5.2%] | 1.5 pts | 1.3 | yes | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 19.91 ms | 18.21 ms | 0.92× [0.92, 0.93] | -8.5% | [-9.1%, -7.9%] | 0.6 pts | 0.8 | yes | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 12 | 34.7 | 31.6 | 0.92× [0.90, 0.94] | -8.6% | [-10.8%, -6.5%] | 2.8 pts | 0.7 | no | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 12 | 2593 | 2544 | 0.99× [0.95, 1.02] | -1.4% | [-4.7%, +2.0%] | 4.1 pts | 1.0 | no | no |
| multi | u64 | 4096 | churn | ordered | baseline | 12 | 49.3 | 46.6 | 0.94× [0.94, 0.95] | -6.2% | [-6.8%, -5.6%] | 0.8 pts | 0.9 | yes | yes |
| multi | u64 | 4096 | build | ordered | baseline | 12 | 5.14 ms | 4.79 ms | 0.93× [0.93, 0.94] | -7.0% | [-7.8%, -6.2%] | 0.9 pts | 0.9 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 46.7 | 44.8 | 0.96× [0.95, 0.97] | -4.3% | [-5.3%, -3.2%] | 1.0 pts | 2.2 | yes | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3499 | 3502 | 1.00× [0.99, 1.01] | -0.0% | [-0.8%, +0.7%] | 0.7 pts | 0.9 | yes | no |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 63.9 | 57.4 | 0.90× [0.88, 0.91] | -11.3% | [-13.1%, -9.4%] | 1.8 pts | 1.8 | yes | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 24.95 ms | 22.77 ms | 0.91× [0.90, 0.91] | -10.3% | [-11.2%, -9.4%] | 0.9 pts | 1.7 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 12 | 123 | 133 | 1.09× [1.06, 1.12] | +8.3% | [+5.8%, +10.8%] | 2.8 pts | 3.7 | no | yes |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 6479 | 7548 | 1.16× [1.15, 1.18] | +13.9% | [+12.8%, +15.0%] | 1.4 pts | 2.6 | yes | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 12 | 288 | 267 | 0.92× [0.90, 0.95] | -8.5% | [-11.7%, -5.3%] | 5.5 pts | 0.8 | no | yes |
| multi | url | 4096 | valuesFor | ordered | baseline | 10 | 85.4 | 85.3 | 1.00× [0.99, 1.02] | +0.2% | [-1.0%, +1.5%] | 1.4 pts | 1.2 | yes | no |
| multi | url | 4096 | valuesBetween | ordered | baseline | 10 | 3748 | 3636 | 0.97× [0.95, 0.99] | -3.4% | [-5.4%, -1.4%] | 3.1 pts | 1.3 | yes | yes |
| multi | url | 4096 | prefix | ordered | baseline | 10 | 180 | 169 | 0.94× [0.93, 0.95] | -6.6% | [-7.3%, -5.8%] | 0.8 pts | 0.7 | yes | yes |
| multi | url | 4096 | churn | ordered | baseline | 10 | 123 | 121 | 0.98× [0.98, 0.99] | -1.7% | [-2.2%, -1.1%] | 0.9 pts | 1.0 | yes | yes |
| multi | url | 4096 | build | ordered | baseline | 10 | 11.54 ms | 11.20 ms | 0.97× [0.96, 0.98] | -3.0% | [-3.7%, -2.2%] | 0.8 pts | 1.4 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 110 | 110 | 1.00× [1.00, 1.01] | +0.4% | [-0.4%, +1.2%] | 0.7 pts | 1.2 | yes | no |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4393 | 4379 | 0.99× [0.98, 1.01] | -0.5% | [-1.6%, +0.6%] | 1.1 pts | 1.8 | yes | no |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 283 | 271 | 0.96× [0.94, 0.98] | -4.3% | [-6.2%, -2.5%] | 1.8 pts | 1.3 | yes | yes |
| multi | url | 16384 | churn | ordered | baseline | 6 | 174 | 175 | 1.01× [1.00, 1.02] | +0.7% | [-0.1%, +1.6%] | 0.8 pts | 1.5 | yes | no |
| multi | url | 16384 | build | ordered | baseline | 6 | 57.96 ms | 57.62 ms | 0.99× [0.99, 1.00] | -0.9% | [-1.3%, -0.4%] | 0.4 pts | 0.9 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 12 | 358 | 416 | 1.16× [1.14, 1.18] | +13.5% | [+12.0%, +15.1%] | 2.3 pts | 3.4 | yes | yes |
| multi | url | 262144 | valuesBetween | ordered | baseline | 12 | 10.6 µs | 11.4 µs | 1.08× [1.07, 1.08] | +7.1% | [+6.4%, +7.8%] | 0.9 pts | 1.6 | yes | yes |
| multi | url | 262144 | prefix | ordered | baseline | 12 | 2333 | 2455 | 1.05× [1.03, 1.07] | +4.7% | [+2.9%, +6.6%] | 1.8 pts | 1.1 | yes | yes |
| multi | url | 262144 | churn | ordered | baseline | 12 | 652 | 649 | 0.98× [0.96, 1.00] | -2.0% | [-3.9%, -0.1%] | 6.3 pts | 1.3 | yes | yes |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 12 | 53.8 | 52.5 | 0.98× [0.97, 0.98] | -2.3% | [-2.9%, -1.6%] | 1.3 pts | 0.6 | yes | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 3164 | 3113 | 0.99× [0.96, 1.01] | -1.3% | [-4.1%, +1.4%] | 3.0 pts | 0.9 | no | no |
| multi | uuid | 4096 | prefix | ordered | baseline | 12 | 89.4 | 85.3 | 0.96× [0.95, 0.97] | -4.3% | [-5.2%, -3.4%] | 1.7 pts | 1.5 | yes | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 12 | 78.0 | 73.2 | 0.94× [0.93, 0.95] | -6.6% | [-7.4%, -5.7%] | 1.2 pts | 1.7 | yes | yes |
| multi | uuid | 4096 | build | ordered | baseline | 12 | 7.63 ms | 7.15 ms | 0.94× [0.93, 0.94] | -6.7% | [-7.3%, -6.1%] | 0.8 pts | 1.0 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 65.3 | 64.0 | 0.98× [0.97, 0.99] | -2.1% | [-2.7%, -1.4%] | 0.6 pts | 1.6 | yes | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3590 | 3581 | 0.99× [0.99, 1.00] | -0.7% | [-1.4%, +0.1%] | 0.7 pts | 1.2 | yes | no |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 108 | 105 | 0.98× [0.97, 0.99] | -2.3% | [-3.2%, -1.4%] | 0.9 pts | 1.3 | yes | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 102 | 97.1 | 0.95× [0.94, 0.96] | -5.0% | [-6.2%, -3.9%] | 1.1 pts | 1.4 | yes | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 36.15 ms | 34.17 ms | 0.95× [0.94, 0.96] | -5.5% | [-6.4%, -4.6%] | 0.9 pts | 1.5 | yes | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 12 | 221 | 242 | 1.10× [1.08, 1.11] | +8.7% | [+7.4%, +10.0%] | 1.4 pts | 2.1 | yes | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 8135 | 9326 | 1.15× [1.13, 1.16] | +12.7% | [+11.9%, +13.6%] | 1.3 pts | 2.1 | yes | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 12 | 532 | 607 | 1.13× [1.12, 1.14] | +11.7% | [+10.7%, +12.6%] | 1.0 pts | 1.9 | yes | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 12 | 441 | 401 | 0.90× [0.86, 0.94] | -10.9% | [-15.9%, -5.8%] | 5.5 pts | 0.8 | no | yes |
| unique | email | 4096 | valuesFor | ordered | baseline | 12 | 22.3 | 21.7 | 0.98× [0.94, 1.02] | -2.4% | [-6.5%, +1.7%] | 4.3 pts | 1.9 | no | no |
| unique | email | 4096 | valuesBetween | ordered | baseline | 12 | 1292 | 1238 | 0.96× [0.95, 0.97] | -4.2% | [-5.1%, -3.4%] | 1.0 pts | 1.1 | yes | yes |
| unique | email | 4096 | prefix | ordered | baseline | 12 | 57.4 | 53.4 | 0.93× [0.93, 0.93] | -7.5% | [-7.8%, -7.2%] | 1.2 pts | 1.4 | yes | yes |
| unique | email | 4096 | churn | ordered | baseline | 12 | 70.6 | 69.1 | 0.98× [0.97, 0.99] | -1.9% | [-2.7%, -1.2%] | 0.9 pts | 1.2 | yes | yes |
| unique | email | 4096 | build | ordered | baseline | 12 | 1.03 ms | 1.01 ms | 0.98× [0.97, 0.98] | -2.2% | [-2.9%, -1.5%] | 0.9 pts | 1.0 | yes | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 37.2 | 37.2 | 0.99× [0.98, 1.01] | -0.5% | [-2.2%, +1.1%] | 1.6 pts | 2.3 | yes | no |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 1622 | 1555 | 0.96× [0.95, 0.97] | -4.1% | [-5.4%, -2.8%] | 1.2 pts | 1.5 | yes | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 71.7 | 67.7 | 0.94× [0.93, 0.95] | -6.2% | [-7.2%, -5.1%] | 1.0 pts | 1.5 | yes | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 88.2 | 87.9 | 1.00× [0.99, 1.01] | -0.1% | [-1.2%, +1.1%] | 1.1 pts | 1.3 | yes | no |
| unique | email | 16384 | build | ordered | baseline | 6 | 5.06 ms | 4.88 ms | 0.97× [0.96, 0.97] | -3.5% | [-4.3%, -2.7%] | 0.8 pts | 1.1 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 12 | 143 | 163 | 1.10× [1.06, 1.14] | +9.3% | [+6.0%, +12.6%] | 6.3 pts | 5.6 | no | yes |
| unique | email | 262144 | valuesBetween | ordered | baseline | 12 | 3604 | 4342 | 1.18× [1.16, 1.21] | +15.5% | [+13.5%, +17.5%] | 2.5 pts | 2.4 | yes | yes |
| unique | email | 262144 | prefix | ordered | baseline | 12 | 186 | 211 | 1.11× [1.10, 1.13] | +10.3% | [+9.2%, +11.3%] | 1.5 pts | 1.7 | yes | yes |
| unique | email | 262144 | churn | ordered | baseline | 12 | 276 | 286 | 1.05× [1.03, 1.07] | +4.8% | [+3.2%, +6.3%] | 2.4 pts | 1.2 | yes | yes |
| unique | path | 4096 | valuesFor | ordered | baseline | 12 | 78.4 | 77.3 | 0.98× [0.96, 1.01] | -1.5% | [-4.5%, +1.4%] | 2.9 pts | 1.8 | no | no |
| unique | path | 4096 | valuesBetween | ordered | baseline | 12 | 2043 | 1976 | 0.98× [0.96, 0.99] | -2.5% | [-4.1%, -0.9%] | 2.1 pts | 0.9 | yes | yes |
| unique | path | 4096 | prefix | ordered | baseline | 12 | 236 | 221 | 0.94× [0.93, 0.95] | -6.3% | [-7.1%, -5.4%] | 1.6 pts | 1.3 | yes | yes |
| unique | path | 4096 | churn | ordered | baseline | 12 | 195 | 194 | 0.99× [0.98, 1.00] | -0.7% | [-1.7%, +0.4%] | 1.5 pts | 1.6 | yes | no |
| unique | path | 4096 | build | ordered | baseline | 12 | 2.64 ms | 2.53 ms | 0.96× [0.95, 0.97] | -4.4% | [-5.4%, -3.5%] | 1.1 pts | 1.3 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 12 | 113 | 110 | 0.99× [0.85, 1.19] | -0.6% | [-17.3%, +16.1%] | 15.3 pts | 15.8 | no | no |
| unique | path | 16384 | valuesBetween | ordered | baseline | 12 | 2566 | 2511 | 1.01× [0.90, 1.15] | +1.1% | [-10.8%, +13.0%] | 11.3 pts | 12.0 | no | no |
| unique | path | 16384 | prefix | ordered | baseline | 12 | 495 | 488 | 0.98× [0.90, 1.08] | -1.7% | [-11.2%, +7.7%] | 8.8 pts | 3.3 | no | no |
| unique | path | 16384 | churn | ordered | baseline | 12 | 247 | 256 | 1.03× [1.02, 1.05] | +3.1% | [+1.8%, +4.5%] | 1.5 pts | 1.7 | yes | yes |
| unique | path | 16384 | build | ordered | baseline | 12 | 12.84 ms | 13.03 ms | 1.02× [1.01, 1.03] | +1.9% | [+0.7%, +3.0%] | 1.8 pts | 2.7 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 12 | 330 | 395 | 1.19× [1.16, 1.21] | +15.7% | [+14.1%, +17.4%] | 1.8 pts | 2.1 | yes | yes |
| unique | path | 262144 | valuesBetween | ordered | baseline | 12 | 6309 | 6825 | 1.09× [1.06, 1.11] | +7.9% | [+5.8%, +10.0%] | 2.4 pts | 2.4 | no | yes |
| unique | path | 262144 | prefix | ordered | baseline | 12 | 4802 | 4909 | 1.01× [0.96, 1.07] | +1.3% | [-3.7%, +6.2%] | 5.2 pts | 1.9 | no | no |
| unique | path | 262144 | churn | ordered | baseline | 12 | 657 | 715 | 1.08× [1.06, 1.10] | +7.4% | [+5.7%, +9.0%] | 1.8 pts | 0.8 | yes | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 12 | 32.1 | 33.5 | 1.01× [0.98, 1.04] | +1.0% | [-1.7%, +3.6%] | 3.7 pts | 0.8 | no | no |
| unique | str | 4096 | valuesBetween | ordered | baseline | 12 | 1636 | 1588 | 0.97× [0.96, 0.98] | -3.4% | [-4.3%, -2.5%] | 1.5 pts | 0.8 | yes | yes |
| unique | str | 4096 | prefix | ordered | baseline | 12 | 3204 | 3124 | 0.97× [0.97, 0.98] | -2.6% | [-3.6%, -1.7%] | 1.5 pts | 1.7 | yes | yes |
| unique | str | 4096 | churn | ordered | baseline | 12 | 93.8 | 89.8 | 0.95× [0.95, 0.96] | -4.9% | [-5.1%, -4.6%] | 1.3 pts | 1.5 | yes | yes |
| unique | str | 4096 | build | ordered | baseline | 12 | 1.35 ms | 1.26 ms | 0.93× [0.92, 0.94] | -7.3% | [-8.2%, -6.4%] | 1.0 pts | 0.6 | yes | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 8 | 49.3 | 50.0 | 1.03× [1.01, 1.05] | +2.6% | [+0.9%, +4.4%] | 1.9 pts | 2.3 | yes | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 8 | 1661 | 1658 | 1.00× [0.99, 1.01] | -0.2% | [-1.1%, +0.7%] | 1.0 pts | 1.5 | yes | no |
| unique | str | 16384 | prefix | ordered | baseline | 8 | 13.5 µs | 13.3 µs | 0.99× [0.97, 1.00] | -1.4% | [-2.6%, -0.2%] | 1.4 pts | 1.1 | yes | yes |
| unique | str | 16384 | churn | ordered | baseline | 8 | 112 | 112 | 1.01× [1.00, 1.01] | +0.7% | [-0.1%, +1.4%] | 2.7 pts | 2.7 | yes | no |
| unique | str | 16384 | build | ordered | baseline | 8 | 6.36 ms | 6.10 ms | 0.96× [0.95, 0.96] | -4.6% | [-5.0%, -4.1%] | 0.5 pts | 0.9 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 12 | 141 | 197 | 1.40× [1.35, 1.47] | +28.8% | [+25.8%, +31.8%] | 2.9 pts | 2.6 | no | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 12 | 3340 | 4831 | 1.44× [1.38, 1.51] | +30.7% | [+27.7%, +33.6%] | 3.3 pts | 3.3 | yes | yes |
| unique | str | 262144 | prefix | ordered | baseline | 12 | 483.8 µs | 740.7 µs | 1.51× [1.47, 1.55] | +33.8% | [+32.0%, +35.6%] | 3.1 pts | 1.4 | yes | yes |
| unique | str | 262144 | churn | ordered | baseline | 12 | 321 | 340 | 1.07× [1.04, 1.10] | +6.3% | [+3.9%, +8.8%] | 3.1 pts | 1.4 | no | yes |
| unique | street | 4096 | valuesFor | ordered | baseline | 12 | 40.2 | 41.4 | 1.03× [1.00, 1.05] | +2.6% | [+0.1%, +5.1%] | 3.8 pts | 1.5 | no | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 12 | 1788 | 1690 | 0.95× [0.93, 0.97] | -5.8% | [-8.0%, -3.5%] | 2.7 pts | 1.3 | no | yes |
| unique | street | 4096 | prefix | ordered | baseline | 12 | 199 | 191 | 0.96× [0.95, 0.97] | -4.2% | [-5.3%, -3.1%] | 1.3 pts | 1.1 | yes | yes |
| unique | street | 4096 | churn | ordered | baseline | 12 | 113 | 107 | 0.95× [0.94, 0.96] | -4.9% | [-6.0%, -3.8%] | 1.4 pts | 1.0 | yes | yes |
| unique | street | 4096 | build | ordered | baseline | 12 | 1.61 ms | 1.48 ms | 0.91× [0.90, 0.93] | -9.3% | [-11.4%, -7.3%] | 2.0 pts | 2.4 | no | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 6 | 66.5 | 66.2 | 0.99× [0.98, 1.00] | -0.8% | [-1.7%, +0.0%] | 0.8 pts | 0.8 | yes | no |
| unique | street | 16384 | valuesBetween | ordered | baseline | 6 | 2178 | 2097 | 0.96× [0.96, 0.97] | -3.9% | [-4.6%, -3.2%] | 0.7 pts | 0.6 | yes | yes |
| unique | street | 16384 | prefix | ordered | baseline | 6 | 635 | 606 | 0.96× [0.95, 0.98] | -3.9% | [-5.2%, -2.5%] | 1.3 pts | 0.7 | yes | yes |
| unique | street | 16384 | churn | ordered | baseline | 6 | 147 | 140 | 0.95× [0.94, 0.96] | -5.5% | [-6.4%, -4.6%] | 0.8 pts | 0.7 | yes | yes |
| unique | street | 16384 | build | ordered | baseline | 6 | 8.19 ms | 7.60 ms | 0.93× [0.92, 0.95] | -7.5% | [-9.3%, -5.7%] | 1.7 pts | 2.8 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 13.2 | 13.0 | 0.98× [0.97, 0.98] | -2.3% | [-3.1%, -1.6%] | 0.7 pts | 0.8 | yes | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1065 | 1007 | 0.95× [0.94, 0.96] | -5.4% | [-6.2%, -4.6%] | 0.8 pts | 0.9 | yes | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 46.0 | 43.6 | 0.95× [0.94, 0.95] | -5.5% | [-6.1%, -4.8%] | 0.6 pts | 0.9 | yes | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 762.6 µs | 723.0 µs | 0.95× [0.94, 0.95] | -5.7% | [-6.3%, -5.1%] | 0.6 pts | 0.7 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 10 | 23.3 | 22.5 | 0.97× [0.95, 0.99] | -2.9% | [-4.7%, -1.1%] | 1.8 pts | 3.9 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 10 | 1595 | 1540 | 0.97× [0.96, 0.98] | -3.1% | [-4.4%, -1.8%] | 1.8 pts | 0.9 | yes | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 10 | 53.6 | 50.0 | 0.94× [0.92, 0.95] | -6.8% | [-8.2%, -5.4%] | 1.4 pts | 1.1 | yes | yes |
| unique | u64 | 16384 | build | ordered | baseline | 10 | 3.52 ms | 3.32 ms | 0.94× [0.93, 0.95] | -6.6% | [-7.5%, -5.7%] | 1.0 pts | 1.3 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 12 | 55.5 | 75.1 | 1.34× [1.29, 1.40] | +25.5% | [+22.6%, +28.3%] | 4.3 pts | 4.0 | no | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 2088 | 3009 | 1.42× [1.37, 1.47] | +29.5% | [+27.2%, +31.8%] | 4.1 pts | 3.4 | yes | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 12 | 206 | 209 | 1.02× [1.00, 1.05] | +2.3% | [-0.4%, +5.1%] | 3.8 pts | 1.2 | no | no |
| unique | url | 4096 | valuesFor | ordered | baseline | 12 | 61.0 | 62.3 | 1.02× [1.00, 1.04] | +1.8% | [-0.1%, +3.6%] | 1.8 pts | 1.0 | yes | no |
| unique | url | 4096 | valuesBetween | ordered | baseline | 12 | 1933 | 1815 | 0.94× [0.93, 0.95] | -6.4% | [-7.9%, -4.9%] | 2.5 pts | 1.3 | yes | yes |
| unique | url | 4096 | prefix | ordered | baseline | 12 | 147 | 136 | 0.92× [0.91, 0.94] | -8.3% | [-9.8%, -6.8%] | 1.8 pts | 1.8 | yes | yes |
| unique | url | 4096 | churn | ordered | baseline | 12 | 155 | 157 | 1.02× [1.01, 1.02] | +1.7% | [+1.1%, +2.3%] | 0.7 pts | 0.7 | yes | yes |
| unique | url | 4096 | build | ordered | baseline | 12 | 2.18 ms | 2.10 ms | 0.97× [0.96, 0.98] | -3.3% | [-4.1%, -2.5%] | 1.7 pts | 1.4 | yes | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 87.9 | 89.2 | 1.02× [1.01, 1.03] | +1.7% | [+0.7%, +2.7%] | 1.0 pts | 1.4 | yes | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 2401 | 2315 | 0.96× [0.95, 0.98] | -3.7% | [-4.8%, -2.5%] | 1.1 pts | 1.2 | yes | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 230 | 224 | 0.97× [0.95, 0.98] | -3.6% | [-4.9%, -2.2%] | 1.3 pts | 1.5 | yes | no |
| unique | url | 16384 | churn | ordered | baseline | 6 | 199 | 206 | 1.05× [1.03, 1.06] | +4.5% | [+3.2%, +5.8%] | 1.2 pts | 1.2 | yes | yes |
| unique | url | 16384 | build | ordered | baseline | 6 | 10.43 ms | 10.43 ms | 1.00× [0.99, 1.01] | +0.2% | [-0.7%, +1.1%] | 0.8 pts | 1.2 | yes | no |
| unique | url | 262144 | valuesFor | ordered | baseline | 12 | 314 | 369 | 1.16× [1.14, 1.18] | +14.0% | [+12.5%, +15.5%] | 2.7 pts | 3.8 | yes | yes |
| unique | url | 262144 | valuesBetween | ordered | baseline | 12 | 6489 | 6779 | 1.04× [1.04, 1.05] | +4.3% | [+3.6%, +4.9%] | 1.1 pts | 1.3 | yes | yes |
| unique | url | 262144 | prefix | ordered | baseline | 12 | 1343 | 1406 | 1.06× [1.04, 1.07] | +5.5% | [+4.1%, +6.8%] | 1.7 pts | 1.3 | yes | yes |
| unique | url | 262144 | churn | ordered | baseline | 12 | 570 | 604 | 1.07× [1.04, 1.11] | +7.0% | [+4.1%, +9.8%] | 3.4 pts | 1.2 | no | yes |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 10 | 26.9 | 26.0 | 0.98× [0.97, 0.99] | -2.4% | [-3.6%, -1.2%] | 2.5 pts | 0.8 | yes | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 10 | 1468 | 1410 | 0.96× [0.94, 0.98] | -4.0% | [-5.9%, -2.1%] | 1.9 pts | 1.1 | yes | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 10 | 68.5 | 63.8 | 0.93× [0.93, 0.93] | -7.6% | [-7.8%, -7.4%] | 0.5 pts | 0.6 | yes | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 10 | 81.7 | 84.5 | 1.03× [1.03, 1.04] | +3.3% | [+2.5%, +4.1%] | 0.8 pts | 1.0 | yes | yes |
| unique | uuid | 4096 | build | ordered | baseline | 10 | 1.15 ms | 1.19 ms | 1.03× [1.03, 1.04] | +3.2% | [+2.5%, +3.9%] | 2.2 pts | 2.0 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 12 | 42.2 | 39.8 | 0.95× [0.94, 0.96] | -5.1% | [-6.2%, -4.0%] | 1.3 pts | 1.8 | yes | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 12 | 1610 | 1580 | 0.98× [0.97, 0.99] | -1.9% | [-2.7%, -1.2%] | 0.9 pts | 1.3 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 12 | 83.9 | 79.4 | 0.95× [0.92, 0.99] | -4.7% | [-8.6%, -0.8%] | 3.7 pts | 6.6 | no | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 12 | 97.4 | 102 | 1.06× [1.05, 1.06] | +5.5% | [+5.1%, +5.8%] | 1.0 pts | 1.1 | yes | yes |
| unique | uuid | 16384 | build | ordered | baseline | 12 | 5.55 ms | 5.69 ms | 1.03× [1.02, 1.03] | +2.8% | [+2.3%, +3.4%] | 0.9 pts | 1.1 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 12 | 160 | 177 | 1.11× [1.09, 1.13] | +10.1% | [+8.5%, +11.7%] | 2.1 pts | 2.5 | yes | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 4256 | 5166 | 1.19× [1.17, 1.22] | +16.2% | [+14.5%, +17.9%] | 2.0 pts | 2.0 | yes | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 12 | 320 | 363 | 1.13× [1.10, 1.16] | +11.3% | [+8.9%, +13.6%] | 2.8 pts | 4.1 | no | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 12 | 321 | 336 | 1.04× [1.02, 1.06] | +3.8% | [+1.6%, +6.0%] | 3.6 pts | 1.4 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.68% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.12% does not clear the 0.36% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.24%, 1.00%] includes zero
- multi email n=16384 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesFor: ordered vs baseline: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 prefix: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-4.74%, 0.74%] includes zero
- multi path n=4096 churn: ordered vs baseline: the pooled difference of -0.21% does not clear the 0.45% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 churn: ordered vs baseline: the pooled interval [-1.02%, 0.60%] includes zero
- multi path n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.23% does not clear the 0.31% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.38%, 0.84%] includes zero
- multi path n=16384 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-7.55%, 0.25%] includes zero
- multi str n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-4.25%, 0.46%] includes zero
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of -1.68% does not clear the 2.92% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 prefix: ordered vs baseline: the pooled interval [-4.93%, 1.57%] includes zero
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.25% does not clear the 0.60% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.95%, 0.46%] includes zero
- multi str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.07% does not clear the 0.80% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.33%, 1.18%] includes zero
- multi street n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi street n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.35% does not clear the 1.47% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-4.74%, 2.03%] includes zero
- multi u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.05% does not clear the 0.57% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.82%, 0.73%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.22% does not clear the 0.61% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.04%, 1.48%] includes zero
- multi url n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.19% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.40%, 1.17%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.50% does not clear the 1.03% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.64%, 0.63%] includes zero
- multi url n=16384 churn: ordered vs baseline: the pooled difference of 0.73% does not clear the 1.09% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=16384 churn: ordered vs baseline: the pooled interval [-0.10%, 1.56%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-4.10%, 1.41%] includes zero
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.66% does not clear the 0.69% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.39%, 0.07%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=4096 valuesFor: ordered vs baseline: the pooled interval [-6.50%, 1.65%] includes zero
- unique email n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.18%, 1.10%] includes zero
- unique email n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 churn: ordered vs baseline: the pooled difference of -0.09% does not clear the 1.34% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=16384 churn: ordered vs baseline: the pooled interval [-1.24%, 1.07%] includes zero
- unique email n=262144 valuesFor: ordered vs baseline: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=4096 valuesFor: ordered vs baseline: the pooled interval [-4.51%, 1.44%] includes zero
- unique path n=4096 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=4096 churn: ordered vs baseline: the pooled interval [-1.71%, 0.41%] includes zero
- unique path n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique path n=16384 valuesFor: ordered vs baseline: the pooled interval [-17.27%, 16.08%] includes zero
- unique path n=16384 valuesFor: ordered vs baseline: the processes scatter 15.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-10.76%, 13.03%] includes zero
- unique path n=16384 valuesBetween: ordered vs baseline: the processes scatter 12.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=16384 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 9 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=16384 prefix: ordered vs baseline: the pooled interval [-11.19%, 7.73%] includes zero
- unique path n=16384 prefix: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=16384 prefix: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=16384 build: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 prefix: ordered vs baseline: the pooled difference of 1.26% does not clear the 6.59% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=262144 prefix: ordered vs baseline: the pooled interval [-3.72%, 6.24%] includes zero
- unique str n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.98% does not clear the 1.17% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.67%, 3.63%] includes zero
- unique str n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.22% does not clear the 0.44% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.12%, 0.68%] includes zero
- unique str n=16384 churn: ordered vs baseline: the pooled interval [-0.10%, 1.40%] includes zero
- unique str n=16384 churn: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.76% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.81% does not clear the 0.81% noise floor, the bound on what the harness reports between identical code in every process
- unique street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.66%, 0.04%] includes zero
- unique street n=16384 build: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 churn: ordered vs baseline: the pooled difference of 2.34% does not clear the 2.68% noise floor, the bound on what the harness reports between identical code in every process
- unique u64 n=262144 churn: ordered vs baseline: the pooled interval [-0.44%, 5.12%] includes zero
- unique url n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.05%, 3.63%] includes zero
- unique url n=16384 prefix: ordered vs baseline: the pooled difference of -3.56% does not clear the 4.81% noise floor, the bound on what the harness reports between identical code in every process
- unique url n=16384 build: ordered vs baseline: the pooled difference of 0.18% does not clear the 0.55% noise floor, the bound on what the harness reports between identical code in every process
- unique url n=16384 build: ordered vs baseline: the pooled interval [-0.70%, 1.05%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique uuid n=16384 prefix: ordered vs baseline: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 prefix: ordered vs baseline: 1 processes resolved A as faster and 11 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 prefix: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
