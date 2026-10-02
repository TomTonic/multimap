| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | u64 | 1048576 | ordered | 3 | 190 | 184 | +327 ms | 94 |
| multi-str | u64 | 1048576 | baseline | 3 | 190 | 184 | +326 ms | 94 |
| multi-str | str | 1048576 | ordered | 3 | 205 | 197 | +423 ms | 100 |
| multi-str | str | 1048576 | baseline | 3 | 205 | 196 | +415 ms | 100 |
| multi-str | uuid | 1048576 | ordered | 3 | 232 | 218 | +418 ms | 113 |
| multi-str | uuid | 1048576 | baseline | 3 | 232 | 218 | +417 ms | 113 |
| multi-str | email | 1048576 | ordered | 3 | 215 | 207 | +410 ms | 104 |
| multi-str | email | 1048576 | baseline | 3 | 215 | 207 | +407 ms | 104 |
| multi-str | url | 1048576 | ordered | 3 | 247 | 238 | +516 ms | 122 |
| multi-str | url | 1048576 | baseline | 3 | 247 | 238 | +504 ms | 122 |
| multi-str | path | 300000 | ordered | 3 | 227 | 218 | +125 ms | 112 |
| multi-str | path | 300000 | baseline | 3 | 227 | 218 | +128 ms | 112 |
| multi-str | street | 212449 | ordered | 3 | 111 | 110 | +55 ms | 53 |
| multi-str | street | 212449 | baseline | 3 | 111 | 111 | +52 ms | 53 |
| unique-str | u64 | 1048576 | ordered | 3 | 41 | 44 | +141 ms | 19 |
| unique-str | u64 | 1048576 | baseline | 3 | 41 | 43 | +137 ms | 19 |
| unique-str | str | 1048576 | ordered | 3 | 56 | 55 | +234 ms | 26 |
| unique-str | str | 1048576 | baseline | 3 | 56 | 54 | +226 ms | 26 |
| unique-str | uuid | 1048576 | ordered | 3 | 84 | 78 | +234 ms | 39 |
| unique-str | uuid | 1048576 | baseline | 3 | 84 | 78 | +235 ms | 39 |
| unique-str | email | 1048576 | ordered | 3 | 67 | 66 | +220 ms | 31 |
| unique-str | email | 1048576 | baseline | 3 | 67 | 66 | +213 ms | 31 |
| unique-str | url | 1048576 | ordered | 3 | 100 | 97 | +312 ms | 48 |
| unique-str | url | 1048576 | baseline | 3 | 100 | 97 | +312 ms | 48 |
| unique-str | path | 300000 | ordered | 3 | 80 | 77 | +68 ms | 38 |
| unique-str | path | 300000 | baseline | 3 | 80 | 79 | +69 ms | 38 |
| unique-str | street | 212449 | ordered | 3 | 59 | 63 | +37 ms | 27 |
| unique-str | street | 212449 | baseline | 3 | 59 | 62 | +38 ms | 27 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
