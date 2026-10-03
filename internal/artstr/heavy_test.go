package artstr

import (
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	if s := os.Getenv("ARTSTR_CROWDED"); s != "" {
		CrowdedRatio, _ = strconv.Atoi(s)
	}
	if s := os.Getenv("ARTSTR_SEED"); s != "" {
		seedBase, _ = strconv.ParseUint(s, 10, 64)
		mapSteps = 30000
	}
	os.Exit(m.Run())
}
