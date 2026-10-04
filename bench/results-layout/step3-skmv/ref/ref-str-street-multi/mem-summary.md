| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | street | 212449 | ordered | 3 | 119 | 111 | +52 ms | 61 |
| multi-str | street | 212449 | hashed | 3 | 160 | 141 | +49 ms | 98 |
| multi-str | street | 212449 | btree-sets | 3 | 367 | 336 | +92 ms | 185 |
| multi-str | street | 212449 | map-sets | 3 | 363 | 342 | +68 ms | 199 |
| multi-str | street | 212449 | baseline | 3 | 119 | 111 | +56 ms | 62 |

- multi-str street: the values are 28 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
