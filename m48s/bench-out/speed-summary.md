| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 183 | 166 | 0.91× [0.90, 0.92] | -9.9% | [-11.0%, -8.8%] | 2.1 pts | 2.4 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 187 | 222 | 1.17× [1.15, 1.20] | +14.8% | [+12.9%, +16.8%] | 2.0 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 7484 | 7321 | 0.98× [0.97, 1.00] | -1.6% | [-2.9%, -0.3%] | 1.0 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 7594 | 12.6 µs | 1.65× [1.61, 1.69] | +39.4% | [+38.0%, +40.8%] | 1.3 pts | 2.0 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 9154 | 9508 | 1.02× [1.00, 1.04] | +2.0% | [+0.1%, +3.8%] | 2.0 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 9723 | 18.4 µs | 1.89× [1.87, 1.90] | +47.0% | [+46.6%, +47.3%] | 0.6 pts | 0.8 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 281 | 238 | 0.85× [0.84, 0.86] | -17.6% | [-18.4%, -16.8%] | 1.5 pts | 2.1 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 281 | 291 | 1.04× [1.03, 1.05] | +4.0% | [+2.8%, +5.1%] | 0.9 pts | 0.7 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 13.93 ms | 10.59 ms | 0.76× [0.76, 0.77] | -30.8% | [-31.5%, -30.2%] | 0.6 pts | 1.1 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 13.87 ms | 13.96 ms | 1.01× [1.00, 1.01] | +0.9% | [+0.4%, +1.3%] | 0.5 pts | 1.2 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 223 | 212 | 0.96× [0.95, 0.96] | -4.7% | [-5.7%, -3.7%] | 0.8 pts | 1.2 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 235 | 318 | 1.38× [1.37, 1.38] | +27.4% | [+27.1%, +27.7%] | 3.2 pts | 2.7 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 7910 | 8182 | 1.04× [1.02, 1.07] | +4.1% | [+2.0%, +6.2%] | 1.4 pts | 1.6 | no | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 8137 | 15.5 µs | 1.90× [1.85, 1.94] | +47.3% | [+46.0%, +48.6%] | 0.9 pts | 0.9 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 8 | 35.5 µs | 38.7 µs | 1.08× [1.07, 1.09] | +7.5% | [+6.7%, +8.3%] | 1.2 pts | 1.2 | yes | no |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 8 | 32.6 µs | 66.6 µs | 2.02× [1.99, 2.05] | +50.5% | [+49.6%, +51.3%] | 0.5 pts | 0.6 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 8 | 340 | 308 | 0.90× [0.89, 0.91] | -11.0% | [-11.9%, -10.1%] | 1.4 pts | 2.3 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 8 | 359 | 411 | 1.15× [1.13, 1.16] | +12.8% | [+11.5%, +14.1%] | 1.1 pts | 1.4 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 8 | 65.20 ms | 53.76 ms | 0.83× [0.82, 0.83] | -21.0% | [-21.5%, -20.4%] | 1.0 pts | 2.6 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 8 | 65.33 ms | 73.02 ms | 1.12× [1.11, 1.13] | +10.6% | [+10.0%, +11.1%] | 0.4 pts | 1.9 | yes | yes |
| natural-str | dirs | 65536 | valuesFor | ordered | baseline | 8 | 310 | 351 | 1.14× [1.12, 1.16] | +12.3% | [+10.6%, +13.9%] | 2.7 pts | 1.1 | yes | yes |
| natural-str | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 346 | 599 | 1.74× [1.68, 1.80] | +42.5% | [+40.6%, +44.3%] | 1.6 pts | 1.0 | yes | yes |
| natural-str | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 9050 | 10.0 µs | 1.12× [1.11, 1.12] | +10.5% | [+10.1%, +10.9%] | 1.3 pts | 1.2 | yes | yes |
| natural-str | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 9395 | 21.7 µs | 2.36× [2.34, 2.39] | +57.7% | [+57.3%, +58.1%] | 1.3 pts | 1.5 | yes | yes |
| natural-str | dirs | 65536 | prefix | ordered | baseline | 8 | 161.9 µs | 192.4 µs | 1.21× [1.19, 1.22] | +17.2% | [+16.2%, +18.3%] | 1.3 pts | 1.2 | yes | yes |
| natural-str | dirs | 65536 | prefix | ordered | btree-sets | 8 | 168.3 µs | 459.1 µs | 2.78× [2.74, 2.81] | +64.0% | [+63.5%, +64.5%] | 1.4 pts | 3.5 | yes | yes |
| natural-str | dirs | 65536 | churn | ordered | baseline | 8 | 546 | 514 | 0.91× [0.89, 0.93] | -10.3% | [-12.5%, -8.1%] | 2.1 pts | 0.3 | no | yes |
| natural-str | dirs | 65536 | churn | ordered | btree-sets | 8 | 581 | 738 | 1.28× [1.24, 1.32] | +21.9% | [+19.7%, +24.0%] | 2.6 pts | 2.7 | yes | yes |
| natural-str | dirs | 65536 | build | ordered | baseline | 8 | 393.38 ms | 372.05 ms | 0.95× [0.93, 0.96] | -5.8% | [-7.1%, -4.4%] | 3.3 pts | 3.6 | yes | yes |
| natural-str | dirs | 65536 | build | ordered | btree-sets | 8 | 391.47 ms | 514.74 ms | 1.31× [1.30, 1.31] | +23.5% | [+23.2%, +23.9%] | 0.4 pts | 1.2 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | baseline | 4 | 129 | 107 | 0.83× [0.82, 0.84] | -20.5% | [-21.6%, -19.4%] | 0.7 pts | 0.9 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 4 | 131 | 188 | 1.43× [1.38, 1.49] | +30.2% | [+27.7%, +32.7%] | 1.6 pts | 2.1 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 4 | 5841 | 6109 | 1.04× [1.02, 1.06] | +3.8% | [+2.0%, +5.7%] | 1.2 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 4 | 5827 | 9131 | 1.56× [1.53, 1.59] | +36.0% | [+34.8%, +37.1%] | 0.7 pts | 1.5 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 4 | 661 | 572 | 0.88× [0.87, 0.89] | -13.5% | [-14.5%, -12.5%] | 0.6 pts | 0.3 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 4 | 667 | 931 | 1.40× [1.37, 1.42] | +28.5% | [+27.2%, +29.8%] | 0.8 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 4 | 206 | 151 | 0.73× [0.73, 0.73] | -37.1% | [-37.4%, -36.7%] | 0.2 pts | 0.4 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 4 | 207 | 250 | 1.21× [1.20, 1.22] | +17.1% | [+16.4%, +17.9%] | 0.5 pts | 0.6 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 4 | 8.32 ms | 5.71 ms | 0.68× [0.68, 0.69] | -46.0% | [-47.6%, -44.5%] | 1.0 pts | 2.4 | yes | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 4 | 8.39 ms | 9.98 ms | 1.19× [1.18, 1.21] | +16.1% | [+15.2%, +17.1%] | 0.6 pts | 1.8 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 8 | 149 | 132 | 0.88× [0.86, 0.90] | -13.7% | [-15.7%, -11.6%] | 1.8 pts | 3.3 | no | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 8 | 153 | 243 | 1.59× [1.57, 1.61] | +37.0% | [+36.2%, +37.9%] | 1.5 pts | 3.4 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 6119 | 6482 | 1.06× [1.05, 1.07] | +5.7% | [+4.8%, +6.5%] | 0.8 pts | 1.5 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 6173 | 10.0 µs | 1.63× [1.58, 1.68] | +38.6% | [+36.8%, +40.5%] | 1.6 pts | 2.4 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 8 | 2131 | 2222 | 1.04× [1.02, 1.05] | +3.5% | [+1.8%, +5.1%] | 1.5 pts | 0.9 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 8 | 2189 | 3711 | 1.74× [1.72, 1.77] | +42.5% | [+41.7%, +43.3%] | 1.1 pts | 0.8 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 8 | 243 | 200 | 0.85× [0.80, 0.89] | -18.2% | [-24.3%, -12.0%] | 5.1 pts | 8.9 | no | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 8 | 255 | 329 | 1.31× [1.29, 1.34] | +23.7% | [+22.2%, +25.1%] | 1.7 pts | 2.0 | yes | yes |
| natural-str | street | 16384 | build | ordered | baseline | 8 | 38.35 ms | 28.47 ms | 0.75× [0.75, 0.75] | -33.6% | [-33.8%, -33.4%] | 2.9 pts | 7.7 | yes | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 8 | 38.44 ms | 49.99 ms | 1.30× [1.30, 1.31] | +23.3% | [+22.8%, +23.8%] | 0.5 pts | 3.4 | yes | yes |
| natural-str | street | 65536 | valuesFor | ordered | baseline | 8 | 192 | 215 | 1.15× [1.06, 1.25] | +12.8% | [+5.6%, +20.1%] | 5.1 pts | 2.9 | no | yes |
| natural-str | street | 65536 | valuesFor | ordered | btree-sets | 8 | 214 | 438 | 2.11× [2.00, 2.24] | +52.6% | [+49.9%, +55.3%] | 2.3 pts | 2.0 | yes | yes |
| natural-str | street | 65536 | valuesBetween | ordered | baseline | 8 | 6596 | 7672 | 1.16× [1.12, 1.21] | +13.8% | [+10.4%, +17.2%] | 2.4 pts | 2.2 | no | yes |
| natural-str | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 6860 | 15.7 µs | 2.34× [2.19, 2.51] | +57.2% | [+54.3%, +60.2%] | 2.1 pts | 3.4 | yes | yes |
| natural-str | street | 65536 | prefix | ordered | baseline | 8 | 8093 | 9637 | 1.20× [1.16, 1.23] | +16.4% | [+14.1%, +18.8%] | 1.6 pts | 2.1 | no | yes |
| natural-str | street | 65536 | prefix | ordered | btree-sets | 8 | 8146 | 17.8 µs | 2.26× [2.12, 2.42] | +55.8% | [+52.8%, +58.8%] | 2.2 pts | 5.2 | yes | yes |
| natural-str | street | 65536 | churn | ordered | baseline | 8 | 371 | 375 | 0.96× [0.94, 0.98] | -4.2% | [-6.8%, -1.5%] | 2.7 pts | 0.5 | no | yes |
| natural-str | street | 65536 | churn | ordered | btree-sets | 8 | 398 | 561 | 1.40× [1.34, 1.47] | +28.6% | [+25.1%, +32.0%] | 2.3 pts | 1.9 | no | yes |
| natural-str | street | 65536 | build | ordered | baseline | 8 | 217.31 ms | 193.72 ms | 0.90× [0.86, 0.94] | -11.2% | [-15.7%, -6.7%] | 2.9 pts | 8.0 | no | yes |
| natural-str | street | 65536 | build | ordered | btree-sets | 8 | 219.49 ms | 318.30 ms | 1.45× [1.42, 1.48] | +31.0% | [+29.5%, +32.4%] | 1.0 pts | 2.6 | yes | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 145 | 140 | 0.97× [0.96, 0.99] | -2.9% | [-4.3%, -1.4%] | 1.0 pts | 0.9 | yes | yes |
| single-value-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 146 | 155 | 1.05× [1.02, 1.08] | +4.7% | [+1.6%, +7.8%] | 3.0 pts | 1.5 | no | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 3910 | 5076 | 1.31× [1.29, 1.32] | +23.4% | [+22.4%, +24.3%] | 0.7 pts | 1.0 | yes | yes |
| single-value-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 3918 | 2061 | 0.53× [0.52, 0.54] | -89.6% | [-93.0%, -86.1%] | 3.1 pts | 1.1 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | baseline | 8 | 4674 | 6469 | 1.35× [1.32, 1.39] | +26.1% | [+24.0%, +28.2%] | 1.8 pts | 0.6 | yes | yes |
| single-value-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 4631 | 2538 | 0.56× [0.54, 0.58] | -78.9% | [-84.9%, -72.9%] | 4.2 pts | 0.6 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | baseline | 8 | 278 | 263 | 0.96× [0.95, 0.97] | -4.6% | [-5.7%, -3.5%] | 0.8 pts | 1.3 | yes | yes |
| single-value-str | dirs | 4096 | churn | ordered | btree-map | 8 | 277 | 229 | 0.83× [0.82, 0.83] | -20.9% | [-21.9%, -20.0%] | 0.6 pts | 0.6 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | baseline | 8 | 4.53 ms | 3.81 ms | 0.84× [0.84, 0.85] | -18.6% | [-19.7%, -17.6%] | 0.8 pts | 1.2 | yes | yes |
| single-value-str | dirs | 4096 | build | ordered | btree-map | 8 | 4.57 ms | 3.31 ms | 0.73× [0.72, 0.73] | -37.7% | [-39.0%, -36.3%] | 1.2 pts | 1.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesFor | ordered | baseline | 4 | 175 | 175 | 1.00× [0.98, 1.01] | -0.2% | [-1.7%, +1.2%] | 0.9 pts | 1.4 | yes | no |
| single-value-str | dirs | 16384 | valuesFor | ordered | btree-map | 4 | 177 | 213 | 1.19× [1.17, 1.21] | +15.9% | [+14.6%, +17.2%] | 0.8 pts | 0.8 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | baseline | 4 | 4057 | 5515 | 1.36× [1.35, 1.38] | +26.7% | [+26.0%, +27.5%] | 0.5 pts | 1.7 | yes | yes |
| single-value-str | dirs | 16384 | valuesBetween | ordered | btree-map | 4 | 4054 | 2297 | 0.57× [0.56, 0.57] | -76.9% | [-79.5%, -74.3%] | 1.6 pts | 0.8 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | baseline | 4 | 18.3 µs | 27.2 µs | 1.47× [1.46, 1.48] | +32.2% | [+31.7%, +32.6%] | 0.3 pts | 0.5 | yes | yes |
| single-value-str | dirs | 16384 | prefix | ordered | btree-map | 4 | 16.4 µs | 9209 | 0.56× [0.55, 0.56] | -80.1% | [-82.7%, -77.4%] | 1.7 pts | 1.4 | yes | yes |
| single-value-str | dirs | 16384 | churn | ordered | baseline | 4 | 326 | 328 | 1.01× [1.00, 1.02] | +0.7% | [-0.2%, +1.6%] | 0.6 pts | 1.7 | yes | no |
| single-value-str | dirs | 16384 | churn | ordered | btree-map | 4 | 327 | 298 | 0.91× [0.91, 0.92] | -9.6% | [-10.4%, -8.9%] | 0.5 pts | 0.4 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | baseline | 4 | 19.98 ms | 17.63 ms | 0.88× [0.88, 0.89] | -13.0% | [-14.0%, -12.1%] | 0.6 pts | 1.3 | yes | yes |
| single-value-str | dirs | 16384 | build | ordered | btree-map | 4 | 20.07 ms | 16.29 ms | 0.81× [0.80, 0.82] | -23.2% | [-24.9%, -21.4%] | 1.1 pts | 2.6 | yes | yes |
| single-value-str | dirs | 65536 | valuesFor | ordered | baseline | 8 | 253 | 321 | 1.26× [1.17, 1.36] | +20.6% | [+14.6%, +26.7%] | 4.6 pts | 3.5 | no | yes |
| single-value-str | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 247 | 324 | 1.35× [1.30, 1.39] | +25.7% | [+23.0%, +28.3%] | 2.9 pts | 1.3 | no | yes |
| single-value-str | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 4580 | 6747 | 1.50× [1.48, 1.52] | +33.4% | [+32.6%, +34.3%] | 1.2 pts | 1.7 | yes | yes |
| single-value-str | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 4342 | 3022 | 0.70× [0.69, 0.72] | -41.9% | [-45.4%, -38.5%] | 5.3 pts | 2.0 | yes | yes |
| single-value-str | dirs | 65536 | prefix | ordered | baseline | 8 | 73.1 µs | 120.8 µs | 1.66× [1.65, 1.68] | +39.8% | [+39.3%, +40.4%] | 0.9 pts | 1.3 | yes | yes |
| single-value-str | dirs | 65536 | prefix | ordered | btree-map | 8 | 72.3 µs | 41.4 µs | 0.59× [0.57, 0.60] | -70.7% | [-75.2%, -66.1%] | 5.9 pts | 1.9 | yes | yes |
| single-value-str | dirs | 65536 | churn | ordered | baseline | 8 | 482 | 545 | 1.13× [1.09, 1.17] | +11.6% | [+8.5%, +14.7%] | 3.4 pts | 1.2 | no | yes |
| single-value-str | dirs | 65536 | churn | ordered | btree-map | 8 | 488 | 527 | 1.09× [1.07, 1.10] | +7.9% | [+6.7%, +9.0%] | 1.3 pts | 1.7 | yes | yes |
| single-value-str | dirs | 65536 | build | ordered | baseline | 8 | 107.55 ms | 105.71 ms | 0.98× [0.98, 0.99] | -1.7% | [-2.4%, -1.0%] | 2.8 pts | 4.6 | yes | yes |
| single-value-str | dirs | 65536 | build | ordered | btree-map | 8 | 100.52 ms | 94.89 ms | 0.94× [0.93, 0.96] | -5.9% | [-7.6%, -4.2%] | 2.0 pts | 2.5 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | baseline | 8 | 99.7 | 90.9 | 0.91× [0.91, 0.92] | -9.4% | [-10.2%, -8.6%] | 0.7 pts | 1.5 | yes | yes |
| single-value-str | street | 4096 | valuesFor | ordered | btree-map | 8 | 102 | 132 | 1.28× [1.25, 1.32] | +22.2% | [+20.0%, +24.3%] | 1.5 pts | 1.4 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 3348 | 4651 | 1.39× [1.38, 1.39] | +27.8% | [+27.5%, +28.2%] | 0.6 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | valuesBetween | ordered | btree-map | 8 | 3278 | 1627 | 0.50× [0.49, 0.50] | -101.5% | [-104.2%, -98.9%] | 1.9 pts | 1.7 | yes | yes |
| single-value-str | street | 4096 | prefix | ordered | baseline | 8 | 485 | 482 | 0.99× [0.98, 1.00] | -1.0% | [-2.3%, +0.2%] | 2.1 pts | 1.1 | yes | no |
| single-value-str | street | 4096 | prefix | ordered | btree-map | 8 | 476 | 260 | 0.54× [0.51, 0.58] | -84.4% | [-95.0%, -73.9%] | 7.5 pts | 1.5 | no | yes |
| single-value-str | street | 4096 | churn | ordered | baseline | 8 | 188 | 165 | 0.88× [0.87, 0.89] | -13.2% | [-14.3%, -12.1%] | 1.0 pts | 1.8 | yes | yes |
| single-value-str | street | 4096 | churn | ordered | btree-map | 8 | 188 | 188 | 1.00× [1.00, 1.01] | +0.3% | [-0.4%, +1.0%] | 0.8 pts | 2.1 | yes | no |
| single-value-str | street | 4096 | build | ordered | baseline | 8 | 2.98 ms | 2.33 ms | 0.78× [0.77, 0.79] | -28.0% | [-29.0%, -27.0%] | 0.6 pts | 0.9 | yes | yes |
| single-value-str | street | 4096 | build | ordered | btree-map | 8 | 3.01 ms | 2.83 ms | 0.95× [0.94, 0.95] | -5.4% | [-6.2%, -4.7%] | 1.1 pts | 1.3 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | baseline | 6 | 118 | 113 | 0.96× [0.95, 0.97] | -4.3% | [-5.4%, -3.1%] | 0.7 pts | 1.8 | yes | yes |
| single-value-str | street | 16384 | valuesFor | ordered | btree-map | 6 | 118 | 178 | 1.50× [1.48, 1.53] | +33.6% | [+32.3%, +34.8%] | 1.2 pts | 3.6 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 3567 | 4962 | 1.39× [1.38, 1.40] | +28.2% | [+27.6%, +28.7%] | 0.4 pts | 0.8 | yes | yes |
| single-value-str | street | 16384 | valuesBetween | ordered | btree-map | 6 | 3504 | 1757 | 0.50× [0.49, 0.51] | -99.7% | [-103.2%, -96.2%] | 2.1 pts | 1.9 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | baseline | 6 | 1271 | 1714 | 1.35× [1.30, 1.40] | +25.9% | [+23.3%, +28.4%] | 1.7 pts | 0.9 | yes | yes |
| single-value-str | street | 16384 | prefix | ordered | btree-map | 6 | 1259 | 757 | 0.60× [0.60, 0.60] | -66.6% | [-67.8%, -65.4%] | 1.3 pts | 0.6 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | baseline | 6 | 217 | 203 | 0.95× [0.93, 0.96] | -5.7% | [-7.6%, -3.8%] | 1.3 pts | 3.4 | yes | yes |
| single-value-str | street | 16384 | churn | ordered | btree-map | 6 | 217 | 241 | 1.11× [1.10, 1.13] | +10.1% | [+9.1%, +11.1%] | 0.8 pts | 1.3 | yes | yes |
| single-value-str | street | 16384 | build | ordered | baseline | 6 | 13.51 ms | 11.01 ms | 0.82× [0.81, 0.82] | -22.5% | [-23.7%, -21.3%] | 0.7 pts | 1.2 | yes | yes |
| single-value-str | street | 16384 | build | ordered | btree-map | 6 | 13.47 ms | 14.18 ms | 1.05× [1.04, 1.06] | +4.9% | [+4.1%, +5.7%] | 0.6 pts | 2.0 | yes | yes |
| single-value-str | street | 65536 | valuesFor | ordered | baseline | 8 | 143 | 156 | 1.11× [1.02, 1.21] | +9.7% | [+2.0%, +17.5%] | 5.1 pts | 8.4 | no | yes |
| single-value-str | street | 65536 | valuesFor | ordered | btree-map | 8 | 145 | 233 | 1.61× [1.60, 1.63] | +38.0% | [+37.4%, +38.7%] | 0.7 pts | 1.6 | yes | yes |
| single-value-str | street | 65536 | valuesBetween | ordered | baseline | 8 | 3664 | 5402 | 1.47× [1.45, 1.50] | +32.1% | [+31.0%, +33.2%] | 0.9 pts | 1.1 | yes | yes |
| single-value-str | street | 65536 | valuesBetween | ordered | btree-map | 8 | 3606 | 1925 | 0.54× [0.51, 0.57] | -85.3% | [-96.3%, -74.4%] | 7.8 pts | 5.0 | no | yes |
| single-value-str | street | 65536 | prefix | ordered | baseline | 8 | 4351 | 7015 | 1.62× [1.60, 1.65] | +38.4% | [+37.5%, +39.4%] | 1.1 pts | 0.6 | yes | yes |
| single-value-str | street | 65536 | prefix | ordered | btree-map | 8 | 4315 | 2723 | 0.63× [0.61, 0.65] | -58.0% | [-63.3%, -52.7%] | 3.5 pts | 1.1 | yes | yes |
| single-value-str | street | 65536 | churn | ordered | baseline | 8 | 271 | 299 | 1.08× [1.07, 1.09] | +7.6% | [+6.8%, +8.3%] | 0.6 pts | 0.6 | yes | yes |
| single-value-str | street | 65536 | churn | ordered | btree-map | 8 | 306 | 363 | 1.20× [1.18, 1.22] | +16.7% | [+15.6%, +17.8%] | 0.9 pts | 1.1 | yes | yes |
| single-value-str | street | 65536 | build | ordered | baseline | 8 | 63.05 ms | 58.41 ms | 0.93× [0.92, 0.93] | -7.9% | [-8.8%, -7.0%] | 0.8 pts | 2.9 | yes | yes |
| single-value-str | street | 65536 | build | ordered | btree-map | 8 | 62.99 ms | 69.71 ms | 1.11× [1.11, 1.11] | +9.9% | [+9.8%, +10.1%] | 0.4 pts | 1.8 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 prefix: ordered vs baseline: the pooled difference of 7.49% does not clear the 14.70% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=16384 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 build: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=65536 prefix: ordered vs baseline: the A/A validations found a systematic difference of -3.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=65536 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -4.87% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=65536 prefix: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -1.07% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=65536 churn: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=65536 build: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=65536 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.52% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 prefix: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=4096 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=4096 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 churn: ordered vs baseline: the processes scatter 8.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 churn: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 build: ordered vs baseline: the processes scatter 7.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 build: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=65536 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=65536 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=65536 prefix: ordered vs btree-sets: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.81% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=65536 build: ordered vs baseline: the processes scatter 8.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=65536 build: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -1.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.49% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.87% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.24% does not clear the 0.61% noise floor, the bound on what the harness reports between identical code in every process
- single-value-str dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.66%, 1.18%] includes zero
- single-value-str dirs n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 churn: ordered vs baseline: the pooled difference of 0.72% does not clear the 1.63% noise floor, the bound on what the harness reports between identical code in every process
- single-value-str dirs n=16384 churn: ordered vs baseline: the pooled interval [-0.20%, 1.64%] includes zero
- single-value-str dirs n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value-str dirs n=16384 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.10% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str dirs n=16384 build: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=65536 valuesBetween: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=65536 build: ordered vs baseline: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str dirs n=65536 build: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value-str dirs n=65536 build: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.85% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=4096 prefix: ordered vs baseline: the pooled interval [-2.32%, 0.24%] includes zero
- single-value-str street n=4096 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -1.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=4096 churn: ordered vs btree-map: the pooled interval [-0.41%, 0.96%] includes zero
- single-value-str street n=4096 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 churn: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=16384 build: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=65536 valuesFor: ordered vs baseline: the processes scatter 8.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=65536 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=65536 valuesBetween: ordered vs btree-map: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=65536 prefix: ordered vs baseline: the A/A validations found a systematic difference of -1.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value-str street n=65536 build: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value-str street n=65536 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.07% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
