| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | u64 | 1048576 | ordered | 3 | 186 | 180 | +326 ms | 92 |
| multi-str | u64 | 1048576 | baseline | 3 | 251 | 244 | +430 ms | 124 |
| multi-str | str | 1048576 | ordered | 3 | 201 | 192 | +424 ms | 98 |
| multi-str | str | 1048576 | baseline | 3 | 279 | 273 | +523 ms | 137 |
| multi-str | uuid | 1048576 | ordered | 3 | 229 | 216 | +422 ms | 111 |
| multi-str | uuid | 1048576 | baseline | 3 | 294 | 288 | +557 ms | 144 |
| multi-str | email | 1048576 | ordered | 3 | 211 | 203 | +424 ms | 103 |
| multi-str | email | 1048576 | baseline | 3 | 277 | 271 | +520 ms | 135 |
| multi-str | url | 1048576 | ordered | 3 | 244 | 235 | +531 ms | 120 |
| multi-str | url | 1048576 | baseline | 3 | 332 | 287 | +625 ms | 163 |
| multi-str | path | 300000 | ordered | 3 | 224 | 215 | +129 ms | 110 |
| multi-str | path | 300000 | baseline | 3 | 334 | 289 | +153 ms | 164 |
| multi-str | street | 212449 | ordered | 3 | 108 | 108 | +53 ms | 52 |
| multi-str | street | 212449 | baseline | 3 | 174 | 174 | +65 ms | 84 |
| unique-str | u64 | 1048576 | ordered | 3 | 41 | 44 | +138 ms | 19 |
| unique-str | u64 | 1048576 | baseline | 3 | 105 | 107 | +234 ms | 51 |
| unique-str | str | 1048576 | ordered | 3 | 56 | 54 | +226 ms | 26 |
| unique-str | str | 1048576 | baseline | 3 | 134 | 134 | +304 ms | 65 |
| unique-str | uuid | 1048576 | ordered | 3 | 84 | 78 | +238 ms | 39 |
| unique-str | uuid | 1048576 | baseline | 3 | 149 | 150 | +335 ms | 71 |
| unique-str | email | 1048576 | ordered | 3 | 67 | 66 | +214 ms | 31 |
| unique-str | email | 1048576 | baseline | 3 | 132 | 133 | +294 ms | 63 |
| unique-str | url | 1048576 | ordered | 3 | 100 | 97 | +313 ms | 48 |
| unique-str | url | 1048576 | baseline | 3 | 187 | 149 | +411 ms | 91 |
| unique-str | path | 300000 | ordered | 3 | 80 | 78 | +68 ms | 38 |
| unique-str | path | 300000 | baseline | 3 | 189 | 152 | +97 ms | 92 |
| unique-str | street | 212449 | ordered | 3 | 59 | 63 | +36 ms | 27 |
| unique-str | street | 212449 | baseline | 3 | 128 | 131 | +52 ms | 61 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
