| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | street | 4096 | valuesFor | ordered | baseline | 6 | 96.2 | 82.5 | 0.86× [0.85, 0.87] | -16.9% | [-18.3%, -15.5%] | 1.3 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 6 | 97.2 | 140 | 1.45× [1.42, 1.47] | +30.8% | [+29.7%, +31.9%] | 1.0 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 6 | 4683 | 5152 | 1.11× [1.09, 1.13] | +10.1% | [+8.5%, +11.6%] | 1.5 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 6 | 4648 | 7131 | 1.54× [1.51, 1.56] | +34.9% | [+33.7%, +36.1%] | 1.1 pts | 1.8 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 6 | 516 | 473 | 0.92× [0.91, 0.93] | -8.9% | [-10.0%, -7.8%] | 1.0 pts | 0.5 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 6 | 520 | 729 | 1.39× [1.35, 1.42] | +27.9% | [+26.1%, +29.7%] | 1.7 pts | 1.4 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 6 | 219 | 124 | 0.57× [0.56, 0.58] | -74.7% | [-77.8%, -71.7%] | 2.9 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 6 | 220 | 207 | 0.93× [0.92, 0.94] | -7.5% | [-8.9%, -6.1%] | 1.3 pts | 0.9 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 6 | 7.97 ms | 4.73 ms | 0.59× [0.58, 0.60] | -68.5% | [-71.5%, -65.5%] | 2.9 pts | 2.1 | yes | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 6 | 8.03 ms | 7.72 ms | 0.97× [0.96, 0.97] | -3.3% | [-3.8%, -2.7%] | 0.5 pts | 0.6 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 6 | 114 | 105 | 0.92× [0.92, 0.93] | -8.4% | [-9.2%, -7.5%] | 0.8 pts | 1.0 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 6 | 121 | 188 | 1.58× [1.55, 1.61] | +36.7% | [+35.5%, +37.9%] | 1.1 pts | 2.3 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 5163 | 5815 | 1.12× [1.10, 1.14] | +10.9% | [+9.4%, +12.3%] | 1.4 pts | 2.4 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 5221 | 8305 | 1.59× [1.57, 1.61] | +37.1% | [+36.3%, +37.9%] | 0.8 pts | 2.1 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 6 | 1725 | 1940 | 1.12× [1.11, 1.14] | +11.0% | [+9.5%, +12.4%] | 1.3 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 6 | 1743 | 3033 | 1.75× [1.71, 1.79] | +42.9% | [+41.6%, +44.3%] | 1.3 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 6 | 238 | 178 | 0.73× [0.72, 0.74] | -37.0% | [-39.0%, -35.0%] | 1.9 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 6 | 240 | 293 | 1.22× [1.21, 1.24] | +18.3% | [+17.4%, +19.2%] | 0.8 pts | 0.9 | yes | yes |
| natural-str | street | 16384 | build | ordered | baseline | 6 | 37.55 ms | 25.25 ms | 0.67× [0.67, 0.68] | -48.6% | [-49.9%, -47.3%] | 1.2 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 6 | 38.18 ms | 45.25 ms | 1.18× [1.16, 1.20] | +15.3% | [+14.1%, +16.4%] | 1.1 pts | 1.5 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | baseline | 8 | 207 | 255 | 1.22× [1.19, 1.24] | +17.9% | [+16.2%, +19.7%] | 2.6 pts | 3.3 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | btree-sets | 8 | 240 | 503 | 2.05× [2.01, 2.10] | +51.3% | [+50.3%, +52.3%] | 1.1 pts | 2.0 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 6667 | 9411 | 1.42× [1.40, 1.44] | +29.8% | [+28.7%, +30.8%] | 1.0 pts | 1.0 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 7076 | 22.8 µs | 3.21× [3.13, 3.28] | +68.8% | [+68.1%, +69.6%] | 0.8 pts | 2.3 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | baseline | 8 | 26.4 µs | 36.0 µs | 1.38× [1.35, 1.41] | +27.5% | [+26.0%, +29.1%] | 1.6 pts | 2.1 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | btree-sets | 8 | 25.8 µs | 74.7 µs | 2.91× [2.85, 2.97] | +65.6% | [+64.9%, +66.3%] | 0.8 pts | 1.4 | yes | yes |
| natural-str | street | 212449 | churn | ordered | baseline | 8 | 551 | 545 | 0.98× [0.96, 1.00] | -2.0% | [-4.2%, +0.2%] | 2.8 pts | 0.7 | no | no |
| natural-str | street | 212449 | churn | ordered | btree-sets | 8 | 632 | 774 | 1.23× [1.21, 1.25] | +18.8% | [+17.4%, +20.3%] | 1.4 pts | 1.6 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str street n=4096 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 prefix: ordered vs baseline: the A/A validations found a systematic difference of +1.15% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=16384 build: ordered vs baseline: the A/A validations found a systematic difference of -0.23% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=212449 valuesFor: ordered vs baseline: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of +0.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 prefix: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 churn: ordered vs baseline: the pooled interval [-4.17%, 0.15%] includes zero
