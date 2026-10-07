| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-ptr | street | 212449 | ordered | 3 | 54 | 52 | +20 ms | 33 |
| natural-ptr | street | 212449 | btree-sets | 3 | 285 | 256 | +76 ms | 144 |
| natural-ptr | street | 212449 | baseline | 3 | 89 | 88 | +45 ms | 46 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
