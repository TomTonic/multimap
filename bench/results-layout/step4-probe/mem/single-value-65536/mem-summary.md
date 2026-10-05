| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value | u64 | 65536 | ordered | 2 | 38 | 9 | +2 ms | 25 |
| single-value | u64 | 65536 | btree-map | 2 | 45 | 37 | +4 ms | 28 |
| single-value | str | 65536 | ordered | 2 | 27 | 2 | +0 ms | 25 |
| single-value | str | 65536 | btree-map | 2 | 64 | 36 | +4 ms | 34 |
| single-value | uuid | 65536 | ordered | 2 | 66 | 8 | +1 ms | 35 |
| single-value | uuid | 65536 | btree-map | 2 | 85 | 37 | +4 ms | 44 |
| single-value | email | 65536 | ordered | 2 | 48 | 3 | -0 ms | 37 |
| single-value | email | 65536 | btree-map | 2 | 66 | 37 | +4 ms | 35 |
| single-value | url | 65536 | ordered | 2 | 77 | 11 | +2 ms | 51 |
| single-value | url | 65536 | btree-map | 2 | 108 | 34 | +5 ms | 56 |
| single-value | path | 65536 | ordered | 2 | 57 | 7 | +1 ms | 41 |
| single-value | path | 65536 | btree-map | 2 | 108 | 36 | +4 ms | 56 |
| single-value | street | 65536 | ordered | 2 | 28 | 4 | +1 ms | 22 |
| single-value | street | 65536 | btree-map | 2 | 54 | 36 | +5 ms | 29 |
| single-value | dirs | 65536 | ordered | 2 | 40 | 5 | +1 ms | 30 |
| single-value | dirs | 65536 | btree-map | 2 | 95 | 36 | +5 ms | 49 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
