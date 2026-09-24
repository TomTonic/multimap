#!/bin/sh
# Where does touching children start to pay off? art-touch against art at
# sizes between the cached 4K and the uncached 1M of run9.sh.
set -e
cd "$(dirname "$0")"
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-touch-size.jsonl
: > "$OUT"
for k in u64 str; do
  for n in 16384 65536 262144; do
    ./results/mmcompare -a art-touch -only art -keys $k -n $n -ops valuesBetween -out $OUT $RT 2> results/mm-touch-size-$k-$n.log
  done
done
echo done
