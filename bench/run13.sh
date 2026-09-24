#!/bin/sh
# run12.sh again with 151 instead of 101 samples per candidate, so that more
# measurements are comparable with one another.
set -e
cd "$(dirname "$0")"
. ./procs.sh
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-lib-touch3.jsonl
: > "$OUT"
for k in u64 str; do
  for n in 4096 1048576; do
    runp results/mm-lib-touch3-$k-$n.log ./results/mmcompare -a lib -only art,art-touch -loopscale 1.5 -repeats 151 -keys $k -n $n -ops valuesBetween -out $OUT
  done
done
echo 'done'
