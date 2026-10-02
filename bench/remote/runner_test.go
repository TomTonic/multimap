package remote

import (
	"os/exec"
	"strings"
	"testing"
)

// arm runs bench/remote/arm-run.sh with args and returns its output.
func arm(t *testing.T, env []string, args ...string) (string, error) {
	t.Helper()
	if _, err := exec.LookPath("bash"); err != nil {
		t.Skip("no bash")
	}
	cmd := exec.Command("bash", append([]string{"arm-run.sh", "--no-fetch", "--queue", "testdata/queue.txt"}, args...)...)
	cmd.Env = append(cmd.Environ(), env...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// TestRunner makes sure that the person who owns the arm64 machine can see
// what a measuring job will do before it runs. It belongs to the job queue
// that lets the agent hand measurements to a second machine
// (docs/redesign/MEASURING.md), and exercises the runner in its dry-run mode
// on a test queue: which job is picked, what the script tells about the
// machine, the duration and the expected end, and the command it would run.
func TestRunner(t *testing.T) {
	for _, tc := range []struct {
		name string
		env  []string
		args []string
		want []string
		not  []string
	}{
		{
			name: "picks the first job without results and says when it ends",
			env:  []string{"ARM_RUN_DONE="},
			args: []string{"--dry-run"},
			want: []string{"job:       t1  ref HEAD  baseline -", "duration:  about 10 min", "expected end", "Do not use the machine", "-suite dev -keys u64 -sizes 4096 -vs hashed -skipmem", "-tags ''"},
		},
		{
			name: "skips the jobs that have results",
			env:  []string{"ARM_RUN_DONE=t1"},
			args: []string{"--dry-run"},
			want: []string{"job:       t2"},
			not:  []string{"job:       t1"},
		},
		{
			name: "builds the baseline and the tags a job asks for",
			env:  []string{"ARM_RUN_DONE=t1"},
			args: []string{"--dry-run", "t2"},
			want: []string{"job:       t2", "tags 'strvals'", "with the baseline", "-tags 'baseline,strvals'", "-suite dev -keys u64,str -sizes 4096,16384 -vs baseline"},
			not:  []string{"# a comment"},
		},
		{
			name: "lists the queue with what is done",
			env:  []string{"ARM_RUN_DONE=t1 t3"},
			args: []string{"--list"},
			want: []string{"t1     done", "t2     open", "t3     done"},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out, err := arm(t, tc.env, tc.args...)
			if err != nil {
				t.Fatalf("%v:\n%s", err, out)
			}
			for _, w := range tc.want {
				if !strings.Contains(out, w) {
					t.Errorf("the output has no %q:\n%s", w, out)
				}
			}
			for _, n := range tc.not {
				if strings.Contains(out, n) {
					t.Errorf("the output has %q:\n%s", n, out)
				}
			}
		})
	}
}

// TestRunnerErrors makes sure that the runner stops with a clear message when
// it has nothing to run, and does not start a job it cannot understand. It
// belongs to the job queue for the arm64 machine (docs/redesign/MEASURING.md):
// the person who runs it should never wait for a benchmark that cannot start.
func TestRunnerErrors(t *testing.T) {
	for _, tc := range []struct {
		name string
		env  []string
		args []string
		want string
	}{
		{"says that no job is open when all have results", []string{"ARM_RUN_DONE=t1 t2 t3"}, []string{"--dry-run"}, "no open job in the queue"},
		{"says that a job is not in the queue", []string{"ARM_RUN_DONE="}, []string{"--dry-run", "nope"}, "no job nope in the queue"},
		{"rejects an unknown option", nil, []string{"--nope"}, "unknown option --nope"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			out, err := arm(t, tc.env, tc.args...)
			if err == nil || !strings.Contains(out, tc.want) {
				t.Errorf("err %v, output %q, want a failure with %q", err, out, tc.want)
			}
		})
	}
}
