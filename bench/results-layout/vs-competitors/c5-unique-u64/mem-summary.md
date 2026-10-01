| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique | u64 | 1048576 | ordered | 3 | 41 | 17 | +68 ms | 19 |
| unique | u64 | 1048576 | btree-map | 3 | 37 | 37 | +78 ms | 19 |
| unique | str | 1048576 | ordered | 3 | 52 | 28 | +123 ms | 23 |
| unique | str | 1048576 | btree-map | 3 | 56 | 37 | +93 ms | 25 |
| unique | uuid | 1048576 | ordered | 3 | 68 | 28 | +115 ms | 31 |
| unique | uuid | 1048576 | btree-map | 3 | 77 | 37 | +95 ms | 35 |
| unique | email | 1048576 | ordered | 3 | 61 | 26 | +109 ms | 28 |
| unique | email | 1048576 | btree-map | 3 | 57 | 37 | +90 ms | 25 |
| unique | url | 1048576 | ordered | 3 | 86 | 36 | +156 ms | 41 |
| unique | url | 1048576 | btree-map | 3 | 101 | 37 | +86 ms | 47 |
| unique | path | 300000 | ordered | 3 | 72 | 38 | +42 ms | 34 |
| unique | path | 300000 | btree-map | 3 | 101 | 36 | +30 ms | 47 |
| unique | street | 212449 | ordered | 3 | 59 | 34 | +24 ms | 27 |
| unique | street | 212449 | btree-map | 3 | 47 | 37 | +19 ms | 20 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
