| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique | email | 4096 | valuesFor | ordered | btree-map | 6 | 23.3 | 97.4 | 4.20× [4.17, 4.24] | +76.2% | [+76.0%, +76.4%] | 0.2 pts | 0.7 | yes | yes |
| unique | email | 4096 | valuesBetween | ordered | btree-map | 6 | 1211 | 440 | 0.36× [0.36, 0.36] | -175.6% | [-176.6%, -174.7%] | 0.9 pts | 0.5 | yes | yes |
| unique | email | 4096 | prefix | ordered | btree-map | 6 | 58.4 | 97.7 | 1.67× [1.66, 1.69] | +40.3% | [+39.9%, +40.7%] | 0.4 pts | 1.1 | yes | yes |
| unique | email | 4096 | churn | ordered | btree-map | 6 | 72.5 | 141 | 1.96× [1.94, 1.98] | +48.9% | [+48.4%, +49.5%] | 0.5 pts | 1.2 | yes | yes |
| unique | email | 4096 | build | ordered | btree-map | 6 | 1.05 ms | 1.95 ms | 1.86× [1.85, 1.87] | +46.2% | [+45.8%, +46.6%] | 0.4 pts | 0.5 | yes | yes |
| unique | email | 16384 | valuesFor | ordered | btree-map | 6 | 38.4 | 133 | 3.47× [3.44, 3.50] | +71.2% | [+71.0%, +71.5%] | 0.2 pts | 1.6 | yes | yes |
| unique | email | 16384 | valuesBetween | ordered | btree-map | 6 | 1577 | 519 | 0.33× [0.33, 0.33] | -204.1% | [-206.4%, -201.9%] | 2.1 pts | 1.3 | yes | yes |
| unique | email | 16384 | prefix | ordered | btree-map | 6 | 71.8 | 135 | 1.88× [1.86, 1.90] | +46.8% | [+46.3%, +47.3%] | 0.5 pts | 1.8 | yes | yes |
| unique | email | 16384 | churn | ordered | btree-map | 6 | 92.5 | 198 | 2.15× [2.08, 2.23] | +53.5% | [+51.9%, +55.1%] | 1.5 pts | 2.9 | yes | yes |
| unique | email | 16384 | build | ordered | btree-map | 6 | 5.16 ms | 10.33 ms | 2.00× [1.98, 2.02] | +49.9% | [+49.4%, +50.4%] | 0.5 pts | 1.7 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | btree-map | 6 | 173 | 248 | 1.45× [1.41, 1.49] | +31.1% | [+29.3%, +32.8%] | 1.7 pts | 1.4 | yes | yes |
| unique | email | 262144 | valuesBetween | ordered | btree-map | 6 | 3234 | 1560 | 0.50× [0.49, 0.51] | -100.7% | [-104.0%, -97.4%] | 3.1 pts | 0.3 | yes | yes |
| unique | email | 262144 | prefix | ordered | btree-map | 6 | 229 | 274 | 1.20× [1.17, 1.22] | +16.3% | [+14.6%, +18.1%] | 1.7 pts | 2.9 | yes | yes |
| unique | email | 262144 | churn | ordered | btree-map | 6 | 366 | 562 | 1.55× [1.52, 1.57] | +35.3% | [+34.2%, +36.4%] | 1.1 pts | 1.1 | yes | yes |
| unique | path | 4096 | valuesFor | ordered | btree-map | 6 | 80.0 | 114 | 1.44× [1.42, 1.46] | +30.6% | [+29.5%, +31.7%] | 1.0 pts | 0.8 | yes | yes |
| unique | path | 4096 | valuesBetween | ordered | btree-map | 6 | 1946 | 557 | 0.28× [0.28, 0.29] | -253.7% | [-259.6%, -247.8%] | 5.6 pts | 1.2 | yes | yes |
| unique | path | 4096 | prefix | ordered | btree-map | 6 | 231 | 143 | 0.61× [0.59, 0.62] | -64.7% | [-68.4%, -60.9%] | 3.6 pts | 1.1 | yes | yes |
| unique | path | 4096 | churn | ordered | btree-map | 6 | 201 | 178 | 0.87× [0.87, 0.88] | -14.3% | [-15.4%, -13.2%] | 1.0 pts | 0.6 | yes | yes |
| unique | path | 4096 | build | ordered | btree-map | 6 | 2.72 ms | 2.53 ms | 0.93× [0.92, 0.94] | -8.0% | [-9.2%, -6.9%] | 1.1 pts | 1.2 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | btree-map | 6 | 114 | 165 | 1.44× [1.40, 1.47] | +30.4% | [+28.7%, +32.1%] | 1.6 pts | 1.9 | yes | yes |
| unique | path | 16384 | valuesBetween | ordered | btree-map | 6 | 2517 | 720 | 0.29× [0.28, 0.30] | -249.4% | [-259.9%, -239.0%] | 10.0 pts | 4.3 | yes | yes |
| unique | path | 16384 | prefix | ordered | btree-map | 6 | 476 | 263 | 0.56× [0.54, 0.57] | -79.6% | [-83.6%, -75.5%] | 3.9 pts | 1.4 | yes | yes |
| unique | path | 16384 | churn | ordered | btree-map | 6 | 257 | 244 | 0.95× [0.94, 0.96] | -5.2% | [-6.2%, -4.3%] | 0.9 pts | 0.8 | yes | yes |
| unique | path | 16384 | build | ordered | btree-map | 6 | 13.25 ms | 13.00 ms | 0.98× [0.97, 0.99] | -2.0% | [-2.8%, -1.2%] | 0.8 pts | 1.0 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | btree-map | 12 | 384 | 412 | 1.08× [1.06, 1.11] | +7.7% | [+5.4%, +10.0%] | 3.2 pts | 2.4 | no | yes |
| unique | path | 262144 | valuesBetween | ordered | btree-map | 12 | 5732 | 2887 | 0.51× [0.50, 0.52] | -96.6% | [-100.5%, -92.8%] | 5.2 pts | 1.6 | yes | yes |
| unique | path | 262144 | prefix | ordered | btree-map | 12 | 5122 | 2376 | 0.45× [0.42, 0.47] | -123.7% | [-135.5%, -111.9%] | 15.2 pts | 0.7 | yes | yes |
| unique | path | 262144 | churn | ordered | btree-map | 12 | 733 | 740 | 1.02× [1.00, 1.04] | +1.6% | [-0.2%, +3.4%] | 1.7 pts | 1.1 | yes | no |
| unique | str | 4096 | valuesFor | ordered | btree-map | 6 | 35.8 | 97.3 | 2.77× [2.70, 2.84] | +63.9% | [+63.0%, +64.8%] | 0.8 pts | 1.3 | yes | yes |
| unique | str | 4096 | valuesBetween | ordered | btree-map | 6 | 1566 | 437 | 0.28× [0.27, 0.28] | -259.1% | [-263.8%, -254.5%] | 4.4 pts | 1.7 | yes | yes |
| unique | str | 4096 | prefix | ordered | btree-map | 6 | 3070 | 849 | 0.27× [0.27, 0.28] | -265.3% | [-268.5%, -262.1%] | 3.1 pts | 1.3 | yes | yes |
| unique | str | 4096 | churn | ordered | btree-map | 6 | 96.8 | 140 | 1.45× [1.44, 1.46] | +30.9% | [+30.4%, +31.3%] | 0.4 pts | 0.7 | yes | yes |
| unique | str | 4096 | build | ordered | btree-map | 6 | 1.38 ms | 1.93 ms | 1.40× [1.38, 1.41] | +28.4% | [+27.8%, +28.9%] | 0.6 pts | 0.5 | yes | yes |
| unique | str | 16384 | valuesFor | ordered | btree-map | 6 | 50.0 | 133 | 2.66× [2.64, 2.68] | +62.4% | [+62.1%, +62.7%] | 0.3 pts | 1.2 | yes | yes |
| unique | str | 16384 | valuesBetween | ordered | btree-map | 6 | 1609 | 512 | 0.32× [0.31, 0.32] | -214.3% | [-218.1%, -210.6%] | 3.6 pts | 1.9 | yes | yes |
| unique | str | 16384 | prefix | ordered | btree-map | 6 | 13.1 µs | 3478 | 0.27× [0.26, 0.27] | -275.0% | [-279.2%, -270.8%] | 4.0 pts | 2.0 | yes | yes |
| unique | str | 16384 | churn | ordered | btree-map | 6 | 117 | 197 | 1.69× [1.67, 1.71] | +40.9% | [+40.3%, +41.5%] | 0.6 pts | 1.0 | yes | yes |
| unique | str | 16384 | build | ordered | btree-map | 6 | 6.51 ms | 10.25 ms | 1.57× [1.56, 1.58] | +36.3% | [+35.9%, +36.7%] | 0.4 pts | 1.0 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | btree-map | 12 | 174 | 244 | 1.38× [1.33, 1.44] | +27.7% | [+24.8%, +30.7%] | 3.6 pts | 3.4 | no | yes |
| unique | str | 262144 | valuesBetween | ordered | btree-map | 12 | 2810 | 1356 | 0.50× [0.47, 0.52] | -101.6% | [-112.1%, -91.2%] | 11.4 pts | 1.0 | no | yes |
| unique | str | 262144 | prefix | ordered | btree-map | 12 | 377.7 µs | 154.0 µs | 0.44× [0.43, 0.45] | -126.7% | [-131.0%, -122.4%] | 18.9 pts | 0.5 | yes | yes |
| unique | str | 262144 | churn | ordered | btree-map | 12 | 395 | 543 | 1.42× [1.40, 1.45] | +29.7% | [+28.5%, +30.8%] | 1.7 pts | 1.2 | yes | yes |
| unique | street | 4096 | valuesFor | ordered | btree-map | 6 | 41.9 | 89.3 | 2.12× [2.07, 2.18] | +52.9% | [+51.8%, +54.1%] | 1.1 pts | 1.8 | yes | yes |
| unique | street | 4096 | valuesBetween | ordered | btree-map | 6 | 1703 | 496 | 0.29× [0.28, 0.30] | -246.4% | [-255.9%, -236.9%] | 9.1 pts | 2.6 | yes | yes |
| unique | street | 4096 | prefix | ordered | btree-map | 6 | 192 | 132 | 0.68× [0.66, 0.70] | -46.8% | [-50.9%, -42.7%] | 3.9 pts | 2.5 | yes | yes |
| unique | street | 4096 | churn | ordered | btree-map | 6 | 115 | 138 | 1.19× [1.17, 1.21] | +16.2% | [+14.8%, +17.7%] | 1.4 pts | 1.4 | yes | yes |
| unique | street | 4096 | build | ordered | btree-map | 6 | 1.62 ms | 2.08 ms | 1.29× [1.27, 1.31] | +22.3% | [+21.0%, +23.5%] | 1.2 pts | 2.1 | yes | yes |
| unique | street | 16384 | valuesFor | ordered | btree-map | 6 | 67.7 | 122 | 1.80× [1.78, 1.82] | +44.4% | [+43.9%, +45.0%] | 0.6 pts | 1.1 | yes | yes |
| unique | street | 16384 | valuesBetween | ordered | btree-map | 6 | 2131 | 574 | 0.27× [0.27, 0.27] | -273.0% | [-277.3%, -268.7%] | 4.1 pts | 1.5 | yes | yes |
| unique | street | 16384 | prefix | ordered | btree-map | 6 | 601 | 273 | 0.45× [0.44, 0.45] | -123.5% | [-126.6%, -120.5%] | 2.9 pts | 0.9 | yes | yes |
| unique | street | 16384 | churn | ordered | btree-map | 6 | 152 | 191 | 1.26× [1.24, 1.29] | +20.9% | [+19.5%, +22.4%] | 1.4 pts | 1.9 | yes | yes |
| unique | street | 16384 | build | ordered | btree-map | 6 | 8.23 ms | 10.78 ms | 1.31× [1.29, 1.34] | +23.9% | [+22.5%, +25.2%] | 1.3 pts | 1.5 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 17.8 | 84.3 | 4.74× [4.69, 4.80] | +78.9% | [+78.7%, +79.1%] | 0.2 pts | 3.3 | yes | yes |
| unique | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 239 | 431 | 1.80× [1.78, 1.82] | +44.4% | [+43.7%, +45.0%] | 0.6 pts | 2.5 | yes | yes |
| unique | u64 | 4096 | churn | ordered | btree-map | 6 | 37.0 | 121 | 3.29× [3.26, 3.32] | +69.6% | [+69.3%, +69.9%] | 0.3 pts | 2.0 | yes | yes |
| unique | u64 | 4096 | build | ordered | btree-map | 6 | 587.0 µs | 1.71 ms | 2.90× [2.87, 2.94] | +65.6% | [+65.1%, +66.0%] | 0.4 pts | 1.8 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | btree-map | 6 | 26.5 | 113 | 4.24× [4.21, 4.28] | +76.4% | [+76.2%, +76.6%] | 0.2 pts | 2.5 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | btree-map | 6 | 284 | 463 | 1.63× [1.62, 1.65] | +38.8% | [+38.3%, +39.3%] | 0.5 pts | 1.3 | yes | yes |
| unique | u64 | 16384 | churn | ordered | btree-map | 6 | 54.3 | 166 | 3.07× [3.04, 3.10] | +67.4% | [+67.1%, +67.8%] | 0.3 pts | 1.8 | yes | yes |
| unique | u64 | 16384 | build | ordered | btree-map | 6 | 3.34 ms | 8.79 ms | 2.63× [2.61, 2.65] | +62.0% | [+61.6%, +62.3%] | 0.3 pts | 1.2 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | btree-map | 6 | 39.4 | 211 | 5.44× [5.38, 5.50] | +81.6% | [+81.4%, +81.8%] | 0.2 pts | 0.6 | yes | yes |
| unique | u64 | 262144 | valuesBetween | ordered | btree-map | 6 | 286 | 700 | 2.44× [2.43, 2.46] | +59.0% | [+58.8%, +59.3%] | 0.2 pts | 0.6 | yes | yes |
| unique | u64 | 262144 | churn | ordered | btree-map | 6 | 179 | 401 | 2.23× [2.19, 2.27] | +55.1% | [+54.4%, +55.9%] | 0.7 pts | 0.7 | yes | yes |
| unique | url | 4096 | valuesFor | ordered | btree-map | 8 | 63.4 | 113 | 1.79× [1.74, 1.84] | +44.1% | [+42.6%, +45.6%] | 1.5 pts | 1.9 | yes | yes |
| unique | url | 4096 | valuesBetween | ordered | btree-map | 8 | 1836 | 551 | 0.30× [0.29, 0.30] | -236.1% | [-239.7%, -232.5%] | 4.6 pts | 1.0 | yes | yes |
| unique | url | 4096 | prefix | ordered | btree-map | 8 | 143 | 125 | 0.89× [0.87, 0.90] | -13.0% | [-14.9%, -11.0%] | 1.9 pts | 1.0 | yes | yes |
| unique | url | 4096 | churn | ordered | btree-map | 8 | 161 | 176 | 1.08× [1.06, 1.09] | +7.2% | [+6.0%, +8.5%] | 1.2 pts | 1.1 | yes | yes |
| unique | url | 4096 | build | ordered | btree-map | 8 | 2.23 ms | 2.39 ms | 1.08× [1.07, 1.09] | +7.4% | [+6.4%, +8.4%] | 0.9 pts | 0.8 | yes | yes |
| unique | url | 16384 | valuesFor | ordered | btree-map | 6 | 90.5 | 163 | 1.82× [1.80, 1.84] | +45.1% | [+44.4%, +45.7%] | 0.7 pts | 1.7 | yes | yes |
| unique | url | 16384 | valuesBetween | ordered | btree-map | 6 | 2374 | 719 | 0.30× [0.30, 0.31] | -231.1% | [-238.2%, -224.0%] | 6.8 pts | 2.3 | yes | yes |
| unique | url | 16384 | prefix | ordered | btree-map | 6 | 212 | 196 | 0.92× [0.91, 0.92] | -9.0% | [-9.8%, -8.2%] | 0.7 pts | 0.7 | yes | yes |
| unique | url | 16384 | churn | ordered | btree-map | 6 | 207 | 239 | 1.15× [1.13, 1.18] | +13.2% | [+11.4%, +15.0%] | 1.7 pts | 1.7 | yes | yes |
| unique | url | 16384 | build | ordered | btree-map | 6 | 10.66 ms | 12.53 ms | 1.17× [1.17, 1.18] | +14.8% | [+14.4%, +15.1%] | 0.4 pts | 0.5 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | btree-map | 6 | 361 | 399 | 1.11× [1.09, 1.13] | +9.9% | [+8.5%, +11.3%] | 1.3 pts | 1.3 | yes | yes |
| unique | url | 262144 | valuesBetween | ordered | btree-map | 6 | 6060 | 2788 | 0.46× [0.44, 0.47] | -119.3% | [-126.2%, -112.4%] | 6.6 pts | 1.9 | yes | yes |
| unique | url | 262144 | prefix | ordered | btree-map | 6 | 1335 | 807 | 0.60× [0.59, 0.61] | -65.6% | [-68.5%, -62.7%] | 2.8 pts | 1.4 | yes | yes |
| unique | url | 262144 | churn | ordered | btree-map | 6 | 657 | 732 | 1.11× [1.09, 1.14] | +10.2% | [+8.4%, +12.1%] | 1.8 pts | 2.6 | yes | yes |
| unique | uuid | 4096 | valuesFor | ordered | btree-map | 6 | 28.9 | 97.0 | 3.18× [2.76, 3.74] | +68.5% | [+63.8%, +73.3%] | 4.5 pts | 14.6 | yes | yes |
| unique | uuid | 4096 | valuesBetween | ordered | btree-map | 6 | 1383 | 433 | 0.31× [0.31, 0.32] | -221.2% | [-226.0%, -216.4%] | 4.5 pts | 2.1 | yes | yes |
| unique | uuid | 4096 | prefix | ordered | btree-map | 6 | 69.3 | 98.9 | 1.43× [1.43, 1.44] | +30.2% | [+30.0%, +30.4%] | 0.2 pts | 0.4 | yes | yes |
| unique | uuid | 4096 | churn | ordered | btree-map | 6 | 83.7 | 138 | 1.64× [1.61, 1.66] | +38.9% | [+37.9%, +39.9%] | 1.0 pts | 1.7 | yes | yes |
| unique | uuid | 4096 | build | ordered | btree-map | 6 | 1.17 ms | 1.93 ms | 1.62× [1.60, 1.65] | +38.4% | [+37.4%, +39.4%] | 1.0 pts | 1.3 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | btree-map | 6 | 43.0 | 135 | 3.15× [3.13, 3.18] | +68.3% | [+68.0%, +68.5%] | 0.2 pts | 1.5 | yes | yes |
| unique | uuid | 16384 | valuesBetween | ordered | btree-map | 6 | 1561 | 536 | 0.34× [0.33, 0.35] | -194.0% | [-200.3%, -187.6%] | 6.1 pts | 4.1 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | btree-map | 6 | 84.0 | 143 | 1.70× [1.69, 1.71] | +41.1% | [+40.8%, +41.5%] | 0.3 pts | 1.4 | yes | yes |
| unique | uuid | 16384 | churn | ordered | btree-map | 6 | 104 | 192 | 1.86× [1.81, 1.92] | +46.3% | [+44.7%, +47.9%] | 1.5 pts | 1.9 | yes | yes |
| unique | uuid | 16384 | build | ordered | btree-map | 6 | 5.54 ms | 10.28 ms | 1.85× [1.83, 1.86] | +45.9% | [+45.5%, +46.3%] | 0.4 pts | 1.3 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | btree-map | 10 | 205 | 291 | 1.40× [1.37, 1.43] | +28.7% | [+27.2%, +30.2%] | 1.5 pts | 1.8 | yes | yes |
| unique | uuid | 262144 | valuesBetween | ordered | btree-map | 10 | 3927 | 2038 | 0.53× [0.51, 0.55] | -89.0% | [-95.6%, -82.4%] | 6.6 pts | 1.2 | yes | yes |
| unique | uuid | 262144 | prefix | ordered | btree-map | 10 | 372 | 404 | 1.10× [1.07, 1.12] | +8.7% | [+6.8%, +10.6%] | 2.8 pts | 2.4 | yes | yes |
| unique | uuid | 262144 | churn | ordered | btree-map | 10 | 363 | 558 | 1.56× [1.54, 1.58] | +35.9% | [+35.1%, +36.7%] | 1.0 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique email n=16384 churn: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 prefix: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=4096 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -1.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique path n=16384 valuesBetween: ordered vs btree-map: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 churn: ordered vs btree-map: the pooled interval [-0.17%, 3.40%] includes zero
- unique str n=4096 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of +0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique str n=16384 prefix: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesFor: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.82% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique str n=262144 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -10.60% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=4096 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.33% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 prefix: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=4096 build: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of +0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 churn: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=262144 churn: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 valuesFor: ordered vs btree-map: the processes scatter 14.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.24% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique uuid n=16384 valuesBetween: ordered vs btree-map: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 prefix: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
