| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | links | 4096 | ordered | 1 | 725 | 40 | +0 ms | 341 |
| natural | links | 4096 | btree-sets | 1 | 1466 | 103 | +1 ms | 683 |
| single-value | links | 4096 | ordered | 1 | 39 | 5 | +0 ms | 30 |
| single-value | links | 4096 | btree-map | 1 | 60 | 38 | -0 ms | 34 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
