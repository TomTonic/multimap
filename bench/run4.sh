#!/bin/sh
# ValuesBetween after the structural upper bound: new ART scan against the
# former one (art-linear) and both B-tree variants.
set -e
cd "$(dirname "$0")"
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-range2.jsonl
: > "$OUT"
for k in u64 str; do
  for n in 4096 1048576; do
    ./results/mmcompare -keys $k -n $n -ops valuesBetween -out $OUT $RT 2> results/mm-range2-$k-$n.log
  done
done
echo done
