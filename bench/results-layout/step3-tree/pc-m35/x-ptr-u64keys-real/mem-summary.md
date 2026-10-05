| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-ptr | u64 | 262144 | ordered | 3 | 125 | 119 | +57 ms | 66 |
| natural-ptr | u64 | 262144 | btree-sets | 3 | 352 | 328 | +116 ms | 181 |
| natural-ptr | u64 | 262144 | baseline | 3 | 124 | 118 | +59 ms | 65 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
