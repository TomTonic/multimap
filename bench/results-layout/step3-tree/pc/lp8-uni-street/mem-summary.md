| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique-str | street | 212449 | ordered | 3 | 67 | 63 | +39 ms | 35 |
| unique-str | street | 212449 | btree-map | 3 | 66 | 49 | +31 ms | 34 |
| unique-str | street | 212449 | ordered-lpage | 3 | 30 | 4 | +1 ms | 20 |
| unique-str | street | 212449 | ordered-lpage-zc | 3 | 30 | 4 | +4 ms | 20 |

- unique-str street: the values are 10 bytes of string per key. Only ordered-lpage holds them in its heap; every other candidate holds 16-byte headers that point into one shared buffer (see toVs), so its heap figure does not count them.

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
