| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | dirs | 4096 | valuesFor | ordered | btree-map | 6 | 77.1 | 111 | 1.45× [1.41, 1.48] | +30.8% | [+29.1%, +32.5%] | 1.6 pts | 1.7 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | btree-map | 6 | 1996 | 547 | 0.27× [0.27, 0.28] | -268.2% | [-276.5%, -260.0%] | 7.9 pts | 2.3 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | btree-map | 6 | 2238 | 647 | 0.29× [0.29, 0.29] | -243.7% | [-248.3%, -239.2%] | 4.4 pts | 1.5 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | btree-map | 6 | 201 | 184 | 0.91× [0.90, 0.93] | -9.5% | [-11.4%, -7.6%] | 1.8 pts | 1.2 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | btree-map | 6 | 2.76 ms | 2.67 ms | 0.97× [0.96, 0.97] | -3.4% | [-3.7%, -3.0%] | 0.4 pts | 0.4 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | btree-map | 6 | 107 | 162 | 1.53× [1.50, 1.57] | +34.7% | [+33.2%, +36.2%] | 1.4 pts | 1.9 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | btree-map | 6 | 2557 | 740 | 0.29× [0.28, 0.29] | -248.0% | [-254.3%, -241.6%] | 6.0 pts | 3.1 | yes | yes |
| unique-str | dirs | 16384 | prefix | ordered | btree-map | 6 | 9747 | 2533 | 0.26× [0.25, 0.26] | -288.6% | [-299.8%, -277.5%] | 10.6 pts | 0.9 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | btree-map | 6 | 257 | 254 | 0.98× [0.97, 1.00] | -1.5% | [-2.8%, -0.3%] | 1.2 pts | 1.4 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | btree-map | 6 | 13.68 ms | 14.16 ms | 1.04× [1.03, 1.05] | +3.5% | [+2.5%, +4.5%] | 0.9 pts | 1.1 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | btree-map | 12 | 204 | 245 | 1.24× [1.21, 1.28] | +19.4% | [+17.1%, +21.7%] | 3.6 pts | 2.3 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | btree-map | 12 | 2999 | 1102 | 0.37× [0.35, 0.38] | -172.2% | [-183.2%, -161.2%] | 11.4 pts | 2.0 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | btree-map | 12 | 72.9 µs | 18.3 µs | 0.24× [0.23, 0.24] | -318.5% | [-328.5%, -308.4%] | 16.7 pts | 1.2 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | btree-map | 12 | 498 | 461 | 0.95× [0.93, 0.96] | -5.6% | [-7.5%, -3.7%] | 2.0 pts | 1.3 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str dirs n=4096 valuesBetween: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.63% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesFor: ordered vs btree-map: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesBetween: ordered vs btree-map: the A/A validations found a systematic difference of -1.37% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
