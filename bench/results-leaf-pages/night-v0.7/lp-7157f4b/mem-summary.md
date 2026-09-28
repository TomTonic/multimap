| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 262144 | ordered | 1 | 154 | 89 | +39 ms | 75 |
| multi | u64 | 262144 | baseline | 1 | 115 | 56 | +13 ms | 68 |
| multi | str | 262144 | ordered | 1 | 176 | 110 | +55 ms | 84 |
| multi | str | 262144 | baseline | 1 | 131 | 39 | +28 ms | 75 |
| multi | uuid | 262144 | ordered | 1 | 193 | 127 | +58 ms | 94 |
| multi | uuid | 262144 | baseline | 1 | 153 | 53 | +27 ms | 90 |
| multi | path | 262144 | ordered | 1 | 233 | 127 | +78 ms | 113 |
| multi | path | 262144 | baseline | 1 | 165 | 45 | +39 ms | 102 |
| multi | street | 212449 | ordered | 1 | 120 | 104 | +47 ms | 56 |
| multi | street | 212449 | baseline | 1 | 62 | 21 | +14 ms | 38 |
| unique | u64 | 262144 | ordered | 1 | 79 | 87 | +34 ms | 38 |
| unique | u64 | 262144 | baseline | 1 | 21 | 2 | +5 ms | 14 |
| unique | str | 262144 | ordered | 1 | 101 | 109 | +46 ms | 47 |
| unique | str | 262144 | baseline | 1 | 43 | 28 | +14 ms | 31 |
| unique | uuid | 262144 | ordered | 1 | 118 | 126 | +51 ms | 57 |
| unique | uuid | 262144 | baseline | 1 | 65 | 49 | +16 ms | 45 |
| unique | path | 262144 | ordered | 1 | 158 | 126 | +71 ms | 76 |
| unique | path | 262144 | baseline | 1 | 76 | 41 | +25 ms | 57 |
| unique | street | 212449 | ordered | 1 | 96 | 104 | +44 ms | 45 |
| unique | street | 212449 | baseline | 1 | 34 | 24 | +10 ms | 24 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
