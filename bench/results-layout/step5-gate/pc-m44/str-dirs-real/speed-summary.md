| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | dirs | 4096 | valuesFor | ordered | baseline | 8 | 139 | 125 | 0.90× [0.89, 0.90] | -11.5% | [-12.2%, -10.7%] | 0.9 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 141 | 171 | 1.20× [1.16, 1.24] | +16.4% | [+13.6%, +19.2%] | 3.1 pts | 3.2 | no | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 6398 | 6540 | 1.02× [1.01, 1.03] | +2.3% | [+1.3%, +3.3%] | 1.1 pts | 1.0 | yes | yes |
| natural-str | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 6413 | 10.5 µs | 1.64× [1.60, 1.67] | +38.9% | [+37.6%, +40.3%] | 1.3 pts | 2.3 | yes | yes |
| natural-str | dirs | 4096 | prefix | ordered | baseline | 8 | 7223 | 7469 | 1.02× [1.00, 1.04] | +2.2% | [+0.1%, +4.3%] | 3.5 pts | 0.9 | no | no |
| natural-str | dirs | 4096 | prefix | ordered | btree-sets | 8 | 7369 | 15.2 µs | 2.02× [1.93, 2.12] | +50.5% | [+48.2%, +52.8%] | 2.2 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | baseline | 8 | 269 | 189 | 0.70× [0.70, 0.71] | -42.0% | [-43.7%, -40.2%] | 2.1 pts | 1.3 | yes | yes |
| natural-str | dirs | 4096 | churn | ordered | btree-sets | 8 | 269 | 245 | 0.91× [0.89, 0.92] | -10.4% | [-12.2%, -8.6%] | 2.3 pts | 1.5 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | baseline | 8 | 12.69 ms | 8.32 ms | 0.66× [0.65, 0.66] | -52.1% | [-52.9%, -51.4%] | 1.1 pts | 0.9 | yes | yes |
| natural-str | dirs | 4096 | build | ordered | btree-sets | 8 | 12.78 ms | 11.43 ms | 0.90× [0.89, 0.90] | -11.7% | [-12.1%, -11.2%] | 0.7 pts | 0.9 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | baseline | 8 | 172 | 161 | 0.93× [0.92, 0.94] | -7.5% | [-8.4%, -6.6%] | 1.3 pts | 1.5 | yes | yes |
| natural-str | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 178 | 233 | 1.32× [1.28, 1.35] | +24.0% | [+22.0%, +26.0%] | 1.9 pts | 2.0 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 6940 | 7371 | 1.06× [1.05, 1.07] | +5.4% | [+4.6%, +6.3%] | 0.9 pts | 1.1 | yes | yes |
| natural-str | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 6963 | 11.5 µs | 1.65× [1.63, 1.67] | +39.5% | [+38.8%, +40.2%] | 0.8 pts | 1.0 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | baseline | 8 | 34.0 µs | 38.4 µs | 1.13× [1.12, 1.14] | +11.5% | [+10.7%, +12.3%] | 0.8 pts | 1.1 | yes | yes |
| natural-str | dirs | 16384 | prefix | ordered | btree-sets | 8 | 30.4 µs | 59.9 µs | 1.96× [1.94, 1.99] | +49.0% | [+48.4%, +49.6%] | 0.7 pts | 2.0 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | baseline | 8 | 328 | 260 | 0.80× [0.80, 0.81] | -24.5% | [-25.1%, -24.0%] | 2.2 pts | 1.9 | yes | yes |
| natural-str | dirs | 16384 | churn | ordered | btree-sets | 8 | 334 | 386 | 1.14× [1.12, 1.17] | +12.5% | [+10.8%, +14.3%] | 1.8 pts | 2.3 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | baseline | 8 | 63.28 ms | 46.15 ms | 0.73× [0.73, 0.74] | -36.3% | [-37.5%, -35.0%] | 1.2 pts | 1.1 | yes | yes |
| natural-str | dirs | 16384 | build | ordered | btree-sets | 8 | 63.25 ms | 67.23 ms | 1.06× [1.05, 1.07] | +5.4% | [+4.6%, +6.3%] | 0.9 pts | 1.0 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | baseline | 8 | 254 | 272 | 1.07× [1.06, 1.09] | +6.5% | [+5.2%, +7.9%] | 1.2 pts | 1.5 | yes | yes |
| natural-str | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 286 | 496 | 1.74× [1.69, 1.78] | +42.5% | [+40.9%, +44.0%] | 1.6 pts | 2.2 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 9232 | 11.0 µs | 1.21× [1.19, 1.23] | +17.3% | [+16.2%, +18.4%] | 1.5 pts | 1.0 | yes | yes |
| natural-str | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 9794 | 23.6 µs | 2.39× [2.33, 2.45] | +58.1% | [+57.1%, +59.1%] | 1.0 pts | 1.2 | yes | yes |
| natural-str | dirs | 86215 | prefix | ordered | baseline | 8 | 208.2 µs | 257.9 µs | 1.19× [1.14, 1.24] | +15.7% | [+12.2%, +19.3%] | 4.3 pts | 1.0 | no | yes |
| natural-str | dirs | 86215 | prefix | ordered | btree-sets | 8 | 214.5 µs | 514.3 µs | 2.53× [2.41, 2.65] | +60.4% | [+58.5%, +62.3%] | 2.0 pts | 0.9 | yes | yes |
| natural-str | dirs | 86215 | churn | ordered | baseline | 8 | 579 | 540 | 0.94× [0.91, 0.97] | -6.4% | [-9.9%, -2.8%] | 3.9 pts | 1.0 | no | yes |
| natural-str | dirs | 86215 | churn | ordered | btree-sets | 8 | 631 | 716 | 1.12× [1.11, 1.14] | +11.0% | [+9.9%, +12.2%] | 1.1 pts | 0.9 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str dirs n=4096 valuesFor: ordered vs btree-sets: the A/A validations found a systematic difference of -0.58% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=4096 valuesFor: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=4096 prefix: ordered vs baseline: the pooled difference of 2.18% does not clear the 5.46% noise floor, the bound on what the harness reports between identical code in every process
- natural-str dirs n=16384 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of -0.45% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -5.09% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str dirs n=16384 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -2.77% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
