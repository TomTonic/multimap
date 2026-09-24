#!/bin/sh
# Is the library's lag behind the prototype's touching scan a matter of heap
# layout? The same pairs with the library built last (as before) and first.
# This deliberately fixes the build order, so it runs one process per order
# with -layoutseed 0 instead of going through runp.
set -e
cd "$(dirname "$0")"
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-order.jsonl
: > "$OUT"
for n in 4096 1048576; do
  for first in false true; do
    # shellcheck disable=SC2086 # RT holds several flags
    ./results/mmcompare -a lib -only art-touch -loopscale 1.5 -repeats 151 -libfirst=$first -keys str -n $n -ops valuesBetween -out $OUT -layoutseed 0 $RT 2> results/mm-order-str-$n-libfirst-$first.log
  done
done
echo 'done'
