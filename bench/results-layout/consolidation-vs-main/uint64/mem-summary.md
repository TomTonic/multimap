| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 117 | 20 | +75 ms | 57 |
| multi | u64 | 1048576 | baseline | 3 | 149 | 83 | +197 ms | 73 |
| multi | str | 1048576 | ordered | 3 | 129 | 31 | +136 ms | 62 |
| multi | str | 1048576 | baseline | 3 | 177 | 111 | +280 ms | 86 |
| multi | uuid | 1048576 | ordered | 3 | 147 | 31 | +124 ms | 71 |
| multi | uuid | 1048576 | baseline | 3 | 193 | 127 | +281 ms | 93 |
| multi | email | 1048576 | ordered | 3 | 137 | 30 | +123 ms | 66 |
| multi | email | 1048576 | baseline | 3 | 176 | 109 | +262 ms | 84 |
| multi | url | 1048576 | ordered | 3 | 162 | 39 | +179 ms | 79 |
| multi | url | 1048576 | baseline | 3 | 230 | 126 | +375 ms | 112 |
| multi | path | 300000 | ordered | 3 | 149 | 41 | +45 ms | 73 |
| multi | path | 300000 | baseline | 3 | 233 | 127 | +90 ms | 114 |
| multi | street | 212449 | ordered | 3 | 82 | 35 | +26 ms | 38 |
| multi | street | 212449 | baseline | 3 | 119 | 104 | +44 ms | 57 |
| unique | u64 | 1048576 | ordered | 3 | 41 | 17 | +66 ms | 19 |
| unique | u64 | 1048576 | baseline | 3 | 73 | 81 | +167 ms | 35 |
| unique | str | 1048576 | ordered | 3 | 52 | 28 | +122 ms | 23 |
| unique | str | 1048576 | baseline | 3 | 102 | 110 | +242 ms | 49 |
| unique | uuid | 1048576 | ordered | 3 | 68 | 28 | +116 ms | 31 |
| unique | uuid | 1048576 | baseline | 3 | 117 | 125 | +248 ms | 55 |
| unique | email | 1048576 | ordered | 3 | 61 | 26 | +111 ms | 28 |
| unique | email | 1048576 | baseline | 3 | 100 | 108 | +224 ms | 47 |
| unique | url | 1048576 | ordered | 3 | 86 | 35 | +157 ms | 41 |
| unique | url | 1048576 | baseline | 3 | 155 | 125 | +327 ms | 75 |
| unique | path | 300000 | ordered | 3 | 72 | 37 | +41 ms | 34 |
| unique | path | 300000 | baseline | 3 | 157 | 126 | +82 ms | 76 |
| unique | street | 212449 | ordered | 3 | 59 | 34 | +22 ms | 26 |
| unique | street | 212449 | baseline | 3 | 96 | 104 | +42 ms | 45 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
