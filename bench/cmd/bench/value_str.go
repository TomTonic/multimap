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
// digits, as a record ID, or, if names is not nil and the number is one of
// them (number v is names[v-1]), that name: a real string of its own length,
// the locality of a street or the name of a file. All of them are views into
// one buffer, so that they cost the garbage collector one object and every
// candidate holds the same string headers, and one number is always the same
// string.
func toVs(u []uint64, names []string) []V {
	if names != nil {
		return toNamed(u, names)
	}
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

// toNamed is toVs for a corpus with names; numbers beyond the names (the
// transient values of the index workloads never get here, see newPairs) are
// 16 hex digits as without names.
func toNamed(u []uint64, names []string) []V {
	offs := make([]int, len(names)+1)
	for i, n := range names {
		offs[i+1] = offs[i] + len(n)
	}
	buf := make([]byte, 0, offs[len(names)])
	for _, n := range names {
		buf = append(buf, n...)
	}
	s := unsafe.String(unsafe.SliceData(buf), len(buf))
	var rest []uint64 // numbers without a name
	out := make([]V, len(u))
	for i, x := range u {
		if x >= 1 && x <= uint64(len(names)) {
			out[i] = s[offs[x-1]:offs[x]]
		} else {
			rest = append(rest, x)
			out[i] = ""
		}
	}
	if len(rest) > 0 {
		hex := toVs(rest, nil)
		for i, x := range u {
			if x < 1 || x > uint64(len(names)) {
				out[i], hex = hex[0], hex[1:]
			}
		}
	}
	return out
}

// weigh returns what the timed loops and the checks add up per value: the
// address of its bytes. It reads only the string header, which the
// candidate holds, not the bytes behind it, and since every candidate holds
// the same headers (see toVs), the sums of two candidates agree exactly when
// they hold the same values.
func weigh(v V) uint64 { return uint64(uintptr(unsafe.Pointer(unsafe.StringData(v)))) }
