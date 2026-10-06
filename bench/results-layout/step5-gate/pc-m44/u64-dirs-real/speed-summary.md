| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| natural | dirs | 4096 | valuesFor | ordered | baseline | 6 | 101 | 86.0 | 0.85× [0.84, 0.86] | -17.6% | [-18.7%, -16.4%] | 1.1 pts | 0.8 | yes | yes |
| natural | dirs | 4096 | valuesFor | ordered | btree-sets | 6 | 103 | 160 | 1.57× [1.55, 1.60] | +36.4% | [+35.4%, +37.4%] | 1.0 pts | 1.1 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | baseline | 6 | 2610 | 3032 | 1.18× [1.15, 1.20] | +15.0% | [+13.4%, +16.6%] | 1.5 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | valuesBetween | ordered | btree-sets | 6 | 2598 | 5413 | 2.07× [2.03, 2.12] | +51.8% | [+50.9%, +52.7%] | 0.9 pts | 1.5 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | baseline | 6 | 2918 | 3356 | 1.13× [1.11, 1.15] | +11.4% | [+10.0%, +12.8%] | 1.4 pts | 0.7 | yes | yes |
| natural | dirs | 4096 | prefix | ordered | btree-sets | 6 | 2927 | 7400 | 2.47× [2.39, 2.55] | +59.5% | [+58.2%, +60.8%] | 1.3 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | churn | ordered | baseline | 6 | 232 | 153 | 0.66× [0.64, 0.67] | -52.4% | [-55.2%, -49.7%] | 2.6 pts | 1.3 | yes | yes |
| natural | dirs | 4096 | churn | ordered | btree-sets | 6 | 230 | 226 | 0.99× [0.97, 1.00] | -1.4% | [-3.2%, +0.5%] | 1.8 pts | 1.3 | yes | no |
| natural | dirs | 4096 | build | ordered | baseline | 6 | 10.93 ms | 7.27 ms | 0.66× [0.66, 0.67] | -50.6% | [-52.0%, -49.3%] | 1.3 pts | 1.2 | yes | yes |
| natural | dirs | 4096 | build | ordered | btree-sets | 6 | 10.96 ms | 10.70 ms | 0.97× [0.96, 0.99] | -2.6% | [-3.8%, -1.3%] | 1.2 pts | 1.4 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | baseline | 6 | 127 | 116 | 0.91× [0.90, 0.92] | -10.0% | [-11.7%, -8.2%] | 1.7 pts | 1.6 | yes | yes |
| natural | dirs | 16384 | valuesFor | ordered | btree-sets | 6 | 129 | 222 | 1.72× [1.71, 1.74] | +42.0% | [+41.6%, +42.5%] | 0.4 pts | 0.6 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | baseline | 6 | 3001 | 3752 | 1.24× [1.23, 1.26] | +19.6% | [+18.8%, +20.4%] | 0.8 pts | 1.2 | yes | yes |
| natural | dirs | 16384 | valuesBetween | ordered | btree-sets | 6 | 3009 | 6153 | 2.04× [2.03, 2.05] | +51.0% | [+50.7%, +51.3%] | 0.3 pts | 0.9 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | baseline | 6 | 11.7 µs | 17.0 µs | 1.42× [1.37, 1.48] | +29.7% | [+26.9%, +32.6%] | 2.7 pts | 0.5 | yes | yes |
| natural | dirs | 16384 | prefix | ordered | btree-sets | 6 | 10.7 µs | 31.3 µs | 2.89× [2.82, 2.96] | +65.4% | [+64.6%, +66.2%] | 0.8 pts | 0.3 | yes | yes |
| natural | dirs | 16384 | churn | ordered | baseline | 6 | 272 | 206 | 0.76× [0.75, 0.77] | -31.3% | [-33.3%, -29.3%] | 1.9 pts | 1.5 | yes | yes |
| natural | dirs | 16384 | churn | ordered | btree-sets | 6 | 274 | 335 | 1.21× [1.19, 1.24] | +17.7% | [+16.2%, +19.2%] | 1.4 pts | 1.6 | yes | yes |
| natural | dirs | 16384 | build | ordered | baseline | 6 | 52.50 ms | 38.54 ms | 0.74× [0.73, 0.75] | -35.5% | [-37.2%, -33.9%] | 1.6 pts | 1.8 | yes | yes |
| natural | dirs | 16384 | build | ordered | btree-sets | 6 | 53.20 ms | 59.92 ms | 1.13× [1.12, 1.14] | +11.4% | [+10.7%, +12.1%] | 0.7 pts | 1.0 | yes | yes |
| natural | dirs | 86215 | valuesFor | ordered | baseline | 8 | 179 | 181 | 1.02× [1.00, 1.04] | +1.6% | [-0.4%, +3.7%] | 1.9 pts | 1.5 | no | no |
| natural | dirs | 86215 | valuesFor | ordered | btree-sets | 8 | 198 | 416 | 2.10× [2.05, 2.15] | +52.3% | [+51.2%, +53.5%] | 1.2 pts | 1.5 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | baseline | 8 | 3429 | 4447 | 1.30× [1.29, 1.31] | +23.1% | [+22.5%, +23.7%] | 1.0 pts | 1.2 | yes | yes |
| natural | dirs | 86215 | valuesBetween | ordered | btree-sets | 8 | 4184 | 12.0 µs | 2.88× [2.78, 2.98] | +65.2% | [+64.1%, +66.4%] | 1.1 pts | 1.5 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | baseline | 8 | 83.7 µs | 118.4 µs | 1.42× [1.40, 1.44] | +29.6% | [+28.6%, +30.6%] | 1.0 pts | 1.2 | yes | yes |
| natural | dirs | 86215 | prefix | ordered | btree-sets | 8 | 87.6 µs | 243.7 µs | 2.80× [2.71, 2.90] | +64.3% | [+63.1%, +65.5%] | 1.3 pts | 0.9 | yes | yes |
| natural | dirs | 86215 | churn | ordered | baseline | 8 | 427 | 420 | 0.95× [0.93, 0.97] | -4.9% | [-7.0%, -2.7%] | 2.6 pts | 0.7 | no | yes |
| natural | dirs | 86215 | churn | ordered | btree-sets | 8 | 508 | 662 | 1.29× [1.29, 1.30] | +22.8% | [+22.3%, +23.3%] | 0.5 pts | 0.5 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- natural dirs n=4096 churn: ordered vs btree-sets: the pooled interval [-3.22%, 0.48%] includes zero
- natural dirs n=86215 valuesFor: ordered vs baseline: the pooled interval [-0.40%, 3.68%] includes zero
