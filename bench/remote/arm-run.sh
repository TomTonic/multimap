#!/usr/bin/env bash
# arm-run.sh: run one job of the benchmark queue on this machine and push the
# results to the branch arm-results.
#
# For the user of the arm64 machine (macOS or Linux; the machine is yours, the
# script only runs when you start it):
#
#   1. Make sure Go (the version of go.mod) and git with push rights are
#      installed, and that you are in a clone of the repository.
#   2. Plug the machine in, close what you can, and do not use it for the
#      duration the script names: a busy machine measures nothing. The script
#      warns about battery power and a high load.
#   3. Run   bench/remote/arm-run.sh          the first job without results
#            bench/remote/arm-run.sh a1       the job a1
#            bench/remote/arm-run.sh --list   the queue and what is done
#
# The script fetches the queue (bench/remote/queue.txt on the branch cacheline)
# and the commit the job names, builds that commit in a temporary worktree (your
# working tree is not touched), runs the benchmark under caffeinate (macOS), and
# commits the results with a description of the machine to the branch
# arm-results of the remote and pushes it. Nothing else is pushed.
#
# Queue format, one job per line, # starts a comment:
#   <id> <ref> <baseline-ref or -> <duration-minutes> [tags=a,b] [env=A=1,B=2] <arguments of cmd/bench>
#   <id> <ref> - <duration-minutes> [env=A=1,B=2] gotest <packages> <flags of go test>
# The ref and the baseline ref are commits or branches of the remote. A
# baseline ref makes the script build the library as of that ref into the
# bench (cmd/mkbaseline) and run it with the build tag baseline; tags= adds
# other build tags, such as strvals; env= sets environment variables for the run.
# A job whose arguments start with gotest runs  go test <packages> <flags>  in
# the commit's worktree instead of cmd/bench (for the microbenchmarks of the
# internal packages, such as  gotest ./internal/vpage ./internal/lpage -run ^$
# -bench BenchmarkGet -benchtime 2s ); its output is run.log. The duration is an
# estimate that the script shows with the expected end.
#
# Options: --dry-run (show the plan, run nothing), --no-push (keep the results
# in a temporary directory), --no-fetch, --yes (do not ask about warnings),
# --queue FILE (read the queue from a file), --branch NAME (branch of the queue,
# default cacheline), --remote NAME (default origin).
# ARM_MACHINE names the machine in the results; ARM_RUN_DONE ("a0 a1") replaces
# the look into arm-results for which jobs are done (for tests).

set -eu
set -o pipefail

remote=origin qbranch=cacheline rbranch=arm-results
dry=0 push=1 fetch=1 yes=0 list=0 queue_file="" want=""

usage() { sed -n '2,/^$/p' "$0" | sed 's/^# \{0,1\}//'; }

while [ $# -gt 0 ]; do
	case "$1" in
	--dry-run) dry=1 ;;
	--no-push) push=0 ;;
	--no-fetch) fetch=0 ;;
	--yes) yes=1 ;;
	--list) list=1 ;;
	--queue) queue_file=$2; shift ;;
	--branch) qbranch=$2; shift ;;
	--remote) remote=$2; shift ;;
	-h | --help) usage; exit 0 ;;
	-*) echo "arm-run.sh: unknown option $1" >&2; exit 2 ;;
	*) want=$1 ;;
	esac
	shift
done

if [ -n "$queue_file" ]; then queue_file=$(cd "$(dirname "$queue_file")" && pwd)/$(basename "$queue_file"); fi
root=$(git rev-parse --show-toplevel)
cd "$root"
if [ "$fetch" = 1 ]; then git fetch -q "$remote"; fi

# queue returns the text of the queue.
queue() {
	if [ -n "$queue_file" ]; then cat "$queue_file"; else git show "$remote/$qbranch:bench/remote/queue.txt"; fi
}

# done_ids lists the ids that have results on the results branch.
done_ids() {
	if [ -n "${ARM_RUN_DONE+x}" ]; then echo "$ARM_RUN_DONE" | tr ' ' '\n'; return; fi
	git ls-tree -r --name-only "$remote/$rbranch" 2>/dev/null | sed -n 's|^\([^/]*\)/DONE$|\1|p' || true
}

# resolve prints the commit of ref, as a branch of the remote or as it is.
resolve() {
	git rev-parse --verify -q "$remote/$1^{commit}" || git rev-parse --verify -q "$1^{commit}" || {
		echo "arm-run.sh: cannot resolve $1" >&2
		return 1
	}
}

# queued prints the jobs of the queue without comments and empty lines.
queued() { queue | sed -e 's/#.*//' -e '/^[[:space:]]*$/d'; }

done_list=$(done_ids)
is_done() { echo "$done_list" | grep -qx -- "$1"; }

if [ "$list" = 1 ]; then
	queued | while read -r id ref base mins args; do
		state=open
		if is_done "$id"; then state=done; fi
		printf '%-6s %-5s %4s min  %s (baseline %s)  %s\n' "$id" "$state" "$mins" "$ref" "$base" "$args"
	done
	exit 0
