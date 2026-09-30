| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 116 | 20 | +85 ms | 57 |
| multi | u64 | 1048576 | baseline | 3 | 116 | 20 | +85 ms | 57 |
| multi | str | 1048576 | ordered | 3 | 129 | 30 | +150 ms | 62 |
| multi | str | 1048576 | baseline | 3 | 140 | 31 | +147 ms | 67 |
| multi | uuid | 1048576 | ordered | 3 | 147 | 31 | +143 ms | 70 |
| multi | uuid | 1048576 | baseline | 3 | 147 | 32 | +140 ms | 71 |
| multi | email | 1048576 | ordered | 3 | 137 | 30 | +132 ms | 65 |
| multi | email | 1048576 | baseline | 3 | 139 | 30 | +131 ms | 66 |
| multi | url | 1048576 | ordered | 3 | 161 | 41 | +193 ms | 78 |
| multi | url | 1048576 | baseline | 3 | 192 | 39 | +195 ms | 93 |
| multi | path | 300000 | ordered | 3 | 149 | 41 | +44 ms | 72 |
| multi | path | 300000 | baseline | 3 | 196 | 41 | +44 ms | 95 |
| multi | street | 212449 | ordered | 3 | 82 | 35 | +24 ms | 38 |
| multi | street | 212449 | baseline | 3 | 85 | 35 | +24 ms | 39 |
| unique | u64 | 1048576 | ordered | 3 | 41 | 17 | +77 ms | 19 |
| unique | u64 | 1048576 | baseline | 3 | 41 | 17 | +78 ms | 19 |
| unique | str | 1048576 | ordered | 3 | 51 | 27 | +135 ms | 23 |
| unique | str | 1048576 | baseline | 3 | 66 | 27 | +132 ms | 30 |
| unique | uuid | 1048576 | ordered | 3 | 68 | 28 | +125 ms | 31 |
| unique | uuid | 1048576 | baseline | 3 | 68 | 28 | +124 ms | 31 |
| unique | email | 1048576 | ordered | 3 | 62 | 26 | +117 ms | 28 |
| unique | email | 1048576 | baseline | 3 | 65 | 26 | +117 ms | 29 |
| unique | url | 1048576 | ordered | 3 | 85 | 36 | +177 ms | 40 |
| unique | url | 1048576 | baseline | 3 | 117 | 36 | +177 ms | 56 |
| unique | path | 300000 | ordered | 3 | 72 | 37 | +40 ms | 34 |
| unique | path | 300000 | baseline | 3 | 121 | 37 | +38 ms | 58 |
| unique | street | 212449 | ordered | 3 | 59 | 34 | +24 ms | 27 |
| unique | street | 212449 | baseline | 3 | 60 | 34 | +24 ms | 27 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
