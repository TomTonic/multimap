| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | email | 4096 | valuesFor | ordered | baseline | 12 | 50.9 | 49.9 | 0.98× [0.97, 0.99] | -1.8% | [-3.1%, -0.5%] | 2.2 pts | 1.0 | yes | yes |
| multi-str | email | 4096 | valuesBetween | ordered | baseline | 12 | 3159 | 2882 | 0.92× [0.90, 0.94] | -8.9% | [-11.5%, -6.4%] | 3.6 pts | 1.3 | no | yes |
| multi-str | email | 4096 | prefix | ordered | baseline | 12 | 80.9 | 76.1 | 0.95× [0.93, 0.96] | -5.8% | [-7.1%, -4.5%] | 1.8 pts | 1.5 | yes | yes |
| multi-str | email | 4096 | churn | ordered | baseline | 12 | 85.9 | 83.6 | 0.97× [0.96, 0.98] | -2.9% | [-3.7%, -2.1%] | 1.0 pts | 1.5 | yes | yes |
| multi-str | email | 4096 | build | ordered | baseline | 12 | 8.66 ms | 8.44 ms | 0.97× [0.96, 0.98] | -3.1% | [-3.6%, -2.5%] | 0.5 pts | 1.1 | yes | yes |
| multi-str | email | 16384 | valuesFor | ordered | baseline | 6 | 63.9 | 62.9 | 0.98× [0.98, 0.99] | -1.6% | [-2.6%, -0.6%] | 0.9 pts | 1.5 | yes | yes |
| multi-str | email | 16384 | valuesBetween | ordered | baseline | 6 | 3796 | 3578 | 0.94× [0.93, 0.95] | -6.4% | [-7.6%, -5.2%] | 1.1 pts | 1.4 | yes | yes |
| multi-str | email | 16384 | prefix | ordered | baseline | 6 | 95.7 | 93.5 | 0.97× [0.97, 0.98] | -2.6% | [-3.4%, -1.8%] | 0.8 pts | 1.3 | yes | yes |
| multi-str | email | 16384 | churn | ordered | baseline | 6 | 126 | 123 | 0.98× [0.97, 0.99] | -2.3% | [-3.2%, -1.5%] | 0.8 pts | 0.9 | yes | yes |
| multi-str | email | 16384 | build | ordered | baseline | 6 | 43.56 ms | 42.74 ms | 0.98× [0.97, 0.99] | -2.2% | [-3.1%, -1.2%] | 0.9 pts | 1.7 | yes | yes |
| multi-str | email | 262144 | valuesFor | ordered | baseline | 12 | 249 | 256 | 1.03× [1.01, 1.05] | +3.2% | [+1.2%, +5.2%] | 2.1 pts | 3.1 | yes | yes |
| multi-str | email | 262144 | valuesBetween | ordered | baseline | 12 | 10.2 µs | 10.0 µs | 0.98× [0.97, 1.00] | -1.6% | [-3.6%, +0.4%] | 2.2 pts | 3.0 | yes | no |
| multi-str | email | 262144 | prefix | ordered | baseline | 12 | 341 | 352 | 1.04× [1.02, 1.05] | +3.4% | [+2.0%, +4.9%] | 1.7 pts | 2.7 | yes | yes |
| multi-str | email | 262144 | churn | ordered | baseline | 12 | 490 | 477 | 0.90× [0.87, 0.94] | -10.6% | [-15.3%, -5.9%] | 7.0 pts | 0.6 | no | yes |
| multi-str | path | 4096 | valuesFor | ordered | baseline | 10 | 108 | 105 | 0.97× [0.97, 0.98] | -2.7% | [-3.5%, -1.8%] | 1.8 pts | 1.5 | yes | yes |
| multi-str | path | 4096 | valuesBetween | ordered | baseline | 10 | 4480 | 3976 | 0.89× [0.88, 0.91] | -12.0% | [-14.0%, -10.0%] | 2.9 pts | 1.1 | yes | yes |
| multi-str | path | 4096 | prefix | ordered | baseline | 10 | 331 | 305 | 0.92× [0.91, 0.94] | -8.2% | [-9.8%, -6.5%] | 1.8 pts | 1.1 | yes | yes |
| multi-str | path | 4096 | churn | ordered | baseline | 10 | 169 | 168 | 1.00× [0.99, 1.01] | -0.4% | [-1.4%, +0.6%] | 1.1 pts | 1.2 | yes | no |
| multi-str | path | 4096 | build | ordered | baseline | 10 | 16.05 ms | 15.87 ms | 0.99× [0.98, 1.00] | -1.3% | [-2.2%, -0.5%] | 1.0 pts | 1.5 | yes | yes |
| multi-str | path | 16384 | valuesFor | ordered | baseline | 12 | 141 | 137 | 0.97× [0.96, 0.98] | -2.9% | [-3.7%, -2.2%] | 2.0 pts | 3.1 | yes | yes |
| multi-str | path | 16384 | valuesBetween | ordered | baseline | 12 | 5246 | 4742 | 0.91× [0.90, 0.91] | -10.4% | [-11.0%, -9.8%] | 0.9 pts | 1.3 | yes | yes |
| multi-str | path | 16384 | prefix | ordered | baseline | 12 | 775 | 720 | 0.91× [0.89, 0.94] | -9.7% | [-12.4%, -6.9%] | 3.0 pts | 1.2 | no | yes |
| multi-str | path | 16384 | churn | ordered | baseline | 12 | 259 | 266 | 1.03× [1.02, 1.03] | +2.8% | [+2.4%, +3.2%] | 1.1 pts | 1.8 | yes | yes |
| multi-str | path | 16384 | build | ordered | baseline | 12 | 83.31 ms | 84.74 ms | 1.02× [1.01, 1.02] | +1.6% | [+0.9%, +2.3%] | 0.9 pts | 1.2 | yes | yes |
| multi-str | path | 262144 | valuesFor | ordered | baseline | 12 | 450 | 495 | 1.09× [1.08, 1.10] | +7.9% | [+7.0%, +8.8%] | 1.7 pts | 1.9 | yes | yes |
| multi-str | path | 262144 | valuesBetween | ordered | baseline | 12 | 13.4 µs | 13.3 µs | 1.00× [0.98, 1.01] | -0.4% | [-2.1%, +1.3%] | 2.1 pts | 3.5 | yes | no |
| multi-str | path | 262144 | prefix | ordered | baseline | 12 | 12.1 µs | 11.4 µs | 0.96× [0.94, 0.98] | -4.3% | [-6.2%, -2.3%] | 3.1 pts | 2.1 | yes | yes |
| multi-str | path | 262144 | churn | ordered | baseline | 12 | 794 | 827 | 1.02× [0.99, 1.04] | +1.6% | [-0.8%, +4.0%] | 3.7 pts | 0.9 | no | no |
| multi-str | str | 4096 | valuesFor | ordered | baseline | 12 | 64.8 | 63.1 | 0.98× [0.97, 1.00] | -1.6% | [-2.9%, -0.4%] | 1.7 pts | 1.1 | yes | yes |
| multi-str | str | 4096 | valuesBetween | ordered | baseline | 12 | 3718 | 3374 | 0.92× [0.90, 0.93] | -9.2% | [-11.0%, -7.4%] | 2.7 pts | 1.0 | yes | yes |
| multi-str | str | 4096 | prefix | ordered | baseline | 12 | 7900 | 7209 | 0.92× [0.89, 0.94] | -9.1% | [-11.8%, -6.4%] | 2.7 pts | 0.6 | no | yes |
| multi-str | str | 4096 | churn | ordered | baseline | 12 | 101 | 101 | 0.99× [0.99, 0.99] | -1.1% | [-1.5%, -0.8%] | 0.8 pts | 1.3 | yes | yes |
| multi-str | str | 4096 | build | ordered | baseline | 12 | 10.38 ms | 9.98 ms | 0.97× [0.96, 0.97] | -3.5% | [-4.0%, -3.1%] | 0.6 pts | 1.2 | yes | yes |
| multi-str | str | 16384 | valuesFor | ordered | baseline | 6 | 75.3 | 77.5 | 1.03× [1.02, 1.03] | +2.6% | [+1.8%, +3.3%] | 0.8 pts | 1.3 | yes | yes |
| multi-str | str | 16384 | valuesBetween | ordered | baseline | 6 | 3936 | 3745 | 0.95× [0.94, 0.95] | -5.6% | [-5.9%, -5.2%] | 0.3 pts | 0.5 | yes | yes |
| multi-str | str | 16384 | prefix | ordered | baseline | 6 | 36.5 µs | 34.6 µs | 0.95× [0.94, 0.96] | -5.6% | [-6.9%, -4.2%] | 1.3 pts | 1.3 | yes | yes |
| multi-str | str | 16384 | churn | ordered | baseline | 6 | 141 | 141 | 1.01× [0.99, 1.02] | +0.7% | [-0.7%, +2.1%] | 1.4 pts | 2.0 | yes | no |
| multi-str | str | 16384 | build | ordered | baseline | 6 | 49.24 ms | 49.46 ms | 1.01× [1.00, 1.02] | +0.9% | [-0.5%, +2.2%] | 1.3 pts | 1.9 | yes | no |
| multi-str | str | 262144 | valuesFor | ordered | baseline | 12 | 274 | 290 | 1.07× [1.05, 1.09] | +6.8% | [+5.2%, +8.4%] | 2.0 pts | 3.4 | yes | yes |
| multi-str | str | 262144 | valuesBetween | ordered | baseline | 12 | 10.4 µs | 10.7 µs | 1.03× [1.02, 1.04] | +3.0% | [+2.3%, +3.8%] | 1.6 pts | 2.3 | yes | yes |
| multi-str | str | 262144 | prefix | ordered | baseline | 12 | 1.66 ms | 1.72 ms | 1.04× [1.03, 1.05] | +3.6% | [+2.6%, +4.6%] | 1.4 pts | 2.9 | yes | yes |
| multi-str | str | 262144 | churn | ordered | baseline | 12 | 515 | 516 | 0.96× [0.94, 0.98] | -4.3% | [-6.8%, -1.8%] | 4.4 pts | 0.9 | no | yes |
| multi-str | street | 4096 | valuesFor | ordered | baseline | 8 | 51.8 | 51.8 | 0.99× [0.98, 1.00] | -1.0% | [-2.0%, +0.1%] | 2.2 pts | 0.8 | yes | no |
| multi-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 2439 | 2190 | 0.89× [0.88, 0.91] | -11.9% | [-13.8%, -10.0%] | 1.9 pts | 0.6 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | baseline | 8 | 235 | 210 | 0.89× [0.88, 0.91] | -11.9% | [-13.9%, -9.9%] | 2.0 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | churn | ordered | baseline | 8 | 108 | 107 | 0.99× [0.98, 1.00] | -1.3% | [-2.3%, -0.3%] | 1.0 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | build | ordered | baseline | 8 | 4.17 ms | 3.97 ms | 0.95× [0.95, 0.96] | -5.2% | [-5.8%, -4.5%] | 0.9 pts | 1.8 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | baseline | 10 | 76.8 | 76.1 | 0.99× [0.98, 1.00] | -0.8% | [-2.1%, +0.4%] | 1.2 pts | 1.7 | yes | no |
| multi-str | street | 16384 | valuesBetween | ordered | baseline | 10 | 3038 | 2814 | 0.93× [0.92, 0.94] | -7.7% | [-8.7%, -6.7%] | 1.2 pts | 1.5 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | baseline | 10 | 879 | 801 | 0.91× [0.89, 0.93] | -10.0% | [-11.9%, -8.1%] | 1.9 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | churn | ordered | baseline | 10 | 152 | 148 | 0.98× [0.96, 0.99] | -2.5% | [-3.8%, -1.3%] | 1.2 pts | 1.3 | yes | yes |
| multi-str | street | 16384 | build | ordered | baseline | 10 | 22.39 ms | 21.70 ms | 0.97× [0.97, 0.97] | -3.1% | [-3.6%, -2.7%] | 0.7 pts | 1.0 | yes | yes |
| multi-str | u64 | 4096 | valuesFor | ordered | baseline | 12 | 37.6 | 35.4 | 0.95× [0.93, 0.97] | -5.3% | [-8.0%, -2.7%] | 5.1 pts | 1.0 | no | yes |
| multi-str | u64 | 4096 | valuesBetween | ordered | baseline | 12 | 2783 | 2606 | 0.94× [0.93, 0.95] | -5.9% | [-7.0%, -4.8%] | 2.4 pts | 0.7 | yes | yes |
| multi-str | u64 | 4096 | churn | ordered | baseline | 12 | 64.6 | 63.3 | 0.97× [0.97, 0.98] | -2.6% | [-3.2%, -2.1%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | u64 | 4096 | build | ordered | baseline | 12 | 6.93 ms | 6.82 ms | 0.98× [0.98, 0.99] | -2.0% | [-2.5%, -1.4%] | 0.8 pts | 1.6 | yes | yes |
| multi-str | u64 | 16384 | valuesFor | ordered | baseline | 6 | 48.1 | 47.4 | 0.99× [0.98, 1.00] | -1.3% | [-2.3%, -0.3%] | 0.9 pts | 1.7 | yes | yes |
| multi-str | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3716 | 3568 | 0.96× [0.96, 0.97] | -4.0% | [-4.6%, -3.4%] | 0.6 pts | 0.8 | yes | yes |
| multi-str | u64 | 16384 | churn | ordered | baseline | 6 | 90.1 | 86.9 | 0.96× [0.96, 0.97] | -3.6% | [-4.6%, -2.6%] | 0.9 pts | 1.5 | yes | yes |
| multi-str | u64 | 16384 | build | ordered | baseline | 6 | 34.20 ms | 33.52 ms | 0.98× [0.97, 0.98] | -2.4% | [-3.0%, -1.8%] | 0.6 pts | 0.7 | yes | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | baseline | 12 | 150 | 156 | 1.04× [1.02, 1.06] | +3.9% | [+2.0%, +5.9%] | 2.5 pts | 4.3 | yes | yes |
| multi-str | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 8717 | 8639 | 0.99× [0.98, 1.01] | -0.7% | [-2.2%, +0.8%] | 2.0 pts | 4.1 | yes | no |
| multi-str | u64 | 262144 | churn | ordered | baseline | 12 | 377 | 365 | 0.90× [0.85, 0.96] | -10.5% | [-17.2%, -3.9%] | 8.0 pts | 0.7 | no | yes |
| multi-str | url | 4096 | valuesFor | ordered | baseline | 6 | 91.7 | 91.7 | 1.01× [1.00, 1.02] | +0.6% | [-0.4%, +1.6%] | 0.9 pts | 0.9 | yes | no |
| multi-str | url | 4096 | valuesBetween | ordered | baseline | 6 | 4421 | 3849 | 0.87× [0.86, 0.88] | -15.1% | [-16.1%, -14.2%] | 0.9 pts | 0.4 | yes | yes |
| multi-str | url | 4096 | prefix | ordered | baseline | 6 | 190 | 177 | 0.93× [0.91, 0.94] | -7.7% | [-9.5%, -6.0%] | 1.7 pts | 1.4 | yes | yes |
| multi-str | url | 4096 | churn | ordered | baseline | 6 | 143 | 140 | 0.99× [0.97, 1.00] | -1.4% | [-2.8%, -0.0%] | 1.3 pts | 1.6 | yes | yes |
| multi-str | url | 4096 | build | ordered | baseline | 6 | 13.74 ms | 13.41 ms | 0.98× [0.98, 0.98] | -2.3% | [-2.5%, -2.0%] | 0.2 pts | 0.3 | yes | yes |
| multi-str | url | 16384 | valuesFor | ordered | baseline | 6 | 115 | 115 | 1.01× [0.99, 1.03] | +1.0% | [-0.6%, +2.6%] | 1.5 pts | 1.8 | yes | no |
| multi-str | url | 16384 | valuesBetween | ordered | baseline | 6 | 5120 | 4529 | 0.88× [0.88, 0.89] | -13.1% | [-13.9%, -12.3%] | 0.8 pts | 1.3 | yes | yes |
| multi-str | url | 16384 | prefix | ordered | baseline | 6 | 302 | 281 | 0.93× [0.93, 0.94] | -7.1% | [-8.1%, -6.0%] | 1.0 pts | 0.6 | yes | yes |
| multi-str | url | 16384 | churn | ordered | baseline | 6 | 222 | 226 | 1.02× [1.01, 1.03] | +1.8% | [+0.7%, +2.9%] | 1.0 pts | 1.0 | yes | yes |
| multi-str | url | 16384 | build | ordered | baseline | 6 | 71.66 ms | 71.65 ms | 1.00× [1.00, 1.01] | +0.2% | [-0.4%, +0.9%] | 0.6 pts | 0.9 | yes | no |
| multi-str | url | 262144 | valuesFor | ordered | baseline | 12 | 415 | 454 | 1.09× [1.08, 1.10] | +7.9% | [+7.0%, +8.9%] | 1.7 pts | 2.7 | yes | yes |
| multi-str | url | 262144 | valuesBetween | ordered | baseline | 12 | 13.6 µs | 13.0 µs | 0.95× [0.95, 0.96] | -4.7% | [-5.3%, -4.1%] | 0.6 pts | 0.8 | yes | yes |
| multi-str | url | 262144 | prefix | ordered | baseline | 12 | 3031 | 2894 | 0.96× [0.96, 0.97] | -3.9% | [-4.7%, -3.2%] | 1.2 pts | 0.8 | yes | yes |
| multi-str | url | 262144 | churn | ordered | baseline | 12 | 726 | 739 | 1.00× [0.95, 1.05] | -0.4% | [-5.1%, +4.4%] | 4.8 pts | 1.1 | no | no |
| multi-str | uuid | 4096 | valuesFor | ordered | baseline | 12 | 58.7 | 56.7 | 0.97× [0.95, 0.98] | -3.5% | [-4.8%, -2.1%] | 1.4 pts | 0.8 | yes | yes |
| multi-str | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 3467 | 3182 | 0.92× [0.90, 0.94] | -8.5% | [-10.8%, -6.2%] | 3.6 pts | 1.0 | no | yes |
| multi-str | uuid | 4096 | prefix | ordered | baseline | 12 | 94.8 | 89.6 | 0.95× [0.94, 0.96] | -5.4% | [-6.2%, -4.6%] | 1.1 pts | 1.2 | yes | yes |
| multi-str | uuid | 4096 | churn | ordered | baseline | 12 | 97.3 | 89.1 | 0.91× [0.91, 0.92] | -9.6% | [-10.3%, -8.9%] | 1.0 pts | 1.3 | yes | yes |
| multi-str | uuid | 4096 | build | ordered | baseline | 12 | 9.46 ms | 8.88 ms | 0.94× [0.94, 0.94] | -6.6% | [-6.9%, -6.2%] | 0.5 pts | 0.8 | yes | yes |
| multi-str | uuid | 16384 | valuesFor | ordered | baseline | 6 | 67.7 | 66.1 | 0.98× [0.97, 0.98] | -2.6% | [-3.0%, -2.1%] | 0.4 pts | 0.7 | yes | yes |
| multi-str | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3954 | 3607 | 0.91× [0.91, 0.92] | -9.4% | [-10.3%, -8.5%] | 0.8 pts | 1.2 | yes | yes |
| multi-str | uuid | 16384 | prefix | ordered | baseline | 6 | 114 | 109 | 0.96× [0.95, 0.97] | -4.3% | [-5.1%, -3.5%] | 0.8 pts | 1.3 | yes | yes |
| multi-str | uuid | 16384 | churn | ordered | baseline | 6 | 137 | 135 | 0.98× [0.98, 0.99] | -1.5% | [-2.5%, -0.5%] | 1.0 pts | 1.7 | yes | yes |
| multi-str | uuid | 16384 | build | ordered | baseline | 6 | 45.88 ms | 44.60 ms | 0.96× [0.95, 0.97] | -4.0% | [-5.3%, -2.7%] | 1.3 pts | 1.7 | yes | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | baseline | 12 | 261 | 265 | 1.02× [1.00, 1.04] | +1.5% | [-0.4%, +3.4%] | 2.5 pts | 4.2 | yes | no |
| multi-str | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 10.8 µs | 10.3 µs | 0.95× [0.94, 0.96] | -5.4% | [-6.4%, -4.5%] | 1.3 pts | 1.7 | yes | yes |
| multi-str | uuid | 262144 | prefix | ordered | baseline | 12 | 678 | 666 | 0.98× [0.96, 1.00] | -1.9% | [-4.0%, +0.1%] | 2.2 pts | 5.6 | no | no |
| multi-str | uuid | 262144 | churn | ordered | baseline | 12 | 518 | 496 | 0.93× [0.91, 0.95] | -7.8% | [-9.9%, -5.6%] | 3.1 pts | 0.5 | no | yes |
| unique-str | email | 4096 | valuesFor | ordered | baseline | 12 | 24.4 | 23.1 | 0.96× [0.94, 0.97] | -4.6% | [-6.6%, -2.7%] | 2.6 pts | 0.9 | yes | yes |
| unique-str | email | 4096 | valuesBetween | ordered | baseline | 12 | 1446 | 1182 | 0.82× [0.81, 0.83] | -21.9% | [-23.0%, -20.7%] | 2.1 pts | 1.6 | yes | yes |
| unique-str | email | 4096 | prefix | ordered | baseline | 12 | 60.2 | 54.4 | 0.91× [0.90, 0.91] | -10.4% | [-11.2%, -9.7%] | 1.2 pts | 1.2 | yes | yes |
| unique-str | email | 4096 | churn | ordered | baseline | 12 | 80.7 | 77.1 | 0.96× [0.96, 0.97] | -3.9% | [-4.6%, -3.3%] | 0.8 pts | 0.8 | yes | yes |
| unique-str | email | 4096 | build | ordered | baseline | 12 | 1.15 ms | 1.10 ms | 0.96× [0.95, 0.96] | -4.7% | [-5.3%, -4.1%] | 1.0 pts | 1.1 | yes | yes |
| unique-str | email | 16384 | valuesFor | ordered | baseline | 8 | 40.8 | 39.2 | 0.96× [0.95, 0.97] | -4.2% | [-4.9%, -3.4%] | 1.1 pts | 2.1 | yes | yes |
| unique-str | email | 16384 | valuesBetween | ordered | baseline | 8 | 1821 | 1540 | 0.85× [0.84, 0.86] | -18.1% | [-19.3%, -16.9%] | 1.3 pts | 1.5 | yes | yes |
| unique-str | email | 16384 | prefix | ordered | baseline | 8 | 74.5 | 69.9 | 0.94× [0.92, 0.95] | -6.9% | [-8.4%, -5.4%] | 1.5 pts | 3.7 | yes | yes |
| unique-str | email | 16384 | churn | ordered | baseline | 8 | 103 | 98.5 | 0.96× [0.95, 0.98] | -3.9% | [-5.7%, -2.1%] | 1.9 pts | 1.4 | yes | yes |
| unique-str | email | 16384 | build | ordered | baseline | 8 | 5.58 ms | 5.24 ms | 0.94× [0.94, 0.94] | -6.6% | [-6.8%, -6.3%] | 0.6 pts | 0.9 | yes | yes |
| unique-str | email | 262144 | valuesFor | ordered | baseline | 12 | 172 | 177 | 1.04× [1.02, 1.05] | +3.6% | [+2.4%, +4.8%] | 1.6 pts | 2.0 | yes | yes |
| unique-str | email | 262144 | valuesBetween | ordered | baseline | 12 | 5253 | 5073 | 0.96× [0.95, 0.97] | -4.1% | [-5.4%, -2.8%] | 2.0 pts | 1.2 | yes | yes |
| unique-str | email | 262144 | prefix | ordered | baseline | 12 | 243 | 255 | 1.04× [1.03, 1.06] | +4.2% | [+2.6%, +5.8%] | 1.8 pts | 2.1 | yes | yes |
| unique-str | email | 262144 | churn | ordered | baseline | 12 | 316 | 310 | 0.97× [0.95, 0.99] | -3.1% | [-5.7%, -0.5%] | 5.7 pts | 1.0 | no | yes |
| unique-str | path | 4096 | valuesFor | ordered | baseline | 12 | 80.6 | 78.9 | 0.97× [0.95, 1.00] | -3.0% | [-5.7%, -0.3%] | 2.7 pts | 1.9 | no | yes |
| unique-str | path | 4096 | valuesBetween | ordered | baseline | 12 | 2443 | 1962 | 0.80× [0.79, 0.81] | -24.9% | [-26.4%, -23.4%] | 2.6 pts | 0.9 | yes | yes |
| unique-str | path | 4096 | prefix | ordered | baseline | 12 | 250 | 223 | 0.89× [0.89, 0.90] | -12.2% | [-12.9%, -11.4%] | 1.5 pts | 1.2 | yes | yes |
| unique-str | path | 4096 | churn | ordered | baseline | 12 | 210 | 204 | 0.97× [0.96, 0.98] | -3.0% | [-3.8%, -2.3%] | 1.0 pts | 0.9 | yes | yes |
| unique-str | path | 4096 | build | ordered | baseline | 12 | 2.77 ms | 2.57 ms | 0.93× [0.92, 0.93] | -7.6% | [-8.2%, -7.0%] | 0.8 pts | 0.9 | yes | yes |
| unique-str | path | 16384 | valuesFor | ordered | baseline | 12 | 118 | 115 | 0.94× [0.88, 1.01] | -6.5% | [-14.0%, +1.0%] | 7.8 pts | 9.3 | no | no |
| unique-str | path | 16384 | valuesBetween | ordered | baseline | 12 | 3190 | 2597 | 0.77× [0.72, 0.84] | -29.1% | [-39.1%, -19.1%] | 10.7 pts | 11.8 | no | yes |
| unique-str | path | 16384 | prefix | ordered | baseline | 12 | 564 | 489 | 0.86× [0.83, 0.89] | -16.7% | [-20.6%, -12.7%] | 4.6 pts | 1.4 | no | yes |
| unique-str | path | 16384 | churn | ordered | baseline | 12 | 273 | 273 | 1.01× [1.00, 1.01] | +0.7% | [+0.4%, +1.1%] | 0.6 pts | 0.8 | yes | yes |
| unique-str | path | 16384 | build | ordered | baseline | 12 | 14.06 ms | 13.69 ms | 0.97× [0.97, 0.98] | -3.0% | [-3.5%, -2.5%] | 1.0 pts | 1.1 | yes | yes |
| unique-str | path | 262144 | valuesFor | ordered | baseline | 12 | 363 | 408 | 1.12× [1.11, 1.14] | +11.1% | [+9.8%, +12.4%] | 1.5 pts | 1.8 | yes | yes |
| unique-str | path | 262144 | valuesBetween | ordered | baseline | 12 | 8485 | 8293 | 0.98× [0.97, 0.98] | -2.3% | [-3.0%, -1.6%] | 0.8 pts | 0.9 | yes | yes |
| unique-str | path | 262144 | prefix | ordered | baseline | 12 | 6518 | 6030 | 0.94× [0.92, 0.96] | -6.9% | [-9.2%, -4.5%] | 3.0 pts | 1.1 | no | yes |
| unique-str | path | 262144 | churn | ordered | baseline | 12 | 711 | 756 | 1.07× [1.05, 1.08] | +6.1% | [+4.5%, +7.8%] | 2.1 pts | 1.1 | yes | yes |
| unique-str | str | 4096 | valuesFor | ordered | baseline | 12 | 36.0 | 33.2 | 0.93× [0.91, 0.95] | -7.3% | [-9.3%, -5.2%] | 3.2 pts | 0.8 | no | yes |
| unique-str | str | 4096 | valuesBetween | ordered | baseline | 12 | 1837 | 1547 | 0.85× [0.83, 0.86] | -18.2% | [-19.9%, -16.4%] | 2.0 pts | 1.1 | yes | yes |
| unique-str | str | 4096 | prefix | ordered | baseline | 12 | 3581 | 2977 | 0.83× [0.83, 0.84] | -20.4% | [-21.2%, -19.6%] | 0.9 pts | 0.7 | yes | yes |
| unique-str | str | 4096 | churn | ordered | baseline | 12 | 106 | 97.1 | 0.92× [0.92, 0.93] | -8.3% | [-9.2%, -7.4%] | 1.3 pts | 1.3 | yes | yes |
| unique-str | str | 4096 | build | ordered | baseline | 12 | 1.51 ms | 1.36 ms | 0.90× [0.89, 0.90] | -11.6% | [-12.4%, -10.7%] | 1.1 pts | 0.8 | yes | yes |
| unique-str | str | 16384 | valuesFor | ordered | baseline | 6 | 50.5 | 51.5 | 1.02× [1.01, 1.03] | +2.1% | [+1.4%, +2.8%] | 0.6 pts | 1.1 | yes | yes |
| unique-str | str | 16384 | valuesBetween | ordered | baseline | 6 | 1895 | 1642 | 0.86× [0.86, 0.87] | -15.6% | [-16.5%, -14.8%] | 0.8 pts | 1.0 | yes | yes |
| unique-str | str | 16384 | prefix | ordered | baseline | 6 | 15.8 µs | 13.4 µs | 0.85× [0.84, 0.86] | -17.8% | [-19.4%, -16.1%] | 1.6 pts | 1.2 | yes | yes |
| unique-str | str | 16384 | churn | ordered | baseline | 6 | 126 | 123 | 0.99× [0.97, 1.00] | -1.4% | [-2.7%, -0.1%] | 1.2 pts | 1.1 | yes | yes |
| unique-str | str | 16384 | build | ordered | baseline | 6 | 6.96 ms | 6.50 ms | 0.94× [0.93, 0.94] | -6.9% | [-7.8%, -5.9%] | 0.9 pts | 1.4 | yes | yes |
| unique-str | str | 262144 | valuesFor | ordered | baseline | 12 | 200 | 210 | 1.04× [1.03, 1.06] | +4.3% | [+3.2%, +5.3%] | 1.4 pts | 1.8 | yes | yes |
| unique-str | str | 262144 | valuesBetween | ordered | baseline | 12 | 5645 | 5573 | 0.99× [0.98, 1.01] | -0.9% | [-2.3%, +0.6%] | 1.8 pts | 1.4 | yes | no |
| unique-str | str | 262144 | prefix | ordered | baseline | 12 | 919.6 µs | 913.7 µs | 0.99× [0.98, 1.00] | -0.6% | [-1.8%, +0.5%] | 1.3 pts | 1.4 | yes | no |
| unique-str | str | 262144 | churn | ordered | baseline | 12 | 380 | 386 | 1.02× [0.97, 1.07] | +1.8% | [-2.6%, +6.1%] | 4.7 pts | 1.7 | no | no |
| unique-str | street | 4096 | valuesFor | ordered | baseline | 12 | 43.9 | 43.3 | 1.00× [0.99, 1.02] | +0.3% | [-1.5%, +2.1%] | 2.0 pts | 0.9 | yes | no |
| unique-str | street | 4096 | valuesBetween | ordered | baseline | 12 | 1942 | 1677 | 0.86× [0.85, 0.86] | -16.7% | [-17.7%, -15.7%] | 2.0 pts | 1.0 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | baseline | 12 | 213 | 185 | 0.87× [0.85, 0.88] | -15.2% | [-17.2%, -13.2%] | 2.3 pts | 1.5 | yes | yes |
| unique-str | street | 4096 | churn | ordered | baseline | 12 | 123 | 117 | 0.96× [0.95, 0.96] | -4.4% | [-4.9%, -3.9%] | 0.7 pts | 0.7 | yes | yes |
| unique-str | street | 4096 | build | ordered | baseline | 12 | 1.74 ms | 1.58 ms | 0.91× [0.90, 0.91] | -10.2% | [-11.0%, -9.3%] | 0.9 pts | 1.7 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | baseline | 6 | 69.9 | 69.3 | 0.98× [0.97, 1.00] | -1.5% | [-3.2%, +0.1%] | 1.5 pts | 1.7 | yes | no |
| unique-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 2389 | 2150 | 0.90× [0.89, 0.91] | -11.0% | [-12.0%, -10.0%] | 0.9 pts | 1.1 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | baseline | 6 | 703 | 621 | 0.89× [0.88, 0.90] | -12.4% | [-13.4%, -11.4%] | 0.9 pts | 0.5 | yes | yes |
| unique-str | street | 16384 | churn | ordered | baseline | 6 | 166 | 159 | 0.96× [0.94, 0.98] | -4.1% | [-5.9%, -2.3%] | 1.7 pts | 1.0 | yes | yes |
| unique-str | street | 16384 | build | ordered | baseline | 6 | 8.73 ms | 8.05 ms | 0.93× [0.92, 0.93] | -8.1% | [-9.2%, -7.0%] | 1.1 pts | 1.9 | yes | yes |
| unique-str | u64 | 4096 | valuesFor | ordered | baseline | 6 | 14.8 | 13.4 | 0.91× [0.90, 0.92] | -9.8% | [-10.8%, -8.8%] | 0.9 pts | 0.9 | yes | yes |
| unique-str | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1161 | 940 | 0.80× [0.79, 0.82] | -24.4% | [-26.3%, -22.4%] | 1.8 pts | 1.9 | yes | yes |
| unique-str | u64 | 4096 | churn | ordered | baseline | 6 | 54.4 | 50.7 | 0.93× [0.92, 0.95] | -7.0% | [-8.5%, -5.4%] | 1.5 pts | 1.6 | yes | yes |
| unique-str | u64 | 4096 | build | ordered | baseline | 6 | 873.7 µs | 807.9 µs | 0.92× [0.92, 0.93] | -8.3% | [-9.1%, -7.5%] | 0.8 pts | 1.7 | yes | yes |
| unique-str | u64 | 16384 | valuesFor | ordered | baseline | 12 | 24.8 | 23.2 | 0.94× [0.93, 0.94] | -6.9% | [-7.5%, -6.3%] | 0.8 pts | 1.6 | yes | yes |
| unique-str | u64 | 16384 | valuesBetween | ordered | baseline | 12 | 1696 | 1520 | 0.90× [0.88, 0.92] | -11.2% | [-14.1%, -8.3%] | 2.7 pts | 1.6 | no | yes |
| unique-str | u64 | 16384 | churn | ordered | baseline | 12 | 65.2 | 61.9 | 0.95× [0.94, 0.97] | -5.0% | [-6.9%, -3.2%] | 2.4 pts | 1.3 | yes | yes |
| unique-str | u64 | 16384 | build | ordered | baseline | 12 | 4.02 ms | 3.73 ms | 0.93× [0.93, 0.93] | -7.6% | [-7.9%, -7.3%] | 0.7 pts | 1.0 | yes | yes |
| unique-str | u64 | 262144 | valuesFor | ordered | baseline | 12 | 86.4 | 88.3 | 1.02× [1.01, 1.03] | +1.7% | [+0.6%, +2.8%] | 1.9 pts | 2.8 | yes | yes |
| unique-str | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 4027 | 4050 | 1.00× [0.99, 1.02] | +0.2% | [-1.1%, +1.5%] | 1.6 pts | 1.4 | yes | no |
| unique-str | u64 | 262144 | churn | ordered | baseline | 12 | 240 | 239 | 0.98× [0.94, 1.02] | -2.1% | [-6.7%, +2.4%] | 5.4 pts | 0.9 | no | no |
| unique-str | url | 4096 | valuesFor | ordered | baseline | 10 | 62.5 | 62.7 | 0.99× [0.98, 1.01] | -0.5% | [-2.4%, +1.3%] | 1.7 pts | 1.1 | yes | no |
| unique-str | url | 4096 | valuesBetween | ordered | baseline | 10 | 2342 | 1809 | 0.78× [0.77, 0.79] | -28.6% | [-29.7%, -27.4%] | 2.4 pts | 0.9 | yes | yes |
| unique-str | url | 4096 | prefix | ordered | baseline | 10 | 151 | 137 | 0.90× [0.88, 0.91] | -11.4% | [-13.2%, -9.6%] | 1.8 pts | 1.3 | yes | yes |
| unique-str | url | 4096 | churn | ordered | baseline | 10 | 172 | 167 | 0.97× [0.97, 0.98] | -2.8% | [-3.5%, -2.0%] | 0.8 pts | 0.7 | yes | yes |
| unique-str | url | 4096 | build | ordered | baseline | 10 | 2.35 ms | 2.15 ms | 0.93× [0.92, 0.93] | -7.8% | [-8.2%, -7.5%] | 0.8 pts | 0.5 | yes | yes |
| unique-str | url | 16384 | valuesFor | ordered | baseline | 6 | 91.7 | 91.4 | 1.00× [0.98, 1.01] | -0.3% | [-1.6%, +1.0%] | 1.2 pts | 2.1 | yes | no |
| unique-str | url | 16384 | valuesBetween | ordered | baseline | 6 | 3075 | 2370 | 0.77× [0.77, 0.78] | -29.6% | [-30.6%, -28.7%] | 0.9 pts | 1.0 | yes | yes |
| unique-str | url | 16384 | prefix | ordered | baseline | 6 | 232 | 210 | 0.91× [0.90, 0.91] | -10.3% | [-11.2%, -9.4%] | 0.9 pts | 0.9 | yes | yes |
| unique-str | url | 16384 | churn | ordered | baseline | 6 | 220 | 218 | 1.00× [0.99, 1.01] | -0.2% | [-1.5%, +1.0%] | 1.1 pts | 1.0 | yes | no |
| unique-str | url | 16384 | build | ordered | baseline | 6 | 11.34 ms | 10.89 ms | 0.96× [0.95, 0.97] | -4.2% | [-5.0%, -3.3%] | 0.8 pts | 1.0 | yes | yes |
| unique-str | url | 262144 | valuesFor | ordered | baseline | 8 | 339 | 386 | 1.14× [1.11, 1.16] | +11.9% | [+10.0%, +13.9%] | 2.7 pts | 3.1 | yes | yes |
| unique-str | url | 262144 | valuesBetween | ordered | baseline | 8 | 8470 | 8118 | 0.96× [0.95, 0.96] | -4.4% | [-5.0%, -3.9%] | 0.9 pts | 1.1 | yes | yes |
| unique-str | url | 262144 | prefix | ordered | baseline | 8 | 1697 | 1577 | 0.94× [0.92, 0.95] | -6.9% | [-8.9%, -4.9%] | 2.0 pts | 1.6 | yes | yes |
| unique-str | url | 262144 | churn | ordered | baseline | 8 | 612 | 643 | 1.05× [1.04, 1.07] | +5.0% | [+3.5%, +6.5%] | 1.7 pts | 0.8 | yes | yes |
| unique-str | uuid | 4096 | valuesFor | ordered | baseline | 12 | 30.3 | 27.0 | 0.90× [0.88, 0.91] | -11.3% | [-13.0%, -9.6%] | 2.0 pts | 0.5 | yes | yes |
| unique-str | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 1641 | 1368 | 0.84× [0.82, 0.85] | -19.5% | [-21.4%, -17.5%] | 3.0 pts | 1.2 | yes | yes |
| unique-str | uuid | 4096 | prefix | ordered | baseline | 12 | 73.3 | 66.0 | 0.90× [0.90, 0.91] | -10.8% | [-11.4%, -10.2%] | 0.8 pts | 1.1 | yes | yes |
| unique-str | uuid | 4096 | churn | ordered | baseline | 12 | 96.6 | 88.6 | 0.91× [0.90, 0.92] | -10.2% | [-11.3%, -9.1%] | 1.4 pts | 0.8 | yes | yes |
| unique-str | uuid | 4096 | build | ordered | baseline | 12 | 1.32 ms | 1.18 ms | 0.89× [0.89, 0.90] | -11.9% | [-12.7%, -11.1%] | 0.9 pts | 0.9 | yes | yes |
| unique-str | uuid | 16384 | valuesFor | ordered | baseline | 6 | 44.7 | 41.3 | 0.92× [0.91, 0.92] | -8.8% | [-9.4%, -8.2%] | 0.6 pts | 0.8 | yes | yes |
| unique-str | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 1942 | 1544 | 0.79× [0.79, 0.80] | -26.0% | [-27.3%, -24.7%] | 1.3 pts | 1.5 | yes | yes |
| unique-str | uuid | 16384 | prefix | ordered | baseline | 6 | 88.5 | 80.8 | 0.92× [0.91, 0.92] | -9.2% | [-9.8%, -8.7%] | 0.5 pts | 1.3 | yes | yes |
| unique-str | uuid | 16384 | churn | ordered | baseline | 6 | 113 | 108 | 0.96× [0.95, 0.97] | -4.4% | [-5.7%, -3.1%] | 1.2 pts | 0.8 | yes | yes |
| unique-str | uuid | 16384 | build | ordered | baseline | 6 | 6.24 ms | 5.63 ms | 0.91× [0.90, 0.91] | -10.1% | [-10.7%, -9.4%] | 0.6 pts | 1.0 | yes | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | baseline | 12 | 181 | 182 | 1.01× [1.00, 1.03] | +1.3% | [-0.2%, +2.8%] | 2.0 pts | 2.5 | yes | no |
| unique-str | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 5782 | 5624 | 0.97× [0.95, 0.99] | -3.6% | [-5.8%, -1.4%] | 2.2 pts | 2.1 | no | yes |
| unique-str | uuid | 262144 | prefix | ordered | baseline | 12 | 401 | 395 | 0.99× [0.98, 1.01] | -0.6% | [-1.8%, +0.6%] | 2.1 pts | 3.7 | yes | no |
| unique-str | uuid | 262144 | churn | ordered | baseline | 12 | 315 | 318 | 1.00× [0.99, 1.01] | -0.2% | [-1.2%, +0.9%] | 2.8 pts | 1.1 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str email n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.24% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str email n=262144 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-3.56%, 0.37%] includes zero
- multi-str email n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str email n=262144 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=4096 churn: ordered vs baseline: the pooled difference of -0.42% does not clear the 0.54% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=4096 churn: ordered vs baseline: the pooled interval [-1.40%, 0.56%] includes zero
- multi-str path n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.21% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str path n=16384 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.38% does not clear the 0.39% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.07%, 1.31%] includes zero
- multi-str path n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 valuesBetween: ordered vs baseline: 4 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str path n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of +1.70% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str path n=262144 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 churn: ordered vs baseline: the pooled difference of 1.59% does not clear the 1.63% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=262144 churn: ordered vs baseline: the pooled interval [-0.81%, 4.00%] includes zero
- multi-str str n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str str n=16384 churn: ordered vs baseline: the pooled interval [-0.71%, 2.13%] includes zero
- multi-str str n=16384 build: ordered vs baseline: the pooled interval [-0.50%, 2.24%] includes zero
- multi-str str n=262144 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.68% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str str n=262144 prefix: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.95% does not clear the 1.01% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.04%, 0.14%] includes zero
- multi-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.06%, 0.37%] includes zero
- multi-str u64 n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -1.04% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.18%, 0.76%] includes zero
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: 3 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.62% does not clear the 0.87% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.38%, 1.61%] includes zero
- multi-str url n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.58%, 2.56%] includes zero
- multi-str url n=16384 build: ordered vs baseline: the pooled difference of 0.24% does not clear the 0.48% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=16384 build: ordered vs baseline: the pooled interval [-0.43%, 0.91%] includes zero
- multi-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 churn: ordered vs baseline: the pooled difference of -0.38% does not clear the 2.02% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=262144 churn: ordered vs baseline: the pooled interval [-5.12%, 4.37%] includes zero
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.43%, 3.41%] includes zero
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str uuid n=262144 prefix: ordered vs baseline: the pooled interval [-3.99%, 0.12%] includes zero
- multi-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 prefix: ordered vs baseline: 1 processes resolved A as faster and 8 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str email n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str email n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=16384 prefix: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str email n=262144 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=16384 valuesFor: ordered vs baseline: the pooled interval [-14.04%, 1.03%] includes zero
- unique-str path n=16384 valuesFor: ordered vs baseline: the processes scatter 9.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=16384 valuesBetween: ordered vs baseline: the processes scatter 11.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.35%, 0.58%] includes zero
- unique-str str n=262144 prefix: ordered vs baseline: the pooled difference of -0.65% does not clear the 1.10% noise floor, the bound on what the harness reports between identical code in every process
- unique-str str n=262144 prefix: ordered vs baseline: the pooled interval [-1.78%, 0.49%] includes zero
- unique-str str n=262144 churn: ordered vs baseline: the pooled interval [-2.60%, 6.13%] includes zero
- unique-str street n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.31% does not clear the 0.59% noise floor, the bound on what the harness reports between identical code in every process
- unique-str street n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.49%, 2.12%] includes zero
- unique-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.17%, 0.08%] includes zero
- unique-str street n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.17% does not clear the 0.67% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.14%, 1.48%] includes zero
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str u64 n=262144 churn: ordered vs baseline: the pooled difference of -2.14% does not clear the 3.67% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=262144 churn: ordered vs baseline: the pooled interval [-6.71%, 2.43%] includes zero
- unique-str url n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.54% does not clear the 0.58% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.35%, 1.27%] includes zero
- unique-str url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.30% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.58%, 0.98%] includes zero
- unique-str url n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=16384 churn: ordered vs baseline: the pooled difference of -0.25% does not clear the 0.66% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=16384 churn: ordered vs baseline: the pooled interval [-1.45%, 0.96%] includes zero
- unique-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.68% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.21%, 2.77%] includes zero
- unique-str uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 prefix: ordered vs baseline: the pooled interval [-1.79%, 0.58%] includes zero
- unique-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 prefix: ordered vs baseline: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str uuid n=262144 churn: ordered vs baseline: the pooled difference of -0.16% does not clear the 1.69% noise floor, the bound on what the harness reports between identical code in every process
- unique-str uuid n=262144 churn: ordered vs baseline: the pooled interval [-1.24%, 0.93%] includes zero
