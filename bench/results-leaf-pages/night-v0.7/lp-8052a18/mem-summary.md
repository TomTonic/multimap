| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 262144 | ordered | 1 | 154 | 89 | +39 ms | 75 |
| multi | u64 | 262144 | baseline | 1 | 152 | 63 | +28 ms | 77 |
| multi | str | 262144 | ordered | 1 | 176 | 110 | +54 ms | 84 |
| multi | str | 262144 | baseline | 1 | 196 | 96 | +69 ms | 96 |
| multi | uuid | 262144 | ordered | 1 | 193 | 127 | +59 ms | 94 |
| multi | uuid | 262144 | baseline | 1 | 221 | 93 | +70 ms | 113 |
| multi | path | 262144 | ordered | 1 | 233 | 127 | +79 ms | 113 |
| multi | path | 262144 | baseline | 1 | 259 | 110 | +93 ms | 130 |
| multi | street | 212449 | ordered | 1 | 119 | 104 | +43 ms | 56 |
| multi | street | 212449 | baseline | 1 | 116 | 72 | +45 ms | 60 |
| unique | u64 | 262144 | ordered | 1 | 79 | 87 | +33 ms | 38 |
| unique | u64 | 262144 | baseline | 1 | 16 | 1 | +0 ms | 10 |
| unique | str | 262144 | ordered | 1 | 101 | 109 | +49 ms | 47 |
| unique | str | 262144 | baseline | 1 | 45 | 41 | +11 ms | 27 |
| unique | uuid | 262144 | ordered | 1 | 118 | 126 | +51 ms | 57 |
| unique | uuid | 262144 | baseline | 1 | 92 | 46 | +28 ms | 51 |
| unique | path | 262144 | ordered | 1 | 158 | 126 | +71 ms | 76 |
| unique | path | 262144 | baseline | 1 | 106 | 42 | +31 ms | 63 |
| unique | street | 212449 | ordered | 1 | 96 | 104 | +40 ms | 45 |
| unique | street | 212449 | baseline | 1 | 48 | 37 | +12 ms | 31 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
