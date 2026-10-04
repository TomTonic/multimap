| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique-str | dirs | 86215 | ordered | 3 | 78 | 72 | +19 ms | 41 |
| unique-str | dirs | 86215 | btree-map | 3 | 108 | 49 | +14 ms | 55 |
| unique-str | dirs | 86215 | ordered-lpage | 3 | 54 | 9 | +3 ms | 36 |
| unique-str | dirs | 86215 | ordered-lpage-zc | 3 | 54 | 9 | +2 ms | 36 |

- unique-str dirs: the values are 15 bytes of string per key. Only ordered-lpage holds them in its heap; every other candidate holds 16-byte headers that point into one shared buffer (see toVs), so its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
