
natural: ordered speed vs main / vs btree-sets (>1 = ordered faster; * = not precise)

| keys | op | 4,096 | 16,384 | 65,536 |
|---|---|---|---|---|
| street | valuesFor | 0.81* / 1.92 | 0.86* / 2.03 | 1.03* / 2.50 |
| street | valuesBetween | 2.11 / 3.85 | 1.93 / 3.38 | 2.17 / 4.85 |
| street | churn | 0.66 / 1.32 | 0.75 / 1.48 | 0.94* / 1.52 |
| street | build | 0.63 / 1.32 | 0.71 / 1.44 | 0.90* / 1.65 |
| dirs | valuesFor | 0.85 / 1.48 | 0.88* / 1.58 | 1.16 / 2.00 |
| dirs | valuesBetween | 1.67 / 2.92 | 1.78 / 3.00 | 2.15 / 4.06 |
| dirs | churn | 0.78 / 1.15 | 0.90* / 1.26 | 1.00* / 1.39 |
| dirs | build | 0.72 / 1.12 | 0.79 / 1.21 | 1.01* / 1.48 |
| links | valuesFor | 0.99 / 1.85 | 1.03 / 1.84 | 1.09 / 1.91 |
| links | valuesBetween | 1.66 / 2.74 | 1.66 / 2.85 | 1.60 / 2.98 |
| links | churn | 0.76 / 1.29 | 0.83* / 1.43 | 0.85 / 1.35 |
| links | build | 0.72 / 1.20 | 0.82 / 1.40 | 1.00 / 1.42 |
| url | valuesFor | 0.87 / 1.63 | 0.91 / 1.72 | 1.12* / 1.97 |
| url | valuesBetween | 1.50 / 2.49 | 1.49 / 2.50 | 1.71 / 3.11 |
| url | churn | 0.74 / 1.11 | 0.85 / 1.23 | 0.90* / 1.24 |
| url | build | 0.70 / 1.07 | 0.77 / 1.20 | 0.96 / 1.39 |

single-value: ordered speed vs main / vs btree-map (>1 = ordered faster; * = not precise)

| keys | op | 4,096 | 16,384 | 65,536 |
|---|---|---|---|---|
| street | valuesFor | 0.83 / 1.55 | 0.92* / 1.73 | 1.03* / 1.81 |
| street | valuesBetween | 2.54 / 0.68 | 2.24 / 0.62 | 2.43 / 0.67 |
| street | churn | 0.90 / 1.10 | 0.97 / 1.19* | 1.12 / 1.28 |
| street | build | 0.77 / 1.02* | 0.83 / 1.14 | 0.98 / 1.20 |
| dirs | valuesFor | 0.88 / 1.17* | 0.92 / 1.32 | 1.17* / 1.39* |
| dirs | valuesBetween | 1.68 / 0.42 | 1.82 / 0.49 | 2.13 / 0.66* |
| dirs | churn | 0.95 / 0.86 | 1.05 / 0.96 | 1.22 / 1.16 |
| dirs | build | 0.80 / 0.78 | 0.88 / 0.87 | 1.04 / 0.99 |
| links | valuesFor | 0.88 / 1.69 | 0.91* / 1.89 | 1.02* / 2.10 |
| links | valuesBetween | 2.02 / 0.67 | 2.12 / 0.72 | 2.21 / 0.73 |
| links | churn | 0.89 / 1.15 | 0.97* / 1.27 | 1.12* / 1.36 |
| links | build | 0.73 / 1.06 | 0.81 / 1.23 | 0.92 / 1.27 |
| url | valuesFor | 0.85 / 1.26 | 0.91 / 1.43 | 1.13* / 1.38 |
| url | valuesBetween | 1.41 / 0.31 | 1.40 / 0.33 | 1.72 / 0.52* |
| url | churn | 0.82 / 0.79 | 0.91* / 0.87 | 1.05* / 1.02 |
| url | build | 0.72 / 0.71 | 0.76 / 0.80 | 0.88 / 0.89 |
