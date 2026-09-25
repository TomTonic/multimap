// Package awake keeps the machine from sleeping while a benchmark runs, and
// tells afterwards whether it slept anyway.
//
// Why: an idle machine goes to sleep even on mains power, because nothing
// touches the keyboard during a run of several hours. The processes are
// suspended, not killed, so the run finishes, but the comparisons in flight
// see a machine that stopped and restarted. Every operating system has its own
// way to hold it awake; this package hides them behind Hold.
package awake

import (
	"fmt"
	"os/exec"
	"time"
)

// Hold asks the operating system not to sleep while the process runs: the
// display may turn off, the system stays up. Call it once at the start of a
// long run and call the returned release when the run ends; release is safe
// to call more than once.
//
// On macOS Hold runs caffeinate, on Linux systemd-inhibit, on Windows it calls
// SetThreadExecutionState. It returns an error when the platform has no such
// mechanism or it is missing; the caller should then warn and carry on, and
// Slept will still tell whether the machine slept.
func Hold(why string) (release func(), err error) { return hold(why) }

// Slept returns how long the machine slept since start, which must come from
// time.Now in this process. It compares the wall clock, which keeps running
// while the machine sleeps, with Go's monotonic clock, which on macOS and
// Linux stops. Use it after each measurement to flag the ones a sleep
// interrupted; it returns 0 on systems whose monotonic clock counts sleep.
func Slept(start time.Time) time.Duration {
	now := time.Now()
	s := now.Round(0).Sub(start.Round(0)) - now.Sub(start)
	if s < time.Second { // clock adjustments and rounding, not sleep
		return 0
	}
	return s
}

// holdCommand starts a helper that holds the machine awake until it is
// killed, for the platforms that offer only a command for it.
func holdCommand(name string, args ...string) (func(), error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return nil, fmt.Errorf("cannot keep the machine awake: %w", err)
	}
	cmd := exec.Command(path, args...)
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("cannot keep the machine awake: %w", err)
	}
	done := false
	return func() {
		if !done {
			done = true
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		}
	}, nil
}
