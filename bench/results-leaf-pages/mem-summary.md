| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 5 | 140 | 52 | +154 ms | 72 |
| multi | u64 | 1048576 | baseline | 5 | 149 | 83 | +219 ms | 72 |
| multi | str | 1048576 | ordered | 5 | 194 | 89 | +386 ms | 96 |
| multi | str | 1048576 | baseline | 5 | 177 | 111 | +306 ms | 86 |
| multi | uuid | 1048576 | ordered | 5 | 223 | 92 | +400 ms | 111 |
| multi | uuid | 1048576 | baseline | 5 | 192 | 127 | +326 ms | 93 |
| multi | email | 1048576 | ordered | 5 | 201 | 90 | +417 ms | 99 |
| multi | email | 1048576 | baseline | 5 | 175 | 109 | +290 ms | 84 |
| multi | url | 1048576 | ordered | 5 | 251 | 99 | +425 ms | 127 |
| multi | url | 1048576 | baseline | 5 | 229 | 118 | +375 ms | 111 |
| multi | path | 300000 | ordered | 5 | 259 | 111 | +108 ms | 130 |
| multi | path | 300000 | baseline | 5 | 232 | 127 | +92 ms | 113 |
| multi | street | 212449 | ordered | 5 | 116 | 73 | +44 ms | 60 |
| multi | street | 212449 | baseline | 5 | 120 | 104 | +45 ms | 56 |
| unique | u64 | 1048576 | ordered | 5 | 16 | 1 | +2 ms | 9 |
| unique | u64 | 1048576 | baseline | 5 | 73 | 81 | +185 ms | 35 |
| unique | str | 1048576 | ordered | 5 | 44 | 32 | +51 ms | 27 |
| unique | str | 1048576 | baseline | 5 | 101 | 109 | +266 ms | 48 |
| unique | uuid | 1048576 | ordered | 5 | 95 | 39 | +128 ms | 54 |
| unique | uuid | 1048576 | baseline | 5 | 117 | 125 | +275 ms | 55 |
| unique | email | 1048576 | ordered | 5 | 72 | 40 | +114 ms | 40 |
| unique | email | 1048576 | baseline | 5 | 100 | 108 | +248 ms | 47 |
| unique | url | 1048576 | ordered | 5 | 117 | 41 | +126 ms | 67 |
| unique | url | 1048576 | baseline | 5 | 154 | 117 | +331 ms | 74 |
| unique | path | 300000 | ordered | 5 | 105 | 41 | +36 ms | 62 |
| unique | path | 300000 | baseline | 5 | 157 | 126 | +84 ms | 76 |
| unique | street | 212449 | ordered | 5 | 48 | 39 | +15 ms | 31 |
| unique | street | 212449 | baseline | 5 | 96 | 104 | +44 ms | 45 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
