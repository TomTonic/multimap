#!/bin/sh
# Runs every benchmark script in order and logs when each starts and ends.
# RT and MEMGC pass extra flags to the rtcompare commands and to memgc, e.g.
#   RT='-loopscale 1.5 -repeats 151' MEMGC='-cyclescale 1.5' ./runall.sh
set -e
cd "$(dirname "$0")"
export RT MEMGC
for s in run.sh run2.sh run3.sh run4.sh run5.sh run6.sh run7.sh run8.sh run9.sh run10.sh run11.sh run12.sh run13.sh run14.sh run15.sh; do
  echo "=== $s start $(date '+%F %T')"
  ./$s
  echo "=== $s end   $(date '+%F %T')"
done
echo "=== all done $(date '+%F %T')"
