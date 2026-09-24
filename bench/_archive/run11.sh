#!/bin/sh
# The library's range scan, which now always touches children ahead (skipping
# nodes with fewer than two children in range), against the prototype's plain
# scan (art) and its unconditional touching scan (art-touch).
set -e
cd "$(dirname "$0")"
. ./procs.sh
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-lib-touch.jsonl
: > "$OUT"
for k in u64 str; do
  for n in 4096 1048576; do
    runp results/mm-lib-touch-$k-$n.log ./results/mmcompare -a lib -only art,art-touch -keys $k -n $n -ops valuesBetween -out $OUT
  done
done
echo 'done'
