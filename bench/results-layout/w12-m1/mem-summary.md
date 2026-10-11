| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | street | 212449 | ordered | 3 | 54 | 5 | +3 ms | 35 |
| natural | street | 212449 | btree-sets | 3 | 285 | 62 | +40 ms | 144 |
| natural | street | 212449 | baseline | 3 | 54 | 5 | +3 ms | 35 |
| natural | dirs | 86215 | ordered | 3 | 73 | 8 | +2 ms | 47 |
| natural | dirs | 86215 | btree-sets | 3 | 335 | 62 | +15 ms | 167 |
| natural | dirs | 86215 | baseline | 3 | 73 | 8 | +2 ms | 47 |
| natural | links | 199519 | ordered | 3 | 761 | 38 | +21 ms | 389 |
| natural | links | 199519 | btree-sets | 3 | 1506 | 78 | +56 ms | 756 |
| natural | links | 199519 | baseline | 3 | 762 | 38 | +21 ms | 389 |
| natural | url | 262144 | ordered | 3 | 150 | 14 | +10 ms | 89 |
| natural | url | 262144 | btree-sets | 3 | 416 | 64 | +46 ms | 209 |
| natural | url | 262144 | baseline | 3 | 150 | 14 | +10 ms | 89 |
| single-value | street | 212449 | ordered | 3 | 27 | 3 | +2 ms | 20 |
| single-value | street | 212449 | btree-map | 3 | 55 | 37 | +13 ms | 28 |
| single-value | street | 212449 | baseline | 3 | 27 | 3 | +2 ms | 20 |
| single-value | dirs | 86215 | ordered | 3 | 40 | 5 | +1 ms | 30 |
| single-value | dirs | 86215 | btree-map | 3 | 96 | 37 | +4 ms | 49 |
| single-value | dirs | 86215 | baseline | 3 | 40 | 5 | +1 ms | 30 |
| single-value | links | 199519 | ordered | 3 | 32 | 4 | +2 ms | 24 |
| single-value | links | 199519 | btree-map | 3 | 57 | 37 | +9 ms | 30 |
| single-value | links | 199519 | baseline | 3 | 32 | 4 | +3 ms | 24 |
| single-value | url | 262144 | ordered | 3 | 72 | 7 | +5 ms | 50 |
| single-value | url | 262144 | btree-map | 3 | 109 | 37 | +15 ms | 55 |
| single-value | url | 262144 | baseline | 3 | 72 | 7 | +5 ms | 50 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
