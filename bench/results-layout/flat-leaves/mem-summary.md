| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 1048576 | ordered | 3 | 116 | 20 | +83 ms | 57 |
| multi | u64 | 1048576 | baseline | 3 | 148 | 82 | +209 ms | 72 |
| multi | str | 1048576 | ordered | 3 | 140 | 31 | +144 ms | 67 |
| multi | str | 1048576 | baseline | 3 | 175 | 109 | +317 ms | 85 |
| multi | uuid | 1048576 | ordered | 3 | 147 | 32 | +139 ms | 71 |
| multi | uuid | 1048576 | baseline | 3 | 191 | 125 | +315 ms | 92 |
| multi | email | 1048576 | ordered | 3 | 139 | 30 | +130 ms | 66 |
| multi | email | 1048576 | baseline | 3 | 174 | 108 | +281 ms | 84 |
| multi | url | 1048576 | ordered | 3 | 190 | 34 | +162 ms | 92 |
| multi | url | 1048576 | baseline | 3 | 228 | 117 | +365 ms | 111 |
| multi | path | 300000 | ordered | 3 | 196 | 41 | +43 ms | 95 |
| multi | path | 300000 | baseline | 3 | 231 | 126 | +89 ms | 112 |
| multi | street | 212449 | ordered | 3 | 85 | 35 | +24 ms | 39 |
| multi | street | 212449 | baseline | 3 | 118 | 103 | +41 ms | 56 |
| unique | u64 | 1048576 | ordered | 3 | 41 | 17 | +73 ms | 19 |
| unique | u64 | 1048576 | baseline | 3 | 73 | 81 | +180 ms | 35 |
| unique | str | 1048576 | ordered | 3 | 66 | 27 | +131 ms | 30 |
| unique | str | 1048576 | baseline | 3 | 99 | 107 | +270 ms | 47 |
| unique | uuid | 1048576 | ordered | 3 | 68 | 28 | +124 ms | 31 |
| unique | uuid | 1048576 | baseline | 3 | 116 | 124 | +268 ms | 55 |
| unique | email | 1048576 | ordered | 3 | 65 | 26 | +117 ms | 29 |
| unique | email | 1048576 | baseline | 3 | 98 | 106 | +239 ms | 46 |
| unique | url | 1048576 | ordered | 3 | 119 | 31 | +143 ms | 57 |
| unique | url | 1048576 | baseline | 3 | 153 | 115 | +322 ms | 74 |
| unique | path | 300000 | ordered | 3 | 121 | 37 | +39 ms | 58 |
| unique | path | 300000 | baseline | 3 | 156 | 124 | +80 ms | 75 |
| unique | street | 212449 | ordered | 3 | 60 | 34 | +23 ms | 27 |
| unique | street | 212449 | baseline | 3 | 95 | 103 | +42 ms | 45 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
