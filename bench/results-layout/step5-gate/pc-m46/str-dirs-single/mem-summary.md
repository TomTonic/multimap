| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| single-value-str | dirs | 86215 | ordered | 3 | 50 | 6 | +3 ms | 34 |
| single-value-str | dirs | 86215 | btree-map | 3 | 108 | 49 | +10 ms | 55 |
| single-value-str | dirs | 86215 | baseline | 3 | 87 | 37 | +10 ms | 46 |

- single-value-str dirs: the values are 15 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
