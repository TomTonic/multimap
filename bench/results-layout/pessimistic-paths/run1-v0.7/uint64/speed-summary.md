| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 6 | 50.6 | 49.1 | 0.97× [0.97, 0.98] | -2.8% | [-3.5%, -2.0%] | 0.7 pts | 0.5 | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 6 | 3020 | 3002 | 0.99× [0.98, 1.01] | -0.6% | [-2.0%, +0.8%] | 1.3 pts | 0.6 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 6 | 79.5 | 76.4 | 0.96× [0.95, 0.97] | -4.5% | [-5.3%, -3.6%] | 0.8 pts | 0.8 | yes |
| multi | email | 4096 | churn | ordered | baseline | 6 | 69.7 | 65.9 | 0.95× [0.94, 0.96] | -5.5% | [-6.5%, -4.4%] | 1.0 pts | 0.6 | yes |
| multi | email | 4096 | build | ordered | baseline | 6 | 5.73 ms | 5.58 ms | 0.98× [0.96, 1.00] | -2.2% | [-4.1%, -0.3%] | 1.8 pts | 1.2 | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 62.6 | 60.8 | 0.97× [0.96, 0.98] | -3.4% | [-4.3%, -2.5%] | 0.9 pts | 0.7 | yes |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3551 | 3513 | 0.99× [0.98, 1.00] | -1.0% | [-1.5%, -0.5%] | 0.5 pts | 0.4 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 95.5 | 93.0 | 0.97× [0.96, 0.98] | -3.0% | [-4.2%, -1.8%] | 1.1 pts | 0.9 | yes |
| multi | email | 16384 | churn | ordered | baseline | 6 | 102 | 97.8 | 0.96× [0.95, 0.97] | -4.5% | [-5.6%, -3.4%] | 1.0 pts | 0.6 | yes |
| multi | email | 16384 | build | ordered | baseline | 6 | 29.65 ms | 28.44 ms | 0.96× [0.95, 0.97] | -4.6% | [-5.6%, -3.6%] | 1.0 pts | 1.3 | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 10 | 268 | 262 | 0.99× [0.98, 1.01] | -0.7% | [-1.9%, +0.5%] | 1.7 pts | 1.7 | yes |
| multi | email | 262144 | valuesBetween | ordered | baseline | 10 | 8227 | 8199 | 0.99× [0.99, 1.00] | -0.6% | [-1.4%, +0.2%] | 1.1 pts | 1.0 | yes |
| multi | email | 262144 | prefix | ordered | baseline | 10 | 328 | 323 | 1.00× [0.98, 1.02] | -0.2% | [-2.0%, +1.6%] | 2.5 pts | 2.2 | yes |
| multi | email | 262144 | churn | ordered | baseline | 10 | 450 | 429 | 0.99× [0.93, 1.05] | -1.3% | [-7.2%, +4.6%] | 8.3 pts | 2.4 | no |
| multi | email | 1048576 | valuesFor | ordered | baseline | 10 | 414 | 412 | 1.00× [0.99, 1.00] | -0.4% | [-1.0%, +0.2%] | 0.9 pts | 0.8 | yes |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 10 | 10.3 µs | 10.2 µs | 0.99× [0.99, 1.00] | -0.8% | [-1.3%, -0.3%] | 0.7 pts | 0.7 | yes |
| multi | email | 1048576 | prefix | ordered | baseline | 10 | 644 | 646 | 1.00× [1.00, 1.01] | +0.3% | [-0.0%, +0.7%] | 0.5 pts | 0.4 | yes |
| multi | email | 1048576 | churn | ordered | baseline | 10 | 648 | 638 | 0.97× [0.93, 1.02] | -2.7% | [-7.5%, +2.1%] | 6.7 pts | 1.3 | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 112 | 108 | 0.96× [0.95, 0.97] | -4.2% | [-5.0%, -3.5%] | 0.7 pts | 0.7 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4264 | 4151 | 0.97× [0.96, 0.99] | -2.7% | [-4.0%, -1.3%] | 1.3 pts | 0.7 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 338 | 336 | 0.99× [0.98, 1.00] | -0.8% | [-1.9%, +0.3%] | 1.0 pts | 0.4 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 164 | 160 | 0.97× [0.96, 0.98] | -2.8% | [-3.7%, -2.0%] | 0.8 pts | 0.6 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 12.19 ms | 11.80 ms | 0.96× [0.96, 0.97] | -3.8% | [-4.5%, -3.1%] | 0.7 pts | 0.8 | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 8 | 141 | 138 | 0.98× [0.97, 0.99] | -1.8% | [-2.7%, -1.0%] | 1.0 pts | 0.8 | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 8 | 4762 | 4769 | 1.00× [1.00, 1.01] | +0.4% | [-0.4%, +1.2%] | 1.0 pts | 0.7 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 8 | 636 | 625 | 1.00× [0.98, 1.01] | -0.3% | [-1.9%, +1.4%] | 2.0 pts | 0.4 | yes |
| multi | path | 16384 | churn | ordered | baseline | 8 | 252 | 255 | 1.01× [1.01, 1.02] | +1.2% | [+0.6%, +1.9%] | 0.7 pts | 0.3 | yes |
| multi | path | 16384 | build | ordered | baseline | 8 | 67.73 ms | 66.59 ms | 0.98× [0.97, 0.99] | -1.9% | [-2.6%, -1.1%] | 0.9 pts | 0.7 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 539 | 565 | 1.04× [1.04, 1.05] | +4.1% | [+3.4%, +4.8%] | 0.9 pts | 0.7 | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 12.4 µs | 13.3 µs | 1.06× [1.05, 1.07] | +5.8% | [+4.9%, +6.7%] | 1.3 pts | 0.8 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 13.2 µs | 13.8 µs | 1.04× [1.02, 1.07] | +4.0% | [+1.8%, +6.2%] | 3.1 pts | 0.2 | no |
| multi | path | 262144 | churn | ordered | baseline | 10 | 928 | 901 | 0.97× [0.95, 1.00] | -2.8% | [-5.3%, -0.4%] | 3.4 pts | 1.4 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 8 | 61.8 | 60.7 | 0.98× [0.98, 0.99] | -1.6% | [-2.2%, -1.0%] | 0.7 pts | 0.6 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 8 | 3547 | 3545 | 1.00× [0.98, 1.01] | -0.2% | [-1.7%, +1.3%] | 1.9 pts | 0.8 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 8 | 8121 | 7996 | 0.99× [0.97, 1.00] | -1.3% | [-2.8%, +0.2%] | 1.8 pts | 0.7 | yes |
| multi | str | 4096 | churn | ordered | baseline | 8 | 84.3 | 80.1 | 0.95× [0.94, 0.95] | -5.7% | [-6.3%, -5.1%] | 0.8 pts | 0.4 | yes |
| multi | str | 4096 | build | ordered | baseline | 8 | 6.87 ms | 6.56 ms | 0.95× [0.95, 0.96] | -4.7% | [-5.7%, -3.8%] | 1.1 pts | 1.1 | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 72.6 | 71.7 | 0.99× [0.98, 1.00] | -1.2% | [-2.1%, -0.4%] | 0.8 pts | 0.8 | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3702 | 3652 | 0.99× [0.98, 1.00] | -1.1% | [-2.1%, -0.2%] | 0.9 pts | 0.7 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 34.7 µs | 34.3 µs | 0.99× [0.98, 1.00] | -1.0% | [-2.1%, +0.2%] | 1.1 pts | 0.8 | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 122 | 118 | 0.97× [0.96, 0.98] | -3.0% | [-4.0%, -2.1%] | 0.9 pts | 0.5 | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 35.31 ms | 34.25 ms | 0.97× [0.96, 0.98] | -3.1% | [-4.3%, -1.9%] | 1.1 pts | 0.9 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 327 | 331 | 1.02× [1.01, 1.03] | +2.1% | [+0.9%, +3.3%] | 1.7 pts | 1.2 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 9301 | 9537 | 1.02× [1.01, 1.04] | +2.4% | [+1.4%, +3.4%] | 1.4 pts | 1.2 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.49 ms | 1.54 ms | 1.03× [1.02, 1.04] | +2.7% | [+1.8%, +3.5%] | 1.2 pts | 1.2 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 493 | 501 | 0.98× [0.95, 1.00] | -2.6% | [-5.6%, +0.5%] | 4.2 pts | 1.4 | no |
| multi | str | 1048576 | valuesFor | ordered | baseline | 10 | 441 | 448 | 1.02× [1.00, 1.04] | +2.1% | [+0.1%, +4.1%] | 2.8 pts | 1.7 | yes |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 10 | 10.4 µs | 10.7 µs | 1.03× [1.02, 1.05] | +3.2% | [+2.1%, +4.4%] | 1.6 pts | 1.5 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 10 | 6.32 ms | 6.49 ms | 1.03× [1.02, 1.04] | +2.9% | [+2.1%, +3.7%] | 1.1 pts | 1.8 | yes |
| multi | str | 1048576 | churn | ordered | baseline | 10 | 717 | 700 | 1.02× [0.96, 1.07] | +1.6% | [-3.7%, +6.9%] | 7.4 pts | 2.7 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 6 | 56.3 | 53.9 | 0.96× [0.96, 0.97] | -4.0% | [-4.5%, -3.4%] | 0.5 pts | 0.3 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 6 | 2491 | 2456 | 0.98× [0.97, 1.00] | -1.6% | [-2.8%, -0.4%] | 1.2 pts | 0.6 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 6 | 276 | 262 | 0.95× [0.93, 0.96] | -5.8% | [-7.4%, -4.2%] | 1.5 pts | 0.8 | yes |
| multi | street | 4096 | churn | ordered | baseline | 6 | 99.8 | 94.2 | 0.95× [0.94, 0.95] | -5.8% | [-6.7%, -4.9%] | 0.9 pts | 0.6 | yes |
| multi | street | 4096 | build | ordered | baseline | 6 | 3.16 ms | 2.96 ms | 0.94× [0.92, 0.95] | -6.9% | [-8.5%, -5.4%] | 1.4 pts | 1.3 | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 80.8 | 77.5 | 0.96× [0.95, 0.97] | -4.2% | [-5.7%, -2.6%] | 1.5 pts | 1.0 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2987 | 2933 | 0.98× [0.97, 0.99] | -2.1% | [-3.2%, -0.9%] | 1.1 pts | 0.9 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 1001 | 963 | 0.97× [0.96, 0.98] | -3.3% | [-4.1%, -2.5%] | 0.8 pts | 0.2 | yes |
| multi | street | 16384 | churn | ordered | baseline | 6 | 141 | 132 | 0.95× [0.94, 0.96] | -5.1% | [-6.5%, -3.7%] | 1.3 pts | 0.8 | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 17.68 ms | 16.70 ms | 0.94× [0.93, 0.95] | -6.7% | [-7.7%, -5.8%] | 0.9 pts | 0.6 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 8 | 37.8 | 36.5 | 0.97× [0.95, 0.98] | -3.2% | [-4.7%, -1.8%] | 1.8 pts | 1.0 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2759 | 2678 | 0.97× [0.96, 0.99] | -2.8% | [-4.3%, -1.3%] | 1.8 pts | 0.8 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 8 | 50.2 | 47.8 | 0.95× [0.94, 0.95] | -5.4% | [-5.9%, -4.9%] | 0.6 pts | 0.4 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 8 | 4.43 ms | 4.24 ms | 0.96× [0.95, 0.96] | -4.6% | [-5.0%, -4.3%] | 0.5 pts | 0.8 | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 48.5 | 46.4 | 0.96× [0.96, 0.97] | -4.0% | [-4.6%, -3.3%] | 0.6 pts | 0.7 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3594 | 3545 | 0.99× [0.98, 0.99] | -1.3% | [-2.0%, -0.6%] | 0.7 pts | 0.5 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 68.9 | 62.8 | 0.92× [0.91, 0.92] | -9.1% | [-10.0%, -8.3%] | 0.8 pts | 0.4 | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 23.00 ms | 21.71 ms | 0.95× [0.94, 0.95] | -5.7% | [-6.2%, -5.3%] | 0.4 pts | 0.5 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 165 | 156 | 0.96× [0.94, 0.97] | -4.5% | [-6.3%, -2.7%] | 2.5 pts | 2.2 | yes |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 7652 | 7677 | 0.99× [0.98, 1.00] | -1.0% | [-2.0%, -0.0%] | 1.4 pts | 1.1 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 328 | 327 | 1.01× [0.96, 1.06] | +0.6% | [-4.7%, +5.8%] | 7.3 pts | 2.5 | no |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 221 | 207 | 0.94× [0.94, 0.95] | -6.2% | [-6.8%, -5.6%] | 0.9 pts | 0.9 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 6889 | 6770 | 0.99× [0.98, 0.99] | -1.5% | [-2.1%, -0.9%] | 0.8 pts | 0.7 | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 10 | 413 | 382 | 0.93× [0.89, 0.97] | -7.5% | [-12.4%, -2.6%] | 6.9 pts | 1.0 | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 6 | 90.9 | 88.7 | 0.98× [0.97, 0.99] | -2.3% | [-3.4%, -1.2%] | 1.0 pts | 1.0 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 6 | 4009 | 3918 | 0.98× [0.97, 0.99] | -2.0% | [-3.1%, -1.0%] | 1.0 pts | 0.6 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 6 | 192 | 188 | 0.98× [0.98, 0.98] | -2.1% | [-2.4%, -1.8%] | 0.3 pts | 0.2 | yes |
| multi | url | 4096 | churn | ordered | baseline | 6 | 133 | 129 | 0.96× [0.95, 0.97] | -3.9% | [-4.8%, -3.0%] | 0.8 pts | 0.5 | yes |
| multi | url | 4096 | build | ordered | baseline | 6 | 9.96 ms | 9.54 ms | 0.96× [0.95, 0.97] | -4.3% | [-5.4%, -3.3%] | 1.0 pts | 1.3 | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 116 | 115 | 0.99× [0.98, 1.00] | -1.0% | [-2.1%, -0.0%] | 1.0 pts | 0.7 | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4548 | 4503 | 0.99× [0.99, 1.00] | -1.0% | [-1.5%, -0.5%] | 0.5 pts | 0.4 | yes |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 280 | 278 | 0.99× [0.98, 1.00] | -1.2% | [-2.2%, -0.2%] | 0.9 pts | 0.5 | yes |
| multi | url | 16384 | churn | ordered | baseline | 6 | 200 | 197 | 0.99× [0.98, 1.00] | -1.4% | [-2.3%, -0.5%] | 0.8 pts | 0.8 | yes |
| multi | url | 16384 | build | ordered | baseline | 6 | 56.30 ms | 54.71 ms | 0.97× [0.96, 0.98] | -3.2% | [-4.2%, -2.1%] | 1.0 pts | 0.8 | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 10 | 451 | 466 | 1.01× [0.98, 1.05] | +1.3% | [-1.9%, +4.4%] | 4.4 pts | 3.8 | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 10 | 12.0 µs | 12.4 µs | 1.03× [1.02, 1.04] | +2.6% | [+1.9%, +3.4%] | 1.1 pts | 0.8 | yes |
| multi | url | 262144 | prefix | ordered | baseline | 10 | 2273 | 2341 | 1.03× [1.01, 1.05] | +2.5% | [+0.6%, +4.4%] | 2.6 pts | 0.5 | yes |
| multi | url | 262144 | churn | ordered | baseline | 10 | 776 | 770 | 1.01× [0.96, 1.07] | +1.4% | [-3.7%, +6.6%] | 7.2 pts | 2.3 | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 10 | 680 | 715 | 1.02× [1.01, 1.04] | +2.3% | [+0.7%, +3.8%] | 2.2 pts | 2.0 | yes |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 10 | 13.4 µs | 13.7 µs | 1.04× [1.03, 1.05] | +3.6% | [+2.6%, +4.6%] | 1.4 pts | 1.1 | yes |
| multi | url | 1048576 | prefix | ordered | baseline | 10 | 10.1 µs | 10.4 µs | 1.05× [1.03, 1.06] | +4.4% | [+3.1%, +5.7%] | 1.8 pts | 0.2 | yes |
| multi | url | 1048576 | churn | ordered | baseline | 10 | 1063 | 1045 | 1.00× [0.97, 1.04] | +0.2% | [-3.2%, +3.6%] | 4.8 pts | 1.5 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 55.2 | 53.5 | 0.97× [0.96, 0.97] | -3.2% | [-3.7%, -2.7%] | 0.5 pts | 0.4 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3368 | 3316 | 0.98× [0.97, 1.00] | -1.7% | [-3.1%, -0.3%] | 1.3 pts | 0.5 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 93.6 | 89.6 | 0.96× [0.95, 0.96] | -4.3% | [-5.1%, -3.6%] | 0.7 pts | 0.7 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 81.1 | 76.0 | 0.93× [0.92, 0.94] | -7.2% | [-8.2%, -6.3%] | 0.9 pts | 0.6 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 6.41 ms | 6.16 ms | 0.96× [0.95, 0.97] | -4.1% | [-4.8%, -3.3%] | 0.8 pts | 0.8 | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 65.1 | 63.0 | 0.97× [0.96, 0.98] | -3.5% | [-4.4%, -2.6%] | 0.9 pts | 0.7 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3665 | 3588 | 0.99× [0.98, 1.00] | -1.5% | [-2.5%, -0.5%] | 1.0 pts | 0.8 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 115 | 110 | 0.96× [0.95, 0.97] | -3.9% | [-4.9%, -2.8%] | 1.0 pts | 0.8 | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 111 | 105 | 0.94× [0.93, 0.95] | -6.6% | [-8.0%, -5.2%] | 1.3 pts | 0.9 | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 32.50 ms | 30.97 ms | 0.95× [0.94, 0.97] | -4.8% | [-6.1%, -3.6%] | 1.2 pts | 0.6 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 297 | 299 | 1.00× [0.99, 1.02] | +0.1% | [-1.4%, +1.6%] | 2.1 pts | 1.8 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 8995 | 9015 | 1.00× [0.99, 1.02] | +0.1% | [-1.4%, +1.6%] | 2.0 pts | 1.5 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 613 | 629 | 1.02× [1.01, 1.03] | +1.9% | [+1.0%, +2.9%] | 1.3 pts | 0.7 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 518 | 497 | 0.96× [0.94, 0.99] | -3.7% | [-6.7%, -0.6%] | 4.3 pts | 0.9 | no |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 10 | 438 | 430 | 0.99× [0.98, 1.01] | -0.7% | [-2.1%, +0.6%] | 1.9 pts | 2.1 | yes |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 10 | 10.7 µs | 10.7 µs | 1.01× [0.99, 1.02] | +0.7% | [-0.6%, +2.1%] | 1.9 pts | 1.7 | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 10 | 1941 | 1964 | 1.01× [1.00, 1.02] | +1.0% | [-0.2%, +2.1%] | 1.6 pts | 1.2 | yes |
| multi | uuid | 1048576 | churn | ordered | baseline | 10 | 714 | 666 | 0.95× [0.91, 0.99] | -5.3% | [-9.9%, -0.6%] | 6.5 pts | 0.8 | no |
| unique | email | 4096 | valuesFor | ordered | baseline | 10 | 27.4 | 25.9 | 0.95× [0.94, 0.97] | -5.0% | [-6.9%, -3.2%] | 2.5 pts | 1.4 | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 10 | 1338 | 1288 | 0.96× [0.95, 0.97] | -4.2% | [-5.3%, -3.2%] | 1.5 pts | 1.0 | yes |
| unique | email | 4096 | prefix | ordered | baseline | 10 | 60.1 | 57.1 | 0.95× [0.94, 0.95] | -5.6% | [-6.3%, -4.9%] | 0.9 pts | 1.0 | yes |
| unique | email | 4096 | churn | ordered | baseline | 10 | 71.1 | 68.0 | 0.96× [0.95, 0.96] | -4.6% | [-5.0%, -4.2%] | 0.5 pts | 0.5 | yes |
| unique | email | 4096 | build | ordered | baseline | 10 | 886.8 µs | 831.2 µs | 0.94× [0.93, 0.95] | -6.7% | [-8.0%, -5.3%] | 1.9 pts | 1.1 | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 41.1 | 38.7 | 0.94× [0.94, 0.95] | -6.3% | [-6.8%, -5.7%] | 0.5 pts | 0.6 | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 1688 | 1642 | 0.97× [0.96, 0.98] | -2.8% | [-3.7%, -1.9%] | 0.8 pts | 0.9 | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 75.0 | 70.8 | 0.95× [0.94, 0.95] | -5.7% | [-6.4%, -5.1%] | 0.6 pts | 0.7 | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 93.5 | 88.5 | 0.95× [0.94, 0.97] | -4.8% | [-6.7%, -2.8%] | 1.8 pts | 0.6 | yes |
| unique | email | 16384 | build | ordered | baseline | 6 | 4.42 ms | 4.13 ms | 0.94× [0.93, 0.95] | -6.5% | [-7.4%, -5.6%] | 0.9 pts | 0.7 | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 6 | 222 | 222 | 1.00× [0.99, 1.01] | -0.3% | [-1.2%, +0.6%] | 0.8 pts | 0.5 | yes |
| unique | email | 262144 | valuesBetween | ordered | baseline | 6 | 4697 | 4679 | 1.00× [0.99, 1.01] | -0.2% | [-1.3%, +0.8%] | 1.0 pts | 0.6 | yes |
| unique | email | 262144 | prefix | ordered | baseline | 6 | 271 | 281 | 1.02× [1.00, 1.03] | +1.8% | [+0.5%, +3.1%] | 1.3 pts | 0.8 | yes |
| unique | email | 262144 | churn | ordered | baseline | 6 | 295 | 293 | 0.99× [0.98, 1.01] | -0.6% | [-2.1%, +0.9%] | 1.5 pts | 0.7 | yes |
| unique | email | 1048576 | valuesFor | ordered | baseline | 10 | 371 | 370 | 1.00× [0.99, 1.00] | -0.4% | [-1.0%, +0.2%] | 0.8 pts | 0.6 | yes |
| unique | email | 1048576 | valuesBetween | ordered | baseline | 10 | 6651 | 6626 | 1.00× [0.99, 1.00] | -0.3% | [-0.7%, +0.0%] | 0.5 pts | 0.5 | yes |
| unique | email | 1048576 | prefix | ordered | baseline | 10 | 479 | 473 | 1.00× [0.99, 1.01] | +0.1% | [-1.2%, +1.4%] | 1.8 pts | 1.2 | yes |
| unique | email | 1048576 | churn | ordered | baseline | 10 | 495 | 506 | 0.99× [0.96, 1.03] | -0.5% | [-4.2%, +3.2%] | 5.2 pts | 2.7 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 6 | 90.1 | 85.3 | 0.95× [0.94, 0.96] | -5.6% | [-6.9%, -4.3%] | 1.3 pts | 0.8 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 6 | 2283 | 2225 | 0.97× [0.96, 0.99] | -2.7% | [-4.4%, -0.9%] | 1.7 pts | 0.7 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 6 | 259 | 251 | 0.98× [0.96, 1.00] | -1.9% | [-3.7%, -0.1%] | 1.7 pts | 0.7 | yes |
| unique | path | 4096 | churn | ordered | baseline | 6 | 208 | 192 | 0.93× [0.92, 0.93] | -8.0% | [-8.8%, -7.1%] | 0.8 pts | 0.5 | yes |
| unique | path | 4096 | build | ordered | baseline | 6 | 2.34 ms | 2.10 ms | 0.90× [0.89, 0.92] | -10.6% | [-11.9%, -9.2%] | 1.3 pts | 0.8 | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 6 | 120 | 115 | 0.97× [0.96, 0.98] | -3.5% | [-4.5%, -2.5%] | 1.0 pts | 0.8 | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 6 | 2704 | 2689 | 1.00× [0.99, 1.00] | -0.2% | [-0.9%, +0.5%] | 0.7 pts | 0.6 | yes |
| unique | path | 16384 | prefix | ordered | baseline | 6 | 450 | 449 | 1.01× [0.99, 1.02] | +0.7% | [-1.0%, +2.4%] | 1.6 pts | 0.3 | yes |
| unique | path | 16384 | churn | ordered | baseline | 6 | 270 | 259 | 0.97× [0.95, 0.98] | -3.6% | [-5.1%, -2.0%] | 1.5 pts | 0.6 | yes |
| unique | path | 16384 | build | ordered | baseline | 6 | 11.51 ms | 10.63 ms | 0.92× [0.91, 0.94] | -8.2% | [-9.8%, -6.5%] | 1.6 pts | 0.9 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 453 | 470 | 1.04× [1.03, 1.05] | +3.9% | [+2.6%, +5.1%] | 1.7 pts | 0.9 | yes |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 8037 | 8630 | 1.10× [1.08, 1.11] | +8.8% | [+7.7%, +9.9%] | 1.6 pts | 1.0 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 5988 | 6333 | 1.06× [1.03, 1.09] | +5.4% | [+2.7%, +8.1%] | 3.8 pts | 0.4 | no |
| unique | path | 262144 | churn | ordered | baseline | 10 | 697 | 706 | 1.02× [1.00, 1.05] | +2.4% | [+0.4%, +4.4%] | 2.8 pts | 1.1 | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 8 | 37.2 | 36.5 | 0.98× [0.97, 0.99] | -2.0% | [-2.9%, -1.1%] | 1.1 pts | 0.7 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 8 | 1719 | 1694 | 0.99× [0.98, 0.99] | -1.5% | [-2.2%, -0.8%] | 0.8 pts | 0.5 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 8 | 3448 | 3351 | 0.97× [0.95, 0.99] | -3.1% | [-4.8%, -1.5%] | 2.0 pts | 0.7 | yes |
| unique | str | 4096 | churn | ordered | baseline | 8 | 93.0 | 88.3 | 0.95× [0.94, 0.95] | -5.8% | [-6.5%, -5.0%] | 0.9 pts | 0.7 | yes |
| unique | str | 4096 | build | ordered | baseline | 8 | 1.13 ms | 1.03 ms | 0.91× [0.90, 0.93] | -9.8% | [-11.6%, -8.0%] | 2.2 pts | 1.0 | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 6 | 49.3 | 48.4 | 0.98× [0.97, 0.99] | -2.1% | [-2.9%, -1.4%] | 0.7 pts | 0.8 | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 1746 | 1692 | 0.97× [0.96, 0.97] | -3.3% | [-3.9%, -2.6%] | 0.6 pts | 0.5 | yes |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 14.7 µs | 14.4 µs | 0.98× [0.96, 0.99] | -2.5% | [-3.8%, -1.1%] | 1.3 pts | 0.8 | yes |
| unique | str | 16384 | churn | ordered | baseline | 6 | 110 | 107 | 0.97× [0.96, 0.98] | -2.9% | [-4.1%, -1.7%] | 1.2 pts | 0.6 | yes |
| unique | str | 16384 | build | ordered | baseline | 6 | 5.27 ms | 4.99 ms | 0.94× [0.94, 0.95] | -5.9% | [-6.8%, -5.1%] | 0.8 pts | 0.7 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 214 | 231 | 1.08× [1.07, 1.09] | +7.2% | [+6.1%, +8.2%] | 1.5 pts | 0.4 | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 4216 | 4775 | 1.14× [1.12, 1.17] | +12.4% | [+10.7%, +14.2%] | 2.4 pts | 0.7 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 636.9 µs | 720.1 µs | 1.15× [1.11, 1.19] | +13.2% | [+10.2%, +16.2%] | 4.2 pts | 0.4 | no |
| unique | str | 262144 | churn | ordered | baseline | 10 | 336 | 336 | 1.01× [0.99, 1.03] | +1.1% | [-1.0%, +3.2%] | 2.9 pts | 0.9 | no |
| unique | str | 1048576 | valuesFor | ordered | baseline | 10 | 385 | 401 | 1.02× [1.02, 1.03] | +2.4% | [+1.6%, +3.3%] | 1.2 pts | 0.8 | yes |
| unique | str | 1048576 | valuesBetween | ordered | baseline | 10 | 6554 | 6934 | 1.05× [1.05, 1.06] | +5.2% | [+4.6%, +5.7%] | 0.8 pts | 0.5 | yes |
| unique | str | 1048576 | prefix | ordered | baseline | 10 | 3.96 ms | 4.23 ms | 1.07× [1.06, 1.07] | +6.1% | [+5.5%, +6.7%] | 0.9 pts | 1.2 | yes |
| unique | str | 1048576 | churn | ordered | baseline | 10 | 569 | 564 | 1.00× [0.98, 1.03] | +0.1% | [-2.5%, +2.7%] | 3.6 pts | 2.1 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 8 | 48.7 | 47.0 | 0.97× [0.96, 0.97] | -3.1% | [-3.7%, -2.6%] | 0.6 pts | 0.4 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 8 | 1968 | 1918 | 0.98× [0.96, 0.99] | -2.5% | [-4.3%, -0.8%] | 2.1 pts | 1.3 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 8 | 233 | 221 | 0.95× [0.93, 0.96] | -5.6% | [-7.4%, -3.7%] | 2.2 pts | 1.2 | yes |
| unique | street | 4096 | churn | ordered | baseline | 8 | 120 | 111 | 0.93× [0.92, 0.93] | -7.9% | [-8.5%, -7.3%] | 0.8 pts | 0.6 | yes |
| unique | street | 4096 | build | ordered | baseline | 8 | 1.44 ms | 1.30 ms | 0.91× [0.90, 0.92] | -10.4% | [-11.4%, -9.3%] | 1.3 pts | 0.8 | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 8 | 72.8 | 69.9 | 0.96× [0.96, 0.96] | -4.1% | [-4.5%, -3.7%] | 0.5 pts | 0.4 | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 8 | 2327 | 2261 | 0.97× [0.97, 0.98] | -3.0% | [-3.6%, -2.3%] | 0.8 pts | 0.7 | yes |
| unique | street | 16384 | prefix | ordered | baseline | 8 | 780 | 762 | 0.98× [0.96, 1.00] | -1.8% | [-3.8%, +0.1%] | 2.3 pts | 0.8 | yes |
| unique | street | 16384 | churn | ordered | baseline | 8 | 153 | 145 | 0.94× [0.93, 0.95] | -6.6% | [-7.9%, -5.3%] | 1.5 pts | 0.7 | yes |
| unique | street | 16384 | build | ordered | baseline | 8 | 7.13 ms | 6.52 ms | 0.91× [0.91, 0.92] | -9.6% | [-10.1%, -9.2%] | 0.5 pts | 0.6 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 10 | 16.1 | 14.7 | 0.92× [0.90, 0.95] | -8.5% | [-11.2%, -5.8%] | 3.8 pts | 1.6 | no |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 10 | 1120 | 1023 | 0.92× [0.91, 0.92] | -9.0% | [-9.8%, -8.2%] | 1.1 pts | 0.9 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 10 | 47.3 | 43.8 | 0.92× [0.92, 0.93] | -8.2% | [-8.6%, -7.7%] | 0.6 pts | 0.6 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 10 | 656.3 µs | 610.1 µs | 0.92× [0.92, 0.93] | -8.2% | [-9.0%, -7.4%] | 1.2 pts | 0.5 | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 10 | 25.8 | 23.4 | 0.91× [0.91, 0.91] | -10.0% | [-10.3%, -9.8%] | 0.3 pts | 0.6 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 10 | 1714 | 1642 | 0.95× [0.95, 0.96] | -4.8% | [-5.4%, -4.2%] | 0.8 pts | 0.6 | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 10 | 53.1 | 47.3 | 0.91× [0.88, 0.93] | -10.2% | [-13.0%, -7.5%] | 3.9 pts | 1.2 | no |
| unique | u64 | 16384 | build | ordered | baseline | 10 | 3.17 ms | 2.92 ms | 0.92× [0.92, 0.93] | -8.3% | [-8.8%, -7.8%] | 0.7 pts | 0.7 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 10 | 97.3 | 93.9 | 0.97× [0.95, 0.98] | -3.3% | [-4.8%, -1.8%] | 2.1 pts | 1.1 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 3328 | 3254 | 0.98× [0.97, 0.99] | -2.5% | [-3.5%, -1.4%] | 1.4 pts | 0.5 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 10 | 198 | 195 | 0.98× [0.96, 1.00] | -2.0% | [-4.1%, +0.1%] | 3.0 pts | 0.8 | no |
| unique | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 151 | 151 | 0.99× [0.98, 1.00] | -1.1% | [-1.9%, -0.4%] | 1.1 pts | 1.3 | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 3413 | 3352 | 0.98× [0.98, 0.99] | -1.8% | [-2.3%, -1.4%] | 0.7 pts | 0.8 | yes |
| unique | u64 | 1048576 | churn | ordered | baseline | 10 | 361 | 357 | 1.00× [0.98, 1.02] | -0.2% | [-2.2%, +1.7%] | 2.7 pts | 1.3 | yes |
| unique | url | 4096 | valuesFor | ordered | baseline | 8 | 69.1 | 67.3 | 0.97× [0.96, 0.98] | -3.0% | [-3.7%, -2.2%] | 0.9 pts | 0.7 | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 8 | 2094 | 2058 | 0.98× [0.96, 1.00] | -1.9% | [-3.7%, -0.2%] | 2.1 pts | 1.2 | yes |
| unique | url | 4096 | prefix | ordered | baseline | 8 | 161 | 156 | 0.97× [0.96, 0.99] | -2.6% | [-3.9%, -1.4%] | 1.5 pts | 1.0 | yes |
| unique | url | 4096 | churn | ordered | baseline | 8 | 162 | 150 | 0.93× [0.92, 0.94] | -7.7% | [-8.6%, -6.8%] | 1.1 pts | 0.7 | yes |
| unique | url | 4096 | build | ordered | baseline | 8 | 1.88 ms | 1.72 ms | 0.92× [0.91, 0.92] | -9.2% | [-10.0%, -8.4%] | 1.0 pts | 0.6 | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 95.6 | 93.5 | 0.98× [0.97, 0.98] | -2.4% | [-3.1%, -1.6%] | 0.7 pts | 0.6 | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 2512 | 2496 | 0.99× [0.98, 1.00] | -0.8% | [-2.0%, +0.3%] | 1.1 pts | 0.9 | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 223 | 221 | 0.98× [0.97, 1.00] | -1.8% | [-3.3%, -0.2%] | 1.5 pts | 0.8 | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 206 | 193 | 0.94× [0.93, 0.96] | -6.0% | [-7.7%, -4.3%] | 1.6 pts | 0.8 | yes |
| unique | url | 16384 | build | ordered | baseline | 6 | 9.35 ms | 8.53 ms | 0.92× [0.90, 0.93] | -9.1% | [-11.1%, -7.2%] | 1.8 pts | 0.9 | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 10 | 424 | 415 | 1.01× [0.97, 1.05] | +0.7% | [-3.0%, +4.3%] | 5.1 pts | 3.0 | no |
| unique | url | 262144 | valuesBetween | ordered | baseline | 10 | 7821 | 8255 | 1.05× [1.04, 1.06] | +4.7% | [+3.9%, +5.5%] | 1.1 pts | 0.7 | yes |
| unique | url | 262144 | prefix | ordered | baseline | 10 | 1466 | 1540 | 1.03× [1.01, 1.06] | +3.4% | [+1.2%, +5.5%] | 3.0 pts | 0.7 | no |
| unique | url | 262144 | churn | ordered | baseline | 10 | 633 | 634 | 0.99× [0.98, 1.01] | -0.5% | [-2.4%, +1.4%] | 2.6 pts | 1.8 | yes |
| unique | url | 1048576 | valuesFor | ordered | baseline | 10 | 704 | 745 | 1.03× [1.00, 1.05] | +2.6% | [+0.3%, +4.9%] | 3.2 pts | 2.2 | no |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 10 | 9871 | 10.3 µs | 1.04× [1.03, 1.05] | +3.6% | [+2.4%, +4.8%] | 1.6 pts | 1.4 | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 10 | 6332 | 6718 | 1.05× [1.03, 1.07] | +4.7% | [+2.8%, +6.7%] | 2.8 pts | 0.3 | yes |
| unique | url | 1048576 | churn | ordered | baseline | 10 | 1037 | 1013 | 0.99× [0.97, 1.01] | -1.3% | [-3.4%, +0.9%] | 3.0 pts | 1.4 | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 6 | 30.6 | 29.1 | 0.95× [0.93, 0.96] | -5.7% | [-7.1%, -4.3%] | 1.3 pts | 1.1 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1565 | 1526 | 0.98× [0.97, 0.99] | -2.3% | [-3.4%, -1.1%] | 1.1 pts | 0.8 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 72.0 | 68.5 | 0.95× [0.94, 0.96] | -5.2% | [-5.9%, -4.6%] | 0.6 pts | 0.7 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 6 | 82.6 | 78.0 | 0.95× [0.94, 0.96] | -5.6% | [-6.7%, -4.5%] | 1.1 pts | 1.0 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 6 | 994.4 µs | 915.7 µs | 0.93× [0.92, 0.94] | -7.4% | [-8.9%, -5.9%] | 1.4 pts | 0.8 | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 6 | 42.9 | 39.6 | 0.92× [0.92, 0.93] | -8.5% | [-9.0%, -8.0%] | 0.5 pts | 0.7 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 1678 | 1618 | 0.96× [0.95, 0.97] | -3.8% | [-5.0%, -2.6%] | 1.1 pts | 1.1 | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 87.5 | 83.5 | 0.96× [0.95, 0.96] | -4.5% | [-4.9%, -4.1%] | 0.3 pts | 0.4 | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 6 | 97.4 | 91.9 | 0.95× [0.93, 0.97] | -5.2% | [-7.1%, -3.3%] | 1.8 pts | 0.8 | yes |
| unique | uuid | 16384 | build | ordered | baseline | 6 | 4.73 ms | 4.47 ms | 0.94× [0.93, 0.95] | -6.2% | [-7.0%, -5.4%] | 0.8 pts | 0.6 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 241 | 237 | 0.98× [0.98, 0.99] | -1.5% | [-2.1%, -1.0%] | 0.8 pts | 0.4 | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 5190 | 5018 | 0.98× [0.97, 0.99] | -2.0% | [-2.9%, -1.2%] | 1.2 pts | 0.7 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 422 | 425 | 1.01× [1.00, 1.02] | +0.9% | [-0.2%, +2.1%] | 1.6 pts | 1.1 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 339 | 334 | 0.99× [0.97, 1.00] | -1.5% | [-3.1%, +0.2%] | 2.3 pts | 0.6 | yes |
| unique | uuid | 1048576 | valuesFor | ordered | baseline | 10 | 371 | 366 | 0.99× [0.98, 1.00] | -1.1% | [-1.9%, -0.4%] | 1.0 pts | 1.3 | yes |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 10 | 7054 | 6965 | 0.99× [0.99, 0.99] | -1.1% | [-1.4%, -0.9%] | 0.4 pts | 0.4 | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 10 | 1323 | 1318 | 1.00× [0.99, 1.01] | -0.3% | [-1.2%, +0.6%] | 1.3 pts | 0.9 | yes |
| unique | uuid | 1048576 | churn | ordered | baseline | 10 | 543 | 533 | 1.00× [0.96, 1.03] | -0.5% | [-4.0%, +3.0%] | 4.9 pts | 2.5 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.59% does not clear the 1.36% median noise floor of the processes
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.00%, 0.81%] includes zero
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.01% does not clear the 1.41% median noise floor of the processes
- multi email n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.68% does not clear the 1.16% median noise floor of the processes
- multi email n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.88%, 0.52%] includes zero
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.58% does not clear the 1.67% median noise floor of the processes
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.39%, 0.24%] includes zero
- multi email n=262144 prefix: ordered vs baseline: the pooled difference of -0.18% does not clear the 1.48% median noise floor of the processes
- multi email n=262144 prefix: ordered vs baseline: the pooled interval [-1.98%, 1.63%] includes zero
- multi email n=262144 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 prefix: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=262144 churn: ordered vs baseline: the pooled difference of -1.27% does not clear the 2.68% median noise floor of the processes
- multi email n=262144 churn: ordered vs baseline: the pooled interval [-7.19%, 4.64%] includes zero
- multi email n=262144 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 churn: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.39% does not clear the 1.05% median noise floor of the processes
- multi email n=1048576 valuesFor: ordered vs baseline: the pooled interval [-1.00%, 0.22%] includes zero
- multi email n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.78% does not clear the 1.87% median noise floor of the processes
- multi email n=1048576 prefix: ordered vs baseline: the pooled difference of 0.32% does not clear the 1.33% median noise floor of the processes
- multi email n=1048576 prefix: ordered vs baseline: the pooled interval [-0.03%, 0.68%] includes zero
- multi email n=1048576 churn: ordered vs baseline: the pooled difference of -2.73% does not clear the 4.37% median noise floor of the processes
- multi email n=1048576 churn: ordered vs baseline: the pooled interval [-7.53%, 2.07%] includes zero
- multi email n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of -0.78% does not clear the 3.93% median noise floor of the processes
- multi path n=4096 prefix: ordered vs baseline: the pooled interval [-1.88%, 0.32%] includes zero
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.41% does not clear the 1.76% median noise floor of the processes
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.41%, 1.23%] includes zero
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of -0.27% does not clear the 7.09% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-1.94%, 1.40%] includes zero
- multi path n=16384 churn: ordered vs baseline: the pooled difference of 1.25% does not clear the 2.67% median noise floor of the processes
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of 3.98% does not clear the 10.21% median noise floor of the processes
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.20% does not clear the 1.59% median noise floor of the processes
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.75%, 1.35%] includes zero
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of -1.31% does not clear the 2.47% median noise floor of the processes
- multi str n=4096 prefix: ordered vs baseline: the pooled interval [-2.82%, 0.20%] includes zero
- multi str n=16384 valuesFor: ordered vs baseline: the pooled difference of -1.23% does not clear the 1.29% median noise floor of the processes
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.12% does not clear the 1.66% median noise floor of the processes
- multi str n=16384 prefix: ordered vs baseline: the pooled difference of -0.97% does not clear the 2.36% median noise floor of the processes
- multi str n=16384 prefix: ordered vs baseline: the pooled interval [-2.15%, 0.20%] includes zero
- multi str n=262144 churn: ordered vs baseline: the pooled interval [-5.57%, 0.45%] includes zero
- multi str n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=1048576 churn: ordered vs baseline: the pooled difference of 1.61% does not clear the 2.52% median noise floor of the processes
- multi str n=1048576 churn: ordered vs baseline: the pooled interval [-3.71%, 6.92%] includes zero
- multi str n=1048576 churn: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 churn: ordered vs baseline: 5 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.62% does not clear the 2.36% median noise floor of the processes
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled difference of -2.08% does not clear the 2.46% median noise floor of the processes
- multi street n=16384 prefix: ordered vs baseline: the pooled difference of -3.31% does not clear the 5.62% median noise floor of the processes
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.32% does not clear the 1.87% median noise floor of the processes
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.99% does not clear the 1.77% median noise floor of the processes
- multi u64 n=262144 churn: ordered vs baseline: the pooled difference of 0.56% does not clear the 2.25% median noise floor of the processes
- multi u64 n=262144 churn: ordered vs baseline: the pooled interval [-4.65%, 5.77%] includes zero
- multi u64 n=262144 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -1.50% does not clear the 1.63% median noise floor of the processes
- multi u64 n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=16384 valuesFor: ordered vs baseline: the pooled difference of -1.04% does not clear the 1.71% median noise floor of the processes
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.96% does not clear the 1.22% median noise floor of the processes
- multi url n=16384 prefix: ordered vs baseline: the pooled difference of -1.23% does not clear the 1.75% median noise floor of the processes
- multi url n=16384 churn: ordered vs baseline: the pooled difference of -1.39% does not clear the 1.85% median noise floor of the processes
- multi url n=262144 valuesFor: ordered vs baseline: the pooled difference of 1.28% does not clear the 1.71% median noise floor of the processes
- multi url n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.86%, 4.43%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs baseline: 5 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of 2.52% does not clear the 6.37% median noise floor of the processes
- multi url n=262144 churn: ordered vs baseline: the pooled difference of 1.43% does not clear the 2.21% median noise floor of the processes
- multi url n=262144 churn: ordered vs baseline: the pooled interval [-3.72%, 6.58%] includes zero
- multi url n=262144 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 churn: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 prefix: ordered vs baseline: the pooled difference of 4.39% does not clear the 8.71% median noise floor of the processes
- multi url n=1048576 churn: ordered vs baseline: the pooled difference of 0.19% does not clear the 3.03% median noise floor of the processes
- multi url n=1048576 churn: ordered vs baseline: the pooled interval [-3.25%, 3.63%] includes zero
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.73% does not clear the 2.35% median noise floor of the processes
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of -1.49% does not clear the 1.54% median noise floor of the processes
- multi uuid n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.14% does not clear the 1.28% median noise floor of the processes
- multi uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.37%, 1.64%] includes zero
- multi uuid n=262144 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.09% does not clear the 1.90% median noise floor of the processes
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.37%, 1.55%] includes zero
- multi uuid n=262144 prefix: ordered vs baseline: the pooled difference of 1.94% does not clear the 2.22% median noise floor of the processes
- multi uuid n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.72% does not clear the 1.41% median noise floor of the processes
- multi uuid n=1048576 valuesFor: ordered vs baseline: the pooled interval [-2.08%, 0.64%] includes zero
- multi uuid n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.72% does not clear the 1.76% median noise floor of the processes
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.62%, 2.07%] includes zero
- multi uuid n=1048576 prefix: ordered vs baseline: the pooled difference of 0.99% does not clear the 1.27% median noise floor of the processes
- multi uuid n=1048576 prefix: ordered vs baseline: the pooled interval [-0.16%, 2.14%] includes zero
- unique email n=262144 valuesFor: ordered vs baseline: the pooled difference of -0.32% does not clear the 1.13% median noise floor of the processes
- unique email n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.19%, 0.55%] includes zero
- unique email n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.22% does not clear the 2.24% median noise floor of the processes
- unique email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.27%, 0.83%] includes zero
- unique email n=262144 churn: ordered vs baseline: the pooled difference of -0.61% does not clear the 2.24% median noise floor of the processes
- unique email n=262144 churn: ordered vs baseline: the pooled interval [-2.14%, 0.93%] includes zero
- unique email n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.40% does not clear the 1.45% median noise floor of the processes
- unique email n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.98%, 0.18%] includes zero
- unique email n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.32% does not clear the 0.97% median noise floor of the processes
- unique email n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.67%, 0.04%] includes zero
- unique email n=1048576 prefix: ordered vs baseline: the pooled difference of 0.11% does not clear the 1.18% median noise floor of the processes
- unique email n=1048576 prefix: ordered vs baseline: the pooled interval [-1.21%, 1.43%] includes zero
- unique email n=1048576 churn: ordered vs baseline: the pooled difference of -0.50% does not clear the 1.30% median noise floor of the processes
- unique email n=1048576 churn: ordered vs baseline: the pooled interval [-4.24%, 3.23%] includes zero
- unique email n=1048576 churn: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=1048576 churn: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=4096 prefix: ordered vs baseline: the pooled difference of -1.92% does not clear the 2.44% median noise floor of the processes
- unique path n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.22% does not clear the 1.36% median noise floor of the processes
- unique path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.92%, 0.48%] includes zero
- unique path n=16384 prefix: ordered vs baseline: the pooled difference of 0.70% does not clear the 7.30% median noise floor of the processes
- unique path n=16384 prefix: ordered vs baseline: the pooled interval [-1.01%, 2.42%] includes zero
- unique path n=262144 prefix: ordered vs baseline: the pooled difference of 5.38% does not clear the 11.76% median noise floor of the processes
- unique str n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.49% does not clear the 1.59% median noise floor of the processes
- unique str n=4096 prefix: ordered vs baseline: the pooled difference of -3.14% does not clear the 3.84% median noise floor of the processes
- unique str n=262144 churn: ordered vs baseline: the pooled difference of 1.12% does not clear the 3.68% median noise floor of the processes
- unique str n=262144 churn: ordered vs baseline: the pooled interval [-0.97%, 3.21%] includes zero
- unique str n=1048576 churn: ordered vs baseline: the pooled difference of 0.13% does not clear the 1.93% median noise floor of the processes
- unique str n=1048576 churn: ordered vs baseline: the pooled interval [-2.46%, 2.71%] includes zero
- unique str n=1048576 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique street n=16384 prefix: ordered vs baseline: the pooled difference of -1.83% does not clear the 4.95% median noise floor of the processes
- unique street n=16384 prefix: ordered vs baseline: the pooled interval [-3.79%, 0.14%] includes zero
- unique u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of -2.47% does not clear the 3.34% median noise floor of the processes
- unique u64 n=262144 churn: ordered vs baseline: the pooled difference of -2.02% does not clear the 2.91% median noise floor of the processes
- unique u64 n=262144 churn: ordered vs baseline: the pooled interval [-4.13%, 0.10%] includes zero
- unique u64 n=1048576 valuesFor: ordered vs baseline: the pooled difference of -1.13% does not clear the 1.33% median noise floor of the processes
- unique u64 n=1048576 churn: ordered vs baseline: the pooled difference of -0.22% does not clear the 1.80% median noise floor of the processes
- unique u64 n=1048576 churn: ordered vs baseline: the pooled interval [-2.15%, 1.71%] includes zero
- unique u64 n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.82% does not clear the 1.74% median noise floor of the processes
- unique url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.95%, 0.32%] includes zero
- unique url n=16384 prefix: ordered vs baseline: the pooled difference of -1.76% does not clear the 1.80% median noise floor of the processes
- unique url n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.69% does not clear the 1.32% median noise floor of the processes
- unique url n=262144 valuesFor: ordered vs baseline: the pooled interval [-2.96%, 4.34%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 prefix: ordered vs baseline: the pooled difference of 3.36% does not clear the 6.21% median noise floor of the processes
- unique url n=262144 churn: ordered vs baseline: the pooled difference of -0.51% does not clear the 1.81% median noise floor of the processes
- unique url n=262144 churn: ordered vs baseline: the pooled interval [-2.39%, 1.38%] includes zero
- unique url n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=1048576 prefix: ordered vs baseline: the pooled difference of 4.74% does not clear the 8.98% median noise floor of the processes
- unique url n=1048576 churn: ordered vs baseline: the pooled difference of -1.29% does not clear the 2.34% median noise floor of the processes
- unique url n=1048576 churn: ordered vs baseline: the pooled interval [-3.45%, 0.87%] includes zero
- unique url n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=262144 prefix: ordered vs baseline: the pooled difference of 0.94% does not clear the 1.68% median noise floor of the processes
- unique uuid n=262144 prefix: ordered vs baseline: the pooled interval [-0.22%, 2.09%] includes zero
- unique uuid n=262144 churn: ordered vs baseline: the pooled difference of -1.48% does not clear the 2.37% median noise floor of the processes
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-3.13%, 0.18%] includes zero
- unique uuid n=1048576 valuesFor: ordered vs baseline: the pooled difference of -1.15% does not clear the 1.45% median noise floor of the processes
- unique uuid n=1048576 prefix: ordered vs baseline: the pooled difference of -0.31% does not clear the 1.27% median noise floor of the processes
- unique uuid n=1048576 prefix: ordered vs baseline: the pooled interval [-1.24%, 0.62%] includes zero
- unique uuid n=1048576 churn: ordered vs baseline: the pooled difference of -0.49% does not clear the 1.74% median noise floor of the processes
- unique uuid n=1048576 churn: ordered vs baseline: the pooled interval [-3.98%, 3.01%] includes zero
- unique uuid n=1048576 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 churn: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
