| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | dirs | 86215 | ordered | 3 | 98 | 37 | +8 ms | 46 |
| multi | dirs | 86215 | hashed | 3 | 170 | 81 | +12 ms | 101 |
| multi | dirs | 86215 | btree-sets | 3 | 327 | 62 | +15 ms | 159 |
| multi | dirs | 86215 | map-sets | 3 | 331 | 66 | +12 ms | 181 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
