| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique | street | 212449 | ordered | 3 | 32 | 4 | -0 ms | 20 |
| unique | street | 212449 | btree-map | 3 | 47 | 37 | +18 ms | 20 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
