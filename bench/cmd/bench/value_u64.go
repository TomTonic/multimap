//go:build !strvals

package main

// V is the value type of every candidate: uint64, or string in a bench built
// with the tag strvals (see value_str.go). Both builds share one source, so
// that the uint64 build compiles exactly as before.
type V = uint64

// valueTag marks results of this build: empty for uint64 values.
const valueTag = ""

// toVs returns the values for the value numbers u; for uint64 they are the
// numbers themselves.
func toVs(u []uint64) []V { return u }

// weigh returns what the timed loops and the checks add up per value: for
// uint64, the value itself.
func weigh(v V) uint64 { return v }
