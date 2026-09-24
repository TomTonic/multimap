#!/bin/sh
# Realistic multimap comparison: value sets per key, iterator results, ART
# multimap against tidwall/btree used both ways (values inline / behind a
# pointer). Then memory and GC of the three multimaps, one process each.
set -e
cd "$(dirname "$0")"
. ./procs.sh
go build -o results/mmcompare ./cmd/mmcompare
go build -o results/memgc ./cmd/memgc
OUT=results/mm.jsonl
: > "$OUT"
for k in u64 str; do
  runp results/mm-$k-4096.log ./results/mmcompare -keys $k -n 4096    -ops valuesFor,valuesBetween,addRemove,build -out $OUT
  runp results/mm-$k-1M.log ./results/mmcompare -keys $k -n 1048576 -ops valuesFor,valuesBetween,addRemove       -out $OUT
done
: > results/mm-memgc.jsonl
for k in u64 str; do
  for impl in none mm-art mm-btree-inline mm-btree-ptr; do
    ./results/memgc -impl $impl -keys $k $MEMGC >> results/mm-memgc.jsonl
  done
done
echo 'done'
