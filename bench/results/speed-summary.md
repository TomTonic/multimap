| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | str | 4096 | valuesFor | ordered | btree-sets | 5 | 67.0 | 192 | 2.85× [2.83, 2.87] | +65.0% | [+64.6%, +65.2%] | 0.2 pts | 1.3 | yes |
| multi | str | 4096 | valuesFor | ordered | hashed | 5 | 66.3 | 42.9 | 0.65× [0.64, 0.68] | -54.3% | [-57.4%, -47.8%] | 3.8 pts | 3.5 | yes |
| multi | str | 4096 | valuesFor | ordered | map-sets | 5 | 66.3 | 102 | 1.54× [1.53, 1.54] | +35.0% | [+34.5%, +35.2%] | 0.3 pts | 0.9 | yes |
| multi | str | 4096 | valuesBetween | ordered | btree-sets | 5 | 3746 | 7963 | 2.13× [2.11, 2.14] | +53.0% | [+52.7%, +53.3%] | 0.3 pts | 0.5 | yes |
| multi | str | 4096 | valuesBetween | ordered | hashed | 5 | 3916 | 58.9 µs | 14.95× [14.71, 15.42] | +93.3% | [+93.2%, +93.5%] | 0.1 pts | 1.7 | yes |
| multi | str | 4096 | valuesBetween | ordered | map-sets | 5 | 3780 | 59.2 µs | 15.60× [15.39, 15.71] | +93.6% | [+93.5%, +93.6%] | 0.1 pts | 0.7 | yes |
| multi | str | 4096 | churn | ordered | btree-sets | 5 | 82.7 | 190 | 2.32× [2.26, 2.35] | +56.9% | [+55.7%, +57.5%] | 0.7 pts | 2.1 | yes |
| multi | str | 4096 | churn | ordered | hashed | 5 | 79.5 | 53.3 | 0.67× [0.67, 0.68] | -49.2% | [-49.9%, -47.2%] | 1.1 pts | 0.8 | yes |
| multi | str | 4096 | churn | ordered | map-sets | 5 | 81.7 | 61.4 | 0.75× [0.74, 0.78] | -33.1% | [-35.1%, -29.0%] | 2.4 pts | 1.5 | yes |
| multi | str | 4096 | build | ordered | btree-sets | 5 | 6.48 ms | 15.49 ms | 2.40× [2.38, 2.41] | +58.3% | [+58.0%, +58.5%] | 0.2 pts | 1.3 | yes |
| multi | str | 4096 | build | ordered | hashed | 5 | 6.47 ms | 4.65 ms | 0.72× [0.71, 0.72] | -39.4% | [-40.6%, -38.4%] | 0.9 pts | 1.7 | yes |
| multi | str | 4096 | build | ordered | map-sets | 5 | 6.47 ms | 5.38 ms | 0.83× [0.83, 0.84] | -20.0% | [-21.0%, -19.4%] | 0.7 pts | 1.3 | yes |
| multi | str | 1048576 | valuesFor | ordered | btree-sets | 27 | 569 | 1277 | 2.25× [2.20, 2.31] | +55.6% | [+54.5%, +56.7%] | 2.7 pts | 2.7 | yes |
| multi | str | 1048576 | valuesFor | ordered | hashed | 27 | 482 | 228 | 0.47× [0.47, 0.48] | -110.7% | [-114.0%, -108.9%] | 6.5 pts | 3.2 | yes |
| multi | str | 1048576 | valuesFor | ordered | map-sets | 27 | 533 | 519 | 0.97× [0.96, 1.00] | -2.6% | [-3.7%, +0.2%] | 5.0 pts | 3.7 | yes |
| multi | str | 1048576 | valuesBetween | ordered | btree-sets | 27 | 13.4 µs | 33.3 µs | 2.52× [2.47, 2.56] | +60.3% | [+59.5%, +60.9%] | 1.8 pts | 6.6 | yes |
| multi | str | 1048576 | churn | ordered | btree-sets | 27 | 749 | 1474 | 1.98× [1.95, 2.03] | +49.4% | [+48.6%, +50.8%] | 2.7 pts | 4.1 | yes |
| multi | str | 1048576 | churn | ordered | hashed | 27 | 625 | 400 | 0.64× [0.63, 0.66] | -57.2% | [-58.7%, -50.5%] | 10.3 pts | 5.8 | yes |
| multi | str | 1048576 | churn | ordered | map-sets | 27 | 690 | 496 | 0.73× [0.71, 0.74] | -37.7% | [-41.4%, -35.0%] | 8.1 pts | 4.8 | yes |
| multi | u64 | 4096 | valuesFor | ordered | btree-sets | 5 | 39.9 | 174 | 4.41× [4.32, 4.50] | +77.3% | [+76.9%, +77.8%] | 0.4 pts | 3.3 | yes |
| multi | u64 | 4096 | valuesFor | ordered | hashed | 5 | 40.6 | 39.6 | 0.98× [0.97, 0.99] | -2.2% | [-3.2%, -1.0%] | 0.9 pts | 0.7 | yes |
| multi | u64 | 4096 | valuesFor | ordered | map-sets | 5 | 39.5 | 98.9 | 2.50× [2.46, 2.57] | +59.9% | [+59.4%, +61.2%] | 0.7 pts | 3.4 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | btree-sets | 5 | 2898 | 7628 | 2.61× [2.53, 2.66] | +61.7% | [+60.5%, +62.5%] | 0.8 pts | 1.8 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | hashed | 5 | 3126 | 57.7 µs | 18.41× [18.17, 18.64] | +94.6% | [+94.5%, +94.6%] | 0.1 pts | 0.8 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | map-sets | 5 | 2931 | 57.9 µs | 19.74× [18.61, 20.34] | +94.9% | [+94.6%, +95.1%] | 0.2 pts | 4.3 | yes |
| multi | u64 | 4096 | churn | ordered | btree-sets | 5 | 44.7 | 168 | 3.73× [3.62, 3.82] | +73.2% | [+72.4%, +73.8%] | 0.6 pts | 2.0 | yes |
| multi | u64 | 4096 | churn | ordered | hashed | 5 | 43.7 | 47.2 | 1.08× [1.07, 1.09] | +7.6% | [+6.6%, +8.6%] | 0.8 pts | 1.1 | yes |
| multi | u64 | 4096 | churn | ordered | map-sets | 5 | 44.0 | 52.9 | 1.20× [1.19, 1.21] | +16.7% | [+15.9%, +17.6%] | 0.7 pts | 0.7 | yes |
| multi | u64 | 4096 | build | ordered | btree-sets | 5 | 3.88 ms | 13.78 ms | 3.55× [3.50, 3.61] | +71.9% | [+71.5%, +72.3%] | 0.3 pts | 2.8 | yes |
| multi | u64 | 4096 | build | ordered | hashed | 5 | 3.87 ms | 4.30 ms | 1.11× [1.10, 1.12] | +10.0% | [+8.9%, +10.9%] | 0.8 pts | 2.3 | yes |
| multi | u64 | 4096 | build | ordered | map-sets | 5 | 3.88 ms | 4.95 ms | 1.28× [1.27, 1.28] | +21.8% | [+21.3%, +22.1%] | 0.3 pts | 0.9 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | btree-sets | 40 | 246 | 1051 | 4.26× [4.23, 4.34] | +76.5% | [+76.3%, +76.9%] | 0.9 pts | 2.6 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | hashed | 40 | 226 | 219 | 0.97× [0.95, 0.97] | -3.6% | [-5.7%, -2.7%] | 4.7 pts | 4.8 | yes |
| multi | u64 | 1048576 | valuesFor | ordered | map-sets | 40 | 237 | 507 | 2.15× [2.09, 2.18] | +53.5% | [+52.1%, +54.0%] | 3.0 pts | 7.4 | yes |
| multi | u64 | 1048576 | valuesBetween | ordered | btree-sets | 40 | 9094 | 31.2 µs | 3.40× [3.39, 3.49] | +70.6% | [+70.5%, +71.3%] | 1.3 pts | 4.5 | yes |
| multi | u64 | 1048576 | churn | ordered | btree-sets | 40 | 442 | 1256 | 2.86× [2.61, 2.91] | +65.0% | [+61.7%, +65.6%] | 6.1 pts | 11.0 | yes |
| multi | u64 | 1048576 | churn | ordered | hashed | 40 | 338 | 356 | 1.06× [1.00, 1.08] | +5.4% | [-0.0%, +7.4%] | 11.5 pts | 7.6 | no |
| multi | u64 | 1048576 | churn | ordered | map-sets | 40 | 417 | 460 | 1.12× [1.05, 1.12] | +11.0% | [+5.1%, +10.6%] | 8.6 pts | 6.6 | no |
| unique | str | 4096 | valuesFor | ordered | btree-map | 5 | 41.3 | 110 | 2.65× [2.63, 2.66] | +62.3% | [+62.0%, +62.4%] | 0.2 pts | 0.8 | yes |
| unique | str | 4096 | valuesBetween | ordered | btree-map | 5 | 1744 | 480 | 0.28× [0.27, 0.28] | -263.2% | [-264.8%, -259.3%] | 2.2 pts | 1.4 | yes |
| unique | str | 4096 | churn | ordered | btree-map | 5 | 96.3 | 149 | 1.55× [1.53, 1.56] | +35.5% | [+34.8%, +35.9%] | 0.4 pts | 1.0 | yes |
| unique | str | 4096 | build | ordered | btree-map | 5 | 1.14 ms | 1.75 ms | 1.54× [1.53, 1.55] | +35.0% | [+34.6%, +35.5%] | 0.4 pts | 0.5 | yes |
| unique | str | 1048576 | valuesFor | ordered | btree-map | 14 | 465 | 792 | 1.71× [1.65, 1.75] | +41.5% | [+39.5%, +42.7%] | 2.8 pts | 2.7 | yes |
| unique | str | 1048576 | valuesBetween | ordered | btree-map | 14 | 7802 | 4083 | 0.52× [0.51, 0.53] | -90.6% | [-95.8%, -88.2%] | 6.6 pts | 5.2 | yes |
| unique | str | 1048576 | churn | ordered | btree-map | 14 | 817 | 1174 | 1.43× [1.36, 1.47] | +30.3% | [+26.2%, +32.1%] | 5.0 pts | 4.8 | yes |
| unique | u64 | 4096 | valuesFor | ordered | btree-map | 5 | 16.6 | 93.5 | 5.63× [5.53, 5.74] | +82.2% | [+81.9%, +82.6%] | 0.3 pts | 1.5 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | btree-map | 5 | 1130 | 477 | 0.42× [0.42, 0.43] | -137.9% | [-140.3%, -135.2%] | 2.0 pts | 2.4 | yes |
| unique | u64 | 4096 | churn | ordered | btree-map | 5 | 46.5 | 131 | 2.82× [2.79, 2.86] | +64.5% | [+64.2%, +65.0%] | 0.3 pts | 2.2 | yes |
| unique | u64 | 4096 | build | ordered | btree-map | 5 | 654.4 µs | 1.54 ms | 2.36× [2.14, 2.49] | +57.6% | [+53.3%, +59.8%] | 2.6 pts | 5.3 | yes |
| unique | u64 | 1048576 | valuesFor | ordered | btree-map | 9 | 178 | 526 | 2.94× [2.91, 2.99] | +66.0% | [+65.6%, +66.5%] | 0.6 pts | 1.3 | yes |
| unique | u64 | 1048576 | valuesBetween | ordered | btree-map | 9 | 3661 | 2349 | 0.64× [0.62, 0.67] | -57.0% | [-60.7%, -50.2%] | 6.8 pts | 3.2 | yes |
| unique | u64 | 1048576 | churn | ordered | btree-map | 9 | 534 | 925 | 1.76× [1.61, 1.82] | +43.3% | [+38.0%, +45.1%] | 4.6 pts | 6.1 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
