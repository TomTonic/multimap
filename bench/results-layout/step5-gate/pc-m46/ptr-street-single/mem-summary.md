| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value-ptr | street | 212449 | ordered | 3 | 28 | 27 | +12 ms | 19 |
| single-value-ptr | street | 212449 | btree-map | 3 | 55 | 37 | +26 ms | 28 |
| single-value-ptr | street | 212449 | baseline | 3 | 67 | 67 | +36 ms | 34 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
