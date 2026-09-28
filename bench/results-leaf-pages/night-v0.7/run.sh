#!/bin/bash
# Overnight run 2026-09-27/28 with rtcompare v0.7.0. A is always main's Ordered
# (bench in the main worktree); B is the page variant (baseline copy) or the
# other candidates.
cd /mnt/c/temp/code/multi_map-bench-main/bench
B=/tmp/claude-1000/-mnt-c-temp-code-multi-map/87b484e1-de3b-4329-a0cb-56af7472d941/scratchpad/bin
O=/tmp/claude-1000/-mnt-c-temp-code-multi-map/87b484e1-de3b-4329-a0cb-56af7472d941/scratchpad/night
COMMON="-suite dev -buildmax 65536"
echo "start $(date +%FT%T)" >> $O/done.txt
$B/bench-vs-e1f6ad3 $COMMON -vs baseline -sizes 4096,16384,262144,1048576 -minprocs 5 -maxprocs 10 -memn 1048576 -memrounds 3 -out $O/lp-e1f6ad3 > $O/lp-e1f6ad3.log 2>&1
echo "lp-e1f6ad3 exit $? $(date +%FT%T)" >> $O/done.txt
for v in 8052a18 afd1137 7157f4b 527d9b7; do
  $B/bench-vs-$v $COMMON -vs baseline -keys u64,str,uuid,path,street -sizes 4096,262144 -minprocs 5 -maxprocs 10 -memn 262144 -memrounds 1 -out $O/lp-$v > $O/lp-$v.log 2>&1
  echo "lp-$v exit $? $(date +%FT%T)" >> $O/done.txt
done
$B/bench-main $COMMON -sizes 4096,16384,262144,1048576 -minprocs 5 -maxprocs 6 -memn 1048576 -memrounds 3 -out $O/main > $O/main.log 2>&1
echo "main exit $? $(date +%FT%T)" >> $O/done.txt
