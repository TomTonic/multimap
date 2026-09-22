#!/bin/sh
# Value container: array spill (vset.Set) against hash spill (vset.HashSpill).
# Memory/GC of the containers alone and of the whole multimaps, then speed.
set -e
cd "$(dirname "$0")"
go build -o results/memgc ./cmd/memgc
go build -o results/vsetcompare ./cmd/vsetcompare
: > results/vset-memgc.jsonl
for impl in none vset-hash vset-array mm-art mm-btree-inline mm-btree-ptr; do
  ./results/memgc -impl $impl -keys u64 >> results/vset-memgc.jsonl
done
for impl in none mm-art mm-btree-inline mm-btree-ptr; do
  ./results/memgc -impl $impl -keys str >> results/vset-memgc.jsonl
done
: > results/vset.jsonl
./results/vsetcompare -n 4096    -out results/vset.jsonl 2> results/vset-4096.log
./results/vsetcompare -n 1048576 -out results/vset.jsonl 2> results/vset-1M.log
echo done
