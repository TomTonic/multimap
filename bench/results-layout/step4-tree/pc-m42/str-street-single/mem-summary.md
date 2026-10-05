| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value-str | street | 212449 | ordered | 3 | 31 | 4 | +4 ms | 22 |
| single-value-str | street | 212449 | btree-map | 3 | 66 | 49 | +29 ms | 34 |
| single-value-str | street | 212449 | baseline | 3 | 69 | 34 | +27 ms | 36 |

- single-value-str street: the values are 10 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
