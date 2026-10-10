| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 94.7 | 79.6 | 0.84× [0.83, 0.85] | -19.2% | [-20.4%, -18.0%] | 1.0 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 96.9 | 161 | 1.66× [1.64, 1.68] | +39.8% | [+39.2%, +40.5%] | 0.9 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 98.0 | 86.2 | 0.88× [0.86, 0.89] | -13.9% | [-15.9%, -11.8%] | 2.6 pts | 1.9 | no | yes |
| natural | dirs | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 98.9 | 93.5 | 0.94× [0.92, 0.96] | -6.2% | [-8.3%, -4.2%] | 2.1 pts | 1.8 | no | yes |
| natural | dirs | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 98.0 | 95.4 | 0.97× [0.97, 0.98] | -2.6% | [-3.4%, -1.8%] | 0.6 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1351 | 2544 | 1.91× [1.87, 1.94] | +47.6% | [+46.6%, +48.5%] | 1.5 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 1353 | 5576 | 4.17× [4.10, 4.24] | +76.0% | [+75.6%, +76.4%] | 0.3 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 1347 | 3274 | 2.49× [2.45, 2.53] | +59.9% | [+59.2%, +60.5%] | 0.9 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 1334 | 2537 | 1.92× [1.88, 1.96] | +48.0% | [+46.9%, +49.0%] | 0.9 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 1312 | 1871 | 1.44× [1.43, 1.45] | +30.5% | [+30.0%, +31.0%] | 1.0 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 200 | 152 | 0.75× [0.74, 0.77] | -32.5% | [-34.9%, -30.1%] | 1.8 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 200 | 241 | 1.20× [1.20, 1.21] | +16.9% | [+16.5%, +17.3%] | 0.9 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | churn | ordered | ordered-mk1 | 8 | 203 | 171 | 0.84× [0.84, 0.84] | -19.0% | [-19.3%, -18.6%] | 0.5 pts | 0.5 | yes | yes |
| natural | dirs | 4096 | churn | ordered | ordered-mk2 | 8 | 198 | 217 | 1.10× [1.10, 1.10] | +9.0% | [+8.7%, +9.3%] | 0.4 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | churn | ordered | ordered-mk4 | 8 | 200 | 215 | 1.08× [1.07, 1.09] | +7.4% | [+6.7%, +8.1%] | 0.6 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 9.88 ms | 6.72 ms | 0.68× [0.67, 0.69] | -47.5% | [-49.0%, -45.9%] | 1.5 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 9.83 ms | 11.26 ms | 1.15× [1.14, 1.16] | +12.9% | [+12.1%, +13.6%] | 0.8 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | build | ordered | ordered-mk1 | 8 | 9.78 ms | 7.76 ms | 0.79× [0.78, 0.80] | -26.0% | [-27.7%, -24.4%] | 1.7 pts | 1.5 | yes | yes |
| natural | dirs | 4096 | build | ordered | ordered-mk2 | 8 | 9.77 ms | 10.66 ms | 1.09× [1.08, 1.11] | +8.3% | [+7.1%, +9.6%] | 1.3 pts | 2.1 | yes | yes |
| natural | dirs | 4096 | build | ordered | ordered-mk4 | 8 | 9.67 ms | 10.56 ms | 1.09× [1.08, 1.09] | +7.8% | [+7.5%, +8.2%] | 0.4 pts | 1.0 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 121 | 111 | 0.92× [0.91, 0.93] | -8.8% | [-9.6%, -8.0%] | 0.8 pts | 0.8 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 124 | 223 | 1.79× [1.72, 1.87] | +44.2% | [+41.9%, +46.4%] | 1.4 pts | 1.9 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 126 | 118 | 0.93× [0.91, 0.96] | -7.1% | [-9.9%, -4.3%] | 2.0 pts | 1.7 | no | yes |
| natural | dirs | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 126 | 123 | 0.97× [0.95, 1.00] | -2.7% | [-5.4%, -0.1%] | 1.9 pts | 1.8 | no | yes |
| natural | dirs | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 126 | 123 | 0.98× [0.96, 1.00] | -2.1% | [-4.3%, +0.2%] | 1.5 pts | 1.5 | no | no |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1651 | 3440 | 2.09× [2.08, 2.10] | +52.2% | [+52.0%, +52.4%] | 0.4 pts | 0.9 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 1673 | 6379 | 3.84× [3.79, 3.88] | +73.9% | [+73.6%, +74.2%] | 0.3 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 1650 | 3959 | 2.41× [2.39, 2.43] | +58.6% | [+58.2%, +58.9%] | 0.4 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 1629 | 3204 | 1.97× [1.96, 1.97] | +49.1% | [+48.9%, +49.3%] | 0.2 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 1623 | 2434 | 1.50× [1.49, 1.51] | +33.5% | [+33.1%, +33.9%] | 0.3 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 248 | 207 | 0.84× [0.83, 0.85] | -18.9% | [-20.4%, -17.4%] | 1.9 pts | 1.7 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 246 | 342 | 1.38× [1.36, 1.40] | +27.5% | [+26.3%, +28.6%] | 1.1 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | churn | ordered | ordered-mk1 | 8 | 250 | 224 | 0.90× [0.89, 0.92] | -10.5% | [-12.6%, -8.5%] | 2.0 pts | 2.7 | no | yes |
| natural | dirs | 16384 | churn | ordered | ordered-mk2 | 8 | 244 | 267 | 1.10× [1.08, 1.12] | +9.2% | [+7.8%, +10.6%] | 1.1 pts | 2.5 | yes | yes |
| natural | dirs | 16384 | churn | ordered | ordered-mk4 | 8 | 249 | 273 | 1.10× [1.09, 1.11] | +9.1% | [+7.9%, +10.3%] | 0.9 pts | 1.6 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 48.49 ms | 37.55 ms | 0.78× [0.77, 0.78] | -28.5% | [-29.4%, -27.6%] | 1.1 pts | 0.6 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 48.99 ms | 63.75 ms | 1.30× [1.28, 1.33] | +23.2% | [+21.9%, +24.5%] | 1.5 pts | 1.8 | yes | yes |
| natural | dirs | 16384 | build | ordered | ordered-mk1 | 8 | 48.08 ms | 40.97 ms | 0.85× [0.84, 0.86] | -17.8% | [-18.9%, -16.6%] | 1.2 pts | 1.0 | yes | yes |
| natural | dirs | 16384 | build | ordered | ordered-mk2 | 8 | 48.22 ms | 53.98 ms | 1.12× [1.11, 1.12] | +10.5% | [+10.1%, +11.0%] | 1.0 pts | 1.4 | yes | yes |
| natural | dirs | 16384 | build | ordered | ordered-mk4 | 8 | 48.19 ms | 53.32 ms | 1.11× [1.11, 1.12] | +10.0% | [+9.5%, +10.4%] | 0.9 pts | 1.2 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 166 | 171 | 1.04× [1.01, 1.09] | +4.3% | [+0.8%, +7.8%] | 2.4 pts | 2.7 | no | yes |
| natural | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 172 | 348 | 2.01× [1.93, 2.10] | +50.3% | [+48.3%, +52.4%] | 1.6 pts | 2.9 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 164 | 159 | 0.97× [0.95, 0.98] | -3.6% | [-5.0%, -2.1%] | 1.3 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 164 | 163 | 1.00× [0.98, 1.02] | +0.1% | [-1.9%, +2.1%] | 1.6 pts | 1.4 | yes | no |
| natural | dirs | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 163 | 162 | 0.99× [0.97, 1.01] | -1.1% | [-3.2%, +1.0%] | 1.8 pts | 1.7 | no | no |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1867 | 3995 | 2.13× [2.12, 2.15] | +53.1% | [+52.8%, +53.5%] | 0.5 pts | 1.4 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 2146 | 8576 | 4.00× [3.85, 4.18] | +75.0% | [+74.0%, +76.1%] | 0.9 pts | 1.8 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1837 | 4315 | 2.35× [2.34, 2.36] | +57.4% | [+57.2%, +57.7%] | 0.3 pts | 0.7 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1833 | 3542 | 1.94× [1.91, 1.96] | +48.4% | [+47.7%, +49.1%] | 0.6 pts | 1.6 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1823 | 2764 | 1.52× [1.51, 1.54] | +34.3% | [+33.7%, +34.9%] | 0.5 pts | 1.3 | yes | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 342 | 364 | 1.05× [1.05, 1.06] | +5.1% | [+4.6%, +5.7%] | 0.9 pts | 0.9 | yes | yes |
| natural | dirs | 65536 | churn | ordered | btree-sets | 8 | 397 | 581 | 1.46× [1.42, 1.49] | +31.3% | [+29.7%, +32.8%] | 1.2 pts | 3.4 | yes | yes |
| natural | dirs | 65536 | churn | ordered | ordered-mk1 | 8 | 333 | 346 | 1.03× [1.01, 1.05] | +2.9% | [+1.1%, +4.8%] | 1.4 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | churn | ordered | ordered-mk2 | 8 | 331 | 379 | 1.14× [1.13, 1.16] | +12.5% | [+11.2%, +13.7%] | 1.3 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | churn | ordered | ordered-mk4 | 8 | 327 | 358 | 1.10× [1.06, 1.14] | +9.1% | [+6.0%, +12.1%] | 1.8 pts | 1.2 | no | yes |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 265.67 ms | 243.02 ms | 0.92× [0.91, 0.93] | -8.4% | [-9.4%, -7.5%] | 1.2 pts | 1.4 | yes | yes |
| natural | dirs | 65536 | build | ordered | btree-sets | 8 | 267.21 ms | 404.16 ms | 1.51× [1.49, 1.53] | +33.8% | [+33.0%, +34.6%] | 0.7 pts | 1.2 | yes | yes |
| natural | dirs | 65536 | build | ordered | ordered-mk1 | 8 | 264.58 ms | 244.66 ms | 0.93× [0.91, 0.94] | -8.1% | [-9.7%, -6.5%] | 1.4 pts | 1.5 | yes | yes |
| natural | dirs | 65536 | build | ordered | ordered-mk2 | 8 | 266.97 ms | 299.68 ms | 1.12× [1.10, 1.14] | +10.8% | [+9.1%, +12.4%] | 1.4 pts | 2.7 | yes | yes |
| natural | dirs | 65536 | build | ordered | ordered-mk4 | 8 | 266.18 ms | 296.06 ms | 1.11× [1.10, 1.12] | +10.0% | [+8.9%, +11.0%] | 1.0 pts | 1.5 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 64.3 | 51.6 | 0.80× [0.79, 0.81] | -24.9% | [-27.1%, -22.8%] | 2.9 pts | 2.8 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 64.9 | 136 | 2.08× [2.06, 2.11] | +52.0% | [+51.4%, +52.5%] | 0.9 pts | 1.8 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 64.5 | 52.9 | 0.82× [0.81, 0.82] | -22.6% | [-23.8%, -21.4%] | 2.8 pts | 2.3 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 64.5 | 58.1 | 0.90× [0.88, 0.91] | -11.7% | [-13.4%, -10.0%] | 2.4 pts | 1.8 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 64.2 | 60.1 | 0.94× [0.93, 0.94] | -6.9% | [-7.6%, -6.2%] | 0.5 pts | 0.4 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 857 | 2058 | 2.45× [2.37, 2.52] | +59.1% | [+57.9%, +60.4%] | 1.0 pts | 1.2 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 861 | 4595 | 5.40× [5.29, 5.51] | +81.5% | [+81.1%, +81.9%] | 0.3 pts | 0.9 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 858 | 2743 | 3.26× [3.20, 3.33] | +69.4% | [+68.8%, +69.9%] | 0.4 pts | 0.8 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 850 | 2160 | 2.58× [2.53, 2.63] | +61.2% | [+60.5%, +61.9%] | 0.8 pts | 1.1 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 842 | 1641 | 1.98× [1.93, 2.03] | +49.4% | [+48.1%, +50.7%] | 0.9 pts | 0.9 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 144 | 87.5 | 0.61× [0.60, 0.61] | -65.1% | [-66.6%, -63.6%] | 1.7 pts | 1.3 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 146 | 188 | 1.30× [1.29, 1.32] | +23.3% | [+22.3%, +24.4%] | 1.2 pts | 1.7 | yes | yes |
| natural | street | 4096 | churn | ordered | ordered-mk1 | 8 | 144 | 107 | 0.74× [0.73, 0.75] | -34.8% | [-36.5%, -33.1%] | 1.2 pts | 0.9 | yes | yes |
| natural | street | 4096 | churn | ordered | ordered-mk2 | 8 | 143 | 144 | 1.00× [0.99, 1.02] | +0.3% | [-1.0%, +1.6%] | 1.0 pts | 1.0 | yes | no |
| natural | street | 4096 | churn | ordered | ordered-mk4 | 8 | 143 | 148 | 1.04× [1.03, 1.05] | +3.5% | [+2.6%, +4.4%] | 0.5 pts | 0.7 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 5.94 ms | 3.44 ms | 0.58× [0.57, 0.59] | -73.1% | [-76.2%, -69.9%] | 3.6 pts | 1.6 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 5.98 ms | 7.51 ms | 1.25× [1.20, 1.31] | +20.3% | [+16.8%, +23.7%] | 2.3 pts | 2.2 | no | yes |
| natural | street | 4096 | build | ordered | ordered-mk1 | 8 | 5.88 ms | 4.26 ms | 0.73× [0.71, 0.74] | -37.3% | [-39.9%, -34.8%] | 1.5 pts | 1.2 | yes | yes |
| natural | street | 4096 | build | ordered | ordered-mk2 | 8 | 5.86 ms | 6.21 ms | 1.06× [1.04, 1.08] | +5.3% | [+3.5%, +7.2%] | 1.1 pts | 1.4 | yes | yes |
| natural | street | 4096 | build | ordered | ordered-mk4 | 8 | 5.86 ms | 6.43 ms | 1.09× [1.06, 1.13] | +8.6% | [+6.0%, +11.3%] | 1.6 pts | 3.0 | no | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 80.2 | 70.9 | 0.88× [0.88, 0.89] | -13.1% | [-14.1%, -12.0%] | 1.6 pts | 1.4 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 81.4 | 184 | 2.25× [2.19, 2.31] | +55.6% | [+54.4%, +56.7%] | 0.9 pts | 2.8 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 80.6 | 72.2 | 0.89× [0.86, 0.93] | -12.1% | [-16.1%, -8.1%] | 2.6 pts | 2.4 | no | yes |
| natural | street | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 80.8 | 76.7 | 0.95× [0.93, 0.96] | -5.7% | [-7.7%, -3.7%] | 1.7 pts | 1.4 | no | yes |
| natural | street | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 80.5 | 77.8 | 0.97× [0.96, 0.99] | -2.7% | [-4.0%, -1.4%] | 1.1 pts | 0.9 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 1163 | 2708 | 2.33× [2.31, 2.36] | +57.1% | [+56.7%, +57.6%] | 0.5 pts | 1.1 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 1177 | 5378 | 4.60× [4.53, 4.67] | +78.3% | [+77.9%, +78.6%] | 0.4 pts | 1.3 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 1157 | 3303 | 2.87× [2.85, 2.88] | +65.1% | [+64.9%, +65.3%] | 0.2 pts | 0.7 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 1144 | 2661 | 2.32× [2.28, 2.37] | +57.0% | [+56.2%, +57.7%] | 0.5 pts | 1.3 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 1136 | 2036 | 1.80× [1.78, 1.82] | +44.5% | [+43.8%, +45.2%] | 0.5 pts | 1.3 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 176 | 127 | 0.72× [0.70, 0.74] | -39.5% | [-43.6%, -35.3%] | 3.6 pts | 2.4 | no | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 177 | 266 | 1.51× [1.48, 1.54] | +33.7% | [+32.2%, +35.2%] | 1.6 pts | 2.4 | yes | yes |
| natural | street | 16384 | churn | ordered | ordered-mk1 | 8 | 175 | 145 | 0.83× [0.82, 0.85] | -20.4% | [-22.6%, -18.2%] | 1.8 pts | 1.5 | no | yes |
| natural | street | 16384 | churn | ordered | ordered-mk2 | 8 | 175 | 185 | 1.06× [1.04, 1.08] | +5.7% | [+4.0%, +7.5%] | 1.5 pts | 2.0 | yes | yes |
| natural | street | 16384 | churn | ordered | ordered-mk4 | 8 | 174 | 186 | 1.07× [1.06, 1.08] | +6.7% | [+6.0%, +7.3%] | 0.6 pts | 0.8 | yes | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 28.62 ms | 18.76 ms | 0.65× [0.64, 0.67] | -53.1% | [-56.2%, -50.0%] | 2.2 pts | 1.1 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 29.02 ms | 41.73 ms | 1.44× [1.42, 1.46] | +30.6% | [+29.6%, +31.7%] | 1.6 pts | 1.1 | yes | yes |
| natural | street | 16384 | build | ordered | ordered-mk1 | 8 | 28.32 ms | 22.48 ms | 0.79× [0.78, 0.80] | -26.6% | [-28.1%, -25.0%] | 1.7 pts | 2.0 | yes | yes |
| natural | street | 16384 | build | ordered | ordered-mk2 | 8 | 28.44 ms | 30.59 ms | 1.08× [1.06, 1.09] | +7.2% | [+6.0%, +8.4%] | 1.1 pts | 1.4 | yes | yes |
| natural | street | 16384 | build | ordered | ordered-mk4 | 8 | 28.29 ms | 31.22 ms | 1.11× [1.10, 1.12] | +9.5% | [+8.7%, +10.4%] | 0.7 pts | 1.1 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 105 | 100.0 | 0.97× [0.93, 1.00] | -3.5% | [-7.3%, +0.4%] | 3.4 pts | 2.5 | no | no |
| natural | street | 65536 | valuesFor | ordered | btree-sets | 8 | 110 | 297 | 2.68× [2.54, 2.83] | +62.7% | [+60.7%, +64.7%] | 1.6 pts | 3.2 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 104 | 99.3 | 0.96× [0.93, 0.98] | -4.5% | [-7.0%, -2.0%] | 1.9 pts | 1.4 | no | yes |
| natural | street | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 105 | 105 | 1.01× [0.98, 1.04] | +0.7% | [-2.2%, +3.7%] | 1.9 pts | 1.4 | no | no |
| natural | street | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 104 | 104 | 1.00× [0.98, 1.01] | -0.5% | [-2.4%, +1.4%] | 1.9 pts | 1.4 | yes | no |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 1342 | 3101 | 2.31× [2.29, 2.34] | +56.8% | [+56.3%, +57.2%] | 0.3 pts | 0.7 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 1522 | 7356 | 4.89× [4.80, 4.99] | +79.6% | [+79.2%, +79.9%] | 0.8 pts | 1.1 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1335 | 3628 | 2.75× [2.71, 2.78] | +63.6% | [+63.1%, +64.0%] | 0.4 pts | 1.4 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1330 | 2997 | 2.26× [2.24, 2.28] | +55.8% | [+55.3%, +56.2%] | 0.4 pts | 1.0 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1309 | 2286 | 1.74× [1.72, 1.77] | +42.7% | [+42.0%, +43.4%] | 0.4 pts | 1.0 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 231 | 211 | 0.92× [0.89, 0.94] | -8.9% | [-11.8%, -6.1%] | 2.3 pts | 3.4 | no | yes |
| natural | street | 65536 | churn | ordered | btree-sets | 8 | 263 | 442 | 1.68× [1.66, 1.70] | +40.3% | [+39.7%, +41.0%] | 1.0 pts | 1.8 | yes | yes |
| natural | street | 65536 | churn | ordered | ordered-mk1 | 8 | 227 | 229 | 1.00× [0.99, 1.01] | +0.1% | [-0.9%, +1.1%] | 2.7 pts | 4.6 | yes | no |
| natural | street | 65536 | churn | ordered | ordered-mk2 | 8 | 229 | 264 | 1.14× [1.12, 1.17] | +12.3% | [+10.5%, +14.2%] | 2.0 pts | 2.4 | yes | yes |
| natural | street | 65536 | churn | ordered | ordered-mk4 | 8 | 226 | 250 | 1.11× [1.09, 1.14] | +10.2% | [+8.1%, +12.4%] | 1.6 pts | 2.8 | no | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 154.44 ms | 119.48 ms | 0.77× [0.76, 0.78] | -29.5% | [-31.2%, -27.7%] | 1.7 pts | 1.3 | yes | yes |
| natural | street | 65536 | build | ordered | btree-sets | 8 | 152.86 ms | 253.67 ms | 1.63× [1.60, 1.67] | +38.8% | [+37.4%, +40.1%] | 1.6 pts | 3.9 | yes | yes |
| natural | street | 65536 | build | ordered | ordered-mk1 | 8 | 154.26 ms | 133.66 ms | 0.87× [0.85, 0.88] | -15.6% | [-17.3%, -13.9%] | 1.9 pts | 2.1 | yes | yes |
| natural | street | 65536 | build | ordered | ordered-mk2 | 8 | 155.25 ms | 170.17 ms | 1.10× [1.08, 1.12] | +9.2% | [+7.7%, +10.7%] | 1.1 pts | 1.9 | yes | yes |
| natural | street | 65536 | build | ordered | ordered-mk4 | 8 | 154.15 ms | 170.42 ms | 1.10× [1.09, 1.12] | +9.4% | [+8.6%, +10.3%] | 0.8 pts | 1.6 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 44.9 | 35.8 | 0.80× [0.79, 0.82] | -24.7% | [-26.9%, -22.6%] | 1.8 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 44.9 | 156 | 3.44× [3.36, 3.53] | +71.0% | [+70.3%, +71.7%] | 0.6 pts | 2.9 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 45.4 | 37.6 | 0.83× [0.82, 0.84] | -20.5% | [-22.0%, -19.0%] | 1.8 pts | 1.6 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 45.2 | 38.5 | 0.85× [0.84, 0.85] | -17.9% | [-18.7%, -17.0%] | 0.6 pts | 0.6 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 45.2 | 38.5 | 0.85× [0.84, 0.87] | -17.5% | [-19.7%, -15.2%] | 1.4 pts | 1.3 | no | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 1768 | 2438 | 1.38× [1.32, 1.44] | +27.5% | [+24.4%, +30.7%] | 1.9 pts | 1.0 | no | yes |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 1768 | 7021 | 4.00× [3.95, 4.04] | +75.0% | [+74.7%, +75.3%] | 0.3 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 1746 | 2149 | 1.24× [1.20, 1.28] | +19.5% | [+16.9%, +22.0%] | 1.7 pts | 0.9 | no | yes |
| natural | u64 | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 1736 | 2023 | 1.16× [1.14, 1.19] | +14.1% | [+12.0%, +16.3%] | 2.1 pts | 1.2 | no | yes |
| natural | u64 | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 1730 | 2026 | 1.18× [1.16, 1.21] | +15.6% | [+13.9%, +17.3%] | 2.0 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 70.9 | 46.7 | 0.65× [0.64, 0.66] | -53.4% | [-55.7%, -51.0%] | 2.4 pts | 1.7 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 8 | 71.3 | 157 | 2.20× [2.17, 2.24] | +54.6% | [+54.0%, +55.3%] | 0.5 pts | 1.1 | yes | yes |
| natural | u64 | 4096 | churn | ordered | ordered-mk1 | 8 | 71.2 | 63.0 | 0.88× [0.87, 0.90] | -13.0% | [-14.5%, -11.6%] | 1.1 pts | 1.4 | yes | yes |
| natural | u64 | 4096 | churn | ordered | ordered-mk2 | 8 | 71.2 | 67.6 | 0.95× [0.94, 0.96] | -5.3% | [-6.0%, -4.6%] | 0.9 pts | 1.2 | yes | yes |
| natural | u64 | 4096 | churn | ordered | ordered-mk4 | 8 | 71.5 | 67.0 | 0.94× [0.94, 0.95] | -6.2% | [-6.7%, -5.7%] | 0.5 pts | 0.7 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 9.10 ms | 4.76 ms | 0.52× [0.52, 0.53] | -90.9% | [-92.2%, -89.6%] | 1.4 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 8 | 9.23 ms | 14.92 ms | 1.62× [1.62, 1.63] | +38.4% | [+38.2%, +38.6%] | 0.5 pts | 1.0 | yes | yes |
| natural | u64 | 4096 | build | ordered | ordered-mk1 | 8 | 9.12 ms | 6.37 ms | 0.70× [0.70, 0.70] | -43.0% | [-43.6%, -42.4%] | 0.7 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | build | ordered | ordered-mk2 | 8 | 9.08 ms | 7.00 ms | 0.77× [0.76, 0.77] | -30.0% | [-31.0%, -29.0%] | 0.9 pts | 1.4 | yes | yes |
| natural | u64 | 4096 | build | ordered | ordered-mk4 | 8 | 9.08 ms | 7.06 ms | 0.78× [0.77, 0.78] | -29.0% | [-30.3%, -27.7%] | 1.0 pts | 1.3 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 8 | 47.0 | 43.1 | 0.92× [0.90, 0.93] | -9.3% | [-11.1%, -7.4%] | 1.6 pts | 2.7 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 8 | 47.2 | 191 | 4.03× [3.99, 4.08] | +75.2% | [+75.0%, +75.5%] | 0.2 pts | 1.5 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 47.5 | 46.0 | 0.97× [0.96, 0.98] | -3.2% | [-4.4%, -1.9%] | 1.0 pts | 1.1 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 47.7 | 47.6 | 1.00× [0.99, 1.01] | +0.0% | [-1.0%, +1.1%] | 0.9 pts | 1.2 | yes | no |
| natural | u64 | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 47.5 | 48.1 | 1.01× [0.99, 1.04] | +1.3% | [-1.2%, +3.7%] | 1.8 pts | 2.6 | no | no |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 2728 | 3408 | 1.25× [1.25, 1.25] | +20.1% | [+20.0%, +20.3%] | 0.2 pts | 0.3 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 8 | 2736 | 7475 | 2.72× [2.70, 2.75] | +63.3% | [+63.0%, +63.6%] | 0.3 pts | 1.4 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 2680 | 3097 | 1.16× [1.14, 1.17] | +13.6% | [+12.5%, +14.7%] | 0.8 pts | 1.3 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 2650 | 2722 | 1.03× [1.01, 1.04] | +2.5% | [+1.4%, +3.5%] | 0.8 pts | 1.5 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 2641 | 2651 | 1.01× [1.00, 1.02] | +0.7% | [-0.2%, +1.5%] | 0.6 pts | 1.3 | yes | no |
| natural | u64 | 16384 | churn | ordered | baseline | 8 | 88.7 | 57.3 | 0.64× [0.63, 0.65] | -55.9% | [-58.6%, -53.3%] | 2.6 pts | 1.1 | yes | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 8 | 96.0 | 234 | 2.44× [2.39, 2.50] | +59.0% | [+58.1%, +59.9%] | 0.8 pts | 1.6 | yes | yes |
| natural | u64 | 16384 | churn | ordered | ordered-mk1 | 8 | 90.1 | 82.5 | 0.92× [0.91, 0.92] | -9.2% | [-10.0%, -8.4%] | 0.5 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | churn | ordered | ordered-mk2 | 8 | 91.6 | 92.7 | 1.01× [1.00, 1.02] | +1.1% | [+0.0%, +2.1%] | 0.8 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | churn | ordered | ordered-mk4 | 8 | 88.5 | 89.8 | 1.01× [1.00, 1.02] | +0.9% | [+0.3%, +1.5%] | 0.7 pts | 1.6 | yes | yes |
| natural | u64 | 16384 | build | ordered | baseline | 8 | 35.66 ms | 22.88 ms | 0.64× [0.64, 0.64] | -56.2% | [-56.8%, -55.5%] | 1.1 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 8 | 35.97 ms | 80.11 ms | 2.22× [2.18, 2.26] | +55.0% | [+54.2%, +55.8%] | 0.6 pts | 1.0 | yes | yes |
| natural | u64 | 16384 | build | ordered | ordered-mk1 | 8 | 35.27 ms | 31.09 ms | 0.88× [0.88, 0.89] | -13.6% | [-14.2%, -12.9%] | 0.5 pts | 0.7 | yes | yes |
| natural | u64 | 16384 | build | ordered | ordered-mk2 | 8 | 35.18 ms | 34.66 ms | 0.98× [0.97, 0.99] | -1.8% | [-2.7%, -1.0%] | 0.6 pts | 1.1 | yes | yes |
| natural | u64 | 16384 | build | ordered | ordered-mk4 | 8 | 35.23 ms | 34.29 ms | 0.97× [0.97, 0.98] | -2.7% | [-2.9%, -2.5%] | 0.7 pts | 1.0 | yes | yes |
| natural | u64 | 65536 | valuesFor | ordered | baseline | 8 | 66.8 | 56.5 | 0.86× [0.82, 0.90] | -16.8% | [-21.9%, -11.7%] | 4.5 pts | 2.9 | no | yes |
| natural | u64 | 65536 | valuesFor | ordered | btree-sets | 8 | 84.1 | 313 | 3.59× [3.35, 3.88] | +72.2% | [+70.1%, +74.2%] | 1.7 pts | 3.5 | yes | yes |
| natural | u64 | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 63.3 | 55.4 | 0.88× [0.86, 0.90] | -14.0% | [-16.6%, -11.4%] | 4.2 pts | 3.7 | no | yes |
| natural | u64 | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 63.7 | 61.0 | 0.96× [0.93, 0.99] | -4.7% | [-7.9%, -1.4%] | 3.1 pts | 3.0 | no | yes |
| natural | u64 | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 65.3 | 63.8 | 0.98× [0.94, 1.03] | -1.6% | [-6.1%, +2.9%] | 4.9 pts | 4.7 | no | no |
| natural | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 2564 | 4163 | 1.63× [1.60, 1.65] | +38.5% | [+37.7%, +39.4%] | 1.2 pts | 2.1 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | btree-sets | 8 | 3045 | 11.8 µs | 3.88× [3.78, 3.98] | +74.2% | [+73.6%, +74.9%] | 0.6 pts | 1.3 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 2497 | 3699 | 1.48× [1.46, 1.51] | +32.7% | [+31.7%, +33.6%] | 1.1 pts | 2.8 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 2468 | 2977 | 1.20× [1.19, 1.22] | +16.8% | [+15.6%, +17.9%] | 0.9 pts | 1.7 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 2439 | 2465 | 1.01× [1.00, 1.02] | +1.0% | [+0.2%, +1.8%] | 1.4 pts | 2.2 | yes | yes |
| natural | u64 | 65536 | churn | ordered | baseline | 8 | 166 | 131 | 0.80× [0.77, 0.82] | -25.5% | [-29.6%, -21.5%] | 3.0 pts | 1.9 | no | yes |
| natural | u64 | 65536 | churn | ordered | btree-sets | 8 | 218 | 413 | 1.88× [1.76, 2.02] | +46.9% | [+43.3%, +50.6%] | 2.6 pts | 4.0 | yes | yes |
| natural | u64 | 65536 | churn | ordered | ordered-mk1 | 8 | 171 | 168 | 0.98× [0.96, 0.99] | -2.4% | [-4.1%, -0.7%] | 1.5 pts | 1.3 | yes | yes |
| natural | u64 | 65536 | churn | ordered | ordered-mk2 | 8 | 170 | 172 | 1.01× [0.99, 1.03] | +0.9% | [-1.0%, +2.8%] | 1.5 pts | 1.2 | yes | no |
| natural | u64 | 65536 | churn | ordered | ordered-mk4 | 8 | 159 | 161 | 1.02× [1.00, 1.03] | +1.6% | [+0.3%, +2.8%] | 1.3 pts | 1.0 | yes | yes |
| natural | u64 | 65536 | build | ordered | baseline | 8 | 205.39 ms | 141.58 ms | 0.69× [0.66, 0.72] | -45.7% | [-52.1%, -39.2%] | 4.7 pts | 2.3 | no | yes |
| natural | u64 | 65536 | build | ordered | btree-sets | 8 | 204.85 ms | 506.47 ms | 2.46× [2.43, 2.49] | +59.4% | [+58.9%, +59.8%] | 0.7 pts | 2.1 | yes | yes |
| natural | u64 | 65536 | build | ordered | ordered-mk1 | 8 | 204.71 ms | 187.97 ms | 0.91× [0.90, 0.92] | -10.0% | [-11.6%, -8.5%] | 1.7 pts | 1.0 | yes | yes |
| natural | u64 | 65536 | build | ordered | ordered-mk2 | 8 | 205.14 ms | 211.66 ms | 1.02× [1.00, 1.03] | +1.9% | [+0.5%, +3.3%] | 1.5 pts | 1.5 | yes | yes |
| natural | u64 | 65536 | build | ordered | ordered-mk4 | 8 | 203.93 ms | 207.50 ms | 1.02× [1.01, 1.03] | +1.8% | [+0.8%, +2.7%] | 0.7 pts | 1.0 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | baseline | 8 | 96.0 | 83.7 | 0.87× [0.86, 0.88] | -14.9% | [-16.8%, -13.0%] | 1.3 pts | 1.3 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | btree-sets | 8 | 98.2 | 189 | 1.92× [1.89, 1.95] | +48.0% | [+47.2%, +48.8%] | 0.6 pts | 1.3 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 100.0 | 85.6 | 0.85× [0.84, 0.87] | -17.1% | [-19.5%, -14.8%] | 2.1 pts | 1.9 | no | yes |
| natural | url | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 101 | 94.7 | 0.94× [0.92, 0.96] | -6.2% | [-8.3%, -4.2%] | 1.4 pts | 1.4 | no | yes |
| natural | url | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 101 | 98.2 | 0.98× [0.95, 1.00] | -2.4% | [-5.1%, +0.2%] | 1.9 pts | 2.4 | no | no |
| natural | url | 4096 | valuesBetween | ordered | baseline | 8 | 2102 | 3657 | 1.76× [1.74, 1.79] | +43.3% | [+42.6%, +44.0%] | 0.7 pts | 0.5 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | btree-sets | 8 | 2095 | 7711 | 3.72× [3.66, 3.78] | +73.1% | [+72.7%, +73.5%] | 0.5 pts | 1.0 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 2071 | 3703 | 1.81× [1.75, 1.87] | +44.7% | [+42.9%, +46.6%] | 1.6 pts | 1.4 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 2048 | 3075 | 1.51× [1.45, 1.56] | +33.6% | [+31.2%, +36.1%] | 1.9 pts | 1.5 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 2018 | 2350 | 1.17× [1.15, 1.20] | +14.8% | [+12.9%, +16.6%] | 1.4 pts | 0.9 | yes | yes |
| natural | url | 4096 | churn | ordered | baseline | 8 | 180 | 128 | 0.72× [0.71, 0.72] | -39.5% | [-40.7%, -38.3%] | 0.8 pts | 0.6 | yes | yes |
| natural | url | 4096 | churn | ordered | btree-sets | 8 | 180 | 220 | 1.24× [1.21, 1.28] | +19.6% | [+17.2%, +21.9%] | 1.9 pts | 1.6 | no | yes |
| natural | url | 4096 | churn | ordered | ordered-mk1 | 8 | 180 | 142 | 0.79× [0.79, 0.80] | -25.8% | [-27.0%, -24.6%] | 1.2 pts | 1.3 | yes | yes |
| natural | url | 4096 | churn | ordered | ordered-mk2 | 8 | 178 | 181 | 1.01× [1.01, 1.01] | +1.0% | [+0.6%, +1.4%] | 0.9 pts | 1.3 | yes | yes |
| natural | url | 4096 | churn | ordered | ordered-mk4 | 8 | 178 | 184 | 1.03× [1.02, 1.04] | +2.8% | [+2.1%, +3.4%] | 1.1 pts | 1.9 | yes | yes |
| natural | url | 4096 | build | ordered | baseline | 8 | 17.93 ms | 11.68 ms | 0.65× [0.64, 0.66] | -53.4% | [-55.2%, -51.5%] | 1.6 pts | 1.5 | yes | yes |
| natural | url | 4096 | build | ordered | btree-sets | 8 | 17.98 ms | 20.73 ms | 1.15× [1.12, 1.18] | +13.1% | [+11.1%, +15.1%] | 1.7 pts | 1.8 | no | yes |
| natural | url | 4096 | build | ordered | ordered-mk1 | 8 | 17.73 ms | 13.33 ms | 0.75× [0.75, 0.76] | -33.1% | [-34.1%, -32.0%] | 0.7 pts | 1.2 | yes | yes |
| natural | url | 4096 | build | ordered | ordered-mk2 | 8 | 17.72 ms | 17.55 ms | 0.99× [0.98, 1.00] | -0.8% | [-1.7%, -0.0%] | 0.7 pts | 1.0 | yes | yes |
| natural | url | 4096 | build | ordered | ordered-mk4 | 8 | 17.70 ms | 18.16 ms | 1.03× [1.02, 1.03] | +2.8% | [+2.3%, +3.2%] | 0.6 pts | 1.2 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | baseline | 8 | 119 | 108 | 0.90× [0.90, 0.91] | -10.8% | [-11.1%, -10.5%] | 0.6 pts | 0.8 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | btree-sets | 8 | 123 | 246 | 2.00× [1.98, 2.03] | +50.1% | [+49.4%, +50.7%] | 0.5 pts | 1.2 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 125 | 111 | 0.89× [0.89, 0.90] | -12.1% | [-12.8%, -11.4%] | 1.1 pts | 1.0 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 125 | 120 | 0.96× [0.95, 0.97] | -4.0% | [-4.9%, -3.0%] | 0.9 pts | 1.2 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 124 | 123 | 0.99× [0.97, 1.01] | -1.2% | [-3.0%, +0.6%] | 1.2 pts | 1.7 | yes | no |
| natural | url | 16384 | valuesBetween | ordered | baseline | 8 | 2649 | 4445 | 1.68× [1.67, 1.69] | +40.5% | [+40.1%, +40.8%] | 0.4 pts | 0.8 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | btree-sets | 8 | 2643 | 8220 | 3.10× [3.06, 3.14] | +67.7% | [+67.3%, +68.1%] | 0.4 pts | 1.7 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 2626 | 4421 | 1.69× [1.68, 1.70] | +40.8% | [+40.5%, +41.1%] | 0.3 pts | 0.8 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 2599 | 3689 | 1.42× [1.42, 1.42] | +29.5% | [+29.4%, +29.6%] | 0.3 pts | 0.8 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 2589 | 2994 | 1.16× [1.15, 1.17] | +13.7% | [+13.1%, +14.2%] | 0.5 pts | 1.1 | yes | yes |
| natural | url | 16384 | churn | ordered | baseline | 8 | 231 | 183 | 0.79× [0.78, 0.80] | -26.8% | [-28.1%, -25.5%] | 1.8 pts | 1.8 | yes | yes |
| natural | url | 16384 | churn | ordered | btree-sets | 8 | 230 | 329 | 1.42× [1.39, 1.46] | +29.8% | [+28.3%, +31.3%] | 1.3 pts | 2.1 | yes | yes |
| natural | url | 16384 | churn | ordered | ordered-mk1 | 8 | 233 | 201 | 0.86× [0.85, 0.86] | -16.8% | [-17.9%, -15.8%] | 1.5 pts | 1.7 | yes | yes |
| natural | url | 16384 | churn | ordered | ordered-mk2 | 8 | 229 | 239 | 1.03× [1.03, 1.04] | +3.3% | [+2.8%, +3.8%] | 0.8 pts | 1.2 | yes | yes |
| natural | url | 16384 | churn | ordered | ordered-mk4 | 8 | 224 | 232 | 1.03× [1.03, 1.04] | +3.1% | [+2.8%, +3.5%] | 0.5 pts | 1.3 | yes | yes |
| natural | url | 16384 | build | ordered | baseline | 8 | 85.51 ms | 60.68 ms | 0.71× [0.70, 0.72] | -40.5% | [-42.3%, -38.7%] | 1.4 pts | 0.9 | yes | yes |
| natural | url | 16384 | build | ordered | btree-sets | 8 | 85.57 ms | 110.70 ms | 1.30× [1.28, 1.32] | +23.0% | [+21.6%, +24.3%] | 1.0 pts | 1.6 | yes | yes |
| natural | url | 16384 | build | ordered | ordered-mk1 | 8 | 84.65 ms | 66.82 ms | 0.79× [0.77, 0.80] | -27.2% | [-29.7%, -24.6%] | 1.5 pts | 1.1 | yes | yes |
| natural | url | 16384 | build | ordered | ordered-mk2 | 8 | 84.77 ms | 85.88 ms | 1.01× [1.01, 1.02] | +1.3% | [+0.6%, +2.1%] | 0.9 pts | 1.5 | yes | yes |
| natural | url | 16384 | build | ordered | ordered-mk4 | 8 | 84.55 ms | 86.65 ms | 1.03× [1.02, 1.04] | +2.8% | [+2.0%, +3.5%] | 0.7 pts | 1.2 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | baseline | 8 | 187 | 208 | 1.13× [1.10, 1.17] | +11.8% | [+8.7%, +14.9%] | 3.4 pts | 2.7 | no | yes |
| natural | url | 65536 | valuesFor | ordered | btree-sets | 8 | 214 | 452 | 2.08× [2.04, 2.11] | +51.9% | [+51.1%, +52.6%] | 0.9 pts | 1.2 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 177 | 176 | 0.99× [0.95, 1.04] | -0.8% | [-5.4%, +3.8%] | 2.9 pts | 2.6 | no | no |
| natural | url | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 174 | 180 | 1.04× [0.97, 1.11] | +3.8% | [-2.6%, +10.2%] | 3.9 pts | 4.0 | no | no |
| natural | url | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 173 | 177 | 1.02× [0.98, 1.06] | +1.9% | [-1.6%, +5.4%] | 3.2 pts | 3.2 | no | no |
| natural | url | 65536 | valuesBetween | ordered | baseline | 8 | 3361 | 5715 | 1.70× [1.66, 1.74] | +41.0% | [+39.6%, +42.4%] | 1.7 pts | 2.1 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | btree-sets | 8 | 4153 | 14.2 µs | 3.38× [3.31, 3.46] | +70.4% | [+69.8%, +71.1%] | 0.6 pts | 1.1 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 3375 | 5460 | 1.63× [1.63, 1.64] | +38.8% | [+38.5%, +39.1%] | 0.6 pts | 0.7 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 3233 | 4597 | 1.41× [1.34, 1.47] | +28.9% | [+25.6%, +32.1%] | 2.2 pts | 2.7 | no | yes |
| natural | url | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 3176 | 3682 | 1.16× [1.13, 1.19] | +13.9% | [+11.9%, +15.9%] | 2.6 pts | 2.9 | no | yes |
| natural | url | 65536 | churn | ordered | baseline | 8 | 380 | 385 | 1.01× [0.98, 1.05] | +1.2% | [-2.4%, +4.8%] | 2.2 pts | 2.6 | no | no |
| natural | url | 65536 | churn | ordered | btree-sets | 8 | 448 | 595 | 1.33× [1.30, 1.36] | +25.0% | [+23.3%, +26.7%] | 2.0 pts | 3.2 | yes | yes |
| natural | url | 65536 | churn | ordered | ordered-mk1 | 8 | 398 | 390 | 0.99× [0.95, 1.03] | -1.2% | [-5.3%, +3.0%] | 3.0 pts | 2.3 | no | no |
| natural | url | 65536 | churn | ordered | ordered-mk2 | 8 | 380 | 404 | 1.07× [1.04, 1.09] | +6.1% | [+3.6%, +8.6%] | 1.6 pts | 1.1 | no | yes |
| natural | url | 65536 | churn | ordered | ordered-mk4 | 8 | 368 | 389 | 1.04× [1.03, 1.06] | +4.3% | [+2.6%, +6.0%] | 3.2 pts | 1.6 | yes | yes |
| natural | url | 65536 | build | ordered | baseline | 8 | 509.41 ms | 451.93 ms | 0.89× [0.88, 0.90] | -12.7% | [-14.3%, -11.1%] | 1.0 pts | 0.9 | yes | yes |
| natural | url | 65536 | build | ordered | btree-sets | 8 | 508.22 ms | 749.36 ms | 1.47× [1.45, 1.50] | +32.0% | [+30.8%, +33.2%] | 1.2 pts | 1.9 | yes | yes |
| natural | url | 65536 | build | ordered | ordered-mk1 | 8 | 512.51 ms | 452.68 ms | 0.88× [0.87, 0.90] | -13.3% | [-15.3%, -11.3%] | 1.5 pts | 1.4 | yes | yes |
| natural | url | 65536 | build | ordered | ordered-mk2 | 8 | 502.51 ms | 528.17 ms | 1.04× [1.03, 1.05] | +4.0% | [+3.0%, +5.1%] | 0.9 pts | 1.0 | yes | yes |
| natural | url | 65536 | build | ordered | ordered-mk4 | 8 | 509.65 ms | 526.23 ms | 1.03× [1.02, 1.04] | +2.7% | [+2.1%, +3.4%] | 1.0 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 8 | 84.9 | 72.1 | 0.85× [0.83, 0.86] | -18.0% | [-20.0%, -15.9%] | 2.5 pts | 2.8 | no | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 86.0 | 108 | 1.26× [1.23, 1.28] | +20.4% | [+18.9%, +22.0%] | 1.2 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 88.4 | 78.2 | 0.89× [0.88, 0.90] | -12.3% | [-14.0%, -10.7%] | 1.8 pts | 2.0 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 88.5 | 84.3 | 0.95× [0.94, 0.96] | -5.2% | [-6.1%, -4.2%] | 1.0 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 88.5 | 86.0 | 0.97× [0.97, 0.98] | -2.8% | [-3.3%, -2.3%] | 0.9 pts | 1.3 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 1009 | 1931 | 1.92× [1.89, 1.94] | +47.8% | [+47.0%, +48.5%] | 1.1 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 987 | 549 | 0.56× [0.54, 0.57] | -79.3% | [-84.3%, -74.2%] | 3.1 pts | 1.4 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 1012 | 2861 | 2.85× [2.81, 2.89] | +64.9% | [+64.4%, +65.3%] | 0.3 pts | 0.6 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 1004 | 2204 | 2.22× [2.18, 2.27] | +55.0% | [+54.1%, +55.9%] | 0.7 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 988 | 1618 | 1.64× [1.60, 1.68] | +38.9% | [+37.5%, +40.4%] | 1.0 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 8 | 207 | 193 | 0.92× [0.92, 0.93] | -8.2% | [-9.0%, -7.5%] | 1.8 pts | 1.3 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 8 | 207 | 184 | 0.89× [0.87, 0.91] | -12.4% | [-14.4%, -10.4%] | 1.9 pts | 1.0 | no | yes |
| single-value | dirs | 4096 | churn | ordered | ordered-mk1 | 8 | 206 | 210 | 1.01× [0.99, 1.03] | +0.9% | [-1.2%, +3.1%] | 2.2 pts | 1.9 | no | no |
| single-value | dirs | 4096 | churn | ordered | ordered-mk2 | 8 | 207 | 292 | 1.42× [1.40, 1.43] | +29.3% | [+28.8%, +29.9%] | 0.6 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | ordered-mk4 | 8 | 203 | 262 | 1.29× [1.28, 1.30] | +22.5% | [+21.8%, +23.1%] | 0.7 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 8 | 3.54 ms | 2.70 ms | 0.77× [0.75, 0.79] | -30.7% | [-34.2%, -27.3%] | 2.5 pts | 0.7 | no | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 8 | 3.59 ms | 2.76 ms | 0.77× [0.75, 0.80] | -29.2% | [-33.2%, -25.2%] | 3.2 pts | 1.1 | no | yes |
| single-value | dirs | 4096 | build | ordered | ordered-mk1 | 8 | 3.51 ms | 3.07 ms | 0.88× [0.86, 0.89] | -14.2% | [-15.9%, -12.6%] | 2.1 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | build | ordered | ordered-mk2 | 8 | 3.43 ms | 4.54 ms | 1.32× [1.29, 1.35] | +24.4% | [+22.5%, +26.2%] | 1.7 pts | 1.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | ordered-mk4 | 8 | 3.38 ms | 4.15 ms | 1.23× [1.21, 1.24] | +18.5% | [+17.3%, +19.6%] | 1.3 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 8 | 106 | 103 | 0.96× [0.95, 0.98] | -3.7% | [-5.6%, -1.9%] | 1.3 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 8 | 107 | 155 | 1.45× [1.40, 1.50] | +30.8% | [+28.5%, +33.2%] | 1.8 pts | 2.2 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 111 | 107 | 0.96× [0.96, 0.97] | -3.6% | [-4.6%, -2.7%] | 1.7 pts | 1.8 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 111 | 111 | 1.00× [0.99, 1.01] | +0.2% | [-0.7%, +1.1%] | 0.9 pts | 1.0 | yes | no |
| single-value | dirs | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 110 | 110 | 1.00× [0.97, 1.02] | -0.3% | [-3.1%, +2.4%] | 1.8 pts | 2.0 | no | no |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 1130 | 2548 | 2.26× [2.22, 2.31] | +55.8% | [+54.9%, +56.7%] | 0.6 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 8 | 1120 | 719 | 0.65× [0.64, 0.65] | -54.6% | [-56.1%, -53.1%] | 1.5 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 1126 | 3365 | 3.00× [2.95, 3.04] | +66.6% | [+66.1%, +67.1%] | 0.4 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 1116 | 2658 | 2.38× [2.35, 2.40] | +57.9% | [+57.5%, +58.4%] | 0.5 pts | 1.5 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 1108 | 2018 | 1.83× [1.81, 1.85] | +45.3% | [+44.8%, +45.9%] | 0.5 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 8 | 246 | 273 | 1.09× [1.07, 1.11] | +8.2% | [+6.8%, +9.6%] | 1.4 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 8 | 246 | 250 | 1.02× [1.00, 1.03] | +1.6% | [-0.1%, +3.4%] | 1.6 pts | 1.1 | yes | no |
| single-value | dirs | 16384 | churn | ordered | ordered-mk1 | 8 | 245 | 268 | 1.10× [1.08, 1.11] | +8.7% | [+7.1%, +10.3%] | 1.1 pts | 1.3 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | ordered-mk2 | 8 | 247 | 356 | 1.45× [1.44, 1.46] | +31.0% | [+30.6%, +31.5%] | 0.5 pts | 0.9 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | ordered-mk4 | 8 | 245 | 330 | 1.34× [1.33, 1.36] | +25.6% | [+25.0%, +26.3%] | 0.7 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 8 | 15.61 ms | 13.73 ms | 0.88× [0.87, 0.89] | -13.7% | [-15.5%, -11.9%] | 1.3 pts | 0.5 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 8 | 16.01 ms | 14.32 ms | 0.89× [0.87, 0.90] | -12.9% | [-15.1%, -10.7%] | 3.0 pts | 1.2 | no | yes |
| single-value | dirs | 16384 | build | ordered | ordered-mk1 | 8 | 15.95 ms | 15.11 ms | 0.95× [0.93, 0.97] | -5.7% | [-7.9%, -3.6%] | 3.1 pts | 1.5 | no | yes |
| single-value | dirs | 16384 | build | ordered | ordered-mk2 | 8 | 15.94 ms | 21.93 ms | 1.40× [1.36, 1.44] | +28.3% | [+26.3%, +30.3%] | 1.9 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | build | ordered | ordered-mk4 | 8 | 15.71 ms | 19.97 ms | 1.28× [1.24, 1.32] | +22.0% | [+19.6%, +24.4%] | 2.1 pts | 1.3 | no | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 142 | 148 | 1.07× [1.00, 1.15] | +6.5% | [-0.0%, +13.1%] | 4.3 pts | 4.1 | no | no |
| single-value | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 138 | 215 | 1.56× [1.54, 1.59] | +36.0% | [+35.0%, +37.0%] | 0.7 pts | 0.9 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 142 | 146 | 1.03× [1.01, 1.05] | +2.6% | [+0.6%, +4.6%] | 2.2 pts | 2.2 | no | yes |
| single-value | dirs | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 142 | 146 | 1.03× [1.03, 1.04] | +3.4% | [+2.7%, +4.0%] | 1.0 pts | 1.0 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 141 | 144 | 1.02× [1.00, 1.04] | +2.3% | [+0.4%, +4.1%] | 1.5 pts | 1.5 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1237 | 2818 | 2.29× [2.28, 2.30] | +56.3% | [+56.1%, +56.5%] | 0.2 pts | 0.4 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 1223 | 897 | 0.74× [0.73, 0.74] | -36.0% | [-36.5%, -35.4%] | 1.5 pts | 1.3 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1225 | 3626 | 2.98× [2.94, 3.02] | +66.4% | [+66.0%, +66.9%] | 0.3 pts | 1.2 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1221 | 2920 | 2.41× [2.38, 2.43] | +58.5% | [+58.0%, +58.9%] | 0.4 pts | 1.2 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1214 | 2260 | 1.86× [1.85, 1.88] | +46.3% | [+45.9%, +46.7%] | 0.3 pts | 1.0 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 319 | 402 | 1.26× [1.25, 1.28] | +20.7% | [+19.8%, +21.7%] | 0.9 pts | 1.0 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | btree-map | 8 | 343 | 376 | 1.10× [1.05, 1.15] | +9.0% | [+4.8%, +13.2%] | 2.9 pts | 3.1 | no | yes |
| single-value | dirs | 65536 | churn | ordered | ordered-mk1 | 8 | 339 | 397 | 1.20× [1.18, 1.21] | +16.5% | [+15.4%, +17.7%] | 2.1 pts | 2.2 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | ordered-mk2 | 8 | 319 | 466 | 1.45× [1.43, 1.48] | +31.2% | [+30.1%, +32.2%] | 0.8 pts | 1.1 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | ordered-mk4 | 8 | 322 | 448 | 1.38× [1.35, 1.41] | +27.5% | [+26.1%, +28.8%] | 1.3 pts | 1.4 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 83.79 ms | 87.12 ms | 1.03× [1.02, 1.05] | +3.3% | [+2.2%, +4.5%] | 1.6 pts | 1.0 | yes | yes |
| single-value | dirs | 65536 | build | ordered | btree-map | 8 | 84.61 ms | 83.73 ms | 0.99× [0.98, 1.00] | -1.0% | [-2.6%, +0.5%] | 1.4 pts | 1.0 | yes | no |
| single-value | dirs | 65536 | build | ordered | ordered-mk1 | 8 | 84.45 ms | 86.31 ms | 1.02× [1.00, 1.04] | +1.7% | [-0.4%, +3.9%] | 1.8 pts | 1.2 | no | no |
| single-value | dirs | 65536 | build | ordered | ordered-mk2 | 8 | 83.77 ms | 112.35 ms | 1.35× [1.33, 1.38] | +26.0% | [+24.6%, +27.4%] | 1.3 pts | 1.5 | yes | yes |
| single-value | dirs | 65536 | build | ordered | ordered-mk4 | 8 | 84.73 ms | 109.55 ms | 1.29× [1.26, 1.31] | +22.2% | [+20.6%, +23.8%] | 1.6 pts | 2.0 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 57.2 | 46.5 | 0.81× [0.78, 0.84] | -23.9% | [-28.5%, -19.3%] | 2.9 pts | 3.6 | no | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 57.5 | 90.8 | 1.58× [1.57, 1.58] | +36.6% | [+36.4%, +36.9%] | 0.4 pts | 0.8 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 59.4 | 49.0 | 0.82× [0.81, 0.84] | -21.5% | [-23.3%, -19.6%] | 1.5 pts | 2.3 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 59.2 | 53.7 | 0.91× [0.90, 0.91] | -10.4% | [-10.7%, -10.1%] | 0.6 pts | 0.8 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 59.4 | 55.4 | 0.93× [0.92, 0.95] | -7.1% | [-8.7%, -5.4%] | 1.0 pts | 1.4 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 606 | 1702 | 2.83× [2.78, 2.87] | +64.6% | [+64.0%, +65.2%] | 0.4 pts | 0.7 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 601 | 502 | 0.84× [0.83, 0.86] | -18.9% | [-21.0%, -16.8%] | 2.2 pts | 1.5 | no | yes |
| single-value | street | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 609 | 2601 | 4.31× [4.26, 4.36] | +76.8% | [+76.5%, +77.1%] | 0.3 pts | 0.8 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 610 | 2030 | 3.37× [3.35, 3.39] | +70.3% | [+70.1%, +70.5%] | 0.3 pts | 0.9 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 606 | 1513 | 2.52× [2.47, 2.57] | +60.4% | [+59.6%, +61.2%] | 0.6 pts | 1.1 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 141 | 113 | 0.81× [0.79, 0.82] | -24.1% | [-26.3%, -21.9%] | 1.7 pts | 1.3 | yes | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 140 | 142 | 1.01× [0.99, 1.02] | +0.9% | [-0.5%, +2.4%] | 1.3 pts | 1.3 | yes | no |
| single-value | street | 4096 | churn | ordered | ordered-mk1 | 8 | 141 | 132 | 0.94× [0.93, 0.94] | -6.8% | [-7.6%, -5.9%] | 1.1 pts | 0.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | ordered-mk2 | 8 | 140 | 195 | 1.39× [1.37, 1.40] | +27.8% | [+26.9%, +28.8%] | 0.7 pts | 1.1 | yes | yes |
| single-value | street | 4096 | churn | ordered | ordered-mk4 | 8 | 141 | 195 | 1.37× [1.34, 1.40] | +27.2% | [+25.6%, +28.8%] | 1.0 pts | 1.4 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.34 ms | 1.69 ms | 0.73× [0.70, 0.77] | -36.5% | [-43.1%, -30.0%] | 4.8 pts | 1.3 | no | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.36 ms | 2.16 ms | 0.93× [0.89, 0.96] | -8.0% | [-12.0%, -4.0%] | 3.0 pts | 1.4 | no | yes |
| single-value | street | 4096 | build | ordered | ordered-mk1 | 8 | 2.33 ms | 2.02 ms | 0.86× [0.84, 0.87] | -16.6% | [-18.8%, -14.4%] | 4.5 pts | 1.7 | no | yes |
| single-value | street | 4096 | build | ordered | ordered-mk2 | 8 | 2.32 ms | 3.15 ms | 1.37× [1.26, 1.49] | +26.9% | [+20.9%, +32.9%] | 3.7 pts | 3.2 | no | yes |
| single-value | street | 4096 | build | ordered | ordered-mk4 | 8 | 2.36 ms | 3.13 ms | 1.34× [1.30, 1.37] | +25.1% | [+23.1%, +27.1%] | 1.6 pts | 1.1 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 68.6 | 64.2 | 0.93× [0.92, 0.94] | -7.3% | [-8.7%, -5.8%] | 1.1 pts | 1.2 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 68.9 | 125 | 1.80× [1.74, 1.86] | +44.5% | [+42.6%, +46.4%] | 1.2 pts | 2.7 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 71.5 | 68.1 | 0.94× [0.90, 0.98] | -6.4% | [-10.7%, -2.2%] | 3.3 pts | 3.7 | no | yes |
| single-value | street | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 71.2 | 71.7 | 1.01× [0.99, 1.03] | +0.9% | [-1.1%, +2.8%] | 1.4 pts | 1.6 | yes | no |
| single-value | street | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 71.2 | 72.2 | 1.01× [1.00, 1.02] | +0.9% | [+0.0%, +1.8%] | 1.1 pts | 1.4 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 786 | 2157 | 2.75× [2.74, 2.76] | +63.6% | [+63.5%, +63.7%] | 0.3 pts | 0.8 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 777 | 612 | 0.79× [0.76, 0.81] | -27.3% | [-31.4%, -23.1%] | 3.2 pts | 2.8 | no | yes |
| single-value | street | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 787 | 2999 | 3.83× [3.73, 3.92] | +73.9% | [+73.2%, +74.5%] | 0.5 pts | 1.9 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 783 | 2365 | 3.04× [3.01, 3.07] | +67.1% | [+66.8%, +67.5%] | 0.3 pts | 0.7 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 774 | 1777 | 2.31× [2.29, 2.32] | +56.6% | [+56.3%, +56.9%] | 0.4 pts | 1.3 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 167 | 159 | 0.93× [0.90, 0.97] | -7.3% | [-11.3%, -3.2%] | 2.6 pts | 1.9 | no | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 166 | 196 | 1.17× [1.16, 1.18] | +14.6% | [+13.7%, +15.4%] | 1.2 pts | 1.5 | yes | yes |
| single-value | street | 16384 | churn | ordered | ordered-mk1 | 8 | 169 | 174 | 1.04× [1.02, 1.06] | +3.7% | [+1.9%, +5.5%] | 1.4 pts | 1.4 | yes | yes |
| single-value | street | 16384 | churn | ordered | ordered-mk2 | 8 | 172 | 249 | 1.45× [1.42, 1.47] | +30.9% | [+29.7%, +32.2%] | 1.0 pts | 1.5 | yes | yes |
| single-value | street | 16384 | churn | ordered | ordered-mk4 | 8 | 169 | 231 | 1.38× [1.37, 1.39] | +27.6% | [+27.0%, +28.1%] | 0.9 pts | 1.3 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 10.84 ms | 8.29 ms | 0.77× [0.75, 0.78] | -30.3% | [-32.5%, -28.2%] | 2.8 pts | 0.7 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 11.24 ms | 11.80 ms | 1.07× [1.02, 1.12] | +6.6% | [+2.1%, +11.1%] | 3.0 pts | 1.5 | no | yes |
| single-value | street | 16384 | build | ordered | ordered-mk1 | 8 | 10.89 ms | 9.66 ms | 0.91× [0.89, 0.93] | -10.2% | [-12.6%, -7.8%] | 4.2 pts | 1.2 | no | yes |
| single-value | street | 16384 | build | ordered | ordered-mk2 | 8 | 10.54 ms | 15.18 ms | 1.41× [1.36, 1.47] | +29.2% | [+26.7%, +31.8%] | 2.0 pts | 1.4 | yes | yes |
| single-value | street | 16384 | build | ordered | ordered-mk4 | 8 | 10.78 ms | 14.61 ms | 1.36× [1.31, 1.41] | +26.3% | [+23.4%, +29.2%] | 2.2 pts | 1.1 | no | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 85.1 | 87.5 | 1.05× [1.03, 1.08] | +5.0% | [+2.9%, +7.0%] | 7.1 pts | 7.9 | no | yes |
| single-value | street | 65536 | valuesFor | ordered | btree-map | 8 | 86.8 | 178 | 2.06× [2.03, 2.08] | +51.4% | [+50.8%, +52.0%] | 0.8 pts | 2.5 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 88.7 | 91.2 | 1.02× [0.99, 1.04] | +1.8% | [-0.5%, +4.2%] | 1.8 pts | 2.3 | no | no |
| single-value | street | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 88.9 | 94.7 | 1.06× [1.05, 1.08] | +6.0% | [+4.3%, +7.6%] | 1.6 pts | 2.3 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 88.4 | 94.0 | 1.06× [1.06, 1.06] | +5.5% | [+5.3%, +5.7%] | 1.5 pts | 1.9 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 838 | 2402 | 2.86× [2.85, 2.88] | +65.1% | [+64.9%, +65.3%] | 0.3 pts | 1.2 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | btree-map | 8 | 834 | 746 | 0.89× [0.88, 0.90] | -12.3% | [-13.4%, -11.2%] | 1.7 pts | 1.7 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 837 | 3265 | 3.90× [3.88, 3.91] | +74.3% | [+74.2%, +74.4%] | 0.1 pts | 0.6 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 835 | 2561 | 3.07× [3.07, 3.08] | +67.5% | [+67.4%, +67.6%] | 0.1 pts | 0.6 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 830 | 1950 | 2.35× [2.33, 2.37] | +57.4% | [+57.0%, +57.8%] | 0.3 pts | 1.3 | yes | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 206 | 232 | 1.14× [1.12, 1.15] | +11.9% | [+10.5%, +13.4%] | 1.6 pts | 1.7 | yes | yes |
| single-value | street | 65536 | churn | ordered | btree-map | 8 | 209 | 265 | 1.27× [1.23, 1.32] | +21.4% | [+18.8%, +24.1%] | 1.7 pts | 1.9 | no | yes |
| single-value | street | 65536 | churn | ordered | ordered-mk1 | 8 | 204 | 240 | 1.19× [1.16, 1.21] | +15.7% | [+13.9%, +17.5%] | 1.9 pts | 1.7 | yes | yes |
| single-value | street | 65536 | churn | ordered | ordered-mk2 | 8 | 203 | 304 | 1.49× [1.46, 1.51] | +32.8% | [+31.7%, +33.8%] | 0.7 pts | 0.9 | yes | yes |
| single-value | street | 65536 | churn | ordered | ordered-mk4 | 8 | 207 | 284 | 1.40× [1.39, 1.41] | +28.4% | [+27.9%, +28.9%] | 1.3 pts | 1.7 | yes | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 54.37 ms | 51.31 ms | 0.94× [0.91, 0.98] | -6.3% | [-10.2%, -2.5%] | 2.4 pts | 1.0 | no | yes |
| single-value | street | 65536 | build | ordered | btree-map | 8 | 55.05 ms | 62.07 ms | 1.13× [1.12, 1.14] | +11.5% | [+10.5%, +12.5%] | 2.1 pts | 1.3 | yes | yes |
| single-value | street | 65536 | build | ordered | ordered-mk1 | 8 | 53.27 ms | 54.37 ms | 1.02× [1.00, 1.04] | +1.6% | [-0.4%, +3.5%] | 1.9 pts | 1.3 | yes | no |
| single-value | street | 65536 | build | ordered | ordered-mk2 | 8 | 54.40 ms | 76.26 ms | 1.39× [1.37, 1.41] | +28.1% | [+27.0%, +29.2%] | 1.5 pts | 1.4 | yes | yes |
| single-value | street | 65536 | build | ordered | ordered-mk4 | 8 | 53.63 ms | 72.88 ms | 1.37× [1.36, 1.38] | +26.9% | [+26.2%, +27.5%] | 1.4 pts | 1.2 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 8 | 32.8 | 16.8 | 0.51× [0.51, 0.52] | -94.7% | [-97.0%, -92.4%] | 2.4 pts | 4.7 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 8 | 32.5 | 87.3 | 2.69× [2.68, 2.70] | +62.8% | [+62.7%, +62.9%] | 0.1 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 33.9 | 19.9 | 0.59× [0.59, 0.59] | -69.8% | [-70.6%, -68.9%] | 0.6 pts | 1.1 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 34.1 | 20.4 | 0.60× [0.59, 0.60] | -67.5% | [-68.5%, -66.4%] | 1.0 pts | 1.8 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 34.1 | 20.4 | 0.60× [0.60, 0.60] | -67.3% | [-67.8%, -66.8%] | 0.9 pts | 1.4 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 279 | 1033 | 3.69× [3.66, 3.72] | +72.9% | [+72.7%, +73.1%] | 0.2 pts | 1.7 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 8 | 278 | 437 | 1.57× [1.54, 1.59] | +36.2% | [+35.2%, +37.3%] | 0.9 pts | 4.4 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 283 | 1547 | 5.46× [5.44, 5.48] | +81.7% | [+81.6%, +81.7%] | 0.1 pts | 1.2 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 283 | 1484 | 5.24× [5.21, 5.26] | +80.9% | [+80.8%, +81.0%] | 0.1 pts | 0.9 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 284 | 1475 | 5.21× [5.19, 5.22] | +80.8% | [+80.7%, +80.9%] | 0.1 pts | 1.2 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 8 | 79.0 | 45.8 | 0.58× [0.55, 0.62] | -71.0% | [-80.3%, -61.7%] | 7.1 pts | 3.2 | no | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 8 | 78.9 | 123 | 1.56× [1.54, 1.59] | +36.0% | [+35.0%, +37.0%] | 0.7 pts | 1.7 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | ordered-mk1 | 8 | 78.8 | 54.3 | 0.69× [0.67, 0.72] | -44.0% | [-48.5%, -39.5%] | 3.0 pts | 2.2 | no | yes |
| single-value | u64 | 4096 | churn | ordered | ordered-mk2 | 8 | 79.0 | 67.9 | 0.86× [0.86, 0.87] | -15.9% | [-16.7%, -15.1%] | 0.8 pts | 0.8 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | ordered-mk4 | 8 | 79.1 | 67.3 | 0.85× [0.84, 0.86] | -17.4% | [-19.2%, -15.7%] | 1.3 pts | 1.4 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 8 | 1.18 ms | 866.6 µs | 0.76× [0.67, 0.88] | -30.9% | [-48.5%, -13.3%] | 13.4 pts | 1.4 | no | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 8 | 1.19 ms | 1.82 ms | 1.53× [1.51, 1.55] | +34.7% | [+33.8%, +35.6%] | 1.1 pts | 1.6 | yes | yes |
| single-value | u64 | 4096 | build | ordered | ordered-mk1 | 8 | 1.18 ms | 985.0 µs | 0.84× [0.77, 0.91] | -19.7% | [-29.6%, -9.8%] | 6.2 pts | 1.4 | no | yes |
| single-value | u64 | 4096 | build | ordered | ordered-mk2 | 8 | 1.18 ms | 1.31 ms | 1.10× [1.05, 1.14] | +8.8% | [+5.1%, +12.4%] | 3.1 pts | 1.1 | no | yes |
| single-value | u64 | 4096 | build | ordered | ordered-mk4 | 8 | 1.18 ms | 1.42 ms | 1.21× [1.14, 1.28] | +17.1% | [+12.4%, +21.8%] | 3.5 pts | 1.0 | no | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 8 | 26.6 | 22.6 | 0.84× [0.83, 0.86] | -18.5% | [-20.9%, -16.2%] | 2.0 pts | 5.0 | no | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 26.0 | 111 | 4.30× [4.24, 4.36] | +76.7% | [+76.4%, +77.1%] | 0.5 pts | 3.6 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 26.8 | 25.2 | 0.94× [0.92, 0.96] | -6.5% | [-8.3%, -4.7%] | 2.1 pts | 3.6 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 26.7 | 26.2 | 0.97× [0.94, 1.00] | -3.2% | [-6.0%, -0.4%] | 2.4 pts | 4.3 | no | yes |
| single-value | u64 | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 26.8 | 26.8 | 0.99× [0.97, 1.00] | -1.2% | [-2.9%, +0.5%] | 2.6 pts | 6.7 | yes | no |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 1716 | 1592 | 0.94× [0.93, 0.94] | -6.9% | [-7.8%, -6.1%] | 1.3 pts | 0.7 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 1688 | 470 | 0.28× [0.28, 0.28] | -258.3% | [-261.6%, -255.0%] | 3.5 pts | 1.7 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 1721 | 2066 | 1.20× [1.18, 1.21] | +16.5% | [+15.6%, +17.4%] | 0.8 pts | 0.8 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 1673 | 1718 | 1.03× [1.02, 1.03] | +2.5% | [+2.3%, +2.6%] | 0.3 pts | 0.7 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 1668 | 1680 | 1.00× [1.00, 1.00] | +0.2% | [+0.0%, +0.5%] | 0.5 pts | 1.2 | yes | no |
| single-value | u64 | 16384 | churn | ordered | baseline | 8 | 74.6 | 52.8 | 0.72× [0.68, 0.76] | -38.8% | [-46.4%, -31.2%] | 7.1 pts | 4.2 | no | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 8 | 79.0 | 173 | 2.19× [2.11, 2.28] | +54.4% | [+52.5%, +56.2%] | 1.5 pts | 2.1 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | ordered-mk1 | 8 | 74.9 | 66.8 | 0.90× [0.88, 0.91] | -11.7% | [-13.5%, -10.0%] | 1.4 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | ordered-mk2 | 8 | 77.6 | 80.2 | 1.03× [1.02, 1.04] | +3.0% | [+2.3%, +3.6%] | 1.2 pts | 1.0 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | ordered-mk4 | 8 | 76.9 | 77.6 | 1.01× [0.99, 1.03] | +0.7% | [-1.0%, +2.5%] | 1.7 pts | 1.6 | yes | no |
| single-value | u64 | 16384 | build | ordered | baseline | 8 | 5.78 ms | 3.73 ms | 0.65× [0.61, 0.71] | -52.8% | [-64.7%, -40.9%] | 12.2 pts | 2.0 | no | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 8 | 5.81 ms | 9.18 ms | 1.57× [1.52, 1.62] | +36.2% | [+34.0%, +38.4%] | 2.4 pts | 1.7 | yes | yes |
| single-value | u64 | 16384 | build | ordered | ordered-mk1 | 8 | 5.79 ms | 4.26 ms | 0.75× [0.73, 0.78] | -32.7% | [-36.9%, -28.4%] | 2.9 pts | 1.1 | no | yes |
| single-value | u64 | 16384 | build | ordered | ordered-mk2 | 8 | 5.78 ms | 5.11 ms | 0.87× [0.85, 0.89] | -14.7% | [-17.5%, -11.9%] | 2.8 pts | 1.5 | no | yes |
| single-value | u64 | 16384 | build | ordered | ordered-mk4 | 8 | 5.81 ms | 4.87 ms | 0.84× [0.83, 0.86] | -18.4% | [-20.3%, -16.4%] | 1.5 pts | 0.7 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | baseline | 8 | 35.5 | 25.3 | 0.71× [0.70, 0.72] | -40.7% | [-42.3%, -39.1%] | 1.0 pts | 1.0 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | btree-map | 8 | 34.5 | 154 | 4.45× [4.36, 4.56] | +77.5% | [+77.1%, +78.0%] | 0.5 pts | 3.0 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 35.5 | 29.6 | 0.84× [0.81, 0.86] | -19.8% | [-23.3%, -16.3%] | 2.5 pts | 3.7 | no | yes |
| single-value | u64 | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 35.8 | 33.9 | 0.95× [0.93, 0.96] | -5.7% | [-7.6%, -3.9%] | 1.5 pts | 2.9 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 35.8 | 35.7 | 1.00× [0.99, 1.01] | -0.1% | [-1.1%, +1.0%] | 1.9 pts | 2.8 | yes | no |
| single-value | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 1491 | 1969 | 1.32× [1.31, 1.33] | +24.2% | [+23.7%, +24.6%] | 0.4 pts | 1.2 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | btree-map | 8 | 1495 | 592 | 0.40× [0.39, 0.40] | -152.0% | [-154.2%, -149.8%] | 2.1 pts | 2.0 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1492 | 2572 | 1.73× [1.72, 1.73] | +42.1% | [+42.0%, +42.1%] | 0.1 pts | 0.5 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1494 | 1901 | 1.27× [1.27, 1.28] | +21.5% | [+21.3%, +21.8%] | 0.3 pts | 1.3 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1481 | 1514 | 1.02× [1.02, 1.03] | +2.2% | [+1.7%, +2.6%] | 0.4 pts | 1.2 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | baseline | 8 | 112 | 92.6 | 0.83× [0.79, 0.87] | -20.5% | [-25.9%, -15.1%] | 3.6 pts | 1.7 | no | yes |
| single-value | u64 | 65536 | churn | ordered | btree-map | 8 | 118 | 235 | 1.98× [1.87, 2.11] | +49.6% | [+46.5%, +52.6%] | 3.1 pts | 3.1 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | ordered-mk1 | 8 | 116 | 113 | 0.96× [0.94, 0.97] | -4.6% | [-6.4%, -2.8%] | 2.5 pts | 2.4 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | ordered-mk2 | 8 | 117 | 137 | 1.17× [1.16, 1.18] | +14.8% | [+14.0%, +15.6%] | 2.2 pts | 1.9 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | ordered-mk4 | 8 | 117 | 121 | 1.02× [1.01, 1.03] | +2.1% | [+1.2%, +3.1%] | 1.2 pts | 0.9 | yes | yes |
| single-value | u64 | 65536 | build | ordered | baseline | 8 | 28.51 ms | 21.52 ms | 0.76× [0.74, 0.79] | -31.3% | [-35.9%, -26.7%] | 5.6 pts | 1.2 | no | yes |
| single-value | u64 | 65536 | build | ordered | btree-map | 8 | 29.70 ms | 52.76 ms | 1.76× [1.72, 1.79] | +43.1% | [+42.0%, +44.2%] | 1.5 pts | 1.0 | yes | yes |
| single-value | u64 | 65536 | build | ordered | ordered-mk1 | 8 | 28.67 ms | 25.28 ms | 0.87× [0.82, 0.93] | -14.8% | [-22.0%, -7.7%] | 6.6 pts | 1.8 | no | yes |
| single-value | u64 | 65536 | build | ordered | ordered-mk2 | 8 | 27.68 ms | 32.07 ms | 1.13× [1.12, 1.14] | +11.7% | [+10.9%, +12.5%] | 2.9 pts | 1.1 | yes | yes |
| single-value | u64 | 65536 | build | ordered | ordered-mk4 | 8 | 28.59 ms | 27.72 ms | 0.99× [0.96, 1.01] | -1.4% | [-4.0%, +1.2%] | 3.0 pts | 1.4 | no | no |
| single-value | url | 4096 | valuesFor | ordered | baseline | 8 | 75.2 | 63.5 | 0.84× [0.82, 0.86] | -18.9% | [-21.5%, -16.4%] | 1.7 pts | 1.9 | no | yes |
| single-value | url | 4096 | valuesFor | ordered | btree-map | 8 | 76.3 | 111 | 1.46× [1.43, 1.49] | +31.3% | [+29.9%, +32.7%] | 1.0 pts | 1.7 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 78.2 | 65.5 | 0.84× [0.82, 0.85] | -19.5% | [-21.5%, -17.5%] | 1.8 pts | 2.3 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 78.3 | 74.4 | 0.95× [0.94, 0.95] | -5.6% | [-6.4%, -4.8%] | 1.1 pts | 1.6 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 78.3 | 76.7 | 0.98× [0.97, 0.98] | -2.3% | [-2.8%, -1.7%] | 1.1 pts | 1.8 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | baseline | 8 | 1164 | 1864 | 1.60× [1.57, 1.63] | +37.6% | [+36.4%, +38.7%] | 1.1 pts | 0.9 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | btree-map | 8 | 1124 | 557 | 0.49× [0.48, 0.51] | -103.1% | [-110.2%, -95.9%] | 5.6 pts | 1.9 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 1160 | 2738 | 2.39× [2.35, 2.43] | +58.2% | [+57.4%, +58.9%] | 0.5 pts | 0.8 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 1151 | 2153 | 1.90× [1.86, 1.95] | +47.4% | [+46.1%, +48.6%] | 0.8 pts | 1.0 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 1124 | 1583 | 1.41× [1.39, 1.44] | +29.1% | [+27.8%, +30.5%] | 0.9 pts | 0.9 | yes | yes |
| single-value | url | 4096 | churn | ordered | baseline | 8 | 210 | 171 | 0.81× [0.80, 0.83] | -22.8% | [-24.9%, -20.8%] | 2.6 pts | 1.0 | yes | yes |
| single-value | url | 4096 | churn | ordered | btree-map | 8 | 209 | 183 | 0.88× [0.87, 0.89] | -13.6% | [-14.7%, -12.4%] | 1.1 pts | 0.8 | yes | yes |
| single-value | url | 4096 | churn | ordered | ordered-mk1 | 8 | 210 | 182 | 0.87× [0.85, 0.89] | -15.4% | [-18.0%, -12.8%] | 2.4 pts | 1.3 | no | yes |
| single-value | url | 4096 | churn | ordered | ordered-mk2 | 8 | 209 | 259 | 1.24× [1.23, 1.25] | +19.2% | [+18.6%, +19.9%] | 1.5 pts | 1.4 | yes | yes |
| single-value | url | 4096 | churn | ordered | ordered-mk4 | 8 | 209 | 242 | 1.15× [1.13, 1.18] | +13.3% | [+11.3%, +15.4%] | 1.6 pts | 1.4 | no | yes |
| single-value | url | 4096 | build | ordered | baseline | 8 | 3.29 ms | 2.32 ms | 0.70× [0.67, 0.73] | -42.8% | [-49.1%, -36.4%] | 7.1 pts | 1.5 | no | yes |
| single-value | url | 4096 | build | ordered | btree-map | 8 | 3.38 ms | 2.57 ms | 0.77× [0.75, 0.80] | -29.3% | [-33.0%, -25.6%] | 3.4 pts | 1.3 | no | yes |
| single-value | url | 4096 | build | ordered | ordered-mk1 | 8 | 3.45 ms | 2.66 ms | 0.78× [0.77, 0.79] | -27.9% | [-30.1%, -25.8%] | 2.6 pts | 0.6 | yes | yes |
| single-value | url | 4096 | build | ordered | ordered-mk2 | 8 | 3.31 ms | 3.83 ms | 1.17× [1.14, 1.19] | +14.2% | [+12.5%, +15.9%] | 1.3 pts | 0.9 | yes | yes |
| single-value | url | 4096 | build | ordered | ordered-mk4 | 8 | 3.28 ms | 3.64 ms | 1.13× [1.12, 1.15] | +11.9% | [+10.5%, +13.2%] | 2.1 pts | 1.2 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | baseline | 8 | 95.5 | 87.8 | 0.92× [0.91, 0.93] | -8.9% | [-10.1%, -7.7%] | 1.8 pts | 2.2 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | btree-map | 8 | 97.1 | 160 | 1.65× [1.63, 1.67] | +39.3% | [+38.6%, +40.0%] | 1.3 pts | 2.3 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 99.2 | 88.8 | 0.89× [0.88, 0.90] | -12.2% | [-13.8%, -10.6%] | 1.8 pts | 1.9 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 99.8 | 96.9 | 0.97× [0.96, 0.98] | -2.8% | [-3.8%, -1.8%] | 2.6 pts | 3.2 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 98.9 | 97.6 | 0.98× [0.98, 0.99] | -1.6% | [-2.4%, -0.7%] | 2.1 pts | 3.3 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | baseline | 8 | 1413 | 2384 | 1.69× [1.66, 1.71] | +40.7% | [+39.7%, +41.6%] | 0.7 pts | 1.2 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | btree-map | 8 | 1391 | 729 | 0.52× [0.52, 0.53] | -91.3% | [-92.4%, -90.2%] | 1.4 pts | 1.0 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 1406 | 3200 | 2.30× [2.27, 2.33] | +56.4% | [+55.9%, +57.0%] | 0.5 pts | 1.3 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 1395 | 2549 | 1.82× [1.81, 1.83] | +45.2% | [+44.9%, +45.5%] | 0.3 pts | 0.8 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 1392 | 1948 | 1.41× [1.40, 1.41] | +28.9% | [+28.4%, +29.3%] | 0.5 pts | 0.9 | yes | yes |
| single-value | url | 16384 | churn | ordered | baseline | 8 | 263 | 238 | 0.90× [0.88, 0.93] | -10.5% | [-13.9%, -7.1%] | 2.6 pts | 1.7 | no | yes |
| single-value | url | 16384 | churn | ordered | btree-map | 8 | 264 | 254 | 0.97× [0.96, 0.98] | -3.2% | [-4.2%, -2.3%] | 0.9 pts | 0.6 | yes | yes |
| single-value | url | 16384 | churn | ordered | ordered-mk1 | 8 | 261 | 239 | 0.91× [0.91, 0.92] | -9.3% | [-9.7%, -9.0%] | 2.4 pts | 1.5 | yes | yes |
| single-value | url | 16384 | churn | ordered | ordered-mk2 | 8 | 262 | 325 | 1.25× [1.23, 1.26] | +19.9% | [+19.0%, +20.7%] | 1.3 pts | 1.3 | yes | yes |
| single-value | url | 16384 | churn | ordered | ordered-mk4 | 8 | 266 | 310 | 1.16× [1.14, 1.19] | +14.1% | [+12.2%, +16.1%] | 1.5 pts | 1.7 | yes | yes |
| single-value | url | 16384 | build | ordered | baseline | 8 | 15.97 ms | 11.50 ms | 0.72× [0.70, 0.74] | -39.4% | [-43.0%, -35.9%] | 3.4 pts | 0.8 | yes | yes |
| single-value | url | 16384 | build | ordered | btree-map | 8 | 16.51 ms | 14.16 ms | 0.86× [0.86, 0.87] | -15.9% | [-16.5%, -15.3%] | 1.7 pts | 0.8 | yes | yes |
| single-value | url | 16384 | build | ordered | ordered-mk1 | 8 | 16.11 ms | 12.72 ms | 0.79× [0.78, 0.80] | -26.2% | [-27.4%, -25.1%] | 2.7 pts | 0.6 | yes | yes |
| single-value | url | 16384 | build | ordered | ordered-mk2 | 8 | 16.18 ms | 19.09 ms | 1.18× [1.12, 1.24] | +15.0% | [+10.6%, +19.4%] | 2.9 pts | 1.3 | no | yes |
| single-value | url | 16384 | build | ordered | ordered-mk4 | 8 | 16.29 ms | 18.32 ms | 1.13× [1.10, 1.17] | +11.8% | [+9.4%, +14.2%] | 1.8 pts | 0.8 | no | yes |
| single-value | url | 65536 | valuesFor | ordered | baseline | 8 | 132 | 148 | 1.15× [1.11, 1.21] | +13.3% | [+9.5%, +17.1%] | 7.6 pts | 5.3 | no | yes |
| single-value | url | 65536 | valuesFor | ordered | btree-map | 8 | 132 | 217 | 1.61× [1.58, 1.64] | +38.0% | [+36.8%, +39.2%] | 2.0 pts | 2.9 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 131 | 127 | 1.01× [0.98, 1.04] | +0.9% | [-2.1%, +4.0%] | 6.5 pts | 4.6 | no | no |
| single-value | url | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 130 | 137 | 1.05× [1.02, 1.08] | +4.7% | [+1.6%, +7.8%] | 3.5 pts | 3.1 | no | yes |
| single-value | url | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 134 | 136 | 1.05× [1.03, 1.08] | +4.8% | [+2.7%, +7.0%] | 2.5 pts | 2.2 | no | yes |
| single-value | url | 65536 | valuesBetween | ordered | baseline | 8 | 1503 | 2677 | 1.76× [1.75, 1.78] | +43.3% | [+42.9%, +43.8%] | 1.0 pts | 1.9 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | btree-map | 8 | 1493 | 900 | 0.61× [0.59, 0.62] | -64.9% | [-68.1%, -61.7%] | 2.4 pts | 1.4 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1524 | 3439 | 2.29× [2.24, 2.33] | +56.3% | [+55.4%, +57.1%] | 1.1 pts | 2.9 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1494 | 2770 | 1.84× [1.82, 1.86] | +45.7% | [+45.2%, +46.3%] | 0.8 pts | 2.4 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1465 | 2119 | 1.43× [1.42, 1.45] | +30.3% | [+29.6%, +31.0%] | 1.1 pts | 2.7 | yes | yes |
| single-value | url | 65536 | churn | ordered | baseline | 8 | 392 | 421 | 1.05× [1.03, 1.06] | +4.3% | [+3.1%, +5.6%] | 2.3 pts | 1.8 | yes | yes |
| single-value | url | 65536 | churn | ordered | btree-map | 8 | 404 | 402 | 1.00× [0.97, 1.05] | +0.4% | [-3.6%, +4.4%] | 2.7 pts | 2.1 | no | no |
| single-value | url | 65536 | churn | ordered | ordered-mk1 | 8 | 370 | 374 | 1.01× [1.00, 1.02] | +1.0% | [+0.1%, +1.9%] | 1.0 pts | 0.7 | yes | yes |
| single-value | url | 65536 | churn | ordered | ordered-mk2 | 8 | 363 | 448 | 1.25× [1.22, 1.28] | +19.8% | [+17.8%, +21.7%] | 1.5 pts | 1.4 | yes | yes |
| single-value | url | 65536 | churn | ordered | ordered-mk4 | 8 | 396 | 476 | 1.18× [1.16, 1.19] | +15.0% | [+13.9%, +16.1%] | 1.5 pts | 1.1 | yes | yes |
| single-value | url | 65536 | build | ordered | baseline | 8 | 89.78 ms | 77.24 ms | 0.83× [0.82, 0.85] | -19.8% | [-21.7%, -18.0%] | 5.2 pts | 2.2 | yes | yes |
| single-value | url | 65536 | build | ordered | btree-map | 8 | 88.92 ms | 84.18 ms | 0.93× [0.88, 0.98] | -7.7% | [-13.8%, -1.5%] | 4.3 pts | 2.8 | no | yes |
| single-value | url | 65536 | build | ordered | ordered-mk1 | 8 | 91.43 ms | 78.50 ms | 0.85× [0.84, 0.86] | -17.8% | [-19.6%, -16.0%] | 3.8 pts | 2.6 | yes | yes |
| single-value | url | 65536 | build | ordered | ordered-mk2 | 8 | 89.07 ms | 100.83 ms | 1.15× [1.12, 1.18] | +12.9% | [+10.5%, +15.3%] | 1.7 pts | 1.7 | no | yes |
| single-value | url | 65536 | build | ordered | ordered-mk4 | 8 | 87.74 ms | 97.83 ms | 1.12× [1.10, 1.13] | +10.4% | [+9.5%, +11.3%] | 1.0 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.80% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 build: ordered vs ordered-mk2: the A/A validations found a systematic difference of -0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 build: ordered vs ordered-mk2: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-4.28%, 0.15%] includes zero
- natural dirs n=16384 churn: ordered vs ordered-mk1: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs ordered-mk2: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs ordered-mk2: the pooled difference of 0.12% does not clear the 0.44% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 valuesFor: ordered vs ordered-mk2: the pooled interval [-1.86%, 2.10%] includes zero
- natural dirs n=65536 valuesFor: ordered vs ordered-mk2: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural dirs n=65536 valuesFor: ordered vs ordered-mk4: the pooled interval [-3.20%, 0.97%] includes zero
- natural dirs n=65536 churn: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 build: ordered vs ordered-mk2: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesFor: ordered vs ordered-mk1: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 churn: ordered vs ordered-mk2: the pooled difference of 0.32% does not clear the 0.42% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=4096 churn: ordered vs ordered-mk2: the pooled interval [-1.01%, 1.65%] includes zero
- natural street n=4096 build: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 build: ordered vs ordered-mk1: the A/A validations found a systematic difference of -0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 build: ordered vs ordered-mk4: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs ordered-mk1: the A/A validations found a systematic difference of +0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=16384 valuesFor: ordered vs ordered-mk1: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs ordered-mk1: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: the pooled interval [-7.29%, 0.36%] includes zero
- natural street n=65536 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs ordered-mk2: the pooled difference of 0.74% does not clear the 0.96% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 valuesFor: ordered vs ordered-mk2: the pooled interval [-2.23%, 3.71%] includes zero
- natural street n=65536 valuesFor: ordered vs ordered-mk4: the pooled difference of -0.46% does not clear the 0.79% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 valuesFor: ordered vs ordered-mk4: the pooled interval [-2.36%, 1.44%] includes zero
- natural street n=65536 churn: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs ordered-mk1: the pooled difference of 0.10% does not clear the 0.37% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 churn: ordered vs ordered-mk1: the pooled interval [-0.92%, 1.11%] includes zero
- natural street n=65536 churn: ordered vs ordered-mk1: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs ordered-mk1: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural street n=65536 churn: ordered vs ordered-mk2: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs ordered-mk4: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs btree-sets: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs ordered-mk1: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesFor: ordered vs ordered-mk2: the pooled difference of 0.04% does not clear the 0.42% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=16384 valuesFor: ordered vs ordered-mk2: the pooled interval [-1.01%, 1.08%] includes zero
- natural u64 n=16384 valuesFor: ordered vs ordered-mk4: the A/A validations found a systematic difference of -0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-1.18%, 3.70%] includes zero
- natural u64 n=16384 valuesFor: ordered vs ordered-mk4: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesFor: ordered vs ordered-mk4: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=16384 valuesBetween: ordered vs ordered-mk4: the pooled interval [-0.20%, 1.53%] includes zero
- natural u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs ordered-mk2: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs ordered-mk4: the pooled interval [-6.12%, 2.87%] includes zero
- natural u64 n=65536 valuesFor: ordered vs ordered-mk4: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs ordered-mk4: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=65536 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesBetween: ordered vs ordered-mk1: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesBetween: ordered vs ordered-mk4: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesBetween: ordered vs ordered-mk4: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 churn: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 churn: ordered vs ordered-mk2: the pooled interval [-0.96%, 2.84%] includes zero
- natural u64 n=65536 churn: ordered vs ordered-mk2: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=65536 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 build: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 build: ordered vs ordered-mk4: the A/A validations found a systematic difference of -0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=4096 valuesFor: ordered vs ordered-mk4: the pooled interval [-5.10%, 0.21%] includes zero
- natural url n=4096 valuesFor: ordered vs ordered-mk4: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-3.00%, 0.63%] includes zero
- natural url n=16384 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesFor: ordered vs ordered-mk1: the pooled interval [-5.45%, 3.75%] includes zero
- natural url n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesFor: ordered vs ordered-mk1: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural url n=65536 valuesFor: ordered vs ordered-mk2: the pooled interval [-2.60%, 10.20%] includes zero
- natural url n=65536 valuesFor: ordered vs ordered-mk2: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesFor: ordered vs ordered-mk4: the pooled interval [-1.65%, 5.41%] includes zero
- natural url n=65536 valuesFor: ordered vs ordered-mk4: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesFor: ordered vs ordered-mk4: 4 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural url n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesBetween: ordered vs ordered-mk2: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesBetween: ordered vs ordered-mk4: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 churn: ordered vs baseline: the pooled interval [-2.41%, 4.81%] includes zero
- natural url n=65536 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 churn: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural url n=65536 churn: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 churn: ordered vs ordered-mk1: the A/A validations found a systematic difference of -0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 churn: ordered vs ordered-mk1: the pooled interval [-5.34%, 2.98%] includes zero
- natural url n=65536 churn: ordered vs ordered-mk1: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 churn: ordered vs ordered-mk1: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural url n=65536 build: ordered vs ordered-mk1: the A/A validations found a systematic difference of -0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 build: ordered vs ordered-mk4: the A/A validations found a systematic difference of +0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 valuesFor: ordered vs ordered-mk1: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 valuesFor: ordered vs ordered-mk2: the A/A validations found a systematic difference of -0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 churn: ordered vs ordered-mk1: the pooled interval [-1.24%, 3.05%] includes zero
- single-value dirs n=4096 churn: ordered vs ordered-mk1: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value dirs n=4096 churn: ordered vs ordered-mk4: the A/A validations found a systematic difference of -0.37% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs ordered-mk2: the pooled difference of 0.21% does not clear the 1.31% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 valuesFor: ordered vs ordered-mk2: the pooled interval [-0.69%, 1.10%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs ordered-mk4: the pooled difference of -0.35% does not clear the 0.41% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-3.10%, 2.41%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs ordered-mk4: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value dirs n=16384 churn: ordered vs btree-map: the pooled interval [-0.08%, 3.35%] includes zero
- single-value dirs n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -1.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=65536 valuesFor: ordered vs baseline: the pooled interval [-0.03%, 13.11%] includes zero
- single-value dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 churn: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 churn: ordered vs ordered-mk1: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: the pooled interval [-2.56%, 0.46%] includes zero
- single-value dirs n=65536 build: ordered vs ordered-mk1: the pooled interval [-0.39%, 3.86%] includes zero
- single-value street n=4096 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.24% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesFor: ordered vs ordered-mk1: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.94% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 churn: ordered vs btree-map: the pooled interval [-0.54%, 2.43%] includes zero
- single-value street n=4096 build: ordered vs ordered-mk2: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs ordered-mk1: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs ordered-mk2: the pooled interval [-1.05%, 2.79%] includes zero
- single-value street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs baseline: the processes scatter 7.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs ordered-mk1: the pooled interval [-0.54%, 4.18%] includes zero
- single-value street n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs ordered-mk2: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.34% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=65536 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.51% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=65536 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=65536 build: ordered vs ordered-mk1: the pooled interval [-0.43%, 3.53%] includes zero
- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesBetween: ordered vs btree-map: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs ordered-mk1: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs ordered-mk1: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=16384 valuesFor: ordered vs ordered-mk2: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-2.93%, 0.49%] includes zero
- single-value u64 n=16384 valuesFor: ordered vs ordered-mk4: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs ordered-mk4: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=16384 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.11% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=16384 valuesBetween: ordered vs ordered-mk4: the pooled difference of 0.24% does not clear the 0.29% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=16384 churn: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 churn: ordered vs ordered-mk4: the pooled difference of 0.74% does not clear the 0.86% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=16384 churn: ordered vs ordered-mk4: the pooled interval [-1.03%, 2.51%] includes zero
- single-value u64 n=16384 build: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.13% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=65536 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk2: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk4: the pooled difference of -0.09% does not clear the 0.27% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk4: the pooled interval [-1.13%, 0.95%] includes zero
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk4: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk4: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=65536 valuesBetween: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs ordered-mk1: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 build: ordered vs baseline: the A/A validations found a systematic difference of -1.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=65536 build: ordered vs ordered-mk4: the pooled difference of -1.37% does not clear the 1.90% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=65536 build: ordered vs ordered-mk4: the pooled interval [-3.96%, 1.21%] includes zero
- single-value url n=4096 valuesFor: ordered vs ordered-mk1: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=4096 churn: ordered vs ordered-mk4: the A/A validations found a systematic difference of -0.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=16384 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 valuesFor: ordered vs ordered-mk2: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 valuesFor: ordered vs ordered-mk2: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value url n=16384 valuesFor: ordered vs ordered-mk4: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 valuesFor: ordered vs ordered-mk4: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value url n=65536 valuesFor: ordered vs baseline: the processes scatter 5.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesFor: ordered vs ordered-mk1: the pooled interval [-2.10%, 3.98%] includes zero
- single-value url n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesFor: ordered vs ordered-mk1: 2 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value url n=65536 valuesFor: ordered vs ordered-mk2: the A/A validations found a systematic difference of -0.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=65536 valuesFor: ordered vs ordered-mk2: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesFor: ordered vs ordered-mk4: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesBetween: ordered vs ordered-mk1: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesBetween: ordered vs ordered-mk2: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesBetween: ordered vs ordered-mk4: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 churn: ordered vs btree-map: the pooled difference of 0.39% does not clear the 0.69% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 churn: ordered vs btree-map: the pooled interval [-3.58%, 4.37%] includes zero
- single-value url n=65536 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 churn: ordered vs btree-map: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value url n=65536 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 build: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 build: ordered vs ordered-mk1: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
