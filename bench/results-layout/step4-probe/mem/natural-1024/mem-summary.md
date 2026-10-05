| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | u64 | 1024 | ordered | 2 | 149 | 22 | +0 ms | 98 |
| natural | u64 | 1024 | btree-sets | 2 | 379 | 87 | +0 ms | 212 |
| natural | str | 1024 | ordered | 2 | 163 | 29 | +0 ms | 101 |
| natural | str | 1024 | btree-sets | 2 | 402 | 93 | +0 ms | 221 |
| natural | uuid | 1024 | ordered | 2 | 187 | 29 | +0 ms | 118 |
| natural | uuid | 1024 | btree-sets | 2 | 422 | 91 | +0 ms | 236 |
| natural | email | 1024 | ordered | 2 | 176 | 28 | +0 ms | 114 |
| natural | email | 1024 | btree-sets | 2 | 408 | 92 | -0 ms | 227 |
| natural | url | 1024 | ordered | 2 | 222 | 35 | +0 ms | 137 |
| natural | url | 1024 | btree-sets | 2 | 446 | 88 | +0 ms | 264 |
| natural | path | 1024 | ordered | 2 | 212 | 39 | +0 ms | 140 |
| natural | path | 1024 | btree-sets | 2 | 451 | 89 | +0 ms | 260 |
| natural | street | 1024 | ordered | 2 | 79 | 23 | +0 ms | 58 |
| natural | street | 1024 | btree-sets | 2 | 282 | 93 | +0 ms | 163 |
| natural | dirs | 1024 | ordered | 2 | 109 | 31 | +0 ms | 99 |
| natural | dirs | 1024 | btree-sets | 2 | 336 | 89 | +1 ms | 214 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
