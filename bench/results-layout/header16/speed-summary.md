| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 6 | 49.2 | 49.3 | 1.00× [1.00, 1.01] | +0.4% | [-0.3%, +1.1%] | 0.7 pts | 0.6 | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 6 | 2927 | 2938 | 1.01× [1.00, 1.02] | +0.9% | [-0.5%, +2.2%] | 1.3 pts | 0.7 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 6 | 77.0 | 76.0 | 0.98× [0.97, 0.99] | -1.8% | [-2.6%, -1.1%] | 0.7 pts | 1.0 | yes |
| multi | email | 4096 | churn | ordered | baseline | 6 | 62.3 | 61.6 | 0.99× [0.98, 1.00] | -1.2% | [-2.0%, -0.4%] | 0.8 pts | 0.5 | yes |
| multi | email | 4096 | build | ordered | baseline | 6 | 5.32 ms | 5.29 ms | 0.99× [0.98, 1.01] | -0.6% | [-2.2%, +1.0%] | 1.5 pts | 1.0 | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 60.1 | 60.2 | 1.01× [1.00, 1.03] | +1.2% | [-0.3%, +2.6%] | 1.3 pts | 1.0 | yes |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3468 | 3507 | 1.01× [1.00, 1.02] | +0.9% | [+0.1%, +1.7%] | 0.8 pts | 0.6 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 94.5 | 93.5 | 0.99× [0.98, 1.00] | -0.9% | [-1.8%, -0.0%] | 0.8 pts | 0.8 | yes |
| multi | email | 16384 | churn | ordered | baseline | 6 | 99.5 | 99.9 | 1.00× [1.00, 1.01] | +0.2% | [-0.2%, +0.5%] | 0.3 pts | 0.2 | yes |
| multi | email | 16384 | build | ordered | baseline | 6 | 26.99 ms | 27.00 ms | 1.00× [0.99, 1.01] | -0.2% | [-1.1%, +0.7%] | 0.9 pts | 1.1 | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 10 | 269 | 282 | 1.04× [1.03, 1.05] | +3.7% | [+2.6%, +4.9%] | 1.6 pts | 1.3 | yes |
| multi | email | 262144 | valuesBetween | ordered | baseline | 10 | 9564 | 9671 | 1.01× [1.01, 1.02] | +1.0% | [+0.5%, +1.5%] | 0.7 pts | 0.7 | yes |
| multi | email | 262144 | prefix | ordered | baseline | 10 | 348 | 350 | 1.01× [1.00, 1.02] | +1.0% | [-0.0%, +2.0%] | 1.4 pts | 1.0 | yes |
| multi | email | 262144 | churn | ordered | baseline | 10 | 387 | 382 | 0.99× [0.96, 1.03] | -0.5% | [-3.8%, +2.8%] | 4.6 pts | 1.8 | no |
| multi | email | 1048576 | valuesFor | ordered | baseline | 10 | 378 | 387 | 1.03× [1.01, 1.04] | +2.5% | [+1.3%, +3.7%] | 1.7 pts | 2.1 | yes |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 10 | 11.3 µs | 11.5 µs | 1.01× [1.00, 1.02] | +1.1% | [+0.3%, +1.9%] | 1.1 pts | 1.0 | yes |
| multi | email | 1048576 | prefix | ordered | baseline | 10 | 653 | 673 | 1.03× [1.02, 1.04] | +3.2% | [+2.2%, +4.2%] | 1.4 pts | 1.3 | yes |
| multi | email | 1048576 | churn | ordered | baseline | 10 | 562 | 565 | 1.01× [0.98, 1.04] | +1.0% | [-1.9%, +3.8%] | 4.0 pts | 1.0 | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 107 | 105 | 0.98× [0.97, 0.98] | -2.6% | [-3.0%, -2.1%] | 0.4 pts | 0.4 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4125 | 4134 | 1.01× [1.00, 1.02] | +0.6% | [-0.4%, +1.5%] | 0.9 pts | 0.5 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 333 | 329 | 0.99× [0.98, 1.00] | -1.1% | [-2.4%, +0.3%] | 1.3 pts | 0.4 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 158 | 159 | 1.00× [1.00, 1.01] | +0.5% | [-0.5%, +1.5%] | 0.9 pts | 0.4 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 11.44 ms | 11.46 ms | 1.00× [1.00, 1.01] | +0.3% | [-0.2%, +0.8%] | 0.5 pts | 0.5 | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 6 | 137 | 136 | 0.98× [0.97, 1.00] | -1.7% | [-3.1%, -0.2%] | 1.4 pts | 1.0 | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 6 | 4691 | 4713 | 1.01× [1.00, 1.02] | +0.7% | [-0.2%, +1.5%] | 0.8 pts | 0.7 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 6 | 614 | 617 | 1.00× [0.99, 1.02] | +0.5% | [-1.3%, +2.2%] | 1.7 pts | 0.3 | yes |
| multi | path | 16384 | churn | ordered | baseline | 6 | 248 | 248 | 1.01× [1.00, 1.02] | +0.6% | [-0.2%, +1.5%] | 0.8 pts | 0.7 | yes |
| multi | path | 16384 | build | ordered | baseline | 6 | 65.46 ms | 66.58 ms | 1.01× [1.01, 1.02] | +1.4% | [+0.6%, +2.2%] | 0.7 pts | 0.8 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 537 | 538 | 1.00× [0.99, 1.02] | +0.3% | [-1.4%, +2.1%] | 2.5 pts | 1.8 | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 13.3 µs | 13.2 µs | 1.00× [0.98, 1.01] | -0.2% | [-1.7%, +1.3%] | 2.2 pts | 1.6 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 12.4 µs | 12.6 µs | 1.02× [0.99, 1.05] | +2.1% | [-0.9%, +5.1%] | 4.2 pts | 0.3 | no |
| multi | path | 262144 | churn | ordered | baseline | 10 | 843 | 810 | 0.94× [0.90, 0.98] | -6.5% | [-10.8%, -2.2%] | 6.0 pts | 3.0 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 6 | 60.4 | 60.9 | 1.00× [1.00, 1.01] | +0.5% | [-0.3%, +1.2%] | 0.7 pts | 0.8 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 6 | 3478 | 3482 | 1.00× [0.99, 1.02] | +0.4% | [-0.9%, +1.7%] | 1.2 pts | 0.5 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 6 | 7822 | 7858 | 1.00× [0.99, 1.01] | -0.1% | [-1.3%, +1.2%] | 1.2 pts | 0.4 | yes |
| multi | str | 4096 | churn | ordered | baseline | 6 | 77.7 | 78.2 | 1.01× [1.00, 1.02] | +0.6% | [-0.2%, +1.5%] | 0.8 pts | 0.6 | yes |
| multi | str | 4096 | build | ordered | baseline | 6 | 6.33 ms | 6.38 ms | 1.00× [0.99, 1.01] | +0.2% | [-0.6%, +1.1%] | 0.8 pts | 0.6 | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 71.2 | 73.6 | 1.03× [1.03, 1.04] | +3.1% | [+2.5%, +3.8%] | 0.6 pts | 0.6 | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3636 | 3679 | 1.01× [1.00, 1.01] | +0.6% | [-0.2%, +1.4%] | 0.8 pts | 0.9 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 34.2 µs | 34.4 µs | 1.00× [0.99, 1.02] | +0.3% | [-1.4%, +2.1%] | 1.7 pts | 1.1 | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 116 | 116 | 1.00× [0.99, 1.01] | +0.3% | [-0.7%, +1.3%] | 1.0 pts | 0.4 | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 32.05 ms | 32.19 ms | 1.01× [1.00, 1.01] | +0.5% | [-0.4%, +1.4%] | 0.9 pts | 1.2 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 310 | 321 | 1.04× [1.02, 1.05] | +3.6% | [+2.4%, +4.8%] | 1.7 pts | 1.7 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 10.2 µs | 10.3 µs | 1.02× [1.00, 1.03] | +1.7% | [+0.3%, +3.0%] | 1.9 pts | 1.5 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.67 ms | 1.70 ms | 1.01× [1.00, 1.03] | +1.4% | [+0.2%, +2.5%] | 1.6 pts | 1.4 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 414 | 423 | 1.02× [0.99, 1.04] | +1.6% | [-1.0%, +4.2%] | 3.6 pts | 1.8 | no |
| multi | str | 1048576 | valuesFor | ordered | baseline | 10 | 424 | 435 | 1.03× [1.02, 1.05] | +3.1% | [+1.8%, +4.4%] | 1.9 pts | 1.9 | yes |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 10 | 11.7 µs | 11.7 µs | 1.00× [1.00, 1.01] | +0.4% | [-0.5%, +1.2%] | 1.2 pts | 1.1 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 10 | 7.16 ms | 7.22 ms | 1.01× [1.00, 1.01] | +0.6% | [-0.0%, +1.2%] | 0.9 pts | 1.8 | yes |
| multi | str | 1048576 | churn | ordered | baseline | 10 | 600 | 581 | 0.98× [0.95, 1.02] | -1.6% | [-4.8%, +1.7%] | 4.6 pts | 2.1 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 6 | 54.5 | 54.3 | 1.00× [0.99, 1.01] | +0.1% | [-1.0%, +1.3%] | 1.1 pts | 0.8 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 6 | 2380 | 2373 | 0.99× [0.98, 1.00] | -0.7% | [-1.9%, +0.5%] | 1.2 pts | 0.7 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 6 | 255 | 257 | 1.00× [0.99, 1.01] | +0.3% | [-0.7%, +1.2%] | 0.9 pts | 0.4 | yes |
| multi | street | 4096 | churn | ordered | baseline | 6 | 89.8 | 89.2 | 0.99× [0.99, 1.00] | -0.6% | [-1.5%, +0.2%] | 0.8 pts | 0.6 | yes |
| multi | street | 4096 | build | ordered | baseline | 6 | 2.78 ms | 2.75 ms | 0.99× [0.98, 1.00] | -0.6% | [-1.5%, +0.3%] | 0.9 pts | 1.0 | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 75.4 | 76.1 | 1.01× [1.00, 1.02] | +0.7% | [-0.2%, +1.6%] | 0.9 pts | 0.6 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2804 | 2820 | 1.00× [1.00, 1.01] | +0.4% | [+0.1%, +0.6%] | 0.2 pts | 0.2 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 905 | 917 | 1.01× [0.99, 1.02] | +0.6% | [-0.8%, +2.1%] | 1.4 pts | 0.5 | yes |
| multi | street | 16384 | churn | ordered | baseline | 6 | 128 | 128 | 1.00× [0.99, 1.02] | +0.2% | [-1.1%, +1.5%] | 1.2 pts | 0.7 | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 15.31 ms | 15.23 ms | 0.99× [0.99, 1.00] | -0.7% | [-1.3%, -0.1%] | 0.6 pts | 0.8 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 8 | 37.0 | 37.0 | 1.00× [0.99, 1.01] | +0.0% | [-0.7%, +0.8%] | 0.9 pts | 0.5 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2566 | 2686 | 1.05× [1.04, 1.05] | +4.3% | [+3.5%, +5.1%] | 1.0 pts | 0.4 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 8 | 44.7 | 44.3 | 0.99× [0.98, 1.01] | -0.8% | [-2.5%, +0.8%] | 2.0 pts | 1.4 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 8 | 3.96 ms | 3.96 ms | 1.00× [0.99, 1.01] | -0.2% | [-1.1%, +0.8%] | 1.2 pts | 0.7 | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 44.3 | 44.4 | 1.01× [1.00, 1.02] | +0.7% | [-0.1%, +1.5%] | 0.8 pts | 1.0 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3432 | 3588 | 1.05× [1.04, 1.06] | +4.6% | [+3.9%, +5.4%] | 0.7 pts | 0.6 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 57.6 | 57.4 | 0.99× [0.98, 1.01] | -0.7% | [-2.0%, +0.5%] | 1.2 pts | 0.6 | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 19.49 ms | 19.16 ms | 0.98× [0.97, 0.99] | -1.9% | [-2.6%, -1.3%] | 0.7 pts | 1.0 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 147 | 150 | 1.02× [1.01, 1.03] | +2.1% | [+1.2%, +3.0%] | 1.3 pts | 1.3 | yes |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 7938 | 8247 | 1.04× [1.03, 1.05] | +3.7% | [+3.1%, +4.3%] | 0.8 pts | 0.7 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 247 | 250 | 1.01× [0.97, 1.06] | +1.3% | [-2.7%, +5.3%] | 5.6 pts | 1.9 | no |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 197 | 201 | 1.01× [1.00, 1.02] | +0.7% | [-0.3%, +1.7%] | 1.4 pts | 1.4 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 7936 | 8110 | 1.02× [1.01, 1.03] | +1.9% | [+1.3%, +2.5%] | 0.9 pts | 0.9 | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 10 | 368 | 367 | 0.96× [0.91, 1.01] | -4.7% | [-10.0%, +0.6%] | 7.4 pts | 0.9 | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 6 | 81.8 | 81.4 | 1.00× [0.99, 1.00] | -0.2% | [-0.8%, +0.3%] | 0.6 pts | 0.5 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 6 | 3677 | 3684 | 0.99× [0.98, 1.01] | -0.7% | [-2.1%, +0.8%] | 1.4 pts | 0.8 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 6 | 405 | 401 | 0.99× [0.97, 1.00] | -1.2% | [-2.7%, +0.3%] | 1.4 pts | 0.8 | yes |
| multi | url | 4096 | churn | ordered | baseline | 6 | 117 | 118 | 1.00× [1.00, 1.01] | +0.2% | [-0.4%, +0.8%] | 0.6 pts | 0.5 | yes |
| multi | url | 4096 | build | ordered | baseline | 6 | 9.07 ms | 9.14 ms | 1.01× [1.00, 1.02] | +0.6% | [-0.4%, +1.5%] | 0.9 pts | 1.1 | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 96.0 | 96.1 | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.6%] | 0.8 pts | 0.9 | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 3934 | 3937 | 1.00× [0.99, 1.01] | +0.1% | [-0.6%, +0.8%] | 0.6 pts | 0.6 | yes |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 1305 | 1297 | 1.00× [0.98, 1.02] | +0.2% | [-1.6%, +2.0%] | 1.7 pts | 0.8 | yes |
| multi | url | 16384 | churn | ordered | baseline | 6 | 173 | 173 | 1.00× [0.99, 1.00] | -0.3% | [-1.0%, +0.5%] | 0.7 pts | 0.5 | yes |
| multi | url | 16384 | build | ordered | baseline | 6 | 47.49 ms | 47.63 ms | 1.00× [1.00, 1.01] | +0.1% | [-0.3%, +0.6%] | 0.4 pts | 0.4 | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 10 | 430 | 435 | 1.01× [1.00, 1.03] | +1.3% | [+0.1%, +2.5%] | 1.7 pts | 1.4 | yes |
| multi | url | 262144 | valuesBetween | ordered | baseline | 10 | 11.1 µs | 11.1 µs | 1.01× [1.00, 1.02] | +0.6% | [-0.5%, +1.6%] | 1.5 pts | 1.3 | yes |
| multi | url | 262144 | prefix | ordered | baseline | 10 | 52.4 µs | 51.5 µs | 0.99× [0.96, 1.02] | -0.7% | [-3.7%, +2.3%] | 4.2 pts | 0.6 | no |
| multi | url | 262144 | churn | ordered | baseline | 10 | 561 | 568 | 1.01× [0.99, 1.03] | +1.0% | [-1.4%, +3.3%] | 3.2 pts | 1.6 | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 10 | 550 | 562 | 1.01× [1.00, 1.03] | +1.4% | [+0.1%, +2.7%] | 1.8 pts | 1.7 | yes |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 10 | 12.5 µs | 12.4 µs | 1.00× [0.99, 1.00] | -0.4% | [-1.0%, +0.2%] | 0.8 pts | 0.7 | yes |
| multi | url | 1048576 | prefix | ordered | baseline | 10 | 233.9 µs | 234.7 µs | 0.99× [0.98, 1.01] | -0.6% | [-2.1%, +0.9%] | 2.1 pts | 0.1 | yes |
| multi | url | 1048576 | churn | ordered | baseline | 10 | 725 | 750 | 1.02× [0.98, 1.05] | +1.6% | [-1.8%, +5.0%] | 4.7 pts | 1.7 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 53.5 | 54.6 | 1.02× [1.01, 1.04] | +2.2% | [+0.9%, +3.5%] | 1.2 pts | 0.9 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3247 | 3308 | 1.01× [1.00, 1.02] | +1.3% | [+0.3%, +2.4%] | 1.0 pts | 0.5 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 90.6 | 88.3 | 0.98× [0.97, 0.99] | -2.2% | [-3.5%, -0.8%] | 1.3 pts | 1.6 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 71.7 | 71.7 | 0.99× [0.98, 1.00] | -1.0% | [-2.2%, +0.2%] | 1.2 pts | 0.8 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 5.92 ms | 5.94 ms | 1.00× [0.99, 1.01] | -0.1% | [-1.2%, +1.0%] | 1.1 pts | 0.8 | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 62.5 | 64.9 | 1.04× [1.03, 1.04] | +3.5% | [+3.0%, +4.0%] | 0.5 pts | 0.5 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3566 | 3599 | 1.01× [1.00, 1.02] | +1.3% | [+0.3%, +2.3%] | 0.9 pts | 0.8 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 112 | 109 | 0.98× [0.97, 0.99] | -2.1% | [-2.8%, -1.3%] | 0.7 pts | 0.6 | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 107 | 107 | 1.00× [0.99, 1.01] | +0.3% | [-0.6%, +1.3%] | 0.9 pts | 0.5 | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 29.72 ms | 29.70 ms | 1.00× [1.00, 1.01] | +0.1% | [-0.4%, +0.7%] | 0.5 pts | 0.7 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 275 | 282 | 1.03× [1.02, 1.03] | +2.4% | [+1.7%, +3.2%] | 1.0 pts | 1.2 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 10.4 µs | 10.5 µs | 1.02× [1.01, 1.03] | +1.5% | [+0.6%, +2.5%] | 1.3 pts | 1.0 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 670 | 690 | 1.03× [1.02, 1.05] | +3.3% | [+2.2%, +4.5%] | 1.6 pts | 1.4 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 407 | 411 | 0.99× [0.96, 1.03] | -0.7% | [-3.9%, +2.5%] | 4.4 pts | 1.5 | no |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 10 | 403 | 407 | 1.02× [1.01, 1.04] | +2.1% | [+0.5%, +3.6%] | 2.2 pts | 2.6 | yes |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 10 | 11.9 µs | 12.1 µs | 1.01× [1.00, 1.02] | +0.7% | [-0.5%, +1.9%] | 1.7 pts | 2.0 | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 10 | 2189 | 2214 | 1.01× [1.00, 1.03] | +1.2% | [-0.1%, +2.5%] | 1.8 pts | 1.9 | yes |
| multi | uuid | 1048576 | churn | ordered | baseline | 10 | 613 | 598 | 0.98× [0.95, 1.02] | -1.6% | [-5.2%, +1.9%] | 5.0 pts | 1.0 | no |
| unique | email | 4096 | valuesFor | ordered | baseline | 8 | 26.0 | 25.8 | 0.99× [0.97, 1.00] | -1.5% | [-3.3%, +0.4%] | 2.2 pts | 1.2 | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 8 | 1230 | 1286 | 1.05× [1.04, 1.06] | +4.5% | [+3.8%, +5.3%] | 0.9 pts | 0.8 | yes |
| unique | email | 4096 | prefix | ordered | baseline | 8 | 57.4 | 55.4 | 0.97× [0.96, 0.97] | -3.5% | [-3.8%, -3.2%] | 0.4 pts | 0.5 | yes |
| unique | email | 4096 | churn | ordered | baseline | 8 | 68.2 | 67.0 | 0.99× [0.98, 0.99] | -1.4% | [-1.9%, -0.9%] | 0.6 pts | 0.5 | yes |
| unique | email | 4096 | build | ordered | baseline | 8 | 834.8 µs | 828.2 µs | 0.99× [0.97, 1.01] | -0.9% | [-2.6%, +0.8%] | 2.0 pts | 0.7 | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 10 | 37.6 | 37.8 | 1.01× [1.00, 1.01] | +0.9% | [+0.4%, +1.4%] | 0.7 pts | 1.1 | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 10 | 1573 | 1608 | 1.02× [1.02, 1.02] | +1.9% | [+1.6%, +2.3%] | 0.5 pts | 0.5 | yes |
| unique | email | 16384 | prefix | ordered | baseline | 10 | 70.8 | 70.0 | 0.99× [0.99, 0.99] | -1.1% | [-1.5%, -0.6%] | 0.6 pts | 0.6 | yes |
| unique | email | 16384 | churn | ordered | baseline | 10 | 87.6 | 86.6 | 1.00× [0.98, 1.02] | -0.3% | [-2.1%, +1.5%] | 2.5 pts | 0.8 | yes |
| unique | email | 16384 | build | ordered | baseline | 10 | 4.12 ms | 4.06 ms | 0.99× [0.98, 0.99] | -1.3% | [-2.0%, -0.7%] | 0.9 pts | 0.7 | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 6 | 179 | 189 | 1.06× [1.05, 1.07] | +5.7% | [+4.5%, +6.9%] | 1.1 pts | 1.2 | yes |
| unique | email | 262144 | valuesBetween | ordered | baseline | 6 | 5089 | 5228 | 1.03× [1.02, 1.04] | +3.1% | [+2.2%, +4.1%] | 0.9 pts | 0.8 | yes |
| unique | email | 262144 | prefix | ordered | baseline | 6 | 273 | 282 | 1.02× [1.00, 1.04] | +1.9% | [+0.2%, +3.6%] | 1.6 pts | 1.0 | yes |
| unique | email | 262144 | churn | ordered | baseline | 6 | 286 | 286 | 1.00× [0.99, 1.02] | +0.2% | [-1.3%, +1.7%] | 1.4 pts | 0.7 | yes |
| unique | email | 1048576 | valuesFor | ordered | baseline | 10 | 304 | 313 | 1.02× [1.01, 1.03] | +2.0% | [+1.3%, +2.7%] | 1.0 pts | 1.3 | yes |
| unique | email | 1048576 | valuesBetween | ordered | baseline | 10 | 6672 | 6820 | 1.02× [1.01, 1.03] | +1.9% | [+1.2%, +2.5%] | 0.9 pts | 1.1 | yes |
| unique | email | 1048576 | prefix | ordered | baseline | 10 | 435 | 455 | 1.05× [1.04, 1.06] | +4.8% | [+3.5%, +6.1%] | 1.8 pts | 1.8 | yes |
| unique | email | 1048576 | churn | ordered | baseline | 10 | 492 | 492 | 1.00× [0.98, 1.03] | +0.4% | [-1.8%, +2.6%] | 3.1 pts | 1.7 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 6 | 86.1 | 83.9 | 0.97× [0.96, 0.99] | -2.7% | [-4.0%, -1.3%] | 1.3 pts | 1.2 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 6 | 2171 | 2184 | 1.00× [0.99, 1.02] | +0.4% | [-0.7%, +1.5%] | 1.0 pts | 0.6 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 6 | 252 | 249 | 0.99× [0.97, 1.01] | -0.9% | [-2.7%, +0.9%] | 1.7 pts | 0.8 | yes |
| unique | path | 4096 | churn | ordered | baseline | 6 | 197 | 194 | 0.98× [0.97, 0.99] | -1.8% | [-3.0%, -0.6%] | 1.1 pts | 0.6 | yes |
| unique | path | 4096 | build | ordered | baseline | 6 | 2.12 ms | 2.09 ms | 0.99× [0.98, 1.00] | -1.4% | [-2.4%, -0.4%] | 0.9 pts | 0.9 | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 10 | 115 | 114 | 0.99× [0.98, 0.99] | -1.5% | [-1.9%, -1.0%] | 0.7 pts | 0.6 | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 10 | 2602 | 2633 | 1.01× [1.01, 1.02] | +1.4% | [+0.9%, +1.8%] | 0.6 pts | 0.7 | yes |
| unique | path | 16384 | prefix | ordered | baseline | 10 | 427 | 425 | 1.01× [0.99, 1.03] | +0.7% | [-1.5%, +2.9%] | 3.0 pts | 0.7 | no |
| unique | path | 16384 | churn | ordered | baseline | 10 | 266 | 262 | 0.99× [0.98, 1.00] | -1.3% | [-2.2%, -0.4%] | 1.3 pts | 0.5 | yes |
| unique | path | 16384 | build | ordered | baseline | 10 | 10.90 ms | 10.78 ms | 0.99× [0.98, 0.99] | -1.3% | [-2.1%, -0.5%] | 1.1 pts | 1.2 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 466 | 473 | 1.01× [0.99, 1.03] | +1.3% | [-0.6%, +3.2%] | 2.6 pts | 1.7 | yes |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 7910 | 8179 | 1.01× [0.99, 1.03] | +1.1% | [-0.7%, +2.9%] | 2.6 pts | 1.6 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 5747 | 5804 | 1.01× [0.96, 1.07] | +1.5% | [-3.8%, +6.8%] | 7.4 pts | 0.7 | no |
| unique | path | 262144 | churn | ordered | baseline | 10 | 674 | 693 | 1.01× [0.99, 1.03] | +0.7% | [-1.0%, +2.5%] | 2.4 pts | 1.4 | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 8 | 36.9 | 37.2 | 1.01× [1.00, 1.02] | +1.1% | [+0.5%, +1.8%] | 0.8 pts | 0.5 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 8 | 1661 | 1659 | 1.00× [0.99, 1.00] | -0.3% | [-1.1%, +0.4%] | 0.9 pts | 0.7 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 8 | 3254 | 3364 | 1.03× [1.01, 1.05] | +2.9% | [+1.1%, +4.8%] | 2.3 pts | 0.8 | yes |
| unique | str | 4096 | churn | ordered | baseline | 8 | 86.9 | 85.9 | 0.98× [0.97, 1.00] | -1.8% | [-3.1%, -0.4%] | 1.6 pts | 1.4 | yes |
| unique | str | 4096 | build | ordered | baseline | 8 | 1.03 ms | 1.02 ms | 0.99× [0.98, 1.00] | -0.7% | [-1.8%, +0.4%] | 1.3 pts | 0.4 | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 6 | 47.0 | 49.5 | 1.06× [1.05, 1.06] | +5.3% | [+4.7%, +6.0%] | 0.6 pts | 0.8 | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 1644 | 1684 | 1.02× [1.01, 1.03] | +2.2% | [+1.3%, +3.0%] | 0.8 pts | 0.8 | yes |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 14.0 µs | 14.4 µs | 1.03× [1.02, 1.04] | +2.6% | [+1.5%, +3.7%] | 1.0 pts | 0.6 | yes |
| unique | str | 16384 | churn | ordered | baseline | 6 | 107 | 107 | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.3%] | 0.7 pts | 0.3 | yes |
| unique | str | 16384 | build | ordered | baseline | 6 | 4.97 ms | 4.97 ms | 1.00× [0.99, 1.01] | -0.2% | [-1.2%, +0.8%] | 0.9 pts | 0.9 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 6 | 233 | 241 | 1.04× [1.02, 1.05] | +3.4% | [+2.0%, +4.9%] | 1.4 pts | 1.4 | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 6 | 5844 | 6002 | 1.02× [1.01, 1.04] | +2.3% | [+0.6%, +4.0%] | 1.6 pts | 1.2 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 6 | 931.3 µs | 951.1 µs | 1.02× [1.01, 1.03] | +2.3% | [+1.5%, +3.1%] | 0.8 pts | 0.3 | yes |
| unique | str | 262144 | churn | ordered | baseline | 6 | 324 | 333 | 1.03× [1.02, 1.05] | +3.4% | [+2.0%, +4.8%] | 1.3 pts | 0.8 | yes |
| unique | str | 1048576 | valuesFor | ordered | baseline | 10 | 356 | 373 | 1.04× [1.04, 1.05] | +4.1% | [+3.5%, +4.7%] | 0.9 pts | 0.9 | yes |
| unique | str | 1048576 | valuesBetween | ordered | baseline | 10 | 7054 | 7131 | 1.01× [1.01, 1.02] | +1.4% | [+0.8%, +2.0%] | 0.8 pts | 0.9 | yes |
| unique | str | 1048576 | prefix | ordered | baseline | 10 | 4.25 ms | 4.34 ms | 1.02× [1.01, 1.02] | +1.8% | [+1.4%, +2.3%] | 0.7 pts | 1.2 | yes |
| unique | str | 1048576 | churn | ordered | baseline | 10 | 531 | 543 | 1.02× [1.00, 1.04] | +1.9% | [-0.3%, +4.1%] | 3.1 pts | 2.1 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 48.5 | 48.6 | 1.01× [0.99, 1.03] | +1.0% | [-0.8%, +2.7%] | 1.7 pts | 1.2 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 1859 | 1900 | 1.02× [1.01, 1.03] | +1.9% | [+0.8%, +2.9%] | 1.0 pts | 0.8 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 216 | 221 | 1.02× [1.01, 1.03] | +1.7% | [+0.7%, +2.6%] | 0.9 pts | 0.4 | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 109 | 107 | 0.99× [0.98, 0.99] | -1.4% | [-2.1%, -0.7%] | 0.7 pts | 0.5 | yes |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.28 ms | 1.26 ms | 0.99× [0.97, 1.00] | -1.3% | [-3.0%, +0.5%] | 1.6 pts | 1.0 | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 6 | 69.0 | 69.5 | 1.01× [1.00, 1.01] | +0.6% | [-0.1%, +1.4%] | 0.7 pts | 0.7 | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 6 | 2191 | 2202 | 1.00× [1.00, 1.01] | +0.4% | [+0.0%, +0.8%] | 0.4 pts | 0.4 | yes |
| unique | street | 16384 | prefix | ordered | baseline | 6 | 746 | 744 | 1.00× [0.99, 1.02] | +0.3% | [-1.2%, +1.8%] | 1.4 pts | 0.6 | yes |
| unique | street | 16384 | churn | ordered | baseline | 6 | 145 | 143 | 0.98× [0.97, 1.00] | -2.0% | [-3.5%, -0.4%] | 1.5 pts | 0.7 | yes |
| unique | street | 16384 | build | ordered | baseline | 6 | 6.43 ms | 6.35 ms | 0.99× [0.99, 0.99] | -1.1% | [-1.5%, -0.7%] | 0.4 pts | 0.4 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 8 | 15.1 | 15.3 | 1.01× [1.00, 1.03] | +1.3% | [-0.0%, +2.6%] | 1.6 pts | 0.7 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 962 | 1050 | 1.09× [1.08, 1.10] | +8.4% | [+7.8%, +8.9%] | 0.7 pts | 0.6 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 8 | 44.1 | 43.4 | 0.98× [0.98, 0.99] | -1.5% | [-2.2%, -0.9%] | 0.7 pts | 0.6 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 8 | 604.3 µs | 589.4 µs | 0.97× [0.96, 0.99] | -2.8% | [-4.5%, -1.0%] | 2.1 pts | 1.0 | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 10 | 22.6 | 22.9 | 1.01× [1.01, 1.02] | +1.3% | [+0.9%, +1.7%] | 0.6 pts | 0.6 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 10 | 1535 | 1661 | 1.09× [1.08, 1.09] | +7.8% | [+7.4%, +8.3%] | 0.6 pts | 0.6 | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 10 | 48.3 | 47.3 | 0.98× [0.96, 0.99] | -2.5% | [-4.2%, -0.7%] | 2.5 pts | 0.7 | yes |
| unique | u64 | 16384 | build | ordered | baseline | 10 | 2.92 ms | 2.80 ms | 0.96× [0.95, 0.97] | -4.1% | [-5.0%, -3.2%] | 1.3 pts | 0.5 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 6 | 86.7 | 88.0 | 1.02× [1.00, 1.03] | +1.5% | [+0.1%, +2.9%] | 1.3 pts | 1.2 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 6 | 3572 | 3783 | 1.05× [1.04, 1.07] | +5.2% | [+3.7%, +6.6%] | 1.4 pts | 1.6 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 6 | 175 | 178 | 1.01× [0.99, 1.03] | +0.8% | [-1.1%, +2.7%] | 1.9 pts | 0.4 | yes |
| unique | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 136 | 140 | 1.02× [1.00, 1.03] | +1.6% | [+0.4%, +2.8%] | 1.7 pts | 1.6 | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 3122 | 3304 | 1.06× [1.05, 1.07] | +5.6% | [+4.7%, +6.5%] | 1.2 pts | 1.4 | yes |
| unique | u64 | 1048576 | churn | ordered | baseline | 10 | 311 | 308 | 1.01× [0.99, 1.04] | +1.3% | [-1.2%, +3.8%] | 3.5 pts | 2.0 | no |
| unique | url | 4096 | valuesFor | ordered | baseline | 8 | 60.3 | 60.6 | 1.00× [0.99, 1.01] | +0.1% | [-0.5%, +0.8%] | 0.8 pts | 0.7 | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 8 | 1791 | 1778 | 0.99× [0.98, 1.00] | -0.7% | [-1.6%, +0.2%] | 1.0 pts | 0.7 | yes |
| unique | url | 4096 | prefix | ordered | baseline | 8 | 243 | 243 | 1.00× [0.99, 1.01] | -0.0% | [-0.6%, +0.5%] | 0.6 pts | 0.4 | yes |
| unique | url | 4096 | churn | ordered | baseline | 8 | 134 | 132 | 0.98× [0.98, 0.99] | -1.8% | [-2.3%, -1.3%] | 0.6 pts | 0.4 | yes |
| unique | url | 4096 | build | ordered | baseline | 8 | 1.60 ms | 1.59 ms | 1.00× [0.98, 1.02] | -0.2% | [-2.2%, +1.7%] | 2.3 pts | 0.8 | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 76.4 | 75.8 | 1.00× [0.99, 1.00] | -0.3% | [-1.0%, +0.4%] | 0.7 pts | 0.8 | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 1963 | 1951 | 1.00× [0.99, 1.00] | -0.3% | [-1.0%, +0.5%] | 0.7 pts | 0.7 | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 648 | 653 | 1.00× [0.98, 1.02] | +0.2% | [-1.8%, +2.2%] | 1.9 pts | 1.2 | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 171 | 170 | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.4%] | 0.9 pts | 0.4 | yes |
| unique | url | 16384 | build | ordered | baseline | 6 | 7.62 ms | 7.43 ms | 0.98× [0.96, 0.99] | -2.4% | [-3.8%, -0.9%] | 1.4 pts | 0.6 | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 10 | 373 | 376 | 1.01× [1.00, 1.02] | +1.0% | [+0.4%, +1.6%] | 0.9 pts | 0.8 | yes |
| unique | url | 262144 | valuesBetween | ordered | baseline | 10 | 6347 | 6367 | 1.00× [0.99, 1.01] | +0.3% | [-0.6%, +1.2%] | 1.2 pts | 0.8 | yes |
| unique | url | 262144 | prefix | ordered | baseline | 10 | 28.0 µs | 28.7 µs | 1.01× [0.99, 1.04] | +1.0% | [-1.3%, +3.4%] | 3.3 pts | 0.5 | no |
| unique | url | 262144 | churn | ordered | baseline | 10 | 523 | 508 | 0.98× [0.97, 0.99] | -2.0% | [-3.0%, -1.1%] | 1.4 pts | 1.1 | yes |
| unique | url | 1048576 | valuesFor | ordered | baseline | 8 | 483 | 488 | 1.00× [0.99, 1.02] | +0.5% | [-0.7%, +1.7%] | 1.4 pts | 1.8 | yes |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 8 | 7647 | 7725 | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +0.8%] | 0.9 pts | 0.9 | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 8 | 150.6 µs | 148.7 µs | 0.99× [0.99, 1.00] | -0.6% | [-1.4%, +0.1%] | 0.9 pts | 0.1 | yes |
| unique | url | 1048576 | churn | ordered | baseline | 8 | 711 | 698 | 1.00× [0.98, 1.02] | -0.1% | [-1.8%, +1.6%] | 2.0 pts | 1.3 | yes |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 6 | 29.6 | 30.1 | 1.01× [1.00, 1.03] | +1.5% | [-0.2%, +3.1%] | 1.6 pts | 0.9 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1465 | 1517 | 1.04× [1.03, 1.05] | +3.6% | [+2.7%, +4.5%] | 0.9 pts | 0.6 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 69.3 | 66.4 | 0.96× [0.95, 0.97] | -4.5% | [-5.8%, -3.3%] | 1.2 pts | 1.4 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 6 | 80.7 | 79.8 | 0.98× [0.98, 0.99] | -1.7% | [-2.3%, -1.1%] | 0.6 pts | 0.5 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 6 | 956.2 µs | 956.5 µs | 0.99× [0.98, 1.01] | -0.9% | [-2.5%, +0.7%] | 1.5 pts | 0.4 | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 6 | 38.4 | 40.7 | 1.06× [1.06, 1.06] | +5.6% | [+5.3%, +6.0%] | 0.4 pts | 0.5 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 1592 | 1645 | 1.03× [1.03, 1.04] | +3.2% | [+2.6%, +3.8%] | 0.6 pts | 0.6 | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 83.0 | 81.8 | 0.98× [0.98, 0.99] | -1.6% | [-2.4%, -0.8%] | 0.8 pts | 0.9 | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 6 | 98.1 | 96.7 | 0.99× [0.98, 1.00] | -0.8% | [-1.7%, +0.1%] | 0.9 pts | 0.3 | yes |
| unique | uuid | 16384 | build | ordered | baseline | 6 | 4.63 ms | 4.59 ms | 0.99× [0.98, 1.00] | -0.9% | [-2.2%, +0.3%] | 1.2 pts | 0.8 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 208 | 213 | 1.02× [1.01, 1.03] | +1.9% | [+0.6%, +3.1%] | 1.8 pts | 1.5 | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 5830 | 5924 | 1.02× [1.01, 1.03] | +2.2% | [+1.3%, +3.1%] | 1.3 pts | 1.1 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 412 | 432 | 1.05× [1.04, 1.06] | +5.0% | [+3.9%, +6.0%] | 1.4 pts | 1.4 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 323 | 325 | 1.00× [0.99, 1.02] | +0.4% | [-1.4%, +2.2%] | 2.5 pts | 1.6 | yes |
| unique | uuid | 1048576 | valuesFor | ordered | baseline | 6 | 338 | 339 | 1.00× [0.98, 1.02] | -0.3% | [-2.1%, +1.6%] | 1.8 pts | 2.1 | yes |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 6 | 7235 | 7284 | 1.00× [0.99, 1.02] | +0.4% | [-0.9%, +1.7%] | 1.2 pts | 1.3 | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 6 | 1362 | 1379 | 1.01× [1.00, 1.03] | +1.0% | [-0.4%, +2.5%] | 1.4 pts | 1.3 | yes |
| unique | uuid | 1048576 | churn | ordered | baseline | 6 | 524 | 528 | 1.00× [0.99, 1.02] | +0.5% | [-1.3%, +2.3%] | 1.7 pts | 0.9 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi email n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.43% does not clear the 1.29% median noise floor of the processes
- multi email n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.28%, 1.13%] includes zero
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.87% does not clear the 1.61% median noise floor of the processes
- multi email n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.47%, 2.20%] includes zero
- multi email n=4096 churn: ordered vs baseline: the pooled difference of -1.18% does not clear the 2.38% median noise floor of the processes
- multi email n=4096 build: ordered vs baseline: the pooled difference of -0.61% does not clear the 2.36% median noise floor of the processes
- multi email n=4096 build: ordered vs baseline: the pooled interval [-2.18%, 0.97%] includes zero
- multi email n=16384 valuesFor: ordered vs baseline: the pooled difference of 1.16% does not clear the 1.65% median noise floor of the processes
- multi email n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.25%, 2.57%] includes zero
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.93% does not clear the 1.50% median noise floor of the processes
- multi email n=16384 prefix: ordered vs baseline: the pooled difference of -0.89% does not clear the 1.23% median noise floor of the processes
- multi email n=16384 churn: ordered vs baseline: the pooled difference of 0.16% does not clear the 2.32% median noise floor of the processes
- multi email n=16384 churn: ordered vs baseline: the pooled interval [-0.20%, 0.51%] includes zero
- multi email n=16384 build: ordered vs baseline: the pooled difference of -0.16% does not clear the 1.11% median noise floor of the processes
- multi email n=16384 build: ordered vs baseline: the pooled interval [-1.07%, 0.74%] includes zero
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled difference of 1.03% does not clear the 1.89% median noise floor of the processes
- multi email n=262144 prefix: ordered vs baseline: the pooled difference of 0.96% does not clear the 1.29% median noise floor of the processes
- multi email n=262144 prefix: ordered vs baseline: the pooled interval [-0.04%, 1.96%] includes zero
- multi email n=262144 churn: ordered vs baseline: the pooled difference of -0.53% does not clear the 2.16% median noise floor of the processes
- multi email n=262144 churn: ordered vs baseline: the pooled interval [-3.81%, 2.75%] includes zero
- multi email n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 1.08% does not clear the 1.84% median noise floor of the processes
- multi email n=1048576 churn: ordered vs baseline: the pooled difference of 0.98% does not clear the 2.41% median noise floor of the processes
- multi email n=1048576 churn: ordered vs baseline: the pooled interval [-1.85%, 3.82%] includes zero
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.56% does not clear the 2.58% median noise floor of the processes
- multi path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.42%, 1.53%] includes zero
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of -1.06% does not clear the 2.58% median noise floor of the processes
- multi path n=4096 prefix: ordered vs baseline: the pooled interval [-2.45%, 0.34%] includes zero
- multi path n=4096 churn: ordered vs baseline: the pooled difference of 0.48% does not clear the 1.30% median noise floor of the processes
- multi path n=4096 churn: ordered vs baseline: the pooled interval [-0.50%, 1.46%] includes zero
- multi path n=4096 build: ordered vs baseline: the pooled difference of 0.29% does not clear the 1.01% median noise floor of the processes
- multi path n=4096 build: ordered vs baseline: the pooled interval [-0.21%, 0.78%] includes zero
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.66% does not clear the 1.48% median noise floor of the processes
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.18%, 1.49%] includes zero
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of 0.47% does not clear the 9.01% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-1.30%, 2.23%] includes zero
- multi path n=16384 churn: ordered vs baseline: the pooled difference of 0.64% does not clear the 1.54% median noise floor of the processes
- multi path n=16384 churn: ordered vs baseline: the pooled interval [-0.23%, 1.51%] includes zero
- multi path n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.32% does not clear the 1.64% median noise floor of the processes
- multi path n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.44%, 2.09%] includes zero
- multi path n=262144 valuesFor: ordered vs baseline: 4 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled difference of -0.21% does not clear the 1.17% median noise floor of the processes
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-1.75%, 1.33%] includes zero
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of 2.12% does not clear the 11.50% median noise floor of the processes
- multi path n=262144 prefix: ordered vs baseline: the pooled interval [-0.88%, 5.13%] includes zero
- multi path n=262144 churn: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.46% does not clear the 1.08% median noise floor of the processes
- multi str n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.29%, 1.21%] includes zero
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.42% does not clear the 2.24% median noise floor of the processes
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.87%, 1.71%] includes zero
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of -0.07% does not clear the 3.62% median noise floor of the processes
- multi str n=4096 prefix: ordered vs baseline: the pooled interval [-1.30%, 1.16%] includes zero
- multi str n=4096 churn: ordered vs baseline: the pooled difference of 0.64% does not clear the 1.67% median noise floor of the processes
- multi str n=4096 churn: ordered vs baseline: the pooled interval [-0.22%, 1.50%] includes zero
- multi str n=4096 build: ordered vs baseline: the pooled difference of 0.25% does not clear the 1.53% median noise floor of the processes
- multi str n=4096 build: ordered vs baseline: the pooled interval [-0.63%, 1.12%] includes zero
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.61% does not clear the 1.60% median noise floor of the processes
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.23%, 1.44%] includes zero
- multi str n=16384 prefix: ordered vs baseline: the pooled difference of 0.35% does not clear the 1.62% median noise floor of the processes
- multi str n=16384 prefix: ordered vs baseline: the pooled interval [-1.44%, 2.14%] includes zero
- multi str n=16384 churn: ordered vs baseline: the pooled difference of 0.30% does not clear the 1.66% median noise floor of the processes
- multi str n=16384 churn: ordered vs baseline: the pooled interval [-0.74%, 1.35%] includes zero
- multi str n=16384 build: ordered vs baseline: the pooled difference of 0.51% does not clear the 0.75% median noise floor of the processes
- multi str n=16384 build: ordered vs baseline: the pooled interval [-0.43%, 1.44%] includes zero
- multi str n=262144 churn: ordered vs baseline: the pooled difference of 1.58% does not clear the 2.17% median noise floor of the processes
- multi str n=262144 churn: ordered vs baseline: the pooled interval [-0.99%, 4.16%] includes zero
- multi str n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.38% does not clear the 1.37% median noise floor of the processes
- multi str n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.47%, 1.24%] includes zero
- multi str n=1048576 prefix: ordered vs baseline: the pooled interval [-0.03%, 1.19%] includes zero
- multi str n=1048576 churn: ordered vs baseline: the pooled difference of -1.56% does not clear the 3.33% median noise floor of the processes
- multi str n=1048576 churn: ordered vs baseline: the pooled interval [-4.84%, 1.72%] includes zero
- multi str n=1048576 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 churn: ordered vs baseline: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.14% does not clear the 1.34% median noise floor of the processes
- multi street n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.03%, 1.31%] includes zero
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.73% does not clear the 1.22% median noise floor of the processes
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.94%, 0.48%] includes zero
- multi street n=4096 prefix: ordered vs baseline: the pooled difference of 0.27% does not clear the 1.70% median noise floor of the processes
- multi street n=4096 prefix: ordered vs baseline: the pooled interval [-0.68%, 1.23%] includes zero
- multi street n=4096 churn: ordered vs baseline: the pooled difference of -0.65% does not clear the 1.27% median noise floor of the processes
- multi street n=4096 churn: ordered vs baseline: the pooled interval [-1.48%, 0.18%] includes zero
- multi street n=4096 build: ordered vs baseline: the pooled difference of -0.63% does not clear the 1.43% median noise floor of the processes
- multi street n=4096 build: ordered vs baseline: the pooled interval [-1.55%, 0.30%] includes zero
- multi street n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.69% does not clear the 1.47% median noise floor of the processes
- multi street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.22%, 1.61%] includes zero
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.39% does not clear the 1.72% median noise floor of the processes
- multi street n=16384 prefix: ordered vs baseline: the pooled difference of 0.64% does not clear the 4.41% median noise floor of the processes
- multi street n=16384 prefix: ordered vs baseline: the pooled interval [-0.83%, 2.10%] includes zero
- multi street n=16384 churn: ordered vs baseline: the pooled difference of 0.21% does not clear the 2.16% median noise floor of the processes
- multi street n=16384 churn: ordered vs baseline: the pooled interval [-1.08%, 1.49%] includes zero
- multi street n=16384 build: ordered vs baseline: the pooled difference of -0.68% does not clear the 0.83% median noise floor of the processes
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.03% does not clear the 1.46% median noise floor of the processes
- multi u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.70%, 0.75%] includes zero
- multi u64 n=4096 churn: ordered vs baseline: the pooled difference of -0.85% does not clear the 1.50% median noise floor of the processes
- multi u64 n=4096 churn: ordered vs baseline: the pooled interval [-2.53%, 0.83%] includes zero
- multi u64 n=4096 build: ordered vs baseline: the pooled difference of -0.18% does not clear the 3.08% median noise floor of the processes
- multi u64 n=4096 build: ordered vs baseline: the pooled interval [-1.15%, 0.79%] includes zero
- multi u64 n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.69% does not clear the 1.31% median noise floor of the processes
- multi u64 n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.12%, 1.49%] includes zero
- multi u64 n=16384 churn: ordered vs baseline: the pooled difference of -0.75% does not clear the 1.90% median noise floor of the processes
- multi u64 n=16384 churn: ordered vs baseline: the pooled interval [-1.99%, 0.50%] includes zero
- multi u64 n=262144 churn: ordered vs baseline: the pooled difference of 1.29% does not clear the 2.86% median noise floor of the processes
- multi u64 n=262144 churn: ordered vs baseline: the pooled interval [-2.75%, 5.32%] includes zero
- multi u64 n=262144 churn: ordered vs baseline: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.72% does not clear the 1.96% median noise floor of the processes
- multi u64 n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.31%, 1.74%] includes zero
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 1.89% does not clear the 2.12% median noise floor of the processes
- multi u64 n=1048576 churn: ordered vs baseline: the pooled difference of -4.67% does not clear the 8.89% median noise floor of the processes
- multi u64 n=1048576 churn: ordered vs baseline: the pooled interval [-9.96%, 0.62%] includes zero
- multi url n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.25% does not clear the 0.54% median noise floor of the processes
- multi url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.82%, 0.33%] includes zero
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.66% does not clear the 1.51% median noise floor of the processes
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.12%, 0.81%] includes zero
- multi url n=4096 prefix: ordered vs baseline: the pooled difference of -1.23% does not clear the 1.68% median noise floor of the processes
- multi url n=4096 prefix: ordered vs baseline: the pooled interval [-2.73%, 0.28%] includes zero
- multi url n=4096 churn: ordered vs baseline: the pooled difference of 0.21% does not clear the 1.54% median noise floor of the processes
- multi url n=4096 churn: ordered vs baseline: the pooled interval [-0.40%, 0.83%] includes zero
- multi url n=4096 build: ordered vs baseline: the pooled difference of 0.56% does not clear the 1.50% median noise floor of the processes
- multi url n=4096 build: ordered vs baseline: the pooled interval [-0.37%, 1.50%] includes zero
- multi url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.18% does not clear the 0.86% median noise floor of the processes
- multi url n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.99%, 0.63%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.13% does not clear the 1.27% median noise floor of the processes
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.55%, 0.81%] includes zero
- multi url n=16384 prefix: ordered vs baseline: the pooled difference of 0.21% does not clear the 2.10% median noise floor of the processes
- multi url n=16384 prefix: ordered vs baseline: the pooled interval [-1.58%, 1.99%] includes zero
- multi url n=16384 churn: ordered vs baseline: the pooled difference of -0.27% does not clear the 2.42% median noise floor of the processes
- multi url n=16384 churn: ordered vs baseline: the pooled interval [-1.00%, 0.47%] includes zero
- multi url n=16384 build: ordered vs baseline: the pooled difference of 0.14% does not clear the 1.40% median noise floor of the processes
- multi url n=16384 build: ordered vs baseline: the pooled interval [-0.31%, 0.59%] includes zero
- multi url n=262144 valuesFor: ordered vs baseline: the pooled difference of 1.29% does not clear the 1.75% median noise floor of the processes
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.58% does not clear the 1.49% median noise floor of the processes
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.47%, 1.63%] includes zero
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of -0.71% does not clear the 7.28% median noise floor of the processes
- multi url n=262144 prefix: ordered vs baseline: the pooled interval [-3.68%, 2.27%] includes zero
- multi url n=262144 churn: ordered vs baseline: the pooled difference of 0.95% does not clear the 2.49% median noise floor of the processes
- multi url n=262144 churn: ordered vs baseline: the pooled interval [-1.36%, 3.27%] includes zero
- multi url n=262144 churn: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=1048576 valuesFor: ordered vs baseline: the pooled difference of 1.36% does not clear the 1.52% median noise floor of the processes
- multi url n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.42% does not clear the 1.82% median noise floor of the processes
- multi url n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.99%, 0.16%] includes zero
- multi url n=1048576 prefix: ordered vs baseline: the pooled difference of -0.58% does not clear the 35.37% median noise floor of the processes
- multi url n=1048576 prefix: ordered vs baseline: the pooled interval [-2.10%, 0.94%] includes zero
- multi url n=1048576 churn: ordered vs baseline: the pooled difference of 1.63% does not clear the 2.42% median noise floor of the processes
- multi url n=1048576 churn: ordered vs baseline: the pooled interval [-1.75%, 5.02%] includes zero
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.35% does not clear the 2.60% median noise floor of the processes
- multi uuid n=4096 churn: ordered vs baseline: the pooled difference of -1.01% does not clear the 1.85% median noise floor of the processes
- multi uuid n=4096 churn: ordered vs baseline: the pooled interval [-2.24%, 0.22%] includes zero
- multi uuid n=4096 build: ordered vs baseline: the pooled difference of -0.11% does not clear the 1.47% median noise floor of the processes
- multi uuid n=4096 build: ordered vs baseline: the pooled interval [-1.24%, 1.02%] includes zero
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of 1.30% does not clear the 2.24% median noise floor of the processes
- multi uuid n=16384 churn: ordered vs baseline: the pooled difference of 0.35% does not clear the 1.79% median noise floor of the processes
- multi uuid n=16384 churn: ordered vs baseline: the pooled interval [-0.57%, 1.27%] includes zero
- multi uuid n=16384 build: ordered vs baseline: the pooled difference of 0.13% does not clear the 1.68% median noise floor of the processes
- multi uuid n=16384 build: ordered vs baseline: the pooled interval [-0.43%, 0.70%] includes zero
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of 1.52% does not clear the 1.68% median noise floor of the processes
- multi uuid n=262144 churn: ordered vs baseline: the pooled difference of -0.68% does not clear the 1.90% median noise floor of the processes
- multi uuid n=262144 churn: ordered vs baseline: the pooled interval [-3.85%, 2.49%] includes zero
- multi uuid n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.73% does not clear the 1.42% median noise floor of the processes
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.48%, 1.94%] includes zero
- multi uuid n=1048576 prefix: ordered vs baseline: the pooled interval [-0.09%, 2.50%] includes zero
- multi uuid n=1048576 prefix: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=1048576 churn: ordered vs baseline: the pooled difference of -1.63% does not clear the 3.85% median noise floor of the processes
- multi uuid n=1048576 churn: ordered vs baseline: the pooled interval [-5.21%, 1.94%] includes zero
- unique email n=4096 valuesFor: ordered vs baseline: the pooled difference of -1.47% does not clear the 2.31% median noise floor of the processes
- unique email n=4096 valuesFor: ordered vs baseline: the pooled interval [-3.30%, 0.35%] includes zero
- unique email n=4096 build: ordered vs baseline: the pooled difference of -0.86% does not clear the 2.49% median noise floor of the processes
- unique email n=4096 build: ordered vs baseline: the pooled interval [-2.57%, 0.84%] includes zero
- unique email n=16384 churn: ordered vs baseline: the pooled difference of -0.31% does not clear the 2.39% median noise floor of the processes
- unique email n=16384 churn: ordered vs baseline: the pooled interval [-2.13%, 1.50%] includes zero
- unique email n=16384 build: ordered vs baseline: the pooled difference of -1.35% does not clear the 1.90% median noise floor of the processes
- unique email n=262144 churn: ordered vs baseline: the pooled difference of 0.19% does not clear the 2.08% median noise floor of the processes
- unique email n=262144 churn: ordered vs baseline: the pooled interval [-1.32%, 1.69%] includes zero
- unique email n=1048576 churn: ordered vs baseline: the pooled difference of 0.43% does not clear the 1.77% median noise floor of the processes
- unique email n=1048576 churn: ordered vs baseline: the pooled interval [-1.78%, 2.64%] includes zero
- unique email n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.42% does not clear the 1.87% median noise floor of the processes
- unique path n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.68%, 1.52%] includes zero
- unique path n=4096 prefix: ordered vs baseline: the pooled difference of -0.90% does not clear the 1.77% median noise floor of the processes
- unique path n=4096 prefix: ordered vs baseline: the pooled interval [-2.70%, 0.90%] includes zero
- unique path n=4096 churn: ordered vs baseline: the pooled difference of -1.83% does not clear the 1.91% median noise floor of the processes
- unique path n=4096 build: ordered vs baseline: the pooled difference of -1.39% does not clear the 1.98% median noise floor of the processes
- unique path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 1.37% does not clear the 1.44% median noise floor of the processes
- unique path n=16384 prefix: ordered vs baseline: the pooled difference of 0.69% does not clear the 6.38% median noise floor of the processes
- unique path n=16384 prefix: ordered vs baseline: the pooled interval [-1.49%, 2.87%] includes zero
- unique path n=16384 churn: ordered vs baseline: the pooled difference of -1.29% does not clear the 2.77% median noise floor of the processes
- unique path n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.57%, 3.18%] includes zero
- unique path n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=262144 valuesBetween: ordered vs baseline: the pooled difference of 1.10% does not clear the 1.43% median noise floor of the processes
- unique path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.73%, 2.93%] includes zero
- unique path n=262144 valuesBetween: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=262144 prefix: ordered vs baseline: the pooled difference of 1.47% does not clear the 12.71% median noise floor of the processes
- unique path n=262144 prefix: ordered vs baseline: the pooled interval [-3.85%, 6.80%] includes zero
- unique path n=262144 churn: ordered vs baseline: the pooled difference of 0.74% does not clear the 1.62% median noise floor of the processes
- unique path n=262144 churn: ordered vs baseline: the pooled interval [-0.99%, 2.48%] includes zero
- unique str n=4096 valuesFor: ordered vs baseline: the pooled difference of 1.12% does not clear the 1.46% median noise floor of the processes
- unique str n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.35% does not clear the 1.74% median noise floor of the processes
- unique str n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.11%, 0.41%] includes zero
- unique str n=4096 build: ordered vs baseline: the pooled difference of -0.71% does not clear the 2.81% median noise floor of the processes
- unique str n=4096 build: ordered vs baseline: the pooled interval [-1.77%, 0.36%] includes zero
- unique str n=16384 valuesBetween: ordered vs baseline: the pooled difference of 2.16% does not clear the 2.17% median noise floor of the processes
- unique str n=16384 churn: ordered vs baseline: the pooled difference of -0.42% does not clear the 2.69% median noise floor of the processes
- unique str n=16384 churn: ordered vs baseline: the pooled interval [-1.11%, 0.28%] includes zero
- unique str n=16384 build: ordered vs baseline: the pooled difference of -0.19% does not clear the 1.05% median noise floor of the processes
- unique str n=16384 build: ordered vs baseline: the pooled interval [-1.18%, 0.80%] includes zero
- unique str n=262144 prefix: ordered vs baseline: the pooled difference of 2.28% does not clear the 4.27% median noise floor of the processes
- unique str n=1048576 churn: ordered vs baseline: the pooled interval [-0.25%, 4.12%] includes zero
- unique str n=1048576 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.95% does not clear the 1.32% median noise floor of the processes
- unique street n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.78%, 2.69%] includes zero
- unique street n=4096 prefix: ordered vs baseline: the pooled difference of 1.66% does not clear the 2.01% median noise floor of the processes
- unique street n=4096 build: ordered vs baseline: the pooled difference of -1.26% does not clear the 1.97% median noise floor of the processes
- unique street n=4096 build: ordered vs baseline: the pooled interval [-2.98%, 0.47%] includes zero
- unique street n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.63% does not clear the 1.25% median noise floor of the processes
- unique street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.10%, 1.36%] includes zero
- unique street n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.40% does not clear the 1.54% median noise floor of the processes
- unique street n=16384 prefix: ordered vs baseline: the pooled difference of 0.31% does not clear the 3.29% median noise floor of the processes
- unique street n=16384 prefix: ordered vs baseline: the pooled interval [-1.17%, 1.80%] includes zero
- unique street n=16384 churn: ordered vs baseline: the pooled difference of -1.95% does not clear the 3.07% median noise floor of the processes
- unique u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of 1.28% does not clear the 2.82% median noise floor of the processes
- unique u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.02%, 2.58%] includes zero
- unique u64 n=16384 churn: ordered vs baseline: the pooled difference of -2.46% does not clear the 2.59% median noise floor of the processes
- unique u64 n=262144 valuesFor: ordered vs baseline: the pooled difference of 1.50% does not clear the 2.24% median noise floor of the processes
- unique u64 n=262144 churn: ordered vs baseline: the pooled difference of 0.80% does not clear the 2.44% median noise floor of the processes
- unique u64 n=262144 churn: ordered vs baseline: the pooled interval [-1.15%, 2.75%] includes zero
- unique u64 n=1048576 churn: ordered vs baseline: the pooled difference of 1.27% does not clear the 1.68% median noise floor of the processes
- unique u64 n=1048576 churn: ordered vs baseline: the pooled interval [-1.23%, 3.76%] includes zero
- unique u64 n=1048576 churn: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 churn: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.12% does not clear the 0.67% median noise floor of the processes
- unique url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.52%, 0.76%] includes zero
- unique url n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.70% does not clear the 1.79% median noise floor of the processes
- unique url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.58%, 0.18%] includes zero
- unique url n=4096 prefix: ordered vs baseline: the pooled difference of -0.01% does not clear the 1.25% median noise floor of the processes
- unique url n=4096 prefix: ordered vs baseline: the pooled interval [-0.55%, 0.52%] includes zero
- unique url n=4096 build: ordered vs baseline: the pooled difference of -0.23% does not clear the 4.22% median noise floor of the processes
- unique url n=4096 build: ordered vs baseline: the pooled interval [-2.18%, 1.73%] includes zero
- unique url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.31% does not clear the 1.06% median noise floor of the processes
- unique url n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.03%, 0.40%] includes zero
- unique url n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.29% does not clear the 1.26% median noise floor of the processes
- unique url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.03%, 0.46%] includes zero
- unique url n=16384 prefix: ordered vs baseline: the pooled difference of 0.20% does not clear the 2.26% median noise floor of the processes
- unique url n=16384 prefix: ordered vs baseline: the pooled interval [-1.79%, 2.20%] includes zero
- unique url n=16384 churn: ordered vs baseline: the pooled difference of -0.60% does not clear the 1.95% median noise floor of the processes
- unique url n=16384 churn: ordered vs baseline: the pooled interval [-1.55%, 0.36%] includes zero
- unique url n=262144 valuesFor: ordered vs baseline: the pooled difference of 1.03% does not clear the 1.54% median noise floor of the processes
- unique url n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.32% does not clear the 1.47% median noise floor of the processes
- unique url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.55%, 1.19%] includes zero
- unique url n=262144 prefix: ordered vs baseline: the pooled difference of 1.04% does not clear the 8.77% median noise floor of the processes
- unique url n=262144 prefix: ordered vs baseline: the pooled interval [-1.34%, 3.43%] includes zero
- unique url n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.49% does not clear the 0.80% median noise floor of the processes
- unique url n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.69%, 1.67%] includes zero
- unique url n=1048576 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.08% does not clear the 1.17% median noise floor of the processes
- unique url n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.68%, 0.83%] includes zero
- unique url n=1048576 prefix: ordered vs baseline: the pooled difference of -0.64% does not clear the 19.17% median noise floor of the processes
- unique url n=1048576 prefix: ordered vs baseline: the pooled interval [-1.39%, 0.11%] includes zero
- unique url n=1048576 churn: ordered vs baseline: the pooled difference of -0.09% does not clear the 2.74% median noise floor of the processes
- unique url n=1048576 churn: ordered vs baseline: the pooled interval [-1.78%, 1.60%] includes zero
- unique uuid n=4096 valuesFor: ordered vs baseline: the pooled difference of 1.46% does not clear the 2.61% median noise floor of the processes
- unique uuid n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.22%, 3.15%] includes zero
- unique uuid n=4096 build: ordered vs baseline: the pooled difference of -0.91% does not clear the 2.94% median noise floor of the processes
- unique uuid n=4096 build: ordered vs baseline: the pooled interval [-2.53%, 0.72%] includes zero
- unique uuid n=16384 churn: ordered vs baseline: the pooled difference of -0.82% does not clear the 2.94% median noise floor of the processes
- unique uuid n=16384 churn: ordered vs baseline: the pooled interval [-1.71%, 0.08%] includes zero
- unique uuid n=16384 build: ordered vs baseline: the pooled difference of -0.92% does not clear the 1.52% median noise floor of the processes
- unique uuid n=16384 build: ordered vs baseline: the pooled interval [-2.19%, 0.35%] includes zero
- unique uuid n=262144 churn: ordered vs baseline: the pooled difference of 0.40% does not clear the 1.98% median noise floor of the processes
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-1.41%, 2.21%] includes zero
- unique uuid n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=1048576 valuesFor: ordered vs baseline: the pooled difference of -0.25% does not clear the 0.88% median noise floor of the processes
- unique uuid n=1048576 valuesFor: ordered vs baseline: the pooled interval [-2.13%, 1.62%] includes zero
- unique uuid n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.42% does not clear the 0.86% median noise floor of the processes
- unique uuid n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.87%, 1.71%] includes zero
- unique uuid n=1048576 prefix: ordered vs baseline: the pooled difference of 1.01% does not clear the 1.40% median noise floor of the processes
- unique uuid n=1048576 prefix: ordered vs baseline: the pooled interval [-0.43%, 2.45%] includes zero
- unique uuid n=1048576 churn: ordered vs baseline: the pooled difference of 0.50% does not clear the 1.75% median noise floor of the processes
- unique uuid n=1048576 churn: ordered vs baseline: the pooled interval [-1.28%, 2.27%] includes zero
