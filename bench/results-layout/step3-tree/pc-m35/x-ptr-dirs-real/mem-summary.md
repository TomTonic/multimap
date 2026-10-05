| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-ptr | dirs | 86215 | ordered | 3 | 108 | 106 | +22 ms | 55 |
| natural-ptr | dirs | 86215 | btree-sets | 3 | 335 | 271 | +29 ms | 167 |
| natural-ptr | dirs | 86215 | baseline | 3 | 106 | 102 | +23 ms | 54 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
