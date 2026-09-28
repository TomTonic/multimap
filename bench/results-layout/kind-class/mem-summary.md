| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 116 | 20 | +83 ms | 57 |
| multi | u64 | 1048576 | baseline | 3 | 116 | 20 | +82 ms | 57 |
| multi | str | 1048576 | ordered | 3 | 140 | 31 | +143 ms | 67 |
| multi | str | 1048576 | baseline | 3 | 140 | 31 | +145 ms | 67 |
| multi | uuid | 1048576 | ordered | 3 | 147 | 32 | +140 ms | 71 |
| multi | uuid | 1048576 | baseline | 3 | 147 | 32 | +139 ms | 71 |
| multi | email | 1048576 | ordered | 3 | 139 | 30 | +129 ms | 66 |
| multi | email | 1048576 | baseline | 3 | 139 | 30 | +130 ms | 66 |
| multi | url | 1048576 | ordered | 3 | 192 | 39 | +187 ms | 93 |
| multi | url | 1048576 | baseline | 3 | 192 | 40 | +188 ms | 93 |
| multi | path | 300000 | ordered | 3 | 196 | 41 | +43 ms | 95 |
| multi | path | 300000 | baseline | 3 | 196 | 41 | +43 ms | 95 |
| multi | street | 212449 | ordered | 3 | 85 | 35 | +24 ms | 39 |
| multi | street | 212449 | baseline | 3 | 85 | 35 | +25 ms | 39 |
| unique | u64 | 1048576 | ordered | 3 | 41 | 17 | +73 ms | 19 |
| unique | u64 | 1048576 | baseline | 3 | 41 | 17 | +72 ms | 19 |
| unique | str | 1048576 | ordered | 3 | 66 | 27 | +128 ms | 30 |
| unique | str | 1048576 | baseline | 3 | 66 | 27 | +128 ms | 30 |
| unique | uuid | 1048576 | ordered | 3 | 68 | 28 | +124 ms | 31 |
| unique | uuid | 1048576 | baseline | 3 | 68 | 28 | +126 ms | 31 |
| unique | email | 1048576 | ordered | 3 | 65 | 26 | +117 ms | 29 |
| unique | email | 1048576 | baseline | 3 | 65 | 26 | +117 ms | 29 |
| unique | url | 1048576 | ordered | 3 | 117 | 37 | +166 ms | 56 |
| unique | url | 1048576 | baseline | 3 | 117 | 37 | +165 ms | 56 |
| unique | path | 300000 | ordered | 3 | 121 | 37 | +38 ms | 58 |
| unique | path | 300000 | baseline | 3 | 121 | 37 | +39 ms | 58 |
| unique | street | 212449 | ordered | 3 | 60 | 34 | +24 ms | 27 |
| unique | street | 212449 | baseline | 3 | 60 | 34 | +23 ms | 27 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
