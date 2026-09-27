| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | baseline | 6 | 66.6 | 49.8 | 0.75× [0.74, 0.76] | -33.9% | [-35.8%, -32.0%] | 1.8 pts | 1.4 | yes |
| multi | email | 4096 | valuesBetween | ordered | baseline | 6 | 3080 | 2916 | 0.95× [0.93, 0.96] | -5.6% | [-7.7%, -4.2%] | 1.7 pts | 0.8 | yes |
| multi | email | 4096 | prefix | ordered | baseline | 6 | 89.5 | 75.5 | 0.84× [0.84, 0.85] | -18.8% | [-19.1%, -17.6%] | 0.8 pts | 0.6 | yes |
| multi | email | 4096 | churn | ordered | baseline | 6 | 87.6 | 58.0 | 0.66× [0.65, 0.67] | -51.9% | [-54.0%, -49.9%] | 1.9 pts | 0.5 | yes |
| multi | email | 4096 | build | ordered | baseline | 6 | 8.76 ms | 4.91 ms | 0.56× [0.55, 0.56] | -78.6% | [-80.3%, -77.8%] | 1.2 pts | 0.7 | yes |
| multi | email | 16384 | valuesFor | ordered | baseline | 6 | 82.2 | 60.4 | 0.73× [0.70, 0.75] | -37.3% | [-42.4%, -33.2%] | 4.4 pts | 2.0 | no |
| multi | email | 16384 | valuesBetween | ordered | baseline | 6 | 3822 | 3509 | 0.91× [0.90, 0.93] | -9.5% | [-10.8%, -7.1%] | 1.7 pts | 1.3 | yes |
| multi | email | 16384 | prefix | ordered | baseline | 6 | 116 | 93.6 | 0.82× [0.79, 0.84] | -22.4% | [-27.0%, -19.3%] | 3.7 pts | 1.5 | no |
| multi | email | 16384 | churn | ordered | baseline | 6 | 140 | 85.5 | 0.61× [0.58, 0.64] | -64.1% | [-72.1%, -56.5%] | 7.4 pts | 1.0 | no |
| multi | email | 16384 | build | ordered | baseline | 6 | 50.67 ms | 25.74 ms | 0.51× [0.51, 0.52] | -95.4% | [-97.1%, -92.8%] | 2.1 pts | 0.9 | yes |
| multi | email | 65536 | valuesFor | ordered | baseline | 6 | 229 | 128 | 0.63× [0.47, 0.77] | -58.0% | [-112.0%, -30.1%] | 39.0 pts | 2.6 | no |
| multi | email | 65536 | valuesBetween | ordered | baseline | 6 | 6954 | 4935 | 0.73× [0.65, 0.83] | -36.4% | [-53.8%, -20.1%] | 16.1 pts | 2.2 | no |
| multi | email | 65536 | prefix | ordered | baseline | 6 | 275 | 166 | 0.60× [0.55, 0.70] | -66.1% | [-80.3%, -43.0%] | 17.8 pts | 2.2 | no |
| multi | email | 65536 | churn | ordered | baseline | 6 | 342 | 195 | 0.57× [0.54, 0.62] | -75.1% | [-86.7%, -62.3%] | 11.6 pts | 2.1 | no |
| multi | email | 65536 | build | ordered | baseline | 6 | 342.53 ms | 191.43 ms | 0.56× [0.54, 0.59] | -77.8% | [-84.6%, -69.2%] | 7.3 pts | 3.5 | no |
| multi | email | 262144 | valuesFor | ordered | baseline | 6 | 436 | 258 | 0.62× [0.54, 0.69] | -60.6% | [-83.8%, -44.4%] | 18.8 pts | 6.6 | no |
| multi | email | 262144 | valuesBetween | ordered | baseline | 6 | 10.7 µs | 9279 | 0.86× [0.84, 0.91] | -15.8% | [-19.0%, -10.2%] | 4.2 pts | 2.2 | no |
| multi | email | 262144 | prefix | ordered | baseline | 6 | 487 | 347 | 0.71× [0.68, 0.78] | -40.7% | [-46.6%, -28.1%] | 8.8 pts | 1.9 | no |
| multi | email | 262144 | churn | ordered | baseline | 6 | 604 | 357 | 0.57× [0.54, 0.63] | -75.6% | [-86.0%, -58.3%] | 13.2 pts | 3.3 | no |
| multi | email | 1048576 | valuesFor | ordered | baseline | 6 | 592 | 410 | 0.71× [0.68, 0.74] | -40.6% | [-47.4%, -35.6%] | 5.6 pts | 2.2 | no |
| multi | email | 1048576 | valuesBetween | ordered | baseline | 6 | 13.5 µs | 11.4 µs | 0.85× [0.83, 0.86] | -18.3% | [-21.2%, -16.6%] | 2.2 pts | 1.5 | no |
| multi | email | 1048576 | prefix | ordered | baseline | 6 | 777 | 648 | 0.84× [0.79, 0.87] | -18.5% | [-26.6%, -14.4%] | 5.8 pts | 2.9 | no |
| multi | email | 1048576 | churn | ordered | baseline | 6 | 923 | 530 | 0.56× [0.54, 0.61] | -77.8% | [-84.5%, -64.6%] | 9.5 pts | 3.4 | no |
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 120 | 104 | 0.87× [0.86, 0.88] | -15.1% | [-16.8%, -14.2%] | 1.2 pts | 1.3 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4280 | 4020 | 0.93× [0.93, 0.95] | -7.0% | [-7.7%, -5.6%] | 1.0 pts | 0.4 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 367 | 320 | 0.87× [0.86, 0.89] | -14.7% | [-16.6%, -12.4%] | 2.0 pts | 0.6 | no |
| multi | path | 4096 | churn | ordered | baseline | 6 | 189 | 151 | 0.80× [0.78, 0.82] | -25.2% | [-28.5%, -22.1%] | 3.0 pts | 1.0 | no |
| multi | path | 4096 | build | ordered | baseline | 6 | 16.68 ms | 11.21 ms | 0.67× [0.66, 0.68] | -49.3% | [-50.5%, -47.9%] | 1.3 pts | 1.1 | yes |
| multi | path | 16384 | valuesFor | ordered | baseline | 6 | 163 | 137 | 0.84× [0.75, 0.89] | -18.9% | [-33.3%, -12.2%] | 10.1 pts | 4.5 | no |
| multi | path | 16384 | valuesBetween | ordered | baseline | 6 | 5255 | 4686 | 0.89× [0.88, 0.91] | -12.4% | [-13.8%, -9.9%] | 1.8 pts | 0.7 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 6 | 738 | 624 | 0.84× [0.82, 0.86] | -18.7% | [-22.4%, -16.0%] | 3.0 pts | 0.5 | no |
| multi | path | 16384 | churn | ordered | baseline | 6 | 361 | 243 | 0.69× [0.66, 0.74] | -44.4% | [-51.9%, -35.3%] | 8.0 pts | 1.0 | no |
| multi | path | 16384 | build | ordered | baseline | 6 | 96.26 ms | 63.71 ms | 0.66× [0.65, 0.67] | -51.3% | [-53.0%, -49.8%] | 1.5 pts | 0.9 | yes |
| multi | path | 65536 | valuesFor | ordered | baseline | 6 | 451 | 315 | 0.67× [0.63, 0.71] | -49.1% | [-58.9%, -41.7%] | 8.2 pts | 1.7 | no |
| multi | path | 65536 | valuesBetween | ordered | baseline | 6 | 10.1 µs | 7358 | 0.72× [0.63, 0.80] | -38.4% | [-59.1%, -25.4%] | 16.0 pts | 2.4 | no |
| multi | path | 65536 | prefix | ordered | baseline | 6 | 2745 | 2264 | 0.80× [0.77, 0.87] | -24.5% | [-30.6%, -14.3%] | 7.8 pts | 0.5 | no |
| multi | path | 65536 | churn | ordered | baseline | 6 | 702 | 446 | 0.63× [0.57, 0.70] | -59.1% | [-74.1%, -43.6%] | 14.5 pts | 2.8 | no |
| multi | path | 65536 | build | ordered | baseline | 6 | 708.93 ms | 488.22 ms | 0.69× [0.68, 0.69] | -45.7% | [-47.8%, -44.4%] | 1.7 pts | 1.5 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 6 | 751 | 568 | 0.74× [0.71, 0.79] | -35.2% | [-40.7%, -26.5%] | 6.8 pts | 2.7 | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 6 | 14.7 µs | 12.6 µs | 0.86× [0.84, 0.89] | -16.3% | [-18.7%, -12.9%] | 2.8 pts | 1.4 | no |
| multi | path | 262144 | prefix | ordered | baseline | 6 | 12.4 µs | 11.1 µs | 0.89× [0.85, 0.92] | -12.0% | [-17.8%, -8.1%] | 4.6 pts | 0.3 | no |
| multi | path | 262144 | churn | ordered | baseline | 6 | 1216 | 814 | 0.68× [0.63, 0.71] | -47.7% | [-58.5%, -41.0%] | 8.4 pts | 2.5 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 9 | 76.9 | 62.7 | 0.82× [0.80, 0.83] | -21.7% | [-24.5%, -20.8%] | 2.4 pts | 2.2 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 9 | 3594 | 3483 | 0.97× [0.95, 0.98] | -3.2% | [-5.3%, -1.9%] | 2.2 pts | 1.0 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 9 | 7645 | 7907 | 1.02× [1.00, 1.04] | +1.6% | [-0.1%, +3.5%] | 2.3 pts | 0.8 | yes |
| multi | str | 4096 | churn | ordered | baseline | 9 | 107 | 76.2 | 0.72× [0.72, 0.73] | -38.6% | [-39.6%, -37.2%] | 1.6 pts | 0.5 | yes |
| multi | str | 4096 | build | ordered | baseline | 9 | 10.19 ms | 6.05 ms | 0.59× [0.59, 0.60] | -68.4% | [-69.3%, -67.2%] | 1.3 pts | 1.1 | yes |
| multi | str | 16384 | valuesFor | ordered | baseline | 9 | 91.5 | 74.6 | 0.82× [0.80, 0.83] | -22.4% | [-25.1%, -21.0%] | 2.7 pts | 1.3 | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 9 | 4077 | 3713 | 0.90× [0.89, 0.91] | -10.6% | [-12.1%, -9.4%] | 1.7 pts | 1.4 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 9 | 38.1 µs | 34.5 µs | 0.91× [0.90, 0.91] | -10.5% | [-11.5%, -9.3%] | 1.4 pts | 0.8 | yes |
| multi | str | 16384 | churn | ordered | baseline | 9 | 151 | 103 | 0.70× [0.67, 0.71] | -43.9% | [-48.3%, -41.8%] | 4.2 pts | 0.9 | yes |
| multi | str | 16384 | build | ordered | baseline | 9 | 55.53 ms | 31.24 ms | 0.57× [0.56, 0.57] | -76.9% | [-78.5%, -75.0%] | 2.3 pts | 1.2 | yes |
| multi | str | 65536 | valuesFor | ordered | baseline | 10 | 258 | 169 | 0.65× [0.60, 0.73] | -54.3% | [-67.1%, -37.8%] | 20.4 pts | 2.0 | no |
| multi | str | 65536 | valuesBetween | ordered | baseline | 10 | 7193 | 5048 | 0.70× [0.66, 0.77] | -42.1% | [-50.4%, -29.9%] | 14.3 pts | 2.0 | no |
| multi | str | 65536 | prefix | ordered | baseline | 10 | 270.5 µs | 195.8 µs | 0.76× [0.68, 0.82] | -31.5% | [-47.0%, -21.2%] | 18.0 pts | 1.7 | no |
| multi | str | 65536 | churn | ordered | baseline | 10 | 368 | 225 | 0.59× [0.55, 0.65] | -68.4% | [-81.3%, -54.7%] | 18.6 pts | 4.1 | no |
| multi | str | 65536 | build | ordered | baseline | 10 | 359.93 ms | 216.22 ms | 0.60× [0.59, 0.61] | -67.0% | [-69.4%, -64.1%] | 3.7 pts | 2.1 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 442 | 329 | 0.75× [0.70, 0.77] | -33.8% | [-43.0%, -30.1%] | 9.0 pts | 2.7 | no |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 12.2 µs | 10.2 µs | 0.84× [0.83, 0.87] | -18.9% | [-21.1%, -15.3%] | 4.0 pts | 2.3 | no |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.96 ms | 1.66 ms | 0.84× [0.83, 0.86] | -18.6% | [-20.3%, -16.7%] | 2.6 pts | 1.5 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 712 | 417 | 0.59× [0.55, 0.63] | -70.0% | [-82.0%, -58.9%] | 16.2 pts | 3.9 | no |
| multi | str | 262144 | build | ordered | baseline | 10 | 2751.45 ms | 1860.80 ms | 0.68× [0.67, 0.68] | -47.6% | [-49.6%, -46.7%] | 2.0 pts | 3.8 | yes |
| multi | str | 1048576 | valuesFor | ordered | baseline | 10 | 548 | 447 | 0.83× [0.80, 0.85] | -20.6% | [-25.4%, -17.2%] | 5.7 pts | 4.0 | no |
| multi | str | 1048576 | valuesBetween | ordered | baseline | 10 | 13.0 µs | 11.6 µs | 0.89× [0.88, 0.90] | -12.2% | [-13.7%, -11.2%] | 1.8 pts | 1.1 | yes |
| multi | str | 1048576 | prefix | ordered | baseline | 10 | 7.77 ms | 7.21 ms | 0.93× [0.92, 0.93] | -7.9% | [-8.3%, -7.7%] | 0.4 pts | 0.9 | yes |
| multi | str | 1048576 | churn | ordered | baseline | 10 | 854 | 530 | 0.62× [0.58, 0.65] | -62.2% | [-72.8%, -53.9%] | 13.2 pts | 4.6 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 5 | 68.8 | 54.3 | 0.79× [0.79, 0.81] | -26.4% | [-27.3%, -24.2%] | 1.3 pts | 0.9 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 5 | 1818 | 2304 | 1.26× [1.24, 1.28] | +20.3% | [+19.6%, +22.0%] | 1.0 pts | 0.6 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 5 | 203 | 249 | 1.22× [1.21, 1.24] | +18.3% | [+17.1%, +19.2%] | 0.8 pts | 0.5 | yes |
| multi | street | 4096 | churn | ordered | baseline | 5 | 121 | 84.6 | 0.70× [0.69, 0.71] | -43.3% | [-44.2%, -40.6%] | 1.5 pts | 0.6 | yes |
| multi | street | 4096 | build | ordered | baseline | 5 | 5.28 ms | 2.62 ms | 0.50× [0.49, 0.50] | -102.0% | [-106.0%, -99.4%] | 2.7 pts | 1.2 | yes |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 90.3 | 76.1 | 0.84× [0.83, 0.85] | -18.9% | [-19.8%, -17.7%] | 1.0 pts | 0.6 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2311 | 2807 | 1.21× [1.19, 1.24] | +17.1% | [+15.6%, +19.5%] | 1.8 pts | 1.5 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 679 | 894 | 1.32× [1.29, 1.37] | +24.5% | [+22.6%, +26.9%] | 2.1 pts | 0.9 | yes |
| multi | street | 16384 | churn | ordered | baseline | 6 | 181 | 122 | 0.69× [0.67, 0.70] | -45.6% | [-48.9%, -42.9%] | 2.9 pts | 0.7 | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 28.08 ms | 14.72 ms | 0.52× [0.52, 0.53] | -90.8% | [-92.6%, -89.6%] | 1.4 pts | 1.1 | yes |
| multi | street | 65536 | valuesFor | ordered | baseline | 6 | 179 | 115 | 0.66× [0.59, 0.72] | -52.1% | [-69.3%, -38.8%] | 14.5 pts | 1.3 | no |
| multi | street | 65536 | valuesBetween | ordered | baseline | 6 | 3798 | 3312 | 0.96× [0.80, 1.03] | -4.5% | [-24.8%, +3.0%] | 13.3 pts | 1.7 | no |
| multi | street | 65536 | prefix | ordered | baseline | 6 | 3927 | 4546 | 1.16× [1.11, 1.22] | +13.5% | [+9.6%, +18.0%] | 4.0 pts | 1.1 | no |
| multi | street | 65536 | churn | ordered | baseline | 6 | 408 | 215 | 0.54× [0.52, 0.57] | -86.2% | [-94.1%, -75.7%] | 8.8 pts | 1.5 | no |
| multi | street | 65536 | build | ordered | baseline | 6 | 188.25 ms | 109.57 ms | 0.58× [0.58, 0.59] | -71.2% | [-72.5%, -69.4%] | 1.5 pts | 0.9 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 5 | 44.3 | 36.9 | 0.82× [0.81, 0.84] | -21.4% | [-23.1%, -19.1%] | 1.6 pts | 1.2 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 5 | 2670 | 2673 | 1.00× [0.99, 1.01] | +0.2% | [-0.5%, +1.3%] | 0.8 pts | 0.4 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 5 | 52.0 | 40.0 | 0.77× [0.76, 0.79] | -29.8% | [-31.4%, -26.9%] | 1.8 pts | 1.1 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 5 | 5.26 ms | 3.61 ms | 0.69× [0.68, 0.69] | -45.6% | [-47.7%, -44.6%] | 1.2 pts | 0.7 | yes |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 5 | 52.5 | 44.4 | 0.85× [0.84, 0.85] | -18.3% | [-19.2%, -17.2%] | 0.8 pts | 1.1 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 5 | 3182 | 3552 | 1.11× [1.11, 1.12] | +10.1% | [+9.8%, +10.8%] | 0.4 pts | 0.4 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 5 | 75.2 | 46.9 | 0.62× [0.62, 0.64] | -60.2% | [-61.8%, -55.3%] | 2.6 pts | 0.8 | yes |
| multi | u64 | 16384 | build | ordered | baseline | 5 | 26.96 ms | 16.89 ms | 0.63× [0.62, 0.64] | -58.3% | [-61.7%, -56.0%] | 2.3 pts | 1.7 | yes |
| multi | u64 | 65536 | valuesFor | ordered | baseline | 10 | 93.0 | 68.1 | 0.75× [0.65, 0.78] | -34.0% | [-54.9%, -27.7%] | 19.0 pts | 2.1 | no |
| multi | u64 | 65536 | valuesBetween | ordered | baseline | 10 | 5972 | 4708 | 0.80× [0.76, 0.87] | -25.1% | [-31.6%, -15.6%] | 11.2 pts | 1.7 | no |
| multi | u64 | 65536 | churn | ordered | baseline | 10 | 188 | 109 | 0.58× [0.54, 0.60] | -72.1% | [-83.6%, -66.4%] | 12.0 pts | 1.9 | no |
| multi | u64 | 65536 | build | ordered | baseline | 10 | 184.54 ms | 110.78 ms | 0.60× [0.59, 0.62] | -66.4% | [-68.9%, -62.1%] | 4.8 pts | 3.5 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 240 | 146 | 0.60× [0.57, 0.65] | -66.3% | [-76.6%, -54.5%] | 15.4 pts | 5.7 | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 9725 | 7987 | 0.81× [0.78, 0.83] | -23.5% | [-27.7%, -20.4%] | 5.1 pts | 2.4 | no |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 349 | 229 | 0.66× [0.60, 0.71] | -51.2% | [-65.6%, -40.0%] | 17.9 pts | 3.5 | no |
| multi | u64 | 262144 | build | ordered | baseline | 10 | 1512.55 ms | 1092.37 ms | 0.72× [0.72, 0.72] | -39.2% | [-39.6%, -38.2%] | 1.0 pts | 1.1 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | baseline | 10 | 285 | 197 | 0.69× [0.67, 0.69] | -45.8% | [-48.3%, -43.9%] | 3.1 pts | 1.2 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | baseline | 10 | 8726 | 8043 | 0.93× [0.92, 0.94] | -7.1% | [-8.7%, -6.2%] | 1.8 pts | 1.5 | yes |
| multi | u64 | 1048576 | churn | ordered | baseline | 10 | 562 | 337 | 0.64× [0.54, 0.69] | -56.3% | [-86.8%, -44.4%] | 29.7 pts | 7.8 | no |
| multi | url | 4096 | valuesFor | ordered | baseline | 6 | 93.4 | 81.8 | 0.87× [0.86, 0.89] | -15.0% | [-16.8%, -12.9%] | 1.8 pts | 2.9 | yes |
| multi | url | 4096 | valuesBetween | ordered | baseline | 6 | 3820 | 3618 | 0.95× [0.94, 0.96] | -5.6% | [-6.6%, -4.5%] | 1.0 pts | 0.6 | yes |
| multi | url | 4096 | prefix | ordered | baseline | 6 | 449 | 402 | 0.90× [0.88, 0.91] | -11.2% | [-13.5%, -9.4%] | 1.9 pts | 1.0 | no |
| multi | url | 4096 | churn | ordered | baseline | 6 | 144 | 111 | 0.77× [0.76, 0.79] | -29.7% | [-31.8%, -27.3%] | 2.1 pts | 0.9 | yes |
| multi | url | 4096 | build | ordered | baseline | 6 | 13.90 ms | 8.75 ms | 0.63× [0.62, 0.64] | -58.8% | [-60.3%, -57.4%] | 1.4 pts | 1.0 | yes |
| multi | url | 16384 | valuesFor | ordered | baseline | 6 | 119 | 97.3 | 0.82× [0.73, 0.88] | -22.2% | [-37.9%, -14.3%] | 11.2 pts | 3.6 | no |
| multi | url | 16384 | valuesBetween | ordered | baseline | 6 | 4392 | 3956 | 0.90× [0.88, 0.92] | -10.6% | [-13.9%, -8.2%] | 2.7 pts | 1.9 | no |
| multi | url | 16384 | prefix | ordered | baseline | 6 | 1506 | 1326 | 0.88× [0.87, 0.90] | -13.7% | [-14.6%, -11.3%] | 1.6 pts | 0.5 | yes |
| multi | url | 16384 | churn | ordered | baseline | 6 | 246 | 161 | 0.66× [0.62, 0.68] | -50.9% | [-60.4%, -46.6%] | 6.6 pts | 1.0 | no |
| multi | url | 16384 | build | ordered | baseline | 6 | 76.96 ms | 45.97 ms | 0.60× [0.58, 0.60] | -67.7% | [-71.1%, -65.8%] | 2.5 pts | 1.8 | yes |
| multi | url | 65536 | valuesFor | ordered | baseline | 6 | 370 | 249 | 0.69× [0.52, 0.84] | -44.2% | [-92.1%, -19.2%] | 34.7 pts | 10.1 | no |
| multi | url | 65536 | valuesBetween | ordered | baseline | 6 | 8296 | 5945 | 0.71× [0.67, 0.82] | -40.2% | [-49.2%, -22.7%] | 12.6 pts | 2.2 | no |
| multi | url | 65536 | prefix | ordered | baseline | 6 | 8503 | 6941 | 0.83× [0.79, 0.89] | -20.0% | [-27.4%, -12.4%] | 7.1 pts | 0.9 | no |
| multi | url | 65536 | churn | ordered | baseline | 6 | 532 | 337 | 0.64× [0.60, 0.71] | -55.4% | [-65.8%, -41.3%] | 11.7 pts | 2.8 | no |
| multi | url | 65536 | build | ordered | baseline | 6 | 555.68 ms | 360.09 ms | 0.64× [0.62, 0.66] | -57.0% | [-62.0%, -51.8%] | 4.8 pts | 2.4 | yes |
| multi | url | 262144 | valuesFor | ordered | baseline | 6 | 562 | 442 | 0.78× [0.76, 0.82] | -28.5% | [-31.9%, -22.3%] | 4.6 pts | 2.2 | no |
| multi | url | 262144 | valuesBetween | ordered | baseline | 6 | 12.6 µs | 10.9 µs | 0.86× [0.84, 0.87] | -15.6% | [-18.6%, -14.6%] | 1.9 pts | 0.9 | no |
| multi | url | 262144 | prefix | ordered | baseline | 6 | 61.2 µs | 52.5 µs | 0.86× [0.85, 0.87] | -15.8% | [-18.0%, -14.7%] | 1.6 pts | 0.3 | yes |
| multi | url | 262144 | churn | ordered | baseline | 6 | 981 | 555 | 0.56× [0.53, 0.63] | -77.9% | [-86.9%, -58.9%] | 13.3 pts | 3.4 | no |
| multi | url | 1048576 | valuesFor | ordered | baseline | 6 | 772 | 563 | 0.73× [0.67, 0.76] | -37.6% | [-50.0%, -31.2%] | 9.0 pts | 2.9 | no |
| multi | url | 1048576 | valuesBetween | ordered | baseline | 6 | 15.7 µs | 12.5 µs | 0.80× [0.77, 0.85] | -25.3% | [-29.7%, -17.5%] | 5.8 pts | 2.4 | no |
| multi | url | 1048576 | prefix | ordered | baseline | 6 | 199.1 µs | 224.5 µs | 1.13× [0.82, 1.39] | +11.8% | [-21.9%, +27.8%] | 23.7 pts | 1.2 | no |
| multi | url | 1048576 | churn | ordered | baseline | 6 | 1223 | 686 | 0.57× [0.51, 0.62] | -76.6% | [-96.7%, -62.0%] | 16.5 pts | 4.6 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 5 | 69.2 | 54.5 | 0.79× [0.79, 0.79] | -27.0% | [-27.2%, -26.4%] | 0.3 pts | 0.3 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 5 | 3332 | 3246 | 0.99× [0.97, 1.00] | -1.4% | [-3.5%, -0.2%] | 1.3 pts | 0.5 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 5 | 97.4 | 88.3 | 0.90× [0.90, 0.91] | -10.9% | [-11.5%, -10.3%] | 0.5 pts | 0.7 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 5 | 93.4 | 67.0 | 0.72× [0.71, 0.73] | -39.6% | [-41.6%, -37.3%] | 1.7 pts | 0.7 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 5 | 9.08 ms | 5.46 ms | 0.60× [0.60, 0.61] | -65.5% | [-66.8%, -64.5%] | 0.9 pts | 0.7 | yes |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 10 | 87.6 | 64.7 | 0.74× [0.73, 0.76] | -34.9% | [-37.7%, -31.7%] | 4.2 pts | 1.4 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 10 | 4027 | 3595 | 0.89× [0.88, 0.90] | -12.4% | [-13.3%, -10.9%] | 1.6 pts | 1.2 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 10 | 133 | 111 | 0.84× [0.83, 0.85] | -19.5% | [-21.2%, -17.0%] | 2.9 pts | 1.0 | no |
| multi | uuid | 16384 | churn | ordered | baseline | 10 | 157 | 97.9 | 0.63× [0.60, 0.66] | -58.5% | [-67.4%, -52.5%] | 10.4 pts | 1.3 | no |
| multi | uuid | 16384 | build | ordered | baseline | 10 | 54.47 ms | 28.41 ms | 0.52× [0.52, 0.53] | -91.5% | [-93.5%, -88.4%] | 3.6 pts | 1.2 | yes |
| multi | uuid | 65536 | valuesFor | ordered | baseline | 10 | 308 | 156 | 0.52× [0.47, 0.57] | -92.4% | [-112.2%, -74.6%] | 26.2 pts | 2.7 | no |
| multi | uuid | 65536 | valuesBetween | ordered | baseline | 10 | 7910 | 5714 | 0.71× [0.67, 0.78] | -41.2% | [-49.6%, -28.4%] | 14.9 pts | 2.1 | no |
| multi | uuid | 65536 | prefix | ordered | baseline | 10 | 344 | 268 | 0.75× [0.73, 0.82] | -32.9% | [-37.3%, -21.7%] | 10.9 pts | 2.3 | no |
| multi | uuid | 65536 | churn | ordered | baseline | 10 | 367 | 214 | 0.62× [0.57, 0.65] | -61.2% | [-76.0%, -54.9%] | 14.8 pts | 2.5 | no |
| multi | uuid | 65536 | build | ordered | baseline | 10 | 404.10 ms | 226.70 ms | 0.57× [0.56, 0.57] | -75.9% | [-79.0%, -74.2%] | 3.4 pts | 1.7 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 6 | 485 | 293 | 0.59× [0.56, 0.68] | -69.5% | [-79.1%, -47.4%] | 15.1 pts | 3.2 | no |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 6 | 12.3 µs | 10.2 µs | 0.85× [0.79, 0.89] | -17.7% | [-27.4%, -12.9%] | 6.9 pts | 3.2 | no |
| multi | uuid | 262144 | prefix | ordered | baseline | 6 | 820 | 688 | 0.83× [0.79, 0.87] | -20.4% | [-27.0%, -14.7%] | 5.9 pts | 2.3 | no |
| multi | uuid | 262144 | churn | ordered | baseline | 6 | 756 | 423 | 0.58× [0.53, 0.68] | -72.3% | [-89.5%, -47.0%] | 20.3 pts | 4.8 | no |
| multi | uuid | 262144 | build | ordered | baseline | 5 | 2787.19 ms | 1892.61 ms | 0.68× [0.67, 0.69] | -47.2% | [-50.0%, -45.5%] | 1.8 pts | 2.9 | yes |
| multi | uuid | 1048576 | valuesFor | ordered | baseline | 6 | 648 | 422 | 0.66× [0.62, 0.78] | -51.6% | [-61.6%, -27.9%] | 16.0 pts | 3.6 | no |
| multi | uuid | 1048576 | valuesBetween | ordered | baseline | 6 | 14.1 µs | 12.1 µs | 0.87× [0.85, 0.90] | -14.9% | [-18.1%, -11.7%] | 3.0 pts | 1.2 | no |
| multi | uuid | 1048576 | prefix | ordered | baseline | 6 | 2537 | 2201 | 0.86× [0.84, 0.89] | -16.4% | [-19.3%, -11.9%] | 3.5 pts | 1.5 | no |
| multi | uuid | 1048576 | churn | ordered | baseline | 6 | 1035 | 582 | 0.55× [0.50, 0.62] | -82.4% | [-99.1%, -61.1%] | 18.1 pts | 6.0 | no |
| unique | email | 4096 | valuesFor | ordered | baseline | 5 | 39.7 | 26.0 | 0.65× [0.63, 0.67] | -54.0% | [-58.0%, -50.3%] | 3.1 pts | 1.5 | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 5 | 345 | 1221 | 3.54× [3.52, 3.58] | +71.7% | [+71.6%, +72.1%] | 0.2 pts | 0.9 | yes |
| unique | email | 4096 | prefix | ordered | baseline | 5 | 60.2 | 55.2 | 0.92× [0.91, 0.93] | -8.8% | [-9.7%, -7.8%] | 0.8 pts | 1.1 | yes |
| unique | email | 4096 | churn | ordered | baseline | 5 | 91.4 | 67.7 | 0.74× [0.73, 0.75] | -35.4% | [-36.3%, -33.8%] | 1.0 pts | 0.6 | yes |
| unique | email | 4096 | build | ordered | baseline | 5 | 1.23 ms | 842.8 µs | 0.69× [0.68, 0.70] | -45.4% | [-47.9%, -43.4%] | 1.8 pts | 0.8 | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 56.6 | 38.1 | 0.67× [0.66, 0.68] | -49.8% | [-52.5%, -47.4%] | 2.4 pts | 2.2 | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 455 | 1560 | 3.43× [3.34, 3.62] | +70.9% | [+70.1%, +72.4%] | 1.1 pts | 3.6 | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 85.9 | 70.3 | 0.82× [0.81, 0.83] | -22.3% | [-23.1%, -20.9%] | 1.0 pts | 1.0 | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 129 | 94.9 | 0.74× [0.72, 0.76] | -35.8% | [-39.0%, -32.3%] | 3.2 pts | 1.0 | yes |
| unique | email | 16384 | build | ordered | baseline | 6 | 6.75 ms | 4.07 ms | 0.60× [0.60, 0.61] | -66.7% | [-67.8%, -64.6%] | 1.5 pts | 1.1 | yes |
| unique | email | 65536 | valuesFor | ordered | baseline | 6 | 70.1 | 53.6 | 0.77× [0.75, 0.80] | -30.4% | [-33.2%, -25.8%] | 3.5 pts | 1.2 | no |
| unique | email | 65536 | valuesBetween | ordered | baseline | 6 | 514 | 1548 | 3.04× [2.89, 3.12] | +67.1% | [+65.4%, +68.0%] | 1.3 pts | 2.0 | yes |
| unique | email | 65536 | prefix | ordered | baseline | 6 | 105 | 103 | 0.98× [0.95, 1.01] | -2.1% | [-5.4%, +1.0%] | 3.0 pts | 0.9 | no |
| unique | email | 65536 | churn | ordered | baseline | 6 | 171 | 160 | 0.93× [0.85, 1.06] | -7.4% | [-17.8%, +6.0%] | 11.4 pts | 1.5 | no |
| unique | email | 65536 | build | ordered | baseline | 6 | 35.30 ms | 28.99 ms | 0.82× [0.81, 0.83] | -21.3% | [-22.7%, -20.6%] | 1.0 pts | 0.8 | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 6 | 267 | 204 | 0.75× [0.65, 0.83] | -34.0% | [-53.5%, -20.0%] | 16.0 pts | 1.9 | no |
| unique | email | 262144 | valuesBetween | ordered | baseline | 6 | 1546 | 4026 | 2.60× [2.55, 2.66] | +61.6% | [+60.8%, +62.3%] | 0.7 pts | 0.6 | yes |
| unique | email | 262144 | prefix | ordered | baseline | 6 | 308 | 277 | 0.93× [0.81, 1.06] | -7.7% | [-23.3%, +5.4%] | 13.7 pts | 3.6 | no |
| unique | email | 262144 | churn | ordered | baseline | 6 | 521 | 364 | 0.71× [0.65, 0.82] | -41.7% | [-54.1%, -22.6%] | 15.0 pts | 3.3 | no |
| unique | email | 1048576 | valuesFor | ordered | baseline | 6 | 431 | 310 | 0.71× [0.68, 0.74] | -40.8% | [-47.2%, -34.4%] | 6.1 pts | 1.6 | no |
| unique | email | 1048576 | valuesBetween | ordered | baseline | 6 | 1913 | 6498 | 3.38× [3.32, 3.43] | +70.4% | [+69.9%, +70.9%] | 0.5 pts | 0.9 | yes |
| unique | email | 1048576 | prefix | ordered | baseline | 6 | 482 | 438 | 0.91× [0.83, 0.99] | -10.4% | [-20.4%, -1.2%] | 9.2 pts | 3.2 | no |
| unique | email | 1048576 | churn | ordered | baseline | 6 | 670 | 555 | 0.80× [0.76, 0.86] | -24.4% | [-31.3%, -15.8%] | 7.4 pts | 2.2 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 5 | 98.1 | 83.6 | 0.85× [0.85, 0.86] | -17.3% | [-18.2%, -16.8%] | 0.6 pts | 0.6 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 5 | 674 | 2023 | 3.01× [2.94, 3.10] | +66.8% | [+66.0%, +67.7%] | 0.7 pts | 1.4 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 5 | 230 | 242 | 1.05× [1.04, 1.07] | +5.1% | [+3.7%, +6.2%] | 1.0 pts | 0.7 | yes |
| unique | path | 4096 | churn | ordered | baseline | 5 | 215 | 199 | 0.92× [0.92, 0.93] | -8.7% | [-9.1%, -8.0%] | 0.4 pts | 0.3 | yes |
| unique | path | 4096 | build | ordered | baseline | 5 | 2.86 ms | 2.06 ms | 0.72× [0.71, 0.73] | -39.6% | [-40.5%, -37.9%] | 1.1 pts | 0.9 | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 6 | 130 | 114 | 0.88× [0.87, 0.89] | -13.6% | [-14.7%, -12.9%] | 0.9 pts | 0.6 | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 6 | 852 | 2565 | 3.01× [2.99, 3.05] | +66.8% | [+66.6%, +67.2%] | 0.3 pts | 0.7 | yes |
| unique | path | 16384 | prefix | ordered | baseline | 6 | 318 | 422 | 1.33× [1.31, 1.35] | +24.7% | [+23.8%, +26.2%] | 1.1 pts | 0.4 | yes |
| unique | path | 16384 | churn | ordered | baseline | 6 | 283 | 261 | 0.93× [0.91, 0.95] | -7.9% | [-9.5%, -5.6%] | 1.9 pts | 1.1 | yes |
| unique | path | 16384 | build | ordered | baseline | 6 | 15.08 ms | 10.90 ms | 0.73× [0.72, 0.73] | -37.1% | [-38.4%, -36.3%] | 1.0 pts | 0.9 | yes |
| unique | path | 65536 | valuesFor | ordered | baseline | 6 | 235 | 218 | 0.91× [0.87, 0.98] | -9.5% | [-14.6%, -1.8%] | 6.1 pts | 1.5 | no |
| unique | path | 65536 | valuesBetween | ordered | baseline | 6 | 1213 | 2909 | 2.42× [2.34, 2.49] | +58.7% | [+57.2%, +59.8%] | 1.3 pts | 1.6 | yes |
| unique | path | 65536 | prefix | ordered | baseline | 6 | 701 | 1197 | 1.71× [1.65, 1.83] | +41.4% | [+39.3%, +45.5%] | 3.0 pts | 0.9 | yes |
| unique | path | 65536 | churn | ordered | baseline | 6 | 459 | 434 | 0.95× [0.92, 0.98] | -5.3% | [-8.5%, -1.6%] | 3.3 pts | 0.7 | no |
| unique | path | 65536 | build | ordered | baseline | 6 | 93.77 ms | 71.91 ms | 0.78× [0.76, 0.80] | -28.5% | [-31.0%, -24.4%] | 3.2 pts | 2.8 | no |
| unique | path | 262144 | valuesFor | ordered | baseline | 6 | 541 | 443 | 0.84× [0.77, 0.98] | -18.8% | [-30.6%, -2.3%] | 13.5 pts | 3.1 | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 6 | 2753 | 6559 | 2.37× [2.30, 2.45] | +57.7% | [+56.5%, +59.2%] | 1.3 pts | 0.7 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 6 | 1900 | 5790 | 3.04× [2.79, 3.26] | +67.1% | [+64.2%, +69.3%] | 2.4 pts | 0.6 | yes |
| unique | path | 262144 | churn | ordered | baseline | 6 | 1063 | 773 | 0.76× [0.65, 0.86] | -31.1% | [-53.5%, -16.1%] | 17.8 pts | 4.6 | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 6 | 45.6 | 37.6 | 0.83× [0.81, 0.84] | -21.1% | [-23.3%, -19.7%] | 1.7 pts | 1.7 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 6 | 384 | 1560 | 4.06× [4.04, 4.10] | +75.4% | [+75.2%, +75.6%] | 0.2 pts | 0.8 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 6 | 593 | 3053 | 5.15× [5.11, 5.19] | +80.6% | [+80.4%, +80.7%] | 0.1 pts | 1.1 | yes |
| unique | str | 4096 | churn | ordered | baseline | 6 | 110 | 86.2 | 0.78× [0.78, 0.79] | -27.6% | [-29.0%, -26.5%] | 1.2 pts | 0.7 | yes |
| unique | str | 4096 | build | ordered | baseline | 6 | 1.45 ms | 1.02 ms | 0.71× [0.69, 0.72] | -41.1% | [-45.0%, -38.8%] | 3.0 pts | 1.0 | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 6 | 58.5 | 49.9 | 0.85× [0.84, 0.86] | -17.4% | [-18.5%, -16.2%] | 1.1 pts | 1.3 | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 6 | 470 | 1653 | 3.51× [3.47, 3.54] | +71.5% | [+71.2%, +71.7%] | 0.3 pts | 1.0 | yes |
| unique | str | 16384 | prefix | ordered | baseline | 6 | 2546 | 13.8 µs | 5.42× [5.40, 5.47] | +81.5% | [+81.5%, +81.7%] | 0.1 pts | 0.6 | yes |
| unique | str | 16384 | churn | ordered | baseline | 6 | 139 | 114 | 0.82× [0.80, 0.85] | -21.7% | [-25.4%, -18.1%] | 3.4 pts | 1.6 | no |
| unique | str | 16384 | build | ordered | baseline | 6 | 7.11 ms | 4.98 ms | 0.70× [0.70, 0.71] | -42.5% | [-43.4%, -41.4%] | 0.9 pts | 0.9 | yes |
| unique | str | 65536 | valuesFor | ordered | baseline | 6 | 90.6 | 72.3 | 0.80× [0.78, 0.84] | -24.2% | [-28.5%, -18.8%] | 4.6 pts | 1.6 | no |
| unique | str | 65536 | valuesBetween | ordered | baseline | 6 | 665 | 1801 | 2.70× [2.60, 2.75] | +63.0% | [+61.5%, +63.7%] | 1.1 pts | 2.0 | yes |
| unique | str | 65536 | prefix | ordered | baseline | 6 | 15.6 µs | 61.2 µs | 3.93× [3.86, 3.97] | +74.5% | [+74.1%, +74.8%] | 0.3 pts | 0.9 | yes |
| unique | str | 65536 | churn | ordered | baseline | 6 | 188 | 171 | 0.93× [0.87, 0.99] | -7.0% | [-14.6%, -0.7%] | 6.6 pts | 1.3 | no |
| unique | str | 65536 | build | ordered | baseline | 6 | 47.84 ms | 32.73 ms | 0.69× [0.68, 0.69] | -45.7% | [-47.0%, -44.1%] | 1.4 pts | 0.7 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 6 | 169 | 221 | 1.37× [1.29, 1.40] | +26.9% | [+22.6%, +28.5%] | 2.8 pts | 0.6 | no |
| unique | str | 262144 | valuesBetween | ordered | baseline | 6 | 1126 | 4478 | 4.02× [3.96, 4.09] | +75.1% | [+74.8%, +75.6%] | 0.4 pts | 0.6 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 6 | 121.2 µs | 650.1 µs | 5.43× [5.29, 5.54] | +81.6% | [+81.1%, +82.0%] | 0.4 pts | 0.5 | yes |
| unique | str | 262144 | churn | ordered | baseline | 6 | 358 | 393 | 1.19× [0.84, 1.51] | +16.0% | [-19.7%, +33.6%] | 25.4 pts | 6.1 | no |
| unique | str | 1048576 | valuesFor | ordered | baseline | 6 | 356 | 347 | 0.97× [0.94, 1.03] | -2.7% | [-6.1%, +3.3%] | 4.4 pts | 1.6 | no |
| unique | str | 1048576 | valuesBetween | ordered | baseline | 6 | 1795 | 6808 | 3.77× [3.68, 3.87] | +73.5% | [+72.9%, +74.2%] | 0.6 pts | 1.8 | yes |
| unique | str | 1048576 | prefix | ordered | baseline | 6 | 902.4 µs | 4.26 ms | 4.71× [4.67, 4.77] | +78.8% | [+78.6%, +79.0%] | 0.2 pts | 1.2 | yes |
| unique | str | 1048576 | churn | ordered | baseline | 6 | 785 | 586 | 0.76× [0.69, 0.85] | -31.5% | [-44.9%, -17.9%] | 12.9 pts | 4.1 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 54.3 | 48.9 | 0.90× [0.89, 0.90] | -11.1% | [-12.0%, -10.6%] | 0.7 pts | 0.6 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 471 | 1772 | 3.75× [3.72, 3.79] | +73.3% | [+73.1%, +73.6%] | 0.2 pts | 0.7 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 109 | 211 | 1.93× [1.91, 1.97] | +48.2% | [+47.6%, +49.2%] | 0.8 pts | 1.4 | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 121 | 108 | 0.90× [0.88, 0.91] | -11.5% | [-13.6%, -9.5%] | 2.0 pts | 1.5 | no |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.86 ms | 1.23 ms | 0.66× [0.66, 0.67] | -50.5% | [-52.1%, -49.0%] | 1.5 pts | 0.9 | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 6 | 70.6 | 70.2 | 1.00× [0.99, 1.00] | -0.3% | [-0.8%, +0.5%] | 0.6 pts | 0.7 | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 6 | 570 | 2151 | 3.75× [3.69, 3.81] | +73.3% | [+72.9%, +73.7%] | 0.4 pts | 1.5 | yes |
| unique | street | 16384 | prefix | ordered | baseline | 6 | 210 | 712 | 3.40× [3.30, 3.45] | +70.6% | [+69.7%, +71.0%] | 0.6 pts | 1.2 | yes |
| unique | street | 16384 | churn | ordered | baseline | 6 | 155 | 146 | 0.95× [0.93, 0.97] | -5.0% | [-7.7%, -3.3%] | 2.1 pts | 1.1 | no |
| unique | street | 16384 | build | ordered | baseline | 6 | 8.79 ms | 6.25 ms | 0.71× [0.70, 0.72] | -40.6% | [-41.9%, -39.7%] | 1.0 pts | 1.3 | yes |
| unique | street | 65536 | valuesFor | ordered | baseline | 6 | 100 | 97.4 | 0.98× [0.95, 0.99] | -2.4% | [-5.7%, -1.2%] | 2.1 pts | 1.1 | no |
| unique | street | 65536 | valuesBetween | ordered | baseline | 6 | 703 | 2359 | 3.36× [3.31, 3.41] | +70.2% | [+69.8%, +70.7%] | 0.4 pts | 0.8 | yes |
| unique | street | 65536 | prefix | ordered | baseline | 6 | 679 | 3040 | 4.48× [4.31, 4.58] | +77.7% | [+76.8%, +78.2%] | 0.7 pts | 1.1 | yes |
| unique | street | 65536 | churn | ordered | baseline | 6 | 234 | 235 | 1.03× [0.97, 1.10] | +3.1% | [-2.9%, +9.1%] | 5.7 pts | 0.8 | no |
| unique | street | 65536 | build | ordered | baseline | 6 | 48.48 ms | 40.31 ms | 0.83× [0.82, 0.84] | -20.6% | [-22.3%, -19.4%] | 1.4 pts | 1.2 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 5 | 16.8 | 15.3 | 0.91× [0.90, 0.92] | -9.8% | [-11.4%, -9.1%] | 0.9 pts | 0.3 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 5 | 242 | 1021 | 4.21× [4.17, 4.25] | +76.2% | [+76.0%, +76.5%] | 0.2 pts | 1.5 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 5 | 44.7 | 42.0 | 0.94× [0.93, 0.95] | -6.2% | [-8.1%, -5.1%] | 1.2 pts | 1.4 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 5 | 591.1 µs | 584.4 µs | 0.99× [0.98, 1.00] | -0.6% | [-1.8%, +0.4%] | 0.9 pts | 0.8 | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 5 | 31.7 | 22.8 | 0.72× [0.71, 0.72] | -38.9% | [-41.0%, -38.1%] | 1.2 pts | 1.2 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 5 | 285 | 1575 | 5.51× [5.47, 5.56] | +81.8% | [+81.7%, +82.0%] | 0.1 pts | 1.0 | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 5 | 65.6 | 45.8 | 0.70× [0.68, 0.72] | -42.9% | [-46.8%, -38.9%] | 3.2 pts | 1.6 | yes |
| unique | u64 | 16384 | build | ordered | baseline | 5 | 3.40 ms | 2.66 ms | 0.79× [0.78, 0.79] | -26.7% | [-27.8%, -26.1%] | 0.7 pts | 0.5 | yes |
| unique | u64 | 65536 | valuesFor | ordered | baseline | 6 | 37.3 | 25.2 | 0.68× [0.66, 0.70] | -47.8% | [-50.6%, -42.5%] | 3.8 pts | 2.8 | yes |
| unique | u64 | 65536 | valuesBetween | ordered | baseline | 6 | 282 | 1923 | 6.81× [6.77, 6.84] | +85.3% | [+85.2%, +85.4%] | 0.1 pts | 0.7 | yes |
| unique | u64 | 65536 | churn | ordered | baseline | 6 | 79.9 | 89.4 | 1.13× [0.99, 1.24] | +11.4% | [-0.8%, +19.5%] | 9.7 pts | 4.1 | no |
| unique | u64 | 65536 | build | ordered | baseline | 6 | 15.70 ms | 14.43 ms | 0.92× [0.88, 0.98] | -8.2% | [-13.1%, -1.9%] | 5.4 pts | 4.4 | no |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 6 | 49.1 | 76.8 | 1.56× [1.51, 1.69] | +36.0% | [+33.7%, +40.7%] | 3.3 pts | 1.3 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 6 | 396 | 2634 | 6.61× [6.39, 6.89] | +84.9% | [+84.4%, +85.5%] | 0.5 pts | 0.8 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 6 | 172 | 210 | 1.15× [1.03, 1.43] | +13.1% | [+3.0%, +30.1%] | 12.9 pts | 2.7 | no |
| unique | u64 | 1048576 | valuesFor | ordered | baseline | 6 | 141 | 137 | 0.97× [0.91, 1.07] | -3.1% | [-9.8%, +6.5%] | 7.8 pts | 2.3 | no |
| unique | u64 | 1048576 | valuesBetween | ordered | baseline | 6 | 773 | 3098 | 4.02× [3.92, 4.14] | +75.1% | [+74.5%, +75.9%] | 0.6 pts | 1.7 | yes |
| unique | u64 | 1048576 | churn | ordered | baseline | 6 | 366 | 366 | 1.05× [0.78, 1.32] | +4.8% | [-28.2%, +24.5%] | 25.1 pts | 5.3 | no |
| unique | url | 4096 | valuesFor | ordered | baseline | 6 | 69.6 | 60.8 | 0.87× [0.87, 0.88] | -14.7% | [-15.4%, -13.8%] | 0.7 pts | 1.0 | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 6 | 544 | 1684 | 3.09× [3.04, 3.13] | +67.6% | [+67.2%, +68.0%] | 0.4 pts | 1.2 | yes |
| unique | url | 4096 | prefix | ordered | baseline | 6 | 172 | 239 | 1.39× [1.37, 1.40] | +28.2% | [+27.3%, +28.8%] | 0.7 pts | 0.9 | yes |
| unique | url | 4096 | churn | ordered | baseline | 6 | 152 | 137 | 0.89× [0.88, 0.91] | -12.2% | [-13.7%, -9.6%] | 1.9 pts | 1.2 | no |
| unique | url | 4096 | build | ordered | baseline | 6 | 2.24 ms | 1.63 ms | 0.72× [0.71, 0.74] | -38.1% | [-40.1%, -34.9%] | 2.5 pts | 0.5 | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 81.6 | 75.0 | 0.92× [0.90, 0.93] | -9.0% | [-11.0%, -8.0%] | 1.4 pts | 1.4 | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 603 | 1883 | 3.12× [3.07, 3.27] | +68.0% | [+67.4%, +69.4%] | 1.0 pts | 1.6 | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 257 | 638 | 2.48× [2.21, 2.63] | +59.6% | [+54.9%, +62.0%] | 3.4 pts | 6.0 | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 187 | 163 | 0.87× [0.86, 0.90] | -15.0% | [-16.7%, -10.8%] | 2.8 pts | 1.6 | no |
| unique | url | 16384 | build | ordered | baseline | 6 | 9.58 ms | 7.50 ms | 0.78× [0.77, 0.79] | -28.1% | [-29.6%, -27.1%] | 1.2 pts | 1.0 | yes |
| unique | url | 65536 | valuesFor | ordered | baseline | 6 | 181 | 164 | 0.95× [0.78, 1.09] | -5.6% | [-28.1%, +8.3%] | 17.4 pts | 4.1 | no |
| unique | url | 65536 | valuesBetween | ordered | baseline | 6 | 855 | 2270 | 2.65× [2.59, 2.72] | +62.3% | [+61.4%, +63.2%] | 0.9 pts | 1.2 | yes |
| unique | url | 65536 | prefix | ordered | baseline | 6 | 788 | 2498 | 3.16× [3.13, 3.20] | +68.3% | [+68.0%, +68.8%] | 0.4 pts | 0.5 | yes |
| unique | url | 65536 | churn | ordered | baseline | 6 | 341 | 331 | 0.98× [0.96, 1.03] | -1.5% | [-4.2%, +2.5%] | 3.2 pts | 0.6 | no |
| unique | url | 65536 | build | ordered | baseline | 6 | 70.22 ms | 58.75 ms | 0.84× [0.81, 0.86] | -19.3% | [-23.8%, -16.3%] | 3.6 pts | 2.2 | no |
| unique | url | 262144 | valuesFor | ordered | baseline | 6 | 388 | 358 | 0.90× [0.84, 1.06] | -10.5% | [-18.5%, +5.4%] | 11.4 pts | 2.6 | no |
| unique | url | 262144 | valuesBetween | ordered | baseline | 6 | 1670 | 5106 | 3.06× [2.96, 3.19] | +67.3% | [+66.2%, +68.6%] | 1.2 pts | 1.1 | yes |
| unique | url | 262144 | prefix | ordered | baseline | 6 | 4859 | 21.3 µs | 4.41× [4.34, 4.46] | +77.3% | [+77.0%, +77.6%] | 0.3 pts | 0.3 | yes |
| unique | url | 262144 | churn | ordered | baseline | 6 | 568 | 595 | 1.04× [0.95, 1.13] | +3.6% | [-5.1%, +11.4%] | 7.9 pts | 2.5 | no |
| unique | url | 1048576 | valuesFor | ordered | baseline | 6 | 639 | 499 | 0.78× [0.75, 0.81] | -28.1% | [-33.0%, -23.6%] | 4.5 pts | 1.4 | no |
| unique | url | 1048576 | valuesBetween | ordered | baseline | 6 | 2633 | 7468 | 2.82× [2.71, 2.88] | +64.5% | [+63.1%, +65.2%] | 1.0 pts | 1.6 | yes |
| unique | url | 1048576 | prefix | ordered | baseline | 6 | 31.8 µs | 132.5 µs | 4.15× [4.09, 4.21] | +75.9% | [+75.5%, +76.3%] | 0.3 pts | 0.4 | yes |
| unique | url | 1048576 | churn | ordered | baseline | 6 | 895 | 784 | 0.88× [0.83, 0.92] | -14.1% | [-20.8%, -8.7%] | 5.8 pts | 1.9 | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 5 | 44.1 | 31.0 | 0.70× [0.68, 0.72] | -43.8% | [-46.5%, -38.8%] | 3.1 pts | 2.2 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 5 | 405 | 1437 | 3.53× [3.49, 3.60] | +71.7% | [+71.3%, +72.2%] | 0.4 pts | 1.3 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 5 | 75.6 | 66.6 | 0.88× [0.88, 0.88] | -13.4% | [-14.2%, -13.1%] | 0.4 pts | 0.5 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 5 | 97.2 | 80.9 | 0.83× [0.82, 0.84] | -20.4% | [-21.8%, -19.5%] | 0.9 pts | 0.7 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 5 | 1.41 ms | 970.9 µs | 0.69× [0.68, 0.70] | -45.6% | [-46.7%, -43.5%] | 1.3 pts | 0.4 | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 6 | 50.4 | 40.8 | 0.81× [0.80, 0.82] | -23.2% | [-25.1%, -22.4%] | 1.3 pts | 1.4 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 422 | 1590 | 3.78× [3.70, 3.86] | +73.5% | [+73.0%, +74.1%] | 0.6 pts | 1.5 | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 6 | 83.9 | 82.0 | 0.98× [0.97, 0.98] | -2.5% | [-2.7%, -2.1%] | 0.3 pts | 0.3 | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 6 | 118 | 108 | 0.90× [0.87, 0.93] | -11.3% | [-15.5%, -8.0%] | 3.6 pts | 1.5 | no |
| unique | uuid | 16384 | build | ordered | baseline | 6 | 5.90 ms | 4.60 ms | 0.78× [0.77, 0.78] | -28.6% | [-29.2%, -27.8%] | 0.7 pts | 0.6 | yes |
| unique | uuid | 65536 | valuesFor | ordered | baseline | 6 | 90.2 | 66.7 | 0.75× [0.69, 0.81] | -33.7% | [-44.1%, -23.5%] | 9.8 pts | 0.9 | no |
| unique | uuid | 65536 | valuesBetween | ordered | baseline | 6 | 719 | 1889 | 2.63× [2.58, 2.67] | +61.9% | [+61.2%, +62.5%] | 0.6 pts | 0.7 | yes |
| unique | uuid | 65536 | prefix | ordered | baseline | 6 | 135 | 141 | 1.05× [0.99, 1.17] | +5.0% | [-1.3%, +14.4%] | 7.5 pts | 2.0 | no |
| unique | uuid | 65536 | churn | ordered | baseline | 6 | 210 | 189 | 0.93× [0.85, 1.04] | -7.5% | [-17.7%, +4.0%] | 10.3 pts | 1.6 | no |
| unique | uuid | 65536 | build | ordered | baseline | 6 | 47.94 ms | 34.40 ms | 0.71× [0.68, 0.74] | -40.3% | [-47.6%, -35.9%] | 5.6 pts | 2.7 | no |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 6 | 255 | 209 | 0.80× [0.74, 0.88] | -25.5% | [-34.8%, -13.5%] | 10.1 pts | 1.5 | no |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 6 | 1347 | 4721 | 3.51× [3.42, 3.60] | +71.5% | [+70.8%, +72.2%] | 0.7 pts | 0.9 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 6 | 275 | 360 | 1.31× [1.18, 1.52] | +23.7% | [+14.9%, +34.4%] | 9.3 pts | 3.4 | no |
| unique | uuid | 262144 | churn | ordered | baseline | 6 | 458 | 422 | 1.00× [0.75, 1.24] | -0.0% | [-34.2%, +19.1%] | 25.4 pts | 5.2 | no |
| unique | uuid | 1048576 | valuesFor | ordered | baseline | 6 | 471 | 330 | 0.70× [0.67, 0.72] | -43.5% | [-49.4%, -39.3%] | 4.8 pts | 1.8 | no |
| unique | uuid | 1048576 | valuesBetween | ordered | baseline | 6 | 2487 | 6989 | 2.82× [2.76, 2.85] | +64.5% | [+63.8%, +64.9%] | 0.5 pts | 0.9 | yes |
| unique | uuid | 1048576 | prefix | ordered | baseline | 6 | 559 | 1289 | 2.33× [2.09, 2.49] | +57.1% | [+52.2%, +59.9%] | 3.7 pts | 2.7 | yes |
| unique | uuid | 1048576 | churn | ordered | baseline | 6 | 842 | 590 | 0.69× [0.64, 0.74] | -44.8% | [-55.7%, -34.8%] | 10.0 pts | 3.9 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
