| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 8 | 107 | 86.2 | 0.81× [0.80, 0.82] | -24.0% | [-25.5%, -22.5%] | 1.5 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 8 | 109 | 161 | 1.49× [1.46, 1.51] | +32.7% | [+31.7%, +33.8%] | 1.1 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 8 | 2841 | 3094 | 1.09× [1.06, 1.12] | +8.0% | [+5.4%, +10.5%] | 2.6 pts | 1.2 | no | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 8 | 2825 | 5544 | 1.97× [1.96, 1.98] | +49.2% | [+48.9%, +49.6%] | 0.6 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 8 | 3262 | 3406 | 1.02× [1.01, 1.03] | +1.9% | [+0.8%, +3.0%] | 1.9 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 8 | 3240 | 7653 | 2.37× [2.34, 2.41] | +57.9% | [+57.3%, +58.4%] | 1.1 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 8 | 226 | 160 | 0.71× [0.70, 0.72] | -40.8% | [-42.2%, -39.5%] | 1.4 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 8 | 226 | 241 | 1.07× [1.04, 1.11] | +6.8% | [+4.0%, +9.7%] | 2.7 pts | 2.3 | no | yes |
| natural | dirs | 4096 | build | ordered | baseline | 8 | 10.80 ms | 7.64 ms | 0.71× [0.70, 0.71] | -41.6% | [-43.1%, -40.0%] | 1.6 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 8 | 10.77 ms | 11.45 ms | 1.06× [1.04, 1.07] | +5.3% | [+3.9%, +6.7%] | 1.4 pts | 1.4 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 134 | 119 | 0.89× [0.87, 0.91] | -12.3% | [-14.8%, -9.9%] | 2.7 pts | 2.6 | no | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 136 | 229 | 1.68× [1.64, 1.72] | +40.5% | [+39.0%, +41.9%] | 1.5 pts | 1.9 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3218 | 3836 | 1.18× [1.17, 1.19] | +15.2% | [+14.2%, +16.3%] | 1.1 pts | 1.7 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 3234 | 6418 | 2.01× [1.94, 2.08] | +50.1% | [+48.4%, +51.9%] | 1.9 pts | 5.1 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 8 | 12.5 µs | 17.1 µs | 1.41× [1.34, 1.48] | +28.9% | [+25.3%, +32.4%] | 4.1 pts | 0.6 | no | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 8 | 12.4 µs | 33.5 µs | 2.77× [2.65, 2.90] | +63.9% | [+62.3%, +65.6%] | 1.7 pts | 0.6 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 270 | 216 | 0.80× [0.79, 0.81] | -25.1% | [-26.6%, -23.6%] | 1.5 pts | 1.3 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 275 | 356 | 1.29× [1.27, 1.31] | +22.7% | [+21.5%, +23.9%] | 1.2 pts | 1.2 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 52.58 ms | 42.43 ms | 0.80× [0.78, 0.82] | -25.7% | [-28.8%, -22.6%] | 3.6 pts | 2.6 | no | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 53.07 ms | 64.65 ms | 1.22× [1.19, 1.25] | +18.0% | [+15.9%, +20.1%] | 2.2 pts | 2.5 | no | yes |
| natural | dirs | 86215 | valuesFor | ordered | baseline | 8 | 193 | 188 | 0.99× [0.89, 1.13] | -0.6% | [-12.8%, +11.6%] | 12.7 pts | 9.2 | no | no |
| natural | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 208 | 418 | 2.01× [1.87, 2.17] | +50.2% | [+46.4%, +53.9%] | 4.2 pts | 4.6 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3851 | 5354 | 1.36× [1.18, 1.60] | +26.6% | [+15.5%, +37.7%] | 12.0 pts | 17.3 | no | yes |
| natural | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 4252 | 12.2 µs | 2.89× [2.69, 3.11] | +65.4% | [+62.8%, +67.9%] | 2.7 pts | 8.4 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | baseline | 8 | 90.0 µs | 128.2 µs | 1.47× [1.27, 1.74] | +32.0% | [+21.6%, +42.4%] | 10.4 pts | 8.4 | no | yes |
| natural | dirs | 86215 | prefix | ordered | btree-sets | 8 | 95.6 µs | 293.3 µs | 3.12× [2.74, 3.63] | +68.0% | [+63.5%, +72.4%] | 4.4 pts | 4.9 | yes | yes |
| natural | dirs | 86215 | churn | ordered | baseline | 8 | 443 | 432 | 0.96× [0.90, 1.03] | -4.1% | [-10.8%, +2.6%] | 6.5 pts | 4.4 | no | no |
| natural | dirs | 86215 | churn | ordered | btree-sets | 8 | 507 | 677 | 1.33× [1.29, 1.37] | +24.6% | [+22.4%, +26.8%] | 2.4 pts | 4.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 2.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-12.77%, 11.58%] includes zero
- natural dirs n=86215 valuesFor: ordered vs baseline: the processes scatter 9.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesFor: ordered vs baseline: 2 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- natural dirs n=86215 valuesFor: ordered vs btree-sets: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 17.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesBetween: ordered vs btree-sets: the processes scatter 8.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 prefix: ordered vs baseline: the A/A validations found a systematic difference of -1.73% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=86215 prefix: ordered vs baseline: the processes scatter 8.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 prefix: ordered vs btree-sets: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 churn: ordered vs baseline: the pooled interval [-10.78%, 2.64%] includes zero
- natural dirs n=86215 churn: ordered vs baseline: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 churn: ordered vs btree-sets: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
