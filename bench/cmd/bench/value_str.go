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
// natural-str and single-value-str.
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

// weigh returns what the timed loops add up per value: its length and its
// first and last byte. It reads the bytes of the string, as a caller that uses
// a value does, and so pays the cache miss that a candidate with a pointer to
// the caller's string pays there and a candidate that holds the bytes in its
// own pages does not. The checks compare values with checkWeigh instead.
func weigh(v V) uint64 {
	if len(v) == 0 {
		return 0
	}
	return uint64(len(v)) + uint64(v[0]) + uint64(v[len(v)-1])
}

// checkWeigh is weigh for the checks that candidates hold the same values: a
// hash of all bytes, since two values may agree in length and end bytes. The
// timed loops do not use it; it costs more than a caller's use of a value.
func checkWeigh(v V) uint64 {
	h := uint64(14695981039346656037)
	for i := range len(v) {
		h = (h ^ uint64(v[i])) * 1099511628211
	}
	return h ^ uint64(len(v))
}

// valueBytes returns the bytes of the string values vals: what a candidate
// that owns its values holds besides the headers.
func valueBytes(vals []V) int {
	n := 0
	for _, v := range vals {
		n += len(v)
	}
	return n
}
