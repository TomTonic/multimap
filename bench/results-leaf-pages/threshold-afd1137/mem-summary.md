| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| multi | u64 | 65536 | ordered | 1 | 159 | 92 | +6 ms | 76 |
| multi | u64 | 65536 | baseline | 1 | 110 | 47 | +2 ms | 63 |
| multi | str | 65536 | ordered | 1 | 177 | 109 | +8 ms | 84 |
| multi | str | 65536 | baseline | 1 | 121 | 52 | +2 ms | 71 |
| multi | uuid | 65536 | ordered | 1 | 194 | 127 | +8 ms | 93 |
| multi | uuid | 65536 | baseline | 1 | 154 | 72 | +3 ms | 93 |
| multi | path | 65536 | ordered | 1 | 235 | 128 | +12 ms | 113 |
| multi | path | 65536 | baseline | 1 | 163 | 59 | +7 ms | 103 |
| multi | street | 65536 | ordered | 1 | 119 | 104 | +9 ms | 56 |
| multi | street | 65536 | baseline | 1 | 60 | 32 | +3 ms | 38 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
