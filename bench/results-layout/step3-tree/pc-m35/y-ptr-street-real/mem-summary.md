| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-ptr | street | 212449 | ordered | 3 | 90 | 88 | +45 ms | 46 |
| natural-ptr | street | 212449 | baseline | 3 | 86 | 84 | +43 ms | 44 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
