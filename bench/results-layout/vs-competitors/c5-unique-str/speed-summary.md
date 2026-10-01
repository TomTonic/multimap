| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | email | 4096 | valuesFor | ordered | btree-map | 6 | 22.9 | 98.1 | 4.28× [4.24, 4.32] | +76.6% | [+76.4%, +76.8%] | 0.2 pts | 1.0 | yes | yes |
| unique-str | email | 4096 | valuesBetween | ordered | btree-map | 6 | 1307 | 446 | 0.34× [0.34, 0.34] | -192.6% | [-193.2%, -192.1%] | 0.5 pts | 0.4 | yes | yes |
| unique-str | email | 4096 | prefix | ordered | btree-map | 6 | 57.6 | 98.9 | 1.72× [1.70, 1.73] | +41.8% | [+41.3%, +42.3%] | 0.5 pts | 1.6 | yes | yes |
| unique-str | email | 4096 | churn | ordered | btree-map | 6 | 76.7 | 147 | 1.91× [1.88, 1.94] | +47.6% | [+46.7%, +48.4%] | 0.8 pts | 1.6 | yes | yes |
| unique-str | email | 4096 | build | ordered | btree-map | 6 | 1.12 ms | 2.02 ms | 1.79× [1.78, 1.80] | +44.1% | [+43.7%, +44.5%] | 0.4 pts | 0.3 | yes | yes |
| unique-str | email | 16384 | valuesFor | ordered | btree-map | 6 | 38.4 | 135 | 3.53× [3.51, 3.55] | +71.7% | [+71.6%, +71.9%] | 0.1 pts | 1.1 | yes | yes |
| unique-str | email | 16384 | valuesBetween | ordered | btree-map | 6 | 1617 | 534 | 0.33× [0.33, 0.33] | -202.6% | [-205.7%, -199.4%] | 3.0 pts | 1.9 | yes | yes |
| unique-str | email | 16384 | prefix | ordered | btree-map | 6 | 72.0 | 136 | 1.89× [1.87, 1.92] | +47.2% | [+46.5%, +47.9%] | 0.7 pts | 2.1 | yes | yes |
| unique-str | email | 16384 | churn | ordered | btree-map | 6 | 98.6 | 205 | 2.11× [2.05, 2.18] | +52.6% | [+51.1%, +54.1%] | 1.4 pts | 2.6 | yes | yes |
| unique-str | email | 16384 | build | ordered | btree-map | 6 | 5.42 ms | 10.68 ms | 1.97× [1.96, 1.98] | +49.2% | [+49.0%, +49.4%] | 0.2 pts | 0.9 | yes | yes |
| unique-str | email | 262144 | valuesFor | ordered | btree-map | 12 | 184 | 260 | 1.43× [1.39, 1.47] | +30.0% | [+28.0%, +32.0%] | 1.9 pts | 2.4 | yes | yes |
| unique-str | email | 262144 | valuesBetween | ordered | btree-map | 12 | 3517 | 1670 | 0.48× [0.47, 0.49] | -108.9% | [-112.5%, -105.4%] | 7.4 pts | 0.9 | yes | yes |
| unique-str | email | 262144 | prefix | ordered | btree-map | 12 | 252 | 286 | 1.16× [1.13, 1.18] | +13.7% | [+11.8%, +15.6%] | 2.7 pts | 2.7 | yes | yes |
| unique-str | email | 262144 | churn | ordered | btree-map | 12 | 381 | 594 | 1.60× [1.56, 1.64] | +37.6% | [+36.0%, +39.2%] | 1.9 pts | 2.3 | yes | yes |
| unique-str | path | 4096 | valuesFor | ordered | btree-map | 6 | 79.4 | 114 | 1.44× [1.42, 1.47] | +30.7% | [+29.4%, +32.1%] | 1.3 pts | 1.0 | yes | yes |
| unique-str | path | 4096 | valuesBetween | ordered | btree-map | 6 | 2078 | 555 | 0.26× [0.26, 0.27] | -279.7% | [-285.9%, -273.5%] | 5.9 pts | 1.0 | yes | yes |
| unique-str | path | 4096 | prefix | ordered | btree-map | 6 | 235 | 142 | 0.61× [0.60, 0.62] | -63.9% | [-66.8%, -61.1%] | 2.7 pts | 0.8 | yes | yes |
| unique-str | path | 4096 | churn | ordered | btree-map | 6 | 205 | 184 | 0.90× [0.89, 0.91] | -10.7% | [-12.0%, -9.4%] | 1.2 pts | 0.8 | yes | yes |
| unique-str | path | 4096 | build | ordered | btree-map | 6 | 2.74 ms | 2.56 ms | 0.94× [0.93, 0.94] | -6.7% | [-7.4%, -5.9%] | 0.7 pts | 0.8 | yes | yes |
| unique-str | path | 16384 | valuesFor | ordered | btree-map | 6 | 114 | 165 | 1.45× [1.43, 1.46] | +30.8% | [+30.1%, +31.5%] | 0.7 pts | 0.9 | yes | yes |
| unique-str | path | 16384 | valuesBetween | ordered | btree-map | 6 | 2630 | 721 | 0.28× [0.27, 0.28] | -263.5% | [-268.9%, -258.2%] | 5.1 pts | 1.7 | yes | yes |
| unique-str | path | 16384 | prefix | ordered | btree-map | 6 | 487 | 260 | 0.54× [0.54, 0.55] | -83.5% | [-86.5%, -80.6%] | 2.8 pts | 0.7 | yes | yes |
| unique-str | path | 16384 | churn | ordered | btree-map | 6 | 265 | 251 | 0.95× [0.94, 0.97] | -5.0% | [-6.7%, -3.3%] | 1.6 pts | 1.4 | yes | yes |
| unique-str | path | 16384 | build | ordered | btree-map | 6 | 13.31 ms | 13.29 ms | 1.00× [0.99, 1.01] | +0.2% | [-0.9%, +1.4%] | 1.1 pts | 1.4 | yes | no |
| unique-str | path | 262144 | valuesFor | ordered | btree-map | 12 | 401 | 428 | 1.08× [1.05, 1.12] | +7.5% | [+4.4%, +10.5%] | 3.2 pts | 2.4 | no | yes |
| unique-str | path | 262144 | valuesBetween | ordered | btree-map | 12 | 6056 | 3044 | 0.51× [0.49, 0.52] | -97.7% | [-102.7%, -92.8%] | 6.0 pts | 2.1 | yes | yes |
| unique-str | path | 262144 | prefix | ordered | btree-map | 12 | 5756 | 2711 | 0.46× [0.44, 0.48] | -119.4% | [-129.1%, -109.8%] | 12.8 pts | 0.5 | yes | yes |
| unique-str | path | 262144 | churn | ordered | btree-map | 12 | 748 | 801 | 1.07× [1.06, 1.07] | +6.2% | [+5.5%, +6.8%] | 0.8 pts | 1.3 | yes | yes |
| unique-str | str | 4096 | valuesFor | ordered | btree-map | 6 | 35.7 | 99.1 | 2.79× [2.75, 2.82] | +64.1% | [+63.7%, +64.6%] | 0.4 pts | 0.9 | yes | yes |
| unique-str | str | 4096 | valuesBetween | ordered | btree-map | 6 | 1630 | 436 | 0.27× [0.27, 0.27] | -274.4% | [-276.7%, -272.1%] | 2.2 pts | 0.8 | yes | yes |
| unique-str | str | 4096 | prefix | ordered | btree-map | 6 | 3257 | 837 | 0.26× [0.26, 0.26] | -288.1% | [-290.2%, -285.9%] | 2.0 pts | 1.4 | yes | yes |
| unique-str | str | 4096 | churn | ordered | btree-map | 6 | 104 | 145 | 1.40× [1.38, 1.41] | +28.3% | [+27.6%, +29.1%] | 0.7 pts | 1.0 | yes | yes |
| unique-str | str | 4096 | build | ordered | btree-map | 6 | 1.48 ms | 1.98 ms | 1.33× [1.32, 1.35] | +25.1% | [+24.0%, +26.1%] | 1.0 pts | 1.2 | yes | yes |
| unique-str | str | 16384 | valuesFor | ordered | btree-map | 6 | 49.9 | 135 | 2.72× [2.70, 2.75] | +63.3% | [+62.9%, +63.6%] | 0.3 pts | 1.6 | yes | yes |
| unique-str | str | 16384 | valuesBetween | ordered | btree-map | 6 | 1685 | 525 | 0.31× [0.31, 0.31] | -222.7% | [-227.3%, -218.1%] | 4.4 pts | 2.8 | yes | yes |
| unique-str | str | 16384 | prefix | ordered | btree-map | 6 | 13.5 µs | 3648 | 0.27× [0.27, 0.27] | -272.3% | [-276.2%, -268.5%] | 3.7 pts | 1.6 | yes | yes |
| unique-str | str | 16384 | churn | ordered | btree-map | 6 | 125 | 204 | 1.65× [1.62, 1.68] | +39.3% | [+38.3%, +40.3%] | 0.9 pts | 1.7 | yes | yes |
| unique-str | str | 16384 | build | ordered | btree-map | 6 | 6.86 ms | 10.61 ms | 1.54× [1.53, 1.55] | +35.1% | [+34.7%, +35.5%] | 0.4 pts | 1.3 | yes | yes |
| unique-str | str | 262144 | valuesFor | ordered | btree-map | 12 | 200 | 262 | 1.32× [1.28, 1.36] | +24.4% | [+22.2%, +26.6%] | 2.7 pts | 2.3 | yes | yes |
| unique-str | str | 262144 | valuesBetween | ordered | btree-map | 12 | 3307 | 1556 | 0.47× [0.46, 0.49] | -111.5% | [-119.3%, -103.6%] | 7.8 pts | 0.7 | yes | yes |
| unique-str | str | 262144 | prefix | ordered | btree-map | 12 | 452.8 µs | 183.7 µs | 0.42× [0.41, 0.44] | -135.4% | [-145.5%, -125.3%] | 15.0 pts | 0.5 | yes | yes |
| unique-str | str | 262144 | churn | ordered | btree-map | 12 | 424 | 594 | 1.42× [1.38, 1.46] | +29.6% | [+27.6%, +31.5%] | 2.0 pts | 1.6 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | btree-map | 8 | 40.3 | 89.8 | 2.23× [2.19, 2.27] | +55.1% | [+54.3%, +55.9%] | 0.8 pts | 1.2 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | btree-map | 8 | 1755 | 493 | 0.28× [0.27, 0.28] | -259.1% | [-264.5%, -253.7%] | 5.1 pts | 1.5 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | btree-map | 8 | 199 | 131 | 0.66× [0.66, 0.67] | -51.1% | [-52.3%, -50.0%] | 1.3 pts | 0.7 | yes | yes |
| unique-str | street | 4096 | churn | ordered | btree-map | 8 | 121 | 142 | 1.17× [1.16, 1.19] | +14.7% | [+13.4%, +15.9%] | 1.8 pts | 1.8 | yes | yes |
| unique-str | street | 4096 | build | ordered | btree-map | 8 | 1.68 ms | 2.10 ms | 1.26× [1.23, 1.29] | +20.7% | [+18.9%, +22.5%] | 2.2 pts | 3.6 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | btree-map | 6 | 67.0 | 125 | 1.88× [1.82, 1.94] | +46.8% | [+45.1%, +48.4%] | 1.6 pts | 3.3 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | btree-map | 6 | 2177 | 574 | 0.26× [0.26, 0.27] | -277.7% | [-284.6%, -270.8%] | 6.6 pts | 2.7 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | btree-map | 6 | 627 | 272 | 0.43× [0.43, 0.44] | -130.0% | [-134.3%, -125.7%] | 4.1 pts | 1.5 | yes | yes |
| unique-str | street | 16384 | churn | ordered | btree-map | 6 | 158 | 200 | 1.27× [1.24, 1.29] | +21.1% | [+19.4%, +22.7%] | 1.6 pts | 2.3 | yes | yes |
| unique-str | street | 16384 | build | ordered | btree-map | 6 | 8.38 ms | 11.09 ms | 1.33× [1.31, 1.36] | +25.0% | [+23.6%, +26.3%] | 1.3 pts | 1.7 | yes | yes |
| unique-str | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 13.6 | 86.0 | 6.32× [6.27, 6.37] | +84.2% | [+84.0%, +84.3%] | 0.1 pts | 1.0 | yes | yes |
| unique-str | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 1090 | 426 | 0.39× [0.39, 0.39] | -155.8% | [-156.8%, -154.9%] | 0.9 pts | 1.1 | yes | yes |
| unique-str | u64 | 4096 | churn | ordered | btree-map | 6 | 50.9 | 126 | 2.48× [2.42, 2.53] | +59.6% | [+58.7%, +60.5%] | 0.9 pts | 2.5 | yes | yes |
| unique-str | u64 | 4096 | build | ordered | btree-map | 6 | 834.9 µs | 1.79 ms | 2.15× [2.11, 2.19] | +53.5% | [+52.6%, +54.3%] | 0.8 pts | 2.4 | yes | yes |
| unique-str | u64 | 16384 | valuesFor | ordered | btree-map | 6 | 23.3 | 115 | 4.98× [4.95, 5.00] | +79.9% | [+79.8%, +80.0%] | 0.1 pts | 1.2 | yes | yes |
| unique-str | u64 | 16384 | valuesBetween | ordered | btree-map | 6 | 1569 | 460 | 0.29× [0.29, 0.30] | -240.1% | [-242.2%, -238.1%] | 1.9 pts | 1.1 | yes | yes |
| unique-str | u64 | 16384 | churn | ordered | btree-map | 6 | 60.8 | 173 | 2.87× [2.81, 2.94] | +65.2% | [+64.4%, +66.0%] | 0.8 pts | 2.1 | yes | yes |
| unique-str | u64 | 16384 | build | ordered | btree-map | 6 | 3.81 ms | 9.21 ms | 2.41× [2.40, 2.43] | +58.6% | [+58.3%, +58.9%] | 0.3 pts | 1.4 | yes | yes |
| unique-str | u64 | 262144 | valuesFor | ordered | btree-map | 6 | 75.7 | 218 | 2.88× [2.76, 3.02] | +65.3% | [+63.7%, +66.9%] | 1.5 pts | 2.3 | yes | yes |
| unique-str | u64 | 262144 | valuesBetween | ordered | btree-map | 6 | 1711 | 837 | 0.49× [0.47, 0.52] | -102.3% | [-111.3%, -93.3%] | 8.6 pts | 1.6 | yes | yes |
| unique-str | u64 | 262144 | churn | ordered | btree-map | 6 | 271 | 399 | 1.51× [1.48, 1.55] | +33.8% | [+32.3%, +35.3%] | 1.4 pts | 1.2 | yes | yes |
| unique-str | url | 4096 | valuesFor | ordered | btree-map | 12 | 63.1 | 113 | 1.82× [1.79, 1.85] | +45.1% | [+44.1%, +46.0%] | 1.2 pts | 1.4 | yes | yes |
| unique-str | url | 4096 | valuesBetween | ordered | btree-map | 12 | 2053 | 547 | 0.26× [0.26, 0.26] | -281.6% | [-284.9%, -278.4%] | 7.2 pts | 1.2 | yes | yes |
| unique-str | url | 4096 | prefix | ordered | btree-map | 12 | 147 | 126 | 0.86× [0.85, 0.88] | -16.1% | [-18.0%, -14.1%] | 1.9 pts | 1.0 | yes | yes |
| unique-str | url | 4096 | churn | ordered | btree-map | 12 | 167 | 182 | 1.09× [1.08, 1.10] | +8.3% | [+7.2%, +9.4%] | 1.2 pts | 1.0 | yes | yes |
| unique-str | url | 4096 | build | ordered | btree-map | 12 | 2.32 ms | 2.45 ms | 1.06× [1.05, 1.07] | +5.6% | [+4.4%, +6.7%] | 1.6 pts | 1.2 | yes | yes |
| unique-str | url | 16384 | valuesFor | ordered | btree-map | 6 | 90.9 | 165 | 1.82× [1.81, 1.82] | +44.9% | [+44.8%, +45.1%] | 0.2 pts | 0.4 | yes | yes |
| unique-str | url | 16384 | valuesBetween | ordered | btree-map | 6 | 2646 | 724 | 0.27× [0.27, 0.28] | -267.2% | [-271.3%, -263.1%] | 3.9 pts | 1.3 | yes | yes |
| unique-str | url | 16384 | prefix | ordered | btree-map | 6 | 218 | 194 | 0.88× [0.88, 0.89] | -13.1% | [-14.2%, -12.1%] | 1.0 pts | 0.9 | yes | yes |
| unique-str | url | 16384 | churn | ordered | btree-map | 6 | 215 | 247 | 1.16× [1.13, 1.18] | +13.5% | [+11.6%, +15.4%] | 1.8 pts | 2.0 | yes | yes |
| unique-str | url | 16384 | build | ordered | btree-map | 6 | 10.96 ms | 13.00 ms | 1.19× [1.18, 1.20] | +15.9% | [+15.2%, +16.6%] | 0.7 pts | 1.2 | yes | yes |
| unique-str | url | 262144 | valuesFor | ordered | btree-map | 8 | 378 | 422 | 1.12× [1.09, 1.14] | +10.3% | [+8.5%, +12.1%] | 1.8 pts | 1.8 | yes | yes |
| unique-str | url | 262144 | valuesBetween | ordered | btree-map | 8 | 6852 | 2969 | 0.43× [0.42, 0.45] | -130.3% | [-137.1%, -123.6%] | 6.4 pts | 2.2 | yes | yes |
| unique-str | url | 262144 | prefix | ordered | btree-map | 8 | 1429 | 870 | 0.62× [0.61, 0.64] | -60.9% | [-64.7%, -57.2%] | 4.1 pts | 1.7 | yes | yes |
| unique-str | url | 262144 | churn | ordered | btree-map | 8 | 678 | 788 | 1.15× [1.14, 1.17] | +13.4% | [+12.0%, +14.7%] | 1.3 pts | 1.5 | yes | yes |
| unique-str | uuid | 4096 | valuesFor | ordered | btree-map | 6 | 28.7 | 99.1 | 3.43× [3.34, 3.53] | +70.9% | [+70.1%, +71.7%] | 0.8 pts | 1.5 | yes | yes |
| unique-str | uuid | 4096 | valuesBetween | ordered | btree-map | 6 | 1465 | 434 | 0.30× [0.29, 0.30] | -238.2% | [-240.5%, -236.0%] | 2.1 pts | 1.1 | yes | yes |
| unique-str | uuid | 4096 | prefix | ordered | btree-map | 6 | 69.1 | 99.9 | 1.45× [1.44, 1.47] | +31.2% | [+30.6%, +31.8%] | 0.6 pts | 1.1 | yes | yes |
| unique-str | uuid | 4096 | churn | ordered | btree-map | 6 | 87.5 | 141 | 1.61× [1.58, 1.64] | +37.9% | [+36.8%, +39.0%] | 1.1 pts | 1.9 | yes | yes |
| unique-str | uuid | 4096 | build | ordered | btree-map | 6 | 1.23 ms | 1.97 ms | 1.59× [1.58, 1.61] | +37.3% | [+36.5%, +38.1%] | 0.7 pts | 0.9 | yes | yes |
| unique-str | uuid | 16384 | valuesFor | ordered | btree-map | 6 | 42.8 | 138 | 3.23× [3.19, 3.26] | +69.0% | [+68.7%, +69.3%] | 0.3 pts | 2.1 | yes | yes |
| unique-str | uuid | 16384 | valuesBetween | ordered | btree-map | 6 | 1618 | 546 | 0.34× [0.33, 0.34] | -196.5% | [-198.5%, -194.4%] | 2.0 pts | 1.6 | yes | yes |
| unique-str | uuid | 16384 | prefix | ordered | btree-map | 6 | 83.8 | 144 | 1.72× [1.71, 1.73] | +41.8% | [+41.6%, +42.0%] | 0.2 pts | 0.9 | yes | yes |
| unique-str | uuid | 16384 | churn | ordered | btree-map | 6 | 108 | 197 | 1.83× [1.79, 1.87] | +45.4% | [+44.2%, +46.6%] | 1.2 pts | 1.7 | yes | yes |
| unique-str | uuid | 16384 | build | ordered | btree-map | 6 | 5.70 ms | 10.63 ms | 1.87× [1.85, 1.88] | +46.4% | [+46.0%, +46.8%] | 0.3 pts | 1.3 | yes | yes |
| unique-str | uuid | 262144 | valuesFor | ordered | btree-map | 8 | 199 | 298 | 1.54× [1.49, 1.61] | +35.2% | [+32.7%, +37.8%] | 2.5 pts | 2.7 | yes | yes |
| unique-str | uuid | 262144 | valuesBetween | ordered | btree-map | 8 | 4003 | 2100 | 0.53× [0.52, 0.54] | -88.8% | [-91.9%, -85.8%] | 3.3 pts | 0.7 | yes | yes |
| unique-str | uuid | 262144 | prefix | ordered | btree-map | 8 | 353 | 419 | 1.20× [1.17, 1.22] | +16.6% | [+14.8%, +18.4%] | 1.7 pts | 1.2 | yes | yes |
| unique-str | uuid | 262144 | churn | ordered | btree-map | 8 | 393 | 624 | 1.61× [1.58, 1.64] | +37.8% | [+36.5%, +39.2%] | 1.6 pts | 1.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str email n=16384 prefix: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=16384 churn: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.80% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str email n=262144 prefix: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str email n=262144 churn: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=16384 build: ordered vs btree-map: the pooled difference of 0.23% does not clear the 0.45% noise floor, the bound on what the harness reports between identical code in every process
- unique-str path n=16384 build: ordered vs btree-map: the pooled interval [-0.89%, 1.35%] includes zero
- unique-str path n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str path n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str str n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str str n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str str n=262144 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str str n=262144 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -5.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=4096 build: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 churn: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=4096 churn: ordered vs btree-map: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of +0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str u64 n=4096 build: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=16384 churn: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str u64 n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str url n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str uuid n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
