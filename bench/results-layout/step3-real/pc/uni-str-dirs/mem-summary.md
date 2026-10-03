| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique-str | dirs | 86215 | ordered | 3 | 69 | 72 | +16 ms | 33 |
| unique-str | dirs | 86215 | btree-map | 3 | 99 | 49 | +11 ms | 47 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
