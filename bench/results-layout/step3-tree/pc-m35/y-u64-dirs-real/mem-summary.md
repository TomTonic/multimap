| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | dirs | 86215 | ordered | 3 | 108 | 38 | +11 ms | 55 |
| natural | dirs | 86215 | baseline | 3 | 106 | 38 | +11 ms | 54 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
