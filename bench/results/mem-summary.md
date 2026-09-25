| keys | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---:|---:|---:|---:|
| u64 | ordered | 5 | 165 | 99 | +158 ms | 80 |
| u64 | hashed | 5 | 177 | 95 | +118 ms | 115 |
| u64 | btree-sets | 5 | 343 | 64 | +220 ms | 172 |
| u64 | map-sets | 5 | 360 | 81 | +197 ms | 206 |
| str | ordered | 5 | 204 | 111 | +235 ms | 100 |
| str | hashed | 5 | 196 | 95 | +116 ms | 121 |
| str | btree-sets | 5 | 363 | 64 | +220 ms | 178 |
| str | map-sets | 5 | 379 | 81 | +196 ms | 212 |

1048576 keys. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
