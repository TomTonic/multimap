| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 117 | 20 | +79 ms | 57 |
| multi | u64 | 1048576 | baseline | 3 | 117 | 20 | +77 ms | 57 |
| multi | str | 1048576 | ordered | 3 | 129 | 31 | +137 ms | 62 |
| multi | str | 1048576 | baseline | 3 | 129 | 31 | +136 ms | 62 |
| multi | uuid | 1048576 | ordered | 3 | 147 | 31 | +126 ms | 71 |
| multi | uuid | 1048576 | baseline | 3 | 147 | 31 | +124 ms | 71 |
| multi | email | 1048576 | ordered | 3 | 137 | 30 | +121 ms | 66 |
| multi | email | 1048576 | baseline | 3 | 137 | 30 | +122 ms | 66 |
| multi | url | 1048576 | ordered | 3 | 162 | 39 | +175 ms | 79 |
| multi | url | 1048576 | baseline | 3 | 162 | 39 | +180 ms | 79 |
| multi | path | 300000 | ordered | 3 | 149 | 41 | +49 ms | 73 |
| multi | path | 300000 | baseline | 3 | 149 | 41 | +45 ms | 72 |
| multi | street | 212449 | ordered | 3 | 82 | 35 | +26 ms | 38 |
| multi | street | 212449 | baseline | 3 | 82 | 35 | +26 ms | 38 |
| unique | u64 | 1048576 | ordered | 3 | 16 | 1 | +4 ms | 9 |
| unique | u64 | 1048576 | baseline | 3 | 41 | 17 | +67 ms | 19 |
| unique | str | 1048576 | ordered | 3 | 44 | 31 | +44 ms | 27 |
| unique | str | 1048576 | baseline | 3 | 52 | 28 | +124 ms | 23 |
| unique | uuid | 1048576 | ordered | 3 | 95 | 40 | +116 ms | 54 |
| unique | uuid | 1048576 | baseline | 3 | 68 | 28 | +115 ms | 31 |
| unique | email | 1048576 | ordered | 3 | 72 | 40 | +109 ms | 40 |
| unique | email | 1048576 | baseline | 3 | 61 | 26 | +110 ms | 28 |
| unique | url | 1048576 | ordered | 3 | 112 | 33 | +105 ms | 65 |
| unique | url | 1048576 | baseline | 3 | 86 | 37 | +152 ms | 41 |
| unique | path | 300000 | ordered | 3 | 105 | 38 | +35 ms | 62 |
| unique | path | 300000 | baseline | 3 | 72 | 37 | +42 ms | 34 |
| unique | street | 212449 | ordered | 3 | 48 | 37 | +14 ms | 31 |
| unique | street | 212449 | baseline | 3 | 59 | 34 | +26 ms | 27 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
