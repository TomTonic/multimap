#!/bin/sh
# Is the library's lag behind the prototype's touching scan a matter of heap
# layout? The same pairs with the library built last (as before) and first.
set -e
cd "$(dirname "$0")"
go build -o results/mmcompare ./cmd/mmcompare
OUT=results/mm-order.jsonl
: > "$OUT"
for n in 4096 1048576; do
  for first in false true; do
    ./results/mmcompare -a lib -only art-touch -loopscale 1.5 -repeats 151 -libfirst=$first -keys str -n $n -ops valuesBetween -out $OUT $RT 2> results/mm-order-str-$n-libfirst-$first.log
  done
done
echo done
