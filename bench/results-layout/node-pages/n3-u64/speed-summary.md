| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 12 | 45.7 | 46.1 | 1.01× [1.00, 1.03] | +1.2% | [-0.4%, +2.9%] | 2.6 pts | 1.1 | yes | no |
| multi | email | 4096 | valuesBetween | ordered | baseline | 12 | 2822 | 2864 | 1.01× [0.99, 1.03] | +1.3% | [-0.7%, +3.3%] | 2.2 pts | 0.9 | no | no |
| multi | email | 4096 | prefix | ordered | baseline | 12 | 77.4 | 76.5 | 0.99× [0.98, 1.00] | -1.1% | [-2.3%, +0.1%] | 1.4 pts | 1.2 | yes | no |
| multi | email | 4096 | churn | ordered | baseline | 12 | 68.6 | 66.7 | 0.97× [0.97, 0.98] | -2.9% | [-3.6%, -2.2%] | 0.9 pts | 1.5 | yes | yes |
| multi | email | 4096 | build | ordered | baseline | 12 | 6.91 ms | 6.80 ms | 0.98× [0.98, 0.99] | -2.0% | [-2.6%, -1.4%] | 0.8 pts | 1.1 | yes | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 59.9 | 59.6 | 1.00× [0.99, 1.00] | -0.3% | [-1.1%, +0.4%] | 0.7 pts | 1.3 | yes | no |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3470 | 3453 | 0.99× [0.99, 1.00] | -0.6% | [-1.0%, -0.1%] | 0.4 pts | 0.7 | yes | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 90.9 | 91.1 | 1.00× [1.00, 1.01] | +0.2% | [-0.3%, +0.7%] | 0.5 pts | 1.0 | yes | no |
| multi | email | 16384 | churn | ordered | baseline | 6 | 93.7 | 90.3 | 0.97× [0.96, 0.97] | -3.5% | [-4.2%, -2.8%] | 0.6 pts | 1.2 | yes | yes |
| multi | email | 16384 | build | ordered | baseline | 6 | 33.71 ms | 33.05 ms | 0.98× [0.97, 0.98] | -2.3% | [-2.8%, -1.7%] | 0.5 pts | 1.0 | yes | yes |
| multi | email | 65536 | valuesFor | ordered | baseline | 12 | 81.3 | 80.4 | 0.98× [0.96, 1.00] | -1.8% | [-3.9%, +0.4%] | 2.8 pts | 3.0 | no | no |
| multi | email | 65536 | valuesBetween | ordered | baseline | 12 | 3869 | 3837 | 0.99× [0.98, 1.01] | -0.7% | [-2.0%, +0.7%] | 2.3 pts | 2.5 | yes | no |
| multi | email | 65536 | prefix | ordered | baseline | 12 | 130 | 128 | 0.99× [0.98, 1.00] | -0.9% | [-2.0%, +0.2%] | 1.4 pts | 2.6 | yes | no |
| multi | email | 65536 | churn | ordered | baseline | 12 | 177 | 174 | 0.97× [0.96, 0.98] | -3.5% | [-4.5%, -2.5%] | 1.5 pts | 0.9 | yes | yes |
| multi | email | 65536 | build | ordered | baseline | 12 | 204.45 ms | 197.67 ms | 0.98× [0.97, 0.98] | -2.6% | [-3.0%, -2.1%] | 1.1 pts | 1.1 | yes | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 12 | 196 | 192 | 0.98× [0.97, 1.00] | -1.6% | [-3.2%, -0.1%] | 2.3 pts | 3.3 | yes | yes |
| multi | email | 262144 | valuesBetween | ordered | baseline | 12 | 7264 | 7217 | 0.99× [0.99, 1.00] | -0.5% | [-1.4%, +0.4%] | 1.3 pts | 1.9 | yes | no |
| multi | email | 262144 | prefix | ordered | baseline | 12 | 259 | 256 | 1.01× [0.99, 1.03] | +0.6% | [-1.2%, +2.5%] | 2.3 pts | 4.4 | yes | no |
| multi | email | 262144 | churn | ordered | baseline | 12 | 399 | 380 | 0.93× [0.89, 0.97] | -7.5% | [-12.4%, -2.7%] | 5.4 pts | 0.8 | no | yes |
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 104 | 103 | 0.98× [0.96, 1.00] | -2.0% | [-4.0%, -0.1%] | 1.8 pts | 1.3 | yes | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 3815 | 3863 | 1.01× [0.99, 1.03] | +1.0% | [-0.7%, +2.6%] | 1.6 pts | 0.7 | yes | no |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 307 | 308 | 1.01× [0.99, 1.02] | +0.7% | [-0.8%, +2.2%] | 1.4 pts | 1.0 | yes | no |
| multi | path | 4096 | churn | ordered | baseline | 6 | 154 | 150 | 0.98× [0.96, 0.99] | -2.3% | [-3.8%, -0.8%] | 1.5 pts | 1.8 | yes | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 14.39 ms | 13.98 ms | 0.97× [0.97, 0.98] | -3.0% | [-3.5%, -2.5%] | 0.5 pts | 0.8 | yes | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 6 | 134 | 133 | 1.01× [0.99, 1.02] | +0.5% | [-1.0%, +2.1%] | 1.4 pts | 1.8 | yes | no |
| multi | path | 16384 | valuesBetween | ordered | baseline | 6 | 4532 | 4567 | 1.01× [1.00, 1.02] | +0.9% | [-0.2%, +1.9%] | 1.0 pts | 1.8 | yes | no |
| multi | path | 16384 | prefix | ordered | baseline | 6 | 707 | 720 | 1.01× [0.99, 1.03] | +1.0% | [-0.7%, +2.7%] | 1.6 pts | 0.6 | yes | no |
| multi | path | 16384 | churn | ordered | baseline | 6 | 210 | 204 | 0.97× [0.96, 0.98] | -2.9% | [-3.6%, -2.2%] | 0.7 pts | 1.2 | yes | yes |
| multi | path | 16384 | build | ordered | baseline | 6 | 71.96 ms | 69.95 ms | 0.97× [0.97, 0.98] | -2.9% | [-3.4%, -2.4%] | 0.5 pts | 1.1 | yes | yes |
| multi | path | 65536 | valuesFor | ordered | baseline | 12 | 205 | 206 | 1.00× [0.99, 1.01] | -0.1% | [-1.3%, +1.2%] | 2.5 pts | 2.4 | yes | no |
| multi | path | 65536 | valuesBetween | ordered | baseline | 12 | 5733 | 5785 | 1.01× [1.00, 1.01] | +0.6% | [-0.2%, +1.4%] | 3.3 pts | 2.8 | yes | no |
| multi | path | 65536 | prefix | ordered | baseline | 12 | 2429 | 2366 | 1.00× [0.99, 1.01] | -0.2% | [-1.5%, +1.0%] | 2.4 pts | 1.1 | yes | no |
| multi | path | 65536 | churn | ordered | baseline | 12 | 394 | 385 | 0.98× [0.94, 1.01] | -2.5% | [-6.0%, +1.0%] | 4.0 pts | 0.9 | no | no |
| multi | path | 65536 | build | ordered | baseline | 12 | 460.94 ms | 449.83 ms | 0.98× [0.97, 0.98] | -2.5% | [-3.3%, -1.8%] | 0.9 pts | 1.0 | yes | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 12 | 391 | 391 | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.3%] | 1.3 pts | 1.5 | yes | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 12 | 10.3 µs | 10.3 µs | 1.00× [0.99, 1.01] | +0.1% | [-0.6%, +0.8%] | 1.1 pts | 1.4 | yes | no |
| multi | path | 262144 | prefix | ordered | baseline | 12 | 9244 | 9150 | 1.00× [0.98, 1.02] | -0.1% | [-2.1%, +1.9%] | 2.0 pts | 0.7 | yes | no |
| multi | path | 262144 | churn | ordered | baseline | 12 | 694 | 683 | 0.95× [0.92, 0.97] | -5.8% | [-8.4%, -3.1%] | 6.0 pts | 0.8 | no | yes |
| multi | str | 4096 | valuesFor | ordered | baseline | 12 | 59.6 | 59.7 | 1.00× [0.99, 1.00] | -0.3% | [-1.1%, +0.5%] | 1.4 pts | 0.8 | yes | no |
| multi | str | 4096 | valuesBetween | ordered | baseline | 12 | 3256 | 3336 | 1.03× [1.03, 1.04] | +3.0% | [+2.5%, +3.5%] | 2.6 pts | 1.1 | yes | yes |
| multi | str | 4096 | prefix | ordered | baseline | 12 | 6728 | 7009 | 1.04× [1.00, 1.09] | +4.0% | [-0.2%, +8.1%] | 4.5 pts | 1.1 | no | no |
| multi | str | 4096 | churn | ordered | baseline | 12 | 85.4 | 83.1 | 0.97× [0.96, 0.98] | -3.0% | [-3.7%, -2.3%] | 0.9 pts | 1.1 | yes | yes |
| multi | str | 4096 | build | ordered | baseline | 12 | 8.52 ms | 8.34 ms | 0.98× [0.97, 0.99] | -2.2% | [-3.3%, -1.1%] | 1.3 pts | 1.8 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 72.3 | 72.0 | 0.99× [0.99, 1.00] | -0.5% | [-1.3%, +0.3%] | 0.8 pts | 1.5 | yes | no |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3602 | 3606 | 1.00× [0.99, 1.01] | -0.2% | [-1.1%, +0.6%] | 0.8 pts | 1.5 | yes | no |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 33.4 µs | 33.7 µs | 1.00× [1.00, 1.01] | +0.3% | [-0.5%, +1.2%] | 0.8 pts | 1.2 | yes | no |
| multi | str | 16384 | churn | ordered | baseline | 6 | 112 | 108 | 0.98× [0.97, 0.99] | -2.2% | [-3.0%, -1.4%] | 0.8 pts | 1.4 | yes | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 39.99 ms | 38.89 ms | 0.97× [0.97, 0.97] | -3.0% | [-3.3%, -2.8%] | 0.2 pts | 0.5 | yes | yes |
| multi | str | 65536 | valuesFor | ordered | baseline | 6 | 97.9 | 96.6 | 0.99× [0.97, 1.01] | -1.0% | [-3.0%, +0.9%] | 1.9 pts | 2.3 | yes | no |
| multi | str | 65536 | valuesBetween | ordered | baseline | 6 | 4057 | 4003 | 0.99× [0.98, 1.01] | -0.7% | [-2.2%, +0.8%] | 1.5 pts | 1.5 | yes | no |
| multi | str | 65536 | prefix | ordered | baseline | 6 | 153.1 µs | 149.3 µs | 0.99× [0.97, 1.01] | -0.9% | [-2.8%, +0.9%] | 1.8 pts | 1.5 | yes | no |
| multi | str | 65536 | churn | ordered | baseline | 6 | 208 | 203 | 0.97× [0.96, 0.99] | -2.6% | [-3.8%, -1.3%] | 1.2 pts | 1.2 | yes | yes |
| multi | str | 65536 | build | ordered | baseline | 6 | 238.42 ms | 229.94 ms | 0.97× [0.95, 0.98] | -3.4% | [-4.9%, -1.8%] | 1.5 pts | 1.4 | yes | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 12 | 215 | 213 | 1.00× [0.98, 1.01] | -0.5% | [-2.2%, +1.3%] | 2.3 pts | 3.0 | yes | no |
| multi | str | 262144 | valuesBetween | ordered | baseline | 12 | 7648 | 7548 | 0.99× [0.97, 1.01] | -1.0% | [-2.7%, +0.7%] | 1.8 pts | 2.5 | yes | no |
| multi | str | 262144 | prefix | ordered | baseline | 12 | 1.25 ms | 1.24 ms | 1.00× [0.99, 1.00] | -0.4% | [-1.3%, +0.5%] | 1.3 pts | 2.1 | yes | no |
| multi | str | 262144 | churn | ordered | baseline | 12 | 423 | 405 | 0.93× [0.90, 0.95] | -7.9% | [-10.7%, -5.1%] | 3.9 pts | 0.6 | no | yes |
| multi | street | 4096 | valuesFor | ordered | baseline | 8 | 48.5 | 48.3 | 1.00× [0.98, 1.01] | -0.5% | [-2.3%, +1.3%] | 1.9 pts | 0.8 | yes | no |
| multi | street | 4096 | valuesBetween | ordered | baseline | 8 | 2241 | 2280 | 1.01× [0.99, 1.03] | +0.9% | [-0.8%, +2.6%] | 2.6 pts | 1.0 | yes | no |
| multi | street | 4096 | prefix | ordered | baseline | 8 | 216 | 224 | 1.04× [1.02, 1.06] | +3.9% | [+2.2%, +5.5%] | 1.7 pts | 1.1 | yes | yes |
| multi | street | 4096 | churn | ordered | baseline | 8 | 95.6 | 94.8 | 0.99× [0.98, 1.00] | -1.0% | [-2.1%, +0.0%] | 1.3 pts | 1.2 | yes | no |
| multi | street | 4096 | build | ordered | baseline | 8 | 3.70 ms | 3.69 ms | 1.00× [0.98, 1.01] | -0.4% | [-1.6%, +0.9%] | 1.1 pts | 1.9 | yes | no |
| multi | street | 16384 | valuesFor | ordered | baseline | 8 | 73.7 | 73.4 | 1.00× [0.99, 1.01] | -0.2% | [-0.9%, +0.6%] | 1.0 pts | 1.2 | yes | no |
| multi | street | 16384 | valuesBetween | ordered | baseline | 8 | 2835 | 2866 | 1.01× [1.00, 1.01] | +0.6% | [-0.0%, +1.2%] | 0.6 pts | 0.8 | yes | no |
| multi | street | 16384 | prefix | ordered | baseline | 8 | 789 | 810 | 1.03× [1.01, 1.05] | +2.8% | [+0.9%, +4.7%] | 2.1 pts | 1.0 | yes | yes |
| multi | street | 16384 | churn | ordered | baseline | 8 | 134 | 130 | 0.97× [0.96, 0.98] | -2.9% | [-3.7%, -2.1%] | 0.9 pts | 1.2 | yes | yes |
| multi | street | 16384 | build | ordered | baseline | 8 | 19.86 ms | 19.67 ms | 0.99× [0.98, 1.00] | -1.0% | [-1.6%, -0.4%] | 0.6 pts | 1.2 | yes | yes |
| multi | street | 65536 | valuesFor | ordered | baseline | 8 | 99.2 | 98.2 | 1.00× [0.98, 1.02] | -0.3% | [-2.2%, +1.6%] | 1.8 pts | 1.9 | yes | no |
| multi | street | 65536 | valuesBetween | ordered | baseline | 8 | 3210 | 3278 | 1.01× [0.99, 1.03] | +1.2% | [-0.7%, +3.0%] | 1.7 pts | 2.0 | yes | no |
| multi | street | 65536 | prefix | ordered | baseline | 8 | 3867 | 3937 | 1.03× [1.02, 1.04] | +2.9% | [+1.6%, +4.3%] | 1.3 pts | 0.9 | yes | yes |
| multi | street | 65536 | churn | ordered | baseline | 8 | 217 | 212 | 0.98× [0.97, 0.99] | -2.0% | [-2.8%, -1.3%] | 0.7 pts | 0.9 | yes | yes |
| multi | street | 65536 | build | ordered | baseline | 8 | 120.86 ms | 118.75 ms | 0.99× [0.97, 1.00] | -1.2% | [-2.7%, +0.2%] | 1.5 pts | 1.9 | yes | no |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 12 | 31.9 | 31.1 | 0.97× [0.94, 1.01] | -3.0% | [-6.9%, +0.9%] | 4.0 pts | 0.7 | no | no |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 12 | 2509 | 2496 | 0.99× [0.97, 1.01] | -0.8% | [-2.7%, +1.1%] | 3.6 pts | 1.0 | yes | no |
| multi | u64 | 4096 | churn | ordered | baseline | 12 | 50.7 | 49.1 | 0.97× [0.96, 0.98] | -3.2% | [-4.2%, -2.3%] | 1.2 pts | 2.4 | yes | yes |
| multi | u64 | 4096 | build | ordered | baseline | 12 | 5.55 ms | 5.13 ms | 0.93× [0.92, 0.93] | -8.0% | [-9.0%, -7.0%] | 1.1 pts | 1.2 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 46.1 | 45.9 | 0.99× [0.99, 1.00] | -0.6% | [-1.1%, -0.0%] | 0.5 pts | 1.3 | yes | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3444 | 3399 | 0.99× [0.98, 1.00] | -1.4% | [-2.3%, -0.4%] | 0.9 pts | 1.2 | yes | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 67.3 | 65.0 | 0.96× [0.96, 0.97] | -3.8% | [-4.3%, -3.3%] | 0.5 pts | 1.1 | yes | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 26.62 ms | 25.14 ms | 0.94× [0.92, 0.96] | -6.3% | [-8.1%, -4.5%] | 1.7 pts | 2.4 | yes | yes |
| multi | u64 | 65536 | valuesFor | ordered | baseline | 12 | 53.6 | 53.2 | 0.99× [0.97, 1.01] | -0.9% | [-3.0%, +1.3%] | 2.2 pts | 2.9 | no | no |
| multi | u64 | 65536 | valuesBetween | ordered | baseline | 12 | 4078 | 4079 | 0.99× [0.98, 1.01] | -0.7% | [-1.9%, +0.6%] | 1.3 pts | 1.8 | yes | no |
| multi | u64 | 65536 | churn | ordered | baseline | 12 | 133 | 129 | 0.98× [0.96, 0.99] | -2.5% | [-3.7%, -1.4%] | 1.5 pts | 1.1 | yes | yes |
| multi | u64 | 65536 | build | ordered | baseline | 12 | 147.55 ms | 141.17 ms | 0.95× [0.94, 0.96] | -5.0% | [-6.2%, -3.9%] | 1.3 pts | 1.4 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 12 | 120 | 119 | 0.98× [0.95, 1.01] | -2.0% | [-4.8%, +0.8%] | 4.7 pts | 6.3 | no | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 6302 | 6288 | 1.00× [0.99, 1.01] | +0.0% | [-1.3%, +1.3%] | 2.6 pts | 3.8 | yes | no |
| multi | u64 | 262144 | churn | ordered | baseline | 12 | 290 | 286 | 0.94× [0.91, 0.97] | -6.4% | [-10.0%, -2.7%] | 5.7 pts | 1.0 | no | yes |
| multi | url | 4096 | valuesFor | ordered | baseline | 12 | 85.5 | 84.9 | 0.99× [0.98, 1.00] | -1.0% | [-2.2%, +0.3%] | 2.2 pts | 1.4 | yes | no |
| multi | url | 4096 | valuesBetween | ordered | baseline | 12 | 3668 | 3768 | 1.03× [1.00, 1.05] | +2.5% | [+0.2%, +4.8%] | 2.4 pts | 1.0 | no | yes |
| multi | url | 4096 | prefix | ordered | baseline | 12 | 177 | 180 | 1.02× [1.01, 1.04] | +2.3% | [+1.2%, +3.5%] | 1.5 pts | 1.4 | yes | yes |
| multi | url | 4096 | churn | ordered | baseline | 12 | 127 | 123 | 0.97× [0.97, 0.98] | -2.8% | [-3.2%, -2.4%] | 0.9 pts | 1.1 | yes | yes |
| multi | url | 4096 | build | ordered | baseline | 12 | 11.86 ms | 11.48 ms | 0.97× [0.97, 0.98] | -3.0% | [-3.5%, -2.5%] | 0.8 pts | 1.7 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 109 | 108 | 0.99× [0.99, 1.00] | -1.0% | [-1.5%, -0.5%] | 0.5 pts | 0.7 | yes | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4336 | 4381 | 1.01× [0.99, 1.03] | +0.9% | [-0.6%, +2.5%] | 1.5 pts | 2.7 | yes | no |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 277 | 280 | 1.01× [1.00, 1.02] | +1.0% | [-0.1%, +2.1%] | 1.1 pts | 1.0 | yes | no |
| multi | url | 16384 | churn | ordered | baseline | 6 | 180 | 174 | 0.97× [0.96, 0.98] | -3.1% | [-3.8%, -2.3%] | 0.7 pts | 1.1 | yes | yes |
| multi | url | 16384 | build | ordered | baseline | 6 | 60.26 ms | 58.53 ms | 0.97× [0.97, 0.97] | -3.0% | [-3.2%, -2.8%] | 0.2 pts | 0.4 | yes | yes |
| multi | url | 65536 | valuesFor | ordered | baseline | 12 | 191 | 188 | 1.00× [0.98, 1.01] | -0.4% | [-1.5%, +0.7%] | 1.6 pts | 1.4 | yes | no |
| multi | url | 65536 | valuesBetween | ordered | baseline | 12 | 5750 | 5687 | 1.00× [0.98, 1.01] | -0.3% | [-1.8%, +1.1%] | 1.9 pts | 1.2 | yes | no |
| multi | url | 65536 | prefix | ordered | baseline | 12 | 675 | 698 | 1.03× [0.99, 1.07] | +2.7% | [-1.0%, +6.5%] | 5.6 pts | 3.8 | no | no |
| multi | url | 65536 | churn | ordered | baseline | 12 | 358 | 351 | 0.97× [0.95, 1.00] | -2.7% | [-4.9%, -0.4%] | 2.7 pts | 1.2 | no | yes |
| multi | url | 65536 | build | ordered | baseline | 12 | 410.60 ms | 399.27 ms | 0.98× [0.97, 0.98] | -2.3% | [-2.9%, -1.6%] | 0.8 pts | 0.7 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 12 | 361 | 356 | 0.99× [0.96, 1.02] | -0.8% | [-3.6%, +2.0%] | 3.4 pts | 4.0 | no | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 12 | 10.4 µs | 10.4 µs | 1.00× [1.00, 1.00] | -0.1% | [-0.4%, +0.2%] | 1.1 pts | 1.3 | yes | no |
| multi | url | 262144 | prefix | ordered | baseline | 12 | 2325 | 2358 | 1.00× [1.00, 1.01] | +0.1% | [-0.4%, +0.5%] | 0.9 pts | 0.9 | yes | no |
| multi | url | 262144 | churn | ordered | baseline | 12 | 632 | 619 | 0.96× [0.92, 0.99] | -4.7% | [-8.4%, -0.9%] | 7.3 pts | 1.2 | no | yes |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 12 | 52.7 | 52.3 | 1.00× [0.98, 1.01] | -0.2% | [-1.7%, +1.3%] | 2.1 pts | 0.9 | yes | no |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 3102 | 3118 | 1.01× [0.98, 1.04] | +0.9% | [-1.9%, +3.7%] | 3.6 pts | 1.3 | no | no |
| multi | uuid | 4096 | prefix | ordered | baseline | 12 | 89.4 | 89.0 | 1.00× [0.99, 1.01] | +0.1% | [-0.6%, +0.7%] | 1.4 pts | 1.2 | yes | no |
| multi | uuid | 4096 | churn | ordered | baseline | 12 | 78.9 | 76.7 | 0.98× [0.96, 0.99] | -2.4% | [-3.7%, -1.2%] | 1.3 pts | 2.0 | yes | yes |
| multi | uuid | 4096 | build | ordered | baseline | 12 | 7.72 ms | 7.60 ms | 0.98× [0.98, 0.99] | -1.6% | [-2.0%, -1.2%] | 0.9 pts | 1.5 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 63.9 | 63.2 | 0.99× [0.99, 0.99] | -1.0% | [-1.3%, -0.6%] | 0.3 pts | 0.8 | yes | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3532 | 3545 | 1.00× [1.00, 1.01] | +0.3% | [-0.4%, +1.0%] | 0.7 pts | 1.5 | yes | no |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 107 | 107 | 1.00× [0.99, 1.01] | +0.2% | [-0.9%, +1.3%] | 1.0 pts | 1.9 | yes | no |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 102 | 100 | 0.98× [0.97, 0.99] | -2.2% | [-2.9%, -1.4%] | 0.7 pts | 1.2 | yes | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 36.31 ms | 35.85 ms | 0.99× [0.98, 0.99] | -1.5% | [-1.9%, -1.0%] | 0.4 pts | 0.9 | yes | yes |
| multi | uuid | 65536 | valuesFor | ordered | baseline | 12 | 97.1 | 94.2 | 0.98× [0.97, 1.00] | -1.7% | [-3.6%, +0.1%] | 2.9 pts | 2.9 | yes | no |
| multi | uuid | 65536 | valuesBetween | ordered | baseline | 12 | 4335 | 4359 | 1.00× [0.98, 1.01] | -0.2% | [-1.6%, +1.1%] | 1.8 pts | 1.7 | yes | no |
| multi | uuid | 65536 | prefix | ordered | baseline | 12 | 183 | 181 | 1.00× [0.97, 1.02] | -0.5% | [-2.6%, +1.7%] | 3.2 pts | 4.1 | no | no |
| multi | uuid | 65536 | churn | ordered | baseline | 12 | 212 | 207 | 0.96× [0.95, 0.98] | -3.8% | [-5.2%, -2.4%] | 2.2 pts | 0.9 | yes | yes |
| multi | uuid | 65536 | build | ordered | baseline | 12 | 236.63 ms | 227.58 ms | 0.97× [0.96, 0.97] | -3.4% | [-4.2%, -2.6%] | 0.9 pts | 0.9 | yes | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 12 | 227 | 221 | 0.98× [0.95, 1.00] | -2.5% | [-4.8%, -0.1%] | 3.0 pts | 4.3 | no | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 8049 | 8050 | 1.00× [0.98, 1.02] | +0.2% | [-1.7%, +2.0%] | 2.1 pts | 2.8 | yes | no |
| multi | uuid | 262144 | prefix | ordered | baseline | 12 | 535 | 529 | 0.99× [0.98, 1.00] | -1.2% | [-2.6%, +0.2%] | 1.8 pts | 3.6 | yes | no |
| multi | uuid | 262144 | churn | ordered | baseline | 12 | 444 | 420 | 0.92× [0.88, 0.96] | -9.2% | [-13.8%, -4.6%] | 5.2 pts | 0.7 | no | yes |
| unique | email | 4096 | valuesFor | ordered | baseline | 12 | 22.5 | 21.9 | 0.97× [0.95, 0.99] | -2.9% | [-4.9%, -1.0%] | 2.4 pts | 1.0 | yes | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 12 | 1225 | 1261 | 1.03× [1.02, 1.03] | +2.7% | [+2.3%, +3.1%] | 1.1 pts | 1.0 | yes | yes |
| unique | email | 4096 | prefix | ordered | baseline | 12 | 58.2 | 57.2 | 0.99× [0.98, 0.99] | -1.4% | [-1.8%, -1.0%] | 0.8 pts | 1.0 | yes | yes |
| unique | email | 4096 | churn | ordered | baseline | 12 | 71.6 | 70.8 | 0.99× [0.98, 1.00] | -1.3% | [-2.2%, -0.4%] | 0.9 pts | 1.4 | yes | yes |
| unique | email | 4096 | build | ordered | baseline | 12 | 1.04 ms | 1.04 ms | 0.99× [0.99, 1.00] | -0.6% | [-1.2%, +0.0%] | 1.0 pts | 1.1 | yes | no |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 38.0 | 37.4 | 0.99× [0.98, 1.00] | -1.1% | [-2.1%, -0.2%] | 0.9 pts | 1.3 | yes | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 1596 | 1596 | 0.99× [0.98, 1.00] | -0.7% | [-1.7%, +0.3%] | 1.0 pts | 1.0 | yes | no |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 72.3 | 71.8 | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.3%] | 0.6 pts | 1.5 | yes | no |
| unique | email | 16384 | churn | ordered | baseline | 6 | 89.4 | 88.4 | 0.98× [0.98, 0.99] | -1.5% | [-1.9%, -1.2%] | 0.4 pts | 0.4 | yes | yes |
| unique | email | 16384 | build | ordered | baseline | 6 | 5.11 ms | 5.05 ms | 0.99× [0.98, 0.99] | -1.0% | [-1.5%, -0.5%] | 0.5 pts | 0.9 | yes | yes |
| unique | email | 65536 | valuesFor | ordered | baseline | 6 | 45.5 | 44.1 | 0.97× [0.96, 0.98] | -3.3% | [-4.1%, -2.5%] | 0.7 pts | 1.4 | yes | yes |
| unique | email | 65536 | valuesBetween | ordered | baseline | 6 | 1535 | 1543 | 1.00× [0.99, 1.01] | +0.1% | [-0.9%, +1.1%] | 0.9 pts | 1.9 | yes | no |
| unique | email | 65536 | prefix | ordered | baseline | 6 | 90.3 | 89.4 | 0.99× [0.98, 1.00] | -0.8% | [-1.6%, -0.1%] | 0.7 pts | 1.0 | yes | yes |
| unique | email | 65536 | churn | ordered | baseline | 6 | 130 | 129 | 0.97× [0.96, 0.99] | -2.7% | [-3.9%, -1.5%] | 1.1 pts | 0.8 | yes | yes |
| unique | email | 65536 | build | ordered | baseline | 6 | 27.15 ms | 26.43 ms | 0.98× [0.96, 0.99] | -2.5% | [-3.8%, -1.1%] | 1.3 pts | 0.8 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 12 | 133 | 130 | 0.99× [0.97, 1.00] | -1.3% | [-2.7%, +0.2%] | 1.8 pts | 1.8 | yes | no |
| unique | email | 262144 | valuesBetween | ordered | baseline | 12 | 3183 | 3188 | 1.02× [1.01, 1.03] | +2.0% | [+1.4%, +2.5%] | 1.5 pts | 1.0 | yes | yes |
| unique | email | 262144 | prefix | ordered | baseline | 12 | 176 | 180 | 1.01× [1.00, 1.03] | +1.3% | [+0.0%, +2.5%] | 1.9 pts | 1.8 | yes | yes |
| unique | email | 262144 | churn | ordered | baseline | 12 | 261 | 263 | 0.99× [0.97, 1.00] | -1.5% | [-3.3%, +0.3%] | 2.6 pts | 1.1 | yes | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 12 | 78.2 | 77.4 | 0.99× [0.98, 1.00] | -1.0% | [-2.4%, +0.3%] | 2.0 pts | 1.3 | yes | no |
| unique | path | 4096 | valuesBetween | ordered | baseline | 12 | 2007 | 2043 | 1.01× [0.98, 1.03] | +0.7% | [-1.6%, +3.0%] | 2.5 pts | 1.3 | no | no |
| unique | path | 4096 | prefix | ordered | baseline | 12 | 233 | 235 | 1.01× [1.00, 1.01] | +0.7% | [-0.0%, +1.5%] | 0.9 pts | 0.8 | yes | no |
| unique | path | 4096 | churn | ordered | baseline | 12 | 201 | 194 | 0.97× [0.96, 0.98] | -3.4% | [-4.2%, -2.5%] | 1.0 pts | 1.1 | yes | yes |
| unique | path | 4096 | build | ordered | baseline | 12 | 2.71 ms | 2.62 ms | 0.97× [0.97, 0.97] | -2.9% | [-3.3%, -2.6%] | 0.9 pts | 0.9 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 12 | 112 | 112 | 1.01× [0.98, 1.03] | +0.5% | [-2.0%, +3.0%] | 3.9 pts | 4.0 | no | no |
| unique | path | 16384 | valuesBetween | ordered | baseline | 12 | 2537 | 2583 | 1.03× [1.02, 1.04] | +3.1% | [+2.4%, +3.8%] | 5.7 pts | 5.6 | yes | yes |
| unique | path | 16384 | prefix | ordered | baseline | 12 | 483 | 487 | 1.03× [1.01, 1.06] | +3.4% | [+1.2%, +5.5%] | 2.7 pts | 1.3 | no | yes |
| unique | path | 16384 | churn | ordered | baseline | 12 | 253 | 248 | 0.97× [0.97, 0.98] | -2.7% | [-3.1%, -2.3%] | 1.1 pts | 1.3 | yes | yes |
| unique | path | 16384 | build | ordered | baseline | 12 | 13.08 ms | 12.71 ms | 0.97× [0.97, 0.98] | -2.7% | [-3.1%, -2.4%] | 0.5 pts | 0.8 | yes | yes |
| unique | path | 65536 | valuesFor | ordered | baseline | 8 | 148 | 149 | 0.99× [0.98, 1.01] | -0.7% | [-2.3%, +0.8%] | 1.5 pts | 1.5 | yes | no |
| unique | path | 65536 | valuesBetween | ordered | baseline | 8 | 2808 | 2847 | 1.01× [1.00, 1.02] | +1.0% | [-0.0%, +2.0%] | 1.0 pts | 1.5 | yes | no |
| unique | path | 65536 | prefix | ordered | baseline | 8 | 1372 | 1420 | 1.02× [1.00, 1.04] | +2.2% | [+0.4%, +4.0%] | 1.7 pts | 1.0 | yes | no |
| unique | path | 65536 | churn | ordered | baseline | 8 | 377 | 364 | 0.97× [0.96, 0.99] | -3.0% | [-4.6%, -1.4%] | 1.7 pts | 1.6 | yes | yes |
| unique | path | 65536 | build | ordered | baseline | 8 | 72.85 ms | 71.18 ms | 0.98× [0.97, 0.98] | -2.3% | [-2.8%, -1.8%] | 0.7 pts | 0.9 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 12 | 316 | 316 | 1.00× [0.98, 1.02] | +0.1% | [-1.6%, +1.9%] | 2.6 pts | 2.5 | yes | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 12 | 5706 | 5813 | 1.02× [1.00, 1.03] | +1.5% | [-0.1%, +3.0%] | 2.2 pts | 1.6 | yes | no |
| unique | path | 262144 | prefix | ordered | baseline | 12 | 4787 | 4994 | 1.02× [1.00, 1.04] | +2.2% | [+0.5%, +4.0%] | 4.4 pts | 1.3 | yes | no |
| unique | path | 262144 | churn | ordered | baseline | 12 | 625 | 617 | 0.96× [0.94, 0.99] | -3.7% | [-6.8%, -0.6%] | 3.2 pts | 0.8 | no | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 12 | 32.9 | 31.5 | 0.97× [0.94, 0.99] | -3.4% | [-5.9%, -0.8%] | 3.7 pts | 0.8 | no | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 12 | 1600 | 1636 | 1.02× [1.01, 1.02] | +1.5% | [+0.7%, +2.4%] | 1.7 pts | 1.1 | yes | yes |
| unique | str | 4096 | prefix | ordered | baseline | 12 | 3092 | 3181 | 1.03× [1.02, 1.03] | +2.5% | [+1.6%, +3.3%] | 1.0 pts | 1.1 | yes | yes |
| unique | str | 4096 | churn | ordered | baseline | 12 | 95.5 | 95.1 | 0.99× [0.99, 1.00] | -0.9% | [-1.5%, -0.4%] | 1.2 pts | 1.3 | yes | yes |
| unique | str | 4096 | build | ordered | baseline | 12 | 1.38 ms | 1.37 ms | 1.00× [0.99, 1.01] | -0.4% | [-1.4%, +0.6%] | 1.0 pts | 0.6 | yes | no |
| unique | str | 16384 | valuesFor | ordered | baseline | 6 | 49.3 | 48.7 | 0.99× [0.98, 1.01] | -0.6% | [-2.5%, +1.2%] | 1.8 pts | 2.1 | yes | no |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 1612 | 1637 | 1.01× [1.00, 1.02] | +1.2% | [+0.4%, +2.1%] | 0.8 pts | 0.8 | yes | no |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 13.3 µs | 13.7 µs | 1.03× [1.02, 1.04] | +3.1% | [+2.3%, +3.8%] | 0.7 pts | 0.6 | yes | yes |
| unique | str | 16384 | churn | ordered | baseline | 6 | 114 | 112 | 0.98× [0.97, 0.99] | -2.0% | [-2.8%, -1.3%] | 0.7 pts | 0.9 | yes | yes |
| unique | str | 16384 | build | ordered | baseline | 6 | 6.51 ms | 6.36 ms | 0.99× [0.98, 0.99] | -1.5% | [-2.2%, -0.9%] | 0.6 pts | 0.8 | yes | yes |
| unique | str | 65536 | valuesFor | ordered | baseline | 8 | 64.2 | 63.8 | 0.99× [0.98, 1.00] | -0.7% | [-1.7%, +0.3%] | 1.0 pts | 1.8 | yes | no |
| unique | str | 65536 | valuesBetween | ordered | baseline | 8 | 1808 | 1800 | 0.99× [0.97, 1.01] | -1.0% | [-2.6%, +0.7%] | 1.5 pts | 2.7 | yes | no |
| unique | str | 65536 | prefix | ordered | baseline | 8 | 62.9 µs | 62.8 µs | 0.99× [0.98, 1.01] | -0.8% | [-2.5%, +1.0%] | 1.7 pts | 3.5 | yes | no |
| unique | str | 65536 | churn | ordered | baseline | 8 | 165 | 160 | 0.99× [0.98, 1.00] | -0.9% | [-1.7%, -0.2%] | 0.7 pts | 0.7 | yes | no |
| unique | str | 65536 | build | ordered | baseline | 8 | 32.96 ms | 32.31 ms | 0.98× [0.97, 0.99] | -2.1% | [-3.2%, -1.1%] | 1.0 pts | 1.2 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 12 | 136 | 135 | 0.98× [0.97, 1.00] | -1.8% | [-3.3%, -0.4%] | 2.5 pts | 1.5 | yes | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 12 | 2716 | 2678 | 0.98× [0.95, 1.02] | -1.9% | [-5.6%, +1.7%] | 4.3 pts | 2.2 | no | no |
| unique | str | 262144 | prefix | ordered | baseline | 12 | 401.8 µs | 380.2 µs | 0.98× [0.93, 1.03] | -2.5% | [-7.7%, +2.6%] | 5.9 pts | 1.9 | no | no |
| unique | str | 262144 | churn | ordered | baseline | 12 | 313 | 309 | 0.98× [0.97, 1.00] | -1.7% | [-3.4%, -0.1%] | 2.2 pts | 0.7 | yes | yes |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 39.9 | 39.0 | 0.97× [0.95, 0.98] | -3.3% | [-4.9%, -1.7%] | 1.5 pts | 0.5 | yes | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 1744 | 1785 | 1.03× [1.01, 1.05] | +2.7% | [+0.9%, +4.5%] | 1.7 pts | 0.9 | yes | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 192 | 197 | 1.02× [1.01, 1.04] | +2.3% | [+0.6%, +4.0%] | 1.6 pts | 1.2 | yes | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 115 | 115 | 1.00× [1.00, 1.01] | +0.3% | [-0.5%, +1.1%] | 0.8 pts | 0.5 | yes | no |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.62 ms | 1.63 ms | 1.00× [1.00, 1.01] | +0.4% | [-0.4%, +1.3%] | 0.8 pts | 1.6 | yes | no |
| unique | street | 16384 | valuesFor | ordered | baseline | 10 | 66.8 | 66.2 | 0.98× [0.98, 0.99] | -1.7% | [-2.5%, -0.9%] | 2.3 pts | 2.3 | yes | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 10 | 2158 | 2190 | 1.01× [1.00, 1.02] | +1.2% | [+0.3%, +2.1%] | 1.7 pts | 1.5 | yes | yes |
| unique | street | 16384 | prefix | ordered | baseline | 10 | 613 | 639 | 1.03× [1.01, 1.05] | +3.2% | [+1.3%, +5.2%] | 2.5 pts | 1.4 | yes | yes |
| unique | street | 16384 | churn | ordered | baseline | 10 | 150 | 149 | 0.99× [0.99, 1.00] | -0.6% | [-0.8%, -0.4%] | 0.6 pts | 0.8 | yes | yes |
| unique | street | 16384 | build | ordered | baseline | 10 | 8.09 ms | 8.09 ms | 0.99× [0.99, 1.00] | -0.7% | [-1.5%, +0.1%] | 0.9 pts | 1.7 | yes | no |
| unique | street | 65536 | valuesFor | ordered | baseline | 12 | 89.2 | 87.9 | 0.98× [0.94, 1.02] | -2.3% | [-6.4%, +1.8%] | 4.2 pts | 4.7 | no | no |
| unique | street | 65536 | valuesBetween | ordered | baseline | 12 | 2433 | 2456 | 1.00× [0.97, 1.04] | +0.3% | [-2.8%, +3.5%] | 2.9 pts | 4.1 | no | no |
| unique | street | 65536 | prefix | ordered | baseline | 12 | 2918 | 2981 | 1.02× [1.00, 1.04] | +2.0% | [-0.1%, +4.2%] | 2.2 pts | 1.5 | no | no |
| unique | street | 65536 | churn | ordered | baseline | 12 | 224 | 222 | 0.99× [0.98, 0.99] | -1.1% | [-1.7%, -0.5%] | 0.9 pts | 1.4 | yes | yes |
| unique | street | 65536 | build | ordered | baseline | 12 | 42.90 ms | 42.50 ms | 0.99× [0.98, 0.99] | -1.4% | [-2.3%, -0.5%] | 0.9 pts | 1.0 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 17.8 | 12.7 | 0.71× [0.71, 0.72] | -40.4% | [-41.4%, -39.3%] | 1.0 pts | 0.8 | yes | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 238 | 981 | 4.12× [4.10, 4.15] | +75.7% | [+75.6%, +75.9%] | 0.1 pts | 1.3 | yes | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 36.6 | 46.1 | 1.27× [1.25, 1.28] | +21.0% | [+20.3%, +21.7%] | 0.7 pts | 1.8 | yes | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 579.9 µs | 770.9 µs | 1.33× [1.32, 1.34] | +24.7% | [+24.2%, +25.1%] | 0.4 pts | 0.9 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 6 | 26.4 | 22.9 | 0.87× [0.85, 0.88] | -15.5% | [-17.1%, -14.0%] | 1.5 pts | 3.2 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 281 | 1462 | 5.19× [5.16, 5.23] | +80.7% | [+80.6%, +80.9%] | 0.1 pts | 1.1 | yes | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 6 | 54.4 | 52.8 | 0.97× [0.96, 0.98] | -3.3% | [-4.4%, -2.2%] | 1.0 pts | 1.7 | yes | yes |
| unique | u64 | 16384 | build | ordered | baseline | 6 | 3.38 ms | 3.57 ms | 1.06× [1.05, 1.07] | +5.2% | [+4.4%, +6.1%] | 0.8 pts | 1.0 | yes | yes |
| unique | u64 | 65536 | valuesFor | ordered | baseline | 10 | 31.4 | 26.0 | 0.83× [0.82, 0.85] | -20.0% | [-22.0%, -18.0%] | 2.0 pts | 4.8 | yes | yes |
| unique | u64 | 65536 | valuesBetween | ordered | baseline | 10 | 279 | 1983 | 7.13× [7.10, 7.17] | +86.0% | [+85.9%, +86.1%] | 0.1 pts | 1.4 | yes | yes |
| unique | u64 | 65536 | churn | ordered | baseline | 10 | 69.5 | 80.3 | 1.17× [1.15, 1.18] | +14.2% | [+12.9%, +15.5%] | 1.9 pts | 2.1 | yes | yes |
| unique | u64 | 65536 | build | ordered | baseline | 10 | 16.87 ms | 18.29 ms | 1.07× [1.05, 1.09] | +6.5% | [+5.0%, +8.1%] | 2.0 pts | 2.5 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 12 | 37.1 | 42.6 | 1.08× [0.99, 1.20] | +7.8% | [-1.3%, +16.9%] | 13.1 pts | 9.7 | no | no |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 306 | 1399 | 4.54× [4.38, 4.71] | +78.0% | [+77.2%, +78.8%] | 1.3 pts | 2.4 | yes | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 12 | 115 | 181 | 1.56× [1.47, 1.66] | +36.0% | [+32.2%, +39.9%] | 4.7 pts | 1.3 | no | yes |
| unique | url | 4096 | valuesFor | ordered | baseline | 10 | 60.9 | 59.0 | 0.98× [0.97, 0.99] | -2.2% | [-3.5%, -0.9%] | 1.8 pts | 0.9 | yes | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 10 | 1883 | 1939 | 1.03× [1.01, 1.04] | +2.7% | [+1.2%, +4.1%] | 1.7 pts | 0.8 | yes | yes |
| unique | url | 4096 | prefix | ordered | baseline | 10 | 142 | 146 | 1.02× [1.01, 1.04] | +2.2% | [+0.9%, +3.4%] | 1.2 pts | 1.1 | yes | yes |
| unique | url | 4096 | churn | ordered | baseline | 10 | 160 | 155 | 0.97× [0.95, 0.98] | -3.5% | [-4.7%, -2.3%] | 1.3 pts | 1.3 | yes | yes |
| unique | url | 4096 | build | ordered | baseline | 10 | 2.23 ms | 2.17 ms | 0.98× [0.96, 0.99] | -2.6% | [-4.5%, -0.6%] | 1.9 pts | 1.3 | yes | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 88.4 | 87.2 | 0.99× [0.98, 0.99] | -1.4% | [-2.2%, -0.7%] | 0.7 pts | 0.8 | yes | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 2363 | 2391 | 1.01× [1.01, 1.02] | +1.2% | [+0.5%, +1.9%] | 0.6 pts | 0.8 | yes | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 211 | 215 | 1.02× [1.00, 1.04] | +2.0% | [+0.3%, +3.7%] | 1.6 pts | 1.6 | yes | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 201 | 196 | 0.97× [0.97, 0.98] | -2.7% | [-3.6%, -1.9%] | 0.8 pts | 1.0 | yes | yes |
| unique | url | 16384 | build | ordered | baseline | 6 | 10.57 ms | 10.31 ms | 0.97× [0.97, 0.98] | -2.8% | [-3.5%, -2.1%] | 0.7 pts | 1.1 | yes | yes |
| unique | url | 65536 | valuesFor | ordered | baseline | 8 | 120 | 120 | 0.99× [0.98, 1.01] | -0.9% | [-2.5%, +0.7%] | 1.5 pts | 1.2 | yes | no |
| unique | url | 65536 | valuesBetween | ordered | baseline | 8 | 2640 | 2662 | 1.00× [0.99, 1.02] | +0.5% | [-1.2%, +2.2%] | 1.6 pts | 1.8 | yes | no |
| unique | url | 65536 | prefix | ordered | baseline | 8 | 432 | 457 | 1.04× [1.02, 1.06] | +3.6% | [+1.6%, +5.6%] | 1.8 pts | 0.9 | yes | yes |
| unique | url | 65536 | churn | ordered | baseline | 8 | 305 | 300 | 0.97× [0.96, 0.99] | -2.7% | [-4.2%, -1.2%] | 1.4 pts | 0.7 | yes | yes |
| unique | url | 65536 | build | ordered | baseline | 8 | 61.99 ms | 60.08 ms | 0.97× [0.96, 0.97] | -3.3% | [-3.9%, -2.7%] | 0.6 pts | 0.9 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 12 | 314 | 308 | 0.99× [0.96, 1.02] | -1.1% | [-4.4%, +2.3%] | 4.0 pts | 3.7 | no | no |
| unique | url | 262144 | valuesBetween | ordered | baseline | 12 | 6217 | 6268 | 0.99× [0.98, 1.01] | -0.6% | [-2.4%, +1.2%] | 2.2 pts | 2.5 | yes | no |
| unique | url | 262144 | prefix | ordered | baseline | 12 | 1284 | 1302 | 1.00× [0.98, 1.03] | +0.4% | [-2.2%, +3.0%] | 3.0 pts | 2.6 | no | no |
| unique | url | 262144 | churn | ordered | baseline | 12 | 553 | 546 | 0.98× [0.96, 1.00] | -1.9% | [-3.8%, +0.1%] | 2.3 pts | 0.9 | yes | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 8 | 26.9 | 26.3 | 0.99× [0.97, 1.00] | -1.5% | [-2.8%, -0.1%] | 1.7 pts | 0.5 | yes | no |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 8 | 1411 | 1447 | 1.01× [0.99, 1.03] | +1.1% | [-0.7%, +2.8%] | 1.7 pts | 1.1 | yes | no |
| unique | uuid | 4096 | prefix | ordered | baseline | 8 | 69.2 | 68.9 | 1.00× [0.99, 1.00] | -0.5% | [-1.2%, +0.2%] | 0.7 pts | 1.0 | yes | no |
| unique | uuid | 4096 | churn | ordered | baseline | 8 | 83.2 | 81.9 | 0.99× [0.98, 1.00] | -1.3% | [-2.4%, -0.2%] | 1.2 pts | 1.3 | yes | yes |
| unique | uuid | 4096 | build | ordered | baseline | 8 | 1.16 ms | 1.15 ms | 0.99× [0.98, 1.00] | -1.2% | [-2.0%, -0.4%] | 1.1 pts | 1.1 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 6 | 42.6 | 41.7 | 0.98× [0.97, 0.99] | -2.2% | [-2.8%, -1.5%] | 0.7 pts | 1.0 | yes | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 1567 | 1591 | 1.01× [1.01, 1.01] | +1.0% | [+0.6%, +1.4%] | 0.4 pts | 0.5 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 84.2 | 84.2 | 1.00× [0.99, 1.00] | -0.1% | [-0.6%, +0.3%] | 0.4 pts | 1.2 | yes | no |
| unique | uuid | 16384 | churn | ordered | baseline | 6 | 98.6 | 96.7 | 0.98× [0.97, 0.99] | -1.7% | [-2.6%, -0.8%] | 0.8 pts | 0.9 | yes | yes |
| unique | uuid | 16384 | build | ordered | baseline | 6 | 5.57 ms | 5.51 ms | 0.99× [0.99, 1.00] | -0.8% | [-1.2%, -0.3%] | 0.5 pts | 0.7 | yes | yes |
| unique | uuid | 65536 | valuesFor | ordered | baseline | 10 | 54.7 | 53.7 | 0.98× [0.97, 0.98] | -2.5% | [-3.3%, -1.7%] | 1.7 pts | 3.1 | yes | yes |
| unique | uuid | 65536 | valuesBetween | ordered | baseline | 10 | 1876 | 1874 | 1.00× [0.99, 1.00] | -0.1% | [-0.5%, +0.3%] | 0.8 pts | 1.5 | yes | no |
| unique | uuid | 65536 | prefix | ordered | baseline | 10 | 107 | 108 | 1.00× [0.99, 1.02] | +0.5% | [-0.6%, +1.6%] | 2.0 pts | 3.6 | yes | no |
| unique | uuid | 65536 | churn | ordered | baseline | 10 | 158 | 157 | 0.99× [0.97, 1.00] | -1.4% | [-3.2%, +0.5%] | 1.9 pts | 1.0 | yes | no |
| unique | uuid | 65536 | build | ordered | baseline | 10 | 33.73 ms | 32.85 ms | 0.98× [0.97, 1.00] | -1.7% | [-3.3%, -0.0%] | 2.0 pts | 1.3 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 12 | 157 | 152 | 0.96× [0.95, 0.97] | -4.1% | [-5.2%, -2.9%] | 1.8 pts | 1.7 | yes | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 3884 | 3878 | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.3%] | 1.4 pts | 0.9 | yes | no |
| unique | uuid | 262144 | prefix | ordered | baseline | 12 | 317 | 313 | 0.99× [0.98, 1.00] | -0.7% | [-1.7%, +0.3%] | 1.5 pts | 2.1 | yes | no |
| unique | uuid | 262144 | churn | ordered | baseline | 12 | 315 | 311 | 0.96× [0.93, 0.98] | -4.6% | [-7.5%, -1.7%] | 3.0 pts | 0.9 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the pooled difference of 1.24% does not clear the 1.36% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.41%, 2.90%] includes zero
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.34% does not clear the 1.41% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.67%, 3.35%] includes zero
- multi email n=4096 prefix: ordered vs baseline: the pooled interval [-2.29%, 0.14%] includes zero
- multi email n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.34% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.13%, 0.44%] includes zero
- multi email n=16384 prefix: ordered vs baseline: the pooled difference of 0.19% does not clear the 0.49% noise floor, the bound on what the harness reports between identical code in every process
- multi email n=16384 prefix: ordered vs baseline: the pooled interval [-0.30%, 0.68%] includes zero
- multi email n=65536 valuesFor: ordered vs baseline: the pooled interval [-3.91%, 0.40%] includes zero
- multi email n=65536 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=65536 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=65536 valuesBetween: ordered vs baseline: the pooled interval [-2.02%, 0.71%] includes zero
- multi email n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=65536 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=65536 prefix: ordered vs baseline: the pooled interval [-1.99%, 0.18%] includes zero
- multi email n=65536 prefix: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.43%, 0.41%] includes zero
- multi email n=262144 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=262144 prefix: ordered vs baseline: the pooled interval [-1.22%, 2.45%] includes zero
- multi email n=262144 prefix: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 prefix: ordered vs baseline: 6 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.97% does not clear the 1.66% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.66%, 2.60%] includes zero
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of 0.68% does not clear the 0.85% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=4096 prefix: ordered vs baseline: the pooled interval [-0.82%, 2.19%] includes zero
- multi path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.53% does not clear the 0.65% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.98%, 2.05%] includes zero
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.16%, 1.93%] includes zero
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of 0.98% does not clear the 2.62% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-0.72%, 2.68%] includes zero
- multi path n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.08% does not clear the 0.54% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.32%, 1.15%] includes zero
- multi path n=65536 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=65536 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=65536 valuesBetween: ordered vs baseline: the pooled interval [-0.20%, 1.43%] includes zero
- multi path n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=65536 valuesBetween: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=65536 prefix: ordered vs baseline: the pooled difference of -0.22% does not clear the 5.02% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=65536 prefix: ordered vs baseline: the pooled interval [-1.49%, 1.05%] includes zero
- multi path n=65536 churn: ordered vs baseline: the pooled interval [-6.05%, 0.98%] includes zero
- multi path n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.64% does not clear the 0.83% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.63%, 0.34%] includes zero
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.08% does not clear the 0.58% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.59%, 0.76%] includes zero
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of -0.07% does not clear the 2.99% noise floor, the bound on what the harness reports between identical code in every process
- multi path n=262144 prefix: ordered vs baseline: the pooled interval [-2.07%, 1.93%] includes zero
- multi str n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.31% does not clear the 1.42% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.11%, 0.49%] includes zero
- multi str n=4096 prefix: ordered vs baseline: the pooled interval [-0.18%, 8.14%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.33%, 0.30%] includes zero
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.24% does not clear the 0.75% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.05%, 0.57%] includes zero
- multi str n=16384 prefix: ordered vs baseline: the pooled interval [-0.50%, 1.19%] includes zero
- multi str n=65536 valuesFor: ordered vs baseline: the pooled interval [-2.98%, 0.91%] includes zero
- multi str n=65536 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=65536 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=65536 valuesBetween: ordered vs baseline: the pooled interval [-2.24%, 0.81%] includes zero
- multi str n=65536 prefix: ordered vs baseline: the pooled interval [-2.77%, 0.94%] includes zero
- multi str n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.49% does not clear the 0.64% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.24%, 1.25%] includes zero
- multi str n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=262144 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.74%, 0.66%] includes zero
- multi str n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=262144 prefix: ordered vs baseline: the pooled difference of -0.41% does not clear the 0.91% noise floor, the bound on what the harness reports between identical code in every process
- multi str n=262144 prefix: ordered vs baseline: the pooled interval [-1.28%, 0.46%] includes zero
- multi str n=262144 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.50% does not clear the 1.53% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.29%, 1.30%] includes zero
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.76%, 2.58%] includes zero
- multi street n=4096 churn: ordered vs baseline: the pooled interval [-2.09%, 0.02%] includes zero
- multi street n=4096 build: ordered vs baseline: the pooled difference of -0.36% does not clear the 0.71% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=4096 build: ordered vs baseline: the pooled interval [-1.57%, 0.85%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.19% does not clear the 0.47% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.95%, 0.56%] includes zero
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.02%, 1.18%] includes zero
- multi street n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.26% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- multi street n=65536 valuesFor: ordered vs baseline: the pooled interval [-2.16%, 1.64%] includes zero
- multi street n=65536 valuesBetween: ordered vs baseline: the pooled interval [-0.70%, 3.04%] includes zero
- multi street n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=65536 build: ordered vs baseline: the pooled interval [-2.65%, 0.17%] includes zero
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-6.92%, 0.91%] includes zero
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.82% does not clear the 1.09% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.74%, 1.11%] includes zero
- multi u64 n=4096 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=16384 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=65536 valuesFor: ordered vs baseline: the pooled interval [-3.03%, 1.27%] includes zero
- multi u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.93%, 0.61%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-4.76%, 0.77%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 6.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.00% does not clear the 0.25% noise floor, the bound on what the harness reports between identical code in every process
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.32%, 1.33%] includes zero
- multi u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesBetween: ordered vs baseline: 5 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.22%, 0.28%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.65%, 2.52%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=16384 prefix: ordered vs baseline: the pooled difference of 0.99% does not clear the 1.06% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=16384 prefix: ordered vs baseline: the pooled interval [-0.13%, 2.11%] includes zero
- multi url n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.43% does not clear the 0.81% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.53%, 0.67%] includes zero
- multi url n=65536 valuesBetween: ordered vs baseline: the pooled difference of -0.35% does not clear the 0.45% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.77%, 1.07%] includes zero
- multi url n=65536 prefix: ordered vs baseline: the pooled interval [-1.04%, 6.50%] includes zero
- multi url n=65536 prefix: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=65536 prefix: ordered vs baseline: 6 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.64%, 2.02%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.10% does not clear the 0.60% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.43%, 0.24%] includes zero
- multi url n=262144 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of 0.05% does not clear the 2.04% noise floor, the bound on what the harness reports between identical code in every process
- multi url n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of -1.08% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=262144 prefix: ordered vs baseline: the pooled interval [-0.40%, 0.51%] includes zero
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.21% does not clear the 1.16% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.73%, 1.30%] includes zero
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.89% does not clear the 1.88% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.93%, 3.71%] includes zero
- multi uuid n=4096 prefix: ordered vs baseline: the pooled difference of 0.07% does not clear the 0.80% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=4096 prefix: ordered vs baseline: the pooled interval [-0.60%, 0.73%] includes zero
- multi uuid n=4096 churn: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.29% does not clear the 0.30% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.45%, 1.03%] includes zero
- multi uuid n=16384 prefix: ordered vs baseline: the pooled difference of 0.20% does not clear the 0.23% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=16384 prefix: ordered vs baseline: the pooled interval [-0.87%, 1.28%] includes zero
- multi uuid n=65536 valuesFor: ordered vs baseline: the pooled interval [-3.56%, 0.14%] includes zero
- multi uuid n=65536 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=65536 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=65536 valuesBetween: ordered vs baseline: the pooled difference of -0.23% does not clear the 0.34% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.61%, 1.14%] includes zero
- multi uuid n=65536 prefix: ordered vs baseline: the pooled difference of -0.45% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=65536 prefix: ordered vs baseline: the pooled interval [-2.65%, 1.74%] includes zero
- multi uuid n=65536 prefix: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=65536 prefix: ordered vs baseline: 2 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.18% does not clear the 0.45% noise floor, the bound on what the harness reports between identical code in every process
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.65%, 2.02%] includes zero
- multi uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesBetween: ordered vs baseline: 4 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 prefix: ordered vs baseline: the pooled interval [-2.55%, 0.18%] includes zero
- multi uuid n=262144 prefix: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 prefix: ordered vs baseline: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=4096 build: ordered vs baseline: the pooled interval [-1.16%, 0.04%] includes zero
- unique email n=4096 build: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.67% does not clear the 0.88% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.68%, 0.34%] includes zero
- unique email n=16384 prefix: ordered vs baseline: the pooled difference of -0.39% does not clear the 0.45% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=16384 prefix: ordered vs baseline: the pooled interval [-1.06%, 0.27%] includes zero
- unique email n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.13% does not clear the 0.23% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=65536 valuesBetween: ordered vs baseline: the pooled interval [-0.87%, 1.12%] includes zero
- unique email n=65536 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.70%, 0.17%] includes zero
- unique email n=262144 churn: ordered vs baseline: the pooled difference of -1.50% does not clear the 1.77% noise floor, the bound on what the harness reports between identical code in every process
- unique email n=262144 churn: ordered vs baseline: the A/A validations found a systematic difference of +1.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=262144 churn: ordered vs baseline: the pooled interval [-3.35%, 0.35%] includes zero
- unique path n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.38%, 0.28%] includes zero
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.71% does not clear the 1.15% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.60%, 3.01%] includes zero
- unique path n=4096 prefix: ordered vs baseline: the pooled difference of 0.73% does not clear the 0.74% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique path n=4096 prefix: ordered vs baseline: the pooled interval [-0.02%, 1.48%] includes zero
- unique path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.53% does not clear the 0.93% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.96%, 3.01%] includes zero
- unique path n=16384 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=16384 valuesBetween: ordered vs baseline: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=65536 valuesFor: ordered vs baseline: the pooled interval [-2.33%, 0.83%] includes zero
- unique path n=65536 valuesBetween: ordered vs baseline: the pooled interval [-0.02%, 1.98%] includes zero
- unique path n=65536 prefix: ordered vs baseline: the pooled difference of 2.22% does not clear the 6.71% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.15% does not clear the 0.85% noise floor, the bound on what the harness reports between identical code in every process
- unique path n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.65%, 1.94%] includes zero
- unique path n=262144 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.07%, 3.05%] includes zero
- unique path n=262144 prefix: ordered vs baseline: the pooled difference of 2.22% does not clear the 6.44% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 build: ordered vs baseline: the pooled difference of -0.43% does not clear the 0.72% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=4096 build: ordered vs baseline: the pooled interval [-1.44%, 0.58%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.51%, 1.23%] includes zero
- unique str n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique str n=16384 valuesBetween: ordered vs baseline: the pooled difference of 1.22% does not clear the 1.22% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.74%, 0.32%] includes zero
- unique str n=65536 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique str n=65536 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique str n=65536 valuesBetween: ordered vs baseline: the pooled interval [-2.58%, 0.66%] includes zero
- unique str n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=65536 prefix: ordered vs baseline: the pooled interval [-2.51%, 1.01%] includes zero
- unique str n=65536 prefix: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=65536 churn: ordered vs baseline: the pooled difference of -0.94% does not clear the 1.03% noise floor, the bound on what the harness reports between identical code in every process
- unique str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-5.57%, 1.67%] includes zero
- unique str n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique str n=262144 prefix: ordered vs baseline: the pooled interval [-7.67%, 2.62%] includes zero
- unique street n=4096 churn: ordered vs baseline: the pooled difference of 0.33% does not clear the 0.92% noise floor, the bound on what the harness reports between identical code in every process
- unique street n=4096 churn: ordered vs baseline: the pooled interval [-0.48%, 1.14%] includes zero
- unique street n=4096 build: ordered vs baseline: the pooled interval [-0.44%, 1.26%] includes zero
- unique street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.96% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=16384 build: ordered vs baseline: the pooled interval [-1.49%, 0.13%] includes zero
- unique street n=16384 build: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique street n=65536 valuesFor: ordered vs baseline: the pooled interval [-6.36%, 1.84%] includes zero
- unique street n=65536 valuesFor: ordered vs baseline: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.34% does not clear the 0.60% noise floor, the bound on what the harness reports between identical code in every process
- unique street n=65536 valuesBetween: ordered vs baseline: the pooled interval [-2.84%, 3.52%] includes zero
- unique street n=65536 valuesBetween: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=65536 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique street n=65536 prefix: ordered vs baseline: the pooled interval [-0.11%, 4.19%] includes zero
- unique u64 n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -1.06% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=65536 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=65536 build: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.29%, 16.87%] includes zero
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 9.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: 10 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.87% does not clear the 1.19% noise floor, the bound on what the harness reports between identical code in every process
- unique url n=65536 valuesFor: ordered vs baseline: the pooled interval [-2.45%, 0.72%] includes zero
- unique url n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.20%, 2.16%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the pooled interval [-4.44%, 2.27%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.38%, 1.15%] includes zero
- unique url n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 prefix: ordered vs baseline: the pooled difference of 0.40% does not clear the 1.26% noise floor, the bound on what the harness reports between identical code in every process
- unique url n=262144 prefix: ordered vs baseline: the pooled interval [-2.19%, 3.00%] includes zero
- unique url n=262144 prefix: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 prefix: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 churn: ordered vs baseline: the pooled interval [-3.81%, 0.07%] includes zero
- unique uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -1.46% does not clear the 1.55% noise floor, the bound on what the harness reports between identical code in every process
- unique uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.66%, 2.83%] includes zero
- unique uuid n=4096 prefix: ordered vs baseline: the pooled interval [-1.22%, 0.25%] includes zero
- unique uuid n=16384 prefix: ordered vs baseline: the pooled difference of -0.14% does not clear the 0.54% noise floor, the bound on what the harness reports between identical code in every process
- unique uuid n=16384 prefix: ordered vs baseline: the pooled interval [-0.60%, 0.33%] includes zero
- unique uuid n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique uuid n=65536 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=65536 valuesBetween: ordered vs baseline: the pooled difference of -0.10% does not clear the 0.20% noise floor, the bound on what the harness reports between identical code in every process
- unique uuid n=65536 valuesBetween: ordered vs baseline: the pooled interval [-0.53%, 0.34%] includes zero
- unique uuid n=65536 prefix: ordered vs baseline: the pooled difference of 0.45% does not clear the 0.48% noise floor, the bound on what the harness reports between identical code in every process
- unique uuid n=65536 prefix: ordered vs baseline: the pooled interval [-0.65%, 1.55%] includes zero
- unique uuid n=65536 prefix: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=65536 prefix: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=65536 churn: ordered vs baseline: the pooled difference of -1.40% does not clear the 1.47% noise floor, the bound on what the harness reports between identical code in every process
- unique uuid n=65536 churn: ordered vs baseline: the pooled interval [-3.25%, 0.46%] includes zero
- unique uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.62% does not clear the 1.14% noise floor, the bound on what the harness reports between identical code in every process
- unique uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.59%, 0.34%] includes zero
- unique uuid n=262144 prefix: ordered vs baseline: the pooled interval [-1.70%, 0.33%] includes zero
- unique uuid n=262144 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 prefix: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
