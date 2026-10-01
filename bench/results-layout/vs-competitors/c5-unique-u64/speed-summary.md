| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique | email | 4096 | valuesFor | ordered | btree-map | 6 | 22.5 | 96.4 | 4.31× [4.26, 4.35] | +76.8% | [+76.5%, +77.0%] | 0.2 pts | 1.1 | yes | yes |
| unique | email | 4096 | valuesBetween | ordered | btree-map | 6 | 1251 | 432 | 0.35× [0.34, 0.35] | -188.9% | [-191.5%, -186.2%] | 2.5 pts | 1.4 | yes | yes |
| unique | email | 4096 | prefix | ordered | btree-map | 6 | 57.1 | 97.8 | 1.71× [1.70, 1.73] | +41.6% | [+41.0%, +42.3%] | 0.6 pts | 1.6 | yes | yes |
| unique | email | 4096 | churn | ordered | btree-map | 6 | 71.2 | 143 | 1.99× [1.94, 2.04] | +49.7% | [+48.4%, +51.0%] | 1.2 pts | 2.3 | yes | yes |
| unique | email | 4096 | build | ordered | btree-map | 6 | 1.04 ms | 1.94 ms | 1.89× [1.88, 1.90] | +47.0% | [+46.7%, +47.3%] | 0.3 pts | 0.3 | yes | yes |
| unique | email | 16384 | valuesFor | ordered | btree-map | 6 | 38.0 | 132 | 3.47× [3.44, 3.51] | +71.2% | [+70.9%, +71.5%] | 0.3 pts | 1.6 | yes | yes |
| unique | email | 16384 | valuesBetween | ordered | btree-map | 6 | 1580 | 511 | 0.32× [0.32, 0.33] | -208.5% | [-211.8%, -205.2%] | 3.1 pts | 2.3 | yes | yes |
| unique | email | 16384 | prefix | ordered | btree-map | 6 | 71.7 | 133 | 1.86× [1.84, 1.88] | +46.3% | [+45.8%, +46.7%] | 0.4 pts | 1.4 | yes | yes |
| unique | email | 16384 | churn | ordered | btree-map | 6 | 88.7 | 198 | 2.24× [2.21, 2.27] | +55.3% | [+54.7%, +55.9%] | 0.6 pts | 1.3 | yes | yes |
| unique | email | 16384 | build | ordered | btree-map | 6 | 5.05 ms | 10.42 ms | 2.06× [2.03, 2.08] | +51.4% | [+50.9%, +52.0%] | 0.5 pts | 1.4 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | btree-map | 12 | 173 | 249 | 1.46× [1.42, 1.50] | +31.3% | [+29.3%, +33.2%] | 2.9 pts | 4.1 | yes | yes |
| unique | email | 262144 | valuesBetween | ordered | btree-map | 12 | 3169 | 1587 | 0.52× [0.49, 0.55] | -93.1% | [-105.9%, -80.2%] | 15.4 pts | 1.7 | no | yes |
| unique | email | 262144 | prefix | ordered | btree-map | 12 | 236 | 273 | 1.19× [1.15, 1.24] | +16.1% | [+13.0%, +19.2%] | 4.2 pts | 4.1 | no | yes |
| unique | email | 262144 | churn | ordered | btree-map | 12 | 363 | 558 | 1.54× [1.52, 1.56] | +35.2% | [+34.4%, +36.1%] | 1.0 pts | 1.3 | yes | yes |
| unique | path | 4096 | valuesFor | ordered | btree-map | 6 | 78.2 | 113 | 1.45× [1.42, 1.48] | +31.1% | [+29.6%, +32.7%] | 1.4 pts | 1.1 | yes | yes |
| unique | path | 4096 | valuesBetween | ordered | btree-map | 6 | 1958 | 550 | 0.28× [0.27, 0.28] | -258.0% | [-263.9%, -252.0%] | 5.7 pts | 1.1 | yes | yes |
| unique | path | 4096 | prefix | ordered | btree-map | 6 | 230 | 141 | 0.61× [0.60, 0.63] | -63.5% | [-67.7%, -59.2%] | 4.1 pts | 1.3 | yes | yes |
| unique | path | 4096 | churn | ordered | btree-map | 6 | 193 | 178 | 0.92× [0.92, 0.93] | -8.6% | [-9.3%, -8.0%] | 0.6 pts | 0.4 | yes | yes |
| unique | path | 4096 | build | ordered | btree-map | 6 | 2.58 ms | 2.50 ms | 0.97× [0.96, 0.97] | -3.4% | [-3.9%, -2.8%] | 0.5 pts | 0.6 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | btree-map | 6 | 112 | 162 | 1.45× [1.43, 1.47] | +31.0% | [+30.1%, +32.0%] | 0.9 pts | 1.1 | yes | yes |
| unique | path | 16384 | valuesBetween | ordered | btree-map | 6 | 2511 | 706 | 0.28× [0.28, 0.29] | -254.8% | [-261.1%, -248.4%] | 6.1 pts | 1.8 | yes | yes |
| unique | path | 16384 | prefix | ordered | btree-map | 6 | 470 | 259 | 0.55× [0.54, 0.56] | -81.4% | [-85.5%, -77.3%] | 3.9 pts | 1.0 | yes | yes |
| unique | path | 16384 | churn | ordered | btree-map | 6 | 247 | 243 | 0.98× [0.96, 0.99] | -2.3% | [-3.9%, -0.7%] | 1.6 pts | 1.5 | yes | yes |
| unique | path | 16384 | build | ordered | btree-map | 6 | 12.63 ms | 12.92 ms | 1.02× [1.02, 1.03] | +2.3% | [+1.7%, +2.9%] | 0.6 pts | 0.8 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | btree-map | 12 | 383 | 403 | 1.06× [1.04, 1.08] | +6.0% | [+4.1%, +7.8%] | 4.5 pts | 3.4 | yes | yes |
| unique | path | 262144 | valuesBetween | ordered | btree-map | 12 | 5741 | 2935 | 0.51× [0.49, 0.53] | -96.5% | [-102.5%, -90.4%] | 6.6 pts | 2.2 | yes | yes |
| unique | path | 262144 | prefix | ordered | btree-map | 12 | 5315 | 2415 | 0.46× [0.44, 0.47] | -119.4% | [-125.1%, -113.7%] | 10.4 pts | 0.5 | yes | yes |
| unique | path | 262144 | churn | ordered | btree-map | 12 | 716 | 750 | 1.03× [1.02, 1.05] | +3.1% | [+1.8%, +4.4%] | 1.6 pts | 1.9 | yes | yes |
| unique | str | 4096 | valuesFor | ordered | btree-map | 8 | 34.0 | 96.5 | 2.85× [2.80, 2.91] | +65.0% | [+64.2%, +65.7%] | 0.8 pts | 1.0 | yes | yes |
| unique | str | 4096 | valuesBetween | ordered | btree-map | 8 | 1567 | 427 | 0.27× [0.27, 0.27] | -266.9% | [-269.5%, -264.2%] | 2.5 pts | 1.1 | yes | yes |
| unique | str | 4096 | prefix | ordered | btree-map | 8 | 3094 | 824 | 0.27× [0.27, 0.27] | -273.3% | [-276.1%, -270.6%] | 2.6 pts | 1.7 | yes | yes |
| unique | str | 4096 | churn | ordered | btree-map | 8 | 94.2 | 139 | 1.48× [1.46, 1.49] | +32.3% | [+31.7%, +32.9%] | 0.7 pts | 0.9 | yes | yes |
| unique | str | 4096 | build | ordered | btree-map | 8 | 1.34 ms | 1.94 ms | 1.44× [1.38, 1.50] | +30.4% | [+27.6%, +33.2%] | 3.2 pts | 2.8 | yes | yes |
| unique | str | 16384 | valuesFor | ordered | btree-map | 6 | 48.9 | 132 | 2.70× [2.66, 2.73] | +62.9% | [+62.4%, +63.4%] | 0.5 pts | 2.2 | yes | yes |
| unique | str | 16384 | valuesBetween | ordered | btree-map | 6 | 1615 | 505 | 0.31× [0.31, 0.32] | -219.0% | [-220.8%, -217.1%] | 1.7 pts | 0.9 | yes | yes |
| unique | str | 16384 | prefix | ordered | btree-map | 6 | 13.1 µs | 3467 | 0.27× [0.26, 0.27] | -276.6% | [-281.4%, -271.7%] | 4.6 pts | 1.9 | yes | yes |
| unique | str | 16384 | churn | ordered | btree-map | 6 | 114 | 198 | 1.74× [1.72, 1.75] | +42.5% | [+41.9%, +43.0%] | 0.5 pts | 0.8 | yes | yes |
| unique | str | 16384 | build | ordered | btree-map | 6 | 6.29 ms | 10.26 ms | 1.64× [1.63, 1.65] | +39.0% | [+38.7%, +39.3%] | 0.3 pts | 0.9 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | btree-map | 12 | 181 | 247 | 1.36× [1.32, 1.40] | +26.3% | [+24.1%, +28.5%] | 3.2 pts | 3.0 | yes | yes |
| unique | str | 262144 | valuesBetween | ordered | btree-map | 12 | 2814 | 1414 | 0.50× [0.50, 0.51] | -98.4% | [-100.9%, -95.9%] | 7.6 pts | 0.8 | yes | yes |
| unique | str | 262144 | prefix | ordered | btree-map | 12 | 390.1 µs | 160.8 µs | 0.44× [0.40, 0.48] | -129.0% | [-149.5%, -108.5%] | 20.7 pts | 0.4 | no | yes |
| unique | str | 262144 | churn | ordered | btree-map | 12 | 400 | 560 | 1.44× [1.43, 1.45] | +30.7% | [+30.3%, +31.1%] | 1.0 pts | 0.9 | yes | yes |
| unique | street | 4096 | valuesFor | ordered | btree-map | 6 | 40.0 | 87.9 | 2.19× [2.13, 2.25] | +54.4% | [+53.1%, +55.6%] | 1.2 pts | 1.8 | yes | yes |
| unique | street | 4096 | valuesBetween | ordered | btree-map | 6 | 1732 | 489 | 0.28× [0.28, 0.29] | -255.3% | [-262.4%, -248.3%] | 6.7 pts | 2.2 | yes | yes |
| unique | street | 4096 | prefix | ordered | btree-map | 6 | 195 | 131 | 0.67× [0.66, 0.68] | -50.2% | [-52.3%, -48.1%] | 2.0 pts | 1.4 | yes | yes |
| unique | street | 4096 | churn | ordered | btree-map | 6 | 114 | 138 | 1.22× [1.19, 1.24] | +17.8% | [+16.2%, +19.3%] | 1.5 pts | 1.4 | yes | yes |
| unique | street | 4096 | build | ordered | btree-map | 6 | 1.59 ms | 2.10 ms | 1.32× [1.29, 1.36] | +24.4% | [+22.4%, +26.4%] | 1.9 pts | 3.4 | yes | yes |
| unique | street | 16384 | valuesFor | ordered | btree-map | 6 | 66.4 | 121 | 1.82× [1.81, 1.84] | +45.2% | [+44.8%, +45.5%] | 0.3 pts | 0.7 | yes | yes |
| unique | street | 16384 | valuesBetween | ordered | btree-map | 6 | 2138 | 566 | 0.26× [0.26, 0.27] | -278.5% | [-281.2%, -275.8%] | 2.6 pts | 0.8 | yes | yes |
| unique | street | 16384 | prefix | ordered | btree-map | 6 | 616 | 268 | 0.43× [0.43, 0.44] | -130.4% | [-133.0%, -127.9%] | 2.4 pts | 0.8 | yes | yes |
| unique | street | 16384 | churn | ordered | btree-map | 6 | 151 | 192 | 1.28× [1.25, 1.30] | +21.6% | [+20.3%, +23.0%] | 1.3 pts | 1.9 | yes | yes |
| unique | street | 16384 | build | ordered | btree-map | 6 | 8.00 ms | 10.91 ms | 1.35× [1.33, 1.38] | +26.2% | [+25.0%, +27.3%] | 1.1 pts | 1.2 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 13.4 | 83.6 | 6.23× [6.14, 6.32] | +83.9% | [+83.7%, +84.2%] | 0.2 pts | 1.7 | yes | yes |
| unique | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 1021 | 422 | 0.41× [0.41, 0.42] | -142.6% | [-144.7%, -140.4%] | 2.0 pts | 1.5 | yes | yes |
| unique | u64 | 4096 | churn | ordered | btree-map | 6 | 46.6 | 121 | 2.61× [2.57, 2.64] | +61.7% | [+61.2%, +62.2%] | 0.5 pts | 1.5 | yes | yes |
| unique | u64 | 4096 | build | ordered | btree-map | 6 | 758.5 µs | 1.70 ms | 2.25× [2.21, 2.28] | +55.5% | [+54.7%, +56.2%] | 0.7 pts | 1.6 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | btree-map | 6 | 23.3 | 112 | 4.79× [4.76, 4.82] | +79.1% | [+79.0%, +79.3%] | 0.1 pts | 1.8 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | btree-map | 6 | 1517 | 450 | 0.30× [0.30, 0.30] | -236.7% | [-238.3%, -235.0%] | 1.6 pts | 1.0 | yes | yes |
| unique | u64 | 16384 | churn | ordered | btree-map | 6 | 55.1 | 166 | 3.00× [2.91, 3.09] | +66.6% | [+65.6%, +67.7%] | 1.0 pts | 2.0 | yes | yes |
| unique | u64 | 16384 | build | ordered | btree-map | 6 | 3.52 ms | 8.90 ms | 2.51× [2.49, 2.54] | +60.2% | [+59.9%, +60.6%] | 0.3 pts | 1.2 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | btree-map | 6 | 67.6 | 210 | 3.06× [2.84, 3.32] | +67.4% | [+64.8%, +69.9%] | 2.4 pts | 3.0 | yes | yes |
| unique | u64 | 262144 | valuesBetween | ordered | btree-map | 6 | 1640 | 806 | 0.49× [0.48, 0.51] | -102.3% | [-109.5%, -95.2%] | 6.8 pts | 1.4 | yes | yes |
| unique | u64 | 262144 | churn | ordered | btree-map | 6 | 269 | 386 | 1.48× [1.43, 1.54] | +32.6% | [+30.1%, +35.0%] | 2.3 pts | 2.0 | yes | yes |
| unique | url | 4096 | valuesFor | ordered | btree-map | 12 | 60.9 | 112 | 1.85× [1.81, 1.90] | +46.0% | [+44.8%, +47.2%] | 1.4 pts | 1.9 | yes | yes |
| unique | url | 4096 | valuesBetween | ordered | btree-map | 12 | 1859 | 543 | 0.29× [0.28, 0.30] | -245.8% | [-253.6%, -238.1%] | 7.5 pts | 1.7 | yes | yes |
| unique | url | 4096 | prefix | ordered | btree-map | 12 | 144 | 126 | 0.87× [0.85, 0.89] | -15.0% | [-17.7%, -12.4%] | 2.8 pts | 1.7 | no | yes |
| unique | url | 4096 | churn | ordered | btree-map | 12 | 156 | 177 | 1.14× [1.13, 1.15] | +12.3% | [+11.4%, +13.1%] | 1.6 pts | 1.7 | yes | yes |
| unique | url | 4096 | build | ordered | btree-map | 12 | 2.16 ms | 2.40 ms | 1.12× [1.11, 1.14] | +10.9% | [+9.8%, +12.0%] | 1.2 pts | 1.0 | yes | yes |
| unique | url | 16384 | valuesFor | ordered | btree-map | 6 | 89.7 | 162 | 1.82× [1.79, 1.85] | +45.0% | [+44.0%, +45.9%] | 0.9 pts | 2.4 | yes | yes |
| unique | url | 16384 | valuesBetween | ordered | btree-map | 6 | 2357 | 708 | 0.30× [0.30, 0.30] | -233.0% | [-237.4%, -228.5%] | 4.3 pts | 1.7 | yes | yes |
| unique | url | 16384 | prefix | ordered | btree-map | 6 | 215 | 193 | 0.90× [0.89, 0.91] | -11.6% | [-12.9%, -10.3%] | 1.2 pts | 1.2 | yes | yes |
| unique | url | 16384 | churn | ordered | btree-map | 6 | 199 | 241 | 1.21× [1.19, 1.23] | +17.4% | [+16.0%, +18.7%] | 1.3 pts | 1.3 | yes | yes |
| unique | url | 16384 | build | ordered | btree-map | 6 | 10.30 ms | 12.66 ms | 1.23× [1.22, 1.24] | +18.6% | [+17.8%, +19.4%] | 0.7 pts | 1.0 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | btree-map | 12 | 357 | 405 | 1.13× [1.11, 1.16] | +11.9% | [+10.0%, +13.7%] | 3.2 pts | 2.9 | yes | yes |
| unique | url | 262144 | valuesBetween | ordered | btree-map | 12 | 6072 | 2846 | 0.47× [0.46, 0.47] | -113.3% | [-115.7%, -110.9%] | 4.3 pts | 1.4 | yes | yes |
| unique | url | 262144 | prefix | ordered | btree-map | 12 | 1321 | 818 | 0.62× [0.61, 0.64] | -60.1% | [-63.7%, -56.4%] | 4.8 pts | 2.0 | yes | yes |
| unique | url | 262144 | churn | ordered | btree-map | 12 | 650 | 741 | 1.13× [1.12, 1.15] | +11.5% | [+10.3%, +12.7%] | 1.6 pts | 1.7 | yes | yes |
| unique | uuid | 4096 | valuesFor | ordered | btree-map | 6 | 28.0 | 96.3 | 3.42× [3.36, 3.49] | +70.8% | [+70.2%, +71.3%] | 0.5 pts | 1.3 | yes | yes |
| unique | uuid | 4096 | valuesBetween | ordered | btree-map | 6 | 1422 | 426 | 0.30× [0.30, 0.31] | -231.5% | [-235.3%, -227.6%] | 3.7 pts | 2.5 | yes | yes |
| unique | uuid | 4096 | prefix | ordered | btree-map | 6 | 68.8 | 99.3 | 1.44× [1.39, 1.49] | +30.6% | [+28.2%, +33.0%] | 2.3 pts | 4.8 | yes | yes |
| unique | uuid | 4096 | churn | ordered | btree-map | 6 | 81.3 | 138 | 1.69× [1.67, 1.71] | +40.9% | [+40.1%, +41.6%] | 0.7 pts | 1.7 | yes | yes |
| unique | uuid | 4096 | build | ordered | btree-map | 6 | 1.15 ms | 1.90 ms | 1.67× [1.64, 1.69] | +40.0% | [+39.1%, +40.9%] | 0.8 pts | 1.1 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | btree-map | 6 | 42.5 | 134 | 3.15× [3.12, 3.19] | +68.3% | [+68.0%, +68.6%] | 0.3 pts | 2.2 | yes | yes |
| unique | uuid | 16384 | valuesBetween | ordered | btree-map | 6 | 1574 | 528 | 0.34× [0.33, 0.34] | -196.6% | [-198.5%, -194.6%] | 1.9 pts | 1.2 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | btree-map | 6 | 83.7 | 142 | 1.68× [1.66, 1.70] | +40.6% | [+39.8%, +41.3%] | 0.7 pts | 2.7 | yes | yes |
| unique | uuid | 16384 | churn | ordered | btree-map | 6 | 100 | 191 | 1.91× [1.86, 1.98] | +47.8% | [+46.1%, +49.4%] | 1.6 pts | 1.9 | yes | yes |
| unique | uuid | 16384 | build | ordered | btree-map | 6 | 5.46 ms | 10.32 ms | 1.90× [1.88, 1.91] | +47.3% | [+46.9%, +47.6%] | 0.4 pts | 0.8 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | btree-map | 12 | 201 | 287 | 1.43× [1.41, 1.45] | +30.1% | [+29.2%, +31.1%] | 1.7 pts | 1.8 | yes | yes |
| unique | uuid | 262144 | valuesBetween | ordered | btree-map | 12 | 3860 | 2084 | 0.55× [0.53, 0.56] | -83.5% | [-87.8%, -79.2%] | 5.2 pts | 1.0 | yes | yes |
| unique | uuid | 262144 | prefix | ordered | btree-map | 12 | 364 | 407 | 1.12× [1.09, 1.14] | +10.5% | [+8.6%, +12.3%] | 2.9 pts | 2.5 | yes | yes |
| unique | uuid | 262144 | churn | ordered | btree-map | 12 | 369 | 571 | 1.58× [1.56, 1.61] | +36.8% | [+35.9%, +37.8%] | 1.0 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique email n=4096 churn: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.52% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=262144 valuesFor: ordered vs btree-map: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.96% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=262144 prefix: ordered vs btree-map: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.93% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique path n=262144 valuesFor: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=4096 build: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique str n=262144 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -9.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.90% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=4096 build: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.44% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=16384 churn: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 churn: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.54% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 prefix: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.19% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique uuid n=4096 prefix: ordered vs btree-map: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 prefix: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 prefix: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
