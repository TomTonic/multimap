| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | dirs | 86215 | ordered | 3 | 149 | 135 | +20 ms | 75 |
| multi-str | dirs | 86215 | btree-sets | 3 | 421 | 315 | +25 ms | 209 |
| multi-str | dirs | 86215 | ordered-lpage | 3 | 149 | 135 | +20 ms | 75 |
| multi-str | dirs | 86215 | ordered-lpage-mv | 3 | 139 | 59 | +6 ms | 79 |
| multi-str | dirs | 86215 | ordered-lpage-mvzc | 3 | 139 | 59 | +6 ms | 79 |

- multi-str dirs: the values are 68 bytes of string per key. Only ordered-lpage holds them in its heap; every other candidate holds 16-byte headers that point into one shared buffer (see toVs), so its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
