| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 4 | 140 | 119 | 0.85× [0.84, 0.86] | -17.0% | [-18.4%, -15.6%] | 0.9 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 4 | 142 | 213 | 1.48× [1.43, 1.54] | +32.5% | [+30.2%, +34.9%] | 1.5 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 4 | 2167 | 3568 | 1.67× [1.62, 1.72] | +40.0% | [+38.2%, +41.8%] | 1.1 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 4 | 2151 | 6269 | 2.92× [2.89, 2.95] | +65.8% | [+65.4%, +66.1%] | 0.2 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 4 | 246 | 190 | 0.78× [0.77, 0.79] | -28.9% | [-30.7%, -27.0%] | 1.2 pts | 1.3 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 4 | 244 | 277 | 1.15× [1.14, 1.16] | +13.1% | [+12.6%, +13.5%] | 0.3 pts | 0.2 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 4 | 11.84 ms | 8.49 ms | 0.72× [0.71, 0.72] | -39.7% | [-41.3%, -38.1%] | 1.0 pts | 1.8 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 4 | 11.86 ms | 13.22 ms | 1.12× [1.10, 1.13] | +10.4% | [+9.4%, +11.4%] | 0.6 pts | 2.2 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 174 | 152 | 0.88× [0.86, 0.90] | -13.8% | [-16.9%, -10.6%] | 2.1 pts | 4.9 | no | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 179 | 280 | 1.58× [1.54, 1.61] | +36.5% | [+35.2%, +37.8%] | 1.0 pts | 1.0 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 2280 | 4054 | 1.78× [1.73, 1.83] | +43.8% | [+42.1%, +45.5%] | 1.2 pts | 2.1 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 2295 | 6875 | 3.00× [2.98, 3.01] | +66.6% | [+66.5%, +66.8%] | 0.2 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 297 | 265 | 0.90× [0.86, 0.94] | -11.2% | [-15.9%, -6.5%] | 4.6 pts | 6.3 | no | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 313 | 391 | 1.26× [1.23, 1.30] | +20.7% | [+18.7%, +22.8%] | 1.9 pts | 2.8 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 55.46 ms | 43.96 ms | 0.79× [0.79, 0.80] | -25.8% | [-26.0%, -25.6%] | 1.1 pts | 5.2 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 55.53 ms | 67.13 ms | 1.21× [1.20, 1.21] | +17.3% | [+17.0%, +17.6%] | 0.2 pts | 2.2 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 238 | 271 | 1.16× [1.15, 1.18] | +14.1% | [+13.2%, +14.9%] | 3.4 pts | 1.5 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 260 | 517 | 2.00× [1.94, 2.07] | +50.1% | [+48.4%, +51.7%] | 2.6 pts | 1.6 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 2598 | 5547 | 2.15× [2.12, 2.18] | +53.4% | [+52.7%, +54.1%] | 0.6 pts | 0.7 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 2741 | 11.1 µs | 4.06× [4.05, 4.08] | +75.4% | [+75.3%, +75.5%] | 0.4 pts | 1.1 | yes | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 490 | 494 | 1.00× [0.97, 1.04] | +0.1% | [-3.5%, +3.6%] | 8.2 pts | 1.2 | no | no |
| natural | dirs | 65536 | churn | ordered | btree-sets | 8 | 526 | 722 | 1.39× [1.37, 1.42] | +28.3% | [+26.9%, +29.7%] | 1.9 pts | 1.6 | yes | yes |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 324.83 ms | 319.02 ms | 1.01× [0.96, 1.08] | +1.4% | [-4.5%, +7.4%] | 5.3 pts | 7.5 | no | no |
| natural | dirs | 65536 | build | ordered | btree-sets | 8 | 323.84 ms | 479.22 ms | 1.48× [1.47, 1.48] | +32.3% | [+32.1%, +32.6%] | 0.6 pts | 1.3 | yes | yes |
| natural | links | 4096 | valuesFor | ordered | baseline | 8 | 254 | 250 | 0.99× [0.98, 1.00] | -1.3% | [-2.4%, -0.2%] | 0.7 pts | 0.5 | yes | yes |
| natural | links | 4096 | valuesFor | ordered | btree-sets | 8 | 261 | 483 | 1.85× [1.84, 1.87] | +46.1% | [+45.8%, +46.4%] | 0.3 pts | 0.5 | yes | yes |
| natural | links | 4096 | valuesBetween | ordered | baseline | 8 | 12.0 µs | 20.0 µs | 1.66× [1.65, 1.68] | +39.9% | [+39.5%, +40.4%] | 0.3 pts | 0.6 | yes | yes |
| natural | links | 4096 | valuesBetween | ordered | btree-sets | 8 | 12.0 µs | 33.0 µs | 2.74× [2.73, 2.75] | +63.5% | [+63.3%, +63.7%] | 0.1 pts | 0.4 | yes | yes |
| natural | links | 4096 | churn | ordered | baseline | 8 | 154 | 116 | 0.76× [0.74, 0.78] | -31.4% | [-34.4%, -28.4%] | 3.8 pts | 4.5 | yes | yes |
| natural | links | 4096 | churn | ordered | btree-sets | 8 | 155 | 198 | 1.29× [1.28, 1.30] | +22.6% | [+22.0%, +23.3%] | 0.7 pts | 2.1 | yes | yes |
| natural | links | 4096 | build | ordered | baseline | 8 | 104.07 ms | 75.32 ms | 0.72× [0.72, 0.73] | -38.0% | [-38.6%, -37.4%] | 0.5 pts | 2.5 | yes | yes |
| natural | links | 4096 | build | ordered | btree-sets | 8 | 103.94 ms | 125.21 ms | 1.20× [1.20, 1.21] | +16.9% | [+16.6%, +17.2%] | 0.3 pts | 2.0 | yes | yes |
| natural | links | 16384 | valuesFor | ordered | baseline | 8 | 359 | 371 | 1.03× [1.02, 1.04] | +3.1% | [+2.4%, +3.9%] | 2.0 pts | 1.4 | yes | yes |
| natural | links | 16384 | valuesFor | ordered | btree-sets | 8 | 370 | 677 | 1.84× [1.81, 1.87] | +45.5% | [+44.6%, +46.4%] | 0.8 pts | 1.0 | yes | yes |
| natural | links | 16384 | valuesBetween | ordered | baseline | 8 | 14.9 µs | 24.9 µs | 1.66× [1.63, 1.68] | +39.7% | [+38.8%, +40.6%] | 0.6 pts | 0.6 | yes | yes |
| natural | links | 16384 | valuesBetween | ordered | btree-sets | 8 | 15.7 µs | 44.7 µs | 2.85× [2.81, 2.90] | +65.0% | [+64.4%, +65.5%] | 0.5 pts | 0.8 | yes | yes |
| natural | links | 16384 | churn | ordered | baseline | 8 | 256 | 216 | 0.83× [0.77, 0.92] | -19.8% | [-30.7%, -9.0%] | 10.6 pts | 3.2 | no | yes |
| natural | links | 16384 | churn | ordered | btree-sets | 8 | 277 | 394 | 1.43× [1.40, 1.45] | +30.0% | [+28.8%, +31.2%] | 1.1 pts | 2.1 | yes | yes |
| natural | links | 16384 | build | ordered | baseline | 8 | 623.86 ms | 505.54 ms | 0.82× [0.81, 0.83] | -21.6% | [-23.0%, -20.2%] | 3.0 pts | 2.9 | yes | yes |
| natural | links | 16384 | build | ordered | btree-sets | 8 | 629.19 ms | 869.21 ms | 1.40× [1.39, 1.41] | +28.6% | [+27.8%, +29.3%] | 0.7 pts | 1.2 | yes | yes |
| natural | links | 65536 | valuesFor | ordered | baseline | 4 | 475 | 510 | 1.09× [1.07, 1.10] | +7.9% | [+6.5%, +9.3%] | 0.9 pts | 0.6 | yes | yes |
| natural | links | 65536 | valuesFor | ordered | btree-sets | 4 | 500 | 951 | 1.91× [1.90, 1.92] | +47.7% | [+47.4%, +47.9%] | 0.1 pts | 0.2 | yes | yes |
| natural | links | 65536 | valuesBetween | ordered | baseline | 4 | 17.2 µs | 27.9 µs | 1.60× [1.55, 1.65] | +37.3% | [+35.3%, +39.4%] | 1.3 pts | 0.8 | yes | yes |
| natural | links | 65536 | valuesBetween | ordered | btree-sets | 4 | 17.9 µs | 54.0 µs | 2.98× [2.90, 3.06] | +66.4% | [+65.5%, +67.3%] | 0.6 pts | 0.9 | yes | yes |
| natural | links | 65536 | churn | ordered | baseline | 4 | 446 | 390 | 0.85× [0.84, 0.86] | -17.3% | [-18.6%, -16.1%] | 0.8 pts | 0.1 | yes | yes |
| natural | links | 65536 | churn | ordered | btree-sets | 4 | 479 | 638 | 1.35× [1.31, 1.40] | +26.0% | [+23.4%, +28.6%] | 1.6 pts | 1.2 | yes | yes |
| natural | links | 65536 | build | ordered | baseline | 4 | 4279.13 ms | 4270.66 ms | 1.00× [0.99, 1.02] | +0.1% | [-1.4%, +1.6%] | 1.0 pts | 1.7 | yes | no |
| natural | links | 65536 | build | ordered | btree-sets | 4 | 4251.86 ms | 6047.14 ms | 1.42× [1.41, 1.44] | +29.8% | [+29.2%, +30.3%] | 0.3 pts | 1.4 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 92.5 | 74.5 | 0.81× [0.79, 0.83] | -23.4% | [-25.8%, -21.0%] | 1.5 pts | 2.4 | no | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 93.1 | 179 | 1.92× [1.91, 1.94] | +48.0% | [+47.6%, +48.5%] | 0.4 pts | 0.6 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 1416 | 2971 | 2.11× [2.10, 2.12] | +52.5% | [+52.3%, +52.8%] | 0.6 pts | 0.7 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 1384 | 5330 | 3.85× [3.81, 3.89] | +74.0% | [+73.8%, +74.3%] | 0.2 pts | 0.9 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 178 | 118 | 0.66× [0.66, 0.66] | -51.0% | [-51.2%, -50.9%] | 0.4 pts | 0.6 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 180 | 235 | 1.32× [1.31, 1.33] | +24.1% | [+23.5%, +24.8%] | 0.7 pts | 1.0 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 6.90 ms | 4.35 ms | 0.63× [0.63, 0.64] | -58.5% | [-59.6%, -57.4%] | 0.9 pts | 1.4 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 6.89 ms | 9.12 ms | 1.32× [1.32, 1.33] | +24.5% | [+24.3%, +24.6%] | 0.3 pts | 1.5 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 114 | 96.7 | 0.86× [0.83, 0.89] | -16.5% | [-20.4%, -12.7%] | 2.5 pts | 4.5 | no | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 114 | 230 | 2.03× [2.01, 2.04] | +50.6% | [+50.2%, +51.0%] | 0.8 pts | 2.4 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 1694 | 3241 | 1.93× [1.90, 1.95] | +48.1% | [+47.3%, +48.8%] | 0.5 pts | 0.9 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 1695 | 5673 | 3.38× [3.33, 3.43] | +70.4% | [+70.0%, +70.9%] | 0.4 pts | 2.0 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 210 | 155 | 0.75× [0.75, 0.76] | -32.7% | [-33.3%, -32.1%] | 6.4 pts | 7.6 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 219 | 324 | 1.48× [1.42, 1.55] | +32.7% | [+29.8%, +35.5%] | 3.4 pts | 4.3 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 32.25 ms | 22.69 ms | 0.71× [0.71, 0.71] | -41.2% | [-41.7%, -40.8%] | 2.9 pts | 6.0 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 32.30 ms | 46.21 ms | 1.44× [1.44, 1.45] | +30.8% | [+30.5%, +31.0%] | 1.2 pts | 5.5 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 147 | 146 | 1.03× [0.97, 1.11] | +3.0% | [-3.6%, +9.6%] | 5.0 pts | 4.6 | no | no |
| natural | street | 65536 | valuesFor | ordered | btree-sets | 8 | 162 | 393 | 2.50× [2.39, 2.61] | +60.0% | [+58.2%, +61.7%] | 1.8 pts | 1.5 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 1893 | 4070 | 2.17× [2.08, 2.26] | +53.8% | [+51.9%, +55.7%] | 1.4 pts | 1.7 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 1982 | 9395 | 4.85× [4.57, 5.17] | +79.4% | [+78.1%, +80.7%] | 0.9 pts | 2.8 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 294 | 272 | 0.94× [0.85, 1.06] | -5.9% | [-17.1%, +5.4%] | 8.6 pts | 2.2 | no | no |
| natural | street | 65536 | churn | ordered | btree-sets | 8 | 345 | 530 | 1.52× [1.47, 1.57] | +34.2% | [+32.1%, +36.3%] | 1.6 pts | 5.2 | yes | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 176.17 ms | 155.82 ms | 0.90× [0.85, 0.95] | -11.2% | [-17.2%, -5.3%] | 3.9 pts | 17.8 | no | yes |
| natural | street | 65536 | build | ordered | btree-sets | 8 | 176.87 ms | 291.07 ms | 1.65× [1.61, 1.70] | +39.5% | [+37.8%, +41.2%] | 1.3 pts | 5.0 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | baseline | 8 | 141 | 123 | 0.87× [0.87, 0.87] | -15.1% | [-15.5%, -14.7%] | 0.4 pts | 0.6 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | btree-sets | 8 | 143 | 231 | 1.63× [1.62, 1.63] | +38.5% | [+38.1%, +38.8%] | 0.6 pts | 1.2 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | baseline | 8 | 3357 | 5048 | 1.50× [1.48, 1.51] | +33.1% | [+32.6%, +33.7%] | 0.5 pts | 0.5 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | btree-sets | 8 | 3342 | 8318 | 2.49× [2.48, 2.50] | +59.8% | [+59.6%, +60.0%] | 0.2 pts | 0.8 | yes | yes |
| natural | url | 4096 | churn | ordered | baseline | 8 | 215 | 156 | 0.74× [0.72, 0.75] | -36.0% | [-39.4%, -32.6%] | 3.5 pts | 4.4 | yes | yes |
| natural | url | 4096 | churn | ordered | btree-sets | 8 | 214 | 238 | 1.11× [1.11, 1.12] | +10.3% | [+9.8%, +10.7%] | 0.5 pts | 0.8 | yes | yes |
| natural | url | 4096 | build | ordered | baseline | 8 | 21.12 ms | 14.79 ms | 0.70× [0.70, 0.70] | -42.7% | [-43.4%, -42.0%] | 0.5 pts | 1.3 | yes | yes |
| natural | url | 4096 | build | ordered | btree-sets | 8 | 21.10 ms | 22.56 ms | 1.07× [1.06, 1.07] | +6.3% | [+5.8%, +6.7%] | 0.4 pts | 1.2 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | baseline | 6 | 168 | 153 | 0.91× [0.91, 0.92] | -9.5% | [-9.8%, -9.2%] | 2.6 pts | 4.1 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | btree-sets | 6 | 174 | 300 | 1.72× [1.69, 1.76] | +42.0% | [+40.7%, +43.2%] | 1.0 pts | 1.0 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | baseline | 6 | 3553 | 5311 | 1.49× [1.48, 1.50] | +33.0% | [+32.5%, +33.4%] | 0.3 pts | 0.7 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | btree-sets | 6 | 3628 | 9038 | 2.50× [2.45, 2.56] | +60.0% | [+59.2%, +60.9%] | 0.7 pts | 1.4 | yes | yes |
| natural | url | 16384 | churn | ordered | baseline | 6 | 280 | 237 | 0.85× [0.83, 0.86] | -18.2% | [-20.0%, -16.4%] | 1.3 pts | 1.8 | yes | yes |
| natural | url | 16384 | churn | ordered | btree-sets | 6 | 296 | 368 | 1.23× [1.23, 1.24] | +18.9% | [+18.4%, +19.4%] | 0.6 pts | 1.0 | yes | yes |
| natural | url | 16384 | build | ordered | baseline | 6 | 94.07 ms | 72.11 ms | 0.77× [0.76, 0.77] | -30.5% | [-31.2%, -29.7%] | 0.8 pts | 3.9 | yes | yes |
| natural | url | 16384 | build | ordered | btree-sets | 6 | 94.14 ms | 113.06 ms | 1.20× [1.20, 1.21] | +16.9% | [+16.5%, +17.4%] | 0.3 pts | 2.9 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | baseline | 8 | 259 | 286 | 1.12× [1.07, 1.17] | +10.5% | [+6.5%, +14.5%] | 2.5 pts | 1.1 | no | yes |
| natural | url | 65536 | valuesFor | ordered | btree-sets | 8 | 284 | 558 | 1.97× [1.92, 2.02] | +49.2% | [+47.8%, +50.6%] | 1.1 pts | 0.8 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | baseline | 8 | 4266 | 7312 | 1.71× [1.68, 1.74] | +41.5% | [+40.5%, +42.5%] | 0.6 pts | 1.1 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | btree-sets | 8 | 4595 | 14.2 µs | 3.11× [3.02, 3.21] | +67.9% | [+66.9%, +68.8%] | 0.7 pts | 1.8 | yes | yes |
| natural | url | 65536 | churn | ordered | baseline | 8 | 492 | 467 | 0.90× [0.87, 0.93] | -11.5% | [-15.3%, -7.7%] | 3.2 pts | 0.4 | no | yes |
| natural | url | 65536 | churn | ordered | btree-sets | 8 | 551 | 690 | 1.24× [1.23, 1.25] | +19.5% | [+18.8%, +20.2%] | 1.7 pts | 0.7 | yes | yes |
| natural | url | 65536 | build | ordered | baseline | 8 | 597.13 ms | 563.72 ms | 0.96× [0.95, 0.96] | -4.6% | [-5.3%, -4.0%] | 2.8 pts | 2.8 | yes | yes |
| natural | url | 65536 | build | ordered | btree-sets | 8 | 594.66 ms | 828.29 ms | 1.39× [1.37, 1.41] | +28.0% | [+27.0%, +28.9%] | 1.3 pts | 2.5 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 123 | 108 | 0.88× [0.87, 0.88] | -14.1% | [-15.0%, -13.1%] | 1.0 pts | 1.6 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 125 | 148 | 1.17× [1.14, 1.21] | +14.7% | [+12.1%, +17.3%] | 1.8 pts | 1.2 | no | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1536 | 2572 | 1.68× [1.62, 1.74] | +40.5% | [+38.5%, +42.5%] | 1.3 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 1486 | 615 | 0.42× [0.41, 0.42] | -140.9% | [-142.5%, -139.3%] | 1.4 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 254 | 241 | 0.95× [0.94, 0.96] | -5.5% | [-6.4%, -4.5%] | 0.8 pts | 1.6 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 251 | 215 | 0.86× [0.85, 0.86] | -16.6% | [-17.0%, -16.2%] | 0.5 pts | 0.7 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 4.08 ms | 3.26 ms | 0.80× [0.80, 0.80] | -25.1% | [-25.6%, -24.6%] | 0.5 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 4.09 ms | 3.17 ms | 0.78× [0.77, 0.78] | -28.8% | [-29.0%, -28.6%] | 0.5 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 6 | 151 | 139 | 0.92× [0.92, 0.93] | -8.5% | [-9.0%, -7.9%] | 0.7 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 6 | 154 | 204 | 1.32× [1.30, 1.35] | +24.5% | [+23.2%, +25.8%] | 1.4 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 1593 | 2900 | 1.82× [1.80, 1.84] | +45.0% | [+44.4%, +45.5%] | 0.4 pts | 0.5 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 6 | 1550 | 758 | 0.49× [0.48, 0.50] | -104.2% | [-107.3%, -101.0%] | 2.4 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 6 | 293 | 305 | 1.05× [1.03, 1.07] | +4.9% | [+3.2%, +6.7%] | 2.0 pts | 5.7 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 6 | 292 | 280 | 0.96× [0.95, 0.97] | -4.1% | [-5.1%, -3.1%] | 0.6 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 6 | 17.82 ms | 15.69 ms | 0.88× [0.88, 0.89] | -13.2% | [-14.0%, -12.3%] | 0.6 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 6 | 17.92 ms | 15.61 ms | 0.87× [0.87, 0.88] | -14.7% | [-15.5%, -13.8%] | 0.5 pts | 1.1 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 196 | 220 | 1.17× [1.04, 1.33] | +14.2% | [+3.6%, +24.8%] | 6.8 pts | 4.8 | no | yes |
| single-value | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 200 | 275 | 1.39× [1.33, 1.45] | +28.0% | [+25.0%, +31.1%] | 1.9 pts | 1.5 | no | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1716 | 3599 | 2.13× [2.06, 2.21] | +53.1% | [+51.5%, +54.8%] | 1.9 pts | 1.3 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 1604 | 1063 | 0.66× [0.59, 0.74] | -52.0% | [-68.6%, -35.5%] | 11.9 pts | 2.8 | no | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 377 | 459 | 1.22× [1.20, 1.23] | +17.8% | [+16.8%, +18.8%] | 1.7 pts | 2.3 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | btree-map | 8 | 414 | 474 | 1.16× [1.14, 1.18] | +13.8% | [+12.1%, +15.5%] | 1.4 pts | 1.9 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 86.10 ms | 89.69 ms | 1.04× [1.04, 1.05] | +4.1% | [+3.7%, +4.5%] | 0.9 pts | 5.9 | yes | yes |
| single-value | dirs | 65536 | build | ordered | btree-map | 8 | 86.06 ms | 84.55 ms | 0.99× [0.97, 1.01] | -1.0% | [-2.8%, +0.9%] | 1.7 pts | 11.6 | yes | no |
| single-value | links | 4096 | valuesFor | ordered | baseline | 4 | 74.9 | 66.2 | 0.88× [0.88, 0.89] | -13.1% | [-14.1%, -12.2%] | 0.6 pts | 2.8 | yes | yes |
| single-value | links | 4096 | valuesFor | ordered | btree-map | 4 | 75.5 | 130 | 1.69× [1.67, 1.72] | +41.0% | [+40.2%, +41.8%] | 0.5 pts | 0.7 | yes | yes |
| single-value | links | 4096 | valuesBetween | ordered | baseline | 4 | 1087 | 2207 | 2.02× [2.01, 2.04] | +50.5% | [+50.2%, +50.9%] | 0.2 pts | 0.3 | yes | yes |
| single-value | links | 4096 | valuesBetween | ordered | btree-map | 4 | 1068 | 712 | 0.67× [0.66, 0.68] | -49.2% | [-51.6%, -46.8%] | 1.5 pts | 0.7 | yes | yes |
| single-value | links | 4096 | churn | ordered | baseline | 4 | 166 | 148 | 0.89× [0.88, 0.89] | -12.4% | [-13.2%, -11.7%] | 0.4 pts | 0.9 | yes | yes |
| single-value | links | 4096 | churn | ordered | btree-map | 4 | 166 | 192 | 1.15× [1.15, 1.15] | +13.0% | [+12.8%, +13.3%] | 0.2 pts | 0.2 | yes | yes |
| single-value | links | 4096 | build | ordered | baseline | 4 | 2.71 ms | 1.99 ms | 0.73× [0.72, 0.75] | -36.5% | [-38.9%, -34.2%] | 1.5 pts | 4.2 | yes | yes |
| single-value | links | 4096 | build | ordered | btree-map | 4 | 2.72 ms | 2.88 ms | 1.06× [1.05, 1.07] | +5.6% | [+4.5%, +6.7%] | 0.7 pts | 2.0 | yes | yes |
| single-value | links | 16384 | valuesFor | ordered | baseline | 8 | 93.6 | 84.3 | 0.91× [0.88, 0.93] | -10.3% | [-13.2%, -7.4%] | 2.0 pts | 6.5 | no | yes |
| single-value | links | 16384 | valuesFor | ordered | btree-map | 8 | 95.1 | 181 | 1.89× [1.86, 1.92] | +47.1% | [+46.2%, +48.0%] | 0.8 pts | 1.4 | yes | yes |
| single-value | links | 16384 | valuesBetween | ordered | baseline | 8 | 1149 | 2445 | 2.12× [2.11, 2.13] | +52.8% | [+52.6%, +53.0%] | 0.3 pts | 0.6 | yes | yes |
| single-value | links | 16384 | valuesBetween | ordered | btree-map | 8 | 1130 | 819 | 0.72× [0.72, 0.73] | -38.2% | [-38.8%, -37.5%] | 0.8 pts | 0.9 | yes | yes |
| single-value | links | 16384 | churn | ordered | baseline | 8 | 195 | 186 | 0.97× [0.92, 1.01] | -3.6% | [-8.1%, +1.0%] | 2.9 pts | 7.6 | no | no |
| single-value | links | 16384 | churn | ordered | btree-map | 8 | 193 | 245 | 1.27× [1.26, 1.28] | +21.5% | [+20.9%, +22.1%] | 0.5 pts | 0.8 | yes | yes |
| single-value | links | 16384 | build | ordered | baseline | 8 | 11.75 ms | 9.54 ms | 0.81× [0.81, 0.82] | -22.7% | [-23.1%, -22.4%] | 0.5 pts | 1.6 | yes | yes |
| single-value | links | 16384 | build | ordered | btree-map | 8 | 11.73 ms | 14.43 ms | 1.23× [1.22, 1.24] | +18.6% | [+17.8%, +19.4%] | 0.6 pts | 2.9 | yes | yes |
| single-value | links | 65536 | valuesFor | ordered | baseline | 8 | 110 | 110 | 1.02× [0.99, 1.05] | +1.8% | [-1.5%, +5.1%] | 2.0 pts | 3.5 | no | no |
| single-value | links | 65536 | valuesFor | ordered | btree-map | 8 | 109 | 234 | 2.10× [2.03, 2.18] | +52.5% | [+50.8%, +54.1%] | 1.1 pts | 5.8 | yes | yes |
| single-value | links | 65536 | valuesBetween | ordered | baseline | 8 | 1256 | 2746 | 2.21× [2.12, 2.31] | +54.7% | [+52.8%, +56.6%] | 1.1 pts | 1.3 | yes | yes |
| single-value | links | 65536 | valuesBetween | ordered | btree-map | 8 | 1219 | 882 | 0.73× [0.73, 0.73] | -37.3% | [-37.8%, -36.9%] | 1.0 pts | 1.3 | yes | yes |
| single-value | links | 65536 | churn | ordered | baseline | 8 | 240 | 266 | 1.12× [1.06, 1.17] | +10.3% | [+5.8%, +14.9%] | 3.3 pts | 3.4 | no | yes |
| single-value | links | 65536 | churn | ordered | btree-map | 8 | 264 | 359 | 1.36× [1.35, 1.37] | +26.6% | [+26.0%, +27.2%] | 0.7 pts | 1.4 | yes | yes |
| single-value | links | 65536 | build | ordered | baseline | 8 | 54.82 ms | 50.67 ms | 0.92× [0.92, 0.93] | -8.1% | [-9.2%, -7.1%] | 0.9 pts | 2.7 | yes | yes |
| single-value | links | 65536 | build | ordered | btree-map | 8 | 54.68 ms | 69.59 ms | 1.27× [1.26, 1.29] | +21.5% | [+20.7%, +22.4%] | 0.7 pts | 2.6 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 81.1 | 67.0 | 0.83× [0.82, 0.84] | -20.4% | [-21.4%, -19.5%] | 0.8 pts | 1.6 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 81.8 | 127 | 1.55× [1.52, 1.57] | +35.4% | [+34.4%, +36.4%] | 0.8 pts | 1.1 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 922 | 2304 | 2.54× [2.51, 2.56] | +60.6% | [+60.1%, +61.0%] | 0.6 pts | 1.0 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 889 | 594 | 0.68× [0.67, 0.68] | -48.1% | [-48.8%, -47.3%] | 1.3 pts | 0.7 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 167 | 149 | 0.90× [0.89, 0.90] | -11.3% | [-12.1%, -10.6%] | 0.6 pts | 0.8 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 167 | 184 | 1.10× [1.09, 1.10] | +8.9% | [+8.5%, +9.4%] | 0.8 pts | 1.4 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.68 ms | 2.04 ms | 0.77× [0.77, 0.78] | -29.3% | [-30.1%, -28.5%] | 1.0 pts | 0.8 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.68 ms | 2.71 ms | 1.02× [0.99, 1.05] | +1.8% | [-1.0%, +4.6%] | 1.8 pts | 3.0 | no | no |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 97.2 | 87.9 | 0.92× [0.89, 0.95] | -9.1% | [-13.0%, -5.2%] | 2.5 pts | 8.6 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 98.0 | 171 | 1.73× [1.71, 1.75] | +42.3% | [+41.7%, +42.9%] | 0.4 pts | 0.9 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1112 | 2484 | 2.24× [2.23, 2.24] | +55.3% | [+55.2%, +55.4%] | 0.3 pts | 0.6 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1083 | 667 | 0.62× [0.61, 0.62] | -62.0% | [-63.3%, -60.7%] | 1.1 pts | 1.2 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 195 | 187 | 0.97× [0.96, 0.98] | -2.8% | [-4.1%, -1.5%] | 2.0 pts | 4.9 | yes | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 195 | 231 | 1.19× [1.17, 1.23] | +16.3% | [+14.2%, +18.5%] | 1.4 pts | 2.3 | no | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 11.83 ms | 9.79 ms | 0.83× [0.82, 0.84] | -20.2% | [-21.3%, -19.0%] | 0.9 pts | 1.6 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 11.82 ms | 13.44 ms | 1.14× [1.13, 1.15] | +12.2% | [+11.6%, +12.7%] | 0.8 pts | 2.7 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 118 | 122 | 1.03× [1.01, 1.06] | +3.3% | [+0.5%, +6.0%] | 1.8 pts | 5.5 | no | yes |
| single-value | street | 65536 | valuesFor | ordered | btree-map | 8 | 118 | 216 | 1.81× [1.73, 1.90] | +44.7% | [+42.2%, +47.3%] | 1.7 pts | 6.3 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 1128 | 2723 | 2.43× [2.39, 2.48] | +58.9% | [+58.2%, +59.6%] | 0.6 pts | 0.9 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | btree-map | 8 | 1105 | 731 | 0.67× [0.66, 0.68] | -49.4% | [-50.9%, -48.0%] | 4.1 pts | 6.4 | yes | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 242 | 277 | 1.12× [1.12, 1.13] | +10.8% | [+10.4%, +11.2%] | 1.0 pts | 0.6 | yes | yes |
| single-value | street | 65536 | churn | ordered | btree-map | 8 | 255 | 321 | 1.28× [1.26, 1.30] | +22.0% | [+20.7%, +23.2%] | 1.0 pts | 2.2 | yes | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 54.21 ms | 52.90 ms | 0.98× [0.96, 0.99] | -2.6% | [-3.7%, -1.5%] | 0.9 pts | 4.2 | yes | yes |
| single-value | street | 65536 | build | ordered | btree-map | 8 | 54.31 ms | 65.49 ms | 1.20× [1.20, 1.21] | +17.0% | [+16.5%, +17.4%] | 0.7 pts | 3.5 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | baseline | 4 | 112 | 95.4 | 0.85× [0.84, 0.86] | -17.1% | [-18.4%, -15.8%] | 0.8 pts | 1.7 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | btree-map | 4 | 112 | 142 | 1.26× [1.23, 1.29] | +20.7% | [+18.9%, +22.5%] | 1.1 pts | 1.2 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | baseline | 4 | 1781 | 2462 | 1.41× [1.37, 1.45] | +29.0% | [+26.8%, +31.1%] | 1.4 pts | 1.0 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | btree-map | 4 | 1728 | 540 | 0.31× [0.31, 0.32] | -218.8% | [-222.5%, -215.2%] | 2.3 pts | 1.7 | yes | yes |
| single-value | url | 4096 | churn | ordered | baseline | 4 | 254 | 206 | 0.82× [0.81, 0.82] | -22.1% | [-22.9%, -21.3%] | 0.5 pts | 0.8 | yes | yes |
| single-value | url | 4096 | churn | ordered | btree-map | 4 | 255 | 201 | 0.79× [0.78, 0.80] | -26.2% | [-27.7%, -24.6%] | 1.0 pts | 1.7 | yes | yes |
| single-value | url | 4096 | build | ordered | baseline | 4 | 3.92 ms | 2.82 ms | 0.72× [0.71, 0.73] | -38.7% | [-40.5%, -37.0%] | 1.1 pts | 1.8 | yes | yes |
| single-value | url | 4096 | build | ordered | btree-map | 4 | 3.93 ms | 2.79 ms | 0.71× [0.69, 0.73] | -40.5% | [-44.2%, -36.9%] | 2.3 pts | 3.2 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | baseline | 8 | 134 | 122 | 0.91× [0.90, 0.92] | -9.9% | [-10.9%, -8.9%] | 0.8 pts | 2.2 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | btree-map | 8 | 135 | 194 | 1.43× [1.42, 1.45] | +30.3% | [+29.5%, +31.0%] | 0.6 pts | 1.4 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | baseline | 8 | 1917 | 2678 | 1.40× [1.38, 1.41] | +28.5% | [+27.8%, +29.3%] | 0.6 pts | 0.8 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | btree-map | 8 | 1889 | 622 | 0.33× [0.33, 0.33] | -203.2% | [-205.5%, -200.9%] | 2.0 pts | 1.6 | yes | yes |
| single-value | url | 16384 | churn | ordered | baseline | 8 | 306 | 268 | 0.91× [0.86, 0.95] | -10.4% | [-15.9%, -5.0%] | 3.9 pts | 3.3 | no | yes |
| single-value | url | 16384 | churn | ordered | btree-map | 8 | 302 | 260 | 0.87× [0.86, 0.87] | -15.3% | [-16.1%, -14.5%] | 0.9 pts | 0.9 | yes | yes |
| single-value | url | 16384 | build | ordered | baseline | 8 | 17.70 ms | 13.39 ms | 0.76× [0.75, 0.76] | -32.0% | [-33.0%, -31.0%] | 0.8 pts | 1.7 | yes | yes |
| single-value | url | 16384 | build | ordered | btree-map | 8 | 17.71 ms | 14.08 ms | 0.80× [0.79, 0.80] | -25.7% | [-26.4%, -25.0%] | 0.7 pts | 2.1 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | baseline | 8 | 195 | 214 | 1.13× [1.09, 1.18] | +11.9% | [+8.2%, +15.5%] | 3.0 pts | 1.2 | no | yes |
| single-value | url | 65536 | valuesFor | ordered | btree-map | 8 | 193 | 269 | 1.38× [1.35, 1.42] | +27.8% | [+25.9%, +29.6%] | 2.4 pts | 1.2 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | baseline | 8 | 2124 | 3624 | 1.72× [1.67, 1.76] | +41.7% | [+40.2%, +43.3%] | 1.2 pts | 0.9 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | btree-map | 8 | 2004 | 1015 | 0.52× [0.49, 0.55] | -94.1% | [-105.5%, -82.8%] | 7.3 pts | 0.7 | no | yes |
| single-value | url | 65536 | churn | ordered | baseline | 8 | 410 | 426 | 1.05× [1.02, 1.07] | +4.5% | [+2.0%, +7.0%] | 1.8 pts | 0.8 | no | yes |
| single-value | url | 65536 | churn | ordered | btree-map | 8 | 445 | 446 | 1.02× [1.01, 1.03] | +1.9% | [+0.8%, +2.9%] | 1.5 pts | 1.7 | yes | yes |
| single-value | url | 65536 | build | ordered | baseline | 8 | 91.09 ms | 79.79 ms | 0.88× [0.87, 0.90] | -13.5% | [-15.5%, -11.6%] | 1.3 pts | 4.8 | yes | yes |
| single-value | url | 65536 | build | ordered | btree-map | 8 | 90.84 ms | 80.23 ms | 0.89× [0.87, 0.90] | -12.7% | [-14.5%, -10.8%] | 1.2 pts | 6.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.74% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=4096 build: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs baseline: the processes scatter 6.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.07% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 build: ordered vs baseline: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.18% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 churn: ordered vs baseline: the pooled difference of 0.05% does not clear the 3.76% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 churn: ordered vs baseline: the pooled interval [-3.52%, 3.63%] includes zero
- natural dirs n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -1.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 build: ordered vs baseline: the pooled interval [-4.47%, 7.36%] includes zero
- natural dirs n=65536 build: ordered vs baseline: the processes scatter 7.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 build: ordered vs baseline: 2 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural links n=4096 churn: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=4096 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=4096 build: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.05% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural links n=16384 churn: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural links n=16384 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=16384 build: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=65536 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=65536 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=65536 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=65536 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=65536 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=65536 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=65536 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=65536 build: ordered vs baseline: the pooled difference of 0.09% does not clear the 0.32% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=65536 build: ordered vs baseline: the pooled interval [-1.44%, 1.61%] includes zero
- natural links n=65536 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=16384 valuesFor: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs baseline: the processes scatter 7.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs btree-sets: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs baseline: the processes scatter 6.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs btree-sets: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: the pooled interval [-3.57%, 9.64%] includes zero
- natural street n=65536 valuesFor: ordered vs baseline: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural street n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -1.76% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 churn: ordered vs baseline: the pooled interval [-17.15%, 5.44%] includes zero
- natural street n=65536 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs btree-sets: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs baseline: the processes scatter 17.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs btree-sets: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.21% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=4096 churn: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 valuesFor: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=16384 build: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 build: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -2.62% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -1.84% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 build: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 build: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.82% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.85% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.49% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 churn: ordered vs baseline: the processes scatter 5.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesBetween: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs baseline: the processes scatter 5.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: the pooled interval [-2.84%, 0.89%] includes zero
- single-value dirs n=65536 build: ordered vs btree-map: the processes scatter 11.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: 2 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value links n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -1.01% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.85% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 build: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 build: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=16384 valuesFor: ordered vs baseline: the processes scatter 6.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.50% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.61% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=16384 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of +0.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=16384 churn: ordered vs baseline: the pooled interval [-8.12%, 1.01%] includes zero
- single-value links n=16384 churn: ordered vs baseline: the processes scatter 7.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=16384 churn: ordered vs baseline: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value links n=16384 build: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.50%, 5.06%] includes zero
- single-value links n=65536 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=65536 valuesFor: ordered vs btree-map: the processes scatter 5.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=65536 churn: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=65536 build: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=65536 build: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 build: ordered vs btree-map: the pooled interval [-0.96%, 4.58%] includes zero
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 8.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs baseline: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs btree-map: the processes scatter 6.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.15% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=65536 valuesBetween: ordered vs btree-map: the processes scatter 6.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 churn: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value url n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.07% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value url n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value url n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value url n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value url n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value url n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value url n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value url n=4096 build: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=16384 churn: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 build: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -1.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=65536 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -6.52% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=65536 build: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 build: ordered vs btree-map: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
