#!/bin/sh
# run11.sh again with 50% longer batches (operations per sample), on an idle
# machine, to confirm the library's touching range scan.
set -e
cd "$(dirname "$0")"
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-lib-touch2.jsonl
: > "$OUT"
for k in u64 str; do
  for n in 4096 1048576; do
    ./results/mmcompare -a lib -only art,art-touch -loopscale 1.5 -keys $k -n $n -ops valuesBetween -out $OUT 2> results/mm-lib-touch2-$k-$n.log
  done
done
echo done
