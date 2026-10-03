| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage | 6 | 75.6 | 115 | 1.51× [1.47, 1.55] | +33.6% | [+31.8%, +35.4%] | 1.7 pts | 2.8 | yes | yes |
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-zc | 6 | 75.4 | 94.7 | 1.27× [1.25, 1.28] | +21.0% | [+20.2%, +21.8%] | 0.7 pts | 1.0 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage | 6 | 2038 | 1800 | 0.89× [0.88, 0.90] | -11.9% | [-13.3%, -10.6%] | 1.3 pts | 0.6 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-zc | 6 | 2015 | 1157 | 0.58× [0.57, 0.58] | -73.7% | [-76.1%, -71.4%] | 2.2 pts | 0.9 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage | 6 | 2276 | 2170 | 0.95× [0.94, 0.97] | -5.2% | [-6.8%, -3.5%] | 1.6 pts | 1.1 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage-zc | 6 | 2261 | 1356 | 0.61× [0.60, 0.62] | -63.8% | [-65.8%, -61.9%] | 1.8 pts | 0.8 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage | 6 | 201 | 254 | 1.25× [1.22, 1.28] | +19.8% | [+17.8%, +21.8%] | 1.9 pts | 2.0 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage-zc | 6 | 205 | 323 | 1.58× [1.55, 1.62] | +36.9% | [+35.6%, +38.1%] | 1.2 pts | 1.1 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage | 6 | 2.85 ms | 4.81 ms | 1.69× [1.66, 1.71] | +40.7% | [+39.9%, +41.5%] | 0.8 pts | 1.5 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage-zc | 6 | 2.78 ms | 5.60 ms | 1.99× [1.98, 2.00] | +49.7% | [+49.4%, +50.1%] | 0.3 pts | 0.7 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage | 12 | 107 | 142 | 1.33× [1.30, 1.35] | +24.7% | [+23.4%, +26.0%] | 2.0 pts | 3.0 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-zc | 12 | 107 | 122 | 1.14× [1.12, 1.16] | +12.2% | [+10.7%, +13.7%] | 1.6 pts | 1.9 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage | 12 | 2601 | 2266 | 0.87× [0.86, 0.88] | -15.1% | [-16.3%, -13.9%] | 1.3 pts | 1.1 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-zc | 12 | 2588 | 1356 | 0.52× [0.52, 0.53] | -90.8% | [-91.3%, -90.3%] | 1.6 pts | 1.1 | yes | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage | 12 | 10.8 µs | 9168 | 0.86× [0.83, 0.88] | -16.6% | [-20.1%, -13.1%] | 5.1 pts | 1.0 | no | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage-zc | 12 | 10.4 µs | 4981 | 0.48× [0.47, 0.48] | -109.0% | [-111.5%, -106.5%] | 5.4 pts | 0.5 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage | 12 | 288 | 339 | 1.17× [1.16, 1.19] | +14.7% | [+13.5%, +15.8%] | 1.2 pts | 1.3 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage-zc | 12 | 306 | 463 | 1.51× [1.49, 1.53] | +33.8% | [+33.0%, +34.6%] | 1.0 pts | 0.9 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage | 12 | 14.16 ms | 22.61 ms | 1.58× [1.55, 1.60] | +36.5% | [+35.6%, +37.5%] | 1.0 pts | 1.1 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage-zc | 12 | 15.02 ms | 29.19 ms | 1.95× [1.93, 1.97] | +48.7% | [+48.1%, +49.3%] | 1.0 pts | 0.7 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage | 12 | 233 | 226 | 0.99× [0.96, 1.03] | -0.8% | [-4.4%, +2.8%] | 3.8 pts | 2.3 | no | no |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-zc | 12 | 209 | 181 | 0.91× [0.88, 0.93] | -10.3% | [-13.0%, -7.5%] | 6.3 pts | 3.1 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage | 12 | 3207 | 2680 | 0.83× [0.81, 0.86] | -19.8% | [-23.4%, -16.2%] | 4.3 pts | 3.1 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-zc | 12 | 3157 | 1732 | 0.54× [0.53, 0.54] | -85.5% | [-87.3%, -83.7%] | 2.8 pts | 1.0 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage | 12 | 82.1 µs | 61.4 µs | 0.77× [0.76, 0.78] | -30.4% | [-32.1%, -28.8%] | 2.4 pts | 0.8 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage-zc | 12 | 77.3 µs | 32.8 µs | 0.43× [0.43, 0.44] | -131.1% | [-132.6%, -129.5%] | 2.1 pts | 0.9 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage | 12 | 539 | 586 | 1.10× [1.08, 1.12] | +8.9% | [+7.1%, +10.7%] | 1.9 pts | 1.5 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage-zc | 12 | 554 | 690 | 1.27× [1.25, 1.28] | +21.0% | [+19.9%, +22.2%] | 1.2 pts | 1.2 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str dirs n=4096 valuesFor: ordered vs ordered-lpage: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=4096 valuesBetween: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of +0.49% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=4096 churn: ordered vs ordered-lpage: the processes scatter 2.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the pooled interval [-4.38%, 2.77%] includes zero
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage: 1 processes resolved A as faster and 4 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesBetween: ordered vs ordered-lpage: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 prefix: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -2.36% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
