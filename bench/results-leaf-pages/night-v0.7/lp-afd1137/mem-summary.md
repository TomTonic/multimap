| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 262144 | ordered | 1 | 155 | 89 | +39 ms | 75 |
| multi | u64 | 262144 | baseline | 1 | 109 | 51 | +9 ms | 64 |
| multi | str | 262144 | ordered | 1 | 176 | 110 | +55 ms | 84 |
| multi | str | 262144 | baseline | 1 | 115 | 37 | +17 ms | 67 |
| multi | uuid | 262144 | ordered | 1 | 193 | 127 | +59 ms | 94 |
| multi | uuid | 262144 | baseline | 1 | 147 | 49 | +24 ms | 88 |
| multi | path | 262144 | ordered | 1 | 233 | 127 | +79 ms | 113 |
| multi | path | 262144 | baseline | 1 | 153 | 47 | +31 ms | 98 |
| multi | street | 212449 | ordered | 1 | 120 | 104 | +49 ms | 56 |
| multi | street | 212449 | baseline | 1 | 59 | 31 | +12 ms | 37 |
| unique | u64 | 262144 | ordered | 1 | 79 | 87 | +34 ms | 38 |
| unique | u64 | 262144 | baseline | 1 | 15 | 1 | +1 ms | 10 |
| unique | str | 262144 | ordered | 1 | 101 | 109 | +47 ms | 47 |
| unique | str | 262144 | baseline | 1 | 44 | 43 | +11 ms | 28 |
| unique | uuid | 262144 | ordered | 1 | 118 | 126 | +50 ms | 57 |
| unique | uuid | 262144 | baseline | 1 | 90 | 43 | +28 ms | 52 |
| unique | path | 262144 | ordered | 1 | 158 | 126 | +71 ms | 76 |
| unique | path | 262144 | baseline | 1 | 104 | 36 | +32 ms | 63 |
| unique | street | 212449 | ordered | 1 | 96 | 104 | +43 ms | 45 |
| unique | street | 212449 | baseline | 1 | 48 | 33 | +12 ms | 31 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
