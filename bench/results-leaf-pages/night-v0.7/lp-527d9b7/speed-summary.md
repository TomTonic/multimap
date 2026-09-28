| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 109 | 119 | 1.09× [1.08, 1.10] | +8.0% | [+7.4%, +8.7%] | 0.6 pts | 0.8 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4107 | 3810 | 0.93× [0.92, 0.93] | -8.1% | [-9.0%, -7.2%] | 0.8 pts | 0.4 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 336 | 405 | 1.21× [1.19, 1.22] | +17.2% | [+16.2%, +18.2%] | 0.9 pts | 0.4 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 169 | 207 | 1.23× [1.21, 1.24] | +18.6% | [+17.5%, +19.7%] | 1.0 pts | 0.9 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 12.04 ms | 18.86 ms | 1.57× [1.56, 1.58] | +36.2% | [+35.8%, +36.6%] | 0.4 pts | 0.8 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 546 | 463 | 0.85× [0.83, 0.87] | -17.3% | [-19.9%, -14.8%] | 3.6 pts | 2.7 | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 13.6 µs | 9948 | 0.74× [0.73, 0.75] | -35.1% | [-36.7%, -33.6%] | 2.1 pts | 1.3 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 12.3 µs | 9503 | 0.76× [0.74, 0.78] | -31.7% | [-35.8%, -27.5%] | 5.8 pts | 0.4 | no |
| multi | path | 262144 | churn | ordered | baseline | 10 | 889 | 883 | 0.98× [0.95, 1.00] | -2.5% | [-4.9%, -0.2%] | 3.3 pts | 1.5 | no |
| multi | str | 4096 | valuesFor | ordered | baseline | 10 | 62.6 | 70.4 | 1.14× [1.12, 1.16] | +12.2% | [+10.4%, +14.0%] | 2.5 pts | 2.3 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 10 | 3462 | 2761 | 0.80× [0.79, 0.81] | -24.8% | [-26.0%, -23.5%] | 1.7 pts | 0.5 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 10 | 7642 | 5835 | 0.76× [0.74, 0.78] | -31.5% | [-34.7%, -28.3%] | 4.5 pts | 0.8 | no |
| multi | str | 4096 | churn | ordered | baseline | 10 | 82.6 | 127 | 1.53× [1.51, 1.55] | +34.7% | [+33.8%, +35.6%] | 1.3 pts | 1.3 | yes |
| multi | str | 4096 | build | ordered | baseline | 10 | 6.55 ms | 12.57 ms | 1.93× [1.92, 1.94] | +48.2% | [+47.9%, +48.5%] | 0.4 pts | 0.8 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 344 | 298 | 0.87× [0.86, 0.88] | -14.6% | [-15.7%, -13.4%] | 1.6 pts | 1.3 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 11.1 µs | 9138 | 0.82× [0.82, 0.83] | -21.5% | [-22.4%, -20.6%] | 1.3 pts | 1.1 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.76 ms | 1.47 ms | 0.83× [0.82, 0.84] | -21.1% | [-22.5%, -19.6%] | 2.0 pts | 1.1 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 478 | 523 | 1.05× [1.03, 1.08] | +5.0% | [+2.5%, +7.4%] | 3.4 pts | 1.9 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 8 | 55.7 | 64.1 | 1.15× [1.14, 1.16] | +13.3% | [+12.4%, +14.1%] | 1.0 pts | 1.1 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 8 | 2329 | 1734 | 0.74× [0.73, 0.76] | -34.4% | [-36.7%, -32.1%] | 2.8 pts | 1.1 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 8 | 256 | 267 | 1.06× [1.03, 1.08] | +5.3% | [+3.4%, +7.2%] | 2.3 pts | 1.0 | yes |
| multi | street | 4096 | churn | ordered | baseline | 8 | 94.6 | 141 | 1.50× [1.46, 1.53] | +33.2% | [+31.7%, +34.8%] | 1.9 pts | 1.8 | yes |
| multi | street | 4096 | build | ordered | baseline | 8 | 2.81 ms | 6.33 ms | 2.25× [2.25, 2.25] | +55.5% | [+55.5%, +55.6%] | 0.1 pts | 0.3 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 6 | 37.1 | 43.1 | 1.16× [1.15, 1.18] | +14.0% | [+12.8%, +15.2%] | 1.1 pts | 1.0 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 2636 | 2631 | 1.00× [0.98, 1.02] | -0.4% | [-2.3%, +1.5%] | 1.8 pts | 0.6 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 6 | 44.9 | 60.4 | 1.35× [1.33, 1.37] | +25.9% | [+24.8%, +26.9%] | 1.0 pts | 0.9 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 6 | 4.04 ms | 5.71 ms | 1.42× [1.40, 1.43] | +29.3% | [+28.7%, +30.0%] | 0.6 pts | 0.8 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 170 | 157 | 0.92× [0.90, 0.94] | -8.5% | [-10.8%, -6.1%] | 3.2 pts | 2.6 | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 8861 | 6809 | 0.77× [0.76, 0.77] | -30.4% | [-31.0%, -29.9%] | 0.8 pts | 0.5 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 309 | 320 | 1.01× [0.98, 1.05] | +1.3% | [-2.3%, +4.9%] | 5.0 pts | 2.5 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 55.0 | 69.4 | 1.26× [1.26, 1.27] | +20.6% | [+20.3%, +21.0%] | 0.3 pts | 0.4 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3280 | 3647 | 1.11× [1.10, 1.12] | +10.0% | [+9.0%, +11.0%] | 0.9 pts | 0.5 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 88.6 | 130 | 1.47× [1.46, 1.47] | +31.9% | [+31.7%, +32.0%] | 0.1 pts | 0.2 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 73.4 | 119 | 1.61× [1.59, 1.63] | +37.9% | [+37.2%, +38.7%] | 0.7 pts | 0.9 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 5.90 ms | 11.21 ms | 1.90× [1.88, 1.91] | +47.3% | [+46.8%, +47.7%] | 0.4 pts | 0.8 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 295 | 281 | 0.96× [0.95, 0.98] | -3.6% | [-5.3%, -2.0%] | 2.4 pts | 2.4 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 10.7 µs | 8320 | 0.78× [0.77, 0.79] | -27.9% | [-29.1%, -26.7%] | 1.7 pts | 1.0 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 691 | 605 | 0.88× [0.87, 0.88] | -14.1% | [-15.2%, -13.0%] | 1.5 pts | 1.0 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 466 | 502 | 1.04× [0.99, 1.10] | +4.0% | [-1.5%, +9.4%] | 7.6 pts | 3.6 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 10 | 86.7 | 95.2 | 1.09× [1.07, 1.11] | +8.3% | [+6.5%, +10.1%] | 2.5 pts | 2.3 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 10 | 2173 | 1832 | 0.84× [0.83, 0.85] | -18.9% | [-20.7%, -17.0%] | 2.6 pts | 1.5 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 10 | 250 | 310 | 1.24× [1.23, 1.26] | +19.5% | [+18.5%, +20.5%] | 1.4 pts | 0.8 | yes |
| unique | path | 4096 | churn | ordered | baseline | 10 | 205 | 249 | 1.21× [1.19, 1.23] | +17.5% | [+16.3%, +18.8%] | 1.8 pts | 1.3 | yes |
| unique | path | 4096 | build | ordered | baseline | 10 | 2.14 ms | 4.10 ms | 1.91× [1.90, 1.92] | +47.7% | [+47.4%, +47.9%] | 0.4 pts | 0.6 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 505 | 385 | 0.78× [0.76, 0.80] | -28.4% | [-31.7%, -25.1%] | 4.6 pts | 2.9 | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 7840 | 4407 | 0.56× [0.55, 0.56] | -79.9% | [-82.7%, -77.1%] | 3.9 pts | 1.8 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 6456 | 3603 | 0.57× [0.55, 0.59] | -76.2% | [-81.8%, -70.5%] | 7.9 pts | 0.5 | yes |
| unique | path | 262144 | churn | ordered | baseline | 10 | 889 | 813 | 0.92× [0.89, 0.94] | -9.0% | [-12.1%, -6.0%] | 4.2 pts | 1.5 | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 6 | 38.9 | 48.9 | 1.26× [1.25, 1.28] | +20.7% | [+19.8%, +21.6%] | 0.9 pts | 0.9 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 6 | 1636 | 934 | 0.57× [0.56, 0.58] | -74.9% | [-78.7%, -71.1%] | 3.6 pts | 1.3 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 6 | 3150 | 1548 | 0.49× [0.48, 0.50] | -103.1% | [-106.6%, -99.7%] | 3.3 pts | 1.3 | yes |
| unique | str | 4096 | churn | ordered | baseline | 6 | 91.1 | 136 | 1.50× [1.48, 1.53] | +33.4% | [+32.2%, +34.6%] | 1.1 pts | 1.2 | yes |
| unique | str | 4096 | build | ordered | baseline | 6 | 1.05 ms | 2.23 ms | 2.11× [2.08, 2.14] | +52.6% | [+51.9%, +53.2%] | 0.6 pts | 0.8 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 225 | 173 | 0.75× [0.74, 0.76] | -33.3% | [-35.6%, -31.0%] | 3.2 pts | 1.7 | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 5147 | 2911 | 0.56× [0.55, 0.57] | -78.1% | [-82.1%, -74.2%] | 5.5 pts | 2.2 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 814.2 µs | 438.6 µs | 0.53× [0.52, 0.55] | -87.9% | [-93.0%, -82.7%] | 7.2 pts | 0.7 | yes |
| unique | str | 262144 | churn | ordered | baseline | 10 | 335 | 316 | 0.96× [0.93, 0.98] | -4.6% | [-7.1%, -2.1%] | 3.5 pts | 2.3 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 8 | 49.0 | 57.6 | 1.18× [1.17, 1.19] | +15.0% | [+14.4%, +15.7%] | 0.8 pts | 0.9 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 8 | 1819 | 1227 | 0.68× [0.67, 0.68] | -48.0% | [-49.0%, -47.0%] | 1.2 pts | 0.6 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 8 | 213 | 226 | 1.06× [1.04, 1.08] | +5.5% | [+4.1%, +7.0%] | 1.7 pts | 1.0 | yes |
| unique | street | 4096 | churn | ordered | baseline | 8 | 110 | 160 | 1.45× [1.43, 1.47] | +31.0% | [+30.2%, +31.8%] | 0.9 pts | 1.3 | yes |
| unique | street | 4096 | build | ordered | baseline | 8 | 1.26 ms | 3.04 ms | 2.42× [2.39, 2.44] | +58.6% | [+58.2%, +59.1%] | 0.6 pts | 1.0 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 15.5 | 16.4 | 1.07× [1.05, 1.09] | +6.4% | [+4.7%, +8.0%] | 1.6 pts | 0.4 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 1012 | 309 | 0.31× [0.30, 0.31] | -227.2% | [-229.8%, -224.5%] | 2.5 pts | 1.4 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 43.9 | 39.6 | 0.90× [0.90, 0.91] | -10.8% | [-11.3%, -10.4%] | 0.4 pts | 0.4 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 606.0 µs | 508.2 µs | 0.84× [0.83, 0.85] | -19.1% | [-20.0%, -18.3%] | 0.8 pts | 0.6 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 10 | 95.3 | 54.8 | 0.57× [0.57, 0.58] | -74.3% | [-76.8%, -71.8%] | 3.5 pts | 0.8 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 3579 | 1031 | 0.29× [0.28, 0.30] | -246.5% | [-255.8%, -237.2%] | 13.0 pts | 1.4 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 10 | 201 | 155 | 0.78× [0.76, 0.80] | -28.5% | [-32.4%, -24.5%] | 5.5 pts | 0.8 | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 6 | 31.1 | 46.4 | 1.49× [1.47, 1.51] | +32.9% | [+32.0%, +33.7%] | 0.8 pts | 1.0 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1503 | 1804 | 1.20× [1.19, 1.21] | +16.7% | [+16.2%, +17.2%] | 0.5 pts | 0.6 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 66.6 | 109 | 1.63× [1.60, 1.65] | +38.5% | [+37.7%, +39.4%] | 0.8 pts | 1.9 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 6 | 79.4 | 122 | 1.53× [1.51, 1.54] | +34.5% | [+33.9%, +35.2%] | 0.6 pts | 0.8 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 6 | 959.6 µs | 2.19 ms | 2.26× [2.16, 2.36] | +55.7% | [+53.7%, +57.7%] | 1.9 pts | 1.4 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 235 | 210 | 0.88× [0.85, 0.91] | -13.7% | [-17.1%, -10.3%] | 4.7 pts | 3.1 | no |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 5590 | 2879 | 0.51× [0.51, 0.52] | -95.8% | [-97.9%, -93.8%] | 2.9 pts | 1.4 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 429 | 311 | 0.73× [0.72, 0.74] | -37.0% | [-38.8%, -35.2%] | 2.6 pts | 1.6 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 326 | 330 | 1.04× [1.00, 1.08] | +3.7% | [-0.4%, +7.7%] | 5.6 pts | 2.2 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi path n=262144 valuesFor: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled difference of -0.38% does not clear the 2.25% median noise floor of the processes
- multi u64 n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.28%, 1.53%] includes zero
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs baseline: the pooled difference of 1.30% does not clear the 1.67% median noise floor of the processes
- multi u64 n=262144 churn: ordered vs baseline: the pooled interval [-2.31%, 4.90%] includes zero
- multi u64 n=262144 churn: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs baseline: 5 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs baseline: the pooled interval [-1.51%, 9.41%] includes zero
- multi uuid n=262144 churn: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs baseline: 7 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique path n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-0.35%, 7.70%] includes zero
- unique uuid n=262144 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
