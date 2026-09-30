| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | email | 4096 | valuesFor | ordered | baseline | 8 | 55.0 | 54.4 | 0.99× [0.97, 1.00] | -1.4% | [-2.9%, +0.1%] | 1.2 pts | 0.8 | yes | no |
| multi-str | email | 4096 | valuesBetween | ordered | baseline | 8 | 3475 | 3347 | 0.96× [0.96, 0.97] | -3.7% | [-4.6%, -2.8%] | 1.9 pts | 0.7 | yes | yes |
| multi-str | email | 4096 | prefix | ordered | baseline | 8 | 84.1 | 81.8 | 0.97× [0.97, 0.98] | -2.6% | [-3.4%, -1.9%] | 1.0 pts | 0.9 | yes | yes |
| multi-str | email | 4096 | churn | ordered | baseline | 8 | 90.7 | 87.8 | 0.97× [0.95, 0.99] | -2.9% | [-4.8%, -0.9%] | 1.2 pts | 0.8 | yes | yes |
| multi-str | email | 4096 | build | ordered | baseline | 8 | 8.95 ms | 8.68 ms | 0.97× [0.96, 0.98] | -3.0% | [-4.0%, -2.0%] | 0.8 pts | 0.9 | yes | yes |
| multi-str | email | 16384 | valuesFor | ordered | baseline | 8 | 66.5 | 65.6 | 0.98× [0.97, 1.00] | -1.6% | [-3.6%, +0.3%] | 1.4 pts | 1.1 | yes | no |
| multi-str | email | 16384 | valuesBetween | ordered | baseline | 8 | 3933 | 3848 | 0.98× [0.97, 0.98] | -2.2% | [-2.6%, -1.8%] | 0.6 pts | 0.5 | yes | yes |
| multi-str | email | 16384 | prefix | ordered | baseline | 8 | 99.9 | 98.2 | 0.98× [0.96, 1.00] | -2.2% | [-4.6%, +0.1%] | 1.7 pts | 0.9 | no | no |
| multi-str | email | 16384 | churn | ordered | baseline | 8 | 161 | 153 | 0.97× [0.96, 0.98] | -3.1% | [-4.6%, -1.6%] | 1.8 pts | 2.0 | yes | yes |
| multi-str | email | 16384 | build | ordered | baseline | 8 | 49.93 ms | 48.06 ms | 0.96× [0.95, 0.97] | -3.9% | [-5.2%, -2.7%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | email | 262144 | valuesFor | ordered | baseline | 8 | 272 | 275 | 1.01× [1.00, 1.03] | +1.4% | [-0.0%, +2.7%] | 0.9 pts | 1.2 | yes | no |
| multi-str | email | 262144 | valuesBetween | ordered | baseline | 8 | 11.0 µs | 11.0 µs | 1.00× [1.00, 1.01] | +0.3% | [-0.3%, +0.9%] | 0.5 pts | 0.9 | yes | no |
| multi-str | email | 262144 | prefix | ordered | baseline | 8 | 386 | 383 | 1.00× [0.99, 1.01] | -0.1% | [-1.2%, +1.1%] | 0.9 pts | 1.0 | yes | no |
| multi-str | email | 262144 | churn | ordered | baseline | 8 | 540 | 527 | 0.96× [0.92, 1.00] | -4.3% | [-8.3%, -0.3%] | 6.7 pts | 0.9 | no | yes |
| multi-str | path | 4096 | valuesFor | ordered | baseline | 8 | 115 | 114 | 0.99× [0.98, 1.00] | -1.0% | [-1.9%, -0.0%] | 0.9 pts | 0.9 | yes | yes |
| multi-str | path | 4096 | valuesBetween | ordered | baseline | 8 | 4915 | 4751 | 0.96× [0.95, 0.97] | -3.8% | [-4.8%, -2.8%] | 1.6 pts | 0.7 | yes | yes |
| multi-str | path | 4096 | prefix | ordered | baseline | 8 | 365 | 344 | 0.98× [0.95, 1.02] | -1.6% | [-5.1%, +2.0%] | 2.5 pts | 0.8 | no | no |
| multi-str | path | 4096 | churn | ordered | baseline | 8 | 187 | 184 | 1.01× [1.00, 1.01] | +0.6% | [+0.4%, +0.8%] | 1.1 pts | 0.6 | yes | no |
| multi-str | path | 4096 | build | ordered | baseline | 8 | 16.89 ms | 16.70 ms | 0.99× [0.98, 0.99] | -1.3% | [-1.6%, -1.0%] | 0.3 pts | 0.8 | yes | yes |
| multi-str | path | 16384 | valuesFor | ordered | baseline | 8 | 148 | 148 | 1.00× [0.98, 1.03] | +0.2% | [-2.5%, +2.8%] | 1.7 pts | 0.9 | no | no |
| multi-str | path | 16384 | valuesBetween | ordered | baseline | 8 | 5507 | 5428 | 0.99× [0.98, 1.00] | -1.4% | [-2.3%, -0.5%] | 0.6 pts | 0.4 | yes | yes |
| multi-str | path | 16384 | prefix | ordered | baseline | 8 | 803 | 833 | 1.02× [0.99, 1.05] | +1.7% | [-1.2%, +4.5%] | 3.3 pts | 1.1 | no | no |
| multi-str | path | 16384 | churn | ordered | baseline | 8 | 322 | 332 | 1.02× [1.00, 1.05] | +2.3% | [-0.4%, +4.9%] | 1.6 pts | 1.6 | no | no |
| multi-str | path | 16384 | build | ordered | baseline | 8 | 96.22 ms | 97.84 ms | 1.02× [1.01, 1.03] | +1.9% | [+1.2%, +2.7%] | 0.8 pts | 1.0 | yes | yes |
| multi-str | path | 262144 | valuesFor | ordered | baseline | 8 | 519 | 559 | 1.07× [1.06, 1.09] | +6.9% | [+5.4%, +8.4%] | 1.6 pts | 1.1 | yes | yes |
| multi-str | path | 262144 | valuesBetween | ordered | baseline | 8 | 14.6 µs | 14.9 µs | 1.02× [1.00, 1.03] | +1.6% | [+0.2%, +2.9%] | 1.2 pts | 1.8 | yes | yes |
| multi-str | path | 262144 | prefix | ordered | baseline | 8 | 12.1 µs | 12.0 µs | 1.01× [1.00, 1.03] | +1.1% | [-0.3%, +2.5%] | 1.9 pts | 1.8 | yes | no |
| multi-str | path | 262144 | churn | ordered | baseline | 8 | 942 | 968 | 0.99× [0.95, 1.03] | -0.9% | [-4.8%, +3.0%] | 2.7 pts | 0.5 | no | no |
| multi-str | str | 4096 | valuesFor | ordered | baseline | 8 | 68.7 | 67.2 | 0.98× [0.96, 1.00] | -2.1% | [-4.2%, -0.1%] | 1.5 pts | 1.4 | no | yes |
| multi-str | str | 4096 | valuesBetween | ordered | baseline | 8 | 4019 | 3830 | 0.95× [0.93, 0.96] | -5.5% | [-7.0%, -4.0%] | 1.6 pts | 0.6 | yes | yes |
| multi-str | str | 4096 | prefix | ordered | baseline | 8 | 9313 | 8710 | 0.95× [0.92, 0.99] | -4.9% | [-9.1%, -0.7%] | 2.8 pts | 0.8 | no | yes |
| multi-str | str | 4096 | churn | ordered | baseline | 8 | 108 | 107 | 0.98× [0.97, 1.00] | -1.6% | [-3.5%, +0.3%] | 1.2 pts | 0.8 | yes | no |
| multi-str | str | 4096 | build | ordered | baseline | 8 | 10.69 ms | 10.33 ms | 0.97× [0.96, 0.97] | -3.4% | [-3.8%, -2.9%] | 0.4 pts | 0.9 | yes | yes |
| multi-str | str | 16384 | valuesFor | ordered | baseline | 8 | 78.6 | 77.6 | 0.99× [0.99, 1.00] | -0.7% | [-1.4%, -0.0%] | 0.9 pts | 0.7 | yes | yes |
| multi-str | str | 16384 | valuesBetween | ordered | baseline | 8 | 4076 | 3973 | 0.97× [0.96, 0.99] | -2.8% | [-4.5%, -1.1%] | 1.1 pts | 1.0 | yes | yes |
| multi-str | str | 16384 | prefix | ordered | baseline | 8 | 38.5 µs | 37.1 µs | 0.97× [0.95, 0.98] | -3.5% | [-5.4%, -1.5%] | 1.4 pts | 1.1 | yes | yes |
| multi-str | str | 16384 | churn | ordered | baseline | 8 | 172 | 171 | 0.99× [0.97, 1.00] | -1.3% | [-2.6%, +0.1%] | 1.4 pts | 1.3 | yes | no |
| multi-str | str | 16384 | build | ordered | baseline | 8 | 54.76 ms | 54.32 ms | 0.99× [0.99, 0.99] | -1.2% | [-1.5%, -0.8%] | 0.6 pts | 0.6 | yes | yes |
| multi-str | str | 262144 | valuesFor | ordered | baseline | 8 | 301 | 313 | 1.04× [1.03, 1.05] | +4.2% | [+3.3%, +5.0%] | 0.6 pts | 0.5 | yes | yes |
| multi-str | str | 262144 | valuesBetween | ordered | baseline | 8 | 11.2 µs | 11.8 µs | 1.05× [1.04, 1.06] | +4.7% | [+3.4%, +6.1%] | 1.1 pts | 1.7 | yes | yes |
| multi-str | str | 262144 | prefix | ordered | baseline | 8 | 1.77 ms | 1.88 ms | 1.06× [1.04, 1.08] | +5.4% | [+3.4%, +7.3%] | 1.2 pts | 2.4 | yes | yes |
| multi-str | str | 262144 | churn | ordered | baseline | 8 | 574 | 570 | 0.96× [0.94, 0.99] | -4.0% | [-6.8%, -1.1%] | 3.7 pts | 0.5 | no | yes |
| multi-str | street | 4096 | valuesFor | ordered | baseline | 8 | 58.5 | 59.1 | 1.02× [0.99, 1.05] | +1.8% | [-1.0%, +4.6%] | 2.0 pts | 1.1 | no | no |
| multi-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 2701 | 2680 | 1.00× [0.98, 1.01] | -0.5% | [-2.3%, +1.4%] | 1.7 pts | 0.7 | yes | no |
| multi-str | street | 4096 | prefix | ordered | baseline | 8 | 265 | 261 | 0.99× [0.96, 1.03] | -0.5% | [-3.9%, +2.9%] | 2.7 pts | 1.2 | no | no |
| multi-str | street | 4096 | churn | ordered | baseline | 8 | 117 | 114 | 0.98× [0.97, 1.00] | -1.6% | [-3.1%, -0.0%] | 1.1 pts | 0.7 | yes | yes |
| multi-str | street | 4096 | build | ordered | baseline | 8 | 4.44 ms | 4.22 ms | 0.95× [0.94, 0.96] | -5.5% | [-6.6%, -4.3%] | 0.8 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | baseline | 8 | 80.2 | 80.6 | 1.01× [0.99, 1.02] | +0.8% | [-0.8%, +2.4%] | 1.5 pts | 1.5 | yes | no |
| multi-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 3195 | 3211 | 1.00× [0.99, 1.02] | +0.3% | [-1.3%, +1.8%] | 1.2 pts | 0.9 | yes | no |
| multi-str | street | 16384 | prefix | ordered | baseline | 8 | 1005 | 1020 | 1.00× [0.98, 1.03] | +0.4% | [-1.7%, +2.6%] | 1.5 pts | 0.6 | no | no |
| multi-str | street | 16384 | churn | ordered | baseline | 8 | 191 | 184 | 0.98× [0.98, 0.99] | -1.7% | [-2.5%, -0.9%] | 1.2 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | build | ordered | baseline | 8 | 24.86 ms | 23.83 ms | 0.96× [0.95, 0.97] | -3.7% | [-4.7%, -2.7%] | 1.0 pts | 1.5 | yes | yes |
| multi-str | u64 | 4096 | valuesFor | ordered | baseline | 8 | 42.6 | 42.3 | 1.00× [0.99, 1.01] | -0.4% | [-1.5%, +0.7%] | 0.8 pts | 0.6 | yes | no |
| multi-str | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 3023 | 2958 | 0.97× [0.95, 0.99] | -3.0% | [-4.9%, -1.1%] | 2.5 pts | 1.0 | yes | yes |
| multi-str | u64 | 4096 | churn | ordered | baseline | 8 | 67.0 | 64.6 | 0.98× [0.97, 0.98] | -2.6% | [-3.0%, -2.1%] | 0.6 pts | 0.4 | yes | yes |
| multi-str | u64 | 4096 | build | ordered | baseline | 8 | 7.14 ms | 6.90 ms | 0.97× [0.95, 0.98] | -3.6% | [-4.7%, -2.4%] | 1.0 pts | 1.8 | yes | yes |
| multi-str | u64 | 16384 | valuesFor | ordered | baseline | 4 | 49.7 | 48.8 | 0.98× [0.97, 0.98] | -2.4% | [-2.9%, -1.8%] | 0.4 pts | 0.3 | yes | yes |
| multi-str | u64 | 16384 | valuesBetween | ordered | baseline | 4 | 3784 | 3768 | 1.00× [0.98, 1.01] | -0.5% | [-2.2%, +1.2%] | 1.1 pts | 0.9 | yes | no |
| multi-str | u64 | 16384 | churn | ordered | baseline | 4 | 113 | 109 | 0.97× [0.96, 0.97] | -3.5% | [-4.2%, -2.8%] | 0.4 pts | 0.6 | yes | yes |
| multi-str | u64 | 16384 | build | ordered | baseline | 4 | 38.70 ms | 37.08 ms | 0.96× [0.95, 0.97] | -4.1% | [-5.2%, -3.0%] | 0.7 pts | 1.0 | yes | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | baseline | 8 | 166 | 165 | 1.00× [0.99, 1.01] | -0.0% | [-0.9%, +0.8%] | 1.0 pts | 1.5 | yes | no |
| multi-str | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 9387 | 9320 | 0.99× [0.99, 1.00] | -0.7% | [-1.1%, -0.3%] | 0.6 pts | 0.8 | yes | no |
| multi-str | u64 | 262144 | churn | ordered | baseline | 8 | 413 | 397 | 0.93× [0.91, 0.95] | -7.4% | [-9.8%, -5.0%] | 3.6 pts | 0.5 | no | yes |
| multi-str | url | 4096 | valuesFor | ordered | baseline | 4 | 95.7 | 96.0 | 1.01× [1.00, 1.02] | +0.8% | [+0.1%, +1.5%] | 0.5 pts | 0.5 | yes | no |
| multi-str | url | 4096 | valuesBetween | ordered | baseline | 4 | 4763 | 4602 | 0.97× [0.96, 0.98] | -3.1% | [-4.3%, -1.9%] | 0.8 pts | 0.4 | yes | yes |
| multi-str | url | 4096 | prefix | ordered | baseline | 4 | 207 | 198 | 0.96× [0.94, 0.97] | -4.4% | [-6.0%, -2.9%] | 1.0 pts | 0.6 | yes | yes |
| multi-str | url | 4096 | churn | ordered | baseline | 4 | 153 | 154 | 1.00× [0.99, 1.02] | +0.2% | [-1.2%, +1.7%] | 0.9 pts | 0.5 | yes | no |
| multi-str | url | 4096 | build | ordered | baseline | 4 | 14.25 ms | 13.91 ms | 0.98× [0.97, 0.98] | -2.6% | [-3.2%, -1.9%] | 0.4 pts | 0.8 | yes | yes |
| multi-str | url | 16384 | valuesFor | ordered | baseline | 8 | 124 | 123 | 1.00× [0.97, 1.02] | -0.3% | [-2.6%, +1.9%] | 1.9 pts | 1.0 | no | no |
| multi-str | url | 16384 | valuesBetween | ordered | baseline | 8 | 5476 | 5241 | 0.96× [0.95, 0.98] | -3.9% | [-5.5%, -2.4%] | 1.0 pts | 0.6 | yes | yes |
| multi-str | url | 16384 | prefix | ordered | baseline | 8 | 336 | 331 | 0.97× [0.95, 0.99] | -2.9% | [-5.1%, -0.6%] | 1.4 pts | 0.5 | no | yes |
| multi-str | url | 16384 | churn | ordered | baseline | 8 | 276 | 279 | 1.01× [1.01, 1.02] | +1.4% | [+1.0%, +1.9%] | 0.5 pts | 0.5 | yes | yes |
| multi-str | url | 16384 | build | ordered | baseline | 8 | 84.62 ms | 84.43 ms | 0.99× [0.98, 1.01] | -0.6% | [-1.9%, +0.6%] | 0.9 pts | 1.2 | yes | no |
| multi-str | url | 262144 | valuesFor | ordered | baseline | 8 | 479 | 509 | 1.07× [1.01, 1.13] | +6.5% | [+1.2%, +11.7%] | 3.4 pts | 2.2 | no | yes |
| multi-str | url | 262144 | valuesBetween | ordered | baseline | 8 | 14.8 µs | 14.6 µs | 0.99× [0.96, 1.01] | -1.5% | [-4.2%, +1.2%] | 2.0 pts | 2.0 | no | no |
| multi-str | url | 262144 | prefix | ordered | baseline | 8 | 3487 | 3444 | 1.00× [0.98, 1.02] | +0.1% | [-1.6%, +1.9%] | 1.9 pts | 1.0 | yes | no |
| multi-str | url | 262144 | churn | ordered | baseline | 8 | 832 | 847 | 0.99× [0.95, 1.03] | -1.4% | [-5.6%, +2.9%] | 3.2 pts | 0.6 | no | no |
| multi-str | uuid | 4096 | valuesFor | ordered | baseline | 8 | 62.0 | 59.0 | 0.95× [0.94, 0.96] | -5.0% | [-6.0%, -3.9%] | 0.8 pts | 0.7 | yes | yes |
| multi-str | uuid | 4096 | valuesBetween | ordered | baseline | 8 | 3769 | 3618 | 0.95× [0.94, 0.96] | -4.9% | [-5.9%, -3.9%] | 1.3 pts | 0.5 | yes | yes |
| multi-str | uuid | 4096 | prefix | ordered | baseline | 8 | 98.9 | 94.9 | 0.95× [0.94, 0.96] | -4.8% | [-5.9%, -3.7%] | 1.1 pts | 1.0 | yes | yes |
| multi-str | uuid | 4096 | churn | ordered | baseline | 8 | 103 | 94.7 | 0.92× [0.90, 0.94] | -8.8% | [-11.3%, -6.3%] | 2.0 pts | 1.3 | no | yes |
| multi-str | uuid | 4096 | build | ordered | baseline | 8 | 9.71 ms | 9.14 ms | 0.94× [0.94, 0.95] | -6.0% | [-6.8%, -5.1%] | 0.7 pts | 1.0 | yes | yes |
| multi-str | uuid | 16384 | valuesFor | ordered | baseline | 8 | 70.1 | 66.3 | 0.94× [0.93, 0.95] | -6.1% | [-7.0%, -5.1%] | 1.3 pts | 1.1 | yes | yes |
| multi-str | uuid | 16384 | valuesBetween | ordered | baseline | 8 | 4106 | 3863 | 0.94× [0.93, 0.96] | -6.2% | [-7.7%, -4.7%] | 1.1 pts | 0.9 | yes | yes |
| multi-str | uuid | 16384 | prefix | ordered | baseline | 8 | 119 | 114 | 0.96× [0.95, 0.96] | -4.6% | [-5.4%, -3.8%] | 1.1 pts | 0.7 | yes | yes |
| multi-str | uuid | 16384 | churn | ordered | baseline | 8 | 167 | 164 | 0.98× [0.96, 1.00] | -2.1% | [-4.0%, -0.1%] | 1.4 pts | 1.7 | yes | yes |
| multi-str | uuid | 16384 | build | ordered | baseline | 8 | 52.83 ms | 51.10 ms | 0.97× [0.95, 0.99] | -3.4% | [-5.2%, -1.5%] | 1.1 pts | 1.5 | yes | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | baseline | 8 | 292 | 290 | 1.00× [0.98, 1.02] | +0.3% | [-1.8%, +2.4%] | 1.6 pts | 1.3 | no | no |
| multi-str | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 11.6 µs | 11.3 µs | 0.97× [0.97, 0.98] | -2.7% | [-3.4%, -2.0%] | 0.8 pts | 0.8 | yes | yes |
| multi-str | uuid | 262144 | prefix | ordered | baseline | 8 | 756 | 734 | 0.98× [0.96, 0.99] | -2.5% | [-3.8%, -1.1%] | 0.8 pts | 1.1 | yes | yes |
| multi-str | uuid | 262144 | churn | ordered | baseline | 8 | 564 | 542 | 0.93× [0.87, 1.00] | -7.4% | [-14.6%, -0.2%] | 6.2 pts | 0.8 | no | yes |
| unique-str | email | 4096 | valuesFor | ordered | baseline | 8 | 30.3 | 29.2 | 0.97× [0.95, 0.99] | -3.3% | [-5.4%, -1.1%] | 2.9 pts | 1.7 | no | yes |
| unique-str | email | 4096 | valuesBetween | ordered | baseline | 8 | 1544 | 1443 | 0.93× [0.90, 0.97] | -7.2% | [-11.0%, -3.4%] | 2.5 pts | 1.5 | no | yes |
| unique-str | email | 4096 | prefix | ordered | baseline | 8 | 61.9 | 59.3 | 0.95× [0.94, 0.96] | -4.8% | [-5.9%, -3.7%] | 0.9 pts | 0.8 | yes | yes |
| unique-str | email | 4096 | churn | ordered | baseline | 8 | 85.3 | 83.1 | 0.98× [0.95, 1.00] | -2.4% | [-5.2%, +0.3%] | 2.2 pts | 1.4 | no | no |
| unique-str | email | 4096 | build | ordered | baseline | 8 | 1.19 ms | 1.12 ms | 0.93× [0.92, 0.95] | -7.0% | [-9.1%, -4.9%] | 2.7 pts | 1.3 | no | yes |
| unique-str | email | 16384 | valuesFor | ordered | baseline | 6 | 42.3 | 40.8 | 0.97× [0.96, 0.98] | -3.4% | [-4.4%, -2.3%] | 0.8 pts | 0.9 | yes | yes |
| unique-str | email | 16384 | valuesBetween | ordered | baseline | 6 | 1894 | 1777 | 0.94× [0.94, 0.95] | -6.1% | [-6.9%, -5.3%] | 0.5 pts | 0.4 | yes | yes |
| unique-str | email | 16384 | prefix | ordered | baseline | 6 | 75.9 | 72.7 | 0.96× [0.94, 0.98] | -4.2% | [-6.0%, -2.4%] | 1.1 pts | 0.7 | yes | yes |
| unique-str | email | 16384 | churn | ordered | baseline | 6 | 112 | 107 | 0.95× [0.93, 0.96] | -5.8% | [-7.0%, -4.5%] | 0.8 pts | 0.6 | yes | yes |
| unique-str | email | 16384 | build | ordered | baseline | 6 | 5.96 ms | 5.56 ms | 0.93× [0.93, 0.94] | -7.1% | [-7.6%, -6.6%] | 0.8 pts | 0.9 | yes | yes |
| unique-str | email | 262144 | valuesFor | ordered | baseline | 6 | 198 | 198 | 0.99× [0.98, 1.00] | -1.0% | [-1.9%, -0.0%] | 1.3 pts | 1.3 | yes | yes |
| unique-str | email | 262144 | valuesBetween | ordered | baseline | 6 | 5815 | 5731 | 0.98× [0.97, 1.00] | -1.8% | [-3.3%, -0.4%] | 1.6 pts | 1.5 | yes | yes |
| unique-str | email | 262144 | prefix | ordered | baseline | 6 | 273 | 273 | 0.99× [0.98, 1.01] | -0.5% | [-2.4%, +1.4%] | 1.7 pts | 1.7 | yes | no |
| unique-str | email | 262144 | churn | ordered | baseline | 6 | 382 | 372 | 0.96× [0.95, 0.98] | -3.7% | [-5.0%, -2.4%] | 2.4 pts | 0.5 | yes | yes |
| unique-str | path | 4096 | valuesFor | ordered | baseline | 8 | 90.9 | 90.1 | 0.99× [0.98, 1.01] | -0.6% | [-1.9%, +0.6%] | 1.1 pts | 0.7 | yes | no |
| unique-str | path | 4096 | valuesBetween | ordered | baseline | 8 | 2784 | 2641 | 0.95× [0.93, 0.97] | -5.5% | [-8.1%, -2.8%] | 2.1 pts | 0.7 | no | yes |
| unique-str | path | 4096 | prefix | ordered | baseline | 8 | 282 | 267 | 0.95× [0.94, 0.97] | -4.8% | [-6.5%, -3.0%] | 1.9 pts | 0.6 | yes | yes |
| unique-str | path | 4096 | churn | ordered | baseline | 8 | 230 | 221 | 0.97× [0.95, 0.98] | -3.5% | [-4.7%, -2.3%] | 1.3 pts | 0.8 | yes | yes |
| unique-str | path | 4096 | build | ordered | baseline | 8 | 2.96 ms | 2.75 ms | 0.94× [0.92, 0.95] | -6.9% | [-8.4%, -5.4%] | 1.1 pts | 1.1 | yes | yes |
| unique-str | path | 16384 | valuesFor | ordered | baseline | 8 | 123 | 122 | 1.00× [0.98, 1.01] | -0.4% | [-1.7%, +1.0%] | 0.9 pts | 0.6 | yes | no |
| unique-str | path | 16384 | valuesBetween | ordered | baseline | 8 | 3367 | 3193 | 0.95× [0.93, 0.97] | -5.1% | [-7.2%, -3.0%] | 1.4 pts | 1.1 | no | yes |
| unique-str | path | 16384 | prefix | ordered | baseline | 8 | 577 | 569 | 0.99× [0.97, 1.01] | -0.9% | [-2.6%, +0.9%] | 1.8 pts | 0.6 | yes | no |
| unique-str | path | 16384 | churn | ordered | baseline | 8 | 329 | 328 | 1.00× [0.99, 1.01] | +0.1% | [-1.1%, +1.3%] | 2.5 pts | 1.2 | yes | no |
| unique-str | path | 16384 | build | ordered | baseline | 8 | 15.06 ms | 14.64 ms | 0.97× [0.96, 0.98] | -3.3% | [-4.6%, -2.0%] | 1.1 pts | 1.2 | yes | yes |
| unique-str | path | 262144 | valuesFor | ordered | baseline | 8 | 425 | 465 | 1.11× [1.07, 1.14] | +9.6% | [+6.8%, +12.4%] | 2.0 pts | 1.5 | no | yes |
| unique-str | path | 262144 | valuesBetween | ordered | baseline | 8 | 9617 | 9533 | 1.01× [0.98, 1.04] | +0.7% | [-2.3%, +3.6%] | 2.5 pts | 2.7 | no | no |
| unique-str | path | 262144 | prefix | ordered | baseline | 8 | 6684 | 6567 | 1.00× [0.96, 1.03] | -0.5% | [-4.1%, +3.2%] | 3.9 pts | 4.0 | no | no |
| unique-str | path | 262144 | churn | ordered | baseline | 8 | 858 | 915 | 1.09× [1.05, 1.13] | +8.1% | [+5.0%, +11.3%] | 2.3 pts | 0.9 | no | yes |
| unique-str | str | 4096 | valuesFor | ordered | baseline | 8 | 42.8 | 41.5 | 0.97× [0.96, 0.99] | -2.6% | [-3.8%, -1.4%] | 1.8 pts | 1.2 | yes | yes |
| unique-str | str | 4096 | valuesBetween | ordered | baseline | 8 | 2015 | 1854 | 0.92× [0.89, 0.95] | -8.4% | [-12.1%, -4.8%] | 2.3 pts | 0.8 | no | yes |
| unique-str | str | 4096 | prefix | ordered | baseline | 8 | 3971 | 3722 | 0.95× [0.92, 0.97] | -5.4% | [-8.2%, -2.6%] | 1.9 pts | 0.5 | no | yes |
| unique-str | str | 4096 | churn | ordered | baseline | 8 | 113 | 106 | 0.93× [0.92, 0.94] | -7.7% | [-8.8%, -6.6%] | 1.7 pts | 1.0 | yes | yes |
| unique-str | str | 4096 | build | ordered | baseline | 8 | 1.60 ms | 1.43 ms | 0.90× [0.87, 0.93] | -11.5% | [-15.6%, -7.4%] | 3.2 pts | 1.1 | no | yes |
| unique-str | str | 16384 | valuesFor | ordered | baseline | 8 | 52.8 | 52.2 | 0.99× [0.97, 1.00] | -1.1% | [-2.6%, +0.3%] | 1.0 pts | 0.7 | yes | no |
| unique-str | str | 16384 | valuesBetween | ordered | baseline | 8 | 1962 | 1833 | 0.93× [0.92, 0.95] | -7.2% | [-8.7%, -5.8%] | 0.9 pts | 0.8 | yes | yes |
| unique-str | str | 16384 | prefix | ordered | baseline | 8 | 17.1 µs | 15.5 µs | 0.92× [0.90, 0.93] | -9.1% | [-10.7%, -7.6%] | 1.0 pts | 0.7 | yes | yes |
| unique-str | str | 16384 | churn | ordered | baseline | 8 | 135 | 132 | 0.98× [0.96, 0.99] | -2.4% | [-4.3%, -0.6%] | 1.8 pts | 1.2 | yes | yes |
| unique-str | str | 16384 | build | ordered | baseline | 8 | 7.34 ms | 6.87 ms | 0.94× [0.93, 0.94] | -6.5% | [-7.2%, -5.9%] | 1.1 pts | 1.1 | yes | yes |
| unique-str | str | 262144 | valuesFor | ordered | baseline | 8 | 231 | 233 | 1.01× [0.98, 1.03] | +0.8% | [-1.8%, +3.3%] | 2.2 pts | 1.5 | no | no |
| unique-str | str | 262144 | valuesBetween | ordered | baseline | 8 | 6167 | 6261 | 1.02× [1.01, 1.03] | +1.9% | [+0.6%, +3.2%] | 1.3 pts | 1.3 | yes | yes |
| unique-str | str | 262144 | prefix | ordered | baseline | 8 | 1.01 ms | 1.02 ms | 1.02× [1.00, 1.04] | +1.7% | [-0.3%, +3.6%] | 1.6 pts | 1.8 | yes | no |
| unique-str | str | 262144 | churn | ordered | baseline | 8 | 417 | 412 | 0.97× [0.93, 1.02] | -2.8% | [-7.5%, +1.9%] | 5.7 pts | 1.6 | no | no |
| unique-str | street | 4096 | valuesFor | ordered | baseline | 4 | 51.0 | 51.8 | 1.02× [1.01, 1.03] | +1.7% | [+0.6%, +2.8%] | 0.7 pts | 0.3 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | baseline | 4 | 2130 | 2137 | 1.00× [0.98, 1.02] | +0.1% | [-1.8%, +1.9%] | 1.2 pts | 0.5 | yes | no |
| unique-str | street | 4096 | prefix | ordered | baseline | 4 | 236 | 231 | 0.99× [0.97, 1.00] | -1.4% | [-2.6%, -0.2%] | 0.7 pts | 0.3 | yes | no |
| unique-str | street | 4096 | churn | ordered | baseline | 4 | 136 | 128 | 0.95× [0.94, 0.95] | -5.7% | [-6.5%, -4.9%] | 0.5 pts | 0.4 | yes | yes |
| unique-str | street | 4096 | build | ordered | baseline | 4 | 1.90 ms | 1.72 ms | 0.90× [0.89, 0.91] | -10.9% | [-11.8%, -10.0%] | 0.6 pts | 0.6 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | baseline | 8 | 72.2 | 72.7 | 1.00× [1.00, 1.01] | +0.4% | [-0.2%, +0.9%] | 0.7 pts | 0.4 | yes | no |
| unique-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 2503 | 2537 | 1.01× [1.01, 1.02] | +1.4% | [+0.5%, +2.2%] | 0.9 pts | 0.8 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | baseline | 8 | 808 | 813 | 1.01× [1.00, 1.02] | +1.1% | [+0.2%, +2.0%] | 3.5 pts | 1.6 | yes | no |
| unique-str | street | 16384 | churn | ordered | baseline | 8 | 180 | 172 | 0.96× [0.94, 0.98] | -4.1% | [-6.0%, -2.2%] | 1.7 pts | 1.0 | yes | yes |
| unique-str | street | 16384 | build | ordered | baseline | 8 | 9.33 ms | 8.66 ms | 0.93× [0.92, 0.94] | -7.4% | [-8.4%, -6.5%] | 0.8 pts | 1.3 | yes | yes |
| unique-str | u64 | 4096 | valuesFor | ordered | baseline | 8 | 18.0 | 17.1 | 0.96× [0.95, 0.97] | -3.8% | [-5.1%, -2.6%] | 3.2 pts | 1.4 | yes | yes |
| unique-str | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 1170 | 1114 | 0.94× [0.94, 0.95] | -5.8% | [-6.5%, -5.1%] | 0.7 pts | 0.5 | yes | yes |
| unique-str | u64 | 4096 | churn | ordered | baseline | 8 | 56.7 | 54.2 | 0.94× [0.93, 0.96] | -6.1% | [-8.0%, -4.1%] | 1.7 pts | 1.0 | yes | yes |
| unique-str | u64 | 4096 | build | ordered | baseline | 8 | 881.6 µs | 832.5 µs | 0.93× [0.91, 0.95] | -7.2% | [-9.7%, -4.7%] | 2.0 pts | 1.0 | no | yes |
| unique-str | u64 | 16384 | valuesFor | ordered | baseline | 6 | 25.8 | 24.5 | 0.95× [0.93, 0.97] | -5.1% | [-7.0%, -3.2%] | 1.4 pts | 1.4 | yes | yes |
| unique-str | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 1719 | 1665 | 0.97× [0.96, 0.98] | -3.2% | [-4.6%, -1.9%] | 0.9 pts | 0.7 | yes | yes |
| unique-str | u64 | 16384 | churn | ordered | baseline | 6 | 71.4 | 66.6 | 0.93× [0.92, 0.95] | -7.1% | [-8.9%, -5.3%] | 1.6 pts | 0.8 | yes | yes |
| unique-str | u64 | 16384 | build | ordered | baseline | 6 | 4.21 ms | 3.97 ms | 0.94× [0.93, 0.95] | -6.4% | [-7.7%, -5.1%] | 1.1 pts | 1.5 | yes | yes |
| unique-str | u64 | 262144 | valuesFor | ordered | baseline | 8 | 99.7 | 98.0 | 0.99× [0.97, 1.01] | -0.9% | [-2.8%, +1.1%] | 1.4 pts | 1.7 | yes | no |
| unique-str | u64 | 262144 | valuesBetween | ordered | baseline | 8 | 4633 | 4533 | 0.98× [0.97, 0.99] | -2.4% | [-3.6%, -1.2%] | 1.8 pts | 2.0 | yes | yes |
| unique-str | u64 | 262144 | churn | ordered | baseline | 8 | 266 | 260 | 0.96× [0.92, 1.00] | -4.3% | [-8.6%, -0.0%] | 3.8 pts | 1.8 | no | yes |
| unique-str | url | 4096 | valuesFor | ordered | baseline | 8 | 71.7 | 72.2 | 1.00× [0.99, 1.02] | +0.3% | [-0.9%, +1.6%] | 1.1 pts | 0.8 | yes | no |
| unique-str | url | 4096 | valuesBetween | ordered | baseline | 8 | 2669 | 2515 | 0.94× [0.93, 0.96] | -5.9% | [-8.1%, -3.7%] | 1.6 pts | 0.6 | no | yes |
| unique-str | url | 4096 | prefix | ordered | baseline | 8 | 167 | 158 | 0.95× [0.93, 0.97] | -5.3% | [-7.8%, -2.8%] | 1.7 pts | 0.9 | no | yes |
| unique-str | url | 4096 | churn | ordered | baseline | 8 | 183 | 180 | 0.97× [0.96, 0.99] | -3.0% | [-4.6%, -1.5%] | 1.0 pts | 0.6 | yes | yes |
| unique-str | url | 4096 | build | ordered | baseline | 8 | 2.43 ms | 2.27 ms | 0.93× [0.90, 0.95] | -7.9% | [-10.7%, -5.2%] | 1.9 pts | 1.2 | no | yes |
| unique-str | url | 16384 | valuesFor | ordered | baseline | 8 | 96.3 | 98.3 | 1.01× [0.99, 1.04] | +1.4% | [-0.9%, +3.7%] | 2.0 pts | 1.4 | no | no |
| unique-str | url | 16384 | valuesBetween | ordered | baseline | 8 | 3237 | 3041 | 0.95× [0.93, 0.96] | -5.7% | [-7.2%, -4.2%] | 1.0 pts | 1.0 | yes | yes |
| unique-str | url | 16384 | prefix | ordered | baseline | 8 | 251 | 246 | 0.97× [0.96, 0.98] | -3.3% | [-4.3%, -2.3%] | 1.0 pts | 0.5 | yes | yes |
| unique-str | url | 16384 | churn | ordered | baseline | 8 | 273 | 275 | 1.01× [0.99, 1.03] | +0.8% | [-1.1%, +2.7%] | 1.7 pts | 0.8 | yes | no |
| unique-str | url | 16384 | build | ordered | baseline | 8 | 12.50 ms | 11.88 ms | 0.95× [0.94, 0.96] | -5.1% | [-6.2%, -4.0%] | 0.7 pts | 0.8 | yes | yes |
| unique-str | url | 262144 | valuesFor | ordered | baseline | 8 | 407 | 442 | 1.09× [1.05, 1.14] | +8.5% | [+5.1%, +11.9%] | 3.4 pts | 2.7 | no | yes |
| unique-str | url | 262144 | valuesBetween | ordered | baseline | 8 | 9492 | 9436 | 0.98× [0.96, 1.00] | -2.0% | [-3.9%, -0.0%] | 1.8 pts | 2.1 | yes | yes |
| unique-str | url | 262144 | prefix | ordered | baseline | 8 | 2061 | 2083 | 1.02× [1.00, 1.04] | +1.8% | [-0.3%, +3.9%] | 2.2 pts | 1.1 | no | no |
| unique-str | url | 262144 | churn | ordered | baseline | 8 | 774 | 801 | 1.02× [1.00, 1.04] | +2.0% | [-0.2%, +4.1%] | 1.8 pts | 0.7 | no | no |
| unique-str | uuid | 4096 | valuesFor | ordered | baseline | 8 | 36.7 | 33.8 | 0.92× [0.91, 0.94] | -8.4% | [-10.0%, -6.7%] | 1.0 pts | 0.6 | yes | yes |
| unique-str | uuid | 4096 | valuesBetween | ordered | baseline | 8 | 1786 | 1645 | 0.92× [0.91, 0.94] | -8.4% | [-10.3%, -6.5%] | 1.4 pts | 0.7 | yes | yes |
| unique-str | uuid | 4096 | prefix | ordered | baseline | 8 | 75.9 | 71.1 | 0.94× [0.93, 0.95] | -6.5% | [-7.4%, -5.7%] | 0.8 pts | 0.7 | yes | yes |
| unique-str | uuid | 4096 | churn | ordered | baseline | 8 | 101 | 92.9 | 0.92× [0.91, 0.94] | -8.2% | [-9.6%, -6.7%] | 1.6 pts | 0.8 | yes | yes |
| unique-str | uuid | 4096 | build | ordered | baseline | 8 | 1.34 ms | 1.24 ms | 0.92× [0.90, 0.94] | -8.4% | [-10.5%, -6.2%] | 1.6 pts | 0.7 | no | yes |
| unique-str | uuid | 16384 | valuesFor | ordered | baseline | 8 | 45.7 | 41.4 | 0.90× [0.89, 0.91] | -10.8% | [-11.8%, -9.9%] | 0.8 pts | 0.8 | yes | yes |
| unique-str | uuid | 16384 | valuesBetween | ordered | baseline | 8 | 2021 | 1746 | 0.87× [0.86, 0.87] | -15.4% | [-15.8%, -14.9%] | 1.2 pts | 1.2 | yes | yes |
| unique-str | uuid | 16384 | prefix | ordered | baseline | 8 | 90.5 | 84.6 | 0.94× [0.92, 0.95] | -6.9% | [-8.6%, -5.2%] | 1.0 pts | 1.0 | yes | yes |
| unique-str | uuid | 16384 | churn | ordered | baseline | 8 | 123 | 118 | 0.97× [0.96, 0.99] | -2.8% | [-4.7%, -0.9%] | 1.7 pts | 0.9 | yes | yes |
| unique-str | uuid | 16384 | build | ordered | baseline | 8 | 6.50 ms | 5.99 ms | 0.92× [0.91, 0.93] | -9.0% | [-10.2%, -7.8%] | 1.2 pts | 1.6 | yes | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | baseline | 8 | 215 | 216 | 1.00× [0.98, 1.02] | +0.2% | [-1.9%, +2.4%] | 1.4 pts | 1.1 | no | no |
| unique-str | uuid | 262144 | valuesBetween | ordered | baseline | 8 | 6252 | 6186 | 0.99× [0.97, 1.00] | -1.5% | [-2.8%, -0.2%] | 0.9 pts | 1.0 | yes | yes |
| unique-str | uuid | 262144 | prefix | ordered | baseline | 8 | 445 | 431 | 0.97× [0.96, 0.99] | -2.8% | [-4.3%, -1.3%] | 1.2 pts | 1.5 | yes | yes |
| unique-str | uuid | 262144 | churn | ordered | baseline | 8 | 388 | 383 | 0.97× [0.95, 0.99] | -3.0% | [-5.3%, -0.7%] | 2.3 pts | 0.5 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str email n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str email n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.86%, 0.09%] includes zero
- multi-str email n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.56%, 0.34%] includes zero
- multi-str email n=16384 prefix: ordered vs baseline: the pooled interval [-4.57%, 0.12%] includes zero
- multi-str email n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.03%, 2.73%] includes zero
- multi-str email n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.30% does not clear the 1.04% noise floor, the bound on what the harness reports between identical code in every process
- multi-str email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.32%, 0.92%] includes zero
- multi-str email n=262144 prefix: ordered vs baseline: the pooled difference of -0.06% does not clear the 0.48% noise floor, the bound on what the harness reports between identical code in every process
- multi-str email n=262144 prefix: ordered vs baseline: the pooled interval [-1.20%, 1.09%] includes zero
- multi-str path n=4096 prefix: ordered vs baseline: the pooled difference of -1.57% does not clear the 4.05% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=4096 prefix: ordered vs baseline: the pooled interval [-5.11%, 1.97%] includes zero
- multi-str path n=4096 churn: ordered vs baseline: the pooled difference of 0.59% does not clear the 1.11% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.18% does not clear the 0.56% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.46%, 2.82%] includes zero
- multi-str path n=16384 prefix: ordered vs baseline: the pooled difference of 1.68% does not clear the 3.22% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=16384 prefix: ordered vs baseline: the pooled interval [-1.18%, 4.54%] includes zero
- multi-str path n=16384 churn: ordered vs baseline: the pooled interval [-0.36%, 4.91%] includes zero
- multi-str path n=262144 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str path n=262144 prefix: ordered vs baseline: the pooled difference of 1.13% does not clear the 3.79% noise floor, the bound on what the harness reports between identical code in every process
- multi-str path n=262144 prefix: ordered vs baseline: the A/A validations found a systematic difference of -2.06% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str path n=262144 prefix: ordered vs baseline: the pooled interval [-0.29%, 2.55%] includes zero
- multi-str path n=262144 churn: ordered vs baseline: the pooled interval [-4.83%, 3.01%] includes zero
- multi-str str n=4096 prefix: ordered vs baseline: the A/A validations found a systematic difference of -0.68% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str str n=4096 churn: ordered vs baseline: the pooled interval [-3.47%, 0.30%] includes zero
- multi-str str n=16384 churn: ordered vs baseline: the pooled interval [-2.61%, 0.08%] includes zero
- multi-str str n=262144 prefix: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.76% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.95%, 4.64%] includes zero
- multi-str street n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.45% does not clear the 1.11% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.34%, 1.43%] includes zero
- multi-str street n=4096 prefix: ordered vs baseline: the pooled difference of -0.53% does not clear the 1.49% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 prefix: ordered vs baseline: the pooled interval [-3.92%, 2.86%] includes zero
- multi-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.82%, 2.42%] includes zero
- multi-str street n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.25% does not clear the 0.59% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.34%, 1.84%] includes zero
- multi-str street n=16384 prefix: ordered vs baseline: the pooled difference of 0.42% does not clear the 1.30% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 prefix: ordered vs baseline: the pooled interval [-1.73%, 2.58%] includes zero
- multi-str u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.37% does not clear the 0.80% noise floor, the bound on what the harness reports between identical code in every process
- multi-str u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.45%, 0.71%] includes zero
- multi-str u64 n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str u64 n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.50% does not clear the 0.74% noise floor, the bound on what the harness reports between identical code in every process
- multi-str u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-2.17%, 1.18%] includes zero
- multi-str u64 n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str u64 n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.05% does not clear the 0.69% noise floor, the bound on what the harness reports between identical code in every process
- multi-str u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.95%, 0.85%] includes zero
- multi-str u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.69% does not clear the 0.71% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.79% does not clear the 1.02% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str url n=4096 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str url n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str url n=4096 churn: ordered vs baseline: the pooled difference of 0.23% does not clear the 2.23% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=4096 churn: ordered vs baseline: the pooled interval [-1.22%, 1.68%] includes zero
- multi-str url n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- multi-str url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.34% does not clear the 0.61% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.59%, 1.90%] includes zero
- multi-str url n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str url n=16384 build: ordered vs baseline: the pooled interval [-1.94%, 0.65%] includes zero
- multi-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-4.17%, 1.20%] includes zero
- multi-str url n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 prefix: ordered vs baseline: the pooled difference of 0.15% does not clear the 3.54% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=262144 prefix: ordered vs baseline: the pooled interval [-1.57%, 1.87%] includes zero
- multi-str url n=262144 churn: ordered vs baseline: the pooled difference of -1.39% does not clear the 1.42% noise floor, the bound on what the harness reports between identical code in every process
- multi-str url n=262144 churn: ordered vs baseline: the pooled interval [-5.64%, 2.87%] includes zero
- multi-str uuid n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.27% does not clear the 0.57% noise floor, the bound on what the harness reports between identical code in every process
- multi-str uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.85%, 2.40%] includes zero
- unique-str email n=4096 churn: ordered vs baseline: the pooled interval [-5.15%, 0.27%] includes zero
- unique-str email n=262144 prefix: ordered vs baseline: the pooled difference of -0.50% does not clear the 0.60% noise floor, the bound on what the harness reports between identical code in every process
- unique-str email n=262144 prefix: ordered vs baseline: the pooled interval [-2.38%, 1.37%] includes zero
- unique-str email n=262144 prefix: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str path n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.62% does not clear the 0.82% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.85%, 0.61%] includes zero
- unique-str path n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of -0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str path n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.36% does not clear the 1.00% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.68%, 0.96%] includes zero
- unique-str path n=16384 prefix: ordered vs baseline: the pooled difference of -0.88% does not clear the 6.37% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=16384 prefix: ordered vs baseline: the pooled interval [-2.63%, 0.87%] includes zero
- unique-str path n=16384 churn: ordered vs baseline: the pooled difference of 0.09% does not clear the 0.75% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=16384 churn: ordered vs baseline: the pooled interval [-1.12%, 1.31%] includes zero
- unique-str path n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.66% does not clear the 0.69% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-2.28%, 3.60%] includes zero
- unique-str path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str path n=262144 prefix: ordered vs baseline: the pooled difference of -0.47% does not clear the 7.10% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=262144 prefix: ordered vs baseline: the pooled interval [-4.15%, 3.21%] includes zero
- unique-str path n=262144 prefix: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.57%, 0.32%] includes zero
- unique-str str n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.80%, 3.34%] includes zero
- unique-str str n=262144 prefix: ordered vs baseline: the pooled interval [-0.25%, 3.60%] includes zero
- unique-str str n=262144 churn: ordered vs baseline: the pooled interval [-7.52%, 1.87%] includes zero
- unique-str street n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str street n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str street n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.07% does not clear the 0.70% noise floor, the bound on what the harness reports between identical code in every process
- unique-str street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.77%, 1.90%] includes zero
- unique-str street n=4096 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str street n=4096 prefix: ordered vs baseline: the pooled difference of -1.40% does not clear the 2.27% noise floor, the bound on what the harness reports between identical code in every process
- unique-str street n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str street n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str street n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.37% does not clear the 0.91% noise floor, the bound on what the harness reports between identical code in every process
- unique-str street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.19%, 0.94%] includes zero
- unique-str street n=16384 prefix: ordered vs baseline: the pooled difference of 1.10% does not clear the 1.81% noise floor, the bound on what the harness reports between identical code in every process
- unique-str u64 n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.80%, 1.07%] includes zero
- unique-str u64 n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.31% does not clear the 0.75% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.93%, 1.55%] includes zero
- unique-str url n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.92%, 3.69%] includes zero
- unique-str url n=16384 churn: ordered vs baseline: the pooled difference of 0.80% does not clear the 0.90% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=16384 churn: ordered vs baseline: the pooled interval [-1.05%, 2.66%] includes zero
- unique-str url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 prefix: ordered vs baseline: the pooled difference of 1.81% does not clear the 5.32% noise floor, the bound on what the harness reports between identical code in every process
- unique-str url n=262144 prefix: ordered vs baseline: the pooled interval [-0.28%, 3.91%] includes zero
- unique-str url n=262144 churn: ordered vs baseline: the pooled interval [-0.16%, 4.12%] includes zero
- unique-str uuid n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.24% does not clear the 0.30% noise floor, the bound on what the harness reports between identical code in every process
- unique-str uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.89%, 2.38%] includes zero
- unique-str uuid n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
