| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value-ptr | dirs | 86215 | ordered | 3 | 77 | 76 | +15 ms | 41 |
| single-value-ptr | dirs | 86215 | baseline | 3 | 68 | 67 | +16 ms | 36 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
