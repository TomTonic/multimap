| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 139 | 119 | 0.86× [0.85, 0.86] | -16.8% | [-17.7%, -16.0%] | 0.7 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 142 | 209 | 1.48× [1.44, 1.52] | +32.2% | [+30.4%, +34.1%] | 1.5 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 139 | 124 | 0.89× [0.88, 0.91] | -11.8% | [-14.2%, -9.4%] | 1.7 pts | 1.1 | no | yes |
| natural | dirs | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 139 | 133 | 0.96× [0.94, 0.97] | -4.6% | [-5.9%, -3.4%] | 1.0 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 139 | 135 | 0.97× [0.96, 0.99] | -3.0% | [-4.6%, -1.4%] | 1.1 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2060 | 3569 | 1.74× [1.71, 1.78] | +42.6% | [+41.4%, +43.8%] | 1.3 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2065 | 6290 | 3.04× [3.03, 3.05] | +67.1% | [+67.0%, +67.2%] | 0.4 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 2085 | 4842 | 2.35× [2.31, 2.39] | +57.5% | [+56.7%, +58.2%] | 0.8 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 2051 | 3846 | 1.89× [1.86, 1.93] | +47.2% | [+46.3%, +48.1%] | 0.9 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 2078 | 2941 | 1.42× [1.40, 1.44] | +29.7% | [+28.8%, +30.6%] | 1.0 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 244 | 188 | 0.78× [0.77, 0.78] | -29.0% | [-29.5%, -28.5%] | 0.6 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 244 | 274 | 1.12× [1.11, 1.12] | +10.6% | [+10.1%, +11.0%] | 0.6 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | churn | ordered | ordered-mk1 | 8 | 243 | 208 | 0.86× [0.86, 0.87] | -16.0% | [-16.8%, -15.1%] | 1.5 pts | 3.0 | yes | yes |
| natural | dirs | 4096 | churn | ordered | ordered-mk2 | 8 | 242 | 260 | 1.08× [1.07, 1.08] | +7.2% | [+6.8%, +7.6%] | 0.5 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | churn | ordered | ordered-mk4 | 8 | 241 | 254 | 1.06× [1.06, 1.07] | +5.7% | [+5.3%, +6.1%] | 0.8 pts | 2.5 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 11.86 ms | 8.54 ms | 0.72× [0.72, 0.72] | -38.8% | [-39.5%, -38.0%] | 0.5 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 11.85 ms | 13.10 ms | 1.10× [1.10, 1.11] | +9.4% | [+9.3%, +9.6%] | 0.3 pts | 1.3 | yes | yes |
| natural | dirs | 4096 | build | ordered | ordered-mk1 | 8 | 11.82 ms | 9.58 ms | 0.81× [0.81, 0.81] | -23.4% | [-24.0%, -22.7%] | 0.8 pts | 1.9 | yes | yes |
| natural | dirs | 4096 | build | ordered | ordered-mk2 | 8 | 11.77 ms | 12.58 ms | 1.07× [1.06, 1.08] | +6.4% | [+5.8%, +7.0%] | 0.5 pts | 1.9 | yes | yes |
| natural | dirs | 4096 | build | ordered | ordered-mk4 | 8 | 11.79 ms | 12.59 ms | 1.07× [1.06, 1.07] | +6.4% | [+6.0%, +6.9%] | 0.3 pts | 1.8 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 174 | 151 | 0.87× [0.87, 0.88] | -14.8% | [-15.6%, -14.0%] | 0.9 pts | 1.8 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 177 | 274 | 1.57× [1.46, 1.69] | +36.1% | [+31.4%, +40.9%] | 2.9 pts | 4.7 | no | yes |
| natural | dirs | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 174 | 159 | 0.92× [0.91, 0.92] | -9.2% | [-9.9%, -8.5%] | 1.7 pts | 2.1 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 174 | 166 | 0.95× [0.95, 0.96] | -4.8% | [-5.3%, -4.3%] | 0.4 pts | 0.6 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 174 | 168 | 0.97× [0.96, 0.98] | -3.2% | [-4.1%, -2.4%] | 0.6 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 2230 | 4076 | 1.83× [1.80, 1.85] | +45.2% | [+44.6%, +45.9%] | 0.5 pts | 0.9 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 2245 | 6851 | 3.06× [3.05, 3.07] | +67.3% | [+67.2%, +67.4%] | 0.5 pts | 2.0 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 2221 | 5201 | 2.35× [2.33, 2.37] | +57.4% | [+57.0%, +57.8%] | 0.3 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 2228 | 4218 | 1.90× [1.89, 1.91] | +47.4% | [+47.1%, +47.6%] | 0.3 pts | 0.7 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 2198 | 3233 | 1.46× [1.45, 1.48] | +31.6% | [+31.0%, +32.3%] | 0.5 pts | 0.8 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 293 | 255 | 0.87× [0.86, 0.88] | -15.0% | [-16.1%, -13.9%] | 1.2 pts | 2.1 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 309 | 378 | 1.23× [1.22, 1.25] | +18.8% | [+17.8%, +19.8%] | 1.2 pts | 1.4 | yes | yes |
| natural | dirs | 16384 | churn | ordered | ordered-mk1 | 8 | 288 | 262 | 0.92× [0.90, 0.94] | -8.8% | [-11.0%, -6.5%] | 2.1 pts | 6.6 | no | yes |
| natural | dirs | 16384 | churn | ordered | ordered-mk2 | 8 | 286 | 312 | 1.10× [1.09, 1.11] | +8.8% | [+8.0%, +9.6%] | 0.5 pts | 1.2 | yes | yes |
| natural | dirs | 16384 | churn | ordered | ordered-mk4 | 8 | 283 | 303 | 1.07× [1.06, 1.08] | +6.6% | [+5.9%, +7.3%] | 0.7 pts | 2.2 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 55.93 ms | 44.41 ms | 0.80× [0.79, 0.80] | -25.5% | [-26.6%, -24.5%] | 0.9 pts | 2.9 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 55.83 ms | 66.77 ms | 1.20× [1.19, 1.20] | +16.4% | [+15.8%, +16.9%] | 0.4 pts | 3.0 | yes | yes |
| natural | dirs | 16384 | build | ordered | ordered-mk1 | 8 | 55.94 ms | 48.01 ms | 0.86× [0.86, 0.86] | -16.3% | [-16.9%, -15.7%] | 1.4 pts | 4.7 | yes | yes |
| natural | dirs | 16384 | build | ordered | ordered-mk2 | 8 | 55.88 ms | 60.55 ms | 1.09× [1.08, 1.09] | +8.1% | [+7.6%, +8.5%] | 0.4 pts | 2.0 | yes | yes |
| natural | dirs | 16384 | build | ordered | ordered-mk4 | 8 | 55.80 ms | 60.15 ms | 1.08× [1.08, 1.08] | +7.6% | [+7.4%, +7.7%] | 0.3 pts | 1.3 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 237 | 269 | 1.14× [1.12, 1.17] | +12.5% | [+10.5%, +14.4%] | 1.4 pts | 0.6 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 260 | 498 | 1.97× [1.91, 2.03] | +49.2% | [+47.6%, +50.8%] | 1.3 pts | 0.9 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 231 | 249 | 1.12× [1.04, 1.20] | +10.5% | [+4.1%, +17.0%] | 4.9 pts | 2.4 | no | yes |
| natural | dirs | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 228 | 244 | 1.09× [1.03, 1.14] | +7.9% | [+3.1%, +12.6%] | 3.8 pts | 2.0 | no | yes |
| natural | dirs | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 227 | 238 | 1.05× [1.03, 1.06] | +4.5% | [+3.0%, +6.0%] | 1.4 pts | 0.8 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 2536 | 5617 | 2.21× [2.16, 2.26] | +54.8% | [+53.7%, +55.8%] | 0.8 pts | 0.9 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 2711 | 11.0 µs | 4.08× [4.03, 4.14] | +75.5% | [+75.2%, +75.8%] | 0.2 pts | 0.8 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 2500 | 6018 | 2.41× [2.36, 2.47] | +58.5% | [+57.6%, +59.5%] | 0.7 pts | 1.2 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 2501 | 4891 | 1.96× [1.93, 2.00] | +49.1% | [+48.1%, +50.1%] | 0.8 pts | 1.5 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 2479 | 3814 | 1.52× [1.50, 1.55] | +34.4% | [+33.4%, +35.4%] | 1.0 pts | 1.4 | yes | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 461 | 465 | 0.97× [0.94, 0.99] | -3.5% | [-6.2%, -0.8%] | 2.4 pts | 0.4 | no | yes |
| natural | dirs | 65536 | churn | ordered | btree-sets | 8 | 516 | 700 | 1.36× [1.34, 1.37] | +26.3% | [+25.5%, +27.1%] | 1.2 pts | 1.0 | yes | yes |
| natural | dirs | 65536 | churn | ordered | ordered-mk1 | 8 | 447 | 430 | 0.93× [0.90, 0.96] | -7.3% | [-10.5%, -4.0%] | 2.4 pts | 0.3 | no | yes |
| natural | dirs | 65536 | churn | ordered | ordered-mk2 | 8 | 452 | 475 | 1.03× [0.99, 1.06] | +2.4% | [-0.9%, +5.8%] | 2.2 pts | 0.3 | no | no |
| natural | dirs | 65536 | churn | ordered | ordered-mk4 | 8 | 446 | 459 | 0.99× [0.96, 1.02] | -1.2% | [-4.6%, +2.3%] | 2.2 pts | 0.3 | no | no |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 326.19 ms | 318.04 ms | 0.97× [0.96, 0.99] | -2.8% | [-4.1%, -1.5%] | 1.1 pts | 2.1 | yes | yes |
| natural | dirs | 65536 | build | ordered | btree-sets | 8 | 326.57 ms | 471.11 ms | 1.44× [1.43, 1.45] | +30.7% | [+30.3%, +31.1%] | 0.5 pts | 1.5 | yes | yes |
| natural | dirs | 65536 | build | ordered | ordered-mk1 | 8 | 326.88 ms | 312.88 ms | 0.96× [0.95, 0.97] | -4.0% | [-4.8%, -3.2%] | 0.9 pts | 2.7 | yes | yes |
| natural | dirs | 65536 | build | ordered | ordered-mk2 | 8 | 327.41 ms | 364.12 ms | 1.11× [1.10, 1.13] | +10.3% | [+9.1%, +11.4%] | 1.0 pts | 2.3 | yes | yes |
| natural | dirs | 65536 | build | ordered | ordered-mk4 | 8 | 327.84 ms | 356.86 ms | 1.10× [1.08, 1.11] | +8.8% | [+7.6%, +10.0%] | 1.1 pts | 2.2 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 91.1 | 73.9 | 0.81× [0.81, 0.81] | -23.3% | [-23.7%, -22.8%] | 0.3 pts | 0.6 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 92.0 | 179 | 1.93× [1.92, 1.94] | +48.2% | [+47.9%, +48.6%] | 0.4 pts | 0.7 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 91.4 | 74.0 | 0.81× [0.80, 0.82] | -23.6% | [-24.7%, -22.4%] | 0.9 pts | 0.6 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 92.2 | 81.8 | 0.89× [0.89, 0.90] | -12.2% | [-12.8%, -11.6%] | 1.0 pts | 0.7 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 92.8 | 86.3 | 0.94× [0.90, 0.98] | -6.5% | [-11.5%, -1.5%] | 2.9 pts | 1.9 | no | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 1332 | 2971 | 2.24× [2.20, 2.29] | +55.4% | [+54.5%, +56.4%] | 0.8 pts | 0.8 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 1340 | 5353 | 4.02× [4.00, 4.04] | +75.1% | [+75.0%, +75.2%] | 0.2 pts | 0.6 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 1352 | 4197 | 3.13× [3.04, 3.23] | +68.1% | [+67.2%, +69.0%] | 0.8 pts | 1.2 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 1354 | 3330 | 2.48× [2.44, 2.53] | +59.8% | [+59.1%, +60.4%] | 0.4 pts | 0.5 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 1333 | 2619 | 1.97× [1.94, 2.00] | +49.2% | [+48.4%, +50.0%] | 1.2 pts | 1.1 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 179 | 117 | 0.66× [0.65, 0.66] | -52.0% | [-52.7%, -51.3%] | 1.1 pts | 2.4 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 178 | 238 | 1.33× [1.33, 1.34] | +25.0% | [+24.7%, +25.3%] | 0.4 pts | 0.7 | yes | yes |
| natural | street | 4096 | churn | ordered | ordered-mk1 | 8 | 178 | 137 | 0.77× [0.77, 0.78] | -29.1% | [-29.4%, -28.8%] | 1.0 pts | 1.8 | yes | yes |
| natural | street | 4096 | churn | ordered | ordered-mk2 | 8 | 179 | 178 | 1.00× [0.99, 1.02] | +0.3% | [-1.2%, +1.8%] | 1.1 pts | 2.1 | yes | no |
| natural | street | 4096 | churn | ordered | ordered-mk4 | 8 | 178 | 182 | 1.02× [1.02, 1.03] | +2.2% | [+1.8%, +2.7%] | 0.5 pts | 1.5 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 7.00 ms | 4.40 ms | 0.63× [0.63, 0.63] | -59.4% | [-59.8%, -59.0%] | 1.3 pts | 2.4 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 6.98 ms | 9.39 ms | 1.33× [1.32, 1.35] | +25.0% | [+24.2%, +25.9%] | 1.0 pts | 3.2 | yes | yes |
| natural | street | 4096 | build | ordered | ordered-mk1 | 8 | 6.99 ms | 5.07 ms | 0.73× [0.72, 0.73] | -37.8% | [-38.9%, -36.7%] | 0.8 pts | 1.7 | yes | yes |
| natural | street | 4096 | build | ordered | ordered-mk2 | 8 | 6.92 ms | 7.06 ms | 1.01× [1.01, 1.02] | +1.5% | [+0.6%, +2.3%] | 0.5 pts | 1.3 | yes | yes |
| natural | street | 4096 | build | ordered | ordered-mk4 | 8 | 6.93 ms | 7.34 ms | 1.06× [1.05, 1.06] | +5.3% | [+4.7%, +6.0%] | 0.5 pts | 1.1 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | baseline | 8 | 113 | 95.1 | 0.85× [0.84, 0.85] | -18.2% | [-19.3%, -17.0%] | 0.8 pts | 1.7 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | btree-sets | 8 | 114 | 232 | 2.06× [2.01, 2.11] | +51.3% | [+50.2%, +52.5%] | 1.0 pts | 2.3 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 113 | 95.0 | 0.84× [0.83, 0.85] | -19.0% | [-20.2%, -17.8%] | 0.9 pts | 1.4 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 113 | 104 | 0.91× [0.90, 0.93] | -9.5% | [-11.3%, -7.6%] | 1.3 pts | 2.2 | yes | yes |
| natural | street | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 113 | 107 | 0.94× [0.94, 0.95] | -6.1% | [-6.5%, -5.7%] | 0.5 pts | 1.1 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | baseline | 8 | 1646 | 3273 | 1.99× [1.98, 2.00] | +49.8% | [+49.4%, +50.1%] | 0.3 pts | 0.4 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 1624 | 5643 | 3.46× [3.43, 3.49] | +71.1% | [+70.9%, +71.4%] | 0.2 pts | 0.9 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 1632 | 4487 | 2.74× [2.72, 2.77] | +63.5% | [+63.2%, +63.9%] | 0.2 pts | 0.7 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 1634 | 3614 | 2.21× [2.21, 2.22] | +54.8% | [+54.7%, +55.0%] | 0.4 pts | 1.0 | yes | yes |
| natural | street | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 1625 | 2810 | 1.72× [1.71, 1.74] | +42.0% | [+41.4%, +42.7%] | 0.5 pts | 0.8 | yes | yes |
| natural | street | 16384 | churn | ordered | baseline | 8 | 208 | 151 | 0.74× [0.71, 0.79] | -34.3% | [-41.7%, -27.0%] | 7.1 pts | 9.8 | no | yes |
| natural | street | 16384 | churn | ordered | btree-sets | 8 | 216 | 305 | 1.43× [1.41, 1.44] | +30.0% | [+29.3%, +30.7%] | 1.0 pts | 1.9 | yes | yes |
| natural | street | 16384 | churn | ordered | ordered-mk1 | 8 | 207 | 173 | 0.84× [0.81, 0.87] | -18.6% | [-22.8%, -14.4%] | 2.6 pts | 5.0 | no | yes |
| natural | street | 16384 | churn | ordered | ordered-mk2 | 8 | 207 | 215 | 1.04× [1.03, 1.04] | +3.7% | [+3.2%, +4.2%] | 0.4 pts | 1.1 | yes | yes |
| natural | street | 16384 | churn | ordered | ordered-mk4 | 8 | 208 | 222 | 1.07× [1.04, 1.09] | +6.4% | [+4.1%, +8.6%] | 1.8 pts | 4.9 | no | yes |
| natural | street | 16384 | build | ordered | baseline | 8 | 32.69 ms | 22.74 ms | 0.69× [0.69, 0.70] | -43.9% | [-44.6%, -43.1%] | 0.5 pts | 1.8 | yes | yes |
| natural | street | 16384 | build | ordered | btree-sets | 8 | 32.76 ms | 46.75 ms | 1.43× [1.42, 1.44] | +29.9% | [+29.4%, +30.4%] | 0.5 pts | 4.0 | yes | yes |
| natural | street | 16384 | build | ordered | ordered-mk1 | 8 | 32.77 ms | 25.70 ms | 0.78× [0.78, 0.79] | -27.5% | [-27.8%, -27.2%] | 0.3 pts | 1.0 | yes | yes |
| natural | street | 16384 | build | ordered | ordered-mk2 | 8 | 32.73 ms | 34.08 ms | 1.04× [1.04, 1.04] | +3.9% | [+3.6%, +4.1%] | 0.3 pts | 1.2 | yes | yes |
| natural | street | 16384 | build | ordered | ordered-mk4 | 8 | 32.72 ms | 35.05 ms | 1.07× [1.07, 1.08] | +6.8% | [+6.3%, +7.2%] | 0.4 pts | 1.3 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 145 | 146 | 1.00× [0.98, 1.02] | +0.2% | [-2.0%, +2.4%] | 1.7 pts | 1.5 | no | no |
| natural | street | 65536 | valuesFor | ordered | btree-sets | 8 | 161 | 407 | 2.55× [2.45, 2.66] | +60.9% | [+59.2%, +62.5%] | 1.2 pts | 1.2 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 144 | 140 | 0.96× [0.96, 0.97] | -4.0% | [-4.6%, -3.4%] | 1.7 pts | 1.2 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 144 | 144 | 1.01× [1.00, 1.01] | +0.6% | [-0.1%, +1.3%] | 0.8 pts | 0.7 | yes | no |
| natural | street | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 143 | 144 | 1.02× [0.99, 1.04] | +1.8% | [-0.7%, +4.3%] | 3.1 pts | 3.4 | no | no |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 1844 | 4069 | 2.21× [2.19, 2.23] | +54.8% | [+54.4%, +55.2%] | 0.4 pts | 0.6 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 1945 | 9315 | 4.83× [4.72, 4.95] | +79.3% | [+78.8%, +79.8%] | 0.4 pts | 1.4 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1822 | 5025 | 2.77× [2.74, 2.80] | +63.9% | [+63.4%, +64.3%] | 0.4 pts | 0.9 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1817 | 4028 | 2.23× [2.21, 2.24] | +55.1% | [+54.7%, +55.4%] | 0.4 pts | 0.8 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1791 | 3128 | 1.75× [1.73, 1.77] | +42.8% | [+42.2%, +43.4%] | 0.6 pts | 1.0 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 286 | 254 | 0.89× [0.86, 0.92] | -12.6% | [-16.7%, -8.6%] | 4.6 pts | 1.2 | no | yes |
| natural | street | 65536 | churn | ordered | btree-sets | 8 | 343 | 520 | 1.51× [1.50, 1.52] | +33.7% | [+33.3%, +34.2%] | 0.8 pts | 2.7 | yes | yes |
| natural | street | 65536 | churn | ordered | ordered-mk1 | 8 | 297 | 302 | 0.99× [0.94, 1.04] | -1.3% | [-6.4%, +3.8%] | 3.9 pts | 0.7 | no | no |
| natural | street | 65536 | churn | ordered | ordered-mk2 | 8 | 300 | 327 | 1.06× [1.06, 1.07] | +6.1% | [+5.6%, +6.5%] | 2.2 pts | 0.5 | yes | yes |
| natural | street | 65536 | churn | ordered | ordered-mk4 | 8 | 302 | 324 | 1.03× [1.01, 1.06] | +3.3% | [+0.8%, +5.8%] | 2.4 pts | 0.5 | no | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 181.93 ms | 158.27 ms | 0.87× [0.86, 0.88] | -15.0% | [-16.3%, -13.6%] | 1.5 pts | 4.5 | yes | yes |
| natural | street | 65536 | build | ordered | btree-sets | 8 | 182.20 ms | 286.39 ms | 1.57× [1.57, 1.58] | +36.3% | [+36.2%, +36.5%] | 0.6 pts | 2.6 | yes | yes |
| natural | street | 65536 | build | ordered | ordered-mk1 | 8 | 182.17 ms | 171.01 ms | 0.94× [0.93, 0.95] | -6.1% | [-7.4%, -4.8%] | 1.3 pts | 6.6 | yes | yes |
| natural | street | 65536 | build | ordered | ordered-mk2 | 8 | 182.12 ms | 207.25 ms | 1.14× [1.13, 1.15] | +12.2% | [+11.2%, +13.1%] | 0.8 pts | 4.1 | yes | yes |
| natural | street | 65536 | build | ordered | ordered-mk4 | 8 | 180.75 ms | 206.40 ms | 1.14× [1.12, 1.16] | +12.2% | [+11.0%, +13.5%] | 1.0 pts | 4.9 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 60.8 | 49.4 | 0.81× [0.81, 0.81] | -23.5% | [-24.0%, -23.0%] | 0.7 pts | 1.0 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 61.2 | 193 | 3.16× [3.14, 3.18] | +68.3% | [+68.1%, +68.5%] | 0.2 pts | 1.2 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 60.9 | 50.8 | 0.84× [0.83, 0.84] | -19.6% | [-20.2%, -19.1%] | 0.5 pts | 0.7 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 60.7 | 51.4 | 0.85× [0.84, 0.85] | -18.0% | [-18.7%, -17.4%] | 0.7 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 60.8 | 51.5 | 0.85× [0.84, 0.85] | -18.0% | [-18.8%, -17.1%] | 0.6 pts | 0.9 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2964 | 3719 | 1.25× [1.24, 1.26] | +20.1% | [+19.4%, +20.9%] | 0.5 pts | 0.7 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 2976 | 8267 | 2.79× [2.78, 2.80] | +64.2% | [+64.0%, +64.3%] | 0.2 pts | 1.5 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 2954 | 3583 | 1.20× [1.17, 1.24] | +16.7% | [+14.3%, +19.1%] | 1.5 pts | 2.0 | no | yes |
| natural | u64 | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 2948 | 3407 | 1.15× [1.15, 1.16] | +13.4% | [+12.8%, +14.0%] | 0.5 pts | 0.8 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 2953 | 3389 | 1.15× [1.15, 1.16] | +13.1% | [+12.8%, +13.5%] | 0.3 pts | 0.5 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 82.4 | 57.3 | 0.69× [0.69, 0.70] | -44.5% | [-45.2%, -43.9%] | 1.0 pts | 0.7 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 8 | 82.8 | 187 | 2.25× [2.25, 2.26] | +55.6% | [+55.5%, +55.7%] | 0.4 pts | 1.8 | yes | yes |
| natural | u64 | 4096 | churn | ordered | ordered-mk1 | 8 | 82.8 | 74.5 | 0.90× [0.90, 0.91] | -11.0% | [-11.6%, -10.4%] | 0.7 pts | 2.0 | yes | yes |
| natural | u64 | 4096 | churn | ordered | ordered-mk2 | 8 | 83.0 | 78.8 | 0.95× [0.94, 0.95] | -5.4% | [-5.8%, -5.0%] | 0.3 pts | 1.0 | yes | yes |
| natural | u64 | 4096 | churn | ordered | ordered-mk4 | 8 | 82.6 | 77.6 | 0.95× [0.94, 0.95] | -5.7% | [-6.4%, -5.1%] | 0.4 pts | 2.2 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 10.56 ms | 5.94 ms | 0.56× [0.56, 0.57] | -77.8% | [-78.9%, -76.7%] | 1.0 pts | 1.9 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 8 | 10.59 ms | 18.31 ms | 1.73× [1.71, 1.75] | +42.2% | [+41.6%, +42.7%] | 0.5 pts | 2.8 | yes | yes |
| natural | u64 | 4096 | build | ordered | ordered-mk1 | 8 | 10.55 ms | 7.43 ms | 0.70× [0.70, 0.71] | -42.0% | [-42.7%, -41.4%] | 0.6 pts | 1.1 | yes | yes |
| natural | u64 | 4096 | build | ordered | ordered-mk2 | 8 | 10.54 ms | 7.94 ms | 0.75× [0.75, 0.75] | -33.0% | [-33.3%, -32.7%] | 0.5 pts | 1.9 | yes | yes |
| natural | u64 | 4096 | build | ordered | ordered-mk4 | 8 | 10.57 ms | 8.03 ms | 0.76× [0.76, 0.76] | -31.8% | [-32.0%, -31.6%] | 0.4 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | baseline | 8 | 60.2 | 53.9 | 0.90× [0.87, 0.93] | -11.6% | [-15.5%, -7.7%] | 2.4 pts | 3.3 | no | yes |
| natural | u64 | 16384 | valuesFor | ordered | btree-sets | 8 | 62.7 | 243 | 3.82× [3.63, 4.04] | +73.8% | [+72.4%, +75.2%] | 1.3 pts | 3.6 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 60.0 | 57.1 | 0.95× [0.93, 0.96] | -5.8% | [-7.0%, -4.6%] | 1.5 pts | 1.7 | yes | yes |
| natural | u64 | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 60.6 | 59.5 | 0.98× [0.96, 1.00] | -1.9% | [-4.2%, +0.4%] | 1.9 pts | 2.0 | no | no |
| natural | u64 | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 61.5 | 60.7 | 0.99× [0.97, 1.01] | -1.1% | [-3.3%, +1.1%] | 1.7 pts | 2.4 | no | no |
| natural | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 3939 | 4491 | 1.14× [1.12, 1.15] | +12.0% | [+11.1%, +12.8%] | 0.7 pts | 1.2 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | btree-sets | 8 | 4015 | 8556 | 2.20× [2.12, 2.28] | +54.5% | [+52.8%, +56.1%] | 1.4 pts | 2.7 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 3951 | 4540 | 1.15× [1.14, 1.16] | +13.2% | [+12.4%, +14.1%] | 0.6 pts | 0.9 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 3912 | 3959 | 1.02× [1.00, 1.03] | +1.7% | [+0.4%, +2.9%] | 0.8 pts | 1.3 | yes | yes |
| natural | u64 | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 3892 | 3880 | 0.99× [0.99, 1.00] | -0.6% | [-1.5%, +0.2%] | 0.5 pts | 1.1 | yes | no |
| natural | u64 | 16384 | churn | ordered | baseline | 8 | 106 | 76.1 | 0.71× [0.69, 0.74] | -40.1% | [-45.6%, -34.6%] | 7.1 pts | 2.7 | no | yes |
| natural | u64 | 16384 | churn | ordered | btree-sets | 8 | 126 | 295 | 2.32× [2.11, 2.57] | +56.9% | [+52.7%, +61.1%] | 3.9 pts | 2.8 | yes | yes |
| natural | u64 | 16384 | churn | ordered | ordered-mk1 | 8 | 117 | 113 | 0.95× [0.94, 0.96] | -5.3% | [-6.8%, -3.9%] | 2.2 pts | 2.1 | yes | yes |
| natural | u64 | 16384 | churn | ordered | ordered-mk2 | 8 | 112 | 113 | 1.01× [1.00, 1.02] | +1.0% | [+0.2%, +1.9%] | 0.5 pts | 1.0 | yes | yes |
| natural | u64 | 16384 | churn | ordered | ordered-mk4 | 8 | 105 | 106 | 1.01× [0.99, 1.02] | +0.6% | [-0.9%, +2.2%] | 1.6 pts | 1.9 | yes | no |
| natural | u64 | 16384 | build | ordered | baseline | 8 | 38.36 ms | 26.52 ms | 0.69× [0.68, 0.70] | -45.0% | [-47.7%, -42.3%] | 2.0 pts | 2.3 | yes | yes |
| natural | u64 | 16384 | build | ordered | btree-sets | 8 | 36.79 ms | 93.41 ms | 2.53× [2.51, 2.56] | +60.5% | [+60.1%, +60.9%] | 0.3 pts | 1.7 | yes | yes |
| natural | u64 | 16384 | build | ordered | ordered-mk1 | 8 | 37.39 ms | 34.21 ms | 0.89× [0.87, 0.91] | -12.7% | [-15.1%, -10.3%] | 1.4 pts | 1.6 | no | yes |
| natural | u64 | 16384 | build | ordered | ordered-mk2 | 8 | 36.74 ms | 36.42 ms | 0.99× [0.98, 0.99] | -1.4% | [-1.7%, -1.1%] | 0.6 pts | 1.9 | yes | yes |
| natural | u64 | 16384 | build | ordered | ordered-mk4 | 8 | 36.93 ms | 36.15 ms | 0.98× [0.97, 0.98] | -2.4% | [-2.8%, -2.1%] | 0.4 pts | 0.9 | yes | yes |
| natural | u64 | 65536 | valuesFor | ordered | baseline | 8 | 92.9 | 79.2 | 0.84× [0.82, 0.85] | -19.6% | [-22.1%, -17.1%] | 4.9 pts | 1.3 | no | yes |
| natural | u64 | 65536 | valuesFor | ordered | btree-sets | 8 | 109 | 389 | 3.60× [3.46, 3.75] | +72.2% | [+71.1%, +73.4%] | 0.9 pts | 1.2 | yes | yes |
| natural | u64 | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 92.7 | 86.2 | 0.94× [0.89, 0.99] | -6.5% | [-12.1%, -0.9%] | 5.9 pts | 1.4 | no | yes |
| natural | u64 | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 89.8 | 86.1 | 0.99× [0.93, 1.06] | -1.0% | [-7.6%, +5.5%] | 4.5 pts | 1.4 | no | no |
| natural | u64 | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 92.5 | 90.3 | 1.00× [0.95, 1.05] | +0.0% | [-4.9%, +4.9%] | 3.5 pts | 1.0 | no | no |
| natural | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 3445 | 5403 | 1.57× [1.55, 1.58] | +36.2% | [+35.5%, +36.8%] | 0.5 pts | 0.8 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | btree-sets | 8 | 3576 | 12.8 µs | 3.57× [3.52, 3.62] | +72.0% | [+71.6%, +72.4%] | 0.3 pts | 0.9 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 3455 | 5311 | 1.53× [1.53, 1.54] | +34.8% | [+34.6%, +35.0%] | 0.7 pts | 0.8 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 3399 | 4157 | 1.23× [1.23, 1.23] | +18.8% | [+18.5%, +19.0%] | 0.6 pts | 0.6 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 3394 | 3434 | 1.01× [1.00, 1.02] | +1.0% | [-0.1%, +2.1%] | 1.1 pts | 1.1 | yes | no |
| natural | u64 | 65536 | churn | ordered | baseline | 8 | 214 | 157 | 0.71× [0.66, 0.78] | -39.9% | [-51.7%, -28.1%] | 11.1 pts | 1.2 | no | yes |
| natural | u64 | 65536 | churn | ordered | btree-sets | 8 | 253 | 472 | 1.89× [1.81, 1.97] | +47.0% | [+44.9%, +49.2%] | 1.9 pts | 7.3 | yes | yes |
| natural | u64 | 65536 | churn | ordered | ordered-mk1 | 8 | 181 | 181 | 0.95× [0.86, 1.05] | -5.7% | [-16.4%, +4.9%] | 8.1 pts | 1.3 | no | no |
| natural | u64 | 65536 | churn | ordered | ordered-mk2 | 8 | 182 | 187 | 0.99× [0.93, 1.05] | -1.5% | [-7.3%, +4.3%] | 4.7 pts | 0.8 | no | no |
| natural | u64 | 65536 | churn | ordered | ordered-mk4 | 8 | 193 | 197 | 0.93× [0.92, 0.95] | -7.5% | [-9.3%, -5.7%] | 2.2 pts | 0.3 | yes | yes |
| natural | u64 | 65536 | build | ordered | baseline | 8 | 256.23 ms | 194.69 ms | 0.76× [0.75, 0.77] | -31.5% | [-33.8%, -29.1%] | 1.6 pts | 2.1 | yes | yes |
| natural | u64 | 65536 | build | ordered | btree-sets | 8 | 256.83 ms | 561.77 ms | 2.19× [2.15, 2.22] | +54.3% | [+53.6%, +55.0%] | 0.5 pts | 2.5 | yes | yes |
| natural | u64 | 65536 | build | ordered | ordered-mk1 | 8 | 254.63 ms | 245.03 ms | 0.96× [0.95, 0.97] | -4.3% | [-5.2%, -3.4%] | 0.7 pts | 1.2 | yes | yes |
| natural | u64 | 65536 | build | ordered | ordered-mk2 | 8 | 254.24 ms | 262.24 ms | 1.03× [1.02, 1.03] | +2.8% | [+2.4%, +3.3%] | 0.7 pts | 1.6 | yes | yes |
| natural | u64 | 65536 | build | ordered | ordered-mk4 | 8 | 253.88 ms | 259.25 ms | 1.02× [1.01, 1.02] | +1.7% | [+1.0%, +2.3%] | 0.5 pts | 0.8 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | baseline | 4 | 144 | 125 | 0.87× [0.86, 0.89] | -14.3% | [-16.1%, -12.6%] | 1.1 pts | 1.5 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | btree-sets | 4 | 142 | 230 | 1.62× [1.60, 1.64] | +38.3% | [+37.6%, +39.1%] | 0.5 pts | 0.8 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | ordered-mk1 | 4 | 141 | 123 | 0.87× [0.85, 0.88] | -15.4% | [-17.2%, -13.6%] | 1.1 pts | 1.8 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | ordered-mk2 | 4 | 141 | 132 | 0.94× [0.93, 0.94] | -6.9% | [-7.2%, -6.6%] | 0.2 pts | 0.3 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | ordered-mk4 | 4 | 140 | 138 | 0.98× [0.97, 0.98] | -2.1% | [-2.6%, -1.6%] | 0.3 pts | 0.5 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | baseline | 4 | 3282 | 5065 | 1.55× [1.54, 1.56] | +35.4% | [+34.9%, +35.8%] | 0.3 pts | 0.3 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | btree-sets | 4 | 3312 | 8400 | 2.57× [2.47, 2.68] | +61.1% | [+59.5%, +62.7%] | 1.0 pts | 2.7 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | ordered-mk1 | 4 | 3290 | 5554 | 1.68× [1.67, 1.70] | +40.6% | [+40.0%, +41.2%] | 0.4 pts | 0.5 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | ordered-mk2 | 4 | 3277 | 4578 | 1.40× [1.39, 1.41] | +28.5% | [+27.8%, +29.3%] | 0.5 pts | 0.6 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | ordered-mk4 | 4 | 3332 | 3773 | 1.14× [1.13, 1.15] | +12.2% | [+11.2%, +13.3%] | 0.7 pts | 0.6 | yes | yes |
| natural | url | 4096 | churn | ordered | baseline | 4 | 213 | 155 | 0.73× [0.72, 0.73] | -37.6% | [-38.9%, -36.4%] | 0.8 pts | 1.3 | yes | yes |
| natural | url | 4096 | churn | ordered | btree-sets | 4 | 216 | 241 | 1.11× [1.09, 1.13] | +10.2% | [+8.6%, +11.9%] | 1.0 pts | 1.8 | yes | yes |
| natural | url | 4096 | churn | ordered | ordered-mk1 | 4 | 212 | 170 | 0.80× [0.80, 0.81] | -24.3% | [-25.0%, -23.7%] | 0.4 pts | 1.0 | yes | yes |
| natural | url | 4096 | churn | ordered | ordered-mk2 | 4 | 214 | 217 | 1.00× [0.99, 1.02] | +0.3% | [-1.2%, +1.8%] | 0.9 pts | 2.5 | yes | no |
| natural | url | 4096 | churn | ordered | ordered-mk4 | 4 | 209 | 213 | 1.02× [1.02, 1.02] | +1.8% | [+1.8%, +1.9%] | 0.0 pts | 0.2 | yes | yes |
| natural | url | 4096 | build | ordered | baseline | 4 | 21.40 ms | 15.00 ms | 0.70× [0.70, 0.71] | -42.6% | [-43.4%, -41.7%] | 0.5 pts | 0.9 | yes | yes |
| natural | url | 4096 | build | ordered | btree-sets | 4 | 21.42 ms | 22.62 ms | 1.05× [1.05, 1.06] | +5.2% | [+4.4%, +5.9%] | 0.5 pts | 1.3 | yes | yes |
| natural | url | 4096 | build | ordered | ordered-mk1 | 4 | 21.36 ms | 16.38 ms | 0.77× [0.76, 0.77] | -30.4% | [-31.0%, -29.8%] | 0.4 pts | 1.4 | yes | yes |
| natural | url | 4096 | build | ordered | ordered-mk2 | 4 | 21.34 ms | 20.86 ms | 0.98× [0.97, 0.98] | -2.2% | [-2.9%, -1.5%] | 0.4 pts | 1.8 | yes | yes |
| natural | url | 4096 | build | ordered | ordered-mk4 | 4 | 21.29 ms | 21.77 ms | 1.02× [1.01, 1.03] | +2.1% | [+1.4%, +2.7%] | 0.4 pts | 1.8 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | baseline | 8 | 168 | 152 | 0.91× [0.87, 0.96] | -9.5% | [-14.4%, -4.6%] | 3.0 pts | 4.7 | no | yes |
| natural | url | 16384 | valuesFor | ordered | btree-sets | 8 | 174 | 296 | 1.70× [1.66, 1.73] | +41.0% | [+39.7%, +42.3%] | 0.9 pts | 1.1 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 167 | 149 | 0.90× [0.89, 0.90] | -11.7% | [-12.7%, -10.8%] | 0.6 pts | 0.8 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 168 | 160 | 0.96× [0.94, 0.98] | -4.2% | [-6.0%, -2.5%] | 1.3 pts | 1.8 | yes | yes |
| natural | url | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 167 | 163 | 0.98× [0.96, 1.01] | -1.7% | [-4.3%, +0.9%] | 1.6 pts | 3.0 | no | no |
| natural | url | 16384 | valuesBetween | ordered | baseline | 8 | 3513 | 5361 | 1.53× [1.51, 1.56] | +34.8% | [+33.8%, +35.8%] | 0.7 pts | 1.3 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | btree-sets | 8 | 3576 | 8924 | 2.52× [2.49, 2.55] | +60.3% | [+59.8%, +60.8%] | 0.4 pts | 0.8 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 3547 | 5882 | 1.67× [1.65, 1.68] | +40.0% | [+39.5%, +40.5%] | 0.4 pts | 0.8 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 3511 | 4867 | 1.38× [1.37, 1.39] | +27.7% | [+27.1%, +28.2%] | 0.4 pts | 0.6 | yes | yes |
| natural | url | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 3510 | 3968 | 1.13× [1.13, 1.14] | +11.8% | [+11.2%, +12.4%] | 0.4 pts | 0.7 | yes | yes |
| natural | url | 16384 | churn | ordered | baseline | 8 | 280 | 234 | 0.85× [0.82, 0.89] | -17.6% | [-22.7%, -12.5%] | 3.9 pts | 4.4 | no | yes |
| natural | url | 16384 | churn | ordered | btree-sets | 8 | 296 | 363 | 1.23× [1.21, 1.24] | +18.4% | [+17.5%, +19.2%] | 1.2 pts | 1.7 | yes | yes |
| natural | url | 16384 | churn | ordered | ordered-mk1 | 8 | 273 | 243 | 0.89× [0.88, 0.90] | -12.8% | [-14.0%, -11.6%] | 1.8 pts | 1.5 | yes | yes |
| natural | url | 16384 | churn | ordered | ordered-mk2 | 8 | 281 | 290 | 1.03× [1.01, 1.05] | +2.9% | [+0.7%, +5.1%] | 1.4 pts | 0.8 | no | yes |
| natural | url | 16384 | churn | ordered | ordered-mk4 | 8 | 276 | 285 | 1.01× [1.00, 1.03] | +1.3% | [-0.5%, +3.1%] | 1.4 pts | 0.6 | yes | no |
| natural | url | 16384 | build | ordered | baseline | 8 | 96.30 ms | 72.88 ms | 0.76× [0.76, 0.76] | -31.9% | [-32.4%, -31.3%] | 0.5 pts | 2.9 | yes | yes |
| natural | url | 16384 | build | ordered | btree-sets | 8 | 96.26 ms | 113.41 ms | 1.18× [1.17, 1.19] | +15.3% | [+14.4%, +16.2%] | 0.6 pts | 5.4 | yes | yes |
| natural | url | 16384 | build | ordered | ordered-mk1 | 8 | 96.41 ms | 77.73 ms | 0.81× [0.80, 0.83] | -22.8% | [-25.5%, -20.0%] | 2.2 pts | 12.1 | no | yes |
| natural | url | 16384 | build | ordered | ordered-mk2 | 8 | 96.45 ms | 97.37 ms | 1.01× [1.00, 1.02] | +1.0% | [+0.1%, +2.0%] | 0.6 pts | 3.7 | yes | yes |
| natural | url | 16384 | build | ordered | ordered-mk4 | 8 | 96.11 ms | 98.11 ms | 1.02× [1.02, 1.03] | +2.0% | [+1.6%, +2.4%] | 0.4 pts | 2.5 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | baseline | 8 | 257 | 281 | 1.11× [1.07, 1.15] | +9.8% | [+6.6%, +12.9%] | 2.1 pts | 0.8 | no | yes |
| natural | url | 65536 | valuesFor | ordered | btree-sets | 8 | 288 | 566 | 2.02× [1.94, 2.11] | +50.5% | [+48.3%, +52.7%] | 2.0 pts | 1.4 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 253 | 273 | 1.10× [1.08, 1.11] | +9.0% | [+7.7%, +10.3%] | 1.9 pts | 0.6 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 251 | 270 | 1.09× [1.07, 1.11] | +8.1% | [+6.5%, +9.7%] | 3.2 pts | 1.1 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 250 | 262 | 1.05× [1.04, 1.06] | +4.7% | [+3.5%, +6.0%] | 2.5 pts | 1.0 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | baseline | 8 | 4266 | 7375 | 1.73× [1.71, 1.74] | +42.1% | [+41.7%, +42.4%] | 0.7 pts | 1.0 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | btree-sets | 8 | 4585 | 14.1 µs | 3.11× [3.04, 3.18] | +67.8% | [+67.1%, +68.6%] | 0.5 pts | 1.2 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 4259 | 7136 | 1.68× [1.66, 1.70] | +40.5% | [+39.7%, +41.3%] | 0.5 pts | 0.7 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 4192 | 5887 | 1.39× [1.37, 1.41] | +28.2% | [+27.3%, +29.1%] | 1.0 pts | 1.2 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 4173 | 4768 | 1.14× [1.12, 1.16] | +12.5% | [+11.0%, +14.1%] | 1.1 pts | 1.1 | yes | yes |
| natural | url | 65536 | churn | ordered | baseline | 8 | 489 | 453 | 0.89× [0.88, 0.91] | -11.8% | [-13.6%, -10.1%] | 1.4 pts | 0.2 | yes | yes |
| natural | url | 65536 | churn | ordered | btree-sets | 8 | 551 | 691 | 1.26× [1.23, 1.29] | +20.6% | [+18.5%, +22.7%] | 1.6 pts | 0.7 | no | yes |
| natural | url | 65536 | churn | ordered | ordered-mk1 | 8 | 486 | 450 | 0.88× [0.87, 0.89] | -13.6% | [-15.0%, -12.2%] | 3.4 pts | 0.3 | yes | yes |
| natural | url | 65536 | churn | ordered | ordered-mk2 | 8 | 477 | 443 | 0.89× [0.85, 0.94] | -11.8% | [-17.2%, -6.4%] | 4.4 pts | 0.4 | no | yes |
| natural | url | 65536 | churn | ordered | ordered-mk4 | 8 | 476 | 424 | 0.88× [0.84, 0.93] | -13.9% | [-19.7%, -8.0%] | 3.9 pts | 0.4 | no | yes |
| natural | url | 65536 | build | ordered | baseline | 8 | 604.97 ms | 566.68 ms | 0.94× [0.92, 0.95] | -6.9% | [-8.3%, -5.5%] | 1.2 pts | 1.6 | yes | yes |
| natural | url | 65536 | build | ordered | btree-sets | 8 | 601.75 ms | 835.72 ms | 1.39× [1.37, 1.40] | +27.9% | [+27.0%, +28.7%] | 0.7 pts | 2.0 | yes | yes |
| natural | url | 65536 | build | ordered | ordered-mk1 | 8 | 606.07 ms | 572.32 ms | 0.94× [0.93, 0.95] | -6.2% | [-7.0%, -5.4%] | 0.7 pts | 1.1 | yes | yes |
| natural | url | 65536 | build | ordered | ordered-mk2 | 8 | 606.79 ms | 634.76 ms | 1.05× [1.04, 1.06] | +4.6% | [+3.6%, +5.6%] | 0.9 pts | 1.4 | yes | yes |
| natural | url | 65536 | build | ordered | ordered-mk4 | 8 | 604.52 ms | 625.86 ms | 1.04× [1.02, 1.05] | +3.6% | [+2.0%, +5.2%] | 1.3 pts | 2.2 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | baseline | 6 | 123 | 107 | 0.87× [0.87, 0.88] | -14.6% | [-15.4%, -13.8%] | 0.5 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | btree-map | 6 | 124 | 143 | 1.15× [1.14, 1.16] | +13.4% | [+12.6%, +14.1%] | 0.5 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | ordered-mk1 | 6 | 123 | 112 | 0.91× [0.89, 0.92] | -10.1% | [-12.0%, -8.2%] | 1.3 pts | 1.2 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | ordered-mk2 | 6 | 124 | 120 | 0.97× [0.96, 0.98] | -3.4% | [-4.4%, -2.3%] | 1.2 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesFor | ordered | ordered-mk4 | 6 | 124 | 122 | 0.98× [0.96, 1.00] | -2.1% | [-3.7%, -0.4%] | 1.0 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 1506 | 2615 | 1.76× [1.70, 1.82] | +43.1% | [+41.2%, +44.9%] | 1.2 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | btree-map | 6 | 1440 | 611 | 0.43× [0.42, 0.43] | -133.9% | [-135.8%, -132.0%] | 2.0 pts | 0.5 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | ordered-mk1 | 6 | 1493 | 4088 | 2.74× [2.71, 2.77] | +63.5% | [+63.1%, +63.9%] | 0.6 pts | 0.9 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | ordered-mk2 | 6 | 1498 | 3148 | 2.11× [2.05, 2.17] | +52.6% | [+51.3%, +54.0%] | 0.9 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | valuesBetween | ordered | ordered-mk4 | 6 | 1476 | 2378 | 1.62× [1.59, 1.66] | +38.4% | [+37.1%, +39.8%] | 1.1 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | baseline | 6 | 251 | 236 | 0.95× [0.94, 0.95] | -5.6% | [-6.0%, -5.1%] | 0.6 pts | 1.1 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | btree-map | 6 | 254 | 215 | 0.85× [0.85, 0.86] | -17.4% | [-18.0%, -16.7%] | 0.7 pts | 0.8 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | ordered-mk1 | 6 | 251 | 265 | 1.06× [1.05, 1.07] | +5.6% | [+5.1%, +6.2%] | 0.5 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | ordered-mk2 | 6 | 254 | 355 | 1.40× [1.39, 1.41] | +28.6% | [+28.1%, +29.1%] | 0.3 pts | 1.0 | yes | yes |
| single-value | dirs | 4096 | churn | ordered | ordered-mk4 | 6 | 253 | 326 | 1.29× [1.28, 1.29] | +22.3% | [+22.0%, +22.5%] | 0.1 pts | 0.7 | yes | yes |
| single-value | dirs | 4096 | build | ordered | baseline | 6 | 4.11 ms | 3.32 ms | 0.81× [0.80, 0.82] | -24.1% | [-25.5%, -22.6%] | 0.9 pts | 1.7 | yes | yes |
| single-value | dirs | 4096 | build | ordered | btree-map | 6 | 4.09 ms | 3.18 ms | 0.78× [0.77, 0.79] | -28.8% | [-30.6%, -27.0%] | 1.1 pts | 1.6 | yes | yes |
| single-value | dirs | 4096 | build | ordered | ordered-mk1 | 6 | 4.08 ms | 3.74 ms | 0.91× [0.90, 0.92] | -9.5% | [-10.6%, -8.4%] | 0.9 pts | 1.5 | yes | yes |
| single-value | dirs | 4096 | build | ordered | ordered-mk2 | 6 | 4.04 ms | 5.23 ms | 1.29× [1.27, 1.31] | +22.5% | [+21.3%, +23.6%] | 0.8 pts | 1.8 | yes | yes |
| single-value | dirs | 4096 | build | ordered | ordered-mk4 | 6 | 4.04 ms | 4.94 ms | 1.21× [1.20, 1.22] | +17.7% | [+17.0%, +18.4%] | 0.5 pts | 1.0 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | baseline | 6 | 151 | 138 | 0.92× [0.90, 0.93] | -9.1% | [-10.9%, -7.3%] | 1.2 pts | 3.0 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | btree-map | 6 | 154 | 200 | 1.29× [1.27, 1.31] | +22.4% | [+21.4%, +23.4%] | 1.1 pts | 1.7 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | ordered-mk1 | 6 | 153 | 146 | 0.96× [0.94, 0.97] | -4.7% | [-6.5%, -2.8%] | 1.2 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | valuesFor | ordered | ordered-mk2 | 6 | 153 | 152 | 1.00× [0.99, 1.01] | -0.3% | [-1.2%, +0.6%] | 0.8 pts | 1.0 | yes | no |
| single-value | dirs | 16384 | valuesFor | ordered | ordered-mk4 | 6 | 153 | 152 | 0.99× [0.98, 1.01] | -0.6% | [-1.8%, +0.5%] | 0.7 pts | 1.0 | yes | no |
| single-value | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 1547 | 2925 | 1.89× [1.87, 1.91] | +47.0% | [+46.5%, +47.6%] | 0.5 pts | 0.7 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | btree-map | 6 | 1524 | 750 | 0.50× [0.49, 0.50] | -101.7% | [-103.9%, -99.6%] | 3.4 pts | 1.2 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | ordered-mk1 | 6 | 1541 | 4375 | 2.85× [2.82, 2.88] | +64.9% | [+64.5%, +65.3%] | 0.3 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | ordered-mk2 | 6 | 1541 | 3458 | 2.26× [2.21, 2.31] | +55.8% | [+54.8%, +56.7%] | 0.8 pts | 1.3 | yes | yes |
| single-value | dirs | 16384 | valuesBetween | ordered | ordered-mk4 | 6 | 1525 | 2642 | 1.74× [1.72, 1.75] | +42.4% | [+41.8%, +42.9%] | 0.5 pts | 0.8 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | baseline | 6 | 292 | 303 | 1.04× [1.03, 1.04] | +3.5% | [+3.3%, +3.7%] | 0.2 pts | 0.6 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | btree-map | 6 | 291 | 277 | 0.96× [0.94, 0.97] | -4.5% | [-5.8%, -3.2%] | 0.8 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | ordered-mk1 | 6 | 293 | 329 | 1.12× [1.11, 1.14] | +11.0% | [+9.6%, +12.4%] | 0.9 pts | 2.6 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | ordered-mk2 | 6 | 301 | 448 | 1.49× [1.47, 1.51] | +32.8% | [+31.8%, +33.9%] | 0.8 pts | 1.4 | yes | yes |
| single-value | dirs | 16384 | churn | ordered | ordered-mk4 | 6 | 291 | 390 | 1.34× [1.34, 1.35] | +25.5% | [+25.2%, +25.7%] | 0.2 pts | 0.4 | yes | yes |
| single-value | dirs | 16384 | build | ordered | baseline | 6 | 18.03 ms | 15.95 ms | 0.88× [0.88, 0.89] | -13.0% | [-13.9%, -12.1%] | 0.9 pts | 2.4 | yes | yes |
| single-value | dirs | 16384 | build | ordered | btree-map | 6 | 18.02 ms | 15.73 ms | 0.87× [0.87, 0.88] | -14.3% | [-15.1%, -13.4%] | 0.7 pts | 1.8 | yes | yes |
| single-value | dirs | 16384 | build | ordered | ordered-mk1 | 6 | 18.05 ms | 17.38 ms | 0.97× [0.95, 0.98] | -3.6% | [-5.3%, -2.0%] | 1.0 pts | 2.9 | yes | yes |
| single-value | dirs | 16384 | build | ordered | ordered-mk2 | 6 | 17.93 ms | 23.86 ms | 1.33× [1.31, 1.35] | +24.9% | [+23.6%, +26.2%] | 0.8 pts | 2.5 | yes | yes |
| single-value | dirs | 16384 | build | ordered | ordered-mk4 | 6 | 17.94 ms | 22.47 ms | 1.25× [1.23, 1.27] | +20.0% | [+18.4%, +21.6%] | 0.9 pts | 3.7 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | baseline | 8 | 197 | 218 | 1.14× [1.09, 1.19] | +12.0% | [+8.3%, +15.8%] | 2.5 pts | 2.0 | no | yes |
| single-value | dirs | 65536 | valuesFor | ordered | btree-map | 8 | 196 | 270 | 1.39× [1.37, 1.41] | +28.0% | [+27.0%, +28.9%] | 1.1 pts | 0.9 | yes | yes |
| single-value | dirs | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 190 | 197 | 1.04× [1.02, 1.07] | +4.1% | [+2.1%, +6.1%] | 1.3 pts | 1.4 | no | yes |
| single-value | dirs | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 190 | 203 | 1.08× [1.01, 1.17] | +7.8% | [+0.8%, +14.8%] | 4.5 pts | 4.7 | no | yes |
| single-value | dirs | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 188 | 195 | 1.04× [1.03, 1.05] | +4.1% | [+3.1%, +5.1%] | 1.1 pts | 1.4 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 1678 | 3644 | 2.25× [2.21, 2.29] | +55.5% | [+54.8%, +56.3%] | 1.3 pts | 1.1 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | btree-map | 8 | 1593 | 1041 | 0.65× [0.62, 0.68] | -54.2% | [-61.1%, -47.3%] | 5.7 pts | 1.2 | no | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1606 | 4600 | 2.89× [2.82, 2.97] | +65.4% | [+64.5%, +66.3%] | 0.7 pts | 1.6 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1591 | 3677 | 2.32× [2.30, 2.33] | +56.8% | [+56.5%, +57.1%] | 0.5 pts | 1.1 | yes | yes |
| single-value | dirs | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1580 | 2856 | 1.82× [1.81, 1.83] | +45.1% | [+44.9%, +45.3%] | 1.0 pts | 1.6 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | baseline | 8 | 380 | 459 | 1.21× [1.18, 1.23] | +17.1% | [+15.3%, +19.0%] | 1.9 pts | 1.7 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | btree-map | 8 | 420 | 472 | 1.14× [1.11, 1.18] | +12.6% | [+9.6%, +15.5%] | 1.8 pts | 2.6 | no | yes |
| single-value | dirs | 65536 | churn | ordered | ordered-mk1 | 8 | 369 | 449 | 1.20× [1.18, 1.21] | +16.4% | [+15.3%, +17.6%] | 0.9 pts | 0.8 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | ordered-mk2 | 8 | 388 | 560 | 1.46× [1.43, 1.50] | +31.5% | [+29.9%, +33.1%] | 1.4 pts | 0.9 | yes | yes |
| single-value | dirs | 65536 | churn | ordered | ordered-mk4 | 8 | 365 | 499 | 1.34× [1.32, 1.36] | +25.4% | [+24.1%, +26.7%] | 1.0 pts | 0.8 | yes | yes |
| single-value | dirs | 65536 | build | ordered | baseline | 8 | 86.56 ms | 90.89 ms | 1.05× [1.04, 1.07] | +5.2% | [+3.9%, +6.4%] | 2.3 pts | 13.1 | yes | yes |
| single-value | dirs | 65536 | build | ordered | btree-map | 8 | 86.51 ms | 86.61 ms | 1.01× [0.99, 1.02] | +0.7% | [-0.7%, +2.1%] | 1.3 pts | 9.5 | yes | no |
| single-value | dirs | 65536 | build | ordered | ordered-mk1 | 8 | 86.40 ms | 91.19 ms | 1.06× [1.05, 1.07] | +5.3% | [+4.3%, +6.2%] | 0.7 pts | 5.8 | yes | yes |
| single-value | dirs | 65536 | build | ordered | ordered-mk2 | 8 | 86.41 ms | 118.57 ms | 1.37× [1.36, 1.39] | +27.1% | [+26.4%, +27.9%] | 0.7 pts | 6.2 | yes | yes |
| single-value | dirs | 65536 | build | ordered | ordered-mk4 | 8 | 86.29 ms | 112.56 ms | 1.30× [1.29, 1.32] | +23.3% | [+22.3%, +24.2%] | 0.7 pts | 6.8 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | baseline | 8 | 80.7 | 66.3 | 0.82× [0.82, 0.83] | -21.5% | [-22.2%, -20.8%] | 0.5 pts | 1.4 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | btree-map | 8 | 81.4 | 124 | 1.52× [1.51, 1.54] | +34.4% | [+33.8%, +35.0%] | 0.4 pts | 0.8 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 80.9 | 66.4 | 0.82× [0.82, 0.83] | -21.3% | [-21.6%, -20.9%] | 1.0 pts | 1.3 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 80.6 | 73.3 | 0.91× [0.90, 0.92] | -9.8% | [-10.9%, -8.8%] | 1.0 pts | 1.2 | yes | yes |
| single-value | street | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 81.5 | 76.7 | 0.94× [0.93, 0.95] | -6.3% | [-7.3%, -5.3%] | 1.2 pts | 1.2 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | baseline | 8 | 858 | 2318 | 2.69× [2.68, 2.70] | +62.8% | [+62.7%, +63.0%] | 0.6 pts | 0.9 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | btree-map | 8 | 852 | 593 | 0.70× [0.69, 0.71] | -42.1% | [-44.0%, -40.1%] | 2.1 pts | 1.0 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 872 | 3761 | 4.32× [4.25, 4.38] | +76.8% | [+76.5%, +77.2%] | 0.3 pts | 0.7 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 870 | 2892 | 3.36× [3.29, 3.43] | +70.2% | [+69.6%, +70.8%] | 0.6 pts | 1.0 | yes | yes |
| single-value | street | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 862 | 2200 | 2.58× [2.50, 2.66] | +61.2% | [+60.1%, +62.4%] | 0.7 pts | 0.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | baseline | 8 | 166 | 146 | 0.89× [0.87, 0.91] | -12.5% | [-14.9%, -10.0%] | 1.8 pts | 3.8 | no | yes |
| single-value | street | 4096 | churn | ordered | btree-map | 8 | 163 | 179 | 1.10× [1.10, 1.10] | +9.1% | [+8.8%, +9.3%] | 0.6 pts | 1.9 | yes | yes |
| single-value | street | 4096 | churn | ordered | ordered-mk1 | 8 | 165 | 164 | 0.99× [0.99, 1.00] | -0.7% | [-1.2%, -0.2%] | 0.5 pts | 0.7 | yes | yes |
| single-value | street | 4096 | churn | ordered | ordered-mk2 | 8 | 166 | 233 | 1.40× [1.39, 1.41] | +28.7% | [+28.2%, +29.2%] | 0.4 pts | 1.0 | yes | yes |
| single-value | street | 4096 | churn | ordered | ordered-mk4 | 8 | 166 | 226 | 1.37× [1.37, 1.37] | +26.9% | [+26.8%, +27.0%] | 0.5 pts | 1.5 | yes | yes |
| single-value | street | 4096 | build | ordered | baseline | 8 | 2.64 ms | 2.03 ms | 0.78× [0.77, 0.79] | -28.8% | [-30.6%, -27.0%] | 1.9 pts | 3.1 | yes | yes |
| single-value | street | 4096 | build | ordered | btree-map | 8 | 2.61 ms | 2.74 ms | 1.05× [1.03, 1.07] | +4.7% | [+2.8%, +6.6%] | 1.6 pts | 2.9 | yes | yes |
| single-value | street | 4096 | build | ordered | ordered-mk1 | 8 | 2.62 ms | 2.30 ms | 0.88× [0.87, 0.89] | -13.5% | [-14.5%, -12.4%] | 1.0 pts | 1.7 | yes | yes |
| single-value | street | 4096 | build | ordered | ordered-mk2 | 8 | 2.60 ms | 3.51 ms | 1.34× [1.33, 1.36] | +25.6% | [+24.7%, +26.6%] | 1.0 pts | 2.2 | yes | yes |
| single-value | street | 4096 | build | ordered | ordered-mk4 | 8 | 2.60 ms | 3.44 ms | 1.32× [1.31, 1.33] | +24.2% | [+23.4%, +25.1%] | 0.7 pts | 1.5 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | baseline | 8 | 96.2 | 86.7 | 0.90× [0.90, 0.91] | -10.9% | [-11.4%, -10.4%] | 0.7 pts | 3.1 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | btree-map | 8 | 97.1 | 170 | 1.73× [1.73, 1.74] | +42.3% | [+42.1%, +42.4%] | 0.4 pts | 0.9 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 97.8 | 86.5 | 0.88× [0.88, 0.89] | -13.1% | [-14.0%, -12.2%] | 1.0 pts | 1.2 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 97.7 | 94.5 | 0.97× [0.96, 0.98] | -3.5% | [-4.6%, -2.4%] | 1.0 pts | 1.4 | yes | yes |
| single-value | street | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 97.4 | 96.2 | 0.99× [0.98, 0.99] | -1.5% | [-2.1%, -0.8%] | 0.7 pts | 1.1 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | baseline | 8 | 1078 | 2493 | 2.32× [2.30, 2.34] | +56.9% | [+56.6%, +57.3%] | 0.4 pts | 0.7 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | btree-map | 8 | 1057 | 674 | 0.64× [0.63, 0.64] | -57.1% | [-58.3%, -55.9%] | 1.2 pts | 0.8 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 1081 | 3937 | 3.69× [3.67, 3.71] | +72.9% | [+72.7%, +73.0%] | 0.1 pts | 0.3 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 1076 | 3080 | 2.88× [2.85, 2.91] | +65.2% | [+64.9%, +65.6%] | 0.4 pts | 0.8 | yes | yes |
| single-value | street | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 1072 | 2366 | 2.20× [2.18, 2.21] | +54.4% | [+54.2%, +54.7%] | 0.3 pts | 0.5 | yes | yes |
| single-value | street | 16384 | churn | ordered | baseline | 8 | 193 | 187 | 0.97× [0.96, 0.97] | -3.6% | [-3.9%, -3.3%] | 1.8 pts | 3.7 | yes | yes |
| single-value | street | 16384 | churn | ordered | btree-map | 8 | 193 | 230 | 1.19× [1.18, 1.19] | +15.7% | [+15.5%, +15.9%] | 1.0 pts | 3.2 | yes | yes |
| single-value | street | 16384 | churn | ordered | ordered-mk1 | 8 | 194 | 204 | 1.06× [1.04, 1.08] | +6.0% | [+4.2%, +7.8%] | 1.3 pts | 2.6 | yes | yes |
| single-value | street | 16384 | churn | ordered | ordered-mk2 | 8 | 194 | 280 | 1.44× [1.43, 1.44] | +30.5% | [+30.3%, +30.8%] | 0.3 pts | 0.9 | yes | yes |
| single-value | street | 16384 | churn | ordered | ordered-mk4 | 8 | 196 | 268 | 1.37× [1.36, 1.38] | +27.2% | [+26.7%, +27.8%] | 0.5 pts | 1.6 | yes | yes |
| single-value | street | 16384 | build | ordered | baseline | 8 | 12.08 ms | 9.89 ms | 0.82× [0.81, 0.83] | -21.6% | [-22.7%, -20.5%] | 1.7 pts | 3.7 | yes | yes |
| single-value | street | 16384 | build | ordered | btree-map | 8 | 11.98 ms | 13.65 ms | 1.14× [1.13, 1.15] | +12.2% | [+11.5%, +12.9%] | 0.9 pts | 3.4 | yes | yes |
| single-value | street | 16384 | build | ordered | ordered-mk1 | 8 | 12.05 ms | 11.08 ms | 0.92× [0.91, 0.93] | -8.8% | [-9.4%, -8.1%] | 0.7 pts | 1.4 | yes | yes |
| single-value | street | 16384 | build | ordered | ordered-mk2 | 8 | 12.08 ms | 16.11 ms | 1.33× [1.31, 1.35] | +24.9% | [+23.9%, +25.9%] | 0.7 pts | 2.5 | yes | yes |
| single-value | street | 16384 | build | ordered | ordered-mk4 | 8 | 12.00 ms | 15.89 ms | 1.32× [1.30, 1.33] | +24.2% | [+23.3%, +25.0%] | 0.6 pts | 1.7 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | baseline | 8 | 119 | 122 | 1.02× [1.01, 1.03] | +2.1% | [+0.8%, +3.4%] | 1.1 pts | 2.3 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | btree-map | 8 | 120 | 220 | 1.83× [1.82, 1.84] | +45.3% | [+45.0%, +45.6%] | 0.5 pts | 1.1 | yes | yes |
| single-value | street | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 120 | 115 | 0.98× [0.96, 1.00] | -2.4% | [-4.5%, -0.3%] | 4.2 pts | 4.5 | no | yes |
| single-value | street | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 119 | 122 | 1.03× [1.00, 1.06] | +2.8% | [-0.1%, +5.6%] | 1.7 pts | 2.0 | no | no |
| single-value | street | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 119 | 122 | 1.03× [1.01, 1.05] | +3.1% | [+1.4%, +4.9%] | 1.1 pts | 1.8 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | baseline | 8 | 1088 | 2729 | 2.53× [2.50, 2.56] | +60.5% | [+60.1%, +61.0%] | 0.6 pts | 0.9 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | btree-map | 8 | 1061 | 745 | 0.71× [0.67, 0.74] | -41.7% | [-48.7%, -34.7%] | 5.3 pts | 4.6 | no | yes |
| single-value | street | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1092 | 4112 | 3.79× [3.75, 3.83] | +73.6% | [+73.4%, +73.9%] | 0.5 pts | 1.5 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1094 | 3193 | 2.93× [2.89, 2.97] | +65.8% | [+65.4%, +66.3%] | 0.5 pts | 1.3 | yes | yes |
| single-value | street | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1088 | 2456 | 2.26× [2.23, 2.29] | +55.8% | [+55.2%, +56.3%] | 0.4 pts | 0.7 | yes | yes |
| single-value | street | 65536 | churn | ordered | baseline | 8 | 245 | 279 | 1.14× [1.11, 1.17] | +12.3% | [+10.0%, +14.5%] | 1.5 pts | 1.0 | no | yes |
| single-value | street | 65536 | churn | ordered | btree-map | 8 | 258 | 327 | 1.28× [1.27, 1.29] | +22.0% | [+21.5%, +22.5%] | 0.7 pts | 1.2 | yes | yes |
| single-value | street | 65536 | churn | ordered | ordered-mk1 | 8 | 256 | 336 | 1.30× [1.21, 1.40] | +22.8% | [+17.2%, +28.4%] | 5.0 pts | 4.7 | no | yes |
| single-value | street | 65536 | churn | ordered | ordered-mk2 | 8 | 243 | 365 | 1.50× [1.49, 1.51] | +33.3% | [+33.0%, +33.6%] | 0.6 pts | 0.6 | yes | yes |
| single-value | street | 65536 | churn | ordered | ordered-mk4 | 8 | 242 | 351 | 1.44× [1.43, 1.46] | +30.8% | [+30.0%, +31.5%] | 1.2 pts | 1.2 | yes | yes |
| single-value | street | 65536 | build | ordered | baseline | 8 | 55.25 ms | 53.73 ms | 0.97× [0.95, 1.00] | -2.6% | [-5.6%, +0.3%] | 2.1 pts | 9.8 | no | no |
| single-value | street | 65536 | build | ordered | btree-map | 8 | 55.69 ms | 66.16 ms | 1.20× [1.18, 1.21] | +16.5% | [+15.5%, +17.6%] | 1.2 pts | 8.8 | yes | yes |
| single-value | street | 65536 | build | ordered | ordered-mk1 | 8 | 55.24 ms | 58.70 ms | 1.05× [1.03, 1.09] | +5.2% | [+2.5%, +7.9%] | 2.2 pts | 9.9 | no | yes |
| single-value | street | 65536 | build | ordered | ordered-mk2 | 8 | 55.02 ms | 79.43 ms | 1.43× [1.38, 1.49] | +30.3% | [+27.4%, +33.1%] | 2.2 pts | 19.8 | yes | yes |
| single-value | street | 65536 | build | ordered | ordered-mk4 | 8 | 54.98 ms | 76.69 ms | 1.38× [1.33, 1.43] | +27.5% | [+24.6%, +30.3%] | 2.1 pts | 18.9 | no | yes |
| single-value | u64 | 4096 | valuesFor | ordered | baseline | 8 | 40.1 | 20.2 | 0.50× [0.50, 0.51] | -98.8% | [-100.8%, -96.9%] | 1.5 pts | 5.9 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | btree-map | 8 | 40.1 | 108 | 2.70× [2.68, 2.71] | +62.9% | [+62.7%, +63.1%] | 0.4 pts | 7.0 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 40.1 | 23.5 | 0.59× [0.58, 0.59] | -70.8% | [-71.5%, -70.0%] | 0.8 pts | 4.5 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 40.4 | 24.3 | 0.60× [0.59, 0.61] | -66.1% | [-68.3%, -63.9%] | 1.9 pts | 9.4 | yes | yes |
| single-value | u64 | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 40.1 | 24.2 | 0.60× [0.60, 0.61] | -66.1% | [-67.4%, -64.9%] | 1.2 pts | 5.6 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 348 | 1231 | 3.54× [3.53, 3.54] | +71.7% | [+71.7%, +71.8%] | 0.0 pts | 1.0 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | btree-map | 8 | 348 | 441 | 1.27× [1.27, 1.27] | +21.2% | [+21.0%, +21.3%] | 0.2 pts | 1.6 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 348 | 1989 | 5.72× [5.71, 5.73] | +82.5% | [+82.5%, +82.6%] | 0.0 pts | 1.8 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 349 | 1842 | 5.29× [5.28, 5.30] | +81.1% | [+81.1%, +81.1%] | 0.0 pts | 2.0 | yes | yes |
| single-value | u64 | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 348 | 1836 | 5.28× [5.27, 5.29] | +81.1% | [+81.0%, +81.1%] | 0.0 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | baseline | 8 | 83.9 | 54.3 | 0.65× [0.64, 0.65] | -54.1% | [-55.5%, -52.8%] | 1.0 pts | 2.1 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | btree-map | 8 | 83.1 | 152 | 1.83× [1.83, 1.84] | +45.5% | [+45.4%, +45.6%] | 0.2 pts | 0.8 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | ordered-mk1 | 8 | 82.3 | 63.0 | 0.77× [0.76, 0.77] | -30.6% | [-31.0%, -30.2%] | 0.6 pts | 2.1 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | ordered-mk2 | 8 | 82.3 | 77.3 | 0.94× [0.94, 0.94] | -6.5% | [-6.8%, -6.1%] | 0.3 pts | 1.2 | yes | yes |
| single-value | u64 | 4096 | churn | ordered | ordered-mk4 | 8 | 82.5 | 77.3 | 0.94× [0.93, 0.94] | -6.7% | [-7.0%, -6.4%] | 0.4 pts | 1.3 | yes | yes |
| single-value | u64 | 4096 | build | ordered | baseline | 8 | 1.25 ms | 917.1 µs | 0.73× [0.71, 0.76] | -36.2% | [-40.4%, -31.9%] | 2.5 pts | 2.7 | no | yes |
| single-value | u64 | 4096 | build | ordered | btree-map | 8 | 1.24 ms | 2.22 ms | 1.79× [1.78, 1.80] | +44.1% | [+43.8%, +44.5%] | 0.3 pts | 0.7 | yes | yes |
| single-value | u64 | 4096 | build | ordered | ordered-mk1 | 8 | 1.25 ms | 1.07 ms | 0.85× [0.83, 0.88] | -17.4% | [-20.6%, -14.2%] | 2.1 pts | 1.9 | no | yes |
| single-value | u64 | 4096 | build | ordered | ordered-mk2 | 8 | 1.24 ms | 1.41 ms | 1.13× [1.12, 1.15] | +11.8% | [+10.4%, +13.1%] | 1.5 pts | 2.7 | yes | yes |
| single-value | u64 | 4096 | build | ordered | ordered-mk4 | 8 | 1.25 ms | 1.48 ms | 1.18× [1.17, 1.20] | +15.5% | [+14.2%, +16.8%] | 1.5 pts | 2.3 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | baseline | 8 | 30.8 | 24.9 | 0.81× [0.80, 0.82] | -23.8% | [-25.3%, -22.3%] | 1.1 pts | 6.0 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | btree-map | 8 | 31.1 | 145 | 4.67× [4.64, 4.70] | +78.6% | [+78.4%, +78.7%] | 0.1 pts | 2.6 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 30.9 | 28.6 | 0.93× [0.92, 0.93] | -8.1% | [-8.7%, -7.4%] | 0.6 pts | 3.4 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 30.9 | 30.4 | 0.99× [0.98, 0.99] | -1.4% | [-1.8%, -1.0%] | 0.3 pts | 1.7 | yes | yes |
| single-value | u64 | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 30.8 | 30.8 | 1.00× [0.99, 1.00] | -0.2% | [-0.7%, +0.3%] | 0.3 pts | 1.7 | yes | no |
| single-value | u64 | 16384 | valuesBetween | ordered | baseline | 8 | 2191 | 1892 | 0.86× [0.85, 0.87] | -15.9% | [-17.1%, -14.7%] | 0.8 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | btree-map | 8 | 2129 | 501 | 0.24× [0.23, 0.24] | -324.6% | [-326.7%, -322.5%] | 1.4 pts | 1.5 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 2185 | 2740 | 1.25× [1.25, 1.26] | +20.1% | [+19.8%, +20.5%] | 0.5 pts | 1.1 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 2154 | 2216 | 1.03× [1.03, 1.03] | +2.9% | [+2.6%, +3.2%] | 0.4 pts | 1.2 | yes | yes |
| single-value | u64 | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 2152 | 2161 | 1.00× [1.00, 1.01] | +0.4% | [+0.2%, +0.6%] | 0.5 pts | 1.4 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | baseline | 8 | 74.0 | 52.6 | 0.71× [0.71, 0.71] | -41.1% | [-41.3%, -40.9%] | 0.3 pts | 0.3 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | btree-map | 8 | 73.0 | 195 | 2.69× [2.65, 2.72] | +62.8% | [+62.3%, +63.2%] | 0.4 pts | 2.0 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | ordered-mk1 | 8 | 73.7 | 65.7 | 0.89× [0.89, 0.90] | -11.9% | [-12.1%, -11.7%] | 0.5 pts | 1.0 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | ordered-mk2 | 8 | 73.5 | 75.3 | 1.03× [1.03, 1.03] | +2.7% | [+2.6%, +2.8%] | 0.2 pts | 0.7 | yes | yes |
| single-value | u64 | 16384 | churn | ordered | ordered-mk4 | 8 | 73.5 | 73.6 | 1.00× [1.00, 1.00] | -0.1% | [-0.4%, +0.1%] | 0.2 pts | 0.9 | yes | no |
| single-value | u64 | 16384 | build | ordered | baseline | 8 | 6.07 ms | 3.90 ms | 0.64× [0.63, 0.65] | -55.7% | [-57.7%, -53.8%] | 1.4 pts | 0.8 | yes | yes |
| single-value | u64 | 16384 | build | ordered | btree-map | 8 | 6.10 ms | 11.28 ms | 1.86× [1.84, 1.87] | +46.1% | [+45.6%, +46.6%] | 0.5 pts | 0.9 | yes | yes |
| single-value | u64 | 16384 | build | ordered | ordered-mk1 | 8 | 6.05 ms | 4.61 ms | 0.76× [0.75, 0.77] | -31.1% | [-33.0%, -29.2%] | 1.7 pts | 1.3 | yes | yes |
| single-value | u64 | 16384 | build | ordered | ordered-mk2 | 8 | 6.02 ms | 5.34 ms | 0.89× [0.87, 0.91] | -12.7% | [-15.2%, -10.2%] | 1.9 pts | 1.2 | no | yes |
| single-value | u64 | 16384 | build | ordered | ordered-mk4 | 8 | 6.01 ms | 5.18 ms | 0.87× [0.85, 0.88] | -15.5% | [-17.4%, -13.7%] | 1.3 pts | 1.4 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | baseline | 8 | 40.5 | 31.2 | 0.79× [0.74, 0.85] | -26.7% | [-35.8%, -17.6%] | 7.0 pts | 1.2 | no | yes |
| single-value | u64 | 65536 | valuesFor | ordered | btree-map | 8 | 38.8 | 188 | 4.77× [4.42, 5.19] | +79.1% | [+77.4%, +80.7%] | 1.1 pts | 8.1 | yes | yes |
| single-value | u64 | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 38.0 | 28.3 | 0.77× [0.75, 0.79] | -30.4% | [-34.0%, -26.8%] | 6.4 pts | 3.7 | no | yes |
| single-value | u64 | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 38.8 | 37.7 | 0.99× [0.97, 1.00] | -1.4% | [-2.6%, -0.3%] | 1.6 pts | 0.6 | yes | no |
| single-value | u64 | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 38.7 | 38.2 | 1.01× [0.99, 1.03] | +0.7% | [-1.1%, +2.6%] | 2.4 pts | 1.3 | yes | no |
| single-value | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 1897 | 2265 | 1.20× [1.19, 1.21] | +16.4% | [+15.7%, +17.2%] | 1.6 pts | 3.3 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | btree-map | 8 | 1869 | 563 | 0.30× [0.30, 0.30] | -231.1% | [-234.0%, -228.2%] | 3.6 pts | 5.5 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 1904 | 3260 | 1.72× [1.70, 1.73] | +41.8% | [+41.3%, +42.3%] | 0.5 pts | 2.0 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 1882 | 2385 | 1.27× [1.26, 1.27] | +21.1% | [+20.8%, +21.4%] | 0.3 pts | 1.1 | yes | yes |
| single-value | u64 | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 1864 | 1898 | 1.01× [1.00, 1.02] | +1.4% | [+0.5%, +2.4%] | 0.6 pts | 3.4 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | baseline | 8 | 123 | 105 | 0.85× [0.83, 0.86] | -18.0% | [-20.0%, -16.0%] | 1.6 pts | 1.3 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | btree-map | 8 | 135 | 269 | 2.00× [1.93, 2.07] | +50.0% | [+48.2%, +51.8%] | 1.4 pts | 2.0 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | ordered-mk1 | 8 | 121 | 122 | 1.00× [0.99, 1.00] | -0.4% | [-1.1%, +0.3%] | 1.1 pts | 1.2 | yes | no |
| single-value | u64 | 65536 | churn | ordered | ordered-mk2 | 8 | 123 | 140 | 1.16× [1.15, 1.18] | +14.1% | [+13.0%, +15.2%] | 1.1 pts | 1.7 | yes | yes |
| single-value | u64 | 65536 | churn | ordered | ordered-mk4 | 8 | 118 | 121 | 1.03× [1.02, 1.03] | +2.4% | [+1.8%, +3.1%] | 0.7 pts | 0.6 | yes | yes |
| single-value | u64 | 65536 | build | ordered | baseline | 8 | 25.30 ms | 19.70 ms | 0.78× [0.77, 0.78] | -28.5% | [-29.5%, -27.4%] | 1.3 pts | 2.2 | yes | yes |
| single-value | u64 | 65536 | build | ordered | btree-map | 8 | 25.31 ms | 57.69 ms | 2.28× [2.26, 2.30] | +56.1% | [+55.7%, +56.4%] | 0.4 pts | 1.8 | yes | yes |
| single-value | u64 | 65536 | build | ordered | ordered-mk1 | 8 | 25.38 ms | 23.53 ms | 0.93× [0.92, 0.93] | -7.9% | [-8.2%, -7.5%] | 0.4 pts | 0.8 | yes | yes |
| single-value | u64 | 65536 | build | ordered | ordered-mk2 | 8 | 25.32 ms | 29.58 ms | 1.17× [1.14, 1.19] | +14.2% | [+12.4%, +15.9%] | 1.6 pts | 3.7 | yes | yes |
| single-value | u64 | 65536 | build | ordered | ordered-mk4 | 8 | 25.26 ms | 25.86 ms | 1.02× [1.00, 1.04] | +1.9% | [-0.1%, +3.9%] | 1.9 pts | 4.2 | no | no |
| single-value | url | 4096 | valuesFor | ordered | baseline | 8 | 111 | 94.8 | 0.85× [0.83, 0.86] | -18.1% | [-20.1%, -16.1%] | 1.3 pts | 3.9 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | btree-map | 8 | 112 | 138 | 1.23× [1.22, 1.25] | +18.8% | [+17.7%, +19.8%] | 0.9 pts | 1.4 | yes | yes |
| single-value | url | 4096 | valuesFor | ordered | ordered-mk1 | 8 | 112 | 95.1 | 0.85× [0.83, 0.86] | -17.9% | [-20.1%, -15.6%] | 1.4 pts | 1.9 | no | yes |
| single-value | url | 4096 | valuesFor | ordered | ordered-mk2 | 8 | 112 | 105 | 0.93× [0.92, 0.95] | -7.1% | [-9.2%, -5.0%] | 1.3 pts | 1.8 | no | yes |
| single-value | url | 4096 | valuesFor | ordered | ordered-mk4 | 8 | 112 | 109 | 0.97× [0.96, 0.99] | -2.9% | [-4.5%, -1.2%] | 1.1 pts | 1.9 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | baseline | 8 | 1717 | 2484 | 1.46× [1.43, 1.48] | +31.3% | [+30.2%, +32.5%] | 1.2 pts | 1.0 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | btree-map | 8 | 1648 | 532 | 0.32× [0.32, 0.33] | -209.4% | [-211.9%, -206.9%] | 1.8 pts | 0.9 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | ordered-mk1 | 8 | 1721 | 3916 | 2.29× [2.26, 2.33] | +56.4% | [+55.7%, +57.0%] | 0.6 pts | 0.9 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | ordered-mk2 | 8 | 1722 | 3070 | 1.80× [1.76, 1.83] | +44.3% | [+43.2%, +45.4%] | 1.0 pts | 1.1 | yes | yes |
| single-value | url | 4096 | valuesBetween | ordered | ordered-mk4 | 8 | 1701 | 2333 | 1.38× [1.36, 1.40] | +27.4% | [+26.5%, +28.3%] | 1.0 pts | 0.9 | yes | yes |
| single-value | url | 4096 | churn | ordered | baseline | 8 | 253 | 206 | 0.82× [0.81, 0.82] | -22.0% | [-22.8%, -21.2%] | 0.9 pts | 1.4 | yes | yes |
| single-value | url | 4096 | churn | ordered | btree-map | 8 | 252 | 198 | 0.79× [0.78, 0.80] | -26.8% | [-27.7%, -25.8%] | 0.9 pts | 1.7 | yes | yes |
| single-value | url | 4096 | churn | ordered | ordered-mk1 | 8 | 252 | 228 | 0.91× [0.90, 0.92] | -9.9% | [-11.1%, -8.7%] | 0.8 pts | 1.5 | yes | yes |
| single-value | url | 4096 | churn | ordered | ordered-mk2 | 8 | 252 | 309 | 1.23× [1.23, 1.23] | +18.8% | [+18.8%, +18.9%] | 0.5 pts | 2.0 | yes | yes |
| single-value | url | 4096 | churn | ordered | ordered-mk4 | 8 | 251 | 290 | 1.16× [1.15, 1.17] | +14.0% | [+13.4%, +14.6%] | 0.4 pts | 1.2 | yes | yes |
| single-value | url | 4096 | build | ordered | baseline | 8 | 3.90 ms | 2.86 ms | 0.73× [0.72, 0.74] | -36.6% | [-39.0%, -34.2%] | 1.6 pts | 2.1 | yes | yes |
| single-value | url | 4096 | build | ordered | btree-map | 8 | 3.92 ms | 2.78 ms | 0.72× [0.71, 0.72] | -39.4% | [-40.7%, -38.1%] | 1.3 pts | 1.7 | yes | yes |
| single-value | url | 4096 | build | ordered | ordered-mk1 | 8 | 3.90 ms | 3.17 ms | 0.81× [0.81, 0.82] | -23.0% | [-23.6%, -22.4%] | 0.8 pts | 1.3 | yes | yes |
| single-value | url | 4096 | build | ordered | ordered-mk2 | 8 | 3.88 ms | 4.53 ms | 1.17× [1.16, 1.17] | +14.3% | [+14.1%, +14.5%] | 1.1 pts | 2.3 | yes | yes |
| single-value | url | 4096 | build | ordered | ordered-mk4 | 8 | 3.86 ms | 4.46 ms | 1.15× [1.13, 1.17] | +12.9% | [+11.5%, +14.3%] | 1.6 pts | 2.3 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | baseline | 8 | 134 | 121 | 0.90× [0.89, 0.92] | -10.8% | [-13.0%, -8.6%] | 1.7 pts | 4.4 | no | yes |
| single-value | url | 16384 | valuesFor | ordered | btree-map | 8 | 135 | 192 | 1.42× [1.40, 1.43] | +29.5% | [+28.7%, +30.2%] | 0.5 pts | 1.3 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | ordered-mk1 | 8 | 134 | 119 | 0.88× [0.87, 0.90] | -13.0% | [-14.7%, -11.3%] | 1.2 pts | 1.5 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | ordered-mk2 | 8 | 135 | 130 | 0.97× [0.96, 0.97] | -3.6% | [-4.1%, -3.1%] | 0.4 pts | 0.6 | yes | yes |
| single-value | url | 16384 | valuesFor | ordered | ordered-mk4 | 8 | 135 | 133 | 0.99× [0.99, 0.99] | -1.0% | [-1.4%, -0.7%] | 0.4 pts | 0.8 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | baseline | 8 | 1876 | 2702 | 1.44× [1.42, 1.46] | +30.6% | [+29.8%, +31.5%] | 0.7 pts | 0.9 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | btree-map | 8 | 1830 | 610 | 0.33× [0.33, 0.34] | -199.4% | [-201.6%, -197.3%] | 2.3 pts | 1.3 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | ordered-mk1 | 8 | 1879 | 4136 | 2.21× [2.20, 2.23] | +54.8% | [+54.5%, +55.2%] | 0.4 pts | 0.8 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | ordered-mk2 | 8 | 1867 | 3250 | 1.74× [1.72, 1.76] | +42.6% | [+42.0%, +43.2%] | 0.5 pts | 0.8 | yes | yes |
| single-value | url | 16384 | valuesBetween | ordered | ordered-mk4 | 8 | 1870 | 2530 | 1.35× [1.32, 1.37] | +25.8% | [+24.5%, +27.0%] | 0.8 pts | 0.9 | yes | yes |
| single-value | url | 16384 | churn | ordered | baseline | 8 | 303 | 264 | 0.89× [0.85, 0.94] | -11.9% | [-17.7%, -6.1%] | 4.2 pts | 7.7 | no | yes |
| single-value | url | 16384 | churn | ordered | btree-map | 8 | 304 | 261 | 0.86× [0.86, 0.87] | -16.1% | [-16.8%, -15.5%] | 0.5 pts | 0.5 | yes | yes |
| single-value | url | 16384 | churn | ordered | ordered-mk1 | 8 | 298 | 281 | 0.94× [0.94, 0.94] | -6.2% | [-6.4%, -5.9%] | 0.6 pts | 1.7 | yes | yes |
| single-value | url | 16384 | churn | ordered | ordered-mk2 | 8 | 299 | 368 | 1.24× [1.22, 1.26] | +19.5% | [+18.3%, +20.7%] | 1.0 pts | 3.0 | yes | yes |
| single-value | url | 16384 | churn | ordered | ordered-mk4 | 8 | 298 | 347 | 1.16× [1.16, 1.17] | +14.1% | [+13.8%, +14.4%] | 0.3 pts | 0.8 | yes | yes |
| single-value | url | 16384 | build | ordered | baseline | 8 | 17.95 ms | 13.52 ms | 0.75× [0.75, 0.76] | -32.8% | [-33.8%, -31.7%] | 1.4 pts | 2.3 | yes | yes |
| single-value | url | 16384 | build | ordered | btree-map | 8 | 17.92 ms | 14.17 ms | 0.79× [0.79, 0.80] | -26.1% | [-26.7%, -25.6%] | 0.6 pts | 1.5 | yes | yes |
| single-value | url | 16384 | build | ordered | ordered-mk1 | 8 | 17.92 ms | 14.68 ms | 0.82× [0.81, 0.83] | -22.0% | [-23.4%, -20.7%] | 1.1 pts | 3.3 | yes | yes |
| single-value | url | 16384 | build | ordered | ordered-mk2 | 8 | 17.87 ms | 20.84 ms | 1.16× [1.15, 1.17] | +14.0% | [+13.3%, +14.8%] | 0.6 pts | 2.0 | yes | yes |
| single-value | url | 16384 | build | ordered | ordered-mk4 | 8 | 17.85 ms | 20.25 ms | 1.12× [1.11, 1.13] | +10.9% | [+10.0%, +11.9%] | 1.0 pts | 3.6 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | baseline | 8 | 194 | 213 | 1.13× [1.11, 1.15] | +11.3% | [+10.0%, +12.7%] | 2.2 pts | 1.0 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | btree-map | 8 | 196 | 269 | 1.39× [1.35, 1.43] | +27.9% | [+25.9%, +29.9%] | 2.0 pts | 1.0 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | ordered-mk1 | 8 | 185 | 196 | 1.05× [1.02, 1.08] | +4.9% | [+2.2%, +7.7%] | 2.1 pts | 0.8 | no | yes |
| single-value | url | 65536 | valuesFor | ordered | ordered-mk2 | 8 | 187 | 206 | 1.10× [1.08, 1.13] | +9.4% | [+7.6%, +11.1%] | 1.3 pts | 0.5 | yes | yes |
| single-value | url | 65536 | valuesFor | ordered | ordered-mk4 | 8 | 189 | 200 | 1.07× [1.05, 1.09] | +6.9% | [+5.2%, +8.6%] | 1.6 pts | 0.7 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | baseline | 8 | 2095 | 3631 | 1.72× [1.68, 1.76] | +41.8% | [+40.4%, +43.2%] | 1.7 pts | 1.3 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | btree-map | 8 | 1977 | 1020 | 0.52× [0.51, 0.52] | -93.6% | [-96.6%, -90.6%] | 6.0 pts | 0.5 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | ordered-mk1 | 8 | 2064 | 4785 | 2.31× [2.25, 2.38] | +56.8% | [+55.5%, +58.0%] | 1.1 pts | 1.1 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | ordered-mk2 | 8 | 2086 | 3762 | 1.80× [1.76, 1.85] | +44.5% | [+43.0%, +45.9%] | 1.2 pts | 1.0 | yes | yes |
| single-value | url | 65536 | valuesBetween | ordered | ordered-mk4 | 8 | 2070 | 2908 | 1.41× [1.39, 1.42] | +28.9% | [+28.3%, +29.6%] | 0.8 pts | 0.6 | yes | yes |
| single-value | url | 65536 | churn | ordered | baseline | 8 | 400 | 413 | 1.04× [1.03, 1.04] | +3.5% | [+3.1%, +4.0%] | 1.1 pts | 0.5 | yes | yes |
| single-value | url | 65536 | churn | ordered | btree-map | 8 | 457 | 451 | 1.00× [0.97, 1.03] | +0.2% | [-2.6%, +3.0%] | 3.3 pts | 3.3 | no | no |
| single-value | url | 65536 | churn | ordered | ordered-mk1 | 8 | 388 | 404 | 1.04× [1.03, 1.05] | +3.8% | [+3.2%, +4.3%] | 1.0 pts | 0.7 | yes | yes |
| single-value | url | 65536 | churn | ordered | ordered-mk2 | 8 | 426 | 518 | 1.23× [1.19, 1.27] | +18.7% | [+16.3%, +21.1%] | 3.0 pts | 1.0 | no | yes |
| single-value | url | 65536 | churn | ordered | ordered-mk4 | 8 | 395 | 467 | 1.16× [1.15, 1.17] | +13.8% | [+13.1%, +14.5%] | 1.6 pts | 1.0 | yes | yes |
| single-value | url | 65536 | build | ordered | baseline | 8 | 92.53 ms | 79.89 ms | 0.86× [0.85, 0.87] | -15.8% | [-17.0%, -14.6%] | 1.1 pts | 6.2 | yes | yes |
| single-value | url | 65536 | build | ordered | btree-map | 8 | 92.63 ms | 81.41 ms | 0.88× [0.87, 0.89] | -13.6% | [-15.3%, -12.0%] | 1.4 pts | 6.4 | yes | yes |
| single-value | url | 65536 | build | ordered | ordered-mk1 | 8 | 92.91 ms | 81.94 ms | 0.89× [0.89, 0.89] | -12.6% | [-12.7%, -12.5%] | 1.2 pts | 4.7 | yes | yes |
| single-value | url | 65536 | build | ordered | ordered-mk2 | 8 | 93.13 ms | 108.59 ms | 1.17× [1.17, 1.18] | +14.9% | [+14.4%, +15.3%] | 1.6 pts | 8.1 | yes | yes |
| single-value | url | 65536 | build | ordered | ordered-mk4 | 8 | 92.96 ms | 105.00 ms | 1.13× [1.13, 1.14] | +11.7% | [+11.3%, +12.1%] | 1.0 pts | 7.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +1.01% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 valuesBetween: ordered vs ordered-mk1: the A/A validations found a systematic difference of +1.14% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.32% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=4096 churn: ordered vs ordered-mk1: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=4096 churn: ordered vs ordered-mk4: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs ordered-mk1: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs ordered-mk1: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs ordered-mk4: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs ordered-mk1: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs ordered-mk2: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs ordered-mk4: the A/A validations found a systematic difference of +0.13% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesFor: ordered vs ordered-mk2: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -2.03% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -1.27% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=65536 churn: ordered vs ordered-mk2: the pooled difference of 2.44% does not clear the 3.01% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 churn: ordered vs ordered-mk2: the pooled interval [-0.94%, 5.82%] includes zero
- natural dirs n=65536 churn: ordered vs ordered-mk4: the pooled difference of -1.16% does not clear the 2.26% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=65536 churn: ordered vs ordered-mk4: the pooled interval [-4.64%, 2.33%] includes zero
- natural dirs n=65536 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 build: ordered vs ordered-mk1: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 build: ordered vs ordered-mk2: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 build: ordered vs ordered-mk4: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesFor: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesFor: ordered vs ordered-mk4: the A/A validations found a systematic difference of +0.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +1.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +1.48% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesBetween: ordered vs ordered-mk1: the A/A validations found a systematic difference of +1.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.99% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesBetween: ordered vs ordered-mk4: the A/A validations found a systematic difference of +1.07% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 churn: ordered vs ordered-mk2: the pooled interval [-1.17%, 1.81%] includes zero
- natural street n=4096 churn: ordered vs ordered-mk2: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 build: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 valuesFor: ordered vs ordered-mk2: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs baseline: the processes scatter 9.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs ordered-mk1: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 churn: ordered vs ordered-mk4: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=16384 build: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: the pooled difference of 0.20% does not clear the 0.56% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 valuesFor: ordered vs baseline: the pooled interval [-2.03%, 2.42%] includes zero
- natural street n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.00% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 valuesFor: ordered vs ordered-mk2: the pooled interval [-0.10%, 1.34%] includes zero
- natural street n=65536 valuesFor: ordered vs ordered-mk4: the pooled interval [-0.68%, 4.30%] includes zero
- natural street n=65536 valuesFor: ordered vs ordered-mk4: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -2.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 churn: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs ordered-mk1: the pooled difference of -1.29% does not clear the 1.83% noise floor, the bound on what the harness reports between identical code in every process
- natural street n=65536 churn: ordered vs ordered-mk1: the pooled interval [-6.40%, 3.82%] includes zero
- natural street n=65536 churn: ordered vs ordered-mk2: the A/A validations found a systematic difference of -2.04% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 churn: ordered vs ordered-mk4: the A/A validations found a systematic difference of -1.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=65536 build: ordered vs baseline: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs ordered-mk1: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs ordered-mk2: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 build: ordered vs ordered-mk4: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=4096 valuesBetween: ordered vs ordered-mk1: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 churn: ordered vs ordered-mk4: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 build: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesFor: ordered vs btree-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesFor: ordered vs ordered-mk2: the pooled interval [-4.20%, 0.41%] includes zero
- natural u64 n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-3.28%, 1.11%] includes zero
- natural u64 n=16384 valuesFor: ordered vs ordered-mk4: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesFor: ordered vs ordered-mk4: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 valuesBetween: ordered vs ordered-mk4: the pooled interval [-1.47%, 0.19%] includes zero
- natural u64 n=16384 churn: ordered vs baseline: the A/A validations found a systematic difference of -1.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=16384 churn: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs ordered-mk1: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=16384 churn: ordered vs ordered-mk4: the pooled difference of 0.62% does not clear the 1.03% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=16384 churn: ordered vs ordered-mk4: the pooled interval [-0.92%, 2.16%] includes zero
- natural u64 n=16384 churn: ordered vs ordered-mk4: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=16384 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 valuesFor: ordered vs ordered-mk2: the pooled difference of -1.05% does not clear the 1.82% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=65536 valuesFor: ordered vs ordered-mk2: the pooled interval [-7.64%, 5.55%] includes zero
- natural u64 n=65536 valuesFor: ordered vs ordered-mk4: the pooled difference of 0.00% does not clear the 1.12% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=65536 valuesFor: ordered vs ordered-mk4: the pooled interval [-4.93%, 4.93%] includes zero
- natural u64 n=65536 valuesBetween: ordered vs ordered-mk4: the pooled interval [-0.12%, 2.14%] includes zero
- natural u64 n=65536 churn: ordered vs btree-sets: the processes scatter 7.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 churn: ordered vs ordered-mk1: the A/A validations found a systematic difference of -1.85% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 churn: ordered vs ordered-mk1: the pooled interval [-16.41%, 4.93%] includes zero
- natural u64 n=65536 churn: ordered vs ordered-mk2: the pooled difference of -1.51% does not clear the 4.17% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=65536 churn: ordered vs ordered-mk2: the pooled interval [-7.33%, 4.31%] includes zero
- natural u64 n=65536 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 build: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=4096 valuesFor: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesFor: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesFor: ordered vs ordered-mk1: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesFor: ordered vs ordered-mk2: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesFor: ordered vs ordered-mk4: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesBetween: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesBetween: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=4096 valuesBetween: ordered vs ordered-mk1: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesBetween: ordered vs ordered-mk2: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 valuesBetween: ordered vs ordered-mk4: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 churn: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 churn: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 churn: ordered vs ordered-mk1: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 churn: ordered vs ordered-mk2: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 churn: ordered vs ordered-mk2: the pooled difference of 0.29% does not clear the 0.42% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=4096 churn: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=4096 churn: ordered vs ordered-mk2: the pooled interval [-1.21%, 1.80%] includes zero
- natural url n=4096 churn: ordered vs ordered-mk2: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=4096 churn: ordered vs ordered-mk2: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural url n=4096 churn: ordered vs ordered-mk4: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 build: ordered vs baseline: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 build: ordered vs btree-sets: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 build: ordered vs ordered-mk1: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 build: ordered vs ordered-mk2: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=4096 build: ordered vs ordered-mk4: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- natural url n=16384 valuesFor: ordered vs baseline: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-4.30%, 0.87%] includes zero
- natural url n=16384 valuesFor: ordered vs ordered-mk4: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=16384 churn: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 churn: ordered vs ordered-mk4: the pooled interval [-0.46%, 3.10%] includes zero
- natural url n=16384 build: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.06% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=16384 build: ordered vs btree-sets: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 build: ordered vs ordered-mk1: the processes scatter 12.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 build: ordered vs ordered-mk2: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=16384 build: ordered vs ordered-mk4: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +1.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 valuesFor: ordered vs ordered-mk2: the A/A validations found a systematic difference of -0.88% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of -3.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -2.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 churn: ordered vs ordered-mk2: the A/A validations found a systematic difference of -1.75% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 churn: ordered vs ordered-mk4: the A/A validations found a systematic difference of -2.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 build: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 build: ordered vs ordered-mk4: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=4096 valuesBetween: ordered vs ordered-mk1: the A/A validations found a systematic difference of +0.86% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +1.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.22% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 valuesFor: ordered vs ordered-mk2: the pooled difference of -0.30% does not clear the 0.62% noise floor, the bound on what the harness reports between identical code in every process
- single-value dirs n=16384 valuesFor: ordered vs ordered-mk2: the pooled interval [-1.18%, 0.57%] includes zero
- single-value dirs n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-1.79%, 0.54%] includes zero
- single-value dirs n=16384 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.72% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 valuesBetween: ordered vs ordered-mk1: the A/A validations found a systematic difference of +0.91% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value dirs n=16384 churn: ordered vs ordered-mk1: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 build: ordered vs ordered-mk1: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 build: ordered vs ordered-mk2: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=16384 build: ordered vs ordered-mk4: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 valuesFor: ordered vs ordered-mk2: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 churn: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs baseline: the processes scatter 13.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: the pooled interval [-0.71%, 2.12%] includes zero
- single-value dirs n=65536 build: ordered vs btree-map: the processes scatter 9.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs btree-map: 5 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value dirs n=65536 build: ordered vs ordered-mk1: the processes scatter 5.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs ordered-mk2: the processes scatter 6.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value dirs n=65536 build: ordered vs ordered-mk4: the processes scatter 6.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 valuesFor: ordered vs ordered-mk1: the A/A validations found a systematic difference of +0.20% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.99% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesBetween: ordered vs ordered-mk1: the A/A validations found a systematic difference of +1.52% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +1.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=4096 churn: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=4096 build: ordered vs ordered-mk2: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 churn: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs ordered-mk1: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 churn: ordered vs ordered-mk2: the A/A validations found a systematic difference of -0.17% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=16384 build: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs ordered-mk2: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=16384 build: ordered vs ordered-mk4: the A/A validations found a systematic difference of +0.27% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=65536 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesFor: ordered vs ordered-mk1: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value street n=65536 valuesFor: ordered vs ordered-mk2: the pooled interval [-0.09%, 5.60%] includes zero
- single-value street n=65536 valuesFor: ordered vs ordered-mk2: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 valuesBetween: ordered vs btree-map: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 churn: ordered vs ordered-mk1: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs baseline: the pooled interval [-5.56%, 0.31%] includes zero
- single-value street n=65536 build: ordered vs baseline: the processes scatter 9.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs btree-map: the processes scatter 8.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs ordered-mk1: the processes scatter 9.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs ordered-mk2: the processes scatter 19.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=65536 build: ordered vs ordered-mk4: the processes scatter 18.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs baseline: the processes scatter 5.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 7.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs ordered-mk1: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.05% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=4096 valuesFor: ordered vs ordered-mk2: the processes scatter 9.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 valuesFor: ordered vs ordered-mk4: the processes scatter 5.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 churn: ordered vs ordered-mk1: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 build: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 build: ordered vs ordered-mk2: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=4096 build: ordered vs ordered-mk4: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 6.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs ordered-mk1: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=16384 valuesFor: ordered vs ordered-mk4: the pooled interval [-0.69%, 0.25%] includes zero
- single-value u64 n=16384 churn: ordered vs ordered-mk4: the pooled difference of -0.13% does not clear the 0.32% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=16384 churn: ordered vs ordered-mk4: the pooled interval [-0.36%, 0.10%] includes zero
- single-value u64 n=65536 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +2.07% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=65536 valuesFor: ordered vs btree-map: the processes scatter 8.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk1: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk2: the pooled difference of -1.45% does not clear the 1.87% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk4: the pooled difference of 0.74% does not clear the 0.86% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=65536 valuesFor: ordered vs ordered-mk4: the pooled interval [-1.08%, 2.57%] includes zero
- single-value u64 n=65536 valuesBetween: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesBetween: ordered vs btree-map: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesBetween: ordered vs ordered-mk1: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 valuesBetween: ordered vs ordered-mk4: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 churn: ordered vs baseline: the A/A validations found a systematic difference of +0.69% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value u64 n=65536 churn: ordered vs ordered-mk1: the pooled difference of -0.42% does not clear the 0.51% noise floor, the bound on what the harness reports between identical code in every process
- single-value u64 n=65536 churn: ordered vs ordered-mk1: the pooled interval [-1.14%, 0.30%] includes zero
- single-value u64 n=65536 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 build: ordered vs ordered-mk2: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 build: ordered vs ordered-mk4: the pooled interval [-0.14%, 3.86%] includes zero
- single-value u64 n=65536 build: ordered vs ordered-mk4: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=65536 build: ordered vs ordered-mk4: 5 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value url n=4096 valuesFor: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=4096 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.89% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=4096 valuesBetween: ordered vs ordered-mk1: the A/A validations found a systematic difference of +1.02% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=4096 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.78% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=4096 valuesBetween: ordered vs ordered-mk4: the A/A validations found a systematic difference of +0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=4096 churn: ordered vs ordered-mk2: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=4096 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=4096 build: ordered vs ordered-mk2: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=4096 build: ordered vs ordered-mk4: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 valuesFor: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 churn: ordered vs baseline: the processes scatter 7.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 churn: ordered vs ordered-mk2: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 build: ordered vs ordered-mk1: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=16384 build: ordered vs ordered-mk4: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 valuesFor: ordered vs ordered-mk2: the A/A validations found a systematic difference of -1.19% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=65536 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -5.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=65536 churn: ordered vs btree-map: the pooled difference of 0.22% does not clear the 0.48% noise floor, the bound on what the harness reports between identical code in every process
- single-value url n=65536 churn: ordered vs btree-map: the pooled interval [-2.61%, 3.05%] includes zero
- single-value url n=65536 churn: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 churn: ordered vs btree-map: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value url n=65536 build: ordered vs baseline: the processes scatter 6.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 build: ordered vs btree-map: the processes scatter 6.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 build: ordered vs ordered-mk1: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 build: ordered vs ordered-mk2: the processes scatter 8.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=65536 build: ordered vs ordered-mk4: the processes scatter 7.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
