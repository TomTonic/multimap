| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | email | 4096 | valuesFor | ordered | btree-sets | 6 | 49.1 | 175 | 3.58× [3.56, 3.59] | +72.0% | [+71.9%, +72.2%] | 0.1 pts | 0.5 | yes |
| multi | email | 4096 | valuesFor | ordered | hashed | 6 | 49.5 | 40.7 | 0.82× [0.82, 0.83] | -21.5% | [-22.1%, -20.8%] | 0.6 pts | 0.3 | yes |
| multi | email | 4096 | valuesFor | ordered | map-sets | 6 | 48.9 | 93.8 | 1.91× [1.89, 1.94] | +47.8% | [+47.1%, +48.4%] | 0.6 pts | 1.4 | yes |
| multi | email | 4096 | valuesBetween | ordered | btree-sets | 6 | 2835 | 7282 | 2.56× [2.54, 2.59] | +61.0% | [+60.6%, +61.4%] | 0.4 pts | 0.7 | yes |
| multi | email | 4096 | valuesBetween | ordered | hashed | 6 | 3052 | 56.0 µs | 18.53× [18.24, 18.83] | +94.6% | [+94.5%, +94.7%] | 0.1 pts | 0.7 | yes |
| multi | email | 4096 | valuesBetween | ordered | map-sets | 6 | 2934 | 55.7 µs | 18.96× [18.63, 19.30] | +94.7% | [+94.6%, +94.8%] | 0.1 pts | 0.9 | yes |
| multi | email | 4096 | prefix | ordered | btree-sets | 6 | 75.1 | 176 | 2.35× [2.34, 2.36] | +57.4% | [+57.3%, +57.6%] | 0.2 pts | 0.6 | yes |
| multi | email | 4096 | prefix | ordered | hashed | 6 | 85.4 | 49.7 µs | 534.06× [445.74, 666.05] | +99.8% | [+99.8%, +99.8%] | 0.0 pts | 1.1 | yes |
| multi | email | 4096 | prefix | ordered | map-sets | 6 | 80.5 | 44.6 µs | 489.62× [391.80, 652.54] | +99.8% | [+99.7%, +99.8%] | 0.0 pts | 1.2 | yes |
| multi | email | 4096 | churn | ordered | btree-sets | 6 | 64.3 | 187 | 2.93× [2.90, 2.97] | +65.9% | [+65.5%, +66.3%] | 0.4 pts | 0.6 | yes |
| multi | email | 4096 | churn | ordered | hashed | 6 | 62.1 | 53.1 | 0.85× [0.84, 0.87] | -17.1% | [-19.0%, -15.3%] | 1.8 pts | 1.1 | yes |
| multi | email | 4096 | churn | ordered | map-sets | 6 | 63.7 | 62.9 | 0.99× [0.97, 1.00] | -1.5% | [-2.9%, -0.1%] | 1.3 pts | 0.6 | yes |
| multi | email | 4096 | build | ordered | btree-sets | 6 | 5.30 ms | 14.11 ms | 2.67× [2.65, 2.69] | +62.6% | [+62.3%, +62.8%] | 0.2 pts | 0.4 | yes |
| multi | email | 4096 | build | ordered | hashed | 6 | 5.23 ms | 4.57 ms | 0.87× [0.86, 0.88] | -14.7% | [-15.7%, -13.7%] | 0.9 pts | 0.4 | yes |
| multi | email | 4096 | build | ordered | map-sets | 6 | 5.26 ms | 5.11 ms | 0.97× [0.96, 0.99] | -2.7% | [-4.2%, -1.2%] | 1.4 pts | 1.1 | yes |
| multi | email | 16384 | valuesFor | ordered | btree-sets | 6 | 61.1 | 221 | 3.62× [3.59, 3.65] | +72.4% | [+72.1%, +72.6%] | 0.2 pts | 0.7 | yes |
| multi | email | 16384 | valuesFor | ordered | hashed | 6 | 60.1 | 46.2 | 0.77× [0.76, 0.78] | -30.0% | [-31.0%, -29.0%] | 0.9 pts | 0.5 | yes |
| multi | email | 16384 | valuesFor | ordered | map-sets | 6 | 60.0 | 104 | 1.73× [1.70, 1.76] | +42.3% | [+41.3%, +43.2%] | 0.9 pts | 1.4 | yes |
| multi | email | 16384 | valuesBetween | ordered | btree-sets | 6 | 3498 | 7884 | 2.26× [2.24, 2.28] | +55.8% | [+55.4%, +56.2%] | 0.4 pts | 0.7 | yes |
| multi | email | 16384 | valuesBetween | ordered | hashed | 6 | 4197 | 231.0 µs | 55.35× [53.63, 57.18] | +98.2% | [+98.1%, +98.3%] | 0.1 pts | 0.9 | yes |
| multi | email | 16384 | valuesBetween | ordered | map-sets | 6 | 4236 | 218.4 µs | 50.13× [46.75, 54.04] | +98.0% | [+97.9%, +98.1%] | 0.1 pts | 1.6 | yes |
| multi | email | 16384 | prefix | ordered | btree-sets | 6 | 94.5 | 231 | 2.44× [2.41, 2.47] | +59.0% | [+58.5%, +59.6%] | 0.5 pts | 1.2 | yes |
| multi | email | 16384 | prefix | ordered | hashed | 6 | 281 | 216.0 µs | 776.71× [740.26, 816.94] | +99.9% | [+99.9%, +99.9%] | 0.0 pts | 0.6 | yes |
| multi | email | 16384 | prefix | ordered | map-sets | 6 | 275 | 196.9 µs | 714.37× [701.55, 727.66] | +99.9% | [+99.9%, +99.9%] | 0.0 pts | 0.2 | yes |
| multi | email | 16384 | churn | ordered | btree-sets | 6 | 103 | 295 | 2.87× [2.80, 2.93] | +65.1% | [+64.3%, +65.9%] | 0.7 pts | 1.2 | yes |
| multi | email | 16384 | churn | ordered | hashed | 6 | 92.9 | 73.0 | 0.78× [0.76, 0.79] | -29.0% | [-31.8%, -26.3%] | 2.6 pts | 1.1 | yes |
| multi | email | 16384 | churn | ordered | map-sets | 6 | 98.7 | 101 | 1.02× [1.01, 1.04] | +2.3% | [+0.6%, +4.0%] | 1.6 pts | 0.8 | yes |
| multi | email | 16384 | build | ordered | btree-sets | 6 | 27.97 ms | 82.24 ms | 2.94× [2.90, 2.98] | +66.0% | [+65.5%, +66.4%] | 0.4 pts | 0.8 | yes |
| multi | email | 16384 | build | ordered | hashed | 6 | 27.15 ms | 22.39 ms | 0.83× [0.82, 0.83] | -20.8% | [-21.6%, -20.1%] | 0.7 pts | 0.6 | yes |
| multi | email | 16384 | build | ordered | map-sets | 6 | 28.14 ms | 30.36 ms | 1.08× [1.07, 1.10] | +7.8% | [+6.3%, +9.3%] | 1.4 pts | 0.7 | yes |
| multi | email | 262144 | valuesFor | ordered | btree-sets | 6 | 323 | 720 | 2.29× [2.19, 2.41] | +56.4% | [+54.3%, +58.5%] | 2.0 pts | 3.7 | yes |
| multi | email | 262144 | valuesFor | ordered | hashed | 6 | 265 | 181 | 0.70× [0.69, 0.71] | -43.3% | [-45.7%, -41.0%] | 2.2 pts | 1.2 | yes |
| multi | email | 262144 | valuesFor | ordered | map-sets | 6 | 280 | 425 | 1.48× [1.42, 1.55] | +32.6% | [+29.6%, +35.6%] | 2.9 pts | 3.4 | yes |
| multi | email | 262144 | valuesBetween | ordered | btree-sets | 6 | 10.2 µs | 26.8 µs | 2.62× [2.56, 2.68] | +61.8% | [+61.0%, +62.7%] | 0.8 pts | 2.1 | yes |
| multi | email | 262144 | prefix | ordered | btree-sets | 6 | 425 | 872 | 2.04× [2.01, 2.08] | +51.1% | [+50.2%, +51.9%] | 0.8 pts | 1.4 | yes |
| multi | email | 262144 | churn | ordered | btree-sets | 6 | 478 | 957 | 2.01× [1.97, 2.05] | +50.2% | [+49.1%, +51.3%] | 1.0 pts | 2.2 | yes |
| multi | email | 262144 | churn | ordered | hashed | 92 | 362 | 310 | 0.84× [0.83, 0.85] | -19.0% | [-20.7%, -17.4%] | 8.1 pts | 3.2 | yes |
| multi | email | 262144 | churn | ordered | map-sets | 30 | 392 | 379 | 0.97× [0.96, 0.99] | -2.7% | [-4.1%, -1.4%] | 3.6 pts | 3.1 | yes |
| multi | email | 1048576 | valuesFor | ordered | btree-sets | 6 | 462 | 1193 | 2.59× [2.53, 2.66] | +61.5% | [+60.5%, +62.4%] | 0.9 pts | 3.1 | yes |
| multi | email | 1048576 | valuesFor | ordered | hashed | 6 | 424 | 223 | 0.53× [0.51, 0.56] | -87.4% | [-94.9%, -79.8%] | 7.2 pts | 3.6 | yes |
| multi | email | 1048576 | valuesFor | ordered | map-sets | 28 | 400 | 478 | 1.17× [1.15, 1.19] | +14.7% | [+13.0%, +16.3%] | 4.3 pts | 5.4 | yes |
| multi | email | 1048576 | valuesBetween | ordered | btree-sets | 6 | 11.7 µs | 30.9 µs | 2.65× [2.59, 2.72] | +62.3% | [+61.4%, +63.3%] | 0.9 pts | 2.5 | yes |
| multi | email | 1048576 | prefix | ordered | btree-sets | 6 | 720 | 1940 | 2.69× [2.61, 2.77] | +62.8% | [+61.7%, +63.9%] | 1.1 pts | 2.3 | yes |
| multi | email | 1048576 | churn | ordered | btree-sets | 6 | 659 | 1526 | 2.26× [2.12, 2.43] | +55.8% | [+52.8%, +58.8%] | 2.9 pts | 4.0 | yes |
| multi | email | 1048576 | churn | ordered | hashed | 28 | 521 | 411 | 0.79× [0.77, 0.80] | -27.1% | [-29.7%, -24.5%] | 6.7 pts | 2.8 | yes |
| multi | email | 1048576 | churn | ordered | map-sets | 28 | 575 | 488 | 0.83× [0.82, 0.85] | -20.1% | [-21.9%, -18.3%] | 4.7 pts | 2.7 | yes |
| multi | path | 4096 | valuesFor | ordered | btree-sets | 6 | 108 | 202 | 1.88× [1.86, 1.89] | +46.7% | [+46.3%, +47.1%] | 0.4 pts | 0.7 | yes |
| multi | path | 4096 | valuesFor | ordered | hashed | 6 | 105 | 47.9 | 0.45× [0.45, 0.46] | -120.0% | [-121.9%, -118.0%] | 1.9 pts | 1.0 | yes |
| multi | path | 4096 | valuesFor | ordered | map-sets | 6 | 105 | 97.9 | 0.93× [0.93, 0.94] | -7.2% | [-7.7%, -6.8%] | 0.4 pts | 0.4 | yes |
| multi | path | 4096 | valuesBetween | ordered | btree-sets | 6 | 4026 | 7671 | 1.89× [1.87, 1.92] | +47.2% | [+46.4%, +48.0%] | 0.7 pts | 0.9 | yes |
| multi | path | 4096 | valuesBetween | ordered | hashed | 6 | 4278 | 65.0 µs | 15.18× [15.04, 15.34] | +93.4% | [+93.3%, +93.5%] | 0.1 pts | 0.5 | yes |
| multi | path | 4096 | valuesBetween | ordered | map-sets | 6 | 4151 | 65.3 µs | 15.71× [15.48, 15.95] | +93.6% | [+93.5%, +93.7%] | 0.1 pts | 0.9 | yes |
| multi | path | 4096 | prefix | ordered | btree-sets | 6 | 331 | 481 | 1.46× [1.44, 1.49] | +31.7% | [+30.3%, +33.0%] | 1.3 pts | 0.5 | yes |
| multi | path | 4096 | prefix | ordered | hashed | 6 | 434 | 59.8 µs | 136.15× [130.19, 142.67] | +99.3% | [+99.2%, +99.3%] | 0.0 pts | 1.1 | yes |
| multi | path | 4096 | prefix | ordered | map-sets | 6 | 407 | 54.5 µs | 133.58× [129.57, 137.85] | +99.3% | [+99.2%, +99.3%] | 0.0 pts | 0.7 | yes |
| multi | path | 4096 | churn | ordered | btree-sets | 6 | 164 | 243 | 1.48× [1.47, 1.50] | +32.7% | [+31.9%, +33.4%] | 0.7 pts | 0.4 | yes |
| multi | path | 4096 | churn | ordered | hashed | 6 | 155 | 67.2 | 0.43× [0.43, 0.44] | -131.8% | [-134.1%, -129.5%] | 2.2 pts | 0.6 | yes |
| multi | path | 4096 | churn | ordered | map-sets | 6 | 157 | 78.0 | 0.51× [0.49, 0.53] | -96.9% | [-105.9%, -88.0%] | 8.5 pts | 1.9 | yes |
| multi | path | 4096 | build | ordered | btree-sets | 6 | 11.44 ms | 17.55 ms | 1.54× [1.53, 1.55] | +35.1% | [+34.6%, +35.7%] | 0.5 pts | 0.7 | yes |
| multi | path | 4096 | build | ordered | hashed | 6 | 11.40 ms | 5.34 ms | 0.47× [0.46, 0.47] | -114.5% | [-116.3%, -112.6%] | 1.8 pts | 0.8 | yes |
| multi | path | 4096 | build | ordered | map-sets | 6 | 11.46 ms | 6.24 ms | 0.54× [0.54, 0.55] | -84.1% | [-86.0%, -82.1%] | 1.9 pts | 0.7 | yes |
| multi | path | 16384 | valuesFor | ordered | btree-sets | 6 | 140 | 262 | 1.88× [1.87, 1.89] | +46.8% | [+46.4%, +47.2%] | 0.4 pts | 0.6 | yes |
| multi | path | 16384 | valuesFor | ordered | hashed | 6 | 134 | 53.2 | 0.40× [0.39, 0.40] | -152.7% | [-154.6%, -150.8%] | 1.8 pts | 0.8 | yes |
| multi | path | 16384 | valuesFor | ordered | map-sets | 8 | 139 | 110 | 0.79× [0.78, 0.80] | -26.8% | [-28.0%, -25.6%] | 1.4 pts | 0.7 | yes |
| multi | path | 16384 | valuesBetween | ordered | btree-sets | 6 | 4724 | 8418 | 1.78× [1.75, 1.80] | +43.7% | [+43.0%, +44.4%] | 0.7 pts | 0.8 | yes |
| multi | path | 16384 | valuesBetween | ordered | hashed | 6 | 5816 | 275.6 µs | 47.67× [46.44, 48.96] | +97.9% | [+97.8%, +98.0%] | 0.1 pts | 0.5 | yes |
| multi | path | 16384 | valuesBetween | ordered | map-sets | 6 | 5990 | 262.0 µs | 43.15× [41.05, 45.48] | +97.7% | [+97.6%, +97.8%] | 0.1 pts | 0.8 | yes |
| multi | path | 16384 | prefix | ordered | btree-sets | 6 | 627 | 997 | 1.62× [1.57, 1.67] | +38.2% | [+36.1%, +40.2%] | 2.0 pts | 0.6 | yes |
| multi | path | 16384 | prefix | ordered | hashed | 6 | 1261 | 269.0 µs | 209.27× [201.78, 217.34] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 0.7 | yes |
| multi | path | 16384 | prefix | ordered | map-sets | 6 | 1233 | 248.6 µs | 203.01× [196.13, 210.38] | +99.5% | [+99.5%, +99.5%] | 0.0 pts | 1.2 | yes |
| multi | path | 16384 | churn | ordered | btree-sets | 6 | 267 | 410 | 1.52× [1.50, 1.54] | +34.3% | [+33.5%, +35.0%] | 0.7 pts | 0.6 | yes |
| multi | path | 16384 | churn | ordered | hashed | 6 | 245 | 104 | 0.43× [0.41, 0.44] | -134.6% | [-141.7%, -127.5%] | 6.8 pts | 1.3 | yes |
| multi | path | 16384 | churn | ordered | map-sets | 8 | 232 | 116 | 0.52× [0.49, 0.54] | -93.7% | [-102.1%, -85.4%] | 10.0 pts | 2.5 | yes |
| multi | path | 16384 | build | ordered | btree-sets | 6 | 67.54 ms | 106.65 ms | 1.58× [1.56, 1.60] | +36.7% | [+36.0%, +37.4%] | 0.7 pts | 1.0 | yes |
| multi | path | 16384 | build | ordered | hashed | 6 | 66.07 ms | 28.46 ms | 0.43× [0.43, 0.44] | -131.1% | [-134.2%, -128.1%] | 2.9 pts | 1.2 | yes |
| multi | path | 16384 | build | ordered | map-sets | 6 | 66.84 ms | 39.77 ms | 0.59× [0.59, 0.60] | -68.4% | [-69.8%, -67.0%] | 1.3 pts | 0.7 | yes |
| multi | path | 262144 | valuesFor | ordered | btree-sets | 6 | 611 | 916 | 1.52× [1.48, 1.56] | +34.2% | [+32.5%, +35.8%] | 1.5 pts | 2.1 | yes |
| multi | path | 262144 | valuesFor | ordered | hashed | 6 | 521 | 218 | 0.42× [0.41, 0.43] | -138.9% | [-145.0%, -132.7%] | 5.9 pts | 2.0 | yes |
| multi | path | 262144 | valuesFor | ordered | map-sets | 6 | 527 | 429 | 0.82× [0.81, 0.82] | -22.4% | [-23.3%, -21.6%] | 0.8 pts | 0.6 | yes |
| multi | path | 262144 | valuesBetween | ordered | btree-sets | 6 | 13.6 µs | 27.5 µs | 2.04× [2.02, 2.06] | +50.9% | [+50.4%, +51.3%] | 0.5 pts | 0.7 | yes |
| multi | path | 262144 | prefix | ordered | btree-sets | 6 | 13.6 µs | 29.9 µs | 2.20× [2.11, 2.29] | +54.5% | [+52.7%, +56.3%] | 1.7 pts | 0.3 | yes |
| multi | path | 262144 | churn | ordered | btree-sets | 6 | 959 | 1191 | 1.25× [1.23, 1.27] | +20.0% | [+18.7%, +21.3%] | 1.2 pts | 1.3 | yes |
| multi | path | 262144 | churn | ordered | hashed | 6 | 802 | 386 | 0.48× [0.46, 0.50] | -109.0% | [-116.1%, -102.0%] | 6.7 pts | 2.0 | yes |
| multi | path | 262144 | churn | ordered | map-sets | 6 | 871 | 494 | 0.57× [0.55, 0.58] | -76.9% | [-82.9%, -71.0%] | 5.7 pts | 2.3 | yes |
| multi | str | 4096 | valuesFor | ordered | btree-sets | 6 | 62.1 | 179 | 2.87× [2.86, 2.89] | +65.2% | [+65.0%, +65.4%] | 0.2 pts | 0.7 | yes |
| multi | str | 4096 | valuesFor | ordered | hashed | 6 | 62.4 | 41.5 | 0.67× [0.66, 0.68] | -49.5% | [-51.2%, -47.9%] | 1.6 pts | 0.8 | yes |
| multi | str | 4096 | valuesFor | ordered | map-sets | 6 | 61.9 | 95.0 | 1.53× [1.52, 1.55] | +34.8% | [+34.1%, +35.5%] | 0.7 pts | 1.1 | yes |
| multi | str | 4096 | valuesBetween | ordered | btree-sets | 6 | 3393 | 7523 | 2.21× [2.19, 2.23] | +54.8% | [+54.3%, +55.2%] | 0.4 pts | 0.5 | yes |
| multi | str | 4096 | valuesBetween | ordered | hashed | 6 | 3680 | 56.2 µs | 15.28× [14.90, 15.67] | +93.5% | [+93.3%, +93.6%] | 0.2 pts | 1.2 | yes |
| multi | str | 4096 | valuesBetween | ordered | map-sets | 6 | 3535 | 56.5 µs | 15.98× [15.68, 16.28] | +93.7% | [+93.6%, +93.9%] | 0.1 pts | 0.9 | yes |
| multi | str | 4096 | prefix | ordered | btree-sets | 6 | 7662 | 18.5 µs | 2.40× [2.35, 2.46] | +58.4% | [+57.4%, +59.4%] | 1.0 pts | 1.2 | yes |
| multi | str | 4096 | prefix | ordered | hashed | 6 | 8298 | 57.5 µs | 6.96× [6.83, 7.10] | +85.6% | [+85.4%, +85.9%] | 0.3 pts | 0.6 | yes |
| multi | str | 4096 | prefix | ordered | map-sets | 6 | 7899 | 65.8 µs | 8.34× [8.14, 8.56] | +88.0% | [+87.7%, +88.3%] | 0.3 pts | 1.0 | yes |
| multi | str | 4096 | churn | ordered | btree-sets | 6 | 82.5 | 194 | 2.31× [2.27, 2.36] | +56.7% | [+55.9%, +57.6%] | 0.8 pts | 1.5 | yes |
| multi | str | 4096 | churn | ordered | hashed | 6 | 80.6 | 55.3 | 0.69× [0.67, 0.71] | -44.7% | [-48.2%, -41.1%] | 3.4 pts | 1.5 | yes |
| multi | str | 4096 | churn | ordered | map-sets | 6 | 78.3 | 59.7 | 0.76× [0.75, 0.78] | -30.8% | [-32.9%, -28.7%] | 2.0 pts | 1.0 | yes |
| multi | str | 4096 | build | ordered | btree-sets | 6 | 6.61 ms | 14.87 ms | 2.23× [2.21, 2.25] | +55.2% | [+54.8%, +55.6%] | 0.4 pts | 0.9 | yes |
| multi | str | 4096 | build | ordered | hashed | 6 | 6.53 ms | 4.71 ms | 0.72× [0.72, 0.73] | -38.2% | [-39.7%, -36.6%] | 1.5 pts | 1.5 | yes |
| multi | str | 4096 | build | ordered | map-sets | 6 | 6.58 ms | 5.31 ms | 0.81× [0.80, 0.82] | -23.3% | [-24.8%, -21.9%] | 1.4 pts | 1.2 | yes |
| multi | str | 16384 | valuesFor | ordered | btree-sets | 6 | 78.7 | 229 | 2.90× [2.84, 2.97] | +65.5% | [+64.7%, +66.3%] | 0.8 pts | 1.9 | yes |
| multi | str | 16384 | valuesFor | ordered | hashed | 6 | 76.4 | 47.4 | 0.62× [0.61, 0.62] | -61.4% | [-62.6%, -60.2%] | 1.2 pts | 0.8 | yes |
| multi | str | 16384 | valuesFor | ordered | map-sets | 6 | 78.0 | 110 | 1.41× [1.39, 1.42] | +28.8% | [+28.3%, +29.4%] | 0.5 pts | 0.6 | yes |
| multi | str | 16384 | valuesBetween | ordered | btree-sets | 6 | 3772 | 8223 | 2.18× [2.15, 2.20] | +54.1% | [+53.6%, +54.6%] | 0.5 pts | 1.0 | yes |
| multi | str | 16384 | valuesBetween | ordered | hashed | 6 | 6053 | 232.4 µs | 38.37× [34.75, 42.82] | +97.4% | [+97.1%, +97.7%] | 0.3 pts | 2.0 | yes |
| multi | str | 16384 | valuesBetween | ordered | map-sets | 6 | 6269 | 223.8 µs | 35.81× [32.53, 39.83] | +97.2% | [+96.9%, +97.5%] | 0.3 pts | 1.8 | yes |
| multi | str | 16384 | prefix | ordered | btree-sets | 6 | 35.3 µs | 81.0 µs | 2.28× [2.24, 2.32] | +56.2% | [+55.4%, +56.9%] | 0.7 pts | 1.0 | yes |
| multi | str | 16384 | prefix | ordered | hashed | 6 | 37.4 µs | 256.0 µs | 6.80× [6.61, 7.01] | +85.3% | [+84.9%, +85.7%] | 0.4 pts | 1.3 | yes |
| multi | str | 16384 | prefix | ordered | map-sets | 6 | 37.4 µs | 297.4 µs | 7.85× [7.43, 8.31] | +87.3% | [+86.5%, +88.0%] | 0.7 pts | 2.8 | yes |
| multi | str | 16384 | churn | ordered | btree-sets | 6 | 151 | 335 | 2.30× [2.21, 2.40] | +56.6% | [+54.7%, +58.4%] | 1.8 pts | 2.2 | yes |
| multi | str | 16384 | churn | ordered | hashed | 6 | 131 | 86.0 | 0.66× [0.65, 0.68] | -50.4% | [-54.3%, -46.4%] | 3.8 pts | 1.0 | yes |
| multi | str | 16384 | churn | ordered | map-sets | 26 | 114 | 98.9 | 0.86× [0.85, 0.88] | -15.8% | [-17.8%, -13.9%] | 4.8 pts | 2.1 | yes |
| multi | str | 16384 | build | ordered | btree-sets | 6 | 36.23 ms | 90.33 ms | 2.48× [2.44, 2.53] | +59.7% | [+59.1%, +60.4%] | 0.7 pts | 1.0 | yes |
| multi | str | 16384 | build | ordered | hashed | 6 | 35.30 ms | 24.19 ms | 0.69× [0.68, 0.70] | -45.5% | [-47.4%, -43.6%] | 1.8 pts | 1.2 | yes |
| multi | str | 16384 | build | ordered | map-sets | 26 | 32.65 ms | 29.17 ms | 0.89× [0.89, 0.90] | -11.9% | [-12.9%, -11.0%] | 2.4 pts | 1.0 | yes |
| multi | str | 262144 | valuesFor | ordered | btree-sets | 6 | 378 | 764 | 2.01× [1.99, 2.03] | +50.3% | [+49.9%, +50.7%] | 0.4 pts | 0.6 | yes |
| multi | str | 262144 | valuesFor | ordered | hashed | 6 | 317 | 192 | 0.60× [0.59, 0.61] | -66.7% | [-69.5%, -63.9%] | 2.7 pts | 1.5 | yes |
| multi | str | 262144 | valuesFor | ordered | map-sets | 6 | 332 | 434 | 1.29× [1.27, 1.32] | +22.7% | [+21.1%, +24.2%] | 1.5 pts | 1.8 | yes |
| multi | str | 262144 | valuesBetween | ordered | btree-sets | 6 | 11.1 µs | 27.3 µs | 2.45× [2.41, 2.50] | +59.2% | [+58.5%, +59.9%] | 0.7 pts | 2.0 | yes |
| multi | str | 262144 | prefix | ordered | btree-sets | 6 | 1.74 ms | 4.50 ms | 2.59× [2.52, 2.67] | +61.4% | [+60.3%, +62.5%] | 1.0 pts | 2.9 | yes |
| multi | str | 262144 | churn | ordered | btree-sets | 6 | 540 | 1004 | 1.90× [1.83, 1.97] | +47.4% | [+45.4%, +49.3%] | 1.8 pts | 2.8 | yes |
| multi | str | 262144 | churn | ordered | hashed | 56 | 405 | 318 | 0.78× [0.76, 0.79] | -28.4% | [-30.8%, -26.0%] | 9.0 pts | 4.3 | yes |
| multi | str | 262144 | churn | ordered | map-sets | 30 | 434 | 384 | 0.89× [0.88, 0.90] | -12.4% | [-13.8%, -11.1%] | 3.6 pts | 2.5 | yes |
| multi | str | 1048576 | valuesFor | ordered | btree-sets | 6 | 532 | 1281 | 2.39× [2.27, 2.51] | +58.1% | [+56.0%, +60.1%] | 2.0 pts | 2.7 | yes |
| multi | str | 1048576 | valuesFor | ordered | hashed | 6 | 451 | 221 | 0.50× [0.48, 0.51] | -101.9% | [-107.9%, -95.9%] | 5.7 pts | 2.4 | yes |
| multi | str | 1048576 | valuesFor | ordered | map-sets | 6 | 469 | 511 | 1.09× [1.07, 1.11] | +8.2% | [+6.6%, +9.8%] | 1.5 pts | 1.6 | yes |
| multi | str | 1048576 | valuesBetween | ordered | btree-sets | 6 | 12.6 µs | 32.8 µs | 2.62× [2.56, 2.68] | +61.8% | [+61.0%, +62.7%] | 0.8 pts | 1.5 | yes |
| multi | str | 1048576 | prefix | ordered | btree-sets | 6 | 7.66 ms | 20.57 ms | 2.67× [2.61, 2.74] | +62.5% | [+61.6%, +63.5%] | 0.9 pts | 2.1 | yes |
| multi | str | 1048576 | churn | ordered | btree-sets | 6 | 795 | 1625 | 2.12× [2.03, 2.21] | +52.8% | [+50.8%, +54.7%] | 1.9 pts | 2.3 | yes |
| multi | str | 1048576 | churn | ordered | hashed | 6 | 584 | 429 | 0.73× [0.71, 0.74] | -37.5% | [-40.6%, -34.4%] | 3.0 pts | 1.3 | yes |
| multi | str | 1048576 | churn | ordered | map-sets | 30 | 628 | 508 | 0.81× [0.80, 0.83] | -23.3% | [-25.6%, -21.0%] | 6.1 pts | 3.6 | yes |
| multi | street | 4096 | valuesFor | ordered | btree-sets | 6 | 55.8 | 141 | 2.54× [2.50, 2.58] | +60.6% | [+60.0%, +61.2%] | 0.6 pts | 1.4 | yes |
| multi | street | 4096 | valuesFor | ordered | hashed | 6 | 53.8 | 20.8 | 0.39× [0.37, 0.40] | -158.6% | [-169.2%, -147.9%] | 10.1 pts | 2.5 | yes |
| multi | street | 4096 | valuesFor | ordered | map-sets | 6 | 55.4 | 69.9 | 1.26× [1.24, 1.27] | +20.6% | [+19.6%, +21.5%] | 0.9 pts | 0.8 | yes |
| multi | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 2278 | 4733 | 2.09× [2.04, 2.14] | +52.2% | [+51.0%, +53.4%] | 1.1 pts | 1.3 | yes |
| multi | street | 4096 | valuesBetween | ordered | hashed | 6 | 2384 | 55.2 µs | 23.27× [22.95, 23.59] | +95.7% | [+95.6%, +95.8%] | 0.1 pts | 0.6 | yes |
| multi | street | 4096 | valuesBetween | ordered | map-sets | 6 | 2321 | 54.8 µs | 23.67× [23.42, 23.93] | +95.8% | [+95.7%, +95.8%] | 0.0 pts | 0.6 | yes |
| multi | street | 4096 | prefix | ordered | btree-sets | 6 | 248 | 585 | 2.38× [2.33, 2.44] | +58.1% | [+57.2%, +59.0%] | 0.9 pts | 1.1 | yes |
| multi | street | 4096 | prefix | ordered | hashed | 6 | 281 | 51.9 µs | 185.29× [181.62, 189.11] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 0.4 | yes |
| multi | street | 4096 | prefix | ordered | map-sets | 6 | 273 | 47.8 µs | 174.27× [170.51, 178.20] | +99.4% | [+99.4%, +99.4%] | 0.0 pts | 0.7 | yes |
| multi | street | 4096 | churn | ordered | btree-sets | 6 | 91.8 | 201 | 2.19× [2.17, 2.21] | +54.3% | [+53.8%, +54.7%] | 0.4 pts | 0.5 | yes |
| multi | street | 4096 | churn | ordered | hashed | 6 | 88.6 | 47.8 | 0.54× [0.52, 0.56] | -84.8% | [-90.6%, -79.0%] | 5.5 pts | 2.0 | yes |
| multi | street | 4096 | churn | ordered | map-sets | 10 | 89.3 | 62.3 | 0.69× [0.68, 0.70] | -45.2% | [-47.9%, -42.4%] | 3.9 pts | 1.8 | yes |
| multi | street | 4096 | build | ordered | btree-sets | 6 | 2.74 ms | 6.21 ms | 2.27× [2.25, 2.28] | +55.9% | [+55.6%, +56.2%] | 0.3 pts | 0.6 | yes |
| multi | street | 4096 | build | ordered | hashed | 6 | 2.74 ms | 1.89 ms | 0.69× [0.67, 0.71] | -45.0% | [-48.2%, -41.8%] | 3.0 pts | 2.0 | yes |
| multi | street | 4096 | build | ordered | map-sets | 10 | 2.73 ms | 2.37 ms | 0.87× [0.85, 0.88] | -15.5% | [-17.5%, -13.5%] | 2.7 pts | 1.5 | yes |
| multi | street | 16384 | valuesFor | ordered | btree-sets | 6 | 77.1 | 188 | 2.41× [2.34, 2.49] | +58.5% | [+57.3%, +59.8%] | 1.2 pts | 2.5 | yes |
| multi | street | 16384 | valuesFor | ordered | hashed | 6 | 75.7 | 29.2 | 0.39× [0.38, 0.39] | -159.1% | [-160.8%, -157.5%] | 1.6 pts | 0.7 | yes |
| multi | street | 16384 | valuesFor | ordered | map-sets | 6 | 76.1 | 80.8 | 1.06× [1.05, 1.07] | +5.8% | [+5.2%, +6.4%] | 0.6 pts | 0.5 | yes |
| multi | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 2760 | 5419 | 1.96× [1.94, 1.99] | +49.1% | [+48.4%, +49.8%] | 0.7 pts | 1.1 | yes |
| multi | street | 16384 | valuesBetween | ordered | hashed | 6 | 3220 | 239.4 µs | 73.61× [71.38, 75.98] | +98.6% | [+98.6%, +98.7%] | 0.0 pts | 1.2 | yes |
| multi | street | 16384 | valuesBetween | ordered | map-sets | 6 | 3211 | 219.8 µs | 68.66× [66.86, 70.56] | +98.5% | [+98.5%, +98.6%] | 0.0 pts | 0.6 | yes |
| multi | street | 16384 | prefix | ordered | btree-sets | 6 | 906 | 2192 | 2.41× [2.36, 2.47] | +58.6% | [+57.6%, +59.6%] | 0.9 pts | 0.8 | yes |
| multi | street | 16384 | prefix | ordered | hashed | 6 | 1320 | 239.7 µs | 178.47× [167.49, 190.99] | +99.4% | [+99.4%, +99.5%] | 0.0 pts | 0.9 | yes |
| multi | street | 16384 | prefix | ordered | map-sets | 6 | 1238 | 215.8 µs | 172.29× [165.87, 179.23] | +99.4% | [+99.4%, +99.4%] | 0.0 pts | 0.5 | yes |
| multi | street | 16384 | churn | ordered | btree-sets | 6 | 133 | 286 | 2.15× [2.13, 2.17] | +53.5% | [+53.1%, +53.9%] | 0.3 pts | 0.3 | yes |
| multi | street | 16384 | churn | ordered | hashed | 6 | 122 | 60.3 | 0.49× [0.48, 0.50] | -102.6% | [-106.3%, -99.0%] | 3.5 pts | 0.9 | yes |
| multi | street | 16384 | churn | ordered | map-sets | 12 | 126 | 88.5 | 0.70× [0.68, 0.72] | -42.9% | [-46.9%, -38.9%] | 6.3 pts | 2.0 | yes |
| multi | street | 16384 | build | ordered | btree-sets | 6 | 15.74 ms | 36.65 ms | 2.32× [2.28, 2.36] | +56.8% | [+56.1%, +57.5%] | 0.7 pts | 1.0 | yes |
| multi | street | 16384 | build | ordered | hashed | 6 | 15.33 ms | 9.99 ms | 0.66× [0.65, 0.66] | -52.6% | [-54.7%, -50.5%] | 2.0 pts | 0.9 | yes |
| multi | street | 16384 | build | ordered | map-sets | 6 | 15.68 ms | 14.38 ms | 0.92× [0.91, 0.93] | -9.2% | [-10.4%, -7.9%] | 1.2 pts | 0.4 | yes |
| multi | u64 | 4096 | valuesFor | ordered | btree-sets | 6 | 35.8 | 161 | 4.48× [4.44, 4.53] | +77.7% | [+77.5%, +77.9%] | 0.2 pts | 0.9 | yes |
| multi | u64 | 4096 | valuesFor | ordered | hashed | 6 | 36.6 | 38.0 | 1.04× [1.03, 1.05] | +3.8% | [+2.9%, +4.7%] | 0.8 pts | 0.4 | yes |
| multi | u64 | 4096 | valuesFor | ordered | map-sets | 6 | 36.3 | 92.2 | 2.54× [2.52, 2.56] | +60.6% | [+60.3%, +60.9%] | 0.3 pts | 0.7 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | btree-sets | 6 | 2590 | 7142 | 2.76× [2.74, 2.79] | +63.8% | [+63.5%, +64.1%] | 0.3 pts | 0.4 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | hashed | 6 | 2784 | 52.9 µs | 18.90× [18.56, 19.26] | +94.7% | [+94.6%, +94.8%] | 0.1 pts | 0.9 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | map-sets | 6 | 2609 | 53.6 µs | 20.44× [20.13, 20.77] | +95.1% | [+95.0%, +95.2%] | 0.1 pts | 1.0 | yes |
| multi | u64 | 4096 | churn | ordered | btree-sets | 6 | 45.8 | 164 | 3.63× [3.57, 3.69] | +72.4% | [+72.0%, +72.9%] | 0.4 pts | 0.9 | yes |
| multi | u64 | 4096 | churn | ordered | hashed | 6 | 44.0 | 46.9 | 1.07× [1.06, 1.08] | +6.7% | [+5.8%, +7.7%] | 0.9 pts | 0.7 | yes |
| multi | u64 | 4096 | churn | ordered | map-sets | 6 | 45.0 | 55.6 | 1.24× [1.22, 1.26] | +19.1% | [+17.9%, +20.3%] | 1.1 pts | 0.8 | yes |
| multi | u64 | 4096 | build | ordered | btree-sets | 6 | 4.07 ms | 12.59 ms | 3.10× [3.05, 3.15] | +67.7% | [+67.2%, +68.3%] | 0.5 pts | 0.9 | yes |
| multi | u64 | 4096 | build | ordered | hashed | 6 | 3.96 ms | 4.31 ms | 1.09× [1.08, 1.09] | +7.9% | [+7.2%, +8.5%] | 0.7 pts | 0.4 | yes |
| multi | u64 | 4096 | build | ordered | map-sets | 6 | 3.99 ms | 4.74 ms | 1.19× [1.18, 1.21] | +16.0% | [+15.0%, +17.1%] | 1.0 pts | 0.6 | yes |
| multi | u64 | 16384 | valuesFor | ordered | btree-sets | 6 | 46.1 | 200 | 4.32× [4.27, 4.38] | +76.9% | [+76.6%, +77.2%] | 0.3 pts | 1.4 | yes |
| multi | u64 | 16384 | valuesFor | ordered | hashed | 6 | 45.5 | 45.0 | 0.99× [0.98, 0.99] | -1.1% | [-1.7%, -0.6%] | 0.5 pts | 0.7 | yes |
| multi | u64 | 16384 | valuesFor | ordered | map-sets | 6 | 45.4 | 105 | 2.32× [2.30, 2.35] | +56.9% | [+56.4%, +57.4%] | 0.4 pts | 1.1 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | btree-sets | 6 | 3603 | 7873 | 2.18× [2.17, 2.20] | +54.2% | [+53.9%, +54.5%] | 0.3 pts | 0.5 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | hashed | 6 | 4705 | 212.6 µs | 45.25× [41.32, 50.00] | +97.8% | [+97.6%, +98.0%] | 0.2 pts | 2.5 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | map-sets | 6 | 4681 | 202.9 µs | 44.36× [40.92, 48.42] | +97.7% | [+97.6%, +97.9%] | 0.2 pts | 1.5 | yes |
| multi | u64 | 16384 | churn | ordered | btree-sets | 6 | 69.6 | 271 | 3.94× [3.79, 4.10] | +74.6% | [+73.6%, +75.6%] | 1.0 pts | 1.4 | yes |
| multi | u64 | 16384 | churn | ordered | hashed | 6 | 60.3 | 66.7 | 1.11× [1.09, 1.14] | +10.2% | [+8.3%, +12.0%] | 1.8 pts | 1.0 | yes |
| multi | u64 | 16384 | churn | ordered | map-sets | 6 | 61.9 | 94.3 | 1.51× [1.48, 1.55] | +34.0% | [+32.4%, +35.5%] | 1.5 pts | 1.0 | yes |
| multi | u64 | 16384 | build | ordered | btree-sets | 6 | 20.78 ms | 78.20 ms | 3.70× [3.62, 3.79] | +73.0% | [+72.3%, +73.6%] | 0.6 pts | 1.1 | yes |
| multi | u64 | 16384 | build | ordered | hashed | 6 | 20.41 ms | 21.85 ms | 1.07× [1.06, 1.08] | +6.7% | [+5.8%, +7.5%] | 0.8 pts | 1.0 | yes |
| multi | u64 | 16384 | build | ordered | map-sets | 6 | 21.13 ms | 29.96 ms | 1.42× [1.38, 1.47] | +29.6% | [+27.4%, +31.8%] | 2.1 pts | 1.4 | yes |
| multi | u64 | 262144 | valuesFor | ordered | btree-sets | 6 | 181 | 650 | 3.62× [3.50, 3.75] | +72.4% | [+71.4%, +73.3%] | 0.9 pts | 1.8 | yes |
| multi | u64 | 262144 | valuesFor | ordered | hashed | 6 | 161 | 174 | 1.09× [1.07, 1.10] | +8.1% | [+6.8%, +9.3%] | 1.2 pts | 0.9 | yes |
| multi | u64 | 262144 | valuesFor | ordered | map-sets | 6 | 169 | 412 | 2.45× [2.40, 2.51] | +59.2% | [+58.3%, +60.1%] | 0.9 pts | 1.8 | yes |
| multi | u64 | 262144 | valuesBetween | ordered | btree-sets | 6 | 9073 | 25.9 µs | 2.84× [2.76, 2.92] | +64.8% | [+63.8%, +65.8%] | 1.0 pts | 2.3 | yes |
| multi | u64 | 262144 | churn | ordered | btree-sets | 6 | 348 | 837 | 2.43× [2.33, 2.54] | +58.9% | [+57.1%, +60.6%] | 1.6 pts | 3.0 | yes |
| multi | u64 | 262144 | churn | ordered | hashed | 30 | 256 | 253 | 0.97× [0.96, 0.99] | -2.7% | [-4.7%, -0.7%] | 5.3 pts | 2.3 | yes |
| multi | u64 | 262144 | churn | ordered | map-sets | 6 | 294 | 341 | 1.16× [1.14, 1.19] | +13.9% | [+12.1%, +15.8%] | 1.8 pts | 1.8 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | btree-sets | 6 | 229 | 1086 | 4.59× [4.44, 4.74] | +78.2% | [+77.5%, +78.9%] | 0.7 pts | 3.9 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | hashed | 30 | 202 | 205 | 1.01× [1.00, 1.02] | +0.6% | [-0.2%, +1.5%] | 2.3 pts | 1.9 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | map-sets | 6 | 212 | 484 | 2.29× [2.23, 2.34] | +56.3% | [+55.2%, +57.3%] | 1.0 pts | 2.3 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | btree-sets | 6 | 8396 | 30.1 µs | 3.55× [3.45, 3.65] | +71.8% | [+71.0%, +72.6%] | 0.7 pts | 2.3 | yes |
| multi | u64 | 1048576 | churn | ordered | btree-sets | 6 | 465 | 1462 | 3.13× [3.07, 3.19] | +68.0% | [+67.5%, +68.6%] | 0.5 pts | 1.9 | yes |
| multi | u64 | 1048576 | churn | ordered | hashed | 30 | 367 | 356 | 0.96× [0.95, 0.98] | -3.7% | [-5.7%, -1.7%] | 5.3 pts | 1.1 | yes |
| multi | u64 | 1048576 | churn | ordered | map-sets | 30 | 392 | 429 | 1.09× [1.08, 1.10] | +8.5% | [+7.5%, +9.5%] | 2.7 pts | 2.2 | yes |
| multi | url | 4096 | valuesFor | ordered | btree-sets | 6 | 92.3 | 199 | 2.15× [2.13, 2.18] | +53.5% | [+53.0%, +54.1%] | 0.5 pts | 1.2 | yes |
| multi | url | 4096 | valuesFor | ordered | hashed | 6 | 91.0 | 46.4 | 0.51× [0.50, 0.52] | -95.8% | [-98.2%, -93.5%] | 2.3 pts | 1.3 | yes |
| multi | url | 4096 | valuesFor | ordered | map-sets | 6 | 91.4 | 100.0 | 1.09× [1.08, 1.10] | +8.5% | [+7.8%, +9.2%] | 0.6 pts | 0.9 | yes |
| multi | url | 4096 | valuesBetween | ordered | btree-sets | 6 | 3869 | 7625 | 1.97× [1.95, 1.98] | +49.2% | [+48.8%, +49.5%] | 0.3 pts | 0.5 | yes |
| multi | url | 4096 | valuesBetween | ordered | hashed | 6 | 4073 | 63.7 µs | 15.58× [15.35, 15.81] | +93.6% | [+93.5%, +93.7%] | 0.1 pts | 0.7 | yes |
| multi | url | 4096 | valuesBetween | ordered | map-sets | 6 | 3949 | 63.8 µs | 16.26× [16.07, 16.46] | +93.9% | [+93.8%, +93.9%] | 0.1 pts | 0.7 | yes |
| multi | url | 4096 | prefix | ordered | btree-sets | 6 | 181 | 262 | 1.45× [1.43, 1.46] | +30.9% | [+30.2%, +31.5%] | 0.6 pts | 0.5 | yes |
| multi | url | 4096 | prefix | ordered | hashed | 6 | 304 | 57.0 µs | 185.85× [172.34, 201.67] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 1.6 | yes |
| multi | url | 4096 | prefix | ordered | map-sets | 6 | 276 | 51.6 µs | 184.25× [172.99, 197.09] | +99.5% | [+99.4%, +99.5%] | 0.0 pts | 1.2 | yes |
| multi | url | 4096 | churn | ordered | btree-sets | 6 | 132 | 230 | 1.76× [1.67, 1.86] | +43.2% | [+40.2%, +46.3%] | 2.9 pts | 2.7 | yes |
| multi | url | 4096 | churn | ordered | hashed | 6 | 128 | 65.3 | 0.52× [0.51, 0.52] | -94.0% | [-96.4%, -91.7%] | 2.2 pts | 0.8 | yes |
| multi | url | 4096 | churn | ordered | map-sets | 6 | 130 | 81.2 | 0.62× [0.61, 0.63] | -61.1% | [-63.1%, -59.1%] | 1.9 pts | 0.6 | yes |
| multi | url | 4096 | build | ordered | btree-sets | 6 | 9.59 ms | 16.92 ms | 1.77× [1.75, 1.78] | +43.4% | [+42.9%, +43.8%] | 0.4 pts | 0.7 | yes |
| multi | url | 4096 | build | ordered | hashed | 6 | 9.45 ms | 5.25 ms | 0.56× [0.55, 0.58] | -78.0% | [-82.9%, -73.1%] | 4.6 pts | 4.3 | yes |
| multi | url | 4096 | build | ordered | map-sets | 6 | 9.53 ms | 6.04 ms | 0.64× [0.63, 0.65] | -57.2% | [-59.5%, -54.8%] | 2.2 pts | 1.0 | yes |
| multi | url | 16384 | valuesFor | ordered | btree-sets | 6 | 122 | 265 | 2.16× [2.13, 2.18] | +53.6% | [+53.1%, +54.1%] | 0.4 pts | 0.6 | yes |
| multi | url | 16384 | valuesFor | ordered | hashed | 6 | 116 | 52.2 | 0.45× [0.45, 0.45] | -122.6% | [-124.3%, -120.9%] | 1.6 pts | 0.6 | yes |
| multi | url | 16384 | valuesFor | ordered | map-sets | 6 | 118 | 111 | 0.94× [0.93, 0.95] | -6.3% | [-7.2%, -5.3%] | 0.9 pts | 0.7 | yes |
| multi | url | 16384 | valuesBetween | ordered | btree-sets | 6 | 4588 | 8501 | 1.84× [1.80, 1.89] | +45.8% | [+44.3%, +47.2%] | 1.4 pts | 1.8 | yes |
| multi | url | 16384 | valuesBetween | ordered | hashed | 6 | 5451 | 269.5 µs | 48.32× [44.70, 52.58] | +97.9% | [+97.8%, +98.1%] | 0.2 pts | 2.3 | yes |
| multi | url | 16384 | valuesBetween | ordered | map-sets | 6 | 5928 | 257.2 µs | 40.73× [35.70, 47.43] | +97.5% | [+97.2%, +97.9%] | 0.3 pts | 2.8 | yes |
| multi | url | 16384 | prefix | ordered | btree-sets | 6 | 283 | 438 | 1.54× [1.50, 1.59] | +35.2% | [+33.2%, +37.1%] | 1.8 pts | 1.1 | yes |
| multi | url | 16384 | prefix | ordered | hashed | 6 | 711 | 260.5 µs | 356.43× [323.14, 397.36] | +99.7% | [+99.7%, +99.7%] | 0.0 pts | 1.2 | yes |
| multi | url | 16384 | prefix | ordered | map-sets | 6 | 703 | 239.6 µs | 338.78× [321.59, 357.91] | +99.7% | [+99.7%, +99.7%] | 0.0 pts | 0.7 | yes |
| multi | url | 16384 | churn | ordered | btree-sets | 6 | 214 | 377 | 1.75× [1.72, 1.79] | +43.0% | [+41.9%, +44.1%] | 1.0 pts | 1.3 | yes |
| multi | url | 16384 | churn | ordered | hashed | 6 | 202 | 98.5 | 0.49× [0.48, 0.50] | -105.1% | [-108.8%, -101.4%] | 3.5 pts | 1.0 | yes |
| multi | url | 16384 | churn | ordered | map-sets | 6 | 205 | 130 | 0.64× [0.63, 0.66] | -55.1% | [-59.6%, -50.6%] | 4.3 pts | 0.8 | yes |
| multi | url | 16384 | build | ordered | btree-sets | 6 | 57.48 ms | 101.36 ms | 1.77× [1.74, 1.79] | +43.3% | [+42.4%, +44.3%] | 0.9 pts | 1.6 | yes |
| multi | url | 16384 | build | ordered | hashed | 6 | 56.92 ms | 29.30 ms | 0.51× [0.51, 0.52] | -95.2% | [-97.0%, -93.4%] | 1.7 pts | 0.8 | yes |
| multi | url | 16384 | build | ordered | map-sets | 6 | 57.46 ms | 37.53 ms | 0.66× [0.64, 0.68] | -51.3% | [-55.7%, -46.9%] | 4.2 pts | 1.9 | yes |
| multi | url | 262144 | valuesFor | ordered | btree-sets | 6 | 576 | 898 | 1.57× [1.53, 1.63] | +36.4% | [+34.4%, +38.5%] | 1.9 pts | 2.5 | yes |
| multi | url | 262144 | valuesFor | ordered | hashed | 6 | 496 | 212 | 0.43× [0.42, 0.44] | -133.1% | [-139.9%, -126.2%] | 6.5 pts | 2.4 | yes |
| multi | url | 262144 | valuesFor | ordered | map-sets | 6 | 490 | 433 | 0.88× [0.87, 0.89] | -13.7% | [-15.4%, -12.0%] | 1.6 pts | 1.2 | yes |
| multi | url | 262144 | valuesBetween | ordered | btree-sets | 6 | 13.3 µs | 27.4 µs | 2.07× [2.04, 2.11] | +51.8% | [+51.0%, +52.6%] | 0.8 pts | 1.3 | yes |
| multi | url | 262144 | prefix | ordered | btree-sets | 6 | 2588 | 4982 | 1.95× [1.88, 2.01] | +48.6% | [+46.9%, +50.3%] | 1.6 pts | 0.5 | yes |
| multi | url | 262144 | churn | ordered | btree-sets | 6 | 858 | 1143 | 1.35× [1.33, 1.37] | +25.9% | [+24.6%, +27.1%] | 1.2 pts | 1.4 | yes |
| multi | url | 262144 | churn | ordered | hashed | 6 | 674 | 388 | 0.57× [0.56, 0.59] | -74.9% | [-79.1%, -70.6%] | 4.0 pts | 1.6 | yes |
| multi | url | 262144 | churn | ordered | map-sets | 6 | 763 | 476 | 0.63× [0.61, 0.64] | -59.4% | [-63.4%, -55.5%] | 3.8 pts | 1.6 | yes |
| multi | url | 1048576 | valuesFor | ordered | btree-sets | 6 | 831 | 1406 | 1.67× [1.62, 1.72] | +40.1% | [+38.1%, +42.0%] | 1.8 pts | 2.5 | yes |
| multi | url | 1048576 | valuesFor | ordered | hashed | 6 | 693 | 236 | 0.34× [0.33, 0.35] | -193.4% | [-200.3%, -186.5%] | 6.6 pts | 2.5 | yes |
| multi | url | 1048576 | valuesFor | ordered | map-sets | 6 | 736 | 510 | 0.70× [0.68, 0.71] | -43.7% | [-47.0%, -40.3%] | 3.2 pts | 1.7 | yes |
| multi | url | 1048576 | valuesBetween | ordered | btree-sets | 6 | 14.7 µs | 31.6 µs | 2.15× [2.12, 2.19] | +53.6% | [+52.9%, +54.3%] | 0.7 pts | 1.2 | yes |
| multi | url | 1048576 | prefix | ordered | btree-sets | 6 | 11.8 µs | 27.6 µs | 2.33× [2.25, 2.43] | +57.2% | [+55.5%, +58.9%] | 1.6 pts | 0.3 | yes |
| multi | url | 1048576 | churn | ordered | btree-sets | 6 | 1125 | 1694 | 1.50× [1.48, 1.53] | +33.5% | [+32.3%, +34.6%] | 1.1 pts | 2.1 | yes |
| multi | url | 1048576 | churn | ordered | hashed | 6 | 993 | 482 | 0.49× [0.47, 0.51] | -105.7% | [-114.9%, -96.4%] | 8.8 pts | 2.7 | yes |
| multi | url | 1048576 | churn | ordered | map-sets | 6 | 1034 | 560 | 0.55× [0.53, 0.57] | -83.4% | [-90.0%, -76.7%] | 6.3 pts | 2.6 | yes |
| multi | uuid | 4096 | valuesFor | ordered | btree-sets | 6 | 55.0 | 181 | 3.29× [3.27, 3.30] | +69.6% | [+69.4%, +69.7%] | 0.1 pts | 0.4 | yes |
| multi | uuid | 4096 | valuesFor | ordered | hashed | 6 | 54.5 | 42.5 | 0.78× [0.78, 0.78] | -28.3% | [-28.7%, -27.9%] | 0.4 pts | 0.3 | yes |
| multi | uuid | 4096 | valuesFor | ordered | map-sets | 6 | 54.8 | 96.4 | 1.76× [1.75, 1.77] | +43.2% | [+42.9%, +43.6%] | 0.3 pts | 0.7 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | btree-sets | 6 | 3262 | 7492 | 2.28× [2.26, 2.31] | +56.2% | [+55.7%, +56.7%] | 0.5 pts | 0.6 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | hashed | 6 | 3501 | 56.4 µs | 16.09× [15.75, 16.44] | +93.8% | [+93.7%, +93.9%] | 0.1 pts | 1.1 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | map-sets | 6 | 3345 | 56.7 µs | 16.98× [16.79, 17.17] | +94.1% | [+94.0%, +94.2%] | 0.1 pts | 0.6 | yes |
| multi | uuid | 4096 | prefix | ordered | btree-sets | 6 | 90.5 | 186 | 2.06× [2.04, 2.07] | +51.4% | [+51.1%, +51.6%] | 0.3 pts | 0.7 | yes |
| multi | uuid | 4096 | prefix | ordered | hashed | 6 | 115 | 50.3 µs | 409.86× [336.23, 524.78] | +99.8% | [+99.7%, +99.8%] | 0.1 pts | 1.1 | yes |
| multi | uuid | 4096 | prefix | ordered | map-sets | 6 | 120 | 45.2 µs | 375.06× [341.63, 415.74] | +99.7% | [+99.7%, +99.8%] | 0.0 pts | 0.5 | yes |
| multi | uuid | 4096 | churn | ordered | btree-sets | 6 | 81.2 | 197 | 2.44× [2.41, 2.47] | +59.0% | [+58.6%, +59.5%] | 0.4 pts | 0.5 | yes |
| multi | uuid | 4096 | churn | ordered | hashed | 6 | 76.0 | 56.0 | 0.74× [0.73, 0.74] | -35.6% | [-36.4%, -34.8%] | 0.8 pts | 0.4 | yes |
| multi | uuid | 4096 | churn | ordered | map-sets | 6 | 79.7 | 70.0 | 0.88× [0.87, 0.90] | -13.5% | [-15.3%, -11.7%] | 1.7 pts | 0.6 | yes |
| multi | uuid | 4096 | build | ordered | btree-sets | 6 | 6.14 ms | 14.90 ms | 2.44× [2.42, 2.46] | +59.0% | [+58.6%, +59.3%] | 0.4 pts | 1.1 | yes |
| multi | uuid | 4096 | build | ordered | hashed | 6 | 6.01 ms | 4.82 ms | 0.80× [0.79, 0.80] | -25.3% | [-26.0%, -24.6%] | 0.6 pts | 0.8 | yes |
| multi | uuid | 4096 | build | ordered | map-sets | 6 | 6.08 ms | 5.65 ms | 0.93× [0.91, 0.94] | -7.9% | [-9.3%, -6.5%] | 1.4 pts | 1.0 | yes |
| multi | uuid | 16384 | valuesFor | ordered | btree-sets | 6 | 66.7 | 229 | 3.40× [3.31, 3.50] | +70.6% | [+69.8%, +71.4%] | 0.8 pts | 2.2 | yes |
| multi | uuid | 16384 | valuesFor | ordered | hashed | 6 | 65.3 | 47.9 | 0.74× [0.73, 0.74] | -35.9% | [-37.3%, -34.4%] | 1.4 pts | 1.2 | yes |
| multi | uuid | 16384 | valuesFor | ordered | map-sets | 6 | 66.5 | 110 | 1.64× [1.61, 1.67] | +39.0% | [+37.8%, +40.2%] | 1.1 pts | 1.4 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | btree-sets | 6 | 3637 | 7998 | 2.20× [2.18, 2.23] | +54.6% | [+54.1%, +55.1%] | 0.4 pts | 0.9 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | hashed | 6 | 4377 | 242.4 µs | 50.03× [42.58, 60.64] | +98.0% | [+97.7%, +98.4%] | 0.3 pts | 3.8 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | map-sets | 6 | 5399 | 230.6 µs | 41.43× [35.82, 49.14] | +97.6% | [+97.2%, +98.0%] | 0.4 pts | 2.9 | yes |
| multi | uuid | 16384 | prefix | ordered | btree-sets | 6 | 114 | 260 | 2.27× [2.24, 2.30] | +55.9% | [+55.3%, +56.5%] | 0.5 pts | 0.7 | yes |
| multi | uuid | 16384 | prefix | ordered | hashed | 6 | 340 | 229.9 µs | 677.74× [665.77, 690.15] | +99.9% | [+99.8%, +99.9%] | 0.0 pts | 0.2 | yes |
| multi | uuid | 16384 | prefix | ordered | map-sets | 6 | 313 | 208.1 µs | 661.52× [645.82, 678.01] | +99.8% | [+99.8%, +99.9%] | 0.0 pts | 0.3 | yes |
| multi | uuid | 16384 | churn | ordered | btree-sets | 6 | 118 | 310 | 2.65× [2.58, 2.72] | +62.2% | [+61.2%, +63.2%] | 1.0 pts | 1.6 | yes |
| multi | uuid | 16384 | churn | ordered | hashed | 6 | 107 | 78.1 | 0.73× [0.72, 0.75] | -36.7% | [-39.3%, -34.1%] | 2.5 pts | 0.7 | yes |
| multi | uuid | 16384 | churn | ordered | map-sets | 6 | 101 | 96.0 | 0.94× [0.93, 0.95] | -6.8% | [-8.1%, -5.5%] | 1.2 pts | 0.8 | yes |
| multi | uuid | 16384 | build | ordered | btree-sets | 6 | 31.42 ms | 86.45 ms | 2.72× [2.68, 2.77] | +63.3% | [+62.7%, +63.8%] | 0.6 pts | 1.0 | yes |
| multi | uuid | 16384 | build | ordered | hashed | 6 | 30.65 ms | 23.97 ms | 0.78× [0.77, 0.79] | -28.2% | [-29.4%, -27.0%] | 1.1 pts | 1.2 | yes |
| multi | uuid | 16384 | build | ordered | map-sets | 6 | 31.51 ms | 32.79 ms | 1.04× [1.02, 1.05] | +3.5% | [+1.9%, +5.1%] | 1.5 pts | 0.7 | yes |
| multi | uuid | 262144 | valuesFor | ordered | btree-sets | 6 | 330 | 786 | 2.40× [2.32, 2.49] | +58.4% | [+57.0%, +59.8%] | 1.4 pts | 2.6 | yes |
| multi | uuid | 262144 | valuesFor | ordered | hashed | 6 | 282 | 195 | 0.68× [0.67, 0.70] | -46.5% | [-49.7%, -43.2%] | 3.1 pts | 1.7 | yes |
| multi | uuid | 262144 | valuesFor | ordered | map-sets | 6 | 291 | 425 | 1.45× [1.41, 1.50] | +31.2% | [+29.3%, +33.1%] | 1.8 pts | 3.0 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | btree-sets | 6 | 10.9 µs | 27.4 µs | 2.50× [2.44, 2.55] | +59.9% | [+59.1%, +60.8%] | 0.8 pts | 1.9 | yes |
| multi | uuid | 262144 | prefix | ordered | btree-sets | 6 | 766 | 1975 | 2.59× [2.54, 2.64] | +61.4% | [+60.6%, +62.1%] | 0.7 pts | 1.6 | yes |
| multi | uuid | 262144 | churn | ordered | btree-sets | 6 | 511 | 997 | 1.96× [1.90, 2.03] | +49.1% | [+47.5%, +50.7%] | 1.6 pts | 2.8 | yes |
| multi | uuid | 262144 | churn | ordered | hashed | 72 | 393 | 334 | 0.85× [0.84, 0.86] | -17.9% | [-19.5%, -16.3%] | 6.8 pts | 2.9 | yes |
| multi | uuid | 262144 | churn | ordered | map-sets | 30 | 436 | 406 | 0.94× [0.93, 0.95] | -6.5% | [-7.8%, -5.3%] | 3.4 pts | 2.0 | yes |
| multi | uuid | 1048576 | valuesFor | ordered | btree-sets | 6 | 448 | 1278 | 2.82× [2.74, 2.91] | +64.5% | [+63.5%, +65.6%] | 1.0 pts | 3.1 | yes |
| multi | uuid | 1048576 | valuesFor | ordered | hashed | 6 | 399 | 220 | 0.56× [0.55, 0.56] | -79.3% | [-81.5%, -77.1%] | 2.1 pts | 1.2 | yes |
| multi | uuid | 1048576 | valuesFor | ordered | map-sets | 6 | 418 | 502 | 1.19× [1.18, 1.21] | +16.2% | [+15.1%, +17.3%] | 1.0 pts | 1.0 | yes |
| multi | uuid | 1048576 | valuesBetween | ordered | btree-sets | 6 | 12.1 µs | 31.1 µs | 2.57× [2.51, 2.63] | +61.0% | [+60.1%, +62.0%] | 0.9 pts | 2.5 | yes |
| multi | uuid | 1048576 | prefix | ordered | btree-sets | 6 | 2317 | 6176 | 2.74× [2.64, 2.85] | +63.5% | [+62.1%, +64.9%] | 1.3 pts | 2.8 | yes |
| multi | uuid | 1048576 | churn | ordered | btree-sets | 6 | 680 | 1545 | 2.26× [2.22, 2.30] | +55.7% | [+54.9%, +56.5%] | 0.8 pts | 2.0 | yes |
| multi | uuid | 1048576 | churn | ordered | hashed | 30 | 555 | 448 | 0.81× [0.79, 0.82] | -24.0% | [-26.3%, -21.6%] | 6.3 pts | 2.9 | yes |
| multi | uuid | 1048576 | churn | ordered | map-sets | 30 | 614 | 514 | 0.83× [0.82, 0.84] | -20.7% | [-21.8%, -19.6%] | 2.9 pts | 1.8 | yes |
| unique | email | 4096 | valuesFor | ordered | btree-map | 6 | 26.1 | 101 | 3.89× [3.83, 3.95] | +74.3% | [+73.9%, +74.7%] | 0.4 pts | 0.9 | yes |
| unique | email | 4096 | valuesBetween | ordered | btree-map | 6 | 1164 | 481 | 0.41× [0.41, 0.42] | -141.4% | [-143.3%, -139.5%] | 1.8 pts | 1.0 | yes |
| unique | email | 4096 | prefix | ordered | btree-map | 6 | 55.3 | 101 | 1.82× [1.81, 1.83] | +45.1% | [+44.8%, +45.5%] | 0.3 pts | 0.9 | yes |
| unique | email | 4096 | churn | ordered | btree-map | 6 | 67.8 | 142 | 2.08× [2.06, 2.11] | +52.0% | [+51.5%, +52.5%] | 0.5 pts | 1.0 | yes |
| unique | email | 4096 | build | ordered | btree-map | 6 | 858.9 µs | 1.66 ms | 1.94× [1.90, 1.98] | +48.5% | [+47.4%, +49.6%] | 1.1 pts | 0.5 | yes |
| unique | email | 16384 | valuesFor | ordered | btree-map | 6 | 38.2 | 136 | 3.55× [3.53, 3.57] | +71.8% | [+71.7%, +72.0%] | 0.1 pts | 0.8 | yes |
| unique | email | 16384 | valuesBetween | ordered | btree-map | 6 | 1508 | 552 | 0.37× [0.36, 0.38] | -171.4% | [-176.4%, -166.3%] | 4.8 pts | 1.9 | yes |
| unique | email | 16384 | prefix | ordered | btree-map | 6 | 69.9 | 140 | 1.99× [1.97, 2.02] | +49.8% | [+49.2%, +50.4%] | 0.6 pts | 1.1 | yes |
| unique | email | 16384 | churn | ordered | btree-map | 6 | 91.4 | 200 | 2.20× [2.10, 2.30] | +54.5% | [+52.5%, +56.5%] | 1.9 pts | 1.4 | yes |
| unique | email | 16384 | build | ordered | btree-map | 6 | 4.11 ms | 8.91 ms | 2.16× [2.14, 2.18] | +53.7% | [+53.4%, +54.1%] | 0.3 pts | 0.6 | yes |
| unique | email | 262144 | valuesFor | ordered | btree-map | 20 | 230 | 290 | 1.30× [1.26, 1.34] | +23.0% | [+20.8%, +25.2%] | 4.7 pts | 4.4 | yes |
| unique | email | 262144 | valuesBetween | ordered | btree-map | 6 | 4945 | 2090 | 0.42× [0.41, 0.44] | -136.6% | [-143.7%, -129.6%] | 6.7 pts | 0.8 | yes |
| unique | email | 262144 | prefix | ordered | btree-map | 6 | 292 | 323 | 1.10× [1.09, 1.12] | +9.5% | [+8.2%, +10.8%] | 1.3 pts | 1.1 | yes |
| unique | email | 262144 | churn | ordered | btree-map | 20 | 333 | 551 | 1.65× [1.58, 1.72] | +39.4% | [+36.8%, +41.9%] | 5.4 pts | 1.2 | yes |
| unique | email | 1048576 | valuesFor | ordered | btree-map | 6 | 397 | 753 | 1.91× [1.89, 1.92] | +47.6% | [+47.2%, +48.0%] | 0.4 pts | 1.1 | yes |
| unique | email | 1048576 | valuesBetween | ordered | btree-map | 6 | 6952 | 3821 | 0.55× [0.54, 0.56] | -81.7% | [-84.3%, -79.1%] | 2.5 pts | 1.6 | yes |
| unique | email | 1048576 | prefix | ordered | btree-map | 6 | 532 | 878 | 1.64× [1.62, 1.66] | +39.0% | [+38.3%, +39.7%] | 0.7 pts | 1.0 | yes |
| unique | email | 1048576 | churn | ordered | btree-map | 6 | 607 | 1061 | 1.75× [1.70, 1.80] | +42.9% | [+41.3%, +44.5%] | 1.6 pts | 2.3 | yes |
| unique | path | 4096 | valuesFor | ordered | btree-map | 6 | 86.0 | 124 | 1.44× [1.43, 1.45] | +30.6% | [+29.9%, +31.2%] | 0.6 pts | 0.8 | yes |
| unique | path | 4096 | valuesBetween | ordered | btree-map | 6 | 1978 | 633 | 0.32× [0.31, 0.32] | -215.7% | [-222.4%, -208.9%] | 6.4 pts | 1.4 | yes |
| unique | path | 4096 | prefix | ordered | btree-map | 6 | 240 | 166 | 0.70× [0.69, 0.70] | -43.9% | [-45.7%, -42.0%] | 1.8 pts | 0.6 | yes |
| unique | path | 4096 | churn | ordered | btree-map | 6 | 198 | 189 | 0.96× [0.94, 0.97] | -4.5% | [-6.0%, -3.0%] | 1.5 pts | 0.8 | yes |
| unique | path | 4096 | build | ordered | btree-map | 6 | 2.09 ms | 2.16 ms | 1.04× [1.03, 1.04] | +3.5% | [+2.8%, +4.1%] | 0.6 pts | 0.5 | yes |
| unique | path | 16384 | valuesFor | ordered | btree-map | 6 | 117 | 172 | 1.48× [1.47, 1.49] | +32.5% | [+32.1%, +32.9%] | 0.4 pts | 0.5 | yes |
| unique | path | 16384 | valuesBetween | ordered | btree-map | 6 | 2536 | 763 | 0.30× [0.30, 0.30] | -233.5% | [-237.2%, -229.8%] | 3.5 pts | 0.8 | yes |
| unique | path | 16384 | prefix | ordered | btree-map | 20 | 414 | 277 | 0.67× [0.66, 0.68] | -49.3% | [-51.2%, -47.3%] | 4.2 pts | 0.7 | yes |
| unique | path | 16384 | churn | ordered | btree-map | 20 | 276 | 259 | 0.95× [0.93, 0.97] | -5.4% | [-7.3%, -3.5%] | 4.1 pts | 2.2 | yes |
| unique | path | 16384 | build | ordered | btree-map | 6 | 11.01 ms | 11.37 ms | 1.04× [1.03, 1.04] | +3.4% | [+2.7%, +4.1%] | 0.6 pts | 0.7 | yes |
| unique | path | 262144 | valuesFor | ordered | btree-map | 30 | 529 | 490 | 0.93× [0.92, 0.95] | -7.0% | [-8.6%, -5.4%] | 4.3 pts | 2.9 | yes |
| unique | path | 262144 | valuesBetween | ordered | btree-map | 6 | 7617 | 3465 | 0.46× [0.45, 0.46] | -119.0% | [-122.1%, -115.8%] | 3.0 pts | 1.5 | yes |
| unique | path | 262144 | prefix | ordered | btree-map | 6 | 5886 | 2633 | 0.45× [0.44, 0.46] | -123.4% | [-127.8%, -119.0%] | 4.2 pts | 0.3 | yes |
| unique | path | 262144 | churn | ordered | btree-map | 52 | 892 | 854 | 0.96× [0.94, 0.97] | -4.5% | [-6.2%, -2.9%] | 5.9 pts | 3.6 | yes |
| unique | str | 4096 | valuesFor | ordered | btree-map | 6 | 38.3 | 102 | 2.67× [2.65, 2.69] | +62.6% | [+62.3%, +62.8%] | 0.2 pts | 0.6 | yes |
| unique | str | 4096 | valuesBetween | ordered | btree-map | 6 | 1529 | 480 | 0.32× [0.31, 0.32] | -217.3% | [-223.5%, -211.0%] | 5.9 pts | 2.1 | yes |
| unique | str | 4096 | prefix | ordered | btree-map | 6 | 2938 | 926 | 0.32× [0.31, 0.32] | -216.7% | [-221.2%, -212.2%] | 4.3 pts | 1.0 | yes |
| unique | str | 4096 | churn | ordered | btree-map | 6 | 87.0 | 141 | 1.62× [1.60, 1.64] | +38.2% | [+37.4%, +38.9%] | 0.7 pts | 1.1 | yes |
| unique | str | 4096 | build | ordered | btree-map | 6 | 1.06 ms | 1.66 ms | 1.56× [1.53, 1.60] | +36.1% | [+34.5%, +37.6%] | 1.5 pts | 0.6 | yes |
| unique | str | 16384 | valuesFor | ordered | btree-map | 6 | 50.4 | 136 | 2.69× [2.67, 2.72] | +62.9% | [+62.6%, +63.2%] | 0.3 pts | 1.1 | yes |
| unique | str | 16384 | valuesBetween | ordered | btree-map | 6 | 1612 | 543 | 0.34× [0.34, 0.34] | -196.3% | [-197.8%, -194.8%] | 1.4 pts | 0.6 | yes |
| unique | str | 16384 | prefix | ordered | btree-map | 6 | 13.6 µs | 3786 | 0.28× [0.28, 0.28] | -257.1% | [-262.4%, -251.8%] | 5.0 pts | 1.4 | yes |
| unique | str | 16384 | churn | ordered | btree-map | 6 | 109 | 198 | 1.79× [1.72, 1.87] | +44.2% | [+41.9%, +46.5%] | 2.2 pts | 2.1 | yes |
| unique | str | 16384 | build | ordered | btree-map | 6 | 5.09 ms | 8.80 ms | 1.73× [1.71, 1.75] | +42.3% | [+41.5%, +43.0%] | 0.7 pts | 1.1 | yes |
| unique | str | 262144 | valuesFor | ordered | btree-map | 14 | 267 | 289 | 1.10× [1.09, 1.12] | +9.5% | [+8.2%, +10.8%] | 2.3 pts | 1.5 | yes |
| unique | str | 262144 | valuesBetween | ordered | btree-map | 6 | 5540 | 1940 | 0.36× [0.34, 0.37] | -181.1% | [-193.3%, -169.0%] | 11.6 pts | 1.9 | yes |
| unique | str | 262144 | prefix | ordered | btree-map | 14 | 848.1 µs | 259.4 µs | 0.31× [0.30, 0.32] | -220.2% | [-228.3%, -212.2%] | 13.9 pts | 0.4 | yes |
| unique | str | 262144 | churn | ordered | btree-map | 14 | 403 | 557 | 1.35× [1.31, 1.40] | +26.2% | [+23.9%, +28.4%] | 3.8 pts | 2.7 | yes |
| unique | str | 1048576 | valuesFor | ordered | btree-map | 6 | 416 | 742 | 1.80× [1.76, 1.84] | +44.5% | [+43.3%, +45.8%] | 1.2 pts | 2.0 | yes |
| unique | str | 1048576 | valuesBetween | ordered | btree-map | 6 | 7175 | 3853 | 0.54× [0.53, 0.54] | -86.6% | [-89.0%, -84.2%] | 2.3 pts | 1.7 | yes |
| unique | str | 1048576 | prefix | ordered | btree-map | 6 | 4.32 ms | 1.83 ms | 0.42× [0.42, 0.42] | -136.9% | [-138.4%, -135.4%] | 1.4 pts | 0.7 | yes |
| unique | str | 1048576 | churn | ordered | btree-map | 6 | 615 | 1084 | 1.75× [1.70, 1.81] | +43.0% | [+41.2%, +44.9%] | 1.8 pts | 2.2 | yes |
| unique | street | 4096 | valuesFor | ordered | btree-map | 6 | 49.8 | 93.9 | 1.90× [1.88, 1.92] | +47.4% | [+46.9%, +47.9%] | 0.5 pts | 0.9 | yes |
| unique | street | 4096 | valuesBetween | ordered | btree-map | 6 | 1702 | 528 | 0.31× [0.30, 0.32] | -221.0% | [-229.2%, -212.8%] | 7.9 pts | 2.2 | yes |
| unique | street | 4096 | prefix | ordered | btree-map | 6 | 205 | 146 | 0.71× [0.70, 0.72] | -40.2% | [-42.0%, -38.3%] | 1.8 pts | 0.9 | yes |
| unique | street | 4096 | churn | ordered | btree-map | 6 | 110 | 147 | 1.33× [1.31, 1.35] | +24.9% | [+23.9%, +25.9%] | 1.0 pts | 0.9 | yes |
| unique | street | 4096 | build | ordered | btree-map | 6 | 1.26 ms | 1.80 ms | 1.43× [1.41, 1.44] | +29.9% | [+29.2%, +30.7%] | 0.7 pts | 0.9 | yes |
| unique | street | 16384 | valuesFor | ordered | btree-map | 6 | 70.9 | 135 | 1.88× [1.82, 1.96] | +46.9% | [+45.0%, +48.9%] | 1.8 pts | 3.0 | yes |
| unique | street | 16384 | valuesBetween | ordered | btree-map | 6 | 2113 | 619 | 0.29× [0.29, 0.30] | -239.6% | [-246.6%, -232.6%] | 6.6 pts | 2.1 | yes |
| unique | street | 16384 | prefix | ordered | btree-map | 6 | 704 | 320 | 0.45× [0.45, 0.46] | -119.9% | [-123.7%, -116.0%] | 3.6 pts | 1.0 | yes |
| unique | street | 16384 | churn | ordered | btree-map | 6 | 147 | 195 | 1.33× [1.30, 1.36] | +24.6% | [+22.9%, +26.3%] | 1.6 pts | 0.9 | yes |
| unique | street | 16384 | build | ordered | btree-map | 6 | 6.46 ms | 9.47 ms | 1.46× [1.45, 1.47] | +31.5% | [+30.9%, +32.1%] | 0.6 pts | 0.5 | yes |
| unique | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 14.8 | 85.6 | 5.76× [5.70, 5.83] | +82.6% | [+82.5%, +82.8%] | 0.2 pts | 0.7 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 933 | 440 | 0.47× [0.47, 0.47] | -112.3% | [-113.5%, -111.1%] | 1.1 pts | 0.7 | yes |
| unique | u64 | 4096 | churn | ordered | btree-map | 6 | 43.6 | 124 | 2.83× [2.80, 2.86] | +64.7% | [+64.3%, +65.1%] | 0.4 pts | 1.4 | yes |
| unique | u64 | 4096 | build | ordered | btree-map | 6 | 609.9 µs | 1.41 ms | 2.29× [2.25, 2.34] | +56.4% | [+55.6%, +57.2%] | 0.8 pts | 0.4 | yes |
| unique | u64 | 16384 | valuesFor | ordered | btree-map | 6 | 22.6 | 113 | 5.00× [4.97, 5.03] | +80.0% | [+79.9%, +80.1%] | 0.1 pts | 0.9 | yes |
| unique | u64 | 16384 | valuesBetween | ordered | btree-map | 6 | 1569 | 468 | 0.30× [0.30, 0.30] | -234.6% | [-236.0%, -233.1%] | 1.4 pts | 0.5 | yes |
| unique | u64 | 16384 | churn | ordered | btree-map | 6 | 45.5 | 161 | 3.52× [3.43, 3.61] | +71.6% | [+70.8%, +72.3%] | 0.7 pts | 1.1 | yes |
| unique | u64 | 16384 | build | ordered | btree-map | 6 | 2.81 ms | 7.51 ms | 2.67× [2.65, 2.69] | +62.6% | [+62.3%, +62.9%] | 0.3 pts | 0.6 | yes |
| unique | u64 | 262144 | valuesFor | ordered | btree-map | 6 | 95.7 | 224 | 2.30× [2.21, 2.39] | +56.5% | [+54.8%, +58.1%] | 1.6 pts | 1.8 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | btree-map | 6 | 3131 | 998 | 0.32× [0.31, 0.32] | -215.0% | [-220.0%, -210.0%] | 4.8 pts | 0.9 | yes |
| unique | u64 | 262144 | churn | ordered | btree-map | 10 | 214 | 354 | 1.63× [1.54, 1.73] | +38.5% | [+35.0%, +42.0%] | 5.0 pts | 1.2 | yes |
| unique | u64 | 1048576 | valuesFor | ordered | btree-map | 6 | 161 | 483 | 2.99× [2.96, 3.02] | +66.5% | [+66.2%, +66.9%] | 0.3 pts | 1.0 | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | btree-map | 38 | 3280 | 2178 | 0.68× [0.66, 0.70] | -46.6% | [-51.3%, -41.9%] | 14.3 pts | 7.9 | no |
| unique | u64 | 1048576 | churn | ordered | btree-map | 6 | 412 | 919 | 2.23× [2.20, 2.26] | +55.2% | [+54.6%, +55.8%] | 0.6 pts | 0.9 | yes |
| unique | url | 4096 | valuesFor | ordered | btree-map | 6 | 72.0 | 123 | 1.71× [1.69, 1.73] | +41.6% | [+40.9%, +42.2%] | 0.6 pts | 1.0 | yes |
| unique | url | 4096 | valuesBetween | ordered | btree-map | 6 | 1857 | 607 | 0.33× [0.32, 0.33] | -206.0% | [-211.3%, -200.6%] | 5.1 pts | 1.1 | yes |
| unique | url | 4096 | prefix | ordered | btree-map | 6 | 150 | 146 | 0.98× [0.96, 0.99] | -2.3% | [-3.8%, -0.9%] | 1.4 pts | 1.3 | yes |
| unique | url | 4096 | churn | ordered | btree-map | 6 | 163 | 185 | 1.13× [1.11, 1.15] | +11.3% | [+9.8%, +12.8%] | 1.4 pts | 0.7 | yes |
| unique | url | 4096 | build | ordered | btree-map | 6 | 1.79 ms | 2.08 ms | 1.15× [1.13, 1.17] | +13.3% | [+11.8%, +14.9%] | 1.5 pts | 1.6 | yes |
| unique | url | 16384 | valuesFor | ordered | btree-map | 6 | 98.7 | 172 | 1.75× [1.72, 1.77] | +42.7% | [+41.9%, +43.5%] | 0.8 pts | 1.4 | yes |
| unique | url | 16384 | valuesBetween | ordered | btree-map | 6 | 2388 | 765 | 0.32× [0.32, 0.33] | -210.9% | [-215.0%, -206.8%] | 3.9 pts | 1.5 | yes |
| unique | url | 16384 | prefix | ordered | btree-map | 12 | 207 | 208 | 0.99× [0.98, 1.00] | -1.0% | [-2.1%, +0.2%] | 1.8 pts | 1.3 | yes |
| unique | url | 16384 | churn | ordered | btree-map | 12 | 226 | 249 | 1.12× [1.10, 1.14] | +10.7% | [+8.9%, +12.6%] | 2.9 pts | 1.7 | yes |
| unique | url | 16384 | build | ordered | btree-map | 12 | 8.89 ms | 10.92 ms | 1.22× [1.20, 1.23] | +17.9% | [+17.0%, +18.9%] | 1.5 pts | 1.5 | yes |
| unique | url | 262144 | valuesFor | ordered | btree-map | 30 | 493 | 483 | 0.98× [0.97, 1.00] | -1.8% | [-3.3%, -0.3%] | 4.0 pts | 2.8 | yes |
| unique | url | 262144 | valuesBetween | ordered | btree-map | 6 | 7924 | 3370 | 0.43× [0.42, 0.44] | -131.5% | [-137.5%, -125.6%] | 5.7 pts | 2.0 | yes |
| unique | url | 262144 | prefix | ordered | btree-map | 30 | 1484 | 940 | 0.64× [0.63, 0.65] | -56.7% | [-58.7%, -54.7%] | 5.2 pts | 0.9 | yes |
| unique | url | 262144 | churn | ordered | btree-map | 130 | 666 | 768 | 1.07× [1.05, 1.10] | +6.9% | [+4.5%, +9.4%] | 14.1 pts | 7.0 | no |
| unique | url | 1048576 | valuesFor | ordered | btree-map | 6 | 742 | 917 | 1.24× [1.22, 1.26] | +19.5% | [+18.2%, +20.7%] | 1.2 pts | 1.3 | yes |
| unique | url | 1048576 | valuesBetween | ordered | btree-map | 6 | 9529 | 4558 | 0.48× [0.47, 0.48] | -109.2% | [-112.0%, -106.4%] | 2.7 pts | 1.4 | yes |
| unique | url | 1048576 | prefix | ordered | btree-map | 6 | 5868 | 3322 | 0.56× [0.54, 0.58] | -78.7% | [-85.9%, -71.4%] | 6.9 pts | 0.5 | yes |
| unique | url | 1048576 | churn | ordered | btree-map | 6 | 1105 | 1260 | 1.16× [1.14, 1.17] | +13.4% | [+12.5%, +14.4%] | 0.9 pts | 0.8 | yes |
| unique | uuid | 4096 | valuesFor | ordered | btree-map | 6 | 30.5 | 101 | 3.33× [3.31, 3.35] | +70.0% | [+69.8%, +70.2%] | 0.2 pts | 0.4 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | btree-map | 6 | 1385 | 473 | 0.34× [0.34, 0.35] | -193.1% | [-196.4%, -189.8%] | 3.2 pts | 1.3 | yes |
| unique | uuid | 4096 | prefix | ordered | btree-map | 6 | 66.7 | 104 | 1.55× [1.54, 1.57] | +35.7% | [+35.1%, +36.2%] | 0.5 pts | 1.1 | yes |
| unique | uuid | 4096 | churn | ordered | btree-map | 6 | 81.1 | 141 | 1.73× [1.71, 1.75] | +42.2% | [+41.5%, +43.0%] | 0.7 pts | 1.0 | yes |
| unique | uuid | 4096 | build | ordered | btree-map | 6 | 969.6 µs | 1.63 ms | 1.70× [1.65, 1.75] | +41.2% | [+39.5%, +42.8%] | 1.6 pts | 0.7 | yes |
| unique | uuid | 16384 | valuesFor | ordered | btree-map | 6 | 40.8 | 139 | 3.41× [3.36, 3.47] | +70.7% | [+70.2%, +71.2%] | 0.4 pts | 2.3 | yes |
| unique | uuid | 16384 | valuesBetween | ordered | btree-map | 6 | 1554 | 556 | 0.36× [0.36, 0.37] | -176.9% | [-181.5%, -172.4%] | 4.3 pts | 1.9 | yes |
| unique | uuid | 16384 | prefix | ordered | btree-map | 6 | 82.0 | 146 | 1.78× [1.77, 1.79] | +43.8% | [+43.4%, +44.3%] | 0.4 pts | 0.7 | yes |
| unique | uuid | 16384 | churn | ordered | btree-map | 6 | 100 | 193 | 1.88× [1.81, 1.96] | +46.8% | [+44.7%, +48.9%] | 2.0 pts | 2.0 | yes |
| unique | uuid | 16384 | build | ordered | btree-map | 6 | 4.65 ms | 8.81 ms | 1.89× [1.87, 1.91] | +47.1% | [+46.6%, +47.6%] | 0.5 pts | 0.9 | yes |
| unique | uuid | 262144 | valuesFor | ordered | btree-map | 6 | 252 | 373 | 1.49× [1.45, 1.52] | +32.7% | [+31.1%, +34.3%] | 1.5 pts | 2.1 | yes |
| unique | uuid | 262144 | valuesBetween | ordered | btree-map | 6 | 5902 | 2726 | 0.47× [0.45, 0.49] | -113.8% | [-123.0%, -104.7%] | 8.7 pts | 2.8 | yes |
| unique | uuid | 262144 | prefix | ordered | btree-map | 8 | 441 | 501 | 1.12× [1.10, 1.14] | +10.6% | [+8.7%, +12.4%] | 2.2 pts | 2.0 | yes |
| unique | uuid | 262144 | churn | ordered | btree-map | 6 | 406 | 640 | 1.59× [1.58, 1.61] | +37.3% | [+36.5%, +38.0%] | 0.7 pts | 0.8 | yes |
| unique | uuid | 1048576 | valuesFor | ordered | btree-map | 6 | 371 | 810 | 2.18× [2.14, 2.23] | +54.2% | [+53.3%, +55.1%] | 0.8 pts | 2.2 | yes |
| unique | uuid | 1048576 | valuesBetween | ordered | btree-map | 6 | 7283 | 4000 | 0.55× [0.54, 0.56] | -82.0% | [-85.3%, -78.6%] | 3.2 pts | 1.7 | yes |
| unique | uuid | 1048576 | prefix | ordered | btree-map | 6 | 1378 | 1343 | 0.98× [0.97, 0.99] | -2.0% | [-2.9%, -1.0%] | 0.9 pts | 1.2 | yes |
| unique | uuid | 1048576 | churn | ordered | btree-map | 6 | 588 | 1080 | 1.82× [1.76, 1.87] | +45.0% | [+43.3%, +46.6%] | 1.6 pts | 3.4 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi email n=4096 churn: ordered vs map-sets: the pooled difference of -1.49% does not clear the 1.80% median noise floor of the processes
- multi email n=16384 churn: ordered vs map-sets: the pooled difference of 2.28% does not clear the 3.03% median noise floor of the processes
- multi email n=262144 valuesFor: ordered vs btree-sets: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesFor: ordered vs map-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 churn: ordered vs hashed: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 churn: ordered vs map-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=262144 churn: ordered vs map-sets: 1 processes resolved A as faster and 15 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi email n=1048576 valuesFor: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=1048576 valuesFor: ordered vs hashed: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=1048576 valuesFor: ordered vs map-sets: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=1048576 valuesBetween: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=1048576 prefix: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=1048576 churn: ordered vs btree-sets: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=1048576 churn: ordered vs hashed: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi email n=1048576 churn: ordered vs map-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=16384 churn: ordered vs map-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=262144 valuesFor: ordered vs hashed: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=262144 churn: ordered vs hashed: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=262144 churn: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 prefix: ordered vs map-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=16384 churn: ordered vs map-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 prefix: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 churn: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 churn: ordered vs hashed: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 churn: ordered vs map-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 valuesFor: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 valuesFor: ordered vs hashed: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 prefix: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=1048576 churn: ordered vs map-sets: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 valuesFor: ordered vs hashed: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 churn: ordered vs hashed: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=4096 build: ordered vs hashed: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=16384 valuesFor: ordered vs hashed: the pooled difference of -1.14% does not clear the 1.20% median noise floor of the processes
- multi u64 n=16384 valuesBetween: ordered vs hashed: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs hashed: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs hashed: 1 processes resolved A as faster and 11 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=1048576 valuesFor: ordered vs btree-sets: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=1048576 valuesFor: ordered vs hashed: the pooled difference of 0.63% does not clear the 1.67% median noise floor of the processes
- multi u64 n=1048576 valuesFor: ordered vs hashed: the pooled interval [-0.22%, 1.49%] includes zero
- multi u64 n=1048576 valuesFor: ordered vs hashed: 6 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=1048576 valuesFor: ordered vs map-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=1048576 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=1048576 churn: ordered vs hashed: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=1048576 churn: ordered vs map-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=4096 churn: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=4096 build: ordered vs hashed: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=16384 valuesBetween: ordered vs hashed: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=16384 valuesBetween: ordered vs map-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=262144 valuesFor: ordered vs hashed: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 valuesFor: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 valuesFor: ordered vs hashed: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 churn: ordered vs hashed: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi url n=1048576 churn: ordered vs map-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=16384 valuesBetween: ordered vs hashed: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=16384 valuesBetween: ordered vs map-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesFor: ordered vs btree-sets: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesFor: ordered vs map-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs hashed: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs map-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 valuesFor: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 valuesBetween: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 prefix: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=1048576 churn: ordered vs hashed: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesFor: ordered vs btree-map: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=1048576 churn: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=16384 churn: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs btree-map: 1 processes resolved A as faster and 24 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=262144 churn: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 churn: ordered vs btree-map: 1 processes resolved A as faster and 25 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique str n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 churn: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=1048576 valuesFor: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=1048576 churn: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=1048576 valuesBetween: ordered vs btree-map: the processes scatter 7.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=16384 prefix: ordered vs btree-map: the pooled difference of -0.97% does not clear the 1.98% median noise floor of the processes
- unique url n=16384 prefix: ordered vs btree-map: the pooled interval [-2.11%, 0.17%] includes zero
- unique url n=262144 valuesFor: ordered vs btree-map: the pooled difference of -1.77% does not clear the 1.94% median noise floor of the processes
- unique url n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs btree-map: 2 processes resolved A as faster and 13 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique url n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 churn: ordered vs btree-map: the processes scatter 7.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 churn: ordered vs btree-map: 71 processes resolved A as faster and 25 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique uuid n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 prefix: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=1048576 churn: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
