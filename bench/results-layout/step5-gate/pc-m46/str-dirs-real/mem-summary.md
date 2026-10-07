| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-str | dirs | 86215 | ordered | 3 | 132 | 52 | +7 ms | 74 |
| natural-str | dirs | 86215 | btree-sets | 3 | 421 | 355 | +35 ms | 210 |
| natural-str | dirs | 86215 | baseline | 3 | 165 | 79 | +15 ms | 83 |

- natural-str dirs: the values are 68 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
