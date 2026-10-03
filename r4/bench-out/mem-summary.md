| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | dirs | 86215 | ordered | 3 | 141 | 135 | +20 ms | 67 |
| multi-str | dirs | 86215 | hashed | 3 | 216 | 163 | +16 ms | 123 |
| multi-str | dirs | 86215 | btree-sets | 3 | 413 | 315 | +25 ms | 201 |
| multi-str | dirs | 86215 | map-sets | 3 | 417 | 318 | +18 ms | 223 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
