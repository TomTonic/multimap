| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 262144 | ordered | 1 | 154 | 89 | +37 ms | 75 |
| multi | u64 | 262144 | baseline | 1 | 115 | 56 | +12 ms | 68 |
| multi | str | 262144 | ordered | 1 | 176 | 110 | +53 ms | 84 |
| multi | str | 262144 | baseline | 1 | 131 | 38 | +26 ms | 75 |
| multi | uuid | 262144 | ordered | 1 | 193 | 127 | +58 ms | 94 |
| multi | uuid | 262144 | baseline | 1 | 153 | 56 | +27 ms | 90 |
| multi | path | 262144 | ordered | 1 | 233 | 127 | +74 ms | 113 |
| multi | path | 262144 | baseline | 1 | 165 | 49 | +34 ms | 102 |
| multi | street | 212449 | ordered | 1 | 119 | 104 | +43 ms | 56 |
| multi | street | 212449 | baseline | 1 | 62 | 23 | +12 ms | 38 |
| unique | u64 | 262144 | ordered | 1 | 79 | 87 | +32 ms | 38 |
| unique | u64 | 262144 | baseline | 1 | 21 | 2 | +4 ms | 14 |
| unique | str | 262144 | ordered | 1 | 101 | 109 | +47 ms | 47 |
| unique | str | 262144 | baseline | 1 | 43 | 30 | +13 ms | 31 |
| unique | uuid | 262144 | ordered | 1 | 118 | 126 | +49 ms | 57 |
| unique | uuid | 262144 | baseline | 1 | 65 | 48 | +15 ms | 45 |
| unique | path | 262144 | ordered | 1 | 157 | 126 | +67 ms | 76 |
| unique | path | 262144 | baseline | 1 | 76 | 47 | +22 ms | 57 |
| unique | street | 212449 | ordered | 1 | 96 | 104 | +42 ms | 45 |
| unique | street | 212449 | baseline | 1 | 34 | 23 | +11 ms | 24 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
