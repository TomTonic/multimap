| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | links | 4096 | ordered | 1 | 838 | 42 | -0 ms | 438 |
| natural | links | 4096 | btree-sets | 1 | 1595 | 103 | +1 ms | 822 |
| single-value | links | 4096 | ordered | 1 | 41 | 6 | -1 ms | 33 |
| single-value | links | 4096 | btree-map | 1 | 60 | 40 | -0 ms | 36 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
