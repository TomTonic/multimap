#!/bin/sh
# What dropping the copy of hash-spilled value sets is worth: the touching
# range scan with the current iteration (art-touch) against the former,
# copying one (art-touch-copy).
set -e
cd "$(dirname "$0")"
. ./procs.sh
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-copy.jsonl
: > "$OUT"
for k in u64 str; do
  for n in 4096 1048576; do
    runp results/mm-copy-$k-$n.log ./results/mmcompare -a art-touch -only art-touch-copy -keys $k -n $n -ops valuesBetween -out $OUT
  done
done
echo 'done'
