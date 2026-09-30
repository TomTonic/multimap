| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi-str | email | 4096 | valuesFor | ordered | baseline | 6 | 54.3 | 53.0 | 0.97× [0.96, 0.99] | -2.6% | [-3.8%, -1.4%] | 1.1 pts | 1.1 | yes |
| multi-str | email | 4096 | valuesBetween | ordered | baseline | 6 | 3315 | 3282 | 0.99× [0.98, 1.00] | -1.2% | [-2.5%, +0.1%] | 1.2 pts | 0.7 | yes |
| multi-str | email | 4096 | prefix | ordered | baseline | 6 | 83.4 | 79.3 | 0.95× [0.95, 0.96] | -5.1% | [-5.7%, -4.5%] | 0.5 pts | 0.7 | yes |
| multi-str | email | 4096 | churn | ordered | baseline | 6 | 83.5 | 81.8 | 0.98× [0.97, 0.99] | -1.7% | [-2.6%, -0.8%] | 0.9 pts | 0.5 | yes |
| multi-str | email | 4096 | build | ordered | baseline | 6 | 7.04 ms | 6.87 ms | 0.98× [0.97, 0.98] | -2.4% | [-3.2%, -1.6%] | 0.8 pts | 0.9 | yes |
| multi-str | email | 16384 | valuesFor | ordered | baseline | 6 | 68.2 | 65.6 | 0.96× [0.95, 0.97] | -3.7% | [-4.8%, -2.6%] | 1.1 pts | 0.7 | yes |
| multi-str | email | 16384 | valuesBetween | ordered | baseline | 6 | 3921 | 3875 | 0.98× [0.97, 1.00] | -1.5% | [-2.9%, -0.2%] | 1.3 pts | 0.8 | yes |
| multi-str | email | 16384 | prefix | ordered | baseline | 6 | 104 | 98.7 | 0.95× [0.94, 0.95] | -5.4% | [-5.9%, -5.0%] | 0.4 pts | 0.3 | yes |
| multi-str | email | 16384 | churn | ordered | baseline | 6 | 139 | 134 | 0.98× [0.96, 1.00] | -2.2% | [-4.1%, -0.3%] | 1.8 pts | 0.9 | yes |
| multi-str | email | 16384 | build | ordered | baseline | 6 | 41.26 ms | 40.16 ms | 0.97× [0.95, 0.98] | -3.3% | [-4.8%, -1.8%] | 1.5 pts | 1.2 | yes |
| multi-str | email | 262144 | valuesFor | ordered | baseline | 10 | 301 | 301 | 1.01× [1.00, 1.02] | +1.0% | [+0.1%, +1.8%] | 1.2 pts | 1.2 | yes |
| multi-str | email | 262144 | valuesBetween | ordered | baseline | 10 | 11.6 µs | 11.4 µs | 0.99× [0.98, 1.00] | -1.0% | [-1.9%, -0.1%] | 1.3 pts | 1.0 | yes |
| multi-str | email | 262144 | prefix | ordered | baseline | 10 | 414 | 406 | 0.98× [0.98, 0.99] | -1.7% | [-2.5%, -0.8%] | 1.2 pts | 1.0 | yes |
| multi-str | email | 262144 | churn | ordered | baseline | 10 | 532 | 529 | 0.98× [0.95, 1.02] | -2.0% | [-5.5%, +1.5%] | 4.9 pts | 0.9 | no |
| multi-str | email | 1048576 | valuesFor | ordered | baseline | 10 | 467 | 470 | 1.01× [1.00, 1.01] | +0.7% | [+0.3%, +1.1%] | 0.5 pts | 0.6 | yes |
| multi-str | email | 1048576 | valuesBetween | ordered | baseline | 10 | 13.6 µs | 13.6 µs | 1.00× [0.99, 1.01] | +0.1% | [-0.9%, +1.1%] | 1.4 pts | 1.2 | yes |
| multi-str | email | 1048576 | prefix | ordered | baseline | 10 | 785 | 773 | 0.99× [0.98, 0.99] | -1.4% | [-2.0%, -0.8%] | 0.8 pts | 0.8 | yes |
| multi-str | email | 1048576 | churn | ordered | baseline | 10 | 745 | 703 | 0.95× [0.93, 0.97] | -5.5% | [-7.8%, -3.3%] | 3.2 pts | 0.6 | no |
| multi-str | path | 4096 | valuesFor | ordered | baseline | 10 | 117 | 114 | 0.97× [0.96, 0.98] | -2.9% | [-3.9%, -1.8%] | 1.5 pts | 1.4 | yes |
| multi-str | path | 4096 | valuesBetween | ordered | baseline | 10 | 4882 | 4787 | 0.98× [0.98, 0.99] | -1.9% | [-2.3%, -1.4%] | 0.6 pts | 0.4 | yes |
| multi-str | path | 4096 | prefix | ordered | baseline | 10 | 368 | 361 | 0.98× [0.97, 1.00] | -1.6% | [-3.4%, +0.2%] | 2.5 pts | 0.8 | yes |
| multi-str | path | 4096 | churn | ordered | baseline | 10 | 181 | 179 | 0.99× [0.99, 1.00] | -0.7% | [-1.1%, -0.2%] | 0.6 pts | 0.4 | yes |
| multi-str | path | 4096 | build | ordered | baseline | 10 | 13.71 ms | 13.31 ms | 0.97× [0.96, 0.98] | -3.2% | [-3.8%, -2.5%] | 0.9 pts | 1.2 | yes |
| multi-str | path | 16384 | valuesFor | ordered | baseline | 10 | 152 | 150 | 0.99× [0.98, 1.00] | -1.0% | [-2.0%, +0.0%] | 1.4 pts | 1.1 | yes |
| multi-str | path | 16384 | valuesBetween | ordered | baseline | 10 | 5490 | 5477 | 1.00× [0.99, 1.00] | -0.4% | [-1.0%, +0.1%] | 0.8 pts | 0.6 | yes |
| multi-str | path | 16384 | prefix | ordered | baseline | 10 | 698 | 712 | 1.02× [0.99, 1.04] | +1.7% | [-0.7%, +4.1%] | 3.3 pts | 0.6 | no |
| multi-str | path | 16384 | churn | ordered | baseline | 10 | 294 | 299 | 1.00× [0.99, 1.02] | +0.4% | [-1.0%, +1.8%] | 1.9 pts | 1.3 | yes |
| multi-str | path | 16384 | build | ordered | baseline | 10 | 81.83 ms | 81.96 ms | 1.00× [0.99, 1.01] | +0.2% | [-0.7%, +1.1%] | 1.2 pts | 1.2 | yes |
| multi-str | path | 262144 | valuesFor | ordered | baseline | 10 | 550 | 588 | 1.07× [1.06, 1.08] | +6.7% | [+6.0%, +7.5%] | 1.0 pts | 0.9 | yes |
| multi-str | path | 262144 | valuesBetween | ordered | baseline | 10 | 15.1 µs | 15.6 µs | 1.02× [1.02, 1.03] | +2.4% | [+1.6%, +3.2%] | 1.1 pts | 0.9 | yes |
| multi-str | path | 262144 | prefix | ordered | baseline | 10 | 12.9 µs | 13.3 µs | 1.02× [1.00, 1.04] | +2.2% | [+0.2%, +4.3%] | 2.9 pts | 0.3 | no |
| multi-str | path | 262144 | churn | ordered | baseline | 10 | 952 | 950 | 0.98× [0.95, 1.01] | -2.3% | [-5.3%, +0.7%] | 4.1 pts | 2.3 | no |
| multi-str | str | 4096 | valuesFor | ordered | baseline | 6 | 66.9 | 64.0 | 0.95× [0.95, 0.96] | -4.7% | [-5.3%, -4.1%] | 0.6 pts | 0.5 | yes |
| multi-str | str | 4096 | valuesBetween | ordered | baseline | 6 | 3922 | 3837 | 0.98× [0.97, 0.98] | -2.5% | [-3.3%, -1.6%] | 0.8 pts | 0.4 | yes |
| multi-str | str | 4096 | prefix | ordered | baseline | 6 | 9056 | 8697 | 0.96× [0.95, 0.97] | -4.4% | [-5.5%, -3.2%] | 1.1 pts | 0.4 | yes |
| multi-str | str | 4096 | churn | ordered | baseline | 6 | 97.9 | 97.1 | 0.99× [0.98, 0.99] | -1.2% | [-1.9%, -0.6%] | 0.6 pts | 0.5 | yes |
| multi-str | str | 4096 | build | ordered | baseline | 6 | 8.29 ms | 7.99 ms | 0.97× [0.96, 0.98] | -3.4% | [-4.3%, -2.4%] | 0.9 pts | 1.1 | yes |
| multi-str | str | 16384 | valuesFor | ordered | baseline | 8 | 78.3 | 76.9 | 0.99× [0.98, 1.00] | -1.4% | [-2.6%, -0.3%] | 1.3 pts | 0.9 | yes |
| multi-str | str | 16384 | valuesBetween | ordered | baseline | 8 | 4100 | 4053 | 0.99× [0.98, 0.99] | -1.4% | [-2.2%, -0.5%] | 1.0 pts | 0.9 | yes |
| multi-str | str | 16384 | prefix | ordered | baseline | 8 | 38.3 µs | 37.5 µs | 0.98× [0.97, 0.99] | -2.3% | [-3.5%, -1.0%] | 1.5 pts | 1.1 | yes |
| multi-str | str | 16384 | churn | ordered | baseline | 8 | 151 | 149 | 1.00× [0.98, 1.01] | -0.5% | [-2.0%, +1.1%] | 1.9 pts | 1.1 | yes |
| multi-str | str | 16384 | build | ordered | baseline | 8 | 46.18 ms | 45.23 ms | 0.98× [0.97, 0.99] | -2.2% | [-3.2%, -1.1%] | 1.2 pts | 0.6 | yes |
| multi-str | str | 262144 | valuesFor | ordered | baseline | 10 | 361 | 370 | 1.03× [1.02, 1.03] | +2.5% | [+1.9%, +3.2%] | 0.9 pts | 0.8 | yes |
| multi-str | str | 262144 | valuesBetween | ordered | baseline | 10 | 11.8 µs | 12.5 µs | 1.06× [1.05, 1.07] | +5.5% | [+4.7%, +6.2%] | 1.1 pts | 0.7 | yes |
| multi-str | str | 262144 | prefix | ordered | baseline | 10 | 1.87 ms | 1.97 ms | 1.05× [1.05, 1.06] | +5.1% | [+4.6%, +5.7%] | 0.7 pts | 0.9 | yes |
| multi-str | str | 262144 | churn | ordered | baseline | 10 | 592 | 562 | 0.97× [0.91, 1.03] | -3.6% | [-9.7%, +2.5%] | 8.5 pts | 1.6 | no |
| multi-str | str | 1048576 | valuesFor | ordered | baseline | 10 | 510 | 520 | 1.03× [1.02, 1.04] | +3.0% | [+2.0%, +4.0%] | 1.4 pts | 1.1 | yes |
| multi-str | str | 1048576 | valuesBetween | ordered | baseline | 10 | 13.1 µs | 14.0 µs | 1.08× [1.07, 1.09] | +7.4% | [+6.6%, +8.1%] | 1.1 pts | 1.0 | yes |
| multi-str | str | 1048576 | prefix | ordered | baseline | 10 | 7.95 ms | 8.59 ms | 1.08× [1.07, 1.08] | +7.2% | [+6.7%, +7.7%] | 0.7 pts | 0.7 | yes |
| multi-str | str | 1048576 | churn | ordered | baseline | 10 | 762 | 760 | 1.00× [0.96, 1.05] | +0.3% | [-3.8%, +4.4%] | 5.7 pts | 1.1 | no |
| multi-str | street | 4096 | valuesFor | ordered | baseline | 6 | 58.8 | 58.3 | 0.99× [0.98, 1.00] | -0.8% | [-1.6%, -0.0%] | 0.7 pts | 0.6 | yes |
| multi-str | street | 4096 | valuesBetween | ordered | baseline | 6 | 2651 | 2666 | 1.01× [1.00, 1.02] | +1.0% | [+0.1%, +1.9%] | 0.9 pts | 0.4 | yes |
| multi-str | street | 4096 | prefix | ordered | baseline | 6 | 292 | 287 | 0.99× [0.98, 1.00] | -1.2% | [-2.2%, -0.3%] | 0.9 pts | 0.3 | yes |
| multi-str | street | 4096 | churn | ordered | baseline | 6 | 108 | 107 | 0.98× [0.98, 0.99] | -1.9% | [-2.4%, -1.3%] | 0.5 pts | 0.3 | yes |
| multi-str | street | 4096 | build | ordered | baseline | 6 | 3.53 ms | 3.34 ms | 0.94× [0.94, 0.95] | -5.9% | [-6.4%, -5.4%] | 0.5 pts | 0.5 | yes |
| multi-str | street | 16384 | valuesFor | ordered | baseline | 6 | 81.2 | 79.9 | 0.99× [0.98, 0.99] | -1.4% | [-2.1%, -0.7%] | 0.7 pts | 0.5 | yes |
| multi-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 3139 | 3205 | 1.02× [1.01, 1.03] | +1.8% | [+0.9%, +2.8%] | 0.9 pts | 0.8 | yes |
| multi-str | street | 16384 | prefix | ordered | baseline | 6 | 1071 | 1090 | 1.02× [1.01, 1.02] | +1.7% | [+1.1%, +2.3%] | 0.6 pts | 0.2 | yes |
| multi-str | street | 16384 | churn | ordered | baseline | 6 | 160 | 158 | 0.99× [0.98, 1.00] | -1.3% | [-2.3%, -0.3%] | 0.9 pts | 0.6 | yes |
| multi-str | street | 16384 | build | ordered | baseline | 6 | 19.90 ms | 19.01 ms | 0.95× [0.95, 0.96] | -4.8% | [-5.2%, -4.4%] | 0.3 pts | 0.4 | yes |
| multi-str | u64 | 4096 | valuesFor | ordered | baseline | 6 | 41.8 | 40.0 | 0.96× [0.95, 0.97] | -4.5% | [-5.7%, -3.3%] | 1.1 pts | 0.8 | yes |
| multi-str | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 2875 | 2886 | 1.00× [1.00, 1.01] | +0.4% | [-0.4%, +1.3%] | 0.8 pts | 0.4 | yes |
| multi-str | u64 | 4096 | churn | ordered | baseline | 6 | 61.4 | 59.9 | 0.97× [0.97, 0.98] | -2.7% | [-3.3%, -2.1%] | 0.5 pts | 0.4 | yes |
| multi-str | u64 | 4096 | build | ordered | baseline | 6 | 5.65 ms | 5.46 ms | 0.96× [0.95, 0.98] | -3.9% | [-5.3%, -2.5%] | 1.3 pts | 1.2 | yes |
| multi-str | u64 | 16384 | valuesFor | ordered | baseline | 6 | 50.8 | 48.8 | 0.96× [0.95, 0.97] | -4.0% | [-4.8%, -3.2%] | 0.8 pts | 0.8 | yes |
| multi-str | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3784 | 3757 | 0.99× [0.98, 1.01] | -0.6% | [-2.0%, +0.7%] | 1.3 pts | 1.1 | yes |
| multi-str | u64 | 16384 | churn | ordered | baseline | 6 | 98.0 | 92.7 | 0.96× [0.95, 0.97] | -4.4% | [-5.5%, -3.3%] | 1.0 pts | 0.4 | yes |
| multi-str | u64 | 16384 | build | ordered | baseline | 6 | 32.53 ms | 31.30 ms | 0.97× [0.96, 0.97] | -3.6% | [-4.4%, -2.8%] | 0.8 pts | 0.4 | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | baseline | 10 | 186 | 184 | 0.98× [0.97, 0.99] | -2.0% | [-3.5%, -0.6%] | 2.0 pts | 2.1 | yes |
| multi-str | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 9689 | 9589 | 1.00× [0.99, 1.01] | -0.3% | [-1.4%, +0.7%] | 1.4 pts | 1.2 | yes |
| multi-str | u64 | 262144 | churn | ordered | baseline | 10 | 413 | 410 | 0.99× [0.95, 1.05] | -0.6% | [-5.6%, +4.5%] | 7.0 pts | 2.2 | no |
| multi-str | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 243 | 237 | 0.98× [0.97, 0.99] | -1.8% | [-2.9%, -0.6%] | 1.6 pts | 1.1 | yes |
| multi-str | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 9364 | 9488 | 1.01× [1.00, 1.02] | +0.8% | [+0.1%, +1.5%] | 1.0 pts | 0.9 | yes |
| multi-str | u64 | 1048576 | churn | ordered | baseline | 10 | 539 | 537 | 0.95× [0.89, 1.01] | -5.5% | [-12.3%, +1.3%] | 9.5 pts | 2.1 | no |
| multi-str | url | 4096 | valuesFor | ordered | baseline | 6 | 96.3 | 95.2 | 0.99× [0.99, 0.99] | -1.0% | [-1.3%, -0.7%] | 0.3 pts | 0.3 | yes |
| multi-str | url | 4096 | valuesBetween | ordered | baseline | 6 | 4692 | 4606 | 0.98× [0.97, 0.99] | -2.2% | [-3.1%, -1.3%] | 0.8 pts | 0.5 | yes |
| multi-str | url | 4096 | prefix | ordered | baseline | 6 | 204 | 199 | 0.97× [0.96, 0.98] | -3.0% | [-4.2%, -1.9%] | 1.1 pts | 0.8 | yes |
| multi-str | url | 4096 | churn | ordered | baseline | 6 | 151 | 149 | 0.99× [0.98, 1.00] | -0.9% | [-2.1%, +0.3%] | 1.1 pts | 0.7 | yes |
| multi-str | url | 4096 | build | ordered | baseline | 6 | 11.52 ms | 11.18 ms | 0.97× [0.96, 0.98] | -2.8% | [-3.7%, -1.9%] | 0.9 pts | 1.0 | yes |
| multi-str | url | 16384 | valuesFor | ordered | baseline | 10 | 126 | 126 | 0.99× [0.98, 1.01] | -0.6% | [-2.1%, +0.8%] | 2.1 pts | 1.3 | yes |
| multi-str | url | 16384 | valuesBetween | ordered | baseline | 10 | 5392 | 5317 | 0.99× [0.98, 1.00] | -1.1% | [-1.9%, -0.4%] | 1.1 pts | 0.8 | yes |
| multi-str | url | 16384 | prefix | ordered | baseline | 10 | 310 | 307 | 0.99× [0.97, 1.01] | -0.9% | [-2.6%, +0.7%] | 2.3 pts | 1.0 | yes |
| multi-str | url | 16384 | churn | ordered | baseline | 10 | 256 | 259 | 1.01× [1.00, 1.02] | +0.8% | [+0.1%, +1.5%] | 1.0 pts | 0.8 | yes |
| multi-str | url | 16384 | build | ordered | baseline | 10 | 70.66 ms | 70.17 ms | 1.00× [0.99, 1.01] | -0.3% | [-1.2%, +0.6%] | 1.2 pts | 1.2 | yes |
| multi-str | url | 262144 | valuesFor | ordered | baseline | 10 | 530 | 556 | 1.06× [1.04, 1.08] | +5.6% | [+4.3%, +7.0%] | 1.9 pts | 1.8 | yes |
| multi-str | url | 262144 | valuesBetween | ordered | baseline | 10 | 15.3 µs | 15.1 µs | 0.99× [0.98, 0.99] | -1.4% | [-1.8%, -0.9%] | 0.6 pts | 0.5 | yes |
| multi-str | url | 262144 | prefix | ordered | baseline | 10 | 2942 | 2976 | 1.01× [0.99, 1.04] | +1.4% | [-0.9%, +3.6%] | 3.1 pts | 0.5 | no |
| multi-str | url | 262144 | churn | ordered | baseline | 10 | 866 | 852 | 0.98× [0.95, 1.02] | -1.6% | [-5.1%, +2.0%] | 4.9 pts | 1.9 | no |
| multi-str | url | 1048576 | valuesFor | ordered | baseline | 10 | 784 | 812 | 1.05× [1.03, 1.08] | +4.9% | [+2.7%, +7.1%] | 3.1 pts | 2.7 | no |
| multi-str | url | 1048576 | valuesBetween | ordered | baseline | 10 | 16.9 µs | 16.8 µs | 1.00× [0.98, 1.01] | -0.3% | [-2.0%, +1.5%] | 2.4 pts | 1.8 | yes |
| multi-str | url | 1048576 | prefix | ordered | baseline | 10 | 12.8 µs | 13.4 µs | 1.03× [1.01, 1.05] | +3.0% | [+1.0%, +5.0%] | 2.8 pts | 0.3 | yes |
| multi-str | url | 1048576 | churn | ordered | baseline | 10 | 1177 | 1195 | 1.03× [0.99, 1.07] | +3.0% | [-1.0%, +7.0%] | 5.6 pts | 2.2 | no |
| multi-str | uuid | 4096 | valuesFor | ordered | baseline | 6 | 61.2 | 57.9 | 0.95× [0.94, 0.95] | -5.7% | [-6.3%, -5.2%] | 0.5 pts | 0.5 | yes |
| multi-str | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3690 | 3578 | 0.97× [0.97, 0.97] | -3.2% | [-3.6%, -2.9%] | 0.3 pts | 0.2 | yes |
| multi-str | uuid | 4096 | prefix | ordered | baseline | 6 | 99.4 | 92.9 | 0.94× [0.93, 0.94] | -6.8% | [-7.6%, -6.0%] | 0.8 pts | 0.9 | yes |
| multi-str | uuid | 4096 | churn | ordered | baseline | 6 | 96.4 | 89.3 | 0.92× [0.91, 0.94] | -8.2% | [-9.6%, -6.8%] | 1.3 pts | 0.9 | yes |
| multi-str | uuid | 4096 | build | ordered | baseline | 6 | 7.76 ms | 7.27 ms | 0.94× [0.93, 0.95] | -6.6% | [-7.6%, -5.5%] | 1.0 pts | 1.0 | yes |
| multi-str | uuid | 16384 | valuesFor | ordered | baseline | 10 | 72.4 | 67.9 | 0.96× [0.93, 0.98] | -4.6% | [-7.0%, -2.2%] | 3.3 pts | 2.4 | no |
| multi-str | uuid | 16384 | valuesBetween | ordered | baseline | 10 | 4156 | 3977 | 0.97× [0.95, 0.98] | -3.5% | [-5.2%, -1.9%] | 2.3 pts | 1.7 | yes |
| multi-str | uuid | 16384 | prefix | ordered | baseline | 10 | 125 | 116 | 0.94× [0.92, 0.96] | -6.3% | [-8.4%, -4.2%] | 2.9 pts | 2.4 | no |
| multi-str | uuid | 16384 | churn | ordered | baseline | 10 | 164 | 156 | 0.96× [0.94, 0.98] | -3.9% | [-5.8%, -2.0%] | 2.7 pts | 0.8 | yes |
| multi-str | uuid | 16384 | build | ordered | baseline | 10 | 46.42 ms | 44.49 ms | 0.97× [0.95, 0.98] | -3.5% | [-5.0%, -2.0%] | 2.1 pts | 1.1 | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | baseline | 10 | 337 | 335 | 0.98× [0.97, 0.99] | -2.0% | [-3.1%, -0.9%] | 1.6 pts | 1.2 | yes |
| multi-str | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 12.2 µs | 11.8 µs | 0.97× [0.96, 0.98] | -3.5% | [-4.5%, -2.4%] | 1.5 pts | 1.1 | yes |
| multi-str | uuid | 262144 | prefix | ordered | baseline | 10 | 812 | 785 | 0.96× [0.95, 0.97] | -4.2% | [-5.3%, -3.1%] | 1.5 pts | 0.9 | yes |
| multi-str | uuid | 262144 | churn | ordered | baseline | 10 | 577 | 554 | 0.97× [0.92, 1.02] | -3.3% | [-8.6%, +2.0%] | 7.4 pts | 1.7 | no |
| multi-str | uuid | 1048576 | valuesFor | ordered | baseline | 10 | 502 | 501 | 1.01× [0.99, 1.02] | +0.6% | [-0.6%, +1.9%] | 1.7 pts | 1.5 | yes |
| multi-str | uuid | 1048576 | valuesBetween | ordered | baseline | 10 | 14.0 µs | 13.6 µs | 0.97× [0.95, 0.99] | -3.1% | [-4.9%, -1.3%] | 2.6 pts | 2.2 | yes |
| multi-str | uuid | 1048576 | prefix | ordered | baseline | 10 | 2704 | 2618 | 0.96× [0.94, 0.97] | -4.4% | [-5.9%, -2.9%] | 2.1 pts | 1.6 | yes |
| multi-str | uuid | 1048576 | churn | ordered | baseline | 10 | 783 | 734 | 0.97× [0.92, 1.02] | -3.2% | [-8.4%, +2.0%] | 7.3 pts | 1.2 | no |
| unique-str | email | 4096 | valuesFor | ordered | baseline | 6 | 29.7 | 28.7 | 0.97× [0.95, 0.98] | -3.5% | [-5.4%, -1.7%] | 1.8 pts | 1.1 | yes |
| unique-str | email | 4096 | valuesBetween | ordered | baseline | 6 | 1520 | 1478 | 0.97× [0.96, 0.98] | -2.8% | [-3.9%, -1.7%] | 1.0 pts | 0.8 | yes |
| unique-str | email | 4096 | prefix | ordered | baseline | 6 | 62.6 | 58.2 | 0.93× [0.92, 0.94] | -7.5% | [-8.5%, -6.4%] | 1.0 pts | 1.1 | yes |
| unique-str | email | 4096 | churn | ordered | baseline | 6 | 79.1 | 75.5 | 0.96× [0.95, 0.96] | -4.4% | [-4.8%, -3.9%] | 0.4 pts | 0.4 | yes |
| unique-str | email | 4096 | build | ordered | baseline | 6 | 970.1 µs | 919.7 µs | 0.95× [0.93, 0.96] | -5.7% | [-7.3%, -4.2%] | 1.5 pts | 0.6 | yes |
| unique-str | email | 16384 | valuesFor | ordered | baseline | 6 | 43.4 | 40.5 | 0.93× [0.93, 0.94] | -7.4% | [-7.9%, -6.8%] | 0.5 pts | 0.6 | yes |
| unique-str | email | 16384 | valuesBetween | ordered | baseline | 6 | 1872 | 1796 | 0.96× [0.95, 0.97] | -4.2% | [-5.0%, -3.5%] | 0.7 pts | 0.7 | yes |
| unique-str | email | 16384 | prefix | ordered | baseline | 6 | 77.8 | 72.2 | 0.93× [0.92, 0.93] | -8.1% | [-9.2%, -7.0%] | 1.0 pts | 0.9 | yes |
| unique-str | email | 16384 | churn | ordered | baseline | 6 | 102 | 98.0 | 0.96× [0.96, 0.97] | -3.9% | [-4.7%, -3.1%] | 0.7 pts | 0.2 | yes |
| unique-str | email | 16384 | build | ordered | baseline | 6 | 4.84 ms | 4.56 ms | 0.94× [0.94, 0.95] | -6.1% | [-6.7%, -5.6%] | 0.5 pts | 0.4 | yes |
| unique-str | email | 262144 | valuesFor | ordered | baseline | 6 | 211 | 212 | 1.00× [1.00, 1.01] | +0.0% | [-0.5%, +0.6%] | 0.5 pts | 0.5 | yes |
| unique-str | email | 262144 | valuesBetween | ordered | baseline | 6 | 6108 | 6017 | 0.99× [0.98, 1.00] | -1.3% | [-2.5%, -0.2%] | 1.1 pts | 0.8 | yes |
| unique-str | email | 262144 | prefix | ordered | baseline | 6 | 290 | 284 | 0.98× [0.96, 0.99] | -2.5% | [-3.8%, -1.2%] | 1.3 pts | 1.1 | yes |
| unique-str | email | 262144 | churn | ordered | baseline | 6 | 342 | 340 | 0.98× [0.97, 1.00] | -1.8% | [-3.5%, -0.2%] | 1.5 pts | 0.9 | yes |
| unique-str | email | 1048576 | valuesFor | ordered | baseline | 8 | 360 | 360 | 1.00× [0.99, 1.01] | +0.3% | [-0.8%, +1.5%] | 1.3 pts | 1.4 | yes |
| unique-str | email | 1048576 | valuesBetween | ordered | baseline | 8 | 7844 | 7713 | 0.98× [0.98, 0.99] | -1.6% | [-2.5%, -0.7%] | 1.1 pts | 1.0 | yes |
| unique-str | email | 1048576 | prefix | ordered | baseline | 8 | 513 | 501 | 0.98× [0.98, 0.99] | -1.7% | [-2.3%, -1.0%] | 0.8 pts | 0.8 | yes |
| unique-str | email | 1048576 | churn | ordered | baseline | 8 | 547 | 541 | 0.99× [0.97, 1.01] | -0.9% | [-2.6%, +0.7%] | 2.0 pts | 1.1 | yes |
| unique-str | path | 4096 | valuesFor | ordered | baseline | 8 | 94.7 | 90.4 | 0.96× [0.95, 0.97] | -4.3% | [-5.4%, -3.3%] | 1.3 pts | 0.9 | yes |
| unique-str | path | 4096 | valuesBetween | ordered | baseline | 8 | 2799 | 2694 | 0.96× [0.94, 0.97] | -4.5% | [-6.4%, -2.7%] | 2.2 pts | 1.1 | yes |
| unique-str | path | 4096 | prefix | ordered | baseline | 8 | 284 | 274 | 0.97× [0.96, 0.98] | -3.5% | [-4.6%, -2.4%] | 1.4 pts | 0.6 | yes |
| unique-str | path | 4096 | churn | ordered | baseline | 8 | 215 | 209 | 0.97× [0.96, 0.97] | -3.5% | [-4.2%, -2.7%] | 0.9 pts | 0.6 | yes |
| unique-str | path | 4096 | build | ordered | baseline | 8 | 2.42 ms | 2.25 ms | 0.93× [0.92, 0.94] | -7.7% | [-8.6%, -6.8%] | 1.1 pts | 0.9 | yes |
| unique-str | path | 16384 | valuesFor | ordered | baseline | 8 | 124 | 122 | 0.98× [0.97, 0.99] | -2.2% | [-2.8%, -1.5%] | 0.8 pts | 0.9 | yes |
| unique-str | path | 16384 | valuesBetween | ordered | baseline | 8 | 3366 | 3206 | 0.95× [0.95, 0.96] | -4.9% | [-5.5%, -4.3%] | 0.7 pts | 0.6 | yes |
| unique-str | path | 16384 | prefix | ordered | baseline | 8 | 489 | 476 | 0.96× [0.95, 0.98] | -3.8% | [-5.6%, -2.1%] | 2.1 pts | 0.4 | yes |
| unique-str | path | 16384 | churn | ordered | baseline | 8 | 291 | 289 | 1.00× [0.99, 1.02] | +0.3% | [-1.5%, +2.0%] | 2.0 pts | 0.8 | yes |
| unique-str | path | 16384 | build | ordered | baseline | 8 | 12.28 ms | 11.95 ms | 0.98× [0.96, 0.99] | -2.6% | [-3.7%, -1.4%] | 1.4 pts | 0.8 | yes |
| unique-str | path | 262144 | valuesFor | ordered | baseline | 10 | 473 | 522 | 1.10× [1.09, 1.11] | +9.0% | [+8.5%, +9.5%] | 0.7 pts | 0.6 | yes |
| unique-str | path | 262144 | valuesBetween | ordered | baseline | 10 | 9759 | 9907 | 1.02× [1.01, 1.02] | +1.6% | [+0.9%, +2.3%] | 0.9 pts | 0.7 | yes |
| unique-str | path | 262144 | prefix | ordered | baseline | 10 | 8850 | 8667 | 0.99× [0.97, 1.01] | -1.2% | [-3.5%, +1.1%] | 3.2 pts | 0.3 | no |
| unique-str | path | 262144 | churn | ordered | baseline | 10 | 915 | 902 | 1.02× [0.99, 1.06] | +2.2% | [-1.1%, +5.5%] | 4.6 pts | 2.0 | no |
| unique-str | str | 4096 | valuesFor | ordered | baseline | 8 | 41.8 | 38.9 | 0.94× [0.92, 0.95] | -6.9% | [-8.5%, -5.3%] | 1.9 pts | 1.3 | yes |
| unique-str | str | 4096 | valuesBetween | ordered | baseline | 8 | 1980 | 1874 | 0.95× [0.93, 0.96] | -5.7% | [-7.2%, -4.1%] | 1.9 pts | 1.1 | yes |
| unique-str | str | 4096 | prefix | ordered | baseline | 8 | 4004 | 3762 | 0.95× [0.93, 0.96] | -5.7% | [-7.4%, -4.1%] | 2.0 pts | 0.6 | yes |
| unique-str | str | 4096 | churn | ordered | baseline | 8 | 101 | 94.7 | 0.94× [0.93, 0.94] | -6.8% | [-7.6%, -6.1%] | 0.9 pts | 0.8 | yes |
| unique-str | str | 4096 | build | ordered | baseline | 8 | 1.24 ms | 1.13 ms | 0.92× [0.90, 0.93] | -9.3% | [-10.7%, -7.9%] | 1.7 pts | 0.5 | yes |
| unique-str | str | 16384 | valuesFor | ordered | baseline | 6 | 52.3 | 49.9 | 0.95× [0.95, 0.96] | -4.7% | [-5.6%, -3.9%] | 0.9 pts | 0.7 | yes |
| unique-str | str | 16384 | valuesBetween | ordered | baseline | 6 | 1960 | 1873 | 0.96× [0.95, 0.97] | -4.4% | [-5.4%, -3.3%] | 1.0 pts | 1.0 | yes |
| unique-str | str | 16384 | prefix | ordered | baseline | 6 | 17.1 µs | 15.9 µs | 0.93× [0.93, 0.94] | -7.1% | [-7.9%, -6.2%] | 0.8 pts | 0.5 | yes |
| unique-str | str | 16384 | churn | ordered | baseline | 6 | 120 | 117 | 0.97× [0.95, 0.99] | -3.3% | [-5.1%, -1.4%] | 1.8 pts | 0.8 | yes |
| unique-str | str | 16384 | build | ordered | baseline | 6 | 5.80 ms | 5.50 ms | 0.94× [0.94, 0.95] | -6.1% | [-6.9%, -5.2%] | 0.8 pts | 0.7 | yes |
| unique-str | str | 262144 | valuesFor | ordered | baseline | 6 | 239 | 238 | 1.00× [0.99, 1.02] | +0.1% | [-1.4%, +1.6%] | 1.4 pts | 1.4 | yes |
| unique-str | str | 262144 | valuesBetween | ordered | baseline | 6 | 6253 | 6357 | 1.01× [1.00, 1.03] | +1.3% | [+0.1%, +2.5%] | 1.2 pts | 1.4 | yes |
| unique-str | str | 262144 | prefix | ordered | baseline | 6 | 1.02 ms | 1.06 ms | 1.03× [1.02, 1.04] | +2.8% | [+2.0%, +3.5%] | 0.7 pts | 0.2 | yes |
| unique-str | str | 262144 | churn | ordered | baseline | 6 | 351 | 352 | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.4%] | 0.7 pts | 0.3 | yes |
| unique-str | str | 1048576 | valuesFor | ordered | baseline | 10 | 386 | 390 | 1.02× [1.01, 1.02] | +1.6% | [+0.9%, +2.4%] | 1.0 pts | 1.1 | yes |
| unique-str | str | 1048576 | valuesBetween | ordered | baseline | 10 | 7292 | 7821 | 1.08× [1.07, 1.08] | +7.0% | [+6.5%, +7.4%] | 0.6 pts | 0.7 | yes |
| unique-str | str | 1048576 | prefix | ordered | baseline | 10 | 4.42 ms | 4.73 ms | 1.06× [1.06, 1.07] | +6.1% | [+5.6%, +6.5%] | 0.6 pts | 1.2 | yes |
| unique-str | str | 1048576 | churn | ordered | baseline | 10 | 569 | 577 | 1.01× [0.98, 1.05] | +1.4% | [-1.6%, +4.3%] | 4.2 pts | 3.3 | no |
| unique-str | street | 4096 | valuesFor | ordered | baseline | 8 | 52.9 | 52.0 | 0.98× [0.96, 1.00] | -1.9% | [-3.7%, -0.1%] | 2.1 pts | 1.2 | yes |
| unique-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 2158 | 2187 | 1.01× [1.00, 1.01] | +0.8% | [+0.3%, +1.4%] | 0.6 pts | 0.4 | yes |
| unique-str | street | 4096 | prefix | ordered | baseline | 8 | 248 | 246 | 0.99× [0.97, 1.01] | -1.2% | [-3.2%, +0.8%] | 2.4 pts | 0.9 | yes |
| unique-str | street | 4096 | churn | ordered | baseline | 8 | 131 | 123 | 0.94× [0.94, 0.95] | -5.9% | [-6.6%, -5.2%] | 0.8 pts | 0.5 | yes |
| unique-str | street | 4096 | build | ordered | baseline | 8 | 1.56 ms | 1.40 ms | 0.90× [0.90, 0.91] | -10.7% | [-11.2%, -10.2%] | 0.6 pts | 0.5 | yes |
| unique-str | street | 16384 | valuesFor | ordered | baseline | 6 | 74.8 | 73.5 | 0.98× [0.98, 0.99] | -1.7% | [-2.2%, -1.3%] | 0.4 pts | 0.4 | yes |
| unique-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 2513 | 2579 | 1.03× [1.02, 1.04] | +2.8% | [+2.3%, +3.4%] | 0.5 pts | 0.6 | yes |
| unique-str | street | 16384 | prefix | ordered | baseline | 6 | 869 | 882 | 1.01× [1.00, 1.02] | +1.1% | [-0.2%, +2.3%] | 1.2 pts | 0.5 | yes |
| unique-str | street | 16384 | churn | ordered | baseline | 6 | 167 | 159 | 0.96× [0.94, 0.97] | -4.6% | [-6.1%, -3.1%] | 1.4 pts | 0.8 | yes |
| unique-str | street | 16384 | build | ordered | baseline | 6 | 7.72 ms | 7.11 ms | 0.92× [0.91, 0.93] | -8.5% | [-9.9%, -7.1%] | 1.3 pts | 1.6 | yes |
| unique-str | u64 | 4096 | valuesFor | ordered | baseline | 10 | 18.6 | 16.8 | 0.91× [0.88, 0.93] | -10.4% | [-13.9%, -7.0%] | 4.8 pts | 1.9 | no |
| unique-str | u64 | 4096 | valuesBetween | ordered | baseline | 10 | 1163 | 1160 | 1.00× [0.99, 1.00] | -0.1% | [-0.6%, +0.5%] | 0.8 pts | 0.8 | yes |
| unique-str | u64 | 4096 | churn | ordered | baseline | 10 | 53.8 | 50.5 | 0.94× [0.94, 0.95] | -6.3% | [-6.9%, -5.8%] | 0.7 pts | 0.7 | yes |
| unique-str | u64 | 4096 | build | ordered | baseline | 10 | 739.2 µs | 697.7 µs | 0.94× [0.94, 0.95] | -5.9% | [-6.5%, -5.3%] | 0.9 pts | 0.5 | yes |
| unique-str | u64 | 16384 | valuesFor | ordered | baseline | 8 | 26.8 | 25.0 | 0.93× [0.92, 0.94] | -7.4% | [-8.4%, -6.3%] | 1.3 pts | 1.4 | yes |
| unique-str | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1731 | 1736 | 1.00× [0.99, 1.01] | +0.1% | [-0.8%, +1.1%] | 1.1 pts | 1.1 | yes |
| unique-str | u64 | 16384 | churn | ordered | baseline | 8 | 59.6 | 56.2 | 0.94× [0.92, 0.95] | -6.7% | [-8.5%, -4.9%] | 2.1 pts | 0.8 | yes |
| unique-str | u64 | 16384 | build | ordered | baseline | 8 | 3.49 ms | 3.29 ms | 0.94× [0.93, 0.95] | -6.5% | [-7.5%, -5.4%] | 1.2 pts | 1.1 | yes |
| unique-str | u64 | 262144 | valuesFor | ordered | baseline | 6 | 108 | 105 | 0.98× [0.96, 0.99] | -2.5% | [-4.2%, -0.7%] | 1.7 pts | 1.3 | yes |
| unique-str | u64 | 262144 | valuesBetween | ordered | baseline | 6 | 4523 | 4484 | 0.99× [0.98, 1.00] | -1.3% | [-2.5%, -0.1%] | 1.1 pts | 1.1 | yes |
| unique-str | u64 | 262144 | churn | ordered | baseline | 6 | 218 | 217 | 0.98× [0.97, 1.00] | -1.7% | [-3.6%, +0.3%] | 1.9 pts | 0.8 | yes |
| unique-str | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 157 | 153 | 0.98× [0.97, 0.98] | -2.2% | [-2.7%, -1.6%] | 0.8 pts | 1.0 | yes |
| unique-str | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 3925 | 3912 | 0.99× [0.99, 1.00] | -0.5% | [-1.2%, +0.2%] | 1.0 pts | 1.2 | yes |
| unique-str | u64 | 1048576 | churn | ordered | baseline | 10 | 362 | 360 | 0.98× [0.96, 1.01] | -1.7% | [-4.4%, +1.0%] | 3.7 pts | 2.2 | no |
| unique-str | url | 4096 | valuesFor | ordered | baseline | 6 | 74.0 | 72.6 | 0.98× [0.97, 0.99] | -1.9% | [-2.7%, -1.1%] | 0.8 pts | 0.6 | yes |
| unique-str | url | 4096 | valuesBetween | ordered | baseline | 6 | 2682 | 2546 | 0.95× [0.93, 0.96] | -5.5% | [-7.1%, -3.9%] | 1.5 pts | 0.7 | yes |
| unique-str | url | 4096 | prefix | ordered | baseline | 6 | 169 | 164 | 0.97× [0.96, 0.98] | -3.1% | [-4.1%, -2.0%] | 1.0 pts | 0.7 | yes |
| unique-str | url | 4096 | churn | ordered | baseline | 6 | 174 | 170 | 0.97× [0.96, 0.98] | -3.0% | [-4.2%, -1.9%] | 1.1 pts | 0.5 | yes |
| unique-str | url | 4096 | build | ordered | baseline | 6 | 2.00 ms | 1.86 ms | 0.93× [0.92, 0.94] | -7.2% | [-8.4%, -5.9%] | 1.2 pts | 0.7 | yes |
| unique-str | url | 16384 | valuesFor | ordered | baseline | 8 | 99.4 | 97.9 | 0.99× [0.98, 0.99] | -1.4% | [-1.8%, -1.0%] | 0.5 pts | 0.5 | yes |
| unique-str | url | 16384 | valuesBetween | ordered | baseline | 8 | 3254 | 3082 | 0.95× [0.94, 0.96] | -5.4% | [-6.1%, -4.7%] | 0.8 pts | 0.7 | yes |
| unique-str | url | 16384 | prefix | ordered | baseline | 8 | 244 | 240 | 0.99× [0.97, 1.00] | -1.5% | [-3.3%, +0.3%] | 2.2 pts | 1.0 | yes |
| unique-str | url | 16384 | churn | ordered | baseline | 8 | 232 | 230 | 0.98× [0.97, 0.99] | -1.6% | [-2.6%, -0.5%] | 1.3 pts | 0.5 | yes |
| unique-str | url | 16384 | build | ordered | baseline | 8 | 10.11 ms | 9.68 ms | 0.95× [0.95, 0.96] | -5.1% | [-5.6%, -4.6%] | 0.6 pts | 0.5 | yes |
| unique-str | url | 262144 | valuesFor | ordered | baseline | 10 | 436 | 465 | 1.08× [1.06, 1.11] | +7.7% | [+5.9%, +9.6%] | 2.6 pts | 2.4 | yes |
| unique-str | url | 262144 | valuesBetween | ordered | baseline | 10 | 9635 | 9532 | 0.99× [0.98, 1.01] | -0.7% | [-2.0%, +0.5%] | 1.7 pts | 1.7 | yes |
| unique-str | url | 262144 | prefix | ordered | baseline | 10 | 2106 | 2137 | 1.02× [0.99, 1.04] | +1.6% | [-0.9%, +4.0%] | 3.5 pts | 0.7 | no |
| unique-str | url | 262144 | churn | ordered | baseline | 10 | 629 | 660 | 1.05× [1.04, 1.06] | +4.5% | [+3.6%, +5.5%] | 1.3 pts | 1.2 | yes |
| unique-str | url | 1048576 | valuesFor | ordered | baseline | 10 | 682 | 737 | 1.08× [1.07, 1.09] | +7.4% | [+6.1%, +8.6%] | 1.7 pts | 1.6 | yes |
| unique-str | url | 1048576 | valuesBetween | ordered | baseline | 10 | 11.1 µs | 11.1 µs | 1.00× [0.99, 1.00] | -0.2% | [-0.8%, +0.3%] | 0.8 pts | 0.9 | yes |
| unique-str | url | 1048576 | prefix | ordered | baseline | 10 | 7278 | 7289 | 1.01× [0.99, 1.03] | +0.6% | [-1.5%, +2.7%] | 3.0 pts | 0.3 | no |
| unique-str | url | 1048576 | churn | ordered | baseline | 10 | 995 | 1007 | 1.02× [1.00, 1.03] | +1.8% | [+0.3%, +3.4%] | 2.2 pts | 1.5 | yes |
| unique-str | uuid | 4096 | valuesFor | ordered | baseline | 8 | 36.1 | 33.3 | 0.92× [0.91, 0.94] | -8.4% | [-10.2%, -6.5%] | 2.2 pts | 1.6 | yes |
| unique-str | uuid | 4096 | valuesBetween | ordered | baseline | 8 | 1792 | 1685 | 0.94× [0.93, 0.95] | -6.6% | [-7.4%, -5.8%] | 1.0 pts | 0.6 | yes |
| unique-str | uuid | 4096 | prefix | ordered | baseline | 8 | 77.4 | 70.6 | 0.91× [0.91, 0.91] | -9.7% | [-10.0%, -9.3%] | 0.4 pts | 0.4 | yes |
| unique-str | uuid | 4096 | churn | ordered | baseline | 8 | 93.3 | 86.4 | 0.93× [0.92, 0.93] | -7.7% | [-8.5%, -7.0%] | 0.9 pts | 0.6 | yes |
| unique-str | uuid | 4096 | build | ordered | baseline | 8 | 1.10 ms | 1.02 ms | 0.92× [0.91, 0.93] | -9.2% | [-10.3%, -8.1%] | 1.3 pts | 0.6 | yes |
| unique-str | uuid | 16384 | valuesFor | ordered | baseline | 8 | 47.0 | 41.5 | 0.89× [0.88, 0.89] | -12.6% | [-13.4%, -11.8%] | 0.9 pts | 0.9 | yes |
| unique-str | uuid | 16384 | valuesBetween | ordered | baseline | 8 | 2023 | 1797 | 0.89× [0.89, 0.89] | -12.3% | [-12.9%, -11.8%] | 0.7 pts | 0.6 | yes |
| unique-str | uuid | 16384 | prefix | ordered | baseline | 8 | 93.1 | 84.7 | 0.91× [0.90, 0.91] | -10.2% | [-11.0%, -9.4%] | 1.0 pts | 1.0 | yes |
| unique-str | uuid | 16384 | churn | ordered | baseline | 8 | 111 | 105 | 0.96× [0.94, 0.97] | -4.6% | [-6.5%, -2.8%] | 2.2 pts | 1.0 | yes |
| unique-str | uuid | 16384 | build | ordered | baseline | 8 | 5.36 ms | 4.90 ms | 0.92× [0.91, 0.92] | -9.3% | [-9.8%, -8.7%] | 0.7 pts | 0.6 | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | baseline | 8 | 227 | 221 | 0.98× [0.97, 0.99] | -2.3% | [-3.3%, -1.2%] | 1.3 pts | 1.0 | yes |
| unique-str | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 6427 | 6201 | 0.97× [0.95, 0.99] | -3.4% | [-5.4%, -1.5%] | 2.3 pts | 2.0 | yes |
| unique-str | uuid | 262144 | prefix | ordered | baseline | 8 | 468 | 453 | 0.96× [0.95, 0.97] | -4.0% | [-5.3%, -2.8%] | 1.5 pts | 1.4 | yes |
| unique-str | uuid | 262144 | churn | ordered | baseline | 8 | 333 | 331 | 1.00× [0.98, 1.02] | +0.1% | [-1.6%, +1.8%] | 2.0 pts | 1.0 | yes |
| unique-str | uuid | 1048576 | valuesFor | ordered | baseline | 8 | 380 | 382 | 1.01× [0.99, 1.02] | +0.7% | [-0.9%, +2.3%] | 1.9 pts | 2.0 | yes |
| unique-str | uuid | 1048576 | valuesBetween | ordered | baseline | 8 | 8003 | 7406 | 0.93× [0.92, 0.94] | -7.8% | [-8.9%, -6.8%] | 1.3 pts | 1.3 | yes |
| unique-str | uuid | 1048576 | prefix | ordered | baseline | 8 | 1521 | 1417 | 0.93× [0.93, 0.93] | -7.5% | [-8.0%, -7.0%] | 0.6 pts | 0.5 | yes |
| unique-str | uuid | 1048576 | churn | ordered | baseline | 8 | 561 | 554 | 0.99× [0.97, 1.01] | -0.8% | [-2.7%, +1.1%] | 2.3 pts | 1.5 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi-str email n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.22% does not clear the 1.63% median noise floor of the processes
- multi-str email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.52%, 0.09%] includes zero
- multi-str email n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.96% does not clear the 1.67% median noise floor of the processes
- multi-str email n=262144 valuesBetween: ordered vs baseline: the pooled difference of -1.01% does not clear the 1.56% median noise floor of the processes
- multi-str email n=262144 churn: ordered vs baseline: the pooled difference of -2.01% does not clear the 2.85% median noise floor of the processes
- multi-str email n=262144 churn: ordered vs baseline: the pooled interval [-5.54%, 1.52%] includes zero
- multi-str email n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.71% does not clear the 1.48% median noise floor of the processes
- multi-str email n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.11% does not clear the 1.08% median noise floor of the processes
- multi-str email n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.93%, 1.14%] includes zero
- multi-str email n=1048576 prefix: ordered vs baseline: the pooled difference of -1.41% does not clear the 1.94% median noise floor of the processes
- multi-str path n=4096 prefix: ordered vs baseline: the pooled difference of -1.60% does not clear the 2.57% median noise floor of the processes
- multi-str path n=4096 prefix: ordered vs baseline: the pooled interval [-3.39%, 0.20%] includes zero
- multi-str path n=4096 churn: ordered vs baseline: the pooled difference of -0.67% does not clear the 1.61% median noise floor of the processes
- multi-str path n=16384 valuesFor: ordered vs baseline: the pooled difference of -1.00% does not clear the 1.77% median noise floor of the processes
- multi-str path n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.02%, 0.02%] includes zero
- multi-str path n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.44% does not clear the 1.04% median noise floor of the processes
- multi-str path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.01%, 0.13%] includes zero
- multi-str path n=16384 prefix: ordered vs baseline: the pooled difference of 1.69% does not clear the 4.40% median noise floor of the processes
- multi-str path n=16384 prefix: ordered vs baseline: the pooled interval [-0.69%, 4.07%] includes zero
- multi-str path n=16384 churn: ordered vs baseline: the pooled difference of 0.38% does not clear the 2.03% median noise floor of the processes
- multi-str path n=16384 churn: ordered vs baseline: the pooled interval [-1.02%, 1.77%] includes zero
- multi-str path n=16384 build: ordered vs baseline: the pooled difference of 0.20% does not clear the 1.48% median noise floor of the processes
- multi-str path n=16384 build: ordered vs baseline: the pooled interval [-0.66%, 1.06%] includes zero
- multi-str path n=262144 prefix: ordered vs baseline: the pooled difference of 2.25% does not clear the 9.52% median noise floor of the processes
- multi-str path n=262144 churn: ordered vs baseline: the pooled interval [-5.25%, 0.66%] includes zero
- multi-str path n=262144 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=4096 churn: ordered vs baseline: the pooled difference of -1.24% does not clear the 2.18% median noise floor of the processes
- multi-str str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.38% does not clear the 2.20% median noise floor of the processes
- multi-str str n=16384 churn: ordered vs baseline: the pooled difference of -0.46% does not clear the 2.06% median noise floor of the processes
- multi-str str n=16384 churn: ordered vs baseline: the pooled interval [-2.01%, 1.09%] includes zero
- multi-str str n=262144 churn: ordered vs baseline: the pooled interval [-9.72%, 2.50%] includes zero
- multi-str str n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str str n=1048576 churn: ordered vs baseline: the pooled difference of 0.30% does not clear the 3.14% median noise floor of the processes
- multi-str str n=1048576 churn: ordered vs baseline: the pooled interval [-3.81%, 4.41%] includes zero
- multi-str street n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.82% does not clear the 1.19% median noise floor of the processes
- multi-str street n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.01% does not clear the 2.16% median noise floor of the processes
- multi-str street n=4096 prefix: ordered vs baseline: the pooled difference of -1.23% does not clear the 3.02% median noise floor of the processes
- multi-str street n=16384 valuesFor: ordered vs baseline: the pooled difference of -1.42% does not clear the 2.17% median noise floor of the processes
- multi-str street n=16384 prefix: ordered vs baseline: the pooled difference of 1.68% does not clear the 3.73% median noise floor of the processes
- multi-str street n=16384 churn: ordered vs baseline: the pooled difference of -1.33% does not clear the 1.86% median noise floor of the processes
- multi-str u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.43% does not clear the 1.66% median noise floor of the processes
- multi-str u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.41%, 1.27%] includes zero
- multi-str u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.64% does not clear the 1.02% median noise floor of the processes
- multi-str u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.99%, 0.72%] includes zero
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.32% does not clear the 1.61% median noise floor of the processes
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.36%, 0.71%] includes zero
- multi-str u64 n=262144 churn: ordered vs baseline: the pooled difference of -0.57% does not clear the 2.59% median noise floor of the processes
- multi-str u64 n=262144 churn: ordered vs baseline: the pooled interval [-5.61%, 4.46%] includes zero
- multi-str u64 n=262144 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 churn: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str u64 n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.78% does not clear the 1.44% median noise floor of the processes
- multi-str u64 n=1048576 churn: ordered vs baseline: the pooled interval [-12.29%, 1.27%] includes zero
- multi-str u64 n=1048576 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=1048576 churn: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str url n=4096 churn: ordered vs baseline: the pooled difference of -0.93% does not clear the 1.42% median noise floor of the processes
- multi-str url n=4096 churn: ordered vs baseline: the pooled interval [-2.12%, 0.25%] includes zero
- multi-str url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.62% does not clear the 1.99% median noise floor of the processes
- multi-str url n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.09%, 0.84%] includes zero
- multi-str url n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.14% does not clear the 1.30% median noise floor of the processes
- multi-str url n=16384 prefix: ordered vs baseline: the pooled difference of -0.94% does not clear the 2.32% median noise floor of the processes
- multi-str url n=16384 prefix: ordered vs baseline: the pooled interval [-2.59%, 0.70%] includes zero
- multi-str url n=16384 churn: ordered vs baseline: the pooled difference of 0.81% does not clear the 1.76% median noise floor of the processes
- multi-str url n=16384 build: ordered vs baseline: the pooled difference of -0.30% does not clear the 1.33% median noise floor of the processes
- multi-str url n=16384 build: ordered vs baseline: the pooled interval [-1.18%, 0.58%] includes zero
- multi-str url n=262144 valuesBetween: ordered vs baseline: the pooled difference of -1.37% does not clear the 1.57% median noise floor of the processes
- multi-str url n=262144 prefix: ordered vs baseline: the pooled difference of 1.37% does not clear the 9.21% median noise floor of the processes
- multi-str url n=262144 prefix: ordered vs baseline: the pooled interval [-0.87%, 3.60%] includes zero
- multi-str url n=262144 churn: ordered vs baseline: the pooled interval [-5.08%, 1.97%] includes zero
- multi-str url n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.27% does not clear the 1.18% median noise floor of the processes
- multi-str url n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-2.00%, 1.46%] includes zero
- multi-str url n=1048576 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str url n=1048576 prefix: ordered vs baseline: the pooled difference of 3.02% does not clear the 12.09% median noise floor of the processes
- multi-str url n=1048576 churn: ordered vs baseline: the pooled interval [-1.04%, 6.96%] includes zero
- multi-str url n=1048576 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=16384 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=16384 prefix: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 churn: ordered vs baseline: the pooled interval [-8.63%, 1.95%] includes zero
- multi-str uuid n=262144 churn: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str uuid n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.61% does not clear the 1.59% median noise floor of the processes
- multi-str uuid n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.63%, 1.85%] includes zero
- multi-str uuid n=1048576 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str uuid n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 churn: ordered vs baseline: the pooled difference of -3.18% does not clear the 3.85% median noise floor of the processes
- multi-str uuid n=1048576 churn: ordered vs baseline: the pooled interval [-8.37%, 2.01%] includes zero
- multi-str uuid n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str email n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.05% does not clear the 1.49% median noise floor of the processes
- unique-str email n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.45%, 0.55%] includes zero
- unique-str email n=262144 churn: ordered vs baseline: the pooled difference of -1.84% does not clear the 2.48% median noise floor of the processes
- unique-str email n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.35% does not clear the 0.95% median noise floor of the processes
- unique-str email n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.77%, 1.47%] includes zero
- unique-str email n=1048576 churn: ordered vs baseline: the pooled difference of -0.93% does not clear the 1.45% median noise floor of the processes
- unique-str email n=1048576 churn: ordered vs baseline: the pooled interval [-2.61%, 0.74%] includes zero
- unique-str path n=16384 prefix: ordered vs baseline: the pooled difference of -3.83% does not clear the 5.52% median noise floor of the processes
- unique-str path n=16384 churn: ordered vs baseline: the pooled difference of 0.26% does not clear the 3.00% median noise floor of the processes
- unique-str path n=16384 churn: ordered vs baseline: the pooled interval [-1.45%, 1.97%] includes zero
- unique-str path n=262144 valuesBetween: ordered vs baseline: the pooled difference of 1.62% does not clear the 1.65% median noise floor of the processes
- unique-str path n=262144 prefix: ordered vs baseline: the pooled difference of -1.20% does not clear the 14.56% median noise floor of the processes
- unique-str path n=262144 prefix: ordered vs baseline: the pooled interval [-3.47%, 1.08%] includes zero
- unique-str path n=262144 churn: ordered vs baseline: the pooled interval [-1.12%, 5.48%] includes zero
- unique-str str n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.07% does not clear the 0.83% median noise floor of the processes
- unique-str str n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.42%, 1.57%] includes zero
- unique-str str n=262144 valuesBetween: ordered vs baseline: the pooled difference of 1.29% does not clear the 1.61% median noise floor of the processes
- unique-str str n=262144 churn: ordered vs baseline: the pooled difference of -0.37% does not clear the 2.08% median noise floor of the processes
- unique-str str n=262144 churn: ordered vs baseline: the pooled interval [-1.14%, 0.40%] includes zero
- unique-str str n=1048576 churn: ordered vs baseline: the pooled difference of 1.37% does not clear the 1.88% median noise floor of the processes
- unique-str str n=1048576 churn: ordered vs baseline: the pooled interval [-1.61%, 4.34%] includes zero
- unique-str str n=1048576 churn: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=1048576 churn: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str street n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.82% does not clear the 2.07% median noise floor of the processes
- unique-str street n=4096 prefix: ordered vs baseline: the pooled difference of -1.19% does not clear the 2.84% median noise floor of the processes
- unique-str street n=4096 prefix: ordered vs baseline: the pooled interval [-3.16%, 0.79%] includes zero
- unique-str street n=16384 prefix: ordered vs baseline: the pooled difference of 1.06% does not clear the 3.74% median noise floor of the processes
- unique-str street n=16384 prefix: ordered vs baseline: the pooled interval [-0.15%, 2.28%] includes zero
- unique-str u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.07% does not clear the 1.48% median noise floor of the processes
- unique-str u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.62%, 0.48%] includes zero
- unique-str u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.12% does not clear the 1.43% median noise floor of the processes
- unique-str u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.83%, 1.07%] includes zero
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of -1.32% does not clear the 1.39% median noise floor of the processes
- unique-str u64 n=262144 churn: ordered vs baseline: the pooled difference of -1.66% does not clear the 2.65% median noise floor of the processes
- unique-str u64 n=262144 churn: ordered vs baseline: the pooled interval [-3.63%, 0.30%] includes zero
- unique-str u64 n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.51% does not clear the 1.65% median noise floor of the processes
- unique-str u64 n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-1.21%, 0.18%] includes zero
- unique-str u64 n=1048576 churn: ordered vs baseline: the pooled difference of -1.69% does not clear the 1.75% median noise floor of the processes
- unique-str u64 n=1048576 churn: ordered vs baseline: the pooled interval [-4.36%, 0.97%] includes zero
- unique-str u64 n=1048576 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str url n=16384 prefix: ordered vs baseline: the pooled difference of -1.48% does not clear the 2.50% median noise floor of the processes
- unique-str url n=16384 prefix: ordered vs baseline: the pooled interval [-3.29%, 0.33%] includes zero
- unique-str url n=16384 churn: ordered vs baseline: the pooled difference of -1.57% does not clear the 2.31% median noise floor of the processes
- unique-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.71% does not clear the 1.18% median noise floor of the processes
- unique-str url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.95%, 0.54%] includes zero
- unique-str url n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str url n=262144 prefix: ordered vs baseline: the pooled difference of 1.55% does not clear the 5.82% median noise floor of the processes
- unique-str url n=262144 prefix: ordered vs baseline: the pooled interval [-0.94%, 4.05%] includes zero
- unique-str url n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.24% does not clear the 1.24% median noise floor of the processes
- unique-str url n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.81%, 0.32%] includes zero
- unique-str url n=1048576 prefix: ordered vs baseline: the pooled difference of 0.61% does not clear the 10.11% median noise floor of the processes
- unique-str url n=1048576 prefix: ordered vs baseline: the pooled interval [-1.50%, 2.72%] includes zero
- unique-str url n=1048576 churn: ordered vs baseline: the pooled difference of 1.81% does not clear the 2.07% median noise floor of the processes
- unique-str url n=1048576 churn: ordered vs baseline: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 churn: ordered vs baseline: the pooled difference of 0.09% does not clear the 2.87% median noise floor of the processes
- unique-str uuid n=262144 churn: ordered vs baseline: the pooled interval [-1.58%, 1.75%] includes zero
- unique-str uuid n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.69% does not clear the 0.71% median noise floor of the processes
- unique-str uuid n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.92%, 2.30%] includes zero
- unique-str uuid n=1048576 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str uuid n=1048576 churn: ordered vs baseline: the pooled difference of -0.83% does not clear the 1.94% median noise floor of the processes
- unique-str uuid n=1048576 churn: ordered vs baseline: the pooled interval [-2.73%, 1.07%] includes zero
