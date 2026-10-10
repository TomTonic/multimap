#!/bin/bash
D=/tmp/claude-1000/wprof; cd /mnt/c/temp/code/multi_map-layout/bench
for k in str email url path street dirs links; do for v in single-value natural; do for n in 4096 16384; do
  MKPROBE=1 MKPROBE_KEYS=$k MKPROBE_VALUES=$v MKPROBE_N=$n MKPROBE_CYCLES=1 MKPROBE_SECS=10 MKPROBE_PROF=$D/write $D/prof.test -test.run 'TestProbe$' -test.count=1 -test.timeout 0 > $D/out-$k-$v-$n.txt 2>&1
  echo "$(date +%H:%M:%S) $k $v $n" >> $D/progress.txt
done; done; done
echo done >> $D/progress.txt
