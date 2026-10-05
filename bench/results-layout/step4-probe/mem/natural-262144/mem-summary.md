| values | keys | n | candidate | rounds | heap B/key | scannable B/key | GC CPU per cycle | heap B/key after removing half the keys |
|---|---|---:|---|---:|---:|---:|---:|---:|
| natural | u64 | 262144 | ordered | 2 | 123 | 22 | +18 ms | 64 |
| natural | u64 | 262144 | btree-sets | 2 | 352 | 84 | +65 ms | 181 |
| natural | str | 262144 | ordered | 2 | 130 | 26 | +23 ms | 66 |
| natural | str | 262144 | btree-sets | 2 | 371 | 84 | +68 ms | 187 |
| natural | uuid | 262144 | ordered | 2 | 162 | 28 | +23 ms | 84 |
| natural | uuid | 262144 | btree-sets | 2 | 392 | 86 | +66 ms | 197 |
| natural | email | 262144 | ordered | 2 | 146 | 25 | +21 ms | 75 |
| natural | email | 262144 | btree-sets | 2 | 373 | 84 | +68 ms | 187 |
| natural | url | 262144 | ordered | 2 | 178 | 39 | +29 ms | 92 |
| natural | url | 262144 | btree-sets | 2 | 416 | 90 | +62 ms | 209 |
| natural | path | 262144 | ordered | 2 | 154 | 36 | +32 ms | 80 |
| natural | path | 262144 | btree-sets | 2 | 416 | 84 | +68 ms | 209 |
| natural | street | 212449 | ordered | 2 | 73 | 21 | +17 ms | 40 |
| natural | street | 212449 | btree-sets | 2 | 285 | 81 | +51 ms | 144 |
| natural | dirs | 86215 | ordered | 2 | 96 | 28 | +7 ms | 52 |
| natural | dirs | 86215 | btree-sets | 2 | 335 | 83 | +19 ms | 167 |

n: keys per candidate. Heap figures exclude the key corpus. GC CPU is per full cycle minus a process that holds only the corpus. After removing every second key, bytes are still per key of the full corpus.
