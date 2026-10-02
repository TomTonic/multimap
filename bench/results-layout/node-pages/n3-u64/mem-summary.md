| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 117 | 20 | +74 ms | 57 |
| multi | u64 | 1048576 | baseline | 3 | 117 | 20 | +74 ms | 57 |
| multi | str | 1048576 | ordered | 3 | 129 | 31 | +137 ms | 62 |
| multi | str | 1048576 | baseline | 3 | 129 | 30 | +137 ms | 62 |
| multi | uuid | 1048576 | ordered | 3 | 147 | 31 | +129 ms | 71 |
| multi | uuid | 1048576 | baseline | 3 | 147 | 31 | +127 ms | 71 |
| multi | email | 1048576 | ordered | 3 | 137 | 30 | +120 ms | 66 |
| multi | email | 1048576 | baseline | 3 | 137 | 30 | +122 ms | 66 |
| multi | url | 1048576 | ordered | 3 | 162 | 40 | +176 ms | 79 |
| multi | url | 1048576 | baseline | 3 | 162 | 39 | +177 ms | 79 |
| multi | path | 300000 | ordered | 3 | 149 | 41 | +45 ms | 73 |
| multi | path | 300000 | baseline | 3 | 149 | 41 | +46 ms | 73 |
| multi | street | 212449 | ordered | 3 | 82 | 35 | +27 ms | 38 |
| multi | street | 212449 | baseline | 3 | 82 | 35 | +26 ms | 38 |
| unique | u64 | 1048576 | ordered | 3 | 16 | 1 | +2 ms | 9 |
| unique | u64 | 1048576 | baseline | 3 | 41 | 17 | +69 ms | 19 |
| unique | str | 1048576 | ordered | 3 | 52 | 28 | +124 ms | 23 |
| unique | str | 1048576 | baseline | 3 | 52 | 28 | +124 ms | 23 |
| unique | uuid | 1048576 | ordered | 3 | 68 | 28 | +117 ms | 31 |
| unique | uuid | 1048576 | baseline | 3 | 68 | 28 | +117 ms | 31 |
| unique | email | 1048576 | ordered | 3 | 61 | 26 | +110 ms | 28 |
| unique | email | 1048576 | baseline | 3 | 61 | 26 | +111 ms | 28 |
| unique | url | 1048576 | ordered | 3 | 86 | 35 | +159 ms | 41 |
| unique | url | 1048576 | baseline | 3 | 86 | 35 | +158 ms | 41 |
| unique | path | 300000 | ordered | 3 | 72 | 38 | +44 ms | 34 |
| unique | path | 300000 | baseline | 3 | 72 | 37 | +43 ms | 34 |
| unique | street | 212449 | ordered | 3 | 59 | 34 | +24 ms | 27 |
| unique | street | 212449 | baseline | 3 | 59 | 34 | +25 ms | 27 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
