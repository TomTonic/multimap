| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural-str | street | 212449 | ordered | 3 | 80 | 36 | +8 ms | 48 |
| natural-str | street | 212449 | btree-sets | 3 | 367 | 301 | +70 ms | 185 |
| natural-str | street | 212449 | baseline | 3 | 114 | 65 | +26 ms | 59 |
| natural-str | dirs | 86215 | ordered | 3 | 132 | 52 | +5 ms | 77 |
| natural-str | dirs | 86215 | btree-sets | 3 | 421 | 315 | +27 ms | 209 |
| natural-str | dirs | 86215 | baseline | 3 | 165 | 78 | +11 ms | 83 |
| single-value-str | street | 212449 | ordered | 3 | 31 | 4 | +3 ms | 23 |
| single-value-str | street | 212449 | btree-map | 3 | 66 | 49 | +18 ms | 34 |
| single-value-str | street | 212449 | baseline | 3 | 69 | 34 | +21 ms | 36 |
| single-value-str | dirs | 86215 | ordered | 3 | 51 | 6 | +2 ms | 37 |
| single-value-str | dirs | 86215 | btree-map | 3 | 107 | 49 | +7 ms | 54 |
| single-value-str | dirs | 86215 | baseline | 3 | 87 | 37 | +8 ms | 46 |

- natural-str street: the values are 28 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.
- natural-str dirs: the values are 68 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.
- single-value-str street: the values are 10 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.
- single-value-str dirs: the values are 15 bytes of string per key. A candidate that copies values into its own objects holds them in its heap; one that holds 16-byte headers points into one shared buffer (see toVs), and its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
