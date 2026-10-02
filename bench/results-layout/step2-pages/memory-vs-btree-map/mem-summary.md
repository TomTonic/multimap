| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique | u64 | 262144 | ordered | 3 | 18 | 1 | +1 ms | 10 |
| unique | u64 | 262144 | btree-map | 3 | 37 | 37 | +21 ms | 19 |
| unique | str | 262144 | ordered | 3 | 31 | 3 | +4 ms | 21 |
| unique | str | 262144 | btree-map | 3 | 56 | 36 | +27 ms | 25 |
| unique | uuid | 262144 | ordered | 3 | 65 | 4 | +2 ms | 44 |
| unique | uuid | 262144 | btree-map | 3 | 77 | 37 | +26 ms | 35 |
| unique | email | 262144 | ordered | 3 | 56 | 7 | +2 ms | 32 |
| unique | email | 262144 | btree-map | 3 | 58 | 37 | +24 ms | 26 |
| unique | url | 262144 | ordered | 3 | 91 | 15 | +3 ms | 60 |
| unique | url | 262144 | btree-map | 3 | 101 | 41 | +13 ms | 47 |
| unique | path | 262144 | ordered | 3 | 67 | 11 | +7 ms | 47 |
| unique | path | 262144 | btree-map | 3 | 101 | 37 | +24 ms | 47 |
| unique | street | 212449 | ordered | 3 | 32 | 4 | +3 ms | 20 |
| unique | street | 212449 | btree-map | 3 | 47 | 37 | +19 ms | 20 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
