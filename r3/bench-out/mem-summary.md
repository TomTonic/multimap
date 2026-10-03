| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | street | 212449 | ordered | 3 | 111 | 108 | +42 ms | 53 |
| multi-str | street | 212449 | hashed | 3 | 153 | 141 | +36 ms | 90 |
| multi-str | street | 212449 | btree-sets | 3 | 359 | 301 | +67 ms | 177 |
| multi-str | street | 212449 | map-sets | 3 | 355 | 298 | +42 ms | 191 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
