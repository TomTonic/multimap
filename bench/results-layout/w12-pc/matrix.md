
natural: ordered speed vs the code before W1+W2 (df110e6) / vs btree-sets (>1 = ordered faster; * = not precise)

| keys | op | 4,096 | 16,384 | 65,536 |
|---|---|---|---|---|
| street | valuesFor | 0.98 / 2.01 | 0.99 / 2.22 | 1.00 / 2.72 |
| street | valuesBetween | 1.04* / 4.90 | 1.02 / 4.22 | 1.02 / 5.08 |
| street | churn | 1.00 / 1.31 | 1.02 / 1.53 | 1.00* / 1.65 |
| street | build | 1.01 / 1.25 | 1.01 / 1.46 | 0.99 / 1.72 |
| dirs | valuesFor | 1.00 / 1.69 | 0.99* / 1.78 | 1.01* / 2.22 |
| dirs | valuesBetween | 1.05 / 3.70 | 1.04 / 3.57 | 1.02 / 4.67 |
| dirs | churn | 0.99 / 1.22* | 1.00 / 1.49* | 1.00* / 1.52 |
| dirs | build | 1.01 / 1.16 | 1.01 / 1.36 | 1.00 / 1.61 |
| links | valuesFor | 1.00 / 1.98 | 1.00 / 1.99 | 0.98 / 1.94 |
| links | valuesBetween | 1.09 / 3.37 | 1.04 / 3.00 | 1.00* / 3.09 |
| links | churn | 0.99 / 1.20 | 0.99 / 1.48 | 0.96* / 1.31 |
| links | build | 1.00 / 1.13 | 1.00 / 1.38 | 1.00 / 1.38 |
| url | valuesFor | 1.00 / 1.83 | 0.99 / 1.94 | 1.00* / 2.06 |
| url | valuesBetween | 1.04 / 3.27 | 1.03 / 2.91 | 1.02 / 3.45 |
| url | churn | 1.01 / 1.20 | 1.00 / 1.50 | 1.00 / 1.33 |
| url | build | 1.01 / 1.15 | 0.99 / 1.34 | 0.99 / 1.46 |

single-value: ordered speed vs the code before W1+W2 (df110e6) / vs btree-map (>1 = ordered faster; * = not precise)

| keys | op | 4,096 | 16,384 | 65,536 |
|---|---|---|---|---|
| street | valuesFor | 0.98* / 1.56 | 1.00 / 1.76 | 0.99 / 2.00 |
| street | valuesBetween | 1.02 / 0.83 | 1.02 / 0.74 | 1.02 / 0.84* |
| street | churn | 1.03 / 1.01 | 1.02 / 1.16 | 1.00 / 1.27* |
| street | build | 1.00 / 0.92 | 1.00 / 1.04 | 0.99 / 1.13 |
| dirs | valuesFor | 0.99 / 1.33 | 0.99 / 1.51 | 0.98* / 1.57 |
| dirs | valuesBetween | 1.03 / 0.57 | 1.02 / 0.65 | 1.01 / 0.77* |
| dirs | churn | 1.00 / 0.91 | 1.01 / 1.01* | 1.01 / 1.22 |
| dirs | build | 0.99 / 0.78 | 1.00 / 0.91 | 0.99 / 1.05* |
| links | valuesFor | 0.99 / 1.83 | 0.99 / 1.97 | 0.99* / 2.20 |
| links | valuesBetween | 1.01 / 0.80 | 1.02 / 0.84 | 1.01 / 0.86 |
| links | churn | 0.95 / 1.01 | 0.95 / 1.17 | 0.96 / 1.30* |
| links | build | 0.99 / 0.96* | 0.99 / 1.11 | 0.99 / 1.14 |
| url | valuesFor | 0.98* / 1.39 | 0.99 / 1.59 | 0.99 / 1.29* |
| url | valuesBetween | 1.00 / 0.49 | 1.01* / 0.52 | 1.01* / 0.62* |
| url | churn | 1.00* / 0.87 | 1.00 / 0.96* | 1.00 / 1.06* |
| url | build | 0.99 / 0.79 | 0.98 / 0.85 | 0.98 / 0.94 |
