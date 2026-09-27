| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 5 | 149 | 83 | +217 ms | 72 |
| multi | u64 | 1048576 | baseline | 5 | 149 | 83 | +219 ms | 72 |
| multi | str | 1048576 | ordered | 5 | 188 | 95 | +355 ms | 92 |
| multi | str | 1048576 | baseline | 5 | 177 | 111 | +307 ms | 86 |
| multi | uuid | 1048576 | ordered | 5 | 208 | 95 | +339 ms | 101 |
| multi | uuid | 1048576 | baseline | 5 | 192 | 127 | +322 ms | 93 |
| multi | email | 1048576 | ordered | 5 | 188 | 93 | +337 ms | 90 |
| multi | email | 1048576 | baseline | 5 | 175 | 109 | +288 ms | 84 |
| multi | url | 1048576 | ordered | 5 | 236 | 97 | +386 ms | 115 |
| multi | url | 1048576 | baseline | 5 | 229 | 118 | +375 ms | 111 |
| multi | path | 300000 | ordered | 5 | 241 | 104 | +91 ms | 117 |
| multi | path | 300000 | baseline | 5 | 232 | 127 | +92 ms | 113 |
| multi | street | 212449 | ordered | 5 | 122 | 100 | +47 ms | 57 |
| multi | street | 212449 | baseline | 5 | 120 | 104 | +45 ms | 56 |
| unique | u64 | 1048576 | ordered | 5 | 16 | 1 | +2 ms | 9 |
| unique | u64 | 1048576 | baseline | 5 | 73 | 81 | +185 ms | 35 |
| unique | str | 1048576 | ordered | 5 | 44 | 32 | +54 ms | 27 |
| unique | str | 1048576 | baseline | 5 | 101 | 109 | +266 ms | 48 |
| unique | uuid | 1048576 | ordered | 5 | 95 | 40 | +128 ms | 54 |
| unique | uuid | 1048576 | baseline | 5 | 117 | 125 | +275 ms | 55 |
| unique | email | 1048576 | ordered | 5 | 72 | 40 | +116 ms | 40 |
| unique | email | 1048576 | baseline | 5 | 100 | 108 | +247 ms | 47 |
| unique | url | 1048576 | ordered | 5 | 117 | 41 | +124 ms | 67 |
| unique | url | 1048576 | baseline | 5 | 154 | 117 | +327 ms | 74 |
| unique | path | 300000 | ordered | 5 | 105 | 41 | +35 ms | 62 |
| unique | path | 300000 | baseline | 5 | 158 | 126 | +83 ms | 76 |
| unique | street | 212449 | ordered | 5 | 48 | 39 | +15 ms | 31 |
| unique | street | 212449 | baseline | 5 | 96 | 104 | +44 ms | 45 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
