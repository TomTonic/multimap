| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 6 | 87.3 | 83.0 | 0.96× [0.94, 0.98] | -4.3% | [-6.1%, -2.4%] | 1.7 pts | 1.4 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 6 | 88.1 | 160 | 1.84× [1.82, 1.87] | +45.8% | [+45.1%, +46.4%] | 0.7 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 3041 | 2657 | 0.87× [0.86, 0.89] | -14.5% | [-16.5%, -12.6%] | 1.9 pts | 0.6 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 6 | 3071 | 5451 | 1.78× [1.76, 1.81] | +43.9% | [+43.1%, +44.8%] | 0.8 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 6 | 3357 | 2873 | 0.85× [0.84, 0.86] | -17.2% | [-18.6%, -15.9%] | 1.3 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 6 | 3413 | 7544 | 2.22× [2.16, 2.28] | +54.9% | [+53.8%, +56.1%] | 1.1 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 6 | 152 | 143 | 0.94× [0.92, 0.95] | -6.8% | [-8.1%, -5.5%] | 1.3 pts | 1.0 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 6 | 152 | 225 | 1.48× [1.46, 1.51] | +32.6% | [+31.6%, +33.6%] | 1.0 pts | 0.9 | yes | yes |
| natural | dirs | 4096 | build | ordered | baseline | 6 | 7.31 ms | 6.71 ms | 0.92× [0.91, 0.93] | -8.6% | [-10.1%, -7.2%] | 1.4 pts | 2.1 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 6 | 7.45 ms | 10.96 ms | 1.49× [1.44, 1.55] | +33.1% | [+30.7%, +35.4%] | 2.2 pts | 2.3 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 8 | 118 | 115 | 0.96× [0.91, 1.00] | -4.7% | [-9.4%, +0.1%] | 4.6 pts | 4.2 | no | no |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 8 | 125 | 224 | 1.85× [1.70, 2.02] | +45.9% | [+41.2%, +50.6%] | 5.1 pts | 6.6 | no | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 8 | 3782 | 3498 | 0.91× [0.86, 0.95] | -10.5% | [-16.0%, -5.0%] | 5.4 pts | 4.7 | no | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 8 | 4025 | 6360 | 1.69× [1.52, 1.90] | +40.8% | [+34.1%, +47.5%] | 7.5 pts | 5.4 | no | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 8 | 17.0 µs | 15.2 µs | 0.89× [0.87, 0.91] | -12.5% | [-15.6%, -9.5%] | 3.1 pts | 0.9 | no | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 8 | 17.6 µs | 36.0 µs | 1.97× [1.86, 2.09] | +49.2% | [+46.3%, +52.1%] | 3.1 pts | 3.2 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 8 | 205 | 195 | 0.96× [0.95, 0.96] | -4.5% | [-5.2%, -3.9%] | 0.8 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 8 | 212 | 330 | 1.56× [1.52, 1.59] | +35.7% | [+34.3%, +37.2%] | 1.4 pts | 1.9 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 8 | 38.83 ms | 36.37 ms | 0.94× [0.93, 0.94] | -6.5% | [-7.2%, -5.9%] | 0.8 pts | 1.1 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 8 | 39.65 ms | 61.18 ms | 1.52× [1.46, 1.58] | +34.3% | [+31.7%, +36.8%] | 2.4 pts | 4.7 | yes | yes |
| natural | dirs | 86215 | valuesFor | ordered | baseline | 8 | 180 | 177 | 0.98× [0.96, 1.01] | -1.7% | [-4.5%, +1.0%] | 2.6 pts | 1.7 | no | no |
| natural | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 221 | 421 | 1.89× [1.85, 1.93] | +47.0% | [+45.8%, +48.2%] | 1.2 pts | 1.4 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 4590 | 4353 | 0.95× [0.93, 0.96] | -5.6% | [-7.3%, -3.8%] | 2.3 pts | 2.6 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 6388 | 12.7 µs | 2.02× [1.98, 2.06] | +50.4% | [+49.5%, +51.4%] | 1.2 pts | 1.5 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | baseline | 8 | 123.2 µs | 113.4 µs | 0.93× [0.92, 0.94] | -7.5% | [-8.6%, -6.4%] | 1.4 pts | 1.0 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | btree-sets | 8 | 127.4 µs | 250.9 µs | 2.03× [2.02, 2.04] | +50.7% | [+50.4%, +50.9%] | 0.9 pts | 0.3 | yes | yes |
| natural | dirs | 86215 | churn | ordered | baseline | 8 | 388 | 378 | 0.97× [0.94, 1.00] | -3.5% | [-6.9%, -0.1%] | 4.0 pts | 1.2 | no | yes |
| natural | dirs | 86215 | churn | ordered | btree-sets | 8 | 504 | 640 | 1.28× [1.26, 1.30] | +21.9% | [+20.7%, +23.1%] | 1.2 pts | 1.0 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=4096 build: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs baseline: the pooled interval [-9.45%, 0.14%] includes zero
- natural dirs n=16384 valuesFor: ordered vs baseline: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesFor: ordered vs btree-sets: the processes scatter 6.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesBetween: ordered vs baseline: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 5.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 prefix: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=16384 churn: ordered vs btree-sets: the A/A validations found a systematic difference of +0.53% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=16384 build: ordered vs btree-sets: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-4.47%, 1.04%] includes zero
- natural dirs n=86215 valuesBetween: ordered vs baseline: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural dirs n=86215 valuesBetween: ordered vs btree-sets: the A/A validations found a systematic difference of +1.02% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural dirs n=86215 prefix: ordered vs btree-sets: the A/A validations found a systematic difference of -3.01% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
