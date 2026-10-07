| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value | u64 | 262144 | ordered | 3 | 24 | 2 | +6 ms | 18 |
| single-value | u64 | 262144 | btree-map | 3 | 45 | 37 | +21 ms | 27 |
| single-value | u64 | 262144 | baseline | 3 | 53 | 21 | +19 ms | 30 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
