| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi-str | u64 | 1048576 | ordered | 3 | 250 | 243 | +465 ms | 123 |
| multi-str | u64 | 1048576 | baseline | 3 | 250 | 243 | +459 ms | 123 |
| multi-str | str | 1048576 | ordered | 3 | 260 | 253 | +547 ms | 127 |
| multi-str | str | 1048576 | baseline | 3 | 276 | 269 | +576 ms | 135 |
| multi-str | uuid | 1048576 | ordered | 3 | 277 | 270 | +558 ms | 135 |
| multi-str | uuid | 1048576 | baseline | 3 | 293 | 286 | +600 ms | 143 |
| multi-str | email | 1048576 | ordered | 3 | 273 | 267 | +552 ms | 133 |
| multi-str | email | 1048576 | baseline | 3 | 275 | 269 | +554 ms | 134 |
| multi-str | url | 1048576 | ordered | 3 | 294 | 287 | +644 ms | 144 |
| multi-str | url | 1048576 | baseline | 3 | 329 | 286 | +694 ms | 161 |
| multi-str | path | 300000 | ordered | 3 | 282 | 275 | +148 ms | 138 |
| multi-str | path | 300000 | baseline | 3 | 332 | 287 | +158 ms | 162 |
| multi-str | street | 212449 | ordered | 3 | 169 | 168 | +64 ms | 81 |
| multi-str | street | 212449 | baseline | 3 | 173 | 173 | +65 ms | 83 |
| unique-str | u64 | 1048576 | ordered | 3 | 105 | 106 | +257 ms | 51 |
| unique-str | u64 | 1048576 | baseline | 3 | 105 | 106 | +259 ms | 51 |
| unique-str | str | 1048576 | ordered | 3 | 115 | 116 | +335 ms | 55 |
| unique-str | str | 1048576 | baseline | 3 | 131 | 132 | +345 ms | 63 |
| unique-str | uuid | 1048576 | ordered | 3 | 132 | 134 | +333 ms | 63 |
| unique-str | uuid | 1048576 | baseline | 3 | 148 | 149 | +368 ms | 71 |
| unique-str | email | 1048576 | ordered | 3 | 128 | 129 | +318 ms | 61 |
| unique-str | email | 1048576 | baseline | 3 | 130 | 132 | +322 ms | 62 |
| unique-str | url | 1048576 | ordered | 3 | 149 | 150 | +399 ms | 72 |
| unique-str | url | 1048576 | baseline | 3 | 184 | 148 | +443 ms | 89 |
| unique-str | path | 300000 | ordered | 3 | 137 | 138 | +87 ms | 67 |
| unique-str | path | 300000 | baseline | 3 | 188 | 151 | +97 ms | 91 |
| unique-str | street | 212449 | ordered | 3 | 123 | 126 | +49 ms | 59 |
| unique-str | street | 212449 | baseline | 3 | 127 | 130 | +49 ms | 61 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
