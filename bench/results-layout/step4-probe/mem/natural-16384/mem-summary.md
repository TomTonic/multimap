| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | u64 | 16384 | ordered | 2 | 128 | 27 | +1 ms | 62 |
| natural | u64 | 16384 | btree-sets | 2 | 350 | 87 | +2 ms | 181 |
| natural | str | 16384 | ordered | 2 | 133 | 27 | +0 ms | 71 |
| natural | str | 16384 | btree-sets | 2 | 371 | 87 | +2 ms | 187 |
| natural | uuid | 16384 | ordered | 2 | 164 | 28 | +1 ms | 85 |
| natural | uuid | 16384 | btree-sets | 2 | 390 | 87 | +2 ms | 197 |
| natural | email | 16384 | ordered | 2 | 152 | 27 | +0 ms | 77 |
| natural | email | 16384 | btree-sets | 2 | 372 | 86 | +2 ms | 187 |
| natural | url | 16384 | ordered | 2 | 187 | 35 | +1 ms | 98 |
| natural | url | 16384 | btree-sets | 2 | 415 | 85 | +3 ms | 209 |
| natural | path | 16384 | ordered | 2 | 170 | 37 | +1 ms | 89 |
| natural | path | 16384 | btree-sets | 2 | 414 | 87 | +2 ms | 208 |
| natural | street | 16384 | ordered | 2 | 74 | 21 | +1 ms | 44 |
| natural | street | 16384 | btree-sets | 2 | 280 | 83 | +2 ms | 146 |
| natural | dirs | 16384 | ordered | 2 | 97 | 29 | +2 ms | 55 |
| natural | dirs | 16384 | btree-sets | 2 | 326 | 85 | +3 ms | 161 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
