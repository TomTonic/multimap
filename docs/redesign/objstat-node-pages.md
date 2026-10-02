# Object statistic of node-pages

Every object of the tree of `multimap.Ordered`, built from the bench corpus of each case
(`node-pages` working state of 2026-10-02, before the parking commit; Go 1.27.1, amd64). Counted after
building, not after `churn`. Value sets inside set leaves (`vset` storage, 2% of the leaves
of multi maps) are not counted. Tool: [prototype/](prototype/).

- **block**: the Go size class the object occupies, including the 8-byte malloc header of
  objects with pointers above 512 bytes.
- **not ×64 / not ×128**: share of objects whose block size is not a multiple of 64 / 128 bytes.
- **line overflow**: expected share of objects that touch more 64-byte lines than their size
  needs, averaged over the positions a block of that class takes in its span.
- **mix**: share of each object kind (N5..N256 inner nodes, +tail with a path tail, R8..R256
  range nodes, page, flat/typed/set leaf).

Profiles: `multi`/`unique` with uint64 values, `-str` with string values (bench tag
`strvals`). `path` and `street` corpora hold at most 300 000 and 212 449 keys; their last rows
are at that maximum.

| case | objects | not ×64 | not ×128 | line overflow | mix |
|---|--:|--:|--:|--:|---|
| u64 multi 4K | 4467 | 55.8 % | 83.3 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; flat 90%; set 2% |
| u64 multi 16K | 18365 | 55.0 % | 88.5 % | 0.0 % | N256 1%; N5 9%; N58 1%; flat 87%; set 2% |
| u64 multi 256K | 323893 | 49.8 % | 86.9 % | 0.0 % | N12 4%; N256 0%; N26 0%; N5 15%; flat 79%; set 2% |
| u64 multi 1M | 1145874 | 56.4 % | 84.4 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; flat 89%; set 2% |
| u64 unique 4K | 195 | 0.0 % | 0.0 % | 0.5 % | R256 1%; page 99% |
| u64 unique 16K | 1015 | 0.0 % | 0.0 % | 0.1 % | R256 0%; R8 25%; page 75% |
| u64 unique 256K | 12406 | 0.0 % | 0.0 % | 0.0 % | R256 0%; R56 2%; page 98% |
| u64 unique 1M | 51364 | 0.0 % | 0.0 % | 0.5 % | R256 1%; R8 0%; page 99% |
| u64 multi-str 4K | 4467 | 91.7 % | 94.2 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; set 3%; typed 89% |
| u64 multi-str 16K | 18365 | 89.2 % | 98.6 % | 0.0 % | N256 1%; N5 9%; N58 1%; set 3%; typed 87% |
| u64 multi-str 256K | 323893 | 80.9 % | 95.7 % | 0.0 % | N12 4%; N256 0%; N26 0%; N5 15%; set 2%; typed 79% |
| u64 multi-str 1M | 1145874 | 91.5 % | 94.3 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; set 3%; typed 89% |
| u64 unique-str 4K | 4467 | 91.7 % | 94.2 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; typed 92% |
| u64 unique-str 16K | 18365 | 89.2 % | 98.6 % | 0.0 % | N256 1%; N5 9%; N58 1%; typed 89% |
| u64 unique-str 256K | 323893 | 80.9 % | 95.7 % | 0.0 % | N12 4%; N256 0%; N26 0%; N5 15%; typed 81% |
| u64 unique-str 1M | 1145874 | 91.5 % | 94.3 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; typed 92% |
| str multi 4K | 5920 | 37.5 % | 88.3 % | 2.0 % | N12 4%; N26 0%; N5 27%; flat 68%; set 2% |
| str multi 16K | 22315 | 40.9 % | 87.0 % | 0.8 % | N12 4%; N26 0%; N5 22%; flat 72%; set 2% |
| str multi 256K | 358598 | 42.0 % | 89.9 % | 0.1 % | N12 1%; N26 1%; N5 25%; flat 71%; set 2% |
| str multi 1M | 1401474 | 43.8 % | 85.4 % | 0.0 % | N12 6%; N26 0%; N5 19%; flat 73%; set 2% |
| str unique 4K | 5920 | 69.2 % | 96.0 % | 3.5 % | N12 4%; N26 0%; N5 27%; flat 69% |
| str unique 16K | 22315 | 73.4 % | 95.4 % | 1.3 % | N12 4%; N26 0%; N5 22%; flat 73% |
| str unique 256K | 358598 | 73.1 % | 98.1 % | 0.1 % | N12 1%; N26 1%; N5 25%; flat 73% |
| str unique 1M | 1401474 | 74.8 % | 93.7 % | 0.0 % | N12 6%; N26 0%; N5 19%; flat 75% |
| str multi-str 4K | 5920 | 69.2 % | 96.0 % | 14.1 % | N12 4%; N26 0%; N5 27%; set 2%; typed 67% |
| str multi-str 16K | 22315 | 73.4 % | 95.4 % | 10.8 % | N12 4%; N26 0%; N5 22%; set 2%; typed 71% |
| str multi-str 256K | 358598 | 73.1 % | 98.1 % | 6.5 % | N12 1%; N26 1%; N5 25%; set 2%; typed 71% |
| str multi-str 1M | 1401474 | 74.8 % | 93.7 % | 5.0 % | N12 6%; N26 0%; N5 19%; set 2%; typed 73% |
| str unique-str 4K | 5920 | 69.2 % | 96.0 % | 27.7 % | N12 4%; N26 0%; N5 27%; typed 69% |
| str unique-str 16K | 22315 | 73.4 % | 95.4 % | 21.5 % | N12 4%; N26 0%; N5 22%; typed 73% |
| str unique-str 256K | 358598 | 73.1 % | 98.1 % | 13.0 % | N12 1%; N26 1%; N5 25%; typed 73% |
| str unique-str 1M | 1401474 | 74.8 % | 93.7 % | 9.9 % | N12 6%; N26 0%; N5 19%; typed 75% |
| uuid multi 4K | 5505 | 49.6 % | 94.6 % | 18.5 % | N12 4%; N26 1%; N5 21%; flat 73%; set 2% |
| uuid multi 16K | 22101 | 49.6 % | 96.1 % | 18.6 % | N12 2%; N26 1%; N5 22%; flat 72%; set 2% |
| uuid multi 256K | 353798 | 49.5 % | 96.1 % | 18.5 % | N12 2%; N26 1%; N5 22%; flat 72%; set 2% |
| uuid multi 1M | 1414691 | 49.6 % | 94.6 % | 18.5 % | N12 4%; N26 1%; N5 21%; flat 72%; set 2% |
| uuid unique 4K | 5505 | 74.1 % | 95.1 % | 37.1 % | N12 4%; N26 1%; N5 21%; flat 74% |
| uuid unique 16K | 22101 | 74.1 % | 96.5 % | 37.0 % | N12 2%; N26 1%; N5 22%; flat 74% |
| uuid unique 256K | 353798 | 74.1 % | 96.6 % | 37.0 % | N12 2%; N26 1%; N5 22%; flat 74% |
| uuid unique 1M | 1414691 | 74.1 % | 95.1 % | 37.1 % | N12 4%; N26 1%; N5 21%; flat 74% |
| uuid multi-str 4K | 5505 | 28.9 % | 93.4 % | 14.5 % | N12 4%; N26 1%; N5 21%; set 2%; typed 72% |
| uuid multi-str 16K | 22101 | 29.9 % | 95.9 % | 15.0 % | N12 2%; N26 1%; N5 22%; set 2%; typed 72% |
| uuid multi-str 256K | 353798 | 31.1 % | 96.5 % | 15.6 % | N12 2%; N26 1%; N5 22%; set 2%; typed 72% |
| uuid multi-str 1M | 1414691 | 31.2 % | 95.1 % | 15.6 % | N12 4%; N26 1%; N5 21%; set 2%; typed 72% |
| uuid unique-str 4K | 5505 | 0.0 % | 95.1 % | 0.0 % | N12 4%; N26 1%; N5 21%; typed 74% |
| uuid unique-str 16K | 22101 | 0.0 % | 96.5 % | 0.0 % | N12 2%; N26 1%; N5 22%; typed 74% |
| uuid unique-str 256K | 353798 | 0.0 % | 96.6 % | 0.0 % | N12 2%; N26 1%; N5 22%; typed 74% |
| uuid unique-str 1M | 1414691 | 0.0 % | 95.1 % | 0.0 % | N12 4%; N26 1%; N5 21%; typed 74% |
| email multi 4K | 5210 | 50.8 % | 91.6 % | 19.5 % | N12 6%; N26 1%; N5 15%; flat 77%; set 2% |
| email multi 16K | 21472 | 49.2 % | 94.5 % | 18.5 % | N12 0%; N26 3%; N5 20%; flat 75%; set 2% |
| email multi 256K | 333601 | 49.8 % | 91.8 % | 17.5 % | N12 4%; N26 2%; N5 16%; flat 77%; set 2% |
| email multi 1M | 1415558 | 46.6 % | 95.2 % | 15.8 % | N12 1%; N26 1%; N5 24%; flat 72%; set 2% |
| email unique 4K | 5210 | 78.6 % | 93.6 % | 33.2 % | N12 6%; N26 1%; N5 15%; flat 79% |
| email unique 16K | 21472 | 76.3 % | 96.7 % | 30.8 % | N12 0%; N26 3%; N5 20%; flat 76% |
| email unique 256K | 333601 | 78.6 % | 94.6 % | 29.1 % | N12 4%; N26 2%; N5 16%; flat 79% |
| email unique 1M | 1415558 | 74.1 % | 98.1 % | 26.0 % | N12 1%; N26 1%; N5 24%; flat 74% |
| email multi-str 4K | 5210 | 72.4 % | 93.6 % | 20.1 % | N12 6%; N26 1%; N5 15%; set 2%; typed 76% |
| email multi-str 16K | 21472 | 71.3 % | 96.7 % | 19.7 % | N12 0%; N26 3%; N5 20%; set 2%; typed 74% |
| email multi-str 256K | 333601 | 75.1 % | 94.6 % | 20.3 % | N12 4%; N26 2%; N5 16%; set 2%; typed 76% |
| email multi-str 1M | 1415558 | 71.5 % | 98.1 % | 19.1 % | N12 1%; N26 1%; N5 24%; set 2%; typed 72% |
| email unique-str 4K | 5210 | 68.2 % | 93.6 % | 34.1 % | N12 6%; N26 1%; N5 15%; typed 79% |
| email unique-str 16K | 21472 | 67.9 % | 96.7 % | 33.9 % | N12 0%; N26 3%; N5 20%; typed 76% |
| email unique-str 256K | 333601 | 72.6 % | 94.6 % | 36.3 % | N12 4%; N26 2%; N5 16%; typed 79% |
| email unique-str 1M | 1415558 | 69.7 % | 98.1 % | 34.8 % | N12 1%; N26 1%; N5 24%; typed 74% |
| url multi 4K | 6064 | 38.7 % | 89.6 % | 4.4 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; flat 66%; set 2% |
| url multi 16K | 24394 | 38.5 % | 89.8 % | 5.0 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; flat 66%; set 2% |
| url multi 256K | 392837 | 39.4 % | 91.3 % | 6.0 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 2%; N58 0%; flat 65%; set 2% |
| url multi 1M | 1566212 | 40.0 % | 91.8 % | 6.4 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 1%; N58 0%; flat 65%; set 2% |
| url unique 4K | 6064 | 44.6 % | 93.6 % | 7.7 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; flat 68% |
| url unique 16K | 24394 | 46.0 % | 93.7 % | 8.7 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; flat 67% |
| url unique 256K | 392837 | 50.7 % | 95.2 % | 10.3 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 2%; N58 0%; flat 67% |
| url unique 1M | 1566212 | 52.7 % | 95.7 % | 11.0 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 1%; N58 0%; flat 67% |
| url multi-str 4K | 6064 | 49.2 % | 90.5 % | 12.8 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; set 20%; typed 47% |
| url multi-str 16K | 24394 | 49.8 % | 91.1 % | 12.5 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; set 18%; typed 50% |
| url multi-str 256K | 392837 | 51.9 % | 92.5 % | 12.1 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 2%; N58 0%; set 13%; typed 54% |
| url multi-str 1M | 1566212 | 53.1 % | 92.8 % | 11.8 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 1%; N58 0%; set 11%; typed 56% |
| url unique-str 4K | 6064 | 50.0 % | 97.6 % | 11.4 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; set 18%; typed 49% |
| url unique-str 16K | 24394 | 49.7 % | 97.2 % | 12.0 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; set 16%; typed 51% |
| url unique-str 256K | 392837 | 51.1 % | 97.4 % | 13.1 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 2%; N58 0%; set 12%; typed 55% |
| url unique-str 1M | 1566212 | 52.2 % | 97.4 % | 13.6 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 1%; N58 0%; set 9%; typed 57% |
| path multi 4K | 6341 | 42.1 % | 94.3 % | 7.5 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 27%; N5+tail 6%; N58 0%; N58+tail 0%; flat 63%; set 2% |
| path multi 16K | 25353 | 42.4 % | 94.5 % | 7.8 % | N12 1%; N12+tail 0%; N26 0%; N26+tail 0%; N5 28%; N5+tail 6%; N58 0%; N58+tail 0%; flat 63%; set 1% |
| path multi 256K | 401574 | 41.4 % | 93.2 % | 6.2 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; flat 64%; set 1% |
| path unique 4K | 6341 | 52.8 % | 96.8 % | 13.2 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 27%; N5+tail 6%; N58 0%; N58+tail 0%; flat 65% |
| path unique 16K | 25353 | 56.8 % | 97.5 % | 13.5 % | N12 1%; N12+tail 0%; N26 0%; N26+tail 0%; N5 28%; N5+tail 6%; N58 0%; N58+tail 0%; flat 65% |
| path unique 256K | 401574 | 62.5 % | 97.7 % | 10.6 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; flat 65% |
| path multi-str 4K | 6341 | 51.9 % | 93.8 % | 12.3 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 27%; N5+tail 6%; N58 0%; N58+tail 0%; set 9%; typed 56% |
| path multi-str 16K | 25353 | 54.8 % | 95.2 % | 12.4 % | N12 1%; N12+tail 0%; N26 0%; N26+tail 0%; N5 28%; N5+tail 6%; N58 0%; N58+tail 0%; set 6%; typed 59% |
| path multi-str 256K | 401574 | 60.6 % | 96.6 % | 11.2 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; set 3%; typed 62% |
| path unique-str 4K | 6341 | 48.5 % | 98.1 % | 12.7 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 27%; N5+tail 6%; N58 0%; N58+tail 0%; set 7%; typed 58% |
| path unique-str 16K | 25353 | 50.7 % | 98.2 % | 14.7 % | N12 1%; N12+tail 0%; N26 0%; N26+tail 0%; N5 28%; N5+tail 6%; N58 0%; N58+tail 0%; set 4%; typed 60% |
| path unique-str 256K | 401574 | 57.3 % | 97.8 % | 16.6 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; set 1%; typed 64% |
| street multi 4K | 5977 | 58.8 % | 94.1 % | 0.6 % | N12 3%; N26 0%; N5 28%; N58 0%; flat 68%; set 0% |
| street multi 16K | 23984 | 59.8 % | 94.6 % | 0.4 % | N12 3%; N26 0%; N5 29%; N58 0%; flat 68%; set 0% |
| street unique 4K | 5977 | 68.5 % | 96.5 % | 0.6 % | N12 3%; N26 0%; N5 28%; N58 0%; flat 69% |
| street unique 16K | 23984 | 68.3 % | 96.9 % | 0.4 % | N12 3%; N26 0%; N5 29%; N58 0%; flat 68% |
| street multi-str 4K | 5977 | 68.4 % | 96.5 % | 9.0 % | N12 3%; N26 0%; N5 28%; N58 0%; set 1%; typed 67% |
| street multi-str 16K | 23984 | 68.2 % | 96.9 % | 7.2 % | N12 3%; N26 0%; N5 29%; N58 0%; set 1%; typed 67% |
| street unique-str 4K | 5977 | 68.4 % | 96.5 % | 10.9 % | N12 3%; N26 0%; N5 28%; N58 0%; typed 69% |
| street unique-str 16K | 23984 | 68.2 % | 96.9 % | 8.6 % | N12 3%; N26 0%; N5 29%; N58 0%; set 0%; typed 68% |
| path multi 300000 | 459351 | 41.3 % | 93.1 % | 6.1 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; flat 64%; set 1% |
| path unique 300000 | 459351 | 62.8 % | 97.7 % | 10.4 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; flat 65% |
| path multi-str 300000 | 459351 | 60.9 % | 96.6 % | 11.1 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; set 3%; typed 62% |
| path unique-str 300000 | 459351 | 57.7 % | 97.8 % | 16.6 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; set 1%; typed 64% |
| street multi 212449 | 314730 | 60.3 % | 95.0 % | 0.2 % | N12 2%; N26 0%; N5 30%; N5+tail 0%; N58 0%; flat 67%; set 0% |
| street unique 212449 | 314730 | 67.5 % | 97.2 % | 0.2 % | N12 2%; N26 0%; N5 30%; N5+tail 0%; N58 0%; flat 68% |
| street multi-str 212449 | 314730 | 67.5 % | 97.2 % | 5.3 % | N12 2%; N26 0%; N5 30%; N5+tail 0%; N58 0%; set 1%; typed 67% |
| street unique-str 212449 | 314730 | 67.5 % | 97.2 % | 6.7 % | N12 2%; N26 0%; N5 30%; N5+tail 0%; N58 0%; set 0%; typed 68% |
