| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value | u64 | 262144 | ordered | 2 | 24 | 2 | +4 ms | 20 |
| single-value | u64 | 262144 | btree-map | 2 | 45 | 37 | +15 ms | 27 |
| single-value | str | 262144 | ordered | 2 | 33 | 4 | +2 ms | 24 |
| single-value | str | 262144 | btree-map | 2 | 64 | 36 | +18 ms | 33 |
| single-value | uuid | 262144 | ordered | 2 | 57 | 4 | +2 ms | 45 |
| single-value | uuid | 262144 | btree-map | 2 | 85 | 37 | +19 ms | 43 |
| single-value | email | 262144 | ordered | 2 | 48 | 6 | +2 ms | 31 |
| single-value | email | 262144 | btree-map | 2 | 66 | 37 | +19 ms | 34 |
| single-value | url | 262144 | ordered | 2 | 72 | 8 | +7 ms | 48 |
| single-value | url | 262144 | btree-map | 2 | 109 | 33 | +18 ms | 55 |
| single-value | path | 262144 | ordered | 2 | 50 | 6 | +5 ms | 36 |
| single-value | path | 262144 | btree-map | 2 | 109 | 37 | +20 ms | 55 |
| single-value | street | 212449 | ordered | 2 | 27 | 3 | +3 ms | 20 |
| single-value | street | 212449 | btree-map | 2 | 55 | 37 | +16 ms | 28 |
| single-value | dirs | 86215 | ordered | 2 | 41 | 5 | +2 ms | 30 |
| single-value | dirs | 86215 | btree-map | 2 | 96 | 37 | +7 ms | 49 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
