| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi | str | 4096 | valuesFor | ordered | btree-sets | 6 | 60.3 | 175 | 2.92× [2.90, 2.94] | +65.7% | [+65.5%, +66.0%] | 0.2 pts | 0.9 | yes | yes |
| multi | str | 4096 | valuesFor | ordered | hashed | 6 | 58.9 | 35.2 | 0.61× [0.60, 0.62] | -64.3% | [-66.7%, -62.0%] | 2.3 pts | 0.5 | yes | yes |
| multi | str | 4096 | valuesFor | ordered | map-sets | 6 | 59.6 | 94.0 | 1.59× [1.57, 1.61] | +36.9% | [+36.1%, +37.8%] | 0.8 pts | 1.3 | yes | yes |
| multi | str | 4096 | valuesBetween | ordered | btree-sets | 6 | 3272 | 7257 | 2.23× [2.21, 2.26] | +55.2% | [+54.7%, +55.7%] | 0.5 pts | 0.7 | yes | yes |
| multi | str | 4096 | valuesBetween | ordered | hashed | 6 | 3402 | 52.4 µs | 15.38× [15.16, 15.60] | +93.5% | [+93.4%, +93.6%] | 0.1 pts | 1.2 | yes | yes |
| multi | str | 4096 | valuesBetween | ordered | map-sets | 6 | 3323 | 53.9 µs | 16.34× [16.08, 16.61] | +93.9% | [+93.8%, +94.0%] | 0.1 pts | 1.1 | yes | yes |
| multi | str | 4096 | prefix | ordered | btree-sets | 6 | 6798 | 18.6 µs | 2.76× [2.72, 2.80] | +63.8% | [+63.2%, +64.3%] | 0.5 pts | 0.5 | yes | yes |
| multi | str | 4096 | prefix | ordered | hashed | 6 | 7260 | 54.8 µs | 7.62× [7.45, 7.81] | +86.9% | [+86.6%, +87.2%] | 0.3 pts | 1.2 | yes | yes |
| multi | str | 4096 | prefix | ordered | map-sets | 6 | 6918 | 64.5 µs | 9.40× [9.19, 9.61] | +89.4% | [+89.1%, +89.6%] | 0.2 pts | 1.0 | yes | yes |
| multi | str | 4096 | churn | ordered | btree-sets | 6 | 83.4 | 178 | 2.14× [2.11, 2.16] | +53.2% | [+52.6%, +53.8%] | 0.6 pts | 1.3 | yes | yes |
| multi | str | 4096 | churn | ordered | hashed | 6 | 81.6 | 57.1 | 0.69× [0.69, 0.70] | -44.0% | [-45.2%, -42.9%] | 1.1 pts | 1.0 | yes | yes |
| multi | str | 4096 | churn | ordered | map-sets | 6 | 82.1 | 61.4 | 0.75× [0.74, 0.75] | -33.7% | [-34.6%, -32.7%] | 0.9 pts | 0.8 | yes | yes |
| multi | str | 4096 | build | ordered | btree-sets | 6 | 8.31 ms | 16.61 ms | 2.00× [1.99, 2.01] | +50.0% | [+49.7%, +50.3%] | 0.3 pts | 0.9 | yes | yes |
| multi | str | 4096 | build | ordered | hashed | 6 | 8.28 ms | 5.66 ms | 0.68× [0.68, 0.68] | -46.6% | [-47.2%, -46.1%] | 0.5 pts | 0.4 | yes | yes |
| multi | str | 4096 | build | ordered | map-sets | 6 | 8.28 ms | 6.04 ms | 0.73× [0.72, 0.74] | -37.2% | [-38.9%, -35.5%] | 1.6 pts | 1.8 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | btree-sets | 12 | 73.1 | 215 | 2.95× [2.93, 2.97] | +66.1% | [+65.9%, +66.3%] | 0.3 pts | 1.6 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | hashed | 12 | 72.6 | 45.3 | 0.62× [0.62, 0.63] | -60.4% | [-61.5%, -59.4%] | 1.6 pts | 1.8 | yes | yes |
| multi | str | 16384 | valuesFor | ordered | map-sets | 12 | 72.1 | 101 | 1.40× [1.39, 1.42] | +28.8% | [+28.2%, +29.3%] | 0.7 pts | 2.0 | yes | yes |
| multi | str | 16384 | valuesBetween | ordered | btree-sets | 12 | 3636 | 7692 | 2.12× [2.11, 2.13] | +52.7% | [+52.5%, +53.0%] | 0.4 pts | 1.3 | yes | yes |
| multi | str | 16384 | valuesBetween | ordered | hashed | 12 | 3952 | 216.4 µs | 54.70× [53.98, 55.44] | +98.2% | [+98.1%, +98.2%] | 0.0 pts | 1.4 | yes | yes |
| multi | str | 16384 | valuesBetween | ordered | map-sets | 12 | 3997 | 209.9 µs | 52.07× [50.43, 53.83] | +98.1% | [+98.0%, +98.1%] | 0.1 pts | 1.6 | yes | yes |
| multi | str | 16384 | prefix | ordered | btree-sets | 12 | 33.6 µs | 76.7 µs | 2.29× [2.29, 2.30] | +56.4% | [+56.2%, +56.5%] | 0.4 pts | 1.2 | yes | yes |
| multi | str | 16384 | prefix | ordered | hashed | 12 | 34.7 µs | 244.6 µs | 6.93× [6.88, 6.98] | +85.6% | [+85.5%, +85.7%] | 0.7 pts | 5.3 | yes | yes |
| multi | str | 16384 | prefix | ordered | map-sets | 12 | 34.3 µs | 282.3 µs | 8.23× [8.17, 8.29] | +87.8% | [+87.8%, +87.9%] | 0.2 pts | 1.3 | yes | yes |
| multi | str | 16384 | churn | ordered | btree-sets | 12 | 112 | 262 | 2.28× [2.25, 2.31] | +56.2% | [+55.6%, +56.7%] | 0.6 pts | 1.4 | yes | yes |
| multi | str | 16384 | churn | ordered | hashed | 12 | 107 | 74.4 | 0.69× [0.68, 0.70] | -44.8% | [-46.3%, -43.2%] | 1.6 pts | 1.1 | yes | yes |
| multi | str | 16384 | churn | ordered | map-sets | 12 | 108 | 88.1 | 0.81× [0.79, 0.83] | -23.4% | [-26.1%, -20.7%] | 3.4 pts | 2.3 | no | yes |
| multi | str | 16384 | build | ordered | btree-sets | 12 | 39.64 ms | 89.20 ms | 2.24× [2.23, 2.26] | +55.5% | [+55.2%, +55.7%] | 0.4 pts | 0.8 | yes | yes |
| multi | str | 16384 | build | ordered | hashed | 12 | 39.24 ms | 26.73 ms | 0.68× [0.68, 0.68] | -46.6% | [-47.2%, -46.0%] | 0.6 pts | 0.7 | yes | yes |
| multi | str | 16384 | build | ordered | map-sets | 12 | 39.63 ms | 31.34 ms | 0.79× [0.77, 0.80] | -27.4% | [-29.6%, -25.2%] | 3.4 pts | 1.5 | yes | yes |
| multi | str | 262144 | valuesFor | ordered | btree-sets | 12 | 272 | 580 | 2.16× [2.11, 2.22] | +53.8% | [+52.6%, +54.9%] | 1.7 pts | 4.4 | yes | yes |
| multi | str | 262144 | valuesFor | ordered | hashed | 12 | 227 | 160 | 0.71× [0.69, 0.72] | -41.5% | [-44.7%, -38.3%] | 4.1 pts | 3.6 | yes | yes |
| multi | str | 262144 | valuesFor | ordered | map-sets | 12 | 243 | 370 | 1.53× [1.50, 1.56] | +34.5% | [+33.2%, +35.9%] | 2.0 pts | 4.0 | yes | yes |
| multi | str | 262144 | valuesBetween | ordered | btree-sets | 12 | 8382 | 23.7 µs | 2.82× [2.79, 2.85] | +64.5% | [+64.1%, +64.9%] | 0.6 pts | 1.9 | yes | yes |
| multi | str | 262144 | prefix | ordered | btree-sets | 12 | 1.32 ms | 3.99 ms | 3.01× [2.95, 3.07] | +66.7% | [+66.0%, +67.4%] | 0.7 pts | 2.0 | yes | yes |
| multi | str | 262144 | churn | ordered | btree-sets | 12 | 491 | 820 | 1.69× [1.67, 1.72] | +41.0% | [+39.9%, +42.0%] | 1.3 pts | 1.3 | yes | yes |
| multi | str | 262144 | churn | ordered | hashed | 12 | 422 | 312 | 0.71× [0.67, 0.75] | -40.9% | [-48.8%, -33.0%] | 7.5 pts | 0.8 | no | yes |
| multi | str | 262144 | churn | ordered | map-sets | 12 | 442 | 360 | 0.80× [0.79, 0.82] | -24.4% | [-26.1%, -22.7%] | 2.5 pts | 0.8 | yes | yes |
| multi | u64 | 4096 | valuesFor | ordered | btree-sets | 12 | 32.8 | 162 | 4.95× [4.89, 5.03] | +79.8% | [+79.5%, +80.1%] | 0.3 pts | 0.9 | yes | yes |
| multi | u64 | 4096 | valuesFor | ordered | hashed | 12 | 31.6 | 31.1 | 0.96× [0.94, 0.99] | -4.0% | [-6.9%, -1.0%] | 3.0 pts | 0.7 | no | yes |
| multi | u64 | 4096 | valuesFor | ordered | map-sets | 12 | 31.3 | 92.7 | 2.95× [2.88, 3.02] | +66.1% | [+65.3%, +66.9%] | 0.8 pts | 1.1 | yes | yes |
| multi | u64 | 4096 | valuesBetween | ordered | btree-sets | 12 | 2551 | 7073 | 2.78× [2.75, 2.82] | +64.1% | [+63.6%, +64.5%] | 0.5 pts | 0.7 | yes | yes |
| multi | u64 | 4096 | valuesBetween | ordered | hashed | 12 | 2690 | 51.1 µs | 19.07× [18.90, 19.25] | +94.8% | [+94.7%, +94.8%] | 0.1 pts | 1.3 | yes | yes |
| multi | u64 | 4096 | valuesBetween | ordered | map-sets | 12 | 2590 | 52.9 µs | 20.51× [20.39, 20.63] | +95.1% | [+95.1%, +95.2%] | 0.1 pts | 1.3 | yes | yes |
| multi | u64 | 4096 | churn | ordered | btree-sets | 12 | 49.2 | 157 | 3.18× [3.14, 3.22] | +68.6% | [+68.2%, +68.9%] | 0.4 pts | 1.3 | yes | yes |
| multi | u64 | 4096 | churn | ordered | hashed | 12 | 49.0 | 51.9 | 1.06× [1.06, 1.07] | +5.9% | [+5.4%, +6.4%] | 1.0 pts | 1.5 | yes | yes |
| multi | u64 | 4096 | churn | ordered | map-sets | 12 | 48.9 | 55.1 | 1.13× [1.12, 1.14] | +11.5% | [+10.7%, +12.2%] | 1.6 pts | 2.1 | yes | yes |
| multi | u64 | 4096 | build | ordered | btree-sets | 12 | 5.14 ms | 14.78 ms | 2.87× [2.85, 2.90] | +65.2% | [+64.9%, +65.5%] | 0.3 pts | 1.1 | yes | yes |
| multi | u64 | 4096 | build | ordered | hashed | 12 | 5.14 ms | 5.25 ms | 1.02× [1.01, 1.03] | +2.0% | [+1.5%, +2.6%] | 0.8 pts | 1.0 | yes | yes |
| multi | u64 | 4096 | build | ordered | map-sets | 12 | 5.13 ms | 5.57 ms | 1.08× [1.08, 1.09] | +7.8% | [+7.1%, +8.4%] | 0.9 pts | 1.0 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 46.0 | 194 | 4.20× [4.17, 4.23] | +76.2% | [+76.0%, +76.4%] | 0.2 pts | 1.1 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | hashed | 6 | 46.1 | 43.3 | 0.94× [0.93, 0.95] | -6.6% | [-7.4%, -5.8%] | 0.8 pts | 1.3 | yes | yes |
| multi | u64 | 16384 | valuesFor | ordered | map-sets | 6 | 45.5 | 99.9 | 2.20× [2.18, 2.21] | +54.5% | [+54.2%, +54.7%] | 0.2 pts | 1.0 | yes | yes |
| multi | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 3470 | 7487 | 2.16× [2.14, 2.18] | +53.7% | [+53.2%, +54.1%] | 0.4 pts | 1.8 | yes | yes |
| multi | u64 | 16384 | valuesBetween | ordered | hashed | 6 | 3629 | 202.9 µs | 55.51× [54.68, 56.35] | +98.2% | [+98.2%, +98.2%] | 0.0 pts | 1.4 | yes | yes |
| multi | u64 | 16384 | valuesBetween | ordered | map-sets | 6 | 3675 | 195.9 µs | 53.21× [52.60, 53.83] | +98.1% | [+98.1%, +98.1%] | 0.0 pts | 0.9 | yes | yes |
| multi | u64 | 16384 | churn | ordered | btree-sets | 6 | 65.1 | 227 | 3.45× [3.41, 3.49] | +71.0% | [+70.7%, +71.4%] | 0.3 pts | 1.0 | yes | yes |
| multi | u64 | 16384 | churn | ordered | hashed | 6 | 62.5 | 66.9 | 1.06× [1.05, 1.07] | +5.7% | [+4.8%, +6.7%] | 0.9 pts | 1.0 | yes | yes |
| multi | u64 | 16384 | churn | ordered | map-sets | 6 | 63.3 | 75.8 | 1.18× [1.15, 1.20] | +15.0% | [+13.4%, +16.7%] | 1.6 pts | 1.9 | yes | yes |
| multi | u64 | 16384 | build | ordered | btree-sets | 6 | 25.23 ms | 78.74 ms | 3.10× [3.07, 3.14] | +67.8% | [+67.4%, +68.1%] | 0.3 pts | 1.0 | yes | yes |
| multi | u64 | 16384 | build | ordered | hashed | 6 | 24.95 ms | 24.86 ms | 0.99× [0.98, 1.00] | -0.6% | [-1.6%, +0.3%] | 0.9 pts | 2.0 | yes | no |
| multi | u64 | 16384 | build | ordered | map-sets | 6 | 25.15 ms | 27.67 ms | 1.11× [1.09, 1.13] | +10.1% | [+8.6%, +11.6%] | 1.4 pts | 0.7 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | btree-sets | 12 | 148 | 526 | 3.56× [3.52, 3.61] | +71.9% | [+71.6%, +72.3%] | 0.5 pts | 2.1 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | hashed | 12 | 130 | 148 | 1.13× [1.11, 1.14] | +11.1% | [+9.7%, +12.6%] | 2.2 pts | 3.5 | yes | yes |
| multi | u64 | 262144 | valuesFor | ordered | map-sets | 12 | 137 | 350 | 2.56× [2.53, 2.59] | +60.9% | [+60.5%, +61.4%] | 0.8 pts | 2.9 | yes | yes |
| multi | u64 | 262144 | valuesBetween | ordered | btree-sets | 12 | 6959 | 22.1 µs | 3.18× [3.16, 3.20] | +68.6% | [+68.3%, +68.8%] | 0.5 pts | 1.7 | yes | yes |
| multi | u64 | 262144 | churn | ordered | btree-sets | 12 | 363 | 669 | 1.88× [1.82, 1.94] | +46.8% | [+45.2%, +48.4%] | 1.7 pts | 2.2 | yes | yes |
| multi | u64 | 262144 | churn | ordered | hashed | 12 | 296 | 264 | 0.89× [0.86, 0.93] | -12.1% | [-16.9%, -7.3%] | 5.8 pts | 0.6 | no | yes |
| multi | u64 | 262144 | churn | ordered | map-sets | 12 | 318 | 313 | 0.98× [0.98, 0.99] | -1.7% | [-1.9%, -1.4%] | 1.5 pts | 0.9 | yes | yes |
| multi | url | 4096 | valuesFor | ordered | btree-sets | 6 | 86.8 | 193 | 2.24× [2.21, 2.28] | +55.4% | [+54.7%, +56.2%] | 0.7 pts | 1.5 | yes | yes |
| multi | url | 4096 | valuesFor | ordered | hashed | 6 | 84.8 | 42.3 | 0.50× [0.49, 0.51] | -99.9% | [-104.2%, -95.5%] | 4.1 pts | 1.1 | yes | yes |
| multi | url | 4096 | valuesFor | ordered | map-sets | 6 | 85.3 | 97.7 | 1.15× [1.14, 1.16] | +13.1% | [+12.1%, +14.2%] | 1.0 pts | 1.6 | yes | yes |
| multi | url | 4096 | valuesBetween | ordered | btree-sets | 6 | 3726 | 7464 | 2.01× [1.98, 2.03] | +50.1% | [+49.4%, +50.8%] | 0.7 pts | 1.1 | yes | yes |
| multi | url | 4096 | valuesBetween | ordered | hashed | 6 | 3869 | 61.1 µs | 15.86× [15.64, 16.09] | +93.7% | [+93.6%, +93.8%] | 0.1 pts | 1.1 | yes | yes |
| multi | url | 4096 | valuesBetween | ordered | map-sets | 6 | 3809 | 62.8 µs | 16.41× [16.14, 16.69] | +93.9% | [+93.8%, +94.0%] | 0.1 pts | 1.3 | yes | yes |
| multi | url | 4096 | prefix | ordered | btree-sets | 6 | 180 | 249 | 1.38× [1.37, 1.39] | +27.6% | [+26.8%, +28.3%] | 0.7 pts | 0.9 | yes | yes |
| multi | url | 4096 | prefix | ordered | hashed | 6 | 227 | 53.1 µs | 239.24× [222.59, 258.59] | +99.6% | [+99.6%, +99.6%] | 0.0 pts | 1.1 | yes | yes |
| multi | url | 4096 | prefix | ordered | map-sets | 6 | 218 | 49.1 µs | 223.00× [205.76, 243.38] | +99.6% | [+99.5%, +99.6%] | 0.0 pts | 1.2 | yes | yes |
| multi | url | 4096 | churn | ordered | btree-sets | 6 | 124 | 207 | 1.68× [1.65, 1.70] | +40.4% | [+39.5%, +41.3%] | 0.9 pts | 1.3 | yes | yes |
| multi | url | 4096 | churn | ordered | hashed | 6 | 121 | 67.2 | 0.55× [0.55, 0.56] | -80.3% | [-82.5%, -78.0%] | 2.2 pts | 2.1 | yes | yes |
| multi | url | 4096 | churn | ordered | map-sets | 6 | 121 | 71.5 | 0.59× [0.58, 0.60] | -69.0% | [-71.2%, -66.7%] | 2.2 pts | 1.1 | yes | yes |
| multi | url | 4096 | build | ordered | btree-sets | 6 | 11.53 ms | 19.85 ms | 1.72× [1.71, 1.73] | +41.8% | [+41.5%, +42.1%] | 0.3 pts | 0.9 | yes | yes |
| multi | url | 4096 | build | ordered | hashed | 6 | 11.48 ms | 6.35 ms | 0.56× [0.55, 0.56] | -80.2% | [-82.7%, -77.6%] | 2.4 pts | 2.5 | yes | yes |
| multi | url | 4096 | build | ordered | map-sets | 6 | 11.48 ms | 6.97 ms | 0.61× [0.59, 0.64] | -63.1% | [-68.9%, -57.4%] | 5.5 pts | 3.5 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | btree-sets | 6 | 112 | 248 | 2.22× [2.17, 2.26] | +54.9% | [+53.9%, +55.8%] | 0.9 pts | 2.3 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | hashed | 6 | 109 | 51.1 | 0.47× [0.46, 0.47] | -113.7% | [-115.7%, -111.7%] | 1.9 pts | 2.0 | yes | yes |
| multi | url | 16384 | valuesFor | ordered | map-sets | 6 | 110 | 105 | 0.96× [0.95, 0.97] | -4.0% | [-5.1%, -2.8%] | 1.1 pts | 1.6 | yes | yes |
| multi | url | 16384 | valuesBetween | ordered | btree-sets | 6 | 4419 | 7967 | 1.81× [1.80, 1.82] | +44.7% | [+44.4%, +45.0%] | 0.3 pts | 0.7 | yes | yes |
| multi | url | 16384 | valuesBetween | ordered | hashed | 6 | 5059 | 259.7 µs | 50.83× [48.63, 53.22] | +98.0% | [+97.9%, +98.1%] | 0.1 pts | 1.4 | yes | yes |
| multi | url | 16384 | valuesBetween | ordered | map-sets | 6 | 5149 | 249.9 µs | 47.81× [46.30, 49.43] | +97.9% | [+97.8%, +98.0%] | 0.1 pts | 0.8 | yes | yes |
| multi | url | 16384 | prefix | ordered | btree-sets | 6 | 287 | 462 | 1.60× [1.58, 1.63] | +37.7% | [+36.8%, +38.6%] | 0.9 pts | 1.2 | yes | yes |
| multi | url | 16384 | prefix | ordered | hashed | 6 | 574 | 250.5 µs | 439.02× [429.43, 449.05] | +99.8% | [+99.8%, +99.8%] | 0.0 pts | 0.4 | yes | yes |
| multi | url | 16384 | prefix | ordered | map-sets | 6 | 524 | 233.2 µs | 438.35× [421.12, 457.05] | +99.8% | [+99.8%, +99.8%] | 0.0 pts | 0.8 | yes | yes |
| multi | url | 16384 | churn | ordered | btree-sets | 6 | 184 | 320 | 1.75× [1.72, 1.78] | +42.8% | [+41.9%, +43.7%] | 0.9 pts | 1.6 | yes | yes |
| multi | url | 16384 | churn | ordered | hashed | 6 | 170 | 88.4 | 0.52× [0.52, 0.53] | -90.8% | [-92.6%, -89.0%] | 1.7 pts | 1.2 | yes | yes |
| multi | url | 16384 | churn | ordered | map-sets | 6 | 175 | 109 | 0.64× [0.62, 0.66] | -56.4% | [-60.9%, -52.0%] | 4.2 pts | 2.0 | yes | yes |
| multi | url | 16384 | build | ordered | btree-sets | 6 | 58.91 ms | 105.95 ms | 1.80× [1.78, 1.81] | +44.3% | [+43.8%, +44.9%] | 0.5 pts | 1.5 | yes | yes |
| multi | url | 16384 | build | ordered | hashed | 6 | 59.05 ms | 31.58 ms | 0.54× [0.53, 0.54] | -86.3% | [-87.4%, -85.3%] | 1.0 pts | 0.9 | yes | yes |
| multi | url | 16384 | build | ordered | map-sets | 6 | 58.94 ms | 36.51 ms | 0.62× [0.61, 0.63] | -60.6% | [-63.0%, -58.1%] | 2.3 pts | 1.3 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | btree-sets | 12 | 420 | 740 | 1.77× [1.75, 1.79] | +43.4% | [+42.8%, +44.1%] | 1.3 pts | 2.1 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | hashed | 12 | 349 | 187 | 0.54× [0.53, 0.54] | -86.3% | [-89.0%, -83.6%] | 3.1 pts | 2.6 | yes | yes |
| multi | url | 262144 | valuesFor | ordered | map-sets | 12 | 361 | 389 | 1.08× [1.06, 1.11] | +7.6% | [+5.5%, +9.7%] | 2.1 pts | 3.6 | no | yes |
| multi | url | 262144 | valuesBetween | ordered | btree-sets | 12 | 10.9 µs | 24.5 µs | 2.25× [2.22, 2.27] | +55.5% | [+55.0%, +56.0%] | 0.7 pts | 2.2 | yes | yes |
| multi | url | 262144 | prefix | ordered | btree-sets | 12 | 2588 | 5385 | 2.08× [2.04, 2.11] | +51.8% | [+50.9%, +52.7%] | 1.2 pts | 0.7 | yes | yes |
| multi | url | 262144 | churn | ordered | btree-sets | 12 | 714 | 957 | 1.37× [1.34, 1.41] | +27.1% | [+25.1%, +29.1%] | 2.2 pts | 1.6 | yes | yes |
| multi | url | 262144 | churn | ordered | hashed | 12 | 621 | 340 | 0.53× [0.52, 0.55] | -87.9% | [-93.4%, -82.4%] | 9.9 pts | 0.7 | yes | yes |
| multi | url | 262144 | churn | ordered | map-sets | 12 | 646 | 409 | 0.63× [0.63, 0.64] | -57.7% | [-59.9%, -55.6%] | 4.0 pts | 1.1 | yes | yes |
| multi | uuid | 4096 | valuesFor | ordered | btree-sets | 6 | 53.3 | 177 | 3.35× [3.32, 3.38] | +70.1% | [+69.9%, +70.4%] | 0.3 pts | 1.1 | yes | yes |
| multi | uuid | 4096 | valuesFor | ordered | hashed | 6 | 51.6 | 36.0 | 0.70× [0.68, 0.72] | -42.8% | [-46.0%, -39.6%] | 3.1 pts | 0.7 | yes | yes |
| multi | uuid | 4096 | valuesFor | ordered | map-sets | 6 | 52.1 | 95.4 | 1.82× [1.81, 1.84] | +45.2% | [+44.8%, +45.5%] | 0.4 pts | 0.8 | yes | yes |
| multi | uuid | 4096 | valuesBetween | ordered | btree-sets | 6 | 3078 | 7342 | 2.41× [2.38, 2.44] | +58.4% | [+57.9%, +59.0%] | 0.5 pts | 0.8 | yes | yes |
| multi | uuid | 4096 | valuesBetween | ordered | hashed | 6 | 3269 | 52.7 µs | 16.23× [16.17, 16.28] | +93.8% | [+93.8%, +93.9%] | 0.0 pts | 0.2 | yes | yes |
| multi | uuid | 4096 | valuesBetween | ordered | map-sets | 6 | 3143 | 54.4 µs | 17.46× [17.11, 17.82] | +94.3% | [+94.2%, +94.4%] | 0.1 pts | 1.2 | yes | yes |
| multi | uuid | 4096 | prefix | ordered | btree-sets | 6 | 89.4 | 180 | 2.01× [2.00, 2.03] | +50.4% | [+50.1%, +50.6%] | 0.3 pts | 0.6 | yes | yes |
| multi | uuid | 4096 | prefix | ordered | hashed | 6 | 109 | 47.3 µs | 437.82× [429.78, 446.18] | +99.8% | [+99.8%, +99.8%] | 0.0 pts | 0.4 | yes | yes |
| multi | uuid | 4096 | prefix | ordered | map-sets | 6 | 109 | 43.8 µs | 400.56× [397.31, 403.87] | +99.8% | [+99.7%, +99.8%] | 0.0 pts | 0.2 | yes | yes |
| multi | uuid | 4096 | churn | ordered | btree-sets | 6 | 78.0 | 179 | 2.28× [2.24, 2.32] | +56.1% | [+55.4%, +56.9%] | 0.7 pts | 1.5 | yes | yes |
| multi | uuid | 4096 | churn | ordered | hashed | 6 | 77.7 | 59.0 | 0.76× [0.76, 0.77] | -31.2% | [-32.2%, -30.3%] | 0.9 pts | 0.9 | yes | yes |
| multi | uuid | 4096 | churn | ordered | map-sets | 6 | 76.6 | 63.3 | 0.83× [0.82, 0.84] | -20.6% | [-21.7%, -19.4%] | 1.1 pts | 1.1 | yes | yes |
| multi | uuid | 4096 | build | ordered | btree-sets | 6 | 7.60 ms | 16.73 ms | 2.19× [2.15, 2.23] | +54.3% | [+53.5%, +55.2%] | 0.8 pts | 2.3 | yes | yes |
| multi | uuid | 4096 | build | ordered | hashed | 6 | 7.65 ms | 5.81 ms | 0.76× [0.76, 0.76] | -31.7% | [-32.4%, -31.1%] | 0.7 pts | 0.6 | yes | yes |
| multi | uuid | 4096 | build | ordered | map-sets | 6 | 7.59 ms | 6.26 ms | 0.82× [0.82, 0.83] | -21.2% | [-22.1%, -20.4%] | 0.8 pts | 0.8 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | btree-sets | 12 | 65.1 | 219 | 3.37× [3.35, 3.38] | +70.3% | [+70.2%, +70.4%] | 0.3 pts | 1.6 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | hashed | 12 | 64.6 | 46.4 | 0.72× [0.72, 0.72] | -38.8% | [-39.5%, -38.0%] | 1.6 pts | 1.7 | yes | yes |
| multi | uuid | 16384 | valuesFor | ordered | map-sets | 12 | 63.9 | 102 | 1.59× [1.58, 1.60] | +37.1% | [+36.8%, +37.5%] | 0.5 pts | 1.3 | yes | yes |
| multi | uuid | 16384 | valuesBetween | ordered | btree-sets | 12 | 3578 | 7683 | 2.15× [2.14, 2.16] | +53.5% | [+53.2%, +53.8%] | 0.3 pts | 1.4 | yes | yes |
| multi | uuid | 16384 | valuesBetween | ordered | hashed | 12 | 3947 | 231.0 µs | 58.16× [56.58, 59.82] | +98.3% | [+98.2%, +98.3%] | 0.1 pts | 1.4 | yes | yes |
| multi | uuid | 16384 | valuesBetween | ordered | map-sets | 12 | 4108 | 221.3 µs | 53.68× [52.10, 55.36] | +98.1% | [+98.1%, +98.2%] | 0.1 pts | 1.5 | yes | yes |
| multi | uuid | 16384 | prefix | ordered | btree-sets | 12 | 107 | 244 | 2.27× [2.26, 2.27] | +55.9% | [+55.8%, +56.0%] | 0.3 pts | 1.4 | yes | yes |
| multi | uuid | 16384 | prefix | ordered | hashed | 12 | 199 | 220.9 µs | 1099.00× [1075.56, 1123.48] | +99.9% | [+99.9%, +99.9%] | 0.0 pts | 0.3 | yes | yes |
| multi | uuid | 16384 | prefix | ordered | map-sets | 12 | 202 | 204.5 µs | 1029.98× [1009.26, 1051.56] | +99.9% | [+99.9%, +99.9%] | 0.0 pts | 0.6 | yes | yes |
| multi | uuid | 16384 | churn | ordered | btree-sets | 12 | 108 | 262 | 2.42× [2.40, 2.44] | +58.7% | [+58.4%, +59.0%] | 0.5 pts | 1.1 | yes | yes |
| multi | uuid | 16384 | churn | ordered | hashed | 12 | 101 | 77.4 | 0.76× [0.75, 0.77] | -31.4% | [-33.6%, -29.2%] | 2.2 pts | 1.5 | yes | yes |
| multi | uuid | 16384 | churn | ordered | map-sets | 12 | 102 | 91.5 | 0.91× [0.88, 0.94] | -10.2% | [-14.2%, -6.3%] | 3.8 pts | 2.9 | no | yes |
| multi | uuid | 16384 | build | ordered | btree-sets | 12 | 36.57 ms | 88.97 ms | 2.43× [2.42, 2.45] | +58.9% | [+58.7%, +59.2%] | 0.5 pts | 1.6 | yes | yes |
| multi | uuid | 16384 | build | ordered | hashed | 12 | 36.23 ms | 27.15 ms | 0.75× [0.75, 0.76] | -33.1% | [-33.9%, -32.2%] | 1.0 pts | 1.3 | yes | yes |
| multi | uuid | 16384 | build | ordered | map-sets | 12 | 36.52 ms | 31.93 ms | 0.87× [0.85, 0.89] | -15.0% | [-17.3%, -12.7%] | 2.6 pts | 1.7 | no | yes |
| multi | uuid | 262144 | valuesFor | ordered | btree-sets | 12 | 271 | 630 | 2.38× [2.32, 2.44] | +57.9% | [+56.8%, +59.0%] | 1.5 pts | 3.8 | yes | yes |
| multi | uuid | 262144 | valuesFor | ordered | hashed | 12 | 225 | 169 | 0.75× [0.74, 0.77] | -32.6% | [-35.0%, -30.2%] | 3.8 pts | 4.3 | yes | yes |
| multi | uuid | 262144 | valuesFor | ordered | map-sets | 12 | 237 | 378 | 1.60× [1.56, 1.64] | +37.4% | [+35.9%, +38.8%] | 1.7 pts | 4.1 | yes | yes |
| multi | uuid | 262144 | valuesBetween | ordered | btree-sets | 12 | 8603 | 24.0 µs | 2.78× [2.75, 2.82] | +64.1% | [+63.6%, +64.5%] | 0.7 pts | 2.7 | yes | yes |
| multi | uuid | 262144 | prefix | ordered | btree-sets | 12 | 599 | 1715 | 2.88× [2.83, 2.94] | +65.3% | [+64.7%, +66.0%] | 0.7 pts | 2.2 | yes | yes |
| multi | uuid | 262144 | churn | ordered | btree-sets | 12 | 512 | 856 | 1.71× [1.67, 1.75] | +41.5% | [+40.3%, +42.8%] | 1.7 pts | 2.3 | yes | yes |
| multi | uuid | 262144 | churn | ordered | hashed | 12 | 437 | 326 | 0.72× [0.69, 0.75] | -39.2% | [-45.1%, -33.3%] | 8.8 pts | 0.6 | no | yes |
| multi | uuid | 262144 | churn | ordered | map-sets | 12 | 458 | 384 | 0.84× [0.83, 0.85] | -18.9% | [-20.7%, -17.2%] | 2.8 pts | 0.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi str n=16384 valuesFor: ordered vs map-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of -0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi str n=16384 prefix: ordered vs hashed: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +0.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi str n=16384 churn: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesFor: ordered vs btree-sets: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesFor: ordered vs hashed: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesFor: ordered vs map-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 churn: ordered vs hashed: the A/A validations found a systematic difference of +2.04% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi u64 n=4096 churn: ordered vs map-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi u64 n=16384 build: ordered vs hashed: the pooled interval [-1.63%, 0.34%] includes zero
- multi u64 n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs hashed: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs map-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=4096 prefix: ordered vs hashed: the A/A validations found a systematic difference of +4.28% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=4096 churn: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=4096 build: ordered vs hashed: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=4096 build: ordered vs map-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +1.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +5.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=16384 churn: ordered vs hashed: the A/A validations found a systematic difference of -0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=16384 churn: ordered vs map-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=16384 build: ordered vs map-sets: the A/A validations found a systematic difference of -0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi url n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs hashed: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs map-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=4096 prefix: ordered vs hashed: the A/A validations found a systematic difference of +4.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi uuid n=4096 build: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +11.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi uuid n=16384 prefix: ordered vs map-sets: the A/A validations found a systematic difference of +9.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi uuid n=16384 churn: ordered vs map-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi uuid n=262144 valuesFor: ordered vs btree-sets: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesFor: ordered vs hashed: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesFor: ordered vs map-sets: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 prefix: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
