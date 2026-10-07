| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | street | 4096 | valuesFor | ordered | baseline | 8 | 98.7 | 81.8 | 0.82× [0.82, 0.83] | -21.2% | [-21.7%, -20.7%] | 0.8 pts | 0.7 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 99.6 | 140 | 1.40× [1.39, 1.42] | +28.8% | [+28.3%, +29.3%] | 0.8 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 4961 | 5293 | 1.07× [1.06, 1.07] | +6.4% | [+5.9%, +7.0%] | 0.7 pts | 0.6 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 4941 | 7307 | 1.48× [1.46, 1.50] | +32.5% | [+31.7%, +33.3%] | 0.8 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 8 | 536 | 482 | 0.90× [0.88, 0.92] | -11.4% | [-13.6%, -9.2%] | 2.2 pts | 1.4 | no | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 8 | 539 | 741 | 1.38× [1.37, 1.40] | +27.8% | [+26.9%, +28.7%] | 0.9 pts | 1.0 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 8 | 232 | 130 | 0.56× [0.55, 0.57] | -79.2% | [-81.9%, -76.4%] | 2.6 pts | 1.2 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 8 | 236 | 209 | 0.89× [0.88, 0.90] | -12.4% | [-14.0%, -10.8%] | 1.7 pts | 1.4 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 8 | 8.87 ms | 5.32 ms | 0.60× [0.58, 0.61] | -67.8% | [-71.8%, -63.8%] | 4.0 pts | 1.2 | yes | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 8 | 8.99 ms | 8.62 ms | 0.97× [0.94, 1.00] | -3.1% | [-6.7%, +0.5%] | 3.9 pts | 3.0 | no | no |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 8 | 116 | 107 | 0.93× [0.91, 0.95] | -7.6% | [-10.4%, -4.8%] | 2.7 pts | 2.8 | no | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 8 | 123 | 196 | 1.65× [1.57, 1.75] | +39.6% | [+36.3%, +42.8%] | 3.3 pts | 4.2 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 8 | 5503 | 6021 | 1.10× [1.08, 1.12] | +9.1% | [+7.5%, +10.7%] | 1.6 pts | 2.1 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 8 | 5588 | 8865 | 1.65× [1.53, 1.79] | +39.4% | [+34.5%, +44.2%] | 5.1 pts | 7.1 | no | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 8 | 1829 | 1977 | 1.07× [1.06, 1.08] | +6.2% | [+5.3%, +7.0%] | 1.5 pts | 0.9 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 8 | 1839 | 3183 | 1.76× [1.69, 1.83] | +43.1% | [+40.8%, +45.4%] | 2.3 pts | 3.2 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 8 | 267 | 197 | 0.73× [0.72, 0.75] | -36.2% | [-38.8%, -33.6%] | 2.6 pts | 1.8 | yes | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 8 | 297 | 322 | 1.11× [1.08, 1.14] | +9.9% | [+7.3%, +12.5%] | 3.1 pts | 2.7 | no | yes |
| natural-str | street | 16384 | build | ordered | baseline | 8 | 44.30 ms | 30.55 ms | 0.70× [0.68, 0.72] | -43.0% | [-47.9%, -38.1%] | 5.2 pts | 3.7 | no | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 8 | 44.60 ms | 49.25 ms | 1.11× [1.08, 1.15] | +10.2% | [+7.2%, +13.2%] | 3.1 pts | 3.4 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | baseline | 8 | 200 | 235 | 1.18× [1.13, 1.23] | +15.0% | [+11.4%, +18.5%] | 5.1 pts | 4.8 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | btree-sets | 8 | 257 | 554 | 2.16× [2.09, 2.23] | +53.7% | [+52.1%, +55.3%] | 2.8 pts | 4.4 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 7169 | 9220 | 1.30× [1.26, 1.35] | +23.3% | [+20.8%, +25.7%] | 3.3 pts | 4.9 | no | yes |
| natural-str | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 7791 | 23.7 µs | 3.02× [2.87, 3.19] | +66.9% | [+65.2%, +68.7%] | 1.8 pts | 5.5 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | baseline | 8 | 26.8 µs | 35.9 µs | 1.32× [1.30, 1.35] | +24.4% | [+23.1%, +25.8%] | 1.6 pts | 3.1 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | btree-sets | 8 | 26.5 µs | 78.6 µs | 2.98× [2.85, 3.12] | +66.4% | [+64.9%, +67.9%] | 2.5 pts | 6.7 | yes | yes |
| natural-str | street | 212449 | churn | ordered | baseline | 8 | 567 | 558 | 0.99× [0.96, 1.02] | -1.1% | [-3.8%, +1.6%] | 2.6 pts | 0.8 | no | no |
| natural-str | street | 212449 | churn | ordered | btree-sets | 8 | 640 | 772 | 1.22× [1.20, 1.25] | +18.3% | [+16.7%, +20.0%] | 1.6 pts | 2.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str street n=4096 build: ordered vs btree-sets: the pooled interval [-6.67%, 0.45%] includes zero
- natural-str street n=4096 build: ordered vs btree-sets: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs baseline: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 4.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 7.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 prefix: ordered vs btree-sets: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 churn: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 build: ordered vs baseline: the processes scatter 3.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 build: ordered vs btree-sets: the processes scatter 3.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 4.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs btree-sets: the processes scatter 4.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 5.5 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs baseline: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs btree-sets: the processes scatter 6.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 churn: ordered vs baseline: the pooled difference of -1.10% does not clear the 1.67% noise floor, the bound on what the harness reports between identical code in every process
- natural-str street n=212449 churn: ordered vs baseline: the pooled interval [-3.83%, 1.64%] includes zero
- natural-str street n=212449 churn: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