fi

line=""
while IFS= read -r l; do
	set -f
	# shellcheck disable=SC2086
	set -- $l
	set +f
	id=${1:-}
	if [ -n "$want" ]; then
		if [ "$id" = "$want" ]; then line=$l; break; fi
	elif ! is_done "$id"; then
		line=$l
		break
	fi
done <<EOF2
$(queued)
EOF2
if [ -z "$line" ]; then
	if [ -n "$want" ]; then echo "arm-run.sh: no job $want in the queue" >&2; else echo "no open job in the queue"; fi
	exit 1
fi
read -r id ref base mins args <<EOF2
$line
EOF2
case "$id" in *[!A-Za-z0-9._-]*) echo "arm-run.sh: bad job id $id" >&2; exit 2 ;; esac
case "$mins" in '' | *[!0-9]*) echo "arm-run.sh: bad duration $mins in job $id" >&2; exit 2 ;; esac
tags=""
case "$args" in tags=*) tags=${args%% *}; tags=${tags#tags=}; case "$args" in *' '*) args=${args#* } ;; *) args="" ;; esac ;; esac
envs=""
case "$args" in env=*) envs=${args%% *}; envs=${envs#env=}; envs=$(echo "$envs" | tr ',' ' '); case "$args" in *' '*) args=${args#* } ;; *) args="" ;; esac ;; esac
gotest=0 gargs=""
case "$args" in gotest | gotest\ *) gotest=1; gargs=${args#gotest}; gargs=${gargs# } ;; esac

# --- the machine -----------------------------------------------------------

os=$(uname -s)
if [ "$os" = Darwin ]; then
	cpu=$(sysctl -n machdep.cpu.brand_string 2>/dev/null || echo unknown)
	line_size=$(sysctl -n hw.cachelinesize 2>/dev/null || echo unknown)
	cores=$(sysctl -n hw.physicalcpu 2>/dev/null || echo unknown)
	perf=$(sysctl -n hw.perflevel0.physicalcpu 2>/dev/null || echo unknown)
	mem=$(( $(sysctl -n hw.memsize 2>/dev/null || echo 0) / 1073741824 ))
	power=$(pmset -g batt 2>/dev/null | head -1 || echo unknown)
else
	cpu=$(sed -n 's/^model name[[:space:]]*: //p;s/^Model[[:space:]]*: //p' /proc/cpuinfo 2>/dev/null | head -1)
	if [ -z "$cpu" ]; then cpu=$(lscpu 2>/dev/null | sed -n 's/^Model name:[[:space:]]*//p' | head -1); fi
	line_size=$(getconf LEVEL1_DCACHE_LINESIZE 2>/dev/null || echo unknown)
	cores=$(nproc 2>/dev/null || echo unknown)
	perf=$cores
	mem=$(( $(sed -n 's/^MemTotal:[[:space:]]*\([0-9]*\) kB/\1/p' /proc/meminfo 2>/dev/null || echo 0) / 1048576 ))
	power="not checked"
fi
machine=${ARM_MACHINE:-${cpu:-arm64} ${mem}GB $os}
load() { LC_ALL=C uptime | sed 's/.*load averages*: *//'; }

# when prints the clock time in $1 minutes.
when() { date -v+"$1"M '+%H:%M' 2>/dev/null || date -d "+$1 minutes" '+%H:%M'; }

warn=""
case "$power" in *Battery*) warn="$warn the machine runs on battery;" ;; esac
# a machine that was busy a moment ago is not at rest, whatever it does now
if [ "$(load | awk -F'[ ,]+' '{ print ($1 >= 2 || $2 >= 3) ? 1 : 0 }')" = 1 ]; then warn="$warn the load average is $(load);"; fi
load_before=$(load)

echo "machine:   $machine (cache line $line_size B, $perf performance cores of $cores, power: $power)"
echo "job:       $id  ref $ref  baseline $base  tags '${tags:-none}'  env '${envs:-none}'"
echo "arguments: $args"
echo "duration:  about $mins min; now $(date '+%H:%M'), expected end $(when "$mins")"
echo "load:      $(load)"
if [ -n "$warn" ]; then echo "WARNING:  $warn do not start now, or accept an invalid measurement."; fi
echo "Do not use the machine until the script is done."

sha=$(resolve "$ref")
basesha=""
if [ "$base" != "-" ]; then basesha=$(resolve "$base"); fi
build_tags=$tags
if [ -n "$basesha" ]; then build_tags="baseline${build_tags:+,$build_tags}"; fi

if [ "$dry" = 1 ]; then
	if [ "$gotest" = 1 ]; then
		echo "dry run: would check out commit $sha in a temporary worktree,"
		echo "dry run: run  [caffeinate -i] ${envs:+env $envs }go test $gargs"
	else
		echo "dry run: would build commit $sha${basesha:+ with the baseline $basesha} in a temporary worktree,"
		echo "dry run: run  [caffeinate -i] bench -tags '$build_tags' $args -out <dir>"
	fi
	echo "dry run: and push the results of $id to $remote/$rbranch"
	exit 0
