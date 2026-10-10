#!/bin/bash
D=/tmp/claude-1000/desc; cd /mnt/c/temp/code/multi_map-layout/bench
for k in str email url path street dirs links; do for n in 4096 16384 65536; do
  MKLOOK=1 MKLOOK_KEYS=$k MKLOOK_VALUES=single-value MKLOOK_N=$n MKLOOK_SECS=10 MKLOOK_PROF=$D/look $D/prof.test -test.run TestLookProbe -test.count=1 > $D/out-$k-$n.txt 2>&1
  echo "$(date +%H:%M:%S) $k $n" >> $D/progress.txt
done; done
echo done >> $D/progress.txt
