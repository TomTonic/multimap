| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | street | 212449 | ordered | 3 | 119 | 111 | +54 ms | 61 |
| multi-str | street | 212449 | btree-sets | 3 | 367 | 337 | +91 ms | 185 |
| multi-str | street | 212449 | ordered-lpage | 3 | 119 | 111 | +53 ms | 61 |

- multi-str street: the values are 28 bytes of string per key. Only ordered-lpage holds them in its heap; every other candidate holds 16-byte headers that point into one shared buffer (see toVs), so its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
