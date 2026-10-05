| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | u64 | 65536 | ordered | 2 | 126 | 25 | +4 ms | 66 |
| natural | u64 | 65536 | btree-sets | 2 | 351 | 85 | +13 ms | 182 |
| natural | str | 65536 | ordered | 2 | 130 | 26 | +2 ms | 67 |
| natural | str | 65536 | btree-sets | 2 | 370 | 83 | +12 ms | 188 |
| natural | uuid | 65536 | ordered | 2 | 163 | 27 | +2 ms | 83 |
| natural | uuid | 65536 | btree-sets | 2 | 392 | 86 | +13 ms | 198 |
| natural | email | 65536 | ordered | 2 | 147 | 25 | +2 ms | 76 |
| natural | email | 65536 | btree-sets | 2 | 372 | 86 | +13 ms | 189 |
| natural | url | 65536 | ordered | 2 | 183 | 36 | +8 ms | 94 |
| natural | url | 65536 | btree-sets | 2 | 414 | 86 | +15 ms | 209 |
| natural | path | 65536 | ordered | 2 | 161 | 36 | +8 ms | 84 |
| natural | path | 65536 | btree-sets | 2 | 415 | 85 | +15 ms | 210 |
| natural | street | 65536 | ordered | 2 | 76 | 21 | +3 ms | 42 |
| natural | street | 65536 | btree-sets | 2 | 287 | 82 | +14 ms | 146 |
| natural | dirs | 65536 | ordered | 2 | 97 | 28 | +4 ms | 52 |
| natural | dirs | 65536 | btree-sets | 2 | 335 | 84 | +11 ms | 168 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
