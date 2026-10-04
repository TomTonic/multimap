//go:build strvals

package main

import (
	"strconv"
	"testing"
	"unsafe"

	"github.com/TomTonic/multimap/bench/keys"
)

// TestStringValues makes sure a bench built with string values checks the
// candidates as strictly as one with integers. It covers the value type of
// the strvals build (value_str.go): every value number becomes its own 16
// hex digits, and every value number is its own string, because the checks
// compare which values a candidate holds, not only how many.
func TestStringValues(t *testing.T) {
	nums := []uint64{0, 1, 0xdeadbeef, 1<<63 | 42}
	vs := toVs(nums, nil)
	seen := map[uint64]bool{}
	for i, v := range vs {
		if got, err := strconv.ParseUint(v, 16, 64); err != nil || got != nums[i] || len(v) != 16 {
			t.Fatalf("value %d is %q, want %d as 16 hex digits", i, v, nums[i])
		}
		if seen[checkWeigh(v)] {
			t.Fatalf("value %q checks like another", v)
		}
		seen[checkWeigh(v)] = true
	}
}

// TestNamedValues makes sure that a bench built with string values gives the
// corpora that have them real strings. It covers toNamed (value_str.go): a
// value number of a corpus with names becomes its name, of its own length,
// the same number is the same string at the same address, a number beyond the
// names falls back to 16 hex digits, and street and directory corpora have
// names of different lengths.
func TestNamedValues(t *testing.T) {
	names := []string{"a", "bb", "dddd"}
	vs := toVs([]uint64{3, 1, 3, 2, 99}, names)
	want := []string{"dddd", "a", "dddd", "bb", "0000000000000063"}
	for i, v := range vs {
		if v != want[i] {
			t.Errorf("value %d is %q, want %q", i, v, want[i])
		}
	}
	if addr(vs[0]) != addr(vs[2]) || addr(vs[0]) == addr(vs[1]) {
		t.Error("equal numbers must be the same string, different numbers different ones")
	}
	for _, kind := range []keys.Kind{keys.Street, keys.Dirs} {
		c := keys.Generate(kind, 2000, 3)
		lengths := map[int]bool{}
		for _, v := range toVs(c.Natural[0], c.Names) {
			lengths[len(v)] = true
		}
		for _, n := range c.Names[:200] {
			lengths[len(n)] = true
		}
		if len(lengths) < 5 {
			t.Errorf("%s: names of only %d different lengths", kind, len(lengths))
		}
	}
}

// addr returns the address of the bytes of v.
func addr(v V) uintptr { return uintptr(unsafe.Pointer(unsafe.StringData(v))) }

// TestWeighReadsBytes makes sure that the timed loops of a bench with string
// values depend on the bytes of a value, as a caller's use of it does: values
// of other length or other end bytes weigh differently, an empty one weighs 0.
func TestWeighReadsBytes(t *testing.T) {
	if weigh("") != 0 {
		t.Error("an empty value must weigh 0")
	}
	if weigh("ab") == weigh("abc") || weigh("ab") == weigh("bb") || weigh("ab") == weigh("aa") {
		t.Error("length, first and last byte must all show in the weight")
	}
	if weigh("ab") != weigh(string([]byte("ab"))) {
		t.Error("a copy of a value must weigh like the value")
	}
}
