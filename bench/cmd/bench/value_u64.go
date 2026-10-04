//go:build !strvals

package main

// V is the value type of every candidate: uint64, or string in a bench built
// with the tag strvals (see value_str.go). Both builds share one source, so
// that the uint64 build compiles exactly as before.
type V = uint64

// valueTag marks results of this build: empty for uint64 values.
const valueTag = ""

// toVs returns the values for the value numbers u; for uint64 they are the
// numbers themselves, and the names that some corpora give them (see
// keys.Corpus.Names) are not used.
func toVs(u []uint64, _ []string) []V { return u }

// weigh returns what the timed loops and the checks add up per value: for
// uint64, the value itself.
func weigh(v V) uint64 { return v }

// checkWeigh is weigh for the checks that candidates hold the same values:
// the value itself.
func checkWeigh(v V) uint64 { return v }

// valueBytes returns the bytes of the string values vals; none for uint64.
func valueBytes(_ []V) int { return 0 }
