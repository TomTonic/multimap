package main

import "github.com/TomTonic/multimap/bench/keys"

// baseline is multimap.Ordered as of another commit, which cmd/mkbaseline
// copies into package baseline. It is a candidate only in a bench built with
// the tag baseline (go run -tags baseline ./cmd/bench), so that Ordered can
// be compared head to head with an earlier version of itself.
const baseline = "baseline"

// kit drives a candidate whose type only a file with a build tag knows. Its
// functions take the candidate as any and assert its concrete type once,
// outside the timed loops, which then call it directly, as the other
// candidates' loops do.
type kit struct {
	ref       string // the commit the candidate was copied from
	build     func(k [][]byte, vals []uint64, offs []int) any
	empty     func() any
	apply     func(m any) func(k [][]byte, ms []mutation)
	valuesFor func(m any, p keys.Set) func(uint64)
	between   func(m any, from, to keys.Set) func(uint64)
	sum       func(m any, key []byte) uint64
	rangeSum  func(m any, from, to []byte) uint64
	keys      func(m any) int
	removeKey func(m any) func(key []byte)
}

// baseKit is the kit of the baseline candidate, or nil in a bench built
// without the tag baseline.
var baseKit *kit
