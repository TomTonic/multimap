//go:build mkstats

package page

// layCalls counts the calls of head.lay in a build with the tag mkstats (all pages of the process, not safe for
// concurrent use), so that a probe can tell how often a write computes the layout of a page
// (docs/redesign/write-path-analysis.md).
var layCalls uint64

// countLay counts one call of head.lay.
func countLay() { layCalls++ }

// LayCalls returns how often head.lay ran since the start of the process.
func LayCalls() uint64 { return layCalls }
