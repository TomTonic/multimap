| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 4 | 133 | 129 | 0.97× [0.95, 0.98] | -3.3% | [-4.9%, -1.8%] | 1.0 pts | 1.5 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 4 | 134 | 208 | 1.55× [1.51, 1.58] | +35.3% | [+33.8%, +36.8%] | 0.9 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 4 | 3801 | 4068 | 1.07× [1.06, 1.09] | +6.6% | [+5.3%, +7.9%] | 0.8 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 4 | 3784 | 6298 | 1.66× [1.66, 1.67] | +39.9% | [+39.8%, +40.0%] | 0.1 pts | 0.1 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 4 | 4443 | 4714 | 1.06× [1.04, 1.08] | +5.6% | [+3.6%, +7.6%] | 1.2 pts | 0.4 | yes | no |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 4 | 4394 | 8456 | 1.89× [1.81, 1.97] | +47.0% | [+44.8%, +49.2%] | 1.4 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 4 | 250 | 203 | 0.81× [0.80, 0.82] | -23.4% | [-24.5%, -22.4%] | 0.7 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 4 | 247 | 275 | 1.10× [1.08, 1.12] | +9.1% | [+7.6%, +10.7%] | 1.0 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 4 | 11.93 ms | 9.40 ms | 0.79× [0.79, 0.79] | -27.0% | [-27.3%, -26.7%] | 0.2 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 4 | 11.93 ms | 13.07 ms | 1.09× [1.08, 1.10] | +8.6% | [+7.6%, +9.5%] | 0.6 pts | 3.5 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 4 | 167 | 162 | 0.97× [0.96, 0.99] | -2.8% | [-4.4%, -1.2%] | 1.0 pts | 1.8 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 4 | 175 | 280 | 1.60× [1.52, 1.68] | +37.4% | [+34.4%, +40.4%] | 1.9 pts | 1.5 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 4 | 4225 | 4509 | 1.06× [1.05, 1.08] | +6.1% | [+5.2%, +7.0%] | 0.6 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 4 | 4247 | 6918 | 1.63× [1.60, 1.66] | +38.7% | [+37.6%, +39.8%] | 0.7 pts | 1.2 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 4 | 20.3 µs | 22.3 µs | 1.13× [1.11, 1.14] | +11.1% | [+10.3%, +12.0%] | 0.5 pts | 0.6 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 4 | 19.7 µs | 35.4 µs | 1.77× [1.76, 1.78] | +43.5% | [+43.2%, +43.8%] | 0.2 pts | 0.6 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 4 | 311 | 262 | 0.84× [0.84, 0.85] | -18.9% | [-19.7%, -18.1%] | 0.5 pts | 1.6 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 4 | 349 | 392 | 1.12× [1.09, 1.14] | +10.4% | [+8.7%, +12.1%] | 1.1 pts | 1.0 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 4 | 57.90 ms | 46.53 ms | 0.81× [0.80, 0.81] | -24.1% | [-24.7%, -23.5%] | 0.4 pts | 1.7 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 4 | 57.78 ms | 66.56 ms | 1.15× [1.14, 1.16] | +13.1% | [+12.4%, +13.7%] | 0.4 pts | 3.5 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 249 | 257 | 1.04× [1.02, 1.07] | +4.0% | [+1.5%, +6.5%] | 2.1 pts | 0.7 | no | yes |
| natural | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 288 | 508 | 1.77× [1.71, 1.83] | +43.5% | [+41.5%, +45.5%] | 2.4 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 5255 | 5847 | 1.11× [1.09, 1.12] | +9.5% | [+8.0%, +11.0%] | 1.6 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 5637 | 11.1 µs | 1.99× [1.97, 2.01] | +49.8% | [+49.2%, +50.3%] | 0.4 pts | 0.6 | yes | yes |
| natural | dirs | 65536 | prefix | ordered | baseline | 8 | 83.3 µs | 92.7 µs | 1.14× [1.12, 1.15] | +12.1% | [+10.8%, +13.4%] | 1.0 pts | 0.6 | yes | yes |
| natural | dirs | 65536 | prefix | ordered | btree-sets | 8 | 86.3 µs | 200.9 µs | 2.34× [2.32, 2.36] | +57.2% | [+56.8%, +57.6%] | 0.6 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 512 | 437 | 0.82× [0.82, 0.83] | -21.3% | [-21.9%, -20.7%] | 3.1 pts | 0.4 | yes | yes |
| natural | dirs | 65536 | churn | ordered | btree-sets | 8 | 593 | 697 | 1.17× [1.16, 1.18] | +14.7% | [+13.8%, +15.6%] | 0.6 pts | 0.8 | yes | yes |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 358.92 ms | 308.28 ms | 0.86× [0.86, 0.86] | -16.4% | [-16.7%, -16.1%] | 0.4 pts | 0.7 | yes | yes |
| natural | dirs | 65536 | build | ordered | btree-sets | 8 | 358.89 ms | 471.03 ms | 1.31× [1.31, 1.31] | +23.7% | [+23.6%, +23.7%] | 0.3 pts | 0.6 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 6 | 84.8 | 76.6 | 0.90× [0.89, 0.91] | -10.6% | [-11.8%, -9.5%] | 0.8 pts | 1.7 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 6 | 85.6 | 180 | 2.10× [2.09, 2.12] | +52.4% | [+52.1%, +52.8%] | 0.3 pts | 0.8 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 6 | 3066 | 3480 | 1.13× [1.12, 1.15] | +11.9% | [+10.9%, +12.8%] | 0.6 pts | 0.4 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 3029 | 5333 | 1.77× [1.75, 1.79] | +43.5% | [+42.9%, +44.0%] | 0.4 pts | 0.9 | yes | yes |
| natural | street | 4096 | prefix | ordered | baseline | 6 | 296 | 338 | 1.15× [1.13, 1.17] | +13.0% | [+11.4%, +14.7%] | 1.1 pts | 0.8 | yes | yes |
| natural | street | 4096 | prefix | ordered | btree-sets | 6 | 298 | 650 | 2.18× [2.16, 2.20] | +54.1% | [+53.7%, +54.5%] | 0.3 pts | 0.6 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 6 | 183 | 135 | 0.74× [0.73, 0.74] | -35.7% | [-37.0%, -34.4%] | 0.8 pts | 1.3 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 6 | 181 | 239 | 1.31× [1.30, 1.32] | +23.4% | [+22.9%, +24.0%] | 0.6 pts | 0.9 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 6 | 7.06 ms | 5.07 ms | 0.72× [0.71, 0.73] | -38.8% | [-40.0%, -37.6%] | 0.9 pts | 2.5 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 6 | 7.06 ms | 9.25 ms | 1.31× [1.30, 1.33] | +23.8% | [+23.1%, +24.5%] | 0.4 pts | 1.6 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 106 | 99.2 | 0.93× [0.93, 0.94] | -7.1% | [-7.8%, -6.5%] | 1.0 pts | 1.8 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 109 | 235 | 2.15× [2.14, 2.16] | +53.4% | [+53.2%, +53.6%] | 0.3 pts | 0.8 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 3349 | 3738 | 1.11× [1.11, 1.12] | +10.3% | [+9.7%, +10.8%] | 0.4 pts | 0.7 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 3426 | 5827 | 1.72× [1.70, 1.74] | +42.0% | [+41.3%, +42.6%] | 0.9 pts | 2.1 | yes | yes |
| natural | street | 16384 | prefix | ordered | baseline | 8 | 1102 | 1243 | 1.12× [1.10, 1.15] | +10.8% | [+8.8%, +12.7%] | 1.5 pts | 0.9 | yes | yes |
| natural | street | 16384 | prefix | ordered | btree-sets | 8 | 1113 | 2282 | 2.04× [2.03, 2.05] | +51.0% | [+50.8%, +51.3%] | 0.6 pts | 1.7 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 221 | 175 | 0.79× [0.79, 0.80] | -25.9% | [-26.6%, -25.2%] | 0.5 pts | 0.9 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 234 | 311 | 1.32× [1.30, 1.35] | +24.4% | [+23.2%, +25.7%] | 0.8 pts | 0.9 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 33.96 ms | 25.78 ms | 0.76× [0.76, 0.76] | -31.4% | [-31.6%, -31.2%] | 0.6 pts | 1.4 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 33.81 ms | 46.72 ms | 1.38× [1.37, 1.39] | +27.6% | [+27.2%, +28.0%] | 0.5 pts | 4.9 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 153 | 150 | 0.98× [0.97, 1.00] | -1.6% | [-3.3%, +0.1%] | 1.6 pts | 1.0 | yes | no |
| natural | street | 65536 | valuesFor | ordered | btree-sets | 8 | 185 | 414 | 2.26× [2.19, 2.34] | +55.8% | [+54.4%, +57.3%] | 2.1 pts | 1.3 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 3901 | 4503 | 1.16× [1.13, 1.18] | +13.6% | [+11.8%, +15.3%] | 1.1 pts | 1.0 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 4182 | 9531 | 2.28× [2.26, 2.30] | +56.1% | [+55.8%, +56.5%] | 0.6 pts | 1.2 | yes | yes |
| natural | street | 65536 | prefix | ordered | baseline | 8 | 4631 | 5407 | 1.17× [1.13, 1.21] | +14.6% | [+11.8%, +17.4%] | 2.2 pts | 0.7 | no | yes |
| natural | street | 65536 | prefix | ordered | btree-sets | 8 | 4856 | 11.4 µs | 2.39× [2.28, 2.51] | +58.2% | [+56.2%, +60.2%] | 1.2 pts | 1.0 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 360 | 309 | 0.85× [0.82, 0.87] | -17.9% | [-21.4%, -14.5%] | 3.0 pts | 0.7 | no | yes |
| natural | street | 65536 | churn | ordered | btree-sets | 8 | 390 | 532 | 1.33× [1.29, 1.37] | +24.7% | [+22.4%, +27.1%] | 1.5 pts | 1.2 | yes | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 204.91 ms | 169.21 ms | 0.82× [0.82, 0.83] | -21.6% | [-22.5%, -20.7%] | 0.7 pts | 3.2 | yes | yes |
| natural | street | 65536 | build | ordered | btree-sets | 8 | 203.89 ms | 286.36 ms | 1.40× [1.39, 1.42] | +28.8% | [+28.2%, +29.4%] | 0.7 pts | 3.7 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 4 | 50.9 | 50.5 | 1.00× [0.99, 1.01] | -0.2% | [-1.1%, +0.7%] | 0.6 pts | 1.0 | yes | no |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 4 | 51.5 | 194 | 3.76× [3.73, 3.78] | +73.4% | [+73.2%, +73.6%] | 0.1 pts | 0.8 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 4 | 3952 | 3983 | 1.01× [1.00, 1.02] | +1.1% | [-0.1%, +2.2%] | 0.7 pts | 0.9 | yes | no |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 4 | 3956 | 8272 | 2.09× [2.06, 2.13] | +52.2% | [+51.4%, +53.0%] | 0.5 pts | 2.0 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 4 | 72.2 | 69.5 | 0.96× [0.95, 0.97] | -4.0% | [-4.9%, -3.0%] | 0.6 pts | 2.0 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 4 | 71.7 | 187 | 2.61× [2.58, 2.63] | +61.6% | [+61.3%, +62.0%] | 0.2 pts | 1.6 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 4 | 7.52 ms | 7.01 ms | 0.93× [0.92, 0.94] | -7.7% | [-8.5%, -6.8%] | 0.5 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 4 | 7.54 ms | 18.30 ms | 2.43× [2.42, 2.44] | +58.8% | [+58.6%, +59.0%] | 0.1 pts | 0.8 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 4 | 56.3 | 54.9 | 0.98× [0.97, 0.99] | -2.3% | [-3.4%, -1.2%] | 0.7 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 4 | 58.8 | 238 | 4.03× [3.94, 4.13] | +75.2% | [+74.6%, +75.8%] | 0.4 pts | 1.9 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 4 | 4630 | 4641 | 1.01× [0.99, 1.02] | +0.7% | [-0.8%, +2.2%] | 0.9 pts | 2.0 | yes | no |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 4 | 4588 | 8420 | 1.84× [1.81, 1.86] | +45.6% | [+44.7%, +46.4%] | 0.5 pts | 1.8 | yes | yes |
| natural | u64 | 16384 | churn | ordered | baseline | 4 | 84.5 | 80.4 | 0.95× [0.94, 0.97] | -4.7% | [-6.6%, -2.8%] | 1.2 pts | 1.7 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 4 | 97.2 | 261 | 2.78× [2.56, 3.05] | +64.0% | [+60.9%, +67.2%] | 2.0 pts | 1.7 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 4 | 30.71 ms | 28.93 ms | 0.94× [0.93, 0.95] | -6.2% | [-7.2%, -5.1%] | 0.7 pts | 1.9 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 4 | 30.78 ms | 88.56 ms | 2.88× [2.85, 2.91] | +65.3% | [+65.0%, +65.7%] | 0.2 pts | 2.0 | yes | yes |
| natural | u64 | 65536 | valuesFor | ordered | baseline | 8 | 90.0 | 84.8 | 0.94× [0.89, 0.98] | -7.0% | [-11.9%, -2.0%] | 3.1 pts | 0.7 | no | yes |
| natural | u64 | 65536 | valuesFor | ordered | btree-sets | 8 | 109 | 391 | 3.55× [3.42, 3.69] | +71.8% | [+70.8%, +72.9%] | 1.0 pts | 1.0 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 5303 | 5436 | 1.03× [1.01, 1.05] | +2.6% | [+0.8%, +4.4%] | 1.1 pts | 1.2 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | btree-sets | 8 | 5517 | 12.9 µs | 2.39× [2.30, 2.48] | +58.1% | [+56.5%, +59.7%] | 1.3 pts | 2.4 | yes | yes |
| natural | u64 | 65536 | churn | ordered | baseline | 8 | 187 | 172 | 0.90× [0.86, 0.94] | -11.0% | [-16.1%, -6.0%] | 5.2 pts | 0.7 | no | yes |
| natural | u64 | 65536 | churn | ordered | btree-sets | 8 | 243 | 485 | 1.94× [1.80, 2.11] | +48.5% | [+44.5%, +52.5%] | 2.5 pts | 6.2 | yes | yes |
| natural | u64 | 65536 | build | ordered | baseline | 8 | 233.77 ms | 223.66 ms | 0.95× [0.94, 0.96] | -5.0% | [-6.3%, -3.7%] | 1.0 pts | 0.9 | yes | yes |
| natural | u64 | 65536 | build | ordered | btree-sets | 8 | 232.67 ms | 562.31 ms | 2.38× [2.25, 2.52] | +58.0% | [+55.6%, +60.4%] | 1.6 pts | 3.8 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 6 | 127 | 117 | 0.92× [0.92, 0.93] | -8.1% | [-8.6%, -7.7%] | 0.4 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 6 | 128 | 148 | 1.16× [1.14, 1.19] | +13.9% | [+11.9%, +15.8%] | 1.3 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 1770 | 3007 | 1.71× [1.67, 1.75] | +41.5% | [+40.3%, +42.7%] | 0.8 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 6 | 1742 | 636 | 0.36× [0.36, 0.37] | -174.0% | [-180.5%, -167.6%] | 4.5 pts | 1.3 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | baseline | 6 | 2031 | 3757 | 1.85× [1.83, 1.88] | +46.1% | [+45.2%, +46.9%] | 0.7 pts | 0.4 | yes | yes |
| single-value | dirs | 4096 | prefix | ordered | btree-map | 6 | 2006 | 725 | 0.36× [0.36, 0.37] | -176.6% | [-180.8%, -172.5%] | 2.5 pts | 0.4 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 6 | 300 | 264 | 0.88× [0.87, 0.89] | -13.6% | [-14.4%, -12.8%] | 0.6 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 6 | 300 | 222 | 0.73× [0.72, 0.74] | -37.1% | [-38.6%, -35.5%] | 1.2 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 6 | 4.34 ms | 3.70 ms | 0.85× [0.85, 0.86] | -17.4% | [-18.1%, -16.7%] | 0.7 pts | 1.2 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 6 | 4.34 ms | 3.14 ms | 0.72× [0.72, 0.72] | -38.6% | [-39.2%, -38.0%] | 0.6 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 4 | 156 | 150 | 0.96× [0.96, 0.97] | -3.7% | [-4.6%, -2.8%] | 0.6 pts | 1.6 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 4 | 157 | 209 | 1.33× [1.29, 1.37] | +24.7% | [+22.3%, +27.1%] | 1.5 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 4 | 1851 | 3344 | 1.82× [1.79, 1.85] | +45.1% | [+44.2%, +45.9%] | 0.5 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 4 | 1837 | 777 | 0.42× [0.42, 0.43] | -135.9% | [-140.2%, -131.7%] | 2.7 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | baseline | 4 | 7497 | 16.3 µs | 2.17× [2.16, 2.17] | +53.8% | [+53.6%, +54.0%] | 0.1 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | prefix | ordered | btree-map | 4 | 6984 | 2367 | 0.32× [0.31, 0.33] | -213.9% | [-223.2%, -204.5%] | 5.9 pts | 0.5 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 4 | 331 | 326 | 0.98× [0.97, 0.99] | -2.1% | [-3.2%, -0.9%] | 0.7 pts | 1.8 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 4 | 335 | 285 | 0.85× [0.84, 0.86] | -17.7% | [-19.2%, -16.2%] | 0.9 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 4 | 18.88 ms | 17.08 ms | 0.91× [0.90, 0.91] | -10.3% | [-10.9%, -9.7%] | 0.4 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 4 | 18.91 ms | 15.53 ms | 0.82× [0.81, 0.83] | -21.8% | [-23.1%, -20.5%] | 0.8 pts | 2.4 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 192 | 207 | 1.10× [1.00, 1.23] | +9.2% | [-0.3%, +18.7%] | 6.7 pts | 9.3 | no | no |
| single-value | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 208 | 294 | 1.42× [1.37, 1.47] | +29.6% | [+27.0%, +32.2%] | 2.2 pts | 1.1 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1993 | 3691 | 1.87× [1.82, 1.92] | +46.4% | [+44.9%, +47.9%] | 1.2 pts | 1.4 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 1975 | 1078 | 0.55× [0.52, 0.57] | -83.3% | [-91.6%, -75.0%] | 10.9 pts | 1.9 | yes | yes |
| single-value | dirs | 65536 | prefix | ordered | baseline | 8 | 28.4 µs | 65.8 µs | 2.31× [2.30, 2.32] | +56.7% | [+56.5%, +56.8%] | 0.4 pts | 2.1 | yes | yes |
| single-value | dirs | 65536 | prefix | ordered | btree-map | 8 | 26.8 µs | 9089 | 0.32× [0.32, 0.33] | -210.3% | [-217.4%, -203.1%] | 11.0 pts | 2.4 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 425 | 460 | 1.10× [1.06, 1.14] | +9.1% | [+5.7%, +12.5%] | 2.1 pts | 1.1 | no | yes |
| single-value | dirs | 65536 | churn | ordered | btree-map | 8 | 468 | 487 | 1.04× [1.03, 1.04] | +3.7% | [+3.1%, +4.2%] | 0.9 pts | 1.6 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 90.98 ms | 90.05 ms | 0.99× [0.99, 0.99] | -0.9% | [-1.2%, -0.6%] | 0.3 pts | 1.7 | yes | yes |
| single-value | dirs | 65536 | build | ordered | btree-map | 8 | 90.93 ms | 85.99 ms | 0.95× [0.93, 0.96] | -5.6% | [-7.3%, -3.9%] | 1.5 pts | 10.3 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 84.3 | 68.7 | 0.82× [0.81, 0.82] | -22.6% | [-23.0%, -22.2%] | 0.6 pts | 0.9 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 84.5 | 128 | 1.49× [1.47, 1.52] | +33.1% | [+31.9%, +34.3%] | 1.0 pts | 1.8 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 1347 | 2698 | 2.00× [1.98, 2.03] | +50.1% | [+49.4%, +50.8%] | 0.6 pts | 1.3 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1335 | 623 | 0.47× [0.47, 0.47] | -113.8% | [-114.9%, -112.6%] | 1.5 pts | 0.9 | yes | yes |
| single-value | street | 4096 | prefix | ordered | baseline | 8 | 253 | 304 | 1.20× [1.19, 1.21] | +16.6% | [+15.9%, +17.2%] | 0.4 pts | 0.8 | yes | yes |
| single-value | street | 4096 | prefix | ordered | btree-map | 8 | 252 | 181 | 0.71× [0.70, 0.72] | -40.5% | [-42.0%, -38.9%] | 2.1 pts | 1.0 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 206 | 165 | 0.80× [0.79, 0.80] | -25.5% | [-26.4%, -24.5%] | 0.9 pts | 1.1 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 207 | 181 | 0.87× [0.86, 0.87] | -15.5% | [-16.5%, -14.5%] | 1.2 pts | 1.3 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.86 ms | 2.36 ms | 0.82× [0.81, 0.83] | -21.5% | [-23.0%, -20.0%] | 1.5 pts | 2.5 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.85 ms | 2.73 ms | 0.95× [0.94, 0.97] | -4.9% | [-6.7%, -3.0%] | 1.3 pts | 2.2 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 99.9 | 89.5 | 0.89× [0.88, 0.90] | -11.8% | [-13.1%, -10.5%] | 0.8 pts | 2.1 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 101 | 172 | 1.69× [1.66, 1.72] | +40.8% | [+39.7%, +41.8%] | 0.9 pts | 2.1 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1484 | 2929 | 1.97× [1.96, 1.99] | +49.3% | [+48.9%, +49.8%] | 0.3 pts | 1.0 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1475 | 703 | 0.47× [0.47, 0.48] | -111.0% | [-113.1%, -108.9%] | 1.6 pts | 1.8 | yes | yes |
| single-value | street | 16384 | prefix | ordered | baseline | 8 | 542 | 991 | 1.84× [1.83, 1.85] | +45.6% | [+45.2%, +46.1%] | 0.4 pts | 0.6 | yes | yes |
| single-value | street | 16384 | prefix | ordered | btree-map | 8 | 543 | 378 | 0.70× [0.69, 0.70] | -43.2% | [-44.3%, -42.1%] | 1.8 pts | 1.1 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 225 | 208 | 0.94× [0.92, 0.96] | -6.8% | [-9.3%, -4.4%] | 2.0 pts | 2.8 | no | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 226 | 231 | 1.03× [1.02, 1.04] | +2.6% | [+1.6%, +3.7%] | 0.8 pts | 0.9 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 12.74 ms | 11.05 ms | 0.87× [0.86, 0.88] | -15.1% | [-16.2%, -14.0%] | 1.1 pts | 3.7 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 12.74 ms | 13.46 ms | 1.06× [1.05, 1.06] | +5.5% | [+5.1%, +6.0%] | 0.3 pts | 0.7 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 122 | 120 | 0.99× [0.95, 1.02] | -1.3% | [-5.0%, +2.4%] | 2.3 pts | 6.7 | no | no |
| single-value | street | 65536 | valuesFor | ordered | btree-map | 8 | 126 | 226 | 1.79× [1.75, 1.84] | +44.2% | [+43.0%, +45.5%] | 1.0 pts | 1.7 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 1525 | 3083 | 2.03× [2.01, 2.05] | +50.7% | [+50.2%, +51.2%] | 0.5 pts | 1.2 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | btree-map | 8 | 1497 | 789 | 0.53× [0.51, 0.55] | -87.4% | [-94.4%, -80.4%] | 5.9 pts | 3.2 | yes | yes |
| single-value | street | 65536 | prefix | ordered | baseline | 8 | 1638 | 3994 | 2.42× [2.39, 2.45] | +58.7% | [+58.1%, +59.2%] | 0.5 pts | 0.6 | yes | yes |
| single-value | street | 65536 | prefix | ordered | btree-map | 8 | 1637 | 1179 | 0.71× [0.70, 0.72] | -40.2% | [-42.4%, -38.0%] | 2.0 pts | 0.9 | yes | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 266 | 297 | 1.11× [1.10, 1.12] | +10.0% | [+9.1%, +10.9%] | 0.8 pts | 0.7 | yes | yes |
| single-value | street | 65536 | churn | ordered | btree-map | 8 | 289 | 332 | 1.14× [1.13, 1.15] | +12.3% | [+11.3%, +13.3%] | 1.3 pts | 1.8 | yes | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 58.69 ms | 57.70 ms | 0.98× [0.98, 0.99] | -1.7% | [-2.4%, -1.0%] | 0.8 pts | 2.9 | yes | yes |
| single-value | street | 65536 | build | ordered | btree-map | 8 | 58.89 ms | 65.39 ms | 1.11× [1.11, 1.11] | +9.9% | [+9.6%, +10.2%] | 0.3 pts | 1.4 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 4 | 46.2 | 23.5 | 0.51× [0.50, 0.51] | -96.7% | [-98.4%, -95.1%] | 1.0 pts | 4.4 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 4 | 46.1 | 109 | 2.37× [2.35, 2.38] | +57.7% | [+57.4%, +58.0%] | 0.2 pts | 2.2 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 4 | 945 | 1530 | 1.62× [1.61, 1.63] | +38.1% | [+37.8%, +38.5%] | 0.2 pts | 2.3 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 4 | 946 | 471 | 0.50× [0.50, 0.50] | -100.3% | [-100.8%, -99.8%] | 0.3 pts | 1.8 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 4 | 75.4 | 67.7 | 0.90× [0.89, 0.90] | -11.5% | [-11.9%, -11.0%] | 0.3 pts | 1.0 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 4 | 76.5 | 153 | 2.00× [1.98, 2.02] | +50.0% | [+49.5%, +50.6%] | 0.3 pts | 0.8 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 4 | 1.15 ms | 1.16 ms | 1.01× [1.00, 1.02] | +1.3% | [+0.3%, +2.2%] | 0.6 pts | 0.8 | yes | no |
| single-value | u64 | 4096 | build | ordered | btree-map | 4 | 1.14 ms | 2.23 ms | 1.94× [1.92, 1.97] | +48.6% | [+47.9%, +49.2%] | 0.4 pts | 1.3 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 4 | 31.5 | 28.5 | 0.91× [0.90, 0.92] | -10.5% | [-11.7%, -9.2%] | 0.8 pts | 3.6 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 4 | 31.6 | 146 | 4.60× [4.58, 4.62] | +78.3% | [+78.2%, +78.4%] | 0.1 pts | 1.7 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 4 | 1929 | 2106 | 1.09× [1.09, 1.10] | +8.5% | [+8.3%, +8.8%] | 0.1 pts | 0.4 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 4 | 1890 | 529 | 0.28× [0.28, 0.28] | -257.0% | [-260.2%, -253.8%] | 2.0 pts | 2.4 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 4 | 79.8 | 69.0 | 0.87× [0.86, 0.87] | -15.2% | [-15.9%, -14.5%] | 0.4 pts | 0.7 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 4 | 78.7 | 195 | 2.49× [2.48, 2.51] | +59.9% | [+59.6%, +60.1%] | 0.2 pts | 0.7 | yes | yes |
| single-value | u64 | 16384 | build | ordered | baseline | 4 | 6.01 ms | 4.93 ms | 0.82× [0.81, 0.83] | -21.9% | [-23.4%, -20.4%] | 0.9 pts | 1.6 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 4 | 6.00 ms | 11.16 ms | 1.87× [1.85, 1.88] | +46.4% | [+45.9%, +46.8%] | 0.3 pts | 0.8 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | baseline | 8 | 38.4 | 28.5 | 0.75× [0.71, 0.78] | -34.0% | [-39.9%, -28.0%] | 4.3 pts | 1.6 | no | yes |
| single-value | u64 | 65536 | valuesFor | ordered | btree-map | 8 | 39.3 | 191 | 4.74× [4.39, 5.16] | +78.9% | [+77.2%, +80.6%] | 1.1 pts | 6.5 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 2004 | 2521 | 1.26× [1.25, 1.26] | +20.6% | [+20.2%, +20.9%] | 0.7 pts | 4.3 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | btree-map | 8 | 1995 | 587 | 0.29× [0.29, 0.30] | -240.7% | [-243.9%, -237.5%] | 2.6 pts | 3.0 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | baseline | 8 | 127 | 124 | 0.99× [0.95, 1.03] | -1.0% | [-4.8%, +2.8%] | 2.5 pts | 2.3 | no | no |
| single-value | u64 | 65536 | churn | ordered | btree-map | 8 | 140 | 266 | 1.93× [1.90, 1.96] | +48.1% | [+47.3%, +48.9%] | 1.0 pts | 1.7 | yes | yes |
| single-value | u64 | 65536 | build | ordered | baseline | 8 | 25.72 ms | 24.03 ms | 0.94× [0.92, 0.96] | -6.8% | [-9.0%, -4.7%] | 1.3 pts | 1.5 | no | yes |
| single-value | u64 | 65536 | build | ordered | btree-map | 8 | 25.78 ms | 57.27 ms | 2.23× [2.19, 2.28] | +55.2% | [+54.4%, +56.1%] | 0.6 pts | 1.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 prefix: ordered vs baseline: the pooled difference of 5.60% does not clear the 10.51% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=4096 prefix: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 build: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +9.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 prefix: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +3.01% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 prefix: ordered vs baseline: the A/A validations found a systematic difference of -4.59% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +2.69% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of +2.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.67% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.12% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.15% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 build: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs btree-sets: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: the pooled difference of -1.63% does not clear the 1.64% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.91% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 valuesFor: ordered vs baseline: the pooled interval [-3.33%, 0.08%] includes zero
- natural street n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +2.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 build: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs btree-sets: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.19% does not clear the 0.25% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.09%, 0.71%] includes zero
- natural u64 n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.12%, 2.24%] includes zero
- natural u64 n=4096 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=4096 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.67% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.82%, 2.16%] includes zero
- natural u64 n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=16384 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural u64 n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 churn: ordered vs btree-sets: the processes scatter 6.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 build: ordered vs btree-sets: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -1.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of -0.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 prefix: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 build: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesFor: ordered vs baseline: the pooled interval [-0.26%, 18.65%] includes zero
- single-value dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 9.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=65536 prefix: ordered vs baseline: the A/A validations found a systematic difference of +8.12% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=65536 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 prefix: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=65536 build: ordered vs btree-map: the processes scatter 10.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -1.52% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 build: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 churn: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs baseline: the pooled interval [-5.04%, 2.41%] includes zero
- single-value street n=65536 valuesFor: ordered vs baseline: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 7 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=65536 valuesBetween: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 build: ordered vs baseline: the pooled difference of 1.27% does not clear the 1.33% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value u64 n=65536 valuesFor: ordered vs btree-map: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesBetween: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesBetween: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs baseline: the pooled interval [-4.84%, 2.77%] includes zero
- single-value u64 n=65536 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
