| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value | u64 | 16384 | ordered | 2 | 51 | 19 | +0 ms | 25 |
| single-value | u64 | 16384 | btree-map | 2 | 45 | 36 | +1 ms | 29 |
| single-value | str | 16384 | ordered | 2 | 37 | 3 | +0 ms | 28 |
| single-value | str | 16384 | btree-map | 2 | 65 | 36 | +1 ms | 36 |
| single-value | uuid | 16384 | ordered | 2 | 63 | 4 | +0 ms | 50 |
| single-value | uuid | 16384 | btree-map | 2 | 85 | 37 | +1 ms | 46 |
| single-value | email | 16384 | ordered | 2 | 63 | 11 | +0 ms | 28 |
| single-value | email | 16384 | btree-map | 2 | 67 | 37 | +1 ms | 36 |
| single-value | url | 16384 | ordered | 2 | 82 | 8 | +0 ms | 55 |
| single-value | url | 16384 | btree-map | 2 | 109 | 37 | +1 ms | 58 |
| single-value | path | 16384 | ordered | 2 | 65 | 7 | +0 ms | 47 |
| single-value | path | 16384 | btree-map | 2 | 109 | 36 | +1 ms | 58 |
| single-value | street | 16384 | ordered | 2 | 31 | 4 | +1 ms | 24 |
| single-value | street | 16384 | btree-map | 2 | 55 | 37 | +1 ms | 31 |
| single-value | dirs | 16384 | ordered | 2 | 49 | 6 | +0 ms | 37 |
| single-value | dirs | 16384 | btree-map | 2 | 96 | 37 | +1 ms | 49 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
