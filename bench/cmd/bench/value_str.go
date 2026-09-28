//go:build strvals

package main

import (
	"unsafe"
)

// V is string in a bench built with the tag strvals: values that hold a
// pointer, as record IDs or names do, which Ordered keeps in set leaves
// rather than flat leaves (see value_u64.go).
type V = string

// valueTag marks results of this build: its value profiles are reported as
// multi-str and unique-str.
const valueTag = "-str"

// toVs returns the values for the value numbers u: each number as 16 hex
// digits, as a record ID. All of them are views into one buffer, so that
// they cost the garbage collector one object and every candidate holds the
// same string headers.
func toVs(u []uint64) []V {
	const hex = "0123456789abcdef"
	buf := make([]byte, 16*len(u))
	for i, x := range u {
		for j := 15; j >= 0; j-- {
			buf[16*i+j] = hex[x&0xf]
			x >>= 4
		}
	}
	s := unsafe.String(unsafe.SliceData(buf), len(buf))
	out := make([]V, len(u))
	for i := range out {
		out[i] = s[16*i : 16*i+16]
	}
	return out
}

// weigh returns what the timed loops and the checks add up per value: the
// address of its bytes. It reads only the string header, which the
// candidate holds, not the bytes behind it, and since every candidate holds
// the same headers (see toVs), the sums of two candidates agree exactly when
// they hold the same values.
func weigh(v V) uint64 { return uint64(uintptr(unsafe.Pointer(unsafe.StringData(v)))) }
