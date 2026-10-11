
natural: ordered speed vs the code before W1+W2 (df110e6) / vs btree-sets (>1 = ordered faster; * = not precise)

| keys | op | 4,096 | 16,384 | 65,536 |
|---|---|---|---|---|
| street | valuesFor | 1.01 / 1.91 | 1.00 / 2.01 | 1.00 / 2.41 |
| street | valuesBetween | 1.04* / 4.00 | 1.01 / 3.44 | 1.01* / 4.88 |
| street | churn | 1.02 / 1.35 | 1.02 / 1.44 | 0.94* / 1.51* |
| street | build | 1.01 / 1.35 | 1.01 / 1.45 | 0.99 / 1.64 |
| dirs | valuesFor | 1.01 / 1.48 | 1.01 / 1.55 | 1.00 / 1.99 |
| dirs | valuesBetween | 1.02* / 3.02 | 1.00 / 3.08 | 1.00 / 4.10 |
| dirs | churn | 1.01 / 1.14 | 1.01 / 1.26 | 0.91* / 1.39 |
| dirs | build | 1.00 / 1.12 | 1.00 / 1.22 | 1.00 / 1.46 |
| links | valuesFor | 0.99 / 1.83 | 1.00 / 1.85 | 0.99 / 1.90 |
| links | valuesBetween | 1.01 / 2.83 | 1.00 / 2.89 | 1.00 / 3.01 |
| links | churn | 1.00 / 1.29 | 0.91* / 1.43 | 0.83* / 1.36 |
| links | build | 1.00 / 1.19 | 1.00 / 1.37 | 1.00 / 1.42 |
| url | valuesFor | 1.01 / 1.61 | 1.01 / 1.69 | 1.00 / 1.94 |
| url | valuesBetween | 1.01 / 2.56 | 1.01 / 2.53 | 1.01 / 3.10 |
| url | churn | 1.01 / 1.12 | 0.99* / 1.23 | 0.87* / 1.26* |
| url | build | 1.01 / 1.07 | 1.00 / 1.20* | 1.00 / 1.38 |

single-value: ordered speed vs the code before W1+W2 (df110e6) / vs btree-map (>1 = ordered faster; * = not precise)

| keys | op | 4,096 | 16,384 | 65,536 |
|---|---|---|---|---|
| street | valuesFor | 1.00 / 1.53 | 1.00 / 1.72 | 1.00 / 1.82 |
| street | valuesBetween | 1.03* / 0.71 | 1.03 / 0.63 | 1.01 / 0.68 |
| street | churn | 1.02 / 1.11 | 1.01 / 1.19 | 0.99 / 1.29 |
| street | build | 1.00 / 1.03 | 1.00 / 1.15 | 0.99 / 1.21 |
| dirs | valuesFor | 1.00 / 1.15 | 1.01 / 1.31 | 1.01 / 1.40* |
| dirs | valuesBetween | 1.00 / 0.43 | 1.00 / 0.49 | 1.01 / 0.67* |
| dirs | churn | 1.00 / 0.86 | 1.01 / 0.97 | 0.98* / 1.15 |
| dirs | build | 0.99 / 0.78 | 1.00 / 0.88 | 0.99 / 1.00 |
| links | valuesFor | 1.00 / 1.70 | 1.00 / 1.89 | 1.00 / 2.11 |
| links | valuesBetween | 1.01 / 0.69 | 0.99 / 0.74 | 1.00 / 0.74 |
| links | churn | 0.99 / 1.16 | 0.99 / 1.27 | 0.98 / 1.36 |
| links | build | 0.99 / 1.06 | 1.00 / 1.22 | 0.99 / 1.27 |
| url | valuesFor | 1.00 / 1.24 | 1.01 / 1.43 | 1.02* / 1.42 |
| url | valuesBetween | 1.02* / 0.32 | 1.02 / 0.33 | 1.01 / 0.53 |
| url | churn | 1.00 / 0.79 | 1.00 / 0.85 | 0.99 / 1.02* |
| url | build | 0.99 / 0.71 | 0.99 / 0.79 | 0.99 / 0.88 |
