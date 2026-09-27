| values | keys | n | operation | A | B | processes | A ns/op | B ns/op | A speed vs B | difference | 95% across processes | sd between | ratio | precise |
|---|---|---:|---|---|---|---:|---:|---:|---|---:|---|---:|---:|---|
| multi | str | 4194304 | valuesFor | ordered | baseline | 3 | 790 | 619 | 0.78× [0.67, 0.92] | -27.5% | [-48.9%, -8.9%] | 8.1 pts | 3.4 | no |
| multi | str | 4194304 | valuesBetween | ordered | baseline | 3 | 16.0 µs | 14.2 µs | 0.89× [0.74, 1.16] | -12.4% | [-35.9%, +13.8%] | 10.0 pts | 3.2 | no |
| multi | str | 4194304 | prefix | ordered | baseline | 3 | 32.67 ms | 35.66 ms | 1.09× [0.99, 1.21] | +8.4% | [-1.0%, +17.2%] | 3.7 pts | 2.4 | no |
| multi | str | 4194304 | churn | ordered | baseline | 3 | 962 | 806 | 0.83× [0.81, 0.85] | -20.9% | [-24.2%, -17.4%] | 1.4 pts | 0.6 | no |
| multi | u64 | 4194304 | valuesFor | ordered | baseline | 3 | 282 | 243 | 0.82× [0.75, 0.91] | -21.4% | [-32.7%, -9.4%] | 4.7 pts | 1.0 | no |
| multi | u64 | 4194304 | valuesBetween | ordered | baseline | 3 | 10.9 µs | 9443 | 0.87× [0.77, 1.06] | -14.5% | [-29.6%, +5.4%] | 7.1 pts | 4.5 | no |
| multi | u64 | 4194304 | churn | ordered | baseline | 3 | 626 | 396 | 0.66× [0.59, 0.75] | -51.7% | [-69.2%, -33.4%] | 7.2 pts | 1.1 | no |
| multi | uuid | 4194304 | valuesFor | ordered | baseline | 3 | 809 | 538 | 0.69× [0.59, 0.87] | -44.6% | [-69.6%, -14.3%] | 11.1 pts | 3.1 | no |
| multi | uuid | 4194304 | valuesBetween | ordered | baseline | 3 | 14.9 µs | 12.4 µs | 0.88× [0.75, 1.08] | -14.2% | [-32.6%, +7.6%] | 8.1 pts | 3.0 | no |
| multi | uuid | 4194304 | prefix | ordered | baseline | 3 | 8604 | 7938 | 0.93× [0.86, 0.99] | -7.7% | [-15.6%, -1.0%] | 2.9 pts | 0.9 | no |
| multi | uuid | 4194304 | churn | ordered | baseline | 3 | 1118 | 768 | 0.80× [0.57, 1.16] | -24.5% | [-76.0%, +13.9%] | 18.1 pts | 6.6 | no |
| unique | str | 4194304 | valuesFor | ordered | baseline | 3 | 564 | 523 | 0.92× [0.75, 1.22] | -8.7% | [-33.4%, +18.0%] | 10.3 pts | 6.3 | no |
| unique | str | 4194304 | valuesBetween | ordered | baseline | 3 | 2613 | 9219 | 3.47× [2.95, 4.79] | +71.2% | [+66.1%, +79.1%] | 2.6 pts | 5.8 | yes |
| unique | str | 4194304 | prefix | ordered | baseline | 3 | 4.33 ms | 22.50 ms | 5.19× [4.44, 6.81] | +80.7% | [+77.5%, +85.3%] | 1.6 pts | 3.5 | yes |
| unique | str | 4194304 | churn | ordered | baseline | 3 | 812 | 763 | 0.94× [0.63, 1.78] | -6.4% | [-58.4%, +43.8%] | 20.6 pts | 12.3 | no |
| unique | u64 | 4194304 | valuesFor | ordered | baseline | 3 | 295 | 176 | 0.58× [0.54, 0.64] | -71.6% | [-86.9%, -55.6%] | 6.3 pts | 1.5 | no |
| unique | u64 | 4194304 | valuesBetween | ordered | baseline | 3 | 1334 | 4357 | 3.23× [3.05, 3.42] | +69.1% | [+67.2%, +70.8%] | 0.7 pts | 1.3 | yes |
| unique | u64 | 4194304 | churn | ordered | baseline | 3 | 627 | 378 | 0.50× [0.34, 2.66] | -98.1% | [-198.0%, +62.4%] | 52.4 pts | 14.1 | no |
| unique | uuid | 4194304 | valuesFor | ordered | baseline | 3 | 650 | 443 | 0.69× [0.65, 0.76] | -45.1% | [-54.9%, -32.0%] | 4.6 pts | 1.4 | no |
| unique | uuid | 4194304 | valuesBetween | ordered | baseline | 3 | 2529 | 7368 | 2.91× [2.74, 3.12] | +65.7% | [+63.4%, +67.9%] | 0.9 pts | 1.6 | yes |
| unique | uuid | 4194304 | prefix | ordered | baseline | 3 | 1228 | 4695 | 3.83× [3.49, 4.10] | +73.9% | [+71.3%, +75.6%] | 0.9 pts | 2.0 | yes |
| unique | uuid | 4194304 | churn | ordered | baseline | 3 | 879 | 593 | 0.67× [0.44, 1.75] | -48.3% | [-126.2%, +42.8%] | 34.0 pts | 19.7 | no |

A speed vs B: how many operations A completes in the time B needs for one (2.00× = twice as fast, 0.50× = half as fast), from the median difference; the bracket is its 95% interval across processes. Difference: rtcompare's relative difference, positive when A is faster. Ratio: spread between processes over the standard error one process reports.
