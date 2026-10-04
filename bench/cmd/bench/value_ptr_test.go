//go:build ptrvals && !strvals

package main

import "testing"

// TestPointerValues makes sure a bench built with pointer values checks the
// candidates as strictly as one with integers. It covers toVs of the ptrvals
// build (value_ptr.go): equal value numbers are the same record, different
// numbers different ones, and the checks tell the records apart by what they
// hold, not by their address.
func TestPointerValues(t *testing.T) {
	vs := toVs([]uint64{5, 7, 5, 0}, nil)
	if vs[0] != vs[2] || vs[0] == vs[1] || vs[3] == nil {
		t.Fatal("equal numbers must be the same record, different numbers different ones")
	}
	if weigh(vs[0]) != 5 || checkWeigh(vs[0]) == checkWeigh(vs[1]) {
		t.Fatalf("weigh = %d, want 5; the checks must tell two records apart", weigh(vs[0]))
	}
}
