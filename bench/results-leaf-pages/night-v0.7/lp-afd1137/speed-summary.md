| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | path | 4096 | valuesFor | ordered | baseline | 6 | 109 | 124 | 1.13× [1.13, 1.14] | +11.9% | [+11.5%, +12.3%] | 0.4 pts | 0.5 | yes |
| multi | path | 4096 | valuesBetween | ordered | baseline | 6 | 4121 | 3435 | 0.83× [0.82, 0.85] | -20.0% | [-21.8%, -18.3%] | 1.6 pts | 0.9 | yes |
| multi | path | 4096 | prefix | ordered | baseline | 6 | 330 | 429 | 1.30× [1.29, 1.32] | +23.1% | [+22.2%, +24.0%] | 0.8 pts | 0.4 | yes |
| multi | path | 4096 | churn | ordered | baseline | 6 | 169 | 221 | 1.31× [1.29, 1.32] | +23.5% | [+22.7%, +24.4%] | 0.8 pts | 0.6 | yes |
| multi | path | 4096 | build | ordered | baseline | 6 | 12.15 ms | 22.23 ms | 1.83× [1.81, 1.85] | +45.4% | [+44.8%, +45.9%] | 0.5 pts | 0.6 | yes |
| multi | path | 262144 | valuesFor | ordered | baseline | 10 | 624 | 544 | 0.89× [0.88, 0.90] | -12.4% | [-14.0%, -10.8%] | 2.2 pts | 1.5 | yes |
| multi | path | 262144 | valuesBetween | ordered | baseline | 10 | 13.9 µs | 9579 | 0.68× [0.67, 0.69] | -47.0% | [-49.0%, -45.0%] | 2.8 pts | 1.7 | yes |
| multi | path | 262144 | prefix | ordered | baseline | 10 | 12.8 µs | 8793 | 0.68× [0.65, 0.70] | -48.1% | [-54.2%, -42.1%] | 8.4 pts | 0.4 | no |
| multi | path | 262144 | churn | ordered | baseline | 10 | 965 | 1000 | 1.03× [1.02, 1.04] | +3.0% | [+2.0%, +4.0%] | 1.4 pts | 0.8 | yes |
| multi | str | 4096 | valuesFor | ordered | baseline | 6 | 61.7 | 72.0 | 1.17× [1.16, 1.18] | +14.7% | [+14.1%, +15.2%] | 0.5 pts | 0.7 | yes |
| multi | str | 4096 | valuesBetween | ordered | baseline | 6 | 3358 | 2531 | 0.75× [0.74, 0.77] | -32.6% | [-35.4%, -29.7%] | 2.7 pts | 0.9 | yes |
| multi | str | 4096 | prefix | ordered | baseline | 6 | 7350 | 5095 | 0.69× [0.68, 0.70] | -44.6% | [-46.7%, -42.5%] | 2.0 pts | 0.4 | yes |
| multi | str | 4096 | churn | ordered | baseline | 6 | 79.6 | 132 | 1.65× [1.63, 1.68] | +39.5% | [+38.6%, +40.3%] | 0.8 pts | 0.9 | yes |
| multi | str | 4096 | build | ordered | baseline | 6 | 6.53 ms | 14.23 ms | 2.17× [2.12, 2.22] | +53.9% | [+52.8%, +55.0%] | 1.0 pts | 2.3 | yes |
| multi | str | 262144 | valuesFor | ordered | baseline | 10 | 365 | 294 | 0.80× [0.77, 0.83] | -25.4% | [-30.3%, -20.6%] | 6.8 pts | 3.6 | no |
| multi | str | 262144 | valuesBetween | ordered | baseline | 10 | 11.1 µs | 6630 | 0.60× [0.60, 0.61] | -65.7% | [-67.1%, -64.2%] | 2.0 pts | 1.1 | yes |
| multi | str | 262144 | prefix | ordered | baseline | 10 | 1.71 ms | 960.2 µs | 0.56× [0.56, 0.57] | -77.2% | [-79.5%, -75.0%] | 3.1 pts | 1.0 | yes |
| multi | str | 262144 | churn | ordered | baseline | 10 | 497 | 557 | 1.10× [1.08, 1.13] | +9.5% | [+7.4%, +11.6%] | 2.9 pts | 2.7 | no |
| multi | street | 4096 | valuesFor | ordered | baseline | 8 | 56.4 | 67.0 | 1.20× [1.18, 1.22] | +16.4% | [+15.0%, +17.8%] | 1.7 pts | 1.7 | yes |
| multi | street | 4096 | valuesBetween | ordered | baseline | 8 | 2251 | 1381 | 0.61× [0.61, 0.62] | -62.7% | [-64.4%, -61.1%] | 1.9 pts | 0.6 | yes |
| multi | street | 4096 | prefix | ordered | baseline | 8 | 247 | 291 | 1.18× [1.15, 1.20] | +14.9% | [+13.3%, +16.6%] | 2.0 pts | 1.3 | yes |
| multi | street | 4096 | churn | ordered | baseline | 8 | 94.4 | 145 | 1.54× [1.52, 1.56] | +35.1% | [+34.4%, +35.8%] | 0.8 pts | 0.8 | yes |
| multi | street | 4096 | build | ordered | baseline | 8 | 2.82 ms | 7.45 ms | 2.65× [2.63, 2.66] | +62.2% | [+62.0%, +62.4%] | 0.3 pts | 0.8 | yes |
| multi | u64 | 4096 | valuesFor | ordered | baseline | 8 | 37.3 | 51.1 | 1.38× [1.36, 1.39] | +27.3% | [+26.6%, +27.9%] | 0.8 pts | 0.8 | yes |
| multi | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2574 | 2097 | 0.82× [0.81, 0.83] | -22.0% | [-24.2%, -19.9%] | 2.6 pts | 0.7 | yes |
| multi | u64 | 4096 | churn | ordered | baseline | 8 | 46.3 | 80.6 | 1.74× [1.72, 1.76] | +42.7% | [+42.0%, +43.3%] | 0.8 pts | 1.0 | yes |
| multi | u64 | 4096 | build | ordered | baseline | 8 | 4.05 ms | 7.37 ms | 1.83× [1.80, 1.86] | +45.3% | [+44.5%, +46.2%] | 1.0 pts | 2.4 | yes |
| multi | u64 | 262144 | valuesFor | ordered | baseline | 10 | 168 | 140 | 0.81× [0.75, 0.89] | -22.8% | [-32.7%, -13.0%] | 13.8 pts | 8.9 | no |
| multi | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 8632 | 5788 | 0.67× [0.66, 0.68] | -49.6% | [-51.8%, -47.4%] | 3.1 pts | 1.6 | yes |
| multi | u64 | 262144 | churn | ordered | baseline | 10 | 293 | 295 | 1.02× [0.98, 1.06] | +1.7% | [-2.5%, +5.9%] | 5.9 pts | 3.8 | no |
| multi | uuid | 4096 | valuesFor | ordered | baseline | 10 | 55.6 | 71.8 | 1.29× [1.29, 1.30] | +22.7% | [+22.2%, +23.3%] | 0.7 pts | 1.1 | yes |
| multi | uuid | 4096 | valuesBetween | ordered | baseline | 10 | 3216 | 2705 | 0.84× [0.83, 0.85] | -19.2% | [-20.9%, -17.4%] | 2.4 pts | 1.1 | yes |
| multi | uuid | 4096 | prefix | ordered | baseline | 10 | 90.2 | 198 | 2.20× [2.15, 2.25] | +54.6% | [+53.6%, +55.5%] | 1.4 pts | 4.1 | yes |
| multi | uuid | 4096 | churn | ordered | baseline | 10 | 78.9 | 129 | 1.64× [1.60, 1.68] | +39.0% | [+37.3%, +40.6%] | 2.3 pts | 1.6 | yes |
| multi | uuid | 4096 | build | ordered | baseline | 10 | 6.14 ms | 14.42 ms | 2.33× [2.30, 2.36] | +57.1% | [+56.6%, +57.6%] | 0.7 pts | 1.1 | yes |
| multi | uuid | 262144 | valuesFor | ordered | baseline | 10 | 325 | 343 | 1.05× [1.04, 1.07] | +5.0% | [+3.8%, +6.1%] | 1.6 pts | 1.7 | yes |
| multi | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 11.2 µs | 7860 | 0.70× [0.69, 0.70] | -43.1% | [-44.0%, -42.2%] | 1.3 pts | 0.7 | yes |
| multi | uuid | 262144 | prefix | ordered | baseline | 10 | 742 | 697 | 0.95× [0.94, 0.96] | -5.5% | [-6.8%, -4.2%] | 1.8 pts | 1.6 | yes |
| multi | uuid | 262144 | churn | ordered | baseline | 10 | 533 | 593 | 1.12× [1.09, 1.15] | +10.8% | [+8.7%, +12.9%] | 2.9 pts | 2.0 | no |
| unique | path | 4096 | valuesFor | ordered | baseline | 6 | 85.1 | 96.2 | 1.13× [1.12, 1.14] | +11.5% | [+10.7%, +12.2%] | 0.7 pts | 0.9 | yes |
| unique | path | 4096 | valuesBetween | ordered | baseline | 6 | 2018 | 751 | 0.37× [0.37, 0.38] | -168.0% | [-172.7%, -163.3%] | 4.5 pts | 1.3 | yes |
| unique | path | 4096 | prefix | ordered | baseline | 6 | 247 | 270 | 1.10× [1.07, 1.12] | +8.9% | [+6.9%, +10.8%] | 1.8 pts | 1.1 | yes |
| unique | path | 4096 | churn | ordered | baseline | 6 | 202 | 218 | 1.08× [1.06, 1.09] | +7.2% | [+6.1%, +8.2%] | 1.0 pts | 0.5 | yes |
| unique | path | 4096 | build | ordered | baseline | 6 | 2.11 ms | 3.48 ms | 1.65× [1.62, 1.69] | +39.6% | [+38.4%, +40.8%] | 1.1 pts | 1.4 | yes |
| unique | path | 262144 | valuesFor | ordered | baseline | 10 | 529 | 447 | 0.84× [0.83, 0.86] | -18.4% | [-21.0%, -15.9%] | 3.6 pts | 2.0 | no |
| unique | path | 262144 | valuesBetween | ordered | baseline | 10 | 7317 | 3023 | 0.42× [0.42, 0.42] | -138.7% | [-140.9%, -136.4%] | 3.2 pts | 1.0 | yes |
| unique | path | 262144 | prefix | ordered | baseline | 10 | 6040 | 2159 | 0.37× [0.36, 0.38] | -171.7% | [-180.1%, -163.3%] | 11.7 pts | 0.3 | yes |
| unique | path | 262144 | churn | ordered | baseline | 10 | 906 | 869 | 0.97× [0.95, 0.98] | -3.4% | [-5.1%, -1.6%] | 2.4 pts | 1.4 | yes |
| unique | str | 4096 | valuesFor | ordered | baseline | 6 | 37.7 | 48.1 | 1.28× [1.26, 1.29] | +21.7% | [+20.8%, +22.5%] | 0.8 pts | 1.0 | yes |
| unique | str | 4096 | valuesBetween | ordered | baseline | 6 | 1518 | 473 | 0.31× [0.31, 0.31] | -222.4% | [-226.2%, -218.5%] | 3.7 pts | 1.2 | yes |
| unique | str | 4096 | prefix | ordered | baseline | 6 | 2946 | 611 | 0.21× [0.21, 0.21] | -383.5% | [-387.4%, -379.5%] | 3.8 pts | 0.8 | yes |
| unique | str | 4096 | churn | ordered | baseline | 6 | 88.5 | 117 | 1.31× [1.29, 1.34] | +23.7% | [+22.2%, +25.1%] | 1.4 pts | 1.5 | yes |
| unique | str | 4096 | build | ordered | baseline | 6 | 1.05 ms | 2.10 ms | 1.98× [1.91, 2.05] | +49.4% | [+47.5%, +51.2%] | 1.8 pts | 1.1 | yes |
| unique | str | 262144 | valuesFor | ordered | baseline | 10 | 276 | 173 | 0.62× [0.61, 0.64] | -60.5% | [-64.5%, -56.5%] | 5.5 pts | 2.6 | yes |
| unique | str | 262144 | valuesBetween | ordered | baseline | 10 | 5178 | 1474 | 0.28× [0.28, 0.29] | -253.6% | [-257.1%, -250.1%] | 4.9 pts | 0.8 | yes |
| unique | str | 262144 | prefix | ordered | baseline | 10 | 754.4 µs | 142.0 µs | 0.19× [0.18, 0.19] | -433.9% | [-449.6%, -418.3%] | 21.9 pts | 0.7 | yes |
| unique | str | 262144 | churn | ordered | baseline | 10 | 368 | 364 | 0.98× [0.96, 1.00] | -1.9% | [-3.9%, +0.1%] | 2.8 pts | 0.2 | no |
| unique | street | 4096 | valuesFor | ordered | baseline | 6 | 50.2 | 56.1 | 1.12× [1.11, 1.13] | +10.6% | [+10.1%, +11.1%] | 0.5 pts | 0.4 | yes |
| unique | street | 4096 | valuesBetween | ordered | baseline | 6 | 1731 | 592 | 0.34× [0.33, 0.35] | -192.4% | [-201.4%, -183.4%] | 8.6 pts | 2.5 | yes |
| unique | street | 4096 | prefix | ordered | baseline | 6 | 209 | 156 | 0.75× [0.74, 0.76] | -33.7% | [-35.3%, -32.0%] | 1.6 pts | 0.9 | yes |
| unique | street | 4096 | churn | ordered | baseline | 6 | 112 | 123 | 1.09× [1.07, 1.10] | +8.2% | [+6.8%, +9.5%] | 1.3 pts | 0.9 | yes |
| unique | street | 4096 | build | ordered | baseline | 6 | 1.29 ms | 2.53 ms | 1.95× [1.93, 1.97] | +48.7% | [+48.3%, +49.1%] | 0.4 pts | 0.5 | yes |
| unique | u64 | 4096 | valuesFor | ordered | baseline | 6 | 15.2 | 22.0 | 1.44× [1.40, 1.48] | +30.5% | [+28.8%, +32.3%] | 1.7 pts | 0.7 | yes |
| unique | u64 | 4096 | valuesBetween | ordered | baseline | 6 | 937 | 319 | 0.34× [0.34, 0.34] | -192.5% | [-195.1%, -190.0%] | 2.4 pts | 1.1 | yes |
| unique | u64 | 4096 | churn | ordered | baseline | 6 | 44.2 | 50.7 | 1.15× [1.13, 1.16] | +12.9% | [+11.9%, +14.0%] | 1.0 pts | 1.3 | yes |
| unique | u64 | 4096 | build | ordered | baseline | 6 | 605.7 µs | 924.9 µs | 1.57× [1.51, 1.64] | +36.3% | [+33.6%, +39.1%] | 2.6 pts | 1.5 | yes |
| unique | u64 | 262144 | valuesFor | ordered | baseline | 10 | 84.5 | 54.4 | 0.63× [0.59, 0.67] | -58.9% | [-68.2%, -49.6%] | 13.0 pts | 4.3 | no |
| unique | u64 | 262144 | valuesBetween | ordered | baseline | 10 | 2601 | 540 | 0.21× [0.20, 0.21] | -387.1% | [-399.2%, -375.1%] | 16.8 pts | 0.9 | yes |
| unique | u64 | 262144 | churn | ordered | baseline | 10 | 167 | 125 | 0.74× [0.71, 0.77] | -35.3% | [-41.2%, -29.4%] | 8.2 pts | 2.8 | no |
| unique | uuid | 4096 | valuesFor | ordered | baseline | 6 | 30.7 | 47.2 | 1.54× [1.52, 1.56] | +35.1% | [+34.3%, +35.9%] | 0.7 pts | 0.9 | yes |
| unique | uuid | 4096 | valuesBetween | ordered | baseline | 6 | 1414 | 472 | 0.33× [0.33, 0.34] | -199.1% | [-201.9%, -196.2%] | 2.7 pts | 0.8 | yes |
| unique | uuid | 4096 | prefix | ordered | baseline | 6 | 68.0 | 120 | 1.76× [1.75, 1.77] | +43.3% | [+42.9%, +43.6%] | 0.3 pts | 0.7 | yes |
| unique | uuid | 4096 | churn | ordered | baseline | 6 | 84.0 | 105 | 1.25× [1.23, 1.26] | +19.7% | [+18.7%, +20.8%] | 1.0 pts | 0.6 | yes |
| unique | uuid | 4096 | build | ordered | baseline | 6 | 982.4 µs | 1.92 ms | 1.94× [1.83, 2.07] | +48.6% | [+45.4%, +51.7%] | 3.0 pts | 1.4 | yes |
| unique | uuid | 262144 | valuesFor | ordered | baseline | 10 | 235 | 239 | 1.02× [1.01, 1.04] | +2.3% | [+1.2%, +3.5%] | 1.6 pts | 1.9 | yes |
| unique | uuid | 262144 | valuesBetween | ordered | baseline | 10 | 4748 | 1592 | 0.33× [0.33, 0.34] | -198.6% | [-200.8%, -196.5%] | 3.0 pts | 0.7 | yes |
| unique | uuid | 262144 | prefix | ordered | baseline | 10 | 419 | 403 | 0.97× [0.96, 0.98] | -3.2% | [-4.7%, -1.7%] | 2.1 pts | 2.2 | yes |
| unique | uuid | 262144 | churn | ordered | baseline | 10 | 342 | 357 | 1.03× [1.00, 1.07] | +3.0% | [-0.1%, +6.1%] | 4.4 pts | 2.0 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference; the bracket is its 95% interval across processes (rtcompare.Combine: a t interval over the per-process differences). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build.

Warnings from pooling:

- multi str n=4096 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 valuesFor: ordered vs baseline: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi str n=262144 churn: ordered vs baseline: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=4096 build: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 8.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 8 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi u64 n=262144 churn: ordered vs baseline: the pooled interval [-2.50%, 5.94%] includes zero
- multi u64 n=262144 churn: ordered vs baseline: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi u64 n=262144 churn: ordered vs baseline: 4 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi uuid n=4096 prefix: ordered vs baseline: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique str n=262144 churn: ordered vs baseline: the pooled difference of -1.91% does not clear the 2.11% median noise floor of the processes
- unique str n=262144 churn: ordered vs baseline: the pooled interval [-3.94%, 0.11%] includes zero
- unique street n=4096 valuesBetween: ordered vs baseline: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 valuesFor: ordered vs baseline: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique u64 n=262144 churn: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 prefix: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique uuid n=262144 churn: ordered vs baseline: the pooled interval [-0.10%, 6.13%] includes zero
- unique uuid n=262144 churn: ordered vs baseline: 3 processes resolved A as faster and 1 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
