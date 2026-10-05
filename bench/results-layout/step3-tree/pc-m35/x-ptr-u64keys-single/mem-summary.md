| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value-ptr | u64 | 262144 | ordered | 3 | 53 | 53 | +30 ms | 30 |
| single-value-ptr | u64 | 262144 | btree-map | 3 | 45 | 37 | +28 ms | 27 |
| single-value-ptr | u64 | 262144 | baseline | 3 | 45 | 45 | +28 ms | 26 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
