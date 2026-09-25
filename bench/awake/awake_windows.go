package awake

import (
	"fmt"
	"runtime"
	"syscall"
)

const (
	esContinuous     = 0x80000000
	esSystemRequired = 0x00000001
)

// hold calls SetThreadExecutionState. The request belongs to the calling
// thread, so one goroutine locked to its thread makes it and clears it again.
func hold(string) (func(), error) {
	proc := syscall.NewLazyDLL("kernel32.dll").NewProc("SetThreadExecutionState")
	if err := proc.Find(); err != nil {
		return nil, fmt.Errorf("cannot keep the machine awake: %w", err)
	}
	errs, stop := make(chan error), make(chan struct{})
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		if r, _, err := proc.Call(esContinuous | esSystemRequired); r == 0 {
			errs <- fmt.Errorf("cannot keep the machine awake: %w", err)
			return
		}
		errs <- nil
		<-stop
		_, _, _ = proc.Call(esContinuous)
	}()
	if err := <-errs; err != nil {
		return nil, err
	}
	done := false
	return func() {
		if !done {
			done = true
			stop <- struct{}{}
		}
	}, nil
}
