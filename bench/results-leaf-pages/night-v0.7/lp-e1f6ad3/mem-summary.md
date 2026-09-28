| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 149 | 83 | +220 ms | 72 |
| multi | u64 | 1048576 | baseline | 3 | 149 | 83 | +220 ms | 72 |
| multi | str | 1048576 | ordered | 3 | 177 | 111 | +314 ms | 86 |
| multi | str | 1048576 | baseline | 3 | 177 | 111 | +307 ms | 86 |
| multi | uuid | 1048576 | ordered | 3 | 192 | 127 | +319 ms | 93 |
| multi | uuid | 1048576 | baseline | 3 | 192 | 127 | +317 ms | 93 |
| multi | email | 1048576 | ordered | 3 | 175 | 109 | +283 ms | 84 |
| multi | email | 1048576 | baseline | 3 | 175 | 109 | +289 ms | 84 |
| multi | url | 1048576 | ordered | 3 | 229 | 118 | +398 ms | 111 |
| multi | url | 1048576 | baseline | 3 | 229 | 118 | +381 ms | 111 |
| multi | path | 300000 | ordered | 3 | 232 | 127 | +93 ms | 113 |
| multi | path | 300000 | baseline | 3 | 232 | 127 | +96 ms | 113 |
| multi | street | 212449 | ordered | 3 | 120 | 104 | +45 ms | 56 |
| multi | street | 212449 | baseline | 3 | 120 | 104 | +47 ms | 56 |
| unique | u64 | 1048576 | ordered | 3 | 73 | 81 | +189 ms | 35 |
| unique | u64 | 1048576 | baseline | 3 | 16 | 1 | +3 ms | 9 |
| unique | str | 1048576 | ordered | 3 | 101 | 109 | +273 ms | 48 |
| unique | str | 1048576 | baseline | 3 | 44 | 33 | +47 ms | 27 |
| unique | uuid | 1048576 | ordered | 3 | 117 | 125 | +285 ms | 55 |
| unique | uuid | 1048576 | baseline | 3 | 95 | 39 | +131 ms | 54 |
| unique | email | 1048576 | ordered | 3 | 100 | 108 | +257 ms | 47 |
| unique | email | 1048576 | baseline | 3 | 72 | 40 | +116 ms | 40 |
| unique | url | 1048576 | ordered | 3 | 154 | 117 | +336 ms | 74 |
| unique | url | 1048576 | baseline | 3 | 117 | 43 | +117 ms | 67 |
| unique | path | 300000 | ordered | 3 | 158 | 126 | +86 ms | 76 |
| unique | path | 300000 | baseline | 3 | 105 | 38 | +35 ms | 62 |
| unique | street | 212449 | ordered | 3 | 96 | 104 | +46 ms | 45 |
| unique | street | 212449 | baseline | 3 | 48 | 40 | +14 ms | 31 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