fi

if [ -n "$warn" ] && [ "$yes" = 0 ] && [ -t 0 ]; then
	printf 'Start anyway? [y/N] '
	read -r answer
	case "$answer" in y | Y | yes) ;; *) echo aborted; exit 1 ;; esac
fi

# --- build and run ---------------------------------------------------------

tmp=$(mktemp -d "${TMPDIR:-/tmp}/arm-run.XXXXXX")
keep=0
cleanup() {
	git worktree remove --force "$tmp/src" >/dev/null 2>&1 || true
	git worktree remove --force "$tmp/res" >/dev/null 2>&1 || true
	git worktree prune >/dev/null 2>&1 || true
	[ "$keep" = 1 ] || rm -rf "$tmp"
}
trap cleanup EXIT

out=$tmp/out
mkdir -p "$out"
git worktree add -q --detach "$tmp/src" "$sha"
if [ "$gotest" = 1 ]; then
	# compile the test binaries now, so that the compiler does not run during the measurement
	(
		cd "$tmp/src"
		set -f
		for p in $gargs; do
			case "$p" in -*) break ;; esac
			go test -c -o /dev/null "$p"
		done
	)
else
	(
		cd "$tmp/src/bench"
		if [ -n "$basesha" ]; then go run ./cmd/mkbaseline -ref "$basesha"; fi
		go build ${build_tags:+-tags "$build_tags"} -o "$tmp/bench.bin" ./cmd/bench
	)
fi
# the build has just loaded the machine: let it settle before measuring
echo "settling for ${ARM_SETTLE:-30} s after the build ..."
sleep "${ARM_SETTLE:-30}"
{
	echo "job:        $id"
	echo "commit:     $sha"
	echo "baseline:   ${basesha:--}"
	echo "tags:       ${build_tags:-none}"
	echo "arguments:  $args"
	echo "machine:    $machine"
	echo "system:     $(uname -srm)"
	echo "cpu:        $cpu ($perf performance cores of $cores)"
	echo "memory:     ${mem} GB"
	echo "cache line: $line_size bytes"
	echo "go:         $(go version)"
	echo "power:      $power"
	echo "load before the build: $load_before"
	echo "load at the start:     $(load)"
	echo "started:    $(date '+%Y-%m-%d %H:%M:%S %z')"
	if [ -n "$warn" ]; then echo "warnings:  $warn"; fi
} >"$out/env.txt"
echo "$args" >"$out/args.txt"

prefix=""
if command -v caffeinate >/dev/null 2>&1; then prefix="caffeinate -i"; fi
status=0
set -f +e
if [ "$gotest" = 1 ]; then
	cd "$tmp/src"
	# shellcheck disable=SC2086
	env $envs $prefix go test $gargs 2>&1 | tee "$out/run.log"
else
	cd "$tmp/src/bench"
	# shellcheck disable=SC2086
	env $envs $prefix "$tmp/bench.bin" $args -out "$out/bench-out" 2>&1 | tee "$out/run.log"
fi
status=${PIPESTATUS[0]}
set +f -e
{
	echo "load end:   $(load)"
	echo "finished:   $(date '+%Y-%m-%d %H:%M:%S %z')"
	echo "exit code:  $status"
} >>"$out/env.txt"
if [ "$status" != 0 ]; then
	keep=1
	echo "arm-run.sh: the benchmark failed (exit $status); results kept in $out" >&2
	exit "$status"
fi
echo ok >"$out/DONE"

# --- publish ---------------------------------------------------------------

cd "$root"
if git rev-parse --verify -q "$remote/$rbranch" >/dev/null; then
	git worktree add -q -B "$rbranch" "$tmp/res" "$remote/$rbranch"
else
	git worktree add -q --detach "$tmp/res" HEAD
	(cd "$tmp/res" && git checkout -q --orphan "$rbranch" && git rm -rq --cached . && find . -mindepth 1 -maxdepth 1 ! -name .git -exec rm -rf {} +)
	cat >"$tmp/res/README.md" <<'EOF2'
# Results of the arm64 jobs

One directory per job of bench/remote/queue.txt (on the branch cacheline): env.txt
describes the machine, args.txt the arguments, run.log the output, bench-out the
results of cmd/bench. The file DONE marks a finished job. This branch holds results only.
EOF2
fi
rm -rf "${tmp:?}/res/$id"
cp -R "$out" "$tmp/res/$id"
(
	cd "$tmp/res"
	git add -A
	git commit -q -m "Add results of job $id from $machine" -m "Commit $sha, baseline ${basesha:--}, arguments: $args"
)
if [ "$push" = 1 ]; then
	if git -C "$tmp/res" push "$remote" "$rbranch"; then
		echo "pushed the results of $id to $remote/$rbranch"
	else
		keep=1
		echo "arm-run.sh: push failed; the results are committed in $tmp/res (branch $rbranch): push them by hand" >&2
		exit 1
	fi
else
	keep=1
	echo "results of $id are in $tmp/res (branch $rbranch), not pushed"
fi
