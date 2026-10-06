| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | u64 | 262144 | ordered | 3 | 123 | 21 | +14 ms | 64 |
| natural | u64 | 262144 | btree-sets | 3 | 352 | 64 | +47 ms | 181 |
| natural | u64 | 262144 | baseline | 3 | 125 | 23 | +16 ms | 66 |
| natural | street | 212449 | ordered | 3 | 73 | 21 | +12 ms | 40 |
| natural | street | 212449 | btree-sets | 3 | 285 | 62 | +39 ms | 144 |
| natural | street | 212449 | baseline | 3 | 90 | 35 | +20 ms | 46 |
| natural | dirs | 86215 | ordered | 3 | 95 | 28 | +6 ms | 51 |
| natural | dirs | 86215 | btree-sets | 3 | 336 | 62 | +15 ms | 167 |
| natural | dirs | 86215 | baseline | 3 | 108 | 37 | +8 ms | 55 |
| single-value | u64 | 262144 | ordered | 3 | 24 | 2 | +3 ms | 20 |
| single-value | u64 | 262144 | btree-map | 3 | 45 | 37 | +14 ms | 27 |
| single-value | u64 | 262144 | baseline | 3 | 53 | 21 | +15 ms | 29 |
| single-value | street | 212449 | ordered | 3 | 27 | 3 | +2 ms | 20 |
| single-value | street | 212449 | btree-map | 3 | 55 | 37 | +13 ms | 28 |
| single-value | street | 212449 | baseline | 3 | 67 | 34 | +20 ms | 35 |
| single-value | dirs | 86215 | ordered | 3 | 40 | 5 | +1 ms | 30 |
| single-value | dirs | 86215 | btree-map | 3 | 96 | 37 | +4 ms | 49 |
| single-value | dirs | 86215 | baseline | 3 | 77 | 37 | +7 ms | 40 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
