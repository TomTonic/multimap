// Package rtopt holds the rtcompare settings that every comparison command
// shares, so that a whole suite can be rerun with longer batches or more
// samples through the same two flags.
package rtopt

import (
	"flag"
	"fmt"
	"os"

	"github.com/TomTonic/rtcompare"
)

// maxQuantizationError is ten times stricter than rtcompare's default: a
// smoke test at the default showed an 18% tie rate.
const maxQuantizationError = 0.0001

var (
	loopScale = flag.Float64("loopscale", 1, "multiply the calibrated operations per batch by this factor")
	repeats   = flag.Int("repeats", 0, "timing samples per candidate (0: rtcompare's default)")
)

// Options returns the CompareOptions for comparing a with b under the
// -loopscale and -repeats flags. Call it after flag.Parse, once per pair:
// with -loopscale it calibrates both candidates the way rtcompare.Compare
// does, takes the larger batch size and scales it. gcBetween collects garbage
// between batches, for operations that allocate a whole structure.
func Options(a, b rtcompare.Candidate, gcBetween bool) rtcompare.CompareOptions {
	c := rtcompare.CollectOptions{MaxQuantizationError: maxQuantizationError, Repeats: *repeats, GCBetween: gcBetween}
	if *loopScale != 1 {
		c.InnerLoops = scaledLoops(a, b, c, *loopScale)
	}
	return rtcompare.CompareOptions{Collect: c}
}

// scaledLoops exists because rtcompare calibrates inside Compare and offers no
// scale factor; longer batches average more cache and scheduler noise into
// every sample.
func scaledLoops(a, b rtcompare.Candidate, c rtcompare.CollectOptions, scale float64) uint64 {
	opt := rtcompare.CalibrationOptions{MaxQuantizationError: c.MaxQuantizationError, GCBetween: c.GCBetween}
	var loops uint64
	for _, x := range []rtcompare.Candidate{a, b} {
		cal, err := rtcompare.CalibrateInnerLoops(x, opt)
		if err != nil {
			fmt.Fprintln(os.Stderr, "calibration:", err)
			os.Exit(1)
		}
		loops = max(loops, cal.InnerLoops)
	}
	return uint64(float64(loops)*scale + 0.5)
}
