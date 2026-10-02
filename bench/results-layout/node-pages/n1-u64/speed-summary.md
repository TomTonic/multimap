| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 12 | 46.6 | 46.5 | 0.98× [0.92, 1.05] | -1.8% | [-8.6%, +5.0%] | 6.8 pts | 2.2 | no | no |
| multi | email | 4096 | valuesBetween | ordered | baseline | 12 | 2851 | 2829 | 0.98× [0.97, 1.00] | -1.6% | [-3.4%, +0.3%] | 2.7 pts | 1.0 | yes | no |
| multi | email | 4096 | prefix | ordered | baseline | 12 | 77.9 | 76.7 | 0.97× [0.92, 1.02] | -3.3% | [-8.2%, +1.6%] | 4.7 pts | 4.2 | no | no |
| multi | email | 4096 | churn | ordered | baseline | 12 | 69.7 | 67.1 | 0.96× [0.95, 0.97] | -4.2% | [-5.2%, -3.2%] | 1.0 pts | 1.6 | yes | yes |
| multi | email | 4096 | build | ordered | baseline | 12 | 7.34 ms | 6.77 ms | 0.92× [0.92, 0.93] | -8.2% | [-8.7%, -7.7%] | 0.7 pts | 1.0 | yes | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 60.7 | 60.3 | 0.99× [0.99, 1.00] | -0.7% | [-1.4%, +0.0%] | 0.7 pts | 1.3 | yes | no |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3521 | 3484 | 0.99× [0.98, 1.00] | -1.1% | [-1.9%, -0.3%] | 0.8 pts | 1.1 | yes | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 91.9 | 91.6 | 1.00× [0.99, 1.02] | +0.1% | [-1.3%, +1.5%] | 1.4 pts | 2.3 | yes | no |
| multi | email | 16384 | churn | ordered | baseline | 6 | 96.4 | 92.7 | 0.96× [0.96, 0.96] | -4.2% | [-4.7%, -3.7%] | 0.5 pts | 0.9 | yes | yes |
| multi | email | 16384 | build | ordered | baseline | 6 | 36.66 ms | 33.83 ms | 0.92× [0.91, 0.92] | -9.0% | [-9.6%, -8.3%] | 0.6 pts | 1.0 | yes | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 12 | 203 | 202 | 0.98× [0.97, 1.00] | -1.6% | [-3.6%, +0.4%] | 2.7 pts | 4.0 | no | no |
| multi | email | 262144 | valuesBetween | ordered | baseline | 12 | 7344 | 7366 | 1.00× [0.99, 1.01] | -0.2% | [-1.4%, +1.0%] | 1.6 pts | 2.7 | yes | no |
| multi | email | 262144 | prefix | ordered | baseline | 12 | 268 | 267 | 1.00× [0.98, 1.02] | -0.0% | [-1.6%, +1.6%] | 2.2 pts | 3.5 | yes | no |
| multi | email | 262144 | churn | ordered | baseline | 12 | 403 | 385 | 0.92× [0.88, 0.95] | -9.0% | [-13.0%, -5.0%] | 4.7 pts | 0.7 | no | yes |
| multi | path | 4096 | valuesFor | ordered | baseline | 10 | 104 | 103 | 1.00× [0.98, 1.01] | -0.3% | [-1.6%, +0.9%] | 1.5 pts | 1.0 | yes | no |
| multi | path | 4096 | valuesBetween | ordered | baseline | 10 | 3883 | 3852 | 0.99× [0.98, 1.01] | -0.6% | [-2.4%, +1.2%] | 2.3 pts | 0.9 | yes | no |
| multi | path | 4096 | prefix | ordered | baseline | 10 | 312 | 308 | 0.99× [0.98, 1.00] | -1.0% | [-1.9%, -0.0%] | 1.6 pts | 1.1 | yes | yes |
| multi | path | 4096 | churn | ordered | baseline | 10 | 153 | 147 | 0.97× [0.96, 0.97] | -3.6% | [-4.5%, -2.7%] | 1.1 pts | 1.2 | yes | yes |
| multi | path | 4096 | build | ordered | baseline | 10 | 14.82 ms | 13.78 ms | 0.94× [0.93, 0.94] | -6.6% | [-7.1%, -6.1%] | 0.6 pts | 0.8 | yes | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 6 | 132 | 134 | 1.00× [0.99, 1.02] | +0.4% | [-1.2%, +2.1%] | 1.5 pts | 2.0 | yes | no |
| multi | path | 16384 | valuesBetween | ordered | baseline | 6 | 4592 | 4595 | 1.01× [1.00, 1.01] | +0.6% | [-0.2%, +1.5%] | 0.8 pts | 1.2 | yes | no |
| multi | path | 16384 | prefix | ordered | baseline | 6 | 709 | 713 | 1.00× [0.98, 1.01] | -0.3% | [-2.0%, +1.4%] | 1.6 pts | 0.6 | yes | no |
| multi | path | 16384 | churn | ordered | baseline | 6 | 210 | 204 | 0.97× [0.96, 0.97] | -3.4% | [-3.9%, -3.0%] | 0.4 pts | 0.8 | yes | yes |
| multi | path | 16384 | build | ordered | baseline | 6 | 73.23 ms | 69.41 ms | 0.95× [0.94, 0.95] | -5.6% | [-6.2%, -5.0%] | 0.6 pts | 1.2 | yes | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 12 | 414 | 412 | 1.00× [0.99, 1.00] | -0.3% | [-0.8%, +0.2%] | 0.8 pts | 0.8 | yes | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 12 | 10.3 µs | 10.3 µs | 1.00× [0.99, 1.00] | -0.3% | [-0.8%, +0.2%] | 1.0 pts | 1.4 | yes | no |
| multi | path | 262144 | prefix | ordered | baseline | 12 | 9313 | 9197 | 1.01× [0.99, 1.02] | +0.6% | [-0.6%, +1.8%] | 2.8 pts | 1.2 | yes | no |
| multi | path | 262144 | churn | ordered | baseline | 12 | 681 | 682 | 0.99× [0.96, 1.02] | -1.3% | [-4.4%, +1.9%] | 4.2 pts | 0.7 | no | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 10 | 60.9 | 60.2 | 0.99× [0.98, 1.00] | -1.1% | [-2.2%, +0.1%] | 1.6 pts | 0.7 | yes | no |
| multi | str | 4096 | valuesBetween | ordered | baseline | 10 | 3333 | 3314 | 0.99× [0.97, 1.01] | -1.1% | [-3.1%, +0.9%] | 2.5 pts | 1.0 | yes | no |
| multi | str | 4096 | prefix | ordered | baseline | 10 | 6999 | 6838 | 0.98× [0.97, 1.00] | -1.8% | [-3.6%, -0.1%] | 3.5 pts | 0.7 | yes | no |
| multi | str | 4096 | churn | ordered | baseline | 10 | 88.3 | 82.4 | 0.94× [0.93, 0.95] | -6.6% | [-7.8%, -5.5%] | 1.2 pts | 1.5 | yes | yes |
| multi | str | 4096 | build | ordered | baseline | 10 | 9.12 ms | 8.34 ms | 0.91× [0.91, 0.92] | -9.6% | [-10.1%, -9.2%] | 0.9 pts | 1.1 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 72.8 | 72.1 | 1.00× [0.98, 1.01] | -0.4% | [-1.8%, +0.9%] | 1.3 pts | 2.2 | yes | no |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3644 | 3601 | 0.99× [0.99, 1.00] | -0.8% | [-1.4%, -0.2%] | 0.6 pts | 1.0 | yes | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 33.9 µs | 33.6 µs | 0.99× [0.98, 1.00] | -1.0% | [-1.7%, -0.3%] | 0.7 pts | 0.9 | yes | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 112 | 108 | 0.95× [0.95, 0.96] | -4.9% | [-5.4%, -4.4%] | 0.5 pts | 0.6 | yes | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 42.13 ms | 38.98 ms | 0.93× [0.92, 0.93] | -7.7% | [-8.4%, -7.0%] | 0.7 pts | 1.4 | yes | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 12 | 219 | 215 | 1.00× [0.98, 1.01] | -0.5% | [-2.1%, +1.2%] | 2.0 pts | 3.0 | yes | no |
| multi | str | 262144 | valuesBetween | ordered | baseline | 12 | 7663 | 7575 | 0.99× [0.98, 1.01] | -0.6% | [-2.3%, +1.0%] | 2.0 pts | 3.1 | yes | no |
| multi | str | 262144 | prefix | ordered | baseline | 12 | 1.25 ms | 1.22 ms | 0.99× [0.98, 1.00] | -0.8% | [-1.8%, +0.3%] | 1.6 pts | 2.7 | yes | no |
| multi | str | 262144 | churn | ordered | baseline | 12 | 420 | 407 | 0.96× [0.93, 0.99] | -4.2% | [-7.0%, -1.4%] | 4.9 pts | 0.8 | no | yes |
| multi | street | 4096 | valuesFor | ordered | baseline | 10 | 48.2 | 47.8 | 1.00× [0.98, 1.01] | -0.2% | [-1.7%, +1.2%] | 1.8 pts | 0.8 | yes | no |
| multi | street | 4096 | valuesBetween | ordered | baseline | 10 | 2238 | 2275 | 1.02× [1.00, 1.04] | +2.1% | [+0.1%, +4.1%] | 2.9 pts | 1.2 | yes | yes |
| multi | street | 4096 | prefix | ordered | baseline | 10 | 226 | 218 | 0.96× [0.95, 0.98] | -3.8% | [-5.5%, -2.2%] | 1.7 pts | 1.2 | yes | yes |
| multi | street | 4096 | churn | ordered | baseline | 10 | 98.5 | 93.5 | 0.95× [0.94, 0.96] | -5.6% | [-6.5%, -4.6%] | 1.2 pts | 1.3 | yes | yes |
| multi | street | 4096 | build | ordered | baseline | 10 | 3.79 ms | 3.62 ms | 0.96× [0.96, 0.97] | -4.1% | [-4.7%, -3.6%] | 0.7 pts | 0.7 | yes | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 74.2 | 73.6 | 1.00× [0.99, 1.00] | -0.5% | [-1.1%, +0.1%] | 0.6 pts | 0.7 | yes | no |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2848 | 2841 | 1.00× [0.99, 1.01] | -0.1% | [-1.1%, +0.9%] | 1.0 pts | 1.3 | yes | no |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 809 | 797 | 0.99× [0.98, 1.00] | -1.4% | [-2.5%, -0.3%] | 1.0 pts | 0.5 | yes | no |
| multi | street | 16384 | churn | ordered | baseline | 6 | 135 | 131 | 0.97× [0.96, 0.98] | -3.1% | [-4.3%, -1.9%] | 1.1 pts | 1.7 | yes | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 21.76 ms | 19.62 ms | 0.90× [0.89, 0.91] | -11.1% | [-12.6%, -9.5%] | 1.5 pts | 1.7 | yes | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 10 | 32.5 | 31.7 | 0.97× [0.95, 0.98] | -3.4% | [-5.2%, -1.6%] | 2.0 pts | 0.3 | yes | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 10 | 2581 | 2504 | 0.97× [0.95, 0.98] | -3.4% | [-5.2%, -1.6%] | 1.8 pts | 0.5 | yes | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 10 | 53.5 | 49.6 | 0.93× [0.92, 0.93] | -8.0% | [-9.0%, -7.0%] | 1.2 pts | 1.6 | yes | yes |
| multi | u64 | 4096 | build | ordered | baseline | 10 | 5.75 ms | 5.17 ms | 0.90× [0.89, 0.92] | -10.6% | [-12.1%, -9.2%] | 1.3 pts | 1.2 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 46.8 | 46.1 | 0.98× [0.98, 0.98] | -1.8% | [-2.0%, -1.5%] | 0.2 pts | 0.5 | yes | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3497 | 3392 | 0.97× [0.96, 0.97] | -3.3% | [-3.7%, -2.9%] | 0.4 pts | 0.5 | yes | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 67.9 | 62.6 | 0.92× [0.91, 0.92] | -9.1% | [-9.7%, -8.4%] | 0.6 pts | 0.7 | yes | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 27.22 ms | 25.06 ms | 0.93× [0.92, 0.93] | -8.1% | [-8.6%, -7.6%] | 0.4 pts | 0.5 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 12 | 120 | 117 | 0.98× [0.95, 1.01] | -2.2% | [-4.9%, +0.5%] | 4.7 pts | 6.3 | no | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 6355 | 6376 | 1.00× [0.98, 1.02] | -0.3% | [-2.3%, +1.7%] | 2.9 pts | 4.2 | yes | no |
| multi | u64 | 262144 | churn | ordered | baseline | 12 | 295 | 280 | 0.93× [0.90, 0.95] | -8.0% | [-11.2%, -4.8%] | 5.1 pts | 0.9 | no | yes |
| multi | url | 4096 | valuesFor | ordered | baseline | 12 | 86.1 | 84.9 | 0.99× [0.98, 1.00] | -0.8% | [-1.5%, -0.1%] | 1.7 pts | 1.3 | yes | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 12 | 3718 | 3710 | 1.01× [0.99, 1.03] | +0.9% | [-1.4%, +3.1%] | 2.3 pts | 1.1 | no | no |
| multi | url | 4096 | prefix | ordered | baseline | 12 | 182 | 180 | 0.99× [0.98, 1.00] | -0.9% | [-1.8%, -0.0%] | 1.2 pts | 1.1 | yes | yes |
| multi | url | 4096 | churn | ordered | baseline | 12 | 128 | 122 | 0.96× [0.95, 0.97] | -4.3% | [-5.7%, -2.9%] | 1.5 pts | 1.8 | yes | yes |
| multi | url | 4096 | build | ordered | baseline | 12 | 12.50 ms | 11.54 ms | 0.92× [0.92, 0.92] | -8.6% | [-9.1%, -8.2%] | 0.5 pts | 0.9 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 109 | 109 | 0.99× [0.99, 1.00] | -0.5% | [-1.1%, -0.0%] | 0.5 pts | 0.7 | yes | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4404 | 4396 | 1.00× [0.98, 1.01] | -0.3% | [-1.9%, +1.2%] | 1.5 pts | 2.5 | yes | no |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 282 | 278 | 0.99× [0.98, 1.00] | -1.0% | [-2.2%, +0.2%] | 1.2 pts | 1.1 | yes | no |
| multi | url | 16384 | churn | ordered | baseline | 6 | 173 | 168 | 0.97× [0.96, 0.99] | -2.6% | [-3.8%, -1.3%] | 1.2 pts | 1.8 | yes | yes |
| multi | url | 16384 | build | ordered | baseline | 6 | 62.98 ms | 58.55 ms | 0.93× [0.92, 0.93] | -8.0% | [-8.7%, -7.4%] | 0.6 pts | 0.9 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 12 | 353 | 349 | 1.00× [0.98, 1.02] | -0.2% | [-2.1%, +1.8%] | 3.4 pts | 3.6 | yes | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 12 | 10.5 µs | 10.6 µs | 1.00× [0.99, 1.01] | -0.0% | [-1.2%, +1.2%] | 1.2 pts | 1.9 | yes | no |
| multi | url | 262144 | prefix | ordered | baseline | 12 | 2344 | 2323 | 1.00× [0.99, 1.00] | -0.4% | [-0.9%, +0.2%] | 1.0 pts | 0.7 | yes | no |
| multi | url | 262144 | churn | ordered | baseline | 12 | 621 | 622 | 0.96× [0.94, 0.98] | -3.9% | [-6.1%, -1.7%] | 6.9 pts | 1.1 | no | yes |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 53.2 | 52.7 | 0.99× [0.98, 1.01] | -0.7% | [-2.3%, +0.8%] | 1.5 pts | 0.6 | yes | no |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3140 | 3113 | 0.98× [0.97, 1.00] | -1.6% | [-2.8%, -0.3%] | 1.2 pts | 0.4 | yes | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 90.0 | 89.6 | 1.00× [0.99, 1.01] | +0.1% | [-1.0%, +1.1%] | 1.0 pts | 0.9 | yes | no |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 80.4 | 77.2 | 0.96× [0.95, 0.97] | -4.4% | [-5.6%, -3.2%] | 1.2 pts | 1.7 | yes | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 8.28 ms | 7.61 ms | 0.92× [0.91, 0.92] | -8.9% | [-9.5%, -8.2%] | 0.6 pts | 0.7 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 64.7 | 63.8 | 0.98× [0.98, 0.99] | -1.6% | [-2.0%, -1.1%] | 0.4 pts | 0.9 | yes | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3576 | 3549 | 0.99× [0.99, 1.00] | -0.8% | [-1.2%, -0.4%] | 0.4 pts | 0.9 | yes | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 108 | 108 | 1.00× [0.99, 1.00] | -0.5% | [-1.1%, +0.2%] | 0.6 pts | 1.1 | yes | no |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 103 | 99.5 | 0.97× [0.97, 0.98] | -3.0% | [-3.6%, -2.3%] | 0.6 pts | 1.1 | yes | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 38.35 ms | 35.79 ms | 0.93× [0.93, 0.94] | -7.0% | [-7.7%, -6.3%] | 0.7 pts | 1.2 | yes | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 12 | 225 | 219 | 0.97× [0.96, 0.99] | -2.6% | [-4.1%, -1.0%] | 2.3 pts | 3.6 | yes | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 8111 | 8138 | 1.00× [0.99, 1.01] | -0.0% | [-1.4%, +1.3%] | 1.9 pts | 2.5 | yes | no |
| multi | uuid | 262144 | prefix | ordered | baseline | 12 | 523 | 534 | 1.01× [1.00, 1.02] | +1.2% | [+0.2%, +2.3%] | 1.5 pts | 3.5 | yes | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 12 | 454 | 433 | 0.94× [0.90, 0.98] | -6.8% | [-11.2%, -2.4%] | 5.1 pts | 0.7 | no | yes |
| unique | email | 4096 | valuesFor | ordered | baseline | 6 | 33.4 | 22.1 | 0.66× [0.64, 0.67] | -52.0% | [-55.4%, -48.6%] | 3.2 pts | 1.1 | yes | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 6 | 342 | 1217 | 3.59× [3.57, 3.61] | +72.1% | [+72.0%, +72.3%] | 0.2 pts | 0.6 | yes | yes |
| unique | email | 4096 | prefix | ordered | baseline | 6 | 57.5 | 57.3 | 1.00× [0.98, 1.01] | -0.2% | [-1.7%, +1.4%] | 1.5 pts | 2.2 | yes | no |
| unique | email | 4096 | churn | ordered | baseline | 6 | 91.8 | 70.3 | 0.76× [0.76, 0.77] | -30.8% | [-31.8%, -29.7%] | 1.0 pts | 0.9 | yes | yes |
| unique | email | 4096 | build | ordered | baseline | 6 | 1.44 ms | 1.02 ms | 0.70× [0.68, 0.72] | -42.5% | [-46.1%, -38.9%] | 3.4 pts | 2.2 | yes | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 53.7 | 37.7 | 0.70× [0.70, 0.71] | -42.4% | [-43.5%, -41.2%] | 1.1 pts | 1.5 | yes | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 437 | 1565 | 3.58× [3.55, 3.61] | +72.0% | [+71.8%, +72.3%] | 0.2 pts | 1.4 | yes | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 81.2 | 72.4 | 0.89× [0.88, 0.89] | -12.7% | [-13.1%, -12.4%] | 0.3 pts | 0.6 | yes | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 128 | 90.8 | 0.71× [0.70, 0.72] | -41.1% | [-42.7%, -39.4%] | 1.5 pts | 1.1 | yes | yes |
| unique | email | 16384 | build | ordered | baseline | 6 | 7.86 ms | 5.04 ms | 0.64× [0.64, 0.65] | -56.0% | [-57.2%, -54.8%] | 1.1 pts | 1.1 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 12 | 154 | 168 | 1.08× [1.05, 1.11] | +7.3% | [+4.6%, +10.0%] | 3.4 pts | 2.5 | no | yes |
| unique | email | 262144 | valuesBetween | ordered | baseline | 12 | 1105 | 2530 | 2.29× [2.22, 2.38] | +56.4% | [+54.9%, +57.9%] | 1.8 pts | 1.0 | yes | yes |
| unique | email | 262144 | prefix | ordered | baseline | 12 | 181 | 210 | 1.16× [1.14, 1.18] | +13.8% | [+12.2%, +15.5%] | 1.8 pts | 2.0 | yes | yes |
| unique | email | 262144 | churn | ordered | baseline | 12 | 405 | 325 | 0.79× [0.78, 0.80] | -27.3% | [-28.8%, -25.7%] | 2.4 pts | 1.0 | yes | yes |
| unique | path | 4096 | valuesFor | ordered | baseline | 6 | 92.2 | 77.9 | 0.85× [0.84, 0.86] | -17.6% | [-18.8%, -16.5%] | 1.1 pts | 0.9 | yes | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 6 | 615 | 1966 | 3.21× [3.19, 3.24] | +68.9% | [+68.6%, +69.2%] | 0.3 pts | 0.6 | yes | yes |
| unique | path | 4096 | prefix | ordered | baseline | 6 | 211 | 235 | 1.11× [1.10, 1.13] | +10.1% | [+9.0%, +11.2%] | 1.1 pts | 1.2 | yes | yes |
| unique | path | 4096 | churn | ordered | baseline | 6 | 223 | 195 | 0.88× [0.87, 0.89] | -13.9% | [-14.9%, -12.9%] | 1.0 pts | 0.8 | yes | yes |
| unique | path | 4096 | build | ordered | baseline | 6 | 3.48 ms | 2.60 ms | 0.75× [0.74, 0.75] | -33.8% | [-34.6%, -33.0%] | 0.8 pts | 0.7 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 6 | 125 | 113 | 0.90× [0.89, 0.92] | -10.9% | [-12.7%, -9.0%] | 1.8 pts | 2.2 | yes | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 6 | 770 | 2536 | 3.30× [3.28, 3.32] | +69.7% | [+69.5%, +69.8%] | 0.2 pts | 0.7 | yes | yes |
| unique | path | 16384 | prefix | ordered | baseline | 6 | 303 | 477 | 1.56× [1.53, 1.59] | +36.0% | [+34.8%, +37.1%] | 1.1 pts | 1.0 | yes | yes |
| unique | path | 16384 | churn | ordered | baseline | 6 | 287 | 250 | 0.87× [0.86, 0.87] | -15.5% | [-16.4%, -14.6%] | 0.9 pts | 0.8 | yes | yes |
| unique | path | 16384 | build | ordered | baseline | 6 | 17.53 ms | 12.88 ms | 0.74× [0.73, 0.74] | -36.0% | [-37.7%, -34.2%] | 1.6 pts | 1.2 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 6 | 342 | 363 | 1.05× [1.03, 1.06] | +4.4% | [+3.1%, +5.7%] | 1.2 pts | 1.2 | yes | yes |
| unique | path | 262144 | valuesBetween | ordered | baseline | 6 | 1993 | 4763 | 2.38× [2.30, 2.47] | +58.0% | [+56.6%, +59.5%] | 1.4 pts | 1.6 | yes | yes |
| unique | path | 262144 | prefix | ordered | baseline | 6 | 1707 | 5586 | 3.32× [3.21, 3.45] | +69.9% | [+68.8%, +71.0%] | 1.0 pts | 0.4 | yes | yes |
| unique | path | 262144 | churn | ordered | baseline | 6 | 703 | 687 | 0.98× [0.97, 1.00] | -1.7% | [-2.8%, -0.5%] | 1.1 pts | 0.5 | yes | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 10 | 40.5 | 31.9 | 0.80× [0.79, 0.81] | -25.3% | [-27.2%, -23.4%] | 3.5 pts | 1.0 | yes | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 10 | 379 | 1582 | 4.21× [4.16, 4.26] | +76.2% | [+75.9%, +76.5%] | 0.3 pts | 1.6 | yes | yes |
| unique | str | 4096 | prefix | ordered | baseline | 10 | 603 | 3092 | 5.14× [5.13, 5.15] | +80.5% | [+80.5%, +80.6%] | 0.1 pts | 0.8 | yes | yes |
| unique | str | 4096 | churn | ordered | baseline | 10 | 112 | 94.3 | 0.85× [0.84, 0.85] | -18.2% | [-19.0%, -17.4%] | 1.1 pts | 0.9 | yes | yes |
| unique | str | 4096 | build | ordered | baseline | 10 | 1.74 ms | 1.35 ms | 0.78× [0.76, 0.79] | -29.0% | [-31.7%, -26.4%] | 2.8 pts | 1.2 | yes | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 8 | 55.9 | 48.6 | 0.87× [0.86, 0.87] | -15.4% | [-16.3%, -14.6%] | 0.9 pts | 1.3 | yes | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 8 | 451 | 1607 | 3.57× [3.56, 3.58] | +72.0% | [+71.9%, +72.1%] | 0.1 pts | 0.5 | yes | yes |
| unique | str | 16384 | prefix | ordered | baseline | 8 | 2445 | 13.2 µs | 5.40× [5.35, 5.45] | +81.5% | [+81.3%, +81.7%] | 0.2 pts | 1.5 | yes | yes |
| unique | str | 16384 | churn | ordered | baseline | 8 | 137 | 114 | 0.84× [0.82, 0.85] | -19.5% | [-21.3%, -17.8%] | 2.0 pts | 1.7 | yes | yes |
| unique | str | 16384 | build | ordered | baseline | 8 | 8.30 ms | 6.29 ms | 0.76× [0.75, 0.77] | -31.5% | [-32.8%, -30.3%] | 1.8 pts | 1.8 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 12 | 107 | 163 | 1.48× [1.44, 1.53] | +32.5% | [+30.4%, +34.6%] | 2.4 pts | 1.9 | yes | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 12 | 768 | 2199 | 2.87× [2.79, 2.95] | +65.1% | [+64.2%, +66.1%] | 1.4 pts | 0.8 | yes | yes |
| unique | str | 262144 | prefix | ordered | baseline | 12 | 70.4 µs | 300.4 µs | 4.27× [4.04, 4.52] | +76.6% | [+75.3%, +77.9%] | 1.6 pts | 0.9 | yes | yes |
| unique | str | 262144 | churn | ordered | baseline | 12 | 315 | 341 | 1.08× [1.05, 1.12] | +7.5% | [+4.6%, +10.4%] | 3.5 pts | 1.7 | no | yes |
| unique | street | 4096 | valuesFor | ordered | baseline | 12 | 48.5 | 38.1 | 0.80× [0.77, 0.83] | -25.5% | [-30.1%, -20.9%] | 5.0 pts | 1.9 | no | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 12 | 444 | 1732 | 3.89× [3.87, 3.92] | +74.3% | [+74.2%, +74.5%] | 0.2 pts | 1.0 | yes | yes |
| unique | street | 4096 | prefix | ordered | baseline | 12 | 104 | 192 | 1.84× [1.83, 1.85] | +45.7% | [+45.5%, +45.9%] | 0.5 pts | 1.4 | yes | yes |
| unique | street | 4096 | churn | ordered | baseline | 12 | 121 | 114 | 0.95× [0.93, 0.96] | -5.8% | [-7.2%, -4.3%] | 1.6 pts | 1.5 | yes | yes |
| unique | street | 4096 | build | ordered | baseline | 12 | 2.16 ms | 1.61 ms | 0.74× [0.74, 0.75] | -34.9% | [-36.0%, -33.8%] | 1.6 pts | 1.3 | yes | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 6 | 66.7 | 66.2 | 0.99× [0.98, 1.00] | -1.2% | [-2.2%, -0.1%] | 1.0 pts | 1.5 | yes | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 6 | 541 | 2160 | 4.01× [3.97, 4.05] | +75.0% | [+74.8%, +75.3%] | 0.2 pts | 1.0 | yes | yes |
| unique | street | 16384 | prefix | ordered | baseline | 6 | 195 | 614 | 3.15× [3.12, 3.19] | +68.3% | [+67.9%, +68.6%] | 0.3 pts | 1.2 | yes | yes |
| unique | street | 16384 | churn | ordered | baseline | 6 | 156 | 151 | 0.96× [0.95, 0.97] | -4.0% | [-5.3%, -2.6%] | 1.3 pts | 1.5 | yes | yes |
| unique | street | 16384 | build | ordered | baseline | 6 | 10.21 ms | 7.99 ms | 0.79× [0.78, 0.79] | -27.3% | [-28.1%, -26.4%] | 0.8 pts | 0.7 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 12 | 12.6 | 12.9 | 1.01× [0.93, 1.10] | +1.2% | [-7.0%, +9.4%] | 7.7 pts | 9.6 | no | no |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 12 | 249 | 960 | 3.85× [3.84, 3.86] | +74.0% | [+73.9%, +74.1%] | 0.1 pts | 1.1 | yes | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 12 | 46.1 | 45.8 | 1.00× [1.00, 1.01] | +0.3% | [-0.2%, +0.9%] | 1.5 pts | 3.2 | yes | no |
| unique | u64 | 4096 | build | ordered | baseline | 12 | 718.6 µs | 766.8 µs | 1.06× [1.06, 1.07] | +6.0% | [+5.8%, +6.3%] | 0.5 pts | 0.7 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 6 | 29.8 | 23.3 | 0.78× [0.77, 0.79] | -28.8% | [-30.2%, -27.3%] | 1.4 pts | 1.7 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 292 | 1456 | 4.98× [4.96, 5.01] | +79.9% | [+79.8%, +80.0%] | 0.1 pts | 0.9 | yes | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 6 | 68.4 | 52.5 | 0.77× [0.75, 0.78] | -30.4% | [-33.1%, -27.8%] | 2.5 pts | 3.8 | yes | yes |
| unique | u64 | 16384 | build | ordered | baseline | 6 | 4.13 ms | 3.52 ms | 0.85× [0.84, 0.86] | -17.5% | [-19.1%, -15.8%] | 1.6 pts | 2.2 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 12 | 41.6 | 42.9 | 1.04× [1.01, 1.07] | +3.9% | [+0.8%, +6.9%] | 7.7 pts | 5.9 | no | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 311 | 1388 | 4.44× [4.40, 4.48] | +77.5% | [+77.3%, +77.7%] | 0.5 pts | 0.9 | yes | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 12 | 117 | 171 | 1.40× [1.36, 1.44] | +28.3% | [+26.3%, +30.4%] | 3.6 pts | 0.8 | yes | yes |
| unique | url | 4096 | valuesFor | ordered | baseline | 12 | 69.9 | 59.4 | 0.86× [0.83, 0.89] | -16.7% | [-20.4%, -13.0%] | 3.7 pts | 2.1 | no | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 12 | 506 | 1846 | 3.66× [3.62, 3.71] | +72.7% | [+72.4%, +73.0%] | 0.6 pts | 1.6 | yes | yes |
| unique | url | 4096 | prefix | ordered | baseline | 12 | 147 | 146 | 0.99× [0.98, 1.00] | -1.1% | [-2.2%, +0.0%] | 1.6 pts | 1.8 | yes | no |
| unique | url | 4096 | churn | ordered | baseline | 12 | 183 | 154 | 0.84× [0.84, 0.85] | -18.8% | [-19.4%, -18.2%] | 1.0 pts | 0.8 | yes | yes |
| unique | url | 4096 | build | ordered | baseline | 12 | 2.79 ms | 2.17 ms | 0.77× [0.76, 0.77] | -30.7% | [-31.6%, -29.7%] | 1.3 pts | 0.6 | yes | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 95.8 | 87.5 | 0.91× [0.91, 0.92] | -9.4% | [-10.4%, -8.5%] | 0.9 pts | 1.5 | yes | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 659 | 2371 | 3.59× [3.55, 3.65] | +72.2% | [+71.8%, +72.6%] | 0.4 pts | 1.4 | yes | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 191 | 215 | 1.13× [1.12, 1.14] | +11.3% | [+10.4%, +12.2%] | 0.8 pts | 1.0 | yes | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 229 | 200 | 0.87× [0.86, 0.88] | -15.2% | [-16.8%, -13.5%] | 1.6 pts | 1.4 | yes | yes |
| unique | url | 16384 | build | ordered | baseline | 6 | 14.59 ms | 10.41 ms | 0.71× [0.70, 0.72] | -40.6% | [-42.7%, -38.4%] | 2.0 pts | 1.1 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 12 | 312 | 340 | 1.06× [1.04, 1.09] | +6.1% | [+3.6%, +8.5%] | 2.9 pts | 3.0 | no | yes |
| unique | url | 262144 | valuesBetween | ordered | baseline | 12 | 1905 | 5174 | 2.76× [2.72, 2.81] | +63.8% | [+63.2%, +64.4%] | 1.2 pts | 1.1 | yes | yes |
| unique | url | 262144 | prefix | ordered | baseline | 12 | 667 | 1219 | 1.81× [1.78, 1.84] | +44.8% | [+43.8%, +45.8%] | 1.2 pts | 1.5 | yes | yes |
| unique | url | 262144 | churn | ordered | baseline | 12 | 613 | 595 | 0.95× [0.93, 0.97] | -5.2% | [-7.1%, -3.4%] | 3.2 pts | 1.3 | yes | yes |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 12 | 37.7 | 26.8 | 0.73× [0.68, 0.78] | -37.4% | [-47.2%, -27.6%] | 9.8 pts | 3.6 | no | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 396 | 1389 | 3.51× [3.48, 3.55] | +71.5% | [+71.3%, +71.8%] | 0.4 pts | 1.6 | yes | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 12 | 71.5 | 68.9 | 0.97× [0.95, 0.98] | -3.4% | [-5.0%, -1.7%] | 1.5 pts | 2.5 | yes | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 12 | 101 | 81.6 | 0.81× [0.81, 0.81] | -23.8% | [-24.2%, -23.5%] | 1.0 pts | 1.0 | yes | yes |
| unique | uuid | 4096 | build | ordered | baseline | 12 | 1.66 ms | 1.14 ms | 0.70× [0.69, 0.70] | -43.8% | [-45.1%, -42.6%] | 2.6 pts | 1.3 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 6 | 47.3 | 41.9 | 0.88× [0.87, 0.89] | -13.6% | [-15.1%, -12.1%] | 1.4 pts | 2.6 | yes | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 404 | 1574 | 3.88× [3.84, 3.92] | +74.2% | [+73.9%, +74.5%] | 0.3 pts | 1.9 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 79.6 | 84.9 | 1.06× [1.05, 1.08] | +6.1% | [+5.1%, +7.1%] | 1.0 pts | 1.9 | yes | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 6 | 123 | 101 | 0.82× [0.81, 0.83] | -22.1% | [-23.5%, -20.7%] | 1.3 pts | 1.0 | yes | yes |
| unique | uuid | 16384 | build | ordered | baseline | 6 | 7.11 ms | 5.43 ms | 0.76× [0.75, 0.77] | -31.5% | [-32.6%, -30.4%] | 1.1 pts | 1.4 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 12 | 177 | 187 | 1.04× [1.02, 1.05] | +3.6% | [+2.2%, +4.9%] | 1.6 pts | 1.5 | yes | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 995 | 2812 | 2.83× [2.73, 2.92] | +64.6% | [+63.4%, +65.8%] | 1.3 pts | 0.8 | yes | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 12 | 186 | 310 | 1.66× [1.63, 1.68] | +39.6% | [+38.6%, +40.6%] | 1.2 pts | 1.6 | yes | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 12 | 364 | 351 | 0.93× [0.92, 0.95] | -7.2% | [-9.1%, -5.3%] | 3.5 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the pooled interval [-8.56%, 5.01%] includes zero
- multi email n=4096 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-3.42%, 0.32%] includes zero
- multi email n=4096 prefix: ordered vs baseline: the pooled interval [-8.23%, 1.62%] includes zero
- multi email n=4096 prefix: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.40%, 0.04%] includes zero
- multi email n=16384 prefix: ordered vs baseline: the pooled difference of 0.11% does not clear the 0.31% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=16384 prefix: ordered vs baseline: the pooled interval [-1.31%, 1.54%] includes zero
- multi email n=16384 prefix: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=16384 prefix: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.62%, 0.39%] includes zero
- multi email n=262144 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.16% does not clear the 0.47% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.36%, 1.03%] includes zero
- multi email n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=262144 prefix: ordered vs baseline: the pooled difference of -0.00% does not clear the 0.36% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=262144 prefix: ordered vs baseline: the pooled interval [-1.58%, 1.58%] includes zero
- multi email n=262144 prefix: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 prefix: ordered vs baseline: 4 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.34% does not clear the 0.71% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.62%, 0.95%] includes zero
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.62% does not clear the 0.99% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.42%, 1.18%] includes zero
- multi path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.45% does not clear the 0.63% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.16%, 2.05%] includes zero
- multi path n=16384 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.61% does not clear the 1.21% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.23%, 1.46%] includes zero
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of -0.29% does not clear the 1.90% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-1.98%, 1.40%] includes zero
- multi path n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.30% does not clear the 0.45% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.76%, 0.16%] includes zero
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.27% does not clear the 0.58% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.76%, 0.21%] includes zero
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of 0.59% does not clear the 5.09% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 prefix: ordered vs baseline: the pooled interval [-0.57%, 1.75%] includes zero
- multi path n=262144 churn: ordered vs baseline: the pooled difference of -1.25% does not clear the 2.52% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 churn: ordered vs baseline: the pooled interval [-4.36%, 1.86%] includes zero
- multi str n=4096 valuesFor: ordered vs baseline: the pooled difference of -1.06% does not clear the 1.19% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.17%, 0.05%] includes zero
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.10% does not clear the 1.25% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-3.09%, 0.88%] includes zero
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of -1.85% does not clear the 2.72% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.44% does not clear the 0.53% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.77%, 0.89%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.24% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi str n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.13%, 1.21%] includes zero
- multi str n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.29%, 1.01%] includes zero
- multi str n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=262144 prefix: ordered vs baseline: the pooled interval [-1.82%, 0.29%] includes zero
- multi str n=262144 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.22% does not clear the 1.79% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.66%, 1.23%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.50% does not clear the 0.73% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.13%, 0.13%] includes zero
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.11% does not clear the 0.63% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.13%, 0.91%] includes zero
- multi street n=16384 prefix: ordered vs baseline: the pooled difference of -1.42% does not clear the 1.87% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-4.93%, 0.51%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 6.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.28% does not clear the 0.42% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.26%, 1.69%] includes zero
- multi u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesBetween: ordered vs baseline: 6 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.88% does not clear the 1.32% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.37%, 3.12%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.33% does not clear the 1.05% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.86%, 1.19%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=16384 prefix: ordered vs baseline: the pooled difference of -1.02% does not clear the 1.76% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=16384 prefix: ordered vs baseline: the pooled interval [-2.22%, 0.19%] includes zero
- multi url n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.16% does not clear the 0.61% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.12%, 1.81%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.01% does not clear the 0.44% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.17%, 1.15%] includes zero
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of -0.37% does not clear the 1.69% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=262144 prefix: ordered vs baseline: the pooled interval [-0.90%, 0.17%] includes zero
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.74% does not clear the 1.25% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.67% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.28%, 0.79%] includes zero
- multi uuid n=4096 prefix: ordered vs baseline: the pooled difference of 0.07% does not clear the 1.04% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 prefix: ordered vs baseline: the pooled interval [-0.96%, 1.11%] includes zero
- multi uuid n=16384 prefix: ordered vs baseline: the pooled interval [-1.11%, 0.19%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.05% does not clear the 0.43% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.44%, 1.35%] includes zero
- multi uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 prefix: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 prefix: ordered vs baseline: 6 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=4096 prefix: ordered vs baseline: the pooled difference of -0.15% does not clear the 0.49% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=4096 prefix: ordered vs baseline: the pooled interval [-1.68%, 1.37%] includes zero
- unique email n=4096 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=4096 prefix: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=4096 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=262144 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -1.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique path n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -1.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique str n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of -1.11% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-7.04%, 9.40%] includes zero
- unique u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 9.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 valuesFor: ordered vs baseline: 10 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=4096 churn: ordered vs baseline: the pooled interval [-0.24%, 0.93%] includes zero
- unique u64 n=4096 churn: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 churn: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of -0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=16384 churn: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 5.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: 6 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=4096 prefix: ordered vs baseline: the pooled interval [-2.18%, 0.00%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 prefix: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 prefix: ordered vs baseline: 1 processes resolved A as faster and 11 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=16384 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
