| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | dirs | 86215 | ordered | 3 | 149 | 139 | +27 ms | 76 |
| multi-str | dirs | 86215 | ordered-lpage-mv | 3 | 158 | 72 | +12 ms | 87 |
| multi-str | dirs | 86215 | ordered-lpage-mvzc | 3 | 156 | 72 | +13 ms | 87 |

- multi-str dirs: the values are 68 bytes of string per key. Only ordered-lpage holds them in its heap; every other candidate holds 16-byte headers that point into one shared buffer (see toVs), so its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
