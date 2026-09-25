| keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| str | 4096 | valuesFor | hashed | map-sets | 5 | 51.6 | 105 | 2.04× [2.01, 2.06] | +50.9% | [+50.3%, +51.3%] | 0.4 pts | 3.2 | yes |
| str | 4096 | valuesFor | ordered | btree-sets | 5 | 86.4 | 206 | 2.38× [2.37, 2.40] | +58.0% | [+57.8%, +58.4%] | 0.2 pts | 2.8 | yes |
| str | 4096 | valuesFor | ordered | hashed | 5 | 84.9 | 50.5 | 0.59× [0.59, 0.60] | -68.6% | [-69.1%, -67.7%] | 0.6 pts | 0.8 | yes |
| str | 4096 | valuesFor | ordered | map-sets | 5 | 86.1 | 105 | 1.22× [1.21, 1.22] | +18.1% | [+17.7%, +18.3%] | 0.3 pts | 1.5 | yes |
| str | 4096 | valuesBetween | ordered | btree-sets | 5 | 4493 | 8300 | 1.84× [1.82, 1.85] | +45.5% | [+45.1%, +46.1%] | 0.4 pts | 1.6 | yes |
| str | 4096 | valuesBetween | ordered | hashed | 5 | 4506 | 66.6 µs | 14.81× [14.68, 14.91] | +93.2% | [+93.2%, +93.3%] | 0.0 pts | 1.6 | yes |
| str | 4096 | addRemove | hashed | map-sets | 5 | 54.3 | 57.4 | 1.05× [1.04, 1.06] | +5.2% | [+4.2%, +5.8%] | 0.7 pts | 10.3 | yes |
| str | 4096 | addRemove | ordered | btree-sets | 5 | 139 | 262 | 1.87× [1.85, 1.89] | +46.6% | [+46.0%, +47.1%] | 0.4 pts | 0.9 | yes |
| str | 4096 | addRemove | ordered | hashed | 5 | 133 | 56.9 | 0.43× [0.42, 0.44] | -134.5% | [-138.0%, -128.6%] | 3.8 pts | 1.4 | yes |
| str | 4096 | addRemove | ordered | map-sets | 5 | 130 | 57.4 | 0.44× [0.44, 0.44] | -126.9% | [-127.8%, -125.6%] | 0.9 pts | 2.7 | yes |
| str | 4096 | build | hashed | map-sets | 5 | 1.33 ms | 1.61 ms | 1.21× [1.20, 1.21] | +17.1% | [+16.4%, +17.6%] | 0.5 pts | 1.6 | yes |
| str | 4096 | build | ordered | btree-sets | 5 | 2.04 ms | 3.22 ms | 1.58× [1.57, 1.59] | +36.6% | [+36.2%, +37.1%] | 0.4 pts | 2.4 | yes |
| str | 4096 | build | ordered | hashed | 5 | 2.04 ms | 1.32 ms | 0.65× [0.64, 0.65] | -54.5% | [-55.7%, -53.8%] | 0.8 pts | 2.1 | yes |
| str | 4096 | build | ordered | map-sets | 5 | 2.05 ms | 1.59 ms | 0.78× [0.77, 0.78] | -28.4% | [-29.9%, -27.6%] | 0.9 pts | 1.9 | yes |
| str | 1048576 | valuesFor | hashed | map-sets | 15 | 202 | 408 | 2.01× [2.00, 2.03] | +50.4% | [+50.0%, +50.7%] | 0.6 pts | 2.5 | yes |
| str | 1048576 | valuesFor | ordered | btree-sets | 15 | 416 | 1169 | 2.82× [2.77, 2.85] | +64.6% | [+63.9%, +64.9%] | 0.8 pts | 3.2 | yes |
| str | 1048576 | valuesFor | ordered | hashed | 15 | 359 | 200 | 0.56× [0.55, 0.56] | -79.2% | [-82.0%, -78.6%] | 3.1 pts | 3.5 | yes |
| str | 1048576 | valuesFor | ordered | map-sets | 15 | 369 | 414 | 1.11× [1.08, 1.13] | +9.9% | [+7.7%, +11.5%] | 3.4 pts | 5.2 | yes |
| str | 1048576 | valuesBetween | ordered | btree-sets | 15 | 9640 | 21.2 µs | 2.21× [2.16, 2.25] | +54.7% | [+53.7%, +55.5%] | 1.6 pts | 6.1 | yes |
| str | 1048576 | addRemove | hashed | map-sets | 15 | 311 | 509 | 1.60× [1.59, 1.66] | +37.6% | [+37.1%, +39.6%] | 2.2 pts | 10.8 | yes |
| str | 1048576 | addRemove | ordered | btree-sets | 15 | 649 | 1235 | 1.90× [1.78, 1.97] | +47.2% | [+43.9%, +49.1%] | 4.8 pts | 14.7 | yes |
| str | 1048576 | addRemove | ordered | hashed | 15 | 583 | 311 | 0.53× [0.52, 0.54] | -87.4% | [-92.0%, -84.6%] | 6.7 pts | 10.6 | yes |
| str | 1048576 | addRemove | ordered | map-sets | 15 | 592 | 511 | 0.85× [0.83, 0.86] | -17.6% | [-20.0%, -16.0%] | 3.6 pts | 8.2 | yes |
| u64 | 4096 | valuesFor | hashed | map-sets | 6 | 49.6 | 103 | 2.07× [2.06, 2.09] | +51.6% | [+51.4%, +52.2%] | 0.4 pts | 2.7 | yes |
| u64 | 4096 | valuesFor | ordered | btree-sets | 6 | 46.1 | 192 | 4.16× [4.14, 4.18] | +76.0% | [+75.8%, +76.1%] | 0.1 pts | 2.1 | yes |
| u64 | 4096 | valuesFor | ordered | hashed | 6 | 45.5 | 48.8 | 1.08× [1.07, 1.09] | +7.2% | [+6.2%, +8.0%] | 0.9 pts | 4.4 | yes |
| u64 | 4096 | valuesFor | ordered | map-sets | 6 | 45.8 | 103 | 2.25× [2.24, 2.26] | +55.6% | [+55.3%, +55.8%] | 0.2 pts | 1.7 | yes |
| u64 | 4096 | valuesBetween | ordered | btree-sets | 6 | 3569 | 8091 | 2.27× [2.26, 2.29] | +55.9% | [+55.7%, +56.2%] | 0.2 pts | 1.2 | yes |
| u64 | 4096 | valuesBetween | ordered | hashed | 6 | 3557 | 63.7 µs | 17.87× [17.72, 18.03] | +94.4% | [+94.4%, +94.5%] | 0.0 pts | 1.5 | yes |
| u64 | 4096 | addRemove | hashed | map-sets | 6 | 53.6 | 54.5 | 1.02× [1.01, 1.03] | +2.0% | [+0.9%, +2.8%] | 0.9 pts | 20.1 | yes |
| u64 | 4096 | addRemove | ordered | btree-sets | 6 | 51.9 | 239 | 4.62× [4.59, 4.66] | +78.3% | [+78.2%, +78.5%] | 0.1 pts | 0.9 | yes |
| u64 | 4096 | addRemove | ordered | hashed | 6 | 50.1 | 53.4 | 1.07× [1.05, 1.09] | +6.3% | [+5.2%, +8.0%] | 1.4 pts | 7.0 | yes |
| u64 | 4096 | addRemove | ordered | map-sets | 6 | 49.8 | 54.6 | 1.11× [1.09, 1.13] | +9.6% | [+8.0%, +11.8%] | 1.8 pts | 13.9 | yes |
| u64 | 4096 | build | hashed | map-sets | 6 | 1.25 ms | 1.53 ms | 1.22× [1.21, 1.24] | +18.1% | [+17.4%, +19.3%] | 0.9 pts | 2.4 | yes |
| u64 | 4096 | build | ordered | btree-sets | 6 | 1.16 ms | 3.13 ms | 2.71× [2.70, 2.73] | +63.1% | [+63.0%, +63.4%] | 0.2 pts | 1.6 | yes |
| u64 | 4096 | build | ordered | hashed | 6 | 1.15 ms | 1.25 ms | 1.08× [1.08, 1.09] | +7.6% | [+7.4%, +7.9%] | 0.2 pts | 0.8 | yes |
| u64 | 4096 | build | ordered | map-sets | 6 | 1.16 ms | 1.53 ms | 1.33× [1.32, 1.34] | +24.6% | [+24.0%, +25.6%] | 0.8 pts | 2.4 | yes |
| u64 | 1048576 | valuesFor | hashed | map-sets | 13 | 178 | 383 | 2.16× [2.13, 2.17] | +53.7% | [+53.1%, +53.8%] | 0.6 pts | 2.3 | yes |
| u64 | 1048576 | valuesFor | ordered | btree-sets | 13 | 184 | 1029 | 5.62× [5.55, 5.64] | +82.2% | [+82.0%, +82.3%] | 0.2 pts | 2.0 | yes |
| u64 | 1048576 | valuesFor | ordered | hashed | 13 | 171 | 176 | 1.02× [1.01, 1.05] | +2.1% | [+0.6%, +4.3%] | 3.1 pts | 5.0 | yes |
| u64 | 1048576 | valuesFor | ordered | map-sets | 13 | 175 | 383 | 2.19× [2.12, 2.23] | +54.3% | [+52.9%, +55.1%] | 1.8 pts | 8.7 | yes |
| u64 | 1048576 | valuesBetween | ordered | btree-sets | 13 | 7380 | 20.4 µs | 2.79× [2.72, 2.82] | +64.1% | [+63.3%, +64.5%] | 1.0 pts | 5.2 | yes |
| u64 | 1048576 | addRemove | hashed | map-sets | 13 | 287 | 493 | 1.72× [1.70, 1.75] | +41.8% | [+41.3%, +42.9%] | 1.3 pts | 5.3 | yes |
| u64 | 1048576 | addRemove | ordered | btree-sets | 13 | 353 | 1103 | 3.11× [3.10, 3.17] | +67.9% | [+67.7%, +68.4%] | 0.6 pts | 1.8 | yes |
| u64 | 1048576 | addRemove | ordered | hashed | 13 | 315 | 281 | 0.89× [0.89, 0.90] | -11.8% | [-12.5%, -11.3%] | 1.0 pts | 1.7 | yes |
| u64 | 1048576 | addRemove | ordered | map-sets | 13 | 322 | 493 | 1.54× [1.50, 1.55] | +34.9% | [+33.3%, +35.5%] | 1.9 pts | 6.3 | yes |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
