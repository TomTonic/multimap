| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-str | street | 212449 | ordered | 3 | 80 | 36 | +12 ms | 46 |
| natural-str | street | 212449 | btree-sets | 3 | 367 | 336 | +93 ms | 185 |
| natural-str | street | 212449 | baseline | 3 | 114 | 65 | +33 ms | 59 |

- natural-str street: the values are 28 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
