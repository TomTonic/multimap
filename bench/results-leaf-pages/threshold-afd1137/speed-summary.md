| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | path | 16384 | valuesFor | ordered | baseline | 8 | 138 | 154 | 1.11× [1.10, 1.13] | +10.2% | [+9.3%, +11.2%] | 1.2 pts | 1.0 | yes |
| multi | path | 16384 | valuesBetween | ordered | baseline | 8 | 4951 | 3766 | 0.76× [0.75, 0.77] | -31.6% | [-32.5%, -30.7%] | 1.0 pts | 0.6 | yes |
| multi | path | 16384 | prefix | ordered | baseline | 8 | 643 | 668 | 1.03× [1.01, 1.05] | +3.2% | [+1.3%, +5.0%] | 2.2 pts | 0.4 | yes |
| multi | path | 16384 | churn | ordered | baseline | 8 | 251 | 286 | 1.14× [1.13, 1.16] | +12.4% | [+11.2%, +13.6%] | 1.5 pts | 1.1 | yes |
| multi | path | 16384 | build | ordered | baseline | 8 | 65.14 ms | 108.98 ms | 1.67× [1.65, 1.68] | +40.0% | [+39.5%, +40.5%] | 0.6 pts | 0.9 | yes |
| multi | path | 32768 | valuesFor | ordered | baseline | 10 | 172 | 184 | 1.04× [1.00, 1.08] | +3.7% | [+0.0%, +7.3%] | 5.1 pts | 2.3 | no |
| multi | path | 32768 | valuesBetween | ordered | baseline | 10 | 5354 | 4219 | 0.78× [0.77, 0.79] | -27.9% | [-29.3%, -26.6%] | 1.9 pts | 1.4 | yes |
| multi | path | 32768 | prefix | ordered | baseline | 10 | 1229 | 1095 | 0.89× [0.87, 0.91] | -12.7% | [-15.4%, -10.0%] | 3.7 pts | 0.4 | no |
| multi | path | 32768 | churn | ordered | baseline | 10 | 368 | 348 | 0.96× [0.95, 0.97] | -4.4% | [-5.7%, -3.1%] | 1.8 pts | 1.2 | yes |
| multi | path | 32768 | build | ordered | baseline | 10 | 184.24 ms | 270.49 ms | 1.48× [1.46, 1.49] | +32.2% | [+31.4%, +33.1%] | 1.1 pts | 1.1 | yes |
| multi | path | 65536 | valuesFor | ordered | baseline | 10 | 312 | 248 | 0.81× [0.79, 0.83] | -24.0% | [-26.9%, -21.1%] | 4.1 pts | 1.8 | no |
| multi | path | 65536 | valuesBetween | ordered | baseline | 10 | 7691 | 5641 | 0.73× [0.71, 0.75] | -37.2% | [-40.7%, -33.7%] | 4.9 pts | 1.0 | yes |
| multi | path | 65536 | prefix | ordered | baseline | 10 | 2295 | 1890 | 0.83× [0.80, 0.86] | -20.9% | [-25.7%, -16.2%] | 6.6 pts | 0.7 | no |
| multi | path | 65536 | churn | ordered | baseline | 10 | 544 | 511 | 0.97× [0.93, 1.01] | -3.2% | [-7.5%, +1.1%] | 6.1 pts | 3.6 | no |
| multi | path | 65536 | build | ordered | baseline | 10 | 544.62 ms | 725.31 ms | 1.32× [1.31, 1.34] | +24.5% | [+23.4%, +25.6%] | 1.6 pts | 1.3 | yes |
| multi | path | 131072 | valuesFor | ordered | baseline | 10 | 446 | 362 | 0.81× [0.80, 0.83] | -22.8% | [-25.2%, -20.4%] | 3.3 pts | 2.1 | no |
| multi | path | 131072 | valuesBetween | ordered | baseline | 10 | 10.8 µs | 7347 | 0.69× [0.68, 0.70] | -45.8% | [-48.0%, -43.5%] | 3.1 pts | 1.2 | yes |
| multi | path | 131072 | prefix | ordered | baseline | 10 | 4398 | 3150 | 0.71× [0.70, 0.73] | -40.3% | [-43.6%, -37.0%] | 4.6 pts | 0.3 | yes |
| multi | path | 131072 | churn | ordered | baseline | 10 | 707 | 704 | 1.02× [0.99, 1.04] | +1.5% | [-1.1%, +4.1%] | 3.6 pts | 2.3 | no |
| multi | str | 16384 | valuesFor | ordered | baseline | 8 | 73.8 | 86.6 | 1.17× [1.16, 1.18] | +14.7% | [+14.0%, +15.3%] | 0.7 pts | 0.8 | yes |
| multi | str | 16384 | valuesBetween | ordered | baseline | 8 | 3930 | 2891 | 0.74× [0.73, 0.74] | -36.0% | [-36.6%, -35.4%] | 0.7 pts | 0.5 | yes |
| multi | str | 16384 | prefix | ordered | baseline | 8 | 37.1 µs | 25.0 µs | 0.67× [0.67, 0.68] | -48.2% | [-49.4%, -47.1%] | 1.4 pts | 0.7 | yes |
| multi | str | 16384 | churn | ordered | baseline | 8 | 113 | 162 | 1.41× [1.36, 1.47] | +29.1% | [+26.4%, +31.9%] | 3.3 pts | 2.4 | yes |
| multi | str | 16384 | build | ordered | baseline | 8 | 32.58 ms | 67.71 ms | 2.07× [2.06, 2.08] | +51.8% | [+51.6%, +52.0%] | 0.3 pts | 0.4 | yes |
| multi | str | 32768 | valuesFor | ordered | baseline | 10 | 81.7 | 96.3 | 1.18× [1.17, 1.19] | +15.1% | [+14.6%, +15.7%] | 0.7 pts | 0.8 | yes |
| multi | str | 32768 | valuesBetween | ordered | baseline | 10 | 4040 | 3183 | 0.79× [0.78, 0.79] | -27.0% | [-27.8%, -26.3%] | 1.1 pts | 0.7 | yes |
| multi | str | 32768 | prefix | ordered | baseline | 10 | 76.6 µs | 56.4 µs | 0.73× [0.73, 0.74] | -36.5% | [-37.8%, -35.1%] | 1.9 pts | 0.9 | yes |
| multi | str | 32768 | churn | ordered | baseline | 10 | 150 | 185 | 1.21× [1.17, 1.25] | +17.4% | [+14.9%, +19.9%] | 3.5 pts | 2.8 | no |
| multi | str | 32768 | build | ordered | baseline | 10 | 86.48 ms | 153.42 ms | 1.74× [1.71, 1.78] | +42.6% | [+41.6%, +43.7%] | 1.4 pts | 1.4 | yes |
| multi | str | 65536 | valuesFor | ordered | baseline | 10 | 149 | 129 | 0.85× [0.78, 0.93] | -18.3% | [-28.8%, -7.8%] | 14.7 pts | 4.1 | no |
| multi | str | 65536 | valuesBetween | ordered | baseline | 10 | 4937 | 3697 | 0.74× [0.73, 0.76] | -34.4% | [-36.5%, -32.2%] | 3.0 pts | 1.3 | yes |
| multi | str | 65536 | prefix | ordered | baseline | 10 | 188.4 µs | 132.4 µs | 0.69× [0.67, 0.71] | -45.8% | [-50.0%, -41.6%] | 5.8 pts | 1.9 | yes |
| multi | str | 65536 | churn | ordered | baseline | 10 | 242 | 240 | 0.98× [0.96, 1.01] | -2.0% | [-4.6%, +0.5%] | 3.6 pts | 1.5 | no |
| multi | str | 65536 | build | ordered | baseline | 10 | 256.07 ms | 371.79 ms | 1.46× [1.45, 1.47] | +31.4% | [+30.8%, +32.0%] | 0.8 pts | 0.7 | yes |
| multi | str | 131072 | valuesFor | ordered | baseline | 10 | 274 | 178 | 0.66× [0.64, 0.69] | -51.0% | [-56.4%, -45.5%] | 7.7 pts | 2.3 | no |
| multi | str | 131072 | valuesBetween | ordered | baseline | 10 | 8261 | 5104 | 0.61× [0.60, 0.62] | -64.6% | [-67.5%, -61.6%] | 4.1 pts | 1.2 | yes |
| multi | str | 131072 | prefix | ordered | baseline | 10 | 622.5 µs | 349.2 µs | 0.56× [0.55, 0.57] | -79.5% | [-83.1%, -75.8%] | 5.1 pts | 0.5 | yes |
| multi | str | 131072 | churn | ordered | baseline | 10 | 352 | 357 | 1.01× [0.97, 1.05] | +1.2% | [-2.6%, +4.9%] | 5.3 pts | 2.2 | no |
| multi | street | 16384 | valuesFor | ordered | baseline | 6 | 76.7 | 83.0 | 1.08× [1.07, 1.09] | +7.4% | [+6.9%, +8.0%] | 0.5 pts | 0.4 | yes |
| multi | street | 16384 | valuesBetween | ordered | baseline | 6 | 2880 | 1626 | 0.56× [0.56, 0.57] | -77.0% | [-78.8%, -75.2%] | 1.7 pts | 0.9 | yes |
| multi | street | 16384 | prefix | ordered | baseline | 6 | 929 | 562 | 0.60× [0.59, 0.61] | -66.5% | [-69.7%, -63.4%] | 3.0 pts | 0.8 | yes |
| multi | street | 16384 | churn | ordered | baseline | 6 | 130 | 179 | 1.37× [1.33, 1.41] | +26.8% | [+24.6%, +29.0%] | 2.1 pts | 1.4 | yes |
| multi | street | 16384 | build | ordered | baseline | 6 | 15.71 ms | 35.31 ms | 2.25× [2.21, 2.28] | +55.5% | [+54.8%, +56.1%] | 0.6 pts | 0.9 | yes |
| multi | street | 32768 | valuesFor | ordered | baseline | 10 | 90.5 | 98.1 | 1.07× [1.03, 1.11] | +6.2% | [+2.9%, +9.5%] | 4.6 pts | 3.7 | no |
| multi | street | 32768 | valuesBetween | ordered | baseline | 10 | 3171 | 1866 | 0.59× [0.58, 0.59] | -70.2% | [-71.8%, -68.5%] | 2.3 pts | 0.8 | yes |
| multi | street | 32768 | prefix | ordered | baseline | 10 | 1995 | 1034 | 0.52× [0.51, 0.53] | -93.1% | [-95.8%, -90.3%] | 3.9 pts | 0.7 | yes |
| multi | street | 32768 | churn | ordered | baseline | 10 | 175 | 202 | 1.15× [1.13, 1.18] | +13.4% | [+11.5%, +15.2%] | 2.6 pts | 1.5 | yes |
| multi | street | 32768 | build | ordered | baseline | 10 | 43.92 ms | 86.21 ms | 1.97× [1.92, 2.01] | +49.2% | [+48.0%, +50.3%] | 1.6 pts | 2.4 | yes |
| multi | street | 65536 | valuesFor | ordered | baseline | 10 | 131 | 118 | 0.86× [0.75, 1.01] | -16.2% | [-33.6%, +1.2%] | 24.3 pts | 8.1 | no |
| multi | street | 65536 | valuesBetween | ordered | baseline | 10 | 3433 | 1998 | 0.57× [0.55, 0.59] | -75.5% | [-81.3%, -69.8%] | 8.1 pts | 2.2 | yes |
| multi | street | 65536 | prefix | ordered | baseline | 10 | 4299 | 2008 | 0.47× [0.46, 0.47] | -114.9% | [-117.5%, -112.2%] | 3.6 pts | 0.6 | yes |
| multi | street | 65536 | churn | ordered | baseline | 10 | 257 | 250 | 0.97× [0.93, 1.00] | -3.4% | [-7.1%, +0.3%] | 5.2 pts | 2.7 | no |
| multi | street | 65536 | build | ordered | baseline | 10 | 114.61 ms | 197.73 ms | 1.73× [1.70, 1.76] | +42.2% | [+41.3%, +43.1%] | 1.3 pts | 1.0 | yes |
| multi | street | 131072 | valuesFor | ordered | baseline | 10 | 234 | 154 | 0.66× [0.63, 0.68] | -52.5% | [-57.9%, -47.1%] | 7.6 pts | 2.2 | no |
| multi | street | 131072 | valuesBetween | ordered | baseline | 10 | 5023 | 2558 | 0.50× [0.50, 0.51] | -98.7% | [-101.1%, -96.3%] | 3.3 pts | 0.7 | yes |
| multi | street | 131072 | prefix | ordered | baseline | 10 | 10.1 µs | 4236 | 0.42× [0.41, 0.42] | -139.8% | [-143.6%, -136.0%] | 5.3 pts | 0.5 | yes |
| multi | street | 131072 | churn | ordered | baseline | 10 | 394 | 361 | 0.95× [0.92, 0.98] | -5.3% | [-9.1%, -1.6%] | 5.2 pts | 1.9 | no |
| multi | u64 | 16384 | valuesFor | ordered | baseline | 6 | 43.9 | 56.5 | 1.29× [1.28, 1.29] | +22.2% | [+21.8%, +22.7%] | 0.4 pts | 0.7 | yes |
| multi | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 3780 | 2444 | 0.65× [0.64, 0.65] | -54.1% | [-55.3%, -53.0%] | 1.1 pts | 0.6 | yes |
| multi | u64 | 16384 | churn | ordered | baseline | 6 | 52.2 | 94.8 | 1.80× [1.75, 1.85] | +44.4% | [+43.0%, +45.9%] | 1.4 pts | 1.8 | yes |
| multi | u64 | 16384 | build | ordered | baseline | 6 | 19.08 ms | 34.66 ms | 1.81× [1.78, 1.83] | +44.6% | [+43.9%, +45.4%] | 0.7 pts | 1.0 | yes |
| multi | u64 | 32768 | valuesFor | ordered | baseline | 8 | 44.9 | 63.3 | 1.40× [1.38, 1.41] | +28.4% | [+27.7%, +29.1%] | 0.8 pts | 1.1 | yes |
| multi | u64 | 32768 | valuesBetween | ordered | baseline | 8 | 4466 | 2639 | 0.59× [0.58, 0.60] | -69.7% | [-71.5%, -68.0%] | 2.1 pts | 1.2 | yes |
| multi | u64 | 32768 | churn | ordered | baseline | 8 | 72.5 | 107 | 1.46× [1.40, 1.52] | +31.4% | [+28.4%, +34.3%] | 3.6 pts | 3.0 | yes |
| multi | u64 | 32768 | build | ordered | baseline | 8 | 49.32 ms | 82.67 ms | 1.68× [1.66, 1.69] | +40.3% | [+39.8%, +40.9%] | 0.6 pts | 0.7 | yes |
| multi | u64 | 65536 | valuesFor | ordered | baseline | 10 | 60.0 | 71.4 | 1.15× [1.09, 1.22] | +13.1% | [+8.0%, +18.2%] | 7.1 pts | 3.8 | no |
| multi | u64 | 65536 | valuesBetween | ordered | baseline | 10 | 4702 | 2928 | 0.62× [0.61, 0.62] | -62.1% | [-63.8%, -60.5%] | 2.3 pts | 0.5 | yes |
| multi | u64 | 65536 | churn | ordered | baseline | 10 | 125 | 133 | 1.06× [1.01, 1.12] | +5.6% | [+0.7%, +10.6%] | 7.0 pts | 5.4 | no |
| multi | u64 | 65536 | build | ordered | baseline | 10 | 147.00 ms | 209.28 ms | 1.41× [1.38, 1.44] | +29.1% | [+27.6%, +30.6%] | 2.1 pts | 2.1 | yes |
| multi | u64 | 131072 | valuesFor | ordered | baseline | 10 | 111 | 93.9 | 0.83× [0.81, 0.86] | -20.2% | [-23.6%, -16.7%] | 4.8 pts | 2.7 | no |
| multi | u64 | 131072 | valuesBetween | ordered | baseline | 10 | 6634 | 3962 | 0.60× [0.59, 0.61] | -65.7% | [-68.6%, -62.8%] | 4.1 pts | 0.7 | yes |
| multi | u64 | 131072 | churn | ordered | baseline | 10 | 204 | 193 | 0.94× [0.89, 0.99] | -6.6% | [-12.7%, -0.6%] | 8.5 pts | 3.4 | no |
| multi | uuid | 16384 | valuesFor | ordered | baseline | 6 | 64.8 | 85.0 | 1.31× [1.30, 1.32] | +23.7% | [+23.1%, +24.3%] | 0.6 pts | 0.6 | yes |
| multi | uuid | 16384 | valuesBetween | ordered | baseline | 6 | 3851 | 2963 | 0.77× [0.76, 0.77] | -30.0% | [-31.0%, -29.1%] | 0.9 pts | 0.6 | yes |
| multi | uuid | 16384 | prefix | ordered | baseline | 6 | 115 | 217 | 1.90× [1.89, 1.91] | +47.3% | [+46.9%, +47.7%] | 0.3 pts | 0.5 | yes |
| multi | uuid | 16384 | churn | ordered | baseline | 6 | 107 | 161 | 1.49× [1.45, 1.54] | +33.0% | [+31.0%, +35.0%] | 1.9 pts | 1.8 | yes |
| multi | uuid | 16384 | build | ordered | baseline | 6 | 31.92 ms | 67.16 ms | 2.10× [2.04, 2.17] | +52.4% | [+51.0%, +53.8%] | 1.4 pts | 1.0 | yes |
| multi | uuid | 32768 | valuesFor | ordered | baseline | 10 | 77.6 | 105 | 1.24× [1.12, 1.40] | +19.5% | [+10.6%, +28.4%] | 12.5 pts | 8.5 | no |
| multi | uuid | 32768 | valuesBetween | ordered | baseline | 10 | 4094 | 3433 | 0.84× [0.83, 0.85] | -19.1% | [-20.2%, -18.1%] | 1.5 pts | 1.0 | yes |
| multi | uuid | 32768 | prefix | ordered | baseline | 10 | 141 | 253 | 1.73× [1.63, 1.85] | +42.3% | [+38.5%, +46.0%] | 5.3 pts | 7.0 | yes |
| multi | uuid | 32768 | churn | ordered | baseline | 10 | 160 | 197 | 1.21× [1.16, 1.27] | +17.3% | [+13.7%, +21.0%] | 5.1 pts | 3.1 | no |
| multi | uuid | 32768 | build | ordered | baseline | 10 | 86.75 ms | 167.00 ms | 1.93× [1.90, 1.97] | +48.3% | [+47.4%, +49.2%] | 1.3 pts | 1.3 | yes |
| multi | uuid | 65536 | valuesFor | ordered | baseline | 10 | 141 | 132 | 0.93× [0.89, 0.96] | -7.8% | [-11.7%, -3.8%] | 5.6 pts | 2.9 | no |
| multi | uuid | 65536 | valuesBetween | ordered | baseline | 10 | 5551 | 4079 | 0.73× [0.72, 0.74] | -37.7% | [-39.5%, -35.9%] | 2.5 pts | 1.0 | yes |
| multi | uuid | 65536 | prefix | ordered | baseline | 10 | 288 | 322 | 1.14× [1.12, 1.16] | +12.1% | [+10.6%, +13.6%] | 2.1 pts | 1.6 | yes |
| multi | uuid | 65536 | churn | ordered | baseline | 10 | 266 | 276 | 1.04× [1.01, 1.06] | +3.7% | [+1.4%, +6.1%] | 3.2 pts | 1.9 | no |
| multi | uuid | 65536 | build | ordered | baseline | 10 | 259.83 ms | 424.17 ms | 1.64× [1.61, 1.68] | +39.1% | [+37.8%, +40.5%] | 1.9 pts | 1.6 | yes |
| multi | uuid | 131072 | valuesFor | ordered | baseline | 10 | 258 | 250 | 0.94× [0.88, 1.01] | -6.3% | [-13.9%, +1.2%] | 10.5 pts | 5.0 | no |
| multi | uuid | 131072 | valuesBetween | ordered | baseline | 10 | 9257 | 5543 | 0.60× [0.59, 0.61] | -66.5% | [-68.9%, -64.2%] | 3.3 pts | 1.1 | yes |
| multi | uuid | 131072 | prefix | ordered | baseline | 10 | 433 | 450 | 1.05× [1.04, 1.06] | +5.0% | [+4.0%, +6.1%] | 1.5 pts | 1.4 | yes |
| multi | uuid | 131072 | churn | ordered | baseline | 10 | 376 | 419 | 1.11× [1.08, 1.15] | +10.0% | [+7.3%, +12.7%] | 3.8 pts | 1.9 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi path n=16384 prefix: ordered vs baseline: the pooled difference of 3.16% does not clear the 5.31% median noise floor of the processes
- multi path n=32768 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=32768 valuesFor: ordered vs baseline: 7 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi path n=65536 churn: ordered vs baseline: the pooled interval [-7.53%, 1.13%] includes zero
- multi path n=65536 churn: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=131072 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=131072 churn: ordered vs baseline: the pooled difference of 1.52% does not clear the 1.65% median noise floor of the processes
- multi path n=131072 churn: ordered vs baseline: the pooled interval [-1.07%, 4.11%] includes zero
- multi path n=131072 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi path n=131072 churn: ordered vs baseline: 2 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi str n=16384 churn: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=32768 churn: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=65536 valuesFor: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=65536 churn: ordered vs baseline: the pooled interval [-4.59%, 0.50%] includes zero
- multi str n=131072 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=131072 churn: ordered vs baseline: the pooled difference of 1.17% does not clear the 2.25% median noise floor of the processes
- multi str n=131072 churn: ordered vs baseline: the pooled interval [-2.59%, 4.94%] includes zero
- multi str n=131072 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=131072 churn: ordered vs baseline: 1 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=32768 valuesFor: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=32768 valuesFor: ordered vs baseline: 8 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=32768 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=65536 valuesFor: ordered vs baseline: the pooled interval [-33.58%, 1.16%] includes zero
- multi street n=65536 valuesFor: ordered vs baseline: the processes scatter 8.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=65536 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 5 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=65536 churn: ordered vs baseline: the pooled interval [-7.15%, 0.34%] includes zero
- multi street n=65536 churn: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi street n=65536 churn: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi street n=131072 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=32768 churn: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=65536 churn: ordered vs baseline: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=65536 churn: ordered vs baseline: 8 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=65536 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=131072 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=131072 churn: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=131072 churn: ordered vs baseline: 1 processes resolved A as faster and 7 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=32768 valuesFor: ordered vs baseline: the processes scatter 8.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=32768 prefix: ordered vs baseline: the processes scatter 7.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=32768 churn: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=65536 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=131072 valuesFor: ordered vs baseline: the pooled interval [-13.85%, 1.22%] includes zero
- multi uuid n=131072 valuesFor: ordered vs baseline: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=131072 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
