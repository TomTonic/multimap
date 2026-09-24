#!/bin/sh
# Follow-up runs: memory/GC per structure (one process each), node search
# strategies, and the flat arena against the pointer ART.
set -e
cd "$(dirname "$0")"
. ./procs.sh
go build -o results/memgc ./cmd/memgc
go build -o results/nodesearch ./cmd/nodesearch
go build -o results/compare ./cmd/compare
: > results/memgc.jsonl
for k in u64 str; do
  for impl in none arena-art arena-flat ptr-art tidwall-btree plar-art plar-hot go-map; do
    ./results/memgc -impl $impl -keys $k $MEMGC >> results/memgc.jsonl
  done
done
: > results/nodesearch.jsonl
runp results/nodesearch.log ./results/nodesearch -out results/nodesearch.jsonl
OUT=results/flat.jsonl
: > "$OUT"
for k in u64 str; do
  runp results/flat-$k-4096.log ./results/compare -a arena-flat -only ptr-art,arena-art -keys $k -n 4096    -ops get,miss,scan,build -out $OUT
  runp results/flat-$k-1M.log ./results/compare -a arena-flat -only ptr-art,arena-art -keys $k -n 1048576 -ops get,miss,scan       -out $OUT
done
echo 'done'
