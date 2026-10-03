| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | inflation | precise | resolved |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|---|
| unique-str | street | 4096 | valuesFor | ordered | ordered-lpage | 6 | 46.9 | 79.3 | 1.71× [1.68, 1.73] | +41.4% | [+40.5%, +42.2%] | 0.8 pts | 2.4 | yes | yes |
| unique-str | street | 4096 | valuesFor | ordered | ordered-lpage-zc | 6 | 46.6 | 64.3 | 1.38× [1.36, 1.41] | +27.7% | [+26.6%, +28.8%] | 1.1 pts | 2.3 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | ordered-lpage | 6 | 1784 | 1309 | 0.74× [0.72, 0.75] | -35.8% | [-39.0%, -32.7%] | 3.0 pts | 1.5 | yes | yes |
| unique-str | street | 4096 | valuesBetween | ordered | ordered-lpage-zc | 6 | 1774 | 928 | 0.52× [0.51, 0.52] | -93.4% | [-96.2%, -90.7%] | 2.6 pts | 1.2 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | ordered-lpage | 6 | 201 | 218 | 1.09× [1.08, 1.10] | +8.5% | [+7.6%, +9.4%] | 0.9 pts | 0.7 | yes | yes |
| unique-str | street | 4096 | prefix | ordered | ordered-lpage-zc | 6 | 200 | 171 | 0.85× [0.84, 0.86] | -18.0% | [-19.5%, -16.5%] | 1.4 pts | 1.3 | yes | yes |
| unique-str | street | 4096 | churn | ordered | ordered-lpage | 6 | 123 | 257 | 2.08× [2.02, 2.14] | +51.9% | [+50.6%, +53.3%] | 1.3 pts | 2.2 | yes | yes |
| unique-str | street | 4096 | churn | ordered | ordered-lpage-zc | 6 | 125 | 310 | 2.49× [2.45, 2.53] | +59.9% | [+59.2%, +60.5%] | 0.6 pts | 0.9 | yes | yes |
| unique-str | street | 4096 | build | ordered | ordered-lpage | 6 | 1.76 ms | 4.38 ms | 2.49× [2.43, 2.55] | +59.8% | [+58.8%, +60.8%] | 0.9 pts | 1.8 | yes | yes |
| unique-str | street | 4096 | build | ordered | ordered-lpage-zc | 6 | 1.74 ms | 5.09 ms | 2.91× [2.88, 2.94] | +65.7% | [+65.3%, +66.0%] | 0.4 pts | 1.1 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | ordered-lpage | 12 | 64.8 | 92.9 | 1.46× [1.41, 1.50] | +31.3% | [+29.2%, +33.5%] | 2.7 pts | 4.3 | yes | yes |
| unique-str | street | 16384 | valuesFor | ordered | ordered-lpage-zc | 12 | 65.2 | 77.2 | 1.20× [1.18, 1.22] | +16.9% | [+15.5%, +18.3%] | 1.9 pts | 2.6 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | ordered-lpage | 12 | 2158 | 1553 | 0.72× [0.72, 0.73] | -38.5% | [-39.8%, -37.2%] | 2.2 pts | 1.6 | yes | yes |
| unique-str | street | 16384 | valuesBetween | ordered | ordered-lpage-zc | 12 | 2164 | 1046 | 0.48× [0.48, 0.49] | -106.2% | [-107.4%, -104.9%] | 1.8 pts | 1.0 | yes | yes |
| unique-str | street | 16384 | prefix | ordered | ordered-lpage | 12 | 637 | 546 | 0.87× [0.85, 0.88] | -15.5% | [-17.7%, -13.4%] | 3.1 pts | 2.0 | no | yes |
| unique-str | street | 16384 | prefix | ordered | ordered-lpage-zc | 12 | 628 | 361 | 0.58× [0.57, 0.58] | -73.5% | [-74.5%, -72.6%] | 1.6 pts | 0.7 | yes | yes |
| unique-str | street | 16384 | churn | ordered | ordered-lpage | 12 | 159 | 294 | 1.84× [1.82, 1.87] | +45.7% | [+45.0%, +46.5%] | 1.1 pts | 1.5 | yes | yes |
| unique-str | street | 16384 | churn | ordered | ordered-lpage-zc | 12 | 170 | 353 | 2.07× [2.02, 2.13] | +51.8% | [+50.5%, +53.0%] | 1.6 pts | 2.0 | yes | yes |
| unique-str | street | 16384 | build | ordered | ordered-lpage | 12 | 8.64 ms | 18.61 ms | 2.15× [2.14, 2.16] | +53.5% | [+53.2%, +53.8%] | 0.3 pts | 0.7 | yes | yes |
| unique-str | street | 16384 | build | ordered | ordered-lpage-zc | 12 | 8.70 ms | 22.45 ms | 2.58× [2.56, 2.59] | +61.2% | [+61.0%, +61.4%] | 0.6 pts | 0.9 | yes | yes |
| unique-str | street | 212449 | valuesFor | ordered | ordered-lpage | 12 | 239 | 202 | 0.85× [0.82, 0.89] | -17.5% | [-22.3%, -12.8%] | 5.1 pts | 3.1 | no | yes |
| unique-str | street | 212449 | valuesFor | ordered | ordered-lpage-zc | 12 | 231 | 171 | 0.76× [0.72, 0.81] | -31.0% | [-39.2%, -22.8%] | 8.2 pts | 4.1 | no | yes |
| unique-str | street | 212449 | valuesBetween | ordered | ordered-lpage | 12 | 4585 | 2711 | 0.63× [0.59, 0.69] | -58.2% | [-70.7%, -45.6%] | 13.6 pts | 4.6 | no | yes |
| unique-str | street | 212449 | valuesBetween | ordered | ordered-lpage-zc | 12 | 3951 | 1788 | 0.45× [0.40, 0.50] | -124.6% | [-149.8%, -99.5%] | 24.9 pts | 5.1 | no | yes |
| unique-str | street | 212449 | prefix | ordered | ordered-lpage | 12 | 13.8 µs | 7962 | 0.59× [0.56, 0.63] | -68.3% | [-77.5%, -59.1%] | 12.9 pts | 3.6 | no | yes |
| unique-str | street | 212449 | prefix | ordered | ordered-lpage-zc | 12 | 13.1 µs | 5070 | 0.39× [0.37, 0.40] | -159.4% | [-166.8%, -152.0%] | 12.6 pts | 1.5 | yes | yes |
| unique-str | street | 212449 | churn | ordered | ordered-lpage | 12 | 529 | 646 | 1.23× [1.21, 1.25] | +18.9% | [+17.6%, +20.1%] | 2.0 pts | 1.8 | yes | yes |
| unique-str | street | 212449 | churn | ordered | ordered-lpage-zc | 12 | 501 | 679 | 1.34× [1.32, 1.37] | +25.6% | [+24.1%, +27.1%] | 2.1 pts | 2.4 | yes | yes |

