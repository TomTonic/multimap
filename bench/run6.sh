#!/bin/sh
# New ART layout (node25 + one 128-byte leaf with 9 inline values, mmart2)
# against the current one (80-byte leaf with 3 inline values, mmart).
set -e
cd "$(dirname "$0")"
go build -o results/mmcompare ./cmd/mmcompare
go build -o results/memgc ./cmd/memgc
OUT=results/mm-v2.jsonl
: > "$OUT"
for k in u64 str; do
  ./results/mmcompare -a art-v2 -only art -keys $k -n 4096    -ops valuesFor,valuesBetween,addRemove,build -out $OUT 2> results/mm-v2-$k-4096.log
  ./results/mmcompare -a art-v2 -only art -keys $k -n 1048576 -ops valuesFor,valuesBetween,addRemove       -out $OUT 2> results/mm-v2-$k-1M.log
done
: > results/mm-v2-memgc.jsonl
for k in u64 str; do
  for impl in none mm-art mm-art-v2; do
    ./results/memgc -impl $impl -keys $k >> results/mm-v2-memgc.jsonl
  done
done
echo done
