#!/bin/sh
# HOT (plar/go-hot-trie) against the ART (ptr-art) with rtcompare, then
# memory/GC of every structure: three round-robin rounds with 50 GC cycles
# each, so the median is robust against the baseline noise of single runs.
set -e
cd "$(dirname "$0")"
. ./procs.sh
go build -o results/compare ./cmd/compare
go build -o results/memgc ./cmd/memgc
OUT=results/hot.jsonl
: > "$OUT"
for k in u64 str; do
  runp results/hot-$k-4096.log ./results/compare -a ptr-art -only plar-hot -keys $k -n 4096    -ops get,miss,build -out $OUT
  runp results/hot-$k-1M.log ./results/compare -a ptr-art -only plar-hot -keys $k -n 1048576 -ops get,miss       -out $OUT
done
: > results/memgc3.jsonl
for round in 1 2 3; do
  for k in u64 str; do
    for impl in none ptr-art tidwall-btree plar-art plar-hot go-map mm-art mm-btree-inline mm-btree-ptr; do
      ./results/memgc -impl $impl -keys $k -cycles 50 $MEMGC >> results/memgc3.jsonl
    done
  done
done
echo 'done'
