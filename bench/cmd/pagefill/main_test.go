package main

import (
	"bytes"
	"strings"
	"testing"
)

// TestRun makes sure that the page fill experiment reports a row for every key
// kind it is asked for, with the figures the redesign's step 1 needs. It
// belongs to the page prototype (docs/redesign, PLAN step 1): integers must
// come out as uniform pages of 16 bytes a key at about 70% fill, strings as
// general pages, and a typo must be reported.
func TestRun(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, []string{"-keys", "u64,uuid", "-n", "20000"}); err != nil {
		t.Fatal(err)
	}
	rows := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(rows) != 6 || !strings.HasPrefix(rows[4], "| u64 | 20000 |") || !strings.HasPrefix(rows[5], "| uuid | 20000 |") {
		t.Fatalf("rows:\n%s", out.String())
	}
	if c := strings.Split(rows[4], " | "); c[6] != "100 %" {
		t.Errorf("integer keys: %s of the pages are uniform, want all", c[6])
	}
	if c := strings.Split(rows[5], " | "); c[6] != "0 %" {
		t.Errorf("uuids: %s of the pages are uniform, want none", c[6])
	}
	if err := run(&out, []string{"-keys", "nope"}); err == nil {
		t.Error("an unknown key kind was accepted")
	}
	if err := run(&out, []string{"-nope"}); err == nil {
		t.Error("an unknown flag was accepted")
	}
}

// TestHelpers makes sure that the pieces of the page fill experiment measure
// what their names say. It belongs to the page prototype (docs/redesign, PLAN
// step 1).
func TestHelpers(t *testing.T) {
	if got := lcp([]byte("abcd"), []byte("abxy")); got != 2 {
		t.Errorf("lcp = %d, want 2", got)
	}
	if got := lcp([]byte("ab"), []byte("abxy")); got != 2 {
		t.Errorf("lcp of a prefix = %d, want 2", got)
	}
	for pages, want := range map[int]int{1: 128, 8: 128, 9: 256, 24: 256, 25: 512, 56: 512, 57: 2112} {
		if got := routerBytes(pages); got != want {
			t.Errorf("routerBytes(%d) = %d, want %d", pages, got, want)
		}
	}
}
