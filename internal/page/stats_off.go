//go:build !mkstats

package page

// countLay counts a call of head.lay in a build with the tag mkstats (stats_on.go); here it costs nothing.
func countLay() {}

// LayCalls returns how often head.lay ran; 0 in a build without the tag mkstats.
func LayCalls() uint64 { return 0 }
