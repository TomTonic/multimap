| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value | u64 | 1024 | ordered | 2 | 40 | 2 | -0 ms | 47 |
| single-value | u64 | 1024 | btree-map | 2 | 57 | 37 | +0 ms | 50 |
| single-value | str | 1024 | ordered | 2 | 56 | 4 | +0 ms | 53 |
| single-value | str | 1024 | btree-map | 2 | 76 | 39 | +0 ms | 59 |
| single-value | uuid | 1024 | ordered | 2 | 76 | 4 | -0 ms | 75 |
| single-value | uuid | 1024 | btree-map | 2 | 96 | 40 | +0 ms | 71 |
| single-value | email | 1024 | ordered | 2 | 68 | 8 | -0 ms | 57 |
| single-value | email | 1024 | btree-map | 2 | 77 | 40 | +0 ms | 54 |
| single-value | url | 1024 | ordered | 2 | 100 | 13 | +0 ms | 87 |
| single-value | url | 1024 | btree-map | 2 | 126 | 37 | +0 ms | 91 |
| single-value | path | 1024 | ordered | 2 | 97 | 12 | +0 ms | 98 |
| single-value | path | 1024 | btree-map | 2 | 123 | 40 | +0 ms | 95 |
| single-value | street | 1024 | ordered | 2 | 42 | 8 | +0 ms | 47 |
| single-value | street | 1024 | btree-map | 2 | 67 | 49 | +0 ms | 55 |
| single-value | dirs | 1024 | ordered | 2 | 58 | 10 | -0 ms | 51 |
| single-value | dirs | 1024 | btree-map | 2 | 61 | 43 | -0 ms | 39 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
