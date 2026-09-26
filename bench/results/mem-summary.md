| values | keys | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---|---:|---:|---:|---:|---:|
| multi | u64 | ordered | 5 | 141 | 64 | +73 ms | 76 |
| multi | u64 | hashed | 5 | 177 | 101 | +216 ms | 115 |
| multi | u64 | btree-sets | 5 | 343 | 73 | +356 ms | 172 |
| multi | u64 | map-sets | 5 | 360 | 86 | +322 ms | 206 |
| multi | str | ordered | 5 | 120 | 30 | +85 ms | 70 |
| multi | str | hashed | 5 | 196 | 101 | +235 ms | 121 |
| multi | str | btree-sets | 5 | 363 | 75 | +368 ms | 178 |
| multi | str | map-sets | 5 | 379 | 85 | +336 ms | 212 |
| unique | u64 | ordered | 5 | 17 | 1 | +4 ms | 11 |
| unique | u64 | btree-map | 5 | 37 | 37 | +57 ms | 19 |
| unique | str | ordered | 5 | 33 | 23 | +38 ms | 25 |
| unique | str | btree-map | 5 | 56 | 37 | +76 ms | 25 |

1048576 keys. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
