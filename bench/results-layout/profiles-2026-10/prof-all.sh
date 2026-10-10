#!/bin/bash
P=/tmp/claude-1000/prof-all; T=/tmp/claude-1000/prof.test
rm -rf $P; mkdir -p $P/out
for k in u64 str uuid email url path street dirs links; do for v in single-value natural; do for n in 4096 16384 65536; do
  c=40
  MKPROBE=1 MKPROBE_KEYS=$k MKPROBE_VALUES=$v MKPROBE_N=$n MKPROBE_CYCLES=$c MKPROBE_PROF=$P/write $T -test.run 'TestProbe$' -test.count=1 > $P/out/probe-$k-$v-$n.txt 2>&1
  MKLOOK=1 MKLOOK_KEYS=$k MKLOOK_VALUES=$v MKLOOK_N=$n MKLOOK_PROF=$P/look $T -test.run TestLookProbe -test.count=1 > $P/out/look-$k-$v-$n.txt 2>&1
  MKRANGE=1 MKRANGE_KEYS=$k MKRANGE_VALUES=$v MKRANGE_N=$n MKRANGE_PROF=$P/range $T -test.run TestRangeProbe -test.count=1 > $P/out/range-$k-$v-$n.txt 2>&1
  echo "$(date +%H:%M:%S) $k $v $n" >> $P/progress.txt
done; done; done
echo done >> $P/progress.txt
