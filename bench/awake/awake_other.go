//go:build !darwin && !linux && !windows

package awake

import (
	"fmt"
	"runtime"
)

func hold(string) (func(), error) {
	return nil, fmt.Errorf("cannot keep the machine awake on %s", runtime.GOOS)
}
