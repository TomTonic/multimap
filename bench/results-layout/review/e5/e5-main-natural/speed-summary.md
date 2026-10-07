| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 96.8 | 80.5 | 0.83× [0.82, 0.84] | -20.8% | [-22.0%, -19.6%] | 1.7 pts | 1.3 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 98.8 | 161 | 1.65× [1.64, 1.67] | +39.4% | [+38.9%, +40.0%] | 0.6 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2632 | 2606 | 1.00× [0.98, 1.03] | +0.2% | [-2.3%, +2.6%] | 2.3 pts | 0.7 | no | no |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2628 | 5590 | 2.12× [2.10, 2.15] | +52.9% | [+52.4%, +53.4%] | 0.7 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 200 | 153 | 0.76× [0.75, 0.77] | -31.8% | [-34.2%, -29.4%] | 2.2 pts | 1.6 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 200 | 236 | 1.20× [1.18, 1.22] | +16.7% | [+15.2%, +18.1%] | 1.5 pts | 1.5 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 9.91 ms | 6.67 ms | 0.67× [0.66, 0.68] | -48.8% | [-51.4%, -46.2%] | 4.1 pts | 1.9 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 9.84 ms | 11.37 ms | 1.15× [1.14, 1.16] | +13.3% | [+12.4%, +14.1%] | 0.9 pts | 0.8 | yes | yes |
| natural | dirs | 65536 | valuesFor | ordered | baseline | 8 | 163 | 166 | 1.04× [0.98, 1.10] | +3.7% | [-1.8%, +9.3%] | 5.4 pts | 4.9 | no | no |
| natural | dirs | 65536 | valuesFor | ordered | btree-sets | 8 | 173 | 361 | 2.10× [2.05, 2.14] | +52.3% | [+51.2%, +53.3%] | 1.1 pts | 1.6 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | baseline | 8 | 3327 | 3967 | 1.20× [1.18, 1.22] | +16.7% | [+15.4%, +17.9%] | 1.2 pts | 1.6 | yes | yes |
| natural | dirs | 65536 | valuesBetween | ordered | btree-sets | 8 | 3628 | 8441 | 2.38× [2.31, 2.47] | +58.1% | [+56.7%, +59.4%] | 1.6 pts | 2.7 | yes | yes |
| natural | dirs | 65536 | churn | ordered | baseline | 8 | 331 | 355 | 1.07× [1.05, 1.08] | +6.4% | [+5.1%, +7.8%] | 1.6 pts | 1.4 | yes | yes |
| natural | dirs | 65536 | churn | ordered | btree-sets | 8 | 391 | 576 | 1.48× [1.47, 1.49] | +32.4% | [+31.9%, +32.9%] | 0.6 pts | 1.1 | yes | yes |
| natural | dirs | 65536 | build | ordered | baseline | 8 | 254.12 ms | 235.96 ms | 0.92× [0.91, 0.93] | -8.5% | [-9.5%, -7.6%] | 1.6 pts | 2.1 | yes | yes |
| natural | dirs | 65536 | build | ordered | btree-sets | 8 | 255.06 ms | 391.42 ms | 1.53× [1.52, 1.54] | +34.7% | [+34.3%, +35.1%] | 0.4 pts | 0.7 | yes | yes |
| natural | street | 4096 | valuesFor | ordered | baseline | 8 | 65.4 | 52.2 | 0.79× [0.77, 0.81] | -27.3% | [-30.7%, -23.9%] | 3.4 pts | 3.0 | no | yes |
| natural | street | 4096 | valuesFor | ordered | btree-sets | 8 | 66.0 | 136 | 2.04× [1.95, 2.14] | +51.0% | [+48.7%, +53.2%] | 2.1 pts | 4.5 | yes | yes |
| natural | street | 4096 | valuesBetween | ordered | baseline | 8 | 1908 | 2169 | 1.15× [1.12, 1.18] | +12.8% | [+10.6%, +15.0%] | 2.2 pts | 1.1 | no | yes |
| natural | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 1890 | 4721 | 2.50× [2.41, 2.58] | +59.9% | [+58.5%, +61.3%] | 1.3 pts | 2.8 | yes | yes |
| natural | street | 4096 | churn | ordered | baseline | 8 | 149 | 91.1 | 0.61× [0.60, 0.62] | -63.7% | [-65.7%, -61.6%] | 2.7 pts | 1.7 | yes | yes |
| natural | street | 4096 | churn | ordered | btree-sets | 8 | 151 | 192 | 1.28× [1.27, 1.29] | +21.9% | [+21.0%, +22.7%] | 0.9 pts | 1.2 | yes | yes |
| natural | street | 4096 | build | ordered | baseline | 8 | 6.09 ms | 3.50 ms | 0.57× [0.57, 0.58] | -74.5% | [-76.2%, -72.9%] | 1.9 pts | 1.0 | yes | yes |
| natural | street | 4096 | build | ordered | btree-sets | 8 | 6.12 ms | 7.82 ms | 1.26× [1.24, 1.29] | +20.9% | [+19.6%, +22.3%] | 1.4 pts | 1.0 | yes | yes |
| natural | street | 65536 | valuesFor | ordered | baseline | 8 | 106 | 103 | 0.98× [0.94, 1.03] | -2.1% | [-6.8%, +2.6%] | 5.8 pts | 4.2 | no | no |
| natural | street | 65536 | valuesFor | ordered | btree-sets | 8 | 112 | 297 | 2.67× [2.61, 2.74] | +62.5% | [+61.6%, +63.5%] | 1.4 pts | 2.7 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | baseline | 8 | 2502 | 3159 | 1.26× [1.25, 1.27] | +20.8% | [+20.0%, +21.6%] | 1.7 pts | 2.4 | yes | yes |
| natural | street | 65536 | valuesBetween | ordered | btree-sets | 8 | 2640 | 7189 | 2.68× [2.65, 2.71] | +62.7% | [+62.3%, +63.2%] | 1.7 pts | 4.4 | yes | yes |
| natural | street | 65536 | churn | ordered | baseline | 8 | 230 | 209 | 0.90× [0.89, 0.92] | -10.7% | [-12.8%, -8.6%] | 2.2 pts | 2.2 | no | yes |
| natural | street | 65536 | churn | ordered | btree-sets | 8 | 249 | 430 | 1.70× [1.65, 1.75] | +41.2% | [+39.4%, +42.9%] | 1.6 pts | 4.9 | yes | yes |
| natural | street | 65536 | build | ordered | baseline | 8 | 151.99 ms | 118.15 ms | 0.78× [0.77, 0.79] | -28.3% | [-29.7%, -26.8%] | 1.8 pts | 1.4 | yes | yes |
| natural | street | 65536 | build | ordered | btree-sets | 8 | 153.00 ms | 246.59 ms | 1.60× [1.58, 1.63] | +37.6% | [+36.6%, +38.5%] | 0.9 pts | 1.7 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | baseline | 8 | 45.3 | 36.6 | 0.81× [0.80, 0.82] | -23.5% | [-24.6%, -22.3%] | 1.4 pts | 1.1 | yes | yes |
| natural | u64 | 4096 | valuesFor | ordered | btree-sets | 8 | 45.3 | 157 | 3.48× [3.44, 3.52] | +71.2% | [+70.9%, +71.6%] | 0.3 pts | 1.8 | yes | yes |
| natural | u64 | 4096 | valuesBetween | ordered | baseline | 8 | 2799 | 2572 | 0.92× [0.90, 0.94] | -8.8% | [-11.1%, -6.5%] | 2.2 pts | 0.7 | no | yes |
| natural | u64 | 4096 | valuesBetween | ordered | btree-sets | 8 | 2782 | 7168 | 2.61× [2.57, 2.66] | +61.7% | [+61.0%, +62.4%] | 0.7 pts | 1.3 | yes | yes |
| natural | u64 | 4096 | churn | ordered | baseline | 8 | 72.5 | 48.3 | 0.66× [0.66, 0.67] | -51.0% | [-52.3%, -49.7%] | 1.6 pts | 1.0 | yes | yes |
| natural | u64 | 4096 | churn | ordered | btree-sets | 8 | 72.7 | 159 | 2.18× [2.17, 2.20] | +54.2% | [+53.8%, +54.6%] | 0.4 pts | 1.0 | yes | yes |
| natural | u64 | 4096 | build | ordered | baseline | 8 | 9.29 ms | 4.84 ms | 0.52× [0.52, 0.52] | -92.7% | [-93.4%, -92.0%] | 1.2 pts | 2.3 | yes | yes |
| natural | u64 | 4096 | build | ordered | btree-sets | 8 | 9.30 ms | 15.15 ms | 1.63× [1.62, 1.65] | +38.8% | [+38.1%, +39.4%] | 0.6 pts | 2.0 | yes | yes |
| natural | u64 | 65536 | valuesFor | ordered | baseline | 8 | 63.0 | 52.2 | 0.84× [0.82, 0.87] | -18.5% | [-22.1%, -15.0%] | 3.5 pts | 3.5 | no | yes |
| natural | u64 | 65536 | valuesFor | ordered | btree-sets | 8 | 81.0 | 312 | 3.73× [3.63, 3.85] | +73.2% | [+72.4%, +74.0%] | 0.8 pts | 1.7 | yes | yes |
| natural | u64 | 65536 | valuesBetween | ordered | baseline | 8 | 4178 | 4135 | 0.99× [0.99, 1.00] | -0.7% | [-1.5%, +0.1%] | 0.8 pts | 1.4 | yes | no |
| natural | u64 | 65536 | valuesBetween | ordered | btree-sets | 8 | 4570 | 10.3 µs | 2.29× [2.22, 2.37] | +56.3% | [+54.9%, +57.7%] | 1.4 pts | 3.4 | yes | yes |
| natural | u64 | 65536 | churn | ordered | baseline | 8 | 157 | 124 | 0.77× [0.76, 0.79] | -29.8% | [-32.4%, -27.2%] | 2.8 pts | 1.9 | yes | yes |
| natural | u64 | 65536 | churn | ordered | btree-sets | 8 | 199 | 407 | 2.04× [2.01, 2.07] | +51.0% | [+50.3%, +51.6%] | 1.0 pts | 2.1 | yes | yes |
| natural | u64 | 65536 | build | ordered | baseline | 8 | 197.77 ms | 140.38 ms | 0.71× [0.70, 0.72] | -40.9% | [-43.1%, -38.8%] | 2.0 pts | 1.3 | yes | yes |
| natural | u64 | 65536 | build | ordered | btree-sets | 8 | 201.51 ms | 490.48 ms | 2.43× [2.39, 2.46] | +58.8% | [+58.2%, +59.4%] | 0.6 pts | 1.4 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | baseline | 8 | 97.5 | 85.2 | 0.87× [0.87, 0.88] | -14.4% | [-15.6%, -13.3%] | 1.1 pts | 1.1 | yes | yes |
| natural | url | 4096 | valuesFor | ordered | btree-sets | 8 | 98.7 | 188 | 1.90× [1.88, 1.92] | +47.4% | [+46.9%, +47.8%] | 0.5 pts | 1.0 | yes | yes |
| natural | url | 4096 | valuesBetween | ordered | baseline | 8 | 3894 | 3731 | 0.97× [0.95, 0.99] | -3.3% | [-5.7%, -0.8%] | 2.6 pts | 1.1 | no | yes |
| natural | url | 4096 | valuesBetween | ordered | btree-sets | 8 | 3882 | 7704 | 1.99× [1.98, 2.00] | +49.8% | [+49.4%, +50.1%] | 0.5 pts | 0.9 | yes | yes |
| natural | url | 4096 | churn | ordered | baseline | 8 | 179 | 128 | 0.71× [0.71, 0.72] | -40.1% | [-41.3%, -38.9%] | 1.3 pts | 1.2 | yes | yes |
| natural | url | 4096 | churn | ordered | btree-sets | 8 | 179 | 218 | 1.23× [1.22, 1.24] | +18.6% | [+17.7%, +19.4%] | 0.9 pts | 1.0 | yes | yes |
| natural | url | 4096 | build | ordered | baseline | 8 | 17.56 ms | 11.53 ms | 0.66× [0.65, 0.67] | -52.1% | [-54.1%, -50.2%] | 1.9 pts | 2.2 | yes | yes |
| natural | url | 4096 | build | ordered | btree-sets | 8 | 17.62 ms | 20.47 ms | 1.16× [1.15, 1.17] | +13.6% | [+12.8%, +14.4%] | 0.9 pts | 1.4 | yes | yes |
| natural | url | 65536 | valuesFor | ordered | baseline | 8 | 186 | 216 | 1.16× [1.14, 1.19] | +14.0% | [+12.0%, +16.1%] | 1.9 pts | 1.7 | no | yes |
| natural | url | 65536 | valuesFor | ordered | btree-sets | 8 | 203 | 424 | 2.09× [2.06, 2.12] | +52.1% | [+51.4%, +52.8%] | 0.7 pts | 0.9 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | baseline | 8 | 5178 | 5613 | 1.08× [1.07, 1.10] | +7.6% | [+6.4%, +8.8%] | 1.4 pts | 1.8 | yes | yes |
| natural | url | 65536 | valuesBetween | ordered | btree-sets | 8 | 6240 | 13.5 µs | 2.20× [2.16, 2.24] | +54.5% | [+53.6%, +55.3%] | 0.8 pts | 1.7 | yes | yes |
| natural | url | 65536 | churn | ordered | baseline | 8 | 380 | 383 | 1.01× [1.00, 1.03] | +1.1% | [-0.4%, +2.7%] | 1.7 pts | 1.1 | yes | no |
| natural | url | 65536 | churn | ordered | btree-sets | 8 | 442 | 591 | 1.33× [1.30, 1.35] | +24.5% | [+23.0%, +26.1%] | 1.8 pts | 4.1 | yes | yes |
| natural | url | 65536 | build | ordered | baseline | 8 | 481.36 ms | 433.66 ms | 0.89× [0.88, 0.90] | -12.5% | [-13.7%, -11.3%] | 1.4 pts | 1.5 | yes | yes |
| natural | url | 65536 | build | ordered | btree-sets | 8 | 479.49 ms | 718.02 ms | 1.50× [1.48, 1.51] | +33.1% | [+32.4%, +33.9%] | 1.0 pts | 2.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 valuesBetween: ordered vs baseline: the pooled difference of 0.16% does not clear the 1.31% noise floor, the bound on what the harness reports between identical code in every process
- natural dirs n=4096 valuesBetween: ordered vs baseline: the pooled interval [-2.27%, 2.59%] includes zero
- natural dirs n=65536 valuesFor: ordered vs baseline: the pooled interval [-1.82%, 9.29%] includes zero
- natural dirs n=65536 valuesFor: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=65536 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesFor: ordered vs baseline: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 4.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: the pooled interval [-6.82%, 2.64%] includes zero
- natural street n=65536 valuesFor: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesFor: ordered vs baseline: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural street n=65536 valuesFor: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural street n=65536 churn: ordered vs btree-sets: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=4096 build: ordered vs baseline: the A/A validations found a systematic difference of +0.04% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=4096 build: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesFor: ordered vs baseline: the A/A validations found a systematic difference of +0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 valuesFor: ordered vs baseline: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 valuesBetween: ordered vs baseline: the pooled interval [-1.51%, 0.05%] includes zero
- natural u64 n=65536 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.49% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural u64 n=65536 valuesBetween: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural u64 n=65536 churn: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=4096 build: ordered vs baseline: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 churn: ordered vs baseline: the pooled interval [-0.41%, 2.66%] includes zero
- natural url n=65536 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.38% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural url n=65536 churn: ordered vs btree-sets: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural url n=65536 build: ordered vs btree-sets: the processes scatter 2.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
