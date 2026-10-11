| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | street | 212449 | ordered | 3 | 54 | 5 | +6 ms | 35 |
| natural | street | 212449 | btree-sets | 3 | 285 | 81 | +51 ms | 144 |
| natural | street | 212449 | baseline | 3 | 54 | 5 | +5 ms | 35 |
| natural | dirs | 86215 | ordered | 3 | 74 | 8 | +2 ms | 48 |
| natural | dirs | 86215 | btree-sets | 3 | 336 | 83 | +18 ms | 168 |
| natural | dirs | 86215 | baseline | 3 | 72 | 8 | +2 ms | 46 |
| natural | links | 199519 | ordered | 3 | 762 | 40 | +31 ms | 390 |
| natural | links | 199519 | btree-sets | 3 | 1506 | 98 | +74 ms | 758 |
| natural | links | 199519 | baseline | 3 | 761 | 40 | +31 ms | 391 |
| natural | url | 262144 | ordered | 3 | 150 | 21 | +11 ms | 89 |
| natural | url | 262144 | btree-sets | 3 | 416 | 87 | +60 ms | 209 |
| natural | url | 262144 | baseline | 3 | 150 | 15 | +13 ms | 89 |
| single-value | street | 212449 | ordered | 3 | 27 | 3 | +4 ms | 20 |
| single-value | street | 212449 | btree-map | 3 | 55 | 37 | +16 ms | 28 |
| single-value | street | 212449 | baseline | 3 | 27 | 3 | +4 ms | 20 |
| single-value | dirs | 86215 | ordered | 3 | 40 | 5 | +2 ms | 30 |
| single-value | dirs | 86215 | btree-map | 3 | 95 | 37 | +6 ms | 48 |
| single-value | dirs | 86215 | baseline | 3 | 40 | 5 | +2 ms | 30 |
| single-value | links | 199519 | ordered | 3 | 32 | 4 | +3 ms | 23 |
| single-value | links | 199519 | btree-map | 3 | 57 | 37 | +13 ms | 30 |
| single-value | links | 199519 | baseline | 3 | 32 | 4 | +3 ms | 24 |
| single-value | url | 262144 | ordered | 3 | 72 | 10 | +6 ms | 50 |
| single-value | url | 262144 | btree-map | 3 | 109 | 38 | +18 ms | 55 |
| single-value | url | 262144 | baseline | 3 | 72 | 12 | +7 ms | 50 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
