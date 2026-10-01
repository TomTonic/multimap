| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | u64 | 1048576 | ordered | 3 | 250 | 244 | +423 ms | 124 |
| multi-str | u64 | 1048576 | baseline | 3 | 251 | 245 | +429 ms | 124 |
| multi-str | str | 1048576 | ordered | 3 | 261 | 254 | +502 ms | 128 |
| multi-str | str | 1048576 | baseline | 3 | 279 | 272 | +518 ms | 137 |
| multi-str | uuid | 1048576 | ordered | 3 | 277 | 270 | +515 ms | 136 |
| multi-str | uuid | 1048576 | baseline | 3 | 294 | 287 | +554 ms | 144 |
| multi-str | email | 1048576 | ordered | 3 | 273 | 267 | +505 ms | 134 |
| multi-str | email | 1048576 | baseline | 3 | 277 | 271 | +512 ms | 135 |
| multi-str | url | 1048576 | ordered | 3 | 296 | 289 | +589 ms | 146 |
| multi-str | url | 1048576 | baseline | 3 | 332 | 287 | +626 ms | 163 |
| multi-str | path | 300000 | ordered | 3 | 282 | 275 | +143 ms | 139 |
| multi-str | path | 300000 | baseline | 3 | 334 | 289 | +156 ms | 164 |
| multi-str | street | 212449 | ordered | 3 | 168 | 169 | +63 ms | 82 |
| multi-str | street | 212449 | baseline | 3 | 174 | 174 | +66 ms | 84 |
| unique-str | u64 | 1048576 | ordered | 3 | 105 | 106 | +231 ms | 51 |
| unique-str | u64 | 1048576 | baseline | 3 | 105 | 107 | +231 ms | 51 |
| unique-str | str | 1048576 | ordered | 3 | 116 | 116 | +306 ms | 55 |
| unique-str | str | 1048576 | baseline | 3 | 134 | 135 | +307 ms | 65 |
| unique-str | uuid | 1048576 | ordered | 3 | 132 | 134 | +306 ms | 63 |
| unique-str | uuid | 1048576 | baseline | 3 | 149 | 150 | +344 ms | 71 |
| unique-str | email | 1048576 | ordered | 3 | 128 | 129 | +294 ms | 61 |
| unique-str | email | 1048576 | baseline | 3 | 132 | 133 | +302 ms | 63 |
| unique-str | url | 1048576 | ordered | 3 | 150 | 151 | +370 ms | 73 |
| unique-str | url | 1048576 | baseline | 3 | 187 | 149 | +408 ms | 91 |
| unique-str | path | 300000 | ordered | 3 | 137 | 138 | +84 ms | 67 |
| unique-str | path | 300000 | baseline | 3 | 189 | 152 | +96 ms | 92 |
| unique-str | street | 212449 | ordered | 3 | 123 | 125 | +48 ms | 59 |
| unique-str | street | 212449 | baseline | 3 | 128 | 131 | +53 ms | 61 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