Regime: serial: one process at a time, each with the machine to itself.

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the pooled difference (Estimate.Ratio, B/A); the bracket is its 95% interval across processes (rtcompare.CombineStaged: a t interval over the per-process differences, with the scatter of the run's first stage, which sized the run once, so that stopping early on a calm scatter cannot narrow it). Difference: rtcompare's relative difference, positive when A is faster. Inflation: spread between processes over the spread one process's interval implies. Precise: the interval is within the asked-for points or share of the difference. Resolved: the interval excludes zero and the difference clears the systematic bias the A/A validations found in every process. Churn and build come from rtcompare's workload package: churn is ns per insertion or deletion in a multimap in use, build ns per whole build; the streams delete and put back long-lived values too (-permchurn).

Warnings from pooling:

- unique-str street n=4096 valuesFor: ordered vs ordered-lpage: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 2.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 churn: ordered vs ordered-lpage: the processes scatter 2.2 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=4096 build: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of +0.41% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=16384 valuesFor: ordered vs ordered-lpage: the processes scatter 4.3 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 2.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=16384 build: ordered vs ordered-lpage-zc: the A/A validations found a systematic difference of -0.56% between identical code, the same in every process; the harness favours one position, and pooling cannot remove that
- unique-str street n=212449 valuesFor: ordered vs ordered-lpage: the processes scatter 3.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesFor: ordered vs ordered-lpage-zc: the processes scatter 4.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesBetween: ordered vs ordered-lpage: the processes scatter 4.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 valuesBetween: ordered vs ordered-lpage-zc: the processes scatter 5.1 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 prefix: ordered vs ordered-lpage: the processes scatter 3.6 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
- unique-str street n=212449 churn: ordered vs ordered-lpage-zc: the processes scatter 2.4 times as widely as one process's interval implies, so a single process's result for this comparison is not to be trusted on its own
