# Object statistic of node-pages (the starting point)

Produced by `go run ./cmd/objstat` in `bench/` (see the tool's package comment for the exact
definitions), on the tree of branch `node-pages` (`cf6944b`), Go 1.27.1, amd64, 2026-10-02. The
index is built from the corpus of each case as the benchmark builds it and counted afterwards,
not after `churn`. The last row of `path` and `street` is at the maximum their corpus allows
(300 000 and 212 449 keys). Profiles: `multi` and `unique` have uint64 values, the `-str`
profiles string values.

Targets of the redesign (STRATEGY.md, R1-R6): 100% of the objects at multiples of 64 bytes, at
least 95% at multiples of 128, no line overflow.

| case | objects | block bytes per key | not x64 | not x128 | line overflow | mix |
|---|--:|--:|--:|--:|--:|---|
| u64 multi 4K | 4478 | 73.8 | 55.7 % | 83.4 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; flat leaf 89%; set leaf 2% |
| u64 multi 16K | 18387 | 81.7 | 54.9 % | 88.5 % | 0.0 % | N256 1%; N5 9%; N58 1%; flat leaf 87%; set leaf 2% |
| u64 multi 256K | 323927 | 76.6 | 49.8 % | 86.9 % | 0.0 % | N12 4%; N256 0%; N26 0%; N5 15%; flat leaf 79%; set leaf 2% |
| u64 multi 1M | 1146008 | 72.8 | 56.4 % | 84.4 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; flat leaf 89%; set leaf 2% |
| u64 multi-str 4K | 4478 | 94.7 | 91.5 % | 94.3 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; set leaf 3%; typed leaf 89% |
| u64 multi-str 16K | 18387 | 102.6 | 89.1 % | 98.6 % | 0.0 % | N256 1%; N5 9%; N58 1%; set leaf 3%; typed leaf 87% |
| u64 multi-str 256K | 323927 | 96.5 | 80.9 % | 95.7 % | 0.0 % | N12 4%; N256 0%; N26 0%; N5 15%; set leaf 2%; typed leaf 79% |
| u64 multi-str 1M | 1146008 | 92.5 | 91.5 % | 94.3 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; set leaf 3%; typed leaf 89% |
| u64 unique 4K | 202 | 24.2 | 0.0 % | 0.0 % | 0.5 % | R256 0%; page 100% |
| u64 unique 16K | 1025 | 26.1 | 0.0 % | 0.0 % | 0.1 % | R256 0%; R8 25%; page 75% |
| u64 unique 256K | 12434 | 24.1 | 0.0 % | 0.0 % | 0.0 % | R256 0%; R56 2%; page 98% |
| u64 unique 1M | 51465 | 24.2 | 0.0 % | 0.0 % | 0.5 % | R256 0%; R8 0%; page 99% |
| u64 unique-str 4K | 4478 | 48.8 | 91.5 % | 94.3 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; typed leaf 91% |
| u64 unique-str 16K | 18387 | 57.7 | 89.1 % | 98.6 % | 0.0 % | N256 1%; N5 9%; N58 1%; typed leaf 89% |
| u64 unique-str 256K | 323927 | 52.6 | 80.9 % | 95.7 % | 0.0 % | N12 4%; N256 0%; N26 0%; N5 15%; typed leaf 81% |
| u64 unique-str 1M | 1146008 | 48.8 | 91.5 % | 94.3 % | 0.0 % | N12 1%; N256 0%; N26 4%; N5 3%; N58 0%; typed leaf 91% |
| str multi 4K | 5947 | 95.1 | 37.7 % | 88.5 % | 2.5 % | N12 3%; N26 0%; N5 27%; flat leaf 67%; set leaf 2% |
| str multi 16K | 22317 | 87.1 | 40.9 % | 87.0 % | 1.0 % | N12 5%; N26 0%; N5 22%; flat leaf 72%; set leaf 2% |
| str multi 256K | 358452 | 85.4 | 41.8 % | 89.9 % | 0.1 % | N12 1%; N26 1%; N5 25%; flat leaf 71%; set leaf 2% |
| str multi 1M | 1400191 | 85.2 | 43.7 % | 85.4 % | 0.0 % | N12 6%; N26 0%; N5 19%; flat leaf 73%; set leaf 2% |
| str multi-str 4K | 5947 | 122.3 | 68.9 % | 96.3 % | 14.3 % | N12 3%; N26 0%; N5 27%; set leaf 2%; typed leaf 67% |
| str multi-str 16K | 22317 | 112.7 | 73.4 % | 95.4 % | 11.3 % | N12 5%; N26 0%; N5 22%; set leaf 2%; typed leaf 71% |
| str multi-str 256K | 358452 | 108.5 | 73.1 % | 98.0 % | 6.9 % | N12 1%; N26 1%; N5 25%; set leaf 2%; typed leaf 71% |
| str multi-str 1M | 1400191 | 107.3 | 74.9 % | 93.7 % | 5.3 % | N12 6%; N26 0%; N5 19%; set leaf 2%; typed leaf 73% |
| str unique 4K | 5947 | 66.9 | 68.9 % | 96.3 % | 4.2 % | N12 3%; N26 0%; N5 27%; flat leaf 69% |
| str unique 16K | 22317 | 60.0 | 73.4 % | 95.4 % | 1.6 % | N12 5%; N26 0%; N5 22%; flat leaf 73% |
| str unique 256K | 358452 | 59.3 | 73.1 % | 98.0 % | 0.1 % | N12 1%; N26 1%; N5 25%; flat leaf 73% |
| str unique 1M | 1400191 | 59.6 | 74.9 % | 93.7 % | 0.0 % | N12 6%; N26 0%; N5 19%; flat leaf 75% |
| str unique-str 4K | 5947 | 77.8 | 68.9 % | 96.3 % | 27.8 % | N12 3%; N26 0%; N5 27%; typed leaf 69% |
| str unique-str 16K | 22317 | 68.9 | 73.4 % | 95.4 % | 22.0 % | N12 5%; N26 0%; N5 22%; typed leaf 73% |
| str unique-str 256K | 358452 | 65.2 | 73.1 % | 98.0 % | 13.7 % | N12 1%; N26 1%; N5 25%; typed leaf 73% |
| str unique-str 1M | 1400191 | 64.1 | 74.9 % | 93.7 % | 10.6 % | N12 6%; N26 0%; N5 19%; typed leaf 75% |
| uuid multi 4K | 5503 | 104.4 | 49.7 % | 94.6 % | 18.5 % | N12 4%; N26 1%; N5 21%; flat leaf 73%; set leaf 2% |
| uuid multi 16K | 22101 | 103.2 | 49.7 % | 96.0 % | 18.6 % | N12 2%; N26 1%; N5 22%; flat leaf 72%; set leaf 2% |
| uuid multi 256K | 353518 | 102.6 | 49.6 % | 96.0 % | 18.5 % | N12 2%; N26 1%; N5 22%; flat leaf 72%; set leaf 2% |
| uuid multi 1M | 1414417 | 103.0 | 49.6 % | 94.6 % | 18.5 % | N12 4%; N26 1%; N5 21%; flat leaf 72%; set leaf 2% |
| uuid multi-str 4K | 5503 | 137.1 | 29.0 % | 93.5 % | 14.5 % | N12 4%; N26 1%; N5 21%; set leaf 2%; typed leaf 72% |
| uuid multi-str 16K | 22101 | 136.1 | 30.0 % | 95.9 % | 15.0 % | N12 2%; N26 1%; N5 22%; set leaf 2%; typed leaf 72% |
| uuid multi-str 256K | 353518 | 134.9 | 31.2 % | 96.5 % | 15.6 % | N12 2%; N26 1%; N5 22%; set leaf 2%; typed leaf 72% |
| uuid multi-str 1M | 1414417 | 135.2 | 31.2 % | 95.1 % | 15.6 % | N12 4%; N26 1%; N5 21%; set leaf 2%; typed leaf 72% |
| uuid unique 4K | 5503 | 75.4 | 74.1 % | 95.1 % | 37.1 % | N12 4%; N26 1%; N5 21%; flat leaf 74% |
| uuid unique 16K | 22101 | 75.6 | 74.1 % | 96.4 % | 37.0 % | N12 2%; N26 1%; N5 22%; flat leaf 74% |
| uuid unique 256K | 353518 | 75.5 | 74.1 % | 96.5 % | 37.1 % | N12 2%; N26 1%; N5 22%; flat leaf 74% |
| uuid unique 1M | 1414417 | 76.0 | 74.1 % | 95.1 % | 37.1 % | N12 4%; N26 1%; N5 21%; flat leaf 74% |
| uuid unique-str 4K | 5503 | 91.3 | 0.0 % | 95.1 % | 0.0 % | N12 4%; N26 1%; N5 21%; typed leaf 74% |
| uuid unique-str 16K | 22101 | 91.6 | 0.0 % | 96.4 % | 0.0 % | N12 2%; N26 1%; N5 22%; typed leaf 74% |
| uuid unique-str 256K | 353518 | 91.5 | 0.0 % | 96.5 % | 0.0 % | N12 2%; N26 1%; N5 22%; typed leaf 74% |
| uuid unique-str 1M | 1414417 | 92.0 | 0.0 % | 95.1 % | 0.0 % | N12 4%; N26 1%; N5 21%; typed leaf 74% |
| email multi 4K | 5196 | 92.9 | 50.7 % | 91.8 % | 19.6 % | N12 6%; N26 1%; N5 15%; flat leaf 77%; set leaf 2% |
| email multi 16K | 21475 | 96.1 | 49.3 % | 94.5 % | 18.6 % | N12 0%; N26 3%; N5 20%; flat leaf 75%; set leaf 2% |
| email multi 256K | 333391 | 92.2 | 49.9 % | 91.8 % | 17.6 % | N12 4%; N26 2%; N5 16%; flat leaf 77%; set leaf 2% |
| email multi 1M | 1415873 | 93.2 | 46.6 % | 95.2 % | 15.7 % | N12 1%; N26 1%; N5 24%; flat leaf 72%; set leaf 2% |
| email multi-str 4K | 5196 | 117.9 | 72.7 % | 93.6 % | 20.2 % | N12 6%; N26 1%; N5 15%; set leaf 2%; typed leaf 76% |
| email multi-str 16K | 21475 | 121.1 | 71.0 % | 96.7 % | 19.6 % | N12 0%; N26 3%; N5 20%; set leaf 2%; typed leaf 74% |
| email multi-str 256K | 333391 | 116.7 | 75.2 % | 94.6 % | 20.3 % | N12 4%; N26 2%; N5 16%; set leaf 2%; typed leaf 76% |
| email multi-str 1M | 1415873 | 117.6 | 71.5 % | 98.1 % | 19.1 % | N12 1%; N26 1%; N5 24%; set leaf 2%; typed leaf 72% |
| email unique 4K | 5196 | 68.6 | 78.8 % | 93.6 % | 33.0 % | N12 6%; N26 1%; N5 15%; flat leaf 79% |
| email unique 16K | 21475 | 72.6 | 76.3 % | 96.7 % | 30.8 % | N12 0%; N26 3%; N5 20%; flat leaf 76% |
| email unique 256K | 333391 | 68.8 | 78.6 % | 94.6 % | 29.1 % | N12 4%; N26 2%; N5 16%; flat leaf 79% |
| email unique 1M | 1415873 | 69.5 | 74.1 % | 98.1 % | 26.0 % | N12 1%; N26 1%; N5 24%; flat leaf 74% |
| email unique-str 4K | 5196 | 73.3 | 68.4 % | 93.6 % | 34.2 % | N12 6%; N26 1%; N5 15%; typed leaf 79% |
| email unique-str 16K | 21475 | 77.5 | 67.4 % | 96.7 % | 33.7 % | N12 0%; N26 3%; N5 20%; typed leaf 76% |
| email unique-str 256K | 333391 | 74.1 | 72.7 % | 94.6 % | 36.3 % | N12 4%; N26 2%; N5 16%; typed leaf 79% |
| email unique-str 1M | 1415873 | 75.2 | 69.7 % | 98.1 % | 34.8 % | N12 1%; N26 1%; N5 24%; typed leaf 74% |
| url multi 4K | 6065 | 132.3 | 37.6 % | 88.7 % | 4.6 % | N12 2%; N12+tail 0%; N26 1%; N5 28%; N5+tail 2%; N58 0%; flat leaf 66%; set leaf 2% |
| url multi 16K | 24400 | 128.2 | 38.0 % | 89.5 % | 5.2 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; flat leaf 66%; set leaf 2% |
| url multi 256K | 392855 | 121.6 | 39.4 % | 91.2 % | 6.0 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 2%; N58 0%; flat leaf 65%; set leaf 2% |
| url multi 1M | 1566213 | 117.9 | 40.0 % | 91.8 % | 6.4 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 1%; N58 0%; flat leaf 65%; set leaf 2% |
| url multi-str 4K | 6065 | 162.4 | 49.9 % | 91.4 % | 13.1 % | N12 2%; N12+tail 0%; N26 1%; N5 28%; N5+tail 2%; N58 0%; set leaf 21%; typed leaf 47% |
| url multi-str 16K | 24400 | 157.5 | 50.5 % | 91.8 % | 12.6 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; set leaf 18%; typed leaf 50% |
| url multi-str 256K | 392855 | 148.9 | 52.1 % | 92.5 % | 12.1 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 2%; N58 0%; set leaf 13%; typed leaf 54% |
| url multi-str 1M | 1566213 | 144.4 | 53.1 % | 92.9 % | 11.8 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 1%; N58 0%; set leaf 11%; typed leaf 56% |
| url unique 4K | 6065 | 107.5 | 45.4 % | 93.0 % | 8.0 % | N12 2%; N12+tail 0%; N26 1%; N5 28%; N5+tail 2%; N58 0%; flat leaf 68% |
| url unique 16K | 24400 | 104.2 | 47.0 % | 93.5 % | 9.0 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; flat leaf 67% |
| url unique 256K | 392855 | 97.8 | 50.6 % | 95.1 % | 10.3 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 2%; N58 0%; flat leaf 67% |
| url unique 1M | 1566213 | 94.2 | 52.7 % | 95.7 % | 10.9 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 1%; N58 0%; flat leaf 67% |
| url unique-str 4K | 6065 | 129.9 | 50.1 % | 97.5 % | 11.8 % | N12 2%; N12+tail 0%; N26 1%; N5 28%; N5+tail 2%; N58 0%; set leaf 19%; typed leaf 48% |
| url unique-str 16K | 24400 | 123.9 | 50.0 % | 97.3 % | 12.1 % | N12 2%; N12+tail 0%; N26 0%; N5 28%; N5+tail 2%; N58 0%; set leaf 16%; typed leaf 51% |
| url unique-str 256K | 392855 | 113.3 | 51.1 % | 97.4 % | 13.1 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 2%; N58 0%; set leaf 12%; typed leaf 55% |
| url unique-str 1M | 1566213 | 107.6 | 52.2 % | 97.4 % | 13.6 % | N12 2%; N12+tail 0%; N26 0%; N5 29%; N5+tail 1%; N58 0%; set leaf 10%; typed leaf 57% |
| path multi 4K | 6317 | 124.3 | 42.1 % | 93.4 % | 7.7 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 27%; N5+tail 6%; N58 0%; flat leaf 63%; set leaf 2% |
| path multi 16K | 25301 | 117.0 | 42.3 % | 94.3 % | 8.0 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 28%; N5+tail 5%; N58 0%; N58+tail 0%; flat leaf 63%; set leaf 1% |
| path multi 256K | 401413 | 105.7 | 41.4 % | 93.2 % | 6.2 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; flat leaf 64%; set leaf 1% |
| path multi 300000 | 459206 | 105.2 | 41.3 % | 93.1 % | 6.1 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; flat leaf 64%; set leaf 1% |
| path multi-str 4K | 6317 | 151.3 | 51.9 % | 93.9 % | 12.5 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 27%; N5+tail 6%; N58 0%; set leaf 9%; typed leaf 56% |
| path multi-str 16K | 25301 | 143.5 | 54.6 % | 95.1 % | 12.7 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 28%; N5+tail 5%; N58 0%; N58+tail 0%; set leaf 6%; typed leaf 59% |
| path multi-str 256K | 401413 | 130.5 | 60.6 % | 96.5 % | 11.2 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; set leaf 3%; typed leaf 62% |
| path multi-str 300000 | 459206 | 129.9 | 60.8 % | 96.6 % | 11.1 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; set leaf 3%; typed leaf 62% |
| path unique 4K | 6317 | 99.2 | 52.9 % | 96.6 % | 13.1 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 27%; N5+tail 6%; N58 0%; flat leaf 65% |
| path unique 16K | 25301 | 92.2 | 56.5 % | 97.3 % | 13.6 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 28%; N5+tail 5%; N58 0%; N58+tail 0%; flat leaf 65% |
| path unique 256K | 401413 | 80.5 | 62.6 % | 97.6 % | 10.7 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; flat leaf 65% |
| path unique 300000 | 459206 | 80.0 | 62.8 % | 97.7 % | 10.4 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; flat leaf 65% |
| path unique-str 4K | 6317 | 111.8 | 48.6 % | 97.9 % | 13.3 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 27%; N5+tail 6%; N58 0%; set leaf 7%; typed leaf 58% |
| path unique-str 16K | 25301 | 102.7 | 50.4 % | 98.1 % | 15.0 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 28%; N5+tail 5%; N58 0%; N58+tail 0%; set leaf 4%; typed leaf 60% |
| path unique-str 256K | 401413 | 88.2 | 57.3 % | 97.8 % | 16.6 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; set leaf 1%; typed leaf 64% |
| path unique-str 300000 | 459206 | 87.6 | 57.6 % | 97.8 % | 16.6 % | N12 2%; N12+tail 0%; N26 0%; N26+tail 0%; N5 29%; N5+tail 3%; N58 0%; N58+tail 0%; set leaf 1%; typed leaf 64% |
| street multi 4K | 5950 | 73.9 | 60.3 % | 94.1 % | 0.7 % | N12 3%; N26 0%; N5 28%; N58 0%; flat leaf 69%; set leaf 0% |
| street multi 16K | 23924 | 74.0 | 60.1 % | 94.6 % | 0.4 % | N12 3%; N26 0%; N5 28%; N5+tail 0%; N58 0%; flat leaf 68%; set leaf 0% |
| street multi 212449 | 314837 | 74.5 | 60.3 % | 95.1 % | 0.2 % | N12 2%; N26 0%; N5 30%; N5+tail 0%; N58 0%; flat leaf 67%; set leaf 0% |
| street multi-str 4K | 5950 | 84.1 | 68.7 % | 96.4 % | 9.3 % | N12 3%; N26 0%; N5 28%; N58 0%; set leaf 1%; typed leaf 68% |
| street multi-str 16K | 23924 | 83.6 | 68.4 % | 96.8 % | 7.2 % | N12 3%; N26 0%; N5 28%; N5+tail 0%; N58 0%; set leaf 1%; typed leaf 68% |
| street multi-str 212449 | 314837 | 82.1 | 67.4 % | 97.2 % | 5.3 % | N12 2%; N26 0%; N5 30%; N5+tail 0%; N58 0%; set leaf 1%; typed leaf 67% |
| street unique 4K | 5950 | 65.3 | 68.8 % | 96.4 % | 0.7 % | N12 3%; N26 0%; N5 28%; N58 0%; flat leaf 69% |
| street unique 16K | 23924 | 65.6 | 68.5 % | 96.8 % | 0.4 % | N12 3%; N26 0%; N5 28%; N5+tail 0%; N58 0%; flat leaf 68% |
| street unique 212449 | 314837 | 66.5 | 67.5 % | 97.2 % | 0.2 % | N12 2%; N26 0%; N5 30%; N5+tail 0%; N58 0%; flat leaf 67% |
| street unique-str 4K | 5950 | 69.6 | 68.7 % | 96.4 % | 10.8 % | N12 3%; N26 0%; N5 28%; N58 0%; typed leaf 69% |
| street unique-str 16K | 23924 | 68.5 | 68.4 % | 96.8 % | 8.6 % | N12 3%; N26 0%; N5 28%; N5+tail 0%; N58 0%; typed leaf 68% |
| street unique-str 212449 | 314837 | 66.9 | 67.4 % | 97.2 % | 6.7 % | N12 2%; N26 0%; N5 30%; N5+tail 0%; N58 0%; set leaf 0%; typed leaf 67% |
