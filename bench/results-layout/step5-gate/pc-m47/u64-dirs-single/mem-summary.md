| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value | dirs | 86215 | ordered | 3 | 40 | 5 | +2 ms | 28 |
| single-value | dirs | 86215 | btree-map | 3 | 96 | 37 | +8 ms | 49 |
| single-value | dirs | 86215 | baseline | 3 | 77 | 37 | +11 ms | 41 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
