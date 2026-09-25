package awake

import (
	"os"
	"strconv"
)

// hold runs caffeinate -i (no idle system sleep); -w ends it with this
// process even if release is never called.
func hold(string) (func(), error) {
	return holdCommand("caffeinate", "-i", "-w", strconv.Itoa(os.Getpid()))
}
