| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | dirs | 4096 | valuesFor | ordered | btree-map | 8 | 118 | 150 | 1.26× [1.23, 1.28] | +20.3% | [+18.6%, +22.0%] | 2.0 pts | 1.6 | yes | yes |
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage | 8 | 118 | 193 | 1.64× [1.63, 1.65] | +39.1% | [+38.8%, +39.4%] | 0.4 pts | 0.8 | yes | yes |
| unique-str | dirs | 4096 | valuesFor | ordered | ordered-lpage-zc | 8 | 117 | 173 | 1.48× [1.47, 1.49] | +32.6% | [+32.2%, +33.0%] | 0.5 pts | 0.9 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | btree-map | 8 | 2823 | 2031 | 0.72× [0.71, 0.73] | -38.8% | [-40.0%, -37.7%] | 1.4 pts | 0.7 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage | 8 | 2834 | 2117 | 0.75× [0.75, 0.76] | -32.6% | [-33.9%, -31.4%] | 1.5 pts | 0.8 | yes | yes |
| unique-str | dirs | 4096 | valuesBetween | ordered | ordered-lpage-zc | 8 | 2800 | 1602 | 0.57× [0.57, 0.58] | -74.3% | [-76.3%, -72.3%] | 2.4 pts | 0.8 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | btree-map | 8 | 3378 | 2547 | 0.75× [0.74, 0.77] | -32.5% | [-34.8%, -30.2%] | 2.8 pts | 0.5 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage | 8 | 3407 | 2467 | 0.73× [0.72, 0.75] | -36.3% | [-38.9%, -33.6%] | 3.2 pts | 0.6 | yes | yes |
| unique-str | dirs | 4096 | prefix | ordered | ordered-lpage-zc | 8 | 3351 | 1780 | 0.54× [0.53, 0.54] | -86.8% | [-88.0%, -85.6%] | 1.4 pts | 0.3 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | btree-map | 8 | 260 | 226 | 0.87× [0.87, 0.87] | -14.9% | [-15.5%, -14.3%] | 0.8 pts | 0.6 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage | 8 | 261 | 352 | 1.34× [1.34, 1.35] | +25.5% | [+25.2%, +25.8%] | 0.3 pts | 0.8 | yes | yes |
| unique-str | dirs | 4096 | churn | ordered | ordered-lpage-zc | 8 | 264 | 414 | 1.56× [1.54, 1.57] | +35.9% | [+35.3%, +36.5%] | 0.7 pts | 1.4 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | btree-map | 8 | 3.55 ms | 3.25 ms | 0.91× [0.91, 0.92] | -9.6% | [-10.1%, -9.1%] | 0.6 pts | 1.8 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage | 8 | 3.57 ms | 6.06 ms | 1.70× [1.69, 1.71] | +41.1% | [+40.8%, +41.5%] | 0.4 pts | 1.9 | yes | yes |
| unique-str | dirs | 4096 | build | ordered | ordered-lpage-zc | 8 | 3.57 ms | 6.81 ms | 1.92× [1.91, 1.93] | +47.9% | [+47.7%, +48.1%] | 0.2 pts | 0.7 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | btree-map | 10 | 151 | 215 | 1.39× [1.35, 1.42] | +28.0% | [+26.2%, +29.7%] | 2.2 pts | 1.7 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage | 10 | 151 | 234 | 1.55× [1.53, 1.57] | +35.6% | [+34.8%, +36.4%] | 1.0 pts | 2.4 | yes | yes |
| unique-str | dirs | 16384 | valuesFor | ordered | ordered-lpage-zc | 10 | 149 | 213 | 1.41× [1.38, 1.45] | +29.3% | [+27.5%, +31.2%] | 2.2 pts | 4.9 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | btree-map | 10 | 3137 | 2273 | 0.73× [0.72, 0.75] | -36.6% | [-39.3%, -33.9%] | 3.2 pts | 3.0 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage | 10 | 3142 | 2273 | 0.73× [0.72, 0.74] | -36.7% | [-38.8%, -34.6%] | 2.7 pts | 2.8 | yes | yes |
| unique-str | dirs | 16384 | valuesBetween | ordered | ordered-lpage-zc | 10 | 3105 | 1723 | 0.56× [0.55, 0.56] | -80.1% | [-81.0%, -79.2%] | 1.1 pts | 0.6 | yes | yes |
| unique-str | dirs | 16384 | prefix | ordered | btree-map | 10 | 14.3 µs | 10.2 µs | 0.71× [0.71, 0.72] | -40.1% | [-41.7%, -38.5%] | 2.3 pts | 3.2 | yes | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage | 10 | 15.1 µs | 9564 | 0.64× [0.63, 0.64] | -56.7% | [-57.9%, -55.6%] | 1.7 pts | 1.2 | yes | yes |
| unique-str | dirs | 16384 | prefix | ordered | ordered-lpage-zc | 10 | 14.6 µs | 6642 | 0.45× [0.45, 0.45] | -121.9% | [-122.7%, -121.1%] | 1.1 pts | 0.7 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | btree-map | 10 | 349 | 307 | 0.89× [0.88, 0.90] | -12.1% | [-13.8%, -10.5%] | 1.9 pts | 1.8 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage | 10 | 365 | 426 | 1.19× [1.16, 1.21] | +15.7% | [+13.7%, +17.6%] | 2.8 pts | 3.6 | yes | yes |
| unique-str | dirs | 16384 | churn | ordered | ordered-lpage-zc | 10 | 378 | 510 | 1.33× [1.30, 1.38] | +25.1% | [+22.9%, +27.3%] | 2.8 pts | 3.8 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | btree-map | 10 | 17.58 ms | 16.64 ms | 0.95× [0.94, 0.96] | -5.1% | [-6.3%, -3.8%] | 1.5 pts | 1.9 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage | 10 | 17.75 ms | 27.69 ms | 1.56× [1.54, 1.57] | +35.8% | [+35.2%, +36.3%] | 1.1 pts | 2.1 | yes | yes |
| unique-str | dirs | 16384 | build | ordered | ordered-lpage-zc | 10 | 17.83 ms | 31.82 ms | 1.80× [1.79, 1.81] | +44.4% | [+44.2%, +44.7%] | 0.5 pts | 0.8 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | btree-map | 24 | 291 | 341 | 1.18× [1.16, 1.21] | +15.6% | [+13.9%, +17.2%] | 3.8 pts | 1.2 | yes | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage | 24 | 253 | 306 | 1.17× [1.12, 1.23] | +14.8% | [+10.8%, +18.9%] | 6.5 pts | 4.1 | no | yes |
| unique-str | dirs | 86215 | valuesFor | ordered | ordered-lpage-zc | 24 | 240 | 280 | 1.13× [1.08, 1.19] | +11.6% | [+7.6%, +15.6%] | 6.2 pts | 4.7 | no | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | btree-map | 24 | 4259 | 3205 | 0.75× [0.74, 0.76] | -33.5% | [-35.9%, -31.1%] | 3.6 pts | 1.1 | yes | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage | 24 | 4113 | 2597 | 0.63× [0.62, 0.64] | -58.3% | [-60.3%, -56.2%] | 2.8 pts | 0.9 | yes | yes |
| unique-str | dirs | 86215 | valuesBetween | ordered | ordered-lpage-zc | 24 | 3696 | 1930 | 0.52× [0.51, 0.53] | -92.0% | [-94.5%, -89.5%] | 5.3 pts | 1.5 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | btree-map | 24 | 84.2 µs | 56.4 µs | 0.68× [0.67, 0.69] | -48.0% | [-50.3%, -45.7%] | 3.8 pts | 1.0 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage | 24 | 88.3 µs | 48.9 µs | 0.57× [0.56, 0.57] | -77.0% | [-79.6%, -74.4%] | 5.3 pts | 1.1 | yes | yes |
| unique-str | dirs | 86215 | prefix | ordered | ordered-lpage-zc | 24 | 82.9 µs | 33.7 µs | 0.41× [0.41, 0.41] | -145.5% | [-146.6%, -144.4%] | 2.0 pts | 1.4 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | btree-map | 24 | 612 | 590 | 0.97× [0.96, 0.97] | -3.5% | [-4.4%, -2.7%] | 1.6 pts | 2.7 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage | 24 | 546 | 577 | 1.03× [1.02, 1.04] | +2.8% | [+1.6%, +4.1%] | 2.2 pts | 0.7 | yes | yes |
| unique-str | dirs | 86215 | churn | ordered | ordered-lpage-zc | 24 | 585 | 665 | 1.11× [1.09, 1.12] | +9.8% | [+8.5%, +11.1%] | 2.6 pts | 1.1 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str dirs n=4096 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.55% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 4.9 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 valuesBetween: ordered vs btree-map: the processes scatter 3.0 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 valuesBetween: ordered vs ordered-lpage: the processes scatter 2.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 prefix: ordered vs btree-map: the processes scatter 3.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 churn: ordered vs btree-map: the A/A validations found a systematic difference of -0.25% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=16384 churn: ordered vs ordered-lpage: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 churn: ordered vs ordered-lpage-zc: the processes scatter 3.8 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=16384 build: ordered vs ordered-lpage: the processes scatter 2.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesFor: ordered vs btree-map: the A/A validations found a systematic difference of +1.08% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the A/A validations found a systematic difference of +0.71% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of +0.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 4.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 valuesFor: ordered vs ordered-lpage-zc: 21 processes resolved A as faster and 2 as slower; each was confident, and the disagreement between them is the layout of memory, not the code
- unique-str dirs n=86215 prefix: ordered vs btree-map: the A/A validations found a systematic difference of +1.47% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 prefix: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of +2.30% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str dirs n=86215 churn: ordered vs btree-map: the processes scatter 2.7 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str dirs n=86215 churn: ordered vs ordered-lpage: the A/A validations found a systematic difference of -0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
