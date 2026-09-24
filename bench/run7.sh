#!/bin/sh
# The library's Ordered multimap (the prototype ported into the module, plus
# node25, deletion and the public iterator API) against the mmart prototype.
set -e
cd "$(dirname "$0")"
. ./procs.sh
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-lib.jsonl
: > "$OUT"
for k in u64 str; do
  runp results/mm-lib-$k-4096.log ./results/mmcompare -a lib -only art -keys $k -n 4096    -ops valuesFor,valuesBetween,addRemove,build -out $OUT
  runp results/mm-lib-$k-1M.log ./results/mmcompare -a lib -only art -keys $k -n 1048576 -ops valuesFor,valuesBetween,addRemove       -out $OUT
done
echo 'done'
