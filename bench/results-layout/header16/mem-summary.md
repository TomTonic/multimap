| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 148 | 82 | +212 ms | 72 |
| multi | u64 | 1048576 | baseline | 3 | 149 | 83 | +212 ms | 72 |
| multi | str | 1048576 | ordered | 3 | 175 | 109 | +308 ms | 85 |
| multi | str | 1048576 | baseline | 3 | 177 | 111 | +302 ms | 86 |
| multi | uuid | 1048576 | ordered | 3 | 191 | 125 | +307 ms | 92 |
| multi | uuid | 1048576 | baseline | 3 | 192 | 127 | +317 ms | 93 |
| multi | email | 1048576 | ordered | 3 | 174 | 108 | +279 ms | 84 |
| multi | email | 1048576 | baseline | 3 | 175 | 109 | +285 ms | 84 |
| multi | url | 1048576 | ordered | 3 | 228 | 117 | +359 ms | 111 |
| multi | url | 1048576 | baseline | 3 | 229 | 118 | +370 ms | 111 |
| multi | path | 300000 | ordered | 3 | 231 | 126 | +89 ms | 112 |
| multi | path | 300000 | baseline | 3 | 232 | 127 | +89 ms | 113 |
| multi | street | 212449 | ordered | 3 | 118 | 103 | +42 ms | 56 |
| multi | street | 212449 | baseline | 3 | 119 | 104 | +44 ms | 56 |
| unique | u64 | 1048576 | ordered | 3 | 73 | 81 | +179 ms | 35 |
| unique | u64 | 1048576 | baseline | 3 | 73 | 81 | +182 ms | 35 |
| unique | str | 1048576 | ordered | 3 | 99 | 107 | +271 ms | 47 |
| unique | str | 1048576 | baseline | 3 | 101 | 109 | +258 ms | 48 |
| unique | uuid | 1048576 | ordered | 3 | 116 | 124 | +268 ms | 55 |
| unique | uuid | 1048576 | baseline | 3 | 117 | 125 | +274 ms | 55 |
| unique | email | 1048576 | ordered | 3 | 98 | 106 | +241 ms | 46 |
| unique | email | 1048576 | baseline | 3 | 100 | 108 | +241 ms | 47 |
| unique | url | 1048576 | ordered | 3 | 153 | 115 | +321 ms | 74 |
| unique | url | 1048576 | baseline | 3 | 154 | 117 | +325 ms | 74 |
| unique | path | 300000 | ordered | 3 | 156 | 124 | +80 ms | 75 |
| unique | path | 300000 | baseline | 3 | 157 | 126 | +81 ms | 76 |
| unique | street | 212449 | ordered | 3 | 95 | 103 | +42 ms | 45 |
| unique | street | 212449 | baseline | 3 | 96 | 104 | +43 ms | 45 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
