| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | u64 | 4096 | ordered | 2 | 129 | 19 | +0 ms | 70 |
| natural | u64 | 4096 | btree-sets | 2 | 368 | 87 | +0 ms | 193 |
| natural | str | 4096 | ordered | 2 | 148 | 32 | +0 ms | 81 |
| natural | str | 4096 | btree-sets | 2 | 387 | 89 | +0 ms | 202 |
| natural | uuid | 4096 | ordered | 2 | 171 | 28 | -0 ms | 92 |
| natural | uuid | 4096 | btree-sets | 2 | 407 | 89 | +0 ms | 211 |
| natural | email | 4096 | ordered | 2 | 158 | 25 | +0 ms | 88 |
| natural | email | 4096 | btree-sets | 2 | 390 | 89 | +0 ms | 200 |
| natural | url | 4096 | ordered | 2 | 200 | 36 | +0 ms | 110 |
| natural | url | 4096 | btree-sets | 2 | 431 | 89 | +1 ms | 221 |
| natural | path | 4096 | ordered | 2 | 188 | 38 | +0 ms | 105 |
| natural | path | 4096 | btree-sets | 2 | 430 | 87 | +1 ms | 224 |
| natural | street | 4096 | ordered | 2 | 77 | 22 | +1 ms | 50 |
| natural | street | 4096 | btree-sets | 2 | 281 | 82 | +1 ms | 152 |
| natural | dirs | 4096 | ordered | 2 | 122 | 29 | +0 ms | 66 |
| natural | dirs | 4096 | btree-sets | 2 | 313 | 82 | +1 ms | 154 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
