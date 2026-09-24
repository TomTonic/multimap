# procs.sh is sourced by the run scripts.
#
# runp LOG COMMAND [ARGS...] runs one comparison command in $PROCS processes
# (default 5), passing -layoutseed 1, 2, ... so that every process builds its
# fixtures in a different heap layout, plus the extra flags in $RT. All
# processes append to the same JSON lines file (-out in ARGS) and write their
# reports to LOG. cmd/summarize pools the processes of each comparison.
#
# Why several processes: rtcompare's interval covers the noise within one
# process only. For large fixtures the heap layout of a process shifts results
# by several points (see README.md), so one process is one sample of layout.
runp() {
  log=$1
  shift
  : > "$log"
  i=1
  while [ "$i" -le "${PROCS:-5}" ]; do
    echo "## process $i (-layoutseed $i)" >> "$log"
    # shellcheck disable=SC2086 # RT holds several flags
    "$@" -layoutseed "$i" $RT 2>> "$log"
    i=$((i + 1))
  done
}
