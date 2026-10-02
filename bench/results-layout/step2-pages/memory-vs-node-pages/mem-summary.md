| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique | u64 | 262144 | ordered | 3 | 18 | 1 | +0 ms | 10 |
| unique | u64 | 262144 | baseline | 3 | 16 | 1 | +0 ms | 10 |
| unique | str | 262144 | ordered | 3 | 31 | 3 | +2 ms | 21 |
| unique | str | 262144 | baseline | 3 | 51 | 27 | +25 ms | 23 |
| unique | uuid | 262144 | ordered | 3 | 65 | 4 | +3 ms | 44 |
| unique | uuid | 262144 | baseline | 3 | 68 | 27 | +24 ms | 32 |
| unique | email | 262144 | ordered | 3 | 56 | 7 | +2 ms | 32 |
| unique | email | 262144 | baseline | 3 | 61 | 25 | +22 ms | 27 |
| unique | url | 262144 | ordered | 3 | 91 | 10 | +2 ms | 60 |
| unique | url | 262144 | baseline | 3 | 90 | 42 | +20 ms | 42 |
| unique | path | 262144 | ordered | 3 | 67 | 11 | +8 ms | 47 |
| unique | path | 262144 | baseline | 3 | 73 | 37 | +37 ms | 34 |
| unique | street | 212449 | ordered | 3 | 32 | 4 | +2 ms | 20 |
| unique | street | 212449 | baseline | 3 | 59 | 34 | +25 ms | 26 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
