| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 76.0 | 177 | 2.33× [2.31, 2.35] | +57.1% | [+56.7%, +57.5%] | 0.5 pts | 2.2 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage | 8 | 75.2 | 75.5 | 1.01× [1.00, 1.01] | +0.6% | [+0.1%, +1.0%] | 0.5 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mv | 8 | 76.4 | 150 | 1.96× [1.95, 1.98] | +49.1% | [+48.8%, +49.4%] | 0.4 pts | 1.0 | yes | yes |
| multi-str | street | 4096 | valuesFor | ordered | ordered-lpage-mvzc | 8 | 76.0 | 130 | 1.70× [1.69, 1.72] | +41.3% | [+40.9%, +41.7%] | 0.5 pts | 1.1 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 3296 | 9010 | 2.73× [2.70, 2.75] | +63.3% | [+63.0%, +63.7%] | 0.4 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage | 8 | 3307 | 3313 | 1.00× [1.00, 1.01] | +0.0% | [-0.4%, +0.5%] | 0.6 pts | 0.4 | yes | no |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mv | 8 | 3280 | 2521 | 0.78× [0.76, 0.79] | -29.0% | [-31.0%, -27.1%] | 2.3 pts | 1.4 | yes | yes |
| multi-str | street | 4096 | valuesBetween | ordered | ordered-lpage-mvzc | 8 | 3259 | 2040 | 0.62× [0.62, 0.63] | -60.0% | [-61.4%, -58.7%] | 1.6 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | btree-sets | 8 | 320 | 924 | 2.89× [2.88, 2.91] | +65.4% | [+65.2%, +65.6%] | 0.2 pts | 0.6 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage | 8 | 315 | 316 | 1.02× [1.00, 1.03] | +1.6% | [+0.4%, +2.7%] | 1.4 pts | 0.7 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mv | 8 | 319 | 347 | 1.10× [1.08, 1.11] | +8.9% | [+7.7%, +10.1%] | 1.4 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | prefix | ordered | ordered-lpage-mvzc | 8 | 317 | 293 | 0.93× [0.92, 0.94] | -7.8% | [-8.9%, -6.6%] | 1.4 pts | 0.8 | yes | yes |
| multi-str | street | 4096 | churn | ordered | btree-sets | 8 | 137 | 248 | 1.81× [1.79, 1.84] | +44.9% | [+44.3%, +45.5%] | 0.7 pts | 1.9 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage | 8 | 136 | 136 | 1.00× [1.00, 1.01] | +0.5% | [+0.3%, +0.7%] | 0.2 pts | 0.9 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mv | 8 | 137 | 290 | 2.12× [2.11, 2.12] | +52.8% | [+52.7%, +52.9%] | 0.1 pts | 0.4 | yes | yes |
| multi-str | street | 4096 | churn | ordered | ordered-lpage-mvzc | 8 | 140 | 370 | 2.64× [2.61, 2.66] | +62.1% | [+61.7%, +62.4%] | 0.4 pts | 1.5 | yes | yes |
| multi-str | street | 4096 | build | ordered | btree-sets | 8 | 5.50 ms | 10.12 ms | 1.85× [1.83, 1.86] | +45.8% | [+45.4%, +46.2%] | 0.5 pts | 1.2 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage | 8 | 5.48 ms | 5.50 ms | 1.01× [1.00, 1.01] | +0.5% | [-0.3%, +1.3%] | 1.0 pts | 1.4 | yes | no |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mv | 8 | 5.51 ms | 12.66 ms | 2.30× [2.29, 2.32] | +56.6% | [+56.3%, +56.8%] | 0.3 pts | 1.3 | yes | yes |
| multi-str | street | 4096 | build | ordered | ordered-lpage-mvzc | 8 | 5.51 ms | 15.41 ms | 2.78× [2.75, 2.81] | +64.0% | [+63.7%, +64.4%] | 0.4 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | btree-sets | 8 | 108 | 265 | 2.39× [2.33, 2.44] | +58.1% | [+57.2%, +59.1%] | 1.1 pts | 1.2 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage | 8 | 98.6 | 99.0 | 1.00× [1.00, 1.01] | +0.2% | [-0.3%, +0.7%] | 0.6 pts | 0.9 | yes | no |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mv | 8 | 100 | 181 | 1.79× [1.76, 1.82] | +44.1% | [+43.2%, +45.0%] | 1.1 pts | 1.8 | yes | yes |
| multi-str | street | 16384 | valuesFor | ordered | ordered-lpage-mvzc | 8 | 99.6 | 158 | 1.57× [1.55, 1.59] | +36.3% | [+35.4%, +37.2%] | 1.1 pts | 2.8 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 3709 | 9915 | 2.70× [2.64, 2.76] | +62.9% | [+62.1%, +63.8%] | 1.0 pts | 2.3 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage | 8 | 3626 | 3621 | 0.99× [0.99, 1.00] | -0.6% | [-1.0%, -0.1%] | 0.6 pts | 0.8 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mv | 8 | 3686 | 2813 | 0.77× [0.76, 0.79] | -29.5% | [-32.0%, -27.0%] | 3.0 pts | 3.6 | yes | yes |
| multi-str | street | 16384 | valuesBetween | ordered | ordered-lpage-mvzc | 8 | 3581 | 2253 | 0.63× [0.62, 0.63] | -59.6% | [-60.8%, -58.5%] | 1.4 pts | 0.9 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | btree-sets | 8 | 1271 | 3834 | 3.02× [2.96, 3.08] | +66.9% | [+66.2%, +67.5%] | 0.8 pts | 1.1 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage | 8 | 1172 | 1187 | 1.00× [0.98, 1.01] | -0.5% | [-1.8%, +0.9%] | 1.6 pts | 0.8 | yes | no |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mv | 8 | 1223 | 990 | 0.81× [0.81, 0.82] | -22.9% | [-23.5%, -22.3%] | 0.7 pts | 0.3 | yes | yes |
| multi-str | street | 16384 | prefix | ordered | ordered-lpage-mvzc | 8 | 1179 | 788 | 0.66× [0.66, 0.67] | -51.0% | [-52.4%, -49.6%] | 1.7 pts | 0.7 | yes | yes |
| multi-str | street | 16384 | churn | ordered | btree-sets | 8 | 232 | 379 | 1.66× [1.62, 1.70] | +39.7% | [+38.2%, +41.2%] | 1.8 pts | 1.3 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage | 8 | 185 | 186 | 1.00× [0.99, 1.01] | +0.0% | [-0.5%, +0.6%] | 0.7 pts | 1.6 | yes | no |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mv | 8 | 191 | 344 | 1.74× [1.65, 1.86] | +42.7% | [+39.2%, +46.2%] | 4.1 pts | 6.0 | yes | yes |
| multi-str | street | 16384 | churn | ordered | ordered-lpage-mvzc | 8 | 226 | 441 | 1.97× [1.89, 2.04] | +49.1% | [+47.2%, +51.1%] | 2.3 pts | 2.8 | yes | yes |
| multi-str | street | 16384 | build | ordered | btree-sets | 8 | 29.62 ms | 53.68 ms | 1.83× [1.80, 1.87] | +45.4% | [+44.4%, +46.4%] | 1.2 pts | 1.6 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage | 8 | 28.91 ms | 32.19 ms | 1.10× [1.09, 1.12] | +9.3% | [+8.0%, +10.6%] | 1.6 pts | 2.1 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mv | 8 | 30.04 ms | 58.80 ms | 1.96× [1.92, 2.01] | +49.1% | [+47.9%, +50.3%] | 1.4 pts | 2.2 | yes | yes |
| multi-str | street | 16384 | build | ordered | ordered-lpage-mvzc | 8 | 29.58 ms | 72.84 ms | 2.44× [2.41, 2.48] | +59.1% | [+58.5%, +59.7%] | 0.7 pts | 1.6 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | btree-sets | 20 | 313 | 681 | 2.19× [2.17, 2.20] | +54.3% | [+53.9%, +54.6%] | 0.9 pts | 0.9 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage | 20 | 260 | 261 | 1.01× [1.00, 1.01] | +0.7% | [-0.0%, +1.4%] | 2.6 pts | 1.7 | yes | no |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mv | 20 | 278 | 308 | 1.12× [1.10, 1.13] | +10.4% | [+9.2%, +11.5%] | 2.4 pts | 1.4 | yes | yes |
| multi-str | street | 212449 | valuesFor | ordered | ordered-lpage-mvzc | 20 | 265 | 276 | 1.03× [1.03, 1.04] | +3.4% | [+2.8%, +3.9%] | 1.7 pts | 0.9 | yes | yes |
| multi-str | street | 212449 | valuesBetween | ordered | btree-sets | 20 | 6960 | 20.3 µs | 2.86× [2.82, 2.90] | +65.0% | [+64.5%, +65.5%] | 0.8 pts | 1.7 | yes | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage | 20 | 6263 | 6184 | 1.00× [0.99, 1.01] | -0.3% | [-1.1%, +0.5%] | 1.7 pts | 1.1 | yes | no |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mv | 20 | 6119 | 3941 | 0.62× [0.61, 0.63] | -60.5% | [-62.6%, -58.5%] | 6.6 pts | 3.0 | yes | yes |
| multi-str | street | 212449 | valuesBetween | ordered | ordered-lpage-mvzc | 20 | 5912 | 3131 | 0.52× [0.52, 0.53] | -91.2% | [-92.2%, -90.1%] | 5.2 pts | 1.9 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | btree-sets | 20 | 24.6 µs | 75.8 µs | 3.10× [3.06, 3.15] | +67.8% | [+67.3%, +68.2%] | 1.0 pts | 0.7 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage | 20 | 23.5 µs | 23.3 µs | 1.02× [1.00, 1.04] | +2.2% | [+0.2%, +4.1%] | 3.3 pts | 1.2 | yes | no |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mv | 20 | 22.1 µs | 12.5 µs | 0.57× [0.56, 0.57] | -76.1% | [-77.6%, -74.6%] | 5.5 pts | 1.6 | yes | yes |
| multi-str | street | 212449 | prefix | ordered | ordered-lpage-mvzc | 20 | 20.6 µs | 9536 | 0.46× [0.45, 0.46] | -118.5% | [-119.9%, -117.1%] | 5.6 pts | 1.8 | yes | yes |
| multi-str | street | 212449 | churn | ordered | btree-sets | 20 | 662 | 903 | 1.38× [1.36, 1.39] | +27.5% | [+26.7%, +28.3%] | 1.6 pts | 0.9 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage | 20 | 581 | 513 | 0.84× [0.83, 0.85] | -19.0% | [-20.4%, -17.5%] | 5.9 pts | 0.5 | yes | yes |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mv | 20 | 598 | 630 | 0.99× [0.97, 1.01] | -1.1% | [-2.7%, +0.6%] | 3.9 pts | 0.4 | yes | no |
| multi-str | street | 212449 | churn | ordered | ordered-lpage-mvzc | 20 | 626 | 715 | 1.11× [1.09, 1.12] | +9.7% | [+8.5%, +10.9%] | 2.4 pts | 0.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str street n=4096 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=4096 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +0.72% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage: the pooled difference of 0.04% does not clear the 1.13% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage: the A/A validations found a systematic difference of -0.59% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=4096 valuesBetween: ordered vs ordered-lpage: the pooled interval [-0.42%, 0.50%] includes zero
- multi-str street n=4096 build: ordered vs ordered-lpage: the pooled interval [-0.28%, 1.33%] includes zero
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage: the pooled difference of 0.18% does not clear the 0.40% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage: the pooled interval [-0.35%, 0.71%] includes zero
- multi-str street n=16384 valuesFor: ordered vs ordered-lpage-mvzc: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 prefix: ordered vs ordered-lpage: the pooled difference of -0.47% does not clear the 1.48% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 prefix: ordered vs ordered-lpage: the pooled interval [-1.83%, 0.88%] includes zero
- multi-str street n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.97% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 churn: ordered vs ordered-lpage: the pooled difference of 0.04% does not clear the 0.49% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=16384 churn: ordered vs ordered-lpage: the pooled interval [-0.53%, 0.62%] includes zero
- multi-str street n=16384 churn: ordered vs ordered-lpage-mv: the processes scatter 6.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 churn: ordered vs ordered-lpage-mvzc: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 build: ordered vs btree-sets: the A/A validations found a systematic difference of -0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=16384 build: ordered vs ordered-lpage: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 build: ordered vs ordered-lpage-mv: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=16384 build: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of -1.02% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of +0.79% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 valuesFor: ordered vs ordered-lpage: the pooled interval [-0.03%, 1.41%] includes zero
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: the pooled difference of -0.28% does not clear the 0.39% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage: the pooled interval [-1.10%, 0.55%] includes zero
- multi-str street n=212449 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str street n=212449 prefix: ordered vs ordered-lpage: the pooled difference of 2.18% does not clear the 2.32% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 prefix: ordered vs ordered-lpage: the A/A validations found a systematic difference of -1.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 prefix: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of +1.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 churn: ordered vs btree-sets: the A/A validations found a systematic difference of -0.71% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 churn: ordered vs ordered-lpage: the A/A validations found a systematic difference of +0.85% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 churn: ordered vs ordered-lpage-mv: the pooled difference of -1.07% does not clear the 2.46% noise floor, the bound on what the harness reports between identical code in every process
- multi-str street n=212449 churn: ordered vs ordered-lpage-mv: the A/A validations found a systematic difference of +1.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str street n=212449 churn: ordered vs ordered-lpage-mv: the pooled interval [-2.71%, 0.57%] includes zero
