| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | street | 212449 | valuesFor | ordered | baseline | 8 | 155 | 209 | 1.33× [1.28, 1.39] | +24.8% | [+21.7%, +27.9%] | 2.7 pts | 2.9 | no | yes |
| natural | street | 212449 | valuesFor | ordered | ordered-mk1 | 8 | 152 | 191 | 1.25× [1.17, 1.35] | +20.3% | [+14.5%, +26.0%] | 3.7 pts | 3.0 | no | yes |
| natural | street | 212449 | valuesFor | ordered | ordered-mk2 | 8 | 148 | 174 | 1.19× [1.17, 1.22] | +16.0% | [+14.2%, +17.7%] | 5.2 pts | 4.7 | yes | yes |
| natural | street | 212449 | valuesFor | ordered | ordered-mk4 | 8 | 144 | 163 | 1.13× [1.09, 1.18] | +11.5% | [+8.1%, +15.0%] | 2.4 pts | 2.0 | no | yes |
| natural | street | 212449 | valuesBetween | ordered | baseline | 8 | 1988 | 5964 | 3.01× [2.92, 3.11] | +66.8% | [+65.7%, +67.8%] | 0.9 pts | 1.6 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | ordered-mk1 | 8 | 1966 | 5830 | 2.91× [2.76, 3.09] | +65.7% | [+63.7%, +67.6%] | 1.6 pts | 2.5 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | ordered-mk2 | 8 | 1855 | 4225 | 2.36× [2.31, 2.42] | +57.7% | [+56.6%, +58.7%] | 1.5 pts | 2.3 | yes | yes |
| natural | street | 212449 | valuesBetween | ordered | ordered-mk4 | 8 | 1746 | 3190 | 1.80× [1.74, 1.85] | +44.3% | [+42.6%, +46.0%] | 1.7 pts | 2.9 | yes | yes |
| natural | street | 212449 | churn | ordered | baseline | 8 | 409 | 431 | 1.04× [1.00, 1.09] | +3.9% | [-0.1%, +8.0%] | 5.1 pts | 1.4 | no | no |
| natural | street | 212449 | churn | ordered | ordered-mk1 | 8 | 409 | 470 | 1.13× [1.08, 1.19] | +11.8% | [+7.3%, +16.2%] | 3.8 pts | 0.9 | no | yes |
| natural | street | 212449 | churn | ordered | ordered-mk2 | 8 | 404 | 476 | 1.18× [1.13, 1.24] | +15.3% | [+11.2%, +19.5%] | 4.8 pts | 1.3 | no | yes |
| natural | street | 212449 | churn | ordered | ordered-mk4 | 8 | 400 | 455 | 1.13× [1.09, 1.16] | +11.2% | [+8.5%, +13.8%] | 3.5 pts | 0.8 | no | yes |
| natural | u64 | 212449 | valuesFor | ordered | baseline | 8 | 121 | 119 | 0.98× [0.95, 1.01] | -2.2% | [-5.4%, +1.0%] | 2.9 pts | 3.4 | no | no |
| natural | u64 | 212449 | valuesFor | ordered | ordered-mk1 | 8 | 125 | 119 | 0.98× [0.93, 1.03] | -2.2% | [-7.6%, +3.1%] | 6.8 pts | 6.8 | no | no |
| natural | u64 | 212449 | valuesFor | ordered | ordered-mk2 | 8 | 117 | 116 | 0.96× [0.93, 0.99] | -3.9% | [-7.3%, -0.6%] | 4.2 pts | 4.2 | no | yes |
| natural | u64 | 212449 | valuesFor | ordered | ordered-mk4 | 8 | 122 | 120 | 1.01× [0.96, 1.06] | +1.0% | [-4.1%, +6.1%] | 3.6 pts | 4.3 | no | no |
| natural | u64 | 212449 | valuesBetween | ordered | baseline | 8 | 2984 | 7316 | 2.47× [2.41, 2.54] | +59.5% | [+58.4%, +60.6%] | 1.5 pts | 3.4 | yes | yes |
| natural | u64 | 212449 | valuesBetween | ordered | ordered-mk1 | 8 | 2948 | 5916 | 2.00× [1.87, 2.14] | +49.9% | [+46.5%, +53.2%] | 3.1 pts | 6.7 | yes | yes |
| natural | u64 | 212449 | valuesBetween | ordered | ordered-mk2 | 8 | 2902 | 5390 | 1.84× [1.82, 1.87] | +45.7% | [+45.0%, +46.4%] | 1.8 pts | 3.0 | yes | yes |
| natural | u64 | 212449 | valuesBetween | ordered | ordered-mk4 | 8 | 2847 | 3935 | 1.38× [1.32, 1.44] | +27.5% | [+24.5%, +30.5%] | 2.6 pts | 3.9 | no | yes |
| natural | u64 | 212449 | churn | ordered | baseline | 8 | 312 | 263 | 0.84× [0.80, 0.88] | -19.6% | [-25.2%, -14.1%] | 4.3 pts | 0.8 | no | yes |
| natural | u64 | 212449 | churn | ordered | ordered-mk1 | 8 | 310 | 309 | 0.99× [0.96, 1.01] | -1.2% | [-3.9%, +1.4%] | 2.6 pts | 0.5 | no | no |
| natural | u64 | 212449 | churn | ordered | ordered-mk2 | 8 | 313 | 322 | 1.03× [1.00, 1.07] | +3.4% | [-0.2%, +6.9%] | 3.9 pts | 0.7 | no | no |
| natural | u64 | 212449 | churn | ordered | ordered-mk4 | 8 | 315 | 321 | 0.98× [0.93, 1.05] | -1.8% | [-8.1%, +4.4%] | 4.7 pts | 1.0 | no | no |
| natural | url | 212449 | valuesFor | ordered | baseline | 8 | 302 | 382 | 1.26× [1.22, 1.29] | +20.5% | [+18.4%, +22.6%] | 1.9 pts | 2.8 | no | yes |
| natural | url | 212449 | valuesFor | ordered | ordered-mk1 | 8 | 303 | 323 | 1.08× [1.06, 1.10] | +7.2% | [+5.5%, +8.9%] | 2.7 pts | 2.5 | yes | yes |
| natural | url | 212449 | valuesFor | ordered | ordered-mk2 | 8 | 299 | 311 | 1.05× [1.02, 1.08] | +4.8% | [+2.3%, +7.2%] | 2.4 pts | 2.4 | no | yes |
| natural | url | 212449 | valuesFor | ordered | ordered-mk4 | 8 | 297 | 300 | 1.02× [1.00, 1.05] | +2.1% | [-0.2%, +4.4%] | 1.6 pts | 2.0 | no | no |
| natural | url | 212449 | valuesBetween | ordered | baseline | 8 | 5810 | 10.8 µs | 1.86× [1.81, 1.91] | +46.2% | [+44.7%, +47.6%] | 1.2 pts | 2.3 | yes | yes |
| natural | url | 212449 | valuesBetween | ordered | ordered-mk1 | 8 | 5708 | 9928 | 1.73× [1.72, 1.74] | +42.2% | [+41.9%, +42.4%] | 0.5 pts | 1.3 | yes | yes |
| natural | url | 212449 | valuesBetween | ordered | ordered-mk2 | 8 | 5617 | 7911 | 1.40× [1.39, 1.42] | +28.8% | [+27.9%, +29.7%] | 1.6 pts | 2.8 | yes | yes |
| natural | url | 212449 | valuesBetween | ordered | ordered-mk4 | 8 | 5628 | 6510 | 1.15× [1.13, 1.16] | +12.8% | [+11.9%, +13.7%] | 2.2 pts | 3.9 | yes | yes |
| natural | url | 212449 | churn | ordered | baseline | 8 | 648 | 633 | 0.97× [0.95, 1.00] | -2.9% | [-5.6%, -0.2%] | 4.0 pts | 0.9 | no | no |
| natural | url | 212449 | churn | ordered | ordered-mk1 | 8 | 645 | 650 | 0.99× [0.95, 1.04] | -0.8% | [-5.4%, +3.7%] | 4.3 pts | 0.8 | no | no |
| natural | url | 212449 | churn | ordered | ordered-mk2 | 8 | 638 | 660 | 1.01× [0.97, 1.06] | +1.1% | [-3.0%, +5.3%] | 3.8 pts | 0.7 | no | no |
| natural | url | 212449 | churn | ordered | ordered-mk4 | 8 | 633 | 651 | 1.02× [0.96, 1.08] | +1.5% | [-4.7%, +7.8%] | 3.6 pts | 0.7 | no | no |
| single-value | street | 212449 | valuesFor | ordered | baseline | 8 | 115 | 172 | 1.48× [1.39, 1.58] | +32.3% | [+28.1%, +36.6%] | 3.9 pts | 4.0 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | ordered-mk1 | 8 | 115 | 158 | 1.40× [1.23, 1.61] | +28.5% | [+18.8%, +38.1%] | 7.1 pts | 6.2 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | ordered-mk2 | 8 | 111 | 141 | 1.26× [1.20, 1.33] | +20.7% | [+16.8%, +24.6%] | 2.8 pts | 3.0 | no | yes |
| single-value | street | 212449 | valuesFor | ordered | ordered-mk4 | 8 | 110 | 136 | 1.24× [1.18, 1.30] | +19.4% | [+15.5%, +23.3%] | 3.0 pts | 3.3 | no | yes |
| single-value | street | 212449 | valuesBetween | ordered | baseline | 8 | 1020 | 3874 | 3.70× [3.59, 3.81] | +73.0% | [+72.1%, +73.8%] | 0.9 pts | 1.5 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | ordered-mk1 | 8 | 946 | 3805 | 4.02× [3.72, 4.37] | +75.1% | [+73.1%, +77.1%] | 1.5 pts | 3.1 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | ordered-mk2 | 8 | 886 | 2781 | 3.21× [3.12, 3.31] | +68.9% | [+68.0%, +69.8%] | 0.8 pts | 2.1 | yes | yes |
| single-value | street | 212449 | valuesBetween | ordered | ordered-mk4 | 8 | 852 | 2095 | 2.45× [2.39, 2.52] | +59.2% | [+58.2%, +60.2%] | 0.8 pts | 1.7 | yes | yes |
| single-value | street | 212449 | churn | ordered | baseline | 8 | 318 | 405 | 1.28× [1.23, 1.33] | +21.7% | [+18.8%, +24.6%] | 2.7 pts | 1.6 | no | yes |
| single-value | street | 212449 | churn | ordered | ordered-mk1 | 8 | 300 | 404 | 1.34× [1.28, 1.39] | +25.1% | [+22.1%, +28.1%] | 2.1 pts | 0.9 | no | yes |
| single-value | street | 212449 | churn | ordered | ordered-mk2 | 8 | 321 | 471 | 1.47× [1.41, 1.53] | +31.9% | [+29.3%, +34.6%] | 2.2 pts | 1.2 | yes | yes |
| single-value | street | 212449 | churn | ordered | ordered-mk4 | 8 | 322 | 460 | 1.39× [1.37, 1.42] | +28.2% | [+27.1%, +29.4%] | 1.4 pts | 0.9 | yes | yes |
| single-value | u64 | 212449 | valuesFor | ordered | baseline | 8 | 45.4 | 46.0 | 1.02× [0.93, 1.14] | +2.0% | [-8.0%, +12.0%] | 7.6 pts | 7.0 | no | no |
| single-value | u64 | 212449 | valuesFor | ordered | ordered-mk1 | 8 | 45.3 | 37.0 | 0.85× [0.75, 0.99] | -17.4% | [-34.0%, -0.8%] | 11.8 pts | 6.3 | no | yes |
| single-value | u64 | 212449 | valuesFor | ordered | ordered-mk2 | 8 | 43.7 | 43.5 | 0.95× [0.86, 1.07] | -4.9% | [-16.8%, +7.0%] | 8.1 pts | 6.2 | no | no |
| single-value | u64 | 212449 | valuesFor | ordered | ordered-mk4 | 8 | 42.6 | 44.8 | 1.05× [1.01, 1.10] | +5.1% | [+1.0%, +9.1%] | 3.2 pts | 4.0 | no | yes |
| single-value | u64 | 212449 | valuesBetween | ordered | baseline | 8 | 721 | 1773 | 2.42× [2.31, 2.53] | +58.6% | [+56.8%, +60.5%] | 1.3 pts | 1.7 | yes | yes |
| single-value | u64 | 212449 | valuesBetween | ordered | ordered-mk1 | 8 | 678 | 1934 | 2.85× [2.80, 2.91] | +65.0% | [+64.3%, +65.6%] | 0.5 pts | 2.1 | yes | yes |
| single-value | u64 | 212449 | valuesBetween | ordered | ordered-mk2 | 8 | 682 | 1849 | 2.78× [2.73, 2.83] | +64.0% | [+63.4%, +64.6%] | 1.0 pts | 3.4 | yes | yes |
| single-value | u64 | 212449 | valuesBetween | ordered | ordered-mk4 | 8 | 671 | 1259 | 1.89× [1.87, 1.91] | +47.0% | [+46.4%, +47.5%] | 0.9 pts | 2.3 | yes | yes |
| single-value | u64 | 212449 | churn | ordered | baseline | 8 | 165 | 188 | 1.10× [1.05, 1.16] | +9.4% | [+4.9%, +13.8%] | 3.3 pts | 1.3 | no | yes |
| single-value | u64 | 212449 | churn | ordered | ordered-mk1 | 8 | 155 | 172 | 1.07× [1.04, 1.10] | +6.5% | [+3.6%, +9.4%] | 2.5 pts | 0.7 | no | yes |
| single-value | u64 | 212449 | churn | ordered | ordered-mk2 | 8 | 170 | 234 | 1.34× [1.31, 1.38] | +25.6% | [+23.5%, +27.7%] | 2.1 pts | 1.1 | yes | yes |
| single-value | u64 | 212449 | churn | ordered | ordered-mk4 | 8 | 165 | 221 | 1.33× [1.31, 1.34] | +24.7% | [+23.9%, +25.5%] | 1.1 pts | 0.5 | yes | yes |
| single-value | url | 212449 | valuesFor | ordered | baseline | 8 | 230 | 322 | 1.40× [1.38, 1.42] | +28.5% | [+27.5%, +29.5%] | 2.6 pts | 3.1 | yes | yes |
| single-value | url | 212449 | valuesFor | ordered | ordered-mk1 | 8 | 225 | 257 | 1.16× [1.07, 1.27] | +14.0% | [+6.8%, +21.1%] | 4.4 pts | 3.8 | no | yes |
| single-value | url | 212449 | valuesFor | ordered | ordered-mk2 | 8 | 216 | 243 | 1.12× [1.07, 1.18] | +10.8% | [+6.1%, +15.5%] | 3.9 pts | 3.0 | no | yes |
| single-value | url | 212449 | valuesFor | ordered | ordered-mk4 | 8 | 230 | 248 | 1.08× [1.04, 1.13] | +7.3% | [+3.5%, +11.2%] | 3.8 pts | 3.4 | no | yes |
| single-value | url | 212449 | valuesBetween | ordered | baseline | 8 | 2775 | 5543 | 2.01× [1.94, 2.10] | +50.3% | [+48.4%, +52.3%] | 2.1 pts | 2.9 | yes | yes |
| single-value | url | 212449 | valuesBetween | ordered | ordered-mk1 | 8 | 2676 | 5879 | 2.22× [2.09, 2.36] | +55.0% | [+52.3%, +57.6%] | 2.1 pts | 3.5 | yes | yes |
| single-value | url | 212449 | valuesBetween | ordered | ordered-mk2 | 8 | 2555 | 4592 | 1.84× [1.77, 1.91] | +45.6% | [+43.6%, +47.7%] | 1.6 pts | 2.4 | yes | yes |
| single-value | url | 212449 | valuesBetween | ordered | ordered-mk4 | 8 | 2319 | 3246 | 1.40× [1.34, 1.47] | +28.6% | [+25.4%, +31.7%] | 2.9 pts | 4.1 | no | yes |
| single-value | url | 212449 | churn | ordered | baseline | 8 | 537 | 574 | 1.08× [1.04, 1.11] | +7.0% | [+3.8%, +10.2%] | 2.1 pts | 1.6 | no | yes |
| single-value | url | 212449 | churn | ordered | ordered-mk1 | 8 | 546 | 555 | 1.02× [1.00, 1.04] | +1.9% | [+0.3%, +3.4%] | 1.6 pts | 0.8 | yes | yes |
| single-value | url | 212449 | churn | ordered | ordered-mk2 | 8 | 538 | 637 | 1.19× [1.18, 1.21] | +16.3% | [+15.1%, +17.4%] | 1.3 pts | 1.0 | yes | yes |
| single-value | url | 212449 | churn | ordered | ordered-mk4 | 8 | 536 | 619 | 1.17× [1.13, 1.20] | +14.2% | [+11.7%, +16.6%] | 1.7 pts | 1.2 | no | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural street n=212449 valuesFor: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs ordered-mk1: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesFor: ordered vs ordered-mk2: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs ordered-mk1: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs ordered-mk2: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 valuesBetween: ordered vs ordered-mk4: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=212449 churn: ordered vs baseline: the pooled interval [-0.13%, 7.98%] includes zero
- natural u64 n=212449 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of -0.16% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=212449 valuesFor: ordered vs baseline: the pooled interval [-5.41%, 0.98%] includes zero
- natural u64 n=212449 valuesFor: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=212449 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=212449 valuesFor: ordered vs ordered-mk1: the pooled interval [-7.57%, 3.11%] includes zero
- natural u64 n=212449 valuesFor: ordered vs ordered-mk1: the processes scatter 6.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=212449 valuesFor: ordered vs ordered-mk1: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=212449 valuesFor: ordered vs ordered-mk2: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=212449 valuesFor: ordered vs ordered-mk2: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=212449 valuesFor: ordered vs ordered-mk4: the pooled interval [-4.12%, 6.08%] includes zero
- natural u64 n=212449 valuesFor: ordered vs ordered-mk4: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=212449 valuesFor: ordered vs ordered-mk4: 3 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural u64 n=212449 valuesBetween: ordered vs baseline: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=212449 valuesBetween: ordered vs ordered-mk1: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=212449 valuesBetween: ordered vs ordered-mk2: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=212449 valuesBetween: ordered vs ordered-mk4: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=212449 churn: ordered vs ordered-mk1: the pooled difference of -1.25% does not clear the 1.99% noise floor, the bound on what the harness reports between identical code in every process
- natural u64 n=212449 churn: ordered vs ordered-mk1: the pooled interval [-3.92%, 1.42%] includes zero
- natural u64 n=212449 churn: ordered vs ordered-mk2: the pooled interval [-0.17%, 6.89%] includes zero
- natural u64 n=212449 churn: ordered vs ordered-mk4: the pooled interval [-8.11%, 4.45%] includes zero
- natural url n=212449 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=212449 valuesFor: ordered vs ordered-mk1: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=212449 valuesFor: ordered vs ordered-mk2: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=212449 valuesFor: ordered vs ordered-mk4: the pooled interval [-0.20%, 4.36%] includes zero
- natural url n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=212449 valuesBetween: ordered vs ordered-mk2: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=212449 valuesBetween: ordered vs ordered-mk4: the processes scatter 3.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=212449 churn: ordered vs baseline: the pooled difference of -2.94% does not clear the 3.23% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=212449 churn: ordered vs ordered-mk1: the pooled difference of -0.83% does not clear the 2.31% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=212449 churn: ordered vs ordered-mk1: the pooled interval [-5.40%, 3.74%] includes zero
- natural url n=212449 churn: ordered vs ordered-mk2: the pooled difference of 1.14% does not clear the 1.48% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=212449 churn: ordered vs ordered-mk2: the pooled interval [-3.01%, 5.29%] includes zero
- natural url n=212449 churn: ordered vs ordered-mk4: the pooled difference of 1.55% does not clear the 3.00% noise floor, the bound on what the harness reports between identical code in every process
- natural url n=212449 churn: ordered vs ordered-mk4: the pooled interval [-4.67%, 7.76%] includes zero
- single-value street n=212449 valuesFor: ordered vs baseline: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs ordered-mk1: the processes scatter 6.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs ordered-mk2: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesFor: ordered vs ordered-mk4: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs ordered-mk1: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value street n=212449 valuesBetween: ordered vs ordered-mk2: the A/A validations found a systematic difference of +0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value street n=212449 valuesBetween: ordered vs ordered-mk2: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=212449 valuesFor: ordered vs baseline: the pooled interval [-7.99%, 11.97%] includes zero
- single-value u64 n=212449 valuesFor: ordered vs baseline: the processes scatter 7.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=212449 valuesFor: ordered vs baseline: 3 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=212449 valuesFor: ordered vs ordered-mk1: the processes scatter 6.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=212449 valuesFor: ordered vs ordered-mk1: 1 processes resolved A as faster and 6 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=212449 valuesFor: ordered vs ordered-mk2: the pooled interval [-16.81%, 6.95%] includes zero
- single-value u64 n=212449 valuesFor: ordered vs ordered-mk2: the processes scatter 6.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=212449 valuesFor: ordered vs ordered-mk2: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- single-value u64 n=212449 valuesFor: ordered vs ordered-mk4: the processes scatter 4.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=212449 valuesBetween: ordered vs ordered-mk1: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=212449 valuesBetween: ordered vs ordered-mk2: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value u64 n=212449 valuesBetween: ordered vs ordered-mk4: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 valuesFor: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 valuesFor: ordered vs ordered-mk1: the A/A validations found a systematic difference of +0.67% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- single-value url n=212449 valuesFor: ordered vs ordered-mk1: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 valuesFor: ordered vs ordered-mk2: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 valuesFor: ordered vs ordered-mk4: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 valuesBetween: ordered vs baseline: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 valuesBetween: ordered vs ordered-mk1: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 valuesBetween: ordered vs ordered-mk2: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 valuesBetween: ordered vs ordered-mk4: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- single-value url n=212449 churn: ordered vs baseline: the A/A validations found a systematic difference of -0.95% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
