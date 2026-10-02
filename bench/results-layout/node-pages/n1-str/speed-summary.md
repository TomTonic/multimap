| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | email | 4096 | valuesFor | ordered | baseline | 12 | 49.8 | 49.6 | 1.01× [1.00, 1.02] | +1.0% | [-0.5%, +2.4%] | 5.6 pts | 2.2 | yes | no |
| multi-str | email | 4096 | valuesBetween | ordered | baseline | 12 | 2952 | 3000 | 1.02× [0.99, 1.05] | +1.7% | [-0.8%, +4.3%] | 3.0 pts | 1.1 | no | no |
| multi-str | email | 4096 | prefix | ordered | baseline | 12 | 79.1 | 78.6 | 1.00× [0.99, 1.01] | -0.2% | [-0.9%, +0.6%] | 0.9 pts | 0.8 | yes | no |
| multi-str | email | 4096 | churn | ordered | baseline | 12 | 88.4 | 85.7 | 0.97× [0.96, 0.98] | -3.2% | [-4.2%, -2.2%] | 1.3 pts | 2.0 | yes | yes |
| multi-str | email | 4096 | build | ordered | baseline | 12 | 9.31 ms | 8.95 ms | 0.96× [0.96, 0.97] | -3.8% | [-4.6%, -3.1%] | 0.8 pts | 1.4 | yes | yes |
| multi-str | email | 16384 | valuesFor | ordered | baseline | 6 | 61.9 | 62.4 | 1.01× [1.00, 1.01] | +0.9% | [+0.4%, +1.4%] | 0.5 pts | 1.1 | yes | yes |
| multi-str | email | 16384 | valuesBetween | ordered | baseline | 6 | 3578 | 3624 | 1.01× [1.00, 1.02] | +1.3% | [+0.3%, +2.3%] | 1.0 pts | 1.5 | yes | yes |
| multi-str | email | 16384 | prefix | ordered | baseline | 6 | 93.3 | 92.9 | 1.00× [0.99, 1.00] | -0.5% | [-0.9%, +0.0%] | 0.5 pts | 0.8 | yes | no |
| multi-str | email | 16384 | churn | ordered | baseline | 6 | 127 | 124 | 0.97× [0.96, 0.98] | -3.1% | [-3.9%, -2.3%] | 0.8 pts | 1.1 | yes | yes |
| multi-str | email | 16384 | build | ordered | baseline | 6 | 46.99 ms | 45.23 ms | 0.96× [0.95, 0.97] | -4.0% | [-5.3%, -2.7%] | 1.2 pts | 1.3 | yes | yes |
| multi-str | email | 262144 | valuesFor | ordered | baseline | 12 | 229 | 227 | 0.98× [0.96, 1.01] | -1.6% | [-4.2%, +1.0%] | 2.9 pts | 4.3 | no | no |
| multi-str | email | 262144 | valuesBetween | ordered | baseline | 12 | 8711 | 8718 | 1.00× [0.98, 1.02] | -0.0% | [-2.2%, +2.1%] | 3.2 pts | 4.2 | no | no |
| multi-str | email | 262144 | prefix | ordered | baseline | 12 | 315 | 315 | 1.00× [0.98, 1.02] | +0.3% | [-1.6%, +2.1%] | 2.6 pts | 4.6 | yes | no |
| multi-str | email | 262144 | churn | ordered | baseline | 12 | 503 | 490 | 0.94× [0.91, 0.97] | -6.8% | [-10.2%, -3.4%] | 6.3 pts | 0.9 | no | yes |
| multi-str | path | 4096 | valuesFor | ordered | baseline | 12 | 107 | 106 | 0.99× [0.99, 1.00] | -0.6% | [-1.1%, -0.0%] | 2.2 pts | 1.8 | yes | no |
| multi-str | path | 4096 | valuesBetween | ordered | baseline | 12 | 4099 | 4147 | 1.01× [0.99, 1.03] | +0.9% | [-0.9%, +2.8%] | 2.4 pts | 1.1 | yes | no |
| multi-str | path | 4096 | prefix | ordered | baseline | 12 | 311 | 318 | 1.03× [1.02, 1.04] | +2.5% | [+1.6%, +3.4%] | 1.1 pts | 0.8 | yes | yes |
| multi-str | path | 4096 | churn | ordered | baseline | 12 | 175 | 169 | 0.97× [0.96, 0.97] | -3.5% | [-4.1%, -2.9%] | 0.9 pts | 1.2 | yes | yes |
| multi-str | path | 4096 | build | ordered | baseline | 12 | 17.42 ms | 16.60 ms | 0.95× [0.95, 0.96] | -5.0% | [-5.7%, -4.3%] | 0.8 pts | 1.2 | yes | yes |
| multi-str | path | 16384 | valuesFor | ordered | baseline | 10 | 137 | 137 | 1.00× [1.00, 1.01] | +0.1% | [-0.4%, +0.6%] | 0.7 pts | 1.0 | yes | no |
| multi-str | path | 16384 | valuesBetween | ordered | baseline | 10 | 4723 | 4783 | 1.01× [1.00, 1.02] | +1.2% | [+0.4%, +2.0%] | 0.9 pts | 1.3 | yes | yes |
| multi-str | path | 16384 | prefix | ordered | baseline | 10 | 740 | 734 | 1.00× [0.98, 1.02] | -0.0% | [-1.9%, +1.9%] | 3.8 pts | 1.1 | yes | no |
| multi-str | path | 16384 | churn | ordered | baseline | 10 | 255 | 249 | 0.97× [0.96, 0.98] | -3.1% | [-3.7%, -2.5%] | 0.6 pts | 0.6 | yes | yes |
| multi-str | path | 16384 | build | ordered | baseline | 10 | 88.04 ms | 84.32 ms | 0.96× [0.95, 0.97] | -4.2% | [-5.1%, -3.4%] | 0.8 pts | 1.2 | yes | yes |
| multi-str | path | 262144 | valuesFor | ordered | baseline | 12 | 414 | 421 | 1.01× [0.99, 1.03] | +0.8% | [-1.2%, +2.8%] | 2.7 pts | 3.3 | no | no |
| multi-str | path | 262144 | valuesBetween | ordered | baseline | 12 | 11.8 µs | 11.8 µs | 1.01× [0.99, 1.03] | +0.8% | [-1.0%, +2.5%] | 2.3 pts | 3.3 | yes | no |
| multi-str | path | 262144 | prefix | ordered | baseline | 12 | 11.3 µs | 11.4 µs | 1.01× [0.99, 1.04] | +1.0% | [-1.5%, +3.4%] | 3.3 pts | 1.5 | no | no |
| multi-str | path | 262144 | churn | ordered | baseline | 12 | 783 | 765 | 0.94× [0.90, 0.99] | -6.4% | [-11.5%, -1.3%] | 5.9 pts | 0.6 | no | yes |
| multi-str | str | 4096 | valuesFor | ordered | baseline | 12 | 64.3 | 64.6 | 1.00× [1.00, 1.01] | +0.2% | [-0.3%, +0.8%] | 1.3 pts | 0.8 | yes | no |
| multi-str | str | 4096 | valuesBetween | ordered | baseline | 12 | 3482 | 3568 | 1.01× [1.00, 1.03] | +1.5% | [-0.3%, +3.3%] | 2.3 pts | 1.1 | yes | no |
| multi-str | str | 4096 | prefix | ordered | baseline | 12 | 7392 | 7549 | 1.03× [1.00, 1.06] | +3.2% | [+0.3%, +6.1%] | 4.7 pts | 1.0 | no | yes |
| multi-str | str | 4096 | churn | ordered | baseline | 12 | 114 | 110 | 0.95× [0.95, 0.96] | -4.8% | [-5.2%, -4.3%] | 1.6 pts | 2.2 | yes | yes |
| multi-str | str | 4096 | build | ordered | baseline | 12 | 11.84 ms | 11.29 ms | 0.95× [0.94, 0.96] | -5.4% | [-6.1%, -4.6%] | 0.9 pts | 1.5 | yes | yes |
| multi-str | str | 16384 | valuesFor | ordered | baseline | 12 | 75.1 | 75.9 | 1.02× [0.98, 1.06] | +1.8% | [-2.3%, +6.0%] | 4.0 pts | 5.7 | no | no |
| multi-str | str | 16384 | valuesBetween | ordered | baseline | 12 | 3766 | 3826 | 1.02× [1.01, 1.02] | +1.5% | [+1.0%, +2.0%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | str | 16384 | prefix | ordered | baseline | 12 | 35.0 µs | 35.2 µs | 1.01× [1.00, 1.01] | +0.6% | [+0.1%, +1.1%] | 0.6 pts | 0.8 | yes | yes |
| multi-str | str | 16384 | churn | ordered | baseline | 12 | 153 | 147 | 0.97× [0.96, 0.98] | -3.3% | [-4.1%, -2.4%] | 0.9 pts | 0.8 | yes | yes |
| multi-str | str | 16384 | build | ordered | baseline | 12 | 56.49 ms | 54.15 ms | 0.96× [0.95, 0.98] | -4.0% | [-5.5%, -2.4%] | 1.5 pts | 1.8 | yes | yes |
| multi-str | str | 262144 | valuesFor | ordered | baseline | 12 | 247 | 245 | 0.99× [0.97, 1.02] | -0.8% | [-3.3%, +1.7%] | 3.2 pts | 4.9 | no | no |
| multi-str | str | 262144 | valuesBetween | ordered | baseline | 12 | 9340 | 9234 | 0.99× [0.97, 1.01] | -0.6% | [-2.6%, +1.4%] | 2.5 pts | 4.1 | no | no |
| multi-str | str | 262144 | prefix | ordered | baseline | 12 | 1.48 ms | 1.48 ms | 1.00× [0.98, 1.02] | -0.3% | [-2.3%, +1.8%] | 2.7 pts | 4.1 | no | no |
| multi-str | str | 262144 | churn | ordered | baseline | 12 | 530 | 510 | 0.95× [0.92, 0.98] | -5.6% | [-9.2%, -1.9%] | 5.4 pts | 0.6 | no | yes |
| multi-str | street | 4096 | valuesFor | ordered | baseline | 12 | 49.7 | 49.0 | 0.98× [0.97, 1.00] | -1.8% | [-3.2%, -0.3%] | 2.3 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | baseline | 12 | 2279 | 2341 | 1.03× [1.01, 1.05] | +2.6% | [+0.7%, +4.4%] | 2.8 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | baseline | 12 | 222 | 228 | 1.03× [1.01, 1.04] | +2.5% | [+1.1%, +3.8%] | 1.9 pts | 1.4 | yes | yes |
| multi-str | street | 4096 | churn | ordered | baseline | 12 | 114 | 110 | 0.97× [0.96, 0.98] | -3.1% | [-4.2%, -2.1%] | 1.1 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | build | ordered | baseline | 12 | 4.58 ms | 4.42 ms | 0.96× [0.95, 0.97] | -3.7% | [-4.8%, -2.6%] | 1.0 pts | 1.2 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | baseline | 6 | 74.7 | 74.2 | 0.99× [0.99, 1.00] | -0.6% | [-1.5%, +0.2%] | 0.8 pts | 1.1 | yes | no |
| multi-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 2881 | 2949 | 1.02× [1.01, 1.03] | +1.9% | [+0.7%, +3.1%] | 1.1 pts | 1.5 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | baseline | 6 | 825 | 845 | 1.03× [1.02, 1.05] | +3.2% | [+1.8%, +4.6%] | 1.3 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | churn | ordered | baseline | 6 | 153 | 149 | 0.97× [0.96, 0.98] | -3.0% | [-4.1%, -1.9%] | 1.0 pts | 1.5 | yes | yes |
| multi-str | street | 16384 | build | ordered | baseline | 6 | 24.03 ms | 23.23 ms | 0.97× [0.96, 0.97] | -3.5% | [-4.5%, -2.6%] | 0.9 pts | 1.5 | yes | yes |
| multi-str | u64 | 4096 | valuesFor | ordered | baseline | 12 | 35.3 | 34.3 | 0.99× [0.97, 1.02] | -0.6% | [-3.0%, +1.8%] | 4.6 pts | 0.9 | no | no |
| multi-str | u64 | 4096 | valuesBetween | ordered | baseline | 12 | 2637 | 2704 | 1.02× [1.00, 1.05] | +2.4% | [+0.2%, +4.7%] | 3.2 pts | 1.0 | no | yes |
| multi-str | u64 | 4096 | churn | ordered | baseline | 12 | 68.9 | 66.7 | 0.97× [0.96, 0.98] | -3.2% | [-3.8%, -2.5%] | 0.8 pts | 1.4 | yes | yes |
| multi-str | u64 | 4096 | build | ordered | baseline | 12 | 7.63 ms | 7.30 ms | 0.95× [0.95, 0.96] | -4.8% | [-5.3%, -4.2%] | 0.7 pts | 0.8 | yes | yes |
| multi-str | u64 | 16384 | valuesFor | ordered | baseline | 6 | 48.3 | 48.0 | 0.99× [0.97, 1.00] | -1.3% | [-2.8%, +0.1%] | 1.4 pts | 2.5 | yes | no |
| multi-str | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3561 | 3641 | 1.02× [1.02, 1.03] | +2.2% | [+1.6%, +2.8%] | 0.6 pts | 0.7 | yes | yes |
| multi-str | u64 | 16384 | churn | ordered | baseline | 6 | 99.4 | 95.8 | 0.96× [0.96, 0.97] | -3.9% | [-4.6%, -3.2%] | 0.6 pts | 0.6 | yes | yes |
| multi-str | u64 | 16384 | build | ordered | baseline | 6 | 37.87 ms | 35.97 ms | 0.97× [0.96, 0.98] | -3.4% | [-4.7%, -2.0%] | 1.3 pts | 0.9 | yes | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | baseline | 12 | 155 | 154 | 0.98× [0.95, 1.02] | -1.8% | [-5.2%, +1.7%] | 4.1 pts | 6.6 | no | no |
| multi-str | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 7776 | 7867 | 1.01× [0.98, 1.03] | +0.6% | [-2.0%, +3.2%] | 2.7 pts | 4.1 | no | no |
| multi-str | u64 | 262144 | churn | ordered | baseline | 12 | 394 | 381 | 0.92× [0.87, 0.98] | -8.8% | [-15.5%, -2.1%] | 7.1 pts | 0.7 | no | yes |
| multi-str | url | 4096 | valuesFor | ordered | baseline | 6 | 88.8 | 89.0 | 1.00× [0.98, 1.02] | +0.2% | [-1.6%, +1.9%] | 1.7 pts | 1.3 | yes | no |
| multi-str | url | 4096 | valuesBetween | ordered | baseline | 6 | 4086 | 4150 | 1.02× [1.00, 1.04] | +2.0% | [+0.4%, +3.5%] | 1.5 pts | 0.7 | yes | yes |
| multi-str | url | 4096 | prefix | ordered | baseline | 6 | 180 | 185 | 1.03× [1.01, 1.04] | +2.4% | [+0.9%, +4.0%] | 1.5 pts | 1.3 | yes | yes |
| multi-str | url | 4096 | churn | ordered | baseline | 6 | 148 | 144 | 0.96× [0.95, 0.97] | -3.7% | [-4.8%, -2.7%] | 1.0 pts | 1.6 | yes | yes |
| multi-str | url | 4096 | build | ordered | baseline | 6 | 14.70 ms | 14.01 ms | 0.96× [0.95, 0.96] | -4.5% | [-5.2%, -3.9%] | 0.7 pts | 0.8 | yes | yes |
| multi-str | url | 16384 | valuesFor | ordered | baseline | 6 | 113 | 113 | 1.00× [0.99, 1.00] | -0.3% | [-0.8%, +0.3%] | 0.5 pts | 0.9 | yes | no |
| multi-str | url | 16384 | valuesBetween | ordered | baseline | 6 | 4766 | 4790 | 1.01× [1.00, 1.02] | +0.8% | [-0.1%, +1.7%] | 0.8 pts | 1.2 | yes | no |
| multi-str | url | 16384 | prefix | ordered | baseline | 6 | 283 | 293 | 1.03× [1.01, 1.05] | +2.9% | [+1.4%, +4.5%] | 1.5 pts | 1.4 | yes | yes |
| multi-str | url | 16384 | churn | ordered | baseline | 6 | 223 | 218 | 0.97× [0.96, 0.97] | -3.4% | [-4.0%, -2.8%] | 0.5 pts | 0.5 | yes | yes |
| multi-str | url | 16384 | build | ordered | baseline | 6 | 75.70 ms | 72.34 ms | 0.96× [0.95, 0.96] | -4.6% | [-5.6%, -3.6%] | 0.9 pts | 1.4 | yes | yes |
| multi-str | url | 262144 | valuesFor | ordered | baseline | 12 | 393 | 393 | 1.01× [0.99, 1.02] | +0.8% | [-0.6%, +2.3%] | 1.6 pts | 2.5 | yes | no |
| multi-str | url | 262144 | valuesBetween | ordered | baseline | 12 | 12.4 µs | 12.4 µs | 1.00× [1.00, 1.01] | +0.5% | [-0.1%, +1.1%] | 1.0 pts | 1.8 | yes | no |
| multi-str | url | 262144 | prefix | ordered | baseline | 12 | 2535 | 2559 | 1.01× [0.99, 1.02] | +0.5% | [-1.0%, +2.0%] | 1.7 pts | 1.2 | yes | no |
| multi-str | url | 262144 | churn | ordered | baseline | 12 | 736 | 725 | 0.96× [0.93, 0.99] | -4.4% | [-7.4%, -1.5%] | 3.5 pts | 0.7 | no | yes |
| multi-str | uuid | 4096 | valuesFor | ordered | baseline | 8 | 55.9 | 56.0 | 1.00× [0.99, 1.01] | -0.2% | [-1.3%, +0.9%] | 1.8 pts | 0.8 | yes | no |
| multi-str | uuid | 4096 | valuesBetween | ordered | baseline | 8 | 3218 | 3288 | 1.01× [0.99, 1.03] | +1.1% | [-0.7%, +3.0%] | 3.4 pts | 1.1 | yes | no |
| multi-str | uuid | 4096 | prefix | ordered | baseline | 8 | 90.9 | 91.3 | 1.00× [0.99, 1.02] | +0.4% | [-0.6%, +1.5%] | 1.3 pts | 1.3 | yes | no |
| multi-str | uuid | 4096 | churn | ordered | baseline | 8 | 96.8 | 93.9 | 0.97× [0.96, 0.98] | -3.5% | [-4.5%, -2.4%] | 1.1 pts | 1.4 | yes | yes |
| multi-str | uuid | 4096 | build | ordered | baseline | 8 | 10.07 ms | 9.66 ms | 0.96× [0.95, 0.97] | -4.2% | [-4.8%, -3.6%] | 0.6 pts | 1.0 | yes | yes |
| multi-str | uuid | 16384 | valuesFor | ordered | baseline | 12 | 65.7 | 66.4 | 1.01× [1.00, 1.02] | +0.8% | [+0.2%, +1.5%] | 0.7 pts | 1.3 | yes | yes |
| multi-str | uuid | 16384 | valuesBetween | ordered | baseline | 12 | 3633 | 3688 | 1.02× [1.01, 1.02] | +1.5% | [+0.9%, +2.1%] | 0.7 pts | 1.3 | yes | yes |
| multi-str | uuid | 16384 | prefix | ordered | baseline | 12 | 109 | 110 | 1.00× [0.97, 1.03] | -0.4% | [-3.4%, +2.7%] | 2.9 pts | 5.5 | no | no |
| multi-str | uuid | 16384 | churn | ordered | baseline | 12 | 134 | 131 | 0.97× [0.97, 0.98] | -3.0% | [-3.5%, -2.5%] | 0.9 pts | 1.0 | yes | yes |
| multi-str | uuid | 16384 | build | ordered | baseline | 12 | 48.09 ms | 46.38 ms | 0.96× [0.96, 0.97] | -3.7% | [-4.5%, -2.8%] | 0.8 pts | 1.4 | yes | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | baseline | 12 | 248 | 249 | 0.99× [0.96, 1.01] | -1.3% | [-3.9%, +1.4%] | 3.5 pts | 5.7 | no | no |
| multi-str | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 9193 | 9213 | 1.00× [0.98, 1.02] | +0.3% | [-1.5%, +2.2%] | 2.3 pts | 3.5 | yes | no |
| multi-str | uuid | 262144 | prefix | ordered | baseline | 12 | 592 | 591 | 1.00× [0.98, 1.02] | -0.0% | [-2.1%, +2.0%] | 2.4 pts | 5.5 | no | no |
| multi-str | uuid | 262144 | churn | ordered | baseline | 12 | 529 | 511 | 0.92× [0.88, 0.97] | -8.4% | [-13.5%, -3.4%] | 6.6 pts | 0.7 | no | yes |
| unique-str | email | 4096 | valuesFor | ordered | baseline | 12 | 22.6 | 22.2 | 0.98× [0.95, 1.01] | -2.3% | [-5.6%, +1.0%] | 4.4 pts | 2.4 | no | no |
| unique-str | email | 4096 | valuesBetween | ordered | baseline | 12 | 1251 | 1316 | 1.05× [1.04, 1.06] | +4.6% | [+3.7%, +5.5%] | 1.2 pts | 1.1 | yes | yes |
| unique-str | email | 4096 | prefix | ordered | baseline | 12 | 57.8 | 57.5 | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.4%] | 1.0 pts | 1.6 | yes | no |
| unique-str | email | 4096 | churn | ordered | baseline | 12 | 78.0 | 75.2 | 0.96× [0.95, 0.97] | -4.5% | [-5.6%, -3.3%] | 1.5 pts | 1.8 | yes | yes |
| unique-str | email | 4096 | build | ordered | baseline | 12 | 1.15 ms | 1.10 ms | 0.95× [0.95, 0.96] | -4.8% | [-5.5%, -4.1%] | 1.8 pts | 1.4 | yes | yes |
| unique-str | email | 16384 | valuesFor | ordered | baseline | 6 | 38.1 | 38.0 | 1.00× [0.99, 1.00] | -0.3% | [-1.1%, +0.4%] | 0.7 pts | 0.9 | yes | no |
| unique-str | email | 16384 | valuesBetween | ordered | baseline | 6 | 1590 | 1641 | 1.03× [1.02, 1.04] | +2.9% | [+1.9%, +4.0%] | 1.0 pts | 1.3 | yes | yes |
| unique-str | email | 16384 | prefix | ordered | baseline | 6 | 71.8 | 72.1 | 1.01× [1.00, 1.02] | +0.7% | [-0.1%, +1.5%] | 0.8 pts | 1.5 | yes | no |
| unique-str | email | 16384 | churn | ordered | baseline | 6 | 96.1 | 92.6 | 0.96× [0.95, 0.97] | -4.1% | [-5.6%, -2.6%] | 1.4 pts | 1.2 | yes | yes |
| unique-str | email | 16384 | build | ordered | baseline | 6 | 5.51 ms | 5.28 ms | 0.96× [0.95, 0.96] | -4.4% | [-5.1%, -3.6%] | 0.7 pts | 0.9 | yes | yes |
| unique-str | email | 262144 | valuesFor | ordered | baseline | 12 | 141 | 140 | 0.98× [0.96, 1.00] | -2.1% | [-3.8%, -0.4%] | 1.9 pts | 1.8 | yes | yes |
| unique-str | email | 262144 | valuesBetween | ordered | baseline | 12 | 3558 | 3578 | 1.01× [1.00, 1.02] | +0.7% | [-0.3%, +1.7%] | 1.3 pts | 0.9 | yes | no |
| unique-str | email | 262144 | prefix | ordered | baseline | 12 | 187 | 189 | 1.01× [1.00, 1.02] | +0.9% | [-0.2%, +1.9%] | 1.2 pts | 2.0 | yes | no |
| unique-str | email | 262144 | churn | ordered | baseline | 12 | 278 | 278 | 0.98× [0.96, 1.00] | -2.4% | [-4.4%, -0.4%] | 3.0 pts | 1.0 | yes | yes |
| unique-str | path | 4096 | valuesFor | ordered | baseline | 6 | 77.4 | 76.6 | 0.99× [0.98, 1.01] | -0.5% | [-2.1%, +1.1%] | 1.5 pts | 0.9 | yes | no |
| unique-str | path | 4096 | valuesBetween | ordered | baseline | 6 | 2090 | 2145 | 1.01× [0.99, 1.03] | +1.1% | [-0.5%, +2.7%] | 1.6 pts | 0.7 | yes | no |
| unique-str | path | 4096 | prefix | ordered | baseline | 6 | 228 | 236 | 1.03× [1.03, 1.04] | +3.1% | [+2.6%, +3.7%] | 0.5 pts | 0.6 | yes | yes |
| unique-str | path | 4096 | churn | ordered | baseline | 6 | 207 | 199 | 0.96× [0.95, 0.97] | -4.1% | [-5.0%, -3.2%] | 0.8 pts | 0.9 | yes | yes |
| unique-str | path | 4096 | build | ordered | baseline | 6 | 2.85 ms | 2.72 ms | 0.96× [0.95, 0.97] | -4.2% | [-5.2%, -3.3%] | 0.9 pts | 0.8 | yes | yes |
| unique-str | path | 16384 | valuesFor | ordered | baseline | 6 | 112 | 111 | 1.00× [0.99, 1.01] | -0.1% | [-1.4%, +1.2%] | 1.2 pts | 1.3 | yes | no |
| unique-str | path | 16384 | valuesBetween | ordered | baseline | 6 | 2616 | 2678 | 1.02× [1.01, 1.03] | +2.0% | [+1.0%, +2.9%] | 0.9 pts | 1.3 | yes | yes |
| unique-str | path | 16384 | prefix | ordered | baseline | 6 | 475 | 489 | 1.02× [1.00, 1.04] | +2.3% | [+0.4%, +4.1%] | 1.8 pts | 0.7 | yes | yes |
| unique-str | path | 16384 | churn | ordered | baseline | 6 | 258 | 251 | 0.97× [0.96, 0.98] | -3.6% | [-4.7%, -2.6%] | 1.0 pts | 1.2 | yes | yes |
| unique-str | path | 16384 | build | ordered | baseline | 6 | 14.02 ms | 13.34 ms | 0.96× [0.95, 0.97] | -4.3% | [-5.7%, -2.9%] | 1.3 pts | 1.5 | yes | yes |
| unique-str | path | 262144 | valuesFor | ordered | baseline | 12 | 342 | 342 | 1.00× [0.99, 1.01] | -0.1% | [-1.1%, +0.9%] | 1.9 pts | 1.7 | yes | no |
| unique-str | path | 262144 | valuesBetween | ordered | baseline | 12 | 6254 | 6310 | 1.00× [0.98, 1.01] | -0.4% | [-1.6%, +0.7%] | 2.5 pts | 2.1 | yes | no |
| unique-str | path | 262144 | prefix | ordered | baseline | 12 | 4777 | 5082 | 1.00× [0.99, 1.02] | +0.4% | [-1.5%, +2.3%] | 3.4 pts | 1.3 | yes | no |
| unique-str | path | 262144 | churn | ordered | baseline | 12 | 657 | 649 | 0.96× [0.94, 0.99] | -3.8% | [-6.9%, -0.8%] | 3.8 pts | 1.1 | no | yes |
| unique-str | str | 4096 | valuesFor | ordered | baseline | 12 | 34.3 | 33.0 | 0.98× [0.96, 1.00] | -2.2% | [-4.4%, -0.0%] | 2.7 pts | 0.7 | no | yes |
| unique-str | str | 4096 | valuesBetween | ordered | baseline | 12 | 1620 | 1661 | 1.02× [1.01, 1.02] | +1.5% | [+0.7%, +2.4%] | 1.1 pts | 0.7 | yes | yes |
| unique-str | str | 4096 | prefix | ordered | baseline | 12 | 3186 | 3254 | 1.02× [1.01, 1.03] | +1.9% | [+1.1%, +2.8%] | 1.1 pts | 1.3 | yes | yes |
| unique-str | str | 4096 | churn | ordered | baseline | 12 | 105 | 100 | 0.96× [0.96, 0.97] | -4.0% | [-4.6%, -3.4%] | 0.9 pts | 1.1 | yes | yes |
| unique-str | str | 4096 | build | ordered | baseline | 12 | 1.52 ms | 1.45 ms | 0.96× [0.95, 0.97] | -4.4% | [-5.1%, -3.6%] | 1.5 pts | 1.0 | yes | yes |
| unique-str | str | 16384 | valuesFor | ordered | baseline | 6 | 49.5 | 48.6 | 0.98× [0.98, 0.99] | -1.7% | [-2.4%, -1.0%] | 0.6 pts | 0.8 | yes | yes |
| unique-str | str | 16384 | valuesBetween | ordered | baseline | 6 | 1638 | 1684 | 1.03× [1.02, 1.04] | +2.5% | [+1.5%, +3.5%] | 1.0 pts | 1.3 | yes | yes |
| unique-str | str | 16384 | prefix | ordered | baseline | 6 | 13.6 µs | 13.8 µs | 1.01× [0.99, 1.02] | +0.7% | [-0.6%, +2.1%] | 1.3 pts | 1.1 | yes | no |
| unique-str | str | 16384 | churn | ordered | baseline | 6 | 124 | 119 | 0.96× [0.95, 0.97] | -4.1% | [-4.9%, -3.3%] | 0.8 pts | 0.8 | yes | yes |
| unique-str | str | 16384 | build | ordered | baseline | 6 | 7.07 ms | 6.75 ms | 0.96× [0.95, 0.96] | -4.5% | [-5.0%, -4.0%] | 0.5 pts | 0.8 | yes | yes |
| unique-str | str | 262144 | valuesFor | ordered | baseline | 12 | 147 | 146 | 0.99× [0.97, 1.01] | -1.0% | [-3.2%, +1.3%] | 2.5 pts | 1.4 | no | no |
| unique-str | str | 262144 | valuesBetween | ordered | baseline | 12 | 3226 | 3247 | 1.00× [0.98, 1.02] | +0.2% | [-1.6%, +2.1%] | 2.2 pts | 1.0 | yes | no |
| unique-str | str | 262144 | prefix | ordered | baseline | 12 | 463.9 µs | 463.0 µs | 0.99× [0.97, 1.01] | -0.7% | [-2.7%, +1.4%] | 3.7 pts | 1.1 | no | no |
| unique-str | str | 262144 | churn | ordered | baseline | 12 | 305 | 299 | 0.98× [0.97, 0.99] | -2.4% | [-3.5%, -1.3%] | 1.7 pts | 0.8 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | baseline | 8 | 40.8 | 39.3 | 0.96× [0.95, 0.97] | -4.5% | [-5.8%, -3.1%] | 1.6 pts | 0.5 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 1759 | 1803 | 1.04× [1.02, 1.06] | +3.7% | [+1.7%, +5.6%] | 2.0 pts | 1.2 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | baseline | 8 | 197 | 199 | 1.01× [1.00, 1.03] | +1.5% | [+0.3%, +2.6%] | 1.4 pts | 1.3 | yes | yes |
| unique-str | street | 4096 | churn | ordered | baseline | 8 | 122 | 119 | 0.97× [0.96, 0.98] | -3.2% | [-4.0%, -2.3%] | 1.1 pts | 1.1 | yes | yes |
| unique-str | street | 4096 | build | ordered | baseline | 8 | 1.75 ms | 1.67 ms | 0.95× [0.94, 0.96] | -5.3% | [-6.6%, -4.0%] | 1.2 pts | 1.7 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | baseline | 12 | 66.5 | 65.6 | 0.99× [0.97, 1.01] | -1.1% | [-3.4%, +1.2%] | 2.3 pts | 2.2 | no | no |
| unique-str | street | 16384 | valuesBetween | ordered | baseline | 12 | 2144 | 2200 | 1.02× [1.01, 1.03] | +2.2% | [+1.3%, +3.1%] | 1.3 pts | 1.3 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | baseline | 12 | 631 | 643 | 1.01× [1.00, 1.03] | +1.3% | [+0.0%, +2.7%] | 1.4 pts | 0.8 | yes | yes |
| unique-str | street | 16384 | churn | ordered | baseline | 12 | 159 | 155 | 0.97× [0.97, 0.98] | -2.7% | [-3.2%, -2.2%] | 0.7 pts | 0.7 | yes | yes |
| unique-str | street | 16384 | build | ordered | baseline | 12 | 8.56 ms | 8.27 ms | 0.96× [0.96, 0.97] | -3.6% | [-3.8%, -3.5%] | 0.4 pts | 0.9 | yes | yes |
| unique-str | u64 | 4096 | valuesFor | ordered | baseline | 6 | 13.8 | 13.1 | 0.95× [0.94, 0.95] | -5.7% | [-6.5%, -4.9%] | 0.8 pts | 1.3 | yes | yes |
| unique-str | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1026 | 1080 | 1.06× [1.05, 1.07] | +5.4% | [+4.5%, +6.3%] | 0.8 pts | 1.1 | yes | yes |
| unique-str | u64 | 4096 | churn | ordered | baseline | 6 | 52.1 | 49.8 | 0.96× [0.95, 0.96] | -4.6% | [-5.0%, -4.3%] | 0.3 pts | 0.5 | yes | yes |
| unique-str | u64 | 4096 | build | ordered | baseline | 6 | 853.2 µs | 806.1 µs | 0.95× [0.94, 0.95] | -5.5% | [-6.2%, -4.7%] | 0.7 pts | 0.8 | yes | yes |
| unique-str | u64 | 16384 | valuesFor | ordered | baseline | 6 | 24.0 | 23.3 | 0.98× [0.96, 0.99] | -2.4% | [-3.9%, -0.9%] | 1.4 pts | 2.2 | yes | yes |
| unique-str | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 1559 | 1596 | 1.02× [1.00, 1.03] | +1.8% | [+0.5%, +3.1%] | 1.3 pts | 0.7 | yes | no |
| unique-str | u64 | 16384 | churn | ordered | baseline | 6 | 61.2 | 58.0 | 0.95× [0.94, 0.97] | -4.9% | [-6.4%, -3.3%] | 1.5 pts | 1.5 | yes | yes |
| unique-str | u64 | 16384 | build | ordered | baseline | 6 | 3.93 ms | 3.78 ms | 0.96× [0.95, 0.97] | -4.1% | [-4.9%, -3.4%] | 0.7 pts | 0.9 | yes | yes |
| unique-str | u64 | 262144 | valuesFor | ordered | baseline | 12 | 57.6 | 54.6 | 0.97× [0.95, 0.99] | -3.0% | [-5.3%, -0.8%] | 4.1 pts | 2.0 | no | yes |
| unique-str | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 1789 | 1820 | 1.01× [0.99, 1.03] | +0.8% | [-1.0%, +2.6%] | 3.4 pts | 1.7 | yes | no |
| unique-str | u64 | 262144 | churn | ordered | baseline | 12 | 198 | 191 | 0.96× [0.94, 0.98] | -4.5% | [-6.8%, -2.2%] | 2.7 pts | 0.8 | no | yes |
| unique-str | url | 4096 | valuesFor | ordered | baseline | 12 | 61.2 | 60.5 | 1.00× [0.97, 1.02] | -0.5% | [-2.8%, +1.9%] | 3.1 pts | 1.8 | no | no |
| unique-str | url | 4096 | valuesBetween | ordered | baseline | 12 | 2090 | 2142 | 1.02× [1.01, 1.03] | +2.1% | [+1.2%, +3.0%] | 2.8 pts | 1.2 | yes | yes |
| unique-str | url | 4096 | prefix | ordered | baseline | 12 | 143 | 148 | 1.04× [1.03, 1.05] | +3.7% | [+2.5%, +4.9%] | 1.4 pts | 1.3 | yes | yes |
| unique-str | url | 4096 | churn | ordered | baseline | 12 | 171 | 164 | 0.96× [0.96, 0.97] | -3.9% | [-4.5%, -3.2%] | 0.9 pts | 0.9 | yes | yes |
| unique-str | url | 4096 | build | ordered | baseline | 12 | 2.38 ms | 2.28 ms | 0.96× [0.95, 0.97] | -4.1% | [-5.1%, -3.0%] | 1.1 pts | 0.8 | yes | yes |
| unique-str | url | 16384 | valuesFor | ordered | baseline | 6 | 89.0 | 88.7 | 0.99× [0.99, 1.00] | -0.8% | [-1.3%, -0.4%] | 0.4 pts | 0.7 | yes | yes |
| unique-str | url | 16384 | valuesBetween | ordered | baseline | 6 | 2636 | 2668 | 1.01× [1.00, 1.01] | +0.8% | [+0.2%, +1.4%] | 0.6 pts | 0.6 | yes | yes |
| unique-str | url | 16384 | prefix | ordered | baseline | 6 | 212 | 220 | 1.04× [1.02, 1.05] | +3.5% | [+2.3%, +4.6%] | 1.1 pts | 1.3 | yes | yes |
| unique-str | url | 16384 | churn | ordered | baseline | 6 | 214 | 208 | 0.97× [0.96, 0.98] | -3.4% | [-4.6%, -2.3%] | 1.1 pts | 1.3 | yes | yes |
| unique-str | url | 16384 | build | ordered | baseline | 6 | 11.39 ms | 10.95 ms | 0.96× [0.95, 0.97] | -3.9% | [-4.9%, -2.9%] | 1.0 pts | 1.5 | yes | yes |
| unique-str | url | 262144 | valuesFor | ordered | baseline | 12 | 319 | 321 | 0.99× [0.96, 1.02] | -1.2% | [-4.0%, +1.6%] | 3.4 pts | 3.7 | no | no |
| unique-str | url | 262144 | valuesBetween | ordered | baseline | 12 | 7169 | 7101 | 1.00× [0.98, 1.02] | +0.0% | [-1.6%, +1.6%] | 1.7 pts | 1.6 | yes | no |
| unique-str | url | 262144 | prefix | ordered | baseline | 12 | 1352 | 1348 | 1.00× [0.98, 1.03] | +0.2% | [-2.1%, +2.6%] | 2.3 pts | 1.9 | no | no |
| unique-str | url | 262144 | churn | ordered | baseline | 12 | 590 | 576 | 0.97× [0.94, 1.01] | -2.7% | [-6.5%, +1.1%] | 4.3 pts | 1.2 | no | no |
| unique-str | uuid | 4096 | valuesFor | ordered | baseline | 12 | 27.2 | 26.9 | 1.00× [0.98, 1.02] | -0.5% | [-2.5%, +1.6%] | 2.8 pts | 1.0 | no | no |
| unique-str | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 1440 | 1496 | 1.03× [1.02, 1.04] | +3.3% | [+2.3%, +4.2%] | 1.0 pts | 0.7 | yes | yes |
| unique-str | uuid | 4096 | prefix | ordered | baseline | 12 | 68.7 | 68.8 | 1.00× [0.99, 1.01] | -0.1% | [-0.8%, +0.6%] | 1.0 pts | 1.2 | yes | no |
| unique-str | uuid | 4096 | churn | ordered | baseline | 12 | 89.0 | 85.5 | 0.96× [0.96, 0.96] | -4.2% | [-4.6%, -3.8%] | 0.7 pts | 0.9 | yes | yes |
| unique-str | uuid | 4096 | build | ordered | baseline | 12 | 1.26 ms | 1.21 ms | 0.95× [0.93, 0.96] | -5.4% | [-7.1%, -3.7%] | 1.9 pts | 1.9 | yes | yes |
| unique-str | uuid | 16384 | valuesFor | ordered | baseline | 8 | 42.2 | 42.3 | 1.00× [0.99, 1.01] | -0.4% | [-1.4%, +0.6%] | 1.0 pts | 1.6 | yes | no |
| unique-str | uuid | 16384 | valuesBetween | ordered | baseline | 8 | 1579 | 1625 | 1.03× [1.02, 1.03] | +2.7% | [+2.3%, +3.2%] | 0.7 pts | 0.8 | yes | yes |
| unique-str | uuid | 16384 | prefix | ordered | baseline | 8 | 83.3 | 84.3 | 1.01× [1.00, 1.02] | +1.0% | [+0.4%, +1.5%] | 0.5 pts | 1.1 | yes | yes |
| unique-str | uuid | 16384 | churn | ordered | baseline | 8 | 105 | 101 | 0.96× [0.95, 0.98] | -3.9% | [-5.7%, -2.1%] | 1.8 pts | 1.5 | yes | yes |
| unique-str | uuid | 16384 | build | ordered | baseline | 8 | 5.91 ms | 5.65 ms | 0.96× [0.96, 0.96] | -4.2% | [-4.7%, -3.7%] | 0.5 pts | 0.8 | yes | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | baseline | 12 | 161 | 156 | 0.97× [0.95, 0.98] | -3.6% | [-4.8%, -2.4%] | 1.6 pts | 1.7 | yes | yes |
| unique-str | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 3919 | 3903 | 1.00× [0.98, 1.01] | -0.3% | [-2.0%, +1.4%] | 1.8 pts | 1.6 | yes | no |
| unique-str | uuid | 262144 | prefix | ordered | baseline | 12 | 300 | 300 | 1.01× [0.99, 1.02] | +0.8% | [-0.7%, +2.2%] | 2.4 pts | 2.8 | yes | no |
| unique-str | uuid | 262144 | churn | ordered | baseline | 12 | 305 | 304 | 0.97× [0.94, 1.00] | -3.0% | [-6.1%, +0.0%] | 3.7 pts | 1.1 | no | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str email n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.48%, 2.41%] includes zero
- multi-str email n=4096 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.82%, 4.32%] includes zero
- multi-str email n=4096 prefix: ordered vs baseline: the pooled difference of -0.16% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- multi-str email n=4096 prefix: ordered vs baseline: the pooled interval [-0.87%, 0.55%] includes zero
- multi-str email n=16384 prefix: ordered vs baseline: the pooled difference of -0.47% does not clear the 0.50% noise floor, the bound on what the harness reports between identical code in every process
- multi-str email n=16384 prefix: ordered vs baseline: the pooled interval [-0.95%, 0.01%] includes zero
- multi-str email n=262144 valuesFor: ordered vs baseline: the pooled interval [-4.22%, 1.00%] includes zero
- multi-str email n=262144 valuesFor: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str email n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.05% does not clear the 0.33% noise floor, the bound on what the harness reports between identical code in every process
- multi-str email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.16%, 2.07%] includes zero
- multi-str email n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 valuesBetween: ordered vs baseline: 6 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str email n=262144 prefix: ordered vs baseline: the pooled interval [-1.56%, 2.10%] includes zero
- multi-str email n=262144 prefix: ordered vs baseline: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 prefix: ordered vs baseline: 4 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str path n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.56% does not clear the 0.85% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.92% does not clear the 0.96% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.94%, 2.78%] includes zero
- multi-str path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.08% does not clear the 0.33% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.43%, 0.60%] includes zero
- multi-str path n=16384 prefix: ordered vs baseline: the pooled difference of -0.02% does not clear the 2.89% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=16384 prefix: ordered vs baseline: the pooled interval [-1.94%, 1.89%] includes zero
- multi-str path n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.25%, 2.83%] includes zero
- multi-str path n=262144 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 valuesFor: ordered vs baseline: 6 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.95%, 2.54%] includes zero
- multi-str path n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 valuesBetween: ordered vs baseline: 5 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str path n=262144 prefix: ordered vs baseline: the pooled difference of 0.97% does not clear the 3.78% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=262144 prefix: ordered vs baseline: the pooled interval [-1.47%, 3.41%] includes zero
- multi-str str n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.21% does not clear the 0.59% noise floor, the bound on what the harness reports between identical code in every process
- multi-str str n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.35%, 0.77%] includes zero
- multi-str str n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.48% does not clear the 1.54% noise floor, the bound on what the harness reports between identical code in every process
- multi-str str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.31%, 3.26%] includes zero
- multi-str str n=4096 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.32%, 6.01%] includes zero
- multi-str str n=16384 valuesFor: ordered vs baseline: the processes scatter 5.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.26%, 1.72%] includes zero
- multi-str str n=262144 valuesFor: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.57%, 1.45%] includes zero
- multi-str str n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str str n=262144 prefix: ordered vs baseline: the pooled difference of -0.29% does not clear the 0.46% noise floor, the bound on what the harness reports between identical code in every process
- multi-str str n=262144 prefix: ordered vs baseline: the pooled interval [-2.34%, 1.76%] includes zero
- multi-str str n=262144 prefix: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 prefix: ordered vs baseline: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.63% does not clear the 0.87% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.48%, 0.21%] includes zero
- multi-str u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.56% does not clear the 1.79% noise floor, the bound on what the harness reports between identical code in every process
- multi-str u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.95%, 1.84%] includes zero
- multi-str u64 n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.13% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.75%, 0.11%] includes zero
- multi-str u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-5.19%, 1.65%] includes zero
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.04%, 3.20%] includes zero
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: 6 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.18% does not clear the 1.43% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.56%, 1.92%] includes zero
- multi-str url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.26% does not clear the 0.57% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.81%, 0.29%] includes zero
- multi-str url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.09%, 1.67%] includes zero
- multi-str url n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.65%, 2.25%] includes zero
- multi-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str url n=262144 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.13%, 1.08%] includes zero
- multi-str url n=262144 prefix: ordered vs baseline: the pooled difference of 0.51% does not clear the 2.03% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=262144 prefix: ordered vs baseline: the pooled interval [-0.95%, 1.97%] includes zero
- multi-str uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.20% does not clear the 0.92% noise floor, the bound on what the harness reports between identical code in every process
- multi-str uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.30%, 0.90%] includes zero
- multi-str uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.15% does not clear the 3.09% noise floor, the bound on what the harness reports between identical code in every process
- multi-str uuid n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -1.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.67%, 2.96%] includes zero
- multi-str uuid n=4096 prefix: ordered vs baseline: the pooled difference of 0.45% does not clear the 0.85% noise floor, the bound on what the harness reports between identical code in every process
- multi-str uuid n=4096 prefix: ordered vs baseline: the pooled interval [-0.61%, 1.50%] includes zero
- multi-str uuid n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str uuid n=16384 prefix: ordered vs baseline: the pooled interval [-3.39%, 2.67%] includes zero
- multi-str uuid n=16384 prefix: ordered vs baseline: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=16384 prefix: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.91%, 1.35%] includes zero
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 5.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.32% does not clear the 0.39% noise floor, the bound on what the harness reports between identical code in every process
- multi-str uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.54%, 2.19%] includes zero
- multi-str uuid n=262144 valuesBetween: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesBetween: ordered vs baseline: 5 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str uuid n=262144 prefix: ordered vs baseline: the pooled difference of -0.02% does not clear the 0.20% noise floor, the bound on what the harness reports between identical code in every process
- multi-str uuid n=262144 prefix: ordered vs baseline: the pooled interval [-2.09%, 2.04%] includes zero
- multi-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 prefix: ordered vs baseline: 6 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str email n=4096 valuesFor: ordered vs baseline: the pooled interval [-5.64%, 0.97%] includes zero
- unique-str email n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=4096 prefix: ordered vs baseline: the pooled interval [-1.64%, 0.37%] includes zero
- unique-str email n=4096 prefix: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str email n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.34% does not clear the 0.72% noise floor, the bound on what the harness reports between identical code in every process
- unique-str email n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.08%, 0.39%] includes zero
- unique-str email n=16384 prefix: ordered vs baseline: the pooled interval [-0.13%, 1.51%] includes zero
- unique-str email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.30%, 1.68%] includes zero
- unique-str email n=262144 prefix: ordered vs baseline: the pooled interval [-0.23%, 1.94%] includes zero
- unique-str email n=262144 prefix: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 prefix: ordered vs baseline: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str path n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.53% does not clear the 2.05% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.14%, 1.08%] includes zero
- unique-str path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.10% does not clear the 1.95% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.55%, 2.74%] includes zero
- unique-str path n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.11% does not clear the 0.96% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.36%, 1.15%] includes zero
- unique-str path n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.09% does not clear the 0.38% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.08%, 0.91%] includes zero
- unique-str path n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str path n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.44% does not clear the 0.72% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.59%, 0.72%] includes zero
- unique-str path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=262144 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str path n=262144 prefix: ordered vs baseline: the pooled difference of 0.40% does not clear the 9.43% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of +4.98% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str path n=262144 prefix: ordered vs baseline: the pooled interval [-1.52%, 2.31%] includes zero
- unique-str str n=16384 prefix: ordered vs baseline: the pooled difference of 0.74% does not clear the 1.50% noise floor, the bound on what the harness reports between identical code in every process
- unique-str str n=16384 prefix: ordered vs baseline: the pooled interval [-0.64%, 2.11%] includes zero
- unique-str str n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.96% does not clear the 1.13% noise floor, the bound on what the harness reports between identical code in every process
- unique-str str n=262144 valuesFor: ordered vs baseline: the pooled interval [-3.19%, 1.26%] includes zero
- unique-str str n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.24% does not clear the 1.11% noise floor, the bound on what the harness reports between identical code in every process
- unique-str str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.61%, 2.09%] includes zero
- unique-str str n=262144 prefix: ordered vs baseline: the pooled difference of -0.66% does not clear the 0.71% noise floor, the bound on what the harness reports between identical code in every process
- unique-str str n=262144 prefix: ordered vs baseline: the pooled interval [-2.69%, 1.37%] includes zero
- unique-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.35%, 1.21%] includes zero
- unique-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str u64 n=16384 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.19% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of 1.81% does not clear the 1.86% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.80% does not clear the 0.86% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.02%, 2.62%] includes zero
- unique-str u64 n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str url n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.46% does not clear the 0.86% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.83%, 1.90%] includes zero
- unique-str url n=262144 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.43% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str url n=262144 valuesFor: ordered vs baseline: the pooled interval [-4.02%, 1.62%] includes zero
- unique-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str url n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.02% does not clear the 0.48% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.58%, 1.62%] includes zero
- unique-str url n=262144 prefix: ordered vs baseline: the pooled difference of 0.22% does not clear the 1.19% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=262144 prefix: ordered vs baseline: the pooled interval [-2.14%, 2.59%] includes zero
- unique-str url n=262144 churn: ordered vs baseline: the pooled interval [-6.51%, 1.13%] includes zero
- unique-str uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.46% does not clear the 2.18% noise floor, the bound on what the harness reports between identical code in every process
- unique-str uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.52%, 1.61%] includes zero
- unique-str uuid n=4096 prefix: ordered vs baseline: the pooled difference of -0.12% does not clear the 0.34% noise floor, the bound on what the harness reports between identical code in every process
- unique-str uuid n=4096 prefix: ordered vs baseline: the pooled interval [-0.81%, 0.57%] includes zero
- unique-str uuid n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.42%, 0.56%] includes zero
- unique-str uuid n=16384 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.31% does not clear the 0.41% noise floor, the bound on what the harness reports between identical code in every process
- unique-str uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.03%, 1.40%] includes zero
- unique-str uuid n=262144 prefix: ordered vs baseline: the pooled interval [-0.69%, 2.22%] includes zero
- unique-str uuid n=262144 prefix: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 prefix: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str uuid n=262144 churn: ordered vs baseline: the pooled interval [-6.05%, 0.05%] includes zero
