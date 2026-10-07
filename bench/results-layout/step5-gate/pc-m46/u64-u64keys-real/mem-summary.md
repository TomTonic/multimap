| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | u64 | 262144 | ordered | 3 | 102 | 7 | +6 ms | 59 |
| natural | u64 | 262144 | btree-sets | 3 | 352 | 84 | +65 ms | 181 |
| natural | u64 | 262144 | baseline | 3 | 125 | 23 | +19 ms | 66 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
