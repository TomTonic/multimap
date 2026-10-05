| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value | u64 | 4096 | ordered | 2 | 23 | 1 | +0 ms | 26 |
| single-value | u64 | 4096 | btree-map | 2 | 48 | 38 | +0 ms | 34 |
| single-value | str | 4096 | ordered | 2 | 37 | 1 | -0 ms | 36 |
| single-value | str | 4096 | btree-map | 2 | 68 | 38 | +0 ms | 39 |
| single-value | uuid | 4096 | ordered | 2 | 77 | 9 | +0 ms | 43 |
| single-value | uuid | 4096 | btree-map | 2 | 88 | 38 | +0 ms | 50 |
| single-value | email | 4096 | ordered | 2 | 47 | 2 | -0 ms | 44 |
| single-value | email | 4096 | btree-map | 2 | 69 | 38 | +0 ms | 39 |
| single-value | url | 4096 | ordered | 2 | 88 | 10 | +0 ms | 63 |
| single-value | url | 4096 | btree-map | 2 | 112 | 39 | +0 ms | 61 |
| single-value | path | 4096 | ordered | 2 | 77 | 9 | +0 ms | 59 |
| single-value | path | 4096 | btree-map | 2 | 113 | 37 | +0 ms | 64 |
| single-value | street | 4096 | ordered | 2 | 34 | 6 | +0 ms | 30 |
| single-value | street | 4096 | btree-map | 2 | 56 | 36 | +0 ms | 37 |
| single-value | dirs | 4096 | ordered | 2 | 47 | 7 | -0 ms | 39 |
| single-value | dirs | 4096 | btree-map | 2 | 79 | 37 | -0 ms | 38 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
