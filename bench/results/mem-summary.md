| values | keys | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---|---:|---:|---:|---:|---:|
| multi | u64 | ordered | 5 | 149 | 83 | +238 ms | 72 |
| multi | u64 | hashed | 5 | 177 | 101 | +235 ms | 115 |
| multi | u64 | btree-sets | 5 | 343 | 73 | +384 ms | 172 |
| multi | u64 | map-sets | 5 | 360 | 85 | +347 ms | 206 |
| multi | str | ordered | 5 | 177 | 111 | +336 ms | 86 |
| multi | str | hashed | 5 | 196 | 101 | +255 ms | 121 |
| multi | str | btree-sets | 5 | 363 | 73 | +400 ms | 178 |
| multi | str | map-sets | 5 | 379 | 86 | +367 ms | 212 |
| unique | u64 | ordered | 5 | 73 | 81 | +202 ms | 35 |
| unique | u64 | btree-map | 5 | 37 | 37 | +64 ms | 19 |
| unique | str | ordered | 5 | 101 | 109 | +288 ms | 48 |
| unique | str | btree-map | 5 | 56 | 37 | +84 ms | 25 |

1048576 keys. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
