| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural-str | street | 4096 | valuesFor | ordered | baseline | 8 | 82.5 | 54.4 | 0.66× [0.65, 0.67] | -51.9% | [-54.7%, -49.1%] | 3.0 pts | 2.1 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | btree-sets | 8 | 83.7 | 141 | 1.68× [1.66, 1.70] | +40.4% | [+39.7%, +41.0%] | 0.6 pts | 1.3 | yes | yes |
| natural-str | street | 4096 | valuesFor | ordered | hashed | 8 | 82.6 | 24.4 | 0.29× [0.29, 0.30] | -240.6% | [-243.8%, -237.3%] | 3.2 pts | 1.4 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | baseline | 8 | 5046 | 2444 | 0.48× [0.47, 0.48] | -109.8% | [-113.4%, -106.2%] | 4.1 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | btree-sets | 8 | 5075 | 7125 | 1.40× [1.37, 1.42] | +28.4% | [+27.1%, +29.6%] | 1.2 pts | 1.7 | yes | yes |
| natural-str | street | 4096 | valuesBetween | ordered | hashed | 8 | 5250 | 54.9 µs | 10.50× [10.36, 10.63] | +90.5% | [+90.4%, +90.6%] | 0.1 pts | 1.5 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | baseline | 8 | 461 | 236 | 0.51× [0.50, 0.52] | -96.2% | [-98.8%, -93.7%] | 2.6 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | btree-sets | 8 | 462 | 718 | 1.54× [1.51, 1.58] | +35.1% | [+33.7%, +36.6%] | 1.7 pts | 1.8 | yes | yes |
| natural-str | street | 4096 | prefix | ordered | hashed | 8 | 520 | 50.5 µs | 96.22× [93.25, 99.38] | +99.0% | [+98.9%, +99.0%] | 0.0 pts | 0.9 | yes | yes |
| natural-str | street | 4096 | churn | ordered | baseline | 8 | 128 | 106 | 0.82× [0.81, 0.84] | -21.3% | [-23.4%, -19.3%] | 2.1 pts | 1.8 | yes | yes |
| natural-str | street | 4096 | churn | ordered | btree-sets | 8 | 131 | 204 | 1.54× [1.51, 1.57] | +35.2% | [+34.0%, +36.5%] | 1.2 pts | 1.1 | yes | yes |
| natural-str | street | 4096 | churn | ordered | hashed | 8 | 127 | 56.7 | 0.45× [0.44, 0.45] | -124.3% | [-126.9%, -121.6%] | 2.7 pts | 1.0 | yes | yes |
| natural-str | street | 4096 | build | ordered | baseline | 8 | 4.91 ms | 4.34 ms | 0.88× [0.86, 0.90] | -14.1% | [-16.6%, -11.5%] | 2.4 pts | 1.6 | no | yes |
| natural-str | street | 4096 | build | ordered | btree-sets | 8 | 5.02 ms | 8.09 ms | 1.62× [1.60, 1.63] | +38.2% | [+37.6%, +38.8%] | 1.0 pts | 0.9 | yes | yes |
| natural-str | street | 4096 | build | ordered | hashed | 8 | 4.93 ms | 2.77 ms | 0.56× [0.56, 0.57] | -77.3% | [-80.0%, -74.6%] | 5.8 pts | 1.6 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | baseline | 6 | 104 | 76.0 | 0.73× [0.72, 0.73] | -37.1% | [-38.0%, -36.2%] | 0.9 pts | 0.8 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | btree-sets | 6 | 107 | 190 | 1.76× [1.73, 1.81] | +43.3% | [+42.1%, +44.6%] | 1.2 pts | 2.7 | yes | yes |
| natural-str | street | 16384 | valuesFor | ordered | hashed | 6 | 105 | 30.1 | 0.29× [0.28, 0.29] | -250.3% | [-257.3%, -243.3%] | 6.7 pts | 3.0 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | baseline | 6 | 5775 | 3018 | 0.53× [0.52, 0.53] | -90.3% | [-93.5%, -87.1%] | 3.0 pts | 2.3 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | btree-sets | 6 | 5732 | 8218 | 1.42× [1.40, 1.45] | +29.6% | [+28.4%, +30.8%] | 1.1 pts | 2.0 | yes | yes |
| natural-str | street | 16384 | valuesBetween | ordered | hashed | 6 | 6422 | 235.4 µs | 36.14× [34.82, 37.57] | +97.2% | [+97.1%, +97.3%] | 0.1 pts | 1.7 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | baseline | 6 | 1870 | 889 | 0.47× [0.46, 0.48] | -111.0% | [-115.3%, -106.8%] | 4.1 pts | 1.2 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | btree-sets | 6 | 1929 | 2997 | 1.58× [1.53, 1.62] | +36.6% | [+34.8%, +38.4%] | 1.7 pts | 1.7 | yes | yes |
| natural-str | street | 16384 | prefix | ordered | hashed | 6 | 2419 | 233.3 µs | 94.90× [91.91, 98.09] | +98.9% | [+98.9%, +99.0%] | 0.0 pts | 0.7 | yes | yes |
| natural-str | street | 16384 | churn | ordered | baseline | 6 | 179 | 149 | 0.83× [0.82, 0.85] | -20.0% | [-21.9%, -18.1%] | 1.8 pts | 1.5 | yes | yes |
| natural-str | street | 16384 | churn | ordered | btree-sets | 6 | 181 | 291 | 1.59× [1.57, 1.61] | +37.0% | [+36.3%, +37.8%] | 0.7 pts | 0.9 | yes | yes |
| natural-str | street | 16384 | churn | ordered | hashed | 6 | 180 | 79.4 | 0.45× [0.44, 0.46] | -122.0% | [-127.2%, -116.8%] | 5.0 pts | 1.9 | yes | yes |
| natural-str | street | 16384 | build | ordered | baseline | 6 | 25.98 ms | 23.47 ms | 0.89× [0.88, 0.91] | -12.2% | [-14.1%, -10.3%] | 1.8 pts | 2.1 | yes | yes |
| natural-str | street | 16384 | build | ordered | btree-sets | 6 | 26.20 ms | 44.97 ms | 1.72× [1.70, 1.73] | +41.7% | [+41.3%, +42.1%] | 0.4 pts | 0.7 | yes | yes |
| natural-str | street | 16384 | build | ordered | hashed | 6 | 26.16 ms | 14.04 ms | 0.54× [0.53, 0.55] | -85.7% | [-88.0%, -83.5%] | 2.1 pts | 0.8 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | baseline | 8 | 267 | 227 | 0.86× [0.84, 0.87] | -16.8% | [-19.3%, -14.4%] | 2.8 pts | 1.7 | no | yes |
| natural-str | street | 212449 | valuesFor | ordered | btree-sets | 8 | 315 | 502 | 1.58× [1.55, 1.61] | +36.5% | [+35.3%, +37.7%] | 1.8 pts | 2.0 | yes | yes |
| natural-str | street | 212449 | valuesFor | ordered | hashed | 8 | 259 | 109 | 0.43× [0.41, 0.44] | -134.8% | [-142.6%, -127.0%] | 8.6 pts | 3.3 | yes | yes |
| natural-str | street | 212449 | valuesBetween | ordered | baseline | 8 | 10.3 µs | 6990 | 0.69× [0.66, 0.72] | -44.6% | [-50.9%, -38.2%] | 6.9 pts | 5.0 | no | yes |
| natural-str | street | 212449 | valuesBetween | ordered | btree-sets | 8 | 11.3 µs | 24.1 µs | 2.15× [2.08, 2.21] | +53.4% | [+51.9%, +54.9%] | 1.6 pts | 3.3 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | baseline | 8 | 37.2 µs | 22.3 µs | 0.61× [0.59, 0.63] | -64.3% | [-70.1%, -58.4%] | 5.7 pts | 1.6 | yes | yes |
| natural-str | street | 212449 | prefix | ordered | btree-sets | 8 | 39.6 µs | 80.5 µs | 2.07× [1.97, 2.18] | +51.8% | [+49.3%, +54.2%] | 2.3 pts | 1.3 | yes | yes |
| natural-str | street | 212449 | churn | ordered | baseline | 8 | 538 | 487 | 0.90× [0.88, 0.92] | -11.6% | [-14.2%, -9.1%] | 3.1 pts | 0.7 | no | yes |
| natural-str | street | 212449 | churn | ordered | btree-sets | 8 | 623 | 756 | 1.21× [1.18, 1.24] | +17.6% | [+15.5%, +19.6%] | 2.1 pts | 2.2 | no | yes |
| natural-str | street | 212449 | churn | ordered | hashed | 8 | 513 | 279 | 0.53× [0.51, 0.55] | -88.5% | [-96.7%, -80.4%] | 7.8 pts | 0.7 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural-str street n=4096 valuesFor: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs btree-sets: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesFor: ordered vs hashed: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs baseline: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 valuesBetween: ordered vs btree-sets: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=16384 prefix: ordered vs hashed: the A/A validations found a systematic difference of +3.07% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=16384 build: ordered vs baseline: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesFor: ordered vs hashed: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs baseline: the A/A validations found a systematic difference of -0.39% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- natural-str street n=212449 valuesBetween: ordered vs baseline: the processes scatter 5.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 valuesBetween: ordered vs btree-sets: the processes scatter 3.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- natural-str street n=212449 churn: ordered vs btree-sets: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
