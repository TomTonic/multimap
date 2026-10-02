| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique | email | 4096 | valuesFor | ordered | baseline | 6 | 34.2 | 28.4 | 0.83× [0.82, 0.84] | -20.3% | [-22.2%, -18.4%] | 1.8 pts | 3.6 | yes | yes |
| unique | email | 4096 | valuesBetween | ordered | baseline | 6 | 347 | 1269 | 3.68× [3.61, 3.74] | +72.8% | [+72.3%, +73.3%] | 0.5 pts | 2.5 | yes | yes |
| unique | email | 4096 | prefix | ordered | baseline | 6 | 59.6 | 57.3 | 0.96× [0.95, 0.97] | -4.5% | [-5.3%, -3.6%] | 0.8 pts | 1.2 | yes | yes |
| unique | email | 4096 | churn | ordered | baseline | 6 | 89.5 | 73.0 | 0.81× [0.81, 0.82] | -22.9% | [-23.8%, -22.1%] | 0.8 pts | 0.9 | yes | yes |
| unique | email | 4096 | build | ordered | baseline | 6 | 1.51 ms | 1.09 ms | 0.72× [0.71, 0.73] | -38.5% | [-40.1%, -36.9%] | 1.5 pts | 1.8 | yes | yes |
| unique | email | 16384 | valuesFor | ordered | baseline | 6 | 46.2 | 37.5 | 0.81× [0.80, 0.81] | -23.8% | [-24.8%, -22.9%] | 0.9 pts | 1.2 | yes | yes |
| unique | email | 16384 | valuesBetween | ordered | baseline | 6 | 441 | 1603 | 3.63× [3.59, 3.66] | +72.4% | [+72.2%, +72.7%] | 0.2 pts | 1.4 | yes | yes |
| unique | email | 16384 | prefix | ordered | baseline | 6 | 79.9 | 72.6 | 0.91× [0.91, 0.91] | -9.9% | [-10.4%, -9.4%] | 0.5 pts | 1.0 | yes | yes |
| unique | email | 16384 | churn | ordered | baseline | 6 | 114 | 91.4 | 0.80× [0.80, 0.81] | -24.6% | [-25.6%, -23.6%] | 1.0 pts | 1.1 | yes | yes |
| unique | email | 16384 | build | ordered | baseline | 6 | 7.25 ms | 5.27 ms | 0.72× [0.72, 0.73] | -38.2% | [-39.0%, -37.5%] | 0.7 pts | 0.6 | yes | yes |
| unique | email | 262144 | valuesFor | ordered | baseline | 12 | 135 | 155 | 1.13× [1.11, 1.15] | +11.7% | [+10.2%, +13.2%] | 3.0 pts | 2.1 | yes | yes |
| unique | email | 262144 | valuesBetween | ordered | baseline | 12 | 1065 | 2995 | 2.79× [2.72, 2.87] | +64.2% | [+63.2%, +65.1%] | 1.9 pts | 1.1 | yes | yes |
| unique | email | 262144 | prefix | ordered | baseline | 12 | 163 | 203 | 1.24× [1.22, 1.26] | +19.4% | [+18.1%, +20.6%] | 1.7 pts | 2.0 | yes | yes |
| unique | email | 262144 | churn | ordered | baseline | 12 | 297 | 278 | 0.87× [0.83, 0.91] | -15.0% | [-19.8%, -10.3%] | 5.6 pts | 1.1 | no | yes |
| unique | path | 4096 | valuesFor | ordered | baseline | 6 | 95.6 | 79.1 | 0.83× [0.82, 0.84] | -20.4% | [-21.6%, -19.2%] | 1.2 pts | 1.3 | yes | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 6 | 786 | 2013 | 2.57× [2.52, 2.62] | +61.1% | [+60.4%, +61.8%] | 0.7 pts | 1.3 | yes | yes |
| unique | path | 4096 | prefix | ordered | baseline | 6 | 258 | 232 | 0.91× [0.90, 0.91] | -10.5% | [-11.0%, -10.0%] | 0.5 pts | 0.5 | yes | yes |
| unique | path | 4096 | churn | ordered | baseline | 6 | 205 | 209 | 1.03× [1.02, 1.05] | +3.3% | [+2.0%, +4.6%] | 1.3 pts | 1.5 | yes | yes |
| unique | path | 4096 | build | ordered | baseline | 6 | 3.66 ms | 2.82 ms | 0.77× [0.76, 0.78] | -30.3% | [-31.9%, -28.7%] | 1.5 pts | 1.4 | yes | yes |
| unique | path | 16384 | valuesFor | ordered | baseline | 6 | 123 | 110 | 0.90× [0.89, 0.91] | -11.4% | [-12.6%, -10.2%] | 1.1 pts | 1.3 | yes | yes |
| unique | path | 16384 | valuesBetween | ordered | baseline | 6 | 963 | 2569 | 2.68× [2.65, 2.71] | +62.7% | [+62.2%, +63.1%] | 0.4 pts | 1.4 | yes | yes |
| unique | path | 16384 | prefix | ordered | baseline | 6 | 359 | 475 | 1.31× [1.28, 1.35] | +23.9% | [+22.0%, +25.7%] | 1.8 pts | 1.0 | yes | yes |
| unique | path | 16384 | churn | ordered | baseline | 6 | 256 | 260 | 1.02× [1.01, 1.03] | +2.0% | [+0.9%, +3.2%] | 1.1 pts | 1.4 | yes | yes |
| unique | path | 16384 | build | ordered | baseline | 6 | 17.28 ms | 13.57 ms | 0.79× [0.78, 0.80] | -26.9% | [-29.0%, -24.8%] | 2.0 pts | 1.5 | yes | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 6 | 310 | 335 | 1.10× [1.08, 1.12] | +9.1% | [+7.8%, +10.4%] | 1.2 pts | 1.0 | yes | yes |
| unique | path | 262144 | valuesBetween | ordered | baseline | 6 | 2169 | 5791 | 2.57× [2.43, 2.72] | +61.0% | [+58.8%, +63.3%] | 2.1 pts | 2.5 | yes | yes |
| unique | path | 262144 | prefix | ordered | baseline | 6 | 1738 | 5381 | 2.90× [2.75, 3.07] | +65.6% | [+63.7%, +67.4%] | 1.8 pts | 0.6 | yes | yes |
| unique | path | 262144 | churn | ordered | baseline | 6 | 598 | 669 | 1.08× [1.06, 1.09] | +7.2% | [+5.7%, +8.6%] | 1.4 pts | 0.3 | yes | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 6 | 43.7 | 39.7 | 0.91× [0.90, 0.92] | -10.0% | [-11.0%, -9.0%] | 1.0 pts | 2.3 | yes | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 6 | 368 | 1606 | 4.37× [4.33, 4.40] | +77.1% | [+76.9%, +77.3%] | 0.2 pts | 1.1 | yes | yes |
| unique | str | 4096 | prefix | ordered | baseline | 6 | 568 | 3191 | 5.60× [5.56, 5.63] | +82.1% | [+82.0%, +82.3%] | 0.1 pts | 2.0 | yes | yes |
| unique | str | 4096 | churn | ordered | baseline | 6 | 98.4 | 104 | 1.06× [1.05, 1.07] | +5.7% | [+4.9%, +6.6%] | 0.8 pts | 1.3 | yes | yes |
| unique | str | 4096 | build | ordered | baseline | 6 | 1.67 ms | 1.48 ms | 0.89× [0.88, 0.89] | -12.7% | [-13.6%, -11.9%] | 0.8 pts | 1.1 | yes | yes |
| unique | str | 16384 | valuesFor | ordered | baseline | 12 | 56.2 | 48.4 | 0.86× [0.85, 0.87] | -16.3% | [-17.3%, -15.3%] | 1.2 pts | 1.9 | yes | yes |
| unique | str | 16384 | valuesBetween | ordered | baseline | 12 | 452 | 1657 | 3.66× [3.64, 3.68] | +72.7% | [+72.6%, +72.8%] | 0.2 pts | 0.9 | yes | yes |
| unique | str | 16384 | prefix | ordered | baseline | 12 | 2419 | 13.5 µs | 5.60× [5.56, 5.64] | +82.1% | [+82.0%, +82.3%] | 0.2 pts | 1.6 | yes | yes |
| unique | str | 16384 | churn | ordered | baseline | 12 | 127 | 122 | 0.96× [0.95, 0.98] | -3.6% | [-5.6%, -1.7%] | 2.0 pts | 2.6 | yes | yes |
| unique | str | 16384 | build | ordered | baseline | 12 | 8.11 ms | 6.89 ms | 0.85× [0.84, 0.85] | -18.1% | [-18.8%, -17.4%] | 0.8 pts | 1.1 | yes | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 105 | 139 | 1.29× [1.25, 1.32] | +22.2% | [+20.2%, +24.3%] | 2.5 pts | 2.1 | yes | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 632 | 2465 | 3.83× [3.70, 3.97] | +73.9% | [+73.0%, +74.8%] | 0.9 pts | 0.7 | yes | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 59.4 µs | 334.2 µs | 5.47× [5.23, 5.73] | +81.7% | [+80.9%, +82.6%] | 0.9 pts | 0.6 | yes | yes |
| unique | str | 262144 | churn | ordered | baseline | 10 | 308 | 366 | 1.15× [1.13, 1.17] | +13.2% | [+11.6%, +14.8%] | 3.3 pts | 1.3 | yes | yes |
| unique | street | 4096 | valuesFor | ordered | baseline | 12 | 47.9 | 46.6 | 0.97× [0.97, 0.97] | -3.0% | [-3.4%, -2.6%] | 0.7 pts | 1.4 | yes | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 12 | 404 | 1781 | 4.40× [4.37, 4.43] | +77.3% | [+77.1%, +77.5%] | 0.3 pts | 1.3 | yes | yes |
| unique | street | 4096 | prefix | ordered | baseline | 12 | 102 | 196 | 1.93× [1.91, 1.94] | +48.1% | [+47.6%, +48.6%] | 0.5 pts | 1.8 | yes | yes |
| unique | street | 4096 | churn | ordered | baseline | 12 | 97.7 | 118 | 1.21× [1.18, 1.24] | +17.5% | [+15.3%, +19.6%] | 2.2 pts | 2.8 | no | yes |
| unique | street | 4096 | build | ordered | baseline | 12 | 1.81 ms | 1.72 ms | 0.94× [0.93, 0.95] | -6.4% | [-7.9%, -4.8%] | 2.1 pts | 1.6 | yes | yes |
| unique | street | 16384 | valuesFor | ordered | baseline | 6 | 60.3 | 65.1 | 1.08× [1.06, 1.09] | +7.3% | [+6.0%, +8.6%] | 1.2 pts | 1.9 | yes | yes |
| unique | street | 16384 | valuesBetween | ordered | baseline | 6 | 476 | 2203 | 4.64× [4.58, 4.70] | +78.4% | [+78.2%, +78.7%] | 0.3 pts | 1.5 | yes | yes |
| unique | street | 16384 | prefix | ordered | baseline | 6 | 177 | 627 | 3.52× [3.51, 3.54] | +71.6% | [+71.5%, +71.7%] | 0.1 pts | 0.5 | yes | yes |
| unique | street | 16384 | churn | ordered | baseline | 6 | 125 | 156 | 1.25× [1.23, 1.27] | +19.9% | [+18.9%, +21.0%] | 1.0 pts | 1.4 | yes | yes |
| unique | street | 16384 | build | ordered | baseline | 6 | 8.60 ms | 8.27 ms | 0.97× [0.96, 0.97] | -3.6% | [-4.1%, -3.1%] | 0.5 pts | 0.6 | yes | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 20.6 | 18.0 | 0.87× [0.87, 0.87] | -15.0% | [-15.4%, -14.6%] | 0.4 pts | 1.4 | yes | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 253 | 241 | 0.95× [0.95, 0.96] | -4.9% | [-5.5%, -4.3%] | 0.6 pts | 1.8 | yes | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 40.6 | 36.2 | 0.89× [0.89, 0.90] | -11.9% | [-12.9%, -11.0%] | 0.9 pts | 1.7 | yes | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 685.8 µs | 589.5 µs | 0.86× [0.85, 0.87] | -16.1% | [-17.0%, -15.1%] | 0.9 pts | 1.6 | yes | yes |
| unique | u64 | 16384 | valuesFor | ordered | baseline | 6 | 28.2 | 26.0 | 0.92× [0.91, 0.93] | -8.3% | [-9.3%, -7.2%] | 1.0 pts | 3.9 | yes | yes |
| unique | u64 | 16384 | valuesBetween | ordered | baseline | 6 | 300 | 286 | 0.95× [0.95, 0.96] | -4.9% | [-5.6%, -4.3%] | 0.6 pts | 1.6 | yes | yes |
| unique | u64 | 16384 | churn | ordered | baseline | 6 | 61.0 | 55.6 | 0.91× [0.91, 0.91] | -10.1% | [-10.4%, -9.8%] | 0.3 pts | 0.6 | yes | yes |
| unique | u64 | 16384 | build | ordered | baseline | 6 | 3.68 ms | 3.43 ms | 0.93× [0.92, 0.95] | -7.0% | [-8.4%, -5.5%] | 1.4 pts | 2.1 | yes | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 12 | 39.8 | 36.7 | 0.91× [0.86, 0.97] | -9.3% | [-15.9%, -2.8%] | 6.5 pts | 3.0 | no | yes |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 12 | 308 | 288 | 0.94× [0.94, 0.94] | -6.6% | [-6.8%, -6.5%] | 0.4 pts | 0.7 | yes | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 12 | 130 | 116 | 0.92× [0.87, 0.96] | -9.0% | [-14.4%, -3.7%] | 6.0 pts | 0.8 | no | yes |
| unique | url | 4096 | valuesFor | ordered | baseline | 6 | 75.4 | 63.3 | 0.83× [0.82, 0.84] | -20.1% | [-21.7%, -18.5%] | 1.5 pts | 2.1 | yes | yes |
| unique | url | 4096 | valuesBetween | ordered | baseline | 6 | 738 | 1905 | 2.59× [2.56, 2.62] | +61.4% | [+61.0%, +61.8%] | 0.4 pts | 0.7 | yes | yes |
| unique | url | 4096 | prefix | ordered | baseline | 6 | 187 | 147 | 0.79× [0.78, 0.80] | -26.6% | [-27.8%, -25.4%] | 1.1 pts | 1.0 | yes | yes |
| unique | url | 4096 | churn | ordered | baseline | 6 | 186 | 168 | 0.90× [0.90, 0.91] | -10.6% | [-11.5%, -9.7%] | 0.8 pts | 0.8 | yes | yes |
| unique | url | 4096 | build | ordered | baseline | 6 | 3.30 ms | 2.33 ms | 0.70× [0.70, 0.71] | -41.9% | [-43.5%, -40.4%] | 1.5 pts | 1.1 | yes | yes |
| unique | url | 16384 | valuesFor | ordered | baseline | 6 | 96.1 | 87.4 | 0.90× [0.89, 0.92] | -10.6% | [-12.0%, -9.1%] | 1.4 pts | 2.0 | yes | yes |
| unique | url | 16384 | valuesBetween | ordered | baseline | 6 | 950 | 2395 | 2.54× [2.52, 2.55] | +60.6% | [+60.3%, +60.8%] | 0.3 pts | 1.0 | yes | yes |
| unique | url | 16384 | prefix | ordered | baseline | 6 | 232 | 219 | 0.94× [0.94, 0.95] | -6.0% | [-6.7%, -5.2%] | 0.7 pts | 0.8 | yes | yes |
| unique | url | 16384 | churn | ordered | baseline | 6 | 224 | 216 | 0.98× [0.96, 0.99] | -2.5% | [-4.4%, -0.6%] | 1.8 pts | 1.6 | yes | yes |
| unique | url | 16384 | build | ordered | baseline | 6 | 15.19 ms | 10.86 ms | 0.72× [0.71, 0.73] | -39.4% | [-40.9%, -37.9%] | 1.4 pts | 1.2 | yes | yes |
| unique | url | 262144 | valuesFor | ordered | baseline | 12 | 297 | 321 | 1.08× [1.07, 1.09] | +7.5% | [+6.5%, +8.4%] | 2.6 pts | 2.4 | yes | yes |
| unique | url | 262144 | valuesBetween | ordered | baseline | 12 | 2409 | 5824 | 2.43× [2.41, 2.45] | +58.8% | [+58.4%, +59.2%] | 1.0 pts | 1.8 | yes | yes |
| unique | url | 262144 | prefix | ordered | baseline | 12 | 774 | 1313 | 1.70× [1.66, 1.73] | +41.1% | [+39.9%, +42.3%] | 1.4 pts | 1.4 | yes | yes |
| unique | url | 262144 | churn | ordered | baseline | 12 | 587 | 598 | 1.02× [0.99, 1.05] | +2.1% | [-0.8%, +4.9%] | 3.3 pts | 0.7 | no | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 12 | 39.1 | 34.4 | 0.88× [0.86, 0.90] | -14.0% | [-16.7%, -11.3%] | 3.2 pts | 6.4 | no | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 12 | 442 | 1444 | 3.28× [3.26, 3.31] | +69.5% | [+69.3%, +69.8%] | 0.4 pts | 1.9 | yes | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 12 | 74.3 | 69.7 | 0.94× [0.93, 0.95] | -6.8% | [-8.0%, -5.5%] | 1.5 pts | 2.7 | yes | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 12 | 94.5 | 84.6 | 0.90× [0.89, 0.91] | -11.3% | [-12.3%, -10.3%] | 1.3 pts | 1.9 | yes | yes |
| unique | uuid | 4096 | build | ordered | baseline | 12 | 1.58 ms | 1.20 ms | 0.76× [0.75, 0.76] | -32.0% | [-32.5%, -31.5%] | 0.7 pts | 1.0 | yes | yes |
| unique | uuid | 16384 | valuesFor | ordered | baseline | 12 | 44.2 | 42.1 | 0.95× [0.93, 0.97] | -5.2% | [-7.1%, -3.3%] | 2.5 pts | 4.8 | yes | yes |
| unique | uuid | 16384 | valuesBetween | ordered | baseline | 12 | 446 | 1634 | 3.65× [3.62, 3.68] | +72.6% | [+72.3%, +72.8%] | 0.3 pts | 2.2 | yes | yes |
| unique | uuid | 16384 | prefix | ordered | baseline | 12 | 81.3 | 86.0 | 1.06× [1.04, 1.07] | +5.2% | [+4.2%, +6.3%] | 1.3 pts | 2.8 | yes | yes |
| unique | uuid | 16384 | churn | ordered | baseline | 12 | 110 | 98.7 | 0.89× [0.88, 0.90] | -12.0% | [-13.2%, -10.9%] | 1.3 pts | 1.7 | yes | yes |
| unique | uuid | 16384 | build | ordered | baseline | 12 | 7.25 ms | 5.67 ms | 0.78× [0.78, 0.79] | -27.9% | [-28.6%, -27.3%] | 1.1 pts | 1.7 | yes | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 12 | 154 | 179 | 1.16× [1.13, 1.19] | +13.8% | [+11.7%, +15.8%] | 2.2 pts | 1.8 | no | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 12 | 1064 | 3167 | 2.99× [2.93, 3.05] | +66.6% | [+65.9%, +67.2%] | 1.1 pts | 1.1 | yes | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 12 | 186 | 293 | 1.59× [1.55, 1.63] | +37.0% | [+35.4%, +38.6%] | 2.0 pts | 2.0 | yes | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 12 | 326 | 326 | 0.99× [0.97, 1.00] | -1.2% | [-2.6%, +0.2%] | 1.8 pts | 0.6 | yes | no |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique email n=4096 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=4096 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique email n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.31% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique email n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique path n=262144 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=4096 valuesFor: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=4096 prefix: ordered vs baseline: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 churn: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=16384 churn: ordered vs baseline: 1 processes resolved A as faster and 11 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique street n=4096 churn: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 valuesFor: ordered vs baseline: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=16384 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.21% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique url n=262144 valuesFor: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique url n=262144 churn: ordered vs baseline: the pooled interval [-0.78%, 4.90%] includes zero
- unique uuid n=4096 valuesFor: ordered vs baseline: the processes scatter 6.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=4096 prefix: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 valuesFor: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +0.11% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique uuid n=16384 prefix: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-2.59%, 0.21%] includes zero
