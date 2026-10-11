| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 103 | 102 | 1.00× [0.99, 1.00] | -0.5% | [-1.2%, +0.3%] | 1.2 pts | 0.7 | yes | no |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 105 | 176 | 1.69× [1.67, 1.70] | +40.8% | [+40.3%, +41.3%] | 0.7 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1478 | 1542 | 1.05× [1.03, 1.06] | +4.4% | [+3.3%, +5.4%] | 1.5 pts | 0.4 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 1509 | 5585 | 3.70× [3.63, 3.76] | +72.9% | [+72.5%, +73.4%] | 0.3 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 211 | 211 | 0.99× [0.97, 1.01] | -1.1% | [-2.9%, +0.6%] | 1.5 pts | 0.9 | yes | no |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 208 | 253 | 1.22× [1.18, 1.26] | +17.9% | [+14.9%, +20.8%] | 2.0 pts | 1.1 | no | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 9.79 ms | 9.87 ms | 1.01× [1.00, 1.02] | +0.9% | [+0.3%, +1.5%] | 0.7 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 9.90 ms | 11.51 ms | 1.16× [1.15, 1.17] | +13.7% | [+13.0%, +14.3%] | 1.0 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 130 | 129 | 0.99× [0.97, 1.01] | -0.9% | [-3.1%, +1.2%] | 1.5 pts | 0.9 | no | no |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 133 | 235 | 1.78× [1.74, 1.82] | +43.9% | [+42.7%, +45.2%] | 1.3 pts | 1.4 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1707 | 1768 | 1.04× [1.02, 1.05] | +3.7% | [+2.4%, +4.9%] | 1.3 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 1778 | 6321 | 3.57× [3.49, 3.66] | +72.0% | [+71.3%, +72.6%] | 0.5 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 260 | 261 | 1.00× [0.99, 1.01] | -0.1% | [-1.2%, +1.1%] | 1.3 pts | 1.1 | yes | no |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 260 | 388 | 1.49× [1.41, 1.57] | +32.7% | [+29.1%, +36.4%] | 2.3 pts | 2.1 | no | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 48.16 ms | 48.39 ms | 1.01× [1.00, 1.01] | +0.6% | [-0.2%, +1.5%] | 0.6 pts | 1.4 | yes | no |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 49.14 ms | 66.33 ms | 1.36× [1.31, 1.40] | +26.2% | [+23.8%, +28.7%] | 1.5 pts | 2.3 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 177 | 177 | 1.01× [0.98, 1.03] | +0.6% | [-2.0%, +3.2%] | 3.1 pts | 1.0 | no | no |
| natural | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 195 | 429 | 2.22× [2.17, 2.27] | +55.0% | [+53.9%, +56.0%] | 1.6 pts | 1.6 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1933 | 1969 | 1.02× [1.02, 1.03] | +2.4% | [+1.9%, +2.9%] | 1.1 pts | 0.7 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 2415 | 11.3 µs | 4.67× [4.44, 4.93] | +78.6% | [+77.5%, +79.7%] | 1.2 pts | 3.9 | yes | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 393 | 398 | 1.00× [0.98, 1.03] | +0.2% | [-2.5%, +3.0%] | 2.2 pts | 0.9 | no | no |
| natural | dirs | 65536 | churn | ordered | btree-sets | 8 | 476 | 713 | 1.52× [1.49, 1.54] | +34.2% | [+33.1%, +35.2%] | 1.1 pts | 1.7 | yes | yes |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 292.52 ms | 291.73 ms | 1.00× [1.00, 1.00] | -0.1% | [-0.5%, +0.3%] | 0.7 pts | 0.8 | yes | no |
| natural | dirs | 65536 | build | ordered | btree-sets | 8 | 292.63 ms | 468.71 ms | 1.61× [1.60, 1.63] | +38.0% | [+37.4%, +38.6%] | 0.8 pts | 1.5 | yes | yes |
| natural | links | 4096 | valuesFor | ordered | baseline | 8 | 201 | 200 | 1.00× [0.98, 1.02] | -0.0% | [-1.9%, +1.9%] | 1.7 pts | 1.2 | yes | no |
| natural | links | 4096 | valuesFor | ordered | btree-sets | 8 | 201 | 400 | 1.98× [1.96, 2.00] | +49.5% | [+49.0%, +50.1%] | 0.5 pts | 0.6 | yes | yes |
| natural | links | 4096 | valuesBetween | ordered | baseline | 8 | 8385 | 9056 | 1.09× [1.06, 1.11] | +7.8% | [+6.1%, +9.6%] | 1.2 pts | 0.8 | yes | yes |
| natural | links | 4096 | valuesBetween | ordered | btree-sets | 8 | 8753 | 29.4 µs | 3.37× [3.33, 3.40] | +70.3% | [+70.0%, +70.6%] | 0.4 pts | 1.0 | yes | yes |
| natural | links | 4096 | churn | ordered | baseline | 8 | 153 | 151 | 0.99× [0.98, 1.01] | -0.7% | [-2.1%, +0.8%] | 1.4 pts | 0.9 | yes | no |
| natural | links | 4096 | churn | ordered | btree-sets | 8 | 145 | 178 | 1.20× [1.19, 1.21] | +16.8% | [+15.9%, +17.6%] | 0.9 pts | 0.7 | yes | yes |
| natural | links | 4096 | build | ordered | baseline | 8 | 95.00 ms | 94.84 ms | 1.00× [0.99, 1.01] | -0.2% | [-1.4%, +1.0%] | 0.8 pts | 2.4 | yes | no |
| natural | links | 4096 | build | ordered | btree-sets | 8 | 95.05 ms | 107.47 ms | 1.13× [1.12, 1.13] | +11.2% | [+10.7%, +11.7%] | 0.5 pts | 1.1 | yes | yes |
| natural | links | 16384 | valuesFor | ordered | baseline | 8 | 273 | 276 | 1.00× [0.98, 1.01] | -0.5% | [-2.2%, +1.3%] | 1.5 pts | 1.0 | yes | no |
| natural | links | 16384 | valuesFor | ordered | btree-sets | 8 | 310 | 618 | 1.99× [1.95, 2.03] | +49.8% | [+48.8%, +50.8%] | 0.7 pts | 0.9 | yes | yes |
| natural | links | 16384 | valuesBetween | ordered | baseline | 8 | 12.6 µs | 13.2 µs | 1.04× [1.04, 1.05] | +4.2% | [+3.6%, +4.7%] | 1.2 pts | 0.6 | yes | yes |
| natural | links | 16384 | valuesBetween | ordered | btree-sets | 8 | 14.5 µs | 42.8 µs | 3.00× [2.94, 3.06] | +66.6% | [+66.0%, +67.3%] | 0.6 pts | 1.2 | yes | yes |
| natural | links | 16384 | churn | ordered | baseline | 8 | 239 | 240 | 0.99× [0.98, 0.99] | -1.2% | [-1.9%, -0.6%] | 1.0 pts | 0.6 | yes | yes |
| natural | links | 16384 | churn | ordered | btree-sets | 8 | 256 | 383 | 1.48× [1.45, 1.52] | +32.6% | [+31.1%, +34.1%] | 1.5 pts | 1.9 | yes | yes |
| natural | links | 16384 | build | ordered | baseline | 8 | 602.54 ms | 603.81 ms | 1.00× [0.99, 1.01] | -0.3% | [-1.4%, +0.7%] | 1.1 pts | 1.7 | yes | no |
| natural | links | 16384 | build | ordered | btree-sets | 8 | 599.39 ms | 831.70 ms | 1.38× [1.35, 1.40] | +27.3% | [+26.1%, +28.5%] | 1.1 pts | 2.0 | yes | yes |
| natural | links | 65536 | valuesFor | ordered | baseline | 8 | 427 | 417 | 0.98× [0.97, 0.99] | -1.9% | [-3.2%, -0.7%] | 1.8 pts | 1.1 | yes | yes |
| natural | links | 65536 | valuesFor | ordered | btree-sets | 8 | 455 | 883 | 1.94× [1.91, 1.96] | +48.4% | [+47.8%, +48.9%] | 0.9 pts | 1.0 | yes | yes |
| natural | links | 65536 | valuesBetween | ordered | baseline | 8 | 18.8 µs | 19.3 µs | 1.00× [0.98, 1.03] | +0.4% | [-2.5%, +3.2%] | 2.8 pts | 1.4 | no | no |
| natural | links | 65536 | valuesBetween | ordered | btree-sets | 8 | 19.6 µs | 60.5 µs | 3.09× [3.06, 3.12] | +67.6% | [+67.3%, +68.0%] | 0.6 pts | 0.9 | yes | yes |
| natural | links | 65536 | churn | ordered | baseline | 8 | 439 | 432 | 0.96× [0.94, 0.99] | -3.9% | [-6.7%, -1.2%] | 2.7 pts | 0.9 | no | yes |
| natural | links | 65536 | churn | ordered | btree-sets | 8 | 481 | 626 | 1.31× [1.29, 1.33] | +23.7% | [+22.6%, +24.9%] | 2.6 pts | 3.1 | yes | yes |
| natural | links | 65536 | build | ordered | baseline | 8 | 4137.82 ms | 4130.30 ms | 1.00× [1.00, 1.00] | -0.1% | [-0.5%, +0.2%] | 0.2 pts | 0.6 | yes | no |
| natural | links | 65536 | build | ordered | btree-sets | 8 | 4126.04 ms | 5746.69 ms | 1.38× [1.34, 1.43] | +27.6% | [+25.3%, +29.9%] | 1.7 pts | 8.7 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 69.2 | 68.1 | 0.98× [0.96, 1.00] | -1.9% | [-3.8%, -0.0%] | 1.3 pts | 0.9 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 70.1 | 141 | 2.01× [1.91, 2.11] | +50.2% | [+47.7%, +52.6%] | 1.6 pts | 2.5 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 941 | 982 | 1.04× [1.00, 1.08] | +3.8% | [+0.4%, +7.2%] | 2.9 pts | 0.8 | no | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 961 | 4711 | 4.90× [4.76, 5.04] | +79.6% | [+79.0%, +80.2%] | 0.5 pts | 1.1 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 156 | 155 | 1.00× [0.99, 1.01] | +0.1% | [-0.7%, +1.0%] | 1.3 pts | 0.9 | yes | no |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 152 | 199 | 1.31× [1.27, 1.34] | +23.4% | [+21.3%, +25.6%] | 1.9 pts | 1.4 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 5.93 ms | 6.00 ms | 1.01× [1.00, 1.02] | +1.1% | [+0.4%, +1.7%] | 0.9 pts | 1.2 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 6.02 ms | 7.51 ms | 1.25× [1.23, 1.27] | +20.1% | [+18.7%, +21.5%] | 0.9 pts | 1.2 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 6 | 84.5 | 83.5 | 0.99× [0.97, 1.00] | -1.3% | [-3.0%, +0.3%] | 1.4 pts | 0.7 | yes | no |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 6 | 85.0 | 190 | 2.22× [2.18, 2.27] | +55.0% | [+54.1%, +55.9%] | 1.2 pts | 1.5 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 6 | 1243 | 1278 | 1.02× [1.00, 1.04] | +2.2% | [+0.4%, +4.0%] | 1.4 pts | 0.8 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 1281 | 5417 | 4.22× [4.11, 4.33] | +76.3% | [+75.7%, +76.9%] | 0.4 pts | 1.2 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 6 | 183 | 186 | 1.02× [1.00, 1.03] | +1.5% | [+0.1%, +2.9%] | 1.0 pts | 1.1 | yes | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 6 | 189 | 287 | 1.53× [1.49, 1.58] | +34.7% | [+32.8%, +36.7%] | 1.8 pts | 2.1 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 6 | 28.02 ms | 28.31 ms | 1.01× [1.00, 1.02] | +1.0% | [+0.3%, +1.7%] | 0.5 pts | 1.3 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 6 | 28.70 ms | 41.49 ms | 1.46× [1.45, 1.47] | +31.5% | [+31.1%, +32.0%] | 0.6 pts | 0.8 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 109 | 108 | 1.00× [0.98, 1.01] | -0.4% | [-1.9%, +1.1%] | 0.9 pts | 0.3 | yes | no |
| natural | street | 65536 | valuesFor | ordered | btree-sets | 8 | 117 | 320 | 2.72× [2.60, 2.85] | +63.2% | [+61.5%, +64.9%] | 1.6 pts | 2.4 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 1399 | 1437 | 1.02× [1.00, 1.04] | +2.1% | [+0.2%, +4.0%] | 1.4 pts | 0.9 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 1532 | 7639 | 5.08× [4.83, 5.36] | +80.3% | [+79.3%, +81.3%] | 1.2 pts | 4.1 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 263 | 264 | 1.00× [0.97, 1.02] | -0.4% | [-3.2%, +2.4%] | 1.8 pts | 1.3 | no | no |
| natural | street | 65536 | churn | ordered | btree-sets | 8 | 307 | 509 | 1.65× [1.57, 1.75] | +39.5% | [+36.3%, +42.7%] | 2.7 pts | 4.9 | yes | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 164.30 ms | 163.14 ms | 0.99× [0.97, 1.01] | -1.2% | [-3.0%, +0.6%] | 1.3 pts | 1.7 | yes | no |
| natural | street | 65536 | build | ordered | btree-sets | 8 | 163.68 ms | 281.29 ms | 1.72× [1.70, 1.74] | +41.8% | [+41.1%, +42.6%] | 0.8 pts | 2.0 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | baseline | 8 | 105 | 105 | 1.00× [0.99, 1.00] | -0.4% | [-1.0%, +0.3%] | 0.7 pts | 0.5 | yes | no |
| natural | url | 4096 | valuesFor | ordered | btree-sets | 8 | 107 | 195 | 1.83× [1.80, 1.86] | +45.3% | [+44.5%, +46.1%] | 0.8 pts | 1.1 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | baseline | 8 | 2290 | 2424 | 1.04× [1.03, 1.06] | +4.2% | [+2.6%, +5.8%] | 2.3 pts | 0.7 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | btree-sets | 8 | 2371 | 7665 | 3.27× [3.15, 3.40] | +69.4% | [+68.2%, +70.6%] | 0.7 pts | 1.3 | yes | yes |
| natural | url | 4096 | churn | ordered | baseline | 8 | 185 | 187 | 1.01× [1.00, 1.01] | +0.5% | [+0.2%, +0.8%] | 1.2 pts | 0.9 | yes | no |
| natural | url | 4096 | churn | ordered | btree-sets | 8 | 190 | 225 | 1.20× [1.17, 1.22] | +16.5% | [+14.7%, +18.2%] | 1.8 pts | 0.7 | yes | yes |
| natural | url | 4096 | build | ordered | baseline | 8 | 17.64 ms | 17.70 ms | 1.01× [1.00, 1.01] | +0.5% | [+0.1%, +0.9%] | 0.4 pts | 0.9 | yes | yes |
| natural | url | 4096 | build | ordered | btree-sets | 8 | 17.82 ms | 20.35 ms | 1.15× [1.14, 1.15] | +12.7% | [+12.3%, +13.0%] | 0.6 pts | 0.9 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | baseline | 8 | 133 | 133 | 0.99× [0.97, 1.01] | -1.3% | [-3.1%, +0.6%] | 2.0 pts | 0.6 | yes | no |
| natural | url | 16384 | valuesFor | ordered | btree-sets | 8 | 136 | 261 | 1.94× [1.90, 1.97] | +48.4% | [+47.4%, +49.3%] | 0.8 pts | 0.7 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | baseline | 8 | 2755 | 2874 | 1.03× [1.02, 1.04] | +2.8% | [+2.2%, +3.4%] | 1.0 pts | 0.5 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | btree-sets | 8 | 2837 | 8302 | 2.91× [2.87, 2.95] | +65.6% | [+65.2%, +66.1%] | 0.5 pts | 1.0 | yes | yes |
| natural | url | 16384 | churn | ordered | baseline | 8 | 256 | 259 | 1.00× [0.99, 1.02] | +0.2% | [-1.2%, +1.6%] | 1.1 pts | 1.0 | yes | no |
| natural | url | 16384 | churn | ordered | btree-sets | 8 | 265 | 402 | 1.50× [1.47, 1.54] | +33.5% | [+31.8%, +35.2%] | 1.7 pts | 2.5 | yes | yes |
| natural | url | 16384 | build | ordered | baseline | 8 | 88.48 ms | 88.06 ms | 0.99× [0.98, 1.01] | -0.6% | [-1.6%, +0.5%] | 0.8 pts | 1.8 | yes | no |
| natural | url | 16384 | build | ordered | btree-sets | 8 | 88.43 ms | 119.90 ms | 1.34× [1.33, 1.36] | +25.6% | [+24.7%, +26.4%] | 0.8 pts | 1.9 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | baseline | 8 | 213 | 214 | 1.00× [0.95, 1.06] | +0.4% | [-5.1%, +6.0%] | 4.9 pts | 1.8 | no | no |
| natural | url | 65536 | valuesFor | ordered | btree-sets | 8 | 251 | 502 | 2.06× [2.01, 2.11] | +51.5% | [+50.4%, +52.7%] | 1.7 pts | 1.7 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | baseline | 8 | 3613 | 3688 | 1.02× [1.00, 1.04] | +2.3% | [+0.4%, +4.1%] | 1.9 pts | 1.1 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | btree-sets | 8 | 5090 | 17.2 µs | 3.45× [3.35, 3.55] | +71.0% | [+70.1%, +71.8%] | 1.0 pts | 2.8 | yes | yes |
| natural | url | 65536 | churn | ordered | baseline | 8 | 458 | 459 | 1.00× [0.99, 1.01] | -0.1% | [-1.0%, +0.8%] | 1.4 pts | 0.6 | yes | no |
| natural | url | 65536 | churn | ordered | btree-sets | 8 | 543 | 724 | 1.33× [1.31, 1.35] | +24.6% | [+23.6%, +25.7%] | 2.8 pts | 4.0 | yes | yes |
| natural | url | 65536 | build | ordered | baseline | 8 | 583.95 ms | 581.84 ms | 0.99× [0.98, 1.00] | -0.8% | [-2.1%, +0.4%] | 1.3 pts | 1.2 | yes | no |
| natural | url | 65536 | build | ordered | btree-sets | 8 | 587.83 ms | 858.57 ms | 1.46× [1.44, 1.49] | +31.7% | [+30.4%, +33.0%] | 1.7 pts | 3.2 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 6 | 91.5 | 90.4 | 0.99× [0.97, 1.01] | -0.9% | [-2.7%, +0.8%] | 1.1 pts | 1.0 | yes | no |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 6 | 92.4 | 123 | 1.33× [1.31, 1.34] | +24.7% | [+23.9%, +25.5%] | 0.8 pts | 0.7 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 1065 | 1107 | 1.03× [1.02, 1.04] | +2.9% | [+1.5%, +4.2%] | 1.6 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 6 | 1046 | 604 | 0.57× [0.56, 0.58] | -75.4% | [-77.6%, -73.1%] | 2.8 pts | 0.6 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 6 | 208 | 209 | 1.00× [0.99, 1.02] | +0.5% | [-0.7%, +1.7%] | 0.9 pts | 0.5 | yes | no |
| single-value | dirs | 4096 | churn | ordered | btree-map | 6 | 209 | 191 | 0.91× [0.90, 0.92] | -9.4% | [-10.6%, -8.2%] | 1.2 pts | 0.6 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 6 | 3.40 ms | 3.40 ms | 0.99× [0.99, 1.00] | -0.6% | [-1.4%, +0.2%] | 0.9 pts | 0.7 | yes | no |
| single-value | dirs | 4096 | build | ordered | btree-map | 6 | 3.49 ms | 2.74 ms | 0.78× [0.77, 0.79] | -27.6% | [-29.4%, -25.9%] | 1.4 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 113 | 113 | 0.99× [0.98, 1.00] | -1.1% | [-2.5%, +0.3%] | 1.3 pts | 0.8 | yes | no |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 115 | 173 | 1.51× [1.48, 1.53] | +33.6% | [+32.5%, +34.7%] | 0.9 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1183 | 1200 | 1.02× [1.01, 1.03] | +2.0% | [+1.1%, +2.9%] | 1.6 pts | 0.7 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 1181 | 762 | 0.65× [0.64, 0.66] | -54.4% | [-56.3%, -52.4%] | 2.4 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 250 | 252 | 1.01× [1.01, 1.02] | +1.4% | [+0.8%, +2.0%] | 0.4 pts | 0.3 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 266 | 269 | 1.01× [0.99, 1.04] | +1.4% | [-1.5%, +4.3%] | 3.8 pts | 1.8 | no | no |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 15.14 ms | 15.04 ms | 1.00× [0.99, 1.00] | -0.3% | [-0.7%, +0.1%] | 0.6 pts | 1.1 | yes | no |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 15.40 ms | 13.94 ms | 0.91× [0.90, 0.92] | -10.0% | [-11.1%, -8.9%] | 0.8 pts | 0.9 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 149 | 149 | 0.98× [0.95, 1.02] | -1.6% | [-4.8%, +1.6%] | 3.1 pts | 1.1 | no | no |
| single-value | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 154 | 237 | 1.57× [1.51, 1.62] | +36.1% | [+34.0%, +38.3%] | 3.0 pts | 1.9 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1294 | 1306 | 1.01× [1.00, 1.02] | +0.9% | [-0.1%, +1.9%] | 1.8 pts | 0.9 | yes | no |
| single-value | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 1305 | 992 | 0.77× [0.72, 0.84] | -29.5% | [-39.9%, -19.2%] | 9.5 pts | 3.7 | no | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 348 | 356 | 1.01× [1.00, 1.02] | +1.0% | [+0.3%, +1.8%] | 0.9 pts | 0.5 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | btree-map | 8 | 374 | 448 | 1.22× [1.20, 1.23] | +17.9% | [+16.8%, +18.9%] | 1.5 pts | 1.3 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 85.52 ms | 84.47 ms | 0.99× [0.98, 1.00] | -1.0% | [-2.1%, +0.2%] | 1.1 pts | 1.0 | yes | no |
| single-value | dirs | 65536 | build | ordered | btree-map | 8 | 83.25 ms | 88.14 ms | 1.05× [1.03, 1.08] | +5.0% | [+2.7%, +7.3%] | 2.0 pts | 2.1 | no | yes |
| single-value | links | 4096 | valuesFor | ordered | baseline | 8 | 54.9 | 54.3 | 0.99× [0.98, 1.00] | -1.0% | [-1.9%, -0.2%] | 1.1 pts | 1.4 | yes | yes |
| single-value | links | 4096 | valuesFor | ordered | btree-map | 8 | 55.2 | 101 | 1.83× [1.82, 1.84] | +45.4% | [+45.1%, +45.8%] | 0.4 pts | 0.8 | yes | yes |
| single-value | links | 4096 | valuesBetween | ordered | baseline | 8 | 748 | 762 | 1.01× [1.00, 1.03] | +1.2% | [-0.3%, +2.7%] | 1.6 pts | 0.5 | yes | no |
| single-value | links | 4096 | valuesBetween | ordered | btree-map | 8 | 747 | 604 | 0.80× [0.79, 0.80] | -25.3% | [-26.2%, -24.4%] | 1.9 pts | 0.6 | yes | yes |
| single-value | links | 4096 | churn | ordered | baseline | 8 | 159 | 151 | 0.95× [0.94, 0.95] | -5.7% | [-6.3%, -5.1%] | 1.5 pts | 1.2 | yes | yes |
| single-value | links | 4096 | churn | ordered | btree-map | 8 | 159 | 161 | 1.01× [0.99, 1.02] | +0.5% | [-0.5%, +1.6%] | 1.1 pts | 0.8 | yes | no |
| single-value | links | 4096 | build | ordered | baseline | 8 | 2.36 ms | 2.34 ms | 0.99× [0.98, 1.00] | -1.1% | [-2.5%, +0.3%] | 1.1 pts | 1.0 | yes | no |
| single-value | links | 4096 | build | ordered | btree-map | 8 | 2.41 ms | 2.30 ms | 0.96× [0.94, 0.98] | -4.0% | [-6.4%, -1.6%] | 1.5 pts | 1.3 | no | yes |
| single-value | links | 16384 | valuesFor | ordered | baseline | 8 | 68.9 | 68.2 | 0.99× [0.99, 0.99] | -1.0% | [-1.5%, -0.5%] | 0.7 pts | 0.5 | yes | no |
| single-value | links | 16384 | valuesFor | ordered | btree-map | 8 | 69.4 | 136 | 1.97× [1.90, 2.05] | +49.3% | [+47.5%, +51.2%] | 1.5 pts | 2.4 | yes | yes |
| single-value | links | 16384 | valuesBetween | ordered | baseline | 8 | 834 | 853 | 1.02× [1.00, 1.04] | +2.3% | [+0.4%, +4.2%] | 1.3 pts | 0.6 | yes | yes |
| single-value | links | 16384 | valuesBetween | ordered | btree-map | 8 | 837 | 702 | 0.84× [0.83, 0.85] | -19.0% | [-20.9%, -17.1%] | 1.4 pts | 0.8 | yes | yes |
| single-value | links | 16384 | churn | ordered | baseline | 8 | 185 | 175 | 0.95× [0.95, 0.96] | -4.8% | [-5.5%, -4.2%] | 1.2 pts | 1.0 | yes | yes |
| single-value | links | 16384 | churn | ordered | btree-map | 8 | 183 | 214 | 1.17× [1.15, 1.18] | +14.4% | [+13.2%, +15.5%] | 1.3 pts | 0.9 | yes | yes |
| single-value | links | 16384 | build | ordered | baseline | 8 | 10.46 ms | 10.32 ms | 0.99× [0.98, 1.00] | -0.7% | [-1.7%, +0.4%] | 0.8 pts | 1.3 | yes | no |
| single-value | links | 16384 | build | ordered | btree-map | 8 | 10.69 ms | 11.89 ms | 1.11× [1.09, 1.13] | +9.9% | [+8.4%, +11.4%] | 1.1 pts | 1.0 | yes | yes |
| single-value | links | 65536 | valuesFor | ordered | baseline | 8 | 87.8 | 87.5 | 0.99× [0.97, 1.01] | -0.9% | [-3.1%, +1.3%] | 1.7 pts | 0.8 | no | no |
| single-value | links | 65536 | valuesFor | ordered | btree-map | 8 | 88.0 | 192 | 2.20× [2.18, 2.22] | +54.5% | [+54.1%, +55.0%] | 0.4 pts | 0.6 | yes | yes |
| single-value | links | 65536 | valuesBetween | ordered | baseline | 8 | 970 | 981 | 1.01× [0.99, 1.03] | +1.0% | [-1.0%, +3.0%] | 1.3 pts | 0.8 | yes | no |
| single-value | links | 65536 | valuesBetween | ordered | btree-map | 8 | 961 | 835 | 0.86× [0.85, 0.87] | -15.8% | [-17.2%, -14.4%] | 1.3 pts | 0.8 | yes | yes |
| single-value | links | 65536 | churn | ordered | baseline | 8 | 238 | 229 | 0.96× [0.96, 0.96] | -4.2% | [-4.7%, -3.7%] | 0.8 pts | 0.6 | yes | yes |
| single-value | links | 65536 | churn | ordered | btree-map | 8 | 243 | 310 | 1.30× [1.24, 1.36] | +23.2% | [+19.6%, +26.7%] | 2.4 pts | 2.9 | no | yes |
| single-value | links | 65536 | build | ordered | baseline | 8 | 56.81 ms | 55.98 ms | 0.99× [0.98, 0.99] | -1.3% | [-1.7%, -0.9%] | 0.6 pts | 1.0 | yes | yes |
| single-value | links | 65536 | build | ordered | btree-map | 8 | 57.35 ms | 65.04 ms | 1.14× [1.12, 1.15] | +12.1% | [+10.8%, +13.3%] | 1.0 pts | 1.9 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 59.4 | 58.7 | 0.98× [0.92, 1.04] | -2.0% | [-8.1%, +4.1%] | 3.9 pts | 5.3 | no | no |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 59.4 | 93.7 | 1.56× [1.48, 1.65] | +36.0% | [+32.6%, +39.3%] | 2.1 pts | 3.7 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 625 | 644 | 1.02× [1.00, 1.04] | +1.9% | [-0.1%, +3.8%] | 1.6 pts | 0.6 | yes | no |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 614 | 512 | 0.83× [0.82, 0.84] | -20.4% | [-21.5%, -19.3%] | 0.9 pts | 0.4 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 137 | 141 | 1.03× [1.01, 1.04] | +2.5% | [+1.2%, +3.8%] | 0.9 pts | 0.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 139 | 142 | 1.01× [1.00, 1.02] | +0.9% | [-0.0%, +1.8%] | 0.7 pts | 0.5 | yes | no |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.26 ms | 2.27 ms | 1.00× [0.99, 1.01] | +0.3% | [-0.6%, +1.2%] | 1.0 pts | 1.1 | yes | no |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.29 ms | 2.11 ms | 0.92× [0.91, 0.93] | -8.6% | [-10.1%, -7.1%] | 1.3 pts | 1.4 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 71.7 | 71.4 | 1.00× [0.99, 1.00] | -0.4% | [-1.2%, +0.4%] | 1.1 pts | 0.8 | yes | no |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 72.4 | 126 | 1.76× [1.70, 1.83] | +43.2% | [+41.2%, +45.3%] | 1.6 pts | 2.3 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 819 | 829 | 1.02× [1.01, 1.04] | +2.1% | [+0.6%, +3.6%] | 1.9 pts | 1.0 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 805 | 588 | 0.74× [0.72, 0.75] | -35.4% | [-38.3%, -32.5%] | 2.4 pts | 1.4 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 166 | 169 | 1.02× [1.00, 1.04] | +1.8% | [-0.1%, +3.7%] | 1.5 pts | 1.1 | yes | no |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 170 | 196 | 1.16× [1.15, 1.17] | +13.9% | [+13.0%, +14.8%] | 2.2 pts | 1.6 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 10.17 ms | 10.16 ms | 1.00× [0.99, 1.00] | -0.3% | [-0.9%, +0.2%] | 0.6 pts | 1.5 | yes | no |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 10.32 ms | 10.67 ms | 1.04× [1.03, 1.05] | +3.7% | [+3.0%, +4.4%] | 0.7 pts | 1.0 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 88.6 | 88.1 | 0.99× [0.98, 1.00] | -0.9% | [-1.9%, +0.2%] | 1.3 pts | 0.7 | yes | no |
| single-value | street | 65536 | valuesFor | ordered | btree-map | 8 | 90.1 | 179 | 2.00× [1.96, 2.04] | +49.9% | [+49.0%, +50.9%] | 1.0 pts | 1.1 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 856 | 868 | 1.02× [1.00, 1.03] | +1.6% | [+0.0%, +3.2%] | 1.7 pts | 1.2 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | btree-map | 8 | 862 | 719 | 0.84× [0.82, 0.86] | -18.8% | [-21.6%, -16.0%] | 2.9 pts | 1.8 | no | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 216 | 219 | 1.00× [0.99, 1.01] | -0.0% | [-0.8%, +0.8%] | 0.9 pts | 1.3 | yes | no |
| single-value | street | 65536 | churn | ordered | btree-map | 8 | 225 | 289 | 1.27× [1.23, 1.30] | +21.1% | [+18.9%, +23.3%] | 1.7 pts | 1.8 | no | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 52.01 ms | 51.60 ms | 0.99× [0.99, 0.99] | -1.2% | [-1.5%, -0.9%] | 0.9 pts | 1.4 | yes | yes |
| single-value | street | 65536 | build | ordered | btree-map | 8 | 53.12 ms | 60.17 ms | 1.13× [1.11, 1.16] | +11.8% | [+9.8%, +13.8%] | 1.8 pts | 1.7 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | baseline | 8 | 84.5 | 83.2 | 0.98× [0.96, 1.01] | -1.5% | [-4.0%, +0.9%] | 1.7 pts | 1.4 | no | no |
| single-value | url | 4096 | valuesFor | ordered | btree-map | 8 | 85.0 | 119 | 1.39× [1.37, 1.42] | +28.2% | [+27.0%, +29.3%] | 0.9 pts | 0.9 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | baseline | 8 | 1248 | 1255 | 1.00× [0.98, 1.02] | +0.1% | [-1.6%, +1.9%] | 2.1 pts | 0.7 | yes | no |
| single-value | url | 4096 | valuesBetween | ordered | btree-map | 8 | 1204 | 599 | 0.49× [0.48, 0.51] | -102.1% | [-108.6%, -95.6%] | 5.7 pts | 1.4 | yes | yes |
| single-value | url | 4096 | churn | ordered | baseline | 8 | 217 | 217 | 1.00× [0.98, 1.03] | +0.5% | [-2.2%, +3.1%] | 1.8 pts | 1.3 | no | no |
| single-value | url | 4096 | churn | ordered | btree-map | 8 | 215 | 188 | 0.87× [0.86, 0.89] | -14.8% | [-16.7%, -12.8%] | 1.8 pts | 0.8 | yes | yes |
| single-value | url | 4096 | build | ordered | baseline | 8 | 3.20 ms | 3.17 ms | 0.99× [0.98, 1.00] | -0.9% | [-1.6%, -0.2%] | 0.8 pts | 0.7 | yes | yes |
| single-value | url | 4096 | build | ordered | btree-map | 8 | 3.22 ms | 2.54 ms | 0.79× [0.78, 0.80] | -26.6% | [-28.1%, -25.1%] | 1.2 pts | 1.0 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | baseline | 8 | 105 | 104 | 0.99× [0.98, 1.01] | -0.8% | [-2.5%, +1.0%] | 1.2 pts | 0.7 | yes | no |
| single-value | url | 16384 | valuesFor | ordered | btree-map | 8 | 106 | 169 | 1.59× [1.55, 1.64] | +37.2% | [+35.6%, +38.9%] | 1.1 pts | 1.1 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | baseline | 8 | 1468 | 1487 | 1.01× [0.99, 1.04] | +1.4% | [-0.7%, +3.5%] | 1.4 pts | 0.8 | no | no |
| single-value | url | 16384 | valuesBetween | ordered | btree-map | 8 | 1450 | 744 | 0.52× [0.50, 0.54] | -92.2% | [-100.5%, -83.8%] | 5.4 pts | 1.8 | yes | yes |
| single-value | url | 16384 | churn | ordered | baseline | 8 | 276 | 276 | 1.00× [0.98, 1.02] | +0.1% | [-1.7%, +1.8%] | 1.3 pts | 0.7 | yes | no |
| single-value | url | 16384 | churn | ordered | btree-map | 8 | 277 | 267 | 0.96× [0.92, 0.99] | -4.7% | [-8.2%, -1.2%] | 2.6 pts | 1.1 | no | yes |
| single-value | url | 16384 | build | ordered | baseline | 8 | 15.49 ms | 15.28 ms | 0.98× [0.97, 0.99] | -2.2% | [-3.3%, -1.1%] | 0.9 pts | 1.3 | yes | yes |
| single-value | url | 16384 | build | ordered | btree-map | 8 | 15.93 ms | 13.53 ms | 0.85× [0.84, 0.86] | -17.6% | [-18.5%, -16.8%] | 1.1 pts | 0.7 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | baseline | 8 | 185 | 184 | 0.99× [0.98, 1.00] | -0.9% | [-1.6%, -0.2%] | 3.9 pts | 1.2 | yes | no |
| single-value | url | 65536 | valuesFor | ordered | btree-map | 8 | 203 | 253 | 1.29× [1.16, 1.44] | +22.4% | [+14.0%, +30.7%] | 7.3 pts | 2.8 | no | yes |
| single-value | url | 65536 | valuesBetween | ordered | baseline | 8 | 1801 | 1845 | 1.01× [0.97, 1.04] | +0.6% | [-3.0%, +4.1%] | 2.5 pts | 1.2 | no | no |
| single-value | url | 65536 | valuesBetween | ordered | btree-map | 8 | 1779 | 1116 | 0.62× [0.58, 0.65] | -62.4% | [-71.3%, -53.4%] | 6.2 pts | 1.7 | no | yes |
| single-value | url | 65536 | churn | ordered | baseline | 8 | 497 | 500 | 1.00× [0.99, 1.02] | +0.3% | [-1.0%, +1.7%] | 1.2 pts | 0.9 | yes | no |
| single-value | url | 65536 | churn | ordered | btree-map | 8 | 478 | 515 | 1.06× [1.00, 1.13] | +5.7% | [-0.2%, +11.5%] | 3.7 pts | 2.6 | no | no |
| single-value | url | 65536 | build | ordered | baseline | 8 | 106.96 ms | 106.31 ms | 0.98× [0.97, 0.99] | -2.4% | [-3.6%, -1.1%] | 1.6 pts | 1.2 | yes | yes |
| single-value | url | 65536 | build | ordered | btree-map | 8 | 105.44 ms | 98.48 ms | 0.94× [0.93, 0.94] | -6.7% | [-7.5%, -5.9%] | 0.7 pts | 0.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.45% does not clear the 0.65% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.18%, 0.28%] includes zero
- natural dirs n=4096 churn: ordered vs baseline: the pooled difference of -1.13% does not clear the 1.14% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=4096 churn: ordered vs baseline: the pooled interval [-2.88%, 0.63%] includes zero
- natural dirs n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.70% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.94% does not clear the 1.20% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.10%, 1.23%] includes zero
- natural dirs n=16384 churn: ordered vs baseline: the pooled difference of -0.06% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=16384 churn: ordered vs baseline: the pooled interval [-1.23%, 1.10%] includes zero
- natural dirs n=16384 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs baseline: the pooled interval [-0.23%, 1.46%] includes zero
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs baseline: the pooled difference of 0.59% does not clear the 1.03% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.97%, 3.16%] includes zero
- natural dirs n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 churn: ordered vs baseline: the pooled difference of 0.23% does not clear the 0.75% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 churn: ordered vs baseline: the pooled interval [-2.52%, 2.98%] includes zero
- natural dirs n=65536 build: ordered vs baseline: the pooled difference of -0.08% does not clear the 0.25% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 build: ordered vs baseline: the pooled interval [-0.46%, 0.31%] includes zero
- natural links n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.01% does not clear the 0.67% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=4096 valuesFor: ordered vs baseline: the pooled interval [-1.93%, 1.91%] includes zero
- natural links n=4096 churn: ordered vs baseline: the pooled difference of -0.66% does not clear the 0.93% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=4096 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural links n=4096 churn: ordered vs baseline: the pooled interval [-2.15%, 0.83%] includes zero
- natural links n=4096 build: ordered vs baseline: the pooled difference of -0.22% does not clear the 0.27% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=4096 build: ordered vs baseline: the pooled interval [-1.42%, 0.97%] includes zero
- natural links n=4096 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.46% does not clear the 1.23% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.24%, 1.33%] includes zero
- natural links n=16384 build: ordered vs baseline: the pooled difference of -0.34% does not clear the 0.39% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=16384 build: ordered vs baseline: the pooled interval [-1.38%, 0.71%] includes zero
- natural links n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.37% does not clear the 1.34% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=65536 valuesBetween: ordered vs baseline: the pooled interval [-2.50%, 3.23%] includes zero
- natural links n=65536 churn: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural links n=65536 build: ordered vs baseline: the pooled difference of -0.15% does not clear the 0.17% noise floor, the bound on what the harness reports between identical code in every process
- natural links n=65536 build: ordered vs baseline: the pooled interval [-0.45%, 0.16%] includes zero
- natural links n=65536 build: ordered vs btree-sets: the processes scatter 8.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 churn: ordered vs baseline: the pooled difference of 0.14% does not clear the 1.02% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=4096 churn: ordered vs baseline: the pooled interval [-0.74%, 1.02%] includes zero
- natural street n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.95%, 0.30%] includes zero
- natural street n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.87% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=16384 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.38% does not clear the 1.48% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.87%, 1.10%] includes zero
- natural street n=65536 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs baseline: the pooled difference of -0.40% does not clear the 0.66% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 churn: ordered vs baseline: the pooled interval [-3.21%, 2.41%] includes zero
- natural street n=65536 churn: ordered vs btree-sets: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs baseline: the pooled interval [-2.97%, 0.62%] includes zero
- natural street n=65536 build: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=4096 valuesFor: ordered vs baseline: the pooled difference of -0.35% does not clear the 0.78% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=4096 valuesFor: ordered vs baseline: the pooled interval [-0.99%, 0.29%] includes zero
- natural url n=4096 churn: ordered vs baseline: the pooled difference of 0.52% does not clear the 0.75% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=4096 build: ordered vs btree-sets: the A/A validations found a systematic difference of +0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=16384 valuesFor: ordered vs baseline: the pooled interval [-3.15%, 0.65%] includes zero
- natural url n=16384 churn: ordered vs baseline: the pooled difference of 0.20% does not clear the 1.00% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=16384 churn: ordered vs baseline: the pooled interval [-1.18%, 1.59%] includes zero
- natural url n=16384 churn: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 build: ordered vs baseline: the pooled interval [-1.62%, 0.51%] includes zero
- natural url n=65536 valuesFor: ordered vs baseline: the pooled difference of 0.44% does not clear the 1.46% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=65536 valuesFor: ordered vs baseline: the pooled interval [-5.14%, 6.02%] includes zero
- natural url n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 churn: ordered vs baseline: the pooled difference of -0.12% does not clear the 1.37% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=65536 churn: ordered vs baseline: the pooled interval [-0.98%, 0.75%] includes zero
- natural url n=65536 churn: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 build: ordered vs baseline: the pooled interval [-2.05%, 0.37%] includes zero
- natural url n=65536 build: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 valuesFor: ordered vs baseline: the pooled interval [-2.70%, 0.81%] includes zero
- single-value dirs n=4096 churn: ordered vs baseline: the pooled difference of 0.48% does not clear the 1.62% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=4096 churn: ordered vs baseline: the pooled interval [-0.73%, 1.70%] includes zero
- single-value dirs n=4096 build: ordered vs baseline: the pooled interval [-1.42%, 0.20%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.45%, 0.31%] includes zero
- single-value dirs n=16384 churn: ordered vs btree-map: the pooled interval [-1.51%, 4.30%] includes zero
- single-value dirs n=16384 build: ordered vs baseline: the pooled difference of -0.31% does not clear the 0.48% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 build: ordered vs baseline: the pooled interval [-0.70%, 0.08%] includes zero
- single-value dirs n=65536 valuesFor: ordered vs baseline: the pooled interval [-4.83%, 1.60%] includes zero
- single-value dirs n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.89% does not clear the 1.17% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=65536 valuesBetween: ordered vs baseline: the pooled interval [-0.12%, 1.90%] includes zero
- single-value dirs n=65536 valuesBetween: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs baseline: the pooled interval [-2.12%, 0.21%] includes zero
- single-value dirs n=65536 build: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.32%, 2.67%] includes zero
- single-value links n=4096 churn: ordered vs btree-map: the pooled interval [-0.53%, 1.56%] includes zero
- single-value links n=4096 build: ordered vs baseline: the pooled interval [-2.47%, 0.31%] includes zero
- single-value links n=16384 valuesFor: ordered vs baseline: the pooled difference of -1.00% does not clear the 1.23% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.80% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value links n=16384 build: ordered vs baseline: the pooled interval [-1.74%, 0.42%] includes zero
- single-value links n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.88% does not clear the 1.10% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=65536 valuesFor: ordered vs baseline: the pooled interval [-3.05%, 1.28%] includes zero
- single-value links n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.98% does not clear the 1.00% noise floor, the bound on what the harness reports between identical code in every process
- single-value links n=65536 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value links n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.00%, 2.96%] includes zero
- single-value links n=65536 churn: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs baseline: the pooled interval [-8.11%, 4.06%] includes zero
- single-value street n=4096 valuesFor: ordered vs baseline: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesBetween: ordered vs baseline: the pooled interval [-0.07%, 3.81%] includes zero
- single-value street n=4096 churn: ordered vs btree-map: the pooled interval [-0.01%, 1.76%] includes zero
- single-value street n=4096 build: ordered vs baseline: the pooled difference of 0.27% does not clear the 0.50% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=4096 build: ordered vs baseline: the pooled interval [-0.64%, 1.18%] includes zero
- single-value street n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.42% does not clear the 0.75% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=16384 valuesFor: ordered vs baseline: the pooled interval [-1.20%, 0.37%] includes zero
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs baseline: the pooled interval [-0.13%, 3.72%] includes zero
- single-value street n=16384 build: ordered vs baseline: the pooled interval [-0.86%, 0.22%] includes zero
- single-value street n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.87% does not clear the 1.21% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.93%, 0.19%] includes zero
- single-value street n=65536 churn: ordered vs baseline: the pooled difference of -0.01% does not clear the 0.68% noise floor, the bound on what the harness reports between identical code in every process
- single-value street n=65536 churn: ordered vs baseline: the pooled interval [-0.84%, 0.82%] includes zero
- single-value url n=4096 valuesFor: ordered vs baseline: the pooled interval [-3.98%, 0.90%] includes zero
- single-value url n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.12% does not clear the 1.26% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=4096 valuesBetween: ordered vs baseline: the pooled interval [-1.64%, 1.87%] includes zero
- single-value url n=4096 churn: ordered vs baseline: the pooled difference of 0.45% does not clear the 0.81% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=4096 churn: ordered vs baseline: the pooled interval [-2.20%, 3.10%] includes zero
- single-value url n=16384 valuesFor: ordered vs baseline: the pooled difference of -0.76% does not clear the 0.80% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=16384 valuesFor: ordered vs baseline: the pooled interval [-2.48%, 0.96%] includes zero
- single-value url n=16384 valuesBetween: ordered vs baseline: the pooled interval [-0.73%, 3.49%] includes zero
- single-value url n=16384 churn: ordered vs baseline: the pooled difference of 0.09% does not clear the 1.04% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=16384 churn: ordered vs baseline: the pooled interval [-1.67%, 1.85%] includes zero
- single-value url n=65536 valuesFor: ordered vs baseline: the pooled difference of -0.92% does not clear the 1.90% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 valuesFor: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesBetween: ordered vs baseline: the pooled difference of 0.55% does not clear the 1.15% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 valuesBetween: ordered vs baseline: the pooled interval [-2.96%, 4.06%] includes zero
- single-value url n=65536 churn: ordered vs baseline: the pooled difference of 0.31% does not clear the 0.97% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 churn: ordered vs baseline: the pooled interval [-1.03%, 1.65%] includes zero
- single-value url n=65536 churn: ordered vs btree-map: the pooled interval [-0.20%, 11.53%] includes zero
- single-value url n=65536 churn: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
