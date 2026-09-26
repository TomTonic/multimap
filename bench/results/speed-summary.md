| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | str | 4096 | valuesFor | ordered | btree-sets | 5 | 81.6 | 175 | 2.14× [2.13, 2.16] | +53.3% | [+53.0%, +53.7%] | 0.3 pts | 2.3 | yes |
| multi | str | 4096 | valuesFor | ordered | hashed | 5 | 80.6 | 38.5 | 0.48× [0.47, 0.49] | -109.5% | [-112.1%, -105.2%] | 2.8 pts | 2.1 | yes |
| multi | str | 4096 | valuesFor | ordered | map-sets | 5 | 80.5 | 93.8 | 1.16× [1.15, 1.17] | +14.1% | [+13.3%, +14.7%] | 0.6 pts | 1.5 | yes |
| multi | str | 4096 | valuesBetween | ordered | btree-sets | 5 | 3063 | 7277 | 2.35× [2.34, 2.38] | +57.5% | [+57.2%, +58.0%] | 0.3 pts | 0.7 | yes |
| multi | str | 4096 | valuesBetween | ordered | hashed | 5 | 3190 | 53.6 µs | 16.77× [16.62, 17.05] | +94.0% | [+94.0%, +94.1%] | 0.1 pts | 1.0 | yes |
| multi | str | 4096 | valuesBetween | ordered | map-sets | 5 | 3102 | 54.1 µs | 17.45× [17.34, 17.63] | +94.3% | [+94.2%, +94.3%] | 0.0 pts | 0.6 | yes |
| multi | str | 4096 | churn | ordered | btree-sets | 5 | 147 | 174 | 1.18× [1.17, 1.19] | +15.3% | [+14.3%, +15.9%] | 0.6 pts | 0.9 | yes |
| multi | str | 4096 | churn | ordered | hashed | 5 | 144 | 50.0 | 0.35× [0.35, 0.35] | -184.9% | [-188.6%, -183.1%] | 2.2 pts | 1.2 | yes |
| multi | str | 4096 | churn | ordered | map-sets | 5 | 144 | 56.8 | 0.39× [0.39, 0.41] | -153.6% | [-158.8%, -142.2%] | 6.7 pts | 2.9 | yes |
| multi | str | 4096 | build | ordered | btree-sets | 5 | 14.51 ms | 14.15 ms | 0.98× [0.97, 0.98] | -2.3% | [-2.9%, -1.9%] | 0.4 pts | 1.4 | yes |
| multi | str | 4096 | build | ordered | hashed | 5 | 14.51 ms | 4.25 ms | 0.29× [0.29, 0.29] | -241.0% | [-243.0%, -240.2%] | 1.1 pts | 1.2 | yes |
| multi | str | 4096 | build | ordered | map-sets | 5 | 14.56 ms | 5.02 ms | 0.34× [0.34, 0.35] | -190.7% | [-192.9%, -188.2%] | 1.9 pts | 1.2 | yes |
| multi | str | 1048576 | valuesFor | ordered | btree-sets | 30 | 497 | 1203 | 2.43× [2.38, 2.45] | +58.9% | [+57.9%, +59.2%] | 1.7 pts | 2.3 | yes |
| multi | str | 1048576 | valuesFor | ordered | hashed | 30 | 406 | 210 | 0.52× [0.50, 0.52] | -93.2% | [-98.4%, -92.2%] | 8.4 pts | 3.1 | yes |
| multi | str | 1048576 | valuesFor | ordered | map-sets | 30 | 478 | 486 | 1.02× [1.01, 1.06] | +1.9% | [+0.7%, +5.4%] | 6.3 pts | 3.5 | no |
| multi | str | 1048576 | valuesBetween | ordered | btree-sets | 30 | 9382 | 30.8 µs | 3.24× [3.20, 3.28] | +69.2% | [+68.8%, +69.5%] | 1.0 pts | 3.7 | yes |
| multi | str | 1048576 | churn | ordered | btree-sets | 30 | 897 | 1376 | 1.56× [1.53, 1.58] | +36.0% | [+34.8%, +36.6%] | 2.4 pts | 2.4 | yes |
| multi | str | 1048576 | churn | ordered | hashed | 30 | 663 | 373 | 0.58× [0.54, 0.58] | -72.9% | [-85.6%, -73.0%] | 16.9 pts | 7.3 | yes |
| multi | str | 1048576 | churn | ordered | map-sets | 30 | 834 | 466 | 0.56× [0.56, 0.58] | -77.5% | [-79.7%, -72.1%] | 10.1 pts | 4.6 | yes |
| multi | u64 | 4096 | valuesFor | ordered | btree-sets | 5 | 45.8 | 164 | 3.59× [3.55, 3.66] | +72.1% | [+71.8%, +72.7%] | 0.4 pts | 2.4 | yes |
| multi | u64 | 4096 | valuesFor | ordered | hashed | 5 | 46.0 | 37.4 | 0.82× [0.81, 0.82] | -22.7% | [-23.6%, -21.4%] | 0.9 pts | 0.8 | yes |
| multi | u64 | 4096 | valuesFor | ordered | map-sets | 5 | 44.9 | 92.3 | 2.08× [2.06, 2.09] | +51.8% | [+51.4%, +52.3%] | 0.4 pts | 1.1 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | btree-sets | 5 | 2837 | 7036 | 2.48× [2.46, 2.52] | +59.7% | [+59.4%, +60.3%] | 0.4 pts | 0.6 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | hashed | 5 | 2992 | 53.3 µs | 18.22× [17.67, 18.48] | +94.5% | [+94.3%, +94.6%] | 0.1 pts | 1.6 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | map-sets | 5 | 2860 | 54.1 µs | 18.89× [18.67, 19.01] | +94.7% | [+94.6%, +94.7%] | 0.0 pts | 0.5 | yes |
| multi | u64 | 4096 | churn | ordered | btree-sets | 5 | 66.6 | 154 | 2.32× [2.30, 2.35] | +56.9% | [+56.5%, +57.4%] | 0.4 pts | 1.0 | yes |
| multi | u64 | 4096 | churn | ordered | hashed | 5 | 62.6 | 44.0 | 0.69× [0.68, 0.70] | -44.8% | [-46.6%, -42.7%] | 1.6 pts | 1.6 | yes |
| multi | u64 | 4096 | churn | ordered | map-sets | 5 | 63.4 | 50.3 | 0.77× [0.76, 0.79] | -29.7% | [-31.4%, -26.5%] | 2.0 pts | 1.5 | yes |
| multi | u64 | 4096 | build | ordered | btree-sets | 5 | 5.76 ms | 12.65 ms | 2.19× [2.17, 2.20] | +54.3% | [+54.0%, +54.5%] | 0.2 pts | 1.1 | yes |
| multi | u64 | 4096 | build | ordered | hashed | 5 | 5.76 ms | 3.90 ms | 0.68× [0.67, 0.69] | -47.4% | [-48.7%, -45.1%] | 1.5 pts | 2.1 | yes |
| multi | u64 | 4096 | build | ordered | map-sets | 5 | 5.86 ms | 4.58 ms | 0.78× [0.78, 0.79] | -27.5% | [-28.5%, -26.0%] | 1.0 pts | 1.7 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | btree-sets | 30 | 282 | 983 | 3.46× [3.39, 3.53] | +71.1% | [+70.5%, +71.6%] | 1.5 pts | 2.9 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | hashed | 30 | 258 | 203 | 0.80× [0.68, 0.85] | -25.6% | [-47.4%, -17.5%] | 40.0 pts | 17.5 | no |
| multi | u64 | 1048576 | valuesFor | ordered | map-sets | 30 | 262 | 473 | 1.82× [1.78, 1.84] | +45.0% | [+44.0%, +45.5%] | 2.1 pts | 2.4 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | btree-sets | 30 | 8992 | 29.2 µs | 3.23× [3.20, 3.29] | +69.1% | [+68.7%, +69.6%] | 1.2 pts | 3.6 | yes |
| multi | u64 | 1048576 | churn | ordered | btree-sets | 30 | 657 | 1218 | 1.82× [1.81, 1.87] | +45.1% | [+44.6%, +46.5%] | 2.5 pts | 2.9 | yes |
| multi | u64 | 1048576 | churn | ordered | hashed | 30 | 490 | 340 | 0.69× [0.66, 0.72] | -44.2% | [-50.8%, -39.1%] | 15.8 pts | 5.7 | no |
| multi | u64 | 1048576 | churn | ordered | map-sets | 30 | 585 | 432 | 0.73× [0.72, 0.74] | -37.2% | [-38.9%, -35.1%] | 5.1 pts | 2.5 | yes |
| unique | str | 4096 | valuesFor | ordered | btree-map | 5 | 60.5 | 102 | 1.68× [1.67, 1.69] | +40.6% | [+40.1%, +40.9%] | 0.3 pts | 2.0 | yes |
| unique | str | 4096 | valuesBetween | ordered | btree-map | 5 | 1235 | 431 | 0.35× [0.35, 0.35] | -186.9% | [-188.3%, -184.5%] | 1.5 pts | 1.7 | yes |
| unique | str | 4096 | churn | ordered | btree-map | 5 | 172 | 138 | 0.80× [0.80, 0.81] | -24.9% | [-25.8%, -23.6%] | 0.9 pts | 1.2 | yes |
| unique | str | 4096 | build | ordered | btree-map | 5 | 2.78 ms | 1.64 ms | 0.59× [0.59, 0.59] | -69.4% | [-70.4%, -69.0%] | 0.6 pts | 0.3 | yes |
| unique | str | 1048576 | valuesFor | ordered | btree-map | 16 | 358 | 748 | 2.11× [2.05, 2.14] | +52.6% | [+51.2%, +53.3%] | 1.9 pts | 2.2 | yes |
| unique | str | 1048576 | valuesBetween | ordered | btree-map | 16 | 3008 | 3524 | 1.17× [1.14, 1.19] | +14.8% | [+12.2%, +16.2%] | 3.8 pts | 3.9 | yes |
| unique | str | 1048576 | churn | ordered | btree-map | 16 | 791 | 1054 | 1.34× [1.33, 1.37] | +25.5% | [+24.6%, +26.8%] | 2.1 pts | 2.0 | yes |
| unique | u64 | 4096 | valuesFor | ordered | btree-map | 5 | 16.4 | 85.8 | 5.23× [5.11, 5.41] | +80.9% | [+80.4%, +81.5%] | 0.4 pts | 2.0 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | btree-map | 5 | 310 | 424 | 1.37× [1.36, 1.38] | +27.0% | [+26.7%, +27.4%] | 0.3 pts | 1.3 | yes |
| unique | u64 | 4096 | churn | ordered | btree-map | 5 | 38.2 | 120 | 3.15× [3.12, 3.23] | +68.3% | [+67.9%, +69.1%] | 0.5 pts | 4.9 | yes |
| unique | u64 | 4096 | build | ordered | btree-map | 5 | 481.5 µs | 1.42 ms | 2.95× [2.88, 2.99] | +66.1% | [+65.3%, +66.5%] | 0.5 pts | 1.3 | yes |
| unique | u64 | 1048576 | valuesFor | ordered | btree-map | 11 | 141 | 493 | 3.50× [3.42, 3.59] | +71.5% | [+70.8%, +72.1%] | 1.0 pts | 1.7 | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | btree-map | 11 | 1085 | 1827 | 1.68× [1.63, 1.89] | +40.4% | [+38.8%, +47.0%] | 6.1 pts | 5.9 | yes |
| unique | u64 | 1048576 | churn | ordered | btree-map | 11 | 426 | 844 | 2.01× [1.96, 2.08] | +50.3% | [+49.0%, +51.9%] | 2.2 pts | 1.9 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
