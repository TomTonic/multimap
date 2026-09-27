| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | str | 4194304 | valuesFor | ordered | baseline | 3 | 829 | 632 | 0.78× [0.68, 0.95] | -27.8% | [-46.4%, -5.1%] | 8.3 pts | 7.1 | no |
| multi | str | 4194304 | valuesBetween | ordered | baseline | 3 | 17.8 µs | 14.4 µs | 0.81× [0.64, 1.08] | -23.1% | [-56.9%, +7.2%] | 12.9 pts | 12.3 | no |
| multi | str | 4194304 | prefix | ordered | baseline | 3 | 39.09 ms | 35.53 ms | 0.90× [0.86, 0.95] | -11.5% | [-15.8%, -5.4%] | 2.1 pts | 1.5 | no |
| multi | str | 4194304 | churn | ordered | baseline | 3 | 1225 | 814 | 0.66× [0.51, 1.00] | -51.2% | [-95.7%, -0.5%] | 19.2 pts | 10.1 | no |
| multi | u64 | 4194304 | valuesFor | ordered | baseline | 3 | 387 | 235 | 0.60× [0.52, 0.71] | -65.6% | [-92.4%, -41.4%] | 10.3 pts | 2.5 | no |
| multi | u64 | 4194304 | valuesBetween | ordered | baseline | 3 | 10.0 µs | 9674 | 0.95× [0.75, 1.16] | -4.7% | [-32.7%, +13.7%] | 9.3 pts | 6.4 | no |
| multi | u64 | 4194304 | churn | ordered | baseline | 3 | 818 | 480 | 0.59× [0.41, 0.82] | -70.5% | [-143.9%, -22.1%] | 24.5 pts | 6.2 | no |
| multi | uuid | 4194304 | valuesFor | ordered | baseline | 3 | 870 | 551 | 0.63× [0.53, 0.83] | -58.0% | [-87.1%, -20.9%] | 13.3 pts | 7.0 | no |
| multi | uuid | 4194304 | valuesBetween | ordered | baseline | 3 | 17.9 µs | 12.6 µs | 0.69× [0.57, 0.95] | -44.1% | [-75.2%, -5.5%] | 14.0 pts | 7.1 | no |
| multi | uuid | 4194304 | prefix | ordered | baseline | 3 | 11.5 µs | 8194 | 0.71× [0.62, 0.85] | -40.8% | [-60.7%, -17.3%] | 8.7 pts | 4.0 | no |
| multi | uuid | 4194304 | churn | ordered | baseline | 3 | 1361 | 817 | 0.60× [0.52, 0.71] | -66.5% | [-93.8%, -40.9%] | 10.6 pts | 3.3 | no |
| unique | str | 4194304 | valuesFor | ordered | baseline | 3 | 522 | 540 | 1.03× [0.80, 1.32] | +3.3% | [-25.7%, +24.2%] | 10.0 pts | 12.3 | no |
| unique | str | 4194304 | valuesBetween | ordered | baseline | 3 | 2447 | 9231 | 3.82× [3.20, 4.43] | +73.8% | [+68.7%, +77.4%] | 1.7 pts | 3.0 | yes |
| unique | str | 4194304 | prefix | ordered | baseline | 3 | 4.08 ms | 22.50 ms | 5.65× [5.12, 6.18] | +82.3% | [+80.5%, +83.8%] | 0.7 pts | 3.0 | yes |
| unique | str | 4194304 | churn | ordered | baseline | 3 | 819 | 806 | 0.95× [0.67, 1.63] | -5.7% | [-49.5%, +38.7%] | 17.7 pts | 7.0 | no |
| unique | u64 | 4194304 | valuesFor | ordered | baseline | 3 | 303 | 178 | 0.59× [0.53, 0.67] | -68.6% | [-89.9%, -49.1%] | 8.2 pts | 1.9 | no |
| unique | u64 | 4194304 | valuesBetween | ordered | baseline | 3 | 1259 | 4273 | 3.39× [2.92, 3.84] | +70.5% | [+65.8%, +74.0%] | 1.6 pts | 2.9 | yes |
| unique | u64 | 4194304 | churn | ordered | baseline | 3 | 670 | 417 | 0.58× [0.41, 1.27] | -73.5% | [-145.9%, +21.0%] | 33.6 pts | 8.2 | no |
| unique | uuid | 4194304 | valuesFor | ordered | baseline | 3 | 571 | 440 | 0.77× [0.66, 0.89] | -29.6% | [-51.8%, -12.8%] | 7.9 pts | 5.8 | no |
| unique | uuid | 4194304 | valuesBetween | ordered | baseline | 3 | 2379 | 7430 | 3.14× [2.60, 3.69] | +68.2% | [+61.5%, +72.9%] | 2.3 pts | 4.2 | yes |
| unique | uuid | 4194304 | prefix | ordered | baseline | 3 | 1138 | 4641 | 4.07× [3.42, 4.81] | +75.5% | [+70.7%, +79.2%] | 1.7 pts | 4.6 | yes |
| unique | uuid | 4194304 | churn | ordered | baseline | 3 | 830 | 669 | 0.86× [0.53, 1.42] | -16.6% | [-89.3%, +29.6%] | 23.9 pts | 9.3 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
