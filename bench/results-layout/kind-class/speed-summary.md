| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 8 | 49.7 | 49.5 | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.4%] | 1.2 pts | 0.9 | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 8 | 2988 | 2951 | 0.99× [0.98, 1.00] | -1.2% | [-2.1%, -0.2%] | 1.1 pts | 0.5 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 8 | 76.3 | 77.0 | 1.01× [1.00, 1.01] | +0.7% | [-0.1%, +1.4%] | 0.9 pts | 1.0 | yes |
| multi | email | 4096 | churn | ordered | baseline | 8 | 63.5 | 64.4 | 1.01× [1.00, 1.01] | +0.9% | [+0.4%, +1.3%] | 0.5 pts | 0.4 | yes |
| multi | email | 4096 | build | ordered | baseline | 8 | 5.49 ms | 5.45 ms | 0.99× [0.97, 1.01] | -1.0% | [-2.8%, +0.7%] | 2.1 pts | 1.7 | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 61.1 | 61.2 | 1.00× [1.00, 1.01] | +0.2% | [-0.4%, +0.8%] | 0.6 pts | 0.5 | yes |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3495 | 3508 | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +0.9%] | 0.8 pts | 0.7 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 92.3 | 92.8 | 1.01× [1.00, 1.02] | +1.0% | [+0.1%, +1.9%] | 0.9 pts | 0.9 | yes |
| multi | email | 16384 | churn | ordered | baseline | 6 | 93.6 | 94.1 | 1.01× [1.00, 1.01] | +0.5% | [-0.3%, +1.4%] | 0.8 pts | 0.5 | yes |
| multi | email | 16384 | build | ordered | baseline | 6 | 28.02 ms | 27.92 ms | 1.00× [0.99, 1.01] | -0.1% | [-1.0%, +0.9%] | 0.9 pts | 0.8 | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 10 | 255 | 252 | 1.00× [0.99, 1.01] | -0.3% | [-1.4%, +0.8%] | 1.5 pts | 1.3 | yes |
| multi | email | 262144 | valuesBetween | ordered | baseline | 10 | 7944 | 7984 | 1.00× [0.99, 1.01] | +0.2% | [-0.7%, +1.1%] | 1.2 pts | 1.1 | yes |
| multi | email | 262144 | prefix | ordered | baseline | 10 | 309 | 310 | 1.01× [1.00, 1.02] | +0.7% | [-0.3%, +1.8%] | 1.4 pts | 0.9 | yes |
| multi | email | 262144 | churn | ordered | baseline | 10 | 389 | 367 | 0.94× [0.91, 0.97] | -6.2% | [-9.4%, -3.1%] | 4.4 pts | 1.6 | no |
| multi | email | 1048576 | valuesFor | ordered | baseline | 10 | 377 | 374 | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.3%] | 1.0 pts | 0.9 | yes |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 10 | 9719 | 9710 | 1.00× [0.99, 1.00] | -0.3% | [-0.7%, +0.2%] | 0.6 pts | 0.6 | yes |
| multi | email | 1048576 | prefix | ordered | baseline | 10 | 580 | 583 | 1.02× [1.00, 1.03] | +1.6% | [+0.5%, +2.8%] | 1.6 pts | 1.6 | yes |
| multi | email | 1048576 | churn | ordered | baseline | 10 | 613 | 617 | 0.99× [0.96, 1.03] | -0.8% | [-4.6%, +2.9%] | 5.2 pts | 0.9 | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 107 | 107 | 0.99× [0.98, 1.01] | -0.7% | [-2.2%, +0.8%] | 1.4 pts | 1.5 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4189 | 4280 | 1.02× [1.01, 1.03] | +2.2% | [+1.1%, +3.3%] | 1.0 pts | 0.6 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 334 | 325 | 0.98× [0.97, 1.00] | -1.7% | [-3.5%, +0.0%] | 1.7 pts | 0.5 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 155 | 155 | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.6%] | 0.8 pts | 0.6 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 11.51 ms | 11.52 ms | 1.00× [0.99, 1.01] | -0.2% | [-1.1%, +0.8%] | 0.9 pts | 1.2 | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 10 | 136 | 137 | 1.00× [1.00, 1.01] | +0.2% | [-0.5%, +1.0%] | 1.0 pts | 0.8 | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 10 | 4767 | 4754 | 1.00× [1.00, 1.01] | +0.3% | [-0.5%, +1.1%] | 1.1 pts | 1.2 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 10 | 616 | 615 | 1.00× [0.98, 1.03] | +0.2% | [-2.1%, +2.4%] | 3.2 pts | 0.5 | no |
| multi | path | 16384 | churn | ordered | baseline | 10 | 243 | 244 | 1.00× [0.99, 1.00] | -0.4% | [-1.2%, +0.4%] | 1.1 pts | 0.6 | yes |
| multi | path | 16384 | build | ordered | baseline | 10 | 65.80 ms | 65.21 ms | 0.99× [0.99, 1.00] | -0.8% | [-1.3%, -0.3%] | 0.7 pts | 0.6 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 478 | 461 | 0.98× [0.96, 1.00] | -1.6% | [-3.6%, +0.4%] | 2.8 pts | 1.9 | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 12.4 µs | 12.4 µs | 0.99× [0.98, 1.01] | -0.9% | [-2.3%, +0.5%] | 2.0 pts | 1.3 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 12.3 µs | 12.4 µs | 1.01× [0.99, 1.03] | +0.7% | [-1.1%, +2.5%] | 2.5 pts | 0.2 | yes |
| multi | path | 262144 | churn | ordered | baseline | 10 | 816 | 829 | 1.00× [0.98, 1.03] | +0.0% | [-2.5%, +2.6%] | 3.6 pts | 1.5 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 8 | 60.9 | 60.9 | 1.00× [0.98, 1.01] | -0.4% | [-1.6%, +0.8%] | 1.5 pts | 1.1 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 8 | 3519 | 3532 | 1.01× [1.00, 1.01] | +0.6% | [+0.0%, +1.2%] | 0.7 pts | 0.4 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 8 | 7915 | 8030 | 1.02× [1.00, 1.04] | +1.8% | [+0.1%, +3.5%] | 2.0 pts | 0.7 | yes |
| multi | str | 4096 | churn | ordered | baseline | 8 | 78.5 | 78.3 | 1.00× [0.99, 1.00] | -0.2% | [-0.8%, +0.4%] | 0.7 pts | 0.5 | yes |
| multi | str | 4096 | build | ordered | baseline | 8 | 6.53 ms | 6.56 ms | 1.00× [1.00, 1.01] | +0.2% | [-0.0%, +0.5%] | 0.3 pts | 0.3 | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 71.5 | 71.3 | 1.00× [0.99, 1.00] | -0.4% | [-0.8%, +0.1%] | 0.4 pts | 0.4 | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3690 | 3674 | 1.00× [0.99, 1.01] | -0.4% | [-1.3%, +0.6%] | 0.9 pts | 0.8 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 34.9 µs | 34.7 µs | 0.99× [0.98, 1.00] | -0.7% | [-1.8%, +0.4%] | 1.0 pts | 0.7 | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 109 | 108 | 1.00× [0.98, 1.02] | +0.3% | [-1.6%, +2.1%] | 1.7 pts | 1.2 | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 32.54 ms | 32.68 ms | 1.00× [1.00, 1.01] | +0.1% | [-0.5%, +0.6%] | 0.5 pts | 0.5 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 269 | 270 | 1.01× [1.00, 1.01] | +0.6% | [-0.2%, +1.5%] | 1.2 pts | 1.1 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 8766 | 8771 | 1.00× [1.00, 1.00] | -0.0% | [-0.5%, +0.5%] | 0.7 pts | 0.5 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.40 ms | 1.39 ms | 1.00× [0.99, 1.00] | -0.1% | [-0.6%, +0.3%] | 0.6 pts | 0.3 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 403 | 418 | 1.03× [1.00, 1.06] | +2.7% | [-0.0%, +5.5%] | 3.9 pts | 1.1 | no |
| multi | str | 1048576 | valuesFor | ordered | baseline | 10 | 400 | 401 | 1.00× [0.99, 1.01] | -0.1% | [-0.8%, +0.6%] | 1.0 pts | 0.8 | yes |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 10 | 9866 | 9850 | 1.00× [0.99, 1.01] | -0.2% | [-1.2%, +0.7%] | 1.3 pts | 1.2 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 10 | 6.10 ms | 6.09 ms | 1.00× [1.00, 1.00] | -0.2% | [-0.4%, +0.1%] | 0.4 pts | 0.7 | yes |
| multi | str | 1048576 | churn | ordered | baseline | 10 | 610 | 605 | 1.00× [0.94, 1.07] | -0.2% | [-6.6%, +6.2%] | 9.0 pts | 2.9 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 6 | 53.2 | 53.6 | 1.00× [0.98, 1.02] | +0.0% | [-1.7%, +1.8%] | 1.7 pts | 1.1 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 6 | 2426 | 2385 | 0.99× [0.98, 1.00] | -0.8% | [-1.8%, +0.1%] | 0.9 pts | 0.4 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 6 | 260 | 262 | 1.00× [0.99, 1.01] | +0.3% | [-0.6%, +1.1%] | 0.8 pts | 0.3 | yes |
| multi | street | 4096 | churn | ordered | baseline | 6 | 92.8 | 92.8 | 1.00× [0.99, 1.00] | -0.1% | [-0.5%, +0.3%] | 0.4 pts | 0.3 | yes |
| multi | street | 4096 | build | ordered | baseline | 6 | 2.90 ms | 2.88 ms | 0.99× [0.98, 1.00] | -0.7% | [-1.7%, +0.3%] | 0.9 pts | 0.7 | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 10 | 76.1 | 76.5 | 1.00× [1.00, 1.01] | +0.2% | [-0.5%, +0.9%] | 1.0 pts | 0.7 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 10 | 2878 | 2870 | 1.00× [0.99, 1.00] | -0.4% | [-0.8%, -0.0%] | 0.5 pts | 0.5 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 10 | 943 | 945 | 1.00× [0.98, 1.02] | -0.0% | [-1.8%, +1.8%] | 2.6 pts | 0.9 | yes |
| multi | street | 16384 | churn | ordered | baseline | 10 | 128 | 128 | 1.00× [0.99, 1.01] | +0.0% | [-0.6%, +0.7%] | 0.9 pts | 0.7 | yes |
| multi | street | 16384 | build | ordered | baseline | 10 | 16.09 ms | 15.98 ms | 0.99× [0.99, 1.00] | -0.6% | [-1.0%, -0.2%] | 0.5 pts | 0.7 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 10 | 36.8 | 36.5 | 0.99× [0.98, 1.00] | -0.9% | [-1.7%, -0.1%] | 1.1 pts | 0.7 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 10 | 2687 | 2655 | 0.99× [0.98, 1.01] | -0.7% | [-1.9%, +0.5%] | 1.7 pts | 0.8 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 10 | 45.2 | 44.9 | 1.00× [0.99, 1.01] | -0.0% | [-0.8%, +0.7%] | 1.1 pts | 0.7 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 10 | 4.09 ms | 4.13 ms | 1.00× [0.98, 1.02] | +0.2% | [-1.7%, +2.1%] | 2.6 pts | 1.4 | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 45.8 | 46.1 | 1.00× [0.99, 1.01] | +0.3% | [-0.6%, +1.3%] | 0.9 pts | 1.1 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3497 | 3476 | 0.99× [0.99, 1.00] | -0.6% | [-1.2%, -0.1%] | 0.5 pts | 0.5 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 56.9 | 56.6 | 1.00× [0.99, 1.01] | -0.1% | [-1.2%, +1.0%] | 1.0 pts | 1.0 | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 20.43 ms | 20.23 ms | 0.99× [0.98, 1.00] | -0.7% | [-1.8%, +0.4%] | 1.1 pts | 0.9 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 139 | 142 | 1.02× [1.00, 1.03] | +1.8% | [+0.2%, +3.4%] | 2.2 pts | 2.0 | yes |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 6913 | 6951 | 1.00× [0.99, 1.02] | +0.4% | [-0.9%, +1.7%] | 1.8 pts | 1.4 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 263 | 261 | 1.01× [0.97, 1.04] | +0.7% | [-2.8%, +4.1%] | 4.8 pts | 1.1 | no |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 203 | 211 | 1.04× [1.03, 1.04] | +3.7% | [+3.2%, +4.2%] | 0.6 pts | 0.6 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 6671 | 6603 | 0.99× [0.99, 1.00] | -0.6% | [-1.1%, +0.0%] | 0.8 pts | 0.7 | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 10 | 373 | 374 | 1.00× [0.96, 1.04] | -0.4% | [-4.3%, +3.4%] | 5.4 pts | 0.6 | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 6 | 89.3 | 89.6 | 1.00× [1.00, 1.01] | +0.3% | [-0.1%, +0.6%] | 0.3 pts | 0.3 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 6 | 3932 | 4029 | 1.02× [1.01, 1.03] | +2.0% | [+1.4%, +2.7%] | 0.6 pts | 0.3 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 6 | 186 | 186 | 0.99× [0.98, 1.01] | -0.6% | [-1.9%, +0.7%] | 1.2 pts | 0.7 | yes |
| multi | url | 4096 | churn | ordered | baseline | 6 | 129 | 127 | 0.99× [0.98, 1.01] | -0.6% | [-1.9%, +0.7%] | 1.2 pts | 0.9 | yes |
| multi | url | 4096 | build | ordered | baseline | 6 | 9.47 ms | 9.44 ms | 1.00× [0.99, 1.00] | -0.3% | [-0.7%, +0.2%] | 0.5 pts | 0.7 | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 114 | 114 | 1.00× [0.98, 1.01] | -0.1% | [-1.6%, +1.4%] | 1.4 pts | 1.0 | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4559 | 4570 | 1.00× [0.99, 1.02] | +0.4% | [-1.0%, +1.8%] | 1.3 pts | 0.9 | yes |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 276 | 275 | 1.00× [0.99, 1.01] | +0.4% | [-0.5%, +1.4%] | 0.9 pts | 0.5 | yes |
| multi | url | 16384 | churn | ordered | baseline | 6 | 209 | 207 | 1.00× [0.98, 1.01] | -0.4% | [-1.7%, +0.9%] | 1.2 pts | 0.5 | yes |
| multi | url | 16384 | build | ordered | baseline | 6 | 53.68 ms | 53.81 ms | 1.00× [0.99, 1.01] | -0.0% | [-0.7%, +0.7%] | 0.7 pts | 0.7 | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 10 | 446 | 443 | 0.99× [0.95, 1.04] | -0.5% | [-5.0%, +3.9%] | 6.2 pts | 4.4 | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 10 | 11.9 µs | 11.9 µs | 1.00× [0.99, 1.01] | -0.3% | [-1.3%, +0.6%] | 1.3 pts | 1.0 | yes |
| multi | url | 262144 | prefix | ordered | baseline | 10 | 2263 | 2271 | 0.99× [0.97, 1.02] | -1.0% | [-3.4%, +1.5%] | 3.4 pts | 0.6 | no |
| multi | url | 262144 | churn | ordered | baseline | 10 | 714 | 707 | 0.98× [0.96, 1.00] | -2.2% | [-4.4%, -0.1%] | 3.0 pts | 1.0 | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 10 | 640 | 649 | 1.00× [0.98, 1.02] | -0.2% | [-2.1%, +1.7%] | 2.7 pts | 2.1 | yes |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 10 | 13.2 µs | 13.2 µs | 1.00× [0.99, 1.01] | +0.3% | [-0.6%, +1.1%] | 1.2 pts | 1.1 | yes |
| multi | url | 1048576 | prefix | ordered | baseline | 10 | 9342 | 9354 | 1.00× [0.98, 1.02] | -0.1% | [-1.8%, +1.6%] | 2.4 pts | 0.2 | yes |
| multi | url | 1048576 | churn | ordered | baseline | 10 | 990 | 1018 | 1.03× [0.99, 1.07] | +2.4% | [-1.4%, +6.3%] | 5.4 pts | 1.4 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 53.1 | 53.0 | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.7%] | 0.8 pts | 0.6 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3307 | 3324 | 1.00× [0.99, 1.01] | -0.0% | [-1.2%, +1.1%] | 1.1 pts | 0.5 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 89.8 | 90.3 | 1.01× [1.00, 1.01] | +0.7% | [+0.0%, +1.4%] | 0.6 pts | 0.7 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 74.3 | 74.2 | 1.00× [0.99, 1.00] | -0.3% | [-0.9%, +0.3%] | 0.6 pts | 0.4 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 6.06 ms | 6.06 ms | 1.00× [0.99, 1.00] | -0.3% | [-0.9%, +0.4%] | 0.6 pts | 0.7 | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 62.8 | 62.8 | 1.00× [0.99, 1.01] | -0.0% | [-0.9%, +0.9%] | 0.9 pts | 0.8 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3629 | 3610 | 1.00× [0.99, 1.00] | -0.3% | [-0.8%, +0.2%] | 0.5 pts | 0.5 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 110 | 110 | 1.01× [1.00, 1.02] | +1.0% | [-0.4%, +2.3%] | 1.3 pts | 1.2 | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 103 | 104 | 1.01× [0.99, 1.02] | +0.5% | [-0.8%, +1.8%] | 1.3 pts | 0.8 | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 30.22 ms | 30.31 ms | 1.00× [0.98, 1.01] | -0.3% | [-1.8%, +1.2%] | 1.5 pts | 1.4 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 263 | 267 | 1.00× [0.98, 1.01] | -0.1% | [-1.7%, +1.4%] | 2.1 pts | 1.7 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 8956 | 8988 | 1.00× [0.99, 1.01] | -0.4% | [-1.4%, +0.5%] | 1.3 pts | 1.0 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 598 | 597 | 0.99× [0.98, 1.01] | -0.6% | [-1.8%, +0.6%] | 1.7 pts | 1.3 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 430 | 429 | 1.02× [0.97, 1.07] | +1.9% | [-3.0%, +6.9%] | 6.9 pts | 2.6 | no |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 10 | 390 | 393 | 1.01× [1.00, 1.02] | +0.7% | [-0.3%, +1.6%] | 1.4 pts | 1.3 | yes |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 10 | 10.3 µs | 10.3 µs | 1.00× [0.99, 1.00] | -0.0% | [-0.5%, +0.4%] | 0.7 pts | 0.6 | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 10 | 1876 | 1880 | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.7%] | 1.2 pts | 0.9 | yes |
| multi | uuid | 1048576 | churn | ordered | baseline | 10 | 621 | 607 | 1.00× [0.95, 1.07] | +0.4% | [-5.7%, +6.5%] | 8.6 pts | 1.4 | no |
| unique | email | 4096 | valuesFor | ordered | baseline | 10 | 25.8 | 25.8 | 1.00× [0.98, 1.01] | -0.5% | [-1.8%, +0.9%] | 1.8 pts | 0.9 | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 10 | 1298 | 1269 | 0.98× [0.97, 0.98] | -2.4% | [-3.2%, -1.6%] | 1.1 pts | 0.8 | yes |
| unique | email | 4096 | prefix | ordered | baseline | 10 | 56.6 | 57.2 | 1.01× [1.00, 1.02] | +1.3% | [+0.2%, +2.4%] | 1.6 pts | 1.9 | yes |
| unique | email | 4096 | churn | ordered | baseline | 10 | 66.9 | 66.9 | 1.00× [0.99, 1.00] | -0.3% | [-0.9%, +0.3%] | 0.8 pts | 0.9 | yes |
| unique | email | 4096 | build | ordered | baseline | 10 | 821.2 µs | 816.1 µs | 0.99× [0.97, 1.01] | -1.0% | [-2.9%, +1.0%] | 2.7 pts | 1.1 | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 38.1 | 37.9 | 0.99× [0.98, 1.01] | -0.5% | [-1.6%, +0.5%] | 1.0 pts | 1.1 | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 1639 | 1620 | 0.99× [0.98, 1.00] | -1.1% | [-1.8%, -0.3%] | 0.7 pts | 0.7 | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 70.3 | 69.8 | 0.99× [0.99, 1.00] | -0.7% | [-1.1%, -0.4%] | 0.3 pts | 0.3 | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 81.0 | 81.6 | 1.01× [1.00, 1.02] | +0.7% | [-0.3%, +1.7%] | 0.9 pts | 0.6 | yes |
| unique | email | 16384 | build | ordered | baseline | 6 | 4.02 ms | 4.03 ms | 1.00× [0.99, 1.01] | +0.0% | [-0.9%, +1.0%] | 0.9 pts | 0.6 | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 10 | 200 | 202 | 1.01× [0.98, 1.04] | +0.7% | [-2.1%, +3.5%] | 3.9 pts | 2.4 | no |
| unique | email | 262144 | valuesBetween | ordered | baseline | 10 | 3914 | 3910 | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +1.0%] | 1.2 pts | 0.8 | yes |
| unique | email | 262144 | prefix | ordered | baseline | 10 | 225 | 221 | 0.99× [0.98, 1.00] | -0.9% | [-1.8%, -0.0%] | 1.2 pts | 0.8 | yes |
| unique | email | 262144 | churn | ordered | baseline | 10 | 236 | 232 | 0.99× [0.98, 1.00] | -0.7% | [-1.8%, +0.3%] | 1.5 pts | 1.2 | yes |
| unique | email | 1048576 | valuesFor | ordered | baseline | 8 | 317 | 313 | 1.00× [0.99, 1.01] | -0.5% | [-1.5%, +0.5%] | 1.2 pts | 1.3 | yes |
| unique | email | 1048576 | valuesBetween | ordered | baseline | 8 | 6477 | 6412 | 0.99× [0.98, 1.00] | -0.8% | [-1.7%, +0.1%] | 1.1 pts | 1.2 | yes |
| unique | email | 1048576 | prefix | ordered | baseline | 8 | 428 | 429 | 1.00× [0.99, 1.01] | +0.3% | [-0.8%, +1.5%] | 1.4 pts | 1.3 | yes |
| unique | email | 1048576 | churn | ordered | baseline | 8 | 435 | 435 | 1.00× [0.98, 1.02] | -0.1% | [-1.8%, +1.7%] | 2.1 pts | 0.9 | yes |
| unique | path | 4096 | valuesFor | ordered | baseline | 6 | 83.7 | 84.4 | 1.00× [0.99, 1.01] | +0.3% | [-0.7%, +1.3%] | 1.0 pts | 0.7 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 6 | 2220 | 2200 | 1.00× [0.99, 1.02] | +0.1% | [-1.3%, +1.6%] | 1.4 pts | 0.8 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 6 | 255 | 253 | 0.99× [0.99, 1.00] | -0.5% | [-1.5%, +0.4%] | 0.9 pts | 0.5 | yes |
| unique | path | 4096 | churn | ordered | baseline | 6 | 185 | 187 | 1.01× [1.00, 1.01] | +0.7% | [+0.2%, +1.3%] | 0.5 pts | 0.4 | yes |
| unique | path | 4096 | build | ordered | baseline | 6 | 2.07 ms | 2.05 ms | 0.99× [0.99, 1.00] | -0.7% | [-1.1%, -0.3%] | 0.4 pts | 0.3 | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 10 | 115 | 114 | 1.00× [0.99, 1.01] | -0.1% | [-1.0%, +0.7%] | 1.2 pts | 1.1 | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 10 | 2697 | 2686 | 1.00× [0.99, 1.00] | -0.1% | [-0.7%, +0.4%] | 0.7 pts | 0.7 | yes |
| unique | path | 16384 | prefix | ordered | baseline | 10 | 428 | 432 | 1.00× [0.98, 1.02] | -0.1% | [-1.8%, +1.6%] | 2.4 pts | 0.5 | yes |
| unique | path | 16384 | churn | ordered | baseline | 10 | 245 | 250 | 1.00× [0.99, 1.02] | +0.3% | [-0.9%, +1.6%] | 1.8 pts | 1.0 | yes |
| unique | path | 16384 | build | ordered | baseline | 10 | 10.38 ms | 10.43 ms | 1.00× [1.00, 1.01] | +0.2% | [-0.2%, +0.7%] | 0.6 pts | 0.5 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 420 | 422 | 1.01× [1.00, 1.03] | +1.3% | [+0.1%, +2.5%] | 1.6 pts | 1.1 | yes |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 8386 | 8529 | 1.01× [0.99, 1.02] | +0.8% | [-0.5%, +2.2%] | 1.9 pts | 1.5 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 5748 | 5876 | 1.01× [0.98, 1.05] | +1.4% | [-2.0%, +4.8%] | 4.7 pts | 0.4 | no |
| unique | path | 262144 | churn | ordered | baseline | 10 | 636 | 638 | 1.00× [0.98, 1.02] | -0.1% | [-1.8%, +1.6%] | 2.4 pts | 2.0 | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 10 | 36.3 | 36.3 | 1.00× [0.99, 1.01] | -0.1% | [-0.8%, +0.6%] | 1.0 pts | 0.7 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 10 | 1723 | 1702 | 0.99× [0.98, 1.00] | -0.8% | [-2.2%, +0.5%] | 1.9 pts | 1.2 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 10 | 3346 | 3356 | 1.00× [0.98, 1.02] | +0.3% | [-1.7%, +2.3%] | 2.9 pts | 1.0 | no |
| unique | str | 4096 | churn | ordered | baseline | 10 | 86.6 | 87.0 | 1.00× [0.99, 1.01] | +0.1% | [-0.5%, +0.7%] | 0.9 pts | 0.8 | yes |
| unique | str | 4096 | build | ordered | baseline | 10 | 1.02 ms | 1.02 ms | 1.00× [0.99, 1.01] | -0.2% | [-1.2%, +0.9%] | 1.5 pts | 0.5 | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 6 | 48.1 | 47.9 | 1.00× [0.99, 1.01] | -0.1% | [-0.9%, +0.6%] | 0.7 pts | 0.8 | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 1696 | 1693 | 1.00× [0.99, 1.00] | -0.3% | [-0.8%, +0.1%] | 0.5 pts | 0.4 | yes |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 14.6 µs | 14.5 µs | 1.00× [0.99, 1.00] | -0.3% | [-0.8%, +0.1%] | 0.5 pts | 0.3 | yes |
| unique | str | 16384 | churn | ordered | baseline | 6 | 103 | 104 | 1.01× [1.00, 1.02] | +1.0% | [-0.1%, +2.1%] | 1.0 pts | 0.7 | yes |
| unique | str | 16384 | build | ordered | baseline | 6 | 4.92 ms | 4.90 ms | 0.99× [0.99, 1.00] | -0.5% | [-1.3%, +0.3%] | 0.8 pts | 0.7 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 206 | 209 | 1.00× [0.99, 1.01] | +0.2% | [-1.0%, +1.4%] | 1.7 pts | 1.0 | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 4467 | 4448 | 0.99× [0.97, 1.01] | -1.0% | [-2.6%, +0.6%] | 2.2 pts | 1.7 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 713.0 µs | 676.6 µs | 0.98× [0.96, 1.01] | -2.1% | [-4.7%, +0.5%] | 3.6 pts | 0.6 | no |
| unique | str | 262144 | churn | ordered | baseline | 10 | 310 | 313 | 1.00× [0.99, 1.01] | +0.3% | [-0.8%, +1.4%] | 1.5 pts | 1.0 | yes |
| unique | str | 1048576 | valuesFor | ordered | baseline | 6 | 347 | 348 | 0.99× [0.98, 1.01] | -0.5% | [-2.4%, +1.4%] | 1.8 pts | 2.1 | yes |
| unique | str | 1048576 | valuesBetween | ordered | baseline | 6 | 6679 | 6681 | 1.00× [0.98, 1.02] | -0.2% | [-1.9%, +1.6%] | 1.7 pts | 1.7 | yes |
| unique | str | 1048576 | prefix | ordered | baseline | 6 | 4.09 ms | 4.10 ms | 1.00× [0.98, 1.01] | -0.4% | [-2.1%, +1.2%] | 1.6 pts | 1.9 | yes |
| unique | str | 1048576 | churn | ordered | baseline | 6 | 474 | 467 | 0.98× [0.98, 0.99] | -1.6% | [-2.3%, -0.9%] | 0.7 pts | 0.4 | yes |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 47.4 | 47.4 | 1.00× [0.99, 1.01] | +0.3% | [-0.8%, +1.4%] | 1.0 pts | 0.8 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 1913 | 1921 | 1.00× [0.98, 1.02] | -0.1% | [-1.7%, +1.5%] | 1.5 pts | 0.9 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 218 | 219 | 1.01× [0.99, 1.02] | +0.8% | [-0.7%, +2.3%] | 1.4 pts | 0.5 | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 110 | 110 | 1.00× [0.99, 1.01] | +0.1% | [-1.2%, +1.4%] | 1.3 pts | 0.8 | yes |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.28 ms | 1.27 ms | 1.00× [0.98, 1.01] | -0.5% | [-2.2%, +1.2%] | 1.6 pts | 1.0 | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 6 | 70.2 | 70.1 | 1.00× [0.99, 1.01] | -0.3% | [-1.2%, +0.5%] | 0.8 pts | 0.7 | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 6 | 2253 | 2237 | 0.99× [0.98, 1.00] | -0.8% | [-1.9%, +0.3%] | 1.1 pts | 0.8 | yes |
| unique | street | 16384 | prefix | ordered | baseline | 6 | 763 | 760 | 1.00× [0.98, 1.01] | -0.4% | [-1.9%, +1.2%] | 1.4 pts | 0.5 | yes |
| unique | street | 16384 | churn | ordered | baseline | 6 | 142 | 143 | 1.01× [1.00, 1.02] | +1.0% | [-0.0%, +2.1%] | 1.0 pts | 0.5 | yes |
| unique | street | 16384 | build | ordered | baseline | 6 | 6.40 ms | 6.41 ms | 1.00× [1.00, 1.01] | +0.2% | [-0.5%, +0.8%] | 0.6 pts | 0.8 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 14.9 | 14.8 | 1.00× [0.99, 1.02] | +0.4% | [-1.2%, +2.0%] | 1.5 pts | 0.7 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1038 | 995 | 0.96× [0.95, 0.97] | -4.2% | [-4.9%, -3.6%] | 0.6 pts | 0.5 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 43.2 | 43.3 | 1.00× [0.99, 1.01] | -0.3% | [-1.5%, +1.0%] | 1.2 pts | 1.2 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 598.9 µs | 598.1 µs | 0.99× [0.98, 1.01] | -0.5% | [-1.9%, +0.9%] | 1.3 pts | 0.8 | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 6 | 23.4 | 23.7 | 1.01× [1.00, 1.03] | +1.3% | [+0.0%, +2.5%] | 1.2 pts | 1.7 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 1612 | 1593 | 0.99× [0.98, 0.99] | -1.4% | [-1.8%, -0.9%] | 0.4 pts | 0.3 | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 6 | 44.2 | 44.4 | 1.01× [0.99, 1.02] | +0.6% | [-0.6%, +1.9%] | 1.2 pts | 0.6 | yes |
| unique | u64 | 16384 | build | ordered | baseline | 6 | 2.85 ms | 2.84 ms | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.6%] | 0.8 pts | 0.6 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 10 | 70.3 | 67.3 | 0.98× [0.96, 1.00] | -2.0% | [-4.1%, +0.1%] | 3.0 pts | 1.1 | no |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 2425 | 2406 | 1.00× [0.99, 1.00] | -0.4% | [-0.9%, +0.1%] | 0.7 pts | 0.2 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 10 | 154 | 152 | 1.01× [0.99, 1.03] | +0.6% | [-1.5%, +2.6%] | 2.9 pts | 0.3 | no |
| unique | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 139 | 138 | 0.99× [0.99, 1.00] | -0.6% | [-1.3%, +0.1%] | 1.0 pts | 1.1 | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 3205 | 3170 | 0.99× [0.98, 1.00] | -1.2% | [-2.0%, -0.5%] | 1.0 pts | 1.2 | yes |
| unique | u64 | 1048576 | churn | ordered | baseline | 10 | 306 | 298 | 0.99× [0.97, 1.01] | -1.3% | [-3.3%, +0.8%] | 2.9 pts | 1.2 | no |
| unique | url | 4096 | valuesFor | ordered | baseline | 6 | 66.2 | 65.6 | 0.99× [0.99, 1.00] | -0.7% | [-1.5%, +0.1%] | 0.8 pts | 0.5 | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 6 | 2045 | 2060 | 1.01× [1.00, 1.01] | +0.7% | [+0.1%, +1.3%] | 0.6 pts | 0.4 | yes |
| unique | url | 4096 | prefix | ordered | baseline | 6 | 156 | 155 | 1.00× [0.99, 1.01] | -0.4% | [-1.5%, +0.7%] | 1.1 pts | 0.8 | yes |
| unique | url | 4096 | churn | ordered | baseline | 6 | 147 | 147 | 1.00× [0.99, 1.01] | +0.3% | [-0.8%, +1.3%] | 1.0 pts | 0.6 | yes |
| unique | url | 4096 | build | ordered | baseline | 6 | 1.66 ms | 1.66 ms | 1.00× [0.99, 1.01] | -0.0% | [-0.9%, +0.9%] | 0.9 pts | 0.5 | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 91.6 | 91.2 | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.6%] | 0.8 pts | 0.8 | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 2487 | 2486 | 1.00× [0.99, 1.01] | -0.1% | [-1.4%, +1.2%] | 1.2 pts | 1.1 | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 218 | 217 | 1.00× [0.99, 1.01] | -0.1% | [-1.2%, +1.0%] | 1.1 pts | 0.6 | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 194 | 192 | 0.99× [0.98, 0.99] | -1.4% | [-2.3%, -0.5%] | 0.8 pts | 0.3 | yes |
| unique | url | 16384 | build | ordered | baseline | 6 | 8.28 ms | 8.32 ms | 1.00× [1.00, 1.01] | +0.3% | [-0.5%, +1.0%] | 0.7 pts | 0.7 | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 10 | 384 | 383 | 0.99× [0.94, 1.04] | -1.4% | [-6.5%, +3.6%] | 7.0 pts | 4.9 | no |
| unique | url | 262144 | valuesBetween | ordered | baseline | 10 | 8017 | 7972 | 1.00× [0.98, 1.02] | +0.1% | [-1.6%, +1.8%] | 2.4 pts | 2.1 | yes |
| unique | url | 262144 | prefix | ordered | baseline | 10 | 1448 | 1436 | 0.99× [0.97, 1.01] | -0.8% | [-3.0%, +1.4%] | 3.1 pts | 0.7 | no |
| unique | url | 262144 | churn | ordered | baseline | 10 | 563 | 562 | 0.99× [0.98, 1.00] | -0.7% | [-1.6%, +0.2%] | 1.2 pts | 1.0 | yes |
| unique | url | 1048576 | valuesFor | ordered | baseline | 10 | 597 | 596 | 1.00× [0.97, 1.04] | +0.1% | [-3.4%, +3.6%] | 4.9 pts | 3.8 | no |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 10 | 9575 | 9646 | 1.00× [0.99, 1.01] | -0.0% | [-1.2%, +1.2%] | 1.7 pts | 1.7 | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 10 | 5834 | 5959 | 1.02× [0.99, 1.05] | +1.7% | [-1.0%, +4.5%] | 3.9 pts | 0.4 | no |
| unique | url | 1048576 | churn | ordered | baseline | 10 | 933 | 944 | 1.01× [0.99, 1.02] | +0.7% | [-0.7%, +2.0%] | 1.9 pts | 1.2 | yes |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 6 | 28.8 | 28.8 | 1.00× [0.99, 1.02] | +0.3% | [-1.0%, +1.7%] | 1.3 pts | 0.8 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1532 | 1494 | 0.98× [0.97, 0.99] | -2.1% | [-3.4%, -0.9%] | 1.2 pts | 0.8 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 68.2 | 68.1 | 1.00× [0.99, 1.00] | -0.1% | [-0.6%, +0.3%] | 0.5 pts | 0.5 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 6 | 76.2 | 76.2 | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +0.8%] | 0.7 pts | 0.7 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 6 | 895.6 µs | 899.9 µs | 1.00× [0.99, 1.01] | +0.1% | [-0.5%, +0.7%] | 0.6 pts | 0.3 | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 8 | 39.3 | 39.0 | 0.99× [0.99, 1.00] | -0.7% | [-1.5%, +0.1%] | 0.9 pts | 1.1 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 8 | 1639 | 1616 | 0.99× [0.98, 0.99] | -1.1% | [-1.7%, -0.5%] | 0.7 pts | 0.6 | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 8 | 83.3 | 82.9 | 1.00× [0.99, 1.01] | -0.4% | [-1.3%, +0.5%] | 1.1 pts | 1.1 | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 8 | 88.9 | 89.5 | 1.01× [1.00, 1.03] | +1.2% | [-0.5%, +2.9%] | 2.0 pts | 1.3 | yes |
| unique | uuid | 16384 | build | ordered | baseline | 8 | 4.34 ms | 4.30 ms | 0.99× [0.99, 1.00] | -0.7% | [-1.4%, +0.1%] | 0.9 pts | 0.9 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 198 | 197 | 1.00× [0.99, 1.01] | -0.0% | [-1.4%, +1.3%] | 1.9 pts | 1.4 | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 4612 | 4635 | 1.00× [0.99, 1.01] | +0.1% | [-0.9%, +1.2%] | 1.5 pts | 1.1 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 372 | 370 | 0.99× [0.98, 1.00] | -0.8% | [-2.0%, +0.4%] | 1.7 pts | 1.8 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 302 | 300 | 0.99× [0.97, 1.00] | -1.5% | [-3.3%, +0.3%] | 2.5 pts | 1.4 | yes |
| unique | uuid | 1048576 | valuesFor | ordered | baseline | 8 | 337 | 339 | 1.01× [1.00, 1.02] | +0.9% | [+0.1%, +1.6%] | 0.9 pts | 1.0 | yes |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 8 | 6828 | 6875 | 1.01× [0.99, 1.02] | +0.6% | [-0.5%, +1.8%] | 1.4 pts | 1.6 | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 8 | 1259 | 1273 | 1.00× [0.99, 1.01] | +0.2% | [-0.9%, +1.3%] | 1.3 pts | 1.2 | yes |
| unique | uuid | 1048576 | churn | ordered | baseline | 8 | 482 | 476 | 0.98× [0.97, 1.00] | -1.8% | [-3.6%, +0.0%] | 2.1 pts | 1.0 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.56% does not clear the 1.08% median noise floor of the processes
- multi email n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.58%, 0.45%] includes zero
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.18% does not clear the 1.47% median noise floor of the processes
- multi email n=4096 prefix: ordered vs baseline: the pooled difference of 0.65% does not clear the 0.81% median noise floor of the processes
- multi email n=4096 prefix: ordered vs baseline: the pooled interval [-0.11%, 1.41%] includes zero
- multi email n=4096 churn: ordered vs baseline: the pooled difference of 0.87% does not clear the 1.94% median noise floor of the processes
- multi email n=4096 build: ordered vs baseline: the pooled difference of -1.03% does not clear the 1.38% median noise floor of the processes
- multi email n=4096 build: ordered vs baseline: the pooled interval [-2.75%, 0.70%] includes zero
- multi email n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.22% does not clear the 1.63% median noise floor of the processes
- multi email n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.41%, 0.84%] includes zero
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.08% does not clear the 0.94% median noise floor of the processes
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.74%, 0.90%] includes zero
- multi email n=16384 prefix: ordered vs baseline: the pooled difference of 0.95% does not clear the 1.28% median noise floor of the processes
- multi email n=16384 churn: ordered vs baseline: the pooled difference of 0.55% does not clear the 2.40% median noise floor of the processes
- multi email n=16384 churn: ordered vs baseline: the pooled interval [-0.32%, 1.41%] includes zero
- multi email n=16384 build: ordered vs baseline: the pooled difference of -0.09% does not clear the 0.88% median noise floor of the processes
- multi email n=16384 build: ordered vs baseline: the pooled interval [-1.05%, 0.86%] includes zero
- multi email n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.26% does not clear the 1.91% median noise floor of the processes
- multi email n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.36%, 0.85%] includes zero
- multi email n=262144 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.22% does not clear the 1.71% median noise floor of the processes
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.67%, 1.10%] includes zero
- multi email n=262144 prefix: ordered vs baseline: the pooled difference of 0.73% does not clear the 1.44% median noise floor of the processes
- multi email n=262144 prefix: ordered vs baseline: the pooled interval [-0.30%, 1.75%] includes zero
- multi email n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.39% does not clear the 1.47% median noise floor of the processes
- multi email n=1048576 valuesFor: ordered vs baseline: the pooled interval [-1.09%, 0.31%] includes zero
- multi email n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.27% does not clear the 1.17% median noise floor of the processes
- multi email n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.72%, 0.18%] includes zero
- multi email n=1048576 churn: ordered vs baseline: the pooled difference of -0.84% does not clear the 6.35% median noise floor of the processes
- multi email n=1048576 churn: ordered vs baseline: the pooled interval [-4.56%, 2.87%] includes zero
- multi path n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.71% does not clear the 0.74% median noise floor of the processes
- multi path n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.17%, 0.76%] includes zero
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of -1.75% does not clear the 6.28% median noise floor of the processes
- multi path n=4096 prefix: ordered vs baseline: the pooled interval [-3.52%, 0.02%] includes zero
- multi path n=4096 churn: ordered vs baseline: the pooled difference of -0.22% does not clear the 2.12% median noise floor of the processes
- multi path n=4096 churn: ordered vs baseline: the pooled interval [-1.03%, 0.60%] includes zero
- multi path n=4096 build: ordered vs baseline: the pooled difference of -0.17% does not clear the 1.18% median noise floor of the processes
- multi path n=4096 build: ordered vs baseline: the pooled interval [-1.09%, 0.76%] includes zero
- multi path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.25% does not clear the 1.17% median noise floor of the processes
- multi path n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.48%, 0.97%] includes zero
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.29% does not clear the 1.66% median noise floor of the processes
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.48%, 1.07%] includes zero
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of 0.17% does not clear the 9.21% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-2.09%, 2.44%] includes zero
- multi path n=16384 churn: ordered vs baseline: the pooled difference of -0.37% does not clear the 1.82% median noise floor of the processes
- multi path n=16384 churn: ordered vs baseline: the pooled interval [-1.19%, 0.45%] includes zero
- multi path n=16384 build: ordered vs baseline: the pooled difference of -0.84% does not clear the 1.41% median noise floor of the processes
- multi path n=262144 valuesFor: ordered vs baseline: the pooled difference of -1.61% does not clear the 1.71% median noise floor of the processes
- multi path n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.65%, 0.43%] includes zero
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.87% does not clear the 2.19% median noise floor of the processes
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.28%, 0.54%] includes zero
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of 0.67% does not clear the 12.96% median noise floor of the processes
- multi path n=262144 prefix: ordered vs baseline: the pooled interval [-1.12%, 2.46%] includes zero
- multi path n=262144 churn: ordered vs baseline: the pooled difference of 0.02% does not clear the 2.32% median noise floor of the processes
- multi path n=262144 churn: ordered vs baseline: the pooled interval [-2.53%, 2.57%] includes zero
- multi str n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.39% does not clear the 0.94% median noise floor of the processes
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.62%, 0.84%] includes zero
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.62% does not clear the 2.04% median noise floor of the processes
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of 1.79% does not clear the 3.84% median noise floor of the processes
- multi str n=4096 churn: ordered vs baseline: the pooled difference of -0.21% does not clear the 1.93% median noise floor of the processes
- multi str n=4096 churn: ordered vs baseline: the pooled interval [-0.79%, 0.37%] includes zero
- multi str n=4096 build: ordered vs baseline: the pooled difference of 0.25% does not clear the 1.45% median noise floor of the processes
- multi str n=4096 build: ordered vs baseline: the pooled interval [-0.01%, 0.51%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.36% does not clear the 1.67% median noise floor of the processes
- multi str n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.79%, 0.08%] includes zero
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.39% does not clear the 1.52% median noise floor of the processes
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.34%, 0.57%] includes zero
- multi str n=16384 prefix: ordered vs baseline: the pooled difference of -0.71% does not clear the 1.52% median noise floor of the processes
- multi str n=16384 prefix: ordered vs baseline: the pooled interval [-1.81%, 0.39%] includes zero
- multi str n=16384 churn: ordered vs baseline: the pooled difference of 0.26% does not clear the 2.22% median noise floor of the processes
- multi str n=16384 churn: ordered vs baseline: the pooled interval [-1.56%, 2.09%] includes zero
- multi str n=16384 build: ordered vs baseline: the pooled difference of 0.08% does not clear the 0.97% median noise floor of the processes
- multi str n=16384 build: ordered vs baseline: the pooled interval [-0.48%, 0.64%] includes zero
- multi str n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.63% does not clear the 1.32% median noise floor of the processes
- multi str n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.22%, 1.47%] includes zero
- multi str n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.00% does not clear the 1.34% median noise floor of the processes
- multi str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.49%, 0.49%] includes zero
- multi str n=262144 prefix: ordered vs baseline: the pooled difference of -0.14% does not clear the 2.29% median noise floor of the processes
- multi str n=262144 prefix: ordered vs baseline: the pooled interval [-0.56%, 0.28%] includes zero
- multi str n=262144 churn: ordered vs baseline: the pooled interval [-0.05%, 5.50%] includes zero
- multi str n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.09% does not clear the 1.47% median noise floor of the processes
- multi str n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.80%, 0.63%] includes zero
- multi str n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.25% does not clear the 1.51% median noise floor of the processes
- multi str n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-1.21%, 0.72%] includes zero
- multi str n=1048576 prefix: ordered vs baseline: the pooled difference of -0.17% does not clear the 0.62% median noise floor of the processes
- multi str n=1048576 prefix: ordered vs baseline: the pooled interval [-0.44%, 0.11%] includes zero
- multi str n=1048576 churn: ordered vs baseline: the pooled difference of -0.19% does not clear the 2.37% median noise floor of the processes
- multi str n=1048576 churn: ordered vs baseline: the pooled interval [-6.61%, 6.24%] includes zero
- multi str n=1048576 churn: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 churn: ordered vs baseline: 3 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.04% does not clear the 2.26% median noise floor of the processes
- multi street n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.71%, 1.78%] includes zero
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.82% does not clear the 1.53% median noise floor of the processes
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.77%, 0.12%] includes zero
- multi street n=4096 prefix: ordered vs baseline: the pooled difference of 0.28% does not clear the 2.84% median noise floor of the processes
- multi street n=4096 prefix: ordered vs baseline: the pooled interval [-0.59%, 1.15%] includes zero
- multi street n=4096 churn: ordered vs baseline: the pooled difference of -0.10% does not clear the 1.27% median noise floor of the processes
- multi street n=4096 churn: ordered vs baseline: the pooled interval [-0.54%, 0.33%] includes zero
- multi street n=4096 build: ordered vs baseline: the pooled difference of -0.70% does not clear the 1.65% median noise floor of the processes
- multi street n=4096 build: ordered vs baseline: the pooled interval [-1.67%, 0.27%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.24% does not clear the 1.27% median noise floor of the processes
- multi street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.45%, 0.94%] includes zero
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.38% does not clear the 1.17% median noise floor of the processes
- multi street n=16384 prefix: ordered vs baseline: the pooled difference of -0.02% does not clear the 3.65% median noise floor of the processes
- multi street n=16384 prefix: ordered vs baseline: the pooled interval [-1.85%, 1.81%] includes zero
- multi street n=16384 churn: ordered vs baseline: the pooled difference of 0.05% does not clear the 1.51% median noise floor of the processes
- multi street n=16384 churn: ordered vs baseline: the pooled interval [-0.61%, 0.71%] includes zero
- multi street n=16384 build: ordered vs baseline: the pooled difference of -0.58% does not clear the 0.69% median noise floor of the processes
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.92% does not clear the 1.67% median noise floor of the processes
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.67% does not clear the 1.38% median noise floor of the processes
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.86%, 0.53%] includes zero
- multi u64 n=4096 churn: ordered vs baseline: the pooled difference of -0.03% does not clear the 1.25% median noise floor of the processes
- multi u64 n=4096 churn: ordered vs baseline: the pooled interval [-0.80%, 0.74%] includes zero
- multi u64 n=4096 build: ordered vs baseline: the pooled difference of 0.18% does not clear the 1.90% median noise floor of the processes
- multi u64 n=4096 build: ordered vs baseline: the pooled interval [-1.71%, 2.06%] includes zero
- multi u64 n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.33% does not clear the 1.21% median noise floor of the processes
- multi u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.65%, 1.30%] includes zero
- multi u64 n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.64% does not clear the 1.51% median noise floor of the processes
- multi u64 n=16384 churn: ordered vs baseline: the pooled difference of -0.12% does not clear the 1.85% median noise floor of the processes
- multi u64 n=16384 churn: ordered vs baseline: the pooled interval [-1.21%, 0.97%] includes zero
- multi u64 n=16384 build: ordered vs baseline: the pooled difference of -0.70% does not clear the 1.67% median noise floor of the processes
- multi u64 n=16384 build: ordered vs baseline: the pooled interval [-1.80%, 0.41%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the pooled difference of 1.78% does not clear the 1.90% median noise floor of the processes
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.39% does not clear the 1.39% median noise floor of the processes
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.93%, 1.70%] includes zero
- multi u64 n=262144 churn: ordered vs baseline: the pooled difference of 0.66% does not clear the 2.10% median noise floor of the processes
- multi u64 n=262144 churn: ordered vs baseline: the pooled interval [-2.76%, 4.09%] includes zero
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.56% does not clear the 1.22% median noise floor of the processes
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-1.15%, 0.02%] includes zero
- multi u64 n=1048576 churn: ordered vs baseline: the pooled difference of -0.44% does not clear the 6.60% median noise floor of the processes
- multi u64 n=1048576 churn: ordered vs baseline: the pooled interval [-4.28%, 3.40%] includes zero
- multi url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.27% does not clear the 1.25% median noise floor of the processes
- multi url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.08%, 0.61%] includes zero
- multi url n=4096 prefix: ordered vs baseline: the pooled difference of -0.64% does not clear the 1.74% median noise floor of the processes
- multi url n=4096 prefix: ordered vs baseline: the pooled interval [-1.93%, 0.66%] includes zero
- multi url n=4096 churn: ordered vs baseline: the pooled difference of -0.63% does not clear the 2.10% median noise floor of the processes
- multi url n=4096 churn: ordered vs baseline: the pooled interval [-1.93%, 0.68%] includes zero
- multi url n=4096 build: ordered vs baseline: the pooled difference of -0.27% does not clear the 1.08% median noise floor of the processes
- multi url n=4096 build: ordered vs baseline: the pooled interval [-0.75%, 0.20%] includes zero
- multi url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.09% does not clear the 1.22% median noise floor of the processes
- multi url n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.58%, 1.40%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.37% does not clear the 1.51% median noise floor of the processes
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.04%, 1.78%] includes zero
- multi url n=16384 prefix: ordered vs baseline: the pooled difference of 0.42% does not clear the 2.53% median noise floor of the processes
- multi url n=16384 prefix: ordered vs baseline: the pooled interval [-0.55%, 1.38%] includes zero
- multi url n=16384 churn: ordered vs baseline: the pooled difference of -0.42% does not clear the 1.83% median noise floor of the processes
- multi url n=16384 churn: ordered vs baseline: the pooled interval [-1.72%, 0.89%] includes zero
- multi url n=16384 build: ordered vs baseline: the pooled difference of -0.03% does not clear the 1.37% median noise floor of the processes
- multi url n=16384 build: ordered vs baseline: the pooled interval [-0.73%, 0.68%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.55% does not clear the 2.07% median noise floor of the processes
- multi url n=262144 valuesFor: ordered vs baseline: the pooled interval [-5.02%, 3.92%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.33% does not clear the 1.53% median noise floor of the processes
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.28%, 0.62%] includes zero
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of -0.96% does not clear the 6.28% median noise floor of the processes
- multi url n=262144 prefix: ordered vs baseline: the pooled interval [-3.41%, 1.49%] includes zero
- multi url n=262144 churn: ordered vs baseline: the pooled difference of -2.23% does not clear the 2.44% median noise floor of the processes
- multi url n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.20% does not clear the 1.19% median noise floor of the processes
- multi url n=1048576 valuesFor: ordered vs baseline: the pooled interval [-2.10%, 1.70%] includes zero
- multi url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.28% does not clear the 1.21% median noise floor of the processes
- multi url n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.56%, 1.11%] includes zero
- multi url n=1048576 prefix: ordered vs baseline: the pooled difference of -0.09% does not clear the 9.27% median noise floor of the processes
- multi url n=1048576 prefix: ordered vs baseline: the pooled interval [-1.82%, 1.65%] includes zero
- multi url n=1048576 churn: ordered vs baseline: the pooled difference of 2.45% does not clear the 2.75% median noise floor of the processes
- multi url n=1048576 churn: ordered vs baseline: the pooled interval [-1.39%, 6.28%] includes zero
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.16% does not clear the 0.82% median noise floor of the processes
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.03%, 0.72%] includes zero
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.00% does not clear the 2.03% median noise floor of the processes
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.16%, 1.15%] includes zero
- multi uuid n=4096 prefix: ordered vs baseline: the pooled difference of 0.68% does not clear the 0.82% median noise floor of the processes
- multi uuid n=4096 churn: ordered vs baseline: the pooled difference of -0.30% does not clear the 1.18% median noise floor of the processes
- multi uuid n=4096 churn: ordered vs baseline: the pooled interval [-0.87%, 0.28%] includes zero
- multi uuid n=4096 build: ordered vs baseline: the pooled difference of -0.27% does not clear the 1.82% median noise floor of the processes
- multi uuid n=4096 build: ordered vs baseline: the pooled interval [-0.94%, 0.41%] includes zero
- multi uuid n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.01% does not clear the 1.93% median noise floor of the processes
- multi uuid n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.94%, 0.92%] includes zero
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.30% does not clear the 1.58% median noise floor of the processes
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.83%, 0.23%] includes zero
- multi uuid n=16384 prefix: ordered vs baseline: the pooled interval [-0.37%, 2.31%] includes zero
- multi uuid n=16384 churn: ordered vs baseline: the pooled difference of 0.52% does not clear the 2.00% median noise floor of the processes
- multi uuid n=16384 churn: ordered vs baseline: the pooled interval [-0.81%, 1.84%] includes zero
- multi uuid n=16384 build: ordered vs baseline: the pooled difference of -0.29% does not clear the 2.13% median noise floor of the processes
- multi uuid n=16384 build: ordered vs baseline: the pooled interval [-1.82%, 1.24%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.15% does not clear the 1.46% median noise floor of the processes
- multi uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.66%, 1.37%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.45% does not clear the 2.09% median noise floor of the processes
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.40%, 0.50%] includes zero
- multi uuid n=262144 prefix: ordered vs baseline: the pooled difference of -0.59% does not clear the 1.65% median noise floor of the processes
- multi uuid n=262144 prefix: ordered vs baseline: the pooled interval [-1.82%, 0.64%] includes zero
- multi uuid n=262144 churn: ordered vs baseline: the pooled difference of 1.94% does not clear the 2.43% median noise floor of the processes
- multi uuid n=262144 churn: ordered vs baseline: the pooled interval [-2.97%, 6.86%] includes zero
- multi uuid n=262144 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.66% does not clear the 1.24% median noise floor of the processes
- multi uuid n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.32%, 1.65%] includes zero
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.05% does not clear the 1.81% median noise floor of the processes
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.53%, 0.43%] includes zero
- multi uuid n=1048576 prefix: ordered vs baseline: the pooled difference of -0.15% does not clear the 1.15% median noise floor of the processes
- multi uuid n=1048576 prefix: ordered vs baseline: the pooled interval [-0.99%, 0.68%] includes zero
- multi uuid n=1048576 churn: ordered vs baseline: the pooled difference of 0.40% does not clear the 3.47% median noise floor of the processes
- multi uuid n=1048576 churn: ordered vs baseline: the pooled interval [-5.74%, 6.54%] includes zero
- multi uuid n=1048576 churn: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.45% does not clear the 1.54% median noise floor of the processes
- unique email n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.77%, 0.86%] includes zero
- unique email n=4096 prefix: ordered vs baseline: the pooled difference of 1.29% does not clear the 1.52% median noise floor of the processes
- unique email n=4096 churn: ordered vs baseline: the pooled difference of -0.29% does not clear the 0.76% median noise floor of the processes
- unique email n=4096 churn: ordered vs baseline: the pooled interval [-0.88%, 0.29%] includes zero
- unique email n=4096 build: ordered vs baseline: the pooled difference of -0.98% does not clear the 2.16% median noise floor of the processes
- unique email n=4096 build: ordered vs baseline: the pooled interval [-2.94%, 0.97%] includes zero
- unique email n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.51% does not clear the 1.00% median noise floor of the processes
- unique email n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.55%, 0.52%] includes zero
- unique email n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.06% does not clear the 1.06% median noise floor of the processes
- unique email n=16384 churn: ordered vs baseline: the pooled difference of 0.72% does not clear the 1.55% median noise floor of the processes
- unique email n=16384 churn: ordered vs baseline: the pooled interval [-0.25%, 1.69%] includes zero
- unique email n=16384 build: ordered vs baseline: the pooled difference of 0.01% does not clear the 1.47% median noise floor of the processes
- unique email n=16384 build: ordered vs baseline: the pooled interval [-0.94%, 0.96%] includes zero
- unique email n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.74% does not clear the 1.62% median noise floor of the processes
- unique email n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.05%, 3.53%] includes zero
- unique email n=262144 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.12% does not clear the 2.02% median noise floor of the processes
- unique email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.73%, 0.97%] includes zero
- unique email n=262144 prefix: ordered vs baseline: the pooled difference of -0.93% does not clear the 1.64% median noise floor of the processes
- unique email n=262144 churn: ordered vs baseline: the pooled difference of -0.73% does not clear the 1.81% median noise floor of the processes
- unique email n=262144 churn: ordered vs baseline: the pooled interval [-1.80%, 0.34%] includes zero
- unique email n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.47% does not clear the 1.04% median noise floor of the processes
- unique email n=1048576 valuesFor: ordered vs baseline: the pooled interval [-1.48%, 0.53%] includes zero
- unique email n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.77% does not clear the 1.44% median noise floor of the processes
- unique email n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-1.65%, 0.12%] includes zero
- unique email n=1048576 prefix: ordered vs baseline: the pooled difference of 0.32% does not clear the 1.41% median noise floor of the processes
- unique email n=1048576 prefix: ordered vs baseline: the pooled interval [-0.83%, 1.46%] includes zero
- unique email n=1048576 churn: ordered vs baseline: the pooled difference of -0.08% does not clear the 2.19% median noise floor of the processes
- unique email n=1048576 churn: ordered vs baseline: the pooled interval [-1.83%, 1.67%] includes zero
- unique email n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.32% does not clear the 1.23% median noise floor of the processes
- unique path n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.70%, 1.33%] includes zero
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.12% does not clear the 1.50% median noise floor of the processes
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.33%, 1.58%] includes zero
- unique path n=4096 prefix: ordered vs baseline: the pooled difference of -0.55% does not clear the 1.77% median noise floor of the processes
- unique path n=4096 prefix: ordered vs baseline: the pooled interval [-1.51%, 0.41%] includes zero
- unique path n=4096 churn: ordered vs baseline: the pooled difference of 0.74% does not clear the 2.56% median noise floor of the processes
- unique path n=4096 build: ordered vs baseline: the pooled difference of -0.70% does not clear the 1.53% median noise floor of the processes
- unique path n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.14% does not clear the 0.93% median noise floor of the processes
- unique path n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.01%, 0.73%] includes zero
- unique path n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.14% does not clear the 1.67% median noise floor of the processes
- unique path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.68%, 0.39%] includes zero
- unique path n=16384 prefix: ordered vs baseline: the pooled difference of -0.09% does not clear the 5.82% median noise floor of the processes
- unique path n=16384 prefix: ordered vs baseline: the pooled interval [-1.79%, 1.60%] includes zero
- unique path n=16384 churn: ordered vs baseline: the pooled difference of 0.34% does not clear the 2.54% median noise floor of the processes
- unique path n=16384 churn: ordered vs baseline: the pooled interval [-0.91%, 1.60%] includes zero
- unique path n=16384 build: ordered vs baseline: the pooled difference of 0.24% does not clear the 1.23% median noise floor of the processes
- unique path n=16384 build: ordered vs baseline: the pooled interval [-0.20%, 0.68%] includes zero
- unique path n=262144 valuesFor: ordered vs baseline: the pooled difference of 1.31% does not clear the 1.32% median noise floor of the processes
- unique path n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.84% does not clear the 1.99% median noise floor of the processes
- unique path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.51%, 2.20%] includes zero
- unique path n=262144 prefix: ordered vs baseline: the pooled difference of 1.39% does not clear the 11.56% median noise floor of the processes
- unique path n=262144 prefix: ordered vs baseline: the pooled interval [-1.99%, 4.78%] includes zero
- unique path n=262144 churn: ordered vs baseline: the pooled difference of -0.12% does not clear the 1.62% median noise floor of the processes
- unique path n=262144 churn: ordered vs baseline: the pooled interval [-1.82%, 1.58%] includes zero
- unique str n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.12% does not clear the 1.75% median noise floor of the processes
- unique str n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.80%, 0.57%] includes zero
- unique str n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.85% does not clear the 1.90% median noise floor of the processes
- unique str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.18%, 0.48%] includes zero
- unique str n=4096 prefix: ordered vs baseline: the pooled difference of 0.29% does not clear the 4.55% median noise floor of the processes
- unique str n=4096 prefix: ordered vs baseline: the pooled interval [-1.75%, 2.34%] includes zero
- unique str n=4096 churn: ordered vs baseline: the pooled difference of 0.10% does not clear the 1.06% median noise floor of the processes
- unique str n=4096 churn: ordered vs baseline: the pooled interval [-0.54%, 0.75%] includes zero
- unique str n=4096 build: ordered vs baseline: the pooled difference of -0.17% does not clear the 2.66% median noise floor of the processes
- unique str n=4096 build: ordered vs baseline: the pooled interval [-1.22%, 0.88%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.14% does not clear the 1.04% median noise floor of the processes
- unique str n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.90%, 0.62%] includes zero
- unique str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.34% does not clear the 1.59% median noise floor of the processes
- unique str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.82%, 0.14%] includes zero
- unique str n=16384 prefix: ordered vs baseline: the pooled difference of -0.34% does not clear the 1.35% median noise floor of the processes
- unique str n=16384 prefix: ordered vs baseline: the pooled interval [-0.82%, 0.13%] includes zero
- unique str n=16384 churn: ordered vs baseline: the pooled difference of 1.02% does not clear the 2.52% median noise floor of the processes
- unique str n=16384 churn: ordered vs baseline: the pooled interval [-0.06%, 2.10%] includes zero
- unique str n=16384 build: ordered vs baseline: the pooled difference of -0.53% does not clear the 2.30% median noise floor of the processes
- unique str n=16384 build: ordered vs baseline: the pooled interval [-1.34%, 0.29%] includes zero
- unique str n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.19% does not clear the 2.39% median noise floor of the processes
- unique str n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.02%, 1.41%] includes zero
- unique str n=262144 valuesBetween: ordered vs baseline: the pooled difference of -1.00% does not clear the 1.95% median noise floor of the processes
- unique str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.61%, 0.60%] includes zero
- unique str n=262144 prefix: ordered vs baseline: the pooled difference of -2.06% does not clear the 8.72% median noise floor of the processes
- unique str n=262144 prefix: ordered vs baseline: the pooled interval [-4.65%, 0.53%] includes zero
- unique str n=262144 churn: ordered vs baseline: the pooled difference of 0.33% does not clear the 1.96% median noise floor of the processes
- unique str n=262144 churn: ordered vs baseline: the pooled interval [-0.76%, 1.43%] includes zero
- unique str n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.52% does not clear the 1.43% median noise floor of the processes
- unique str n=1048576 valuesFor: ordered vs baseline: the pooled interval [-2.39%, 1.36%] includes zero
- unique str n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.15% does not clear the 1.19% median noise floor of the processes
- unique str n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-1.94%, 1.63%] includes zero
- unique str n=1048576 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique str n=1048576 prefix: ordered vs baseline: the pooled difference of -0.42% does not clear the 0.71% median noise floor of the processes
- unique str n=1048576 prefix: ordered vs baseline: the pooled interval [-2.06%, 1.22%] includes zero
- unique str n=1048576 churn: ordered vs baseline: the pooled difference of -1.64% does not clear the 2.06% median noise floor of the processes
- unique street n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.26% does not clear the 1.26% median noise floor of the processes
- unique street n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.83%, 1.35%] includes zero
- unique street n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.08% does not clear the 2.00% median noise floor of the processes
- unique street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.70%, 1.53%] includes zero
- unique street n=4096 prefix: ordered vs baseline: the pooled difference of 0.80% does not clear the 2.84% median noise floor of the processes
- unique street n=4096 prefix: ordered vs baseline: the pooled interval [-0.67%, 2.28%] includes zero
- unique street n=4096 churn: ordered vs baseline: the pooled difference of 0.07% does not clear the 1.14% median noise floor of the processes
- unique street n=4096 churn: ordered vs baseline: the pooled interval [-1.25%, 1.38%] includes zero
- unique street n=4096 build: ordered vs baseline: the pooled difference of -0.49% does not clear the 1.09% median noise floor of the processes
- unique street n=4096 build: ordered vs baseline: the pooled interval [-2.16%, 1.18%] includes zero
- unique street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.33% does not clear the 0.85% median noise floor of the processes
- unique street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.18%, 0.52%] includes zero
- unique street n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.83% does not clear the 1.42% median noise floor of the processes
- unique street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.95%, 0.29%] includes zero
- unique street n=16384 prefix: ordered vs baseline: the pooled difference of -0.35% does not clear the 3.88% median noise floor of the processes
- unique street n=16384 prefix: ordered vs baseline: the pooled interval [-1.86%, 1.16%] includes zero
- unique street n=16384 churn: ordered vs baseline: the pooled difference of 1.04% does not clear the 1.80% median noise floor of the processes
- unique street n=16384 churn: ordered vs baseline: the pooled interval [-0.00%, 2.08%] includes zero
- unique street n=16384 build: ordered vs baseline: the pooled difference of 0.16% does not clear the 1.04% median noise floor of the processes
- unique street n=16384 build: ordered vs baseline: the pooled interval [-0.46%, 0.78%] includes zero
- unique u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.41% does not clear the 2.81% median noise floor of the processes
- unique u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.19%, 2.02%] includes zero
- unique u64 n=4096 churn: ordered vs baseline: the pooled difference of -0.26% does not clear the 1.18% median noise floor of the processes
- unique u64 n=4096 churn: ordered vs baseline: the pooled interval [-1.50%, 0.99%] includes zero
- unique u64 n=4096 build: ordered vs baseline: the pooled difference of -0.54% does not clear the 2.84% median noise floor of the processes
- unique u64 n=4096 build: ordered vs baseline: the pooled interval [-1.94%, 0.86%] includes zero
- unique u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.39% does not clear the 2.22% median noise floor of the processes
- unique u64 n=16384 churn: ordered vs baseline: the pooled difference of 0.63% does not clear the 2.34% median noise floor of the processes
- unique u64 n=16384 churn: ordered vs baseline: the pooled interval [-0.59%, 1.86%] includes zero
- unique u64 n=16384 build: ordered vs baseline: the pooled difference of -0.21% does not clear the 2.52% median noise floor of the processes
- unique u64 n=16384 build: ordered vs baseline: the pooled interval [-1.02%, 0.61%] includes zero
- unique u64 n=262144 valuesFor: ordered vs baseline: the pooled difference of -1.99% does not clear the 2.55% median noise floor of the processes
- unique u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-4.12%, 0.14%] includes zero
- unique u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.41% does not clear the 3.78% median noise floor of the processes
- unique u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.91%, 0.09%] includes zero
- unique u64 n=262144 churn: ordered vs baseline: the pooled difference of 0.58% does not clear the 3.30% median noise floor of the processes
- unique u64 n=262144 churn: ordered vs baseline: the pooled interval [-1.49%, 2.65%] includes zero
- unique u64 n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.64% does not clear the 0.95% median noise floor of the processes
- unique u64 n=1048576 valuesFor: ordered vs baseline: the pooled interval [-1.33%, 0.05%] includes zero
- unique u64 n=1048576 churn: ordered vs baseline: the pooled difference of -1.25% does not clear the 1.97% median noise floor of the processes
- unique u64 n=1048576 churn: ordered vs baseline: the pooled interval [-3.35%, 0.85%] includes zero
- unique u64 n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.66% does not clear the 1.24% median noise floor of the processes
- unique url n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.46%, 0.14%] includes zero
- unique url n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.74% does not clear the 1.33% median noise floor of the processes
- unique url n=4096 prefix: ordered vs baseline: the pooled difference of -0.40% does not clear the 1.59% median noise floor of the processes
- unique url n=4096 prefix: ordered vs baseline: the pooled interval [-1.51%, 0.71%] includes zero
- unique url n=4096 churn: ordered vs baseline: the pooled difference of 0.30% does not clear the 1.05% median noise floor of the processes
- unique url n=4096 churn: ordered vs baseline: the pooled interval [-0.75%, 1.35%] includes zero
- unique url n=4096 build: ordered vs baseline: the pooled difference of -0.02% does not clear the 3.07% median noise floor of the processes
- unique url n=4096 build: ordered vs baseline: the pooled interval [-0.94%, 0.90%] includes zero
- unique url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.22% does not clear the 1.35% median noise floor of the processes
- unique url n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.05%, 0.60%] includes zero
- unique url n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.09% does not clear the 1.84% median noise floor of the processes
- unique url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.39%, 1.21%] includes zero
- unique url n=16384 prefix: ordered vs baseline: the pooled difference of -0.13% does not clear the 1.64% median noise floor of the processes
- unique url n=16384 prefix: ordered vs baseline: the pooled interval [-1.24%, 0.99%] includes zero
- unique url n=16384 churn: ordered vs baseline: the pooled difference of -1.40% does not clear the 1.80% median noise floor of the processes
- unique url n=16384 build: ordered vs baseline: the pooled difference of 0.28% does not clear the 0.92% median noise floor of the processes
- unique url n=16384 build: ordered vs baseline: the pooled interval [-0.48%, 1.04%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the pooled interval [-6.48%, 3.60%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.08% does not clear the 1.52% median noise floor of the processes
- unique url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.63%, 1.80%] includes zero
- unique url n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 prefix: ordered vs baseline: the pooled difference of -0.77% does not clear the 3.77% median noise floor of the processes
- unique url n=262144 prefix: ordered vs baseline: the pooled interval [-2.98%, 1.43%] includes zero
- unique url n=262144 churn: ordered vs baseline: the pooled difference of -0.69% does not clear the 2.06% median noise floor of the processes
- unique url n=262144 churn: ordered vs baseline: the pooled interval [-1.56%, 0.18%] includes zero
- unique url n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.11% does not clear the 1.11% median noise floor of the processes
- unique url n=1048576 valuesFor: ordered vs baseline: the pooled interval [-3.39%, 3.61%] includes zero
- unique url n=1048576 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=1048576 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.00% does not clear the 1.14% median noise floor of the processes
- unique url n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-1.23%, 1.23%] includes zero
- unique url n=1048576 prefix: ordered vs baseline: the pooled difference of 1.73% does not clear the 8.63% median noise floor of the processes
- unique url n=1048576 prefix: ordered vs baseline: the pooled interval [-1.04%, 4.50%] includes zero
- unique url n=1048576 churn: ordered vs baseline: the pooled difference of 0.66% does not clear the 2.13% median noise floor of the processes
- unique url n=1048576 churn: ordered vs baseline: the pooled interval [-0.70%, 2.03%] includes zero
- unique uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.30% does not clear the 2.06% median noise floor of the processes
- unique uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.05%, 1.66%] includes zero
- unique uuid n=4096 prefix: ordered vs baseline: the pooled difference of -0.13% does not clear the 1.26% median noise floor of the processes
- unique uuid n=4096 prefix: ordered vs baseline: the pooled interval [-0.60%, 0.35%] includes zero
- unique uuid n=4096 churn: ordered vs baseline: the pooled difference of 0.06% does not clear the 1.17% median noise floor of the processes
- unique uuid n=4096 churn: ordered vs baseline: the pooled interval [-0.67%, 0.80%] includes zero
- unique uuid n=4096 build: ordered vs baseline: the pooled difference of 0.06% does not clear the 2.10% median noise floor of the processes
- unique uuid n=4096 build: ordered vs baseline: the pooled interval [-0.55%, 0.67%] includes zero
- unique uuid n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.70% does not clear the 0.89% median noise floor of the processes
- unique uuid n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.49%, 0.10%] includes zero
- unique uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.12% does not clear the 1.31% median noise floor of the processes
- unique uuid n=16384 prefix: ordered vs baseline: the pooled difference of -0.39% does not clear the 1.11% median noise floor of the processes
- unique uuid n=16384 prefix: ordered vs baseline: the pooled interval [-1.29%, 0.50%] includes zero
- unique uuid n=16384 prefix: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=16384 churn: ordered vs baseline: the pooled difference of 1.22% does not clear the 2.99% median noise floor of the processes
- unique uuid n=16384 churn: ordered vs baseline: the pooled interval [-0.49%, 2.94%] includes zero
- unique uuid n=16384 build: ordered vs baseline: the pooled difference of -0.66% does not clear the 1.64% median noise floor of the processes
- unique uuid n=16384 build: ordered vs baseline: the pooled interval [-1.43%, 0.10%] includes zero
- unique uuid n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.02% does not clear the 1.53% median noise floor of the processes
- unique uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.37%, 1.32%] includes zero
- unique uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.11% does not clear the 2.11% median noise floor of the processes
- unique uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.93%, 1.16%] includes zero
- unique uuid n=262144 prefix: ordered vs baseline: the pooled difference of -0.80% does not clear the 1.98% median noise floor of the processes
- unique uuid n=262144 prefix: ordered vs baseline: the pooled interval [-2.03%, 0.43%] includes zero
- unique uuid n=262144 churn: ordered vs baseline: the pooled difference of -1.48% does not clear the 1.80% median noise floor of the processes
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-3.30%, 0.35%] includes zero
- unique uuid n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.86% does not clear the 1.45% median noise floor of the processes
- unique uuid n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.61% does not clear the 1.03% median noise floor of the processes
- unique uuid n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.54%, 1.76%] includes zero
- unique uuid n=1048576 prefix: ordered vs baseline: the pooled difference of 0.23% does not clear the 0.93% median noise floor of the processes
- unique uuid n=1048576 prefix: ordered vs baseline: the pooled interval [-0.86%, 1.32%] includes zero
- unique uuid n=1048576 churn: ordered vs baseline: the pooled difference of -1.76% does not clear the 2.01% median noise floor of the processes
- unique uuid n=1048576 churn: ordered vs baseline: the pooled interval [-3.55%, 0.03%] includes zero
