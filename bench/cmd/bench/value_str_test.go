//go:build strvals

package main

import (
	"strconv"
	"testing"
)

// TestStringValues makes sure a bench built with string values checks the
// candidates as strictly as one with integers. It covers the value type of
// the strvals build (value_str.go): every value number becomes its own 16
// hex digits, and weigh tells apart two strings with equal bytes, because
// the checks compare which values a candidate holds, not only how many.
func TestStringValues(t *testing.T) {
	nums := []uint64{0, 1, 0xdeadbeef, 1<<63 | 42}
	vs := toVs(nums)
	seen := map[uint64]bool{}
	for i, v := range vs {
		if got, err := strconv.ParseUint(v, 16, 64); err != nil || got != nums[i] || len(v) != 16 {
			t.Fatalf("value %d is %q, want %d as 16 hex digits", i, v, nums[i])
		}
		if seen[weigh(v)] {
			t.Fatalf("value %q weighs like another", v)
		}
		seen[weigh(v)] = true
	}
	if again := toVs(nums[:1]); again[0] == vs[0] && weigh(again[0]) == weigh(vs[0]) {
		t.Fatal("a copy of a value weighs like the value itself")
	}
}
