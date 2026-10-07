| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 144 | 134 | 0.93× [0.93, 0.94] | -7.0% | [-7.9%, -6.1%] | 1.4 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 146 | 218 | 1.50× [1.47, 1.54] | +33.5% | [+31.8%, +35.1%] | 1.0 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 3492 | 4194 | 1.20× [1.18, 1.22] | +16.8% | [+15.3%, +18.3%] | 1.1 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 3496 | 6487 | 1.86× [1.83, 1.89] | +46.3% | [+45.5%, +47.1%] | 0.6 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 8 | 4035 | 4986 | 1.22× [1.21, 1.23] | +17.8% | [+17.1%, +18.4%] | 1.0 pts | 0.3 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 8 | 3976 | 8824 | 2.20× [2.12, 2.30] | +54.6% | [+52.8%, +56.5%] | 1.2 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 250 | 207 | 0.84× [0.83, 0.84] | -19.6% | [-20.7%, -18.4%] | 0.7 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 254 | 291 | 1.16× [1.13, 1.19] | +13.6% | [+11.2%, +15.9%] | 1.6 pts | 1.5 | no | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 12.32 ms | 9.83 ms | 0.80× [0.79, 0.81] | -25.2% | [-26.9%, -23.6%] | 1.1 pts | 1.8 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 12.30 ms | 13.62 ms | 1.11× [1.10, 1.12] | +9.9% | [+9.4%, +10.4%] | 0.6 pts | 1.5 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 174 | 164 | 0.94× [0.93, 0.95] | -6.2% | [-7.7%, -4.7%] | 1.1 pts | 2.2 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 180 | 279 | 1.58× [1.50, 1.67] | +36.6% | [+33.2%, +40.0%] | 2.5 pts | 2.3 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3599 | 4527 | 1.26× [1.25, 1.27] | +20.7% | [+20.0%, +21.5%] | 1.3 pts | 2.0 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3617 | 6892 | 1.93× [1.81, 2.06] | +48.1% | [+44.8%, +51.4%] | 2.0 pts | 4.9 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 8 | 17.0 µs | 23.3 µs | 1.38× [1.36, 1.40] | +27.6% | [+26.6%, +28.5%] | 0.9 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 8 | 16.2 µs | 36.0 µs | 2.18× [2.17, 2.20] | +54.2% | [+53.8%, +54.5%] | 0.2 pts | 0.9 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 289 | 255 | 0.90× [0.86, 0.95] | -10.9% | [-16.2%, -5.6%] | 3.7 pts | 8.7 | no | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 308 | 380 | 1.24× [1.20, 1.29] | +19.6% | [+16.7%, +22.5%] | 2.2 pts | 2.5 | no | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 55.26 ms | 46.40 ms | 0.84× [0.84, 0.84] | -19.0% | [-19.2%, -18.9%] | 0.3 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 55.22 ms | 66.76 ms | 1.22× [1.19, 1.25] | +17.8% | [+15.6%, +20.0%] | 1.3 pts | 8.3 | no | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 226 | 254 | 1.13× [1.11, 1.16] | +11.9% | [+9.9%, +13.8%] | 2.0 pts | 1.1 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 254 | 502 | 2.00× [1.98, 2.03] | +50.1% | [+49.5%, +50.6%] | 0.6 pts | 0.4 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 4118 | 5675 | 1.38× [1.36, 1.40] | +27.6% | [+26.6%, +28.7%] | 1.1 pts | 1.5 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 4279 | 11.1 µs | 2.60× [2.59, 2.62] | +61.6% | [+61.4%, +61.8%] | 0.2 pts | 0.5 | yes | yes |
| natural | dirs | 65536 | prefix | ordered | baseline | 8 | 61.8 µs | 91.4 µs | 1.45× [1.44, 1.46] | +31.0% | [+30.5%, +31.6%] | 0.8 pts | 1.2 | yes | yes |
| natural | dirs | 65536 | prefix | ordered | btree-sets | 8 | 66.6 µs | 201.2 µs | 3.06× [3.05, 3.08] | +67.3% | [+67.2%, +67.5%] | 0.2 pts | 0.8 | yes | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 451 | 441 | 0.94× [0.92, 0.96] | -6.4% | [-8.6%, -4.2%] | 3.1 pts | 0.4 | no | yes |
| natural | dirs | 65536 | churn | ordered | btree-sets | 8 | 512 | 709 | 1.39× [1.35, 1.43] | +27.9% | [+25.8%, +29.9%] | 1.5 pts | 1.3 | yes | yes |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 320.57 ms | 305.69 ms | 0.95× [0.95, 0.96] | -5.1% | [-5.7%, -4.4%] | 0.5 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | build | ordered | btree-sets | 8 | 321.52 ms | 471.74 ms | 1.47× [1.46, 1.47] | +31.8% | [+31.4%, +32.2%] | 0.4 pts | 1.8 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 4 | 92.6 | 77.4 | 0.84× [0.83, 0.84] | -19.5% | [-20.6%, -18.4%] | 0.7 pts | 0.7 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 4 | 92.7 | 183 | 1.97× [1.96, 1.98] | +49.2% | [+49.0%, +49.4%] | 0.1 pts | 0.3 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 4 | 2493 | 3472 | 1.39× [1.38, 1.40] | +28.2% | [+27.7%, +28.8%] | 0.3 pts | 0.5 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 4 | 2471 | 5353 | 2.18× [2.16, 2.20] | +54.1% | [+53.6%, +54.6%] | 0.3 pts | 1.0 | yes | yes |
| natural | street | 4096 | prefix | ordered | baseline | 4 | 324 | 344 | 1.06× [1.05, 1.08] | +5.7% | [+4.4%, +7.0%] | 0.8 pts | 0.9 | yes | yes |
| natural | street | 4096 | prefix | ordered | btree-sets | 4 | 324 | 646 | 2.01× [1.99, 2.04] | +50.3% | [+49.8%, +50.9%] | 0.3 pts | 0.7 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 4 | 179 | 136 | 0.76× [0.76, 0.77] | -31.2% | [-32.4%, -30.1%] | 0.7 pts | 1.3 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 4 | 180 | 240 | 1.34× [1.32, 1.36] | +25.2% | [+24.2%, +26.2%] | 0.6 pts | 0.8 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 4 | 6.90 ms | 5.05 ms | 0.73× [0.73, 0.74] | -36.5% | [-37.4%, -35.6%] | 0.6 pts | 1.9 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 4 | 6.93 ms | 9.35 ms | 1.35× [1.33, 1.36] | +25.7% | [+24.9%, +26.5%] | 0.5 pts | 1.6 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 113 | 98.7 | 0.88× [0.85, 0.91] | -13.7% | [-17.4%, -10.0%] | 2.4 pts | 4.9 | no | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 114 | 233 | 2.05× [1.96, 2.15] | +51.3% | [+49.1%, +53.5%] | 1.4 pts | 3.5 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 2739 | 3756 | 1.37× [1.35, 1.39] | +27.0% | [+26.1%, +27.9%] | 0.6 pts | 1.2 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 2750 | 5666 | 2.07× [2.01, 2.15] | +51.8% | [+50.2%, +53.4%] | 1.0 pts | 3.3 | yes | yes |
| natural | street | 16384 | prefix | ordered | baseline | 8 | 949 | 1247 | 1.33× [1.30, 1.36] | +24.8% | [+23.3%, +26.3%] | 1.0 pts | 0.7 | yes | yes |
| natural | street | 16384 | prefix | ordered | btree-sets | 8 | 948 | 2259 | 2.41× [2.37, 2.44] | +58.5% | [+57.9%, +59.1%] | 0.7 pts | 1.1 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 207 | 169 | 0.82× [0.81, 0.83] | -21.9% | [-22.9%, -21.0%] | 1.0 pts | 1.8 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 213 | 304 | 1.43× [1.42, 1.43] | +30.0% | [+29.8%, +30.3%] | 0.4 pts | 0.5 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 32.53 ms | 25.59 ms | 0.79× [0.79, 0.79] | -27.2% | [-27.3%, -27.1%] | 0.2 pts | 1.0 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 32.63 ms | 47.06 ms | 1.45× [1.43, 1.46] | +30.8% | [+30.3%, +31.4%] | 0.5 pts | 2.5 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 144 | 144 | 1.02× [1.01, 1.03] | +1.8% | [+0.8%, +2.8%] | 4.1 pts | 3.8 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | btree-sets | 8 | 163 | 414 | 2.52× [2.42, 2.63] | +60.3% | [+58.7%, +61.9%] | 1.1 pts | 1.0 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 3020 | 4380 | 1.46× [1.44, 1.48] | +31.4% | [+30.5%, +32.3%] | 1.4 pts | 1.7 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 3168 | 9471 | 3.02× [2.97, 3.08] | +66.9% | [+66.3%, +67.5%] | 1.5 pts | 3.5 | yes | yes |
| natural | street | 65536 | prefix | ordered | baseline | 8 | 3457 | 5334 | 1.59× [1.55, 1.63] | +37.1% | [+35.5%, +38.6%] | 1.8 pts | 0.9 | yes | yes |
| natural | street | 65536 | prefix | ordered | btree-sets | 8 | 3526 | 11.2 µs | 3.29× [3.21, 3.38] | +69.6% | [+68.8%, +70.4%] | 1.1 pts | 1.3 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 295 | 295 | 0.98× [0.95, 1.02] | -1.6% | [-5.0%, +1.8%] | 4.1 pts | 0.9 | no | no |
| natural | street | 65536 | churn | ordered | btree-sets | 8 | 339 | 529 | 1.55× [1.52, 1.58] | +35.3% | [+34.0%, +36.6%] | 1.3 pts | 5.0 | yes | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 179.33 ms | 168.06 ms | 0.95× [0.95, 0.95] | -5.3% | [-5.8%, -4.8%] | 2.2 pts | 2.0 | yes | yes |
| natural | street | 65536 | build | ordered | btree-sets | 8 | 179.28 ms | 289.00 ms | 1.63× [1.61, 1.65] | +38.6% | [+37.9%, +39.3%] | 1.0 pts | 4.2 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 4 | 60.5 | 50.7 | 0.83× [0.83, 0.84] | -19.9% | [-20.2%, -19.6%] | 0.2 pts | 0.3 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 4 | 61.6 | 196 | 3.21× [3.16, 3.25] | +68.8% | [+68.4%, +69.2%] | 0.3 pts | 1.5 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 4 | 4012 | 4006 | 1.00× [0.99, 1.00] | -0.5% | [-1.2%, +0.3%] | 0.5 pts | 0.6 | yes | no |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 4 | 4017 | 8243 | 2.06× [2.05, 2.08] | +51.5% | [+51.2%, +51.9%] | 0.2 pts | 0.8 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 4 | 83.1 | 69.4 | 0.84× [0.83, 0.85] | -19.3% | [-20.4%, -18.3%] | 0.7 pts | 1.5 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 4 | 82.6 | 187 | 2.27× [2.25, 2.29] | +56.0% | [+55.6%, +56.3%] | 0.2 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 4 | 10.56 ms | 7.00 ms | 0.66× [0.66, 0.67] | -50.9% | [-51.5%, -50.2%] | 0.4 pts | 1.1 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 4 | 10.55 ms | 18.58 ms | 1.76× [1.74, 1.78] | +43.3% | [+42.7%, +44.0%] | 0.4 pts | 1.6 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 8 | 58.5 | 55.1 | 0.95× [0.92, 0.97] | -5.6% | [-8.3%, -2.9%] | 1.7 pts | 2.8 | no | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 8 | 61.0 | 241 | 3.96× [3.91, 4.00] | +74.7% | [+74.4%, +75.0%] | 0.3 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 4637 | 4611 | 0.99× [0.99, 1.00] | -0.6% | [-0.9%, -0.2%] | 0.5 pts | 1.5 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 8 | 4635 | 8372 | 1.82× [1.82, 1.83] | +45.2% | [+45.0%, +45.4%] | 1.6 pts | 5.2 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 8 | 94.3 | 79.2 | 0.84× [0.83, 0.86] | -19.0% | [-21.1%, -16.9%] | 1.3 pts | 1.5 | no | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 8 | 108 | 262 | 2.41× [2.33, 2.49] | +58.4% | [+57.0%, +59.9%] | 3.6 pts | 3.2 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 8 | 34.47 ms | 29.05 ms | 0.84× [0.84, 0.84] | -18.9% | [-19.2%, -18.6%] | 0.7 pts | 2.5 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 8 | 34.57 ms | 89.08 ms | 2.58× [2.56, 2.60] | +61.2% | [+60.9%, +61.5%] | 0.4 pts | 3.4 | yes | yes |
| natural | u64 | 65536 | valuesFor | ordered | baseline | 8 | 89.1 | 82.2 | 0.94× [0.90, 0.97] | -6.7% | [-10.8%, -2.6%] | 3.5 pts | 1.0 | no | yes |
| natural | u64 | 65536 | valuesFor | ordered | btree-sets | 8 | 109 | 389 | 3.59× [3.32, 3.91] | +72.2% | [+69.9%, +74.4%] | 1.6 pts | 1.7 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 5169 | 5462 | 1.05× [1.05, 1.06] | +5.2% | [+4.7%, +5.7%] | 0.3 pts | 0.4 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | btree-sets | 8 | 5366 | 13.0 µs | 2.49× [2.32, 2.67] | +59.8% | [+57.0%, +62.6%] | 1.8 pts | 4.3 | yes | yes |
| natural | u64 | 65536 | churn | ordered | baseline | 8 | 198 | 171 | 0.83× [0.81, 0.86] | -20.2% | [-24.0%, -16.3%] | 7.0 pts | 0.9 | no | yes |
| natural | u64 | 65536 | churn | ordered | btree-sets | 8 | 253 | 475 | 1.83× [1.72, 1.95] | +45.3% | [+41.8%, +48.8%] | 2.5 pts | 11.5 | yes | yes |
| natural | u64 | 65536 | build | ordered | baseline | 8 | 250.79 ms | 219.50 ms | 0.88× [0.84, 0.93] | -13.4% | [-19.3%, -7.5%] | 3.6 pts | 10.5 | no | yes |
| natural | u64 | 65536 | build | ordered | btree-sets | 8 | 251.11 ms | 560.04 ms | 2.21× [2.14, 2.28] | +54.7% | [+53.2%, +56.2%] | 1.0 pts | 4.7 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 4 | 123 | 118 | 0.95× [0.94, 0.96] | -5.0% | [-6.1%, -3.9%] | 0.7 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 4 | 124 | 145 | 1.16× [1.14, 1.18] | +13.8% | [+11.9%, +15.6%] | 1.1 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 4 | 1830 | 3039 | 1.66× [1.62, 1.70] | +39.6% | [+38.1%, +41.1%] | 1.0 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 4 | 1785 | 611 | 0.34× [0.34, 0.35] | -190.7% | [-195.1%, -186.2%] | 2.8 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 4 | 2186 | 3844 | 1.74× [1.68, 1.79] | +42.4% | [+40.6%, +44.2%] | 1.1 pts | 0.6 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | btree-map | 4 | 2167 | 673 | 0.31× [0.30, 0.32] | -220.4% | [-228.0%, -212.9%] | 4.7 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 4 | 251 | 261 | 1.05× [1.04, 1.05] | +4.5% | [+4.2%, +4.7%] | 0.2 pts | 0.6 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 4 | 252 | 217 | 0.86× [0.85, 0.87] | -16.4% | [-17.8%, -14.9%] | 0.9 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 4 | 4.02 ms | 3.69 ms | 0.92× [0.91, 0.93] | -8.7% | [-9.6%, -7.8%] | 0.6 pts | 1.3 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 4 | 4.02 ms | 3.14 ms | 0.78× [0.77, 0.79] | -27.7% | [-29.2%, -26.2%] | 0.9 pts | 1.7 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 4 | 152 | 150 | 0.99× [0.98, 1.00] | -0.8% | [-2.0%, +0.5%] | 0.8 pts | 1.5 | yes | no |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 4 | 153 | 204 | 1.33× [1.30, 1.36] | +24.8% | [+22.9%, +26.7%] | 1.2 pts | 1.3 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 4 | 1885 | 3358 | 1.79× [1.78, 1.80] | +44.1% | [+43.7%, +44.4%] | 0.2 pts | 0.4 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 4 | 1863 | 743 | 0.40× [0.39, 0.41] | -150.1% | [-154.0%, -146.1%] | 2.5 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 4 | 7765 | 16.3 µs | 2.08× [2.07, 2.08] | +51.8% | [+51.7%, +51.9%] | 0.1 pts | 0.3 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | btree-map | 4 | 7878 | 2302 | 0.30× [0.29, 0.31] | -236.6% | [-245.5%, -227.8%] | 5.6 pts | 0.3 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 4 | 291 | 323 | 1.11× [1.10, 1.12] | +9.7% | [+9.0%, +10.5%] | 0.5 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 4 | 290 | 278 | 0.96× [0.95, 0.98] | -3.9% | [-5.3%, -2.4%] | 0.9 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 4 | 17.57 ms | 17.06 ms | 0.97× [0.97, 0.98] | -2.6% | [-3.4%, -1.8%] | 0.5 pts | 1.7 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 4 | 17.51 ms | 15.52 ms | 0.89× [0.88, 0.89] | -12.9% | [-13.9%, -11.8%] | 0.6 pts | 2.0 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 188 | 204 | 1.09× [1.07, 1.11] | +8.4% | [+6.7%, +10.1%] | 1.1 pts | 1.5 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 195 | 270 | 1.39× [1.36, 1.41] | +27.8% | [+26.6%, +29.1%] | 1.0 pts | 0.8 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 2032 | 3690 | 1.83× [1.80, 1.87] | +45.4% | [+44.5%, +46.4%] | 0.9 pts | 1.3 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 2011 | 984 | 0.49× [0.47, 0.52] | -103.5% | [-114.9%, -92.1%] | 10.9 pts | 1.9 | no | yes |
| single-value | dirs | 65536 | prefix | ordered | baseline | 8 | 29.4 µs | 64.6 µs | 2.25× [2.24, 2.26] | +55.6% | [+55.3%, +55.8%] | 0.2 pts | 1.0 | yes | yes |
| single-value | dirs | 65536 | prefix | ordered | btree-map | 8 | 28.1 µs | 8398 | 0.30× [0.27, 0.34] | -231.8% | [-269.8%, -193.8%] | 24.5 pts | 14.0 | no | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 371 | 452 | 1.21× [1.20, 1.22] | +17.4% | [+17.0%, +17.8%] | 1.3 pts | 1.0 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | btree-map | 8 | 405 | 464 | 1.15× [1.14, 1.16] | +13.3% | [+12.6%, +14.0%] | 1.3 pts | 1.9 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 84.79 ms | 89.48 ms | 1.06× [1.05, 1.06] | +5.2% | [+5.1%, +5.4%] | 0.2 pts | 1.5 | yes | yes |
| single-value | dirs | 65536 | build | ordered | btree-map | 8 | 84.72 ms | 84.89 ms | 1.00× [0.99, 1.02] | +0.4% | [-1.0%, +1.8%] | 1.1 pts | 7.8 | yes | no |
| single-value | street | 4096 | valuesFor | ordered | baseline | 4 | 79.9 | 68.8 | 0.86× [0.86, 0.86] | -16.6% | [-16.9%, -16.2%] | 0.2 pts | 0.4 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 4 | 80.5 | 127 | 1.58× [1.54, 1.62] | +36.6% | [+35.0%, +38.3%] | 1.0 pts | 1.1 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 4 | 1393 | 2735 | 1.97× [1.94, 2.00] | +49.2% | [+48.5%, +50.0%] | 0.5 pts | 0.9 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 4 | 1363 | 590 | 0.43× [0.43, 0.44] | -131.1% | [-132.9%, -129.2%] | 1.2 pts | 0.6 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 4 | 266 | 307 | 1.16× [1.15, 1.16] | +13.5% | [+13.1%, +13.9%] | 0.3 pts | 0.6 | yes | yes |
| single-value | street | 4096 | prefix | ordered | btree-map | 4 | 264 | 175 | 0.66× [0.65, 0.67] | -52.0% | [-54.2%, -49.7%] | 1.4 pts | 0.8 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 4 | 162 | 163 | 1.01× [1.00, 1.01] | +0.9% | [+0.3%, +1.4%] | 0.4 pts | 1.0 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 4 | 163 | 180 | 1.10× [1.09, 1.12] | +9.4% | [+8.5%, +10.3%] | 0.6 pts | 1.1 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 4 | 2.56 ms | 2.35 ms | 0.92× [0.91, 0.92] | -8.9% | [-9.8%, -8.1%] | 0.5 pts | 1.5 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 4 | 2.57 ms | 2.70 ms | 1.05× [1.03, 1.07] | +4.8% | [+3.3%, +6.3%] | 0.9 pts | 1.7 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 95.9 | 89.7 | 0.93× [0.92, 0.95] | -7.0% | [-8.4%, -5.6%] | 1.1 pts | 3.2 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 96.2 | 171 | 1.77× [1.76, 1.78] | +43.5% | [+43.1%, +43.9%] | 0.9 pts | 2.9 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1520 | 2929 | 1.93× [1.92, 1.93] | +48.1% | [+47.9%, +48.2%] | 0.1 pts | 0.4 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1510 | 670 | 0.44× [0.44, 0.45] | -125.2% | [-126.2%, -124.1%] | 2.3 pts | 2.3 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 8 | 561 | 989 | 1.79× [1.77, 1.80] | +44.0% | [+43.6%, +44.4%] | 0.3 pts | 0.4 | yes | yes |
| single-value | street | 16384 | prefix | ordered | btree-map | 8 | 560 | 363 | 0.65× [0.65, 0.65] | -54.3% | [-54.9%, -53.8%] | 1.5 pts | 1.0 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 191 | 204 | 1.07× [1.04, 1.11] | +6.8% | [+3.9%, +9.7%] | 1.9 pts | 6.9 | no | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 191 | 228 | 1.20× [1.18, 1.22] | +16.5% | [+15.1%, +17.8%] | 1.2 pts | 2.0 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 11.79 ms | 11.05 ms | 0.94× [0.94, 0.95] | -6.3% | [-6.8%, -5.8%] | 0.6 pts | 2.2 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 11.75 ms | 13.50 ms | 1.15× [1.14, 1.16] | +12.9% | [+12.2%, +13.6%] | 0.6 pts | 2.2 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 118 | 120 | 1.02× [1.00, 1.05] | +2.2% | [+0.1%, +4.4%] | 1.4 pts | 3.8 | no | yes |
| single-value | street | 65536 | valuesFor | ordered | btree-map | 8 | 121 | 222 | 1.85× [1.83, 1.87] | +46.0% | [+45.5%, +46.4%] | 0.5 pts | 1.7 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 1564 | 3099 | 2.00× [1.98, 2.02] | +50.0% | [+49.6%, +50.5%] | 1.2 pts | 3.2 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | btree-map | 8 | 1555 | 785 | 0.50× [0.49, 0.51] | -99.6% | [-103.1%, -96.1%] | 4.4 pts | 1.7 | yes | yes |
| single-value | street | 65536 | prefix | ordered | baseline | 8 | 1700 | 3984 | 2.36× [2.32, 2.41] | +57.7% | [+57.0%, +58.5%] | 0.7 pts | 0.6 | yes | yes |
| single-value | street | 65536 | prefix | ordered | btree-map | 8 | 1720 | 1142 | 0.67× [0.63, 0.72] | -48.8% | [-58.0%, -39.7%] | 5.7 pts | 2.0 | no | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 239 | 298 | 1.24× [1.23, 1.25] | +19.4% | [+18.7%, +20.1%] | 1.0 pts | 1.0 | yes | yes |
| single-value | street | 65536 | churn | ordered | btree-map | 8 | 254 | 320 | 1.27× [1.27, 1.28] | +21.5% | [+21.1%, +21.8%] | 0.6 pts | 1.0 | yes | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 54.30 ms | 57.65 ms | 1.07× [1.06, 1.07] | +6.2% | [+5.6%, +6.9%] | 0.9 pts | 5.0 | yes | yes |
| single-value | street | 65536 | build | ordered | btree-map | 8 | 54.16 ms | 65.41 ms | 1.21× [1.20, 1.21] | +17.1% | [+16.6%, +17.5%] | 0.3 pts | 2.8 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 4 | 40.0 | 23.5 | 0.59× [0.58, 0.59] | -70.4% | [-71.3%, -69.5%] | 0.6 pts | 2.7 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 4 | 40.1 | 108 | 2.69× [2.62, 2.77] | +62.9% | [+61.9%, +63.8%] | 0.6 pts | 12.4 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 4 | 996 | 1545 | 1.56× [1.55, 1.56] | +35.7% | [+35.5%, +35.9%] | 0.1 pts | 1.4 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 4 | 995 | 441 | 0.44× [0.44, 0.44] | -125.5% | [-125.8%, -125.1%] | 0.2 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 4 | 82.5 | 67.6 | 0.82× [0.81, 0.83] | -22.1% | [-23.4%, -20.8%] | 0.8 pts | 2.0 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 4 | 83.0 | 152 | 1.84× [1.83, 1.85] | +45.7% | [+45.5%, +45.8%] | 0.1 pts | 0.4 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 4 | 1.25 ms | 1.17 ms | 0.93× [0.92, 0.94] | -7.4% | [-8.9%, -6.0%] | 0.9 pts | 1.0 | yes | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 4 | 1.24 ms | 2.21 ms | 1.77× [1.74, 1.81] | +43.6% | [+42.6%, +44.7%] | 0.7 pts | 1.7 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 4 | 30.7 | 28.5 | 0.93× [0.92, 0.94] | -7.7% | [-8.8%, -6.5%] | 0.7 pts | 3.6 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 4 | 30.9 | 145 | 4.70× [4.67, 4.74] | +78.7% | [+78.6%, +78.9%] | 0.1 pts | 2.2 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 4 | 2004 | 2150 | 1.08× [1.06, 1.09] | +7.0% | [+5.9%, +8.2%] | 0.7 pts | 1.4 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 4 | 1940 | 500 | 0.26× [0.26, 0.26] | -287.2% | [-292.1%, -282.3%] | 3.1 pts | 4.3 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 4 | 74.0 | 68.2 | 0.92× [0.91, 0.93] | -8.6% | [-9.4%, -7.7%] | 0.6 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 4 | 73.9 | 194 | 2.65× [2.60, 2.70] | +62.3% | [+61.6%, +62.9%] | 0.4 pts | 1.9 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 4 | 5.88 ms | 4.97 ms | 0.85× [0.84, 0.86] | -17.6% | [-19.1%, -16.2%] | 0.9 pts | 0.8 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 4 | 5.94 ms | 11.14 ms | 1.89× [1.86, 1.91] | +47.0% | [+46.3%, +47.7%] | 0.5 pts | 1.1 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | baseline | 8 | 38.1 | 28.1 | 0.76× [0.72, 0.81] | -31.5% | [-39.3%, -23.8%] | 6.6 pts | 4.0 | no | yes |
| single-value | u64 | 65536 | valuesFor | ordered | btree-map | 8 | 38.8 | 189 | 4.84× [4.62, 5.09] | +79.4% | [+78.4%, +80.3%] | 0.6 pts | 4.9 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 2074 | 2568 | 1.23× [1.22, 1.24] | +18.9% | [+18.2%, +19.6%] | 0.8 pts | 4.0 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | btree-map | 8 | 2060 | 562 | 0.27× [0.27, 0.28] | -266.6% | [-270.3%, -262.9%] | 2.5 pts | 3.2 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | baseline | 8 | 121 | 121 | 1.01× [0.99, 1.02] | +0.6% | [-0.9%, +2.2%] | 1.4 pts | 1.7 | yes | no |
| single-value | u64 | 65536 | churn | ordered | btree-map | 8 | 131 | 262 | 2.01× [1.97, 2.06] | +50.3% | [+49.1%, +51.4%] | 0.8 pts | 1.4 | yes | yes |
| single-value | u64 | 65536 | build | ordered | baseline | 8 | 24.47 ms | 24.13 ms | 0.99× [0.97, 1.00] | -1.5% | [-3.1%, +0.1%] | 1.1 pts | 1.6 | yes | no |
| single-value | u64 | 65536 | build | ordered | btree-map | 8 | 24.40 ms | 57.26 ms | 2.34× [2.31, 2.37] | +57.3% | [+56.8%, +57.9%] | 0.7 pts | 2.5 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of -8.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 churn: ordered vs baseline: the processes scatter 8.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 8.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.69% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.69% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 prefix: ordered vs baseline: the A/A validations found a systematic difference of +4.49% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +2.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -1.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 prefix: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=16384 valuesFor: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=16384 build: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +0.98% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs baseline: the pooled difference of -1.57% does not clear the 2.28% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 churn: ordered vs baseline: the pooled interval [-4.99%, 1.85%] includes zero
- natural street n=65536 churn: ordered vs btree-sets: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs btree-sets: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.46% does not clear the 0.60% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.21%, 0.29%] includes zero
- natural u64 n=4096 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 build: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 build: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +2.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 churn: ordered vs btree-sets: the processes scatter 11.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 build: ordered vs baseline: the A/A validations found a systematic difference of +0.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 build: ordered vs baseline: the processes scatter 10.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 build: ordered vs btree-sets: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.76% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.96%, 0.45%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of -0.08% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=65536 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=65536 prefix: ordered vs baseline: the A/A validations found a systematic difference of +6.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=65536 prefix: ordered vs btree-map: the processes scatter 14.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: the pooled interval [-0.97%, 1.80%] includes zero
- single-value dirs n=65536 build: ordered vs btree-map: the processes scatter 7.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -1.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.68% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 churn: ordered vs baseline: the processes scatter 6.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.19% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesBetween: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 prefix: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs baseline: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 12.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 churn: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesBetween: ordered vs btree-map: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs btree-map: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesBetween: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesBetween: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs baseline: the pooled interval [-0.95%, 2.15%] includes zero
- single-value u64 n=65536 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=65536 build: ordered vs baseline: the pooled interval [-3.08%, 0.13%] includes zero
- single-value u64 n=65536 build: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
