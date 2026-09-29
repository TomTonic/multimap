| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | u64 | 1048576 | ordered | 3 | 250 | 243 | +461 ms | 123 |
| multi-str | u64 | 1048576 | hashed | 3 | 262 | 256 | +349 ms | 157 |
| multi-str | u64 | 1048576 | btree-sets | 3 | 452 | 421 | +801 ms | 226 |
| multi-str | u64 | 1048576 | map-sets | 3 | 468 | 437 | +748 ms | 261 |
| multi-str | str | 1048576 | ordered | 3 | 276 | 270 | +568 ms | 135 |
| multi-str | str | 1048576 | hashed | 3 | 282 | 256 | +368 ms | 163 |
| multi-str | str | 1048576 | btree-sets | 3 | 472 | 421 | +818 ms | 232 |
| multi-str | str | 1048576 | map-sets | 3 | 488 | 436 | +765 ms | 266 |
| multi-str | uuid | 1048576 | ordered | 3 | 293 | 286 | +598 ms | 143 |
| multi-str | uuid | 1048576 | hashed | 3 | 302 | 256 | +369 ms | 173 |
| multi-str | uuid | 1048576 | btree-sets | 3 | 492 | 421 | +828 ms | 242 |
| multi-str | uuid | 1048576 | map-sets | 3 | 508 | 435 | +762 ms | 277 |
| multi-str | email | 1048576 | ordered | 3 | 275 | 269 | +561 ms | 134 |
| multi-str | email | 1048576 | hashed | 3 | 283 | 256 | +370 ms | 164 |
| multi-str | email | 1048576 | btree-sets | 3 | 473 | 423 | +824 ms | 233 |
| multi-str | email | 1048576 | map-sets | 3 | 489 | 436 | +768 ms | 267 |
| multi-str | url | 1048576 | ordered | 3 | 329 | 286 | +674 ms | 161 |
| multi-str | url | 1048576 | hashed | 3 | 325 | 256 | +389 ms | 185 |
| multi-str | url | 1048576 | btree-sets | 3 | 514 | 418 | +835 ms | 254 |
| multi-str | url | 1048576 | map-sets | 3 | 531 | 435 | +806 ms | 288 |
| multi-str | path | 300000 | ordered | 3 | 332 | 287 | +157 ms | 162 |
| multi-str | path | 300000 | hashed | 3 | 319 | 249 | +95 ms | 178 |
| multi-str | path | 300000 | btree-sets | 3 | 515 | 436 | +173 ms | 254 |
| multi-str | path | 300000 | map-sets | 3 | 525 | 446 | +161 ms | 281 |
| multi-str | street | 212449 | ordered | 3 | 173 | 173 | +65 ms | 82 |
| multi-str | street | 212449 | hashed | 3 | 153 | 141 | +50 ms | 88 |
| multi-str | street | 212449 | btree-sets | 3 | 359 | 336 | +92 ms | 175 |
| multi-str | street | 212449 | map-sets | 3 | 355 | 341 | +74 ms | 189 |
| unique-str | u64 | 1048576 | ordered | 3 | 105 | 106 | +251 ms | 51 |
| unique-str | u64 | 1048576 | btree-map | 3 | 48 | 48 | +85 ms | 25 |
| unique-str | str | 1048576 | ordered | 3 | 131 | 132 | +347 ms | 63 |
| unique-str | str | 1048576 | btree-map | 3 | 68 | 48 | +115 ms | 31 |
| unique-str | uuid | 1048576 | ordered | 3 | 148 | 149 | +377 ms | 71 |
| unique-str | uuid | 1048576 | btree-map | 3 | 88 | 48 | +115 ms | 41 |
| unique-str | email | 1048576 | ordered | 3 | 130 | 132 | +318 ms | 62 |
| unique-str | email | 1048576 | btree-map | 3 | 69 | 48 | +111 ms | 32 |
| unique-str | url | 1048576 | ordered | 3 | 184 | 148 | +444 ms | 89 |
| unique-str | url | 1048576 | btree-map | 3 | 111 | 49 | +122 ms | 52 |
| unique-str | path | 300000 | ordered | 3 | 188 | 151 | +97 ms | 91 |
| unique-str | path | 300000 | btree-map | 3 | 112 | 48 | +34 ms | 53 |
| unique-str | street | 212449 | ordered | 3 | 127 | 130 | +50 ms | 60 |
| unique-str | street | 212449 | btree-map | 3 | 58 | 49 | +23 ms | 26 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
