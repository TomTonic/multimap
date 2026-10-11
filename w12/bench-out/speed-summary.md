| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 141 | 142 | 1.01× [0.99, 1.02] | +0.6% | [-0.7%, +1.9%] | 1.1 pts | 1.0 | yes | no |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 142 | 209 | 1.48× [1.46, 1.50] | +32.4% | [+31.5%, +33.3%] | 1.1 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2115 | 2120 | 1.02× [0.99, 1.05] | +1.8% | [-1.4%, +5.0%] | 2.4 pts | 0.7 | no | no |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2106 | 6289 | 3.02× [2.99, 3.04] | +66.9% | [+66.6%, +67.2%] | 0.3 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 244 | 246 | 1.01× [1.00, 1.01] | +0.7% | [+0.2%, +1.2%] | 0.4 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 244 | 276 | 1.14× [1.11, 1.16] | +12.1% | [+10.3%, +13.9%] | 1.7 pts | 1.6 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 11.72 ms | 11.73 ms | 1.00× [1.00, 1.01] | +0.3% | [+0.0%, +0.6%] | 0.5 pts | 2.0 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 11.71 ms | 13.06 ms | 1.12× [1.11, 1.12] | +10.4% | [+9.9%, +10.9%] | 0.4 pts | 1.7 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 4 | 175 | 176 | 1.01× [0.99, 1.02] | +0.5% | [-0.7%, +1.8%] | 0.8 pts | 1.5 | yes | no |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 4 | 179 | 278 | 1.55× [1.49, 1.61] | +35.4% | [+32.8%, +37.9%] | 1.6 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 4 | 2218 | 2226 | 1.00× [0.98, 1.01] | -0.3% | [-1.8%, +1.2%] | 1.0 pts | 0.6 | yes | no |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 4 | 2243 | 6858 | 3.08× [3.01, 3.16] | +67.5% | [+66.7%, +68.3%] | 0.5 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 4 | 284 | 286 | 1.01× [1.00, 1.01] | +0.5% | [+0.0%, +1.0%] | 0.3 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 4 | 306 | 384 | 1.26× [1.23, 1.28] | +20.4% | [+18.6%, +22.1%] | 1.1 pts | 1.0 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 4 | 54.91 ms | 55.11 ms | 1.00× [1.00, 1.01] | +0.5% | [+0.1%, +0.8%] | 0.2 pts | 1.7 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 4 | 54.79 ms | 66.67 ms | 1.22× [1.21, 1.22] | +17.8% | [+17.2%, +18.3%] | 0.4 pts | 2.4 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 225 | 225 | 1.00× [0.99, 1.01] | +0.0% | [-1.1%, +1.1%] | 0.7 pts | 0.8 | yes | no |
| natural | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 262 | 510 | 1.99× [1.92, 2.06] | +49.8% | [+48.0%, +51.5%] | 1.5 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 2487 | 2487 | 1.00× [0.99, 1.02] | +0.3% | [-1.2%, +1.8%] | 1.2 pts | 0.7 | yes | no |
| natural | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 2700 | 11.1 µs | 4.10× [4.04, 4.15] | +75.6% | [+75.3%, +75.9%] | 0.3 pts | 0.9 | yes | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 451 | 420 | 0.91× [0.84, 0.99] | -10.5% | [-19.7%, -1.2%] | 5.7 pts | 0.7 | no | yes |
| natural | dirs | 65536 | churn | ordered | btree-sets | 8 | 513 | 717 | 1.39× [1.36, 1.43] | +28.1% | [+26.3%, +29.9%] | 1.7 pts | 1.5 | yes | yes |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 321.63 ms | 322.22 ms | 1.00× [1.00, 1.00] | +0.0% | [-0.2%, +0.2%] | 0.4 pts | 1.4 | yes | no |
| natural | dirs | 65536 | build | ordered | btree-sets | 8 | 322.49 ms | 472.87 ms | 1.46× [1.46, 1.47] | +31.6% | [+31.3%, +31.9%] | 0.6 pts | 2.0 | yes | yes |
| natural | links | 4096 | valuesFor | ordered | baseline | 4 | 256 | 256 | 0.99× [0.98, 1.00] | -0.8% | [-1.8%, +0.2%] | 0.6 pts | 0.4 | yes | no |
| natural | links | 4096 | valuesFor | ordered | btree-sets | 4 | 261 | 479 | 1.83× [1.82, 1.84] | +45.4% | [+45.0%, +45.7%] | 0.2 pts | 0.3 | yes | yes |
| natural | links | 4096 | valuesBetween | ordered | baseline | 4 | 12.1 µs | 12.0 µs | 1.01× [1.00, 1.02] | +0.7% | [-0.1%, +1.6%] | 0.5 pts | 0.5 | yes | no |
| natural | links | 4096 | valuesBetween | ordered | btree-sets | 4 | 11.8 µs | 33.2 µs | 2.83× [2.75, 2.91] | +64.6% | [+63.6%, +65.6%] | 0.6 pts | 1.5 | yes | yes |
| natural | links | 4096 | churn | ordered | baseline | 4 | 155 | 155 | 1.00× [0.99, 1.00] | -0.3% | [-0.7%, +0.1%] | 0.3 pts | 1.0 | yes | no |
| natural | links | 4096 | churn | ordered | btree-sets | 4 | 154 | 197 | 1.29× [1.28, 1.30] | +22.5% | [+22.1%, +22.9%] | 0.3 pts | 0.5 | yes | yes |
| natural | links | 4096 | build | ordered | baseline | 4 | 103.79 ms | 103.55 ms | 1.00× [1.00, 1.00] | -0.1% | [-0.5%, +0.2%] | 0.2 pts | 2.2 | yes | no |
| natural | links | 4096 | build | ordered | btree-sets | 4 | 103.72 ms | 123.02 ms | 1.19× [1.18, 1.19] | +15.6% | [+15.4%, +15.8%] | 0.1 pts | 1.2 | yes | yes |
| natural | links | 16384 | valuesFor | ordered | baseline | 8 | 357 | 354 | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.6%] | 1.2 pts | 0.8 | yes | no |
| natural | links | 16384 | valuesFor | ordered | btree-sets | 8 | 370 | 687 | 1.85× [1.81, 1.89] | +45.8% | [+44.7%, +47.0%] | 1.1 pts | 1.3 | yes | yes |
| natural | links | 16384 | valuesBetween | ordered | baseline | 8 | 14.6 µs | 14.7 µs | 1.00× [1.00, 1.01] | +0.1% | [-0.4%, +0.5%] | 1.2 pts | 0.7 | yes | no |
| natural | links | 16384 | valuesBetween | ordered | btree-sets | 8 | 15.4 µs | 44.4 µs | 2.89× [2.85, 2.93] | +65.4% | [+64.9%, +65.9%] | 0.4 pts | 0.9 | yes | yes |
| natural | links | 16384 | churn | ordered | baseline | 8 | 248 | 230 | 0.91× [0.89, 0.93] | -9.9% | [-12.1%, -7.7%] | 2.4 pts | 0.4 | no | yes |
| natural | links | 16384 | churn | ordered | btree-sets | 8 | 256 | 362 | 1.43× [1.42, 1.44] | +30.2% | [+29.8%, +30.6%] | 1.0 pts | 2.6 | yes | yes |
| natural | links | 16384 | build | ordered | baseline | 8 | 625.96 ms | 620.61 ms | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.3%] | 0.7 pts | 1.3 | yes | no |
| natural | links | 16384 | build | ordered | btree-sets | 8 | 630.12 ms | 863.38 ms | 1.37× [1.35, 1.38] | +26.7% | [+26.1%, +27.4%] | 0.9 pts | 1.7 | yes | yes |
| natural | links | 65536 | valuesFor | ordered | baseline | 8 | 473 | 469 | 0.99× [0.98, 1.01] | -0.6% | [-2.4%, +1.2%] | 1.3 pts | 0.8 | yes | no |
| natural | links | 65536 | valuesFor | ordered | btree-sets | 8 | 498 | 945 | 1.90× [1.89, 1.91] | +47.3% | [+47.0%, +47.7%] | 0.3 pts | 0.4 | yes | yes |
| natural | links | 65536 | valuesBetween | ordered | baseline | 8 | 16.8 µs | 16.9 µs | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +0.8%] | 1.0 pts | 0.5 | yes | no |
| natural | links | 65536 | valuesBetween | ordered | btree-sets | 8 | 17.7 µs | 53.8 µs | 3.01× [2.92, 3.11] | +66.8% | [+65.8%, +67.9%] | 0.7 pts | 1.0 | yes | yes |
| natural | links | 65536 | churn | ordered | baseline | 8 | 436 | 357 | 0.83× [0.71, 0.98] | -21.0% | [-39.9%, -2.1%] | 13.6 pts | 0.9 | no | yes |
| natural | links | 65536 | churn | ordered | btree-sets | 8 | 470 | 628 | 1.36× [1.34, 1.39] | +26.6% | [+25.1%, +28.0%] | 1.0 pts | 0.8 | yes | yes |
| natural | links | 65536 | build | ordered | baseline | 8 | 4268.67 ms | 4269.68 ms | 1.00× [1.00, 1.00] | -0.2% | [-0.5%, +0.2%] | 0.3 pts | 0.5 | yes | no |
| natural | links | 65536 | build | ordered | btree-sets | 8 | 4262.43 ms | 6024.86 ms | 1.42× [1.40, 1.43] | +29.3% | [+28.6%, +30.1%] | 0.4 pts | 2.1 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 92.3 | 93.1 | 1.01× [0.99, 1.02] | +0.8% | [-0.6%, +2.2%] | 1.1 pts | 0.8 | yes | no |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 92.9 | 178 | 1.91× [1.90, 1.93] | +47.8% | [+47.3%, +48.3%] | 0.5 pts | 1.1 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 1330 | 1351 | 1.04× [1.01, 1.06] | +3.6% | [+1.3%, +5.8%] | 2.2 pts | 0.6 | no | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 1337 | 5319 | 4.00× [3.99, 4.01] | +75.0% | [+74.9%, +75.1%] | 0.1 pts | 0.4 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 177 | 180 | 1.02× [1.01, 1.02] | +1.5% | [+0.7%, +2.3%] | 0.6 pts | 1.2 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 175 | 236 | 1.35× [1.34, 1.36] | +25.9% | [+25.5%, +26.3%] | 0.3 pts | 0.5 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 6.74 ms | 6.82 ms | 1.01× [1.01, 1.01] | +1.2% | [+1.1%, +1.4%] | 0.2 pts | 1.0 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 6.72 ms | 9.08 ms | 1.35× [1.35, 1.35] | +25.9% | [+25.7%, +26.1%] | 0.2 pts | 1.0 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 4 | 116 | 116 | 1.00× [0.99, 1.01] | +0.3% | [-0.5%, +1.2%] | 0.5 pts | 1.3 | yes | no |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 4 | 114 | 226 | 2.01× [1.92, 2.09] | +50.1% | [+48.0%, +52.2%] | 1.3 pts | 2.9 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 4 | 1634 | 1649 | 1.01× [1.00, 1.03] | +1.0% | [-0.4%, +2.5%] | 0.9 pts | 0.6 | yes | no |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 4 | 1630 | 5634 | 3.44× [3.41, 3.48] | +71.0% | [+70.6%, +71.3%] | 0.2 pts | 1.3 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 4 | 203 | 208 | 1.02× [1.01, 1.02] | +1.7% | [+1.3%, +2.1%] | 0.2 pts | 0.9 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 4 | 209 | 301 | 1.44× [1.42, 1.46] | +30.5% | [+29.4%, +31.6%] | 0.7 pts | 1.4 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 4 | 31.70 ms | 31.86 ms | 1.01× [1.00, 1.01] | +0.6% | [+0.4%, +0.9%] | 0.2 pts | 1.0 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 4 | 31.65 ms | 45.84 ms | 1.45× [1.45, 1.45] | +31.1% | [+30.9%, +31.2%] | 0.1 pts | 0.8 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 153 | 152 | 1.00× [0.99, 1.00] | -0.3% | [-1.0%, +0.5%] | 0.8 pts | 0.9 | yes | no |
| natural | street | 65536 | valuesFor | ordered | btree-sets | 8 | 160 | 382 | 2.41× [2.31, 2.53] | +58.5% | [+56.7%, +60.4%] | 2.0 pts | 1.9 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 1827 | 1843 | 1.01× [0.98, 1.03] | +0.6% | [-1.7%, +2.9%] | 1.4 pts | 0.8 | no | no |
| natural | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 1925 | 9335 | 4.88× [4.80, 4.97] | +79.5% | [+79.2%, +79.9%] | 0.5 pts | 1.4 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 284 | 275 | 0.94× [0.89, 0.99] | -6.7% | [-12.1%, -1.3%] | 3.7 pts | 0.7 | no | yes |
| natural | street | 65536 | churn | ordered | btree-sets | 8 | 352 | 552 | 1.51× [1.43, 1.60] | +33.8% | [+29.9%, +37.6%] | 2.4 pts | 4.0 | no | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 177.76 ms | 177.16 ms | 0.99× [0.98, 1.01] | -0.6% | [-1.9%, +0.7%] | 0.9 pts | 2.6 | yes | no |
| natural | street | 65536 | build | ordered | btree-sets | 8 | 177.04 ms | 289.31 ms | 1.64× [1.62, 1.65] | +38.9% | [+38.3%, +39.4%] | 1.0 pts | 3.1 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | baseline | 6 | 141 | 142 | 1.01× [1.00, 1.01] | +0.7% | [+0.1%, +1.2%] | 0.4 pts | 0.7 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | btree-sets | 6 | 142 | 228 | 1.61× [1.60, 1.63] | +38.0% | [+37.4%, +38.6%] | 0.6 pts | 1.0 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | baseline | 6 | 3273 | 3286 | 1.01× [0.99, 1.03] | +1.0% | [-0.8%, +2.9%] | 1.2 pts | 0.7 | yes | no |
| natural | url | 4096 | valuesBetween | ordered | btree-sets | 6 | 3262 | 8342 | 2.56× [2.55, 2.57] | +60.9% | [+60.7%, +61.1%] | 0.2 pts | 0.5 | yes | yes |
| natural | url | 4096 | churn | ordered | baseline | 6 | 214 | 215 | 1.01× [1.00, 1.01] | +0.6% | [+0.4%, +0.7%] | 0.2 pts | 0.4 | yes | yes |
| natural | url | 4096 | churn | ordered | btree-sets | 6 | 213 | 236 | 1.12× [1.10, 1.14] | +10.3% | [+8.7%, +12.0%] | 1.1 pts | 2.5 | yes | yes |
| natural | url | 4096 | build | ordered | baseline | 6 | 20.84 ms | 20.97 ms | 1.01× [1.01, 1.01] | +0.6% | [+0.5%, +0.7%] | 0.3 pts | 1.5 | yes | yes |
| natural | url | 4096 | build | ordered | btree-sets | 6 | 20.91 ms | 22.34 ms | 1.07× [1.06, 1.08] | +6.4% | [+5.7%, +7.1%] | 0.4 pts | 1.7 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | baseline | 8 | 167 | 169 | 1.01× [1.00, 1.01] | +0.6% | [+0.4%, +0.8%] | 0.5 pts | 1.3 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | btree-sets | 8 | 174 | 296 | 1.69× [1.66, 1.72] | +40.8% | [+39.9%, +41.7%] | 0.9 pts | 0.9 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | baseline | 8 | 3504 | 3508 | 1.01× [1.00, 1.01] | +0.8% | [+0.3%, +1.3%] | 0.7 pts | 0.7 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | btree-sets | 8 | 3557 | 8955 | 2.53× [2.52, 2.54] | +60.4% | [+60.3%, +60.6%] | 0.4 pts | 0.8 | yes | yes |
| natural | url | 16384 | churn | ordered | baseline | 8 | 279 | 279 | 0.99× [0.96, 1.02] | -0.9% | [-3.7%, +2.0%] | 1.9 pts | 1.0 | no | no |
| natural | url | 16384 | churn | ordered | btree-sets | 8 | 297 | 368 | 1.23× [1.23, 1.24] | +18.8% | [+18.4%, +19.2%] | 0.4 pts | 0.9 | yes | yes |
| natural | url | 16384 | build | ordered | baseline | 8 | 94.43 ms | 94.48 ms | 1.00× [1.00, 1.00] | -0.1% | [-0.3%, +0.1%] | 0.3 pts | 1.5 | yes | no |
| natural | url | 16384 | build | ordered | btree-sets | 8 | 94.31 ms | 112.33 ms | 1.20× [1.16, 1.23] | +16.3% | [+14.0%, +18.6%] | 1.4 pts | 10.0 | no | yes |
| natural | url | 65536 | valuesFor | ordered | baseline | 8 | 252 | 250 | 1.00× [0.99, 1.02] | +0.3% | [-1.0%, +1.5%] | 0.9 pts | 0.5 | yes | no |
| natural | url | 65536 | valuesFor | ordered | btree-sets | 8 | 284 | 547 | 1.94× [1.88, 2.00] | +48.4% | [+46.9%, +50.0%] | 1.5 pts | 1.0 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | baseline | 8 | 4106 | 4127 | 1.01× [1.00, 1.02] | +1.1% | [+0.2%, +2.1%] | 1.1 pts | 0.9 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | btree-sets | 8 | 4525 | 14.0 µs | 3.10× [3.10, 3.11] | +67.8% | [+67.7%, +67.8%] | 0.2 pts | 0.3 | yes | yes |
| natural | url | 65536 | churn | ordered | baseline | 8 | 479 | 429 | 0.87× [0.84, 0.89] | -15.2% | [-18.4%, -12.0%] | 2.8 pts | 0.3 | no | yes |
| natural | url | 65536 | churn | ordered | btree-sets | 8 | 543 | 688 | 1.26× [1.23, 1.29] | +20.6% | [+18.5%, +22.8%] | 1.6 pts | 0.7 | no | yes |
| natural | url | 65536 | build | ordered | baseline | 8 | 596.26 ms | 597.16 ms | 1.00× [0.99, 1.01] | -0.1% | [-0.7%, +0.5%] | 0.4 pts | 0.6 | yes | no |
| natural | url | 65536 | build | ordered | btree-sets | 8 | 593.26 ms | 820.88 ms | 1.38× [1.37, 1.39] | +27.7% | [+27.2%, +28.2%] | 0.3 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 6 | 123 | 124 | 1.00× [0.99, 1.02] | +0.5% | [-1.5%, +2.4%] | 1.3 pts | 1.6 | yes | no |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 6 | 124 | 143 | 1.15× [1.14, 1.17] | +13.3% | [+12.2%, +14.5%] | 1.0 pts | 0.7 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 1509 | 1523 | 1.00× [0.99, 1.01] | +0.0% | [-0.8%, +0.8%] | 1.1 pts | 0.3 | yes | no |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 6 | 1464 | 619 | 0.43× [0.42, 0.43] | -134.3% | [-136.1%, -132.5%] | 1.4 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 6 | 252 | 253 | 1.00× [0.99, 1.01] | -0.0% | [-0.9%, +0.9%] | 0.6 pts | 1.0 | yes | no |
| single-value | dirs | 4096 | churn | ordered | btree-map | 6 | 253 | 215 | 0.86× [0.85, 0.86] | -16.5% | [-17.3%, -15.7%] | 0.5 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 6 | 3.99 ms | 3.95 ms | 0.99× [0.98, 0.99] | -1.4% | [-1.8%, -0.9%] | 0.4 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 6 | 4.01 ms | 3.12 ms | 0.78× [0.78, 0.78] | -28.4% | [-29.0%, -27.8%] | 0.4 pts | 0.5 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 4 | 152 | 152 | 1.01× [1.00, 1.01] | +0.6% | [+0.2%, +1.0%] | 0.2 pts | 0.4 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 4 | 153 | 200 | 1.31× [1.29, 1.34] | +23.9% | [+22.7%, +25.2%] | 0.8 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 4 | 1543 | 1557 | 1.00× [0.99, 1.01] | -0.1% | [-1.5%, +1.2%] | 0.9 pts | 0.3 | yes | no |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 4 | 1507 | 746 | 0.49× [0.48, 0.50] | -103.9% | [-108.6%, -99.1%] | 3.0 pts | 1.3 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 4 | 298 | 299 | 1.01× [1.00, 1.01] | +0.7% | [+0.1%, +1.2%] | 0.4 pts | 1.1 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 4 | 296 | 284 | 0.97× [0.95, 0.98] | -3.5% | [-5.0%, -1.9%] | 1.0 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 4 | 17.63 ms | 17.58 ms | 1.00× [0.99, 1.01] | -0.2% | [-1.0%, +0.7%] | 0.5 pts | 2.0 | yes | no |
| single-value | dirs | 16384 | build | ordered | btree-map | 4 | 17.66 ms | 15.48 ms | 0.88× [0.87, 0.88] | -13.7% | [-14.3%, -13.2%] | 0.4 pts | 1.6 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 186 | 186 | 1.01× [1.00, 1.01] | +0.5% | [-0.3%, +1.3%] | 0.5 pts | 1.0 | yes | no |
| single-value | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 198 | 270 | 1.40× [1.32, 1.50] | +28.7% | [+24.1%, +33.4%] | 3.1 pts | 2.5 | no | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1597 | 1620 | 1.01× [1.00, 1.01] | +0.7% | [+0.2%, +1.3%] | 2.4 pts | 1.0 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 1580 | 1003 | 0.67× [0.57, 0.79] | -50.0% | [-74.1%, -25.9%] | 16.5 pts | 3.5 | no | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 383 | 383 | 0.98× [0.94, 1.01] | -2.3% | [-5.9%, +1.4%] | 2.6 pts | 0.8 | no | no |
| single-value | dirs | 65536 | churn | ordered | btree-map | 8 | 410 | 467 | 1.15× [1.14, 1.17] | +13.3% | [+12.3%, +14.3%] | 1.9 pts | 2.6 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 85.73 ms | 85.13 ms | 0.99× [0.99, 1.00] | -0.8% | [-1.2%, -0.4%] | 0.3 pts | 1.5 | yes | yes |
| single-value | dirs | 65536 | build | ordered | btree-map | 8 | 85.72 ms | 85.20 ms | 1.00× [0.98, 1.02] | -0.0% | [-1.9%, +1.9%] | 1.2 pts | 5.6 | yes | no |
| single-value | links | 4096 | valuesFor | ordered | baseline | 4 | 75.6 | 75.6 | 1.00× [0.99, 1.01] | +0.2% | [-0.8%, +1.2%] | 0.6 pts | 1.9 | yes | no |
| single-value | links | 4096 | valuesFor | ordered | btree-map | 4 | 75.9 | 130 | 1.70× [1.67, 1.73] | +41.2% | [+40.1%, +42.3%] | 0.7 pts | 1.1 | yes | yes |
| single-value | links | 4096 | valuesBetween | ordered | baseline | 4 | 1046 | 1044 | 1.01× [0.99, 1.02] | +0.8% | [-0.7%, +2.3%] | 1.0 pts | 0.3 | yes | no |
| single-value | links | 4096 | valuesBetween | ordered | btree-map | 4 | 1035 | 709 | 0.69× [0.67, 0.70] | -45.8% | [-49.4%, -42.1%] | 2.3 pts | 0.9 | yes | yes |
| single-value | links | 4096 | churn | ordered | baseline | 4 | 172 | 170 | 0.99× [0.98, 0.99] | -1.4% | [-2.1%, -0.7%] | 0.4 pts | 1.1 | yes | yes |
| single-value | links | 4096 | churn | ordered | btree-map | 4 | 166 | 193 | 1.16× [1.14, 1.18] | +13.7% | [+12.4%, +15.1%] | 0.9 pts | 1.1 | yes | yes |
| single-value | links | 4096 | build | ordered | baseline | 4 | 2.71 ms | 2.69 ms | 0.99× [0.99, 1.00] | -0.5% | [-0.8%, -0.3%] | 0.2 pts | 0.5 | yes | yes |
| single-value | links | 4096 | build | ordered | btree-map | 4 | 2.69 ms | 2.84 ms | 1.06× [1.05, 1.07] | +5.6% | [+4.8%, +6.4%] | 0.5 pts | 1.4 | yes | yes |
| single-value | links | 16384 | valuesFor | ordered | baseline | 4 | 92.7 | 93.0 | 1.00× [0.99, 1.01] | +0.1% | [-0.6%, +0.8%] | 0.4 pts | 1.6 | yes | no |
| single-value | links | 16384 | valuesFor | ordered | btree-map | 4 | 93.8 | 177 | 1.89× [1.87, 1.91] | +47.1% | [+46.4%, +47.8%] | 0.4 pts | 0.9 | yes | yes |
| single-value | links | 16384 | valuesBetween | ordered | baseline | 4 | 1108 | 1113 | 0.99× [0.98, 1.01] | -0.6% | [-2.1%, +0.9%] | 0.9 pts | 0.5 | yes | no |
| single-value | links | 16384 | valuesBetween | ordered | btree-map | 4 | 1098 | 809 | 0.74× [0.73, 0.75] | -35.3% | [-37.7%, -33.0%] | 1.5 pts | 1.0 | yes | yes |
| single-value | links | 16384 | churn | ordered | baseline | 4 | 196 | 194 | 0.99× [0.99, 0.99] | -1.2% | [-1.5%, -1.0%] | 0.2 pts | 0.5 | yes | yes |
| single-value | links | 16384 | churn | ordered | btree-map | 4 | 189 | 240 | 1.27× [1.26, 1.28] | +21.4% | [+20.9%, +21.8%] | 0.3 pts | 0.9 | yes | yes |
| single-value | links | 16384 | build | ordered | baseline | 4 | 11.69 ms | 11.67 ms | 1.00× [1.00, 1.00] | -0.1% | [-0.5%, +0.3%] | 0.2 pts | 0.6 | yes | no |
| single-value | links | 16384 | build | ordered | btree-map | 4 | 11.58 ms | 14.21 ms | 1.22× [1.22, 1.23] | +18.3% | [+17.8%, +18.8%] | 0.3 pts | 1.5 | yes | yes |
| single-value | links | 65536 | valuesFor | ordered | baseline | 8 | 106 | 106 | 1.00× [1.00, 1.01] | +0.3% | [-0.2%, +0.8%] | 0.4 pts | 1.4 | yes | no |
| single-value | links | 65536 | valuesFor | ordered | btree-map | 8 | 107 | 229 | 2.11× [2.09, 2.12] | +52.5% | [+52.2%, +52.9%] | 1.2 pts | 4.5 | yes | yes |
| single-value | links | 65536 | valuesBetween | ordered | baseline | 8 | 1201 | 1211 | 1.00× [0.98, 1.02] | +0.2% | [-1.8%, +2.2%] | 1.8 pts | 0.9 | yes | no |
| single-value | links | 65536 | valuesBetween | ordered | btree-map | 8 | 1187 | 888 | 0.74× [0.73, 0.76] | -34.3% | [-36.2%, -32.3%] | 1.3 pts | 1.1 | yes | yes |
| single-value | links | 65536 | churn | ordered | baseline | 8 | 242 | 241 | 0.98× [0.96, 1.00] | -1.8% | [-3.7%, +0.1%] | 1.5 pts | 0.6 | yes | no |
| single-value | links | 65536 | churn | ordered | btree-map | 8 | 260 | 354 | 1.36× [1.32, 1.40] | +26.6% | [+24.4%, +28.8%] | 1.4 pts | 2.7 | yes | yes |
| single-value | links | 65536 | build | ordered | baseline | 8 | 54.61 ms | 53.85 ms | 0.99× [0.98, 0.99] | -1.4% | [-1.6%, -1.2%] | 0.2 pts | 1.2 | yes | yes |
| single-value | links | 65536 | build | ordered | btree-map | 8 | 54.52 ms | 69.40 ms | 1.27× [1.27, 1.28] | +21.4% | [+21.1%, +21.7%] | 0.4 pts | 5.1 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 79.9 | 79.9 | 1.00× [1.00, 1.00] | -0.1% | [-0.5%, +0.3%] | 0.7 pts | 1.0 | yes | no |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 80.0 | 123 | 1.53× [1.52, 1.54] | +34.6% | [+34.0%, +35.2%] | 0.6 pts | 0.8 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 849 | 859 | 1.03× [0.99, 1.06] | +2.5% | [-0.9%, +5.9%] | 2.6 pts | 0.8 | no | no |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 824 | 579 | 0.71× [0.69, 0.72] | -41.6% | [-44.0%, -39.2%] | 2.1 pts | 0.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 161 | 163 | 1.02× [1.01, 1.02] | +1.5% | [+1.0%, +2.0%] | 0.3 pts | 0.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 161 | 178 | 1.11× [1.10, 1.11] | +9.9% | [+9.5%, +10.2%] | 0.5 pts | 0.8 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.50 ms | 2.50 ms | 1.00× [1.00, 1.01] | +0.1% | [-0.5%, +0.6%] | 0.4 pts | 1.2 | yes | no |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.49 ms | 2.57 ms | 1.03× [1.03, 1.04] | +3.1% | [+2.7%, +3.5%] | 0.3 pts | 0.9 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 4 | 95.5 | 95.7 | 1.00× [1.00, 1.01] | +0.1% | [-0.5%, +0.7%] | 0.4 pts | 1.6 | yes | no |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 4 | 95.8 | 166 | 1.72× [1.69, 1.75] | +41.9% | [+40.8%, +43.0%] | 0.7 pts | 1.3 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 4 | 1062 | 1070 | 1.03× [1.01, 1.04] | +2.8% | [+1.4%, +4.2%] | 0.9 pts | 0.4 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 4 | 1042 | 654 | 0.63× [0.62, 0.64] | -59.2% | [-61.3%, -57.2%] | 1.3 pts | 1.1 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 4 | 190 | 191 | 1.01× [1.01, 1.01] | +1.0% | [+0.7%, +1.3%] | 0.2 pts | 0.5 | yes | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 4 | 189 | 225 | 1.19× [1.19, 1.20] | +16.1% | [+15.7%, +16.6%] | 0.3 pts | 0.6 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 4 | 11.43 ms | 11.40 ms | 1.00× [1.00, 1.00] | +0.0% | [-0.4%, +0.5%] | 0.3 pts | 1.1 | yes | no |
| single-value | street | 16384 | build | ordered | btree-map | 4 | 11.43 ms | 13.07 ms | 1.15× [1.14, 1.15] | +12.7% | [+12.0%, +13.3%] | 0.4 pts | 2.1 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 117 | 116 | 1.00× [1.00, 1.01] | +0.1% | [-0.3%, +0.5%] | 0.4 pts | 1.3 | yes | no |
| single-value | street | 65536 | valuesFor | ordered | btree-map | 8 | 119 | 214 | 1.82× [1.79, 1.84] | +45.0% | [+44.3%, +45.8%] | 0.7 pts | 3.0 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 1078 | 1087 | 1.01× [0.99, 1.03] | +0.8% | [-1.1%, +2.7%] | 2.0 pts | 1.3 | yes | no |
| single-value | street | 65536 | valuesBetween | ordered | btree-map | 8 | 1065 | 719 | 0.68× [0.67, 0.68] | -47.6% | [-48.9%, -46.4%] | 0.9 pts | 1.2 | yes | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 233 | 234 | 0.99× [0.97, 1.00] | -1.4% | [-3.1%, +0.4%] | 1.6 pts | 0.8 | yes | no |
| single-value | street | 65536 | churn | ordered | btree-map | 8 | 253 | 321 | 1.29× [1.28, 1.30] | +22.3% | [+21.7%, +22.9%] | 0.7 pts | 1.6 | yes | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 53.92 ms | 53.55 ms | 0.99× [0.99, 0.99] | -0.8% | [-0.9%, -0.7%] | 0.1 pts | 0.6 | yes | yes |
| single-value | street | 65536 | build | ordered | btree-map | 8 | 54.03 ms | 65.17 ms | 1.21× [1.20, 1.21] | +17.0% | [+16.6%, +17.4%] | 0.4 pts | 3.1 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | baseline | 8 | 112 | 112 | 1.00× [1.00, 1.01] | +0.2% | [-0.4%, +0.7%] | 0.5 pts | 1.1 | yes | no |
| single-value | url | 4096 | valuesFor | ordered | btree-map | 8 | 112 | 139 | 1.24× [1.22, 1.26] | +19.4% | [+18.2%, +20.7%] | 0.9 pts | 1.1 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | baseline | 8 | 1723 | 1745 | 1.02× [1.00, 1.05] | +2.3% | [+0.2%, +4.4%] | 2.3 pts | 0.8 | no | yes |
| single-value | url | 4096 | valuesBetween | ordered | btree-map | 8 | 1662 | 533 | 0.32× [0.32, 0.32] | -211.2% | [-212.3%, -210.1%] | 1.6 pts | 0.8 | yes | yes |
| single-value | url | 4096 | churn | ordered | baseline | 8 | 254 | 254 | 1.00× [1.00, 1.00] | -0.0% | [-0.4%, +0.4%] | 0.6 pts | 1.4 | yes | no |
| single-value | url | 4096 | churn | ordered | btree-map | 8 | 253 | 199 | 0.79× [0.78, 0.79] | -26.9% | [-27.6%, -26.1%] | 0.6 pts | 0.8 | yes | yes |
| single-value | url | 4096 | build | ordered | baseline | 8 | 3.86 ms | 3.80 ms | 0.99× [0.97, 1.01] | -1.1% | [-2.8%, +0.6%] | 1.1 pts | 1.5 | yes | no |
| single-value | url | 4096 | build | ordered | btree-map | 8 | 3.87 ms | 2.76 ms | 0.71× [0.71, 0.72] | -39.9% | [-41.0%, -38.9%] | 0.8 pts | 1.0 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | baseline | 8 | 134 | 135 | 1.01× [1.00, 1.01] | +0.5% | [-0.2%, +1.2%] | 0.5 pts | 1.6 | yes | no |
| single-value | url | 16384 | valuesFor | ordered | btree-map | 8 | 135 | 194 | 1.43× [1.41, 1.44] | +29.9% | [+29.3%, +30.4%] | 0.7 pts | 2.5 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | baseline | 8 | 1884 | 1925 | 1.02× [1.00, 1.05] | +2.4% | [+0.4%, +4.4%] | 1.6 pts | 0.8 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | btree-map | 8 | 1828 | 612 | 0.33× [0.33, 0.34] | -199.1% | [-200.0%, -198.3%] | 1.0 pts | 0.8 | yes | yes |
| single-value | url | 16384 | churn | ordered | baseline | 8 | 311 | 310 | 1.00× [1.00, 1.00] | +0.2% | [+0.1%, +0.3%] | 0.2 pts | 0.6 | yes | no |
| single-value | url | 16384 | churn | ordered | btree-map | 8 | 305 | 260 | 0.85× [0.85, 0.86] | -17.1% | [-18.2%, -16.0%] | 0.9 pts | 0.9 | yes | yes |
| single-value | url | 16384 | build | ordered | baseline | 8 | 17.70 ms | 17.48 ms | 0.99× [0.98, 1.00] | -1.3% | [-2.2%, -0.4%] | 0.6 pts | 1.6 | yes | yes |
| single-value | url | 16384 | build | ordered | btree-map | 8 | 17.74 ms | 14.03 ms | 0.79× [0.79, 0.80] | -26.2% | [-26.7%, -25.7%] | 0.4 pts | 1.1 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | baseline | 8 | 187 | 191 | 1.02× [0.99, 1.04] | +1.7% | [-0.5%, +3.9%] | 2.3 pts | 1.2 | no | no |
| single-value | url | 65536 | valuesFor | ordered | btree-map | 8 | 194 | 274 | 1.42× [1.40, 1.44] | +29.5% | [+28.4%, +30.6%] | 1.6 pts | 0.8 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | baseline | 8 | 2088 | 2114 | 1.01× [0.99, 1.03] | +1.0% | [-0.7%, +2.7%] | 1.9 pts | 0.7 | yes | no |
| single-value | url | 65536 | valuesBetween | ordered | btree-map | 8 | 1989 | 1035 | 0.53× [0.51, 0.55] | -89.9% | [-96.7%, -83.1%] | 4.3 pts | 0.4 | yes | yes |
| single-value | url | 65536 | churn | ordered | baseline | 8 | 403 | 406 | 0.99× [0.99, 1.00] | -0.5% | [-1.3%, +0.2%] | 0.9 pts | 0.5 | yes | no |
| single-value | url | 65536 | churn | ordered | btree-map | 8 | 442 | 446 | 1.02× [0.99, 1.05] | +2.1% | [-0.5%, +4.8%] | 1.7 pts | 1.5 | no | no |
| single-value | url | 65536 | build | ordered | baseline | 8 | 92.17 ms | 91.26 ms | 0.99× [0.98, 1.00] | -1.0% | [-1.8%, -0.3%] | 0.8 pts | 2.9 | yes | yes |
| single-value | url | 65536 | build | ordered | btree-map | 8 | 91.98 ms | 80.66 ms | 0.88× [0.87, 0.89] | -13.8% | [-14.8%, -12.9%] | 0.7 pts | 4.5 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.62% does not clear the 0.78% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.66%, 1.90%] includes zero
- natural dirs n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.40%, 5.05%] includes zero
- natural dirs n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.94% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.54% does not clear the 0.62% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.70%, 1.78%] includes zero
- natural dirs n=16384 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.32% does not clear the 0.54% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.85%, 1.20%] includes zero
- natural dirs n=16384 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs baseline: the pooled difference of 0.00% does not clear the 0.31% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.07%, 1.07%] includes zero
- natural dirs n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.42% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.28% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.21%, 1.78%] includes zero
- natural dirs n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -1.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 build: ordered vs baseline: the pooled difference of 0.01% does not clear the 0.36% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 build: ordered vs baseline: the pooled interval [-0.22%, 0.23%] includes zero
- natural links n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.83% does not clear the 1.98% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.81%, 0.15%] includes zero
- natural links n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.73% does not clear the 2.37% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.10%, 1.55%] includes zero
- natural links n=4096 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 churn: ordered vs baseline: the pooled difference of -0.31% does not clear the 0.39% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=4096 churn: ordered vs baseline: the pooled interval [-0.75%, 0.12%] includes zero
- natural links n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=4096 build: ordered vs baseline: the pooled difference of -0.13% does not clear the 0.13% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=4096 build: ordered vs baseline: the pooled interval [-0.46%, 0.21%] includes zero
- natural links n=4096 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=4096 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural links n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.21% does not clear the 0.74% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.02%, 0.59%] includes zero
- natural links n=16384 valuesBetween: ordered vs baseline: the pooled difference of 0.08% does not clear the 1.23% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.39%, 0.54%] includes zero
- natural links n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural links n=16384 churn: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=16384 build: ordered vs baseline: the pooled difference of -0.40% does not clear the 0.88% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=16384 build: ordered vs baseline: the pooled interval [-1.08%, 0.29%] includes zero
- natural links n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.58% does not clear the 0.66% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=65536 valuesFor: ordered vs baseline: the pooled interval [-2.41%, 1.25%] includes zero
- natural links n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.06% does not clear the 1.80% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=65536 valuesBetween: ordered vs baseline: the pooled interval [-0.71%, 0.82%] includes zero
- natural links n=65536 build: ordered vs baseline: the pooled difference of -0.15% does not clear the 0.17% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=65536 build: ordered vs baseline: the pooled interval [-0.46%, 0.16%] includes zero
- natural links n=65536 build: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.57%, 2.16%] includes zero
- natural street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.97% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.51%, 1.17%] includes zero
- natural street n=16384 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.44%, 2.46%] includes zero
- natural street n=16384 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=16384 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=16384 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural street n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.25% does not clear the 0.46% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 valuesFor: ordered vs baseline: the pooled interval [-0.98%, 0.47%] includes zero
- natural street n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.74%, 2.89%] includes zero
- natural street n=65536 churn: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs baseline: the pooled difference of -0.63% does not clear the 0.70% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 build: ordered vs baseline: the pooled interval [-1.94%, 0.68%] includes zero
- natural street n=65536 build: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.82%, 2.88%] includes zero
- natural url n=4096 churn: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 churn: ordered vs baseline: the pooled interval [-3.71%, 1.98%] includes zero
- natural url n=16384 build: ordered vs baseline: the pooled difference of -0.07% does not clear the 0.09% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=16384 build: ordered vs baseline: the pooled interval [-0.28%, 0.14%] includes zero
- natural url n=16384 build: ordered vs btree-sets: the processes scatter 10.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesFor: ordered vs baseline: the pooled difference of 0.27% does not clear the 0.83% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=65536 valuesFor: ordered vs baseline: the pooled interval [-0.96%, 1.49%] includes zero
- natural url n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -1.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -2.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 build: ordered vs baseline: the pooled difference of -0.10% does not clear the 0.30% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=65536 build: ordered vs baseline: the pooled interval [-0.71%, 0.51%] includes zero
- single-value dirs n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.45% does not clear the 0.60% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.48%, 2.39%] includes zero
- single-value dirs n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.04% does not clear the 1.34% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.84% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.76%, 0.85%] includes zero
- single-value dirs n=4096 churn: ordered vs baseline: the pooled difference of -0.01% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=4096 churn: ordered vs baseline: the pooled interval [-0.88%, 0.86%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.14% does not clear the 1.83% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 valuesBetween: ordered vs baseline: the pooled interval [-1.50%, 1.23%] includes zero
- single-value dirs n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=16384 build: ordered vs baseline: the pooled difference of -0.16% does not clear the 0.26% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 build: ordered vs baseline: the pooled interval [-1.01%, 0.68%] includes zero
- single-value dirs n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value dirs n=65536 valuesFor: ordered vs baseline: the pooled interval [-0.27%, 1.27%] includes zero
- single-value dirs n=65536 valuesFor: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -2.05% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=65536 valuesBetween: ordered vs btree-map: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 churn: ordered vs baseline: the pooled interval [-5.93%, 1.38%] includes zero
- single-value dirs n=65536 churn: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: the pooled difference of -0.03% does not clear the 0.18% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=65536 build: ordered vs btree-map: the pooled interval [-1.93%, 1.87%] includes zero
- single-value dirs n=65536 build: ordered vs btree-map: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: 2 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value links n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.18% does not clear the 0.26% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.84%, 1.21%] includes zero
- single-value links n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.80% does not clear the 2.39% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.75%, 2.35%] includes zero
- single-value links n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.08% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.10% does not clear the 0.12% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.60%, 0.80%] includes zero
- single-value links n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=16384 valuesBetween: ordered vs baseline: the pooled difference of -0.64% does not clear the 1.84% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=16384 valuesBetween: ordered vs baseline: the pooled interval [-2.15%, 0.87%] includes zero
- single-value links n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=16384 build: ordered vs baseline: the pooled difference of -0.09% does not clear the 0.25% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=16384 build: ordered vs baseline: the pooled interval [-0.46%, 0.27%] includes zero
- single-value links n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value links n=65536 valuesFor: ordered vs baseline: the pooled difference of 0.29% does not clear the 0.33% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=65536 valuesFor: ordered vs baseline: the pooled interval [-0.23%, 0.82%] includes zero
- single-value links n=65536 valuesFor: ordered vs btree-map: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.18% does not clear the 0.75% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.80%, 2.15%] includes zero
- single-value links n=65536 churn: ordered vs baseline: the pooled interval [-3.74%, 0.10%] includes zero
- single-value links n=65536 churn: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=65536 build: ordered vs btree-map: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.09% does not clear the 0.52% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=4096 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.27% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.46%, 0.27%] includes zero
- single-value street n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.93%, 5.86%] includes zero
- single-value street n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of +0.77% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 build: ordered vs baseline: the pooled difference of 0.08% does not clear the 0.31% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=4096 build: ordered vs baseline: the pooled interval [-0.46%, 0.63%] includes zero
- single-value street n=16384 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 valuesFor: ordered vs baseline: the pooled difference of 0.09% does not clear the 0.17% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.47%, 0.65%] includes zero
- single-value street n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 build: ordered vs baseline: the pooled difference of 0.02% does not clear the 0.15% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of +0.12% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 build: ordered vs baseline: the pooled interval [-0.45%, 0.49%] includes zero
- single-value street n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- single-value street n=16384 build: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs baseline: the pooled interval [-0.33%, 0.51%] includes zero
- single-value street n=65536 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=65536 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.80% does not clear the 0.83% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.13%, 2.73%] includes zero
- single-value street n=65536 churn: ordered vs baseline: the pooled difference of -1.35% does not clear the 1.41% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=65536 churn: ordered vs baseline: the pooled interval [-3.11%, 0.41%] includes zero
- single-value street n=65536 build: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=4096 valuesFor: ordered vs baseline: the pooled difference of 0.16% does not clear the 0.26% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.41%, 0.72%] includes zero
- single-value url n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -1.10% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=4096 churn: ordered vs baseline: the pooled difference of -0.00% does not clear the 1.19% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=4096 churn: ordered vs baseline: the pooled interval [-0.40%, 0.39%] includes zero
- single-value url n=4096 build: ordered vs baseline: the pooled interval [-2.78%, 0.56%] includes zero
- single-value url n=16384 valuesFor: ordered vs baseline: the pooled interval [-0.16%, 1.19%] includes zero
- single-value url n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 churn: ordered vs baseline: the pooled difference of 0.19% does not clear the 0.44% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 valuesFor: ordered vs baseline: the pooled difference of 1.71% does not clear the 1.72% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 valuesFor: ordered vs baseline: the pooled interval [-0.53%, 3.94%] includes zero
- single-value url n=65536 valuesBetween: ordered vs baseline: the pooled difference of 1.03% does not clear the 1.08% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 valuesBetween: ordered vs baseline: the pooled interval [-0.68%, 2.75%] includes zero
- single-value url n=65536 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -4.83% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=65536 churn: ordered vs baseline: the pooled difference of -0.51% does not clear the 1.69% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 churn: ordered vs baseline: the pooled interval [-1.27%, 0.24%] includes zero
- single-value url n=65536 churn: ordered vs btree-map: the pooled interval [-0.51%, 4.79%] includes zero
- single-value url n=65536 build: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 build: ordered vs btree-map: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
