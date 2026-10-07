| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value | street | 212449 | ordered | 3 | 27 | 3 | +4 ms | 18 |
| single-value | street | 212449 | btree-map | 3 | 55 | 37 | +19 ms | 28 |
| single-value | street | 212449 | baseline | 3 | 67 | 34 | +27 ms | 35 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
