| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 109 | 119 | 1.09× [1.08, 1.10] | +8.1% | [+7.5%, +8.8%] | 0.6 pts | 0.6 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4123 | 3872 | 0.94× [0.93, 0.95] | -6.2% | [-7.7%, -4.8%] | 1.4 pts | 0.7 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 331 | 394 | 1.20× [1.17, 1.23] | +16.5% | [+14.7%, +18.4%] | 1.7 pts | 1.0 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 170 | 208 | 1.24× [1.21, 1.26] | +19.1% | [+17.3%, +20.9%] | 1.7 pts | 1.4 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 12.02 ms | 18.86 ms | 1.56× [1.54, 1.58] | +35.9% | [+35.0%, +36.8%] | 0.9 pts | 1.3 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 598 | 516 | 0.86× [0.84, 0.88] | -16.3% | [-18.7%, -13.8%] | 3.4 pts | 1.8 | no |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 13.6 µs | 10.2 µs | 0.74× [0.73, 0.75] | -34.4% | [-36.2%, -32.7%] | 2.5 pts | 1.7 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 13.0 µs | 9831 | 0.76× [0.74, 0.78] | -30.9% | [-34.4%, -27.4%] | 4.9 pts | 0.3 | no |
| multi | path | 262144 | churn | ordered | baseline | 10 | 945 | 929 | 1.00× [0.98, 1.02] | -0.3% | [-2.2%, +1.6%] | 2.7 pts | 1.2 | yes |
| multi | str | 4096 | valuesFor | ordered | baseline | 8 | 62.6 | 70.9 | 1.13× [1.13, 1.14] | +11.7% | [+11.1%, +12.2%] | 0.6 pts | 0.9 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 8 | 3414 | 2817 | 0.83× [0.81, 0.84] | -21.0% | [-22.9%, -19.1%] | 2.3 pts | 0.7 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 8 | 7528 | 5918 | 0.79× [0.77, 0.80] | -27.0% | [-29.4%, -24.5%] | 2.9 pts | 0.6 | yes |
| multi | str | 4096 | churn | ordered | baseline | 8 | 83.6 | 127 | 1.51× [1.48, 1.55] | +33.9% | [+32.5%, +35.4%] | 1.7 pts | 1.6 | yes |
| multi | str | 4096 | build | ordered | baseline | 8 | 6.50 ms | 12.73 ms | 1.95× [1.95, 1.96] | +48.8% | [+48.7%, +49.0%] | 0.2 pts | 0.4 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 340 | 296 | 0.87× [0.86, 0.88] | -15.0% | [-16.5%, -13.4%] | 2.2 pts | 1.4 | yes |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 11.0 µs | 9203 | 0.83× [0.83, 0.84] | -19.9% | [-20.7%, -19.2%] | 1.0 pts | 0.8 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.76 ms | 1.48 ms | 0.85× [0.84, 0.85] | -18.3% | [-19.3%, -17.2%] | 1.5 pts | 1.2 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 483 | 514 | 1.09× [1.03, 1.16] | +8.4% | [+2.8%, +13.9%] | 7.7 pts | 5.4 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 6 | 56.3 | 64.5 | 1.15× [1.14, 1.17] | +13.1% | [+12.1%, +14.2%] | 1.0 pts | 0.9 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 6 | 2285 | 1747 | 0.76× [0.75, 0.77] | -31.3% | [-33.0%, -29.6%] | 1.6 pts | 0.6 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 6 | 249 | 269 | 1.09× [1.07, 1.10] | +7.9% | [+6.7%, +9.1%] | 1.1 pts | 0.5 | yes |
| multi | street | 4096 | churn | ordered | baseline | 6 | 93.8 | 141 | 1.51× [1.50, 1.52] | +33.7% | [+33.2%, +34.1%] | 0.4 pts | 0.5 | yes |
| multi | street | 4096 | build | ordered | baseline | 6 | 2.81 ms | 6.38 ms | 2.27× [2.26, 2.29] | +56.0% | [+55.7%, +56.3%] | 0.2 pts | 0.7 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 6 | 36.6 | 43.3 | 1.18× [1.16, 1.21] | +15.5% | [+13.7%, +17.2%] | 1.7 pts | 1.2 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 2591 | 2656 | 1.03× [1.02, 1.04] | +2.9% | [+1.8%, +4.0%] | 1.0 pts | 0.4 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 6 | 45.2 | 61.5 | 1.36× [1.34, 1.37] | +26.3% | [+25.6%, +27.0%] | 0.7 pts | 0.6 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 6 | 4.04 ms | 5.83 ms | 1.45× [1.44, 1.45] | +30.8% | [+30.4%, +31.3%] | 0.4 pts | 0.9 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 170 | 154 | 0.92× [0.90, 0.94] | -8.8% | [-11.1%, -6.5%] | 3.2 pts | 2.3 | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 8756 | 6728 | 0.77× [0.76, 0.78] | -29.5% | [-31.5%, -27.5%] | 2.8 pts | 1.8 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 290 | 302 | 1.00× [0.95, 1.05] | -0.3% | [-5.7%, +5.1%] | 7.6 pts | 4.1 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 6 | 56.4 | 71.4 | 1.29× [1.25, 1.33] | +22.4% | [+20.2%, +24.7%] | 2.1 pts | 2.5 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 3303 | 3753 | 1.14× [1.11, 1.16] | +12.1% | [+10.3%, +14.0%] | 1.8 pts | 1.2 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 6 | 90.8 | 132 | 1.47× [1.44, 1.50] | +31.9% | [+30.6%, +33.3%] | 1.3 pts | 3.1 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 6 | 76.7 | 125 | 1.61× [1.57, 1.64] | +37.8% | [+36.4%, +39.2%] | 1.3 pts | 1.2 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 6 | 6.08 ms | 11.54 ms | 1.90× [1.89, 1.92] | +47.5% | [+47.1%, +47.8%] | 0.3 pts | 1.2 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 330 | 315 | 0.96× [0.95, 0.97] | -4.1% | [-5.6%, -2.7%] | 2.1 pts | 2.1 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 11.0 µs | 8540 | 0.79× [0.78, 0.79] | -27.3% | [-28.6%, -25.9%] | 1.8 pts | 1.4 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 727 | 635 | 0.88× [0.87, 0.88] | -14.2% | [-15.1%, -13.3%] | 1.2 pts | 0.9 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 494 | 541 | 1.09× [1.06, 1.12] | +8.1% | [+5.6%, +10.6%] | 3.5 pts | 2.3 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 10 | 87.3 | 95.9 | 1.10× [1.09, 1.11] | +9.1% | [+8.6%, +9.5%] | 0.6 pts | 0.6 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 10 | 2130 | 1823 | 0.85× [0.84, 0.86] | -17.4% | [-18.5%, -16.4%] | 1.5 pts | 0.9 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 10 | 251 | 309 | 1.23× [1.22, 1.24] | +18.8% | [+17.9%, +19.6%] | 1.1 pts | 0.6 | yes |
| unique | path | 4096 | churn | ordered | baseline | 10 | 204 | 251 | 1.22× [1.19, 1.25] | +18.0% | [+16.2%, +19.9%] | 2.6 pts | 1.4 | yes |
| unique | path | 4096 | build | ordered | baseline | 10 | 2.16 ms | 4.15 ms | 1.92× [1.91, 1.93] | +47.9% | [+47.6%, +48.3%] | 0.5 pts | 0.9 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 523 | 405 | 0.78× [0.76, 0.81] | -27.7% | [-31.4%, -24.0%] | 5.2 pts | 2.9 | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 8024 | 4507 | 0.56× [0.55, 0.56] | -80.0% | [-82.4%, -77.6%] | 3.4 pts | 1.2 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 6872 | 3834 | 0.56× [0.54, 0.57] | -79.8% | [-85.6%, -74.0%] | 8.1 pts | 0.5 | yes |
| unique | path | 262144 | churn | ordered | baseline | 10 | 921 | 817 | 0.90× [0.88, 0.92] | -11.3% | [-14.0%, -8.6%] | 3.8 pts | 1.8 | no |
| unique | str | 4096 | valuesFor | ordered | baseline | 6 | 38.1 | 48.3 | 1.26× [1.26, 1.27] | +20.9% | [+20.3%, +21.4%] | 0.5 pts | 0.7 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 6 | 1554 | 922 | 0.59× [0.59, 0.59] | -69.7% | [-70.9%, -68.4%] | 1.2 pts | 0.5 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 6 | 3008 | 1546 | 0.51× [0.50, 0.52] | -94.8% | [-98.6%, -90.9%] | 3.6 pts | 1.3 | yes |
| unique | str | 4096 | churn | ordered | baseline | 6 | 89.5 | 137 | 1.53× [1.51, 1.55] | +34.7% | [+34.0%, +35.5%] | 0.7 pts | 1.0 | yes |
| unique | str | 4096 | build | ordered | baseline | 6 | 1.06 ms | 2.24 ms | 2.09× [2.04, 2.14] | +52.1% | [+51.0%, +53.3%] | 1.1 pts | 1.3 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 266 | 198 | 0.74× [0.73, 0.75] | -34.8% | [-36.4%, -33.2%] | 2.2 pts | 1.3 | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 6032 | 3368 | 0.56× [0.55, 0.57] | -78.6% | [-82.0%, -75.3%] | 4.7 pts | 1.8 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 942.9 µs | 505.0 µs | 0.54× [0.53, 0.55] | -86.4% | [-89.8%, -83.0%] | 4.7 pts | 0.9 | yes |
| unique | str | 262144 | churn | ordered | baseline | 10 | 344 | 339 | 0.98× [0.95, 1.02] | -1.8% | [-5.8%, +2.2%] | 5.6 pts | 1.7 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 50.6 | 58.7 | 1.16× [1.15, 1.18] | +13.9% | [+13.0%, +14.9%] | 0.9 pts | 1.2 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 1794 | 1229 | 0.69× [0.68, 0.70] | -45.6% | [-47.8%, -43.5%] | 2.1 pts | 0.9 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 212 | 227 | 1.08× [1.06, 1.09] | +7.2% | [+5.7%, +8.6%] | 1.4 pts | 0.9 | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 114 | 163 | 1.44× [1.42, 1.45] | +30.4% | [+29.7%, +31.2%] | 0.7 pts | 0.8 | yes |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.29 ms | 3.13 ms | 2.43× [2.41, 2.45] | +58.9% | [+58.5%, +59.2%] | 0.3 pts | 0.6 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 10 | 15.0 | 17.3 | 1.15× [1.11, 1.19] | +12.9% | [+9.5%, +16.2%] | 4.7 pts | 1.4 | no |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 10 | 942 | 309 | 0.33× [0.32, 0.33] | -206.2% | [-208.9%, -203.5%] | 3.8 pts | 1.9 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 10 | 43.7 | 40.9 | 0.94× [0.93, 0.94] | -6.6% | [-7.3%, -5.9%] | 0.9 pts | 1.1 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 10 | 608.1 µs | 522.7 µs | 0.86× [0.85, 0.87] | -16.2% | [-17.4%, -15.1%] | 1.6 pts | 0.9 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 10 | 83.3 | 43.8 | 0.53× [0.52, 0.55] | -87.2% | [-92.0%, -82.3%] | 6.8 pts | 2.2 | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 2737 | 874 | 0.32× [0.31, 0.33] | -215.8% | [-224.7%, -206.8%] | 12.5 pts | 2.1 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 10 | 182 | 131 | 0.74× [0.72, 0.76] | -35.5% | [-39.5%, -31.6%] | 5.5 pts | 1.5 | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 6 | 31.0 | 47.4 | 1.52× [1.48, 1.57] | +34.4% | [+32.3%, +36.4%] | 1.9 pts | 1.9 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1483 | 1818 | 1.23× [1.22, 1.24] | +18.6% | [+17.9%, +19.4%] | 0.7 pts | 0.9 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 67.3 | 111 | 1.66× [1.62, 1.70] | +39.6% | [+38.1%, +41.1%] | 1.4 pts | 2.9 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 6 | 85.8 | 126 | 1.51× [1.44, 1.58] | +33.7% | [+30.5%, +36.9%] | 3.1 pts | 2.2 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 6 | 980.1 µs | 2.26 ms | 2.31× [2.29, 2.33] | +56.7% | [+56.3%, +57.1%] | 0.4 pts | 0.7 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 234 | 203 | 0.87× [0.85, 0.89] | -14.7% | [-17.5%, -11.9%] | 3.9 pts | 3.4 | no |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 5420 | 2848 | 0.51× [0.51, 0.52] | -94.9% | [-97.8%, -92.1%] | 4.0 pts | 1.3 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 402 | 294 | 0.73× [0.72, 0.74] | -36.7% | [-38.9%, -34.5%] | 3.1 pts | 1.9 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 334 | 342 | 1.02× [0.99, 1.06] | +2.4% | [-1.2%, +5.9%] | 4.9 pts | 3.2 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi path n=262144 churn: ordered vs baseline: the pooled difference of -0.33% does not clear the 3.49% median noise floor of the processes
- multi path n=262144 churn: ordered vs baseline: the pooled interval [-2.23%, 1.58%] includes zero
- multi str n=262144 churn: ordered vs baseline: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 churn: ordered vs baseline: 9 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs baseline: the pooled difference of -0.33% does not clear the 2.09% median noise floor of the processes
- multi u64 n=262144 churn: ordered vs baseline: the pooled interval [-5.74%, 5.07%] includes zero
- multi u64 n=262144 churn: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=4096 valuesFor: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=4096 prefix: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi uuid n=262144 churn: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 churn: ordered vs baseline: the pooled difference of -1.79% does not clear the 2.67% median noise floor of the processes
- unique str n=262144 churn: ordered vs baseline: the pooled interval [-5.79%, 2.20%] includes zero
- unique str n=262144 churn: ordered vs baseline: 1 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 prefix: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-1.16%, 5.91%] includes zero
- unique uuid n=262144 churn: ordered vs baseline: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 churn: ordered vs baseline: 5 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
