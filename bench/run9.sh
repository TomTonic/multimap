#!/bin/sh
# Range scan that touches all children of a node before descending, so their
# cache misses overlap (art-touch), against the plain scan and btree-inline.
set -e
cd "$(dirname "$0")"
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-touch.jsonl
: > "$OUT"
for k in u64 str; do
  for n in 4096 1048576; do
    ./results/mmcompare -a art-touch -only art,btree-inline -keys $k -n $n -ops valuesBetween -out $OUT $RT 2> results/mm-touch-$k-$n.log
  done
done
echo done
