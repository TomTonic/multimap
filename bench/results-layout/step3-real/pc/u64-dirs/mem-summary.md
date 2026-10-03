| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | dirs | 86215 | ordered | 3 | 98 | 38 | +12 ms | 46 |
| multi | dirs | 86215 | hashed | 3 | 170 | 88 | +11 ms | 100 |
| multi | dirs | 86215 | btree-sets | 3 | 327 | 82 | +20 ms | 159 |
| multi | dirs | 86215 | map-sets | 3 | 331 | 88 | +16 ms | 180 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
