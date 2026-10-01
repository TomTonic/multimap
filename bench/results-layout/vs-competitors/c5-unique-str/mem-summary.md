| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique-str | u64 | 1048576 | ordered | 3 | 41 | 43 | +139 ms | 19 |
| unique-str | u64 | 1048576 | btree-map | 3 | 48 | 48 | +110 ms | 25 |
| unique-str | str | 1048576 | ordered | 3 | 56 | 55 | +228 ms | 26 |
| unique-str | str | 1048576 | btree-map | 3 | 68 | 48 | +132 ms | 31 |
| unique-str | uuid | 1048576 | ordered | 3 | 84 | 77 | +237 ms | 39 |
| unique-str | uuid | 1048576 | btree-map | 3 | 89 | 48 | +134 ms | 41 |
| unique-str | email | 1048576 | ordered | 3 | 67 | 66 | +215 ms | 31 |
| unique-str | email | 1048576 | btree-map | 3 | 69 | 48 | +133 ms | 32 |
| unique-str | url | 1048576 | ordered | 3 | 100 | 98 | +313 ms | 48 |
| unique-str | url | 1048576 | btree-map | 3 | 112 | 48 | +126 ms | 53 |
| unique-str | path | 300000 | ordered | 3 | 80 | 78 | +68 ms | 38 |
| unique-str | path | 300000 | btree-map | 3 | 112 | 48 | +43 ms | 53 |
| unique-str | street | 212449 | ordered | 3 | 59 | 63 | +38 ms | 27 |
| unique-str | street | 212449 | btree-map | 3 | 58 | 49 | +30 ms | 26 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
