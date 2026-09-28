| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 6 | 50.1 | 52.1 | 1.04× [1.04, 1.05] | +3.9% | [+3.4%, +4.3%] | 0.4 pts | 0.4 | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 6 | 3001 | 2964 | 0.98× [0.97, 1.00] | -1.7% | [-2.9%, -0.4%] | 1.1 pts | 0.6 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 6 | 76.5 | 76.8 | 1.00× [1.00, 1.01] | +0.4% | [+0.1%, +0.7%] | 0.3 pts | 0.4 | yes |
| multi | email | 4096 | churn | ordered | baseline | 6 | 65.8 | 70.2 | 1.07× [1.06, 1.08] | +6.6% | [+5.7%, +7.6%] | 0.9 pts | 0.5 | yes |
| multi | email | 4096 | build | ordered | baseline | 6 | 5.36 ms | 5.95 ms | 1.11× [1.10, 1.12] | +9.8% | [+9.2%, +10.4%] | 0.6 pts | 0.9 | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 8 | 61.6 | 65.7 | 1.06× [1.05, 1.07] | +5.6% | [+4.8%, +6.4%] | 0.9 pts | 0.9 | yes |
| multi | email | 16384 | valuesBetween | ordered | baseline | 8 | 3604 | 3612 | 1.00× [0.99, 1.01] | -0.0% | [-0.8%, +0.8%] | 0.9 pts | 0.8 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 8 | 96.4 | 96.0 | 1.01× [0.99, 1.02] | +0.7% | [-0.9%, +2.2%] | 1.8 pts | 0.7 | yes |
| multi | email | 16384 | churn | ordered | baseline | 8 | 117 | 123 | 1.05× [1.04, 1.06] | +5.0% | [+4.1%, +6.0%] | 1.1 pts | 0.6 | yes |
| multi | email | 16384 | build | ordered | baseline | 8 | 30.26 ms | 33.64 ms | 1.12× [1.11, 1.14] | +11.0% | [+9.8%, +12.1%] | 1.4 pts | 0.9 | yes |
| multi | email | 262144 | valuesFor | ordered | baseline | 10 | 261 | 282 | 1.07× [1.07, 1.08] | +6.8% | [+6.1%, +7.5%] | 1.0 pts | 1.0 | yes |
| multi | email | 262144 | valuesBetween | ordered | baseline | 10 | 9918 | 9928 | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +0.9%] | 1.2 pts | 1.0 | yes |
| multi | email | 262144 | prefix | ordered | baseline | 10 | 355 | 353 | 0.99× [0.99, 1.00] | -0.6% | [-1.3%, +0.2%] | 1.1 pts | 0.8 | yes |
| multi | email | 262144 | churn | ordered | baseline | 10 | 398 | 405 | 1.02× [1.00, 1.05] | +2.1% | [-0.4%, +4.7%] | 3.6 pts | 1.7 | no |
| multi | email | 1048576 | valuesFor | ordered | baseline | 10 | 423 | 436 | 1.05× [1.03, 1.07] | +4.7% | [+3.0%, +6.4%] | 2.4 pts | 1.9 | yes |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 10 | 11.7 µs | 11.8 µs | 1.00× [0.99, 1.02] | +0.4% | [-0.9%, +1.6%] | 1.7 pts | 1.8 | yes |
| multi | email | 1048576 | prefix | ordered | baseline | 10 | 689 | 692 | 1.00× [0.99, 1.02] | +0.1% | [-1.4%, +1.6%] | 2.2 pts | 1.7 | yes |
| multi | email | 1048576 | churn | ordered | baseline | 10 | 620 | 601 | 0.99× [0.96, 1.03] | -0.5% | [-3.8%, +2.8%] | 4.6 pts | 1.6 | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 106 | 109 | 1.03× [1.03, 1.04] | +3.4% | [+2.9%, +3.9%] | 0.5 pts | 0.5 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4132 | 4227 | 1.03× [1.01, 1.04] | +2.5% | [+1.1%, +3.9%] | 1.3 pts | 0.8 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 328 | 342 | 1.04× [1.03, 1.06] | +3.9% | [+2.5%, +5.3%] | 1.3 pts | 0.6 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 158 | 166 | 1.04× [1.03, 1.06] | +4.3% | [+3.2%, +5.3%] | 1.0 pts | 0.6 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 11.55 ms | 12.61 ms | 1.09× [1.08, 1.10] | +8.1% | [+7.2%, +9.0%] | 0.8 pts | 0.9 | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 10 | 137 | 140 | 1.02× [1.02, 1.02] | +2.0% | [+1.6%, +2.4%] | 0.6 pts | 0.6 | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 10 | 4762 | 4804 | 1.01× [1.00, 1.01] | +0.8% | [+0.1%, +1.5%] | 1.0 pts | 0.7 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 10 | 622 | 632 | 1.01× [0.98, 1.04] | +0.7% | [-2.1%, +3.5%] | 3.9 pts | 0.7 | no |
| multi | path | 16384 | churn | ordered | baseline | 10 | 259 | 262 | 1.01× [1.00, 1.03] | +1.4% | [+0.3%, +2.6%] | 1.6 pts | 1.1 | yes |
| multi | path | 16384 | build | ordered | baseline | 10 | 68.44 ms | 72.19 ms | 1.06× [1.05, 1.06] | +5.2% | [+4.4%, +6.1%] | 1.2 pts | 0.9 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 566 | 570 | 1.01× [0.99, 1.03] | +0.8% | [-0.9%, +2.6%] | 2.4 pts | 1.8 | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 13.8 µs | 13.8 µs | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +0.9%] | 1.1 pts | 0.8 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 12.8 µs | 12.6 µs | 1.00× [0.97, 1.03] | +0.1% | [-2.8%, +3.0%] | 4.1 pts | 0.3 | no |
| multi | path | 262144 | churn | ordered | baseline | 10 | 873 | 871 | 0.98× [0.95, 1.02] | -2.1% | [-5.8%, +1.7%] | 5.2 pts | 2.5 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 6 | 62.1 | 64.3 | 1.03× [1.03, 1.04] | +3.4% | [+2.6%, +4.2%] | 0.7 pts | 0.6 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 6 | 3494 | 3530 | 1.01× [1.00, 1.03] | +1.4% | [+0.1%, +2.7%] | 1.3 pts | 0.5 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 6 | 7768 | 7766 | 1.00× [0.98, 1.02] | +0.2% | [-1.6%, +1.9%] | 1.7 pts | 0.4 | yes |
| multi | str | 4096 | churn | ordered | baseline | 6 | 83.7 | 87.7 | 1.05× [1.04, 1.06] | +4.8% | [+3.5%, +6.1%] | 1.2 pts | 0.8 | yes |
| multi | str | 4096 | build | ordered | baseline | 6 | 6.48 ms | 7.32 ms | 1.13× [1.12, 1.14] | +11.4% | [+10.8%, +12.0%] | 0.6 pts | 1.0 | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 6 | 76.3 | 77.7 | 1.02× [1.02, 1.03] | +2.3% | [+1.6%, +3.0%] | 0.7 pts | 0.7 | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 6 | 3740 | 3774 | 1.01× [1.00, 1.01] | +0.8% | [+0.3%, +1.2%] | 0.4 pts | 0.3 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 6 | 35.1 µs | 35.2 µs | 1.00× [0.99, 1.01] | -0.1% | [-1.4%, +1.1%] | 1.2 pts | 1.0 | yes |
| multi | str | 16384 | churn | ordered | baseline | 6 | 128 | 134 | 1.04× [1.04, 1.05] | +4.3% | [+3.7%, +4.9%] | 0.5 pts | 0.3 | yes |
| multi | str | 16384 | build | ordered | baseline | 6 | 35.19 ms | 38.41 ms | 1.08× [1.07, 1.10] | +7.8% | [+6.7%, +9.0%] | 1.1 pts | 0.8 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 338 | 342 | 1.02× [1.01, 1.03] | +1.9% | [+0.9%, +3.0%] | 1.5 pts | 1.1 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 10.9 µs | 10.9 µs | 1.01× [0.99, 1.02] | +0.6% | [-0.6%, +1.7%] | 1.6 pts | 1.4 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.71 ms | 1.73 ms | 1.01× [1.00, 1.03] | +1.4% | [+0.3%, +2.5%] | 1.5 pts | 1.3 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 427 | 436 | 1.02× [0.99, 1.05] | +2.2% | [-0.8%, +5.1%] | 4.1 pts | 2.7 | no |
| multi | str | 1048576 | valuesFor | ordered | baseline | 10 | 461 | 463 | 1.01× [1.00, 1.03] | +1.4% | [-0.5%, +3.2%] | 2.6 pts | 2.2 | yes |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 10 | 12.0 µs | 12.1 µs | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +0.9%] | 1.1 pts | 1.0 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 10 | 7.43 ms | 7.45 ms | 1.01× [1.00, 1.01] | +0.5% | [+0.0%, +1.0%] | 0.7 pts | 1.3 | yes |
| multi | str | 1048576 | churn | ordered | baseline | 10 | 631 | 648 | 1.04× [1.01, 1.08] | +4.1% | [+0.5%, +7.7%] | 5.0 pts | 1.5 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 8 | 55.0 | 57.6 | 1.05× [1.04, 1.06] | +4.8% | [+4.1%, +5.5%] | 0.9 pts | 0.6 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 8 | 2372 | 2393 | 1.00× [0.99, 1.02] | +0.5% | [-0.6%, +1.5%] | 1.2 pts | 0.5 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 8 | 256 | 257 | 1.00× [0.99, 1.02] | +0.3% | [-1.1%, +1.8%] | 1.7 pts | 0.8 | yes |
| multi | street | 4096 | churn | ordered | baseline | 8 | 91.6 | 96.7 | 1.06× [1.05, 1.06] | +5.4% | [+4.7%, +6.1%] | 0.8 pts | 0.6 | yes |
| multi | street | 4096 | build | ordered | baseline | 8 | 2.77 ms | 3.34 ms | 1.21× [1.20, 1.22] | +17.3% | [+16.7%, +17.9%] | 0.7 pts | 0.8 | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 8 | 78.1 | 82.5 | 1.06× [1.05, 1.07] | +5.4% | [+4.4%, +6.4%] | 1.2 pts | 0.9 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 8 | 2894 | 2912 | 1.01× [1.00, 1.01] | +0.5% | [-0.3%, +1.3%] | 1.0 pts | 0.6 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 8 | 961 | 957 | 1.00× [0.99, 1.02] | +0.3% | [-1.1%, +1.6%] | 1.6 pts | 0.6 | yes |
| multi | street | 16384 | churn | ordered | baseline | 8 | 144 | 152 | 1.06× [1.04, 1.08] | +5.5% | [+3.9%, +7.0%] | 1.9 pts | 1.1 | yes |
| multi | street | 16384 | build | ordered | baseline | 8 | 16.24 ms | 18.41 ms | 1.14× [1.13, 1.15] | +12.2% | [+11.5%, +12.8%] | 0.8 pts | 0.8 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 8 | 36.6 | 38.7 | 1.06× [1.05, 1.07] | +5.7% | [+4.9%, +6.4%] | 0.9 pts | 0.8 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2690 | 2685 | 1.00× [0.99, 1.01] | +0.1% | [-0.8%, +1.1%] | 1.2 pts | 0.4 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 8 | 45.5 | 48.9 | 1.07× [1.05, 1.09] | +6.4% | [+4.7%, +8.0%] | 2.0 pts | 1.3 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 8 | 4.00 ms | 4.57 ms | 1.14× [1.13, 1.15] | +12.4% | [+11.9%, +13.0%] | 0.7 pts | 1.2 | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 45.1 | 47.4 | 1.05× [1.05, 1.06] | +4.9% | [+4.5%, +5.3%] | 0.4 pts | 0.7 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3625 | 3615 | 1.00× [0.99, 1.00] | -0.2% | [-0.6%, +0.2%] | 0.4 pts | 0.4 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 60.4 | 66.5 | 1.11× [1.10, 1.12] | +10.0% | [+9.2%, +10.9%] | 0.8 pts | 0.6 | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 20.41 ms | 22.76 ms | 1.11× [1.10, 1.13] | +10.3% | [+9.3%, +11.2%] | 0.9 pts | 0.9 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 160 | 167 | 1.03× [1.03, 1.04] | +3.4% | [+2.7%, +4.0%] | 0.9 pts | 0.9 | yes |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 8566 | 8680 | 1.01× [1.00, 1.01] | +0.7% | [+0.2%, +1.3%] | 0.8 pts | 0.7 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 269 | 299 | 1.08× [1.04, 1.12] | +7.2% | [+3.8%, +10.7%] | 4.8 pts | 2.1 | no |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 209 | 215 | 1.03× [1.02, 1.04] | +3.1% | [+2.0%, +4.1%] | 1.4 pts | 1.5 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 8410 | 8404 | 1.01× [1.00, 1.03] | +1.2% | [-0.4%, +2.9%] | 2.3 pts | 1.9 | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 10 | 396 | 419 | 1.05× [1.02, 1.08] | +4.5% | [+1.5%, +7.5%] | 4.2 pts | 1.8 | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 6 | 82.8 | 85.4 | 1.03× [1.03, 1.04] | +3.2% | [+3.0%, +3.4%] | 0.2 pts | 0.2 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 6 | 3708 | 3760 | 1.01× [1.00, 1.03] | +1.5% | [-0.4%, +3.3%] | 1.7 pts | 1.1 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 6 | 409 | 421 | 1.03× [1.02, 1.03] | +2.6% | [+2.0%, +3.1%] | 0.5 pts | 0.3 | yes |
| multi | url | 4096 | churn | ordered | baseline | 6 | 121 | 128 | 1.06× [1.05, 1.07] | +5.7% | [+5.0%, +6.4%] | 0.7 pts | 0.4 | yes |
| multi | url | 4096 | build | ordered | baseline | 6 | 9.23 ms | 10.10 ms | 1.09× [1.08, 1.10] | +8.4% | [+7.8%, +9.0%] | 0.6 pts | 0.6 | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 98.0 | 101 | 1.03× [1.02, 1.03] | +2.8% | [+2.4%, +3.2%] | 0.4 pts | 0.4 | yes |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4024 | 4049 | 1.01× [1.00, 1.02] | +1.0% | [-0.2%, +2.2%] | 1.2 pts | 0.9 | yes |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 1319 | 1355 | 1.03× [1.02, 1.04] | +2.9% | [+1.9%, +4.0%] | 1.0 pts | 0.6 | yes |
| multi | url | 16384 | churn | ordered | baseline | 6 | 181 | 190 | 1.05× [1.05, 1.06] | +5.0% | [+4.3%, +5.8%] | 0.7 pts | 0.4 | yes |
| multi | url | 16384 | build | ordered | baseline | 6 | 49.42 ms | 53.45 ms | 1.08× [1.08, 1.09] | +7.8% | [+7.1%, +8.5%] | 0.7 pts | 0.5 | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 10 | 467 | 467 | 1.01× [1.00, 1.02] | +1.1% | [+0.2%, +2.1%] | 1.3 pts | 1.0 | yes |
| multi | url | 262144 | valuesBetween | ordered | baseline | 10 | 11.7 µs | 11.7 µs | 1.00× [1.00, 1.01] | +0.2% | [-0.3%, +0.8%] | 0.7 pts | 0.6 | yes |
| multi | url | 262144 | prefix | ordered | baseline | 10 | 52.7 µs | 54.1 µs | 1.01× [0.99, 1.04] | +1.2% | [-1.5%, +3.8%] | 3.7 pts | 0.5 | no |
| multi | url | 262144 | churn | ordered | baseline | 10 | 642 | 622 | 0.99× [0.96, 1.02] | -1.4% | [-4.6%, +1.9%] | 4.5 pts | 2.3 | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 10 | 556 | 565 | 1.01× [0.99, 1.02] | +0.8% | [-0.8%, +2.3%] | 2.1 pts | 2.1 | yes |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 10 | 12.6 µs | 12.7 µs | 1.00× [1.00, 1.01] | +0.4% | [-0.4%, +1.2%] | 1.1 pts | 1.0 | yes |
| multi | url | 1048576 | prefix | ordered | baseline | 10 | 241.5 µs | 241.0 µs | 1.00× [0.99, 1.01] | -0.1% | [-1.4%, +1.2%] | 1.8 pts | 0.1 | yes |
| multi | url | 1048576 | churn | ordered | baseline | 10 | 775 | 763 | 0.99× [0.97, 1.02] | -0.5% | [-3.1%, +2.0%] | 3.6 pts | 1.6 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 55.3 | 57.7 | 1.04× [1.03, 1.05] | +3.8% | [+3.3%, +4.3%] | 0.5 pts | 0.5 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3319 | 3347 | 1.00× [0.98, 1.02] | -0.2% | [-2.0%, +1.6%] | 1.7 pts | 0.8 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 88.8 | 89.7 | 1.01× [1.00, 1.02] | +0.9% | [+0.2%, +1.6%] | 0.6 pts | 0.6 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 75.3 | 79.9 | 1.06× [1.05, 1.07] | +6.0% | [+5.1%, +6.9%] | 0.9 pts | 0.4 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 6.03 ms | 6.71 ms | 1.11× [1.10, 1.12] | +10.0% | [+9.4%, +10.6%] | 0.6 pts | 0.7 | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 67.8 | 70.5 | 1.04× [1.03, 1.05] | +3.8% | [+2.5%, +5.1%] | 1.3 pts | 0.7 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3705 | 3701 | 1.00× [0.99, 1.01] | -0.0% | [-0.9%, +0.9%] | 0.8 pts | 0.5 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 112 | 113 | 1.01× [1.00, 1.01] | +0.6% | [+0.0%, +1.2%] | 0.6 pts | 0.4 | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 133 | 141 | 1.06× [1.05, 1.07] | +5.8% | [+4.9%, +6.7%] | 0.8 pts | 0.3 | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 32.49 ms | 35.63 ms | 1.10× [1.08, 1.11] | +8.8% | [+7.7%, +10.0%] | 1.1 pts | 0.8 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 316 | 336 | 1.08× [1.06, 1.10] | +7.3% | [+5.9%, +8.7%] | 1.9 pts | 1.9 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 11.1 µs | 11.2 µs | 1.00× [1.00, 1.01] | +0.4% | [-0.3%, +1.1%] | 1.0 pts | 0.9 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 732 | 738 | 1.00× [0.99, 1.01] | -0.1% | [-1.0%, +0.7%] | 1.2 pts | 1.0 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 460 | 475 | 1.02× [0.99, 1.04] | +1.5% | [-0.9%, +3.9%] | 3.3 pts | 1.3 | no |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 10 | 427 | 445 | 1.04× [1.02, 1.06] | +3.8% | [+1.9%, +5.7%] | 2.7 pts | 3.3 | yes |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 10 | 12.4 µs | 12.3 µs | 1.00× [0.98, 1.01] | -0.4% | [-1.5%, +0.8%] | 1.6 pts | 1.5 | yes |
| multi | uuid | 1048576 | prefix | ordered | baseline | 10 | 2287 | 2272 | 1.01× [0.99, 1.03] | +1.0% | [-1.3%, +3.4%] | 3.3 pts | 2.9 | no |
| multi | uuid | 1048576 | churn | ordered | baseline | 10 | 615 | 630 | 1.01× [0.96, 1.07] | +1.3% | [-3.9%, +6.5%] | 7.3 pts | 1.5 | no |
| unique | email | 4096 | valuesFor | ordered | baseline | 6 | 26.3 | 39.8 | 1.52× [1.49, 1.55] | +34.2% | [+33.0%, +35.4%] | 1.1 pts | 1.3 | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 6 | 1269 | 337 | 0.27× [0.26, 0.27] | -277.0% | [-279.3%, -274.7%] | 2.2 pts | 0.5 | yes |
| unique | email | 4096 | prefix | ordered | baseline | 6 | 56.0 | 59.7 | 1.07× [1.06, 1.08] | +6.4% | [+5.7%, +7.0%] | 0.6 pts | 0.7 | yes |
| unique | email | 4096 | churn | ordered | baseline | 6 | 67.5 | 89.6 | 1.33× [1.31, 1.35] | +24.8% | [+23.8%, +25.8%] | 1.0 pts | 1.0 | yes |
| unique | email | 4096 | build | ordered | baseline | 6 | 848.8 µs | 1.23 ms | 1.46× [1.43, 1.49] | +31.5% | [+30.3%, +32.7%] | 1.1 pts | 0.6 | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 37.6 | 56.4 | 1.50× [1.49, 1.51] | +33.3% | [+32.8%, +33.9%] | 0.5 pts | 1.1 | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 1591 | 443 | 0.28× [0.28, 0.28] | -260.0% | [-263.1%, -256.8%] | 3.0 pts | 1.1 | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 70.3 | 84.6 | 1.20× [1.20, 1.21] | +16.8% | [+16.4%, +17.3%] | 0.4 pts | 0.5 | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 94.1 | 125 | 1.33× [1.29, 1.36] | +24.5% | [+22.5%, +26.6%] | 2.0 pts | 0.9 | yes |
| unique | email | 16384 | build | ordered | baseline | 6 | 4.07 ms | 6.81 ms | 1.68× [1.67, 1.69] | +40.4% | [+40.2%, +40.7%] | 0.2 pts | 0.3 | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 10 | 218 | 216 | 1.01× [0.98, 1.04] | +1.0% | [-1.7%, +3.7%] | 3.8 pts | 2.9 | no |
| unique | email | 262144 | valuesBetween | ordered | baseline | 10 | 3985 | 1495 | 0.37× [0.37, 0.38] | -167.6% | [-173.2%, -162.0%] | 7.8 pts | 2.1 | yes |
| unique | email | 262144 | prefix | ordered | baseline | 10 | 277 | 229 | 0.83× [0.82, 0.85] | -20.0% | [-22.1%, -17.9%] | 3.0 pts | 1.8 | no |
| unique | email | 262144 | churn | ordered | baseline | 10 | 322 | 365 | 1.15× [1.13, 1.18] | +13.2% | [+11.2%, +15.3%] | 2.8 pts | 2.0 | no |
| unique | email | 1048576 | valuesFor | ordered | baseline | 10 | 381 | 397 | 1.06× [1.05, 1.07] | +5.6% | [+4.5%, +6.7%] | 1.6 pts | 1.8 | yes |
| unique | email | 1048576 | valuesBetween | ordered | baseline | 10 | 6804 | 1955 | 0.29× [0.28, 0.29] | -248.1% | [-251.0%, -245.3%] | 4.0 pts | 1.4 | yes |
| unique | email | 1048576 | prefix | ordered | baseline | 10 | 503 | 421 | 0.85× [0.82, 0.87] | -18.3% | [-21.3%, -15.3%] | 4.2 pts | 3.6 | no |
| unique | email | 1048576 | churn | ordered | baseline | 10 | 611 | 631 | 1.02× [1.00, 1.04] | +1.9% | [-0.4%, +4.2%] | 3.2 pts | 1.9 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 8 | 87.8 | 102 | 1.16× [1.15, 1.17] | +13.7% | [+13.3%, +14.2%] | 0.6 pts | 0.8 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 8 | 2117 | 674 | 0.32× [0.32, 0.32] | -213.7% | [-216.3%, -211.2%] | 3.0 pts | 0.6 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 8 | 250 | 229 | 0.92× [0.91, 0.93] | -8.8% | [-10.0%, -7.5%] | 1.5 pts | 0.7 | yes |
| unique | path | 4096 | churn | ordered | baseline | 8 | 205 | 223 | 1.09× [1.07, 1.10] | +8.0% | [+6.6%, +9.4%] | 1.7 pts | 1.0 | yes |
| unique | path | 4096 | build | ordered | baseline | 8 | 2.15 ms | 3.03 ms | 1.41× [1.39, 1.43] | +29.0% | [+28.1%, +29.9%] | 1.0 pts | 1.5 | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 10 | 115 | 129 | 1.12× [1.12, 1.13] | +10.7% | [+10.4%, +11.1%] | 0.5 pts | 0.5 | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 10 | 2618 | 829 | 0.32× [0.31, 0.32] | -215.9% | [-219.1%, -212.8%] | 4.4 pts | 1.4 | yes |
| unique | path | 16384 | prefix | ordered | baseline | 10 | 423 | 313 | 0.74× [0.73, 0.75] | -35.2% | [-37.4%, -33.0%] | 3.1 pts | 0.5 | yes |
| unique | path | 16384 | churn | ordered | baseline | 10 | 284 | 289 | 1.02× [1.00, 1.05] | +2.4% | [-0.1%, +4.9%] | 3.5 pts | 1.4 | no |
| unique | path | 16384 | build | ordered | baseline | 10 | 11.39 ms | 15.76 ms | 1.37× [1.36, 1.39] | +27.1% | [+26.3%, +27.9%] | 1.1 pts | 0.8 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 529 | 481 | 0.89× [0.86, 0.91] | -12.5% | [-15.8%, -9.3%] | 4.5 pts | 3.1 | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 7495 | 2901 | 0.38× [0.37, 0.39] | -162.3% | [-167.2%, -157.4%] | 6.8 pts | 1.6 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 6004 | 1970 | 0.32× [0.30, 0.34] | -212.9% | [-233.2%, -192.7%] | 28.3 pts | 0.8 | yes |
| unique | path | 262144 | churn | ordered | baseline | 10 | 902 | 893 | 0.99× [0.96, 1.01] | -1.4% | [-4.1%, +1.3%] | 3.8 pts | 2.2 | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 6 | 38.3 | 46.1 | 1.20× [1.20, 1.21] | +16.9% | [+16.4%, +17.4%] | 0.5 pts | 0.7 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 6 | 1613 | 375 | 0.23× [0.23, 0.23] | -332.4% | [-336.4%, -328.3%] | 3.9 pts | 0.9 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 6 | 3157 | 563 | 0.18× [0.18, 0.18] | -461.7% | [-466.5%, -456.8%] | 4.6 pts | 1.1 | yes |
| unique | str | 4096 | churn | ordered | baseline | 6 | 91.3 | 114 | 1.25× [1.24, 1.27] | +20.3% | [+19.1%, +21.4%] | 1.1 pts | 0.9 | yes |
| unique | str | 4096 | build | ordered | baseline | 6 | 1.05 ms | 1.48 ms | 1.41× [1.39, 1.43] | +29.2% | [+28.3%, +30.2%] | 0.9 pts | 0.8 | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 10 | 51.5 | 60.7 | 1.17× [1.17, 1.18] | +14.9% | [+14.3%, +15.5%] | 0.8 pts | 1.0 | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 10 | 1709 | 469 | 0.27× [0.27, 0.28] | -264.4% | [-266.9%, -261.9%] | 3.5 pts | 1.0 | yes |
| unique | str | 16384 | prefix | ordered | baseline | 10 | 14.2 µs | 2508 | 0.18× [0.18, 0.18] | -464.3% | [-471.1%, -457.5%] | 9.5 pts | 1.1 | yes |
| unique | str | 16384 | churn | ordered | baseline | 10 | 121 | 140 | 1.18× [1.16, 1.21] | +15.3% | [+13.6%, +17.1%] | 2.4 pts | 0.9 | yes |
| unique | str | 16384 | build | ordered | baseline | 10 | 5.18 ms | 7.52 ms | 1.46× [1.45, 1.46] | +31.3% | [+31.1%, +31.5%] | 0.3 pts | 0.5 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 271 | 161 | 0.59× [0.57, 0.61] | -69.5% | [-75.7%, -63.4%] | 8.6 pts | 2.7 | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 5074 | 1241 | 0.25× [0.24, 0.25] | -302.1% | [-309.7%, -294.5%] | 10.7 pts | 1.6 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 747.2 µs | 133.0 µs | 0.19× [0.18, 0.19] | -438.8% | [-456.2%, -421.4%] | 24.3 pts | 0.7 | yes |
| unique | str | 262144 | churn | ordered | baseline | 10 | 354 | 311 | 0.89× [0.85, 0.93] | -12.3% | [-17.0%, -7.6%] | 6.6 pts | 4.0 | no |
| unique | str | 1048576 | valuesFor | ordered | baseline | 10 | 393 | 335 | 0.85× [0.84, 0.87] | -17.2% | [-19.2%, -15.2%] | 2.8 pts | 1.8 | yes |
| unique | str | 1048576 | valuesBetween | ordered | baseline | 10 | 7072 | 1868 | 0.26× [0.26, 0.27] | -277.8% | [-280.1%, -275.4%] | 3.3 pts | 0.9 | yes |
| unique | str | 1048576 | prefix | ordered | baseline | 10 | 4.37 ms | 927.7 µs | 0.21× [0.21, 0.21] | -371.6% | [-375.7%, -367.5%] | 5.7 pts | 1.4 | yes |
| unique | str | 1048576 | churn | ordered | baseline | 10 | 594 | 622 | 1.03× [1.01, 1.06] | +3.2% | [+0.6%, +5.9%] | 3.7 pts | 2.9 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 49.5 | 55.4 | 1.13× [1.12, 1.14] | +11.2% | [+10.4%, +12.0%] | 0.8 pts | 0.9 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 1855 | 478 | 0.26× [0.25, 0.26] | -288.5% | [-294.8%, -282.1%] | 6.1 pts | 1.3 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 216 | 107 | 0.50× [0.49, 0.51] | -100.1% | [-104.4%, -95.9%] | 4.1 pts | 1.9 | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 112 | 123 | 1.09× [1.08, 1.11] | +8.7% | [+7.5%, +9.9%] | 1.1 pts | 1.0 | yes |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.29 ms | 1.94 ms | 1.50× [1.50, 1.51] | +33.5% | [+33.2%, +33.9%] | 0.3 pts | 0.3 | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 10 | 71.6 | 72.1 | 1.01× [1.01, 1.02] | +1.3% | [+0.8%, +1.8%] | 0.7 pts | 0.9 | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 10 | 2221 | 574 | 0.26× [0.26, 0.26] | -286.4% | [-288.4%, -284.5%] | 2.7 pts | 0.8 | yes |
| unique | street | 16384 | prefix | ordered | baseline | 10 | 729 | 206 | 0.29× [0.28, 0.29] | -250.7% | [-255.7%, -245.8%] | 6.9 pts | 1.0 | yes |
| unique | street | 16384 | churn | ordered | baseline | 10 | 157 | 157 | 0.99× [0.96, 1.02] | -1.1% | [-4.4%, +2.2%] | 4.6 pts | 2.1 | no |
| unique | street | 16384 | build | ordered | baseline | 10 | 6.45 ms | 9.16 ms | 1.41× [1.40, 1.42] | +29.2% | [+28.7%, +29.7%] | 0.7 pts | 1.2 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 10 | 16.6 | 20.2 | 1.18× [1.13, 1.23] | +15.4% | [+11.8%, +19.0%] | 5.0 pts | 2.3 | no |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 10 | 1048 | 241 | 0.23× [0.23, 0.23] | -335.4% | [-338.6%, -332.2%] | 4.5 pts | 1.8 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 10 | 43.9 | 46.5 | 1.06× [1.05, 1.07] | +5.8% | [+5.0%, +6.5%] | 1.0 pts | 1.2 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 10 | 612.9 µs | 629.4 µs | 1.03× [1.02, 1.05] | +3.3% | [+2.2%, +4.5%] | 1.6 pts | 1.1 | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 6 | 22.9 | 32.2 | 1.40× [1.39, 1.42] | +28.7% | [+28.0%, +29.4%] | 0.7 pts | 1.1 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 1618 | 289 | 0.18× [0.18, 0.18] | -461.0% | [-462.4%, -459.6%] | 1.3 pts | 0.3 | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 6 | 44.5 | 64.0 | 1.44× [1.41, 1.47] | +30.5% | [+29.2%, +31.8%] | 1.2 pts | 1.2 | yes |
| unique | u64 | 16384 | build | ordered | baseline | 6 | 2.80 ms | 3.55 ms | 1.27× [1.26, 1.28] | +21.4% | [+20.8%, +22.0%] | 0.6 pts | 0.6 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 10 | 82.9 | 46.9 | 0.57× [0.55, 0.59] | -76.7% | [-83.3%, -70.1%] | 9.2 pts | 2.7 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 2561 | 380 | 0.15× [0.14, 0.15] | -588.5% | [-617.2%, -559.8%] | 40.1 pts | 2.5 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 10 | 164 | 115 | 0.70× [0.67, 0.73] | -42.6% | [-48.8%, -36.5%] | 8.7 pts | 1.1 | no |
| unique | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 151 | 126 | 0.83× [0.81, 0.85] | -20.9% | [-23.7%, -18.0%] | 4.0 pts | 3.4 | no |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 3350 | 774 | 0.23× [0.23, 0.23] | -337.6% | [-344.4%, -330.8%] | 9.5 pts | 2.0 | yes |
| unique | u64 | 1048576 | churn | ordered | baseline | 10 | 341 | 276 | 0.83× [0.77, 0.90] | -21.0% | [-30.4%, -11.7%] | 13.1 pts | 3.5 | no |
| unique | url | 4096 | valuesFor | ordered | baseline | 8 | 61.3 | 69.2 | 1.13× [1.12, 1.14] | +11.6% | [+11.0%, +12.2%] | 0.8 pts | 1.2 | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 8 | 1704 | 536 | 0.31× [0.31, 0.32] | -217.8% | [-220.8%, -214.7%] | 3.7 pts | 1.1 | yes |
| unique | url | 4096 | prefix | ordered | baseline | 8 | 243 | 169 | 0.70× [0.69, 0.70] | -43.1% | [-44.3%, -41.9%] | 1.5 pts | 1.0 | yes |
| unique | url | 4096 | churn | ordered | baseline | 8 | 138 | 153 | 1.09× [1.07, 1.11] | +8.1% | [+6.2%, +10.1%] | 2.4 pts | 1.4 | yes |
| unique | url | 4096 | build | ordered | baseline | 8 | 1.59 ms | 2.21 ms | 1.39× [1.37, 1.41] | +28.2% | [+27.1%, +29.2%] | 1.2 pts | 0.6 | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 10 | 77.6 | 82.7 | 1.07× [1.06, 1.08] | +6.6% | [+5.8%, +7.3%] | 1.0 pts | 1.2 | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 10 | 1943 | 565 | 0.29× [0.28, 0.30] | -245.6% | [-252.3%, -238.9%] | 9.3 pts | 3.7 | yes |
| unique | url | 16384 | prefix | ordered | baseline | 10 | 658 | 257 | 0.39× [0.38, 0.40] | -156.2% | [-161.5%, -151.0%] | 7.3 pts | 2.7 | yes |
| unique | url | 16384 | churn | ordered | baseline | 10 | 187 | 199 | 1.06× [1.03, 1.08] | +5.3% | [+3.2%, +7.4%] | 3.0 pts | 1.2 | no |
| unique | url | 16384 | build | ordered | baseline | 10 | 7.97 ms | 10.36 ms | 1.29× [1.27, 1.31] | +22.4% | [+21.4%, +23.4%] | 1.4 pts | 1.1 | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 10 | 426 | 350 | 0.83× [0.81, 0.86] | -20.4% | [-23.8%, -16.9%] | 4.9 pts | 3.8 | no |
| unique | url | 262144 | valuesBetween | ordered | baseline | 10 | 6040 | 1861 | 0.30× [0.30, 0.31] | -228.0% | [-231.0%, -225.0%] | 4.2 pts | 0.5 | yes |
| unique | url | 262144 | prefix | ordered | baseline | 10 | 25.5 µs | 5384 | 0.21× [0.21, 0.22] | -369.9% | [-376.0%, -363.8%] | 8.5 pts | 0.4 | yes |
| unique | url | 262144 | churn | ordered | baseline | 10 | 622 | 582 | 0.96× [0.94, 0.99] | -3.8% | [-6.3%, -1.4%] | 3.4 pts | 2.4 | no |
| unique | url | 1048576 | valuesFor | ordered | baseline | 6 | 537 | 574 | 1.06× [1.06, 1.07] | +6.0% | [+5.3%, +6.7%] | 0.7 pts | 0.8 | yes |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 6 | 7732 | 2688 | 0.35× [0.34, 0.35] | -187.3% | [-192.5%, -182.1%] | 5.0 pts | 2.4 | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 6 | 136.5 µs | 32.9 µs | 0.24× [0.24, 0.24] | -317.1% | [-324.0%, -310.3%] | 6.5 pts | 0.2 | yes |
| unique | url | 1048576 | churn | ordered | baseline | 6 | 818 | 901 | 1.09× [1.08, 1.11] | +8.7% | [+7.5%, +9.9%] | 1.2 pts | 0.4 | yes |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 8 | 32.1 | 44.6 | 1.38× [1.33, 1.43] | +27.6% | [+25.0%, +30.2%] | 3.1 pts | 3.7 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 8 | 1530 | 406 | 0.26× [0.26, 0.27] | -279.0% | [-284.6%, -273.4%] | 6.7 pts | 1.4 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 8 | 67.1 | 74.1 | 1.11× [1.10, 1.11] | +9.7% | [+9.1%, +10.3%] | 0.7 pts | 1.1 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 8 | 84.0 | 99.0 | 1.18× [1.17, 1.19] | +15.4% | [+14.7%, +16.1%] | 0.8 pts | 0.5 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 8 | 967.5 µs | 1.45 ms | 1.50× [1.49, 1.51] | +33.3% | [+32.9%, +33.6%] | 0.4 pts | 0.6 | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 10 | 41.8 | 51.0 | 1.23× [1.21, 1.24] | +18.4% | [+17.5%, +19.2%] | 1.2 pts | 1.9 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 10 | 1659 | 411 | 0.25× [0.24, 0.26] | -296.5% | [-310.9%, -282.2%] | 20.1 pts | 5.0 | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 10 | 82.6 | 82.5 | 1.00× [0.99, 1.01] | +0.0% | [-0.9%, +1.0%] | 1.3 pts | 1.6 | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 10 | 108 | 119 | 1.09× [1.06, 1.12] | +8.5% | [+5.9%, +11.1%] | 3.6 pts | 1.2 | no |
| unique | uuid | 16384 | build | ordered | baseline | 10 | 4.79 ms | 6.25 ms | 1.31× [1.30, 1.32] | +23.6% | [+23.2%, +24.0%] | 0.6 pts | 0.7 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 253 | 266 | 1.02× [0.98, 1.05] | +1.7% | [-1.9%, +5.2%] | 4.9 pts | 2.6 | no |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 5428 | 1470 | 0.27× [0.27, 0.27] | -269.1% | [-273.2%, -265.0%] | 5.7 pts | 1.0 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 442 | 249 | 0.57× [0.56, 0.58] | -76.4% | [-79.2%, -73.5%] | 3.9 pts | 1.8 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 373 | 394 | 1.05× [1.00, 1.10] | +4.8% | [+0.5%, +9.1%] | 6.1 pts | 1.2 | no |
| unique | uuid | 1048576 | valuesFor | ordered | baseline | 8 | 361 | 475 | 1.33× [1.31, 1.35] | +24.8% | [+23.9%, +25.7%] | 1.1 pts | 1.1 | yes |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 8 | 6996 | 2448 | 0.35× [0.35, 0.36] | -183.3% | [-188.2%, -178.4%] | 5.8 pts | 2.6 | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 8 | 1308 | 519 | 0.40× [0.38, 0.42] | -149.6% | [-161.3%, -138.0%] | 13.9 pts | 5.1 | yes |
| unique | uuid | 1048576 | churn | ordered | baseline | 8 | 595 | 666 | 1.14× [1.12, 1.16] | +12.2% | [+10.4%, +14.1%] | 2.2 pts | 2.4 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi email n=4096 valuesBetween: ordered vs baseline: the pooled difference of -1.66% does not clear the 2.44% median noise floor of the processes
- multi email n=4096 prefix: ordered vs baseline: the pooled difference of 0.41% does not clear the 1.72% median noise floor of the processes
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.03% does not clear the 1.81% median noise floor of the processes
- multi email n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.81%, 0.76%] includes zero
- multi email n=16384 prefix: ordered vs baseline: the pooled difference of 0.65% does not clear the 1.24% median noise floor of the processes
- multi email n=16384 prefix: ordered vs baseline: the pooled interval [-0.88%, 2.18%] includes zero
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.10% does not clear the 1.59% median noise floor of the processes
- multi email n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.74%, 0.93%] includes zero
- multi email n=262144 prefix: ordered vs baseline: the pooled difference of -0.55% does not clear the 1.29% median noise floor of the processes
- multi email n=262144 prefix: ordered vs baseline: the pooled interval [-1.31%, 0.21%] includes zero
- multi email n=262144 churn: ordered vs baseline: the pooled interval [-0.41%, 4.69%] includes zero
- multi email n=262144 churn: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.36% does not clear the 1.64% median noise floor of the processes
- multi email n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.88%, 1.61%] includes zero
- multi email n=1048576 prefix: ordered vs baseline: the pooled difference of 0.10% does not clear the 1.27% median noise floor of the processes
- multi email n=1048576 prefix: ordered vs baseline: the pooled interval [-1.44%, 1.64%] includes zero
- multi email n=1048576 churn: ordered vs baseline: the pooled difference of -0.54% does not clear the 2.28% median noise floor of the processes
- multi email n=1048576 churn: ordered vs baseline: the pooled interval [-3.83%, 2.75%] includes zero
- multi email n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=4096 prefix: ordered vs baseline: the pooled difference of 3.91% does not clear the 4.46% median noise floor of the processes
- multi path n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.76% does not clear the 1.39% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled difference of 0.72% does not clear the 5.55% median noise floor of the processes
- multi path n=16384 prefix: ordered vs baseline: the pooled interval [-2.10%, 3.54%] includes zero
- multi path n=16384 churn: ordered vs baseline: the pooled difference of 1.43% does not clear the 2.17% median noise floor of the processes
- multi path n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.84% does not clear the 1.61% median noise floor of the processes
- multi path n=262144 valuesFor: ordered vs baseline: the pooled interval [-0.90%, 2.57%] includes zero
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.11% does not clear the 1.72% median noise floor of the processes
- multi path n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.67%, 0.89%] includes zero
- multi path n=262144 prefix: ordered vs baseline: the pooled difference of 0.09% does not clear the 10.36% median noise floor of the processes
- multi path n=262144 prefix: ordered vs baseline: the pooled interval [-2.84%, 3.01%] includes zero
- multi path n=262144 churn: ordered vs baseline: the pooled interval [-5.80%, 1.65%] includes zero
- multi path n=262144 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.39% does not clear the 1.95% median noise floor of the processes
- multi str n=4096 prefix: ordered vs baseline: the pooled difference of 0.15% does not clear the 3.26% median noise floor of the processes
- multi str n=4096 prefix: ordered vs baseline: the pooled interval [-1.64%, 1.94%] includes zero
- multi str n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.78% does not clear the 1.71% median noise floor of the processes
- multi str n=16384 prefix: ordered vs baseline: the pooled difference of -0.13% does not clear the 1.57% median noise floor of the processes
- multi str n=16384 prefix: ordered vs baseline: the pooled interval [-1.39%, 1.12%] includes zero
- multi str n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.55% does not clear the 1.69% median noise floor of the processes
- multi str n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.59%, 1.69%] includes zero
- multi str n=262144 churn: ordered vs baseline: the pooled interval [-0.75%, 5.07%] includes zero
- multi str n=262144 churn: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.46%, 3.19%] includes zero
- multi str n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.13% does not clear the 1.14% median noise floor of the processes
- multi str n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.68%, 0.94%] includes zero
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.49% does not clear the 1.93% median noise floor of the processes
- multi street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.56%, 1.53%] includes zero
- multi street n=4096 prefix: ordered vs baseline: the pooled difference of 0.33% does not clear the 2.44% median noise floor of the processes
- multi street n=4096 prefix: ordered vs baseline: the pooled interval [-1.10%, 1.77%] includes zero
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.52% does not clear the 1.48% median noise floor of the processes
- multi street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.28%, 1.32%] includes zero
- multi street n=16384 prefix: ordered vs baseline: the pooled difference of 0.27% does not clear the 4.90% median noise floor of the processes
- multi street n=16384 prefix: ordered vs baseline: the pooled interval [-1.10%, 1.64%] includes zero
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.13% does not clear the 2.04% median noise floor of the processes
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.83%, 1.10%] includes zero
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.21% does not clear the 1.35% median noise floor of the processes
- multi u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.59%, 0.17%] includes zero
- multi u64 n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.73% does not clear the 1.31% median noise floor of the processes
- multi u64 n=262144 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 1.21% does not clear the 1.64% median noise floor of the processes
- multi u64 n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.43%, 2.86%] includes zero
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled difference of 1.47% does not clear the 2.23% median noise floor of the processes
- multi url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.36%, 3.30%] includes zero
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled difference of 1.00% does not clear the 1.86% median noise floor of the processes
- multi url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.21%, 2.20%] includes zero
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.25% does not clear the 1.29% median noise floor of the processes
- multi url n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.27%, 0.76%] includes zero
- multi url n=262144 prefix: ordered vs baseline: the pooled difference of 1.16% does not clear the 7.88% median noise floor of the processes
- multi url n=262144 prefix: ordered vs baseline: the pooled interval [-1.47%, 3.78%] includes zero
- multi url n=262144 churn: ordered vs baseline: the pooled difference of -1.39% does not clear the 2.12% median noise floor of the processes
- multi url n=262144 churn: ordered vs baseline: the pooled interval [-4.64%, 1.87%] includes zero
- multi url n=262144 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=1048576 valuesFor: ordered vs baseline: the pooled difference of 0.77% does not clear the 1.80% median noise floor of the processes
- multi url n=1048576 valuesFor: ordered vs baseline: the pooled interval [-0.75%, 2.30%] includes zero
- multi url n=1048576 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi url n=1048576 valuesBetween: ordered vs baseline: the pooled difference of 0.41% does not clear the 1.33% median noise floor of the processes
- multi url n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-0.38%, 1.21%] includes zero
- multi url n=1048576 prefix: ordered vs baseline: the pooled difference of -0.09% does not clear the 52.22% median noise floor of the processes
- multi url n=1048576 prefix: ordered vs baseline: the pooled interval [-1.40%, 1.22%] includes zero
- multi url n=1048576 churn: ordered vs baseline: the pooled difference of -0.52% does not clear the 2.31% median noise floor of the processes
- multi url n=1048576 churn: ordered vs baseline: the pooled interval [-3.09%, 2.04%] includes zero
- multi url n=1048576 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.19% does not clear the 2.55% median noise floor of the processes
- multi uuid n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.00%, 1.61%] includes zero
- multi uuid n=4096 prefix: ordered vs baseline: the pooled difference of 0.91% does not clear the 1.62% median noise floor of the processes
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.01% does not clear the 1.23% median noise floor of the processes
- multi uuid n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.89%, 0.87%] includes zero
- multi uuid n=16384 prefix: ordered vs baseline: the pooled difference of 0.62% does not clear the 1.71% median noise floor of the processes
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled difference of 0.41% does not clear the 1.30% median noise floor of the processes
- multi uuid n=262144 valuesBetween: ordered vs baseline: the pooled interval [-0.30%, 1.13%] includes zero
- multi uuid n=262144 prefix: ordered vs baseline: the pooled difference of -0.15% does not clear the 1.29% median noise floor of the processes
- multi uuid n=262144 prefix: ordered vs baseline: the pooled interval [-1.03%, 0.74%] includes zero
- multi uuid n=262144 churn: ordered vs baseline: the pooled difference of 1.48% does not clear the 2.02% median noise floor of the processes
- multi uuid n=262144 churn: ordered vs baseline: the pooled interval [-0.91%, 3.88%] includes zero
- multi uuid n=1048576 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the pooled difference of -0.37% does not clear the 1.52% median noise floor of the processes
- multi uuid n=1048576 valuesBetween: ordered vs baseline: the pooled interval [-1.53%, 0.78%] includes zero
- multi uuid n=1048576 valuesBetween: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=1048576 prefix: ordered vs baseline: the pooled difference of 1.01% does not clear the 1.76% median noise floor of the processes
- multi uuid n=1048576 prefix: ordered vs baseline: the pooled interval [-1.33%, 3.36%] includes zero
- multi uuid n=1048576 prefix: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 churn: ordered vs baseline: the pooled difference of 1.30% does not clear the 2.61% median noise floor of the processes
- multi uuid n=1048576 churn: ordered vs baseline: the pooled interval [-3.93%, 6.53%] includes zero
- unique email n=262144 valuesFor: ordered vs baseline: the pooled difference of 0.99% does not clear the 1.72% median noise floor of the processes
- unique email n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.71%, 3.68%] includes zero
- unique email n=262144 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique email n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=1048576 prefix: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=1048576 churn: ordered vs baseline: the pooled interval [-0.35%, 4.24%] includes zero
- unique path n=16384 churn: ordered vs baseline: the pooled difference of 2.40% does not clear the 2.98% median noise floor of the processes
- unique path n=16384 churn: ordered vs baseline: the pooled interval [-0.11%, 4.91%] includes zero
- unique path n=262144 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 churn: ordered vs baseline: the pooled difference of -1.40% does not clear the 1.49% median noise floor of the processes
- unique path n=262144 churn: ordered vs baseline: the pooled interval [-4.12%, 1.31%] includes zero
- unique path n=262144 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 churn: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=1048576 churn: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=1048576 churn: ordered vs baseline: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique street n=16384 churn: ordered vs baseline: the pooled difference of -1.09% does not clear the 2.47% median noise floor of the processes
- unique street n=16384 churn: ordered vs baseline: the pooled interval [-4.35%, 2.18%] includes zero
- unique street n=16384 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 churn: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=16384 valuesBetween: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=16384 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 valuesFor: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 valuesBetween: ordered vs baseline: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 prefix: ordered vs baseline: the pooled difference of 0.03% does not clear the 0.61% median noise floor of the processes
- unique uuid n=16384 prefix: ordered vs baseline: the pooled interval [-0.93%, 0.98%] includes zero
- unique uuid n=16384 prefix: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=262144 valuesFor: ordered vs baseline: the pooled interval [-1.85%, 5.15%] includes zero
- unique uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 prefix: ordered vs baseline: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
