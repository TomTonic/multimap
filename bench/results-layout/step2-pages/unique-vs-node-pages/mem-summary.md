| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| unique | u64 | 262144 | ordered | 3 | 18 | 1 | +1 ms | 12 |
| unique | u64 | 262144 | baseline | 3 | 16 | 1 | +1 ms | 10 |
| unique | str | 262144 | ordered | 3 | 31 | 3 | +1 ms | 23 |
| unique | str | 262144 | baseline | 3 | 51 | 27 | +25 ms | 23 |
| unique | uuid | 262144 | ordered | 3 | 65 | 4 | +3 ms | 46 |
| unique | uuid | 262144 | baseline | 3 | 68 | 27 | +25 ms | 32 |
| unique | email | 262144 | ordered | 3 | 56 | 7 | +2 ms | 38 |
| unique | email | 262144 | baseline | 3 | 61 | 25 | +21 ms | 27 |
| unique | url | 262144 | ordered | 3 | 91 | 15 | +1 ms | 66 |
| unique | url | 262144 | baseline | 3 | 90 | 45 | +17 ms | 42 |
| unique | path | 262144 | ordered | 3 | 67 | 11 | +6 ms | 51 |
| unique | path | 262144 | baseline | 3 | 73 | 38 | +37 ms | 34 |
| unique | street | 212449 | ordered | 3 | 32 | 4 | +2 ms | 23 |
| unique | street | 212449 | baseline | 3 | 59 | 34 | +26 ms | 26 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
