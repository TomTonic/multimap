package awake

import (
	"runtime"
	"testing"
	"time"
)

// TestHold makes sure a benchmark run can keep the machine awake and let it
// sleep again afterwards. It covers the platform-specific holder of package
// awake: Hold succeeds where the platform's mechanism exists, and release can
// be called twice.
func TestHold(t *testing.T) {
	release, err := Hold("test")
	if err != nil {
		if runtime.GOOS == "darwin" || runtime.GOOS == "windows" {
			t.Fatal(err) // the mechanism is part of the system
		}
		t.Skipf("no mechanism on this machine: %v", err)
	}
	release()
	release()
}

// TestSlept makes sure a run that did not sleep is not flagged as having
// slept. It covers the sleep detection of package awake, which compares the
// wall clock with the monotonic clock: without a sleep in between, both agree.
func TestSlept(t *testing.T) {
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	if s := Slept(start); s != 0 {
		t.Errorf("Slept = %v without sleep", s)
	}
}
