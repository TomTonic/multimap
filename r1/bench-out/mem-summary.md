| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | street | 212449 | ordered | 3 | 82 | 35 | +21 ms | 38 |
| multi | street | 212449 | hashed | 3 | 114 | 73 | +29 ms | 70 |
| multi | street | 212449 | btree-sets | 3 | 277 | 62 | +39 ms | 136 |
| multi | street | 212449 | map-sets | 3 | 274 | 58 | +29 ms | 150 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
