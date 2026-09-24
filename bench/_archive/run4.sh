#!/bin/sh
# ValuesBetween after the structural upper bound: new ART scan against the
# former one (art-linear) and both B-tree variants.
set -e
cd "$(dirname "$0")"
. ./procs.sh
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-range2.jsonl
: > "$OUT"
for k in u64 str; do
  for n in 4096 1048576; do
    runp results/mm-range2-$k-$n.log ./results/mmcompare -keys $k -n $n -ops valuesBetween -out $OUT
  done
done
echo 'done'
