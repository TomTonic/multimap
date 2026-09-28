| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 149 | 83 | +215 ms | 72 |
| multi | u64 | 1048576 | hashed | 3 | 177 | 101 | +218 ms | 115 |
| multi | u64 | 1048576 | btree-sets | 3 | 343 | 74 | +351 ms | 172 |
| multi | u64 | 1048576 | map-sets | 3 | 360 | 85 | +316 ms | 206 |
| multi | str | 1048576 | ordered | 3 | 177 | 111 | +305 ms | 86 |
| multi | str | 1048576 | hashed | 3 | 196 | 101 | +236 ms | 121 |
| multi | str | 1048576 | btree-sets | 3 | 363 | 72 | +364 ms | 178 |
| multi | str | 1048576 | map-sets | 3 | 379 | 87 | +340 ms | 212 |
| multi | uuid | 1048576 | ordered | 3 | 192 | 127 | +325 ms | 93 |
| multi | uuid | 1048576 | hashed | 3 | 217 | 101 | +229 ms | 131 |
| multi | uuid | 1048576 | btree-sets | 3 | 383 | 74 | +359 ms | 188 |
| multi | uuid | 1048576 | map-sets | 3 | 400 | 86 | +327 ms | 222 |
| multi | email | 1048576 | ordered | 3 | 175 | 109 | +285 ms | 84 |
| multi | email | 1048576 | hashed | 3 | 197 | 101 | +234 ms | 121 |
| multi | email | 1048576 | btree-sets | 3 | 364 | 74 | +363 ms | 178 |
| multi | email | 1048576 | map-sets | 3 | 380 | 87 | +334 ms | 213 |
| multi | url | 1048576 | ordered | 3 | 229 | 118 | +370 ms | 111 |
| multi | url | 1048576 | hashed | 3 | 242 | 101 | +233 ms | 143 |
| multi | url | 1048576 | btree-sets | 3 | 408 | 73 | +369 ms | 200 |
| multi | url | 1048576 | map-sets | 3 | 425 | 85 | +341 ms | 235 |
| multi | path | 300000 | ordered | 3 | 232 | 127 | +90 ms | 113 |
| multi | path | 300000 | hashed | 3 | 234 | 96 | +59 ms | 136 |
| multi | path | 300000 | btree-sets | 3 | 406 | 82 | +79 ms | 200 |
| multi | path | 300000 | map-sets | 3 | 417 | 91 | +77 ms | 227 |
| multi | street | 212449 | ordered | 3 | 119 | 104 | +45 ms | 56 |
| multi | street | 212449 | hashed | 3 | 114 | 81 | +38 ms | 70 |
| multi | street | 212449 | btree-sets | 3 | 277 | 82 | +50 ms | 134 |
| multi | street | 212449 | map-sets | 3 | 273 | 79 | +44 ms | 148 |
| unique | u64 | 1048576 | ordered | 3 | 73 | 81 | +179 ms | 35 |
| unique | u64 | 1048576 | btree-map | 3 | 37 | 37 | +59 ms | 19 |
| unique | str | 1048576 | ordered | 3 | 101 | 109 | +264 ms | 48 |
| unique | str | 1048576 | btree-map | 3 | 56 | 37 | +76 ms | 25 |
| unique | uuid | 1048576 | ordered | 3 | 117 | 125 | +276 ms | 55 |
| unique | uuid | 1048576 | btree-map | 3 | 77 | 37 | +77 ms | 35 |
| unique | email | 1048576 | ordered | 3 | 100 | 108 | +243 ms | 47 |
| unique | email | 1048576 | btree-map | 3 | 57 | 37 | +73 ms | 25 |
| unique | url | 1048576 | ordered | 3 | 154 | 117 | +325 ms | 74 |
| unique | url | 1048576 | btree-map | 3 | 102 | 37 | +78 ms | 48 |
| unique | path | 300000 | ordered | 3 | 157 | 126 | +80 ms | 76 |
| unique | path | 300000 | btree-map | 3 | 101 | 36 | +22 ms | 47 |
| unique | street | 212449 | ordered | 3 | 96 | 104 | +43 ms | 45 |
| unique | street | 212449 | btree-map | 3 | 46 | 37 | +15 ms | 20 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
