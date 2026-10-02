| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique | email | 4096 | valuesFor | ordered | btree-map | 6 | 34.0 | 100 | 2.93× [2.90, 2.96] | +65.9% | [+65.6%, +66.2%] | 0.3 pts | 3.2 | yes | yes |
| unique | email | 4096 | valuesBetween | ordered | btree-map | 6 | 349 | 435 | 1.25× [1.24, 1.26] | +20.0% | [+19.3%, +20.6%] | 0.6 pts | 1.8 | yes | yes |
| unique | email | 4096 | prefix | ordered | btree-map | 6 | 59.4 | 99.3 | 1.66× [1.65, 1.68] | +39.9% | [+39.3%, +40.5%] | 0.6 pts | 2.0 | yes | yes |
| unique | email | 4096 | churn | ordered | btree-map | 6 | 87.0 | 143 | 1.64× [1.61, 1.66] | +38.8% | [+37.9%, +39.8%] | 1.0 pts | 2.0 | yes | yes |
| unique | email | 4096 | build | ordered | btree-map | 6 | 1.47 ms | 1.99 ms | 1.36× [1.35, 1.37] | +26.5% | [+25.9%, +27.1%] | 0.6 pts | 1.3 | yes | yes |
| unique | email | 16384 | valuesFor | ordered | btree-map | 6 | 47.3 | 135 | 2.85× [2.79, 2.90] | +64.9% | [+64.2%, +65.6%] | 0.7 pts | 3.7 | yes | yes |
| unique | email | 16384 | valuesBetween | ordered | btree-map | 6 | 442 | 546 | 1.22× [1.20, 1.25] | +18.3% | [+16.5%, +20.1%] | 1.7 pts | 5.5 | yes | yes |
| unique | email | 16384 | prefix | ordered | btree-map | 6 | 80.2 | 138 | 1.72× [1.69, 1.75] | +41.7% | [+40.7%, +42.8%] | 1.0 pts | 3.6 | yes | yes |
| unique | email | 16384 | churn | ordered | btree-map | 6 | 111 | 200 | 1.79× [1.75, 1.83] | +44.2% | [+43.0%, +45.4%] | 1.1 pts | 2.4 | yes | yes |
| unique | email | 16384 | build | ordered | btree-map | 6 | 7.12 ms | 10.76 ms | 1.50× [1.48, 1.52] | +33.5% | [+32.6%, +34.3%] | 0.8 pts | 1.5 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | btree-map | 12 | 178 | 256 | 1.46× [1.44, 1.49] | +31.7% | [+30.4%, +33.0%] | 2.3 pts | 2.2 | yes | yes |
| unique | email | 262144 | valuesBetween | ordered | btree-map | 12 | 1023 | 1296 | 1.27× [1.22, 1.33] | +21.5% | [+18.4%, +24.7%] | 4.4 pts | 0.8 | no | yes |
| unique | email | 262144 | prefix | ordered | btree-map | 12 | 215 | 300 | 1.44× [1.42, 1.47] | +30.8% | [+29.4%, +32.2%] | 1.8 pts | 1.6 | yes | yes |
| unique | email | 262144 | churn | ordered | btree-map | 12 | 428 | 584 | 1.37× [1.35, 1.39] | +27.1% | [+26.1%, +28.1%] | 1.4 pts | 1.2 | yes | yes |
| unique | path | 4096 | valuesFor | ordered | btree-map | 6 | 97.3 | 111 | 1.15× [1.13, 1.16] | +12.7% | [+11.5%, +13.9%] | 1.1 pts | 1.3 | yes | yes |
| unique | path | 4096 | valuesBetween | ordered | btree-map | 6 | 777 | 552 | 0.70× [0.68, 0.72] | -42.9% | [-46.3%, -39.4%] | 3.3 pts | 1.5 | yes | yes |
| unique | path | 4096 | prefix | ordered | btree-map | 6 | 255 | 142 | 0.56× [0.54, 0.57] | -79.1% | [-84.1%, -74.1%] | 4.8 pts | 1.7 | yes | yes |
| unique | path | 4096 | churn | ordered | btree-map | 6 | 199 | 179 | 0.90× [0.88, 0.91] | -11.5% | [-13.2%, -9.9%] | 1.6 pts | 1.5 | yes | yes |
| unique | path | 4096 | build | ordered | btree-map | 6 | 3.63 ms | 2.56 ms | 0.70× [0.69, 0.71] | -42.1% | [-44.2%, -40.0%] | 2.0 pts | 1.9 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | btree-map | 6 | 125 | 161 | 1.28× [1.26, 1.31] | +22.1% | [+20.5%, +23.6%] | 1.5 pts | 1.9 | yes | yes |
| unique | path | 16384 | valuesBetween | ordered | btree-map | 6 | 962 | 716 | 0.74× [0.74, 0.75] | -34.8% | [-35.6%, -34.0%] | 0.8 pts | 0.6 | yes | yes |
| unique | path | 16384 | prefix | ordered | btree-map | 6 | 360 | 258 | 0.72× [0.71, 0.73] | -38.6% | [-40.2%, -37.1%] | 1.5 pts | 0.7 | yes | yes |
| unique | path | 16384 | churn | ordered | btree-map | 6 | 259 | 249 | 0.96× [0.95, 0.97] | -4.0% | [-5.4%, -2.6%] | 1.3 pts | 1.3 | yes | yes |
| unique | path | 16384 | build | ordered | btree-map | 6 | 17.19 ms | 13.32 ms | 0.78× [0.77, 0.79] | -28.5% | [-29.7%, -27.3%] | 1.1 pts | 0.7 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | btree-map | 12 | 364 | 421 | 1.18× [1.15, 1.21] | +15.1% | [+12.9%, +17.4%] | 3.0 pts | 2.1 | no | yes |
| unique | path | 262144 | valuesBetween | ordered | btree-map | 12 | 2112 | 2612 | 1.25× [1.22, 1.28] | +20.0% | [+18.1%, +21.8%] | 2.7 pts | 2.4 | yes | yes |
| unique | path | 262144 | prefix | ordered | btree-map | 12 | 1719 | 2193 | 1.26× [1.22, 1.31] | +20.9% | [+18.2%, +23.5%] | 3.7 pts | 0.6 | no | yes |
| unique | path | 262144 | churn | ordered | btree-map | 12 | 672 | 761 | 1.12× [1.12, 1.13] | +11.1% | [+10.4%, +11.7%] | 1.1 pts | 1.1 | yes | yes |
| unique | str | 4096 | valuesFor | ordered | btree-map | 6 | 43.8 | 99.1 | 2.26× [2.25, 2.27] | +55.7% | [+55.5%, +55.9%] | 0.2 pts | 1.2 | yes | yes |
| unique | str | 4096 | valuesBetween | ordered | btree-map | 6 | 369 | 425 | 1.16× [1.15, 1.17] | +13.6% | [+12.9%, +14.2%] | 0.6 pts | 1.5 | yes | yes |
| unique | str | 4096 | prefix | ordered | btree-map | 6 | 569 | 825 | 1.45× [1.44, 1.46] | +31.1% | [+30.7%, +31.5%] | 0.4 pts | 2.6 | yes | yes |
| unique | str | 4096 | churn | ordered | btree-map | 6 | 96.6 | 141 | 1.46× [1.45, 1.48] | +31.5% | [+30.9%, +32.2%] | 0.6 pts | 1.3 | yes | yes |
| unique | str | 4096 | build | ordered | btree-map | 6 | 1.64 ms | 1.96 ms | 1.19× [1.18, 1.20] | +16.0% | [+15.4%, +16.7%] | 0.6 pts | 1.2 | yes | yes |
| unique | str | 16384 | valuesFor | ordered | btree-map | 6 | 57.3 | 130 | 2.28× [2.25, 2.31] | +56.1% | [+55.6%, +56.6%] | 0.5 pts | 2.7 | yes | yes |
| unique | str | 16384 | valuesBetween | ordered | btree-map | 6 | 454 | 503 | 1.10× [1.09, 1.12] | +9.5% | [+8.2%, +10.8%] | 1.2 pts | 2.6 | yes | yes |
| unique | str | 16384 | prefix | ordered | btree-map | 6 | 2382 | 3441 | 1.44× [1.42, 1.46] | +30.4% | [+29.4%, +31.4%] | 0.9 pts | 3.4 | yes | yes |
| unique | str | 16384 | churn | ordered | btree-map | 6 | 125 | 199 | 1.61× [1.59, 1.62] | +37.7% | [+37.3%, +38.2%] | 0.5 pts | 0.9 | yes | yes |
| unique | str | 16384 | build | ordered | btree-map | 6 | 8.06 ms | 10.45 ms | 1.29× [1.28, 1.31] | +22.7% | [+21.8%, +23.6%] | 0.8 pts | 1.2 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | btree-map | 8 | 153 | 257 | 1.73× [1.66, 1.80] | +42.1% | [+39.7%, +44.5%] | 2.7 pts | 2.0 | yes | yes |
| unique | str | 262144 | valuesBetween | ordered | btree-map | 8 | 650 | 1171 | 1.83× [1.75, 1.92] | +45.4% | [+42.9%, +48.0%] | 2.6 pts | 0.7 | yes | yes |
| unique | str | 262144 | prefix | ordered | btree-map | 8 | 60.0 µs | 127.4 µs | 2.16× [1.95, 2.41] | +53.6% | [+48.8%, +58.5%] | 5.1 pts | 0.7 | yes | yes |
| unique | str | 262144 | churn | ordered | btree-map | 8 | 385 | 579 | 1.49× [1.47, 1.51] | +32.7% | [+31.9%, +33.6%] | 1.0 pts | 1.3 | yes | yes |
| unique | street | 4096 | valuesFor | ordered | btree-map | 12 | 48.1 | 90.3 | 1.88× [1.86, 1.89] | +46.7% | [+46.2%, +47.1%] | 0.6 pts | 2.7 | yes | yes |
| unique | street | 4096 | valuesBetween | ordered | btree-map | 12 | 400 | 488 | 1.22× [1.21, 1.24] | +18.1% | [+17.2%, +19.0%] | 1.2 pts | 1.5 | yes | yes |
| unique | street | 4096 | prefix | ordered | btree-map | 12 | 102 | 129 | 1.27× [1.26, 1.27] | +21.0% | [+20.8%, +21.3%] | 0.8 pts | 1.8 | yes | yes |
| unique | street | 4096 | churn | ordered | btree-map | 12 | 95.7 | 136 | 1.42× [1.40, 1.44] | +29.5% | [+28.7%, +30.4%] | 1.1 pts | 1.7 | yes | yes |
| unique | street | 4096 | build | ordered | btree-map | 12 | 1.78 ms | 2.11 ms | 1.20× [1.17, 1.23] | +16.7% | [+14.8%, +18.5%] | 2.4 pts | 2.8 | yes | yes |
| unique | street | 16384 | valuesFor | ordered | btree-map | 6 | 61.4 | 124 | 2.01× [1.97, 2.05] | +50.3% | [+49.3%, +51.3%] | 1.0 pts | 2.8 | yes | yes |
| unique | street | 16384 | valuesBetween | ordered | btree-map | 6 | 477 | 587 | 1.22× [1.20, 1.25] | +18.2% | [+16.8%, +19.7%] | 1.4 pts | 1.9 | yes | yes |
| unique | street | 16384 | prefix | ordered | btree-map | 6 | 178 | 267 | 1.50× [1.48, 1.52] | +33.3% | [+32.3%, +34.2%] | 0.9 pts | 2.0 | yes | yes |
| unique | street | 16384 | churn | ordered | btree-map | 6 | 122 | 191 | 1.55× [1.53, 1.57] | +35.5% | [+34.7%, +36.2%] | 0.7 pts | 1.8 | yes | yes |
| unique | street | 16384 | build | ordered | btree-map | 6 | 8.64 ms | 11.18 ms | 1.29× [1.28, 1.30] | +22.7% | [+22.1%, +23.3%] | 0.6 pts | 0.5 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | btree-map | 6 | 20.6 | 88.0 | 4.27× [4.25, 4.29] | +76.6% | [+76.5%, +76.7%] | 0.1 pts | 1.7 | yes | yes |
| unique | u64 | 4096 | valuesBetween | ordered | btree-map | 6 | 253 | 422 | 1.67× [1.66, 1.68] | +40.0% | [+39.7%, +40.4%] | 0.3 pts | 2.0 | yes | yes |
| unique | u64 | 4096 | churn | ordered | btree-map | 6 | 40.9 | 121 | 2.96× [2.94, 2.99] | +66.2% | [+65.9%, +66.5%] | 0.3 pts | 1.0 | yes | yes |
| unique | u64 | 4096 | build | ordered | btree-map | 6 | 689.2 µs | 1.74 ms | 2.52× [2.51, 2.54] | +60.4% | [+60.1%, +60.6%] | 0.3 pts | 1.7 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | btree-map | 6 | 28.1 | 111 | 3.95× [3.88, 4.02] | +74.7% | [+74.2%, +75.1%] | 0.4 pts | 4.6 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | btree-map | 6 | 300 | 453 | 1.51× [1.50, 1.52] | +33.7% | [+33.2%, +34.2%] | 0.5 pts | 2.1 | yes | yes |
| unique | u64 | 16384 | churn | ordered | btree-map | 6 | 60.2 | 168 | 2.80× [2.76, 2.84] | +64.3% | [+63.7%, +64.8%] | 0.5 pts | 1.8 | yes | yes |
| unique | u64 | 16384 | build | ordered | btree-map | 6 | 3.68 ms | 8.98 ms | 2.43× [2.41, 2.46] | +58.9% | [+58.5%, +59.3%] | 0.4 pts | 1.4 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | btree-map | 6 | 56.9 | 211 | 3.72× [3.44, 4.05] | +73.1% | [+70.9%, +75.3%] | 2.1 pts | 1.3 | yes | yes |
| unique | u64 | 262144 | valuesBetween | ordered | btree-map | 6 | 315 | 715 | 2.25× [2.21, 2.29] | +55.5% | [+54.8%, +56.3%] | 0.7 pts | 1.2 | yes | yes |
| unique | u64 | 262144 | churn | ordered | btree-map | 6 | 219 | 428 | 1.96× [1.91, 2.02] | +49.1% | [+47.7%, +50.5%] | 1.3 pts | 1.1 | yes | yes |
| unique | url | 4096 | valuesFor | ordered | btree-map | 6 | 76.6 | 111 | 1.45× [1.44, 1.46] | +31.0% | [+30.3%, +31.6%] | 0.6 pts | 1.3 | yes | yes |
| unique | url | 4096 | valuesBetween | ordered | btree-map | 6 | 737 | 527 | 0.69× [0.68, 0.71] | -43.9% | [-46.3%, -41.5%] | 2.3 pts | 1.0 | yes | yes |
| unique | url | 4096 | prefix | ordered | btree-map | 6 | 185 | 125 | 0.68× [0.67, 0.68] | -48.1% | [-49.7%, -46.5%] | 1.5 pts | 0.8 | yes | yes |
| unique | url | 4096 | churn | ordered | btree-map | 6 | 182 | 177 | 0.98× [0.96, 1.00] | -2.1% | [-3.7%, -0.4%] | 1.6 pts | 1.6 | yes | yes |
| unique | url | 4096 | build | ordered | btree-map | 6 | 3.22 ms | 2.45 ms | 0.76× [0.75, 0.76] | -31.7% | [-32.5%, -30.9%] | 0.7 pts | 0.8 | yes | yes |
| unique | url | 16384 | valuesFor | ordered | btree-map | 6 | 98.9 | 159 | 1.58× [1.53, 1.64] | +36.8% | [+34.5%, +39.0%] | 2.2 pts | 5.0 | yes | yes |
| unique | url | 16384 | valuesBetween | ordered | btree-map | 6 | 931 | 705 | 0.76× [0.75, 0.76] | -32.2% | [-33.1%, -31.3%] | 0.9 pts | 0.9 | yes | yes |
| unique | url | 16384 | prefix | ordered | btree-map | 6 | 232 | 194 | 0.84× [0.83, 0.85] | -18.9% | [-20.3%, -17.5%] | 1.3 pts | 1.0 | yes | yes |
| unique | url | 16384 | churn | ordered | btree-map | 6 | 218 | 239 | 1.09× [1.08, 1.10] | +8.1% | [+7.1%, +9.1%] | 1.0 pts | 1.4 | yes | yes |
| unique | url | 16384 | build | ordered | btree-map | 6 | 15.01 ms | 12.74 ms | 0.85× [0.84, 0.86] | -18.1% | [-19.7%, -16.5%] | 1.5 pts | 1.4 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | btree-map | 12 | 335 | 412 | 1.23× [1.21, 1.24] | +18.5% | [+17.4%, +19.6%] | 1.8 pts | 1.5 | yes | yes |
| unique | url | 262144 | valuesBetween | ordered | btree-map | 12 | 2375 | 2612 | 1.10× [1.07, 1.12] | +8.9% | [+6.8%, +11.0%] | 3.0 pts | 3.1 | no | yes |
| unique | url | 262144 | prefix | ordered | btree-map | 12 | 773 | 832 | 1.07× [1.05, 1.10] | +6.9% | [+4.8%, +9.1%] | 3.6 pts | 4.2 | no | yes |
| unique | url | 262144 | churn | ordered | btree-map | 12 | 641 | 734 | 1.14× [1.12, 1.16] | +12.4% | [+10.6%, +14.2%] | 2.2 pts | 1.4 | yes | yes |
| unique | uuid | 4096 | valuesFor | ordered | btree-map | 6 | 39.7 | 99.1 | 2.51× [2.48, 2.55] | +60.2% | [+59.6%, +60.7%] | 0.5 pts | 4.6 | yes | yes |
| unique | uuid | 4096 | valuesBetween | ordered | btree-map | 6 | 444 | 429 | 0.96× [0.95, 0.98] | -3.8% | [-5.6%, -2.1%] | 1.7 pts | 3.3 | yes | yes |
| unique | uuid | 4096 | prefix | ordered | btree-map | 6 | 74.2 | 100 | 1.36× [1.33, 1.40] | +26.5% | [+24.5%, +28.5%] | 1.9 pts | 4.7 | yes | yes |
| unique | uuid | 4096 | churn | ordered | btree-map | 6 | 94.3 | 138 | 1.47× [1.46, 1.48] | +32.0% | [+31.6%, +32.3%] | 0.3 pts | 0.6 | yes | yes |
| unique | uuid | 4096 | build | ordered | btree-map | 6 | 1.57 ms | 1.97 ms | 1.26× [1.24, 1.27] | +20.4% | [+19.3%, +21.5%] | 1.1 pts | 2.3 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | btree-map | 12 | 44.8 | 137 | 3.04× [2.99, 3.09] | +67.1% | [+66.5%, +67.7%] | 0.6 pts | 5.1 | yes | yes |
| unique | uuid | 16384 | valuesBetween | ordered | btree-map | 12 | 448 | 556 | 1.23× [1.20, 1.26] | +18.7% | [+16.8%, +20.7%] | 2.4 pts | 6.6 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | btree-map | 12 | 80.7 | 146 | 1.79× [1.76, 1.82] | +44.2% | [+43.3%, +45.2%] | 1.0 pts | 3.7 | yes | yes |
| unique | uuid | 16384 | churn | ordered | btree-map | 12 | 112 | 192 | 1.73× [1.72, 1.73] | +42.0% | [+41.8%, +42.3%] | 0.4 pts | 0.9 | yes | yes |
| unique | uuid | 16384 | build | ordered | btree-map | 12 | 7.30 ms | 10.60 ms | 1.45× [1.44, 1.46] | +31.0% | [+30.4%, +31.6%] | 0.7 pts | 1.2 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | btree-map | 12 | 193 | 312 | 1.66× [1.56, 1.77] | +39.7% | [+35.7%, +43.6%] | 3.7 pts | 2.9 | yes | yes |
| unique | uuid | 262144 | valuesBetween | ordered | btree-map | 12 | 1078 | 1736 | 1.63× [1.53, 1.75] | +38.7% | [+34.5%, +42.9%] | 4.2 pts | 2.4 | no | yes |
| unique | uuid | 262144 | prefix | ordered | btree-map | 12 | 206 | 393 | 1.92× [1.81, 2.03] | +47.9% | [+44.9%, +50.8%] | 2.9 pts | 3.2 | yes | yes |
| unique | uuid | 262144 | churn | ordered | btree-map | 12 | 412 | 620 | 1.50× [1.48, 1.53] | +33.5% | [+32.4%, +34.5%] | 1.2 pts | 1.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique email n=4096 valuesFor: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=4096 churn: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 valuesFor: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 valuesBetween: ordered vs btree-map: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 prefix: ordered vs btree-map: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 churn: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=4096 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.70% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique path n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.40% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique str n=4096 prefix: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 prefix: ordered vs btree-map: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -9.94% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 valuesFor: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=4096 prefix: ordered vs btree-map: the A/A validations found a systematic difference of -0.35% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique street n=4096 build: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 valuesFor: ordered vs btree-map: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 prefix: ordered vs btree-map: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.57% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique u64 n=16384 valuesFor: ordered vs btree-map: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 valuesBetween: ordered vs btree-map: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of +0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=16384 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of -0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=16384 valuesFor: ordered vs btree-map: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 valuesBetween: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 prefix: ordered vs btree-map: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 valuesFor: ordered vs btree-map: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 valuesBetween: ordered vs btree-map: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 prefix: ordered vs btree-map: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 build: ordered vs btree-map: the A/A validations found a systematic difference of -0.29% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique uuid n=4096 build: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 valuesFor: ordered vs btree-map: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 valuesBetween: ordered vs btree-map: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 prefix: ordered vs btree-map: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +0.48% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique uuid n=262144 valuesFor: ordered vs btree-map: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesBetween: ordered vs btree-map: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 prefix: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
