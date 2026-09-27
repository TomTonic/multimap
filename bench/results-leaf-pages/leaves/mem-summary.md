| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 5 | 149 | 83 | +217 ms | 72 |
| multi | u64 | 1048576 | baseline | 5 | 149 | 83 | +218 ms | 72 |
| multi | str | 1048576 | ordered | 5 | 225 | 159 | +344 ms | 110 |
| multi | str | 1048576 | baseline | 5 | 177 | 111 | +307 ms | 86 |
| multi | uuid | 1048576 | ordered | 5 | 224 | 159 | +358 ms | 109 |
| multi | uuid | 1048576 | baseline | 5 | 192 | 127 | +318 ms | 93 |
| multi | email | 1048576 | ordered | 5 | 222 | 157 | +316 ms | 108 |
| multi | email | 1048576 | baseline | 5 | 175 | 109 | +285 ms | 84 |
| multi | url | 1048576 | ordered | 5 | 227 | 161 | +345 ms | 110 |
| multi | url | 1048576 | baseline | 5 | 229 | 118 | +375 ms | 111 |
| multi | path | 300000 | ordered | 5 | 242 | 156 | +95 ms | 117 |
| multi | path | 300000 | baseline | 5 | 232 | 127 | +92 ms | 113 |
| multi | street | 212449 | ordered | 5 | 132 | 116 | +47 ms | 62 |
| multi | street | 212449 | baseline | 5 | 119 | 104 | +46 ms | 56 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
