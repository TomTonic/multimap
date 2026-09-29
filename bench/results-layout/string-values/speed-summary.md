| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi-str | email | 4096 | valuesFor | ordered | btree-sets | 6 | 52.9 | 178 | 3.34× [3.31, 3.37] | +70.1% | [+69.8%, +70.4%] | 0.3 pts | 1.0 | yes |
| multi-str | email | 4096 | valuesFor | ordered | hashed | 6 | 53.0 | 42.7 | 0.81× [0.80, 0.81] | -23.9% | [-25.0%, -22.8%] | 1.1 pts | 0.7 | yes |
| multi-str | email | 4096 | valuesFor | ordered | map-sets | 6 | 52.7 | 95.3 | 1.81× [1.80, 1.83] | +44.9% | [+44.5%, +45.2%] | 0.3 pts | 0.8 | yes |
| multi-str | email | 4096 | valuesBetween | ordered | btree-sets | 6 | 3271 | 7353 | 2.25× [2.22, 2.28] | +55.5% | [+54.9%, +56.2%] | 0.6 pts | 1.0 | yes |
| multi-str | email | 4096 | valuesBetween | ordered | hashed | 6 | 3412 | 55.3 µs | 16.25× [16.12, 16.37] | +93.8% | [+93.8%, +93.9%] | 0.0 pts | 0.6 | yes |
| multi-str | email | 4096 | valuesBetween | ordered | map-sets | 6 | 3307 | 56.3 µs | 17.04× [16.85, 17.23] | +94.1% | [+94.1%, +94.2%] | 0.1 pts | 0.8 | yes |
| multi-str | email | 4096 | prefix | ordered | btree-sets | 6 | 79.3 | 178 | 2.25× [2.23, 2.26] | +55.5% | [+55.2%, +55.7%] | 0.2 pts | 0.9 | yes |
| multi-str | email | 4096 | prefix | ordered | hashed | 6 | 90.5 | 48.7 µs | 523.72× [479.81, 576.48] | +99.8% | [+99.8%, +99.8%] | 0.0 pts | 0.5 | yes |
| multi-str | email | 4096 | prefix | ordered | map-sets | 6 | 86.4 | 45.1 µs | 508.84× [480.62, 540.59] | +99.8% | [+99.8%, +99.8%] | 0.0 pts | 0.3 | yes |
| multi-str | email | 4096 | churn | ordered | btree-sets | 6 | 85.2 | 199 | 2.35× [2.32, 2.38] | +57.4% | [+57.0%, +57.9%] | 0.5 pts | 0.6 | yes |
| multi-str | email | 4096 | churn | ordered | hashed | 6 | 81.6 | 67.2 | 0.81× [0.81, 0.82] | -22.8% | [-24.2%, -21.5%] | 1.3 pts | 0.8 | yes |
| multi-str | email | 4096 | churn | ordered | map-sets | 8 | 82.8 | 72.5 | 0.87× [0.86, 0.88] | -15.1% | [-16.7%, -13.6%] | 1.8 pts | 0.6 | yes |
| multi-str | email | 4096 | build | ordered | btree-sets | 6 | 6.95 ms | 14.98 ms | 2.16× [2.14, 2.18] | +53.7% | [+53.2%, +54.2%] | 0.5 pts | 1.2 | yes |
| multi-str | email | 4096 | build | ordered | hashed | 6 | 6.84 ms | 5.87 ms | 0.86× [0.85, 0.86] | -16.5% | [-17.3%, -15.7%] | 0.8 pts | 0.7 | yes |
| multi-str | email | 4096 | build | ordered | map-sets | 6 | 6.88 ms | 5.92 ms | 0.86× [0.85, 0.87] | -16.2% | [-17.1%, -15.3%] | 0.9 pts | 0.5 | yes |
| multi-str | email | 16384 | valuesFor | ordered | btree-sets | 6 | 66.5 | 227 | 3.41× [3.37, 3.45] | +70.7% | [+70.3%, +71.0%] | 0.3 pts | 0.9 | yes |
| multi-str | email | 16384 | valuesFor | ordered | hashed | 6 | 65.6 | 48.4 | 0.74× [0.73, 0.75] | -35.0% | [-36.6%, -33.5%] | 1.5 pts | 0.8 | yes |
| multi-str | email | 16384 | valuesFor | ordered | map-sets | 6 | 66.6 | 109 | 1.64× [1.61, 1.68] | +39.2% | [+37.8%, +40.6%] | 1.4 pts | 1.4 | yes |
| multi-str | email | 16384 | valuesBetween | ordered | btree-sets | 6 | 3890 | 8214 | 2.14× [2.08, 2.19] | +53.2% | [+52.0%, +54.4%] | 1.1 pts | 1.8 | yes |
| multi-str | email | 16384 | valuesBetween | ordered | hashed | 6 | 4675 | 227.8 µs | 48.70× [46.30, 51.37] | +97.9% | [+97.8%, +98.1%] | 0.1 pts | 1.1 | yes |
| multi-str | email | 16384 | valuesBetween | ordered | map-sets | 6 | 5049 | 221.7 µs | 42.91× [39.87, 46.44] | +97.7% | [+97.5%, +97.8%] | 0.2 pts | 1.3 | yes |
| multi-str | email | 16384 | prefix | ordered | btree-sets | 6 | 101 | 241 | 2.38× [2.35, 2.41] | +58.0% | [+57.5%, +58.6%] | 0.5 pts | 0.7 | yes |
| multi-str | email | 16384 | prefix | ordered | hashed | 6 | 304 | 211.9 µs | 700.50× [672.61, 730.82] | +99.9% | [+99.9%, +99.9%] | 0.0 pts | 0.5 | yes |
| multi-str | email | 16384 | prefix | ordered | map-sets | 6 | 290 | 198.4 µs | 680.89× [662.34, 700.52] | +99.9% | [+99.8%, +99.9%] | 0.0 pts | 0.3 | yes |
| multi-str | email | 16384 | churn | ordered | btree-sets | 6 | 141 | 314 | 2.22× [2.19, 2.26] | +55.0% | [+54.3%, +55.8%] | 0.7 pts | 1.0 | yes |
| multi-str | email | 16384 | churn | ordered | hashed | 12 | 124 | 93.9 | 0.77× [0.76, 0.78] | -30.0% | [-31.4%, -28.6%] | 2.2 pts | 1.2 | yes |
| multi-str | email | 16384 | churn | ordered | map-sets | 12 | 132 | 124 | 0.92× [0.91, 0.94] | -8.2% | [-10.0%, -6.3%] | 2.9 pts | 1.6 | yes |
| multi-str | email | 16384 | build | ordered | btree-sets | 6 | 41.46 ms | 91.47 ms | 2.21× [2.17, 2.25] | +54.8% | [+54.0%, +55.6%] | 0.8 pts | 1.2 | yes |
| multi-str | email | 16384 | build | ordered | hashed | 6 | 41.64 ms | 33.79 ms | 0.81× [0.80, 0.82] | -23.5% | [-24.3%, -22.7%] | 0.8 pts | 0.6 | yes |
| multi-str | email | 16384 | build | ordered | map-sets | 12 | 40.84 ms | 39.06 ms | 0.95× [0.93, 0.96] | -5.8% | [-7.5%, -4.0%] | 2.8 pts | 1.4 | yes |
| multi-str | email | 262144 | valuesFor | ordered | btree-sets | 6 | 329 | 733 | 2.24× [2.17, 2.31] | +55.3% | [+54.0%, +56.7%] | 1.3 pts | 2.2 | yes |
| multi-str | email | 262144 | valuesFor | ordered | hashed | 6 | 279 | 193 | 0.69× [0.67, 0.70] | -45.6% | [-48.5%, -42.7%] | 2.8 pts | 1.5 | yes |
| multi-str | email | 262144 | valuesFor | ordered | map-sets | 6 | 280 | 430 | 1.52× [1.48, 1.56] | +34.2% | [+32.5%, +35.9%] | 1.6 pts | 2.7 | yes |
| multi-str | email | 262144 | valuesBetween | ordered | btree-sets | 6 | 11.3 µs | 28.0 µs | 2.49× [2.44, 2.55] | +59.9% | [+58.9%, +60.8%] | 0.9 pts | 1.8 | yes |
| multi-str | email | 262144 | prefix | ordered | btree-sets | 6 | 437 | 935 | 2.13× [2.08, 2.18] | +53.0% | [+52.0%, +54.0%] | 1.0 pts | 1.7 | yes |
| multi-str | email | 262144 | churn | ordered | btree-sets | 6 | 602 | 1022 | 1.72× [1.67, 1.76] | +41.8% | [+40.2%, +43.3%] | 1.4 pts | 1.9 | yes |
| multi-str | email | 262144 | churn | ordered | hashed | 102 | 453 | 398 | 0.87× [0.86, 0.88] | -14.8% | [-16.4%, -13.1%] | 8.4 pts | 2.9 | yes |
| multi-str | email | 262144 | churn | ordered | map-sets | 30 | 515 | 464 | 0.90× [0.89, 0.90] | -11.5% | [-12.4%, -10.7%] | 2.3 pts | 1.4 | yes |
| multi-str | email | 1048576 | valuesFor | ordered | btree-sets | 6 | 538 | 1310 | 2.44× [2.36, 2.52] | +59.0% | [+57.7%, +60.3%] | 1.3 pts | 1.7 | yes |
| multi-str | email | 1048576 | valuesFor | ordered | hashed | 20 | 404 | 220 | 0.54× [0.53, 0.55] | -85.8% | [-89.5%, -82.1%] | 7.9 pts | 3.9 | yes |
| multi-str | email | 1048576 | valuesFor | ordered | map-sets | 20 | 424 | 504 | 1.16× [1.14, 1.18] | +13.9% | [+12.1%, +15.6%] | 3.7 pts | 4.4 | yes |
| multi-str | email | 1048576 | valuesBetween | ordered | btree-sets | 6 | 13.7 µs | 33.7 µs | 2.45× [2.37, 2.54] | +59.2% | [+57.8%, +60.6%] | 1.4 pts | 2.7 | yes |
| multi-str | email | 1048576 | prefix | ordered | btree-sets | 6 | 807 | 2072 | 2.55× [2.47, 2.64] | +60.8% | [+59.6%, +62.1%] | 1.2 pts | 3.0 | yes |
| multi-str | email | 1048576 | churn | ordered | btree-sets | 6 | 786 | 1597 | 2.04× [1.99, 2.09] | +51.0% | [+49.9%, +52.1%] | 1.1 pts | 1.8 | yes |
| multi-str | email | 1048576 | churn | ordered | hashed | 20 | 644 | 506 | 0.78× [0.76, 0.79] | -28.7% | [-31.5%, -25.9%] | 6.0 pts | 2.6 | yes |
| multi-str | email | 1048576 | churn | ordered | map-sets | 20 | 694 | 563 | 0.81× [0.80, 0.83] | -23.4% | [-25.7%, -21.1%] | 5.0 pts | 2.9 | yes |
| multi-str | path | 4096 | valuesFor | ordered | btree-sets | 6 | 115 | 205 | 1.78× [1.76, 1.80] | +43.8% | [+43.1%, +44.5%] | 0.7 pts | 1.1 | yes |
| multi-str | path | 4096 | valuesFor | ordered | hashed | 6 | 114 | 50.1 | 0.44× [0.44, 0.44] | -127.4% | [-128.6%, -126.3%] | 1.1 pts | 0.5 | yes |
| multi-str | path | 4096 | valuesFor | ordered | map-sets | 6 | 113 | 101 | 0.89× [0.88, 0.89] | -12.8% | [-13.6%, -12.0%] | 0.7 pts | 0.6 | yes |
| multi-str | path | 4096 | valuesBetween | ordered | btree-sets | 6 | 4739 | 7817 | 1.65× [1.64, 1.66] | +39.3% | [+39.0%, +39.7%] | 0.3 pts | 0.4 | yes |
| multi-str | path | 4096 | valuesBetween | ordered | hashed | 6 | 4904 | 64.1 µs | 13.03× [12.94, 13.13] | +92.3% | [+92.3%, +92.4%] | 0.1 pts | 0.4 | yes |
| multi-str | path | 4096 | valuesBetween | ordered | map-sets | 6 | 4850 | 65.3 µs | 13.48× [13.38, 13.59] | +92.6% | [+92.5%, +92.6%] | 0.1 pts | 0.5 | yes |
| multi-str | path | 4096 | prefix | ordered | btree-sets | 6 | 358 | 481 | 1.34× [1.31, 1.38] | +25.6% | [+23.8%, +27.4%] | 1.7 pts | 0.7 | yes |
| multi-str | path | 4096 | prefix | ordered | hashed | 6 | 487 | 58.7 µs | 118.58× [112.57, 125.27] | +99.2% | [+99.1%, +99.2%] | 0.0 pts | 1.4 | yes |
| multi-str | path | 4096 | prefix | ordered | map-sets | 6 | 467 | 54.9 µs | 116.66× [111.96, 121.77] | +99.1% | [+99.1%, +99.2%] | 0.0 pts | 0.9 | yes |
| multi-str | path | 4096 | churn | ordered | btree-sets | 6 | 188 | 253 | 1.35× [1.34, 1.36] | +26.0% | [+25.5%, +26.6%] | 0.5 pts | 0.3 | yes |
| multi-str | path | 4096 | churn | ordered | hashed | 6 | 175 | 80.5 | 0.46× [0.45, 0.46] | -118.1% | [-120.9%, -115.3%] | 2.7 pts | 1.1 | yes |
| multi-str | path | 4096 | churn | ordered | map-sets | 6 | 176 | 90.1 | 0.52× [0.50, 0.54] | -93.6% | [-100.6%, -86.5%] | 6.7 pts | 1.5 | yes |
| multi-str | path | 4096 | build | ordered | btree-sets | 6 | 13.60 ms | 18.44 ms | 1.36× [1.35, 1.36] | +26.3% | [+26.1%, +26.5%] | 0.2 pts | 0.3 | yes |
| multi-str | path | 4096 | build | ordered | hashed | 6 | 13.54 ms | 6.70 ms | 0.50× [0.49, 0.50] | -101.7% | [-102.9%, -100.4%] | 1.2 pts | 0.6 | yes |
| multi-str | path | 4096 | build | ordered | map-sets | 6 | 13.71 ms | 7.24 ms | 0.52× [0.51, 0.53] | -91.3% | [-94.8%, -87.8%] | 3.3 pts | 1.2 | yes |
| multi-str | path | 16384 | valuesFor | ordered | btree-sets | 6 | 156 | 275 | 1.76× [1.72, 1.79] | +43.0% | [+41.9%, +44.1%] | 1.0 pts | 0.8 | yes |
| multi-str | path | 16384 | valuesFor | ordered | hashed | 6 | 146 | 54.8 | 0.38× [0.37, 0.39] | -162.8% | [-170.8%, -154.9%] | 7.6 pts | 2.5 | yes |
| multi-str | path | 16384 | valuesFor | ordered | map-sets | 6 | 152 | 118 | 0.78× [0.77, 0.80] | -27.4% | [-29.9%, -25.0%] | 2.3 pts | 1.1 | yes |
| multi-str | path | 16384 | valuesBetween | ordered | btree-sets | 6 | 5572 | 8882 | 1.58× [1.56, 1.61] | +36.9% | [+35.9%, +37.8%] | 0.9 pts | 0.9 | yes |
| multi-str | path | 16384 | valuesBetween | ordered | hashed | 6 | 6822 | 271.4 µs | 38.29× [34.66, 42.77] | +97.4% | [+97.1%, +97.7%] | 0.3 pts | 2.1 | yes |
| multi-str | path | 16384 | valuesBetween | ordered | map-sets | 6 | 6904 | 265.9 µs | 36.85× [32.82, 42.02] | +97.3% | [+97.0%, +97.6%] | 0.3 pts | 2.3 | yes |
| multi-str | path | 16384 | prefix | ordered | btree-sets | 6 | 745 | 1065 | 1.47× [1.42, 1.52] | +31.9% | [+29.7%, +34.1%] | 2.1 pts | 0.5 | yes |
| multi-str | path | 16384 | prefix | ordered | hashed | 6 | 1456 | 264.8 µs | 182.29× [173.76, 191.70] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 0.8 | yes |
| multi-str | path | 16384 | prefix | ordered | map-sets | 6 | 1415 | 251.5 µs | 178.20× [172.87, 183.87] | +99.4% | [+99.4%, +99.5%] | 0.0 pts | 0.6 | yes |
| multi-str | path | 16384 | churn | ordered | btree-sets | 6 | 305 | 421 | 1.37× [1.34, 1.39] | +26.9% | [+25.6%, +28.1%] | 1.2 pts | 1.2 | yes |
| multi-str | path | 16384 | churn | ordered | hashed | 6 | 289 | 138 | 0.48× [0.46, 0.50] | -109.6% | [-117.3%, -102.0%] | 7.3 pts | 1.2 | yes |
| multi-str | path | 16384 | churn | ordered | map-sets | 6 | 267 | 147 | 0.55× [0.53, 0.57] | -82.1% | [-87.5%, -76.8%] | 5.1 pts | 1.3 | yes |
| multi-str | path | 16384 | build | ordered | btree-sets | 6 | 85.27 ms | 114.99 ms | 1.34× [1.33, 1.36] | +25.6% | [+24.6%, +26.7%] | 1.0 pts | 1.1 | yes |
| multi-str | path | 16384 | build | ordered | hashed | 6 | 84.66 ms | 42.05 ms | 0.49× [0.48, 0.50] | -104.2% | [-107.4%, -100.9%] | 3.1 pts | 1.2 | yes |
| multi-str | path | 16384 | build | ordered | map-sets | 6 | 84.30 ms | 46.19 ms | 0.55× [0.53, 0.57] | -81.3% | [-87.1%, -75.5%] | 5.5 pts | 1.9 | yes |
| multi-str | path | 262144 | valuesFor | ordered | btree-sets | 6 | 658 | 996 | 1.51× [1.47, 1.55] | +33.6% | [+31.8%, +35.4%] | 1.7 pts | 2.0 | yes |
| multi-str | path | 262144 | valuesFor | ordered | hashed | 6 | 533 | 228 | 0.42× [0.41, 0.43] | -138.0% | [-141.9%, -134.0%] | 3.8 pts | 1.6 | yes |
| multi-str | path | 262144 | valuesFor | ordered | map-sets | 6 | 571 | 458 | 0.81× [0.79, 0.82] | -24.1% | [-26.5%, -21.8%] | 2.2 pts | 1.5 | yes |
| multi-str | path | 262144 | valuesBetween | ordered | btree-sets | 6 | 15.3 µs | 29.1 µs | 1.89× [1.86, 1.91] | +47.0% | [+46.3%, +47.7%] | 0.7 pts | 0.9 | yes |
| multi-str | path | 262144 | prefix | ordered | btree-sets | 6 | 16.8 µs | 32.5 µs | 1.96× [1.93, 1.98] | +48.9% | [+48.3%, +49.6%] | 0.6 pts | 0.1 | yes |
| multi-str | path | 262144 | churn | ordered | btree-sets | 6 | 1052 | 1279 | 1.22× [1.21, 1.24] | +18.1% | [+17.3%, +19.0%] | 0.8 pts | 1.0 | yes |
| multi-str | path | 262144 | churn | ordered | hashed | 6 | 796 | 463 | 0.58× [0.56, 0.59] | -72.8% | [-77.0%, -68.5%] | 4.0 pts | 1.3 | yes |
| multi-str | path | 262144 | churn | ordered | map-sets | 6 | 913 | 547 | 0.60× [0.58, 0.62] | -66.7% | [-72.4%, -60.9%] | 5.5 pts | 2.6 | yes |
| multi-str | str | 4096 | valuesFor | ordered | btree-sets | 6 | 64.6 | 178 | 2.75× [2.74, 2.76] | +63.6% | [+63.6%, +63.7%] | 0.1 pts | 0.4 | yes |
| multi-str | str | 4096 | valuesFor | ordered | hashed | 6 | 63.8 | 43.0 | 0.67× [0.67, 0.68] | -48.7% | [-49.8%, -47.6%] | 1.1 pts | 0.7 | yes |
| multi-str | str | 4096 | valuesFor | ordered | map-sets | 6 | 63.7 | 95.5 | 1.49× [1.47, 1.51] | +32.8% | [+31.8%, +33.8%] | 1.0 pts | 1.7 | yes |
| multi-str | str | 4096 | valuesBetween | ordered | btree-sets | 6 | 3836 | 7463 | 1.95× [1.93, 1.96] | +48.6% | [+48.2%, +49.0%] | 0.4 pts | 0.7 | yes |
| multi-str | str | 4096 | valuesBetween | ordered | hashed | 6 | 3964 | 54.9 µs | 13.74× [13.61, 13.87] | +92.7% | [+92.7%, +92.8%] | 0.1 pts | 0.5 | yes |
| multi-str | str | 4096 | valuesBetween | ordered | map-sets | 6 | 3897 | 56.0 µs | 14.41× [14.21, 14.62] | +93.1% | [+93.0%, +93.2%] | 0.1 pts | 0.9 | yes |
| multi-str | str | 4096 | prefix | ordered | btree-sets | 6 | 8787 | 18.6 µs | 2.12× [2.10, 2.13] | +52.7% | [+52.4%, +53.1%] | 0.4 pts | 0.5 | yes |
| multi-str | str | 4096 | prefix | ordered | hashed | 6 | 9046 | 56.5 µs | 6.25× [6.17, 6.33] | +84.0% | [+83.8%, +84.2%] | 0.2 pts | 0.7 | yes |
| multi-str | str | 4096 | prefix | ordered | map-sets | 6 | 8825 | 65.4 µs | 7.39× [7.21, 7.58] | +86.5% | [+86.1%, +86.8%] | 0.3 pts | 1.5 | yes |
| multi-str | str | 4096 | churn | ordered | btree-sets | 6 | 100 | 198 | 1.99× [1.97, 2.00] | +49.6% | [+49.2%, +50.1%] | 0.4 pts | 0.5 | yes |
| multi-str | str | 4096 | churn | ordered | hashed | 6 | 97.2 | 66.1 | 0.69× [0.67, 0.70] | -45.9% | [-48.9%, -43.0%] | 2.8 pts | 1.5 | yes |
| multi-str | str | 4096 | churn | ordered | map-sets | 6 | 97.8 | 72.2 | 0.74× [0.73, 0.75] | -35.2% | [-37.5%, -32.9%] | 2.2 pts | 0.9 | yes |
| multi-str | str | 4096 | build | ordered | btree-sets | 6 | 8.09 ms | 14.99 ms | 1.85× [1.84, 1.86] | +46.0% | [+45.6%, +46.3%] | 0.3 pts | 0.8 | yes |
| multi-str | str | 4096 | build | ordered | hashed | 6 | 7.98 ms | 5.87 ms | 0.74× [0.73, 0.74] | -35.9% | [-36.8%, -35.1%] | 0.8 pts | 0.7 | yes |
| multi-str | str | 4096 | build | ordered | map-sets | 6 | 8.06 ms | 5.95 ms | 0.74× [0.72, 0.75] | -36.0% | [-38.9%, -33.1%] | 2.8 pts | 1.2 | yes |
| multi-str | str | 16384 | valuesFor | ordered | btree-sets | 6 | 78.4 | 229 | 2.93× [2.89, 2.96] | +65.8% | [+65.4%, +66.2%] | 0.4 pts | 0.9 | yes |
| multi-str | str | 16384 | valuesFor | ordered | hashed | 6 | 75.6 | 47.7 | 0.63× [0.63, 0.64] | -58.2% | [-59.1%, -57.3%] | 0.9 pts | 0.6 | yes |
| multi-str | str | 16384 | valuesFor | ordered | map-sets | 6 | 77.5 | 109 | 1.43× [1.40, 1.46] | +29.9% | [+28.3%, +31.5%] | 1.5 pts | 1.7 | yes |
| multi-str | str | 16384 | valuesBetween | ordered | btree-sets | 6 | 4072 | 8285 | 2.03× [2.00, 2.06] | +50.8% | [+50.1%, +51.5%] | 0.7 pts | 1.2 | yes |
| multi-str | str | 16384 | valuesBetween | ordered | hashed | 6 | 4646 | 224.1 µs | 46.87× [43.10, 51.36] | +97.9% | [+97.7%, +98.1%] | 0.2 pts | 1.6 | yes |
| multi-str | str | 16384 | valuesBetween | ordered | map-sets | 6 | 4721 | 218.1 µs | 45.19× [43.00, 47.61] | +97.8% | [+97.7%, +97.9%] | 0.1 pts | 2.2 | yes |
| multi-str | str | 16384 | prefix | ordered | btree-sets | 6 | 38.4 µs | 82.9 µs | 2.16× [2.13, 2.18] | +53.6% | [+53.1%, +54.1%] | 0.5 pts | 0.6 | yes |
| multi-str | str | 16384 | prefix | ordered | hashed | 6 | 39.2 µs | 252.6 µs | 6.41× [6.34, 6.50] | +84.4% | [+84.2%, +84.6%] | 0.2 pts | 0.9 | yes |
| multi-str | str | 16384 | prefix | ordered | map-sets | 6 | 39.1 µs | 296.1 µs | 7.56× [7.37, 7.75] | +86.8% | [+86.4%, +87.1%] | 0.3 pts | 1.4 | yes |
| multi-str | str | 16384 | churn | ordered | btree-sets | 6 | 150 | 306 | 2.03× [2.01, 2.05] | +50.6% | [+50.2%, +51.1%] | 0.5 pts | 0.7 | yes |
| multi-str | str | 16384 | churn | ordered | hashed | 6 | 144 | 99.1 | 0.69× [0.68, 0.71] | -43.9% | [-47.9%, -39.9%] | 3.8 pts | 0.9 | yes |
| multi-str | str | 16384 | churn | ordered | map-sets | 18 | 148 | 121 | 0.82× [0.81, 0.84] | -21.6% | [-23.7%, -19.6%] | 4.1 pts | 1.9 | yes |
| multi-str | str | 16384 | build | ordered | btree-sets | 6 | 44.68 ms | 88.69 ms | 1.98× [1.95, 2.01] | +49.5% | [+48.7%, +50.3%] | 0.7 pts | 1.4 | yes |
| multi-str | str | 16384 | build | ordered | hashed | 6 | 44.84 ms | 32.83 ms | 0.73× [0.72, 0.74] | -37.1% | [-38.5%, -35.6%] | 1.4 pts | 1.0 | yes |
| multi-str | str | 16384 | build | ordered | map-sets | 18 | 45.24 ms | 38.74 ms | 0.86× [0.85, 0.87] | -16.6% | [-17.8%, -15.4%] | 2.3 pts | 1.3 | yes |
| multi-str | str | 262144 | valuesFor | ordered | btree-sets | 6 | 356 | 696 | 1.97× [1.94, 1.99] | +49.2% | [+48.5%, +49.8%] | 0.6 pts | 1.3 | yes |
| multi-str | str | 262144 | valuesFor | ordered | hashed | 6 | 304 | 190 | 0.62× [0.61, 0.64] | -60.5% | [-63.6%, -57.4%] | 2.9 pts | 1.7 | yes |
| multi-str | str | 262144 | valuesFor | ordered | map-sets | 6 | 322 | 430 | 1.31× [1.28, 1.34] | +23.6% | [+22.0%, +25.3%] | 1.6 pts | 1.6 | yes |
| multi-str | str | 262144 | valuesBetween | ordered | btree-sets | 6 | 12.2 µs | 28.1 µs | 2.31× [2.26, 2.37] | +56.8% | [+55.8%, +57.8%] | 0.9 pts | 1.8 | yes |
| multi-str | str | 262144 | prefix | ordered | btree-sets | 6 | 1.90 ms | 4.60 ms | 2.41× [2.37, 2.45] | +58.5% | [+57.9%, +59.2%] | 0.6 pts | 1.4 | yes |
| multi-str | str | 262144 | churn | ordered | btree-sets | 6 | 601 | 992 | 1.67× [1.62, 1.72] | +40.1% | [+38.3%, +41.8%] | 1.7 pts | 2.5 | yes |
| multi-str | str | 262144 | churn | ordered | hashed | 110 | 487 | 405 | 0.84× [0.82, 0.85] | -19.6% | [-21.6%, -17.7%] | 10.4 pts | 4.9 | yes |
| multi-str | str | 262144 | churn | ordered | map-sets | 30 | 534 | 466 | 0.87× [0.86, 0.88] | -14.9% | [-16.0%, -13.7%] | 3.0 pts | 2.0 | yes |
| multi-str | str | 1048576 | valuesFor | ordered | btree-sets | 6 | 504 | 1230 | 2.42× [2.35, 2.50] | +58.7% | [+57.5%, +60.0%] | 1.2 pts | 2.9 | yes |
| multi-str | str | 1048576 | valuesFor | ordered | hashed | 6 | 423 | 219 | 0.52× [0.51, 0.52] | -94.1% | [-96.9%, -91.2%] | 2.7 pts | 1.5 | yes |
| multi-str | str | 1048576 | valuesFor | ordered | map-sets | 6 | 441 | 498 | 1.13× [1.11, 1.15] | +11.2% | [+9.6%, +12.7%] | 1.4 pts | 1.5 | yes |
| multi-str | str | 1048576 | valuesBetween | ordered | btree-sets | 6 | 13.6 µs | 32.8 µs | 2.40× [2.33, 2.48] | +58.4% | [+57.1%, +59.7%] | 1.2 pts | 3.1 | yes |
| multi-str | str | 1048576 | prefix | ordered | btree-sets | 6 | 8.49 ms | 20.54 ms | 2.45× [2.39, 2.53] | +59.2% | [+58.1%, +60.4%] | 1.1 pts | 3.5 | yes |
| multi-str | str | 1048576 | churn | ordered | btree-sets | 6 | 794 | 1570 | 1.97× [1.95, 1.99] | +49.3% | [+48.7%, +49.9%] | 0.6 pts | 1.4 | yes |
| multi-str | str | 1048576 | churn | ordered | hashed | 30 | 658 | 508 | 0.77× [0.74, 0.80] | -30.7% | [-35.8%, -25.6%] | 13.7 pts | 6.1 | no |
| multi-str | str | 1048576 | churn | ordered | map-sets | 30 | 717 | 565 | 0.79× [0.78, 0.80] | -26.4% | [-27.6%, -25.1%] | 3.4 pts | 2.4 | yes |
| multi-str | street | 4096 | valuesFor | ordered | btree-sets | 6 | 59.2 | 145 | 2.45× [2.42, 2.48] | +59.2% | [+58.7%, +59.7%] | 0.5 pts | 1.0 | yes |
| multi-str | street | 4096 | valuesFor | ordered | hashed | 6 | 56.9 | 22.4 | 0.39× [0.39, 0.40] | -154.8% | [-159.7%, -150.0%] | 4.6 pts | 1.0 | yes |
| multi-str | street | 4096 | valuesFor | ordered | map-sets | 6 | 58.1 | 70.1 | 1.21× [1.20, 1.22] | +17.1% | [+16.3%, +17.9%] | 0.7 pts | 0.7 | yes |
| multi-str | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 2731 | 4754 | 1.75× [1.71, 1.79] | +42.8% | [+41.5%, +44.1%] | 1.2 pts | 2.0 | yes |
| multi-str | street | 4096 | valuesBetween | ordered | hashed | 6 | 2823 | 54.6 µs | 19.30× [19.04, 19.58] | +94.8% | [+94.7%, +94.9%] | 0.1 pts | 0.9 | yes |
| multi-str | street | 4096 | valuesBetween | ordered | map-sets | 6 | 2753 | 55.0 µs | 20.03× [19.85, 20.22] | +95.0% | [+95.0%, +95.1%] | 0.0 pts | 0.5 | yes |
| multi-str | street | 4096 | prefix | ordered | btree-sets | 6 | 292 | 600 | 2.05× [2.00, 2.09] | +51.1% | [+50.0%, +52.2%] | 1.1 pts | 1.3 | yes |
| multi-str | street | 4096 | prefix | ordered | hashed | 6 | 332 | 51.3 µs | 154.28× [151.08, 157.63] | +99.4% | [+99.3%, +99.4%] | 0.0 pts | 0.5 | yes |
| multi-str | street | 4096 | prefix | ordered | map-sets | 6 | 326 | 48.0 µs | 147.64× [142.48, 153.18] | +99.3% | [+99.3%, +99.3%] | 0.0 pts | 0.8 | yes |
| multi-str | street | 4096 | churn | ordered | btree-sets | 6 | 109 | 206 | 1.88× [1.86, 1.90] | +46.8% | [+46.2%, +47.4%] | 0.6 pts | 0.8 | yes |
| multi-str | street | 4096 | churn | ordered | hashed | 6 | 105 | 54.6 | 0.52× [0.50, 0.54] | -93.0% | [-99.9%, -86.1%] | 6.6 pts | 2.4 | yes |
| multi-str | street | 4096 | churn | ordered | map-sets | 6 | 106 | 72.7 | 0.68× [0.65, 0.70] | -48.1% | [-52.9%, -43.3%] | 4.6 pts | 1.5 | yes |
| multi-str | street | 4096 | build | ordered | btree-sets | 6 | 3.36 ms | 6.58 ms | 1.95× [1.94, 1.96] | +48.7% | [+48.4%, +49.0%] | 0.3 pts | 0.4 | yes |
| multi-str | street | 4096 | build | ordered | hashed | 6 | 3.36 ms | 2.28 ms | 0.68× [0.67, 0.69] | -47.0% | [-48.7%, -45.3%] | 1.6 pts | 0.9 | yes |
| multi-str | street | 4096 | build | ordered | map-sets | 6 | 3.36 ms | 2.73 ms | 0.81× [0.80, 0.83] | -22.8% | [-24.9%, -20.6%] | 2.0 pts | 0.7 | yes |
| multi-str | street | 16384 | valuesFor | ordered | btree-sets | 6 | 82.0 | 189 | 2.32× [2.27, 2.37] | +56.9% | [+55.9%, +57.9%] | 1.0 pts | 1.5 | yes |
| multi-str | street | 16384 | valuesFor | ordered | hashed | 6 | 79.3 | 29.7 | 0.37× [0.37, 0.38] | -167.1% | [-170.4%, -163.8%] | 3.1 pts | 1.3 | yes |
| multi-str | street | 16384 | valuesFor | ordered | map-sets | 6 | 80.6 | 82.9 | 1.03× [1.01, 1.05] | +3.1% | [+1.2%, +4.9%] | 1.8 pts | 1.2 | yes |
| multi-str | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 3219 | 5552 | 1.73× [1.70, 1.77] | +42.3% | [+41.2%, +43.4%] | 1.0 pts | 1.4 | yes |
| multi-str | street | 16384 | valuesBetween | ordered | hashed | 6 | 3826 | 228.3 µs | 59.35× [55.37, 63.95] | +98.3% | [+98.2%, +98.4%] | 0.1 pts | 1.2 | yes |
| multi-str | street | 16384 | valuesBetween | ordered | map-sets | 6 | 4002 | 226.9 µs | 55.28× [52.16, 58.80] | +98.2% | [+98.1%, +98.3%] | 0.1 pts | 0.9 | yes |
| multi-str | street | 16384 | prefix | ordered | btree-sets | 6 | 1089 | 2255 | 2.05× [2.03, 2.08] | +51.3% | [+50.7%, +51.9%] | 0.6 pts | 0.3 | yes |
| multi-str | street | 16384 | prefix | ordered | hashed | 6 | 1579 | 229.1 µs | 147.06× [142.44, 151.99] | +99.3% | [+99.3%, +99.3%] | 0.0 pts | 0.6 | yes |
| multi-str | street | 16384 | prefix | ordered | map-sets | 6 | 1526 | 220.9 µs | 143.16× [139.75, 146.74] | +99.3% | [+99.3%, +99.3%] | 0.0 pts | 0.4 | yes |
| multi-str | street | 16384 | churn | ordered | btree-sets | 6 | 162 | 296 | 1.80× [1.74, 1.85] | +44.3% | [+42.7%, +46.0%] | 1.6 pts | 1.7 | yes |
| multi-str | street | 16384 | churn | ordered | hashed | 6 | 149 | 71.4 | 0.48× [0.47, 0.49] | -107.4% | [-111.6%, -103.3%] | 4.0 pts | 1.3 | yes |
| multi-str | street | 16384 | churn | ordered | map-sets | 16 | 152 | 104 | 0.69× [0.67, 0.71] | -45.4% | [-49.4%, -41.3%] | 7.6 pts | 2.7 | yes |
| multi-str | street | 16384 | build | ordered | btree-sets | 6 | 20.21 ms | 39.97 ms | 1.98× [1.93, 2.02] | +49.4% | [+48.3%, +50.5%] | 1.1 pts | 1.2 | yes |
| multi-str | street | 16384 | build | ordered | hashed | 6 | 19.45 ms | 12.89 ms | 0.66× [0.65, 0.67] | -50.7% | [-53.0%, -48.3%] | 2.2 pts | 1.2 | yes |
| multi-str | street | 16384 | build | ordered | map-sets | 16 | 19.28 ms | 16.05 ms | 0.84× [0.83, 0.84] | -19.7% | [-20.6%, -18.8%] | 1.7 pts | 0.7 | yes |
| multi-str | u64 | 4096 | valuesFor | ordered | btree-sets | 6 | 40.1 | 161 | 4.03× [3.98, 4.08] | +75.2% | [+74.9%, +75.5%] | 0.3 pts | 1.0 | yes |
| multi-str | u64 | 4096 | valuesFor | ordered | hashed | 6 | 40.1 | 39.6 | 0.99× [0.98, 1.00] | -1.2% | [-2.6%, +0.2%] | 1.3 pts | 1.0 | yes |
| multi-str | u64 | 4096 | valuesFor | ordered | map-sets | 6 | 39.8 | 93.4 | 2.35× [2.33, 2.37] | +57.5% | [+57.2%, +57.9%] | 0.3 pts | 0.9 | yes |
| multi-str | u64 | 4096 | valuesBetween | ordered | btree-sets | 6 | 2885 | 7255 | 2.52× [2.48, 2.55] | +60.2% | [+59.7%, +60.8%] | 0.5 pts | 0.8 | yes |
| multi-str | u64 | 4096 | valuesBetween | ordered | hashed | 6 | 3044 | 52.3 µs | 17.26× [17.03, 17.50] | +94.2% | [+94.1%, +94.3%] | 0.1 pts | 1.0 | yes |
| multi-str | u64 | 4096 | valuesBetween | ordered | map-sets | 6 | 2907 | 53.7 µs | 18.59× [18.45, 18.73] | +94.6% | [+94.6%, +94.7%] | 0.0 pts | 0.4 | yes |
| multi-str | u64 | 4096 | churn | ordered | btree-sets | 6 | 59.9 | 169 | 2.82× [2.80, 2.83] | +64.5% | [+64.3%, +64.7%] | 0.2 pts | 0.4 | yes |
| multi-str | u64 | 4096 | churn | ordered | hashed | 6 | 59.3 | 59.5 | 1.01× [0.99, 1.02] | +0.8% | [-0.5%, +2.1%] | 1.3 pts | 0.9 | yes |
| multi-str | u64 | 4096 | churn | ordered | map-sets | 8 | 59.2 | 63.1 | 1.06× [1.04, 1.08] | +5.7% | [+4.1%, +7.2%] | 1.9 pts | 1.1 | yes |
| multi-str | u64 | 4096 | build | ordered | btree-sets | 6 | 5.45 ms | 13.19 ms | 2.41× [2.39, 2.43] | +58.5% | [+58.2%, +58.8%] | 0.3 pts | 0.7 | yes |
| multi-str | u64 | 4096 | build | ordered | hashed | 6 | 5.42 ms | 5.48 ms | 1.01× [1.00, 1.03] | +1.5% | [+0.3%, +2.6%] | 1.1 pts | 1.2 | yes |
| multi-str | u64 | 4096 | build | ordered | map-sets | 6 | 5.52 ms | 5.51 ms | 0.99× [0.98, 1.01] | -0.7% | [-2.4%, +1.1%] | 1.7 pts | 1.0 | yes |
| multi-str | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 49.7 | 201 | 4.03× [4.01, 4.06] | +75.2% | [+75.1%, +75.3%] | 0.1 pts | 0.5 | yes |
| multi-str | u64 | 16384 | valuesFor | ordered | hashed | 6 | 48.7 | 45.9 | 0.94× [0.93, 0.94] | -6.5% | [-7.0%, -5.9%] | 0.5 pts | 0.6 | yes |
| multi-str | u64 | 16384 | valuesFor | ordered | map-sets | 6 | 48.5 | 106 | 2.17× [2.13, 2.21] | +53.8% | [+53.0%, +54.7%] | 0.8 pts | 1.5 | yes |
| multi-str | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 3794 | 7947 | 2.10× [2.09, 2.11] | +52.4% | [+52.0%, +52.7%] | 0.3 pts | 0.5 | yes |
| multi-str | u64 | 16384 | valuesBetween | ordered | hashed | 6 | 4295 | 206.6 µs | 47.36× [45.04, 49.93] | +97.9% | [+97.8%, +98.0%] | 0.1 pts | 1.7 | yes |
| multi-str | u64 | 16384 | valuesBetween | ordered | map-sets | 6 | 4533 | 200.2 µs | 44.29× [41.13, 47.98] | +97.7% | [+97.6%, +97.9%] | 0.2 pts | 1.8 | yes |
| multi-str | u64 | 16384 | churn | ordered | btree-sets | 6 | 93.8 | 261 | 2.80× [2.71, 2.90] | +64.3% | [+63.1%, +65.6%] | 1.2 pts | 1.9 | yes |
| multi-str | u64 | 16384 | churn | ordered | hashed | 10 | 86.1 | 86.3 | 1.00× [0.98, 1.02] | +0.0% | [-1.8%, +1.8%] | 2.5 pts | 1.2 | yes |
| multi-str | u64 | 16384 | churn | ordered | map-sets | 6 | 87.4 | 99.0 | 1.13× [1.11, 1.14] | +11.2% | [+9.8%, +12.6%] | 1.3 pts | 1.0 | yes |
| multi-str | u64 | 16384 | build | ordered | btree-sets | 6 | 30.19 ms | 78.14 ms | 2.56× [2.51, 2.61] | +60.9% | [+60.1%, +61.7%] | 0.7 pts | 1.3 | yes |
| multi-str | u64 | 16384 | build | ordered | hashed | 6 | 30.79 ms | 29.00 ms | 0.95× [0.94, 0.96] | -5.2% | [-6.5%, -3.9%] | 1.2 pts | 1.5 | yes |
| multi-str | u64 | 16384 | build | ordered | map-sets | 10 | 31.56 ms | 35.70 ms | 1.12× [1.11, 1.13] | +10.9% | [+10.1%, +11.7%] | 1.1 pts | 0.6 | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | btree-sets | 6 | 184 | 598 | 3.30× [3.20, 3.40] | +69.7% | [+68.7%, +70.6%] | 0.9 pts | 2.6 | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | hashed | 6 | 167 | 175 | 1.05× [1.04, 1.06] | +5.0% | [+4.2%, +5.8%] | 0.7 pts | 0.7 | yes |
| multi-str | u64 | 262144 | valuesFor | ordered | map-sets | 6 | 171 | 403 | 2.34× [2.31, 2.38] | +57.3% | [+56.7%, +57.9%] | 0.6 pts | 1.1 | yes |
| multi-str | u64 | 262144 | valuesBetween | ordered | btree-sets | 6 | 9350 | 25.6 µs | 2.73× [2.66, 2.80] | +63.3% | [+62.4%, +64.2%] | 0.9 pts | 2.1 | yes |
| multi-str | u64 | 262144 | churn | ordered | btree-sets | 6 | 443 | 809 | 1.84× [1.77, 1.92] | +45.7% | [+43.5%, +47.9%] | 2.1 pts | 3.4 | yes |
| multi-str | u64 | 262144 | churn | ordered | hashed | 6 | 370 | 352 | 0.96× [0.95, 0.98] | -3.9% | [-5.8%, -1.9%] | 1.9 pts | 0.4 | yes |
| multi-str | u64 | 262144 | churn | ordered | map-sets | 6 | 392 | 391 | 0.99× [0.98, 1.01] | -0.7% | [-2.5%, +1.1%] | 1.7 pts | 1.4 | yes |
| multi-str | u64 | 1048576 | valuesFor | ordered | btree-sets | 6 | 242 | 1006 | 4.13× [4.07, 4.20] | +75.8% | [+75.4%, +76.2%] | 0.4 pts | 1.3 | yes |
| multi-str | u64 | 1048576 | valuesFor | ordered | hashed | 6 | 210 | 211 | 1.00× [0.99, 1.02] | +0.3% | [-1.1%, +1.6%] | 1.3 pts | 1.3 | yes |
| multi-str | u64 | 1048576 | valuesFor | ordered | map-sets | 6 | 218 | 484 | 2.22× [2.14, 2.31] | +55.0% | [+53.4%, +56.7%] | 1.6 pts | 3.1 | yes |
| multi-str | u64 | 1048576 | valuesBetween | ordered | btree-sets | 6 | 8975 | 30.3 µs | 3.33× [3.24, 3.43] | +70.0% | [+69.1%, +70.8%] | 0.8 pts | 2.9 | yes |
| multi-str | u64 | 1048576 | churn | ordered | btree-sets | 6 | 543 | 1427 | 2.59× [2.51, 2.68] | +61.4% | [+60.1%, +62.7%] | 1.2 pts | 4.0 | yes |
| multi-str | u64 | 1048576 | churn | ordered | hashed | 60 | 477 | 448 | 0.94× [0.93, 0.96] | -6.2% | [-7.9%, -4.6%] | 6.5 pts | 1.2 | yes |
| multi-str | u64 | 1048576 | churn | ordered | map-sets | 30 | 505 | 493 | 0.97× [0.96, 0.98] | -2.8% | [-4.1%, -1.6%] | 3.4 pts | 2.2 | yes |
| multi-str | url | 4096 | valuesFor | ordered | btree-sets | 6 | 97.1 | 199 | 2.04× [2.01, 2.06] | +50.9% | [+50.3%, +51.5%] | 0.6 pts | 1.6 | yes |
| multi-str | url | 4096 | valuesFor | ordered | hashed | 6 | 95.9 | 49.4 | 0.51× [0.51, 0.52] | -95.3% | [-98.0%, -92.5%] | 2.6 pts | 1.7 | yes |
| multi-str | url | 4096 | valuesFor | ordered | map-sets | 6 | 95.8 | 100 | 1.04× [1.02, 1.06] | +3.8% | [+2.1%, +5.5%] | 1.6 pts | 1.7 | yes |
| multi-str | url | 4096 | valuesBetween | ordered | btree-sets | 6 | 4570 | 7673 | 1.68× [1.65, 1.70] | +40.3% | [+39.5%, +41.2%] | 0.8 pts | 1.1 | yes |
| multi-str | url | 4096 | valuesBetween | ordered | hashed | 6 | 4730 | 62.2 µs | 13.10× [12.94, 13.26] | +92.4% | [+92.3%, +92.5%] | 0.1 pts | 0.9 | yes |
| multi-str | url | 4096 | valuesBetween | ordered | map-sets | 6 | 4645 | 63.7 µs | 13.80× [13.66, 13.95] | +92.8% | [+92.7%, +92.8%] | 0.1 pts | 0.7 | yes |
| multi-str | url | 4096 | prefix | ordered | btree-sets | 6 | 197 | 262 | 1.33× [1.32, 1.34] | +24.8% | [+24.2%, +25.4%] | 0.6 pts | 0.6 | yes |
| multi-str | url | 4096 | prefix | ordered | hashed | 6 | 294 | 55.2 µs | 181.97× [170.37, 195.27] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 1.2 | yes |
| multi-str | url | 4096 | prefix | ordered | map-sets | 6 | 319 | 51.5 µs | 162.66× [156.76, 169.02] | +99.4% | [+99.4%, +99.4%] | 0.0 pts | 0.8 | yes |
| multi-str | url | 4096 | churn | ordered | btree-sets | 6 | 157 | 246 | 1.56× [1.55, 1.57] | +36.0% | [+35.5%, +36.4%] | 0.4 pts | 0.3 | yes |
| multi-str | url | 4096 | churn | ordered | hashed | 6 | 149 | 80.9 | 0.54× [0.54, 0.55] | -84.2% | [-86.9%, -81.5%] | 2.6 pts | 0.8 | yes |
| multi-str | url | 4096 | churn | ordered | map-sets | 6 | 148 | 87.1 | 0.59× [0.58, 0.60] | -69.3% | [-72.2%, -66.5%] | 2.7 pts | 0.8 | yes |
| multi-str | url | 4096 | build | ordered | btree-sets | 6 | 11.42 ms | 17.89 ms | 1.57× [1.56, 1.58] | +36.1% | [+35.7%, +36.5%] | 0.4 pts | 0.6 | yes |
| multi-str | url | 4096 | build | ordered | hashed | 6 | 11.33 ms | 6.70 ms | 0.59× [0.58, 0.60] | -69.0% | [-71.2%, -66.9%] | 2.1 pts | 1.4 | yes |
| multi-str | url | 4096 | build | ordered | map-sets | 6 | 11.40 ms | 7.00 ms | 0.62× [0.61, 0.63] | -62.6% | [-65.3%, -59.9%] | 2.5 pts | 1.2 | yes |
| multi-str | url | 16384 | valuesFor | ordered | btree-sets | 6 | 133 | 269 | 2.03× [1.97, 2.09] | +50.7% | [+49.2%, +52.1%] | 1.4 pts | 1.5 | yes |
| multi-str | url | 16384 | valuesFor | ordered | hashed | 6 | 126 | 55.5 | 0.44× [0.44, 0.45] | -125.6% | [-128.6%, -122.6%] | 2.8 pts | 0.8 | yes |
| multi-str | url | 16384 | valuesFor | ordered | map-sets | 18 | 136 | 121 | 0.91× [0.89, 0.92] | -10.3% | [-12.3%, -8.4%] | 3.9 pts | 1.3 | yes |
| multi-str | url | 16384 | valuesBetween | ordered | btree-sets | 6 | 5451 | 8739 | 1.62× [1.58, 1.66] | +38.2% | [+36.8%, +39.6%] | 1.3 pts | 1.2 | yes |
| multi-str | url | 16384 | valuesBetween | ordered | hashed | 6 | 6302 | 264.0 µs | 41.87× [41.48, 42.27] | +97.6% | [+97.6%, +97.6%] | 0.0 pts | 0.2 | yes |
| multi-str | url | 16384 | valuesBetween | ordered | map-sets | 6 | 6982 | 259.8 µs | 36.95× [34.71, 39.50] | +97.3% | [+97.1%, +97.5%] | 0.2 pts | 1.6 | yes |
| multi-str | url | 16384 | prefix | ordered | btree-sets | 6 | 321 | 461 | 1.44× [1.42, 1.45] | +30.3% | [+29.7%, +31.0%] | 0.6 pts | 0.3 | yes |
| multi-str | url | 16384 | prefix | ordered | hashed | 6 | 741 | 254.0 µs | 337.89× [318.70, 359.54] | +99.7% | [+99.7%, +99.7%] | 0.0 pts | 0.7 | yes |
| multi-str | url | 16384 | prefix | ordered | map-sets | 6 | 741 | 239.8 µs | 320.37× [291.92, 354.97] | +99.7% | [+99.7%, +99.7%] | 0.0 pts | 1.2 | yes |
| multi-str | url | 16384 | churn | ordered | btree-sets | 6 | 252 | 398 | 1.55× [1.48, 1.62] | +35.4% | [+32.5%, +38.4%] | 2.8 pts | 2.3 | yes |
| multi-str | url | 16384 | churn | ordered | hashed | 6 | 254 | 137 | 0.54× [0.53, 0.57] | -83.6% | [-90.1%, -77.0%] | 6.3 pts | 2.4 | yes |
| multi-str | url | 16384 | churn | ordered | map-sets | 18 | 234 | 142 | 0.62× [0.61, 0.63] | -60.4% | [-63.3%, -57.5%] | 5.8 pts | 2.4 | yes |
| multi-str | url | 16384 | build | ordered | btree-sets | 6 | 71.56 ms | 109.60 ms | 1.54× [1.53, 1.55] | +34.9% | [+34.5%, +35.4%] | 0.4 pts | 0.5 | yes |
| multi-str | url | 16384 | build | ordered | hashed | 6 | 71.98 ms | 40.55 ms | 0.57× [0.56, 0.57] | -76.7% | [-79.1%, -74.2%] | 2.3 pts | 0.9 | yes |
| multi-str | url | 16384 | build | ordered | map-sets | 6 | 71.43 ms | 46.18 ms | 0.64× [0.63, 0.65] | -56.2% | [-58.7%, -53.6%] | 2.4 pts | 0.6 | yes |
| multi-str | url | 262144 | valuesFor | ordered | btree-sets | 6 | 586 | 903 | 1.56× [1.53, 1.59] | +35.8% | [+34.5%, +37.1%] | 1.2 pts | 2.3 | yes |
| multi-str | url | 262144 | valuesFor | ordered | hashed | 6 | 497 | 223 | 0.45× [0.43, 0.47] | -123.6% | [-134.0%, -113.2%] | 9.9 pts | 4.0 | yes |
| multi-str | url | 262144 | valuesFor | ordered | map-sets | 6 | 525 | 454 | 0.87× [0.86, 0.89] | -14.3% | [-16.1%, -12.5%] | 1.7 pts | 1.1 | yes |
| multi-str | url | 262144 | valuesBetween | ordered | btree-sets | 6 | 14.8 µs | 29.2 µs | 1.96× [1.91, 2.02] | +49.0% | [+47.6%, +50.5%] | 1.4 pts | 1.9 | yes |
| multi-str | url | 262144 | prefix | ordered | btree-sets | 6 | 3255 | 6007 | 1.84× [1.79, 1.90] | +45.7% | [+44.0%, +47.4%] | 1.6 pts | 0.4 | yes |
| multi-str | url | 262144 | churn | ordered | btree-sets | 12 | 908 | 1169 | 1.29× [1.27, 1.31] | +22.6% | [+21.4%, +23.8%] | 1.9 pts | 2.2 | yes |
| multi-str | url | 262144 | churn | ordered | hashed | 12 | 750 | 470 | 0.61× [0.59, 0.63] | -63.1% | [-68.6%, -57.6%] | 8.6 pts | 3.6 | yes |
| multi-str | url | 262144 | churn | ordered | map-sets | 12 | 791 | 535 | 0.67× [0.66, 0.68] | -48.5% | [-51.0%, -46.1%] | 3.9 pts | 2.6 | yes |
| multi-str | url | 1048576 | valuesFor | ordered | btree-sets | 6 | 844 | 1385 | 1.60× [1.52, 1.69] | +37.5% | [+34.1%, +40.9%] | 3.2 pts | 4.3 | yes |
| multi-str | url | 1048576 | valuesFor | ordered | hashed | 6 | 715 | 248 | 0.35× [0.33, 0.36] | -188.6% | [-198.6%, -178.6%] | 9.5 pts | 3.9 | yes |
| multi-str | url | 1048576 | valuesFor | ordered | map-sets | 6 | 753 | 522 | 0.70× [0.68, 0.71] | -43.8% | [-46.1%, -41.5%] | 2.2 pts | 1.1 | yes |
| multi-str | url | 1048576 | valuesBetween | ordered | btree-sets | 6 | 16.8 µs | 33.9 µs | 1.99× [1.94, 2.03] | +49.7% | [+48.6%, +50.8%] | 1.1 pts | 1.7 | yes |
| multi-str | url | 1048576 | prefix | ordered | btree-sets | 6 | 15.2 µs | 30.9 µs | 2.05× [2.00, 2.10] | +51.2% | [+50.1%, +52.3%] | 1.1 pts | 0.2 | yes |
| multi-str | url | 1048576 | churn | ordered | btree-sets | 6 | 1293 | 1792 | 1.39× [1.35, 1.42] | +27.9% | [+26.0%, +29.7%] | 1.7 pts | 2.1 | yes |
| multi-str | url | 1048576 | churn | ordered | hashed | 12 | 1113 | 566 | 0.51× [0.48, 0.53] | -96.7% | [-106.2%, -87.3%] | 14.9 pts | 4.2 | yes |
| multi-str | url | 1048576 | churn | ordered | map-sets | 6 | 1286 | 697 | 0.54× [0.52, 0.57] | -83.9% | [-92.0%, -75.9%] | 7.7 pts | 3.6 | yes |
| multi-str | uuid | 4096 | valuesFor | ordered | btree-sets | 6 | 58.2 | 179 | 3.06× [3.04, 3.08] | +67.3% | [+67.1%, +67.5%] | 0.2 pts | 1.0 | yes |
| multi-str | uuid | 4096 | valuesFor | ordered | hashed | 6 | 57.3 | 43.4 | 0.76× [0.75, 0.77] | -32.1% | [-33.5%, -30.7%] | 1.4 pts | 0.9 | yes |
| multi-str | uuid | 4096 | valuesFor | ordered | map-sets | 6 | 57.7 | 96.2 | 1.67× [1.65, 1.70] | +40.3% | [+39.3%, +41.2%] | 0.9 pts | 1.8 | yes |
| multi-str | uuid | 4096 | valuesBetween | ordered | btree-sets | 6 | 3552 | 7434 | 2.10× [2.08, 2.11] | +52.3% | [+51.8%, +52.7%] | 0.4 pts | 0.7 | yes |
| multi-str | uuid | 4096 | valuesBetween | ordered | hashed | 6 | 3741 | 54.8 µs | 14.89× [14.33, 15.50] | +93.3% | [+93.0%, +93.5%] | 0.3 pts | 2.2 | yes |
| multi-str | uuid | 4096 | valuesBetween | ordered | map-sets | 6 | 3618 | 56.6 µs | 15.84× [15.40, 16.30] | +93.7% | [+93.5%, +93.9%] | 0.2 pts | 1.8 | yes |
| multi-str | uuid | 4096 | prefix | ordered | btree-sets | 6 | 92.4 | 185 | 2.00× [1.98, 2.01] | +49.9% | [+49.6%, +50.2%] | 0.3 pts | 0.7 | yes |
| multi-str | uuid | 4096 | prefix | ordered | hashed | 6 | 112 | 48.5 µs | 428.36× [387.63, 478.66] | +99.8% | [+99.7%, +99.8%] | 0.0 pts | 0.5 | yes |
| multi-str | uuid | 4096 | prefix | ordered | map-sets | 6 | 116 | 45.1 µs | 373.88× [327.16, 436.17] | +99.7% | [+99.7%, +99.8%] | 0.0 pts | 0.9 | yes |
| multi-str | uuid | 4096 | churn | ordered | btree-sets | 6 | 91.9 | 198 | 2.17× [2.13, 2.22] | +54.0% | [+53.0%, +55.0%] | 1.0 pts | 1.3 | yes |
| multi-str | uuid | 4096 | churn | ordered | hashed | 6 | 89.1 | 67.2 | 0.76× [0.75, 0.76] | -32.3% | [-33.0%, -31.6%] | 0.7 pts | 0.4 | yes |
| multi-str | uuid | 4096 | churn | ordered | map-sets | 6 | 89.1 | 75.1 | 0.84× [0.83, 0.85] | -18.9% | [-20.2%, -17.6%] | 1.2 pts | 0.6 | yes |
| multi-str | uuid | 4096 | build | ordered | btree-sets | 6 | 7.36 ms | 15.01 ms | 2.03× [2.02, 2.04] | +50.8% | [+50.6%, +51.1%] | 0.2 pts | 0.6 | yes |
| multi-str | uuid | 4096 | build | ordered | hashed | 6 | 7.33 ms | 6.00 ms | 0.82× [0.81, 0.82] | -22.2% | [-22.8%, -21.5%] | 0.6 pts | 0.4 | yes |
| multi-str | uuid | 4096 | build | ordered | map-sets | 6 | 7.31 ms | 6.06 ms | 0.83× [0.82, 0.84] | -20.4% | [-21.7%, -19.1%] | 1.2 pts | 0.8 | yes |
| multi-str | uuid | 16384 | valuesFor | ordered | btree-sets | 6 | 69.4 | 235 | 3.35× [3.29, 3.43] | +70.2% | [+69.6%, +70.8%] | 0.6 pts | 1.1 | yes |
| multi-str | uuid | 16384 | valuesFor | ordered | hashed | 6 | 65.3 | 48.9 | 0.75× [0.74, 0.75] | -33.8% | [-34.7%, -33.0%] | 0.8 pts | 0.5 | yes |
| multi-str | uuid | 16384 | valuesFor | ordered | map-sets | 6 | 68.6 | 111 | 1.62× [1.59, 1.65] | +38.3% | [+37.1%, +39.5%] | 1.1 pts | 1.0 | yes |
| multi-str | uuid | 16384 | valuesBetween | ordered | btree-sets | 6 | 3972 | 8399 | 2.11× [2.09, 2.14] | +52.6% | [+52.1%, +53.2%] | 0.5 pts | 0.7 | yes |
| multi-str | uuid | 16384 | valuesBetween | ordered | hashed | 6 | 5392 | 236.9 µs | 43.73× [40.66, 47.30] | +97.7% | [+97.5%, +97.9%] | 0.2 pts | 1.5 | yes |
| multi-str | uuid | 16384 | valuesBetween | ordered | map-sets | 6 | 5251 | 230.9 µs | 42.64× [37.83, 48.85] | +97.7% | [+97.4%, +98.0%] | 0.3 pts | 2.4 | yes |
| multi-str | uuid | 16384 | prefix | ordered | btree-sets | 6 | 121 | 269 | 2.23× [2.22, 2.23] | +55.1% | [+54.9%, +55.2%] | 0.2 pts | 0.2 | yes |
| multi-str | uuid | 16384 | prefix | ordered | hashed | 6 | 341 | 223.3 µs | 656.58× [641.24, 672.67] | +99.8% | [+99.8%, +99.9%] | 0.0 pts | 0.3 | yes |
| multi-str | uuid | 16384 | prefix | ordered | map-sets | 6 | 328 | 208.3 µs | 644.91× [621.70, 669.93] | +99.8% | [+99.8%, +99.9%] | 0.0 pts | 0.4 | yes |
| multi-str | uuid | 16384 | churn | ordered | btree-sets | 6 | 155 | 328 | 2.12× [2.05, 2.19] | +52.8% | [+51.2%, +54.4%] | 1.5 pts | 1.5 | yes |
| multi-str | uuid | 16384 | churn | ordered | hashed | 6 | 139 | 102 | 0.74× [0.73, 0.75] | -35.8% | [-37.8%, -33.9%] | 1.9 pts | 0.9 | yes |
| multi-str | uuid | 16384 | churn | ordered | map-sets | 14 | 147 | 130 | 0.88× [0.86, 0.89] | -14.2% | [-16.0%, -12.3%] | 3.3 pts | 1.2 | yes |
| multi-str | uuid | 16384 | build | ordered | btree-sets | 6 | 44.74 ms | 94.27 ms | 2.13× [2.09, 2.18] | +53.1% | [+52.1%, +54.1%] | 1.0 pts | 1.1 | yes |
| multi-str | uuid | 16384 | build | ordered | hashed | 6 | 44.36 ms | 35.21 ms | 0.80× [0.79, 0.80] | -25.7% | [-26.5%, -24.9%] | 0.8 pts | 0.5 | yes |
| multi-str | uuid | 16384 | build | ordered | map-sets | 14 | 43.33 ms | 40.86 ms | 0.93× [0.92, 0.95] | -7.3% | [-9.3%, -5.4%] | 3.3 pts | 1.7 | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | btree-sets | 6 | 340 | 791 | 2.34× [2.28, 2.41] | +57.3% | [+56.1%, +58.6%] | 1.2 pts | 2.2 | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | hashed | 30 | 288 | 204 | 0.71× [0.70, 0.71] | -41.6% | [-42.9%, -40.3%] | 3.5 pts | 2.2 | yes |
| multi-str | uuid | 262144 | valuesFor | ordered | map-sets | 6 | 319 | 446 | 1.42× [1.37, 1.46] | +29.4% | [+27.1%, +31.7%] | 2.2 pts | 2.3 | yes |
| multi-str | uuid | 262144 | valuesBetween | ordered | btree-sets | 6 | 11.8 µs | 28.6 µs | 2.43× [2.38, 2.50] | +58.9% | [+57.9%, +59.9%] | 1.0 pts | 1.8 | yes |
| multi-str | uuid | 262144 | prefix | ordered | btree-sets | 6 | 790 | 2041 | 2.59× [2.56, 2.63] | +61.4% | [+60.9%, +61.9%] | 0.5 pts | 1.0 | yes |
| multi-str | uuid | 262144 | churn | ordered | btree-sets | 6 | 596 | 1057 | 1.79× [1.73, 1.85] | +44.1% | [+42.2%, +46.0%] | 1.8 pts | 3.2 | yes |
| multi-str | uuid | 262144 | churn | ordered | hashed | 108 | 487 | 417 | 0.85× [0.84, 0.87] | -17.2% | [-18.9%, -15.5%] | 9.1 pts | 3.7 | yes |
| multi-str | uuid | 262144 | churn | ordered | map-sets | 30 | 540 | 488 | 0.91× [0.90, 0.92] | -10.1% | [-11.2%, -9.0%] | 3.0 pts | 2.0 | yes |
| multi-str | uuid | 1048576 | valuesFor | ordered | btree-sets | 6 | 473 | 1264 | 2.65× [2.60, 2.71] | +62.3% | [+61.5%, +63.1%] | 0.8 pts | 2.4 | yes |
| multi-str | uuid | 1048576 | valuesFor | ordered | hashed | 6 | 420 | 231 | 0.55× [0.54, 0.56] | -81.7% | [-84.7%, -78.7%] | 2.8 pts | 1.7 | yes |
| multi-str | uuid | 1048576 | valuesFor | ordered | map-sets | 30 | 448 | 509 | 1.12× [1.11, 1.14] | +10.9% | [+9.6%, +12.3%] | 3.6 pts | 4.4 | yes |
| multi-str | uuid | 1048576 | valuesBetween | ordered | btree-sets | 6 | 13.0 µs | 32.2 µs | 2.50× [2.40, 2.61] | +60.0% | [+58.3%, +61.7%] | 1.6 pts | 3.4 | yes |
| multi-str | uuid | 1048576 | prefix | ordered | btree-sets | 6 | 2392 | 6381 | 2.65× [2.54, 2.77] | +62.3% | [+60.7%, +63.9%] | 1.5 pts | 3.6 | yes |
| multi-str | uuid | 1048576 | churn | ordered | btree-sets | 6 | 757 | 1563 | 2.09× [2.05, 2.13] | +52.1% | [+51.2%, +53.0%] | 0.9 pts | 2.6 | yes |
| multi-str | uuid | 1048576 | churn | ordered | hashed | 60 | 650 | 520 | 0.80× [0.79, 0.82] | -24.9% | [-27.2%, -22.6%] | 8.9 pts | 4.3 | yes |
| multi-str | uuid | 1048576 | churn | ordered | map-sets | 30 | 731 | 602 | 0.81× [0.80, 0.82] | -23.3% | [-25.2%, -21.4%] | 5.1 pts | 3.1 | yes |
| unique-str | email | 4096 | valuesFor | ordered | btree-map | 4 | 29.8 | 104 | 3.49× [3.46, 3.53] | +71.4% | [+71.1%, +71.7%] | 0.2 pts | 0.7 | yes |
| unique-str | email | 4096 | valuesBetween | ordered | btree-map | 4 | 1471 | 474 | 0.32× [0.32, 0.33] | -210.2% | [-212.8%, -207.6%] | 1.6 pts | 0.5 | yes |
| unique-str | email | 4096 | prefix | ordered | btree-map | 4 | 58.6 | 104 | 1.78× [1.76, 1.80] | +43.7% | [+43.0%, +44.3%] | 0.4 pts | 1.1 | yes |
| unique-str | email | 4096 | churn | ordered | btree-map | 4 | 76.1 | 146 | 1.91× [1.90, 1.93] | +47.8% | [+47.3%, +48.2%] | 0.3 pts | 0.6 | yes |
| unique-str | email | 4096 | build | ordered | btree-map | 4 | 933.0 µs | 1.74 ms | 1.85× [1.82, 1.88] | +46.0% | [+45.0%, +46.9%] | 0.6 pts | 0.5 | yes |
| unique-str | email | 16384 | valuesFor | ordered | btree-map | 4 | 41.2 | 140 | 3.37× [3.32, 3.43] | +70.4% | [+69.9%, +70.9%] | 0.3 pts | 1.9 | yes |
| unique-str | email | 16384 | valuesBetween | ordered | btree-map | 4 | 1788 | 561 | 0.31× [0.31, 0.32] | -219.1% | [-222.2%, -215.9%] | 2.0 pts | 0.7 | yes |
| unique-str | email | 16384 | prefix | ordered | btree-map | 4 | 71.7 | 142 | 1.98× [1.95, 2.01] | +49.4% | [+48.6%, +50.3%] | 0.5 pts | 1.5 | yes |
| unique-str | email | 16384 | churn | ordered | btree-map | 4 | 95.1 | 208 | 2.14× [2.03, 2.27] | +53.3% | [+50.6%, +55.9%] | 1.7 pts | 3.0 | yes |
| unique-str | email | 16384 | build | ordered | btree-map | 4 | 4.60 ms | 9.25 ms | 2.01× [2.00, 2.03] | +50.3% | [+49.9%, +50.8%] | 0.3 pts | 0.5 | yes |
| unique-str | email | 262144 | valuesFor | ordered | btree-map | 14 | 240 | 308 | 1.30× [1.26, 1.34] | +23.2% | [+20.9%, +25.5%] | 3.9 pts | 3.4 | yes |
| unique-str | email | 262144 | valuesBetween | ordered | btree-map | 6 | 5621 | 2295 | 0.41× [0.39, 0.42] | -144.6% | [-153.3%, -136.0%] | 8.2 pts | 2.2 | yes |
| unique-str | email | 262144 | prefix | ordered | btree-map | 6 | 323 | 356 | 1.12× [1.10, 1.14] | +10.9% | [+9.4%, +12.5%] | 1.5 pts | 1.2 | yes |
| unique-str | email | 262144 | churn | ordered | btree-map | 6 | 415 | 632 | 1.52× [1.47, 1.58] | +34.4% | [+32.2%, +36.6%] | 2.1 pts | 0.5 | yes |
| unique-str | email | 1048576 | valuesFor | ordered | btree-map | 4 | 418 | 787 | 1.90× [1.85, 1.95] | +47.2% | [+45.8%, +48.7%] | 0.9 pts | 2.1 | yes |
| unique-str | email | 1048576 | valuesBetween | ordered | btree-map | 4 | 7424 | 3940 | 0.53× [0.53, 0.54] | -88.4% | [-90.3%, -86.5%] | 1.2 pts | 0.6 | yes |
| unique-str | email | 1048576 | prefix | ordered | btree-map | 4 | 545 | 902 | 1.67× [1.64, 1.69] | +40.0% | [+38.9%, +41.0%] | 0.6 pts | 1.0 | yes |
| unique-str | email | 1048576 | churn | ordered | btree-map | 4 | 604 | 1080 | 1.81× [1.73, 1.89] | +44.6% | [+42.1%, +47.2%] | 1.6 pts | 3.0 | yes |
| unique-str | path | 4096 | valuesFor | ordered | btree-map | 6 | 91.7 | 127 | 1.38× [1.36, 1.40] | +27.4% | [+26.5%, +28.3%] | 0.9 pts | 0.9 | yes |
| unique-str | path | 4096 | valuesBetween | ordered | btree-map | 6 | 2535 | 617 | 0.24× [0.24, 0.25] | -309.4% | [-314.1%, -304.6%] | 4.5 pts | 0.5 | yes |
| unique-str | path | 4096 | prefix | ordered | btree-map | 6 | 266 | 168 | 0.63× [0.63, 0.64] | -57.9% | [-59.0%, -56.8%] | 1.1 pts | 0.3 | yes |
| unique-str | path | 4096 | churn | ordered | btree-map | 6 | 210 | 194 | 0.93× [0.91, 0.94] | -7.9% | [-9.4%, -6.5%] | 1.4 pts | 0.7 | yes |
| unique-str | path | 4096 | build | ordered | btree-map | 6 | 2.26 ms | 2.26 ms | 1.00× [0.99, 1.01] | -0.3% | [-1.0%, +0.5%] | 0.7 pts | 0.6 | yes |
| unique-str | path | 16384 | valuesFor | ordered | btree-map | 6 | 124 | 179 | 1.44× [1.42, 1.46] | +30.7% | [+29.8%, +31.6%] | 0.8 pts | 0.9 | yes |
| unique-str | path | 16384 | valuesBetween | ordered | btree-map | 6 | 3156 | 775 | 0.25× [0.24, 0.25] | -305.9% | [-310.4%, -301.5%] | 4.2 pts | 1.0 | yes |
| unique-str | path | 16384 | prefix | ordered | btree-map | 6 | 483 | 282 | 0.59× [0.57, 0.61] | -68.7% | [-74.6%, -62.7%] | 5.7 pts | 0.9 | yes |
| unique-str | path | 16384 | churn | ordered | btree-map | 10 | 284 | 265 | 0.93× [0.91, 0.94] | -8.0% | [-9.8%, -6.1%] | 2.6 pts | 1.5 | yes |
| unique-str | path | 16384 | build | ordered | btree-map | 6 | 12.05 ms | 12.17 ms | 1.01× [1.00, 1.02] | +0.7% | [-0.1%, +1.6%] | 0.8 pts | 0.6 | yes |
| unique-str | path | 262144 | valuesFor | ordered | btree-map | 28 | 554 | 525 | 0.95× [0.94, 0.97] | -4.9% | [-6.8%, -2.9%] | 5.0 pts | 3.6 | yes |
| unique-str | path | 262144 | valuesBetween | ordered | btree-map | 6 | 9383 | 3812 | 0.41× [0.40, 0.43] | -141.4% | [-148.1%, -134.7%] | 6.4 pts | 3.1 | yes |
| unique-str | path | 262144 | prefix | ordered | btree-map | 6 | 8255 | 3445 | 0.42× [0.40, 0.45] | -135.5% | [-148.2%, -122.7%] | 12.1 pts | 0.4 | yes |
| unique-str | path | 262144 | churn | ordered | btree-map | 6 | 911 | 873 | 0.95× [0.94, 0.97] | -4.9% | [-6.4%, -3.4%] | 1.4 pts | 1.1 | yes |
| unique-str | str | 4096 | valuesFor | ordered | btree-map | 4 | 39.7 | 103 | 2.60× [2.54, 2.67] | +61.5% | [+60.6%, +62.5%] | 0.6 pts | 1.5 | yes |
| unique-str | str | 4096 | valuesBetween | ordered | btree-map | 4 | 1844 | 467 | 0.25× [0.25, 0.26] | -295.4% | [-300.8%, -290.0%] | 3.4 pts | 0.9 | yes |
| unique-str | str | 4096 | prefix | ordered | btree-map | 4 | 3624 | 876 | 0.24× [0.23, 0.25] | -318.1% | [-329.1%, -307.0%] | 7.0 pts | 1.5 | yes |
| unique-str | str | 4096 | churn | ordered | btree-map | 4 | 95.2 | 145 | 1.53× [1.51, 1.54] | +34.4% | [+33.9%, +35.0%] | 0.4 pts | 0.5 | yes |
| unique-str | str | 4096 | build | ordered | btree-map | 4 | 1.16 ms | 1.72 ms | 1.49× [1.47, 1.52] | +33.1% | [+31.8%, +34.3%] | 0.8 pts | 0.5 | yes |
| unique-str | str | 16384 | valuesFor | ordered | btree-map | 4 | 50.7 | 140 | 2.75× [2.73, 2.77] | +63.6% | [+63.4%, +63.9%] | 0.2 pts | 0.5 | yes |
| unique-str | str | 16384 | valuesBetween | ordered | btree-map | 4 | 1878 | 547 | 0.29× [0.28, 0.30] | -242.8% | [-250.9%, -234.7%] | 5.1 pts | 2.0 | yes |
| unique-str | str | 16384 | prefix | ordered | btree-map | 4 | 16.0 µs | 3763 | 0.24× [0.23, 0.24] | -325.1% | [-336.3%, -313.8%] | 7.1 pts | 1.9 | yes |
| unique-str | str | 16384 | churn | ordered | btree-map | 4 | 116 | 207 | 1.76× [1.70, 1.83] | +43.3% | [+41.2%, +45.3%] | 1.3 pts | 1.5 | yes |
| unique-str | str | 16384 | build | ordered | btree-map | 4 | 5.42 ms | 9.16 ms | 1.69× [1.66, 1.72] | +40.8% | [+39.8%, +41.8%] | 0.6 pts | 1.0 | yes |
| unique-str | str | 262144 | valuesFor | ordered | btree-map | 6 | 280 | 318 | 1.14× [1.13, 1.15] | +12.3% | [+11.5%, +13.2%] | 0.8 pts | 0.7 | yes |
| unique-str | str | 262144 | valuesBetween | ordered | btree-map | 6 | 6437 | 2434 | 0.37× [0.36, 0.38] | -171.0% | [-181.6%, -160.3%] | 10.1 pts | 1.5 | yes |
| unique-str | str | 262144 | prefix | ordered | btree-map | 6 | 1.01 ms | 302.3 µs | 0.29× [0.28, 0.31] | -239.2% | [-254.7%, -223.6%] | 14.8 pts | 0.7 | yes |
| unique-str | str | 262144 | churn | ordered | btree-map | 6 | 440 | 647 | 1.46× [1.41, 1.51] | +31.4% | [+29.1%, +33.7%] | 2.2 pts | 1.4 | yes |
| unique-str | str | 1048576 | valuesFor | ordered | btree-map | 4 | 408 | 774 | 1.86× [1.79, 1.94] | +46.3% | [+44.3%, +48.3%] | 1.3 pts | 2.2 | yes |
| unique-str | str | 1048576 | valuesBetween | ordered | btree-map | 4 | 7693 | 3972 | 0.51× [0.50, 0.53] | -94.4% | [-99.0%, -89.8%] | 2.9 pts | 1.7 | yes |
| unique-str | str | 1048576 | prefix | ordered | btree-map | 4 | 4.64 ms | 1.82 ms | 0.39× [0.39, 0.40] | -154.7% | [-159.4%, -149.9%] | 3.0 pts | 1.2 | yes |
| unique-str | str | 1048576 | churn | ordered | btree-map | 4 | 636 | 1152 | 1.80× [1.68, 1.95] | +44.5% | [+40.3%, +48.7%] | 2.6 pts | 3.4 | yes |
| unique-str | street | 4096 | valuesFor | ordered | btree-map | 6 | 52.4 | 95.4 | 1.83× [1.81, 1.85] | +45.3% | [+44.9%, +45.8%] | 0.5 pts | 0.8 | yes |
| unique-str | street | 4096 | valuesBetween | ordered | btree-map | 6 | 2120 | 515 | 0.24× [0.24, 0.25] | -311.3% | [-316.5%, -306.2%] | 4.9 pts | 1.1 | yes |
| unique-str | street | 4096 | prefix | ordered | btree-map | 6 | 236 | 146 | 0.62× [0.61, 0.63] | -62.4% | [-64.9%, -59.9%] | 2.4 pts | 0.7 | yes |
| unique-str | street | 4096 | churn | ordered | btree-map | 6 | 122 | 150 | 1.23× [1.22, 1.25] | +19.0% | [+18.1%, +19.9%] | 0.9 pts | 0.7 | yes |
| unique-str | street | 4096 | build | ordered | btree-map | 6 | 1.44 ms | 1.91 ms | 1.32× [1.29, 1.36] | +24.4% | [+22.2%, +26.7%] | 2.1 pts | 2.4 | yes |
| unique-str | street | 16384 | valuesFor | ordered | btree-map | 6 | 74.1 | 132 | 1.81× [1.76, 1.86] | +44.7% | [+43.2%, +46.1%] | 1.4 pts | 2.6 | yes |
| unique-str | street | 16384 | valuesBetween | ordered | btree-map | 6 | 2542 | 611 | 0.24× [0.24, 0.25] | -315.7% | [-325.3%, -306.0%] | 9.2 pts | 2.5 | yes |
| unique-str | street | 16384 | prefix | ordered | btree-map | 6 | 849 | 317 | 0.38× [0.37, 0.39] | -165.5% | [-171.5%, -159.5%] | 5.7 pts | 1.1 | yes |
| unique-str | street | 16384 | churn | ordered | btree-map | 14 | 159 | 203 | 1.26× [1.23, 1.29] | +20.4% | [+18.4%, +22.4%] | 3.5 pts | 2.2 | yes |
| unique-str | street | 16384 | build | ordered | btree-map | 6 | 7.25 ms | 9.88 ms | 1.36× [1.34, 1.38] | +26.5% | [+25.6%, +27.4%] | 0.9 pts | 0.7 | yes |
| unique-str | u64 | 4096 | valuesFor | ordered | btree-map | 4 | 16.5 | 88.2 | 5.35× [5.14, 5.58] | +81.3% | [+80.5%, +82.1%] | 0.5 pts | 2.2 | yes |
| unique-str | u64 | 4096 | valuesBetween | ordered | btree-map | 4 | 1146 | 429 | 0.37× [0.37, 0.38] | -167.0% | [-172.3%, -161.7%] | 3.3 pts | 1.9 | yes |
| unique-str | u64 | 4096 | churn | ordered | btree-map | 4 | 50.7 | 128 | 2.52× [2.50, 2.54] | +60.3% | [+60.0%, +60.6%] | 0.2 pts | 0.6 | yes |
| unique-str | u64 | 4096 | build | ordered | btree-map | 4 | 698.3 µs | 1.48 ms | 2.11× [2.08, 2.14] | +52.6% | [+51.8%, +53.3%] | 0.5 pts | 0.3 | yes |
| unique-str | u64 | 16384 | valuesFor | ordered | btree-map | 4 | 24.8 | 117 | 4.81× [4.45, 5.22] | +79.2% | [+77.5%, +80.9%] | 1.0 pts | 7.1 | yes |
| unique-str | u64 | 16384 | valuesBetween | ordered | btree-map | 4 | 1675 | 470 | 0.28× [0.27, 0.29] | -254.1% | [-266.2%, -241.9%] | 7.6 pts | 3.0 | yes |
| unique-str | u64 | 16384 | churn | ordered | btree-map | 4 | 56.2 | 172 | 3.03× [2.88, 3.18] | +67.0% | [+65.3%, +68.6%] | 1.0 pts | 1.3 | yes |
| unique-str | u64 | 16384 | build | ordered | btree-map | 4 | 3.34 ms | 7.95 ms | 2.38× [2.35, 2.40] | +57.9% | [+57.5%, +58.3%] | 0.3 pts | 0.4 | yes |
| unique-str | u64 | 262144 | valuesFor | ordered | btree-map | 6 | 116 | 237 | 2.03× [1.98, 2.08] | +50.7% | [+49.5%, +51.9%] | 1.1 pts | 2.0 | yes |
| unique-str | u64 | 262144 | valuesBetween | ordered | btree-map | 6 | 4198 | 1161 | 0.28× [0.27, 0.28] | -262.3% | [-272.8%, -251.9%] | 10.0 pts | 1.6 | yes |
| unique-str | u64 | 262144 | churn | ordered | btree-map | 6 | 291 | 424 | 1.46× [1.41, 1.52] | +31.7% | [+29.3%, +34.0%] | 2.2 pts | 1.9 | yes |
| unique-str | u64 | 1048576 | valuesFor | ordered | btree-map | 6 | 172 | 538 | 3.18× [3.08, 3.29] | +68.6% | [+67.5%, +69.6%] | 1.0 pts | 2.9 | yes |
| unique-str | u64 | 1048576 | valuesBetween | ordered | btree-map | 28 | 3868 | 2567 | 0.68× [0.65, 0.70] | -48.0% | [-52.7%, -43.3%] | 12.1 pts | 7.4 | yes |
| unique-str | u64 | 1048576 | churn | ordered | btree-map | 6 | 444 | 996 | 2.26× [2.16, 2.37] | +55.8% | [+53.8%, +57.8%] | 1.9 pts | 3.2 | yes |
| unique-str | url | 4096 | valuesFor | ordered | btree-map | 6 | 73.2 | 123 | 1.68× [1.64, 1.71] | +40.3% | [+39.2%, +41.4%] | 1.1 pts | 1.8 | yes |
| unique-str | url | 4096 | valuesBetween | ordered | btree-map | 6 | 2439 | 589 | 0.24× [0.24, 0.25] | -310.8% | [-317.1%, -304.6%] | 6.0 pts | 1.1 | yes |
| unique-str | url | 4096 | prefix | ordered | btree-map | 6 | 163 | 144 | 0.89× [0.87, 0.90] | -12.8% | [-14.4%, -11.3%] | 1.5 pts | 1.2 | yes |
| unique-str | url | 4096 | churn | ordered | btree-map | 6 | 170 | 187 | 1.10× [1.09, 1.10] | +8.8% | [+8.3%, +9.3%] | 0.5 pts | 0.3 | yes |
| unique-str | url | 4096 | build | ordered | btree-map | 6 | 1.90 ms | 2.10 ms | 1.11× [1.09, 1.12] | +9.7% | [+8.5%, +10.9%] | 1.1 pts | 0.7 | yes |
| unique-str | url | 16384 | valuesFor | ordered | btree-map | 6 | 100 | 172 | 1.71× [1.69, 1.74] | +41.7% | [+40.9%, +42.5%] | 0.8 pts | 1.5 | yes |
| unique-str | url | 16384 | valuesBetween | ordered | btree-map | 6 | 3019 | 750 | 0.25× [0.25, 0.26] | -299.6% | [-307.4%, -291.8%] | 7.4 pts | 1.9 | yes |
| unique-str | url | 16384 | prefix | ordered | btree-map | 6 | 234 | 209 | 0.89× [0.88, 0.90] | -12.3% | [-13.9%, -10.8%] | 1.5 pts | 0.7 | yes |
| unique-str | url | 16384 | churn | ordered | btree-map | 6 | 242 | 261 | 1.09× [1.06, 1.11] | +7.9% | [+6.0%, +9.8%] | 1.8 pts | 0.8 | yes |
| unique-str | url | 16384 | build | ordered | btree-map | 6 | 9.81 ms | 11.54 ms | 1.18× [1.17, 1.20] | +15.5% | [+14.2%, +16.7%] | 1.2 pts | 1.0 | yes |
| unique-str | url | 262144 | valuesFor | ordered | btree-map | 30 | 524 | 522 | 0.99× [0.98, 1.01] | -0.6% | [-2.0%, +0.7%] | 3.6 pts | 2.6 | yes |
| unique-str | url | 262144 | valuesBetween | ordered | btree-map | 6 | 9071 | 3700 | 0.40× [0.39, 0.42] | -147.5% | [-154.7%, -140.3%] | 6.8 pts | 2.9 | yes |
| unique-str | url | 262144 | prefix | ordered | btree-map | 6 | 1856 | 1130 | 0.60× [0.58, 0.63] | -65.4% | [-71.2%, -59.6%] | 5.5 pts | 0.7 | yes |
| unique-str | url | 262144 | churn | ordered | btree-map | 110 | 922 | 824 | 0.93× [0.92, 0.95] | -7.2% | [-9.1%, -5.2%] | 10.3 pts | 4.5 | yes |
| unique-str | url | 1048576 | valuesFor | ordered | btree-map | 6 | 748 | 922 | 1.24× [1.22, 1.26] | +19.2% | [+17.8%, +20.5%] | 1.3 pts | 1.8 | yes |
| unique-str | url | 1048576 | valuesBetween | ordered | btree-map | 6 | 10.7 µs | 4668 | 0.44× [0.43, 0.44] | -128.8% | [-131.3%, -126.3%] | 2.3 pts | 1.4 | yes |
| unique-str | url | 1048576 | prefix | ordered | btree-map | 6 | 6289 | 3279 | 0.52× [0.51, 0.54] | -91.2% | [-96.4%, -85.9%] | 5.0 pts | 0.3 | yes |
| unique-str | url | 1048576 | churn | ordered | btree-map | 6 | 1183 | 1342 | 1.14× [1.11, 1.16] | +12.0% | [+10.1%, +13.8%] | 1.8 pts | 1.5 | yes |
| unique-str | uuid | 4096 | valuesFor | ordered | btree-map | 4 | 33.5 | 103 | 3.09× [2.98, 3.20] | +67.6% | [+66.5%, +68.8%] | 0.7 pts | 2.0 | yes |
| unique-str | uuid | 4096 | valuesBetween | ordered | btree-map | 4 | 1643 | 461 | 0.28× [0.28, 0.28] | -256.6% | [-259.0%, -254.2%] | 1.5 pts | 0.6 | yes |
| unique-str | uuid | 4096 | prefix | ordered | btree-map | 4 | 70.7 | 106 | 1.50× [1.48, 1.53] | +33.4% | [+32.3%, +34.6%] | 0.7 pts | 1.0 | yes |
| unique-str | uuid | 4096 | churn | ordered | btree-map | 4 | 85.6 | 146 | 1.69× [1.65, 1.73] | +40.8% | [+39.4%, +42.3%] | 0.9 pts | 1.3 | yes |
| unique-str | uuid | 4096 | build | ordered | btree-map | 4 | 1.03 ms | 1.70 ms | 1.65× [1.58, 1.73] | +39.5% | [+36.7%, +42.2%] | 1.7 pts | 1.2 | yes |
| unique-str | uuid | 16384 | valuesFor | ordered | btree-map | 6 | 41.9 | 143 | 3.40× [3.35, 3.45] | +70.6% | [+70.2%, +71.0%] | 0.4 pts | 1.8 | yes |
| unique-str | uuid | 16384 | valuesBetween | ordered | btree-map | 6 | 1774 | 563 | 0.32× [0.32, 0.32] | -214.6% | [-217.0%, -212.3%] | 2.2 pts | 1.0 | yes |
| unique-str | uuid | 16384 | prefix | ordered | btree-map | 6 | 86.0 | 150 | 1.75× [1.73, 1.76] | +42.8% | [+42.3%, +43.3%] | 0.4 pts | 0.7 | yes |
| unique-str | uuid | 16384 | churn | ordered | btree-map | 6 | 106 | 201 | 1.86× [1.76, 1.97] | +46.3% | [+43.3%, +49.2%] | 2.8 pts | 2.9 | yes |
| unique-str | uuid | 16384 | build | ordered | btree-map | 6 | 4.89 ms | 9.31 ms | 1.89× [1.88, 1.91] | +47.2% | [+46.8%, +47.6%] | 0.4 pts | 0.6 | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | btree-map | 6 | 268 | 404 | 1.55× [1.49, 1.61] | +35.5% | [+33.1%, +38.0%] | 2.3 pts | 1.2 | yes |
| unique-str | uuid | 262144 | valuesBetween | ordered | btree-map | 6 | 6151 | 2830 | 0.46× [0.45, 0.48] | -116.3% | [-124.6%, -108.0%] | 7.9 pts | 2.5 | yes |
| unique-str | uuid | 262144 | prefix | ordered | btree-map | 22 | 467 | 537 | 1.16× [1.14, 1.19] | +14.1% | [+12.1%, +16.1%] | 4.5 pts | 4.4 | yes |
| unique-str | uuid | 262144 | churn | ordered | btree-map | 6 | 436 | 710 | 1.64× [1.61, 1.67] | +38.9% | [+37.8%, +40.0%] | 1.1 pts | 1.2 | yes |
| unique-str | uuid | 1048576 | valuesFor | ordered | btree-map | 6 | 397 | 822 | 2.08× [2.03, 2.12] | +51.8% | [+50.9%, +52.8%] | 0.9 pts | 2.1 | yes |
| unique-str | uuid | 1048576 | valuesBetween | ordered | btree-map | 6 | 7262 | 4015 | 0.56× [0.55, 0.57] | -79.6% | [-82.2%, -77.0%] | 2.5 pts | 1.5 | yes |
| unique-str | uuid | 1048576 | prefix | ordered | btree-map | 6 | 1394 | 1411 | 1.02× [1.00, 1.04] | +2.0% | [+0.1%, +3.8%] | 1.8 pts | 2.0 | yes |
| unique-str | uuid | 1048576 | churn | ordered | btree-map | 6 | 612 | 1148 | 1.88× [1.79, 1.97] | +46.7% | [+44.2%, +49.2%] | 2.4 pts | 4.1 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi-str email n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 valuesFor: ordered vs map-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 churn: ordered vs hashed: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=262144 churn: ordered vs hashed: 1 processes resolved A as faster and 87 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str email n=1048576 valuesFor: ordered vs hashed: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=1048576 valuesFor: ordered vs map-sets: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=1048576 valuesBetween: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=1048576 prefix: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=1048576 churn: ordered vs hashed: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str email n=1048576 churn: ordered vs map-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=16384 valuesFor: ordered vs hashed: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=16384 valuesBetween: ordered vs hashed: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=16384 valuesBetween: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str path n=262144 churn: ordered vs map-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=16384 valuesBetween: ordered vs map-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 churn: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 churn: ordered vs hashed: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=262144 churn: ordered vs map-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=1048576 valuesFor: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=1048576 valuesBetween: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=1048576 prefix: ordered vs btree-sets: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=1048576 churn: ordered vs hashed: the processes scatter 6.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str str n=1048576 churn: ordered vs map-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 churn: ordered vs hashed: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 churn: ordered vs map-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=4096 valuesFor: ordered vs hashed: the pooled difference of -1.16% does not clear the 2.07% median noise floor of the processes
- multi-str u64 n=4096 valuesFor: ordered vs hashed: the pooled interval [-2.56%, 0.23%] includes zero
- multi-str u64 n=4096 churn: ordered vs hashed: the pooled difference of 0.79% does not clear the 1.51% median noise floor of the processes
- multi-str u64 n=4096 churn: ordered vs hashed: the pooled interval [-0.54%, 2.11%] includes zero
- multi-str u64 n=4096 build: ordered vs hashed: the pooled difference of 1.45% does not clear the 2.04% median noise floor of the processes
- multi-str u64 n=4096 build: ordered vs map-sets: the pooled difference of -0.68% does not clear the 1.87% median noise floor of the processes
- multi-str u64 n=4096 build: ordered vs map-sets: the pooled interval [-2.43%, 1.06%] includes zero
- multi-str u64 n=16384 churn: ordered vs hashed: the pooled difference of 0.01% does not clear the 2.46% median noise floor of the processes
- multi-str u64 n=16384 churn: ordered vs hashed: the pooled interval [-1.76%, 1.79%] includes zero
- multi-str u64 n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 churn: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=262144 churn: ordered vs map-sets: the pooled difference of -0.69% does not clear the 1.96% median noise floor of the processes
- multi-str u64 n=262144 churn: ordered vs map-sets: the pooled interval [-2.51%, 1.14%] includes zero
- multi-str u64 n=262144 churn: ordered vs map-sets: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str u64 n=1048576 valuesFor: ordered vs hashed: the pooled difference of 0.26% does not clear the 1.59% median noise floor of the processes
- multi-str u64 n=1048576 valuesFor: ordered vs hashed: the pooled interval [-1.08%, 1.61%] includes zero
- multi-str u64 n=1048576 valuesFor: ordered vs map-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=1048576 valuesBetween: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=1048576 churn: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str u64 n=1048576 churn: ordered vs hashed: 3 processes resolved A as faster and 8 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str u64 n=1048576 churn: ordered vs map-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=16384 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=16384 churn: ordered vs hashed: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=16384 churn: ordered vs map-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 valuesFor: ordered vs hashed: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 churn: ordered vs hashed: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=262144 churn: ordered vs map-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 valuesFor: ordered vs btree-sets: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 valuesFor: ordered vs hashed: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 churn: ordered vs hashed: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str url n=1048576 churn: ordered vs map-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=4096 valuesBetween: ordered vs hashed: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=16384 valuesBetween: ordered vs map-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesFor: ordered vs hashed: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 valuesFor: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 churn: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 churn: ordered vs hashed: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=262144 churn: ordered vs hashed: 2 processes resolved A as faster and 102 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str uuid n=1048576 valuesFor: ordered vs btree-sets: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 valuesFor: ordered vs map-sets: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 valuesBetween: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 prefix: ordered vs btree-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 churn: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 churn: ordered vs hashed: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str uuid n=1048576 churn: ordered vs map-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=4096 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=16384 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=16384 churn: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=262144 valuesFor: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=1048576 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=1048576 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=1048576 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=1048576 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=1048576 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str email n=1048576 churn: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=4096 build: ordered vs btree-map: the pooled difference of -0.26% does not clear the 1.78% median noise floor of the processes
- unique-str path n=4096 build: ordered vs btree-map: the pooled interval [-1.05%, 0.52%] includes zero
- unique-str path n=16384 build: ordered vs btree-map: the pooled difference of 0.72% does not clear the 1.89% median noise floor of the processes
- unique-str path n=16384 build: ordered vs btree-map: the pooled interval [-0.14%, 1.58%] includes zero
- unique-str path n=262144 valuesFor: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=262144 valuesFor: ordered vs btree-map: 2 processes resolved A as faster and 20 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str path n=262144 valuesBetween: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=4096 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=16384 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=1048576 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=1048576 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=1048576 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=1048576 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=1048576 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str str n=1048576 churn: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 build: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 churn: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str u64 n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str u64 n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str u64 n=16384 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 7.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=16384 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str u64 n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=16384 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str u64 n=16384 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str u64 n=1048576 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=1048576 valuesBetween: ordered vs btree-map: the processes scatter 7.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=1048576 churn: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesFor: ordered vs btree-map: the pooled difference of -0.64% does not clear the 1.57% median noise floor of the processes
- unique-str url n=262144 valuesFor: ordered vs btree-map: the pooled interval [-1.97%, 0.68%] includes zero
- unique-str url n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesFor: ordered vs btree-map: 2 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str url n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 churn: ordered vs btree-map: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 churn: ordered vs btree-map: 15 processes resolved A as faster and 54 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str uuid n=4096 valuesFor: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str uuid n=4096 valuesBetween: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str uuid n=4096 prefix: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str uuid n=4096 churn: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str uuid n=4096 build: ordered vs btree-map: only 4 processes were combined; below five the interval is wide and the heterogeneity figures are rough
- unique-str uuid n=16384 churn: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 prefix: ordered vs btree-map: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=1048576 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=1048576 churn: ordered vs btree-map: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
