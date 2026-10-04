| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| multi-str | dirs | 4096 | valuesFor | ordered | btree-sets | 12 | 85.3 | 162 | 1.91× [1.87, 1.95] | +47.7% | [+46.6%, +48.7%] | 1.1 pts | 1.8 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage | 12 | 84.2 | 84.7 | 1.00× [0.99, 1.02] | +0.5% | [-0.8%, +1.7%] | 1.7 pts | 1.4 | yes | no |
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mv | 12 | 83.8 | 148 | 1.78× [1.75, 1.80] | +43.7% | [+42.9%, +44.4%] | 0.8 pts | 1.2 | yes | yes |
| multi-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-mvzc | 12 | 84.4 | 120 | 1.42× [1.40, 1.45] | +29.8% | [+28.7%, +30.9%] | 1.2 pts | 1.3 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 12 | 2702 | 10.2 µs | 3.82× [3.77, 3.87] | +73.8% | [+73.4%, +74.2%] | 0.5 pts | 1.1 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage | 12 | 2716 | 2793 | 1.02× [1.00, 1.05] | +2.2% | [-0.2%, +4.6%] | 3.2 pts | 1.4 | no | no |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mv | 12 | 2697 | 3693 | 1.37× [1.33, 1.41] | +26.9% | [+24.8%, +29.0%] | 2.2 pts | 1.5 | yes | yes |
| multi-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-mvzc | 12 | 2702 | 2259 | 0.85× [0.83, 0.87] | -17.5% | [-20.1%, -15.0%] | 3.3 pts | 1.2 | no | yes |
| multi-str | dirs | 4096 | prefix | ordered | btree-sets | 12 | 2826 | 14.5 µs | 5.06× [4.97, 5.15] | +80.2% | [+79.9%, +80.6%] | 0.5 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage | 12 | 2855 | 2956 | 1.03× [1.02, 1.04] | +3.1% | [+2.2%, +4.1%] | 1.3 pts | 0.7 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mv | 12 | 2874 | 4289 | 1.50× [1.46, 1.54] | +33.3% | [+31.5%, +35.0%] | 1.6 pts | 0.9 | yes | yes |
| multi-str | dirs | 4096 | prefix | ordered | ordered-lpage-mvzc | 12 | 2838 | 2576 | 0.90× [0.89, 0.92] | -10.6% | [-12.3%, -8.8%] | 1.9 pts | 1.4 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | btree-sets | 12 | 167 | 236 | 1.42× [1.41, 1.43] | +29.4% | [+28.8%, +29.9%] | 0.9 pts | 1.0 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage | 12 | 166 | 166 | 1.01× [1.00, 1.02] | +0.8% | [-0.3%, +2.0%] | 1.5 pts | 1.4 | yes | no |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mv | 12 | 169 | 286 | 1.71× [1.67, 1.75] | +41.4% | [+40.1%, +42.8%] | 1.5 pts | 2.5 | yes | yes |
| multi-str | dirs | 4096 | churn | ordered | ordered-lpage-mvzc | 12 | 171 | 369 | 2.16× [2.12, 2.20] | +53.7% | [+52.9%, +54.6%] | 1.0 pts | 1.6 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | btree-sets | 12 | 8.22 ms | 11.93 ms | 1.45× [1.44, 1.46] | +31.2% | [+30.7%, +31.7%] | 0.8 pts | 0.7 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage | 12 | 8.09 ms | 8.34 ms | 1.04× [1.04, 1.04] | +3.7% | [+3.5%, +3.9%] | 0.5 pts | 0.8 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mv | 12 | 8.16 ms | 16.50 ms | 2.04× [2.00, 2.08] | +51.0% | [+50.0%, +51.9%] | 0.9 pts | 1.9 | yes | yes |
| multi-str | dirs | 4096 | build | ordered | ordered-lpage-mvzc | 12 | 8.30 ms | 19.98 ms | 2.41× [2.36, 2.47] | +58.6% | [+57.6%, +59.5%] | 0.9 pts | 1.7 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | btree-sets | 12 | 121 | 225 | 1.87× [1.84, 1.91] | +46.6% | [+45.7%, +47.5%] | 1.0 pts | 1.4 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage | 12 | 118 | 118 | 1.00× [0.98, 1.02] | +0.0% | [-1.8%, +1.8%] | 1.9 pts | 1.8 | yes | no |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mv | 12 | 119 | 192 | 1.55× [1.53, 1.58] | +35.6% | [+34.5%, +36.7%] | 2.8 pts | 3.5 | yes | yes |
| multi-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-mvzc | 12 | 118 | 153 | 1.28× [1.27, 1.30] | +22.1% | [+21.1%, +23.2%] | 2.3 pts | 2.7 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 12 | 3556 | 11.2 µs | 3.16× [3.12, 3.21] | +68.4% | [+67.9%, +68.8%] | 1.0 pts | 2.1 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage | 12 | 3553 | 3587 | 1.01× [1.00, 1.02] | +0.7% | [-0.2%, +1.5%] | 1.0 pts | 1.1 | yes | no |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mv | 12 | 3525 | 4132 | 1.17× [1.16, 1.18] | +14.5% | [+13.7%, +15.4%] | 1.1 pts | 1.5 | yes | yes |
| multi-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-mvzc | 12 | 3567 | 2600 | 0.73× [0.73, 0.74] | -36.2% | [-37.7%, -34.7%] | 1.9 pts | 1.5 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | btree-sets | 12 | 16.8 µs | 60.9 µs | 3.84× [3.76, 3.92] | +73.9% | [+73.4%, +74.5%] | 0.7 pts | 0.9 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage | 12 | 15.7 µs | 16.0 µs | 1.01× [1.00, 1.02] | +0.9% | [+0.1%, +1.7%] | 1.8 pts | 0.6 | yes | no |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mv | 12 | 16.8 µs | 20.4 µs | 1.20× [1.19, 1.21] | +16.3% | [+15.6%, +17.0%] | 1.9 pts | 1.2 | yes | yes |
| multi-str | dirs | 16384 | prefix | ordered | ordered-lpage-mvzc | 12 | 15.6 µs | 10.6 µs | 0.67× [0.64, 0.70] | -50.2% | [-57.0%, -43.4%] | 7.5 pts | 0.7 | no | yes |
| multi-str | dirs | 16384 | churn | ordered | btree-sets | 12 | 254 | 377 | 1.44× [1.43, 1.45] | +30.6% | [+30.0%, +31.2%] | 2.6 pts | 3.1 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage | 12 | 234 | 234 | 1.00× [0.99, 1.00] | -0.4% | [-0.7%, -0.0%] | 0.5 pts | 0.7 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mv | 12 | 253 | 359 | 1.41× [1.38, 1.45] | +29.2% | [+27.6%, +30.9%] | 3.8 pts | 5.5 | yes | yes |
| multi-str | dirs | 16384 | churn | ordered | ordered-lpage-mvzc | 12 | 263 | 453 | 1.66× [1.64, 1.68] | +39.7% | [+39.1%, +40.4%] | 3.8 pts | 5.2 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | btree-sets | 12 | 44.58 ms | 66.80 ms | 1.49× [1.47, 1.51] | +32.9% | [+32.0%, +33.7%] | 1.0 pts | 1.5 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage | 12 | 45.78 ms | 49.31 ms | 1.08× [1.08, 1.09] | +7.6% | [+7.0%, +8.1%] | 1.3 pts | 1.2 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mv | 12 | 45.71 ms | 80.76 ms | 1.76× [1.74, 1.78] | +43.2% | [+42.7%, +43.7%] | 1.1 pts | 2.0 | yes | yes |
| multi-str | dirs | 16384 | build | ordered | ordered-lpage-mvzc | 12 | 44.67 ms | 99.34 ms | 2.21× [2.20, 2.22] | +54.7% | [+54.5%, +55.0%] | 0.5 pts | 1.1 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | btree-sets | 12 | 361 | 567 | 1.58× [1.55, 1.62] | +36.8% | [+35.3%, +38.3%] | 1.5 pts | 1.1 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage | 12 | 317 | 318 | 1.00× [0.99, 1.02] | +0.4% | [-1.5%, +2.4%] | 2.5 pts | 1.3 | yes | no |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mv | 12 | 328 | 354 | 1.09× [1.07, 1.11] | +8.0% | [+6.5%, +9.6%] | 2.3 pts | 1.2 | yes | yes |
| multi-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-mvzc | 12 | 317 | 285 | 0.91× [0.89, 0.93] | -10.2% | [-12.4%, -8.1%] | 2.3 pts | 1.1 | no | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 12 | 9620 | 31.6 µs | 3.28× [3.22, 3.35] | +69.5% | [+68.9%, +70.1%] | 0.7 pts | 1.0 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage | 12 | 8467 | 8501 | 1.00× [0.97, 1.03] | -0.1% | [-2.6%, +2.5%] | 3.1 pts | 2.4 | no | no |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mv | 12 | 8434 | 6757 | 0.80× [0.78, 0.81] | -25.5% | [-27.5%, -23.5%] | 3.5 pts | 2.2 | yes | yes |
| multi-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-mvzc | 12 | 7816 | 4495 | 0.56× [0.55, 0.58] | -77.1% | [-82.3%, -71.9%] | 7.5 pts | 3.3 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | btree-sets | 12 | 148.2 µs | 867.9 µs | 5.82× [5.56, 6.09] | +82.8% | [+82.0%, +83.6%] | 1.0 pts | 0.9 | yes | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage | 12 | 133.1 µs | 143.8 µs | 1.06× [1.00, 1.11] | +5.2% | [+0.2%, +10.3%] | 6.7 pts | 0.7 | no | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mv | 12 | 141.7 µs | 130.7 µs | 0.94× [0.90, 1.00] | -6.0% | [-11.6%, -0.3%] | 7.7 pts | 1.0 | no | yes |
| multi-str | dirs | 86215 | prefix | ordered | ordered-lpage-mvzc | 12 | 154.1 µs | 86.2 µs | 0.56× [0.54, 0.59] | -77.0% | [-84.6%, -69.5%] | 11.9 pts | 0.9 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | btree-sets | 12 | 671 | 873 | 1.29× [1.27, 1.32] | +22.6% | [+21.2%, +24.0%] | 1.9 pts | 2.3 | yes | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage | 12 | 628 | 621 | 0.94× [0.91, 0.98] | -6.0% | [-10.3%, -1.7%] | 5.3 pts | 1.0 | no | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mv | 12 | 637 | 740 | 1.15× [1.12, 1.18] | +13.3% | [+11.0%, +15.5%] | 2.9 pts | 1.3 | no | yes |
| multi-str | dirs | 86215 | churn | ordered | ordered-lpage-mvzc | 12 | 637 | 835 | 1.28× [1.25, 1.32] | +22.0% | [+19.8%, +24.2%] | 2.3 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- multi-str dirs n=4096 valuesFor: ordered vs ordered-lpage: the pooled difference of 0.47% does not clear the 1.05% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=4096 valuesFor: ordered vs ordered-lpage: the pooled interval [-0.80%, 1.74%] includes zero
- multi-str dirs n=4096 valuesFor: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of -0.46% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=4096 valuesBetween: ordered vs ordered-lpage: the pooled interval [-0.20%, 4.65%] includes zero
- multi-str dirs n=4096 churn: ordered vs ordered-lpage: the pooled interval [-0.29%, 1.98%] includes zero
- multi-str dirs n=4096 churn: ordered vs ordered-lpage-mv: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=4096 build: ordered vs ordered-lpage-mvzc: the A/A validations found a systematic difference of +0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the pooled difference of 0.01% does not clear the 0.49% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the pooled interval [-1.80%, 1.82%] includes zero
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage: 1 processes resolved A as faster and 3 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage-mv: the processes scatter 3.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesFor: ordered vs ordered-lpage-mvzc: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 valuesBetween: ordered vs ordered-lpage: the pooled interval [-0.16%, 1.51%] includes zero
- multi-str dirs n=16384 prefix: ordered vs ordered-lpage: the pooled difference of 0.90% does not clear the 12.23% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=16384 prefix: ordered vs ordered-lpage: the A/A validations found a systematic difference of -6.64% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=16384 churn: ordered vs btree-sets: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 churn: ordered vs ordered-lpage-mv: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=16384 churn: ordered vs ordered-lpage-mvzc: the processes scatter 5.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the pooled difference of 0.44% does not clear the 0.82% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the pooled interval [-1.51%, 2.39%] includes zero
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the pooled difference of -0.09% does not clear the 0.57% noise floor, the bound on what the harness reports between identical code in every process
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the pooled interval [-2.64%, 2.46%] includes zero
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: 2 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage-mv: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 valuesBetween: ordered vs ordered-lpage-mvzc: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- multi-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of +1.26% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- multi-str dirs n=86215 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
