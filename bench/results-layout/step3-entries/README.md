# Entries with several values: how many would fit one object of 512 bytes

Measured on 2026-10-03 to size the single-key page of PLAN step 3: the object that holds one entry
with all its values inline. The question: for entries with two or more values, how many would need the
fallback (a value set, or an object beyond the grid), as a function of how many values the page's
header has room for.

Tool: `go run ./cmd/objstat -keys u64,uuid,email,url,path,street -values multi -sizes 16384,262144
-entries -strvals=false` (`-entries` is new; the table is in [objstat-entries.md](objstat-entries.md)).
The tree is built from the corpus as the benchmark builds it; every leaf reports its number of values
and the length of its key remainder (`art.Object.Values`, `Remainder`).

**Model.** An entry *fits* for a header of room N if it has at most N values and
`header(N) + remainder + values * valueBytes <= 512`, with `header(N) = 8 * ceil((2 + N) / 8)` (type and
remainder length, one length byte per value, rounded up to a multiple of 8) and `valueBytes` 8 for
`uint64`, 16 for the 16-hex-digit string values. "remainder > 503" counts remainders that do not fit
behind a header of 8 bytes at all (they would be oversized objects).

## Result

Share of the entries with several values that fit without the fallback.

| N values in the header | header | synthetic skew (u64, uuid, email, url, path; 256K keys) | street names (natural counts, 212K keys) |
|--:|--:|--:|--:|
| 3 | 8 B | 46.6 % | 64.7 % |
| 4 | 8 B | 69.9 % | 73.7 % |
| 5 | 8 B | 72.0 % | 78.8 % |
| 6 | 8 B | 73.9 % | 82.4 % |
| 7 | 16 B | 75.9 % | 84.8 % |
| 8 | 16 B | 77.9 % | 86.7 % |
| 10 | 16 B | 82.0 % | 89.2 % |
| 12 | 16 B | 86.1 % | 90.9 % |
| 14 | 16 B | 90.1 % | 92.2 % |
| 16 | 24 B | 94.0 % | 93.3 % |
| 64 | 72 B | 95.2 % | - |

In terms of all keys of a case (half of the keys of the skewed profile have one value, 79% of the
street names):

| N | keys that need the fallback, skewed profile | street names |
|--:|--:|--:|
| 6 | 13.1 % | 3.6 % |
| 10 | 9.0 % | 2.2 % |
| 14 | 5.0 % | 1.6 % |

## Reading

- **The count decides, the bytes do not.** The rows of all key kinds, `uint64` and string values, and
  the 16K and 256K sizes are identical to the first decimal, because the number of values per key
  comes from the profile and not from the key. The extreme sums are far below 512 bytes
  (300 bytes of key + 16 values of 8 bytes + a header of 16 is 444), except a 300-byte key with 16
  string values.
- **98% is not reachable by a bigger header.** The profile has 6% of its multi-value entries at 17
  values or more, mostly many more: even a header with room for 64 values (72 bytes) covers 95.2%.
  About 5% of the multi-value entries (2.4% of all keys) have more values than a page of 512
  bytes can hold, whatever the header, and take the fallback.
- **The knee is at 4 for the skewed profile** (46.6 -> 69.9 % from 3 to 4, then 2 points per step
  until 14) and at 6 for the street names (natural counts, 79% of the names have one locality).
- **Only the street names have natural counts.** The other six kinds use the skewed profile of the
  bench (`-values multi`: 50% one value, 35% 2-4, 12% 5-16, 3% 17-200), which is synthetic by
  construction.

## Caveats

- The remainders are those of today's leaves, from their base on. In the tree of the target the
  single-key page hangs closer to the root, so its remainder is longer: at most the whole key, 300
  bytes for `path`. That cannot change the table except for the case above, but it makes the
  oversized objects (remainder above the page) more likely for keys beyond 500 bytes; the corpora have
  none.
- All leaves are counted, not only those a tree with the single-key page would have: with half of
  the keys multi-valued, today's tree falls back everywhere (`settle`), so every key is a leaf.
  Only leaves with two or more values enter the table.
- The model charges one byte per value for its length (the user's sketch). For fixed-size values a
  count byte would do and the header would shrink; the table's bytes would then bind earlier only for
  values of 16 bytes and more.
