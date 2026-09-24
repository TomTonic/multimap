#!/bin/sh
# Runs every rtcompare scenario, one process per scenario, and collects the
# JSON lines in results/results.jsonl (human-readable reports in results/*.log).
set -e
cd "$(dirname "$0")"
. ./procs.sh
mkdir -p results
go build -o results/compare ./cmd/compare
OUT=results/results.jsonl
: > "$OUT"
for k in u64 str; do
  runp results/$k-4096.log ./results/compare -keys $k -n 4096    -ops get,miss,scan,build -out $OUT
  runp results/$k-1M.log ./results/compare -keys $k -n 1048576 -ops get,miss,scan       -out $OUT
  runp results/$k-65536.log ./results/compare -keys $k -n 65536   -ops build               -out $OUT
done
echo 'done'
